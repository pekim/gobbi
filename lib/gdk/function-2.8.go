// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CairoCreate is a wrapper around the C function gdk_cairo_create.
func CairoCreate(window *Window) *cairo.Context {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gdk_cairo_create(c_window)
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_cairo_rectangle : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Unsupported : gdk_cairo_region : no return generator

// Unsupported : gdk_cairo_set_source_color : no return generator

// Unsupported : gdk_cairo_set_source_pixbuf : no return generator
