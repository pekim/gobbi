// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

func (recv *AppLaunchContext) AppLaunchContext() *gio.AppLaunchContext {}

func (recv *Cursor) Object() *gobject.Object {}

func (recv *Device) Object() *gobject.Object {}

func (recv *DeviceManager) Object() *gobject.Object {}

func (recv *DeviceTool) Object() *gobject.Object {}

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets : no param type

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

func (recv *Display) Object() *gobject.Object {}

func (recv *DisplayManager) Object() *gobject.Object {}

func (recv *DragContext) Object() *gobject.Object {}

func (recv *DrawingContext) Object() *gobject.Object {}

func (recv *FrameClock) Object() *gobject.Object {}

func (recv *GLContext) Object() *gobject.Object {}

func (recv *Keymap) Object() *gobject.Object {}

func (recv *Monitor) Object() *gobject.Object {}

func (recv *Screen) Object() *gobject.Object {}

func (recv *Seat) Object() *gobject.Object {}

func (recv *Visual) Object() *gobject.Object {}

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

func (recv *Window) Object() *gobject.Object {}
