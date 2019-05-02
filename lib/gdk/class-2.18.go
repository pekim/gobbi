// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	GdkWindow * window_pickEmbeddedChildHandler(GObject *, gdouble, gdouble, gpointer);

	static gulong Window_signal_connect_pick_embedded_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pick-embedded-child", G_CALLBACK(window_pickEmbeddedChildHandler), data);
	}

*/
import "C"

// Unsupported signal 'from-embedder' for Window : unsupported parameter offscreen_x : direction is 'out'

type signalWindowPickEmbeddedChildDetail struct {
	callback  WindowSignalPickEmbeddedChildCallback
	handlerID C.gulong
}

var signalWindowPickEmbeddedChildId int
var signalWindowPickEmbeddedChildMap = make(map[int]signalWindowPickEmbeddedChildDetail)
var signalWindowPickEmbeddedChildLock sync.RWMutex

// WindowSignalPickEmbeddedChildCallback is a callback function for a 'pick-embedded-child' signal emitted from a Window.
type WindowSignalPickEmbeddedChildCallback func(x float64, y float64) Window

/*
ConnectPickEmbeddedChild connects the callback to the 'pick-embedded-child' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectPickEmbeddedChild to remove it.
*/
func (recv *Window) ConnectPickEmbeddedChild(callback WindowSignalPickEmbeddedChildCallback) int {
	signalWindowPickEmbeddedChildLock.Lock()
	defer signalWindowPickEmbeddedChildLock.Unlock()

	signalWindowPickEmbeddedChildId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_pick_embedded_child(instance, C.gpointer(uintptr(signalWindowPickEmbeddedChildId)))

	detail := signalWindowPickEmbeddedChildDetail{callback, handlerID}
	signalWindowPickEmbeddedChildMap[signalWindowPickEmbeddedChildId] = detail

	return signalWindowPickEmbeddedChildId
}

/*
DisconnectPickEmbeddedChild disconnects a callback from the 'pick-embedded-child' signal for the Window.

The connectionID should be a value returned from a call to ConnectPickEmbeddedChild.
*/
func (recv *Window) DisconnectPickEmbeddedChild(connectionID int) {
	signalWindowPickEmbeddedChildLock.Lock()
	defer signalWindowPickEmbeddedChildLock.Unlock()

	detail, exists := signalWindowPickEmbeddedChildMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowPickEmbeddedChildMap, connectionID)
}

//export window_pickEmbeddedChildHandler
func window_pickEmbeddedChildHandler(_ *C.GObject, c_x C.gdouble, c_y C.gdouble, data C.gpointer) *C.GdkWindow {
	signalWindowPickEmbeddedChildLock.RLock()
	defer signalWindowPickEmbeddedChildLock.RUnlock()

	x := float64(c_x)

	y := float64(c_y)

	index := int(uintptr(data))
	callback := signalWindowPickEmbeddedChildMap[index].callback
	retGo := callback(x, y)
	retC :=
		(*C.GdkWindow)(retGo.ToC())
	return retC
}

// Unsupported signal 'to-embedder' for Window : unsupported parameter embedder_x : direction is 'out'

// EnsureNative is a wrapper around the C function gdk_window_ensure_native.
func (recv *Window) EnsureNative() bool {
	retC := C.gdk_window_ensure_native((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Flush is a wrapper around the C function gdk_window_flush.
func (recv *Window) Flush() {
	C.gdk_window_flush((*C.GdkWindow)(recv.native))

	return
}

// GeometryChanged is a wrapper around the C function gdk_window_geometry_changed.
func (recv *Window) GeometryChanged() {
	C.gdk_window_geometry_changed((*C.GdkWindow)(recv.native))

	return
}

// GetCursor is a wrapper around the C function gdk_window_get_cursor.
func (recv *Window) GetCursor() *Cursor {
	retC := C.gdk_window_get_cursor((*C.GdkWindow)(recv.native))
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetRootCoords is a wrapper around the C function gdk_window_get_root_coords.
func (recv *Window) GetRootCoords(x int32, y int32) (int32, int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_root_x C.gint

	var c_root_y C.gint

	C.gdk_window_get_root_coords((*C.GdkWindow)(recv.native), c_x, c_y, &c_root_x, &c_root_y)

	rootX := (int32)(c_root_x)

	rootY := (int32)(c_root_y)

	return rootX, rootY
}

// IsDestroyed is a wrapper around the C function gdk_window_is_destroyed.
func (recv *Window) IsDestroyed() bool {
	retC := C.gdk_window_is_destroyed((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Restack is a wrapper around the C function gdk_window_restack.
func (recv *Window) Restack(sibling *Window, above bool) {
	c_sibling := (*C.GdkWindow)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GdkWindow)(sibling.ToC())
	}

	c_above :=
		boolToGboolean(above)

	C.gdk_window_restack((*C.GdkWindow)(recv.native), c_sibling, c_above)

	return
}
