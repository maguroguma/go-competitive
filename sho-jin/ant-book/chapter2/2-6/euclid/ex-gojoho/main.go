package main

import "fmt"

func main() {
	// x, yはでたらめな数で良い（extgcdの呼び出しスタックの奥底で1, 0にセットされるので関係ない）
	x, y := 1<<60, 1<<60
	d := extgcd(111, 30, &x, &y)
	fmt.Println(d, x, y)
}

// 以下はdrken流コード(https://qiita.com/drken/items/b97ff231e43bce50199a)
// 返り値: a, bの最大公約数
// ax + by = gcd(a, b) を満たす (x, y) が格納される
func extgcd(a, b int, x, y *int) int {
	// 再帰の終了条件
	if b == 0 {
		*x, *y = 1, 0
		return a
	}

	d := extgcd(b, a%b, y, x)
	*y -= (a / b) * (*x)
	return d
}

// 以下は蟻本流コード
// func extgcd(a, b int, x, y *int) int {
// 	d := a
// 	if b != 0 {
// 		d = extgcd(b, a%b, y, x)
// 		*y -= (a / b) * (*x)
// 	} else {
// 		*x, *y = 1, 0
// 	}
// 	return d
// }
