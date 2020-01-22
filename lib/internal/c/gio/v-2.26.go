// Code generated - DO NOT EDIT.
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

static GDBusMessage* c_g_dbus_message_new_method_error(GDBusMessage* method_call_message, const gchar* error_name, const gchar* error_message_format) {
    return g_dbus_message_new_method_error(method_call_message, error_name, error_message_format, NULL);
}
*/
/*

static GDBusMessage* c_g_dbus_message_new_method_error_valist(GDBusMessage* method_call_message, const gchar* error_name, const gchar* error_message_format) {
    return g_dbus_message_new_method_error_valist(method_call_message, error_name, error_message_format, NULL);
}
*/
/*

static void c_g_dbus_method_invocation_return_error(GDBusMethodInvocation* invocation, GQuark domain, gint code, const gchar* format) {
    return g_dbus_method_invocation_return_error(invocation, domain, code, format, NULL);
}
*/
/*

static void c_g_dbus_method_invocation_return_error_valist(GDBusMethodInvocation* invocation, GQuark domain, gint code, const gchar* format) {
    return g_dbus_method_invocation_return_error_valist(invocation, domain, code, format, NULL);
}
*/
/*

static void c_g_settings_get(GSettings* settings, const gchar* key, const gchar* format) {
    return g_settings_get(settings, key, format, NULL);
}
*/
/*

static gboolean c_g_settings_set(GSettings* settings, const gchar* key, const gchar* format) {
    return g_settings_set(settings, key, format, NULL);
}
*/
import "C"

type CredentialsClass C.GCredentialsClass
type DBusAnnotationInfo C.GDBusAnnotationInfo
type DBusArgInfo C.GDBusArgInfo
type DBusErrorEntry C.GDBusErrorEntry
type DBusInterfaceInfo C.GDBusInterfaceInfo
type DBusInterfaceVTable C.GDBusInterfaceVTable
type DBusMethodInfo C.GDBusMethodInfo
type DBusNodeInfo C.GDBusNodeInfo
type DBusPropertyInfo C.GDBusPropertyInfo
type DBusProxyClass C.GDBusProxyClass
type DBusSignalInfo C.GDBusSignalInfo
type DBusSubtreeVTable C.GDBusSubtreeVTable

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
type ProxyAddressClass C.GProxyAddressClass
type ProxyInterface C.GProxyInterface

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type TlsClientConnectionInterface C.GTlsClientConnectionInterface
type TlsServerConnectionInterface C.GTlsServerConnectionInterface
type UnixCredentialsMessageClass C.GUnixCredentialsMessageClass

func Fn_g_dbus_annotation_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusAnnotationInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_annotation_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_annotation_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusAnnotationInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_annotation_info_unref(cValueInstance)
}

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

func Fn_g_dbus_arg_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusArgInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_arg_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_arg_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusArgInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_arg_info_unref(cValueInstance)
}

func Fn_g_dbus_interface_info_generate_xml(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	C.g_dbus_interface_info_generate_xml(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_interface_info_lookup_method(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_method(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_signal(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_signal(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_unref(cValueInstance)
}

func Fn_g_dbus_method_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_info_unref(cValueInstance)
}

func Fn_g_dbus_node_info_new_for_xml(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_node_info_new_for_xml(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_generate_xml(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	C.g_dbus_node_info_generate_xml(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_node_info_lookup_interface(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_node_info_lookup_interface(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_node_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_node_info_unref(cValueInstance)
}

func Fn_g_dbus_property_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusPropertyInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_property_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_property_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusPropertyInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_property_info_unref(cValueInstance)
}

func Fn_g_dbus_signal_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusSignalInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_signal_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_signal_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusSignalInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_signal_info_unref(cValueInstance)
}

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

func Fn_g_bus_get_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_bus_get_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_sync(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

func Fn_g_bus_own_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameOwnerFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_own_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

func Fn_g_bus_own_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) uint {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameOwnerFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (*C.GClosure)(unsafe.Pointer(param5))

	ret := C.g_bus_own_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return (uint)(ret)
}

func Fn_g_bus_unown_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unown_name(cValue0)
}

func Fn_g_bus_unwatch_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unwatch_name(cValue0)
}

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

func Fn_g_bus_watch_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameWatcherFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_watch_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

func Fn_g_bus_watch_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameWatcherFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_watch_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

func Fn_g_dbus_address_get_for_bus_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_for_bus_sync(cValue0, cValue1, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

func Fn_g_dbus_address_get_stream_finish(param0 unsafe.Pointer, param1 *string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_stream_finish(cValue0, cValue1, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_address_get_stream_sync(param0 string, param1 *string, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_stream_sync(cValue0, cValue1, cValue2, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

func Fn_g_dbus_generate_guid() string {
	ret := C.g_dbus_generate_guid()

	return C.GoString(ret)
}

func Fn_g_dbus_is_address(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_address(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_guid(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_guid(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_interface_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_interface_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_member_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_member_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_supported_address(param0 string, error unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_is_supported_address(cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_is_unique_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_unique_name(cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

// UNSUPPORTED : g_application_add_main_option_entries : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : g_application_command_line_get_environ : no array length

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

func Fn_g_credentials_new() unsafe.Pointer {
	ret := C.g_credentials_new()

	return unsafe.Pointer(ret)
}

func Fn_g_credentials_get_native(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GCredentialsType)(param0)

	ret := C.g_credentials_get_native(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_credentials_get_unix_user(paramInstance unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_credentials_get_unix_user(cValueInstance, cError)

	return (uint)(ret)
}

func Fn_g_credentials_is_same_user(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_credentials_is_same_user(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_set_native(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GCredentialsType)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_credentials_set_native(cValueInstance, cValue0, cValue1)
}

func Fn_g_credentials_set_unix_user(paramInstance unsafe.Pointer, param0 uint, error unsafe.Pointer) bool {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.uid_t)(param0)

	cError := (**C.GError)(error)

	ret := C.g_credentials_set_unix_user(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	ret := C.g_credentials_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_auth_observer_new() unsafe.Pointer {
	ret := C.g_dbus_auth_observer_new()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_auth_observer_authorize_authenticated_peer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusAuthObserver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GCredentials)(unsafe.Pointer(param1))

	ret := C.g_dbus_auth_observer_authorize_authenticated_peer(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_sync(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GDBusConnectionFlags)(param1)

	cValue2 := (*C.GDBusAuthObserver)(unsafe.Pointer(param2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_sync(cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_sync(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GDBusConnectionFlags)(param2)

	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

func Fn_g_dbus_connection_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_call_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (*C.GVariantType)(unsafe.Pointer(param5))

	cValue6 := (C.GDBusCallFlags)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (*C.GCancellable)(unsafe.Pointer(param8))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

func Fn_g_dbus_connection_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_close_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_sync(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_emit_signal(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_emit_signal(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

func Fn_g_dbus_connection_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_flush_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_sync(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_capabilities(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_capabilities(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_connection_get_exit_on_close(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_exit_on_close(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_guid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_guid(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_get_peer_credentials(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_peer_credentials(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_unique_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_unique_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_is_closed(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

func Fn_g_dbus_connection_remove_filter(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_remove_filter(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_send_message(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *uint32, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusSendMessageFlags)(param1)

	cValue2 := (*C.guint32)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

func Fn_g_dbus_connection_send_message_with_reply_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_send_message_with_reply_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *uint32, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusSendMessageFlags)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_set_exit_on_close(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_dbus_connection_set_exit_on_close(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

func Fn_g_dbus_connection_signal_unsubscribe(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_signal_unsubscribe(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_start_message_processing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_start_message_processing(cValueInstance)
}

func Fn_g_dbus_connection_unregister_object(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_connection_unregister_object(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_unregister_subtree(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_connection_unregister_subtree(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

func Fn_g_dbus_message_new() unsafe.Pointer {
	ret := C.g_dbus_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_from_blob(param0 []uint8, param1 uint64, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.guchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (C.GDBusCapabilityFlags)(param2)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_new_from_blob(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_call(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.g_dbus_message_new_method_call(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_signal(param0 string, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_message_new_signal(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_copy(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_copy(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_arg0(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_arg0(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_body(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_body(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_destination(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_destination(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_error_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_error_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_message_get_header(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageHeaderField)(param0)

	ret := C.g_dbus_message_get_header(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : no array length

func Fn_g_dbus_message_get_interface(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_interface(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_locked(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_locked(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_message_get_member(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_member(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_message_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_message_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_message_get_num_unix_fds(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_num_unix_fds(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_reply_serial(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_reply_serial(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_sender(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_sender(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_serial(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_serial(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_signature(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_signature(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_unix_fd_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_unix_fd_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_lock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_lock(cValueInstance)
}

func Fn_g_dbus_message_new_method_error(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_g_dbus_message_new_method_error(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_error_literal(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_message_new_method_error_literal(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_error_valist(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_g_dbus_message_new_method_error_valist(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_reply(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_new_method_reply(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_print(paramInstance unsafe.Pointer, param0 uint) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_message_print(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_dbus_message_set_body(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_dbus_message_set_body(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_destination(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_destination(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_error_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_error_name(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageFlags)(param0)

	C.g_dbus_message_set_flags(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_header(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageHeaderField)(param0)

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_dbus_message_set_header(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_message_set_interface(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_member(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_member(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_message_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageType)(param0)

	C.g_dbus_message_set_message_type(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_num_unix_fds(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_num_unix_fds(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_path(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_reply_serial(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_reply_serial(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_sender(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_sender(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_serial(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_serial(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_signature(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_signature(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_unix_fd_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GUnixFDList)(unsafe.Pointer(param0))

	C.g_dbus_message_set_unix_fd_list(cValueInstance, cValue0)
}

func Fn_g_dbus_message_to_blob(paramInstance unsafe.Pointer, param0 *uint64, param1 int, error unsafe.Pointer) []uint8 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusCapabilityFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_to_blob(cValueInstance, cValue0, cValue1, cError)

	retLen := int(*cValue0)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_dbus_message_to_gerror(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_to_gerror(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_message_bytes_needed(param0 []uint8, param1 uint64, error unsafe.Pointer) uint64 {
	cValue0 := (*C.guchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_bytes_needed(cValue0, cValue1, cError)

	return (uint64)(ret)
}

func Fn_g_dbus_method_invocation_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_interface_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_interface_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_message(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_message(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_method_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_method_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_parameters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_parameters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_sender(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_sender(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_user_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_user_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_dbus_method_invocation_return_dbus_error(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_method_invocation_return_dbus_error(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_method_invocation_return_error(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.c_g_dbus_method_invocation_return_error(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_dbus_method_invocation_return_error_literal(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_method_invocation_return_error_literal(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_dbus_method_invocation_return_error_valist(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.c_g_dbus_method_invocation_return_error_valist(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_dbus_method_invocation_return_gerror(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_return_gerror(cValueInstance, cValue0)
}

func Fn_g_dbus_method_invocation_return_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_return_value(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : parameter 'get_proxy_type_func' is callback

func Fn_g_dbus_proxy_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_for_bus_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_sync(param0 int, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (C.GDBusProxyFlags)(param1)

	cValue2 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_for_bus_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_sync(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusProxyFlags)(param1)

	cValue2 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call : parameter 'callback' is callback

func Fn_g_dbus_proxy_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_call_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (C.GDBusCallFlags)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : parameter 'callback' is callback

func Fn_g_dbus_proxy_get_cached_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_proxy_get_cached_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : no array length

func Fn_g_dbus_proxy_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_default_timeout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_default_timeout(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_interface_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_interface_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_interface_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_interface_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name_owner(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_name_owner(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_set_cached_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_dbus_proxy_set_cached_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_proxy_set_default_timeout(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_dbus_proxy_set_default_timeout(cValueInstance, cValue0)
}

func Fn_g_dbus_proxy_set_interface_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param0))

	C.g_dbus_proxy_set_interface_info(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_proxy_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : parameter 'callback' is callback

func Fn_g_dbus_server_new_sync(param0 string, param1 int, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GDBusServerFlags)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_server_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_server_get_client_address(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_client_address(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_server_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_server_get_guid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_guid(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_server_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_server_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_start(cValueInstance)
}

func Fn_g_dbus_server_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_stop(cValueInstance)
}

// UNSUPPORTED : g_data_input_stream_read_line : no array length

// UNSUPPORTED : g_data_input_stream_read_line_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : no array length

// UNSUPPORTED : g_data_input_stream_read_until_async : parameter 'callback' is callback

func Fn_g_data_input_stream_read_upto(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_upto(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_upto_async : parameter 'callback' is callback

// UNSUPPORTED : g_desktop_app_info_get_keywords : no array length

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_list_actions : no array length

// UNSUPPORTED : g_desktop_app_info_search : array has no type

// UNSUPPORTED : g_file_enumerator_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_enumerator_next_files_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_io_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_info_get_attribute_stringv : no array length

// UNSUPPORTED : g_file_info_list_attributes : no array length

// UNSUPPORTED : g_file_info_set_attribute_stringv : parameter 'attr_value' is array parameter without length parameter

// UNSUPPORTED : g_file_input_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_output_stream_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_filename_completer_get_completions : no array length

// UNSUPPORTED : g_io_module_new : blacklisted

// UNSUPPORTED : g_io_module_load : blacklisted

// UNSUPPORTED : g_io_module_unload : blacklisted

// UNSUPPORTED : g_io_module_query : blacklisted

// UNSUPPORTED : g_io_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_io_stream_splice_async : parameter 'callback' is callback

// UNSUPPORTED : g_inet_address_new_from_bytes : parameter 'bytes' is array parameter without length parameter

// UNSUPPORTED : g_input_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read : blacklisted

// UNSUPPORTED : g_input_stream_read_all : blacklisted

// UNSUPPORTED : g_input_stream_read_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_skip_async : parameter 'callback' is callback

// UNSUPPORTED : g_list_store_insert_sorted : parameter 'compare_func' is callback

// UNSUPPORTED : g_list_store_sort : parameter 'compare_func' is callback

// UNSUPPORTED : g_memory_input_stream_new_from_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_input_stream_add_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_output_stream_new : parameter 'realloc_function' is callback

func Fn_g_memory_output_stream_steal_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_steal_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_network_address_get_scheme(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_scheme(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_address_parse_uri(param0 string, param1 uint16, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse_uri(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_get_scheme(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_scheme(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_set_scheme(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_network_service_set_scheme(cValueInstance, cValue0)
}

// UNSUPPORTED : g_output_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_flush_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_splice_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_async : parameter 'callback' is callback

func Fn_g_permission_acquire(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_acquire_async : parameter 'callback' is callback

func Fn_g_permission_acquire_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_permission_get_allowed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_allowed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_acquire(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_can_acquire(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_release(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_can_release(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_impl_update(paramInstance unsafe.Pointer, param0 bool, param1 bool, param2 bool) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	C.g_permission_impl_update(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_permission_release(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_release(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_release_async : parameter 'callback' is callback

func Fn_g_permission_release_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_release_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_proxy_address_new(param0 unsafe.Pointer, param1 uint16, param2 string, param3 string, param4 uint16, param5 string, param6 string) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.guint16)(param4)

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (*C.gchar)(C.CString(param6))
	defer C.free(unsafe.Pointer(cValue6))

	ret := C.g_proxy_address_new(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_address_get_destination_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_destination_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_destination_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_destination_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_proxy_address_get_password(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_password(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_protocol(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_protocol(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_username(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_username(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_address_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_records_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_service_async : parameter 'callback' is callback

func Fn_g_settings_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	ret := C.g_settings_new_with_backend(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend_and_path(param0 string, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_settings_new_with_backend_and_path(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_path(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_settings_new_with_path(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_bind(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GSettingsBindFlags)(param3)

	C.g_settings_bind(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

func Fn_g_settings_bind_writable(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 bool) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := toCBool(param3)

	C.g_settings_bind_writable(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_settings_delay(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_delay(cValueInstance)
}

func Fn_g_settings_get(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_g_settings_get(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_get_boolean(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_boolean(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_settings_get_child(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_double(paramInstance unsafe.Pointer, param0 string) float64 {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_double(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_g_settings_get_enum(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_enum(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_settings_get_flags(paramInstance unsafe.Pointer, param0 string) uint {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_flags(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_g_settings_get_has_unapplied(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_get_has_unapplied(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_settings_get_int(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_int(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

func Fn_g_settings_get_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_get_strv : no array length

func Fn_g_settings_get_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_is_writable(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_is_writable(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_children : no array length

// UNSUPPORTED : g_settings_list_keys : no array length

func Fn_g_settings_set(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.c_g_settings_set(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_boolean(paramInstance unsafe.Pointer, param0 string, param1 bool) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.g_settings_set_boolean(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	ret := C.g_settings_set_double(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.g_settings_set_int(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_string(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_settings_set_string(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_set_strv : parameter 'value' is array parameter without length parameter

func Fn_g_settings_set_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	ret := C.g_settings_set_value(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_relocatable_schemas : no array length

// UNSUPPORTED : g_settings_list_schemas : no array length

func Fn_g_settings_unbind(param0 unsafe.Pointer, param1 string) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_settings_unbind(cValue0, cValue1)
}

// UNSUPPORTED : g_settings_backend_changed : blacklisted

// UNSUPPORTED : g_settings_backend_changed_tree : blacklisted

// UNSUPPORTED : g_settings_backend_keys_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_flatten_tree : blacklisted

// UNSUPPORTED : g_settings_backend_get_default : blacklisted

// UNSUPPORTED : g_simple_async_result_new : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_from_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_take_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_run_in_thread : parameter 'func' is callback

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

func Fn_g_simple_permission_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.g_simple_permission_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_credentials(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_credentials(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_timeout(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_timeout(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_receive_with_blocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 bool, param3 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive_with_blocking(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_with_blocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 bool, param3 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_with_blocking(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return (uint64)(ret)
}

func Fn_g_socket_set_timeout(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_socket_set_timeout(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_uri(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_uri_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_enable_proxy(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_enable_proxy(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_client_get_timeout(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_timeout(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_client_set_enable_proxy(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_socket_client_set_enable_proxy(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_timeout(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_socket_client_set_timeout(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_newv : parameter 'argv' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_communicate_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_communicate_utf8_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_check_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : parameter 'child_setup' is callback

// UNSUPPORTED : g_subprocess_launcher_set_environ : parameter 'env' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_launcher_spawnv : parameter 'argv' is array parameter without length parameter

// UNSUPPORTED : g_task_new : parameter 'callback' is callback

// UNSUPPORTED : g_task_attach_source : parameter 'callback' is callback

// UNSUPPORTED : g_task_return_pointer : parameter 'result_destroy' is callback

// UNSUPPORTED : g_task_run_in_thread : parameter 'task_func' is callback

// UNSUPPORTED : g_task_run_in_thread_sync : parameter 'task_func' is callback

// UNSUPPORTED : g_task_set_task_data : parameter 'task_data_destroy' is callback

// UNSUPPORTED : g_task_report_error : parameter 'callback' is callback

// UNSUPPORTED : g_task_report_new_error : parameter 'callback' is callback

// UNSUPPORTED : g_themed_icon_get_names : no array length

// UNSUPPORTED : g_tls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_verify_chain_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_ask_password_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

func Fn_g_unix_connection_receive_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_credentials(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_send_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_credentials(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

func Fn_g_unix_credentials_message_new() unsafe.Pointer {
	ret := C.g_unix_credentials_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_new_with_credentials(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))

	ret := C.g_unix_credentials_message_new_with_credentials(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_get_credentials(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixCredentialsMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_credentials_message_get_credentials(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_is_supported() bool {
	ret := C.g_unix_credentials_message_is_supported()

	return toGoBool(ret)
}

func Fn_g_unix_socket_address_new_with_type(param0 []int8, param1 int, param2 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GUnixSocketAddressType)(param2)

	ret := C.g_unix_socket_address_new_with_type(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_get_address_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_address_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

func Fn_g_zlib_compressor_get_file_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))

	ret := C.g_zlib_compressor_get_file_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_set_file_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	C.g_zlib_compressor_set_file_info(cValueInstance, cValue0)
}

func Fn_g_zlib_decompressor_get_file_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GZlibDecompressor)(unsafe.Pointer(paramInstance))

	ret := C.g_zlib_decompressor_get_file_info(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_action_group_list_actions : no array length

// UNSUPPORTED : g_app_info_get_supported_types : no array length

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

// UNSUPPORTED : g_drive_poll_for_media : parameter 'callback' is callback

// UNSUPPORTED : g_drive_start : parameter 'callback' is callback

// UNSUPPORTED : g_drive_stop : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_dtls_connection_shutdown_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_append_to_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_copy : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_copy_async : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_create_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_create_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_delete_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_eject_mountable_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_file_enumerate_children_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_find_enclosing_mount_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents : blacklisted

// UNSUPPORTED : g_file_load_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents_finish : blacklisted

// UNSUPPORTED : g_file_load_partial_contents_async : parameter 'read_more_callback' is callback

// UNSUPPORTED : g_file_load_partial_contents_finish : blacklisted

// UNSUPPORTED : g_file_make_directory_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage_async : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_mount_enclosing_volume : parameter 'callback' is callback

// UNSUPPORTED : g_file_mount_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_move : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_open_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_poll_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_default_handler_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_filesystem_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_query_info_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_attributes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_display_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_start_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_stop_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_trash_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_loadable_icon_load_async : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject : parameter 'callback' is callback

// UNSUPPORTED : g_mount_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

func Fn_g_proxy_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GProxyAddress)(unsafe.Pointer(param1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

func Fn_g_proxy_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_supports_hostname(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_supports_hostname(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_proxy_resolver_is_supported(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GProxyResolver)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_resolver_is_supported(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

func Fn_g_socket_connectable_proxy_enumerate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnectable)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connectable_proxy_enumerate(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback
