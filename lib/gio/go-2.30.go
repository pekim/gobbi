// This is a generated file - DO NOT EDIT
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56
// +build gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Unsupported : g_dbus_object_manager_client_new_finish : return type :

// Unsupported : g_dbus_object_manager_client_new_for_bus_finish : return type :

// Unsupported : g_dbus_object_manager_client_new_for_bus_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_client_new_sync : unsupported parameter get_proxy_type_func : no type generator for DBusProxyTypeFunc (GDBusProxyTypeFunc) for param get_proxy_type_func

// Unsupported : g_dbus_object_manager_server_new : return type :

// Unsupported : g_dbus_object_proxy_new : return type :

// Unsupported : g_dbus_object_skeleton_new : return type :

// Unsupported : g_tls_password_new : return type :

const NETWORK_MONITOR_EXTENSION_POINT_NAME string = "gio-network-monitor"

type IOModuleScopeFlags int

const (
	IO_MODULE_SCOPE_NONE             IOModuleScopeFlags = 0
	IO_MODULE_SCOPE_BLOCK_DUPLICATES IOModuleScopeFlags = 1
)

type TlsDatabaseLookupFlags int

const (
	TLS_DATABASE_LOOKUP_NONE    TlsDatabaseLookupFlags = 0
	TLS_DATABASE_LOOKUP_KEYPAIR TlsDatabaseLookupFlags = 1
)

type TlsInteractionResult int

const (
	TLS_INTERACTION_UNHANDLED TlsInteractionResult = 0
	TLS_INTERACTION_HANDLED   TlsInteractionResult = 1
	TLS_INTERACTION_FAILED    TlsInteractionResult = 2
)

// Unsupported : g_dbus_gvalue_to_gvariant : return type :

// Unsupported : g_io_modules_load_all_in_directory_with_scope : return type :

// Unsupported : g_tls_file_database_new : return type :
