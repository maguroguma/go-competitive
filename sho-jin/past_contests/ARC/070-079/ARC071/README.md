# ARC071 過去問感想

Last Change: 2020-03-05 22:33:01.

## E問題（@2020-03-05）

試験管なしの青Diff問題だがなんとか自力ACできた（nuipさんwriterの青Diffは何故か解ける）。

解答PDFでは「変換可能なグループ」という見方をしていたが、
自分は `AB -> BA` のようにスワップが可能なことに目をつけた。

`AB -> BBB -> BAAB -> BAAAA -> BA`

よって、変換を駆使して `S', T'` を構成する `A, B` の数がそれぞれ等しくなればよい、というふうに考えた。

まず、 `S'` の `A, B` の数は、ダブリングの操作を両方等しい数行いさえすれば、
いくらでも増やせるので、とりあえず `A, B` の数を `100000` 増やす。
これによって、確実に `S'` のほうが `T'` よりも `A, B` 両方の数が大きくなる。
このあと、適当に `BB->A` の変換を行えば、 `A` の数を `S', T'` で揃えることができる。
`A` の数を揃えた後は、 `BBB` を消すことでしか `B` の数を揃えられないので、
結局 `B` の数の差が3の倍数かどうか、を調べるだけで良い。

