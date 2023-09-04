package main

import (
	"fmt"
	"os"

	"resume/pkg/app"
	"resume/pkg/html"
	"resume/pkg/yaml"
)

const (
	yamlResume   = `data/resume.yaml`
	htmlOutput   = `resume.html`
	htmlTemplate = `html/content.html`
)

func main() {
	var loader = yaml.NewLoader(yamlResume)
	var writer = html.NewWriter(htmlOutput, htmlTemplate)
	if err := app.NewApp(&app.AppConfiguration{
		Loader: loader,
		Writer: writer,
	}).MangéLaTête(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
