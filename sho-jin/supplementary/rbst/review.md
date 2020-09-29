# Setが必要な問題の復習

Last Change: 2020-04-09 00:10:01.

## RBST等で解く問題

- [AGC005 B.Minimum Sum](https://atcoder.jp/contests/agc005/tasks/agc005_b)
  - Setの使い方も重要だがそれ以上に、よくある数列の問題の **「各要素が解に対して何回寄与するのかを数えるタイプの問題」** としても重要な典型問題。
  - `Insert, MinGreater, MaxLess` の操作が必要、番兵の事も考えると任意の整数値をkeyにできる方が好ましい。
  - 要素数 `200000` 程度だと500ms程度。
  - [画像解説](./images/AGC005-B.jpg)
- [ABC140 E.Second Sum](https://atcoder.jp/contests/abc140/tasks/abc140_e)
  - ↑の問題の上位互換、詰めの部分が結構大変、とはいえ「2番めに大きい・小さい・短い」などは典型問題の1つと考えたい。
  - `Insert, MinGreater, MaxLess` の操作が必要、こちらは番兵を使わないほうが楽と思われる。
  - [画像解説](./images/ABC140-E.jpg)
- [ABC128 E.Roadwork](https://atcoder.jp/contests/abc128/tasks/abc128_e)
  - Set意外の要素が強く、そしてそこが難しいので、Set操作の練習には向かないかもしれない。
  - イベントソートの代表問題と思っておく。
  - 任意の整数範囲のkeyについて `Insert, Delete, FindMinimum` の操作が必要。
  - 手持ちのTreapでは1200ms弱かかっているので改良したい。
    - この問題は要素数が `200000*2` なのでそこまで気にする必要はないかもしれないが。。
  - [画像解説](./images/ABC128-E.jpg)
- [Donutsプロコンチャレンジ2015 C.行列のできるドーナツ屋](https://atcoder.jp/contests/donuts-2015/tasks/donuts_2015_3)
  - Setを使う技法といもす法（あるいは遅延セグ木）が必要。
    - `Insert, MinGreater` が必要。
  - Set操作パートは比較的簡単だと思う。
- [CODE FESTIVAL 2014 予選B D.登山家](https://atcoder.jp/contests/code-festival-2014-qualb/tasks/code_festival_qualB_d)
  - Setを使って解く場合は、同じ高さのものの扱いについてちょっと工夫する必要がある。
  - 想定解法はDPやstackを使うもの。

## BIT等で解く問題

[こちら](../BIT/README.md)を参照されたい。

