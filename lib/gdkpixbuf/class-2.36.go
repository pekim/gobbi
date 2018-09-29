// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.36

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

// CopyOptions is a wrapper around the C function gdk_pixbuf_copy_options.
func (recv *Pixbuf) CopyOptions(destPixbuf *Pixbuf) bool {
	c_dest_pixbuf := (*C.GdkPixbuf)(destPixbuf.ToC())

	retC := C.gdk_pixbuf_copy_options((*C.GdkPixbuf)(recv.native), c_dest_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels : no return type

// Unsupported : gdk_pixbuf_get_pixels_with_length : unsupported parameter length : no type generator for guint, guint*

// Unsupported : gdk_pixbuf_read_pixels : no return generator

// RemoveOption is a wrapper around the C function gdk_pixbuf_remove_option.
func (recv *Pixbuf) RemoveOption(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_remove_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_save : unsupported parameter error : record with indirection level of 2

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_stream : unsupported parameter error : record with indirection level of 2

// Unsupported : gdk_pixbuf_save_to_stream_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gdk_pixbuf_save_to_streamv : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_savev : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_loader_write : unsupported parameter buf : no param type

// Unsupported : PixbufSimpleAnimIter : no CType
