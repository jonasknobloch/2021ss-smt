package ibm

type model struct {
	VE []string                      // English vocabulary
	VF []string                      // French vocabulary
	t  map[string]map[string]float64 // dictionary t(f|e)
	l  func(m, l int) float64        // length model ε(l|m)
}

func NewModel(VE, VF []string, t map[string]map[string]float64, l func(m, l int) float64) *model {
	return &model{
		VE: VE,
		VF: VF,
		t:  t,
		l:  l,
	}
}
