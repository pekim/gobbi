// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

func (recv *Credentials) toC() *C.GCredentials {

	return recv.native
}

// Unsupported : g_credentials_new : no return generator

// GetNative is a wrapper around the C function g_credentials_get_native.
func (recv *Credentials) GetNative(nativeType CredentialsType) uintptr {
	c_native_type := (C.GCredentialsType)(nativeType)

	retC := C.g_credentials_get_native((*C.GCredentials)(recv.native), c_native_type)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_credentials_get_unix_pid : no return generator

// Unsupported : g_credentials_get_unix_user : no return generator

// Unsupported : g_credentials_is_same_user : unsupported parameter other_credentials : no type generator for Credentials, GCredentials*

// Unsupported : g_credentials_set_native : no return generator

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

func (recv *DBusAuthObserver) toC() *C.GDBusAuthObserver {

	return recv.native
}

// Unsupported : g_dbus_auth_observer_new : no return generator

// Unsupported : g_dbus_auth_observer_allow_mechanism : no return generator

// Unsupported : g_dbus_auth_observer_authorize_authenticated_peer : unsupported parameter stream : no type generator for IOStream, GIOStream*

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

func (recv *DBusConnection) toC() *C.GDBusConnection {

	return recv.native
}

// Unsupported : g_dbus_connection_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_new_for_address_sync : unsupported parameter flags : no type generator for DBusConnectionFlags, GDBusConnectionFlags

// Unsupported : g_dbus_connection_new_sync : unsupported parameter stream : no type generator for IOStream, GIOStream*

// Unsupported : g_dbus_connection_add_filter : unsupported parameter filter_function : no type generator for DBusMessageFilterFunction, GDBusMessageFilterFunction

// Unsupported : g_dbus_connection_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_call_with_unix_fd_list_finish : unsupported parameter out_fd_list : no type generator for UnixFDList, GUnixFDList**

// Unsupported : g_dbus_connection_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_close : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_connection_close_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_close_sync : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_connection_emit_signal : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_connection_export_action_group : unsupported parameter action_group : no type generator for ActionGroup, GActionGroup*

// Unsupported : g_dbus_connection_export_menu_model : unsupported parameter menu : no type generator for MenuModel, GMenuModel*

// Unsupported : g_dbus_connection_flush : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_connection_flush_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_flush_sync : unsupported parameter cancellable : no type generator for Cancellable, GCancellable*

// Unsupported : g_dbus_connection_get_capabilities : no return generator

// Unsupported : g_dbus_connection_get_exit_on_close : no return generator

// GetGuid is a wrapper around the C function g_dbus_connection_get_guid.
func (recv *DBusConnection) GetGuid() string {
	retC := C.g_dbus_connection_get_guid((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_connection_get_peer_credentials : no return generator

// Unsupported : g_dbus_connection_get_stream : no return generator

// GetUniqueName is a wrapper around the C function g_dbus_connection_get_unique_name.
func (recv *DBusConnection) GetUniqueName() string {
	retC := C.g_dbus_connection_get_unique_name((*C.GDBusConnection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_connection_is_closed : no return generator

// Unsupported : g_dbus_connection_register_object : unsupported parameter user_data_free_func : no type generator for GLib.DestroyNotify, GDestroyNotify

// Unsupported : g_dbus_connection_register_subtree : unsupported parameter flags : no type generator for DBusSubtreeFlags, GDBusSubtreeFlags

// Unsupported : g_dbus_connection_remove_filter : no return generator

// Unsupported : g_dbus_connection_send_message : unsupported parameter message : no type generator for DBusMessage, GDBusMessage*

// Unsupported : g_dbus_connection_send_message_with_reply : unsupported parameter message : no type generator for DBusMessage, GDBusMessage*

// Unsupported : g_dbus_connection_send_message_with_reply_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_connection_send_message_with_reply_sync : unsupported parameter message : no type generator for DBusMessage, GDBusMessage*

// Unsupported : g_dbus_connection_set_exit_on_close : unsupported parameter exit_on_close : no type generator for gboolean, gboolean

// Unsupported : g_dbus_connection_signal_subscribe : unsupported parameter flags : no type generator for DBusSignalFlags, GDBusSignalFlags

// Unsupported : g_dbus_connection_signal_unsubscribe : no return generator

// Unsupported : g_dbus_connection_start_message_processing : no return generator

// Unsupported : g_dbus_connection_unexport_action_group : no return generator

// Unsupported : g_dbus_connection_unexport_menu_model : no return generator

// Unsupported : g_dbus_connection_unregister_object : no return generator

// Unsupported : g_dbus_connection_unregister_subtree : no return generator

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

func (recv *DBusMessage) toC() *C.GDBusMessage {

	return recv.native
}

// Unsupported : g_dbus_message_new : no return generator

// Unsupported : g_dbus_message_new_from_blob : unsupported parameter blob : no param type

// Unsupported : g_dbus_message_new_method_call : no return generator

// Unsupported : g_dbus_message_new_signal : no return generator

// Unsupported : g_dbus_message_copy : no return generator

// GetArg0 is a wrapper around the C function g_dbus_message_get_arg0.
func (recv *DBusMessage) GetArg0() string {
	retC := C.g_dbus_message_get_arg0((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_body : return type : Blacklisted record : GVariant

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

// Unsupported : g_dbus_message_get_flags : no return generator

// Unsupported : g_dbus_message_get_header : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_message_get_header_fields : no return type

// GetInterface is a wrapper around the C function g_dbus_message_get_interface.
func (recv *DBusMessage) GetInterface() string {
	retC := C.g_dbus_message_get_interface((*C.GDBusMessage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_message_get_locked : no return generator

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

// Unsupported : g_dbus_message_get_unix_fd_list : no return generator

// Unsupported : g_dbus_message_lock : no return generator

// Unsupported : g_dbus_message_new_method_error : unsupported parameter ... : varargs

// Unsupported : g_dbus_message_new_method_error_literal : no return generator

// Unsupported : g_dbus_message_new_method_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_dbus_message_new_method_reply : no return generator

// Print is a wrapper around the C function g_dbus_message_print.
func (recv *DBusMessage) Print(indent uint32) string {
	c_indent := (C.guint)(indent)

	retC := C.g_dbus_message_print((*C.GDBusMessage)(recv.native), c_indent)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_dbus_message_set_body : unsupported parameter body : Blacklisted record : GVariant

// Unsupported : g_dbus_message_set_byte_order : no return generator

// Unsupported : g_dbus_message_set_destination : no return generator

// Unsupported : g_dbus_message_set_error_name : no return generator

// Unsupported : g_dbus_message_set_flags : unsupported parameter flags : no type generator for DBusMessageFlags, GDBusMessageFlags

// Unsupported : g_dbus_message_set_header : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_dbus_message_set_interface : no return generator

// Unsupported : g_dbus_message_set_member : no return generator

// Unsupported : g_dbus_message_set_message_type : no return generator

// Unsupported : g_dbus_message_set_num_unix_fds : no return generator

// Unsupported : g_dbus_message_set_path : no return generator

// Unsupported : g_dbus_message_set_reply_serial : no return generator

// Unsupported : g_dbus_message_set_sender : no return generator

// Unsupported : g_dbus_message_set_serial : no return generator

// Unsupported : g_dbus_message_set_signature : no return generator

// Unsupported : g_dbus_message_set_unix_fd_list : unsupported parameter fd_list : no type generator for UnixFDList, GUnixFDList*

// Unsupported : g_dbus_message_to_blob : unsupported parameter out_size : no type generator for gsize, gsize*

// Unsupported : g_dbus_message_to_gerror : no return generator

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

func (recv *DBusMethodInvocation) toC() *C.GDBusMethodInvocation {

	return recv.native
}

// Unsupported : g_dbus_method_invocation_get_connection : no return generator

// GetInterfaceName is a wrapper around the C function g_dbus_method_invocation_get_interface_name.
func (recv *DBusMethodInvocation) GetInterfaceName() string {
	retC := C.g_dbus_method_invocation_get_interface_name((*C.GDBusMethodInvocation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_method_invocation_get_message : no return generator

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
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_dbus_method_invocation_return_dbus_error : no return generator

// Unsupported : g_dbus_method_invocation_return_error : unsupported parameter ... : varargs

// Unsupported : g_dbus_method_invocation_return_error_literal : no return generator

// Unsupported : g_dbus_method_invocation_return_error_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : g_dbus_method_invocation_return_gerror : no return generator

// Unsupported : g_dbus_method_invocation_return_value : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_method_invocation_return_value_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_method_invocation_take_error : no return generator

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

func (recv *DBusProxy) toC() *C.GDBusProxy {

	return recv.native
}

// Unsupported : g_dbus_proxy_new_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_new_for_bus_sync : unsupported parameter flags : no type generator for DBusProxyFlags, GDBusProxyFlags

// Unsupported : g_dbus_proxy_new_sync : unsupported parameter connection : no type generator for DBusConnection, GDBusConnection*

// Unsupported : g_dbus_proxy_call : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_finish : unsupported parameter res : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_dbus_proxy_call_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_finish : unsupported parameter out_fd_list : no type generator for UnixFDList, GUnixFDList**

// Unsupported : g_dbus_proxy_call_with_unix_fd_list_sync : unsupported parameter parameters : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property : return type : Blacklisted record : GVariant

// Unsupported : g_dbus_proxy_get_cached_property_names : no return type

// Unsupported : g_dbus_proxy_get_connection : no return generator

// GetDefaultTimeout is a wrapper around the C function g_dbus_proxy_get_default_timeout.
func (recv *DBusProxy) GetDefaultTimeout() int32 {
	retC := C.g_dbus_proxy_get_default_timeout((*C.GDBusProxy)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_dbus_proxy_get_flags : no return generator

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

// Unsupported : g_dbus_proxy_set_default_timeout : no return generator

// Unsupported : g_dbus_proxy_set_interface_info : no return generator

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

func (recv *DBusServer) toC() *C.GDBusServer {

	return recv.native
}

// Unsupported : g_dbus_server_new_sync : unsupported parameter flags : no type generator for DBusServerFlags, GDBusServerFlags

// GetClientAddress is a wrapper around the C function g_dbus_server_get_client_address.
func (recv *DBusServer) GetClientAddress() string {
	retC := C.g_dbus_server_get_client_address((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_server_get_flags : no return generator

// GetGuid is a wrapper around the C function g_dbus_server_get_guid.
func (recv *DBusServer) GetGuid() string {
	retC := C.g_dbus_server_get_guid((*C.GDBusServer)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_dbus_server_is_active : no return generator

// Unsupported : g_dbus_server_start : no return generator

// Unsupported : g_dbus_server_stop : no return generator

// ProxyAddress is a wrapper around the C record GProxyAddress.
type ProxyAddress struct {
	native *C.GProxyAddress
	// parent_instance : no type generator for InetSocketAddress, GInetSocketAddress
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

func (recv *ProxyAddress) toC() *C.GProxyAddress {

	return recv.native
}

// Unsupported : g_proxy_address_new : unsupported parameter inetaddr : no type generator for InetAddress, GInetAddress*

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

// UnixCredentialsMessage is a wrapper around the C record GUnixCredentialsMessage.
type UnixCredentialsMessage struct {
	native *C.GUnixCredentialsMessage
	// parent_instance : no type generator for SocketControlMessage, GSocketControlMessage
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

func (recv *UnixCredentialsMessage) toC() *C.GUnixCredentialsMessage {

	return recv.native
}

// Unsupported : g_unix_credentials_message_new : no return generator

// Unsupported : g_unix_credentials_message_new_with_credentials : unsupported parameter credentials : no type generator for Credentials, GCredentials*

// Unsupported : g_unix_credentials_message_get_credentials : no return generator
