// Code generated - DO NOT EDIT.
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"fmt"
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

	static GDBusMessage* _g_dbus_message_new_method_error(GDBusMessage* method_call_message, const gchar* error_name, const gchar* error_message_format) {
		return g_dbus_message_new_method_error(method_call_message, error_name, error_message_format);
    }
*/
/*

	static void _g_dbus_method_invocation_return_error(GDBusMethodInvocation* invocation, GQuark domain, gint code, const gchar* format) {
		return g_dbus_method_invocation_return_error(invocation, domain, code, format);
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
/*

	static void _g_settings_get(GSettings* settings, const gchar* key, const gchar* format) {
		return g_settings_get(settings, key, format);
    }
*/
/*

	static gboolean _g_settings_set(GSettings* settings, const gchar* key, const gchar* format) {
		return g_settings_set(settings, key, format);
    }
*/
/*

	static void _g_dbus_error_set_dbus_error(GError** error, const gchar* dbus_error_name, const gchar* dbus_error_message, const gchar* format) {
		return g_dbus_error_set_dbus_error(error, dbus_error_name, dbus_error_message, format);
    }
*/
import "C"

type BusNameOwnerFlags C.GBusNameOwnerFlags

const (
	BUS_NAME_OWNER_FLAGS_NONE              BusNameOwnerFlags = 0
	BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT BusNameOwnerFlags = 1
	BUS_NAME_OWNER_FLAGS_REPLACE           BusNameOwnerFlags = 2
	BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE      BusNameOwnerFlags = 4
)

type BusNameWatcherFlags C.GBusNameWatcherFlags

const (
	BUS_NAME_WATCHER_FLAGS_NONE       BusNameWatcherFlags = 0
	BUS_NAME_WATCHER_FLAGS_AUTO_START BusNameWatcherFlags = 1
)

type DBusCallFlags C.GDBusCallFlags

const (
	DBUS_CALL_FLAGS_NONE                            DBusCallFlags = 0
	DBUS_CALL_FLAGS_NO_AUTO_START                   DBusCallFlags = 1
	DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusCallFlags = 2
)

type DBusCapabilityFlags C.GDBusCapabilityFlags

const (
	DBUS_CAPABILITY_FLAGS_NONE            DBusCapabilityFlags = 0
	DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING DBusCapabilityFlags = 1
)

type DBusConnectionFlags C.GDBusConnectionFlags

const (
	DBUS_CONNECTION_FLAGS_NONE                           DBusConnectionFlags = 0
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT          DBusConnectionFlags = 1
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER          DBusConnectionFlags = 2
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusConnectionFlags = 4
	DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION         DBusConnectionFlags = 8
	DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING       DBusConnectionFlags = 16
)

type DBusMessageFlags C.GDBusMessageFlags

const (
	DBUS_MESSAGE_FLAGS_NONE                            DBusMessageFlags = 0
	DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED               DBusMessageFlags = 1
	DBUS_MESSAGE_FLAGS_NO_AUTO_START                   DBusMessageFlags = 2
	DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusMessageFlags = 4
)

type DBusPropertyInfoFlags C.GDBusPropertyInfoFlags

const (
	DBUS_PROPERTY_INFO_FLAGS_NONE     DBusPropertyInfoFlags = 0
	DBUS_PROPERTY_INFO_FLAGS_READABLE DBusPropertyInfoFlags = 1
	DBUS_PROPERTY_INFO_FLAGS_WRITABLE DBusPropertyInfoFlags = 2
)

type DBusProxyFlags C.GDBusProxyFlags

const (
	DBUS_PROXY_FLAGS_NONE                              DBusProxyFlags = 0
	DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES            DBusProxyFlags = 1
	DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS            DBusProxyFlags = 2
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START                 DBusProxyFlags = 4
	DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES        DBusProxyFlags = 8
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION DBusProxyFlags = 16
)

type DBusSendMessageFlags C.GDBusSendMessageFlags

const (
	DBUS_SEND_MESSAGE_FLAGS_NONE            DBusSendMessageFlags = 0
	DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL DBusSendMessageFlags = 1
)

type DBusServerFlags C.GDBusServerFlags

const (
	DBUS_SERVER_FLAGS_NONE                           DBusServerFlags = 0
	DBUS_SERVER_FLAGS_RUN_IN_THREAD                  DBusServerFlags = 1
	DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusServerFlags = 2
)

type DBusSignalFlags C.GDBusSignalFlags

const (
	DBUS_SIGNAL_FLAGS_NONE                 DBusSignalFlags = 0
	DBUS_SIGNAL_FLAGS_NO_MATCH_RULE        DBusSignalFlags = 1
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE DBusSignalFlags = 2
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH      DBusSignalFlags = 4
)

type DBusSubtreeFlags C.GDBusSubtreeFlags

const (
	DBUS_SUBTREE_FLAGS_NONE                           DBusSubtreeFlags = 0
	DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES DBusSubtreeFlags = 1
)

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

// CredentialsNew is a wrapper around the C function g_credentials_new.
func CredentialsNew() *Credentials {
	retC := C.g_credentials_new()
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_credentials_get_native : no return generator

// GetUnixUser is a wrapper around the C function g_credentials_get_unix_user.
func (recv *Credentials) GetUnixUser() (uint32, error) {
	var cThrowableError *C.GError

	retC := C.g_credentials_get_unix_user((*C.GCredentials)(recv.native), &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// IsSameUser is a wrapper around the C function g_credentials_is_same_user.
func (recv *Credentials) IsSameUser(otherCredentials *Credentials) (bool, error) {
	c_other_credentials := (*C.GCredentials)(C.NULL)
	if otherCredentials != nil {
		c_other_credentials = (*C.GCredentials)(otherCredentials.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_credentials_is_same_user((*C.GCredentials)(recv.native), c_other_credentials, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_credentials_set_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// SetUnixUser is a wrapper around the C function g_credentials_set_unix_user.
func (recv *Credentials) SetUnixUser(uid uint32) (bool, error) {
	c_uid := (C.uid_t)(uid)

	var cThrowableError *C.GError

	retC := C.g_credentials_set_unix_user((*C.GCredentials)(recv.native), c_uid, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToString is a wrapper around the C function g_credentials_to_string.
func (recv *Credentials) ToString() string {
	retC := C.g_credentials_to_string((*C.GCredentials)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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

// DBusAuthObserverNew is a wrapper around the C function g_dbus_auth_observer_new.
func DBusAuthObserverNew() *DBusAuthObserver {
	retC := C.g_dbus_auth_observer_new()
	retGo := DBusAuthObserverNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AuthorizeAuthenticatedPeer is a wrapper around the C function g_dbus_auth_observer_authorize_authenticated_peer.
func (recv *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) bool {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_credentials := (*C.GCredentials)(C.NULL)
	if credentials != nil {
		c_credentials = (*C.GCredentials)(credentials.ToC())
	}

	retC := C.g_dbus_auth_observer_authorize_authenticated_peer((*C.GDBusAuthObserver)(recv.native), c_stream, c_credentials)
	retGo := retC == C.TRUE

	return retGo
}

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

// DBusConnectionNewFinish is a wrapper around the C function g_dbus_connection_new_finish.
func DBusConnectionNewFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewForAddressFinish is a wrapper around the C function g_dbus_connection_new_for_address_finish.
func DBusConnectionNewForAddressFinish(res *AsyncResult) (*DBusConnection, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_finish(c_res, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewForAddressSync is a wrapper around the C function g_dbus_connection_new_for_address_sync.
func DBusConnectionNewForAddressSync(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_sync(c_address, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusConnectionNewSync is a wrapper around the C function g_dbus_connection_new_sync.
func DBusConnectionNewSync(stream *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_sync(c_stream, c_guid, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// g_dbus_connection_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_dbus_connection_new_for_address : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction (GDBusMessageFilterFunction) for param filter_function

// Unsupported : g_dbus_connection_call : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CallFinish is a wrapper around the C function g_dbus_connection_call_finish.
func (recv *DBusConnection) CallFinish(res *AsyncResult) (*glib.Variant, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_call_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CallSync is a wrapper around the C function g_dbus_connection_call_sync.
func (recv *DBusConnection) CallSync(busName string, objectPath string, interfaceName string, methodName string, parameters *glib.Variant, replyType *glib.VariantType, flags DBusCallFlags, timeoutMsec int32, cancellable *Cancellable) (*glib.Variant, error) {
	c_bus_name := C.CString(busName)
	defer C.free(unsafe.Pointer(c_bus_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_method_name := C.CString(methodName)
	defer C.free(unsafe.Pointer(c_method_name))

	c_parameters := (*C.GVariant)(C.NULL)
	if parameters != nil {
		c_parameters = (*C.GVariant)(parameters.ToC())
	}

	c_reply_type := (*C.GVariantType)(C.NULL)
	if replyType != nil {
		c_reply_type = (*C.GVariantType)(replyType.ToC())
	}

	c_flags := (C.GDBusCallFlags)(flags)

	c_timeout_msec := (C.gint)(timeoutMsec)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_call_sync((*C.GDBusConnection)(recv.native), c_bus_name, c_object_path, c_interface_name, c_method_name, c_parameters, c_reply_type, c_flags, c_timeout_msec, c_cancellable, &cThrowableError)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dbus_connection_close : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_dbus_connection_close_finish.
func (recv *DBusConnection) CloseFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_close_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CloseSync is a wrapper around the C function g_dbus_connection_close_sync.
func (recv *DBusConnection) CloseSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_close_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EmitSignal is a wrapper around the C function g_dbus_connection_emit_signal.
func (recv *DBusConnection) EmitSignal(destinationBusName string, objectPath string, interfaceName string, signalName string, parameters *glib.Variant) (bool, error) {
	c_destination_bus_name := C.CString(destinationBusName)
	defer C.free(unsafe.Pointer(c_destination_bus_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_signal_name := C.CString(signalName)
	defer C.free(unsafe.Pointer(c_signal_name))

	c_parameters := (*C.GVariant)(C.NULL)
	if parameters != nil {
		c_parameters = (*C.GVariant)(parameters.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_emit_signal((*C.GDBusConnection)(recv.native), c_destination_bus_name, c_object_path, c_interface_name, c_signal_name, c_parameters, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// FlushFinish is a wrapper around the C function g_dbus_connection_flush_finish.
func (recv *DBusConnection) FlushFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_flush_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// FlushSync is a wrapper around the C function g_dbus_connection_flush_sync.
func (recv *DBusConnection) FlushSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_flush_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetCapabilities is a wrapper around the C function g_dbus_connection_get_capabilities.
func (recv *DBusConnection) GetCapabilities() DBusCapabilityFlags {
	retC := C.g_dbus_connection_get_capabilities((*C.GDBusConnection)(recv.native))
	retGo := (DBusCapabilityFlags)(retC)

	return retGo
}

// GetExitOnClose is a wrapper around the C function g_dbus_connection_get_exit_on_close.
func (recv *DBusConnection) GetExitOnClose() bool {
	retC := C.g_dbus_connection_get_exit_on_close((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGuid is a wrapper around the C function g_dbus_connection_get_guid.
func (recv *DBusConnection) GetGuid() string {
	retC := C.g_dbus_connection_get_guid((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPeerCredentials is a wrapper around the C function g_dbus_connection_get_peer_credentials.
func (recv *DBusConnection) GetPeerCredentials() *Credentials {
	retC := C.g_dbus_connection_get_peer_credentials((*C.GDBusConnection)(recv.native))
	var retGo (*Credentials)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CredentialsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStream is a wrapper around the C function g_dbus_connection_get_stream.
func (recv *DBusConnection) GetStream() *IOStream {
	retC := C.g_dbus_connection_get_stream((*C.GDBusConnection)(recv.native))
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUniqueName is a wrapper around the C function g_dbus_connection_get_unique_name.
func (recv *DBusConnection) GetUniqueName() string {
	retC := C.g_dbus_connection_get_unique_name((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsClosed is a wrapper around the C function g_dbus_connection_is_closed.
func (recv *DBusConnection) IsClosed() bool {
	retC := C.g_dbus_connection_is_closed((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// RemoveFilter is a wrapper around the C function g_dbus_connection_remove_filter.
func (recv *DBusConnection) RemoveFilter(filterId uint32) {
	c_filter_id := (C.guint)(filterId)

	C.g_dbus_connection_remove_filter((*C.GDBusConnection)(recv.native), c_filter_id)

	return
}

// SendMessage is a wrapper around the C function g_dbus_connection_send_message.
func (recv *DBusConnection) SendMessage(message *DBusMessage, flags DBusSendMessageFlags) (bool, uint32, error) {
	c_message := (*C.GDBusMessage)(C.NULL)
	if message != nil {
		c_message = (*C.GDBusMessage)(message.ToC())
	}

	c_flags := (C.GDBusSendMessageFlags)(flags)

	var c_out_serial C.guint32

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message((*C.GDBusConnection)(recv.native), c_message, c_flags, &c_out_serial, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outSerial := (uint32)(c_out_serial)

	return retGo, outSerial, goError
}

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SendMessageWithReplyFinish is a wrapper around the C function g_dbus_connection_send_message_with_reply_finish.
func (recv *DBusConnection) SendMessageWithReplyFinish(res *AsyncResult) (*DBusMessage, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message_with_reply_finish((*C.GDBusConnection)(recv.native), c_res, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SendMessageWithReplySync is a wrapper around the C function g_dbus_connection_send_message_with_reply_sync.
func (recv *DBusConnection) SendMessageWithReplySync(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int32, cancellable *Cancellable) (*DBusMessage, uint32, error) {
	c_message := (*C.GDBusMessage)(C.NULL)
	if message != nil {
		c_message = (*C.GDBusMessage)(message.ToC())
	}

	c_flags := (C.GDBusSendMessageFlags)(flags)

	c_timeout_msec := (C.gint)(timeoutMsec)

	var c_out_serial C.guint32

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message_with_reply_sync((*C.GDBusConnection)(recv.native), c_message, c_flags, c_timeout_msec, &c_out_serial, c_cancellable, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	outSerial := (uint32)(c_out_serial)

	return retGo, outSerial, goError
}

// SetExitOnClose is a wrapper around the C function g_dbus_connection_set_exit_on_close.
func (recv *DBusConnection) SetExitOnClose(exitOnClose bool) {
	c_exit_on_close :=
		boolToGboolean(exitOnClose)

	C.g_dbus_connection_set_exit_on_close((*C.GDBusConnection)(recv.native), c_exit_on_close)

	return
}

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback (GDBusSignalCallback) for param callback

// SignalUnsubscribe is a wrapper around the C function g_dbus_connection_signal_unsubscribe.
func (recv *DBusConnection) SignalUnsubscribe(subscriptionId uint32) {
	c_subscription_id := (C.guint)(subscriptionId)

	C.g_dbus_connection_signal_unsubscribe((*C.GDBusConnection)(recv.native), c_subscription_id)

	return
}

// StartMessageProcessing is a wrapper around the C function g_dbus_connection_start_message_processing.
func (recv *DBusConnection) StartMessageProcessing() {
	C.g_dbus_connection_start_message_processing((*C.GDBusConnection)(recv.native))

	return
}

// UnregisterObject is a wrapper around the C function g_dbus_connection_unregister_object.
func (recv *DBusConnection) UnregisterObject(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_object((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

// UnregisterSubtree is a wrapper around the C function g_dbus_connection_unregister_subtree.
func (recv *DBusConnection) UnregisterSubtree(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_subtree((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

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

// DBusMessageNew is a wrapper around the C function g_dbus_message_new.
func DBusMessageNew() *DBusMessage {
	retC := C.g_dbus_message_new()
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMessageNewFromBlob is a wrapper around the C function g_dbus_message_new_from_blob.
func DBusMessageNewFromBlob(blob []uint8, capabilities DBusCapabilityFlags) (*DBusMessage, error) {
	c_blob_array := make([]C.guint8, len(blob)+1, len(blob)+1)
	for i, item := range blob {
		c := (C.guint8)(item)
		c_blob_array[i] = c
	}
	c_blob_array[len(blob)] = 0
	c_blob_arrayPtr := &c_blob_array[0]
	c_blob := (*C.guchar)(unsafe.Pointer(c_blob_arrayPtr))

	c_blob_len := (C.gsize)(len(blob))

	c_capabilities := (C.GDBusCapabilityFlags)(capabilities)

	var cThrowableError *C.GError

	retC := C.g_dbus_message_new_from_blob(c_blob, c_blob_len, c_capabilities, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusMessageNewMethodCall is a wrapper around the C function g_dbus_message_new_method_call.
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) *DBusMessage {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	retC := C.g_dbus_message_new_method_call(c_name, c_path, c_interface_, c_method)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMessageNewSignal is a wrapper around the C function g_dbus_message_new_signal.
func DBusMessageNewSignal(path string, interface_ string, signal string) *DBusMessage {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_signal := C.CString(signal)
	defer C.free(unsafe.Pointer(c_signal))

	retC := C.g_dbus_message_new_signal(c_path, c_interface_, c_signal)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusMessageBytesNeeded is a wrapper around the C function g_dbus_message_bytes_needed.
func DBusMessageBytesNeeded(blob []uint8) (int64, error) {
	c_blob_array := make([]C.guint8, len(blob)+1, len(blob)+1)
	for i, item := range blob {
		c := (C.guint8)(item)
		c_blob_array[i] = c
	}
	c_blob_array[len(blob)] = 0
	c_blob_arrayPtr := &c_blob_array[0]
	c_blob := (*C.guchar)(unsafe.Pointer(c_blob_arrayPtr))

	c_blob_len := (C.gsize)(len(blob))

	var cThrowableError *C.GError

	retC := C.g_dbus_message_bytes_needed(c_blob, c_blob_len, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Copy is a wrapper around the C function g_dbus_message_copy.
func (recv *DBusMessage) Copy() (*DBusMessage, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_copy((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetArg0 is a wrapper around the C function g_dbus_message_get_arg0.
func (recv *DBusMessage) GetArg0() string {
	retC := C.g_dbus_message_get_arg0((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetBody is a wrapper around the C function g_dbus_message_get_body.
func (recv *DBusMessage) GetBody() *glib.Variant {
	retC := C.g_dbus_message_get_body((*C.GDBusMessage)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetByteOrder is a wrapper around the C function g_dbus_message_get_byte_order.
func (recv *DBusMessage) GetByteOrder() DBusMessageByteOrder {
	retC := C.g_dbus_message_get_byte_order((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageByteOrder)(retC)

	return retGo
}

// GetDestination is a wrapper around the C function g_dbus_message_get_destination.
func (recv *DBusMessage) GetDestination() string {
	retC := C.g_dbus_message_get_destination((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetErrorName is a wrapper around the C function g_dbus_message_get_error_name.
func (recv *DBusMessage) GetErrorName() string {
	retC := C.g_dbus_message_get_error_name((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_message_get_flags.
func (recv *DBusMessage) GetFlags() DBusMessageFlags {
	retC := C.g_dbus_message_get_flags((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageFlags)(retC)

	return retGo
}

// GetHeader is a wrapper around the C function g_dbus_message_get_header.
func (recv *DBusMessage) GetHeader(headerField DBusMessageHeaderField) *glib.Variant {
	c_header_field := (C.GDBusMessageHeaderField)(headerField)

	retC := C.g_dbus_message_get_header((*C.GDBusMessage)(recv.native), c_header_field)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_get_header_fields : array return type :

// GetInterface is a wrapper around the C function g_dbus_message_get_interface.
func (recv *DBusMessage) GetInterface() string {
	retC := C.g_dbus_message_get_interface((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLocked is a wrapper around the C function g_dbus_message_get_locked.
func (recv *DBusMessage) GetLocked() bool {
	retC := C.g_dbus_message_get_locked((*C.GDBusMessage)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMember is a wrapper around the C function g_dbus_message_get_member.
func (recv *DBusMessage) GetMember() string {
	retC := C.g_dbus_message_get_member((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMessageType is a wrapper around the C function g_dbus_message_get_message_type.
func (recv *DBusMessage) GetMessageType() DBusMessageType {
	retC := C.g_dbus_message_get_message_type((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageType)(retC)

	return retGo
}

// GetNumUnixFds is a wrapper around the C function g_dbus_message_get_num_unix_fds.
func (recv *DBusMessage) GetNumUnixFds() uint32 {
	retC := C.g_dbus_message_get_num_unix_fds((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetPath is a wrapper around the C function g_dbus_message_get_path.
func (recv *DBusMessage) GetPath() string {
	retC := C.g_dbus_message_get_path((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetReplySerial is a wrapper around the C function g_dbus_message_get_reply_serial.
func (recv *DBusMessage) GetReplySerial() uint32 {
	retC := C.g_dbus_message_get_reply_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSender is a wrapper around the C function g_dbus_message_get_sender.
func (recv *DBusMessage) GetSender() string {
	retC := C.g_dbus_message_get_sender((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSerial is a wrapper around the C function g_dbus_message_get_serial.
func (recv *DBusMessage) GetSerial() uint32 {
	retC := C.g_dbus_message_get_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSignature is a wrapper around the C function g_dbus_message_get_signature.
func (recv *DBusMessage) GetSignature() string {
	retC := C.g_dbus_message_get_signature((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUnixFdList is a wrapper around the C function g_dbus_message_get_unix_fd_list.
func (recv *DBusMessage) GetUnixFdList() *UnixFDList {
	retC := C.g_dbus_message_get_unix_fd_list((*C.GDBusMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lock is a wrapper around the C function g_dbus_message_lock.
func (recv *DBusMessage) Lock() {
	C.g_dbus_message_lock((*C.GDBusMessage)(recv.native))

	return
}

// NewMethodError is a wrapper around the C function g_dbus_message_new_method_error.
func (recv *DBusMessage) NewMethodError(errorName string, errorMessageFormat string, args ...interface{}) *DBusMessage {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	goFormattedString := fmt.Sprintf(errorMessageFormat, args...)
	c_error_message_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_error_message_format))

	retC := C._g_dbus_message_new_method_error((*C.GDBusMessage)(recv.native), c_error_name, c_error_message_format)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NewMethodErrorLiteral is a wrapper around the C function g_dbus_message_new_method_error_literal.
func (recv *DBusMessage) NewMethodErrorLiteral(errorName string, errorMessage string) *DBusMessage {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	retC := C.g_dbus_message_new_method_error_literal((*C.GDBusMessage)(recv.native), c_error_name, c_error_message)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// NewMethodReply is a wrapper around the C function g_dbus_message_new_method_reply.
func (recv *DBusMessage) NewMethodReply() *DBusMessage {
	retC := C.g_dbus_message_new_method_reply((*C.GDBusMessage)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Print is a wrapper around the C function g_dbus_message_print.
func (recv *DBusMessage) Print(indent uint32) string {
	c_indent := (C.guint)(indent)

	retC := C.g_dbus_message_print((*C.GDBusMessage)(recv.native), c_indent)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// SetBody is a wrapper around the C function g_dbus_message_set_body.
func (recv *DBusMessage) SetBody(body *glib.Variant) {
	c_body := (*C.GVariant)(C.NULL)
	if body != nil {
		c_body = (*C.GVariant)(body.ToC())
	}

	C.g_dbus_message_set_body((*C.GDBusMessage)(recv.native), c_body)

	return
}

// SetByteOrder is a wrapper around the C function g_dbus_message_set_byte_order.
func (recv *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	c_byte_order := (C.GDBusMessageByteOrder)(byteOrder)

	C.g_dbus_message_set_byte_order((*C.GDBusMessage)(recv.native), c_byte_order)

	return
}

// SetDestination is a wrapper around the C function g_dbus_message_set_destination.
func (recv *DBusMessage) SetDestination(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_destination((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetErrorName is a wrapper around the C function g_dbus_message_set_error_name.
func (recv *DBusMessage) SetErrorName(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_error_name((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetFlags is a wrapper around the C function g_dbus_message_set_flags.
func (recv *DBusMessage) SetFlags(flags DBusMessageFlags) {
	c_flags := (C.GDBusMessageFlags)(flags)

	C.g_dbus_message_set_flags((*C.GDBusMessage)(recv.native), c_flags)

	return
}

// SetHeader is a wrapper around the C function g_dbus_message_set_header.
func (recv *DBusMessage) SetHeader(headerField DBusMessageHeaderField, value *glib.Variant) {
	c_header_field := (C.GDBusMessageHeaderField)(headerField)

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_dbus_message_set_header((*C.GDBusMessage)(recv.native), c_header_field, c_value)

	return
}

// SetInterface is a wrapper around the C function g_dbus_message_set_interface.
func (recv *DBusMessage) SetInterface(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_interface((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetMember is a wrapper around the C function g_dbus_message_set_member.
func (recv *DBusMessage) SetMember(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_member((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetMessageType is a wrapper around the C function g_dbus_message_set_message_type.
func (recv *DBusMessage) SetMessageType(type_ DBusMessageType) {
	c_type := (C.GDBusMessageType)(type_)

	C.g_dbus_message_set_message_type((*C.GDBusMessage)(recv.native), c_type)

	return
}

// SetNumUnixFds is a wrapper around the C function g_dbus_message_set_num_unix_fds.
func (recv *DBusMessage) SetNumUnixFds(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_num_unix_fds((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetPath is a wrapper around the C function g_dbus_message_set_path.
func (recv *DBusMessage) SetPath(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_path((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetReplySerial is a wrapper around the C function g_dbus_message_set_reply_serial.
func (recv *DBusMessage) SetReplySerial(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_reply_serial((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetSender is a wrapper around the C function g_dbus_message_set_sender.
func (recv *DBusMessage) SetSender(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_sender((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetSerial is a wrapper around the C function g_dbus_message_set_serial.
func (recv *DBusMessage) SetSerial(serial uint32) {
	c_serial := (C.guint32)(serial)

	C.g_dbus_message_set_serial((*C.GDBusMessage)(recv.native), c_serial)

	return
}

// SetSignature is a wrapper around the C function g_dbus_message_set_signature.
func (recv *DBusMessage) SetSignature(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_signature((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetUnixFdList is a wrapper around the C function g_dbus_message_set_unix_fd_list.
func (recv *DBusMessage) SetUnixFdList(fdList *UnixFDList) {
	c_fd_list := (*C.GUnixFDList)(C.NULL)
	if fdList != nil {
		c_fd_list = (*C.GUnixFDList)(fdList.ToC())
	}

	C.g_dbus_message_set_unix_fd_list((*C.GDBusMessage)(recv.native), c_fd_list)

	return
}

// Unsupported : g_dbus_message_to_blob : array return type :

// ToGerror is a wrapper around the C function g_dbus_message_to_gerror.
func (recv *DBusMessage) ToGerror() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_to_gerror((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

// GetConnection is a wrapper around the C function g_dbus_method_invocation_get_connection.
func (recv *DBusMethodInvocation) GetConnection() *DBusConnection {
	retC := C.g_dbus_method_invocation_get_connection((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInterfaceName is a wrapper around the C function g_dbus_method_invocation_get_interface_name.
func (recv *DBusMethodInvocation) GetInterfaceName() string {
	retC := C.g_dbus_method_invocation_get_interface_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMessage is a wrapper around the C function g_dbus_method_invocation_get_message.
func (recv *DBusMethodInvocation) GetMessage() *DBusMessage {
	retC := C.g_dbus_method_invocation_get_message((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMethodInfo is a wrapper around the C function g_dbus_method_invocation_get_method_info.
func (recv *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	retC := C.g_dbus_method_invocation_get_method_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMethodName is a wrapper around the C function g_dbus_method_invocation_get_method_name.
func (recv *DBusMethodInvocation) GetMethodName() string {
	retC := C.g_dbus_method_invocation_get_method_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_method_invocation_get_object_path.
func (recv *DBusMethodInvocation) GetObjectPath() string {
	retC := C.g_dbus_method_invocation_get_object_path((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetParameters is a wrapper around the C function g_dbus_method_invocation_get_parameters.
func (recv *DBusMethodInvocation) GetParameters() *glib.Variant {
	retC := C.g_dbus_method_invocation_get_parameters((*C.GDBusMethodInvocation)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSender is a wrapper around the C function g_dbus_method_invocation_get_sender.
func (recv *DBusMethodInvocation) GetSender() string {
	retC := C.g_dbus_method_invocation_get_sender((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_method_invocation_get_user_data : no return generator

// ReturnDbusError is a wrapper around the C function g_dbus_method_invocation_return_dbus_error.
func (recv *DBusMethodInvocation) ReturnDbusError(errorName string, errorMessage string) {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	C.g_dbus_method_invocation_return_dbus_error((*C.GDBusMethodInvocation)(recv.native), c_error_name, c_error_message)

	return
}

// ReturnError is a wrapper around the C function g_dbus_method_invocation_return_error.
func (recv *DBusMethodInvocation) ReturnError(domain glib.Quark, code int32, format string, args ...interface{}) {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_dbus_method_invocation_return_error((*C.GDBusMethodInvocation)(recv.native), c_domain, c_code, c_format)

	return
}

// ReturnErrorLiteral is a wrapper around the C function g_dbus_method_invocation_return_error_literal.
func (recv *DBusMethodInvocation) ReturnErrorLiteral(domain glib.Quark, code int32, message string) {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_dbus_method_invocation_return_error_literal((*C.GDBusMethodInvocation)(recv.native), c_domain, c_code, c_message)

	return
}

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// ReturnGerror is a wrapper around the C function g_dbus_method_invocation_return_gerror.
func (recv *DBusMethodInvocation) ReturnGerror(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_dbus_method_invocation_return_gerror((*C.GDBusMethodInvocation)(recv.native), c_error)

	return
}

// ReturnValue is a wrapper around the C function g_dbus_method_invocation_return_value.
func (recv *DBusMethodInvocation) ReturnValue(parameters *glib.Variant) {
	c_parameters := (*C.GVariant)(C.NULL)
	if parameters != nil {
		c_parameters = (*C.GVariant)(parameters.ToC())
	}

	C.g_dbus_method_invocation_return_value((*C.GDBusMethodInvocation)(recv.native), c_parameters)

	return
}

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

// DBusProxyNewFinish is a wrapper around the C function g_dbus_proxy_new_finish.
func DBusProxyNewFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewForBusFinish is a wrapper around the C function g_dbus_proxy_new_for_bus_finish.
func DBusProxyNewForBusFinish(res *AsyncResult) (*DBusProxy, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_finish(c_res, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewForBusSync is a wrapper around the C function g_dbus_proxy_new_for_bus_sync.
func DBusProxyNewForBusSync(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_bus_type := (C.GBusType)(busType)

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_sync(c_bus_type, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// DBusProxyNewSync is a wrapper around the C function g_dbus_proxy_new_sync.
func DBusProxyNewSync(connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_sync(c_connection, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// g_dbus_proxy_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_dbus_proxy_new_for_bus : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_dbus_proxy_call : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CallFinish is a wrapper around the C function g_dbus_proxy_call_finish.
func (recv *DBusProxy) CallFinish(res *AsyncResult) (*glib.Variant, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_call_finish((*C.GDBusProxy)(recv.native), c_res, &cThrowableError)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CallSync is a wrapper around the C function g_dbus_proxy_call_sync.
func (recv *DBusProxy) CallSync(methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int32, cancellable *Cancellable) (*glib.Variant, error) {
	c_method_name := C.CString(methodName)
	defer C.free(unsafe.Pointer(c_method_name))

	c_parameters := (*C.GVariant)(C.NULL)
	if parameters != nil {
		c_parameters = (*C.GVariant)(parameters.ToC())
	}

	c_flags := (C.GDBusCallFlags)(flags)

	c_timeout_msec := (C.gint)(timeoutMsec)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_call_sync((*C.GDBusProxy)(recv.native), c_method_name, c_parameters, c_flags, c_timeout_msec, c_cancellable, &cThrowableError)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetCachedProperty is a wrapper around the C function g_dbus_proxy_get_cached_property.
func (recv *DBusProxy) GetCachedProperty(propertyName string) *glib.Variant {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.g_dbus_proxy_get_cached_property((*C.GDBusProxy)(recv.native), c_property_name)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetCachedPropertyNames is a wrapper around the C function g_dbus_proxy_get_cached_property_names.
func (recv *DBusProxy) GetCachedPropertyNames() []string {
	retC := C.g_dbus_proxy_get_cached_property_names((*C.GDBusProxy)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetConnection is a wrapper around the C function g_dbus_proxy_get_connection.
func (recv *DBusProxy) GetConnection() *DBusConnection {
	retC := C.g_dbus_proxy_get_connection((*C.GDBusProxy)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultTimeout is a wrapper around the C function g_dbus_proxy_get_default_timeout.
func (recv *DBusProxy) GetDefaultTimeout() int32 {
	retC := C.g_dbus_proxy_get_default_timeout((*C.GDBusProxy)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_proxy_get_flags.
func (recv *DBusProxy) GetFlags() DBusProxyFlags {
	retC := C.g_dbus_proxy_get_flags((*C.GDBusProxy)(recv.native))
	retGo := (DBusProxyFlags)(retC)

	return retGo
}

// GetInterfaceInfo is a wrapper around the C function g_dbus_proxy_get_interface_info.
func (recv *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_proxy_get_interface_info((*C.GDBusProxy)(recv.native))
	var retGo (*DBusInterfaceInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetInterfaceName is a wrapper around the C function g_dbus_proxy_get_interface_name.
func (recv *DBusProxy) GetInterfaceName() string {
	retC := C.g_dbus_proxy_get_interface_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_dbus_proxy_get_name.
func (recv *DBusProxy) GetName() string {
	retC := C.g_dbus_proxy_get_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNameOwner is a wrapper around the C function g_dbus_proxy_get_name_owner.
func (recv *DBusProxy) GetNameOwner() string {
	retC := C.g_dbus_proxy_get_name_owner((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_proxy_get_object_path.
func (recv *DBusProxy) GetObjectPath() string {
	retC := C.g_dbus_proxy_get_object_path((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetCachedProperty is a wrapper around the C function g_dbus_proxy_set_cached_property.
func (recv *DBusProxy) SetCachedProperty(propertyName string, value *glib.Variant) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_dbus_proxy_set_cached_property((*C.GDBusProxy)(recv.native), c_property_name, c_value)

	return
}

// SetDefaultTimeout is a wrapper around the C function g_dbus_proxy_set_default_timeout.
func (recv *DBusProxy) SetDefaultTimeout(timeoutMsec int32) {
	c_timeout_msec := (C.gint)(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout((*C.GDBusProxy)(recv.native), c_timeout_msec)

	return
}

// SetInterfaceInfo is a wrapper around the C function g_dbus_proxy_set_interface_info.
func (recv *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	c_info := (*C.GDBusInterfaceInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GDBusInterfaceInfo)(info.ToC())
	}

	C.g_dbus_proxy_set_interface_info((*C.GDBusProxy)(recv.native), c_info)

	return
}

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

// DBusServerNewSync is a wrapper around the C function g_dbus_server_new_sync.
func DBusServerNewSync(address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusServer, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusServerFlags)(flags)

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_observer := (*C.GDBusAuthObserver)(C.NULL)
	if observer != nil {
		c_observer = (*C.GDBusAuthObserver)(observer.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_server_new_sync(c_address, c_flags, c_guid, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusServerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetClientAddress is a wrapper around the C function g_dbus_server_get_client_address.
func (recv *DBusServer) GetClientAddress() string {
	retC := C.g_dbus_server_get_client_address((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_server_get_flags.
func (recv *DBusServer) GetFlags() DBusServerFlags {
	retC := C.g_dbus_server_get_flags((*C.GDBusServer)(recv.native))
	retGo := (DBusServerFlags)(retC)

	return retGo
}

// GetGuid is a wrapper around the C function g_dbus_server_get_guid.
func (recv *DBusServer) GetGuid() string {
	retC := C.g_dbus_server_get_guid((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsActive is a wrapper around the C function g_dbus_server_is_active.
func (recv *DBusServer) IsActive() bool {
	retC := C.g_dbus_server_is_active((*C.GDBusServer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Start is a wrapper around the C function g_dbus_server_start.
func (recv *DBusServer) Start() {
	C.g_dbus_server_start((*C.GDBusServer)(recv.native))

	return
}

// Stop is a wrapper around the C function g_dbus_server_stop.
func (recv *DBusServer) Stop() {
	C.g_dbus_server_stop((*C.GDBusServer)(recv.native))

	return
}

// ReadUpto is a wrapper around the C function g_data_input_stream_read_upto.
func (recv *DataInputStream) ReadUpto(stopChars string, stopCharsLen int64, cancellable *Cancellable) (string, uint64, error) {
	c_stop_chars := C.CString(stopChars)
	defer C.free(unsafe.Pointer(c_stop_chars))

	c_stop_chars_len := (C.gssize)(stopCharsLen)

	var c_length C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_upto((*C.GDataInputStream)(recv.native), c_stop_chars, c_stop_chars_len, &c_length, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_memory_output_stream_steal_data : no return generator

// NetworkAddressParseUri is a wrapper around the C function g_network_address_parse_uri.
func NetworkAddressParseUri(uri string, defaultPort uint16) (*NetworkAddress, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_default_port := (C.guint16)(defaultPort)

	var cThrowableError *C.GError

	retC := C.g_network_address_parse_uri(c_uri, c_default_port, &cThrowableError)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetScheme is a wrapper around the C function g_network_address_get_scheme.
func (recv *NetworkAddress) GetScheme() string {
	retC := C.g_network_address_get_scheme((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetScheme is a wrapper around the C function g_network_service_get_scheme.
func (recv *NetworkService) GetScheme() string {
	retC := C.g_network_service_get_scheme((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetScheme is a wrapper around the C function g_network_service_set_scheme.
func (recv *NetworkService) SetScheme(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.g_network_service_set_scheme((*C.GNetworkService)(recv.native), c_scheme)

	return
}

// Acquire is a wrapper around the C function g_permission_acquire.
func (recv *Permission) Acquire(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_permission_acquire((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_permission_acquire_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AcquireFinish is a wrapper around the C function g_permission_acquire_finish.
func (recv *Permission) AcquireFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_acquire_finish((*C.GPermission)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAllowed is a wrapper around the C function g_permission_get_allowed.
func (recv *Permission) GetAllowed() bool {
	retC := C.g_permission_get_allowed((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanAcquire is a wrapper around the C function g_permission_get_can_acquire.
func (recv *Permission) GetCanAcquire() bool {
	retC := C.g_permission_get_can_acquire((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanRelease is a wrapper around the C function g_permission_get_can_release.
func (recv *Permission) GetCanRelease() bool {
	retC := C.g_permission_get_can_release((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ImplUpdate is a wrapper around the C function g_permission_impl_update.
func (recv *Permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	c_allowed :=
		boolToGboolean(allowed)

	c_can_acquire :=
		boolToGboolean(canAcquire)

	c_can_release :=
		boolToGboolean(canRelease)

	C.g_permission_impl_update((*C.GPermission)(recv.native), c_allowed, c_can_acquire, c_can_release)

	return
}

// Release is a wrapper around the C function g_permission_release.
func (recv *Permission) Release(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_permission_release((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_permission_release_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReleaseFinish is a wrapper around the C function g_permission_release_finish.
func (recv *Permission) ReleaseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_release_finish((*C.GPermission)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

// ProxyAddressNew is a wrapper around the C function g_proxy_address_new.
func ProxyAddressNew(inetaddr *InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) *ProxyAddress {
	c_inetaddr := (*C.GInetAddress)(C.NULL)
	if inetaddr != nil {
		c_inetaddr = (*C.GInetAddress)(inetaddr.ToC())
	}

	c_port := (C.guint16)(port)

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_dest_hostname := C.CString(destHostname)
	defer C.free(unsafe.Pointer(c_dest_hostname))

	c_dest_port := (C.guint16)(destPort)

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.g_proxy_address_new(c_inetaddr, c_port, c_protocol, c_dest_hostname, c_dest_port, c_username, c_password)
	retGo := ProxyAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetDestinationHostname is a wrapper around the C function g_proxy_address_get_destination_hostname.
func (recv *ProxyAddress) GetDestinationHostname() string {
	retC := C.g_proxy_address_get_destination_hostname((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDestinationPort is a wrapper around the C function g_proxy_address_get_destination_port.
func (recv *ProxyAddress) GetDestinationPort() uint16 {
	retC := C.g_proxy_address_get_destination_port((*C.GProxyAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetPassword is a wrapper around the C function g_proxy_address_get_password.
func (recv *ProxyAddress) GetPassword() string {
	retC := C.g_proxy_address_get_password((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProtocol is a wrapper around the C function g_proxy_address_get_protocol.
func (recv *ProxyAddress) GetProtocol() string {
	retC := C.g_proxy_address_get_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUsername is a wrapper around the C function g_proxy_address_get_username.
func (recv *ProxyAddress) GetUsername() string {
	retC := C.g_proxy_address_get_username((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SettingsNew is a wrapper around the C function g_settings_new.
func SettingsNew(schemaId string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	retC := C.g_settings_new(c_schema_id)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithBackend is a wrapper around the C function g_settings_new_with_backend.
func SettingsNewWithBackend(schemaId string, backend *SettingsBackend) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	retC := C.g_settings_new_with_backend(c_schema_id, c_backend)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithBackendAndPath is a wrapper around the C function g_settings_new_with_backend_and_path.
func SettingsNewWithBackendAndPath(schemaId string, backend *SettingsBackend, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_backend_and_path(c_schema_id, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsNewWithPath is a wrapper around the C function g_settings_new_with_path.
func SettingsNewWithPath(schemaId string, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_path(c_schema_id, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SettingsListSchemas is a wrapper around the C function g_settings_list_schemas.
func SettingsListSchemas() []string {
	retC := C.g_settings_list_schemas()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// SettingsUnbind is a wrapper around the C function g_settings_unbind.
func SettingsUnbind(object *gobject.Object, property string) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	C.g_settings_unbind(c_object, c_property)

	return
}

// Bind is a wrapper around the C function g_settings_bind.
func (recv *Settings) Bind(key string, object *gobject.Object, property string, flags SettingsBindFlags) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_flags := (C.GSettingsBindFlags)(flags)

	C.g_settings_bind((*C.GSettings)(recv.native), c_key, c_object, c_property, c_flags)

	return
}

// Unsupported : g_settings_bind_with_mapping : unsupported parameter get_mapping : no type generator for SettingsBindGetMapping (GSettingsBindGetMapping) for param get_mapping

// BindWritable is a wrapper around the C function g_settings_bind_writable.
func (recv *Settings) BindWritable(key string, object *gobject.Object, property string, inverted bool) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_inverted :=
		boolToGboolean(inverted)

	C.g_settings_bind_writable((*C.GSettings)(recv.native), c_key, c_object, c_property, c_inverted)

	return
}

// Delay is a wrapper around the C function g_settings_delay.
func (recv *Settings) Delay() {
	C.g_settings_delay((*C.GSettings)(recv.native))

	return
}

// Get is a wrapper around the C function g_settings_get.
func (recv *Settings) Get(key string, format string, args ...interface{}) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_settings_get((*C.GSettings)(recv.native), c_key, c_format)

	return
}

// GetBoolean is a wrapper around the C function g_settings_get_boolean.
func (recv *Settings) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_boolean((*C.GSettings)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetChild is a wrapper around the C function g_settings_get_child.
func (recv *Settings) GetChild(name string) *Settings {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_get_child((*C.GSettings)(recv.native), c_name)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDouble is a wrapper around the C function g_settings_get_double.
func (recv *Settings) GetDouble(key string) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_double((*C.GSettings)(recv.native), c_key)
	retGo := (float64)(retC)

	return retGo
}

// GetEnum is a wrapper around the C function g_settings_get_enum.
func (recv *Settings) GetEnum(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_enum((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_settings_get_flags.
func (recv *Settings) GetFlags(key string) uint32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_flags((*C.GSettings)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// GetHasUnapplied is a wrapper around the C function g_settings_get_has_unapplied.
func (recv *Settings) GetHasUnapplied() bool {
	retC := C.g_settings_get_has_unapplied((*C.GSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInt is a wrapper around the C function g_settings_get_int.
func (recv *Settings) GetInt(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// GetString is a wrapper around the C function g_settings_get_string.
func (recv *Settings) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_string((*C.GSettings)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetStrv is a wrapper around the C function g_settings_get_strv.
func (recv *Settings) GetStrv(key string) []string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_strv((*C.GSettings)(recv.native), c_key)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetValue is a wrapper around the C function g_settings_get_value.
func (recv *Settings) GetValue(key string) *glib.Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_value((*C.GSettings)(recv.native), c_key)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsWritable is a wrapper around the C function g_settings_is_writable.
func (recv *Settings) IsWritable(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_is_writable((*C.GSettings)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Set is a wrapper around the C function g_settings_set.
func (recv *Settings) Set(key string, format string, args ...interface{}) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_settings_set((*C.GSettings)(recv.native), c_key, c_format)
	retGo := retC == C.TRUE

	return retGo
}

// SetBoolean is a wrapper around the C function g_settings_set_boolean.
func (recv *Settings) SetBoolean(key string, value bool) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	retC := C.g_settings_set_boolean((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetDouble is a wrapper around the C function g_settings_set_double.
func (recv *Settings) SetDouble(key string, value float64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	retC := C.g_settings_set_double((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetInt is a wrapper around the C function g_settings_set_int.
func (recv *Settings) SetInt(key string, value int32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	retC := C.g_settings_set_int((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetString is a wrapper around the C function g_settings_set_string.
func (recv *Settings) SetString(key string, value string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.g_settings_set_string((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetStrv is a wrapper around the C function g_settings_set_strv.
func (recv *Settings) SetStrv(key string, value []string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value_array := make([]*C.gchar, len(value)+1, len(value)+1)
	for i, item := range value {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_value_array[i] = c
	}
	c_value_array[len(value)] = nil
	c_value_arrayPtr := &c_value_array[0]
	c_value := (**C.gchar)(unsafe.Pointer(c_value_arrayPtr))

	retC := C.g_settings_set_strv((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetValue is a wrapper around the C function g_settings_set_value.
func (recv *Settings) SetValue(key string, value *glib.Variant) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_settings_set_value((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// g_settings_backend_flatten_tree : unsupported parameter keys : output array param keys
// Unsupported : g_settings_backend_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_changed_tree : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_keys_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Unsupported : g_settings_backend_path_changed : unsupported parameter origin_tag : no type generator for gpointer (gpointer) for param origin_tag

// Blacklisted : g_settings_backend_path_writable_changed

// Blacklisted : g_settings_backend_writable_changed

// SimplePermissionNew is a wrapper around the C function g_simple_permission_new.
func SimplePermissionNew(allowed bool) *SimplePermission {
	c_allowed :=
		boolToGboolean(allowed)

	retC := C.g_simple_permission_new(c_allowed)
	retGo := SimplePermissionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetCredentials is a wrapper around the C function g_socket_get_credentials.
func (recv *Socket) GetCredentials() (*Credentials, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_credentials((*C.GSocket)(recv.native), &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetTimeout is a wrapper around the C function g_socket_get_timeout.
func (recv *Socket) GetTimeout() uint32 {
	retC := C.g_socket_get_timeout((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// ReceiveWithBlocking is a wrapper around the C function g_socket_receive_with_blocking.
func (recv *Socket) ReceiveWithBlocking(buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive_with_blocking((*C.GSocket)(recv.native), c_buffer, c_size, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SendWithBlocking is a wrapper around the C function g_socket_send_with_blocking.
func (recv *Socket) SendWithBlocking(buffer []uint8, blocking bool, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_blocking :=
		boolToGboolean(blocking)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send_with_blocking((*C.GSocket)(recv.native), c_buffer, c_size, c_blocking, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetTimeout is a wrapper around the C function g_socket_set_timeout.
func (recv *Socket) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_set_timeout((*C.GSocket)(recv.native), c_timeout)

	return
}

// ConnectToUri is a wrapper around the C function g_socket_client_connect_to_uri.
func (recv *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_uri((*C.GSocketClient)(recv.native), c_uri, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectToUriFinish is a wrapper around the C function g_socket_client_connect_to_uri_finish.
func (recv *SocketClient) ConnectToUriFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_uri_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetEnableProxy is a wrapper around the C function g_socket_client_get_enable_proxy.
func (recv *SocketClient) GetEnableProxy() bool {
	retC := C.g_socket_client_get_enable_proxy((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTimeout is a wrapper around the C function g_socket_client_get_timeout.
func (recv *SocketClient) GetTimeout() uint32 {
	retC := C.g_socket_client_get_timeout((*C.GSocketClient)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// SetEnableProxy is a wrapper around the C function g_socket_client_set_enable_proxy.
func (recv *SocketClient) SetEnableProxy(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.g_socket_client_set_enable_proxy((*C.GSocketClient)(recv.native), c_enable)

	return
}

// SetTimeout is a wrapper around the C function g_socket_client_set_timeout.
func (recv *SocketClient) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_client_set_timeout((*C.GSocketClient)(recv.native), c_timeout)

	return
}

// ReceiveCredentials is a wrapper around the C function g_unix_connection_receive_credentials.
func (recv *UnixConnection) ReceiveCredentials(cancellable *Cancellable) (*Credentials, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SendCredentials is a wrapper around the C function g_unix_connection_send_credentials.
func (recv *UnixConnection) SendCredentials(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

// UnixCredentialsMessageNew is a wrapper around the C function g_unix_credentials_message_new.
func UnixCredentialsMessageNew() *UnixCredentialsMessage {
	retC := C.g_unix_credentials_message_new()
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixCredentialsMessageNewWithCredentials is a wrapper around the C function g_unix_credentials_message_new_with_credentials.
func UnixCredentialsMessageNewWithCredentials(credentials *Credentials) *UnixCredentialsMessage {
	c_credentials := (*C.GCredentials)(C.NULL)
	if credentials != nil {
		c_credentials = (*C.GCredentials)(credentials.ToC())
	}

	retC := C.g_unix_credentials_message_new_with_credentials(c_credentials)
	retGo := UnixCredentialsMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixCredentialsMessageIsSupported is a wrapper around the C function g_unix_credentials_message_is_supported.
func UnixCredentialsMessageIsSupported() bool {
	retC := C.g_unix_credentials_message_is_supported()
	retGo := retC == C.TRUE

	return retGo
}

// GetCredentials is a wrapper around the C function g_unix_credentials_message_get_credentials.
func (recv *UnixCredentialsMessage) GetCredentials() *Credentials {
	retC := C.g_unix_credentials_message_get_credentials((*C.GUnixCredentialsMessage)(recv.native))
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnixSocketAddressNewWithType is a wrapper around the C function g_unix_socket_address_new_with_type.
func UnixSocketAddressNewWithType(path []rune, type_ UnixSocketAddressType) *UnixSocketAddress {
	c_path_array := make([]C.gchar, len(path)+1, len(path)+1)
	for i, item := range path {
		c := (C.gchar)(item)
		c_path_array[i] = c
	}
	c_path_array[len(path)] = 0
	c_path_arrayPtr := &c_path_array[0]
	c_path := (*C.gchar)(unsafe.Pointer(c_path_arrayPtr))

	c_path_len := (C.gint)(len(path))

	c_type := (C.GUnixSocketAddressType)(type_)

	retC := C.g_unix_socket_address_new_with_type(c_path, c_path_len, c_type)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetAddressType is a wrapper around the C function g_unix_socket_address_get_address_type.
func (recv *UnixSocketAddress) GetAddressType() UnixSocketAddressType {
	retC := C.g_unix_socket_address_get_address_type((*C.GUnixSocketAddress)(recv.native))
	retGo := (UnixSocketAddressType)(retC)

	return retGo
}

// GetFileInfo is a wrapper around the C function g_zlib_compressor_get_file_info.
func (recv *ZlibCompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_compressor_get_file_info((*C.GZlibCompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFileInfo is a wrapper around the C function g_zlib_compressor_set_file_info.
func (recv *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	c_file_info := (*C.GFileInfo)(C.NULL)
	if fileInfo != nil {
		c_file_info = (*C.GFileInfo)(fileInfo.ToC())
	}

	C.g_zlib_compressor_set_file_info((*C.GZlibCompressor)(recv.native), c_file_info)

	return
}

// GetFileInfo is a wrapper around the C function g_zlib_decompressor_get_file_info.
func (recv *ZlibDecompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_decompressor_get_file_info((*C.GZlibDecompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

const PROXY_EXTENSION_POINT_NAME string = C.G_PROXY_EXTENSION_POINT_NAME

type BusType C.GBusType

const (
	BUS_TYPE_STARTER BusType = -1
	BUS_TYPE_NONE    BusType = 0
	BUS_TYPE_SYSTEM  BusType = 1
	BUS_TYPE_SESSION BusType = 2
)

type CredentialsType C.GCredentialsType

const (
	CREDENTIALS_TYPE_INVALID              CredentialsType = 0
	CREDENTIALS_TYPE_LINUX_UCRED          CredentialsType = 1
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED     CredentialsType = 2
	CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED CredentialsType = 3
	CREDENTIALS_TYPE_SOLARIS_UCRED        CredentialsType = 4
	CREDENTIALS_TYPE_NETBSD_UNPCBID       CredentialsType = 5
)

type DBusError C.GDBusError

const (
	DBUS_ERROR_FAILED                           DBusError = 0
	DBUS_ERROR_NO_MEMORY                        DBusError = 1
	DBUS_ERROR_SERVICE_UNKNOWN                  DBusError = 2
	DBUS_ERROR_NAME_HAS_NO_OWNER                DBusError = 3
	DBUS_ERROR_NO_REPLY                         DBusError = 4
	DBUS_ERROR_IO_ERROR                         DBusError = 5
	DBUS_ERROR_BAD_ADDRESS                      DBusError = 6
	DBUS_ERROR_NOT_SUPPORTED                    DBusError = 7
	DBUS_ERROR_LIMITS_EXCEEDED                  DBusError = 8
	DBUS_ERROR_ACCESS_DENIED                    DBusError = 9
	DBUS_ERROR_AUTH_FAILED                      DBusError = 10
	DBUS_ERROR_NO_SERVER                        DBusError = 11
	DBUS_ERROR_TIMEOUT                          DBusError = 12
	DBUS_ERROR_NO_NETWORK                       DBusError = 13
	DBUS_ERROR_ADDRESS_IN_USE                   DBusError = 14
	DBUS_ERROR_DISCONNECTED                     DBusError = 15
	DBUS_ERROR_INVALID_ARGS                     DBusError = 16
	DBUS_ERROR_FILE_NOT_FOUND                   DBusError = 17
	DBUS_ERROR_FILE_EXISTS                      DBusError = 18
	DBUS_ERROR_UNKNOWN_METHOD                   DBusError = 19
	DBUS_ERROR_TIMED_OUT                        DBusError = 20
	DBUS_ERROR_MATCH_RULE_NOT_FOUND             DBusError = 21
	DBUS_ERROR_MATCH_RULE_INVALID               DBusError = 22
	DBUS_ERROR_SPAWN_EXEC_FAILED                DBusError = 23
	DBUS_ERROR_SPAWN_FORK_FAILED                DBusError = 24
	DBUS_ERROR_SPAWN_CHILD_EXITED               DBusError = 25
	DBUS_ERROR_SPAWN_CHILD_SIGNALED             DBusError = 26
	DBUS_ERROR_SPAWN_FAILED                     DBusError = 27
	DBUS_ERROR_SPAWN_SETUP_FAILED               DBusError = 28
	DBUS_ERROR_SPAWN_CONFIG_INVALID             DBusError = 29
	DBUS_ERROR_SPAWN_SERVICE_INVALID            DBusError = 30
	DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND          DBusError = 31
	DBUS_ERROR_SPAWN_PERMISSIONS_INVALID        DBusError = 32
	DBUS_ERROR_SPAWN_FILE_INVALID               DBusError = 33
	DBUS_ERROR_SPAWN_NO_MEMORY                  DBusError = 34
	DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN          DBusError = 35
	DBUS_ERROR_INVALID_SIGNATURE                DBusError = 36
	DBUS_ERROR_INVALID_FILE_CONTENT             DBusError = 37
	DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN DBusError = 38
	DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN           DBusError = 39
	DBUS_ERROR_OBJECT_PATH_IN_USE               DBusError = 40
	DBUS_ERROR_UNKNOWN_OBJECT                   DBusError = 41
	DBUS_ERROR_UNKNOWN_INTERFACE                DBusError = 42
	DBUS_ERROR_UNKNOWN_PROPERTY                 DBusError = 43
	DBUS_ERROR_PROPERTY_READ_ONLY               DBusError = 44
)

// DBusErrorEncodeGerror is a wrapper around the C function g_dbus_error_encode_gerror.
func DBusErrorEncodeGerror(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_encode_gerror(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorGetRemoteError is a wrapper around the C function g_dbus_error_get_remote_error.
func DBusErrorGetRemoteError(error *glib.Error) string {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_get_remote_error(c_error)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorIsRemoteError is a wrapper around the C function g_dbus_error_is_remote_error.
func DBusErrorIsRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_is_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DBusErrorNewForDbusError is a wrapper around the C function g_dbus_error_new_for_dbus_error.
func DBusErrorNewForDbusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	retC := C.g_dbus_error_new_for_dbus_error(c_dbus_error_name, c_dbus_error_message)
	retGo := glib.ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DBusErrorRegisterError is a wrapper around the C function g_dbus_error_register_error.
func DBusErrorRegisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_register_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

// g_dbus_error_register_error_domain : unsupported parameter entries :
// DBusErrorSetDbusError is a wrapper around the C function g_dbus_error_set_dbus_error.
func DBusErrorSetDbusError(error *glib.Error, dbusErrorName string, dbusErrorMessage string, format string, args ...interface{}) {
	c_error := (**C.GError)(C.NULL)
	if error != nil {
		c_error = (**C.GError)(error.ToC())
	}

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	c_dbus_error_message := C.CString(dbusErrorMessage)
	defer C.free(unsafe.Pointer(c_dbus_error_message))

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_dbus_error_set_dbus_error(c_error, c_dbus_error_name, c_dbus_error_message, c_format)

	return
}

// g_dbus_error_set_dbus_error_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// DBusErrorStripRemoteError is a wrapper around the C function g_dbus_error_strip_remote_error.
func DBusErrorStripRemoteError(error *glib.Error) bool {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	retC := C.g_dbus_error_strip_remote_error(c_error)
	retGo := retC == C.TRUE

	return retGo
}

// DBusErrorUnregisterError is a wrapper around the C function g_dbus_error_unregister_error.
func DBusErrorUnregisterError(errorDomain glib.Quark, errorCode int32, dbusErrorName string) bool {
	c_error_domain := (C.GQuark)(errorDomain)

	c_error_code := (C.gint)(errorCode)

	c_dbus_error_name := C.CString(dbusErrorName)
	defer C.free(unsafe.Pointer(c_dbus_error_name))

	retC := C.g_dbus_error_unregister_error(c_error_domain, c_error_code, c_dbus_error_name)
	retGo := retC == C.TRUE

	return retGo
}

type DBusMessageByteOrder C.GDBusMessageByteOrder

const (
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN    DBusMessageByteOrder = 66
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN DBusMessageByteOrder = 108
)

type DBusMessageHeaderField C.GDBusMessageHeaderField

const (
	DBUS_MESSAGE_HEADER_FIELD_INVALID      DBusMessageHeaderField = 0
	DBUS_MESSAGE_HEADER_FIELD_PATH         DBusMessageHeaderField = 1
	DBUS_MESSAGE_HEADER_FIELD_INTERFACE    DBusMessageHeaderField = 2
	DBUS_MESSAGE_HEADER_FIELD_MEMBER       DBusMessageHeaderField = 3
	DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME   DBusMessageHeaderField = 4
	DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL DBusMessageHeaderField = 5
	DBUS_MESSAGE_HEADER_FIELD_DESTINATION  DBusMessageHeaderField = 6
	DBUS_MESSAGE_HEADER_FIELD_SENDER       DBusMessageHeaderField = 7
	DBUS_MESSAGE_HEADER_FIELD_SIGNATURE    DBusMessageHeaderField = 8
	DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS DBusMessageHeaderField = 9
)

type DBusMessageType C.GDBusMessageType

const (
	DBUS_MESSAGE_TYPE_INVALID       DBusMessageType = 0
	DBUS_MESSAGE_TYPE_METHOD_CALL   DBusMessageType = 1
	DBUS_MESSAGE_TYPE_METHOD_RETURN DBusMessageType = 2
	DBUS_MESSAGE_TYPE_ERROR         DBusMessageType = 3
	DBUS_MESSAGE_TYPE_SIGNAL        DBusMessageType = 4
)

type UnixSocketAddressType C.GUnixSocketAddressType

const (
	UNIX_SOCKET_ADDRESS_INVALID         UnixSocketAddressType = 0
	UNIX_SOCKET_ADDRESS_ANONYMOUS       UnixSocketAddressType = 1
	UNIX_SOCKET_ADDRESS_PATH            UnixSocketAddressType = 2
	UNIX_SOCKET_ADDRESS_ABSTRACT        UnixSocketAddressType = 3
	UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED UnixSocketAddressType = 4
)

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
func DbusIsAddress(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_address(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsGuid is a wrapper around the C function g_dbus_is_guid.
func DbusIsGuid(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_guid(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsInterfaceName is a wrapper around the C function g_dbus_is_interface_name.
func DbusIsInterfaceName(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_interface_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsMemberName is a wrapper around the C function g_dbus_is_member_name.
func DbusIsMemberName(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_member_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsName is a wrapper around the C function g_dbus_is_name.
func DbusIsName(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// DbusIsSupportedAddress is a wrapper around the C function g_dbus_is_supported_address.
func DbusIsSupportedAddress(string_ string) (bool, error) {
	c_string := C.CString(string_)
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
func DbusIsUniqueName(string_ string) bool {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_dbus_is_unique_name(c_string)
	retGo := retC == C.TRUE

	return retGo
}

// Proxy is a wrapper around the C record GProxy.
type Proxy struct {
	native *C.GProxy
}

func ProxyNewFromC(u unsafe.Pointer) *Proxy {
	c := (*C.GProxy)(u)
	if c == nil {
		return nil
	}

	g := &Proxy{native: c}

	return g
}

func (recv *Proxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Proxy with another Proxy, and returns true if they represent the same GObject.
func (recv *Proxy) Equals(other *Proxy) bool {
	return other.ToC() == recv.ToC()
}

// ProxyGetDefaultForProtocol is a wrapper around the C function g_proxy_get_default_for_protocol.
func ProxyGetDefaultForProtocol(protocol string) *Proxy {
	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	retC := C.g_proxy_get_default_for_protocol(c_protocol)
	retGo := ProxyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Connect is a wrapper around the C function g_proxy_connect.
func (recv *Proxy) Connect(connection *IOStream, proxyAddress *ProxyAddress, cancellable *Cancellable) (*IOStream, error) {
	c_connection := (*C.GIOStream)(C.NULL)
	if connection != nil {
		c_connection = (*C.GIOStream)(connection.ToC())
	}

	c_proxy_address := (*C.GProxyAddress)(C.NULL)
	if proxyAddress != nil {
		c_proxy_address = (*C.GProxyAddress)(proxyAddress.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_proxy_connect((*C.GProxy)(recv.native), c_connection, c_proxy_address, c_cancellable, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_proxy_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectFinish is a wrapper around the C function g_proxy_connect_finish.
func (recv *Proxy) ConnectFinish(result *AsyncResult) (*IOStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_proxy_connect_finish((*C.GProxy)(recv.native), c_result, &cThrowableError)
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SupportsHostname is a wrapper around the C function g_proxy_supports_hostname.
func (recv *Proxy) SupportsHostname() bool {
	retC := C.g_proxy_supports_hostname((*C.GProxy)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ProxyResolver is a wrapper around the C record GProxyResolver.
type ProxyResolver struct {
	native *C.GProxyResolver
}

func ProxyResolverNewFromC(u unsafe.Pointer) *ProxyResolver {
	c := (*C.GProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolver{native: c}

	return g
}

func (recv *ProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyResolver with another ProxyResolver, and returns true if they represent the same GObject.
func (recv *ProxyResolver) Equals(other *ProxyResolver) bool {
	return other.ToC() == recv.ToC()
}

// ProxyResolverGetDefault is a wrapper around the C function g_proxy_resolver_get_default.
func ProxyResolverGetDefault() *ProxyResolver {
	retC := C.g_proxy_resolver_get_default()
	retGo := ProxyResolverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsSupported is a wrapper around the C function g_proxy_resolver_is_supported.
func (recv *ProxyResolver) IsSupported() bool {
	retC := C.g_proxy_resolver_is_supported((*C.GProxyResolver)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Lookup is a wrapper around the C function g_proxy_resolver_lookup.
func (recv *ProxyResolver) Lookup(uri string, cancellable *Cancellable) ([]string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_proxy_resolver_lookup((*C.GProxyResolver)(recv.native), c_uri, c_cancellable, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_proxy_resolver_lookup_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupFinish is a wrapper around the C function g_proxy_resolver_lookup_finish.
func (recv *ProxyResolver) LookupFinish(result *AsyncResult) ([]string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_proxy_resolver_lookup_finish((*C.GProxyResolver)(recv.native), c_result, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ProxyEnumerate is a wrapper around the C function g_socket_connectable_proxy_enumerate.
func (recv *SocketConnectable) ProxyEnumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_proxy_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// Equals compares this CredentialsClass with another CredentialsClass, and returns true if they represent the same GObject.
func (recv *CredentialsClass) Equals(other *CredentialsClass) bool {
	return other.ToC() == recv.ToC()
}

// DBusAnnotationInfo is a wrapper around the C record GDBusAnnotationInfo.
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

// DBusArgInfo is a wrapper around the C record GDBusArgInfo.
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

// Equals compares this DBusErrorEntry with another DBusErrorEntry, and returns true if they represent the same GObject.
func (recv *DBusErrorEntry) Equals(other *DBusErrorEntry) bool {
	return other.ToC() == recv.ToC()
}

// DBusInterfaceInfo is a wrapper around the C record GDBusInterfaceInfo.
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

// Equals compares this DBusInterfaceVTable with another DBusInterfaceVTable, and returns true if they represent the same GObject.
func (recv *DBusInterfaceVTable) Equals(other *DBusInterfaceVTable) bool {
	return other.ToC() == recv.ToC()
}

// DBusMethodInfo is a wrapper around the C record GDBusMethodInfo.
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

// DBusNodeInfo is a wrapper around the C record GDBusNodeInfo.
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

// Equals compares this DBusNodeInfo with another DBusNodeInfo, and returns true if they represent the same GObject.
func (recv *DBusNodeInfo) Equals(other *DBusNodeInfo) bool {
	return other.ToC() == recv.ToC()
}

// DBusNodeInfoNewForXml is a wrapper around the C function g_dbus_node_info_new_for_xml.
func DBusNodeInfoNewForXml(xmlData string) (*DBusNodeInfo, error) {
	c_xml_data := C.CString(xmlData)
	defer C.free(unsafe.Pointer(c_xml_data))

	var cThrowableError *C.GError

	retC := C.g_dbus_node_info_new_for_xml(c_xml_data, &cThrowableError)
	retGo := DBusNodeInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

// DBusPropertyInfo is a wrapper around the C record GDBusPropertyInfo.
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

// Equals compares this DBusProxyClass with another DBusProxyClass, and returns true if they represent the same GObject.
func (recv *DBusProxyClass) Equals(other *DBusProxyClass) bool {
	return other.ToC() == recv.ToC()
}

// DBusSignalInfo is a wrapper around the C record GDBusSignalInfo.
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

// Equals compares this DBusSubtreeVTable with another DBusSubtreeVTable, and returns true if they represent the same GObject.
func (recv *DBusSubtreeVTable) Equals(other *DBusSubtreeVTable) bool {
	return other.ToC() == recv.ToC()
}

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

// Equals compares this ProxyAddressClass with another ProxyAddressClass, and returns true if they represent the same GObject.
func (recv *ProxyAddressClass) Equals(other *ProxyAddressClass) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this ProxyInterface with another ProxyInterface, and returns true if they represent the same GObject.
func (recv *ProxyInterface) Equals(other *ProxyInterface) bool {
	return other.ToC() == recv.ToC()
}

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

// Equals compares this TlsClientConnectionInterface with another TlsClientConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsClientConnectionInterface) Equals(other *TlsClientConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this TlsServerConnectionInterface with another TlsServerConnectionInterface, and returns true if they represent the same GObject.
func (recv *TlsServerConnectionInterface) Equals(other *TlsServerConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this UnixCredentialsMessageClass with another UnixCredentialsMessageClass, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessageClass) Equals(other *UnixCredentialsMessageClass) bool {
	return other.ToC() == recv.ToC()
}
