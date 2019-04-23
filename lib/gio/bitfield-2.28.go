// This is a generated file - DO NOT EDIT
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

type ApplicationFlags C.GApplicationFlags

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

// Blacklisted : GIOStreamSpliceFlags

// Blacklisted : GTlsCertificateFlags
