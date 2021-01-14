package csum

// Imos manages event processing.
type Imos struct {
	ts []int64
}

func NewImos(maxT int) *Imos {
	im := new(Imos)
	im.ts = make([]int64, maxT+1)
	return im
}

// AddEvent process an event by adding an event value to timestamp t.
func (im *Imos) AddEvent(t int, ev int64) {
	im.ts[t] += ev
}

// Build simulates all registered events, and then return results.
func (im *Imos) Build() []int64 {
	n := len(im.ts)
	for i := 1; i < n; i++ {
		im.ts[i] += im.ts[i-1]
	}

	res := make([]int64, len(im.ts))
	copy(res, im.ts)

	return res
}
