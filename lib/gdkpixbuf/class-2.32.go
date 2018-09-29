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
	c_data := (*C.GBytes)(data.ToC())

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

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

// GetOptions is a wrapper around the C function gdk_pixbuf_get_options.
func (recv *Pixbuf) GetOptions() *glib.HashTable {
	retC := C.gdk_pixbuf_get_options((*C.GdkPixbuf)(recv.native))
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels : no return type

// Unsupported : gdk_pixbuf_get_pixels_with_length : unsupported parameter length : no type generator for guint, guint*

// ReadPixelBytes is a wrapper around the C function gdk_pixbuf_read_pixel_bytes.
func (recv *Pixbuf) ReadPixelBytes() *glib.Bytes {
	retC := C.gdk_pixbuf_read_pixel_bytes((*C.GdkPixbuf)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : PixbufSimpleAnimIter : no CType
