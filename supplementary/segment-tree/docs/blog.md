<!-- セグメントツリー（Go） -->

セグメントツリーをもう少し取り回しが効くようにしたいなぁと思った((解説で「セグメントツリーでも解けます」とあるときにテンション下がるのは嫌ですよね。))ので、
他の方々のブログ等を参考にしながら書き直してみました((言語を変えただけで、基本的に写経です。))。


<!-- vim-markdown-toc Marked -->

* [実装](#実装)
  * [通常（遅延伝搬なし）](#通常（遅延伝搬なし）)
    * [例: yukicoder No.875 Range Mindex Query](#例:-yukicoder-no.875-range-mindex-query)
  * [遅延伝搬あり](#遅延伝搬あり)
    * [例: yukicoder No.876 Range Compress Query](#例:-yukicoder-no.876-range-compress-query)
* [参考](#参考)
  * [tsutajさんの解説](#tsutajさんの解説)
  * [ei1333さん・beetさんの解説](#ei1333さん・beetさんの解説)

<!-- vim-markdown-toc -->

## 実装

詳細は後述しますが、tsutajさん・ei1333さん・beetさんのお三方の解説を大いに参考にさせていただきました。

変数名やメソッド名はそれぞれの記事から拝借しているため、若干キメラになっていて読みにくいかもしれません。

また、数学的な側面の理解が浅いので、モノイド・作用素モノイド周りの命名が特に気持ち悪いかもしれません((アドリブでいじるときに、自分にとって意味を思い出しやすいように意識して書いたらこうなりました))。

### 通常（遅延伝搬なし）

```go
type T int // (T, f): Monoid

type SegmentTree struct {
  sz   int              // minimum power of 2
  data []T              // elements in T
  f    func(lv, rv T) T // T <> T -> T
  ti   T                // identity element of Monoid
}

func NewSegmentTree(
  n int, f func(lv, rv T) T, ti T,
) *SegmentTree {
  st := new(SegmentTree)
  st.ti = ti
  st.f = f

  st.sz = 1
  for st.sz < n {
    st.sz *= 2
  }

  st.data = make([]T, 2*st.sz-1)
  for i := 0; i < 2*st.sz-1; i++ {
    st.data[i] = st.ti
  }

  return st
}

func (st *SegmentTree) Set(k int, x T) {
  st.data[k+(st.sz-1)] = x
}

func (st *SegmentTree) Build() {
  for i := st.sz - 2; i >= 0; i-- {
    st.data[i] = st.f(st.data[2*i+1], st.data[2*i+2])
  }
}

func (st *SegmentTree) Update(k int, x T) {
  k += st.sz - 1
  st.data[k] = x

  for k > 0 {
    k = (k - 1) / 2
    st.data[k] = st.f(st.data[2*k+1], st.data[2*k+2])
  }
}

func (st *SegmentTree) Query(a, b int) T {
  return st.query(a, b, 0, 0, st.sz)
}

func (st *SegmentTree) query(a, b, k, l, r int) T {
  if r <= a || b <= l {
    return st.ti
  }

  if a <= l && r <= b {
    return st.data[k]
  }

  lv := st.query(a, b, 2*k+1, l, (l+r)/2)
  rv := st.query(a, b, 2*k+2, (l+r)/2, r)
  return st.f(lv, rv)
}

func (st *SegmentTree) Get(k int) T {
  return st.data[k+(st.sz-1)]
}
```

おそらくは、モノイドの定義（ `T, f, ti` ）を適切に書き換えることだけに注力すれば、
うまく動くのではないかと思います((AOJにある典型例（RMQ, RSQ）は検証済み。))((`T` の型が複雑になるときのインスタンス化の速度への影響が、まだ十分に検証しきれていません。))。

#### 例: yukicoder No.875 Range Mindex Query

[問題のURL](https://yukicoder.me/problems/no/875)

yukicoderの解説にもある通り、
最小値に加えて、最小値が入っているインデックスをもたせた構造体を要素の型とすればOKです。

以下はコードの抜粋です（提出は[こちら](https://yukicoder.me/submissions/419460)）。

```go
type T struct {
	v   int
	idx int
}

func main() {
	n, q := ReadInt2()
	A := ReadIntSlice(n)

	f := func(lv, rv T) T {
		t := T{}
		if lv.v < rv.v {
			t.v = lv.v
			t.idx = lv.idx
		} else {
			t.v = rv.v
			t.idx = rv.idx
		}
		return t
	}
	ti := T{v: 1<<31 - 1, idx: -1}
	st := NewSegmentTree(n, f, ti)
	for i := 0; i < n; i++ {
		st.Set(i, T{v: A[i], idx: i})
	}
	st.Build()

	for i := 0; i < q; i++ {
		c, l, r := ReadInt3()
		if c == 1 {
			ol := st.Get(l - 1)
			or := st.Get(r - 1)
			ol.idx, or.idx = or.idx, ol.idx
			st.Update(l-1, or)
			st.Update(r-1, ol)
		} else {
			e := st.Query(l-1, r)
			fmt.Println(e.idx + 1)
		}
	}
}
```

### 遅延伝搬あり

```go
// Assumption: T == E
type T int // (T, f): Monoid
type E int // (E, h): Operator Monoid

type LazySegmentTree struct {
  sz   int
  data []T
  lazy []E
  f    func(lv, rv T) T        // T <> T -> T
  g    func(to T, from E) T    // T <> E -> T (assignment operator)
  h    func(to, from E) E      // E <> E -> E (assignment operator)
  p    func(e E, length int) E // E <> N -> E
  ti   T
  ei   E
}

func NewLazySegmentTree(
  n int,
  f func(lv, rv T) T, g func(to T, from E) T,
  h func(to, from E) E, p func(e E, length int) E,
  ti T, ei E,
) *LazySegmentTree {
  lst := new(LazySegmentTree)
  lst.f, lst.g, lst.h, lst.p = f, g, h, p
  lst.ti, lst.ei = ti, ei

  lst.sz = 1
  for lst.sz < n {
    lst.sz *= 2
  }

  lst.data = make([]T, 2*lst.sz-1)
  lst.lazy = make([]E, 2*lst.sz-1)
  for i := 0; i < 2*lst.sz-1; i++ {
    lst.data[i] = lst.ti
    lst.lazy[i] = lst.ei
  }

  return lst
}

func (lst *LazySegmentTree) Set(k int, x T) {
  lst.data[k+(lst.sz-1)] = x
}

func (lst *LazySegmentTree) Build() {
  for i := lst.sz - 2; i >= 0; i-- {
    lst.data[i] = lst.f(lst.data[2*i+1], lst.data[2*i+2])
  }
}

func (lst *LazySegmentTree) propagate(k, length int) {
  if lst.lazy[k] != lst.ei {
    if k < lst.sz-1 {
      lst.lazy[2*k+1] = lst.h(lst.lazy[2*k+1], lst.lazy[k])
      lst.lazy[2*k+2] = lst.h(lst.lazy[2*k+2], lst.lazy[k])
    }
    lst.data[k] = lst.g(lst.data[k], lst.p(lst.lazy[k], length))
    lst.lazy[k] = lst.ei
  }
}

func (lst *LazySegmentTree) Update(a, b int, x E) T {
  return lst.update(a, b, x, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) update(a, b int, x E, k, l, r int) T {
  lst.propagate(k, r-l)

  if r <= a || b <= l {
    return lst.data[k]
  }

  if a <= l && r <= b {
    lst.lazy[k] = lst.h(lst.lazy[k], x)
    lst.propagate(k, r-l)
    return lst.data[k]
  }

  lv := lst.update(a, b, x, 2*k+1, l, (l+r)/2)
  rv := lst.update(a, b, x, 2*k+2, (l+r)/2, r)
  lst.data[k] = lst.f(lv, rv)
  return lst.data[k]
}

func (lst *LazySegmentTree) Query(a, b int) T {
  return lst.query(a, b, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) query(a, b, k, l, r int) T {
  lst.propagate(k, r-l)

  if r <= a || b <= l {
    return lst.ti
  }

  if a <= l && r <= b {
    return lst.data[k]
  }

  lv := lst.query(a, b, 2*k+1, l, (l+r)/2)
  rv := lst.query(a, b, 2*k+2, (l+r)/2, r)
  return lst.f(lv, rv)
}

func (lst *LazySegmentTree) Get(k int) T {
  return lst.Query(k, k+1)
}
```

大抵はそうだろうということで、混乱しないように `Assumption: T == E` みたいなことを書いてしまいました。
しかし、次の例に示すように、別にそうとも限らないですね。。

こちらもモノイド、作用素モノイド、
そしてそれらに関する関数オブジェクトを適切に決定することに注力しさえすればいいように書いたつもりです。

#### 例: yukicoder No.876 Range Compress Query

[問題のURL](https://yukicoder.me/problems/no/876)

yukicoderの解説での想定解法は「階差数列に着目して区間加算を2箇所の点加算で済むようにする」
というようなもの((そちらの解法で解き直していないので間違っているかもしれません。))のようです。
しかし、区間加算を真に受けて遅延伝搬セグメントツリーでも解けます。

与えられた定義で計算される圧縮値の他、区間の両端の値をもたせた構造体を要素の型とすればOKです。
単位元は少し注意が必要です。

以下はコードの抜粋です（提出は[こちら](https://yukicoder.me/submissions/419477)）。

```go
const INF_BIT60 = 1 << 60

func main() {
	n, q := ReadInt2()
	A := ReadIntSlice(n)

	f := func(lv, rv T) T {
		t := T{}
		t.v = lv.v + rv.v
		if lv.r >= INF_BIT60 || rv.l >= INF_BIT60 {
		} else if lv.r != rv.l {
			t.v++
		}
		t.l, t.r = lv.l, rv.r
		return t
	}
	g := func(to T, from E) T {
		to.l += int(from)
		to.r += int(from)
		return to
	}
	h := func(to, from E) E {
		return to + from
	}
	p := func(e E, length int) E {
		return e
	}
	ti := T{v: 0, l: INF_BIT60, r: INF_BIT60}
	ei := 0
	lst := NewLazySegmentTree(n, f, g, h, p, ti, E(ei))

	for i := 0; i < n; i++ {
		lst.Set(i, T{v: 0, l: A[i], r: A[i]})
	}
	lst.Build()

	for i := 0; i < q; i++ {
		c := ReadInt()
		if c == 1 {
			l, r, x := ReadInt3()
			lst.Update(l-1, r, E(x))
		} else {
			l, r := ReadInt2()
			t := lst.Query(l-1, r)
			fmt.Println(t.v + 1)
		}
	}
}

type T struct {
	v, l, r int
}

type E int
```

## 参考

学習の高速道路が整備されていてありがたいなぁと感じました。

### tsutajさんの解説

[tsutajさんの解説記事（遅延伝搬なし）](http://tsutaj.hatenablog.com/entry/2017/03/29/204841)

[tsutajさんの解説記事（遅延伝搬あり）](http://tsutaj.hatenablog.com/entry/2017/03/30/224339)

いずれもとてもわかりやすく、初学時にセグメントツリーの仕組みを理解するのにはかどりました((蟻本だけでは不安だったところを補完するのに活用させていただきました))。

セグメント木の抽象化だけに絞るのであれば、tsutajさんのブログを理解すれば十分な気がします。

### ei1333さん・beetさんの解説

[ei1333さんの旧ブログ記事での解説](https://ei1333.github.io/algorithm/segment-tree.html)

[beetさんの旧記事での解説](http://beet-aizu.hatenablog.com/entry/2017/12/01/225955)

お二方とも旧ブログの方を参照してしまって恐縮です。
特にbeetさんの記事の方は本当に丁寧に説明されていてわかりやすかったです。

---

少数の定義を渡すだけで正しく動いてくれて感動したので、個人的には満足しています。

まだ解いた問題が少なすぎるので「これでOK」とは言い難いですが、
とりあえずは運用して様子見してみようと思います。

