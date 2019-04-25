// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

type ResourceError int

const (
	RESOURCE_ERROR_NOT_FOUND ResourceError = 0
	RESOURCE_ERROR_INTERNAL  ResourceError = 1
)

// ResourceErrorQuark is a wrapper around the C function g_resource_error_quark.
func ResourceErrorQuark() glib.Quark {
	retC := C.g_resource_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type SocketClientEvent int

const (
	SOCKET_CLIENT_RESOLVING         SocketClientEvent = 0
	SOCKET_CLIENT_RESOLVED          SocketClientEvent = 1
	SOCKET_CLIENT_CONNECTING        SocketClientEvent = 2
	SOCKET_CLIENT_CONNECTED         SocketClientEvent = 3
	SOCKET_CLIENT_PROXY_NEGOTIATING SocketClientEvent = 4
	SOCKET_CLIENT_PROXY_NEGOTIATED  SocketClientEvent = 5
	SOCKET_CLIENT_TLS_HANDSHAKING   SocketClientEvent = 6
	SOCKET_CLIENT_TLS_HANDSHAKED    SocketClientEvent = 7
	SOCKET_CLIENT_COMPLETE          SocketClientEvent = 8
)
