package sentences

import (
	"github.com/jeremy-spitzig/conjugation-generator/verbs"
	"io"
	"path/filepath"
	"text/template"
)

type Templates struct {
	template *template.Template
}

type templateData struct{
	Verb verbs.Verb
	Tense interface{}
}

func LoadTemplates(td string) (*Templates, error) {
	glob := filepath.Join(td, "*.tmpl")
	t, err := template.New("").ParseGlob(glob)
	if err != nil {
		return nil, err
	}
	return &Templates{t}, nil
}

func (t *Templates) Execute(v verbs.Verb, c *verbs.Conjugation, w io.Writer) error {
	for _, t := range t.template.Templates() {
		tn := t.Name()[:len(t.Name()) - 5]
		vt, err := c.Tense(tn)
		if err != nil {
			return err
		}
		var td templateData

		switch vt.(type){
		case *verbs.VerbTense:
			td = templateData{v, *(vt.(*verbs.VerbTense))}
		case *verbs.InfinitiveVerbTense:
			td = templateData{v, *(vt.(*verbs.InfinitiveVerbTense))}
		case *verbs.ImperativeVerbTense:
			td = templateData{v, *(vt.(*verbs.ImperativeVerbTense))}
		case string:
			td = templateData{v, vt.(string)}
		}
		t.Execute(w, td)
	}
	return nil
}