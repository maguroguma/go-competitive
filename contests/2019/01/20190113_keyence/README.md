# KEYENCE Programming Contest 2019感想

- A問題はいろいろやり方がありそうでちょっと迷ったが、intが4つだけなのでソートしてハードコーディングした。
  - 解説放送でもソートが真っ先に出ていたので、多分良かったのだと思う。
- B問題でやらかしてしまった。2WAの35分とかそれぐらい。
  - 文字列処理だが、先入観をもって臨んでしまったために、ちゃんと整理できていなかった。
  - WAしたあとも落ち着いて整理したが、作りたい文字列 `"keyence"` の間に不純物が入り込んでいる場合のうまい処理がわからなかった。
  - ~~**連結の仕方を全パターン試す** という解き方が正解らしい。~~
  - **切り取りのパターンを全パターン試して、残った文字列を連結してチェック** が模範解答。
  - **復習する価値の高いB問題。**
- C問題は400点としては自分の経験の中では最易に感じられた。
  - シンプルな貪欲法だと思った。
- D問題は考察は案外正しかったようだが、実装し切るイメージは沸かなかった。
  - 惜しかったとかでは一切ない。
  - 練習のために二分探索を使って解いてみたが、実際のところは特に必要ではなかった。
    - `O(nm)` のループの中で `O(Max(logn, logm))` の二分探索を行っていたため、計算量はギリギリだったと思う。
  - **ロジカルに考えることを諦めなければ正解にたどり着ける可能性もなかったとは言えない。**
    - 難しそうに見えてよくよく考えると非常にクリアに実装できる言い問題だと思った。
    - 解き切るには経験値があまりにも足りていないので演習を続けましょう。
