package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const version = "0.0.3"

func main() {
	versionFlag := flag.Bool("version", false, "Display the version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("turtle version %s\n", version)
		os.Exit(0)
	}

	if err := turtle(); err != nil {
		if errors.Is(err, ErrMissingTarget) {
			fmt.Fprintln(os.Stderr, "Usage: turtle <target>")
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
