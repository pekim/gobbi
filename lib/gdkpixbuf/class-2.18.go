// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import gobject "github.com/pekim/gobbi/lib/gobject"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

func (recv *Pixbuf) Object() *gobject.Object {}

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

func (recv *PixbufAnimation) Object() *gobject.Object {}

func (recv *PixbufAnimationIter) Object() *gobject.Object {}

func (recv *PixbufLoader) Object() *gobject.Object {}

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

func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {}

// Unsupported : PixbufSimpleAnimIter : no CType
