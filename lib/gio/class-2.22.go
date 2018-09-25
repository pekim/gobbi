// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Socket is a wrapper around the C record GSocket.
type Socket struct {
	native *C.GSocket
	// parent_instance : record
	// priv : record
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.GSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketNew is a wrapper around the C function g_socket_new.
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) (*Socket, error) {
	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	var cThrowableError *C.GError

	retC := C.g_socket_new(c_family, c_type, c_protocol, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SocketNewFromFd is a wrapper around the C function g_socket_new_from_fd.
func SocketNewFromFd(fd int32) (*Socket, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_socket_new_from_fd(c_fd, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Accept is a wrapper around the C function g_socket_accept.
func (recv *Socket) Accept(cancellable *Cancellable) (*Socket, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_accept((*C.GSocket)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Bind is a wrapper around the C function g_socket_bind.
func (recv *Socket) Bind(address *SocketAddress, allowReuse bool) (bool, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

	c_allow_reuse :=
		boolToGboolean(allowReuse)

	var cThrowableError *C.GError

	retC := C.g_socket_bind((*C.GSocket)(recv.native), c_address, c_allow_reuse, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// CheckConnectResult is a wrapper around the C function g_socket_check_connect_result.
func (recv *Socket) CheckConnectResult() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_check_connect_result((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Close is a wrapper around the C function g_socket_close.
func (recv *Socket) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_close((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ConditionCheck is a wrapper around the C function g_socket_condition_check.
func (recv *Socket) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_socket_condition_check((*C.GSocket)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// ConditionWait is a wrapper around the C function g_socket_condition_wait.
func (recv *Socket) ConditionWait(condition glib.IOCondition, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_condition_wait((*C.GSocket)(recv.native), c_condition, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Connect is a wrapper around the C function g_socket_connect.
func (recv *Socket) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_connect((*C.GSocket)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ConnectionFactoryCreateConnection is a wrapper around the C function g_socket_connection_factory_create_connection.
func (recv *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	retC := C.g_socket_connection_factory_create_connection((*C.GSocket)(recv.native))
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateSource is a wrapper around the C function g_socket_create_source.
func (recv *Socket) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_socket_create_source((*C.GSocket)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBlocking is a wrapper around the C function g_socket_get_blocking.
func (recv *Socket) GetBlocking() bool {
	retC := C.g_socket_get_blocking((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFamily is a wrapper around the C function g_socket_get_family.
func (recv *Socket) GetFamily() SocketFamily {
	retC := C.g_socket_get_family((*C.GSocket)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetFd is a wrapper around the C function g_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	retC := C.g_socket_get_fd((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetKeepalive is a wrapper around the C function g_socket_get_keepalive.
func (recv *Socket) GetKeepalive() bool {
	retC := C.g_socket_get_keepalive((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetListenBacklog is a wrapper around the C function g_socket_get_listen_backlog.
func (recv *Socket) GetListenBacklog() int32 {
	retC := C.g_socket_get_listen_backlog((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function g_socket_get_local_address.
func (recv *Socket) GetLocalAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_local_address((*C.GSocket)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_get_option : unsupported parameter value : no type generator for gint, gint*

// GetProtocol is a wrapper around the C function g_socket_get_protocol.
func (recv *Socket) GetProtocol() SocketProtocol {
	retC := C.g_socket_get_protocol((*C.GSocket)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// GetRemoteAddress is a wrapper around the C function g_socket_get_remote_address.
func (recv *Socket) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_remote_address((*C.GSocket)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetSocketType is a wrapper around the C function g_socket_get_socket_type.
func (recv *Socket) GetSocketType() SocketType {
	retC := C.g_socket_get_socket_type((*C.GSocket)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// IsClosed is a wrapper around the C function g_socket_is_closed.
func (recv *Socket) IsClosed() bool {
	retC := C.g_socket_is_closed((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsConnected is a wrapper around the C function g_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	retC := C.g_socket_is_connected((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listen is a wrapper around the C function g_socket_listen.
func (recv *Socket) Listen() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_listen((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_receive : unsupported parameter buffer : no param type

// Unsupported : g_socket_receive_from : unsupported parameter buffer : no param type

// Unsupported : g_socket_receive_message : unsupported parameter vectors : no param type

// Unsupported : g_socket_receive_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_receive_with_blocking : unsupported parameter buffer : no param type

// Unsupported : g_socket_send : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_message : unsupported parameter vectors : no param type

// Unsupported : g_socket_send_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_send_to : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_with_blocking : unsupported parameter buffer : no param type

// Unsupported : g_socket_set_blocking : no return generator

// Unsupported : g_socket_set_broadcast : no return generator

// Unsupported : g_socket_set_keepalive : no return generator

// Unsupported : g_socket_set_listen_backlog : no return generator

// Unsupported : g_socket_set_multicast_loopback : no return generator

// Unsupported : g_socket_set_multicast_ttl : no return generator

// Unsupported : g_socket_set_timeout : no return generator

// Unsupported : g_socket_set_ttl : no return generator

// Shutdown is a wrapper around the C function g_socket_shutdown.
func (recv *Socket) Shutdown(shutdownRead bool, shutdownWrite bool) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	var cThrowableError *C.GError

	retC := C.g_socket_shutdown((*C.GSocket)(recv.native), c_shutdown_read, c_shutdown_write, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SpeaksIpv4 is a wrapper around the C function g_socket_speaks_ipv4.
func (recv *Socket) SpeaksIpv4() bool {
	retC := C.g_socket_speaks_ipv4((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SocketClient is a wrapper around the C record GSocketClient.
type SocketClient struct {
	native *C.GSocketClient
	// parent_instance : record
	// priv : record
}

func SocketClientNewFromC(u unsafe.Pointer) *SocketClient {
	c := (*C.GSocketClient)(u)
	if c == nil {
		return nil
	}

	g := &SocketClient{native: c}

	return g
}

func (recv *SocketClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClientNew is a wrapper around the C function g_socket_client_new.
func SocketClientNew() *SocketClient {
	retC := C.g_socket_client_new()
	retGo := SocketClientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_socket_client_add_application_proxy : no return generator

// Unsupported : g_socket_client_connect : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_async : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ConnectToHost is a wrapper around the C function g_socket_client_connect_to_host.
func (recv *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_host_and_port := C.CString(hostAndPort)
	defer C.free(unsafe.Pointer(c_host_and_port))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host((*C.GSocketClient)(recv.native), c_host_and_port, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_host_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_uri_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetFamily is a wrapper around the C function g_socket_client_get_family.
func (recv *SocketClient) GetFamily() SocketFamily {
	retC := C.g_socket_client_get_family((*C.GSocketClient)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function g_socket_client_get_local_address.
func (recv *SocketClient) GetLocalAddress() *SocketAddress {
	retC := C.g_socket_client_get_local_address((*C.GSocketClient)(recv.native))
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProtocol is a wrapper around the C function g_socket_client_get_protocol.
func (recv *SocketClient) GetProtocol() SocketProtocol {
	retC := C.g_socket_client_get_protocol((*C.GSocketClient)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// Unsupported : g_socket_client_get_proxy_resolver : no return generator

// GetSocketType is a wrapper around the C function g_socket_client_get_socket_type.
func (recv *SocketClient) GetSocketType() SocketType {
	retC := C.g_socket_client_get_socket_type((*C.GSocketClient)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// Unsupported : g_socket_client_set_enable_proxy : no return generator

// Unsupported : g_socket_client_set_family : no return generator

// Unsupported : g_socket_client_set_local_address : no return generator

// Unsupported : g_socket_client_set_protocol : no return generator

// Unsupported : g_socket_client_set_proxy_resolver : unsupported parameter proxy_resolver : no type generator for ProxyResolver, GProxyResolver*

// Unsupported : g_socket_client_set_socket_type : no return generator

// Unsupported : g_socket_client_set_timeout : no return generator

// Unsupported : g_socket_client_set_tls : no return generator

// Unsupported : g_socket_client_set_tls_validation_flags : no return generator

// SocketConnection is a wrapper around the C record GSocketConnection.
type SocketConnection struct {
	native *C.GSocketConnection
	// parent_instance : record
	// priv : record
}

func SocketConnectionNewFromC(u unsafe.Pointer) *SocketConnection {
	c := (*C.GSocketConnection)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnection{native: c}

	return g
}

func (recv *SocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_connection_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetLocalAddress is a wrapper around the C function g_socket_connection_get_local_address.
func (recv *SocketConnection) GetLocalAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_local_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetRemoteAddress is a wrapper around the C function g_socket_connection_get_remote_address.
func (recv *SocketConnection) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_remote_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetSocket is a wrapper around the C function g_socket_connection_get_socket.
func (recv *SocketConnection) GetSocket() *Socket {
	retC := C.g_socket_connection_get_socket((*C.GSocketConnection)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SocketListener is a wrapper around the C record GSocketListener.
type SocketListener struct {
	native *C.GSocketListener
	// parent_instance : record
	// priv : record
}

func SocketListenerNewFromC(u unsafe.Pointer) *SocketListener {
	c := (*C.GSocketListener)(u)
	if c == nil {
		return nil
	}

	g := &SocketListener{native: c}

	return g
}

func (recv *SocketListener) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketListenerNew is a wrapper around the C function g_socket_listener_new.
func SocketListenerNew() *SocketListener {
	retC := C.g_socket_listener_new()
	retGo := SocketListenerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Accept is a wrapper around the C function g_socket_listener_accept.
func (recv *SocketListener) Accept(cancellable *Cancellable) (*SocketConnection, **gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// AcceptSocket is a wrapper around the C function g_socket_listener_accept_socket.
func (recv *SocketListener) AcceptSocket(cancellable *Cancellable) (*Socket, **gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_socket((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_socket_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// AddAddress is a wrapper around the C function g_socket_listener_add_address.
func (recv *SocketListener) AddAddress(address *SocketAddress, type_ SocketType, protocol SocketProtocol, sourceObject *gobject.Object) (bool, **SocketAddress, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	c_source_object := (*C.GObject)(sourceObject.ToC())

	var c_effective_address *C.GSocketAddress

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_address((*C.GSocketListener)(recv.native), c_address, c_type, c_protocol, c_source_object, &c_effective_address, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	effectiveAddress := SocketAddressNewFromC(unsafe.Pointer(c_effective_address))

	return retGo, effectiveAddress, goThrowableError
}

// AddInetPort is a wrapper around the C function g_socket_listener_add_inet_port.
func (recv *SocketListener) AddInetPort(port uint16, sourceObject *gobject.Object) (bool, error) {
	c_port := (C.guint16)(port)

	c_source_object := (*C.GObject)(sourceObject.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_inet_port((*C.GSocketListener)(recv.native), c_port, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// AddSocket is a wrapper around the C function g_socket_listener_add_socket.
func (recv *SocketListener) AddSocket(socket *Socket, sourceObject *gobject.Object) (bool, error) {
	c_socket := (*C.GSocket)(socket.ToC())

	c_source_object := (*C.GObject)(sourceObject.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_socket((*C.GSocketListener)(recv.native), c_socket, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_listener_close : no return generator

// Unsupported : g_socket_listener_set_backlog : no return generator

// SocketService is a wrapper around the C record GSocketService.
type SocketService struct {
	native *C.GSocketService
	// parent_instance : record
	// priv : record
}

func SocketServiceNewFromC(u unsafe.Pointer) *SocketService {
	c := (*C.GSocketService)(u)
	if c == nil {
		return nil
	}

	g := &SocketService{native: c}

	return g
}

func (recv *SocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketServiceNew is a wrapper around the C function g_socket_service_new.
func SocketServiceNew() *SocketService {
	retC := C.g_socket_service_new()
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsActive is a wrapper around the C function g_socket_service_is_active.
func (recv *SocketService) IsActive() bool {
	retC := C.g_socket_service_is_active((*C.GSocketService)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_socket_service_start : no return generator

// Unsupported : g_socket_service_stop : no return generator

// TcpConnection is a wrapper around the C record GTcpConnection.
type TcpConnection struct {
	native *C.GTcpConnection
	// parent_instance : record
	// priv : record
}

func TcpConnectionNewFromC(u unsafe.Pointer) *TcpConnection {
	c := (*C.GTcpConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnection{native: c}

	return g
}

func (recv *TcpConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetGracefulDisconnect is a wrapper around the C function g_tcp_connection_get_graceful_disconnect.
func (recv *TcpConnection) GetGracefulDisconnect() bool {
	retC := C.g_tcp_connection_get_graceful_disconnect((*C.GTcpConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_tcp_connection_set_graceful_disconnect : no return generator

// ThreadedSocketService is a wrapper around the C record GThreadedSocketService.
type ThreadedSocketService struct {
	native *C.GThreadedSocketService
	// parent_instance : record
	// priv : record
}

func ThreadedSocketServiceNewFromC(u unsafe.Pointer) *ThreadedSocketService {
	c := (*C.GThreadedSocketService)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketService{native: c}

	return g
}

func (recv *ThreadedSocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadedSocketServiceNew is a wrapper around the C function g_threaded_socket_service_new.
func ThreadedSocketServiceNew(maxThreads int32) *SocketService {
	c_max_threads := (C.int)(maxThreads)

	retC := C.g_threaded_socket_service_new(c_max_threads)
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}
