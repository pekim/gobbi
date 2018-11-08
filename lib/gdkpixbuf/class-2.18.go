// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no type generator for guint8 (guchar) for array param data

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no type generator for guint8 (guint8) for array param data

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no type generator for utf8 (char*) for array param data

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult (GAsyncResult*) for param async_result

// GetLoop is a wrapper around the C function gdk_pixbuf_simple_anim_get_loop.
func (recv *PixbufSimpleAnim) GetLoop() bool {
	retC := C.gdk_pixbuf_simple_anim_get_loop((*C.GdkPixbufSimpleAnim)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLoop is a wrapper around the C function gdk_pixbuf_simple_anim_set_loop.
func (recv *PixbufSimpleAnim) SetLoop(loop bool) {
	c_loop :=
		boolToGboolean(loop)

	C.gdk_pixbuf_simple_anim_set_loop((*C.GdkPixbufSimpleAnim)(recv.native), c_loop)

	return
}

// Unsupported : PixbufSimpleAnimIter : no CType
