# CODE FESTIVAL 2017 final 過去問感想

Last Change: 2020-02-15 01:56:29.

- A問題は300点で制約も小さいからとても簡単に見えるが、ちょっと頭をひねる必要はある。
  - Aが入る4箇所について「Aを入れるか入れないか？」をビット探索で全探索する。
  - 場合分けを真面目に書くと16通り分岐させる必要があるのでビット全探索をすべき。
  - ちょっと思いつくまでに時間がかかりすぎ。
  - **典型: 全探索できる小さな制約のときは、for文で舐めるような全探索だけでなくビット全探索もすぐに想起できるように。**
- B問題は回文問題で気合を入れて臨んだが、かなりあっけなかった。
  - 特殊なアルゴリズムは不要な、しかも考察は簡単な問題。
  - 結局の所、同じ文字を並べてはいけないこと、3文字の回文も作れないことから、異なる3文字を並べ続けなければならない。
  - 各文字の出現回数をカウントし、最小出現回数のものをそれぞれから引く。
    - このときにすべての文字について差が1以下であればOK、そうでないならNG。

## C問題（@2020-02-15）

わからなくて答え観てしまったけど、「やられた。。」という感想。
実際に参加してたらかなりショックを受けてたかもしれない。

入力のバリエーションの小ささにもっと注目すべきだったかもしれない。
忘れた頃に解き直したい。

とりあえず実装に関しては、なるべく通常のビット全探索に落とすことを目指したほうが、バグりにくくていいと思う。

