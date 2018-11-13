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

// Creates a new pixbuf animation by loading an image from an resource.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set.
/*

C function

gdk_pixbuf_animation_new_from_resource
*/
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

// Creates a new animation by loading it from an input stream.
//
// The file format is detected automatically. If %NULL is returned, then
// @error will be set. The @cancellable can be used to abort the operation
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. Other possible errors are in
// the #GDK_PIXBUF_ERROR and %G_IO_ERROR domains.
//
// The stream is not closed.
/*

C function

gdk_pixbuf_animation_new_from_stream
*/
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
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes an asynchronous pixbuf animation creation operation started with
// gdk_pixbuf_animation_new_from_stream_async().
/*

C function

gdk_pixbuf_animation_new_from_stream_finish
*/
func PixbufAnimationNewFromStreamFinish(asyncResult *gio.AsyncResult) (*PixbufAnimation, error) {
	c_async_result := (*C.GAsyncResult)(asyncResult.ToC())

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_stream_finish(c_async_result, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
