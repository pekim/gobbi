// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "C"

// GetByteLength is a wrapper around the C function gdk_pixbuf_get_byte_length.
func (recv *Pixbuf) GetByteLength() uint64 {
	retC := C.gdk_pixbuf_get_byte_length((*C.GdkPixbuf)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels_with_length : array return type :
