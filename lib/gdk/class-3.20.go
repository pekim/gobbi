// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void display_seatAddedHandler(GObject *, GdkSeat *, gpointer);

	static gulong Display_signal_connect_seat_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "seat-added", G_CALLBACK(display_seatAddedHandler), data);
	}

*/
/*

	void display_seatRemovedHandler(GObject *, GdkSeat *, gpointer);

	static gulong Display_signal_connect_seat_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "seat-removed", G_CALLBACK(display_seatRemovedHandler), data);
	}

*/
/*

	void dragcontext_dndFinishedHandler(GObject *, gpointer);

	static gulong DragContext_signal_connect_dnd_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "dnd-finished", G_CALLBACK(dragcontext_dndFinishedHandler), data);
	}

*/
import "C"

// Returns the #GdkSeat the device belongs to.
/*

C function

gdk_device_get_seat
*/
func (recv *Device) GetSeat() *Seat {
	retC := C.gdk_device_get_seat((*C.GdkDevice)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalDisplaySeatAddedDetail struct {
	callback  DisplaySignalSeatAddedCallback
	handlerID C.gulong
}

var signalDisplaySeatAddedId int
var signalDisplaySeatAddedMap = make(map[int]signalDisplaySeatAddedDetail)
var signalDisplaySeatAddedLock sync.Mutex

// DisplaySignalSeatAddedCallback is a callback function for a 'seat-added' signal emitted from a Display.
type DisplaySignalSeatAddedCallback func(seat *Seat)

/*
ConnectSeatAdded connects the callback to the 'seat-added' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectSeatAdded to remove it.
*/
func (recv *Display) ConnectSeatAdded(callback DisplaySignalSeatAddedCallback) int {
	signalDisplaySeatAddedLock.Lock()
	defer signalDisplaySeatAddedLock.Unlock()

	signalDisplaySeatAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_seat_added(instance, C.gpointer(uintptr(signalDisplaySeatAddedId)))

	detail := signalDisplaySeatAddedDetail{callback, handlerID}
	signalDisplaySeatAddedMap[signalDisplaySeatAddedId] = detail

	return signalDisplaySeatAddedId
}

/*
DisconnectSeatAdded disconnects a callback from the 'seat-added' signal for the Display.

The connectionID should be a value returned from a call to ConnectSeatAdded.
*/
func (recv *Display) DisconnectSeatAdded(connectionID int) {
	signalDisplaySeatAddedLock.Lock()
	defer signalDisplaySeatAddedLock.Unlock()

	detail, exists := signalDisplaySeatAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplaySeatAddedMap, connectionID)
}

//export display_seatAddedHandler
func display_seatAddedHandler(_ *C.GObject, c_seat *C.GdkSeat, data C.gpointer) {
	seat := SeatNewFromC(unsafe.Pointer(c_seat))

	index := int(uintptr(data))
	callback := signalDisplaySeatAddedMap[index].callback
	callback(seat)
}

type signalDisplaySeatRemovedDetail struct {
	callback  DisplaySignalSeatRemovedCallback
	handlerID C.gulong
}

var signalDisplaySeatRemovedId int
var signalDisplaySeatRemovedMap = make(map[int]signalDisplaySeatRemovedDetail)
var signalDisplaySeatRemovedLock sync.Mutex

// DisplaySignalSeatRemovedCallback is a callback function for a 'seat-removed' signal emitted from a Display.
type DisplaySignalSeatRemovedCallback func(seat *Seat)

/*
ConnectSeatRemoved connects the callback to the 'seat-removed' signal for the Display.

The returned value represents the connection, and may be passed to DisconnectSeatRemoved to remove it.
*/
func (recv *Display) ConnectSeatRemoved(callback DisplaySignalSeatRemovedCallback) int {
	signalDisplaySeatRemovedLock.Lock()
	defer signalDisplaySeatRemovedLock.Unlock()

	signalDisplaySeatRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Display_signal_connect_seat_removed(instance, C.gpointer(uintptr(signalDisplaySeatRemovedId)))

	detail := signalDisplaySeatRemovedDetail{callback, handlerID}
	signalDisplaySeatRemovedMap[signalDisplaySeatRemovedId] = detail

	return signalDisplaySeatRemovedId
}

/*
DisconnectSeatRemoved disconnects a callback from the 'seat-removed' signal for the Display.

The connectionID should be a value returned from a call to ConnectSeatRemoved.
*/
func (recv *Display) DisconnectSeatRemoved(connectionID int) {
	signalDisplaySeatRemovedLock.Lock()
	defer signalDisplaySeatRemovedLock.Unlock()

	detail, exists := signalDisplaySeatRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDisplaySeatRemovedMap, connectionID)
}

//export display_seatRemovedHandler
func display_seatRemovedHandler(_ *C.GObject, c_seat *C.GdkSeat, data C.gpointer) {
	seat := SeatNewFromC(unsafe.Pointer(c_seat))

	index := int(uintptr(data))
	callback := signalDisplaySeatRemovedMap[index].callback
	callback(seat)
}

// Returns the default #GdkSeat for this display.
/*

C function

gdk_display_get_default_seat
*/
func (recv *Display) GetDefaultSeat() *Seat {
	retC := C.gdk_display_get_default_seat((*C.GdkDisplay)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the list of seats known to @display.
/*

C function

gdk_display_list_seats
*/
func (recv *Display) ListSeats() *glib.List {
	retC := C.gdk_display_list_seats((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'action-changed' for DragContext : unsupported parameter action : type DragAction :

// Unsupported signal 'cancel' for DragContext : unsupported parameter reason : type DragCancelReason :

type signalDragContextDndFinishedDetail struct {
	callback  DragContextSignalDndFinishedCallback
	handlerID C.gulong
}

var signalDragContextDndFinishedId int
var signalDragContextDndFinishedMap = make(map[int]signalDragContextDndFinishedDetail)
var signalDragContextDndFinishedLock sync.Mutex

// DragContextSignalDndFinishedCallback is a callback function for a 'dnd-finished' signal emitted from a DragContext.
type DragContextSignalDndFinishedCallback func()

/*
ConnectDndFinished connects the callback to the 'dnd-finished' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectDndFinished to remove it.
*/
func (recv *DragContext) ConnectDndFinished(callback DragContextSignalDndFinishedCallback) int {
	signalDragContextDndFinishedLock.Lock()
	defer signalDragContextDndFinishedLock.Unlock()

	signalDragContextDndFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_dnd_finished(instance, C.gpointer(uintptr(signalDragContextDndFinishedId)))

	detail := signalDragContextDndFinishedDetail{callback, handlerID}
	signalDragContextDndFinishedMap[signalDragContextDndFinishedId] = detail

	return signalDragContextDndFinishedId
}

/*
DisconnectDndFinished disconnects a callback from the 'dnd-finished' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectDndFinished.
*/
func (recv *DragContext) DisconnectDndFinished(connectionID int) {
	signalDragContextDndFinishedLock.Lock()
	defer signalDragContextDndFinishedLock.Unlock()

	detail, exists := signalDragContextDndFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextDndFinishedMap, connectionID)
}

//export dragcontext_dndFinishedHandler
func dragcontext_dndFinishedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalDragContextDndFinishedMap[index].callback
	callback()
}

// Unsupported signal 'drop-performed' for DragContext : unsupported parameter time : type gint :

// Returns the window on which the drag icon should be rendered
// during the drag operation. Note that the window may not be
// available until the drag operation has begun. GDK will move
// the window in accordance with the ongoing drag operation.
// The window is owned by @context and will be destroyed when
// the drag operation is over.
/*

C function

gdk_drag_context_get_drag_window
*/
func (recv *DragContext) GetDragWindow() *Window {
	retC := C.gdk_drag_context_get_drag_window((*C.GdkDragContext)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Requests the drag and drop operation to be managed by @context.
// When a drag and drop operation becomes managed, the #GdkDragContext
// will internally handle all input and source-side #GdkEventDND events
// as required by the windowing system.
//
// Once the drag and drop operation is managed, the drag context will
// emit the following signals:
// - The #GdkDragContext::action-changed signal whenever the final action
// to be performed by the drag and drop operation changes.
// - The #GdkDragContext::drop-performed signal after the user performs
// the drag and drop gesture (typically by releasing the mouse button).
// - The #GdkDragContext::dnd-finished signal after the drag and drop
// operation concludes (after all #GdkSelection transfers happen).
// - The #GdkDragContext::cancel signal if the drag and drop operation is
// finished but doesn't happen over an accepting destination, or is
// cancelled through other means.
/*

C function

gdk_drag_context_manage_dnd
*/
func (recv *DragContext) ManageDnd(ipcWindow *Window, actions DragAction) bool {
	c_ipc_window := (*C.GdkWindow)(C.NULL)
	if ipcWindow != nil {
		c_ipc_window = (*C.GdkWindow)(ipcWindow.ToC())
	}

	c_actions := (C.GdkDragAction)(actions)

	retC := C.gdk_drag_context_manage_dnd((*C.GdkDragContext)(recv.native), c_ipc_window, c_actions)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the position of the drag window that will be kept
// under the cursor hotspot. Initially, the hotspot is at the
// top left corner of the drag window.
/*

C function

gdk_drag_context_set_hotspot
*/
func (recv *DragContext) SetHotspot(hotX int32, hotY int32) {
	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gdk_drag_context_set_hotspot((*C.GdkDragContext)(recv.native), c_hot_x, c_hot_y)

	return
}

// Whether the #GdkGLContext is in legacy mode or not.
//
// The #GdkGLContext must be realized before calling this function.
//
// When realizing a GL context, GDK will try to use the OpenGL 3.2 core
// profile; this profile removes all the OpenGL API that was deprecated
// prior to the 3.2 version of the specification. If the realization is
// successful, this function will return %FALSE.
//
// If the underlying OpenGL implementation does not support core profiles,
// GDK will fall back to a pre-3.2 compatibility profile, and this function
// will return %TRUE.
//
// You can use the value returned by this function to decide which kind
// of OpenGL API to use, or whether to do extension discovery, or what
// kind of shader programs to load.
/*

C function

gdk_gl_context_is_legacy
*/
func (recv *GLContext) IsLegacy() bool {
	retC := C.gdk_gl_context_is_legacy((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
