# heap tree by golang

ソースを読むと計算量についても記述されている。

`n = h.Len()`

- `Init`: `O(n)`
- `Push`: `O(log(n))`
- `Pop`: `O(log(n))`
- `Remove`: `O(log(n))`
- `Fix`: `O(log(n))`
  - ただし `Remove -> Pop` よりも若干高速。
