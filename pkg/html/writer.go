package html

import (
	"html/template"
	"os"

	"github.com/go-yaml/yaml"

	"resume/pkg/app"
)

const (
	ErrorCodeHtmlOutputFileOpen  app.ErrorCode = `HtmlOutputFileOpen`
	ErrorCodeHtmlOutputFileClose               = `HtmlOutputFileClose`
	ErrorCodeHtmlTemplateParse                 = `HtmlTemplateParse`
	ErrorCodeHtmlTemplateExecute               = `HtmlTemplateExecute`
	ErrorCodeHtmlStyleDecode                   = `HtmlStyleDecode`
	ErrorCodeHtmlStyleFileOpen                 = `HtmlStyleFileOpen`
	ErrorCodeHtmlStyleFileClose                = `HtmlStyleFileClose`
)

type Writer struct {
	output   string
	style    string
	template string
}

type Style struct {
	Palette struct {
		Background template.CSS `json:"background" yaml:"background"`
		Text       template.CSS `json:"text"       yaml:"text"`
		Borders    template.CSS `json:"borders"    yaml:"borders"`
		Hilite     template.CSS `json:"hilite"     yaml:"hilite"`
	} `json:"palette" yaml:"palette"`
}

func NewWriter(output, style, template string) Writer {
	var writer = Writer{
		output:   output,
		style:    style,
		template: template,
	}
	return writer
}

func (writer Writer) Write(resume *app.Resume) app.Error {
	if style, err := writer.loadStyle(); err != nil {
		return err
	} else if err := writer.doTemplate(resume, style); err != nil {
		return err
	}
	return nil
}

func (writer Writer) loadStyle() (*Style, app.Error) {
	var style *Style
	if file, err := os.Open(writer.style); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleFileOpen, err)
	} else if err := yaml.NewDecoder(file).Decode(&style); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleDecode, err)
	} else if err := file.Close(); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleFileClose, err)
	} else {
		return style, nil
	}
}

func (writer Writer) doTemplate(resume *app.Resume, style *Style) app.Error {
	var data = &struct {
		Resume *app.Resume
		Style  *Style
	}{
		Resume: resume,
		Style:  style,
	}
	if mold, err := template.ParseFiles(writer.template); err != nil {
		return app.NewError(ErrorCodeHtmlTemplateParse, err)
	} else if file, err := os.Create(writer.output); err != nil {
		return app.NewError(ErrorCodeHtmlOutputFileOpen, err)
	} else if err := mold.Execute(file, data); err != nil {
		return app.NewError(ErrorCodeHtmlTemplateExecute, err)
	} else if err := file.Close(); err != nil {
		return app.NewError(ErrorCodeHtmlOutputFileClose, err)
	}
	return nil
}
