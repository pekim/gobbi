package main

import (
	"fmt"
	"github.com/pekim/gobbi/example/subclass4/sc"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
)

type MyWidget struct {
	instance *sc.DrawingAreaDerived
}

func main() {
	gtk.Init(os.Args)

	class := sc.DrawingAreaDerive("test_widget")
	fmt.Println(class)

	myWidget := MyWidget{}
	myWidget.instance = class.New(&myWidget)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Container().Add(myWidget.instance.DrawingArea().Widget())
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
