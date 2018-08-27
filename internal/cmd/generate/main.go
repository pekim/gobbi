package main

import (
	"gobbi2/internal/generate"
)

func main() {
	gir := generate.GirNew("GLib-2.0")
	gir.Generate()
}
