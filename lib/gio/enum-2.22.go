// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import glib "github.com/pekim/gobbi/lib/glib"

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

// ResolverErrorQuark is a wrapper around the C function g_resolver_error_quark.
func ResolverErrorQuark() glib.Quark {
	retC := C.g_resolver_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

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
