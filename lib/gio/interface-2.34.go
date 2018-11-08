// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_app_info_get_supported_types : no return type

// IsTagged is a wrapper around the C function g_async_result_is_tagged.
func (recv *AsyncResult) IsTagged(sourceTag uintptr) bool {
	c_source_tag := (C.gpointer)(sourceTag)

	retC := C.g_async_result_is_tagged((*C.GAsyncResult)(recv.native), c_source_tag)
	retGo := retC == C.TRUE

	return retGo
}

// LegacyPropagateError is a wrapper around the C function g_async_result_legacy_propagate_error.
func (recv *AsyncResult) LegacyPropagateError() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_async_result_legacy_propagate_error((*C.GAsyncResult)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_drive_get_symbolic_icon : no return generator

// Unsupported : g_file_delete_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_delete_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// Unsupported : g_mount_get_symbolic_icon : no return generator

// Unsupported : g_volume_get_symbolic_icon : no return generator
