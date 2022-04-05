package yaml

import "testing"

func Test_YamlFileConfig(t *testing.T) {
	file := "config.yaml"

	js := NewYamlProvider(file)

	dat := make(map[string]interface{})

	js.Load(dat)

	if val, ok := dat["Foo"]; !ok || val.(string) != "Bar" {
		t.Fatal("Not valid")
	}
}
