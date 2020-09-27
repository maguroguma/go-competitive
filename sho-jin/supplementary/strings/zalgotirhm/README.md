# Z-Algorithm

Last Change: 2020-09-27 19:32:44.

## 概要

ある文字列 `S` について、 `S[i:] (i = 0, 1, ..., |S|-1)` を考える。

すべての `S[i:]` について、「元の文字列 `S` との最長の共通接頭辞の長さ」を求めたい。
Z-Algorithmは、 `O(|S|)` でこれを実現する。

※愚直に実装すると、明らかに `O(|S|^2)` になってしまう。

※アウトプットの配列のことを、Z配列などと呼ぶこともあるらしい。

## 参考

個人的にはそこそこ納得するのが難しいアルゴリズムだと思い、
図示してくれていたりと工夫されている文献は多かったが、なかなか納得できなかった。

その中でも、[かつっぱさんの動画解説](https://www.youtube.com/watch?v=f6ct5PQHqM0&feature=youtu.be)
はかなり丁寧でわかりやすかった。

## アルゴリズムのポイント

（境界などの細かい部分は別にして、）一度スキャンしたところは2回以上スキャンしないように工夫する。
すなわち、以前までの確定したZ配列の情報を利用する。

一方で、未知の部分、すなわちまだスキャンできていない先の文字列については、都度素直にスキャンしていく。

## 実装のポイント

かつっぱさんの動画中の実装をGoにそのまま移植した。

最大のポイントは、スキャン済みの部分をスキップする部分。
この部分で3つの場合分けを考える必要があるが、
特にif文を用いなくても、不要なスキャンがなくなるように（例えばすぐにbreakしたりする）
コードが組まれている。

`*same` の値の変化に着目することが最重要。

```go
// ZAlgorithm calculates Z-values of a slice S.
// Z-values show a maximum length of prefix between S and S[i:] for each i.
// Time complexity: O(n)
func ZAlgorithm(S []rune) (Z []int) {
	_min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(S)
	Z = make([]int, n)

	// 探索済みの最右
	from, last := -1, -1
	for i := 1; i < n; i++ {
		same := &Z[i]

		// Z-Algorithm!
		if from != -1 {
			// Z[i-from]: 先頭から少し進めた部分のZ値
			// (last-i): 以前にスキャンしたことのある部分の長さ
			*same = _min(Z[i-from], last-i)
			*same = _max(0, *same)
		}

		// 素直なスキャンを末尾まで行う
		for i+(*same) < n && S[*same] == S[i+(*same)] {
			(*same)++
		}

		// 最右を更新する
		if last < i+(*same) {
			last = i + (*same)
			from = i
		}
	}

	Z[0] = n
	return Z
}
```
