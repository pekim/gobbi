// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.2 gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gdk_pixbuf_set_option

// Obtains the available information about the format of the
// currently loading image file.
/*

C function : gdk_pixbuf_loader_get_format
*/
func (recv *PixbufLoader) GetFormat() *PixbufFormat {
	retC := C.gdk_pixbuf_loader_get_format((*C.GdkPixbufLoader)(recv.native))
	var retGo (*PixbufFormat)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufFormatNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Causes the image to be scaled while it is loaded. The desired
// image size can be determined relative to the original size of
// the image by calling gdk_pixbuf_loader_set_size() from a
// signal handler for the ::size-prepared signal.
//
// Attempts to set the desired image size  are ignored after the
// emission of the ::size-prepared signal.
/*

C function : gdk_pixbuf_loader_set_size
*/
func (recv *PixbufLoader) SetSize(width int32, height int32) {
	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_pixbuf_loader_set_size((*C.GdkPixbufLoader)(recv.native), c_width, c_height)

	return
}
