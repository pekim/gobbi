package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"time"
)

var libraries = generate.RepositorySpecs{
	{Version: "1.0", Name: "Atk"},
	//{Version: "1.0", Name: "cairo"},
	//{Version: "2.0", Name: "freetype2"},
	{Version: "3.0", Name: "Gdk"},
	{Version: "2.0", Name: "GdkPixbuf"},
	{Version: "2.0", Name: "Gio"},
	{Version: "2.0", Name: "GLib"},
	{Version: "2.0", Name: "GObject"},
	{Version: "3.0", Name: "Gtk"},
	//{Version: "3.0", Name: "GtkSource"},
	//{Version: "4.0", Name: "JavaScriptCore"},
	{Version: "1.0", Name: "Pango"},
	//{Version: "1.0", Name: "PangoCairo"},
	//{Version: "1.0", Name: "PangoFT2"},
	//{Version: "2.4", Name: "Soup"},
	//{Version: "4.0", Name: "WebKit2"},
	{Version: "2.0", Name: "xlib"},
}

func generateLibraries() {
	start := time.Now()

	libraries.Generate()

	end := time.Now()
	fmt.Printf("generation %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
