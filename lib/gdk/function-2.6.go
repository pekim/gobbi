// This is a generated file - DO NOT EDIT
// +build gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Returns whether the dropped data has been successfully
// transferred. This function is intended to be used while
// handling a %GDK_DROP_FINISHED event, its return value is
// meaningless at other times.
/*

C function

gdk_drag_drop_succeeded
*/
func DragDropSucceeded(context *DragContext) bool {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	retC := C.gdk_drag_drop_succeeded(c_context)
	retGo := retC == C.TRUE

	return retGo
}
