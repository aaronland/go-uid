package uid

import (
	"context"
	"testing"
)

func TestMultiUID(t *testing.T) {

	ctx := context.Background()

	str_id, err := NewStringUID(ctx, "testing")

	if err != nil {
		t.Fatalf("Failed to create string UID, %v", err)
	}

	int_id, err := NewInt64UID(ctx, int64(10))

	if err != nil {
		t.Fatalf("Failed to create int64 UID, %v", err)
	}

	m := NewMultiUID(ctx, str_id, int_id)

	expected := "StringUID#testing Int64UID#10"

	if m.String() != expected {
		t.Fatalf("Unexpected string value, %v", m.String())
	}

}
