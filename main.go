package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremy-spitzig/conjugation-generator/sentences"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

func main() {
	ifn := os.Getenv("VERB_INPUT_FILE")
	ofn := os.Getenv("SENTENCE_OUTPUT_FILE")
	if ifn == "" {
		ifn = "./input.json"
	}
	if ofn == "" {
		ofn = "./output.csv"
	}
	verbs, err := verbs.ReadVerbs(ifn)
	if err != nil {
		log.Panicln(err)
	}

	of, err := os.OpenFile(ofn, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer of.Close()

	for _, v := range verbs {
		gs, err := sentences.GenerateSentences(v)
		if err != nil {
			log.Panicln(err)
		}
		for _, s := range gs {
			fmt.Fprint(of, s.English, ";", s.Português, "\n")
		}
	}
}
