// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Flags for use with g_io_module_scope_new().
type IOModuleScopeFlags C.GIOModuleScopeFlags

const (
	// No module scan flags
	IO_MODULE_SCOPE_NONE IOModuleScopeFlags = 0
	// When using this scope to load or
	// scan modules, automatically block a modules which has the same base
	// basename as previously loaded module.
	IO_MODULE_SCOPE_BLOCK_DUPLICATES IOModuleScopeFlags = 1
)

// Flags for g_tls_database_lookup_certificate_for_handle(),
// g_tls_database_lookup_certificate_issuer(),
// and g_tls_database_lookup_certificates_issued_by().
type TlsDatabaseLookupFlags C.GTlsDatabaseLookupFlags

const (
	// No lookup flags
	TLS_DATABASE_LOOKUP_NONE TlsDatabaseLookupFlags = 0
	// Restrict lookup to certificates that have
	// a private key.
	TLS_DATABASE_LOOKUP_KEYPAIR TlsDatabaseLookupFlags = 1
)

// #GTlsInteractionResult is returned by various functions in #GTlsInteraction
// when finishing an interaction request.
type TlsInteractionResult C.GTlsInteractionResult

const (
	// The interaction was unhandled (i.e. not
	// implemented).
	TLS_INTERACTION_UNHANDLED TlsInteractionResult = 0
	// The interaction completed, and resulting data
	// is available.
	TLS_INTERACTION_HANDLED TlsInteractionResult = 1
	// The interaction has failed, or was cancelled.
	// and the operation should be aborted.
	TLS_INTERACTION_FAILED TlsInteractionResult = 2
)
