// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// GEmblemOrigin is used to add information about the origin of the emblem
// to #GEmblem.
type EmblemOrigin C.GEmblemOrigin

const (
	// Emblem of unknown origin
	EMBLEM_ORIGIN_UNKNOWN EmblemOrigin = 0

	// Emblem adds device-specific information
	EMBLEM_ORIGIN_DEVICE EmblemOrigin = 1

	// Emblem depicts live metadata, such as "readonly"
	EMBLEM_ORIGIN_LIVEMETADATA EmblemOrigin = 2

	// Emblem comes from a user-defined tag, e.g. set by nautilus (in the future)
	EMBLEM_ORIGIN_TAG EmblemOrigin = 3
)
