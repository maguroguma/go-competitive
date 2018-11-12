package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type bufReader struct {
	r   *bufio.Reader
	buf []byte
	i   int
}

var reader = &bufReader{
	bufio.NewReader(os.Stdin),
	make([]byte, 0),
	0,
}

func (r *bufReader) readLine() {
	if r.i < len(r.buf) {
		return
	}
	r.buf = make([]byte, 0)
	r.i = 0
	for {
		line, isPrefix, err := r.r.ReadLine()
		if err != nil {
			panic(err)
		}
		r.buf = append(r.buf, line...)
		if !isPrefix {
			break
		}
	}
}

func (r *bufReader) next() string {
	r.readLine()
	from := r.i
	for ; r.i < len(r.buf); r.i++ {
		if r.buf[r.i] == ' ' {
			break
		}
	}
	s := string(r.buf[from:r.i])
	r.i++
	return s
}

func (r *bufReader) nextLine() string {
	r.readLine()
	s := string(r.buf[r.i:])
	r.i = len(r.buf)
	return s
}

var writer = bufio.NewWriter(os.Stdout)

func next() string {
	return reader.next()
}

func nextInt64() int64 {
	i, err := strconv.ParseInt(reader.next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func nextInt() int {
	return int(nextInt64())
}

func nextLine() string {
	return reader.nextLine()
}

func out(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func max(x, y int) int {
	return int(max64(int64(x), int64(y)))
}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func min(x, y int) int {
	return int(min64(int64(x), int64(y)))
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func abs(x int) int {
	return int(abs64(int64(x)))
}

func joinInt64s(a []int64, sep string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(b, sep)
}

func joinInts(a []int, sep string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

func divUp64(x, y int64) int64 {
	return (x + y - 1) / y
}

func divUp(x, y int) int {
	return int(divUp64(int64(x), int64(y)))
}

func gcd64(x, y int64) int64 {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func gcd(x, y int) int {
	return int(gcd64(int64(x), int64(y)))
}

func lcm64(x, y int64) int64 {
	return x * y / gcd64(x, y)
}

func lcm(x, y int) int {
	return int(lcm64(int64(x), int64(y)))
}

func pow64(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func pow(x, y int) int {
	return int(pow64(int64(x), int64(y)))
}

/*******************************************************************/

var n int
var A []int

func main() {
	n := nextInt()
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
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
