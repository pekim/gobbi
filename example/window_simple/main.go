package main

import (
	"github.com/pekim/gobbi"
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

	gtk.Init()

	window := gtk.WindowNew(gtk.WindowType_Toplevel)
	window.SetTitle("qaz")
	window.SetDefaultSize(400, 300)
	window.ShowAll()

	gtk.Main()
}
