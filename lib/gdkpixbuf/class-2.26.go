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

// Creates a new pixbuf by loading an image from an resource.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set.
/*

C function : gdk_pixbuf_new_from_resource
*/
func PixbufNewFromResource(resourcePath string) (*Pixbuf, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_resource(c_resource_path, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new pixbuf by loading an image from an resource.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set.
//
// The image will be scaled to fit in the requested size, optionally
// preserving the image's aspect ratio. When preserving the aspect ratio,
// a @width of -1 will cause the image to be scaled to the exact given
// height, and a @height of -1 will cause the image to be scaled to the
// exact given width. When not preserving aspect ratio, a @width or
// @height of -1 means to not scale the image at all in that dimension.
//
// The stream is not closed.
/*

C function : gdk_pixbuf_new_from_resource_at_scale
*/
func PixbufNewFromResourceAtScale(resourcePath string, width int32, height int32, preserveAspectRatio bool) (*Pixbuf, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_resource_at_scale(c_resource_path, c_width, c_height, c_preserve_aspect_ratio, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns the length of the pixel data, in bytes.
/*

C function : gdk_pixbuf_get_byte_length
*/
func (recv *Pixbuf) GetByteLength() uint64 {
	retC := C.gdk_pixbuf_get_byte_length((*C.GdkPixbuf)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels_with_length : no return type
