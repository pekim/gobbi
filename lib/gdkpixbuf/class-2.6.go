// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	"unsafe"
)

// Flip is a wrapper around the C function gdk_pixbuf_flip.
func (recv *Pixbuf) Flip(horizontal bool) *Pixbuf {
	c_horizontal :=
		boolToGboolean(horizontal)

	retC := C.gdk_pixbuf_flip((*C.GdkPixbuf)(recv.native), c_horizontal)
	var retGo (*Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RotateSimple is a wrapper around the C function gdk_pixbuf_rotate_simple.
func (recv *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf {
	c_angle := (C.GdkPixbufRotation)(angle)

	retC := C.gdk_pixbuf_rotate_simple((*C.GdkPixbuf)(recv.native), c_angle)
	var retGo (*Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
