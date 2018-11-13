// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.2 gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Returns a description of the format.
/*

C function

gdk_pixbuf_format_get_description
*/
func (recv *PixbufFormat) GetDescription() string {
	retC := C.gdk_pixbuf_format_get_description((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_format_get_extensions : no return type

// Unsupported : gdk_pixbuf_format_get_mime_types : no return type

// Returns the name of the format.
/*

C function

gdk_pixbuf_format_get_name
*/
func (recv *PixbufFormat) GetName() string {
	retC := C.gdk_pixbuf_format_get_name((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns whether pixbufs can be saved in the given format.
/*

C function

gdk_pixbuf_format_is_writable
*/
func (recv *PixbufFormat) IsWritable() bool {
	retC := C.gdk_pixbuf_format_is_writable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
