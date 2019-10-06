初めてCodeforcesに参加しました。

CodeforcesはAtCoderほど日本語の記事が見当たらなかったので、そのあたりモチベーションにしつつブログの練習も兼ねて参加録をつけていこうと思います。

<!-- TOC -->

- [全体](#%e5%85%a8%e4%bd%93)
- [A. Dawid and Bags of Candies](#a-dawid-and-bags-of-candies)
	- [問題](#%e5%95%8f%e9%a1%8c)
	- [解答](#%e8%a7%a3%e7%ad%94)
- [B. Ania and Minimizing](#b-ania-and-minimizing)
	- [問題](#%e5%95%8f%e9%a1%8c-1)
	- [解答](#%e8%a7%a3%e7%ad%94-1)
- [C. Anadi and Domino](#c-anadi-and-domino)
	- [問題](#%e5%95%8f%e9%a1%8c-2)
	- [解答](#%e8%a7%a3%e7%ad%94-2)
- [D. Marcin and Training Camp](#d-marcin-and-training-camp)
	- [問題](#%e5%95%8f%e9%a1%8c-3)
	- [解答](#%e8%a7%a3%e7%ad%94-3)
- [感想](#%e6%84%9f%e6%83%b3)

<!-- /TOC -->

<a id="markdown-全体" name="全体"></a>
## 全体

A, Bの2完でRate 1511スタートでした。
こちらも青色を目指して頑張っていこうと思います。

Cはpretestは通りましたが、ちゃんと全探索しておらずsystem testで撃沈しました。
Dもぱっと浮かんだ解法を後日実装したところ、撃沈していたので、この2問を重点的に復習していきます。

---

<a id="markdown-a-dawid-and-bags-of-candies" name="a-dawid-and-bags-of-candies"></a>
## A. Dawid and Bags of Candies

[問題URL](https://codeforces.com/contest/1230/problem/A)

<a id="markdown-問題" name="問題"></a>
### 問題

Dawidはキャンディの入った4つの袋を持っている。
`i` 番目の袋は `A[i]` 個のキャンディを含んでいる。
Dawidはこれらの袋を二人の友人のいずれかにそれぞれ配る（必ず配らなければならず、自身が保持してはいけない）。

二人の友人に同じ数のアメを配ることは可能か判定せよ。

制約: `1 <= A[i] <= 100`

<a id="markdown-解答" name="解答"></a>
### 解答

それぞれの袋の配り方を全探索したいと考えた。

袋はどちらか一方に所属することになるので、ビット全探索を書いた。

```go
var A []int

func main() {
	A = ReadIntSlice(4)

	for i := 0; i < (1 << uint(4)); i++ {
		one, two := 0, 0
		for j := 0; j < 4; j++ {
			if NthBit(i, j) == 1 {
				one += A[j]
			} else {
				two += A[j]
			}
		}
		if one == two {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
```

<a id="markdown-b-ania-and-minimizing" name="b-ania-and-minimizing"></a>
## B. Ania and Minimizing

[問題URL](https://codeforces.com/contest/1230/problem/B)

<a id="markdown-問題-1" name="問題-1"></a>
### 問題

Aniaは大きな整数 `S` を持っている。
この `S` は、10進数表現で `n` 桁となる（上位桁の桁埋めのための0（leading zeroes）はない）。

Aniaは最大で `k` 個の桁を変えることができる（ただし、leading zeroesを作ってはいけない）。
Aniaができるだけ小さい数を作ろうとするとき、その値は何になるか答えよ。

制約: `1 <= n <= 200000, 0 <= k <= n`

<a id="markdown-解答-1" name="解答-1"></a>
### 解答

`S` がとても大きいので文字列のまま処理する方針で考える。

気持ちとしては上位桁を小さくするほうが得なので、上の桁から0にしていきたいが、条件から「最上位桁を0とするのはご法度」となる。

なので、最上位桁は泣く泣く1とし、それ以降を可能なだけ0で置き換えていくことにする。

シミュレーションする際は、変える必要のない桁を変えて `k` を無駄に減らしてしまったり、ループを抜けるところを間違えて変える桁の過不足が発生しないように注意した。

```go
var n, k int
var S []rune

func main() {
	n, k = ReadInt2()
	S = ReadRuneSlice()

	if k == 0 {
		fmt.Println(string(S))
		return
	}

	if n == 1 {
		S[0] = '0'
		k--
	} else {
		if S[0] != '1' {
			S[0] = '1'
			k--
		}
	}

	if k == 0 {
		fmt.Println(string(S))
		return
	}

	// この時点で k > 0
	for i := 1; i < n; i++ {
		if S[i] != '0' {
			S[i] = '0'
			k--
		}

		if k == 0 {
			break
		}
	}

	fmt.Println(string(S))
}
```

leading zeroes というのはあまり馴染みがなかったので覚えておきたいところ。

<a id="markdown-c-anadi-and-domino" name="c-anadi-and-domino"></a>
## C. Anadi and Domino

[問題URL](https://codeforces.com/contest/1230/problem/C)

<a id="markdown-問題-2" name="問題-2"></a>
### 問題

Anadiはドミノの1セットを持っている。
それぞれのドミノは、両サイドに1から6のサイコロの面が描かれており、それぞれの種類について1ピースあるため、全部で21個のドミノがある。

※URLの画像が親切なので詳しくはそちらをご参照ください。

また、Anadiは自己ループ・多重辺の無い無向グラフを持っており、そのグラフの辺にドミノを置こうとしている。
ドミノを置く際のルールは以下のようになっている。

- ドミノの両サイドが、辺を結ぶノードを向くようにする。
- 複数のドミノがあるノードを指すように置かれる場合、そのノードが刺されるドミノの数字は等しくなければならない。

Anadiは最大で何個のドミノを置けるか答えよ。

制約: `1 <= n <= 7, 1 <= m <= n*(n-1)/2` （それぞれノード、辺の数）

<a id="markdown-解答-2" name="解答-2"></a>
### 解答

よくわからないまま、サンプル4のノード数7の完全グラフのケースについて、色々とお絵かきをしながら考えた。

考えるうちに、ノードが6個以下の場合は、すべての辺に対してなんらかのドミノが置けることがわかった（1からnの番号を素直にノードに振ってやった後、その番号に向くようにドミノを選んで置いていけば良い）。

なので、 `n <= 6` のケースでは `m` をそのまま出力すれば良い。

ノード数が7のときだけは、2つのノードは互いにドミノの数字をかぶらせる必要が出てくる。

**ここで、ドミノの数字をかぶらせるノードペアが同じであれば、それぞれのノードに向けるドミノの数字のパターンが変わろうとも、置けるドミノの数は変わらない、事がわかる**
（使うドミノが数字に合わせて変わるだけで、使いたいドミノがセットの中にない、ということは起こらない。）

なので、ノード数が7の場合は、「ドミノの数をかぶらせるノードのペアの選び方」を全探索する。

ノード間の辺の有無や、ある種類のドミノの使用済みかどうかの判断はそれぞれともに行列を利用して行っている。

```go
var n, m int

var adjMatrix [10][10]int
var dominoMatrix [10][10]int

func main() {
	n, m = ReadInt2()
	for i := 0; i < m; i++ {
		a, b := ReadInt2()
		// 常にaをbより小さくする
		if a > b {
			a, b = b, a
		}
		adjMatrix[a][b] = 1
		adjMatrix[b][a] = 1
	}

	if n <= 6 {
		fmt.Println(m)
	} else {
		fmt.Println(sub())
	}
}

func sub() int {
	res := 0

	// i, jを同じ数字とする
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			// ノードにサイコロの目を割り振る
			memo := make(map[int]int)
			dice := 1
			for k := 1; k <= n; k++ {
				if k == j {
					continue
				}
				memo[k] = dice
				dice++
			}
			memo[j] = memo[i]

			tmp := 0
			initialize()
			for l := 1; l <= n; l++ {
				for m := l + 1; m <= n; m++ {
					ll, mm := memo[l], memo[m]
					if ll > mm {
						ll, mm = mm, ll
					}
					if adjMatrix[l][m] == 1 && dominoMatrix[ll][mm] == 1 {
						tmp++
						dominoMatrix[ll][mm] = 0
					}
				}
			}

			ChMax(&res, tmp)
		}
	}

	return res
}

func initialize() {
	for i := 1; i <= 6; i++ {
		dominoMatrix[i][i] = 1
		for j := i + 1; j <= 6; j++ {
			dominoMatrix[i][j] = 1
		}
	}
}
```

コンテスト中は、「数字をかぶらせる必要のあるノード番号を7とし、
7を割り当てるノードを全部調べる」というようなことを考えていたら、
見るからにバグりそうなコードが出来上がってしまった（実際にsystem testで撃沈した）。

ヤバそうなコードを書いているときは、たいてい予想通りどこかでひっかかるので、
コンテスト中も途中で考えを捨てる勇気を持ちたい。

とはいえ、改めて整理して書いてみても、それなりに複雑なコードになってしまった。

ごちゃごちゃ考えずに、`n` の値によらずに全探索する方法もあるらしい。
そちらのほうが確実そうな一方で、実装方法がよくわかっていない（のでどなたか教えて下さい）。

<a id="markdown-d-marcin-and-training-camp" name="d-marcin-and-training-camp"></a>
## D. Marcin and Training Camp

[問題URL](https://codeforces.com/contest/1230/problem/D)

<a id="markdown-問題-3" name="問題-3"></a>
### 問題

Marcinは大学の講師で、n人の合宿に参加したい生徒を抱えている。
Marcinは、互いにcalmlyに共同作業できる生徒たちを送ろうとしているとしている。

生徒たちは1からn番の番号が振られており、それぞれ `A[i], B[i]` で特徴づけられている。

- `A[i]`: このビット表現は、0から59のアルゴリズムを、`i` 番目の生徒が知っているかどうかを示している（1ならば知っている）。
- `B[i]`: `i` 番目の生徒のスキルの高さを示している（高ければ高いほどよい）。

生徒たちは、「あるアルゴリズムについて、自分は知っているが相手は知らない」というようなアルゴリズムが1つでもある場合に、相手のことを見下す。
2人以上のグループを組んだとき、相手のことを見下す生徒が一人もいない場合に、そのグループはcalmlyに共同作業できる。

共同作業できるグループを作る際、そのグループを構成する生徒のスキルレベルの和を最大化する場合、スキルレベルの和はいくらになるか答えよ。

<a id="markdown-解答-3" name="解答-3"></a>
### 解答

グリーディに考える。

グループが作れる条件を、「少なくとも1組のペアが、互いに知っているアルゴリズムの集合が同じで、互いを対等だと思っている」ことだと考える。

そこで、まずは対等な人同士でグループを作ってしまう。
対等な人のグループは、Aでソートすれば見つかる。

また、この対等な人たちでのみ構成されるグループに対して、まだグループが組めていない生徒を組み込んでいくことを考える。
具体的には、グループを表すビット集合と、ある生徒のビット集合のビットOR演算をとり、ビット集合がグループのビット集合のままであれば、
その生徒はグループに組み込むことができる。
逆に、グループのビット集合が変化してしまう場合、その生徒はそのグループの人全員を見下してしまうため、そのグループには組み込めない。

最後に、出来上がった全てのグループを併合しても問題ないため、グループに所属している生徒のスキルレベルの和を計算する。

和を計算するときは、32ビット整数の範囲を超えてしまうことに注意。

ソート部分の計算量が `O(n * log(n))` で、対等な人たちのグループに残った人の組み込みを考える部分の計算量が `O(n^2)` となる。

```go
var n int
var A []int64
var B []int

type Student struct {
	key int64
	a   int64
	b   int
	idx int
}
type StudentList []*Student
type byKey struct {
	StudentList
}

func (l StudentList) Len() int {
	return len(l)
}
func (l StudentList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.StudentList[i].key < l.StudentList[j].key
}

var flags []bool

func main() {
	n = ReadInt()
	A = ReadInt64Slice(n)
	B = ReadIntSlice(n)
	flags = make([]bool, n)

	if n == 1 {
		fmt.Println(0)
		return
	}

	L := make(StudentList, 0, 200000)
	for i := 0; i < n; i++ {
		L = append(L, &Student{key: A[i], a: A[i], b: B[i], idx: i})
	}
	sort.Stable(byKey{L})

	// 2以上のサイズのメモ
	memo := make(map[int64]int)
	for i := 0; i < len(L); i++ {
		if i == 0 {
			if L[i].a == L[i+1].a {
				memo[L[i].a] = 1
			}
		} else if i == len(L)-1 {
			if L[i-1].a == L[i].a {
				memo[L[i].a] = 1
			}
		} else {
			if L[i-1].a == L[i].a || L[i].a == L[i+1].a {
				memo[L[i].a] = 1
			}
		}
	}

	for bits := range memo {
		for i := 0; i < n; i++ {
			if bits|A[i] == bits {
				flags[i] = true
			}
		}
	}

	sum := int64(0)
	for i := 0; i < n; i++ {
		if flags[i] {
			sum += int64(B[i])
		}
	}

	fmt.Println(sum)
}
```

対等な人たちのグループを併合してから、そのグループのアルゴリズムセットが包含するような生徒を取り込む、というようにすると誤りになる
（もとの個々のグループの人全員を見下してしまう生徒を組み込んでしまう場合がある）。

やむを得ず貪欲法を選択することになった場合は、ちゃんと整理しきってから提出したい。

これもどうやら全探索的なやり方があるらしいので、コンテスト中はできる限りそういったものが選択できるようにしたい。

---

<a id="markdown-感想" name="感想"></a>
## 感想

細かい部分でなれない実装が要求される感じがあり、本質的でない部分で消耗してしまった気がします。
最近のABCのE問題以降では、これぐらいの実装が要求されるものも少なくないように感じるため、慣れていきたいところです。

また、無証明の貪欲は危険だということも体験できたので、証明をスムーズにできる訓練もしていきたいと思います。
