package main

import (
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(gtk.LabelNew("Label 1").Widget())
	container.Add(gtk.LabelNew("Label 2 ++").Widget())
	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
