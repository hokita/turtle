package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := turtle(); err != nil {
		if errors.Is(err, ErrMissingTarget) {
			fmt.Fprintln(os.Stderr, "Usage: turtle <target>")
		} else {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
