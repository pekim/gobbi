package main

import (
	"github.com/pekim/gobbi/internal/generate"
)

func main() {
	girs := generate.GirNewRoot("Gio", "2.0")

	for _, gir := range girs {
		gir.Generate()
	}
}
