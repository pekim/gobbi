// This is a generated file - DO NOT EDIT
// +build gio_2.56

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

// Determines if @device_path is considered a block device path which is only
// used in implementation of the OS. This is primarily used for hiding
// mounted volumes that are intended as APIs for programs to read, and system
// administrators at a shell; rather than something that should, for example,
// appear in a GUI. For example, the Linux `/proc` filesystem.
//
// The list of device paths considered ‘system’ ones may change over time.
/*

C function

g_unix_is_system_device_path
*/
func UnixIsSystemDevicePath(devicePath string) bool {
	c_device_path := C.CString(devicePath)
	defer C.free(unsafe.Pointer(c_device_path))

	retC := C.g_unix_is_system_device_path(c_device_path)
	retGo := retC == C.TRUE

	return retGo
}

// Determines if @fs_type is considered a type of file system which is only
// used in implementation of the OS. This is primarily used for hiding
// mounted volumes that are intended as APIs for programs to read, and system
// administrators at a shell; rather than something that should, for example,
// appear in a GUI. For example, the Linux `/proc` filesystem.
//
// The list of file system types considered ‘system’ ones may change over time.
/*

C function

g_unix_is_system_fs_type
*/
func UnixIsSystemFsType(fsType string) bool {
	c_fs_type := C.CString(fsType)
	defer C.free(unsafe.Pointer(c_fs_type))

	retC := C.g_unix_is_system_fs_type(c_fs_type)
	retGo := retC == C.TRUE

	return retGo
}
