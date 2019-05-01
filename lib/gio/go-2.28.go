// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

type ApplicationFlags int

const (
	APPLICATION_FLAGS_NONE           ApplicationFlags = 0
	APPLICATION_IS_SERVICE           ApplicationFlags = 1
	APPLICATION_IS_LAUNCHER          ApplicationFlags = 2
	APPLICATION_HANDLES_OPEN         ApplicationFlags = 4
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = 8
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags = 16
	APPLICATION_NON_UNIQUE           ApplicationFlags = 32
	APPLICATION_CAN_OVERRIDE_APP_ID  ApplicationFlags = 64
)

type IOStreamSpliceFlags int

const (
	IO_STREAM_SPLICE_NONE          IOStreamSpliceFlags = 0
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1
	IO_STREAM_SPLICE_CLOSE_STREAM2 IOStreamSpliceFlags = 2
	IO_STREAM_SPLICE_WAIT_FOR_BOTH IOStreamSpliceFlags = 4
)

type TlsCertificateFlags int

const (
	TLS_CERTIFICATE_UNKNOWN_CA    TlsCertificateFlags = 1
	TLS_CERTIFICATE_BAD_IDENTITY  TlsCertificateFlags = 2
	TLS_CERTIFICATE_NOT_ACTIVATED TlsCertificateFlags = 4
	TLS_CERTIFICATE_EXPIRED       TlsCertificateFlags = 8
	TLS_CERTIFICATE_REVOKED       TlsCertificateFlags = 16
	TLS_CERTIFICATE_INSECURE      TlsCertificateFlags = 32
	TLS_CERTIFICATE_GENERIC_ERROR TlsCertificateFlags = 64
	TLS_CERTIFICATE_VALIDATE_ALL  TlsCertificateFlags = 127
)

// Unsupported : g_application_new : return type :

// Unsupported : g_simple_action_new : return type :

// Unsupported : g_simple_action_new_stateful : return type :

// Unsupported : g_simple_action_group_new : return type :

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_tcp_wrapper_connection_new : return type :

// Unsupported : g_tls_certificate_new_from_file : return type :

// Unsupported : g_tls_certificate_new_from_files : return type :

// Unsupported : g_tls_certificate_new_from_pem : return type :

type TlsAuthenticationMode int

const (
	TLS_AUTHENTICATION_NONE      TlsAuthenticationMode = 0
	TLS_AUTHENTICATION_REQUESTED TlsAuthenticationMode = 1
	TLS_AUTHENTICATION_REQUIRED  TlsAuthenticationMode = 2
)

type TlsError int

const (
	TLS_ERROR_UNAVAILABLE          TlsError = 0
	TLS_ERROR_MISC                 TlsError = 1
	TLS_ERROR_BAD_CERTIFICATE      TlsError = 2
	TLS_ERROR_NOT_TLS              TlsError = 3
	TLS_ERROR_HANDSHAKE            TlsError = 4
	TLS_ERROR_CERTIFICATE_REQUIRED TlsError = 5
	TLS_ERROR_EOF                  TlsError = 6
)

// g_tls_error_quark : return type :
type TlsRehandshakeMode int

const (
	TLS_REHANDSHAKE_NEVER    TlsRehandshakeMode = 0
	TLS_REHANDSHAKE_SAFELY   TlsRehandshakeMode = 1
	TLS_REHANDSHAKE_UNSAFELY TlsRehandshakeMode = 2
)

// Unsupported : g_app_info_get_fallback_for_type : return type :

// Unsupported : g_app_info_get_recommended_for_type : return type :

// Unsupported : g_memory_settings_backend_new : return type :

// Unsupported : g_null_settings_backend_new : return type :

// Unsupported : g_pollable_source_new : return type :

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_tls_backend_get_default : return type :

// Unsupported : g_tls_client_connection_new : return type :

// Unsupported : g_tls_error_quark : return type :

// Unsupported : g_tls_server_connection_new : return type :
