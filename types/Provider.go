package types

type SourceType int

type Provider interface {
	Load(map[string]interface{}) error
}
