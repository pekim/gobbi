package main

import (
	"gobbi2/internal/generate"
)

func main() {
	gir := generate.GirNewRoot("GObject", "2.0")
	gir.Generate()
}
