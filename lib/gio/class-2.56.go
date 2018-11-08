// This is a generated file - DO NOT EDIT
// +build gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// SetOptionContextDescription is a wrapper around the C function g_application_set_option_context_description.
func (recv *Application) SetOptionContextDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_application_set_option_context_description((*C.GApplication)(recv.native), c_description)

	return
}

// SetOptionContextParameterString is a wrapper around the C function g_application_set_option_context_parameter_string.
func (recv *Application) SetOptionContextParameterString(parameterString string) {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	C.g_application_set_option_context_parameter_string((*C.GApplication)(recv.native), c_parameter_string)

	return
}

// SetOptionContextSummary is a wrapper around the C function g_application_set_option_context_summary.
func (recv *Application) SetOptionContextSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_application_set_option_context_summary((*C.GApplication)(recv.native), c_summary)

	return
}

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter (GConverter*) for param converter

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter (GConverter*) for param converter

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no type generator for guint8 () for array param blob

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// GetLocaleString is a wrapper around the C function g_desktop_app_info_get_locale_string.
func (recv *DesktopAppInfo) GetLocaleString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_locale_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon (GIcon*) for param icon

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File (GFile*) for param file

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no type generator for guint8 () for array param bytes

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc (GReallocFunc) for param realloc_function

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// JoinMulticastGroupSsm is a wrapper around the C function g_socket_join_multicast_group_ssm.
func (recv *Socket) JoinMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific := (*C.GInetAddress)(sourceSpecific.ToC())

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LeaveMulticastGroupSsm is a wrapper around the C function g_socket_leave_multicast_group_ssm.
func (recv *Socket) LeaveMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(group.ToC())

	c_source_specific := (*C.GInetAddress)(sourceSpecific.ToC())

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no type generator for filename () for array param argv

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no type generator for utf8 (char*) for array param iconnames

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no type generator for gint (gint) for array param fds

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no type generator for gchar () for array param path

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no type generator for gchar () for array param path
