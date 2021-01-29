package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremy-spitzig/conjugation-generator/sentences"
	"github.com/urfave/cli/v2"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "input",
				Value: "./default-input.json",
				Usage: "Load input from `FILE`",
			},
			&cli.StringFlag{
				Name:  "outputFileBase",
				Value: "output",
				Usage: "The base of the filenames for output files",
			},
			&cli.StringFlag{
				Name:  "outputFileExtension",
				Value: ".csv",
				Usage: "The extension of the filenames for output files",
			},
			&cli.StringFlag{
				Name:  "modelsDir",
				Value: "./default-verb-models",
				Usage: "Directory containing verb model files",
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {

	md := c.String("modelsDir")
	m, err := verbs.LoadModels(md)

	if err != nil {
		return err
	}

	ifn := c.String("input")
	ofn := c.String("outputFileBase")
	ofe := c.String("outputFileExtension")

	verbs, err := verbs.ReadVerbs(ifn)
	if err != nil {
		return err
	}

	index := 1

	of, err := getFile(ofn, ofe, index)
	if err != nil {
		return err
	}

	lines := 0

	for _, v := range verbs {
		c, err := m.Conjugate(v)
		if err != nil {
			return err
		}
		gs, err := sentences.GenerateSentences(v, c)
		if err != nil {
			return err
		}
		for _, s := range gs {
			if lines >= 2000 {
				lines = 0
				index++
				of, err = getFile(ofn, ofe, index)
				if err != nil {
					return err
				}
			}
			fmt.Fprint(of, s.English, ";", s.Português, "\n")
			lines++
		}
	}
	of.Close()
	return nil
}

func getFile(ofn, ofe string, index int) (*os.File, error) {
	of, err := os.OpenFile(fmt.Sprintf("%s-%d%s", ofn, index, ofe), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return of, nil
}
