package ibm

import "fmt"

type Sample struct {
	E, F []string
}

func (m *model) Train(d []Sample, n int) {
	c := make(map[string]map[string]float64)

	for i := 0; i < n; i++ {
		for _, v := range m.VF {
			c[v] = make(map[string]float64)

			for _, u := range m.VE {
				c[v][u] = 0
			}
		}

		for _, s := range d {
			for _, e := range s.E {
				sum := float64(0)

				for _, f := range s.F {
					sum = sum + m.t[f][e]
				}

				for _, f := range s.F {
					c[f][e] = c[f][e] + m.t[f][e]/sum
				}
			}
		}

		fmt.Printf("c: %v\n", c)

		for _, v := range m.VF {
			var sum float64

			for _, u := range m.VE {
				sum = sum + c[v][u]
			}

			for _, u := range m.VE {
				m.t[v][u] = c[v][u] / sum
			}
		}

		fmt.Printf("t: %v\n", m.t)
	}
}
