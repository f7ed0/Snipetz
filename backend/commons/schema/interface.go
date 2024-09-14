package schema

type ValidModel interface {
	AllFieldValid() bool
}
