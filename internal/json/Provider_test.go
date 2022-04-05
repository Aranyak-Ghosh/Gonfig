package json

import "testing"

func Test_JsonFileConfig(t *testing.T) {
	file := "config.json"

	js := NewJsonFileProvider(file)

	dat := make(map[string]interface{})

	js.Load(dat)

	if val, ok := dat["Foo"]; !ok || val.(string) != "Bar" {
		t.Fatal("Not valid")
	}
}

func Test_JsonStringConfig(t *testing.T) {
	data := "{\"Foo\":\"Bar\"}"

	js := NewJsonStringProvider(data)

	dat := make(map[string]interface{})

	js.Load(dat)

	if val, ok := dat["Foo"]; !ok || val.(string) != "Bar" {
		t.Fatal("Not valid")
	}
}
