方針は悪くはなかったけど、C2で慣れないことをしてしまってバグらせて破滅してしまいました（pretestでは露呈せず、system testでREして発覚）。

悲劇を繰り返さないように、ちゃんと記録を残しておきます。

※D, Eともに面白そう、かつ勉強にもなりそうなので、取り組み次第追記します。

<!-- TOC -->

- [A. Yet Another Dividing into Teams](#a-yet-another-dividing-into-teams)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84)
  - [解答](#%e8%a7%a3%e7%ad%94)
- [B. Books Exchange (hard version)](#b-books-exchange-hard-version)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-1)
  - [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Good Numbers](#c-good-numbers)
  - [問題の要約](#%e5%95%8f%e9%a1%8c%e3%81%ae%e8%a6%81%e7%b4%84-2)
  - [easy versionの解答](#easy-version%e3%81%ae%e8%a7%a3%e7%ad%94)
  - [hard versionの解答](#hard-version%e3%81%ae%e8%a7%a3%e7%ad%94)
  - [Editorialのスマートな解法](#editorial%e3%81%ae%e3%82%b9%e3%83%9e%e3%83%bc%e3%83%88%e3%81%aa%e8%a7%a3%e6%b3%95)

<!-- /TOC -->

<a id="markdown-a-yet-another-dividing-into-teams" name="a-yet-another-dividing-into-teams"></a>
## A. Yet Another Dividing into Teams

[問題URL](https://codeforces.com/contest/1249/problem/A)

<a id="markdown-問題の要約" name="問題の要約"></a>
### 問題の要約

整数からなる数列が与えられるので、2つのできるだけ少ない数のチームに分けたい。

ただし、同じチーム内に整数同士の差の絶対値が1となるようなペアが存在してはいけない。

<a id="markdown-解答" name="解答"></a>
### 解答

数列をソートしたときに、隣同士の差がすべて2以上ならば、全員同じチームに入れて問題ないので答えは `1` 。
隣同士以外の差の絶対値は確実に2以上になるので、これで問題は起きない。

一方で、隣同士で差が1となるペアが一組でも存在したら、それらは別々のチームに分けなければならない。
また、2組以上存在しても、それらが別のチームに分かれるように配分すれば問題ないので、この場合は `2` とすればよい。

```go
var q int

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		n := ReadInt()
		A := ReadIntSlice(n)

		sort.Sort(sort.IntSlice(A))

		diffNum := 0
		for i := 0; i < n-1; i++ {
			if AbsInt(A[i]-A[i+1]) == 1 {
				diffNum++
			}
		}

		if diffNum > 0 {
			fmt.Println(2)
		} else {
			fmt.Println(1)
		}
	}
}
```

<a id="markdown-b-books-exchange-hard-version" name="b-books-exchange-hard-version"></a>
## B. Books Exchange (hard version)

[問題URL](https://codeforces.com/contest/1249/problem/B2)

<a id="markdown-問題の要約-1" name="問題の要約-1"></a>
### 問題の要約

「1からN番まで採番された子どもたちが、次の日自分が持っている本を上げる相手が記された配列」が与えられる。

もともと自分が持っている本がいずれまた自分に帰ってくるが、その日数を答えよ、という問題。

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

どの子供たちもなんらかのループに組み込まれるので、そのループの長さを答えれば良い。

easy, hardともに再帰関数を使う同じコードを提出した。

easyの方は制約的に、一度調べたループももう一度調べるような効率の悪いコードを書いても通りそうだが、
hardの方はそれではTLEするはず。

```go
var q int
var n int
var P []int
var flags []bool
var answers []int

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		n = ReadInt()
		P = []int{-1}
		tmp := ReadIntSlice(n)
		P = append(P, tmp...)
		flags = make([]bool, n+1)
		for j := 0; j <= n; j++ {
			flags[j] = false
		}
		answers = make([]int, n+1)

		for j := 1; j <= n; j++ {
			if flags[j] {
				continue
			}
			sub(j, 0)
		}
		fmt.Println(PrintIntsLine(answers[1:]...))
	}
}

func sub(id, num int) int {
	if flags[id] {
		return num
	}

	flags[id] = true
	nextId := P[id]
	answers[id] = sub(nextId, num+1)
	return answers[id]
}
```

コンテスト後にtwitter観て気づきましたが、Union-Findでも求まりますね。

ここ数ヶ月Union Find使った記憶がないので、そろそろ手段として思い出しておかないとまずそうです。

<a id="markdown-c-good-numbers" name="c-good-numbers"></a>
## C. Good Numbers

[問題URL](https://codeforces.com/contest/1249/problem/C2)

<a id="markdown-問題の要約-2" name="問題の要約-2"></a>
### 問題の要約

3の冪数（`1, 3, 9, 27, ...`）の部分和で表される数（同じ数値は2回以上使ってはいけない）をGood Numbersと定義する。

与えられた `n` に対して、 `n <= m` を満たす最小のGood Number `m` を答えよ、という問題。

<a id="markdown-easy-versionの解答" name="easy-versionの解答"></a>
### easy versionの解答

`3^9` を計算してみると `19683` であるため、答えるべきGood Numberを考えるにあたって、
`3^10` 以上の冪数は考える必要がない。

また、 `3^9` までの冪数を使った部分和の数は、bit全探索で `2^10` 個全て列挙できる。

列挙したGood Numbersを配列に集めてソートし、二分探索すれば境界となる `m` を計算できる。

```go
var q int
var n int
var G []int
var pows [20]int

func main() {
	q = ReadInt()

	for i := 0; i < 10; i++ {
		pows[i] = PowInt(3, i)
	}

	G = []int{}
	// すべてのgood numbersを集めておく
	for i := 0; i < 1<<10; i++ {
		val := 0
		for j := 0; j < 10; j++ {
			if NthBit(i, j) == 1 {
				val += pows[j]
			}
		}

		G = append(G, val)
	}

	sort.Sort(sort.IntSlice(G))
	for i := 0; i < q; i++ {
		n = ReadInt()

		// m は中央を意味する何らかの値
		isOK := func(m int) bool {
			if G[m] >= n {
				return true
			}
			return false
		}

		ng, ok := -1, len(G)
		for int(math.Abs(float64(ok-ng))) > 1 {
			mid := (ok + ng) / 2
			if isOK(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}
		fmt.Println(G[ok])
	}
}
```

よくよく考えたら、easyの制約ではクエリごとにGood Numbers配列すべて線形探索しても間に合いますね。

<a id="markdown-hard-versionの解答" name="hard-versionの解答"></a>
### hard versionの解答

今度は `1 <= n <= 10^18` と制約がとても大きくなっている。

同じ要領で全探索しようと思い、 `3^39` を計算すると `4*10^18` 以上と確認できる。

よって `3^40` 以上の冪数は考えなくてよいが、easyの要領でGood Numbersを全列挙しようとしても、
`2^40` 個になってしまうため、同じ方法ではメモリも時間も足りない。

ここで、`2^20` ならば全列挙できることに着目し、半分全列挙を考える。

まず、 `3^0 ~ 3^19` までの冪数を用いたGood Numbersを全列挙し、ソートする。
次に、 `3^20 ~ 3^39` までの冪数を用いたGood Numbersを全列挙し、ソートする。
これらから互いに1つずつ選んで和を取ると、すべてのGood Numbersを列挙できる。

蟻本に従って、大きい方のGood Numbersすべてに対して、小さい方のGood Numbersの境界を探る二分探索をするやり方だと、
1クエリあたりおおよそ `2^20 * log(2^20)` ステップと見積もれるので、最大で500個のクエリに対してこれでは間に合わないとわかる。

よくよく考えると、大きい方のGood Numbersについても単調性があることがわかる。
すなわち、「ある大きい方のGood Numberを採用したときに、
小さい方のGood Numberから何かしら選んで `n` 以上とできるならば、
大きい方のGood Numbersからさらに大きいものを選んでも `n` 以上の和を達成することは可能である」と言える。

よって結局、それぞれのGood Numbersの配列を二分探索することでも、題意に沿うGood Numbersが計算できる。

二分探索の条件部分でさらに二分探索をすることになるため、ある `n` に対する計算ステップは、
大体 `log((log2^20)^2)` と見積もれる。

全体の、計算量は `O(q * (log(2^(n/2)))^2)` で間に合う。

```go
var q int
var n int64

// var G []int
var befG, aftG DirRange
var pows [50]int64

type DirRange []int64

func (a DirRange) Len() int           { return len(a) }
func (a DirRange) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DirRange) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	q = ReadInt()

	for i := 0; i < 40; i++ {
		pows[i] = PowInt64(3, int64(i))
	}

	befG, aftG = make(DirRange, 0, 1000000), make(DirRange, 0, 1000000)
	// すべてのgood numbersを集めておく
	for i := 0; i < 1<<20; i++ {
		val := int64(0)
		for j := 0; j < 20; j++ {
			if NthBit(i, j) == 1 {
				val += pows[j]
			}
		}

		befG = append(befG, val)
	}
	for i := 0; i < 1<<20; i++ {
		val := int64(0)
		for j := 0; j < 20; j++ {
			if NthBit(i, j) == 1 {
				val += pows[20+j]
			}
		}

		aftG = append(aftG, val)
	}

	sort.Sort(befG)
	sort.Sort(aftG)

	for i := 0; i < q; i++ {
		n = ReadInt64()

		aftIdx := decideAftIdx()
		befIdx := decideBefIdx(aftIdx)

		fmt.Println(aftG[aftIdx] + befG[befIdx])
	}
}

func decideAftIdx() int {
	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		befIdx := decideBefIdx(m)
		if befIdx < len(befG) {
			return true
		}
		return false
	}

	ng, ok := -1, len(aftG)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func decideBefIdx(m int) int {
	aftVal := aftG[m]

	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		s := aftVal + befG[m]
		if s >= n {
			return true
		}
		return false
	}

	ng, ok := -1, len(befG)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// PowInt is integer version of math.Pow
// PowInt calculate a power by Binary Power (二分累乗法(O(log e))).
func PowInt64(a, e int64) int64 {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := PowInt64(a, halfE)
		return half * half
	}

	return a * PowInt64(a, e-1)
}
```

コンテスト中に提出したコードでは、「二分探索の条件部分でさらに二分探索する」部分でサボった実装をしてしまい、
特定のケースで落ちる書き方をしていました。

せっかくわざわざ蟻本開いて初めて実践で使ってみたのに、残念。

<a id="markdown-editorialのスマートな解法" name="editorialのスマートな解法"></a>
### Editorialのスマートな解法

通している人数からしてももっと簡単なやり方があるのだと思い、こちらもさらっておきます。

Good Numberとは、3進数表現したときに、各桁がすべて `0, 1` のいずれかで表される数と言い換えられます。

なので、与えられた `n` を3進数表現したとき、 `2` の桁が存在しなければそれがそのまま答えになります。
一方で、 `2` がどこかの桁に登場したときは、貪欲法の感覚（というよりは桁DPでよく考えるアレ）で、
`n` より大きい最小の数を構成することを考えます。

このために、以下の図に示すように、 `n` の3進数表現で2が登場する最上位の桁を調べます。
この桁以下の `2` はすべて潰す必要があるためです。



また、現在の `n` より大きい数値とするために、調べた `2` の最上位桁のさらに上に存在する、
`0` の桁で最初に出会うもの（最下位桁のもの）を調べます。
この桁を `1` とすることで、あらたに `2` の桁は登場させずに、 `n` よりも大きなGood Numberを構成できます。
さらに、立てた `1` よりも小さい桁はすべて `0` としてしまっても、依然として `n` よりは大きいGood Numberとできるので、
これが答えとなります。

```go
var q int
var pows [40]int64

func main() {
	q = ReadInt()

	pows[0] = 1
	for i := 1; i < 40; i++ {
		pows[i] = pows[i-1] * 3
	}

	for tc := 0; tc < q; tc++ {
		inn := ReadInt64()
		n := inn

		ternary := [40]int{}
		for i := 0; i < 40; i++ {
			ternary[i] = int(n % 3)
			n /= 3
		}

		twoIdx, zeroIdx := -1, -1
		for i := 0; i < 40; i++ {
			if ternary[i] == 2 {
				twoIdx = i
			}
		}
		if twoIdx == -1 {
			fmt.Println(inn)
			continue
		}

		for i := twoIdx + 1; i < 40; i++ {
			if ternary[i] == 0 {
				zeroIdx = i
				ternary[i] = 1
				break
			}
		}

		ans := int64(0)
		for i := zeroIdx; i < 40; i++ {
			ans += pows[i] * int64(ternary[i])
		}

		fmt.Println(ans)
	}
}
```

こういう見方が冷静にできないとだめだなぁと反省しました。

---

ものにもよるだろうけど、C以降でのeasyとhardでは、一旦思考をリセットしたほうが良さそう。
