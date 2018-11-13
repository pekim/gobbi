// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Unsupported signal 'from-embedder' for Window : unsupported parameter embedder_x : type gdouble :

// Unsupported signal 'pick-embedded-child' for Window : unsupported parameter x : type gdouble :

// Unsupported signal 'to-embedder' for Window : unsupported parameter offscreen_x : type gdouble :

// Tries to ensure that there is a window-system native window for this
// GdkWindow. This may fail in some situations, returning %FALSE.
//
// Offscreen window and children of them can never have native windows.
//
// Some backends may not support native child windows.
/*

C function : gdk_window_ensure_native
*/
func (recv *Window) EnsureNative() bool {
	retC := C.gdk_window_ensure_native((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This function does nothing.
/*

C function : gdk_window_flush
*/
func (recv *Window) Flush() {
	C.gdk_window_flush((*C.GdkWindow)(recv.native))

	return
}

// This function informs GDK that the geometry of an embedded
// offscreen window has changed. This is necessary for GDK to keep
// track of which offscreen window the pointer is in.
/*

C function : gdk_window_geometry_changed
*/
func (recv *Window) GeometryChanged() {
	C.gdk_window_geometry_changed((*C.GdkWindow)(recv.native))

	return
}

// Retrieves a #GdkCursor pointer for the cursor currently set on the
// specified #GdkWindow, or %NULL.  If the return value is %NULL then
// there is no custom cursor set on the specified window, and it is
// using the cursor for its parent window.
/*

C function : gdk_window_get_cursor
*/
func (recv *Window) GetCursor() *Cursor {
	retC := C.gdk_window_get_cursor((*C.GdkWindow)(recv.native))
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Obtains the position of a window position in root
// window coordinates. This is similar to
// gdk_window_get_origin() but allows you to pass
// in any position in the window, not just the origin.
/*

C function : gdk_window_get_root_coords
*/
func (recv *Window) GetRootCoords(x int32, y int32) (int32, int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_root_x C.gint

	var c_root_y C.gint

	C.gdk_window_get_root_coords((*C.GdkWindow)(recv.native), c_x, c_y, &c_root_x, &c_root_y)

	rootX := (int32)(c_root_x)

	rootY := (int32)(c_root_y)

	return rootX, rootY
}

// Check to see if a window is destroyed..
/*

C function : gdk_window_is_destroyed
*/
func (recv *Window) IsDestroyed() bool {
	retC := C.gdk_window_is_destroyed((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Changes the position of  @window in the Z-order (stacking order), so that
// it is above @sibling (if @above is %TRUE) or below @sibling (if @above is
// %FALSE).
//
// If @sibling is %NULL, then this either raises (if @above is %TRUE) or
// lowers the window.
//
// If @window is a toplevel, the window manager may choose to deny the
// request to move the window in the Z-order, gdk_window_restack() only
// requests the restack, does not guarantee it.
/*

C function : gdk_window_restack
*/
func (recv *Window) Restack(sibling *Window, above bool) {
	c_sibling := (*C.GdkWindow)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GdkWindow)(sibling.ToC())
	}

	c_above :=
		boolToGboolean(above)

	C.gdk_window_restack((*C.GdkWindow)(recv.native), c_sibling, c_above)

	return
}
