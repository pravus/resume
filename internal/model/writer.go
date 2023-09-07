package model

type Writer interface {
	Write(*Resume) Error
}
