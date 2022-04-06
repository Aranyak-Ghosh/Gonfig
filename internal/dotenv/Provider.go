package dotenv

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Aranyak-Ghosh/golist"
	"github.com/Aranyak-Ghosh/gonfig/types"
	"github.com/joho/godotenv"
)

type dotEnvProvider struct {
	fileName string
}

var _ types.Provider = (*dotEnvProvider)(nil)

const separator = "__"
const blank = " \t\r\n"

func isNullOrEmpty(str string) bool {
	return strings.Trim(str, blank) == ""
}

func (de *dotEnvProvider) Load(data map[string]any) error {

	if !isNullOrEmpty(de.fileName) {
		if err := godotenv.Load(de.fileName); err != nil {
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

		if len(vars) >= 2 {
			handleKeyNesting(vars[0], strings.Join(vars[1:], separator), data)
		}

	}

	return nil
}

func handleKeyNesting(key string, val string, dat any) {
	isDataMap := isMap(dat)
	if isDataMap {
		handleKeyNestingForMap(key, val, dat.(map[string]any))
	} else {
		handleKeyNestingForSlice(key, val, dat.(*golist.List[any]))
	}
}

func handleKeyNestingForMap(key, val string, dat map[string]any) {
	var nestedKeys = strings.Split(key, separator)

	if len(nestedKeys) > 1 {
		if _, ok := dat[nestedKeys[0]]; !ok {
			if _, err := strconv.Atoi(nestedKeys[1]); err != nil {
				dat[nestedKeys[0]] = make(map[string]any)
			} else {
				dat[nestedKeys[0]] = new(golist.List[any])
			}
		}
		handleKeyNesting(strings.Join(nestedKeys[1:], separator), val, dat[nestedKeys[0]])
	} else {
		dat[key] = val
	}
}

func handleKeyNestingForSlice(key, val string, dat *golist.List[any]) {
	var nestedKeys = strings.Split(key, separator)

	if len(nestedKeys) == 1 {
		dat.Append(val)
	} else {
		nestedVal := make(map[string]any)
		handleKeyNesting(strings.Join(nestedKeys[1:], separator), val, nestedVal)
		dat.Append(nestedVal)
	}
}

func isMap(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

// Pass an empty string to load default env file
func NewDotEnvProvider(fileName string) types.Provider {
	return &dotEnvProvider{
		fileName: fileName,
	}
}
