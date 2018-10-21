package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	window.Container().Add(container.Widget())

	clickCount := 0
	button := gtk.ButtonNewWithLabel("Click me")
	button.Widget().SetHalign(gtk.GTK_ALIGN_CENTER)
	button.ConnectClicked(func() {
		clickCount++
		button.SetLabel(fmt.Sprintf("clicked %d times", clickCount))
	})
	container.Add(button.Widget())

	window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println("key pressed :", event.String, event.Keyval, event.State)
		return false
	})

	connectId2 := window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		fmt.Println("Should not see this message.")
		return false
	})
	window.Widget().DisconnectKeyPressEvent(connectId2)

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
