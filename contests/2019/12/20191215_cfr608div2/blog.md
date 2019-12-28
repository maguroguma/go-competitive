<!-- Codeforces Round No.608 (Div.2) 参加記録 (A〜C解答) -->

<!-- TOC -->

- [A. Suits](#a-suits)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Blocks](#b-blocks)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Shawarma Tent](#c-shawarma-tent)
	- [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

※Dはちょっと難しそうなので手を付けるのは当分先になりそうですが、
Eがシンプルな見た目で面白そうだったため、近々追記するかもしれません。

<a id="markdown-a-suits" name="a-suits"></a>
## A. Suits

[問題のURL](https://codeforces.com/contest/1271/problem/A)

<a id="markdown-解答" name="解答"></a>
### 解答

2つのセットでジャケットが共通しており、ほかは独立している。
セットに組み込まれるアイテムはすべて1つずつであるため、
ジャケットはできるだけ高いセットで先に使ってしまい、
作れなくなったら、他方の安い方をできる限り作れば良い。

```go
var a, b, c, d, e, f int

func main() {
	a, b, c, d = ReadInt4()
	e, f = ReadInt2()

	bc := Min(b, c)

	ans := 0
	if e > f {
		// aの方を使う
		m := Min(a, d)
		ans += e * m
		d -= m
		if d > 0 {
			mm := Min(bc, d)
			ans += f * mm
		}
	} else {
		// bcの方を使う
		m := Min(bc, d)
		ans += f * m
		d -= m
		if d > 0 {
			mm := Min(a, d)
			ans += e * mm
		}
	}

	fmt.Println(ans)
}
```

これもDiv.2のAの中では大分簡単な方な気がします。

<a id="markdown-b-blocks" name="b-blocks"></a>
## B. Blocks

[問題のURL](https://codeforces.com/contest/1271/problem/B)

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

初期の個数が偶数のものを反転させれば良い。

なぜなら、2つのブロックを反転させるとき、反転させて個数をゼロにしたい方の色のブロックの個数は、

1. 2個減る
2. 変わらない（1個減って1個増える）

のいずれかであることから、依然として偶数のままとなる。

最終的に選んだ色を0個にしなければならないことを考えると、
初期の個数が偶数の方しか選ぶことが出来ない。

具体的な反転回数については、左から愚直に選んだ色を反転させるようにすればよく、
高々 `n` 回で十分である。

```go
var n int
var S []rune

func main() {
	n = ReadInt()
	S = ReadRuneSlice()

	// 最初から全部いっしょならOK
	memo := make(map[rune]int)
	memo['B'] = 0
	memo['W'] = 0
	for i := 0; i < n; i++ {
		memo[S[i]]++
	}
	if memo['B'] == n || memo['W'] == n {
		fmt.Println(0)
		return
	}

	// BもWも奇数なら無理
	if memo['B']%2 == 1 && memo['W']%2 == 1 {
		fmt.Println(-1)
		return
	}

	answers := []int{}
	if memo['B']%2 == 0 {
		// BをWにする
		for i := 0; i < n-1; i++ {
			if S[i] == 'B' {
				answers = append(answers, i+1)
				S[i] = 'W'
				S[i+1] = rev(S[i+1])
			}
		}
	} else {
		// WをBにする
		for i := 0; i < n-1; i++ {
			if S[i] == 'W' {
				answers = append(answers, i+1)
				S[i] = 'B'
				S[i+1] = rev(S[i+1])
			}
		}
	}

	fmt.Println(len(answers))
	fmt.Println(PrintIntsLine(answers...))
}

func rev(r rune) rune {
	if r == 'B' {
		return 'W'
	} else {
		return 'B'
	}
}
```

結構似たような問題は解いた気がするのに、かなり時間がかかってしまいました。

精進不足。

<a id="markdown-c-shawarma-tent" name="c-shawarma-tent"></a>
## C. Shawarma Tent

[問題のURL](https://codeforces.com/contest/1271/problem/C)

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

よくよく考えると、学校の1マス左・右・下・上のいずれかにテントを配置するのが最適です。

それぞれに配置した場合の、買い食いを行う可能性のある生徒の数を数え、
最大となるものを選べばよいです。

```go
var n int
var sx, sy int
var X, Y []int

func main() {
	n = ReadInt()
	sx, sy = ReadInt2()

	X, Y = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x, y := ReadInt2()
		X[i], Y[i] = x, y
	}

	l, r, u, d := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		x, y := X[i], Y[i]

		if x <= sx-1 {
			l++
		}
		if x >= sx+1 {
			r++
		}
		if y <= sy-1 {
			d++
		}
		if y >= sy+1 {
			u++
		}
	}

	maxi := Max(l, r, u, d)
	fmt.Println(maxi)
	if l == maxi {
		fmt.Println(sx-1, sy)
	} else if r == maxi {
		fmt.Println(sx+1, sy)
	} else if d == maxi {
		fmt.Println(sx, sy-1)
	} else {
		fmt.Println(sx, sy+1)
	}
}
```

コンテスト中に通した解法では、片方の軸を学校に合わせれば良いということしか気づけず、
素直に区間で一番重なりが多い部分を求めるべく、家の位置のソート・ランレングス圧縮・累積和とかいう、
闇の深いコードを書いてしまいました（なぜ通ったのか）。

問題文がいかついと複雑に考えてしまうような気がするので、もうちょっと冷静に問題を俯瞰したいところです。

