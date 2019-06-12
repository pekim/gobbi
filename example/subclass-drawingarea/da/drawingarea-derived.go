package da

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"math"
	"unsafe"
)

type DrawingAreaDerived struct {
	class  *DrawingAreaDerivedClass
	native DrawingAreaDerivedNative
}

func (d *DrawingAreaDerived) Draw(cr *cairo.Context) {
	widget := d.DrawingArea().Widget()

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

	styleContext := widget.GetStyleContext()
	colour := styleContext.GetColor(styleContext.GetState())

	// use the widget's context's colour for the source pattern
	gdk.CairoSetSourceRgba(cr, colour)

	// fill the path (that describes a circle)
	cr.Fill()
}

// DrawingArea upcasts to *DrawingArea
func (recv *DrawingAreaDerived) DrawingArea() *gtk.DrawingArea {
	return gtk.DrawingAreaNewFromC(unsafe.Pointer(recv.native))

}
