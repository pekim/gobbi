// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Starts a drag and creates a new drag context for it.
//
// This function is called by the drag source.
/*

C function

gdk_drag_begin_from_point
*/
func DragBeginFromPoint(window *Window, device *Device, targets *glib.List, xRoot int32, yRoot int32) *DragContext {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	retC := C.gdk_drag_begin_from_point(c_window, c_device, c_targets, c_x_root, c_y_root)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inform GDK if the drop ended successfully. Passing %FALSE
// for @success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should
// be the last call before dropping the reference to the
// @context.
//
// The #GdkDragContext will only take the first gdk_drag_drop_done()
// call as effective, if this function is called multiple times,
// all subsequent calls will be ignored.
/*

C function

gdk_drag_drop_done
*/
func DragDropDone(context *DragContext, success bool) {
	c_context := (*C.GdkDragContext)(C.NULL)
	if context != nil {
		c_context = (*C.GdkDragContext)(context.ToC())
	}

	c_success :=
		boolToGboolean(success)

	C.gdk_drag_drop_done(c_context, c_success)

	return
}
