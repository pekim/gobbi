// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Takes an existing pixbuf and checks for the presence of an
// associated "orientation" option, which may be provided by the
// jpeg loader (which reads the exif orientation tag) or the
// tiff loader (which reads the tiff orientation tag, and
// compensates it for the partial transforms performed by
// libtiff). If an orientation option/tag is present, the
// appropriate transform will be performed so that the pixbuf
// is oriented correctly.
/*

C function

gdk_pixbuf_apply_embedded_orientation
*/
func (recv *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	retC := C.gdk_pixbuf_apply_embedded_orientation((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}
