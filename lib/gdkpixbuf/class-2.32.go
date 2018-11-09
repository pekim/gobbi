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

	return retGo
}

// GetOptions is a wrapper around the C function gdk_pixbuf_get_options.
func (recv *Pixbuf) GetOptions() *glib.HashTable {
	retC := C.gdk_pixbuf_get_options((*C.GdkPixbuf)(recv.native))
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReadPixelBytes is a wrapper around the C function gdk_pixbuf_read_pixel_bytes.
func (recv *Pixbuf) ReadPixelBytes() *glib.Bytes {
	retC := C.gdk_pixbuf_read_pixel_bytes((*C.GdkPixbuf)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReadPixels is a wrapper around the C function gdk_pixbuf_read_pixels.
func (recv *Pixbuf) ReadPixels() uint8 {
	retC := C.gdk_pixbuf_read_pixels((*C.GdkPixbuf)(recv.native))
	retGo := (uint8)(retC)

	return retGo
}
