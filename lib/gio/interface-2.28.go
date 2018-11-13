// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_action_activate : unsupported parameter parameter : Blacklisted record : GVariant

// Checks if @action is currently enabled.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
/*

C function : g_action_get_enabled
*/
func (recv *Action) GetEnabled() bool {
	retC := C.g_action_get_enabled((*C.GAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Queries the name of @action.
/*

C function : g_action_get_name
*/
func (recv *Action) GetName() string {
	retC := C.g_action_get_name((*C.GAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_action_get_parameter_type : return type : Blacklisted record : GVariantType

// Unsupported : g_action_get_state : return type : Blacklisted record : GVariant

// Unsupported : g_action_get_state_hint : return type : Blacklisted record : GVariant

// Unsupported : g_action_get_state_type : return type : Blacklisted record : GVariantType

// Unsupported signal 'action-added' for ActionGroup : unsupported parameter action_name : type utf8 :

// Unsupported signal 'action-enabled-changed' for ActionGroup : unsupported parameter action_name : type utf8 :

// Unsupported signal 'action-removed' for ActionGroup : unsupported parameter action_name : type utf8 :

// Unsupported signal 'action-state-changed' for ActionGroup : unsupported parameter action_name : type utf8 :

// Emits the #GActionGroup::action-added signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
/*

C function : g_action_group_action_added
*/
func (recv *ActionGroup) ActionAdded(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_added((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// Emits the #GActionGroup::action-enabled-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
/*

C function : g_action_group_action_enabled_changed
*/
func (recv *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_enabled :=
		boolToGboolean(enabled)

	C.g_action_group_action_enabled_changed((*C.GActionGroup)(recv.native), c_action_name, c_enabled)

	return
}

// Emits the #GActionGroup::action-removed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
/*

C function : g_action_group_action_removed
*/
func (recv *ActionGroup) ActionRemoved(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_removed((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// Unsupported : g_action_group_action_state_changed : unsupported parameter state : Blacklisted record : GVariant

// Unsupported : g_action_group_activate_action : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : g_action_group_change_action_state : unsupported parameter value : Blacklisted record : GVariant

// Checks if the named action within @action_group is currently enabled.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
/*

C function : g_action_group_get_action_enabled
*/
func (recv *ActionGroup) GetActionEnabled(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_enabled((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_action_group_get_action_parameter_type : return type : Blacklisted record : GVariantType

// Unsupported : g_action_group_get_action_state : return type : Blacklisted record : GVariant

// Unsupported : g_action_group_get_action_state_hint : return type : Blacklisted record : GVariant

// Unsupported : g_action_group_get_action_state_type : return type : Blacklisted record : GVariantType

// Checks if the named action exists within @action_group.
/*

C function : g_action_group_has_action
*/
func (recv *ActionGroup) HasAction(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_has_action((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_action_group_list_actions : no return type

// PollableInputStream is a wrapper around the C record GPollableInputStream.
type PollableInputStream struct {
	native *C.GPollableInputStream
}

func PollableInputStreamNewFromC(u unsafe.Pointer) *PollableInputStream {
	c := (*C.GPollableInputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStream{native: c}

	return g
}

func (recv *PollableInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Checks if @stream is actually pollable. Some classes may implement
// #GPollableInputStream but have only certain instances of that class
// be pollable. If this method returns %FALSE, then the behavior of
// other #GPollableInputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant;
// a stream cannot switch from pollable to non-pollable or vice versa.
/*

C function : g_pollable_input_stream_can_poll
*/
func (recv *PollableInputStream) CanPoll() bool {
	retC := C.g_pollable_input_stream_can_poll((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Creates a #GSource that triggers when @stream can be read, or
// @cancellable is triggered or an error occurs. The callback on the
// source is of the #GPollableSourceFunc type.
//
// As with g_pollable_input_stream_is_readable(), it is possible that
// the stream may not actually be readable even after the source
// triggers, so you should use g_pollable_input_stream_read_nonblocking()
// rather than g_input_stream_read() from the callback.
/*

C function : g_pollable_input_stream_create_source
*/
func (recv *PollableInputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_input_stream_create_source((*C.GPollableInputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if @stream can be read.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_input_stream_read()
// after this returns %TRUE would still block. To guarantee
// non-blocking behavior, you should always use
// g_pollable_input_stream_read_nonblocking(), which will return a
// %G_IO_ERROR_WOULD_BLOCK error rather than blocking.
/*

C function : g_pollable_input_stream_is_readable
*/
func (recv *PollableInputStream) IsReadable() bool {
	retC := C.g_pollable_input_stream_is_readable((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Attempts to read up to @count bytes from @stream into @buffer, as
// with g_input_stream_read(). If @stream is not currently readable,
// this will immediately return %G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_input_stream_create_source() to create a #GSource
// that will be triggered when @stream is readable.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
/*

C function : g_pollable_input_stream_read_nonblocking
*/
func (recv *PollableInputStream) ReadNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_input_stream_read_nonblocking((*C.GPollableInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PollableOutputStream is a wrapper around the C record GPollableOutputStream.
type PollableOutputStream struct {
	native *C.GPollableOutputStream
}

func PollableOutputStreamNewFromC(u unsafe.Pointer) *PollableOutputStream {
	c := (*C.GPollableOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStream{native: c}

	return g
}

func (recv *PollableOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Checks if @stream is actually pollable. Some classes may implement
// #GPollableOutputStream but have only certain instances of that
// class be pollable. If this method returns %FALSE, then the behavior
// of other #GPollableOutputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant;
// a stream cannot switch from pollable to non-pollable or vice versa.
/*

C function : g_pollable_output_stream_can_poll
*/
func (recv *PollableOutputStream) CanPoll() bool {
	retC := C.g_pollable_output_stream_can_poll((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Creates a #GSource that triggers when @stream can be written, or
// @cancellable is triggered or an error occurs. The callback on the
// source is of the #GPollableSourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that
// the stream may not actually be writable even after the source
// triggers, so you should use g_pollable_output_stream_write_nonblocking()
// rather than g_output_stream_write() from the callback.
/*

C function : g_pollable_output_stream_create_source
*/
func (recv *PollableOutputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_output_stream_create_source((*C.GPollableOutputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_output_stream_write()
// after this returns %TRUE would still block. To guarantee
// non-blocking behavior, you should always use
// g_pollable_output_stream_write_nonblocking(), which will return a
// %G_IO_ERROR_WOULD_BLOCK error rather than blocking.
/*

C function : g_pollable_output_stream_is_writable
*/
func (recv *PollableOutputStream) IsWritable() bool {
	retC := C.g_pollable_output_stream_is_writable((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Attempts to write up to @count bytes from @buffer to @stream, as
// with g_output_stream_write(). If @stream is not currently writable,
// this will immediately return %G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource
// that will be triggered when @stream is writable.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
//
// Also note that if %G_IO_ERROR_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you send the same @buffer and @count.
/*

C function : g_pollable_output_stream_write_nonblocking
*/
func (recv *PollableOutputStream) WriteNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_pollable_output_stream_write_nonblocking((*C.GPollableOutputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// TlsBackend is a wrapper around the C record GTlsBackend.
type TlsBackend struct {
	native *C.GTlsBackend
}

func TlsBackendNewFromC(u unsafe.Pointer) *TlsBackend {
	c := (*C.GTlsBackend)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackend{native: c}

	return g
}

func (recv *TlsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gets the #GType of @backend's #GTlsCertificate implementation.
/*

C function : g_tls_backend_get_certificate_type
*/
func (recv *TlsBackend) GetCertificateType() gobject.Type {
	retC := C.g_tls_backend_get_certificate_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Gets the #GType of @backend's #GTlsClientConnection implementation.
/*

C function : g_tls_backend_get_client_connection_type
*/
func (recv *TlsBackend) GetClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Gets the #GType of @backend's #GTlsServerConnection implementation.
/*

C function : g_tls_backend_get_server_connection_type
*/
func (recv *TlsBackend) GetServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Checks if TLS is supported; if this returns %FALSE for the default
// #GTlsBackend, it means no "real" TLS backend is available.
/*

C function : g_tls_backend_supports_tls
*/
func (recv *TlsBackend) SupportsTls() bool {
	retC := C.g_tls_backend_supports_tls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TlsClientConnection is a wrapper around the C record GTlsClientConnection.
type TlsClientConnection struct {
	native *C.GTlsClientConnection
}

func TlsClientConnectionNewFromC(u unsafe.Pointer) *TlsClientConnection {
	c := (*C.GTlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnection{native: c}

	return g
}

func (recv *TlsClientConnection) ToC() unsafe.Pointer {

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

C function : g_tls_client_connection_get_accepted_cas
*/
func (recv *TlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_tls_client_connection_get_accepted_cas((*C.GTlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @conn's expected server identity
/*

C function : g_tls_client_connection_get_server_identity
*/
func (recv *TlsClientConnection) GetServerIdentity() *SocketConnectable {
	retC := C.g_tls_client_connection_get_server_identity((*C.GTlsClientConnection)(recv.native))
	retGo := SocketConnectableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets whether @conn will force the lowest-supported TLS protocol
// version rather than attempt to negotiate the highest mutually-
// supported version of TLS; see g_tls_client_connection_set_use_ssl3().
/*

C function : g_tls_client_connection_get_use_ssl3
*/
func (recv *TlsClientConnection) GetUseSsl3() bool {
	retC := C.g_tls_client_connection_get_use_ssl3((*C.GTlsClientConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets @conn's validation flags
/*

C function : g_tls_client_connection_get_validation_flags
*/
func (recv *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_tls_client_connection_get_validation_flags((*C.GTlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Sets @conn's expected server identity, which is used both to tell
// servers on virtual hosts which certificate to present, and also
// to let @conn know what name to look for in the certificate when
// performing %G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
/*

C function : g_tls_client_connection_set_server_identity
*/
func (recv *TlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	C.g_tls_client_connection_set_server_identity((*C.GTlsClientConnection)(recv.native), c_identity)

	return
}

// If @use_ssl3 is %TRUE, this forces @conn to use the lowest-supported
// TLS protocol version rather than trying to properly negotiate the
// highest mutually-supported protocol version with the peer. This can
// be used when talking to broken TLS servers that exhibit protocol
// version intolerance.
//
// Be aware that SSL 3.0 is generally disabled by the #GTlsBackend, so
// the lowest-supported protocol version is probably not SSL 3.0.
/*

C function : g_tls_client_connection_set_use_ssl3
*/
func (recv *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	c_use_ssl3 :=
		boolToGboolean(useSsl3)

	C.g_tls_client_connection_set_use_ssl3((*C.GTlsClientConnection)(recv.native), c_use_ssl3)

	return
}

// Sets @conn's validation flags, to override the default set of
// checks performed when validating a server certificate. By default,
// %G_TLS_CERTIFICATE_VALIDATE_ALL is used.
/*

C function : g_tls_client_connection_set_validation_flags
*/
func (recv *TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_tls_client_connection_set_validation_flags((*C.GTlsClientConnection)(recv.native), c_flags)

	return
}

// TlsServerConnection is a wrapper around the C record GTlsServerConnection.
type TlsServerConnection struct {
	native *C.GTlsServerConnection
}

func TlsServerConnectionNewFromC(u unsafe.Pointer) *TlsServerConnection {
	c := (*C.GTlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnection{native: c}

	return g
}

func (recv *TlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
