package dotenv

import (
	"os"
	"reflect"
	"strings"

	"github.com/Aranyak-Ghosh/gonfig/types"
	"github.com/joho/godotenv"
)

type dotEnvProvider struct {
	fileName string
	filePath string
}

var _ types.Provider = (*dotEnvProvider)(nil)

const separator = "__"
const blank = " \t\r\n"

func isNullOrEmpty(str string) bool {
	return strings.Trim(str, blank) == ""
}

func (de *dotEnvProvider) Load(data map[string]interface{}) error {
	path := string(append([]byte(de.filePath), de.fileName...))

	if !isNullOrEmpty(path) {
		if err := godotenv.Load(path); err != nil {
			return err
		}
	} else {
		if err := godotenv.Load(); err != nil {
			return err
		}
	}
	rawData := os.Environ()

	for _, d := range rawData {
		vars := strings.Split(d, "=")

		if len(vars) == 2 {
			handleNesting(vars[0], vars[1], data)
		} else if len(vars) > 2 {
			handleNesting(vars[0], strings.Join(vars[1:], separator), data)
		}

	}

	return nil
}

func handleNesting(key string, val string, dat map[string]interface{}) {
	var nestedKeys = strings.Split(key, separator)

	if len(nestedKeys) > 1 {
		if v, ok := dat[nestedKeys[0]]; !ok {
			dat[nestedKeys[0]] = make(map[string]interface{})
		} else {
			inType := reflect.TypeOf(v)

			if inType.Kind() != reflect.Map {
				dat[nestedKeys[0]] = make(map[string]interface{})
			}
		}
		handleNesting(strings.Join(nestedKeys[1:], separator), val, dat[nestedKeys[0]].(map[string]interface{}))
	} else {
		spl := strings.Split(val, separator)
		if len(spl) > 1 {
			dat[key] = spl
		} else {
			dat[key] = spl[0]
		}
	}
}

func NewDotEnvProvider(fileName string, filePath string) types.Provider {
	return &dotEnvProvider{
		fileName: fileName,
		filePath: filePath,
	}
}
