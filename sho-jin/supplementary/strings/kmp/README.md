# KMP法（クヌース・モリス・プラット法）

Last Change: 2020-10-09 02:27:45.

結構難しいと思うが、わかってくるとZ-Algorithmと似たような考え方になってくる気もする
（パターン文字列のprefixに着目するあたり）。

## 参考

- [wikipedia](https://ja.wikipedia.org/wiki/%E3%82%AF%E3%83%8C%E3%83%BC%E3%82%B9%E2%80%93%E3%83%A2%E3%83%AA%E3%82%B9%E2%80%93%E3%83%97%E3%83%A9%E3%83%83%E3%83%88%E6%B3%95#:~:text=%E3%82%AF%E3%83%8C%E3%83%BC%E3%82%B9%E2%80%93%E3%83%A2%E3%83%AA%E3%82%B9%E2%80%93%E3%83%97%E3%83%A9%E3%83%83%E3%83%88%E6%B3%95%EF%BC%88,%E5%8C%96%E3%81%99%E3%82%8B%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A7%E3%81%82%E3%82%8B%E3%80%82)
  - 前処理部分を除けば、KMP法の概要を知る分には、悪くない。とはいえ、KMP法の本体は前処理部分だと思うので、そこについては別の資料をあたったほうがわかりやすい。
- [kimiyukiさんのブログ](https://wiki.kimiyuki.net/Knuth-Morris-Pratt%E6%B3%95)
  - かなり端折られているが、このブログのおかげで前処理≒borderということに気づけた。リンク集も大変参考になる。
- [deve68さんのブログ](https://deve68.hatenadiary.org/entry/20120117/1326749583)
  - 競技プログラマな方ではないかもしれないが、borderについての解説がとてつもなくわかりやすかった。
  - クローズされると悲しくなりそうなので、いくつか抜粋してメモさせていただく。
- [すぬけさんのブログ（KMP法）](https://snuke.hatenablog.com/entry/2014/12/01/235807)
  - コードを含めて、最後はこちらを参考にさせていただいた。
  - 前処理部分のコードしかなかったため、テキスト検索のコードについてはdeve68さんのブログのものを参考にした。
- [すぬけさんのブログ（文字列の周期性判定）](https://snuke.hatenablog.com/entry/2015/04/05/184819)
  - これはおまけだが、せっかくなので履修した。

---

## deve68さんのブログより

文字列の **border, tagged border** という概念について知る。

### borderとは

**ある文字列のprefixとsuffixが一致している場合、一致しているprefixとsuffixをその文字列のborderと呼ぶ。**

以下は注意点。

- prefixとsuffixは重複する部分があっても良い。
- ただし、その文字列自身は除外する。
- borderは空文字でもOK。
- 空文字にはborderが存在しない。
  - 長さで言うならば `-1` となる。
  - borderが空文字列のときは、長さは `0` となる。
- **最も幅広いborderをtagged borderと呼ぶ。**

```
assert taggedBorderOf("ababaa") == "a"
assert taggedBorderOf("ababa") == "aba" // prefix と suffix で a を共有している
assert taggedBorderOf("abab") == "ab"
assert taggedBorderOf("aba") == "a"
assert taggedBorderOf("ab") == ""       // border が空文字の場合もある
assert taggedBorderOf("a") == ""
assert taggedBorderOf("") == null       // 空文字に border は存在しない
```

### borderの持つ性質

**ある文字列のborderのborderもまた、ある文字列のborderである。**

```
assert bordersOf("aabaa") == ["aa", "a", ""]
assert bordersOf("aa") == ["a", ""]
assert bordersOf("a") == [""]

// border は再帰的な構造になっている
assert bordersOf("aabaa") == ["aa", *bordersOf("aa")]
```

---

## すぬけさんのブログより

KMP法では、「パターン文字列 `P` のすべてのprefix `P[:i]` に対して、tagged borderを効率的に求める」
ことが肝要となる。
また、本題はテキスト中のパターン文字列の高速な検索だが、これについても
**上述の前処理にあたるtagged borderの高速計算が理解できれば、同時に自然と理解できる。**

### アルゴリズム理解のコツ

アルゴリズムのコードは非常に短い。
が、これだけで何をやっているかを完全に理解するのは結構大変（少なくとも自分は大変だった）。

```cpp
A[0] = -1;
int j = -1;
for (int i = 0; i < S.size(); i++) {
  while (j >= 0 && S[i] != S[j]) j = A[j];
  j++;
  A[i+1] = j;
}
```

理解のコツは、以下のようなところだと思う。

- `j` という変数がtagged borderの長さを表しており、それと同時に、パターン `P` の次の照合先となる（0-indexな）インデックスを表している。
- `j = A[j]` というのが不一致時の挙動で、これは `j` を前方に戻していることに相当する。
  - そして、戻している先は、「不一致が生じたパターン文字列の直前部分のtagged border」といえる。
  - これは、borderを再帰的に処理しているような形になっている。
  - これによって、次に調べるべきところまで戻っても、tagged border分は照合が済んでいるような形になっている。
    - **よって `i` の方は後ろに下げる必要がない。**

アルゴリズムは2重ループのため `O(N^2)` に見えなくもないが、
やはり `i` が後ろに下がらないことで効率的になっていることがわかる。

なぜなら、 `j = A[j]` によって `j` は後ろに戻ることがあるが、
戻らない場合は `i` と同時にインクリメントされることになる。
そして、 `j` の最小値は `-1` である。

これらのことから、「 `j` が戻れる量は最大でも `N` までである」とわかり、計算量は `O(N)` となる。

---

## 文字列の周期

一応、せっかくなのでついでに覚えておく。

「最小の周期長」＝「 `k` 文字ずらしたものが元の文字列と一致するような最小の `k (k > 0)` 」

※ [この問題](https://atcoder.jp/contests/jag2015summer-day2/tasks/icpc2015summer_day2_f)の説明を読んだ感じ、
`k` については任意の正の整数で定義できるっぽい。

そして、あるパターン文字列のすべてのprefixについて、最小の周期長は `O(N)` でもとまる。

具体的には、 `i - KMP[i] (1-index)` が答えとなる。
これは、KMP法の前処理、すなわちborderの定義からも自然と導かれる。

---

## wikipediaの解説より

※以下は等幅フォントじゃないとみづらい。

```
m: 01234567890123456789012
S: ABC ABCDAB ABCDABCDABDE
W: ABCDABD
i: 0123456

m: 01234567890123456789012
S: ABC ABCDAB ABCDABCDABDE
W:     ABCDABD
i:     0123456

m: 01234567890123456789012
S: ABC ABCDAB ABCDABCDABDE
W:         ABCDABD
i:         0123456

m: 01234567890123456789012
S: ABC ABCDAB ABCDABCDABDE
W:            ABCDABD
i:            0123456

m: 01234567890123456789012
S: ABC ABCDAB ABCDABCDABDE
W:                ABCDABD
i:                0123456
```

## すぬけさんのブログより

```
aabaabaaa
_010123452
```

