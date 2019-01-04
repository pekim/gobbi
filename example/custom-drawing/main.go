package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
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
	window.SetTitle("Custom drawing")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Widget().SetValign(gtk.GTK_ALIGN_CENTER)
	window.Container().Add(container.Widget())

	label := gtk.LabelNew("custom drawing")
	container.Add(label.Widget())

	da := gtk.DrawingAreaNew()
	//bgColour := gdk.RGBA{Red: 1.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	//da.Widget().OverrideBackgroundColor(gtk.GTK_STATE_FLAG_NORMAL, &bgColour)
	da.Widget().SetHalign(gtk.GTK_ALIGN_CENTER)
	da.Widget().SetSizeRequest(50, 50)
	da.Widget().ConnectDraw(func(cr *cairo.Context) bool {
		alloc := da.Widget().GetAllocation()
		fmt.Println(alloc.Height, alloc.Width)

		gtk.RenderBackground(da.Widget().GetStyleContext(), cr,
			0, 0, float64(alloc.Width), float64(alloc.Height))

		return false
	})
	container.Add(da.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
