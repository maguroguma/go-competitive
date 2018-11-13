package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSliceSort(t *testing.T) {
	actual := sortIntSlice([]int{3, 1, 5, 4, 2, 5})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 5}, actual)
}

func TestStringSliceSort(t *testing.T) {
	actual := sortStringSlice([]string{"bubble", "basic", "actual", "zsh", "example"})
	assert.Equal(t, []string{"actual", "basic", "bubble", "example", "zsh"}, actual)
}

func TestIntSliceReverseSort(t *testing.T) {
	actual := reverseSortIntSlice([]int{3, 1, 5, 4, 2, 5})
	assert.Equal(t, []int{5, 5, 4, 3, 2, 1}, actual)
}

func TestStringSliceReverseSort(t *testing.T) {
	actual := reverseSortStringSlice([]string{"bubble", "basic", "actual", "zsh", "example"})
	assert.Equal(t, []string{"zsh", "example", "bubble", "basic", "actual"}, actual)
}

func TestSortStructByName(t *testing.T) {
	personSlice := []Person{
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
		{"first", 1},
	}
	expected := []Person{
		{"first", 1},
		{"fourth", 4},
		{"second", 2},
		{"third", 3},
	}
	actual := sortStructByName(personSlice)
	assert.Equal(t, expected, actual)
}

func TestReverseSortStructByName(t *testing.T) {
	personSlice := []Person{
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
		{"first", 1},
	}
	expected := []Person{
		{"third", 3},
		{"second", 2},
		{"fourth", 4},
		{"first", 1},
	}
	actual := reverseSortStructByName(personSlice)
	assert.Equal(t, expected, actual)
}

func TestSortStructByNumber(t *testing.T) {
	personSlice := []Person{
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
		{"first", 1},
	}
	expected := []Person{
		{"first", 1},
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
	}
	actual := sortStructByNumber(personSlice)
	assert.Equal(t, expected, actual)
}

func TestSortStructByNameAndNumber(t *testing.T) {
	personSlice := []Person{
		{"third", 3},
		{"third", 33},
		{"second", 2},
		{"fourth", 444},
		{"second", 222},
		{"third", 333},
		{"first", 11},
		{"fourth", 4},
		{"fourth", 44},
		{"first", 111},
		{"second", 22},
		{"first", 1},
	}
	expected := []Person{
		{"first", 1},
		{"first", 11},
		{"first", 111},
		{"fourth", 4},
		{"fourth", 44},
		{"fourth", 444},
		{"second", 2},
		{"second", 22},
		{"second", 222},
		{"third", 3},
		{"third", 33},
		{"third", 333},
	}
	actual := sortStructByNameAndNumber(personSlice)
	assert.Equal(t, expected, actual)
}

// at first, I could not expect the result of this test
func TestReverseSortStructByNameAndNumber(t *testing.T) {
	personSlice := []Person{
		{"third", 3},
		{"third", 33},
		{"second", 2},
		{"fourth", 444},
		{"second", 222},
		{"third", 333},
		{"first", 11},
		{"fourth", 4},
		{"fourth", 44},
		{"first", 111},
		{"second", 22},
		{"first", 1},
	}
	expected := []Person{
		{"third", 333},
		{"third", 33},
		{"third", 3},
		{"second", 222},
		{"second", 22},
		{"second", 2},
		{"fourth", 444},
		{"fourth", 44},
		{"fourth", 4},
		{"first", 111},
		{"first", 11},
		{"first", 1},
	}
	actual := reverseSortStructByNameAndNumber(personSlice)
	assert.Equal(t, expected, actual)
}
