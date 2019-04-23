// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	gboolean dtlsconnection_acceptCertificateHandler(GObject *, GTlsCertificate *, GTlsCertificateFlags, gpointer);

	static gulong DtlsConnection_signal_connect_accept_certificate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accept-certificate", G_CALLBACK(dtlsconnection_acceptCertificateHandler), data);
	}

*/
import "C"

// DatagramBased is a wrapper around the C record GDatagramBased.
type DatagramBased struct {
	native *C.GDatagramBased
}

func DatagramBasedNewFromC(u unsafe.Pointer) *DatagramBased {
	c := (*C.GDatagramBased)(u)
	if c == nil {
		return nil
	}

	g := &DatagramBased{native: c}

	return g
}

func (recv *DatagramBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DatagramBased with another DatagramBased, and returns true if they represent the same GObject.
func (recv *DatagramBased) Equals(other *DatagramBased) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_datagram_based_condition_check

// Blacklisted : g_datagram_based_condition_wait

// Blacklisted : g_datagram_based_create_source

// Unsupported : g_datagram_based_receive_messages : unsupported parameter messages :

// Unsupported : g_datagram_based_send_messages : unsupported parameter messages :

// DtlsClientConnection is a wrapper around the C record GDtlsClientConnection.
type DtlsClientConnection struct {
	native *C.GDtlsClientConnection
}

func DtlsClientConnectionNewFromC(u unsafe.Pointer) *DtlsClientConnection {
	c := (*C.GDtlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsClientConnection{native: c}

	return g
}

func (recv *DtlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DtlsClientConnection with another DtlsClientConnection, and returns true if they represent the same GObject.
func (recv *DtlsClientConnection) Equals(other *DtlsClientConnection) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dtls_client_connection_new

// Blacklisted : g_dtls_client_connection_get_accepted_cas

// Blacklisted : g_dtls_client_connection_get_server_identity

// Blacklisted : g_dtls_client_connection_get_validation_flags

// Blacklisted : g_dtls_client_connection_set_server_identity

// Blacklisted : g_dtls_client_connection_set_validation_flags

// DtlsConnection is a wrapper around the C record GDtlsConnection.
type DtlsConnection struct {
	native *C.GDtlsConnection
}

func DtlsConnectionNewFromC(u unsafe.Pointer) *DtlsConnection {
	c := (*C.GDtlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsConnection{native: c}

	return g
}

func (recv *DtlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DtlsConnection with another DtlsConnection, and returns true if they represent the same GObject.
func (recv *DtlsConnection) Equals(other *DtlsConnection) bool {
	return other.ToC() == recv.ToC()
}

type signalDtlsConnectionAcceptCertificateDetail struct {
	callback  DtlsConnectionSignalAcceptCertificateCallback
	handlerID C.gulong
}

var signalDtlsConnectionAcceptCertificateId int
var signalDtlsConnectionAcceptCertificateMap = make(map[int]signalDtlsConnectionAcceptCertificateDetail)
var signalDtlsConnectionAcceptCertificateLock sync.RWMutex

// DtlsConnectionSignalAcceptCertificateCallback is a callback function for a 'accept-certificate' signal emitted from a DtlsConnection.
type DtlsConnectionSignalAcceptCertificateCallback func(peerCert *TlsCertificate, errors TlsCertificateFlags) bool

/*
ConnectAcceptCertificate connects the callback to the 'accept-certificate' signal for the DtlsConnection.

The returned value represents the connection, and may be passed to DisconnectAcceptCertificate to remove it.
*/
func (recv *DtlsConnection) ConnectAcceptCertificate(callback DtlsConnectionSignalAcceptCertificateCallback) int {
	signalDtlsConnectionAcceptCertificateLock.Lock()
	defer signalDtlsConnectionAcceptCertificateLock.Unlock()

	signalDtlsConnectionAcceptCertificateId++
	instance := C.gpointer(recv.native)
	handlerID := C.DtlsConnection_signal_connect_accept_certificate(instance, C.gpointer(uintptr(signalDtlsConnectionAcceptCertificateId)))

	detail := signalDtlsConnectionAcceptCertificateDetail{callback, handlerID}
	signalDtlsConnectionAcceptCertificateMap[signalDtlsConnectionAcceptCertificateId] = detail

	return signalDtlsConnectionAcceptCertificateId
}

/*
DisconnectAcceptCertificate disconnects a callback from the 'accept-certificate' signal for the DtlsConnection.

The connectionID should be a value returned from a call to ConnectAcceptCertificate.
*/
func (recv *DtlsConnection) DisconnectAcceptCertificate(connectionID int) {
	signalDtlsConnectionAcceptCertificateLock.Lock()
	defer signalDtlsConnectionAcceptCertificateLock.Unlock()

	detail, exists := signalDtlsConnectionAcceptCertificateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDtlsConnectionAcceptCertificateMap, connectionID)
}

//export dtlsconnection_acceptCertificateHandler
func dtlsconnection_acceptCertificateHandler(_ *C.GObject, c_peer_cert *C.GTlsCertificate, c_errors C.GTlsCertificateFlags, data C.gpointer) C.gboolean {
	signalDtlsConnectionAcceptCertificateLock.RLock()
	defer signalDtlsConnectionAcceptCertificateLock.RUnlock()

	peerCert := TlsCertificateNewFromC(unsafe.Pointer(c_peer_cert))

	errors := TlsCertificateFlags(c_errors)

	index := int(uintptr(data))
	callback := signalDtlsConnectionAcceptCertificateMap[index].callback
	retGo := callback(peerCert, errors)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : g_dtls_connection_close

// Unsupported : g_dtls_connection_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dtls_connection_close_finish

// Blacklisted : g_dtls_connection_emit_accept_certificate

// Blacklisted : g_dtls_connection_get_certificate

// Blacklisted : g_dtls_connection_get_database

// Blacklisted : g_dtls_connection_get_interaction

// Blacklisted : g_dtls_connection_get_peer_certificate

// Blacklisted : g_dtls_connection_get_peer_certificate_errors

// Blacklisted : g_dtls_connection_get_rehandshake_mode

// Blacklisted : g_dtls_connection_get_require_close_notify

// Blacklisted : g_dtls_connection_handshake

// Unsupported : g_dtls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dtls_connection_handshake_finish

// Blacklisted : g_dtls_connection_set_certificate

// Blacklisted : g_dtls_connection_set_database

// Blacklisted : g_dtls_connection_set_interaction

// Blacklisted : g_dtls_connection_set_rehandshake_mode

// Blacklisted : g_dtls_connection_set_require_close_notify

// Blacklisted : g_dtls_connection_shutdown

// Unsupported : g_dtls_connection_shutdown_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_dtls_connection_shutdown_finish

// DtlsServerConnection is a wrapper around the C record GDtlsServerConnection.
type DtlsServerConnection struct {
	native *C.GDtlsServerConnection
}

func DtlsServerConnectionNewFromC(u unsafe.Pointer) *DtlsServerConnection {
	c := (*C.GDtlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsServerConnection{native: c}

	return g
}

func (recv *DtlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DtlsServerConnection with another DtlsServerConnection, and returns true if they represent the same GObject.
func (recv *DtlsServerConnection) Equals(other *DtlsServerConnection) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_dtls_server_connection_new

// Blacklisted : g_socket_connectable_to_string

// Blacklisted : g_tls_backend_get_dtls_client_connection_type

// Blacklisted : g_tls_backend_get_dtls_server_connection_type

// Blacklisted : g_tls_backend_supports_dtls
