# 動的計画法

## LCS(Longest Common Subsequence)

螺旋本の説明が今まで見てきたものの中で一番親切に思えた。

※単に昔よりも力がついたからだけかもしれない。

```go
// 疑似コード
if i == 0 || j == 0 {
  c[i][j] = 0
}

if i, j > 0 && X[i] == Y[i] {
  c[i][j] = c[i-1][j-1] + 1
} else {
  c[i][j] = max(c[i][j-1], c[i-1][j])
}
```

