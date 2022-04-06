package gonfig

import (
	"fmt"
	"strings"

	"github.com/Aranyak-Ghosh/golist"
	"github.com/Aranyak-Ghosh/gonfig/internal/dotenv"
	"github.com/Aranyak-Ghosh/gonfig/internal/json"
	"github.com/Aranyak-Ghosh/gonfig/internal/yaml"
	"github.com/mitchellh/mapstructure"

	"github.com/Aranyak-Ghosh/gonfig/types"
)

type configManager struct {
	providers *golist.List[types.Provider]
	config    map[string]any
}

const (
	DotEnv types.SourceType = iota
	YAML
	JSON
	// INI
	// TOML
	// HCL
)

const separator = "__"

type Provider struct {
	ProviderType types.SourceType
	FileName     string
	FileContent  string
}

type ConfigManager interface {
	AddProvider(provider Provider) error
	GetConfig(string) (any, error)
	MapConfig(string, any) error
	ReloadConfig() error
}

var _ ConfigManager = (*configManager)(nil)

func (cm *configManager) AddProvider(provider Provider) error {
	var err error = nil
	switch provider.ProviderType {
	case DotEnv:
		pr := dotenv.NewDotEnvProvider(provider.FileName)
		err = cm.loadProvider(pr)
	case YAML:
		if !isNullOrEmpty(provider.FileName) {
			pr := yaml.NewYamlFileProvider(provider.FileName)
			err = cm.loadProvider(pr)
		} else if !isNullOrEmpty(provider.FileContent) {
			pr := yaml.NewYamlStringProvider(provider.FileContent)
			err = cm.loadProvider(pr)
		}
	case JSON:
		if !isNullOrEmpty(provider.FileName) {
			pr := json.NewJsonFileProvider(provider.FileName)
			err = cm.loadProvider(pr)
		} else if !isNullOrEmpty(provider.FileContent) {
			pr := json.NewJsonStringProvider(provider.FileContent)
			err = cm.loadProvider(pr)
		}
	}
	return err
}
func (cm *configManager) GetConfig(key string) (any, error) {
	keys := strings.Split(key, separator)
	var res any = cm.config
	var ok bool
	for _, k := range keys {
		res, ok = res.(map[string]any)[k]
		if !ok {
			return nil, fmt.Errorf("Key not found")
		}
	}

	return res, nil
}
func (cm *configManager) MapConfig(key string, out any) error {
	keys := strings.Split(key, separator)
	var res any = cm.config
	var ok bool
	for _, k := range keys {
		res, ok = res.(map[string]any)[k]
		if !ok {
			return fmt.Errorf("Key not found")
		}
	}
	return mapstructure.Decode(res, &out)
}
func (cm *configManager) ReloadConfig() error {
	var e error = nil
	for _, val := range *cm.providers {
		err := cm.loadProvider(val)
		if err != nil {
			e = fmt.Errorf("%w", err)
		}
	}
	if e != nil {
		e = fmt.Errorf("One of more errors occured while loading the config: %w", e)
	}
	return e
}

func NewConfigManager() *configManager {
	return &configManager{
		providers: new(golist.List[types.Provider]),
		config:    make(map[string]any),
	}
}

var blankString = " \t\r\n"

func isNullOrEmpty(data string) bool {
	return strings.Trim(data, blankString) == ""
}

func (cm *configManager) loadProvider(pr types.Provider) error {
	if cm.config == nil {
		cm.config = make(map[string]any)
	}
	err := pr.Load(cm.config)
	if err != nil {
		return err
	}
	cm.providers.Append(pr)
	return nil
}
