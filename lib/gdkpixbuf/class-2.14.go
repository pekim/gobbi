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

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// PixbufNewFromStream is a wrapper around the C function gdk_pixbuf_new_from_stream.
func PixbufNewFromStream(stream *gio.InputStream, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(stream.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream(c_stream, c_cancellable, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufNewFromStreamAtScale is a wrapper around the C function gdk_pixbuf_new_from_stream_at_scale.
func PixbufNewFromStreamAtScale(stream *gio.InputStream, width int32, height int32, preserveAspectRatio bool, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(stream.ToC())

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream_at_scale(c_stream, c_width, c_height, c_preserve_aspect_ratio, c_cancellable, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data :

// Unsupported : gdk_pixbuf_save_to_stream : unsupported parameter error : record with indirection level of 2

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// Unsupported : PixbufSimpleAnimIter : no CType
