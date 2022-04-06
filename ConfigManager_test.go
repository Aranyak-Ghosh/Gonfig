package gonfig

import "testing"

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

	var expected = "Bars"

	err := cm.AddProvider(Provider{
		ProviderType: JSON,
		FileName:     "appsettings.json",
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
