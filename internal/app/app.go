package app

import (
	"io"
	"resume/internal/model"
)

type App struct {
	loader model.Loader
	writer model.Writer
	output io.Writer
}

type Dependencies struct {
	Loader model.Loader
	Writer model.Writer
	Output io.Writer
}

const (
	ErrorCodeAppResumeLoad  model.ErrorCode = "AppResumeLoad"
	ErrorCodeAppResumeWrite model.ErrorCode = "AppResumeWrite"
)

func NewApp(deps Dependencies) *App {
	output := deps.Output
	if output == nil {
		output = io.Discard
	}
	return &App{loader: deps.Loader, writer: deps.Writer, output: output}
}

// Run loads the source data and renders the resume to the configured output.
func (app *App) Run() model.Error {
	resume, err := app.loader.Load()
	if err != nil {
		return model.NewError(ErrorCodeAppResumeLoad, err)
	}
	if err := app.writer.Write(app.output, resume); err != nil {
		return model.NewError(ErrorCodeAppResumeWrite, err)
	}
	return nil
}
