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

// Checks if @res has the given @source_tag (generally a function
// pointer indicating the function @res was created by).
/*

C function

g_async_result_is_tagged
*/
func (recv *AsyncResult) IsTagged(sourceTag uintptr) bool {
	c_source_tag := (C.gpointer)(sourceTag)

	retC := C.g_async_result_is_tagged((*C.GAsyncResult)(recv.native), c_source_tag)
	retGo := retC == C.TRUE

	return retGo
}

// If @res is a #GSimpleAsyncResult, this is equivalent to
// g_simple_async_result_propagate_error(). Otherwise it returns
// %FALSE.
//
// This can be used for legacy error handling in async *_finish()
// wrapper functions that traditionally handled #GSimpleAsyncResult
// error returns themselves rather than calling into the virtual method.
// This should not be used in new code; #GAsyncResult errors that are
// set by virtual methods should also be extracted by virtual methods,
// to enable subclasses to chain up correctly.
/*

C function

g_async_result_legacy_propagate_error
*/
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

// Gets the icon for @drive.
/*

C function

g_drive_get_symbolic_icon
*/
func (recv *Drive) GetSymbolicIcon() *Icon {
	retC := C.g_drive_get_symbolic_icon((*C.GDrive)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_delete_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes deleting a file started with g_file_delete_async().
/*

C function

g_file_delete_finish
*/
func (recv *File) DeleteFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_delete_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the symbolic icon for @mount.
/*

C function

g_mount_get_symbolic_icon
*/
func (recv *Mount) GetSymbolicIcon() *Icon {
	retC := C.g_mount_get_symbolic_icon((*C.GMount)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the symbolic icon for @volume.
/*

C function

g_volume_get_symbolic_icon
*/
func (recv *Volume) GetSymbolicIcon() *Icon {
	retC := C.g_volume_get_symbolic_icon((*C.GVolume)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}
