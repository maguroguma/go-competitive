package modint

// Originated from @ccppjsrb
// e.g.: https://atcoder.jp/contests/arc106/submissions/17669427

// Mod constants.
const (
	Mod1000000007 = 1000000007
	Mod998244353  = 998244353
)

var _mod Mint
var _fmod func(Mint) Mint

// Mint treats the modular arithmetic
type Mint int64

// SetMod sets the mod. It must be called first.
func SetMod(newmod Mint) {
	switch newmod {
	case Mod1000000007:
		_fmod = staticMod1000000007
	case Mod998244353:
		_fmod = staticMod998244353
	default:
		_mod = newmod
		_fmod = dynamicMod
	}
}
func dynamicMod(m Mint) Mint {
	m %= _mod
	if m < 0 {
		return m + _mod
	}
	return m
}
func staticMod1000000007(m Mint) Mint {
	m %= Mod1000000007
	if m < 0 {
		return m + Mod1000000007
	}
	return m
}
func staticMod998244353(m Mint) Mint {
	m %= Mod998244353
	if m < 0 {
		return m + Mod998244353
	}
	return m
}

// Mod returns m % mod.
func (m Mint) Mod() Mint {
	return _fmod(m)
}

// Inv returns modular multiplicative inverse
func (m Mint) Inv() Mint {
	return m.Pow(Mint(0).Sub(2))
}

// Pow returns m^n
func (m Mint) Pow(n Mint) Mint {
	p := Mint(1)
	for n > 0 {
		if n&1 == 1 {
			p.MulAs(m)
		}
		m.MulAs(m)
		n >>= 1
	}
	return p
}

// Add returns m+x
func (m Mint) Add(x Mint) Mint {
	return (m + x).Mod()
}

// Sub returns m-x
func (m Mint) Sub(x Mint) Mint {
	return (m - x).Mod()
}

// Mul returns m*x
func (m Mint) Mul(x Mint) Mint {
	return (m * x).Mod()
}

// Div returns m/x
func (m Mint) Div(x Mint) Mint {
	return m.Mul(x.Inv())
}

// AddAs assigns *m + x to *m and returns m
func (m *Mint) AddAs(x Mint) *Mint {
	*m = m.Add(x)
	return m
}

// SubAs assigns *m - x to *m and returns m
func (m *Mint) SubAs(x Mint) *Mint {
	*m = m.Sub(x)
	return m
}

// MulAs assigns *m * x to *m and returns m
func (m *Mint) MulAs(x Mint) *Mint {
	*m = m.Mul(x)
	return m
}

// DivAs assigns *m / x to *m and returns m
func (m *Mint) DivAs(x Mint) *Mint {
	*m = m.Div(x)
	return m
}

// cf := NewCombFactorial(2000000)
// maxNum == "maximum n" * 2 (for H(n,r))
// res := cf.C(n, r) 	// 組み合わせ
// res := cf.H(n, r) 	// 重複組合せ
// res := cf.P(n, r) 	// 順列

type CombFactorial struct {
	maxNum Mint
	fact   func(x Mint) Mint
	invf   func(x Mint) Mint
}

func NewCombFactorial(maxNum Mint) *CombFactorial {
	cf := new(CombFactorial)
	cf.maxNum = maxNum
	cf.initCF()

	return cf
}

func (c *CombFactorial) initCF() {
	var i Mint

	factTable := make([]Mint, c.maxNum+50)
	invfTable := make([]Mint, c.maxNum+50)

	factTable[0] = 1
	invfTable[0] = factTable[0].Inv()
	for i = 1; i <= c.maxNum; i++ {
		val := factTable[i-1].Mul(Mint(i))
		factTable[i] = val
		invfTable[i] = factTable[i].Inv()
	}

	c.fact = func(x Mint) Mint { return factTable[x] }
	c.invf = func(x Mint) Mint { return invfTable[x] }
}
func (c *CombFactorial) C(n, r Mint) Mint {
	var res Mint

	res = Mint(1).
		Mul(c.fact(n)).
		Mul(c.invf(r)).
		Mul(c.invf(n - r))

	return res
}
func (c *CombFactorial) P(n, r Mint) Mint {
	var res Mint

	res = 1
	res.MulAs(c.fact(n)).MulAs(c.invf(n - r))

	return res
}
func (c *CombFactorial) H(n, r Mint) Mint {
	return c.C(n-1+r, r)
}
