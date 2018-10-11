// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

func (recv *Pixbuf) Object() *gobject.Object {}

// PixbufAnimationNewFromResource is a wrapper around the C function gdk_pixbuf_animation_new_from_resource.
func PixbufAnimationNewFromResource(resourcePath string) (*PixbufAnimation, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_resource(c_resource_path, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufAnimationNewFromStream is a wrapper around the C function gdk_pixbuf_animation_new_from_stream.
func PixbufAnimationNewFromStream(stream *gio.InputStream, cancellable *gio.Cancellable) (*PixbufAnimation, error) {
	c_stream := (*C.GInputStream)(stream.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_stream(c_stream, c_cancellable, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

func (recv *PixbufAnimation) Object() *gobject.Object {}

func (recv *PixbufAnimationIter) Object() *gobject.Object {}

func (recv *PixbufLoader) Object() *gobject.Object {}

func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {}

// Unsupported : PixbufSimpleAnimIter : no CType
