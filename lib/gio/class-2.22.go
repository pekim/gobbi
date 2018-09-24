// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// Socket is a wrapper around the C record GSocket.
type Socket struct {
	native *C.GSocket
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *Socket) toC() *C.GSocket {

	return recv.native
}

// Unsupported : g_socket_new : no return generator

// Unsupported : g_socket_new_from_fd : no return generator

// Unsupported : g_socket_accept : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_bind : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// CheckConnectResult is a wrapper around the C function g_socket_check_connect_result.
func (recv *Socket) CheckConnectResult() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_check_connect_result((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
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

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
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

// Unsupported : g_socket_condition_timed_wait : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_condition_wait : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_connect : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_connection_factory_create_connection : no return generator

// Unsupported : g_socket_create_source : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// GetBlocking is a wrapper around the C function g_socket_get_blocking.
func (recv *Socket) GetBlocking() bool {
	retC := C.g_socket_get_blocking((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_socket_get_credentials : no return generator

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

// Unsupported : g_socket_get_local_address : no return generator

// Unsupported : g_socket_get_option : unsupported parameter value : no type generator for gint, gint*

// GetProtocol is a wrapper around the C function g_socket_get_protocol.
func (recv *Socket) GetProtocol() SocketProtocol {
	retC := C.g_socket_get_protocol((*C.GSocket)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// Unsupported : g_socket_get_remote_address : no return generator

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

// Unsupported : g_socket_join_multicast_group : unsupported parameter group : no type generator for InetAddress, GInetAddress*

// Unsupported : g_socket_join_multicast_group_ssm : unsupported parameter group : no type generator for InetAddress, GInetAddress*

// Unsupported : g_socket_leave_multicast_group : unsupported parameter group : no type generator for InetAddress, GInetAddress*

// Unsupported : g_socket_leave_multicast_group_ssm : unsupported parameter group : no type generator for InetAddress, GInetAddress*

// Listen is a wrapper around the C function g_socket_listen.
func (recv *Socket) Listen() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_listen((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_receive : unsupported parameter buffer : no param type

// Unsupported : g_socket_receive_from : unsupported parameter address : no type generator for SocketAddress, GSocketAddress**

// Unsupported : g_socket_receive_message : unsupported parameter address : no type generator for SocketAddress, GSocketAddress**

// Unsupported : g_socket_receive_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_receive_with_blocking : unsupported parameter buffer : no param type

// Unsupported : g_socket_send : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_message : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_send_messages : unsupported parameter messages : no param type

// Unsupported : g_socket_send_to : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

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
	c_shutdown_read := C.FALSE
	if shutdownRead {
		c_shutdown_read = C.TRUE
	}

	c_shutdown_write := C.FALSE
	if shutdownWrite {
		c_shutdown_write = C.TRUE
	}

	var cThrowableError *C.GError

	retC := C.g_socket_shutdown((*C.GSocket)(recv.native), c_shutdown_read, c_shutdown_write, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
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
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *SocketClient) toC() *C.GSocketClient {

	return recv.native
}

// Unsupported : g_socket_client_new : no return generator

// Unsupported : g_socket_client_add_application_proxy : no return generator

// Unsupported : g_socket_client_connect : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_async : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_host : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_host_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_service : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_client_connect_to_uri : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_client_connect_to_uri_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetFamily is a wrapper around the C function g_socket_client_get_family.
func (recv *SocketClient) GetFamily() SocketFamily {
	retC := C.g_socket_client_get_family((*C.GSocketClient)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Unsupported : g_socket_client_get_local_address : no return generator

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

// Unsupported : g_socket_client_set_local_address : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_client_set_protocol : no return generator

// Unsupported : g_socket_client_set_proxy_resolver : unsupported parameter proxy_resolver : no type generator for ProxyResolver, GProxyResolver*

// Unsupported : g_socket_client_set_socket_type : no return generator

// Unsupported : g_socket_client_set_timeout : no return generator

// Unsupported : g_socket_client_set_tls : no return generator

// Unsupported : g_socket_client_set_tls_validation_flags : no return generator

// SocketConnection is a wrapper around the C record GSocketConnection.
type SocketConnection struct {
	native *C.GSocketConnection
	// parent_instance : no type generator for IOStream, GIOStream
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

func (recv *SocketConnection) toC() *C.GSocketConnection {

	return recv.native
}

// Unsupported : g_socket_connection_connect : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_connection_connect_async : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_connection_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_connection_get_local_address : no return generator

// Unsupported : g_socket_connection_get_remote_address : no return generator

// Unsupported : g_socket_connection_get_socket : no return generator

// SocketListener is a wrapper around the C record GSocketListener.
type SocketListener struct {
	native *C.GSocketListener
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *SocketListener) toC() *C.GSocketListener {

	return recv.native
}

// Unsupported : g_socket_listener_new : no return generator

// Unsupported : g_socket_listener_accept : unsupported parameter source_object : no type generator for GObject.Object, GObject**

// Unsupported : g_socket_listener_accept_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_listener_accept_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_accept_socket : unsupported parameter source_object : no type generator for GObject.Object, GObject**

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_socket_listener_accept_socket_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_add_address : unsupported parameter address : no type generator for SocketAddress, GSocketAddress*

// Unsupported : g_socket_listener_add_any_inet_port : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_socket_listener_add_inet_port : unsupported parameter source_object : no type generator for GObject.Object, GObject*

// Unsupported : g_socket_listener_add_socket : unsupported parameter socket : no type generator for Socket, GSocket*

// Unsupported : g_socket_listener_close : no return generator

// Unsupported : g_socket_listener_set_backlog : no return generator

// SocketService is a wrapper around the C record GSocketService.
type SocketService struct {
	native *C.GSocketService
	// parent_instance : no type generator for SocketListener, GSocketListener
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

func (recv *SocketService) toC() *C.GSocketService {

	return recv.native
}

// Unsupported : g_socket_service_new : no return generator

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
	// parent_instance : no type generator for SocketConnection, GSocketConnection
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

func (recv *TcpConnection) toC() *C.GTcpConnection {

	return recv.native
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
	// parent_instance : no type generator for SocketService, GSocketService
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

func (recv *ThreadedSocketService) toC() *C.GThreadedSocketService {

	return recv.native
}

// Unsupported : g_threaded_socket_service_new : no return generator
