package main

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"math"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

const css = `
	.custom {
		color: red;
		background-color: green;

		border-radius: 16px;
	}

	.custom.over {
		color: green;
		background-color: red;
	}
`

const class = "custom"
const hoverClass = "over"

func main() {
	gtk.Init(os.Args)

	// create the custom widget
	da := gtk.DrawingAreaNew()
	da.Widget().SetHalign(gtk.GTK_ALIGN_CENTER)
	da.Widget().SetSizeRequest(150, 150)

	// style the widget
	cssProvider := gtk.CssProviderNew()
	_, err := cssProvider.LoadFromData(([]uint8)(css))
	if err != nil {
		panic(err)
	}
	da.Widget().GetStyleContext().AddProvider(cssProvider.StyleProvider(), uint32(gtk.STYLE_PROVIDER_PRIORITY_APPLICATION))
	da.Widget().GetStyleContext().AddClass(class)

	// draw widget when required
	da.Widget().ConnectDraw(func(cr *cairo.Context) bool {
		return draw(cr, da.Widget())
	})

	// add a class on mouse enter, and remove the class on mouse leave
	da.Widget().AddEvents(int32(gdk.GDK_ENTER_NOTIFY_MASK | gdk.GDK_LEAVE_NOTIFY_MASK))
	da.Widget().ConnectEnterNotifyEvent(func(event *gdk.EventCrossing) bool {
		da.Widget().GetStyleContext().AddClass(hoverClass)
		return false
	})
	da.Widget().ConnectLeaveNotifyEvent(func(event *gdk.EventCrossing) bool {
		da.Widget().GetStyleContext().RemoveClass(hoverClass)
		return false
	})

	// create window
	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Custom drawing")
	window.SetDefaultSize(300, 300)

	// create a container - just used to centre the custome widget
	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Widget().SetValign(gtk.GTK_ALIGN_CENTER)

	// add the container and custom drawn widget to the window
	container.Add(da.Widget())
	window.Container().Add(container.Widget())

	// show the window until it's destroyed
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()
	gtk.Main()
}

func draw(cr *cairo.Context, widget *gtk.Widget) bool {
	// find the dimensions that the widget's  been allocated, and
	// therefore the size of the area to draw in
	alloc := widget.GetAllocation()
	height := float64(alloc.Height)
	width := float64(alloc.Width)

	// render background first
	gtk.RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	// an arc that describes a circle to the path
	cr.Arc(width/2.0, height/2.0, math.Min(width, height)/2.0, 0, 2*math.Pi)

	// use the widget's context's colour for the source pattern
	styleContext := widget.GetStyleContext()
	colour := styleContext.GetColor(styleContext.GetState())
	gdk.CairoSetSourceRgba(cr, colour)

	// fill the path (that describes a circle)
	cr.Fill()

	return false
}
