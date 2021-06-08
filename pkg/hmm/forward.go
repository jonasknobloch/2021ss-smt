package hmm

func Forward(V []string, h *hmm) (float64, map[int]map[string]float64) {
	T := make(map[int]map[string]float64)

	for i := range V {
		T[i+1] = make(map[string]float64)
	}

	for _, q := range h.states {
		T[1][q] = h.PEmission[q][V[0]] * h.PTransition[h.final][q]
	}

	for t := 2; t <= len(V); t++ {
		for _, q := range h.states {
			var sum float64

			for _, qq := range h.states {
				sum = sum + h.PTransition[qq][q]*T[t-1][qq]
			}

			T[t][q] = h.PEmission[q][V[t-1]] * sum
		}
	}

	var sum float64

	for _, q := range h.states {
		sum = sum + h.PTransition[q][h.final]*T[len(V)][q]
	}

	return sum, T
}
