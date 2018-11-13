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

// GResourceFlags give information about a particular file inside a resource
// bundle.
type ResourceFlags C.GResourceFlags

const (
	// No flags set.
	RESOURCE_FLAGS_NONE ResourceFlags = 0
	// The file is compressed.
	RESOURCE_FLAGS_COMPRESSED ResourceFlags = 1
)

// GResourceLookupFlags determine how resource path lookups are handled.
type ResourceLookupFlags C.GResourceLookupFlags

const (
	// No flags set.
	RESOURCE_LOOKUP_FLAGS_NONE ResourceLookupFlags = 0
)
