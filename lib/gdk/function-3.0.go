// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import cairo "github.com/pekim/gobbi/lib/cairo"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Sets the specified #GdkRGBA as the source color of @cr.
/*

C function

gdk_cairo_set_source_rgba
*/
func CairoSetSourceRgba(cr *cairo.Context, rgba *RGBA) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gdk_cairo_set_source_rgba(c_cr, c_rgba)

	return
}

// Disables multidevice support in GDK. This call must happen prior
// to gdk_display_open(), gtk_init(), gtk_init_with_args() or
// gtk_init_check() in order to take effect.
//
// Most common GTK+ applications won’t ever need to call this. Only
// applications that do mixed GDK/Xlib calls could want to disable
// multidevice support if such Xlib code deals with input devices in
// any way and doesn’t observe the presence of XInput 2.
/*

C function

gdk_disable_multidevice
*/
func DisableMultidevice() {
	C.gdk_disable_multidevice()

	return
}

// Removes an error trap pushed with gdk_error_trap_push(), but
// without bothering to wait and see whether an error occurred.  If an
// error arrives later asynchronously that was triggered while the
// trap was pushed, that error will be ignored.
/*

C function

gdk_error_trap_pop_ignored
*/
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()

	return
}

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1
