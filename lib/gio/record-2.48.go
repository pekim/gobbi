// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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

// Provides an interface for socket-like objects which have datagram semantics,
// following the Berkeley sockets API. The interface methods are thin wrappers
// around the corresponding virtual methods, and no pre-processing of inputs is
// implemented — so implementations of this API must handle all functionality
// documented in the interface methods.
/*

C record/class : GDatagramBasedInterface
*/
type DatagramBasedInterface struct {
	native *C.GDatagramBasedInterface
	// g_iface : record
	// no type for receive_messages
	// no type for send_messages
	// no type for create_source
	// no type for condition_check
	// no type for condition_wait
}

func DatagramBasedInterfaceNewFromC(u unsafe.Pointer) *DatagramBasedInterface {
	c := (*C.GDatagramBasedInterface)(u)
	if c == nil {
		return nil
	}

	g := &DatagramBasedInterface{native: c}

	return g
}

func (recv *DatagramBasedInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// vtable for a #GDtlsClientConnection implementation.
/*

C record/class : GDtlsClientConnectionInterface
*/
type DtlsClientConnectionInterface struct {
	native *C.GDtlsClientConnectionInterface
	// g_iface : record
}

func DtlsClientConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsClientConnectionInterface {
	c := (*C.GDtlsClientConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsClientConnectionInterface{native: c}

	return g
}

func (recv *DtlsClientConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Virtual method table for a #GDtlsConnection implementation.
/*

C record/class : GDtlsConnectionInterface
*/
type DtlsConnectionInterface struct {
	native *C.GDtlsConnectionInterface
	// g_iface : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// no type for shutdown
	// no type for shutdown_async
	// no type for shutdown_finish
}

func DtlsConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsConnectionInterface {
	c := (*C.GDtlsConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsConnectionInterface{native: c}

	return g
}

func (recv *DtlsConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// vtable for a #GDtlsServerConnection implementation.
/*

C record/class : GDtlsServerConnectionInterface
*/
type DtlsServerConnectionInterface struct {
	native *C.GDtlsServerConnectionInterface
	// g_iface : record
}

func DtlsServerConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsServerConnectionInterface {
	c := (*C.GDtlsServerConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsServerConnectionInterface{native: c}

	return g
}

func (recv *DtlsServerConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Structure used for scatter/gather data input when receiving multiple
// messages or packets in one go. You generally pass in an array of empty
// #GInputVectors and the operation will use all the buffers as if they
// were one buffer, and will set @bytes_received to the total number of bytes
// received across all #GInputVectors.
//
// This structure closely mirrors `struct mmsghdr` and `struct msghdr` from
// the POSIX sockets API (see `man 2 recvmmsg`).
//
// If @address is non-%NULL then it is set to the source address the message
// was received from, and the caller must free it afterwards.
//
// If @control_messages is non-%NULL then it is set to an array of control
// messages received with the message (if any), and the caller must free it
// afterwards. @num_control_messages is set to the number of elements in
// this array, which may be zero.
//
// Flags relevant to this message will be returned in @flags. For example,
// `MSG_EOR` or `MSG_TRUNC`.
/*

C record/class : GInputMessage
*/
type InputMessage struct {
	native *C.GInputMessage
	// address : record
	// no type for vectors
	NumVectors    uint32
	BytesReceived uint64
	Flags         int32
	// no type for control_messages
	// num_control_messages : guint* with indirection level of 1
}

func InputMessageNewFromC(u unsafe.Pointer) *InputMessage {
	c := (*C.GInputMessage)(u)
	if c == nil {
		return nil
	}

	g := &InputMessage{
		BytesReceived: (uint64)(c.bytes_received),
		Flags:         (int32)(c.flags),
		NumVectors:    (uint32)(c.num_vectors),
		native:        c,
	}

	return g
}

func (recv *InputMessage) ToC() unsafe.Pointer {
	recv.native.num_vectors =
		(C.guint)(recv.NumVectors)
	recv.native.bytes_received =
		(C.gsize)(recv.BytesReceived)
	recv.native.flags =
		(C.gint)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}
