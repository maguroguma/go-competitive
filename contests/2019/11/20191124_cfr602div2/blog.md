
<!-- TOC -->

- [A. Math Problem](#a-math-problem)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Box](#b-box)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Messy](#c-messy)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D1. Optimal Subsequences (Easy Version)](#d1-optimal-subsequences-easy-version)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-math-problem" name="a-math-problem"></a>
## A. Math Problem

[問題のURL](https://codeforces.com/contest/1262/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

`n` 個の数直線上の区間が与えられる。

ここに、ある区間1つを加える。
ただし、この区間は `n` 個の与えられたすべての区間と、少なくとも1つの共有点（いずれの区間にも含まれる点）を持たなければならない。

加えるべき区間の最小の長さを答えよ、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

区間を並べた時の様子をイメージする。

このとき、最小の右端をとる区間が少なくとも1つ存在する（複数存在する場合もある）。
これを `l` とする。
また、最大の左端をとる区間が少なくとも1つ存在する（複数存在する場合もある）。
これを `r` とする。

求めるべき区間は、区間 `l, r` と共有点を保つ必要があるため、
少なくとも `l` の右端と `r` の左端は区間の端点とする必要がある。
このようにして出来上がる区間は、他の `n-2` 個の区間とも自然と共有点を持つことになるため、
この区間の長さが答えとなる。

区間の長さは、最小の右端を `minR` 、最大の左端を `maxL` とすると、
`maxL - minR` で求まる。

ただし、この値が負になる場合は、最小の右端の値が最大の左端の値を上回っている。
このときは、 `[maxL, minR]` の任意の値が、 `n` 個すべての区間と共有点を持つことになるため、答えは `0` とすればよい。

```go
var t int
var n int
var L, R []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		L, R = make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			l, r := ReadInt2()
			L[i], R[i] = l, r
		}
		solve()
	}
}

func solve() {
	minR, maxL := 1000000000+5, 0
	for i := 0; i < n; i++ {
		ChMin(&minR, R[i])
		ChMax(&maxL, L[i])
	}
	fmt.Println(Max(maxL-minR, 0))
}
```

<a id="markdown-b-box" name="b-box"></a>
## B. Box

[問題のURL](https://codeforces.com/contest/1262/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

ある `1, 2, ..., n` の `n` 個の異なる数値の順列 `P` が、箱を開けるためのコードとなっている。

`P` はわからないが、情報として prefix maximums `Q` が与えられている。

※prefix maximumsは、 `q[1] = p[1], q[2] = max(p[1], p[2]), q[3] = max(p[1], p[2], p[3]), ...` のように定義されるもの。

これを元に、元の順列 `P` としてありうるものを1つ構築せよ、という問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

まず、prefix maximumsの定義から、 `P[1] == Q[1]` であるとわかる。
また、 `i >= 2` に関して、 `Q[i-1] < Q[i]` となる場合は、 `P[i] == Q[i]` であるとわかる。

これらの自明な情報から、 `P` に関して確定した部分だけを先に埋めておく。
そして、未確定の部分を前から順番に確定させていくこと、および矛盾について考える。

前から順番に見るときに、未確定部分に選べる数値に関する必要条件は、

1. `Q[i]` 以下であること
2. `[1, n]` の数値のうち、未使用のものであること

の2つである。

よって、未使用の `[1, n]` の数値かつ `Q[i]` 以下のものを選ぶことになる。
条件に当てはまるものを適当に全探索してしまうと、 `O(n^2)` になってしまうため、少し工夫が必要となる。

実は、未使用のもののなかから選ぶ際には、小さいものから順番に選んでよい。
これは、以下のように考えることで正当性を示せる。

> ある条件に合致する `P` について、自明パートを終えた後の未確定部分についての部分列を考える。
> この部分列は、自明パートで残った未使用の数値の順列となる。
> prefix maximumsの定義から `Q[j] >= Q[i] (j > i)` と単調増加するため、
> この部分列については昇順ソートしたものも必ず条件を満たす。

小さいものから割り当てていき、prefix maximumsの条件に抵触してしまった場合は、どうあがいても条件を満たす `P` は作れないので、 `-1` を出力する。

以上の方針に従って実装すれば良い。
自明パートを終えた後の部分は、未使用の要素を小さい順にqueueに放り込むのが簡単だと思う。

```go
var t int
var n int
var Q []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		Q = ReadIntSlice(n)

		solve()
	}
}

func solve() {
	memo := make([]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = 1
	}

	answers := make([]int, n)
	answers[0] = Q[0]
	memo[Q[0]] = 0
	for i := 1; i < n; i++ {
		if Q[i-1] < Q[i] {
			answers[i] = Q[i]
			memo[Q[i]] = 0
		}
	}

	// 未使用のものを小さい順にスライスに詰める
	unused := []int{}
	for i := 1; i <= n; i++ {
		if memo[i] == 1 {
			unused = append(unused, i)
		}
	}

	for i := 0; i < n; i++ {
		// 未割り当てならば、小さいものを割り当てる、矛盾したらアウト
		if answers[i] == 0 {
			// 次の未使用のものを取り出す
			c := unused[0]
			unused = unused[1:]
			if Q[i] >= c {
				answers[i] = c
			} else {
				fmt.Println(-1)
				return
			}
		}
	}

	fmt.Println(PrintIntsLine(answers...))
}
```

途中の貪欲法でよくあるタイプの証明は、コンテスト中にはここまではっきりとは言語化できていなかったので、もう少しこういった証明には慣れたいところです。

また本番では、メモ用配列をそのまま使って未使用のもののポインタを都度移動させるという、しゃくとり法っぽい実装をしてしまいました。
しゃくとり法に慣れていないのが悪いんですが、ちょっと実装に手間取ってしまったので、もう少し立ち止まってからコーディングすべきでした。

<a id="markdown-c-messy" name="c-messy"></a>
## C. Messy

[問題のURL](https://codeforces.com/contest/1262/problem/A)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

偶数の値 `n` 文字の丸括弧列 `S` が与えられる。
この丸括弧列の `(, )` の数は等しいものとする。

ここで、1回の操作で `S` の `l` 文字目から `r` 文字目までの部分を反転させることができる。
※この反転という操作は、選んだ区間の括弧の種類を逆転させるという意味ではなく（そうすると開き括弧と綴じ括弧の総数が変化してしまう）、
列の順番の反転のことである。

また、 `k` が同時に与えられるため、 `S` のprefixのうち `k` 個のprefixが正しい括弧列となるようにしたい。

操作が最大で `n` 回可能であるときに、 `k` 個のprefixが正しい括弧列となるようにするための、操作手順を提示せよ、という問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

まず、 `n` 回が操作の上限であることを見逃してしまい、難問のように感じてしまったのが大反省（Outputの `m` という数値に途中から引っ張られてしまった？）。

一見、制約が強そうに見える反転という操作も、操作回数が十分であれば、どんな列にでも組み替えてしまえる、ということを把握したい
（前から順番に注目して、そこに置きたい文字を後ろのほうを探索して見つけ、そこを始点・終点として反転すれば、始点については欲しい文字が手に入る、
これを以降の文字についても同様に行えば、任意のほしい文字列が手に入る）。

さて、目的の括弧列としては `()()()...(((...)))` の形とすればよい。

前半部分の構築方法としては、
注目している文字のインデックスの偶奇からあるべき括弧の種類を調べ、
一致しているのであればスルー、不一致であれば後ろの方に存在するあるべき括弧を見つけ、始点と終点で反転すれば良い。

また、後半部分の構築方法としては、
前半分が `(` となるように、後半分が `)` となるように、同様の反転処理を行えば良い。

前半の `()()...()` は、 `()` が合計 `k-1` 個作る必要がある。
そのため、前半部分の長さは `2*(k-1)` となる。
前半の長さと全体長 `n` から後半の長さもわかるため、上述の前半分・後半分の処理を分けることができる。

反転処理は、別に関数として切り分けておくと再利用性も高くスッキリする。

```go
var t int
var n, k int
var S []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n, k = ReadInt2()
		S = ReadRuneSlice()

		solve()
	}
}

func solve() {
	answers := [][]int{}

	for i := 0; i < 2*(k-1); i++ {
		if i%2 == 0 && S[i] == ')' {
			answers = append(answers, sub(i, '('))
		} else if i%2 == 1 && S[i] == '(' {
			answers = append(answers, sub(i, ')'))
		}
	}

	tmp := (n - 2*(k-1)) / 2
	for i := 2 * (k - 1); i < 2*(k-1)+tmp; i++ {
		if S[i] == ')' {
			answers = append(answers, sub(i, '('))
		}
	}

	fmt.Println(len(answers))
	for i := 0; i < len(answers); i++ {
		fmt.Println(answers[i][0]+1, answers[i][1]+1)
	}
}

// sからスタートしてrを発見したところで反転する
// さらに始点と終点をスライスで返す
func sub(s int, r rune) []int {
	t := s + 1
	for i := s + 1; i < n; i++ {
		if S[i] == r {
			t = i
			break
		}
	}
	res := []int{s, t}

	// [s, t]を反転
	rev := Reverse(S[s : t+1])
	for i := 0; i < len(rev); i++ {
		S[s+i] = rev[i]
	}

	return res
}

func Reverse(A []rune) []rune {
	res := []rune{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}
```

`Reverse` みたいなのは地味にスニペットに用意しておくと、多少は快適です。

<a id="markdown-d1-optimal-subsequences-easy-version" name="d1-optimal-subsequences-easy-version"></a>
## D1. Optimal Subsequences (Easy Version)

[問題のURL](https://codeforces.com/contest/1262/problem/D1)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

長さが `n` のある数列 `A` が与えられる。
この数列 `A` の長さ `k` の部分列を考える。

以下の条件を満たすとき、この部分列はoptimalであるとする。

1. 考えられる任意の長さ `k` の部分列のうち、部分列を構成する要素の和が最も大きい。
2. 1を満たす部分列の中で、辞書順最小である。

`m` 個のリクエストが与えられるので、それらすべてに答える。

リクエスト `j` は `(k[j], pos[j])` で与えられる。
これに対して、長さ `k[j]` のoptimalな部分列に対し、1-basedで位置 `pos[j]` の数値で答える。

Easy Versionの制約

- `1 <= n, m <= 100`

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

すべての `k (1 <= k <= 100)` について、optimalな部分列を前計算しておくことを考える。

制約が小さく制限時間も3secと長いので、よっぽど変なことをしない限りTLEしないだろうということで、
素直に考えていく。

まず、1つ目の条件を満たす必要があることを考えると、 `A` を降順ソートしたとき、
先頭 `k` 個の要素を部分列が含む必要がある。
optimalな部分列が含むべき要素がわかったので、これが辞書順最小となるような部分列を取得することを考える。

降順ソート後の `A` の先頭 `k` 個の配列を `topk` とする。
`topk` の構成要素と一致する辞書順最小の部分列を得るには、 `topk` の構成要素をできるだけ元の配列 `A` の前の方から貪欲に選択すれば良い。

以下のコードは、ソートや愚直な全探索を駆使して、ある `k` についてのoptimalな部分列を計算している。
`subsub` 関数が `topk` から辞書順最小の部分を求める関数で、ここが `O(kn)` となり一番ネックとなる部分である。
`k` は `[1, n]` 全てについて求めているため、全体で `O(n^3)` となる。

```go
var n int

var A []int
var m int
var k, pos int

var answers [105][]int

var sA []int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)
	m = ReadInt()

	sA = make([]int, n)
	for i := 0; i < n; i++ {
		sA[i] = A[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sA)))

	sub()

	for i := 0; i < m; i++ {
		k, pos = ReadInt2()
		pos--

		fmt.Println(answers[k][pos])
	}
}

func sub() {
	for k := 1; k <= n; k++ {
		answers[k] = make([]int, k)

		topk := sA[:k]
		tmp := subsub(k, topk)

		sort.Sort(sort.IntSlice(tmp))

		for i := 0; i < k; i++ {
			answers[k][i] = A[tmp[i]]
		}
	}
}

// Aのidxの配列を返す
func subsub(k int, topk []int) []int {
	res := make([]int, k)
	memo := make([]bool, n)
	for i := 0; i < k; i++ {
		// Aと照合させる
		for j := 0; j < n; j++ {
			if topk[i] == A[j] && !memo[j] {
				res[i] = j
				memo[j] = true
				break
			}
		}
	}
	return res
}
```

---

最近誤読というより重要な制約とかの見落としが多いので、強めに意識しないといけない。
