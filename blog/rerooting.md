次に出題されたら絶対に落としたくないという気持ちから、再帰関数による全方位木DPを抽象化したコードを書いてみました。

セグメント木の際と同様に、型 `T` を都度書き換えないと行けないのがかっこ悪いですが。。
どなたかいい方法を御存知でしたら教えて下さい。

```go
type T int

type ReRooting struct {
	n int
	G [][]int

	ti      T
	dp, res []T
	merge   func(l, r T) T
	addNode func(t T, idx int) T
}

func NewReRooting(
	n int, AG [][]int, ti T, merge func(l, r T) T, addNode func(t T, idx int) T,
) *ReRooting {
	s := new(ReRooting)
	s.n, s.G, s.ti, s.merge, s.addNode = n, AG, ti, merge, addNode
	s.dp, s.res = make([]T, n), make([]T, n)

	s.Solve()

	return s
}

func (s *ReRooting) Solve() {
	s.inOrder(0, -1)
	s.reroot(0, -1, s.ti)
}

func (s *ReRooting) Query(idx int) T {
	return s.res[idx]
}

func (s *ReRooting) inOrder(cid, pid int) T {
	res := s.ti

	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}

		res = s.merge(res, s.inOrder(nid, cid))
	}
	res = s.addNode(res, cid)
	s.dp[cid] = res

	return s.dp[cid]
}

func (s *ReRooting) reroot(cid, pid int, parentValue T) {
	childValues := []T{}
	nexts := []int{}
	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}
		childValues = append(childValues, s.dp[nid])
		nexts = append(nexts, nid)
	}

	// result of cid
	rootValue := s.ti
	for _, v := range childValues {
		rootValue = s.merge(rootValue, v)
	}
	rootValue = s.merge(rootValue, parentValue)
	rootValue = s.addNode(rootValue, cid)
	s.res[cid] = rootValue

	// for children
	accum := s.merge(s.ti, parentValue)
	length := len(childValues)
	if length == 0 {
		return
	}
	if length == 1 {
		s.reroot(nexts[0], cid, s.addNode(accum, cid))
		return
	}

	// cid has more than one child
	R, L := make([]T, length), make([]T, length)
	L[0] = s.merge(s.ti, childValues[0])
	for i := 1; i < length; i++ {
		L[i] = s.merge(L[i-1], childValues[i])
	}
	R[length-1] = s.merge(s.ti, childValues[length-1])
	for i := length - 2; i >= 0; i-- {
		R[i] = s.merge(R[i+1], childValues[i])
	}

	for i, nid := range nexts {
		if i == 0 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, R[1]), cid))
		} else if i == length-1 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, L[length-2]), cid))
		} else {
			s.reroot(nid, cid, s.addNode(s.merge(accum, s.merge(L[i-1], R[i+1])), cid))
		}
	}
}
```

## 参考

[ei1333さんによる全方位木DPの解説記事](https://ei1333.hateblo.jp/entry/2017/04/10/224413)

[ABC160-Fのすぬけさんによる解説](https://www.youtube.com/watch?v=zG1L4vYuGrg&t=5950s)

お二人の解説によって、全方位木DPがすっきりと理解できました。
解説を読んだ（聴いた）後、まずは抽象化を考えずに、素直にDFSを2回やる手法で、アドホックなコードを書いてみました。

いきなり抽象化から入るのは大変かもしれないので、まずはこれらを参照するのが良いかと思います。

[keymoonさんによる全方位木DPの抽象化解説記事](https://qiita.com/keymoon/items/2a52f1b0fb7ef67fb89e)

こちらの記事を受けて[Goによる実装（写経）](https://qiita.com/_maguroguma/items/5f51da6eac7829929229)を書いてみることで、
今回の再帰関数による抽象化も書くことができました。
keymoonさん、ありがとうございます！

