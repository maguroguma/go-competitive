Bを難しく感じてしまったので、要点を整理して類題に備えたいところ。

あと何故かCでGoの気の利かせたスライス確保をしたら謎のTLE2回出してしまったので、これからはやらないように。

※Bは想定解法がもっとスマートなはずなので、Editorialが公開され次第、追記します。
※Dは面白そうなので、取り組み次第追記します。

## A. Broken Keyboard

[問題URL](https://codeforces.com/contest/1251/problem/A)

### 問題の要約

26文字のキーボードについていくつかの壊れたキーがあるので、故障していないキーだけを列挙する問題。

壊れたキーを叩くと、その文字について確実に2回タイプされてしまう、というのをヒントに解く。

### 解答

「確実に故障していない」と断定できるのは、連続で奇数回タイプされた文字だけで、
それを見つけ出して列挙すればいい。

素直に与えられた文字配列を舐めてもいいが、バグが怖いのでランレングス圧縮のスニペットを貼って利用した。

最後の出力についても、同じ文字を2回以上出力してはいけない、ソートして出力する、スペースをいれないなど、
制約が多いので注意。

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

AGCで太陽したときに作ったスニペットですが、直近のABCでも呼び出したりと、
何かと出番があります。

## B. Binary Palindromes

[問題URL](https://codeforces.com/contest/1251/problem/B)

### 問題の要約

0と1からなる複数の文字列が与えられる。

複数の文字列間、あるいは単一の文字列間でもどちらでも良いので、各文字を交換することが何回でもできる。

できるだけ多くの回文を作ろうとしたとき、最大で何個作れるか、という問題。

### 解答

まず、全文字列に渡る0、1の数をそれぞれ集計しておく。

回文を作るときは0でも1でも好きな方を2個ずつとり、それを先頭と末尾に配置する形で作っていく。
奇数長の場合は、最後に真ん中の数を0か1のいずれか一方選ぶことになる。

このように考えると、偶数長の回文を考える場合は、1にしろ0にしろ、その総数を2ずつしか減らさないので、
総数の偶奇は変化しない。
よって、必ず偶数長のものはすべて作ることができる。

また、奇数長の回文を作るときは、最後に真ん中の文字を作るために、
残っている数が奇数の文字を選択する。

このような方法をとったときに、最後の奇数長の回文ができるかできないかのいずれかとなるため、


```go
var q int
var n int
var S [][]rune

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		n = ReadInt()
		S = [][]rune{}
		for j := 0; j < n; j++ {
			S = append(S, ReadRuneSlice())
		}

		one, zero := 0, 0
		for j := 0; j < n; j++ {
			for k := 0; k < len(S[j]); k++ {
				if S[j][k] == '1' {
					one++
				} else {
					zero++
				}
			}
		}

		evens := []int{}
		odds := []int{}
		for j := 0; j < n; j++ {
			if len(S[j])%2 == 0 {
				evens = append(evens, len(S[j]))
			} else {
				odds = append(odds, len(S[j]))
			}
		}
		sort.Sort(sort.IntSlice(evens))
		sort.Sort(sort.IntSlice(odds))

		ans := solve(evens, odds, one, zero)
		fmt.Println(ans)
	}
}

func solve(evens, odds []int, one, zero int) int {
	ans := 0

	// 偶数から
	for i := 0; i < len(evens); i++ {
		enum := evens[i]
		for enum > 0 {
			if one < 2 && zero < 2 {
				break
			}

			enum -= 2
			if one > zero {
				one -= 2
			} else {
				zero -= 2
			}
		}

		if enum == 0 {
			ans++
		}
	}

	// 次に奇数
	for i := 0; i < len(odds); i++ {
		onum := odds[i]
		if one == 0 && zero == 0 {
			break
		}
		onum--
		if one%2 == 1 {
			one--
		} else {
			zero--
		}

		// onumは偶数が確定
		for onum > 0 {
			if one < 2 && zero < 2 {
				break
			}

			onum -= 2
			if one > zero {
				one -= 2
			} else {
				zero -= 2
			}
		}

		if onum == 0 {
			ans++
		}
	}

	return ans
}
```

1と0どちらかを使うかの判定で多く残っている方を採用していますが、
おそらくこの部分は意味がないはず。

嘘解法じゃないよね。。？

## C. Minimize The Integer

[問題URL](https://codeforces.com/contest/1251/problem/C)

### 問題の要約

とても大きい桁数の整数値が与えられる。

隣同士の桁の偶奇が異なる場合は、それらを互いに交換してよい。

この操作が何回でもできるとき、与えられた数は最小でいくらになるか、という問題。

### 解答

交換についてもっと直感的に把握するために、例えば、サンプルの `0709` という数をそれぞれ偶数・奇数というカテゴリでのみ考えると、
`EOEO` というような列になる。

`OO` や `EE` などは交換できないことを考えると、奇数の列だけ考えたとき、それらの相対位置は変えられないことがわかる（偶数列についても同じ）。

なので、まずは奇数と偶数だけを、順番を崩さないようにそれぞれ別にかき集めておく。

あとはこれらを上位桁からどう配置していくかだが、貪欲に上位桁には小さい方を置く方針で良い。
マージソートの要領で、偶数・奇数それぞれを使い尽くすように併合していく。

```go
var t int
var S []rune

func main() {
	t = ReadInt()

	for i := 0; i < t; i++ {
		S = ReadRuneSlice()

		// evens := make([]rune, 0, 300000+5)
		// odds := make([]rune, 0, 300000+5)
		evens := []rune{}
		odds := []rune{}
		for j := 0; j < len(S); j++ {
			if (S[j]-'0')%2 == 0 {
				evens = append(evens, S[j])
			} else {
				odds = append(odds, S[j])
			}
		}

		// answers := make([]rune, 0, 300000+5)
		answers := []rune{}
		e, o := 0, 0
		for e < len(evens) || o < len(odds) {
			if e == len(evens) {
				answers = append(answers, odds[o])
				o++
			} else if o == len(odds) {
				answers = append(answers, evens[e])
				e++
			} else if evens[e] < odds[o] {
				answers = append(answers, evens[e])
				e++
			} else {
				answers = append(answers, odds[o])
				o++
			}
		}

		fmt.Println(string(answers))
	}
}
```

コメントアウトした内容で提出したらなぜかTLEしてしまいました。
可変長配列のメモリ再取得がない分早いと思って書いたのですが、裏目に出ました。
少なくともCodeforcesでは、いつもどおり何も考えずに0容量で確保したほうが良さそうです。

絶対にBよりこっちのほうが簡単だと思うけど、同じようにBにこだわってC以降見れなかった人が多いのかも。

