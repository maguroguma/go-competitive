package main

import (
	"fmt"
)

func main() {
	A := []int{}

	// fmt.Println(A[:1]) // panic
	fmt.Println(A[0:])
	// fmt.Println(A[1:]) // panic

	A = []int{1}
	// fmt.Println(A[:2]) // panic
	fmt.Println(A[1:])
	// fmt.Println(A[2:]) // panic
}
