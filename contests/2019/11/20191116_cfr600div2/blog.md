<!-- TOC -->

- [A. Single Push](#a-single-push)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Silly Mistake](#b-silly-mistake)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Sweets Eating](#c-sweets-eating)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)
  - [公式editorialの解法](#%e5%85%ac%e5%bc%8feditorial%e3%81%ae%e8%a7%a3%e6%b3%95)
- [D. Harmonious Graph](#d-harmonious-graph)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
  - [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-single-push" name="a-single-push"></a>
## A. Single Push

[問題のURL](https://codeforces.com/contest/1253/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

与えられた配列 `A` に対して、1度だけ任意の連続区間に対してある正の整数加算することが許される。
操作は行わなくても良い。

これによって、もう一方の与えられた配列 `B` に等しくすることができるか判定する問題。

<a id="markdown-解答" name="解答"></a>
### 解答

すべての要素に関して `diff[i] = B[i] - A[i]` を計算しておく。
この `diff` 配列が `[0, ..., 0, k, ..., k, 0, ..., 0], k >= 0` のようになっていればよい。

判定を簡単にするために、 `diff` 配列に対して[ランレングス圧縮](http://e-words.jp/w/%E3%83%A9%E3%83%B3%E3%83%AC%E3%83%B3%E3%82%B0%E3%82%B9%E5%9C%A7%E7%B8%AE.html)を施す。
圧縮後の配列に対して、

1. 負の整数が検出されたらアウト
2. 正の整数が2つ以上検出されたらアウト
3. そうでないならセーフ

のように判定すればよい。

```go
var t int
var n int
var A, B []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A, B = ReadIntSlice(n), ReadIntSlice(n)

		solve()
	}
}

func solve() {
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = B[i] - A[i]
	}

	pressed, _ := RunLengthEncoding(diff)

	positive := 0
	for i := 0; i < len(pressed); i++ {
		if pressed[i] < 0 {
			fmt.Println("NO")
			return
		}

		if pressed[i] > 0 {
			positive++
		}
	}

	if positive == 0 || positive == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []int) ([]int, []int) {
	runes := []int{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		// 1文字目の場合保持
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			// 直前の文字と一致していればインクリメント
			l++
		} else {
			// 不一致のタイミングで追加し、長さをリセットする
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []int, L []int) []int {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}
```

<a id="markdown-b-silly-mistake" name="b-silly-mistake"></a>
## B. Silly Mistake

[問題のURL](https://codeforces.com/contest/1253/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

あるオフィスの入退場記録が配列で表されている。
正の整数が入場、負の整数が退場を表している。

この記録には満たすべき性質があり、以下の3つがある。

1. 各メンバは入場の前に退場することはない。
2. 各メンバは1日に複数回入退場してはならない。
3. 各メンバは入場したら必ずその日のうちに退場しなければならない。

このような性質を満たすように、入退場記録を1日ごとに正しく分割せよ、という問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

特に日数を小さくしたり大きくしたり、という制約はないので、
「オフィスが空になったら、すぐさま一日をリセットする」というふうにシミュレーションするのが良い。
（リセットしたほうが、あるメンバの複数回入場のチェックが楽になるので、このほうが簡単。）

各入退場イベントを処理・管理するにあたり、オフィスに現在いる人、およびその日の各メンバの入退場記録を `map` で管理する。
（メンバは固定で `10^6` なので、固定長配列だと空室判定や一日のリセットに時間がかかり間に合わない。）

処理する中で問題の3つの制約に抵触したらその時点で `-1` を出力すれば良い。

```go
var n int
var A []int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)

	if n%2 == 1 {
		fmt.Println(-1)
		return
	}

	memo := make(map[int]int)  // オフィス
	times := make(map[int]int) // ある一日の入場回数
	answers := []int{}
	count := 0
	for i := 0; i < n; i++ {
		a := A[i]

		if a < 0 {
			a = -a
			// 退出
			if _, ok := memo[a]; ok {
				delete(memo, a)
			} else {
				// 入場前退出のためアウト
				fmt.Println(-1)
				return
			}
		} else {
			// 入場
			if times[a] == 0 {
				memo[a] = 1
				times[a] = 1
			} else {
				// 一日に2回登場したのでアウト
				fmt.Println(-1)
				return
			}
		}

		count++
		// 空室判定
		if len(memo) == 0 {
			// 空室なので次の日へリセット
			answers = append(answers, count)
			count = 0
			memo = make(map[int]int)
			times = make(map[int]int)
		}
	}

	if len(memo) != 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(len(answers))
	fmt.Println(PrintIntsLine(answers...))
}
```

一番最後の空室判定を忘れて1WA出してしまったのが大反省。

競技プログラミングでGolangの `map` の `delete` を行ったのは何気に初めてな気がする。

<a id="markdown-c-sweets-eating" name="c-sweets-eating"></a>
## C. Sweets Eating

[問題のURL](https://codeforces.com/contest/1253/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

けいおん

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

食べるケーキ `k` 個は、ケーキ列を昇順ソートした上でその前 `k` 個でよい、というのはすぐに分かる。
また、食べる順番についても、カロリーが大きい方から小さい方を選ぶ形でよい、というのもすぐに分かる。

これを愚直に各 `k` について線形スキャンする形で行うと、トータルの計算量が `O(n^2)` になって間に合わない。

そこで、直前の結果を利用して高速に計算できないかを考える。
図のように（ `m = 2` のケース）、 `k` が1増えるごとに、追加して食べるケーキの番号から `m` 飛びのケーキについても砂糖を加算する必要があるとわかる。



この `m` 飛びの累積和を事前に `O(n)` で計算しておけば、 `k` のときの答えを利用して `k+1` が計算できる。

結局、ソートがネックになるため、計算量は `O(nlogn)` 。

```go
var n, m int
var A []int
var answers []int64

func main() {
	n, m = ReadInt2()
	A = ReadIntSlice(n)
	answers = make([]int64, n)

	sort.Sort(sort.IntSlice(A))

	memo := make([]int64, n)
	for i := 0; i < m; i++ {
		memo[i] = int64(A[i])
	}
	for i := m; i < n; i++ {
		memo[i] = memo[i-m] + int64(A[i])
	}

	answers[0] = int64(A[0])
	for i := 1; i < n; i++ {
		answers[i] = answers[i-1] + memo[i]
	}
	fmt.Println(PrintIntsLine(answers...))
}
```

`m` 飛びの累積和の計算というのを初めてやったので、最初やり方が分からずに時間を食ってしまった。
ちょっと筋の悪い手法だった気がする。

<a id="markdown-公式editorialの解法" name="公式editorialの解法"></a>
### 公式editorialの解法

直前の結果ではなく、 `m` 個前の結果を利用しましょうという方法。

`m` 個前の結果を基準として考えると、新たに追加して食べるケーキを含めて、
それまでのケーキすべての累積和をそのまま加算することで、
`k` 個の場合の答えが得られる。



```go
var n, m int
var A []int
var answers []int64

func main() {
	n, m = ReadInt2()
	A = ReadIntSlice(n)
	answers = make([]int64, n)

	sort.Sort(sort.IntSlice(A))

	sums := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + int64(A[i])
	}

	for i := 0; i < n; i++ {
		if i < m {
			answers[i] = sums[i+1]
			continue
		}

		answers[i] = answers[i-m] + sums[i+1]
	}

	for i := 0; i < n; i++ {
		if i == n {
			fmt.Printf("%d\n", answers[i])
		} else {
			fmt.Printf("%d ", answers[i])
		}
	}
}
```

すでに解いた部分問題を可能な限り利用とするのはDP考える上でも重要だと思うので、
こういった視点が他の問題でも持てるようになりたい。

<a id="markdown-d-harmonious-graph" name="d-harmonious-graph"></a>
## D. Harmonious Graph

[問題のURL](https://codeforces.com/contest/1253/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

`n` 頂点 `m` 辺からなる無向グラフが与えられる。

また、すべての `(l, m, r), 1 <= l < m < r <= n` について、 `l, r` 間にパスが存在するときには `l, m` にもパスが存在する場合、
そのグラフは harmonious であるという。

与えられた無向グラフに対して辺を追加して harmonious にするためには、最小で何本の辺を足す必要があるか求めよ、という問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

与えられたグラフに対し、DFSなりUnion Find木を使うなりして連結成分を計算する。
さらに、各連結成分を構成するノードのIDについて、最小のものと最大のもの（それぞれ `l, r` とする）を調べておく。

すると、 harmonious である状態とは、各連結成分を `[l, r]` の区間とみなすと、
区間の交差がない状態であると言える。
よって、交差している区間同士をマージすればよく、そのためには連結成分を構成する適当な2点間に1本辺を足すだけで良い。

結局、区間をマージするたびに答えをインクリメントする、という処理を高速に行えば良い。
このためには、 `[l, r]` を `l` 基準で昇順ソート（※ノード番号の小さい順にDFSすれば、自然とソートされる）し、その順に区間をスキャンしていけば良い。
具体的には、暫定の `r` の最大値を保持しながら、次の `l` がその最大値以下であれば交差していると判定できる。

```go
var n, m int
var G [200000 + 5][]int
var colors [200000 + 5]int
var left, right int

func main() {
	n, m = ReadInt2()

	for i := 0; i < m; i++ {
		x, y := ReadInt2()
		x--
		y--
		G[x] = append(G[x], y)
		G[y] = append(G[y], x)
	}

	L := make(ComponentList, 0)
	for i := 0; i < n; i++ {
		colors[i] = -1
	}
	for i := 0; i < n; i++ {
		if colors[i] == -1 {
			left, right = i, i
			dfs(i, i)
			L = append(L, &Component{key: left, l: left, r: right})
		}
	}

	ans := 0
	biggest := L[0].r
	for i := 1; i < len(L); i++ {
		if L[i].l <= biggest {
			ans++
			ChMax(&biggest, L[i].r)
		} else {
			biggest = L[i].r
		}
	}

	fmt.Println(ans)
}

func dfs(i, c int) {
	colors[i] = c

	for _, nid := range G[i] {
		if colors[nid] == -1 {
			ChMin(&left, nid)
			ChMax(&right, nid)
			dfs(nid, c)
		}
	}
}

type Component struct {
	key  int
	l, r int
}
type ComponentList []*Component
```

「交差する区間をマージする」という読み替えが出来なかったので反省。

加えて、区間を効率的にマージしていく方法も、簡単だけど意外と自分で考えるのは難しかったので、
ちゃんと覚えておきたいところです。

実装については、Union Find木を使うとちょっとだけ遅くなるけど、
人によってはそちらのほうが簡単に実装できるかも（自分は両方試したところDFSのほうがスッキリしました）。

---

こどふぉをやっていると、特に有名な名前はついていないけど時々必要になる実装手法（？）のようなものがよく出る気がするので、
このあたりもちゃんと定着させていきたいところ。
