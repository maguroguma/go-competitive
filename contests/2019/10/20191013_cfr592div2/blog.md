Cが少しでもいけそうと思ってしまった時点で負けでした。

※Bは通ったものの、正当性の保証が全くできていないのでコードのみ記載します。

※Cは後日追記します（2019-10-14追記）。

<!-- TOC -->

- [A: Pens and Pencils](#a-pens-and-pencils)
  - [問題](#%e5%95%8f%e9%a1%8c)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Rooms and Staircases](#b-rooms-and-staircases)
  - [問題](#%e5%95%8f%e9%a1%8c-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. The Football Season](#c-the-football-season)
  - [問題](#%e5%95%8f%e9%a1%8c-2)
  - [解答（2019-10-14追記）](#%e8%a7%a3%e7%ad%942019-10-14%e8%bf%bd%e8%a8%98)
- [D. Paint the Tree](#d-paint-the-tree)
  - [問題](#%e5%95%8f%e9%a1%8c-3)
  - [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-pens-and-pencils" name="a-pens-and-pencils"></a>
## A: Pens and Pencils

[問題URL](https://codeforces.com/contest/1244/problem/A)

<a id="markdown-問題" name="問題"></a>
### 問題

※あまりにも問題文が冗長なので、要点のみ。

`a` 個の講義と `b` 個の実践演習がある。
1本のペンで `c` 個の講義のノートを取る事ができ、1本の鉛筆で `d` 個の実践演習の設計図を描ける。
全部の講義と実践演習をこなしたいが、筆箱の容量はペンと鉛筆合わせて `k` 本までしか入らない。

すべての講義と実践演習をこなせるか判定し、こなせる場合は筆箱に収まるペンと鉛筆の本数をそれぞれ答えよ。

テストケースは `t` 個あるため、それぞれについて答えよ。

制約:

- `1 <= t <= 100`
- `1 <= a, b, c, d, k <= 100`

<a id="markdown-解答" name="解答"></a>
### 解答

答えはなんでも良いので最小化する必要はないが、最小限必要な本数をそれぞれ答えるほうが簡単なので、それを求める。

`a` 個の講義をこなすために必要なペンの本数と、 `b` 個の実践演習をこなすために必要な鉛筆の本数は、
それぞれ `Ceil(a/c), Ceil(b/d)` で求まる。

それの合計が `k` を超えていなければ、それぞれ出力すれば良い。

```go
var t int

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		a, b, c, d := ReadInt4()
		k := ReadInt()

		pen := (a + (c - 1)) / c
		pencil := (b + (d - 1)) / d
		if pen+pencil > k {
			fmt.Println(-1)
		} else {
			fmt.Println(pen, pencil)
		}
	}
}
```

オーバーフローとゼロ除算に注意するぐらい。

<a id="markdown-b-rooms-and-staircases" name="b-rooms-and-staircases"></a>
## B. Rooms and Staircases

[問題URL](https://codeforces.com/contest/1244/problem/B)

<a id="markdown-問題-1" name="問題-1"></a>
### 問題

※問題URLの図がわかりやすいため、ここでは要点のみ。

Nikolayは2階建ての家に住んでいる。家のそれぞれのフロアには `n` 個の部屋があり、一列に並んでいる。
左から `1 ~ n` の番号が振られている。

同じフロアの隣り合っている部屋はすべて移動可能である。
また、いくつかの部屋には1階と2階をつなぐ階段がある場合があり、その場合は、各階同じ番号の部屋を行き来できる。

Nikolayは以下の2つのルールを守りながら、できるだけ多くの部屋を回りたい。
回れる部屋の最大個数を答えよ。

- スタートする部屋は好きに選ぶことができる。
- 一度通った部屋は二度と通ることはできない。

`t` 個のテストケースに対して答えよ。

制約:

- `1 <= t <= 100`
- `1 <= n <= 1000`

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

※とりあえずはコードのみ。

いずれかのフロアの端の部屋から出発したほうがよく、出発地点からできるだけ遠くの階段のある部屋まで移動する。
階段を使ったあとは、また別フロアの出発地点の部屋番号の部屋まで戻る。

複数の出発地点のうち、回ることができた部屋数が大きい方を出力する。

```go
var t int

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		n := ReadInt()
		S := ReadRuneSlice()

		allZero := true
		for i := 0; i < n; i++ {
			if S[i] == '1' {
				allZero = false
				break
			}
		}
		if allZero {
			fmt.Println(n)
			continue
		}

		if S[0] == '1' || S[n-1] == '1' {
			fmt.Println(2 * n)
			continue
		}

		minIdx, maxIdx := 0, n-1
		for j := 0; j < n; j++ {
			if S[j] == '1' {
				minIdx = j
				break
			}
		}
		for j := n - 1; j >= 0; j-- {
			if S[j] == '1' {
				maxIdx = j
				break
			}
		}

		res := Max(2*(maxIdx+1), 2*(n-(minIdx)))
		fmt.Println(res)
	}
}
```

正当性はどうやって保証するのだろう。。？

<a id="markdown-c-the-football-season" name="c-the-football-season"></a>
## C. The Football Season

[問題URL](https://codeforces.com/contest/1244/problem/C)

<a id="markdown-問題-2" name="問題-2"></a>
### 問題

Berlandのフットボールのシーズンが終わった。
Berlandフットボールのルールによると、それぞれの試合は2つのチームで争われる。
試合結果は、引き分けかどちらか片方のチームの勝利・敗北のいずれかである。
チームが勝つと勝利チームに `w` ポイントが入り、負けると敗北チームは点数を獲得できない。
また、引き分けの場合は両チームに `d` ポイントが入る。

マネージャは結果をまとめようとしたが、詳細を紛失してしまい、
残されているのは `n` 試合で `p` ポイント獲得した、という情報のみだった。

これらの情報から `x, y, z` 、すなわち勝利数、敗北数、引き分け数を求めて出力せよ。
答えが存在しない場合は `-1` を出力せよ。

制約: `1 <= n <= 10^12, 0 <= p <= 10^17, 1 <= d < w <= 10^5`

<a id="markdown-解答2019-10-14追記" name="解答2019-10-14追記"></a>
### 解答（2019-10-14追記）

制約が大きすぎるため、愚直に全探索するとTLEしてしまう。

以下の2つを抑えると、全探索の範囲を大幅に小さくできる。

1. `w > d` のため、 `x` をできるだけ大きくした方が良い、すなわち、 `y` は小さい値から探索すれば良い。
2. 1と併せて考えると、 `y < w` までの探索で打ち切って良い。

1はほとんど自明だが、2は以下のようにして考えられる。

> `y = m*w + y' (m >= 0, 0 <= y' < w)` とすると、 `y` 試合引き分けとなったことによって取得できる点数は `d*y = d*m*w + y'*d` である。
> ここで、得点 `d*m*w` は `d*m` 試合勝った場合の得点と同じである。
> 1より、 `y` 試合引き分けと考えるよりも `y'` 試合引き分けと考えたほうが良いため、 `0 <= y < w` で十分。

`y` が決まれば `x, z` も与えられた方程式から順に決まるため、それぞれ0以上の整数であればそれを出力すれば良い。。

```go
var n, p, w, d int64

func main() {
	n, p, w, d = ReadInt64_4()

	for y := int64(0); y < w; y++ {
		if (p-d*y)%w == 0 {
			x := (p - d*y) / w
			z := n - x - y

			if x >= 0 && y >= 0 && z >= 0 {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println(-1)
}
```

Codeforcesの本コンテストのブログのやりとりを参考にしました。
（多分、整数問題たくさん解いてきた人には典型なんだろうなぁ。）

本番中は拡張ユークリッドの互除法に固執してしまい、全く手が出ませんでした。

さっさと諦めるのがベストだったので、愚直に問題と向き合うだけではなく、選球眼も磨いていきたいところです。

<a id="markdown-d-paint-the-tree" name="d-paint-the-tree"></a>
## D. Paint the Tree

[問題URL](https://codeforces.com/contest/1244/problem/D)

<a id="markdown-問題-3" name="問題-3"></a>
### 問題

`n` 頂点からなる木が与えられる。
木は、無向で連結で閉路のないグラフである。

木の各頂点を、3つの色のうち1つを選んで塗らなければならない。
各頂点について、いずれかの色で塗る場合のコストがわかっているとする。

色を塗る際は、以下のルールに従う必要がある。

> 木の任意の3つの頂点として、 `(x, y, z), x != y, y != z, z != x, xとyは1つの辺でつながっている, yとzは1つの辺でつながっている`
> というものを考えたとき、
> `x, y, z` はそれぞれ違う色で塗られなければならない。

このようなルールのもとですべての頂点を塗るときの、最小コストを求めよ。
また、塗り方が存在しない場合は `-1` を出力せよ。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

一見、色々なことをやらなければいけないように見えるが、
以下の図に示すように、ある頂点の次数が3以上になるともうアウトになってしまうことがわかる。

<figure class="figure-image figure-image-fotolife" title="次数3以上の頂点が存在するとうまく塗れない">[f:id:maguroguma:20191014002327j:plain]<figcaption>次数3以上の頂点が存在するとうまく塗れない</figcaption></figure>

よって、すべてのノードの次数は2以下でなければならず、ノードが一列に並んだような木しか考えなくて良い。

次に、その塗り方も単純で、次数が1のノードから塗り始めて、つながっているまだ塗っていないノードを順次塗っていくと、
最初の2つの色を決めた時点で、以降は自動的に塗るべき色が定まることがわかる。

<figure class="figure-image figure-image-fotolife" title="次数1のノードから塗り方を決めていく">[f:id:maguroguma:20191014002025j:plain]<figcaption>次数1のノードから塗り方を決めていく</figcaption></figure>

~~図にも示したとおり、塗り方はせいぜい12通りしかないため、その~~ 塗り方をすべて試して最小の塗り方はどれか、を調べれば良い。

訂正:
[くる](https://profile.hatena.ne.jp/ningenMe/)さんより訂正いただきました。
順方向の6通りだけで十分網羅できています（2つの連続ノードの塗り方は、2色の選び方が `Comb(3, 2)` 通りでその2色の並べ方が `2!` 通りのため積で6通り、そして2つ決まればすべてが決まるので6通りで十分）。

実装部分が本体だと思うので、コードについても補足。

次数を調べたりする部分は素直にやればよいが、色の塗り方の持ち方などは色々な方針がある。
なるべくバグらせないように気をつけながら、実装方針として以下のように考えた。

1. 毎回根から探索するのは大変なので、ある根から見た頂点のつながっている順番を配列で持っておく（ `orders` ）。
2. 12通りの塗り方について、各ノードの色をすべて保存しておく（ `howto` ）。
3. 各塗り方については前処理に従ってコストの計算のみに集中する。

```go
var n int
var C [4][]int
var edges [100000 + 5][]int
var degins [100000 + 5]int
var howto [20][100000 + 5]int
var orders [100000 + 5]int

func main() {
	n = ReadInt()
	C[1] = ReadIntSlice(n)
	C[2] = ReadIntSlice(n)
	C[3] = ReadIntSlice(n)
	for i := 0; i < n-1; i++ {
		u, v := ReadInt2()
		u--
		v--
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	for i := 0; i < n; i++ {
		for _, e := range edges[i] {
			degins[e]++
		}
	}

	for i := 0; i < n; i++ {
		if degins[i] > 2 {
			fmt.Println(-1)
			return
		}
	}

	roots := []int{}
	for i := 0; i < n; i++ {
		if degins[i] == 1 {
			roots = append(roots, i)
		}
	}

	dfs(roots[0], -1, 0)

	temp := [][]int{
		[]int{1, 2, 3}, []int{1, 3, 2},
		[]int{2, 1, 3}, []int{2, 3, 1},
		[]int{3, 1, 2}, []int{3, 2, 1},
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < n; j++ {
			howto[i][orders[j]] = temp[i][j%3]
		}
	}
	for i := 6; i < 12; i++ {
		for j := 0; j < n; j++ {
			howto[i][orders[n-1-j]] = temp[i-6][j%3]
		}
	}

	ans := int64(1 << 60)
	ansOrderIdx := -1
	for i := 0; i < 12; i++ {
		tmpAns := int64(0)

		for j := 0; j < n; j++ {
			c := howto[i][j]
			tmpAns += int64(C[c][j])
		}

		if ans >= tmpAns {
			ans = tmpAns
			ansOrderIdx = i
		}
	}

	fmt.Println(ans)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Println(howto[ansOrderIdx][i])
		} else {
			fmt.Printf("%d ", howto[ansOrderIdx][i])
		}
	}
}

func dfs(cid, pid, curIdx int) {
	orders[curIdx] = cid
	for _, e := range edges[cid] {
		if e == pid {
			continue
		}
		dfs(e, cid, curIdx+1)
	}
}
```

実装ゲーでした。
強い人達は「冗長、虚無」といった感想だったようですが、
自分レベルだと実装面ではできるだけバグらせにくい実装を考える労力が必要で、
やる価値はあるかと思います。

---

個人的な体感難易度は `A < D < B << C` でした。

Codeforcesはフルフィードバックはないことを加味すると、自分の性格上、AtCoder以上に早解き勝負に期待できないので、
すこしでも詰まったら積極的に次の問題にシフトしたほうが良さそうです。
