// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
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

	gboolean socketservice_incomingHandler(GObject *, GSocketConnection *, GObject *, gpointer);

	static gulong SocketService_signal_connect_incoming(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "incoming", G_CALLBACK(socketservice_incomingHandler), data);
	}

*/
/*

	gboolean threadedsocketservice_runHandler(GObject *, GSocketConnection *, GObject *, gpointer);

	static gulong ThreadedSocketService_signal_connect_run(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "run", G_CALLBACK(threadedsocketservice_runHandler), data);
	}

*/
/*

	void volumemonitor_driveStopButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_stop_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-stop-button", G_CALLBACK(volumemonitor_driveStopButtonHandler), data);
	}

*/
import "C"

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Blacklisted : g_cancellable_disconnect

// Blacklisted : g_cancellable_make_pollfd

// Blacklisted : g_cancellable_release_fd

// Blacklisted : g_file_io_stream_get_etag

// Blacklisted : g_file_io_stream_query_info

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_io_stream_query_info_finish

// Blacklisted : g_file_info_get_attribute_stringv

// Blacklisted : g_file_info_has_namespace

// Blacklisted : g_file_info_set_attribute_status

// Blacklisted : g_io_stream_clear_pending

// Blacklisted : g_io_stream_close

// Unsupported : g_io_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_io_stream_close_finish

// Blacklisted : g_io_stream_get_input_stream

// Blacklisted : g_io_stream_get_output_stream

// Blacklisted : g_io_stream_has_pending

// Blacklisted : g_io_stream_is_closed

// Blacklisted : g_io_stream_set_pending

// Blacklisted : g_inet_address_new_any

// Blacklisted : g_inet_address_new_from_bytes

// Blacklisted : g_inet_address_new_from_string

// Blacklisted : g_inet_address_new_loopback

// Blacklisted : g_inet_address_get_family

// Blacklisted : g_inet_address_get_is_any

// Blacklisted : g_inet_address_get_is_link_local

// Blacklisted : g_inet_address_get_is_loopback

// Blacklisted : g_inet_address_get_is_mc_global

// Blacklisted : g_inet_address_get_is_mc_link_local

// Blacklisted : g_inet_address_get_is_mc_node_local

// Blacklisted : g_inet_address_get_is_mc_org_local

// Blacklisted : g_inet_address_get_is_mc_site_local

// Blacklisted : g_inet_address_get_is_multicast

// Blacklisted : g_inet_address_get_is_site_local

// Blacklisted : g_inet_address_get_native_size

// Blacklisted : g_inet_address_to_bytes

// Blacklisted : g_inet_address_to_string

// Blacklisted : g_inet_socket_address_new

// Blacklisted : g_inet_socket_address_get_address

// Blacklisted : g_inet_socket_address_get_port

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter processes :

// Blacklisted : g_network_address_new

// Blacklisted : g_network_address_parse

// Blacklisted : g_network_address_get_hostname

// Blacklisted : g_network_address_get_port

// Blacklisted : g_network_service_new

// Blacklisted : g_network_service_get_domain

// Blacklisted : g_network_service_get_protocol

// Blacklisted : g_network_service_get_service

// Blacklisted : g_resolver_free_addresses

// Blacklisted : g_resolver_free_targets

// Blacklisted : g_resolver_get_default

// Blacklisted : g_resolver_lookup_by_address

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_resolver_lookup_by_address_finish

// Blacklisted : g_resolver_lookup_by_name

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_resolver_lookup_by_name_finish

// Blacklisted : g_resolver_lookup_service

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_resolver_lookup_service_finish

// Blacklisted : g_resolver_set_default

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

// Blacklisted : g_socket_new

// Blacklisted : g_socket_new_from_fd

// Blacklisted : g_socket_accept

// Blacklisted : g_socket_bind

// Blacklisted : g_socket_check_connect_result

// Blacklisted : g_socket_close

// Blacklisted : g_socket_condition_check

// Blacklisted : g_socket_condition_wait

// Blacklisted : g_socket_connect

// Blacklisted : g_socket_connection_factory_create_connection

// Blacklisted : g_socket_create_source

// Blacklisted : g_socket_get_blocking

// Blacklisted : g_socket_get_family

// Blacklisted : g_socket_get_fd

// Blacklisted : g_socket_get_keepalive

// Blacklisted : g_socket_get_listen_backlog

// Blacklisted : g_socket_get_local_address

// Blacklisted : g_socket_get_protocol

// Blacklisted : g_socket_get_remote_address

// Blacklisted : g_socket_get_socket_type

// Blacklisted : g_socket_is_closed

// Blacklisted : g_socket_is_connected

// Blacklisted : g_socket_listen

// Blacklisted : g_socket_receive

// Blacklisted : g_socket_receive_from

// Unsupported : g_socket_receive_message : unsupported parameter vectors :

// Blacklisted : g_socket_send

// Unsupported : g_socket_send_message : unsupported parameter vectors :

// Blacklisted : g_socket_send_to

// Blacklisted : g_socket_set_blocking

// Blacklisted : g_socket_set_keepalive

// Blacklisted : g_socket_set_listen_backlog

// Blacklisted : g_socket_shutdown

// Blacklisted : g_socket_speaks_ipv4

// Unsupported : g_socket_address_new_from_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// Blacklisted : g_socket_address_get_family

// Blacklisted : g_socket_address_get_native_size

// Unsupported : g_socket_address_to_native : unsupported parameter dest : no type generator for gpointer (gpointer) for param dest

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketClient) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClient with another SocketClient, and returns true if they represent the same GObject.
func (recv *SocketClient) Equals(other *SocketClient) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketClient) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketClient.
func CastToSocketClient(object *gobject.Object) *SocketClient {
	return SocketClientNewFromC(object.ToC())
}

// Blacklisted : g_socket_client_new

// Blacklisted : g_socket_client_add_application_proxy

// Blacklisted : g_socket_client_connect

// Unsupported : g_socket_client_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_client_connect_finish

// Blacklisted : g_socket_client_connect_to_host

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_client_connect_to_host_finish

// Blacklisted : g_socket_client_connect_to_service

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_client_connect_to_service_finish

// Blacklisted : g_socket_client_get_family

// Blacklisted : g_socket_client_get_local_address

// Blacklisted : g_socket_client_get_protocol

// Blacklisted : g_socket_client_get_socket_type

// Blacklisted : g_socket_client_set_family

// Blacklisted : g_socket_client_set_local_address

// Blacklisted : g_socket_client_set_protocol

// Blacklisted : g_socket_client_set_socket_type

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnection with another SocketConnection, and returns true if they represent the same GObject.
func (recv *SocketConnection) Equals(other *SocketConnection) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *SocketConnection) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketConnection) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to SocketConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketConnection.
func CastToSocketConnection(object *gobject.Object) *SocketConnection {
	return SocketConnectionNewFromC(object.ToC())
}

// Blacklisted : g_socket_connection_factory_lookup_type

// Blacklisted : g_socket_connection_factory_register_type

// Blacklisted : g_socket_connection_get_local_address

// Blacklisted : g_socket_connection_get_remote_address

// Blacklisted : g_socket_connection_get_socket

// Blacklisted : g_socket_control_message_deserialize

// Blacklisted : g_socket_control_message_get_level

// Blacklisted : g_socket_control_message_get_msg_type

// Blacklisted : g_socket_control_message_get_size

// Unsupported : g_socket_control_message_serialize : unsupported parameter data : no type generator for gpointer (gpointer) for param data

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketListener) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketListener) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketListener with another SocketListener, and returns true if they represent the same GObject.
func (recv *SocketListener) Equals(other *SocketListener) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketListener) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketListener.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketListener.
func CastToSocketListener(object *gobject.Object) *SocketListener {
	return SocketListenerNewFromC(object.ToC())
}

// Blacklisted : g_socket_listener_new

// Blacklisted : g_socket_listener_accept

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_listener_accept_finish

// Blacklisted : g_socket_listener_accept_socket

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_socket_listener_accept_socket_finish

// Blacklisted : g_socket_listener_add_address

// Blacklisted : g_socket_listener_add_inet_port

// Blacklisted : g_socket_listener_add_socket

// Blacklisted : g_socket_listener_close

// Blacklisted : g_socket_listener_set_backlog

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketService with another SocketService, and returns true if they represent the same GObject.
func (recv *SocketService) Equals(other *SocketService) bool {
	return other.ToC() == recv.ToC()
}

// SocketListener upcasts to *SocketListener
func (recv *SocketService) SocketListener() *SocketListener {
	return SocketListenerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketService) Object() *gobject.Object {
	return recv.SocketListener().Object()
}

// CastToWidget down casts any arbitrary Object to SocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketService.
func CastToSocketService(object *gobject.Object) *SocketService {
	return SocketServiceNewFromC(object.ToC())
}

type signalSocketServiceIncomingDetail struct {
	callback  SocketServiceSignalIncomingCallback
	handlerID C.gulong
}

var signalSocketServiceIncomingId int
var signalSocketServiceIncomingMap = make(map[int]signalSocketServiceIncomingDetail)
var signalSocketServiceIncomingLock sync.RWMutex

// SocketServiceSignalIncomingCallback is a callback function for a 'incoming' signal emitted from a SocketService.
type SocketServiceSignalIncomingCallback func(connection *SocketConnection, sourceObject *gobject.Object) bool

/*
ConnectIncoming connects the callback to the 'incoming' signal for the SocketService.

The returned value represents the connection, and may be passed to DisconnectIncoming to remove it.
*/
func (recv *SocketService) ConnectIncoming(callback SocketServiceSignalIncomingCallback) int {
	signalSocketServiceIncomingLock.Lock()
	defer signalSocketServiceIncomingLock.Unlock()

	signalSocketServiceIncomingId++
	instance := C.gpointer(recv.native)
	handlerID := C.SocketService_signal_connect_incoming(instance, C.gpointer(uintptr(signalSocketServiceIncomingId)))

	detail := signalSocketServiceIncomingDetail{callback, handlerID}
	signalSocketServiceIncomingMap[signalSocketServiceIncomingId] = detail

	return signalSocketServiceIncomingId
}

/*
DisconnectIncoming disconnects a callback from the 'incoming' signal for the SocketService.

The connectionID should be a value returned from a call to ConnectIncoming.
*/
func (recv *SocketService) DisconnectIncoming(connectionID int) {
	signalSocketServiceIncomingLock.Lock()
	defer signalSocketServiceIncomingLock.Unlock()

	detail, exists := signalSocketServiceIncomingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketServiceIncomingMap, connectionID)
}

//export socketservice_incomingHandler
func socketservice_incomingHandler(_ *C.GObject, c_connection *C.GSocketConnection, c_source_object *C.GObject, data C.gpointer) C.gboolean {
	signalSocketServiceIncomingLock.RLock()
	defer signalSocketServiceIncomingLock.RUnlock()

	connection := SocketConnectionNewFromC(unsafe.Pointer(c_connection))

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	index := int(uintptr(data))
	callback := signalSocketServiceIncomingMap[index].callback
	retGo := callback(connection, sourceObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_socket_service_new

// Blacklisted : g_socket_service_is_active

// Blacklisted : g_socket_service_start

// Blacklisted : g_socket_service_stop

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TcpConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TcpConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpConnection with another TcpConnection, and returns true if they represent the same GObject.
func (recv *TcpConnection) Equals(other *TcpConnection) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromC(unsafe.Pointer(recv.native))
}

// IOStream upcasts to *IOStream
func (recv *TcpConnection) IOStream() *IOStream {
	return recv.SocketConnection().IOStream()
}

// Object upcasts to *Object
func (recv *TcpConnection) Object() *gobject.Object {
	return recv.SocketConnection().Object()
}

// CastToWidget down casts any arbitrary Object to TcpConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TcpConnection.
func CastToTcpConnection(object *gobject.Object) *TcpConnection {
	return TcpConnectionNewFromC(object.ToC())
}

// Blacklisted : g_tcp_connection_get_graceful_disconnect

// Blacklisted : g_tcp_connection_set_graceful_disconnect

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ThreadedSocketService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ThreadedSocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadedSocketService with another ThreadedSocketService, and returns true if they represent the same GObject.
func (recv *ThreadedSocketService) Equals(other *ThreadedSocketService) bool {
	return other.ToC() == recv.ToC()
}

// SocketService upcasts to *SocketService
func (recv *ThreadedSocketService) SocketService() *SocketService {
	return SocketServiceNewFromC(unsafe.Pointer(recv.native))
}

// SocketListener upcasts to *SocketListener
func (recv *ThreadedSocketService) SocketListener() *SocketListener {
	return recv.SocketService().SocketListener()
}

// Object upcasts to *Object
func (recv *ThreadedSocketService) Object() *gobject.Object {
	return recv.SocketService().Object()
}

// CastToWidget down casts any arbitrary Object to ThreadedSocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThreadedSocketService.
func CastToThreadedSocketService(object *gobject.Object) *ThreadedSocketService {
	return ThreadedSocketServiceNewFromC(object.ToC())
}

type signalThreadedSocketServiceRunDetail struct {
	callback  ThreadedSocketServiceSignalRunCallback
	handlerID C.gulong
}

var signalThreadedSocketServiceRunId int
var signalThreadedSocketServiceRunMap = make(map[int]signalThreadedSocketServiceRunDetail)
var signalThreadedSocketServiceRunLock sync.RWMutex

// ThreadedSocketServiceSignalRunCallback is a callback function for a 'run' signal emitted from a ThreadedSocketService.
type ThreadedSocketServiceSignalRunCallback func(connection *SocketConnection, sourceObject *gobject.Object) bool

/*
ConnectRun connects the callback to the 'run' signal for the ThreadedSocketService.

The returned value represents the connection, and may be passed to DisconnectRun to remove it.
*/
func (recv *ThreadedSocketService) ConnectRun(callback ThreadedSocketServiceSignalRunCallback) int {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	signalThreadedSocketServiceRunId++
	instance := C.gpointer(recv.native)
	handlerID := C.ThreadedSocketService_signal_connect_run(instance, C.gpointer(uintptr(signalThreadedSocketServiceRunId)))

	detail := signalThreadedSocketServiceRunDetail{callback, handlerID}
	signalThreadedSocketServiceRunMap[signalThreadedSocketServiceRunId] = detail

	return signalThreadedSocketServiceRunId
}

/*
DisconnectRun disconnects a callback from the 'run' signal for the ThreadedSocketService.

The connectionID should be a value returned from a call to ConnectRun.
*/
func (recv *ThreadedSocketService) DisconnectRun(connectionID int) {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	detail, exists := signalThreadedSocketServiceRunMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalThreadedSocketServiceRunMap, connectionID)
}

//export threadedsocketservice_runHandler
func threadedsocketservice_runHandler(_ *C.GObject, c_connection *C.GSocketConnection, c_source_object *C.GObject, data C.gpointer) C.gboolean {
	signalThreadedSocketServiceRunLock.RLock()
	defer signalThreadedSocketServiceRunLock.RUnlock()

	connection := SocketConnectionNewFromC(unsafe.Pointer(c_connection))

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	index := int(uintptr(data))
	callback := signalThreadedSocketServiceRunMap[index].callback
	retGo := callback(connection, sourceObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_threaded_socket_service_new

// Blacklisted : g_unix_connection_receive_fd

// Blacklisted : g_unix_connection_send_fd

// Blacklisted : g_unix_fd_message_new

// Blacklisted : g_unix_fd_message_append_fd

// Unsupported : g_unix_fd_message_steal_fds : array return type :

// Blacklisted : g_unix_socket_address_new

// Blacklisted : g_unix_socket_address_abstract_names_supported

// Blacklisted : g_unix_socket_address_get_is_abstract

// Blacklisted : g_unix_socket_address_get_path

// Blacklisted : g_unix_socket_address_get_path_len

type signalVolumeMonitorDriveStopButtonDetail struct {
	callback  VolumeMonitorSignalDriveStopButtonCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveStopButtonId int
var signalVolumeMonitorDriveStopButtonMap = make(map[int]signalVolumeMonitorDriveStopButtonDetail)
var signalVolumeMonitorDriveStopButtonLock sync.RWMutex

// VolumeMonitorSignalDriveStopButtonCallback is a callback function for a 'drive-stop-button' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveStopButtonCallback func(drive *Drive)

/*
ConnectDriveStopButton connects the callback to the 'drive-stop-button' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveStopButton to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveStopButton(callback VolumeMonitorSignalDriveStopButtonCallback) int {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	signalVolumeMonitorDriveStopButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_stop_button(instance, C.gpointer(uintptr(signalVolumeMonitorDriveStopButtonId)))

	detail := signalVolumeMonitorDriveStopButtonDetail{callback, handlerID}
	signalVolumeMonitorDriveStopButtonMap[signalVolumeMonitorDriveStopButtonId] = detail

	return signalVolumeMonitorDriveStopButtonId
}

/*
DisconnectDriveStopButton disconnects a callback from the 'drive-stop-button' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveStopButton.
*/
func (recv *VolumeMonitor) DisconnectDriveStopButton(connectionID int) {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	detail, exists := signalVolumeMonitorDriveStopButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveStopButtonMap, connectionID)
}

//export volumemonitor_driveStopButtonHandler
func volumemonitor_driveStopButtonHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveStopButtonLock.RLock()
	defer signalVolumeMonitorDriveStopButtonLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveStopButtonMap[index].callback
	callback(drive)
}
