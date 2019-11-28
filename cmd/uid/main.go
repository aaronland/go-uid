package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-uid"
	"log"
)

func main() {

	uri := flag.String("uri", "", "...")

	flag.Parse()

	ctx := context.Background()
	pr, err := uid.NewProvider(ctx, *uri)

	if err != nil {
		log.Fatal(err)
	}

	id, err := pr.UID()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}
