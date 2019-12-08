今回もランレングス符号化が活躍してくれた。

<!-- TOC -->

- [A. Beautiful String](#a-beautiful-string)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Beautiful Numbers](#b-beautiful-numbers)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Beautiful Regional Contest](#c-beautiful-regional-contest)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Beautiful Sequence](#d-beautiful-sequence)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-beautiful-string" name="a-beautiful-string"></a>
## A. Beautiful String

[問題のURL](https://codeforces.com/contest/1265/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

`a, b, c, ?` の4文字からなる文字列が与えられる。

`?` の部分を `a, b, c` のいずれかに置き換えて、同じ文字が連続しないように文字列を構築せよ、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

基本的には、文字列を舐めて前後の文字と違うものを素直に選び続ければ良い。
効率的な実装を心がけたい。

面倒なコーナーケースは潰しておきたいと思ったので、まずは1文字だけの場合を処理する。

2文字以上の場合も、まずは構築不可能な場合、すなわち、最初から連続する部分が存在しないかをチェックする。
このチェックにはランレングス符号化が便利。
`a, b, c` のいずれかが2文字以上続いていたらアウトと判定すれば良い。

それらのチェックが終わったら後は純粋に構築していくだけなので、バグに気をつけながら適宜関数に切り分けたりしてコーディングしていく。

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
	// 1文字のケース
	if len(S) == 1 && S[0] != '?' {
		fmt.Println(string(S))
		return
	}
	if len(S) == 1 && S[0] == '?' {
		fmt.Println("a")
		return
	}

	// 以降は2文字以上

	// 可能かどうか
	pressed, nums := RunLengthEncoding(S)
	for i := 0; i < len(pressed); i++ {
		r := pressed[i]
		cnt := nums[i]

		if r != '?' && cnt > 1 {
			fmt.Println(-1)
			return
		}
	}

	for i := 0; i < len(S); i++ {
		// 決まっていたらパス
		if S[i] != '?' {
			continue
		}

		if i == 0 {
			// 次を見る
			next := S[i+1]
			S[i] = sub(next)
		} else if i == len(S)-1 {
			// 前を見る
			before := S[i-1]
			S[i] = sub(before)
		} else {
			// 前後を見る
			next, before := S[i+1], S[i-1]
			S[i] = subsub(next, before)
		}
	}

	fmt.Println(string(S))
}

// rはa, b, c, ?のいずれかで、r以外のa, b, cのいずれかを返す
// ?だったらaを返す
func sub(r rune) rune {
	if r == 'a' {
		return 'b'
	} else if r == 'b' {
		return 'c'
	} else if r == 'c' {
		return 'a'
	} else {
		return 'a'
	}
}

func subsub(r, s rune) rune {
	memo := make(map[rune]bool)
	memo[r] = true
	memo[s] = true

	isA := memo['a']
	isB := memo['b']
	isC := memo['c']

	if !isA {
		return 'a'
	} else if !isB {
		return 'b'
	} else if !isC {
		return 'c'
	}

	return 'a'
}

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []rune) ([]rune, []int) {
	runes := []rune{}
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
```

Aから20分くらいかかって幸先が良くないですが、特別詰まったわけではないと思い込みつつ先に進みました。

<a id="markdown-b-beautiful-numbers" name="b-beautiful-numbers"></a>
## B. Beautiful Numbers

[問題のURL](https://codeforces.com/contest/1265/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

`1, 2, .., n` までの順列が配列の形で与えられる。

ここから長さ `m` の連続する部分列を切り出したとき、 `1, 2, .., m` の順列になっているかを、
すべての `m` について判定する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

`O(n^2)` は許されない制約であるため、全探索的な愚直な解法は認められない。

元の順列の `1, 2, .., m` の要素の位置を `pos[1], pos[2], .., pos[m]` とする。
ここで `max(pos[i]) - min(pos[i]) + 1 (1 <= i <= m)` によって、 `1, 2, .., m` の要素を含む連続部分配列のうち、
長さが最小のものの長さが手に入る。

この長さが `m` であれば `m` の順列をなす連続部分列が取得できるし、
`m` より大きくなってしまうならば、不純物が混じってしまうため、目的の連続部分列は取得できない（必要十分条件になっている）。

よって、上述の `max, min` を都度更新していけばすべての `m` について高速に判定できる。

計算量は全体で `O(n)` 。

```go
var t int
var n int
var P []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		P = ReadIntSlice(n)

		solve()
	}
}

func solve() {
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		p := P[i]
		pos[p] = i
	}

	answers := make([]rune, n)

	l, r := pos[1], pos[1]
	answers[0] = '1'
	for i := 2; i <= n; i++ {
		ChMin(&l, pos[i])
		ChMax(&r, pos[i])

		dist := r - l + 1
		if dist == i {
			answers[i-1] = '1'
		} else {
			answers[i-1] = '0'
		}
	}

	fmt.Println(string(answers))
}

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
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

<a id="markdown-c-beautiful-regional-contest" name="c-beautiful-regional-contest"></a>
## C. Beautiful Regional Contest

[問題のURL](https://codeforces.com/contest/1265/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

あるプログラミングコンテストにおいて、 `n` 人の解答問題数が配列で与えられるので、
それぞれの人に金・銀・銅メダルを授与することを考える。

以下の条件を満たすように、それぞれのメダルの数を求める問題。

1. それぞれのメダルの数は正の整数である。
2. 金メダルの数は、銀・銅メダルよりも大きい必要がある。ただし、銀メダルと銅メダルの数量関係は不問である。
3. 金メダルを受賞する参加者は、銀メダルを受賞する参加者よりも多くの問題を解いていなければならない。
4. 銀メダルを...、銅メダルを...。
5. 銅メダルを...、メダルを受賞していない参加者よりも多くの問題を解いていなければならない。
6. 3つのメダルの合計個数は、 `Floor(n/2)` 以下である必要がある。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

満たすべき条件が多いが、素直にそれらを満たすように貪欲的に解を構築していく方針で考える。

条件を整理していく。

条件3, 4, 5より、メダルの種類の境界（メダルを授与されない参加者は透明のメダルでももらうと思っておく）において、
問題の正解数に1以上の差がある必要がある、すなわち、正解数が異なっている必要がある。

よって、ここでもランレングス符号化を施しておくと見通しが良くなる。
以降は、ランレングス符号化が行われている前提で話をすすめる。

（※当然ながら、正答数が多いひとを差し置いて、それより正答数が小さい人にメダルを授与することはできないことに注意する。）

条件1より、符号化後に要素数が3以上である必要があるため、そうでない場合は不可能として終了する。

条件2より、金メダルの受賞者は少ないほうが良いため、正解数が最大の人のみとする。

ここからは、銀メダルと銅メダルの受賞者数を、条件2, 6に気をつけながら増やしていけば良い。
銀メダルについては、金メダルの数より大きくなりさえすればよい。
銅メダルについても、まずは金メダルの数よりは大きくし、まだ条件6に対して余裕があるのであれば、
さらに追加していけば良い。

```go
var t int
var n int
var P []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		P = ReadIntSlice(n)

		solve()
	}
}

func solve() {
	g, s, b := 0, 0, 0
	limit := n / 2

	_, counts := RunLengthEncoding(P)
	// 圧縮後に最低長さ3は必要
	if len(counts) < 3 {
		fmt.Println(0, 0, 0)
		return
	}

	// gを決める
	g += counts[0]

	// sを決める、伸ばす
	idx := 1
	for idx < len(counts) {
		// gを超えたらbreak
		if s > g {
			break
		}

		s += counts[idx]
		idx++
	}

	// bを決める、伸ばす
	for idx < len(counts) {
		// gを超えたらbreak
		if b > g {
			break
		}

		b += counts[idx]
		idx++
	}

	if g == 0 || s == 0 || b == 0 {
		fmt.Println(0, 0, 0)
		return
	}
	if s <= g || b <= g {
		fmt.Println(0, 0, 0)
		return
	}
	sum := g + s + b
	if sum > limit {
		fmt.Println(0, 0, 0)
		return
	}

	// bをできるだけ伸ばす
	for idx < len(counts) {
		// 足した結果がlimitを超えるなら足さないでbreak
		if sum+counts[idx] > limit {
			break
		}

		b += counts[idx]
		sum += counts[idx]
		idx++
	}

	fmt.Println(g, s, b)
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
```

<a id="markdown-d-beautiful-sequence" name="d-beautiful-sequence"></a>
## D. Beautiful Sequence

[問題のURL](https://codeforces.com/contest/1265/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

`0, 1, 2, 3` の数字がそれぞれ `a, b, c, d` 個与えられるので、
これらをすべて使い切って1列に並べる。

出来上がった数列に関して、隣り合う数字の差をすべて `1` とすることができるか判定し、
可能ならば具体的に構築せよ、という問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

判定だけなら結構簡単かもしれないが、構築も要求されるとすごく難しい。

（以下はeditorialそのままです。）

まず、作るべき数列に関して、偶数は `0, 0, .., 0, 2, 2, .., 2` と並べるのが最適となる。
こうすることで、 `3` を置くことができる箇所を最大化できる一方で、 `1` はどの間でも置くことができる。

※この時点で少し難しい気がする。
構築可能なときに必要な `1` の数が変わっておらず、かつ `3` を受け入れる箇所を最大化出来ているから最適、と読めば良い？

この後は、まず `3` を `2` の間に並べて、隙間の数が足りないならば端に置く、
続いて `1` を残っている隙間に入れて、足りなければやはり端に置く、とすれば良い。

しかし、具体的に条件分岐させて一発での構築を目指すと、かなり複雑になると思う（というより私は諦めました）。

editorialの実装例では、以下のように開始する数に関して全探索を行うような賢い実装を行っている。

```go
var A map[int]int

func main() {
	a, b, c, d := ReadInt4()
	total := a + b + c + d

	for i := 0; i < 4; i++ {
		A = make(map[int]int)
		A[0], A[1], A[2], A[3] = a, b, c, d
		answers := []int{}

		if A[i] > 0 {
			last := i
			answers = append(answers, last)
			A[last]--
			for {
				if A[last-1] > 0 {
					A[last-1]--
					answers = append(answers, last-1)
					last--
				} else if A[last+1] > 0 {
					A[last+1]--
					answers = append(answers, last+1)
					last++
				} else {
					break
				}
			}

			if len(answers) == total {
				fmt.Println("YES")
				fmt.Println(PrintIntsLine(answers...))
				return
			}
		}
	}

	fmt.Println("NO")
}
```

配列ではなく辞書を使うことで `last-1` の負数判定をなくしたりと、きれいな実装の参考になりました。

---

AtCoder緑ぐらいの人たちでランレングス符号化を用意していない人が居たら、ぜひ用意してみていただきたいです。
意外と便利なので。

