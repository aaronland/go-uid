package uid

import (
	"context"
)

type StringProvider[T uid] struct {
	Provider[T]
}

func NewStringProvider[T uid](ctx context.Context, uri string) (Provider[T], error) {

	pr := &StringProvider[T]{}
	return pr, nil
}

func (pr *StringProvider[T]) NewId() (Identifier[T], error) {

	s := "debug"
	return NewUniqueIdentifier(s)
}
