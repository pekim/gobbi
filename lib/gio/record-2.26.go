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

// Class structure for #GCredentials.
/*

C record/class : GCredentialsClass
*/
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

// Information about an annotation.
/*

C record/class : GDBusAnnotationInfo
*/
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

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_annotation_info_ref
*/
func (recv *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	retC := C.g_dbus_annotation_info_ref((*C.GDBusAnnotationInfo)(recv.native))
	retGo := DBusAnnotationInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_annotation_info_unref
*/
func (recv *DBusAnnotationInfo) Unref() {
	C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(recv.native))

	return
}

// Information about an argument for a method or a signal.
/*

C record/class : GDBusArgInfo
*/
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

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_arg_info_ref
*/
func (recv *DBusArgInfo) Ref() *DBusArgInfo {
	retC := C.g_dbus_arg_info_ref((*C.GDBusArgInfo)(recv.native))
	retGo := DBusArgInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_arg_info_unref
*/
func (recv *DBusArgInfo) Unref() {
	C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(recv.native))

	return
}

// Struct used in g_dbus_error_register_error_domain().
/*

C record/class : GDBusErrorEntry
*/
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

// Information about a D-Bus interface.
/*

C record/class : GDBusInterfaceInfo
*/
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

// Appends an XML representation of @info (and its children) to @string_builder.
//
// This function is typically used for generating introspection XML
// documents at run-time for handling the
// `org.freedesktop.DBus.Introspectable.Introspect`
// method.
/*

C function : g_dbus_interface_info_generate_xml
*/
func (recv *DBusInterfaceInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(C.NULL)
	if stringBuilder != nil {
		c_string_builder = (*C.GString)(stringBuilder.ToC())
	}

	C.g_dbus_interface_info_generate_xml((*C.GDBusInterfaceInfo)(recv.native), c_indent, c_string_builder)

	return
}

// Looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on @info.
/*

C function : g_dbus_interface_info_lookup_method
*/
func (recv *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_method((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
/*

C function : g_dbus_interface_info_lookup_property
*/
func (recv *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_property((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
/*

C function : g_dbus_interface_info_lookup_signal
*/
func (recv *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_interface_info_lookup_signal((*C.GDBusInterfaceInfo)(recv.native), c_name)
	retGo := DBusSignalInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_interface_info_ref
*/
func (recv *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_info_ref((*C.GDBusInterfaceInfo)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_interface_info_unref
*/
func (recv *DBusInterfaceInfo) Unref() {
	C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// Virtual table for handling properties and method calls for a D-Bus
// interface.
//
// Since 2.38, if you want to handle getting/setting D-Bus properties
// asynchronously, give %NULL as your get_property() or set_property()
// function. The D-Bus call will be directed to your @method_call function,
// with the provided @interface_name set to "org.freedesktop.DBus.Properties".
//
// Ownership of the #GDBusMethodInvocation object passed to the
// method_call() function is transferred to your handler; you must
// call one of the methods of #GDBusMethodInvocation to return a reply
// (possibly empty), or an error. These functions also take ownership
// of the passed-in invocation object, so unless the invocation
// object has otherwise been referenced, it will be then be freed.
// Calling one of these functions may be done within your
// method_call() implementation but it also can be done at a later
// point to handle the method asynchronously.
//
// The usual checks on the validity of the calls is performed. For
// `Get` calls, an error is automatically returned if the property does
// not exist or the permissions do not allow access. The same checks are
// performed for `Set` calls, and the provided value is also checked for
// being the correct type.
//
// For both `Get` and `Set` calls, the #GDBusMethodInvocation
// passed to the @method_call handler can be queried with
// g_dbus_method_invocation_get_property_info() to get a pointer
// to the #GDBusPropertyInfo of the property.
//
// If you have readable properties specified in your interface info,
// you must ensure that you either provide a non-%NULL @get_property()
// function or provide implementations of both the `Get` and `GetAll`
// methods on org.freedesktop.DBus.Properties interface in your @method_call
// function. Note that the required return type of the `Get` call is
// `(v)`, not the type of the property. `GetAll` expects a return value
// of type `a{sv}`.
//
// If you have writable properties specified in your interface info,
// you must ensure that you either provide a non-%NULL @set_property()
// function or provide an implementation of the `Set` call. If implementing
// the call, you must return the value of type %G_VARIANT_TYPE_UNIT.
/*

C record/class : GDBusInterfaceVTable
*/
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

// Information about a method on an D-Bus interface.
/*

C record/class : GDBusMethodInfo
*/
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

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_method_info_ref
*/
func (recv *DBusMethodInfo) Ref() *DBusMethodInfo {
	retC := C.g_dbus_method_info_ref((*C.GDBusMethodInfo)(recv.native))
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_method_info_unref
*/
func (recv *DBusMethodInfo) Unref() {
	C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(recv.native))

	return
}

// Information about nodes in a remote object hierarchy.
/*

C record/class : GDBusNodeInfo
*/
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

// Parses @xml_data and returns a #GDBusNodeInfo representing the data.
//
// The introspection XML must contain exactly one top-level
// <node> element.
//
// Note that this routine is using a
// [GMarkup][glib-Simple-XML-Subset-Parser.description]-based
// parser that only accepts a subset of valid XML documents.
/*

C function : g_dbus_node_info_new_for_xml
*/
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

// Appends an XML representation of @info (and its children) to @string_builder.
//
// This function is typically used for generating introspection XML documents at run-time for
// handling the `org.freedesktop.DBus.Introspectable.Introspect`  method.
/*

C function : g_dbus_node_info_generate_xml
*/
func (recv *DBusNodeInfo) GenerateXml(indent uint32, stringBuilder *glib.String) {
	c_indent := (C.guint)(indent)

	c_string_builder := (*C.GString)(C.NULL)
	if stringBuilder != nil {
		c_string_builder = (*C.GString)(stringBuilder.ToC())
	}

	C.g_dbus_node_info_generate_xml((*C.GDBusNodeInfo)(recv.native), c_indent, c_string_builder)

	return
}

// Looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
/*

C function : g_dbus_node_info_lookup_interface
*/
func (recv *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_dbus_node_info_lookup_interface((*C.GDBusNodeInfo)(recv.native), c_name)
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_node_info_ref
*/
func (recv *DBusNodeInfo) Ref() *DBusNodeInfo {
	retC := C.g_dbus_node_info_ref((*C.GDBusNodeInfo)(recv.native))
	retGo := DBusNodeInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_node_info_unref
*/
func (recv *DBusNodeInfo) Unref() {
	C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(recv.native))

	return
}

// Information about a D-Bus property on a D-Bus interface.
/*

C record/class : GDBusPropertyInfo
*/
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

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_property_info_ref
*/
func (recv *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	retC := C.g_dbus_property_info_ref((*C.GDBusPropertyInfo)(recv.native))
	retGo := DBusPropertyInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_property_info_unref
*/
func (recv *DBusPropertyInfo) Unref() {
	C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(recv.native))

	return
}

// Class structure for #GDBusProxy.
/*

C record/class : GDBusProxyClass
*/
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

// Information about a signal on a D-Bus interface.
/*

C record/class : GDBusSignalInfo
*/
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

// If @info is statically allocated does nothing. Otherwise increases
// the reference count.
/*

C function : g_dbus_signal_info_ref
*/
func (recv *DBusSignalInfo) Ref() *DBusSignalInfo {
	retC := C.g_dbus_signal_info_ref((*C.GDBusSignalInfo)(recv.native))
	retGo := DBusSignalInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0,
// the memory used is freed.
/*

C function : g_dbus_signal_info_unref
*/
func (recv *DBusSignalInfo) Unref() {
	C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(recv.native))

	return
}

// Virtual table for handling subtrees registered with g_dbus_connection_register_subtree().
/*

C record/class : GDBusSubtreeVTable
*/
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

// Class structure for #GProxyAddress.
/*

C record/class : GProxyAddressClass
*/
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

// Provides an interface for handling proxy connection and payload.
/*

C record/class : GProxyInterface
*/
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

// vtable for a #GTlsClientConnection implementation.
/*

C record/class : GTlsClientConnectionInterface
*/
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

// vtable for a #GTlsServerConnection implementation.
/*

C record/class : GTlsServerConnectionInterface
*/
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

// Class structure for #GUnixCredentialsMessage.
/*

C record/class : GUnixCredentialsMessageClass
*/
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
