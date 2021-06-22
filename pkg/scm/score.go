package scm

import (
	"fmt"
	"math"
)

func (s *scm) scoreLM(l, k int, e []string) (p float64) {
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

	p = 1

	y := []string{s.final}
	y = append(y, e...)

	for i := 0; i < l-k; i++ {
		y = append(y, "")
	}

	y = append(y, s.final)

	for i, v := range y {
		if i+1 > len(y)-1 {
			break
		}

		if v != "" && y[i+1] == "" {
			p = p * max([]string{v}, s.VE)
			continue
		}

		if v == "" && y[i+1] == "" {
			p = p * max(s.VE, s.VE)
			continue
		}

		if v == "" && y[i+1] != "" {
			p = p * max(s.VE, []string{y[i+1]})
			continue
		}

		p = p * s.b[v][y[i+1]]
	}

	return
}

func (s *scm) scoreTM(l, k int, e, f []string) (p float64) {
	p = 1

	for _, ff := range f {
		var max float64
		for _, ee := range s.VE {
			if p := s.t[ee][ff]; p > max {
				max = p
			}
		}

		var sum float64
		for _, ee := range e[:k] {
			sum = sum + s.t[ee][ff]
		}

		p = p * (sum + float64(l-k)*max)
	}

	if l > 0 {
		p = p * 1 / math.Pow(float64(l), float64(len(f)))
	}

	p = p * s.l(l, len(f))

	return
}

func (s *scm) score(e, f []string) float64 {

	// f = input sentence
	// e = prefix of potential translation (k = len(e))
	// l = predicted length of translation (lMax >= l >= k)

	k := len(e)

	if e[len(e)-1] == s.final {
		k--
	}

	if k > s.lMax {
		fmt.Printf("Hypothesis %s exceeds maximal allowed length of %d\n", e, s.lMax)
		return 0
	}

	var max float64
	for l := k; l <= s.lMax; l++ {
		ll := l

		if e[len(e)-1] == s.final {
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

	return max
}
