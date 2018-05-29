package main

import (
	"log"

	"github.com/podhmo/go-traceable/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	http.Get("https://example.com")
	return nil
}
