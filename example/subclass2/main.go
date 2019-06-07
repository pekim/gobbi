package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
)

type MyWidget struct {
	instance *gtk.EntryDerived
}

func (w *MyWidget) Draw(cr *cairo.Context) bool {
	fmt.Println("draw!")
	return true
}

func main() {
	gtk.Init(os.Args)

	//sc.SubclassCreate()
	//gobject.SubclassCreate()
	//gobject.FfiClosure()

	class := gtk.EntryDerive("test_widget")
	fmt.Println(class)

	myWidget := MyWidget{}
	myWidget.instance = class.New(&myWidget)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Container().Add(myWidget.instance.Entry().Widget())
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
