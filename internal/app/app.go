package app

import (
	"os"

	"resume/internal/model"
)

type App struct {
	loader model.Loader
	writer model.Writer
}

type Dependencies struct {
	Loader model.Loader
	Writer model.Writer
}

const (
	ErrorCodeAppResumeLoad  model.ErrorCode = `AppResumeLoad`
	ErrorCodeAppResumeWrite model.ErrorCode = `AppResumeWrite`
)

func NewApp(deps *Dependencies) *App {
	var app = &App{
		loader: deps.Loader,
		writer: deps.Writer,
	}
	return app
}

func (app *App) MangéLaTête() model.Error {
	if resume, err := app.loader.Load(); err != nil {
		return model.NewError(ErrorCodeAppResumeLoad, err)
	} else if err := app.writer.Write(os.Stdout, resume); err != nil {
		return model.NewError(ErrorCodeAppResumeWrite, err)
	} else {
		return nil
	}
}
