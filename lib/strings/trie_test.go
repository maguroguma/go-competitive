package strings

import (
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie('a')

	trie.Insert("fire")
	if !trie.Find("fire") {
		t.Errorf("got %v, want %v", false, true)
	}
	if !trie.FindStartWith("fi") {
		t.Errorf("got %v, want %v", false, true)
	}

	trie.Insert("fireman")
	trie.Insert("firearm")
	trie.Insert("firework")
	trie.Insert("algo")

	trie.Insert("fireman")
	trie.Insert("firearm")
	trie.Insert("firework")
	trie.Insert("algo")

	println(trie.CountWord())
	println(trie.SizeTrie())

	if trie.CountWord() != 9 {
		t.Errorf("got %v, want %v", trie.CountWord(), 9)
	}
	if trie.SizeTrie() != 18+1 {
		t.Errorf("got %v, want %v", trie.SizeTrie(), 18+1)
	}
}

// https://atcoder.jp/contests/agc047/tasks/agc047_b
func TestTraverse(t *testing.T) {
	n := 6
	T := []string{"b", "a", "abc", "c", "d", "ab"}

	S := []string{}
	for i := 0; i < n; i++ {
		rev := []rune{}
		org := []rune(T[i])
		for j := len(org) - 1; j >= 0; j-- {
			rev = append(rev, org[j])
		}

		S = append(S, string(rev))
	}

	for _, str := range S {
		println(str)
	}

	sort.Slice(S, func(i, j int) bool {
		return len(S[i]) < len(S[j])
	})

	trie := NewTrie('a')

	ans := 0
	for i := 0; i < n; i++ {
		memo := [26]int{}
		for _, r := range S[i] {
			memo[r-'a']++
		}

		op := func(cid int, c int) {
			for i := 0; i < _TRIE_CHAR_SIZE; i++ {
				nid := trie.nodes[cid].next[i]
				if nid != -1 && trie.IsAccept(nid) && memo[i] > 0 {
					ans++
				}
			}

			memo[c]--
		}
		trie.Traverse(S[i][:len(S[i])-1], op)

		trie.Insert(S[i])
	}

	if ans != 5 {
		t.Errorf("got %v, want %v", ans, 5)
	}
}
