package main

import (
	"github.com/pekim/gobbi/internal/generate"
	"os"
)

func main() {
	name := os.Args[1]
	version := os.Args[2]

	noFormat := false
	if len(os.Args) > 3 && os.Args[3] == "noformat" {
		noFormat = true
	}

	generate.FromRoot(name, version, noFormat)
}
