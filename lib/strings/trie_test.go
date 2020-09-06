package strings

import "testing"

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
}
