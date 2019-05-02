// Code generated - DO NOT EDIT.
// +build gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetScrollLockState is a wrapper around the C function gdk_keymap_get_scroll_lock_state.
func (recv *Keymap) GetScrollLockState() bool {
	retC := C.gdk_keymap_get_scroll_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPassThrough is a wrapper around the C function gdk_window_get_pass_through.
func (recv *Window) GetPassThrough() bool {
	retC := C.gdk_window_get_pass_through((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPassThrough is a wrapper around the C function gdk_window_set_pass_through.
func (recv *Window) SetPassThrough(passThrough bool) {
	c_pass_through :=
		boolToGboolean(passThrough)

	C.gdk_window_set_pass_through((*C.GdkWindow)(recv.native), c_pass_through)

	return
}
