// Code generated - DO NOT EDIT.
// +build gio_2.38

package gio

import "unsafe"

// MENU_ATTRIBUTE_ICON is a representation of the C constant G_MENU_ATTRIBUTE_ICON.
//
// since 2.38
const MENU_ATTRIBUTE_ICON = "icon"

// FileMeasureFlags is a representation of the C bitfield GFileMeasureFlags.
type FileMeasureFlags int

// FileMeasureFlags_none is a representation of the C bitfield member G_FILE_MEASURE_NONE.
const FileMeasureFlags_none = FileMeasureFlags(0)

// FileMeasureFlags_report_any_error is a representation of the C bitfield member G_FILE_MEASURE_REPORT_ANY_ERROR.
const FileMeasureFlags_report_any_error = FileMeasureFlags(2)

// FileMeasureFlags_apparent_size is a representation of the C bitfield member G_FILE_MEASURE_APPARENT_SIZE.
const FileMeasureFlags_apparent_size = FileMeasureFlags(4)

// FileMeasureFlags_no_xdev is a representation of the C bitfield member G_FILE_MEASURE_NO_XDEV.
const FileMeasureFlags_no_xdev = FileMeasureFlags(8)

// UNSUPPORTED : g_action_parse_detailed_name : throws

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UNSUPPORTED : g_unix_mount_at : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_for : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// BytesIcon is a representation of the C record GBytesIcon.
//
// since 2.38
type BytesIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBytesIcon that represents the BytesIcon.
func (recv *BytesIcon) ToC() unsafe.Pointer {
	return recv.native
}

// BytesIconNewFromC creates a new BytesIcon from a pointer to the C GBytesIcon that represents the BytesIcon.
func BytesIconNewFromC(native unsafe.Pointer) *BytesIcon {
	return &BytesIcon{native: native}
}

// PropertyAction is a representation of the C record GPropertyAction.
//
// since 2.38
type PropertyAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPropertyAction that represents the PropertyAction.
func (recv *PropertyAction) ToC() unsafe.Pointer {
	return recv.native
}

// PropertyActionNewFromC creates a new PropertyAction from a pointer to the C GPropertyAction that represents the PropertyAction.
func PropertyActionNewFromC(native unsafe.Pointer) *PropertyAction {
	return &PropertyAction{native: native}
}
