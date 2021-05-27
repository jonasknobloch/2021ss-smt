package main

import (
	"errors"
	"fmt"
)

type hmm struct {
	states       []string                      // Q
	observations []string                      // V
	final        string                        // #
	pTransition  map[string]map[string]float64 // t
	pEmission    map[string]map[string]float64 // e
}

func newHMM(Q []string, V []string, f string) *hmm {
	h := hmm{
		states:       Q,
		observations: V,
		final:        f,
		pTransition:  make(map[string]map[string]float64),
		pEmission:    make(map[string]map[string]float64),
	}

	h.pTransition[h.final] = make(map[string]float64)

	for _, q := range h.states {
		h.pTransition[q] = make(map[string]float64)
		h.pEmission[q] = make(map[string]float64)
	}

	for _, q1 := range h.states {
		h.pTransition[h.final][q1] = 0
		h.pTransition[q1][h.final] = 0

		for _, q2 := range h.states {
			h.pTransition[q1][q2] = 0
		}
	}

	for _, q := range h.states {
		for _, v := range h.observations {
			h.pEmission[q][v] = 0
		}
	}

	return &h
}

func (h *hmm) validate() error {
	Qf := append(h.states, h.final)

	for _, q1 := range Qf {
		var sum float64

		for _, q2 := range Qf {
			sum = sum + h.pTransition[q1][q2]
		}

		if sum != float64(1) {
			return errors.New("invalid transition probabilities for " + q1)
		}
	}

	for _, q := range h.states {
		var sum float64

		for _, v := range h.observations {
			sum = sum + h.pEmission[q][v]
		}

		if sum != float64(1) {
			return errors.New("invalid emission probabilities for " + q)
		}
	}

	return nil
}

func main() {
	Q := []string{
		"q1", "q2",
	}

	V := []string{
		"a", "b", "y",
	}

	h := newHMM(Q, V, "#")

	h.pTransition["#"]["q1"] = 1
	h.pTransition["q1"]["q1"] = 0.5
	h.pTransition["q1"]["q2"] = 0.5
	h.pTransition["q2"]["q2"] = 0.5
	h.pTransition["q2"]["#"] = 0.5

	h.pEmission["q1"]["a"] = 0.5
	h.pEmission["q1"]["b"] = 0.5
	h.pEmission["q2"]["b"] = 0.5
	h.pEmission["q2"]["y"] = 0.5

	if err := h.validate(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("valid hmm initialized")

	fmt.Printf("[forward] P(V=aby) = %f\n", forward([]string{"a", "b", "y"}, h))
	fmt.Printf("[forward] P(V=abby) = %f\n", forward([]string{"a", "b", "b", "y"}, h))

	fmt.Printf("[backward] P(V=aby) = %f\n", backward([]string{"a", "b", "y"}, h))
	fmt.Printf("[backward] P(V=abby) = %f\n", backward([]string{"a", "b", "b", "y"}, h))
}

func forward(V []string, h *hmm) float64 {
	T := make(map[int]map[string]float64)

	for i := range V {
		T[i+1] = make(map[string]float64)
	}

	for _, q := range h.states {
		T[1][q] = h.pEmission[q][V[0]] * h.pTransition[h.final][q]
	}

	for t := 2; t <= len(V); t++ {
		for _, q := range h.states {
			var sum float64

			for _, qq := range h.states {
				sum = sum + h.pTransition[qq][q]*T[t-1][qq]
			}

			T[t][q] = h.pEmission[q][V[t-1]] * sum
		}
	}

	var sum float64

	for _, q := range h.states {
		sum = sum + h.pTransition[q][h.final]*T[len(V)][q]
	}

	return sum
}

func backward(V []string, h *hmm) float64 {
	S := make(map[int]map[string]float64)

	for i := range V {
		S[i+1] = make(map[string]float64)
	}

	for _, q := range h.states {
		S[len(V)][q] = h.pTransition[q][h.final]
	}

	for t := len(V) - 1; t >= 1; t-- {
		for _, q := range h.states {
			var sum float64

			for _, qq := range h.states {
				sum = sum + h.pEmission[qq][V[t]]*h.pTransition[q][qq]*S[t+1][qq]
			}

			S[t][q] = sum
		}
	}

	var sum float64

	for _, q := range h.states {
		sum = sum + h.pEmission[q][V[0]]*h.pTransition[h.final][q]*S[1][q]
	}

	return sum
}
