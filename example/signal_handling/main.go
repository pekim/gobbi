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
	gtk.Init()

	//gobbi.SetTraceHandler(func(message string) {
	//	fmt.Print(message)
	//})

	window := gtk.WindowNew(gtk.WindowType_Toplevel)
	window.SetTitle("A window title")
	window.SetDefaultSize(400, 300)

	container := gtk.BoxNew(gtk.Orientation_Vertical, 20).Container()
	window.Container().Add(container.Widget())

	clickCount := 0
	button := gtk.ButtonNewWithLabel("Click me")
	button.Widget().SetHalign(gtk.Align_Center)
	button.ConnectClicked(func(_ *gtk.Button) {
		clickCount++
		button.SetLabel(fmt.Sprintf("clicked %d times", clickCount))
	})
	container.Add(button.Widget())

	label := gtk.LabelNew("or press keys (and check stdout in terminal)")
	container.Add(label.Widget())

	window.Widget().ConnectKeyPressEvent(func(w *gtk.Widget, event *gdk.EventKey) bool {
		message := fmt.Sprintf("key pressed : %s  %d  %d", event.FieldString(), event.FieldKeyval(), event.FieldState())
		label.SetText(message)
		fmt.Println(message)

		return gdk.EVENT_PROPAGATE
	})

	connectId2 := window.Widget().ConnectKeyPressEvent(func(_ *gtk.Widget, event *gdk.EventKey) bool {
		fmt.Println("Should not see this message.")
		return false
	})
	window.Widget().DisconnectSignal(connectId2)

	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })
	window.Widget().ShowAll()

	gtk.Main()
}
