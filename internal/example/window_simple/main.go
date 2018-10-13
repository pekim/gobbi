package main

import (
	"github.com/pekim/gobbi/lib/gtk"
)

func main() {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Qaz qwerty +++")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(gtk.LabelNew("Label 1").Widget())
	container.Add(gtk.LabelNew("Label 2 ++").Widget())
	window.Container().Add(container.Widget())

	window.Widget().ShowAll()

	// Will not terminate when the window is closed.
	// Can be fixed when signals are supported.
	gtk.Main()
}
