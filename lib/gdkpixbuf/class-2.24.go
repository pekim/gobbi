// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// gdk_pixbuf_new_from_stream_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
// gdk_pixbuf_new_from_stream_at_scale_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
// PixbufSaveToStreamFinish is a wrapper around the C function gdk_pixbuf_save_to_stream_finish.
func PixbufSaveToStreamFinish(asyncResult *gio.AsyncResult) (bool, error) {
	c_async_result := (*C.GAsyncResult)(asyncResult.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_save_to_stream_finish(c_async_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : gdk_pixbuf_save_to_stream_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
