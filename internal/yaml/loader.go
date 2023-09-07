package yaml

import (
	"os"

	"github.com/go-yaml/yaml"

	"resume/internal/model"
)

const (
	ErrorCodeYamlDecode    model.ErrorCode = `YamlDecode`
	ErrorCodeYamlFileOpen  model.ErrorCode = `YamlFileOpen`
	ErrorCodeYamlFileClose model.ErrorCode = `YamlFileClose`
)

type Loader struct {
	fileName string
}

var _ model.Loader = (*Loader)(nil)

func NewLoader(fileName string) Loader {
	var loader = Loader{
		fileName: fileName,
	}
	return loader
}

func (loader Loader) Load() (*model.Resume, model.Error) {
	var resume *model.Resume
	if file, err := os.Open(loader.fileName); err != nil {
		return nil, model.NewError(ErrorCodeYamlFileOpen, err)
	} else if err := yaml.NewDecoder(file).Decode(&resume); err != nil {
		return nil, model.NewError(ErrorCodeYamlDecode, err)
	} else if err := file.Close(); err != nil {
		return nil, model.NewError(ErrorCodeYamlFileClose, err)
	} else {
		return resume, nil
	}
}
