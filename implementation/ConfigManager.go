package implementation

import "github.com/Aranyak-Ghosh/gonfig/types"

type configManager struct {
	providers []types.Provider
	config    any
}

type ConfigManager interface {
	AddProvider(provider types.Provider)
}
