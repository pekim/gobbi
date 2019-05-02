// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"C"
	cairo "github.com/pekim/gobbi/lib/cairo"
)

// CairoDrawFromGl is a wrapper around the C function gdk_cairo_draw_from_gl.
func CairoDrawFromGl(cr *cairo.Context, window *Window, source int32, sourceType int32, bufferScale int32, x int32, y int32, width int32, height int32) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_source := (C.int)(source)

	c_source_type := (C.int)(sourceType)

	c_buffer_scale := (C.int)(bufferScale)

	c_x := (C.int)(x)

	c_y := (C.int)(y)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_cairo_draw_from_gl(c_cr, c_window, c_source, c_source_type, c_buffer_scale, c_x, c_y, c_width, c_height)

	return
}
