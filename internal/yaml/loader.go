package yaml

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"

	"resume/internal/model"
)

const (
	ErrorCodeYamlDecode   model.ErrorCode = "YamlDecode"
	ErrorCodeYamlFileOpen model.ErrorCode = "YamlFileOpen"
)

type Loader struct{ fileName string }

var _ model.Loader = Loader{}

func NewLoader(fileName string) Loader { return Loader{fileName: fileName} }

func (loader Loader) Load() (*model.Resume, model.Error) {
	file, err := os.Open(loader.fileName)
	if err != nil {
		return nil, model.NewError(ErrorCodeYamlFileOpen, fmt.Errorf("%s: %w", loader.fileName, err))
	}
	defer file.Close()

	var resume model.Resume
	if err := yaml.NewDecoder(file).Decode(&resume); err != nil {
		return nil, model.NewError(ErrorCodeYamlDecode, fmt.Errorf("%s: %w", loader.fileName, err))
	}
	if resume.Contact.FirstName == "" && resume.Contact.LastName == "" {
		return nil, model.NewError(ErrorCodeYamlDecode, fmt.Errorf("resume contact name is required"))
	}
	return &resume, nil
}
