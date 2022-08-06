package uid

import (
	"context"
	"fmt"
	"testing"
)

func TestStringProvider(t *testing.T) {

	ctx := context.Background()

	str := "helloworld"

	uri := fmt.Sprintf("%s://?string=%s", STRING_SCHEME, str)

	pr, err := NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create string UID, %v", err)
	}

	if id.String() != str {
		t.Fatalf("Invalid ID for string provider")
	}
}
