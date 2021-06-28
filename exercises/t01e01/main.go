package main

import (
	"fmt"
	"github.com/jonasknobloch/2021ss-smt/pkg/ibm"
	"math"
	"strings"
)

func main() {
	VE := []string{
		"kra",
		"ban",
		"las",
		"gha",
	}

	VF := []string{
		"du",
		"su",
		"ur",
		"fur",
	}

	t := make(map[string]map[string]float64)

	for _, vf := range VF {
		t[vf] = make(map[string]float64)
	}

	t["du"]["kra"] = 0.2
	t["du"]["ban"] = 0.4
	t["du"]["las"] = 0.4
	t["du"]["gha"] = 0

	t["su"]["kra"] = 0
	t["su"]["ban"] = 0.1
	t["su"]["las"] = 0.8
	t["su"]["gha"] = 0.1

	t["ur"]["kra"] = 0.3
	t["ur"]["ban"] = 0.4
	t["ur"]["las"] = 0.25
	t["ur"]["gha"] = 0.25

	t["fur"]["kra"] = 0.4
	t["fur"]["ban"] = 0.3
	t["fur"]["las"] = 0.1
	t["fur"]["gha"] = 0.2

	l1 := func(m, l int) float64 {
		if m == l {
			return 0.5
		}

		if math.Abs(float64(m-l)) == 1 {
			return 0.25
		}

		return 0
	}

	l2 := func(m, l int) float64 {
		if l > 0 {
			return math.Pow(2, -float64(l))
		}

		return 0
	}

	l3 := func(m, l int) float64 {
		if m == l {
			return 1
		}

		return 0
	}

	ma := ibm.NewModel(VE, VF, t, l1)

	fa1 := strings.Split("du su ur", " ")
	ea1 := strings.Split("kra las gha", " ")

	fmt.Printf("P(%v, %v) = %f\n", fa1, ea1, ma.P(fa1, ea1))

	fa2 := strings.Split("du su", " ")
	ea2 := strings.Split("kra las gha", " ")

	fmt.Printf("P(%v, %v) = %f\n", fa2, ea2, ma.P(fa2, ea2))

	fa3 := strings.Split("su du su", " ")
	ea3 := strings.Split("gha gha", " ")

	fmt.Printf("P(%v, %v) = %f\n", fa2, ea3, ma.P(fa3, ea3))

	fb1 := strings.Split("du", " ")
	fb2 := strings.Split("du su ur", " ")
	fb3 := strings.Split("fur du fur su fur", " ")

	fmt.Println()

	mb := ibm.NewModel(VE, VF, t, l2)

	eb1, pb1 := mb.Decode(fb1)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb1, eb1, pb1)

	eb2, pb2 := mb.Decode(fb2)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb2, eb2, pb2)

	eb3, pb3 := mb.Decode(fb3)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb3, eb3, pb3)

	fmt.Println()

	mc := ibm.NewModel(VE, VF, t, l3)

	ec1, pc1 := mc.Decode(fb1)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb1, ec1, pc1)

	ec2, pc2 := mc.Decode(fb2)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb2, ec2, pc2)

	ec3, pc3 := mc.Decode(fb3)
	fmt.Printf("Decode(%v) = %v (%f)\n", fb3, ec3, pc3)
}
