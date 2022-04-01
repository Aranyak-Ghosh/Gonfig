package types

type SourceType int

const (
	ENV SourceType = iota
	DotEnv
	YAML
	JSON
)

type Provider struct {
	Type     SourceType
	FilePath string
}
