package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"time"
)

func main() {
	generateLibraries()
	outputScc()
}

func generateLibraries() {
	start := time.Now()
	generate.FromRoot("Gtk", "3.0")
	//generate.FromRoot("PangoCairo", "1.0")
	end := time.Now()

	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
