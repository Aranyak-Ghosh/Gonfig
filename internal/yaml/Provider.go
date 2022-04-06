package yaml

import (
	"io/ioutil"
	"strings"

	"github.com/Aranyak-Ghosh/gonfig/types"
	"gopkg.in/yaml.v2"
)

type yamlProvider struct {
	fileName    string
	fileContent string
}

var _ types.Provider = (*yamlProvider)(nil)

const blank = " \t\r\n"

func isNullOrEmpty(str string) bool {
	return strings.Trim(str, blank) == ""
}

func (yp *yamlProvider) Load(res map[string]interface{}) error {
	if !isNullOrEmpty(yp.fileName) {
		if data, err := ioutil.ReadFile(yp.fileName); err != nil {
			return err
		} else {
			err := yaml.Unmarshal(data, &res)
			if err != nil {
				return err
			}
		}
	} else if !isNullOrEmpty(yp.fileContent) {
		err := yaml.Unmarshal([]byte(yp.fileContent), &res)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewYamlFileProvider(filename string) types.Provider {
	return &yamlProvider{
		fileName: filename,
	}
}

func NewYamlStringProvider(content string) types.Provider {
	return &yamlProvider{
		fileContent: content,
	}
}
