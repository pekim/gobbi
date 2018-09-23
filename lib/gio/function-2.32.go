// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// Unsupported : g_file_new_tmp : unsupported parameter iostream : no type generator for FileIOStream, GFileIOStream**

// Unsupported : g_network_monitor_get_default : no return generator

// ResourceErrorQuark is a wrapper around the C function g_resource_error_quark.
func ResourceErrorQuark() glib.Quark {
	retC := C.g_resource_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// ResourceLoad is a wrapper around the C function g_resource_load.
func ResourceLoad(filename string) (*Resource, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_resource_load(c_filename, &cThrowableError)
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resources_enumerate_children : unsupported parameter lookup_flags : no type generator for ResourceLookupFlags, GResourceLookupFlags

// Unsupported : g_resources_get_info : unsupported parameter lookup_flags : no type generator for ResourceLookupFlags, GResourceLookupFlags

// Unsupported : g_resources_lookup_data : unsupported parameter lookup_flags : no type generator for ResourceLookupFlags, GResourceLookupFlags

// Unsupported : g_resources_open_stream : unsupported parameter lookup_flags : no type generator for ResourceLookupFlags, GResourceLookupFlags

// Unsupported : g_resources_register : no return generator

// Unsupported : g_resources_unregister : no return generator

// SettingsSchemaSourceGetDefault is a wrapper around the C function g_settings_schema_source_get_default.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_get_default()
	retGo := SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}
