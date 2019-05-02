// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// RequestSelectionNotification is a wrapper around the C function gdk_display_request_selection_notification.
func (recv *Display) RequestSelectionNotification(selection *Atom) bool {
	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gdk_display_request_selection_notification((*C.GdkDisplay)(recv.native), c_selection)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets :

// SupportsClipboardPersistence is a wrapper around the C function gdk_display_supports_clipboard_persistence.
func (recv *Display) SupportsClipboardPersistence() bool {
	retC := C.gdk_display_supports_clipboard_persistence((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsSelectionNotification is a wrapper around the C function gdk_display_supports_selection_notification.
func (recv *Display) SupportsSelectionNotification() bool {
	retC := C.gdk_display_supports_selection_notification((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ConfigureFinished is a wrapper around the C function gdk_window_configure_finished.
func (recv *Window) ConfigureFinished() {
	C.gdk_window_configure_finished((*C.GdkWindow)(recv.native))

	return
}

// EnableSynchronizedConfigure is a wrapper around the C function gdk_window_enable_synchronized_configure.
func (recv *Window) EnableSynchronizedConfigure() {
	C.gdk_window_enable_synchronized_configure((*C.GdkWindow)(recv.native))

	return
}

// SetFocusOnMap is a wrapper around the C function gdk_window_set_focus_on_map.
func (recv *Window) SetFocusOnMap(focusOnMap bool) {
	c_focus_on_map :=
		boolToGboolean(focusOnMap)

	C.gdk_window_set_focus_on_map((*C.GdkWindow)(recv.native), c_focus_on_map)

	return
}

// DragDropSucceeded is a wrapper around the C function gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	retC := C.gdk_drag_drop_succeeded(c_context)
	retGo := retC == C.TRUE

	return retGo
}

// EventOwnerChange is a wrapper around the C record GdkEventOwnerChange.
type EventOwnerChange struct {
	native *C.GdkEventOwnerChange
	Type   EventType
	// window : record
	SendEvent int8
	// owner : record
	Reason OwnerChange
	// selection : record
	Time          uint32
	SelectionTime uint32
}

func EventOwnerChangeNewFromC(u unsafe.Pointer) *EventOwnerChange {
	c := (*C.GdkEventOwnerChange)(u)
	if c == nil {
		return nil
	}

	g := &EventOwnerChange{
		Reason:        (OwnerChange)(c.reason),
		SelectionTime: (uint32)(c.selection_time),
		SendEvent:     (int8)(c.send_event),
		Time:          (uint32)(c.time),
		Type:          (EventType)(c._type),
		native:        c,
	}

	return g
}

func (recv *EventOwnerChange) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GdkEventType)(recv.Type)
	recv.native.send_event =
		(C.gint8)(recv.SendEvent)
	recv.native.reason =
		(C.GdkOwnerChange)(recv.Reason)
	recv.native.time =
		(C.guint32)(recv.Time)
	recv.native.selection_time =
		(C.guint32)(recv.SelectionTime)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EventOwnerChange with another EventOwnerChange, and returns true if they represent the same GObject.
func (recv *EventOwnerChange) Equals(other *EventOwnerChange) bool {
	return other.ToC() == recv.ToC()
}
