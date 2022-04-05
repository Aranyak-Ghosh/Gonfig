package gonfig

import (
	"github.com/Aranyak-Ghosh/golist"
	"github.com/Aranyak-Ghosh/gonfig/types"
)

type configManager struct {
	providers golist.List[types.Provider]
	config    map[string]interface{}
}

type ConfigManager interface {
	AddProvider(provider types.Provider)
	ParseConfig(any)
	ReloadConfig(any)
}

var _ ConfigManager = (*configManager)(nil)

func (cm *configManager) AddProvider(provider types.Provider) {}
func (cm *configManager) ParseConfig(any)                     {}
func (cm *configManager) ReloadConfig(any)                    {}
