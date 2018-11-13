// This is a generated file - DO NOT EDIT
// +build gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// A key in the "preview" namespace for getting a #GIcon that can be
// used to get preview of the file. For example, it may be a low
// resolution thumbnail without metadata. Corresponding
// #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.  The value
// for this key should contain a #GIcon.
const FILE_ATTRIBUTE_PREVIEW_ICON string = C.G_FILE_ATTRIBUTE_PREVIEW_ICON

// A key in the "standard" namespace for getting the amount of disk space
// that is consumed by the file (in bytes).  This will generally be larger
// than the file size (due to block size overhead) but can occasionally be
// smaller (for example, for sparse files).
// Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.
const FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE string = C.G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE
