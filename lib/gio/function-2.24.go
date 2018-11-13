// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

g_io_modules_scan_all_in_directory
*/
func IoModulesScanAllInDirectory(dirname string) {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	C.g_io_modules_scan_all_in_directory(c_dirname)

	return
}
