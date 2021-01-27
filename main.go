package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

func main() {
	fn := os.Getenv("VERB_INPUT_FILE")
	if fn == "" {
		fn = "./input.json"
	}
	verbs, err := verbs.ReadVerbs(fn)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(verbs)
}
