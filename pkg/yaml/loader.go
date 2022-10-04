package yaml

import (
	"os"

	"github.com/go-yaml/yaml"

	"resume/pkg/app"
)

const (
	ErrorCodeYamlDecode    app.ErrorCode = `YamlDecode`
	ErrorCodeYamlFileOpen                = `YamlFileOpen`
	ErrorCodeYamlFileClose               = `YamlFileClose`
)

type Loader struct {
	fileName string
}

func NewLoader(fileName string) *Loader {
	var loader = &Loader{
		fileName: fileName,
	}
	return loader
}

func (loader *Loader) Load() (*app.Resume, app.Error) {
	var resume *app.Resume
	if file, err := os.Open(loader.fileName); err != nil {
		return nil, app.NewError(ErrorCodeYamlFileOpen, err)
	} else if err := yaml.NewDecoder(file).Decode(&resume); err != nil {
		return nil, app.NewError(ErrorCodeYamlDecode, err)
	} else if err := file.Close(); err != nil {
		return nil, app.NewError(ErrorCodeYamlFileClose, err)
	} else {
		return resume, nil
	}
}
