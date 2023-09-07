package html

import (
	"html/template"
	"os"

	"resume/internal/model"
)

const (
	ErrorCodeHtmlOutputFileOpen  model.ErrorCode = `HtmlOutputFileOpen`
	ErrorCodeHtmlOutputFileClose model.ErrorCode = `HtmlOutputFileClose`
	ErrorCodeHtmlTemplateParse   model.ErrorCode = `HtmlTemplateParse`
	ErrorCodeHtmlTemplateExecute model.ErrorCode = `HtmlTemplateExecute`
)

type Writer struct {
	output   string
	template string
}

var _ model.Writer = (*Writer)(nil)

func NewWriter(output, template string) Writer {
	var writer = Writer{
		output:   output,
		template: template,
	}
	return writer
}

func (writer Writer) Write(resume *model.Resume) model.Error {
	var data = &struct {
		Resume *model.Resume
	}{
		Resume: resume,
	}
	if mold, err := template.ParseFiles(writer.template); err != nil {
		return model.NewError(ErrorCodeHtmlTemplateParse, err)
	} else if file, err := os.Create(writer.output); err != nil {
		return model.NewError(ErrorCodeHtmlOutputFileOpen, err)
	} else if err := mold.Execute(file, data); err != nil {
		return model.NewError(ErrorCodeHtmlTemplateExecute, err)
	} else if err := file.Close(); err != nil {
		return model.NewError(ErrorCodeHtmlOutputFileClose, err)
	}
	return nil
}
