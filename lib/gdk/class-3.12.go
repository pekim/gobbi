// This is a generated file - DO NOT EDIT
// +build gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Gets information about which window the given pointer device is in, based on events
// that have been received so far from the display server. If another application
// has a pointer grab, or this application has a grab with owner_events = %FALSE,
// %NULL may be returned even if the pointer is physically over one of this
// application's windows.
/*

C function

gdk_device_get_last_event_window
*/
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

// Get the current event compression setting for this window.
/*

C function

gdk_window_get_event_compression
*/
func (recv *Window) GetEventCompression() bool {
	retC := C.gdk_window_get_event_compression((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines whether or not extra unprocessed motion events in
// the event queue can be discarded. If %TRUE only the most recent
// event will be delivered.
//
// Some types of applications, e.g. paint programs, need to see all
// motion events and will benefit from turning off event compression.
//
// By default, event compression is enabled.
/*

C function

gdk_window_set_event_compression
*/
func (recv *Window) SetEventCompression(eventCompression bool) {
	c_event_compression :=
		boolToGboolean(eventCompression)

	C.gdk_window_set_event_compression((*C.GdkWindow)(recv.native), c_event_compression)

	return
}

// Newer GTK+ windows using client-side decorations use extra geometry
// around their frames for effects like shadows and invisible borders.
// Window managers that want to maximize windows or snap to edges need
// to know where the extents of the actual frame lie, so that users
// don’t feel like windows are snapping against random invisible edges.
//
// Note that this property is automatically updated by GTK+, so this
// function should only be used by applications which do not use GTK+
// to create toplevel windows.
/*

C function

gdk_window_set_shadow_width
*/
func (recv *Window) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	c_left := (C.gint)(left)

	c_right := (C.gint)(right)

	c_top := (C.gint)(top)

	c_bottom := (C.gint)(bottom)

	C.gdk_window_set_shadow_width((*C.GdkWindow)(recv.native), c_left, c_right, c_top, c_bottom)

	return
}
