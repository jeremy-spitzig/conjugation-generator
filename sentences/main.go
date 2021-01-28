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

	return indicativePresentTenseSentences(v, c), nil

}

func indicativePresentTenseSentences(v verbs.Verb, c *verbs.Conjugation) []Sentence {
	sentences := make([]Sentence, len(subjects))
	var es string
	var ps string
	for i, subject := range subjects {
		if subject.englishPerson == third && !subject.plural {
			es = subject.english + " " + v.ThirdPersonPresent + " " + v.ExampleComplement
		} else {
			es = subject.english + " " + v.Present + " " + v.ExampleComplement
		}

		ps = portuguêsSentence(subject, v, c.IndicativoPresente)
		sentences[i] = Sentence{es, ps}
	}
	return sentences
}

func portuguêsSentence(s subject, v verbs.Verb, t verbs.VerbTense) string {
	switch {
	case s.portuguesePerson == first && !s.plural:
		return s.português + " " + t.FirstPersonSingular + " " + v.PortugueseExampleComplement
	case s.portuguesePerson == second && !s.plural:
		return s.português + " " + t.SecondPersonSingular + " " + v.PortugueseExampleComplement
	case s.portuguesePerson == third && !s.plural:
		return s.português + " " + t.ThirdPersonSingular + " " + v.PortugueseExampleComplement
	case s.portuguesePerson == first && s.plural:
		return s.português + " " + t.FirstPersonPlural + " " + v.PortugueseExampleComplement
	case s.portuguesePerson == second && s.plural:
		return s.português + " " + t.SecondPersonPlural + " " + v.PortugueseExampleComplement
	case s.portuguesePerson == third && s.plural:
		return s.português + " " + t.ThirdPersonPlural + " " + v.PortugueseExampleComplement
	}
	return ""
}
