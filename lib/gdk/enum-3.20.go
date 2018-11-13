// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Used in #GdkDragContext to the reason of a cancelled DND operation.
type DragCancelReason C.GdkDragCancelReason

const (
	// There is no suitable drop target.
	GDK_DRAG_CANCEL_NO_TARGET DragCancelReason = 0

	// Drag cancelled by the user
	GDK_DRAG_CANCEL_USER_CANCELLED DragCancelReason = 1

	// Unspecified error.
	GDK_DRAG_CANCEL_ERROR DragCancelReason = 2
)
