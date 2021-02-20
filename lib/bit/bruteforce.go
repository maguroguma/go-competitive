package bit

// originated from ktateish@AtCoder
// e.g.: https://atcoder.jp/contests/abc167/submissions/13042361

// BruteForceByBits01(bitsNum, fn) calls fn with []int for each n-bit 0/1 pattern
func BruteForceByBits01(bitsNum int, fn func(bits []int)) {
	// e.g.
	// BruteForceByBits01(10, func(b []int) { fmt.Println(b) }
	// -> [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [0, 1, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [1, 1, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
	//    ...
	N := 1 << uint(bitsNum)
	a := make([]int, bitsNum)
	for i := 0; i < N; i++ {
		for j := 0; j < bitsNum; j++ {
			k := bitsNum - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = 0
			} else {
				a[k] = 1
			}
		}
		fn(a)
	}
}

// BruteForceByBitsTF(bitsNum, fn) calls fn with []bool for each n-bit true/false pattern
func BruteForceByBitsTF(bitsNum int, fn func(bitFlags []bool)) {
	// e.g.
	// BruteForceByBitsTF(10, func(b []bool) { fmt.Println(b) }
	// -> [false, false, false, false, false, false, false, false, false, false]
	//    [true, false, false, false, false, false, false, false, false, false]
	//    [false, true, false, false, false, false, false, false, false, false]
	//    ...
	N := 1 << uint(bitsNum)
	a := make([]bool, bitsNum)
	for i := 0; i < N; i++ {
		for j := 0; j < bitsNum; j++ {
			k := bitsNum - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = false
			} else {
				a[k] = true
			}
		}
		fn(a)
	}
}
