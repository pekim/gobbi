// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// Enumeration describing how a drive can be started/stopped.
type DriveStartStopType C.GDriveStartStopType

const (
	// Unknown or drive doesn't support
	// start/stop.
	DRIVE_START_STOP_TYPE_UNKNOWN DriveStartStopType = 0

	// The stop method will physically
	// shut down the drive and e.g. power down the port the drive is
	// attached to.
	DRIVE_START_STOP_TYPE_SHUTDOWN DriveStartStopType = 1

	// The start/stop methods are used
	// for connecting/disconnect to the drive over the network.
	DRIVE_START_STOP_TYPE_NETWORK DriveStartStopType = 2

	// The start/stop methods will
	// assemble/disassemble a virtual drive from several physical
	// drives.
	DRIVE_START_STOP_TYPE_MULTIDISK DriveStartStopType = 3

	// The start/stop methods will
	// unlock/lock the disk (for example using the ATA <quote>SECURITY
	// UNLOCK DEVICE</quote> command)
	DRIVE_START_STOP_TYPE_PASSWORD DriveStartStopType = 4
)

// An error code used with %G_RESOLVER_ERROR in a #GError returned
// from a #GResolver routine.
type ResolverError C.GResolverError

const (
	// the requested name/address/service was not
	// found
	RESOLVER_ERROR_NOT_FOUND ResolverError = 0

	// the requested information could not
	// be looked up due to a network error or similar problem
	RESOLVER_ERROR_TEMPORARY_FAILURE ResolverError = 1

	// unknown error
	RESOLVER_ERROR_INTERNAL ResolverError = 2
)

// The protocol family of a #GSocketAddress. (These values are
// identical to the system defines %AF_INET, %AF_INET6 and %AF_UNIX,
// if available.)
type SocketFamily C.GSocketFamily

const (
	// no address family
	SOCKET_FAMILY_INVALID SocketFamily = 0

	// the UNIX domain family
	SOCKET_FAMILY_UNIX SocketFamily = 1

	// the IPv4 family
	SOCKET_FAMILY_IPV4 SocketFamily = 2

	// the IPv6 family
	SOCKET_FAMILY_IPV6 SocketFamily = 10
)

// A protocol identifier is specified when creating a #GSocket, which is a
// family/type specific identifier, where 0 means the default protocol for
// the particular family/type.
//
// This enum contains a set of commonly available and used protocols. You
// can also pass any other identifiers handled by the platform in order to
// use protocols not listed here.
type SocketProtocol C.GSocketProtocol

const (
	// The protocol type is unknown
	SOCKET_PROTOCOL_UNKNOWN SocketProtocol = -1

	// The default protocol for the family/type
	SOCKET_PROTOCOL_DEFAULT SocketProtocol = 0

	// TCP over IP
	SOCKET_PROTOCOL_TCP SocketProtocol = 6

	// UDP over IP
	SOCKET_PROTOCOL_UDP SocketProtocol = 17

	// SCTP over IP
	SOCKET_PROTOCOL_SCTP SocketProtocol = 132
)

// Flags used when creating a #GSocket. Some protocols may not implement
// all the socket types.
type SocketType C.GSocketType

const (
	// Type unknown or wrong
	SOCKET_TYPE_INVALID SocketType = 0

	// Reliable connection-based byte streams (e.g. TCP).
	SOCKET_TYPE_STREAM SocketType = 1

	// Connectionless, unreliable datagram passing.
	// (e.g. UDP)
	SOCKET_TYPE_DATAGRAM SocketType = 2

	// Reliable connection-based passing of datagrams
	// of fixed maximum length (e.g. SCTP).
	SOCKET_TYPE_SEQPACKET SocketType = 3
)
