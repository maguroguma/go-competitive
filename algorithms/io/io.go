package io

// PrintIntsLine returns integers string delimited by a space.
func PrintIntsLine(A ...int) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		if i == len(A)-1 {
			res = append(res)
		}
	}
}
