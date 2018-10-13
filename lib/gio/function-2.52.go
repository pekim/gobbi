// This is a generated file - DO NOT EDIT
// +build gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_action_parse_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_action_print_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_app_info_create_from_commandline : no return generator

// Unsupported : g_app_info_get_default_for_type : no return generator

// Unsupported : g_app_info_get_default_for_uri_scheme : no return generator

// Unsupported : g_app_info_launch_default_for_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_app_info_launch_default_for_uri_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_async_initable_newv_async : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_bus_get_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback, GBusAcquiredCallback

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback, GBusNameAcquiredCallback

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback, GBusNameAppearedCallback

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback, GBusNameAppearedCallback

// Unsupported : g_content_type_get_icon : no return generator

// Unsupported : g_content_type_get_symbolic_icon : no return generator

// Unsupported : g_content_type_guess : unsupported parameter data : no param type

// Unsupported : g_content_type_guess_for_tree : unsupported parameter root : no type generator for File, GFile*

// ContentTypeIsMimeType is a wrapper around the C function g_content_type_is_mime_type.
func ContentTypeIsMimeType(type_ string, mimeType string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.g_content_type_is_mime_type(c_type, c_mime_type)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_address_get_stream_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations : no param type

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries : no param type

// Unsupported : g_dbus_gvalue_to_gvariant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_dbus_gvariant_to_gvalue : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_dtls_client_connection_new : unsupported parameter base_socket : no type generator for DatagramBased, GDatagramBased*

// Unsupported : g_dtls_server_connection_new : unsupported parameter base_socket : no type generator for DatagramBased, GDatagramBased*

// Unsupported : g_file_new_for_commandline_arg : no return generator

// Unsupported : g_file_new_for_commandline_arg_and_cwd : no return generator

// Unsupported : g_file_new_for_path : no return generator

// Unsupported : g_file_new_for_uri : no return generator

// Unsupported : g_file_new_tmp : unsupported parameter iostream : record with indirection level of 2

// Unsupported : g_file_parse_name : no return generator

// Unsupported : g_icon_deserialize : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_icon_new_for_string : no return generator

// Unsupported : g_initable_newv : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_io_extension_point_implement : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc, GIOSchedulerJobFunc

// Unsupported : g_network_monitor_get_default : no return generator

// Unsupported : g_pollable_stream_read : unsupported parameter buffer : no param type

// Unsupported : g_pollable_stream_write : unsupported parameter buffer : no param type

// Unsupported : g_pollable_stream_write_all : unsupported parameter buffer : no param type

// Unsupported : g_proxy_get_default_for_protocol : no return generator

// Unsupported : g_proxy_resolver_get_default : no return generator

// Unsupported : g_resources_enumerate_children : no return type

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_backend_get_default : no return generator

// Unsupported : g_tls_client_connection_new : unsupported parameter server_identity : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_tls_file_database_new : no return generator

// Unsupported : g_tls_server_connection_new : no return generator

// UnixMountFor is a wrapper around the C function g_unix_mount_for.
func UnixMountFor(filePath string) (*UnixMountEntry, uint64) {
	c_file_path := C.CString(filePath)
	defer C.free(unsafe.Pointer(c_file_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_for(c_file_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// Unsupported : g_unix_mount_guess_icon : no return generator

// Unsupported : g_unix_mount_guess_symbolic_icon : no return generator
