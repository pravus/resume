package main

import (
	"fmt"
	"os"

	"resume/pkg/app"
	"resume/pkg/html"
	"resume/pkg/yaml"
)

const (
	dataResume = `data/resume.yaml`

	templateHtml = `html/content.html`
	templateCss  = `html/content.css`
	dataStyle    = `html/style.yaml`

	outputHtml = `resume.html`
	outputCss  = `resume.css`
)

func main() {
	var loader = yaml.NewLoader(dataResume)
	var writer = html.NewWriter(&html.WriterConfiguration{
		Output: &html.WriterFilePair{
			Html: outputHtml,
			Css:  outputCss,
		},
		Style: &html.WriterStyle{
			FileName:   dataStyle,
			Stylesheet: outputCss,
		},
		Template: &html.WriterFilePair{
			Html: templateHtml,
			Css:  templateCss,
		},
	})

	if err := app.NewApp(&app.AppConfiguration{
		Loader: loader,
		Writer: writer,
	}).MangéLaTête(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
