// Code generated - DO NOT EDIT.
// +build gio_2.26

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// PROXY_EXTENSION_POINT_NAME is a representation of the C constant G_PROXY_EXTENSION_POINT_NAME.
//
// since 2.26
const PROXY_EXTENSION_POINT_NAME = "gio-proxy"

// BusNameOwnerFlags is a representation of the C bitfield GBusNameOwnerFlags.
type BusNameOwnerFlags int

// BusNameOwnerFlags_none is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_NONE.
const BusNameOwnerFlags_none = BusNameOwnerFlags(0)

// BusNameOwnerFlags_allow_replacement is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT.
const BusNameOwnerFlags_allow_replacement = BusNameOwnerFlags(1)

// BusNameOwnerFlags_replace is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_REPLACE.
const BusNameOwnerFlags_replace = BusNameOwnerFlags(2)

// BusNameOwnerFlags_do_not_queue is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE.
const BusNameOwnerFlags_do_not_queue = BusNameOwnerFlags(4)

// BusNameWatcherFlags is a representation of the C bitfield GBusNameWatcherFlags.
type BusNameWatcherFlags int

// BusNameWatcherFlags_none is a representation of the C bitfield member G_BUS_NAME_WATCHER_FLAGS_NONE.
const BusNameWatcherFlags_none = BusNameWatcherFlags(0)

// BusNameWatcherFlags_auto_start is a representation of the C bitfield member G_BUS_NAME_WATCHER_FLAGS_AUTO_START.
const BusNameWatcherFlags_auto_start = BusNameWatcherFlags(1)

// DBusCallFlags is a representation of the C bitfield GDBusCallFlags.
type DBusCallFlags int

// DBusCallFlags_none is a representation of the C bitfield member G_DBUS_CALL_FLAGS_NONE.
const DBusCallFlags_none = DBusCallFlags(0)

// DBusCallFlags_no_auto_start is a representation of the C bitfield member G_DBUS_CALL_FLAGS_NO_AUTO_START.
const DBusCallFlags_no_auto_start = DBusCallFlags(1)

// DBusCallFlags_allow_interactive_authorization is a representation of the C bitfield member G_DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
const DBusCallFlags_allow_interactive_authorization = DBusCallFlags(2)

// DBusCapabilityFlags is a representation of the C bitfield GDBusCapabilityFlags.
type DBusCapabilityFlags int

// DBusCapabilityFlags_none is a representation of the C bitfield member G_DBUS_CAPABILITY_FLAGS_NONE.
const DBusCapabilityFlags_none = DBusCapabilityFlags(0)

// DBusCapabilityFlags_unix_fd_passing is a representation of the C bitfield member G_DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING.
const DBusCapabilityFlags_unix_fd_passing = DBusCapabilityFlags(1)

// DBusConnectionFlags is a representation of the C bitfield GDBusConnectionFlags.
type DBusConnectionFlags int

// DBusConnectionFlags_none is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_NONE.
const DBusConnectionFlags_none = DBusConnectionFlags(0)

// DBusConnectionFlags_authentication_client is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT.
const DBusConnectionFlags_authentication_client = DBusConnectionFlags(1)

// DBusConnectionFlags_authentication_server is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER.
const DBusConnectionFlags_authentication_server = DBusConnectionFlags(2)

// DBusConnectionFlags_authentication_allow_anonymous is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
const DBusConnectionFlags_authentication_allow_anonymous = DBusConnectionFlags(4)

// DBusConnectionFlags_message_bus_connection is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION.
const DBusConnectionFlags_message_bus_connection = DBusConnectionFlags(8)

// DBusConnectionFlags_delay_message_processing is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING.
const DBusConnectionFlags_delay_message_processing = DBusConnectionFlags(16)

// DBusMessageFlags is a representation of the C bitfield GDBusMessageFlags.
type DBusMessageFlags int

// DBusMessageFlags_none is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NONE.
const DBusMessageFlags_none = DBusMessageFlags(0)

// DBusMessageFlags_no_reply_expected is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED.
const DBusMessageFlags_no_reply_expected = DBusMessageFlags(1)

// DBusMessageFlags_no_auto_start is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NO_AUTO_START.
const DBusMessageFlags_no_auto_start = DBusMessageFlags(2)

// DBusMessageFlags_allow_interactive_authorization is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
const DBusMessageFlags_allow_interactive_authorization = DBusMessageFlags(4)

// DBusPropertyInfoFlags is a representation of the C bitfield GDBusPropertyInfoFlags.
type DBusPropertyInfoFlags int

// DBusPropertyInfoFlags_none is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_NONE.
const DBusPropertyInfoFlags_none = DBusPropertyInfoFlags(0)

// DBusPropertyInfoFlags_readable is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_READABLE.
const DBusPropertyInfoFlags_readable = DBusPropertyInfoFlags(1)

// DBusPropertyInfoFlags_writable is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_WRITABLE.
const DBusPropertyInfoFlags_writable = DBusPropertyInfoFlags(2)

// DBusProxyFlags is a representation of the C bitfield GDBusProxyFlags.
type DBusProxyFlags int

// DBusProxyFlags_none is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_NONE.
const DBusProxyFlags_none = DBusProxyFlags(0)

// DBusProxyFlags_do_not_load_properties is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES.
const DBusProxyFlags_do_not_load_properties = DBusProxyFlags(1)

// DBusProxyFlags_do_not_connect_signals is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS.
const DBusProxyFlags_do_not_connect_signals = DBusProxyFlags(2)

// DBusProxyFlags_do_not_auto_start is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START.
const DBusProxyFlags_do_not_auto_start = DBusProxyFlags(4)

// DBusProxyFlags_get_invalidated_properties is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES.
const DBusProxyFlags_get_invalidated_properties = DBusProxyFlags(8)

// DBusProxyFlags_do_not_auto_start_at_construction is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION.
const DBusProxyFlags_do_not_auto_start_at_construction = DBusProxyFlags(16)

// DBusSendMessageFlags is a representation of the C bitfield GDBusSendMessageFlags.
type DBusSendMessageFlags int

// DBusSendMessageFlags_none is a representation of the C bitfield member G_DBUS_SEND_MESSAGE_FLAGS_NONE.
const DBusSendMessageFlags_none = DBusSendMessageFlags(0)

// DBusSendMessageFlags_preserve_serial is a representation of the C bitfield member G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL.
const DBusSendMessageFlags_preserve_serial = DBusSendMessageFlags(1)

// DBusServerFlags is a representation of the C bitfield GDBusServerFlags.
type DBusServerFlags int

// DBusServerFlags_none is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_NONE.
const DBusServerFlags_none = DBusServerFlags(0)

// DBusServerFlags_run_in_thread is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_RUN_IN_THREAD.
const DBusServerFlags_run_in_thread = DBusServerFlags(1)

// DBusServerFlags_authentication_allow_anonymous is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
const DBusServerFlags_authentication_allow_anonymous = DBusServerFlags(2)

// DBusSignalFlags is a representation of the C bitfield GDBusSignalFlags.
type DBusSignalFlags int

// DBusSignalFlags_none is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_NONE.
const DBusSignalFlags_none = DBusSignalFlags(0)

// DBusSignalFlags_no_match_rule is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_NO_MATCH_RULE.
const DBusSignalFlags_no_match_rule = DBusSignalFlags(1)

// DBusSignalFlags_match_arg0_namespace is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE.
const DBusSignalFlags_match_arg0_namespace = DBusSignalFlags(2)

// DBusSignalFlags_match_arg0_path is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH.
const DBusSignalFlags_match_arg0_path = DBusSignalFlags(4)

// DBusSubtreeFlags is a representation of the C bitfield GDBusSubtreeFlags.
type DBusSubtreeFlags int

// DBusSubtreeFlags_none is a representation of the C bitfield member G_DBUS_SUBTREE_FLAGS_NONE.
const DBusSubtreeFlags_none = DBusSubtreeFlags(0)

// DBusSubtreeFlags_dispatch_to_unenumerated_nodes is a representation of the C bitfield member G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES.
const DBusSubtreeFlags_dispatch_to_unenumerated_nodes = DBusSubtreeFlags(1)

// BusType is a representation of the C enumeration GBusType.
type BusType int

// BusType_starter is a representation of the C enumeration member G_BUS_TYPE_STARTER.
const BusType_starter = BusType(-1)

// BusType_none is a representation of the C enumeration member G_BUS_TYPE_NONE.
const BusType_none = BusType(0)

// BusType_system is a representation of the C enumeration member G_BUS_TYPE_SYSTEM.
const BusType_system = BusType(1)

// BusType_session is a representation of the C enumeration member G_BUS_TYPE_SESSION.
const BusType_session = BusType(2)

// CredentialsType is a representation of the C enumeration GCredentialsType.
type CredentialsType int

// CredentialsType_invalid is a representation of the C enumeration member G_CREDENTIALS_TYPE_INVALID.
const CredentialsType_invalid = CredentialsType(0)

// CredentialsType_linux_ucred is a representation of the C enumeration member G_CREDENTIALS_TYPE_LINUX_UCRED.
const CredentialsType_linux_ucred = CredentialsType(1)

// CredentialsType_freebsd_cmsgcred is a representation of the C enumeration member G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
const CredentialsType_freebsd_cmsgcred = CredentialsType(2)

// CredentialsType_openbsd_sockpeercred is a representation of the C enumeration member G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
const CredentialsType_openbsd_sockpeercred = CredentialsType(3)

// CredentialsType_solaris_ucred is a representation of the C enumeration member G_CREDENTIALS_TYPE_SOLARIS_UCRED.
const CredentialsType_solaris_ucred = CredentialsType(4)

// CredentialsType_netbsd_unpcbid is a representation of the C enumeration member G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
const CredentialsType_netbsd_unpcbid = CredentialsType(5)

// DBusError is a representation of the C enumeration GDBusError.
type DBusError int

// DBusError_failed is a representation of the C enumeration member G_DBUS_ERROR_FAILED.
const DBusError_failed = DBusError(0)

// DBusError_no_memory is a representation of the C enumeration member G_DBUS_ERROR_NO_MEMORY.
const DBusError_no_memory = DBusError(1)

// DBusError_service_unknown is a representation of the C enumeration member G_DBUS_ERROR_SERVICE_UNKNOWN.
const DBusError_service_unknown = DBusError(2)

// DBusError_name_has_no_owner is a representation of the C enumeration member G_DBUS_ERROR_NAME_HAS_NO_OWNER.
const DBusError_name_has_no_owner = DBusError(3)

// DBusError_no_reply is a representation of the C enumeration member G_DBUS_ERROR_NO_REPLY.
const DBusError_no_reply = DBusError(4)

// DBusError_io_error is a representation of the C enumeration member G_DBUS_ERROR_IO_ERROR.
const DBusError_io_error = DBusError(5)

// DBusError_bad_address is a representation of the C enumeration member G_DBUS_ERROR_BAD_ADDRESS.
const DBusError_bad_address = DBusError(6)

// DBusError_not_supported is a representation of the C enumeration member G_DBUS_ERROR_NOT_SUPPORTED.
const DBusError_not_supported = DBusError(7)

// DBusError_limits_exceeded is a representation of the C enumeration member G_DBUS_ERROR_LIMITS_EXCEEDED.
const DBusError_limits_exceeded = DBusError(8)

// DBusError_access_denied is a representation of the C enumeration member G_DBUS_ERROR_ACCESS_DENIED.
const DBusError_access_denied = DBusError(9)

// DBusError_auth_failed is a representation of the C enumeration member G_DBUS_ERROR_AUTH_FAILED.
const DBusError_auth_failed = DBusError(10)

// DBusError_no_server is a representation of the C enumeration member G_DBUS_ERROR_NO_SERVER.
const DBusError_no_server = DBusError(11)

// DBusError_timeout is a representation of the C enumeration member G_DBUS_ERROR_TIMEOUT.
const DBusError_timeout = DBusError(12)

// DBusError_no_network is a representation of the C enumeration member G_DBUS_ERROR_NO_NETWORK.
const DBusError_no_network = DBusError(13)

// DBusError_address_in_use is a representation of the C enumeration member G_DBUS_ERROR_ADDRESS_IN_USE.
const DBusError_address_in_use = DBusError(14)

// DBusError_disconnected is a representation of the C enumeration member G_DBUS_ERROR_DISCONNECTED.
const DBusError_disconnected = DBusError(15)

// DBusError_invalid_args is a representation of the C enumeration member G_DBUS_ERROR_INVALID_ARGS.
const DBusError_invalid_args = DBusError(16)

// DBusError_file_not_found is a representation of the C enumeration member G_DBUS_ERROR_FILE_NOT_FOUND.
const DBusError_file_not_found = DBusError(17)

// DBusError_file_exists is a representation of the C enumeration member G_DBUS_ERROR_FILE_EXISTS.
const DBusError_file_exists = DBusError(18)

// DBusError_unknown_method is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_METHOD.
const DBusError_unknown_method = DBusError(19)

// DBusError_timed_out is a representation of the C enumeration member G_DBUS_ERROR_TIMED_OUT.
const DBusError_timed_out = DBusError(20)

// DBusError_match_rule_not_found is a representation of the C enumeration member G_DBUS_ERROR_MATCH_RULE_NOT_FOUND.
const DBusError_match_rule_not_found = DBusError(21)

// DBusError_match_rule_invalid is a representation of the C enumeration member G_DBUS_ERROR_MATCH_RULE_INVALID.
const DBusError_match_rule_invalid = DBusError(22)

// DBusError_spawn_exec_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_EXEC_FAILED.
const DBusError_spawn_exec_failed = DBusError(23)

// DBusError_spawn_fork_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FORK_FAILED.
const DBusError_spawn_fork_failed = DBusError(24)

// DBusError_spawn_child_exited is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CHILD_EXITED.
const DBusError_spawn_child_exited = DBusError(25)

// DBusError_spawn_child_signaled is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CHILD_SIGNALED.
const DBusError_spawn_child_signaled = DBusError(26)

// DBusError_spawn_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FAILED.
const DBusError_spawn_failed = DBusError(27)

// DBusError_spawn_setup_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SETUP_FAILED.
const DBusError_spawn_setup_failed = DBusError(28)

// DBusError_spawn_config_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CONFIG_INVALID.
const DBusError_spawn_config_invalid = DBusError(29)

// DBusError_spawn_service_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SERVICE_INVALID.
const DBusError_spawn_service_invalid = DBusError(30)

// DBusError_spawn_service_not_found is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND.
const DBusError_spawn_service_not_found = DBusError(31)

// DBusError_spawn_permissions_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID.
const DBusError_spawn_permissions_invalid = DBusError(32)

// DBusError_spawn_file_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FILE_INVALID.
const DBusError_spawn_file_invalid = DBusError(33)

// DBusError_spawn_no_memory is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_NO_MEMORY.
const DBusError_spawn_no_memory = DBusError(34)

// DBusError_unix_process_id_unknown is a representation of the C enumeration member G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN.
const DBusError_unix_process_id_unknown = DBusError(35)

// DBusError_invalid_signature is a representation of the C enumeration member G_DBUS_ERROR_INVALID_SIGNATURE.
const DBusError_invalid_signature = DBusError(36)

// DBusError_invalid_file_content is a representation of the C enumeration member G_DBUS_ERROR_INVALID_FILE_CONTENT.
const DBusError_invalid_file_content = DBusError(37)

// DBusError_selinux_security_context_unknown is a representation of the C enumeration member G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN.
const DBusError_selinux_security_context_unknown = DBusError(38)

// DBusError_adt_audit_data_unknown is a representation of the C enumeration member G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN.
const DBusError_adt_audit_data_unknown = DBusError(39)

// DBusError_object_path_in_use is a representation of the C enumeration member G_DBUS_ERROR_OBJECT_PATH_IN_USE.
const DBusError_object_path_in_use = DBusError(40)

// DBusError_unknown_object is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_OBJECT.
const DBusError_unknown_object = DBusError(41)

// DBusError_unknown_interface is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_INTERFACE.
const DBusError_unknown_interface = DBusError(42)

// DBusError_unknown_property is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_PROPERTY.
const DBusError_unknown_property = DBusError(43)

// DBusError_property_read_only is a representation of the C enumeration member G_DBUS_ERROR_PROPERTY_READ_ONLY.
const DBusError_property_read_only = DBusError(44)

// DBusMessageByteOrder is a representation of the C enumeration GDBusMessageByteOrder.
type DBusMessageByteOrder int

// DBusMessageByteOrder_big_endian is a representation of the C enumeration member G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN.
const DBusMessageByteOrder_big_endian = DBusMessageByteOrder(66)

// DBusMessageByteOrder_little_endian is a representation of the C enumeration member G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN.
const DBusMessageByteOrder_little_endian = DBusMessageByteOrder(108)

// DBusMessageHeaderField is a representation of the C enumeration GDBusMessageHeaderField.
type DBusMessageHeaderField int

// DBusMessageHeaderField_invalid is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_INVALID.
const DBusMessageHeaderField_invalid = DBusMessageHeaderField(0)

// DBusMessageHeaderField_path is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_PATH.
const DBusMessageHeaderField_path = DBusMessageHeaderField(1)

// DBusMessageHeaderField_interface is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE.
const DBusMessageHeaderField_interface = DBusMessageHeaderField(2)

// DBusMessageHeaderField_member is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_MEMBER.
const DBusMessageHeaderField_member = DBusMessageHeaderField(3)

// DBusMessageHeaderField_error_name is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME.
const DBusMessageHeaderField_error_name = DBusMessageHeaderField(4)

// DBusMessageHeaderField_reply_serial is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL.
const DBusMessageHeaderField_reply_serial = DBusMessageHeaderField(5)

// DBusMessageHeaderField_destination is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION.
const DBusMessageHeaderField_destination = DBusMessageHeaderField(6)

// DBusMessageHeaderField_sender is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_SENDER.
const DBusMessageHeaderField_sender = DBusMessageHeaderField(7)

// DBusMessageHeaderField_signature is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE.
const DBusMessageHeaderField_signature = DBusMessageHeaderField(8)

// DBusMessageHeaderField_num_unix_fds is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS.
const DBusMessageHeaderField_num_unix_fds = DBusMessageHeaderField(9)

// DBusMessageType is a representation of the C enumeration GDBusMessageType.
type DBusMessageType int

// DBusMessageType_invalid is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_INVALID.
const DBusMessageType_invalid = DBusMessageType(0)

// DBusMessageType_method_call is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_METHOD_CALL.
const DBusMessageType_method_call = DBusMessageType(1)

// DBusMessageType_method_return is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_METHOD_RETURN.
const DBusMessageType_method_return = DBusMessageType(2)

// DBusMessageType_error is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_ERROR.
const DBusMessageType_error = DBusMessageType(3)

// DBusMessageType_signal is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_SIGNAL.
const DBusMessageType_signal = DBusMessageType(4)

// UnixSocketAddressType is a representation of the C enumeration GUnixSocketAddressType.
type UnixSocketAddressType int

// UnixSocketAddressType_invalid is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_INVALID.
const UnixSocketAddressType_invalid = UnixSocketAddressType(0)

// UnixSocketAddressType_anonymous is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ANONYMOUS.
const UnixSocketAddressType_anonymous = UnixSocketAddressType(1)

// UnixSocketAddressType_path is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_PATH.
const UnixSocketAddressType_path = UnixSocketAddressType(2)

// UnixSocketAddressType_abstract is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ABSTRACT.
const UnixSocketAddressType_abstract = UnixSocketAddressType(3)

// UnixSocketAddressType_abstract_padded is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
const UnixSocketAddressType_abstract_padded = UnixSocketAddressType(4)

// UNSUPPORTED : g_action_parse_detailed_name : throws

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// BusOwnNameOnConnectionWithClosures wraps the C function g_bus_own_name_on_connection_with_closures.
//
// since 2.26
func BusOwnNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint {
	sys_connection := connection.ToC()
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAcquiredClosure := nameAcquiredClosure.ToC()
	sys_nameLostClosure := nameLostClosure.ToC()
	retSys := gio.Fn_g_bus_own_name_on_connection_with_closures(sys_connection, sys_name, sys_flags, sys_nameAcquiredClosure, sys_nameLostClosure)
	ret := retSys

	return ret
}

// BusOwnNameWithClosures wraps the C function g_bus_own_name_with_closures.
//
// since 2.26
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint {
	sys_busType := (int)(busType)
	sys_name := name
	sys_flags := (int)(flags)
	sys_busAcquiredClosure := busAcquiredClosure.ToC()
	sys_nameAcquiredClosure := nameAcquiredClosure.ToC()
	sys_nameLostClosure := nameLostClosure.ToC()
	retSys := gio.Fn_g_bus_own_name_with_closures(sys_busType, sys_name, sys_flags, sys_busAcquiredClosure, sys_nameAcquiredClosure, sys_nameLostClosure)
	ret := retSys

	return ret
}

// BusUnownName wraps the C function g_bus_unown_name.
//
// since 2.26
func BusUnownName(ownerId uint) {
	sys_ownerId := ownerId
	gio.Fn_g_bus_unown_name(sys_ownerId)
}

// BusUnwatchName wraps the C function g_bus_unwatch_name.
//
// since 2.26
func BusUnwatchName(watcherId uint) {
	sys_watcherId := watcherId
	gio.Fn_g_bus_unwatch_name(sys_watcherId)
}

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// BusWatchNameOnConnectionWithClosures wraps the C function g_bus_watch_name_on_connection_with_closures.
//
// since 2.26
func BusWatchNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint {
	sys_connection := connection.ToC()
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAppearedClosure := nameAppearedClosure.ToC()
	sys_nameVanishedClosure := nameVanishedClosure.ToC()
	retSys := gio.Fn_g_bus_watch_name_on_connection_with_closures(sys_connection, sys_name, sys_flags, sys_nameAppearedClosure, sys_nameVanishedClosure)
	ret := retSys

	return ret
}

// BusWatchNameWithClosures wraps the C function g_bus_watch_name_with_closures.
//
// since 2.26
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint {
	sys_busType := (int)(busType)
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAppearedClosure := nameAppearedClosure.ToC()
	sys_nameVanishedClosure := nameVanishedClosure.ToC()
	retSys := gio.Fn_g_bus_watch_name_with_closures(sys_busType, sys_name, sys_flags, sys_nameAppearedClosure, sys_nameVanishedClosure)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// DbusGenerateGuid wraps the C function g_dbus_generate_guid.
//
// since 2.26
func DbusGenerateGuid() string {
	retSys := gio.Fn_g_dbus_generate_guid()
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// DbusIsAddress wraps the C function g_dbus_is_address.
//
// since 2.26
func DbusIsAddress(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_address(sys_string_)
	ret := retSys

	return ret
}

// DbusIsGuid wraps the C function g_dbus_is_guid.
//
// since 2.26
func DbusIsGuid(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_guid(sys_string_)
	ret := retSys

	return ret
}

// DbusIsInterfaceName wraps the C function g_dbus_is_interface_name.
//
// since 2.26
func DbusIsInterfaceName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_interface_name(sys_string_)
	ret := retSys

	return ret
}

// DbusIsMemberName wraps the C function g_dbus_is_member_name.
//
// since 2.26
func DbusIsMemberName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_member_name(sys_string_)
	ret := retSys

	return ret
}

// DbusIsName wraps the C function g_dbus_is_name.
//
// since 2.26
func DbusIsName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_name(sys_string_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dbus_is_supported_address : throws

// DbusIsUniqueName wraps the C function g_dbus_is_unique_name.
//
// since 2.26
func DbusIsUniqueName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_unique_name(sys_string_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UNSUPPORTED : g_unix_mount_at : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_for : has [in]out param, time_read

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// CredentialsClass is a representation of the C record GCredentialsClass.
//
// since 2.26
type CredentialsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCredentialsClass that represents the CredentialsClass.
func (recv *CredentialsClass) ToC() unsafe.Pointer {
	return recv.native
}

// CredentialsClassNewFromC creates a new CredentialsClass from a pointer to the C GCredentialsClass that represents the CredentialsClass.
func CredentialsClassNewFromC(native unsafe.Pointer) *CredentialsClass {
	return &CredentialsClass{native: native}
}

// DBusAnnotationInfo is a representation of the C record GDBusAnnotationInfo.
//
// since 2.26
type DBusAnnotationInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusAnnotationInfo that represents the DBusAnnotationInfo.
func (recv *DBusAnnotationInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusAnnotationInfoNewFromC creates a new DBusAnnotationInfo from a pointer to the C GDBusAnnotationInfo that represents the DBusAnnotationInfo.
func DBusAnnotationInfoNewFromC(native unsafe.Pointer) *DBusAnnotationInfo {
	return &DBusAnnotationInfo{native: native}
}

// DBusArgInfo is a representation of the C record GDBusArgInfo.
//
// since 2.26
type DBusArgInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusArgInfo that represents the DBusArgInfo.
func (recv *DBusArgInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusArgInfoNewFromC creates a new DBusArgInfo from a pointer to the C GDBusArgInfo that represents the DBusArgInfo.
func DBusArgInfoNewFromC(native unsafe.Pointer) *DBusArgInfo {
	return &DBusArgInfo{native: native}
}

// DBusErrorEntry is a representation of the C record GDBusErrorEntry.
//
// since 2.26
type DBusErrorEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusErrorEntry that represents the DBusErrorEntry.
func (recv *DBusErrorEntry) ToC() unsafe.Pointer {
	return recv.native
}

// DBusErrorEntryNewFromC creates a new DBusErrorEntry from a pointer to the C GDBusErrorEntry that represents the DBusErrorEntry.
func DBusErrorEntryNewFromC(native unsafe.Pointer) *DBusErrorEntry {
	return &DBusErrorEntry{native: native}
}

// DBusInterfaceInfo is a representation of the C record GDBusInterfaceInfo.
//
// since 2.26
type DBusInterfaceInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceInfo that represents the DBusInterfaceInfo.
func (recv *DBusInterfaceInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceInfoNewFromC creates a new DBusInterfaceInfo from a pointer to the C GDBusInterfaceInfo that represents the DBusInterfaceInfo.
func DBusInterfaceInfoNewFromC(native unsafe.Pointer) *DBusInterfaceInfo {
	return &DBusInterfaceInfo{native: native}
}

// DBusInterfaceVTable is a representation of the C record GDBusInterfaceVTable.
//
// since 2.26
type DBusInterfaceVTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceVTable that represents the DBusInterfaceVTable.
func (recv *DBusInterfaceVTable) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceVTableNewFromC creates a new DBusInterfaceVTable from a pointer to the C GDBusInterfaceVTable that represents the DBusInterfaceVTable.
func DBusInterfaceVTableNewFromC(native unsafe.Pointer) *DBusInterfaceVTable {
	return &DBusInterfaceVTable{native: native}
}

// DBusMethodInfo is a representation of the C record GDBusMethodInfo.
//
// since 2.26
type DBusMethodInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMethodInfo that represents the DBusMethodInfo.
func (recv *DBusMethodInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMethodInfoNewFromC creates a new DBusMethodInfo from a pointer to the C GDBusMethodInfo that represents the DBusMethodInfo.
func DBusMethodInfoNewFromC(native unsafe.Pointer) *DBusMethodInfo {
	return &DBusMethodInfo{native: native}
}

// DBusNodeInfo is a representation of the C record GDBusNodeInfo.
//
// since 2.26
type DBusNodeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusNodeInfo that represents the DBusNodeInfo.
func (recv *DBusNodeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusNodeInfoNewFromC creates a new DBusNodeInfo from a pointer to the C GDBusNodeInfo that represents the DBusNodeInfo.
func DBusNodeInfoNewFromC(native unsafe.Pointer) *DBusNodeInfo {
	return &DBusNodeInfo{native: native}
}

// DBusPropertyInfo is a representation of the C record GDBusPropertyInfo.
//
// since 2.26
type DBusPropertyInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusPropertyInfo that represents the DBusPropertyInfo.
func (recv *DBusPropertyInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusPropertyInfoNewFromC creates a new DBusPropertyInfo from a pointer to the C GDBusPropertyInfo that represents the DBusPropertyInfo.
func DBusPropertyInfoNewFromC(native unsafe.Pointer) *DBusPropertyInfo {
	return &DBusPropertyInfo{native: native}
}

// DBusProxyClass is a representation of the C record GDBusProxyClass.
//
// since 2.26
type DBusProxyClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxyClass that represents the DBusProxyClass.
func (recv *DBusProxyClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyClassNewFromC creates a new DBusProxyClass from a pointer to the C GDBusProxyClass that represents the DBusProxyClass.
func DBusProxyClassNewFromC(native unsafe.Pointer) *DBusProxyClass {
	return &DBusProxyClass{native: native}
}

// DBusSignalInfo is a representation of the C record GDBusSignalInfo.
//
// since 2.26
type DBusSignalInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusSignalInfo that represents the DBusSignalInfo.
func (recv *DBusSignalInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusSignalInfoNewFromC creates a new DBusSignalInfo from a pointer to the C GDBusSignalInfo that represents the DBusSignalInfo.
func DBusSignalInfoNewFromC(native unsafe.Pointer) *DBusSignalInfo {
	return &DBusSignalInfo{native: native}
}

// DBusSubtreeVTable is a representation of the C record GDBusSubtreeVTable.
//
// since 2.26
type DBusSubtreeVTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusSubtreeVTable that represents the DBusSubtreeVTable.
func (recv *DBusSubtreeVTable) ToC() unsafe.Pointer {
	return recv.native
}

// DBusSubtreeVTableNewFromC creates a new DBusSubtreeVTable from a pointer to the C GDBusSubtreeVTable that represents the DBusSubtreeVTable.
func DBusSubtreeVTableNewFromC(native unsafe.Pointer) *DBusSubtreeVTable {
	return &DBusSubtreeVTable{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// ProxyAddressClass is a representation of the C record GProxyAddressClass.
//
// since 2.26
type ProxyAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressClass that represents the ProxyAddressClass.
func (recv *ProxyAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressClassNewFromC creates a new ProxyAddressClass from a pointer to the C GProxyAddressClass that represents the ProxyAddressClass.
func ProxyAddressClassNewFromC(native unsafe.Pointer) *ProxyAddressClass {
	return &ProxyAddressClass{native: native}
}

// ProxyInterface is a representation of the C record GProxyInterface.
//
// since 2.26
type ProxyInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyInterface that represents the ProxyInterface.
func (recv *ProxyInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyInterfaceNewFromC creates a new ProxyInterface from a pointer to the C GProxyInterface that represents the ProxyInterface.
func ProxyInterfaceNewFromC(native unsafe.Pointer) *ProxyInterface {
	return &ProxyInterface{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// TlsClientConnectionInterface is a representation of the C record GTlsClientConnectionInterface.
//
// since 2.26
type TlsClientConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsClientConnectionInterface that represents the TlsClientConnectionInterface.
func (recv *TlsClientConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsClientConnectionInterfaceNewFromC creates a new TlsClientConnectionInterface from a pointer to the C GTlsClientConnectionInterface that represents the TlsClientConnectionInterface.
func TlsClientConnectionInterfaceNewFromC(native unsafe.Pointer) *TlsClientConnectionInterface {
	return &TlsClientConnectionInterface{native: native}
}

// TlsServerConnectionInterface is a representation of the C record GTlsServerConnectionInterface.
//
// since 2.26
type TlsServerConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsServerConnectionInterface that represents the TlsServerConnectionInterface.
func (recv *TlsServerConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsServerConnectionInterfaceNewFromC creates a new TlsServerConnectionInterface from a pointer to the C GTlsServerConnectionInterface that represents the TlsServerConnectionInterface.
func TlsServerConnectionInterfaceNewFromC(native unsafe.Pointer) *TlsServerConnectionInterface {
	return &TlsServerConnectionInterface{native: native}
}

// UnixCredentialsMessageClass is a representation of the C record GUnixCredentialsMessageClass.
//
// since 2.26
type UnixCredentialsMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessageClass that represents the UnixCredentialsMessageClass.
func (recv *UnixCredentialsMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessageClassNewFromC creates a new UnixCredentialsMessageClass from a pointer to the C GUnixCredentialsMessageClass that represents the UnixCredentialsMessageClass.
func UnixCredentialsMessageClassNewFromC(native unsafe.Pointer) *UnixCredentialsMessageClass {
	return &UnixCredentialsMessageClass{native: native}
}

// Credentials is a representation of the C record GCredentials.
//
// since 2.26
type Credentials struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCredentials that represents the Credentials.
func (recv *Credentials) ToC() unsafe.Pointer {
	return recv.native
}

// CredentialsNewFromC creates a new Credentials from a pointer to the C GCredentials that represents the Credentials.
func CredentialsNewFromC(native unsafe.Pointer) *Credentials {
	return &Credentials{native: native}
}

// DBusAuthObserver is a representation of the C record GDBusAuthObserver.
//
// since 2.26
type DBusAuthObserver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusAuthObserver that represents the DBusAuthObserver.
func (recv *DBusAuthObserver) ToC() unsafe.Pointer {
	return recv.native
}

// DBusAuthObserverNewFromC creates a new DBusAuthObserver from a pointer to the C GDBusAuthObserver that represents the DBusAuthObserver.
func DBusAuthObserverNewFromC(native unsafe.Pointer) *DBusAuthObserver {
	return &DBusAuthObserver{native: native}
}

// DBusConnection is a representation of the C record GDBusConnection.
//
// since 2.26
type DBusConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusConnection that represents the DBusConnection.
func (recv *DBusConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DBusConnectionNewFromC creates a new DBusConnection from a pointer to the C GDBusConnection that represents the DBusConnection.
func DBusConnectionNewFromC(native unsafe.Pointer) *DBusConnection {
	return &DBusConnection{native: native}
}

// DBusMessage is a representation of the C record GDBusMessage.
//
// since 2.26
type DBusMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMessage that represents the DBusMessage.
func (recv *DBusMessage) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMessageNewFromC creates a new DBusMessage from a pointer to the C GDBusMessage that represents the DBusMessage.
func DBusMessageNewFromC(native unsafe.Pointer) *DBusMessage {
	return &DBusMessage{native: native}
}

// DBusMethodInvocation is a representation of the C record GDBusMethodInvocation.
//
// since 2.26
type DBusMethodInvocation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMethodInvocation that represents the DBusMethodInvocation.
func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMethodInvocationNewFromC creates a new DBusMethodInvocation from a pointer to the C GDBusMethodInvocation that represents the DBusMethodInvocation.
func DBusMethodInvocationNewFromC(native unsafe.Pointer) *DBusMethodInvocation {
	return &DBusMethodInvocation{native: native}
}

// DBusProxy is a representation of the C record GDBusProxy.
//
// since 2.26
type DBusProxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxy that represents the DBusProxy.
func (recv *DBusProxy) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyNewFromC creates a new DBusProxy from a pointer to the C GDBusProxy that represents the DBusProxy.
func DBusProxyNewFromC(native unsafe.Pointer) *DBusProxy {
	return &DBusProxy{native: native}
}

// DBusServer is a representation of the C record GDBusServer.
//
// since 2.26
type DBusServer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusServer that represents the DBusServer.
func (recv *DBusServer) ToC() unsafe.Pointer {
	return recv.native
}

// DBusServerNewFromC creates a new DBusServer from a pointer to the C GDBusServer that represents the DBusServer.
func DBusServerNewFromC(native unsafe.Pointer) *DBusServer {
	return &DBusServer{native: native}
}

// ProxyAddress is a representation of the C record GProxyAddress.
//
// since 2.26
type ProxyAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddress that represents the ProxyAddress.
func (recv *ProxyAddress) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressNewFromC creates a new ProxyAddress from a pointer to the C GProxyAddress that represents the ProxyAddress.
func ProxyAddressNewFromC(native unsafe.Pointer) *ProxyAddress {
	return &ProxyAddress{native: native}
}

// UnixCredentialsMessage is a representation of the C record GUnixCredentialsMessage.
//
// since 2.26
type UnixCredentialsMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessage that represents the UnixCredentialsMessage.
func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessageNewFromC creates a new UnixCredentialsMessage from a pointer to the C GUnixCredentialsMessage that represents the UnixCredentialsMessage.
func UnixCredentialsMessageNewFromC(native unsafe.Pointer) *UnixCredentialsMessage {
	return &UnixCredentialsMessage{native: native}
}

// Proxy is a representation of the C interface GProxy.
//
// since 2.26
type Proxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxy that represents the Proxy.
func (recv *Proxy) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyNewFromC creates a new Proxy from a pointer to the C GProxy that represents the Proxy.
func ProxyNewFromC(native unsafe.Pointer) *Proxy {
	return &Proxy{native: native}
}

// ProxyResolver is a representation of the C interface GProxyResolver.
//
// since 2.26
type ProxyResolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyResolver that represents the ProxyResolver.
func (recv *ProxyResolver) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyResolverNewFromC creates a new ProxyResolver from a pointer to the C GProxyResolver that represents the ProxyResolver.
func ProxyResolverNewFromC(native unsafe.Pointer) *ProxyResolver {
	return &ProxyResolver{native: native}
}
