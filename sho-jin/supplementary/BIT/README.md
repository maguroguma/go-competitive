# BITの活用

## 参考

- [けんちょんさんのQiita記事](https://qiita.com/drken/items/1b7e6e459c24a83bb7fd)
- [hosさんのスライド](http://hos.ac/slides/20140319_bit.pdf)
- [yosupoさんのブログ記事](https://yosupo.hatenablog.com/entry/2014/12/09/222401)
- [kazumaさんのブログ記事](http://kazuma8128.hatenablog.com/entry/2018/06/20/210631)

## 例題

- [ARC033 C.データ構造](https://arc033.contest.atcoder.jp/tasks/arc033_3)
  - シンプルで座圧なども不要のため、verify用によい。
- [yukicoder649 ここでちょっとQK!](https://yukicoder.me/problems/no/649)
  - 座圧が要求されるため、そちらを含めてverifyするのによい。
- [CFR602 D2. Optimal Subsequences (Hard Version)](https://codeforces.com/contest/1262/problem/D2)
  - 位置が重要になるため座圧は不要だが、クエリ先読みやもろもろ構造体のソートが必要になるため、整理するのが面倒。
  - 書いてみると意外とシンプルにはなる。
- [ARC028 B.特別賞](https://atcoder.jp/contests/arc028/tasks/arc028_2)
  - とてもシンプルなので最初の練習に良い。もしくはverify用に。
- [Chokudai SpeedRun 001 J.転倒数](https://atcoder.jp/contests/chokudai_S001/tasks/chokudai_S001_j)
  - 長らく蟻本の説明をうまく噛み砕けていなかったが、数列を左からBITの集合に追加していく、と考えると自然に理解できる。

## BIT上の二分探索

BIT上の二分探索と言われると「どういうこと？」となってしまうが、
要は **元の配列の先頭からの累積和に関して二分探索** しているということ。

別の表現をすると **「1-basedなインデックス `i` で、 `[1, i]` の要素の総和が `w` 以上となる最小の `i` を探索する」** ということ。

※累積和に関して、二分探索の前提となる単調性が必要であるため、暗黙的に各要素は非負であることを仮定する。
とはいえ、以下に示すOrdered Setとしての活用がメインとなると思うので、あまり気にしなくて良い。

### 何が嬉しいか？

条件はあるが **Ordered Setのような使い方ができる。**

これは、BITに適用する元配列を、 `[1, D]` までの値の要素数とすることで、二分探索のキーを `k` としたときに `k` 番目の値が求められることによる。

よって、管理する値の最大値分のメモリが必要となる。
最大値が `10^9` のようなときでも、取りうる最大値がわかっているのであれば、座標圧縮と併用することで、同様に利用できる。

[けんちょんさんのQiita記事](https://qiita.com/drken/items/1b7e6e459c24a83bb7fd)で詳しく整理されている。

※BITは基本的に1-basedであることを念頭において考えること。

※残念ながら、具体的な格納されている値に関する検索はできない（多分）。

### アルゴリズム、ロジック、コード

Goだと以下のような感じになる。

```go
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
```

何をやっているかというと、
**先頭からの累積和を、できるだけ長い区間から足してよいかを都度判断する、**
ということをしている。

流れとしてはBITの上方の区間が長いノードから順番に見ており、
`k` というのは現在注目している区間の長さを意味している。

各区間は、担当する区間の和を持っており、それがkey値よりも小さい場合は足してから右側の次の短い長さの区間を見る必要がある。
逆にkey値よりも大きい場合は、その区間和は足さずに左側の次の短い長さの区間を見る必要がある。

if文の条件の細部がちょっと難しいが、これについてはあまり気にしないほうがいい（実験してみると正しいことは確認できる）。
