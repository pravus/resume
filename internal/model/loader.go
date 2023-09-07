package model

type Loader interface {
	Load() (*Resume, Error)
}
