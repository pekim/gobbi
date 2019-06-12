package da

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gtk"
	"math"
	"unsafe"
)

// DrawingAreaDerived represent an instance of the subclassed widget.
type DrawingAreaDerived struct {
	native DrawingAreaDerivedNative

	red   float64
	green float64
	blue  float64
}

func (d *DrawingAreaDerived) init() {
	// default to a mid grey
	d.SetColour(0.5, 0.5, 0.5)
}

func (d *DrawingAreaDerived) SetColour(r, g, b float64) {
	d.red = r
	d.green = g
	d.blue = b
}

func (d *DrawingAreaDerived) Draw(cr *cairo.Context) {
	widget := d.DrawingArea().Widget()

	// find the dimensions that the widget's been allocated, and
	// therefore the size of the area to draw in
	alloc := widget.GetAllocation()
	height := float64(alloc.Height)
	width := float64(alloc.Width)

	// render background first
	gtk.RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	// an arc that describes a circle to the path
	cr.Arc(width/2.0, height/2.0, math.Min(width, height)/2.0, 0, 2*math.Pi)

	// the circle's fill colour
	cr.SetSourceRGB(d.red, d.green, d.blue)

	// fill the path (that describes a circle)
	cr.Fill()
}

// DrawingArea upcasts to *DrawingArea
func (recv *DrawingAreaDerived) DrawingArea() *gtk.DrawingArea {
	return gtk.DrawingAreaNewFromC(unsafe.Pointer(recv.native))
}
