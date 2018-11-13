// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Returns %TRUE if the save option specified by @option_key is supported when
// saving a pixbuf using the module implementing @format.
// See gdk_pixbuf_save() for more information about option keys.
/*

C function

gdk_pixbuf_format_is_save_option_supported
*/
func (recv *PixbufFormat) IsSaveOptionSupported(optionKey string) bool {
	c_option_key := C.CString(optionKey)
	defer C.free(unsafe.Pointer(c_option_key))

	retC := C.gdk_pixbuf_format_is_save_option_supported((*C.GdkPixbufFormat)(recv.native), c_option_key)
	retGo := retC == C.TRUE

	return retGo
}
