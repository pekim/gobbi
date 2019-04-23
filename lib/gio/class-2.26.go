// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
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
/*

	gboolean dbusauthobserver_authorizeAuthenticatedPeerHandler(GObject *, GIOStream *, GCredentials *, gpointer);

	static gulong DBusAuthObserver_signal_connect_authorize_authenticated_peer(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "authorize-authenticated-peer", G_CALLBACK(dbusauthobserver_authorizeAuthenticatedPeerHandler), data);
	}

*/
/*

	void dbusconnection_closedHandler(GObject *, gboolean, GError *, gpointer);

	static gulong DBusConnection_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(dbusconnection_closedHandler), data);
	}

*/
/*

	void dbusproxy_gSignalHandler(GObject *, gchar*, gchar*, GVariant *, gpointer);

	static gulong DBusProxy_signal_connect_g_signal(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "g-signal", G_CALLBACK(dbusproxy_gSignalHandler), data);
	}

*/
/*

	gboolean dbusserver_newConnectionHandler(GObject *, GDBusConnection *, gpointer);

	static gulong DBusServer_signal_connect_new_connection(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "new-connection", G_CALLBACK(dbusserver_newConnectionHandler), data);
	}

*/
import "C"

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native *C.GCredentials
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	c := (*C.GCredentials)(u)
	if c == nil {
		return nil
	}

	g := &Credentials{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Credentials) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Credentials with another Credentials, and returns true if they represent the same GObject.
func (recv *Credentials) Equals(other *Credentials) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Credentials) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Credentials.
// Exercise care, as this is a potentially dangerous function if the Object is not a Credentials.
func CastToCredentials(object *gobject.Object) *Credentials {
	return CredentialsNewFromC(object.ToC())
}

// Blacklisted : g_credentials_new

// Unsupported : g_credentials_get_native : no return generator

// Blacklisted : g_credentials_get_unix_user

// Blacklisted : g_credentials_is_same_user

// Unsupported : g_credentials_set_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// Blacklisted : g_credentials_set_unix_user

// Blacklisted : g_credentials_to_string

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native *C.GDBusAuthObserver
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	c := (*C.GDBusAuthObserver)(u)
	if c == nil {
		return nil
	}

	g := &DBusAuthObserver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusAuthObserver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusAuthObserver with another DBusAuthObserver, and returns true if they represent the same GObject.
func (recv *DBusAuthObserver) Equals(other *DBusAuthObserver) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusAuthObserver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusAuthObserver.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusAuthObserver.
func CastToDBusAuthObserver(object *gobject.Object) *DBusAuthObserver {
	return DBusAuthObserverNewFromC(object.ToC())
}

type signalDBusAuthObserverAuthorizeAuthenticatedPeerDetail struct {
	callback  DBusAuthObserverSignalAuthorizeAuthenticatedPeerCallback
	handlerID C.gulong
}

var signalDBusAuthObserverAuthorizeAuthenticatedPeerId int
var signalDBusAuthObserverAuthorizeAuthenticatedPeerMap = make(map[int]signalDBusAuthObserverAuthorizeAuthenticatedPeerDetail)
var signalDBusAuthObserverAuthorizeAuthenticatedPeerLock sync.RWMutex

// DBusAuthObserverSignalAuthorizeAuthenticatedPeerCallback is a callback function for a 'authorize-authenticated-peer' signal emitted from a DBusAuthObserver.
type DBusAuthObserverSignalAuthorizeAuthenticatedPeerCallback func(stream *IOStream, credentials *Credentials) bool

/*
ConnectAuthorizeAuthenticatedPeer connects the callback to the 'authorize-authenticated-peer' signal for the DBusAuthObserver.

The returned value represents the connection, and may be passed to DisconnectAuthorizeAuthenticatedPeer to remove it.
*/
func (recv *DBusAuthObserver) ConnectAuthorizeAuthenticatedPeer(callback DBusAuthObserverSignalAuthorizeAuthenticatedPeerCallback) int {
	signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.Lock()
	defer signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.Unlock()

	signalDBusAuthObserverAuthorizeAuthenticatedPeerId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusAuthObserver_signal_connect_authorize_authenticated_peer(instance, C.gpointer(uintptr(signalDBusAuthObserverAuthorizeAuthenticatedPeerId)))

	detail := signalDBusAuthObserverAuthorizeAuthenticatedPeerDetail{callback, handlerID}
	signalDBusAuthObserverAuthorizeAuthenticatedPeerMap[signalDBusAuthObserverAuthorizeAuthenticatedPeerId] = detail

	return signalDBusAuthObserverAuthorizeAuthenticatedPeerId
}

/*
DisconnectAuthorizeAuthenticatedPeer disconnects a callback from the 'authorize-authenticated-peer' signal for the DBusAuthObserver.

The connectionID should be a value returned from a call to ConnectAuthorizeAuthenticatedPeer.
*/
func (recv *DBusAuthObserver) DisconnectAuthorizeAuthenticatedPeer(connectionID int) {
	signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.Lock()
	defer signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.Unlock()

	detail, exists := signalDBusAuthObserverAuthorizeAuthenticatedPeerMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusAuthObserverAuthorizeAuthenticatedPeerMap, connectionID)
}

//export dbusauthobserver_authorizeAuthenticatedPeerHandler
func dbusauthobserver_authorizeAuthenticatedPeerHandler(_ *C.GObject, c_stream *C.GIOStream, c_credentials *C.GCredentials, data C.gpointer) C.gboolean {
	signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.RLock()
	defer signalDBusAuthObserverAuthorizeAuthenticatedPeerLock.RUnlock()

	stream := IOStreamNewFromC(unsafe.Pointer(c_stream))

	credentials := CredentialsNewFromC(unsafe.Pointer(c_credentials))

	index := int(uintptr(data))
	callback := signalDBusAuthObserverAuthorizeAuthenticatedPeerMap[index].callback
	retGo := callback(stream, credentials)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dbus_auth_observer_new

// Blacklisted : g_dbus_auth_observer_authorize_authenticated_peer

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native *C.GDBusConnection
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	c := (*C.GDBusConnection)(u)
	if c == nil {
		return nil
	}

	g := &DBusConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusConnection with another DBusConnection, and returns true if they represent the same GObject.
func (recv *DBusConnection) Equals(other *DBusConnection) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusConnection.
func CastToDBusConnection(object *gobject.Object) *DBusConnection {
	return DBusConnectionNewFromC(object.ToC())
}

type signalDBusConnectionClosedDetail struct {
	callback  DBusConnectionSignalClosedCallback
	handlerID C.gulong
}

var signalDBusConnectionClosedId int
var signalDBusConnectionClosedMap = make(map[int]signalDBusConnectionClosedDetail)
var signalDBusConnectionClosedLock sync.RWMutex

// DBusConnectionSignalClosedCallback is a callback function for a 'closed' signal emitted from a DBusConnection.
type DBusConnectionSignalClosedCallback func(remotePeerVanished bool, error *glib.Error)

/*
ConnectClosed connects the callback to the 'closed' signal for the DBusConnection.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *DBusConnection) ConnectClosed(callback DBusConnectionSignalClosedCallback) int {
	signalDBusConnectionClosedLock.Lock()
	defer signalDBusConnectionClosedLock.Unlock()

	signalDBusConnectionClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusConnection_signal_connect_closed(instance, C.gpointer(uintptr(signalDBusConnectionClosedId)))

	detail := signalDBusConnectionClosedDetail{callback, handlerID}
	signalDBusConnectionClosedMap[signalDBusConnectionClosedId] = detail

	return signalDBusConnectionClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the DBusConnection.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *DBusConnection) DisconnectClosed(connectionID int) {
	signalDBusConnectionClosedLock.Lock()
	defer signalDBusConnectionClosedLock.Unlock()

	detail, exists := signalDBusConnectionClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusConnectionClosedMap, connectionID)
}

//export dbusconnection_closedHandler
func dbusconnection_closedHandler(_ *C.GObject, c_remote_peer_vanished C.gboolean, c_error *C.GError, data C.gpointer) {
	signalDBusConnectionClosedLock.RLock()
	defer signalDBusConnectionClosedLock.RUnlock()

	remotePeerVanished := c_remote_peer_vanished == C.TRUE

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	index := int(uintptr(data))
	callback := signalDBusConnectionClosedMap[index].callback
	callback(remotePeerVanished, error)
}

// Blacklisted : g_dbus_connection_new_finish

// Blacklisted : g_dbus_connection_new_for_address_finish

// Blacklisted : g_dbus_connection_new_for_address_sync

// Blacklisted : g_dbus_connection_new_sync

// g_dbus_connection_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_dbus_connection_new_for_address : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction (GDBusMessageFilterFunction) for param filter_function

// Unsupported : g_dbus_connection_call : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_connection_call_finish

// Blacklisted : g_dbus_connection_call_sync

// Unsupported : g_dbus_connection_close : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_connection_close_finish

// Blacklisted : g_dbus_connection_close_sync

// Blacklisted : g_dbus_connection_emit_signal

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_connection_flush_finish

// Blacklisted : g_dbus_connection_flush_sync

// Blacklisted : g_dbus_connection_get_capabilities

// Blacklisted : g_dbus_connection_get_exit_on_close

// Blacklisted : g_dbus_connection_get_guid

// Blacklisted : g_dbus_connection_get_peer_credentials

// Blacklisted : g_dbus_connection_get_stream

// Blacklisted : g_dbus_connection_get_unique_name

// Blacklisted : g_dbus_connection_is_closed

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Blacklisted : g_dbus_connection_remove_filter

// Blacklisted : g_dbus_connection_send_message

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_connection_send_message_with_reply_finish

// Blacklisted : g_dbus_connection_send_message_with_reply_sync

// Blacklisted : g_dbus_connection_set_exit_on_close

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback (GDBusSignalCallback) for param callback

// Blacklisted : g_dbus_connection_signal_unsubscribe

// Blacklisted : g_dbus_connection_start_message_processing

// Blacklisted : g_dbus_connection_unregister_object

// Blacklisted : g_dbus_connection_unregister_subtree

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native *C.GDBusMessage
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	c := (*C.GDBusMessage)(u)
	if c == nil {
		return nil
	}

	g := &DBusMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusMessage with another DBusMessage, and returns true if they represent the same GObject.
func (recv *DBusMessage) Equals(other *DBusMessage) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMessage.
func CastToDBusMessage(object *gobject.Object) *DBusMessage {
	return DBusMessageNewFromC(object.ToC())
}

// Blacklisted : g_dbus_message_new

// Blacklisted : g_dbus_message_new_from_blob

// Blacklisted : g_dbus_message_new_method_call

// Blacklisted : g_dbus_message_new_signal

// Blacklisted : g_dbus_message_bytes_needed

// Blacklisted : g_dbus_message_copy

// Blacklisted : g_dbus_message_get_arg0

// Blacklisted : g_dbus_message_get_body

// Blacklisted : g_dbus_message_get_byte_order

// Blacklisted : g_dbus_message_get_destination

// Blacklisted : g_dbus_message_get_error_name

// Blacklisted : g_dbus_message_get_flags

// Blacklisted : g_dbus_message_get_header

// Unsupported : g_dbus_message_get_header_fields : array return type :

// Blacklisted : g_dbus_message_get_interface

// Blacklisted : g_dbus_message_get_locked

// Blacklisted : g_dbus_message_get_member

// Blacklisted : g_dbus_message_get_message_type

// Blacklisted : g_dbus_message_get_num_unix_fds

// Blacklisted : g_dbus_message_get_path

// Blacklisted : g_dbus_message_get_reply_serial

// Blacklisted : g_dbus_message_get_sender

// Blacklisted : g_dbus_message_get_serial

// Blacklisted : g_dbus_message_get_signature

// Blacklisted : g_dbus_message_get_unix_fd_list

// Blacklisted : g_dbus_message_lock

// Blacklisted : g_dbus_message_new_method_error

// Blacklisted : g_dbus_message_new_method_error_literal

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : g_dbus_message_new_method_reply

// Blacklisted : g_dbus_message_print

// Blacklisted : g_dbus_message_set_body

// Blacklisted : g_dbus_message_set_byte_order

// Blacklisted : g_dbus_message_set_destination

// Blacklisted : g_dbus_message_set_error_name

// Blacklisted : g_dbus_message_set_flags

// Blacklisted : g_dbus_message_set_header

// Blacklisted : g_dbus_message_set_interface

// Blacklisted : g_dbus_message_set_member

// Blacklisted : g_dbus_message_set_message_type

// Blacklisted : g_dbus_message_set_num_unix_fds

// Blacklisted : g_dbus_message_set_path

// Blacklisted : g_dbus_message_set_reply_serial

// Blacklisted : g_dbus_message_set_sender

// Blacklisted : g_dbus_message_set_serial

// Blacklisted : g_dbus_message_set_signature

// Blacklisted : g_dbus_message_set_unix_fd_list

// Unsupported : g_dbus_message_to_blob : array return type :

// Blacklisted : g_dbus_message_to_gerror

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native *C.GDBusMethodInvocation
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	c := (*C.GDBusMethodInvocation)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMethodInvocation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusMethodInvocation with another DBusMethodInvocation, and returns true if they represent the same GObject.
func (recv *DBusMethodInvocation) Equals(other *DBusMethodInvocation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusMethodInvocation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusMethodInvocation.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMethodInvocation.
func CastToDBusMethodInvocation(object *gobject.Object) *DBusMethodInvocation {
	return DBusMethodInvocationNewFromC(object.ToC())
}

// Blacklisted : g_dbus_method_invocation_get_connection

// Blacklisted : g_dbus_method_invocation_get_interface_name

// Blacklisted : g_dbus_method_invocation_get_message

// Blacklisted : g_dbus_method_invocation_get_method_info

// Blacklisted : g_dbus_method_invocation_get_method_name

// Blacklisted : g_dbus_method_invocation_get_object_path

// Blacklisted : g_dbus_method_invocation_get_parameters

// Blacklisted : g_dbus_method_invocation_get_sender

// Unsupported : g_dbus_method_invocation_get_user_data : no return generator

// Blacklisted : g_dbus_method_invocation_return_dbus_error

// Blacklisted : g_dbus_method_invocation_return_error

// Blacklisted : g_dbus_method_invocation_return_error_literal

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Blacklisted : g_dbus_method_invocation_return_gerror

// Blacklisted : g_dbus_method_invocation_return_value

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native *C.GDBusProxy
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	c := (*C.GDBusProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxy{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusProxy) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusProxy with another DBusProxy, and returns true if they represent the same GObject.
func (recv *DBusProxy) Equals(other *DBusProxy) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusProxy) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusProxy.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusProxy.
func CastToDBusProxy(object *gobject.Object) *DBusProxy {
	return DBusProxyNewFromC(object.ToC())
}

// Unsupported signal 'g-properties-changed' for DBusProxy : unsupported parameter invalidated_properties :

type signalDBusProxyGSignalDetail struct {
	callback  DBusProxySignalGSignalCallback
	handlerID C.gulong
}

var signalDBusProxyGSignalId int
var signalDBusProxyGSignalMap = make(map[int]signalDBusProxyGSignalDetail)
var signalDBusProxyGSignalLock sync.RWMutex

// DBusProxySignalGSignalCallback is a callback function for a 'g-signal' signal emitted from a DBusProxy.
type DBusProxySignalGSignalCallback func(senderName string, signalName string, parameters *glib.Variant)

/*
ConnectGSignal connects the callback to the 'g-signal' signal for the DBusProxy.

The returned value represents the connection, and may be passed to DisconnectGSignal to remove it.
*/
func (recv *DBusProxy) ConnectGSignal(callback DBusProxySignalGSignalCallback) int {
	signalDBusProxyGSignalLock.Lock()
	defer signalDBusProxyGSignalLock.Unlock()

	signalDBusProxyGSignalId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusProxy_signal_connect_g_signal(instance, C.gpointer(uintptr(signalDBusProxyGSignalId)))

	detail := signalDBusProxyGSignalDetail{callback, handlerID}
	signalDBusProxyGSignalMap[signalDBusProxyGSignalId] = detail

	return signalDBusProxyGSignalId
}

/*
DisconnectGSignal disconnects a callback from the 'g-signal' signal for the DBusProxy.

The connectionID should be a value returned from a call to ConnectGSignal.
*/
func (recv *DBusProxy) DisconnectGSignal(connectionID int) {
	signalDBusProxyGSignalLock.Lock()
	defer signalDBusProxyGSignalLock.Unlock()

	detail, exists := signalDBusProxyGSignalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusProxyGSignalMap, connectionID)
}

//export dbusproxy_gSignalHandler
func dbusproxy_gSignalHandler(_ *C.GObject, c_sender_name *C.gchar, c_signal_name *C.gchar, c_parameters *C.GVariant, data C.gpointer) {
	signalDBusProxyGSignalLock.RLock()
	defer signalDBusProxyGSignalLock.RUnlock()

	senderName := C.GoString(c_sender_name)

	signalName := C.GoString(c_signal_name)

	parameters := glib.VariantNewFromC(unsafe.Pointer(c_parameters))

	index := int(uintptr(data))
	callback := signalDBusProxyGSignalMap[index].callback
	callback(senderName, signalName, parameters)
}

// Blacklisted : g_dbus_proxy_new_finish

// Blacklisted : g_dbus_proxy_new_for_bus_finish

// Blacklisted : g_dbus_proxy_new_for_bus_sync

// Blacklisted : g_dbus_proxy_new_sync

// g_dbus_proxy_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_dbus_proxy_new_for_bus : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_dbus_proxy_call : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dbus_proxy_call_finish

// Blacklisted : g_dbus_proxy_call_sync

// Blacklisted : g_dbus_proxy_get_cached_property

// Blacklisted : g_dbus_proxy_get_cached_property_names

// Blacklisted : g_dbus_proxy_get_connection

// Blacklisted : g_dbus_proxy_get_default_timeout

// Blacklisted : g_dbus_proxy_get_flags

// Blacklisted : g_dbus_proxy_get_interface_info

// Blacklisted : g_dbus_proxy_get_interface_name

// Blacklisted : g_dbus_proxy_get_name

// Blacklisted : g_dbus_proxy_get_name_owner

// Blacklisted : g_dbus_proxy_get_object_path

// Blacklisted : g_dbus_proxy_set_cached_property

// Blacklisted : g_dbus_proxy_set_default_timeout

// Blacklisted : g_dbus_proxy_set_interface_info

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native *C.GDBusServer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	c := (*C.GDBusServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusServer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusServer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusServer with another DBusServer, and returns true if they represent the same GObject.
func (recv *DBusServer) Equals(other *DBusServer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusServer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusServer.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusServer.
func CastToDBusServer(object *gobject.Object) *DBusServer {
	return DBusServerNewFromC(object.ToC())
}

type signalDBusServerNewConnectionDetail struct {
	callback  DBusServerSignalNewConnectionCallback
	handlerID C.gulong
}

var signalDBusServerNewConnectionId int
var signalDBusServerNewConnectionMap = make(map[int]signalDBusServerNewConnectionDetail)
var signalDBusServerNewConnectionLock sync.RWMutex

// DBusServerSignalNewConnectionCallback is a callback function for a 'new-connection' signal emitted from a DBusServer.
type DBusServerSignalNewConnectionCallback func(connection *DBusConnection) bool

/*
ConnectNewConnection connects the callback to the 'new-connection' signal for the DBusServer.

The returned value represents the connection, and may be passed to DisconnectNewConnection to remove it.
*/
func (recv *DBusServer) ConnectNewConnection(callback DBusServerSignalNewConnectionCallback) int {
	signalDBusServerNewConnectionLock.Lock()
	defer signalDBusServerNewConnectionLock.Unlock()

	signalDBusServerNewConnectionId++
	instance := C.gpointer(recv.native)
	handlerID := C.DBusServer_signal_connect_new_connection(instance, C.gpointer(uintptr(signalDBusServerNewConnectionId)))

	detail := signalDBusServerNewConnectionDetail{callback, handlerID}
	signalDBusServerNewConnectionMap[signalDBusServerNewConnectionId] = detail

	return signalDBusServerNewConnectionId
}

/*
DisconnectNewConnection disconnects a callback from the 'new-connection' signal for the DBusServer.

The connectionID should be a value returned from a call to ConnectNewConnection.
*/
func (recv *DBusServer) DisconnectNewConnection(connectionID int) {
	signalDBusServerNewConnectionLock.Lock()
	defer signalDBusServerNewConnectionLock.Unlock()

	detail, exists := signalDBusServerNewConnectionMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDBusServerNewConnectionMap, connectionID)
}

//export dbusserver_newConnectionHandler
func dbusserver_newConnectionHandler(_ *C.GObject, c_connection *C.GDBusConnection, data C.gpointer) C.gboolean {
	signalDBusServerNewConnectionLock.RLock()
	defer signalDBusServerNewConnectionLock.RUnlock()

	connection := DBusConnectionNewFromC(unsafe.Pointer(c_connection))

	index := int(uintptr(data))
	callback := signalDBusServerNewConnectionMap[index].callback
	retGo := callback(connection)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dbus_server_new_sync

// Blacklisted : g_dbus_server_get_client_address

// Blacklisted : g_dbus_server_get_flags

// Blacklisted : g_dbus_server_get_guid

// Blacklisted : g_dbus_server_is_active

// Blacklisted : g_dbus_server_start

// Blacklisted : g_dbus_server_stop

// Blacklisted : g_data_input_stream_read_upto

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_memory_output_stream_steal_data : no return generator

// Blacklisted : g_network_address_parse_uri

// Blacklisted : g_network_address_get_scheme

// Blacklisted : g_network_service_get_scheme

// Blacklisted : g_network_service_set_scheme

// Blacklisted : g_permission_acquire

// Unsupported : g_permission_acquire_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_permission_acquire_finish

// Blacklisted : g_permission_get_allowed

// Blacklisted : g_permission_get_can_acquire

// Blacklisted : g_permission_get_can_release

// Blacklisted : g_permission_impl_update

// Blacklisted : g_permission_release

// Unsupported : g_permission_release_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_permission_release_finish

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	c := (*C.GProxyAddress)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyAddress with another ProxyAddress, and returns true if they represent the same GObject.
func (recv *ProxyAddress) Equals(other *ProxyAddress) bool {
	return other.ToC() == recv.ToC()
}

// InetSocketAddress upcasts to *InetSocketAddress
func (recv *ProxyAddress) InetSocketAddress() *InetSocketAddress {
	return InetSocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// SocketAddress upcasts to *SocketAddress
func (recv *ProxyAddress) SocketAddress() *SocketAddress {
	return recv.InetSocketAddress().SocketAddress()
}

// Object upcasts to *Object
func (recv *ProxyAddress) Object() *gobject.Object {
	return recv.InetSocketAddress().Object()
}

// CastToWidget down casts any arbitrary Object to ProxyAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProxyAddress.
func CastToProxyAddress(object *gobject.Object) *ProxyAddress {
	return ProxyAddressNewFromC(object.ToC())
}

// Blacklisted : g_proxy_address_new

// Blacklisted : g_proxy_address_get_destination_hostname

// Blacklisted : g_proxy_address_get_destination_port

// Blacklisted : g_proxy_address_get_password

// Blacklisted : g_proxy_address_get_protocol

// Blacklisted : g_proxy_address_get_username

// Blacklisted : g_settings_new

// Blacklisted : g_settings_new_with_backend

// Blacklisted : g_settings_new_with_backend_and_path

// Blacklisted : g_settings_new_with_path

// Blacklisted : g_settings_list_schemas

// Blacklisted : g_settings_unbind

// Blacklisted : g_settings_bind

// Unsupported : g_settings_bind_with_mapping : unsupported parameter get_mapping : no type generator for SettingsBindGetMapping (GSettingsBindGetMapping) for param get_mapping

// Blacklisted : g_settings_bind_writable

// Blacklisted : g_settings_delay

// Blacklisted : g_settings_get

// Blacklisted : g_settings_get_boolean

// Blacklisted : g_settings_get_child

// Blacklisted : g_settings_get_double

// Blacklisted : g_settings_get_enum

// Blacklisted : g_settings_get_flags

// Blacklisted : g_settings_get_has_unapplied

// Blacklisted : g_settings_get_int

// Blacklisted : g_settings_get_string

// Blacklisted : g_settings_get_strv

// Blacklisted : g_settings_get_value

// Blacklisted : g_settings_is_writable

// Blacklisted : g_settings_set

// Blacklisted : g_settings_set_boolean

// Blacklisted : g_settings_set_double

// Blacklisted : g_settings_set_int

// Blacklisted : g_settings_set_string

// Blacklisted : g_settings_set_strv

// Blacklisted : g_settings_set_value

// g_settings_backend_flatten_tree : unsupported parameter keys : output array param keys
// Unsupported : g_settings_backend_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_changed_tree : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_keys_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_path_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Blacklisted : g_settings_backend_path_writable_changed

// Blacklisted : g_settings_backend_writable_changed

// Blacklisted : g_simple_permission_new

// Blacklisted : g_socket_get_credentials

// Blacklisted : g_socket_get_timeout

// Blacklisted : g_socket_receive_with_blocking

// Blacklisted : g_socket_send_with_blocking

// Blacklisted : g_socket_set_timeout

// Blacklisted : g_socket_client_connect_to_uri

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_client_connect_to_uri_finish

// Blacklisted : g_socket_client_get_enable_proxy

// Blacklisted : g_socket_client_get_timeout

// Blacklisted : g_socket_client_set_enable_proxy

// Blacklisted : g_socket_client_set_timeout

// Blacklisted : g_unix_connection_receive_credentials

// Blacklisted : g_unix_connection_send_credentials

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	c := (*C.GUnixCredentialsMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixCredentialsMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixCredentialsMessage with another UnixCredentialsMessage, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessage) Equals(other *UnixCredentialsMessage) bool {
	return other.ToC() == recv.ToC()
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixCredentialsMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixCredentialsMessage) Object() *gobject.Object {
	return recv.SocketControlMessage().Object()
}

// CastToWidget down casts any arbitrary Object to UnixCredentialsMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixCredentialsMessage.
func CastToUnixCredentialsMessage(object *gobject.Object) *UnixCredentialsMessage {
	return UnixCredentialsMessageNewFromC(object.ToC())
}

// Blacklisted : g_unix_credentials_message_new

// Blacklisted : g_unix_credentials_message_new_with_credentials

// Blacklisted : g_unix_credentials_message_is_supported

// Blacklisted : g_unix_credentials_message_get_credentials

// Blacklisted : g_unix_socket_address_new_with_type

// Blacklisted : g_unix_socket_address_get_address_type

// Blacklisted : g_zlib_compressor_get_file_info

// Blacklisted : g_zlib_compressor_set_file_info

// Blacklisted : g_zlib_decompressor_get_file_info
