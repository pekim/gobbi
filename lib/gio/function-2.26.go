// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Unsupported : g_bus_get : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_bus_get().
//
// The returned object is a singleton, that is, shared with other
// callers of g_bus_get() and g_bus_get_sync() for @bus_type. In the
// event that you need a private message bus connection, use
// g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned #GDBusConnection object will (usually) have
// the #GDBusConnection:exit-on-close property set to %TRUE.
/*

C function : g_bus_get_finish
*/
func BusGetFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_bus_get_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously connects to the message bus specified by @bus_type.
// Note that the returned object may shared with other callers,
// e.g. if two separate parts of a process calls this function with
// the same @bus_type, they will share the same object.
//
// This is a synchronous failable function. See g_bus_get() and
// g_bus_get_finish() for the asynchronous version.
//
// The returned object is a singleton, that is, shared with other
// callers of g_bus_get() and g_bus_get_sync() for @bus_type. In the
// event that you need a private message bus connection, use
// g_dbus_address_get_for_bus_sync() and
// g_dbus_connection_new_for_address().
//
// Note that the returned #GDBusConnection object will (usually) have
// the #GDBusConnection:exit-on-close property set to %TRUE.
/*

C function : g_bus_get_sync
*/
func BusGetSync(busType BusType, cancellable *Cancellable) (*DBusConnection, error) {
	c_bus_type := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_bus_get_sync(c_bus_type, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bus_own_name : unsupported parameter bus_acquired_handler : no type generator for BusAcquiredCallback (GBusAcquiredCallback) for param bus_acquired_handler

// Unsupported : g_bus_own_name_on_connection : unsupported parameter name_acquired_handler : no type generator for BusNameAcquiredCallback (GBusNameAcquiredCallback) for param name_acquired_handler

// Version of g_bus_own_name_on_connection() using closures instead of
// callbacks for easier binding in other languages.
/*

C function : g_bus_own_name_on_connection_with_closures
*/
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

// Version of g_bus_own_name() using closures instead of callbacks for
// easier binding in other languages.
/*

C function : g_bus_own_name_with_closures
*/
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

// Stops owning a name.
/*

C function : g_bus_unown_name
*/
func BusUnownName(ownerId uint32) {
	c_owner_id := (C.guint)(ownerId)

	C.g_bus_unown_name(c_owner_id)

	return
}

// Stops watching a name.
/*

C function : g_bus_unwatch_name
*/
func BusUnwatchName(watcherId uint32) {
	c_watcher_id := (C.guint)(watcherId)

	C.g_bus_unwatch_name(c_watcher_id)

	return
}

// Unsupported : g_bus_watch_name : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Unsupported : g_bus_watch_name_on_connection : unsupported parameter name_appeared_handler : no type generator for BusNameAppearedCallback (GBusNameAppearedCallback) for param name_appeared_handler

// Version of g_bus_watch_name_on_connection() using closures instead of callbacks for
// easier binding in other languages.
/*

C function : g_bus_watch_name_on_connection_with_closures
*/
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

// Version of g_bus_watch_name() using closures instead of callbacks for
// easier binding in other languages.
/*

C function : g_bus_watch_name_with_closures
*/
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

// Synchronously looks up the D-Bus address for the well-known message
// bus instance specified by @bus_type. This may involve using various
// platform specific mechanisms.
//
// The returned address will be in the
// [D-Bus address format](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
/*

C function : g_dbus_address_get_for_bus_sync
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_address_get_stream : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an operation started with g_dbus_address_get_stream().
/*

C function : g_dbus_address_get_stream_finish
*/
func DbusAddressGetStreamFinish(res *AsyncResult) (*IOStream, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_out_guid *C.gchar

	var cThrowableError *C.GError

	retC := C.g_dbus_address_get_stream_finish(c_res, &c_out_guid, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outGuid := C.GoString(c_out_guid)
	defer C.free(unsafe.Pointer(c_out_guid))

	return retGo, outGuid, goThrowableError
}

// Synchronously connects to an endpoint specified by @address and
// sets up the connection so it is in a state to run the client-side
// of the D-Bus authentication conversation. @address must be in the
// [D-Bus address format](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This is a synchronous failable function. See
// g_dbus_address_get_stream() for the asynchronous version.
/*

C function : g_dbus_address_get_stream_sync
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outGuid := C.GoString(c_out_guid)
	defer C.free(unsafe.Pointer(c_out_guid))

	return retGo, outGuid, goThrowableError
}

// Unsupported : g_dbus_annotation_info_lookup : unsupported parameter annotations :

// Creates a D-Bus error name to use for @error. If @error matches
// a registered error (cf. g_dbus_error_register_error()), the corresponding
// D-Bus error name will be returned.
//
// Otherwise the a name of the form
// `org.gtk.GDBus.UnmappedGError.Quark._ESCAPED_QUARK_NAME.Code_ERROR_CODE`
// will be used. This allows other GDBus applications to map the error
// on the wire back to a #GError using g_dbus_error_new_for_dbus_error().
//
// This function is typically only used in object mappings to put a
// #GError on the wire. Regular applications should not use it.
/*

C function : g_dbus_error_encode_gerror
*/
func DbusErrorEncodeGerror(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_encode_gerror(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the D-Bus error name used for @error, if any.
//
// This function is guaranteed to return a D-Bus error name for all
// #GErrors returned from functions handling remote method calls
// (e.g. g_dbus_connection_call_finish()) unless
// g_dbus_error_strip_remote_error() has been used on @error.
/*

C function : g_dbus_error_get_remote_error
*/
func DbusErrorGetRemoteError(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_get_remote_error(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Checks if @error represents an error received via D-Bus from a remote peer. If so,
// use g_dbus_error_get_remote_error() to get the name of the error.
/*

C function : g_dbus_error_is_remote_error
*/
func DbusErrorIsRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_is_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// Creates a #GError based on the contents of @dbus_error_name and
// @dbus_error_message.
//
// Errors registered with g_dbus_error_register_error() will be looked
// up using @dbus_error_name and if a match is found, the error domain
// and code is used. Applications can use g_dbus_error_get_remote_error()
// to recover @dbus_error_name.
//
// If a match against a registered error is not found and the D-Bus
// error name is in a form as returned by g_dbus_error_encode_gerror()
// the error domain and code encoded in the name is used to
// create the #GError. Also, @dbus_error_name is added to the error message
// such that it can be recovered with g_dbus_error_get_remote_error().
//
// Otherwise, a #GError with the error code %G_IO_ERROR_DBUS_ERROR
// in the #G_IO_ERROR error domain is returned. Also, @dbus_error_name is
// added to the error message such that it can be recovered with
// g_dbus_error_get_remote_error().
//
// In all three cases, @dbus_error_name can always be recovered from the
// returned #GError using the g_dbus_error_get_remote_error() function
// (unless g_dbus_error_strip_remote_error() hasn't been used on the returned error).
//
// This function is typically only used in object mappings to prepare
// #GError instances for applications. Regular applications should not use
// it.
/*

C function : g_dbus_error_new_for_dbus_error
*/
func DbusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	retC := C.g_dbus_error_new_for_dbus_error(c_dbus_error_name, c_dbus_error_message)
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates an association to map between @dbus_error_name and
// #GErrors specified by @error_domain and @error_code.
//
// This is typically done in the routine that returns the #GQuark for
// an error domain.
/*

C function : g_dbus_error_register_error
*/
func DbusErrorRegisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_register_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_error_register_error_domain : unsupported parameter entries :

// Looks for extra information in the error message used to recover
// the D-Bus error name and strips it if found. If stripped, the
// message field in @error will correspond exactly to what was
// received on the wire.
//
// This is typically used when presenting errors to the end user.
/*

C function : g_dbus_error_strip_remote_error
*/
func DbusErrorStripRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_strip_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// Destroys an association previously set up with g_dbus_error_register_error().
/*

C function : g_dbus_error_unregister_error
*/
func DbusErrorUnregisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_unregister_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// Generate a D-Bus GUID that can be used with
// e.g. g_dbus_connection_new().
//
// See the D-Bus specification regarding what strings are valid D-Bus
// GUID (for example, D-Bus GUIDs are not RFC-4122 compliant).
/*

C function : g_dbus_generate_guid
*/
func DbusGenerateGuid() string {
	retC := C.g_dbus_generate_guid()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Checks if @string is a
// [D-Bus address](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This doesn't check if @string is actually supported by #GDBusServer
// or #GDBusConnection - use g_dbus_is_supported_address() to do more
// checks.
/*

C function : g_dbus_is_address
*/
func DbusIsAddress(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_address(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @string is a D-Bus GUID.
//
// See the D-Bus specification regarding what strings are valid D-Bus
// GUID (for example, D-Bus GUIDs are not RFC-4122 compliant).
/*

C function : g_dbus_is_guid
*/
func DbusIsGuid(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_guid(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @string is a valid D-Bus interface name.
/*

C function : g_dbus_is_interface_name
*/
func DbusIsInterfaceName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_interface_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @string is a valid D-Bus member (e.g. signal or method) name.
/*

C function : g_dbus_is_member_name
*/
func DbusIsMemberName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_member_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @string is a valid D-Bus bus name (either unique or well-known).
/*

C function : g_dbus_is_name
*/
func DbusIsName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Like g_dbus_is_address() but also checks if the library supports the
// transports in @string and that key/value pairs for each transport
// are valid. See the specification of the
// [D-Bus address format](https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
/*

C function : g_dbus_is_supported_address
*/
func DbusIsSupportedAddress(string string) (bool, error) {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	var cThrowableError *C.GError

	retC := C.g_dbus_is_supported_address(c_string, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if @string is a valid D-Bus unique bus name.
/*

C function : g_dbus_is_unique_name
*/
func DbusIsUniqueName(string string) bool {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_unique_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Lookup "gio-proxy" extension point for a proxy implementation that supports
// specified protocol.
/*

C function : g_proxy_get_default_for_protocol
*/
func ProxyGetDefaultForProtocol(protocol string) *Proxy {
	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	retC := C.g_proxy_get_default_for_protocol(c_protocol)
	retGo := ProxyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the default #GProxyResolver for the system.
/*

C function : g_proxy_resolver_get_default
*/
func ProxyResolverGetDefault() *ProxyResolver {
	retC := C.g_proxy_resolver_get_default()
	retGo := ProxyResolverNewFromC(unsafe.Pointer(retC))

	return retGo
}
