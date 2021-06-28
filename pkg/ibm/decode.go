package ibm

import (
	"math"
)

func (m *model) P(f, e []string) float64 {
	prod := float64(1)

	for _, e := range e {
		var sum float64

		for _, f := range f {
			sum = sum + m.t[f][e]
		}

		prod = prod * sum
	}

	return 1 / math.Pow(float64(len(f)), float64(len(e))) * m.l(len(f), len(e)) * prod
}

func (m *model) Decode(f []string) (e []string, p float64) {
	var Y [][]string // hypothesis set

	var kleene func(y []string, k int)

	kleene = func(y []string, k int) {
		cy := make([]string, len(y))
		copy(cy, y)
		Y = append(Y, cy)

		if k == 0 {
			return
		}

		for i := 0; i < len(m.VE); i++ {
			kleene(append(y, m.VE[i]), k-1)
		}
	}

	kleene([]string{}, len(f))

	for _, y := range Y {
		if pp := m.P(f, y); pp > p {
			e = y  // arg
			p = pp // max
		}
	}

	return
}
