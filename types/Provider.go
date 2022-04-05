package types

type SourceType int

const (
	DotEnv SourceType = iota
	YAML
	JSON
	INI
	TOML
	HCL
)

type Provider interface {
	Load(map[string]interface{}) error
}
