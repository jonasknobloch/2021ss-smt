package scm

type scm struct {
	VE    []string                      // English vocabulary
	VF    []string                      // French vocabulary
	final string                        // #
	b     map[string]map[string]float64 // bigram model b(e'|e)
	t     map[string]map[string]float64 // dictionary t(f|e)
	l     func(l, m int) float64        // length model ε(m|l)
	lMax  int                           // max scoring length
	YMax  int                           // max hypotheses
}

func NewSCM(VE, VF []string, final string, b, t map[string]map[string]float64, l func(l, m int) float64, lMax, YMax int) *scm {
	return &scm{
		VE:    VE,
		VF:    VF,
		final: final,
		b:     b,
		t:     t,
		l:     l,
		lMax:  lMax,
		YMax:  YMax,
	}
}
