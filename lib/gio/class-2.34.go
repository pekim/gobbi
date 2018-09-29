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

// Unsupported : g_app_launch_context_get_display : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Unsupported : g_app_launch_context_get_environment : no return type

// Unsupported : g_app_launch_context_get_startup_notify_id : unsupported parameter info : no type generator for AppInfo, GAppInfo*

// Unsupported : g_application_add_main_option : unsupported parameter short_name : no type generator for gchar, char

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries : no param type

// GetDbusConnection is a wrapper around the C function g_application_get_dbus_connection.
func (recv *Application) GetDbusConnection() *DBusConnection {
	retC := C.g_application_get_dbus_connection((*C.GApplication)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDbusObjectPath is a wrapper around the C function g_application_get_dbus_object_path.
func (recv *Application) GetDbusObjectPath() string {
	retC := C.g_application_get_dbus_object_path((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_open : unsupported parameter files : no param type

// Unsupported : g_application_run : unsupported parameter argv : no param type

// Unsupported : g_application_set_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// Unsupported : g_application_command_line_create_file_for_arg : no return generator

// Unsupported : g_application_command_line_get_arguments : unsupported parameter argc : no type generator for gint, int*

// Unsupported : g_application_command_line_get_environ : no return type

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// GetStdin is a wrapper around the C function g_application_command_line_get_stdin.
func (recv *ApplicationCommandLine) GetStdin() *InputStream {
	retC := C.g_application_command_line_get_stdin((*C.GApplicationCommandLine)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// AllowMechanism is a wrapper around the C function g_dbus_auth_observer_allow_mechanism.
func (recv *DBusAuthObserver) AllowMechanism(mechanism string) bool {
	c_mechanism := C.CString(mechanism)
	defer C.free(unsafe.Pointer(c_mechanism))

	retC := C.g_dbus_auth_observer_allow_mechanism((*C.GDBusAuthObserver)(recv.native), c_mechanism)
	retGo := retC == C.TRUE

	return retGo
}

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

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_flush_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// GetLastSerial is a wrapper around the C function g_dbus_connection_get_last_serial.
func (recv *DBusConnection) GetLastSerial() uint32 {
	retC := C.g_dbus_connection_get_last_serial((*C.GDBusConnection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_send_message : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_send_message_with_reply_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_send_message_with_reply_sync : unsupported parameter out_serial : no type generator for guint32, volatile guint32*

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback, GDBusSignalCallback

// Unsupported : g_dbus_interface_skeleton_get_properties : return type : Blacklisted record : GVariant

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

// IsExported is a wrapper around the C function g_dbus_object_manager_server_is_exported.
func (recv *DBusObjectManagerServer) IsExported(object *DBusObjectSkeleton) bool {
	c_object := (*C.GDBusObjectSkeleton)(object.ToC())

	retC := C.g_dbus_object_manager_server_is_exported((*C.GDBusObjectManagerServer)(recv.native), c_object)
	retGo := retC == C.TRUE

	return retGo
}

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

// GetStartupWmClass is a wrapper around the C function g_desktop_app_info_get_startup_wm_class.
func (recv *DesktopAppInfo) GetStartupWmClass() string {
	retC := C.g_desktop_app_info_get_startup_wm_class((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

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

// Unsupported : g_input_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_async : unsupported parameter buffer : no param type

// Unsupported : g_input_stream_read_all_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_async : unsupported parameter buffer : no param type

// ReadBytes is a wrapper around the C function g_input_stream_read_bytes.
func (recv *InputStream) ReadBytes(count uint64, cancellable *Cancellable) (*glib.Bytes, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_bytes((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_read_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_read_bytes_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_read_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_input_stream_skip_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_input_stream_skip_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_list_store_insert_sorted : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_sort : unsupported parameter compare_func : no type generator for GLib.CompareDataFunc, GCompareDataFunc

// Unsupported : g_list_store_splice : unsupported parameter additions : no param type

// MemoryInputStreamNewFromBytes is a wrapper around the C function g_memory_input_stream_new_from_bytes.
func MemoryInputStreamNewFromBytes(bytes *glib.Bytes) *InputStream {
	c_bytes := (*C.GBytes)(bytes.ToC())

	retC := C.g_memory_input_stream_new_from_bytes(c_bytes)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// AddBytes is a wrapper around the C function g_memory_input_stream_add_bytes.
func (recv *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	c_bytes := (*C.GBytes)(bytes.ToC())

	C.g_memory_input_stream_add_bytes((*C.GMemoryInputStream)(recv.native), c_bytes)

	return
}

// Unsupported : g_memory_input_stream_add_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// StealAsBytes is a wrapper around the C function g_memory_output_stream_steal_as_bytes.
func (recv *MemoryOutputStream) StealAsBytes() *glib.Bytes {
	retC := C.g_memory_output_stream_steal_as_bytes((*C.GMemoryOutputStream)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_attribute_iter_get_next : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_get_value : return type : Blacklisted record : GVariant

// MenuItemNewFromModel is a wrapper around the C function g_menu_item_new_from_model.
func MenuItemNewFromModel(model *MenuModel, itemIndex int32) *MenuItem {
	c_model := (*C.GMenuModel)(model.ToC())

	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_item_new_from_model(c_model, c_item_index)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// GetLink is a wrapper around the C function g_menu_item_get_link.
func (recv *MenuItem) GetLink(link string) *MenuModel {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_item_get_link((*C.GMenuItem)(recv.native), c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_set_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_action_and_target_value : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_attribute_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_menu_link_iter_get_next : unsupported parameter value : record with indirection level of 2

// Unsupported : g_menu_model_get_item_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_model_get_item_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

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

// GetDestinationProtocol is a wrapper around the C function g_proxy_address_get_destination_protocol.
func (recv *ProxyAddress) GetDestinationProtocol() string {
	retC := C.g_proxy_address_get_destination_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function g_proxy_address_get_uri.
func (recv *ProxyAddress) GetUri() string {
	retC := C.g_proxy_address_get_uri((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_address_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_name_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// LookupRecords is a wrapper around the C function g_resolver_lookup_records.
func (recv *Resolver) LookupRecords(rrname string, recordType ResolverRecordType, cancellable *Cancellable) (*glib.List, error) {
	c_rrname := C.CString(rrname)
	defer C.free(unsafe.Pointer(c_rrname))

	c_record_type := (C.GResolverRecordType)(recordType)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_records((*C.GResolver)(recv.native), c_rrname, c_record_type, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_records_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_records_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

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

// Unsupported : g_simple_async_result_set_error : unsupported parameter ... : varargs

// Unsupported : g_simple_async_result_set_error_va : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter destroy_op_res : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// Unsupported : g_socket_get_option : unsupported parameter value : no type generator for gint, gint*

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

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_connection_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

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

// TestDBus is a wrapper around the C record GTestDBus.
type TestDBus struct {
	native *C.GTestDBus
}

func TestDBusNewFromC(u unsafe.Pointer) *TestDBus {
	c := (*C.GTestDBus)(u)
	if c == nil {
		return nil
	}

	g := &TestDBus{native: c}

	return g
}

func (recv *TestDBus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TestDBusNew is a wrapper around the C function g_test_dbus_new.
func TestDBusNew(flags TestDBusFlags) *TestDBus {
	c_flags := (C.GTestDBusFlags)(flags)

	retC := C.g_test_dbus_new(c_flags)
	retGo := TestDBusNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddServiceDir is a wrapper around the C function g_test_dbus_add_service_dir.
func (recv *TestDBus) AddServiceDir(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_test_dbus_add_service_dir((*C.GTestDBus)(recv.native), c_path)

	return
}

// Down is a wrapper around the C function g_test_dbus_down.
func (recv *TestDBus) Down() {
	C.g_test_dbus_down((*C.GTestDBus)(recv.native))

	return
}

// GetBusAddress is a wrapper around the C function g_test_dbus_get_bus_address.
func (recv *TestDBus) GetBusAddress() string {
	retC := C.g_test_dbus_get_bus_address((*C.GTestDBus)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_test_dbus_get_flags.
func (recv *TestDBus) GetFlags() TestDBusFlags {
	retC := C.g_test_dbus_get_flags((*C.GTestDBus)(recv.native))
	retGo := (TestDBusFlags)(retC)

	return retGo
}

// Stop is a wrapper around the C function g_test_dbus_stop.
func (recv *TestDBus) Stop() {
	C.g_test_dbus_stop((*C.GTestDBus)(recv.native))

	return
}

// Up is a wrapper around the C function g_test_dbus_up.
func (recv *TestDBus) Up() {
	C.g_test_dbus_up((*C.GTestDBus)(recv.native))

	return
}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_themed_icon_get_names : no return type

// IsSame is a wrapper around the C function g_tls_certificate_is_same.
func (recv *TlsCertificate) IsSame(certTwo *TlsCertificate) bool {
	c_cert_two := (*C.GTlsCertificate)(certTwo.ToC())

	retC := C.g_tls_certificate_is_same((*C.GTlsCertificate)(recv.native), c_cert_two)
	retGo := retC == C.TRUE

	return retGo
}

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
