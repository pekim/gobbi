// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// PollableSourceNewFull is a wrapper around the C function g_pollable_source_new_full.
func PollableSourceNewFull(pollableStream uintptr, childSource *glib.Source, cancellable *Cancellable) *glib.Source {
	c_pollable_stream := (C.gpointer)(pollableStream)

	c_child_source := (*C.GSource)(childSource.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_pollable_source_new_full(c_pollable_stream, c_child_source, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_pollable_stream_read : unsupported parameter buffer : no param type

// Unsupported : g_pollable_stream_write : unsupported parameter buffer : no param type

// Unsupported : g_pollable_stream_write_all : unsupported parameter buffer : no param type

// Unsupported : g_unix_mount_guess_symbolic_icon : no return generator
