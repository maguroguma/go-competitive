package csum

import (
	"reflect"
	"testing"
)

func TestImos(t *testing.T) {
	actual := []int64{0, 1, 3, 3, 3, 2, 2}

	maxT := 6
	im := NewImos(maxT)

	im.AddEvent(1, 1)
	im.AddEvent(2, 2)
	im.AddEvent(5, -1)

	expected := im.Build()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
