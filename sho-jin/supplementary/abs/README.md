# 絶対値記号を外すタイプ

Last Change: 2020-09-26 15:24:47.

高校数学だと2乗して外すことが多かった気がするけど、競技プログラミング、「というか計算機に全探索させられる」
という文脈のもとでは、
**愚直にに `|A| = max(-A, A)` の全パターンを試す** という方法のほうがうまくいくことが多い気がする。

## 問題集

- [ABC147 E.Balanced Path](https://atcoder.jp/contests/abc147/tasks/abc147_e)
- [codeFlyer final B.交通費](https://atcoder.jp/contests/bitflyer2018-final/tasks/bitflyer2018_final_b)
- [ABC100 D.Patisserie ABC](https://atcoder.jp/contests/abc100/tasks/abc100_d)
  - [アルメリアさんの解説](https://betrue12.hateblo.jp/entry/2018/06/17/132624)
- [AGC036 A.Triangle](https://atcoder.jp/contests/agc036/tasks/agc036_a)
- [ABC178 E.Dist Max](https://atcoder.jp/contests/abc178/tasks/abc178_e)
  - [45度回転](https://kagamiz.hatenablog.com/entry/2014/12/21/213931)と併せて理解したい。

### マンハッタン距離と45度回転

45度回転を他に応用するのはちょっと難しそうなので、とりあえずマンハッタン距離を考えるということに絞る。

マンハッタン距離は、普通の2次元座標系では、原点中心に正方形が斜めにしたような等距離線が描かれる。
（ひし形と呼びたくなるが、別に斜めってるからひし形ということは一切ない。正方形である。）

この通常2次元座標系を45度回転した世界では、等距離線を表す正方形が、もとのx軸y軸に並行になるように、
向きが正されるようなイメージになる。

このような座標変換は、45度回転の行列変換を考えると `(x', y') = sqrt(2) * (x-y, x+y)` のようになるが、
（※この問題を背景とする限りに置いては、とりあえず実数の係数は忘れてしまって良い。）

このような変換後の座標系に置いて、2点間の距離は
`max(|x1' - x2'|, |y1' - y2'|)`
に変換される。

