// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// Unsupported : g_bus_get : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_bus_get_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_bus_get_sync : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback, GBusAcquiredCallback

// Unsupported : g_bus_own_name_on_connection : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_bus_own_name_on_connection_with_closures : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// BusOwnNameWithClosures is a wrapper around the C function g_bus_own_name_with_closures.
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_bus_acquired_closure := busAcquiredClosure.toC()

	c_name_acquired_closure := nameAcquiredClosure.toC()

	c_name_lost_closure := nameLostClosure.toC()

	retC := C.g_bus_own_name_with_closures(c_bus_type, c_name, c_flags, c_bus_acquired_closure, c_name_acquired_closure, c_name_lost_closure)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_bus_unown_name : no return generator

// Unsupported : g_bus_unwatch_name : no return generator

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback, GBusNameAppearedCallback

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_bus_watch_name_on_connection_with_closures : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// BusWatchNameWithClosures is a wrapper around the C function g_bus_watch_name_with_closures.
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_name_appeared_closure := nameAppearedClosure.toC()

	c_name_vanished_closure := nameVanishedClosure.toC()

	retC := C.g_bus_watch_name_with_closures(c_bus_type, c_name, c_flags, c_name_appeared_closure, c_name_vanished_closure)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_dbus_address_get_for_bus_sync : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_address_get_stream : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_address_get_stream_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_address_get_stream_sync : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations : no param type

// DbusErrorEncodeGerror is a wrapper around the C function g_dbus_error_encode_gerror.
func DbusErrorEncodeGerror(error *glib.Error) string {
	c_error := error.toC()

	retC := C.g_dbus_error_encode_gerror(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DbusErrorGetRemoteError is a wrapper around the C function g_dbus_error_get_remote_error.
func DbusErrorGetRemoteError(error *glib.Error) string {
	c_error := error.toC()

	retC := C.g_dbus_error_get_remote_error(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_error_is_remote_error : no return generator

// DbusErrorNewForDbusError is a wrapper around the C function g_dbus_error_new_for_dbus_error.
func DbusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	retC := C.g_dbus_error_new_for_dbus_error(c_dbus_error_name, c_dbus_error_message)
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_error_register_error : no return generator

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter quark_volatile : no type generator for gsize, volatile gsize*

// Unsupported : g_dbus_error_strip_remote_error : no return generator

// Unsupported : g_dbus_error_unregister_error : no return generator

// DbusGenerateGuid is a wrapper around the C function g_dbus_generate_guid.
func DbusGenerateGuid() string {
	retC := C.g_dbus_generate_guid()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_is_address : no return generator

// Unsupported : g_dbus_is_guid : no return generator

// Unsupported : g_dbus_is_interface_name : no return generator

// Unsupported : g_dbus_is_member_name : no return generator

// Unsupported : g_dbus_is_name : no return generator

// Unsupported : g_dbus_is_supported_address : no return generator

// Unsupported : g_dbus_is_unique_name : no return generator

// Unsupported : g_proxy_get_default_for_protocol : no return generator

// Unsupported : g_proxy_resolver_get_default : no return generator
