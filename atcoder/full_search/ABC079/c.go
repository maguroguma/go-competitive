package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a, b, c, d int
var input string

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	inputInt, _ := strconv.Atoi(input)
	a = inputInt / 1000
	b = inputInt / 100 % 10
	c = inputInt / 10 % 10
	d = inputInt / 1 % 10
	ints := []int{a, b, c, d}

	// forループによるbit全探索
	for i := 0; i < (1 << 3); i++ {
		answer := ""
		sum := ints[0]
		answer += strconv.FormatInt(int64(ints[0]), 10)
		for j := 0; j < 3; j++ {
			bit := 1 & (i >> byte(j))
			if bit == 1 {
				sum += ints[j+1]
				answer += "+" + strconv.FormatInt(int64(ints[j+1]), 10)
			} else {
				sum -= ints[j+1]
				answer += "-" + strconv.FormatInt(int64(ints[j+1]), 10)
			}
		}

		if sum == 7 {
			fmt.Println(answer + "=7")
			return
		}
	}
}
