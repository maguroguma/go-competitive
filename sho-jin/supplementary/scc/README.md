# 強連結成分分解

Last Change: 2020-10-16 01:38:51.

ACLに入っていたので、インターフェースだけでも理解する。

## 参考

- [強連結成分分解＆トポロジカルソート](https://hcpc-hokudai.github.io/archive/graph_scc_001.pdf)
  - すごい力作に見える。

## ACLの概要

Goの移植版については[EmptyBoxさんのこちら](https://qiita.com/EmptyBox_0/items/2f8e3cf7bd44e0f789d5#scc)を拝借する。

- 構造体を `new` したあとはそこに一つ一つ有向辺を足していくような形で構築していく。
- `Scc` メソッドで、強連結成分分解が得られる。
- 強連結成分一つ一つのnode idリストの順番は不定。
- 強連結成分の順序はトポロジカルソートされて出力される。

