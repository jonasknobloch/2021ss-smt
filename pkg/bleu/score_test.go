package bleu

import (
	"strings"
	"testing"
)

func TestErl(t *testing.T) {
	e := strings.Split("a b c d e", " ")

	r1 := strings.Split("a b c d e f g", " ")
	r2 := strings.Split("a b c d", " ")
	r3 := strings.Split("a b c d e f g h", " ")
	r4 := strings.Split("a b c d e f", " ")

	if erl := erl(e, r1, r2, r3, r4); erl != 4 {
		t.Errorf("unexpected effective references length: got %d want %d", erl, 4)
	}
}
