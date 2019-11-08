package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/pekim/gobbi/internal/generate"
	"os"
	"sync"
	"time"
)

type library struct {
	version string
	name    string
}

var libraries = []library{
	{version: "1.0", name: "Atk"},
	{version: "1.0", name: "cairo"},
	{version: "2.0", name: "freetype2"},
	{version: "3.0", name: "Gdk"},
	{version: "2.0", name: "GdkPixbuf"},
	{version: "2.0", name: "Gio"},
	{version: "2.0", name: "GLib"},
	{version: "2.0", name: "GObject"},
	{version: "3.0", name: "Gtk"},
	{version: "3.0", name: "GtkSource"},
	{version: "4.0", name: "JavaScriptCore"},
	{version: "1.0", name: "Pango"},
	{version: "1.0", name: "PangoCairo"},
	{version: "1.0", name: "PangoFT2"},
	{version: "2.4", name: "Soup"},
	{version: "4.0", name: "WebKit2"},
}

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

	for _, lib := range libraries {
		fmt.Printf("%-16s %s\n", lib.name, lib.version)
	}
	fmt.Println()

	var wg sync.WaitGroup
	for _, lib := range libraries {
		wg.Add(1)

		go func(l library) {
			generate.FromRoot(l.name, l.version, noFormat, addToTotal, oneDone)
			wg.Done()
		}(lib)
	}
	wg.Wait()

	end := time.Now()
	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}

var lock sync.Mutex
var total int
var done int
var previousLength int

func addToTotal(n int) {
	lock.Lock()
	defer lock.Unlock()

	total += n
	updateProgress()
}

func oneDone() {
	lock.Lock()
	defer lock.Unlock()

	done++
	updateProgress()
}

func updateProgress() {
	tm.MoveCursorBackward(previousLength)
	previousLength, _ = tm.Printf("%d/%d", done, total)
	tm.Flush()
}
