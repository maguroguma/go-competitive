# ABC058過去問感想

- A問題、B問題は特になし
- C問題はやり方は手順を踏んで考えればほぼ自明だけど、実装はちょっと工夫が必要
    - `map` を使った方法と `rune` を `int32` のエイリアスであることを利用した配列アクセスの方法の2つが思い浮かんだが、後者を選んだ
        - Goであんまり `map` を使っていないので、こちらも試しに実装したい
            - 実際に `map` バージョンも作ろうとしたが、初期化周りのチェックが恐ろしく面倒だったので投げた
- D問題は算数・数学系の問題
    - 出されるとできなくて落ち込む系
    - とはいえ典型問題の香りもするので、ヒントとなりそうな点を列挙してみる
        - とりあえず全探索を基本として、数式で和を表現する
        - 高校数学にならい、因数分解してみてひらめかないか試してみる
            - 実際に、因数分解できた時点で、まだ不十分ではあるものの計算量は激減している
        - もともとの問題の制限から、`O(N)` まで落とせればいけるということに到達する
        - `O(N)` に落とす最後のステップを考える
