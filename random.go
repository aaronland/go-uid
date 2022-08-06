package uid

import (
	"context"
	"github.com/aaronland/go-string/random"
)

const RANDOM_SCHEME string = "random"

func init() {
	ctx := context.Background()
	RegisterProvider(ctx, RANDOM_SCHEME, NewRandomProvider)
}

type RandomProvider struct {
	Provider
}

func NewRandomProvider(ctx context.Context, uri string) (Provider, error) {
	pr := &RandomProvider{}
	return pr, nil
}

func (pr *RandomProvider) UID(...interface{}) (UID, error) {

	opts := random.DefaultOptions()
	opts.Length = 16
	opts.Chars = 16
	opts.ASCII = true
	opts.AlphaNumeric = true

	s, err := random.String(opts)

	if err != nil {
		return nil, err
	}

	return NewStringUID(s)
}
