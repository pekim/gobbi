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

	void Display_seatAddedHandler();

	static gulong Display_signal_connect_seat_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "seat-added", Display_seatAddedHandler, data);
	}

*/
/*

	void Display_seatRemovedHandler();

	static gulong Display_signal_connect_seat_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "seat-removed", Display_seatRemovedHandler, data);
	}

*/
/*

	void DragContext_actionChangedHandler();

	static gulong DragContext_signal_connect_action_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-changed", DragContext_actionChangedHandler, data);
	}

*/
/*

	void DragContext_cancelHandler();

	static gulong DragContext_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", DragContext_cancelHandler, data);
	}

*/
/*

	void DragContext_dndFinishedHandler();

	static gulong DragContext_signal_connect_dnd_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "dnd-finished", DragContext_dndFinishedHandler, data);
	}

*/
/*

	void DragContext_dropPerformedHandler();

	static gulong DragContext_signal_connect_drop_performed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drop-performed", DragContext_dropPerformedHandler, data);
	}

*/
import "C"

// GetSeat is a wrapper around the C function gdk_device_get_seat.
func (recv *Device) GetSeat() *Seat {
	retC := C.gdk_device_get_seat((*C.GdkDevice)(recv.native))
	retGo := SeatNewFromC(unsafe.Pointer(retC))

	return retGo
}

var signalDisplaySeatAddedId int
var signalDisplaySeatAddedMap = make(map[int]DisplaySignalSeatAddedCallback)
var signalDisplaySeatAddedLock sync.Mutex

// DisplaySignalSeatAddedCallback is a callback function for a 'seat-added' signal emitted from a Display.
type DisplaySignalSeatAddedCallback func(seat *Seat)

func (recv *Display) ConnectSeatAdded() {}

func Display_seatAddedHandler() {}

var signalDisplaySeatRemovedId int
var signalDisplaySeatRemovedMap = make(map[int]DisplaySignalSeatRemovedCallback)
var signalDisplaySeatRemovedLock sync.Mutex

// DisplaySignalSeatRemovedCallback is a callback function for a 'seat-removed' signal emitted from a Display.
type DisplaySignalSeatRemovedCallback func(seat *Seat)

func (recv *Display) ConnectSeatRemoved() {}

func Display_seatRemovedHandler() {}

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

var signalDragContextActionChangedId int
var signalDragContextActionChangedMap = make(map[int]DragContextSignalActionChangedCallback)
var signalDragContextActionChangedLock sync.Mutex

// DragContextSignalActionChangedCallback is a callback function for a 'action-changed' signal emitted from a DragContext.
type DragContextSignalActionChangedCallback func(action DragAction)

func (recv *DragContext) ConnectActionChanged() {}

func DragContext_actionChangedHandler() {}

var signalDragContextCancelId int
var signalDragContextCancelMap = make(map[int]DragContextSignalCancelCallback)
var signalDragContextCancelLock sync.Mutex

// DragContextSignalCancelCallback is a callback function for a 'cancel' signal emitted from a DragContext.
type DragContextSignalCancelCallback func(reason DragCancelReason)

func (recv *DragContext) ConnectCancel() {}

func DragContext_cancelHandler() {}

var signalDragContextDndFinishedId int
var signalDragContextDndFinishedMap = make(map[int]DragContextSignalDndFinishedCallback)
var signalDragContextDndFinishedLock sync.Mutex

// DragContextSignalDndFinishedCallback is a callback function for a 'dnd-finished' signal emitted from a DragContext.
type DragContextSignalDndFinishedCallback func()

func (recv *DragContext) ConnectDndFinished() {}

func DragContext_dndFinishedHandler() {}

var signalDragContextDropPerformedId int
var signalDragContextDropPerformedMap = make(map[int]DragContextSignalDropPerformedCallback)
var signalDragContextDropPerformedLock sync.Mutex

// DragContextSignalDropPerformedCallback is a callback function for a 'drop-performed' signal emitted from a DragContext.
type DragContextSignalDropPerformedCallback func(time int32)

func (recv *DragContext) ConnectDropPerformed() {}

func DragContext_dropPerformedHandler() {}

// GetDragWindow is a wrapper around the C function gdk_drag_context_get_drag_window.
func (recv *DragContext) GetDragWindow() *Window {
	retC := C.gdk_drag_context_get_drag_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ManageDnd is a wrapper around the C function gdk_drag_context_manage_dnd.
func (recv *DragContext) ManageDnd(ipcWindow *Window, actions DragAction) bool {
	c_ipc_window := (*C.GdkWindow)(ipcWindow.ToC())

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
