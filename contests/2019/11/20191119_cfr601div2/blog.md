素の状態でB問題の嘘解法に疑問を持てず、もやもや。

<!-- TOC -->

- [A. Changing Volume](#a-changing-volume)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Fridge Lockers](#b-fridge-lockers)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
  - [嘘解法の反例](#%e5%98%98%e8%a7%a3%e6%b3%95%e3%81%ae%e5%8f%8d%e4%be%8b)
- [C. League of Leesins](#c-league-of-leesins)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-changing-volume" name="a-changing-volume"></a>
## A. Changing Volume

[問題URL](https://codeforces.com/contest/1255/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

テレビのボリューム `a` を `b` に変化させたい。

`-5, -2, -1, +1, +2, +5` の6つのボタンがあるので、できるだけ少ない回数でボリュームを合わせる場合、
最小で何回操作が必要か答える問題。

<a id="markdown-解答" name="解答"></a>
### 解答

まず、操作は「ボリュームを上げ続ける」か「ボリュームを下げ続ける」のいずれかで良い。
（`+1 -> -1` は意味がなく +`2 -> -1` は `+1` 一回で代替可能で、 `+5 -> -2, +5 -> -1` もそれぞれ `+1 -> +2, +2 -> +2` の操作で代替可能なため。）
この場合、できるだけ1回の操作で目的の `b` により近づけるボタンを選択すれば良い。

また、用意されているボタンは対称性があるため、 `diff = Abs(b - a)` という差の絶対値を考える。
そうすると、 `diff` に対して超過しないようにボタンの絶対値で引ける回数を調べればよく、
これはそれぞれの商とあまりを考えることで求められる。

```go
var t int
var a, b int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		a, b = ReadInt2()

		solve()
	}
}

func solve() {
	diff := AbsInt(b - a)
	ans := 0
	ans += diff / 5
	diff %= 5
	ans += diff / 2
	diff %= 2
	ans += diff

	fmt.Println(ans)
}
```

<a id="markdown-b-fridge-lockers" name="b-fridge-lockers"></a>
## B. Fridge Lockers

[問題URL](https://codeforces.com/contest/1255/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

`n` 人のが持つ `n` 個の冷蔵庫同士が鎖で繋がれており、さらに両端に鍵が取り付けられている。
自身の冷蔵庫に繋がれた鎖は自由に外すことができる（両端が外せるので繋がれた相手の冷蔵庫の鎖も外せる）。

すべてのメンバーについて、「自分が外せる鎖をすべて外したときに、他のすべての冷蔵庫に何らかの鎖が取り付けられている」
という状態が成り立つときに、privateであると呼ぶ。

各冷蔵庫にはコストが振られており、冷蔵庫 `i, j` 同士を鎖で結びつける場合、 `A[i] + A[j]` のコストがかかるとする。
同じ冷蔵庫のペアに、鎖を何本重ねてつなげても良い。
`m` 本の鎖が与えられたとき、privateな状態にするために必要な最小のコストを求め、privateに出来ない場合は `-1` を出力する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

ある冷蔵庫に着目したときに、その冷蔵庫につながっている鎖をすべて外して、なお他のすべての冷蔵庫に鎖がつながっていなければならないとすると、
すべての冷蔵庫は「少なくとも自分以外の異なる2つの冷蔵庫と鎖で繋がれていなければならない」と主張できる。
（もし、自分以外の1つの冷蔵庫としかつながっていない場合、その他方の冷蔵庫に鎖を外されたら、自分につながる鎖はなくなってしまう。）

このような状態を満たすためには、 `1 -> 2 -> ... -> n -> 1` とループさせれば自然と構築できる。
このとき必要になる鎖は `n` 本であるため、最低 `n` 本の鎖があればprivateな状態は作れる。

ただし、 `n = 2` の場合は、自分以外と繋げられる冷蔵庫の数は1個のみなので、この場合は鎖が何本あってもprivateには出来ない。

また、 `n < m` のときは、できるだけつなげるコストが小さい冷蔵庫同士を繰り返し鎖で結べばよく、
それは `A` についてソートして前2つを選べば良い(※)。

。。と思いきや、（※）部分は嘘解法。

今回はコンテスト中に `m > n` の制約が取っ払われたため、このようなケースはジャッジされなくなった。

```go
var t int
var n, m int
var A []int

type Edge struct {
	key          int
	nodeId, cost int
}
type EdgeList []*Edge
type byKey struct {
	EdgeList
}

func (l EdgeList) Len() int {
	return len(l)
}
func (l EdgeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.EdgeList[i].key < l.EdgeList[j].key
}

// how to use
// L := make(EdgeList, 0, 200000+5)
// L = append(L, &Edge{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n, m = ReadInt2()
		A = ReadIntSlice(n)

		solve()
	}
}

func solve() {
	if m < n {
		fmt.Println(-1)
		return
	}

	if n == 2 {
		fmt.Println(-1)
		return
	}

	sumA := Sum(A...)
	cost := 2 * sumA

	L := make(EdgeList, 0)
	for i := 0; i < n; i++ {
		L = append(L, &Edge{key: A[i], nodeId: i + 1, cost: A[i]})
	}
	sort.Stable(byKey{L})

	mm := m - n
	cost += mm * (L[0].cost + L[1].cost)

	answers := [][]int{}
	for i := 1; i <= n; i++ {
		if i == n {
			answers = append(answers, []int{i, 1})
		} else {
			answers = append(answers, []int{i, i + 1})
		}
	}
	for i := 0; i < mm; i++ { // この部分は嘘解法に含まれるが、コンテスト後半のテストケースでは実行されることはない
		answers = append(answers, []int{L[0].nodeId, L[1].nodeId})
	}

	fmt.Println(cost)
	for i := 0; i < m; i++ {
		fmt.Printf("%d %d\n", answers[i][0], answers[i][1])
	}
}
```

<a id="markdown-嘘解法の反例" name="嘘解法の反例"></a>
### 嘘解法の反例

[ここ](https://codeforces.com/blog/entry/71562)で `m > n` のケースが議論されている。
嘘解法をざっくりと否定するならば以下のような感じでしょうか。

> 各冷蔵庫について、自身がループに組み込まれさえすれば良い。
> 鎖が `n` 本しかない場合は1つのループですべてをつなげればよいが、
> たくさんある場合には、できるだけコストの小さい冷蔵庫とループが組めるようにするほうが良い。

[Um_nik氏による正しい証明はこちらから。](https://codeforces.com/blog/entry/71562?#comment-559266)）

たしかにこれは、なまじ慎重で賢明な人ほど考え込んで損をしている、というケースがかなりありそう。
激遅3完でも比較的マシなパフォーマンスが出たのはこの辺も影響していたのかも。。？

これに気づかずに突っ走ってしまった自分も大分問題な気がするけど、
Div2のBだからあまり考えなかったし、フルフィードバックならWAの後もう少し考え込むだろう、
ということにして、とりあえずは深入りしないようにします。

<a id="markdown-c-league-of-leesins" name="c-league-of-leesins"></a>
## C. League of Leesins

[問題URL](https://codeforces.com/contest/1255/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

`1, 2, ..., n`  で構成される順列 `P` が与えられる。

この数列を前から順番に連続するtripleに分割し、 `n-2` 個のtripleからなる配列を生成する。
このtripleの配列中でtripleの順番をデタラメに入れ替え、
さらに、個々のtripleの中で数字の順番を入れ替える。

今、 `n-2` 個のtripleが与えられるので、元の順列 `P` を復元せよ、という問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

とりあえず、元の順列からtripleが切り分けられる様子を図示すると、
ある数字の出現頻度が大きなヒントになりそうなことがすぐに分かる。



ほとんど（ `n-4` 個）の数字の出現頻度は3回となり、1回、2回がそれぞれ2つずつ、という分布になる。
例えば、出現頻度が1のある数字を考えたとき、その数字を含むtripleは唯一で、
さらにその中に出現頻度が2回である数字の片方が含まれる。

このtripleを左から並べ始めるとすると、次に並べるべきは、1つ目のtripleに含まれる頻度2回の数字および頻度1回の数字を含む、
別のtripleであると決定できる。
使った数字の頻度を減らして考えると、以降も同様に考えられるため、左から順番に元の数列を復元していくことができる。

（難易度はぜんぜん違うけど、[JSCのC問題](https://atcoder.jp/contests/jsc2019-qual/tasks/jsc2019_qual_c)のように、
一見複雑だけど実は左から順番に再帰的に決まっていく、というのはよく見る。）

続いて具体的な実装方法について考える。

上述の考察を踏まえると、1つ前に選んだtripleが左から `(a, b, c)` とすると、
次に選択すべきtripleは `(b, c, x)` のような形、すなわち「 `b, c` を含み、かつ `a` を含まないtriple」となる。

目的のtripleを探すのにすべてのtripleを探すわけには行かないので、一工夫が必要。
とはいえ、それほど面倒なことは考えずに、「ある数字 `b` を含む高々3つのtriple」の中からすべて調べれば良い。
これは、バケット法の要領でtripleを管理しておけば、ある数字 `b` のバケットの中を全探索する形で探索できる。



肝心の数列の復元部分については、目的のtripleを見つけるたびに `x` に該当する数字を都度appendしていくようにすればよい。

ここまで整理してもあまりいい実装にはならなかったので、
適宜関数を切り分けてバグらせないように注意する。

```go
var n int
var T [][]int

var book [100000 + 5][][]int // バケット

func main() {
	n = ReadInt()
	for i := 0; i < n-2; i++ {
		T = append(T, ReadIntSlice(3))
	}

	// 頻度をカウント
	memo := make([]int, n+1)
	for i := 0; i < n-2; i++ {
		for j := 0; j < 3; j++ {
			q := T[i][j]
			memo[q]++
		}
	}

	ones := []int{}
	twos := []int{}
	for i := 1; i < n+1; i++ {
		if memo[i] == 1 {
			ones = append(ones, i)
		} else if memo[i] == 2 {
			twos = append(twos, i)
		}
	}

	// バケットでtripleを管理
	for i := 0; i < n-2; i++ {
		a, b, c := T[i][0], T[i][1], T[i][2]
		book[a] = append(book[a], []int{a, b, c})
		book[b] = append(book[b], []int{a, b, c})
		book[c] = append(book[c], []int{a, b, c})
	}

	// 先頭のトリプルを見つけ、元の配列の3番目までを復元する
	targetInt := ones[0]
	firstTriple := book[targetInt][0]
	answers := make([]int, 3)
	for i := 0; i < 3; i++ {
		tmp := firstTriple[i]
		if memo[tmp] == 1 {
			answers[0] = tmp
		} else if memo[tmp] == 2 {
			answers[1] = tmp
		} else {
			answers[2] = tmp
		}
	}

	// 完全に復元されるまでループ
	for len(answers) < n {
		l := len(answers)
		ex, inc1, inc2 := answers[l-3], answers[l-2], answers[l-1]
		// 高々3個のtripleから該当するものを見つける
		for _, tri := range book[inc2] {
			if sub(tri, ex, inc1, inc2) {
				// inc1, inc2以外の値をanswersに追加する
				answers = append(answers, subsubsub(tri, inc1, inc2))
				break
			}
		}
	}

	fmt.Println(PrintIntsLine(answers...))
}

// triが目的のものだったらtrue
func sub(tri []int, ex, inc1, inc2 int) bool {
	// exを含んだらfalse
	if subsub(tri, ex) {
		return false
	}

	// inc1を含まないならfalse
	if !subsub(tri, inc1) {
		return false
	}
	// inc2を含まないならfalse
	if !subsub(tri, inc2) {
		return false
	}

	return true
}

// triがtargetを含んだらtrue
func subsub(tri []int, target int) bool {
	for _, v := range tri {
		if v == target {
			return true
		}
	}
	return false
}

// inc1, inc2以外の値を返す
func subsubsub(tri []int, inc1, inc2 int) int {
	for _, v := range tri {
		if v != inc1 && v != inc2 {
			return v
		}
	}
	return -1
}
```

所望のtripleの検索方法の部分で何故か難しく考えてしまい、時間を浪費してしまったのが大反省。
しかしながら、競技でこういった実装・発想をした経験がなかったので、1つ1つ引き出しを増やしていくしか無いのかも。

---

未だにグローバルに変数を置いたり、命名を適当にやってしまうことに抵抗があります。
