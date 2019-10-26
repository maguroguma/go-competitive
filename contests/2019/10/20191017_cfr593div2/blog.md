
Cの構築難しい。。と思ってたらみんなはやすやすと通していて驚きました。

構築は簡単なものでも刺さらないとずっと解けないので、筋の良い考え方のパターンをためていきたいところ。

※Dは実装方法が参考になりそうなので、取り組み次第追記します。

<!-- TOC -->

- [A. Stones](#a-stones)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Alice and the List of Presents](#b-alice-and-the-list-of-presents)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Labs](#c-labs)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
  - [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-stones" name="a-stones"></a>
## A. Stones

[問題URL](https://codeforces.com/contest/1236/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

3つの石が積まれた山から、2通りの決まった石の取り除き方を実行したとき、最大で何個取れるか？という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

2番目の方法でできる限りたくさん取って、取れなくなったら今度は1番目の方法でできる限りたくさん取る、でOK。

以下は本番中に考えた直感的な証明。

> それぞれ山をA, B, Cとすると、Cから石を取得する方法は2番目の「Bから1つCから2つ」という方法しかない。
> もう片方の石のとり方は「Aから1つBから2つ」というものなので、Bの山の減り方に着目すると、先に2番目の方法で取り尽くしたほうが良い。

制約が小さいので、シミュレーションで書きました。

```go
var t int

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		a, b, c := ReadInt3()

		ans := 0
		for b >= 1 && c >= 2 {
			b--
			c -= 2
			ans += 3
		}
		for a >= 1 && b >= 2 {
			a--
			b -= 2
			ans += 3
		}
		fmt.Println(ans)
	}
}
```

<a id="markdown-b-alice-and-the-list-of-presents" name="b-alice-and-the-list-of-presents"></a>
## B. Alice and the List of Presents

[問題URL](https://codeforces.com/contest/1236/problem/B)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

`n` 種類のプレゼントを `m` 人の子どもたちに振り分ける方法は何通りか？という問題。

ただし、ルールが2つあって、

1. ある種類のプレゼントについて、ある1人に割り振る数は0個か1個のいずれか。
2. ある種類のプレゼントについて、必ず最低1人には1個割り振る必要がある。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

ある1種類のプレゼントについて考えたとき、の `m` 人の子どもたちへの割り当て方を、
`m` 人の子供をビット列として見立てるとスッキリする。

ビットが立っているときにプレゼントを1個割り振ると考えると、
`1 ~ (2^m-1)` の割り振り方がある（2の制約から、すべてのビットが0の場合はNG）。

よって、あるプレゼントについては割り振り方が `(2^m - 1)` 通りあり、
どのプレゼントについてもこれは同じなので、
積事象を考えて `(2^m - 1) ^ n` が答え。

`modpow` をスニペットにしているので、それを呼び出して終わり。。

※64bit整数で計算しないとオーバーフローでWAするので注意しましょう（1敗）。

```go
const MOD = 1000000000 + 7

var n, m int64

func main() {
	n, m = ReadInt64_2()

	tmp := modpow(2, m, MOD)
	tmp = NegativeMod(tmp-1, MOD)
	ans := modpow(tmp, n, MOD)
	ans %= MOD
	fmt.Println(ans)
}

// ModInv returns $a^{-1} mod m$ by Fermat's little theorem.
// O(1), but C is nearly equal to 30 (when m is 1000000000+7).
func ModInv(a, m int64) int64 {
	return modpow(a, m-2, m)
}

func modpow(a, e, m int64) int64 {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := modpow(a, halfE, m)
		return half * half % m
	}

	return a * modpow(a, e-1, m) % m
}

// NegativeMod can calculate a right residual whether value is positive or negative.
func NegativeMod(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}
```

<a id="markdown-c-labs" name="c-labs"></a>
## C. Labs

[問題URL](https://codeforces.com/contest/1236/problem/C)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

`n^2` 個のラボを `n` 個のラボからなるグループ `n` 個に分割したとき、
問題で定義された6パターンの関数値の最小値が最大となるような分割方法を答える問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

※後半かなり雑です。

あるグループA, B間の関数値を考えたとき、 `n * n` 通りの大小関係の比較が起こることから、
水の流れる総量は `n^2` となる。

なので `f(A, B) = x` のとき `f(B, A) = n^2 - x` となる。

よって、最大化した最小値は `floor(n^2 / 2)` としかならないはず。
このような構築方法は無いかと考える。

。。ここまではコンテスト中かなり時間がかかったものの、気づけて、
そこからは無証明の推測で解きました。

`n = 4` の場合を考えると、もとの行列で縦のグループを観たとき、
左から1番目のグループと2番目のグループを比較すると、目的の値に対して不足分があります。

足りないのは `floor(n / 2)` で、上 `floor(n / 2)` 行を反転させれば辻褄があう、と考えました。



一応、コンテス中は `n = 5` のケースも考えてあっていることを確認し提出しました。

実際には、行ごとに考えたときに矢印の流れが均等になっている（半分の行では矢印が右向きに流れて、もう半分は左向きに流れる）ことに気づくことができれば確信を持てそうです。
下の行から上の行へは矢印は伸びず、逆を考えた場合は必ず矢印が伸びるので、行間では特にグループ間の差を考える必要はありません。

```go
var n int
var answers [][]int

func main() {
	n = ReadInt()

	r := n / 2
	answers = make([][]int, n)
	for i := 0; i < n; i++ {
		answers[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			answers[i][j] = i*n + (j + 1)
		}
	}
	for i := 0; i < r; i++ {
		answers[i] = Reverse(answers[i])
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Printf("%d\n", answers[i][j])
			} else {
				fmt.Printf("%d ", answers[i][j])
			}
		}
	}
}

func Reverse(A []int) []int {
	res := []int{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}
```

あくまでもコンテスト中の思考整理のために書いた方法なので、
模範解答はEditorialのもとの行列をジグザグにスキャンして各グループに配置する方法がベストかと思います。

---

「ジグザグに見る」っていうのをパターン化するのは賢くはないが、対症療法にはなりえるかも？
