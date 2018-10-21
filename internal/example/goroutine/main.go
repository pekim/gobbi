package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/glib"
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
	"time"
)

func init() {
	// Ensure that the ui's main thread is locked to the main thread.
	runtime.LockOSThread()
}

func main() {
	gtk.Init([]string{})

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	counter := 0
	label := gtk.LabelNew("")
	window.Container().Add(label.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for range ticker.C {
			counter++

			// It is important to only update the ui on the main thread.
			// Use IdleAdd to do this.
			glib.IdleAdd(func() bool {
				label.SetText(fmt.Sprintf("count : %d", counter))
				return false
			})
		}
	}()

	gtk.Main()
}