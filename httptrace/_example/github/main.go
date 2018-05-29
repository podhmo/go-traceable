package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	client := github.NewClient(nil)
	ctx := context.Background()
	owner := "podhmo"
	repo := "go-traceable"
	result, _, err := client.Repositories.ListReleases(ctx, owner, repo, nil)
	pp.Println(result, err)
	return nil
}
