# 問題集

Last Change: 2020-03-02 12:07:16.

## 目標

- Codeforcesの2000点付近の問題を解く。
- yukicoderのセグ木コンを解く。
- 任意の通常セグメントツリーのモノイドを一般化したものの完成。
  - 区間加算と区間更新が分けられるといいかもしれない。
- 任意の遅延評価セグメントツリーのモノイドを一般化したものの完成。
  - 同じく、区間加算と区間更新が分けられるといいかもしれない。
    - できるのか。。？

## verify用問題集

- [AOJのクエリ問題集](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2)
  - [RMQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_A)
  - [RSQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_B)
  - [RMQ-RUQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_F)
  - [RSQ-RAQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_G)
  - [RMQ-RAQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_H)
  - [RMQ-RUQ](https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_I)
    - 最低限の検証用に。
- [ABC125 C.GCD on Blackboard](https://atcoder.jp/contests/abc125/tasks/abc125_c)
  - 区間GCDでモノイドを定義してみよう。
- [ABC024 B.自動ドア](https://atcoder.jp/contests/abc024/tasks/abc024_b)
  - RAQと点取得ができればできるはず。

## 基礎以上

- [yukicoderのセグメントツリーコンテスト](https://yukicoder.me/contests/232)
  - [Range Mindex Query](https://yukicoder.me/problems/no/875)
    - 要素に最小値のほか、最小値が入っているインデックスをもたせた構造体とすれば、自然に定義できる。
  - [Range Compress Query](https://yukicoder.me/problems/no/876)
    - 与えられた定義で計算される圧縮値の他、区間の両端の値をもたせた構造体とすれば、正しくセグメントツリーを定義できる。
      - モノイドの単位元がなにか？がちょっと難しかった。
    - 想定解法は階差数列を取ることで、そうすると遅延伝搬は不要になる。
- DP関係
  - [日経2020予選 D.Shortest Path on a Line](https://atcoder.jp/contests/nikkei2019-2-qual/tasks/nikkei2019_2_qual_d)
    - なぜか本番中に解けたもの。
  - [ARC026 C.蛍光灯](https://atcoder.jp/contests/arc026/tasks/arc026_3)
    - 日経のものとほぼ同じらしい（2020-01-26時点で未AC）
  - [ABC146 F.Sugoroku](https://atcoder.jp/contests/abc146/tasks/abc146_f)
    - 貪欲法でも解けるが、DPでやる方法を理解してから出ないと、正当性の直観的な理解が難しい気がする。
    - ダイクストラ法の辞書順最小の経路復元とも類似している。
- その他
  - [ABC157 E.Simple String Queries](https://atcoder.jp/contests/abc157/tasks/abc157_e)
    - CodeforcesのDiv3でも全く同じものが2019年10月頃に出題されていた。
    - BITをたくさん使っても解ける。

## キュレーション

- [Codeforcesの問題集](https://codeforces.com/blog/entry/22616)
  - 4年前にまとめられたものなので、ちょっと古いのかも。
- [はまやんはまやんさんの問題集](https://www.hamayanhamayan.com/entry/2017/07/08/173120)

