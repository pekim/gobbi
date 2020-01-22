// Code generated - DO NOT EDIT.
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

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

// DatagramBasedInterface is a representation of the C record GDatagramBasedInterface.
//
// since 2.48
type DatagramBasedInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDatagramBasedInterface that represents the DatagramBasedInterface.
func (recv *DatagramBasedInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DatagramBasedInterfaceNewFromC creates a new DatagramBasedInterface from a pointer to the C GDatagramBasedInterface that represents the DatagramBasedInterface.
func DatagramBasedInterfaceNewFromC(native unsafe.Pointer) *DatagramBasedInterface {
	return &DatagramBasedInterface{native: native}
}

// DtlsClientConnectionInterface is a representation of the C record GDtlsClientConnectionInterface.
//
// since 2.48
type DtlsClientConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsClientConnectionInterface that represents the DtlsClientConnectionInterface.
func (recv *DtlsClientConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsClientConnectionInterfaceNewFromC creates a new DtlsClientConnectionInterface from a pointer to the C GDtlsClientConnectionInterface that represents the DtlsClientConnectionInterface.
func DtlsClientConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsClientConnectionInterface {
	return &DtlsClientConnectionInterface{native: native}
}

// DtlsConnectionInterface is a representation of the C record GDtlsConnectionInterface.
//
// since 2.48
type DtlsConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsConnectionInterface that represents the DtlsConnectionInterface.
func (recv *DtlsConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsConnectionInterfaceNewFromC creates a new DtlsConnectionInterface from a pointer to the C GDtlsConnectionInterface that represents the DtlsConnectionInterface.
func DtlsConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsConnectionInterface {
	return &DtlsConnectionInterface{native: native}
}

// DtlsServerConnectionInterface is a representation of the C record GDtlsServerConnectionInterface.
//
// since 2.48
type DtlsServerConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsServerConnectionInterface that represents the DtlsServerConnectionInterface.
func (recv *DtlsServerConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsServerConnectionInterfaceNewFromC creates a new DtlsServerConnectionInterface from a pointer to the C GDtlsServerConnectionInterface that represents the DtlsServerConnectionInterface.
func DtlsServerConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsServerConnectionInterface {
	return &DtlsServerConnectionInterface{native: native}
}

// InputMessage is a representation of the C record GInputMessage.
//
// since 2.48
type InputMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputMessage that represents the InputMessage.
func (recv *InputMessage) ToC() unsafe.Pointer {
	return recv.native
}

// InputMessageNewFromC creates a new InputMessage from a pointer to the C GInputMessage that represents the InputMessage.
func InputMessageNewFromC(native unsafe.Pointer) *InputMessage {
	return &InputMessage{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// DatagramBased is a representation of the C interface GDatagramBased.
//
// since 2.48
type DatagramBased struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDatagramBased that represents the DatagramBased.
func (recv *DatagramBased) ToC() unsafe.Pointer {
	return recv.native
}

// DatagramBasedNewFromC creates a new DatagramBased from a pointer to the C GDatagramBased that represents the DatagramBased.
func DatagramBasedNewFromC(native unsafe.Pointer) *DatagramBased {
	return &DatagramBased{native: native}
}

// DtlsClientConnection is a representation of the C interface GDtlsClientConnection.
//
// since 2.48
type DtlsClientConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsClientConnection that represents the DtlsClientConnection.
func (recv *DtlsClientConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsClientConnectionNewFromC creates a new DtlsClientConnection from a pointer to the C GDtlsClientConnection that represents the DtlsClientConnection.
func DtlsClientConnectionNewFromC(native unsafe.Pointer) *DtlsClientConnection {
	return &DtlsClientConnection{native: native}
}

// DtlsConnection is a representation of the C interface GDtlsConnection.
//
// since 2.48
type DtlsConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsConnection that represents the DtlsConnection.
func (recv *DtlsConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsConnectionNewFromC creates a new DtlsConnection from a pointer to the C GDtlsConnection that represents the DtlsConnection.
func DtlsConnectionNewFromC(native unsafe.Pointer) *DtlsConnection {
	return &DtlsConnection{native: native}
}

// DtlsServerConnection is a representation of the C interface GDtlsServerConnection.
//
// since 2.48
type DtlsServerConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsServerConnection that represents the DtlsServerConnection.
func (recv *DtlsServerConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsServerConnectionNewFromC creates a new DtlsServerConnection from a pointer to the C GDtlsServerConnection that represents the DtlsServerConnection.
func DtlsServerConnectionNewFromC(native unsafe.Pointer) *DtlsServerConnection {
	return &DtlsServerConnection{native: native}
}
