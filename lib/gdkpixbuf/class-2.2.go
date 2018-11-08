// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.2 gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

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

// SetOption is a wrapper around the C function gdk_pixbuf_set_option.
func (recv *Pixbuf) SetOption(key string, value string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.gdk_pixbuf_set_option((*C.GdkPixbuf)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// GetFormat is a wrapper around the C function gdk_pixbuf_loader_get_format.
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

// SetSize is a wrapper around the C function gdk_pixbuf_loader_set_size.
func (recv *PixbufLoader) SetSize(width int32, height int32) {
	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_pixbuf_loader_set_size((*C.GdkPixbufLoader)(recv.native), c_width, c_height)

	return
}

// Unsupported : PixbufSimpleAnimIter : no CType
