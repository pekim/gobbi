// This is a generated file - DO NOT EDIT
// +build gio_2.54 gio_2.56

package gio

import "unsafe"

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

// UnixMountCopy is a wrapper around the C function g_unix_mount_copy.
func UnixMountCopy(mountEntry *UnixMountEntry) *UnixMountEntry {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_copy(c_mount_entry)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	return retGo
}
