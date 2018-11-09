// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// WriteBytes is a wrapper around the C function gdk_pixbuf_loader_write_bytes.
func (recv *PixbufLoader) WriteBytes(buffer *glib.Bytes) (bool, error) {
	c_buffer := (*C.GBytes)(buffer.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_write_bytes((*C.GdkPixbufLoader)(recv.native), c_buffer, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
