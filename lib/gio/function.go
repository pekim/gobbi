// Code generated - DO NOT EDIT.

package gio

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'g_action_name_is_valid' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_action_parse_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_action_print_detailed_name' : parameter 'target_value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_app_info_create_from_commandline' : parameter 'commandline' of type 'filename' not supported

// UNSUPPORTED : C value 'g_app_info_get_all' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_all_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_type' : parameter 'must_support_uris' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_app_info_get_default_for_uri_scheme' : return type 'AppInfo' not supported

// UNSUPPORTED : C value 'g_app_info_get_fallback_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_get_recommended_for_type' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri' : parameter 'context' of type 'AppLaunchContext' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_async' : parameter 'context' of type 'AppLaunchContext' not supported

// UNSUPPORTED : C value 'g_app_info_launch_default_for_uri_finish' : parameter 'result' of type 'AsyncResult' not supported

var appInfoResetTypeAssociationsFunction *gi.Function
var appInfoResetTypeAssociationsFunction_Once sync.Once

func appInfoResetTypeAssociationsFunction_Set() {
	appInfoResetTypeAssociationsFunction_Once.Do(func() {
		appInfoResetTypeAssociationsFunction = gi.FunctionInvokerNew("Gio", "app_info_reset_type_associations")
	})
}

var appInfoResetTypeAssociationsInvoker *gi.Function

// AppInfoResetTypeAssociations is a representation of the C type g_app_info_reset_type_associations.
func AppInfoResetTypeAssociations(contentType string) {
	appInfoResetTypeAssociationsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(contentType)

	appInfoResetTypeAssociationsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_async_initable_newv_async' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_bus_get' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_get_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_bus_get_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_own_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_own_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

var busUnownNameFunction *gi.Function
var busUnownNameFunction_Once sync.Once

func busUnownNameFunction_Set() {
	busUnownNameFunction_Once.Do(func() {
		busUnownNameFunction = gi.FunctionInvokerNew("Gio", "bus_unown_name")
	})
}

var busUnownNameInvoker *gi.Function

// BusUnownName is a representation of the C type g_bus_unown_name.
func BusUnownName(ownerId uint32) {
	busUnownNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(ownerId)

	busUnownNameFunction.Invoke(inArgs[:], nil)

}

var busUnwatchNameFunction *gi.Function
var busUnwatchNameFunction_Once sync.Once

func busUnwatchNameFunction_Set() {
	busUnwatchNameFunction_Once.Do(func() {
		busUnwatchNameFunction = gi.FunctionInvokerNew("Gio", "bus_unwatch_name")
	})
}

var busUnwatchNameInvoker *gi.Function

// BusUnwatchName is a representation of the C type g_bus_unwatch_name.
func BusUnwatchName(watcherId uint32) {
	busUnwatchNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(watcherId)

	busUnwatchNameFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bus_watch_name' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_on_connection_with_closures' : parameter 'connection' of type 'DBusConnection' not supported

// UNSUPPORTED : C value 'g_bus_watch_name_with_closures' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_content_type_can_be_executable' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_content_type_equals' : return type 'gboolean' not supported

var contentTypeFromMimeTypeFunction *gi.Function
var contentTypeFromMimeTypeFunction_Once sync.Once

func contentTypeFromMimeTypeFunction_Set() {
	contentTypeFromMimeTypeFunction_Once.Do(func() {
		contentTypeFromMimeTypeFunction = gi.FunctionInvokerNew("Gio", "content_type_from_mime_type")
	})
}

var contentTypeFromMimeTypeInvoker *gi.Function

// ContentTypeFromMimeType is a representation of the C type g_content_type_from_mime_type.
func ContentTypeFromMimeType(mimeType string) string {
	contentTypeFromMimeTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	ret := contentTypeFromMimeTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var contentTypeGetDescriptionFunction *gi.Function
var contentTypeGetDescriptionFunction_Once sync.Once

func contentTypeGetDescriptionFunction_Set() {
	contentTypeGetDescriptionFunction_Once.Do(func() {
		contentTypeGetDescriptionFunction = gi.FunctionInvokerNew("Gio", "content_type_get_description")
	})
}

var contentTypeGetDescriptionInvoker *gi.Function

// ContentTypeGetDescription is a representation of the C type g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	contentTypeGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	ret := contentTypeGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var contentTypeGetGenericIconNameFunction *gi.Function
var contentTypeGetGenericIconNameFunction_Once sync.Once

func contentTypeGetGenericIconNameFunction_Set() {
	contentTypeGetGenericIconNameFunction_Once.Do(func() {
		contentTypeGetGenericIconNameFunction = gi.FunctionInvokerNew("Gio", "content_type_get_generic_icon_name")
	})
}

var contentTypeGetGenericIconNameInvoker *gi.Function

// ContentTypeGetGenericIconName is a representation of the C type g_content_type_get_generic_icon_name.
func ContentTypeGetGenericIconName(type_ string) string {
	contentTypeGetGenericIconNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	ret := contentTypeGetGenericIconNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_get_icon' : return type 'Icon' not supported

var contentTypeGetMimeDirsFunction *gi.Function
var contentTypeGetMimeDirsFunction_Once sync.Once

func contentTypeGetMimeDirsFunction_Set() {
	contentTypeGetMimeDirsFunction_Once.Do(func() {
		contentTypeGetMimeDirsFunction = gi.FunctionInvokerNew("Gio", "content_type_get_mime_dirs")
	})
}

var contentTypeGetMimeDirsInvoker *gi.Function

// ContentTypeGetMimeDirs is a representation of the C type g_content_type_get_mime_dirs.
func ContentTypeGetMimeDirs() {
	contentTypeGetMimeDirsFunction_Set()

	contentTypeGetMimeDirsFunction.Invoke(nil, nil)

}

var contentTypeGetMimeTypeFunction *gi.Function
var contentTypeGetMimeTypeFunction_Once sync.Once

func contentTypeGetMimeTypeFunction_Set() {
	contentTypeGetMimeTypeFunction_Once.Do(func() {
		contentTypeGetMimeTypeFunction = gi.FunctionInvokerNew("Gio", "content_type_get_mime_type")
	})
}

var contentTypeGetMimeTypeInvoker *gi.Function

// ContentTypeGetMimeType is a representation of the C type g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	contentTypeGetMimeTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(type_)

	ret := contentTypeGetMimeTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_content_type_get_symbolic_icon' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_content_type_guess' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_content_type_guess_for_tree' : parameter 'root' of type 'File' not supported

// UNSUPPORTED : C value 'g_content_type_is_a' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_content_type_is_mime_type' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_content_type_is_unknown' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_content_type_set_mime_dirs' : parameter 'dirs' has no type

// UNSUPPORTED : C value 'g_content_types_get_registered' : return type 'GLib.List' not supported

var dbusAddressEscapeValueFunction *gi.Function
var dbusAddressEscapeValueFunction_Once sync.Once

func dbusAddressEscapeValueFunction_Set() {
	dbusAddressEscapeValueFunction_Once.Do(func() {
		dbusAddressEscapeValueFunction = gi.FunctionInvokerNew("Gio", "dbus_address_escape_value")
	})
}

var dbusAddressEscapeValueInvoker *gi.Function

// DbusAddressEscapeValue is a representation of the C type g_dbus_address_escape_value.
func DbusAddressEscapeValue(string_ string) string {
	dbusAddressEscapeValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := dbusAddressEscapeValueFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_address_get_for_bus_sync' : parameter 'bus_type' of type 'BusType' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream' : parameter 'cancellable' of type 'Cancellable' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_finish' : parameter 'res' of type 'AsyncResult' not supported

// UNSUPPORTED : C value 'g_dbus_address_get_stream_sync' : parameter 'cancellable' of type 'Cancellable' not supported

// UNSUPPORTED : C value 'g_dbus_annotation_info_lookup' : parameter 'annotations' has no type

// UNSUPPORTED : C value 'g_dbus_error_encode_gerror' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_get_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_is_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_new_for_dbus_error' : return type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_dbus_error_register_error_domain' : parameter 'quark_volatile' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_dbus_error_strip_remote_error' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'g_dbus_error_unregister_error' : parameter 'error_domain' of type 'GLib.Quark' not supported

var dbusGenerateGuidFunction *gi.Function
var dbusGenerateGuidFunction_Once sync.Once

func dbusGenerateGuidFunction_Set() {
	dbusGenerateGuidFunction_Once.Do(func() {
		dbusGenerateGuidFunction = gi.FunctionInvokerNew("Gio", "dbus_generate_guid")
	})
}

var dbusGenerateGuidInvoker *gi.Function

// DbusGenerateGuid is a representation of the C type g_dbus_generate_guid.
func DbusGenerateGuid() string {
	dbusGenerateGuidFunction_Set()

	ret := dbusGenerateGuidFunction.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_dbus_gvalue_to_gvariant' : parameter 'gvalue' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'g_dbus_gvariant_to_gvalue' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_dbus_is_address' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_guid' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_interface_name' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_member_name' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_name' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_supported_address' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dbus_is_unique_name' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_dtls_client_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_dtls_server_connection_new' : parameter 'base_socket' of type 'DatagramBased' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_commandline_arg_and_cwd' : parameter 'arg' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_path' : parameter 'path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_new_for_uri' : return type 'File' not supported

// UNSUPPORTED : C value 'g_file_new_tmp' : parameter 'tmpl' of type 'filename' not supported

// UNSUPPORTED : C value 'g_file_parse_name' : return type 'File' not supported

// UNSUPPORTED : C value 'g_icon_deserialize' : parameter 'value' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'g_icon_hash' : parameter 'icon' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_icon_new_for_string' : return type 'Icon' not supported

// UNSUPPORTED : C value 'g_initable_newv' : parameter 'object_type' of type 'GType' not supported

// UNSUPPORTED : C value 'g_io_error_from_errno' : return type 'IOErrorEnum' not supported

// UNSUPPORTED : C value 'g_io_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_io_extension_point_implement' : parameter 'type' of type 'GType' not supported

var ioExtensionPointLookupFunction *gi.Function
var ioExtensionPointLookupFunction_Once sync.Once

func ioExtensionPointLookupFunction_Set() {
	ioExtensionPointLookupFunction_Once.Do(func() {
		ioExtensionPointLookupFunction = gi.FunctionInvokerNew("Gio", "io_extension_point_lookup")
	})
}

var ioExtensionPointLookupInvoker *gi.Function

// IoExtensionPointLookup is a representation of the C type g_io_extension_point_lookup.
func IoExtensionPointLookup(name string) *IOExtensionPoint {
	ioExtensionPointLookupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := ioExtensionPointLookupFunction.Invoke(inArgs[:], nil)

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo
}

var ioExtensionPointRegisterFunction *gi.Function
var ioExtensionPointRegisterFunction_Once sync.Once

func ioExtensionPointRegisterFunction_Set() {
	ioExtensionPointRegisterFunction_Once.Do(func() {
		ioExtensionPointRegisterFunction = gi.FunctionInvokerNew("Gio", "io_extension_point_register")
	})
}

var ioExtensionPointRegisterInvoker *gi.Function

// IoExtensionPointRegister is a representation of the C type g_io_extension_point_register.
func IoExtensionPointRegister(name string) *IOExtensionPoint {
	ioExtensionPointRegisterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := ioExtensionPointRegisterFunction.Invoke(inArgs[:], nil)

	retGo := &IOExtensionPoint{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_load_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory' : parameter 'dirname' of type 'filename' not supported

// UNSUPPORTED : C value 'g_io_modules_scan_all_in_directory_with_scope' : parameter 'dirname' of type 'filename' not supported

var ioSchedulerCancelAllJobsFunction *gi.Function
var ioSchedulerCancelAllJobsFunction_Once sync.Once

func ioSchedulerCancelAllJobsFunction_Set() {
	ioSchedulerCancelAllJobsFunction_Once.Do(func() {
		ioSchedulerCancelAllJobsFunction = gi.FunctionInvokerNew("Gio", "io_scheduler_cancel_all_jobs")
	})
}

var ioSchedulerCancelAllJobsInvoker *gi.Function

// IoSchedulerCancelAllJobs is a representation of the C type g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	ioSchedulerCancelAllJobsFunction_Set()

	ioSchedulerCancelAllJobsFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_io_scheduler_push_job' : parameter 'job_func' of type 'IOSchedulerJobFunc' not supported

// UNSUPPORTED : C value 'g_keyfile_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_memory_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_network_monitor_get_default' : return type 'NetworkMonitor' not supported

var networkingInitFunction *gi.Function
var networkingInitFunction_Once sync.Once

func networkingInitFunction_Set() {
	networkingInitFunction_Once.Do(func() {
		networkingInitFunction = gi.FunctionInvokerNew("Gio", "networking_init")
	})
}

var networkingInitInvoker *gi.Function

// NetworkingInit is a representation of the C type g_networking_init.
func NetworkingInit() {
	networkingInitFunction_Set()

	networkingInitFunction.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'g_null_settings_backend_new' : return type 'SettingsBackend' not supported

// UNSUPPORTED : C value 'g_pollable_source_new' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_source_new_full' : parameter 'pollable_stream' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'g_pollable_stream_read' : parameter 'stream' of type 'InputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_pollable_stream_write_all' : parameter 'stream' of type 'OutputStream' not supported

// UNSUPPORTED : C value 'g_proxy_get_default_for_protocol' : return type 'Proxy' not supported

// UNSUPPORTED : C value 'g_proxy_resolver_get_default' : return type 'ProxyResolver' not supported

// UNSUPPORTED : C value 'g_resolver_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'g_resource_load' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_resources_enumerate_children' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_get_info' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_lookup_data' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_open_stream' : parameter 'lookup_flags' of type 'ResourceLookupFlags' not supported

// UNSUPPORTED : C value 'g_resources_register' : parameter 'resource' of type 'Resource' not supported

// UNSUPPORTED : C value 'g_resources_unregister' : parameter 'resource' of type 'Resource' not supported

var settingsSchemaSourceGetDefaultFunction *gi.Function
var settingsSchemaSourceGetDefaultFunction_Once sync.Once

func settingsSchemaSourceGetDefaultFunction_Set() {
	settingsSchemaSourceGetDefaultFunction_Once.Do(func() {
		settingsSchemaSourceGetDefaultFunction = gi.FunctionInvokerNew("Gio", "settings_schema_source_get_default")
	})
}

var settingsSchemaSourceGetDefaultInvoker *gi.Function

// SettingsSchemaSourceGetDefault is a representation of the C type g_settings_schema_source_get_default.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	settingsSchemaSourceGetDefaultFunction_Set()

	ret := settingsSchemaSourceGetDefaultFunction.Invoke(nil, nil)

	retGo := &SettingsSchemaSource{native: ret.Pointer()}

	return retGo
}

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

// UNSUPPORTED : C value 'g_unix_is_system_device_path' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_is_system_fs_type' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_unix_mount_at' : parameter 'mount_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_compare' : parameter 'mount1' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_copy' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_for' : parameter 'file_path' of type 'filename' not supported

// UNSUPPORTED : C value 'g_unix_mount_free' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_device_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_fs_type' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_mount_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_options' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

// UNSUPPORTED : C value 'g_unix_mount_get_root_path' : parameter 'mount_entry' of type 'UnixMountEntry' not supported

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
