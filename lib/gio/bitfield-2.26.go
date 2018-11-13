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

// Flags used in g_bus_own_name().
type BusNameOwnerFlags C.GBusNameOwnerFlags

const (
	// No flags set.
	BUS_NAME_OWNER_FLAGS_NONE BusNameOwnerFlags = 0

	// Allow another message bus connection to claim the name.
	BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT BusNameOwnerFlags = 1

	// If another message bus connection owns the name and have
	// specified #G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT, then take the name from the other connection.
	BUS_NAME_OWNER_FLAGS_REPLACE BusNameOwnerFlags = 2

	// If another message bus connection owns the name, immediately
	// return an error from g_bus_own_name() rather than entering the waiting queue for that name. (Since 2.54)
	BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE BusNameOwnerFlags = 4
)

// Flags used in g_bus_watch_name().
type BusNameWatcherFlags C.GBusNameWatcherFlags

const (
	// No flags set.
	BUS_NAME_WATCHER_FLAGS_NONE BusNameWatcherFlags = 0

	// If no-one owns the name when
	// beginning to watch the name, ask the bus to launch an owner for the
	// name.
	BUS_NAME_WATCHER_FLAGS_AUTO_START BusNameWatcherFlags = 1
)

// Flags used in g_dbus_connection_call() and similar APIs.
type DBusCallFlags C.GDBusCallFlags

const (
	// No flags set.
	DBUS_CALL_FLAGS_NONE DBusCallFlags = 0

	// The bus must not launch
	// an owner for the destination name in response to this method
	// invocation.
	DBUS_CALL_FLAGS_NO_AUTO_START DBusCallFlags = 1

	// the caller is prepared to
	// wait for interactive authorization. Since 2.46.
	DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusCallFlags = 2
)

// Capabilities negotiated with the remote peer.
type DBusCapabilityFlags C.GDBusCapabilityFlags

const (
	// No flags set.
	DBUS_CAPABILITY_FLAGS_NONE DBusCapabilityFlags = 0

	// The connection
	// supports exchanging UNIX file descriptors with the remote peer.
	DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING DBusCapabilityFlags = 1
)

// Flags used when creating a new #GDBusConnection.
type DBusConnectionFlags C.GDBusConnectionFlags

const (
	// No flags set.
	DBUS_CONNECTION_FLAGS_NONE DBusConnectionFlags = 0

	// Perform authentication against server.
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT DBusConnectionFlags = 1

	// Perform authentication against client.
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER DBusConnectionFlags = 2

	// When
	// authenticating as a server, allow the anonymous authentication
	// method.
	DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusConnectionFlags = 4

	// Pass this flag if connecting to a peer that is a
	// message bus. This means that the Hello() method will be invoked as part of the connection setup.
	DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION DBusConnectionFlags = 8

	// If set, processing of D-Bus messages is
	// delayed until g_dbus_connection_start_message_processing() is called.
	DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING DBusConnectionFlags = 16
)

// Message flags used in #GDBusMessage.
type DBusMessageFlags C.GDBusMessageFlags

const (
	// No flags set.
	DBUS_MESSAGE_FLAGS_NONE DBusMessageFlags = 0

	// A reply is not expected.
	DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED DBusMessageFlags = 1

	// The bus must not launch an
	// owner for the destination name in response to this message.
	DBUS_MESSAGE_FLAGS_NO_AUTO_START DBusMessageFlags = 2

	// If set on a method
	// call, this flag means that the caller is prepared to wait for interactive
	// authorization. Since 2.46.
	DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION DBusMessageFlags = 4
)

// Flags describing the access control of a D-Bus property.
type DBusPropertyInfoFlags C.GDBusPropertyInfoFlags

const (
	// No flags set.
	DBUS_PROPERTY_INFO_FLAGS_NONE DBusPropertyInfoFlags = 0

	// Property is readable.
	DBUS_PROPERTY_INFO_FLAGS_READABLE DBusPropertyInfoFlags = 1

	// Property is writable.
	DBUS_PROPERTY_INFO_FLAGS_WRITABLE DBusPropertyInfoFlags = 2
)

// Flags used when constructing an instance of a #GDBusProxy derived class.
type DBusProxyFlags C.GDBusProxyFlags

const (
	// No flags set.
	DBUS_PROXY_FLAGS_NONE DBusProxyFlags = 0

	// Don't load properties.
	DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES DBusProxyFlags = 1

	// Don't connect to signals on the remote object.
	DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS DBusProxyFlags = 2

	// If the proxy is for a well-known name,
	// do not ask the bus to launch an owner during proxy initialization or a method call.
	// This flag is only meaningful in proxies for well-known names.
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START DBusProxyFlags = 4

	// If set, the property value for any __invalidated property__ will be (asynchronously) retrieved upon receiving the [`PropertiesChanged`](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-properties) D-Bus signal and the property will not cause emission of the #GDBusProxy::g-properties-changed signal. When the value is received the #GDBusProxy::g-properties-changed signal is emitted for the property along with the retrieved value. Since 2.32.
	DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES DBusProxyFlags = 8

	// If the proxy is for a well-known name,
	// do not ask the bus to launch an owner during proxy initialization, but allow it to be
	// autostarted by a method call. This flag is only meaningful in proxies for well-known names,
	// and only if %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START is not also specified.
	DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION DBusProxyFlags = 16
)

// Flags used when sending #GDBusMessages on a #GDBusConnection.
type DBusSendMessageFlags C.GDBusSendMessageFlags

const (
	// No flags set.
	DBUS_SEND_MESSAGE_FLAGS_NONE DBusSendMessageFlags = 0

	// Do not automatically
	// assign a serial number from the #GDBusConnection object when
	// sending a message.
	DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL DBusSendMessageFlags = 1
)

// Flags used when creating a #GDBusServer.
type DBusServerFlags C.GDBusServerFlags

const (
	// No flags set.
	DBUS_SERVER_FLAGS_NONE DBusServerFlags = 0

	// All #GDBusServer::new-connection
	// signals will run in separated dedicated threads (see signal for
	// details).
	DBUS_SERVER_FLAGS_RUN_IN_THREAD DBusServerFlags = 1

	// Allow the anonymous
	// authentication method.
	DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS DBusServerFlags = 2
)

// Flags used when subscribing to signals via g_dbus_connection_signal_subscribe().
type DBusSignalFlags C.GDBusSignalFlags

const (
	// No flags set.
	DBUS_SIGNAL_FLAGS_NONE DBusSignalFlags = 0

	// Don't actually send the AddMatch
	// D-Bus call for this signal subscription.  This gives you more control
	// over which match rules you add (but you must add them manually).
	DBUS_SIGNAL_FLAGS_NO_MATCH_RULE DBusSignalFlags = 1

	// Match first arguments that
	// contain a bus or interface name with the given namespace.
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE DBusSignalFlags = 2

	// Match first arguments that
	// contain an object path that is either equivalent to the given path,
	// or one of the paths is a subpath of the other.
	DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH DBusSignalFlags = 4
)

// Flags passed to g_dbus_connection_register_subtree().
type DBusSubtreeFlags C.GDBusSubtreeFlags

const (
	// No flags set.
	DBUS_SUBTREE_FLAGS_NONE DBusSubtreeFlags = 0

	// Method calls to objects not in the enumerated range
	// will still be dispatched. This is useful if you want
	// to dynamically spawn objects in the subtree.
	DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES DBusSubtreeFlags = 1
)
