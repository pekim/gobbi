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

	void window_fromEmbedderHandler(GObject *, gdouble, gdouble, gpointer, gpointer, gpointer);

	static gulong Window_signal_connect_from_embedder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "from-embedder", G_CALLBACK(window_fromEmbedderHandler), data);
	}

*/
/*

	GdkWindow * window_pickEmbeddedChildHandler(GObject *, gdouble, gdouble, gpointer);

	static gulong Window_signal_connect_pick_embedded_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pick-embedded-child", G_CALLBACK(window_pickEmbeddedChildHandler), data);
	}

*/
/*

	void window_toEmbedderHandler(GObject *, gdouble, gdouble, gpointer, gpointer, gpointer);

	static gulong Window_signal_connect_to_embedder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "to-embedder", G_CALLBACK(window_toEmbedderHandler), data);
	}

*/
import "C"

type signalWindowFromEmbedderDetail struct {
	callback  WindowSignalFromEmbedderCallback
	handlerID C.gulong
}

var signalWindowFromEmbedderId int
var signalWindowFromEmbedderMap = make(map[int]signalWindowFromEmbedderDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_from_embedder(instance, C.gpointer(uintptr(signalWindowFromEmbedderId)))

	detail := signalWindowFromEmbedderDetail{callback, handlerID}
	signalWindowFromEmbedderMap[signalWindowFromEmbedderId] = detail

	return signalWindowFromEmbedderId
}

/*
DisconnectFromEmbedder disconnects a callback from the 'from-embedder' signal for the Window.

The connectionID should be a value returned from a call to ConnectFromEmbedder.
*/
func (recv *Window) DisconnectFromEmbedder(connectionID int) {
	signalWindowFromEmbedderLock.Lock()
	defer signalWindowFromEmbedderLock.Unlock()

	detail, exists := signalWindowFromEmbedderMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowFromEmbedderMap, connectionID)
}

//export window_fromEmbedderHandler
func window_fromEmbedderHandler(_ *C.GObject, c_embedder_x C.gdouble, c_embedder_y C.gdouble, c_offscreen_x C.gpointer, c_offscreen_y C.gpointer, data C.gpointer) {
	embedderX := float64(c_embedder_x)

	embedderY := float64(c_embedder_y)

	offscreenX := uintptr(c_offscreen_x)

	offscreenY := uintptr(c_offscreen_y)

	index := int(uintptr(data))
	callback := signalWindowFromEmbedderMap[index].callback
	callback(embedderX, embedderY, offscreenX, offscreenY)
}

type signalWindowPickEmbeddedChildDetail struct {
	callback  WindowSignalPickEmbeddedChildCallback
	handlerID C.gulong
}

var signalWindowPickEmbeddedChildId int
var signalWindowPickEmbeddedChildMap = make(map[int]signalWindowPickEmbeddedChildDetail)
var signalWindowPickEmbeddedChildLock sync.Mutex

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
	x := float64(c_x)

	y := float64(c_y)

	index := int(uintptr(data))
	callback := signalWindowPickEmbeddedChildMap[index].callback
	retGo := callback(x, y)
	retC :=
		(*C.GdkWindow)(retGo.ToC())
	return retC
}

type signalWindowToEmbedderDetail struct {
	callback  WindowSignalToEmbedderCallback
	handlerID C.gulong
}

var signalWindowToEmbedderId int
var signalWindowToEmbedderMap = make(map[int]signalWindowToEmbedderDetail)
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
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_to_embedder(instance, C.gpointer(uintptr(signalWindowToEmbedderId)))

	detail := signalWindowToEmbedderDetail{callback, handlerID}
	signalWindowToEmbedderMap[signalWindowToEmbedderId] = detail

	return signalWindowToEmbedderId
}

/*
DisconnectToEmbedder disconnects a callback from the 'to-embedder' signal for the Window.

The connectionID should be a value returned from a call to ConnectToEmbedder.
*/
func (recv *Window) DisconnectToEmbedder(connectionID int) {
	signalWindowToEmbedderLock.Lock()
	defer signalWindowToEmbedderLock.Unlock()

	detail, exists := signalWindowToEmbedderMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowToEmbedderMap, connectionID)
}

//export window_toEmbedderHandler
func window_toEmbedderHandler(_ *C.GObject, c_offscreen_x C.gdouble, c_offscreen_y C.gdouble, c_embedder_x C.gpointer, c_embedder_y C.gpointer, data C.gpointer) {
	offscreenX := float64(c_offscreen_x)

	offscreenY := float64(c_offscreen_y)

	embedderX := uintptr(c_embedder_x)

	embedderY := uintptr(c_embedder_y)

	index := int(uintptr(data))
	callback := signalWindowToEmbedderMap[index].callback
	callback(offscreenX, offscreenY, embedderX, embedderY)
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
