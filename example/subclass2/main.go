package main

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"github.com/pekim/gobbi/lib/gtk"
	"math"
	"os"
)

type MyWidget struct {
	instance *gtk.DrawingAreaDerived
}

func (w *MyWidget) Draw(cr *cairo.Context) bool {
	widget := w.instance.DrawingArea().Widget()

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

	return true
}

func main() {
	gtk.Init(os.Args)

	class := gtk.DrawingAreaDerive("test_widget", &MyWidget{})

	myWidget := MyWidget{}
	myWidget.instance = class.New(&myWidget)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	window.Container().Add(myWidget.instance.DrawingArea().Widget())
	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
