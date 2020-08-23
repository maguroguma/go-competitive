package arithmetic

import "math"

func IsSquareNumber(n int) bool {
	if n == 0 {
		return true
	}

	if n < 0 {
		panic("argument should be not negative number")
	}

	bs := func(initOK, initNG int, isOK func(mid int) bool) (ok int) {
		ng := initNG
		ok = initOK
		for int(math.Abs(float64(ok-ng))) > 1 {
			mid := (ok + ng) / 2
			if isOK(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}

		return ok
	}

	ok := bs(1<<30, 0, func(m int) bool {
		return m*m >= n
	})

	return ok*ok == n
}

func IsCubeNumber(n int) bool {
	if n == 0 {
		return true
	}

	if n < 0 {
		panic("argument should be not negative number")
	}

	bs := func(initOK, initNG int, isOK func(mid int) bool) (ok int) {
		ng := initNG
		ok = initOK
		for int(math.Abs(float64(ok-ng))) > 1 {
			mid := (ok + ng) / 2
			if isOK(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}

		return ok
	}

	ok := bs(1<<20, 0, func(m int) bool {
		return m*m*m >= n
	})

	return ok*ok*ok == n
}
