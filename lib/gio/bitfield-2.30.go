// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Flags describing the behavior of a #GDBusInterfaceSkeleton instance.
type DBusInterfaceSkeletonFlags C.GDBusInterfaceSkeletonFlags

const (
	// No flags set.
	DBUS_INTERFACE_SKELETON_FLAGS_NONE DBusInterfaceSkeletonFlags = 0
	// Each method invocation is handled in
	// a thread dedicated to the invocation. This means that the method implementation can use blocking IO
	// without blocking any other part of the process. It also means that the method implementation must
	// use locking to access data structures used by other threads.
	DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD DBusInterfaceSkeletonFlags = 1
)

// Flags used when constructing a #GDBusObjectManagerClient.
type DBusObjectManagerClientFlags C.GDBusObjectManagerClientFlags

const (
	// No flags set.
	DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE DBusObjectManagerClientFlags = 0
	// If not set and the
	// manager is for a well-known name, then request the bus to launch
	// an owner for the name if no-one owns the name. This flag can only
	// be used in managers for well-known names.
	DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START DBusObjectManagerClientFlags = 1
)

// Flags for g_tls_database_verify_chain().
type TlsDatabaseVerifyFlags C.GTlsDatabaseVerifyFlags

const (
	// No verification flags
	TLS_DATABASE_VERIFY_NONE TlsDatabaseVerifyFlags = 0
)

// Various flags for the password.
type TlsPasswordFlags C.GTlsPasswordFlags

const (
	// No flags
	TLS_PASSWORD_NONE TlsPasswordFlags = 0
	// The password was wrong, and the user should retry.
	TLS_PASSWORD_RETRY TlsPasswordFlags = 2
	// Hint to the user that the password has been
	// wrong many times, and the user may not have many chances left.
	TLS_PASSWORD_MANY_TRIES TlsPasswordFlags = 4
	// Hint to the user that this is the last try to get
	// this password right.
	TLS_PASSWORD_FINAL_TRY TlsPasswordFlags = 8
)
