// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
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

	void Window_pickEmbeddedChildHandler();

	static gulong Window_signal_connect_pick_embedded_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pick-embedded-child", Window_pickEmbeddedChildHandler, data);
	}

*/
/*

	void Window_toEmbedderHandler();

	static gulong Window_signal_connect_to_embedder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "to-embedder", Window_toEmbedderHandler, data);
	}

*/
import "C"

var signalWindowFromEmbedderId int
var signalWindowFromEmbedderMap = make(map[int]WindowSignalFromEmbedderCallback)
var signalWindowFromEmbedderLock sync.Mutex

// WindowSignalFromEmbedderCallback is a callback function for a 'from-embedder' signal emitted from a Window.
type WindowSignalFromEmbedderCallback func(embedderX float64, embedderY float64)

func (recv *Window) ConnectFromEmbedder(callback WindowSignalFromEmbedderCallback) {
	signalWindowFromEmbedderLock.Lock()
	defer signalWindowFromEmbedderLock.Unlock()

	signalWindowFromEmbedderId++
	signalWindowFromEmbedderMap[signalWindowFromEmbedderId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Window_signal_connect_from_embedder(instance, C.gpointer(uintptr(signalWindowFromEmbedderId)))
}

//export Window_fromEmbedderHandler
func Window_fromEmbedderHandler() {}

var signalWindowPickEmbeddedChildId int
var signalWindowPickEmbeddedChildMap = make(map[int]WindowSignalPickEmbeddedChildCallback)
var signalWindowPickEmbeddedChildLock sync.Mutex

// WindowSignalPickEmbeddedChildCallback is a callback function for a 'pick-embedded-child' signal emitted from a Window.
type WindowSignalPickEmbeddedChildCallback func(x float64, y float64) Window

func (recv *Window) ConnectPickEmbeddedChild(callback WindowSignalPickEmbeddedChildCallback) {
	signalWindowPickEmbeddedChildLock.Lock()
	defer signalWindowPickEmbeddedChildLock.Unlock()

	signalWindowPickEmbeddedChildId++
	signalWindowPickEmbeddedChildMap[signalWindowPickEmbeddedChildId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Window_signal_connect_pick_embedded_child(instance, C.gpointer(uintptr(signalWindowPickEmbeddedChildId)))
}

//export Window_pickEmbeddedChildHandler
func Window_pickEmbeddedChildHandler() {}

var signalWindowToEmbedderId int
var signalWindowToEmbedderMap = make(map[int]WindowSignalToEmbedderCallback)
var signalWindowToEmbedderLock sync.Mutex

// WindowSignalToEmbedderCallback is a callback function for a 'to-embedder' signal emitted from a Window.
type WindowSignalToEmbedderCallback func(offscreenX float64, offscreenY float64)

func (recv *Window) ConnectToEmbedder(callback WindowSignalToEmbedderCallback) {
	signalWindowToEmbedderLock.Lock()
	defer signalWindowToEmbedderLock.Unlock()

	signalWindowToEmbedderId++
	signalWindowToEmbedderMap[signalWindowToEmbedderId] = callback

	instance := C.gpointer(recv.Object().ToC())
	C.Window_signal_connect_to_embedder(instance, C.gpointer(uintptr(signalWindowToEmbedderId)))
}

//export Window_toEmbedderHandler
func Window_toEmbedderHandler() {}

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
