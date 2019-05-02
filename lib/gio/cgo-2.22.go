// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
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

// InetAddressNewAny is a wrapper around the C function g_inet_address_new_any.
func InetAddressNewAny(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_any(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewFromBytes is a wrapper around the C function g_inet_address_new_from_bytes.
func InetAddressNewFromBytes(bytes []uint8, family SocketFamily) *InetAddress {
	c_bytes_array := make([]C.guint8, len(bytes)+1, len(bytes)+1)
	for i, item := range bytes {
		c := (C.guint8)(item)
		c_bytes_array[i] = c
	}
	c_bytes_array[len(bytes)] = 0
	c_bytes_arrayPtr := &c_bytes_array[0]
	c_bytes := (*C.guint8)(unsafe.Pointer(c_bytes_arrayPtr))

	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_from_bytes(c_bytes, c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewFromString is a wrapper around the C function g_inet_address_new_from_string.
func InetAddressNewFromString(string_ string) *InetAddress {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_inet_address_new_from_string(c_string)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewLoopback is a wrapper around the C function g_inet_address_new_loopback.
func InetAddressNewLoopback(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_loopback(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetSocketAddressNew is a wrapper around the C function g_inet_socket_address_new.
func InetSocketAddressNew(address *InetAddress, port uint16) *InetSocketAddress {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	c_port := (C.guint16)(port)

	retC := C.g_inet_socket_address_new(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// NetworkAddressNew is a wrapper around the C function g_network_address_new.
func NetworkAddressNew(hostname string, port uint16) *NetworkAddress {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	retC := C.g_network_address_new(c_hostname, c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// NetworkServiceNew is a wrapper around the C function g_network_service_new.
func NetworkServiceNew(service string, protocol string, domain string) *NetworkService {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	retC := C.g_network_service_new(c_service, c_protocol, c_domain)
	retGo := NetworkServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

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

// SocketNew is a wrapper around the C function g_socket_new.
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) (*Socket, error) {
	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	var cThrowableError *C.GError

	retC := C.g_socket_new(c_family, c_type, c_protocol, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

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

// SocketNewFromFd is a wrapper around the C function g_socket_new_from_fd.
func SocketNewFromFd(fd int32) (*Socket, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_socket_new_from_fd(c_fd, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

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

// Unsupported : g_socket_address_new_from_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

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

// SocketClientNew is a wrapper around the C function g_socket_client_new.
func SocketClientNew() *SocketClient {
	retC := C.g_socket_client_new()
	retGo := SocketClientNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

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

// SocketListenerNew is a wrapper around the C function g_socket_listener_new.
func SocketListenerNew() *SocketListener {
	retC := C.g_socket_listener_new()
	retGo := SocketListenerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

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

// SocketServiceNew is a wrapper around the C function g_socket_service_new.
func SocketServiceNew() *SocketService {
	retC := C.g_socket_service_new()
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

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

// ThreadedSocketServiceNew is a wrapper around the C function g_threaded_socket_service_new.
func ThreadedSocketServiceNew(maxThreads int32) *ThreadedSocketService {
	c_max_threads := (C.int)(maxThreads)

	retC := C.g_threaded_socket_service_new(c_max_threads)
	retGo := ThreadedSocketServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixFDMessageNew is a wrapper around the C function g_unix_fd_message_new.
func UnixFDMessageNew() *UnixFDMessage {
	retC := C.g_unix_fd_message_new()
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixSocketAddressNew is a wrapper around the C function g_unix_socket_address_new.
func UnixSocketAddressNew(path string) *UnixSocketAddress {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_unix_socket_address_new(c_path)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_initable_newv : unsupported parameter parameters :

// AsyncInitable is a wrapper around the C record GAsyncInitable.
type AsyncInitable struct {
	native *C.GAsyncInitable
}

func AsyncInitableNewFromC(u unsafe.Pointer) *AsyncInitable {
	c := (*C.GAsyncInitable)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitable{native: c}

	return g
}

func (recv *AsyncInitable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Initable is a wrapper around the C record GInitable.
type Initable struct {
	native *C.GInitable
}

func InitableNewFromC(u unsafe.Pointer) *Initable {
	c := (*C.GInitable)(u)
	if c == nil {
		return nil
	}

	g := &Initable{native: c}

	return g
}

func (recv *Initable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncInitableIface is a wrapper around the C record GAsyncInitableIface.
type AsyncInitableIface struct {
	native *C.GAsyncInitableIface
	// g_iface : record
	// no type for init_async
	// no type for init_finish
}

func AsyncInitableIfaceNewFromC(u unsafe.Pointer) *AsyncInitableIface {
	c := (*C.GAsyncInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitableIface{native: c}

	return g
}

func (recv *AsyncInitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitableIface is a wrapper around the C record GInitableIface.
type InitableIface struct {
	native *C.GInitableIface
	// g_iface : record
	// no type for init
}

func InitableIfaceNewFromC(u unsafe.Pointer) *InitableIface {
	c := (*C.GInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &InitableIface{native: c}

	return g
}

func (recv *InitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputVector is a wrapper around the C record GInputVector.
type InputVector struct {
	native *C.GInputVector
	// buffer : no type generator for gpointer, gpointer
	Size uint64
}

func InputVectorNewFromC(u unsafe.Pointer) *InputVector {
	c := (*C.GInputVector)(u)
	if c == nil {
		return nil
	}

	g := &InputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *InputVector) ToC() unsafe.Pointer {
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// OutputVector is a wrapper around the C record GOutputVector.
type OutputVector struct {
	native *C.GOutputVector
	// buffer : no type generator for gpointer, gconstpointer
	Size uint64
}

func OutputVectorNewFromC(u unsafe.Pointer) *OutputVector {
	c := (*C.GOutputVector)(u)
	if c == nil {
		return nil
	}

	g := &OutputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *OutputVector) ToC() unsafe.Pointer {
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// SrvTargetNew is a wrapper around the C function g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	c_priority := (C.guint16)(priority)

	c_weight := (C.guint16)(weight)

	retC := C.g_srv_target_new(c_hostname, c_port, c_priority, c_weight)
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}
