package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

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

/*******************************************************************/

var n int
var A []int

func main() {
	n, _ := strconv.Atoi(readLine())
	A := make([]int, 0, 1000000)
	str := readLine()
	strSlices := strings.Split(str, " ")
	for _, s := range strSlices {
		i, _ := strconv.Atoi(s)
		A = append(A, i)
	}
	mod := 1000000000 + 7

	counter := make([]int, n)
	if n%2 == 1 { // 奇数（最小値は1）
		for i := 0; i < n; i += 2 {
			if i == 0 {
				counter[i] = 1
			} else {
				counter[i] = 2
			}
		}
	} else { // 偶数（最小値は2）
		for i := 1; i < n; i += 2 {
			counter[i] = 2
		}
	}

	// 判定
	for i := 0; i < n; i++ {
		idx := A[i]
		counter[idx]--
	}
	for _, e := range counter {
		if e != 0 {
			fmt.Println(0)
			return
		}
	}

	// 答えは2^(n/2)
	index := n / 2
	answer := 1
	for i := 0; i < index; i++ {
		answer = answer * 2 % mod
	}
	fmt.Println(answer)
}
