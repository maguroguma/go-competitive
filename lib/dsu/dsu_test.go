package dsu

import (
	"reflect"
	"testing"
)

func TestDisjointSetUnion(t *testing.T) {
	d := NewDisjointSetUnion(4)
	d.Merge(0, 1)
	if !d.Same(0, 1) {
		t.FailNow()
	}
	d.Merge(1, 2)
	if !d.Same(0, 2) {
		t.FailNow()
	}
	if !d.Same(1, 2) {
		t.FailNow()
	}
	if !(d.Size(0) == 3) {
		t.Fatal(d.Size(0))
	}
	if d.Same(0, 3) {
		t.FailNow()
	}
	if g := d.Groups(); !reflect.DeepEqual(g, [][]int{{0, 1, 2}, {3}}) {
		t.Fatal(g)
	}
}

func TestGroups(t *testing.T) {
	d := NewDisjointSetUnion(9)
	d.Merge(2, 1)
	d.Merge(0, 1)
	d.Merge(3, 5)
	d.Merge(4, 5)
	d.Merge(6, 7)
	d.Merge(6, 8)

	expected := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	actual := d.Groups()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
