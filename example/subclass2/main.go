package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
)

type MyWidget struct {
	instance *gtk.DrawingAreaDerived
}

func (w *MyWidget) Draw(cr *cairo.Context) bool {
	fmt.Println("draw!")
	return true
}

func main() {
	gtk.Init(os.Args)

	class := gtk.DrawingAreaDerive("test_widget", &MyWidget{})
	fmt.Println(class)

	myWidget := MyWidget{}
	myWidget.instance = class.New()

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Container().Add(myWidget.instance.DrawingArea().Widget())
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
