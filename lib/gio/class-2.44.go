// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_enumerator_iterate : unsupported parameter out_info : record with indirection level of 2

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_input_stream_read_all_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

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

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

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

// Unsupported : g_list_store_splice : unsupported parameter additions : no param type

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// NetworkAddressNewLoopback is a wrapper around the C function g_network_address_new_loopback.
func NetworkAddressNewLoopback(port uint16) *NetworkAddress {
	c_port := (C.guint16)(port)

	retC := C.g_network_address_new_loopback(c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_output_stream_write_all_async : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_set_state_hint : unsupported parameter state_hint : Blacklisted record : GVariant

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

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

// SimpleIOStreamNew is a wrapper around the C function g_simple_io_stream_new.
func SimpleIOStreamNew(inputStream *InputStream, outputStream *OutputStream) *IOStream {
	c_input_stream := (*C.GInputStream)(inputStream.ToC())

	c_output_stream := (*C.GOutputStream)(outputStream.ToC())

	retC := C.g_simple_io_stream_new(c_input_stream, c_output_stream)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_socket_send_messages : unsupported parameter messages : no param type

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// GetCompleted is a wrapper around the C function g_task_get_completed.
func (recv *Task) GetCompleted() bool {
	retC := C.g_task_get_completed((*C.GTask)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type
