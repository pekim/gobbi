// This is a generated file - DO NOT EDIT
// +build gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

type BusNameOwnerFlags C.GBusNameOwnerFlags

const (
	BUS_NAME_OWNER_FLAGS_NONE              BusNameOwnerFlags = 0
	BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT BusNameOwnerFlags = 1
	BUS_NAME_OWNER_FLAGS_REPLACE           BusNameOwnerFlags = 2
	BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE      BusNameOwnerFlags = 4
)

type BusNameWatcherFlags C.GBusNameWatcherFlags

const (
	BUS_NAME_WATCHER_FLAGS_NONE       BusNameWatcherFlags = 0
	BUS_NAME_WATCHER_FLAGS_AUTO_START BusNameWatcherFlags = 1
)

type DBusCallFlags C.GDBusCallFlags

const (
	DBUS_CALL_FLAGS_NONE                            DBusCallFlags = 0
	DBUS_CALL_FLAGS_NO_AUTO_START                   DBusCallFlags = 1
	DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusCallFlags = 2
)

type DBusCapabilityFlags C.GDBusCapabilityFlags

const (
	DBUS_CAPABILITY_FLAGS_NONE            DBusCapabilityFlags = 0
	DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING DBusCapabilityFlags = 1
)

type DBusConnectionFlags C.GDBusConnectionFlags

const (
	DBUS_CONNECTION_FLAGS_NONE                           DBusConnectionFlags = 0
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT          DBusConnectionFlags = 1
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER          DBusConnectionFlags = 2
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusConnectionFlags = 4
	DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION         DBusConnectionFlags = 8
	DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING       DBusConnectionFlags = 16
)

type DBusMessageFlags C.GDBusMessageFlags

const (
	DBUS_MESSAGE_FLAGS_NONE                            DBusMessageFlags = 0
	DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED               DBusMessageFlags = 1
	DBUS_MESSAGE_FLAGS_NO_AUTO_START                   DBusMessageFlags = 2
	DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusMessageFlags = 4
)

type DBusPropertyInfoFlags C.GDBusPropertyInfoFlags

const (
	DBUS_PROPERTY_INFO_FLAGS_NONE     DBusPropertyInfoFlags = 0
	DBUS_PROPERTY_INFO_FLAGS_READABLE DBusPropertyInfoFlags = 1
	DBUS_PROPERTY_INFO_FLAGS_WRITABLE DBusPropertyInfoFlags = 2
)

type DBusProxyFlags C.GDBusProxyFlags

const (
	DBUS_PROXY_FLAGS_NONE                              DBusProxyFlags = 0
	DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES            DBusProxyFlags = 1
	DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS            DBusProxyFlags = 2
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START                 DBusProxyFlags = 4
	DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES        DBusProxyFlags = 8
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION DBusProxyFlags = 16
)

type DBusSendMessageFlags C.GDBusSendMessageFlags

const (
	DBUS_SEND_MESSAGE_FLAGS_NONE            DBusSendMessageFlags = 0
	DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL DBusSendMessageFlags = 1
)

type DBusServerFlags C.GDBusServerFlags

const (
	DBUS_SERVER_FLAGS_NONE                           DBusServerFlags = 0
	DBUS_SERVER_FLAGS_RUN_IN_THREAD                  DBusServerFlags = 1
	DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusServerFlags = 2
)

type DBusSignalFlags C.GDBusSignalFlags

const (
	DBUS_SIGNAL_FLAGS_NONE                 DBusSignalFlags = 0
	DBUS_SIGNAL_FLAGS_NO_MATCH_RULE        DBusSignalFlags = 1
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE DBusSignalFlags = 2
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH      DBusSignalFlags = 4
)

type DBusSubtreeFlags C.GDBusSubtreeFlags

const (
	DBUS_SUBTREE_FLAGS_NONE                           DBusSubtreeFlags = 0
	DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES DBusSubtreeFlags = 1
)
