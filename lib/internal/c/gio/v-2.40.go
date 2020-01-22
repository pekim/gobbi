// Code generated - DO NOT EDIT.
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

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
/*

static gboolean c_g_output_stream_printf(GOutputStream* stream, gsize* bytes_written, GCancellable* cancellable, GError** error, const gchar* format) {
    return g_output_stream_printf(stream, bytes_written, cancellable, error, format, NULL);
}
*/
/*

static gboolean c_g_output_stream_vprintf(GOutputStream* stream, gsize* bytes_written, GCancellable* cancellable, GError** error, const gchar* format) {
    return g_output_stream_vprintf(stream, bytes_written, cancellable, error, format, NULL);
}
*/
/*

static GSubprocess* c_g_subprocess_launcher_spawn(GSubprocessLauncher* self, GError** error, const gchar* argv0) {
    return g_subprocess_launcher_spawn(self, error, argv0, NULL);
}
*/
import "C"

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

func Fn_g_settings_schema_get_key(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_schema_get_key(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_has_key(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_schema_has_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

func Fn_g_settings_schema_key_get_default_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_default_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_get_range(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_range(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_get_value_type(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_value_type(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_range_check(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.g_settings_schema_key_range_check(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_settings_schema_key_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_key_unref(cValueInstance)
}

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

func Fn_g_app_info_monitor_get() unsafe.Pointer {
	ret := C.g_app_info_monitor_get()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

// UNSUPPORTED : g_application_add_main_option_entries : parameter 'entries' is array parameter without length parameter

func Fn_g_application_add_option_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GOptionGroup)(unsafe.Pointer(param0))

	C.g_application_add_option_group(cValueInstance, cValue0)
}

func Fn_g_application_send_notification(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GNotification)(unsafe.Pointer(param1))

	C.g_application_send_notification(cValueInstance, cValue0, cValue1)
}

func Fn_g_application_withdraw_notification(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_withdraw_notification(cValueInstance, cValue0)
}

// UNSUPPORTED : g_application_command_line_get_environ : no array length

func Fn_g_application_command_line_get_options_dict(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_options_dict(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_message_get_header_fields : no array length

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_proxy_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : no array length

// UNSUPPORTED : g_dbus_proxy_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line : no array length

// UNSUPPORTED : g_data_input_stream_read_line_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : no array length

// UNSUPPORTED : g_data_input_stream_read_until_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_upto_async : parameter 'callback' is callback

// UNSUPPORTED : g_desktop_app_info_get_keywords : no array length

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_list_actions : no array length

// UNSUPPORTED : g_desktop_app_info_search : array has no type

// UNSUPPORTED : g_file_enumerator_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_enumerator_next_files_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_io_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_info_get_attribute_stringv : no array length

// UNSUPPORTED : g_file_info_list_attributes : no array length

// UNSUPPORTED : g_file_info_set_attribute_stringv : parameter 'attr_value' is array parameter without length parameter

// UNSUPPORTED : g_file_input_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_output_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_filename_completer_get_completions : no array length

// UNSUPPORTED : g_io_module_new : blacklisted

// UNSUPPORTED : g_io_module_load : blacklisted

// UNSUPPORTED : g_io_module_unload : blacklisted

// UNSUPPORTED : g_io_module_query : blacklisted

// UNSUPPORTED : g_io_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_io_stream_splice_async : parameter 'callback' is callback

// UNSUPPORTED : g_inet_address_new_from_bytes : parameter 'bytes' is array parameter without length parameter

func Fn_g_inet_socket_address_new_from_string(param0 string, param1 uint) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	ret := C.g_inet_socket_address_new_from_string(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_input_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read : blacklisted

// UNSUPPORTED : g_input_stream_read_all : blacklisted

// UNSUPPORTED : g_input_stream_read_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_skip_async : parameter 'callback' is callback

// UNSUPPORTED : g_list_store_insert_sorted : parameter 'compare_func' is callback

// UNSUPPORTED : g_list_store_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_memory_input_stream_new_from_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_input_stream_add_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_output_stream_new : parameter 'realloc_function' is callback

func Fn_g_notification_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_notification_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_notification_add_button(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_notification_add_button(cValueInstance, cValue0, cValue1)
}

func Fn_g_notification_add_button_with_target_value(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_notification_add_button_with_target_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_notification_set_body(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_body(cValueInstance, cValue0)
}

func Fn_g_notification_set_default_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_default_action(cValueInstance, cValue0)
}

func Fn_g_notification_set_default_action_and_target_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_notification_set_default_action_and_target_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_notification_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_notification_set_icon(cValueInstance, cValue0)
}

func Fn_g_notification_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_title(cValueInstance, cValue0)
}

func Fn_g_notification_set_urgent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_notification_set_urgent(cValueInstance, cValue0)
}

// UNSUPPORTED : g_output_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_flush_async : parameter 'callback' is callback

func Fn_g_output_stream_printf(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 string) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cValue2 := (**C.GError)(unsafe.Pointer(param2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.c_g_output_stream_printf(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_splice_async : parameter 'callback' is callback

func Fn_g_output_stream_vprintf(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 string) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cValue2 := (**C.GError)(unsafe.Pointer(param2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.c_g_output_stream_vprintf(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_write_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_async : parameter 'callback' is callback

// UNSUPPORTED : g_permission_acquire_async : parameter 'callback' is callback

// UNSUPPORTED : g_permission_release_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_address_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_records_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_service_async : parameter 'callback' is callback

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

func Fn_g_settings_get_default_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_default_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

// UNSUPPORTED : g_settings_get_strv : no array length

func Fn_g_settings_get_user_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_user_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_list_children : no array length

// UNSUPPORTED : g_settings_list_keys : no array length

// UNSUPPORTED : g_settings_set_strv : parameter 'value' is array parameter without length parameter

// UNSUPPORTED : g_settings_list_relocatable_schemas : no array length

// UNSUPPORTED : g_settings_list_schemas : no array length

// UNSUPPORTED : g_settings_backend_changed : blacklisted

// UNSUPPORTED : g_settings_backend_changed_tree : blacklisted

// UNSUPPORTED : g_settings_backend_keys_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_flatten_tree : blacklisted

// UNSUPPORTED : g_settings_backend_get_default : blacklisted

// UNSUPPORTED : g_simple_async_result_new : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_from_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_take_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_run_in_thread : parameter 'func' is callback

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_newv : parameter 'argv' is array parameter without length parameter

func Fn_g_subprocess_communicate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cValue2 := (**C.GBytes)(unsafe.Pointer(param2))

	cValue3 := (**C.GBytes)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_communicate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_communicate_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_communicate_utf8_async : parameter 'callback' is callback

func Fn_g_subprocess_force_exit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	C.g_subprocess_force_exit(cValueInstance)
}

func Fn_g_subprocess_get_exit_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_exit_status(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_get_identifier(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_identifier(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_subprocess_get_if_exited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_if_exited(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_if_signaled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_if_signaled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_status(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_get_stderr_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stderr_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_stdin_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stdin_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_stdout_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stdout_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_successful(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_successful(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_term_sig(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_term_sig(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_send_signal(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_send_signal(cValueInstance, cValue0)
}

func Fn_g_subprocess_wait(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_wait_async : parameter 'callback' is callback

func Fn_g_subprocess_wait_check(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_check(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_wait_check_async : parameter 'callback' is callback

func Fn_g_subprocess_wait_check_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_check_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_subprocess_wait_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_subprocess_launcher_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GSubprocessFlags)(param0)

	ret := C.g_subprocess_launcher_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_launcher_getenv(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_subprocess_launcher_getenv(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : parameter 'child_setup' is callback

func Fn_g_subprocess_launcher_set_cwd(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_cwd(cValueInstance, cValue0)
}

// UNSUPPORTED : g_subprocess_launcher_set_environ : parameter 'env' is array parameter without length parameter

func Fn_g_subprocess_launcher_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSubprocessFlags)(param0)

	C.g_subprocess_launcher_set_flags(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stderr_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stderr_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stdin_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stdin_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stdout_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stdout_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_setenv(paramInstance unsafe.Pointer, param0 string, param1 string, param2 bool) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := toCBool(param2)

	C.g_subprocess_launcher_setenv(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_subprocess_launcher_spawn(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GError)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_g_subprocess_launcher_spawn(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_subprocess_launcher_spawnv : parameter 'argv' is array parameter without length parameter

func Fn_g_subprocess_launcher_take_stderr_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stderr_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_take_stdin_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stdin_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_take_stdout_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stdout_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_unsetenv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_unsetenv(cValueInstance, cValue0)
}

// UNSUPPORTED : g_task_new : parameter 'callback' is callback

// UNSUPPORTED : g_task_attach_source : parameter 'callback' is callback

// UNSUPPORTED : g_task_return_pointer : parameter 'result_destroy' is callback

// UNSUPPORTED : g_task_run_in_thread : parameter 'task_func' is callback

// UNSUPPORTED : g_task_run_in_thread_sync : parameter 'task_func' is callback

// UNSUPPORTED : g_task_set_task_data : parameter 'task_data_destroy' is callback

// UNSUPPORTED : g_task_report_error : parameter 'callback' is callback

// UNSUPPORTED : g_task_report_new_error : parameter 'callback' is callback

// UNSUPPORTED : g_themed_icon_get_names : no array length

// UNSUPPORTED : g_tls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_verify_chain_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_ask_password_async : parameter 'callback' is callback

func Fn_g_tls_interaction_invoke_request_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsConnection)(unsafe.Pointer(param0))

	cValue1 := (C.GTlsCertificateRequestFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_invoke_request_certificate(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_request_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsConnection)(unsafe.Pointer(param0))

	cValue1 := (C.GTlsCertificateRequestFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_request_certificate(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

func Fn_g_tls_interaction_request_certificate_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_request_certificate_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

// UNSUPPORTED : g_action_group_list_actions : no array length

// UNSUPPORTED : g_app_info_get_supported_types : no array length

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

// UNSUPPORTED : g_drive_poll_for_media : parameter 'callback' is callback

// UNSUPPORTED : g_drive_start : parameter 'callback' is callback

// UNSUPPORTED : g_drive_stop : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_dtls_connection_shutdown_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_append_to_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_copy : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_copy_async : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_create_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_create_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_delete_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_file_enumerate_children_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_find_enclosing_mount_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents : blacklisted

// UNSUPPORTED : g_file_load_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents_finish : blacklisted

// UNSUPPORTED : g_file_load_partial_contents_async : parameter 'read_more_callback' is callback

// UNSUPPORTED : g_file_load_partial_contents_finish : blacklisted

// UNSUPPORTED : g_file_make_directory_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage_async : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_mount_enclosing_volume : parameter 'callback' is callback

// UNSUPPORTED : g_file_mount_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_move : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_open_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_poll_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_default_handler_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_filesystem_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_attributes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_display_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_start_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_stop_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_trash_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_loadable_icon_load_async : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback
