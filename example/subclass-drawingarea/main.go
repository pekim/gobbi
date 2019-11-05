package main

import (
	"github.com/pekim/gobbi/example/subclass-drawingarea/da"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
)

func main() {
	gtk.Init(os.Args)

	// register the subclass
	daClass := da.DrawingAreaDerive()

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	hbox := gtk.BoxNew(gtk.GTK_ORIENTATION_HORIZONTAL, 10)

	da1 := daClass.New()
	hbox.PackStart(da1.DrawingArea().Widget(), true, true, 0)

	da2 := daClass.New()
	da2.SetColour(0.7, 0.2, 0.2)
	hbox.PackStart(da2.DrawingArea().Widget(), true, true, 0)

	window.Container().Add(hbox.Widget())
	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })
	window.Widget().ShowAll()

	gtk.Main()
}
