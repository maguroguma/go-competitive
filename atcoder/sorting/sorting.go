package sorting

import "sort"

// Person is struct consists of int, float64, string fields.
// Example of sorting structure.
type Person struct {
	age    int
	height float64
	name   string
}
type people []*Person
type byInt struct {
	people
}
type byFloat64 struct {
	people
}
type byString struct {
	people
}

func (p people) Len() int {
	return len(p)
}
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (b byInt) Less(i, j int) bool {
	return b.people[i].age < b.people[j].age
}
func (b byFloat64) Less(i, j int) bool {
	return b.people[i].height < b.people[j].height
}
func (b byString) Less(i, j int) bool {
	return b.people[i].name < b.people[j].name
}

func sortByInt(p []*Person) {
	sort.Sort(byInt{people(p)})
}
func sortByFloat64(p []*Person) {
	sort.Sort(byFloat64{people(p)})
}
func sortByString(p []*Person) {
	sort.Sort(byString{people(p)})
}
func sortDescByInt(p []*Person) {
	sort.Sort(sort.Reverse(byInt{people(p)}))
}
func sortDescByFloat64(p []*Person) {
	sort.Sort(sort.Reverse(byFloat64{people(p)}))
}
func sortDescByString(p []*Person) {
	sort.Sort(sort.Reverse(byString{people(p)}))
}
