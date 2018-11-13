// This is a generated file - DO NOT EDIT
// +build gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import cairo "github.com/pekim/gobbi/lib/cairo"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// This is the main way to draw GL content in GTK+. It takes a render buffer ID
// (@source_type == #GL_RENDERBUFFER) or a texture id (@source_type == #GL_TEXTURE)
// and draws it onto @cr with an OVER operation, respecting the current clip.
// The top left corner of the rectangle specified by @x, @y, @width and @height
// will be drawn at the current (0,0) position of the cairo_t.
//
// This will work for *all* cairo_t, as long as @window is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a window with
// no special effects applied to @cr it will however use a more efficient
// approach.
//
// For #GL_RENDERBUFFER the code will always fall back to software for buffers
// with alpha components, so make sure you use #GL_TEXTURE if using alpha.
//
// Calling this may change the current GL context.
/*

C function : gdk_cairo_draw_from_gl
*/
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
