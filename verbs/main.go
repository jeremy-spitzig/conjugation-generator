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
		Infinitivo:                                 infinitiveVerbTense(stem, "ar", "ares", "armos", "ardes", "arem"),
		Imperativo:                                 imperativeVerbTense(stem, "a", "e", "emos", "ai", "em"),
		ImperativoNegativo:                         negativeImperativeVerbTense(stem, "a", "e", "emos", "ai", "em"),
		Gerúndio:                                   stem + "ando",
		Particípio:                                 stem + "ado",
	}
}

func verbTense(stem, fpsSuffix, spsSuffix, tpsSuffix, fppSuffix, sppSuffix, tppSuffix string) VerbTense {
	return VerbTense{stem + fpsSuffix, stem + spsSuffix, stem + tpsSuffix, stem + fppSuffix, stem + sppSuffix, stem + tppSuffix}
}

func compoundVerbTense(stem, fpsPrefix, spsPrefix, tpsPrefix, fppPrefix, sppPrefix, tppPrefix, fpsSuffix, spsSuffix, tpsSuffix, fppSuffix, sppSuffix, tppSuffix string) VerbTense {
	return VerbTense{fpsPrefix + stem + fpsSuffix, spsPrefix + stem + spsSuffix, tpsPrefix + stem + tpsSuffix, fppPrefix + stem + fppSuffix, sppPrefix + stem + sppSuffix, tppPrefix + stem + tppSuffix}
}

func compoundVerbTenseSimple(stem, fpsPrefix, spsPrefix, tpsPrefix, fppPrefix, sppPrefix, tppPrefix, suffix string) VerbTense {
	return compoundVerbTense(stem, fpsPrefix, spsPrefix, tpsPrefix, fppPrefix, sppPrefix, tppPrefix, suffix, suffix, suffix, suffix, suffix, suffix)
}

func infinitiveVerbTense(stem, imp, spsSuffix, fppSuffix, sppSuffix, tppSuffix string) InfinitiveVerbTense {
	return InfinitiveVerbTense{stem + imp, stem + spsSuffix, stem + fppSuffix, stem + sppSuffix, stem + tppSuffix}
}

func imperativeVerbTense(stem, spsSuffix, tpsSuffix, fppSuffix, sppSuffix, tppSuffix string) ImperativeVerbTense {
	return ImperativeVerbTense{stem + spsSuffix, stem + tpsSuffix, stem + fppSuffix, stem + sppSuffix, stem + tppSuffix}
}

func negativeImperativeVerbTense(stem, spsSuffix, tpsSuffix, fppSuffix, sppSuffix, tppSuffix string) ImperativeVerbTense {
	return ImperativeVerbTense{"não " + stem + spsSuffix, "não " + stem + tpsSuffix, "não " + stem + fppSuffix, "não " + stem + sppSuffix, "não " + stem + tppSuffix}
}
