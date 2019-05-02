// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.32 gdkpixbuf_2.36

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

// PixbufNewFromBytes is a wrapper around the C function gdk_pixbuf_new_from_bytes.
func PixbufNewFromBytes(data *glib.Bytes, colorspace Colorspace, hasAlpha bool, bitsPerSample int32, width int32, height int32, rowstride int32) *Pixbuf {
	c_data := (*C.GBytes)(C.NULL)
	if data != nil {
		c_data = (*C.GBytes)(data.ToC())
	}

	c_colorspace := (C.GdkColorspace)(colorspace)

	c_has_alpha :=
		boolToGboolean(hasAlpha)

	c_bits_per_sample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_rowstride := (C.int)(rowstride)

	retC := C.gdk_pixbuf_new_from_bytes(c_data, c_colorspace, c_has_alpha, c_bits_per_sample, c_width, c_height, c_rowstride)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}
