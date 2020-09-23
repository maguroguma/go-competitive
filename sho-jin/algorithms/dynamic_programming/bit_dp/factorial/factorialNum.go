package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

// ReadInt returns an integer.
func ReadInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
}

/*******************************************************************/

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26

var n int
var dp [1 << uint(20)]int

// インプットされた整数の階乗の値（すべての順列の場合の数）を計算する
func main() {
  n = ReadInt()

  dp[0] = 1 // 数え上げなので、1からスタートする

  iMax := 1 << uint(n)
  // n匹のうさぎをビットで表すため、2^n-1までループを回す
  for i := 0; i < iMax; i++ {
    // 0 ~ 2^n-1 の各ビット列に対して、うさぎn匹分について調べる
    for usagiNth := 0; usagiNth < n; usagiNth++ {
      usabit := 1 << uint(usagiNth)
      if i & usabit != 0 {
        // 調べているうさぎのビットが立っていたら、そのうさぎのビットが立つ前のDPテーブルの値を加算する
        dp[i] += dp[i - usabit]
      }
    }
  }

  fmt.Println(dp[1 << uint(n) - 1])
}

