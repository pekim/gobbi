// Code generated - DO NOT EDIT.
// +build gio_2.60 gio_2.62

package gio

import gio "github.com/pekim/gobbi/lib/internal/c/gio"

// FILE_ATTRIBUTE_DOS_IS_MOUNTPOINT is a representation of the C constant G_FILE_ATTRIBUTE_DOS_IS_MOUNTPOINT.
//
// since 2.60
const FILE_ATTRIBUTE_DOS_IS_MOUNTPOINT = "dos::is-mountpoint"

// FILE_ATTRIBUTE_DOS_REPARSE_POINT_TAG is a representation of the C constant G_FILE_ATTRIBUTE_DOS_REPARSE_POINT_TAG.
//
// since 2.60
const FILE_ATTRIBUTE_DOS_REPARSE_POINT_TAG = "dos::reparse-point-tag"

// ResolverNameLookupFlags is a representation of the C bitfield GResolverNameLookupFlags.
type ResolverNameLookupFlags int

// ResolverNameLookupFlags_default is a representation of the C bitfield member G_RESOLVER_NAME_LOOKUP_FLAGS_DEFAULT.
const ResolverNameLookupFlags_default = ResolverNameLookupFlags(0)

// ResolverNameLookupFlags_ipv4_only is a representation of the C bitfield member G_RESOLVER_NAME_LOOKUP_FLAGS_IPV4_ONLY.
const ResolverNameLookupFlags_ipv4_only = ResolverNameLookupFlags(1)

// ResolverNameLookupFlags_ipv6_only is a representation of the C bitfield member G_RESOLVER_NAME_LOOKUP_FLAGS_IPV6_ONLY.
const ResolverNameLookupFlags_ipv6_only = ResolverNameLookupFlags(2)

// PollableReturn is a representation of the C enumeration GPollableReturn.
type PollableReturn int

// PollableReturn_failed is a representation of the C enumeration member G_POLLABLE_RETURN_FAILED.
const PollableReturn_failed = PollableReturn(0)

// PollableReturn_ok is a representation of the C enumeration member G_POLLABLE_RETURN_OK.
const PollableReturn_ok = PollableReturn(1)

// PollableReturn_would_block is a representation of the C enumeration member G_POLLABLE_RETURN_WOULD_BLOCK.
const PollableReturn_would_block = PollableReturn(-27)

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

// UnixMountGetRootPath wraps the C function g_unix_mount_get_root_path.
//
// since 2.60
func UnixMountGetRootPath(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_root_path(sys_mountEntry)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted
