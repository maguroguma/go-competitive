package bitope

func RandInt() int {
	var tx, ty, tz, tw = 123456789, 362436069, 521288629, 88675123
	tt := (tx ^ (tx << 11))
	tx = ty
	ty = tz
	tz = tw
	tw = (tw ^ (tw >> 19)) ^ (tt ^ (tt >> 8))
	return tw
}
