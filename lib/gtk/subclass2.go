package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr);
void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data);
*/
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gdk"
	"math"
	"unsafe"
)

type DrawingAreaVirtualDraw interface {
	Draw(cr *cairo.Context) bool
}

//type DrawingAreaVirtualFunctions struct {
//	draw DrawingAreaDrawFunc
//}

type DrawingAreaDerivedClass struct {
	name string
	//virtualFunctions interface{}
	gtype C.GType
}

type DrawingAreaDerived struct {
	class  *DrawingAreaDerivedClass
	native *C.GtkDrawingArea
}

func DrawingAreaDerive(name string, virtualFunctions interface{}) *DrawingAreaDerivedClass {
	var vf C.GtkDrawingAreaClass
	vfWidget := (*C.GtkWidgetClass)(unsafe.Pointer(&vf))
	if _, ok := virtualFunctions.(DrawingAreaVirtualDraw); ok {
		vfWidget.draw = (*[0]byte)(C.drawing_area_vf_draw)
	}
	fmt.Println("vfw", vfWidget)

	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkDrawingAreaClass
	//typeInfo.class_data = (C.gconstpointer)(&vf)
	typeInfo.instance_size = C.sizeof_GtkDrawingArea
	typeInfo.class_init = C.GClassInitFunc(C.drawing_area_class_init)

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_DRAWING_AREA, cTypeName, &typeInfo, 0)

	class := &DrawingAreaDerivedClass{
		name: name,
		//virtualFunctions: virtualFunctions,
		gtype: gtype,
	}

	return class
}

func (c *DrawingAreaDerivedClass) New() *DrawingAreaDerived {
	//f, ok := virtualFunctions.(DrawingAreaVirtualDraw)
	//fmt.Println("draw func :", ok)
	//f.Draw(nil)

	native := (*C.GtkDrawingArea)(C.g_object_newv(c.gtype, 0, nil))

	instance := &DrawingAreaDerived{
		class:  c,
		native: native,
	}

	return instance
}

// DrawingArea upcasts to *DrawingArea
func (recv *DrawingAreaDerived) DrawingArea() *DrawingArea {
	return DrawingAreaNewFromC(unsafe.Pointer(recv.native))

}

//export DrawingAreaClassInit
func DrawingAreaClassInit(class *C.GtkDrawingAreaClass, data unsafe.Pointer) {
	widgetClass := (*C.GtkWidgetClass)(unsafe.Pointer(class))

	vf := (*C.GtkDrawingAreaClass)(data)
	vfWidget := (*C.GtkWidgetClass)(unsafe.Pointer(&vf))
	fmt.Println("vfw", vfWidget)

	if vfWidget.draw != nil {
		widgetClass.draw = vfWidget.draw
	}

	//if _, ok := virtualFunctions.(DrawingAreaVirtualDraw); ok {
	//widgetClass.draw = (*[0]byte)(C.drawing_area_vf_draw)
	//}
}

//export DrawingAreaDraw
func DrawingAreaDraw(daC *C.GtkDrawingArea, contextC *C.cairo_t) C.gboolean {
	da := DrawingAreaNewFromC(unsafe.Pointer(daC))

	widget := da.Widget()
	cr := cairo.ContextNewFromC(unsafe.Pointer(contextC))

	// find the dimensions that the widget's  been allocated, and
	// therefore the size of the area to draw in
	alloc := widget.GetAllocation()
	height := float64(alloc.Height)
	width := float64(alloc.Width)

	// render background first
	RenderBackground(widget.GetStyleContext(), cr,
		0, 0, width, height)

	// an arc that describes a circle to the path
	cr.Arc(width/2.0, height/2.0, math.Min(width, height)/2.0, 0, 2*math.Pi)

	styleContext := widget.GetStyleContext()
	colour := styleContext.GetColor(styleContext.GetState())

	// use the widget's context's colour for the source pattern
	gdk.CairoSetSourceRgba(cr, colour)

	// fill the path (that describes a circle)
	cr.Fill()

	return C.FALSE
}
