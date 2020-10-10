package strings

type KMP struct {
	pattern  []rune
	kmpTable []int // KMP table for a pattern
}

// NewKMP returns a KMP instance for a pattern string.
// The instance has the KMP table for its pattern..
func NewKMP(pattern []rune) *KMP {
	kmp := new(KMP)

	kmp.pattern = pattern
	kmp.buildKMPTable()

	return kmp
}

// Search find all indices of a text that match a pattern.
// Indices are 0-index.
// Time complexity: O(len(text))
func (kmp *KMP) Search(text []rune) []int {
	res := []int{}

	if len(kmp.pattern) == 0 {
		return res
	}

	j := 0
	for i := 0; i < len(text); i++ {
		for j >= 0 && text[i] != kmp.pattern[j] {
			j = kmp.kmpTable[j]
		}
		j++
		if j == len(kmp.pattern) {
			res = append(res, i-j+1)
			j = kmp.kmpTable[j]
		}
	}

	return res
}

// Periods returns periods of strings P[:i-1] for all i < len(P).
// Time complexity: O(len(P))
func (kmp *KMP) Periods() []int {
	res := make([]int, len(kmp.kmpTable)-1)

	for i := 1; i < len(kmp.kmpTable); i++ {
		res[i-1] = i - kmp.kmpTable[i]
	}

	return res
}

// buildKMPTable calculates "tagged borders".
// Time complexity: O(len(pattern))
func (kmp *KMP) buildKMPTable() {
	T := make([]int, len(kmp.pattern)+1)

	j := -1
	T[0] = j
	for i := 0; i < len(kmp.pattern); i++ {
		for j >= 0 && kmp.pattern[i] != kmp.pattern[j] {
			j = T[j]
		}
		j++
		T[i+1] = j
	}

	kmp.kmpTable = T
}
