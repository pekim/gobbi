package main

import (
	"github.com/pekim/gobbi/example/subclass-drawingarea/sc"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
)

func main() {
	gtk.Init(os.Args)

	daClass := sc.DrawingAreaDerive()
	da := daClass.New()

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Container().Add(da.DrawingArea().Widget())
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
