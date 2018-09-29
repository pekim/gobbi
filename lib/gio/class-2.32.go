// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_app_launch_context_get_display : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Unsupported : g_app_launch_context_get_environment : no return type

// Unsupported : g_app_launch_context_get_startup_notify_id : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Setenv is a wrapper around the C function g_app_launch_context_setenv.
func (recv *AppLaunchContext) Setenv(variable string, value string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_app_launch_context_setenv((*C.GAppLaunchContext)(recv.native), c_variable, c_value)

	return
}

// Unsetenv is a wrapper around the C function g_app_launch_context_unsetenv.
func (recv *AppLaunchContext) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_app_launch_context_unsetenv((*C.GAppLaunchContext)(recv.native), c_variable)

	return
}

// Unsupported : g_application_add_main_option : unsupported parameter short_name : no type generator for gchar, char

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries : no param type

// Unsupported : g_application_open : unsupported parameter files : no param type

// Quit is a wrapper around the C function g_application_quit.
func (recv *Application) Quit() {
	C.g_application_quit((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_run : unsupported parameter argv : no param type

// Unsupported : g_application_set_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// SetDefault is a wrapper around the C function g_application_set_default.
func (recv *Application) SetDefault() {
	C.g_application_set_default((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_command_line_create_file_for_arg : no return generator

// Unsupported : g_application_command_line_get_arguments : unsupported parameter argc : no type generator for gint, int*

// Unsupported : g_application_command_line_get_environ : no return type

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// Unsupported : g_application_command_line_print : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_printerr : unsupported parameter ... : varargs

// Unsupported : g_buffered_input_stream_fill_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_buffered_input_stream_fill_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_buffered_input_stream_peek : unsupported parameter buffer : no param type

// Unsupported : g_buffered_input_stream_peek_buffer : unsupported parameter count : no type generator for gsize, gsize*

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback, GCallback

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_input_stream_get_converter : no return generator

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_get_converter : no return generator

// Unsupported : g_credentials_get_unix_pid : no return generator

// Unsupported : g_credentials_get_unix_user : no return generator

// Unsupported : g_credentials_set_unix_user : unsupported parameter uid : no type generator for guint, uid_t

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction, GDBusMessageFilterFunction

// Unsupported : g_dbus_connection_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list_finish : unsupported parameter out_fd_list : record with indirection level of 2

// Unsupported : g_dbus_connection_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_close : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_close_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_emit_signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_export_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// ExportMenuModel is a wrapper around the C function g_dbus_connection_export_menu_model.
func (recv *DBusConnection) ExportMenuModel(objectPath string, menu *MenuModel) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_menu := (*C.GMenuModel)(menu.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_export_menu_model((*C.GDBusConnection)(recv.native), c_object_path, c_menu, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_flush_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_send_message : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_send_message_with_reply_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_send_message_with_reply_sync : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback, GDBusSignalCallback

// UnexportActionGroup is a wrapper around the C function g_dbus_connection_unexport_action_group.
func (recv *DBusConnection) UnexportActionGroup(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_action_group((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// UnexportMenuModel is a wrapper around the C function g_dbus_connection_unexport_menu_model.
func (recv *DBusConnection) UnexportMenuModel(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_menu_model((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// GetConnections is a wrapper around the C function g_dbus_interface_skeleton_get_connections.
func (recv *DBusInterfaceSkeleton) GetConnections() *glib.List {
	retC := C.g_dbus_interface_skeleton_get_connections((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_interface_skeleton_get_properties : return type : Blacklisted record : GVariant

// HasConnection is a wrapper around the C function g_dbus_interface_skeleton_has_connection.
func (recv *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	c_connection := (*C.GDBusConnection)(connection.ToC())

	retC := C.g_dbus_interface_skeleton_has_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)
	retGo := retC == C.TRUE

	return retGo
}

// UnexportFromConnection is a wrapper around the C function g_dbus_interface_skeleton_unexport_from_connection.
func (recv *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	c_connection := (*C.GDBusConnection)(connection.ToC())

	C.g_dbus_interface_skeleton_unexport_from_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)

	return
}

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_message_get_body : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_message_get_header : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_message_get_header_fields : no return type

// Unsupported : g_dbus_message_new_method_error : unsupported parameter ... : varargs

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_dbus_message_set_body : unsupported parameter body : Blacklisted record : GVariant

// Unsupported : g_dbus_message_set_header : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_dbus_message_to_blob : unsupported parameter out_size : no type generator for gsize, gsize*

// Unsupported : g_dbus_method_invocation_get_parameters : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_method_invocation_return_error : unsupported parameter ... : varargs

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_dbus_method_invocation_return_value : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_method_invocation_return_value_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_finish : unsupported parameter out_fd_list : record with indirection level of 2

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property_names : no return type

// Unsupported : g_dbus_proxy_set_cached_property : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_data_input_stream_read_line : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_line_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_data_input_stream_read_line_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_line_finish_utf8 : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_line_utf8 : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_until : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_until_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_data_input_stream_read_until_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_data_input_stream_read_upto : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_data_input_stream_read_upto_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_desktop_app_info_get_keywords : no return type

// Unsupported : g_desktop_app_info_launch_uris_as_manager : unsupported parameter user_setup : no type generator for GLib.SpawnChildSetupFunc, GSpawnChildSetupFunc

// Unsupported : g_desktop_app_info_list_actions : no return type

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_get_icon : no return generator

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_get_icon : no return generator

// Unsupported : g_file_enumerator_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_enumerator_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_enumerator_get_child : no return generator

// Unsupported : g_file_enumerator_get_container : no return generator

// Unsupported : g_file_enumerator_iterate : unsupported parameter out_info : record with indirection level of 2

// Unsupported : g_file_enumerator_next_files_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_enumerator_next_files_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_io_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_file_icon_get_file : no return generator

// Unsupported : g_file_info_get_attribute_data : unsupported parameter type : GFileAttributeType* with indirection level of 1

// Unsupported : g_file_info_get_attribute_stringv : no return type

// Unsupported : g_file_info_get_icon : no return generator

// Unsupported : g_file_info_get_symbolic_icon : no return generator

// Unsupported : g_file_info_list_attributes : no return type

// Unsupported : g_file_info_set_attribute_stringv : unsupported parameter attr_value : no param type

// Unsupported : g_file_info_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_info_set_symbolic_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_input_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_input_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_monitor_emit_event : unsupported parameter child : no type generator for File, GFile*

// Unsupported : g_file_output_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_output_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_filename_completer_get_completions : no return type

// Unsupported : g_io_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_io_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_io_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_inet_address_to_bytes : no return generator

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native *C.GInetAddressMask
	// parent_instance : record
	// Private : priv
}

func InetAddressMaskNewFromC(u unsafe.Pointer) *InetAddressMask {
	c := (*C.GInetAddressMask)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMask{native: c}

	return g
}

func (recv *InetAddressMask) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddressMaskNew is a wrapper around the C function g_inet_address_mask_new.
func InetAddressMaskNew(addr *InetAddress, length uint32) (*InetAddressMask, error) {
	c_addr := (*C.GInetAddress)(addr.ToC())

	c_length := (C.guint)(length)

	var cThrowableError *C.GError

	retC := C.g_inet_address_mask_new(c_addr, c_length, &cThrowableError)
	retGo := InetAddressMaskNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// InetAddressMaskNewFromString is a wrapper around the C function g_inet_address_mask_new_from_string.
func InetAddressMaskNewFromString(maskString string) (*InetAddressMask, error) {
	c_mask_string := C.CString(maskString)
	defer C.free(unsafe.Pointer(c_mask_string))

	var cThrowableError *C.GError

	retC := C.g_inet_address_mask_new_from_string(c_mask_string, &cThrowableError)
	retGo := InetAddressMaskNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Equal is a wrapper around the C function g_inet_address_mask_equal.
func (recv *InetAddressMask) Equal(mask2 *InetAddressMask) bool {
	c_mask2 := (*C.GInetAddressMask)(mask2.ToC())

	retC := C.g_inet_address_mask_equal((*C.GInetAddressMask)(recv.native), c_mask2)
	retGo := retC == C.TRUE

	return retGo
}

// GetAddress is a wrapper around the C function g_inet_address_mask_get_address.
func (recv *InetAddressMask) GetAddress() *InetAddress {
	retC := C.g_inet_address_mask_get_address((*C.GInetAddressMask)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function g_inet_address_mask_get_family.
func (recv *InetAddressMask) GetFamily() SocketFamily {
	retC := C.g_inet_address_mask_get_family((*C.GInetAddressMask)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLength is a wrapper around the C function g_inet_address_mask_get_length.
func (recv *InetAddressMask) GetLength() uint32 {
	retC := C.g_inet_address_mask_get_length((*C.GInetAddressMask)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Matches is a wrapper around the C function g_inet_address_mask_matches.
func (recv *InetAddressMask) Matches(address *InetAddress) bool {
	c_address := (*C.GInetAddress)(address.ToC())

	retC := C.g_inet_address_mask_matches((*C.GInetAddressMask)(recv.native), c_address)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function g_inet_address_mask_to_string.
func (recv *InetAddressMask) ToString() string {
	retC := C.g_inet_address_mask_to_string((*C.GInetAddressMask)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetFlowinfo is a wrapper around the C function g_inet_socket_address_get_flowinfo.
func (recv *InetSocketAddress) GetFlowinfo() uint32 {
	retC := C.g_inet_socket_address_get_flowinfo((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetScopeId is a wrapper around the C function g_inet_socket_address_get_scope_id.
func (recv *InetSocketAddress) GetScopeId() uint32 {
	retC := C.g_inet_socket_address_get_scope_id((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_input_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_read_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_skip_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_skip_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_splice : unsupported parameter additions : no param type

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_input_stream_add_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Menu is a wrapper around the C record GMenu.
type Menu struct {
	native *C.GMenu
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuNew is a wrapper around the C function g_menu_new.
func MenuNew() *Menu {
	retC := C.g_menu_new()
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function g_menu_append.
func (recv *Menu) Append(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_append((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// AppendItem is a wrapper around the C function g_menu_append_item.
func (recv *Menu) AppendItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(item.ToC())

	C.g_menu_append_item((*C.GMenu)(recv.native), c_item)

	return
}

// AppendSection is a wrapper around the C function g_menu_append_section.
func (recv *Menu) AppendSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(section.ToC())

	C.g_menu_append_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// AppendSubmenu is a wrapper around the C function g_menu_append_submenu.
func (recv *Menu) AppendSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(submenu.ToC())

	C.g_menu_append_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Freeze is a wrapper around the C function g_menu_freeze.
func (recv *Menu) Freeze() {
	C.g_menu_freeze((*C.GMenu)(recv.native))

	return
}

// Insert is a wrapper around the C function g_menu_insert.
func (recv *Menu) Insert(position int32, label string, detailedAction string) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_insert((*C.GMenu)(recv.native), c_position, c_label, c_detailed_action)

	return
}

// InsertItem is a wrapper around the C function g_menu_insert_item.
func (recv *Menu) InsertItem(position int32, item *MenuItem) {
	c_position := (C.gint)(position)

	c_item := (*C.GMenuItem)(item.ToC())

	C.g_menu_insert_item((*C.GMenu)(recv.native), c_position, c_item)

	return
}

// InsertSection is a wrapper around the C function g_menu_insert_section.
func (recv *Menu) InsertSection(position int32, label string, section *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(section.ToC())

	C.g_menu_insert_section((*C.GMenu)(recv.native), c_position, c_label, c_section)

	return
}

// InsertSubmenu is a wrapper around the C function g_menu_insert_submenu.
func (recv *Menu) InsertSubmenu(position int32, label string, submenu *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(submenu.ToC())

	C.g_menu_insert_submenu((*C.GMenu)(recv.native), c_position, c_label, c_submenu)

	return
}

// Prepend is a wrapper around the C function g_menu_prepend.
func (recv *Menu) Prepend(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_prepend((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// PrependItem is a wrapper around the C function g_menu_prepend_item.
func (recv *Menu) PrependItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(item.ToC())

	C.g_menu_prepend_item((*C.GMenu)(recv.native), c_item)

	return
}

// PrependSection is a wrapper around the C function g_menu_prepend_section.
func (recv *Menu) PrependSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(section.ToC())

	C.g_menu_prepend_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// PrependSubmenu is a wrapper around the C function g_menu_prepend_submenu.
func (recv *Menu) PrependSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(submenu.ToC())

	C.g_menu_prepend_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Remove is a wrapper around the C function g_menu_remove.
func (recv *Menu) Remove(position int32) {
	c_position := (C.gint)(position)

	C.g_menu_remove((*C.GMenu)(recv.native), c_position)

	return
}

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native *C.GMenuAttributeIter
	// parent_instance : record
	// priv : record
}

func MenuAttributeIterNewFromC(u unsafe.Pointer) *MenuAttributeIter {
	c := (*C.GMenuAttributeIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIter{native: c}

	return g
}

func (recv *MenuAttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetName is a wrapper around the C function g_menu_attribute_iter_get_name.
func (recv *MenuAttributeIter) GetName() string {
	retC := C.g_menu_attribute_iter_get_name((*C.GMenuAttributeIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_menu_attribute_iter_get_next : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_get_value : return type : Blacklisted record : GVariant

// Next is a wrapper around the C function g_menu_attribute_iter_next.
func (recv *MenuAttributeIter) Next() bool {
	retC := C.g_menu_attribute_iter_next((*C.GMenuAttributeIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// MenuItem is a wrapper around the C record GMenuItem.
type MenuItem struct {
	native *C.GMenuItem
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemNew is a wrapper around the C function g_menu_item_new.
func MenuItemNew(label string, detailedAction string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	retC := C.g_menu_item_new(c_label, c_detailed_action)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewSection is a wrapper around the C function g_menu_item_new_section.
func MenuItemNewSection(label string, section *MenuModel) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(section.ToC())

	retC := C.g_menu_item_new_section(c_label, c_section)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewSubmenu is a wrapper around the C function g_menu_item_new_submenu.
func MenuItemNewSubmenu(label string, submenu *MenuModel) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(submenu.ToC())

	retC := C.g_menu_item_new_submenu(c_label, c_submenu)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Unsupported : g_menu_item_set_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_action_and_target_value : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_attribute_value : unsupported parameter value : Blacklisted record : GVariant

// SetDetailedAction is a wrapper around the C function g_menu_item_set_detailed_action.
func (recv *MenuItem) SetDetailedAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_item_set_detailed_action((*C.GMenuItem)(recv.native), c_detailed_action)

	return
}

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// SetLabel is a wrapper around the C function g_menu_item_set_label.
func (recv *MenuItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.g_menu_item_set_label((*C.GMenuItem)(recv.native), c_label)

	return
}

// SetLink is a wrapper around the C function g_menu_item_set_link.
func (recv *MenuItem) SetLink(link string, model *MenuModel) {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	c_model := (*C.GMenuModel)(model.ToC())

	C.g_menu_item_set_link((*C.GMenuItem)(recv.native), c_link, c_model)

	return
}

// SetSection is a wrapper around the C function g_menu_item_set_section.
func (recv *MenuItem) SetSection(section *MenuModel) {
	c_section := (*C.GMenuModel)(section.ToC())

	C.g_menu_item_set_section((*C.GMenuItem)(recv.native), c_section)

	return
}

// SetSubmenu is a wrapper around the C function g_menu_item_set_submenu.
func (recv *MenuItem) SetSubmenu(submenu *MenuModel) {
	c_submenu := (*C.GMenuModel)(submenu.ToC())

	C.g_menu_item_set_submenu((*C.GMenuItem)(recv.native), c_submenu)

	return
}

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native *C.GMenuLinkIter
	// parent_instance : record
	// priv : record
}

func MenuLinkIterNewFromC(u unsafe.Pointer) *MenuLinkIter {
	c := (*C.GMenuLinkIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIter{native: c}

	return g
}

func (recv *MenuLinkIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetName is a wrapper around the C function g_menu_link_iter_get_name.
func (recv *MenuLinkIter) GetName() string {
	retC := C.g_menu_link_iter_get_name((*C.GMenuLinkIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_menu_link_iter_get_next : unsupported parameter value : record with indirection level of 2

// GetValue is a wrapper around the C function g_menu_link_iter_get_value.
func (recv *MenuLinkIter) GetValue() *MenuModel {
	retC := C.g_menu_link_iter_get_value((*C.GMenuLinkIter)(recv.native))
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Next is a wrapper around the C function g_menu_link_iter_next.
func (recv *MenuLinkIter) Next() bool {
	retC := C.g_menu_link_iter_next((*C.GMenuLinkIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native *C.GMenuModel
	// parent_instance : record
	// priv : record
}

func MenuModelNewFromC(u unsafe.Pointer) *MenuModel {
	c := (*C.GMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &MenuModel{native: c}

	return g
}

func (recv *MenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_menu_model_get_item_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_model_get_item_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// GetItemLink is a wrapper around the C function g_menu_model_get_item_link.
func (recv *MenuModel) GetItemLink(itemIndex int32, link string) *MenuModel {
	c_item_index := (C.gint)(itemIndex)

	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_model_get_item_link((*C.GMenuModel)(recv.native), c_item_index, c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNItems is a wrapper around the C function g_menu_model_get_n_items.
func (recv *MenuModel) GetNItems() int32 {
	retC := C.g_menu_model_get_n_items((*C.GMenuModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsMutable is a wrapper around the C function g_menu_model_is_mutable.
func (recv *MenuModel) IsMutable() bool {
	retC := C.g_menu_model_is_mutable((*C.GMenuModel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ItemsChanged is a wrapper around the C function g_menu_model_items_changed.
func (recv *MenuModel) ItemsChanged(position int32, removed int32, added int32) {
	c_position := (C.gint)(position)

	c_removed := (C.gint)(removed)

	c_added := (C.gint)(added)

	C.g_menu_model_items_changed((*C.GMenuModel)(recv.native), c_position, c_removed, c_added)

	return
}

// IterateItemAttributes is a wrapper around the C function g_menu_model_iterate_item_attributes.
func (recv *MenuModel) IterateItemAttributes(itemIndex int32) *MenuAttributeIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_attributes((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuAttributeIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterateItemLinks is a wrapper around the C function g_menu_model_iterate_item_links.
func (recv *MenuModel) IterateItemLinks(itemIndex int32) *MenuLinkIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_links((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuLinkIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_notification_add_button_with_target : unsupported parameter ... : varargs

// Unsupported : g_notification_add_button_with_target_value : unsupported parameter target : Blacklisted record : GVariant

// Unsupported : g_notification_set_default_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_notification_set_default_action_and_target_value : unsupported parameter target : Blacklisted record : GVariant

// Unsupported : g_notification_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_output_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_output_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_flush_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_output_stream_flush_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_printf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_output_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_output_stream_splice_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_vprintf : unsupported parameter bytes_written : no type generator for gsize, gsize*

// Unsupported : g_output_stream_write : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all_async : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_write_async : unsupported parameter buffer : no param type

// Unsupported : g_output_stream_write_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_output_stream_write_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_output_stream_write_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_permission_acquire_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_permission_acquire_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_permission_release_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_permission_release_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_address_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_name_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_records_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SettingsNewFull is a wrapper around the C function g_settings_new_full.
func SettingsNewFull(schema *SettingsSchema, backend *SettingsBackend, path string) *Settings {
	c_schema := (*C.GSettingsSchema)(schema.ToC())

	c_backend := (*C.GSettingsBackend)(backend.ToC())

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_full(c_schema, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_settings_bind_with_mapping : unsupported parameter get_mapping : no type generator for SettingsBindGetMapping, GSettingsBindGetMapping

// Unsupported : g_settings_create_action : no return generator

// Unsupported : g_settings_get : unsupported parameter ... : varargs

// Unsupported : g_settings_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_mapped : unsupported parameter mapping : no type generator for SettingsGetMapping, GSettingsGetMapping

// Unsupported : g_settings_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_strv : no return type

// Unsupported : g_settings_get_user_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_list_children : no return type

// Unsupported : g_settings_list_keys : no return type

// Unsupported : g_settings_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_set : unsupported parameter ... : varargs

// Unsupported : g_settings_set_strv : unsupported parameter value : no param type

// Unsupported : g_settings_set_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_backend_keys_changed : unsupported parameter items : no param type

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_set_state : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_simple_action_set_state_hint : unsupported parameter state_hint : Blacklisted record : GVariant

// Unsupported : g_simple_action_group_add_entries : unsupported parameter entries : no param type

// Unsupported : g_simple_action_group_insert : unsupported parameter action : no type generator for Action, GAction*

// Unsupported : g_simple_action_group_lookup : no return generator

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_run_in_thread : unsupported parameter func : no type generator for SimpleAsyncThreadFunc, GSimpleAsyncThreadFunc

// SetCheckCancellable is a wrapper around the C function g_simple_async_result_set_check_cancellable.
func (recv *SimpleAsyncResult) SetCheckCancellable(checkCancellable *Cancellable) {
	c_check_cancellable := (*C.GCancellable)(checkCancellable.ToC())

	C.g_simple_async_result_set_check_cancellable((*C.GSimpleAsyncResult)(recv.native), c_check_cancellable)

	return
}

// Unsupported : g_simple_async_result_set_error : unsupported parameter ... : varargs

// Unsupported : g_simple_async_result_set_error_va : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter destroy_op_res : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// ConditionTimedWait is a wrapper around the C function g_socket_condition_timed_wait.
func (recv *Socket) ConditionTimedWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_condition_timed_wait((*C.GSocket)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetAvailableBytes is a wrapper around the C function g_socket_get_available_bytes.
func (recv *Socket) GetAvailableBytes() int64 {
	retC := C.g_socket_get_available_bytes((*C.GSocket)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetBroadcast is a wrapper around the C function g_socket_get_broadcast.
func (recv *Socket) GetBroadcast() bool {
	retC := C.g_socket_get_broadcast((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMulticastLoopback is a wrapper around the C function g_socket_get_multicast_loopback.
func (recv *Socket) GetMulticastLoopback() bool {
	retC := C.g_socket_get_multicast_loopback((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMulticastTtl is a wrapper around the C function g_socket_get_multicast_ttl.
func (recv *Socket) GetMulticastTtl() uint32 {
	retC := C.g_socket_get_multicast_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_socket_get_option : unsupported parameter value : no type generator for gint, gint*

// GetTtl is a wrapper around the C function g_socket_get_ttl.
func (recv *Socket) GetTtl() uint32 {
	retC := C.g_socket_get_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// JoinMulticastGroup is a wrapper around the C function g_socket_join_multicast_group.
func (recv *Socket) JoinMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LeaveMulticastGroup is a wrapper around the C function g_socket_leave_multicast_group.
func (recv *Socket) LeaveMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_receive : unsupported parameter buffer : no param type

// Unsupported : g_socket_receive_from : unsupported parameter address : record with indirection level of 2

// Unsupported : g_socket_receive_message : unsupported parameter address : record with indirection level of 2

// Unsupported : g_socket_receive_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_receive_with_blocking : unsupported parameter buffer : no param type

// Unsupported : g_socket_send : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_message : unsupported parameter vectors : no param type

// Unsupported : g_socket_send_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_send_to : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_with_blocking : unsupported parameter buffer : no param type

// SetBroadcast is a wrapper around the C function g_socket_set_broadcast.
func (recv *Socket) SetBroadcast(broadcast bool) {
	c_broadcast :=
		boolToGboolean(broadcast)

	C.g_socket_set_broadcast((*C.GSocket)(recv.native), c_broadcast)

	return
}

// SetMulticastLoopback is a wrapper around the C function g_socket_set_multicast_loopback.
func (recv *Socket) SetMulticastLoopback(loopback bool) {
	c_loopback :=
		boolToGboolean(loopback)

	C.g_socket_set_multicast_loopback((*C.GSocket)(recv.native), c_loopback)

	return
}

// SetMulticastTtl is a wrapper around the C function g_socket_set_multicast_ttl.
func (recv *Socket) SetMulticastTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_multicast_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

// SetTtl is a wrapper around the C function g_socket_set_ttl.
func (recv *Socket) SetTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

// Unsupported : g_socket_address_enumerator_next_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_address_enumerator_next_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_async : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_host_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_uri_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_get_proxy_resolver : no return generator

// Unsupported : g_socket_client_set_proxy_resolver : unsupported parameter proxy_resolver : no type generator for ProxyResolver, GProxyResolver*

// Connect is a wrapper around the C function g_socket_connection_connect.
func (recv *SocketConnection) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_connection_connect((*C.GSocketConnection)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_connection_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// IsConnected is a wrapper around the C function g_socket_connection_is_connected.
func (recv *SocketConnection) IsConnected() bool {
	retC := C.g_socket_connection_is_connected((*C.GSocketConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_socket_listener_accept : unsupported parameter source_object : record with indirection level of 2

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_accept_socket : unsupported parameter source_object : record with indirection level of 2

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_socket_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_add_address : unsupported parameter effective_address : record with indirection level of 2

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_subprocess_communicate : unsupported parameter stdout_buf : record with indirection level of 2

// Unsupported : g_subprocess_communicate_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_communicate_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_communicate_utf8_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_communicate_utf8_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_wait_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_wait_check_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_subprocess_wait_check_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_wait_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_subprocess_launcher_set_child_setup : unsupported parameter child_setup : no type generator for GLib.SpawnChildSetupFunc, GSpawnChildSetupFunc

// Unsupported : g_subprocess_launcher_set_environ : unsupported parameter env : no param type

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_launcher_spawnv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_task_return_new_error : unsupported parameter ... : varargs

// Unsupported : g_task_return_pointer : unsupported parameter result_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc, GTaskThreadFunc

// Unsupported : g_task_set_task_data : unsupported parameter task_data_destroy : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_themed_icon_get_names : no return type

// Unsupported : g_tls_certificate_verify : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_tls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_connection_handshake_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_lookup_certificate_for_handle_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_database_lookup_certificate_for_handle_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_lookup_certificate_issuer_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_database_lookup_certificate_issuer_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_lookup_certificates_issued_by : unsupported parameter issuer_raw_dn : no param type

// Unsupported : g_tls_database_lookup_certificates_issued_by_async : unsupported parameter issuer_raw_dn : no param type

// Unsupported : g_tls_database_lookup_certificates_issued_by_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_verify_chain : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_tls_database_verify_chain_async : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_tls_database_verify_chain_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_interaction_ask_password_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_interaction_ask_password_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_interaction_request_certificate_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_password_get_value : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_tls_password_set_value : unsupported parameter value : no param type

// Unsupported : g_tls_password_set_value_full : unsupported parameter value : no param type

// Unsupported : g_unix_connection_receive_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_unix_connection_receive_credentials_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_unix_connection_send_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_unix_connection_send_credentials_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_fd_list_peek_fds : unsupported parameter length : no type generator for gint, gint*

// Unsupported : g_unix_fd_list_steal_fds : unsupported parameter length : no type generator for gint, gint*

// Unsupported : g_unix_fd_message_steal_fds : unsupported parameter length : no type generator for gint, gint*

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported : g_vfs_get_file_for_path : no return generator

// Unsupported : g_vfs_get_file_for_uri : no return generator

// Unsupported : g_vfs_get_supported_uri_schemes : no return type

// Unsupported : g_vfs_parse_name : no return generator

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc, GVfsFileLookupFunc

// Unsupported : g_volume_monitor_get_mount_for_uuid : no return generator

// Unsupported : g_volume_monitor_get_volume_for_uuid : no return generator
