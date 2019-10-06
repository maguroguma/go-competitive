<a id="markdown-全体" name="全体"></a>
## 全体

結果は前回と同じくA, Bの2完でしたが、通すべきCを通せなかったので順位もパフォーマンスも惨憺たる事に。。

Dも難易度的にはちょうど良さそうで解いてみたいのですが、
方針としてコンテスト中の集中力がある状態で目を通した問題のみ時間を書けて復習する、としているので、とりあえずはCまでを解きます。

<!-- TOC -->

- [全体](#%e5%85%a8%e4%bd%93)
- [A. Distinct Digits](#a-distinct-digits)
	- [問題](#%e5%95%8f%e9%a1%8c)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Filling the Grid](#b-filling-the-grid)
	- [問題](#%e5%95%8f%e9%a1%8c-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Primes and Multiplication](#c-primes-and-multiplication)
	- [問題](#%e5%95%8f%e9%a1%8c-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [感想](#%e6%84%9f%e6%83%b3)

<!-- /TOC -->

---

<a id="markdown-a-distinct-digits" name="a-distinct-digits"></a>
## A. Distinct Digits

[問題URL](https://codeforces.com/contest/1228/problem/A)

<a id="markdown-問題" name="問題"></a>
### 問題

以下の2つの条件を満たす `x` が存在すれば、どれでも良いので1つ出力せよ。
存在しない場合は `-1` を出力せよ。

- `l <= x <= r`
- `x` のすべての桁の数が異なる

制約: `1 <= l <= r <= 10^5`

<a id="markdown-解答" name="解答"></a>
### 解答

調べるべき範囲が十分狭いので、全探索を行う。

```go
var l, r int

func main() {
	l, r = ReadInt2()

	for i := l; i <= r; i++ {
		if sub(i) {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}

// すべての桁の数字が異なるかどうか？
func sub(n int) bool {
	nn := n
	memo := [10]int{}

	for n > 0 {
		memo[n%10] = 1
		n /= 10
	}

	res := 0
	for i := 0; i < 10; i++ {
		res += memo[i]
	}

	if res == decimalLength(nn) {
		return true
	}
	return false
}

// nの10進数表現の桁数
func decimalLength(n int) int {
	res := 0
	for n > 0 {
		res++
		n /= 10
	}
	return res
}
```

なんということは無い問題なんですが、個人的にこのような繰り返し整数除算するコードを書くのなんか苦手です（素因数分解とか）。

このときもやたらと石橋を叩いて時間を使いすぎたので、そろそろ自信を持ちたい。。

<a id="markdown-b-filling-the-grid" name="b-filling-the-grid"></a>
## B. Filling the Grid

[問題URL](https://codeforces.com/contest/1228/problem/B)

<a id="markdown-問題-1" name="問題-1"></a>
### 問題

※長いので意訳

各行・列に与えられた数値が1つのみのロジックパズル（マリオのピクロスみたいなやつ）について、
完成後のパターンとして考えられるものの数を `10^9+7` で割ったあまりで答えよ。

制約: `1 <= h, w <= 10^3`

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

素直にシミュレーションを行う。

行方向について決定するときは素直に決めればよいが、
続いて縦方向について決定するときは、矛盾のチェックを逐次行う。

矛盾が最後までなければ、未確定のマスの数の分だけ2倍していき、都度MODを取れば良い。

```go
var h, w int
var R, C []int

var cells [1000 + 5][1000 + 5]int  // 1: 黒, 0: 白
var dones [1000 + 5][1000 + 5]bool // true: 確定

func main() {
	h, w = ReadInt2()
	R = ReadIntSlice(h)
	C = ReadIntSlice(w)

	// 行は無責任に決める
	for i := 0; i < h; i++ {
		r := R[i]
		for j := 0; j < r; j++ {
			cells[i][j] = 1
			dones[i][j] = true
		}

		// 白で確定させる
		cells[i][r] = 0
		dones[i][r] = true
	}

	// 列はチェックしながら決める
	for j := 0; j < w; j++ {
		c := C[j]
		for i := 0; i < c; i++ {
			if dones[i][j] {
				// すでに決定済みかつ、白でないとダメなら矛盾
				// 黒だったならOK
				if cells[i][j] == 0 {
					fmt.Println(0)
					return
				}
			} else {
				cells[i][j] = 1
				dones[i][j] = true
			}
		}

		// 白で確定させる
		if dones[c][j] {
			// すでに決定済みかつ、黒でないとダメなら矛盾
			// 白だったならOK
			if cells[c][j] == 1 {
				fmt.Println(0)
				return
			}
		} else {
			cells[c][j] = 0
			dones[c][j] = true
		}
	}

	ans := int64(1)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !dones[i][j] {
				ans *= int64(2)
				ans %= int64(MOD)
			}
		}
	}
	fmt.Println(ans)
}
```

これまた途中でバグを生んでしまったり、ものすごく時間を使ってしまった。

何気なく `i, j` をイテレータとして使うにしても、列方向は `j` で固定するなど、
こういった部分で混乱を減らし、極力バグを生みにくくなるように工夫していきたい。

<a id="markdown-c-primes-and-multiplication" name="c-primes-and-multiplication"></a>
## C. Primes and Multiplication

[問題URL](https://codeforces.com/contest/1228/problem/C)

<a id="markdown-問題-2" name="問題-2"></a>
### 問題

※これも長いので省略

与えられた定義のもとで `f(x, 1) * f(x, 2) * ... * f(x, n) % (10^9+7)` を求めよ。

制約: `2 <= x <= 10^9, 1 <= n <= 10^18`

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

自分の解答が非常に説明しづらいので、大部分を図示する。

<figure class="figure-image figure-image-fotolife" title="Cの解法">[f:id:maguroguma:20191005033448j:plain]<figcaption>Cの解法</figcaption></figure>

素直に定義式に従って与えられた式を分解していくと、たくさんの `g(?, ?)` の積を計算することになる。
特に `n` が非常に大きく、これをまともに全部計算するわけには行かないので、図示した赤枠の縦方向について効率的に計算できないか考える。
求め方は図中の文章通りで、説明の都合上ところどころ一般化して文字式で置いたりしているが、
コンテスト中は適宜適当な数値を当てはめて考えるとわかりやすいと思う（「例示は理解の試金石」((結城浩さんの数学ガールに登場するらしいですね。いい言葉だと思います。))）。

指数の計算をする部分は、指数部が大きすぎるので二分累乗法が必要。
逆元のスニペットを作るときに同時に作ることになると思うので、一緒に揃えておきたい。

```go
var x, n int64

func main() {
	x, n = ReadInt64_2()

	primes := TrialDivision(int(x))

	ans := int64(1)
	for p := range primes {
		prod := int64(1)
		for {
			if isOverflow(prod, int64(p)) {
				break
			}

			prod *= int64(p)
			if n < prod {
				break
			}
			num := int64(n / prod)

			tmp := modpow(int64(p), num, MOD)
			ans *= tmp
			ans %= MOD
		}
	}

	fmt.Println(ans % MOD)
}

func isOverflow(i, j int64) bool {
	return !(i < math.MaxInt64/j)
}

func modpow(a, e, m int64) int64 {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := modpow(a, halfE, m)
		return half * half % m
	}

	return a * modpow(a, e-1, m) % m
}

// TrialDivision returns the result of prime factorization of integer N.
func TrialDivision(n int) map[int]int {
	if n <= 1 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}

	p := map[int]int{}
	for i := 2; i*i <= n; i++ {
		exp := 0
		for n%i == 0 {
			exp++
			n /= i
		}

		if exp == 0 {
			continue
		}
		p[i] = exp
	}
	if n > 1 {
		p[n] = 1
	}

	return p
}
```

上記のコードでオーバーフローチェックをしていますが、本番中に40分かけてもこれに気づけませんでした。。

初歩的とはいえ、このようなつまり方をしたのは初めてだったので、早い段階で自分の中の地雷処理ができたとポジティブに考えておきます。。

---

<a id="markdown-感想" name="感想"></a>
## 感想

百歩譲って未経験のタイプのオーバーフローを解決できなかったのはともかく、
A，Bの実装が遅すぎたのも問題でした。

競技プログラミングの筋トレが全く足りていないので、引き続きCodeforcesで鍛えていきたいと思います。
