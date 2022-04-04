package uid

import (
	"context"
)

func init() {
	ctx := context.Background()
	RegisterProvider(ctx, "null", NewNullProvider)
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

func (n *NullUID) String() string {
	return ""
}
