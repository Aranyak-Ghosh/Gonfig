package types

type SourceType int

type Provider interface {
	Load(map[string]any) error
}
