package main

import (
	"fmt"
	"github.com/rkoesters/xdg/desktop"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v <desktop file> <uris>\n", os.Args[0])
		os.Exit(1)
	}

	entry, err := desktop.NewFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening '%v': %v\n", os.Args[1], err)
		os.Exit(1)
	}

	var uris []string
	if len(os.Args) > 2 {
		uris = os.Args[2:]
	}

	err = entry.Launch(uris...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error launching '%v': %v\n", os.Args[1], err)
		os.Exit(1)
	}
}
