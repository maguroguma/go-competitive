# 問題集

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

## 基礎以上

- [yukicoderのセグメントツリーコンテスト](https://yukicoder.me/contests/232)
  - [Range Mindex Query](https://yukicoder.me/problems/no/875)
    - 要素に最小値のほか、最小値が入っているインデックスをもたせた構造体とすれば、自然に定義できる。
  - [Range Compress Query](https://yukicoder.me/problems/no/876)
    - 与えられた定義で計算される圧縮値の他、区間の両端の値をもたせた構造体とすれば、正しくセグメントツリーを定義できる。
      - モノイドの単位元がなにか？がちょっと難しかった。
    - 想定解法は階差数列を取ることで、そうすると遅延伝搬は不要になる。

## キュレーション

- [Codeforcesの問題集](https://codeforces.com/blog/entry/22616)
  - 4年前にまとめられたものなので、ちょっと古いのかも。
- [はまやんはまやんさんの問題集](https://www.hamayanhamayan.com/entry/2017/07/08/173120)

