// This is a generated file - DO NOT EDIT
// +build gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter (GConverter*) for param converter

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter (GConverter*) for param converter

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no type generator for guint8 () for array param blob

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File (GFile*) for param file

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no type generator for guint8 () for array param bytes

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc (GReallocFunc) for param realloc_function

// GetInt64 is a wrapper around the C function g_settings_get_int64.
func (recv *Settings) GetInt64(key string) int64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int64((*C.GSettings)(recv.native), c_key)
	retGo := (int64)(retC)

	return retGo
}

// GetUint64 is a wrapper around the C function g_settings_get_uint64.
func (recv *Settings) GetUint64(key string) uint64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_uint64((*C.GSettings)(recv.native), c_key)
	retGo := (uint64)(retC)

	return retGo
}

// SetInt64 is a wrapper around the C function g_settings_set_int64.
func (recv *Settings) SetInt64(key string, value int64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint64)(value)

	retC := C.g_settings_set_int64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetUint64 is a wrapper around the C function g_settings_set_uint64.
func (recv *Settings) SetUint64(key string, value uint64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint64)(value)

	retC := C.g_settings_set_uint64((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no type generator for filename () for array param argv

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no type generator for utf8 (char*) for array param iconnames

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no type generator for gint (gint) for array param fds

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no type generator for gchar () for array param path

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no type generator for gchar () for array param path

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc (GVfsFileLookupFunc) for param uri_func

// UnregisterUriScheme is a wrapper around the C function g_vfs_unregister_uri_scheme.
func (recv *Vfs) UnregisterUriScheme(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.g_vfs_unregister_uri_scheme((*C.GVfs)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}
