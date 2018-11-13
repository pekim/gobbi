// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Creates a new pixbuf by loading an image from an input stream.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set. The @cancellable can be used to abort the operation
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
// the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.
//
// The stream is not closed.
/*

C function : gdk_pixbuf_new_from_stream
*/
func PixbufNewFromStream(stream *gio.InputStream, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream(c_stream, c_cancellable, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new pixbuf by loading an image from an input stream.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set. The @cancellable can be used to abort the operation
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
// the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.
//
// The image will be scaled to fit in the requested size, optionally
// preserving the image's aspect ratio.
//
// When preserving the aspect ratio, a @width of -1 will cause the image to be
// scaled to the exact given height, and a @height of -1 will cause the image
// to be scaled to the exact given width. If both @width and @height are
// given, this function will behave as if the smaller of the two values
// is passed as -1.
//
// When not preserving aspect ratio, a @width or @height of -1 means to not
// scale the image at all in that dimension.
//
// The stream is not closed.
/*

C function : gdk_pixbuf_new_from_stream_at_scale
*/
func PixbufNewFromStreamAtScale(stream *gio.InputStream, width int32, height int32, preserveAspectRatio bool, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream_at_scale(c_stream, c_width, c_height, c_preserve_aspect_ratio, c_cancellable, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_save_to_stream : unsupported parameter ... : varargs
