// This is a generated file - DO NOT EDIT
// +build gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Retrieves a pixel from @window to force the windowing
// system to carry out any pending rendering commands.
//
// This function is intended to be used to synchronize with rendering
// pipelines, to benchmark windowing system rendering operations.
/*

C function

gdk_test_render_sync
*/
func TestRenderSync(window *Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gdk_test_render_sync(c_window)

	return
}

// This function is intended to be used in GTK+ test programs.
// It will warp the mouse pointer to the given (@x,@y) coordinates
// within @window and simulate a button press or release event.
// Because the mouse pointer needs to be warped to the target
// location, use of this function outside of test programs that
// run in their own virtual windowing system (e.g. Xvfb) is not
// recommended.
//
// Also, gdk_test_simulate_button() is a fairly low level function,
// for most testing purposes, gtk_test_widget_click() is the right
// function to call which will generate a button press event followed
// by its accompanying button release event.
/*

C function

gdk_test_simulate_button
*/
func TestSimulateButton(window *Window, x int32, y int32, button uint32, modifiers ModifierType, buttonPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_button_pressrelease := (C.GdkEventType)(buttonPressrelease)

	retC := C.gdk_test_simulate_button(c_window, c_x, c_y, c_button, c_modifiers, c_button_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// This function is intended to be used in GTK+ test programs.
// If (@x,@y) are > (-1,-1), it will warp the mouse pointer to
// the given (@x,@y) coordinates within @window and simulate a
// key press or release event.
//
// When the mouse pointer is warped to the target location, use
// of this function outside of test programs that run in their
// own virtual windowing system (e.g. Xvfb) is not recommended.
// If (@x,@y) are passed as (-1,-1), the mouse pointer will not
// be warped and @window origin will be used as mouse pointer
// location for the event.
//
// Also, gdk_test_simulate_key() is a fairly low level function,
// for most testing purposes, gtk_test_widget_send_key() is the
// right function to call which will generate a key press event
// followed by its accompanying key release event.
/*

C function

gdk_test_simulate_key
*/
func TestSimulateKey(window *Window, x int32, y int32, keyval uint32, modifiers ModifierType, keyPressrelease EventType) bool {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_key_pressrelease := (C.GdkEventType)(keyPressrelease)

	retC := C.gdk_test_simulate_key(c_window, c_x, c_y, c_keyval, c_modifiers, c_key_pressrelease)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_threads_add_timeout_seconds : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : gdk_threads_add_timeout_seconds_full : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function
