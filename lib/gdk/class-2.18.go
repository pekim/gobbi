// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"fmt"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void Window_fromEmbedderHandler();

	static gulong Window_signal_connect_from_embedder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "from-embedder", Window_fromEmbedderHandler, data);
	}

*/
/*

	void Window_toEmbedderHandler();

	static gulong Window_signal_connect_to_embedder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "to-embedder", Window_toEmbedderHandler, data);
	}

*/
import "C"

// Unsupported signal 'create-surface' for Window : return value cairo.Surface :

var signalWindowFromEmbedderId int
var signalWindowFromEmbedderMap = make(map[int]WindowSignalFromEmbedderCallback)
var signalWindowFromEmbedderLock sync.Mutex

// WindowSignalFromEmbedderCallback is a callback function for a 'from-embedder' signal emitted from a Window.
type WindowSignalFromEmbedderCallback func(embedderX float64, embedderY float64)

/*
ConnectFromEmbedder connects the callback to the 'from-embedder' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectFromEmbedder to remove it.
*/
func (recv *Window) ConnectFromEmbedder(callback WindowSignalFromEmbedderCallback) int {
	signalWindowFromEmbedderLock.Lock()
	defer signalWindowFromEmbedderLock.Unlock()

	signalWindowFromEmbedderId++
	signalWindowFromEmbedderMap[signalWindowFromEmbedderId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Window_signal_connect_from_embedder(instance, C.gpointer(uintptr(signalWindowFromEmbedderId)))
	return int(retC)
}

/*
DisconnectFromEmbedder disconnects a callback from the 'from-embedder' signal for the Window.

The connectionID should be a value returned from a call to ConnectFromEmbedder.
*/
func (recv *Window) DisconnectFromEmbedder(connectionID int) {
	_, exists := signalWindowFromEmbedderMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWindowFromEmbedderMap, connectionID)
}

//export Window_fromEmbedderHandler
func Window_fromEmbedderHandler() {
	fmt.Println("cb")
}

// Unsupported signal 'pick-embedded-child' for Window : return value Window :

var signalWindowToEmbedderId int
var signalWindowToEmbedderMap = make(map[int]WindowSignalToEmbedderCallback)
var signalWindowToEmbedderLock sync.Mutex

// WindowSignalToEmbedderCallback is a callback function for a 'to-embedder' signal emitted from a Window.
type WindowSignalToEmbedderCallback func(offscreenX float64, offscreenY float64)

/*
ConnectToEmbedder connects the callback to the 'to-embedder' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectToEmbedder to remove it.
*/
func (recv *Window) ConnectToEmbedder(callback WindowSignalToEmbedderCallback) int {
	signalWindowToEmbedderLock.Lock()
	defer signalWindowToEmbedderLock.Unlock()

	signalWindowToEmbedderId++
	signalWindowToEmbedderMap[signalWindowToEmbedderId] = callback

	instance := C.gpointer(recv.Object().ToC())
	retC := C.Window_signal_connect_to_embedder(instance, C.gpointer(uintptr(signalWindowToEmbedderId)))
	return int(retC)
}

/*
DisconnectToEmbedder disconnects a callback from the 'to-embedder' signal for the Window.

The connectionID should be a value returned from a call to ConnectToEmbedder.
*/
func (recv *Window) DisconnectToEmbedder(connectionID int) {
	_, exists := signalWindowToEmbedderMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.Object().ToC())
	C.g_signal_handler_disconnect(instance, C.gulong(connectionID))
	delete(signalWindowToEmbedderMap, connectionID)
}

//export Window_toEmbedderHandler
func Window_toEmbedderHandler() {
	fmt.Println("cb")
}

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
	retGo := CursorNewFromC(unsafe.Pointer(retC))

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
	c_sibling := (*C.GdkWindow)(sibling.ToC())

	c_above :=
		boolToGboolean(above)

	C.gdk_window_restack((*C.GdkWindow)(recv.native), c_sibling, c_above)

	return
}
