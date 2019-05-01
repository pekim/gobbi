// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

type DBusInterfaceSkeletonFlags int

const (
	DBUS_INTERFACE_SKELETON_FLAGS_NONE                                DBusInterfaceSkeletonFlags = 0
	DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD DBusInterfaceSkeletonFlags = 1
)

type DBusObjectManagerClientFlags int

const (
	DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE              DBusObjectManagerClientFlags = 0
	DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START DBusObjectManagerClientFlags = 1
)

type TlsDatabaseVerifyFlags int

const (
	TLS_DATABASE_VERIFY_NONE TlsDatabaseVerifyFlags = 0
)

type TlsPasswordFlags int

const (
	TLS_PASSWORD_NONE       TlsPasswordFlags = 0
	TLS_PASSWORD_RETRY      TlsPasswordFlags = 2
	TLS_PASSWORD_MANY_TRIES TlsPasswordFlags = 4
	TLS_PASSWORD_FINAL_TRY  TlsPasswordFlags = 8
)
