# ABC157 感想

ペナ連発して60分ぐらい5完といったところだが、そこそこハマりどころが合ったようで、
パフォーマンスは1800ちょいと十分。

反省点も復習箇所もたくさんあるので、しっかり勉強すること。

- A問題は切り上げの練習。ただよく考えたら分子は2だけなので、偶奇性に着目してインクリメントすれば、普通の切り捨てわり算だけで良さそう。
- B問題は9マスのビンゴゲームのシミュレーション。ちょっと考えたが特に工夫せずにやった。
- C問題は効率的な構築をやってしまい、コーナーケースにハマってしまった。
  - 本当は全探索をすべき。作った数字を文字列として扱い、条件を確認すべき。
  - 構築のまま通すにしても、コーナーケースは典型的で気づくべきものなので、反省が残る。
- D問題は連結成分のサイズを取得できるUnionFindの問題。
  - 色々と考察が甘いままコードを書いてしまったのは、今にしてみれば悪手だった。
- E問題はかつてこどふぉにでたものそのままだったので、コピペした。
  - クエリの形まで同じだったのは衝撃だったが、BITのサイズは変更する必要があり1REしてしまった。
    - これを考慮してもフルスクラッチで書くよりははるかに早く済んだ。
- F問題は幾何に面食らってしまったが、割と楽しそうな問題なので、しっかり復習したい。

