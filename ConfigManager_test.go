package gonfig

import (
	"reflect"
	"testing"

	"github.com/Aranyak-Ghosh/golist"
	"github.com/mitchellh/mapstructure"
)

func TestAddEnvProvider(t *testing.T) {
	cm := NewConfigManager()

	var expected = "Bars"

	err := cm.AddProvider(Provider{
		ProviderType: DotEnv,
		FileName:     ".env.tests",
	})

	if err != nil {
		t.Errorf("Failed to add dotenv provider, %v", err)
	}

	dat, err := cm.GetConfig("Foo")

	if err != nil {
		t.Errorf("Failed to get expected key %s from dotenv provider, %v", "Foo", err)
	}

	if dat.(string) != expected {
		t.Errorf("Expected %s but got %v", expected, dat)
	}

}

func TestAddJsonFileProvider(t *testing.T) {
	cm := NewConfigManager()

	type conf struct {
		FooAgain []string
	}

	var got golist.List[string]

	var expected = []string{"Bars"}

	err := cm.AddProvider(Provider{
		ProviderType: JSON,
		FileName:     "appsettings.json",
	})

	if err != nil {
		t.Errorf("Failed to add dotenv provider, %v", err)
	}

	dat, err := cm.GetConfig("Foo__FooAgain")

	if err != nil {
		t.Errorf("Failed to get expected key %s from dotenv provider, %v", "Foo", err)
	}

	mapstructure.Decode(dat, &got)

	if reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %s but got %v", expected, dat)
	}
}

func TestMapJsonFileProvider(t *testing.T) {
	cm := NewConfigManager()

	type conf struct {
		FooAgain []string
	}

	var got golist.List[string]

	var expected = []string{"Bars"}

	err := cm.AddProvider(Provider{
		ProviderType: JSON,
		FileName:     "appsettings.json",
	})

	if err != nil {
		t.Errorf("Failed to add dotenv provider, %v", err)
	}

	err = cm.MapConfig("Foo__FooAgain", &got)

	if err != nil {
		t.Errorf("Failed to get expected key %s from dotenv provider, %v", "Foo", err)
	}

	if reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %s but got %v", expected, got)
	}
}

func TestAddYamlFileProvider(t *testing.T) {
	cm := NewConfigManager()

	var expected = "Bars"

	err := cm.AddProvider(Provider{
		ProviderType: YAML,
		FileName:     "appsettings.yaml",
	})

	if err != nil {
		t.Errorf("Failed to add dotenv provider, %v", err)
	}

	dat, err := cm.GetConfig("Foo")

	if err != nil {
		t.Errorf("Failed to get expected key %s from dotenv provider, %v", "Foo", err)
	}

	if dat.(string) != expected {
		t.Errorf("Expected %s but got %v", expected, dat)
	}
}
