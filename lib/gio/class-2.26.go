// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Credentials is a wrapper around the C record GCredentials.
type Credentials struct {
	native *C.GCredentials
}

func CredentialsNewFromC(u unsafe.Pointer) *Credentials {
	c := (*C.GCredentials)(u)
	if c == nil {
		return nil
	}

	g := &Credentials{native: c}

	return g
}

func (recv *Credentials) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CredentialsNew is a wrapper around the C function g_credentials_new.
func CredentialsNew() *Credentials {
	retC := C.g_credentials_new()
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNative is a wrapper around the C function g_credentials_get_native.
func (recv *Credentials) GetNative(nativeType CredentialsType) uintptr {
	c_native_type := (C.GCredentialsType)(nativeType)

	retC := C.g_credentials_get_native((*C.GCredentials)(recv.native), c_native_type)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_credentials_get_unix_user : no return generator

// IsSameUser is a wrapper around the C function g_credentials_is_same_user.
func (recv *Credentials) IsSameUser(otherCredentials *Credentials) (bool, error) {
	c_other_credentials := (*C.GCredentials)(otherCredentials.ToC())

	var cThrowableError *C.GError

	retC := C.g_credentials_is_same_user((*C.GCredentials)(recv.native), c_other_credentials, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetNative is a wrapper around the C function g_credentials_set_native.
func (recv *Credentials) SetNative(nativeType CredentialsType, native uintptr) {
	c_native_type := (C.GCredentialsType)(nativeType)

	c_native := (C.gpointer)(native)

	C.g_credentials_set_native((*C.GCredentials)(recv.native), c_native_type, c_native)

	return
}

// Unsupported : g_credentials_set_unix_user : unsupported parameter uid : no type generator for guint, uid_t

// ToString is a wrapper around the C function g_credentials_to_string.
func (recv *Credentials) ToString() string {
	retC := C.g_credentials_to_string((*C.GCredentials)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// DBusAuthObserver is a wrapper around the C record GDBusAuthObserver.
type DBusAuthObserver struct {
	native *C.GDBusAuthObserver
}

func DBusAuthObserverNewFromC(u unsafe.Pointer) *DBusAuthObserver {
	c := (*C.GDBusAuthObserver)(u)
	if c == nil {
		return nil
	}

	g := &DBusAuthObserver{native: c}

	return g
}

func (recv *DBusAuthObserver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusAuthObserverNew is a wrapper around the C function g_dbus_auth_observer_new.
func DBusAuthObserverNew() *DBusAuthObserver {
	retC := C.g_dbus_auth_observer_new()
	retGo := DBusAuthObserverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AuthorizeAuthenticatedPeer is a wrapper around the C function g_dbus_auth_observer_authorize_authenticated_peer.
func (recv *DBusAuthObserver) AuthorizeAuthenticatedPeer(stream *IOStream, credentials *Credentials) bool {
	c_stream := (*C.GIOStream)(stream.ToC())

	c_credentials := (*C.GCredentials)(credentials.ToC())

	retC := C.g_dbus_auth_observer_authorize_authenticated_peer((*C.GDBusAuthObserver)(recv.native), c_stream, c_credentials)
	retGo := retC == C.TRUE

	return retGo
}

// DBusConnection is a wrapper around the C record GDBusConnection.
type DBusConnection struct {
	native *C.GDBusConnection
}

func DBusConnectionNewFromC(u unsafe.Pointer) *DBusConnection {
	c := (*C.GDBusConnection)(u)
	if c == nil {
		return nil
	}

	g := &DBusConnection{native: c}

	return g
}

func (recv *DBusConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// DBusConnectionNewForAddressSync is a wrapper around the C function g_dbus_connection_new_for_address_sync.
func DBusConnectionNewForAddressSync(address string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(observer.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_for_address_sync(c_address, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// DBusConnectionNewSync is a wrapper around the C function g_dbus_connection_new_sync.
func DBusConnectionNewSync(stream *IOStream, guid string, flags DBusConnectionFlags, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusConnection, error) {
	c_stream := (*C.GIOStream)(stream.ToC())

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(observer.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_new_sync(c_stream, c_guid, c_flags, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction, GDBusMessageFilterFunction

// Unsupported : g_dbus_connection_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_close : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_close_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// CloseSync is a wrapper around the C function g_dbus_connection_close_sync.
func (recv *DBusConnection) CloseSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_close_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_connection_emit_signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_flush : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_flush_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// FlushSync is a wrapper around the C function g_dbus_connection_flush_sync.
func (recv *DBusConnection) FlushSync(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_flush_sync((*C.GDBusConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetCapabilities is a wrapper around the C function g_dbus_connection_get_capabilities.
func (recv *DBusConnection) GetCapabilities() DBusCapabilityFlags {
	retC := C.g_dbus_connection_get_capabilities((*C.GDBusConnection)(recv.native))
	retGo := (DBusCapabilityFlags)(retC)

	return retGo
}

// GetExitOnClose is a wrapper around the C function g_dbus_connection_get_exit_on_close.
func (recv *DBusConnection) GetExitOnClose() bool {
	retC := C.g_dbus_connection_get_exit_on_close((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGuid is a wrapper around the C function g_dbus_connection_get_guid.
func (recv *DBusConnection) GetGuid() string {
	retC := C.g_dbus_connection_get_guid((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPeerCredentials is a wrapper around the C function g_dbus_connection_get_peer_credentials.
func (recv *DBusConnection) GetPeerCredentials() *Credentials {
	retC := C.g_dbus_connection_get_peer_credentials((*C.GDBusConnection)(recv.native))
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStream is a wrapper around the C function g_dbus_connection_get_stream.
func (recv *DBusConnection) GetStream() *IOStream {
	retC := C.g_dbus_connection_get_stream((*C.GDBusConnection)(recv.native))
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUniqueName is a wrapper around the C function g_dbus_connection_get_unique_name.
func (recv *DBusConnection) GetUniqueName() string {
	retC := C.g_dbus_connection_get_unique_name((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsClosed is a wrapper around the C function g_dbus_connection_is_closed.
func (recv *DBusConnection) IsClosed() bool {
	retC := C.g_dbus_connection_is_closed((*C.GDBusConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// RemoveFilter is a wrapper around the C function g_dbus_connection_remove_filter.
func (recv *DBusConnection) RemoveFilter(filterId uint32) {
	c_filter_id := (C.guint)(filterId)

	C.g_dbus_connection_remove_filter((*C.GDBusConnection)(recv.native), c_filter_id)

	return
}

// SendMessage is a wrapper around the C function g_dbus_connection_send_message.
func (recv *DBusConnection) SendMessage(message *DBusMessage, flags DBusSendMessageFlags) (bool, *uint32, error) {
	c_message := (*C.GDBusMessage)(message.ToC())

	c_flags := (C.GDBusSendMessageFlags)(flags)

	var c_out_serial C.guint32

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message((*C.GDBusConnection)(recv.native), c_message, c_flags, &c_out_serial, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outSerial := (*uint32)(&c_out_serial)

	return retGo, outSerial, goThrowableError
}

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_dbus_connection_send_message_with_reply_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// SendMessageWithReplySync is a wrapper around the C function g_dbus_connection_send_message_with_reply_sync.
func (recv *DBusConnection) SendMessageWithReplySync(message *DBusMessage, flags DBusSendMessageFlags, timeoutMsec int32, cancellable *Cancellable) (*DBusMessage, *uint32, error) {
	c_message := (*C.GDBusMessage)(message.ToC())

	c_flags := (C.GDBusSendMessageFlags)(flags)

	c_timeout_msec := (C.gint)(timeoutMsec)

	var c_out_serial C.guint32

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_send_message_with_reply_sync((*C.GDBusConnection)(recv.native), c_message, c_flags, c_timeout_msec, &c_out_serial, c_cancellable, &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	outSerial := (*uint32)(&c_out_serial)

	return retGo, outSerial, goThrowableError
}

// SetExitOnClose is a wrapper around the C function g_dbus_connection_set_exit_on_close.
func (recv *DBusConnection) SetExitOnClose(exitOnClose bool) {
	c_exit_on_close :=
		boolToGboolean(exitOnClose)

	C.g_dbus_connection_set_exit_on_close((*C.GDBusConnection)(recv.native), c_exit_on_close)

	return
}

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter callback : no type generator for DBusSignalCallback, GDBusSignalCallback

// SignalUnsubscribe is a wrapper around the C function g_dbus_connection_signal_unsubscribe.
func (recv *DBusConnection) SignalUnsubscribe(subscriptionId uint32) {
	c_subscription_id := (C.guint)(subscriptionId)

	C.g_dbus_connection_signal_unsubscribe((*C.GDBusConnection)(recv.native), c_subscription_id)

	return
}

// StartMessageProcessing is a wrapper around the C function g_dbus_connection_start_message_processing.
func (recv *DBusConnection) StartMessageProcessing() {
	C.g_dbus_connection_start_message_processing((*C.GDBusConnection)(recv.native))

	return
}

// UnregisterObject is a wrapper around the C function g_dbus_connection_unregister_object.
func (recv *DBusConnection) UnregisterObject(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_object((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

// UnregisterSubtree is a wrapper around the C function g_dbus_connection_unregister_subtree.
func (recv *DBusConnection) UnregisterSubtree(registrationId uint32) bool {
	c_registration_id := (C.guint)(registrationId)

	retC := C.g_dbus_connection_unregister_subtree((*C.GDBusConnection)(recv.native), c_registration_id)
	retGo := retC == C.TRUE

	return retGo
}

// DBusMessage is a wrapper around the C record GDBusMessage.
type DBusMessage struct {
	native *C.GDBusMessage
}

func DBusMessageNewFromC(u unsafe.Pointer) *DBusMessage {
	c := (*C.GDBusMessage)(u)
	if c == nil {
		return nil
	}

	g := &DBusMessage{native: c}

	return g
}

func (recv *DBusMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMessageNew is a wrapper around the C function g_dbus_message_new.
func DBusMessageNew() *DBusMessage {
	retC := C.g_dbus_message_new()
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// DBusMessageNewMethodCall is a wrapper around the C function g_dbus_message_new_method_call.
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) *DBusMessage {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_method := C.CString(method)
	defer C.free(unsafe.Pointer(c_method))

	retC := C.g_dbus_message_new_method_call(c_name, c_path, c_interface_, c_method)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DBusMessageNewSignal is a wrapper around the C function g_dbus_message_new_signal.
func DBusMessageNewSignal(path string, interface_ string, signal string) *DBusMessage {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := C.CString(interface_)
	defer C.free(unsafe.Pointer(c_interface_))

	c_signal := C.CString(signal)
	defer C.free(unsafe.Pointer(c_signal))

	retC := C.g_dbus_message_new_signal(c_path, c_interface_, c_signal)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_dbus_message_copy.
func (recv *DBusMessage) Copy() (*DBusMessage, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_copy((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetArg0 is a wrapper around the C function g_dbus_message_get_arg0.
func (recv *DBusMessage) GetArg0() string {
	retC := C.g_dbus_message_get_arg0((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_body : return type : Blacklisted record : GVariant

// GetByteOrder is a wrapper around the C function g_dbus_message_get_byte_order.
func (recv *DBusMessage) GetByteOrder() DBusMessageByteOrder {
	retC := C.g_dbus_message_get_byte_order((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageByteOrder)(retC)

	return retGo
}

// GetDestination is a wrapper around the C function g_dbus_message_get_destination.
func (recv *DBusMessage) GetDestination() string {
	retC := C.g_dbus_message_get_destination((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetErrorName is a wrapper around the C function g_dbus_message_get_error_name.
func (recv *DBusMessage) GetErrorName() string {
	retC := C.g_dbus_message_get_error_name((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_message_get_flags.
func (recv *DBusMessage) GetFlags() DBusMessageFlags {
	retC := C.g_dbus_message_get_flags((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageFlags)(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_header : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_message_get_header_fields : no return type

// GetInterface is a wrapper around the C function g_dbus_message_get_interface.
func (recv *DBusMessage) GetInterface() string {
	retC := C.g_dbus_message_get_interface((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLocked is a wrapper around the C function g_dbus_message_get_locked.
func (recv *DBusMessage) GetLocked() bool {
	retC := C.g_dbus_message_get_locked((*C.GDBusMessage)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMember is a wrapper around the C function g_dbus_message_get_member.
func (recv *DBusMessage) GetMember() string {
	retC := C.g_dbus_message_get_member((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMessageType is a wrapper around the C function g_dbus_message_get_message_type.
func (recv *DBusMessage) GetMessageType() DBusMessageType {
	retC := C.g_dbus_message_get_message_type((*C.GDBusMessage)(recv.native))
	retGo := (DBusMessageType)(retC)

	return retGo
}

// GetNumUnixFds is a wrapper around the C function g_dbus_message_get_num_unix_fds.
func (recv *DBusMessage) GetNumUnixFds() uint32 {
	retC := C.g_dbus_message_get_num_unix_fds((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetPath is a wrapper around the C function g_dbus_message_get_path.
func (recv *DBusMessage) GetPath() string {
	retC := C.g_dbus_message_get_path((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetReplySerial is a wrapper around the C function g_dbus_message_get_reply_serial.
func (recv *DBusMessage) GetReplySerial() uint32 {
	retC := C.g_dbus_message_get_reply_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSender is a wrapper around the C function g_dbus_message_get_sender.
func (recv *DBusMessage) GetSender() string {
	retC := C.g_dbus_message_get_sender((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSerial is a wrapper around the C function g_dbus_message_get_serial.
func (recv *DBusMessage) GetSerial() uint32 {
	retC := C.g_dbus_message_get_serial((*C.GDBusMessage)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSignature is a wrapper around the C function g_dbus_message_get_signature.
func (recv *DBusMessage) GetSignature() string {
	retC := C.g_dbus_message_get_signature((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUnixFdList is a wrapper around the C function g_dbus_message_get_unix_fd_list.
func (recv *DBusMessage) GetUnixFdList() *UnixFDList {
	retC := C.g_dbus_message_get_unix_fd_list((*C.GDBusMessage)(recv.native))
	retGo := UnixFDListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lock is a wrapper around the C function g_dbus_message_lock.
func (recv *DBusMessage) Lock() {
	C.g_dbus_message_lock((*C.GDBusMessage)(recv.native))

	return
}

// Unsupported : g_dbus_message_new_method_error : unsupported parameter ... : varargs

// NewMethodErrorLiteral is a wrapper around the C function g_dbus_message_new_method_error_literal.
func (recv *DBusMessage) NewMethodErrorLiteral(errorName string, errorMessage string) *DBusMessage {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	retC := C.g_dbus_message_new_method_error_literal((*C.GDBusMessage)(recv.native), c_error_name, c_error_message)
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// NewMethodReply is a wrapper around the C function g_dbus_message_new_method_reply.
func (recv *DBusMessage) NewMethodReply() *DBusMessage {
	retC := C.g_dbus_message_new_method_reply((*C.GDBusMessage)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Print is a wrapper around the C function g_dbus_message_print.
func (recv *DBusMessage) Print(indent uint32) string {
	c_indent := (C.guint)(indent)

	retC := C.g_dbus_message_print((*C.GDBusMessage)(recv.native), c_indent)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_set_body : unsupported parameter body : Blacklisted record : GVariant

// SetByteOrder is a wrapper around the C function g_dbus_message_set_byte_order.
func (recv *DBusMessage) SetByteOrder(byteOrder DBusMessageByteOrder) {
	c_byte_order := (C.GDBusMessageByteOrder)(byteOrder)

	C.g_dbus_message_set_byte_order((*C.GDBusMessage)(recv.native), c_byte_order)

	return
}

// SetDestination is a wrapper around the C function g_dbus_message_set_destination.
func (recv *DBusMessage) SetDestination(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_destination((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetErrorName is a wrapper around the C function g_dbus_message_set_error_name.
func (recv *DBusMessage) SetErrorName(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_error_name((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetFlags is a wrapper around the C function g_dbus_message_set_flags.
func (recv *DBusMessage) SetFlags(flags DBusMessageFlags) {
	c_flags := (C.GDBusMessageFlags)(flags)

	C.g_dbus_message_set_flags((*C.GDBusMessage)(recv.native), c_flags)

	return
}

// Unsupported : g_dbus_message_set_header : unsupported parameter value : Blacklisted record : GVariant

// SetInterface is a wrapper around the C function g_dbus_message_set_interface.
func (recv *DBusMessage) SetInterface(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_interface((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetMember is a wrapper around the C function g_dbus_message_set_member.
func (recv *DBusMessage) SetMember(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_member((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetMessageType is a wrapper around the C function g_dbus_message_set_message_type.
func (recv *DBusMessage) SetMessageType(type_ DBusMessageType) {
	c_type := (C.GDBusMessageType)(type_)

	C.g_dbus_message_set_message_type((*C.GDBusMessage)(recv.native), c_type)

	return
}

// SetNumUnixFds is a wrapper around the C function g_dbus_message_set_num_unix_fds.
func (recv *DBusMessage) SetNumUnixFds(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_num_unix_fds((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetPath is a wrapper around the C function g_dbus_message_set_path.
func (recv *DBusMessage) SetPath(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_path((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetReplySerial is a wrapper around the C function g_dbus_message_set_reply_serial.
func (recv *DBusMessage) SetReplySerial(value uint32) {
	c_value := (C.guint32)(value)

	C.g_dbus_message_set_reply_serial((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetSender is a wrapper around the C function g_dbus_message_set_sender.
func (recv *DBusMessage) SetSender(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_sender((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetSerial is a wrapper around the C function g_dbus_message_set_serial.
func (recv *DBusMessage) SetSerial(serial uint32) {
	c_serial := (C.guint32)(serial)

	C.g_dbus_message_set_serial((*C.GDBusMessage)(recv.native), c_serial)

	return
}

// SetSignature is a wrapper around the C function g_dbus_message_set_signature.
func (recv *DBusMessage) SetSignature(value string) {
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_signature((*C.GDBusMessage)(recv.native), c_value)

	return
}

// SetUnixFdList is a wrapper around the C function g_dbus_message_set_unix_fd_list.
func (recv *DBusMessage) SetUnixFdList(fdList *UnixFDList) {
	c_fd_list := (*C.GUnixFDList)(fdList.ToC())

	C.g_dbus_message_set_unix_fd_list((*C.GDBusMessage)(recv.native), c_fd_list)

	return
}

// Unsupported : g_dbus_message_to_blob : no return type

// ToGerror is a wrapper around the C function g_dbus_message_to_gerror.
func (recv *DBusMessage) ToGerror() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_dbus_message_to_gerror((*C.GDBusMessage)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// DBusMethodInvocation is a wrapper around the C record GDBusMethodInvocation.
type DBusMethodInvocation struct {
	native *C.GDBusMethodInvocation
}

func DBusMethodInvocationNewFromC(u unsafe.Pointer) *DBusMethodInvocation {
	c := (*C.GDBusMethodInvocation)(u)
	if c == nil {
		return nil
	}

	g := &DBusMethodInvocation{native: c}

	return g
}

func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GetConnection is a wrapper around the C function g_dbus_method_invocation_get_connection.
func (recv *DBusMethodInvocation) GetConnection() *DBusConnection {
	retC := C.g_dbus_method_invocation_get_connection((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInterfaceName is a wrapper around the C function g_dbus_method_invocation_get_interface_name.
func (recv *DBusMethodInvocation) GetInterfaceName() string {
	retC := C.g_dbus_method_invocation_get_interface_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMessage is a wrapper around the C function g_dbus_method_invocation_get_message.
func (recv *DBusMethodInvocation) GetMessage() *DBusMessage {
	retC := C.g_dbus_method_invocation_get_message((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMethodInfo is a wrapper around the C function g_dbus_method_invocation_get_method_info.
func (recv *DBusMethodInvocation) GetMethodInfo() *DBusMethodInfo {
	retC := C.g_dbus_method_invocation_get_method_info((*C.GDBusMethodInvocation)(recv.native))
	retGo := DBusMethodInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMethodName is a wrapper around the C function g_dbus_method_invocation_get_method_name.
func (recv *DBusMethodInvocation) GetMethodName() string {
	retC := C.g_dbus_method_invocation_get_method_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_method_invocation_get_object_path.
func (recv *DBusMethodInvocation) GetObjectPath() string {
	retC := C.g_dbus_method_invocation_get_object_path((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_method_invocation_get_parameters : return type : Blacklisted record : GVariant

// GetSender is a wrapper around the C function g_dbus_method_invocation_get_sender.
func (recv *DBusMethodInvocation) GetSender() string {
	retC := C.g_dbus_method_invocation_get_sender((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUserData is a wrapper around the C function g_dbus_method_invocation_get_user_data.
func (recv *DBusMethodInvocation) GetUserData() uintptr {
	retC := C.g_dbus_method_invocation_get_user_data((*C.GDBusMethodInvocation)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// ReturnDbusError is a wrapper around the C function g_dbus_method_invocation_return_dbus_error.
func (recv *DBusMethodInvocation) ReturnDbusError(errorName string, errorMessage string) {
	c_error_name := C.CString(errorName)
	defer C.free(unsafe.Pointer(c_error_name))

	c_error_message := C.CString(errorMessage)
	defer C.free(unsafe.Pointer(c_error_message))

	C.g_dbus_method_invocation_return_dbus_error((*C.GDBusMethodInvocation)(recv.native), c_error_name, c_error_message)

	return
}

// Unsupported : g_dbus_method_invocation_return_error : unsupported parameter ... : varargs

// ReturnErrorLiteral is a wrapper around the C function g_dbus_method_invocation_return_error_literal.
func (recv *DBusMethodInvocation) ReturnErrorLiteral(domain glib.Quark, code int32, message string) {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	C.g_dbus_method_invocation_return_error_literal((*C.GDBusMethodInvocation)(recv.native), c_domain, c_code, c_message)

	return
}

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// ReturnGerror is a wrapper around the C function g_dbus_method_invocation_return_gerror.
func (recv *DBusMethodInvocation) ReturnGerror(error *glib.Error) {
	c_error := (*C.GError)(error.ToC())

	C.g_dbus_method_invocation_return_gerror((*C.GDBusMethodInvocation)(recv.native), c_error)

	return
}

// Unsupported : g_dbus_method_invocation_return_value : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_object_manager_client_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc, GDBusProxyTypeFunc

// DBusProxy is a wrapper around the C record GDBusProxy.
type DBusProxy struct {
	native *C.GDBusProxy
	// Private : parent_instance
	// Private : priv
}

func DBusProxyNewFromC(u unsafe.Pointer) *DBusProxy {
	c := (*C.GDBusProxy)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxy{native: c}

	return g
}

func (recv *DBusProxy) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// DBusProxyNewForBusSync is a wrapper around the C function g_dbus_proxy_new_for_bus_sync.
func DBusProxyNewForBusSync(busType BusType, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_bus_type := (C.GBusType)(busType)

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(info.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_for_bus_sync(c_bus_type, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// DBusProxyNewSync is a wrapper around the C function g_dbus_proxy_new_sync.
func DBusProxyNewSync(connection *DBusConnection, flags DBusProxyFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable *Cancellable) (*DBusProxy, error) {
	c_connection := (*C.GDBusConnection)(connection.ToC())

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(info.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_interface_name := C.CString(interfaceName)
	defer C.free(unsafe.Pointer(c_interface_name))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_proxy_new_sync(c_connection, c_flags, c_info, c_name, c_object_path, c_interface_name, c_cancellable, &cThrowableError)
	retGo := DBusProxyNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_dbus_proxy_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property_names : no return type

// GetConnection is a wrapper around the C function g_dbus_proxy_get_connection.
func (recv *DBusProxy) GetConnection() *DBusConnection {
	retC := C.g_dbus_proxy_get_connection((*C.GDBusProxy)(recv.native))
	retGo := DBusConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultTimeout is a wrapper around the C function g_dbus_proxy_get_default_timeout.
func (recv *DBusProxy) GetDefaultTimeout() int32 {
	retC := C.g_dbus_proxy_get_default_timeout((*C.GDBusProxy)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_proxy_get_flags.
func (recv *DBusProxy) GetFlags() DBusProxyFlags {
	retC := C.g_dbus_proxy_get_flags((*C.GDBusProxy)(recv.native))
	retGo := (DBusProxyFlags)(retC)

	return retGo
}

// GetInterfaceInfo is a wrapper around the C function g_dbus_proxy_get_interface_info.
func (recv *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo {
	retC := C.g_dbus_proxy_get_interface_info((*C.GDBusProxy)(recv.native))
	retGo := DBusInterfaceInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInterfaceName is a wrapper around the C function g_dbus_proxy_get_interface_name.
func (recv *DBusProxy) GetInterfaceName() string {
	retC := C.g_dbus_proxy_get_interface_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_dbus_proxy_get_name.
func (recv *DBusProxy) GetName() string {
	retC := C.g_dbus_proxy_get_name((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNameOwner is a wrapper around the C function g_dbus_proxy_get_name_owner.
func (recv *DBusProxy) GetNameOwner() string {
	retC := C.g_dbus_proxy_get_name_owner((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetObjectPath is a wrapper around the C function g_dbus_proxy_get_object_path.
func (recv *DBusProxy) GetObjectPath() string {
	retC := C.g_dbus_proxy_get_object_path((*C.GDBusProxy)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_proxy_set_cached_property : unsupported parameter value : Blacklisted record : GVariant

// SetDefaultTimeout is a wrapper around the C function g_dbus_proxy_set_default_timeout.
func (recv *DBusProxy) SetDefaultTimeout(timeoutMsec int32) {
	c_timeout_msec := (C.gint)(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout((*C.GDBusProxy)(recv.native), c_timeout_msec)

	return
}

// SetInterfaceInfo is a wrapper around the C function g_dbus_proxy_set_interface_info.
func (recv *DBusProxy) SetInterfaceInfo(info *DBusInterfaceInfo) {
	c_info := (*C.GDBusInterfaceInfo)(info.ToC())

	C.g_dbus_proxy_set_interface_info((*C.GDBusProxy)(recv.native), c_info)

	return
}

// DBusServer is a wrapper around the C record GDBusServer.
type DBusServer struct {
	native *C.GDBusServer
}

func DBusServerNewFromC(u unsafe.Pointer) *DBusServer {
	c := (*C.GDBusServer)(u)
	if c == nil {
		return nil
	}

	g := &DBusServer{native: c}

	return g
}

func (recv *DBusServer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusServerNewSync is a wrapper around the C function g_dbus_server_new_sync.
func DBusServerNewSync(address string, flags DBusServerFlags, guid string, observer *DBusAuthObserver, cancellable *Cancellable) (*DBusServer, error) {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusServerFlags)(flags)

	c_guid := C.CString(guid)
	defer C.free(unsafe.Pointer(c_guid))

	c_observer := (*C.GDBusAuthObserver)(observer.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_server_new_sync(c_address, c_flags, c_guid, c_observer, c_cancellable, &cThrowableError)
	retGo := DBusServerNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetClientAddress is a wrapper around the C function g_dbus_server_get_client_address.
func (recv *DBusServer) GetClientAddress() string {
	retC := C.g_dbus_server_get_client_address((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_dbus_server_get_flags.
func (recv *DBusServer) GetFlags() DBusServerFlags {
	retC := C.g_dbus_server_get_flags((*C.GDBusServer)(recv.native))
	retGo := (DBusServerFlags)(retC)

	return retGo
}

// GetGuid is a wrapper around the C function g_dbus_server_get_guid.
func (recv *DBusServer) GetGuid() string {
	retC := C.g_dbus_server_get_guid((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// IsActive is a wrapper around the C function g_dbus_server_is_active.
func (recv *DBusServer) IsActive() bool {
	retC := C.g_dbus_server_is_active((*C.GDBusServer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Start is a wrapper around the C function g_dbus_server_start.
func (recv *DBusServer) Start() {
	C.g_dbus_server_start((*C.GDBusServer)(recv.native))

	return
}

// Stop is a wrapper around the C function g_dbus_server_stop.
func (recv *DBusServer) Stop() {
	C.g_dbus_server_stop((*C.GDBusServer)(recv.native))

	return
}

// ReadUpto is a wrapper around the C function g_data_input_stream_read_upto.
func (recv *DataInputStream) ReadUpto(stopChars string, stopCharsLen int64, cancellable *Cancellable) (string, *uint64, error) {
	c_stop_chars := C.CString(stopChars)
	defer C.free(unsafe.Pointer(c_stop_chars))

	c_stop_chars_len := (C.gssize)(stopCharsLen)

	var c_length C.gsize

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_upto((*C.GDataInputStream)(recv.native), c_stop_chars, c_stop_chars_len, &c_length, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (*uint64)(&c_length)

	return retGo, length, goThrowableError
}

// Unsupported : g_data_input_stream_read_upto_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// StealData is a wrapper around the C function g_memory_output_stream_steal_data.
func (recv *MemoryOutputStream) StealData() uintptr {
	retC := C.g_memory_output_stream_steal_data((*C.GMemoryOutputStream)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetScheme is a wrapper around the C function g_network_address_get_scheme.
func (recv *NetworkAddress) GetScheme() string {
	retC := C.g_network_address_get_scheme((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetScheme is a wrapper around the C function g_network_service_get_scheme.
func (recv *NetworkService) GetScheme() string {
	retC := C.g_network_service_get_scheme((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetScheme is a wrapper around the C function g_network_service_set_scheme.
func (recv *NetworkService) SetScheme(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.g_network_service_set_scheme((*C.GNetworkService)(recv.native), c_scheme)

	return
}

// Acquire is a wrapper around the C function g_permission_acquire.
func (recv *Permission) Acquire(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_acquire((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_permission_acquire_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_permission_acquire_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetAllowed is a wrapper around the C function g_permission_get_allowed.
func (recv *Permission) GetAllowed() bool {
	retC := C.g_permission_get_allowed((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanAcquire is a wrapper around the C function g_permission_get_can_acquire.
func (recv *Permission) GetCanAcquire() bool {
	retC := C.g_permission_get_can_acquire((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCanRelease is a wrapper around the C function g_permission_get_can_release.
func (recv *Permission) GetCanRelease() bool {
	retC := C.g_permission_get_can_release((*C.GPermission)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ImplUpdate is a wrapper around the C function g_permission_impl_update.
func (recv *Permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	c_allowed :=
		boolToGboolean(allowed)

	c_can_acquire :=
		boolToGboolean(canAcquire)

	c_can_release :=
		boolToGboolean(canRelease)

	C.g_permission_impl_update((*C.GPermission)(recv.native), c_allowed, c_can_acquire, c_can_release)

	return
}

// Release is a wrapper around the C function g_permission_release.
func (recv *Permission) Release(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_permission_release((*C.GPermission)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_permission_release_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_permission_release_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : record
	// Private : priv
}

func ProxyAddressNewFromC(u unsafe.Pointer) *ProxyAddress {
	c := (*C.GProxyAddress)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddress{native: c}

	return g
}

func (recv *ProxyAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressNew is a wrapper around the C function g_proxy_address_new.
func ProxyAddressNew(inetaddr *InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) *SocketAddress {
	c_inetaddr := (*C.GInetAddress)(inetaddr.ToC())

	c_port := (C.guint16)(port)

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_dest_hostname := C.CString(destHostname)
	defer C.free(unsafe.Pointer(c_dest_hostname))

	c_dest_port := (C.guint16)(destPort)

	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	retC := C.g_proxy_address_new(c_inetaddr, c_port, c_protocol, c_dest_hostname, c_dest_port, c_username, c_password)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDestinationHostname is a wrapper around the C function g_proxy_address_get_destination_hostname.
func (recv *ProxyAddress) GetDestinationHostname() string {
	retC := C.g_proxy_address_get_destination_hostname((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDestinationPort is a wrapper around the C function g_proxy_address_get_destination_port.
func (recv *ProxyAddress) GetDestinationPort() uint16 {
	retC := C.g_proxy_address_get_destination_port((*C.GProxyAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetPassword is a wrapper around the C function g_proxy_address_get_password.
func (recv *ProxyAddress) GetPassword() string {
	retC := C.g_proxy_address_get_password((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProtocol is a wrapper around the C function g_proxy_address_get_protocol.
func (recv *ProxyAddress) GetProtocol() string {
	retC := C.g_proxy_address_get_protocol((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUsername is a wrapper around the C function g_proxy_address_get_username.
func (recv *ProxyAddress) GetUsername() string {
	retC := C.g_proxy_address_get_username((*C.GProxyAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SettingsNew is a wrapper around the C function g_settings_new.
func SettingsNew(schemaId string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	retC := C.g_settings_new(c_schema_id)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SettingsNewWithBackend is a wrapper around the C function g_settings_new_with_backend.
func SettingsNewWithBackend(schemaId string, backend *SettingsBackend) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(backend.ToC())

	retC := C.g_settings_new_with_backend(c_schema_id, c_backend)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SettingsNewWithBackendAndPath is a wrapper around the C function g_settings_new_with_backend_and_path.
func SettingsNewWithBackendAndPath(schemaId string, backend *SettingsBackend, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_backend := (*C.GSettingsBackend)(backend.ToC())

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_backend_and_path(c_schema_id, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SettingsNewWithPath is a wrapper around the C function g_settings_new_with_path.
func SettingsNewWithPath(schemaId string, path string) *Settings {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_with_path(c_schema_id, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Bind is a wrapper around the C function g_settings_bind.
func (recv *Settings) Bind(key string, object uintptr, property string, flags SettingsBindFlags) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_flags := (C.GSettingsBindFlags)(flags)

	C.g_settings_bind((*C.GSettings)(recv.native), c_key, c_object, c_property, c_flags)

	return
}

// Unsupported : g_settings_bind_with_mapping : unsupported parameter get_mapping : no type generator for SettingsBindGetMapping, GSettingsBindGetMapping

// BindWritable is a wrapper around the C function g_settings_bind_writable.
func (recv *Settings) BindWritable(key string, object uintptr, property string, inverted bool) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_inverted :=
		boolToGboolean(inverted)

	C.g_settings_bind_writable((*C.GSettings)(recv.native), c_key, c_object, c_property, c_inverted)

	return
}

// Delay is a wrapper around the C function g_settings_delay.
func (recv *Settings) Delay() {
	C.g_settings_delay((*C.GSettings)(recv.native))

	return
}

// Unsupported : g_settings_get : unsupported parameter ... : varargs

// GetBoolean is a wrapper around the C function g_settings_get_boolean.
func (recv *Settings) GetBoolean(key string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_boolean((*C.GSettings)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// GetChild is a wrapper around the C function g_settings_get_child.
func (recv *Settings) GetChild(name string) *Settings {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_get_child((*C.GSettings)(recv.native), c_name)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDouble is a wrapper around the C function g_settings_get_double.
func (recv *Settings) GetDouble(key string) float64 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_double((*C.GSettings)(recv.native), c_key)
	retGo := (float64)(retC)

	return retGo
}

// GetEnum is a wrapper around the C function g_settings_get_enum.
func (recv *Settings) GetEnum(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_enum((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_settings_get_flags.
func (recv *Settings) GetFlags(key string) uint32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_flags((*C.GSettings)(recv.native), c_key)
	retGo := (uint32)(retC)

	return retGo
}

// GetHasUnapplied is a wrapper around the C function g_settings_get_has_unapplied.
func (recv *Settings) GetHasUnapplied() bool {
	retC := C.g_settings_get_has_unapplied((*C.GSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInt is a wrapper around the C function g_settings_get_int.
func (recv *Settings) GetInt(key string) int32 {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_int((*C.GSettings)(recv.native), c_key)
	retGo := (int32)(retC)

	return retGo
}

// GetString is a wrapper around the C function g_settings_get_string.
func (recv *Settings) GetString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_string((*C.GSettings)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_settings_get_strv : no return type

// Unsupported : g_settings_get_value : return type : Blacklisted record : GVariant

// IsWritable is a wrapper around the C function g_settings_is_writable.
func (recv *Settings) IsWritable(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_is_writable((*C.GSettings)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_set : unsupported parameter ... : varargs

// SetBoolean is a wrapper around the C function g_settings_set_boolean.
func (recv *Settings) SetBoolean(key string, value bool) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	retC := C.g_settings_set_boolean((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetDouble is a wrapper around the C function g_settings_set_double.
func (recv *Settings) SetDouble(key string, value float64) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	retC := C.g_settings_set_double((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetInt is a wrapper around the C function g_settings_set_int.
func (recv *Settings) SetInt(key string, value int32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	retC := C.g_settings_set_int((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetString is a wrapper around the C function g_settings_set_string.
func (recv *Settings) SetString(key string, value string) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	retC := C.g_settings_set_string((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_settings_set_strv : unsupported parameter value : no param type

// Unsupported : g_settings_set_value : unsupported parameter value : Blacklisted record : GVariant

// Blacklisted : g_settings_backend_changed

// Blacklisted : g_settings_backend_changed_tree

// Unsupported : g_settings_backend_keys_changed : unsupported parameter items : no param type

// Blacklisted : g_settings_backend_path_changed

// Blacklisted : g_settings_backend_path_writable_changed

// Blacklisted : g_settings_backend_writable_changed

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// SimplePermissionNew is a wrapper around the C function g_simple_permission_new.
func SimplePermissionNew(allowed bool) *Permission {
	c_allowed :=
		boolToGboolean(allowed)

	retC := C.g_simple_permission_new(c_allowed)
	retGo := PermissionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCredentials is a wrapper around the C function g_socket_get_credentials.
func (recv *Socket) GetCredentials() (*Credentials, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_credentials((*C.GSocket)(recv.native), &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetTimeout is a wrapper around the C function g_socket_get_timeout.
func (recv *Socket) GetTimeout() uint32 {
	retC := C.g_socket_get_timeout((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_socket_receive_with_blocking : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_with_blocking : unsupported parameter buffer : no param type

// SetTimeout is a wrapper around the C function g_socket_set_timeout.
func (recv *Socket) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_set_timeout((*C.GSocket)(recv.native), c_timeout)

	return
}

// ConnectToUri is a wrapper around the C function g_socket_client_connect_to_uri.
func (recv *SocketClient) ConnectToUri(uri string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_uri((*C.GSocketClient)(recv.native), c_uri, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_uri_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_uri_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetEnableProxy is a wrapper around the C function g_socket_client_get_enable_proxy.
func (recv *SocketClient) GetEnableProxy() bool {
	retC := C.g_socket_client_get_enable_proxy((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTimeout is a wrapper around the C function g_socket_client_get_timeout.
func (recv *SocketClient) GetTimeout() uint32 {
	retC := C.g_socket_client_get_timeout((*C.GSocketClient)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// SetEnableProxy is a wrapper around the C function g_socket_client_set_enable_proxy.
func (recv *SocketClient) SetEnableProxy(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.g_socket_client_set_enable_proxy((*C.GSocketClient)(recv.native), c_enable)

	return
}

// SetTimeout is a wrapper around the C function g_socket_client_set_timeout.
func (recv *SocketClient) SetTimeout(timeout uint32) {
	c_timeout := (C.guint)(timeout)

	C.g_socket_client_set_timeout((*C.GSocketClient)(recv.native), c_timeout)

	return
}

// Unsupported : g_subprocess_new : unsupported parameter error : record with indirection level of 2

// Unsupported : g_subprocess_newv : unsupported parameter argv : no param type

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// ReceiveCredentials is a wrapper around the C function g_unix_connection_receive_credentials.
func (recv *UnixConnection) ReceiveCredentials(cancellable *Cancellable) (*Credentials, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SendCredentials is a wrapper around the C function g_unix_connection_send_credentials.
func (recv *UnixConnection) SendCredentials(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_credentials((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : record
	// priv : record
}

func UnixCredentialsMessageNewFromC(u unsafe.Pointer) *UnixCredentialsMessage {
	c := (*C.GUnixCredentialsMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessage{native: c}

	return g
}

func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessageNew is a wrapper around the C function g_unix_credentials_message_new.
func UnixCredentialsMessageNew() *SocketControlMessage {
	retC := C.g_unix_credentials_message_new()
	retGo := SocketControlMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnixCredentialsMessageNewWithCredentials is a wrapper around the C function g_unix_credentials_message_new_with_credentials.
func UnixCredentialsMessageNewWithCredentials(credentials *Credentials) *SocketControlMessage {
	c_credentials := (*C.GCredentials)(credentials.ToC())

	retC := C.g_unix_credentials_message_new_with_credentials(c_credentials)
	retGo := SocketControlMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCredentials is a wrapper around the C function g_unix_credentials_message_get_credentials.
func (recv *UnixCredentialsMessage) GetCredentials() *Credentials {
	retC := C.g_unix_credentials_message_get_credentials((*C.GUnixCredentialsMessage)(recv.native))
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// GetAddressType is a wrapper around the C function g_unix_socket_address_get_address_type.
func (recv *UnixSocketAddress) GetAddressType() UnixSocketAddressType {
	retC := C.g_unix_socket_address_get_address_type((*C.GUnixSocketAddress)(recv.native))
	retGo := (UnixSocketAddressType)(retC)

	return retGo
}

// GetFileInfo is a wrapper around the C function g_zlib_compressor_get_file_info.
func (recv *ZlibCompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_compressor_get_file_info((*C.GZlibCompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFileInfo is a wrapper around the C function g_zlib_compressor_set_file_info.
func (recv *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	c_file_info := (*C.GFileInfo)(fileInfo.ToC())

	C.g_zlib_compressor_set_file_info((*C.GZlibCompressor)(recv.native), c_file_info)

	return
}

// GetFileInfo is a wrapper around the C function g_zlib_decompressor_get_file_info.
func (recv *ZlibDecompressor) GetFileInfo() *FileInfo {
	retC := C.g_zlib_decompressor_get_file_info((*C.GZlibDecompressor)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}
