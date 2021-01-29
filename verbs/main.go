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
	Model                       string `json:"model"`
}

type VerbTense struct {
	FirstPersonSingular  string
	SecondPersonSingular string
	ThirdPersonSingular  string
	FirstPersonPlural    string
	SecondPersonPlural   string
	ThirdPersonPlural    string
}

type InfinitiveVerbTense struct {
	Impersonal           string
	SecondPersonSingular string
	FirstPersonPlural    string
	SecondPersonPlural   string
	ThirdPersonPlural    string
}

type ImperativeVerbTense struct {
	SecondPersonSingular string
	ThirdPersonSingular  string
	FirstPersonPlural    string
	SecondPersonPlural   string
	ThirdPersonPlural    string
}

type Conjugation struct {
	IndicativoPresente                         VerbTense //Finished
	IndicativoPretéritoPerfeito                VerbTense //Finished
	IndicativoPretéritoImperfeito              VerbTense //Finished
	IndicativoPretéritoMaisQuePerfeito         VerbTense
	IndicativoPretéritoPerfeitoComposto        VerbTense
	IndicativoPretéritoMaisQuePerfeitoComposto VerbTense
	IndicativoPretéritoMaisQuePerfeitoAnterior VerbTense
	IndicativoFuturoDoPresenteSimples          VerbTense //Finished
	IndicativoFuturoCompostoComIr              VerbTense //Finished
	IndicativoFuturoDoPresenteComposto         VerbTense
	SubjuntivoPresente                         VerbTense //Finished
	SubjuntivoPretéritoPerfeito                VerbTense
	SubjuntivoPretéritoImperfeito              VerbTense
	SubjuntivoPretéritoMaisQuePerfeitoComposto VerbTense
	SubjuntivoFuturo                           VerbTense //Finished
	SubjuntivoFuturoComposto                   VerbTense
	CondicionalFuturoDoPretéritoSimples        VerbTense
	FuturoDoPretéritoComposto                  VerbTense
	Infinitivo                                 InfinitiveVerbTense
	Imperativo                                 ImperativeVerbTense
	ImperativoNegativo                         ImperativeVerbTense
	Gerúndio                                   string
	Particípio                                 string
}

// ReadVerbs reads in verb from the supplied filename
func ReadVerbs(fn string) ([]Verb, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var verbs []Verb

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&verbs)
	if err != nil {
		return nil, err
	}
	return verbs, nil
}
