// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// ConditionCheck is a wrapper around the C function g_datagram_based_condition_check.
func (recv *DatagramBased) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_datagram_based_condition_check((*C.GDatagramBased)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// ConditionWait is a wrapper around the C function g_datagram_based_condition_wait.
func (recv *DatagramBased) ConditionWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_datagram_based_condition_wait((*C.GDatagramBased)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// CreateSource is a wrapper around the C function g_datagram_based_create_source.
func (recv *DatagramBased) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_datagram_based_create_source((*C.GDatagramBased)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetAcceptedCas is a wrapper around the C function g_dtls_client_connection_get_accepted_cas.
func (recv *DtlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_dtls_client_connection_get_accepted_cas((*C.GDtlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dtls_client_connection_get_server_identity : no return generator

// GetValidationFlags is a wrapper around the C function g_dtls_client_connection_get_validation_flags.
func (recv *DtlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_dtls_client_connection_get_validation_flags((*C.GDtlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Unsupported : g_dtls_client_connection_set_server_identity : unsupported parameter identity : no type generator for SocketConnectable (GSocketConnectable*) for param identity

// SetValidationFlags is a wrapper around the C function g_dtls_client_connection_set_validation_flags.
func (recv *DtlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_dtls_client_connection_set_validation_flags((*C.GDtlsClientConnection)(recv.native), c_flags)

	return
}

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

// Close is a wrapper around the C function g_dtls_connection_close.
func (recv *DtlsConnection) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_close((*C.GDtlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dtls_connection_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dtls_connection_close_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// EmitAcceptCertificate is a wrapper around the C function g_dtls_connection_emit_accept_certificate.
func (recv *DtlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	c_peer_cert := (*C.GTlsCertificate)(peerCert.ToC())

	c_errors := (C.GTlsCertificateFlags)(errors)

	retC := C.g_dtls_connection_emit_accept_certificate((*C.GDtlsConnection)(recv.native), c_peer_cert, c_errors)
	retGo := retC == C.TRUE

	return retGo
}

// GetCertificate is a wrapper around the C function g_dtls_connection_get_certificate.
func (recv *DtlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDatabase is a wrapper around the C function g_dtls_connection_get_database.
func (recv *DtlsConnection) GetDatabase() *TlsDatabase {
	retC := C.g_dtls_connection_get_database((*C.GDtlsConnection)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInteraction is a wrapper around the C function g_dtls_connection_get_interaction.
func (recv *DtlsConnection) GetInteraction() *TlsInteraction {
	retC := C.g_dtls_connection_get_interaction((*C.GDtlsConnection)(recv.native))
	retGo := TlsInteractionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificate is a wrapper around the C function g_dtls_connection_get_peer_certificate.
func (recv *DtlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_peer_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificateErrors is a wrapper around the C function g_dtls_connection_get_peer_certificate_errors.
func (recv *DtlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_dtls_connection_get_peer_certificate_errors((*C.GDtlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// GetRehandshakeMode is a wrapper around the C function g_dtls_connection_get_rehandshake_mode.
func (recv *DtlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_dtls_connection_get_rehandshake_mode((*C.GDtlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// GetRequireCloseNotify is a wrapper around the C function g_dtls_connection_get_require_close_notify.
func (recv *DtlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_dtls_connection_get_require_close_notify((*C.GDtlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Handshake is a wrapper around the C function g_dtls_connection_handshake.
func (recv *DtlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_handshake((*C.GDtlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dtls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dtls_connection_handshake_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

// SetCertificate is a wrapper around the C function g_dtls_connection_set_certificate.
func (recv *DtlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(certificate.ToC())

	C.g_dtls_connection_set_certificate((*C.GDtlsConnection)(recv.native), c_certificate)

	return
}

// SetDatabase is a wrapper around the C function g_dtls_connection_set_database.
func (recv *DtlsConnection) SetDatabase(database *TlsDatabase) {
	c_database := (*C.GTlsDatabase)(database.ToC())

	C.g_dtls_connection_set_database((*C.GDtlsConnection)(recv.native), c_database)

	return
}

// SetInteraction is a wrapper around the C function g_dtls_connection_set_interaction.
func (recv *DtlsConnection) SetInteraction(interaction *TlsInteraction) {
	c_interaction := (*C.GTlsInteraction)(interaction.ToC())

	C.g_dtls_connection_set_interaction((*C.GDtlsConnection)(recv.native), c_interaction)

	return
}

// SetRehandshakeMode is a wrapper around the C function g_dtls_connection_set_rehandshake_mode.
func (recv *DtlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_dtls_connection_set_rehandshake_mode((*C.GDtlsConnection)(recv.native), c_mode)

	return
}

// SetRequireCloseNotify is a wrapper around the C function g_dtls_connection_set_require_close_notify.
func (recv *DtlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_dtls_connection_set_require_close_notify((*C.GDtlsConnection)(recv.native), c_require_close_notify)

	return
}

// Shutdown is a wrapper around the C function g_dtls_connection_shutdown.
func (recv *DtlsConnection) Shutdown(shutdownRead bool, shutdownWrite bool, cancellable *Cancellable) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_shutdown((*C.GDtlsConnection)(recv.native), c_shutdown_read, c_shutdown_write, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dtls_connection_shutdown_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_dtls_connection_shutdown_finish : unsupported parameter result : no type generator for AsyncResult (GAsyncResult*) for param result

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

// ToString is a wrapper around the C function g_socket_connectable_to_string.
func (recv *SocketConnectable) ToString() string {
	retC := C.g_socket_connectable_to_string((*C.GSocketConnectable)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDtlsClientConnectionType is a wrapper around the C function g_tls_backend_get_dtls_client_connection_type.
func (recv *TlsBackend) GetDtlsClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetDtlsServerConnectionType is a wrapper around the C function g_tls_backend_get_dtls_server_connection_type.
func (recv *TlsBackend) GetDtlsServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// SupportsDtls is a wrapper around the C function g_tls_backend_supports_dtls.
func (recv *TlsBackend) SupportsDtls() bool {
	retC := C.g_tls_backend_supports_dtls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
