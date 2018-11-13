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

// The client authentication mode for a #GTlsServerConnection.
type TlsAuthenticationMode C.GTlsAuthenticationMode

const (
	// client authentication not required
	TLS_AUTHENTICATION_NONE TlsAuthenticationMode = 0

	// client authentication is requested
	TLS_AUTHENTICATION_REQUESTED TlsAuthenticationMode = 1

	// client authentication is required
	TLS_AUTHENTICATION_REQUIRED TlsAuthenticationMode = 2
)

// An error code used with %G_TLS_ERROR in a #GError returned from a
// TLS-related routine.
type TlsError C.GTlsError

const (
	// No TLS provider is available
	TLS_ERROR_UNAVAILABLE TlsError = 0

	// Miscellaneous TLS error
	TLS_ERROR_MISC TlsError = 1

	// A certificate could not be parsed
	TLS_ERROR_BAD_CERTIFICATE TlsError = 2

	// The TLS handshake failed because the
	// peer does not seem to be a TLS server.
	TLS_ERROR_NOT_TLS TlsError = 3

	// The TLS handshake failed because the
	// peer's certificate was not acceptable.
	TLS_ERROR_HANDSHAKE TlsError = 4

	// The TLS handshake failed because
	// the server requested a client-side certificate, but none was
	// provided. See g_tls_connection_set_certificate().
	TLS_ERROR_CERTIFICATE_REQUIRED TlsError = 5

	// The TLS connection was closed without proper
	// notice, which may indicate an attack. See
	// g_tls_connection_set_require_close_notify().
	TLS_ERROR_EOF TlsError = 6
)

// When to allow rehandshaking. See
// g_tls_connection_set_rehandshake_mode().
type TlsRehandshakeMode C.GTlsRehandshakeMode

const (
	// Never allow rehandshaking
	TLS_REHANDSHAKE_NEVER TlsRehandshakeMode = 0

	// Allow safe rehandshaking only
	TLS_REHANDSHAKE_SAFELY TlsRehandshakeMode = 1

	// Allow unsafe rehandshaking
	TLS_REHANDSHAKE_UNSAFELY TlsRehandshakeMode = 2
)
