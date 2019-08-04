package main

// セグメント木を持つグローバル配列
var n int
var dat [2*100000 - 1]int

const INT_MAX = 1 << 60

// 初期化
func initialize(n_ int) {
	// 簡単のため、要素数を2のべき乗に
	n = 1
	for n < n_ {
		n *= 2
	}

	// すべての値をINT_MAXに
	for i := 0; i < 2*n-1; i++ {
		dat[i] = INT_MAX
	}
}

// k番目の値（0-based index）をaに変更
func update(k, a int) {
	// 葉の節点
	k += n - 1
	dat[k] = a

	// 登りながら更新
	for k > 0 {
		k = (k - 1) / 2
		dat[k] = Min(dat[k*2+1], dat[k*2+2])
	}
}

// [a, b) の最小値を求める
// 後ろの方の引数は、計算の簡単のための引数
// kは節点の番号、l, rはその節点が[l, r)に対応づいていることを表す
// したがって、外からはquery(a, b, 0, 0, n)として呼ぶ（再帰呼び出し）
func query(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しなければ、INT_MAX
	if r <= a || b <= l {
		return INT_MAX
	}

	// [a, b)が[l, r)を完全に含んでいれば、この節点の値
	if a <= l && r <= b {
		return dat[k]
	} else {
		// そうでなければ、2つの子の最小値
		vl := query(a, b, k*2+1, l, (l+r)/2)
		vr := query(a, b, k*2+2, (l+r)/2, r)
		return Min(vl, vr)
	}
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}
