package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-uid"
	"log"
)

func main() {

	dsn := flag.String("dsn", "", "...")

	flag.Parse()

	pr, err := uid.NewUIDProviderWithDSN(*dsn)

	if err != nil {
		log.Fatal(err)
	}

	id, err := pr.UID()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}
