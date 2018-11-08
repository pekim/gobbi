// This is a generated file - DO NOT EDIT

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

// AppInfoGetAll is a wrapper around the C function g_app_info_get_all.
func AppInfoGetAll() *glib.List {
	retC := C.g_app_info_get_all()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetAllForType is a wrapper around the C function g_app_info_get_all_for_type.
func AppInfoGetAllForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_all_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_app_info_get_default_for_type : no return generator

// Unsupported : g_app_info_get_default_for_uri_scheme : no return generator

// AppInfoLaunchDefaultForUri is a wrapper around the C function g_app_info_launch_default_for_uri.
func AppInfoLaunchDefaultForUri(uri string, context *AppLaunchContext) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_context := (*C.GAppLaunchContext)(context.ToC())

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_default_for_uri(c_uri, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_app_info_launch_default_for_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_app_info_launch_default_for_uri_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// Unsupported : g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_bus_get_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback (GBusAcquiredCallback) for param bus_acquired_handler

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback (GBusNameAcquiredCallback) for param name_acquired_handler

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// ContentTypeCanBeExecutable is a wrapper around the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_can_be_executable(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeEquals is a wrapper around the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	c_type1 := C.CString(type1)
	defer C.free(unsafe.Pointer(c_type1))

	c_type2 := C.CString(type2)
	defer C.free(unsafe.Pointer(c_type2))

	retC := C.g_content_type_equals(c_type1, c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeGetDescription is a wrapper around the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_description(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_content_type_get_icon : no return generator

// ContentTypeGetMimeType is a wrapper around the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_mime_type(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_content_type_get_symbolic_icon : no return generator

// Unsupported : g_content_type_guess : unsupported parameter data : no type generator for guint8 (guchar) for array param data

// Unsupported : g_content_type_guess_for_tree : unsupported parameter root : no type generator for File (GFile*) for param root

// ContentTypeIsA is a wrapper around the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_supertype := C.CString(supertype)
	defer C.free(unsafe.Pointer(c_supertype))

	retC := C.g_content_type_is_a(c_type, c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeIsUnknown is a wrapper around the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_is_unknown(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypesGetRegistered is a wrapper around the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retC := C.g_content_types_get_registered()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dbus_address_get_stream_finish : unsupported parameter res : no type generator for AsyncResult (GAsyncResult*) for param res

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations : no type generator for DBusAnnotationInfo (GDBusAnnotationInfo*) for array param annotations

// DbusErrorQuark is a wrapper around the C function g_dbus_error_quark.
func DbusErrorQuark() glib.Quark {
	retC := C.g_dbus_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries : no type generator for DBusErrorEntry (GDBusErrorEntry) for array param entries

// Unsupported : g_dbus_gvalue_to_gvariant : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_dbus_gvariant_to_gvalue : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_dtls_client_connection_new : unsupported parameter base_socket : no type generator for DatagramBased (GDatagramBased*) for param base_socket

// Unsupported : g_dtls_server_connection_new : unsupported parameter base_socket : no type generator for DatagramBased (GDatagramBased*) for param base_socket

// Unsupported : g_file_new_for_commandline_arg : no return generator

// Unsupported : g_file_new_for_commandline_arg_and_cwd : no return generator

// Unsupported : g_file_new_for_path : no return generator

// Unsupported : g_file_new_for_uri : no return generator

// Unsupported : g_file_new_tmp : unsupported parameter iostream : record with indirection level of 2

// Unsupported : g_file_parse_name : no return generator

// Unsupported : g_icon_deserialize : unsupported parameter value : Blacklisted record : GVariant

// IconHash is a wrapper around the C function g_icon_hash.
func IconHash(icon uintptr) uint32 {
	c_icon := (C.gconstpointer)(icon)

	retC := C.g_icon_hash(c_icon)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_icon_new_for_string : no return generator

// Unsupported : g_initable_newv : unsupported parameter parameters : no type generator for GObject.Parameter (GParameter) for array param parameters

// IoErrorFromErrno is a wrapper around the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	c_err_no := (C.gint)(errNo)

	retC := C.g_io_error_from_errno(c_err_no)
	retGo := (IOErrorEnum)(retC)

	return retGo
}

// IoErrorQuark is a wrapper around the C function g_io_error_quark.
func IoErrorQuark() glib.Quark {
	retC := C.g_io_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// IoExtensionPointImplement is a wrapper around the C function g_io_extension_point_implement.
func IoExtensionPointImplement(extensionPointName string, type_ gobject.Type, extensionName string, priority int32) *IOExtension {
	c_extension_point_name := C.CString(extensionPointName)
	defer C.free(unsafe.Pointer(c_extension_point_name))

	c_type := (C.GType)(type_)

	c_extension_name := C.CString(extensionName)
	defer C.free(unsafe.Pointer(c_extension_name))

	c_priority := (C.gint)(priority)

	retC := C.g_io_extension_point_implement(c_extension_point_name, c_type, c_extension_name, c_priority)
	retGo := IOExtensionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoExtensionPointLookup is a wrapper around the C function g_io_extension_point_lookup.
func IoExtensionPointLookup(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_lookup(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoExtensionPointRegister is a wrapper around the C function g_io_extension_point_register.
func IoExtensionPointRegister(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_register(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoModulesLoadAllInDirectory is a wrapper around the C function g_io_modules_load_all_in_directory.
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	retC := C.g_io_modules_load_all_in_directory(c_dirname)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoSchedulerCancelAllJobs is a wrapper around the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()

	return
}

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Blacklisted : g_keyfile_settings_backend_new

// Unsupported : g_network_monitor_get_default : no return generator

// Unsupported : g_pollable_stream_read : unsupported parameter buffer : no type generator for guint8 () for array param buffer

// Unsupported : g_pollable_stream_write : unsupported parameter buffer : no type generator for guint8 () for array param buffer

// Unsupported : g_pollable_stream_write_all : unsupported parameter buffer : no type generator for guint8 () for array param buffer

// Unsupported : g_proxy_get_default_for_protocol : no return generator

// Unsupported : g_proxy_resolver_get_default : no return generator

// Unsupported : g_resources_enumerate_children : no return type

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_tls_backend_get_default : no return generator

// Unsupported : g_tls_client_connection_new : unsupported parameter server_identity : no type generator for SocketConnectable (GSocketConnectable*) for param server_identity

// Unsupported : g_tls_file_database_new : no return generator

// Unsupported : g_tls_server_connection_new : no return generator

// UnixIsMountPathSystemInternal is a wrapper around the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	retC := C.g_unix_is_mount_path_system_internal(c_mount_path)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountAt is a wrapper around the C function g_unix_mount_at.
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_at(c_mount_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountCompare is a wrapper around the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	c_mount1 := (*C.GUnixMountEntry)(mount1.ToC())

	c_mount2 := (*C.GUnixMountEntry)(mount2.ToC())

	retC := C.g_unix_mount_compare(c_mount1, c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// UnixMountFree is a wrapper around the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	C.g_unix_mount_free(c_mount_entry)

	return
}

// UnixMountGetDevicePath is a wrapper around the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_get_device_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetFsType is a wrapper around the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_get_fs_type(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetMountPath is a wrapper around the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_get_mount_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGuessCanEject is a wrapper around the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_guess_can_eject(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_unix_mount_guess_icon : no return generator

// UnixMountGuessName is a wrapper around the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_guess_name(c_mount_entry)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnixMountGuessShouldDisplay is a wrapper around the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_guess_should_display(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_unix_mount_guess_symbolic_icon : no return generator

// UnixMountIsReadonly is a wrapper around the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_is_readonly(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountIsSystemInternal is a wrapper around the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(mountEntry.ToC())

	retC := C.g_unix_mount_is_system_internal(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsChangedSince is a wrapper around the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mount_points_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsGet is a wrapper around the C function g_unix_mount_points_get.
func UnixMountPointsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mount_points_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountsChangedSince is a wrapper around the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mounts_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountsGet is a wrapper around the C function g_unix_mounts_get.
func UnixMountsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mounts_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}
