// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

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

// Unsupported : gdk_pixbuf_get_pixels : no return type

// Unsupported : gdk_pixbuf_get_pixels_with_length : unsupported parameter length : no type generator for guint, guint*

// Unsupported : gdk_pixbuf_read_pixels : no return generator

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

// PixbufSimpleAnimNew is a wrapper around the C function gdk_pixbuf_simple_anim_new.
func PixbufSimpleAnimNew(width int32, height int32, rate float32) *PixbufSimpleAnim {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_rate := (C.gfloat)(rate)

	retC := C.gdk_pixbuf_simple_anim_new(c_width, c_height, c_rate)
	retGo := PixbufSimpleAnimNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddFrame is a wrapper around the C function gdk_pixbuf_simple_anim_add_frame.
func (recv *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gdk_pixbuf_simple_anim_add_frame((*C.GdkPixbufSimpleAnim)(recv.native), c_pixbuf)

	return
}

// Unsupported : PixbufSimpleAnimIter : no CType
