package main

import "fmt"

func main() {
	fmt.Println(gcd(6, 15))
	fmt.Println(gcd(6, 18))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
