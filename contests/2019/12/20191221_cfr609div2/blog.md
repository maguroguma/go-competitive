<!-- Codeforces Round No.609 Div.2 参加記録 (A〜D解答) -->

青になれました！やったね！

「Codeforcesの青はAtCoderの水色」ぐらいの言説を見た気がするので、
「まぁ青ぐらいなら楽勝だろう」と思ったのですが、想像の3倍ぐらいは難しかったです。

多分青を維持するのは今の自分にはきつくて、
自分の体感では、Codeforcesの1600はAtCoderの1400（水色の真ん中）ぐらいなんだろうなぁと感じています。

とはいえ、ミスは減ってきて実装力も向上しているとは思うので、
今後もしばらくはこの調子で頑張りたいところです。

<!-- TOC -->

- [A. Equation](#a-equation)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Modulo Equality](#b-modulo-equality)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Long Beautiful Integer](#c-long-beautiful-integer)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Domino for Young](#d-domino-for-young)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-equation" name="a-equation"></a>
## A. Equation

[問題のURL](https://codeforces.com/contest/1269/problem/A)

<a id="markdown-解答" name="解答"></a>
### 解答

式変形すると `a = n + b` であるため、小さい方である `b` が決まれば `a` も自然と決まる。
そして、 `a, b` がともに合成数になるようにする。

`2` 以外の偶数はすべて合成数であることを踏まえると、

1. `n` が偶数ならば、 `b` を適当な `2` 以上の偶数とすれば `n + b` も当然 `2` 以上の偶数となる。
2. `n` が奇数ならば、 `b` を適当な合成数である奇数とすれば `n + b` は `2` 以上の偶数となる。

であり、 `a, b` を選択できる。

```go
var n int

func main() {
	n = ReadInt()

	if n%2 == 0 {
		b := 4
		a := n + b
		fmt.Println(a, b)
	} else {
		b := 9
		a := n + b
		fmt.Println(a, b)
	}
}
```

こんな問題に9分もかかってしまいました。

というより、公式Editorialを見ると、「`9*n, 8*n` でOK」とのこと、たしかに。

こどふぉのDiv2のA問題って、
`O(1)` 算数をしないといけなかったり、真面目に全探索しないといけなかったりで、
なかなか難しいです。

簡単と思える日が来てほしい。

<a id="markdown-b-modulo-equality" name="b-modulo-equality"></a>
## B. Modulo Equality

[問題のURL](https://codeforces.com/contest/1269/problem/B)

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

`A, B` それぞれについて、登場する値のカウントをメモしておく。

答えが `x` のとき、すべての `A[i]` に対して、
`(AにおけるA[i]の個数) == (Bにおける(b = (A[i] + x) % m)の個数)` となっている必要がある。

よって、このような `x` としてあり得るものをすべて列挙し、
それが条件を満たすか確かめれば良い。
ただし、求められている `x` は条件を満たす最小のものであるため、見つけるたびに小さい方で更新する。

`x` としてありうるものは、ある適当な `A` の値（以下のコード中では `A[0]` としている）を取り出し、
これが `B` のどの要素に遷移するのか？を考えることで列挙できる。

また、各 `x` に対して、すべての `A` の要素の遷移先の個数が、
`A` の遷移前の個数と一致するかをチェックすれば良い。

以上より、トータルの計算量は `O(n^2)` となり、十分高速である。

```go
var n, m int
var A, B []int

func main() {
	n, m = ReadInt2()
	A = ReadIntSlice(n)
	B = ReadIntSlice(n)

	amemo, bmemo := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		amemo[A[i]]++
		bmemo[B[i]]++
	}

	ans := -1
	aval, anum := A[0], amemo[A[0]]
	for bval, bnum := range bmemo {
		if anum != bnum {
			continue
		}

		var x int
		if bval > aval {
			x = bval - aval
		} else {
			x = bval + m - aval
		}

		isOK := true
		for av, an := range amemo {
			bv := (av + x) % m
			bn := bmemo[bv]
			if an == bn {
				continue
			} else {
				isOK = false
				break
			}
		}

		if isOK {
			if ans == -1 {
				ans = x
			} else {
				ChMin(&ans, x)
			}
		}
	}

	if ans == m {
		ans = 0
	}

	fmt.Println(ans)
}
```

`A[i], B[i]` の範囲が `10^9` なので固定長配列ではなく辞書を使う必要があるわけですが、
あまり慣れていない操作だったので時間がかかってしまいました。

そろそろこどふぉも100問弱ぐらいACしているはずですが、まだまだ足りてないなぁと感じます。

<a id="markdown-c-long-beautiful-integer" name="c-long-beautiful-integer"></a>
## C. Long Beautiful Integer

[問題のURL](https://codeforces.com/contest/1269/problem/C)

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

自由が効くのは最初の `k` 桁までであり、以降の桁は `k` 桁までの設定に従属する。

できるだけ小さいBeautifulな数値を作ろうとすると、
貪欲に上の桁から `A` の桁をコピーする必要があるため、まずはそのように桁を埋めていく。

このようにして出来る数値が `A` 以上であるならば、それが答えとなる。

もし `A` 未満となってしまう場合は、自由が効く桁の一番下の桁、すなわち `k` 桁目から微調整することを考える。

`k` 桁目までは `A` のコピーになっているため、できるだけ下の方の桁で `+1` できれば良い
（必ず `A` よりも大きくなる）。
ただし、注目桁が `9` である場合は繰り上がりを考える必要がある。
その場合は、 `9` の桁を `0` にした後、上の桁に移動してその桁の数字を `+1` することを再考する。

1桁でも変更ができたらその時点で終了し、出来上がった数を出力すれば良い。

`k` 桁目までを変更する際は、下の方の桁も合わせて修正する必要があることには注意する。

```go
var n, k int
var A []rune

func main() {
	n, k = ReadInt2()
	A = ReadRuneSlice()

	isBeau := true
	for i := 0; i < n; i++ {
		if i+k < n {
			if A[i] != A[i+k] {
				isBeau = false
				break
			}
		} else {
			break
		}
	}
	if isBeau {
		fmt.Println(len(A))
		fmt.Println(string(A))
		return
	}

	B := make([]rune, n)
	for i := 0; i < n; i++ {
		B[i] = 'x'
	}

	// iは起点
	for i := 0; i < k; i++ {
		digit := A[i]

		// jは飛び飛び
		for j := i; j < n; j += k {
			B[j] = digit
		}
	}

	isSmall := false
	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			continue
		} else if A[i] < B[i] {
			break
		} else {
			isSmall = true
			break
		}
	}
	if !isSmall {
		fmt.Println(len(B))
		fmt.Println(string(B))
		return
	}

	// 大きくなるように正しく調整する
	for i := k - 1; i >= 0; i-- {
		if B[i] == '9' {
			// iからk飛びで0に更新
			for j := i; j < n; j += k {
				B[j] = '0'
			}
		} else {
			// iからk飛びで更新
			for j := i; j < n; j += k {
				B[j]++
			}
			break
		}
	}

	fmt.Println(len(B))
	fmt.Println(string(B))
}
```

桁DPでよくやる考え方の基本、という感じがしました。

多分、最初の部分の「`A` がすでにBeautifulかどうか？」の判定は要らないと思います。

桁数を出力するのを忘れて1WA、 `isSmall` の判定部分でバグっており2WAしてしまい、
もったいないことをしてしまいました。

実装がちょっと長くなってくるとこういったミスが増えてくるので、
なにか対策を考えます。

<a id="markdown-d-domino-for-young" name="d-domino-for-young"></a>
## D. Domino for Young

[問題のURL](https://codeforces.com/contest/1269/problem/D)

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

わからなかったため、公式Editorialの手法をなぞりました。

与えられたヤング図形を市松模様（チェッカーフラグ）のように塗ります。

その時の黒と白の数を数えて、小さい方が答えです。

小さい方の色を仮に黒とすると、黒は隣接する白とで1つのドミノによって専有され、
また、すべての未使用の黒は未使用の白とペアにすることができるため、とのこと。

賢い（けど、上の説明はちょっと大雑把すぎるかもしれないです）。

```go
var n int
var A []int64
var b, w int64

func main() {
	n = ReadInt()
	A = ReadInt64Slice(n)

	b, w = 0, 0
	for i := 0; i < n; i++ {
		l := A[i]
		if i%2 == 0 {
			// Bから塗り始める
			b += (l + (2 - 1)) / 2
			w += l / 2
		} else {
			// Wから塗り始める
			w += (l + (2 - 1)) / 2
			b += l / 2
		}
	}

	fmt.Println(Min(b, w))
}
```

---

青になれて嬉しい気持ちはありますが、目標はまだまだ上に持てているので、
引き続き頑張っていく所存です。

