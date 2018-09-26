// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// DragDropSucceeded is a wrapper around the C function gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) bool {
	c_context := (*C.GdkDragContext)(context.ToC())

	retC := C.gdk_drag_drop_succeeded(c_context)
	retGo := retC == C.TRUE

	return retGo
}
