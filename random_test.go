package uid

import (
	"context"
	"fmt"
	"testing"
)

func TestRandomProvider(t *testing.T) {

	ctx := context.Background()

	uri := fmt.Sprintf("%s://", RANDOM_SCHEME)

	pr, err := NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create random UID, %v", err)
	}

	if id.String() == "" {
		t.Fatalf("Invalid ID for random provider")
	}
}
