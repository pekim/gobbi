package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"time"
)

func main() {
	start := time.Now()
	generate.FromRoot("Gtk", "3.0")
	fmt.Printf("\ngeneration %.0fms\n\n", time.Now().Sub(start).Seconds()*1000)

}
