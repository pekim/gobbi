// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// CursorNewForDisplay is a wrapper around the C function gdk_cursor_new_for_display.
func CursorNewForDisplay(display *Display, cursorType CursorType) *Cursor {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_cursor_type := (C.GdkCursorType)(cursorType)

	retC := C.gdk_cursor_new_for_display(c_display, c_cursor_type)
	retGo := CursorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter protocol : GdkDragProtocol* with indirection level of 1

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter list : output array param list
