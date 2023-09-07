package html

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"resume/internal/model"
)

const (
	ErrorCodeHtmlTemplateParse   model.ErrorCode = `HtmlTemplateParse`
	ErrorCodeHtmlTemplateExecute model.ErrorCode = `HtmlTemplateExecute`
	ErrorCodeHtmlWrite           model.ErrorCode = `HtmlWrite`
)

type Writer struct {
	template string
}

var _ model.Writer = (*Writer)(nil)

func NewWriter(template string) Writer {
	var writer = Writer{
		template: template,
	}
	return writer
}

func (writer Writer) Write(w io.Writer, resume *model.Resume) model.Error {
	var data = &struct {
		Resume *model.Resume
	}{
		Resume: resume,
	}
	body := bytes.NewBuffer([]byte{})
	if mold, err := template.ParseFiles(writer.template); err != nil {
		return model.NewError(ErrorCodeHtmlTemplateParse, err)
	} else if err := mold.Execute(body, data); err != nil {
		return model.NewError(ErrorCodeHtmlTemplateExecute, err)
	} else if size := int64(body.Len()); size <= 0 {
		return model.NewError(ErrorCodeHtmlTemplateExecute, fmt.Errorf(`template output empty`))
	} else if count, err := io.Copy(w, body); err != nil {
		return model.NewError(ErrorCodeHtmlWrite, err)
	} else if count != size {
		return model.NewError(ErrorCodeHtmlWrite, fmt.Errorf(`expected to write %d bytes but wrote %d`, size, count))
	}
	return nil
}
