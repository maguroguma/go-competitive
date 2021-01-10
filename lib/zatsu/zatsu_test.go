package zatsu

import "testing"

func TestCompress(t *testing.T) {
	A := []int64{1, 10, 100, 1000, 10000}

	comp := NewCompress()
	comp.Add(A...)
	comp.Build()

	for i, a := range A {
		actual := comp.Get(a)
		if actual != int64(i) {
			t.Errorf("got %v, want %v", actual, i)
		}
		actual = comp.InvGet(int64(i))
		if actual != int64(a) {
			t.Errorf("got %v, want %v", actual, a)
		}
	}
}

func TestCompressKind(t *testing.T) {
	A := []int64{1, 10, 100, 1, 10, 100}
	expected := 3

	comp := NewCompress()
	comp.Add(A...)
	comp.Build()

	actual := comp.Kind()
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

// https://atcoder.jp/contests/abc036/tasks/abc036_c
func TestABC036C(t *testing.T) {
	A := []int64{3, 3, 1, 6, 1}
	expected := []int64{1, 1, 0, 2, 0}

	comp := NewCompress()
	comp.Add(A...)
	comp.Build()

	for i, a := range A {
		actual := comp.Get(a)
		if actual != expected[i] {
			t.Errorf("got %v, want %v", actual, expected[i])
		}
	}
}
