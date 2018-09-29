// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// BusGetSync is a wrapper around the C function g_bus_get_sync.
func BusGetSync(busType BusType, cancellable *Cancellable) (*DBusConnection, error) {
	c_bus_type := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_bus_get_sync(c_bus_type, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback, GBusAcquiredCallback

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback, GBusNameAcquiredCallback

// BusOwnNameOnConnectionWithClosures is a wrapper around the C function g_bus_own_name_on_connection_with_closures.
func BusOwnNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	c_connection := (*C.GDBusConnection)(connection.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_name_acquired_closure := (*C.GClosure)(nameAcquiredClosure.ToC())

	c_name_lost_closure := (*C.GClosure)(nameLostClosure.ToC())

	retC := C.g_bus_own_name_on_connection_with_closures(c_connection, c_name, c_flags, c_name_acquired_closure, c_name_lost_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusOwnNameWithClosures is a wrapper around the C function g_bus_own_name_with_closures.
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_bus_acquired_closure := (*C.GClosure)(busAcquiredClosure.ToC())

	c_name_acquired_closure := (*C.GClosure)(nameAcquiredClosure.ToC())

	c_name_lost_closure := (*C.GClosure)(nameLostClosure.ToC())

	retC := C.g_bus_own_name_with_closures(c_bus_type, c_name, c_flags, c_bus_acquired_closure, c_name_acquired_closure, c_name_lost_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusUnownName is a wrapper around the C function g_bus_unown_name.
func BusUnownName(ownerId uint32) {
	c_owner_id := (C.guint)(ownerId)

	C.g_bus_unown_name(c_owner_id)

	return
}

// BusUnwatchName is a wrapper around the C function g_bus_unwatch_name.
func BusUnwatchName(watcherId uint32) {
	c_watcher_id := (C.guint)(watcherId)

	C.g_bus_unwatch_name(c_watcher_id)

	return
}

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback, GBusNameAppearedCallback

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback, GBusNameAppearedCallback

// BusWatchNameOnConnectionWithClosures is a wrapper around the C function g_bus_watch_name_on_connection_with_closures.
func BusWatchNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	c_connection := (*C.GDBusConnection)(connection.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_name_appeared_closure := (*C.GClosure)(nameAppearedClosure.ToC())

	c_name_vanished_closure := (*C.GClosure)(nameVanishedClosure.ToC())

	retC := C.g_bus_watch_name_on_connection_with_closures(c_connection, c_name, c_flags, c_name_appeared_closure, c_name_vanished_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusWatchNameWithClosures is a wrapper around the C function g_bus_watch_name_with_closures.
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_name_appeared_closure := (*C.GClosure)(nameAppearedClosure.ToC())

	c_name_vanished_closure := (*C.GClosure)(nameVanishedClosure.ToC())

	retC := C.g_bus_watch_name_with_closures(c_bus_type, c_name, c_flags, c_name_appeared_closure, c_name_vanished_closure)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_content_type_get_icon : no return generator

// Unsupported : g_content_type_get_symbolic_icon : no return generator

// Unsupported : g_content_type_guess : unsupported parameter data : no param type

// Unsupported : g_content_type_guess_for_tree : unsupported parameter root : no type generator for File, GFile*

// DbusAddressGetForBusSync is a wrapper around the C function g_dbus_address_get_for_bus_sync.
func DbusAddressGetForBusSync(busType BusType, cancellable *Cancellable) (string, error) {
	c_bus_type := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_for_bus_sync(c_bus_type, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_address_get_stream_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// DbusAddressGetStreamSync is a wrapper around the C function g_dbus_address_get_stream_sync.
func DbusAddressGetStreamSync(address string, cancellable *Cancellable) (*IOStream, string, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	var c_out_guid *C.gchar

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_stream_sync(c_address, &c_out_guid, c_cancellable, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outGuid := C.GoString(c_out_guid)
	defer C.free(unsafe.Pointer(c_out_guid))

	return retGo, outGuid, goThrowableError
}

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations : no param type

// DbusErrorEncodeGerror is a wrapper around the C function g_dbus_error_encode_gerror.
func DbusErrorEncodeGerror(error *glib.Error) string {
	c_error := (*C.GError)(error.ToC())

	retC := C.g_dbus_error_encode_gerror(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DbusErrorGetRemoteError is a wrapper around the C function g_dbus_error_get_remote_error.
func DbusErrorGetRemoteError(error *glib.Error) string {
	c_error := (*C.GError)(error.ToC())

	retC := C.g_dbus_error_get_remote_error(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DbusErrorIsRemoteError is a wrapper around the C function g_dbus_error_is_remote_error.
func DbusErrorIsRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(error.ToC())

	retC := C.g_dbus_error_is_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DbusErrorNewForDbusError is a wrapper around the C function g_dbus_error_new_for_dbus_error.
func DbusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	retC := C.g_dbus_error_new_for_dbus_error(c_dbus_error_name, c_dbus_error_message)
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DbusErrorRegisterError is a wrapper around the C function g_dbus_error_register_error.
func DbusErrorRegisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_register_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter quark_volatile : no type generator for gsize, volatile gsize*

// DbusErrorStripRemoteError is a wrapper around the C function g_dbus_error_strip_remote_error.
func DbusErrorStripRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(error.ToC())

	retC := C.g_dbus_error_strip_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DbusErrorUnregisterError is a wrapper around the C function g_dbus_error_unregister_error.
func DbusErrorUnregisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_unregister_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// DbusGenerateGuid is a wrapper around the C function g_dbus_generate_guid.
func DbusGenerateGuid() string {
	retC := C.g_dbus_generate_guid()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_gvalue_to_gvariant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_dbus_gvariant_to_gvalue : unsupported parameter value : Blacklisted record : GVariant

// DbusIsAddress is a wrapper around the C function g_dbus_is_address.
func DbusIsAddress(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_address(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsGuid is a wrapper around the C function g_dbus_is_guid.
func DbusIsGuid(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_guid(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsInterfaceName is a wrapper around the C function g_dbus_is_interface_name.
func DbusIsInterfaceName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_interface_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsMemberName is a wrapper around the C function g_dbus_is_member_name.
func DbusIsMemberName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_member_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsName is a wrapper around the C function g_dbus_is_name.
func DbusIsName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsSupportedAddress is a wrapper around the C function g_dbus_is_supported_address.
func DbusIsSupportedAddress(string string) (bool, error) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	var cThrowableError *C.GError

	retC := C.g_dbus_is_supported_address(c_string, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// DbusIsUniqueName is a wrapper around the C function g_dbus_is_unique_name.
func DbusIsUniqueName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_unique_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

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

// Unsupported : g_resources_get_info : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_tls_backend_get_default : no return generator

// Unsupported : g_tls_client_connection_new : unsupported parameter server_identity : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_tls_file_database_new : no return generator

// Unsupported : g_tls_server_connection_new : no return generator

// Unsupported : g_unix_mount_at : unsupported parameter time_read : no type generator for guint64, guint64*

// Unsupported : g_unix_mount_for : unsupported parameter time_read : no type generator for guint64, guint64*

// Unsupported : g_unix_mount_guess_icon : no return generator

// Unsupported : g_unix_mount_guess_symbolic_icon : no return generator

// Unsupported : g_unix_mount_points_get : unsupported parameter time_read : no type generator for guint64, guint64*

// Unsupported : g_unix_mounts_get : unsupported parameter time_read : no type generator for guint64, guint64*
