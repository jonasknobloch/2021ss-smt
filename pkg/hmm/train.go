package hmm

import (
	"fmt"
)

type Sample struct {
	V []string
	N float64
}

func (h *hmm) Train(c []Sample, i int) (t, e []map[string]map[string]float64) {
	U := func(V []string, t int, qi, qj string) (u float64) {
		p, T := h.Forward(V)
		_, S := h.Backward(V)

		if t == len(V)-1 {
			u = T[t][qi] * h.PTransition[qi][qj] * h.PEmission[qj][V[len(V)-1]] * h.PTransition[qj][h.final] / p
		}

		u = T[t][qi] * h.PTransition[qi][qj] * h.PEmission[qj][V[t]] * S[t+1][qj] / p

		fmt.Printf("U(%s, %d, %s, %s) = %f\n", V, t, qi, qj, u)
		fmt.Printf("T(%s, %d, %s) = %f\n", V, t, qi, T[t][qi])
		fmt.Printf("S(%s, %d, %s) = %f\n", V, t, qi, S[t][qi])
		fmt.Printf("P(%s) = %f\n\n", V, p)

		return
	}

	R := func(V []string, t int, q string) (r float64) {
		p, T := h.Forward(V)
		_, S := h.Backward(V)

		if t == len(V) {
			r = T[len(V)][q] * h.PTransition[q][h.final] / p
		}

		r = T[t][q] * S[t][q] / p

		fmt.Printf("R(%s, %d, %s) = %f\n", V, t, q, r)
		fmt.Printf("T(%s, %d, %s) = %f\n", V, t, q, T[t][q])
		fmt.Printf("S(%s, %d, %s) = %f\n", V, t, q, S[t][q])
		fmt.Printf("P(%s) = %f\n\n", V, p)

		return
	}

	t = make([]map[string]map[string]float64, i+1)
	e = make([]map[string]map[string]float64, i+1)

	t[0] = h.PTransition
	e[0] = h.PEmission

	QS := append(h.states, h.final)

	cTr := make(map[string]map[string]float64)
	cEm := make(map[string]map[string]float64)

	for _, q := range QS {
		cTr[q] = make(map[string]float64)
		cEm[q] = make(map[string]float64)
	}

	for k := 1; k <= i; k++ {
		for _, q := range QS {
			for _, qq := range QS {
				cTr[q][qq] = 0
			}

			for _, v := range h.observations {
				cEm[q][v] = 0
			}
		}

		fmt.Printf("k = %d\n\n", k)

		for _, s := range c {
			for t := 1; t < len(s.V); t++ {
				for _, qj := range h.states {
					for _, qi := range h.states {
						cTr[qj][qi] = cTr[qj][qi] + s.N*U(s.V, t, qi, qj)
					}
				}
			}

			for t := 1; t <= len(s.V); t++ {
				for _, q := range h.states {
					cEm[q][s.V[t-1]] = cEm[q][s.V[t-1]] + s.N*R(s.V, t, q)
				}
			}

			for _, q := range h.states {
				cTr[q][h.final] = cTr[q][h.final] + s.N*R(s.V, 1, q)
				cTr[h.final][q] = cTr[h.final][q] + s.N*R(s.V, len(s.V), q)
			}
		}

		t[k] = make(map[string]map[string]float64)

		for _, qj := range QS {
			t[k][qj] = make(map[string]float64)

			for _, qi := range QS {
				var sum float64

				for _, q := range QS {
					sum = sum + cTr[q][qi]
				}

				t[k][qj][qi] = cTr[qj][qi] / sum
			}
		}

		e[k] = make(map[string]map[string]float64)

		for _, q := range h.states {
			e[k][q] = make(map[string]float64)

			for _, v := range h.observations {
				var sum float64

				for _, vv := range h.observations {
					sum = sum + cEm[q][vv]
				}

				e[k][q][v] = cEm[q][v] / sum
			}
		}

		h.PTransition = t[k]
		h.PEmission = e[k]
	}

	return
}
