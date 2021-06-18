package scm

import (
	"fmt"
	"math"
)

func (s *scm) scoreLM(l, k int, e []string) (p float64) {
	p = s.b[s.final][e[0]]

	if k == 0 {
		return p
	}

	for i := 1; i < k; i++ {
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
	switch l - k {
	case 0:
		a := s.b[e[k-1]][s.final]
		p = p * a
	case 1:
		a := max(e[k-1:], s.VE)
		b := max(s.VE, []string{s.final})
		p = p * a * b
	default:
		a := max(e[k-1:], s.VE)
		b := math.Pow(max(s.VE, s.VE), float64(l-k-1))
		c := max(s.VE, []string{s.final})
		p = p * a * b * c
	}

	return
}

func (s *scm) scoreTM(l, k int, e, f []string) (p float64) {
	if k == 0 {
		return
	}

	p = 1

	for _, ff := range f {
		var max float64
		for _, ee := range e[:k] {
			if p := s.t[ee][ff]; p > max {
				max = p
			}
		}

		var sum float64
		for _, ee := range e[:k] {
			a := s.t[ee][ff]
			b := float64(l-k) * max
			sum = sum + a + b
		}

		p = p * sum // sum can be 0
	}

	// TODO: formatting
	a := 1 / math.Pow(float64(l), float64(len(f))) // l >= k -> l > 1
	b := s.l(l, len(f))
	p = p * a * b

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
		k := len(e)
		ll := l

		if e[len(e)-1] == s.final {
			k--
			ll = k
		}

		lm := s.scoreLM(ll, k, e)
		fmt.Printf("scoreLM(%d, %v) = %f\n", ll, e, lm)

		tm := s.scoreTM(ll, k, e, f)
		fmt.Printf("scoreTM(%d, %v) = %f\n", ll, e, tm)

		if p := lm * tm; p > max {
			max = p
		}
	}

	fmt.Printf("score(%v) = %f\n", e, max)

	return max
}
