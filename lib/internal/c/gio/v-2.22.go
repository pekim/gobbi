// Code generated - DO NOT EDIT.
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

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
import "C"

type AsyncInitableIface C.GAsyncInitableIface
type InitableIface C.GInitableIface
type InputVector C.GInputVector

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
type OutputVector C.GOutputVector

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

func Fn_g_srv_target_new(param0 string, param1 uint16, param2 uint16, param3 uint16) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	cValue3 := (C.guint16)(param3)

	ret := C.g_srv_target_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	C.g_srv_target_free(cValueInstance)
}

func Fn_g_srv_target_get_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_srv_target_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_priority(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_priority(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_weight(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_weight(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	ret := C.g_srv_target_list_sort(cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

// UNSUPPORTED : g_application_add_main_option_entries : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : g_application_command_line_get_environ : no array length

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

func Fn_g_cancellable_disconnect(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gulong)(param0)

	C.g_cancellable_disconnect(cValueInstance, cValue0)
}

func Fn_g_cancellable_make_pollfd(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GPollFD)(unsafe.Pointer(param0))

	ret := C.g_cancellable_make_pollfd(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_cancellable_release_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_release_fd(cValueInstance)
}

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

func Fn_g_file_io_stream_get_etag(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_file_io_stream_get_etag(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_io_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_io_stream_query_info_async : parameter 'callback' is callback

func Fn_g_file_io_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_info_get_attribute_stringv : no array length

func Fn_g_file_info_has_namespace(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_has_namespace(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_info_list_attributes : no array length

func Fn_g_file_info_set_attribute_status(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileAttributeStatus)(param1)

	ret := C.g_file_info_set_attribute_status(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_info_set_attribute_stringv : parameter 'attr_value' is array parameter without length parameter

// UNSUPPORTED : g_file_input_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_output_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_filename_completer_get_completions : no array length

// UNSUPPORTED : g_io_module_new : blacklisted

// UNSUPPORTED : g_io_module_load : blacklisted

// UNSUPPORTED : g_io_module_unload : blacklisted

// UNSUPPORTED : g_io_module_query : blacklisted

func Fn_g_io_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_clear_pending(cValueInstance)
}

func Fn_g_io_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_close_async : parameter 'callback' is callback

func Fn_g_io_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_io_stream_get_input_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_get_input_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_get_output_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_get_output_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_io_stream_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_io_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_set_pending(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_splice_async : parameter 'callback' is callback

func Fn_g_inet_address_new_any(param0 int) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	ret := C.g_inet_address_new_any(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_inet_address_new_from_bytes : parameter 'bytes' is array parameter without length parameter

func Fn_g_inet_address_new_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_inet_address_new_from_string(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_new_loopback(param0 int) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	ret := C.g_inet_address_new_loopback(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_inet_address_get_is_any(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_any(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_link_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_link_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_loopback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_loopback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_global(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_global(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_link_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_link_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_node_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_node_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_org_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_org_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_site_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_site_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_multicast(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_multicast(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_site_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_site_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_native_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_native_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_inet_address_to_bytes(paramInstance unsafe.Pointer) *uint8 {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_to_bytes(cValueInstance)

	return (*uint8)(ret)
}

func Fn_g_inet_address_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_inet_socket_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_port(cValueInstance)

	return (uint16)(ret)
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

func Fn_g_network_address_new(param0 string, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_network_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_get_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_address_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_network_address_parse(param0 string, param1 uint16, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_new(param0 string, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_network_service_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_get_domain(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_domain(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_get_protocol(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_protocol(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_get_service(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_service(cValueInstance)

	return C.GoString(ret)
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

func Fn_g_resolver_lookup_by_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_address_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_by_address_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address_finish(cValueInstance, cValue0, cError)

	return C.GoString(ret)
}

func Fn_g_resolver_lookup_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_by_name_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_records_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_service_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resolver_set_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	C.g_resolver_set_default(cValueInstance)
}

func Fn_g_resolver_free_addresses(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.g_resolver_free_addresses(cValue0)
}

func Fn_g_resolver_free_targets(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.g_resolver_free_targets(cValue0)
}

func Fn_g_resolver_get_default() unsafe.Pointer {
	ret := C.g_resolver_get_default()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

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

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

func Fn_g_socket_new(param0 int, param1 int, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.GSocketProtocol)(param2)

	cError := (**C.GError)(error)

	ret := C.g_socket_new(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_new_from_fd(param0 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_socket_new_from_fd(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_accept(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_accept(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_bind(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_bind(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_check_connect_result(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_check_connect_result(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_socket_close(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_close(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_socket_condition_check(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	ret := C.g_socket_condition_check(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_socket_condition_wait(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_condition_wait(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_connect(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connection_factory_create_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_factory_create_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_create_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	ret := C.g_socket_create_source(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_blocking(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_blocking(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_keepalive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_keepalive(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_listen_backlog(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_listen_backlog(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_local_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_protocol(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_protocol(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_remote_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_socket_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_socket_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_is_connected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_is_connected(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_listen(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_listen(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_socket_receive(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

func Fn_g_socket_receive_from(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gsize)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive_from(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return (uint64)(ret)
}

func Fn_g_socket_receive_message(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 []InputVector, param2 int, param3 *[]unsafe.Pointer, param4 *int, param5 *int, param6 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GInputVector)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gint)(param2)

	var cValue3ArrayPointer (**C.GSocketControlMessage)
	cValue3 := &cValue3ArrayPointer

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive_message(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	param3OutLen := int(*cValue4)
	param3Out := make([]unsafe.Pointer, param3OutLen, param3OutLen)
	if param3OutLen > 0 {
		param3Out = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(cValue3ArrayPointer))[:param3OutLen:param3OutLen]
	}
	*param3 = param3Out

	return (uint64)(ret)
}

func Fn_g_socket_send(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_send(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_message(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []OutputVector, param2 int, param3 []unsafe.Pointer, param4 int, param5 int, param6 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GOutputVector)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gint)(param2)

	cValue3 := (**C.GSocketControlMessage)(unsafe.Pointer(&param3[0]))

	cValue4 := (C.gint)(param4)

	cValue5 := (C.gint)(param5)

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_message(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gsize)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_to(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return (uint64)(ret)
}

func Fn_g_socket_set_blocking(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_socket_set_blocking(cValueInstance, cValue0)
}

func Fn_g_socket_set_keepalive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_socket_set_keepalive(cValueInstance, cValue0)
}

func Fn_g_socket_set_listen_backlog(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_socket_set_listen_backlog(cValueInstance, cValue0)
}

func Fn_g_socket_shutdown(paramInstance unsafe.Pointer, param0 bool, param1 bool, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_shutdown(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_speaks_ipv4(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_speaks_ipv4(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_address_new_from_native(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gsize)(param1)

	ret := C.g_socket_address_new_from_native(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_address_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_address_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_address_get_native_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_address_get_native_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_address_to_native(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_address_to_native(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

func Fn_g_socket_client_new() unsafe.Pointer {
	ret := C.g_socket_client_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_host(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_host_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_service_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

func Fn_g_socket_client_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_get_local_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_local_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_protocol(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_protocol(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_get_socket_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_socket_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_set_family(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSocketFamily)(param0)

	C.g_socket_client_set_family(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_local_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	C.g_socket_client_set_local_address(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_protocol(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSocketProtocol)(param0)

	C.g_socket_client_set_protocol(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_socket_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSocketType)(param0)

	C.g_socket_client_set_socket_type(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

func Fn_g_socket_connection_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_local_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_remote_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_socket(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_get_socket(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_factory_lookup_type(param0 int, param1 int, param2 int) uint64 {
	cValue0 := (C.GSocketFamily)(param0)

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.gint)(param2)

	ret := C.g_socket_connection_factory_lookup_type(cValue0, cValue1, cValue2)

	return (uint64)(ret)
}

func Fn_g_socket_connection_factory_register_type(param0 uint64, param1 int, param2 int, param3 int) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GSocketFamily)(param1)

	cValue2 := (C.GSocketType)(param2)

	cValue3 := (C.gint)(param3)

	C.g_socket_connection_factory_register_type(cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_socket_control_message_get_level(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_level(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_msg_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_msg_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_control_message_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.g_socket_control_message_serialize(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

func Fn_g_socket_listener_new() unsafe.Pointer {
	ret := C.g_socket_listener_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_accept(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

func Fn_g_socket_listener_accept_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_accept_socket(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_socket(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

func Fn_g_socket_listener_accept_socket_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_socket_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_add_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.GSocketProtocol)(param2)

	cValue3 := (*C.GObject)(unsafe.Pointer(param3))

	cValue4 := (**C.GSocketAddress)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_address(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_add_inet_port(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_inet_port(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_add_socket(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocket)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_socket(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	C.g_socket_listener_close(cValueInstance)
}

func Fn_g_socket_listener_set_backlog(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	C.g_socket_listener_set_backlog(cValueInstance, cValue0)
}

func Fn_g_socket_service_new() unsafe.Pointer {
	ret := C.g_socket_service_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_service_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_service_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_service_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_start(cValueInstance)
}

func Fn_g_socket_service_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_stop(cValueInstance)
}

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

func Fn_g_tcp_connection_get_graceful_disconnect(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tcp_connection_get_graceful_disconnect(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tcp_connection_set_graceful_disconnect(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_tcp_connection_set_graceful_disconnect(cValueInstance, cValue0)
}

// UNSUPPORTED : g_themed_icon_get_names : no array length

func Fn_g_threaded_socket_service_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.g_threaded_socket_service_new(cValue0)

	return unsafe.Pointer(ret)
}

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

func Fn_g_unix_connection_receive_fd(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_fd(cValueInstance, cValue0, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_send_fd(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_fd(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_unix_fd_message_new() unsafe.Pointer {
	ret := C.g_unix_fd_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_append_fd(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_message_append_fd(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_unix_fd_message_steal_fds(paramInstance unsafe.Pointer, param0 *int) []int {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.g_unix_fd_message_steal_fds(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_unix_socket_address_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_unix_socket_address_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_get_is_abstract(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_is_abstract(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_socket_address_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_socket_address_get_path_len(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_path_len(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_unix_socket_address_abstract_names_supported() bool {
	ret := C.g_unix_socket_address_abstract_names_supported()

	return toGoBool(ret)
}

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

// UNSUPPORTED : g_action_group_list_actions : no array length

// UNSUPPORTED : g_app_info_get_supported_types : no array length

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

func Fn_g_async_initable_init_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAsyncInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_init_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_async_initable_new_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAsyncInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_new_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_can_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_start_degraded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_start_degraded(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_stop(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_stop(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

func Fn_g_drive_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

func Fn_g_drive_get_start_stop_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_start_stop_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : g_drive_poll_for_media : parameter 'callback' is callback

// UNSUPPORTED : g_drive_start : parameter 'callback' is callback

func Fn_g_drive_start_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_start_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_stop : parameter 'callback' is callback

func Fn_g_drive_stop_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_stop_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dtls_connection_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_dtls_connection_shutdown_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_append_to_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_copy : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_copy_async : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_create_async : parameter 'callback' is callback

func Fn_g_file_create_readwrite(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileCreateFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_readwrite_async : parameter 'callback' is callback

func Fn_g_file_create_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_delete_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable_with_operation : parameter 'callback' is callback

func Fn_g_file_eject_mountable_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

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

func Fn_g_file_open_readwrite(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_open_readwrite_async : parameter 'callback' is callback

func Fn_g_file_open_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_poll_mountable : parameter 'callback' is callback

func Fn_g_file_poll_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_poll_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_query_default_handler_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_filesystem_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : parameter 'callback' is callback

func Fn_g_file_replace_readwrite(paramInstance unsafe.Pointer, param0 string, param1 bool, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	cValue2 := (C.GFileCreateFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_replace_readwrite_async : parameter 'callback' is callback

func Fn_g_file_replace_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_set_attributes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_display_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_start_mountable : parameter 'callback' is callback

func Fn_g_file_start_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_start_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_stop_mountable : parameter 'callback' is callback

func Fn_g_file_stop_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_stop_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_supports_thread_contexts(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_supports_thread_contexts(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_trash_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable_with_operation : parameter 'callback' is callback

func Fn_g_file_unmount_mountable_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_initable_init(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_initable_init(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_loadable_icon_load_async : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject_with_operation : parameter 'callback' is callback

func Fn_g_mount_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

func Fn_g_mount_unmount_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

func Fn_g_socket_connectable_enumerate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnectable)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connectable_enumerate(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

func Fn_g_volume_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback
