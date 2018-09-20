package main

import (
	"gobbi2/internal/generate"
)

func main() {
	girs := generate.GirNewRoot("GObject", "2.0")

	for _, gir := range girs {
		gir.Generate()
	}
}
