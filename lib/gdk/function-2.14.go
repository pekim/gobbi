// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// TestRenderSync is a wrapper around the C function gdk_test_render_sync.
func TestRenderSync(window *Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gdk_test_render_sync(c_window)

	return
}

// TestSimulateButton is a wrapper around the C function gdk_test_simulate_button.
func TestSimulateButton(window *Window, x int32, y int32, button uint32, modifiers ModifierType, buttonPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(window.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_button_pressrelease := (C.GdkEventType)(buttonPressrelease)

	retC := C.gdk_test_simulate_button(c_window, c_x, c_y, c_button, c_modifiers, c_button_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// TestSimulateKey is a wrapper around the C function gdk_test_simulate_key.
func TestSimulateKey(window *Window, x int32, y int32, keyval uint32, modifiers ModifierType, keyPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(window.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_key_pressrelease := (C.GdkEventType)(keyPressrelease)

	retC := C.gdk_test_simulate_key(c_window, c_x, c_y, c_keyval, c_modifiers, c_key_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc, GSourceFunc
