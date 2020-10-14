package bit

// originated from ktateish@AtCoder
// e.g.: https://atcoder.jp/contests/abc167/submissions/13042361

// BruteForceByBits01(n, fn) calls fn with []int for each n-bit 0/1 pattern
// e.g.
// BruteForceByBits01(10, func(b []int) { fmt.Println(b) }
// -> [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//    [0, 1, 0, 0, 0, 0, 0, 0, 0, 0]
//    [1, 1, 0, 0, 0, 0, 0, 0, 0, 0]
//    [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
//    ...
func BruteForceByBits01(n int, fn func([]int)) {
	N := 1 << uint(n)
	a := make([]int, n)
	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			k := n - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = 0
			} else {
				a[k] = 1
			}
		}
		fn(a)
	}
}

// BruteForceByBitsTF(n, fn) calls fn with []bool for each n-bit true/false pattern
// e.g.
// BruteForceByBitsTF(10, func(b []bool) { fmt.Println(b) }
// -> [false, false, false, false, false, false, false, false, false, false]
//    [true, false, false, false, false, false, false, false, false, false]
//    [false, true, false, false, false, false, false, false, false, false]
//    ...
func BruteForceByBitsTF(n int, fn func([]bool)) {
	N := 1 << uint(n)
	a := make([]bool, n)
	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			k := n - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = false
			} else {
				a[k] = true
			}
		}
		fn(a)
	}
}
