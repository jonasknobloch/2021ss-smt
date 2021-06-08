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
