// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Creates a Cairo context for drawing to @window.
//
// Note that calling cairo_reset_clip() on the resulting #cairo_t will
// produce undefined results, so avoid it at all costs.
//
// Typically, this function is used to draw on a #GdkWindow out of the paint
// cycle of the toolkit; this should be avoided, as it breaks various assumptions
// and optimizations.
//
// If you are drawing on a native #GdkWindow in response to a %GDK_EXPOSE event
// you should use gdk_window_begin_draw_frame() and gdk_drawing_context_get_cairo_context()
// instead. GTK will automatically do this for you when drawing a widget.
/*

C function

gdk_cairo_create
*/
func CairoCreate(window *Window) *cairo.Context {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gdk_cairo_create(c_window)
	retGo := cairo.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_cairo_rectangle : unsupported parameter rectangle : Blacklisted record : GdkRectangle

// Adds the given region to the current path of @cr.
/*

C function

gdk_cairo_region
*/
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gdk_cairo_region(c_cr, c_region)

	return
}

// Sets the specified #GdkColor as the source color of @cr.
/*

C function

gdk_cairo_set_source_color
*/
func CairoSetSourceColor(cr *cairo.Context, color *Color) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gdk_cairo_set_source_color(c_cr, c_color)

	return
}

// Sets the given pixbuf as the source pattern for @cr.
//
// The pattern has an extend mode of %CAIRO_EXTEND_NONE and is aligned
// so that the origin of @pixbuf is @pixbuf_x, @pixbuf_y.
/*

C function

gdk_cairo_set_source_pixbuf
*/
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_pixbuf_x := (C.gdouble)(pixbufX)

	c_pixbuf_y := (C.gdouble)(pixbufY)

	C.gdk_cairo_set_source_pixbuf(c_cr, c_pixbuf, c_pixbuf_x, c_pixbuf_y)

	return
}
