package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jeremy-spitzig/conjugation-generator/sentences"
	"github.com/urfave/cli/v2"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language-pack",
				Aliases: []string{"lp"},
				Value: "./default-language-pack/",
				Usage: "Load input from `FILE`",
			},
			&cli.StringFlag{
				Name:  "outputFile",
				Aliases: []string{"o"},
				Value: "output.csv",
				Usage: "The output file name",
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

	lpd := c.String("language-pack")

	md := filepath.Join(lpd, "models")
	m, err := verbs.LoadModels(md)
	if err != nil {
		return err
	}

	td := filepath.Join(lpd, "templates")
	t, err := sentences.LoadTemplates(td)
	if err != nil {
		return err
	}

	vfn := filepath.Join(lpd, "verbs.json")
	verbs, err := verbs.ReadVerbs(vfn)
	if err != nil {
		return err
	}

	ofn := c.String("outputFile")
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
