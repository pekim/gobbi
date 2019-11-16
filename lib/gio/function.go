// Code generated - DO NOT EDIT.

package gio

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'g_action_name_is_valid' : parameter 'action_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : parameter 'detailed_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_action_print_detailed_name' : parameter 'action_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : parameter 'commandline' of type 'filename' not supported

// UNSUPPORTED : C value 'g_app_info_get_all' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : parameter 'content_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : parameter 'content_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : parameter 'uri_scheme' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : parameter 'content_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : parameter 'content_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri' : parameter 'uri' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : parameter 'uri' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_finish' : parameter 'result' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_app_info_reset_type_associations' : parameter 'content_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_async_initable_newv_async' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_bus_get' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_get_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_bus_get_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_own_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

var busUnownNameInvoker *gi.Function

// BusUnownName is a representation of the C type g_bus_unown_name.
func BusUnownName(ownerId uint32) {
	if busUnownNameInvoker == nil {
		busUnownNameInvoker = gi.FunctionInvokerNew("Gio", "bus_unown_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(ownerId)

	busUnownNameInvoker.Invoke(inArgs[:])
}

var busUnwatchNameInvoker *gi.Function

// BusUnwatchName is a representation of the C type g_bus_unwatch_name.
func BusUnwatchName(watcherId uint32) {
	if busUnwatchNameInvoker == nil {
		busUnwatchNameInvoker = gi.FunctionInvokerNew("Gio", "bus_unwatch_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(watcherId)

	busUnwatchNameInvoker.Invoke(inArgs[:])
}

// UNSUPPORTED : C value 'g_bus_watch_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_content_type_can_be_executable' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_equals' : parameter 'type1' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_from_mime_type' : parameter 'mime_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_get_description' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_get_generic_icon_name' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_get_icon' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_get_mime_type' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_guess' : parameter 'filename' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_guess_for_tree' : parameter 'root' of type 'File' not supported

// UNSUPPORTED : C value 'g_content_type_is_a' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_is_mime_type' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_type_is_unknown' : parameter 'type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_content_types_get_registered' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_dbus_address_escape_value' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : parameter 'address' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_sync' : parameter 'address' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : parameter 'annotations' has no type

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : parameter 'dbus_error_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : parameter 'error_domain_quark_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_unregister_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

var dbusGenerateGuidInvoker *gi.Function

// DbusGenerateGuid is a representation of the C type g_dbus_generate_guid.
func DbusGenerateGuid() string {
	if dbusGenerateGuidInvoker == nil {
		dbusGenerateGuidInvoker = gi.FunctionInvokerNew("Gio", "dbus_generate_guid")
	}

	ret := dbusGenerateGuidInvoker.Invoke(nil)
	return ret.String(true)
}

// UNSUPPORTED : C value 'g_dbus_gvalue_to_gvariant' : parameter 'gvalue' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'g_dbus_gvariant_to_gvalue' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_is_address' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_guid' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_interface_name' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_member_name' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_name' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_supported_address' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dbus_is_unique_name' : parameter 'string' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_path' : parameter 'path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_uri' : parameter 'uri' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_file_new_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_parse_name' : parameter 'parse_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_icon_deserialize' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_icon_hash' : parameter 'icon' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_icon_new_for_string' : parameter 'str' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_initable_newv' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_io_error_from_errno' : return type 'IOErrorEnum' not supported

// UNSUPPORTED : C value 'g_io_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_io_extension_point_implement' : parameter 'extension_point_name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_io_extension_point_lookup' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_io_extension_point_register' : parameter 'name' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

var ioSchedulerCancelAllJobsInvoker *gi.Function

// IoSchedulerCancelAllJobs is a representation of the C type g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	if ioSchedulerCancelAllJobsInvoker == nil {
		ioSchedulerCancelAllJobsInvoker = gi.FunctionInvokerNew("Gio", "io_scheduler_cancel_all_jobs")
	}

	ioSchedulerCancelAllJobsInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : parameter 'job_func' of type 'IOSchedulerJobFunc' not supported

// UNSUPPORTED : C value 'g_keyfile_settings_backend_new' : parameter 'filename' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_memory_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_network_monitor_get_default' : return type 'NetworkMonitor' not supported

var networkingInitInvoker *gi.Function

// NetworkingInit is a representation of the C type g_networking_init.
func NetworkingInit() {
	if networkingInitInvoker == nil {
		networkingInitInvoker = gi.FunctionInvokerNew("Gio", "networking_init")
	}

	networkingInitInvoker.Invoke(nil)
}

// UNSUPPORTED : C value 'g_null_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_pollable_source_new' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_source_new_full' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_stream_read' : parameter 'stream' of type 'InputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : parameter 'protocol' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_proxy_resolver_get_default' : return type 'ProxyResolver' not supported

// UNSUPPORTED : C value 'g_resolver_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_load' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_resources_enumerate_children' : parameter 'path' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_resources_get_info' : parameter 'path' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_resources_lookup_data' : parameter 'path' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_resources_open_stream' : parameter 'path' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_resources_register' : parameter 'resource' of type 'Resource' not supported

// UNSUPPORTED : C value 'g_resources_unregister' : parameter 'resource' of type 'Resource' not supported

// UNSUPPORTED : C value 'g_settings_schema_source_get_default' : return type 'SettingsSchemaSource' not supported

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_srv_target_list_sort' : parameter 'targets' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_tls_backend_get_default' : return type 'TlsBackend' not supported

// UNSUPPORTED : C value 'g_tls_client_connection_new' : parameter 'base_io_stream' of type 'IOStream' not supported

// UNSUPPORTED : C value 'g_tls_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_tls_file_database_new' : parameter 'anchors' of type 'filename' not supported

// UNSUPPORTED : C value 'g_tls_server_connection_new' : parameter 'base_io_stream' of type 'IOStream' not supported

// UNSUPPORTED : C value 'g_unix_is_mount_path_system_internal' : parameter 'mount_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_is_system_device_path' : parameter 'device_path' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_unix_is_system_fs_type' : parameter 'fs_type' of type 'utf8' not supported

// UNSUPPORTED : C value 'g_unix_mount_at' : parameter 'mount_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_compare' : parameter 'mount1' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_copy' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_for' : parameter 'file_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_free' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_device_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_fs_type' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_mount_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_can_eject' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_icon' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_name' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_should_display' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_guess_symbolic_icon' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_is_readonly' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_is_system_internal' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_points_changed_since' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_points_get' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_unix_mounts_changed_since' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mounts_get' : return type 'GLib.List' not supported
