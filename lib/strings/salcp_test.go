package strings

import (
	"reflect"
	"testing"
)

func TestSuffixArrayString(t *testing.T) {
	text := "missisippi"
	sa := SuffixArrayString(text)

	expected := []string{
		"i",
		"ippi",
		"isippi",
		"issisippi",
		"missisippi",
		"pi",
		"ppi",
		"sippi",
		"sisippi",
		"ssisippi",
	}

	actual := []string{}
	for _, idx := range sa {
		actual = append(actual, text[idx:])
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestPatternMatchBySuffixArrayOnly(t *testing.T) {
	text := "missisippi"
	tsa := SuffixArrayString(text)

	pattern := "is"

	exLeft, exRight := 2, 3
	acLeft, acRight := MatchBySA(text, tsa, pattern)

	if !(exLeft == acLeft && exRight == acRight) {
		t.Errorf("got (%v, %v), want (%v, %v)", acLeft, acRight, exLeft, exRight)
	}
}
