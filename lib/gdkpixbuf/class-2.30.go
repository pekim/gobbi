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

// This will cause a pixbuf loader to parse a buffer inside a #GBytes
// for an image.  It will return %TRUE if the data was loaded successfully,
// and %FALSE if an error occurred.  In the latter case, the loader
// will be closed, and will not accept further writes. If %FALSE is
// returned, @error will be set to an error from the #GDK_PIXBUF_ERROR
// or #G_FILE_ERROR domains.
//
// See also: gdk_pixbuf_loader_write()
/*

C function

gdk_pixbuf_loader_write_bytes
*/
func (recv *PixbufLoader) WriteBytes(buffer *glib.Bytes) (bool, error) {
	c_buffer := (*C.GBytes)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GBytes)(buffer.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_write_bytes((*C.GdkPixbufLoader)(recv.native), c_buffer, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
