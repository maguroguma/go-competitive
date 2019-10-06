package main

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []rune) ([]rune, []int) {
	runes := []rune{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		// 1文字目の場合保持
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			// 直前の文字と一致していればインクリメント
			l++
		} else {
			// 不一致のタイミングで追加し、長さをリセットする
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []rune, L []int) []rune {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []rune{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}

// RunLengthEncodingIntVer returns encoded slice of an input.
func RunLengthEncodingIntVer(S []int) ([]int, []int) {
	runes := []int{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			l++
		} else {
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecodingIntVer decodes RLE results.
func RunLengthDecodingIntVer(S []int, L []int) []int {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}
