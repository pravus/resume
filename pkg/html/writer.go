package html

import (
	"html/template"
	"os"

	"resume/pkg/app"
)

const (
	ErrorCodeHtmlOutputFileOpen  app.ErrorCode = `HtmlOutputFileOpen`
	ErrorCodeHtmlOutputFileClose app.ErrorCode = `HtmlOutputFileClose`
	ErrorCodeHtmlTemplateParse   app.ErrorCode = `HtmlTemplateParse`
	ErrorCodeHtmlTemplateExecute app.ErrorCode = `HtmlTemplateExecute`
)

type Writer struct {
	output   string
	template string
}

var _ app.Writer = (*Writer)(nil)

func NewWriter(output, template string) Writer {
	var writer = Writer{
		output:   output,
		template: template,
	}
	return writer
}

func (writer Writer) Write(resume *app.Resume) app.Error {
	var data = &struct {
		Resume *app.Resume
	}{
		Resume: resume,
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
