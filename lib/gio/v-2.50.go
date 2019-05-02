// This is a generated file - DO NOT EDIT
// +build gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// GetInt64 is a wrapper around the C function g_settings_get_int64.
func (recv *Settings) GetInt64(key string) int64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int64((*C.GSettings)(recv.native), c_key)
	retGo := (int64)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_settings_get_uint64.
func (recv *Settings) GetUint64(key string) uint64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_uint64((*C.GSettings)(recv.native), c_key)
	retGo := (uint64)(retC)

	return retGo
}

// SetInt64 is a wrapper around the C function g_settings_set_int64.
func (recv *Settings) SetInt64(key string, value int64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	retC := C.g_settings_set_int64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetUint64 is a wrapper around the C function g_settings_set_uint64.
func (recv *Settings) SetUint64(key string, value uint64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	retC := C.g_settings_set_uint64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc (GVfsFileLookupFunc) for param uri_func

// UnregisterUriScheme is a wrapper around the C function g_vfs_unregister_uri_scheme.
func (recv *Vfs) UnregisterUriScheme(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.g_vfs_unregister_uri_scheme((*C.GVfs)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_app_info_launch_default_for_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// g_app_info_launch_default_for_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// AppInfoLaunchDefaultForUriFinish is a wrapper around the C function g_app_info_launch_default_for_uri_finish.
func AppInfoLaunchDefaultForUriFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_default_for_uri_finish(c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IsRemovable is a wrapper around the C function g_drive_is_removable.
func (recv *Drive) IsRemovable() bool {
	retC := C.g_drive_is_removable((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
