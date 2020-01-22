// Code generated - DO NOT EDIT.
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

// FILE_ATTRIBUTE_TRASH_DELETION_DATE is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_DELETION_DATE.
//
// since 2.24
const FILE_ATTRIBUTE_TRASH_DELETION_DATE = "trash::deletion-date"

// FILE_ATTRIBUTE_TRASH_ORIG_PATH is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_ORIG_PATH.
//
// since 2.24
const FILE_ATTRIBUTE_TRASH_ORIG_PATH = "trash::orig-path"

// ConverterFlags is a representation of the C bitfield GConverterFlags.
type ConverterFlags int

// ConverterFlags_none is a representation of the C bitfield member G_CONVERTER_NO_FLAGS.
const ConverterFlags_none = ConverterFlags(0)

// ConverterFlags_input_at_end is a representation of the C bitfield member G_CONVERTER_INPUT_AT_END.
const ConverterFlags_input_at_end = ConverterFlags(1)

// ConverterFlags_flush is a representation of the C bitfield member G_CONVERTER_FLUSH.
const ConverterFlags_flush = ConverterFlags(2)

// ConverterResult is a representation of the C enumeration GConverterResult.
type ConverterResult int

// ConverterResult_error is a representation of the C enumeration member G_CONVERTER_ERROR.
const ConverterResult_error = ConverterResult(0)

// ConverterResult_converted is a representation of the C enumeration member G_CONVERTER_CONVERTED.
const ConverterResult_converted = ConverterResult(1)

// ConverterResult_finished is a representation of the C enumeration member G_CONVERTER_FINISHED.
const ConverterResult_finished = ConverterResult(2)

// ConverterResult_flushed is a representation of the C enumeration member G_CONVERTER_FLUSHED.
const ConverterResult_flushed = ConverterResult(3)

// ZlibCompressorFormat is a representation of the C enumeration GZlibCompressorFormat.
type ZlibCompressorFormat int

// ZlibCompressorFormat_zlib is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_ZLIB.
const ZlibCompressorFormat_zlib = ZlibCompressorFormat(0)

// ZlibCompressorFormat_gzip is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_GZIP.
const ZlibCompressorFormat_gzip = ZlibCompressorFormat(1)

// ZlibCompressorFormat_raw is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_RAW.
const ZlibCompressorFormat_raw = ZlibCompressorFormat(2)

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

// ConverterIface is a representation of the C record GConverterIface.
//
// since 2.24
type ConverterIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterIface that represents the ConverterIface.
func (recv *ConverterIface) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterIfaceNewFromC creates a new ConverterIface from a pointer to the C GConverterIface that represents the ConverterIface.
func ConverterIfaceNewFromC(native unsafe.Pointer) *ConverterIface {
	return &ConverterIface{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// Converter is a representation of the C interface GConverter.
//
// since 2.24
type Converter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverter that represents the Converter.
func (recv *Converter) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterNewFromC creates a new Converter from a pointer to the C GConverter that represents the Converter.
func ConverterNewFromC(native unsafe.Pointer) *Converter {
	return &Converter{native: native}
}

// FileDescriptorBased is a representation of the C interface GFileDescriptorBased.
//
// since 2.24
type FileDescriptorBased struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileDescriptorBased that represents the FileDescriptorBased.
func (recv *FileDescriptorBased) ToC() unsafe.Pointer {
	return recv.native
}

// FileDescriptorBasedNewFromC creates a new FileDescriptorBased from a pointer to the C GFileDescriptorBased that represents the FileDescriptorBased.
func FileDescriptorBasedNewFromC(native unsafe.Pointer) *FileDescriptorBased {
	return &FileDescriptorBased{native: native}
}
