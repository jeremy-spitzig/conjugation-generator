package verbs

import (
	"encoding/json"
	"fmt"
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
	Type                        string `json:"type"`
}

type VerbTense struct {
	FirstPersonSingular  string
	SecondPersonSingular string
	ThirdPersonSingular  string
	FirstPersonPlural    string
	SecondPersonPlural   string
	ThirdPersonPlural    string
}

type Conjugation struct {
	IndicativoPresente                         VerbTense
	IndicativoPretéritoPerfeito                VerbTense
	IndicativoPretéritoImperfeito              VerbTense
	IndicativoPretéritoMaisQuePerfeito         VerbTense
	IndicativoPretéritoPerfeitoComposto        VerbTense
	IndicativoPretéritoMaisQuePerfeitoComposto VerbTense
	IndicativoPretéritoMaisQuePerfeitoAnterior VerbTense
	IndicativoFuturoDoPresenteSimples          VerbTense
	IndicativoFuturoCompostoComIr              VerbTense
	IndicativoFuturoDoPresenteComposto         VerbTense
	SubjuntivoPresente                         VerbTense
	SubjuntivoPretéritoPerfeito                VerbTense
	SubjuntivoPretéritoImperfeito              VerbTense
	SubjuntivoPretéritoMaisQuePerfeitoComposto VerbTense
	SubjuntivoFuturo                           VerbTense
	SubjuntivoFuturoComposto                   VerbTense
	CondicionalFuturoDoPretéritoSimples        VerbTense
	FuturoDoPretéritoComposto                  VerbTense
	// Infinitivo                                 VerbTense
	// Imperativo                                 VerbTense
	// ImperativoNegativo                         VerbTense
	Gerúndio   string
	Particípio string
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

func (v Verb) Conjugate() (*Conjugation, error) {
	// Stem = infinitive - last 2 characters (ar/er/ir)
	// TODO: Maybe better to pull stem out to a separate attribute for irregulars,
	// They're going to have specialized conjugation anyway, so might not be worth it
	stem := v.Infinitive[:len(v.Infinitive)-2]
	switch v.Type {
	case "ar":
		return arConjugation(stem), nil
	}
	return nil, fmt.Errorf("Unrecognized verb type %s", v.Type)
}

func arConjugation(stem string) *Conjugation {
	return &Conjugation{
		IndicativoPresente:                         verbTense(stem, "o", "as", "a", "amos", "ais", "am"),
		IndicativoPretéritoPerfeito:                verbTense(stem, "ei", "aste", "ou", "ámos", "astes", "aram"),
		IndicativoPretéritoImperfeito:              verbTense(stem, "ava", "avas", "ava", "ávamos", "áveis", "avam"),
		IndicativoPretéritoMaisQuePerfeito:         verbTense(stem, "ara", "aras", "ara", "áramos", "áreis", "aram"),
		IndicativoPretéritoPerfeitoComposto:        compoundVerbTenseSimple(stem, "tenho ", "tens ", "tem ", "temos ", "tendes ", "têm ", "ado"),
		IndicativoPretéritoMaisQuePerfeitoComposto: compoundVerbTenseSimple(stem, "tinha ", "tinhas ", "tinha ", "tínhamos ", "tínheis ", "tinham ", "ado"),
		IndicativoPretéritoMaisQuePerfeitoAnterior: compoundVerbTenseSimple(stem, "tivera ", "tiveras ", "tivera ", "tivévamos ", "tivéreis ", "tiveram ", "ado"),
		IndicativoFuturoDoPresenteSimples:          verbTense(stem, "arei", "arás", "ará", "aremos", "areis", "arão"),
		IndicativoFuturoCompostoComIr:              compoundVerbTenseSimple(stem, "vou ", "vais ", "vai ", "vamos ", "ides ", "vão ", "ar"),
		IndicativoFuturoDoPresenteComposto:         compoundVerbTenseSimple(stem, "tenha ", "tenhas ", "tenha ", "tenhamos ", "tenhais ", "tenham ", "ado"),
		SubjuntivoPresente:                         verbTense(stem, "e", "es", "e", "emos", "eis", "em"),
		SubjuntivoPretéritoPerfeito:                compoundVerbTenseSimple(stem, "tenha ", "tenhas ", "tenha ", "tenhamos ", "tenhais ", "tenham ", "ado"),
		SubjuntivoPretéritoImperfeito:              verbTense(stem, "asse", "asses", "asse", "ássemos", "ásseis", "assem"),
		SubjuntivoPretéritoMaisQuePerfeitoComposto: compoundVerbTenseSimple(stem, "tivesse ", "tivesses ", "tivesse ", "tivéssemos ", "tivésseis ", "tivessem ", "ado"),
		SubjuntivoFuturo:                           verbTense(stem, "ar", "ares", "ar", "armos", "ardes", "arem"),
		SubjuntivoFuturoComposto:                   compoundVerbTenseSimple(stem, "tiver ", "tiveres ", "tiver ", "tivermos ", "tiverdes ", "tiverem ", "ado"),
		CondicionalFuturoDoPretéritoSimples:        verbTense(stem, "aria", "arias", "aria", "aríamos", "aríeis", "ariam"),
		FuturoDoPretéritoComposto:                  compoundVerbTenseSimple(stem, "teria ", "terias ", "teria ", "teríamos ", "teríeis ", "teriam ", "ado"),
		// TODO: Restructure these into new struct(s?)
		// Infinitivo:                                 verbTense(stem, "ar", "ares", "ar", "aremos", "areis", "arão"),
		// Imperativo:                                 verbTense(stem, "arei", "arás", "ará", "aremos", "areis", "arão"),
		// ImperativoNegativo:                         compoundVerbTenseSimple(stem, "arei", "arás", "ará", "aremos", "areis", "arão"),
		Gerúndio:   stem + "ando",
		Particípio: stem + "ado",
	}
}

func verbTense(stem string, fpsSuffix string, spsSuffix, tpsSuffix string, fppSuffix string, sppSuffix string, tppSuffix string) VerbTense {
	return VerbTense{stem + fpsSuffix, stem + spsSuffix, stem + tpsSuffix, stem + fppSuffix, stem + sppSuffix, stem + tppSuffix}
}

func compoundVerbTense(stem string, fpsPrefix string, spsPrefix, tpsPrefix string, fppPrefix string, sppPrefix string, tppPrefix string, fpsSuffix string, spsSuffix, tpsSuffix string, fppSuffix string, sppSuffix string, tppSuffix string) VerbTense {
	return VerbTense{fpsPrefix + stem + fpsSuffix, spsPrefix + stem + spsSuffix, tpsPrefix + stem + tpsSuffix, fppPrefix + stem + fppSuffix, sppPrefix + stem + sppSuffix, tppPrefix + stem + tppSuffix}
}

func compoundVerbTenseSimple(stem string, fpsPrefix string, spsPrefix, tpsPrefix string, fppPrefix string, sppPrefix string, tppPrefix string, suffix string) VerbTense {
	return VerbTense{fpsPrefix + stem + suffix, spsPrefix + stem + suffix, tpsPrefix + stem + suffix, fppPrefix + stem + suffix, sppPrefix + stem + suffix, tppPrefix + stem + suffix}
}
