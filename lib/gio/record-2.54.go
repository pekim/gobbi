// This is a generated file - DO NOT EDIT
// +build gio_2.54 gio_2.56

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

// Makes a copy of @mount_point.
/*

C function

g_unix_mount_point_copy
*/
func (recv *UnixMountPoint) Copy() *UnixMountPoint {
	retC := C.g_unix_mount_point_copy((*C.GUnixMountPoint)(recv.native))
	retGo := UnixMountPointNewFromC(unsafe.Pointer(retC))

	return retGo
}
