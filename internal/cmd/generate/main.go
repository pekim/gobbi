package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"time"
)

type library struct {
	version string
	name    string
}

var libraries = []library{
	//{version: "1.0", name: "Atk"},
	//{version: "1.0", name: "cairo"},
	//{version: "2.0", name: "freetype2"},
	//{version: "3.0", name: "Gdk"},
	//{version: "2.0", name: "GdkPixbuf"},
	//{version: "2.0", name: "Gio"},
	{version: "2.0", name: "GLib"},
	{version: "2.0", name: "GObject"},
	//{version: "3.0", name: "Gtk"},
	//{version: "3.0", name: "GtkSource"},
	//{version: "4.0", name: "JavaScriptCore"},
	//{version: "1.0", name: "Pango"},
	//{version: "1.0", name: "PangoCairo"},
	//{version: "1.0", name: "PangoFT2"},
	//{version: "2.4", name: "Soup"},
	//{version: "4.0", name: "WebKit2"},
}

func main() {
	start := time.Now()

	for _, lib := range libraries {
		fmt.Printf("%-16s %s\n", lib.name, lib.version)
		generate.ForRepository(lib.name, lib.version)
	}

	end := time.Now()
	fmt.Println()
	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
