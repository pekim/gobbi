// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// BindBusyProperty is a wrapper around the C function g_application_bind_busy_property.
func (recv *Application) BindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_bind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// GetIsBusy is a wrapper around the C function g_application_get_is_busy.
func (recv *Application) GetIsBusy() bool {
	retC := C.g_application_get_is_busy((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnbindBusyProperty is a wrapper around the C function g_application_unbind_busy_property.
func (recv *Application) UnbindBusyProperty(object uintptr, property string) {
	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_application_unbind_busy_property((*C.GApplication)(recv.native), c_object, c_property)

	return
}

// Unsupported : g_file_enumerator_iterate : unsupported parameter out_info : record with indirection level of 2

// Unsupported : g_input_stream_read_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_input_stream_read_all_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// ListStoreNew is a wrapper around the C function g_list_store_new.
func ListStoreNew(itemType gobject.Type) *ListStore {
	c_item_type := (C.GType)(itemType)

	retC := C.g_list_store_new(c_item_type)
	retGo := ListStoreNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function g_list_store_append.
func (recv *ListStore) Append(item uintptr) {
	c_item := (C.gpointer)(item)

	C.g_list_store_append((*C.GListStore)(recv.native), c_item)

	return
}

// Insert is a wrapper around the C function g_list_store_insert.
func (recv *ListStore) Insert(position uint32, item uintptr) {
	c_position := (C.guint)(position)

	c_item := (C.gpointer)(item)

	C.g_list_store_insert((*C.GListStore)(recv.native), c_position, c_item)

	return
}

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc (GCompareDataFunc) for param compare_func

// Remove is a wrapper around the C function g_list_store_remove.
func (recv *ListStore) Remove(position uint32) {
	c_position := (C.guint)(position)

	C.g_list_store_remove((*C.GListStore)(recv.native), c_position)

	return
}

// RemoveAll is a wrapper around the C function g_list_store_remove_all.
func (recv *ListStore) RemoveAll() {
	C.g_list_store_remove_all((*C.GListStore)(recv.native))

	return
}

// Unsupported : g_list_store_splice : unsupported parameter additions :

// NetworkAddressNewLoopback is a wrapper around the C function g_network_address_new_loopback.
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	c_port := (C.guint16)(port)

	retC := C.g_network_address_new_loopback(c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_output_stream_write_all_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_output_stream_write_all_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// Unsupported : g_simple_action_set_state_hint : unsupported parameter state_hint : Blacklisted record : GVariant

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

	return g
}

func (recv *SimpleIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStream upcasts to *IOStream
func (recv *SimpleIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimpleIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitary Object to SimpleIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleIOStream.
func CastToSimpleIOStream(object *gobject.Object) *SimpleIOStream {
	return SimpleIOStreamNewFromC(object.ToC())
}

// SimpleIOStreamNew is a wrapper around the C function g_simple_io_stream_new.
func SimpleIOStreamNew(inputStream *InputStream, outputStream *OutputStream) *SimpleIOStream {
	c_input_stream := (*C.GInputStream)(inputStream.ToC())

	c_output_stream := (*C.GOutputStream)(outputStream.ToC())

	retC := C.g_simple_io_stream_new(c_input_stream, c_output_stream)
	retGo := SimpleIOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_socket_send_messages : unsupported parameter messages :

// GetCompleted is a wrapper around the C function g_task_get_completed.
func (recv *Task) GetCompleted() bool {
	retC := C.g_task_get_completed((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
