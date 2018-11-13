// This is a generated file - DO NOT EDIT
// +build gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns %TRUE if gdk_window_set_composited() can be used
// to redirect drawing on the window using compositing.
//
// Currently this only works on X11 with XComposite and
// XDamage extensions available.
/*

C function

gdk_display_supports_composite
*/
func (recv *Display) SupportsComposite() bool {
	retC := C.gdk_display_supports_composite((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Determines if keyboard layouts for both right-to-left and left-to-right
// languages are in use.
/*

C function

gdk_keymap_have_bidi_layouts
*/
func (recv *Keymap) HaveBidiLayouts() bool {
	retC := C.gdk_keymap_have_bidi_layouts((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Emits a short beep associated to @window in the appropriate
// display, if supported. Otherwise, emits a short beep on
// the display just as gdk_display_beep().
/*

C function

gdk_window_beep
*/
func (recv *Window) Beep() {
	C.gdk_window_beep((*C.GdkWindow)(recv.native))

	return
}

// Sets a #GdkWindow as composited, or unsets it. Composited
// windows do not automatically have their contents drawn to
// the screen. Drawing is redirected to an offscreen buffer
// and an expose event is emitted on the parent of the composited
// window. It is the responsibility of the parent’s expose handler
// to manually merge the off-screen content onto the screen in
// whatever way it sees fit.
//
// It only makes sense for child windows to be composited; see
// gdk_window_set_opacity() if you need translucent toplevel
// windows.
//
// An additional effect of this call is that the area of this
// window is no longer clipped from regions marked for
// invalidation on its parent. Draws done on the parent
// window are also no longer clipped by the child.
//
// This call is only supported on some systems (currently,
// only X11 with new enough Xcomposite and Xdamage extensions).
// You must call gdk_display_supports_composite() to check if
// setting a window as composited is supported before
// attempting to do so.
/*

C function

gdk_window_set_composited
*/
func (recv *Window) SetComposited(composited bool) {
	c_composited :=
		boolToGboolean(composited)

	C.gdk_window_set_composited((*C.GdkWindow)(recv.native), c_composited)

	return
}

// Set @window to render as partially transparent,
// with opacity 0 being fully transparent and 1 fully opaque. (Values
// of the opacity parameter are clamped to the [0,1] range.)
//
// For toplevel windows this depends on support from the windowing system
// that may not always be there. For instance, On X11, this works only on
// X screens with a compositing manager running. On Wayland, there is no
// per-window opacity value that the compositor would apply. Instead, use
// `gdk_window_set_opaque_region (window, NULL)` to tell the compositor
// that the entire window is (potentially) non-opaque, and draw your content
// with alpha, or use gtk_widget_set_opacity() to set an overall opacity
// for your widgets.
//
// For child windows this function only works for non-native windows.
//
// For setting up per-pixel alpha topelevels, see gdk_screen_get_rgba_visual(),
// and for non-toplevels, see gdk_window_set_composited().
//
// Support for non-toplevel windows was added in 3.8.
/*

C function

gdk_window_set_opacity
*/
func (recv *Window) SetOpacity(opacity float64) {
	c_opacity := (C.gdouble)(opacity)

	C.gdk_window_set_opacity((*C.GdkWindow)(recv.native), c_opacity)

	return
}

// When using GTK+, typically you should use gtk_window_set_startup_id()
// instead of this low-level function.
/*

C function

gdk_window_set_startup_id
*/
func (recv *Window) SetStartupId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_window_set_startup_id((*C.GdkWindow)(recv.native), c_startup_id)

	return
}
