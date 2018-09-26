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

// DragBeginFromPoint is a wrapper around the C function gdk_drag_begin_from_point.
func DragBeginFromPoint(window *Window, device *Device, targets *glib.List, xRoot int32, yRoot int32) *DragContext {
	c_window := (*C.GdkWindow)(window.ToC())

	c_device := (*C.GdkDevice)(device.ToC())

	c_targets := (*C.GList)(targets.ToC())

	c_x_root := (C.gint)(xRoot)

	c_y_root := (C.gint)(yRoot)

	retC := C.gdk_drag_begin_from_point(c_window, c_device, c_targets, c_x_root, c_y_root)
	retGo := DragContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_drag_drop_done : no return generator
