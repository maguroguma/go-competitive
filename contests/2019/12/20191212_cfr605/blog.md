<!-- Codeforces Round No.605 (Div.3) 参加記録 (A〜E解答) -->

<!-- TOC -->

- [A. Three Friends](#a-three-friends)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Snow Walking Robot](#b-snow-walking-robot)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Yet Another Broken Keyboard](#c-yet-another-broken-keyboard)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Remove One Element](#d-remove-one-element)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)
	- [DPによる別解](#dp%e3%81%ab%e3%82%88%e3%82%8b%e5%88%a5%e8%a7%a3)
- [E. Nearest Opposite Parity](#e-nearest-opposite-parity)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-4)
	- [解答](#%e8%a7%a3%e7%ad%94-4)

<!-- /TOC -->

<a id="markdown-a-three-friends" name="a-three-friends"></a>
## A. Three Friends

[問題のURL](https://codeforces.com/contest/1272/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

数直線上に3人がそれぞれ座標 `a, b, c` にいる。
それぞれは、自身の座標を `+1, -1, 0` 変位させることができる。

ここで、総合距離を `|a-b| + |a-c| + |b-c|` とする。

ありうる総合距離の最小値を求めよ、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

まず、一般性を失わずに `a <= b <= c` となるようにする。

この状態では、総合距離は `(b-a) + (c-b) + (c-a) = 2*(c-a)` と表すことができる。
よって、移動後もこの関係が保たれているのであれば `b` はないものとして考えることができる。

よって、 `a, c` の位置関係に注意しながら、それぞれを右に動かすか、左に動かすか判断すれば良い。

オーバーフローには注意する。

```go
var q int
var a, b, c int64

func main() {
	q = ReadInt()

	for tc := 0; tc < q; tc++ {
		a, b, c = ReadInt64_3()

		solve()
	}
}

func solve() {
	A := LongInt{a, b, c}
	sort.Sort(A)

	aa, _, cc := A[0], A[1], A[2]
	if aa < cc {
		aa++
	}
	if aa < cc {
		cc--
	}
	ans := 2 * (cc - aa)

	fmt.Println(ans)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

type LongInt []int64

func (a LongInt) Len() int           { return len(a) }
func (a LongInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a LongInt) Less(i, j int) bool { return a[i] < a[j] }
```

。。こんな考え込まなくても、素直に全探索するべきですね（27通りしか無いので）。

通ってみればなんてことはないですが、こういった算数はpretestではびびりながら出しているので、
提出に14分かかっています。

今回のコンテストで一番反省すべき点だと思います。

<a id="markdown-b-snow-walking-robot" name="b-snow-walking-robot"></a>
## B. Snow Walking Robot

[問題のURL](https://codeforces.com/contest/1272/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

`L, R, U, D` の命令を与えると、それぞれ左、右、上、下に1マス進むロボットがある。

ただし、このロボットは一度通ったマスを通ると壊れてしまう。

ロボットを壊さずに、自分の家から出発させて、命令の最後には自分の家に戻ってくるようにしたい。

今、命令列が与えられるので、この列のいくつかの命令を消去することで、上記の操作を成立させたい。

最大で何個の命令を残すことができるか、またその時の命令列を出力せよ、という問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

基本的には、上下と左右に分けて考えて、例えば `UU..URR..RDD..DLL..L` のように
長方形を描くように移動させることを考えればよい。

このように移動させるためには、上への移動回数と下への移動回数が釣り合う必要があるため、
それぞれのカウントの最小値を採用すれば良い。
左右に関しても一緒である。

ただし、上下の移動ができない場合、あるいは左右の移動ができない場合は、
長方形ではなく線分となってしまい、同じ道を通ってしまうことをケアする必要がある。

この場合はコーナーケースとして別に扱う。

```go
var q int
var S []rune

func main() {
	q = ReadInt()

	for tc := 0; tc < q; tc++ {
		S = ReadRuneSlice()

		solve()
	}
}

func solve() {
	memo := make(map[rune]int)
	memo['L'], memo['R'], memo['U'], memo['D'] = 0, 0, 0, 0

	for _, r := range S {
		memo[r]++
	}

	lrTime := Min(memo['R'], memo['L'])
	udTime := Min(memo['U'], memo['D'])

	if lrTime == 0 && udTime == 0 {
		fmt.Println(0)
		fmt.Println("")
		return
	}

	if lrTime == 0 {
		// UDを1回ずつ
		fmt.Println(2)
		fmt.Println("UD")
		return
	}

	if udTime == 0 {
		// LRを1回ずつ
		fmt.Println(2)
		fmt.Println("LR")
		return
	}

	fmt.Println(lrTime*2 + udTime*2)
	answers := []rune{}
	for i := 0; i < lrTime; i++ {
		answers = append(answers, 'R')
	}
	for i := 0; i < udTime; i++ {
		answers = append(answers, 'D')
	}
	for i := 0; i < lrTime; i++ {
		answers = append(answers, 'L')
	}
	for i := 0; i < udTime; i++ {
		answers = append(answers, 'U')
	}
	fmt.Println(string(answers))
}
```

<a id="markdown-c-yet-another-broken-keyboard" name="c-yet-another-broken-keyboard"></a>
## C. Yet Another Broken Keyboard

[問題のURL](https://codeforces.com/contest/1272/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

ある長さが `n` の文字列が与えられたため、タイピングの練習のため、この文字列のすべての連続部分文字列をタイプしようとした。

しかし、いくつかのキーが壊れているため、いくつかの連続部分文字列は不完全な状態で出力されてしまった。

`n*(n+1)/2` 個の全連続部分文字列のうち、正しくタイピングできたものはいくつか求めよ、という問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

まず、与えられた文字列に対して、壊れているキーのインデックスに目印をつける。

左から文字列をスキャンしたとき、今見ている文字を右端とする部分文字列でタイピングできるものの数は、
以下の図に示すように、タイプ出来ない文字から数えた文字の長さ分だけである。



タイプ不可の文字に出会うたびにカウンターをリセットすることで、これらの数値はすべて記録できるため、
最後に総和を計算すれば良い。

オーバーフローには注意する。

```go
var n, k int
var S []rune
var Avail []rune

func main() {
	n, k = ReadInt2()
	S = ReadRuneSlice()
	for i := 0; i < k; i++ {
		tmp := ReadRuneSlice()
		Avail = append(Avail, tmp[0])
	}

	memo := make([]int, n)
	for i := 0; i < n; i++ {
		if !isAvail(S[i]) {
			memo[i] = -1
		}
	}

	cur := int64(0)
	sums := make([]int64, n)
	for i := 0; i < n; i++ {
		if memo[i] == -1 {
			// アウトなのでリセット
			cur = 0
		} else {
			// 継続
			cur++
		}
		sums[i] = cur
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		ans += sums[i]
	}
	fmt.Println(ans)
}

func isAvail(r rune) bool {
	for i := 0; i < len(Avail); i++ {
		if Avail[i] == r {
			return true
		}
	}
	return false
}
```

。。こんなにめんどくさいことしなくても、タイプ不可能のものとそうでないものを2値化して、
ランレングス圧縮したものを使ったほうが簡単でした。

そこそこ時間をロスしてしまったので、これも結構な失敗でした。

<a id="markdown-d-remove-one-element" name="d-remove-one-element"></a>
## D. Remove One Element

[問題のURL](https://codeforces.com/contest/1272/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

長さ `n` の整数配列 `A` が与えられる。

この整数配列の好きな要素を1個削除するか、もしくはそのままでも良い。

この状態の配列に対し、狭義単調増加である連続部分列を考える。

ありうる最も長い連続部分列の長さを求めよ、という問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

配列 `A` の要素に関して、消す要素を全探索することを考える。
すなわち、ある要素を削除したとき、影響を受けるのはその前後だけであり、
その際に作られる連続部分列がより大きいものとなるのなら更新していく、ということを考える。

まずは、もとの配列に対して、その要素が属する連続部分列の長さを計算する。

更に、数列間の階差を計算しておく。

以下の図に示すように、ある要素を削除したとき、新たに比較すべき階差は元の階差の和で表される。



よって、階差の和が正となる場合にのみ、削除される要素の前後の連続部分列の長さを足し合わせる形で、新たに出来上がる連続部分列の長さを計算する。

```go
var n int
var A []int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)

	// 階差
	diff := []int{} // len(diff) == n-1
	for i := 1; i < n; i++ {
		diff = append(diff, A[i]-A[i-1])
	}

	// 自身が含まれる増加列の長さ
	L := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			L[i] = 1
			continue
		}

		if A[i] > A[i-1] {
			L[i] = L[i-1] + 1
		} else {
			L[i] = 1
		}
	}
	cur, l := 0, 0
	for i := n - 1; i >= 0; i-- {
		if cur == 0 {
			cur, l = L[i], L[i]
		}
		L[i] = l
		cur--
	}

	ans := 0
	// 何もしない場合の最大値で初期化
	for i := 0; i < n; i++ {
		ChMax(&ans, L[i])
	}

	// すべて試す
	for i := 1; i < len(diff); i++ {
		if diff[i] > 0 && diff[i-1] > 0 {
			continue
		}
		if diff[i] <= 0 && diff[i-1] <= 0 {
			continue
		}

		if diff[i]+diff[i-1] > 0 {
			// この場合のみ意味がある
			ChMax(&ans, L[i-1]+L[i+1]-1)
		}
	}
	fmt.Println(ans)
}

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}
```

本番では階差数列のみをうまく扱って答えを出そうとしたらぐちゃぐちゃになってしまい、
結果バグったコードによりpretestで撃沈してしまいました。
また、そうでなくとも、階差のインデックスと元の配列のインデックスの調整で身長になる必要があり、あまりいい実装が出来ませんでした（というかよくよく見るといろいろな部分が省ける）。

> まずは、もとの配列に対して、その要素が属する連続部分列の長さを計算する。

この部分も結構面倒なのですが、「Union Findを使って増加列を同じグループとし、
それぞれが属するグループのサイズを見れば良い」という方法を
[くるさん](https://profile.hatena.ne.jp/ningenMe/)にご教示いただきました。

あまりUnion Findを活用できていなかったと思うので、これからはもう少し視野を広げてみようと思います
（というよりは問題にたくさん触れるべきなのですが）。

<a id="markdown-dpによる別解" name="dpによる別解"></a>
### DPによる別解

こちらも同じくくるさんの解法を参考にしました。

`dp1[i]: i番目の数値が属する連続部分列に関して、何個目かを記録する`

`dp2[i]: i番目以前の数値を1つ削除したときの、i番目の数値が属する連続部分列に関して、何個目かの最大値を記録する`

以下のコードに示すように、 `dp2 -> dp1` のような遷移がないため、正しく動作します。

```go
var n int
var A []int

var dp1, dp2 [200000 + 5]int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)

	for i := 0; i < n; i++ {
		dp1[i] = 1
	}

	for i := 0; i < n; i++ {
		if i-1 >= 0 && A[i] > A[i-1] {
			ChMax(&dp1[i], dp1[i-1]+1)
			ChMax(&dp2[i], dp2[i-1]+1)
		}
		if i-2 >= 0 && A[i] > A[i-2] {
			ChMax(&dp2[i], dp1[i-2]+1)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ChMax(&ans, dp1[i])
		ChMax(&ans, dp2[i])
	}
	fmt.Println(ans)
}
```

なんと呼ぶべきかはわからないですが、
こういった片方向の遷移を行うものは過去のABCあたりでも解いたことがあるような気がします。

いずれにせよ、頻出パターンの1つではあると思うので、定着させていきたいところです。

<a id="markdown-e-nearest-opposite-parity" name="e-nearest-opposite-parity"></a>
## E. Nearest Opposite Parity

[問題のURL](https://codeforces.com/contest/1272/problem/E)

<a id="markdown-問題の概要-4" name="問題の概要-4"></a>
### 問題の概要

長さ `n` の整数配列 `A` が与えられる。

ある位置 `i` からは、1回の移動で `i - A[i]` もしくは `i + A[i]` へと移動できる。

ある開始位置から移動をはじめて、開始位置の `A[i]` とは偶奇が反対のところへたどり着くためには、
最小で何回の移動が必要か、すべての位置に関して答えよ。

また、到達不可能であるならば `-1` を出力せよ、という問題。

<a id="markdown-解答-4" name="解答-4"></a>
### 解答

公式Editorialをなぞったものです。

まず、ある位置 `i` から `j` まで移動できるとき、これを逆向きに考えて `j -> i` とエッジを張ります。

このようにして出来上がるグラフの辺の数は、最大でも `2*n` であるため、全探索が間に合います。

そこで、偶数の要素をスタート地点とした多点BFSを考え、到達位置へのステップ数を記録します。
探索が完了したら、奇数の要素のステップ数を確認すれば、それが答えとなっています。

今度はこれを偶数と奇数を入れ替えて行うことで、すべての位置について答えを求めることが出来ます。

```go
var n int
var A []int

var G [200000 + 5][]int
var answers []int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)

	for i := 0; i < n; i++ {
		lid, rid := i-A[i], i+A[i]
		if lid >= 0 {
			G[lid] = append(G[lid], i)
		}
		if rid < n {
			G[rid] = append(G[rid], i)
		}
	}

	oddIdxs, evenIdxs := []int{}, []int{}
	for i := 0; i < n; i++ {
		if A[i]%2 == 0 {
			evenIdxs = append(evenIdxs, i)
		} else {
			oddIdxs = append(oddIdxs, i)
		}
	}

	answers = make([]int, n)
	for i := 0; i < n; i++ {
		answers[i] = -1
	}

	bfs(oddIdxs, evenIdxs)
	bfs(evenIdxs, oddIdxs)

	fmt.Println(PrintIntsLine(answers...))
}

func bfs(starts, ends []int) {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF_BIT30
	}

	queue := []int{}
	for _, sidx := range starts {
		dist[sidx] = 0
		queue = append(queue, sidx)
	}

	for len(queue) > 0 {
		cid := queue[0]
		queue = queue[1:]

		for _, nid := range G[cid] {
			if dist[nid] == INF_BIT30 {
				dist[nid] = dist[cid] + 1
				queue = append(queue, nid)
			}
		}
	}

	for _, eidx := range ends {
		if dist[eidx] != INF_BIT30 {
			answers[eidx] = dist[eidx]
		}
	}
}
```

ICPC形式？のため、Dが解けずにコンテスト中は見ることすらできませんでしたが、
復習の際にDFSから考え始めたので論外でした。

Dijkstraしかり、「最短経路」が効かれているのだからBFSを検討しないと土俵にも立てていませんでした。

また、別に逆向きのグラフじゃなくても順方向のグラフで同じことをすればいけるのでは？とか考えましたが、
探索済みの箇所で衝突してしまうため無理でした。

確かAtCoderの[うほょじご](https://atcoder.jp/contests/mujin-pc-2018/tasks/mujin_pc_2018_d)という問題が、
想定解法がこれと似たようなコンセプトだったような気がします（違ったらすみません）。
「逆から考える」っていう典型と見るべきなのでしょうか？

---

そろそろ過去に解けなかった問題もとき直したいところなのですが、時間が。。
