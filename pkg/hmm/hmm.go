package hmm

import (
	"errors"
)

type hmm struct {
	states       []string                      // Q
	observations []string                      // V
	final        string                        // #
	PTransition  map[string]map[string]float64 // t
	PEmission    map[string]map[string]float64 // e
}

func NewHMM(Q []string, V []string, f string) *hmm {
	h := hmm{
		states:       Q,
		observations: V,
		final:        f,
		PTransition:  make(map[string]map[string]float64),
		PEmission:    make(map[string]map[string]float64),
	}

	h.PTransition[h.final] = make(map[string]float64)

	for _, q := range h.states {
		h.PTransition[q] = make(map[string]float64)
		h.PEmission[q] = make(map[string]float64)
	}

	for _, q1 := range h.states {
		h.PTransition[h.final][q1] = 0
		h.PTransition[q1][h.final] = 0

		for _, q2 := range h.states {
			h.PTransition[q1][q2] = 0
		}
	}

	for _, q := range h.states {
		for _, v := range h.observations {
			h.PEmission[q][v] = 0
		}
	}

	return &h
}

func (h *hmm) Validate() error {
	Qf := append(h.states, h.final)

	for _, q1 := range Qf {
		var sum float64

		for _, q2 := range Qf {
			sum = sum + h.PTransition[q1][q2]
		}

		if sum != float64(1) {
			return errors.New("invalid transition probabilities for " + q1)
		}
	}

	for _, q := range h.states {
		var sum float64

		for _, v := range h.observations {
			sum = sum + h.PEmission[q][v]
		}

		if sum != float64(1) {
			return errors.New("invalid emission probabilities for " + q)
		}
	}

	return nil
}

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

func Forward(V []string, h *hmm) float64 {
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

	return sum
}

func Backward(V []string, h *hmm) float64 {
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

	return sum
}

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
