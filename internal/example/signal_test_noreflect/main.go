package main

//#include <glib-object.h>
import "C"

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
	window.SetDefaultSize(250, 100)
	window.Widget().ShowAll()

	connectDestroy(window.Widget(), func() {
		gtk.MainQuit()
	})

	connectKeyPressEvent(window.Widget(), func(event *gdk.EventKey) {
		fmt.Println("kp", event)
	})

	window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println("kp, for real!", event)
		return false
	})

	gtk.Main()
}
