package uid

type Provider[T uid] interface {
	NewId() (Identifier[T], error)
}
