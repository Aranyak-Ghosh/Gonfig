package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Aranyak-Ghosh/gonfig/types"
)

type jsonFileProvider struct {
	fileName string
}

type jsonStringProvider struct {
	serializedJson string
}

type jsonType int

const (
	jSONObject jsonType = iota
	jSONArray
	invalidJson
)

var _ types.Provider = (*jsonFileProvider)(nil)
var _ types.Provider = (*jsonStringProvider)(nil)

const blank = " \t\r\n"

func isNullOrEmpty(str string) bool {
	return strings.Trim(str, blank) == ""
}

func isValidJson(data []byte) jsonType {
	trimmed := bytes.Trim(data, blank)

	switch trimmed[0] {
	case '[':
		return jSONArray
	case '{':
		return jSONObject
	default:
		return invalidJson
	}
}

func (jp *jsonFileProvider) Load(res map[string]any) error {

	if !isNullOrEmpty(jp.fileName) {
		if data, err := ioutil.ReadFile(jp.fileName); err != nil {
			return err
		} else {
			switch isValidJson(data) {
			case jSONObject:
				json.Unmarshal(data, &res)
			case jSONArray:
				return fmt.Errorf("Array json not supported")
			default:
				return fmt.Errorf("Not a valid json body")
			}
		}

	}

	return nil
}

func (jp *jsonStringProvider) Load(res map[string]any) error {

	if !isNullOrEmpty(jp.serializedJson) {
		data := []byte(jp.serializedJson)
		switch isValidJson(data) {
		case jSONObject:
			json.Unmarshal(data, &res)
		case jSONArray:
			return fmt.Errorf("Array json not supported")
		default:
			return fmt.Errorf("Not a valid json body")

		}

	}
	return nil
}

func NewJsonFileProvider(filename string) types.Provider {
	return &jsonFileProvider{fileName: filename}
}

func NewJsonStringProvider(jsonString string) types.Provider {
	return &jsonStringProvider{serializedJson: jsonString}
}
