// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
/*

	void SocketListener_eventHandler();

	static gulong SocketListener_signal_connect_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "event", SocketListener_eventHandler, data);
	}

*/
import "C"

// Unsupported signal : unsupported parameter info : no type generator for AppInfo,

// Unsupported signal : unsupported parameter files : no param type

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// RegisterObjectWithClosures is a wrapper around the C function g_dbus_connection_register_object_with_closures.
func (recv *DBusConnection) RegisterObjectWithClosures(objectPath string, interfaceInfo *DBusInterfaceInfo, methodCallClosure *gobject.Closure, getPropertyClosure *gobject.Closure, setPropertyClosure *gobject.Closure) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_info := (*C.GDBusInterfaceInfo)(interfaceInfo.ToC())

	c_method_call_closure := (*C.GClosure)(methodCallClosure.ToC())

	c_get_property_closure := (*C.GClosure)(getPropertyClosure.ToC())

	c_set_property_closure := (*C.GClosure)(setPropertyClosure.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_register_object_with_closures((*C.GDBusConnection)(recv.native), c_object_path, c_interface_info, c_method_call_closure, c_get_property_closure, c_set_property_closure, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported signal : unsupported parameter changed_properties : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported signal : unsupported parameter changed_properties : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported signal : unsupported parameter file : no type generator for File,

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Unsupported signal : unsupported parameter choices : no param type

// Unsupported signal : unsupported parameter processes : no param type

// Unsupported signal : unsupported parameter keys : no param type

// Unsupported signal : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported signal : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported signal : unsupported parameter connectable : no type generator for SocketConnectable,

// SocketListenerSignalEventCallback is a callback function for a 'event' signal emitted from a SocketListener.
type SocketListenerSignalEventCallback func(event SocketListenerEvent, socket *Socket)

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter drive : no type generator for Drive,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter mount : no type generator for Mount,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,

// Unsupported signal : unsupported parameter volume : no type generator for Volume,
