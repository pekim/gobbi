// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Creates a new #GdkPixbuf out of in-memory readonly image data.
// Currently only RGB images with 8 bits per sample are supported.
// This is the #GBytes variant of gdk_pixbuf_new_from_data().
/*

C function : gdk_pixbuf_new_from_bytes
*/
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

	return retGo
}

// Returns a #GHashTable with a list of all the options that may have been
// attached to the @pixbuf when it was loaded, or that may have been
// attached by another function using gdk_pixbuf_set_option().
//
// See gdk_pixbuf_get_option() for more details.
/*

C function : gdk_pixbuf_get_options
*/
func (recv *Pixbuf) GetOptions() *glib.HashTable {
	retC := C.gdk_pixbuf_get_options((*C.GdkPixbuf)(recv.native))
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

/*

C function : gdk_pixbuf_read_pixel_bytes
*/
func (recv *Pixbuf) ReadPixelBytes() *glib.Bytes {
	retC := C.gdk_pixbuf_read_pixel_bytes((*C.GdkPixbuf)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_pixbuf_read_pixels
