// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Unsupported : g_inet_address_new_any : return type :

// Unsupported : g_inet_address_new_from_bytes : return type :

// Unsupported : g_inet_address_new_from_string : return type :

// Unsupported : g_inet_address_new_loopback : return type :

// Unsupported : g_inet_socket_address_new : return type :

// Unsupported : g_network_address_new : return type :

// Unsupported : g_network_service_new : return type :

// Socket is a wrapper around the C record GSocket.
type Socket struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	if u == nil {
		return nil
	}

	g := &Socket{native: u}

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_socket_new : return type :

// Unsupported : g_socket_new_from_fd : return type :

// Unsupported : g_socket_address_new_from_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// SocketClient is a wrapper around the C record GSocketClient.
type SocketClient struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketClientNewFromC(u unsafe.Pointer) *SocketClient {
	if u == nil {
		return nil
	}

	g := &SocketClient{native: u}

	return g
}

func (recv *SocketClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_socket_client_new : return type :

// SocketConnection is a wrapper around the C record GSocketConnection.
type SocketConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketConnectionNewFromC(u unsafe.Pointer) *SocketConnection {
	if u == nil {
		return nil
	}

	g := &SocketConnection{native: u}

	return g
}

func (recv *SocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketListener is a wrapper around the C record GSocketListener.
type SocketListener struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketListenerNewFromC(u unsafe.Pointer) *SocketListener {
	if u == nil {
		return nil
	}

	g := &SocketListener{native: u}

	return g
}

func (recv *SocketListener) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_socket_listener_new : return type :

// SocketService is a wrapper around the C record GSocketService.
type SocketService struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketServiceNewFromC(u unsafe.Pointer) *SocketService {
	if u == nil {
		return nil
	}

	g := &SocketService{native: u}

	return g
}

func (recv *SocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_socket_service_new : return type :

// TcpConnection is a wrapper around the C record GTcpConnection.
type TcpConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TcpConnectionNewFromC(u unsafe.Pointer) *TcpConnection {
	if u == nil {
		return nil
	}

	g := &TcpConnection{native: u}

	return g
}

func (recv *TcpConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadedSocketService is a wrapper around the C record GThreadedSocketService.
type ThreadedSocketService struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func ThreadedSocketServiceNewFromC(u unsafe.Pointer) *ThreadedSocketService {
	if u == nil {
		return nil
	}

	g := &ThreadedSocketService{native: u}

	return g
}

func (recv *ThreadedSocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_threaded_socket_service_new : return type :

// Unsupported : g_unix_fd_message_new : return type :

// Unsupported : g_unix_socket_address_new : return type :
