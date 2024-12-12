package main

import (
	"github.com/pol-cova/minimal_api/cmd/minicli"
	"os"
)

func main() {
	if err := minicli.Execute(); err != nil {
		os.Exit(1)
	}
}
