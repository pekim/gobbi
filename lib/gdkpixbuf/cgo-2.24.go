// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufNewFromStreamFinish is a wrapper around the C function gdk_pixbuf_new_from_stream_finish.
func PixbufNewFromStreamFinish(asyncResult *gio.AsyncResult) (*Pixbuf, error) {
	c_async_result := (*C.GAsyncResult)(asyncResult.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream_finish(c_async_result, &cThrowableError)
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
