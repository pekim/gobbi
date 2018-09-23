// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// ContentTypeGetGenericIconName is a wrapper around the C function g_content_type_get_generic_icon_name.
func ContentTypeGetGenericIconName(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_generic_icon_name(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_content_type_get_symbolic_icon : no return generator

// Unsupported : g_pollable_source_new_full : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_pollable_stream_read : unsupported parameter stream : no type generator for InputStream, GInputStream*

// Unsupported : g_pollable_stream_write : unsupported parameter stream : no type generator for OutputStream, GOutputStream*

// Unsupported : g_pollable_stream_write_all : unsupported parameter stream : no type generator for OutputStream, GOutputStream*

// Unsupported : g_unix_mount_guess_symbolic_icon : no return generator
