# dwango5th B. Sum AND Subarrays

Last Change: 2020-05-04 22:29:34.

## @2020-05-04

一発AC。

なんか解説PDFが異様に難しく見える。

単純に上のビットから貪欲に建てられるかどうか調べていけば良い。
また、探索パートは素直に `O(N^2)` 個のスライスを正直にスキャンしても間に合う。

下位のビットに進むに当たり、調べるスライスのサイズを絞り込んでいくことで、正しい最大値が見つかる。

