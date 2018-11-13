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

// Creates an image surface with the same contents as
// the pixbuf.
/*

C function

gdk_cairo_surface_create_from_pixbuf
*/
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

// Sets a list of backends that GDK should try to use.
//
// This can be be useful if your application does not
// work with certain GDK backends.
//
// By default, GDK tries all included backends.
//
// For example,
// |[<!-- language="C" -->
// gdk_set_allowed_backends ("wayland,quartz,*");
// ]|
// instructs GDK to try the Wayland backend first,
// followed by the Quartz backend, and then all
// others.
//
// If the `GDK_BACKEND` environment variable
// is set, it determines what backends are tried in what
// order, while still respecting the set of allowed backends
// that are specified by this function.
//
// The possible backend names are x11, win32, quartz,
// broadway, wayland. You can also include a * in the
// list to try all remaining backends.
//
// This call must happen prior to gdk_display_open(),
// gtk_init(), gtk_init_with_args() or gtk_init_check()
// in order to take effect.
/*

C function

gdk_set_allowed_backends
*/
func SetAllowedBackends(backends string) {
	c_backends := C.CString(backends)
	defer C.free(unsafe.Pointer(c_backends))

	C.gdk_set_allowed_backends(c_backends)

	return
}
