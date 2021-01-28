package sentences

import (
	"log"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

type Sentence struct {
	English   string
	Português string
}

type person int

const (
	first person = iota
	second
	third
)

type subject struct {
	english          string
	português        string
	englishPerson    person
	portuguesePerson person
	plural           bool
}

var subjects []subject = []subject{
	{"I", "Eu", first, first, false},
	{"You", "Você", second, third, false},
	{"He", "Ele", third, third, false},
	{"She", "Ela", third, third, false},
	{"We", "Nós", first, first, true},
	{"Y'all", "Vocês", second, third, true},
	{"The guys", "Eles", third, third, true},
	{"The gals", "Elas", third, third, true},
}

func GenerateSentences(v verbs.Verb) ([]Sentence, error) {

	c, err := v.Conjugate()

	if err != nil {
		log.Printf("Failed to conjugate %s\n", v.PortugueseInfinitive)
		return nil, err
	}

	return append(append(
		indicativoPresenteSentences(v, c),
		indicativoPretéritoPerfeitoSentences(v, c)...),
		indicativoPretéritoImperfeitoSentences(v, c)...,
	), nil

}

func indicativoPresenteSentences(v verbs.Verb, c *verbs.Conjugation) []Sentence {
	et := verbs.VerbTense{v.Present, v.Present, v.ThirdPersonPresent, v.Present, v.Present, v.Present}
	return sentences(v, et, c.IndicativoPresente, "", "", "", "")
}

func indicativoPretéritoPerfeitoSentences(v verbs.Verb, c *verbs.Conjugation) []Sentence {
	et := verbs.VerbTense{v.Past, v.Past, v.Past, v.Past, v.Past, v.Past}
	return sentences(v, et, c.IndicativoPretéritoPerfeito, "", "", "", "")
}

func indicativoPretéritoImperfeitoSentences(v verbs.Verb, c *verbs.Conjugation) []Sentence {
	vf := "used to " + v.Present
	et := verbs.VerbTense{vf, vf, vf, vf, vf, vf}
	return sentences(v, et, c.IndicativoPretéritoImperfeito, "As a child, ", "", "Como uma criança, ", "")
}

func sentences(v verbs.Verb, et verbs.VerbTense, pt verbs.VerbTense, epfx string, esfx string, ppfx string, psfx string) []Sentence {
	sentences := make([]Sentence, len(subjects))
	var es string
	var ps string
	for i, s := range subjects {
		es = sentence(s.english, s.englishPerson, s.plural, v.ExampleComplement, et, epfx, esfx)
		ps = sentence(s.português, s.portuguesePerson, s.plural, v.PortugueseExampleComplement, pt, ppfx, psfx)
		sentences[i] = Sentence{es, ps}
	}
	return sentences
}

func sentence(s string, p person, plural bool, c string, t verbs.VerbTense, pfx string, sfx string) string {
	switch {
	case p == first && !plural:
		return pfx + s + " " + t.FirstPersonSingular + " " + c + sfx
	case p == second && !plural:
		return pfx + s + " " + t.SecondPersonSingular + " " + c + sfx
	case p == third && !plural:
		return pfx + s + " " + t.ThirdPersonSingular + " " + c + sfx
	case p == first && plural:
		return pfx + s + " " + t.FirstPersonPlural + " " + c + sfx
	case p == second && plural:
		return pfx + s + " " + t.SecondPersonPlural + " " + c + sfx
	case p == third && plural:
		return pfx + s + " " + t.ThirdPersonPlural + " " + c + sfx
	}
	return ""

}
