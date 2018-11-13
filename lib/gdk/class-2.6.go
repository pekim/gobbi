// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_display_request_selection_notification : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gdk_display_store_clipboard : unsupported parameter targets :

// Returns whether the speicifed display supports clipboard
// persistance; i.e. if it’s possible to store the clipboard data after an
// application has quit. On X11 this checks if a clipboard daemon is
// running.
/*

C function

gdk_display_supports_clipboard_persistence
*/
func (recv *Display) SupportsClipboardPersistence() bool {
	retC := C.gdk_display_supports_clipboard_persistence((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether #GdkEventOwnerChange events will be
// sent when the owner of a selection changes.
/*

C function

gdk_display_supports_selection_notification
*/
func (recv *Display) SupportsSelectionNotification() bool {
	retC := C.gdk_display_supports_selection_notification((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Does nothing, present only for compatiblity.
/*

C function

gdk_window_configure_finished
*/
func (recv *Window) ConfigureFinished() {
	C.gdk_window_configure_finished((*C.GdkWindow)(recv.native))

	return
}

// Does nothing, present only for compatiblity.
/*

C function

gdk_window_enable_synchronized_configure
*/
func (recv *Window) EnableSynchronizedConfigure() {
	C.gdk_window_enable_synchronized_configure((*C.GdkWindow)(recv.native))

	return
}

// Setting @focus_on_map to %FALSE hints the desktop environment that the
// window doesn’t want to receive input focus when it is mapped.
// focus_on_map should be turned off for windows that aren’t triggered
// interactively (such as popups from network activity).
//
// On X, it is the responsibility of the window manager to interpret
// this hint. Window managers following the freedesktop.org window
// manager extension specification should respect it.
/*

C function

gdk_window_set_focus_on_map
*/
func (recv *Window) SetFocusOnMap(focusOnMap bool) {
	c_focus_on_map :=
		boolToGboolean(focusOnMap)

	C.gdk_window_set_focus_on_map((*C.GdkWindow)(recv.native), c_focus_on_map)

	return
}
