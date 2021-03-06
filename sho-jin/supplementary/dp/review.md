# 今まで解いたDPの復習

Last Change: 2020-06-28 18:53:46.


<!-- vim-markdown-toc GFM -->

* [問題集](#問題集)
  * [AtCoder](#atcoder)
  * [yukicoder](#yukicoder)
  * [Codeforces](#codeforces)
  * [HackerRank](#hackerrank)

<!-- vim-markdown-toc -->

## 問題集

以下のコマンドで検索した。

```shell
rg -i '(dp|動的計画法)' -g '!*.go' -g '!*.snip' --files-with-matches
```

### AtCoder

- [CODE FESTIVAL 2018A C.半分](https://atcoder.jp/contests/code-festival-2018-quala/tasks/code_festival_2018_quala_c)
- [SoundHound2018本線 B.Neutralize](https://atcoder.jp/contests/soundhound2018-summer-final-open/tasks/soundhound2018_summer_final_b)
- [日経2019本選エキシビション F.コラッツ問題](https://atcoder.jp/contests/nikkei2019-ex/tasks/nikkei2019ex_e)
  - DPに分類すべきかどうか怪しいところ。
- [CODE THANKS FESTIVAL 2018 E.Union](https://atcoder.jp/contests/code-thanks-festival-2018/tasks/code_thanks_festival_2018_e)
- [Tenka1 2016 予選B B.天下一魔力発電](https://atcoder.jp/contests/tenka1-2016-qualb/tasks/tenka1_2016_qualB_b)
- [ABC113 D.Number of Amidakuji](https://atcoder.jp/contests/abc113/tasks/abc113_d)
- [ABC044 C.高橋君とカード](https://atcoder.jp/contests/abc044/tasks/arc060_a)
- [ABC060 D.Simple Knapsack](https://atcoder.jp/contests/abc060/tasks/arc073_b)
  - 見た目がDPなのにDPじゃない方法でも解けてしまう問題。
- [ABC040 C.柱柱柱](https://atcoder.jp/contests/abc040/tasks/abc040_c)
- [ABC041 D.徒競走](https://atcoder.jp/contests/abc041/tasks/abc041_d)
- [ABC036 D.塗り絵](https://atcoder.jp/contests/abc036/tasks/abc036_d)
- [ABC099 C.Strange Bank](https://atcoder.jp/contests/abc099/tasks/abc099_c)
- [ABC104 D.We Love ABC](https://atcoder.jp/contests/abc104/tasks/abc104_d)
- [ABC082 D.FT Robot](https://atcoder.jp/contests/abc082/tasks/arc087_b)
- [ABC029 D.1](https://atcoder.jp/contests/abc029/tasks/abc029_d)
- [ABC054 D.Mixing Experiment](https://atcoder.jp/contests/abc054/tasks/abc054_d)
- [ABC007 D.禁止された数字](https://atcoder.jp/contests/abc007/tasks/abc007_4)
- [AGC030 B.Tree Burning](https://atcoder.jp/contests/agc030/tasks/agc030_b)
  - 部分点がDP。
  - ※本問題集作成時点で未AC。
- [AGC021 A.Digit Sum 2](https://atcoder.jp/contests/agc021/tasks/agc021_a)
- [ABC122 D.We Like AGC](https://atcoder.jp/contests/abc122/tasks/abc122_d)
- [エクサウィザーズ2019 D.Modulo Operations](https://atcoder.jp/contests/exawizards2019/tasks/exawizards2019_d)
  - ※本問題集作成時点で未AC。
  - `2020-02-05` にAC
- [AGC031 B.Reversi](https://atcoder.jp/contests/agc031/tasks/agc031_b)
- [ABC125 D.Flipping Signs](https://atcoder.jp/contests/abc125/tasks/abc125_d)
- [みんなのプロコン2019 D.Ears](https://atcoder.jp/contests/yahoo-procon2019-qual/tasks/yahoo_procon2019_qual_d)
  - ※本問題集作成時点で未AC。
- [ABC117 D.XXOR](https://atcoder.jp/contests/abc117/tasks/abc117_d)
- [ABC118 D.Match Matching](https://atcoder.jp/contests/abc118/tasks/abc118_d)
  - DPでも2種類の解き方があるので、両方復習したい。
- [キーエンス2020 D.Swap and Flip](https://atcoder.jp/contests/keyence2020/tasks/keyence2020_d)
  - ※本問題集作成時点で未AC
  - `2020-03-09` にAC
- [ABC153 E.Crested Ibis vs Monster](https://atcoder.jp/contests/abc153/tasks/abc153_e)
- [ABC145 E.All-you-can-eat](https://atcoder.jp/contests/abc145/tasks/abc145_e)
  - 模範解答はDPではないが、嘘解法に気をつけながらDPで解くことも重要。
- [第二回全国統一プログラミング王決定戦予選 D.Shortest Path on a Line](https://atcoder.jp/contests/nikkei2019-2-qual/tasks/nikkei2019_2_qual_d)
- [ABC135 D.Digits Parade](https://atcoder.jp/contests/abc135/tasks/abc135_d)
- [ABC146 F.Sugoroku](https://atcoder.jp/contests/abc146/tasks/abc146_f)
- [ABC142 E.Get Everything](https://atcoder.jp/contests/abc142/tasks/abc142_e)
- [ABC141 E.Who Says a Pun?](https://atcoder.jp/contests/abc141/tasks/abc141_e)
- [ABC129 C.Typical Stairs](https://atcoder.jp/contests/abc129/tasks/abc129_c)
- [ABC129 E.Sum Equals Xor](https://atcoder.jp/contests/abc129/tasks/abc129_e)
- [diverta 2019 Programming Contest 2 D.Squirrel Merchant](https://atcoder.jp/contests/diverta2019-2/tasks/diverta2019_2_d)
  - ※本問題集作成時点で未AC
  - `2020-02-04` にAC
- [ABC130 E.Common Subsequence](https://atcoder.jp/contests/abc130/tasks/abc130_e)
- [三井住友信託銀行プログラミングコンテスト2019 C.100 to 105](https://atcoder.jp/contests/sumitrust2019/tasks/sumitb2019_c)
- [ABC011 C.123引き算](https://atcoder.jp/contests/abc011/tasks/abc011_3)
- [ARC002 C.コマンド入力](https://atcoder.jp/contests/arc002/tasks/arc002_3)
  - 貪欲でもACしたが、解説スライドに「テストケースが弱い」とか書かれているあたり、嘘解法の可能性が高い。
  - 正攻法はすごろく系のDP。
- [ABC154 E.Almost Everywhere Zero](https://atcoder.jp/contests/abc154/tasks/abc154_e)
- [AGC033 B.LRUD Game](https://atcoder.jp/contests/agc033/tasks/agc033_b)
  - 変種だが、ゲーム系のDPと捉えられる。
- [ABC021 C.正直者の高橋くん](https://atcoder.jp/contests/abc021/tasks/abc021_c)
  - 初見では若干非効率なDPで解いてしまったので、想定解法で解いてみること。
- [ABC155 E.Payment](https://atcoder.jp/contests/abc155/tasks/abc155_e)
  - 変種の桁DP(?)
- [ABC015 D.高橋くんの苦悩](https://atcoder.jp/contests/abc015/tasks/abc015_4)
  - ナップザックDPの変種。簡単。
- [ARC043 B.難易度](https://atcoder.jp/contests/arc043/tasks/arc043_b?lang=ja)
  - 想定解法はDPじゃなかったっぽいが、DPで解けたので。
- [AGC037 A.Dividing a String](https://atcoder.jp/contests/agc037/tasks/agc037_a)
  - 貪欲法で解いたが、想定解法のDPで解けるべき。
- [CODE FESTIVAL 2015 あさぷろ Middle B.ヘイホー君と削除](https://atcoder.jp/contests/code-festival-2015-morning-middle/tasks/cf_2015_morning_easy_d)
  - LCSの練習用、、としたいところだがスニペット貼るだけで終わってしまう。
- [天下一プログラマーコンテスト2014予選B B.エターナルスタティックファイナル](https://atcoder.jp/contests/tenka1-2014-qualb/tasks/tenka1_2014_qualB_b)
  - 想定解法はわからないが、割と自然にDPで解ける。
- [AGC043 A.Range Flip Find Route](https://atcoder.jp/contests/agc043/tasks/agc043_a)
  - DPは答えを求める部分だがそれは超簡単、それ以外の考察のほうが圧倒的に典型感もあり重要。
- [ABC017 D.サプリメント](https://atcoder.jp/contests/abc017/tasks/abc017_4)
  - 累積和によるDP高速化。しゃくとり法も前処理で必要。
  - 数え上げDPの練習にも良い。
- [CODE FESTIVAL 2014 予選B D.登山家](https://atcoder.jp/contests/code-festival-2014-qualb/tasks/code_festival_qualB_d)
  - 自分はSetを使って解いた
- [ARC027 C.最高のトッピングにしような](https://atcoder.jp/contests/arc027/tasks/arc027_3)
  - 貪欲の考察から正しくDPテーブルを定義する必要がある。
- [ABC006 D.トランプ挿入ソート](https://atcoder.jp/contests/abc006/tasks/abc006_4)
  - LISが答えになる問題。
- [ABC162 F.Select Half](https://atcoder.jp/contests/abc162/tasks/abc162_f)
  - 耳DPの系譜らしい？
- [ARC042 C.おやつ](https://atcoder.jp/contests/arc042/tasks/arc042_c)
  - 一工夫必要なナップサックDPという感じだが、思わぬ落とし穴が合った。
  - 「遷移を取りつつグローバルな解を更新する」という手法も検討したい。
- [ABC163 E.Active Infants](https://atcoder.jp/contests/abc163/tasks/abc163_e)
  - `O(n^2)` のDP。
  - 区間DPっぽいらしい？
- [CODE FESTIVAL 2015 予選A B.とても長い数列](https://atcoder.jp/contests/code-festival-2015-quala/tasks/codefestival_2015_qualA_b)
  - 非常に簡単なDP。

### yukicoder

- [yukicoder No.286 Modulo Discount Store](https://yukicoder.me/problems/no/286)
- [yukicoder No.107 モンスター](https://yukicoder.me/problems/no/107)
- [yukicoder No.845 最長の切符](https://yukicoder.me/problems/no/845)
- [yukicoder No.134 走れ！サブロー君](https://yukicoder.me/problems/no/134)
- [yukicoder No.698 ペアでチームを作ろう](https://yukicoder.me/problems/no/698)
- [yukicoder No.357 品物の並び替え（Middle）](https://yukicoder.me/problems/no/357)
- [yukicoder No.771 しおり](https://yukicoder.me/problems/no/771)

### Codeforces

- [Codeforces Round No.597 C.Constanze's Machine](https://yukicoder.me/problems/no/771)
- [Codeforces Round No.591 D.]()
  - DPじゃないかもしれない。。
- [Codeforces Round No.594 C.Ivan the Fool and the Probability Theory](https://codeforces.com/contest/1248/problem/C)
- [Codeforces Round No.605 D.Remove One Element](https://codeforces.com/contest/1272/problem/D)
- [Educational Codeforces Round 90 D. Maximum Sum on Even Positions](https://codeforces.com/problemset/problem/1373/D)
  - おそらく「耳DP」と呼ばれるもの。

### HackerRank

- [ゆるふわ競プロオンサイトNo.3 Div.2 Sweets Distribution(Easy)](https://www.hackerrank.com/contests/yfkpo3-2/challenges/sweets-distribution-easy)
  - We love ABCっぽい雰囲気のDPだと思ったが、耳DPがこれに近いらしい。
  - DPテーブルの持ち方を意識したい。

