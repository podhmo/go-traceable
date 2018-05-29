package main

import (
	"fmt"

	"github.com/podhmo/go-traceable/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
