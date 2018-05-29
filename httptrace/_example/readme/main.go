package main

import (
	"fmt"

	"github.com/podhmo/go-traceable/httptrace"
)

func main() {
	resp, err := httptrace.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
