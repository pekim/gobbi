// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Describes an event occurring on a #GSocketListener. See the
// #GSocketListener::event signal for more details.
//
// Additional values may be added to this type in the future.
type SocketListenerEvent C.GSocketListenerEvent

const (
	// The listener is about to bind a socket.
	SOCKET_LISTENER_BINDING SocketListenerEvent = 0

	// The listener has bound a socket.
	SOCKET_LISTENER_BOUND SocketListenerEvent = 1

	// The listener is about to start
	// listening on this socket.
	SOCKET_LISTENER_LISTENING SocketListenerEvent = 2

	// The listener is now listening on
	// this socket.
	SOCKET_LISTENER_LISTENED SocketListenerEvent = 3
)
