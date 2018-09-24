// This is a generated file - DO NOT EDIT
// +build gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_unix_mount_for : unsupported parameter time_read : no type generator for guint64, guint64*
