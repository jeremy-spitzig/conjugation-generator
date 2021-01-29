package languagepack

import "io"

type ModelInput struct {
	Name string
	Reader io.Reader
}

type TemplateInput struct {
	Name string
	Reader io.Reader
}

type LanguagePack interface {
	Verbs() (io.Reader, error)
	Models() ([]ModelInput, error)
	Templates() ([]TemplateInput, error)
	Close()
}
