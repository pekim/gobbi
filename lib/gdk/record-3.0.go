// This is a generated file - DO NOT EDIT
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Makes a copy of a #GdkRGBA.
//
// The result must be freed through gdk_rgba_free().
/*

C function : gdk_rgba_copy
*/
func (recv *RGBA) Copy() *RGBA {
	retC := C.gdk_rgba_copy((*C.GdkRGBA)(recv.native))
	retGo := RGBANewFromC(unsafe.Pointer(retC))

	return retGo
}

// Compares two RGBA colors.
/*

C function : gdk_rgba_equal
*/
func (recv *RGBA) Equal(p2 uintptr) bool {
	c_p2 := (C.gconstpointer)(p2)

	retC := C.gdk_rgba_equal((C.gconstpointer)(recv.native), c_p2)
	retGo := retC == C.TRUE

	return retGo
}

// Frees a #GdkRGBA created with gdk_rgba_copy()
/*

C function : gdk_rgba_free
*/
func (recv *RGBA) Free() {
	C.gdk_rgba_free((*C.GdkRGBA)(recv.native))

	return
}

// A hash function suitable for using for a hash
// table that stores #GdkRGBAs.
/*

C function : gdk_rgba_hash
*/
func (recv *RGBA) Hash() uint32 {
	retC := C.gdk_rgba_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Parses a textual representation of a color, filling in
// the @red, @green, @blue and @alpha fields of the @rgba #GdkRGBA.
//
// The string can be either one of:
// - A standard name (Taken from the X11 rgb.txt file).
// - A hexadecimal value in the form “\#rgb”, “\#rrggbb”,
// “\#rrrgggbbb” or ”\#rrrrggggbbbb”
// - A RGB color in the form “rgb(r,g,b)” (In this case the color will
// have full opacity)
// - A RGBA color in the form “rgba(r,g,b,a)”
//
// Where “r”, “g”, “b” and “a” are respectively the red, green, blue and
// alpha color values. In the last two cases, r g and b are either integers
// in the range 0 to 255 or percentage values in the range 0% to 100%, and
// a is a floating point value in the range 0 to 1.
/*

C function : gdk_rgba_parse
*/
func (recv *RGBA) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.gdk_rgba_parse((*C.GdkRGBA)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

// Returns a textual specification of @rgba in the form
// `rgb (r, g, b)` or
// `rgba (r, g, b, a)`,
// where “r”, “g”, “b” and “a” represent the red, green,
// blue and alpha values respectively. r, g, and b are
// represented as integers in the range 0 to 255, and a
// is represented as floating point value in the range 0 to 1.
//
// These string forms are string forms those supported by
// the CSS3 colors module, and can be parsed by gdk_rgba_parse().
//
// Note that this string representation may lose some
// precision, since r, g and b are represented as 8-bit
// integers. If this is a concern, you should use a
// different representation.
/*

C function : gdk_rgba_to_string
*/
func (recv *RGBA) ToString() string {
	retC := C.gdk_rgba_to_string((*C.GdkRGBA)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
