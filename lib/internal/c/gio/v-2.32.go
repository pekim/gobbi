// Code generated - DO NOT EDIT.
// +build gio_2.32

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

static void c_g_menu_item_set_action_and_target(GMenuItem* menu_item, const gchar* action, const gchar* format_string) {
    return g_menu_item_set_action_and_target(menu_item, action, format_string, NULL);
}
*/
/*

static void c_g_menu_item_set_attribute(GMenuItem* menu_item, const gchar* attribute, const gchar* format_string) {
    return g_menu_item_set_attribute(menu_item, attribute, format_string, NULL);
}
*/
/*

static gboolean c_g_menu_model_get_item_attribute(GMenuModel* model, gint item_index, const gchar* attribute, const gchar* format_string) {
    return g_menu_model_get_item_attribute(model, item_index, attribute, format_string, NULL);
}
*/
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type ActionMapInterface C.GActionMapInterface

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
type NetworkMonitorInterface C.GNetworkMonitorInterface
type RemoteActionGroupInterface C.GRemoteActionGroupInterface
type Resource C.GResource

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type SettingsSchema C.GSettingsSchema
type SettingsSchemaSource C.GSettingsSchemaSource

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

func Fn_g_file_attribute_matcher_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_matcher_to_string(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

func Fn_g_resource_new_from_data(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resource_new_from_data(cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resource_enumerate_children : no array length

func Fn_g_resource_get_info(paramInstance unsafe.Pointer, param0 string, param1 int, param2 *uint64, param3 *uint32, error unsafe.Pointer) bool {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resource_get_info(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_resource_lookup_data(paramInstance unsafe.Pointer, param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resource_lookup_data(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_open_stream(paramInstance unsafe.Pointer, param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resource_open_stream(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	ret := C.g_resource_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	C.g_resource_unref(cValueInstance)
}

func Fn_g_resource_load(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_resource_load(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_get_path(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

func Fn_g_settings_schema_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_unref(cValueInstance)
}

func Fn_g_settings_schema_source_new_from_directory(param0 string, param1 unsafe.Pointer, param2 bool, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsSchemaSource)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cError := (**C.GError)(error)

	ret := C.g_settings_schema_source_new_from_directory(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

func Fn_g_settings_schema_source_lookup(paramInstance unsafe.Pointer, param0 string, param1 bool) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.g_settings_schema_source_lookup(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_source_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_source_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_source_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_source_unref(cValueInstance)
}

func Fn_g_settings_schema_source_get_default() unsafe.Pointer {
	ret := C.g_settings_schema_source_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_static_resource_fini(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	C.g_static_resource_fini(cValueInstance)
}

func Fn_g_static_resource_get_resource(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	ret := C.g_static_resource_get_resource(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_static_resource_init(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	C.g_static_resource_init(cValueInstance)
}

func Fn_g_unix_mount_point_get_options(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_get_options(cValueInstance)

	return C.GoString(ret)
}

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

func Fn_g_resources_get_info(param0 string, param1 int, param2 *uint64, param3 *uint32, error unsafe.Pointer) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resources_get_info(cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_resources_lookup_data(param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resources_lookup_data(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resources_open_stream(param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resources_open_stream(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resources_register(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_register(cValue0)
}

func Fn_g_resources_unregister(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_unregister(cValue0)
}

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

func Fn_g_app_launch_context_setenv(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_app_launch_context_setenv(cValueInstance, cValue0, cValue1)
}

func Fn_g_app_launch_context_unsetenv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_launch_context_unsetenv(cValueInstance, cValue0)
}

// UNSUPPORTED : g_application_add_main_option_entries : parameter 'entries' is array parameter without length parameter

func Fn_g_application_quit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_quit(cValueInstance)
}

func Fn_g_application_set_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_set_default(cValueInstance)
}

func Fn_g_application_get_default() unsafe.Pointer {
	ret := C.g_application_get_default()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_application_command_line_get_environ : no array length

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

func Fn_g_dbus_action_group_get(param0 unsafe.Pointer, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_action_group_get(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

func Fn_g_dbus_connection_export_action_group(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GActionGroup)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_export_action_group(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_g_dbus_connection_export_menu_model(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_export_menu_model(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

func Fn_g_dbus_connection_unexport_action_group(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unexport_action_group(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_unexport_menu_model(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unexport_menu_model(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

func Fn_g_dbus_interface_skeleton_get_connections(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_connections(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_has_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	ret := C.g_dbus_interface_skeleton_has_connection(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_interface_skeleton_unexport_from_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_interface_skeleton_unexport_from_connection(cValueInstance, cValue0)
}

func Fn_g_dbus_menu_model_get(param0 unsafe.Pointer, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_menu_model_get(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

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

func Fn_g_inet_address_mask_new(param0 unsafe.Pointer, param1 uint, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cError := (**C.GError)(error)

	ret := C.g_inet_address_mask_new(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_new_from_string(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_inet_address_mask_new_from_string(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddressMask)(unsafe.Pointer(param0))

	ret := C.g_inet_address_mask_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_inet_address_mask_get_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_inet_address_mask_get_length(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_length(cValueInstance)

	return (uint)(ret)
}

func Fn_g_inet_address_mask_matches(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	ret := C.g_inet_address_mask_matches(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_inet_address_mask_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_inet_socket_address_get_flowinfo(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_flowinfo(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_inet_socket_address_get_scope_id(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_scope_id(cValueInstance)

	return (uint32)(ret)
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

func Fn_g_menu_new() unsafe.Pointer {
	ret := C.g_menu_new()

	return unsafe.Pointer(ret)
}

func Fn_g_menu_append(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_append(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_append_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuItem)(unsafe.Pointer(param0))

	C.g_menu_append_item(cValueInstance, cValue0)
}

func Fn_g_menu_append_section(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_append_section(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_append_submenu(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_append_submenu(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_freeze(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	C.g_menu_freeze(cValueInstance)
}

func Fn_g_menu_insert(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_menu_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_insert_item(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GMenuItem)(unsafe.Pointer(param1))

	C.g_menu_insert_item(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_insert_section(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GMenuModel)(unsafe.Pointer(param2))

	C.g_menu_insert_section(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_insert_submenu(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GMenuModel)(unsafe.Pointer(param2))

	C.g_menu_insert_submenu(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_prepend(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_prepend(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_prepend_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuItem)(unsafe.Pointer(param0))

	C.g_menu_prepend_item(cValueInstance, cValue0)
}

func Fn_g_menu_prepend_section(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_prepend_section(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_prepend_submenu(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_prepend_submenu(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_remove(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_menu_remove(cValueInstance, cValue0)
}

func Fn_g_menu_attribute_iter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_menu_attribute_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (**C.GVariant)(unsafe.Pointer(param1))

	ret := C.g_menu_attribute_iter_get_next(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_g_menu_attribute_iter_get_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_get_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_attribute_iter_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_next(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_menu_item_new(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_menu_item_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_new_section(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_item_new_section(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_new_submenu(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_item_new_submenu(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_set_action_and_target(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_g_menu_item_set_action_and_target(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_action_and_target_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_item_set_action_and_target_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_attribute(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_g_menu_item_set_attribute(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_item_set_attribute_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_detailed_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_menu_item_set_detailed_action(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_menu_item_set_label(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_link(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_item_set_link(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_section(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.g_menu_item_set_section(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_submenu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.g_menu_item_set_submenu(cValueInstance, cValue0)
}

func Fn_g_menu_link_iter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_menu_link_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (**C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_link_iter_get_next(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_g_menu_link_iter_get_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_get_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_link_iter_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_next(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_menu_model_get_item_attribute(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) bool {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.c_g_menu_model_get_item_attribute(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_g_menu_model_get_item_attribute_value(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GVariantType)(unsafe.Pointer(param2))

	ret := C.g_menu_model_get_item_attribute_value(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_get_item_link(paramInstance unsafe.Pointer, param0 int, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_menu_model_get_item_link(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_get_n_items(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_model_get_n_items(cValueInstance)

	return (int)(ret)
}

func Fn_g_menu_model_is_mutable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_model_is_mutable(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_menu_model_items_changed(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.g_menu_model_items_changed(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_model_iterate_item_attributes(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.g_menu_model_iterate_item_attributes(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_iterate_item_links(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.g_menu_model_iterate_item_links(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_output_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_flush_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_splice_async : parameter 'callback' is callback

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

func Fn_g_settings_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.GSettingsSchema)(unsafe.Pointer(param0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_settings_new_full(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

func Fn_g_settings_create_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_create_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

// UNSUPPORTED : g_settings_get_strv : no array length

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

func Fn_g_simple_async_result_set_check_cancellable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_simple_async_result_set_check_cancellable(cValueInstance, cValue0)
}

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

func Fn_g_socket_condition_timed_wait(paramInstance unsafe.Pointer, param0 int, param1 int64, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (C.gint64)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_condition_timed_wait(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_get_available_bytes(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_available_bytes(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_get_broadcast(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_broadcast(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_multicast_loopback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_multicast_loopback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_multicast_ttl(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_multicast_ttl(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_get_ttl(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_ttl(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_join_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cError := (**C.GError)(error)

	ret := C.g_socket_join_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_leave_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cError := (**C.GError)(error)

	ret := C.g_socket_leave_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_set_broadcast(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_socket_set_broadcast(cValueInstance, cValue0)
}

func Fn_g_socket_set_multicast_loopback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_socket_set_multicast_loopback(cValueInstance, cValue0)
}

func Fn_g_socket_set_multicast_ttl(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_socket_set_multicast_ttl(cValueInstance, cValue0)
}

func Fn_g_socket_set_ttl(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_socket_set_ttl(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

func Fn_g_socket_connection_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_connect(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

func Fn_g_socket_connection_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_connect_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connection_is_connected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_is_connected(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_newv : parameter 'argv' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_communicate_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_communicate_utf8_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_check_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : parameter 'child_setup' is callback

// UNSUPPORTED : g_subprocess_launcher_set_environ : parameter 'env' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_launcher_spawnv : parameter 'argv' is array parameter without length parameter

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

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_receive_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_credentials_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_send_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_credentials_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

// UNSUPPORTED : g_action_group_list_actions : no array length

func Fn_g_action_group_query_action(paramInstance unsafe.Pointer, param0 string, param1 *bool, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) bool {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gboolean)(unsafe.Pointer(param1))

	cValue2 := (**C.GVariantType)(unsafe.Pointer(param2))

	cValue3 := (**C.GVariantType)(unsafe.Pointer(param3))

	cValue4 := (**C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (**C.GVariant)(unsafe.Pointer(param5))

	ret := C.g_action_group_query_action(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_g_action_map_add_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAction)(unsafe.Pointer(param0))

	C.g_action_map_add_action(cValueInstance, cValue0)
}

func Fn_g_action_map_add_action_entries(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GActionEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gpointer)(param2)

	C.g_action_map_add_action_entries(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_action_map_lookup_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_map_lookup_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_map_remove_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_action_map_remove_action(cValueInstance, cValue0)
}

// UNSUPPORTED : g_app_info_get_supported_types : no array length

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

func Fn_g_dbus_interface_dup_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_dup_object(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

func Fn_g_drive_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

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

func Fn_g_mount_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

func Fn_g_network_monitor_can_reach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

func Fn_g_network_monitor_get_network_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_network_monitor_get_network_available(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

func Fn_g_remote_action_group_activate_action_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GRemoteActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_remote_action_group_activate_action_full(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_remote_action_group_change_action_state_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GRemoteActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_remote_action_group_change_action_state_full(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

func Fn_g_volume_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback
