# Suffix ArrayとLCP Array

Last Change: 2020-10-11 15:50:20.

@2020-10-11時点で、Suffix ArrayおよびLCP Arrayを求めるアルゴリズムは確認していない。  
インタフェースと利用方法の概要だけを抑える。

## 参考

- [ACL String](https://atcoder.github.io/ac-library/production/document_ja/string.html)
  - ACLはSuffix ArrayとLCP Arrayの生成までしか担保していない。パターン検索とかは自分でやってという感じ？
- [tiqwablogさんの記事](https://blog.tiqwab.com/2018/09/17/suffix-array.html)
  - Suffix Arrayがテキスト検索にどのように使えるの？というのが最初の節でコンパクトに纏められている。
- [wikipedia](https://ja.wikipedia.org/wiki/%E6%8E%A5%E5%B0%BE%E8%BE%9E%E9%85%8D%E5%88%97)
  - 具体的な擬似コードなどはないが、その分概要はわかりやすい。
- [niuezさんのブログ](https://niuez.hatenablog.com/entry/2019/12/16/203739)
  - C++のコードとともにパターン検索のコードも書いてくれているので、参考にさせていただく。

---

## ACLより

### Suffix Array

文字列の長さ `S` を `n` とすると、その文字列のsuffixは `n` 個ある。  
`S[0:], S[1:], S[2:], ..., S[n-1:]`, 一般化して `S[i:]`  
Suffix Arrayとは、この `n` 個のsuffixを辞書順に並び替えたときの、 `i` を格納した長さ `n` の `[]int` である。

### LCP Array

文字列 `S` とその文字列のSuffix Array `sa` から求まる配列。

文字列の長さ `S` を `n` とすると、その文字列のLCP Arrayの長さは `n-1` である。  
LCP Arrayとは何かというと、
`i` 番目の要素が、suffix文字列 `S[sa[i]:]` とsuffix文字列 `S[sa[i+1]:]` のLCPの長さである配列のことを指す。

※LCPとはLongest Common Prefixのこと。

---

## Suffix Array（およびLCP Array）をつかったパターン検索

以下はwikipediaからの引用。

> 接尾辞配列においては、文字列が出現する位置を求めることはつまり、
> **その文字列で始まっている接尾辞を求めることと同じである。**
> 接尾辞配列は辞書順にソートされているので、検索対象となる文字列の検索には、高速な二分探索アルゴリズムが利用できる。
> `m` を検索文字列の長さとすると、単純な実装では二分探索で `O(m * logn)` の計算時間になる。

二分探索の判定パートで、毎回 `m` 文字だけ照合が必要になるから、このような計算量になる。

### LCP Arrayを用いた検索

うまくやると `O(m + logn)` になるらしい。  
しかしながら、LCP Arrayの計算が必要なわけで、その分を償却するのはけっこう大変な気が。。

