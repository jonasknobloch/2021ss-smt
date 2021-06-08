package hmm

func (h *hmm) Backward(V []string) (float64, map[int]map[string]float64) {
	S := make(map[int]map[string]float64)

	for i := range V {
		S[i+1] = make(map[string]float64)
	}

	for _, q := range h.states {
		S[len(V)][q] = h.PTransition[q][h.final]
	}

	for t := len(V) - 1; t >= 1; t-- {
		for _, q := range h.states {
			var sum float64

			for _, qq := range h.states {
				sum = sum + h.PEmission[qq][V[t]]*h.PTransition[q][qq]*S[t+1][qq]
			}

			S[t][q] = sum
		}
	}

	var sum float64

	for _, q := range h.states {
		sum = sum + h.PEmission[q][V[0]]*h.PTransition[h.final][q]*S[1][q]
	}

	return sum, S
}
