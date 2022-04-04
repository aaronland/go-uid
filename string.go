package uid

import (
	"context"
	"errors"
	"net/url"
)

func init() {
	ctx := context.Background()
	RegisterProvider(ctx, "string", NewStringProvider)
}

type StringProvider struct {
	Provider
	string string
}

type StringUID struct {
	UID
	string string
}

func NewStringProvider(ctx context.Context, uri string) (Provider, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()
	s := q.Get("string")

	if s == "" {
		return nil, errors.New("Empty string")
	}

	pr := &StringProvider{
		string: s,
	}

	return pr, nil
}

func (pr *StringProvider) UID(...interface{}) (UID, error) {
	return NewStringUID(pr.string)
}

func NewStringUID(s string) (UID, error) {

	u := StringUID{
		string: s,
	}

	return &u, nil
}

// where is UID() ?

func (u *StringUID) String() string {
	return u.string
}
