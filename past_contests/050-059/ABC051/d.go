package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)
// readLine can read long line string (at least 10^5)
func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	return readLine()
}
*/

var sc = bufio.NewScanner(os.Stdin)

// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	sc.Scan()
	return sc.Text()
}

// NextIntsLine reads a line text, that consists of **ONLY INTEGERS DELIMITED BY SPACES**, from stdin.
// And then returns intergers slice.
func NextIntsLine() []int {
	ints := []int{}
	intsStr := NextLine()
	tmp := strings.Split(intsStr, " ")
	for _, s := range tmp {
		integer, _ := strconv.Atoi(s)
		ints = append(ints, integer)
	}
	return ints
}

// NextStringsLine reads a line text, that consists of **STRINGS DELIMITED BY SPACES**, from stdin.
// And then returns strings slice.
func NextStringsLine() []string {
	str := NextLine()
	return strings.Split(str, " ")
}

// NextRunesLine reads a line text, that consists of **ONLY CHARACTERS ARRANGED CONTINUOUSLY**, from stdin.
// Ant then returns runes slice.
func NextRunesLine() []rune {
	return []rune(NextLine())
}

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

// PowInt is integer version of math.Pow
func PowInt(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}
	fa := float64(a)
	fe := float64(e)
	fanswer := math.Pow(fa, fe)
	return int(fanswer)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	fa := float64(a)
	fanswer := math.Abs(fa)
	return int(fanswer)
}

// DeleteElement returns a *NEW* slice, that have the same and minimum length and capacity.
// DeleteElement makes a new slice by using easy slice literal.
func DeleteElement(s []int, i int) []int {
	if i < 0 || len(s) <= i {
		panic(errors.New("[index error]"))
	}
	// appendのみの実装
	n := make([]int, 0, len(s)-1)
	n = append(n, s[:i]...)
	n = append(n, s[i+1:]...)
	return n
}

// Concat returns a *NEW* slice, that have the same and minimum length and capacity.
func Concat(s, t []rune) []rune {
	n := make([]rune, 0, len(s)+len(t))
	n = append(n, s...)
	n = append(n, t...)
	return n
}

// UpperRune is rune version of `strings.ToUpper()`.
func UpperRune(r rune) rune {
	str := strings.ToUpper(string(r))
	return []rune(str)[0]
}

// LowerRune is rune version of `strings.ToLower()`.
func LowerRune(r rune) rune {
	str := strings.ToLower(string(r))
	return []rune(str)[0]
}

// ToggleRune returns a upper case if an input is a lower case, v.v.
func ToggleRune(r rune) rune {
	var str string
	if 'a' <= r && r <= 'z' {
		str = strings.ToUpper(string(r))
	} else if 'A' <= r && r <= 'Z' {
		str = strings.ToLower(string(r))
	} else {
		str = string(r)
	}
	return []rune(str)[0]
}

// ToggleString iteratively calls ToggleRune, and returns the toggled string.
func ToggleString(s string) string {
	inputRunes := []rune(s)
	outputRunes := make([]rune, 0, len(inputRunes))
	for _, r := range inputRunes {
		outputRunes = append(outputRunes, ToggleRune(r))
	}
	return string(outputRunes)
}

// Strtoi is a wrapper of `strconv.Atoi()`.
// If `strconv.Atoi()` returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

// sort package (snippets)
//sort.Sort(sort.IntSlice(s))
//sort.Sort(sort.Reverse(sort.IntSlice(s)))
//sort.Sort(sort.Float64Slice(s))
//sort.Sort(sort.StringSlice(s))

// copy function
//a = []int{0, 1, 2}
//b = make([]int, len(a))
//copy(b, a)

/*******************************************************************/

var n, m int
var adjMatrix [110][110]*path
var A, B, C []int

type path struct {
	cost  int
	order []int
}

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				p := &path{cost: 0, order: []int{i}}
				adjMatrix[i][j] = p
			} else {
				p := &path{cost: 1000000, order: []int{i}}
				adjMatrix[i][j] = p
			}
		}
	}

	A, B, C := []int{}, []int{}, []int{}
	for loop := 0; loop < m; loop++ {
		tmp = NextIntsLine()
		a, b, c := tmp[0], tmp[1], tmp[2]
		//p := &path{cost: c, order: []int{a, b}}
		//adjMatrix[a][b].cost = c
		adjMatrix[a][b].cost = Min(adjMatrix[a][b].cost, c)
		adjMatrix[a][b].order = []int{a, b}
		//p = &path{cost: c, order: []int{b, a}}
		//adjMatrix[b][a].cost = c
		adjMatrix[b][a].cost = Min(adjMatrix[b][a].cost, c)
		adjMatrix[b][a].order = []int{b, a}

		A = append(A, a)
		B = append(B, b)
		C = append(C, c)
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				newCost := adjMatrix[i][k].cost + adjMatrix[k][j].cost
				if adjMatrix[i][j].cost > newCost {
					first := adjMatrix[i][k].order
					last := adjMatrix[k][j].order
					newOrder := append(first, last[1:]...)
					//adjMatrix[i][j] = &path{cost: newCost, order: newOrder}
					adjMatrix[i][j].cost = newCost
					adjMatrix[i][j].order = newOrder
				}
			}
		}
	}

	ans := m
	for idx := 0; idx < m; idx++ {
		start, end, cost := A[idx], B[idx], C[idx]
		flag := false
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				leftCost := adjMatrix[i][start].cost + cost + adjMatrix[end][j].cost
				rightCost := adjMatrix[i][j].cost
				if leftCost == rightCost {
					ans--
					flag = true
					break
				}
				//				leftCost = adjMatrix[i][end].cost + cost + adjMatrix[start][j].cost
				//				if leftCost == rightCost {
				//					ans--
				//					flag = true
				//					break
				//				}
			}
			if flag {
				break
			}
		}
	}
	fmt.Println(ans)
}
