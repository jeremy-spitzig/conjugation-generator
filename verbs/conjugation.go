package verbs

import (
	"fmt"
	"regexp"
)

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
	conjugation map[string]interface{}
}

func (c *Conjugation) Tense(t string) (interface{}, error) {
	if t, ok := c.conjugation[t]; ok {
		return t, nil
	} else {
		return nil, fmt.Errorf("Verb tense not found: %s", t)
	}
}

func (ms *Models) Conjugate(v Verb) (*Conjugation, error) {
	if m, ok := ms.models[v.Model]; ok {
		conjugation := map[string]interface{}{}
		capture, err := regexp.Compile(m.Capture)
		if err != nil {
			return nil, err
		}
		exec := func(f func()(interface{}, error)) interface{} {
			if err != nil {
				return nil
			}
			var r interface{}
			r, err = f()
			return r
		}
		for key, tense := range m.Tenses {
			switch tense.TenseType {
			case "regular":
				conjugation[key] = exec(func() (interface{}, error) {return conjugateVerbTense(v, tense, capture)})
			case "infinitive":
				conjugation[key] = exec(func() (interface{}, error) {return conjugateInfinitiveVerbTense(v, tense, capture)})
			case "imperative":
				conjugation[key] = exec(func() (interface{}, error) {return conjugateImperativeVerbTense(v, tense, capture)})
			case "gerund":
				conjugation[key] = capture.ReplaceAllString(v.PortugueseInfinitive, tense.Forms[0])
			case "participle":
				conjugation[key] = capture.ReplaceAllString(v.PortugueseInfinitive, tense.Forms[0])
			}
			if err != nil {
				return nil, err
			}
		}
		return &Conjugation{ conjugation }, nil
	} else {
		return nil, fmt.Errorf("Unrecognized verb type %s", v.Model)
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
