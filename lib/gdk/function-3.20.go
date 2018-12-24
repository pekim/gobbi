// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// DragBeginFromPoint is a wrapper around the C function gdk_drag_begin_from_point.
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

// DragDropDone is a wrapper around the C function gdk_drag_drop_done.
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
