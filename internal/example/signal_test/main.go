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

	connectDestroy(window.Object(), func() {
		gtk.MainQuit()
	})

	connectKeyPressEvent(window.Object(), func(event *gdk.EventKey) {
		fmt.Println("kp", event)
	})

	gtk.Main()
}
