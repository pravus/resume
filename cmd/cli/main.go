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
	resume   = `data/resume.yaml`
	template = `html/content.html`
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
		loader := yaml.NewLoader(_prefix(resume))
		deps.Loader = loader
	}

	// writer
	{
		writer := html.NewWriter(_prefix(template))
		deps.Writer = writer
	}

	// main
	if err := app.NewApp(deps).MangéLaTête(); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
}
