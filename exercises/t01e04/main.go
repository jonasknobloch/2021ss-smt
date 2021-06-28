package main

import (
	"fmt"
	"github.com/jonasknobloch/2021ss-smt/pkg/bleu"
	"strings"
)

func main() {
	er := strings.Split("As I have said , this activity causes enormous damage .", " ")
	ed := strings.Split("As I said before , this activity causes enormous damage .", " ")
	eg := strings.Split("As I said , this activity causes enormous damage .", " ")

	fmt.Printf("BLEU(e_deepl, e_r) = %f\n", bleu.Score(ed, er))
	fmt.Printf("BLEU(e_google, e_r) = %f\n", bleu.Score(eg, er))
	fmt.Printf("BLEU(e_google, e_r, e_deepl) = %f\n", bleu.Score(eg, er, ed))
}
