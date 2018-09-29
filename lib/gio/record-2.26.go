// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
	native *C.GDBusAnnotationInfo
	// ref_count : no type generator for gint, volatile gint
	Key   string
	Value string
	// no type for annotations
}

func DBusAnnotationInfoNewFromC(u unsafe.Pointer) *DBusAnnotationInfo {
	c := (*C.GDBusAnnotationInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusAnnotationInfo{
		Key:    C.GoString(c.key),
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *DBusAnnotationInfo) ToC() unsafe.Pointer {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

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

// DBusArgInfo is a wrapper around the C record GDBusArgInfo.
type DBusArgInfo struct {
	native *C.GDBusArgInfo
	// ref_count : no type generator for gint, volatile gint
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
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusArgInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)

	return (unsafe.Pointer)(recv.native)
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

// DBusInterfaceInfo is a wrapper around the C record GDBusInterfaceInfo.
type DBusInterfaceInfo struct {
	native *C.GDBusInterfaceInfo
	// ref_count : no type generator for gint, volatile gint
	Name string
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
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *DBusInterfaceInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
}

// GenerateXml is a wrapper around the C function g_dbus_interface_info_generate_xml.
func (recv *DBusInterfaceInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(stringBuilder.ToC())

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
	native *C.GDBusMethodInfo
	// ref_count : no type generator for gint, volatile gint
	Name string
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
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *DBusMethodInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
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

// DBusNodeInfo is a wrapper around the C record GDBusNodeInfo.
type DBusNodeInfo struct {
	native *C.GDBusNodeInfo
	// ref_count : no type generator for gint, volatile gint
	Path string
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
		Path:   C.GoString(c.path),
		native: c,
	}

	return g
}

func (recv *DBusNodeInfo) ToC() unsafe.Pointer {
	recv.native.path =
		C.CString(recv.Path)

	return (unsafe.Pointer)(recv.native)
}

// DBusNodeInfoNewForXml is a wrapper around the C function g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) (*DBusNodeInfo, error) {
	c_xml_data := C.CString(xmlData)
	defer C.free(unsafe.Pointer(c_xml_data))

	var cThrowableError *C.GError

	retC := C.g_dbus_node_info_new_for_xml(c_xml_data, &cThrowableError)
	retGo := DBusNodeInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GenerateXml is a wrapper around the C function g_dbus_node_info_generate_xml.
func (recv *DBusNodeInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(stringBuilder.ToC())

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

// DBusPropertyInfo is a wrapper around the C record GDBusPropertyInfo.
type DBusPropertyInfo struct {
	native *C.GDBusPropertyInfo
	// ref_count : no type generator for gint, volatile gint
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
		Signature: C.GoString(c.signature),
		native:    c,
	}

	return g
}

func (recv *DBusPropertyInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.signature =
		C.CString(recv.Signature)
	recv.native.flags =
		(C.GDBusPropertyInfoFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
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
	native *C.GDBusSignalInfo
	// ref_count : no type generator for gint, volatile gint
	Name string
	// no type for args
	// no type for annotations
}

func DBusSignalInfoNewFromC(u unsafe.Pointer) *DBusSignalInfo {
	c := (*C.GDBusSignalInfo)(u)
	if c == nil {
		return nil
	}

	g := &DBusSignalInfo{
		Name:   C.GoString(c.name),
		native: c,
	}

	return g
}

func (recv *DBusSignalInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)

	return (unsafe.Pointer)(recv.native)
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

// Unsupported : g_io_extension_get_type : no return generator

// Unsupported : g_io_extension_point_get_required_type : no return generator

// Unsupported : g_io_extension_point_set_required_type : unsupported parameter type : no type generator for GType, GType

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc, GSourceFunc

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

// Unsupported : g_resource_enumerate_children : no return type

// Unsupported : g_resource_get_info : unsupported parameter size : no type generator for gsize, gsize*

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// Unsupported : g_settings_schema_list_children : no return type

// Unsupported : g_settings_schema_list_keys : no return type

// Unsupported : g_settings_schema_key_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_schema_key_get_value_type : return type : Blacklisted record : GVariantType

// Unsupported : g_settings_schema_key_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : no param type

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

// Unsupported : g_unix_mount_point_guess_icon : no return generator

// Unsupported : g_unix_mount_point_guess_symbolic_icon : no return generator
