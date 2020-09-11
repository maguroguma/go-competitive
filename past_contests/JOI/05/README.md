# 日本情報オリンピック過去問

Last Change: 2020-09-12 02:35:32.

難しかったものや癖のあるものだけメモしていく。

## [E. 品質検査 ★](https://atcoder.jp/contests/joi2007yo/tasks/joi2007yo_e)

なんとなく難しかったが、翌々考えたら解答PDFのにある通り、ほとんどの場合が不明としてしまって良い。

これは、「不明なものはすべて異常としてしまう」や「1つだけ正常にして後はすべて異常としてしまう」が
trueとなりえることを考えるとわかる。

「極端な状況を考える」というのは重要かもしれない。
最近は見ない類の問題だが、論理として重要なので覚えておきたい。

## [B. 最長の階段](https://atcoder.jp/contests/joi2007ho/tasks/joi2007ho_b)

くる式Union Find活用法によってきれいな実装ができる。

## [B. factorial ★](https://atcoder.jp/contests/joisc2007/tasks/joisc2007_factor)

個人的に苦手というか、多分ABCで出たら茶から緑ぐらいなんだけど自分は遅くて悲しい思いをする、というタイプ。

[らてさんと同じ解き方](https://lattemalta.hatenablog.jp/entry/2015/09/16/032833)をしていると思うが、
多分二分探索は不要な解法がある気がする。
とはいえ、このような素数の数え方（？）は過去のABCでも見た気がするので、もっとスムーズにできるようにしたい。

**`n!` の中に含まれる素因数 `2` の個数は、 `n/2 + n/4 + n/8 + ...` で数えられる。**

## [C. mall](https://atcoder.jp/contests/joisc2007/tasks/joisc2007_mall)

2次元累積和をやるだけ。

なんで急にTL6secとかになったんだろう。。？

## [D. building](https://atcoder.jp/contests/joisc2007/tasks/joisc2007_buildi)

LISを求めるだけ。
しかも制約的に `O(N^2)` の愚直なDPが通る。

## [G. Fiber ★](https://atcoder.jp/contests/joisc2007/tasks/joisc2007_fiber)

Union Findで連結成分の数を求めて、それを `x` とすると、答えは `x-1` になる。

簡単ではあるが、木グラフのように `(森の数) - 1` みたいなことをやる難しめの問題もあったので、
一応見逃しやすい典型として覚えておきたい。

## [B. 共通部分文字列 ★★](https://atcoder.jp/contests/joi2008ho/tasks/joi2008ho_b)

この問題はよく覚えておいたほうがいのかもしれない。

ABC141-EのようなDPを練習のつもりでやってみたが、メモリ制約が厳しいため `int16` を用いる必要がある。

あるいは、ローリングハッシュでも解けると思われる。

一方で、解説PDFの手法はかなり特殊。
役に立つかはわからないが、一度なぞってみると良いかもしれない。

## [A. IOIOI ★★](https://atcoder.jp/contests/joi2009ho/tasks/joi2009ho_a)

これも解くだけなら適当にローリングハッシュを使えば良い。
KMP法などの練習にもいいらしいので、それの練習としてみたい。

## [D. カード並べ](https://atcoder.jp/contests/joi2010yo/tasks/joi2010yo_d)

`P(n, k)` の文字列型のpermutationパターンを全列挙するライブラリを使えば良い。

