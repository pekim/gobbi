// Code generated - DO NOT EDIT.
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// NETWORK_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_NETWORK_MONITOR_EXTENSION_POINT_NAME.
//
// since 2.30
const NETWORK_MONITOR_EXTENSION_POINT_NAME = "gio-network-monitor"

// DBusInterfaceSkeletonFlags is a representation of the C bitfield GDBusInterfaceSkeletonFlags.
type DBusInterfaceSkeletonFlags int

// DBusInterfaceSkeletonFlags_none is a representation of the C bitfield member G_DBUS_INTERFACE_SKELETON_FLAGS_NONE.
const DBusInterfaceSkeletonFlags_none = DBusInterfaceSkeletonFlags(0)

// DBusInterfaceSkeletonFlags_handle_method_invocations_in_thread is a representation of the C bitfield member G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD.
const DBusInterfaceSkeletonFlags_handle_method_invocations_in_thread = DBusInterfaceSkeletonFlags(1)

// DBusObjectManagerClientFlags is a representation of the C bitfield GDBusObjectManagerClientFlags.
type DBusObjectManagerClientFlags int

// DBusObjectManagerClientFlags_none is a representation of the C bitfield member G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE.
const DBusObjectManagerClientFlags_none = DBusObjectManagerClientFlags(0)

// DBusObjectManagerClientFlags_do_not_auto_start is a representation of the C bitfield member G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START.
const DBusObjectManagerClientFlags_do_not_auto_start = DBusObjectManagerClientFlags(1)

// TlsDatabaseVerifyFlags is a representation of the C bitfield GTlsDatabaseVerifyFlags.
type TlsDatabaseVerifyFlags int

// TlsDatabaseVerifyFlags_none is a representation of the C bitfield member G_TLS_DATABASE_VERIFY_NONE.
const TlsDatabaseVerifyFlags_none = TlsDatabaseVerifyFlags(0)

// TlsPasswordFlags is a representation of the C bitfield GTlsPasswordFlags.
type TlsPasswordFlags int

// TlsPasswordFlags_none is a representation of the C bitfield member G_TLS_PASSWORD_NONE.
const TlsPasswordFlags_none = TlsPasswordFlags(0)

// TlsPasswordFlags_retry is a representation of the C bitfield member G_TLS_PASSWORD_RETRY.
const TlsPasswordFlags_retry = TlsPasswordFlags(2)

// TlsPasswordFlags_many_tries is a representation of the C bitfield member G_TLS_PASSWORD_MANY_TRIES.
const TlsPasswordFlags_many_tries = TlsPasswordFlags(4)

// TlsPasswordFlags_final_try is a representation of the C bitfield member G_TLS_PASSWORD_FINAL_TRY.
const TlsPasswordFlags_final_try = TlsPasswordFlags(8)

// IOModuleScopeFlags is a representation of the C enumeration GIOModuleScopeFlags.
type IOModuleScopeFlags int

// IOModuleScopeFlags_none is a representation of the C enumeration member G_IO_MODULE_SCOPE_NONE.
const IOModuleScopeFlags_none = IOModuleScopeFlags(0)

// IOModuleScopeFlags_block_duplicates is a representation of the C enumeration member G_IO_MODULE_SCOPE_BLOCK_DUPLICATES.
const IOModuleScopeFlags_block_duplicates = IOModuleScopeFlags(1)

// TlsDatabaseLookupFlags is a representation of the C enumeration GTlsDatabaseLookupFlags.
type TlsDatabaseLookupFlags int

// TlsDatabaseLookupFlags_none is a representation of the C enumeration member G_TLS_DATABASE_LOOKUP_NONE.
const TlsDatabaseLookupFlags_none = TlsDatabaseLookupFlags(0)

// TlsDatabaseLookupFlags_keypair is a representation of the C enumeration member G_TLS_DATABASE_LOOKUP_KEYPAIR.
const TlsDatabaseLookupFlags_keypair = TlsDatabaseLookupFlags(1)

// TlsInteractionResult is a representation of the C enumeration GTlsInteractionResult.
type TlsInteractionResult int

// TlsInteractionResult_unhandled is a representation of the C enumeration member G_TLS_INTERACTION_UNHANDLED.
const TlsInteractionResult_unhandled = TlsInteractionResult(0)

// TlsInteractionResult_handled is a representation of the C enumeration member G_TLS_INTERACTION_HANDLED.
const TlsInteractionResult_handled = TlsInteractionResult(1)

// TlsInteractionResult_failed is a representation of the C enumeration member G_TLS_INTERACTION_FAILED.
const TlsInteractionResult_failed = TlsInteractionResult(2)

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

// DbusGvalueToGvariant wraps the C function g_dbus_gvalue_to_gvariant.
//
// since 2.30
func DbusGvalueToGvariant(gvalue *gobject.Value, type_ *glib.VariantType) *glib.Variant {
	sys_gvalue := gvalue.ToC()
	sys_type_ := type_.ToC()
	retSys := gio.Fn_g_dbus_gvalue_to_gvariant(sys_gvalue, sys_type_)
	ret := glib.VariantNewFromC(retSys)

	return ret
}

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

// DBusInterfaceIface is a representation of the C record GDBusInterfaceIface.
//
// since 2.30
type DBusInterfaceIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceIface that represents the DBusInterfaceIface.
func (recv *DBusInterfaceIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceIfaceNewFromC creates a new DBusInterfaceIface from a pointer to the C GDBusInterfaceIface that represents the DBusInterfaceIface.
func DBusInterfaceIfaceNewFromC(native unsafe.Pointer) *DBusInterfaceIface {
	return &DBusInterfaceIface{native: native}
}

// DBusInterfaceSkeletonClass is a representation of the C record GDBusInterfaceSkeletonClass.
//
// since 2.30
type DBusInterfaceSkeletonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeletonClass that represents the DBusInterfaceSkeletonClass.
func (recv *DBusInterfaceSkeletonClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonClassNewFromC creates a new DBusInterfaceSkeletonClass from a pointer to the C GDBusInterfaceSkeletonClass that represents the DBusInterfaceSkeletonClass.
func DBusInterfaceSkeletonClassNewFromC(native unsafe.Pointer) *DBusInterfaceSkeletonClass {
	return &DBusInterfaceSkeletonClass{native: native}
}

// DBusObjectIface is a representation of the C record GDBusObjectIface.
//
// since 2.30
type DBusObjectIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectIface that represents the DBusObjectIface.
func (recv *DBusObjectIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectIfaceNewFromC creates a new DBusObjectIface from a pointer to the C GDBusObjectIface that represents the DBusObjectIface.
func DBusObjectIfaceNewFromC(native unsafe.Pointer) *DBusObjectIface {
	return &DBusObjectIface{native: native}
}

// DBusObjectManagerClientClass is a representation of the C record GDBusObjectManagerClientClass.
//
// since 2.30
type DBusObjectManagerClientClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClientClass that represents the DBusObjectManagerClientClass.
func (recv *DBusObjectManagerClientClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientClassNewFromC creates a new DBusObjectManagerClientClass from a pointer to the C GDBusObjectManagerClientClass that represents the DBusObjectManagerClientClass.
func DBusObjectManagerClientClassNewFromC(native unsafe.Pointer) *DBusObjectManagerClientClass {
	return &DBusObjectManagerClientClass{native: native}
}

// DBusObjectManagerIface is a representation of the C record GDBusObjectManagerIface.
//
// since 2.30
type DBusObjectManagerIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerIface that represents the DBusObjectManagerIface.
func (recv *DBusObjectManagerIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerIfaceNewFromC creates a new DBusObjectManagerIface from a pointer to the C GDBusObjectManagerIface that represents the DBusObjectManagerIface.
func DBusObjectManagerIfaceNewFromC(native unsafe.Pointer) *DBusObjectManagerIface {
	return &DBusObjectManagerIface{native: native}
}

// DBusObjectManagerServerClass is a representation of the C record GDBusObjectManagerServerClass.
//
// since 2.30
type DBusObjectManagerServerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServerClass that represents the DBusObjectManagerServerClass.
func (recv *DBusObjectManagerServerClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerClassNewFromC creates a new DBusObjectManagerServerClass from a pointer to the C GDBusObjectManagerServerClass that represents the DBusObjectManagerServerClass.
func DBusObjectManagerServerClassNewFromC(native unsafe.Pointer) *DBusObjectManagerServerClass {
	return &DBusObjectManagerServerClass{native: native}
}

// DBusObjectProxyClass is a representation of the C record GDBusObjectProxyClass.
//
// since 2.30
type DBusObjectProxyClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxyClass that represents the DBusObjectProxyClass.
func (recv *DBusObjectProxyClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyClassNewFromC creates a new DBusObjectProxyClass from a pointer to the C GDBusObjectProxyClass that represents the DBusObjectProxyClass.
func DBusObjectProxyClassNewFromC(native unsafe.Pointer) *DBusObjectProxyClass {
	return &DBusObjectProxyClass{native: native}
}

// DBusObjectSkeletonClass is a representation of the C record GDBusObjectSkeletonClass.
//
// since 2.30
type DBusObjectSkeletonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeletonClass that represents the DBusObjectSkeletonClass.
func (recv *DBusObjectSkeletonClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonClassNewFromC creates a new DBusObjectSkeletonClass from a pointer to the C GDBusObjectSkeletonClass that represents the DBusObjectSkeletonClass.
func DBusObjectSkeletonClassNewFromC(native unsafe.Pointer) *DBusObjectSkeletonClass {
	return &DBusObjectSkeletonClass{native: native}
}

// IOModuleScope is a representation of the C record GIOModuleScope.
//
// since 2.30
type IOModuleScope struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModuleScope that represents the IOModuleScope.
func (recv *IOModuleScope) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleScopeNewFromC creates a new IOModuleScope from a pointer to the C GIOModuleScope that represents the IOModuleScope.
func IOModuleScopeNewFromC(native unsafe.Pointer) *IOModuleScope {
	return &IOModuleScope{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// TlsDatabaseClass is a representation of the C record GTlsDatabaseClass.
//
// since 2.30
type TlsDatabaseClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabaseClass that represents the TlsDatabaseClass.
func (recv *TlsDatabaseClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabaseClassNewFromC creates a new TlsDatabaseClass from a pointer to the C GTlsDatabaseClass that represents the TlsDatabaseClass.
func TlsDatabaseClassNewFromC(native unsafe.Pointer) *TlsDatabaseClass {
	return &TlsDatabaseClass{native: native}
}

// TlsInteractionClass is a representation of the C record GTlsInteractionClass.
//
// since 2.30
type TlsInteractionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteractionClass that represents the TlsInteractionClass.
func (recv *TlsInteractionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionClassNewFromC creates a new TlsInteractionClass from a pointer to the C GTlsInteractionClass that represents the TlsInteractionClass.
func TlsInteractionClassNewFromC(native unsafe.Pointer) *TlsInteractionClass {
	return &TlsInteractionClass{native: native}
}

// DBusInterfaceSkeleton is a representation of the C record GDBusInterfaceSkeleton.
//
// since 2.30
type DBusInterfaceSkeleton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeleton that represents the DBusInterfaceSkeleton.
func (recv *DBusInterfaceSkeleton) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonNewFromC creates a new DBusInterfaceSkeleton from a pointer to the C GDBusInterfaceSkeleton that represents the DBusInterfaceSkeleton.
func DBusInterfaceSkeletonNewFromC(native unsafe.Pointer) *DBusInterfaceSkeleton {
	return &DBusInterfaceSkeleton{native: native}
}

// DBusObjectManagerClient is a representation of the C record GDBusObjectManagerClient.
//
// since 2.30
type DBusObjectManagerClient struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClient that represents the DBusObjectManagerClient.
func (recv *DBusObjectManagerClient) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientNewFromC creates a new DBusObjectManagerClient from a pointer to the C GDBusObjectManagerClient that represents the DBusObjectManagerClient.
func DBusObjectManagerClientNewFromC(native unsafe.Pointer) *DBusObjectManagerClient {
	return &DBusObjectManagerClient{native: native}
}

// DBusObjectManagerServer is a representation of the C record GDBusObjectManagerServer.
//
// since 2.30
type DBusObjectManagerServer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServer that represents the DBusObjectManagerServer.
func (recv *DBusObjectManagerServer) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerNewFromC creates a new DBusObjectManagerServer from a pointer to the C GDBusObjectManagerServer that represents the DBusObjectManagerServer.
func DBusObjectManagerServerNewFromC(native unsafe.Pointer) *DBusObjectManagerServer {
	return &DBusObjectManagerServer{native: native}
}

// DBusObjectProxy is a representation of the C record GDBusObjectProxy.
//
// since 2.30
type DBusObjectProxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxy that represents the DBusObjectProxy.
func (recv *DBusObjectProxy) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyNewFromC creates a new DBusObjectProxy from a pointer to the C GDBusObjectProxy that represents the DBusObjectProxy.
func DBusObjectProxyNewFromC(native unsafe.Pointer) *DBusObjectProxy {
	return &DBusObjectProxy{native: native}
}

// DBusObjectSkeleton is a representation of the C record GDBusObjectSkeleton.
//
// since 2.30
type DBusObjectSkeleton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeleton that represents the DBusObjectSkeleton.
func (recv *DBusObjectSkeleton) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonNewFromC creates a new DBusObjectSkeleton from a pointer to the C GDBusObjectSkeleton that represents the DBusObjectSkeleton.
func DBusObjectSkeletonNewFromC(native unsafe.Pointer) *DBusObjectSkeleton {
	return &DBusObjectSkeleton{native: native}
}

// TlsDatabase is a representation of the C record GTlsDatabase.
//
// since 2.30
type TlsDatabase struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabase that represents the TlsDatabase.
func (recv *TlsDatabase) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabaseNewFromC creates a new TlsDatabase from a pointer to the C GTlsDatabase that represents the TlsDatabase.
func TlsDatabaseNewFromC(native unsafe.Pointer) *TlsDatabase {
	return &TlsDatabase{native: native}
}

// TlsInteraction is a representation of the C record GTlsInteraction.
//
// since 2.30
type TlsInteraction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteraction that represents the TlsInteraction.
func (recv *TlsInteraction) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionNewFromC creates a new TlsInteraction from a pointer to the C GTlsInteraction that represents the TlsInteraction.
func TlsInteractionNewFromC(native unsafe.Pointer) *TlsInteraction {
	return &TlsInteraction{native: native}
}

// TlsPassword is a representation of the C record GTlsPassword.
//
// since 2.30
type TlsPassword struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPassword that represents the TlsPassword.
func (recv *TlsPassword) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordNewFromC creates a new TlsPassword from a pointer to the C GTlsPassword that represents the TlsPassword.
func TlsPasswordNewFromC(native unsafe.Pointer) *TlsPassword {
	return &TlsPassword{native: native}
}

// DBusInterface is a representation of the C interface GDBusInterface.
//
// since 2.30
type DBusInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterface that represents the DBusInterface.
func (recv *DBusInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceNewFromC creates a new DBusInterface from a pointer to the C GDBusInterface that represents the DBusInterface.
func DBusInterfaceNewFromC(native unsafe.Pointer) *DBusInterface {
	return &DBusInterface{native: native}
}

// TlsFileDatabase is a representation of the C interface GTlsFileDatabase.
//
// since 2.30
type TlsFileDatabase struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsFileDatabase that represents the TlsFileDatabase.
func (recv *TlsFileDatabase) ToC() unsafe.Pointer {
	return recv.native
}

// TlsFileDatabaseNewFromC creates a new TlsFileDatabase from a pointer to the C GTlsFileDatabase that represents the TlsFileDatabase.
func TlsFileDatabaseNewFromC(native unsafe.Pointer) *TlsFileDatabase {
	return &TlsFileDatabase{native: native}
}
