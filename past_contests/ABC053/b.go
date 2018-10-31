package main

import (
	"bufio"
	"fmt"
	"os"
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

var s string

func main() {
	s = readLine()
	runes := []rune(s)
	head, tail := -1, 0
	for i, r := range runes {
		if head == -1 && r == 'A' {
			head = i
		}
		if r == 'Z' {
			tail = i
		}
	}
	fmt.Println(tail - head + 1)
}
