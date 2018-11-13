package main

import (
	"sort"
)

func sortIntSlice(A []int) []int {
	sort.Sort(sort.IntSlice(A)) // sort.IntSlice(A) is type cast
	return A
}
func reverseSortIntSlice(A []int) []int {
	temp := sort.IntSlice(A)
	sort.Sort(sort.Reverse(temp))
	return A
}

func sortStringSlice(A []string) []string {
	sort.Sort(sort.StringSlice(A)) // sort.StringSlice(A) is type cast
	return A
}
func reverseSortStringSlice(A []string) []string {
	temp := sort.StringSlice(A)
	sort.Sort(sort.Reverse(temp))
	return A
}

type Person struct {
	name   string
	number int
}

// type name is necessary when you want to define method
type people []Person
type byName struct { // type extension (`p people` is not correct)
	people
}
type byNumber struct {
	people
}
type byNameAndNumber struct {
	people
}

// Len is the number of elements in the collection
func (pSlice people) Len() int {
	return len(pSlice)
}

// Swap swaps the elements with indices i and j
func (pSlice people) Swap(i, j int) {
	pSlice[i], pSlice[j] = pSlice[j], pSlice[i]
}

// Less returns whether the element with index i should sort before the element with index j
// Customize this method to realize different sort!
//func (pSlice people) Less(i, j int) bool {
//	return pSlice[i].name < pSlice[j].name
//}
func (b byName) Less(i, j int) bool {
	return b.people[i].name < b.people[j].name
}
func (b byNumber) Less(i, j int) bool {
	return b.people[i].number < b.people[j].number
}
func (b byNameAndNumber) Less(i, j int) bool {
	if b.people[i].name < b.people[j].name {
		return true
	} else if b.people[i].name == b.people[j].name {
		return b.people[i].number < b.people[j].number
	} else {
		return false
	}
}

func sortStructByName(pSlice []Person) []Person {
	sort.Sort(byName{people(pSlice)})
	return pSlice
}

func reverseSortStructByName(pSlice []Person) []Person {
	sort.Sort(sort.Reverse(byName{people(pSlice)}))
	return pSlice
}

func sortStructByNumber(pSlice []Person) []Person {
	sort.Sort(byNumber{people(pSlice)})
	return pSlice
}

func sortStructByNameAndNumber(pSlice []Person) []Person {
	sort.Sort(byNameAndNumber{people(pSlice)})
	return pSlice
}

func reverseSortStructByNameAndNumber(pSlice []Person) []Person {
	sort.Sort(sort.Reverse(byNameAndNumber{people(pSlice)}))
	return pSlice
}

func main() {
}
