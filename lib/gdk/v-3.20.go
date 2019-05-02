// Code generated - DO NOT EDIT.
// +build gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

	void dragcontext_actionChangedHandler(GObject *, GdkDragAction, gpointer);

	static gulong DragContext_signal_connect_action_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-changed", G_CALLBACK(dragcontext_actionChangedHandler), data);
	}

*/
/*

	void dragcontext_cancelHandler(GObject *, GdkDragCancelReason, gpointer);

	static gulong DragContext_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", G_CALLBACK(dragcontext_cancelHandler), data);
	}

*/
/*

	void dragcontext_dndFinishedHandler(GObject *, gpointer);

	static gulong DragContext_signal_connect_dnd_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "dnd-finished", G_CALLBACK(dragcontext_dndFinishedHandler), data);
	}

*/
/*

	void dragcontext_dropPerformedHandler(GObject *, gint, gpointer);

	static gulong DragContext_signal_connect_drop_performed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drop-performed", G_CALLBACK(dragcontext_dropPerformedHandler), data);
	}

*/
import "C"

type SeatCapabilities C.GdkSeatCapabilities

const (
	GDK_SEAT_CAPABILITY_NONE          SeatCapabilities = 0
	GDK_SEAT_CAPABILITY_POINTER       SeatCapabilities = 1
	GDK_SEAT_CAPABILITY_TOUCH         SeatCapabilities = 2
	GDK_SEAT_CAPABILITY_TABLET_STYLUS SeatCapabilities = 4
	GDK_SEAT_CAPABILITY_KEYBOARD      SeatCapabilities = 8
	GDK_SEAT_CAPABILITY_ALL_POINTING  SeatCapabilities = 7
	GDK_SEAT_CAPABILITY_ALL           SeatCapabilities = 15
)

// GetSeat is a wrapper around the C function gdk_device_get_seat.
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
var signalDisplaySeatAddedLock sync.RWMutex

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
	signalDisplaySeatAddedLock.RLock()
	defer signalDisplaySeatAddedLock.RUnlock()

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
var signalDisplaySeatRemovedLock sync.RWMutex

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
	signalDisplaySeatRemovedLock.RLock()
	defer signalDisplaySeatRemovedLock.RUnlock()

	seat := SeatNewFromC(unsafe.Pointer(c_seat))

	index := int(uintptr(data))
	callback := signalDisplaySeatRemovedMap[index].callback
	callback(seat)
}

// GetDefaultSeat is a wrapper around the C function gdk_display_get_default_seat.
func (recv *Display) GetDefaultSeat() *Seat {
	retC := C.gdk_display_get_default_seat((*C.GdkDisplay)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListSeats is a wrapper around the C function gdk_display_list_seats.
func (recv *Display) ListSeats() *glib.List {
	retC := C.gdk_display_list_seats((*C.GdkDisplay)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalDragContextActionChangedDetail struct {
	callback  DragContextSignalActionChangedCallback
	handlerID C.gulong
}

var signalDragContextActionChangedId int
var signalDragContextActionChangedMap = make(map[int]signalDragContextActionChangedDetail)
var signalDragContextActionChangedLock sync.RWMutex

// DragContextSignalActionChangedCallback is a callback function for a 'action-changed' signal emitted from a DragContext.
type DragContextSignalActionChangedCallback func(action DragAction)

/*
ConnectActionChanged connects the callback to the 'action-changed' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectActionChanged to remove it.
*/
func (recv *DragContext) ConnectActionChanged(callback DragContextSignalActionChangedCallback) int {
	signalDragContextActionChangedLock.Lock()
	defer signalDragContextActionChangedLock.Unlock()

	signalDragContextActionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_action_changed(instance, C.gpointer(uintptr(signalDragContextActionChangedId)))

	detail := signalDragContextActionChangedDetail{callback, handlerID}
	signalDragContextActionChangedMap[signalDragContextActionChangedId] = detail

	return signalDragContextActionChangedId
}

/*
DisconnectActionChanged disconnects a callback from the 'action-changed' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectActionChanged.
*/
func (recv *DragContext) DisconnectActionChanged(connectionID int) {
	signalDragContextActionChangedLock.Lock()
	defer signalDragContextActionChangedLock.Unlock()

	detail, exists := signalDragContextActionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextActionChangedMap, connectionID)
}

//export dragcontext_actionChangedHandler
func dragcontext_actionChangedHandler(_ *C.GObject, c_action C.GdkDragAction, data C.gpointer) {
	signalDragContextActionChangedLock.RLock()
	defer signalDragContextActionChangedLock.RUnlock()

	action := DragAction(c_action)

	index := int(uintptr(data))
	callback := signalDragContextActionChangedMap[index].callback
	callback(action)
}

type signalDragContextCancelDetail struct {
	callback  DragContextSignalCancelCallback
	handlerID C.gulong
}

var signalDragContextCancelId int
var signalDragContextCancelMap = make(map[int]signalDragContextCancelDetail)
var signalDragContextCancelLock sync.RWMutex

// DragContextSignalCancelCallback is a callback function for a 'cancel' signal emitted from a DragContext.
type DragContextSignalCancelCallback func(reason DragCancelReason)

/*
ConnectCancel connects the callback to the 'cancel' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectCancel to remove it.
*/
func (recv *DragContext) ConnectCancel(callback DragContextSignalCancelCallback) int {
	signalDragContextCancelLock.Lock()
	defer signalDragContextCancelLock.Unlock()

	signalDragContextCancelId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_cancel(instance, C.gpointer(uintptr(signalDragContextCancelId)))

	detail := signalDragContextCancelDetail{callback, handlerID}
	signalDragContextCancelMap[signalDragContextCancelId] = detail

	return signalDragContextCancelId
}

/*
DisconnectCancel disconnects a callback from the 'cancel' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectCancel.
*/
func (recv *DragContext) DisconnectCancel(connectionID int) {
	signalDragContextCancelLock.Lock()
	defer signalDragContextCancelLock.Unlock()

	detail, exists := signalDragContextCancelMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextCancelMap, connectionID)
}

//export dragcontext_cancelHandler
func dragcontext_cancelHandler(_ *C.GObject, c_reason C.GdkDragCancelReason, data C.gpointer) {
	signalDragContextCancelLock.RLock()
	defer signalDragContextCancelLock.RUnlock()

	reason := DragCancelReason(c_reason)

	index := int(uintptr(data))
	callback := signalDragContextCancelMap[index].callback
	callback(reason)
}

type signalDragContextDndFinishedDetail struct {
	callback  DragContextSignalDndFinishedCallback
	handlerID C.gulong
}

var signalDragContextDndFinishedId int
var signalDragContextDndFinishedMap = make(map[int]signalDragContextDndFinishedDetail)
var signalDragContextDndFinishedLock sync.RWMutex

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
	signalDragContextDndFinishedLock.RLock()
	defer signalDragContextDndFinishedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDragContextDndFinishedMap[index].callback
	callback()
}

type signalDragContextDropPerformedDetail struct {
	callback  DragContextSignalDropPerformedCallback
	handlerID C.gulong
}

var signalDragContextDropPerformedId int
var signalDragContextDropPerformedMap = make(map[int]signalDragContextDropPerformedDetail)
var signalDragContextDropPerformedLock sync.RWMutex

// DragContextSignalDropPerformedCallback is a callback function for a 'drop-performed' signal emitted from a DragContext.
type DragContextSignalDropPerformedCallback func(time int32)

/*
ConnectDropPerformed connects the callback to the 'drop-performed' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectDropPerformed to remove it.
*/
func (recv *DragContext) ConnectDropPerformed(callback DragContextSignalDropPerformedCallback) int {
	signalDragContextDropPerformedLock.Lock()
	defer signalDragContextDropPerformedLock.Unlock()

	signalDragContextDropPerformedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_drop_performed(instance, C.gpointer(uintptr(signalDragContextDropPerformedId)))

	detail := signalDragContextDropPerformedDetail{callback, handlerID}
	signalDragContextDropPerformedMap[signalDragContextDropPerformedId] = detail

	return signalDragContextDropPerformedId
}

/*
DisconnectDropPerformed disconnects a callback from the 'drop-performed' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectDropPerformed.
*/
func (recv *DragContext) DisconnectDropPerformed(connectionID int) {
	signalDragContextDropPerformedLock.Lock()
	defer signalDragContextDropPerformedLock.Unlock()

	detail, exists := signalDragContextDropPerformedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextDropPerformedMap, connectionID)
}

//export dragcontext_dropPerformedHandler
func dragcontext_dropPerformedHandler(_ *C.GObject, c_time C.gint, data C.gpointer) {
	signalDragContextDropPerformedLock.RLock()
	defer signalDragContextDropPerformedLock.RUnlock()

	time := int32(c_time)

	index := int(uintptr(data))
	callback := signalDragContextDropPerformedMap[index].callback
	callback(time)
}

// GetDragWindow is a wrapper around the C function gdk_drag_context_get_drag_window.
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

// ManageDnd is a wrapper around the C function gdk_drag_context_manage_dnd.
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

// SetHotspot is a wrapper around the C function gdk_drag_context_set_hotspot.
func (recv *DragContext) SetHotspot(hotX int32, hotY int32) {
	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gdk_drag_context_set_hotspot((*C.GdkDragContext)(recv.native), c_hot_x, c_hot_y)

	return
}

// IsLegacy is a wrapper around the C function gdk_gl_context_is_legacy.
func (recv *GLContext) IsLegacy() bool {
	retC := C.gdk_gl_context_is_legacy((*C.GdkGLContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type DragCancelReason C.GdkDragCancelReason

const (
	GDK_DRAG_CANCEL_NO_TARGET      DragCancelReason = 0
	GDK_DRAG_CANCEL_USER_CANCELLED DragCancelReason = 1
	GDK_DRAG_CANCEL_ERROR          DragCancelReason = 2
)

// DragBeginFromPoint is a wrapper around the C function gdk_drag_begin_from_point.
func DragBeginFromPoint(window *Window, device *Device, targets *glib.List, xRoot int32, yRoot int32) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	retC := C.gdk_drag_begin_from_point(c_window, c_device, c_targets, c_x_root, c_y_root)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragDropDone is a wrapper around the C function gdk_drag_drop_done.
func DragDropDone(context *DragContext, success bool) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	C.gdk_drag_drop_done(c_context, c_success)

	return
}

// Equal is a wrapper around the C function gdk_rectangle_equal.
func (recv *Rectangle) Equal(rect2 *Rectangle) bool {
	c_rect2 := (*C.GdkRectangle)(C.NULL)
	if rect2 != nil {
		c_rect2 = (*C.GdkRectangle)(rect2.ToC())
	}

	retC := C.gdk_rectangle_equal((*C.GdkRectangle)(recv.native), c_rect2)
	retGo := retC == C.TRUE

	return retGo
}
