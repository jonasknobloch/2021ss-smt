package hmm

func Naive(V []string, h *hmm) float64 {
	var QQ [][]string

	var perm func(QQ []string, k int)

	perm = func(Q []string, k int) {
		if k == 0 {
			CQ := make([]string, len(Q))
			copy(CQ, Q)
			QQ = append(QQ, CQ)

			return
		}

		for i := 0; i < len(h.states); i++ {
			perm(append(Q, h.states[i]), k-1)
		}
	}

	perm([]string{}, len(V))

	var sum float64

	for _, Q := range QQ {
		prd := float64(1)

		for t := 2; t <= len(Q); t++ {
			prd = prd * h.PTransition[Q[t-2]][Q[t-1]] * h.PEmission[Q[t-1]][V[t-1]]
		}

		sum = sum + h.PTransition[h.final][Q[0]]*h.PTransition[Q[len(Q)-1]][h.final]*h.PEmission[Q[0]][V[0]]*prd
	}

	return sum
}
