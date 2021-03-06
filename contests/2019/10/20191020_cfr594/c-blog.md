<!-- Codeforces Round No.594 Div.2 C復習 -->

以前コンテストに参加して解けなかったものの復習です。

公式Editorialがハイコンテクスト過ぎてよくわからなかったのと、
数え上げの方法の典型度合いがものすごく高い気がしたので、別記事として書きました。

## C. Ivan the Fool and the Probability Theory

[問題のURL](https://codeforces.com/contest/1248/problem/C)

### 問題

`n` 行 `m` 列からなるグリッドを、黒・白の2色で塗り分けることを考える。

ただし、塗り方は「辺に関して隣接しているグリッドについて、同じ色の隣接グリッドがたかだか1つまで」とする。

このような塗り方は何通り存在するか、 `10^9 + 7` で割ったあまりで答えよ。

### 解答

一見複雑だが、「左上から少数のグリッドの色を確定させると、他のグリッドの色は自動的に決まる」というようなことはよくある（と思う）。

とりあえず、1行目を条件に反しないように適当に塗ってみると（同じ色が連続するのは2回までとする）、
以下の図のように、同じ色が連続してしまうと、その列を中心として次の行は一意に定まってしまう。



よって、このようなバターンの総数を数え上げる必要がある（※）。

また、同じ色が連続せずに、黒・白が交互に塗られる場合、
以下の図のように、次の行の1列目の色を決めた時点で、他の列の色が確定してしまう。



よって、このような場合も結局、1列目の色の塗り方の総数を数え上げる必要があり、
これは、（※）の部分と同じようにして数え上げることができる。

黒・白が2つまで連続して良いという条件のもとでの、1次元グリッドの塗り分けの総数を求めることに集中できる。

とりあえず、最初の色を黒として樹形図を書いてみる。
以下の図に示すように、よくよく観察すると、
色に関係なく2つ前と1つ前の通り数の和となるように遷移する（フィボナッチ数列）ため、
`dp[i-2] + dp[i-1] = dp[i]` のようにまとめられる。

以上をもとにコーディングしていく。

同じ色が隣接することがない塗り方1通りを2重に数えない、
左上を白とした場合も数えるために2倍する、あたりを忘れないように注意する。

```go
var n, m int64
var dp [100000 + 5]int64

func main() {
	n, m = ReadInt64_2()

	dp[0], dp[1] = 1, 2
	for i := 2; i <= 100000+3; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= MOD
	}

	var ans int64
	ans = dp[m-1] - 1
	ans += dp[n-1]
	ans %= MOD
	ans *= 2
	ans %= MOD

	fmt.Println(ans)
}
```

上記のフィボナッチ数を考えるあたりって自明なんでしょうか？
自分にはあまり自明には感じられないので、説明がかなり冗長かもしれません。

