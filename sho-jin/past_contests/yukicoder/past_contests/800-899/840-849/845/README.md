# 845: 最長の切符

- bitDPの守備範囲を知れる問題。Nが20以下の順列全探索系が出たらbitDPを検討したい。
  - **典型: bitDPはN個全部を使った順列（階乗）だけでなく、部分集合の順列に対しても計算できる。**
- グラフが与えられるわけだが、問題設定的にノード間の最大コストのエッジ以外無視しなければいけないところで引っかかった。

