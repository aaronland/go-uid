package uid

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestYmdProvider(t *testing.T) {

	ctx := context.Background()

	now := time.Now()
	ymd := now.Format("20060102")

	uri := fmt.Sprintf("%s://", YMD_SCHEME)

	pr, err := NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create ymd UID, %v", err)
	}

	if id.String() != ymd {
		t.Fatalf("Invalid ID for ymd provider: %s (expected %s)", ymd, id.String())
	}
}
