package strings

// Nodes are managed by slice(vector).

const (
	// assume words consisting of only lower case or upper case
	_TRIE_CHAR_SIZE = 26
)

// Operate do something in a trie node while traversing automaton.
type Operate func(curNodeID int, c int)

// NewTrie returns a trie managing words starting from base character.
// e.g.: NewTrie('a')
func NewTrie(base rune) *Trie {
	t := new(Trie)

	t.root = 0
	t.nodes = append(t.nodes, newTrieNode(t.root))
	t.base = base
	t.noop = func(curNodeID int, c int) {}

	return t
}

// Insert a word.
func (t *Trie) Insert(word string) {
	t._insert(word, t.nodes[0].common)
}

// Find returns whether the trie has the word or not.
func (t *Trie) Find(word string) bool {
	return t._search(word, false, t.noop)
}

// FindStartWith returns whether the trie has the word having the prefix or not.
func (t *Trie) FindStartWith(prefix string) bool {
	return t._search(prefix, true, t.noop)
}

// Traverse walk on the tree while operating something.
// Operation function receives current node id and current node offset from a base.
func (t *Trie) Traverse(word string, op Operate) {
	t._search(word, false, op)
}

// CountWord returns the number of the words that the trie has.
// CountWord can NOT count the number of UNIQUE words.
func (t *Trie) CountWord() int {
	return t.nodes[0].common
}

// SizeTrie returns the number of the nodes that the trie has.
func (t *Trie) SizeTrie() int {
	return len(t.nodes)
}

// IsAccept returns whether a trie node says accept or not.
func (t *Trie) IsAccept(curNodeID int) bool {
	return len(t.nodes[curNodeID].accept) > 0
}

type Trie struct {
	nodes []*trieNode // nodes managed by the trie
	root  int         // root node id
	base  rune        // base character
	noop  Operate     // do nothing
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
	curNodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := &t.nodes[curNodeID].next[c]
		if *nextID == -1 {
			// add nodes when there is not the next node
			*nextID = len(t.nodes)
			t.nodes = append(t.nodes, newTrieNode(c))
		}
		t.nodes[curNodeID].common++
		curNodeID = *nextID
	}
	t.nodes[curNodeID].common++
	t.nodes[curNodeID].accept = append(t.nodes[curNodeID].accept, wordID)
}

func (t *Trie) _search(word string, isPrefix bool, op Operate) bool {
	curNodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		// operate something (do nothing if op is noop)
		op(curNodeID, c)

		nextNodeID := t.nodes[curNodeID].next[c]
		if nextNodeID == -1 {
			return false
		}
		curNodeID = nextNodeID
	}

	if isPrefix {
		return true
	}

	// check whether the word is accepted or not
	return t.IsAccept(curNodeID)
}
