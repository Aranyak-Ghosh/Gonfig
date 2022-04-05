package types

type SourceType int

const (
	ENV SourceType = iota
	DotEnv
	YAML
	JSON
	INI
	TOML
)

type Provider interface {
	Load(map[string]interface{}) error
}
