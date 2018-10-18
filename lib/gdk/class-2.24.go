// This is a generated file - DO NOT EDIT
// +build gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetNKeys is a wrapper around the C function gdk_device_get_n_keys.
func (recv *Device) GetNKeys() int32 {
	retC := C.gdk_device_get_n_keys((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported signal 'create-surface' for Window : return value cairo.Surface :

// Unsupported signal 'pick-embedded-child' for Window : return value Window :

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
