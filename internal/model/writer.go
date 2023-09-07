package model

import (
	"io"
)

type Writer interface {
	Write(io.Writer, *Resume) Error
}
