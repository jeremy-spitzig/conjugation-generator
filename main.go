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
	ofe := os.Getenv("SENTENCE_OUTPUT_EXT")
	if ifn == "" {
		ifn = "./input.json"
	}
	if ofn == "" {
		ofn = "./output"
	}
	if ofe == "" {
		ofe = ".csv"
	}
	verbs, err := verbs.ReadVerbs(ifn)
	if err != nil {
		log.Panicln(err)
	}

	index := 1

	of, err := getFile(ofn, ofe, index)
	if err != nil {
		log.Panicln(err)
	}

	lines := 0

	for _, v := range verbs {
		gs, err := sentences.GenerateSentences(v)
		if err != nil {
			log.Panicln(err)
		}
		for _, s := range gs {
			if lines >= 2000 {
				lines = 0
				index++
				of, err = getFile(ofn, ofe, index)
				if err != nil {
					log.Panicln(err)
				}
			}
			fmt.Fprint(of, s.English, ";", s.Português, "\n")
			lines++
		}
	}
	of.Close()
}

func getFile(ofn, ofe string, index int) (*os.File, error) {
	of, err := os.OpenFile(fmt.Sprintf("%s-%d%s", ofn, index, ofe), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return of, nil
}
