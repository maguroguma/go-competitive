<!-- Codeforces Round No.602 Div.2 D2復習 -->

コンテスト中に解けなかったものの復習です。

いろいろな解法（というよりも解くために用いるツールが多様）がありますが、
BITを使った方法が一番自分にとって与し易かったため、BITで解きました。

- [D2. Optimal Subsequences (Hard Version)](#d2-optimal-subsequences-hard-version)
  - [問題](#%e5%95%8f%e9%a1%8c)
  - [解答](#%e8%a7%a3%e7%ad%94)
  - [BITを使って k 番目の要素を取得するテクニック](#bit%e3%82%92%e4%bd%bf%e3%81%a3%e3%81%a6-k-%e7%95%aa%e7%9b%ae%e3%81%ae%e8%a6%81%e7%b4%a0%e3%82%92%e5%8f%96%e5%be%97%e3%81%99%e3%82%8b%e3%83%86%e3%82%af%e3%83%8b%e3%83%83%e3%82%af)
    - [BIT上の二分探索とは？](#bit%e4%b8%8a%e3%81%ae%e4%ba%8c%e5%88%86%e6%8e%a2%e7%b4%a2%e3%81%a8%e3%81%af)

## D2. Optimal Subsequences (Hard Version)

[問題のURL](https://codeforces.com/contest/1262/problem/D2)

### 問題

`n` 要素からなる数列 `A` が与えられる。

この数列に対して部分列（間の要素を好きなだけ削除し、残ったものの順番を変えずに得られる数列）を考える。

ある部分列の長さ `k` が与えられたとき、以下の条件を満たす部分列はoptimalであるという。

1. 考えられる長さ `k` のあらゆる部分列の中で、部分列の要素の総和が最大となる。
2. 1の条件を満たす部分列の中で、元の数列の位置に関して辞書順最小である。

一方で、 `m` 個のクエリが与えられる。

各クエリは `k, pos` の2つの1以上の整数からなり、このクエリに対して、
「長さ `k` のoptimalな部分列における `pos` 番目の値」を答える必要がある。

すべてのクエリに対して、それらの順番どおりに答えよ。

制約:

- `1 <= n, m <= 2*10^5`
- `1 <= A[i] <= 10^9`
- `1 <= k <= n, 1 <= pos <= k`

### 解答

Easyバージョンは制約が小さいため、愚直な解法で通りますが、こちらは賢くクエリを処理する必要があります。

まず整理すると、
1の条件よりoptimalな部分列は、元の数列を降順にソートしその先頭 `k` 個の数を選択したものとなります。
ただし、2の条件より、同じ要素が元の数列に複数存在する場合、できるだけ元の数列において前の方の位置に登場したものを優先的に選ぶ必要があります。

よって、元の数列 `A` を、まずは要素の大きさを基準に降順となるようにし、値が同じ場合は元の数列における位置に関して昇順となるようにソートします。
各クエリに答えるためには、まずこのソートした列に対して、先頭の `k` 要素を取得します。
そして、取得した `k` 要素の中から、位置に関して `pos` 番目に小さい要素の値を出力すればOKです。



ただし、これを愚直に行うためには、 `k` 要素取得するたびにそれらを位置に関して昇順ソートする必要があります。
これは許容できないため、工夫が必要です。

まず、クエリをすべて先読みし、 `k` が小さい順に答えていくという工夫ができます。
こうすることで、（許されるなら）ソートしたい集合が、単純に新しい要素がappendされていくだけとなり、シンプルになります。

しかしながら、それでも `k` が大きくなるたびにソートするわけには行かないため、
「要素の追加」と「特定の順番の値の取得」を高速に行う必要があります。

このような操作は「BIT上の二分探索」を活用することで可能なため（詳しくは後述）、各クエリを `O(logn)` で処理できます。

よって、全体で `O(m*logn)` で解くことができます。

※1500msecぐらいかかりました。

```go
var n int
var A []int
var m int

type BinaryIndexedTree struct {
	bit     []int
	n       int
	minPow2 int
}

// n(>=1) is number of elements of original data
func NewBIT(n int) *BinaryIndexedTree {
	newBit := new(BinaryIndexedTree)

	newBit.bit = make([]int, n+1)
	newBit.n = n

	newBit.minPow2 = 1
	for {
		if (newBit.minPow2 << 1) > n {
			break
		}
		newBit.minPow2 <<= 1
	}

	return newBit
}

// Sum of [1, i](1-based)
func (b *BinaryIndexedTree) Sum(i int) int {
	s := 0

	for i > 0 {
		s += b.bit[i]
		i -= i & (-i)
	}

	return s
}

// Add x to i(1-based)
func (b *BinaryIndexedTree) Add(i, x int) {
	for i <= b.n {
		b.bit[i] += x
		i += i & (-i)
	}
}

// LowerBound returns minimum i such that bit.Sum(i) >= w.
func (b *BinaryIndexedTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

	x := 0
	for k := b.minPow2; k > 0; k /= 2 {
		if x+k <= b.n && b.bit[x+k] < w {
			w -= b.bit[x+k]
			x += k
		}
	}

	return x + 1
}

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)
	m = ReadInt()

	// クエリを先読みしてkで昇順ソート
	L := make(QueryList, 0)
	for i := 0; i < m; i++ {
		k, pos := ReadInt2()
		L = append(L, &Query{id: i, k: k, pos: pos}) // idは0-basedで格納
	}
	sort.Stable(L)

	// Aを値で降順→位置で昇順にソート
	S := make(ElementList, 0)
	for i := 0; i < n; i++ {
		S = append(S, &Element{val: A[i], pos: i}) // posは0-basedで格納
	}
	sort.Stable(S)

	// BITをOrderedSetとして扱う
	bit := NewBIT(n)
	ck := 0
	// クエリの回答保持用
	answers := make([]int, m)
	// クエリをkの小さい順に処理していく
	for i := 0; i < m; i++ {
		q := L[i]
		// BITに格納した要素数がk個になるまで追加する
		for ; ck < q.k; ck++ {
			e := S[ck]
			bit.Add(e.pos+1, 1)
		}

		// BIT中のpos番目を回答する
		ans := bit.LowerBound(q.pos)
		answers[q.id] = A[ans-1]
	}

	// まとめて回答
	for i := 0; i < m; i++ {
		fmt.Println(answers[i])
	}
}

type Query struct {
	id, k, pos int
}
type QueryList []*Query

func (l QueryList) Len() int {
	return len(l)
}
func (l QueryList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l QueryList) Less(i, j int) bool {
	return l[i].k < l[j].k
}

type Element struct {
	val, pos int
}
type ElementList []*Element

func (l ElementList) Len() int {
	return len(l)
}
func (l ElementList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ElementList) Less(i, j int) bool {
	if l[i].val > l[j].val {
		return true
	} else if l[i].val < l[j].val {
		return false
	} else {
		return l[i].pos < l[j].pos
	}
}
```

### BITを使って `k` 番目の要素を取得するテクニック

今回行ったような操作は競技プログラミングにおいては典型テクニックのようで、
例えば[けんちょんさんのQiitaの記事](https://qiita.com/drken/items/1b7e6e459c24a83bb7fd#整理)などでも
詳しく解説されています（他にも色々わかりやすくまとめている方がたくさんおられました）。

今回は、位置に関してBITで管理するため、要素数の制約から座標圧縮する必要はありませんでした。

#### BIT上の二分探索とは？

[hosさんのBIT解説PDF](http://hos.ac/slides/20140319_bit.pdf)において最後の方で説明がなされています。

。。が、自分には最初何をやっているのかよくわかりませんでした。

自分なりに議論を補間しつつゆっくりと追っていくと理解できたので、忘れた頃の未来の自分を第一の対象として、理解の道筋を残しておこうと思います。

まず、「BIT上の二分探索ってそもそも何？」となってしまいましたが、これは「元の配列の先頭からの累積和に関する二分探索」となります。
別の表現をすると **「1-basedなインデックス `i` で、 `[1, i]` の要素の総和が `w` 以上となる最小の `i` を探索する」** といえます。

二分探索自体の計算量が対数オーダーで、BITを使って累積和を求めるのも対数オーダーであるため、
計算量が `log` の2乗になるというのは納得できます。

しかし、BITの木構造をうまく利用することで、この二分探索も全体で `O(logn)` に落とすことができる、とのことです。
それとともに与えられたアルゴリズムが以下のものですが、これまた初見時はよくわかりませんでした。

```go
func (b *BinaryIndexedTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

  // b.minPow2は、n以下の最小の2べき
  x := 0
	for k := b.minPow2; k > 0; k /= 2 {
		if x+k <= b.n && b.bit[x+k] < w {
			w -= b.bit[x+k]
			x += k
		}
	}

	return x + 1
}
```

これは要約すると、「先頭からの累積和を、できるだけ長い区間から足して良いかを都度判断する」ということを行っています。

流れとしては、BITの上方の区間が長いノードから順番に見ており、
`k` というのは現在注目している区間の長さを意味しています。

各区間は、担当する区間の和を持っており、それがkey値 `w` よりも小さい場合は、足してから右側の次の短い長さの区間を見る必要があります。
逆にkey値よりも大きい場合は、その区間和は足さずに左側の次の短い長さの区間を見る必要があります。

以下は、ある具体例におけるアルゴリズムの様子を図示したものです（あまりわかりやすくできませんでしたが。。）。



結局、木の高さ分下ることで探索が終了するため、計算量は `O(logn)` となります。

---

クエリの先読みと、 `k` 番目に小さい要素の取得、という2つの典型テクニックが要求される問題だったと思います。

引用した記事でも言及されている平衡二分探索木についても、なるべく早く実装してみたいと思います。

次に出会ったときはちゃんと倒したいところです。
