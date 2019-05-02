// This is a generated file - DO NOT EDIT
// +build gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// SupportsComposite is a wrapper around the C function gdk_display_supports_composite.
func (recv *Display) SupportsComposite() bool {
	retC := C.gdk_display_supports_composite((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HaveBidiLayouts is a wrapper around the C function gdk_keymap_have_bidi_layouts.
func (recv *Keymap) HaveBidiLayouts() bool {
	retC := C.gdk_keymap_have_bidi_layouts((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Beep is a wrapper around the C function gdk_window_beep.
func (recv *Window) Beep() {
	C.gdk_window_beep((*C.GdkWindow)(recv.native))

	return
}

// SetComposited is a wrapper around the C function gdk_window_set_composited.
func (recv *Window) SetComposited(composited bool) {
	c_composited :=
		boolToGboolean(composited)

	C.gdk_window_set_composited((*C.GdkWindow)(recv.native), c_composited)

	return
}

// SetOpacity is a wrapper around the C function gdk_window_set_opacity.
func (recv *Window) SetOpacity(opacity float64) {
	c_opacity := (C.gdouble)(opacity)

	C.gdk_window_set_opacity((*C.GdkWindow)(recv.native), c_opacity)

	return
}

// SetStartupId is a wrapper around the C function gdk_window_set_startup_id.
func (recv *Window) SetStartupId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_window_set_startup_id((*C.GdkWindow)(recv.native), c_startup_id)

	return
}
