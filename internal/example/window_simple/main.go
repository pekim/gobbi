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

	connectId1 := window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println(1, event.Keyval, event.String)
		return false
	})

	connectId2 := window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println(2, event.Keyval, event.String)
		return false
	})

	fmt.Println(connectId1, connectId2)
	window.Widget().DisconnectKeyPressEvent(connectId2)

	window.Widget().ShowAll()
	gtk.Main()
}
