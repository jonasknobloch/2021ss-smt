package bleu

import (
	"math"
	"strings"
)

// modified n-gram precision score
func p(n int, e []string, r ...[]string) float64 {

	// R_{n,e}(g) number of occurrences of n-grams g in e
	// max(R_{n, e1}(g), R_{n, e2}(g)) -> higher number of occurrences

	g := func(n int, e []string) map[string]int {
		R := make(map[string]int)

		for i := 0; i+(n-1) < len(e); i++ {
			k := strings.Join(e[i:i+n], "")
			R[k] = R[k] + 1
		}

		return R
	}

	R := make(map[string]int)

	for _, r := range r {
		for k, v := range g(n, r) {
			if R[k] < v {
				R[k] = v
			}
		}
	}

	E := g(n, e)

	var c int

	for k := range E {
		if R[k] > 0 {
			c++
		}
	}

	return float64(c) / float64(len(E))
}

// combined modified precision score
func cp(e []string, r ...[]string) float64 {
	n := 4

	var sum float64

	for i := 1; i <= n; i++ {
		sum = sum + math.Log(p(i, e, r...))
	}

	return math.Exp(sum / float64(n))
}

// effective references length
func erl(e []string, r ...[]string) int {
	var min int
	var max int

	for _, r := range r {
		if ler := len(e) + len(r); max < ler {
			max = ler
		}
	}

	min = max
	l := []int{}

	// 0 <= j < k instead of 0 < j <= k
	for j := 0; j < len(r); j++ {
		k := int(math.Abs(float64(len(e)) - float64(len(r[j]))))

		if k < min {
			min = k
			l = []int{j}
		}

		if k == min {
			l = append(l, j)
		}
	}

	min = max
	var i int

	for _, j := range l {
		k := len(r[j])

		if k < min {
			min = k
			i = j
		}
	}

	return len(r[i])
}

// brevity penalty
func bp(e []string, r ...[]string) float64 {
	erl := erl(e, r...)

	if len(e) > erl {
		return 0
	}

	return math.Exp(float64(1) - float64(erl)/float64(len(e)))
}

func Score(e []string, r ...[]string) float64 {
	return bp(e, r...) * cp(e, r...)
}
