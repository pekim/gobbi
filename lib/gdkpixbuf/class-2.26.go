// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufNewFromResource is a wrapper around the C function gdk_pixbuf_new_from_resource.
func PixbufNewFromResource(resourcePath string) (*Pixbuf, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_resource(c_resource_path, &cThrowableError)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufNewFromResourceAtScale is a wrapper around the C function gdk_pixbuf_new_from_resource_at_scale.
func PixbufNewFromResourceAtScale(resourcePath string, width int32, height int32, preserveAspectRatio bool) (*Pixbuf, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_resource_at_scale(c_resource_path, c_width, c_height, c_preserve_aspect_ratio, &cThrowableError)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetByteLength is a wrapper around the C function gdk_pixbuf_get_byte_length.
func (recv *Pixbuf) GetByteLength() uint64 {
	retC := C.gdk_pixbuf_get_byte_length((*C.GdkPixbuf)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels_with_length : no return type
