package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	sc.Scan()
	return sc.Text()
}

// NextIntsLine reads a line text, that consists of **ONLY INTEGERS DELIMITED BY SPACES**, from stdin.
// And then returns intergers slice.
func NextIntsLine() []int {
	ints := []int{}
	intsStr := NextLine()
	tmp := strings.Split(intsStr, " ")
	for _, s := range tmp {
		integer, _ := strconv.Atoi(s)
		ints = append(ints, integer)
	}
	return ints
}

/*******************************************************************/

// https://atc001.contest.atcoder.jp/tasks/unionfind_a

var n, q int
var P, A, B []int

// 親の番号
// par[i] == i ならば根（はじめは全部の頂点が根）
var par [100005]int

func main() {
	tmp := NextIntsLine()
	n, q = tmp[0], tmp[1]
	for i := 0; i < q; i++ {
		tmp = NextIntsLine()
		P = append(P, tmp[0])
		A = append(A, tmp[1])
		B = append(B, tmp[2])
	}

	// 初期化
	initialize(n)
	for i := 0; i < q; i++ {
		if P[i] == 0 {
			unite(A[i], B[i])
		} else {
			b := same(A[i], B[i])
			if b {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

// n要素で初期化
func initialize(n int) {
	for i := 0; i < n; i++ {
		par[i] = i
	}
}

// 木の根を求める
// 経路圧縮: 上向きにたどって再帰的に根を調べる際に、調べたら辺を根に直接つなぎ直す（xの親を根に変える）
func root(x int) int {
	if par[x] == x {
		return x
	} else {
		par[x] = root(par[x])
		return par[x] // **経路圧縮**
	}
}

// xとyが同じ集合に属するか否か
func same(x, y int) bool {
	return root(x) == root(y)
}

// xとyの属する集合を併合
func unite(x, y int) {
	x = root(x)
	y = root(y)
	if x == y {
		return
	}

	par[x] = y
}
