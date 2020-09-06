package strings

const (
	// assume words consisting of only lower case or upper case
	_TRIE_CHAR_SIZE = 26
)

// NewTrie returns a trie managing words starting from base character.
// Nodes are managed by slice(vector).
// e.g.: NewTrie('a')
func NewTrie(base rune) *Trie {
	t := new(Trie)

	t.root = 0
	t.nodes = append(t.nodes, newTrieNode(t.root))
	t.base = base

	return t
}

// Insert a word.
func (t *Trie) Insert(word string) {
	t._insert(word, t.nodes[0].common)
}

// Find returns whether the trie has the word or not.
func (t *Trie) Find(word string) bool {
	return t._search(word, false)
}

// FindStartWith returns whether the trie has the word having the prefix or not.
func (t *Trie) FindStartWith(prefix string) bool {
	return t._search(prefix, true)
}

// CountWord returns the number of the words that the trie has.
// CountWord cannot count the number of UNIQUE words.
func (t *Trie) CountWord() int {
	return t.nodes[0].common
}

// SizeTrie returns the number of the nodes that the trie has.
func (t *Trie) SizeTrie() int {
	return len(t.nodes)
}

type Trie struct {
	nodes []*trieNode // nodes managed by the trie
	root  int         // root node id
	base  rune        // base character
}

type trieNode struct {
	next   [_TRIE_CHAR_SIZE]int // child id that a character ith has (NIL == -1)
	accept []int                // accept has string id(s) whose last character is equal to this node
	c      int                  // offset of this node's character from the base character
	common int                  // number of strings that share this node
}

func newTrieNode(c int) *trieNode {
	tn := new(trieNode)

	tn.c = c
	tn.common = 0
	for i := 0; i < _TRIE_CHAR_SIZE; i++ {
		tn.next[i] = -1
	}

	return tn
}

func (t *Trie) _insert(word string, wordID int) {
	nodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := &t.nodes[nodeID].next[c]
		if *nextID == -1 {
			// add nodes when there is not the next node
			*nextID = len(t.nodes)
			t.nodes = append(t.nodes, newTrieNode(c))
		}
		t.nodes[nodeID].common++
		nodeID = *nextID
	}
	t.nodes[nodeID].common++
	t.nodes[nodeID].accept = append(t.nodes[nodeID].accept, wordID)
}

func (t *Trie) _search(word string, isPrefix bool) bool {
	nodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := t.nodes[nodeID].next[c]
		if nextID == -1 {
			return false
		}
		nodeID = nextID
	}

	if isPrefix {
		return true
	}

	// check whether the word is accepted or not
	return len(t.nodes[nodeID].accept) > 0
}
