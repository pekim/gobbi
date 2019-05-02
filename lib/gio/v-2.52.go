// This is a generated file - DO NOT EDIT
// +build gio_2.52 gio_2.54 gio_2.56

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

const FILE_ATTRIBUTE_RECENT_MODIFIED string = C.G_FILE_ATTRIBUTE_RECENT_MODIFIED

// ContentTypeIsMimeType is a wrapper around the C function g_content_type_is_mime_type.
func ContentTypeIsMimeType(type_ string, mimeType string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.g_content_type_is_mime_type(c_type, c_mime_type)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountFor is a wrapper around the C function g_unix_mount_for.
func UnixMountFor(filePath string) (*UnixMountEntry, uint64) {
	c_file_path := C.CString(filePath)
	defer C.free(unsafe.Pointer(c_file_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_for(c_file_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}
