<!-- Codeforces Round No.607 (Div.2) 参加記録 (A〜D解答) -->

<!-- TOC -->

- [A. Suffix Three](#a-suffix-three)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Azamon Web Services](#b-azamon-web-services)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Cut and Paste](#c-cut-and-paste)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Beingawesomeism](#d-beingawesomeism)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-suffix-three" name="a-suffix-three"></a>
## A. Suffix Three

[問題のURL](https://codeforces.com/contest/1281/problem/A)

<a id="markdown-解答" name="解答"></a>
### 解答

よくよく見ると末尾の2文字だけを見れば判定できるので、そこだけを見れば良い。

```go
var t int
var S []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		S = ReadRuneSlice()

		solve()
	}
}

func solve() {
	n := len(S)
	str := S[n-2:]
	if string(str) == "po" {
		fmt.Println("FILIPINO")
	} else if string(str) == "su" {
		fmt.Println("JAPANESE")
	} else {
		fmt.Println("KOREAN")
	}
}
```

ここまで簡単なAはDiv.3でも見たことがないので新鮮でした。

<a id="markdown-b-azamon-web-services" name="b-azamon-web-services"></a>
## B. Azamon Web Services

[問題のURL](https://codeforces.com/contest/1281/problem/B)

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

1回までの入れ替えを許可されている文字列 `S` について、辞書順最小のものを求める。
それが `C` よりも辞書順で小さければ、それを出力すれば良い。
それが `C` よりも辞書順で大きいのならば、impossibleである。

`O(n^2)` が許される制約なので、愚直に考える。

できるだけ小さい文字をできるだけ前方に配置するのが良いので、
探索開始位置を先頭から末尾に1ずつ移動させていき、チェック位置の文字よりも小さいものが、
末尾までに存在すればそれと入れ替えるようにする。

ただし、最小のものが複数ある場合は、できる限り後の物を選ぶ必要があることに注意する。

```go
var t int
var S, C []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		S = ReadRuneSlice()
		C = ReadRuneSlice()

		solve()
	}
}

const impossible = "---"

func solve() {
	n := len(S)
	lidx, ridx := -1, -1
	for i := 0; i < n; i++ {
		midx, minc := 5005, rune(100)
		// できるだけ後方から最小の文字を見つける
		for j := n - 1; j >= i+1; j-- {
			if minc > S[j] {
				midx, minc = j, S[j]
			}
		}

		if minc < S[i] {
			lidx, ridx = i, midx
			break
		}
	}

	if lidx == -1 && ridx == -1 {
		if string(S) < string(C) {
			fmt.Println(string(S))
		} else {
			fmt.Println(impossible)
		}
		return
	}

	S[lidx], S[ridx] = S[ridx], S[lidx]

	if string(S) < string(C) {
		fmt.Println(string(S))
	} else {
		fmt.Println(impossible)
	}
}
```

最初、最小のものを前から選ぶようなコードを書いてしまい、pretest2でWAしてしまって頭を抱えてしまいました。

何気なく `WAWA` という文字列を考えてたら（某プロゲーマーが浮かんだのかWAに怯えたのかは不明）
間違いに気づけたので良かったですが、
偶然以外何者でもないので、もう少し簡単なものでもいいから文字列問題に取り組むべきだなぁと思いました。

<a id="markdown-c-cut-and-paste" name="c-cut-and-paste"></a>
## C. Cut and Paste

[問題のURL](https://codeforces.com/contest/1281/problem/C)

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

`x` の値が初期状態の `S` の長さ以下であるときは素直にシミュレーションができる。

なぜなら、カットのあとすぐ行われるペーストは、切り出したものを少なくとも1回はペーストすることになるため、
`x` 文字目までは連続部分文字列が変わらず、長さだけに集中すれば良いからである。

ここで、 `x` の制約が `10^6` までであることを踏まえると、
`S` の長さが `x` に達するまでは文字列連結が必要だが、それ以降は不要となる。

後は実装を頑張るだけである。

```go
var t int
var x int64
var S []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		x = ReadInt64()
		S = ReadRuneSlice()

		solve()
	}
}

var ans int64
var clip int64

func solve() {
	ans = int64(len(S))
	clip = 0

	for i := int64(0); i < x; i++ {
		c := int64(S[i] - '0')
		cut := S[i+1:]

		clip = NegativeMod(ans-(i+1), MOD)
		clip %= MOD
		ans += (c - 1) * (clip)
		ans %= MOD

		// 長さが十分だったらそれ以上は連結しない
		if int64(len(S)) < x {
			// 最大c-1回分追加で連結する
		OUTER:
			for j := int64(0); j < c-1; j++ {
				for k := 0; k < len(cut); k++ {
					S = append(S, cut[k])

					// x以上になったら連結しない
					if int64(len(S)) >= x {
						break OUTER
					}
				}
			}
		}
	}

	fmt.Println(ans)
}

// NegativeMod can calculate a right residual whether value is positive or negative.
func NegativeMod(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}
```

添字が読みづらくて題意を把握するのに時間がかかってしまいました。

最近活用し始めましたが、Goのラベル付きfor文は競技プログラミングでは結構便利ですね
（実開発では闇雲に使うと怒られそうですが）。

<a id="markdown-d-beingawesomeism" name="d-beingawesomeism"></a>
## D. Beingawesomeism

[問題のURL](https://codeforces.com/contest/1281/problem/D)

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

問題文がえげつないが、要はサンプル等で示された動きで `A` で塗りつぶしたい、ということ。

以下の図のように、1マスでも `A` が存在すればせいぜい4回の動作で塗りつぶしが可能となることを踏まえる。



少ない回数ほど条件は厳しいと考えられるので、
0, 1, 2, 3, 4回それぞれを図示して考えていく。



まず、図示していないが、全てが `P` の場合はimpossibleである。
また、すべてが `A` の場合は当然0回である。

次に、1回で済む場合を考えると、これは端っこの列もしくは行がすべて `A` のときである。
反対側へ向かう1回の移動で塗りつぶしが可能となる。

次に、2回で済む場合は、四隅のいずれかが `A` である場合と、
ある行もしくはある列についてすべてが `A` の場合である。
図に示す矢印の方向に移動を行えば、塗りつぶしが可能であることがわかる。

最後に、四隅を除いた端っこの行もしくは列に `A` が1文字でも存在すれば、3回の移動で済む。

いずれにも当てはまらない場合は、はじめの図に示したような4回の移動が必要となる。

```go
var t int
var r, c int
var S [][]rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		r, c = ReadInt2()
		S = [][]rune{}
		for i := 0; i < r; i++ {
			row := ReadRuneSlice()
			S = append(S, row)
		}

		solve()
	}
}

func solve() {
	anum := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if S[i][j] == 'A' {
				anum++
			}
		}
	}
	if anum == 0 {
		fmt.Println("MORTAL")
		return
	}
	if anum == r*c {
		fmt.Println(0)
		return
	}

	if subRow(0) || subRow(r-1) || subCol(0) || subCol(c-1) {
		fmt.Println(1)
		return
	}

	b := false
	for i := 1; i < r-1; i++ {
		b = b || subRow(i)
	}
	for i := 1; i < c-1; i++ {
		b = b || subCol(i)
	}
	if b {
		fmt.Println(2)
		return
	}
	if S[0][0] == 'A' || S[0][c-1] == 'A' || S[r-1][0] == 'A' || S[r-1][c-1] == 'A' {
		fmt.Println(2)
		return
	}

	b = false
	for i := 1; i < c-1; i++ {
		b = b || (S[0][i] == 'A')
		b = b || (S[r-1][i] == 'A')
	}
	for i := 1; i < r-1; i++ {
		b = b || (S[i][0] == 'A')
		b = b || (S[i][c-1] == 'A')
	}
	if b {
		fmt.Println(3)
		return
	}

	fmt.Println(4)
}

func subRow(rowId int) bool {
	for i := 0; i < c; i++ {
		if S[rowId][i] != 'A' {
			return false
		}
	}
	return true
}

func subCol(colId int) bool {
	for i := 0; i < r; i++ {
		if S[i][colId] != 'A' {
			return false
		}
	}
	return true
}
```

pretestが弱くてなかなかの地獄だったようですが、
コンテスト外とはいえ1発で通ったので良かったです（残業後だったので通らなかったらやはり地獄だった）。

Cを解き終わった時点で満身創痍だったのでコンテスト中は触れられませんでしたが、
やっておくべきでした。

