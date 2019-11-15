// Code generated - DO NOT EDIT.

package gio

import gi "github.com/pekim/gobbi/internal/gi"

// UNSUPPORTED : C value 'g_action_name_is_valid' : non trivial function

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : non trivial function

// UNSUPPORTED : C value 'g_action_print_detailed_name' : non trivial function

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : non trivial function

var appInfoGetAllInvoker *gi.Function

// AppInfoGetAll is a representation of the C type g_app_info_get_all.
func AppInfoGetAll() {
	if appInfoGetAllInvoker == nil {
		appInfoGetAllInvoker = gi.FunctionInvokerNew("Gio", "app_info_get_all")
	}

	appInfoGetAllInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : non trivial function

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : non trivial function

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : non trivial function

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : non trivial function

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : non trivial function

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri' : non trivial function

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : non trivial function

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_finish' : non trivial function

// UNSUPPORTED : C value 'g_app_info_reset_type_associations' : non trivial function

// UNSUPPORTED : C value 'g_async_initable_newv_async' : non trivial function

// UNSUPPORTED : C value 'g_bus_get' : non trivial function

// UNSUPPORTED : C value 'g_bus_get_finish' : non trivial function

// UNSUPPORTED : C value 'g_bus_get_sync' : non trivial function

// UNSUPPORTED : C value 'g_bus_own_name' : non trivial function

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : non trivial function

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : non trivial function

// UNSUPPORTED : C value 'g_bus_own_name_with_closures' : non trivial function

// UNSUPPORTED : C value 'g_bus_unown_name' : non trivial function

// UNSUPPORTED : C value 'g_bus_unwatch_name' : non trivial function

// UNSUPPORTED : C value 'g_bus_watch_name' : non trivial function

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : non trivial function

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : non trivial function

// UNSUPPORTED : C value 'g_bus_watch_name_with_closures' : non trivial function

// UNSUPPORTED : C value 'g_content_type_can_be_executable' : non trivial function

// UNSUPPORTED : C value 'g_content_type_equals' : non trivial function

// UNSUPPORTED : C value 'g_content_type_from_mime_type' : non trivial function

// UNSUPPORTED : C value 'g_content_type_get_description' : non trivial function

// UNSUPPORTED : C value 'g_content_type_get_generic_icon_name' : non trivial function

// UNSUPPORTED : C value 'g_content_type_get_icon' : non trivial function

// UNSUPPORTED : C value 'g_content_type_get_mime_type' : non trivial function

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : non trivial function

// UNSUPPORTED : C value 'g_content_type_guess' : non trivial function

// UNSUPPORTED : C value 'g_content_type_guess_for_tree' : non trivial function

// UNSUPPORTED : C value 'g_content_type_is_a' : non trivial function

// UNSUPPORTED : C value 'g_content_type_is_mime_type' : non trivial function

// UNSUPPORTED : C value 'g_content_type_is_unknown' : non trivial function

var contentTypesGetRegisteredInvoker *gi.Function

// ContentTypesGetRegistered is a representation of the C type g_content_types_get_registered.
func ContentTypesGetRegistered() {
	if contentTypesGetRegisteredInvoker == nil {
		contentTypesGetRegisteredInvoker = gi.FunctionInvokerNew("Gio", "content_types_get_registered")
	}

	contentTypesGetRegisteredInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_dbus_address_escape_value' : non trivial function

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : non trivial function

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : non trivial function

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : non trivial function

// UNSUPPORTED : C value 'g_dbus_address_get_stream_sync' : non trivial function

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : non trivial function

var dbusErrorQuarkInvoker *gi.Function

// DbusErrorQuark is a representation of the C type g_dbus_error_quark.
func DbusErrorQuark() {
	if dbusErrorQuarkInvoker == nil {
		dbusErrorQuarkInvoker = gi.FunctionInvokerNew("Gio", "dbus_error_quark")
	}

	dbusErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_dbus_error_register_error' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : non trivial function

// UNSUPPORTED : C value 'g_dbus_error_unregister_error' : non trivial function

var dbusGenerateGuidInvoker *gi.Function

// DbusGenerateGuid is a representation of the C type g_dbus_generate_guid.
func DbusGenerateGuid() {
	if dbusGenerateGuidInvoker == nil {
		dbusGenerateGuidInvoker = gi.FunctionInvokerNew("Gio", "dbus_generate_guid")
	}

	dbusGenerateGuidInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_dbus_gvalue_to_gvariant' : non trivial function

// UNSUPPORTED : C value 'g_dbus_gvariant_to_gvalue' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_address' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_guid' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_interface_name' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_member_name' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_name' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_supported_address' : non trivial function

// UNSUPPORTED : C value 'g_dbus_is_unique_name' : non trivial function

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : non trivial function

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : non trivial function

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : non trivial function

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : non trivial function

// UNSUPPORTED : C value 'g_file_new_for_path' : non trivial function

// UNSUPPORTED : C value 'g_file_new_for_uri' : non trivial function

// UNSUPPORTED : C value 'g_file_new_tmp' : non trivial function

// UNSUPPORTED : C value 'g_file_parse_name' : non trivial function

// UNSUPPORTED : C value 'g_icon_deserialize' : non trivial function

// UNSUPPORTED : C value 'g_icon_hash' : non trivial function

// UNSUPPORTED : C value 'g_icon_new_for_string' : non trivial function

// UNSUPPORTED : C value 'g_initable_newv' : non trivial function

// UNSUPPORTED : C value 'g_io_error_from_errno' : non trivial function

var ioErrorQuarkInvoker *gi.Function

// IoErrorQuark is a representation of the C type g_io_error_quark.
func IoErrorQuark() {
	if ioErrorQuarkInvoker == nil {
		ioErrorQuarkInvoker = gi.FunctionInvokerNew("Gio", "io_error_quark")
	}

	ioErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_io_extension_point_implement' : non trivial function

// UNSUPPORTED : C value 'g_io_extension_point_lookup' : non trivial function

// UNSUPPORTED : C value 'g_io_extension_point_register' : non trivial function

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : non trivial function

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : non trivial function

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory' : non trivial function

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory_with_scope' : non trivial function

var ioSchedulerCancelAllJobsInvoker *gi.Function

// IoSchedulerCancelAllJobs is a representation of the C type g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	if ioSchedulerCancelAllJobsInvoker == nil {
		ioSchedulerCancelAllJobsInvoker = gi.FunctionInvokerNew("Gio", "io_scheduler_cancel_all_jobs")
	}

	ioSchedulerCancelAllJobsInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : non trivial function

// UNSUPPORTED : C value 'g_keyfile_settings_backend_new' : non trivial function

var memorySettingsBackendNewInvoker *gi.Function

// MemorySettingsBackendNew is a representation of the C type g_memory_settings_backend_new.
func MemorySettingsBackendNew() {
	if memorySettingsBackendNewInvoker == nil {
		memorySettingsBackendNewInvoker = gi.FunctionInvokerNew("Gio", "memory_settings_backend_new")
	}

	memorySettingsBackendNewInvoker.Invoke()
}

var networkMonitorGetDefaultInvoker *gi.Function

// NetworkMonitorGetDefault is a representation of the C type g_network_monitor_get_default.
func NetworkMonitorGetDefault() {
	if networkMonitorGetDefaultInvoker == nil {
		networkMonitorGetDefaultInvoker = gi.FunctionInvokerNew("Gio", "network_monitor_get_default")
	}

	networkMonitorGetDefaultInvoker.Invoke()
}

var networkingInitInvoker *gi.Function

// NetworkingInit is a representation of the C type g_networking_init.
func NetworkingInit() {
	if networkingInitInvoker == nil {
		networkingInitInvoker = gi.FunctionInvokerNew("Gio", "networking_init")
	}

	networkingInitInvoker.Invoke()
}

var nullSettingsBackendNewInvoker *gi.Function

// NullSettingsBackendNew is a representation of the C type g_null_settings_backend_new.
func NullSettingsBackendNew() {
	if nullSettingsBackendNewInvoker == nil {
		nullSettingsBackendNewInvoker = gi.FunctionInvokerNew("Gio", "null_settings_backend_new")
	}

	nullSettingsBackendNewInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_pollable_source_new' : non trivial function

// UNSUPPORTED : C value 'g_pollable_source_new_full' : non trivial function

// UNSUPPORTED : C value 'g_pollable_stream_read' : non trivial function

// UNSUPPORTED : C value 'g_pollable_stream_write' : non trivial function

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : non trivial function

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : non trivial function

var proxyResolverGetDefaultInvoker *gi.Function

// ProxyResolverGetDefault is a representation of the C type g_proxy_resolver_get_default.
func ProxyResolverGetDefault() {
	if proxyResolverGetDefaultInvoker == nil {
		proxyResolverGetDefaultInvoker = gi.FunctionInvokerNew("Gio", "proxy_resolver_get_default")
	}

	proxyResolverGetDefaultInvoker.Invoke()
}

var resolverErrorQuarkInvoker *gi.Function

// ResolverErrorQuark is a representation of the C type g_resolver_error_quark.
func ResolverErrorQuark() {
	if resolverErrorQuarkInvoker == nil {
		resolverErrorQuarkInvoker = gi.FunctionInvokerNew("Gio", "resolver_error_quark")
	}

	resolverErrorQuarkInvoker.Invoke()
}

var resourceErrorQuarkInvoker *gi.Function

// ResourceErrorQuark is a representation of the C type g_resource_error_quark.
func ResourceErrorQuark() {
	if resourceErrorQuarkInvoker == nil {
		resourceErrorQuarkInvoker = gi.FunctionInvokerNew("Gio", "resource_error_quark")
	}

	resourceErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_resource_load' : non trivial function

// UNSUPPORTED : C value 'g_resources_enumerate_children' : non trivial function

// UNSUPPORTED : C value 'g_resources_get_info' : non trivial function

// UNSUPPORTED : C value 'g_resources_lookup_data' : non trivial function

// UNSUPPORTED : C value 'g_resources_open_stream' : non trivial function

// UNSUPPORTED : C value 'g_resources_register' : non trivial function

// UNSUPPORTED : C value 'g_resources_unregister' : non trivial function

var settingsSchemaSourceGetDefaultInvoker *gi.Function

// SettingsSchemaSourceGetDefault is a representation of the C type g_settings_schema_source_get_default.
func SettingsSchemaSourceGetDefault() {
	if settingsSchemaSourceGetDefaultInvoker == nil {
		settingsSchemaSourceGetDefaultInvoker = gi.FunctionInvokerNew("Gio", "settings_schema_source_get_default")
	}

	settingsSchemaSourceGetDefaultInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_simple_async_report_error_in_idle' : non trivial function

// UNSUPPORTED : C value 'g_simple_async_report_gerror_in_idle' : non trivial function

// UNSUPPORTED : C value 'g_simple_async_report_take_gerror_in_idle' : non trivial function

// UNSUPPORTED : C value 'g_srv_target_list_sort' : non trivial function

var tlsBackendGetDefaultInvoker *gi.Function

// TlsBackendGetDefault is a representation of the C type g_tls_backend_get_default.
func TlsBackendGetDefault() {
	if tlsBackendGetDefaultInvoker == nil {
		tlsBackendGetDefaultInvoker = gi.FunctionInvokerNew("Gio", "tls_backend_get_default")
	}

	tlsBackendGetDefaultInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_tls_client_connection_new' : non trivial function

var tlsErrorQuarkInvoker *gi.Function

// TlsErrorQuark is a representation of the C type g_tls_error_quark.
func TlsErrorQuark() {
	if tlsErrorQuarkInvoker == nil {
		tlsErrorQuarkInvoker = gi.FunctionInvokerNew("Gio", "tls_error_quark")
	}

	tlsErrorQuarkInvoker.Invoke()
}

// UNSUPPORTED : C value 'g_tls_file_database_new' : non trivial function

// UNSUPPORTED : C value 'g_tls_server_connection_new' : non trivial function

// UNSUPPORTED : C value 'g_unix_is_mount_path_system_internal' : non trivial function

// UNSUPPORTED : C value 'g_unix_is_system_device_path' : non trivial function

// UNSUPPORTED : C value 'g_unix_is_system_fs_type' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_at' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_compare' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_copy' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_for' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_free' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_get_device_path' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_get_fs_type' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_get_mount_path' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_guess_can_eject' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_guess_icon' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_guess_name' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_guess_should_display' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_guess_symbolic_icon' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_is_readonly' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_is_system_internal' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_points_changed_since' : non trivial function

// UNSUPPORTED : C value 'g_unix_mount_points_get' : non trivial function

// UNSUPPORTED : C value 'g_unix_mounts_changed_since' : non trivial function

// UNSUPPORTED : C value 'g_unix_mounts_get' : non trivial function
