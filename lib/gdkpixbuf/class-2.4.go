// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufNewFromFileAtSize is a wrapper around the C function gdk_pixbuf_new_from_file_at_size.
func PixbufNewFromFileAtSize(filename string, width int32, height int32) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file_at_size(c_filename, c_width, c_height, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PixbufGetFileInfo is a wrapper around the C function gdk_pixbuf_get_file_info.
func PixbufGetFileInfo(filename string) (*PixbufFormat, int32, int32) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var c_width C.gint

	var c_height C.gint

	retC := C.gdk_pixbuf_get_file_info(c_filename, &c_width, &c_height)
	var retGo (*PixbufFormat)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufFormatNewFromC(unsafe.Pointer(retC))
	}

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// PixbufLoaderNewWithMimeType is a wrapper around the C function gdk_pixbuf_loader_new_with_mime_type.
func PixbufLoaderNewWithMimeType(mimeType string) (*PixbufLoader, error) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_mime_type(c_mime_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
