// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// DBusInterfaceSkeleton is a wrapper around the C record GDBusInterfaceSkeleton.
type DBusInterfaceSkeleton struct {
	native *C.GDBusInterfaceSkeleton
	// Private : parent_instance
	// Private : priv
}

func DBusInterfaceSkeletonNewFromC(u unsafe.Pointer) *DBusInterfaceSkeleton {
	c := (*C.GDBusInterfaceSkeleton)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeleton{native: c}

	return g
}

func (recv *DBusInterfaceSkeleton) toC() *C.GDBusInterfaceSkeleton {

	return recv.native
}

// Unsupported : g_dbus_interface_skeleton_export : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_dbus_interface_skeleton_flush : no return generator

// Unsupported : g_dbus_interface_skeleton_get_connection : no return generator

// GetFlags is a wrapper around the C function g_dbus_interface_skeleton_get_flags.
func (recv *DBusInterfaceSkeleton) GetFlags() DBusInterfaceSkeletonFlags {
	retC := C.g_dbus_interface_skeleton_get_flags((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := (DBusInterfaceSkeletonFlags)(retC)

	return retGo
}

// GetInfo is a wrapper around the C function g_dbus_interface_skeleton_get_info.
func (recv *DBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_skeleton_get_info((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_interface_skeleton_get_object_path.
func (recv *DBusInterfaceSkeleton) GetObjectPath() string {
	retC := C.g_dbus_interface_skeleton_get_object_path((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_interface_skeleton_get_properties : return type : Blacklisted record : GVariant

// GetVtable is a wrapper around the C function g_dbus_interface_skeleton_get_vtable.
func (recv *DBusInterfaceSkeleton) GetVtable() *DBusInterfaceVTable {
	retC := C.g_dbus_interface_skeleton_get_vtable((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := DBusInterfaceVTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_interface_skeleton_has_connection : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_dbus_interface_skeleton_set_flags : no return generator

// Unsupported : g_dbus_interface_skeleton_unexport : no return generator

// Unsupported : g_dbus_interface_skeleton_unexport_from_connection : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// DBusObjectManagerClient is a wrapper around the C record GDBusObjectManagerClient.
type DBusObjectManagerClient struct {
	native *C.GDBusObjectManagerClient
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerClientNewFromC(u unsafe.Pointer) *DBusObjectManagerClient {
	c := (*C.GDBusObjectManagerClient)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClient{native: c}

	return g
}

func (recv *DBusObjectManagerClient) toC() *C.GDBusObjectManagerClient {

	return recv.native
}

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_dbus_object_manager_client_get_connection : no return generator

// GetFlags is a wrapper around the C function g_dbus_object_manager_client_get_flags.
func (recv *DBusObjectManagerClient) GetFlags() DBusObjectManagerClientFlags {
	retC := C.g_dbus_object_manager_client_get_flags((*C.GDBusObjectManagerClient)(recv.native))
	retGo := (DBusObjectManagerClientFlags)(retC)

	return retGo
}

// GetName is a wrapper around the C function g_dbus_object_manager_client_get_name.
func (recv *DBusObjectManagerClient) GetName() string {
	retC := C.g_dbus_object_manager_client_get_name((*C.GDBusObjectManagerClient)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNameOwner is a wrapper around the C function g_dbus_object_manager_client_get_name_owner.
func (recv *DBusObjectManagerClient) GetNameOwner() string {
	retC := C.g_dbus_object_manager_client_get_name_owner((*C.GDBusObjectManagerClient)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusObjectManagerServer is a wrapper around the C record GDBusObjectManagerServer.
type DBusObjectManagerServer struct {
	native *C.GDBusObjectManagerServer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerServerNewFromC(u unsafe.Pointer) *DBusObjectManagerServer {
	c := (*C.GDBusObjectManagerServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServer{native: c}

	return g
}

func (recv *DBusObjectManagerServer) toC() *C.GDBusObjectManagerServer {

	return recv.native
}

// Unsupported : g_dbus_object_manager_server_new : no return generator

// Unsupported : g_dbus_object_manager_server_export : unsupported parameter object : no type generator for DBusObjectSkeleton, GDBusObjectSkeleton*

// Unsupported : g_dbus_object_manager_server_export_uniquely : unsupported parameter object : no type generator for DBusObjectSkeleton, GDBusObjectSkeleton*

// Unsupported : g_dbus_object_manager_server_get_connection : no return generator

// Unsupported : g_dbus_object_manager_server_is_exported : unsupported parameter object : no type generator for DBusObjectSkeleton, GDBusObjectSkeleton*

// Unsupported : g_dbus_object_manager_server_set_connection : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unexport is a wrapper around the C function g_dbus_object_manager_server_unexport.
func (recv *DBusObjectManagerServer) Unexport(objectPath string) bool {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	retC := C.g_dbus_object_manager_server_unexport((*C.GDBusObjectManagerServer)(recv.native), c_object_path)
	retGo := retC == C.TRUE

	return retGo
}

// DBusObjectProxy is a wrapper around the C record GDBusObjectProxy.
type DBusObjectProxy struct {
	native *C.GDBusObjectProxy
	// Private : parent_instance
	// Private : priv
}

func DBusObjectProxyNewFromC(u unsafe.Pointer) *DBusObjectProxy {
	c := (*C.GDBusObjectProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxy{native: c}

	return g
}

func (recv *DBusObjectProxy) toC() *C.GDBusObjectProxy {

	return recv.native
}

// Unsupported : g_dbus_object_proxy_new : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_dbus_object_proxy_get_connection : no return generator

// DBusObjectSkeleton is a wrapper around the C record GDBusObjectSkeleton.
type DBusObjectSkeleton struct {
	native *C.GDBusObjectSkeleton
	// Private : parent_instance
	// Private : priv
}

func DBusObjectSkeletonNewFromC(u unsafe.Pointer) *DBusObjectSkeleton {
	c := (*C.GDBusObjectSkeleton)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeleton{native: c}

	return g
}

func (recv *DBusObjectSkeleton) toC() *C.GDBusObjectSkeleton {

	return recv.native
}

// Unsupported : g_dbus_object_skeleton_new : no return generator

// Unsupported : g_dbus_object_skeleton_add_interface : unsupported parameter interface_ : no type generator for DBusInterfaceSkeleton, GDBusInterfaceSkeleton*

// Unsupported : g_dbus_object_skeleton_flush : no return generator

// Unsupported : g_dbus_object_skeleton_remove_interface : unsupported parameter interface_ : no type generator for DBusInterfaceSkeleton, GDBusInterfaceSkeleton*

// Unsupported : g_dbus_object_skeleton_remove_interface_by_name : no return generator

// Unsupported : g_dbus_object_skeleton_set_object_path : no return generator

// TlsDatabase is a wrapper around the C record GTlsDatabase.
type TlsDatabase struct {
	native *C.GTlsDatabase
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func TlsDatabaseNewFromC(u unsafe.Pointer) *TlsDatabase {
	c := (*C.GTlsDatabase)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabase{native: c}

	return g
}

func (recv *TlsDatabase) toC() *C.GTlsDatabase {

	return recv.native
}

// Unsupported : g_tls_database_create_certificate_handle : unsupported parameter certificate : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_database_lookup_certificate_for_handle : unsupported parameter interaction : no type generator for TlsInteraction, GTlsInteraction*

// Unsupported : g_tls_database_lookup_certificate_for_handle_async : unsupported parameter interaction : no type generator for TlsInteraction, GTlsInteraction*

// Unsupported : g_tls_database_lookup_certificate_for_handle_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_lookup_certificate_issuer : unsupported parameter certificate : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_database_lookup_certificate_issuer_async : unsupported parameter certificate : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_database_lookup_certificate_issuer_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_lookup_certificates_issued_by : unsupported parameter issuer_raw_dn : no param type

// Unsupported : g_tls_database_lookup_certificates_issued_by_async : unsupported parameter issuer_raw_dn : no param type

// Unsupported : g_tls_database_lookup_certificates_issued_by_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_database_verify_chain : unsupported parameter chain : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_database_verify_chain_async : unsupported parameter chain : no type generator for TlsCertificate, GTlsCertificate*

// Unsupported : g_tls_database_verify_chain_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// TlsInteraction is a wrapper around the C record GTlsInteraction.
type TlsInteraction struct {
	native *C.GTlsInteraction
	// Private : parent_instance
	// Private : priv
}

func TlsInteractionNewFromC(u unsafe.Pointer) *TlsInteraction {
	c := (*C.GTlsInteraction)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteraction{native: c}

	return g
}

func (recv *TlsInteraction) toC() *C.GTlsInteraction {

	return recv.native
}

// Unsupported : g_tls_interaction_ask_password : unsupported parameter password : no type generator for TlsPassword, GTlsPassword*

// Unsupported : g_tls_interaction_ask_password_async : unsupported parameter password : no type generator for TlsPassword, GTlsPassword*

// Unsupported : g_tls_interaction_ask_password_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_tls_interaction_invoke_ask_password : unsupported parameter password : no type generator for TlsPassword, GTlsPassword*

// Unsupported : g_tls_interaction_invoke_request_certificate : unsupported parameter connection : no type generator for TlsConnection, GTlsConnection*

// Unsupported : g_tls_interaction_request_certificate : unsupported parameter connection : no type generator for TlsConnection, GTlsConnection*

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter connection : no type generator for TlsConnection, GTlsConnection*

// Unsupported : g_tls_interaction_request_certificate_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// TlsPassword is a wrapper around the C record GTlsPassword.
type TlsPassword struct {
	native *C.GTlsPassword
	// parent_instance : no type generator for GObject.Object, GObject
	// priv : record
}

func TlsPasswordNewFromC(u unsafe.Pointer) *TlsPassword {
	c := (*C.GTlsPassword)(u)
	if c == nil {
		return nil
	}

	g := &TlsPassword{native: c}

	return g
}

func (recv *TlsPassword) toC() *C.GTlsPassword {

	return recv.native
}

// Unsupported : g_tls_password_new : no return generator

// GetDescription is a wrapper around the C function g_tls_password_get_description.
func (recv *TlsPassword) GetDescription() string {
	retC := C.g_tls_password_get_description((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_tls_password_get_flags.
func (recv *TlsPassword) GetFlags() TlsPasswordFlags {
	retC := C.g_tls_password_get_flags((*C.GTlsPassword)(recv.native))
	retGo := (TlsPasswordFlags)(retC)

	return retGo
}

// Unsupported : g_tls_password_get_value : unsupported parameter length : no type generator for gsize, gsize*

// GetWarning is a wrapper around the C function g_tls_password_get_warning.
func (recv *TlsPassword) GetWarning() string {
	retC := C.g_tls_password_get_warning((*C.GTlsPassword)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_tls_password_set_description : no return generator

// Unsupported : g_tls_password_set_flags : no return generator

// Unsupported : g_tls_password_set_value : unsupported parameter value : no param type

// Unsupported : g_tls_password_set_value_full : unsupported parameter value : no param type

// Unsupported : g_tls_password_set_warning : no return generator
