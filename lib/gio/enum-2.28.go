// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

type TlsAuthenticationMode C.GTlsAuthenticationMode

const (
	TLS_AUTHENTICATION_NONE      TlsAuthenticationMode = 0
	TLS_AUTHENTICATION_REQUESTED TlsAuthenticationMode = 1
	TLS_AUTHENTICATION_REQUIRED  TlsAuthenticationMode = 2
)

type TlsError C.GTlsError

const (
	TLS_ERROR_UNAVAILABLE          TlsError = 0
	TLS_ERROR_MISC                 TlsError = 1
	TLS_ERROR_BAD_CERTIFICATE      TlsError = 2
	TLS_ERROR_NOT_TLS              TlsError = 3
	TLS_ERROR_HANDSHAKE            TlsError = 4
	TLS_ERROR_CERTIFICATE_REQUIRED TlsError = 5
	TLS_ERROR_EOF                  TlsError = 6
)

type TlsRehandshakeMode C.GTlsRehandshakeMode

const (
	TLS_REHANDSHAKE_NEVER    TlsRehandshakeMode = 0
	TLS_REHANDSHAKE_SAFELY   TlsRehandshakeMode = 1
	TLS_REHANDSHAKE_UNSAFELY TlsRehandshakeMode = 2
)