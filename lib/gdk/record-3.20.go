// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Equal is a wrapper around the C function gdk_rectangle_equal.
func (recv *Rectangle) Equal(rect2 *Rectangle) bool {
	c_rect2 := (*C.GdkRectangle)(C.NULL)
	if rect2 != nil {
		c_rect2 = (*C.GdkRectangle)(rect2.ToC())
	}

	retC := C.gdk_rectangle_equal((*C.GdkRectangle)(recv.native), c_rect2)
	retGo := retC == C.TRUE

	return retGo
}
