package strings

import (
	"math"
	"sort"
)

// originated from:
// https://qiita.com/EmptyBox_0/items/2f8e3cf7bd44e0f789d5#strings
// https://atcoder.github.io/ac-library/production/document_ja/string.html

func SuffixArrayString(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return _saIs(s2, 255)
}

func LcpArrayString(s string, sa []int) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return LcpArrayIntSlice(s2, sa)
}

// MatchBySA finds all matches between a text and a pattern by using Suffix Array.
// Time complexity: O(|P| * log|T|)
func MatchBySA(text string, tsa []int, pattern string) (left, right int) {
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
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	left = bs(len(tsa), -1, func(mid int) bool {
		L := tsa[mid]
		R := min(tsa[mid]+len(pattern), len(text))
		substr := text[L:R]
		return pattern <= substr
	})
	right = bs(-1, len(tsa), func(mid int) bool {
		L := tsa[mid]
		R := min(tsa[mid]+len(pattern), len(text))
		substr := text[L:R]
		return pattern >= substr
	})

	return
}

func SuffixArrayIntSlice(s []int) []int {
	n := len(s)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(l, r int) bool { return s[l] < s[r] })
	s2 := make([]int, n)
	now := 0
	for i := 0; i < n; i++ {
		if i != 0 && s[idx[i-1]] != s[idx[i]] {
			now++
		}
		s2[idx[i]] = now
	}
	return _saIs(s2, now)
}

func SuffixArrayLimitedIntSlice(s []int, upper int) []int {
	sa := _saIs(s, upper)
	return sa
}

func LcpArrayIntSlice(s, sa []int) []int {
	n := len(s)
	rnk := make([]int, n)
	for i := 0; i < n; i++ {
		rnk[sa[i]] = i
	}
	lcp := make([]int, n-1)
	h := 0
	for i := 0; i < n; i++ {
		if h > 0 {
			h--
		}
		if rnk[i] == 0 {
			continue
		}
		j := sa[rnk[i]-1]
		for ; j+h < n && i+h < n; h++ {
			if s[j+h] != s[i+h] {
				break
			}
		}
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func _saIs(s []int, upper int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if s[0] < s[1] {
			return []int{0, 1}
		} else {
			return []int{1, 0}
		}
	}
	sa := make([]int, n)
	ls := make([]bool, n)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			ls[i] = ls[i+1]
		} else {
			ls[i] = s[i] < s[i+1]
		}
	}
	sumL, sumS := make([]int, upper+1), make([]int, upper+1)
	for i := 0; i < n; i++ {
		if !ls[i] {
			sumS[s[i]]++
		} else {
			sumL[s[i]+1]++
		}
	}
	for i := 0; i <= upper; i++ {
		sumS[i] += sumL[i]
		if i < upper {
			sumL[i+1] += sumS[i]
		}
	}
	induce := func(lms []int) {
		for i := 0; i < n; i++ {
			sa[i] = -1
		}
		buf := make([]int, upper+1)
		copy(buf, sumS)
		for _, d := range lms {
			if d == n {
				continue
			}
			sa[buf[s[d]]] = d
			buf[s[d]]++
		}
		copy(buf, sumL)
		sa[buf[s[n-1]]] = n - 1
		buf[s[n-1]]++
		for i := 0; i < n; i++ {
			v := sa[i]
			if v >= 1 && !ls[v-1] {
				sa[buf[s[v-1]]] = v - 1
				buf[s[v-1]]++
			}
		}
		copy(buf, sumL)
		for i := n - 1; i >= 0; i-- {
			v := sa[i]
			if v >= 1 && ls[v-1] {
				buf[s[v-1]+1]--
				sa[buf[s[v-1]+1]] = v - 1
			}
		}
	}
	lmsMap := make([]int, n+1)
	for i, _ := range lmsMap {
		lmsMap[i] = -1
	}
	m := 0
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lmsMap[i] = m
			m++
		}
	}
	lms := make([]int, 0, m)
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lms = append(lms, i)
		}
	}
	induce(lms)
	if m != 0 {
		sortedLms := make([]int, 0, m)
		for _, v := range sa {
			if lmsMap[v] != -1 {
				sortedLms = append(sortedLms, v)
			}
		}
		recS := make([]int, m)
		recUpper := 0
		recS[lmsMap[sortedLms[0]]] = 0
		for i := 1; i < m; i++ {
			l := sortedLms[i-1]
			r := sortedLms[i]
			endL, endR := n, n
			if lmsMap[l]+1 < m {
				endL = lms[lmsMap[l]+1]
			}
			if lmsMap[r]+1 < m {
				endR = lms[lmsMap[r]+1]
			}
			same := true
			if endL-l != endR-r {
				same = false
			} else {
				for l < endL {
					if s[l] != s[r] {
						break
					}
					l++
					r++
				}
				if l == n || s[l] != s[r] {
					same = false
				}
			}
			if !same {
				recUpper++
			}
			recS[lmsMap[sortedLms[i]]] = recUpper
		}
		recSa := _saIs(recS, recUpper)
		for i := 0; i < m; i++ {
			sortedLms[i] = lms[recSa[i]]
		}
		induce(sortedLms)
	}
	return sa
}
