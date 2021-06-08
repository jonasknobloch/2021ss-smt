package main

import (
	"fmt"
	"github.com/jonasknobloch/2021ss-smt/pkg/hmm"
)

func main() {
	Q := []string{
		"q1", "q2",
	}

	V := []string{
		"a", "b", "y",
	}

	h := hmm.NewHMM(Q, V, "#")

	h.PTransition["#"]["q1"] = 1
	h.PTransition["q1"]["q1"] = 0.5
	h.PTransition["q1"]["q2"] = 0.5
	h.PTransition["q2"]["q2"] = 0.5
	h.PTransition["q2"]["#"] = 0.5

	h.PEmission["q1"]["a"] = 0.5
	h.PEmission["q1"]["b"] = 0.5
	h.PEmission["q2"]["b"] = 0.5
	h.PEmission["q2"]["y"] = 0.5

	if err := h.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("valid hmm initialized")

	f1, _ := hmm.Forward([]string{"a", "b", "y"}, h)
	fmt.Printf("[forward] P(V=aby) = %f\n", f1)

	f2, _ := hmm.Forward([]string{"a", "b", "b", "y"}, h)
	fmt.Printf("[forward] P(V=abby) = %f\n", f2)

	f3, _ := hmm.Forward([]string{"a", "b", "b", "b", "y"}, h)
	fmt.Printf("[forward] P(V=abbby) = %f\n", f3)

	b1, _ := hmm.Backward([]string{"a", "b", "y"}, h)
	fmt.Printf("[backward] P(V=aby) = %f\n", b1)

	b2, _ := hmm.Backward([]string{"a", "b", "b", "y"}, h)
	fmt.Printf("[backward] P(V=abby) = %f\n", b2)

	b3, _ := hmm.Backward([]string{"a", "b", "b", "b", "y"}, h)
	fmt.Printf("[backward] P(V=abbby) = %f\n", b3)

	fmt.Printf("[naive] P(V=aby) = %f\n", hmm.Naive([]string{"a", "b", "y"}, h))
	fmt.Printf("[naive] P(V=abby) = %f\n", hmm.Naive([]string{"a", "b", "b", "y"}, h))
	fmt.Printf("[naive] P(V=abbby) = %f\n", hmm.Naive([]string{"a", "b", "b", "b", "y"}, h))

	fmt.Printf("[viterbi] P(V=aby) = %s\n", hmm.Viterbi([]string{"a", "b", "y"}, h))
	fmt.Printf("[viterbi] P(V=abby) = %s\n", hmm.Viterbi([]string{"a", "b", "b", "y"}, h))
	fmt.Printf("[viterbi] P(V=abbby) = %s\n", hmm.Viterbi([]string{"a", "b", "b", "b", "y"}, h))
}
