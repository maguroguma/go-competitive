package random

import (
	"math/rand"
)

// BoundedInt returns an integers that is between min and max ([min, max]).
func BoundedInt(min, max int) int {
	delta := max - min
	r := rand.Intn(delta + 1)
	return r + min
}

// BoundedFloat64 returns a float number that is between min and max ( [min, max) ).
func BoundedFloat64(min, max float64) float64 {
	delta := max - min
	r := rand.Float64()
	return r*delta + min
}
