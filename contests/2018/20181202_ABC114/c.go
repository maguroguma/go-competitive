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

var N int

func main() {
	tmp := NextIntsLine()
	N = tmp[0]

	ans := 0
	digitNum := getDigitNum(N)
	limitedDigits := []rune{'3', '5', '7'}
	//	for d := 3; d <= digitNum; d++ {
	// d桁の7,5,3で構成される数をすべて作ってチェックする
	//		numRunes := make([]rune, d)
	//		for counter := 0; counter < d; counter++ {
	//			numRunes[counter] = limitedDigits[0]
	//			for i := 0; i < 3; i++ {
	//				for j := 0; j < 3; j++ {
	//					for k := 0; k < 3; k++ {
	//
	//					}
	//				}
	//			}
	//		}
	//	}
	maxTimes := PowInt(3, digitNum)
	times := 0
	for i := 0; i < 3; i++ {
		//		numRunes := []rune{}
		//		numRunes = append(numRunes, limitedDigits[i])
		for j := 0; j < 3; j++ {
			//			numRunes = append(numRunes, limitedDigits[j])
			for k := 0; k < 3; k++ {
				//				numRunes = append(numRunes, limitedDigits[k])
				for l := 0; l < 3; l++ {
					//					numRunes = append(numRunes, limitedDigits[l])
					for m := 0; m < 3; m++ {
						//						numRunes = append(numRunes, limitedDigits[m])
						for n := 0; n < 3; n++ {
							//							numRunes = append(numRunes, limitedDigits[n])
							for o := 0; o < 3; o++ {
								//								numRunes = append(numRunes, limitedDigits[o])
								for p := 0; p < 3; p++ {
									//									numRunes = append(numRunes, limitedDigits[p])
									for q := 0; q < 3; q++ {
										//										numRunes = append(numRunes, limitedDigits[q])
										for r := 0; r < 3; r++ {
											//											numRunes = append(numRunes, limitedDigits[r])
											numRunes := []rune{}
											numRunes = append(numRunes, limitedDigits[i], limitedDigits[j], limitedDigits[k], limitedDigits[l], limitedDigits[m], limitedDigits[n], limitedDigits[o], limitedDigits[p], limitedDigits[q], limitedDigits[r])
											for d := 3; d <= digitNum; d++ {
												checkedRunes := numRunes[len(numRunes)-d : len(numRunes)]
												times++
												//fmt.Println(string(checkedRunes))
												if sub(checkedRunes) {
													tmpNum, _ := strconv.Atoi(string(checkedRunes))
													if tmpNum <= N {
														ans++
													}
												}
												if times == maxTimes {
													fmt.Println(ans)
													fmt.Println(times)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	//fmt.Println(ans)
}

// runeスライスが753数かどうか
func sub(S []rune) bool {
	c3, c5, c7 := 0, 0, 0
	for _, r := range S {
		//digit, _ := strconv.Atoi(string(r))
		if r == '7' {
			c7++
		} else if r == '5' {
			c5++
		} else if r == '3' {
			c3++
		}
	}

	if c3 >= 1 && c5 >= 1 && c7 >= 1 {
		return true
	} else {
		return false
	}
}

func getDigitNum(n int) int {
	dnum := 0
	for n > 0 {
		n /= 10
		dnum++
	}
	return dnum
}
