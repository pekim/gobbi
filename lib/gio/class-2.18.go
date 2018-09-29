// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_application_open : unsupported parameter files : no param type

// Unsupported : g_application_run : unsupported parameter argv : no param type

// Unsupported : g_application_set_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

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

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_flush_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

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

// DesktopAppInfoNewFromKeyfile is a wrapper around the C function g_desktop_app_info_new_from_keyfile.
func DesktopAppInfoNewFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	c_key_file := (*C.GKeyFile)(keyFile.ToC())

	retC := C.g_desktop_app_info_new_from_keyfile(c_key_file)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_desktop_app_info_get_keywords : no return type

// Unsupported : g_desktop_app_info_launch_uris_as_manager : unsupported parameter user_setup : no type generator for GLib.SpawnChildSetupFunc, GSpawnChildSetupFunc

// Unsupported : g_desktop_app_info_list_actions : no return type

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_get_icon : no return generator

// GetOrigin is a wrapper around the C function g_emblem_get_origin.
func (recv *Emblem) GetOrigin() EmblemOrigin {
	retC := C.g_emblem_get_origin((*C.GEmblem)(recv.native))
	retGo := (EmblemOrigin)(retC)

	return retGo
}

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// AddEmblem is a wrapper around the C function g_emblemed_icon_add_emblem.
func (recv *EmblemedIcon) AddEmblem(emblem *Emblem) {
	c_emblem := (*C.GEmblem)(emblem.ToC())

	C.g_emblemed_icon_add_emblem((*C.GEmblemedIcon)(recv.native), c_emblem)

	return
}

// GetEmblems is a wrapper around the C function g_emblemed_icon_get_emblems.
func (recv *EmblemedIcon) GetEmblems() *glib.List {
	retC := C.g_emblemed_icon_get_emblems((*C.GEmblemedIcon)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetDataSize is a wrapper around the C function g_memory_output_stream_get_data_size.
func (recv *MemoryOutputStream) GetDataSize() uint64 {
	retC := C.g_memory_output_stream_get_data_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_menu_attribute_iter_get_next : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_get_value : return type : Blacklisted record : GVariant

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

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

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_address_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_name_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

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

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// Unsupported : g_themed_icon_get_names : no return type

// PrependName is a wrapper around the C function g_themed_icon_prepend_name.
func (recv *ThemedIcon) PrependName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_prepend_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
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

// SetRateLimit is a wrapper around the C function g_unix_mount_monitor_set_rate_limit.
func (recv *UnixMountMonitor) SetRateLimit(limitMsec int32) {
	c_limit_msec := (C.int)(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit((*C.GUnixMountMonitor)(recv.native), c_limit_msec)

	return
}

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// Unsupported : g_vfs_get_file_for_path : no return generator

// Unsupported : g_vfs_get_file_for_uri : no return generator

// Unsupported : g_vfs_get_supported_uri_schemes : no return type

// Unsupported : g_vfs_parse_name : no return generator

// Unsupported : g_vfs_register_uri_scheme : unsupported parameter uri_func : no type generator for VfsFileLookupFunc, GVfsFileLookupFunc

// Unsupported : g_volume_monitor_get_mount_for_uuid : no return generator

// Unsupported : g_volume_monitor_get_volume_for_uuid : no return generator
