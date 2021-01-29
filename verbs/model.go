package verbs

import (
	"encoding/json"
	"github.com/jeremy-spitzig/conjugation-generator/languagepack"
	"io"
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

// Loading Models

func LoadModels(pack languagepack.LanguagePack) (*Models, error) {
	ms := map[string]model{}
	mis, err := pack.Models()
	if err != nil {
		return nil, err
	}

	for _, mi := range mis {
		m, err := read(mi.Reader)
		if err != nil {
			return nil, err
		}
		ms[mi.Name] = *m
	}

	return &Models{ms}, nil
}

func read(r io.Reader) (*model, error) {
	var m model
	d := json.NewDecoder(r)
	err := d.Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
