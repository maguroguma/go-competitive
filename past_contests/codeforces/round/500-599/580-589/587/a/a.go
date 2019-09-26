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
	r.Buffer(make([]byte, 1024), int(1e+9))
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
	return int(readInt())
}
func ReadInt2() (int, int) {
	return int(readInt()), int(readInt())
}
func ReadInt3() (int, int, int) {
	return int(readInt()), int(readInt()), int(readInt())
}
func ReadInt4() (int, int, int, int) {
	return int(readInt()), int(readInt()), int(readInt()), int(readInt())
}

func readInt() int32 {
	i, err := strconv.ParseInt(ReadString(), 0, 32)
	if err != nil {
		panic(err.Error())
	}
	return int32(i)
}

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadFloat returns an float.
func ReadFloat() float32 {
	return float32(readFloat())
}

func readFloat() float32 {
	f, err := strconv.ParseFloat(ReadString(), 32)
	if err != nil {
		panic(err.Error())
	}
	return float32(f)
}

// ReadFloatSlice returns an float slice that has n float.
func ReadFloatSlice(n int) []float32 {
	b := make([]float32, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
}

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()
//a := ReadFloat()
//A := ReadFloatSlice(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*******************************************************************/

// const MOD = 1000000000 + 7
// const ALPHABET_NUM = 26
// const INF_INT = math.MaxInt
// const INF_BIT60 = 1 << 60

var n int
var S []rune

func main() {
	n = ReadInt()
	S = ReadRuneSlice()

	anum, bnum := 0, 0
	opnum := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'a' {
			if anum+1-bnum >= 2 {
				S[i] = 'b'
				opnum++
				bnum++
			} else {
				anum++
			}
		} else {
			if bnum+1-anum >= 2 {
				S[i] = 'a'
				opnum++
				anum++
			} else {
				bnum++
			}
		}
	}

	fmt.Println(opnum)
	fmt.Println(string(S))
}

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/
