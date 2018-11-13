// This is a generated file - DO NOT EDIT
// +build gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for 64-bit integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a int64 type in the schema for @settings.
/*

C function : g_settings_get_int64
*/
func (recv *Settings) GetInt64(key string) int64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int64((*C.GSettings)(recv.native), c_key)
	retGo := (int64)(retC)

	return retGo
}

// Gets the value that is stored at @key in @settings.
//
// A convenience variant of g_settings_get() for 64-bit unsigned
// integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a uint64 type in the schema for @settings.
/*

C function : g_settings_get_uint64
*/
func (recv *Settings) GetUint64(key string) uint64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_uint64((*C.GSettings)(recv.native), c_key)
	retGo := (uint64)(retC)

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for 64-bit integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a int64 type in the schema for @settings.
/*

C function : g_settings_set_int64
*/
func (recv *Settings) SetInt64(key string, value int64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	retC := C.g_settings_set_int64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Sets @key in @settings to @value.
//
// A convenience variant of g_settings_set() for 64-bit unsigned
// integers.
//
// It is a programmer error to give a @key that isn't specified as
// having a uint64 type in the schema for @settings.
/*

C function : g_settings_set_uint64
*/
func (recv *Settings) SetUint64(key string, value uint64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	retC := C.g_settings_set_uint64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc (GVfsFileLookupFunc) for param uri_func

// Unregisters the URI handler for @scheme previously registered with
// g_vfs_register_uri_scheme().
/*

C function : g_vfs_unregister_uri_scheme
*/
func (recv *Vfs) UnregisterUriScheme(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.g_vfs_unregister_uri_scheme((*C.GVfs)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}
