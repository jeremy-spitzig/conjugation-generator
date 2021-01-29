package languagepack

import (
	"github.com/go-git/go-git/v5"
	"io"
	"io/ioutil"
	"os"
)

type gitLanguagePack struct {
	fs LanguagePack
	td string
}

func NewGit(url string) (LanguagePack, error) {
	d, err := ioutil.TempDir("", "git-clone")
	if err != nil {
		return nil, err
	}
	_, err = git.PlainClone(d, false, &git.CloneOptions{
		URL: url,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}
	fs := NewFileSystem(d)
	return &gitLanguagePack{fs, d}, nil
}

func (g *gitLanguagePack) Verbs() (io.Reader, error) {
	return g.fs.Verbs()
}

func (g *gitLanguagePack) Models() ([]ModelInput, error) {
	return g.fs.Models()
}

func (g *gitLanguagePack) Templates() ([]TemplateInput, error) {
	return g.fs.Templates()
}

func (g *gitLanguagePack) Close() {
	g.fs.Close()
	os.RemoveAll(g.td)
}
