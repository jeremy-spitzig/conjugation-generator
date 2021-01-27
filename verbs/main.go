package verbs

import (
	"encoding/json"
	"os"
)

// Verb represents information about a verb in english and portuguese
type Verb struct {
	PortugueseInfinitive        string `json:"portugueseInfinitive"`
	Irregular                   bool   `json:"irregular"`
	Infinitive                  string `json:"infinitive"`
	Present                     string `json:"present"`
	ThirdPersonPresent          string `json:"thirdPersonPresent"`
	Gerund                      string `json:"gerund"`
	Past                        string `json:"past"`
	ExampleComplement           string `json:"exampleComplement"`
	PortugueseExampleComplement string `json:"portugueseExampleComplement"`
}

// ReadVerbs reads in verb from the supplied filename
func ReadVerbs(fn string) ([]Verb, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	var verbs []Verb

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&verbs)
	if err != nil {
		return nil, err
	}
	return verbs, nil
}
