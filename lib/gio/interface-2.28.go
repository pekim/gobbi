// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// GetEnabled is a wrapper around the C function g_action_get_enabled.
func (recv *Action) GetEnabled() bool {
	retC := C.g_action_get_enabled((*C.GAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function g_action_get_name.
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

// ActionAdded is a wrapper around the C function g_action_group_action_added.
func (recv *ActionGroup) ActionAdded(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_added((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// ActionEnabledChanged is a wrapper around the C function g_action_group_action_enabled_changed.
func (recv *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_enabled :=
		boolToGboolean(enabled)

	C.g_action_group_action_enabled_changed((*C.GActionGroup)(recv.native), c_action_name, c_enabled)

	return
}

// ActionRemoved is a wrapper around the C function g_action_group_action_removed.
func (recv *ActionGroup) ActionRemoved(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_removed((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// Unsupported : g_action_group_action_state_changed : unsupported parameter state : Blacklisted record : GVariant

// Unsupported : g_action_group_activate_action : unsupported parameter parameter : Blacklisted record : GVariant

// Unsupported : g_action_group_change_action_state : unsupported parameter value : Blacklisted record : GVariant

// GetActionEnabled is a wrapper around the C function g_action_group_get_action_enabled.
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

// HasAction is a wrapper around the C function g_action_group_has_action.
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

// CanPoll is a wrapper around the C function g_pollable_input_stream_can_poll.
func (recv *PollableInputStream) CanPoll() bool {
	retC := C.g_pollable_input_stream_can_poll((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_input_stream_create_source.
func (recv *PollableInputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_pollable_input_stream_create_source((*C.GPollableInputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsReadable is a wrapper around the C function g_pollable_input_stream_is_readable.
func (recv *PollableInputStream) IsReadable() bool {
	retC := C.g_pollable_input_stream_is_readable((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_pollable_input_stream_read_nonblocking : unsupported parameter buffer : no param type

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

// CanPoll is a wrapper around the C function g_pollable_output_stream_can_poll.
func (recv *PollableOutputStream) CanPoll() bool {
	retC := C.g_pollable_output_stream_can_poll((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_output_stream_create_source.
func (recv *PollableOutputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_pollable_output_stream_create_source((*C.GPollableOutputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsWritable is a wrapper around the C function g_pollable_output_stream_is_writable.
func (recv *PollableOutputStream) IsWritable() bool {
	retC := C.g_pollable_output_stream_is_writable((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_pollable_output_stream_write_nonblocking : unsupported parameter buffer : no param type

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

// Unsupported : g_tls_backend_get_certificate_type : no return generator

// Unsupported : g_tls_backend_get_client_connection_type : no return generator

// Unsupported : g_tls_backend_get_server_connection_type : no return generator

// SupportsTls is a wrapper around the C function g_tls_backend_supports_tls.
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

// GetAcceptedCas is a wrapper around the C function g_tls_client_connection_get_accepted_cas.
func (recv *TlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_tls_client_connection_get_accepted_cas((*C.GTlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_tls_client_connection_get_server_identity : no return generator

// GetUseSsl3 is a wrapper around the C function g_tls_client_connection_get_use_ssl3.
func (recv *TlsClientConnection) GetUseSsl3() bool {
	retC := C.g_tls_client_connection_get_use_ssl3((*C.GTlsClientConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValidationFlags is a wrapper around the C function g_tls_client_connection_get_validation_flags.
func (recv *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_tls_client_connection_get_validation_flags((*C.GTlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Unsupported : g_tls_client_connection_set_server_identity : unsupported parameter identity : no type generator for SocketConnectable, GSocketConnectable*

// SetUseSsl3 is a wrapper around the C function g_tls_client_connection_set_use_ssl3.
func (recv *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	c_use_ssl3 :=
		boolToGboolean(useSsl3)

	C.g_tls_client_connection_set_use_ssl3((*C.GTlsClientConnection)(recv.native), c_use_ssl3)

	return
}

// SetValidationFlags is a wrapper around the C function g_tls_client_connection_set_validation_flags.
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
