package uid

import (
	"context"
	"github.com/aaronland/go-roster"
	"net/url"
)

type Provider interface {
	Open(context.Context, string) error
	UID(...interface{}) (UID, error)
}

type ProviderInitializationFunc func(ctx context.Context, uri string) (Provider, error)

var providers_roster roster.Roster

func RegisterProvider(ctx context.Context, scheme string, init_func ProviderInitializationFunc) error {

	err := ensureProviderRoster()

	if err != nil {
		return err
	}

	return providers_roster.Register(ctx, scheme, init_func)
}

func ensureProviderRoster() error {

	if providers_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		providers_roster = r
	}

	return nil
}

func NewProvider(ctx context.Context, uri string) (Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	err = ensureProviderRoster()

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := providers_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(ProviderInitializationFunc)
	return init_func(ctx, uri)
}

func Providers() []string {
	ctx := context.Background()
	return providers_roster.Drivers(ctx)
}
