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

// A #GDatagramBased is a networking interface for representing datagram-based
// communications. It is a more or less direct mapping of the core parts of the
// BSD socket API in a portable GObject interface. It is implemented by
// #GSocket, which wraps the UNIX socket API on UNIX and winsock2 on Windows.
//
// #GDatagramBased is entirely platform independent, and is intended to be used
// alongside higher-level networking APIs such as #GIOStream.
//
// It uses vectored scatter/gather I/O by default, allowing for many messages
// to be sent or received in a single call. Where possible, implementations of
// the interface should take advantage of vectored I/O to minimise processing
// or system calls. For example, #GSocket uses recvmmsg() and sendmmsg() where
// possible. Callers should take advantage of scatter/gather I/O (the use of
// multiple buffers per message) to avoid unnecessary copying of data to
// assemble or disassemble a message.
//
// Each #GDatagramBased operation has a timeout parameter which may be negative
// for blocking behaviour, zero for non-blocking behaviour, or positive for
// timeout behaviour. A blocking operation blocks until finished or there is an
// error. A non-blocking operation will return immediately with a
// %G_IO_ERROR_WOULD_BLOCK error if it cannot make progress. A timeout operation
// will block until the operation is complete or the timeout expires; if the
// timeout expires it will return what progress it made, or
// %G_IO_ERROR_TIMED_OUT if no progress was made. To know when a call would
// successfully run you can call g_datagram_based_condition_check() or
// g_datagram_based_condition_wait(). You can also use
// g_datagram_based_create_source() and attach it to a #GMainContext to get
// callbacks when I/O is possible.
//
// When running a non-blocking operation applications should always be able to
// handle getting a %G_IO_ERROR_WOULD_BLOCK error even when some other function
// said that I/O was possible. This can easily happen in case of a race
// condition in the application, but it can also happen for other reasons. For
// instance, on Windows a socket is always seen as writable until a write
// returns %G_IO_ERROR_WOULD_BLOCK.
//
// As with #GSocket, #GDatagramBaseds can be either connection oriented (for
// example, SCTP) or connectionless (for example, UDP). #GDatagramBaseds must be
// datagram-based, not stream-based. The interface does not cover connection
// establishment — use methods on the underlying type to establish a connection
// before sending and receiving data through the #GDatagramBased API. For
// connectionless socket types the target/source address is specified or
// received in each I/O operation.
//
// Like most other APIs in GLib, #GDatagramBased is not inherently thread safe.
// To use a #GDatagramBased concurrently from multiple threads, you must
// implement your own locking.
/*

C type

GDatagramBased
*/
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

// Checks on the readiness of @datagram_based to perform operations. The
// operations specified in @condition are checked for and masked against the
// currently-satisfied conditions on @datagram_based. The result is returned.
//
// %G_IO_IN will be set in the return value if data is available to read with
// g_datagram_based_receive_messages(), or if the connection is closed remotely
// (EOS); and if the datagram_based has not been closed locally using some
// implementation-specific method (such as g_socket_close() or
// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with @shutdown_read set, if it’s a #GSocket, for
// example), all calls to this function will return %G_IO_ERROR_CLOSED.
//
// %G_IO_OUT will be set if it is expected that at least one byte can be sent
// using g_datagram_based_send_messages() without blocking. It will not be set
// if the datagram_based has been closed locally.
//
// %G_IO_HUP will be set if the connection has been closed locally.
//
// %G_IO_ERR will be set if there was an asynchronous error in transmitting data
// previously enqueued using g_datagram_based_send_messages().
//
// Note that on Windows, it is possible for an operation to return
// %G_IO_ERROR_WOULD_BLOCK even immediately after
// g_datagram_based_condition_check() has claimed that the #GDatagramBased is
// ready for writing. Rather than calling g_datagram_based_condition_check() and
// then writing to the #GDatagramBased if it succeeds, it is generally better to
// simply try writing right away, and try again later if the initial attempt
// returns %G_IO_ERROR_WOULD_BLOCK.
//
// It is meaningless to specify %G_IO_ERR or %G_IO_HUP in @condition; these
// conditions will always be set in the output if they are true. Apart from
// these flags, the output is guaranteed to be masked by @condition.
//
// This call never blocks.
/*

C function

g_datagram_based_condition_check
*/
func (recv *DatagramBased) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_datagram_based_condition_check((*C.GDatagramBased)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// Waits for up to @timeout microseconds for condition to become true on
// @datagram_based. If the condition is met, %TRUE is returned.
//
// If @cancellable is cancelled before the condition is met, or if @timeout is
// reached before the condition is met, then %FALSE is returned and @error is
// set appropriately (%G_IO_ERROR_CANCELLED or %G_IO_ERROR_TIMED_OUT).
/*

C function

g_datagram_based_condition_wait
*/
func (recv *DatagramBased) ConditionWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_datagram_based_condition_wait((*C.GDatagramBased)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GSource that can be attached to a #GMainContext to monitor for
// the availability of the specified @condition on the #GDatagramBased. The
// #GSource keeps a reference to the @datagram_based.
//
// The callback on the source is of the #GDatagramBasedSourceFunc type.
//
// It is meaningless to specify %G_IO_ERR or %G_IO_HUP in @condition; these
// conditions will always be reported in the callback if they are true.
//
// If non-%NULL, @cancellable can be used to cancel the source, which will
// cause the source to trigger, reporting the current condition (which is
// likely 0 unless cancellation happened at the same time as a condition
// change). You can check for this in the callback using
// g_cancellable_is_cancelled().
/*

C function

g_datagram_based_create_source
*/
func (recv *DatagramBased) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_datagram_based_create_source((*C.GDatagramBased)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datagram_based_receive_messages : unsupported parameter messages :

// Unsupported : g_datagram_based_send_messages : unsupported parameter messages :

// #GDtlsClientConnection is the client-side subclass of
// #GDtlsConnection, representing a client-side DTLS connection.
/*

C type

GDtlsClientConnection
*/
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

// Gets the list of distinguished names of the Certificate Authorities
// that the server will accept certificates from. This will be set
// during the TLS handshake if the server requests a certificate.
// Otherwise, it will be %NULL.
//
// Each item in the list is a #GByteArray which contains the complete
// subject DN of the certificate authority.
/*

C function

g_dtls_client_connection_get_accepted_cas
*/
func (recv *DtlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_dtls_client_connection_get_accepted_cas((*C.GDtlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @conn's expected server identity
/*

C function

g_dtls_client_connection_get_server_identity
*/
func (recv *DtlsClientConnection) GetServerIdentity() *SocketConnectable {
	retC := C.g_dtls_client_connection_get_server_identity((*C.GDtlsClientConnection)(recv.native))
	retGo := SocketConnectableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @conn's validation flags
/*

C function

g_dtls_client_connection_get_validation_flags
*/
func (recv *DtlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_dtls_client_connection_get_validation_flags((*C.GDtlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Sets @conn's expected server identity, which is used both to tell
// servers on virtual hosts which certificate to present, and also
// to let @conn know what name to look for in the certificate when
// performing %G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
/*

C function

g_dtls_client_connection_set_server_identity
*/
func (recv *DtlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	C.g_dtls_client_connection_set_server_identity((*C.GDtlsClientConnection)(recv.native), c_identity)

	return
}

// Sets @conn's validation flags, to override the default set of
// checks performed when validating a server certificate. By default,
// %G_TLS_CERTIFICATE_VALIDATE_ALL is used.
/*

C function

g_dtls_client_connection_set_validation_flags
*/
func (recv *DtlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_dtls_client_connection_set_validation_flags((*C.GDtlsClientConnection)(recv.native), c_flags)

	return
}

// #GDtlsConnection is the base DTLS connection class type, which wraps
// a #GDatagramBased and provides DTLS encryption on top of it. Its
// subclasses, #GDtlsClientConnection and #GDtlsServerConnection,
// implement client-side and server-side DTLS, respectively.
//
// For TLS support, see #GTlsConnection.
//
// As DTLS is datagram based, #GDtlsConnection implements #GDatagramBased,
// presenting a datagram-socket-like API for the encrypted connection. This
// operates over a base datagram connection, which is also a #GDatagramBased
// (#GDtlsConnection:base-socket).
//
// To close a DTLS connection, use g_dtls_connection_close().
//
// Neither #GDtlsServerConnection or #GDtlsClientConnection set the peer address
// on their base #GDatagramBased if it is a #GSocket — it is up to the caller to
// do that if they wish. If they do not, and g_socket_close() is called on the
// base socket, the #GDtlsConnection will not raise a %G_IO_ERROR_NOT_CONNECTED
// error on further I/O.
/*

C type

GDtlsConnection
*/
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

// Close the DTLS connection. This is equivalent to calling
// g_dtls_connection_shutdown() to shut down both sides of the connection.
//
// Closing a #GDtlsConnection waits for all buffered but untransmitted data to
// be sent before it completes. It then sends a `close_notify` DTLS alert to the
// peer and may wait for a `close_notify` to be received from the peer. It does
// not close the underlying #GDtlsConnection:base-socket; that must be closed
// separately.
//
// Once @conn is closed, all other operations will return %G_IO_ERROR_CLOSED.
// Closing a #GDtlsConnection multiple times will not return an error.
//
// #GDtlsConnections will be automatically closed when the last reference is
// dropped, but you might want to call this function to make sure resources are
// released as early as possible.
//
// If @cancellable is cancelled, the #GDtlsConnection may be left
// partially-closed and any pending untransmitted data may be lost. Call
// g_dtls_connection_close() again to complete closing the #GDtlsConnection.
/*

C function

g_dtls_connection_close
*/
func (recv *DtlsConnection) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

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

// Finish an asynchronous TLS close operation. See g_dtls_connection_close()
// for more information.
/*

C function

g_dtls_connection_close_finish
*/
func (recv *DtlsConnection) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_close_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Used by #GDtlsConnection implementations to emit the
// #GDtlsConnection::accept-certificate signal.
/*

C function

g_dtls_connection_emit_accept_certificate
*/
func (recv *DtlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	c_peer_cert := (*C.GTlsCertificate)(C.NULL)
	if peerCert != nil {
		c_peer_cert = (*C.GTlsCertificate)(peerCert.ToC())
	}

	c_errors := (C.GTlsCertificateFlags)(errors)

	retC := C.g_dtls_connection_emit_accept_certificate((*C.GDtlsConnection)(recv.native), c_peer_cert, c_errors)
	retGo := retC == C.TRUE

	return retGo
}

// Gets @conn's certificate, as set by
// g_dtls_connection_set_certificate().
/*

C function

g_dtls_connection_get_certificate
*/
func (recv *DtlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the certificate database that @conn uses to verify
// peer certificates. See g_dtls_connection_set_database().
/*

C function

g_dtls_connection_get_database
*/
func (recv *DtlsConnection) GetDatabase() *TlsDatabase {
	retC := C.g_dtls_connection_get_database((*C.GDtlsConnection)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Get the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords. If %NULL is returned, then
// no user interaction will occur for this connection.
/*

C function

g_dtls_connection_get_interaction
*/
func (recv *DtlsConnection) GetInteraction() *TlsInteraction {
	retC := C.g_dtls_connection_get_interaction((*C.GDtlsConnection)(recv.native))
	retGo := TlsInteractionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @conn's peer's certificate after the handshake has completed.
// (It is not set during the emission of
// #GDtlsConnection::accept-certificate.)
/*

C function

g_dtls_connection_get_peer_certificate
*/
func (recv *DtlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_peer_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the errors associated with validating @conn's peer's
// certificate, after the handshake has completed. (It is not set
// during the emission of #GDtlsConnection::accept-certificate.)
/*

C function

g_dtls_connection_get_peer_certificate_errors
*/
func (recv *DtlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_dtls_connection_get_peer_certificate_errors((*C.GDtlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Gets @conn rehandshaking mode. See
// g_dtls_connection_set_rehandshake_mode() for details.
/*

C function

g_dtls_connection_get_rehandshake_mode
*/
func (recv *DtlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_dtls_connection_get_rehandshake_mode((*C.GDtlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// Tests whether or not @conn expects a proper TLS close notification
// when the connection is closed. See
// g_dtls_connection_set_require_close_notify() for details.
/*

C function

g_dtls_connection_get_require_close_notify
*/
func (recv *DtlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_dtls_connection_get_require_close_notify((*C.GDtlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method;
// although the connection needs to perform a handshake after
// connecting (or after sending a "STARTTLS"-type command) and may
// need to rehandshake later if the server requests it,
// #GDtlsConnection will handle this for you automatically when you try
// to send or receive data on the connection. However, you can call
// g_dtls_connection_handshake() manually if you want to know for sure
// whether the initial handshake succeeded or failed (as opposed to
// just immediately trying to write to @conn, in which
// case if it fails, it may not be possible to tell if it failed
// before or after completing the handshake).
//
// Likewise, on the server side, although a handshake is necessary at
// the beginning of the communication, you do not need to call this
// function explicitly unless you want clearer error reporting.
// However, you may call g_dtls_connection_handshake() later on to
// renegotiate parameters (encryption methods, etc) with the client.
//
// #GDtlsConnection::accept_certificate may be emitted during the
// handshake.
/*

C function

g_dtls_connection_handshake
*/
func (recv *DtlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

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

// Finish an asynchronous TLS handshake operation. See
// g_dtls_connection_handshake() for more information.
/*

C function

g_dtls_connection_handshake_finish
*/
func (recv *DtlsConnection) HandshakeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_handshake_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This sets the certificate that @conn will present to its peer
// during the TLS handshake. For a #GDtlsServerConnection, it is
// mandatory to set this, and that will normally be done at construct
// time.
//
// For a #GDtlsClientConnection, this is optional. If a handshake fails
// with %G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server
// requires a certificate, and if you try connecting again, you should
// call this method first. You can call
// g_dtls_client_connection_get_accepted_cas() on the failed connection
// to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with
// or without a certificate; in that case, if you don't provide a
// certificate, you can tell that the server requested one by the fact
// that g_dtls_client_connection_get_accepted_cas() will return
// non-%NULL.)
/*

C function

g_dtls_connection_set_certificate
*/
func (recv *DtlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	C.g_dtls_connection_set_certificate((*C.GDtlsConnection)(recv.native), c_certificate)

	return
}

// Sets the certificate database that is used to verify peer certificates.
// This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to %NULL, then
// peer certificate validation will always set the
// %G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// #GDtlsConnection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// #GDtlsClientConnection:validation-flags).
/*

C function

g_dtls_connection_set_database
*/
func (recv *DtlsConnection) SetDatabase(database *TlsDatabase) {
	c_database := (*C.GTlsDatabase)(C.NULL)
	if database != nil {
		c_database = (*C.GTlsDatabase)(database.ToC())
	}

	C.g_dtls_connection_set_database((*C.GDtlsConnection)(recv.native), c_database)

	return
}

// Set the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of
// #GTlsInteraction. %NULL can also be provided if no user interaction
// should occur for this connection.
/*

C function

g_dtls_connection_set_interaction
*/
func (recv *DtlsConnection) SetInteraction(interaction *TlsInteraction) {
	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	C.g_dtls_connection_set_interaction((*C.GDtlsConnection)(recv.native), c_interaction)

	return
}

// Sets how @conn behaves with respect to rehandshaking requests.
//
// %G_TLS_REHANDSHAKE_NEVER means that it will never agree to
// rehandshake after the initial handshake is complete. (For a client,
// this means it will refuse rehandshake requests from the server, and
// for a server, this means it will close the connection with an error
// if the client attempts to rehandshake.)
//
// %G_TLS_REHANDSHAKE_SAFELY means that the connection will allow a
// rehandshake only if the other end of the connection supports the
// TLS `renegotiation_info` extension. This is the default behavior,
// but means that rehandshaking will not work against older
// implementations that do not support that extension.
//
// %G_TLS_REHANDSHAKE_UNSAFELY means that the connection will allow
// rehandshaking even without the `renegotiation_info` extension. On
// the server side in particular, this is not recommended, since it
// leaves the server open to certain attacks. However, this mode is
// necessary if you need to allow renegotiation with older client
// software.
/*

C function

g_dtls_connection_set_rehandshake_mode
*/
func (recv *DtlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_dtls_connection_set_rehandshake_mode((*C.GDtlsConnection)(recv.native), c_mode)

	return
}

// Sets whether or not @conn expects a proper TLS close notification
// before the connection is closed. If this is %TRUE (the default),
// then @conn will expect to receive a TLS close notification from its
// peer before the connection is closed, and will return a
// %G_TLS_ERROR_EOF error if the connection is closed without proper
// notification (since this may indicate a network error, or
// man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the
// connection was closed cleanly based on application-level data
// (because the application-level data includes a length field, or is
// somehow self-delimiting); in this case, the close notify is
// redundant and may be omitted. You
// can use g_dtls_connection_set_require_close_notify() to tell @conn
// to allow an "unannounced" connection close, in which case the close
// will show up as a 0-length read, as in a non-TLS
// #GDatagramBased, and it is up to the application to check that
// the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the
// connection; when the application calls g_dtls_connection_close_async() on
// @conn itself, this will send a close notification regardless of the
// setting of this property. If you explicitly want to do an unclean
// close, you can close @conn's #GDtlsConnection:base-socket rather
// than closing @conn itself.
/*

C function

g_dtls_connection_set_require_close_notify
*/
func (recv *DtlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_dtls_connection_set_require_close_notify((*C.GDtlsConnection)(recv.native), c_require_close_notify)

	return
}

// Shut down part or all of a DTLS connection.
//
// If @shutdown_read is %TRUE then the receiving side of the connection is shut
// down, and further reading is disallowed. Subsequent calls to
// g_datagram_based_receive_messages() will return %G_IO_ERROR_CLOSED.
//
// If @shutdown_write is %TRUE then the sending side of the connection is shut
// down, and further writing is disallowed. Subsequent calls to
// g_datagram_based_send_messages() will return %G_IO_ERROR_CLOSED.
//
// It is allowed for both @shutdown_read and @shutdown_write to be TRUE — this
// is equivalent to calling g_dtls_connection_close().
//
// If @cancellable is cancelled, the #GDtlsConnection may be left
// partially-closed and any pending untransmitted data may be lost. Call
// g_dtls_connection_shutdown() again to complete closing the #GDtlsConnection.
/*

C function

g_dtls_connection_shutdown
*/
func (recv *DtlsConnection) Shutdown(shutdownRead bool, shutdownWrite bool, cancellable *Cancellable) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

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

// Finish an asynchronous TLS shutdown operation. See
// g_dtls_connection_shutdown() for more information.
/*

C function

g_dtls_connection_shutdown_finish
*/
func (recv *DtlsConnection) ShutdownFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_shutdown_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GDtlsServerConnection is the server-side subclass of #GDtlsConnection,
// representing a server-side DTLS connection.
/*

C type

GDtlsServerConnection
*/
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

// Format a #GSocketConnectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the #GSocketConnectable implementation does not support string formatting,
// the implementation’s type name will be returned as a fallback.
/*

C function

g_socket_connectable_to_string
*/
func (recv *SocketConnectable) ToString() string {
	retC := C.g_socket_connectable_to_string((*C.GSocketConnectable)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GType of @backend’s #GDtlsClientConnection implementation.
/*

C function

g_tls_backend_get_dtls_client_connection_type
*/
func (recv *TlsBackend) GetDtlsClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Gets the #GType of @backend’s #GDtlsServerConnection implementation.
/*

C function

g_tls_backend_get_dtls_server_connection_type
*/
func (recv *TlsBackend) GetDtlsServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Checks if DTLS is supported. DTLS support may not be available even if TLS
// support is available, and vice-versa.
/*

C function

g_tls_backend_supports_dtls
*/
func (recv *TlsBackend) SupportsDtls() bool {
	retC := C.g_tls_backend_supports_dtls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
