<!-- TOC -->

- [A. Heating](#a-heating)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Obtain Two Zeroes](#b-obtain-two-zeroes)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Infinite Fence](#c-infinite-fence)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. A Game with Traps](#d-a-game-with-traps)
	- [問題の概要](#%e5%95%8f%e9%a1%8c%e3%81%ae%e6%a6%82%e8%a6%81-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)

<!-- /TOC -->

<a id="markdown-a-heating" name="a-heating"></a>
## A. Heating

[問題URL](https://codeforces.com/contest/1260/problem/A)

<a id="markdown-問題の概要" name="問題の概要"></a>
### 問題の概要

`k` セクションある1つの暖房器具を設置するとコストが `k^2` かかる。
最大 `c` 個の暖房器具を設置することで合計 `sum` セクション確保したい。
このときに必要となるコストの最小値はいくらか計算せよ、という問題。

<a id="markdown-解答" name="解答"></a>
### 解答

英語とサンプルでやるべきことが理解できたら80％ぐらいACだと思う（heating radiatorはともかくsectionって何？）。

[以前のこどふぉで解いた問題](https://maguroguma.hatenablog.com/?page=1572109201#b-grow-the-tree)の、逆に今度は小さくするバージョン、という感じ。

結論から言うと、できるだけ `c` 個の暖房器具のセクションが均等になるように選べば良い。
直感的には、2次元の場合のマンハッタン距離を考えたとき、
マンハッタン距離が同じ点の中では `x=y` の点がユークリッド距離が一番小さくなる、というのを多次元に考えている。

本番中に書いたコードでは、 `c >= sum` の場合を例外的に扱っているが、多分これは不要。

```go
var n int
var C, S []int64

func main() {
	n = ReadInt()
	C, S = make([]int64, n), make([]int64, n)
	for i := 0; i < n; i++ {
		c, s := ReadInt64_2()
		C[i], S[i] = c, s
	}

	for i := 0; i < n; i++ {
		c, sum := C[i], S[i]

		if c >= sum {
			fmt.Println(sum)
			continue
		}

		x := sum / c
		m := sum % c
		ans := x*x*(c-m) + (x+1)*(x+1)*m
		fmt.Println(ans)
	}
}
```

15分はかかり過ぎだが、英語が難解すぎた。

<a id="markdown-b-obtain-two-zeroes" name="b-obtain-two-zeroes"></a>
## B. Obtain Two Zeroes

[問題URL](https://codeforces.com/contest/1260/problem/B)

<a id="markdown-問題の概要-1" name="問題の概要-1"></a>
### 問題の概要

非負整数 `a, b` が与えられる。

以下の2つの操作、

1. `a := a - 2*x, b := b - x`
2. `a := a - x, b := b - 2*x`

を何回でも実行可能なとき、両方を同時に `0` にすることはできるか判定する問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

まず、それぞれの操作において `x = 1` に操作を限定して良い
（ `x = n` とした場合、選んだ操作を `x=1` で `n` 回実行した、と等しく考えることができる）。

それぞれの操作について、操作後の `a, b` の和を計算すると、ともに `a + b - 3` となる。
つまり、いずれの操作を選んだとしても `a, b` の和は `3` ずつ減っていくこととなる。

よって、両方同時に `0` にするためには、 `a, b` の和が `3` の倍数であることが必要条件となる。
そして `a, b` の和が `3` の倍数のとき、 `(a+b)/3 = m` とすると、 `m` は両方同時に `0` にするために必要な操作回数となる。

また、1の操作を行った場合、 `a` の方を `b` よりも `1` 多く減らせることになり、2の操作ではその逆となる。
つまり、 `Abs(a - b) = diff` とすると、 `diff` 分は両方同時に `0` とするために差を詰める必要があり、
これには `diff` 回のいずれか片方の操作が必要となる。

よって `diff > m` ならば先に片方が0に到達してしまうのでNOとなる。

一方で、 `diff <= m` ならば `diff` 回の操作で `a, b` を等しくすることが出来、
残った回数で `a+b` をピッタリ `0` にできることも保証されているので、結局、 `diff > m` ならばNO、そうでないならばYESとすればよい。
（コンテスト中はもう少し複雑に考えてしまった。）

```go
var t int
var a, b int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		a, b = ReadInt2()

		solve()
	}
}

func solve() {
	if (a+b)%3 != 0 {
		fmt.Println("NO")
		return
	}

	m := (a + b) / 3
	diff := AbsInt(a - b)

	if diff > m {
		fmt.Println("NO")
	} else {
    fmt.Println("YES")
  }
}
```

最近AtCoderで和を考えると不変量が見える〜みたいなのが2回ほどあったので、その反省が活かせて嬉しい。
と思ってたら、フレンドがみんな難なく通していてやっぱり悲しい。

<a id="markdown-c-infinite-fence" name="c-infinite-fence"></a>
## C. Infinite Fence

[問題URL](https://codeforces.com/contest/1260/problem/C)

<a id="markdown-問題の概要-2" name="問題の概要-2"></a>
### 問題の概要

無限に左から右へと並べられた板の列があり、左から `0, 1, ...` と採番されている。

`r, b` の2つの正の整数が与えられ、以下のルールに基づいて、この板を塗っていく。

- 板の番号が `r` で割り切れるのならば、その板を赤で塗る。
- 板の番号が `b` で割り切れるのならば、その板を青で塗る。
- 板の番号が `r, b` のいずれでも割り切れるのならば、その板を赤か青の好きな色で塗る。
- いずれでもないならば、その板は塗ってはいけない。

このようにして無限の板を塗っていき、塗られていない板を除外する。

このとき、連続する `k` 個の板が同じ色で塗られることを避けられるかどうかを判定する問題。

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

とりあえず図を描いて考えてみると、 `r, b` の公約数を周期として、同じようなパターンが続くことがイメージできる。
そこで、最大公約数までをもう少し掘り下げて考えてみることにする。

まず、 `r >= b` であると仮定する（そうでない場合は、2つの数を入れ替えて考えれば良い）。

例えば、 `r = 5, b = 2` のケースを図示すると、連続する列の長さが長くなるのは、値が小さい青色の方であるとわかる。

<figure class="figure-image figure-image-fotolife" title="r=5, b=2のサンプル例">[f:id:maguroguma:20191130172137j:plain]<figcaption>r=5, b=2のサンプル例</figcaption></figure>

よって、 `r >= b` で一般化して、ある赤と赤に囲まれた区間を考える。
図のように、どのような区間でも青は `b` 間隔で並び、右端の赤が `b` の約数である、すなわち `r, b` の公約数であったとしても、
それは赤として考えれば良い（そのほうが `k` 個連続するのを避けるためには有利であるため）。

<figure class="figure-image figure-image-fotolife" title="一般化して一部分をフォーカス">[f:id:maguroguma:20191130172216j:plain]<figcaption>一般化して一部分をフォーカス</figcaption></figure>

ここで、板全体を俯瞰したときに、青の連続する列の長さが最長となるのは、区間の左端の赤と青の板の番号の差が最小となる場合であるとわかる。

サンプルをいくつか試すと、「なんとなく `Gcd(a, b)` が最小値となりそう」とわかる。

。。コンテスト中はこれ以上時間を費やせず、とりあえずWAしたらまた考えようという気持ちでsubmitしたら通ってしまった。

```go
var t int
var r, b, k int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		r, b, k = ReadInt3()

		solve()
	}
}

const obey = "OBEY"
const rebel = "REBEL"

func solve() {
	if r < b {
		r, b = b, r
	}

	g := Gcd(r, b)
	m := (r-1-g)/b + 1
	if m < k {
		fmt.Println(obey)
	} else {
		fmt.Println(rebel)
	}
}
```

以下は、コンテスト後に書いた簡単な証明です（雑な手書きですみません）。

<figure class="figure-image figure-image-fotolife" title="仮説部分の正当性の証明">[f:id:maguroguma:20191130172252j:plain]<figcaption>仮説部分の正当性の証明</figcaption></figure>

なんとかコンテスト中に解けたけど、時間がかなりかかったし難しい。。

具体的な例で仮説を立て、抽象化して検証する、場合によってはまた具体的な例に戻る、というのを自分は特別意図せずやっていますが、
もう少しスマートに短時間でできれば良いなぁとは思います（賢くなりたい）。

<a id="markdown-d-a-game-with-traps" name="d-a-game-with-traps"></a>
## D. A Game with Traps

[問題URL](https://codeforces.com/contest/1260/problem/D)

<a id="markdown-問題の概要-3" name="問題の概要-3"></a>
### 問題の概要

※大分端折っているので、詳しくは本文をご参照ください。

`0, 1, .., n+1` のマス目を初期位置 `0` からゴール `n+1` を軍隊を引き連れながら目指す。

ただし、道中にはトラップがあるため、軍隊の兵士がそれを踏んで死なないようにして進まなければならない。
一方で、兵士1人1人に設定されている `agility` のパラメータが、トラップの威力以上である場合には、兵士はそのトラップの影響を受けない。

トラップは、トラップの先にある解除用ボタンの位置まで到達することで解除できる。

軍隊は、自身と一緒にしか移動できず、1マス移動するには1秒かかる。

`t` 秒間与えられたときに、ゴールまで引率できる兵士の数の最大値はいくらかを堪える問題。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

引率できる兵士はagilityの高い順であり、あるagilityの兵士がゴールできるならば、それ以上のagilityの兵士も全員ゴールできる。

よって、ゴールできるagilityの最小値を二分探索で求めることを考える。

あるagilityの兵士がゴールにたどり着けるかどうかの判定は、以下のようにして可能である。

あるトラップに対して、区間 `[l, r]` を考えると、この区間に関しては、少なくとも3回通過する必要がある。
すなわち、トラップの解除に向かうとき、元の位置に戻るとき、軍隊を引率して進行するときの3回である。

トラップの区間が交差している場合は、逐次トラップの区間を3回通過する必要はなく、
交差する区間をすべてマージして、その区間を3回通過するほうが、経過時間は短くなる。

よって、現在注目中のagilityに対して、影響を考慮すべきトラップの区間をマージすることで、
必要な時間を計算できる。
すなわち、マージ後の区間に関して、全区間が含んでいるマス目の合計数を `T` とすると、
`n + 1 + 2*T` で計算できる。

```go
var m, n, k, t int
var A []int

type Trap struct {
	key     int
	l, r, d int
}
type TrapList []*Trap
type byKey struct {
	TrapList
}

func (l TrapList) Len() int {
	return len(l)
}
func (l TrapList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.TrapList[i].key < l.TrapList[j].key
}

// how to use
// L := make(TrapList, 0, 200000+5)
// L = append(L, &Trap{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

var L TrapList

func main() {
	m, n, k, t = ReadInt4()
	A = ReadIntSlice(m)
	L = make(TrapList, 0)
	for i := 0; i < k; i++ {
		l, r, d := ReadInt3()
		L = append(L, &Trap{key: l, l: l, r: r, d: d})
	}
	maxAgility := Max(A...)

	// 区間の左端で昇順ソート
	sort.Stable(byKey{L})

	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		if C(m) {
			return true
		}
		return false
	}

	ng, ok := -1, maxAgility+1
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	minAgility := ok

	num := 0
	for i := 0; i < m; i++ {
		if A[i] >= minAgility {
			num++
		}
	}
	fmt.Println(num)
}

func C(m int) bool {
	segments := []Trap{}
	l, r := 0, -1
	for i := 0; i < len(L); i++ {
		t := L[i]
		if t.d <= m {
			continue
		}

		if r == -1 {
			l, r = t.l, t.r
			continue
		}

		if r >= t.l-1 {
			// マージして継続
			ChMax(&r, t.r)
		} else {
			// マージせず中断して追加
			segments = append(segments, Trap{l: l, r: r})
			l, r = t.l, t.r
		}
	}
	if r != -1 {
		segments = append(segments, Trap{l: l, r: r})
	}

	time := 1 + n
	for _, seg := range segments {
		time += 2 * (seg.r - seg.l + 1)
	}

	return time <= t
}
```

本質部分である「トラップの区間は少なくとも3回通過する必要がある」という部分が整理できずに、コンテスト中は解くことが出来ませんでした。

二分探索の判定部分に関しては、区間更新可能な遅延評価ありのセグメント木や、
いもす法を使うことでもOKです（セグ木解法は `O(N(logN)^2)` なので、ちょっと危なそうですが）。

---

1次不定方程式とか中国剰余定理ともう少し仲良くなりたい。
