# ARC100 過去問感想

Last Change: 2020-03-11 23:40:56.

## D問題（@2020-03-11）

有名問題。
DPをやりたくなるような見た目をしているが、素朴に考えるとある性質に気づく。

ある数値 `S` を2つの数値 `a, b` に分けると、大きい方を `a` とすると、必ず、
`a >= S/2, b <= S/2` となる。

よって、真ん中で切ったときは、その両端での切り方はできるだけ半分に近いほうがよい、とわかる。

なので、真ん中の切り方をすべて全探索すれば良い。

。。が、左と右の半分を効率的にそれぞれ計算する方法がわからなかった。
ので、解説PDFを見た。

**しゃくとり法を使えば良い。**

ただし、今回のしゃくとり法では片方の端点の伸ばし方のみに集中すれば良い。

左側だけを考えると、半分に切ったときの切り口の位置は、真ん中の位置を右にずらした後、左に戻すことはない。
つまり **単調性を有する** といえる。
（PDFでは、左側の切り口を `F[i]` という関数に置いて、その関数について広義の単調性を主張している、説明がうまい。）

。。実装が非常に汚くなってしまったので、しゃくとり法を復習した後にまた解き直したい。

