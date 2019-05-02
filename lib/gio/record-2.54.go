// This is a generated file - DO NOT EDIT
// +build gio_2.54 gio_2.56

package gio

import (
	"C"
	"unsafe"
)

// Copy is a wrapper around the C function g_unix_mount_point_copy.
func (recv *UnixMountPoint) Copy() *UnixMountPoint {
	retC := C.g_unix_mount_point_copy((*C.GUnixMountPoint)(recv.native))
	retGo := UnixMountPointNewFromC(unsafe.Pointer(retC))

	return retGo
}
