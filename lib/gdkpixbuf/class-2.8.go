// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

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

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

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
