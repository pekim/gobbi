// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	"sync"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	cairo_surface_t * window_createSurfaceHandler(GObject *, gint, gint, gpointer);

	static gulong Window_signal_connect_create_surface(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-surface", G_CALLBACK(window_createSurfaceHandler), data);
	}

*/
import "C"

// Blacklisted : gdk_device_get_associated_device

// Blacklisted : gdk_device_get_axis_value

// Blacklisted : gdk_device_get_device_type

// Blacklisted : gdk_device_get_display

// Blacklisted : gdk_device_get_n_axes

// Blacklisted : gdk_device_get_position

// Blacklisted : gdk_device_get_window_at_position

// Blacklisted : gdk_device_get_window_at_position_double

// Blacklisted : gdk_device_grab

// Blacklisted : gdk_device_list_axes

// Blacklisted : gdk_device_ungrab

// Blacklisted : gdk_device_warp

// Blacklisted : gdk_device_manager_get_client_pointer

// Blacklisted : gdk_device_manager_get_display

// Blacklisted : gdk_device_manager_list_devices

// Blacklisted : gdk_display_get_app_launch_context

// Blacklisted : gdk_display_get_device_manager

// Blacklisted : gdk_display_has_pending

// Blacklisted : gdk_display_notify_startup_complete

// Blacklisted : gdk_display_manager_open_display

// Blacklisted : gdk_drag_context_get_dest_window

// Blacklisted : gdk_drag_context_get_protocol

// Blacklisted : gdk_keymap_get_num_lock_state

type signalWindowCreateSurfaceDetail struct {
	callback  WindowSignalCreateSurfaceCallback
	handlerID C.gulong
}

var signalWindowCreateSurfaceId int
var signalWindowCreateSurfaceMap = make(map[int]signalWindowCreateSurfaceDetail)
var signalWindowCreateSurfaceLock sync.RWMutex

// WindowSignalCreateSurfaceCallback is a callback function for a 'create-surface' signal emitted from a Window.
type WindowSignalCreateSurfaceCallback func(width int32, height int32) cairo.Surface

/*
ConnectCreateSurface connects the callback to the 'create-surface' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectCreateSurface to remove it.
*/
func (recv *Window) ConnectCreateSurface(callback WindowSignalCreateSurfaceCallback) int {
	signalWindowCreateSurfaceLock.Lock()
	defer signalWindowCreateSurfaceLock.Unlock()

	signalWindowCreateSurfaceId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_create_surface(instance, C.gpointer(uintptr(signalWindowCreateSurfaceId)))

	detail := signalWindowCreateSurfaceDetail{callback, handlerID}
	signalWindowCreateSurfaceMap[signalWindowCreateSurfaceId] = detail

	return signalWindowCreateSurfaceId
}

/*
DisconnectCreateSurface disconnects a callback from the 'create-surface' signal for the Window.

The connectionID should be a value returned from a call to ConnectCreateSurface.
*/
func (recv *Window) DisconnectCreateSurface(connectionID int) {
	signalWindowCreateSurfaceLock.Lock()
	defer signalWindowCreateSurfaceLock.Unlock()

	detail, exists := signalWindowCreateSurfaceMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowCreateSurfaceMap, connectionID)
}

//export window_createSurfaceHandler
func window_createSurfaceHandler(_ *C.GObject, c_width C.gint, c_height C.gint, data C.gpointer) *C.cairo_surface_t {
	signalWindowCreateSurfaceLock.RLock()
	defer signalWindowCreateSurfaceLock.RUnlock()

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalWindowCreateSurfaceMap[index].callback
	retGo := callback(width, height)
	retC :=
		(*C.cairo_surface_t)(retGo.ToC())
	return retC
}

// Blacklisted : gdk_window_get_device_cursor

// Blacklisted : gdk_window_get_device_events

// Unsupported : gdk_window_get_device_position : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Blacklisted : gdk_window_get_drag_protocol

// Blacklisted : gdk_window_get_support_multidevice

// Blacklisted : gdk_window_set_device_cursor

// Blacklisted : gdk_window_set_device_events

// Blacklisted : gdk_window_set_source_events

// Blacklisted : gdk_window_set_support_multidevice
