package verbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type modelTense struct {
	TenseType string   `json:"type"`
	Forms     []string `json:"forms"`
}

type model struct {
	Capture string                `json:"capture"`
	Tenses  map[string]modelTense `json:"tenses"`
}

type Models struct {
	models map[string]model
}

type modelFile struct {
	fileName  string
	modelName string
}

// Models methods

func (ms *Models) Conjugate(v Verb) (*Conjugation, error) {
	if m, ok := ms.models[v.Type]; ok {
		capture, err := regexp.Compile(m.Capture)
		if err != nil {
			return nil, err
		}
		dcvt := func(v Verb, m modelTense, capture *regexp.Regexp) *VerbTense {
			if err != nil {
				return nil
			}
			var vt *VerbTense
			vt, err = conjugateVerbTense(v, m, capture)
			return vt
		}
		dcinfvt := func(v Verb, m modelTense, capture *regexp.Regexp) *InfinitiveVerbTense {
			if err != nil {
				return nil
			}
			var vt *InfinitiveVerbTense
			vt, err = conjugateInfinitiveVerbTense(v, m, capture)
			return vt
		}
		dcimpvt := func(v Verb, m modelTense, capture *regexp.Regexp) *ImperativeVerbTense {
			if err != nil {
				return nil
			}
			var vt *ImperativeVerbTense
			vt, err = conjugateImperativeVerbTense(v, m, capture)
			return vt
		}
		return &Conjugation{
			IndicativoPresente:                         *dcvt(v, m.Tenses["indicativoPresente"], capture),
			IndicativoPretéritoPerfeito:                *dcvt(v, m.Tenses["indicativoPretéritoPerfeito"], capture),
			IndicativoPretéritoImperfeito:              *dcvt(v, m.Tenses["indicativoPretéritoImperfeito"], capture),
			IndicativoPretéritoMaisQuePerfeito:         *dcvt(v, m.Tenses["indicativoPretéritoMaisQuePerfeito"], capture),
			IndicativoPretéritoPerfeitoComposto:        *dcvt(v, m.Tenses["indicativoPretéritoPerfeitoComposto"], capture),
			IndicativoPretéritoMaisQuePerfeitoComposto: *dcvt(v, m.Tenses["indicativoPretéritoMaisQuePerfeitoComposto"], capture),
			IndicativoPretéritoMaisQuePerfeitoAnterior: *dcvt(v, m.Tenses["indicativoPretéritoMaisQuePerfeitoAnterior"], capture),
			IndicativoFuturoDoPresenteSimples:          *dcvt(v, m.Tenses["indicativoFuturoDoPresenteSimples"], capture),
			IndicativoFuturoCompostoComIr:              *dcvt(v, m.Tenses["indicativoFuturoCompostoComIr"], capture),
			IndicativoFuturoDoPresenteComposto:         *dcvt(v, m.Tenses["indicativoFuturoDoPresenteComposto"], capture),
			SubjuntivoPresente:                         *dcvt(v, m.Tenses["subjuntivoPresente"], capture),
			SubjuntivoPretéritoPerfeito:                *dcvt(v, m.Tenses["subjuntivoPretéritoPerfeito"], capture),
			SubjuntivoPretéritoImperfeito:              *dcvt(v, m.Tenses["subjuntivoPretéritoImperfeito"], capture),
			SubjuntivoPretéritoMaisQuePerfeitoComposto: *dcvt(v, m.Tenses["subjuntivoPretéritoMaisQuePerfeitoComposto"], capture),
			SubjuntivoFuturo:                           *dcvt(v, m.Tenses["subjuntivoFuturo"], capture),
			SubjuntivoFuturoComposto:                   *dcvt(v, m.Tenses["subjuntivoFuturoComposto"], capture),
			CondicionalFuturoDoPretéritoSimples:        *dcvt(v, m.Tenses["condicionalFuturoDoPretéritoSimples"], capture),
			FuturoDoPretéritoComposto:                  *dcvt(v, m.Tenses["futuroDoPretéritoComposto"], capture),
			Infinitivo:                                 *dcinfvt(v, m.Tenses["infinitivo"], capture),
			Imperativo:                                 *dcimpvt(v, m.Tenses["imperativo"], capture),
			ImperativoNegativo:                         *dcimpvt(v, m.Tenses["imperativoNegativo"], capture),
			Gerúndio:                                   capture.ReplaceAllString(v.PortugueseInfinitive, m.Tenses["gerúndio"].Forms[0]),
			Particípio:                                 capture.ReplaceAllString(v.PortugueseInfinitive, m.Tenses["particípio"].Forms[0]),
		}, nil
	} else {
		return nil, fmt.Errorf("Unrecognized verb type %s", v.Type)
	}
}

func conjugateVerbTense(v Verb, m modelTense, capture *regexp.Regexp) (*VerbTense, error) {
	if len(m.Forms) != 6 {
		return nil, fmt.Errorf("Incorrect model structure")
	}
	return &VerbTense{
		FirstPersonSingular:  capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[0]),
		SecondPersonSingular: capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[1]),
		ThirdPersonSingular:  capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[2]),
		FirstPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[3]),
		SecondPersonPlural:   capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[4]),
		ThirdPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[5]),
	}, nil
}

func conjugateInfinitiveVerbTense(v Verb, m modelTense, capture *regexp.Regexp) (*InfinitiveVerbTense, error) {
	if len(m.Forms) != 5 {
		return nil, fmt.Errorf("Incorrect model structure")
	}
	return &InfinitiveVerbTense{
		Impersonal:           capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[0]),
		SecondPersonSingular: capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[1]),
		FirstPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[2]),
		SecondPersonPlural:   capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[3]),
		ThirdPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[4]),
	}, nil
}

func conjugateImperativeVerbTense(v Verb, m modelTense, capture *regexp.Regexp) (*ImperativeVerbTense, error) {
	if len(m.Forms) != 5 {
		return nil, fmt.Errorf("Incorrect model structure")
	}
	return &ImperativeVerbTense{
		SecondPersonSingular: capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[0]),
		ThirdPersonSingular:  capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[1]),
		FirstPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[2]),
		SecondPersonPlural:   capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[3]),
		ThirdPersonPlural:    capture.ReplaceAllString(v.PortugueseInfinitive, m.Forms[4]),
	}, nil
}

// Loading Models

func LoadModels(md string) (*Models, error) {
	ms := map[string]model{}
	mfs, err := listModelFiles(md)
	if err != nil {
		return nil, err
	}

	for _, mf := range mfs {
		m, err := readFile(mf.fileName)
		if err != nil {
			return nil, err
		}
		ms[mf.modelName] = *m
	}

	return &Models{ms}, nil
}

func readFile(fn string) (*model, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var m model
	d := json.NewDecoder(f)
	err = d.Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func listModelFiles(md string) ([]modelFile, error) {
	var files []modelFile
	fileInfo, err := ioutil.ReadDir(md)
	if err != nil {
		return files, err
	}

	const modelSuffix = ".model.json"
	for _, file := range fileInfo {
		fn := file.Name()
		if strings.HasSuffix(fn, modelSuffix) {
			files = append(files, modelFile{
				modelName: strings.TrimSuffix(fn, modelSuffix),
				fileName:  filepath.Join(md, fn),
			})
		}
	}
	return files, nil
}
