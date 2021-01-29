package verbs

import (
	"encoding/json"
	"github.com/jeremy-spitzig/conjugation-generator/languagepack"
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
	Model                       string `json:"model"`
}

// ReadVerbs reads in verb from the supplied filename
func ReadVerbs(pack languagepack.LanguagePack) ([]Verb, error) {
	r, err := pack.Verbs()
	if err != nil {
		return nil, err
	}

	var verbs []Verb
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&verbs)
	if err != nil {
		return nil, err
	}

	return verbs, nil
}
