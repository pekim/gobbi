package main

import (
	"fmt"
	"github.com/boyter/scc/processor"
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
	fmt.Printf("\ngeneration %.0fms\n\n", time.Now().Sub(start).Seconds()*1000)
}

func outputScc() {
	processor.DirFilePaths = []string{"./lib/"}
	processor.PathBlacklist = ".git"
	processor.Cocomo = true
	processor.Process()
}
