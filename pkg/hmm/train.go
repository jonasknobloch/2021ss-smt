package hmm

type Sample struct {
	V []string
	N float64
}

func (h *hmm) Train(c []Sample, i int) (t, e []map[string]map[string]float64) {
	U := func(V []string, t int, qi, qj string) float64 {
		p, T := h.Forward(V)
		_, S := h.Backward(V)

		if t == len(V)-1 {
			return T[t][qi] * h.PTransition[qi][qj] * h.PEmission[qj][V[len(V)-1]] * h.PTransition[qj][h.final] / p
		}

		return T[t][qi] * h.PTransition[qi][qj] * h.PEmission[qj][V[t]] * S[t+1][qj] / p
	}

	R := func(V []string, t int, q string) float64 {
		p, T := h.Forward(V)
		_, S := h.Backward(V)

		if t == len(V) {
			return T[len(V)][q] * h.PTransition[q][h.final] / p
		}

		return T[t][q] * S[t][q] / p
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
