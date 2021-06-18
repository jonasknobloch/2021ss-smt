package scm

import "fmt"

func (s *scm) Decode(f []string) (e []string) {
	Y := make([][]string, 0, 5)
	y := []string{}

	for len(y) == 0 || y[len(y)-1] != s.final {
		for _, w := range append(s.VE, s.final) {
			Y = append(Y, append(y, w))

			if len(Y) > s.YMax {
				var wk int
				var ws float64

				ws = 1

				for k, y := range Y {
					if p := s.score(y, f); p < ws {
						wk = k
						ws = p
					}
				}

				fmt.Printf("Pruning worst hypothesis %v with a score of %f\n", Y[wk], ws)

				Y[wk] = Y[len(Y)-1]
				Y = Y[:len(Y)-1]
			}
		}

		var bk int
		var bs float64

		for k, y := range Y {
			p := s.score(y, f)
			if p > bs {
				bk = k
				bs = p
			}
		}

		fmt.Printf("Selecting best hypothesis %v with a score of %f\n", Y[bk], bs)

		y = Y[bk]
		Y[bk] = Y[len(Y)-1]
		Y = Y[:len(Y)-1]
	}

	return y[:len(y)-1]
}
