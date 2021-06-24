package main

import (
	"github.com/jonasknobloch/2021ss-smt/pkg/ibm"
)

func main() {
	VE := []string{
		"Kinder",
		"lasst",
		"uns",
		"spielen",
	}

	VF := []string{
		"children",
		"let's",
		"play",
	}

	t := make(map[string]map[string]float64)

	for _, vf := range VF {
		t[vf] = make(map[string]float64)
	}

	for _, vf := range VF {
		for _, ve := range VE {
			t[vf][ve] = 1 / float64(len(VE))
			t[vf][ve] = 1 / float64(len(VE))
		}
	}

	m := ibm.NewModel(VE, VF, t, nil)

	d := []ibm.Sample{
		{
			E: []string{
				"Kinder",
				"spielen",
			},
			F: []string{
				"children",
				"play",
			},
		},
		{
			E: []string{
				"lasst",
				"uns",
				"spielen",
			},
			F: []string{
				"let's",
				"play",
			},
		},
	}

	m.Train(d, 2)
}
