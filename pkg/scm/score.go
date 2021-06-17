package scm

import (
	"fmt"
	"math"
)

func (s *scm) scoreLM(l int, e []string) (p float64) {
	p = s.b[s.final][e[0]]

	for i := 1; i < len(e); i++ {
		p = p * s.b[e[i-1]][e[i]]
	}

	max := func(k1, k2 []string) (m float64) {
		for _, i := range k1 {
			for _, j := range k2 {
				if p := s.b[i][j]; p > m {
					m = p
				}
			}
		}

		return
	}

	// TODO: formatting
	switch l - len(e) {
	case 0:
		a := s.b[e[len(e)-1]][s.final]
		p = p * a
	case 1:
		a := max(e[len(e)-1:], s.VE)
		b := max(s.VE, []string{s.final})
		p = p * a * b
	default:
		a := max(e[len(e)-1:], s.VE)
		b := math.Pow(max(s.VE, s.VE), float64(l-len(e)-1))
		c := max(s.VE, []string{s.final})
		p = p * a * b * c
	}

	fmt.Printf("scoreLM(%d, %v) = %f\n", l, e, p)

	return
}

func (s *scm) scoreTM(l int, e, f []string) (p float64) {
	p = 1

	for _, ff := range f {
		var max float64
		for _, ee := range e {
			if p := s.t[ee][ff]; p > max {
				max = p
			}
		}

		var sum float64
		for _, ee := range e {
			a := s.t[ee][ff]
			b := float64(l-len(e)) * max
			sum = sum + a + b
		}

		p = p * sum // sum can be 0
	}

	// TODO: formatting
	a := 1 / math.Pow(float64(l), float64(len(f)))
	b := s.l(l, len(f))
	p = p * a * b

	fmt.Printf("scoreTM(%d, %v) = %f\n", l, e, p)

	return
}

func (s *scm) score(e, f []string) float64 {

	// input french sentence
	// h(f) = english sentence
	// e_1 ... e_k k >= 0 prefix of potential translation
	// l = full length of translation
	// m = length of french input
	// l_max >= l >= k

	var max float64
	for l := len(e); l <= s.lMax; l++ {
		if p := s.scoreLM(l, e) * s.scoreTM(l, e, f); p > max {
			max = p
		}
	}

	fmt.Printf("score(%v) = %f\n", e, max)

	return max
}
