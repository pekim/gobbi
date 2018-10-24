// This is a generated file - DO NOT EDIT
// +build gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetLastEventWindow is a wrapper around the C function gdk_device_get_last_event_window.
func (recv *Device) GetLastEventWindow() *Window {
	retC := C.gdk_device_get_last_event_window((*C.GdkDevice)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetEventCompression is a wrapper around the C function gdk_window_get_event_compression.
func (recv *Window) GetEventCompression() bool {
	retC := C.gdk_window_get_event_compression((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetEventCompression is a wrapper around the C function gdk_window_set_event_compression.
func (recv *Window) SetEventCompression(eventCompression bool) {
	c_event_compression :=
		boolToGboolean(eventCompression)

	C.gdk_window_set_event_compression((*C.GdkWindow)(recv.native), c_event_compression)

	return
}

// SetShadowWidth is a wrapper around the C function gdk_window_set_shadow_width.
func (recv *Window) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	c_left := (C.gint)(left)

	c_right := (C.gint)(right)

	c_top := (C.gint)(top)

	c_bottom := (C.gint)(bottom)

	C.gdk_window_set_shadow_width((*C.GdkWindow)(recv.native), c_left, c_right, c_top, c_bottom)

	return
}
