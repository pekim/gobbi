// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_action_change_state : unsupported parameter value : Blacklisted record : GVariant

// DBusInterface is a wrapper around the C record GDBusInterface.
type DBusInterface struct {
	native *C.GDBusInterface
}

func DBusInterfaceNewFromC(u unsafe.Pointer) *DBusInterface {
	c := (*C.GDBusInterface)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterface{native: c}

	return g
}

func (recv *DBusInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetInfo is a wrapper around the C function g_dbus_interface_get_info.
func (recv *DBusInterface) GetInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_interface_get_info((*C.GDBusInterface)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_interface_get_object : no return generator

// Unsupported : g_dbus_interface_set_object : unsupported parameter object : no type generator for DBusObject, GDBusObject*

// Unsupported signal 'interface-added' for DBusObject : unsupported parameter interface : no type generator for DBusInterface,

// Unsupported signal 'interface-removed' for DBusObject : unsupported parameter interface : no type generator for DBusInterface,

// Unsupported : g_dbus_object_get_interface : no return generator

// GetInterfaces is a wrapper around the C function g_dbus_object_get_interfaces.
func (recv *DBusObject) GetInterfaces() *glib.List {
	retC := C.g_dbus_object_get_interfaces((*C.GDBusObject)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_object_get_object_path.
func (recv *DBusObject) GetObjectPath() string {
	retC := C.g_dbus_object_get_object_path((*C.GDBusObject)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported signal 'interface-added' for DBusObjectManager : unsupported parameter object : no type generator for DBusObject,

// Unsupported signal 'interface-removed' for DBusObjectManager : unsupported parameter object : no type generator for DBusObject,

// Unsupported signal 'object-added' for DBusObjectManager : unsupported parameter object : no type generator for DBusObject,

// Unsupported signal 'object-removed' for DBusObjectManager : unsupported parameter object : no type generator for DBusObject,

// Unsupported : g_dbus_object_manager_get_interface : no return generator

// Unsupported : g_dbus_object_manager_get_object : no return generator

// GetObjectPath is a wrapper around the C function g_dbus_object_manager_get_object_path.
func (recv *DBusObjectManager) GetObjectPath() string {
	retC := C.g_dbus_object_manager_get_object_path((*C.GDBusObjectManager)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetObjects is a wrapper around the C function g_dbus_object_manager_get_objects.
func (recv *DBusObjectManager) GetObjects() *glib.List {
	retC := C.g_dbus_object_manager_get_objects((*C.GDBusObjectManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultDatabase is a wrapper around the C function g_tls_backend_get_default_database.
func (recv *TlsBackend) GetDefaultDatabase() *TlsDatabase {
	retC := C.g_tls_backend_get_default_database((*C.GTlsBackend)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_tls_backend_get_file_database_type : no return generator

// TlsFileDatabase is a wrapper around the C record GTlsFileDatabase.
type TlsFileDatabase struct {
	native *C.GTlsFileDatabase
}

func TlsFileDatabaseNewFromC(u unsafe.Pointer) *TlsFileDatabase {
	c := (*C.GTlsFileDatabase)(u)
	if c == nil {
		return nil
	}

	g := &TlsFileDatabase{native: c}

	return g
}

func (recv *TlsFileDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
