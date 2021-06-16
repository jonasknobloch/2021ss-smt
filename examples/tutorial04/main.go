package main

import (
	"fmt"
	"github.com/jonasknobloch/2021ss-smt/pkg/hmm"
)

func main() {
	Q := []string{
		"DT",
		"JJ",
		"NN",
	}

	V := []string{
		"A",
		"The",
		"tall",
		"boy",
		"house",
	}

	f := "#"

	h := hmm.NewHMM(Q, V, f)

	QS := append(Q, f)

	for _, q1 := range Q {
		for _, q2 := range Q {
			h.PTransition[q1][q2] = 1 / float64(len(QS))
		}

		h.PTransition[f][q1] = 1 / float64(len(Q))
		h.PTransition[q1][f] = 1 / float64(len(Q))
	}

	h.PEmission["DT"]["A"] = 0.35
	h.PEmission["DT"]["The"] = 0.35
	h.PEmission["DT"]["tall"] = 0.1
	h.PEmission["DT"]["boy"] = 0.1
	h.PEmission["DT"]["house"] = 0.1

	h.PEmission["JJ"]["A"] = 0.1
	h.PEmission["JJ"]["The"] = 0.1
	h.PEmission["JJ"]["tall"] = 0.6
	h.PEmission["JJ"]["boy"] = 0.1
	h.PEmission["JJ"]["house"] = 0.1

	h.PEmission["NN"]["A"] = 0.1
	h.PEmission["NN"]["The"] = 0.1
	h.PEmission["NN"]["tall"] = 0.1
	h.PEmission["NN"]["boy"] = 0.35
	h.PEmission["NN"]["house"] = 0.35

	c := []hmm.Sample{
		{
			V: []string{
				"The",
				"house",
			},
			N: 1,
		},
		{
			V: []string{
				"A",
				"tall",
				"boy",
			},
			N: 1,
		},
	}

	t, e := h.Train(c, 10)

	csv := func(d []map[string]map[string]float64, cs []struct{ a, b string }) {
		fmt.Print("Iteration")
		for _, c := range cs {
			fmt.Printf(",[%s]->[%s]", c.a, c.b)
		}
		fmt.Print("\n")

		for i, v := range d {
			fmt.Printf("%d", i)
			for _, c := range cs {
				fmt.Printf(",%f", v[c.a][c.b])
			}
			fmt.Print("\n")
		}
	}

	var ct, ce []struct{ a, b string }

	for _, q := range QS {
		for _, qq := range QS {
			ct = append(ct, struct{ a, b string }{a: q, b: qq})
		}
	}

	fmt.Printf("Transition Probabilities\n")

	csv(t, ct)

	for _, q := range Q {
		for _, v := range V {
			ce = append(ce, struct{ a, b string }{a: q, b: v})
		}
	}

	fmt.Printf("\nEmission Probabilities\n")

	csv(e, ce)
}
