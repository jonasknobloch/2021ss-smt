package hmm

func Viterbi(V []string, h *hmm) []string {
	T := make(map[int]map[string]float64)

	for i := range V {
		T[i+1] = make(map[string]float64)
	}

	psi := make(map[int]map[string]string)

	for t := 2; t <= len(V); t++ {
		psi[t] = make(map[string]string)
	}

	for _, q := range h.states {
		T[1][q] = h.PEmission[q][V[0]] * h.PTransition[h.final][q]
	}

	for t := 2; t <= len(V); t++ {
		for _, q := range h.states {
			var max float64
			var argMax string

			for _, qq := range h.states {
				if p := h.PTransition[qq][q] * T[t-1][qq]; p > max {
					max = p
					argMax = qq
				}
			}

			T[t][q] = h.PEmission[q][V[t-1]] * max
			psi[t][q] = argMax
		}
	}

	var max float64
	var argMax string

	for _, q := range h.states {
		if p := h.PTransition[q][h.final] * T[len(V)][q]; p > max {
			max = p
			argMax = q
		}
	}

	Q := make([]string, len(V))
	Q[len(V)-1] = argMax

	for t := len(V) - 1; t >= 1; t-- {
		Q[t-1] = psi[t+1][Q[t]]
	}

	return Q
}
