# ARC017 過去問感想

Last Change: 2020-04-03 00:05:23.

## A問題（@2020-02-09）

素数判定するだけ。
制約的には別に `O(sart(n))` でなくとももっと愚直にやって良い。

問題文の冗長さがこどふぉっぽい。

## B問題（@2020-02-10）

指定された長さの狭義単調増加列の数を求める問題。

階差を `1, 0, -1` の情報に圧縮した上で、累積和を考えた。

尺取法はまだできない。

階差に着目する手法だと `k=1` のケースがコーナーケースとなり、正しく数えられないので、
これは個別に処理する必要がある。1WA。

**階差数列を考えるときは、特にコーナーケースに注意したほうが良さそう。**
というよりも、階差を考えること自体が、競技プログラミングでは悪手な気がする。

解説スライドにある、狭義単調増加列に区切って、その中で長さ `k` のものを `O(1)` でカウントするのが
一番賢いと思う。

## C問題（@2020-04-03）

半分全列挙で一番簡単な問題かもしれない。

解答スライドを観ると二分探索するように書かれているが、今回は値が `X-a` に一致するものの数だけわかればいいので、
ハッシュマップだけで十分なのがよい。

半分全列挙のエッセンスに集中できるのが良いと思う。

