package uid

import (
	"context"
	"github.com/aaronland/go-uid/driver"
	"net/url"
	"log"
)

var roster driver.Roster

func init() {

	log.Println("UID INIT")
	r, err := driver.NewDefaultRoster()

	if err != nil {
		panic(err)
	}

	log.Println("HELLO", r)
	roster = r
}

func RegisterProvider(ctx context.Context, name string, pr Provider) error {
	return roster.Register(ctx, name, pr)
}

func NewProvider(ctx context.Context, uri string) (Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	pr := i.(Provider)
	err = pr.Open(ctx, uri)

	if err != nil {
		return nil, err
	}

	return pr, nil
}

type Provider interface {
	Open(context.Context, string) error
	UID(...interface{}) (UID, error) // NOT SURE ABOUT THIS...
}

type UID interface {
	String() string
}
