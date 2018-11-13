// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be polled.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started
degraded.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be stopped.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)
is automatically polled for media.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC string = C.G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC

/*
A key in the "mountable" namespace for getting the #GDriveStartStopType.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE

/*
A key in the "mountable" namespace for getting the unix device file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE
