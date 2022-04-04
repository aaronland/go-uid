package uid

import (
	"testing"
	"fmt"
)

func TestUniqueIdentifier(t *testing.T) {

	u, err := NewUniqueIdentifier(12345)
	
	if err != nil {
		t.Fatalf("Failed to create new thing, %v", err)
	}
	
	fmt.Printf("HELLO %T", u)
}
