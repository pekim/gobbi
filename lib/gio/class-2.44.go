// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

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

// Blacklisted : g_application_bind_busy_property

// Blacklisted : g_application_get_is_busy

// Blacklisted : g_application_unbind_busy_property

// Blacklisted : g_file_enumerator_iterate

// Unsupported : g_input_stream_read_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_input_stream_read_all_finish

// Blacklisted : g_list_store_new

// Blacklisted : g_list_store_append

// Blacklisted : g_list_store_insert

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// Blacklisted : g_list_store_remove

// Blacklisted : g_list_store_remove_all

// Unsupported : g_list_store_splice : unsupported parameter additions :

// Blacklisted : g_network_address_new_loopback

// Unsupported : g_output_stream_write_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_output_stream_write_all_finish

// Blacklisted : g_simple_action_set_state_hint

// SimpleIOStream is a wrapper around the C record GSimpleIOStream.
type SimpleIOStream struct {
	native *C.GSimpleIOStream
}

func SimpleIOStreamNewFromC(u unsafe.Pointer) *SimpleIOStream {
	c := (*C.GSimpleIOStream)(u)
	if c == nil {
		return nil
	}

	g := &SimpleIOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleIOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleIOStream with another SimpleIOStream, and returns true if they represent the same GObject.
func (recv *SimpleIOStream) Equals(other *SimpleIOStream) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to SimpleIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleIOStream.
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromC(object.ToC())
}

// Blacklisted : g_simple_io_stream_new

// Unsupported : g_socket_send_messages : unsupported parameter messages :

// Blacklisted : g_task_get_completed

// Blacklisted : g_unix_mount_monitor_get
