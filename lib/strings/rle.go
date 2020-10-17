package strings

// RunLengthEncoding returns encoded slice of an input.
// Time complexity: O(|S|)
func RunLengthEncoding(S []rune) (comp []rune, cnts []int) {
	comp = []rune{}
	cnts = []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			l++
		} else {
			comp = append(comp, S[i-1])
			cnts = append(cnts, l)
			l = 1
		}
	}
	comp = append(comp, S[len(S)-1])
	cnts = append(cnts, l)

	return
}

// RunLengthDecoding decodes RLE results.
// Time complexity: O(|S|)
func RunLengthDecoding(comp []rune, cnts []int) (S []rune) {
	if len(comp) != len(cnts) {
		panic("S, L are not RunLengthEncoding results")
	}

	S = []rune{}

	for i := 0; i < len(comp); i++ {
		for j := 0; j < cnts[i]; j++ {
			S = append(S, comp[i])
		}
	}

	return
}
