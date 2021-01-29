package verbs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
