// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// CredentialsClass is a wrapper around the C record GCredentialsClass.
type CredentialsClass struct {
	native unsafe.Pointer
}

// DBusAnnotationInfo is a wrapper around the C record GDBusAnnotationInfo.
type DBusAnnotationInfo struct {
	native   unsafe.Pointer
	RefCount int32
	Key      string
	Value    string
	// no type for annotations
}

// DBusArgInfo is a wrapper around the C record GDBusArgInfo.
type DBusArgInfo struct {
	native    unsafe.Pointer
	RefCount  int32
	Name      string
	Signature string
	// no type for annotations
}

// DBusErrorEntry is a wrapper around the C record GDBusErrorEntry.
type DBusErrorEntry struct {
	native        unsafe.Pointer
	ErrorCode     int32
	DbusErrorName string
}

// DBusInterfaceInfo is a wrapper around the C record GDBusInterfaceInfo.
type DBusInterfaceInfo struct {
	native   unsafe.Pointer
	RefCount int32
	Name     string
	// no type for methods
	// no type for signals
	// no type for properties
	// no type for annotations
}

// DBusInterfaceVTable is a wrapper around the C record GDBusInterfaceVTable.
type DBusInterfaceVTable struct {
	native unsafe.Pointer
	// method_call : no type generator for DBusInterfaceMethodCallFunc, GDBusInterfaceMethodCallFunc
	// get_property : no type generator for DBusInterfaceGetPropertyFunc, GDBusInterfaceGetPropertyFunc
	// set_property : no type generator for DBusInterfaceSetPropertyFunc, GDBusInterfaceSetPropertyFunc
	// Private : padding
}

// DBusMethodInfo is a wrapper around the C record GDBusMethodInfo.
type DBusMethodInfo struct {
	native   unsafe.Pointer
	RefCount int32
	Name     string
	// no type for in_args
	// no type for out_args
	// no type for annotations
}

// DBusNodeInfo is a wrapper around the C record GDBusNodeInfo.
type DBusNodeInfo struct {
	native   unsafe.Pointer
	RefCount int32
	Path     string
	// no type for interfaces
	// no type for nodes
	// no type for annotations
}

// DBusPropertyInfo is a wrapper around the C record GDBusPropertyInfo.
type DBusPropertyInfo struct {
	native    unsafe.Pointer
	RefCount  int32
	Name      string
	Signature string
	Flags     DBusPropertyInfoFlags
	// no type for annotations
}

// DBusProxyClass is a wrapper around the C record GDBusProxyClass.
type DBusProxyClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for g_properties_changed
	// no type for g_signal
	// Private : padding
}

// DBusSignalInfo is a wrapper around the C record GDBusSignalInfo.
type DBusSignalInfo struct {
	native   unsafe.Pointer
	RefCount int32
	Name     string
	// no type for args
	// no type for annotations
}

// DBusSubtreeVTable is a wrapper around the C record GDBusSubtreeVTable.
type DBusSubtreeVTable struct {
	native unsafe.Pointer
	// enumerate : no type generator for DBusSubtreeEnumerateFunc, GDBusSubtreeEnumerateFunc
	// introspect : no type generator for DBusSubtreeIntrospectFunc, GDBusSubtreeIntrospectFunc
	// dispatch : no type generator for DBusSubtreeDispatchFunc, GDBusSubtreeDispatchFunc
	// Private : padding
}

// ProxyAddressClass is a wrapper around the C record GProxyAddressClass.
type ProxyAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ProxyInterface is a wrapper around the C record GProxyInterface.
type ProxyInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for connect
	// no type for connect_async
	// no type for connect_finish
	// no type for supports_hostname
}

// TlsClientConnectionInterface is a wrapper around the C record GTlsClientConnectionInterface.
type TlsClientConnectionInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for copy_session_state
}

// TlsServerConnectionInterface is a wrapper around the C record GTlsServerConnectionInterface.
type TlsServerConnectionInterface struct {
	native unsafe.Pointer
	// g_iface : record
}

// UnixCredentialsMessageClass is a wrapper around the C record GUnixCredentialsMessageClass.
type UnixCredentialsMessageClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}
