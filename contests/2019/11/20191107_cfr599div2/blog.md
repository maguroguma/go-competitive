B2が本当にわからなかった。

Cは解けたけど、この手の問題はつい最近もこどふぉで出会ったので、もう少し筋よく考えてさっと答えたいところ。

<!-- TOC -->

- [A. Maximum Square](#a-maximum-square)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B1. Character Swap (Easy Version)](#b1-character-swap-easy-version)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [B2. Character Swap (Hard Version)](#b2-character-swap-hard-version)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [C. Tile Painting](#c-tile-painting)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)
	- [公式editorialの証明](#%e5%85%ac%e5%bc%8feditorial%e3%81%ae%e8%a8%bc%e6%98%8e)

<!-- /TOC -->

<a id="markdown-a-maximum-square" name="a-maximum-square"></a>
## A. Maximum Square

[問題URL](https://codeforces.com/contest/1243/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

（問題ページの図が詳しいのでそちらをご参照ください。）

`n` 個の `a[i] * 1` の縦長の板が与えられるので、それらの好きな組み合わせを好きな順番で横にくっつける。
上の出っ張ったところを切ってできるだけ大きい正方形を作りたい。

最大で一辺の長さはいくらにできるか、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

一辺の長さを `x` にしようと思うと、縦の長さが `x` 以上の板が `x` 枚以上必要となる。
板の縦の長さの最大値が `1000` であることを踏まえ、配列を使って各板の長さについて枚数を数えて記憶しておく。
この配列に対して、反対側から累積和を計算することで、ある長さ以上の板の枚数を `O(1)` で取得できる。

最大の値を答えることが目的なので、大きいところからチェックしていけば良い。
全体で `O(n)` で解ける。

```go
var k int
var n int
var A []int

func main() {
	k = ReadInt()

	for tc := 0; tc < k; tc++ {
		n = ReadInt()
		A = ReadIntSlice(n)
		solve()
	}
}

func solve() {
	cnt := [1005]int{}
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	sums := make([]int, 1005)
	for i := 1000; i >= 0; i-- {
		sums[i] = sums[i+1] + cnt[i]
	}

	for i := 1000; i >= 0; i-- {
		if sums[i] >= i {
			fmt.Println(i)
			return
		}
	}
}
```

テストケースも少ないのでもっと愚直にやってもいいと思うけど、
特に思いつきませんでした。

<a id="markdown-b1-character-swap-easy-version" name="b1-character-swap-easy-version"></a>
## B1. Character Swap (Easy Version)

[問題URL](https://codeforces.com/contest/1243/problem/A)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

長さ `n` の文字列 `S, T` が与えられる。
`S, T` は異なることが保証される。

この文字列に対して、ある `1 <= i, j <= n` について `S` の `i` 文字目と `T` の `j` 文字目を交換して良い。
`i, j` は同じでも異なってもどちらでも良いが、必ず一度交換する必要がある。

文字列 `S, T` を等しくできるかどうか判定する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

`S, T` について、同じ位置の文字を比較したときの異なる個数を考える。
これが `1` 個の場合や、 `3` 個以上の場合は1回のみの交換ではどのようにしても等しくすることはできない。

よって、文字列を等しくできる可能性があるのは、異なる個数がちょうど `2` 個の場合である。

そしてこのような状況で考えられる交換は、異なる位置の番号を小さい順に `i, j` とすると、

- `S` の `i` 文字目と `T` の `j` 文字目の交換
- `S` の `j` 文字目と `T` の `i` 文字目の交換

のいずれかである。
この交換の後に `S, T` が等しくなるための条件はそれぞれ、

- `T[i]` と `T[j]` を比較することになるため `T[i] == T[j]` かつ、 `S[j]` と `S[i]` を比較することになるため `S[j] == S[i]`
- `S[i]` と `S[j]` を比較することになるため `S[i] == S[j]` かつ、 `T[i]` と `T[j]` を比較することになるため `T[i] == T[j]`

結局条件はどちらも同じなので、この条件をチェックすれば良い。

```go
var k int
var n int
var S, T []rune

func main() {
	k = ReadInt()
	for tc := 0; tc < k; tc++ {
		n = ReadInt()
		S = ReadRuneSlice()
		T = ReadRuneSlice()
		solve()
	}
}

func solve() {
	diff := 0
	for i := 0; i < n; i++ {
		if S[i] != T[i] {
			diff++
		}
	}

	if diff == 0 {
		fmt.Println("Yes")
		return
	}
	if diff != 2 {
		fmt.Println("No")
		return
	}

	pos := [2]int{}
	j := 0
	for i := 0; i < n; i++ {
		if S[i] != T[i] {
			pos[j] = i
			j++
		}
	}

	i, j := pos[0], pos[1]
	if S[i] == S[j] && T[i] == T[j] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
```

<a id="markdown-b2-character-swap-hard-version" name="b2-character-swap-hard-version"></a>
## B2. Character Swap (Hard Version)

[問題URL](https://codeforces.com/contest/1243/problem/B2)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

設定はEasyと似ているが、制約がまず全く異なる。

- テストケース `1 <= k <= 1000`
- 文字列長 `2 <= n <= 50`

また、操作回数は `2*n` 回までなら何回でもよい。

さらに、判定結果だけではなく、等しくできるのであればその構築手順まで出力する必要がある。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

※コンテスト中全くわからなかったので、公式editorialの内容そのままです。

まず、等しくできるための必要条件として「すべてのアルファベットについて、 `S, T` に渡って登場回数が偶数であること」が挙げられる
（あるアルファベットが奇数個である場合、 `S, T` のいずれかで、どこかの位置で最終的に余ってしまう）。

また、必要条件が満たされる場合、以下のような手順で文字列を等しくすることができる。

> `i = 1...n` の `i` について、 `S[i] != T[i]` である場合、以下のいずれかが必ず成り立つので、それに応じた操作を行う。
> `j > i` の `j` について、 `S[i] == S[j]` となる場合、 `S[j], T[i]` を交換する。
> あるいは、 `j > i` の `j` について、 `S[i] == T[j]` となる場合、まず `S[j], T[j]` を交換し、その後 `S[j], T[i]` を交換する。
> それぞれの操作を行うことで、 `i` 番目の文字を揃えることができる。
> これを `n` まで行えば文字列を互いに等しくできる。

それぞれの操作について最大で `2` 回までで住むため、 `2*n` という操作回数はこれを行うために十分である。

```go
const ALPHABET_NUM = 26

var k int
var n int
var S, T []rune

func main() {
	k = ReadInt()
	for tc := 0; tc < k; tc++ {
		n = ReadInt()
		S = ReadRuneSlice()
		T = ReadRuneSlice()
		solve()
	}
}

func solve() {
	memo := make([]rune, ALPHABET_NUM)
	for i := 0; i < n; i++ {
		s, t := S[i], T[i]
		memo[s-'a']++
		memo[t-'a']++
	}

	for i := 0; i < len(memo); i++ {
		if memo[i]%2 == 1 {
			fmt.Println("No")
			return
		}
	}

	answers := [][]int{}
	for i := 0; i < n; i++ {
		if S[i] == T[i] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if S[i] == S[j] {
				S[j], T[i] = T[i], S[j]
				answers = append(answers, []int{j, i})
				break
			} else if S[i] == T[j] {
				S[j], T[j] = T[j], S[j]
				S[j], T[i] = T[i], S[j]
				answers = append(answers, []int{j, j})
				answers = append(answers, []int{j, i})
				break
			}
		}
	}

	fmt.Println("Yes")
	fmt.Println(len(answers))
	for i := 0; i < len(answers); i++ {
		fmt.Printf("%d %d\n", answers[i][0]+1, answers[i][1]+1)
	}
}
```

Easyバージョンの解法に囚われすぎて、一歩引いて考えることができませんでした（異なった位置だけ考えると構築がエグい。。とか考えてしまいました）。

- 明らかな必要条件を整理してみる
- 操作回数が固定だったことをもう少し考えてみる

落ち着いてこのあたりができればもう少し違った結果だったかもしれません。

Easyバージョンに引っ張られすぎてHardの思考の幅が狭くなる、というのは以前にも合ったので、もう少し意識的に取り組んだほうが良さそうです。

<a id="markdown-c-tile-painting" name="c-tile-painting"></a>
## C. Tile Painting

[問題URL](https://codeforces.com/contest/1243/problem/C)

<a id="markdown-問題の要約-3" name="問題の要約-3"></a>
### 問題の要約

`n` 枚の連続するタイルがあり、左から `1, 2, ..., n` と採番されている。
これらのタイルを複数の色で塗り分けることを考える。

任意のタイル `i, j` について、 `|j - i|` が `n` の `1` 以外の約数である場合、タイル `i, j` は互いに同じ色である場合に、
`n` 枚のタイルは「芸術的」であるとする。

「芸術的」な塗り方を考えたときに、最大で何色の色で塗り分けることができるか求めよ、という問題。

制約: `1 <= n <= 10^12`

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

※コンテスト中に確信を持って解けたわけではなく、本節の内容は思考過程の整理と反省という意味合いが強いです。

問題文を理解するのがちょっと大変だった。

サンプルを見ると、素数 `5` に関しては `n` がそのまま答えとなっている。
実際に、あるタイル `i` に素数を加算すると存在しないタイルの番号になるため、すべてのタイルを異なる色で塗ることが可能だとわかる。

制約的にも素数判定が `O(sqrt(n))` で許されるため、とりあえず早期returnで書き出しておく。

他の場合はどうなるか？
1つ目のサンプル `4` について考えると、2色で塗り分けることが可能となっている。
なんとなく、平方数だと同じようなことが言えそうだとわかり、適当に実験してみると、そもそも素因数分解したときに1つの素数で分解できるならば、
同じように `n = p^q` であれば `p` が答えとなる、とわかる。

では、素因数分解した結果がそれ以外の場合（複数の素数からなる場合）はどうか？
まず、 `1` が答えになりそうだと予想が立てられる。
感覚的には、素数同士は互いに素であるため、例えばタイル `1, 2` を考えたとき、その素数を適当に加算した組み合わせはどこかで衝突しそう、というもの。
実際にいくつか試して衝突は確認した。

上述の衝突が `n` 以下で必ず起こると主張できれば確信を持って提出できるが、正直コンテスト中はこれ以上詰めきれなかった。
やたらとAC者数が多かったので思い切って投げたら、そのコードが最終的にsystem testもACとなった。

素数判定や素因数分解（試し割り法）は、ともに `O(sqrt(n))` であるため、計算量は問題ない。

```go
var n int64

func main() {
	n = ReadInt64()

	if n == 1 {
		fmt.Println(n)
		return
	}

	if IsPrime(n) {
		fmt.Println(n)
		return
	}

	memo := TrialDivision(n)
	if len(memo) == 1 {
		for k := range memo {
			fmt.Println(k)
			return
		}
	}

	fmt.Println(1)
}

// TrialDivision returns the result of prime factorization of integer N.
func TrialDivision(n int64) map[int64]int {
	if n <= 1 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}

	p := map[int64]int{}
	for i := int64(2); i*i <= n; i++ {
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

// IsPrime judges whether an argument integer is a prime number or not.
func IsPrime(n int64) bool {
	if n == 1 {
		return false
	}

	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
```

`O(sqrt(n))` が許される場合は、素数判定・約数列挙、素因数分解も同時に手段として考慮すべきなのかもしれません。

<a id="markdown-公式editorialの証明" name="公式editorialの証明"></a>
### 公式editorialの証明

まず、 `n = p^q` と表せる場合は、 `i, j <= n` の異なるタイル `i, j` について、
`i` と同じ色で塗るべきタイルの番号は `i + k*p` 、また `j` と同じ色で塗るべきタイルの番号は `j + k'*p` と表すことができる。
それぞれを `p` で割ったあまりは `i, j` となるため、 `|(i+k*p) - (j+k'*p)|` は `p` の倍数とはならず、よって `n` の約数とはなりえない。
そのため、 タイル `i, j` から好きな `n` の約数分番号を進めても決して衝突することはないので、すべて異なる色で塗り分けられる。

また、 `n` が2つ以上の異なる素数からなる場合、[中国の剰余定理](https://mathtrain.jp/chinese)によって、
`i, j <= n` の異なるタイル `i, j` が `n` の約数分番号を進めたときに、 `n` 以下で必ず衝突すると主張できる。

以下は2元の場合の中国の剰余定理。

> `gcd(n1, n2) = 1` のとき、連立合同式合同式 `x = a (mod n1), x = b (mod n2)` を満たす `x` が、
> **`[0, n1*n2)` の範囲にただ1つ存在する。**

ここで、 `gcd(p, q) == 1` を満たす `p, q >= 2` を用いて、 `n = p*q` と表す（ `n` を素因数分解したときに2つ以上の異なる素数からなる場合には、必ずこのような表現が可能である）。
また、適当なタイルの番号 `a, b <= n` について考えると、上記の連立合同式を満たす `x` が `n` 以下（実際にはもっと狭い範囲）に必ず存在すると言える。
よって、タイル `a, b` からそれぞれに対して固有の適当な回数分 `n1, n2` すすめると、あるタイル `x <= n` で衝突することとなる。
そのため、すべてのタイルは同じ色で塗る必要がある。

---

整数問題が得意になる気配が感じられません。
