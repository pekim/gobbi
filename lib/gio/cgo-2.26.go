// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native *C.GCredentials
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	c := (*C.GCredentials)(u)
	if c == nil {
		return nil
	}

	g := &Credentials{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Credentials) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native *C.GDBusAuthObserver
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	c := (*C.GDBusAuthObserver)(u)
	if c == nil {
		return nil
	}

	g := &DBusAuthObserver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusAuthObserver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native *C.GDBusConnection
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	c := (*C.GDBusConnection)(u)
	if c == nil {
		return nil
	}

	g := &DBusConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native *C.GDBusMessage
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	c := (*C.GDBusMessage)(u)
	if c == nil {
		return nil
	}

	g := &DBusMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native *C.GDBusMethodInvocation
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	c := (*C.GDBusMethodInvocation)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMethodInvocation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native *C.GDBusProxy
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	c := (*C.GDBusProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxy{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusProxy) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native *C.GDBusServer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	c := (*C.GDBusServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusServer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusServer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	c := (*C.GProxyAddress)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	c := (*C.GUnixCredentialsMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixCredentialsMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_bus_get_finish : return type :

// Unsupported : g_bus_get_sync : return type :

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback (GBusAcquiredCallback) for param bus_acquired_handler

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback (GBusNameAcquiredCallback) for param name_acquired_handler

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_dbus_address_get_for_bus_sync : return type :

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dbus_address_get_stream_finish : return type :

// Unsupported : g_dbus_address_get_stream_sync : return type :

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations :

// Unsupported : g_dbus_error_encode_gerror : return type :

// Unsupported : g_dbus_error_get_remote_error : return type :

// Unsupported : g_dbus_error_is_remote_error : return type :

// Unsupported : g_dbus_error_new_for_dbus_error : return type :

// Unsupported : g_dbus_error_register_error : return type :

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries :

// Unsupported : g_dbus_error_strip_remote_error : return type :

// Unsupported : g_dbus_error_unregister_error : return type :

// Unsupported : g_dbus_generate_guid : return type :

// Unsupported : g_dbus_is_address : return type :

// Unsupported : g_dbus_is_guid : return type :

// Unsupported : g_dbus_is_interface_name : return type :

// Unsupported : g_dbus_is_member_name : return type :

// Unsupported : g_dbus_is_name : return type :

// Unsupported : g_dbus_is_supported_address : return type :

// Unsupported : g_dbus_is_unique_name : return type :

// Unsupported : g_proxy_get_default_for_protocol : return type :

// Unsupported : g_proxy_resolver_get_default : return type :

// Proxy is a wrapper around the C record GProxy.
type Proxy struct {
	native *C.GProxy
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	c := (*C.GProxy)(u)
	if c == nil {
		return nil
	}

	g := &Proxy{native: c}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyResolver is a wrapper around the C record GProxyResolver.
type ProxyResolver struct {
	native *C.GProxyResolver
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	c := (*C.GProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolver{native: c}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CredentialsClass is a wrapper around the C record GCredentialsClass.
type CredentialsClass struct {
	native *C.GCredentialsClass
}

func CredentialsClassNewFromC(u unsafe.Pointer) *CredentialsClass {
	c := (*C.GCredentialsClass)(u)
	if c == nil {
		return nil
	}

	g := &CredentialsClass{native: c}

	return g
}

func (recv *CredentialsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAnnotationInfo is a wrapper around the C record GDBusAnnotationInfo.
type DBusAnnotationInfo struct {
	native   *C.GDBusAnnotationInfo
	RefCount int32
	Key      string
	Value    string
	// no type for annotations
}

func DBusAnnotationInfoNewFromC(u unsafe.Pointer) *DBusAnnotationInfo {
	c := (*C.GDBusAnnotationInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusAnnotationInfo{
		Key:      C.GoString(c.key),
		RefCount: (int32)(c.ref_count),
		Value:    C.GoString(c.value),
		native:   c,
	}

	return g
}

func (recv *DBusAnnotationInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusArgInfo is a wrapper around the C record GDBusArgInfo.
type DBusArgInfo struct {
	native    *C.GDBusArgInfo
	RefCount  int32
	Name      string
	Signature string
	// no type for annotations
}

func DBusArgInfoNewFromC(u unsafe.Pointer) *DBusArgInfo {
	c := (*C.GDBusArgInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusArgInfo{
		Name:      C.GoString(c.name),
		RefCount:  (int32)(c.ref_count),
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusArgInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusErrorEntry is a wrapper around the C record GDBusErrorEntry.
type DBusErrorEntry struct {
	native        *C.GDBusErrorEntry
	ErrorCode     int32
	DbusErrorName string
}

func DBusErrorEntryNewFromC(u unsafe.Pointer) *DBusErrorEntry {
	c := (*C.GDBusErrorEntry)(u)
	if c == nil {
		return nil
	}

	g := &DBusErrorEntry{
		DbusErrorName: C.GoString(c.dbus_error_name),
		ErrorCode:     (int32)(c.error_code),
		native:        c,
	}

	return g
}

func (recv *DBusErrorEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusInterfaceInfo is a wrapper around the C record GDBusInterfaceInfo.
type DBusInterfaceInfo struct {
	native   *C.GDBusInterfaceInfo
	RefCount int32
	Name     string
	// no type for methods
	// no type for signals
	// no type for properties
	// no type for annotations
}

func DBusInterfaceInfoNewFromC(u unsafe.Pointer) *DBusInterfaceInfo {
	c := (*C.GDBusInterfaceInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusInterfaceInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusInterfaceVTable is a wrapper around the C record GDBusInterfaceVTable.
type DBusInterfaceVTable struct {
	native *C.GDBusInterfaceVTable
	// method_call : no type generator for DBusInterfaceMethodCallFunc, GDBusInterfaceMethodCallFunc
	// get_property : no type generator for DBusInterfaceGetPropertyFunc, GDBusInterfaceGetPropertyFunc
	// set_property : no type generator for DBusInterfaceSetPropertyFunc, GDBusInterfaceSetPropertyFunc
	// Private : padding
}

func DBusInterfaceVTableNewFromC(u unsafe.Pointer) *DBusInterfaceVTable {
	c := (*C.GDBusInterfaceVTable)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceVTable{native: c}

	return g
}

func (recv *DBusInterfaceVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMethodInfo is a wrapper around the C record GDBusMethodInfo.
type DBusMethodInfo struct {
	native   *C.GDBusMethodInfo
	RefCount int32
	Name     string
	// no type for in_args
	// no type for out_args
	// no type for annotations
}

func DBusMethodInfoNewFromC(u unsafe.Pointer) *DBusMethodInfo {
	c := (*C.GDBusMethodInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusMethodInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusNodeInfo is a wrapper around the C record GDBusNodeInfo.
type DBusNodeInfo struct {
	native   *C.GDBusNodeInfo
	RefCount int32
	Path     string
	// no type for interfaces
	// no type for nodes
	// no type for annotations
}

func DBusNodeInfoNewFromC(u unsafe.Pointer) *DBusNodeInfo {
	c := (*C.GDBusNodeInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusNodeInfo{
		Path:     C.GoString(c.path),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusNodeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusPropertyInfo is a wrapper around the C record GDBusPropertyInfo.
type DBusPropertyInfo struct {
	native    *C.GDBusPropertyInfo
	RefCount  int32
	Name      string
	Signature string
	Flags     DBusPropertyInfoFlags
	// no type for annotations
}

func DBusPropertyInfoNewFromC(u unsafe.Pointer) *DBusPropertyInfo {
	c := (*C.GDBusPropertyInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusPropertyInfo{
		Flags:     (DBusPropertyInfoFlags)(c.flags),
		Name:      C.GoString(c.name),
		RefCount:  (int32)(c.ref_count),
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusPropertyInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxyClass is a wrapper around the C record GDBusProxyClass.
type DBusProxyClass struct {
	native *C.GDBusProxyClass
	// Private : parent_class
	// no type for g_properties_changed
	// no type for g_signal
	// Private : padding
}

func DBusProxyClassNewFromC(u unsafe.Pointer) *DBusProxyClass {
	c := (*C.GDBusProxyClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxyClass{native: c}

	return g
}

func (recv *DBusProxyClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusSignalInfo is a wrapper around the C record GDBusSignalInfo.
type DBusSignalInfo struct {
	native   *C.GDBusSignalInfo
	RefCount int32
	Name     string
	// no type for args
	// no type for annotations
}

func DBusSignalInfoNewFromC(u unsafe.Pointer) *DBusSignalInfo {
	c := (*C.GDBusSignalInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusSignalInfo{
		Name:     C.GoString(c.name),
		RefCount: (int32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *DBusSignalInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusSubtreeVTable is a wrapper around the C record GDBusSubtreeVTable.
type DBusSubtreeVTable struct {
	native *C.GDBusSubtreeVTable
	// enumerate : no type generator for DBusSubtreeEnumerateFunc, GDBusSubtreeEnumerateFunc
	// introspect : no type generator for DBusSubtreeIntrospectFunc, GDBusSubtreeIntrospectFunc
	// dispatch : no type generator for DBusSubtreeDispatchFunc, GDBusSubtreeDispatchFunc
	// Private : padding
}

func DBusSubtreeVTableNewFromC(u unsafe.Pointer) *DBusSubtreeVTable {
	c := (*C.GDBusSubtreeVTable)(u)
	if c == nil {
		return nil
	}

	g := &DBusSubtreeVTable{native: c}

	return g
}

func (recv *DBusSubtreeVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressClass is a wrapper around the C record GProxyAddressClass.
type ProxyAddressClass struct {
	native *C.GProxyAddressClass
	// parent_class : record
}

func ProxyAddressClassNewFromC(u unsafe.Pointer) *ProxyAddressClass {
	c := (*C.GProxyAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressClass{native: c}

	return g
}

func (recv *ProxyAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyInterface is a wrapper around the C record GProxyInterface.
type ProxyInterface struct {
	native *C.GProxyInterface
	// g_iface : record
	// no type for connect
	// no type for connect_async
	// no type for connect_finish
	// no type for supports_hostname
}

func ProxyInterfaceNewFromC(u unsafe.Pointer) *ProxyInterface {
	c := (*C.GProxyInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyInterface{native: c}

	return g
}

func (recv *ProxyInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsClientConnectionInterface is a wrapper around the C record GTlsClientConnectionInterface.
type TlsClientConnectionInterface struct {
	native *C.GTlsClientConnectionInterface
	// g_iface : record
	// no type for copy_session_state
}

func TlsClientConnectionInterfaceNewFromC(u unsafe.Pointer) *TlsClientConnectionInterface {
	c := (*C.GTlsClientConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnectionInterface{native: c}

	return g
}

func (recv *TlsClientConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsServerConnectionInterface is a wrapper around the C record GTlsServerConnectionInterface.
type TlsServerConnectionInterface struct {
	native *C.GTlsServerConnectionInterface
	// g_iface : record
}

func TlsServerConnectionInterfaceNewFromC(u unsafe.Pointer) *TlsServerConnectionInterface {
	c := (*C.GTlsServerConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnectionInterface{native: c}

	return g
}

func (recv *TlsServerConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessageClass is a wrapper around the C record GUnixCredentialsMessageClass.
type UnixCredentialsMessageClass struct {
	native *C.GUnixCredentialsMessageClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func UnixCredentialsMessageClassNewFromC(u unsafe.Pointer) *UnixCredentialsMessageClass {
	c := (*C.GUnixCredentialsMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessageClass{native: c}

	return g
}

func (recv *UnixCredentialsMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
