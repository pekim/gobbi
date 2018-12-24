// This is a generated file - DO NOT EDIT
// +build gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CairoCreate is a wrapper around the C function gdk_cairo_create.
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

// CairoRegion is a wrapper around the C function gdk_cairo_region.
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

// CairoSetSourceColor is a wrapper around the C function gdk_cairo_set_source_color.
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

// CairoSetSourcePixbuf is a wrapper around the C function gdk_cairo_set_source_pixbuf.
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
