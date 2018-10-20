package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Qaz qwerty +++")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(gtk.LabelNew("Label 1").Widget())
	container.Add(gtk.LabelNew("Label 2 ++").Widget())
	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(func() {
		gtk.MainQuit()
	})

	window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println(event.Keyval, event.String)
		return false
	})

	window.Widget().ShowAll()
	gtk.Main()
}
