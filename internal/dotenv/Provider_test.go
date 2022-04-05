package dotenv

import "testing"

func Test_LoadDefaultEnv(t *testing.T) {
	provider := NewDotEnvProvider("", "")

	if provider == nil {
		t.Error("provider is nil")
	}

	mapData := make(map[string]interface{})

	if err := provider.Load(mapData); err != nil {
		t.Errorf("error loading data: %s", err)
	}

	t.Logf("mapData: %+v", mapData)
}

func Test_LoadSpecificEnv(t *testing.T) {
	provider := NewDotEnvProvider(".env.named", "")

	if provider == nil {
		t.Error("provider is nil")
	}

	mapData := make(map[string]interface{})

	if err := provider.Load(mapData); err != nil {
		t.Errorf("error loading data: %s", err)
	}

	t.Logf("mapData: %+v", mapData)
}
