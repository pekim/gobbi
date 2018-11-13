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

// An enumeration for well-known message buses.
type BusType C.GBusType

const (
	// An alias for the message bus that activated the process, if any.
	BUS_TYPE_STARTER BusType = -1
	// Not a message bus.
	BUS_TYPE_NONE BusType = 0
	// The system-wide message bus.
	BUS_TYPE_SYSTEM BusType = 1
	// The login session message bus.
	BUS_TYPE_SESSION BusType = 2
)

// Enumeration describing different kinds of native credential types.
type CredentialsType C.GCredentialsType

const (
	// Indicates an invalid native credential type.
	CREDENTIALS_TYPE_INVALID CredentialsType = 0
	// The native credentials type is a struct ucred.
	CREDENTIALS_TYPE_LINUX_UCRED CredentialsType = 1
	// The native credentials type is a struct cmsgcred.
	CREDENTIALS_TYPE_FREEBSD_CMSGCRED CredentialsType = 2
	// The native credentials type is a struct sockpeercred. Added in 2.30.
	CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED CredentialsType = 3
	// The native credentials type is a ucred_t. Added in 2.40.
	CREDENTIALS_TYPE_SOLARIS_UCRED CredentialsType = 4
	// The native credentials type is a struct unpcbid.
	CREDENTIALS_TYPE_NETBSD_UNPCBID CredentialsType = 5
)

// Error codes for the %G_DBUS_ERROR error domain.
type DBusError C.GDBusError

const (
	/*
	   A generic error; "something went wrong" - see the error message for
	   more.
	*/
	DBUS_ERROR_FAILED DBusError = 0
	// There was not enough memory to complete an operation.
	DBUS_ERROR_NO_MEMORY DBusError = 1
	/*
	   The bus doesn't know how to launch a service to supply the bus name
	   you wanted.
	*/
	DBUS_ERROR_SERVICE_UNKNOWN DBusError = 2
	/*
	   The bus name you referenced doesn't exist (i.e. no application owns
	   it).
	*/
	DBUS_ERROR_NAME_HAS_NO_OWNER DBusError = 3
	// No reply to a message expecting one, usually means a timeout occurred.
	DBUS_ERROR_NO_REPLY DBusError = 4
	// Something went wrong reading or writing to a socket, for example.
	DBUS_ERROR_IO_ERROR DBusError = 5
	// A D-Bus bus address was malformed.
	DBUS_ERROR_BAD_ADDRESS DBusError = 6
	// Requested operation isn't supported (like ENOSYS on UNIX).
	DBUS_ERROR_NOT_SUPPORTED DBusError = 7
	// Some limited resource is exhausted.
	DBUS_ERROR_LIMITS_EXCEEDED DBusError = 8
	// Security restrictions don't allow doing what you're trying to do.
	DBUS_ERROR_ACCESS_DENIED DBusError = 9
	// Authentication didn't work.
	DBUS_ERROR_AUTH_FAILED DBusError = 10
	/*
	   Unable to connect to server (probably caused by ECONNREFUSED on a
	   socket).
	*/
	DBUS_ERROR_NO_SERVER DBusError = 11
	/*
	   Certain timeout errors, possibly ETIMEDOUT on a socket.  Note that
	   %G_DBUS_ERROR_NO_REPLY is used for message reply timeouts. Warning:
	   this is confusingly-named given that %G_DBUS_ERROR_TIMED_OUT also
	   exists. We can't fix it for compatibility reasons so just be
	   careful.
	*/
	DBUS_ERROR_TIMEOUT DBusError = 12
	// No network access (probably ENETUNREACH on a socket).
	DBUS_ERROR_NO_NETWORK DBusError = 13
	// Can't bind a socket since its address is in use (i.e. EADDRINUSE).
	DBUS_ERROR_ADDRESS_IN_USE DBusError = 14
	// The connection is disconnected and you're trying to use it.
	DBUS_ERROR_DISCONNECTED DBusError = 15
	// Invalid arguments passed to a method call.
	DBUS_ERROR_INVALID_ARGS DBusError = 16
	// Missing file.
	DBUS_ERROR_FILE_NOT_FOUND DBusError = 17
	// Existing file and the operation you're using does not silently overwrite.
	DBUS_ERROR_FILE_EXISTS DBusError = 18
	// Method name you invoked isn't known by the object you invoked it on.
	DBUS_ERROR_UNKNOWN_METHOD DBusError = 19
	/*
	   Certain timeout errors, e.g. while starting a service. Warning: this is
	   confusingly-named given that %G_DBUS_ERROR_TIMEOUT also exists. We
	   can't fix it for compatibility reasons so just be careful.
	*/
	DBUS_ERROR_TIMED_OUT DBusError = 20
	// Tried to remove or modify a match rule that didn't exist.
	DBUS_ERROR_MATCH_RULE_NOT_FOUND DBusError = 21
	// The match rule isn't syntactically valid.
	DBUS_ERROR_MATCH_RULE_INVALID DBusError = 22
	// While starting a new process, the exec() call failed.
	DBUS_ERROR_SPAWN_EXEC_FAILED DBusError = 23
	// While starting a new process, the fork() call failed.
	DBUS_ERROR_SPAWN_FORK_FAILED DBusError = 24
	// While starting a new process, the child exited with a status code.
	DBUS_ERROR_SPAWN_CHILD_EXITED DBusError = 25
	// While starting a new process, the child exited on a signal.
	DBUS_ERROR_SPAWN_CHILD_SIGNALED DBusError = 26
	// While starting a new process, something went wrong.
	DBUS_ERROR_SPAWN_FAILED DBusError = 27
	// We failed to setup the environment correctly.
	DBUS_ERROR_SPAWN_SETUP_FAILED DBusError = 28
	// We failed to setup the config parser correctly.
	DBUS_ERROR_SPAWN_CONFIG_INVALID DBusError = 29
	// Bus name was not valid.
	DBUS_ERROR_SPAWN_SERVICE_INVALID DBusError = 30
	// Service file not found in system-services directory.
	DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND DBusError = 31
	// Permissions are incorrect on the setuid helper.
	DBUS_ERROR_SPAWN_PERMISSIONS_INVALID DBusError = 32
	// Service file invalid (Name, User or Exec missing).
	DBUS_ERROR_SPAWN_FILE_INVALID DBusError = 33
	// Tried to get a UNIX process ID and it wasn't available.
	DBUS_ERROR_SPAWN_NO_MEMORY DBusError = 34
	// Tried to get a UNIX process ID and it wasn't available.
	DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN DBusError = 35
	// A type signature is not valid.
	DBUS_ERROR_INVALID_SIGNATURE DBusError = 36
	// A file contains invalid syntax or is otherwise broken.
	DBUS_ERROR_INVALID_FILE_CONTENT DBusError = 37
	// Asked for SELinux security context and it wasn't available.
	DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN DBusError = 38
	// Asked for ADT audit data and it wasn't available.
	DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN DBusError = 39
	// There's already an object with the requested object path.
	DBUS_ERROR_OBJECT_PATH_IN_USE DBusError = 40
	// Object you invoked a method on isn't known. Since 2.42
	DBUS_ERROR_UNKNOWN_OBJECT DBusError = 41
	// Interface you invoked a method on isn't known by the object. Since 2.42
	DBUS_ERROR_UNKNOWN_INTERFACE DBusError = 42
	// Property you tried to access isn't known by the object. Since 2.42
	DBUS_ERROR_UNKNOWN_PROPERTY DBusError = 43
	// Property you tried to set is read-only. Since 2.42
	DBUS_ERROR_PROPERTY_READ_ONLY DBusError = 44
)

// Enumeration used to describe the byte order of a D-Bus message.
type DBusMessageByteOrder C.GDBusMessageByteOrder

const (
	// The byte order is big endian.
	DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN DBusMessageByteOrder = 66
	// The byte order is little endian.
	DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN DBusMessageByteOrder = 108
)

// Header fields used in #GDBusMessage.
type DBusMessageHeaderField C.GDBusMessageHeaderField

const (
	// Not a valid header field.
	DBUS_MESSAGE_HEADER_FIELD_INVALID DBusMessageHeaderField = 0
	// The object path.
	DBUS_MESSAGE_HEADER_FIELD_PATH DBusMessageHeaderField = 1
	// The interface name.
	DBUS_MESSAGE_HEADER_FIELD_INTERFACE DBusMessageHeaderField = 2
	// The method or signal name.
	DBUS_MESSAGE_HEADER_FIELD_MEMBER DBusMessageHeaderField = 3
	// The name of the error that occurred.
	DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME DBusMessageHeaderField = 4
	// The serial number the message is a reply to.
	DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL DBusMessageHeaderField = 5
	// The name the message is intended for.
	DBUS_MESSAGE_HEADER_FIELD_DESTINATION DBusMessageHeaderField = 6
	// Unique name of the sender of the message (filled in by the bus).
	DBUS_MESSAGE_HEADER_FIELD_SENDER DBusMessageHeaderField = 7
	// The signature of the message body.
	DBUS_MESSAGE_HEADER_FIELD_SIGNATURE DBusMessageHeaderField = 8
	// The number of UNIX file descriptors that accompany the message.
	DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS DBusMessageHeaderField = 9
)

// Message types used in #GDBusMessage.
type DBusMessageType C.GDBusMessageType

const (
	// Message is of invalid type.
	DBUS_MESSAGE_TYPE_INVALID DBusMessageType = 0
	// Method call.
	DBUS_MESSAGE_TYPE_METHOD_CALL DBusMessageType = 1
	// Method reply.
	DBUS_MESSAGE_TYPE_METHOD_RETURN DBusMessageType = 2
	// Error reply.
	DBUS_MESSAGE_TYPE_ERROR DBusMessageType = 3
	// Signal emission.
	DBUS_MESSAGE_TYPE_SIGNAL DBusMessageType = 4
)

/*
The type of name used by a #GUnixSocketAddress.
%G_UNIX_SOCKET_ADDRESS_PATH indicates a traditional unix domain
socket bound to a filesystem path. %G_UNIX_SOCKET_ADDRESS_ANONYMOUS
indicates a socket not bound to any name (eg, a client-side socket,
or a socket created with socketpair()).

For abstract sockets, there are two incompatible ways of naming
them; the man pages suggest using the entire `struct sockaddr_un`
as the name, padding the unused parts of the %sun_path field with
zeroes; this corresponds to %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
However, many programs instead just use a portion of %sun_path, and
pass an appropriate smaller length to bind() or connect(). This is
%G_UNIX_SOCKET_ADDRESS_ABSTRACT.
*/
type UnixSocketAddressType C.GUnixSocketAddressType

const (
	// invalid
	UNIX_SOCKET_ADDRESS_INVALID UnixSocketAddressType = 0
	// anonymous
	UNIX_SOCKET_ADDRESS_ANONYMOUS UnixSocketAddressType = 1
	// a filesystem path
	UNIX_SOCKET_ADDRESS_PATH UnixSocketAddressType = 2
	// an abstract name
	UNIX_SOCKET_ADDRESS_ABSTRACT UnixSocketAddressType = 3
	/*
	   an abstract name, 0-padded
	     to the full length of a unix socket name
	*/
	UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED UnixSocketAddressType = 4
)
