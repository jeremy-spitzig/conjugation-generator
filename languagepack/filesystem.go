package languagepack

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type fileSystem struct {
	rootDirectoryName 	string
	openFiles 			[]*os.File
}

func (fs *fileSystem) addOpenFile(f *os.File) {
	fs.openFiles = append(fs.openFiles, f)
}

func NewFileSystem(rootDirectoryName string) LanguagePack {
	return &fileSystem{rootDirectoryName, []*os.File{}}
}

func (fs *fileSystem) Verbs() (io.Reader, error) {
	vfn := filepath.Join(fs.rootDirectoryName, "verbs.json")
	f, err := os.Open(vfn)
	if err != nil {
		return nil, err
	}
	fs.addOpenFile(f)
	return f, nil
}

func (fs *fileSystem) Models() ([]ModelInput, error) {
	md := filepath.Join(fs.rootDirectoryName, "models")

	var mis []ModelInput
	fileInfo, err := ioutil.ReadDir(md)
	if err != nil {
		return nil, err
	}

	const modelSuffix = ".model.json"
	for _, file := range fileInfo {
		fn := file.Name()
		if strings.HasSuffix(fn, modelSuffix) {
			f, err := os.Open(filepath.Join(md, fn))
			if err != nil {
				return nil, err
			}
			fs.addOpenFile(f)
			mis = append(mis, ModelInput{
				Name: strings.TrimSuffix(fn, modelSuffix),
				Reader: f,
			})
		}
	}
	return mis, nil
}

func (fs *fileSystem) Templates() ([]TemplateInput, error) {
	td := filepath.Join(fs.rootDirectoryName, "templates")

	var tis []TemplateInput
	fileInfo, err := ioutil.ReadDir(td)
	if err != nil {
		return nil, err
	}

	const templateSuffix = ".tmpl"
	for _, file := range fileInfo {
		fn := file.Name()
		if strings.HasSuffix(fn, templateSuffix) {
			f, err := os.Open(filepath.Join(td, fn))
			if err != nil {
				return nil, err
			}
			fs.addOpenFile(f)
			tis = append(tis, TemplateInput{
				Name: strings.TrimSuffix(fn, templateSuffix),
				Reader: f,
			})
		}
	}
	return tis, nil
}

func (fs *fileSystem) Close() {
	for _, f := range fs.openFiles {
		f.Close()
	}
}
