// Code generated - DO NOT EDIT.
// +build soup_2.50

package soup

import (
	"fmt"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <libsoup/soup.h>
// #include <libsoup/soup-password-manager.h>
// #include <stdlib.h>
/*

	void authmanager_authenticateHandler(GObject *, SoupMessage *, SoupAuth *, gboolean, gpointer);

	static gulong AuthManager_signal_connect_authenticate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "authenticate", G_CALLBACK(authmanager_authenticateHandler), data);
	}

*/
/*

	void cookiejar_changedHandler(GObject *, SoupCookie *, SoupCookie *, gpointer);

	static gulong CookieJar_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(cookiejar_changedHandler), data);
	}

*/
/*

	void message_contentSniffedHandler(GObject *, gchar*, GHashTable *, gpointer);

	static gulong Message_signal_connect_content_sniffed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "content-sniffed", G_CALLBACK(message_contentSniffedHandler), data);
	}

*/
/*

	void message_finishedHandler(GObject *, gpointer);

	static gulong Message_signal_connect_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "finished", G_CALLBACK(message_finishedHandler), data);
	}

*/
/*

	void message_gotBodyHandler(GObject *, gpointer);

	static gulong Message_signal_connect_got_body(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-body", G_CALLBACK(message_gotBodyHandler), data);
	}

*/
/*

	void message_gotChunkHandler(GObject *, SoupBuffer *, gpointer);

	static gulong Message_signal_connect_got_chunk(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-chunk", G_CALLBACK(message_gotChunkHandler), data);
	}

*/
/*

	void message_gotHeadersHandler(GObject *, gpointer);

	static gulong Message_signal_connect_got_headers(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-headers", G_CALLBACK(message_gotHeadersHandler), data);
	}

*/
/*

	void message_gotInformationalHandler(GObject *, gpointer);

	static gulong Message_signal_connect_got_informational(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-informational", G_CALLBACK(message_gotInformationalHandler), data);
	}

*/
/*

	void message_networkEventHandler(GObject *, GSocketClientEvent, GIOStream *, gpointer);

	static gulong Message_signal_connect_network_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "network-event", G_CALLBACK(message_networkEventHandler), data);
	}

*/
/*

	void message_restartedHandler(GObject *, gpointer);

	static gulong Message_signal_connect_restarted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "restarted", G_CALLBACK(message_restartedHandler), data);
	}

*/
/*

	void message_startingHandler(GObject *, gpointer);

	static gulong Message_signal_connect_starting(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "starting", G_CALLBACK(message_startingHandler), data);
	}

*/
/*

	void message_wroteBodyHandler(GObject *, gpointer);

	static gulong Message_signal_connect_wrote_body(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrote-body", G_CALLBACK(message_wroteBodyHandler), data);
	}

*/
/*

	void message_wroteBodyDataHandler(GObject *, SoupBuffer *, gpointer);

	static gulong Message_signal_connect_wrote_body_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrote-body-data", G_CALLBACK(message_wroteBodyDataHandler), data);
	}

*/
/*

	void message_wroteChunkHandler(GObject *, gpointer);

	static gulong Message_signal_connect_wrote_chunk(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrote-chunk", G_CALLBACK(message_wroteChunkHandler), data);
	}

*/
/*

	void message_wroteHeadersHandler(GObject *, gpointer);

	static gulong Message_signal_connect_wrote_headers(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrote-headers", G_CALLBACK(message_wroteHeadersHandler), data);
	}

*/
/*

	void message_wroteInformationalHandler(GObject *, gpointer);

	static gulong Message_signal_connect_wrote_informational(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "wrote-informational", G_CALLBACK(message_wroteInformationalHandler), data);
	}

*/
/*

	void server_requestAbortedHandler(GObject *, SoupMessage *, SoupClientContext *, gpointer);

	static gulong Server_signal_connect_request_aborted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-aborted", G_CALLBACK(server_requestAbortedHandler), data);
	}

*/
/*

	void server_requestFinishedHandler(GObject *, SoupMessage *, SoupClientContext *, gpointer);

	static gulong Server_signal_connect_request_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-finished", G_CALLBACK(server_requestFinishedHandler), data);
	}

*/
/*

	void server_requestReadHandler(GObject *, SoupMessage *, SoupClientContext *, gpointer);

	static gulong Server_signal_connect_request_read(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-read", G_CALLBACK(server_requestReadHandler), data);
	}

*/
/*

	void server_requestStartedHandler(GObject *, SoupMessage *, SoupClientContext *, gpointer);

	static gulong Server_signal_connect_request_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-started", G_CALLBACK(server_requestStartedHandler), data);
	}

*/
/*

	void session_authenticateHandler(GObject *, SoupMessage *, SoupAuth *, gboolean, gpointer);

	static gulong Session_signal_connect_authenticate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "authenticate", G_CALLBACK(session_authenticateHandler), data);
	}

*/
/*

	void session_connectionCreatedHandler(GObject *, GObject *, gpointer);

	static gulong Session_signal_connect_connection_created(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connection-created", G_CALLBACK(session_connectionCreatedHandler), data);
	}

*/
/*

	void session_requestQueuedHandler(GObject *, SoupMessage *, gpointer);

	static gulong Session_signal_connect_request_queued(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-queued", G_CALLBACK(session_requestQueuedHandler), data);
	}

*/
/*

	void session_requestStartedHandler(GObject *, SoupMessage *, SoupSocket *, gpointer);

	static gulong Session_signal_connect_request_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-started", G_CALLBACK(session_requestStartedHandler), data);
	}

*/
/*

	void session_requestUnqueuedHandler(GObject *, SoupMessage *, gpointer);

	static gulong Session_signal_connect_request_unqueued(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "request-unqueued", G_CALLBACK(session_requestUnqueuedHandler), data);
	}

*/
/*

	void session_tunnelingHandler(GObject *, GObject *, gpointer);

	static gulong Session_signal_connect_tunneling(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tunneling", G_CALLBACK(session_tunnelingHandler), data);
	}

*/
/*

	void socket_disconnectedHandler(GObject *, gpointer);

	static gulong Socket_signal_connect_disconnected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnected", G_CALLBACK(socket_disconnectedHandler), data);
	}

*/
/*

	void socket_eventHandler(GObject *, GSocketClientEvent, GIOStream *, gpointer);

	static gulong Socket_signal_connect_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "event", G_CALLBACK(socket_eventHandler), data);
	}

*/
/*

	void socket_newConnectionHandler(GObject *, SoupSocket *, gpointer);

	static gulong Socket_signal_connect_new_connection(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "new-connection", G_CALLBACK(socket_newConnectionHandler), data);
	}

*/
/*

	void socket_readableHandler(GObject *, gpointer);

	static gulong Socket_signal_connect_readable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "readable", G_CALLBACK(socket_readableHandler), data);
	}

*/
/*

	void socket_writableHandler(GObject *, gpointer);

	static gulong Socket_signal_connect_writable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "writable", G_CALLBACK(socket_writableHandler), data);
	}

*/
/*

	void websocketconnection_closedHandler(GObject *, gpointer);

	static gulong WebsocketConnection_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(websocketconnection_closedHandler), data);
	}

*/
/*

	void websocketconnection_closingHandler(GObject *, gpointer);

	static gulong WebsocketConnection_signal_connect_closing(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closing", G_CALLBACK(websocketconnection_closingHandler), data);
	}

*/
/*

	void websocketconnection_errorHandler(GObject *, GError *, gpointer);

	static gulong WebsocketConnection_signal_connect_error(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "error", G_CALLBACK(websocketconnection_errorHandler), data);
	}

*/
/*

	void websocketconnection_messageHandler(GObject *, gint, GBytes *, gpointer);

	static gulong WebsocketConnection_signal_connect_message(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "message", G_CALLBACK(websocketconnection_messageHandler), data);
	}

*/
/*

	static char* _soup_xmlrpc_build_fault(int fault_code, const char* fault_format) {
		return soup_xmlrpc_build_fault(fault_code, fault_format);
    }
*/
/*

	static void _soup_xmlrpc_set_fault(SoupMessage* msg, int fault_code, const char* fault_format) {
		return soup_xmlrpc_set_fault(msg, fault_code, fault_format);
    }
*/
import "C"

type Cacheability C.SoupCacheability

const (
	SOUP_CACHE_CACHEABLE   Cacheability = 1
	SOUP_CACHE_UNCACHEABLE Cacheability = 2
	SOUP_CACHE_INVALIDATES Cacheability = 4
	SOUP_CACHE_VALIDATES   Cacheability = 8
)

type Expectation C.SoupExpectation

const (
	SOUP_EXPECTATION_UNRECOGNIZED Expectation = 1
	SOUP_EXPECTATION_CONTINUE     Expectation = 2
)

type MessageFlags C.SoupMessageFlags

const (
	SOUP_MESSAGE_NO_REDIRECT              MessageFlags = 2
	SOUP_MESSAGE_CAN_REBUILD              MessageFlags = 4
	SOUP_MESSAGE_OVERWRITE_CHUNKS         MessageFlags = 8
	SOUP_MESSAGE_CONTENT_DECODED          MessageFlags = 16
	SOUP_MESSAGE_CERTIFICATE_TRUSTED      MessageFlags = 32
	SOUP_MESSAGE_NEW_CONNECTION           MessageFlags = 64
	SOUP_MESSAGE_IDEMPOTENT               MessageFlags = 128
	SOUP_MESSAGE_IGNORE_CONNECTION_LIMITS MessageFlags = 256
	SOUP_MESSAGE_DO_NOT_USE_AUTH_CACHE    MessageFlags = 512
)

type ServerListenOptions C.SoupServerListenOptions

const (
	SOUP_SERVER_LISTEN_HTTPS     ServerListenOptions = 1
	SOUP_SERVER_LISTEN_IPV4_ONLY ServerListenOptions = 2
	SOUP_SERVER_LISTEN_IPV6_ONLY ServerListenOptions = 4
)

// Address is a wrapper around the C record SoupAddress.
type Address struct {
	native *C.SoupAddress
	// parent : record
}

func AddressNewFromC(u unsafe.Pointer) *Address {
	c := (*C.SoupAddress)(u)
	if c == nil {
		return nil
	}

	g := &Address{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Address) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Address) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Address with another Address, and returns true if they represent the same GObject.
func (recv *Address) Equals(other *Address) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Address) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Address.
// Exercise care, as this is a potentially dangerous function if the Object is not a Address.
func CastToAddress(object *gobject.Object) *Address {
	return AddressNewFromC(object.ToC())
}

// AddressNew is a wrapper around the C function soup_address_new.
func AddressNew(name string, port uint32) *Address {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_port := (C.guint)(port)

	retC := C.soup_address_new(c_name, c_port)
	retGo := AddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddressNewAny is a wrapper around the C function soup_address_new_any.
func AddressNewAny(family AddressFamily, port uint32) *Address {
	c_family := (C.SoupAddressFamily)(family)

	c_port := (C.guint)(port)

	retC := C.soup_address_new_any(c_family, c_port)
	var retGo (*Address)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AddressNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : soup_address_new_from_sockaddr

// EqualByIp is a wrapper around the C function soup_address_equal_by_ip.
func (recv *Address) EqualByIp(addr2 *Address) bool {
	c_addr2 := (C.gconstpointer)(C.NULL)
	if addr2 != nil {
		c_addr2 = (C.gconstpointer)(addr2.ToC())
	}

	retC := C.soup_address_equal_by_ip((C.gconstpointer)(recv.native), c_addr2)
	retGo := retC == C.TRUE

	return retGo
}

// EqualByName is a wrapper around the C function soup_address_equal_by_name.
func (recv *Address) EqualByName(addr2 *Address) bool {
	c_addr2 := (C.gconstpointer)(C.NULL)
	if addr2 != nil {
		c_addr2 = (C.gconstpointer)(addr2.ToC())
	}

	retC := C.soup_address_equal_by_name((C.gconstpointer)(recv.native), c_addr2)
	retGo := retC == C.TRUE

	return retGo
}

// GetGsockaddr is a wrapper around the C function soup_address_get_gsockaddr.
func (recv *Address) GetGsockaddr() *gio.SocketAddress {
	retC := C.soup_address_get_gsockaddr((*C.SoupAddress)(recv.native))
	retGo := gio.SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function soup_address_get_name.
func (recv *Address) GetName() string {
	retC := C.soup_address_get_name((*C.SoupAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPhysical is a wrapper around the C function soup_address_get_physical.
func (recv *Address) GetPhysical() string {
	retC := C.soup_address_get_physical((*C.SoupAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function soup_address_get_port.
func (recv *Address) GetPort() uint32 {
	retC := C.soup_address_get_port((*C.SoupAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSockaddr is a wrapper around the C function soup_address_get_sockaddr.
func (recv *Address) GetSockaddr(len int32) uintptr {
	c_len := (C.int)(len)

	retC := C.soup_address_get_sockaddr((*C.SoupAddress)(recv.native), &c_len)
	retGo := (uintptr)(unsafe.Pointer(&retC))

	return retGo
}

// HashByIp is a wrapper around the C function soup_address_hash_by_ip.
func (recv *Address) HashByIp() uint32 {
	retC := C.soup_address_hash_by_ip((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// HashByName is a wrapper around the C function soup_address_hash_by_name.
func (recv *Address) HashByName() uint32 {
	retC := C.soup_address_hash_by_name((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsResolved is a wrapper around the C function soup_address_is_resolved.
func (recv *Address) IsResolved() bool {
	retC := C.soup_address_is_resolved((*C.SoupAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : soup_address_resolve_async : unsupported parameter callback : no type generator for AddressCallback (SoupAddressCallback) for param callback

// ResolveSync is a wrapper around the C function soup_address_resolve_sync.
func (recv *Address) ResolveSync(cancellable *gio.Cancellable) uint32 {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.soup_address_resolve_sync((*C.SoupAddress)(recv.native), c_cancellable)
	retGo := (uint32)(retC)

	return retGo
}

// SocketConnectable returns the SocketConnectable interface implemented by Address
func (recv *Address) SocketConnectable() *gio.SocketConnectable {
	return gio.SocketConnectableNewFromC(recv.ToC())
}

// Auth is a wrapper around the C record SoupAuth.
type Auth struct {
	native *C.SoupAuth
	// parent : record
	Realm string
}

func AuthNewFromC(u unsafe.Pointer) *Auth {
	c := (*C.SoupAuth)(u)
	if c == nil {
		return nil
	}

	g := &Auth{
		Realm:  C.GoString(c.realm),
		native: c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Auth) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Auth) ToC() unsafe.Pointer {
	recv.native.realm =
		C.CString(recv.Realm)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Auth with another Auth, and returns true if they represent the same GObject.
func (recv *Auth) Equals(other *Auth) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Auth) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Auth.
// Exercise care, as this is a potentially dangerous function if the Object is not a Auth.
func CastToAuth(object *gobject.Object) *Auth {
	return AuthNewFromC(object.ToC())
}

// AuthNew is a wrapper around the C function soup_auth_new.
func AuthNew(type_ gobject.Type, msg *Message, authHeader string) *Auth {
	c_type := (C.GType)(type_)

	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_auth_header := C.CString(authHeader)
	defer C.free(unsafe.Pointer(c_auth_header))

	retC := C.soup_auth_new(c_type, c_msg, c_auth_header)
	var retGo (*Auth)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AuthNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Authenticate is a wrapper around the C function soup_auth_authenticate.
func (recv *Auth) Authenticate(username string, password string) {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.soup_auth_authenticate((*C.SoupAuth)(recv.native), c_username, c_password)

	return
}

// FreeProtectionSpace is a wrapper around the C function soup_auth_free_protection_space.
func (recv *Auth) FreeProtectionSpace(space *glib.SList) {
	c_space := (*C.GSList)(C.NULL)
	if space != nil {
		c_space = (*C.GSList)(space.ToC())
	}

	C.soup_auth_free_protection_space((*C.SoupAuth)(recv.native), c_space)

	return
}

// GetAuthorization is a wrapper around the C function soup_auth_get_authorization.
func (recv *Auth) GetAuthorization(msg *Message) string {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_auth_get_authorization((*C.SoupAuth)(recv.native), c_msg)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetHost is a wrapper around the C function soup_auth_get_host.
func (recv *Auth) GetHost() string {
	retC := C.soup_auth_get_host((*C.SoupAuth)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInfo is a wrapper around the C function soup_auth_get_info.
func (recv *Auth) GetInfo() string {
	retC := C.soup_auth_get_info((*C.SoupAuth)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetProtectionSpace is a wrapper around the C function soup_auth_get_protection_space.
func (recv *Auth) GetProtectionSpace(sourceUri *URI) *glib.SList {
	c_source_uri := (*C.SoupURI)(C.NULL)
	if sourceUri != nil {
		c_source_uri = (*C.SoupURI)(sourceUri.ToC())
	}

	retC := C.soup_auth_get_protection_space((*C.SoupAuth)(recv.native), c_source_uri)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRealm is a wrapper around the C function soup_auth_get_realm.
func (recv *Auth) GetRealm() string {
	retC := C.soup_auth_get_realm((*C.SoupAuth)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSavedPassword is a wrapper around the C function soup_auth_get_saved_password.
func (recv *Auth) GetSavedPassword(user string) string {
	c_user := C.CString(user)
	defer C.free(unsafe.Pointer(c_user))

	retC := C.soup_auth_get_saved_password((*C.SoupAuth)(recv.native), c_user)
	retGo := C.GoString(retC)

	return retGo
}

// GetSavedUsers is a wrapper around the C function soup_auth_get_saved_users.
func (recv *Auth) GetSavedUsers() *glib.SList {
	retC := C.soup_auth_get_saved_users((*C.SoupAuth)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSchemeName is a wrapper around the C function soup_auth_get_scheme_name.
func (recv *Auth) GetSchemeName() string {
	retC := C.soup_auth_get_scheme_name((*C.SoupAuth)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// HasSavedPassword is a wrapper around the C function soup_auth_has_saved_password.
func (recv *Auth) HasSavedPassword(username string, password string) {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.soup_auth_has_saved_password((*C.SoupAuth)(recv.native), c_username, c_password)

	return
}

// IsAuthenticated is a wrapper around the C function soup_auth_is_authenticated.
func (recv *Auth) IsAuthenticated() bool {
	retC := C.soup_auth_is_authenticated((*C.SoupAuth)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsForProxy is a wrapper around the C function soup_auth_is_for_proxy.
func (recv *Auth) IsForProxy() bool {
	retC := C.soup_auth_is_for_proxy((*C.SoupAuth)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsReady is a wrapper around the C function soup_auth_is_ready.
func (recv *Auth) IsReady(msg *Message) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_auth_is_ready((*C.SoupAuth)(recv.native), c_msg)
	retGo := retC == C.TRUE

	return retGo
}

// SavePassword is a wrapper around the C function soup_auth_save_password.
func (recv *Auth) SavePassword(username string, password string) {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.soup_auth_save_password((*C.SoupAuth)(recv.native), c_username, c_password)

	return
}

// Update is a wrapper around the C function soup_auth_update.
func (recv *Auth) Update(msg *Message, authHeader string) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_auth_header := C.CString(authHeader)
	defer C.free(unsafe.Pointer(c_auth_header))

	retC := C.soup_auth_update((*C.SoupAuth)(recv.native), c_msg, c_auth_header)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : AuthBasic : no CType

// Unsupported : AuthDigest : no CType

// AuthDomain is a wrapper around the C record SoupAuthDomain.
type AuthDomain struct {
	native *C.SoupAuthDomain
	// parent : record
}

func AuthDomainNewFromC(u unsafe.Pointer) *AuthDomain {
	c := (*C.SoupAuthDomain)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomain{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AuthDomain) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AuthDomain) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomain with another AuthDomain, and returns true if they represent the same GObject.
func (recv *AuthDomain) Equals(other *AuthDomain) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AuthDomain) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AuthDomain.
// Exercise care, as this is a potentially dangerous function if the Object is not a AuthDomain.
func CastToAuthDomain(object *gobject.Object) *AuthDomain {
	return AuthDomainNewFromC(object.ToC())
}

// Accepts is a wrapper around the C function soup_auth_domain_accepts.
func (recv *AuthDomain) Accepts(msg *Message) string {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_auth_domain_accepts((*C.SoupAuthDomain)(recv.native), c_msg)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AddPath is a wrapper around the C function soup_auth_domain_add_path.
func (recv *AuthDomain) AddPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.soup_auth_domain_add_path((*C.SoupAuthDomain)(recv.native), c_path)

	return
}

// Unsupported : soup_auth_domain_basic_set_auth_callback : unsupported parameter callback : no type generator for AuthDomainBasicAuthCallback (SoupAuthDomainBasicAuthCallback) for param callback

// Challenge is a wrapper around the C function soup_auth_domain_challenge.
func (recv *AuthDomain) Challenge(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_auth_domain_challenge((*C.SoupAuthDomain)(recv.native), c_msg)

	return
}

// CheckPassword is a wrapper around the C function soup_auth_domain_check_password.
func (recv *AuthDomain) CheckPassword(msg *Message, username string, password string) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.soup_auth_domain_check_password((*C.SoupAuthDomain)(recv.native), c_msg, c_username, c_password)
	retGo := retC == C.TRUE

	return retGo
}

// Covers is a wrapper around the C function soup_auth_domain_covers.
func (recv *AuthDomain) Covers(msg *Message) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_auth_domain_covers((*C.SoupAuthDomain)(recv.native), c_msg)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : soup_auth_domain_digest_set_auth_callback : unsupported parameter callback : no type generator for AuthDomainDigestAuthCallback (SoupAuthDomainDigestAuthCallback) for param callback

// GetRealm is a wrapper around the C function soup_auth_domain_get_realm.
func (recv *AuthDomain) GetRealm() string {
	retC := C.soup_auth_domain_get_realm((*C.SoupAuthDomain)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// RemovePath is a wrapper around the C function soup_auth_domain_remove_path.
func (recv *AuthDomain) RemovePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.soup_auth_domain_remove_path((*C.SoupAuthDomain)(recv.native), c_path)

	return
}

// Unsupported : soup_auth_domain_set_filter : unsupported parameter filter : no type generator for AuthDomainFilter (SoupAuthDomainFilter) for param filter

// Unsupported : soup_auth_domain_set_generic_auth_callback : unsupported parameter auth_callback : no type generator for AuthDomainGenericAuthCallback (SoupAuthDomainGenericAuthCallback) for param auth_callback

// TryGenericAuthCallback is a wrapper around the C function soup_auth_domain_try_generic_auth_callback.
func (recv *AuthDomain) TryGenericAuthCallback(msg *Message, username string) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	retC := C.soup_auth_domain_try_generic_auth_callback((*C.SoupAuthDomain)(recv.native), c_msg, c_username)
	retGo := retC == C.TRUE

	return retGo
}

// AuthDomainBasic is a wrapper around the C record SoupAuthDomainBasic.
type AuthDomainBasic struct {
	native *C.SoupAuthDomainBasic
	// parent : record
}

func AuthDomainBasicNewFromC(u unsafe.Pointer) *AuthDomainBasic {
	c := (*C.SoupAuthDomainBasic)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomainBasic{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AuthDomainBasic) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AuthDomainBasic) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomainBasic with another AuthDomainBasic, and returns true if they represent the same GObject.
func (recv *AuthDomainBasic) Equals(other *AuthDomainBasic) bool {
	return other.ToC() == recv.ToC()
}

// AuthDomain upcasts to *AuthDomain
func (recv *AuthDomainBasic) AuthDomain() *AuthDomain {
	return AuthDomainNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *AuthDomainBasic) Object() *gobject.Object {
	return recv.AuthDomain().Object()
}

// CastToWidget down casts any arbitrary Object to AuthDomainBasic.
// Exercise care, as this is a potentially dangerous function if the Object is not a AuthDomainBasic.
func CastToAuthDomainBasic(object *gobject.Object) *AuthDomainBasic {
	return AuthDomainBasicNewFromC(object.ToC())
}

// Unsupported : soup_auth_domain_basic_new : unsupported parameter ... : varargs

// AuthDomainDigest is a wrapper around the C record SoupAuthDomainDigest.
type AuthDomainDigest struct {
	native *C.SoupAuthDomainDigest
	// parent : record
}

func AuthDomainDigestNewFromC(u unsafe.Pointer) *AuthDomainDigest {
	c := (*C.SoupAuthDomainDigest)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomainDigest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AuthDomainDigest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AuthDomainDigest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomainDigest with another AuthDomainDigest, and returns true if they represent the same GObject.
func (recv *AuthDomainDigest) Equals(other *AuthDomainDigest) bool {
	return other.ToC() == recv.ToC()
}

// AuthDomain upcasts to *AuthDomain
func (recv *AuthDomainDigest) AuthDomain() *AuthDomain {
	return AuthDomainNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *AuthDomainDigest) Object() *gobject.Object {
	return recv.AuthDomain().Object()
}

// CastToWidget down casts any arbitrary Object to AuthDomainDigest.
// Exercise care, as this is a potentially dangerous function if the Object is not a AuthDomainDigest.
func CastToAuthDomainDigest(object *gobject.Object) *AuthDomainDigest {
	return AuthDomainDigestNewFromC(object.ToC())
}

// Unsupported : soup_auth_domain_digest_new : unsupported parameter ... : varargs

// AuthDomainDigestEncodePassword is a wrapper around the C function soup_auth_domain_digest_encode_password.
func AuthDomainDigestEncodePassword(username string, realm string, password string) string {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_realm := C.CString(realm)
	defer C.free(unsafe.Pointer(c_realm))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.soup_auth_domain_digest_encode_password(c_username, c_realm, c_password)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// AuthManager is a wrapper around the C record SoupAuthManager.
type AuthManager struct {
	native *C.SoupAuthManager
	// parent : record
	// priv : record
}

func AuthManagerNewFromC(u unsafe.Pointer) *AuthManager {
	c := (*C.SoupAuthManager)(u)
	if c == nil {
		return nil
	}

	g := &AuthManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AuthManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AuthManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthManager with another AuthManager, and returns true if they represent the same GObject.
func (recv *AuthManager) Equals(other *AuthManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AuthManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AuthManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a AuthManager.
func CastToAuthManager(object *gobject.Object) *AuthManager {
	return AuthManagerNewFromC(object.ToC())
}

type signalAuthManagerAuthenticateDetail struct {
	callback  AuthManagerSignalAuthenticateCallback
	handlerID C.gulong
}

var signalAuthManagerAuthenticateId int
var signalAuthManagerAuthenticateMap = make(map[int]signalAuthManagerAuthenticateDetail)
var signalAuthManagerAuthenticateLock sync.RWMutex

// AuthManagerSignalAuthenticateCallback is a callback function for a 'authenticate' signal emitted from a AuthManager.
type AuthManagerSignalAuthenticateCallback func(targetObject *AuthManager, msg *Message, auth *Auth, retrying bool)

/*
ConnectAuthenticate connects the callback to the 'authenticate' signal for the AuthManager.

The returned value represents the connection, and may be passed to DisconnectAuthenticate to remove it.
*/
func (recv *AuthManager) ConnectAuthenticate(callback AuthManagerSignalAuthenticateCallback) int {
	signalAuthManagerAuthenticateLock.Lock()
	defer signalAuthManagerAuthenticateLock.Unlock()

	signalAuthManagerAuthenticateId++
	instance := C.gpointer(recv.native)
	handlerID := C.AuthManager_signal_connect_authenticate(instance, C.gpointer(uintptr(signalAuthManagerAuthenticateId)))

	detail := signalAuthManagerAuthenticateDetail{callback, handlerID}
	signalAuthManagerAuthenticateMap[signalAuthManagerAuthenticateId] = detail

	return signalAuthManagerAuthenticateId
}

/*
DisconnectAuthenticate disconnects a callback from the 'authenticate' signal for the AuthManager.

The connectionID should be a value returned from a call to ConnectAuthenticate.
*/
func (recv *AuthManager) DisconnectAuthenticate(connectionID int) {
	signalAuthManagerAuthenticateLock.Lock()
	defer signalAuthManagerAuthenticateLock.Unlock()

	detail, exists := signalAuthManagerAuthenticateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAuthManagerAuthenticateMap, connectionID)
}

//export authmanager_authenticateHandler
func authmanager_authenticateHandler(c_targetObject *C.GObject, c_msg *C.SoupMessage, c_auth *C.SoupAuth, c_retrying C.gboolean, data C.gpointer) {
	signalAuthManagerAuthenticateLock.RLock()
	defer signalAuthManagerAuthenticateLock.RUnlock()

	msg := MessageNewFromC(unsafe.Pointer(c_msg))

	auth := AuthNewFromC(unsafe.Pointer(c_auth))

	retrying := c_retrying == C.TRUE

	targetObject := AuthManagerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalAuthManagerAuthenticateMap[index].callback
	callback(targetObject, msg, auth, retrying)
}

// UseAuth is a wrapper around the C function soup_auth_manager_use_auth.
func (recv *AuthManager) UseAuth(uri *URI, auth *Auth) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_auth := (*C.SoupAuth)(C.NULL)
	if auth != nil {
		c_auth = (*C.SoupAuth)(auth.ToC())
	}

	C.soup_auth_manager_use_auth((*C.SoupAuthManager)(recv.native), c_uri, c_auth)

	return
}

// SessionFeature returns the SessionFeature interface implemented by AuthManager
func (recv *AuthManager) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// Unsupported : AuthNTLM : no CType

// Unsupported : AuthNegotiate : no CType

// Cache is a wrapper around the C record SoupCache.
type Cache struct {
	native *C.SoupCache
	// parent_instance : record
	// priv : record
}

func CacheNewFromC(u unsafe.Pointer) *Cache {
	c := (*C.SoupCache)(u)
	if c == nil {
		return nil
	}

	g := &Cache{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Cache) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Cache) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cache with another Cache, and returns true if they represent the same GObject.
func (recv *Cache) Equals(other *Cache) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Cache) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Cache.
// Exercise care, as this is a potentially dangerous function if the Object is not a Cache.
func CastToCache(object *gobject.Object) *Cache {
	return CacheNewFromC(object.ToC())
}

// CacheNew is a wrapper around the C function soup_cache_new.
func CacheNew(cacheDir string, cacheType CacheType) *Cache {
	c_cache_dir := C.CString(cacheDir)
	defer C.free(unsafe.Pointer(c_cache_dir))

	c_cache_type := (C.SoupCacheType)(cacheType)

	retC := C.soup_cache_new(c_cache_dir, c_cache_type)
	retGo := CacheNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Clear is a wrapper around the C function soup_cache_clear.
func (recv *Cache) Clear() {
	C.soup_cache_clear((*C.SoupCache)(recv.native))

	return
}

// Dump is a wrapper around the C function soup_cache_dump.
func (recv *Cache) Dump() {
	C.soup_cache_dump((*C.SoupCache)(recv.native))

	return
}

// Flush is a wrapper around the C function soup_cache_flush.
func (recv *Cache) Flush() {
	C.soup_cache_flush((*C.SoupCache)(recv.native))

	return
}

// GetMaxSize is a wrapper around the C function soup_cache_get_max_size.
func (recv *Cache) GetMaxSize() uint32 {
	retC := C.soup_cache_get_max_size((*C.SoupCache)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Load is a wrapper around the C function soup_cache_load.
func (recv *Cache) Load() {
	C.soup_cache_load((*C.SoupCache)(recv.native))

	return
}

// SetMaxSize is a wrapper around the C function soup_cache_set_max_size.
func (recv *Cache) SetMaxSize(maxSize uint32) {
	c_max_size := (C.guint)(maxSize)

	C.soup_cache_set_max_size((*C.SoupCache)(recv.native), c_max_size)

	return
}

// SessionFeature returns the SessionFeature interface implemented by Cache
func (recv *Cache) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// ContentDecoder is a wrapper around the C record SoupContentDecoder.
type ContentDecoder struct {
	native *C.SoupContentDecoder
	// parent : record
	// priv : record
}

func ContentDecoderNewFromC(u unsafe.Pointer) *ContentDecoder {
	c := (*C.SoupContentDecoder)(u)
	if c == nil {
		return nil
	}

	g := &ContentDecoder{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContentDecoder) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContentDecoder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentDecoder with another ContentDecoder, and returns true if they represent the same GObject.
func (recv *ContentDecoder) Equals(other *ContentDecoder) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ContentDecoder) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ContentDecoder.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContentDecoder.
func CastToContentDecoder(object *gobject.Object) *ContentDecoder {
	return ContentDecoderNewFromC(object.ToC())
}

// SessionFeature returns the SessionFeature interface implemented by ContentDecoder
func (recv *ContentDecoder) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// ContentSniffer is a wrapper around the C record SoupContentSniffer.
type ContentSniffer struct {
	native *C.SoupContentSniffer
	// parent : record
	// priv : record
}

func ContentSnifferNewFromC(u unsafe.Pointer) *ContentSniffer {
	c := (*C.SoupContentSniffer)(u)
	if c == nil {
		return nil
	}

	g := &ContentSniffer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContentSniffer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContentSniffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentSniffer with another ContentSniffer, and returns true if they represent the same GObject.
func (recv *ContentSniffer) Equals(other *ContentSniffer) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ContentSniffer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ContentSniffer.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContentSniffer.
func CastToContentSniffer(object *gobject.Object) *ContentSniffer {
	return ContentSnifferNewFromC(object.ToC())
}

// ContentSnifferNew is a wrapper around the C function soup_content_sniffer_new.
func ContentSnifferNew() *ContentSniffer {
	retC := C.soup_content_sniffer_new()
	retGo := ContentSnifferNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBufferSize is a wrapper around the C function soup_content_sniffer_get_buffer_size.
func (recv *ContentSniffer) GetBufferSize() uint64 {
	retC := C.soup_content_sniffer_get_buffer_size((*C.SoupContentSniffer)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Sniff is a wrapper around the C function soup_content_sniffer_sniff.
func (recv *ContentSniffer) Sniff(msg *Message, buffer *Buffer) (string, *glib.HashTable) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_buffer := (*C.SoupBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.SoupBuffer)(buffer.ToC())
	}

	var c_params *C.GHashTable

	retC := C.soup_content_sniffer_sniff((*C.SoupContentSniffer)(recv.native), c_msg, c_buffer, &c_params)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	params := glib.HashTableNewFromC(unsafe.Pointer(c_params))

	return retGo, params
}

// SessionFeature returns the SessionFeature interface implemented by ContentSniffer
func (recv *ContentSniffer) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// CookieJar is a wrapper around the C record SoupCookieJar.
type CookieJar struct {
	native *C.SoupCookieJar
	// parent : record
}

func CookieJarNewFromC(u unsafe.Pointer) *CookieJar {
	c := (*C.SoupCookieJar)(u)
	if c == nil {
		return nil
	}

	g := &CookieJar{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CookieJar) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CookieJar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJar with another CookieJar, and returns true if they represent the same GObject.
func (recv *CookieJar) Equals(other *CookieJar) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *CookieJar) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to CookieJar.
// Exercise care, as this is a potentially dangerous function if the Object is not a CookieJar.
func CastToCookieJar(object *gobject.Object) *CookieJar {
	return CookieJarNewFromC(object.ToC())
}

type signalCookieJarChangedDetail struct {
	callback  CookieJarSignalChangedCallback
	handlerID C.gulong
}

var signalCookieJarChangedId int
var signalCookieJarChangedMap = make(map[int]signalCookieJarChangedDetail)
var signalCookieJarChangedLock sync.RWMutex

// CookieJarSignalChangedCallback is a callback function for a 'changed' signal emitted from a CookieJar.
type CookieJarSignalChangedCallback func(targetObject *CookieJar, oldCookie *Cookie, newCookie *Cookie)

/*
ConnectChanged connects the callback to the 'changed' signal for the CookieJar.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *CookieJar) ConnectChanged(callback CookieJarSignalChangedCallback) int {
	signalCookieJarChangedLock.Lock()
	defer signalCookieJarChangedLock.Unlock()

	signalCookieJarChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CookieJar_signal_connect_changed(instance, C.gpointer(uintptr(signalCookieJarChangedId)))

	detail := signalCookieJarChangedDetail{callback, handlerID}
	signalCookieJarChangedMap[signalCookieJarChangedId] = detail

	return signalCookieJarChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the CookieJar.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *CookieJar) DisconnectChanged(connectionID int) {
	signalCookieJarChangedLock.Lock()
	defer signalCookieJarChangedLock.Unlock()

	detail, exists := signalCookieJarChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCookieJarChangedMap, connectionID)
}

//export cookiejar_changedHandler
func cookiejar_changedHandler(c_targetObject *C.GObject, c_old_cookie *C.SoupCookie, c_new_cookie *C.SoupCookie, data C.gpointer) {
	signalCookieJarChangedLock.RLock()
	defer signalCookieJarChangedLock.RUnlock()

	oldCookie := CookieNewFromC(unsafe.Pointer(c_old_cookie))

	newCookie := CookieNewFromC(unsafe.Pointer(c_new_cookie))

	targetObject := CookieJarNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCookieJarChangedMap[index].callback
	callback(targetObject, oldCookie, newCookie)
}

// CookieJarNew is a wrapper around the C function soup_cookie_jar_new.
func CookieJarNew() *CookieJar {
	retC := C.soup_cookie_jar_new()
	retGo := CookieJarNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddCookie is a wrapper around the C function soup_cookie_jar_add_cookie.
func (recv *CookieJar) AddCookie(cookie *Cookie) {
	c_cookie := (*C.SoupCookie)(C.NULL)
	if cookie != nil {
		c_cookie = (*C.SoupCookie)(cookie.ToC())
	}

	C.soup_cookie_jar_add_cookie((*C.SoupCookieJar)(recv.native), c_cookie)

	return
}

// AddCookieWithFirstParty is a wrapper around the C function soup_cookie_jar_add_cookie_with_first_party.
func (recv *CookieJar) AddCookieWithFirstParty(firstParty *URI, cookie *Cookie) {
	c_first_party := (*C.SoupURI)(C.NULL)
	if firstParty != nil {
		c_first_party = (*C.SoupURI)(firstParty.ToC())
	}

	c_cookie := (*C.SoupCookie)(C.NULL)
	if cookie != nil {
		c_cookie = (*C.SoupCookie)(cookie.ToC())
	}

	C.soup_cookie_jar_add_cookie_with_first_party((*C.SoupCookieJar)(recv.native), c_first_party, c_cookie)

	return
}

// AllCookies is a wrapper around the C function soup_cookie_jar_all_cookies.
func (recv *CookieJar) AllCookies() *glib.SList {
	retC := C.soup_cookie_jar_all_cookies((*C.SoupCookieJar)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DeleteCookie is a wrapper around the C function soup_cookie_jar_delete_cookie.
func (recv *CookieJar) DeleteCookie(cookie *Cookie) {
	c_cookie := (*C.SoupCookie)(C.NULL)
	if cookie != nil {
		c_cookie = (*C.SoupCookie)(cookie.ToC())
	}

	C.soup_cookie_jar_delete_cookie((*C.SoupCookieJar)(recv.native), c_cookie)

	return
}

// GetAcceptPolicy is a wrapper around the C function soup_cookie_jar_get_accept_policy.
func (recv *CookieJar) GetAcceptPolicy() CookieJarAcceptPolicy {
	retC := C.soup_cookie_jar_get_accept_policy((*C.SoupCookieJar)(recv.native))
	retGo := (CookieJarAcceptPolicy)(retC)

	return retGo
}

// GetCookieList is a wrapper around the C function soup_cookie_jar_get_cookie_list.
func (recv *CookieJar) GetCookieList(uri *URI, forHttp bool) *glib.SList {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_for_http :=
		boolToGboolean(forHttp)

	retC := C.soup_cookie_jar_get_cookie_list((*C.SoupCookieJar)(recv.native), c_uri, c_for_http)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCookies is a wrapper around the C function soup_cookie_jar_get_cookies.
func (recv *CookieJar) GetCookies(uri *URI, forHttp bool) string {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_for_http :=
		boolToGboolean(forHttp)

	retC := C.soup_cookie_jar_get_cookies((*C.SoupCookieJar)(recv.native), c_uri, c_for_http)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsPersistent is a wrapper around the C function soup_cookie_jar_is_persistent.
func (recv *CookieJar) IsPersistent() bool {
	retC := C.soup_cookie_jar_is_persistent((*C.SoupCookieJar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Save is a wrapper around the C function soup_cookie_jar_save.
func (recv *CookieJar) Save() {
	C.soup_cookie_jar_save((*C.SoupCookieJar)(recv.native))

	return
}

// SetAcceptPolicy is a wrapper around the C function soup_cookie_jar_set_accept_policy.
func (recv *CookieJar) SetAcceptPolicy(policy CookieJarAcceptPolicy) {
	c_policy := (C.SoupCookieJarAcceptPolicy)(policy)

	C.soup_cookie_jar_set_accept_policy((*C.SoupCookieJar)(recv.native), c_policy)

	return
}

// SetCookie is a wrapper around the C function soup_cookie_jar_set_cookie.
func (recv *CookieJar) SetCookie(uri *URI, cookie string) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_cookie := C.CString(cookie)
	defer C.free(unsafe.Pointer(c_cookie))

	C.soup_cookie_jar_set_cookie((*C.SoupCookieJar)(recv.native), c_uri, c_cookie)

	return
}

// SetCookieWithFirstParty is a wrapper around the C function soup_cookie_jar_set_cookie_with_first_party.
func (recv *CookieJar) SetCookieWithFirstParty(uri *URI, firstParty *URI, cookie string) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_first_party := (*C.SoupURI)(C.NULL)
	if firstParty != nil {
		c_first_party = (*C.SoupURI)(firstParty.ToC())
	}

	c_cookie := C.CString(cookie)
	defer C.free(unsafe.Pointer(c_cookie))

	C.soup_cookie_jar_set_cookie_with_first_party((*C.SoupCookieJar)(recv.native), c_uri, c_first_party, c_cookie)

	return
}

// SessionFeature returns the SessionFeature interface implemented by CookieJar
func (recv *CookieJar) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// CookieJarDB is a wrapper around the C record SoupCookieJarDB.
type CookieJarDB struct {
	native *C.SoupCookieJarDB
	// parent : record
}

func CookieJarDBNewFromC(u unsafe.Pointer) *CookieJarDB {
	c := (*C.SoupCookieJarDB)(u)
	if c == nil {
		return nil
	}

	g := &CookieJarDB{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CookieJarDB) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CookieJarDB) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJarDB with another CookieJarDB, and returns true if they represent the same GObject.
func (recv *CookieJarDB) Equals(other *CookieJarDB) bool {
	return other.ToC() == recv.ToC()
}

// CookieJar upcasts to *CookieJar
func (recv *CookieJarDB) CookieJar() *CookieJar {
	return CookieJarNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CookieJarDB) Object() *gobject.Object {
	return recv.CookieJar().Object()
}

// CastToWidget down casts any arbitrary Object to CookieJarDB.
// Exercise care, as this is a potentially dangerous function if the Object is not a CookieJarDB.
func CastToCookieJarDB(object *gobject.Object) *CookieJarDB {
	return CookieJarDBNewFromC(object.ToC())
}

// CookieJarDBNew is a wrapper around the C function soup_cookie_jar_db_new.
func CookieJarDBNew(filename string, readOnly bool) *CookieJarDB {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_read_only :=
		boolToGboolean(readOnly)

	retC := C.soup_cookie_jar_db_new(c_filename, c_read_only)
	retGo := CookieJarDBNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by CookieJarDB
func (recv *CookieJarDB) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// CookieJarText is a wrapper around the C record SoupCookieJarText.
type CookieJarText struct {
	native *C.SoupCookieJarText
	// parent : record
}

func CookieJarTextNewFromC(u unsafe.Pointer) *CookieJarText {
	c := (*C.SoupCookieJarText)(u)
	if c == nil {
		return nil
	}

	g := &CookieJarText{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CookieJarText) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CookieJarText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJarText with another CookieJarText, and returns true if they represent the same GObject.
func (recv *CookieJarText) Equals(other *CookieJarText) bool {
	return other.ToC() == recv.ToC()
}

// CookieJar upcasts to *CookieJar
func (recv *CookieJarText) CookieJar() *CookieJar {
	return CookieJarNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CookieJarText) Object() *gobject.Object {
	return recv.CookieJar().Object()
}

// CastToWidget down casts any arbitrary Object to CookieJarText.
// Exercise care, as this is a potentially dangerous function if the Object is not a CookieJarText.
func CastToCookieJarText(object *gobject.Object) *CookieJarText {
	return CookieJarTextNewFromC(object.ToC())
}

// CookieJarTextNew is a wrapper around the C function soup_cookie_jar_text_new.
func CookieJarTextNew(filename string, readOnly bool) *CookieJarText {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_read_only :=
		boolToGboolean(readOnly)

	retC := C.soup_cookie_jar_text_new(c_filename, c_read_only)
	retGo := CookieJarTextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SessionFeature returns the SessionFeature interface implemented by CookieJarText
func (recv *CookieJarText) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// Logger is a wrapper around the C record SoupLogger.
type Logger struct {
	native *C.SoupLogger
	// parent : record
}

func LoggerNewFromC(u unsafe.Pointer) *Logger {
	c := (*C.SoupLogger)(u)
	if c == nil {
		return nil
	}

	g := &Logger{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Logger) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Logger) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Logger with another Logger, and returns true if they represent the same GObject.
func (recv *Logger) Equals(other *Logger) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Logger) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Logger.
// Exercise care, as this is a potentially dangerous function if the Object is not a Logger.
func CastToLogger(object *gobject.Object) *Logger {
	return LoggerNewFromC(object.ToC())
}

// LoggerNew is a wrapper around the C function soup_logger_new.
func LoggerNew(level LoggerLogLevel, maxBodySize int32) *Logger {
	c_level := (C.SoupLoggerLogLevel)(level)

	c_max_body_size := (C.int)(maxBodySize)

	retC := C.soup_logger_new(c_level, c_max_body_size)
	retGo := LoggerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Attach is a wrapper around the C function soup_logger_attach.
func (recv *Logger) Attach(session *Session) {
	c_session := (*C.SoupSession)(C.NULL)
	if session != nil {
		c_session = (*C.SoupSession)(session.ToC())
	}

	C.soup_logger_attach((*C.SoupLogger)(recv.native), c_session)

	return
}

// Detach is a wrapper around the C function soup_logger_detach.
func (recv *Logger) Detach(session *Session) {
	c_session := (*C.SoupSession)(C.NULL)
	if session != nil {
		c_session = (*C.SoupSession)(session.ToC())
	}

	C.soup_logger_detach((*C.SoupLogger)(recv.native), c_session)

	return
}

// Unsupported : soup_logger_set_printer : unsupported parameter printer : no type generator for LoggerPrinter (SoupLoggerPrinter) for param printer

// Unsupported : soup_logger_set_request_filter : unsupported parameter request_filter : no type generator for LoggerFilter (SoupLoggerFilter) for param request_filter

// Unsupported : soup_logger_set_response_filter : unsupported parameter response_filter : no type generator for LoggerFilter (SoupLoggerFilter) for param response_filter

// SessionFeature returns the SessionFeature interface implemented by Logger
func (recv *Logger) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// Message is a wrapper around the C record SoupMessage.
type Message struct {
	native *C.SoupMessage
	// parent : record
	Method       string
	StatusCode   uint32
	ReasonPhrase string
	// request_body : record
	// request_headers : record
	// response_body : record
	// response_headers : record
}

func MessageNewFromC(u unsafe.Pointer) *Message {
	c := (*C.SoupMessage)(u)
	if c == nil {
		return nil
	}

	g := &Message{
		Method:       C.GoString(c.method),
		ReasonPhrase: C.GoString(c.reason_phrase),
		StatusCode:   (uint32)(c.status_code),
		native:       c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Message) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Message) ToC() unsafe.Pointer {
	recv.native.method =
		C.CString(recv.Method)
	recv.native.status_code =
		(C.guint)(recv.StatusCode)
	recv.native.reason_phrase =
		C.CString(recv.ReasonPhrase)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Message with another Message, and returns true if they represent the same GObject.
func (recv *Message) Equals(other *Message) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Message) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Message.
// Exercise care, as this is a potentially dangerous function if the Object is not a Message.
func CastToMessage(object *gobject.Object) *Message {
	return MessageNewFromC(object.ToC())
}

type signalMessageContentSniffedDetail struct {
	callback  MessageSignalContentSniffedCallback
	handlerID C.gulong
}

var signalMessageContentSniffedId int
var signalMessageContentSniffedMap = make(map[int]signalMessageContentSniffedDetail)
var signalMessageContentSniffedLock sync.RWMutex

// MessageSignalContentSniffedCallback is a callback function for a 'content-sniffed' signal emitted from a Message.
type MessageSignalContentSniffedCallback func(targetObject *Message, type_ string, params *glib.HashTable)

/*
ConnectContentSniffed connects the callback to the 'content-sniffed' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectContentSniffed to remove it.
*/
func (recv *Message) ConnectContentSniffed(callback MessageSignalContentSniffedCallback) int {
	signalMessageContentSniffedLock.Lock()
	defer signalMessageContentSniffedLock.Unlock()

	signalMessageContentSniffedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_content_sniffed(instance, C.gpointer(uintptr(signalMessageContentSniffedId)))

	detail := signalMessageContentSniffedDetail{callback, handlerID}
	signalMessageContentSniffedMap[signalMessageContentSniffedId] = detail

	return signalMessageContentSniffedId
}

/*
DisconnectContentSniffed disconnects a callback from the 'content-sniffed' signal for the Message.

The connectionID should be a value returned from a call to ConnectContentSniffed.
*/
func (recv *Message) DisconnectContentSniffed(connectionID int) {
	signalMessageContentSniffedLock.Lock()
	defer signalMessageContentSniffedLock.Unlock()

	detail, exists := signalMessageContentSniffedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageContentSniffedMap, connectionID)
}

//export message_contentSniffedHandler
func message_contentSniffedHandler(c_targetObject *C.GObject, c_type *C.gchar, c_params *C.GHashTable, data C.gpointer) {
	signalMessageContentSniffedLock.RLock()
	defer signalMessageContentSniffedLock.RUnlock()

	type_ := C.GoString(c_type)

	params := glib.HashTableNewFromC(unsafe.Pointer(c_params))

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageContentSniffedMap[index].callback
	callback(targetObject, type_, params)
}

type signalMessageFinishedDetail struct {
	callback  MessageSignalFinishedCallback
	handlerID C.gulong
}

var signalMessageFinishedId int
var signalMessageFinishedMap = make(map[int]signalMessageFinishedDetail)
var signalMessageFinishedLock sync.RWMutex

// MessageSignalFinishedCallback is a callback function for a 'finished' signal emitted from a Message.
type MessageSignalFinishedCallback func(targetObject *Message)

/*
ConnectFinished connects the callback to the 'finished' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectFinished to remove it.
*/
func (recv *Message) ConnectFinished(callback MessageSignalFinishedCallback) int {
	signalMessageFinishedLock.Lock()
	defer signalMessageFinishedLock.Unlock()

	signalMessageFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_finished(instance, C.gpointer(uintptr(signalMessageFinishedId)))

	detail := signalMessageFinishedDetail{callback, handlerID}
	signalMessageFinishedMap[signalMessageFinishedId] = detail

	return signalMessageFinishedId
}

/*
DisconnectFinished disconnects a callback from the 'finished' signal for the Message.

The connectionID should be a value returned from a call to ConnectFinished.
*/
func (recv *Message) DisconnectFinished(connectionID int) {
	signalMessageFinishedLock.Lock()
	defer signalMessageFinishedLock.Unlock()

	detail, exists := signalMessageFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageFinishedMap, connectionID)
}

//export message_finishedHandler
func message_finishedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageFinishedLock.RLock()
	defer signalMessageFinishedLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageFinishedMap[index].callback
	callback(targetObject)
}

type signalMessageGotBodyDetail struct {
	callback  MessageSignalGotBodyCallback
	handlerID C.gulong
}

var signalMessageGotBodyId int
var signalMessageGotBodyMap = make(map[int]signalMessageGotBodyDetail)
var signalMessageGotBodyLock sync.RWMutex

// MessageSignalGotBodyCallback is a callback function for a 'got-body' signal emitted from a Message.
type MessageSignalGotBodyCallback func(targetObject *Message)

/*
ConnectGotBody connects the callback to the 'got-body' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectGotBody to remove it.
*/
func (recv *Message) ConnectGotBody(callback MessageSignalGotBodyCallback) int {
	signalMessageGotBodyLock.Lock()
	defer signalMessageGotBodyLock.Unlock()

	signalMessageGotBodyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_got_body(instance, C.gpointer(uintptr(signalMessageGotBodyId)))

	detail := signalMessageGotBodyDetail{callback, handlerID}
	signalMessageGotBodyMap[signalMessageGotBodyId] = detail

	return signalMessageGotBodyId
}

/*
DisconnectGotBody disconnects a callback from the 'got-body' signal for the Message.

The connectionID should be a value returned from a call to ConnectGotBody.
*/
func (recv *Message) DisconnectGotBody(connectionID int) {
	signalMessageGotBodyLock.Lock()
	defer signalMessageGotBodyLock.Unlock()

	detail, exists := signalMessageGotBodyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageGotBodyMap, connectionID)
}

//export message_gotBodyHandler
func message_gotBodyHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageGotBodyLock.RLock()
	defer signalMessageGotBodyLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageGotBodyMap[index].callback
	callback(targetObject)
}

type signalMessageGotChunkDetail struct {
	callback  MessageSignalGotChunkCallback
	handlerID C.gulong
}

var signalMessageGotChunkId int
var signalMessageGotChunkMap = make(map[int]signalMessageGotChunkDetail)
var signalMessageGotChunkLock sync.RWMutex

// MessageSignalGotChunkCallback is a callback function for a 'got-chunk' signal emitted from a Message.
type MessageSignalGotChunkCallback func(targetObject *Message, chunk *Buffer)

/*
ConnectGotChunk connects the callback to the 'got-chunk' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectGotChunk to remove it.
*/
func (recv *Message) ConnectGotChunk(callback MessageSignalGotChunkCallback) int {
	signalMessageGotChunkLock.Lock()
	defer signalMessageGotChunkLock.Unlock()

	signalMessageGotChunkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_got_chunk(instance, C.gpointer(uintptr(signalMessageGotChunkId)))

	detail := signalMessageGotChunkDetail{callback, handlerID}
	signalMessageGotChunkMap[signalMessageGotChunkId] = detail

	return signalMessageGotChunkId
}

/*
DisconnectGotChunk disconnects a callback from the 'got-chunk' signal for the Message.

The connectionID should be a value returned from a call to ConnectGotChunk.
*/
func (recv *Message) DisconnectGotChunk(connectionID int) {
	signalMessageGotChunkLock.Lock()
	defer signalMessageGotChunkLock.Unlock()

	detail, exists := signalMessageGotChunkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageGotChunkMap, connectionID)
}

//export message_gotChunkHandler
func message_gotChunkHandler(c_targetObject *C.GObject, c_chunk *C.SoupBuffer, data C.gpointer) {
	signalMessageGotChunkLock.RLock()
	defer signalMessageGotChunkLock.RUnlock()

	chunk := BufferNewFromC(unsafe.Pointer(c_chunk))

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageGotChunkMap[index].callback
	callback(targetObject, chunk)
}

type signalMessageGotHeadersDetail struct {
	callback  MessageSignalGotHeadersCallback
	handlerID C.gulong
}

var signalMessageGotHeadersId int
var signalMessageGotHeadersMap = make(map[int]signalMessageGotHeadersDetail)
var signalMessageGotHeadersLock sync.RWMutex

// MessageSignalGotHeadersCallback is a callback function for a 'got-headers' signal emitted from a Message.
type MessageSignalGotHeadersCallback func(targetObject *Message)

/*
ConnectGotHeaders connects the callback to the 'got-headers' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectGotHeaders to remove it.
*/
func (recv *Message) ConnectGotHeaders(callback MessageSignalGotHeadersCallback) int {
	signalMessageGotHeadersLock.Lock()
	defer signalMessageGotHeadersLock.Unlock()

	signalMessageGotHeadersId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_got_headers(instance, C.gpointer(uintptr(signalMessageGotHeadersId)))

	detail := signalMessageGotHeadersDetail{callback, handlerID}
	signalMessageGotHeadersMap[signalMessageGotHeadersId] = detail

	return signalMessageGotHeadersId
}

/*
DisconnectGotHeaders disconnects a callback from the 'got-headers' signal for the Message.

The connectionID should be a value returned from a call to ConnectGotHeaders.
*/
func (recv *Message) DisconnectGotHeaders(connectionID int) {
	signalMessageGotHeadersLock.Lock()
	defer signalMessageGotHeadersLock.Unlock()

	detail, exists := signalMessageGotHeadersMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageGotHeadersMap, connectionID)
}

//export message_gotHeadersHandler
func message_gotHeadersHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageGotHeadersLock.RLock()
	defer signalMessageGotHeadersLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageGotHeadersMap[index].callback
	callback(targetObject)
}

type signalMessageGotInformationalDetail struct {
	callback  MessageSignalGotInformationalCallback
	handlerID C.gulong
}

var signalMessageGotInformationalId int
var signalMessageGotInformationalMap = make(map[int]signalMessageGotInformationalDetail)
var signalMessageGotInformationalLock sync.RWMutex

// MessageSignalGotInformationalCallback is a callback function for a 'got-informational' signal emitted from a Message.
type MessageSignalGotInformationalCallback func(targetObject *Message)

/*
ConnectGotInformational connects the callback to the 'got-informational' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectGotInformational to remove it.
*/
func (recv *Message) ConnectGotInformational(callback MessageSignalGotInformationalCallback) int {
	signalMessageGotInformationalLock.Lock()
	defer signalMessageGotInformationalLock.Unlock()

	signalMessageGotInformationalId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_got_informational(instance, C.gpointer(uintptr(signalMessageGotInformationalId)))

	detail := signalMessageGotInformationalDetail{callback, handlerID}
	signalMessageGotInformationalMap[signalMessageGotInformationalId] = detail

	return signalMessageGotInformationalId
}

/*
DisconnectGotInformational disconnects a callback from the 'got-informational' signal for the Message.

The connectionID should be a value returned from a call to ConnectGotInformational.
*/
func (recv *Message) DisconnectGotInformational(connectionID int) {
	signalMessageGotInformationalLock.Lock()
	defer signalMessageGotInformationalLock.Unlock()

	detail, exists := signalMessageGotInformationalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageGotInformationalMap, connectionID)
}

//export message_gotInformationalHandler
func message_gotInformationalHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageGotInformationalLock.RLock()
	defer signalMessageGotInformationalLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageGotInformationalMap[index].callback
	callback(targetObject)
}

type signalMessageNetworkEventDetail struct {
	callback  MessageSignalNetworkEventCallback
	handlerID C.gulong
}

var signalMessageNetworkEventId int
var signalMessageNetworkEventMap = make(map[int]signalMessageNetworkEventDetail)
var signalMessageNetworkEventLock sync.RWMutex

// MessageSignalNetworkEventCallback is a callback function for a 'network-event' signal emitted from a Message.
type MessageSignalNetworkEventCallback func(targetObject *Message, event gio.SocketClientEvent, connection *gio.IOStream)

/*
ConnectNetworkEvent connects the callback to the 'network-event' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectNetworkEvent to remove it.
*/
func (recv *Message) ConnectNetworkEvent(callback MessageSignalNetworkEventCallback) int {
	signalMessageNetworkEventLock.Lock()
	defer signalMessageNetworkEventLock.Unlock()

	signalMessageNetworkEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_network_event(instance, C.gpointer(uintptr(signalMessageNetworkEventId)))

	detail := signalMessageNetworkEventDetail{callback, handlerID}
	signalMessageNetworkEventMap[signalMessageNetworkEventId] = detail

	return signalMessageNetworkEventId
}

/*
DisconnectNetworkEvent disconnects a callback from the 'network-event' signal for the Message.

The connectionID should be a value returned from a call to ConnectNetworkEvent.
*/
func (recv *Message) DisconnectNetworkEvent(connectionID int) {
	signalMessageNetworkEventLock.Lock()
	defer signalMessageNetworkEventLock.Unlock()

	detail, exists := signalMessageNetworkEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageNetworkEventMap, connectionID)
}

//export message_networkEventHandler
func message_networkEventHandler(c_targetObject *C.GObject, c_event C.GSocketClientEvent, c_connection *C.GIOStream, data C.gpointer) {
	signalMessageNetworkEventLock.RLock()
	defer signalMessageNetworkEventLock.RUnlock()

	event := gio.SocketClientEvent(c_event)

	connection := gio.IOStreamNewFromC(unsafe.Pointer(c_connection))

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageNetworkEventMap[index].callback
	callback(targetObject, event, connection)
}

type signalMessageRestartedDetail struct {
	callback  MessageSignalRestartedCallback
	handlerID C.gulong
}

var signalMessageRestartedId int
var signalMessageRestartedMap = make(map[int]signalMessageRestartedDetail)
var signalMessageRestartedLock sync.RWMutex

// MessageSignalRestartedCallback is a callback function for a 'restarted' signal emitted from a Message.
type MessageSignalRestartedCallback func(targetObject *Message)

/*
ConnectRestarted connects the callback to the 'restarted' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectRestarted to remove it.
*/
func (recv *Message) ConnectRestarted(callback MessageSignalRestartedCallback) int {
	signalMessageRestartedLock.Lock()
	defer signalMessageRestartedLock.Unlock()

	signalMessageRestartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_restarted(instance, C.gpointer(uintptr(signalMessageRestartedId)))

	detail := signalMessageRestartedDetail{callback, handlerID}
	signalMessageRestartedMap[signalMessageRestartedId] = detail

	return signalMessageRestartedId
}

/*
DisconnectRestarted disconnects a callback from the 'restarted' signal for the Message.

The connectionID should be a value returned from a call to ConnectRestarted.
*/
func (recv *Message) DisconnectRestarted(connectionID int) {
	signalMessageRestartedLock.Lock()
	defer signalMessageRestartedLock.Unlock()

	detail, exists := signalMessageRestartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageRestartedMap, connectionID)
}

//export message_restartedHandler
func message_restartedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageRestartedLock.RLock()
	defer signalMessageRestartedLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageRestartedMap[index].callback
	callback(targetObject)
}

type signalMessageStartingDetail struct {
	callback  MessageSignalStartingCallback
	handlerID C.gulong
}

var signalMessageStartingId int
var signalMessageStartingMap = make(map[int]signalMessageStartingDetail)
var signalMessageStartingLock sync.RWMutex

// MessageSignalStartingCallback is a callback function for a 'starting' signal emitted from a Message.
type MessageSignalStartingCallback func(targetObject *Message)

/*
ConnectStarting connects the callback to the 'starting' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectStarting to remove it.
*/
func (recv *Message) ConnectStarting(callback MessageSignalStartingCallback) int {
	signalMessageStartingLock.Lock()
	defer signalMessageStartingLock.Unlock()

	signalMessageStartingId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_starting(instance, C.gpointer(uintptr(signalMessageStartingId)))

	detail := signalMessageStartingDetail{callback, handlerID}
	signalMessageStartingMap[signalMessageStartingId] = detail

	return signalMessageStartingId
}

/*
DisconnectStarting disconnects a callback from the 'starting' signal for the Message.

The connectionID should be a value returned from a call to ConnectStarting.
*/
func (recv *Message) DisconnectStarting(connectionID int) {
	signalMessageStartingLock.Lock()
	defer signalMessageStartingLock.Unlock()

	detail, exists := signalMessageStartingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageStartingMap, connectionID)
}

//export message_startingHandler
func message_startingHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageStartingLock.RLock()
	defer signalMessageStartingLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageStartingMap[index].callback
	callback(targetObject)
}

type signalMessageWroteBodyDetail struct {
	callback  MessageSignalWroteBodyCallback
	handlerID C.gulong
}

var signalMessageWroteBodyId int
var signalMessageWroteBodyMap = make(map[int]signalMessageWroteBodyDetail)
var signalMessageWroteBodyLock sync.RWMutex

// MessageSignalWroteBodyCallback is a callback function for a 'wrote-body' signal emitted from a Message.
type MessageSignalWroteBodyCallback func(targetObject *Message)

/*
ConnectWroteBody connects the callback to the 'wrote-body' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectWroteBody to remove it.
*/
func (recv *Message) ConnectWroteBody(callback MessageSignalWroteBodyCallback) int {
	signalMessageWroteBodyLock.Lock()
	defer signalMessageWroteBodyLock.Unlock()

	signalMessageWroteBodyId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_wrote_body(instance, C.gpointer(uintptr(signalMessageWroteBodyId)))

	detail := signalMessageWroteBodyDetail{callback, handlerID}
	signalMessageWroteBodyMap[signalMessageWroteBodyId] = detail

	return signalMessageWroteBodyId
}

/*
DisconnectWroteBody disconnects a callback from the 'wrote-body' signal for the Message.

The connectionID should be a value returned from a call to ConnectWroteBody.
*/
func (recv *Message) DisconnectWroteBody(connectionID int) {
	signalMessageWroteBodyLock.Lock()
	defer signalMessageWroteBodyLock.Unlock()

	detail, exists := signalMessageWroteBodyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageWroteBodyMap, connectionID)
}

//export message_wroteBodyHandler
func message_wroteBodyHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageWroteBodyLock.RLock()
	defer signalMessageWroteBodyLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageWroteBodyMap[index].callback
	callback(targetObject)
}

type signalMessageWroteBodyDataDetail struct {
	callback  MessageSignalWroteBodyDataCallback
	handlerID C.gulong
}

var signalMessageWroteBodyDataId int
var signalMessageWroteBodyDataMap = make(map[int]signalMessageWroteBodyDataDetail)
var signalMessageWroteBodyDataLock sync.RWMutex

// MessageSignalWroteBodyDataCallback is a callback function for a 'wrote-body-data' signal emitted from a Message.
type MessageSignalWroteBodyDataCallback func(targetObject *Message, chunk *Buffer)

/*
ConnectWroteBodyData connects the callback to the 'wrote-body-data' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectWroteBodyData to remove it.
*/
func (recv *Message) ConnectWroteBodyData(callback MessageSignalWroteBodyDataCallback) int {
	signalMessageWroteBodyDataLock.Lock()
	defer signalMessageWroteBodyDataLock.Unlock()

	signalMessageWroteBodyDataId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_wrote_body_data(instance, C.gpointer(uintptr(signalMessageWroteBodyDataId)))

	detail := signalMessageWroteBodyDataDetail{callback, handlerID}
	signalMessageWroteBodyDataMap[signalMessageWroteBodyDataId] = detail

	return signalMessageWroteBodyDataId
}

/*
DisconnectWroteBodyData disconnects a callback from the 'wrote-body-data' signal for the Message.

The connectionID should be a value returned from a call to ConnectWroteBodyData.
*/
func (recv *Message) DisconnectWroteBodyData(connectionID int) {
	signalMessageWroteBodyDataLock.Lock()
	defer signalMessageWroteBodyDataLock.Unlock()

	detail, exists := signalMessageWroteBodyDataMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageWroteBodyDataMap, connectionID)
}

//export message_wroteBodyDataHandler
func message_wroteBodyDataHandler(c_targetObject *C.GObject, c_chunk *C.SoupBuffer, data C.gpointer) {
	signalMessageWroteBodyDataLock.RLock()
	defer signalMessageWroteBodyDataLock.RUnlock()

	chunk := BufferNewFromC(unsafe.Pointer(c_chunk))

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageWroteBodyDataMap[index].callback
	callback(targetObject, chunk)
}

type signalMessageWroteChunkDetail struct {
	callback  MessageSignalWroteChunkCallback
	handlerID C.gulong
}

var signalMessageWroteChunkId int
var signalMessageWroteChunkMap = make(map[int]signalMessageWroteChunkDetail)
var signalMessageWroteChunkLock sync.RWMutex

// MessageSignalWroteChunkCallback is a callback function for a 'wrote-chunk' signal emitted from a Message.
type MessageSignalWroteChunkCallback func(targetObject *Message)

/*
ConnectWroteChunk connects the callback to the 'wrote-chunk' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectWroteChunk to remove it.
*/
func (recv *Message) ConnectWroteChunk(callback MessageSignalWroteChunkCallback) int {
	signalMessageWroteChunkLock.Lock()
	defer signalMessageWroteChunkLock.Unlock()

	signalMessageWroteChunkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_wrote_chunk(instance, C.gpointer(uintptr(signalMessageWroteChunkId)))

	detail := signalMessageWroteChunkDetail{callback, handlerID}
	signalMessageWroteChunkMap[signalMessageWroteChunkId] = detail

	return signalMessageWroteChunkId
}

/*
DisconnectWroteChunk disconnects a callback from the 'wrote-chunk' signal for the Message.

The connectionID should be a value returned from a call to ConnectWroteChunk.
*/
func (recv *Message) DisconnectWroteChunk(connectionID int) {
	signalMessageWroteChunkLock.Lock()
	defer signalMessageWroteChunkLock.Unlock()

	detail, exists := signalMessageWroteChunkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageWroteChunkMap, connectionID)
}

//export message_wroteChunkHandler
func message_wroteChunkHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageWroteChunkLock.RLock()
	defer signalMessageWroteChunkLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageWroteChunkMap[index].callback
	callback(targetObject)
}

type signalMessageWroteHeadersDetail struct {
	callback  MessageSignalWroteHeadersCallback
	handlerID C.gulong
}

var signalMessageWroteHeadersId int
var signalMessageWroteHeadersMap = make(map[int]signalMessageWroteHeadersDetail)
var signalMessageWroteHeadersLock sync.RWMutex

// MessageSignalWroteHeadersCallback is a callback function for a 'wrote-headers' signal emitted from a Message.
type MessageSignalWroteHeadersCallback func(targetObject *Message)

/*
ConnectWroteHeaders connects the callback to the 'wrote-headers' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectWroteHeaders to remove it.
*/
func (recv *Message) ConnectWroteHeaders(callback MessageSignalWroteHeadersCallback) int {
	signalMessageWroteHeadersLock.Lock()
	defer signalMessageWroteHeadersLock.Unlock()

	signalMessageWroteHeadersId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_wrote_headers(instance, C.gpointer(uintptr(signalMessageWroteHeadersId)))

	detail := signalMessageWroteHeadersDetail{callback, handlerID}
	signalMessageWroteHeadersMap[signalMessageWroteHeadersId] = detail

	return signalMessageWroteHeadersId
}

/*
DisconnectWroteHeaders disconnects a callback from the 'wrote-headers' signal for the Message.

The connectionID should be a value returned from a call to ConnectWroteHeaders.
*/
func (recv *Message) DisconnectWroteHeaders(connectionID int) {
	signalMessageWroteHeadersLock.Lock()
	defer signalMessageWroteHeadersLock.Unlock()

	detail, exists := signalMessageWroteHeadersMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageWroteHeadersMap, connectionID)
}

//export message_wroteHeadersHandler
func message_wroteHeadersHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageWroteHeadersLock.RLock()
	defer signalMessageWroteHeadersLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageWroteHeadersMap[index].callback
	callback(targetObject)
}

type signalMessageWroteInformationalDetail struct {
	callback  MessageSignalWroteInformationalCallback
	handlerID C.gulong
}

var signalMessageWroteInformationalId int
var signalMessageWroteInformationalMap = make(map[int]signalMessageWroteInformationalDetail)
var signalMessageWroteInformationalLock sync.RWMutex

// MessageSignalWroteInformationalCallback is a callback function for a 'wrote-informational' signal emitted from a Message.
type MessageSignalWroteInformationalCallback func(targetObject *Message)

/*
ConnectWroteInformational connects the callback to the 'wrote-informational' signal for the Message.

The returned value represents the connection, and may be passed to DisconnectWroteInformational to remove it.
*/
func (recv *Message) ConnectWroteInformational(callback MessageSignalWroteInformationalCallback) int {
	signalMessageWroteInformationalLock.Lock()
	defer signalMessageWroteInformationalLock.Unlock()

	signalMessageWroteInformationalId++
	instance := C.gpointer(recv.native)
	handlerID := C.Message_signal_connect_wrote_informational(instance, C.gpointer(uintptr(signalMessageWroteInformationalId)))

	detail := signalMessageWroteInformationalDetail{callback, handlerID}
	signalMessageWroteInformationalMap[signalMessageWroteInformationalId] = detail

	return signalMessageWroteInformationalId
}

/*
DisconnectWroteInformational disconnects a callback from the 'wrote-informational' signal for the Message.

The connectionID should be a value returned from a call to ConnectWroteInformational.
*/
func (recv *Message) DisconnectWroteInformational(connectionID int) {
	signalMessageWroteInformationalLock.Lock()
	defer signalMessageWroteInformationalLock.Unlock()

	detail, exists := signalMessageWroteInformationalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMessageWroteInformationalMap, connectionID)
}

//export message_wroteInformationalHandler
func message_wroteInformationalHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalMessageWroteInformationalLock.RLock()
	defer signalMessageWroteInformationalLock.RUnlock()

	targetObject := MessageNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalMessageWroteInformationalMap[index].callback
	callback(targetObject)
}

// MessageNew is a wrapper around the C function soup_message_new.
func MessageNew(method string, uriString string) *Message {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri_string := C.CString(uriString)
	defer C.free(unsafe.Pointer(c_uri_string))

	retC := C.soup_message_new(c_method, c_uri_string)
	var retGo (*Message)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MessageNewFromC(unsafe.Pointer(retC))
	}

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// MessageNewFromUri is a wrapper around the C function soup_message_new_from_uri.
func MessageNewFromUri(method string, uri *URI) *Message {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	retC := C.soup_message_new_from_uri(c_method, c_uri)
	retGo := MessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : soup_message_add_header_handler : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Unsupported : soup_message_add_status_code_handler : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// ContentSniffed is a wrapper around the C function soup_message_content_sniffed.
func (recv *Message) ContentSniffed(contentType string, params *glib.HashTable) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_params := (*C.GHashTable)(C.NULL)
	if params != nil {
		c_params = (*C.GHashTable)(params.ToC())
	}

	C.soup_message_content_sniffed((*C.SoupMessage)(recv.native), c_content_type, c_params)

	return
}

// DisableFeature is a wrapper around the C function soup_message_disable_feature.
func (recv *Message) DisableFeature(featureType gobject.Type) {
	c_feature_type := (C.GType)(featureType)

	C.soup_message_disable_feature((*C.SoupMessage)(recv.native), c_feature_type)

	return
}

// Finished is a wrapper around the C function soup_message_finished.
func (recv *Message) Finished() {
	C.soup_message_finished((*C.SoupMessage)(recv.native))

	return
}

// GetAddress is a wrapper around the C function soup_message_get_address.
func (recv *Message) GetAddress() *Address {
	retC := C.soup_message_get_address((*C.SoupMessage)(recv.native))
	retGo := AddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFirstParty is a wrapper around the C function soup_message_get_first_party.
func (recv *Message) GetFirstParty() *URI {
	retC := C.soup_message_get_first_party((*C.SoupMessage)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFlags is a wrapper around the C function soup_message_get_flags.
func (recv *Message) GetFlags() MessageFlags {
	retC := C.soup_message_get_flags((*C.SoupMessage)(recv.native))
	retGo := (MessageFlags)(retC)

	return retGo
}

// GetHttpVersion is a wrapper around the C function soup_message_get_http_version.
func (recv *Message) GetHttpVersion() HTTPVersion {
	retC := C.soup_message_get_http_version((*C.SoupMessage)(recv.native))
	retGo := (HTTPVersion)(retC)

	return retGo
}

// GetHttpsStatus is a wrapper around the C function soup_message_get_https_status.
func (recv *Message) GetHttpsStatus() (bool, *gio.TlsCertificate, gio.TlsCertificateFlags) {
	var c_certificate *C.GTlsCertificate

	var c_errors C.GTlsCertificateFlags

	retC := C.soup_message_get_https_status((*C.SoupMessage)(recv.native), &c_certificate, &c_errors)
	retGo := retC == C.TRUE

	certificate := gio.TlsCertificateNewFromC(unsafe.Pointer(c_certificate))

	errors := (gio.TlsCertificateFlags)(c_errors)

	return retGo, certificate, errors
}

// GetPriority is a wrapper around the C function soup_message_get_priority.
func (recv *Message) GetPriority() MessagePriority {
	retC := C.soup_message_get_priority((*C.SoupMessage)(recv.native))
	retGo := (MessagePriority)(retC)

	return retGo
}

// GetSoupRequest is a wrapper around the C function soup_message_get_soup_request.
func (recv *Message) GetSoupRequest() *Request {
	retC := C.soup_message_get_soup_request((*C.SoupMessage)(recv.native))
	retGo := RequestNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function soup_message_get_uri.
func (recv *Message) GetUri() *URI {
	retC := C.soup_message_get_uri((*C.SoupMessage)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// GotBody is a wrapper around the C function soup_message_got_body.
func (recv *Message) GotBody() {
	C.soup_message_got_body((*C.SoupMessage)(recv.native))

	return
}

// GotChunk is a wrapper around the C function soup_message_got_chunk.
func (recv *Message) GotChunk(chunk *Buffer) {
	c_chunk := (*C.SoupBuffer)(C.NULL)
	if chunk != nil {
		c_chunk = (*C.SoupBuffer)(chunk.ToC())
	}

	C.soup_message_got_chunk((*C.SoupMessage)(recv.native), c_chunk)

	return
}

// GotHeaders is a wrapper around the C function soup_message_got_headers.
func (recv *Message) GotHeaders() {
	C.soup_message_got_headers((*C.SoupMessage)(recv.native))

	return
}

// GotInformational is a wrapper around the C function soup_message_got_informational.
func (recv *Message) GotInformational() {
	C.soup_message_got_informational((*C.SoupMessage)(recv.native))

	return
}

// IsKeepalive is a wrapper around the C function soup_message_is_keepalive.
func (recv *Message) IsKeepalive() bool {
	retC := C.soup_message_is_keepalive((*C.SoupMessage)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Restarted is a wrapper around the C function soup_message_restarted.
func (recv *Message) Restarted() {
	C.soup_message_restarted((*C.SoupMessage)(recv.native))

	return
}

// Unsupported : soup_message_set_chunk_allocator : unsupported parameter allocator : no type generator for ChunkAllocator (SoupChunkAllocator) for param allocator

// SetFirstParty is a wrapper around the C function soup_message_set_first_party.
func (recv *Message) SetFirstParty(firstParty *URI) {
	c_first_party := (*C.SoupURI)(C.NULL)
	if firstParty != nil {
		c_first_party = (*C.SoupURI)(firstParty.ToC())
	}

	C.soup_message_set_first_party((*C.SoupMessage)(recv.native), c_first_party)

	return
}

// SetFlags is a wrapper around the C function soup_message_set_flags.
func (recv *Message) SetFlags(flags MessageFlags) {
	c_flags := (C.SoupMessageFlags)(flags)

	C.soup_message_set_flags((*C.SoupMessage)(recv.native), c_flags)

	return
}

// SetHttpVersion is a wrapper around the C function soup_message_set_http_version.
func (recv *Message) SetHttpVersion(version HTTPVersion) {
	c_version := (C.SoupHTTPVersion)(version)

	C.soup_message_set_http_version((*C.SoupMessage)(recv.native), c_version)

	return
}

// SetPriority is a wrapper around the C function soup_message_set_priority.
func (recv *Message) SetPriority(priority MessagePriority) {
	c_priority := (C.SoupMessagePriority)(priority)

	C.soup_message_set_priority((*C.SoupMessage)(recv.native), c_priority)

	return
}

// SetRedirect is a wrapper around the C function soup_message_set_redirect.
func (recv *Message) SetRedirect(statusCode uint32, redirectUri string) {
	c_status_code := (C.guint)(statusCode)

	c_redirect_uri := C.CString(redirectUri)
	defer C.free(unsafe.Pointer(c_redirect_uri))

	C.soup_message_set_redirect((*C.SoupMessage)(recv.native), c_status_code, c_redirect_uri)

	return
}

// SetRequest is a wrapper around the C function soup_message_set_request.
func (recv *Message) SetRequest(contentType string, reqUse MemoryUse, reqBody []uint8) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_req_use := (C.SoupMemoryUse)(reqUse)

	c_req_body_array := make([]C.guint8, len(reqBody)+1, len(reqBody)+1)
	for i, item := range reqBody {
		c := (C.guint8)(item)
		c_req_body_array[i] = c
	}
	c_req_body_array[len(reqBody)] = 0
	c_req_body_arrayPtr := &c_req_body_array[0]
	c_req_body := (*C.char)(unsafe.Pointer(c_req_body_arrayPtr))

	c_req_length := (C.gsize)(len(reqBody))

	C.soup_message_set_request((*C.SoupMessage)(recv.native), c_content_type, c_req_use, c_req_body, c_req_length)

	return
}

// SetResponse is a wrapper around the C function soup_message_set_response.
func (recv *Message) SetResponse(contentType string, respUse MemoryUse, respBody []uint8) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_resp_use := (C.SoupMemoryUse)(respUse)

	c_resp_body_array := make([]C.guint8, len(respBody)+1, len(respBody)+1)
	for i, item := range respBody {
		c := (C.guint8)(item)
		c_resp_body_array[i] = c
	}
	c_resp_body_array[len(respBody)] = 0
	c_resp_body_arrayPtr := &c_resp_body_array[0]
	c_resp_body := (*C.char)(unsafe.Pointer(c_resp_body_arrayPtr))

	c_resp_length := (C.gsize)(len(respBody))

	C.soup_message_set_response((*C.SoupMessage)(recv.native), c_content_type, c_resp_use, c_resp_body, c_resp_length)

	return
}

// SetStatus is a wrapper around the C function soup_message_set_status.
func (recv *Message) SetStatus(statusCode uint32) {
	c_status_code := (C.guint)(statusCode)

	C.soup_message_set_status((*C.SoupMessage)(recv.native), c_status_code)

	return
}

// SetStatusFull is a wrapper around the C function soup_message_set_status_full.
func (recv *Message) SetStatusFull(statusCode uint32, reasonPhrase string) {
	c_status_code := (C.guint)(statusCode)

	c_reason_phrase := C.CString(reasonPhrase)
	defer C.free(unsafe.Pointer(c_reason_phrase))

	C.soup_message_set_status_full((*C.SoupMessage)(recv.native), c_status_code, c_reason_phrase)

	return
}

// SetUri is a wrapper around the C function soup_message_set_uri.
func (recv *Message) SetUri(uri *URI) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	C.soup_message_set_uri((*C.SoupMessage)(recv.native), c_uri)

	return
}

// Starting is a wrapper around the C function soup_message_starting.
func (recv *Message) Starting() {
	C.soup_message_starting((*C.SoupMessage)(recv.native))

	return
}

// WroteBody is a wrapper around the C function soup_message_wrote_body.
func (recv *Message) WroteBody() {
	C.soup_message_wrote_body((*C.SoupMessage)(recv.native))

	return
}

// WroteBodyData is a wrapper around the C function soup_message_wrote_body_data.
func (recv *Message) WroteBodyData(chunk *Buffer) {
	c_chunk := (*C.SoupBuffer)(C.NULL)
	if chunk != nil {
		c_chunk = (*C.SoupBuffer)(chunk.ToC())
	}

	C.soup_message_wrote_body_data((*C.SoupMessage)(recv.native), c_chunk)

	return
}

// WroteChunk is a wrapper around the C function soup_message_wrote_chunk.
func (recv *Message) WroteChunk() {
	C.soup_message_wrote_chunk((*C.SoupMessage)(recv.native))

	return
}

// WroteHeaders is a wrapper around the C function soup_message_wrote_headers.
func (recv *Message) WroteHeaders() {
	C.soup_message_wrote_headers((*C.SoupMessage)(recv.native))

	return
}

// WroteInformational is a wrapper around the C function soup_message_wrote_informational.
func (recv *Message) WroteInformational() {
	C.soup_message_wrote_informational((*C.SoupMessage)(recv.native))

	return
}

// MultipartInputStream is a wrapper around the C record SoupMultipartInputStream.
type MultipartInputStream struct {
	native *C.SoupMultipartInputStream
	// parent_instance : record
	// Private : priv
}

func MultipartInputStreamNewFromC(u unsafe.Pointer) *MultipartInputStream {
	c := (*C.SoupMultipartInputStream)(u)
	if c == nil {
		return nil
	}

	g := &MultipartInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MultipartInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MultipartInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MultipartInputStream with another MultipartInputStream, and returns true if they represent the same GObject.
func (recv *MultipartInputStream) Equals(other *MultipartInputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *MultipartInputStream) FilterInputStream() *gio.FilterInputStream {
	return gio.FilterInputStreamNewFromC(unsafe.Pointer(recv.native))
}

// InputStream upcasts to *InputStream
func (recv *MultipartInputStream) InputStream() *gio.InputStream {
	return recv.FilterInputStream().InputStream()
}

// Object upcasts to *Object
func (recv *MultipartInputStream) Object() *gobject.Object {
	return recv.FilterInputStream().Object()
}

// CastToWidget down casts any arbitrary Object to MultipartInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a MultipartInputStream.
func CastToMultipartInputStream(object *gobject.Object) *MultipartInputStream {
	return MultipartInputStreamNewFromC(object.ToC())
}

// MultipartInputStreamNew is a wrapper around the C function soup_multipart_input_stream_new.
func MultipartInputStreamNew(msg *Message, baseStream *gio.InputStream) *MultipartInputStream {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.soup_multipart_input_stream_new(c_msg, c_base_stream)
	retGo := MultipartInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetHeaders is a wrapper around the C function soup_multipart_input_stream_get_headers.
func (recv *MultipartInputStream) GetHeaders() *MessageHeaders {
	retC := C.soup_multipart_input_stream_get_headers((*C.SoupMultipartInputStream)(recv.native))
	var retGo (*MessageHeaders)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MessageHeadersNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// NextPart is a wrapper around the C function soup_multipart_input_stream_next_part.
func (recv *MultipartInputStream) NextPart(cancellable *gio.Cancellable) (*gio.InputStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_multipart_input_stream_next_part((*C.SoupMultipartInputStream)(recv.native), c_cancellable, &cThrowableError)
	var retGo (*gio.InputStream)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.InputStreamNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : soup_multipart_input_stream_next_part_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// NextPartFinish is a wrapper around the C function soup_multipart_input_stream_next_part_finish.
func (recv *MultipartInputStream) NextPartFinish(result *gio.AsyncResult) (*gio.InputStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.soup_multipart_input_stream_next_part_finish((*C.SoupMultipartInputStream)(recv.native), c_result, &cThrowableError)
	var retGo (*gio.InputStream)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.InputStreamNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PollableInputStream returns the PollableInputStream interface implemented by MultipartInputStream
func (recv *MultipartInputStream) PollableInputStream() *gio.PollableInputStream {
	return gio.PollableInputStreamNewFromC(recv.ToC())
}

// ProxyResolverDefault is a wrapper around the C record SoupProxyResolverDefault.
type ProxyResolverDefault struct {
	native *C.SoupProxyResolverDefault
	// parent : record
}

func ProxyResolverDefaultNewFromC(u unsafe.Pointer) *ProxyResolverDefault {
	c := (*C.SoupProxyResolverDefault)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolverDefault{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyResolverDefault) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyResolverDefault) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyResolverDefault with another ProxyResolverDefault, and returns true if they represent the same GObject.
func (recv *ProxyResolverDefault) Equals(other *ProxyResolverDefault) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ProxyResolverDefault) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ProxyResolverDefault.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProxyResolverDefault.
func CastToProxyResolverDefault(object *gobject.Object) *ProxyResolverDefault {
	return ProxyResolverDefaultNewFromC(object.ToC())
}

// ProxyURIResolver returns the ProxyURIResolver interface implemented by ProxyResolverDefault
func (recv *ProxyResolverDefault) ProxyURIResolver() *ProxyURIResolver {
	return ProxyURIResolverNewFromC(recv.ToC())
}

// SessionFeature returns the SessionFeature interface implemented by ProxyResolverDefault
func (recv *ProxyResolverDefault) SessionFeature() *SessionFeature {
	return SessionFeatureNewFromC(recv.ToC())
}

// Request is a wrapper around the C record SoupRequest.
type Request struct {
	native *C.SoupRequest
	// parent : record
	// priv : record
}

func RequestNewFromC(u unsafe.Pointer) *Request {
	c := (*C.SoupRequest)(u)
	if c == nil {
		return nil
	}

	g := &Request{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Request) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Request) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Request with another Request, and returns true if they represent the same GObject.
func (recv *Request) Equals(other *Request) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Request) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Request.
// Exercise care, as this is a potentially dangerous function if the Object is not a Request.
func CastToRequest(object *gobject.Object) *Request {
	return RequestNewFromC(object.ToC())
}

// GetContentLength is a wrapper around the C function soup_request_get_content_length.
func (recv *Request) GetContentLength() int64 {
	retC := C.soup_request_get_content_length((*C.SoupRequest)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetContentType is a wrapper around the C function soup_request_get_content_type.
func (recv *Request) GetContentType() string {
	retC := C.soup_request_get_content_type((*C.SoupRequest)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSession is a wrapper around the C function soup_request_get_session.
func (recv *Request) GetSession() *Session {
	retC := C.soup_request_get_session((*C.SoupRequest)(recv.native))
	retGo := SessionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function soup_request_get_uri.
func (recv *Request) GetUri() *URI {
	retC := C.soup_request_get_uri((*C.SoupRequest)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// Send is a wrapper around the C function soup_request_send.
func (recv *Request) Send(cancellable *gio.Cancellable) (*gio.InputStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_request_send((*C.SoupRequest)(recv.native), c_cancellable, &cThrowableError)
	retGo := gio.InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : soup_request_send_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SendFinish is a wrapper around the C function soup_request_send_finish.
func (recv *Request) SendFinish(result *gio.AsyncResult) (*gio.InputStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.soup_request_send_finish((*C.SoupRequest)(recv.native), c_result, &cThrowableError)
	retGo := gio.InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Initable returns the Initable interface implemented by Request
func (recv *Request) Initable() *gio.Initable {
	return gio.InitableNewFromC(recv.ToC())
}

// RequestData is a wrapper around the C record SoupRequestData.
type RequestData struct {
	native *C.SoupRequestData
	// parent : record
	// priv : record
}

func RequestDataNewFromC(u unsafe.Pointer) *RequestData {
	c := (*C.SoupRequestData)(u)
	if c == nil {
		return nil
	}

	g := &RequestData{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RequestData) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RequestData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestData with another RequestData, and returns true if they represent the same GObject.
func (recv *RequestData) Equals(other *RequestData) bool {
	return other.ToC() == recv.ToC()
}

// Request upcasts to *Request
func (recv *RequestData) Request() *Request {
	return RequestNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *RequestData) Object() *gobject.Object {
	return recv.Request().Object()
}

// CastToWidget down casts any arbitrary Object to RequestData.
// Exercise care, as this is a potentially dangerous function if the Object is not a RequestData.
func CastToRequestData(object *gobject.Object) *RequestData {
	return RequestDataNewFromC(object.ToC())
}

// Initable returns the Initable interface implemented by RequestData
func (recv *RequestData) Initable() *gio.Initable {
	return gio.InitableNewFromC(recv.ToC())
}

// RequestFile is a wrapper around the C record SoupRequestFile.
type RequestFile struct {
	native *C.SoupRequestFile
	// parent : record
	// priv : record
}

func RequestFileNewFromC(u unsafe.Pointer) *RequestFile {
	c := (*C.SoupRequestFile)(u)
	if c == nil {
		return nil
	}

	g := &RequestFile{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RequestFile) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RequestFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestFile with another RequestFile, and returns true if they represent the same GObject.
func (recv *RequestFile) Equals(other *RequestFile) bool {
	return other.ToC() == recv.ToC()
}

// Request upcasts to *Request
func (recv *RequestFile) Request() *Request {
	return RequestNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *RequestFile) Object() *gobject.Object {
	return recv.Request().Object()
}

// CastToWidget down casts any arbitrary Object to RequestFile.
// Exercise care, as this is a potentially dangerous function if the Object is not a RequestFile.
func CastToRequestFile(object *gobject.Object) *RequestFile {
	return RequestFileNewFromC(object.ToC())
}

// GetFile is a wrapper around the C function soup_request_file_get_file.
func (recv *RequestFile) GetFile() *gio.File {
	retC := C.soup_request_file_get_file((*C.SoupRequestFile)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Initable returns the Initable interface implemented by RequestFile
func (recv *RequestFile) Initable() *gio.Initable {
	return gio.InitableNewFromC(recv.ToC())
}

// RequestHTTP is a wrapper around the C record SoupRequestHTTP.
type RequestHTTP struct {
	native *C.SoupRequestHTTP
	// parent : record
	// priv : record
}

func RequestHTTPNewFromC(u unsafe.Pointer) *RequestHTTP {
	c := (*C.SoupRequestHTTP)(u)
	if c == nil {
		return nil
	}

	g := &RequestHTTP{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RequestHTTP) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RequestHTTP) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestHTTP with another RequestHTTP, and returns true if they represent the same GObject.
func (recv *RequestHTTP) Equals(other *RequestHTTP) bool {
	return other.ToC() == recv.ToC()
}

// Request upcasts to *Request
func (recv *RequestHTTP) Request() *Request {
	return RequestNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *RequestHTTP) Object() *gobject.Object {
	return recv.Request().Object()
}

// CastToWidget down casts any arbitrary Object to RequestHTTP.
// Exercise care, as this is a potentially dangerous function if the Object is not a RequestHTTP.
func CastToRequestHTTP(object *gobject.Object) *RequestHTTP {
	return RequestHTTPNewFromC(object.ToC())
}

// GetMessage is a wrapper around the C function soup_request_http_get_message.
func (recv *RequestHTTP) GetMessage() *Message {
	retC := C.soup_request_http_get_message((*C.SoupRequestHTTP)(recv.native))
	retGo := MessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Initable returns the Initable interface implemented by RequestHTTP
func (recv *RequestHTTP) Initable() *gio.Initable {
	return gio.InitableNewFromC(recv.ToC())
}

// Blacklisted : SoupRequester

// Server is a wrapper around the C record SoupServer.
type Server struct {
	native *C.SoupServer
	// parent : record
}

func ServerNewFromC(u unsafe.Pointer) *Server {
	c := (*C.SoupServer)(u)
	if c == nil {
		return nil
	}

	g := &Server{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Server) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Server) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Server with another Server, and returns true if they represent the same GObject.
func (recv *Server) Equals(other *Server) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Server) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Server.
// Exercise care, as this is a potentially dangerous function if the Object is not a Server.
func CastToServer(object *gobject.Object) *Server {
	return ServerNewFromC(object.ToC())
}

type signalServerRequestAbortedDetail struct {
	callback  ServerSignalRequestAbortedCallback
	handlerID C.gulong
}

var signalServerRequestAbortedId int
var signalServerRequestAbortedMap = make(map[int]signalServerRequestAbortedDetail)
var signalServerRequestAbortedLock sync.RWMutex

// ServerSignalRequestAbortedCallback is a callback function for a 'request-aborted' signal emitted from a Server.
type ServerSignalRequestAbortedCallback func(targetObject *Server, message *Message, client *ClientContext)

/*
ConnectRequestAborted connects the callback to the 'request-aborted' signal for the Server.

The returned value represents the connection, and may be passed to DisconnectRequestAborted to remove it.
*/
func (recv *Server) ConnectRequestAborted(callback ServerSignalRequestAbortedCallback) int {
	signalServerRequestAbortedLock.Lock()
	defer signalServerRequestAbortedLock.Unlock()

	signalServerRequestAbortedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Server_signal_connect_request_aborted(instance, C.gpointer(uintptr(signalServerRequestAbortedId)))

	detail := signalServerRequestAbortedDetail{callback, handlerID}
	signalServerRequestAbortedMap[signalServerRequestAbortedId] = detail

	return signalServerRequestAbortedId
}

/*
DisconnectRequestAborted disconnects a callback from the 'request-aborted' signal for the Server.

The connectionID should be a value returned from a call to ConnectRequestAborted.
*/
func (recv *Server) DisconnectRequestAborted(connectionID int) {
	signalServerRequestAbortedLock.Lock()
	defer signalServerRequestAbortedLock.Unlock()

	detail, exists := signalServerRequestAbortedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalServerRequestAbortedMap, connectionID)
}

//export server_requestAbortedHandler
func server_requestAbortedHandler(c_targetObject *C.GObject, c_message *C.SoupMessage, c_client *C.SoupClientContext, data C.gpointer) {
	signalServerRequestAbortedLock.RLock()
	defer signalServerRequestAbortedLock.RUnlock()

	message := MessageNewFromC(unsafe.Pointer(c_message))

	client := ClientContextNewFromC(unsafe.Pointer(c_client))

	targetObject := ServerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalServerRequestAbortedMap[index].callback
	callback(targetObject, message, client)
}

type signalServerRequestFinishedDetail struct {
	callback  ServerSignalRequestFinishedCallback
	handlerID C.gulong
}

var signalServerRequestFinishedId int
var signalServerRequestFinishedMap = make(map[int]signalServerRequestFinishedDetail)
var signalServerRequestFinishedLock sync.RWMutex

// ServerSignalRequestFinishedCallback is a callback function for a 'request-finished' signal emitted from a Server.
type ServerSignalRequestFinishedCallback func(targetObject *Server, message *Message, client *ClientContext)

/*
ConnectRequestFinished connects the callback to the 'request-finished' signal for the Server.

The returned value represents the connection, and may be passed to DisconnectRequestFinished to remove it.
*/
func (recv *Server) ConnectRequestFinished(callback ServerSignalRequestFinishedCallback) int {
	signalServerRequestFinishedLock.Lock()
	defer signalServerRequestFinishedLock.Unlock()

	signalServerRequestFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Server_signal_connect_request_finished(instance, C.gpointer(uintptr(signalServerRequestFinishedId)))

	detail := signalServerRequestFinishedDetail{callback, handlerID}
	signalServerRequestFinishedMap[signalServerRequestFinishedId] = detail

	return signalServerRequestFinishedId
}

/*
DisconnectRequestFinished disconnects a callback from the 'request-finished' signal for the Server.

The connectionID should be a value returned from a call to ConnectRequestFinished.
*/
func (recv *Server) DisconnectRequestFinished(connectionID int) {
	signalServerRequestFinishedLock.Lock()
	defer signalServerRequestFinishedLock.Unlock()

	detail, exists := signalServerRequestFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalServerRequestFinishedMap, connectionID)
}

//export server_requestFinishedHandler
func server_requestFinishedHandler(c_targetObject *C.GObject, c_message *C.SoupMessage, c_client *C.SoupClientContext, data C.gpointer) {
	signalServerRequestFinishedLock.RLock()
	defer signalServerRequestFinishedLock.RUnlock()

	message := MessageNewFromC(unsafe.Pointer(c_message))

	client := ClientContextNewFromC(unsafe.Pointer(c_client))

	targetObject := ServerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalServerRequestFinishedMap[index].callback
	callback(targetObject, message, client)
}

type signalServerRequestReadDetail struct {
	callback  ServerSignalRequestReadCallback
	handlerID C.gulong
}

var signalServerRequestReadId int
var signalServerRequestReadMap = make(map[int]signalServerRequestReadDetail)
var signalServerRequestReadLock sync.RWMutex

// ServerSignalRequestReadCallback is a callback function for a 'request-read' signal emitted from a Server.
type ServerSignalRequestReadCallback func(targetObject *Server, message *Message, client *ClientContext)

/*
ConnectRequestRead connects the callback to the 'request-read' signal for the Server.

The returned value represents the connection, and may be passed to DisconnectRequestRead to remove it.
*/
func (recv *Server) ConnectRequestRead(callback ServerSignalRequestReadCallback) int {
	signalServerRequestReadLock.Lock()
	defer signalServerRequestReadLock.Unlock()

	signalServerRequestReadId++
	instance := C.gpointer(recv.native)
	handlerID := C.Server_signal_connect_request_read(instance, C.gpointer(uintptr(signalServerRequestReadId)))

	detail := signalServerRequestReadDetail{callback, handlerID}
	signalServerRequestReadMap[signalServerRequestReadId] = detail

	return signalServerRequestReadId
}

/*
DisconnectRequestRead disconnects a callback from the 'request-read' signal for the Server.

The connectionID should be a value returned from a call to ConnectRequestRead.
*/
func (recv *Server) DisconnectRequestRead(connectionID int) {
	signalServerRequestReadLock.Lock()
	defer signalServerRequestReadLock.Unlock()

	detail, exists := signalServerRequestReadMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalServerRequestReadMap, connectionID)
}

//export server_requestReadHandler
func server_requestReadHandler(c_targetObject *C.GObject, c_message *C.SoupMessage, c_client *C.SoupClientContext, data C.gpointer) {
	signalServerRequestReadLock.RLock()
	defer signalServerRequestReadLock.RUnlock()

	message := MessageNewFromC(unsafe.Pointer(c_message))

	client := ClientContextNewFromC(unsafe.Pointer(c_client))

	targetObject := ServerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalServerRequestReadMap[index].callback
	callback(targetObject, message, client)
}

type signalServerRequestStartedDetail struct {
	callback  ServerSignalRequestStartedCallback
	handlerID C.gulong
}

var signalServerRequestStartedId int
var signalServerRequestStartedMap = make(map[int]signalServerRequestStartedDetail)
var signalServerRequestStartedLock sync.RWMutex

// ServerSignalRequestStartedCallback is a callback function for a 'request-started' signal emitted from a Server.
type ServerSignalRequestStartedCallback func(targetObject *Server, message *Message, client *ClientContext)

/*
ConnectRequestStarted connects the callback to the 'request-started' signal for the Server.

The returned value represents the connection, and may be passed to DisconnectRequestStarted to remove it.
*/
func (recv *Server) ConnectRequestStarted(callback ServerSignalRequestStartedCallback) int {
	signalServerRequestStartedLock.Lock()
	defer signalServerRequestStartedLock.Unlock()

	signalServerRequestStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Server_signal_connect_request_started(instance, C.gpointer(uintptr(signalServerRequestStartedId)))

	detail := signalServerRequestStartedDetail{callback, handlerID}
	signalServerRequestStartedMap[signalServerRequestStartedId] = detail

	return signalServerRequestStartedId
}

/*
DisconnectRequestStarted disconnects a callback from the 'request-started' signal for the Server.

The connectionID should be a value returned from a call to ConnectRequestStarted.
*/
func (recv *Server) DisconnectRequestStarted(connectionID int) {
	signalServerRequestStartedLock.Lock()
	defer signalServerRequestStartedLock.Unlock()

	detail, exists := signalServerRequestStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalServerRequestStartedMap, connectionID)
}

//export server_requestStartedHandler
func server_requestStartedHandler(c_targetObject *C.GObject, c_message *C.SoupMessage, c_client *C.SoupClientContext, data C.gpointer) {
	signalServerRequestStartedLock.RLock()
	defer signalServerRequestStartedLock.RUnlock()

	message := MessageNewFromC(unsafe.Pointer(c_message))

	client := ClientContextNewFromC(unsafe.Pointer(c_client))

	targetObject := ServerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalServerRequestStartedMap[index].callback
	callback(targetObject, message, client)
}

// Unsupported : soup_server_new : unsupported parameter ... : varargs

// AcceptIostream is a wrapper around the C function soup_server_accept_iostream.
func (recv *Server) AcceptIostream(stream *gio.IOStream, localAddr *gio.SocketAddress, remoteAddr *gio.SocketAddress) (bool, error) {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_local_addr := (*C.GSocketAddress)(C.NULL)
	if localAddr != nil {
		c_local_addr = (*C.GSocketAddress)(localAddr.ToC())
	}

	c_remote_addr := (*C.GSocketAddress)(C.NULL)
	if remoteAddr != nil {
		c_remote_addr = (*C.GSocketAddress)(remoteAddr.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_server_accept_iostream((*C.SoupServer)(recv.native), c_stream, c_local_addr, c_remote_addr, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddAuthDomain is a wrapper around the C function soup_server_add_auth_domain.
func (recv *Server) AddAuthDomain(authDomain *AuthDomain) {
	c_auth_domain := (*C.SoupAuthDomain)(C.NULL)
	if authDomain != nil {
		c_auth_domain = (*C.SoupAuthDomain)(authDomain.ToC())
	}

	C.soup_server_add_auth_domain((*C.SoupServer)(recv.native), c_auth_domain)

	return
}

// Unsupported : soup_server_add_early_handler : unsupported parameter callback : no type generator for ServerCallback (SoupServerCallback) for param callback

// Unsupported : soup_server_add_handler : unsupported parameter callback : no type generator for ServerCallback (SoupServerCallback) for param callback

// Unsupported : soup_server_add_websocket_handler : unsupported parameter callback : no type generator for ServerWebsocketCallback (SoupServerWebsocketCallback) for param callback

// Disconnect is a wrapper around the C function soup_server_disconnect.
func (recv *Server) Disconnect() {
	C.soup_server_disconnect((*C.SoupServer)(recv.native))

	return
}

// GetAsyncContext is a wrapper around the C function soup_server_get_async_context.
func (recv *Server) GetAsyncContext() *glib.MainContext {
	retC := C.soup_server_get_async_context((*C.SoupServer)(recv.native))
	var retGo (*glib.MainContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.MainContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetListener is a wrapper around the C function soup_server_get_listener.
func (recv *Server) GetListener() *Socket {
	retC := C.soup_server_get_listener((*C.SoupServer)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetListeners is a wrapper around the C function soup_server_get_listeners.
func (recv *Server) GetListeners() *glib.SList {
	retC := C.soup_server_get_listeners((*C.SoupServer)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPort is a wrapper around the C function soup_server_get_port.
func (recv *Server) GetPort() uint32 {
	retC := C.soup_server_get_port((*C.SoupServer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetUris is a wrapper around the C function soup_server_get_uris.
func (recv *Server) GetUris() *glib.SList {
	retC := C.soup_server_get_uris((*C.SoupServer)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsHttps is a wrapper around the C function soup_server_is_https.
func (recv *Server) IsHttps() bool {
	retC := C.soup_server_is_https((*C.SoupServer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listen is a wrapper around the C function soup_server_listen.
func (recv *Server) Listen(address *gio.SocketAddress, options ServerListenOptions) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_options := (C.SoupServerListenOptions)(options)

	var cThrowableError *C.GError

	retC := C.soup_server_listen((*C.SoupServer)(recv.native), c_address, c_options, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ListenAll is a wrapper around the C function soup_server_listen_all.
func (recv *Server) ListenAll(port uint32, options ServerListenOptions) (bool, error) {
	c_port := (C.guint)(port)

	c_options := (C.SoupServerListenOptions)(options)

	var cThrowableError *C.GError

	retC := C.soup_server_listen_all((*C.SoupServer)(recv.native), c_port, c_options, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ListenFd is a wrapper around the C function soup_server_listen_fd.
func (recv *Server) ListenFd(fd int32, options ServerListenOptions) (bool, error) {
	c_fd := (C.int)(fd)

	c_options := (C.SoupServerListenOptions)(options)

	var cThrowableError *C.GError

	retC := C.soup_server_listen_fd((*C.SoupServer)(recv.native), c_fd, c_options, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ListenLocal is a wrapper around the C function soup_server_listen_local.
func (recv *Server) ListenLocal(port uint32, options ServerListenOptions) (bool, error) {
	c_port := (C.guint)(port)

	c_options := (C.SoupServerListenOptions)(options)

	var cThrowableError *C.GError

	retC := C.soup_server_listen_local((*C.SoupServer)(recv.native), c_port, c_options, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ListenSocket is a wrapper around the C function soup_server_listen_socket.
func (recv *Server) ListenSocket(socket *gio.Socket, options ServerListenOptions) (bool, error) {
	c_socket := (*C.GSocket)(C.NULL)
	if socket != nil {
		c_socket = (*C.GSocket)(socket.ToC())
	}

	c_options := (C.SoupServerListenOptions)(options)

	var cThrowableError *C.GError

	retC := C.soup_server_listen_socket((*C.SoupServer)(recv.native), c_socket, c_options, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PauseMessage is a wrapper around the C function soup_server_pause_message.
func (recv *Server) PauseMessage(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_server_pause_message((*C.SoupServer)(recv.native), c_msg)

	return
}

// Quit is a wrapper around the C function soup_server_quit.
func (recv *Server) Quit() {
	C.soup_server_quit((*C.SoupServer)(recv.native))

	return
}

// RemoveAuthDomain is a wrapper around the C function soup_server_remove_auth_domain.
func (recv *Server) RemoveAuthDomain(authDomain *AuthDomain) {
	c_auth_domain := (*C.SoupAuthDomain)(C.NULL)
	if authDomain != nil {
		c_auth_domain = (*C.SoupAuthDomain)(authDomain.ToC())
	}

	C.soup_server_remove_auth_domain((*C.SoupServer)(recv.native), c_auth_domain)

	return
}

// RemoveHandler is a wrapper around the C function soup_server_remove_handler.
func (recv *Server) RemoveHandler(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.soup_server_remove_handler((*C.SoupServer)(recv.native), c_path)

	return
}

// Run is a wrapper around the C function soup_server_run.
func (recv *Server) Run() {
	C.soup_server_run((*C.SoupServer)(recv.native))

	return
}

// RunAsync is a wrapper around the C function soup_server_run_async.
func (recv *Server) RunAsync() {
	C.soup_server_run_async((*C.SoupServer)(recv.native))

	return
}

// SetSslCertFile is a wrapper around the C function soup_server_set_ssl_cert_file.
func (recv *Server) SetSslCertFile(sslCertFile string, sslKeyFile string) (bool, error) {
	c_ssl_cert_file := C.CString(sslCertFile)
	defer C.free(unsafe.Pointer(c_ssl_cert_file))

	c_ssl_key_file := C.CString(sslKeyFile)
	defer C.free(unsafe.Pointer(c_ssl_key_file))

	var cThrowableError *C.GError

	retC := C.soup_server_set_ssl_cert_file((*C.SoupServer)(recv.native), c_ssl_cert_file, c_ssl_key_file, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnpauseMessage is a wrapper around the C function soup_server_unpause_message.
func (recv *Server) UnpauseMessage(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_server_unpause_message((*C.SoupServer)(recv.native), c_msg)

	return
}

// Session is a wrapper around the C record SoupSession.
type Session struct {
	native *C.SoupSession
	// parent : record
}

func SessionNewFromC(u unsafe.Pointer) *Session {
	c := (*C.SoupSession)(u)
	if c == nil {
		return nil
	}

	g := &Session{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Session) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Session) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Session with another Session, and returns true if they represent the same GObject.
func (recv *Session) Equals(other *Session) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Session) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Session.
// Exercise care, as this is a potentially dangerous function if the Object is not a Session.
func CastToSession(object *gobject.Object) *Session {
	return SessionNewFromC(object.ToC())
}

type signalSessionAuthenticateDetail struct {
	callback  SessionSignalAuthenticateCallback
	handlerID C.gulong
}

var signalSessionAuthenticateId int
var signalSessionAuthenticateMap = make(map[int]signalSessionAuthenticateDetail)
var signalSessionAuthenticateLock sync.RWMutex

// SessionSignalAuthenticateCallback is a callback function for a 'authenticate' signal emitted from a Session.
type SessionSignalAuthenticateCallback func(targetObject *Session, msg *Message, auth *Auth, retrying bool)

/*
ConnectAuthenticate connects the callback to the 'authenticate' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectAuthenticate to remove it.
*/
func (recv *Session) ConnectAuthenticate(callback SessionSignalAuthenticateCallback) int {
	signalSessionAuthenticateLock.Lock()
	defer signalSessionAuthenticateLock.Unlock()

	signalSessionAuthenticateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_authenticate(instance, C.gpointer(uintptr(signalSessionAuthenticateId)))

	detail := signalSessionAuthenticateDetail{callback, handlerID}
	signalSessionAuthenticateMap[signalSessionAuthenticateId] = detail

	return signalSessionAuthenticateId
}

/*
DisconnectAuthenticate disconnects a callback from the 'authenticate' signal for the Session.

The connectionID should be a value returned from a call to ConnectAuthenticate.
*/
func (recv *Session) DisconnectAuthenticate(connectionID int) {
	signalSessionAuthenticateLock.Lock()
	defer signalSessionAuthenticateLock.Unlock()

	detail, exists := signalSessionAuthenticateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionAuthenticateMap, connectionID)
}

//export session_authenticateHandler
func session_authenticateHandler(c_targetObject *C.GObject, c_msg *C.SoupMessage, c_auth *C.SoupAuth, c_retrying C.gboolean, data C.gpointer) {
	signalSessionAuthenticateLock.RLock()
	defer signalSessionAuthenticateLock.RUnlock()

	msg := MessageNewFromC(unsafe.Pointer(c_msg))

	auth := AuthNewFromC(unsafe.Pointer(c_auth))

	retrying := c_retrying == C.TRUE

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionAuthenticateMap[index].callback
	callback(targetObject, msg, auth, retrying)
}

type signalSessionConnectionCreatedDetail struct {
	callback  SessionSignalConnectionCreatedCallback
	handlerID C.gulong
}

var signalSessionConnectionCreatedId int
var signalSessionConnectionCreatedMap = make(map[int]signalSessionConnectionCreatedDetail)
var signalSessionConnectionCreatedLock sync.RWMutex

// SessionSignalConnectionCreatedCallback is a callback function for a 'connection-created' signal emitted from a Session.
type SessionSignalConnectionCreatedCallback func(targetObject *Session, connection *gobject.Object)

/*
ConnectConnectionCreated connects the callback to the 'connection-created' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectConnectionCreated to remove it.
*/
func (recv *Session) ConnectConnectionCreated(callback SessionSignalConnectionCreatedCallback) int {
	signalSessionConnectionCreatedLock.Lock()
	defer signalSessionConnectionCreatedLock.Unlock()

	signalSessionConnectionCreatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_connection_created(instance, C.gpointer(uintptr(signalSessionConnectionCreatedId)))

	detail := signalSessionConnectionCreatedDetail{callback, handlerID}
	signalSessionConnectionCreatedMap[signalSessionConnectionCreatedId] = detail

	return signalSessionConnectionCreatedId
}

/*
DisconnectConnectionCreated disconnects a callback from the 'connection-created' signal for the Session.

The connectionID should be a value returned from a call to ConnectConnectionCreated.
*/
func (recv *Session) DisconnectConnectionCreated(connectionID int) {
	signalSessionConnectionCreatedLock.Lock()
	defer signalSessionConnectionCreatedLock.Unlock()

	detail, exists := signalSessionConnectionCreatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionConnectionCreatedMap, connectionID)
}

//export session_connectionCreatedHandler
func session_connectionCreatedHandler(c_targetObject *C.GObject, c_connection *C.GObject, data C.gpointer) {
	signalSessionConnectionCreatedLock.RLock()
	defer signalSessionConnectionCreatedLock.RUnlock()

	connection := gobject.ObjectNewFromC(unsafe.Pointer(c_connection))

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionConnectionCreatedMap[index].callback
	callback(targetObject, connection)
}

type signalSessionRequestQueuedDetail struct {
	callback  SessionSignalRequestQueuedCallback
	handlerID C.gulong
}

var signalSessionRequestQueuedId int
var signalSessionRequestQueuedMap = make(map[int]signalSessionRequestQueuedDetail)
var signalSessionRequestQueuedLock sync.RWMutex

// SessionSignalRequestQueuedCallback is a callback function for a 'request-queued' signal emitted from a Session.
type SessionSignalRequestQueuedCallback func(targetObject *Session, msg *Message)

/*
ConnectRequestQueued connects the callback to the 'request-queued' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectRequestQueued to remove it.
*/
func (recv *Session) ConnectRequestQueued(callback SessionSignalRequestQueuedCallback) int {
	signalSessionRequestQueuedLock.Lock()
	defer signalSessionRequestQueuedLock.Unlock()

	signalSessionRequestQueuedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_request_queued(instance, C.gpointer(uintptr(signalSessionRequestQueuedId)))

	detail := signalSessionRequestQueuedDetail{callback, handlerID}
	signalSessionRequestQueuedMap[signalSessionRequestQueuedId] = detail

	return signalSessionRequestQueuedId
}

/*
DisconnectRequestQueued disconnects a callback from the 'request-queued' signal for the Session.

The connectionID should be a value returned from a call to ConnectRequestQueued.
*/
func (recv *Session) DisconnectRequestQueued(connectionID int) {
	signalSessionRequestQueuedLock.Lock()
	defer signalSessionRequestQueuedLock.Unlock()

	detail, exists := signalSessionRequestQueuedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionRequestQueuedMap, connectionID)
}

//export session_requestQueuedHandler
func session_requestQueuedHandler(c_targetObject *C.GObject, c_msg *C.SoupMessage, data C.gpointer) {
	signalSessionRequestQueuedLock.RLock()
	defer signalSessionRequestQueuedLock.RUnlock()

	msg := MessageNewFromC(unsafe.Pointer(c_msg))

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionRequestQueuedMap[index].callback
	callback(targetObject, msg)
}

type signalSessionRequestStartedDetail struct {
	callback  SessionSignalRequestStartedCallback
	handlerID C.gulong
}

var signalSessionRequestStartedId int
var signalSessionRequestStartedMap = make(map[int]signalSessionRequestStartedDetail)
var signalSessionRequestStartedLock sync.RWMutex

// SessionSignalRequestStartedCallback is a callback function for a 'request-started' signal emitted from a Session.
type SessionSignalRequestStartedCallback func(targetObject *Session, msg *Message, socket *Socket)

/*
ConnectRequestStarted connects the callback to the 'request-started' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectRequestStarted to remove it.
*/
func (recv *Session) ConnectRequestStarted(callback SessionSignalRequestStartedCallback) int {
	signalSessionRequestStartedLock.Lock()
	defer signalSessionRequestStartedLock.Unlock()

	signalSessionRequestStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_request_started(instance, C.gpointer(uintptr(signalSessionRequestStartedId)))

	detail := signalSessionRequestStartedDetail{callback, handlerID}
	signalSessionRequestStartedMap[signalSessionRequestStartedId] = detail

	return signalSessionRequestStartedId
}

/*
DisconnectRequestStarted disconnects a callback from the 'request-started' signal for the Session.

The connectionID should be a value returned from a call to ConnectRequestStarted.
*/
func (recv *Session) DisconnectRequestStarted(connectionID int) {
	signalSessionRequestStartedLock.Lock()
	defer signalSessionRequestStartedLock.Unlock()

	detail, exists := signalSessionRequestStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionRequestStartedMap, connectionID)
}

//export session_requestStartedHandler
func session_requestStartedHandler(c_targetObject *C.GObject, c_msg *C.SoupMessage, c_socket *C.SoupSocket, data C.gpointer) {
	signalSessionRequestStartedLock.RLock()
	defer signalSessionRequestStartedLock.RUnlock()

	msg := MessageNewFromC(unsafe.Pointer(c_msg))

	socket := SocketNewFromC(unsafe.Pointer(c_socket))

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionRequestStartedMap[index].callback
	callback(targetObject, msg, socket)
}

type signalSessionRequestUnqueuedDetail struct {
	callback  SessionSignalRequestUnqueuedCallback
	handlerID C.gulong
}

var signalSessionRequestUnqueuedId int
var signalSessionRequestUnqueuedMap = make(map[int]signalSessionRequestUnqueuedDetail)
var signalSessionRequestUnqueuedLock sync.RWMutex

// SessionSignalRequestUnqueuedCallback is a callback function for a 'request-unqueued' signal emitted from a Session.
type SessionSignalRequestUnqueuedCallback func(targetObject *Session, msg *Message)

/*
ConnectRequestUnqueued connects the callback to the 'request-unqueued' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectRequestUnqueued to remove it.
*/
func (recv *Session) ConnectRequestUnqueued(callback SessionSignalRequestUnqueuedCallback) int {
	signalSessionRequestUnqueuedLock.Lock()
	defer signalSessionRequestUnqueuedLock.Unlock()

	signalSessionRequestUnqueuedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_request_unqueued(instance, C.gpointer(uintptr(signalSessionRequestUnqueuedId)))

	detail := signalSessionRequestUnqueuedDetail{callback, handlerID}
	signalSessionRequestUnqueuedMap[signalSessionRequestUnqueuedId] = detail

	return signalSessionRequestUnqueuedId
}

/*
DisconnectRequestUnqueued disconnects a callback from the 'request-unqueued' signal for the Session.

The connectionID should be a value returned from a call to ConnectRequestUnqueued.
*/
func (recv *Session) DisconnectRequestUnqueued(connectionID int) {
	signalSessionRequestUnqueuedLock.Lock()
	defer signalSessionRequestUnqueuedLock.Unlock()

	detail, exists := signalSessionRequestUnqueuedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionRequestUnqueuedMap, connectionID)
}

//export session_requestUnqueuedHandler
func session_requestUnqueuedHandler(c_targetObject *C.GObject, c_msg *C.SoupMessage, data C.gpointer) {
	signalSessionRequestUnqueuedLock.RLock()
	defer signalSessionRequestUnqueuedLock.RUnlock()

	msg := MessageNewFromC(unsafe.Pointer(c_msg))

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionRequestUnqueuedMap[index].callback
	callback(targetObject, msg)
}

type signalSessionTunnelingDetail struct {
	callback  SessionSignalTunnelingCallback
	handlerID C.gulong
}

var signalSessionTunnelingId int
var signalSessionTunnelingMap = make(map[int]signalSessionTunnelingDetail)
var signalSessionTunnelingLock sync.RWMutex

// SessionSignalTunnelingCallback is a callback function for a 'tunneling' signal emitted from a Session.
type SessionSignalTunnelingCallback func(targetObject *Session, connection *gobject.Object)

/*
ConnectTunneling connects the callback to the 'tunneling' signal for the Session.

The returned value represents the connection, and may be passed to DisconnectTunneling to remove it.
*/
func (recv *Session) ConnectTunneling(callback SessionSignalTunnelingCallback) int {
	signalSessionTunnelingLock.Lock()
	defer signalSessionTunnelingLock.Unlock()

	signalSessionTunnelingId++
	instance := C.gpointer(recv.native)
	handlerID := C.Session_signal_connect_tunneling(instance, C.gpointer(uintptr(signalSessionTunnelingId)))

	detail := signalSessionTunnelingDetail{callback, handlerID}
	signalSessionTunnelingMap[signalSessionTunnelingId] = detail

	return signalSessionTunnelingId
}

/*
DisconnectTunneling disconnects a callback from the 'tunneling' signal for the Session.

The connectionID should be a value returned from a call to ConnectTunneling.
*/
func (recv *Session) DisconnectTunneling(connectionID int) {
	signalSessionTunnelingLock.Lock()
	defer signalSessionTunnelingLock.Unlock()

	detail, exists := signalSessionTunnelingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSessionTunnelingMap, connectionID)
}

//export session_tunnelingHandler
func session_tunnelingHandler(c_targetObject *C.GObject, c_connection *C.GObject, data C.gpointer) {
	signalSessionTunnelingLock.RLock()
	defer signalSessionTunnelingLock.RUnlock()

	connection := gobject.ObjectNewFromC(unsafe.Pointer(c_connection))

	targetObject := SessionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSessionTunnelingMap[index].callback
	callback(targetObject, connection)
}

// SessionNew is a wrapper around the C function soup_session_new.
func SessionNew() *Session {
	retC := C.soup_session_new()
	retGo := SessionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : soup_session_new_with_options : unsupported parameter ... : varargs

// Abort is a wrapper around the C function soup_session_abort.
func (recv *Session) Abort() {
	C.soup_session_abort((*C.SoupSession)(recv.native))

	return
}

// AddFeature is a wrapper around the C function soup_session_add_feature.
func (recv *Session) AddFeature(feature *SessionFeature) {
	c_feature := (*C.SoupSessionFeature)(feature.ToC())

	C.soup_session_add_feature((*C.SoupSession)(recv.native), c_feature)

	return
}

// AddFeatureByType is a wrapper around the C function soup_session_add_feature_by_type.
func (recv *Session) AddFeatureByType(featureType gobject.Type) {
	c_feature_type := (C.GType)(featureType)

	C.soup_session_add_feature_by_type((*C.SoupSession)(recv.native), c_feature_type)

	return
}

// CancelMessage is a wrapper around the C function soup_session_cancel_message.
func (recv *Session) CancelMessage(msg *Message, statusCode uint32) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_status_code := (C.guint)(statusCode)

	C.soup_session_cancel_message((*C.SoupSession)(recv.native), c_msg, c_status_code)

	return
}

// GetAsyncContext is a wrapper around the C function soup_session_get_async_context.
func (recv *Session) GetAsyncContext() *glib.MainContext {
	retC := C.soup_session_get_async_context((*C.SoupSession)(recv.native))
	var retGo (*glib.MainContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.MainContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFeature is a wrapper around the C function soup_session_get_feature.
func (recv *Session) GetFeature(featureType gobject.Type) *SessionFeature {
	c_feature_type := (C.GType)(featureType)

	retC := C.soup_session_get_feature((*C.SoupSession)(recv.native), c_feature_type)
	var retGo (*SessionFeature)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SessionFeatureNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFeatureForMessage is a wrapper around the C function soup_session_get_feature_for_message.
func (recv *Session) GetFeatureForMessage(featureType gobject.Type, msg *Message) *SessionFeature {
	c_feature_type := (C.GType)(featureType)

	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_session_get_feature_for_message((*C.SoupSession)(recv.native), c_feature_type, c_msg)
	var retGo (*SessionFeature)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SessionFeatureNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFeatures is a wrapper around the C function soup_session_get_features.
func (recv *Session) GetFeatures(featureType gobject.Type) *glib.SList {
	c_feature_type := (C.GType)(featureType)

	retC := C.soup_session_get_features((*C.SoupSession)(recv.native), c_feature_type)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasFeature is a wrapper around the C function soup_session_has_feature.
func (recv *Session) HasFeature(featureType gobject.Type) bool {
	c_feature_type := (C.GType)(featureType)

	retC := C.soup_session_has_feature((*C.SoupSession)(recv.native), c_feature_type)
	retGo := retC == C.TRUE

	return retGo
}

// PauseMessage is a wrapper around the C function soup_session_pause_message.
func (recv *Session) PauseMessage(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_session_pause_message((*C.SoupSession)(recv.native), c_msg)

	return
}

// Unsupported : soup_session_prefetch_dns : unsupported parameter callback : no type generator for AddressCallback (SoupAddressCallback) for param callback

// PrepareForUri is a wrapper around the C function soup_session_prepare_for_uri.
func (recv *Session) PrepareForUri(uri *URI) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	C.soup_session_prepare_for_uri((*C.SoupSession)(recv.native), c_uri)

	return
}

// Unsupported : soup_session_queue_message : unsupported parameter callback : no type generator for SessionCallback (SoupSessionCallback) for param callback

// RedirectMessage is a wrapper around the C function soup_session_redirect_message.
func (recv *Session) RedirectMessage(msg *Message) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_session_redirect_message((*C.SoupSession)(recv.native), c_msg)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveFeature is a wrapper around the C function soup_session_remove_feature.
func (recv *Session) RemoveFeature(feature *SessionFeature) {
	c_feature := (*C.SoupSessionFeature)(feature.ToC())

	C.soup_session_remove_feature((*C.SoupSession)(recv.native), c_feature)

	return
}

// RemoveFeatureByType is a wrapper around the C function soup_session_remove_feature_by_type.
func (recv *Session) RemoveFeatureByType(featureType gobject.Type) {
	c_feature_type := (C.GType)(featureType)

	C.soup_session_remove_feature_by_type((*C.SoupSession)(recv.native), c_feature_type)

	return
}

// Request is a wrapper around the C function soup_session_request.
func (recv *Session) Request(uriString string) (*Request, error) {
	c_uri_string := C.CString(uriString)
	defer C.free(unsafe.Pointer(c_uri_string))

	var cThrowableError *C.GError

	retC := C.soup_session_request((*C.SoupSession)(recv.native), c_uri_string, &cThrowableError)
	retGo := RequestNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequestHttp is a wrapper around the C function soup_session_request_http.
func (recv *Session) RequestHttp(method string, uriString string) (*RequestHTTP, error) {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri_string := C.CString(uriString)
	defer C.free(unsafe.Pointer(c_uri_string))

	var cThrowableError *C.GError

	retC := C.soup_session_request_http((*C.SoupSession)(recv.native), c_method, c_uri_string, &cThrowableError)
	retGo := RequestHTTPNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequestHttpUri is a wrapper around the C function soup_session_request_http_uri.
func (recv *Session) RequestHttpUri(method string, uri *URI) (*RequestHTTP, error) {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_session_request_http_uri((*C.SoupSession)(recv.native), c_method, c_uri, &cThrowableError)
	retGo := RequestHTTPNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequestUri is a wrapper around the C function soup_session_request_uri.
func (recv *Session) RequestUri(uri *URI) (*Request, error) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_session_request_uri((*C.SoupSession)(recv.native), c_uri, &cThrowableError)
	retGo := RequestNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequeueMessage is a wrapper around the C function soup_session_requeue_message.
func (recv *Session) RequeueMessage(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_session_requeue_message((*C.SoupSession)(recv.native), c_msg)

	return
}

// Send is a wrapper around the C function soup_session_send.
func (recv *Session) Send(msg *Message, cancellable *gio.Cancellable) (*gio.InputStream, error) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_session_send((*C.SoupSession)(recv.native), c_msg, c_cancellable, &cThrowableError)
	retGo := gio.InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : soup_session_send_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SendFinish is a wrapper around the C function soup_session_send_finish.
func (recv *Session) SendFinish(result *gio.AsyncResult) (*gio.InputStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.soup_session_send_finish((*C.SoupSession)(recv.native), c_result, &cThrowableError)
	retGo := gio.InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SendMessage is a wrapper around the C function soup_session_send_message.
func (recv *Session) SendMessage(msg *Message) uint32 {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_session_send_message((*C.SoupSession)(recv.native), c_msg)
	retGo := (uint32)(retC)

	return retGo
}

// StealConnection is a wrapper around the C function soup_session_steal_connection.
func (recv *Session) StealConnection(msg *Message) *gio.IOStream {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_session_steal_connection((*C.SoupSession)(recv.native), c_msg)
	retGo := gio.IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnpauseMessage is a wrapper around the C function soup_session_unpause_message.
func (recv *Session) UnpauseMessage(msg *Message) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_session_unpause_message((*C.SoupSession)(recv.native), c_msg)

	return
}

// Unsupported : soup_session_websocket_connect_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WebsocketConnectFinish is a wrapper around the C function soup_session_websocket_connect_finish.
func (recv *Session) WebsocketConnectFinish(result *gio.AsyncResult) (*WebsocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.soup_session_websocket_connect_finish((*C.SoupSession)(recv.native), c_result, &cThrowableError)
	retGo := WebsocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WouldRedirect is a wrapper around the C function soup_session_would_redirect.
func (recv *Session) WouldRedirect(msg *Message) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_session_would_redirect((*C.SoupSession)(recv.native), c_msg)
	retGo := retC == C.TRUE

	return retGo
}

// SessionAsync is a wrapper around the C record SoupSessionAsync.
type SessionAsync struct {
	native *C.SoupSessionAsync
	// parent : record
}

func SessionAsyncNewFromC(u unsafe.Pointer) *SessionAsync {
	c := (*C.SoupSessionAsync)(u)
	if c == nil {
		return nil
	}

	g := &SessionAsync{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SessionAsync) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SessionAsync) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionAsync with another SessionAsync, and returns true if they represent the same GObject.
func (recv *SessionAsync) Equals(other *SessionAsync) bool {
	return other.ToC() == recv.ToC()
}

// Session upcasts to *Session
func (recv *SessionAsync) Session() *Session {
	return SessionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SessionAsync) Object() *gobject.Object {
	return recv.Session().Object()
}

// CastToWidget down casts any arbitrary Object to SessionAsync.
// Exercise care, as this is a potentially dangerous function if the Object is not a SessionAsync.
func CastToSessionAsync(object *gobject.Object) *SessionAsync {
	return SessionAsyncNewFromC(object.ToC())
}

// SessionAsyncNew is a wrapper around the C function soup_session_async_new.
func SessionAsyncNew() *SessionAsync {
	retC := C.soup_session_async_new()
	retGo := SessionAsyncNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : soup_session_async_new_with_options : unsupported parameter ... : varargs

// SessionSync is a wrapper around the C record SoupSessionSync.
type SessionSync struct {
	native *C.SoupSessionSync
	// parent : record
}

func SessionSyncNewFromC(u unsafe.Pointer) *SessionSync {
	c := (*C.SoupSessionSync)(u)
	if c == nil {
		return nil
	}

	g := &SessionSync{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SessionSync) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SessionSync) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionSync with another SessionSync, and returns true if they represent the same GObject.
func (recv *SessionSync) Equals(other *SessionSync) bool {
	return other.ToC() == recv.ToC()
}

// Session upcasts to *Session
func (recv *SessionSync) Session() *Session {
	return SessionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SessionSync) Object() *gobject.Object {
	return recv.Session().Object()
}

// CastToWidget down casts any arbitrary Object to SessionSync.
// Exercise care, as this is a potentially dangerous function if the Object is not a SessionSync.
func CastToSessionSync(object *gobject.Object) *SessionSync {
	return SessionSyncNewFromC(object.ToC())
}

// SessionSyncNew is a wrapper around the C function soup_session_sync_new.
func SessionSyncNew() *SessionSync {
	retC := C.soup_session_sync_new()
	retGo := SessionSyncNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : soup_session_sync_new_with_options : unsupported parameter ... : varargs

// Socket is a wrapper around the C record SoupSocket.
type Socket struct {
	native *C.SoupSocket
	// parent : record
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.SoupSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Socket) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Socket with another Socket, and returns true if they represent the same GObject.
func (recv *Socket) Equals(other *Socket) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Socket.
// Exercise care, as this is a potentially dangerous function if the Object is not a Socket.
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromC(object.ToC())
}

type signalSocketDisconnectedDetail struct {
	callback  SocketSignalDisconnectedCallback
	handlerID C.gulong
}

var signalSocketDisconnectedId int
var signalSocketDisconnectedMap = make(map[int]signalSocketDisconnectedDetail)
var signalSocketDisconnectedLock sync.RWMutex

// SocketSignalDisconnectedCallback is a callback function for a 'disconnected' signal emitted from a Socket.
type SocketSignalDisconnectedCallback func(targetObject *Socket)

/*
ConnectDisconnected connects the callback to the 'disconnected' signal for the Socket.

The returned value represents the connection, and may be passed to DisconnectDisconnected to remove it.
*/
func (recv *Socket) ConnectDisconnected(callback SocketSignalDisconnectedCallback) int {
	signalSocketDisconnectedLock.Lock()
	defer signalSocketDisconnectedLock.Unlock()

	signalSocketDisconnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Socket_signal_connect_disconnected(instance, C.gpointer(uintptr(signalSocketDisconnectedId)))

	detail := signalSocketDisconnectedDetail{callback, handlerID}
	signalSocketDisconnectedMap[signalSocketDisconnectedId] = detail

	return signalSocketDisconnectedId
}

/*
DisconnectDisconnected disconnects a callback from the 'disconnected' signal for the Socket.

The connectionID should be a value returned from a call to ConnectDisconnected.
*/
func (recv *Socket) DisconnectDisconnected(connectionID int) {
	signalSocketDisconnectedLock.Lock()
	defer signalSocketDisconnectedLock.Unlock()

	detail, exists := signalSocketDisconnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketDisconnectedMap, connectionID)
}

//export socket_disconnectedHandler
func socket_disconnectedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalSocketDisconnectedLock.RLock()
	defer signalSocketDisconnectedLock.RUnlock()

	targetObject := SocketNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSocketDisconnectedMap[index].callback
	callback(targetObject)
}

type signalSocketEventDetail struct {
	callback  SocketSignalEventCallback
	handlerID C.gulong
}

var signalSocketEventId int
var signalSocketEventMap = make(map[int]signalSocketEventDetail)
var signalSocketEventLock sync.RWMutex

// SocketSignalEventCallback is a callback function for a 'event' signal emitted from a Socket.
type SocketSignalEventCallback func(targetObject *Socket, event gio.SocketClientEvent, connection *gio.IOStream)

/*
ConnectEvent connects the callback to the 'event' signal for the Socket.

The returned value represents the connection, and may be passed to DisconnectEvent to remove it.
*/
func (recv *Socket) ConnectEvent(callback SocketSignalEventCallback) int {
	signalSocketEventLock.Lock()
	defer signalSocketEventLock.Unlock()

	signalSocketEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Socket_signal_connect_event(instance, C.gpointer(uintptr(signalSocketEventId)))

	detail := signalSocketEventDetail{callback, handlerID}
	signalSocketEventMap[signalSocketEventId] = detail

	return signalSocketEventId
}

/*
DisconnectEvent disconnects a callback from the 'event' signal for the Socket.

The connectionID should be a value returned from a call to ConnectEvent.
*/
func (recv *Socket) DisconnectEvent(connectionID int) {
	signalSocketEventLock.Lock()
	defer signalSocketEventLock.Unlock()

	detail, exists := signalSocketEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketEventMap, connectionID)
}

//export socket_eventHandler
func socket_eventHandler(c_targetObject *C.GObject, c_event C.GSocketClientEvent, c_connection *C.GIOStream, data C.gpointer) {
	signalSocketEventLock.RLock()
	defer signalSocketEventLock.RUnlock()

	event := gio.SocketClientEvent(c_event)

	connection := gio.IOStreamNewFromC(unsafe.Pointer(c_connection))

	targetObject := SocketNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSocketEventMap[index].callback
	callback(targetObject, event, connection)
}

type signalSocketNewConnectionDetail struct {
	callback  SocketSignalNewConnectionCallback
	handlerID C.gulong
}

var signalSocketNewConnectionId int
var signalSocketNewConnectionMap = make(map[int]signalSocketNewConnectionDetail)
var signalSocketNewConnectionLock sync.RWMutex

// SocketSignalNewConnectionCallback is a callback function for a 'new-connection' signal emitted from a Socket.
type SocketSignalNewConnectionCallback func(targetObject *Socket, new *Socket)

/*
ConnectNewConnection connects the callback to the 'new-connection' signal for the Socket.

The returned value represents the connection, and may be passed to DisconnectNewConnection to remove it.
*/
func (recv *Socket) ConnectNewConnection(callback SocketSignalNewConnectionCallback) int {
	signalSocketNewConnectionLock.Lock()
	defer signalSocketNewConnectionLock.Unlock()

	signalSocketNewConnectionId++
	instance := C.gpointer(recv.native)
	handlerID := C.Socket_signal_connect_new_connection(instance, C.gpointer(uintptr(signalSocketNewConnectionId)))

	detail := signalSocketNewConnectionDetail{callback, handlerID}
	signalSocketNewConnectionMap[signalSocketNewConnectionId] = detail

	return signalSocketNewConnectionId
}

/*
DisconnectNewConnection disconnects a callback from the 'new-connection' signal for the Socket.

The connectionID should be a value returned from a call to ConnectNewConnection.
*/
func (recv *Socket) DisconnectNewConnection(connectionID int) {
	signalSocketNewConnectionLock.Lock()
	defer signalSocketNewConnectionLock.Unlock()

	detail, exists := signalSocketNewConnectionMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketNewConnectionMap, connectionID)
}

//export socket_newConnectionHandler
func socket_newConnectionHandler(c_targetObject *C.GObject, c_new *C.SoupSocket, data C.gpointer) {
	signalSocketNewConnectionLock.RLock()
	defer signalSocketNewConnectionLock.RUnlock()

	new := SocketNewFromC(unsafe.Pointer(c_new))

	targetObject := SocketNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSocketNewConnectionMap[index].callback
	callback(targetObject, new)
}

type signalSocketReadableDetail struct {
	callback  SocketSignalReadableCallback
	handlerID C.gulong
}

var signalSocketReadableId int
var signalSocketReadableMap = make(map[int]signalSocketReadableDetail)
var signalSocketReadableLock sync.RWMutex

// SocketSignalReadableCallback is a callback function for a 'readable' signal emitted from a Socket.
type SocketSignalReadableCallback func(targetObject *Socket)

/*
ConnectReadable connects the callback to the 'readable' signal for the Socket.

The returned value represents the connection, and may be passed to DisconnectReadable to remove it.
*/
func (recv *Socket) ConnectReadable(callback SocketSignalReadableCallback) int {
	signalSocketReadableLock.Lock()
	defer signalSocketReadableLock.Unlock()

	signalSocketReadableId++
	instance := C.gpointer(recv.native)
	handlerID := C.Socket_signal_connect_readable(instance, C.gpointer(uintptr(signalSocketReadableId)))

	detail := signalSocketReadableDetail{callback, handlerID}
	signalSocketReadableMap[signalSocketReadableId] = detail

	return signalSocketReadableId
}

/*
DisconnectReadable disconnects a callback from the 'readable' signal for the Socket.

The connectionID should be a value returned from a call to ConnectReadable.
*/
func (recv *Socket) DisconnectReadable(connectionID int) {
	signalSocketReadableLock.Lock()
	defer signalSocketReadableLock.Unlock()

	detail, exists := signalSocketReadableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketReadableMap, connectionID)
}

//export socket_readableHandler
func socket_readableHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalSocketReadableLock.RLock()
	defer signalSocketReadableLock.RUnlock()

	targetObject := SocketNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSocketReadableMap[index].callback
	callback(targetObject)
}

type signalSocketWritableDetail struct {
	callback  SocketSignalWritableCallback
	handlerID C.gulong
}

var signalSocketWritableId int
var signalSocketWritableMap = make(map[int]signalSocketWritableDetail)
var signalSocketWritableLock sync.RWMutex

// SocketSignalWritableCallback is a callback function for a 'writable' signal emitted from a Socket.
type SocketSignalWritableCallback func(targetObject *Socket)

/*
ConnectWritable connects the callback to the 'writable' signal for the Socket.

The returned value represents the connection, and may be passed to DisconnectWritable to remove it.
*/
func (recv *Socket) ConnectWritable(callback SocketSignalWritableCallback) int {
	signalSocketWritableLock.Lock()
	defer signalSocketWritableLock.Unlock()

	signalSocketWritableId++
	instance := C.gpointer(recv.native)
	handlerID := C.Socket_signal_connect_writable(instance, C.gpointer(uintptr(signalSocketWritableId)))

	detail := signalSocketWritableDetail{callback, handlerID}
	signalSocketWritableMap[signalSocketWritableId] = detail

	return signalSocketWritableId
}

/*
DisconnectWritable disconnects a callback from the 'writable' signal for the Socket.

The connectionID should be a value returned from a call to ConnectWritable.
*/
func (recv *Socket) DisconnectWritable(connectionID int) {
	signalSocketWritableLock.Lock()
	defer signalSocketWritableLock.Unlock()

	detail, exists := signalSocketWritableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketWritableMap, connectionID)
}

//export socket_writableHandler
func socket_writableHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalSocketWritableLock.RLock()
	defer signalSocketWritableLock.RUnlock()

	targetObject := SocketNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalSocketWritableMap[index].callback
	callback(targetObject)
}

// Unsupported : soup_socket_new : unsupported parameter ... : varargs

// Unsupported : soup_socket_connect_async : unsupported parameter callback : no type generator for SocketCallback (SoupSocketCallback) for param callback

// ConnectSync is a wrapper around the C function soup_socket_connect_sync.
func (recv *Socket) ConnectSync(cancellable *gio.Cancellable) uint32 {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.soup_socket_connect_sync((*C.SoupSocket)(recv.native), c_cancellable)
	retGo := (uint32)(retC)

	return retGo
}

// Disconnect is a wrapper around the C function soup_socket_disconnect.
func (recv *Socket) Disconnect() {
	C.soup_socket_disconnect((*C.SoupSocket)(recv.native))

	return
}

// GetFd is a wrapper around the C function soup_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	retC := C.soup_socket_get_fd((*C.SoupSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function soup_socket_get_local_address.
func (recv *Socket) GetLocalAddress() *Address {
	retC := C.soup_socket_get_local_address((*C.SoupSocket)(recv.native))
	retGo := AddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRemoteAddress is a wrapper around the C function soup_socket_get_remote_address.
func (recv *Socket) GetRemoteAddress() *Address {
	retC := C.soup_socket_get_remote_address((*C.SoupSocket)(recv.native))
	retGo := AddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsConnected is a wrapper around the C function soup_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	retC := C.soup_socket_is_connected((*C.SoupSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSsl is a wrapper around the C function soup_socket_is_ssl.
func (recv *Socket) IsSsl() bool {
	retC := C.soup_socket_is_ssl((*C.SoupSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listen is a wrapper around the C function soup_socket_listen.
func (recv *Socket) Listen() bool {
	retC := C.soup_socket_listen((*C.SoupSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Read is a wrapper around the C function soup_socket_read.
func (recv *Socket) Read(buffer []uint8, cancellable *gio.Cancellable) (SocketIOStatus, uint64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (C.gpointer)(unsafe.Pointer(c_buffer_arrayPtr))

	c_len := (C.gsize)(len(buffer))

	var c_nread C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_socket_read((*C.SoupSocket)(recv.native), c_buffer, c_len, &c_nread, c_cancellable, &cThrowableError)
	retGo := (SocketIOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	nread := (uint64)(c_nread)

	return retGo, nread, goError
}

// Blacklisted : soup_socket_read_until

// StartProxySsl is a wrapper around the C function soup_socket_start_proxy_ssl.
func (recv *Socket) StartProxySsl(sslHost string, cancellable *gio.Cancellable) bool {
	c_ssl_host := C.CString(sslHost)
	defer C.free(unsafe.Pointer(c_ssl_host))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.soup_socket_start_proxy_ssl((*C.SoupSocket)(recv.native), c_ssl_host, c_cancellable)
	retGo := retC == C.TRUE

	return retGo
}

// StartSsl is a wrapper around the C function soup_socket_start_ssl.
func (recv *Socket) StartSsl(cancellable *gio.Cancellable) bool {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.soup_socket_start_ssl((*C.SoupSocket)(recv.native), c_cancellable)
	retGo := retC == C.TRUE

	return retGo
}

// Write is a wrapper around the C function soup_socket_write.
func (recv *Socket) Write(buffer []uint8, cancellable *gio.Cancellable) (SocketIOStatus, uint64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (C.gconstpointer)(unsafe.Pointer(c_buffer_arrayPtr))

	c_len := (C.gsize)(len(buffer))

	var c_nwrote C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_socket_write((*C.SoupSocket)(recv.native), c_buffer, c_len, &c_nwrote, c_cancellable, &cThrowableError)
	retGo := (SocketIOStatus)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	nwrote := (uint64)(c_nwrote)

	return retGo, nwrote, goError
}

// Initable returns the Initable interface implemented by Socket
func (recv *Socket) Initable() *gio.Initable {
	return gio.InitableNewFromC(recv.ToC())
}

// WebsocketConnection is a wrapper around the C record SoupWebsocketConnection.
type WebsocketConnection struct {
	native *C.SoupWebsocketConnection
	// parent : record
	// Private : pv
}

func WebsocketConnectionNewFromC(u unsafe.Pointer) *WebsocketConnection {
	c := (*C.SoupWebsocketConnection)(u)
	if c == nil {
		return nil
	}

	g := &WebsocketConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebsocketConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebsocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsocketConnection with another WebsocketConnection, and returns true if they represent the same GObject.
func (recv *WebsocketConnection) Equals(other *WebsocketConnection) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WebsocketConnection) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WebsocketConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebsocketConnection.
func CastToWebsocketConnection(object *gobject.Object) *WebsocketConnection {
	return WebsocketConnectionNewFromC(object.ToC())
}

type signalWebsocketConnectionClosedDetail struct {
	callback  WebsocketConnectionSignalClosedCallback
	handlerID C.gulong
}

var signalWebsocketConnectionClosedId int
var signalWebsocketConnectionClosedMap = make(map[int]signalWebsocketConnectionClosedDetail)
var signalWebsocketConnectionClosedLock sync.RWMutex

// WebsocketConnectionSignalClosedCallback is a callback function for a 'closed' signal emitted from a WebsocketConnection.
type WebsocketConnectionSignalClosedCallback func(targetObject *WebsocketConnection)

/*
ConnectClosed connects the callback to the 'closed' signal for the WebsocketConnection.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *WebsocketConnection) ConnectClosed(callback WebsocketConnectionSignalClosedCallback) int {
	signalWebsocketConnectionClosedLock.Lock()
	defer signalWebsocketConnectionClosedLock.Unlock()

	signalWebsocketConnectionClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebsocketConnection_signal_connect_closed(instance, C.gpointer(uintptr(signalWebsocketConnectionClosedId)))

	detail := signalWebsocketConnectionClosedDetail{callback, handlerID}
	signalWebsocketConnectionClosedMap[signalWebsocketConnectionClosedId] = detail

	return signalWebsocketConnectionClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the WebsocketConnection.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *WebsocketConnection) DisconnectClosed(connectionID int) {
	signalWebsocketConnectionClosedLock.Lock()
	defer signalWebsocketConnectionClosedLock.Unlock()

	detail, exists := signalWebsocketConnectionClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebsocketConnectionClosedMap, connectionID)
}

//export websocketconnection_closedHandler
func websocketconnection_closedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebsocketConnectionClosedLock.RLock()
	defer signalWebsocketConnectionClosedLock.RUnlock()

	targetObject := WebsocketConnectionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebsocketConnectionClosedMap[index].callback
	callback(targetObject)
}

type signalWebsocketConnectionClosingDetail struct {
	callback  WebsocketConnectionSignalClosingCallback
	handlerID C.gulong
}

var signalWebsocketConnectionClosingId int
var signalWebsocketConnectionClosingMap = make(map[int]signalWebsocketConnectionClosingDetail)
var signalWebsocketConnectionClosingLock sync.RWMutex

// WebsocketConnectionSignalClosingCallback is a callback function for a 'closing' signal emitted from a WebsocketConnection.
type WebsocketConnectionSignalClosingCallback func(targetObject *WebsocketConnection)

/*
ConnectClosing connects the callback to the 'closing' signal for the WebsocketConnection.

The returned value represents the connection, and may be passed to DisconnectClosing to remove it.
*/
func (recv *WebsocketConnection) ConnectClosing(callback WebsocketConnectionSignalClosingCallback) int {
	signalWebsocketConnectionClosingLock.Lock()
	defer signalWebsocketConnectionClosingLock.Unlock()

	signalWebsocketConnectionClosingId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebsocketConnection_signal_connect_closing(instance, C.gpointer(uintptr(signalWebsocketConnectionClosingId)))

	detail := signalWebsocketConnectionClosingDetail{callback, handlerID}
	signalWebsocketConnectionClosingMap[signalWebsocketConnectionClosingId] = detail

	return signalWebsocketConnectionClosingId
}

/*
DisconnectClosing disconnects a callback from the 'closing' signal for the WebsocketConnection.

The connectionID should be a value returned from a call to ConnectClosing.
*/
func (recv *WebsocketConnection) DisconnectClosing(connectionID int) {
	signalWebsocketConnectionClosingLock.Lock()
	defer signalWebsocketConnectionClosingLock.Unlock()

	detail, exists := signalWebsocketConnectionClosingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebsocketConnectionClosingMap, connectionID)
}

//export websocketconnection_closingHandler
func websocketconnection_closingHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebsocketConnectionClosingLock.RLock()
	defer signalWebsocketConnectionClosingLock.RUnlock()

	targetObject := WebsocketConnectionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebsocketConnectionClosingMap[index].callback
	callback(targetObject)
}

type signalWebsocketConnectionErrorDetail struct {
	callback  WebsocketConnectionSignalErrorCallback
	handlerID C.gulong
}

var signalWebsocketConnectionErrorId int
var signalWebsocketConnectionErrorMap = make(map[int]signalWebsocketConnectionErrorDetail)
var signalWebsocketConnectionErrorLock sync.RWMutex

// WebsocketConnectionSignalErrorCallback is a callback function for a 'error' signal emitted from a WebsocketConnection.
type WebsocketConnectionSignalErrorCallback func(targetObject *WebsocketConnection, error *glib.Error)

/*
ConnectError connects the callback to the 'error' signal for the WebsocketConnection.

The returned value represents the connection, and may be passed to DisconnectError to remove it.
*/
func (recv *WebsocketConnection) ConnectError(callback WebsocketConnectionSignalErrorCallback) int {
	signalWebsocketConnectionErrorLock.Lock()
	defer signalWebsocketConnectionErrorLock.Unlock()

	signalWebsocketConnectionErrorId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebsocketConnection_signal_connect_error(instance, C.gpointer(uintptr(signalWebsocketConnectionErrorId)))

	detail := signalWebsocketConnectionErrorDetail{callback, handlerID}
	signalWebsocketConnectionErrorMap[signalWebsocketConnectionErrorId] = detail

	return signalWebsocketConnectionErrorId
}

/*
DisconnectError disconnects a callback from the 'error' signal for the WebsocketConnection.

The connectionID should be a value returned from a call to ConnectError.
*/
func (recv *WebsocketConnection) DisconnectError(connectionID int) {
	signalWebsocketConnectionErrorLock.Lock()
	defer signalWebsocketConnectionErrorLock.Unlock()

	detail, exists := signalWebsocketConnectionErrorMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebsocketConnectionErrorMap, connectionID)
}

//export websocketconnection_errorHandler
func websocketconnection_errorHandler(c_targetObject *C.GObject, c_error *C.GError, data C.gpointer) {
	signalWebsocketConnectionErrorLock.RLock()
	defer signalWebsocketConnectionErrorLock.RUnlock()

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	targetObject := WebsocketConnectionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebsocketConnectionErrorMap[index].callback
	callback(targetObject, error)
}

type signalWebsocketConnectionMessageDetail struct {
	callback  WebsocketConnectionSignalMessageCallback
	handlerID C.gulong
}

var signalWebsocketConnectionMessageId int
var signalWebsocketConnectionMessageMap = make(map[int]signalWebsocketConnectionMessageDetail)
var signalWebsocketConnectionMessageLock sync.RWMutex

// WebsocketConnectionSignalMessageCallback is a callback function for a 'message' signal emitted from a WebsocketConnection.
type WebsocketConnectionSignalMessageCallback func(targetObject *WebsocketConnection, type_ int32, message *glib.Bytes)

/*
ConnectMessage connects the callback to the 'message' signal for the WebsocketConnection.

The returned value represents the connection, and may be passed to DisconnectMessage to remove it.
*/
func (recv *WebsocketConnection) ConnectMessage(callback WebsocketConnectionSignalMessageCallback) int {
	signalWebsocketConnectionMessageLock.Lock()
	defer signalWebsocketConnectionMessageLock.Unlock()

	signalWebsocketConnectionMessageId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebsocketConnection_signal_connect_message(instance, C.gpointer(uintptr(signalWebsocketConnectionMessageId)))

	detail := signalWebsocketConnectionMessageDetail{callback, handlerID}
	signalWebsocketConnectionMessageMap[signalWebsocketConnectionMessageId] = detail

	return signalWebsocketConnectionMessageId
}

/*
DisconnectMessage disconnects a callback from the 'message' signal for the WebsocketConnection.

The connectionID should be a value returned from a call to ConnectMessage.
*/
func (recv *WebsocketConnection) DisconnectMessage(connectionID int) {
	signalWebsocketConnectionMessageLock.Lock()
	defer signalWebsocketConnectionMessageLock.Unlock()

	detail, exists := signalWebsocketConnectionMessageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebsocketConnectionMessageMap, connectionID)
}

//export websocketconnection_messageHandler
func websocketconnection_messageHandler(c_targetObject *C.GObject, c_type C.gint, c_message *C.GBytes, data C.gpointer) {
	signalWebsocketConnectionMessageLock.RLock()
	defer signalWebsocketConnectionMessageLock.RUnlock()

	type_ := int32(c_type)

	message := glib.BytesNewFromC(unsafe.Pointer(c_message))

	targetObject := WebsocketConnectionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebsocketConnectionMessageMap[index].callback
	callback(targetObject, type_, message)
}

// WebsocketConnectionNew is a wrapper around the C function soup_websocket_connection_new.
func WebsocketConnectionNew(stream *gio.IOStream, uri *URI, type_ WebsocketConnectionType, origin string, protocol string) *WebsocketConnection {
	c_stream := (*C.GIOStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GIOStream)(stream.ToC())
	}

	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_type := (C.SoupWebsocketConnectionType)(type_)

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	retC := C.soup_websocket_connection_new(c_stream, c_uri, c_type, c_origin, c_protocol)
	retGo := WebsocketConnectionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Close is a wrapper around the C function soup_websocket_connection_close.
func (recv *WebsocketConnection) Close(code uint32, data string) {
	c_code := (C.gushort)(code)

	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	C.soup_websocket_connection_close((*C.SoupWebsocketConnection)(recv.native), c_code, c_data)

	return
}

// GetCloseCode is a wrapper around the C function soup_websocket_connection_get_close_code.
func (recv *WebsocketConnection) GetCloseCode() uint32 {
	retC := C.soup_websocket_connection_get_close_code((*C.SoupWebsocketConnection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetCloseData is a wrapper around the C function soup_websocket_connection_get_close_data.
func (recv *WebsocketConnection) GetCloseData() string {
	retC := C.soup_websocket_connection_get_close_data((*C.SoupWebsocketConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetConnectionType is a wrapper around the C function soup_websocket_connection_get_connection_type.
func (recv *WebsocketConnection) GetConnectionType() WebsocketConnectionType {
	retC := C.soup_websocket_connection_get_connection_type((*C.SoupWebsocketConnection)(recv.native))
	retGo := (WebsocketConnectionType)(retC)

	return retGo
}

// GetIoStream is a wrapper around the C function soup_websocket_connection_get_io_stream.
func (recv *WebsocketConnection) GetIoStream() *gio.IOStream {
	retC := C.soup_websocket_connection_get_io_stream((*C.SoupWebsocketConnection)(recv.native))
	retGo := gio.IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOrigin is a wrapper around the C function soup_websocket_connection_get_origin.
func (recv *WebsocketConnection) GetOrigin() string {
	retC := C.soup_websocket_connection_get_origin((*C.SoupWebsocketConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProtocol is a wrapper around the C function soup_websocket_connection_get_protocol.
func (recv *WebsocketConnection) GetProtocol() string {
	retC := C.soup_websocket_connection_get_protocol((*C.SoupWebsocketConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetState is a wrapper around the C function soup_websocket_connection_get_state.
func (recv *WebsocketConnection) GetState() WebsocketState {
	retC := C.soup_websocket_connection_get_state((*C.SoupWebsocketConnection)(recv.native))
	retGo := (WebsocketState)(retC)

	return retGo
}

// GetUri is a wrapper around the C function soup_websocket_connection_get_uri.
func (recv *WebsocketConnection) GetUri() *URI {
	retC := C.soup_websocket_connection_get_uri((*C.SoupWebsocketConnection)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// SendBinary is a wrapper around the C function soup_websocket_connection_send_binary.
func (recv *WebsocketConnection) SendBinary(data []uint8) {
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	C.soup_websocket_connection_send_binary((*C.SoupWebsocketConnection)(recv.native), c_data, c_length)

	return
}

// SendText is a wrapper around the C function soup_websocket_connection_send_text.
func (recv *WebsocketConnection) SendText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.soup_websocket_connection_send_text((*C.SoupWebsocketConnection)(recv.native), c_text)

	return
}

const ADDRESS_ANY_PORT int32 = C.SOUP_ADDRESS_ANY_PORT
const ADDRESS_FAMILY string = C.SOUP_ADDRESS_FAMILY
const ADDRESS_NAME string = C.SOUP_ADDRESS_NAME
const ADDRESS_PHYSICAL string = C.SOUP_ADDRESS_PHYSICAL
const ADDRESS_PORT string = C.SOUP_ADDRESS_PORT
const ADDRESS_PROTOCOL string = C.SOUP_ADDRESS_PROTOCOL
const ADDRESS_SOCKADDR string = C.SOUP_ADDRESS_SOCKADDR
const AUTH_DOMAIN_ADD_PATH string = C.SOUP_AUTH_DOMAIN_ADD_PATH
const AUTH_DOMAIN_BASIC_AUTH_CALLBACK string = C.SOUP_AUTH_DOMAIN_BASIC_AUTH_CALLBACK
const AUTH_DOMAIN_BASIC_AUTH_DATA string = C.SOUP_AUTH_DOMAIN_BASIC_AUTH_DATA
const AUTH_DOMAIN_BASIC_H int32 = C.SOUP_AUTH_DOMAIN_BASIC_H
const AUTH_DOMAIN_DIGEST_AUTH_CALLBACK string = C.SOUP_AUTH_DOMAIN_DIGEST_AUTH_CALLBACK
const AUTH_DOMAIN_DIGEST_AUTH_DATA string = C.SOUP_AUTH_DOMAIN_DIGEST_AUTH_DATA
const AUTH_DOMAIN_DIGEST_H int32 = C.SOUP_AUTH_DOMAIN_DIGEST_H
const AUTH_DOMAIN_FILTER string = C.SOUP_AUTH_DOMAIN_FILTER
const AUTH_DOMAIN_FILTER_DATA string = C.SOUP_AUTH_DOMAIN_FILTER_DATA
const AUTH_DOMAIN_GENERIC_AUTH_CALLBACK string = C.SOUP_AUTH_DOMAIN_GENERIC_AUTH_CALLBACK
const AUTH_DOMAIN_GENERIC_AUTH_DATA string = C.SOUP_AUTH_DOMAIN_GENERIC_AUTH_DATA
const AUTH_DOMAIN_H int32 = C.SOUP_AUTH_DOMAIN_H
const AUTH_DOMAIN_PROXY string = C.SOUP_AUTH_DOMAIN_PROXY
const AUTH_DOMAIN_REALM string = C.SOUP_AUTH_DOMAIN_REALM
const AUTH_DOMAIN_REMOVE_PATH string = C.SOUP_AUTH_DOMAIN_REMOVE_PATH
const AUTH_H int32 = C.SOUP_AUTH_H
const AUTH_HOST string = C.SOUP_AUTH_HOST
const AUTH_IS_AUTHENTICATED string = C.SOUP_AUTH_IS_AUTHENTICATED
const AUTH_IS_FOR_PROXY string = C.SOUP_AUTH_IS_FOR_PROXY
const AUTH_MANAGER_H int32 = C.SOUP_AUTH_MANAGER_H
const AUTH_REALM string = C.SOUP_AUTH_REALM
const AUTH_SCHEME_NAME string = C.SOUP_AUTH_SCHEME_NAME
const CACHE_H int32 = C.SOUP_CACHE_H
const CHAR_HTTP_CTL int32 = C.SOUP_CHAR_HTTP_CTL
const CHAR_HTTP_SEPARATOR int32 = C.SOUP_CHAR_HTTP_SEPARATOR
const CHAR_URI_GEN_DELIMS int32 = C.SOUP_CHAR_URI_GEN_DELIMS
const CHAR_URI_PERCENT_ENCODED int32 = C.SOUP_CHAR_URI_PERCENT_ENCODED
const CHAR_URI_SUB_DELIMS int32 = C.SOUP_CHAR_URI_SUB_DELIMS
const CONTENT_DECODER_H int32 = C.SOUP_CONTENT_DECODER_H
const CONTENT_SNIFFER_H int32 = C.SOUP_CONTENT_SNIFFER_H
const COOKIE_H int32 = C.SOUP_COOKIE_H
const COOKIE_JAR_ACCEPT_POLICY string = C.SOUP_COOKIE_JAR_ACCEPT_POLICY
const COOKIE_JAR_DB_FILENAME string = C.SOUP_COOKIE_JAR_DB_FILENAME
const COOKIE_JAR_DB_H int32 = C.SOUP_COOKIE_JAR_DB_H
const COOKIE_JAR_H int32 = C.SOUP_COOKIE_JAR_H
const COOKIE_JAR_READ_ONLY string = C.SOUP_COOKIE_JAR_READ_ONLY
const COOKIE_JAR_TEXT_FILENAME string = C.SOUP_COOKIE_JAR_TEXT_FILENAME
const COOKIE_JAR_TEXT_H int32 = C.SOUP_COOKIE_JAR_TEXT_H
const COOKIE_MAX_AGE_ONE_DAY int32 = C.SOUP_COOKIE_MAX_AGE_ONE_DAY
const COOKIE_MAX_AGE_ONE_HOUR int32 = C.SOUP_COOKIE_MAX_AGE_ONE_HOUR
const COOKIE_MAX_AGE_ONE_WEEK int32 = C.SOUP_COOKIE_MAX_AGE_ONE_WEEK

// Blacklisted : COOKIE_MAX_AGE_ONE_YEAR

const DATE_H int32 = C.SOUP_DATE_H
const FORM_H int32 = C.SOUP_FORM_H
const FORM_MIME_TYPE_MULTIPART string = C.SOUP_FORM_MIME_TYPE_MULTIPART
const FORM_MIME_TYPE_URLENCODED string = C.SOUP_FORM_MIME_TYPE_URLENCODED
const HEADERS_H int32 = C.SOUP_HEADERS_H
const LOGGER_H int32 = C.SOUP_LOGGER_H
const MAJOR_VERSION int32 = C.SOUP_MAJOR_VERSION
const MESSAGE_BODY_H int32 = C.SOUP_MESSAGE_BODY_H
const MESSAGE_FIRST_PARTY string = C.SOUP_MESSAGE_FIRST_PARTY
const MESSAGE_FLAGS string = C.SOUP_MESSAGE_FLAGS
const MESSAGE_H int32 = C.SOUP_MESSAGE_H
const MESSAGE_HEADERS_H int32 = C.SOUP_MESSAGE_HEADERS_H
const MESSAGE_HTTP_VERSION string = C.SOUP_MESSAGE_HTTP_VERSION
const MESSAGE_METHOD string = C.SOUP_MESSAGE_METHOD
const MESSAGE_PRIORITY string = C.SOUP_MESSAGE_PRIORITY
const MESSAGE_REASON_PHRASE string = C.SOUP_MESSAGE_REASON_PHRASE
const MESSAGE_REQUEST_BODY string = C.SOUP_MESSAGE_REQUEST_BODY
const MESSAGE_REQUEST_BODY_DATA string = C.SOUP_MESSAGE_REQUEST_BODY_DATA
const MESSAGE_REQUEST_HEADERS string = C.SOUP_MESSAGE_REQUEST_HEADERS
const MESSAGE_RESPONSE_BODY string = C.SOUP_MESSAGE_RESPONSE_BODY
const MESSAGE_RESPONSE_BODY_DATA string = C.SOUP_MESSAGE_RESPONSE_BODY_DATA
const MESSAGE_RESPONSE_HEADERS string = C.SOUP_MESSAGE_RESPONSE_HEADERS
const MESSAGE_SERVER_SIDE string = C.SOUP_MESSAGE_SERVER_SIDE
const MESSAGE_STATUS_CODE string = C.SOUP_MESSAGE_STATUS_CODE
const MESSAGE_TLS_CERTIFICATE string = C.SOUP_MESSAGE_TLS_CERTIFICATE
const MESSAGE_TLS_ERRORS string = C.SOUP_MESSAGE_TLS_ERRORS
const MESSAGE_URI string = C.SOUP_MESSAGE_URI
const METHOD_H int32 = C.SOUP_METHOD_H
const MICRO_VERSION int32 = C.SOUP_MICRO_VERSION
const MINOR_VERSION int32 = C.SOUP_MINOR_VERSION
const MISC_H int32 = C.SOUP_MISC_H
const MULTIPART_H int32 = C.SOUP_MULTIPART_H
const MULTIPART_INPUT_STREAM_H int32 = C.SOUP_MULTIPART_INPUT_STREAM_H
const PASSWORD_MANAGER_H int32 = C.SOUP_PASSWORD_MANAGER_H
const PROXY_RESOLVER_DEFAULT_H int32 = C.SOUP_PROXY_RESOLVER_DEFAULT_H
const PROXY_URI_RESOLVER_H int32 = C.SOUP_PROXY_URI_RESOLVER_H

// Blacklisted : REQUESTER_H

const REQUEST_DATA_H int32 = C.SOUP_REQUEST_DATA_H
const REQUEST_FILE_H int32 = C.SOUP_REQUEST_FILE_H
const REQUEST_H int32 = C.SOUP_REQUEST_H
const REQUEST_HTTP_H int32 = C.SOUP_REQUEST_HTTP_H
const REQUEST_SESSION string = C.SOUP_REQUEST_SESSION
const REQUEST_URI string = C.SOUP_REQUEST_URI
const SERVER_ASYNC_CONTEXT string = C.SOUP_SERVER_ASYNC_CONTEXT
const SERVER_H int32 = C.SOUP_SERVER_H
const SERVER_HTTPS_ALIASES string = C.SOUP_SERVER_HTTPS_ALIASES
const SERVER_HTTP_ALIASES string = C.SOUP_SERVER_HTTP_ALIASES
const SERVER_INTERFACE string = C.SOUP_SERVER_INTERFACE
const SERVER_PORT string = C.SOUP_SERVER_PORT
const SERVER_RAW_PATHS string = C.SOUP_SERVER_RAW_PATHS
const SERVER_SERVER_HEADER string = C.SOUP_SERVER_SERVER_HEADER
const SERVER_SSL_CERT_FILE string = C.SOUP_SERVER_SSL_CERT_FILE
const SERVER_SSL_KEY_FILE string = C.SOUP_SERVER_SSL_KEY_FILE
const SERVER_TLS_CERTIFICATE string = C.SOUP_SERVER_TLS_CERTIFICATE
const SESSION_ACCEPT_LANGUAGE string = C.SOUP_SESSION_ACCEPT_LANGUAGE
const SESSION_ACCEPT_LANGUAGE_AUTO string = C.SOUP_SESSION_ACCEPT_LANGUAGE_AUTO
const SESSION_ADD_FEATURE string = C.SOUP_SESSION_ADD_FEATURE
const SESSION_ADD_FEATURE_BY_TYPE string = C.SOUP_SESSION_ADD_FEATURE_BY_TYPE
const SESSION_ASYNC_CONTEXT string = C.SOUP_SESSION_ASYNC_CONTEXT
const SESSION_ASYNC_H int32 = C.SOUP_SESSION_ASYNC_H
const SESSION_FEATURE_H int32 = C.SOUP_SESSION_FEATURE_H
const SESSION_H int32 = C.SOUP_SESSION_H
const SESSION_HTTPS_ALIASES string = C.SOUP_SESSION_HTTPS_ALIASES
const SESSION_HTTP_ALIASES string = C.SOUP_SESSION_HTTP_ALIASES
const SESSION_IDLE_TIMEOUT string = C.SOUP_SESSION_IDLE_TIMEOUT
const SESSION_LOCAL_ADDRESS string = C.SOUP_SESSION_LOCAL_ADDRESS
const SESSION_MAX_CONNS string = C.SOUP_SESSION_MAX_CONNS
const SESSION_MAX_CONNS_PER_HOST string = C.SOUP_SESSION_MAX_CONNS_PER_HOST
const SESSION_PROXY_RESOLVER string = C.SOUP_SESSION_PROXY_RESOLVER
const SESSION_PROXY_URI string = C.SOUP_SESSION_PROXY_URI
const SESSION_REMOVE_FEATURE_BY_TYPE string = C.SOUP_SESSION_REMOVE_FEATURE_BY_TYPE
const SESSION_SSL_CA_FILE string = C.SOUP_SESSION_SSL_CA_FILE
const SESSION_SSL_STRICT string = C.SOUP_SESSION_SSL_STRICT
const SESSION_SSL_USE_SYSTEM_CA_FILE string = C.SOUP_SESSION_SSL_USE_SYSTEM_CA_FILE
const SESSION_SYNC_H int32 = C.SOUP_SESSION_SYNC_H
const SESSION_TIMEOUT string = C.SOUP_SESSION_TIMEOUT
const SESSION_TLS_DATABASE string = C.SOUP_SESSION_TLS_DATABASE
const SESSION_TLS_INTERACTION string = C.SOUP_SESSION_TLS_INTERACTION
const SESSION_USER_AGENT string = C.SOUP_SESSION_USER_AGENT
const SESSION_USE_NTLM string = C.SOUP_SESSION_USE_NTLM
const SESSION_USE_THREAD_CONTEXT string = C.SOUP_SESSION_USE_THREAD_CONTEXT
const SOCKET_ASYNC_CONTEXT string = C.SOUP_SOCKET_ASYNC_CONTEXT
const SOCKET_FLAG_NONBLOCKING string = C.SOUP_SOCKET_FLAG_NONBLOCKING
const SOCKET_H int32 = C.SOUP_SOCKET_H
const SOCKET_IS_SERVER string = C.SOUP_SOCKET_IS_SERVER
const SOCKET_LOCAL_ADDRESS string = C.SOUP_SOCKET_LOCAL_ADDRESS
const SOCKET_REMOTE_ADDRESS string = C.SOUP_SOCKET_REMOTE_ADDRESS
const SOCKET_SSL_CREDENTIALS string = C.SOUP_SOCKET_SSL_CREDENTIALS
const SOCKET_SSL_FALLBACK string = C.SOUP_SOCKET_SSL_FALLBACK
const SOCKET_SSL_STRICT string = C.SOUP_SOCKET_SSL_STRICT
const SOCKET_TIMEOUT string = C.SOUP_SOCKET_TIMEOUT
const SOCKET_TLS_CERTIFICATE string = C.SOUP_SOCKET_TLS_CERTIFICATE
const SOCKET_TLS_ERRORS string = C.SOUP_SOCKET_TLS_ERRORS
const SOCKET_TRUSTED_CERTIFICATE string = C.SOUP_SOCKET_TRUSTED_CERTIFICATE
const SOCKET_USE_THREAD_CONTEXT string = C.SOUP_SOCKET_USE_THREAD_CONTEXT
const STATUS_H int32 = C.SOUP_STATUS_H
const TYPES_H int32 = C.SOUP_TYPES_H
const URI_H int32 = C.SOUP_URI_H
const VALUE_UTILS_H int32 = C.SOUP_VALUE_UTILS_H
const VERSION_MIN_REQUIRED int32 = C.SOUP_VERSION_MIN_REQUIRED
const XMLRPC_H int32 = C.SOUP_XMLRPC_H
const XMLRPC_OLD_H int32 = C.SOUP_XMLRPC_OLD_H

type AddressFamily C.SoupAddressFamily

const (
	SOUP_ADDRESS_FAMILY_INVALID AddressFamily = -1
	SOUP_ADDRESS_FAMILY_IPV4    AddressFamily = 2
	SOUP_ADDRESS_FAMILY_IPV6    AddressFamily = 10
)

type CacheResponse C.SoupCacheResponse

const (
	SOUP_CACHE_RESPONSE_FRESH            CacheResponse = 0
	SOUP_CACHE_RESPONSE_NEEDS_VALIDATION CacheResponse = 1
	SOUP_CACHE_RESPONSE_STALE            CacheResponse = 2
)

type CacheType C.SoupCacheType

const (
	SOUP_CACHE_SINGLE_USER CacheType = 0
	SOUP_CACHE_SHARED      CacheType = 1
)

type ConnectionState C.SoupConnectionState

const (
	SOUP_CONNECTION_NEW                 ConnectionState = 0
	SOUP_CONNECTION_CONNECTING          ConnectionState = 1
	SOUP_CONNECTION_IDLE                ConnectionState = 2
	SOUP_CONNECTION_IN_USE              ConnectionState = 3
	SOUP_CONNECTION_REMOTE_DISCONNECTED ConnectionState = 4
	SOUP_CONNECTION_DISCONNECTED        ConnectionState = 5
)

type CookieJarAcceptPolicy C.SoupCookieJarAcceptPolicy

const (
	SOUP_COOKIE_JAR_ACCEPT_ALWAYS         CookieJarAcceptPolicy = 0
	SOUP_COOKIE_JAR_ACCEPT_NEVER          CookieJarAcceptPolicy = 1
	SOUP_COOKIE_JAR_ACCEPT_NO_THIRD_PARTY CookieJarAcceptPolicy = 2
)

type DateFormat C.SoupDateFormat

const (
	SOUP_DATE_HTTP            DateFormat = 1
	SOUP_DATE_COOKIE          DateFormat = 2
	SOUP_DATE_RFC2822         DateFormat = 3
	SOUP_DATE_ISO8601_COMPACT DateFormat = 4
	SOUP_DATE_ISO8601_FULL    DateFormat = 5
	SOUP_DATE_ISO8601         DateFormat = 5
	SOUP_DATE_ISO8601_XMLRPC  DateFormat = 6
)

type Encoding C.SoupEncoding

const (
	SOUP_ENCODING_UNRECOGNIZED   Encoding = 0
	SOUP_ENCODING_NONE           Encoding = 1
	SOUP_ENCODING_CONTENT_LENGTH Encoding = 2
	SOUP_ENCODING_EOF            Encoding = 3
	SOUP_ENCODING_CHUNKED        Encoding = 4
	SOUP_ENCODING_BYTERANGES     Encoding = 5
)

type HTTPVersion C.SoupHTTPVersion

const (
	SOUP_HTTP_1_0 HTTPVersion = 0
	SOUP_HTTP_1_1 HTTPVersion = 1
)

type KnownStatusCode C.SoupKnownStatusCode

const (
	SOUP_KNOWN_STATUS_CODE_NONE                            KnownStatusCode = 0
	SOUP_KNOWN_STATUS_CODE_CANCELLED                       KnownStatusCode = 1
	SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE                    KnownStatusCode = 2
	SOUP_KNOWN_STATUS_CODE_CANT_RESOLVE_PROXY              KnownStatusCode = 3
	SOUP_KNOWN_STATUS_CODE_CANT_CONNECT                    KnownStatusCode = 4
	SOUP_KNOWN_STATUS_CODE_CANT_CONNECT_PROXY              KnownStatusCode = 5
	SOUP_KNOWN_STATUS_CODE_SSL_FAILED                      KnownStatusCode = 6
	SOUP_KNOWN_STATUS_CODE_IO_ERROR                        KnownStatusCode = 7
	SOUP_KNOWN_STATUS_CODE_MALFORMED                       KnownStatusCode = 8
	SOUP_KNOWN_STATUS_CODE_TRY_AGAIN                       KnownStatusCode = 9
	SOUP_KNOWN_STATUS_CODE_TOO_MANY_REDIRECTS              KnownStatusCode = 10
	SOUP_KNOWN_STATUS_CODE_TLS_FAILED                      KnownStatusCode = 11
	SOUP_KNOWN_STATUS_CODE_CONTINUE                        KnownStatusCode = 100
	SOUP_KNOWN_STATUS_CODE_SWITCHING_PROTOCOLS             KnownStatusCode = 101
	SOUP_KNOWN_STATUS_CODE_PROCESSING                      KnownStatusCode = 102
	SOUP_KNOWN_STATUS_CODE_OK                              KnownStatusCode = 200
	SOUP_KNOWN_STATUS_CODE_CREATED                         KnownStatusCode = 201
	SOUP_KNOWN_STATUS_CODE_ACCEPTED                        KnownStatusCode = 202
	SOUP_KNOWN_STATUS_CODE_NON_AUTHORITATIVE               KnownStatusCode = 203
	SOUP_KNOWN_STATUS_CODE_NO_CONTENT                      KnownStatusCode = 204
	SOUP_KNOWN_STATUS_CODE_RESET_CONTENT                   KnownStatusCode = 205
	SOUP_KNOWN_STATUS_CODE_PARTIAL_CONTENT                 KnownStatusCode = 206
	SOUP_KNOWN_STATUS_CODE_MULTI_STATUS                    KnownStatusCode = 207
	SOUP_KNOWN_STATUS_CODE_MULTIPLE_CHOICES                KnownStatusCode = 300
	SOUP_KNOWN_STATUS_CODE_MOVED_PERMANENTLY               KnownStatusCode = 301
	SOUP_KNOWN_STATUS_CODE_FOUND                           KnownStatusCode = 302
	SOUP_KNOWN_STATUS_CODE_MOVED_TEMPORARILY               KnownStatusCode = 302
	SOUP_KNOWN_STATUS_CODE_SEE_OTHER                       KnownStatusCode = 303
	SOUP_KNOWN_STATUS_CODE_NOT_MODIFIED                    KnownStatusCode = 304
	SOUP_KNOWN_STATUS_CODE_USE_PROXY                       KnownStatusCode = 305
	SOUP_KNOWN_STATUS_CODE_NOT_APPEARING_IN_THIS_PROTOCOL  KnownStatusCode = 306
	SOUP_KNOWN_STATUS_CODE_TEMPORARY_REDIRECT              KnownStatusCode = 307
	SOUP_KNOWN_STATUS_CODE_BAD_REQUEST                     KnownStatusCode = 400
	SOUP_KNOWN_STATUS_CODE_UNAUTHORIZED                    KnownStatusCode = 401
	SOUP_KNOWN_STATUS_CODE_PAYMENT_REQUIRED                KnownStatusCode = 402
	SOUP_KNOWN_STATUS_CODE_FORBIDDEN                       KnownStatusCode = 403
	SOUP_KNOWN_STATUS_CODE_NOT_FOUND                       KnownStatusCode = 404
	SOUP_KNOWN_STATUS_CODE_METHOD_NOT_ALLOWED              KnownStatusCode = 405
	SOUP_KNOWN_STATUS_CODE_NOT_ACCEPTABLE                  KnownStatusCode = 406
	SOUP_KNOWN_STATUS_CODE_PROXY_AUTHENTICATION_REQUIRED   KnownStatusCode = 407
	SOUP_KNOWN_STATUS_CODE_PROXY_UNAUTHORIZED              KnownStatusCode = 407
	SOUP_KNOWN_STATUS_CODE_REQUEST_TIMEOUT                 KnownStatusCode = 408
	SOUP_KNOWN_STATUS_CODE_CONFLICT                        KnownStatusCode = 409
	SOUP_KNOWN_STATUS_CODE_GONE                            KnownStatusCode = 410
	SOUP_KNOWN_STATUS_CODE_LENGTH_REQUIRED                 KnownStatusCode = 411
	SOUP_KNOWN_STATUS_CODE_PRECONDITION_FAILED             KnownStatusCode = 412
	SOUP_KNOWN_STATUS_CODE_REQUEST_ENTITY_TOO_LARGE        KnownStatusCode = 413
	SOUP_KNOWN_STATUS_CODE_REQUEST_URI_TOO_LONG            KnownStatusCode = 414
	SOUP_KNOWN_STATUS_CODE_UNSUPPORTED_MEDIA_TYPE          KnownStatusCode = 415
	SOUP_KNOWN_STATUS_CODE_REQUESTED_RANGE_NOT_SATISFIABLE KnownStatusCode = 416
	SOUP_KNOWN_STATUS_CODE_INVALID_RANGE                   KnownStatusCode = 416
	SOUP_KNOWN_STATUS_CODE_EXPECTATION_FAILED              KnownStatusCode = 417
	SOUP_KNOWN_STATUS_CODE_UNPROCESSABLE_ENTITY            KnownStatusCode = 422
	SOUP_KNOWN_STATUS_CODE_LOCKED                          KnownStatusCode = 423
	SOUP_KNOWN_STATUS_CODE_FAILED_DEPENDENCY               KnownStatusCode = 424
	SOUP_KNOWN_STATUS_CODE_INTERNAL_SERVER_ERROR           KnownStatusCode = 500
	SOUP_KNOWN_STATUS_CODE_NOT_IMPLEMENTED                 KnownStatusCode = 501
	SOUP_KNOWN_STATUS_CODE_BAD_GATEWAY                     KnownStatusCode = 502
	SOUP_KNOWN_STATUS_CODE_SERVICE_UNAVAILABLE             KnownStatusCode = 503
	SOUP_KNOWN_STATUS_CODE_GATEWAY_TIMEOUT                 KnownStatusCode = 504
	SOUP_KNOWN_STATUS_CODE_HTTP_VERSION_NOT_SUPPORTED      KnownStatusCode = 505
	SOUP_KNOWN_STATUS_CODE_INSUFFICIENT_STORAGE            KnownStatusCode = 507
	SOUP_KNOWN_STATUS_CODE_NOT_EXTENDED                    KnownStatusCode = 510
)

type LoggerLogLevel C.SoupLoggerLogLevel

const (
	SOUP_LOGGER_LOG_NONE    LoggerLogLevel = 0
	SOUP_LOGGER_LOG_MINIMAL LoggerLogLevel = 1
	SOUP_LOGGER_LOG_HEADERS LoggerLogLevel = 2
	SOUP_LOGGER_LOG_BODY    LoggerLogLevel = 3
)

type MemoryUse C.SoupMemoryUse

const (
	SOUP_MEMORY_STATIC    MemoryUse = 0
	SOUP_MEMORY_TAKE      MemoryUse = 1
	SOUP_MEMORY_COPY      MemoryUse = 2
	SOUP_MEMORY_TEMPORARY MemoryUse = 3
)

type MessageHeadersType C.SoupMessageHeadersType

const (
	SOUP_MESSAGE_HEADERS_REQUEST   MessageHeadersType = 0
	SOUP_MESSAGE_HEADERS_RESPONSE  MessageHeadersType = 1
	SOUP_MESSAGE_HEADERS_MULTIPART MessageHeadersType = 2
)

type MessagePriority C.SoupMessagePriority

const (
	SOUP_MESSAGE_PRIORITY_VERY_LOW  MessagePriority = 0
	SOUP_MESSAGE_PRIORITY_LOW       MessagePriority = 1
	SOUP_MESSAGE_PRIORITY_NORMAL    MessagePriority = 2
	SOUP_MESSAGE_PRIORITY_HIGH      MessagePriority = 3
	SOUP_MESSAGE_PRIORITY_VERY_HIGH MessagePriority = 4
)

type RequestError C.SoupRequestError

const (
	SOUP_REQUEST_ERROR_BAD_URI                RequestError = 0
	SOUP_REQUEST_ERROR_UNSUPPORTED_URI_SCHEME RequestError = 1
	SOUP_REQUEST_ERROR_PARSING                RequestError = 2
	SOUP_REQUEST_ERROR_ENCODING               RequestError = 3
)

// RequestErrorQuark is a wrapper around the C function soup_request_error_quark.
func RequestErrorQuark() glib.Quark {
	retC := C.soup_request_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Blacklisted : SoupRequesterError

type SocketIOStatus C.SoupSocketIOStatus

const (
	SOUP_SOCKET_OK          SocketIOStatus = 0
	SOUP_SOCKET_WOULD_BLOCK SocketIOStatus = 1
	SOUP_SOCKET_EOF         SocketIOStatus = 2
	SOUP_SOCKET_ERROR       SocketIOStatus = 3
)

type Status C.SoupStatus

const (
	SOUP_STATUS_NONE                            Status = 0
	SOUP_STATUS_CANCELLED                       Status = 1
	SOUP_STATUS_CANT_RESOLVE                    Status = 2
	SOUP_STATUS_CANT_RESOLVE_PROXY              Status = 3
	SOUP_STATUS_CANT_CONNECT                    Status = 4
	SOUP_STATUS_CANT_CONNECT_PROXY              Status = 5
	SOUP_STATUS_SSL_FAILED                      Status = 6
	SOUP_STATUS_IO_ERROR                        Status = 7
	SOUP_STATUS_MALFORMED                       Status = 8
	SOUP_STATUS_TRY_AGAIN                       Status = 9
	SOUP_STATUS_TOO_MANY_REDIRECTS              Status = 10
	SOUP_STATUS_TLS_FAILED                      Status = 11
	SOUP_STATUS_CONTINUE                        Status = 100
	SOUP_STATUS_SWITCHING_PROTOCOLS             Status = 101
	SOUP_STATUS_PROCESSING                      Status = 102
	SOUP_STATUS_OK                              Status = 200
	SOUP_STATUS_CREATED                         Status = 201
	SOUP_STATUS_ACCEPTED                        Status = 202
	SOUP_STATUS_NON_AUTHORITATIVE               Status = 203
	SOUP_STATUS_NO_CONTENT                      Status = 204
	SOUP_STATUS_RESET_CONTENT                   Status = 205
	SOUP_STATUS_PARTIAL_CONTENT                 Status = 206
	SOUP_STATUS_MULTI_STATUS                    Status = 207
	SOUP_STATUS_MULTIPLE_CHOICES                Status = 300
	SOUP_STATUS_MOVED_PERMANENTLY               Status = 301
	SOUP_STATUS_FOUND                           Status = 302
	SOUP_STATUS_MOVED_TEMPORARILY               Status = 302
	SOUP_STATUS_SEE_OTHER                       Status = 303
	SOUP_STATUS_NOT_MODIFIED                    Status = 304
	SOUP_STATUS_USE_PROXY                       Status = 305
	SOUP_STATUS_NOT_APPEARING_IN_THIS_PROTOCOL  Status = 306
	SOUP_STATUS_TEMPORARY_REDIRECT              Status = 307
	SOUP_STATUS_BAD_REQUEST                     Status = 400
	SOUP_STATUS_UNAUTHORIZED                    Status = 401
	SOUP_STATUS_PAYMENT_REQUIRED                Status = 402
	SOUP_STATUS_FORBIDDEN                       Status = 403
	SOUP_STATUS_NOT_FOUND                       Status = 404
	SOUP_STATUS_METHOD_NOT_ALLOWED              Status = 405
	SOUP_STATUS_NOT_ACCEPTABLE                  Status = 406
	SOUP_STATUS_PROXY_AUTHENTICATION_REQUIRED   Status = 407
	SOUP_STATUS_PROXY_UNAUTHORIZED              Status = 407
	SOUP_STATUS_REQUEST_TIMEOUT                 Status = 408
	SOUP_STATUS_CONFLICT                        Status = 409
	SOUP_STATUS_GONE                            Status = 410
	SOUP_STATUS_LENGTH_REQUIRED                 Status = 411
	SOUP_STATUS_PRECONDITION_FAILED             Status = 412
	SOUP_STATUS_REQUEST_ENTITY_TOO_LARGE        Status = 413
	SOUP_STATUS_REQUEST_URI_TOO_LONG            Status = 414
	SOUP_STATUS_UNSUPPORTED_MEDIA_TYPE          Status = 415
	SOUP_STATUS_REQUESTED_RANGE_NOT_SATISFIABLE Status = 416
	SOUP_STATUS_INVALID_RANGE                   Status = 416
	SOUP_STATUS_EXPECTATION_FAILED              Status = 417
	SOUP_STATUS_UNPROCESSABLE_ENTITY            Status = 422
	SOUP_STATUS_LOCKED                          Status = 423
	SOUP_STATUS_FAILED_DEPENDENCY               Status = 424
	SOUP_STATUS_INTERNAL_SERVER_ERROR           Status = 500
	SOUP_STATUS_NOT_IMPLEMENTED                 Status = 501
	SOUP_STATUS_BAD_GATEWAY                     Status = 502
	SOUP_STATUS_SERVICE_UNAVAILABLE             Status = 503
	SOUP_STATUS_GATEWAY_TIMEOUT                 Status = 504
	SOUP_STATUS_HTTP_VERSION_NOT_SUPPORTED      Status = 505
	SOUP_STATUS_INSUFFICIENT_STORAGE            Status = 507
	SOUP_STATUS_NOT_EXTENDED                    Status = 510
)

// StatusGetPhrase is a wrapper around the C function soup_status_get_phrase.
func StatusGetPhrase(statusCode uint32) string {
	c_status_code := (C.guint)(statusCode)

	retC := C.soup_status_get_phrase(c_status_code)
	retGo := C.GoString(retC)

	return retGo
}

// StatusProxify is a wrapper around the C function soup_status_proxify.
func StatusProxify(statusCode uint32) uint32 {
	c_status_code := (C.guint)(statusCode)

	retC := C.soup_status_proxify(c_status_code)
	retGo := (uint32)(retC)

	return retGo
}

type TLDError C.SoupTLDError

const (
	SOUP_TLD_ERROR_INVALID_HOSTNAME   TLDError = 0
	SOUP_TLD_ERROR_IS_IP_ADDRESS      TLDError = 1
	SOUP_TLD_ERROR_NOT_ENOUGH_DOMAINS TLDError = 2
	SOUP_TLD_ERROR_NO_BASE_DOMAIN     TLDError = 3
)

// TLDErrorQuark is a wrapper around the C function soup_tld_error_quark.
func TLDErrorQuark() glib.Quark {
	retC := C.soup_tld_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type WebsocketCloseCode C.SoupWebsocketCloseCode

const (
	SOUP_WEBSOCKET_CLOSE_NORMAL           WebsocketCloseCode = 1000
	SOUP_WEBSOCKET_CLOSE_GOING_AWAY       WebsocketCloseCode = 1001
	SOUP_WEBSOCKET_CLOSE_PROTOCOL_ERROR   WebsocketCloseCode = 1002
	SOUP_WEBSOCKET_CLOSE_UNSUPPORTED_DATA WebsocketCloseCode = 1003
	SOUP_WEBSOCKET_CLOSE_NO_STATUS        WebsocketCloseCode = 1005
	SOUP_WEBSOCKET_CLOSE_ABNORMAL         WebsocketCloseCode = 1006
	SOUP_WEBSOCKET_CLOSE_BAD_DATA         WebsocketCloseCode = 1007
	SOUP_WEBSOCKET_CLOSE_POLICY_VIOLATION WebsocketCloseCode = 1008
	SOUP_WEBSOCKET_CLOSE_TOO_BIG          WebsocketCloseCode = 1009
	SOUP_WEBSOCKET_CLOSE_NO_EXTENSION     WebsocketCloseCode = 1010
	SOUP_WEBSOCKET_CLOSE_SERVER_ERROR     WebsocketCloseCode = 1011
	SOUP_WEBSOCKET_CLOSE_TLS_HANDSHAKE    WebsocketCloseCode = 1015
)

type WebsocketConnectionType C.SoupWebsocketConnectionType

const (
	SOUP_WEBSOCKET_CONNECTION_UNKNOWN WebsocketConnectionType = 0
	SOUP_WEBSOCKET_CONNECTION_CLIENT  WebsocketConnectionType = 1
	SOUP_WEBSOCKET_CONNECTION_SERVER  WebsocketConnectionType = 2
)

type WebsocketDataType C.SoupWebsocketDataType

const (
	SOUP_WEBSOCKET_DATA_TEXT   WebsocketDataType = 1
	SOUP_WEBSOCKET_DATA_BINARY WebsocketDataType = 2
)

type WebsocketError C.SoupWebsocketError

const (
	SOUP_WEBSOCKET_ERROR_FAILED        WebsocketError = 0
	SOUP_WEBSOCKET_ERROR_NOT_WEBSOCKET WebsocketError = 1
	SOUP_WEBSOCKET_ERROR_BAD_HANDSHAKE WebsocketError = 2
	SOUP_WEBSOCKET_ERROR_BAD_ORIGIN    WebsocketError = 3
)

// WebsocketErrorGetQuark is a wrapper around the C function soup_websocket_error_get_quark.
func WebsocketErrorGetQuark() glib.Quark {
	retC := C.soup_websocket_error_get_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type WebsocketState C.SoupWebsocketState

const (
	SOUP_WEBSOCKET_STATE_OPEN    WebsocketState = 1
	SOUP_WEBSOCKET_STATE_CLOSING WebsocketState = 2
	SOUP_WEBSOCKET_STATE_CLOSED  WebsocketState = 3
)

type XMLRPCError C.SoupXMLRPCError

const (
	SOUP_XMLRPC_ERROR_ARGUMENTS XMLRPCError = 0
	SOUP_XMLRPC_ERROR_RETVAL    XMLRPCError = 1
)

// XMLRPCErrorQuark is a wrapper around the C function soup_xmlrpc_error_quark.
func XMLRPCErrorQuark() glib.Quark {
	retC := C.soup_xmlrpc_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type XMLRPCFault C.SoupXMLRPCFault

const (
	SOUP_XMLRPC_FAULT_PARSE_ERROR_NOT_WELL_FORMED                XMLRPCFault = -32700
	SOUP_XMLRPC_FAULT_PARSE_ERROR_UNSUPPORTED_ENCODING           XMLRPCFault = -32701
	SOUP_XMLRPC_FAULT_PARSE_ERROR_INVALID_CHARACTER_FOR_ENCODING XMLRPCFault = -32702
	SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_XML_RPC               XMLRPCFault = -32600
	SOUP_XMLRPC_FAULT_SERVER_ERROR_REQUESTED_METHOD_NOT_FOUND    XMLRPCFault = -32601
	SOUP_XMLRPC_FAULT_SERVER_ERROR_INVALID_METHOD_PARAMETERS     XMLRPCFault = -32602
	SOUP_XMLRPC_FAULT_SERVER_ERROR_INTERNAL_XML_RPC_ERROR        XMLRPCFault = -32603
	SOUP_XMLRPC_FAULT_APPLICATION_ERROR                          XMLRPCFault = -32500
	SOUP_XMLRPC_FAULT_SYSTEM_ERROR                               XMLRPCFault = -32400
	SOUP_XMLRPC_FAULT_TRANSPORT_ERROR                            XMLRPCFault = -32300
)

// XMLRPCFaultQuark is a wrapper around the C function soup_xmlrpc_fault_quark.
func XMLRPCFaultQuark() glib.Quark {
	retC := C.soup_xmlrpc_fault_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : soup_add_completion : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : soup_add_idle : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// Unsupported : soup_add_io_watch : unsupported parameter function : no type generator for GLib.IOFunc (GIOFunc) for param function

// Unsupported : soup_add_timeout : unsupported parameter function : no type generator for GLib.SourceFunc (GSourceFunc) for param function

// CheckVersion is a wrapper around the C function soup_check_version.
func CheckVersion(major uint32, minor uint32, micro uint32) bool {
	c_major := (C.guint)(major)

	c_minor := (C.guint)(minor)

	c_micro := (C.guint)(micro)

	retC := C.soup_check_version(c_major, c_minor, c_micro)
	retGo := retC == C.TRUE

	return retGo
}

// CookiesFree is a wrapper around the C function soup_cookies_free.
func CookiesFree(cookies *glib.SList) {
	c_cookies := (*C.GSList)(C.NULL)
	if cookies != nil {
		c_cookies = (*C.GSList)(cookies.ToC())
	}

	C.soup_cookies_free(c_cookies)

	return
}

// CookiesFromRequest is a wrapper around the C function soup_cookies_from_request.
func CookiesFromRequest(msg *Message) *glib.SList {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_cookies_from_request(c_msg)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CookiesFromResponse is a wrapper around the C function soup_cookies_from_response.
func CookiesFromResponse(msg *Message) *glib.SList {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	retC := C.soup_cookies_from_response(c_msg)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CookiesToCookieHeader is a wrapper around the C function soup_cookies_to_cookie_header.
func CookiesToCookieHeader(cookies *glib.SList) string {
	c_cookies := (*C.GSList)(C.NULL)
	if cookies != nil {
		c_cookies = (*C.GSList)(cookies.ToC())
	}

	retC := C.soup_cookies_to_cookie_header(c_cookies)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// CookiesToRequest is a wrapper around the C function soup_cookies_to_request.
func CookiesToRequest(cookies *glib.SList, msg *Message) {
	c_cookies := (*C.GSList)(C.NULL)
	if cookies != nil {
		c_cookies = (*C.GSList)(cookies.ToC())
	}

	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_cookies_to_request(c_cookies, c_msg)

	return
}

// CookiesToResponse is a wrapper around the C function soup_cookies_to_response.
func CookiesToResponse(cookies *glib.SList, msg *Message) {
	c_cookies := (*C.GSList)(C.NULL)
	if cookies != nil {
		c_cookies = (*C.GSList)(cookies.ToC())
	}

	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	C.soup_cookies_to_response(c_cookies, c_msg)

	return
}

// FormDecode is a wrapper around the C function soup_form_decode.
func FormDecode(encodedForm string) *glib.HashTable {
	c_encoded_form := C.CString(encodedForm)
	defer C.free(unsafe.Pointer(c_encoded_form))

	retC := C.soup_form_decode(c_encoded_form)
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FormDecodeMultipart is a wrapper around the C function soup_form_decode_multipart.
func FormDecodeMultipart(msg *Message, fileControlName string) (*glib.HashTable, string, string, *Buffer) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_file_control_name := C.CString(fileControlName)
	defer C.free(unsafe.Pointer(c_file_control_name))

	var c_filename *C.char

	var c_content_type *C.char

	var c_file *C.SoupBuffer

	retC := C.soup_form_decode_multipart(c_msg, c_file_control_name, &c_filename, &c_content_type, &c_file)
	var retGo (*glib.HashTable)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.HashTableNewFromC(unsafe.Pointer(retC))
	}

	filename := C.GoString(c_filename)
	defer C.free(unsafe.Pointer(c_filename))

	contentType := C.GoString(c_content_type)
	defer C.free(unsafe.Pointer(c_content_type))

	file := BufferNewFromC(unsafe.Pointer(c_file))

	return retGo, filename, contentType, file
}

// Unsupported : soup_form_encode : unsupported parameter ... : varargs

// FormEncodeDatalist is a wrapper around the C function soup_form_encode_datalist.
func FormEncodeDatalist(formDataSet *glib.Data) string {
	c_form_data_set := (**C.GData)(C.NULL)
	if formDataSet != nil {
		c_form_data_set = (**C.GData)(formDataSet.ToC())
	}

	retC := C.soup_form_encode_datalist(c_form_data_set)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormEncodeHash is a wrapper around the C function soup_form_encode_hash.
func FormEncodeHash(formDataSet *glib.HashTable) string {
	c_form_data_set := (*C.GHashTable)(C.NULL)
	if formDataSet != nil {
		c_form_data_set = (*C.GHashTable)(formDataSet.ToC())
	}

	retC := C.soup_form_encode_hash(c_form_data_set)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_form_encode_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : soup_form_request_new : unsupported parameter ... : varargs

// FormRequestNewFromDatalist is a wrapper around the C function soup_form_request_new_from_datalist.
func FormRequestNewFromDatalist(method string, uri string, formDataSet *glib.Data) *Message {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_form_data_set := (**C.GData)(C.NULL)
	if formDataSet != nil {
		c_form_data_set = (**C.GData)(formDataSet.ToC())
	}

	retC := C.soup_form_request_new_from_datalist(c_method, c_uri, c_form_data_set)
	retGo := MessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FormRequestNewFromHash is a wrapper around the C function soup_form_request_new_from_hash.
func FormRequestNewFromHash(method string, uri string, formDataSet *glib.HashTable) *Message {
	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_form_data_set := (*C.GHashTable)(C.NULL)
	if formDataSet != nil {
		c_form_data_set = (*C.GHashTable)(formDataSet.ToC())
	}

	retC := C.soup_form_request_new_from_hash(c_method, c_uri, c_form_data_set)
	retGo := MessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FormRequestNewFromMultipart is a wrapper around the C function soup_form_request_new_from_multipart.
func FormRequestNewFromMultipart(uri string, multipart *Multipart) *Message {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_multipart := (*C.SoupMultipart)(C.NULL)
	if multipart != nil {
		c_multipart = (*C.SoupMultipart)(multipart.ToC())
	}

	retC := C.soup_form_request_new_from_multipart(c_uri, c_multipart)
	retGo := MessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMajorVersion is a wrapper around the C function soup_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.soup_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function soup_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.soup_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function soup_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.soup_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// HeaderContains is a wrapper around the C function soup_header_contains.
func HeaderContains(header string, token string) bool {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	c_token := C.CString(token)
	defer C.free(unsafe.Pointer(c_token))

	retC := C.soup_header_contains(c_header, c_token)
	retGo := retC == C.TRUE

	return retGo
}

// HeaderFreeList is a wrapper around the C function soup_header_free_list.
func HeaderFreeList(list *glib.SList) {
	c_list := (*C.GSList)(C.NULL)
	if list != nil {
		c_list = (*C.GSList)(list.ToC())
	}

	C.soup_header_free_list(c_list)

	return
}

// HeaderFreeParamList is a wrapper around the C function soup_header_free_param_list.
func HeaderFreeParamList(paramList *glib.HashTable) {
	c_param_list := (*C.GHashTable)(C.NULL)
	if paramList != nil {
		c_param_list = (*C.GHashTable)(paramList.ToC())
	}

	C.soup_header_free_param_list(c_param_list)

	return
}

// HeaderGStringAppendParam is a wrapper around the C function soup_header_g_string_append_param.
func HeaderGStringAppendParam(string_ *glib.String, name string, value string) {
	c_string := (*C.GString)(C.NULL)
	if string_ != nil {
		c_string = (*C.GString)(string_.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.soup_header_g_string_append_param(c_string, c_name, c_value)

	return
}

// HeaderGStringAppendParamQuoted is a wrapper around the C function soup_header_g_string_append_param_quoted.
func HeaderGStringAppendParamQuoted(string_ *glib.String, name string, value string) {
	c_string := (*C.GString)(C.NULL)
	if string_ != nil {
		c_string = (*C.GString)(string_.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.soup_header_g_string_append_param_quoted(c_string, c_name, c_value)

	return
}

// HeaderParseList is a wrapper around the C function soup_header_parse_list.
func HeaderParseList(header string) *glib.SList {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	retC := C.soup_header_parse_list(c_header)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HeaderParseParamList is a wrapper around the C function soup_header_parse_param_list.
func HeaderParseParamList(header string) *glib.HashTable {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	retC := C.soup_header_parse_param_list(c_header)
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HeaderParseQualityList is a wrapper around the C function soup_header_parse_quality_list.
func HeaderParseQualityList(header string) (*glib.SList, *glib.SList) {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	var c_unacceptable *C.GSList

	retC := C.soup_header_parse_quality_list(c_header, &c_unacceptable)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	unacceptable := glib.SListNewFromC(unsafe.Pointer(c_unacceptable))

	return retGo, unacceptable
}

// HeaderParseSemiParamList is a wrapper around the C function soup_header_parse_semi_param_list.
func HeaderParseSemiParamList(header string) *glib.HashTable {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	retC := C.soup_header_parse_semi_param_list(c_header)
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HeadersParse is a wrapper around the C function soup_headers_parse.
func HeadersParse(str string, len int32, dest *MessageHeaders) bool {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.int)(len)

	c_dest := (*C.SoupMessageHeaders)(C.NULL)
	if dest != nil {
		c_dest = (*C.SoupMessageHeaders)(dest.ToC())
	}

	retC := C.soup_headers_parse(c_str, c_len, c_dest)
	retGo := retC == C.TRUE

	return retGo
}

// HeadersParseRequest is a wrapper around the C function soup_headers_parse_request.
func HeadersParseRequest(str string, len int32, reqHeaders *MessageHeaders) (uint32, string, string, HTTPVersion) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.int)(len)

	c_req_headers := (*C.SoupMessageHeaders)(C.NULL)
	if reqHeaders != nil {
		c_req_headers = (*C.SoupMessageHeaders)(reqHeaders.ToC())
	}

	var c_req_method *C.char

	var c_req_path *C.char

	var c_ver C.SoupHTTPVersion

	retC := C.soup_headers_parse_request(c_str, c_len, c_req_headers, &c_req_method, &c_req_path, &c_ver)
	retGo := (uint32)(retC)

	reqMethod := C.GoString(c_req_method)
	defer C.free(unsafe.Pointer(c_req_method))

	reqPath := C.GoString(c_req_path)
	defer C.free(unsafe.Pointer(c_req_path))

	ver := (HTTPVersion)(c_ver)

	return retGo, reqMethod, reqPath, ver
}

// HeadersParseResponse is a wrapper around the C function soup_headers_parse_response.
func HeadersParseResponse(str string, len int32, headers *MessageHeaders) (bool, HTTPVersion, uint32, string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_len := (C.int)(len)

	c_headers := (*C.SoupMessageHeaders)(C.NULL)
	if headers != nil {
		c_headers = (*C.SoupMessageHeaders)(headers.ToC())
	}

	var c_ver C.SoupHTTPVersion

	var c_status_code C.guint

	var c_reason_phrase *C.char

	retC := C.soup_headers_parse_response(c_str, c_len, c_headers, &c_ver, &c_status_code, &c_reason_phrase)
	retGo := retC == C.TRUE

	ver := (HTTPVersion)(c_ver)

	statusCode := (uint32)(c_status_code)

	reasonPhrase := C.GoString(c_reason_phrase)
	defer C.free(unsafe.Pointer(c_reason_phrase))

	return retGo, ver, statusCode, reasonPhrase
}

// HeadersParseStatusLine is a wrapper around the C function soup_headers_parse_status_line.
func HeadersParseStatusLine(statusLine string) (bool, HTTPVersion, uint32, string) {
	c_status_line := C.CString(statusLine)
	defer C.free(unsafe.Pointer(c_status_line))

	var c_ver C.SoupHTTPVersion

	var c_status_code C.guint

	var c_reason_phrase *C.char

	retC := C.soup_headers_parse_status_line(c_status_line, &c_ver, &c_status_code, &c_reason_phrase)
	retGo := retC == C.TRUE

	ver := (HTTPVersion)(c_ver)

	statusCode := (uint32)(c_status_code)

	reasonPhrase := C.GoString(c_reason_phrase)
	defer C.free(unsafe.Pointer(c_reason_phrase))

	return retGo, ver, statusCode, reasonPhrase
}

// HttpErrorQuark is a wrapper around the C function soup_http_error_quark.
func HttpErrorQuark() glib.Quark {
	retC := C.soup_http_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// StrCaseEqual is a wrapper around the C function soup_str_case_equal.
func StrCaseEqual(v1 uintptr, v2 uintptr) bool {
	c_v1 := (C.gconstpointer)(v1)

	c_v2 := (C.gconstpointer)(v2)

	retC := C.soup_str_case_equal(c_v1, c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// StrCaseHash is a wrapper around the C function soup_str_case_hash.
func StrCaseHash(key uintptr) uint32 {
	c_key := (C.gconstpointer)(key)

	retC := C.soup_str_case_hash(c_key)
	retGo := (uint32)(retC)

	return retGo
}

// TldDomainIsPublicSuffix is a wrapper around the C function soup_tld_domain_is_public_suffix.
func TldDomainIsPublicSuffix(domain string) bool {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	retC := C.soup_tld_domain_is_public_suffix(c_domain)
	retGo := retC == C.TRUE

	return retGo
}

// TldGetBaseDomain is a wrapper around the C function soup_tld_get_base_domain.
func TldGetBaseDomain(hostname string) (string, error) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	var cThrowableError *C.GError

	retC := C.soup_tld_get_base_domain(c_hostname, &cThrowableError)
	retGo := C.GoString(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : soup_value_array_append : unsupported parameter ... : varargs

// Unsupported : soup_value_array_append_vals : unsupported parameter ... : varargs

// Unsupported : soup_value_array_from_args : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : soup_value_array_get_nth : unsupported parameter ... : varargs

// Unsupported : soup_value_array_insert : unsupported parameter ... : varargs

// ValueArrayNew is a wrapper around the C function soup_value_array_new.
func ValueArrayNew() *gobject.ValueArray {
	retC := C.soup_value_array_new()
	retGo := gobject.ValueArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_value_array_new_with_vals : unsupported parameter ... : varargs

// Unsupported : soup_value_array_to_args : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : soup_value_hash_insert : unsupported parameter ... : varargs

// Unsupported : soup_value_hash_insert_vals : unsupported parameter ... : varargs

// ValueHashInsertValue is a wrapper around the C function soup_value_hash_insert_value.
func ValueHashInsertValue(hash *glib.HashTable, key string, value *gobject.Value) {
	c_hash := (*C.GHashTable)(C.NULL)
	if hash != nil {
		c_hash = (*C.GHashTable)(hash.ToC())
	}

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.soup_value_hash_insert_value(c_hash, c_key, c_value)

	return
}

// Unsupported : soup_value_hash_lookup : unsupported parameter ... : varargs

// Unsupported : soup_value_hash_lookup_vals : unsupported parameter ... : varargs

// ValueHashNew is a wrapper around the C function soup_value_hash_new.
func ValueHashNew() *glib.HashTable {
	retC := C.soup_value_hash_new()
	retGo := glib.HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_value_hash_new_with_vals : unsupported parameter ... : varargs

// WebsocketClientPrepareHandshake is a wrapper around the C function soup_websocket_client_prepare_handshake.
func WebsocketClientPrepareHandshake(msg *Message, origin string, protocols []string) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	c_protocols_array := make([]*C.char, len(protocols)+1, len(protocols)+1)
	for i, item := range protocols {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_protocols_array[i] = c
	}
	c_protocols_array[len(protocols)] = nil
	c_protocols_arrayPtr := &c_protocols_array[0]
	c_protocols := (**C.char)(unsafe.Pointer(c_protocols_arrayPtr))

	C.soup_websocket_client_prepare_handshake(c_msg, c_origin, c_protocols)

	return
}

// WebsocketClientVerifyHandshake is a wrapper around the C function soup_websocket_client_verify_handshake.
func WebsocketClientVerifyHandshake(msg *Message) (bool, error) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	var cThrowableError *C.GError

	retC := C.soup_websocket_client_verify_handshake(c_msg, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WebsocketServerCheckHandshake is a wrapper around the C function soup_websocket_server_check_handshake.
func WebsocketServerCheckHandshake(msg *Message, origin string, protocols []string) (bool, error) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	c_protocols_array := make([]*C.char, len(protocols)+1, len(protocols)+1)
	for i, item := range protocols {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_protocols_array[i] = c
	}
	c_protocols_array[len(protocols)] = nil
	c_protocols_arrayPtr := &c_protocols_array[0]
	c_protocols := (**C.char)(unsafe.Pointer(c_protocols_arrayPtr))

	var cThrowableError *C.GError

	retC := C.soup_websocket_server_check_handshake(c_msg, c_origin, c_protocols, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WebsocketServerProcessHandshake is a wrapper around the C function soup_websocket_server_process_handshake.
func WebsocketServerProcessHandshake(msg *Message, expectedOrigin string, protocols []string) bool {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_expected_origin := C.CString(expectedOrigin)
	defer C.free(unsafe.Pointer(c_expected_origin))

	c_protocols_array := make([]*C.char, len(protocols)+1, len(protocols)+1)
	for i, item := range protocols {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_protocols_array[i] = c
	}
	c_protocols_array[len(protocols)] = nil
	c_protocols_arrayPtr := &c_protocols_array[0]
	c_protocols := (**C.char)(unsafe.Pointer(c_protocols_arrayPtr))

	retC := C.soup_websocket_server_process_handshake(c_msg, c_expected_origin, c_protocols)
	retGo := retC == C.TRUE

	return retGo
}

// XmlrpcBuildFault is a wrapper around the C function soup_xmlrpc_build_fault.
func XmlrpcBuildFault(faultCode int32, faultFormat string, args ...interface{}) string {
	c_fault_code := (C.int)(faultCode)

	goFormattedString := fmt.Sprintf(faultFormat, args...)
	c_fault_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_fault_format))

	retC := C._soup_xmlrpc_build_fault(c_fault_code, c_fault_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_xmlrpc_build_method_call : unsupported parameter params :

// XmlrpcBuildMethodResponse is a wrapper around the C function soup_xmlrpc_build_method_response.
func XmlrpcBuildMethodResponse(value *gobject.Value) string {
	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	retC := C.soup_xmlrpc_build_method_response(c_value)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_xmlrpc_extract_method_call : unsupported parameter ... : varargs

// Unsupported : soup_xmlrpc_extract_method_response : unsupported parameter ... : varargs

// XmlrpcParseMethodCall is a wrapper around the C function soup_xmlrpc_parse_method_call.
func XmlrpcParseMethodCall(methodCall string, length int32) (bool, string, *gobject.ValueArray) {
	c_method_call := C.CString(methodCall)
	defer C.free(unsafe.Pointer(c_method_call))

	c_length := (C.int)(length)

	var c_method_name *C.char

	var c_params *C.GValueArray

	retC := C.soup_xmlrpc_parse_method_call(c_method_call, c_length, &c_method_name, &c_params)
	retGo := retC == C.TRUE

	methodName := C.GoString(c_method_name)
	defer C.free(unsafe.Pointer(c_method_name))

	params := gobject.ValueArrayNewFromC(unsafe.Pointer(c_params))

	return retGo, methodName, params
}

// XmlrpcParseMethodResponse is a wrapper around the C function soup_xmlrpc_parse_method_response.
func XmlrpcParseMethodResponse(methodResponse string, length int32) (bool, *gobject.Value, error) {
	c_method_response := C.CString(methodResponse)
	defer C.free(unsafe.Pointer(c_method_response))

	c_length := (C.int)(length)

	var c_value C.GValue

	var cThrowableError *C.GError

	retC := C.soup_xmlrpc_parse_method_response(c_method_response, c_length, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value, goError
}

// Unsupported : soup_xmlrpc_request_new : unsupported parameter ... : varargs

// XmlrpcSetFault is a wrapper around the C function soup_xmlrpc_set_fault.
func XmlrpcSetFault(msg *Message, faultCode int32, faultFormat string, args ...interface{}) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_fault_code := (C.int)(faultCode)

	goFormattedString := fmt.Sprintf(faultFormat, args...)
	c_fault_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_fault_format))

	C._soup_xmlrpc_set_fault(c_msg, c_fault_code, c_fault_format)

	return
}

// Unsupported : soup_xmlrpc_set_response : unsupported parameter ... : varargs

// PasswordManager is a wrapper around the C record SoupPasswordManager.
type PasswordManager struct {
	native *C.SoupPasswordManager
}

func PasswordManagerNewFromC(u unsafe.Pointer) *PasswordManager {
	c := (*C.SoupPasswordManager)(u)
	if c == nil {
		return nil
	}

	g := &PasswordManager{native: c}

	return g
}

func (recv *PasswordManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PasswordManager with another PasswordManager, and returns true if they represent the same GObject.
func (recv *PasswordManager) Equals(other *PasswordManager) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : soup_password_manager_get_passwords_async : unsupported parameter callback : no type generator for PasswordManagerCallback (SoupPasswordManagerCallback) for param callback

// GetPasswordsSync is a wrapper around the C function soup_password_manager_get_passwords_sync.
func (recv *PasswordManager) GetPasswordsSync(msg *Message, auth *Auth, cancellable *gio.Cancellable) {
	c_msg := (*C.SoupMessage)(C.NULL)
	if msg != nil {
		c_msg = (*C.SoupMessage)(msg.ToC())
	}

	c_auth := (*C.SoupAuth)(C.NULL)
	if auth != nil {
		c_auth = (*C.SoupAuth)(auth.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	C.soup_password_manager_get_passwords_sync((*C.SoupPasswordManager)(recv.native), c_msg, c_auth, c_cancellable)

	return
}

// ProxyURIResolver is a wrapper around the C record SoupProxyURIResolver.
type ProxyURIResolver struct {
	native *C.SoupProxyURIResolver
}

func ProxyURIResolverNewFromC(u unsafe.Pointer) *ProxyURIResolver {
	c := (*C.SoupProxyURIResolver)(u)
	if c == nil {
		return nil
	}

	g := &ProxyURIResolver{native: c}

	return g
}

func (recv *ProxyURIResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyURIResolver with another ProxyURIResolver, and returns true if they represent the same GObject.
func (recv *ProxyURIResolver) Equals(other *ProxyURIResolver) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : soup_proxy_uri_resolver_get_proxy_uri_async : unsupported parameter callback : no type generator for ProxyURIResolverCallback (SoupProxyURIResolverCallback) for param callback

// GetProxyUriSync is a wrapper around the C function soup_proxy_uri_resolver_get_proxy_uri_sync.
func (recv *ProxyURIResolver) GetProxyUriSync(uri *URI, cancellable *gio.Cancellable) (uint32, *URI) {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_proxy_uri *C.SoupURI

	retC := C.soup_proxy_uri_resolver_get_proxy_uri_sync((*C.SoupProxyURIResolver)(recv.native), c_uri, c_cancellable, &c_proxy_uri)
	retGo := (uint32)(retC)

	proxyUri := URINewFromC(unsafe.Pointer(c_proxy_uri))

	return retGo, proxyUri
}

// SessionFeature is a wrapper around the C record SoupSessionFeature.
type SessionFeature struct {
	native *C.SoupSessionFeature
}

func SessionFeatureNewFromC(u unsafe.Pointer) *SessionFeature {
	c := (*C.SoupSessionFeature)(u)
	if c == nil {
		return nil
	}

	g := &SessionFeature{native: c}

	return g
}

func (recv *SessionFeature) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionFeature with another SessionFeature, and returns true if they represent the same GObject.
func (recv *SessionFeature) Equals(other *SessionFeature) bool {
	return other.ToC() == recv.ToC()
}

// AddFeature is a wrapper around the C function soup_session_feature_add_feature.
func (recv *SessionFeature) AddFeature(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.soup_session_feature_add_feature((*C.SoupSessionFeature)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Attach is a wrapper around the C function soup_session_feature_attach.
func (recv *SessionFeature) Attach(session *Session) {
	c_session := (*C.SoupSession)(C.NULL)
	if session != nil {
		c_session = (*C.SoupSession)(session.ToC())
	}

	C.soup_session_feature_attach((*C.SoupSessionFeature)(recv.native), c_session)

	return
}

// Detach is a wrapper around the C function soup_session_feature_detach.
func (recv *SessionFeature) Detach(session *Session) {
	c_session := (*C.SoupSession)(C.NULL)
	if session != nil {
		c_session = (*C.SoupSession)(session.ToC())
	}

	C.soup_session_feature_detach((*C.SoupSessionFeature)(recv.native), c_session)

	return
}

// HasFeature is a wrapper around the C function soup_session_feature_has_feature.
func (recv *SessionFeature) HasFeature(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.soup_session_feature_has_feature((*C.SoupSessionFeature)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveFeature is a wrapper around the C function soup_session_feature_remove_feature.
func (recv *SessionFeature) RemoveFeature(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.soup_session_feature_remove_feature((*C.SoupSessionFeature)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// AddressClass is a wrapper around the C record SoupAddressClass.
type AddressClass struct {
	native *C.SoupAddressClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func AddressClassNewFromC(u unsafe.Pointer) *AddressClass {
	c := (*C.SoupAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &AddressClass{native: c}

	return g
}

func (recv *AddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AddressClass with another AddressClass, and returns true if they represent the same GObject.
func (recv *AddressClass) Equals(other *AddressClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthClass is a wrapper around the C record SoupAuthClass.
type AuthClass struct {
	native *C.SoupAuthClass
	// parent_class : record
	SchemeName string
	Strength   uint32
	// no type for update
	// no type for get_protection_space
	// no type for authenticate
	// no type for is_authenticated
	// no type for get_authorization
	// no type for is_ready
	// no type for can_authenticate
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func AuthClassNewFromC(u unsafe.Pointer) *AuthClass {
	c := (*C.SoupAuthClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthClass{
		SchemeName: C.GoString(c.scheme_name),
		Strength:   (uint32)(c.strength),
		native:     c,
	}

	return g
}

func (recv *AuthClass) ToC() unsafe.Pointer {
	recv.native.scheme_name =
		C.CString(recv.SchemeName)
	recv.native.strength =
		(C.guint)(recv.Strength)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthClass with another AuthClass, and returns true if they represent the same GObject.
func (recv *AuthClass) Equals(other *AuthClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthDomainBasicClass is a wrapper around the C record SoupAuthDomainBasicClass.
type AuthDomainBasicClass struct {
	native *C.SoupAuthDomainBasicClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func AuthDomainBasicClassNewFromC(u unsafe.Pointer) *AuthDomainBasicClass {
	c := (*C.SoupAuthDomainBasicClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomainBasicClass{native: c}

	return g
}

func (recv *AuthDomainBasicClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomainBasicClass with another AuthDomainBasicClass, and returns true if they represent the same GObject.
func (recv *AuthDomainBasicClass) Equals(other *AuthDomainBasicClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthDomainClass is a wrapper around the C record SoupAuthDomainClass.
type AuthDomainClass struct {
	native *C.SoupAuthDomainClass
	// parent_class : record
	// no type for accepts
	// no type for challenge
	// no type for check_password
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func AuthDomainClassNewFromC(u unsafe.Pointer) *AuthDomainClass {
	c := (*C.SoupAuthDomainClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomainClass{native: c}

	return g
}

func (recv *AuthDomainClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomainClass with another AuthDomainClass, and returns true if they represent the same GObject.
func (recv *AuthDomainClass) Equals(other *AuthDomainClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthDomainDigestClass is a wrapper around the C record SoupAuthDomainDigestClass.
type AuthDomainDigestClass struct {
	native *C.SoupAuthDomainDigestClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func AuthDomainDigestClassNewFromC(u unsafe.Pointer) *AuthDomainDigestClass {
	c := (*C.SoupAuthDomainDigestClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthDomainDigestClass{native: c}

	return g
}

func (recv *AuthDomainDigestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthDomainDigestClass with another AuthDomainDigestClass, and returns true if they represent the same GObject.
func (recv *AuthDomainDigestClass) Equals(other *AuthDomainDigestClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthManagerClass is a wrapper around the C record SoupAuthManagerClass.
type AuthManagerClass struct {
	native *C.SoupAuthManagerClass
	// parent_class : record
	// no type for authenticate
}

func AuthManagerClassNewFromC(u unsafe.Pointer) *AuthManagerClass {
	c := (*C.SoupAuthManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthManagerClass{native: c}

	return g
}

func (recv *AuthManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthManagerClass with another AuthManagerClass, and returns true if they represent the same GObject.
func (recv *AuthManagerClass) Equals(other *AuthManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthManagerPrivate is a wrapper around the C record SoupAuthManagerPrivate.
type AuthManagerPrivate struct {
	native *C.SoupAuthManagerPrivate
}

func AuthManagerPrivateNewFromC(u unsafe.Pointer) *AuthManagerPrivate {
	c := (*C.SoupAuthManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AuthManagerPrivate{native: c}

	return g
}

func (recv *AuthManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthManagerPrivate with another AuthManagerPrivate, and returns true if they represent the same GObject.
func (recv *AuthManagerPrivate) Equals(other *AuthManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Buffer is a wrapper around the C record SoupBuffer.
type Buffer struct {
	native *C.SoupBuffer
	// data : gchar* with indirection level of 1
	Length uint64
}

func BufferNewFromC(u unsafe.Pointer) *Buffer {
	c := (*C.SoupBuffer)(u)
	if c == nil {
		return nil
	}

	g := &Buffer{
		Length: (uint64)(c.length),
		native: c,
	}

	return g
}

func (recv *Buffer) ToC() unsafe.Pointer {
	recv.native.length =
		(C.gsize)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Buffer with another Buffer, and returns true if they represent the same GObject.
func (recv *Buffer) Equals(other *Buffer) bool {
	return other.ToC() == recv.ToC()
}

// BufferNew is a wrapper around the C function soup_buffer_new.
func BufferNew(use MemoryUse, data []uint8) *Buffer {
	c_use := (C.SoupMemoryUse)(use)

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.soup_buffer_new(c_use, c_data, c_length)
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BufferNewTake is a wrapper around the C function soup_buffer_new_take.
func BufferNewTake(data []uint8) *Buffer {
	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.soup_buffer_new_take(c_data, c_length)
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_buffer_new_with_owner : unsupported parameter owner_dnotify : no type generator for GLib.DestroyNotify (GDestroyNotify) for param owner_dnotify

// Copy is a wrapper around the C function soup_buffer_copy.
func (recv *Buffer) Copy() *Buffer {
	retC := C.soup_buffer_copy((*C.SoupBuffer)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function soup_buffer_free.
func (recv *Buffer) Free() {
	C.soup_buffer_free((*C.SoupBuffer)(recv.native))

	return
}

// GetAsBytes is a wrapper around the C function soup_buffer_get_as_bytes.
func (recv *Buffer) GetAsBytes() *glib.Bytes {
	retC := C.soup_buffer_get_as_bytes((*C.SoupBuffer)(recv.native))
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : soup_buffer_get_data : unsupported parameter data : output array param data

// GetOwner is a wrapper around the C function soup_buffer_get_owner.
func (recv *Buffer) GetOwner() uintptr {
	retC := C.soup_buffer_get_owner((*C.SoupBuffer)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// NewSubbuffer is a wrapper around the C function soup_buffer_new_subbuffer.
func (recv *Buffer) NewSubbuffer(offset uint64, length uint64) *Buffer {
	c_offset := (C.gsize)(offset)

	c_length := (C.gsize)(length)

	retC := C.soup_buffer_new_subbuffer((*C.SoupBuffer)(recv.native), c_offset, c_length)
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CacheClass is a wrapper around the C record SoupCacheClass.
type CacheClass struct {
	native *C.SoupCacheClass
	// parent_class : record
	// no type for get_cacheability
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
}

func CacheClassNewFromC(u unsafe.Pointer) *CacheClass {
	c := (*C.SoupCacheClass)(u)
	if c == nil {
		return nil
	}

	g := &CacheClass{native: c}

	return g
}

func (recv *CacheClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CacheClass with another CacheClass, and returns true if they represent the same GObject.
func (recv *CacheClass) Equals(other *CacheClass) bool {
	return other.ToC() == recv.ToC()
}

// CachePrivate is a wrapper around the C record SoupCachePrivate.
type CachePrivate struct {
	native *C.SoupCachePrivate
}

func CachePrivateNewFromC(u unsafe.Pointer) *CachePrivate {
	c := (*C.SoupCachePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CachePrivate{native: c}

	return g
}

func (recv *CachePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CachePrivate with another CachePrivate, and returns true if they represent the same GObject.
func (recv *CachePrivate) Equals(other *CachePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ClientContext is a wrapper around the C record SoupClientContext.
type ClientContext struct {
	native *C.SoupClientContext
}

func ClientContextNewFromC(u unsafe.Pointer) *ClientContext {
	c := (*C.SoupClientContext)(u)
	if c == nil {
		return nil
	}

	g := &ClientContext{native: c}

	return g
}

func (recv *ClientContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ClientContext with another ClientContext, and returns true if they represent the same GObject.
func (recv *ClientContext) Equals(other *ClientContext) bool {
	return other.ToC() == recv.ToC()
}

// GetAddress is a wrapper around the C function soup_client_context_get_address.
func (recv *ClientContext) GetAddress() *Address {
	retC := C.soup_client_context_get_address((*C.SoupClientContext)(recv.native))
	var retGo (*Address)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AddressNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetAuthDomain is a wrapper around the C function soup_client_context_get_auth_domain.
func (recv *ClientContext) GetAuthDomain() *AuthDomain {
	retC := C.soup_client_context_get_auth_domain((*C.SoupClientContext)(recv.native))
	var retGo (*AuthDomain)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AuthDomainNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetAuthUser is a wrapper around the C function soup_client_context_get_auth_user.
func (recv *ClientContext) GetAuthUser() string {
	retC := C.soup_client_context_get_auth_user((*C.SoupClientContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetGsocket is a wrapper around the C function soup_client_context_get_gsocket.
func (recv *ClientContext) GetGsocket() *gio.Socket {
	retC := C.soup_client_context_get_gsocket((*C.SoupClientContext)(recv.native))
	var retGo (*gio.Socket)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.SocketNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetHost is a wrapper around the C function soup_client_context_get_host.
func (recv *ClientContext) GetHost() string {
	retC := C.soup_client_context_get_host((*C.SoupClientContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function soup_client_context_get_local_address.
func (recv *ClientContext) GetLocalAddress() *gio.SocketAddress {
	retC := C.soup_client_context_get_local_address((*C.SoupClientContext)(recv.native))
	var retGo (*gio.SocketAddress)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.SocketAddressNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetRemoteAddress is a wrapper around the C function soup_client_context_get_remote_address.
func (recv *ClientContext) GetRemoteAddress() *gio.SocketAddress {
	retC := C.soup_client_context_get_remote_address((*C.SoupClientContext)(recv.native))
	var retGo (*gio.SocketAddress)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.SocketAddressNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSocket is a wrapper around the C function soup_client_context_get_socket.
func (recv *ClientContext) GetSocket() *Socket {
	retC := C.soup_client_context_get_socket((*C.SoupClientContext)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StealConnection is a wrapper around the C function soup_client_context_steal_connection.
func (recv *ClientContext) StealConnection() *gio.IOStream {
	retC := C.soup_client_context_steal_connection((*C.SoupClientContext)(recv.native))
	retGo := gio.IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Connection is a wrapper around the C record SoupConnection.
type Connection struct {
	native *C.SoupConnection
}

func ConnectionNewFromC(u unsafe.Pointer) *Connection {
	c := (*C.SoupConnection)(u)
	if c == nil {
		return nil
	}

	g := &Connection{native: c}

	return g
}

func (recv *Connection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Connection with another Connection, and returns true if they represent the same GObject.
func (recv *Connection) Equals(other *Connection) bool {
	return other.ToC() == recv.ToC()
}

// ContentDecoderClass is a wrapper around the C record SoupContentDecoderClass.
type ContentDecoderClass struct {
	native *C.SoupContentDecoderClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
	// no type for _libsoup_reserved5
}

func ContentDecoderClassNewFromC(u unsafe.Pointer) *ContentDecoderClass {
	c := (*C.SoupContentDecoderClass)(u)
	if c == nil {
		return nil
	}

	g := &ContentDecoderClass{native: c}

	return g
}

func (recv *ContentDecoderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentDecoderClass with another ContentDecoderClass, and returns true if they represent the same GObject.
func (recv *ContentDecoderClass) Equals(other *ContentDecoderClass) bool {
	return other.ToC() == recv.ToC()
}

// ContentDecoderPrivate is a wrapper around the C record SoupContentDecoderPrivate.
type ContentDecoderPrivate struct {
	native *C.SoupContentDecoderPrivate
}

func ContentDecoderPrivateNewFromC(u unsafe.Pointer) *ContentDecoderPrivate {
	c := (*C.SoupContentDecoderPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContentDecoderPrivate{native: c}

	return g
}

func (recv *ContentDecoderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentDecoderPrivate with another ContentDecoderPrivate, and returns true if they represent the same GObject.
func (recv *ContentDecoderPrivate) Equals(other *ContentDecoderPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContentSnifferClass is a wrapper around the C record SoupContentSnifferClass.
type ContentSnifferClass struct {
	native *C.SoupContentSnifferClass
	// parent_class : record
	// no type for sniff
	// no type for get_buffer_size
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
	// no type for _libsoup_reserved5
}

func ContentSnifferClassNewFromC(u unsafe.Pointer) *ContentSnifferClass {
	c := (*C.SoupContentSnifferClass)(u)
	if c == nil {
		return nil
	}

	g := &ContentSnifferClass{native: c}

	return g
}

func (recv *ContentSnifferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentSnifferClass with another ContentSnifferClass, and returns true if they represent the same GObject.
func (recv *ContentSnifferClass) Equals(other *ContentSnifferClass) bool {
	return other.ToC() == recv.ToC()
}

// ContentSnifferPrivate is a wrapper around the C record SoupContentSnifferPrivate.
type ContentSnifferPrivate struct {
	native *C.SoupContentSnifferPrivate
}

func ContentSnifferPrivateNewFromC(u unsafe.Pointer) *ContentSnifferPrivate {
	c := (*C.SoupContentSnifferPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContentSnifferPrivate{native: c}

	return g
}

func (recv *ContentSnifferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContentSnifferPrivate with another ContentSnifferPrivate, and returns true if they represent the same GObject.
func (recv *ContentSnifferPrivate) Equals(other *ContentSnifferPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Cookie is a wrapper around the C record SoupCookie.
type Cookie struct {
	native *C.SoupCookie
	Name   string
	Value  string
	Domain string
	Path   string
	// expires : record
	Secure   bool
	HttpOnly bool
}

func CookieNewFromC(u unsafe.Pointer) *Cookie {
	c := (*C.SoupCookie)(u)
	if c == nil {
		return nil
	}

	g := &Cookie{
		Domain:   C.GoString(c.domain),
		HttpOnly: c.http_only == C.TRUE,
		Name:     C.GoString(c.name),
		Path:     C.GoString(c.path),
		Secure:   c.secure == C.TRUE,
		Value:    C.GoString(c.value),
		native:   c,
	}

	return g
}

func (recv *Cookie) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.value =
		C.CString(recv.Value)
	recv.native.domain =
		C.CString(recv.Domain)
	recv.native.path =
		C.CString(recv.Path)
	recv.native.secure =
		boolToGboolean(recv.Secure)
	recv.native.http_only =
		boolToGboolean(recv.HttpOnly)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cookie with another Cookie, and returns true if they represent the same GObject.
func (recv *Cookie) Equals(other *Cookie) bool {
	return other.ToC() == recv.ToC()
}

// CookieNew is a wrapper around the C function soup_cookie_new.
func CookieNew(name string, value string, domain string, path string, maxAge int32) *Cookie {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_max_age := (C.int)(maxAge)

	retC := C.soup_cookie_new(c_name, c_value, c_domain, c_path, c_max_age)
	retGo := CookieNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CookieParse is a wrapper around the C function soup_cookie_parse.
func CookieParse(header string, origin *URI) *Cookie {
	c_header := C.CString(header)
	defer C.free(unsafe.Pointer(c_header))

	c_origin := (*C.SoupURI)(C.NULL)
	if origin != nil {
		c_origin = (*C.SoupURI)(origin.ToC())
	}

	retC := C.soup_cookie_parse(c_header, c_origin)
	var retGo (*Cookie)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CookieNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// AppliesToUri is a wrapper around the C function soup_cookie_applies_to_uri.
func (recv *Cookie) AppliesToUri(uri *URI) bool {
	c_uri := (*C.SoupURI)(C.NULL)
	if uri != nil {
		c_uri = (*C.SoupURI)(uri.ToC())
	}

	retC := C.soup_cookie_applies_to_uri((*C.SoupCookie)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Copy is a wrapper around the C function soup_cookie_copy.
func (recv *Cookie) Copy() *Cookie {
	retC := C.soup_cookie_copy((*C.SoupCookie)(recv.native))
	retGo := CookieNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DomainMatches is a wrapper around the C function soup_cookie_domain_matches.
func (recv *Cookie) DomainMatches(host string) bool {
	c_host := C.CString(host)
	defer C.free(unsafe.Pointer(c_host))

	retC := C.soup_cookie_domain_matches((*C.SoupCookie)(recv.native), c_host)
	retGo := retC == C.TRUE

	return retGo
}

// Equal is a wrapper around the C function soup_cookie_equal.
func (recv *Cookie) Equal(cookie2 *Cookie) bool {
	c_cookie2 := (*C.SoupCookie)(C.NULL)
	if cookie2 != nil {
		c_cookie2 = (*C.SoupCookie)(cookie2.ToC())
	}

	retC := C.soup_cookie_equal((*C.SoupCookie)(recv.native), c_cookie2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function soup_cookie_free.
func (recv *Cookie) Free() {
	C.soup_cookie_free((*C.SoupCookie)(recv.native))

	return
}

// GetDomain is a wrapper around the C function soup_cookie_get_domain.
func (recv *Cookie) GetDomain() string {
	retC := C.soup_cookie_get_domain((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetExpires is a wrapper around the C function soup_cookie_get_expires.
func (recv *Cookie) GetExpires() *Date {
	retC := C.soup_cookie_get_expires((*C.SoupCookie)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHttpOnly is a wrapper around the C function soup_cookie_get_http_only.
func (recv *Cookie) GetHttpOnly() bool {
	retC := C.soup_cookie_get_http_only((*C.SoupCookie)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function soup_cookie_get_name.
func (recv *Cookie) GetName() string {
	retC := C.soup_cookie_get_name((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPath is a wrapper around the C function soup_cookie_get_path.
func (recv *Cookie) GetPath() string {
	retC := C.soup_cookie_get_path((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSecure is a wrapper around the C function soup_cookie_get_secure.
func (recv *Cookie) GetSecure() bool {
	retC := C.soup_cookie_get_secure((*C.SoupCookie)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValue is a wrapper around the C function soup_cookie_get_value.
func (recv *Cookie) GetValue() string {
	retC := C.soup_cookie_get_value((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetDomain is a wrapper around the C function soup_cookie_set_domain.
func (recv *Cookie) SetDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.soup_cookie_set_domain((*C.SoupCookie)(recv.native), c_domain)

	return
}

// SetExpires is a wrapper around the C function soup_cookie_set_expires.
func (recv *Cookie) SetExpires(expires *Date) {
	c_expires := (*C.SoupDate)(C.NULL)
	if expires != nil {
		c_expires = (*C.SoupDate)(expires.ToC())
	}

	C.soup_cookie_set_expires((*C.SoupCookie)(recv.native), c_expires)

	return
}

// SetHttpOnly is a wrapper around the C function soup_cookie_set_http_only.
func (recv *Cookie) SetHttpOnly(httpOnly bool) {
	c_http_only :=
		boolToGboolean(httpOnly)

	C.soup_cookie_set_http_only((*C.SoupCookie)(recv.native), c_http_only)

	return
}

// SetMaxAge is a wrapper around the C function soup_cookie_set_max_age.
func (recv *Cookie) SetMaxAge(maxAge int32) {
	c_max_age := (C.int)(maxAge)

	C.soup_cookie_set_max_age((*C.SoupCookie)(recv.native), c_max_age)

	return
}

// SetName is a wrapper around the C function soup_cookie_set_name.
func (recv *Cookie) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.soup_cookie_set_name((*C.SoupCookie)(recv.native), c_name)

	return
}

// SetPath is a wrapper around the C function soup_cookie_set_path.
func (recv *Cookie) SetPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.soup_cookie_set_path((*C.SoupCookie)(recv.native), c_path)

	return
}

// SetSecure is a wrapper around the C function soup_cookie_set_secure.
func (recv *Cookie) SetSecure(secure bool) {
	c_secure :=
		boolToGboolean(secure)

	C.soup_cookie_set_secure((*C.SoupCookie)(recv.native), c_secure)

	return
}

// SetValue is a wrapper around the C function soup_cookie_set_value.
func (recv *Cookie) SetValue(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.soup_cookie_set_value((*C.SoupCookie)(recv.native), c_value)

	return
}

// ToCookieHeader is a wrapper around the C function soup_cookie_to_cookie_header.
func (recv *Cookie) ToCookieHeader() string {
	retC := C.soup_cookie_to_cookie_header((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToSetCookieHeader is a wrapper around the C function soup_cookie_to_set_cookie_header.
func (recv *Cookie) ToSetCookieHeader() string {
	retC := C.soup_cookie_to_set_cookie_header((*C.SoupCookie)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// CookieJarClass is a wrapper around the C record SoupCookieJarClass.
type CookieJarClass struct {
	native *C.SoupCookieJarClass
	// parent_class : record
	// no type for save
	// no type for is_persistent
	// no type for changed
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
}

func CookieJarClassNewFromC(u unsafe.Pointer) *CookieJarClass {
	c := (*C.SoupCookieJarClass)(u)
	if c == nil {
		return nil
	}

	g := &CookieJarClass{native: c}

	return g
}

func (recv *CookieJarClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJarClass with another CookieJarClass, and returns true if they represent the same GObject.
func (recv *CookieJarClass) Equals(other *CookieJarClass) bool {
	return other.ToC() == recv.ToC()
}

// CookieJarDBClass is a wrapper around the C record SoupCookieJarDBClass.
type CookieJarDBClass struct {
	native *C.SoupCookieJarDBClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func CookieJarDBClassNewFromC(u unsafe.Pointer) *CookieJarDBClass {
	c := (*C.SoupCookieJarDBClass)(u)
	if c == nil {
		return nil
	}

	g := &CookieJarDBClass{native: c}

	return g
}

func (recv *CookieJarDBClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJarDBClass with another CookieJarDBClass, and returns true if they represent the same GObject.
func (recv *CookieJarDBClass) Equals(other *CookieJarDBClass) bool {
	return other.ToC() == recv.ToC()
}

// CookieJarTextClass is a wrapper around the C record SoupCookieJarTextClass.
type CookieJarTextClass struct {
	native *C.SoupCookieJarTextClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func CookieJarTextClassNewFromC(u unsafe.Pointer) *CookieJarTextClass {
	c := (*C.SoupCookieJarTextClass)(u)
	if c == nil {
		return nil
	}

	g := &CookieJarTextClass{native: c}

	return g
}

func (recv *CookieJarTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieJarTextClass with another CookieJarTextClass, and returns true if they represent the same GObject.
func (recv *CookieJarTextClass) Equals(other *CookieJarTextClass) bool {
	return other.ToC() == recv.ToC()
}

// Date is a wrapper around the C record SoupDate.
type Date struct {
	native *C.SoupDate
	Year   int32
	Month  int32
	Day    int32
	Hour   int32
	Minute int32
	Second int32
	Utc    bool
	Offset int32
}

func DateNewFromC(u unsafe.Pointer) *Date {
	c := (*C.SoupDate)(u)
	if c == nil {
		return nil
	}

	g := &Date{
		Day:    (int32)(c.day),
		Hour:   (int32)(c.hour),
		Minute: (int32)(c.minute),
		Month:  (int32)(c.month),
		Offset: (int32)(c.offset),
		Second: (int32)(c.second),
		Utc:    c.utc == C.TRUE,
		Year:   (int32)(c.year),
		native: c,
	}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {
	recv.native.year =
		(C.int)(recv.Year)
	recv.native.month =
		(C.int)(recv.Month)
	recv.native.day =
		(C.int)(recv.Day)
	recv.native.hour =
		(C.int)(recv.Hour)
	recv.native.minute =
		(C.int)(recv.Minute)
	recv.native.second =
		(C.int)(recv.Second)
	recv.native.utc =
		boolToGboolean(recv.Utc)
	recv.native.offset =
		(C.int)(recv.Offset)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Date with another Date, and returns true if they represent the same GObject.
func (recv *Date) Equals(other *Date) bool {
	return other.ToC() == recv.ToC()
}

// DateNew is a wrapper around the C function soup_date_new.
func DateNew(year int32, month int32, day int32, hour int32, minute int32, second int32) *Date {
	c_year := (C.int)(year)

	c_month := (C.int)(month)

	c_day := (C.int)(day)

	c_hour := (C.int)(hour)

	c_minute := (C.int)(minute)

	c_second := (C.int)(second)

	retC := C.soup_date_new(c_year, c_month, c_day, c_hour, c_minute, c_second)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewFromNow is a wrapper around the C function soup_date_new_from_now.
func DateNewFromNow(offsetSeconds int32) *Date {
	c_offset_seconds := (C.int)(offsetSeconds)

	retC := C.soup_date_new_from_now(c_offset_seconds)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateNewFromString is a wrapper around the C function soup_date_new_from_string.
func DateNewFromString(dateString string) *Date {
	c_date_string := C.CString(dateString)
	defer C.free(unsafe.Pointer(c_date_string))

	retC := C.soup_date_new_from_string(c_date_string)
	var retGo (*Date)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DateNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DateNewFromTimeT is a wrapper around the C function soup_date_new_from_time_t.
func DateNewFromTimeT(when int64) *Date {
	c_when := (C.time_t)(when)

	retC := C.soup_date_new_from_time_t(c_when)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function soup_date_copy.
func (recv *Date) Copy() *Date {
	retC := C.soup_date_copy((*C.SoupDate)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function soup_date_free.
func (recv *Date) Free() {
	C.soup_date_free((*C.SoupDate)(recv.native))

	return
}

// GetDay is a wrapper around the C function soup_date_get_day.
func (recv *Date) GetDay() int32 {
	retC := C.soup_date_get_day((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHour is a wrapper around the C function soup_date_get_hour.
func (recv *Date) GetHour() int32 {
	retC := C.soup_date_get_hour((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinute is a wrapper around the C function soup_date_get_minute.
func (recv *Date) GetMinute() int32 {
	retC := C.soup_date_get_minute((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMonth is a wrapper around the C function soup_date_get_month.
func (recv *Date) GetMonth() int32 {
	retC := C.soup_date_get_month((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetOffset is a wrapper around the C function soup_date_get_offset.
func (recv *Date) GetOffset() int32 {
	retC := C.soup_date_get_offset((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSecond is a wrapper around the C function soup_date_get_second.
func (recv *Date) GetSecond() int32 {
	retC := C.soup_date_get_second((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUtc is a wrapper around the C function soup_date_get_utc.
func (recv *Date) GetUtc() int32 {
	retC := C.soup_date_get_utc((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetYear is a wrapper around the C function soup_date_get_year.
func (recv *Date) GetYear() int32 {
	retC := C.soup_date_get_year((*C.SoupDate)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsPast is a wrapper around the C function soup_date_is_past.
func (recv *Date) IsPast() bool {
	retC := C.soup_date_is_past((*C.SoupDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function soup_date_to_string.
func (recv *Date) ToString(format DateFormat) string {
	c_format := (C.SoupDateFormat)(format)

	retC := C.soup_date_to_string((*C.SoupDate)(recv.native), c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ToTimeT is a wrapper around the C function soup_date_to_time_t.
func (recv *Date) ToTimeT() int64 {
	retC := C.soup_date_to_time_t((*C.SoupDate)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// ToTimeval is a wrapper around the C function soup_date_to_timeval.
func (recv *Date) ToTimeval() *glib.TimeVal {
	var c_time C.GTimeVal

	C.soup_date_to_timeval((*C.SoupDate)(recv.native), &c_time)

	time := glib.TimeValNewFromC(unsafe.Pointer(&c_time))

	return time
}

// LoggerClass is a wrapper around the C record SoupLoggerClass.
type LoggerClass struct {
	native *C.SoupLoggerClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func LoggerClassNewFromC(u unsafe.Pointer) *LoggerClass {
	c := (*C.SoupLoggerClass)(u)
	if c == nil {
		return nil
	}

	g := &LoggerClass{native: c}

	return g
}

func (recv *LoggerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LoggerClass with another LoggerClass, and returns true if they represent the same GObject.
func (recv *LoggerClass) Equals(other *LoggerClass) bool {
	return other.ToC() == recv.ToC()
}

// MessageBody is a wrapper around the C record SoupMessageBody.
type MessageBody struct {
	native *C.SoupMessageBody
	Data   string
	Length int64
}

func MessageBodyNewFromC(u unsafe.Pointer) *MessageBody {
	c := (*C.SoupMessageBody)(u)
	if c == nil {
		return nil
	}

	g := &MessageBody{
		Data:   C.GoString(c.data),
		Length: (int64)(c.length),
		native: c,
	}

	return g
}

func (recv *MessageBody) ToC() unsafe.Pointer {
	recv.native.data =
		C.CString(recv.Data)
	recv.native.length =
		(C.goffset)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageBody with another MessageBody, and returns true if they represent the same GObject.
func (recv *MessageBody) Equals(other *MessageBody) bool {
	return other.ToC() == recv.ToC()
}

// MessageBodyNew is a wrapper around the C function soup_message_body_new.
func MessageBodyNew() *MessageBody {
	retC := C.soup_message_body_new()
	retGo := MessageBodyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function soup_message_body_append.
func (recv *MessageBody) Append(use MemoryUse, data []uint8) {
	c_use := (C.SoupMemoryUse)(use)

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gconstpointer)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	C.soup_message_body_append((*C.SoupMessageBody)(recv.native), c_use, c_data, c_length)

	return
}

// AppendBuffer is a wrapper around the C function soup_message_body_append_buffer.
func (recv *MessageBody) AppendBuffer(buffer *Buffer) {
	c_buffer := (*C.SoupBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.SoupBuffer)(buffer.ToC())
	}

	C.soup_message_body_append_buffer((*C.SoupMessageBody)(recv.native), c_buffer)

	return
}

// AppendTake is a wrapper around the C function soup_message_body_append_take.
func (recv *MessageBody) AppendTake(data []uint8) {
	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	C.soup_message_body_append_take((*C.SoupMessageBody)(recv.native), c_data, c_length)

	return
}

// Complete is a wrapper around the C function soup_message_body_complete.
func (recv *MessageBody) Complete() {
	C.soup_message_body_complete((*C.SoupMessageBody)(recv.native))

	return
}

// Flatten is a wrapper around the C function soup_message_body_flatten.
func (recv *MessageBody) Flatten() *Buffer {
	retC := C.soup_message_body_flatten((*C.SoupMessageBody)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function soup_message_body_free.
func (recv *MessageBody) Free() {
	C.soup_message_body_free((*C.SoupMessageBody)(recv.native))

	return
}

// GetAccumulate is a wrapper around the C function soup_message_body_get_accumulate.
func (recv *MessageBody) GetAccumulate() bool {
	retC := C.soup_message_body_get_accumulate((*C.SoupMessageBody)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetChunk is a wrapper around the C function soup_message_body_get_chunk.
func (recv *MessageBody) GetChunk(offset int64) *Buffer {
	c_offset := (C.goffset)(offset)

	retC := C.soup_message_body_get_chunk((*C.SoupMessageBody)(recv.native), c_offset)
	var retGo (*Buffer)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BufferNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GotChunk is a wrapper around the C function soup_message_body_got_chunk.
func (recv *MessageBody) GotChunk(chunk *Buffer) {
	c_chunk := (*C.SoupBuffer)(C.NULL)
	if chunk != nil {
		c_chunk = (*C.SoupBuffer)(chunk.ToC())
	}

	C.soup_message_body_got_chunk((*C.SoupMessageBody)(recv.native), c_chunk)

	return
}

// SetAccumulate is a wrapper around the C function soup_message_body_set_accumulate.
func (recv *MessageBody) SetAccumulate(accumulate bool) {
	c_accumulate :=
		boolToGboolean(accumulate)

	C.soup_message_body_set_accumulate((*C.SoupMessageBody)(recv.native), c_accumulate)

	return
}

// Truncate is a wrapper around the C function soup_message_body_truncate.
func (recv *MessageBody) Truncate() {
	C.soup_message_body_truncate((*C.SoupMessageBody)(recv.native))

	return
}

// WroteChunk is a wrapper around the C function soup_message_body_wrote_chunk.
func (recv *MessageBody) WroteChunk(chunk *Buffer) {
	c_chunk := (*C.SoupBuffer)(C.NULL)
	if chunk != nil {
		c_chunk = (*C.SoupBuffer)(chunk.ToC())
	}

	C.soup_message_body_wrote_chunk((*C.SoupMessageBody)(recv.native), c_chunk)

	return
}

// MessageClass is a wrapper around the C record SoupMessageClass.
type MessageClass struct {
	native *C.SoupMessageClass
	// parent_class : record
	// no type for wrote_informational
	// no type for wrote_headers
	// no type for wrote_chunk
	// no type for wrote_body
	// no type for got_informational
	// no type for got_headers
	// no type for got_chunk
	// no type for got_body
	// no type for restarted
	// no type for finished
	// no type for starting
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
}

func MessageClassNewFromC(u unsafe.Pointer) *MessageClass {
	c := (*C.SoupMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &MessageClass{native: c}

	return g
}

func (recv *MessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageClass with another MessageClass, and returns true if they represent the same GObject.
func (recv *MessageClass) Equals(other *MessageClass) bool {
	return other.ToC() == recv.ToC()
}

// MessageHeaders is a wrapper around the C record SoupMessageHeaders.
type MessageHeaders struct {
	native *C.SoupMessageHeaders
}

func MessageHeadersNewFromC(u unsafe.Pointer) *MessageHeaders {
	c := (*C.SoupMessageHeaders)(u)
	if c == nil {
		return nil
	}

	g := &MessageHeaders{native: c}

	return g
}

func (recv *MessageHeaders) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageHeaders with another MessageHeaders, and returns true if they represent the same GObject.
func (recv *MessageHeaders) Equals(other *MessageHeaders) bool {
	return other.ToC() == recv.ToC()
}

// MessageHeadersNew is a wrapper around the C function soup_message_headers_new.
func MessageHeadersNew(type_ MessageHeadersType) *MessageHeaders {
	c_type := (C.SoupMessageHeadersType)(type_)

	retC := C.soup_message_headers_new(c_type)
	retGo := MessageHeadersNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append is a wrapper around the C function soup_message_headers_append.
func (recv *MessageHeaders) Append(name string, value string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.soup_message_headers_append((*C.SoupMessageHeaders)(recv.native), c_name, c_value)

	return
}

// CleanConnectionHeaders is a wrapper around the C function soup_message_headers_clean_connection_headers.
func (recv *MessageHeaders) CleanConnectionHeaders() {
	C.soup_message_headers_clean_connection_headers((*C.SoupMessageHeaders)(recv.native))

	return
}

// Clear is a wrapper around the C function soup_message_headers_clear.
func (recv *MessageHeaders) Clear() {
	C.soup_message_headers_clear((*C.SoupMessageHeaders)(recv.native))

	return
}

// Unsupported : soup_message_headers_foreach : unsupported parameter func : no type generator for MessageHeadersForeachFunc (SoupMessageHeadersForeachFunc) for param func

// Free is a wrapper around the C function soup_message_headers_free.
func (recv *MessageHeaders) Free() {
	C.soup_message_headers_free((*C.SoupMessageHeaders)(recv.native))

	return
}

// FreeRanges is a wrapper around the C function soup_message_headers_free_ranges.
func (recv *MessageHeaders) FreeRanges(ranges *Range) {
	c_ranges := (*C.SoupRange)(C.NULL)
	if ranges != nil {
		c_ranges = (*C.SoupRange)(ranges.ToC())
	}

	C.soup_message_headers_free_ranges((*C.SoupMessageHeaders)(recv.native), c_ranges)

	return
}

// Get is a wrapper around the C function soup_message_headers_get.
func (recv *MessageHeaders) Get(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.soup_message_headers_get((*C.SoupMessageHeaders)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// GetContentDisposition is a wrapper around the C function soup_message_headers_get_content_disposition.
func (recv *MessageHeaders) GetContentDisposition() (bool, string, *glib.HashTable) {
	var c_disposition *C.char

	var c_params *C.GHashTable

	retC := C.soup_message_headers_get_content_disposition((*C.SoupMessageHeaders)(recv.native), &c_disposition, &c_params)
	retGo := retC == C.TRUE

	disposition := C.GoString(c_disposition)
	defer C.free(unsafe.Pointer(c_disposition))

	params := glib.HashTableNewFromC(unsafe.Pointer(c_params))

	return retGo, disposition, params
}

// GetContentLength is a wrapper around the C function soup_message_headers_get_content_length.
func (recv *MessageHeaders) GetContentLength() int64 {
	retC := C.soup_message_headers_get_content_length((*C.SoupMessageHeaders)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetContentRange is a wrapper around the C function soup_message_headers_get_content_range.
func (recv *MessageHeaders) GetContentRange() (bool, int64, int64, int64) {
	var c_start C.goffset

	var c_end C.goffset

	var c_total_length C.goffset

	retC := C.soup_message_headers_get_content_range((*C.SoupMessageHeaders)(recv.native), &c_start, &c_end, &c_total_length)
	retGo := retC == C.TRUE

	start := (int64)(c_start)

	end := (int64)(c_end)

	totalLength := (int64)(c_total_length)

	return retGo, start, end, totalLength
}

// GetContentType is a wrapper around the C function soup_message_headers_get_content_type.
func (recv *MessageHeaders) GetContentType() (string, *glib.HashTable) {
	var c_params *C.GHashTable

	retC := C.soup_message_headers_get_content_type((*C.SoupMessageHeaders)(recv.native), &c_params)
	retGo := C.GoString(retC)

	params := glib.HashTableNewFromC(unsafe.Pointer(c_params))

	return retGo, params
}

// GetEncoding is a wrapper around the C function soup_message_headers_get_encoding.
func (recv *MessageHeaders) GetEncoding() Encoding {
	retC := C.soup_message_headers_get_encoding((*C.SoupMessageHeaders)(recv.native))
	retGo := (Encoding)(retC)

	return retGo
}

// GetExpectations is a wrapper around the C function soup_message_headers_get_expectations.
func (recv *MessageHeaders) GetExpectations() Expectation {
	retC := C.soup_message_headers_get_expectations((*C.SoupMessageHeaders)(recv.native))
	retGo := (Expectation)(retC)

	return retGo
}

// GetHeadersType is a wrapper around the C function soup_message_headers_get_headers_type.
func (recv *MessageHeaders) GetHeadersType() MessageHeadersType {
	retC := C.soup_message_headers_get_headers_type((*C.SoupMessageHeaders)(recv.native))
	retGo := (MessageHeadersType)(retC)

	return retGo
}

// GetList is a wrapper around the C function soup_message_headers_get_list.
func (recv *MessageHeaders) GetList(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.soup_message_headers_get_list((*C.SoupMessageHeaders)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// GetOne is a wrapper around the C function soup_message_headers_get_one.
func (recv *MessageHeaders) GetOne(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.soup_message_headers_get_one((*C.SoupMessageHeaders)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : soup_message_headers_get_ranges : unsupported parameter ranges : output array param ranges

// HeaderContains is a wrapper around the C function soup_message_headers_header_contains.
func (recv *MessageHeaders) HeaderContains(name string, token string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_token := C.CString(token)
	defer C.free(unsafe.Pointer(c_token))

	retC := C.soup_message_headers_header_contains((*C.SoupMessageHeaders)(recv.native), c_name, c_token)
	retGo := retC == C.TRUE

	return retGo
}

// HeaderEquals is a wrapper around the C function soup_message_headers_header_equals.
func (recv *MessageHeaders) HeaderEquals(name string, value string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.soup_message_headers_header_equals((*C.SoupMessageHeaders)(recv.native), c_name, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Remove is a wrapper around the C function soup_message_headers_remove.
func (recv *MessageHeaders) Remove(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.soup_message_headers_remove((*C.SoupMessageHeaders)(recv.native), c_name)

	return
}

// Replace is a wrapper around the C function soup_message_headers_replace.
func (recv *MessageHeaders) Replace(name string, value string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.soup_message_headers_replace((*C.SoupMessageHeaders)(recv.native), c_name, c_value)

	return
}

// SetContentDisposition is a wrapper around the C function soup_message_headers_set_content_disposition.
func (recv *MessageHeaders) SetContentDisposition(disposition string, params *glib.HashTable) {
	c_disposition := C.CString(disposition)
	defer C.free(unsafe.Pointer(c_disposition))

	c_params := (*C.GHashTable)(C.NULL)
	if params != nil {
		c_params = (*C.GHashTable)(params.ToC())
	}

	C.soup_message_headers_set_content_disposition((*C.SoupMessageHeaders)(recv.native), c_disposition, c_params)

	return
}

// SetContentLength is a wrapper around the C function soup_message_headers_set_content_length.
func (recv *MessageHeaders) SetContentLength(contentLength int64) {
	c_content_length := (C.goffset)(contentLength)

	C.soup_message_headers_set_content_length((*C.SoupMessageHeaders)(recv.native), c_content_length)

	return
}

// SetContentRange is a wrapper around the C function soup_message_headers_set_content_range.
func (recv *MessageHeaders) SetContentRange(start int64, end int64, totalLength int64) {
	c_start := (C.goffset)(start)

	c_end := (C.goffset)(end)

	c_total_length := (C.goffset)(totalLength)

	C.soup_message_headers_set_content_range((*C.SoupMessageHeaders)(recv.native), c_start, c_end, c_total_length)

	return
}

// SetContentType is a wrapper around the C function soup_message_headers_set_content_type.
func (recv *MessageHeaders) SetContentType(contentType string, params *glib.HashTable) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_params := (*C.GHashTable)(C.NULL)
	if params != nil {
		c_params = (*C.GHashTable)(params.ToC())
	}

	C.soup_message_headers_set_content_type((*C.SoupMessageHeaders)(recv.native), c_content_type, c_params)

	return
}

// SetEncoding is a wrapper around the C function soup_message_headers_set_encoding.
func (recv *MessageHeaders) SetEncoding(encoding Encoding) {
	c_encoding := (C.SoupEncoding)(encoding)

	C.soup_message_headers_set_encoding((*C.SoupMessageHeaders)(recv.native), c_encoding)

	return
}

// SetExpectations is a wrapper around the C function soup_message_headers_set_expectations.
func (recv *MessageHeaders) SetExpectations(expectations Expectation) {
	c_expectations := (C.SoupExpectation)(expectations)

	C.soup_message_headers_set_expectations((*C.SoupMessageHeaders)(recv.native), c_expectations)

	return
}

// SetRange is a wrapper around the C function soup_message_headers_set_range.
func (recv *MessageHeaders) SetRange(start int64, end int64) {
	c_start := (C.goffset)(start)

	c_end := (C.goffset)(end)

	C.soup_message_headers_set_range((*C.SoupMessageHeaders)(recv.native), c_start, c_end)

	return
}

// SetRanges is a wrapper around the C function soup_message_headers_set_ranges.
func (recv *MessageHeaders) SetRanges(ranges *Range, length int32) {
	c_ranges := (*C.SoupRange)(C.NULL)
	if ranges != nil {
		c_ranges = (*C.SoupRange)(ranges.ToC())
	}

	c_length := (C.int)(length)

	C.soup_message_headers_set_ranges((*C.SoupMessageHeaders)(recv.native), c_ranges, c_length)

	return
}

// MessageHeadersIter is a wrapper around the C record SoupMessageHeadersIter.
type MessageHeadersIter struct {
	native *C.SoupMessageHeadersIter
	// Private : dummy
}

func MessageHeadersIterNewFromC(u unsafe.Pointer) *MessageHeadersIter {
	c := (*C.SoupMessageHeadersIter)(u)
	if c == nil {
		return nil
	}

	g := &MessageHeadersIter{native: c}

	return g
}

func (recv *MessageHeadersIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageHeadersIter with another MessageHeadersIter, and returns true if they represent the same GObject.
func (recv *MessageHeadersIter) Equals(other *MessageHeadersIter) bool {
	return other.ToC() == recv.ToC()
}

// MessageHeadersIterInit is a wrapper around the C function soup_message_headers_iter_init.
func MessageHeadersIterInit(hdrs *MessageHeaders) *MessageHeadersIter {
	var c_iter C.SoupMessageHeadersIter

	c_hdrs := (*C.SoupMessageHeaders)(C.NULL)
	if hdrs != nil {
		c_hdrs = (*C.SoupMessageHeaders)(hdrs.ToC())
	}

	C.soup_message_headers_iter_init(&c_iter, c_hdrs)

	iter := MessageHeadersIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Next is a wrapper around the C function soup_message_headers_iter_next.
func (recv *MessageHeadersIter) Next() (bool, string, string) {
	var c_name *C.char

	var c_value *C.char

	retC := C.soup_message_headers_iter_next((*C.SoupMessageHeadersIter)(recv.native), &c_name, &c_value)
	retGo := retC == C.TRUE

	name := C.GoString(c_name)

	value := C.GoString(c_value)

	return retGo, name, value
}

// MessageQueue is a wrapper around the C record SoupMessageQueue.
type MessageQueue struct {
	native *C.SoupMessageQueue
}

func MessageQueueNewFromC(u unsafe.Pointer) *MessageQueue {
	c := (*C.SoupMessageQueue)(u)
	if c == nil {
		return nil
	}

	g := &MessageQueue{native: c}

	return g
}

func (recv *MessageQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageQueue with another MessageQueue, and returns true if they represent the same GObject.
func (recv *MessageQueue) Equals(other *MessageQueue) bool {
	return other.ToC() == recv.ToC()
}

// MessageQueueItem is a wrapper around the C record SoupMessageQueueItem.
type MessageQueueItem struct {
	native *C.SoupMessageQueueItem
}

func MessageQueueItemNewFromC(u unsafe.Pointer) *MessageQueueItem {
	c := (*C.SoupMessageQueueItem)(u)
	if c == nil {
		return nil
	}

	g := &MessageQueueItem{native: c}

	return g
}

func (recv *MessageQueueItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MessageQueueItem with another MessageQueueItem, and returns true if they represent the same GObject.
func (recv *MessageQueueItem) Equals(other *MessageQueueItem) bool {
	return other.ToC() == recv.ToC()
}

// Multipart is a wrapper around the C record SoupMultipart.
type Multipart struct {
	native *C.SoupMultipart
}

func MultipartNewFromC(u unsafe.Pointer) *Multipart {
	c := (*C.SoupMultipart)(u)
	if c == nil {
		return nil
	}

	g := &Multipart{native: c}

	return g
}

func (recv *Multipart) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Multipart with another Multipart, and returns true if they represent the same GObject.
func (recv *Multipart) Equals(other *Multipart) bool {
	return other.ToC() == recv.ToC()
}

// MultipartNew is a wrapper around the C function soup_multipart_new.
func MultipartNew(mimeType string) *Multipart {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.soup_multipart_new(c_mime_type)
	retGo := MultipartNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MultipartNewFromMessage is a wrapper around the C function soup_multipart_new_from_message.
func MultipartNewFromMessage(headers *MessageHeaders, body *MessageBody) *Multipart {
	c_headers := (*C.SoupMessageHeaders)(C.NULL)
	if headers != nil {
		c_headers = (*C.SoupMessageHeaders)(headers.ToC())
	}

	c_body := (*C.SoupMessageBody)(C.NULL)
	if body != nil {
		c_body = (*C.SoupMessageBody)(body.ToC())
	}

	retC := C.soup_multipart_new_from_message(c_headers, c_body)
	var retGo (*Multipart)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MultipartNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// AppendFormFile is a wrapper around the C function soup_multipart_append_form_file.
func (recv *Multipart) AppendFormFile(controlName string, filename string, contentType string, body *Buffer) {
	c_control_name := C.CString(controlName)
	defer C.free(unsafe.Pointer(c_control_name))

	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_body := (*C.SoupBuffer)(C.NULL)
	if body != nil {
		c_body = (*C.SoupBuffer)(body.ToC())
	}

	C.soup_multipart_append_form_file((*C.SoupMultipart)(recv.native), c_control_name, c_filename, c_content_type, c_body)

	return
}

// AppendFormString is a wrapper around the C function soup_multipart_append_form_string.
func (recv *Multipart) AppendFormString(controlName string, data string) {
	c_control_name := C.CString(controlName)
	defer C.free(unsafe.Pointer(c_control_name))

	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	C.soup_multipart_append_form_string((*C.SoupMultipart)(recv.native), c_control_name, c_data)

	return
}

// AppendPart is a wrapper around the C function soup_multipart_append_part.
func (recv *Multipart) AppendPart(headers *MessageHeaders, body *Buffer) {
	c_headers := (*C.SoupMessageHeaders)(C.NULL)
	if headers != nil {
		c_headers = (*C.SoupMessageHeaders)(headers.ToC())
	}

	c_body := (*C.SoupBuffer)(C.NULL)
	if body != nil {
		c_body = (*C.SoupBuffer)(body.ToC())
	}

	C.soup_multipart_append_part((*C.SoupMultipart)(recv.native), c_headers, c_body)

	return
}

// Free is a wrapper around the C function soup_multipart_free.
func (recv *Multipart) Free() {
	C.soup_multipart_free((*C.SoupMultipart)(recv.native))

	return
}

// GetLength is a wrapper around the C function soup_multipart_get_length.
func (recv *Multipart) GetLength() int32 {
	retC := C.soup_multipart_get_length((*C.SoupMultipart)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPart is a wrapper around the C function soup_multipart_get_part.
func (recv *Multipart) GetPart(part int32) (bool, *MessageHeaders, *Buffer) {
	c_part := (C.int)(part)

	var c_headers *C.SoupMessageHeaders

	var c_body *C.SoupBuffer

	retC := C.soup_multipart_get_part((*C.SoupMultipart)(recv.native), c_part, &c_headers, &c_body)
	retGo := retC == C.TRUE

	headers := MessageHeadersNewFromC(unsafe.Pointer(c_headers))

	body := BufferNewFromC(unsafe.Pointer(c_body))

	return retGo, headers, body
}

// ToMessage is a wrapper around the C function soup_multipart_to_message.
func (recv *Multipart) ToMessage(destHeaders *MessageHeaders, destBody *MessageBody) {
	c_dest_headers := (*C.SoupMessageHeaders)(C.NULL)
	if destHeaders != nil {
		c_dest_headers = (*C.SoupMessageHeaders)(destHeaders.ToC())
	}

	c_dest_body := (*C.SoupMessageBody)(C.NULL)
	if destBody != nil {
		c_dest_body = (*C.SoupMessageBody)(destBody.ToC())
	}

	C.soup_multipart_to_message((*C.SoupMultipart)(recv.native), c_dest_headers, c_dest_body)

	return
}

// MultipartInputStreamClass is a wrapper around the C record SoupMultipartInputStreamClass.
type MultipartInputStreamClass struct {
	native *C.SoupMultipartInputStreamClass
	// parent_class : record
}

func MultipartInputStreamClassNewFromC(u unsafe.Pointer) *MultipartInputStreamClass {
	c := (*C.SoupMultipartInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MultipartInputStreamClass{native: c}

	return g
}

func (recv *MultipartInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MultipartInputStreamClass with another MultipartInputStreamClass, and returns true if they represent the same GObject.
func (recv *MultipartInputStreamClass) Equals(other *MultipartInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// MultipartInputStreamPrivate is a wrapper around the C record SoupMultipartInputStreamPrivate.
type MultipartInputStreamPrivate struct {
	native *C.SoupMultipartInputStreamPrivate
}

func MultipartInputStreamPrivateNewFromC(u unsafe.Pointer) *MultipartInputStreamPrivate {
	c := (*C.SoupMultipartInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MultipartInputStreamPrivate{native: c}

	return g
}

func (recv *MultipartInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MultipartInputStreamPrivate with another MultipartInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *MultipartInputStreamPrivate) Equals(other *MultipartInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PasswordManagerInterface is a wrapper around the C record SoupPasswordManagerInterface.
type PasswordManagerInterface struct {
	native *C.SoupPasswordManagerInterface
	// base : record
	// no type for get_passwords_async
	// no type for get_passwords_sync
}

func PasswordManagerInterfaceNewFromC(u unsafe.Pointer) *PasswordManagerInterface {
	c := (*C.SoupPasswordManagerInterface)(u)
	if c == nil {
		return nil
	}

	g := &PasswordManagerInterface{native: c}

	return g
}

func (recv *PasswordManagerInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PasswordManagerInterface with another PasswordManagerInterface, and returns true if they represent the same GObject.
func (recv *PasswordManagerInterface) Equals(other *PasswordManagerInterface) bool {
	return other.ToC() == recv.ToC()
}

// ProxyResolverDefaultClass is a wrapper around the C record SoupProxyResolverDefaultClass.
type ProxyResolverDefaultClass struct {
	native *C.SoupProxyResolverDefaultClass
	// parent_class : record
}

func ProxyResolverDefaultClassNewFromC(u unsafe.Pointer) *ProxyResolverDefaultClass {
	c := (*C.SoupProxyResolverDefaultClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolverDefaultClass{native: c}

	return g
}

func (recv *ProxyResolverDefaultClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyResolverDefaultClass with another ProxyResolverDefaultClass, and returns true if they represent the same GObject.
func (recv *ProxyResolverDefaultClass) Equals(other *ProxyResolverDefaultClass) bool {
	return other.ToC() == recv.ToC()
}

// ProxyURIResolverInterface is a wrapper around the C record SoupProxyURIResolverInterface.
type ProxyURIResolverInterface struct {
	native *C.SoupProxyURIResolverInterface
	// base : record
	// no type for get_proxy_uri_async
	// no type for get_proxy_uri_sync
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func ProxyURIResolverInterfaceNewFromC(u unsafe.Pointer) *ProxyURIResolverInterface {
	c := (*C.SoupProxyURIResolverInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyURIResolverInterface{native: c}

	return g
}

func (recv *ProxyURIResolverInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyURIResolverInterface with another ProxyURIResolverInterface, and returns true if they represent the same GObject.
func (recv *ProxyURIResolverInterface) Equals(other *ProxyURIResolverInterface) bool {
	return other.ToC() == recv.ToC()
}

// Range is a wrapper around the C record SoupRange.
type Range struct {
	native *C.SoupRange
	Start  int64
	End    int64
}

func RangeNewFromC(u unsafe.Pointer) *Range {
	c := (*C.SoupRange)(u)
	if c == nil {
		return nil
	}

	g := &Range{
		End:    (int64)(c.end),
		Start:  (int64)(c.start),
		native: c,
	}

	return g
}

func (recv *Range) ToC() unsafe.Pointer {
	recv.native.start =
		(C.goffset)(recv.Start)
	recv.native.end =
		(C.goffset)(recv.End)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Range with another Range, and returns true if they represent the same GObject.
func (recv *Range) Equals(other *Range) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : SoupRequestClass

// RequestDataClass is a wrapper around the C record SoupRequestDataClass.
type RequestDataClass struct {
	native *C.SoupRequestDataClass
	// parent : record
}

func RequestDataClassNewFromC(u unsafe.Pointer) *RequestDataClass {
	c := (*C.SoupRequestDataClass)(u)
	if c == nil {
		return nil
	}

	g := &RequestDataClass{native: c}

	return g
}

func (recv *RequestDataClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestDataClass with another RequestDataClass, and returns true if they represent the same GObject.
func (recv *RequestDataClass) Equals(other *RequestDataClass) bool {
	return other.ToC() == recv.ToC()
}

// RequestDataPrivate is a wrapper around the C record SoupRequestDataPrivate.
type RequestDataPrivate struct {
	native *C.SoupRequestDataPrivate
}

func RequestDataPrivateNewFromC(u unsafe.Pointer) *RequestDataPrivate {
	c := (*C.SoupRequestDataPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RequestDataPrivate{native: c}

	return g
}

func (recv *RequestDataPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestDataPrivate with another RequestDataPrivate, and returns true if they represent the same GObject.
func (recv *RequestDataPrivate) Equals(other *RequestDataPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RequestFileClass is a wrapper around the C record SoupRequestFileClass.
type RequestFileClass struct {
	native *C.SoupRequestFileClass
	// parent : record
}

func RequestFileClassNewFromC(u unsafe.Pointer) *RequestFileClass {
	c := (*C.SoupRequestFileClass)(u)
	if c == nil {
		return nil
	}

	g := &RequestFileClass{native: c}

	return g
}

func (recv *RequestFileClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestFileClass with another RequestFileClass, and returns true if they represent the same GObject.
func (recv *RequestFileClass) Equals(other *RequestFileClass) bool {
	return other.ToC() == recv.ToC()
}

// RequestFilePrivate is a wrapper around the C record SoupRequestFilePrivate.
type RequestFilePrivate struct {
	native *C.SoupRequestFilePrivate
}

func RequestFilePrivateNewFromC(u unsafe.Pointer) *RequestFilePrivate {
	c := (*C.SoupRequestFilePrivate)(u)
	if c == nil {
		return nil
	}

	g := &RequestFilePrivate{native: c}

	return g
}

func (recv *RequestFilePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestFilePrivate with another RequestFilePrivate, and returns true if they represent the same GObject.
func (recv *RequestFilePrivate) Equals(other *RequestFilePrivate) bool {
	return other.ToC() == recv.ToC()
}

// RequestHTTPClass is a wrapper around the C record SoupRequestHTTPClass.
type RequestHTTPClass struct {
	native *C.SoupRequestHTTPClass
	// parent : record
}

func RequestHTTPClassNewFromC(u unsafe.Pointer) *RequestHTTPClass {
	c := (*C.SoupRequestHTTPClass)(u)
	if c == nil {
		return nil
	}

	g := &RequestHTTPClass{native: c}

	return g
}

func (recv *RequestHTTPClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestHTTPClass with another RequestHTTPClass, and returns true if they represent the same GObject.
func (recv *RequestHTTPClass) Equals(other *RequestHTTPClass) bool {
	return other.ToC() == recv.ToC()
}

// RequestHTTPPrivate is a wrapper around the C record SoupRequestHTTPPrivate.
type RequestHTTPPrivate struct {
	native *C.SoupRequestHTTPPrivate
}

func RequestHTTPPrivateNewFromC(u unsafe.Pointer) *RequestHTTPPrivate {
	c := (*C.SoupRequestHTTPPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RequestHTTPPrivate{native: c}

	return g
}

func (recv *RequestHTTPPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestHTTPPrivate with another RequestHTTPPrivate, and returns true if they represent the same GObject.
func (recv *RequestHTTPPrivate) Equals(other *RequestHTTPPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RequestPrivate is a wrapper around the C record SoupRequestPrivate.
type RequestPrivate struct {
	native *C.SoupRequestPrivate
}

func RequestPrivateNewFromC(u unsafe.Pointer) *RequestPrivate {
	c := (*C.SoupRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &RequestPrivate{native: c}

	return g
}

func (recv *RequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RequestPrivate with another RequestPrivate, and returns true if they represent the same GObject.
func (recv *RequestPrivate) Equals(other *RequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : SoupRequesterClass

// Blacklisted : SoupRequesterPrivate

// ServerClass is a wrapper around the C record SoupServerClass.
type ServerClass struct {
	native *C.SoupServerClass
	// parent_class : record
	// no type for request_started
	// no type for request_read
	// no type for request_finished
	// no type for request_aborted
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func ServerClassNewFromC(u unsafe.Pointer) *ServerClass {
	c := (*C.SoupServerClass)(u)
	if c == nil {
		return nil
	}

	g := &ServerClass{native: c}

	return g
}

func (recv *ServerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ServerClass with another ServerClass, and returns true if they represent the same GObject.
func (recv *ServerClass) Equals(other *ServerClass) bool {
	return other.ToC() == recv.ToC()
}

// SessionAsyncClass is a wrapper around the C record SoupSessionAsyncClass.
type SessionAsyncClass struct {
	native *C.SoupSessionAsyncClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func SessionAsyncClassNewFromC(u unsafe.Pointer) *SessionAsyncClass {
	c := (*C.SoupSessionAsyncClass)(u)
	if c == nil {
		return nil
	}

	g := &SessionAsyncClass{native: c}

	return g
}

func (recv *SessionAsyncClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionAsyncClass with another SessionAsyncClass, and returns true if they represent the same GObject.
func (recv *SessionAsyncClass) Equals(other *SessionAsyncClass) bool {
	return other.ToC() == recv.ToC()
}

// SessionClass is a wrapper around the C record SoupSessionClass.
type SessionClass struct {
	native *C.SoupSessionClass
	// parent_class : record
	// no type for request_started
	// no type for authenticate
	// no type for queue_message
	// no type for requeue_message
	// no type for send_message
	// no type for cancel_message
	// no type for auth_required
	// no type for flush_queue
	// no type for kick
	// no type for _libsoup_reserved4
}

func SessionClassNewFromC(u unsafe.Pointer) *SessionClass {
	c := (*C.SoupSessionClass)(u)
	if c == nil {
		return nil
	}

	g := &SessionClass{native: c}

	return g
}

func (recv *SessionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionClass with another SessionClass, and returns true if they represent the same GObject.
func (recv *SessionClass) Equals(other *SessionClass) bool {
	return other.ToC() == recv.ToC()
}

// SessionFeatureInterface is a wrapper around the C record SoupSessionFeatureInterface.
type SessionFeatureInterface struct {
	native *C.SoupSessionFeatureInterface
	// parent : record
	// no type for attach
	// no type for detach
	// no type for request_queued
	// no type for request_started
	// no type for request_unqueued
	// no type for add_feature
	// no type for remove_feature
	// no type for has_feature
}

func SessionFeatureInterfaceNewFromC(u unsafe.Pointer) *SessionFeatureInterface {
	c := (*C.SoupSessionFeatureInterface)(u)
	if c == nil {
		return nil
	}

	g := &SessionFeatureInterface{native: c}

	return g
}

func (recv *SessionFeatureInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionFeatureInterface with another SessionFeatureInterface, and returns true if they represent the same GObject.
func (recv *SessionFeatureInterface) Equals(other *SessionFeatureInterface) bool {
	return other.ToC() == recv.ToC()
}

// SessionSyncClass is a wrapper around the C record SoupSessionSyncClass.
type SessionSyncClass struct {
	native *C.SoupSessionSyncClass
	// parent_class : record
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func SessionSyncClassNewFromC(u unsafe.Pointer) *SessionSyncClass {
	c := (*C.SoupSessionSyncClass)(u)
	if c == nil {
		return nil
	}

	g := &SessionSyncClass{native: c}

	return g
}

func (recv *SessionSyncClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SessionSyncClass with another SessionSyncClass, and returns true if they represent the same GObject.
func (recv *SessionSyncClass) Equals(other *SessionSyncClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketClass is a wrapper around the C record SoupSocketClass.
type SocketClass struct {
	native *C.SoupSocketClass
	// parent_class : record
	// no type for readable
	// no type for writable
	// no type for disconnected
	// no type for new_connection
	// no type for _libsoup_reserved1
	// no type for _libsoup_reserved2
	// no type for _libsoup_reserved3
	// no type for _libsoup_reserved4
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.SoupSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same GObject.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.ToC() == recv.ToC()
}

// URI is a wrapper around the C record SoupURI.
type URI struct {
	native   *C.SoupURI
	Scheme   string
	User     string
	Password string
	Host     string
	Port     uint32
	Path     string
	Query    string
	Fragment string
}

func URINewFromC(u unsafe.Pointer) *URI {
	c := (*C.SoupURI)(u)
	if c == nil {
		return nil
	}

	g := &URI{
		Fragment: C.GoString(c.fragment),
		Host:     C.GoString(c.host),
		Password: C.GoString(c.password),
		Path:     C.GoString(c.path),
		Port:     (uint32)(c.port),
		Query:    C.GoString(c.query),
		Scheme:   C.GoString(c.scheme),
		User:     C.GoString(c.user),
		native:   c,
	}

	return g
}

func (recv *URI) ToC() unsafe.Pointer {
	recv.native.scheme =
		C.CString(recv.Scheme)
	recv.native.user =
		C.CString(recv.User)
	recv.native.password =
		C.CString(recv.Password)
	recv.native.host =
		C.CString(recv.Host)
	recv.native.port =
		(C.guint)(recv.Port)
	recv.native.path =
		C.CString(recv.Path)
	recv.native.query =
		C.CString(recv.Query)
	recv.native.fragment =
		C.CString(recv.Fragment)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URI with another URI, and returns true if they represent the same GObject.
func (recv *URI) Equals(other *URI) bool {
	return other.ToC() == recv.ToC()
}

// URINew is a wrapper around the C function soup_uri_new.
func URINew(uriString string) *URI {
	c_uri_string := C.CString(uriString)
	defer C.free(unsafe.Pointer(c_uri_string))

	retC := C.soup_uri_new(c_uri_string)
	var retGo (*URI)
	if retC == nil {
		retGo = nil
	} else {
		retGo = URINewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// URIDecode is a wrapper around the C function soup_uri_decode.
func URIDecode(part string) string {
	c_part := C.CString(part)
	defer C.free(unsafe.Pointer(c_part))

	retC := C.soup_uri_decode(c_part)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// URIEncode is a wrapper around the C function soup_uri_encode.
func URIEncode(part string, escapeExtra string) string {
	c_part := C.CString(part)
	defer C.free(unsafe.Pointer(c_part))

	c_escape_extra := C.CString(escapeExtra)
	defer C.free(unsafe.Pointer(c_escape_extra))

	retC := C.soup_uri_encode(c_part, c_escape_extra)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// URINormalize is a wrapper around the C function soup_uri_normalize.
func URINormalize(part string, unescapeExtra string) string {
	c_part := C.CString(part)
	defer C.free(unsafe.Pointer(c_part))

	c_unescape_extra := C.CString(unescapeExtra)
	defer C.free(unsafe.Pointer(c_unescape_extra))

	retC := C.soup_uri_normalize(c_part, c_unescape_extra)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function soup_uri_copy.
func (recv *URI) Copy() *URI {
	retC := C.soup_uri_copy((*C.SoupURI)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// CopyHost is a wrapper around the C function soup_uri_copy_host.
func (recv *URI) CopyHost() *URI {
	retC := C.soup_uri_copy_host((*C.SoupURI)(recv.native))
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function soup_uri_equal.
func (recv *URI) Equal(uri2 *URI) bool {
	c_uri2 := (*C.SoupURI)(C.NULL)
	if uri2 != nil {
		c_uri2 = (*C.SoupURI)(uri2.ToC())
	}

	retC := C.soup_uri_equal((*C.SoupURI)(recv.native), c_uri2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function soup_uri_free.
func (recv *URI) Free() {
	C.soup_uri_free((*C.SoupURI)(recv.native))

	return
}

// GetFragment is a wrapper around the C function soup_uri_get_fragment.
func (recv *URI) GetFragment() string {
	retC := C.soup_uri_get_fragment((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetHost is a wrapper around the C function soup_uri_get_host.
func (recv *URI) GetHost() string {
	retC := C.soup_uri_get_host((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPassword is a wrapper around the C function soup_uri_get_password.
func (recv *URI) GetPassword() string {
	retC := C.soup_uri_get_password((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPath is a wrapper around the C function soup_uri_get_path.
func (recv *URI) GetPath() string {
	retC := C.soup_uri_get_path((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function soup_uri_get_port.
func (recv *URI) GetPort() uint32 {
	retC := C.soup_uri_get_port((*C.SoupURI)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetQuery is a wrapper around the C function soup_uri_get_query.
func (recv *URI) GetQuery() string {
	retC := C.soup_uri_get_query((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetScheme is a wrapper around the C function soup_uri_get_scheme.
func (recv *URI) GetScheme() string {
	retC := C.soup_uri_get_scheme((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUser is a wrapper around the C function soup_uri_get_user.
func (recv *URI) GetUser() string {
	retC := C.soup_uri_get_user((*C.SoupURI)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// HostEqual is a wrapper around the C function soup_uri_host_equal.
func (recv *URI) HostEqual(v2 *URI) bool {
	c_v2 := (C.gconstpointer)(C.NULL)
	if v2 != nil {
		c_v2 = (C.gconstpointer)(v2.ToC())
	}

	retC := C.soup_uri_host_equal((C.gconstpointer)(recv.native), c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// HostHash is a wrapper around the C function soup_uri_host_hash.
func (recv *URI) HostHash() uint32 {
	retC := C.soup_uri_host_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// NewWithBase is a wrapper around the C function soup_uri_new_with_base.
func (recv *URI) NewWithBase(uriString string) *URI {
	c_uri_string := C.CString(uriString)
	defer C.free(unsafe.Pointer(c_uri_string))

	retC := C.soup_uri_new_with_base((*C.SoupURI)(recv.native), c_uri_string)
	retGo := URINewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFragment is a wrapper around the C function soup_uri_set_fragment.
func (recv *URI) SetFragment(fragment string) {
	c_fragment := C.CString(fragment)
	defer C.free(unsafe.Pointer(c_fragment))

	C.soup_uri_set_fragment((*C.SoupURI)(recv.native), c_fragment)

	return
}

// SetHost is a wrapper around the C function soup_uri_set_host.
func (recv *URI) SetHost(host string) {
	c_host := C.CString(host)
	defer C.free(unsafe.Pointer(c_host))

	C.soup_uri_set_host((*C.SoupURI)(recv.native), c_host)

	return
}

// SetPassword is a wrapper around the C function soup_uri_set_password.
func (recv *URI) SetPassword(password string) {
	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.soup_uri_set_password((*C.SoupURI)(recv.native), c_password)

	return
}

// SetPath is a wrapper around the C function soup_uri_set_path.
func (recv *URI) SetPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.soup_uri_set_path((*C.SoupURI)(recv.native), c_path)

	return
}

// SetPort is a wrapper around the C function soup_uri_set_port.
func (recv *URI) SetPort(port uint32) {
	c_port := (C.guint)(port)

	C.soup_uri_set_port((*C.SoupURI)(recv.native), c_port)

	return
}

// SetQuery is a wrapper around the C function soup_uri_set_query.
func (recv *URI) SetQuery(query string) {
	c_query := C.CString(query)
	defer C.free(unsafe.Pointer(c_query))

	C.soup_uri_set_query((*C.SoupURI)(recv.native), c_query)

	return
}

// Unsupported : soup_uri_set_query_from_fields : unsupported parameter ... : varargs

// SetQueryFromForm is a wrapper around the C function soup_uri_set_query_from_form.
func (recv *URI) SetQueryFromForm(form *glib.HashTable) {
	c_form := (*C.GHashTable)(C.NULL)
	if form != nil {
		c_form = (*C.GHashTable)(form.ToC())
	}

	C.soup_uri_set_query_from_form((*C.SoupURI)(recv.native), c_form)

	return
}

// SetScheme is a wrapper around the C function soup_uri_set_scheme.
func (recv *URI) SetScheme(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.soup_uri_set_scheme((*C.SoupURI)(recv.native), c_scheme)

	return
}

// SetUser is a wrapper around the C function soup_uri_set_user.
func (recv *URI) SetUser(user string) {
	c_user := C.CString(user)
	defer C.free(unsafe.Pointer(c_user))

	C.soup_uri_set_user((*C.SoupURI)(recv.native), c_user)

	return
}

// ToString is a wrapper around the C function soup_uri_to_string.
func (recv *URI) ToString(justPathAndQuery bool) string {
	c_just_path_and_query :=
		boolToGboolean(justPathAndQuery)

	retC := C.soup_uri_to_string((*C.SoupURI)(recv.native), c_just_path_and_query)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UsesDefaultPort is a wrapper around the C function soup_uri_uses_default_port.
func (recv *URI) UsesDefaultPort() bool {
	retC := C.soup_uri_uses_default_port((*C.SoupURI)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WebsocketConnectionClass is a wrapper around the C record SoupWebsocketConnectionClass.
type WebsocketConnectionClass struct {
	native *C.SoupWebsocketConnectionClass
	// parent : record
	// no type for message
	// no type for error
	// no type for closing
	// no type for closed
	// no type for pong
}

func WebsocketConnectionClassNewFromC(u unsafe.Pointer) *WebsocketConnectionClass {
	c := (*C.SoupWebsocketConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &WebsocketConnectionClass{native: c}

	return g
}

func (recv *WebsocketConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsocketConnectionClass with another WebsocketConnectionClass, and returns true if they represent the same GObject.
func (recv *WebsocketConnectionClass) Equals(other *WebsocketConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// WebsocketConnectionPrivate is a wrapper around the C record SoupWebsocketConnectionPrivate.
type WebsocketConnectionPrivate struct {
	native *C.SoupWebsocketConnectionPrivate
}

func WebsocketConnectionPrivateNewFromC(u unsafe.Pointer) *WebsocketConnectionPrivate {
	c := (*C.SoupWebsocketConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebsocketConnectionPrivate{native: c}

	return g
}

func (recv *WebsocketConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsocketConnectionPrivate with another WebsocketConnectionPrivate, and returns true if they represent the same GObject.
func (recv *WebsocketConnectionPrivate) Equals(other *WebsocketConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}
