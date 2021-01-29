package main

import (
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
				Aliases: []string{"i"},
				Value: "./default-input.json",
				Usage: "Load input from `FILE`",
			},
			&cli.StringFlag{
				Name:  "outputFile",
				Aliases: []string{"o"},
				Value: "output.csv",
				Usage: "The output file name",
			},
			&cli.StringFlag{
				Name:  "modelsDir",
				Aliases: []string{"m"},
				Value: "./default-verb-models",
				Usage: "Directory containing verb model files",
			},
			&cli.StringFlag{
				Name:  "templatesDir",
				Aliases: []string{"t"},
				Value: "./default-sentence-templates",
				Usage: "Directory containing sentence template files",
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

	td := c.String("templatesDir")
	t, err := sentences.LoadTemplates(td)

	if err != nil {
		return err
	}

	ifn := c.String("input")
	ofn := c.String("outputFile")

	verbs, err := verbs.ReadVerbs(ifn)
	if err != nil {
		return err
	}

	of, err := os.OpenFile(ofn, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer of.Close()

	for _, v := range verbs {
		c, err := m.Conjugate(v)
		if err != nil {
			return err
		}
		err = t.Execute(v, c, of)
		if err != nil {
			return err
		}
	}
	return nil
}
