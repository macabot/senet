package main

import (
	"fmt"
	"os"

	"github.com/macabot/senet/internal/app/static"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Output directory is required as first argument.")
		os.Exit(1)
	}
	outputDir := os.Args[1]

	if err := static.GeneratePages(outputDir); err != nil {
		panic(err)
	}
}
