package prime

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var (
	osa *Osa_kSieve
)

const (
	MAX_N = 1 << uint(20)
)

func TestMain(m *testing.M) {
	println("before all...")

	osa = NewOsa_kSieve(MAX_N)
	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestFactors(t *testing.T) {
	testcases := []struct {
		a        int
		expected []int
	}{
		{a: 120, expected: []int{2, 2, 2, 3, 5}},
		{a: 81, expected: []int{3, 3, 3, 3}},
		{a: 121, expected: []int{11, 11}},
		{a: 13, expected: []int{13}},
		{a: 2, expected: []int{2}},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d testcases", i)
		t.Run(subTitle, func(t *testing.T) {
			actual := osa.Factors(tc.a)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	testcases := []struct {
		a        int
		expected bool
	}{
		{a: 2, expected: true},
		{a: 13, expected: true},
		{a: 81, expected: false},
		{a: 121, expected: false},
		{a: 120, expected: false},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d testcases", i)
		t.Run(subTitle, func(t *testing.T) {
			actual := osa.IsPrime(tc.a)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}
