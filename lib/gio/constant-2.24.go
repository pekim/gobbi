// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// A key in the "trash" namespace.  When requested against
// items in `trash:///`, will return the date and time when the file
// was trashed. The format of the returned string is YYYY-MM-DDThh:mm:ss.
// Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
const FILE_ATTRIBUTE_TRASH_DELETION_DATE string = C.G_FILE_ATTRIBUTE_TRASH_DELETION_DATE

// A key in the "trash" namespace.  When requested against
// items in `trash:///`, will return the original path to the file before it
// was trashed. Corresponding #GFileAttributeType is
// %G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
const FILE_ATTRIBUTE_TRASH_ORIG_PATH string = C.G_FILE_ATTRIBUTE_TRASH_ORIG_PATH
