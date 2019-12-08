<!-- ランレングス符号化で易問をさらに楽にする（緑ぐらいまでの人向け） -->

<!-- TOC -->

- [TL;DR](#tldr)
- [きっかけ](#%e3%81%8d%e3%81%a3%e3%81%8b%e3%81%91)
- [アルゴリズム](#%e3%82%a2%e3%83%ab%e3%82%b4%e3%83%aa%e3%82%ba%e3%83%a0)
- [実装例](#%e5%ae%9f%e8%a3%85%e4%be%8b)
- [活用事例](#%e6%b4%bb%e7%94%a8%e4%ba%8b%e4%be%8b)
	- [ABC143 C. Slimes](#abc143-c-slimes)
	- [Educational Codeforces Round 75 A. Broken Keyboard](#educational-codeforces-round-75-a-broken-keyboard)
	- [Codefources Round 604 A. Beautiful String](#codefources-round-604-a-beautiful-string)
	- [Codefources Round 600 Div.2 A. Single Push](#codefources-round-600-div2-a-single-push)
	- [Codeforces Round 604 C. Beautiful Regional Contest](#codeforces-round-604-c-beautiful-regional-contest)
- [まとめ](#%e3%81%be%e3%81%a8%e3%82%81)

<!-- /TOC -->

<a id="markdown-tldr" name="tldr"></a>
## TL;DR

- AtCoder緑か自分（1300）程度までの人向けです。
- ランレングス符号化を使いやすい形で用意しておくと、意外と高頻度で便利に使えるよという話。
- 10月以降のコンテストの問題で役立ったものを、活用事例として紹介します。

<a id="markdown-きっかけ" name="きっかけ"></a>
## きっかけ

AGC039のA問題が60分ぐらい使ってもWAが取れずに、Bへ向かうも解けず、初めての太陽をやらかしてしまいました。

https://atcoder.jp/contests/agc039/tasks/agc039_a

それなりにレートも落ちてしまっただけに、復習して確実に経験値にしようと、A問題から復習に力を入れました。

人によって好みはわかれるかもしれませんが、このA問題がランレングス符号化を行うと見通しがよくなるもので、
このタイミングでようやくスニペットをこしらえることになりました
（アルゴリズムの存在は知っていても、必要性を感じなかったので用意していなかった）。

すると、以降のコンテストでやたらと使用頻度が高くなり、
無駄に用意しているスニペット集の中でもかなり目立つものになってきました。

なので、自分程度までの人だったら同じように利用できるのではないかと思い、本記事を書くことにしました。

<a id="markdown-アルゴリズム" name="アルゴリズム"></a>
## アルゴリズム

色々とバリエーションはありますが、
本記事では[wikipediaの該当ページ](https://ja.wikipedia.org/wiki/%E9%80%A3%E9%95%B7%E5%9C%A7%E7%B8%AE)で
一番最初に説明されている、最も基本的なものを指しています。

`AAAAABBBBBBBBBAAA` という文字列を `5A9B3A` というように符号化するもので、
「連長圧縮」という日本語の通り直感的なものです。

AtCoderユーザは皆さん優秀なので、アルゴリズムの名前は知らなくても自然と使ってた、という人も多そうです。

メインの焦点は文字列だと思いますが、競技プログラミングでよく登場する整数配列についても、
しばしば便利に使えます。

<a id="markdown-実装例" name="実装例"></a>
## 実装例

言語はGolangですが、アルゴリズム自体はシンプルなので、他言語を使っている方も読む分にはほとんど問題ないかと思います。

ポイントとしては、 **「文字とカウントを分離して別々の配列として出力する」** ことぐらいです
（もっと便利なシグニチャがあったら教えていただきたいです）。

`RunLengthDecoding(RunLengthEncoding(S)) == S` となるように復号関数も一応用意していますが、
特に使ったことはありません。

```go
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

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []rune, L []int) []rune {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []rune{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}
```

C++などだったらジェネリクスを使うのが妥当かと思います。

Golangではジェネリクスがないので、自分の場合、この手の問題ではスニペットで必要な型をその時々に指定しています
（ `interface{}` だといちいち型アサーションするのが面倒）。

https://github.com/my0k/go-competitive/blob/master/snippets/runLengthEncoding.snip

実際のところ、Golangも全然詳しくないので、もっといい方法があれば是非教えていただきたいです。

<a id="markdown-活用事例" name="活用事例"></a>
## 活用事例

ここからは、実際に活用できた問題と具体的な活用方法について書いていきます。

<a id="markdown-abc143-c-slimes" name="abc143-c-slimes"></a>
### ABC143 C. Slimes

[問題のURL](https://atcoder.jp/contests/abc143/tasks/abc143_c)

まさにランレングス符号化そのものです。

スニペット呼び出して、関数に放り込んで、出力のいずれかの配列の長さを提示して終わりです。

他の解き方でも簡単に解けるとは思いますが、バグの心配もなく貼るだけで終わるならばそれが一番だと思います。

```go
var n int
var S []rune

func main() {
	n = ReadInt()
	S = ReadRuneSlice()

	comp, _ := RunLengthEncoding(S)
	fmt.Println(len(comp))
}
```

<a id="markdown-educational-codeforces-round-75-a-broken-keyboard" name="educational-codeforces-round-75-a-broken-keyboard"></a>
### Educational Codeforces Round 75 A. Broken Keyboard

[問題URL](https://codeforces.com/contest/1251/problem/A)

問題を要約すると、おおよそ以下のような感じです。

> 26文字のキーボードについていくつかの壊れたキーがあるので、故障していないキーだけを列挙する問題。
>
> 壊れたキーを叩くと、その文字について確実に2回タイプされてしまう、というのをヒントに解く。

「確実に故障していない」と断定できるのは、連続で奇数回タイプされた文字だけなので、
それを見つけるために文字列をfor文でスキャンしてカウントを取る、という方法でも簡単に解けます。

しかしながら、フルフィードバックではないCodeforcesではとにかくバグの出る余地をなくしたいので、
実績のあるスニペットを持ち出せると多少は安心できます。

```go
var t int
var S []rune

type DirRange []rune

func (a DirRange) Len() int           { return len(a) }
func (a DirRange) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DirRange) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		S = ReadRuneSlice()
		memo := make(map[rune]int)

		SS, L := RunLengthEncoding(S)
		for i := 0; i < len(L); i++ {
			if L[i]%2 == 1 {
				memo[SS[i]] = 1
			}
		}

		answers := DirRange{}
		for key := range memo {
			answers = append(answers, key)
		}

		sort.Sort(answers)
		fmt.Println(string(answers))
	}
}
```

<a id="markdown-codefources-round-604-a-beautiful-string" name="codefources-round-604-a-beautiful-string"></a>
### Codefources Round 604 A. Beautiful String

[問題のURL](https://codeforces.com/contest/1265/problem/A)

問題を要約すると、おおよそ以下のような感じです。

> `a, b, c, ?` の4文字からなる文字列が与えられる。
>
> `?` の部分を `a, b, c` のいずれかに置き換えて、同じ文字が連続しないように文字列を構築せよ、という問題。

この問題も、やること自体は非常に簡単かつすぐに思いつくもので、
「文字列をスキャンして前後の文字と違うものを素直に選び続ければ良い」というだけです。

「最初から2つ以上連続する `a, b, c` が存在するとアウト」という部分を先に処理しておくと、
後半の構築に集中することが出来、バグらせる確率が多少減らせます。

文字列をスキャンして前後を見る、としても良いんですが、
先頭と末尾は処理が変わったり、、とか考えるのも面倒で、ランレングス符号化したものを見ると簡単です。
（結局構築時にそれは考える必要があるんですが）。

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

  // 後は素直に構築するだけ
}
```

<a id="markdown-codefources-round-600-div2-a-single-push" name="codefources-round-600-div2-a-single-push"></a>
### Codefources Round 600 Div.2 A. Single Push

[問題のURL](https://codeforces.com/contest/1253/problem/A)

問題を要約すると、おおよそ以下のような感じです。

> 与えられた配列 `A` に対して、1度だけ任意の連続区間に対してある正の整数加算することが許される。
> 操作は行わなくても良い。
>
> これによって、もう一方の与えられた配列 `B` に等しくすることができるか？

これを達成するためには、
すべての要素に関して `diff[i] = B[i] - A[i]` を計算しておき、
この `diff` 配列が `[0, ..., 0, k, ..., k, 0, ..., 0], k >= 0` のような形になっていればよい、
というのは割と簡単にわかると思います。

しかしながら、これをバグらせずに判定するのは、
いくつかの場合分けが必要だったり、それなりに神経を使うような気がします。

そこで、 **整数配列に対するランレングス符号化** を考えます。

圧縮後の配列に対して、

1. 負の整数が検出されたらアウト
2. 正の整数が2つ以上検出されたらアウト
3. そうでないならセーフ

のように判定すればよい、というふうな実装方針が立ちます。

前後の `0` 埋めがないケースの場合分けなどが不要となり、それなりに実装の負荷は小さくなっているのではないでしょうか。

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
```

<a id="markdown-codeforces-round-604-c-beautiful-regional-contest" name="codeforces-round-604-c-beautiful-regional-contest"></a>
### Codeforces Round 604 C. Beautiful Regional Contest

[問題のURL](https://codeforces.com/contest/1265/problem/C)

問題を要約すると、おおよそ以下のような感じです。

> あるプログラミングコンテストにおいて、 `n` 人の解答問題数が配列で与えられるので、
> それぞれの人に金・銀・銅メダルを授与することを考える。
>
> 以下の条件を満たすように、それぞれのメダルの数を求める問題。
>
> 1. それぞれのメダルの数は正の整数である。
> 2. 金メダルの数は、銀・銅メダルよりも大きい必要がある。ただし、銀メダルと銅メダルの数量関係は不問である。
> 3. 金メダルを受賞する参加者は、銀メダルを受賞する参加者よりも多くの問題を解いていなければならない。
> 4. 銀メダルを...、銅メダルを...。
> 5. 銅メダルを...、メダルを受賞していない参加者よりも多くの問題を解いていなければならない。
> 6. 3つのメダルの合計個数は、 `Floor(n/2)` 以下である必要がある。

考えるポイントが多いだけに、ここでも前処理として整数配列に対するランレングス符号化を行っておくと、
問題の本質部分に集中できて楽です。

条件3, 4, 5を踏まえると、メダルの種類の境界において、
問題の正解数に1以上の差がある必要がある、すなわち、正解数が異なっている必要があるとわかります。

なので、各メダルの枚数を加算するにあたっては、ランレングス符号化を行った後の長さが使えます。

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
```

<a id="markdown-まとめ" name="まとめ"></a>
## まとめ

ランレングス圧縮を用意しておくと、簡単な問題がさらに楽になったり、
多かれ少なかれバグの心配が減ったりと、結構便利だよ、ということを伝えるための内容でした。

正直見返してみると、若干過剰適合というか、「金槌を手にして多くの問題が釘に見えている」という状態に落ちているような気もします。

とはいえ、それで損した感じはなく、少なくとも毒にはなっていないだろうということで、今回このように紹介してみました。

また、問題例のほとんどがCodeforcesである通り、最近こどふぉにも積極的に参加していますので、
近いレベル帯のフレンド（ライバル）募集中です
（知っている人がみんな紫以上でDiv1に行っちゃったり、Div2一緒に出ても大差つけられてしまって悲しい）。

