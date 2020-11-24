# 確率DP

Last Change: 2020-11-25 01:52:08.

期待値DPが一番の勉強目的だったが、確率DPや期待値の線形性といったトピックも一緒に学ぶ。

## 参考

- [確率DPを極めよう](https://compro.tsutaj.com//archive/180220_probability_dp.pdf)
- [けんちょん氏のABC078-C HSIの解説](https://drken1215.hatenablog.com/entry/2019/03/23/175300)
- [「〜がすべて終わるまでの試行回数の期待値」を求める一般的なフレームワーク](https://drken1215.hatenablog.com/entry/2019/03/23/214500)
- [「期待値の線形性」についての解説と、それを用いる問題のまとめ](https://qiita.com/drken/items/3fe1613c44de1f3bfbe1)
  - かなり易しいところから始めてくれるので、迷ったら読み返すと良い。

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

### [SoundHound Inc. Programming Contest 2018 C.Ordinary Beauty](https://atcoder.jp/contests/soundhound2018-summer-qual/tasks/soundhound2018_summer_qual_c)

「期待値の線形性」の最初の練習に。  
`m-1` 個の区間1つ1つが01をとる確率変数だと思って、独立に計算して最後に足し合わせる。

### [ABC008 C.コイン](https://atcoder.jp/contests/abc008/tasks/abc008_3)

エラトステネスのざると一緒に勉強したい問題。  
ただし、ざるよりも確率の計算部分で頭を柔軟にする必要がある。

コイン一つ一つについて01をとる確率変数だと考えればよいのはすぐわかる。  
それぞれについて、「自身を除いた約数について、自分より左側に偶数個並ぶ確率」を求めることになるが、
これを素直に確率の定義に従って「分母は `n!` で分子は〜」と考え始めると泥沼。  
確率を正確に捉えてもっとシンプルに考えれば良い。

この場合は、自身および約数以外はないものとして考えればよい。  
また、約数についてもすべて同質のものと考えて1列に並べてしまい、それらの間に自身を挿入する、という状況を考える。  
すると約数の数を `m` としたときに `Ceil(m+1, 1) / (m+1)` で求まることがわかる。

場合の数に比べて、確率は考え方によっていくらでも簡単になったり難しくなったりするので、そこは訓練が必要。

※なんか2018年7月とかに通してた。こんな教育的問題を雑に消費してしまっていたなんて。。

### [ABC078 C.HSI](https://atcoder.jp/contests/abc078/tasks/arc085_a)

「確率 `p` で起こる事象について、それが起こるまでの期待値は `1/p` 」という超典型。  
できれば `e = (e+1)*(1-p) + p` という漸化式ももう暗記してしまうのが良い。
