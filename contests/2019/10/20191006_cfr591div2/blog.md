初めての unrated codeforces を体験してしまいました。

<a id="markdown-全体" name="全体"></a>
## 全体

<!-- TOC -->

- [全体](#%e5%85%a8%e4%bd%93)
- [A. CME](#a-cme)
  - [問題](#%e5%95%8f%e9%a1%8c)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Strings Equalization](#b-strings-equalization)
  - [問題](#%e5%95%8f%e9%a1%8c-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Save the Nature](#c-save-the-nature)
  - [問題](#%e5%95%8f%e9%a1%8c-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)
- [感想](#%e6%84%9f%e6%83%b3)

<!-- /TOC -->A, B, Cがコンテスト後に確認したところちゃんと通っており、Dは嘘解法だったようでpretestで弾かれていました。

Dも勉強になりそうな雰囲気なので、後日解説ACの上追記していきます。

<a id="markdown-a-cme" name="a-cme"></a>
## A. CME

[問題URL](https://codeforces.com/contest/1241/problem/A)

<a id="markdown-問題" name="問題"></a>
### 問題

※長いので意訳。

`q` 個のクエリに対して `n` が与えられる。

`n` 本のマッチ棒を必ず使い切り `a + b = c` という数式を満たすようにマッチ棒を `a, b, c` に割り当てる。

ただし、 `a, b > 0` でなければならない。

数式を作る際にマッチ棒が足りない場合、好きな本数買い足すことができる。

各 `n` に対して、買い足す必要のあるマッチ棒の本数の最小値をそれぞれ答えよ。

制約: `1 <= q <= 100, 2 <= n <= 10^9`

<a id="markdown-解答" name="解答"></a>
### 解答

必要なマッチの総本数は明らかに偶数であり、逆に偶数本マッチがあればCMEを作れる。

なので、 `n` が偶数ならば `0` 、奇数ならば `1` とすれば良い。

ただし、 `n = 2` のときは `a + b = 1` となり、いずれかが `0` となって条件を満たせなくなる。
これを避けるために、左辺と右辺にそれぞれ `1` ずつ足す必要があり、`2` 本必要となる。

```go
var q int

func main() {
	q = ReadInt()
	for i := 0; i < q; i++ {
		n := ReadInt()

		if n%2 == 0 {
			if n/2 == 1 {
				fmt.Println(2)
			} else {
				fmt.Println(0)
			}
		} else {
			fmt.Println(1)
		}
	}
}
```

上で「明らかに」とか言ってますが、コンテスト中は結構数式こねくり回してますし、
前日のAGCで太陽拝んでしまった後遺症でコーナーケース探りまくっていて、
提出に15分かかっています。

<a id="markdown-b-strings-equalization" name="b-strings-equalization"></a>
## B. Strings Equalization

[問題URL](https://codeforces.com/contest/1241/problem/B)

<a id="markdown-問題-1" name="問題-1"></a>
### 問題

2つの英小文字からなる、同じ長さの文字列 `s, t` が与えられる。
「操作」は何回行っても良い（0回でも良い）。

操作では、どちらの文字列についてでも、2つの隣り合った文字に関して、1つ目に選んだ文字を2つ目に選んだ文字に代入して良い。

`q` 個のクエリが与えられるので、それぞれの `s, t` について、任意回数の操作のもとで `s, t` を等しくできるかどうか判定せよ。

制約: `1 <= q <= 100, 1 <= |s| = |t| <= 100`

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

制約が小さく、文字列を全部舐めても大丈夫なことを抑える。

「何回でも操作ができる」というのが強力で、ある文字列中に存在する1文字によって、
それを左右に伝搬する形でその文字列の他の文字列を上書きしてしまえる。

よって、 `s, t` について舐めて、「それぞれに同じ文字が1文字でも存在すればOK、なければNG」とする。

```go
var q int

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		S := ReadRuneSlice()
		T := ReadRuneSlice()

		smemo := [ALPHABET_NUM]int{}
		tmemo := [ALPHABET_NUM]int{}

		for j := 0; j < len(S); j++ {
			s, t := S[j], T[j]
			smemo[s-'a']++
			tmemo[t-'a']++
		}

		flag := false
		for j := 0; j < ALPHABET_NUM; j++ {
			if smemo[j] > 0 && tmemo[j] > 0 {
				flag = true
				break
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
```

これも10分弱悩まされてしまったのが悔しい。

<a id="markdown-c-save-the-nature" name="c-save-the-nature"></a>
## C. Save the Nature

[問題URL](https://codeforces.com/contest/1241/problem/C)

<a id="markdown-問題-2" name="問題-2"></a>
### 問題

※長いので意訳。

`n` 枚のチケットを好きな順番で売ることができる。
売上の一部は以下の規則のもとで、環境保全基金に寄付される。

- `a` の倍数番目で売られたチケットの `x` ％が寄付される。
- `b` の倍数番目で売られたチケットの `y` ％が寄付される。
- `a, b` の公倍数番目で売られたチケットの `x + y` ％が寄付される。

できるだけ少ない枚数のチケット販売で目標トータル寄付金額 `k` を寄付したい。

目標を達成するためのチケット販売枚数の最小値を答えよ。
また、全チケットを販売しても目標が達成できない場合は `-1` を出力せよ。

クエリが `q` 個与えらるため、それぞれについて答えよ。

制約:

- `1 <= q <= 100`
- `1 <= n <= 2 * 10^5`
- `100 <= p[i] <= 10^9, p[i] % 100 = 0`
- `1 <= a, b <= n, 1 <= x, y <= 100, x + y <= 100`
- `1 <= k <= 10^14`
- チケットの全クエリの合計は `2 * 10^5` を超えない。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

当然ながら、各クエリについてすべてのチケットの順列を総当りすることはできない。

また、チケットの金額が高いものほど、寄付割合の高い順番に配置したいが、
できるだけ少ない枚数のチケットで目標を達成したいという縛りがあるため、
貪欲に `a, b` の公倍数番目に高いチケットを配置するわけにはいかない。

とりあえず「答えを `m` 枚として、目標を達成できる最適なチケットの並べ方」をイメージしてみる。

このような並び方を考えたとき、 `m` 枚のチケットは、すべてのチケットを高い順に並べたときの金額の上位 `m` 枚となるはずである。
（仮に、必ずしも高い順に `m` 枚となっていなくても目標を達成できる場合、上位 `m` 枚に入らないチケットと上位 `m` 枚のチケットと交換しても、金額的に損しないので変わらずに目標を達成できる。）

また、 `m` 枚のチケットで目標が達成できるとき、 `m + 1` 以上のチケットでも当然目標が達成できる。

よって、二分探索で `m` の境界を探索してやればよい。

`m` 枚で条件を達成できるかの判定は、少し横着をして `n * logn` で行うことにした。

※大雑把な見積もりでも、

`logn[i] * n[i]logn[i] = n[i](logn[i])^2 <= n[i](logn)^2 (n = Sum(n[i]))` より、

`Sum(n[i](logn[i])^2) <= (logn)^2 * Sum(n[i]) = n * (logn)^2` なので、間に合いそうと判断。

```go
var q int

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		n := ReadInt()
		P := ReadIntSlice(n)
		for j := 0; j < n; j++ {
			P[j] /= 100
		}
		x, a := ReadInt2()
		y, b := ReadInt2()
		k := ReadInt64()

		sort.Sort(sort.Reverse(sort.IntSlice(P)))
		percents := make([]int, n)
		for j := 0; j < n; j++ {
			if (j+1)%a == 0 && (j+1)%b == 0 {
				percents[j] = x + y
			} else if (j+1)%a == 0 {
				percents[j] = x
			} else if (j+1)%b == 0 {
				percents[j] = y
			}
		}

		isOK := func(m int) bool {
			tmp := make([]int, m)
			for h := 0; h < m; h++ {
				tmp[h] = percents[h]
			}
			sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

			val := int64(0)
			for h := 0; h < m; h++ {
				val += int64(tmp[h]) * int64(P[h])
			}

			if val >= k {
				return true
			}
			return false
		}

		ng, ok := -1, n
		for int(math.Abs(float64(ok-ng))) > 1 {
			mid := (ok + ng) / 2
			if isOK(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}

		if isOK(ok) {
			fmt.Println(ok)
		} else {
			fmt.Println(-1)
		}
	}
}
```

正しい考察にたどり着くまでに結構時間がかかってしまったので、この速度を上げるように努めたいです。

---

<a id="markdown-感想" name="感想"></a>
## 感想

Cとかは別にクエリ形式にする必要ないのでは。。？

社会人として、日曜0時スタートのコンテストが unrated になるのはなかなか心にくるものがありますが、
問題自体は楽しいので引き続き参加していきたいです。
