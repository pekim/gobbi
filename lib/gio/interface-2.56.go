// This is a generated file - DO NOT EDIT
// +build gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Loads the contents of @file and returns it as #GBytes.
//
// If @file is a resource:// based URI, the resulting bytes will reference the
// embedded resource instead of a copy. Otherwise, this is equivalent to calling
// g_file_load_contents() and g_bytes_new_take().
//
// For resources, @etag_out will be set to %NULL.
//
// The data contained in the resulting #GBytes is always zero-terminated, but
// this is not included in the #GBytes length. The resulting #GBytes should be
// freed with g_bytes_unref() when no longer in use.
/*

C function

g_file_load_bytes
*/
func (recv *File) LoadBytes(cancellable *Cancellable) (*glib.Bytes, string, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes((*C.GFile)(recv.native), c_cancellable, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goThrowableError
}

// Unsupported : g_file_load_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Completes an asynchronous request to g_file_load_bytes_async().
//
// For resources, @etag_out will be set to %NULL.
//
// The data contained in the resulting #GBytes is always zero-terminated, but
// this is not included in the #GBytes length. The resulting #GBytes should be
// freed with g_bytes_unref() when no longer in use.
//
// See g_file_load_bytes() for more information.
/*

C function

g_file_load_bytes_finish
*/
func (recv *File) LoadBytesFinish(result *AsyncResult) (*glib.Bytes, string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes_finish((*C.GFile)(recv.native), c_result, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goThrowableError
}

// Exactly like g_file_get_path(), but caches the result via
// g_object_set_qdata_full().  This is useful for example in C
// applications which mix `g_file_*` APIs with native ones.  It
// also avoids an extra duplicated string when possible, so will be
// generally more efficient.
//
// This call does no blocking I/O.
/*

C function

g_file_peek_path
*/
func (recv *File) PeekPath() string {
	retC := C.g_file_peek_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
