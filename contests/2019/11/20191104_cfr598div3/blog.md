「軽量サイトが動いてさえいればこどふぉはrated」ということを学びました。

<!-- TOC -->

- [A. Payment Without Change](#a-payment-without-change)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Minimize the Permutation](#b-minimize-the-permutation)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Platforms Jumping](#c-platforms-jumping)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Binary String Minimizing](#d-binary-string-minimizing)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-payment-without-change" name="a-payment-without-change"></a>
## A. Payment Without Change

[問題URL](https://codeforces.com/contest/1256/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

`n` 円の硬貨 `a` 枚と `1` 円の硬貨 `b` 枚でちょうど `S` 円払えるか判定する問題。

<a id="markdown-解答" name="解答"></a>
### 解答

貪欲的に、 `n` 円で払えるだけ払って差額は1円で払う方法を考える。

`x = Floor(S/n)` が `S` 円に対してできるだけ多くの `n` 円硬貨で払う場合の `n` 円硬貨の枚数。

`x <= a` ならば `n` 円硬貨の手持ちの枚数は十分なので、差額の分 `1` 円硬貨があるかどうかを判定すれば良い。

```go
var q int

func main() {
	q = ReadInt()

	for tc := 0; tc < q; tc++ {
		a, b, n, s := ReadInt64_4()
		x := s / n

		if x <= a && (s-x*n) <= b {
			fmt.Println("YES")
		} else if x > a && (s-a*n) <= b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
```

公式editorialのほうがシンプルでいいですね。

<a id="markdown-b-minimize-the-permutation" name="b-minimize-the-permutation"></a>
## B. Minimize the Permutation

[問題URL](https://codeforces.com/contest/1256/problem/B)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

`1` から `n` までの順列が与えられる。

`1` から `n-1` までの値を取る `i` を考え、各 `i` は最大1回まで選べるとする。
選んだ `i` に対して `i+1` と要素を交換できる。
このような操作は、 `i` を選ぶ順番については自由に選んで良いものとする。

このような操作を行ったとき、できるだけ辞書式順序を小さくしようとした場合、最小はどの様になるか求める問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

選べる範囲の中から最小のものを見つけ、それを移動できる限り前まで運ぶようなシミュレーションを行えば良い。

そうすると、最初は必ず `1` を選ぶ事になり、これを一番前まで運ぶことにより、
`i` として `1` から `(1の初期位置)` までを1回使ったことになる。
あるいは、最初から `1` が1番目にいる場合は、 `i` として `1` を選ぶ必要はないが、
その後、その位置と別の数字がスワップされることはない。

`1` の初期位置を `idx` とすると、次は `[idx, n]` の中から最小のものを選んで、
それを可能な限り前に運べば良い（ただし、 `1` の初期位置が最初から `1` だった場合は、 `[idx+1, n]` から探索する必要がある）。

正直、思いつくのは簡単で実装がちょっと大変という感じだと思う。

```go
var q int

func main() {
	q = ReadInt()
	for tc := 0; tc < q; tc++ {
		n := ReadInt()
		A := ReadIntSlice(n)

		s, t := 0, n
		for s < n-1 {
			idx := findMinimum(A, s, t)
			for i := idx - 1; i >= s; i-- {
				A[i], A[i+1] = A[i+1], A[i]
			}

			if idx == s {
				s++
			} else {
				s = idx
			}
		}
		fmt.Println(PrintIntsLine(A...))
	}
}

func findMinimum(A []int, s, t int) int {
	mini := 1000
	res := -1
	for i := s; i < t; i++ {
		if mini > A[i] {
			mini = A[i]
			res = i
		}
	}
	return res
}
```

コンテスト中はバブルソートの転倒数とかそれらしい問題かと思ったら全然違ってWA連発していました。
（誤読に気づいたときには既にこどふぉが落ちていてとてもつらかった。）

<a id="markdown-c-platforms-jumping" name="c-platforms-jumping"></a>
## C. Platforms Jumping

[問題URL](https://codeforces.com/contest/1256/problem/C)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

`[1, n]` が水たまりの道に対して、 `m` 枚の板を使って、
ジャンプ力 `[1, d]` の人が水たまりに足をつけずに足場 `0` から 足場 `n+1` まで到達できるか判定し、
到達できるのなら板の置き方まで構築する問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

板の情報がランレングス圧縮っぽいと考えた。

すると、最終的に求められる結果が、板の配列の情報を解凍し、さらに `m+1` 箇所に `0` を複数個挿入する形であると捉えた。

板をすべて使い切ることを考えると、挿入する `0` の数は `n - (板の長さの総和)` で計算できる。

また、 `m+1` 個の複数個の `0` が挿入される区間は、ジャンプで飛び越える必要がある区間である。
そのため、最大で `d-1` 個までの `0` しか挿入できない。

一方で、 `m+1` 箇所に均等に `0` を挿入したとしても、
少なくともある1箇所で最大 `Ceil((0の個数) / (m+1))` 個の `0` を挿入する必要がある。

よって、この値が `d-1` 以下であれば対岸まで到達可能と判定できる。

構築方法としては、 `m+1` 箇所に対してジャンプで飛び越えられるギリギリまで `0` を挿入し、
挿入すべき `0` が尽きたならば、あとは板の番号だけをプッシュしていけばよい。

```go
var n, m, d int
var C []int

func main() {
	n, m, d = ReadInt3()
	C = ReadIntSlice(m)

	csum := Sum(C...)

	zeroNum := n - csum
	avg := (zeroNum + ((m + 1) - 1)) / (m + 1)

	id := 1
	if avg <= (d - 1) {
		fmt.Println("YES")
		answers := []int{}
		for i := 0; i < m; i++ {
			// 0を埋める
			for j := 0; j < d-1; j++ {
				if zeroNum == 0 {
					break
				}
				answers = append(answers, 0)
				zeroNum--
			}

			// C[i]個埋める
			for j := 0; j < C[i]; j++ {
				answers = append(answers, id)
			}

			id++
		}
		for zeroNum > 0 {
			answers = append(answers, 0)
			zeroNum--
		}
		fmt.Println(PrintIntsLine(answers...))
	} else {
		fmt.Println("NO")
	}
}
```

解けてたんだから諦めずに軽量サイトに投げればよかった。

<a id="markdown-d-binary-string-minimizing" name="d-binary-string-minimizing"></a>
## D. Binary String Minimizing

[問題URL](https://codeforces.com/contest/1256/problem/D)

<a id="markdown-問題の要約-3" name="問題の要約-3"></a>
### 問題の要約

2進数の文字列が与えられるので、隣同士のスワップ操作を最大 `k` 回まで行えるとき、
構築可能な辞書順最小の文字列を答える問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

素直にスワップをシミュレーションするとTLEするので工夫する。

左から文字列を走査したときに出会った `0` に関して、
その `0` を優先的にできるだけ左に移動させるように、貪欲的に考えれば良い。

文字列を走査する過程で暫定で見つけた `0` の数をカウントして保持しておくことで、
`0` を見つけるたびにその `0` を一番前に持ってくるのに必要なステップ数が `(0の初期位置) - (0の暫定登場回数)` で求まるので、
これを前処理して記憶しておく。

さらに、このステップ数に関して累積和を計算しておくと、
`k` と比較することで、確実に前に寄せることのできる `0` の個数がわかる。

一旦、左に寄せ切ることのできる `0` だけを左に寄せて、あとは余分なスワップ操作を行わないで手に入る配列を構築する。
この配列に対して、残った手数で、左に寄せられなかった中で一番左に残存する `0` を、できる限りスワップして左に寄せてやれば良い。
これらは `O(n)` で可能なので間に合う。

```go
var q int
var n, k int64
var S []rune

func main() {
	q = ReadInt()
	for tc := 0; tc < q; tc++ {
		n, k = ReadInt64_2()
		S = ReadRuneSlice()

		solve()
	}
}

func solve() {
	zeroNum := 0
	for i := int64(0); i < n; i++ {
		if S[i] == '0' {
			zeroNum++
		}
	}
	zeros := make([]int64, zeroNum)

	j := int64(0)
	for i := int64(0); i < n; i++ {
		if S[i] == '0' {
			zeros[j] = i - j
			j++
		}
	}

	zeroSums := make([]int64, zeroNum+1)
	for i := 0; i < zeroNum; i++ {
		zeroSums[i+1] = zeroSums[i] + zeros[i]
	}

	// c: 確実に先頭に持ってくることができる0の数
	c := 0
	for i := 0; i < len(zeroSums); i++ {
		if zeroSums[i] > k {
			break
		} else {
			c = i
		}
	}

	answers := []rune{}
	for i := 0; i < c; i++ {
		answers = append(answers, '0')
	}
	skipped := 0
	for i := int64(0); i < n; i++ {
		if S[i] == '0' && skipped < c {
			skipped++
		} else {
			answers = append(answers, S[i])
		}
	}

	if c == zeroNum {
		fmt.Println(string(answers))
		return
	}

	k -= zeroSums[c]
	for i := int64(c); i < n; i++ {
		if answers[i] == '0' {
			for j := i; j > 0; j-- {
				if k == 0 {
					break
				}
				if answers[j-1] > answers[j] {
					answers[j], answers[j-1] = answers[j-1], answers[j]
					k--
				}
			}

			fmt.Println(string(answers))
			return
		}
	}
}
```

実装方法は色々ありそうですが、自分としては一番しっくり来るものを選んだつもりです。

---

色々文句をつけたいコンテストだったけど、
このレベルをサイトが落ちるまでに早解きできない自分が悪いと思って修行します。

