こどふぉの算数が苦手とかそういうレベルじゃなく。

<!-- TOC -->

- [A. Two Rival Students](#a-two-rival-students)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Magic Stick](#b-magic-stick)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Dominated Subarray](#c-dominated-subarray)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Yet Another Monster Killing Problem](#d-yet-another-monster-killing-problem)
  - [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
  - [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-two-rival-students" name="a-two-rival-students"></a>
## A. Two Rival Students

[問題URL](https://codeforces.com/contest/1257/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

`n` 人の横一列に並んだ生徒の中に2人のライバルが居るので、それらの学生をできるだけ互いに引き離したい。
一回の操作で隣り合う2人の生徒の位置を入れ替えることができる、とする場合に、
最大 `x` 回の操作でどれだけ引き離すことができるか求める問題。

<a id="markdown-解答" name="解答"></a>
### 解答

`n` も `x` もテストケースの数も高々100なので、シミュレーションが十分間に合う。
頑張れば `O(1)` でも求まると思うが、バグらせたくないのでシミュレーションを書いた。

```go
var t int
var n, x, a, b int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n, x, a, b = ReadInt4()
		solve()
	}
}

func solve() {
	if a > b {
		a, b = b, a
	}

	for {
		if x == 0 || a == 1 {
			break
		}

		a--
		x--
	}

	for {
		if x == 0 || b == n {
			break
		}

		b++
		x--
	}

	fmt.Println(AbsInt(b - a))
}
```

<a id="markdown-b-magic-stick" name="b-magic-stick"></a>
## B. Magic Stick

[問題URL](https://codeforces.com/contest/1257/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

ある正の整数について、2種類の魔法をかけることができる。
それぞれの魔法の結果、正の整数は以下のように変化する。

1. `a` が偶数ならば、 `a -> a/2*3` とできる。
2. `a` が1より大きいならば、 `a -> a-1` とできる。

ある2つの正の整数 `x, y` が与えられるので、魔法を好きな順番で何回でも使っても良いので、
`x` から `y` を生み出せるか、を判定する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

ゴールは `x` を `y` よりも大きくすること、というのはすぐに分かる。

よくよく考えると（残念ながら自分はよくよく考えないと気づけなかった）、
`a` が4以上であるならば、2種類の魔法を適当に繰り返し用いれば（偶数なら `1` 、奇数なら `2` ）、
いくらでも数を大きくできることに気づく。

特殊なのは `1, 2, 3` の3つだけで、 `1` は魔法をかけることができず、 `2, 3` はそれぞれループしてしまう。

このことに注意して場合分けを行えば良い。

```go
var t int
var x, y int64

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		x, y = ReadInt64_2()
		solve()
	}
}

func solve() {
	if x >= y {
		fmt.Println("YES")
		return
	}

	if x == 1 {
		if x >= y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		return
	}

	if x == 2 || x == 3 {
		if y <= 3 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		return
	}

	fmt.Println("YES")
}
```

問題を考えているうちに、勘違いして別の問題を解き始めてしまい、死ぬほど時間を取られてしまいました。

<a id="markdown-c-dominated-subarray" name="c-dominated-subarray"></a>
## C. Dominated Subarray

[問題URL](https://codeforces.com/contest/1257/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

2要素以上の長さをもつ配列に対して、その配列の中である1つの整数の頻度が狭義で1位（同率1位ではない）となる場合、
その配列を dominated subarray と呼ぶ。

ある配列 `A` が与えられるので、最小の長さを持つ dominated subarray の配列長を答えよ。
また、 dominated subarray が存在しない場合は `-1` を出力せよ。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

制約的に `O(nlogn)` が間に合うと考え、長さを決め打ちし、その長さ以下の dominated subarray が存在するかどうかを、二分探索で探索することを考えた。
ある長さについて、その長さ（以下）の dominated subarray が存在するかどうかの判定に関しては `O(n)` かかってもよい（配列の全探索のようなことが許容される）。

ある長さ `m` を決めたとき、配列 `A` に対して長さ `m` の固定長のスライディングウィンドウを考える。
このスライディングウィンドウを subarray と考え、この subarray 中の各整数値の頻度をカウントする。
すると、いずれかのカウントが2以上となった時点で、長さ `m` 以下の dominated subarray が存在することを主張できる。
なぜなら、カウントが2回の整数を `a` としたとき、 `[a, ..., a]` のように両端が `a` のような subarray が dominated subarray となるためである。

```go
var t int
var n int
var A []int
var cnt []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A = ReadIntSlice(n)
		solve()
	}
}

func solve() {
	if n == 1 {
		fmt.Println(-1)
		return
	}

	cnt = make([]int, n+1)
	flag := true
	for i := 0; i < n; i++ {
		a := A[i]
		if cnt[a] > 0 {
			// 2回登場するものがあれば答えはある
			flag = false
			break
		}
		cnt[a]++
	}

	if flag {
		// すべて異なる数ならNO
		fmt.Println(-1)
		return
	}

	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		// リセット
		cnt = make([]int, n+1)
		// 初期化
		for i := 0; i < m; i++ {
			a := A[i]
			if cnt[a] > 0 {
				return true
			}
			cnt[a]++
		}
		// スライド
		for i := m; i < n; i++ {
			j := i - m
			cnt[A[j]]--
			a := A[i]
			if cnt[a] > 0 {
				return true
			}
			cnt[a]++
		}

		return false
	}

	ng, ok := 1, n
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
}
```

dominated subarray の性質にもっと早い段階で気づけたら、こんな面倒なことやらずに、もっと楽な方法が取れたと思う。

<a id="markdown-d-yet-another-monster-killing-problem" name="d-yet-another-monster-killing-problem"></a>
## D. Yet Another Monster Killing Problem

[問題URL](https://codeforces.com/contest/1257/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

あるゲームにおいて、 `m` 体のヒーローが、ある1つのダンジョンを攻略しようとしている。
各ヒーローは、力とスタミナの2つのパラメータを有している。

また、ダンジョンには `n` 体のモンスターがおり、各モンスターは力のパラメータのみを有している。
モンスターは決まった順番に並んでおり、ヒーロー達はこの順番にモンスターと出会い、討伐していく。

このゲームは一日単位のターン制で、一日にある1体のヒーローのみがダンジョンに潜入できる。
ダンジョンではモンスターと一体ずつ戦い、ヒーローの力がモンスターの力以上だった場合、
ヒーローはそのモンスターを討伐できる。
ただし、1日に連続して討伐できるモンスターは、潜入したヒーローのスタミナの値までである。
また、ヒーローの力が出会ったモンスターの力未満だった場合、ヒーローは帰還し一日は終了する。

このような設定のもとで、ダンジョンを攻略するのに要する最短の日数はいくらか。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

素直に考えて、生き残っているモンスターは先頭からもれなく倒していく必要があるため、
その日のダンジョンの状況において、できるだけたくさんのモンスターを倒せるようなヒーローを、都度選択するのが最善となる。

一日ごとにすべてのヒーローを全探索して、最もモンスターを多く倒せるヒーローを選択できればよいが、
最悪のケースでは、1日にモンスター1体しか倒すことができず、このようなヒーロー列の全探索をモンスターの数だけ行う事になってしまう。
計算量は `O(n*m)` となるため、今回の制約では間に合わない。

そこで、モンスターの列に対して起点を設定し、さらに1つずつ目標ラインを上げていき、その区間のモンスターをすべて倒せるヒーローが居るかどうかを、
`O(logn)` で求められないか？と考えてみる。

この区間のモンスターを倒す条件は、

1. ヒーローのスタミナが区間の長さ以上であること。
2. ヒーローの力が、区間内のモンスターの力の最大値以上であること。

1については、ヒーローを予めスタミナでソートしておけば、条件を満たすヒーローたちは二分探索で高速に検索できる。
2については、選択対象のヒーローは前述の二分探索で求めたヒーローよりスタミナが大きいヒーロー群なので、
ヒーロー列の後半部分に対して、予め力の「累積Max」とでも言うべきものを前計算しておけば、 `O(1)` で取得できる。
また、区間内のモンスターの力の最大値については、目標ラインを上げる際に最大値を更新することで、 `O(1)` で取得できる。

以上より、必要なパーツは揃ったので、手順に従って実装していく。
（個人的には）目標ラインを上げたり、1日後のモンスター列の起点を更新する部分がバグりやすいと感じたので、
適宜別の関数として切り分けるなど、工夫をする。

計算量はヒーロー列のソートと最大 `n` 回ヒーローの選択が行われることから `O((m+n)*log(m))` 。

```go
type Hero struct {
	key  int
	p, s int
}
type HeroList []*Hero
type byKey struct {
	HeroList
}

func (l HeroList) Len() int {
	return len(l)
}
func (l HeroList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.HeroList[i].key < l.HeroList[j].key
}

// how to use
// L := make(HeroList, 0, 200000+5)
// L = append(L, &Hero{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

var t int
var n, m int
var A, P, S []int
var L HeroList // ヒーロー構造体の配列
var M []int    // [idx, m-1] のヒーロー区間におけるpowerの最大値を記憶

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A = ReadIntSlice(n)
		m = ReadInt()
		P, S = make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			p, s := ReadInt2()
			P[i], S[i] = p, s
		}

		solve()
	}
}

func solve() {
	L = make(HeroList, 0)
	for i := 0; i < m; i++ {
		p, s := P[i], S[i]
		L = append(L, &Hero{key: s, p: p, s: s})
	}
	sort.Stable(byKey{L}) // endurance順に昇順ソートする

	// 後半部分のpowerに関する累積Maxを計算しておく
	M = make([]int, m)
	M[m-1] = L[m-1].p
	for i := m - 2; i >= 0; i-- {
		M[i] = Max(M[i+1], L[i].p)
	}

	l := 0
	ans := 0
	for l < n {
		length := sub(l)
		if length == -1 {
			fmt.Println(-1)
			return
		} else {
			l += length
			ans++
		}
	}

	fmt.Println(ans)
}

// モンスターの配列に対して、A[s]から数えて何体倒せるかを計算する関数
// -1を返したら1体も倒せない（=失敗）
func sub(s int) int {
	maxMP := A[s]
	length := -1
	for l := 0; s+l < n; l++ {
		maxMP = Max(maxMP, A[s+l]) // A[s]から現在見ているところまでのモンスターのpowerの最大値
		idx := sub2(l + 1)
		if idx == m {
			break
		}

		maxHP := M[idx] // l+1以上のenduranceを持つヒーローの中のpowerの最大値
		if maxHP >= maxMP {
			length = l + 1
		} else {
			break
		}
	}

	return length
}

// enduranceがl以上となるギリギリのインデックスを二分探索で計算
// mを返したらそのようなヒーローは存在しない
func sub2(l int) int {
	// m は中央を意味する何らかの値
	isOK := func(mid int) bool {
		if L[mid].s >= l {
			return true
		}
		return false
	}

	ng, ok := -1, m
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}
```

コンテスト中は日数について二分探索する、という方法がまっさきに思い浮かび、
時間がなかったことも合って固執してしまったのが失敗でした。
まずは基本に従って「素直に愚直に考えてから計算量を落とす」という思考もちゃんと視野に入れないとダメですね。

また、コンテスト後に「セグ木を使った」とか「セグ木2本使って判定した」とかのツイートが見られましたが、
活用方法がちょっとわかりませんでした。
コードが劇的に書きやすくなるとかだったら活用したいところですが、ライブラリ整理もしっかりできていないと難しそう。

---

Dみたいな問題の安定感を高めていきたいところ。
