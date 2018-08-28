package main

import (
	"gobbi2/internal/generate"
)

func main() {
	gir := generate.GirNew("Gtk-3.0")
	gir.Generate()
}
