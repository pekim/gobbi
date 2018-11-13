// This is a generated file - DO NOT EDIT
// +build gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns whether the Scroll Lock modifer is locked.
/*

C function : gdk_keymap_get_scroll_lock_state
*/
func (recv *Keymap) GetScrollLockState() bool {
	retC := C.gdk_keymap_get_scroll_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether input to the window is passed through to the window
// below.
//
// See gdk_window_set_pass_through() for details
/*

C function : gdk_window_get_pass_through
*/
func (recv *Window) GetPassThrough() bool {
	retC := C.gdk_window_get_pass_through((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether input to the window is passed through to the window
// below.
//
// The default value of this is %FALSE, which means that pointer
// events that happen inside the window are send first to the window,
// but if the event is not selected by the event mask then the event
// is sent to the parent window, and so on up the hierarchy.
//
// If @pass_through is %TRUE then such pointer events happen as if the
// window wasn't there at all, and thus will be sent first to any
// windows below @window. This is useful if the window is used in a
// transparent fashion. In the terminology of the web this would be called
// "pointer-events: none".
//
// Note that a window with @pass_through %TRUE can still have a subwindow
// without pass through, so you can get events on a subset of a window. And in
// that cases you would get the in-between related events such as the pointer
// enter/leave events on its way to the destination window.
/*

C function : gdk_window_set_pass_through
*/
func (recv *Window) SetPassThrough(passThrough bool) {
	c_pass_through :=
		boolToGboolean(passThrough)

	C.gdk_window_set_pass_through((*C.GdkWindow)(recv.native), c_pass_through)

	return
}
