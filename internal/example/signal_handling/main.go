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
	window.SetDefaultSize(400, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 20).Container()
	window.Container().Add(container.Widget())

	clickCount := 0
	button := gtk.ButtonNewWithLabel("Click me")
	button.Widget().SetHalign(gtk.GTK_ALIGN_CENTER)
	button.ConnectClicked(func() {
		clickCount++
		button.SetLabel(fmt.Sprintf("clicked %d times", clickCount))
	})
	container.Add(button.Widget())

	label := gtk.LabelNew("or press keys (and check stdout in terminal)")
	container.Add(label.Widget())

	window.Widget().ConnectKeyPressEvent(func(event *gdk.EventKey) bool {
		message := fmt.Sprintf("key pressed : %s  %d  %d", event.String, event.Keyval, event.State)
		label.SetText(message)
		fmt.Println(message)

		return gdk.EVENT_PROPAGATE
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
