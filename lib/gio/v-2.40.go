// Code generated - DO NOT EDIT.
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

// FILE_ATTRIBUTE_THUMBNAIL_IS_VALID is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID.
//
// since 2.40
const FILE_ATTRIBUTE_THUMBNAIL_IS_VALID = "thumbnail::is-valid"

// SubprocessFlags is a representation of the C bitfield GSubprocessFlags.
type SubprocessFlags int

// SubprocessFlags_none is a representation of the C bitfield member G_SUBPROCESS_FLAGS_NONE.
const SubprocessFlags_none = SubprocessFlags(0)

// SubprocessFlags_stdin_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDIN_PIPE.
const SubprocessFlags_stdin_pipe = SubprocessFlags(1)

// SubprocessFlags_stdin_inherit is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDIN_INHERIT.
const SubprocessFlags_stdin_inherit = SubprocessFlags(2)

// SubprocessFlags_stdout_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDOUT_PIPE.
const SubprocessFlags_stdout_pipe = SubprocessFlags(4)

// SubprocessFlags_stdout_silence is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDOUT_SILENCE.
const SubprocessFlags_stdout_silence = SubprocessFlags(8)

// SubprocessFlags_stderr_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_PIPE.
const SubprocessFlags_stderr_pipe = SubprocessFlags(16)

// SubprocessFlags_stderr_silence is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_SILENCE.
const SubprocessFlags_stderr_silence = SubprocessFlags(32)

// SubprocessFlags_stderr_merge is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_MERGE.
const SubprocessFlags_stderr_merge = SubprocessFlags(64)

// SubprocessFlags_inherit_fds is a representation of the C bitfield member G_SUBPROCESS_FLAGS_INHERIT_FDS.
const SubprocessFlags_inherit_fds = SubprocessFlags(128)

// TlsCertificateRequestFlags is a representation of the C enumeration GTlsCertificateRequestFlags.
type TlsCertificateRequestFlags int

// TlsCertificateRequestFlags_none is a representation of the C enumeration member G_TLS_CERTIFICATE_REQUEST_NONE.
const TlsCertificateRequestFlags_none = TlsCertificateRequestFlags(0)

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

// AppInfoMonitor is a representation of the C record GAppInfoMonitor.
//
// since 2.40
type AppInfoMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfoMonitor that represents the AppInfoMonitor.
func (recv *AppInfoMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoMonitorNewFromC creates a new AppInfoMonitor from a pointer to the C GAppInfoMonitor that represents the AppInfoMonitor.
func AppInfoMonitorNewFromC(native unsafe.Pointer) *AppInfoMonitor {
	return &AppInfoMonitor{native: native}
}

// Notification is a representation of the C record GNotification.
//
// since 2.40
type Notification struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNotification that represents the Notification.
func (recv *Notification) ToC() unsafe.Pointer {
	return recv.native
}

// NotificationNewFromC creates a new Notification from a pointer to the C GNotification that represents the Notification.
func NotificationNewFromC(native unsafe.Pointer) *Notification {
	return &Notification{native: native}
}

// Subprocess is a representation of the C record GSubprocess.
//
// since 2.40
type Subprocess struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSubprocess that represents the Subprocess.
func (recv *Subprocess) ToC() unsafe.Pointer {
	return recv.native
}

// SubprocessNewFromC creates a new Subprocess from a pointer to the C GSubprocess that represents the Subprocess.
func SubprocessNewFromC(native unsafe.Pointer) *Subprocess {
	return &Subprocess{native: native}
}

// SubprocessLauncher is a representation of the C record GSubprocessLauncher.
//
// since 2.40
type SubprocessLauncher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSubprocessLauncher that represents the SubprocessLauncher.
func (recv *SubprocessLauncher) ToC() unsafe.Pointer {
	return recv.native
}

// SubprocessLauncherNewFromC creates a new SubprocessLauncher from a pointer to the C GSubprocessLauncher that represents the SubprocessLauncher.
func SubprocessLauncherNewFromC(native unsafe.Pointer) *SubprocessLauncher {
	return &SubprocessLauncher{native: native}
}
