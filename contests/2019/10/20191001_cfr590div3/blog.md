# Codeforces Round No.590 (Div.3) 参加記録（A〜E解答）

## 全体

残り時間15分ぐらいで後回しにしたCが解けて5完でした。

Cに時間がかかりすぎてしまったのは悔やまれますが、実力的に何回参加しても好転はしないと思います。

Eはコンテスト中に目を通せて「面白そう」と思ったので、後日解き直しました（解説ACになりましたが）。

---

## A. Equalize Prices Again

[問題URL](https://codeforces.com/contest/1234/problem/A)

### 問題

※やたらと長いので意訳。

各値段が `A[i]` コインである `n` 個の商品のすべての値段を等しくしたい。
もともとの商品の値段合計よりも小さくならないという条件のもとで、金額はいくらに設定すればよいか。

このようなクエリが `q` 個与えられるので、それぞれについて答えよ。

制約: `1 <= q <= 100, 1 <= n[i] <= 100, 1 <= A[i] <= 10^7`

### 解答

各クエリの情報を読み込みながら、都度 `Floor(sum / n)` を計算すれば良い。

```go
var q int

func main() {
	q = ReadInt()
	for i := 0; i < q; i++ {
		n := ReadInt()
		A := ReadIntSlice(n)

		sum := Sum(A...)
		fmt.Println((sum + (n - 1)) / n)
	}
}
```

特に言うことはないと思います。

## B1. Social Network (easy version)

[問題URL](https://codeforces.com/contest/1234/problem/B1)

### 問題

※問題文があまりにも冗長なので省略。

### 解答

要するに、以下を守りながら、ディスプレイを配列でモデル化し、シミュレーションを行えば良い。

- 次に処理するメッセージがすでにディスプレイに表示されているのならば、何もしない。
- 次に処理するメッセージが表示されていない場合は、そのメッセージを配列の先頭に挿入する。
- ディスプレイサイズが `k` を超える場合、配列の最後尾のメッセージを配列から除外する。

最後に、シミュレーション終了時の配列を出力すればOK。

なお、easy versionは制約が小さいので、配列の調整やメッセージの存在チェックを適当に都度スキャンしても通る。
実際に、コンテスト中は完全なシミュレーションに集中したので、効率化を一切考えないコードを書いた。

```go
var n, k int
var A []int
var Ids []int

func main() {
	n, k = ReadInt2()
	A = ReadIntSlice(n)

	Ids = []int{}
	for i := 0; i < n; i++ {
		a := A[i]

		if isExist(a) {
			continue
		}

		newIds := []int{}
		newIds = append(newIds, a)
		if len(Ids) < k {
			newIds = append(newIds, Ids...)
		} else {
			newIds = append(newIds, Ids[:len(Ids)-1]...)
		}
		Ids = newIds
	}

	fmt.Println(len(Ids))
	fmt.Println(PrintIntsLine(Ids...))
}

func isExist(id int) bool {
	for i := 0; i < len(Ids); i++ {
		if id == Ids[i] {
			return true
		}
	}
	return false
}
```

正直英文読解が最大の敵だと思いました
（サンプル見るまで確信が持てませんでしたし、なんならサンプル見るまで勘違いしている部分もありました）。

## B2. Social Network (hard version)

[問題URL](https://codeforces.com/contest/1234/problem/B2)

### 問題

※easy versionと問題設定は全く同じ。

制約: `1 <= n, k <= 2 * 10^5, 1 <= Id[i] <= 10^9`

### 解答

総メッセージ数とディスプレイサイズが大きくなったため、easy versionのコードは通らなくなった。
具体的には、

- 存在チェックの際、ディスプレイをスキャンしている
- メッセージの出し入れの部分を、新しい配列に順番が正しくなるようにすべて詰め直している

部分がまずい。

それぞれ、

- 存在チェックを `map` で行う
- ディスプレイの向きを逆向きにすることでQueueでディスプレイをモデル化できるため、出し入れが `O(1)` で可能になる。

というふうに変更する。

```go
var n, k int
var A []int
var Ids []int
var memo map[int]int

func main() {
	n, k = ReadInt2()
	A = ReadIntSlice(n)

	memo = make(map[int]int)
	for i := 0; i < n; i++ {
		memo[A[i]] = 0
	}

	Ids = []int{}
	for i := 0; i < n; i++ {
		a := A[i]

		if isExist(a) {
			continue
		}

		// 追加
		memo[a] = 1

		Ids = append(Ids, a)
		if len(Ids) <= k {
			continue
		} else {
			memo[Ids[0]] = 0
			Ids = Ids[1:]
		}
	}

	fmt.Println(len(Ids))
	reverseIds := rev()
	fmt.Println(PrintIntsLine(reverseIds...))
}

func isExist(id int) bool {
	if value, ok := memo[id]; ok {
		if value == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func rev() []int {
	res := []int{}

	for i := len(Ids) - 1; i >= 0; i-- {
		res = append(res, Ids[i])
	}

	return res
}
```

## C. Pipes

[問題URL](https://codeforces.com/contest/1234/problem/C)

### 問題

2行n列のマス目に対して、各マスに6種類のパイプが設置されている。
各マスのパイプを時計回り・反時計回りに90度単位で何回も回転させて良い。
左上から流入する水を、右下に流出させることが可能か答えよ。

また、このようなクエリは `q` 個与えられ、それぞれについて `n` およびパイプの配置が与えられるので、
それら全てに答えよ。

※パイプの種類や水の流し方に関する制約は、問題文の図が詳しいのでそちらをご参照ください。

制約: `1 <= q <= 10^4, 1 <= n <= 2 * 10^5(n の総和は 2 * 10^5 を超えないことが保証される)`

### 解答

行列の列の総和がboundされているため、素直にクエリごとに各マスのパイプをチェックして問題ない。

整理すると、問題がシンプルになる。

- パイプ1, 2は互いに同じで、パイプ3, 4, 5, 6はそれぞれ互いに同じ
- 水を左方向に流す意味はない
- 3, 4, 5, 6のパイプを経由して行方向を上・下に移動した後に1, 2のパイプに出会った場合は、確実に失敗する
  - それぞれ行列を下・上に踏み外すか、水が流れなくなるかのいずれかのため

以上から、初期位置を `(0, 0)` とし、出会ったパイプの種類によって座標を変化させていけば良い。

```go
var q int
var n int
var Rows [2][]rune

func main() {
	q = ReadInt()

	for i := 0; i < q; i++ {
		n = ReadInt()
		Rows[0] = ReadRuneSlice()
		Rows[1] = ReadRuneSlice()
		// 現在地
		h, w := 0, 0

		for w < n {
			if Rows[h][w] == '1' || Rows[h][w] == '2' {
				w++
			} else {
				if h == 0 && (Rows[1][w] == '1' || Rows[1][w] == '2') {
					fmt.Println("NO")
					break
				}
				if h == 1 && (Rows[0][w] == '1' || Rows[0][w] == '2') {
					fmt.Println("NO")
					break
				}

				if h == 0 {
					h++
					w++
				} else {
					h--
					w++
				}
			}

			if w == n {
				if h == 0 {
					fmt.Println("NO")
				} else {
					fmt.Println("YES")
				}
			}
		}
	}
}
```

It is guaranteed that the sum of `n` over all queries does not exceed `2 * 10^5`
を見逃したため、無理だと思って先にDを解きました。

実装で手間取りすぎてこの問題だけで1時間弱使ってしまった結果を見ると、正しい動きでした。

## D. Distinct Characters Queries

[問題URL](https://codeforces.com/contest/1234/problem/D)

### 問題

英小文字のみからなる文字列 `s` と、この文字列に関する `q` 個のクエリが与えられる。

`s[l;r]` を `s` の部分文字列 `s[l], s[l+1], ..., s[r]` とする。

クエリは2つのタイプからなる。

1. `1 pos c`: `s[pos]` を `c` で置換する
2. `2 l r`: 部分文字列 `s[l;r]` 内の異なる文字の個数を計算する

各クエリに答えよ。

制約: `1 <= |s| <= 10^5, 1 <= q <= 10^5`

### 解答

26種類の文字について累積和が計算できれば、与えられた区間内の各文字の個数は計算できるので、
1個以上存在する文字の種類数が答えになる。

しかしながら、あるインデックスについて文字の置換が途中で割り込むため、通常の累積和では再計算していては間に合わない。

BITが要件を満たしており、26個に増やすのも容易に思えたため、BITを採用した。

```go
const ALPHABET_NUM = 26

var S []rune
var q int

// [1, n]
var bit [ALPHABET_NUM][100000 + 5]int
var lenS int

func sum(i, alpha int) int {
	s := 0

	for i > 0 {
		s += bit[alpha][i]
		i -= i & (-i)
	}

	return s
}

func add(i, alpha, x int) {
	for i <= lenS {
		bit[alpha][i] += x
		i += i & (-i)
	}
}

func main() {
	S = ReadRuneSlice()
	q = ReadInt()
	lenS = len(S)

	for i := 0; i < len(S); i++ {
		r := S[i]
		c := int(r - 'a')
		add(i+1, c, 1)
	}

	for i := 0; i < q; i++ {
		query := ReadInt()
		if query == 1 {
			idx := ReadInt()
			R := ReadRuneSlice()
			newc := R[0]
			newcint := int(newc - 'a')
			oldc := S[idx-1]
			oldcint := int(oldc - 'a')

			add(idx, newcint, 1)
			add(idx, oldcint, -1)
			S[idx-1] = newc
		} else {
			l, r := ReadInt2()
			res := 0
			for alpha := 0; alpha < ALPHABET_NUM; alpha++ {
				ss := sum(r, alpha) - sum(l-1, alpha)
				if ss > 0 {
					res++
				}
			}
			fmt.Println(res)
		}
	}
}
```

26個の累積和を計算し終わった後に点更新できない事に気づき、紙の蟻本を取りに行きました。

あと提出し終わった後、長さ0の文字列がケースに入っていないかとても心配でした
（プログラムがコケるから無いだろうと決め込みましたが。。）。

## E. Special Permutations

[問題URL](https://codeforces.com/contest/1234/problem/E)

### 問題

数列 `P(i, n)` を `[i, 1, 2, ..., i-1, i+1, ..., n]` のように定義する。
すなわち、 `n` の順列とほとんど同じだが、 `i` のみが先頭に移動したものである。

また、数列 `X: X[1], X[2], ..., X[m] (1 <= X[i] <= n)` が与えられる。

`pos(P, val)` は数列Pが与えられたときの、数列P中における `val` が位置する1-basedのインデックスである。

ここで、関数 $f(P) = \sum_{i=1}^{m-1} |pos(P, X[i]) - pos(P, X[i+1])|$ を定義する。

$f(P_{1}^{n}), f(P_{2}^{n}), ..., f(P_{n}^{n})$ を求めよ。

### 解答

問題文でも言われている通り、先頭に何が移動しようが、もとの昇順に並んだ数列 `P(1, n)` とほとんど同じであることを踏まえると、
関数 `f` の値もそれぞれ似通いそう、すなわち `f(P(1, n), n)` を事前にもとめておいて、
それを適切に調整すればうまく求められそうな気はする。

実際、以下の図のように場合分けによって、項間の距離が計算できる。



手順としては、以下のようなものとなる。

1. 昇順に並んだ数列について関数値を求め、これをベースとする。
  - これは、単純に `X` の階差の和で求まる。
2. 各 `X[i]` について `X[i]` が先頭に来る場合の `X[i-1]` と `X[i+1]` との距離（ `A[i]` とする）を計算しておく。
  - 先述の図示した内容に従って場合分けすれば良い。
3. `X` 中における `1 <= i <= n` な `i` が登場する箇所をリスト形式（ `C[i]` とする）で記憶しておく。
  - 後半でベースをもとに `f(P(i, n), n)` を計算する際、リスト `C[i]` の中身についてのみ、2で求めた値を参照する。
  - もともとベースを計算するために用いた階差が、2で求めた値に置き換わることとなる。
4. すべての `X` に置ける隣接項のペア（公式tutorialでは `segment` というワードが使われている）について、 `1 <= i <= n` な `i` が間に挟まるものの数を数える。
  - `cnt[i]` とすると、これが求まればベースから `cnt[i]` を引くことで、「 `i` が先頭に来ることによってグループ1とグループ2の間の距離が1縮まる」という部分が、すべての `X` の隣接項ペアについて考慮できたことになる。
  - 普通にやると `O(m * n)` なので、imos法により `O(m + n)` で賢く行う。
5. これまでに前処理したデータを利用し、ベースを調整することで各 `f(P(i, n), n)` を求める。
  - おおまかな方針は以下のようになる。
    - `X[idx]` が先頭に移動する場合は、ベースを求めるのに使った階差の代わりに `A[idx]` を使うようにする。
    - `cnt[i]` をベースから引く。

コードの最後で2重ループとなっているが、 `C` の要素数は全部で `m` であるため、
計算量は `O(n + m)` となって十分間に合う。

```go
var n, m int
var X []int

var A, diff []int
var C [][]int
var cnt []int

func main() {
	n, m = ReadInt2()
	X = ReadIntSlice(m)

	// Xの階差を計算する、和がf(P1(n))となる
	diff = make([]int, m-1)
	for i := 0; i < m-1; i++ {
		diff[i] = AbsInt(X[i] - X[i+1])
	}

	// X[i]が数列の先頭に来る場合の、前後のxとの位置の差の絶対値を計算しておく
	A = make([]int, m)
	for i := 0; i < m; i++ {
		if i < m-1 {
			if X[i+1] > X[i] {
				A[i] += X[i+1] - 1
			} else if X[i+1] < X[i] {
				A[i] += X[i+1]
			}
		}

		if i > 0 {
			if X[i-1] > X[i] {
				A[i] += X[i-1] - 1
			} else if X[i-1] < X[i] {
				A[i] += X[i-1]
			}
		}
	}

	// X中における、1<=i<=nのiが登場する位置を記憶しておく
	C = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		C[i] = []int{}
	}
	for i := 0; i < m; i++ {
		C[X[i]] = append(C[X[i]], i)
	}

	// cnt[i]: 1<=i<=nのiについて、
	// Xにおけるすべての隣接項間のうち、隣接項間にiが挟まるものの個数
	cnt = make([]int, n+1)
	for i := 0; i < m-1; i++ {
		if AbsInt(X[i]-X[i+1]) < 2 {
			continue
		} else {
			cnt[Min(X[i], X[i+1])+1]++
			cnt[Max(X[i], X[i+1])]--
		}
	}
	for i := 0; i < n; i++ {
		cnt[i+1] = cnt[i+1] + cnt[i]
	}

	base := int64(0)
	for _, d := range diff {
		base += int64(d)
	}
	answers := []int64{}
	for i := 1; i <= n; i++ {
		ans := base
		for _, idx := range C[i] {
			if idx == 0 {
				ans += int64(A[idx] - diff[idx])
			} else if idx < m-1 {
				ans += int64(A[idx] - diff[idx] - diff[idx-1])
			} else {
				ans += int64(A[idx] - diff[idx-1])
			}
		}
		ans -= int64(cnt[i])
		answers = append(answers, ans)
	}

	fmt.Println(PrintIntsLine(answers...))
}

// PrintIntsLine returns integers string delimited by a space.
func PrintIntsLine(A ...int64) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		// str := strconv.Itoa(A[i])
		str := strconv.FormatInt(A[i], 10)
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}
```

図示したところの、「グループ1とグループ2の要素間はもともとの距離から1縮まる」の部分が、
なぜか自力では最後の方まで抜けてしまっていました。。
気づいてもimos法の適用までには至れず、解説ACとなりました。

個々に要求される知識や実装はそれほど難しいわけではないですが、
しっかり落ち着いて解ききるにはまだ地の力足りてないなぁ、という気がします。

筋トレしましょう。

---

## 感想

Codeforcesってクエリ問題多いんですかね？

今回Cでやらかした制約の見落としなんかはこれからもありそうなので気をつけたいと思います。

