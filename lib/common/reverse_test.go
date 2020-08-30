package common

import (
	"reflect"
	"testing"
)

func TestReverseMyself(t *testing.T) {
	testcases := []struct {
		title    string
		arg, res []interface{}
	}{
		{
			"int",
			[]interface{}{0, 1, 2, 3, 4},
			[]interface{}{4, 3, 2, 1, 0},
		},
		{
			"float64",
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4},
			[]interface{}{4.4, 3.3, 2.2, 1.1, 0.0},
		},
		{
			"rune",
			[]interface{}{'a', 'b', 'c', 'd', 'e'},
			[]interface{}{'e', 'd', 'c', 'b', 'a'},
		},
		{
			"string",
			[]interface{}{"aaa", "bbb", "ccc", "ddd", "eee"},
			[]interface{}{"eee", "ddd", "ccc", "bbb", "aaa"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			ReverseMyself(tc.arg)
			if !reflect.DeepEqual(tc.arg, tc.res) {
				t.Errorf("got %v, want %v", tc.arg, tc.res)
			}
		})
	}
}

func TestRotateMyself(t *testing.T) {
	testcases := []struct {
		title    string
		times    int
		arg, res []interface{}
	}{
		{
			"int",
			2,
			[]interface{}{0, 1, 2, 3, 4},
			[]interface{}{2, 3, 4, 0, 1},
		},
		{
			"int full",
			5,
			[]interface{}{0, 1, 2, 3, 4},
			[]interface{}{0, 1, 2, 3, 4},
		},
		{
			"float64",
			2,
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4},
			[]interface{}{2.2, 3.3, 4.4, 0.0, 1.1},
		},
		{
			"rune",
			2,
			[]interface{}{'a', 'b', 'c', 'd', 'e'},
			[]interface{}{'c', 'd', 'e', 'a', 'b'},
		},
		{
			"string",
			2,
			[]interface{}{"aaa", "bbb", "ccc", "ddd", "eee"},
			[]interface{}{"ccc", "ddd", "eee", "aaa", "bbb"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			RotateMyself(tc.arg, tc.times)
			if !reflect.DeepEqual(tc.arg, tc.res) {
				t.Errorf("got %v, want %v", tc.arg, tc.res)
			}
		})
	}
}
