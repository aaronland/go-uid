package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-uid"
	"log"
)

func main() {

	uri := flag.String("uri", "", "A valid aaronland/go-uid URI.")

	flag.Parse()

	ctx := context.Background()
	pr, err := uid.NewProvider(ctx, *uri)

	if err != nil {
		log.Fatalf("Failed to create provider, %v", err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		log.Fatalf("Failed to create UID, %v", err)
	}

	fmt.Println(id.String())
}
