// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// Equals compares this CredentialsClass with another CredentialsClass, and returns true if they represent the same GObject.
func (recv *CredentialsClass) Equals(other *CredentialsClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusAnnotationInfo with another DBusAnnotationInfo, and returns true if they represent the same GObject.
func (recv *DBusAnnotationInfo) Equals(other *DBusAnnotationInfo) bool {
	return other.ToC() == recv.ToC()
}

// g_dbus_annotation_info_lookup : unsupported parameter annotations :
// Ref is a wrapper around the C function g_dbus_annotation_info_ref.
func (recv *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	retC := C.g_dbus_annotation_info_ref((*C.GDBusAnnotationInfo)(recv.native))
	retGo := DBusAnnotationInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_annotation_info_unref.
func (recv *DBusAnnotationInfo) Unref() {
	C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(recv.native))

	return
}

// Equals compares this DBusArgInfo with another DBusArgInfo, and returns true if they represent the same GObject.
func (recv *DBusArgInfo) Equals(other *DBusArgInfo) bool {
	return other.ToC() == recv.ToC()
}

// Ref is a wrapper around the C function g_dbus_arg_info_ref.
func (recv *DBusArgInfo) Ref() *DBusArgInfo {
	retC := C.g_dbus_arg_info_ref((*C.GDBusArgInfo)(recv.native))
	retGo := DBusArgInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_arg_info_unref.
func (recv *DBusArgInfo) Unref() {
	C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(recv.native))

	return
}

// Equals compares this DBusErrorEntry with another DBusErrorEntry, and returns true if they represent the same GObject.
func (recv *DBusErrorEntry) Equals(other *DBusErrorEntry) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusInterfaceInfo with another DBusInterfaceInfo, and returns true if they represent the same GObject.
func (recv *DBusInterfaceInfo) Equals(other *DBusInterfaceInfo) bool {
	return other.ToC() == recv.ToC()
}

// GenerateXml is a wrapper around the C function g_dbus_interface_info_generate_xml.
func (recv *DBusInterfaceInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(C.NULL)
	if stringBuilder != nil {
		c_string_builder = (*C.GString)(stringBuilder.ToC())
	}

	C.g_dbus_interface_info_generate_xml((*C.GDBusInterfaceInfo)(recv.native), c_indent, c_string_builder)

	return
}

// LookupMethod is a wrapper around the C function g_dbus_interface_info_lookup_method.
func (recv *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_method((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LookupProperty is a wrapper around the C function g_dbus_interface_info_lookup_property.
func (recv *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_property((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LookupSignal is a wrapper around the C function g_dbus_interface_info_lookup_signal.
func (recv *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_signal((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusSignalInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_dbus_interface_info_ref.
func (recv *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_info_ref((*C.GDBusInterfaceInfo)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_interface_info_unref.
func (recv *DBusInterfaceInfo) Unref() {
	C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// Equals compares this DBusInterfaceVTable with another DBusInterfaceVTable, and returns true if they represent the same GObject.
func (recv *DBusInterfaceVTable) Equals(other *DBusInterfaceVTable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusMethodInfo with another DBusMethodInfo, and returns true if they represent the same GObject.
func (recv *DBusMethodInfo) Equals(other *DBusMethodInfo) bool {
	return other.ToC() == recv.ToC()
}

// Ref is a wrapper around the C function g_dbus_method_info_ref.
func (recv *DBusMethodInfo) Ref() *DBusMethodInfo {
	retC := C.g_dbus_method_info_ref((*C.GDBusMethodInfo)(recv.native))
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_method_info_unref.
func (recv *DBusMethodInfo) Unref() {
	C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(recv.native))

	return
}

// Equals compares this DBusNodeInfo with another DBusNodeInfo, and returns true if they represent the same GObject.
func (recv *DBusNodeInfo) Equals(other *DBusNodeInfo) bool {
	return other.ToC() == recv.ToC()
}

// GenerateXml is a wrapper around the C function g_dbus_node_info_generate_xml.
func (recv *DBusNodeInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(C.NULL)
	if stringBuilder != nil {
		c_string_builder = (*C.GString)(stringBuilder.ToC())
	}

	C.g_dbus_node_info_generate_xml((*C.GDBusNodeInfo)(recv.native), c_indent, c_string_builder)

	return
}

// LookupInterface is a wrapper around the C function g_dbus_node_info_lookup_interface.
func (recv *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_node_info_lookup_interface((*C.GDBusNodeInfo)(recv.native), c_name)
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_dbus_node_info_ref.
func (recv *DBusNodeInfo) Ref() *DBusNodeInfo {
	retC := C.g_dbus_node_info_ref((*C.GDBusNodeInfo)(recv.native))
	retGo := DBusNodeInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_node_info_unref.
func (recv *DBusNodeInfo) Unref() {
	C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(recv.native))

	return
}

// Equals compares this DBusPropertyInfo with another DBusPropertyInfo, and returns true if they represent the same GObject.
func (recv *DBusPropertyInfo) Equals(other *DBusPropertyInfo) bool {
	return other.ToC() == recv.ToC()
}

// Ref is a wrapper around the C function g_dbus_property_info_ref.
func (recv *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	retC := C.g_dbus_property_info_ref((*C.GDBusPropertyInfo)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_property_info_unref.
func (recv *DBusPropertyInfo) Unref() {
	C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(recv.native))

	return
}

// Equals compares this DBusProxyClass with another DBusProxyClass, and returns true if they represent the same GObject.
func (recv *DBusProxyClass) Equals(other *DBusProxyClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusSignalInfo with another DBusSignalInfo, and returns true if they represent the same GObject.
func (recv *DBusSignalInfo) Equals(other *DBusSignalInfo) bool {
	return other.ToC() == recv.ToC()
}

// Ref is a wrapper around the C function g_dbus_signal_info_ref.
func (recv *DBusSignalInfo) Ref() *DBusSignalInfo {
	retC := C.g_dbus_signal_info_ref((*C.GDBusSignalInfo)(recv.native))
	retGo := DBusSignalInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_dbus_signal_info_unref.
func (recv *DBusSignalInfo) Unref() {
	C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(recv.native))

	return
}

// Equals compares this DBusSubtreeVTable with another DBusSubtreeVTable, and returns true if they represent the same GObject.
func (recv *DBusSubtreeVTable) Equals(other *DBusSubtreeVTable) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ProxyAddressClass with another ProxyAddressClass, and returns true if they represent the same GObject.
func (recv *ProxyAddressClass) Equals(other *ProxyAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ProxyInterface with another ProxyInterface, and returns true if they represent the same GObject.
func (recv *ProxyInterface) Equals(other *ProxyInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TlsClientConnectionInterface with another TlsClientConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsClientConnectionInterface) Equals(other *TlsClientConnectionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TlsServerConnectionInterface with another TlsServerConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsServerConnectionInterface) Equals(other *TlsServerConnectionInterface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this UnixCredentialsMessageClass with another UnixCredentialsMessageClass, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessageClass) Equals(other *UnixCredentialsMessageClass) bool {
	return other.ToC() == recv.ToC()
}
