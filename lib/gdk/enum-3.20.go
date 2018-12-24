// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

type DragCancelReason C.GdkDragCancelReason

const (
	GDK_DRAG_CANCEL_NO_TARGET      DragCancelReason = 0
	GDK_DRAG_CANCEL_USER_CANCELLED DragCancelReason = 1
	GDK_DRAG_CANCEL_ERROR          DragCancelReason = 2
)
