package uid

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("Failed to parse string, %w", err)
	}

	q := u.Query()
	s := q.Get("string")

	if s == "" {
		return nil, fmt.Errorf("Empty string")
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

func (u *StringUID) Value() any {
	return u.string
}
