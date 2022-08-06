package uid

import (
	"context"
)

const NULL_SCHEME string = "null"

func init() {
	ctx := context.Background()
	RegisterProvider(ctx, NULL_SCHEME, NewNullProvider)
}

type NullProvider struct {
	Provider
}

type NullUID struct {
	UID
}

func NewNullProvider(ctx context.Context, uri string) (Provider, error) {
	pr := &NullProvider{}
	return pr, nil
}

func (n *NullProvider) UID(...interface{}) (UID, error) {
	return NewNullUID()
}

func NewNullUID() (UID, error) {
	n := &NullUID{}
	return n, nil
}

func (n *NullUID) value() any {
	return ""
}
