// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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
// #include <stdlib.h>
import "C"

// Unsupported : g_dbus_gvalue_to_gvariant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_dbus_gvariant_to_gvalue : unsupported parameter value : Blacklisted record : GVariant

// IoModulesLoadAllInDirectoryWithScope is a wrapper around the C function g_io_modules_load_all_in_directory_with_scope.
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

// IoModulesScanAllInDirectoryWithScope is a wrapper around the C function g_io_modules_scan_all_in_directory_with_scope.
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
