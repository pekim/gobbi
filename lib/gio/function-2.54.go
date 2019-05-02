// This is a generated file - DO NOT EDIT
// +build gio_2.54 gio_2.56

package gio

import (
	"C"
	"unsafe"
)

// UnixMountCopy is a wrapper around the C function g_unix_mount_copy.
func UnixMountCopy(mountEntry *UnixMountEntry) *UnixMountEntry {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_copy(c_mount_entry)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}
