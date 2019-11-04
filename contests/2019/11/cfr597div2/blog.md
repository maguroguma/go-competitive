最近関数型プログラミングの勉強ばかりしていて、競技の方が疎かになっていました。

ところどころ脳が停止していたり、不要なWA出してしまったりは避けられなかったのかなと。

※Dは自分にとって貴重なMSTの典型問題なので、どこかで復習し次第追記いたします。

<!-- TOC -->

- [A. Good ol' Numbers Coloring](#a-good-ol-numbers-coloring)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Restricted RPS](#b-restricted-rps)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Constanze's Machine](#c-constanzes-machine)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-good-ol-numbers-coloring" name="a-good-ol-numbers-coloring"></a>
## A. Good ol' Numbers Coloring

[問題URL](https://codeforces.com/contest/1245/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

非負整数に対して黒か白か色を付与していく。

付与の仕方は、

1. 0は必ず白
2. `i-a`, `i-b` が負数なら `i` は黒
3. `i-a` が白なら `i` は白
4. `i-b` が白なら `i` は白
5. いずれでもないなら `i` は黒

という風にすべての非負整数を塗るとき、黒の数は無限か有限かを判定する問題。

<a id="markdown-解答" name="解答"></a>
### 解答

全然頭が働かず、なんとなく `a, b` が互いに素ならいずれ全部白くなりそう、と直感的に判断し無証明でsubmitしてしまいました。

```go
var t int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		a, b := ReadInt2()

		if Gcd(a, b) == 1 {
			fmt.Println("Finite")
		} else {
			fmt.Println("Infinite")
		}
	}
}
```

これでは良くないので、公式editorialの内容を自分なりに噛み砕いて和訳したものを置いておきます。

`a, b` が互いに素な場合の証明の要約は、
「ある程度大きな数に関して `a` 飛びの小さい数を見ると、その中に `b` の倍数が必ず1つは含まれている」
というものです。



コンテスト中にここまできっちり証明するのは流石にナンセンスなはずで、
類題をたくさん解いて早く確信に至ることができる、というのがあるべき姿な気がします。

<a id="markdown-b-restricted-rps" name="b-restricted-rps"></a>
## B. Restricted RPS

[問題URL](https://codeforces.com/contest/1245/problem/A)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

[限定じゃんけん](https://ja.wikipedia.org/wiki/%E8%B3%AD%E5%8D%9A%E9%BB%99%E7%A4%BA%E9%8C%B2%E3%82%AB%E3%82%A4%E3%82%B8#%E7%AC%AC1%E7%AB%A0%E3%80%8C%E5%B8%8C%E6%9C%9B%E3%81%AE%E8%88%B9%E3%80%8D)で
相手の出す手の順番がわかっているので、勝負回数の過半数に勝利できるようにする。
勝利できるのであれば、その時の手順を構築して出力する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

貪欲に相手の手に対して勝てる手を出していく方針で良い。

事前に相手の手をそれぞれカウントしておいて全体で勝利できるかを判断する。

手の構築については、貪欲に勝利手を割り振る部分と、余った手を適当に割り振る部分に分けるのが楽だと思う。

```go
var t int
var n int
var a, b, c int
var S []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		a, b, c = ReadInt3()
		S = ReadRuneSlice()

		need := (n + (2 - 1)) / 2

		r, p, s := 0, 0, 0
		for i := 0; i < len(S); i++ {
			if S[i] == 'R' {
				r++
			} else if S[i] == 'P' {
				p++
			} else {
				s++
			}
		}

		wins := 0
		wins += Min(a, s)
		wins += Min(b, r)
		wins += Min(c, p)

		if wins >= need {
			fmt.Println("YES")

			answers := make([]rune, n)
			for i := 0; i < n; i++ {
				answers[i] = 'x'
			}

			// 勝てるやつを割り振る
			for i := 0; i < len(S); i++ {
				if S[i] == 'R' && b > 0 {
					answers[i] = 'P'
					b--
				} else if S[i] == 'P' && c > 0 {
					answers[i] = 'S'
					c--
				} else if S[i] == 'S' && a > 0 {
					answers[i] = 'R'
					a--
				}
			}

			// あまりを割り振る
			for i := 0; i < n; i++ {
				if answers[i] == 'x' {
					if a > 0 {
						answers[i] = 'R'
						a--
					} else if b > 0 {
						answers[i] = 'P'
						b--
					} else if c > 0 {
						answers[i] = 'S'
						c--
					}
				}
			}

			fmt.Println(string(answers))
		} else {
			fmt.Println("NO")
		}
	}
}
```

20分は時間かかり過ぎだし、条件式で横着したせいでWAするし、ひどかった。

<a id="markdown-c-constanzes-machine" name="c-constanzes-machine"></a>
## C. Constanze's Machine

[問題URL](https://codeforces.com/contest/1245/problem/C)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

音声入力によって得られた文字列が与えられるが、
機器が細工されてしまったため `w` は `uu` と、 `m` は `nn` と入力されるようになってしまっている。

本来入力されたと考えられる文字列の通り数を `10^9 + 7` で割ったあまりを答える問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

`i` 文字目まで見たときの通り数がわかっていれば、それをヒントに `i+1` 文字目まで見たときの通り数もわかりそう、
ということでDPを考える。

問題のポイントとして `uuu` や `nnn` については `uw, wu` や `nm, mn` のように、連結させる方法によって複数の解釈が可能なところなので、
これをフラグ管理して考えた。

すなわち、

`dp[i+1][0]: i文字目までみて直前の文字が連結されてできたものではない（直前の文字がw, mではない）場合の数`

`dp[i+1][1]: i文字目までみて直前の文字が連結されてできたもの（直前の文字がw, mのいずれか）の場合の数`

とDPテーブルを定義する。

すると、 `dp[i+1][0]` については、現在注目している `i` 番目の文字がいかなる文字であったとしても、考慮される必要がある
（注目中の文字が `u, n` のいずれかであっても、それらをそのまま使うという選択肢が取れる）ため、
`dp[i+1][0] += dp[i][0] + dp[i][1]` という遷移が行われる。

また、直前の `i-1` 番目の文字と注目中の `i` 番目の文字がともに `u, n` である場合は、連結を考慮する必要がある。
連結が可能なのは、直前の `i-1` 番目の文字が連結に使われずに `u, n` として単体で残っていることなので、
`dp[i+1][1] += dp[i][0]` という遷移でよい。

問題の設定上、 `w, m` のいずれかが登場したら0通りとすることに注意。

また、最後に答えを出力するときに `dp[n][0] + dp[n][1]` と直接出力してしまわないように注意（1敗、本コンテストで一番反省すべきところ）。

```go
var S []rune
var dp [100000 + 5][2]int64

func main() {
	S = ReadRuneSlice()
	n := len(S)

	for i := 0; i < n; i++ {
		if S[i] == 'm' || S[i] == 'w' {
			fmt.Println(0)
			return
		}
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		dp[i+1][0] += dp[i][0] + dp[i][1]
		dp[i+1][0] %= MOD

		if i == 0 {
			continue
		}

		if S[i] == 'u' && S[i-1] == 'u' {
			dp[i+1][1] += dp[i][0]
			dp[i+1][1] %= MOD
		} else if S[i] == 'n' && S[i-1] == 'n' {
			dp[i+1][1] += dp[i][0]
			dp[i+1][1] %= MOD
		}
	}

	ans := dp[n][0] + dp[n][1]
	ans %= MOD
	fmt.Println(ans)
}
```

editorialを見ると、たしかに連結を考慮するパターンは前2つを見る形でよくて、
全体的にスマートですね。

---

競技プログラミングもたくさんやりたいが、最近業務開発指向の学習も楽しくて時間の使い方が悩ましい。
