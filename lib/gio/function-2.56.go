// This is a generated file - DO NOT EDIT
// +build gio_2.56

package gio

import "unsafe"

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

// UnixIsSystemDevicePath is a wrapper around the C function g_unix_is_system_device_path.
func UnixIsSystemDevicePath(devicePath string) bool {
	c_device_path := C.CString(devicePath)
	defer C.free(unsafe.Pointer(c_device_path))

	retC := C.g_unix_is_system_device_path(c_device_path)
	retGo := retC == C.TRUE

	return retGo
}

// UnixIsSystemFsType is a wrapper around the C function g_unix_is_system_fs_type.
func UnixIsSystemFsType(fsType string) bool {
	c_fs_type := C.CString(fsType)
	defer C.free(unsafe.Pointer(c_fs_type))

	retC := C.g_unix_is_system_fs_type(c_fs_type)
	retGo := retC == C.TRUE

	return retGo
}
