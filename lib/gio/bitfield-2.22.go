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

// Flags used when starting a drive.
type DriveStartFlags C.GDriveStartFlags

const (
	// No flags set.
	DRIVE_START_NONE DriveStartFlags = 0
)

/*
Flags used in g_socket_receive_message() and g_socket_send_message().
The flags listed in the enum are some commonly available flags, but the
values used for them are the same as on the platform, and any other flags
are passed in/out as is. So to use a platform specific flag, just include
the right system header and pass in the flag.
*/
type SocketMsgFlags C.GSocketMsgFlags

const (
	// No flags.
	SOCKET_MSG_NONE SocketMsgFlags = 0
	// Request to send/receive out of band data.
	SOCKET_MSG_OOB SocketMsgFlags = 1
	/*
	   Read data from the socket without removing it from
	       the queue.
	*/
	SOCKET_MSG_PEEK SocketMsgFlags = 2
	/*
	   Don't use a gateway to send out the packet,
	       only send to hosts on directly connected networks.
	*/
	SOCKET_MSG_DONTROUTE SocketMsgFlags = 4
)
