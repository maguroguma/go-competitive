点数的にCを解きたかったけど、ちゃんとD1に移って得点を確保できたのは、コンテストムーブとしては評価してあげたい。

※Cの数え上げはAtCoderでも活きそうな価値の高そうな香りを感じるので、取り組み次第追記します。

※D2はDiv.1レベルの人もかなり苦戦している問題のようで、おそらくしばらくは取り組まないかと思います。

<!-- TOC -->

- [A. Integer Points](#a-integer-points)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Grow The Tree](#b-grow-the-tree)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [D1. The World Is Just a Programming Task (Easy Version)](#d1-the-world-is-just-a-programming-task-easy-version)
	- [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
	- [公式Editorialの解法について](#%e5%85%ac%e5%bc%8feditorial%e3%81%ae%e8%a7%a3%e6%b3%95%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6)
		- [なぜこれでうまくいくか？](#%e3%81%aa%e3%81%9c%e3%81%93%e3%82%8c%e3%81%a7%e3%81%86%e3%81%be%e3%81%8f%e3%81%84%e3%81%8f%e3%81%8b)

<!-- /TOC -->

<a id="markdown-a-integer-points" name="a-integer-points"></a>
## A. Integer Points

[問題URL](https://codeforces.com/contest/1248/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

与えられた直線群の格子点の数を求める問題。

<a id="markdown-解答" name="解答"></a>
### 解答

問題だけみると難しくも見えそうだが、与えられた直線が傾き1か-1のいずれかで切片が異なるだけなので、
整理するととても簡単になる。

平行な直線は交わらないので無視して、傾きが異なる者同士を方程式にして解いてみると、
`x = (q-p)/2` となる。

この `x` が整数値となれば同時に `y` も整数値となるため、交点は格子点となる。

ただし、それぞれの傾きの直線の数 `n, m` の数は大きいので、すべての `p, q` の組み合わせについて知らべるとTLEしてしまう。

`p, q` の組み合わせを考えたとき、それぞれの偶奇性が一致していれば格子点になると読み替えられるので、
配列 `P, Q` の偶数・奇数の数をそれぞれカウントして、それらの掛け算で答えを求める。

```go
var t int

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		n := ReadInt()
		P := ReadIntSlice(n)
		m := ReadInt()
		Q := ReadIntSlice(m)

		peven := 0
		qeven := 0
		for j := 0; j < n; j++ {
			if P[j]%2 == 0 {
				peven++
			}
		}
		for j := 0; j < m; j++ {
			if Q[j]%2 == 0 {
				qeven++
			}
		}

		ans := int64(peven)*int64(qeven) + int64(n-peven)*int64(m-qeven)

		fmt.Println(ans)
	}
}
```

`10^5` に引っ張られて32bitでのオーバーフローの可能性を捨てないようにしましょう（1敗）。

<a id="markdown-b-grow-the-tree" name="b-grow-the-tree"></a>
## B. Grow The Tree

[問題URL](https://codeforces.com/contest/1248/problem/B)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

与えられた棒を使って、2次元平面上に端点同士を繋げる形で、またx, y軸にそれぞれ平行となるように配置して「木」を作る。
木の片方の端点を原点に置くとき、もう片方の端点が原点からできるだけ遠くになるように配置したとき、
その長さの2乗を答える問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

x, y軸に平行になるように棒を設置するにあたって、一度決めた方向とは逆に伸ばすことは、明らかに無意味である。

このようにすると、距離は三平方の定理で `(x軸方向に伸びた距離)^2 + (y軸方向に伸びた距離)^2` で求まる。

なんとなくサンプルを見ても、長い棒はできるだけ同じ軸に平行に伸ばす方向に設置したく成るし、実際にそうすれば良い。

```go
var n int
var A []int

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)

	sort.Sort(sort.IntSlice(A))

	l, r := 0, n-1
	hori, verti := int64(0), int64(0)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			// 後から
			a := A[r]
			r--
			hori += int64(a)
		} else {
			// 前から
			a := A[l]
			l++
			verti += int64(a)
		}
	}

	ans := hori*hori + verti*verti
	fmt.Println(ans)
}
```

簡易的な証明として、はじめに決めたルール通りに一度決めた方向に伸ばすようにした場合、
x, y軸に伸ばす距離の和は、どのような並べ方をしても一定である。

x, y軸に伸ばす距離を `a, b` また和を `c = a + b` とする。

`(a+1)^2 + (b-1)^2` を考えると `a^2 + 2*a + 1 + b^2 - 2*b + 1 = a^2 + b^2 + 2*(a - b + 1)` となるが、
これは `a >= b` ならばもとの `a, b` の状態で計算される原点からの距離よりも確実に大きくなる。

よって、一方向についてより長いものを優先的に、もう片方は短いものを優先的に選ぶ、上のアルゴリズムで最大化できる。

。。コンテスト中は、なんとなく一般的に成り立ちそうだな、ぐらいで流してすぐに提出してしまいましたが、
これぐらいならそのムーブで正解かなと思います。

<a id="markdown-d1-the-world-is-just-a-programming-task-easy-version" name="d1-the-world-is-just-a-programming-task-easy-version"></a>
## D1. The World Is Just a Programming Task (Easy Version)

[問題URL](https://codeforces.com/contest/1248/problem/D1)

なんだこの問題名は。

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

与えられた丸括弧からなる文字列について、「一度だけ」好きな位置間で交換して良い（同じ場所を交換しても良い＝交換しなくても良い）。

この状態で "cyclical shift" なるものを考える。
これは、後 `k` 文字を前 `n-k` 文字の前に移動させてできる文字列のことである。

括弧文字列の美しさを、「すべてのcyclical shiftに対して、正しい括弧列を形成するものの数」と定義する

ある与えられた括弧文字列に対して、前述の交換を施し、考えられる美しさの最大値を答えよ。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

交換するインデックスについては1-basedなのに、cyclical shiftを考えるときは0-basedっぽいことに注意。

制約が `n <= 500` と小さめなので、全探索を検討する。

すべての交換のペアが `500^500` であり、交換の結果できた文字列の美しさの計算が `O(n)` でできれば良い。

基本的には、スタックを使った字句解析を行うことで正しい括弧列をつくるcyclical shiftの数を求めた。
もろもろを計算するために、括弧用の `char` 型とインデックス用の `int` 型で構造体を定義し、
このインスタンスを用いたスタックを使う。



図に示すように、スタックに残った開き括弧と閉じ括弧について着目する。
このインデックスの間に存在する、正しい括弧列の一番外側のものがcyclical shiftの文字列切り出しの起点となる場合、
cyclical shiftは全体で見て正しい括弧列となる（ただし、図に示すように、いくつかの場合分けが必要）。
また、スタックに残っている開き括弧の最初の位置で切り出しても正しい括弧列となるため、
コード中では最後にインクリメントして返している。

上の説明における、「正しい括弧列の一番外側のもの（図中の赤線のインデックス）」は、
スタックで処理する過程で開き括弧が消えるとき、スタックトップが閉じ括弧、もしくはスタックが空となる場合に、
その開き括弧のインデックスを記憶しておけば良い（コード中の変数 `memo` に該当）。

```go
var n int
var S []rune

func main() {
	n = ReadInt()
	S = ReadRuneSlice()

	ans := 0
	l, r := 0, 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 入れ替える
			S[i], S[j] = S[j], S[i]

			tmp := sub()
			if ans < tmp {
				ans = tmp
				l, r = i, j
			}

			// もとに戻す
			S[i], S[j] = S[j], S[i]
		}
	}
	fmt.Println(ans)
	fmt.Println(l+1, r+1)
}

type Char struct {
	idx int
	c   rune
}

func sub() int {
	// スタックで字句解析する
	s := make([]Char, 0, 1000)
	memo := []int{}
	for i := 0; i < n; i++ {
		r := S[i]
		char := Char{idx: i, c: r}

		if len(s) == 0 || r == '(' {
			s = append(s, char)
			continue
		}

		if r == ')' && s[len(s)-1].c == '(' {
			popped := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 || s[len(s)-1].c == ')' {
				memo = append(memo, popped.idx)
			}
		} else {
			s = append(s, char)
		}
	}

	// スタックに何も残っていないときはmemoの長さ
	if len(s) == 0 {
		return len(memo)
	}

	// 残ったカッコの数をチェック
	op, cl := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i].c == '(' {
			op++
		} else {
			cl++
		}
	}
	// 数が一致しないとダメ
	if op != cl {
		return 0
	}

	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i].c == '(' {
			r = s[i].idx
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i].c == ')' {
			l = s[i].idx
			break
		}
	}

	ans := 0
	for i := 0; i < len(memo); i++ {
		if l < memo[i] && memo[i] <= r {
			ans++
		}
	}

	return ans + 1
}
```

解いてる人の人数的にCより高得点がほしかった。

<a id="markdown-公式editorialの解法について" name="公式editorialの解法について"></a>
### 公式Editorialの解法について

正直コンテスト中に書いてて辛かったので、もっとスマートに考えるべきだと思い、Editorialも参照しました。

> Note that the answer to the question about the number of cyclic shifts,
> which are correct bracket sequences, **equals the number of minimal prefix balances.**

具体的にはEditorialにもあるように、閉じ括弧を `-1` 、開き括弧を `1` としたときに、
全体文字列の累積和を計算し、累積和中の最小値の数が答えとなります。

以下の図に示すように、累積和が最小値を取るところでcyclical shiftを考えると、たしかにそれは全体で正しい括弧列となります。



<a id="markdown-なぜこれでうまくいくか" name="なぜこれでうまくいくか"></a>
#### なぜこれでうまくいくか？

正直、自分はぱっと理解できなかったので、もう少し掘り下げました。

前提として開き括弧と閉じ括弧の数は等しいため、この累積和は最終的に `0` となります。
つまり、ある部分で累積和がマイナス、すなわち閉じ括弧が過剰になったとしても、
その後には過剰分の閉じ括弧を消化できる量の開き括弧が「余分に」存在することが保証されています。

また、cyclic shiftを考えたときに、消化されない余分な閉じ括弧が文字列全体で前方に移動させることはできません。
そのような位置で切り出さないために、
余分な閉じ括弧の数が極大となる「累積和全体での最小値」について着目しなければなりません。

---

括弧列に関してEditorialのような扱い方をしたことがなかったので、覚えておきたいところ。

（[天下一魔力発電](https://tenka1-2016-qualb.contest.atcoder.jp/tasks/tenka1_2016_qualB_b)とかそんなやり方だったっけ。。？）
