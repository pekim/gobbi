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

// LoadBytes is a wrapper around the C function g_file_load_bytes.
func (recv *File) LoadBytes(cancellable *Cancellable) (*glib.Bytes, string, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

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

// Unsupported : g_file_load_bytes_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// PeekPath is a wrapper around the C function g_file_peek_path.
func (recv *File) PeekPath() string {
	retC := C.g_file_peek_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
