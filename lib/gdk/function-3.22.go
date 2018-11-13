// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Retrieves the #GdkDrawingContext that created the Cairo
// context @cr.
/*

C function : gdk_cairo_get_drawing_context
*/
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	retC := C.gdk_cairo_get_drawing_context(c_cr)
	var retGo (*DrawingContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DrawingContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Creates a #PangoContext for @display.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for
// the widget you intend to render text onto.
//
// The newly created context will have the default font options
// (see #cairo_font_options_t) for the display; if these options
// change it will not be updated. Using gtk_widget_get_pango_context()
// is more convenient if you want to keep a context around and track
// changes to the font rendering settings.
/*

C function : gdk_pango_context_get_for_display
*/
func PangoContextGetForDisplay(display *Display) *pango.Context {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gdk_pango_context_get_for_display(c_display)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}
