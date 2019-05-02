// This is a generated file - DO NOT EDIT
// +build gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetNKeys is a wrapper around the C function gdk_device_get_n_keys.
func (recv *Device) GetNKeys() int32 {
	retC := C.gdk_device_get_n_keys((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_window_get_display.
func (recv *Window) GetDisplay() *Display {
	retC := C.gdk_window_get_display((*C.GdkWindow)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHeight is a wrapper around the C function gdk_window_get_height.
func (recv *Window) GetHeight() int32 {
	retC := C.gdk_window_get_height((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetScreen is a wrapper around the C function gdk_window_get_screen.
func (recv *Window) GetScreen() *Screen {
	retC := C.gdk_window_get_screen((*C.GdkWindow)(recv.native))
	retGo := ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisual is a wrapper around the C function gdk_window_get_visual.
func (recv *Window) GetVisual() *Visual {
	retC := C.gdk_window_get_visual((*C.GdkWindow)(recv.native))
	retGo := VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gdk_window_get_width.
func (recv *Window) GetWidth() int32 {
	retC := C.gdk_window_get_width((*C.GdkWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// CairoSetSourceWindow is a wrapper around the C function gdk_cairo_set_source_window.
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
