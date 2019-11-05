package main

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/glib"
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/pango"
	"github.com/pekim/gobbi/lib/pangocairo"
	"math"
	"os"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

const css = `
	.custom {
		color: #c03030;
		background-color: #c0c0f0;
	}
`

const class = "custom"

func main() {
	gtk.Init(os.Args)

	// create the custom widget
	da := gtk.DrawingAreaNew()
	da.Widget().SetHalign(gtk.GTK_ALIGN_FILL)
	da.Widget().SetValign(gtk.GTK_ALIGN_FILL)

	// style the widget
	cssProvider := gtk.CssProviderNew()
	_, err := cssProvider.LoadFromData(([]uint8)(css))
	if err != nil {
		panic(err)
	}
	da.Widget().GetStyleContext().AddProvider(cssProvider.StyleProvider(), uint32(gtk.STYLE_PROVIDER_PRIORITY_APPLICATION))
	da.Widget().GetStyleContext().AddClass(class)

	// draw widget when required
	angle := 0.0
	da.Widget().ConnectDraw(func(widget *gtk.Widget, cr *cairo.Context) bool {
		return draw(cr, widget, angle)
	})

	// create window
	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Custom drawing")
	window.SetDefaultSize(300, 300)

	window.Container().Add(da.Widget())

	// show the window until it's destroyed
	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })
	window.Widget().ShowAll()

	go func() {
		ticker := time.NewTicker(time.Second / 60)

		for {
			<-ticker.C

			angle += 2 * math.Pi / 60 / 3

			glib.IdleAddOnce(func() {
				da.Widget().QueueDraw()
			})
		}
	}()

	gtk.Main()
}

func draw(cr *cairo.Context, widget *gtk.Widget, angle float64) bool {
	// find the dimensions that the widget's  been allocated, and
	// therefore the size of the area to draw in
	alloc := widget.GetAllocation()
	height := float64(alloc.Height)
	width := float64(alloc.Width)

	// render background first
	gtk.RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	// Setup the font to use for text.
	fd := pango.FontDescriptionNew()
	fd.SetSize(24 * pango.SCALE)

	// Layout the text.
	layout := widget.CreatePangoLayout("Hello\nworld")
	layout.SetFontDescription(fd)

	// Get the extents of the text.
	textWidth, textHeight := layout.GetPixelSize()

	x := (width) / 2
	y := (height) / 2

	cr.Translate(x, y)
	cr.Rotate(angle)

	// Draw the text.
	cr.MoveTo(0-(float64(textWidth)/2), 0-(float64(textHeight)/2))
	pangocairo.ShowLayout(cr, layout)

	return false
}
