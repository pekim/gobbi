// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Base type for D-Bus interfaces.
/*

C record/class : GDBusInterfaceIface
*/
type DBusInterfaceIface struct {
	native *C.GDBusInterfaceIface
	// parent_iface : record
	// no type for get_info
	// no type for get_object
	// no type for set_object
	// no type for dup_object
}

func DBusInterfaceIfaceNewFromC(u unsafe.Pointer) *DBusInterfaceIface {
	c := (*C.GDBusInterfaceIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceIface{native: c}

	return g
}

func (recv *DBusInterfaceIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(),
// g_dbus_interface_info_lookup_signal() and
// g_dbus_interface_info_lookup_property().
//
// If this has already been called with @info, the existing cache is
// used and its use count is increased.
//
// Note that @info cannot be modified until
// g_dbus_interface_info_cache_release() is called.
/*

C function : g_dbus_interface_info_cache_build
*/
func (recv *DBusInterfaceInfo) CacheBuild() {
	C.g_dbus_interface_info_cache_build((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// Decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the
// resources used by the cache if the usage count drops to zero.
/*

C function : g_dbus_interface_info_cache_release
*/
func (recv *DBusInterfaceInfo) CacheRelease() {
	C.g_dbus_interface_info_cache_release((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// Class structure for #GDBusInterfaceSkeleton.
/*

C record/class : GDBusInterfaceSkeletonClass
*/
type DBusInterfaceSkeletonClass struct {
	native *C.GDBusInterfaceSkeletonClass
	// parent_class : record
	// no type for get_info
	// no type for get_vtable
	// no type for get_properties
	// no type for flush
	// Private : vfunc_padding
	// no type for g_authorize_method
	// Private : signal_padding
}

func DBusInterfaceSkeletonClassNewFromC(u unsafe.Pointer) *DBusInterfaceSkeletonClass {
	c := (*C.GDBusInterfaceSkeletonClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeletonClass{native: c}

	return g
}

func (recv *DBusInterfaceSkeletonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Base object type for D-Bus objects.
/*

C record/class : GDBusObjectIface
*/
type DBusObjectIface struct {
	native *C.GDBusObjectIface
	// parent_iface : record
	// no type for get_object_path
	// no type for get_interfaces
	// no type for get_interface
	// no type for interface_added
	// no type for interface_removed
}

func DBusObjectIfaceNewFromC(u unsafe.Pointer) *DBusObjectIface {
	c := (*C.GDBusObjectIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectIface{native: c}

	return g
}

func (recv *DBusObjectIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Class structure for #GDBusObjectManagerClient.
/*

C record/class : GDBusObjectManagerClientClass
*/
type DBusObjectManagerClientClass struct {
	native *C.GDBusObjectManagerClientClass
	// parent_class : record
	// no type for interface_proxy_signal
	// no type for interface_proxy_properties_changed
	// Private : padding
}

func DBusObjectManagerClientClassNewFromC(u unsafe.Pointer) *DBusObjectManagerClientClass {
	c := (*C.GDBusObjectManagerClientClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClientClass{native: c}

	return g
}

func (recv *DBusObjectManagerClientClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Base type for D-Bus object managers.
/*

C record/class : GDBusObjectManagerIface
*/
type DBusObjectManagerIface struct {
	native *C.GDBusObjectManagerIface
	// parent_iface : record
	// no type for get_object_path
	// no type for get_objects
	// no type for get_object
	// no type for get_interface
	// no type for object_added
	// no type for object_removed
	// no type for interface_added
	// no type for interface_removed
}

func DBusObjectManagerIfaceNewFromC(u unsafe.Pointer) *DBusObjectManagerIface {
	c := (*C.GDBusObjectManagerIface)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerIface{native: c}

	return g
}

func (recv *DBusObjectManagerIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Class structure for #GDBusObjectManagerServer.
/*

C record/class : GDBusObjectManagerServerClass
*/
type DBusObjectManagerServerClass struct {
	native *C.GDBusObjectManagerServerClass
	// parent_class : record
	// Private : padding
}

func DBusObjectManagerServerClassNewFromC(u unsafe.Pointer) *DBusObjectManagerServerClass {
	c := (*C.GDBusObjectManagerServerClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServerClass{native: c}

	return g
}

func (recv *DBusObjectManagerServerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Class structure for #GDBusObjectProxy.
/*

C record/class : GDBusObjectProxyClass
*/
type DBusObjectProxyClass struct {
	native *C.GDBusObjectProxyClass
	// parent_class : record
	// Private : padding
}

func DBusObjectProxyClassNewFromC(u unsafe.Pointer) *DBusObjectProxyClass {
	c := (*C.GDBusObjectProxyClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxyClass{native: c}

	return g
}

func (recv *DBusObjectProxyClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Class structure for #GDBusObjectSkeleton.
/*

C record/class : GDBusObjectSkeletonClass
*/
type DBusObjectSkeletonClass struct {
	native *C.GDBusObjectSkeletonClass
	// parent_class : record
	// no type for authorize_method
	// Private : padding
}

func DBusObjectSkeletonClassNewFromC(u unsafe.Pointer) *DBusObjectSkeletonClass {
	c := (*C.GDBusObjectSkeletonClass)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeletonClass{native: c}

	return g
}

func (recv *DBusObjectSkeletonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Represents a scope for loading IO modules. A scope can be used for blocking
// duplicate modules, or blocking a module you don't want to load.
//
// The scope can be used with g_io_modules_load_all_in_directory_with_scope()
// or g_io_modules_scan_all_in_directory_with_scope().
/*

C record/class : GIOModuleScope
*/
type IOModuleScope struct {
	native *C.GIOModuleScope
}

func IOModuleScopeNewFromC(u unsafe.Pointer) *IOModuleScope {
	c := (*C.GIOModuleScope)(u)
	if c == nil {
		return nil
	}

	g := &IOModuleScope{native: c}

	return g
}

func (recv *IOModuleScope) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Block modules with the given @basename from being loaded when
// this scope is used with g_io_modules_scan_all_in_directory_with_scope()
// or g_io_modules_load_all_in_directory_with_scope().
/*

C function : g_io_module_scope_block
*/
func (recv *IOModuleScope) Block(basename string) {
	c_basename := C.CString(basename)
	defer C.free(unsafe.Pointer(c_basename))

	C.g_io_module_scope_block((*C.GIOModuleScope)(recv.native), c_basename)

	return
}

// Free a module scope.
/*

C function : g_io_module_scope_free
*/
func (recv *IOModuleScope) Free() {
	C.g_io_module_scope_free((*C.GIOModuleScope)(recv.native))

	return
}

// The class for #GTlsDatabase. Derived classes should implement the various
// virtual methods. _async and _finish methods have a default
// implementation that runs the corresponding sync method in a thread.
/*

C record/class : GTlsDatabaseClass
*/
type TlsDatabaseClass struct {
	native *C.GTlsDatabaseClass
	// parent_class : record
	// no type for verify_chain
	// no type for verify_chain_async
	// no type for verify_chain_finish
	// no type for create_certificate_handle
	// no type for lookup_certificate_for_handle
	// no type for lookup_certificate_for_handle_async
	// no type for lookup_certificate_for_handle_finish
	// no type for lookup_certificate_issuer
	// no type for lookup_certificate_issuer_async
	// no type for lookup_certificate_issuer_finish
	// no type for lookup_certificates_issued_by
	// no type for lookup_certificates_issued_by_async
	// no type for lookup_certificates_issued_by_finish
	// Private : padding
}

func TlsDatabaseClassNewFromC(u unsafe.Pointer) *TlsDatabaseClass {
	c := (*C.GTlsDatabaseClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabaseClass{native: c}

	return g
}

func (recv *TlsDatabaseClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The class for #GTlsInteraction. Derived classes implement the various
// virtual interaction methods to handle TLS interactions.
//
// Derived classes can choose to implement whichever interactions methods they'd
// like to support by overriding those virtual methods in their class
// initialization function. If a derived class implements an async method,
// it must also implement the corresponding finish method.
//
// The synchronous interaction methods should implement to display modal dialogs,
// and the asynchronous methods to display modeless dialogs.
//
// If the user cancels an interaction, then the result should be
// %G_TLS_INTERACTION_FAILED and the error should be set with a domain of
// %G_IO_ERROR and code of %G_IO_ERROR_CANCELLED.
/*

C record/class : GTlsInteractionClass
*/
type TlsInteractionClass struct {
	native *C.GTlsInteractionClass
	// Private : parent_class
	// no type for ask_password
	// no type for ask_password_async
	// no type for ask_password_finish
	// no type for request_certificate
	// no type for request_certificate_async
	// no type for request_certificate_finish
	// Private : padding
}

func TlsInteractionClassNewFromC(u unsafe.Pointer) *TlsInteractionClass {
	c := (*C.GTlsInteractionClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteractionClass{native: c}

	return g
}

func (recv *TlsInteractionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
