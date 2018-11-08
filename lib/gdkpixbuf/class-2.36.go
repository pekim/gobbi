// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no type generator for guint8 (guchar) for array param data

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no type generator for guint8 (guint8) for array param data

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no type generator for utf8 (char*) for array param data

// CopyOptions is a wrapper around the C function gdk_pixbuf_copy_options.
func (recv *Pixbuf) CopyOptions(destPixbuf *Pixbuf) bool {
	c_dest_pixbuf := (*C.GdkPixbuf)(destPixbuf.ToC())

	retC := C.gdk_pixbuf_copy_options((*C.GdkPixbuf)(recv.native), c_dest_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveOption is a wrapper around the C function gdk_pixbuf_remove_option.
func (recv *Pixbuf) RemoveOption(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_remove_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_save_to_streamv : unsupported parameter option_keys : no type generator for utf8 (char*) for array param option_keys

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter option_keys : no type generator for utf8 (gchar*) for array param option_keys

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// Unsupported : PixbufSimpleAnimIter : no CType
