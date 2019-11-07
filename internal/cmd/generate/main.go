package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"os"
	"time"
)

func main() {
	generateLibraries()
	outputScc()
}

func generateLibraries() {
	noFormat := false
	if len(os.Args) >= 2 && os.Args[1] == "noformat" {
		noFormat = true
	}

	start := time.Now()

	generate.FromRoot("Atk", "1.0", noFormat)
	generate.FromRoot("cairo", "1.0", noFormat)
	generate.FromRoot("freetype2", "2.0", noFormat)
	generate.FromRoot("Gdk", "3.0", noFormat)
	generate.FromRoot("GdkPixbuf", "2.0", noFormat)
	generate.FromRoot("Gio", "2.0", noFormat)
	generate.FromRoot("GLib", "2.0", noFormat)
	generate.FromRoot("GObject", "2.0", noFormat)
	generate.FromRoot("Gtk", "3.0", noFormat)
	generate.FromRoot("GtkSource", "3.0", noFormat)
	generate.FromRoot("Pango", "1.0", noFormat)
	generate.FromRoot("PangoCairo", "1.0", noFormat)
	generate.FromRoot("PangoFT2", "1.0", noFormat)

	end := time.Now()
	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
