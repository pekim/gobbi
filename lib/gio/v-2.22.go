// Code generated - DO NOT EDIT.
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56 gio_2.58 gio_2.60 gio_2.62

package gio

import "unsafe"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL = "mountable::can-poll"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_START is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START = "mountable::can-start"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED = "mountable::can-start-degraded"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP = "mountable::can-stop"

// FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC = "mountable::is-media-check-automatic"

// FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE = "mountable::start-stop-type"

// FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE = "mountable::unix-device-file"

// DriveStartFlags is a representation of the C bitfield GDriveStartFlags.
type DriveStartFlags int

// DriveStartFlags_none is a representation of the C bitfield member G_DRIVE_START_NONE.
const DriveStartFlags_none = DriveStartFlags(0)

// SocketMsgFlags is a representation of the C bitfield GSocketMsgFlags.
type SocketMsgFlags int

// SocketMsgFlags_none is a representation of the C bitfield member G_SOCKET_MSG_NONE.
const SocketMsgFlags_none = SocketMsgFlags(0)

// SocketMsgFlags_oob is a representation of the C bitfield member G_SOCKET_MSG_OOB.
const SocketMsgFlags_oob = SocketMsgFlags(1)

// SocketMsgFlags_peek is a representation of the C bitfield member G_SOCKET_MSG_PEEK.
const SocketMsgFlags_peek = SocketMsgFlags(2)

// SocketMsgFlags_dontroute is a representation of the C bitfield member G_SOCKET_MSG_DONTROUTE.
const SocketMsgFlags_dontroute = SocketMsgFlags(4)

// DriveStartStopType is a representation of the C enumeration GDriveStartStopType.
type DriveStartStopType int

// DriveStartStopType_unknown is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_UNKNOWN.
const DriveStartStopType_unknown = DriveStartStopType(0)

// DriveStartStopType_shutdown is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_SHUTDOWN.
const DriveStartStopType_shutdown = DriveStartStopType(1)

// DriveStartStopType_network is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_NETWORK.
const DriveStartStopType_network = DriveStartStopType(2)

// DriveStartStopType_multidisk is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_MULTIDISK.
const DriveStartStopType_multidisk = DriveStartStopType(3)

// DriveStartStopType_password is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_PASSWORD.
const DriveStartStopType_password = DriveStartStopType(4)

// ResolverError is a representation of the C enumeration GResolverError.
type ResolverError int

// ResolverError_not_found is a representation of the C enumeration member G_RESOLVER_ERROR_NOT_FOUND.
const ResolverError_not_found = ResolverError(0)

// ResolverError_temporary_failure is a representation of the C enumeration member G_RESOLVER_ERROR_TEMPORARY_FAILURE.
const ResolverError_temporary_failure = ResolverError(1)

// ResolverError_internal is a representation of the C enumeration member G_RESOLVER_ERROR_INTERNAL.
const ResolverError_internal = ResolverError(2)

// SocketFamily is a representation of the C enumeration GSocketFamily.
type SocketFamily int

// SocketFamily_invalid is a representation of the C enumeration member G_SOCKET_FAMILY_INVALID.
const SocketFamily_invalid = SocketFamily(0)

// SocketFamily_unix is a representation of the C enumeration member G_SOCKET_FAMILY_UNIX.
const SocketFamily_unix = SocketFamily(1)

// SocketFamily_ipv4 is a representation of the C enumeration member G_SOCKET_FAMILY_IPV4.
const SocketFamily_ipv4 = SocketFamily(2)

// SocketFamily_ipv6 is a representation of the C enumeration member G_SOCKET_FAMILY_IPV6.
const SocketFamily_ipv6 = SocketFamily(10)

// SocketProtocol is a representation of the C enumeration GSocketProtocol.
type SocketProtocol int

// SocketProtocol_unknown is a representation of the C enumeration member G_SOCKET_PROTOCOL_UNKNOWN.
const SocketProtocol_unknown = SocketProtocol(-1)

// SocketProtocol_default is a representation of the C enumeration member G_SOCKET_PROTOCOL_DEFAULT.
const SocketProtocol_default = SocketProtocol(0)

// SocketProtocol_tcp is a representation of the C enumeration member G_SOCKET_PROTOCOL_TCP.
const SocketProtocol_tcp = SocketProtocol(6)

// SocketProtocol_udp is a representation of the C enumeration member G_SOCKET_PROTOCOL_UDP.
const SocketProtocol_udp = SocketProtocol(17)

// SocketProtocol_sctp is a representation of the C enumeration member G_SOCKET_PROTOCOL_SCTP.
const SocketProtocol_sctp = SocketProtocol(132)

// SocketType is a representation of the C enumeration GSocketType.
type SocketType int

// SocketType_invalid is a representation of the C enumeration member G_SOCKET_TYPE_INVALID.
const SocketType_invalid = SocketType(0)

// SocketType_stream is a representation of the C enumeration member G_SOCKET_TYPE_STREAM.
const SocketType_stream = SocketType(1)

// SocketType_datagram is a representation of the C enumeration member G_SOCKET_TYPE_DATAGRAM.
const SocketType_datagram = SocketType(2)

// SocketType_seqpacket is a representation of the C enumeration member G_SOCKET_TYPE_SEQPACKET.
const SocketType_seqpacket = SocketType(3)

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

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

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

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

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

// AsyncInitableIface is a representation of the C record GAsyncInitableIface.
//
// since 2.22
type AsyncInitableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncInitableIface that represents the AsyncInitableIface.
func (recv *AsyncInitableIface) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncInitableIfaceNewFromC creates a new AsyncInitableIface from a pointer to the C GAsyncInitableIface that represents the AsyncInitableIface.
func AsyncInitableIfaceNewFromC(native unsafe.Pointer) *AsyncInitableIface {
	return &AsyncInitableIface{native: native}
}

// InitableIface is a representation of the C record GInitableIface.
//
// since 2.22
type InitableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitableIface that represents the InitableIface.
func (recv *InitableIface) ToC() unsafe.Pointer {
	return recv.native
}

// InitableIfaceNewFromC creates a new InitableIface from a pointer to the C GInitableIface that represents the InitableIface.
func InitableIfaceNewFromC(native unsafe.Pointer) *InitableIface {
	return &InitableIface{native: native}
}

// InputVector is a representation of the C record GInputVector.
//
// since 2.22
type InputVector struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputVector that represents the InputVector.
func (recv *InputVector) ToC() unsafe.Pointer {
	return recv.native
}

// InputVectorNewFromC creates a new InputVector from a pointer to the C GInputVector that represents the InputVector.
func InputVectorNewFromC(native unsafe.Pointer) *InputVector {
	return &InputVector{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// OutputVector is a representation of the C record GOutputVector.
//
// since 2.22
type OutputVector struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputVector that represents the OutputVector.
func (recv *OutputVector) ToC() unsafe.Pointer {
	return recv.native
}

// OutputVectorNewFromC creates a new OutputVector from a pointer to the C GOutputVector that represents the OutputVector.
func OutputVectorNewFromC(native unsafe.Pointer) *OutputVector {
	return &OutputVector{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// FileIOStream is a representation of the C record GFileIOStream.
//
// since 2.22
type FileIOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStream that represents the FileIOStream.
func (recv *FileIOStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamNewFromC creates a new FileIOStream from a pointer to the C GFileIOStream that represents the FileIOStream.
func FileIOStreamNewFromC(native unsafe.Pointer) *FileIOStream {
	return &FileIOStream{native: native}
}

// IOStream is a representation of the C record GIOStream.
//
// since 2.22
type IOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStream that represents the IOStream.
func (recv *IOStream) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamNewFromC creates a new IOStream from a pointer to the C GIOStream that represents the IOStream.
func IOStreamNewFromC(native unsafe.Pointer) *IOStream {
	return &IOStream{native: native}
}

// Socket is a representation of the C record GSocket.
//
// since 2.22
type Socket struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocket that represents the Socket.
func (recv *Socket) ToC() unsafe.Pointer {
	return recv.native
}

// SocketNewFromC creates a new Socket from a pointer to the C GSocket that represents the Socket.
func SocketNewFromC(native unsafe.Pointer) *Socket {
	return &Socket{native: native}
}

// SocketClient is a representation of the C record GSocketClient.
//
// since 2.22
type SocketClient struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClient that represents the SocketClient.
func (recv *SocketClient) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientNewFromC creates a new SocketClient from a pointer to the C GSocketClient that represents the SocketClient.
func SocketClientNewFromC(native unsafe.Pointer) *SocketClient {
	return &SocketClient{native: native}
}

// SocketConnection is a representation of the C record GSocketConnection.
//
// since 2.22
type SocketConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnection that represents the SocketConnection.
func (recv *SocketConnection) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionNewFromC creates a new SocketConnection from a pointer to the C GSocketConnection that represents the SocketConnection.
func SocketConnectionNewFromC(native unsafe.Pointer) *SocketConnection {
	return &SocketConnection{native: native}
}

// SocketControlMessage is a representation of the C record GSocketControlMessage.
//
// since 2.22
type SocketControlMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessage that represents the SocketControlMessage.
func (recv *SocketControlMessage) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessageNewFromC creates a new SocketControlMessage from a pointer to the C GSocketControlMessage that represents the SocketControlMessage.
func SocketControlMessageNewFromC(native unsafe.Pointer) *SocketControlMessage {
	return &SocketControlMessage{native: native}
}

// SocketListener is a representation of the C record GSocketListener.
//
// since 2.22
type SocketListener struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListener that represents the SocketListener.
func (recv *SocketListener) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerNewFromC creates a new SocketListener from a pointer to the C GSocketListener that represents the SocketListener.
func SocketListenerNewFromC(native unsafe.Pointer) *SocketListener {
	return &SocketListener{native: native}
}

// SocketService is a representation of the C record GSocketService.
//
// since 2.22
type SocketService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketService that represents the SocketService.
func (recv *SocketService) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServiceNewFromC creates a new SocketService from a pointer to the C GSocketService that represents the SocketService.
func SocketServiceNewFromC(native unsafe.Pointer) *SocketService {
	return &SocketService{native: native}
}

// TcpConnection is a representation of the C record GTcpConnection.
//
// since 2.22
type TcpConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnection that represents the TcpConnection.
func (recv *TcpConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionNewFromC creates a new TcpConnection from a pointer to the C GTcpConnection that represents the TcpConnection.
func TcpConnectionNewFromC(native unsafe.Pointer) *TcpConnection {
	return &TcpConnection{native: native}
}

// ThreadedSocketService is a representation of the C record GThreadedSocketService.
//
// since 2.22
type ThreadedSocketService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketService that represents the ThreadedSocketService.
func (recv *ThreadedSocketService) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServiceNewFromC creates a new ThreadedSocketService from a pointer to the C GThreadedSocketService that represents the ThreadedSocketService.
func ThreadedSocketServiceNewFromC(native unsafe.Pointer) *ThreadedSocketService {
	return &ThreadedSocketService{native: native}
}

// UnixConnection is a representation of the C record GUnixConnection.
//
// since 2.22
type UnixConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnection that represents the UnixConnection.
func (recv *UnixConnection) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionNewFromC creates a new UnixConnection from a pointer to the C GUnixConnection that represents the UnixConnection.
func UnixConnectionNewFromC(native unsafe.Pointer) *UnixConnection {
	return &UnixConnection{native: native}
}

// AsyncInitable is a representation of the C interface GAsyncInitable.
//
// since 2.22
type AsyncInitable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncInitable that represents the AsyncInitable.
func (recv *AsyncInitable) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncInitableNewFromC creates a new AsyncInitable from a pointer to the C GAsyncInitable that represents the AsyncInitable.
func AsyncInitableNewFromC(native unsafe.Pointer) *AsyncInitable {
	return &AsyncInitable{native: native}
}

// Initable is a representation of the C interface GInitable.
//
// since 2.22
type Initable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitable that represents the Initable.
func (recv *Initable) ToC() unsafe.Pointer {
	return recv.native
}

// InitableNewFromC creates a new Initable from a pointer to the C GInitable that represents the Initable.
func InitableNewFromC(native unsafe.Pointer) *Initable {
	return &Initable{native: native}
}
