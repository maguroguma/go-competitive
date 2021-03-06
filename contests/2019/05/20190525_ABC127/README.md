# ABC127 感想

2ペナ4完、トータル時間約46分だった。

どちらも
ちゃんとほんの少し注意するだけで防げたものなので、


- A問題は特に気をつけるところのないif文の練習問題。
- B問題は特に気をつけるところのないfor文の練習問題。
- C問題はかなり易しめの300点問題。
  - コーナーケースと呼んでいいかもわからないような簡単なケースで1ペナ出してしまったので大反省。
- D問題は簡単そうだけど実装で苦戦してしまった。
  - あとあとで思ったが、変化後の配列を長さNまで作ってしまってそれを比較することに集中したほうがバグらせにくそうだった。
    - 解説PDFもこの方法を紹介していた。
    - **一度に複数のことをやろうとするとバグらせる確率が高くなる。**
      - 確か、リーダブルコードでも避けるようにしようと書かれていた内容のはず。
  - 解説放送ではpriority queueを使った方法を紹介していたので、こちらも書いてみることにする。
  - **典型チック: 「値を変化させる」の読み替え**
- E問題は色々と時間を使って考えたが、本番中は何もわからなかった。
  - すぬけさんの解説放送を聞いて学習した。
  - 整理するのは大変だが、コードにすると比較的シンプルになるのは面白い。
    - 組み合わせは、高校数学の約分済みの手計算用の式を使うのではなく、定義の式を使ったほうが実装が楽な気がする。
      - というか細かい数字を気にしなくて良いので精神的に良い。
  - **典型: 全体の動きを追うのではなく、個別の要素に着目する。**
  - **典型: 足し算の順番を変えて計算量を落とす（≒圧縮する）。**

## 解説放送メモ

### D問題

結論: 構造体priority queueの練習問題としてよかった。

- 当たり前だけど、2回同じカードを書き換えるのは無駄。
  - ココらへんをちゃんと踏まえておくのは正しい考察には大事。
- カードを塗り替えるというよりは、用意されたカードを並び替えて、大きなN個の数字をとってくる、という読み替えもできる。
  - こういう見方をすると、確かにpriority queueを使おう、という問題になる。
  - **値と枚数のペアを構造体にして、priority queueに突っ込む。**

**実際に書いてみたところ、PQを使う方法がコードは一番スッキリしており、解法としても見通しが良いので、PQが慣れているのであればこの方法を選びたい。**

**一回あたりのpush, popの計算量 `O(log(要素数))` というの改めて抑えておく。**

