// This is a generated file - DO NOT EDIT
// +build gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// g_file_new_build_filename : unsupported parameter ... : varargs
// LoadBytes is a wrapper around the C function g_file_load_bytes.
func (recv *File) LoadBytes(cancellable *Cancellable) (*glib.Bytes, string, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes((*C.GFile)(recv.native), c_cancellable, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goError
}

// Unsupported : g_file_load_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadBytesFinish is a wrapper around the C function g_file_load_bytes_finish.
func (recv *File) LoadBytesFinish(result *AsyncResult) (*glib.Bytes, string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes_finish((*C.GFile)(recv.native), c_result, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goError
}

// PeekPath is a wrapper around the C function g_file_peek_path.
func (recv *File) PeekPath() string {
	retC := C.g_file_peek_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
