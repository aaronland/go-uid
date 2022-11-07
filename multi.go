package uid

import (
	"context"
	"fmt"
	"strings"
)

type MultiUID struct {
	UID
	uids []UID
}

func NewMultiUID(ctx context.Context, uids ...UID) UID {

	r := &MultiUID{
		uids: uids,
	}

	return r
}

func (r *MultiUID) Value() any {
	return r.uids
}

func (r *MultiUID) String() string {

	pairs := make([]string, len(r.uids))

	for idx, uid := range r.uids {
		pairs[idx] = fmt.Sprintf("%T#%s", uid, uid.String())
	}

	return strings.Join(pairs, " ")
}
