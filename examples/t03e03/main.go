package main

import (
	"fmt"
	"github.com/jonasknobloch/2021ss-smt/pkg/scm"
)

func main() {
	VE := []string{"a", "b"}
	VF := []string{"α", "β", "γ"}

	b := make(map[string]map[string]float64)
	t := make(map[string]map[string]float64)

	final := "#"
	YMax := 5

	b["#"] = make(map[string]float64)

	for _, ve := range VE {
		b[ve] = make(map[string]float64)
		t[ve] = make(map[string]float64)
	}

	b["#"]["#"] = 0.5
	b["#"]["a"] = 0.5
	b["#"]["b"] = 0
	b["a"]["#"] = 0
	b["a"]["a"] = 0.5
	b["a"]["b"] = 0.5
	b["b"]["#"] = 0.5
	b["b"]["a"] = 0
	b["b"]["b"] = 0.5

	t["a"]["α"] = 0.5
	t["a"]["β"] = 0
	t["a"]["γ"] = 0.5
	t["b"]["α"] = 0
	t["b"]["β"] = 0.5
	t["b"]["γ"] = 0.5

	l := func(l, m int) float64 {
		if m == l {
			return 0.5
		}

		if l-m == 1 || m-l == 1 {
			return 0.25
		}

		return 0
	}

	lMax := 3

	s := scm.NewSCM(VE, VF, final, b, t, l, lMax, YMax)
	f := []string{"β", "γ"}
	fmt.Println(s.Decode(f))
}
