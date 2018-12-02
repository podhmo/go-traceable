package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
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
	result, _, err := client.Repositories.ListContributors(ctx, owner, repo, nil)
	if err != nil {
		return err
	}
	_ = result
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.SetIndent("", "  ")
	// if err := encoder.Encode(result); err != nil {
	//     return err
	// }
	return nil
}
