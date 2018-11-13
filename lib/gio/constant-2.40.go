// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

/*
A key in the "thumbnail" namespace for checking whether the thumbnail is outdated.
This attribute is %TRUE if the thumbnail is up-to-date with the file it represents,
and %FALSE if the file has been modified since the thumbnail was generated.

If %G_FILE_ATTRIBUTE_THUMBNAILING_FAILED is %TRUE and this attribute is %FALSE,
it indicates that thumbnailing may be attempted again and may succeed.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_THUMBNAIL_IS_VALID string = C.G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID
