package main

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

const css = `
	.custom {
		color: #c03030;
		background-color: #c0c0f0;

		border-radius: 16px;
	}

	.custom.over {
		color: #c0c0f0;
		background-color: #c03030;
	}
`

const class = "custom"
const hoverClass = "over"

func main() {
	gtk.Init()

	// create the custom widget
	da := gtk.DrawingAreaNew()
	da.Widget().SetHalign(gtk.Align_Center)
	da.Widget().SetSizeRequest(150, 150)

	// style the widget
	cssProvider := gtk.CssProviderNew()
	cssProvider.LoadFromData(css, int32(len(css)))
	da.Widget().GetStyleContext().AddProvider(cssProvider.StyleProvider(), uint32(gtk.STYLE_PROVIDER_PRIORITY_APPLICATION))
	da.Widget().GetStyleContext().AddClass(class)

	// draw widget when required
	da.Widget().ConnectDraw(func(widget *gtk.Widget, cr *cairo.Context) bool {
		return draw(cr, widget)
	})

	// add a class on mouse enter, and remove the class on mouse leave
	da.Widget().AddEvents(int32(gdk.EventMask_EnterNotifyMask | gdk.EventMask_LeaveNotifyMask))
	da.Widget().ConnectEnterNotifyEvent(func(widget *gtk.Widget, event *gdk.EventCrossing) bool {
		widget.GetStyleContext().AddClass(hoverClass)
		return false
	})
	da.Widget().ConnectLeaveNotifyEvent(func(widget *gtk.Widget, event *gdk.EventCrossing) bool {
		widget.GetStyleContext().RemoveClass(hoverClass)
		return false
	})

	// create window
	window := gtk.WindowNew(gtk.WindowType_Toplevel)
	window.SetTitle("Custom drawing")
	window.SetDefaultSize(300, 300)

	// create a container - just used to centre the custome widget
	container := gtk.BoxNew(gtk.Orientation_Vertical, 10).Container()
	container.Widget().SetValign(gtk.Align_Center)

	// add the container and custom drawn widget to the window
	container.Add(da.Widget())
	window.Container().Add(container.Widget())

	// show the window until it's destroyed
	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })
	window.Widget().ShowAll()
	gtk.Main()
}

func draw(cr *cairo.Context, widget *gtk.Widget) bool {
	// find the dimensions that the widget's  been allocated, and
	// therefore the size of the area to draw in
	//alloc := widget.GetAllocation()
	//height := float64(alloc.Height)
	//width := float64(alloc.Width)
	height := float64(widget.GetAllocatedHeight())
	width := float64(widget.GetAllocatedWidth())

	// render background first
	gtk.RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	//// an arc that describes a circle to the path
	//cr.Arc(width/2.0, height/2.0, math.Min(width, height)/2.0, 0, 2*math.Pi)
	//
	//styleContext := widget.GetStyleContext()
	//colour := styleContext.GetColor(styleContext.GetState())
	//backGroundColour := styleContext.GetBackgroundColor(styleContext.GetState())
	//
	//// use the widget's context's colour for the source pattern
	//gdk.CairoSetSourceRgba(cr, colour)
	//
	//// fill the path (that describes a circle)
	//cr.Fill()
	//
	//// Setup the font to use for text.
	//fd := pango.FontDescriptionNew()
	////fd.SetFamily("Monospace")
	////fd.SetWeight(pango.PANGO_WEIGHT_BOLD)
	//fd.SetSize(24 * pango.SCALE)
	//
	//// Layout the text.
	//layout := widget.CreatePangoLayout("Hello\nworld")
	//layout.SetFontDescription(fd)
	//
	//// Get the extents of the text.
	//textWidth, textHeight := layout.GetPixelSize()
	//// And position the text in the middle of the widget.
	//x := (width - float64(textWidth)) / 2
	//y := (height - float64(textHeight)) / 2
	//
	//// Draw the text.
	//cr.SetSourceRGBA(backGroundColour.FieldRed(), backGroundColour.FieldGreen(), backGroundColour.FieldBlue(), backGroundColour.FieldAlpha())
	//cr.MoveTo(x, y)
	//pangocairo.ShowLayout(cr, layout)

	return false
}
