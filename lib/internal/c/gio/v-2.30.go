// Code generated - DO NOT EDIT.
// +build gio_2.30

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

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type DBusInterfaceIface C.GDBusInterfaceIface
type DBusInterfaceSkeletonClass C.GDBusInterfaceSkeletonClass
type DBusObjectIface C.GDBusObjectIface
type DBusObjectManagerClientClass C.GDBusObjectManagerClientClass
type DBusObjectManagerIface C.GDBusObjectManagerIface
type DBusObjectManagerServerClass C.GDBusObjectManagerServerClass
type DBusObjectProxyClass C.GDBusObjectProxyClass
type DBusObjectSkeletonClass C.GDBusObjectSkeletonClass
type IOModuleScope C.GIOModuleScope

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type TlsDatabaseClass C.GTlsDatabaseClass
type TlsInteractionClass C.GTlsInteractionClass

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

func Fn_g_dbus_interface_info_cache_build(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_cache_build(cValueInstance)
}

func Fn_g_dbus_interface_info_cache_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_cache_release(cValueInstance)
}

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

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

func Fn_g_dbus_gvalue_to_gvariant(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	ret := C.g_dbus_gvalue_to_gvariant(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_gvariant_to_gvalue(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_dbus_gvariant_to_gvalue(cValue0, cValue1)
}

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

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

func Fn_g_dbus_connection_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, param9 *unsafe.Pointer, param10 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (*C.GVariantType)(unsafe.Pointer(param5))

	cValue6 := (C.GDBusCallFlags)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (*C.GUnixFDList)(unsafe.Pointer(param8))

	cValue9 := (**C.GUnixFDList)(unsafe.Pointer(param9))

	cValue10 := (*C.GCancellable)(unsafe.Pointer(param10))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

func Fn_g_dbus_interface_skeleton_export(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_interface_skeleton_export(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_interface_skeleton_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_flush(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_interface_skeleton_get_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_interface_skeleton_get_properties(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_properties(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_vtable(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_vtable(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusInterfaceSkeletonFlags)(param0)

	C.g_dbus_interface_skeleton_set_flags(cValueInstance, cValue0)
}

func Fn_g_dbus_interface_skeleton_unexport(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_unexport(cValueInstance)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : no array length

func Fn_g_dbus_method_invocation_return_value_with_unix_fd_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	cValue1 := (*C.GUnixFDList)(unsafe.Pointer(param1))

	C.g_dbus_method_invocation_return_value_with_unix_fd_list(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_method_invocation_take_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_take_error(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_client_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_object_manager_client_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_client_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_object_manager_client_new_for_bus_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : parameter 'get_proxy_type_func' is callback

func Fn_g_dbus_object_manager_client_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_client_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_object_manager_client_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_client_get_name_owner(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_name_owner(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_dbus_object_manager_client_new : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : parameter 'get_proxy_type_func' is callback

func Fn_g_dbus_object_manager_server_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_server_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_server_export(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_export(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_export_uniquely(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_export_uniquely(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_server_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_server_unexport(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_server_unexport(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_object_proxy_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_object_proxy_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_proxy_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_proxy_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_skeleton_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_skeleton_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_skeleton_add_interface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_skeleton_add_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_skeleton_flush(cValueInstance)
}

func Fn_g_dbus_object_skeleton_remove_interface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_skeleton_remove_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_remove_interface_by_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_skeleton_remove_interface_by_name(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_set_object_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_skeleton_set_object_path(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_proxy_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : parameter 'callback' is callback

func Fn_g_dbus_proxy_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (C.GDBusCallFlags)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.GUnixFDList)(unsafe.Pointer(param4))

	cValue5 := (**C.GUnixFDList)(unsafe.Pointer(param5))

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : no array length

// UNSUPPORTED : g_dbus_proxy_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line : no array length

// UNSUPPORTED : g_data_input_stream_read_line_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : no array length

func Fn_g_data_input_stream_read_line_finish_utf8(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_line_finish_utf8(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_read_line_utf8(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_line_utf8(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_until_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_upto_async : parameter 'callback' is callback

// UNSUPPORTED : g_desktop_app_info_get_keywords : no array length

func Fn_g_desktop_app_info_get_nodisplay(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_nodisplay(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_desktop_app_info_get_show_in(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_show_in(cValueInstance, cValue0)

	return toGoBool(ret)
}

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

func Fn_g_inet_address_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	ret := C.g_inet_address_equal(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

// UNSUPPORTED : g_settings_get_strv : no array length

func Fn_g_settings_get_uint(paramInstance unsafe.Pointer, param0 string) uint {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_uint(cValueInstance, cValue0)

	return (uint)(ret)
}

// UNSUPPORTED : g_settings_list_children : no array length

// UNSUPPORTED : g_settings_list_keys : no array length

// UNSUPPORTED : g_settings_set_strv : parameter 'value' is array parameter without length parameter

func Fn_g_settings_set_uint(paramInstance unsafe.Pointer, param0 string, param1 uint) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	ret := C.g_settings_set_uint(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

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

func Fn_g_simple_action_set_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_simple_action_set_state(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_add_entries(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GActionEntry)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gpointer)(param2)

	C.g_simple_action_group_add_entries(cValueInstance, cValue0, cValue1, cValue2)
}

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

func Fn_g_tls_connection_get_database(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_database(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_interaction(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_interaction(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

func Fn_g_tls_connection_set_database(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsDatabase)(unsafe.Pointer(param0))

	C.g_tls_connection_set_database(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_interaction(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsInteraction)(unsafe.Pointer(param0))

	C.g_tls_connection_set_interaction(cValueInstance, cValue0)
}

func Fn_g_tls_database_create_certificate_handle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	ret := C.g_tls_database_create_certificate_handle(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_tls_database_lookup_certificate_for_handle(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))

	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_for_handle(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : parameter 'callback' is callback

func Fn_g_tls_database_lookup_certificate_for_handle_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_for_handle_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_database_lookup_certificate_issuer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))

	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_issuer(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : parameter 'callback' is callback

func Fn_g_tls_database_lookup_certificate_issuer_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_issuer_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : parameter 'issuer_raw_dn' is array parameter without length parameter

func Fn_g_tls_database_lookup_certificates_issued_by_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificates_issued_by_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_database_verify_chain(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int, param5 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GSocketConnectable)(unsafe.Pointer(param2))

	cValue3 := (*C.GTlsInteraction)(unsafe.Pointer(param3))

	cValue4 := (C.GTlsDatabaseVerifyFlags)(param4)

	cValue5 := (*C.GCancellable)(unsafe.Pointer(param5))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_verify_chain(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_database_verify_chain_async : parameter 'callback' is callback

func Fn_g_tls_database_verify_chain_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_verify_chain_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_ask_password(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_interaction_ask_password_async : parameter 'callback' is callback

func Fn_g_tls_interaction_ask_password_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_ask_password_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_invoke_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_invoke_ask_password(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

func Fn_g_tls_password_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_tls_password_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_tls_password_get_value(paramInstance unsafe.Pointer, param0 *uint64) *uint8 {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.g_tls_password_get_value(cValueInstance, cValue0)

	return (*uint8)(ret)
}

func Fn_g_tls_password_get_warning(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_warning(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_tls_password_set_description(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_tls_password_set_description(cValueInstance, cValue0)
}

func Fn_g_tls_password_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GTlsPasswordFlags)(param0)

	C.g_tls_password_set_flags(cValueInstance, cValue0)
}

func Fn_g_tls_password_set_value(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gssize)(param1)

	C.g_tls_password_set_value(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

func Fn_g_tls_password_set_warning(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_tls_password_set_warning(cValueInstance, cValue0)
}

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

func Fn_g_action_change_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_action_change_state(cValueInstance, cValue0)
}

// UNSUPPORTED : g_action_group_list_actions : no array length

// UNSUPPORTED : g_app_info_get_supported_types : no array length

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

func Fn_g_dbus_interface_get_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_get_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_set_object(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusObject)(unsafe.Pointer(param0))

	C.g_dbus_interface_set_object(cValueInstance, cValue0)
}

func Fn_g_dbus_object_get_interface(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_get_interface(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_get_interfaces(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_get_interfaces(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_get_interface(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_object_manager_get_interface(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_get_object(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_get_objects(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_get_objects(cValueInstance)

	return unsafe.Pointer(ret)
}

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

func Fn_g_tls_backend_get_default_database(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_default_database(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_backend_get_file_database_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_file_database_type(cValueInstance)

	return (uint64)(ret)
}

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback
