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

func (n *NullProvider) New() T {
	return ""
}
