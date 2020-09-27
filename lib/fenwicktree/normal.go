package fenwicktree

type FenwickTree struct {
	dat     []int
	n       int
	minPow2 int
}

// n(>=1) is number of elements of original data
func NewFenwickTree(n int) *FenwickTree {
	ft := new(FenwickTree)

	ft.dat = make([]int, n+1)
	ft.n = n

	ft.minPow2 = 1
	for {
		if (ft.minPow2 << 1) > n {
			break
		}
		ft.minPow2 <<= 1
	}

	return ft
}

// Add x to i.
// i is 0-index.
func (ft *FenwickTree) Add(i int, x int) {
	ft._add(i+1, x)
}

// RangeSum calculates a range sum of [l, r).
// l, r are 0-index.
func (ft *FenwickTree) RangeSum(l, r int) int {
	res := ft._cumsum(r) - ft._cumsum(l)

	return res
}

// LowerBound returns minimum i such that bit.Sum(i) >= w.
// LowerBound returns n, if the total sum is less than w.
func (ft *FenwickTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

	x := 0
	for k := ft.minPow2; k > 0; k /= 2 {
		if x+k <= ft.n && ft.dat[x+k] < w {
			w -= ft.dat[x+k]
			x += k
		}
	}

	return x
}

// private: Sum of [1, i](1-based)
func (ft *FenwickTree) _cumsum(i int) int {
	s := 0

	for i > 0 {
		s += ft.dat[i]
		i -= i & (-i)
	}

	return s
}

// private: Add x to i(1-based)
func (ft *FenwickTree) _add(i, x int) {
	for i <= ft.n {
		ft.dat[i] += x
		i += i & (-i)
	}
}
