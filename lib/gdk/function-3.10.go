// This is a generated file - DO NOT EDIT
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

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

// CairoSurfaceCreateFromPixbuf is a wrapper around the C function gdk_cairo_surface_create_from_pixbuf.
func CairoSurfaceCreateFromPixbuf(pixbuf *gdkpixbuf.Pixbuf, scale int32, forWindow *Window) *cairo.Surface {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	c_scale := (C.int)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	retC := C.gdk_cairo_surface_create_from_pixbuf(c_pixbuf, c_scale, c_for_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAllowedBackends is a wrapper around the C function gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	c_backends := C.CString(backends)
	defer C.free(unsafe.Pointer(c_backends))

	C.gdk_set_allowed_backends(c_backends)

	return
}
