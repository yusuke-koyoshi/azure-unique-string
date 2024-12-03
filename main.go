package main

import (
	"fmt"
	"os"

	"github.com/yusuke-koyoshi/go-unique-string"
)

func main() {
	if len(os.Args) == 1 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s string1 [string2]...\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("%s\n", unique_string.GenerateUniqueString(os.Args[1:]...))
}
