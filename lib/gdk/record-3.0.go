// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Copy is a wrapper around the C function gdk_rgba_copy.
func (recv *RGBA) Copy() *RGBA {
	retC := C.gdk_rgba_copy((*C.GdkRGBA)(recv.native))
	retGo := RGBANewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function gdk_rgba_equal.
func (recv *RGBA) Equal(p2 *RGBA) bool {
	c_p2 := (C.gconstpointer)(C.NULL)
	if p2 != nil {
		c_p2 = (C.gconstpointer)(p2.ToC())
	}

	retC := C.gdk_rgba_equal((C.gconstpointer)(recv.native), c_p2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gdk_rgba_free.
func (recv *RGBA) Free() {
	C.gdk_rgba_free((*C.GdkRGBA)(recv.native))

	return
}

// Hash is a wrapper around the C function gdk_rgba_hash.
func (recv *RGBA) Hash() uint32 {
	retC := C.gdk_rgba_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Parse is a wrapper around the C function gdk_rgba_parse.
func (recv *RGBA) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.gdk_rgba_parse((*C.GdkRGBA)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function gdk_rgba_to_string.
func (recv *RGBA) ToString() string {
	retC := C.gdk_rgba_to_string((*C.GdkRGBA)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
