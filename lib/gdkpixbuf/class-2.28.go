// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufAnimationNewFromResource is a wrapper around the C function gdk_pixbuf_animation_new_from_resource.
func PixbufAnimationNewFromResource(resourcePath string) (*PixbufAnimation, error) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_resource(c_resource_path, &cThrowableError)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufAnimationNewFromStream is a wrapper around the C function gdk_pixbuf_animation_new_from_stream.
func PixbufAnimationNewFromStream(stream *gio.InputStream, cancellable *gio.Cancellable) (*PixbufAnimation, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_stream(c_stream, c_cancellable, &cThrowableError)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufAnimationNewFromStreamFinish is a wrapper around the C function gdk_pixbuf_animation_new_from_stream_finish.
func PixbufAnimationNewFromStreamFinish(asyncResult *gio.AsyncResult) (*PixbufAnimation, error) {
	c_async_result := (*C.GAsyncResult)(asyncResult.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_stream_finish(c_async_result, &cThrowableError)
	retGPointer := (C.gpointer)(retC)
	nonFloatingRef := C.g_object_is_floating(retGPointer) == C.FALSE
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	if nonFloatingRef {
		C.g_object_unref(retGPointer)
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// gdk_pixbuf_animation_new_from_stream_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback
