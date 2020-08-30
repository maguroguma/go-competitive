package common

import (
	"math/rand"
	"reflect"
)

func ReverseMyself(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func RotateMyself(slice interface{}, n int) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	if length == 0 || n == 0 {
		return
	}
	n = (length + n) % length
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	for i, j := n, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func ShuffleMyself(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	for i := 0; i < length; i++ {
		j := rand.Int()%(length-i) + i
		swap(i, j)
	}
}
