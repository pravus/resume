package main

import (
	"flag"
	"fmt"
	"os"

	"resume/internal/app"
	"resume/internal/html"
	"resume/internal/yaml"
)

const (
	yamlResume   = `data/resume.yaml`
	htmlOutput   = `resume.html`
	htmlTemplate = `html/content.html`
)

type flags struct {
	prefix *string
}

func main() {
	flags := flags{
		prefix: flag.String(`prefix`, ``, `sets the template prefix`),
	}
	flag.Parse()

	// helpers
	_prefix := func(name string) string {
		if *flags.prefix != `` {
			return *flags.prefix + `/` + name
		}
		return name
	}

	// dependencies
	deps := &app.Dependencies{}

	// loader
	{
		loader := yaml.NewLoader(_prefix(yamlResume))
		deps.Loader = loader
	}

	// writer
	{
		writer := html.NewWriter(htmlOutput, _prefix(htmlTemplate))
		deps.Writer = writer
	}

	// main
	if err := app.NewApp(deps).MangéLaTête(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
