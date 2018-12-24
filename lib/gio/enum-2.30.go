// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

type IOModuleScopeFlags C.GIOModuleScopeFlags

const (
	IO_MODULE_SCOPE_NONE             IOModuleScopeFlags = 0
	IO_MODULE_SCOPE_BLOCK_DUPLICATES IOModuleScopeFlags = 1
)

type TlsDatabaseLookupFlags C.GTlsDatabaseLookupFlags

const (
	TLS_DATABASE_LOOKUP_NONE    TlsDatabaseLookupFlags = 0
	TLS_DATABASE_LOOKUP_KEYPAIR TlsDatabaseLookupFlags = 1
)

type TlsInteractionResult C.GTlsInteractionResult

const (
	TLS_INTERACTION_UNHANDLED TlsInteractionResult = 0
	TLS_INTERACTION_HANDLED   TlsInteractionResult = 1
	TLS_INTERACTION_FAILED    TlsInteractionResult = 2
)
