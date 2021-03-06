package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	S, T := "strange", "orange"

	InitRollingHashConfing()
	sh, th := NewRollingHash(S), NewRollingHash(T)

	for i := 0; i < len(S); i++ {
		for j := i + 1; j <= len(S); j++ {
			for a := 0; a < len(T); a++ {
				for b := a + 1; b <= len(T); b++ {
					slen, tlen := j-i, b-a
					ssub, tsub := string(S[i:j]), string(T[a:b])
					sval, tval := sh.Slice(i, slen), th.Slice(a, tlen)

					if sval == tval {
						fmt.Printf("S: %s\tT: %s -> %d\n", ssub, tsub, sval)
					}
				}
			}
		}
	}
}

// rolling hash
// reference: https://atcoder.jp/contests/abc141/submissions/7717102

const (
	MASK30       uint64 = (1 << 30) - 1
	MASK31       uint64 = (1 << 31) - 1
	R_MOD        uint64 = (1 << 61) - 1
	POSITIVIZER  uint64 = R_MOD * ((1 << 3) - 1)
	MAX_S_LENGTH        = 200000 + 50
)

var (
	Base    uint64
	PowMemo []uint64
)

type RollingHash struct {
	hash []uint64
}

func InitRollingHashConfing() {
	rand.Seed(time.Now().UnixNano())

	Base = uint64(rand.Int31n(math.MaxInt32-129)) + uint64(129)
	PowMemo = make([]uint64, MAX_S_LENGTH)
	PowMemo[0] = 1
	for i := 1; i < len(PowMemo); i++ {
		PowMemo[i] = CalcMod(Mul(PowMemo[i-1], Base))
	}
}

func NewRollingHash(s string) *RollingHash {
	rlh := new(RollingHash)

	rlh.hash = make([]uint64, len(s)+1)
	for i := 0; i < len(s); i++ {
		rlh.hash[i+1] = CalcMod(Mul(rlh.hash[i], Base) + uint64(s[i]))
	}

	return rlh
}

func (rlh *RollingHash) Slice(begin, length int) uint64 {
	return CalcMod(rlh.hash[begin+length] + POSITIVIZER - Mul(rlh.hash[begin], PowMemo[length]))
}

func Mul(l, r uint64) uint64 {
	var lu uint64 = l >> 31
	var ld uint64 = l & MASK31
	var ru uint64 = r >> 31
	var rd uint64 = r & MASK31
	var middleBit uint64 = ld*ru + lu*rd

	return ((lu * ru) << 1) + ld*rd + ((middleBit & MASK30) << 31) + (middleBit >> 30)
}

func CalcMod(val uint64) uint64 {
	val = (val & R_MOD) + (val >> 61)
	if val > R_MOD {
		val -= R_MOD
	}
	return val
}
