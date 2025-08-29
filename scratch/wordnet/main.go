package main

import (
	"core/sys/wordnet"
	"fmt"
)

func main() {
	a := wordnet.Lemma["aardvark"]
	if len(a.Adjective.Senses) == 0 {
		fmt.Println()
	}
}
