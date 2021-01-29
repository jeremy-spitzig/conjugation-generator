package main

import (
	"github.com/jeremy-spitzig/conjugation-generator/languagepack"
	"github.com/jeremy-spitzig/conjugation-generator/sentences"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"

	"github.com/jeremy-spitzig/conjugation-generator/verbs"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "language-pack",
				Aliases: []string{"lp"},
				Value: "./default-language-pack/",
				Usage: "Load language pack from `SOURCE`",
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

	lps := c.String("language-pack")
	var lp languagepack.LanguagePack
	if(strings.HasPrefix(lps, "https://")) {
		glp, err := languagepack.NewGit(lps)
		if err != nil {
			return err
		}
		lp = glp
	} else {
		lp = languagepack.NewFileSystem(lps)
	}

	defer lp.Close()

	m, err := verbs.LoadModels(lp)
	if err != nil {
		return err
	}

	t, err := sentences.LoadTemplates(lp)
	if err != nil {
		return err
	}

	verbs, err := verbs.ReadVerbs(lp)
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
