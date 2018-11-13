// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

/*
An error code used with %G_RESOURCE_ERROR in a #GError returned
from a #GResource routine.
*/
type ResourceError C.GResourceError

const (
	// no file was found at the requested path
	RESOURCE_ERROR_NOT_FOUND ResourceError = 0
	// unknown error
	RESOURCE_ERROR_INTERNAL ResourceError = 1
)

/*
Describes an event occurring on a #GSocketClient. See the
#GSocketClient::event signal for more details.

Additional values may be added to this type in the future.
*/
type SocketClientEvent C.GSocketClientEvent

const (
	// The client is doing a DNS lookup.
	SOCKET_CLIENT_RESOLVING SocketClientEvent = 0
	// The client has completed a DNS lookup.
	SOCKET_CLIENT_RESOLVED SocketClientEvent = 1
	/*
	   The client is connecting to a remote
	     host (either a proxy or the destination server).
	*/
	SOCKET_CLIENT_CONNECTING SocketClientEvent = 2
	/*
	   The client has connected to a remote
	     host.
	*/
	SOCKET_CLIENT_CONNECTED SocketClientEvent = 3
	/*
	   The client is negotiating
	     with a proxy to connect to the destination server.
	*/
	SOCKET_CLIENT_PROXY_NEGOTIATING SocketClientEvent = 4
	/*
	   The client has negotiated
	     with the proxy server.
	*/
	SOCKET_CLIENT_PROXY_NEGOTIATED SocketClientEvent = 5
	/*
	   The client is performing a
	     TLS handshake.
	*/
	SOCKET_CLIENT_TLS_HANDSHAKING SocketClientEvent = 6
	/*
	   The client has performed a
	     TLS handshake.
	*/
	SOCKET_CLIENT_TLS_HANDSHAKED SocketClientEvent = 7
	/*
	   The client is done with a particular
	     #GSocketConnectable.
	*/
	SOCKET_CLIENT_COMPLETE SocketClientEvent = 8
)
