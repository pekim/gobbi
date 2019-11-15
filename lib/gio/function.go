// Code generated - DO NOT EDIT.

package gio

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'g_action_name_is_valid' : has parameters

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : has parameters

// UNSUPPORTED : C value 'g_action_print_detailed_name' : has parameters

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : has parameters

// UNSUPPORTED : C value 'g_app_info_get_all' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : has parameters

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : has parameters

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : has parameters

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : has parameters

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : has parameters

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri' : has parameters

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : has parameters

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_finish' : has parameters

// UNSUPPORTED : C value 'g_app_info_reset_type_associations' : has parameters

// UNSUPPORTED : C value 'g_async_initable_newv_async' : has parameters

// UNSUPPORTED : C value 'g_bus_get' : has parameters

// UNSUPPORTED : C value 'g_bus_get_finish' : has parameters

// UNSUPPORTED : C value 'g_bus_get_sync' : has parameters

// UNSUPPORTED : C value 'g_bus_own_name' : has parameters

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : has parameters

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : has parameters

// UNSUPPORTED : C value 'g_bus_own_name_with_closures' : has parameters

// UNSUPPORTED : C value 'g_bus_unown_name' : has parameters

// UNSUPPORTED : C value 'g_bus_unwatch_name' : has parameters

// UNSUPPORTED : C value 'g_bus_watch_name' : has parameters

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : has parameters

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : has parameters

// UNSUPPORTED : C value 'g_bus_watch_name_with_closures' : has parameters

// UNSUPPORTED : C value 'g_content_type_can_be_executable' : has parameters

// UNSUPPORTED : C value 'g_content_type_equals' : has parameters

// UNSUPPORTED : C value 'g_content_type_from_mime_type' : has parameters

// UNSUPPORTED : C value 'g_content_type_get_description' : has parameters

// UNSUPPORTED : C value 'g_content_type_get_generic_icon_name' : has parameters

// UNSUPPORTED : C value 'g_content_type_get_icon' : has parameters

// UNSUPPORTED : C value 'g_content_type_get_mime_type' : has parameters

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : has parameters

// UNSUPPORTED : C value 'g_content_type_guess' : has parameters

// UNSUPPORTED : C value 'g_content_type_guess_for_tree' : has parameters

// UNSUPPORTED : C value 'g_content_type_is_a' : has parameters

// UNSUPPORTED : C value 'g_content_type_is_mime_type' : has parameters

// UNSUPPORTED : C value 'g_content_type_is_unknown' : has parameters

// UNSUPPORTED : C value 'g_content_types_get_registered' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_dbus_address_escape_value' : has parameters

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : has parameters

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : has parameters

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : has parameters

// UNSUPPORTED : C value 'g_dbus_address_get_stream_sync' : has parameters

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : has parameters

// UNSUPPORTED : C value 'g_dbus_error_unregister_error' : has parameters

var dbusGenerateGuidInvoker *gi.Function

// DbusGenerateGuid is a representation of the C type g_dbus_generate_guid.
func DbusGenerateGuid() string {
	if dbusGenerateGuidInvoker == nil {
		dbusGenerateGuidInvoker = gi.FunctionInvokerNew("Gio", "dbus_generate_guid")
	}

	ret := dbusGenerateGuidInvoker.Invoke()
	return ret.String()
}

// UNSUPPORTED : C value 'g_dbus_gvalue_to_gvariant' : has parameters

// UNSUPPORTED : C value 'g_dbus_gvariant_to_gvalue' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_address' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_guid' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_interface_name' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_member_name' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_name' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_supported_address' : has parameters

// UNSUPPORTED : C value 'g_dbus_is_unique_name' : has parameters

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : has parameters

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : has parameters

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : has parameters

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : has parameters

// UNSUPPORTED : C value 'g_file_new_for_path' : has parameters

// UNSUPPORTED : C value 'g_file_new_for_uri' : has parameters

// UNSUPPORTED : C value 'g_file_new_tmp' : has parameters

// UNSUPPORTED : C value 'g_file_parse_name' : has parameters

// UNSUPPORTED : C value 'g_icon_deserialize' : has parameters

// UNSUPPORTED : C value 'g_icon_hash' : has parameters

// UNSUPPORTED : C value 'g_icon_new_for_string' : has parameters

// UNSUPPORTED : C value 'g_initable_newv' : has parameters

// UNSUPPORTED : C value 'g_io_error_from_errno' : has parameters

// UNSUPPORTED : C value 'g_io_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_io_extension_point_implement' : has parameters

// UNSUPPORTED : C value 'g_io_extension_point_lookup' : has parameters

// UNSUPPORTED : C value 'g_io_extension_point_register' : has parameters

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : has parameters

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : has parameters

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory' : has parameters

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory_with_scope' : has parameters

// UNSUPPORTED : C value 'g_io_scheduler_cancel_all_jobs' : return type 'none' not supported

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : has parameters

// UNSUPPORTED : C value 'g_keyfile_settings_backend_new' : has parameters

// UNSUPPORTED : C value 'g_memory_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_network_monitor_get_default' : return type 'NetworkMonitor' not supported

// UNSUPPORTED : C value 'g_networking_init' : return type 'none' not supported

// UNSUPPORTED : C value 'g_null_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_pollable_source_new' : has parameters

// UNSUPPORTED : C value 'g_pollable_source_new_full' : has parameters

// UNSUPPORTED : C value 'g_pollable_stream_read' : has parameters

// UNSUPPORTED : C value 'g_pollable_stream_write' : has parameters

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : has parameters

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : has parameters

// UNSUPPORTED : C value 'g_proxy_resolver_get_default' : return type 'ProxyResolver' not supported

// UNSUPPORTED : C value 'g_resolver_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_load' : has parameters

// UNSUPPORTED : C value 'g_resources_enumerate_children' : has parameters

// UNSUPPORTED : C value 'g_resources_get_info' : has parameters

// UNSUPPORTED : C value 'g_resources_lookup_data' : has parameters

// UNSUPPORTED : C value 'g_resources_open_stream' : has parameters

// UNSUPPORTED : C value 'g_resources_register' : has parameters

// UNSUPPORTED : C value 'g_resources_unregister' : has parameters

// UNSUPPORTED : C value 'g_settings_schema_source_get_default' : return type 'SettingsSchemaSource' not supported

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : has parameters

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : has parameters

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : has parameters

// UNSUPPORTED : C value 'g_srv_target_list_sort' : has parameters

// UNSUPPORTED : C value 'g_tls_backend_get_default' : return type 'TlsBackend' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_new' : has parameters

// UNSUPPORTED : C value 'g_tls_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_tls_file_database_new' : has parameters

// UNSUPPORTED : C value 'g_tls_server_connection_new' : has parameters

// UNSUPPORTED : C value 'g_unix_is_mount_path_system_internal' : has parameters

// UNSUPPORTED : C value 'g_unix_is_system_device_path' : has parameters

// UNSUPPORTED : C value 'g_unix_is_system_fs_type' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_at' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_compare' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_copy' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_for' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_free' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_get_device_path' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_get_fs_type' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_get_mount_path' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_guess_can_eject' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_guess_icon' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_guess_name' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_guess_should_display' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_guess_symbolic_icon' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_is_readonly' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_is_system_internal' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_points_changed_since' : has parameters

// UNSUPPORTED : C value 'g_unix_mount_points_get' : has parameters

// UNSUPPORTED : C value 'g_unix_mounts_changed_since' : has parameters

// UNSUPPORTED : C value 'g_unix_mounts_get' : has parameters
