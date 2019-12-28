<!-- Educational Codeforces Round No.78 参加記録 (A〜C解答) -->

いつもの自分だったらBの算数で詰まって終了だったので、
成長を喜びたいところ。

<!-- TOC -->

- [A. Shuffle Hashing](#a-shuffle-hashing)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. A and B](#b-a-and-b)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Berry Jam](#c-berry-jam)
	- [解答](#%e8%a7%a3%e7%ad%94-2)

<!-- /TOC -->

<a id="markdown-a-shuffle-hashing" name="a-shuffle-hashing"></a>
## A. Shuffle Hashing

[問題のURL](https://codeforces.com/contest/1278/problem/A)

<a id="markdown-解答" name="解答"></a>
### 解答

ハッシュ文字列からパスワード分の長さの分だけ文字列を切り出して調べる、というのを全探索すれば良い。

切り出した部分とパスワードが一致するかどうかの判定は、文字配列をソートして
文字列として一致するかどうかを調べるのが簡単（だと思う）。

```go
var t int
var P, H []rune

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		P, H = ReadRuneSlice(), ReadRuneSlice()

		solve()
	}
}

func solve() {
	if len(P) > len(H) {
		fmt.Println("NO")
		return
	}

	pL := RuneList{}
	for i := 0; i < len(P); i++ {
		pL = append(pL, P[i])
	}
	sort.Sort(pL)

	plen := len(P)
	for i := 0; i < len(H); i++ {
		if i+plen > len(H) {
			break
		}

		tmp := H[i : i+plen]
		hL := RuneList{}
		for j := 0; j < len(tmp); j++ {
			hL = append(hL, tmp[j])
		}
		sort.Sort(hL)

		if string(pL) == string(hL) {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}

type RuneList []rune

func (a RuneList) Len() int           { return len(a) }
func (a RuneList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RuneList) Less(i, j int) bool { return a[i] < a[j] }
```

ハッシュからパスワードを切り出す全探索の部分がコーディングに時間かかりすぎてしまっていたので、
ちょっと良くないですね。
どうしよう。

<a id="markdown-b-a-and-b" name="b-a-and-b"></a>
## B. A and B

[問題のURL](https://codeforces.com/contest/1278/problem/B)

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

2つの数 `a, b` を等しくするためには、最終的に2数の差 `diff = abs(a-b)` を埋める必要がある。

必要な操作回数を `k` 回としたときに、2つの数に加える数値の総和は `k*(k+1)/2` である。
よって、この総和を2つの数に分割して、その2つの数の差が `diff` とすることができるのかを考える。

総和は好きなように分割できる一方で、分割した2つの数の差の偶奇は1種類のみである。

例: `sum == 6` のとき `(0, 6), (1, 5), (2, 4), (3, 3)` が分割の全パターンで、2つの数の差はすべて偶数となる。

以上から、 `k` を小さいところから全探索して、 `diff%2 == sum%2 && diff <= sum` を満たすものを見つければ良い。

総和の形から1テストケースあたりの計算量は `O(Sqrt(Max(a, b)))` なので間に合う。

```go
var t int
var a, b int64

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		a, b = ReadInt64_2()

		solve()
	}
}

func solve() {
	diff := AbsInt(a - b)

	if diff == 0 {
		fmt.Println(0)
		return
	}

	for k := int64(1); ; k++ {
		sum := k * (k + 1) / 2

		if diff%2 == sum%2 && diff <= sum {
			fmt.Println(k)
			return
		}
	}
}
```

> 総和は好きなように分割できる一方で、分割した2つの数の差の偶奇は1種類のみである。

これって自明ですかね？
私は紙に色々書いているうちに「それはそうだ」と納得できる、ぐらいなんですが。

<a id="markdown-c-berry-jam" name="c-berry-jam"></a>
## C. Berry Jam

[問題のURL](https://codeforces.com/contest/1278/problem/C)

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

以下の図に示すように、いちごとブルーベリーそれぞれを `1, -1` に変換する（RはいちごでBはブルーベリー）。



ここで、与えられた配列の真ん中で分割する。
扱いやすくするため、食べられる順番に従うように、左側については元の配列の内側から並ぶように、配列を構築する。

この2つの配列に対して、累積和を考える。
このようにすると、問題は「左の累積和から1つ、右の累積和から1つ数を選んで、全体の合計と等しくなるようにせよ」
と読み替えることができる。
また、答えは左のインデックス、右のインデックスの和が小さいほど望ましい。

このように考えると、それぞれの配列についてある累積和を取るインデックスは、その最小のもののみに興味があるといえる
（同じ数の累積和というのは、食べ残ったRとBの数の差が同じであり、等価であると言えるから（DPっぽい考え方））。

以下の実装では、それぞれの累積和に対する最小のインデックスを、左右の配列分それぞれ辞書で用意した
（負数もまとめて入れたかったため）。

あとは、あり得る和のパターンの全探索を、左配列基準および右配列基準で2スキャンして、ベストなものを見つける。

```go
var t int
var n int
var A []int

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A = ReadIntSlice(2 * n)

		solve()
	}
}

func solve() {
	E := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if A[i] == 2 {
			E[i] = -1
		} else {
			E[i] = 1
		}
	}
	sum := Sum(E...)

	if sum == 0 {
		fmt.Println(0)
		return
	}

	if sum < 0 {
		for i := 0; i < len(E); i++ {
			E[i] = -E[i]
		}
	}
	sum = Sum(E...)

	R, L := []int{}, []int{}
	for i := n - 1; i >= 0; i-- {
		L = append(L, E[i])
	}
	for i := 1; i < n; i++ {
		L[i] += L[i-1]
	}
	for i := n; i < 2*n; i++ {
		R = append(R, E[i])
	}
	for i := 1; i < n; i++ {
		R[i] += R[i-1]
	}

	leftMemo, rightMemo := make(map[int]int), make(map[int]int)
	leftMemo[0], rightMemo[0] = -1, -1
	for i := 0; i < n; i++ {
		a := L[i]
		if _, ok := leftMemo[a]; ok {
			continue
		} else {
			leftMemo[a] = i
		}
	}
	for i := 0; i < n; i++ {
		a := R[i]
		if _, ok := rightMemo[a]; ok {
			continue
		} else {
			rightMemo[a] = i
		}
	}

	ans := INF_BIT30
	for lv := 0; lv <= sum; lv++ {
		rv := sum - lv
		lidx, lok := leftMemo[lv]
		ridx, rok := rightMemo[rv]
		if lok && rok {
			ChMin(&ans, lidx+ridx+2)
		}
	}
	for rv := 0; rv <= sum; rv++ {
		lv := sum - rv
		lidx, lok := leftMemo[lv]
		ridx, rok := rightMemo[rv]
		if lok && rok {
			ChMin(&ans, lidx+ridx+2)
		}
	}
	fmt.Println(ans)
}
```

コンテスト後に時間を使って自分で考えた方法ですが、良い解法なのかはわかりません。
正直変換手数も多くてなんかまどろっこしい気がします。
editorialが公開された上で、模範解答が面白ければ、そちらも追記するかもしれません。

---

今回はなんとかイーブンに持ち込めました。

なんとか粘って今年の残りあと3回のこどふぉで青に持ち込みたいところです。

