package segment

import "sort"

// MergeSegments returns merged source segments.
// Segments(Sections) are closed section.
// e.g.: [l, r], [l', r'] are merged if r >= l'
func MergeSegments(srcSegments [][2]int) [][2]int {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	res := [][2]int{}
	isInitialized := false

	if len(srcSegments) == 0 {
		return res
	}

	// current segment
	curL, curR := 0, 0

	// sort asc by LEFT coordinate
	sort.Slice(srcSegments, func(i, j int) bool {
		return srcSegments[i][0] < srcSegments[j][0]
	})

	for i := 0; i < len(srcSegments); i++ {
		seg := srcSegments[i]

		if !isInitialized {
			curL, curR = seg[0], seg[1]
			isInitialized = true
			continue
		}

		if curR >= seg[0] {
			// merge and continue
			_chmax(&curR, seg[1])
		} else {
			// do not merge, and add it to result
			res = append(res, [2]int{curL, curR})
			curL, curR = seg[0], seg[1]
		}
	}

	if isInitialized {
		res = append(res, [2]int{curL, curR})
	}

	return res
}
