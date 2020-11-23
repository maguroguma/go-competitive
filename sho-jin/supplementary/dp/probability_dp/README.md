# 確率DP

Last Change: 2020-11-23 23:16:33.

期待値DPが一番の勉強目的だったが、確率DPや期待値の線形性といったトピックも一緒に学ぶ。

## 参考

- [確率DPを極めよう](https://compro.tsutaj.com//archive/180220_probability_dp.pdf)
- [けんちょん氏のABC078-C HSIの解説](https://drken1215.hatenablog.com/entry/2019/03/23/175300)
- [「〜がすべて終わるまでの試行回数の期待値」を求める一般的なフレームワーク](https://drken1215.hatenablog.com/entry/2019/03/23/214500)

## スタータパック

[ZRK+94さんのツイートより](https://twitter.com/Zen_Re_Kkyo/status/1135152582194651136)より

- [AtCoder社の給料](https://atcoder.jp/contests/abc003/tasks/abc003_1)
- [HSI](https://atcoder.jp/contests/abc078/tasks/arc085_a)
- [Theme Color](https://atcoder.jp/contests/code-festival-2018-final/tasks/code_festival_2018_final_b)
- [デフレゲーム](https://atcoder.jp/contests/tkppc3/tasks/tkppc3_e)
- [Ordinary Beauty](https://atcoder.jp/contests/soundhound2018-summer-qual/tasks/soundhound2018_summer_qual_c)
- [Removing Blocks](https://atcoder.jp/contests/agc028/tasks/agc028_b)
- [Modulo Operations](https://atcoder.jp/contests/exawizards2019/tasks/exawizards2019_d)

## コツ？

- 状態遷移図を書いてみることは意味がありそう。

---

## 問題例

### [yukicoder No.144 エラトステネスのざる](https://yukicoder.me/problems/no/144)

「期待値の線形性」を理解するのにとてもよい題材。  
解説はリンク先のPDFが非常にわかりやすいので、適宜そちらを参照すべし。

**「主客転倒」の概念も絡んでいる気がする。**

とはいえ、「そうなるから、そう」という感覚が正直拭いきれない。  
`X(S)` の定義が巧くて、自分でできるようになるのだろうか？という気持ちになる。

### [yukicoder No.108 トリプルカードコンプ](https://yukicoder.me/problems/no/108)

まずはDPテーブルの定義の仕方に壁がある（とはいえものすごい典型なのだとは思う）。 
個人的には、その次の「3枚以上引いてるカードは引かない場合の条件付き期待値、および条件付き確率」を考えている部分に壁がある。

ほぼ写経でコードを書いてしまったため、DPテーブルの定義は独自のものにアレンジしてみたが、自分のコードのほうが直感的でわかりやすい気がする。

**同じ状態にループしてしまうような遷移は省いて考える」というのが肝なのか？**

**「確率 `p` で起こる事象について、それが起こるまでの期待値は `1/p` 」というのは超典型なので暗記。**

参考: https://drken1215.hatenablog.com/entry/2019/03/23/175300

### [EDPC J.Sushi](https://atcoder.jp/contests/dp/tasks/dp_j)

トリプルカードコンプと全く同じ問題と思って良い。

### [ABC184 D.increment of coins](https://atcoder.jp/contests/abc184/tasks/abc184_d)

これもトリプルカードコンプやSushiと同じ。  
ただし、条件付き期待値や条件付き確率を考えなくてもよいあたり、この問題は確かに易しめだった。

