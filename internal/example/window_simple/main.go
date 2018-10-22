package main

import (
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

/*
A very simple example, that creates a window, sets its title,
and shows it.
The application quits when the window is closed.
*/

func init() {
	// Ensure that the ui's main thread is locked to the main thread.
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
