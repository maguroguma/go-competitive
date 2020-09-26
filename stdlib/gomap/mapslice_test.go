package gomap

import (
	"reflect"
	"sort"
	"testing"
)

// TestMapSliceValue: map[string][]int というデータ構造は正しい
func TestMapSliceValue(t *testing.T) {
	objects := []struct {
		name  string
		value int
	}{
		{"BBB", 5}, {"BBB", 4}, {"BBB", 3}, {"BBB", 2}, {"BBB", 1},
		{"CCC", 5}, {"CCC", 4}, {"CCC", 3}, {"CCC", 2}, {"CCC", 1},
		{"AAA", 5}, {"AAA", 4}, {"AAA", 3}, {"AAA", 2}, {"AAA", 1},
	}

	expected := []int{1, 2, 3, 4, 5}

	memo := make(map[string][]int)
	for _, v := range objects {
		memo[v.name] = append(memo[v.name], v.value)
	}

	for _, S := range memo {
		sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })
	}

	keys := []string{"AAA", "BBB", "CCC"}

	for _, key := range keys {
		actual := memo[key]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, want %v", actual, expected)
		}
	}
}
