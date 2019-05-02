// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Equals compares this DBusInterfaceIface with another DBusInterfaceIface, and returns true if they represent the same GObject.
func (recv *DBusInterfaceIface) Equals(other *DBusInterfaceIface) bool {
	return other.ToC() == recv.ToC()
}

// CacheBuild is a wrapper around the C function g_dbus_interface_info_cache_build.
func (recv *DBusInterfaceInfo) CacheBuild() {
	C.g_dbus_interface_info_cache_build((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// CacheRelease is a wrapper around the C function g_dbus_interface_info_cache_release.
func (recv *DBusInterfaceInfo) CacheRelease() {
	C.g_dbus_interface_info_cache_release((*C.GDBusInterfaceInfo)(recv.native))

	return
}

// Equals compares this DBusInterfaceSkeletonClass with another DBusInterfaceSkeletonClass, and returns true if they represent the same GObject.
func (recv *DBusInterfaceSkeletonClass) Equals(other *DBusInterfaceSkeletonClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectIface with another DBusObjectIface, and returns true if they represent the same GObject.
func (recv *DBusObjectIface) Equals(other *DBusObjectIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectManagerClientClass with another DBusObjectManagerClientClass, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerClientClass) Equals(other *DBusObjectManagerClientClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectManagerIface with another DBusObjectManagerIface, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerIface) Equals(other *DBusObjectManagerIface) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectManagerServerClass with another DBusObjectManagerServerClass, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerServerClass) Equals(other *DBusObjectManagerServerClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectProxyClass with another DBusObjectProxyClass, and returns true if they represent the same GObject.
func (recv *DBusObjectProxyClass) Equals(other *DBusObjectProxyClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this DBusObjectSkeletonClass with another DBusObjectSkeletonClass, and returns true if they represent the same GObject.
func (recv *DBusObjectSkeletonClass) Equals(other *DBusObjectSkeletonClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this IOModuleScope with another IOModuleScope, and returns true if they represent the same GObject.
func (recv *IOModuleScope) Equals(other *IOModuleScope) bool {
	return other.ToC() == recv.ToC()
}

// IOModuleScopeNew is a wrapper around the C function g_io_module_scope_new.
func IOModuleScopeNew(flags IOModuleScopeFlags) *IOModuleScope {
	c_flags := (C.GIOModuleScopeFlags)(flags)

	retC := C.g_io_module_scope_new(c_flags)
	retGo := IOModuleScopeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Block is a wrapper around the C function g_io_module_scope_block.
func (recv *IOModuleScope) Block(basename string) {
	c_basename := C.CString(basename)
	defer C.free(unsafe.Pointer(c_basename))

	C.g_io_module_scope_block((*C.GIOModuleScope)(recv.native), c_basename)

	return
}

// Free is a wrapper around the C function g_io_module_scope_free.
func (recv *IOModuleScope) Free() {
	C.g_io_module_scope_free((*C.GIOModuleScope)(recv.native))

	return
}

// Equals compares this TlsDatabaseClass with another TlsDatabaseClass, and returns true if they represent the same GObject.
func (recv *TlsDatabaseClass) Equals(other *TlsDatabaseClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this TlsInteractionClass with another TlsInteractionClass, and returns true if they represent the same GObject.
func (recv *TlsInteractionClass) Equals(other *TlsInteractionClass) bool {
	return other.ToC() == recv.ToC()
}
