# ダブリングの練習

Last Change: 2020-09-27 03:17:53.

ダブリングは応用範囲が広いので、できるだけ諳んじて書けるようにしておきたい。

## 問題集

- [ABC179 E.Sequence Sum](https://atcoder.jp/contests/abc179/tasks/abc179_e)
  - ジャンプ先だけでなく、ジャンプして跨いだ区間の総和を同様に管理しておく。
- [ABC175 D.Moving Piece](https://atcoder.jp/contests/abc175/tasks/abc175_d)
  - 1回以上k回以下という制約のせいで難しい。
  - [けんちょんさんのブログ](https://drken1215.hatenablog.com/entry/2020/08/17/182700)が参考になる。
    - ジャンプ区間の和だけでなく、ジャンプしなかった場合も考慮した `all` というダブリング配列も計算しておく。
- [ABC167 D.Teleporter](https://atcoder.jp/contests/abc167/tasks/abc167_d)
- [ABC013 D.阿弥陀](https://atcoder.jp/contests/abc013/tasks/abc013_4)
- [AGC036 B.Do Not Duplicate](https://atcoder.jp/contests/agc036/tasks/agc036_b)

## コツ？

`A[d][v]` というダブリング配列は **頂点 `v` から出発したときの（ `2^d` 進んだときの）...** というように、
起点が `v` であることを意識するのが良いかもしれない。

