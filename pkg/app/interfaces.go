package app

type Loader interface {
	Load() (*Resume, Error)
}

type Writer interface {
	Write(*Resume) Error
}
