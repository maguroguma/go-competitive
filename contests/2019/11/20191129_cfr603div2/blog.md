1問目からいきなり難問を置かないでください。。
青が見えていたのにまた遠のいてしまいました。

<!-- TOC -->

- [A. Sweet Problem](#a-sweet-problem)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. PIN Codes](#b-pin-codes)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Everyone is a Winner!](#c-everyone-is-a-winner)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Secret Passwords](#d-secret-passwords)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-sweet-problem" name="a-sweet-problem"></a>
## A. Sweet Problem

[問題のURL](https://codeforces.com/contest/1263/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

赤、緑、青色のキャンディがそれぞれ `r, g, b` 個ある。

1日に、これらのキャンディから、必ず異なった色のキャンディを1つずつ食べる必要がある。

最大何日食べることができるか計算せよ、という問題。

制約: `1 <= r, g, b <= 10^8`

<a id="markdown-解答" name="解答"></a>
### 解答

場合分けして考える。

まず、 `r, g, b` を降順ソートして大きい順に `a, b, c` とする。

`a >= b + c` の場合を考える。

この場合は、 `b + c` が最適である。
この場合には `a` のキャンディが無限にあると考えても差し支えないので、
常に片方は `a` のキャンディを選び、もう片方は `b, c` の残っている方のキャンディを選べばよい。
`b, c` を選ぶ食べ方は損をするだけとなる。

次に、 `a < b + c` の場合を考える。

この場合は、すべてのキャンディを食べ尽くす方法と、
1つのキャンディが残る食べる方法のいずれかが存在する。
これは、明らかに最適な食べ方である。

まず、 `a - b = diff` とし、この差の分だけ `a` から減らし、 `c` を一緒に選ぶ。
`a - b < c` から、これは必ず可能である。
すると、キャンディの分布は `(a, b, c) -> (b, b, c-diff)` となる。
この後は、

1. `a, b` のキャンディでより多く残っている方、およびもう一方を `c` から選ぶ
2. `c` が尽きているならば `a, b` をそれぞれ選ぶ

この手順は `a, b` のキャンディの数が等しい状態からはじまるので、
`c%2 == 0` ならば `c` が尽きたときには `a, b` が等しくなっているため、すべてのキャンディを食べられる。
`c%2 == 1` ならば `c` が尽きたときには `Abs(a-b) = 1` であるため、1つキャンディが残ることになる。

```go
var t int
var r, g, b int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		r, g, b = ReadInt3()

		solve()
	}
}

func solve() {
	A := []int{r, g, b}
	// 降順ソート
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	sum := A[1] + A[2]
	if sum <= A[0] {
		fmt.Println(sum)
	} else {
		if A[0] > A[1] {
			ans := 0

			diff := A[0] - A[1]
			A[0] -= diff
			ans += diff
			A[2] -= diff
			tmp := (A[2] + (2 - 1)) / 2
			ans += A[2] + A[0] - tmp
			fmt.Println(ans)
		} else {
			tmp := (A[2] + (2 - 1)) / 2
			ans := A[2] + A[0] - tmp
			fmt.Println(ans)
		}
	}
}
```

1問目にしては難しすぎでは？
それにしてもみんなさっと解いていたので驚きました。

<a id="markdown-b-pin-codes" name="b-pin-codes"></a>
## B. PIN Codes

[問題のURL](https://codeforces.com/contest/1263/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

各桁が `0, 1, ..., 9` のいずれかであり、4桁のPINコードを持つ銀行カードが `n` 枚ある。

今、 `n` 枚全てのカードのPINコードについて、すべての異なるカードのペアでPINコードが異なるようにしたい。

あるカードのPINコードの1桁を任意の数字に変更するときの操作回数を1回とするとき、
最小何回の操作で目的を達成できるか求めよ、という問題。

制約: `2 <= n <= 10`

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

カードに `0-based` で番号をふる（入力で与えられる順に自然に採番する）。

あるカードに着目して、それより前の番号のカードすべてと異なるようにすることを考える。

カードの数は少ないので、 `O(n)` の全探索で前のカードをすべてチェックすれば良い。
もし、前の番号のカードと一致するならば必ず少なくとも1桁は変更する必要がある。
このときに変更する桁は、PINコードの最下位桁とする。
ここで、変更後のPINコード全体が、現在注目中のカードの番号以降のカードのPINコード全体と衝突しないように、
すでに使用されている最下位桁の数値をメモとして保持しておく。
未使用の数字を変更後の最下位桁として採用すれば、いずれのカードとも衝突することはない。
また、重要な制約として、カードは最大でも `10` 枚しか存在し得ないため、このような未使用の桁は必ず少なくとも1つは存在する。

これを素直にシミュレーションすれば良い。

```go
var t int
var n int
var P [][]rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()

		P = [][]rune{}
		for i := 0; i < n; i++ {
			row := ReadRuneSlice()
			P = append(P, row)
		}

		solve()
	}
}

var used [10]int

func solve() {
	num := 0
	used = [10]int{}

	for i := 0; i < n; i++ {
		used[P[i][3]-'0'] = 1
	}

	for i := 0; i < n; i++ {
		A := P[i]
		// 直前のものすべてと比較する
		for j := 0; j < i; j++ {
			B := P[j]
			// 等しいものがあればAの最下位桁を未使用のものに変更
			if isEqual(A, B) {
				u := findMinUnused()
				A[3] = rune(u) + '0'
				num++
				used[u] = 1
				break
			} else {
				// 前のものと異なるなら何もしなくて良い
				continue
			}
		}
	}

	fmt.Println(num)
	for i := 0; i < n; i++ {
		fmt.Println(string(P[i]))
	}
}

func findMinUnused() int {
	for i := 0; i < 10; i++ {
		if used[i] == 0 {
			return i
		}
	}
	return -1
}

func isEqual(a, b []rune) bool {
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

制約をうまく使えず、ものすごく難しい考え方をして時間を浪費した挙げ句、
実装ミス（最下位桁で使用済みのものを最初にチェックしなかった）によりsystem test落ちという最悪ムーブをしてしまいました。

<a id="markdown-c-everyone-is-a-winner" name="c-everyone-is-a-winner"></a>
## C. Everyone is a Winner!

[問題のURL](https://codeforces.com/contest/1263/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

レート `n` をコンテストの参加者 `k` 人に振り分けるとき、1人あたりの受け取るレートは `Floor(n/k)` とする。

`k` を任意の正の整数とするとき、考えられる1人あたりのレートをすべて列挙せよ、という問題。

制約: `1 <= n <= 10^9`

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

素直に考えるのであれば、 `k = 1, 2, ..., n+1` として `Floor(n/k)` を計算すれば、
その列が答えとなる。

しかし、制約的にそれは出来ないため、効率的に行うことを考える。

厳密には異なるが、約数列挙に雰囲気が似ているため、 `O(sqrt(n))` ベースの手法に焦点を当てつつ考える。

なんとなく、 `k = 1, 2, ..., sqrt(n)` まで考えて、 `Floor(n/k) = q` だけでなく、 `k` 自体も答えになりそうだ、と仮説を立てる。
この仮説は、以下のようにして正当性が示せる（手書きですみません）。

<figure class="figure-image figure-image-fotolife" title="仮説の正当性の証明">[f:id:maguroguma:20191130185949j:plain]<figcaption>仮説の正当性の証明</figcaption></figure>

よって、 `k = 1, 2, ..., sqrt(n)` の `k` に対する `Floor(n/k) = q` は答えになるとともに、 `k` 自体も答えとなる。
一方で、 `k > sqrt(n)` の `k` については、 `Floor(n/k) < sqrt(n)` となり、これらはすでに答えとして追加されている。

よって、答えは `O(sqrt(n))` で列挙できる。
また、出力に際して、答えの要素数も `O(sqrt(n))` であるため、配列に詰めた後にソートすればよく、
これがネックとなり計算量は1テストケースあたり `O(sqrt(n)*log(n))` となり、間に合う。

```go
var t int
var n int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()

		solve()
	}
}

func solve() {
	memo := make(map[int]int)
	memo[0] = 1

	for k := 1; k*k <= n; k++ {
		memo[n/k] = 1
		memo[k] = 1
	}

	answers := []int{}
	for k := range memo {
		answers = append(answers, k)
	}
	sort.Sort(sort.IntSlice(answers))

	fmt.Println(len(answers))
	fmt.Println(PrintIntsLine(answers...))
}
```

例によって、コンテスト中はここまできっちり考えずに、ｴｲﾔしてしまいました。

<a id="markdown-d-secret-passwords" name="d-secret-passwords"></a>
## D. Secret Passwords

[問題のURL](https://codeforces.com/contest/1263/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

あるシステムのパスワードのリストには、 `n` 個の小文字のアルファベットのみからなるパスワードが書かれている。

このシステムのパスワードは、ある2つのパスワードを比較したとき、2つのパスワードに共通して存在する文字が1文字でもある場合、
この2つのパスワードは等しいと判定される。
さらに、推移律も存在し、パスワード `a, b` が等しく `b, c` が等しい場合、 `a, c` も等しいと判定される。

この中には、1つだけadminのパスワードが含まれているため、adminのパスワードと等しいものを手に入れるために、
最低限必要なパスワードの数を答えよ、という問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

個々のパスワードをグラフの頂点と考え、等しければエッジを張る、というふうに考える。

まず、すべてのパスワードをスキャンし、特定のアルファベットが検出されたら、アルファベットのグループに加える。
UnionFindによって、各アルファベットについてグループに存在するものをすべて併合する。
すべてのアルファベットについて併合が終わると、推移律を満たす形で等しいと判定されるパスワードも同じグループに属するようになる。
よって、最後に連結成分の数を出力すれば、それば答えになっている。

公式editorialにあるように、二部グラフ上で考えるのも良さそう
（ `a, b, .., z` を片方のグループ、パスワードをもう片方のグループとし、パスワードと文字を結んでいき、同じくDFS等で連結成分の数を数える。
こちらのほうが計算量は若干小さい）。

```go
var n int
var S [][]rune

var memo [30][]int

func main() {
	n = ReadInt()
	S = make([][]rune, n)
	for i := 0; i < n; i++ {
		S[i] = ReadRuneSlice()
	}

	for i := 0; i < n; i++ {
		for _, r := range S[i] {
			tmp := r - 'a'
			memo[tmp] = append(memo[tmp], i)
		}
	}

	uf := NewUnionFind(n)

	for i := 0; i < ALPHABET_NUM; i++ {
		if len(memo[i]) == 0 {
			continue
		}

		top := memo[i][0]
		for j := 1; j < len(memo[i]); j++ {
			uf.Unite(top, memo[i][j])
		}
	}

	fmt.Println(uf.CcNum())
}

// 0-based
// uf := NewUnionFind(n)
// uf.Root(x) 			// Get root node of the node x
// uf.Unite(x, y) 	// Unite node x and node y
// uf.Same(x, y) 		// Judge x and y are in the same connected component.
// uf.CcSize(x) 		// Get size of the connected component including node x
// uf.CcNum() 			// Get number of connected components

// UnionFind provides disjoint set algorithm.
// Node id starts from 0 (0-based setting).
type UnionFind struct {
	parents []int
}

// NewUnionFind returns a pointer of a new instance of UnionFind.
func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parents = make([]int, n)

	for i := 0; i < n; i++ {
		uf.parents[i] = -1
	}

	return uf
}

// Root method returns root node of an argument node.
// Root method is a recursive function.
func (uf *UnionFind) Root(x int) int {
	if uf.parents[x] < 0 {
		return x
	}

	// route compression
	uf.parents[x] = uf.Root(uf.parents[x])
	return uf.parents[x]
}

// Unite method merges a set including x and a set including y.
func (uf *UnionFind) Unite(x, y int) bool {
	xp := uf.Root(x)
	yp := uf.Root(y)

	if xp == yp {
		return false
	}

	// merge: xp -> yp
	// merge larger set to smaller set
	if uf.CcSize(xp) > uf.CcSize(yp) {
		xp, yp = yp, xp
	}
	// update set size
	uf.parents[yp] += uf.parents[xp]
	// finally, merge
	uf.parents[xp] = yp

	return true
}

// Same method returns whether x is in the set including y or not.
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// CcSize method returns the size of a set including an argument node.
func (uf *UnionFind) CcSize(x int) int {
	return -uf.parents[uf.Root(x)]
}

// CcNum method returns the number of connected components.
// Time complextity is O(n)
func (uf *UnionFind) CcNum() int {
	res := 0
	for i := 0; i < len(uf.parents); i++ {
		if uf.parents[i] < 0 {
			res++
		}
	}
	return res
}
```

後日復習したら簡単だった。。

確かコンテスト中は前半でペースを乱されたのも合って、問題文の英語が読めなくて解けなかったという感じでした
（そんなセキュリティunkなシステムあるかよ、的な感じで）。

問題文の本質的な部分を読み取ることに集中する訓練が、まだまだ足りないなと感じました。

---

時間が許すのなら、AtCoder始めたときのようにこどふぉも200問ぐらいまとめて解きたいです。
