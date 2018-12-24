// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CairoGetDrawingContext is a wrapper around the C function gdk_cairo_get_drawing_context.
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

// PangoContextGetForDisplay is a wrapper around the C function gdk_pango_context_get_for_display.
func PangoContextGetForDisplay(display *Display) *pango.Context {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	retC := C.gdk_pango_context_get_for_display(c_display)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}
