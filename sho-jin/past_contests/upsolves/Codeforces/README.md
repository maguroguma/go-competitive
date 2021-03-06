# Codeforcesのupsolveした問題たち

Last Change: 2020-12-30 21:17:22.

## [Raif Round 1 E.Carrots for Rabbits](https://codeforces.com/contest/1428/problem/E)

解説は画像のとおりだが、本問題における `F(l, p)` という関数と、この問題で使用する性質は他の問題でも応用できそうなので、
丸暗記してしまってもいいのかもしれない。

![](./raifround-e.jpg)

追記: 肝となる不等式の証明も追加。

![](./raifround-e-proof.jpg)

## [CR681 div2 D.Extreme Subtraction](https://codeforces.com/contest/1443/problem/D)

Editorialがちょっと読みづらいが、結局やりたいことは当初自分が考えていた内容と近いものだった。  
おそらく、自分の解法は、増加列なるべき配列Bが、正しく増加列とならないような実装になっていた。

もう少し自分の方針を具体化すべきだったかもしれない。

![](./cr681div2D.jpg)

## [CR687 div2 D.XOR-gun](https://codeforces.com/contest/1457/problem/D)

最上位ビットに注目すると、最上位ビットが同じものが3つ連続している連続部分列があると、答えは必ず1となる。  
なぜなら、3つのうち後ろ2つをXORすると、最上位ビットが打ち消されて、必ず残った1つがXORされた結果を上回るからである。

。。ここはコンテスト中に気づけたのに、そこから次の単純かつ強力な事実に気づけなかった。。  
今回の制約ではせいぜい32ビット考えれば十分で、よくよく考えると `n` が大きければ上述のような3つの連続部分列は存在せざるを得ない。  
これは、鳩の巣原理的に考えれば自然とわかる。

参考: [高校数学の美しい物語: 鳩の巣原理を使う数学オリンピックの問題](https://mathtrain.jp/pigeon)

よって、ここでは `n < = 100` という制約のもとで解ければ十分となる。  
結果の作り方としては以下の2パターンで十分。

- ある1つの連続部分列をXORして、もとの数列のその前後との大小関係を比較する
- 2つの連続部分列で、それらが隣接している場合、その2つの連続部分列のXOR同士の大小関係を比較する

巧くやれば実装もきれいになる。

**この手の「実はよく考えると巨大な制約は考えなくて良くなる」というのは、最近のAtCoderでもちらほら見かける気がするので、**
**考察の1つの手札として持っておきたい。**

## [CR688 div2 C.Triangles](https://codeforces.com/contest/1453/problem/C)

やりたいことはぼんやりわかるが、かなりやりたくない気持ちになる実装系の問題。  
実際は「底辺（行・列に並行な三角形の辺）を作るためにしか、あるマスの変更は考えなくて良い」という事に気づけると、
かなり前向きになれる。  
この性質は、仮に底辺に含まれない他の唯一の点を作るためにセルの桁を変更すると仮定すると、
そのような変更するセルは、面積を変えずに同じく行・列のいずれかに平行にできることから主張できる（等積変形というやつ？）。

実装としては、以下をかき集めることが大事になると思う。

- 各数字について行ごとの最小・最大の番号
- 各数字について列ごとの最小・最大の番号

`chmax, chmin` 関数を豊富に使ってできるだけ楽をしたい。  
といっても、後半部分の面積計算部分は頭がこんがらがってしまう。。

## [ECR101 C.Building a Fence](https://codeforces.com/contest/1469/problem/C)

本番では想定とは異なった、やや正当性が疑わしい方法（貪欲だがおそらく嘘ではない、と思う。。）で通したので、
改めて、想定解法の区間がテーマのやり方で解き直した。

結局の所、左から順に「現在注目している区画のフェンスのおける位置の範囲」というのが特定できるので、それを愚直に実装してやれば良い。  
具体的には、直前の動かせる範囲が区間 `[cl, cr]` であるとき、現在注目中の区画の動ける範囲は `[H[i], H[i]+(k-1)]` のようになるため、
これら2つの区間の共通部分が空集合にならなければよい。  
この共通部分は `[max(cl, nl), min(cr, nr)]` で計算でき、これが `l > r` となったならば `NO` になる。

区間の扱い方がもっとスムーズにできるようにならないと、また痛い目を見てしまう。。  
[ACLBC-B](https://atcoder.jp/contests/abl/tasks/abl_b)も思い出そう。

### 区間の交差判定について

以前は[きゅうりさんのこのツイート](https://twitter.com/kyuridenamida/status/338366544776146944)を調べて急いでコードを書いたが、
`[max(cl, nl), min(cr, nr)]` が矛盾しないことを考えたほうが良さそう。

[りんごさんのコード](https://atcoder.jp/contests/abl/submissions/17016923)でもそのように書いている。

## [ECR101 D.Ceil Divisions](https://codeforces.com/contest/1469/problem/D)

構築だが、想定解法が面白かった。

以下のようにして解くと、 `n+5` 回という厳しい回数に十分足りる。

1. 現在の最大値 `x` に対して `y >= Kiriage(x, y)` を満たす **最小の** `y` を見つける。
2. `[y+1, x-1]` までを `x` で除算してすべて1にする。
3. `x` を `y` で2回除算して1にする
4. `x <- y` としてstep1に戻る。

`y >= Kiriage(x, y)` というのは `x` の平方数に対して天井関数を適用したものに該当する。  
これは最大値に対して平方根を再帰的に繰り返すような形になり、かなりの速さで小さくなっていくため、制約の回数に足りることになる。  

```python
>>> math.sqrt(200000)
447.21359549995793
>>> math.sqrt(448)
21.166010488516726
>>> math.sqrt(22)
4.69041575982343
>>> math.sqrt(5)
2.23606797749979
>>> math.sqrt(3)
1.7320508075688772
```

よって逆に、 **繰り返し自身に自身をかけていくことを繰り返すと、凄まじいスピード（指数よりもあるかに早い）で大きくなっていく** ということも意味する。  
あまりこのようなイメージが無かったので、覚えておきたい。

