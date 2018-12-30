// This is a generated file - DO NOT EDIT
// +build gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_application_add_main_option : unsupported parameter short_name : no type generator for gchar (char) for param short_name

// GetResourceBasePath is a wrapper around the C function g_application_get_resource_base_path.
func (recv *Application) GetResourceBasePath() string {
	retC := C.g_application_get_resource_base_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetResourceBasePath is a wrapper around the C function g_application_set_resource_base_path.
func (recv *Application) SetResourceBasePath(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.g_application_set_resource_base_path((*C.GApplication)(recv.native), c_resource_path)

	return
}

// DesktopAppInfoGetImplementations is a wrapper around the C function g_desktop_app_info_get_implementations.
func DesktopAppInfoGetImplementations(interface_ string) *glib.List {
	c_interface := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface))

	retC := C.g_desktop_app_info_get_implementations(c_interface)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}
