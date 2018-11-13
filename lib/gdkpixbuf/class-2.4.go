// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Creates a new pixbuf by loading an image from a file.
// The file format is detected automatically. If %NULL is returned, then
// @error will be set. Possible errors are in the #GDK_PIXBUF_ERROR and
// #G_FILE_ERROR domains.
//
// The image will be scaled to fit in the requested size, preserving
// the image's aspect ratio. Note that the returned pixbuf may be smaller
// than @width x @height, if the aspect ratio requires it. To load
// and image at the requested size, regardless of aspect ratio, use
// gdk_pixbuf_new_from_file_at_scale().
/*

C function : gdk_pixbuf_new_from_file_at_size
*/
func PixbufNewFromFileAtSize(filename string, width int32, height int32) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file_at_size(c_filename, c_width, c_height, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// Creates a new pixbuf loader object that always attempts to parse
// image data as if it were an image of mime type @mime_type, instead of
// identifying the type automatically. Useful if you want an error if
// the image isn't the expected mime type, for loading image formats
// that can't be reliably identified by looking at the data, or if
// the user manually forces a specific mime type.
//
// The list of supported mime types depends on what image loaders
// are installed, but typically "image/png", "image/jpeg", "image/gif",
// "image/tiff" and "image/x-xpixmap" are among the supported mime types.
// To obtain the full list of supported mime types, call
// gdk_pixbuf_format_get_mime_types() on each of the #GdkPixbufFormat
// structs returned by gdk_pixbuf_get_formats().
/*

C function : gdk_pixbuf_loader_new_with_mime_type
*/
func PixbufLoaderNewWithMimeType(mimeType string) (*PixbufLoader, error) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_mime_type(c_mime_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
