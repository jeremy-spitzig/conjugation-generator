package sentences

import (
	"github.com/jeremy-spitzig/conjugation-generator/languagepack"
	"github.com/jeremy-spitzig/conjugation-generator/verbs"
	"io"
	"strings"
	"text/template"
)

type Templates struct {
	template *template.Template
}

type templateData struct{
	Verb verbs.Verb
	Tense interface{}
}

func LoadTemplates(pack languagepack.LanguagePack) (*Templates, error) {
	t := template.New("")
	tis, err := pack.Templates()
	if err != nil {
		return nil, err
	}
	for _, ti := range tis {
		buf := new(strings.Builder)
		io.Copy(buf, ti.Reader)
		t.New(ti.Name).Parse(buf.String())
	}
	return &Templates{t}, nil
}

func (t *Templates) Execute(v verbs.Verb, c *verbs.Conjugation, w io.Writer) error {
	for _, t := range t.template.Templates() {
		tn := t.Name()
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