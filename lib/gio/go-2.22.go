// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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
// #include <stdlib.h>
import "C"

type DriveStartFlags int

const (
	DRIVE_START_NONE DriveStartFlags = 0
)

type SocketMsgFlags int

const (
	SOCKET_MSG_NONE      SocketMsgFlags = 0
	SOCKET_MSG_OOB       SocketMsgFlags = 1
	SOCKET_MSG_PEEK      SocketMsgFlags = 2
	SOCKET_MSG_DONTROUTE SocketMsgFlags = 4
)

// Unsupported : g_inet_address_new_any : return type :

// Unsupported : g_inet_address_new_from_bytes : return type :

// Unsupported : g_inet_address_new_from_string : return type :

// Unsupported : g_inet_address_new_loopback : return type :

// Unsupported : g_inet_socket_address_new : return type :

// Unsupported : g_network_address_new : return type :

// Unsupported : g_network_service_new : return type :

// Unsupported : g_socket_new : return type :

// Unsupported : g_socket_new_from_fd : return type :

// Unsupported : g_socket_address_new_from_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// Unsupported : g_socket_client_new : return type :

// Unsupported : g_socket_listener_new : return type :

// Unsupported : g_socket_service_new : return type :

// Unsupported : g_threaded_socket_service_new : return type :

// Unsupported : g_unix_fd_message_new : return type :

// Unsupported : g_unix_socket_address_new : return type :

const FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL string = "mountable::can-poll"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START string = "mountable::can-start"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED string = "mountable::can-start-degraded"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP string = "mountable::can-stop"
const FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC string = "mountable::is-media-check-automatic"
const FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE string = "mountable::start-stop-type"
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE string = "mountable::unix-device-file"

type DriveStartStopType int

const (
	DRIVE_START_STOP_TYPE_UNKNOWN   DriveStartStopType = 0
	DRIVE_START_STOP_TYPE_SHUTDOWN  DriveStartStopType = 1
	DRIVE_START_STOP_TYPE_NETWORK   DriveStartStopType = 2
	DRIVE_START_STOP_TYPE_MULTIDISK DriveStartStopType = 3
	DRIVE_START_STOP_TYPE_PASSWORD  DriveStartStopType = 4
)

type ResolverError int

const (
	RESOLVER_ERROR_NOT_FOUND         ResolverError = 0
	RESOLVER_ERROR_TEMPORARY_FAILURE ResolverError = 1
	RESOLVER_ERROR_INTERNAL          ResolverError = 2
)

// g_resolver_error_quark : return type :
type SocketFamily int

const (
	SOCKET_FAMILY_INVALID SocketFamily = 0
	SOCKET_FAMILY_UNIX    SocketFamily = 1
	SOCKET_FAMILY_IPV4    SocketFamily = 2
	SOCKET_FAMILY_IPV6    SocketFamily = 10
)

type SocketProtocol int

const (
	SOCKET_PROTOCOL_UNKNOWN SocketProtocol = -1
	SOCKET_PROTOCOL_DEFAULT SocketProtocol = 0
	SOCKET_PROTOCOL_TCP     SocketProtocol = 6
	SOCKET_PROTOCOL_UDP     SocketProtocol = 17
	SOCKET_PROTOCOL_SCTP    SocketProtocol = 132
)

type SocketType int

const (
	SOCKET_TYPE_INVALID   SocketType = 0
	SOCKET_TYPE_STREAM    SocketType = 1
	SOCKET_TYPE_DATAGRAM  SocketType = 2
	SOCKET_TYPE_SEQPACKET SocketType = 3
)

// Unsupported : g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_initable_newv : unsupported parameter parameters :

// Unsupported : g_resolver_error_quark : return type :

// Unsupported : g_srv_target_list_sort : return type :

// Unsupported : g_srv_target_new : return type :
