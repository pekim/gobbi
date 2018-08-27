package main

import (
	"gobbi2/internal/generate"
)

func main() {
	gir := generate.GirNew()
	gir.Generate()
}
