// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// A key in the "standard" namespace for checking if a file is
// volatile. This is meant for opaque, non-POSIX-like backends to
// indicate that the URI is not persistent. Applications should look
// at #G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET for the persistent URI.
//
// Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
const FILE_ATTRIBUTE_STANDARD_IS_VOLATILE string = C.G_FILE_ATTRIBUTE_STANDARD_IS_VOLATILE
