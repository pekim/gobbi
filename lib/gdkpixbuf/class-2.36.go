// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Copy the key/value pair options attached to a #GdkPixbuf to another.
// This is useful to keep original metadata after having manipulated
// a file. However be careful to remove metadata which you've already
// applied, such as the "orientation" option after rotating the image.
/*

C function

gdk_pixbuf_copy_options
*/
func (recv *Pixbuf) CopyOptions(destPixbuf *Pixbuf) bool {
	c_dest_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if destPixbuf != nil {
		c_dest_pixbuf = (*C.GdkPixbuf)(destPixbuf.ToC())
	}

	retC := C.gdk_pixbuf_copy_options((*C.GdkPixbuf)(recv.native), c_dest_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// Remove the key/value pair option attached to a #GdkPixbuf.
/*

C function

gdk_pixbuf_remove_option
*/
func (recv *Pixbuf) RemoveOption(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_remove_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_save_to_streamv : unsupported parameter option_keys :

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter option_keys :
