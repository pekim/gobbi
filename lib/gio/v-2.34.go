// Code generated - DO NOT EDIT.
// +build gio_2.34

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
//
// since 2.34
const FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON = "standard::symbolic-icon"

// TestDBusFlags is a representation of the C bitfield GTestDBusFlags.
type TestDBusFlags int

// TestDBusFlags_none is a representation of the C bitfield member G_TEST_DBUS_NONE.
const TestDBusFlags_none = TestDBusFlags(0)

// ResolverRecordType is a representation of the C enumeration GResolverRecordType.
type ResolverRecordType int

// ResolverRecordType_srv is a representation of the C enumeration member G_RESOLVER_RECORD_SRV.
const ResolverRecordType_srv = ResolverRecordType(1)

// ResolverRecordType_mx is a representation of the C enumeration member G_RESOLVER_RECORD_MX.
const ResolverRecordType_mx = ResolverRecordType(2)

// ResolverRecordType_txt is a representation of the C enumeration member G_RESOLVER_RECORD_TXT.
const ResolverRecordType_txt = ResolverRecordType(3)

// ResolverRecordType_soa is a representation of the C enumeration member G_RESOLVER_RECORD_SOA.
const ResolverRecordType_soa = ResolverRecordType(4)

// ResolverRecordType_ns is a representation of the C enumeration member G_RESOLVER_RECORD_NS.
const ResolverRecordType_ns = ResolverRecordType(5)

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

// ContentTypeGetGenericIconName wraps the C function g_content_type_get_generic_icon_name.
//
// since 2.34
func ContentTypeGetGenericIconName(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_generic_icon_name(sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// ContentTypeGetSymbolicIcon wraps the C function g_content_type_get_symbolic_icon.
//
// since 2.34
func ContentTypeGetSymbolicIcon(type_ string) *Icon {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_symbolic_icon(sys_type_)
	ret := IconNewFromC(retSys)

	return ret
}

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

// PollableSourceNewFull wraps the C function g_pollable_source_new_full.
//
// since 2.34
func PollableSourceNewFull(pollableStream unsafe.Pointer, childSource *glib.Source, cancellable *Cancellable) *glib.Source {
	sys_pollableStream := pollableStream
	sys_childSource := childSource.ToC()
	sys_cancellable := cancellable.ToC()
	retSys := gio.Fn_g_pollable_source_new_full(sys_pollableStream, sys_childSource, sys_cancellable)
	ret := glib.SourceNewFromC(retSys)

	return ret
}

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

// UnixMountGuessSymbolicIcon wraps the C function g_unix_mount_guess_symbolic_icon.
//
// since 2.34
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) *Icon {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_symbolic_icon(sys_mountEntry)
	ret := IconNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// TestDBus is a representation of the C record GTestDBus.
//
// since 2.34
type TestDBus struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTestDBus that represents the TestDBus.
func (recv *TestDBus) ToC() unsafe.Pointer {
	return recv.native
}

// TestDBusNewFromC creates a new TestDBus from a pointer to the C GTestDBus that represents the TestDBus.
func TestDBusNewFromC(native unsafe.Pointer) *TestDBus {
	return &TestDBus{native: native}
}
