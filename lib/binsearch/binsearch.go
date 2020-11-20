package binsearch

type jf func(mid int64) bool

func BinarySearch(initOK, initNG int64, isOK jf) (ok int64) {
	_abs := func(a int64) int64 {
		if a < 0 {
			return -a
		}
		return a
	}

	ng := initNG
	ok = initOK
	for _abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

type jffloat64 func(mid float64) bool

func BinarySearchFloat64(initOK, initNG float64, isOK jffloat64) (ok float64) {
	ng := initNG
	ok = initOK
	for i := 0; i < 100; i++ {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}
