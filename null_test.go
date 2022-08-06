package uid

import (
	"context"
	"fmt"
	"testing"
)

func TestNullProvider(t *testing.T) {

	ctx := context.Background()

	uri := fmt.Sprintf("%s://", NULL_SCHEME)

	pr, err := NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create null UID, %v", err)
	}

	if id.String() != "" {
		t.Fatalf("Invalid ID for null provider")
	}
}
