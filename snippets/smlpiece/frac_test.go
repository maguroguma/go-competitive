package smlpiece

import (
	"fmt"
	"testing"
)

type bunsu struct {
	bunshi, bunbo int
}

func TestFloorFrac(t *testing.T) {
	testcases := []struct {
		tname    string
		input    bunsu
		expected int
	}{
		{
			tname:    "1/3 => 0",
			input:    bunsu{bunshi: 1, bunbo: 3},
			expected: 0,
		},
		{
			tname:    "1/-3 => -1",
			input:    bunsu{bunshi: 1, bunbo: -3},
			expected: -1,
		},
		{
			tname:    "-1/3 => -1",
			input:    bunsu{bunshi: -1, bunbo: 3},
			expected: -1,
		},
		{
			tname:    "-1/-3 => 0",
			input:    bunsu{bunshi: -1, bunbo: -3},
			expected: 0,
		},
		{
			tname:    "6/3 => 2",
			input:    bunsu{bunshi: 6, bunbo: 3},
			expected: 2,
		},
		{
			tname:    "6/-3 => -2",
			input:    bunsu{bunshi: 6, bunbo: -3},
			expected: -2,
		},
		{
			tname:    "-6/3 => -2",
			input:    bunsu{bunshi: -6, bunbo: 3},
			expected: -2,
		},
		{
			tname:    "-6/-3 => 2",
			input:    bunsu{bunshi: -6, bunbo: -3},
			expected: 2,
		},
		{
			tname:    "0/3 => 0",
			input:    bunsu{bunshi: 0, bunbo: 3},
			expected: 0,
		},
		{
			tname:    "0/-3 => 0",
			input:    bunsu{bunshi: 0, bunbo: -3},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("No.%2d: %s", i, tc.tname), func(t *testing.T) {
			actual := FloorFrac(tc.input.bunshi, tc.input.bunbo)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestCeilFrac(t *testing.T) {
	testcases := []struct {
		tname    string
		input    bunsu
		expected int
	}{
		{
			tname:    "1/3 => 1",
			input:    bunsu{bunshi: 1, bunbo: 3},
			expected: 1,
		},
		{
			tname:    "1/-3 => 0",
			input:    bunsu{bunshi: 1, bunbo: -3},
			expected: 0,
		},
		{
			tname:    "-1/3 => 0",
			input:    bunsu{bunshi: -1, bunbo: 3},
			expected: 0,
		},
		{
			tname:    "-1/-3 => 1",
			input:    bunsu{bunshi: -1, bunbo: -3},
			expected: 1,
		},
		{
			tname:    "6/3 => 2",
			input:    bunsu{bunshi: 6, bunbo: 3},
			expected: 2,
		},
		{
			tname:    "6/-3 => -2",
			input:    bunsu{bunshi: 6, bunbo: -3},
			expected: -2,
		},
		{
			tname:    "-6/3 => -2",
			input:    bunsu{bunshi: -6, bunbo: 3},
			expected: -2,
		},
		{
			tname:    "-6/-3 => 2",
			input:    bunsu{bunshi: -6, bunbo: -3},
			expected: 2,
		},
		{
			tname:    "0/3 => 0",
			input:    bunsu{bunshi: 0, bunbo: 3},
			expected: 0,
		},
		{
			tname:    "0/-3 => 0",
			input:    bunsu{bunshi: 0, bunbo: -3},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("No.%2d: %s", i, tc.tname), func(t *testing.T) {
			actual := CeilFrac(tc.input.bunshi, tc.input.bunbo)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}
