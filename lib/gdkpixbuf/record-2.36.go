// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_format_get_extensions : no return type

// Unsupported : gdk_pixbuf_format_get_mime_types : no return type

// IsSaveOptionSupported is a wrapper around the C function gdk_pixbuf_format_is_save_option_supported.
func (recv *PixbufFormat) IsSaveOptionSupported(optionKey string) bool {
	c_option_key := C.CString(optionKey)
	defer C.free(unsafe.Pointer(c_option_key))

	retC := C.gdk_pixbuf_format_is_save_option_supported((*C.GdkPixbufFormat)(recv.native), c_option_key)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : GdkPixdata
