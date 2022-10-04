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

type HtmlResume struct {
	Resume     *app.Resume
	Stylesheet string
}

type HtmlStyle struct {
	Palette struct {
		Background string `json:"background" yaml:"background"`
		Text       string `json:"text"       yaml:"text"`
		Borders    string `json:"borders"    yaml:"borders"`
		Hilite     string `json:"hilite"     yaml:"hilite"`
	} `json:"palette" yaml:"palette"`
}

type Writer struct {
	output   *WriterFilePair
	style    *WriterStyle
	template *WriterFilePair
}

type WriterConfiguration struct {
	Output   *WriterFilePair
	Style    *WriterStyle
	Template *WriterFilePair
}

type WriterFilePair struct {
	Html string
	Css  string
}

type WriterStyle struct {
	FileName   string
	Stylesheet string
}

func NewWriter(configuration *WriterConfiguration) *Writer {
	var writer = &Writer{
		output:   configuration.Output,
		style:    configuration.Style,
		template: configuration.Template,
	}
	return writer
}

func (writer *Writer) Write(resume *app.Resume) app.Error {
	var htmlResume = &HtmlResume{
		Resume:     resume,
		Stylesheet: writer.style.Stylesheet,
	}

	if style, err := loadStyle(writer.style.FileName); err != nil {
		return err
	} else if err := doTemplate(writer.template.Html, htmlResume, writer.output.Html); err != nil {
		return err
	} else if err := doTemplate(writer.template.Css, style, writer.output.Css); err != nil {
		return err
	} else {
		return nil
	}
}

func loadStyle(fileName string) (*HtmlStyle, app.Error) {
	var style *HtmlStyle
	if file, err := os.Open(fileName); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleFileOpen, err)
	} else if err := yaml.NewDecoder(file).Decode(&style); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleDecode, err)
	} else if err := file.Close(); err != nil {
		return nil, app.NewError(ErrorCodeHtmlStyleFileClose, err)
	} else {
		return style, nil
	}
}

func doTemplate(templateFileName string, templateData interface{}, outputFileName string) app.Error {
	if mold, err := template.ParseFiles(templateFileName); err != nil {
		return app.NewError(ErrorCodeHtmlTemplateParse, err)
	} else if file, err := os.Create(outputFileName); err != nil {
		return app.NewError(ErrorCodeHtmlOutputFileOpen, err)
	} else if err := mold.Execute(file, templateData); err != nil {
		return app.NewError(ErrorCodeHtmlTemplateExecute, err)
	} else if err := file.Close(); err != nil {
		return app.NewError(ErrorCodeHtmlOutputFileClose, err)
	} else {
		return nil
	}
}
