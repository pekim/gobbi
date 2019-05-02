// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	"unsafe"
)

// ApplyEmbeddedOrientation is a wrapper around the C function gdk_pixbuf_apply_embedded_orientation.
func (recv *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	retC := C.gdk_pixbuf_apply_embedded_orientation((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}
