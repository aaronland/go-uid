package uid

import (
	"testing"
	"fmt"
)

func TestStringUID(t *testing.T) {

	u, err := NewStringProvider(12345)

	if err != nil {
		t.Fatalf("Failed to create new thing, %v", err)
	}
	
	fmt.Printf("HELLO %T", u)
}

