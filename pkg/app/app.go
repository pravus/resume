package app

type App struct {
	loader Loader
	writer Writer
}

type AppConfiguration struct {
	Loader Loader
	Writer Writer
}

const (
	ErrorCodeAppResumeLoad  ErrorCode = `AppResumeLoad`
	ErrorCodeAppResumeWrite           = `AppResumeWrite`
)

func NewApp(configuration *AppConfiguration) *App {
	var app = &App{
		loader: configuration.Loader,
		writer: configuration.Writer,
	}
	return app
}

func (app *App) MangéLaTête() Error {
	if resume, err := app.loader.Load(); err != nil {
		return NewError(ErrorCodeAppResumeLoad, err)
	} else if err := app.writer.Write(resume); err != nil {
		return NewError(ErrorCodeAppResumeWrite, err)
	} else {
		return nil
	}
}
