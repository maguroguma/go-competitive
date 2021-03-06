package strings

const (
	_PMA_CHAR_SIZE = 26
)

type PMA struct {
	root *pmaNode
	base rune
}

type pmaNode struct {
	next    [_PMA_CHAR_SIZE]*pmaNode
	accept  []int
	failure *pmaNode
}

// NewPMA returns Pattern Matching Automaton by pattern strings (dictionary).
func NewPMA(patterns [][]rune, base rune) *PMA {
	root := new(pmaNode)
	root.failure = root

	pma := new(PMA)
	pma.root = root
	pma.base = base

	// build Trie by pattern strings
	var curNode *pmaNode
	for i := 0; i < len(patterns); i++ {
		curNode = root
		for _, r := range patterns[i] {
			c := r - base
			if curNode.next[c] == nil {
				curNode.next[c] = new(pmaNode)
			}
			curNode = curNode.next[c]
		}
		curNode.accept = append(curNode.accept, i)
	}

	// Aho-Corasick method
	que := []*pmaNode{}
	// Initialize queue
	for i := 0; i < _PMA_CHAR_SIZE; i++ {
		if root.next[i] == nil {
			root.next[i] = root
		} else {
			root.next[i].failure = root
			que = append(que, root.next[i])
		}
	}
	// BFS
	for len(que) > 0 {
		curNode = que[0]
		que = que[1:]

		for i := 0; i < _PMA_CHAR_SIZE; i++ {
			if curNode.next[i] == nil {
				continue
			}

			beforeFailNode := curNode.failure
			for beforeFailNode.next[i] == nil {
				beforeFailNode = beforeFailNode.failure
			}

			curNode.next[i].failure = beforeFailNode.next[i]
			curNode.next[i].accept = pma._setUnion(curNode.next[i].accept, beforeFailNode.next[i].accept)

			que = append(que, curNode.next[i])
		}
	}

	return pma
}

// Match returns all matched patterns in the text.
// res[i] has word indices that matches the last character in i index (0-index).
// res[i] can have multiple word indices, and in that case,
//  its order is not deterministic.
// res has shallow copy of accept slices, so you should not update the res.
func (pma *PMA) Match(text []rune) [][]int {
	res := make([][]int, len(text))

	curNode := pma.root
	for i, r := range text {
		c := r - pma.base

		// use the failure link
		for curNode.next[c] == nil {
			curNode = curNode.failure // like an epsilon transition
		}
		curNode = curNode.next[c] // consume a character

		// check pattern match
		res[i] = curNode.accept
	}

	return res
}

func (pma *PMA) _setUnion(A, B []int) []int {
	res := []int{}

	memo := map[int]bool{}
	for _, a := range A {
		memo[a] = true
	}
	for _, b := range B {
		memo[b] = true
	}

	for k := range memo {
		res = append(res, k)
	}

	return res
}
