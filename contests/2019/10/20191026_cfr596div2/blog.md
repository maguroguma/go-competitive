Cはちょっと今の自分には難しかった気がするので、せめてBをスムーズに通したかった。

AtCoderの水色帯の人も（詰めが甘くなってしまいsystem testで落としてしまった人は多そうですが、）
本質的な部分は捉えられていてすごいなぁと思いました。

<!-- TOC -->

- [A. Forgetting Things](#a-forgetting-things)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. TV Subscriptions](#b-tv-subscriptions)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. p-binary](#c-p-binary)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-forgetting-things" name="a-forgetting-things"></a>
## A. Forgetting Things

[問題URL](https://codeforces.com/contest/1247/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

`a + 1 = b` という数式に関して、10進数表記の最上位桁のみがわかっているので、
適当な `a, b` を答えよ、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

最上位桁の桁上りが起こるか起こらないか、をベースに考えれば良い。

桁上りが起こる場合は `a = x9` などとすればよく、桁上りが起こらない場合は `a = x0` などとすればよい。
`a` が決まれば `b` も自動的に決まる。

`b` の最上位桁が `a` よりも小さかったり、最上位桁が `b` のほうが2以上大きかったりしたら、impossibleとする。

ただし、9からの桁上りのみ例外的に扱わなければならないので注意。

※pretestが優しくて教えてくれましたが、これは入っていなくても文句が言えない初歩的なものな気がするので、要反省案件ですね。。

```go
var da, db int

func main() {
	da, db = ReadInt2()

	if da == 9 && db == 1 {
		fmt.Println(9, 10)
		return
	}

	if da > db {
		fmt.Println(-1)
	} else if da == db {
		a := da * 10
		b := a + 1
		fmt.Println(a, b)
	} else {
		if db-da == 1 {
			a := da*10 + 9
			b := a + 1
			fmt.Println(a, b)
		} else {
			fmt.Println(-1)
		}
	}
}
```

<a id="markdown-b-tv-subscriptions" name="b-tv-subscriptions"></a>
## B. TV Subscriptions

[問題URL](https://codeforces.com/contest/1247/problem/B)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

`n` 日分のTV番組情報が与えられるので、ある数のショーのみ購読し、
連続 `d` 日間TVを視聴できるようにしたい。

できるだけ少ない数のショーのみを購読しようとした場合、どれだけ購読すればよいか、という問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

easyバージョンは、番組情報配列を素直に `d` 個分切り出して、都度その中に含まれるショーの種類数をカウントすれば良い。

hardバージョンは制約的にそれは許されないため、スライディングウィンドウを考える。

すなわち、毎回注目区間に含まれるショーの種類数を数え直すのではなく、
区間を1つずらした分の差分のみを着目し、端点のみ観て種類数の情報を更新すれば良い。

```go
var t int
var n, k, d int
var A []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n, k, d = ReadInt3()
		A = ReadIntSlice(n)

		subsc := [1000000 + 5]int{} /* ここ危ない！ */
		ans := 0
		for i := 0; i < d; i++ {
			ch := A[i]
			if subsc[ch] == 0 {
				ans++
			}
			subsc[ch]++
		}

		l := 0
		tmpAns := ans
		for r := d; r < n; r++ {
			subsc[A[l]]--
			if subsc[A[l]] == 0 {
				tmpAns--
			}
			l++

			if subsc[A[r]] == 0 {
				tmpAns++
			}
			subsc[A[r]]++

			ChMin(&ans, tmpAns)
		}

		fmt.Println(ans)
	}
}
```

コンテスト中に提出したコードは、各テストケースごとに巨大な固定長配列を取得し直しているため、
1900msec程度かかってしまい、危ないところでした。
固定長配列の部分を `make(map[int]int)` と辞書にするだけで150msec程度になりました。
AtCoderではなかなか気にしない部分だったので、これからはメモリ確保にも気を使いたいと思います。

また、尺取法が苦手すぎて、なにか複雑なことをやらなければいけない、と思いこんでしまい、徒に時間を消耗してしまいました。

今回のはウィンドウが固定長で尺取虫の動きをしていないので、しゃくとり法と呼ぶのは良くない気がします。

<a id="markdown-c-p-binary" name="c-p-binary"></a>
## C. p-binary

[問題URL](https://codeforces.com/contest/1247/problem/C)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

`2^x + p (x >= 0, -1000 <= p <= 1000)` で表される数をp-binaryと定義したとき、
ある任意の正整数 `n` を重複ありのp-binaryの和で表したい。

項の数をできるだけ小さくなるように表したとき、その数はいくらになるか、という問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

※[すぬけさんの解説文](https://scrapbox.io/snuke/Codeforces_Round_%23596)を読み解いて考えた内容です。

p-binary `x` 個を用いて `n` を和で表現するとき、 `n = n' + x*p (n' は重複ありの2べきの数)` と表すことができる。

式を変形して `n - x*p = n'` とすると、 **重複を許した2べきの数をちょうど `x` 個使い、その和で `n - x*p` を作れるか？** という問題に読み替えることができる。
よって、 `n - x*p` が0以下となる場合は、そのような `x (>= 1)` は答えとして不適である。

ここで `popcount(n - x*p)` を考える。
当然ながら、この数は `n - x*p` を2べきの和で表すときの項数の最小値である。

また、 `n - x*p` の各bitを分解することで、 `n - x*p` は様々な2べきの和で表すことができる。

例えば `2^3` のbitが立っている場合、 `2^3 = 8 * 2^0` であり、これらを足し合わせることで項数を `1 ~ 2^3` まで自由に変化させることができる。
これは、任意の立っているbitに対しても同じように分解が可能である。
よって、 `n - x*p` を重複を許して2べきの数の和で表すとき、その項数は `popcount(n - x*p)` から `n - x*p` まで自由に動かすことができる。

以上より、ある `x` を決め打つとき、 `popcount(n - x*p) <= x <= n - x*p` を満たしていれば、その `x` は答えの候補となる。
また、 `popcount(n - x*p)` は大きくても高々30程度であるため、 `x` を小さいところから調べていけばすぐに停止する。

```go
var n, p int
var bits [40]int

func main() {
	n, p = ReadInt2()

	for x := 0; ; x++ {
		sum := n - p*x

		if sum <= 0 {
			fmt.Println(-1)
			return
		}

		if PopCount(sum) <= x && x <= sum {
			fmt.Println(x)
			return
		}
	}
}
```

---

Codeforcesの数学はAtCoderよりも苦手な気がします。

