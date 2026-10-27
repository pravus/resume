package html

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"

	"resume/internal/model"
)

const (
	ErrorCodeHtmlTemplateParse   model.ErrorCode = "HtmlTemplateParse"
	ErrorCodeHtmlTemplateExecute model.ErrorCode = "HtmlTemplateExecute"
	ErrorCodeHtmlWrite           model.ErrorCode = "HtmlWrite"
)

type Writer struct{ templatePath string }

var _ model.Writer = Writer{}

func NewWriter(templatePath string) Writer { return Writer{templatePath: templatePath} }

func (writer Writer) Write(dst io.Writer, resume *model.Resume) model.Error {
	source, err := os.ReadFile(writer.templatePath)
	if err != nil {
		return model.NewError(ErrorCodeHtmlTemplateParse, fmt.Errorf("%s: %w", writer.templatePath, err))
	}
	tmpl, err := template.New("resume").Parse(string(source))
	if err != nil {
		return model.NewError(ErrorCodeHtmlTemplateParse, err)
	}

	// Render before writing so a template execution error cannot leave partial HTML.
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, struct{ Resume *model.Resume }{Resume: resume}); err != nil {
		return model.NewError(ErrorCodeHtmlTemplateExecute, err)
	}
	if rendered.Len() == 0 {
		return model.NewError(ErrorCodeHtmlTemplateExecute, fmt.Errorf("template output is empty"))
	}
	if _, err := io.Copy(dst, &rendered); err != nil {
		return model.NewError(ErrorCodeHtmlWrite, err)
	}
	return nil
}
