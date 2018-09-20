package main

import (
	"gobbi2/internal/generate"
)

func main() {
	gir := generate.GirNew("GObject-2.0")
	gir.Generate()
}
