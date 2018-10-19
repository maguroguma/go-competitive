package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test複合体のソート(t *testing.T) {
	people := []*Person{
		&Person{age: 0, height: 1.0, name: "tdd"},
		&Person{age: 1, height: 0.5, name: "ddd"},
		&Person{age: 2, height: 0.7, name: "bdd"},
	}
	sortByInt(people)
	assert.Equal(t, []int{0, 1, 2}, []int{people[0].age, people[1].age, people[2].age})
	sortDescByInt(people)
	assert.Equal(t, []int{2, 1, 0}, []int{people[0].age, people[1].age, people[2].age})
	sortByFloat64(people)
	assert.Equal(t, []float64{0.5, 0.7, 1.0}, []float64{people[0].height, people[1].height, people[2].height})
	sortDescByFloat64(people)
	assert.Equal(t, []float64{1.0, 0.7, 0.5}, []float64{people[0].height, people[1].height, people[2].height})
	sortDescByString(people)
	assert.Equal(t, []string{"tdd", "ddd", "bdd"}, []string{people[0].name, people[1].name, people[2].name})
}
