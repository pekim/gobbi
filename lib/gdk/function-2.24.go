// This is a generated file - DO NOT EDIT
// +build gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import cairo "github.com/pekim/gobbi/lib/cairo"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Sets the given window as the source pattern for @cr.
//
// The pattern has an extend mode of %CAIRO_EXTEND_NONE and is aligned
// so that the origin of @window is @x, @y. The window contains all its
// subwindows when rendering.
//
// Note that the contents of @window are undefined outside of the
// visible part of @window, so use this function with care.
/*

C function : gdk_cairo_set_source_window
*/
func CairoSetSourceWindow(cr *cairo.Context, window *Window, x float64, y float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	C.gdk_cairo_set_source_window(c_cr, c_window, c_x, c_y)

	return
}
