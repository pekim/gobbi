package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"os"
	"time"
)

func main() {
	generateLibraries()
	outputScc()
}

func generateLibraries() {
	noFormat := false
	if len(os.Args) >= 2 && os.Args[1] == "noformat" {
		noFormat = true
	}

	start := time.Now()
	generate.FromRoot("Gtk", "3.0", noFormat)
	generate.FromRoot("PangoCairo", "1.0", noFormat)
	end := time.Now()

	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
