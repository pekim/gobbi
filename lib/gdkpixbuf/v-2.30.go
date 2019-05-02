// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

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

// WriteBytes is a wrapper around the C function gdk_pixbuf_loader_write_bytes.
func (recv *PixbufLoader) WriteBytes(buffer *glib.Bytes) (bool, error) {
	c_buffer := (*C.GBytes)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GBytes)(buffer.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_write_bytes((*C.GdkPixbufLoader)(recv.native), c_buffer, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}
