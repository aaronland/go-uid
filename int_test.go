package uid

import (
	"context"
	"strconv"
	"testing"
)

func TestInt64UID(t *testing.T) {

	ctx := context.Background()

	i := int64(10)

	id, err := NewInt64UID(ctx, i)

	if err != nil {
		t.Fatalf("Failed to create int64 UID, %v", err)
	}

	if id.String() != strconv.FormatInt(i, 10) {
		t.Fatalf("Invalid ID for int64 UID")
	}
}
