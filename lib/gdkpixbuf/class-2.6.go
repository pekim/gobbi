// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufNewFromFileAtScale is a wrapper around the C function gdk_pixbuf_new_from_file_at_scale.
func PixbufNewFromFileAtScale(filename string, width int32, height int32, preserveAspectRatio bool) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file_at_scale(c_filename, c_width, c_height, c_preserve_aspect_ratio, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

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
