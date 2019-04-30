// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// DBusInterfaceSkeleton is a wrapper around the C record GDBusInterfaceSkeleton.
type DBusInterfaceSkeleton struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusInterfaceSkeletonNewFromC(u unsafe.Pointer) *DBusInterfaceSkeleton {
	if u == nil {
		return nil
	}

	g := &DBusInterfaceSkeleton{native: u}

	return g
}

func (recv *DBusInterfaceSkeleton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectManagerClient is a wrapper around the C record GDBusObjectManagerClient.
type DBusObjectManagerClient struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerClientNewFromC(u unsafe.Pointer) *DBusObjectManagerClient {
	if u == nil {
		return nil
	}

	g := &DBusObjectManagerClient{native: u}

	return g
}

func (recv *DBusObjectManagerClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_object_manager_client_new_finish : return type :

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : return type :

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// DBusObjectManagerServer is a wrapper around the C record GDBusObjectManagerServer.
type DBusObjectManagerServer struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectManagerServerNewFromC(u unsafe.Pointer) *DBusObjectManagerServer {
	if u == nil {
		return nil
	}

	g := &DBusObjectManagerServer{native: u}

	return g
}

func (recv *DBusObjectManagerServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_object_manager_server_new : return type :

// DBusObjectProxy is a wrapper around the C record GDBusObjectProxy.
type DBusObjectProxy struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectProxyNewFromC(u unsafe.Pointer) *DBusObjectProxy {
	if u == nil {
		return nil
	}

	g := &DBusObjectProxy{native: u}

	return g
}

func (recv *DBusObjectProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_object_proxy_new : return type :

// DBusObjectSkeleton is a wrapper around the C record GDBusObjectSkeleton.
type DBusObjectSkeleton struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func DBusObjectSkeletonNewFromC(u unsafe.Pointer) *DBusObjectSkeleton {
	if u == nil {
		return nil
	}

	g := &DBusObjectSkeleton{native: u}

	return g
}

func (recv *DBusObjectSkeleton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_object_skeleton_new : return type :

// TlsDatabase is a wrapper around the C record GTlsDatabase.
type TlsDatabase struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TlsDatabaseNewFromC(u unsafe.Pointer) *TlsDatabase {
	if u == nil {
		return nil
	}

	g := &TlsDatabase{native: u}

	return g
}

func (recv *TlsDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsInteraction is a wrapper around the C record GTlsInteraction.
type TlsInteraction struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func TlsInteractionNewFromC(u unsafe.Pointer) *TlsInteraction {
	if u == nil {
		return nil
	}

	g := &TlsInteraction{native: u}

	return g
}

func (recv *TlsInteraction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsPassword is a wrapper around the C record GTlsPassword.
type TlsPassword struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TlsPasswordNewFromC(u unsafe.Pointer) *TlsPassword {
	if u == nil {
		return nil
	}

	g := &TlsPassword{native: u}

	return g
}

func (recv *TlsPassword) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_tls_password_new : return type :
