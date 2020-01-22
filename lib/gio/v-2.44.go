// Code generated - DO NOT EDIT.
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

// NetworkConnectivity is a representation of the C enumeration GNetworkConnectivity.
type NetworkConnectivity int

// NetworkConnectivity_local is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_LOCAL.
const NetworkConnectivity_local = NetworkConnectivity(1)

// NetworkConnectivity_limited is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_LIMITED.
const NetworkConnectivity_limited = NetworkConnectivity(2)

// NetworkConnectivity_portal is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_PORTAL.
const NetworkConnectivity_portal = NetworkConnectivity(3)

// NetworkConnectivity_full is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_FULL.
const NetworkConnectivity_full = NetworkConnectivity(4)

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

// ListModelInterface is a representation of the C record GListModelInterface.
//
// since 2.44
type ListModelInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListModelInterface that represents the ListModelInterface.
func (recv *ListModelInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ListModelInterfaceNewFromC creates a new ListModelInterface from a pointer to the C GListModelInterface that represents the ListModelInterface.
func ListModelInterfaceNewFromC(native unsafe.Pointer) *ListModelInterface {
	return &ListModelInterface{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// OutputMessage is a representation of the C record GOutputMessage.
//
// since 2.44
type OutputMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputMessage that represents the OutputMessage.
func (recv *OutputMessage) ToC() unsafe.Pointer {
	return recv.native
}

// OutputMessageNewFromC creates a new OutputMessage from a pointer to the C GOutputMessage that represents the OutputMessage.
func OutputMessageNewFromC(native unsafe.Pointer) *OutputMessage {
	return &OutputMessage{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// SimpleIOStream is a representation of the C record GSimpleIOStream.
//
// since 2.44
type SimpleIOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleIOStream that represents the SimpleIOStream.
func (recv *SimpleIOStream) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleIOStreamNewFromC creates a new SimpleIOStream from a pointer to the C GSimpleIOStream that represents the SimpleIOStream.
func SimpleIOStreamNewFromC(native unsafe.Pointer) *SimpleIOStream {
	return &SimpleIOStream{native: native}
}
