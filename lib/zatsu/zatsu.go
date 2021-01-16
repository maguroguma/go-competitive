package zatsu

import (
	"sort"
)

// NewCompress returns a compress algorithm.
func NewCompress() *Compress {
	c := new(Compress)
	c.xs = []int64{}
	c.cs = []int64{}

	return c
}

// Add can add any number of elements.
// Time complexity: O(1)
func (c *Compress) Add(X ...int64) {
	c.xs = append(c.xs, X...)
}

// Build compresses input elements by sorting.
// Time complexity: O(NlogN)
func (c *Compress) Build() {
	sort.Slice(c.xs, func(i, j int) bool {
		return c.xs[i] < c.xs[j]
	})

	if len(c.xs) == 0 {
		panic("Compress doesn't have any elements")
	}

	c.cs = append(c.cs, c.xs[0])
	for i := 1; i < len(c.xs); i++ {
		if c.xs[i-1] == c.xs[i] {
			continue
		}
		c.cs = append(c.cs, c.xs[i])
	}
}

// Get returns index that is equal to by binary search.
// Results are in [0, len(c.cs)).
// Time complexity: O(logN)
func (c *Compress) Get(x int64) int64 {
	_abs := func(a int64) int64 {
		if a < 0 {
			return -a
		}
		return a
	}

	var ng, ok = int64(-1), int64(len(c.cs))
	for _abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if c.cs[mid] >= x {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// InvGet returns original value that equals to i (compressed values).
// InvGet accepts [0, len(c.cs))
// Time complexity: O(1)
func (c *Compress) InvGet(i int64) int64 {
	if !(0 <= i && i < int64(len(c.cs))) {
		panic("i is out of range")
	}
	return c.cs[i]
}

// Kind returns number of different values, that is len(c.cs).
// Time complexity: O(1)
func (c *Compress) Kind() int {
	return len(c.cs)
}

type Compress struct {
	xs []int64 // sorted original values
	cs []int64 // sorted and compressed original values
}
