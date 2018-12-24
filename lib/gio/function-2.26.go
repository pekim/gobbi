// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// BusGetFinish is a wrapper around the C function g_bus_get_finish.
func BusGetFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_bus_get_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// BusGetSync is a wrapper around the C function g_bus_get_sync.
func BusGetSync(busType BusType, cancellable *Cancellable) (*DBusConnection, error) {
	c_bus_type := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_bus_get_sync(c_bus_type, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback (GBusAcquiredCallback) for param bus_acquired_handler

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback (GBusNameAcquiredCallback) for param name_acquired_handler

// BusOwnNameOnConnectionWithClosures is a wrapper around the C function g_bus_own_name_on_connection_with_closures.
func BusOwnNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_name_acquired_closure := (*C.GClosure)(C.NULL)
	if nameAcquiredClosure != nil {
		c_name_acquired_closure = (*C.GClosure)(nameAcquiredClosure.ToC())
	}

	c_name_lost_closure := (*C.GClosure)(C.NULL)
	if nameLostClosure != nil {
		c_name_lost_closure = (*C.GClosure)(nameLostClosure.ToC())
	}

	retC := C.g_bus_own_name_on_connection_with_closures(c_connection, c_name, c_flags, c_name_acquired_closure, c_name_lost_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusOwnNameWithClosures is a wrapper around the C function g_bus_own_name_with_closures.
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_bus_acquired_closure := (*C.GClosure)(C.NULL)
	if busAcquiredClosure != nil {
		c_bus_acquired_closure = (*C.GClosure)(busAcquiredClosure.ToC())
	}

	c_name_acquired_closure := (*C.GClosure)(C.NULL)
	if nameAcquiredClosure != nil {
		c_name_acquired_closure = (*C.GClosure)(nameAcquiredClosure.ToC())
	}

	c_name_lost_closure := (*C.GClosure)(C.NULL)
	if nameLostClosure != nil {
		c_name_lost_closure = (*C.GClosure)(nameLostClosure.ToC())
	}

	retC := C.g_bus_own_name_with_closures(c_bus_type, c_name, c_flags, c_bus_acquired_closure, c_name_acquired_closure, c_name_lost_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusUnownName is a wrapper around the C function g_bus_unown_name.
func BusUnownName(ownerId uint32) {
	c_owner_id := (C.guint)(ownerId)

	C.g_bus_unown_name(c_owner_id)

	return
}

// BusUnwatchName is a wrapper around the C function g_bus_unwatch_name.
func BusUnwatchName(watcherId uint32) {
	c_watcher_id := (C.guint)(watcherId)

	C.g_bus_unwatch_name(c_watcher_id)

	return
}

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// BusWatchNameOnConnectionWithClosures is a wrapper around the C function g_bus_watch_name_on_connection_with_closures.
func BusWatchNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_name_appeared_closure := (*C.GClosure)(C.NULL)
	if nameAppearedClosure != nil {
		c_name_appeared_closure = (*C.GClosure)(nameAppearedClosure.ToC())
	}

	c_name_vanished_closure := (*C.GClosure)(C.NULL)
	if nameVanishedClosure != nil {
		c_name_vanished_closure = (*C.GClosure)(nameVanishedClosure.ToC())
	}

	retC := C.g_bus_watch_name_on_connection_with_closures(c_connection, c_name, c_flags, c_name_appeared_closure, c_name_vanished_closure)
	retGo := (uint32)(retC)

	return retGo
}

// BusWatchNameWithClosures is a wrapper around the C function g_bus_watch_name_with_closures.
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint32 {
	c_bus_type := (C.GBusType)(busType)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_name_appeared_closure := (*C.GClosure)(C.NULL)
	if nameAppearedClosure != nil {
		c_name_appeared_closure = (*C.GClosure)(nameAppearedClosure.ToC())
	}

	c_name_vanished_closure := (*C.GClosure)(C.NULL)
	if nameVanishedClosure != nil {
		c_name_vanished_closure = (*C.GClosure)(nameVanishedClosure.ToC())
	}

	retC := C.g_bus_watch_name_with_closures(c_bus_type, c_name, c_flags, c_name_appeared_closure, c_name_vanished_closure)
	retGo := (uint32)(retC)

	return retGo
}

// DbusAddressGetForBusSync is a wrapper around the C function g_dbus_address_get_for_bus_sync.
func DbusAddressGetForBusSync(busType BusType, cancellable *Cancellable) (string, error) {
	c_bus_type := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_for_bus_sync(c_bus_type, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// DbusAddressGetStreamFinish is a wrapper around the C function g_dbus_address_get_stream_finish.
func DbusAddressGetStreamFinish(res *AsyncResult) (*IOStream, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_out_guid *C.gchar

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_stream_finish(c_res, &c_out_guid, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outGuid := C.GoString(c_out_guid)
	defer C.free(unsafe.Pointer(c_out_guid))

	return retGo, outGuid, goError
}

// DbusAddressGetStreamSync is a wrapper around the C function g_dbus_address_get_stream_sync.
func DbusAddressGetStreamSync(address string, cancellable *Cancellable) (*IOStream, string, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	var c_out_guid *C.gchar

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_stream_sync(c_address, &c_out_guid, c_cancellable, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outGuid := C.GoString(c_out_guid)
	defer C.free(unsafe.Pointer(c_out_guid))

	return retGo, outGuid, goError
}

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations :

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries :

// DbusGenerateGuid is a wrapper around the C function g_dbus_generate_guid.
func DbusGenerateGuid() string {
	retC := C.g_dbus_generate_guid()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DbusIsAddress is a wrapper around the C function g_dbus_is_address.
func DbusIsAddress(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_address(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsGuid is a wrapper around the C function g_dbus_is_guid.
func DbusIsGuid(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_guid(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsInterfaceName is a wrapper around the C function g_dbus_is_interface_name.
func DbusIsInterfaceName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_interface_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsMemberName is a wrapper around the C function g_dbus_is_member_name.
func DbusIsMemberName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_member_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsName is a wrapper around the C function g_dbus_is_name.
func DbusIsName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsSupportedAddress is a wrapper around the C function g_dbus_is_supported_address.
func DbusIsSupportedAddress(string string) (bool, error) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	var cThrowableError *C.GError

	retC := C.g_dbus_is_supported_address(c_string, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DbusIsUniqueName is a wrapper around the C function g_dbus_is_unique_name.
func DbusIsUniqueName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_unique_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}
