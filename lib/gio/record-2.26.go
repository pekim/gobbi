// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Equals compares this CredentialsClass with another CredentialsClass, and returns true if they represent the same GObject.
func (recv *CredentialsClass) Equals(other *CredentialsClass) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusAnnotationInfo with another DBusAnnotationInfo, and returns true if they represent the same GObject.
func (recv *DBusAnnotationInfo) Equals(other *DBusAnnotationInfo) bool {
	return other.ToC() == recv.ToC()
}

// g_dbus_annotation_info_lookup : unsupported parameter annotations :
// Blacklisted : g_dbus_annotation_info_ref

// Blacklisted : g_dbus_annotation_info_unref

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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusArgInfo with another DBusArgInfo, and returns true if they represent the same GObject.
func (recv *DBusArgInfo) Equals(other *DBusArgInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_arg_info_ref

// Blacklisted : g_dbus_arg_info_unref

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
	recv.native.error_code =
		(C.gint)(recv.ErrorCode)
	recv.native.dbus_error_name =
		C.CString(recv.DbusErrorName)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusErrorEntry with another DBusErrorEntry, and returns true if they represent the same GObject.
func (recv *DBusErrorEntry) Equals(other *DBusErrorEntry) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusInterfaceInfo with another DBusInterfaceInfo, and returns true if they represent the same GObject.
func (recv *DBusInterfaceInfo) Equals(other *DBusInterfaceInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_interface_info_generate_xml

// Blacklisted : g_dbus_interface_info_lookup_method

// Blacklisted : g_dbus_interface_info_lookup_property

// Blacklisted : g_dbus_interface_info_lookup_signal

// Blacklisted : g_dbus_interface_info_ref

// Blacklisted : g_dbus_interface_info_unref

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

// Equals compares this DBusInterfaceVTable with another DBusInterfaceVTable, and returns true if they represent the same GObject.
func (recv *DBusInterfaceVTable) Equals(other *DBusInterfaceVTable) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusMethodInfo with another DBusMethodInfo, and returns true if they represent the same GObject.
func (recv *DBusMethodInfo) Equals(other *DBusMethodInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_method_info_ref

// Blacklisted : g_dbus_method_info_unref

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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.path =
		C.CString(recv.Path)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusNodeInfo with another DBusNodeInfo, and returns true if they represent the same GObject.
func (recv *DBusNodeInfo) Equals(other *DBusNodeInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_node_info_new_for_xml

// Blacklisted : g_dbus_node_info_generate_xml

// Blacklisted : g_dbus_node_info_lookup_interface

// Blacklisted : g_dbus_node_info_ref

// Blacklisted : g_dbus_node_info_unref

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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)
	recv.native.flags =
		(C.GDBusPropertyInfoFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusPropertyInfo with another DBusPropertyInfo, and returns true if they represent the same GObject.
func (recv *DBusPropertyInfo) Equals(other *DBusPropertyInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_property_info_ref

// Blacklisted : g_dbus_property_info_unref

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

// Equals compares this DBusProxyClass with another DBusProxyClass, and returns true if they represent the same GObject.
func (recv *DBusProxyClass) Equals(other *DBusProxyClass) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.ref_count =
		(C.gint)(recv.RefCount)
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusSignalInfo with another DBusSignalInfo, and returns true if they represent the same GObject.
func (recv *DBusSignalInfo) Equals(other *DBusSignalInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dbus_signal_info_ref

// Blacklisted : g_dbus_signal_info_unref

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

// Equals compares this DBusSubtreeVTable with another DBusSubtreeVTable, and returns true if they represent the same GObject.
func (recv *DBusSubtreeVTable) Equals(other *DBusSubtreeVTable) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this ProxyAddressClass with another ProxyAddressClass, and returns true if they represent the same GObject.
func (recv *ProxyAddressClass) Equals(other *ProxyAddressClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this ProxyInterface with another ProxyInterface, and returns true if they represent the same GObject.
func (recv *ProxyInterface) Equals(other *ProxyInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TlsClientConnectionInterface with another TlsClientConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsClientConnectionInterface) Equals(other *TlsClientConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TlsServerConnectionInterface with another TlsServerConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsServerConnectionInterface) Equals(other *TlsServerConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this UnixCredentialsMessageClass with another UnixCredentialsMessageClass, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessageClass) Equals(other *UnixCredentialsMessageClass) bool {
	return other.ToC() == recv.ToC()
}
