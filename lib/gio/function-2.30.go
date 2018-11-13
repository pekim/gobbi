// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_dbus_gvalue_to_gvariant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_dbus_gvariant_to_gvalue : unsupported parameter value : Blacklisted record : GVariant

// Loads all the modules in the specified directory.
//
// If don't require all modules to be initialized (and thus registering
// all gtypes) then you can use g_io_modules_scan_all_in_directory()
// which allows delayed/lazy loading of modules.
/*

C function

g_io_modules_load_all_in_directory_with_scope
*/
func IoModulesLoadAllInDirectoryWithScope(dirname string, scope *IOModuleScope) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	c_scope := (*C.GIOModuleScope)(C.NULL)
	if scope != nil {
		c_scope = (*C.GIOModuleScope)(scope.ToC())
	}

	retC := C.g_io_modules_load_all_in_directory_with_scope(c_dirname, c_scope)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Scans all the modules in the specified directory, ensuring that
// any extension point implemented by a module is registered.
//
// This may not actually load and initialize all the types in each
// module, some modules may be lazily loaded and initialized when
// an extension point it implementes is used with e.g.
// g_io_extension_point_get_extensions() or
// g_io_extension_point_get_extension_by_name().
//
// If you need to guarantee that all types are loaded in all the modules,
// use g_io_modules_load_all_in_directory().
/*

C function

g_io_modules_scan_all_in_directory_with_scope
*/
func IoModulesScanAllInDirectoryWithScope(dirname string, scope *IOModuleScope) {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	c_scope := (*C.GIOModuleScope)(C.NULL)
	if scope != nil {
		c_scope = (*C.GIOModuleScope)(scope.ToC())
	}

	C.g_io_modules_scan_all_in_directory_with_scope(c_dirname, c_scope)

	return
}

// Creates a new #GTlsFileDatabase which uses anchor certificate authorities
// in @anchors to verify certificate chains.
//
// The certificates in @anchors must be PEM encoded.
/*

C function

g_tls_file_database_new
*/
func TlsFileDatabaseNew(anchors string) (*TlsFileDatabase, error) {
	c_anchors := C.CString(anchors)
	defer C.free(unsafe.Pointer(c_anchors))

	var cThrowableError *C.GError

	retC := C.g_tls_file_database_new(c_anchors, &cThrowableError)
	retGo := TlsFileDatabaseNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
