package main

import (
	"fmt"
	"github.com/pekim/gobbi"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
)

func init() {
	// Ensure that the ui's main thread is locked to the main thread.
	runtime.LockOSThread()
}

func main() {
	gobbi.SetErrorHandler(func(err error) {
		panic(err)
	})

	gobbi.SetTraceHandler(func(message string) {
		fmt.Print(message)
	})

	gtk.Init()

	window := gtk.WindowNew(gtk.WindowType_Toplevel)
	window.SetTitle("gobbi")
	window.SetDefaultSize(400, 300)
	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })

	window.Widget().GetWindow().SetEvents(gdk.EventMask_KeyPressMask)
	window.Widget().ConnectKeyPressEvent(func(instance *gtk.Widget, event *gdk.EventKey) bool {
		fmt.Println(event.FieldKeyval(), event.FieldString())
		return true
	})

	window.Widget().ShowAll()

	gtk.Main()
}
