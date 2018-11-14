# `gio` enums

## `BusType`

An enumeration for well-known message buses.

C - `GBusType`

### BUS_TYPE_STARTER = -1
An alias for the message bus that activated the process, if any.

### BUS_TYPE_NONE = 0
Not a message bus.

### BUS_TYPE_SYSTEM = 1
The system-wide message bus.

### BUS_TYPE_SESSION = 2
The login session message bus.


## `ConverterResult`

Results returned from g_converter_convert().

C - `GConverterResult`

### CONVERTER_ERROR = 0
There was an error during conversion.

### CONVERTER_CONVERTED = 1
Some data was consumed or produced

### CONVERTER_FINISHED = 2
The conversion is finished

### CONVERTER_FLUSHED = 3
Flushing is finished


## `CredentialsType`

Enumeration describing different kinds of native credential types.

C - `GCredentialsType`

### CREDENTIALS_TYPE_INVALID = 0
Indicates an invalid native credential type.

### CREDENTIALS_TYPE_LINUX_UCRED = 1
The native credentials type is a struct ucred.

### CREDENTIALS_TYPE_FREEBSD_CMSGCRED = 2
The native credentials type is a struct cmsgcred.

### CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED = 3
The native credentials type is a struct sockpeercred. Added in 2.30.

### CREDENTIALS_TYPE_SOLARIS_UCRED = 4
The native credentials type is a ucred_t. Added in 2.40.

### CREDENTIALS_TYPE_NETBSD_UNPCBID = 5
The native credentials type is a struct unpcbid.


## `DBusError`

Error codes for the %G_DBUS_ERROR error domain.

C - `GDBusError`

### DBUS_ERROR_FAILED = 0
A generic error; "something went wrong" - see the error message for
more.

### DBUS_ERROR_NO_MEMORY = 1
There was not enough memory to complete an operation.

### DBUS_ERROR_SERVICE_UNKNOWN = 2
The bus doesn't know how to launch a service to supply the bus name
you wanted.

### DBUS_ERROR_NAME_HAS_NO_OWNER = 3
The bus name you referenced doesn't exist (i.e. no application owns
it).

### DBUS_ERROR_NO_REPLY = 4
No reply to a message expecting one, usually means a timeout occurred.

### DBUS_ERROR_IO_ERROR = 5
Something went wrong reading or writing to a socket, for example.

### DBUS_ERROR_BAD_ADDRESS = 6
A D-Bus bus address was malformed.

### DBUS_ERROR_NOT_SUPPORTED = 7
Requested operation isn't supported (like ENOSYS on UNIX).

### DBUS_ERROR_LIMITS_EXCEEDED = 8
Some limited resource is exhausted.

### DBUS_ERROR_ACCESS_DENIED = 9
Security restrictions don't allow doing what you're trying to do.

### DBUS_ERROR_AUTH_FAILED = 10
Authentication didn't work.

### DBUS_ERROR_NO_SERVER = 11
Unable to connect to server (probably caused by ECONNREFUSED on a
socket).

### DBUS_ERROR_TIMEOUT = 12
Certain timeout errors, possibly ETIMEDOUT on a socket.  Note that
%G_DBUS_ERROR_NO_REPLY is used for message reply timeouts. Warning:
this is confusingly-named given that %G_DBUS_ERROR_TIMED_OUT also
exists. We can't fix it for compatibility reasons so just be
careful.

### DBUS_ERROR_NO_NETWORK = 13
No network access (probably ENETUNREACH on a socket).

### DBUS_ERROR_ADDRESS_IN_USE = 14
Can't bind a socket since its address is in use (i.e. EADDRINUSE).

### DBUS_ERROR_DISCONNECTED = 15
The connection is disconnected and you're trying to use it.

### DBUS_ERROR_INVALID_ARGS = 16
Invalid arguments passed to a method call.

### DBUS_ERROR_FILE_NOT_FOUND = 17
Missing file.

### DBUS_ERROR_FILE_EXISTS = 18
Existing file and the operation you're using does not silently overwrite.

### DBUS_ERROR_UNKNOWN_METHOD = 19
Method name you invoked isn't known by the object you invoked it on.

### DBUS_ERROR_TIMED_OUT = 20
Certain timeout errors, e.g. while starting a service. Warning: this is
confusingly-named given that %G_DBUS_ERROR_TIMEOUT also exists. We
can't fix it for compatibility reasons so just be careful.

### DBUS_ERROR_MATCH_RULE_NOT_FOUND = 21
Tried to remove or modify a match rule that didn't exist.

### DBUS_ERROR_MATCH_RULE_INVALID = 22
The match rule isn't syntactically valid.

### DBUS_ERROR_SPAWN_EXEC_FAILED = 23
While starting a new process, the exec() call failed.

### DBUS_ERROR_SPAWN_FORK_FAILED = 24
While starting a new process, the fork() call failed.

### DBUS_ERROR_SPAWN_CHILD_EXITED = 25
While starting a new process, the child exited with a status code.

### DBUS_ERROR_SPAWN_CHILD_SIGNALED = 26
While starting a new process, the child exited on a signal.

### DBUS_ERROR_SPAWN_FAILED = 27
While starting a new process, something went wrong.

### DBUS_ERROR_SPAWN_SETUP_FAILED = 28
We failed to setup the environment correctly.

### DBUS_ERROR_SPAWN_CONFIG_INVALID = 29
We failed to setup the config parser correctly.

### DBUS_ERROR_SPAWN_SERVICE_INVALID = 30
Bus name was not valid.

### DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND = 31
Service file not found in system-services directory.

### DBUS_ERROR_SPAWN_PERMISSIONS_INVALID = 32
Permissions are incorrect on the setuid helper.

### DBUS_ERROR_SPAWN_FILE_INVALID = 33
Service file invalid (Name, User or Exec missing).

### DBUS_ERROR_SPAWN_NO_MEMORY = 34
Tried to get a UNIX process ID and it wasn't available.

### DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN = 35
Tried to get a UNIX process ID and it wasn't available.

### DBUS_ERROR_INVALID_SIGNATURE = 36
A type signature is not valid.

### DBUS_ERROR_INVALID_FILE_CONTENT = 37
A file contains invalid syntax or is otherwise broken.

### DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN = 38
Asked for SELinux security context and it wasn't available.

### DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN = 39
Asked for ADT audit data and it wasn't available.

### DBUS_ERROR_OBJECT_PATH_IN_USE = 40
There's already an object with the requested object path.

### DBUS_ERROR_UNKNOWN_OBJECT = 41
Object you invoked a method on isn't known. Since 2.42

### DBUS_ERROR_UNKNOWN_INTERFACE = 42
Interface you invoked a method on isn't known by the object. Since 2.42

### DBUS_ERROR_UNKNOWN_PROPERTY = 43
Property you tried to access isn't known by the object. Since 2.42

### DBUS_ERROR_PROPERTY_READ_ONLY = 44
Property you tried to set is read-only. Since 2.42


## `DBusMessageByteOrder`

Enumeration used to describe the byte order of a D-Bus message.

C - `GDBusMessageByteOrder`

### DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN = 66
The byte order is big endian.

### DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN = 108
The byte order is little endian.


## `DBusMessageHeaderField`

Header fields used in #GDBusMessage.

C - `GDBusMessageHeaderField`

### DBUS_MESSAGE_HEADER_FIELD_INVALID = 0
Not a valid header field.

### DBUS_MESSAGE_HEADER_FIELD_PATH = 1
The object path.

### DBUS_MESSAGE_HEADER_FIELD_INTERFACE = 2
The interface name.

### DBUS_MESSAGE_HEADER_FIELD_MEMBER = 3
The method or signal name.

### DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME = 4
The name of the error that occurred.

### DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL = 5
The serial number the message is a reply to.

### DBUS_MESSAGE_HEADER_FIELD_DESTINATION = 6
The name the message is intended for.

### DBUS_MESSAGE_HEADER_FIELD_SENDER = 7
Unique name of the sender of the message (filled in by the bus).

### DBUS_MESSAGE_HEADER_FIELD_SIGNATURE = 8
The signature of the message body.

### DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS = 9
The number of UNIX file descriptors that accompany the message.


## `DBusMessageType`

Message types used in #GDBusMessage.

C - `GDBusMessageType`

### DBUS_MESSAGE_TYPE_INVALID = 0
Message is of invalid type.

### DBUS_MESSAGE_TYPE_METHOD_CALL = 1
Method call.

### DBUS_MESSAGE_TYPE_METHOD_RETURN = 2
Method reply.

### DBUS_MESSAGE_TYPE_ERROR = 3
Error reply.

### DBUS_MESSAGE_TYPE_SIGNAL = 4
Signal emission.


## `DataStreamByteOrder`

#GDataStreamByteOrder is used to ensure proper endianness of streaming data sources
across various machine architectures.

C - `GDataStreamByteOrder`

### DATA_STREAM_BYTE_ORDER_BIG_ENDIAN = 0
Selects Big Endian byte order.

### DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN = 1
Selects Little Endian byte order.

### DATA_STREAM_BYTE_ORDER_HOST_ENDIAN = 2
Selects endianness based on host machine's architecture.


## `DataStreamNewlineType`

#GDataStreamNewlineType is used when checking for or setting the line endings for a given file.

C - `GDataStreamNewlineType`

### DATA_STREAM_NEWLINE_TYPE_LF = 0
Selects "LF" line endings, common on most modern UNIX platforms.

### DATA_STREAM_NEWLINE_TYPE_CR = 1
Selects "CR" line endings.

### DATA_STREAM_NEWLINE_TYPE_CR_LF = 2
Selects "CR, LF" line ending, common on Microsoft Windows.

### DATA_STREAM_NEWLINE_TYPE_ANY = 3
Automatically try to handle any line ending type.


## `DriveStartStopType`

Enumeration describing how a drive can be started/stopped.

C - `GDriveStartStopType`

### DRIVE_START_STOP_TYPE_UNKNOWN = 0
Unknown or drive doesn't support
   start/stop.

### DRIVE_START_STOP_TYPE_SHUTDOWN = 1
The stop method will physically
   shut down the drive and e.g. power down the port the drive is
   attached to.

### DRIVE_START_STOP_TYPE_NETWORK = 2
The start/stop methods are used
   for connecting/disconnect to the drive over the network.

### DRIVE_START_STOP_TYPE_MULTIDISK = 3
The start/stop methods will
   assemble/disassemble a virtual drive from several physical
   drives.

### DRIVE_START_STOP_TYPE_PASSWORD = 4
The start/stop methods will
   unlock/lock the disk (for example using the ATA <quote>SECURITY
   UNLOCK DEVICE</quote> command)


## `EmblemOrigin`

GEmblemOrigin is used to add information about the origin of the emblem
to #GEmblem.

C - `GEmblemOrigin`

### EMBLEM_ORIGIN_UNKNOWN = 0
Emblem of unknown origin

### EMBLEM_ORIGIN_DEVICE = 1
Emblem adds device-specific information

### EMBLEM_ORIGIN_LIVEMETADATA = 2
Emblem depicts live metadata, such as "readonly"

### EMBLEM_ORIGIN_TAG = 3
Emblem comes from a user-defined tag, e.g. set by nautilus (in the future)


## `FileAttributeStatus`

Used by g_file_set_attributes_from_info() when setting file attributes.

C - `GFileAttributeStatus`

### FILE_ATTRIBUTE_STATUS_UNSET = 0
Attribute value is unset (empty).

### FILE_ATTRIBUTE_STATUS_SET = 1
Attribute value is set.

### FILE_ATTRIBUTE_STATUS_ERROR_SETTING = 2
Indicates an error in setting the value.


## `FileAttributeType`

The data types for file attributes.

C - `GFileAttributeType`

### FILE_ATTRIBUTE_TYPE_INVALID = 0
indicates an invalid or uninitalized type.

### FILE_ATTRIBUTE_TYPE_STRING = 1
a null terminated UTF8 string.

### FILE_ATTRIBUTE_TYPE_BYTE_STRING = 2
a zero terminated string of non-zero bytes.

### FILE_ATTRIBUTE_TYPE_BOOLEAN = 3
a boolean value.

### FILE_ATTRIBUTE_TYPE_UINT32 = 4
an unsigned 4-byte/32-bit integer.

### FILE_ATTRIBUTE_TYPE_INT32 = 5
a signed 4-byte/32-bit integer.

### FILE_ATTRIBUTE_TYPE_UINT64 = 6
an unsigned 8-byte/64-bit integer.

### FILE_ATTRIBUTE_TYPE_INT64 = 7
a signed 8-byte/64-bit integer.

### FILE_ATTRIBUTE_TYPE_OBJECT = 8
a #GObject.

### FILE_ATTRIBUTE_TYPE_STRINGV = 9
a %NULL terminated char **. Since 2.22


## `FileMonitorEvent`

Specifies what type of event a monitor event is.

C - `GFileMonitorEvent`

### FILE_MONITOR_EVENT_CHANGED = 0
a file changed.

### FILE_MONITOR_EVENT_CHANGES_DONE_HINT = 1
a hint that this was probably the last change in a set of changes.

### FILE_MONITOR_EVENT_DELETED = 2
a file was deleted.

### FILE_MONITOR_EVENT_CREATED = 3
a file was created.

### FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED = 4
a file attribute was changed.

### FILE_MONITOR_EVENT_PRE_UNMOUNT = 5
the file location will soon be unmounted.

### FILE_MONITOR_EVENT_UNMOUNTED = 6
the file location was unmounted.

### FILE_MONITOR_EVENT_MOVED = 7
the file was moved -- only sent if the
  (deprecated) %G_FILE_MONITOR_SEND_MOVED flag is set

### FILE_MONITOR_EVENT_RENAMED = 8
the file was renamed within the
  current directory -- only sent if the %G_FILE_MONITOR_WATCH_MOVES
  flag is set.  Since: 2.46.

### FILE_MONITOR_EVENT_MOVED_IN = 9
the file was moved into the
  monitored directory from another location -- only sent if the
  %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46.

### FILE_MONITOR_EVENT_MOVED_OUT = 10
the file was moved out of the
  monitored directory to another location -- only sent if the
  %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46


## `FileType`

Indicates the file's on-disk type.

C - `GFileType`

### FILE_TYPE_UNKNOWN = 0
File's type is unknown.

### FILE_TYPE_REGULAR = 1
File handle represents a regular file.

### FILE_TYPE_DIRECTORY = 2
File handle represents a directory.

### FILE_TYPE_SYMBOLIC_LINK = 3
File handle represents a symbolic link
   (Unix systems).

### FILE_TYPE_SPECIAL = 4
File is a "special" file, such as a socket, fifo,
   block device, or character device.

### FILE_TYPE_SHORTCUT = 5
File is a shortcut (Windows systems).

### FILE_TYPE_MOUNTABLE = 6
File is a mountable location.


## `FilesystemPreviewType`

Indicates a hint from the file system whether files should be
previewed in a file manager. Returned as the value of the key
&num;G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.

C - `GFilesystemPreviewType`

### FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS = 0
Only preview files if user has explicitly requested it.

### FILESYSTEM_PREVIEW_TYPE_IF_LOCAL = 1
Preview files if user has requested preview of "local" files.

### FILESYSTEM_PREVIEW_TYPE_NEVER = 2
Never preview files.


## `IOErrorEnum`

Error codes returned by GIO functions.

Note that this domain may be extended in future GLib releases. In
general, new error codes either only apply to new APIs, or else
replace %G_IO_ERROR_FAILED in cases that were not explicitly
distinguished before. You should therefore avoid writing code like
|[<!-- language="C" -->
if (g_error_matches (error, G_IO_ERROR, G_IO_ERROR_FAILED))
  {
    // Assume that this is EPRINTERONFIRE
    ...
  }
]|
but should instead treat all unrecognized error codes the same as
&num;G_IO_ERROR_FAILED.

C - `GIOErrorEnum`

### IO_ERROR_FAILED = 0
Generic error condition for when an operation fails
    and no more specific #GIOErrorEnum value is defined.

### IO_ERROR_NOT_FOUND = 1
File not found.

### IO_ERROR_EXISTS = 2
File already exists.

### IO_ERROR_IS_DIRECTORY = 3
File is a directory.

### IO_ERROR_NOT_DIRECTORY = 4
File is not a directory.

### IO_ERROR_NOT_EMPTY = 5
File is a directory that isn't empty.

### IO_ERROR_NOT_REGULAR_FILE = 6
File is not a regular file.

### IO_ERROR_NOT_SYMBOLIC_LINK = 7
File is not a symbolic link.

### IO_ERROR_NOT_MOUNTABLE_FILE = 8
File cannot be mounted.

### IO_ERROR_FILENAME_TOO_LONG = 9
Filename is too many characters.

### IO_ERROR_INVALID_FILENAME = 10
Filename is invalid or contains invalid characters.

### IO_ERROR_TOO_MANY_LINKS = 11
File contains too many symbolic links.

### IO_ERROR_NO_SPACE = 12
No space left on drive.

### IO_ERROR_INVALID_ARGUMENT = 13
Invalid argument.

### IO_ERROR_PERMISSION_DENIED = 14
Permission denied.

### IO_ERROR_NOT_SUPPORTED = 15
Operation (or one of its parameters) not supported

### IO_ERROR_NOT_MOUNTED = 16
File isn't mounted.

### IO_ERROR_ALREADY_MOUNTED = 17
File is already mounted.

### IO_ERROR_CLOSED = 18
File was closed.

### IO_ERROR_CANCELLED = 19
Operation was cancelled. See #GCancellable.

### IO_ERROR_PENDING = 20
Operations are still pending.

### IO_ERROR_READ_ONLY = 21
File is read only.

### IO_ERROR_CANT_CREATE_BACKUP = 22
Backup couldn't be created.

### IO_ERROR_WRONG_ETAG = 23
File's Entity Tag was incorrect.

### IO_ERROR_TIMED_OUT = 24
Operation timed out.

### IO_ERROR_WOULD_RECURSE = 25
Operation would be recursive.

### IO_ERROR_BUSY = 26
File is busy.

### IO_ERROR_WOULD_BLOCK = 27
Operation would block.

### IO_ERROR_HOST_NOT_FOUND = 28
Host couldn't be found (remote operations).

### IO_ERROR_WOULD_MERGE = 29
Operation would merge files.

### IO_ERROR_FAILED_HANDLED = 30
Operation failed and a helper program has
    already interacted with the user. Do not display any error dialog.

### IO_ERROR_TOO_MANY_OPEN_FILES = 31
The current process has too many files
    open and can't open any more. Duplicate descriptors do count toward
    this limit. Since 2.20

### IO_ERROR_NOT_INITIALIZED = 32
The object has not been initialized. Since 2.22

### IO_ERROR_ADDRESS_IN_USE = 33
The requested address is already in use. Since 2.22

### IO_ERROR_PARTIAL_INPUT = 34
Need more input to finish operation. Since 2.24

### IO_ERROR_INVALID_DATA = 35
The input data was invalid. Since 2.24

### IO_ERROR_DBUS_ERROR = 36
A remote object generated an error that
    doesn't correspond to a locally registered #GError error
    domain. Use g_dbus_error_get_remote_error() to extract the D-Bus
    error name and g_dbus_error_strip_remote_error() to fix up the
    message so it matches what was received on the wire. Since 2.26.

### IO_ERROR_HOST_UNREACHABLE = 37
Host unreachable. Since 2.26

### IO_ERROR_NETWORK_UNREACHABLE = 38
Network unreachable. Since 2.26

### IO_ERROR_CONNECTION_REFUSED = 39
Connection refused. Since 2.26

### IO_ERROR_PROXY_FAILED = 40
Connection to proxy server failed. Since 2.26

### IO_ERROR_PROXY_AUTH_FAILED = 41
Proxy authentication failed. Since 2.26

### IO_ERROR_PROXY_NEED_AUTH = 42
Proxy server needs authentication. Since 2.26

### IO_ERROR_PROXY_NOT_ALLOWED = 43
Proxy connection is not allowed by ruleset.
    Since 2.26

### IO_ERROR_BROKEN_PIPE = 44
Broken pipe. Since 2.36

### IO_ERROR_CONNECTION_CLOSED = 44
Connection closed by peer. Note that this
    is the same code as %G_IO_ERROR_BROKEN_PIPE; before 2.44 some
    "connection closed" errors returned %G_IO_ERROR_BROKEN_PIPE, but others
    returned %G_IO_ERROR_FAILED. Now they should all return the same
    value, which has this more logical name. Since 2.44.

### IO_ERROR_NOT_CONNECTED = 45
Transport endpoint is not connected. Since 2.44

### IO_ERROR_MESSAGE_TOO_LARGE = 46
Message too large. Since 2.48.


## `IOModuleScopeFlags`

Flags for use with g_io_module_scope_new().

C - `GIOModuleScopeFlags`

### IO_MODULE_SCOPE_NONE = 0
No module scan flags

### IO_MODULE_SCOPE_BLOCK_DUPLICATES = 1
When using this scope to load or
    scan modules, automatically block a modules which has the same base
    basename as previously loaded module.


## `MountOperationResult`

#GMountOperationResult is returned as a result when a request for
information is send by the mounting operation.

C - `GMountOperationResult`

### MOUNT_OPERATION_HANDLED = 0
The request was fulfilled and the
    user specified data is now available

### MOUNT_OPERATION_ABORTED = 1
The user requested the mount operation
    to be aborted

### MOUNT_OPERATION_UNHANDLED = 2
The request was unhandled (i.e. not
    implemented)


## `NetworkConnectivity`

The host's network connectivity state, as reported by #GNetworkMonitor.

C - `GNetworkConnectivity`

### NETWORK_CONNECTIVITY_LOCAL = 1
The host is not configured with a
  route to the Internet; it may or may not be connected to a local
  network.

### NETWORK_CONNECTIVITY_LIMITED = 2
The host is connected to a network, but
  does not appear to be able to reach the full Internet, perhaps
  due to upstream network problems.

### NETWORK_CONNECTIVITY_PORTAL = 3
The host is behind a captive portal and
  cannot reach the full Internet.

### NETWORK_CONNECTIVITY_FULL = 4
The host is connected to a network, and
  appears to be able to reach the full Internet.


## `NotificationPriority`

Priority levels for #GNotifications.

C - `GNotificationPriority`

### NOTIFICATION_PRIORITY_NORMAL = 0
the default priority, to be used for the
  majority of notifications (for example email messages, software updates,
  completed download/sync operations)

### NOTIFICATION_PRIORITY_LOW = 1
for notifications that do not require
  immediate attention - typically used for contextual background
  information, such as contact birthdays or local weather

### NOTIFICATION_PRIORITY_HIGH = 2
for events that require more attention,
  usually because responses are time-sensitive (for example chat and SMS
  messages or alarms)

### NOTIFICATION_PRIORITY_URGENT = 3
for urgent notifications, or notifications
  that require a response in a short space of time (for example phone calls
  or emergency warnings)


## `PasswordSave`

#GPasswordSave is used to indicate the lifespan of a saved password.

&num;Gvfs stores passwords in the Gnome keyring when this flag allows it
to, and later retrieves it again from there.

C - `GPasswordSave`

### PASSWORD_SAVE_NEVER = 0
never save a password.

### PASSWORD_SAVE_FOR_SESSION = 1
save a password for the session.

### PASSWORD_SAVE_PERMANENTLY = 2
save a password permanently.


## `ResolverError`

An error code used with %G_RESOLVER_ERROR in a #GError returned
from a #GResolver routine.

C - `GResolverError`

### RESOLVER_ERROR_NOT_FOUND = 0
the requested name/address/service was not
    found

### RESOLVER_ERROR_TEMPORARY_FAILURE = 1
the requested information could not
    be looked up due to a network error or similar problem

### RESOLVER_ERROR_INTERNAL = 2
unknown error


## `ResolverRecordType`

The type of record that g_resolver_lookup_records() or
g_resolver_lookup_records_async() should retrieve. The records are returned
as lists of #GVariant tuples. Each record type has different values in
the variant tuples returned.

%G_RESOLVER_RECORD_SRV records are returned as variants with the signature
'(qqqs)', containing a guint16 with the priority, a guint16 with the
weight, a guint16 with the port, and a string of the hostname.

%G_RESOLVER_RECORD_MX records are returned as variants with the signature
'(qs)', representing a guint16 with the preference, and a string containing
the mail exchanger hostname.

%G_RESOLVER_RECORD_TXT records are returned as variants with the signature
'(as)', representing an array of the strings in the text record.

%G_RESOLVER_RECORD_SOA records are returned as variants with the signature
'(ssuuuuu)', representing a string containing the primary name server, a
string containing the administrator, the serial as a guint32, the refresh
interval as guint32, the retry interval as a guint32, the expire timeout
as a guint32, and the ttl as a guint32.

%G_RESOLVER_RECORD_NS records are returned as variants with the signature
'(s)', representing a string of the hostname of the name server.

C - `GResolverRecordType`

### RESOLVER_RECORD_SRV = 1
lookup DNS SRV records for a domain

### RESOLVER_RECORD_MX = 2
lookup DNS MX records for a domain

### RESOLVER_RECORD_TXT = 3
lookup DNS TXT records for a name

### RESOLVER_RECORD_SOA = 4
lookup DNS SOA records for a zone

### RESOLVER_RECORD_NS = 5
lookup DNS NS records for a domain


## `ResourceError`

An error code used with %G_RESOURCE_ERROR in a #GError returned
from a #GResource routine.

C - `GResourceError`

### RESOURCE_ERROR_NOT_FOUND = 0
no file was found at the requested path

### RESOURCE_ERROR_INTERNAL = 1
unknown error


## `SocketClientEvent`

Describes an event occurring on a #GSocketClient. See the
&num;GSocketClient::event signal for more details.

Additional values may be added to this type in the future.

C - `GSocketClientEvent`

### SOCKET_CLIENT_RESOLVING = 0
The client is doing a DNS lookup.

### SOCKET_CLIENT_RESOLVED = 1
The client has completed a DNS lookup.

### SOCKET_CLIENT_CONNECTING = 2
The client is connecting to a remote
  host (either a proxy or the destination server).

### SOCKET_CLIENT_CONNECTED = 3
The client has connected to a remote
  host.

### SOCKET_CLIENT_PROXY_NEGOTIATING = 4
The client is negotiating
  with a proxy to connect to the destination server.

### SOCKET_CLIENT_PROXY_NEGOTIATED = 5
The client has negotiated
  with the proxy server.

### SOCKET_CLIENT_TLS_HANDSHAKING = 6
The client is performing a
  TLS handshake.

### SOCKET_CLIENT_TLS_HANDSHAKED = 7
The client has performed a
  TLS handshake.

### SOCKET_CLIENT_COMPLETE = 8
The client is done with a particular
  #GSocketConnectable.


## `SocketFamily`

The protocol family of a #GSocketAddress. (These values are
identical to the system defines %AF_INET, %AF_INET6 and %AF_UNIX,
if available.)

C - `GSocketFamily`

### SOCKET_FAMILY_INVALID = 0
no address family

### SOCKET_FAMILY_UNIX = 1
the UNIX domain family

### SOCKET_FAMILY_IPV4 = 2
the IPv4 family

### SOCKET_FAMILY_IPV6 = 10
the IPv6 family


## `SocketListenerEvent`

Describes an event occurring on a #GSocketListener. See the
&num;GSocketListener::event signal for more details.

Additional values may be added to this type in the future.

C - `GSocketListenerEvent`

### SOCKET_LISTENER_BINDING = 0
The listener is about to bind a socket.

### SOCKET_LISTENER_BOUND = 1
The listener has bound a socket.

### SOCKET_LISTENER_LISTENING = 2
The listener is about to start
   listening on this socket.

### SOCKET_LISTENER_LISTENED = 3
The listener is now listening on
  this socket.


## `SocketProtocol`

A protocol identifier is specified when creating a #GSocket, which is a
family/type specific identifier, where 0 means the default protocol for
the particular family/type.

This enum contains a set of commonly available and used protocols. You
can also pass any other identifiers handled by the platform in order to
use protocols not listed here.

C - `GSocketProtocol`

### SOCKET_PROTOCOL_UNKNOWN = -1
The protocol type is unknown

### SOCKET_PROTOCOL_DEFAULT = 0
The default protocol for the family/type

### SOCKET_PROTOCOL_TCP = 6
TCP over IP

### SOCKET_PROTOCOL_UDP = 17
UDP over IP

### SOCKET_PROTOCOL_SCTP = 132
SCTP over IP


## `SocketType`

Flags used when creating a #GSocket. Some protocols may not implement
all the socket types.

C - `GSocketType`

### SOCKET_TYPE_INVALID = 0
Type unknown or wrong

### SOCKET_TYPE_STREAM = 1
Reliable connection-based byte streams (e.g. TCP).

### SOCKET_TYPE_DATAGRAM = 2
Connectionless, unreliable datagram passing.
    (e.g. UDP)

### SOCKET_TYPE_SEQPACKET = 3
Reliable connection-based passing of datagrams
    of fixed maximum length (e.g. SCTP).


## `TlsAuthenticationMode`

The client authentication mode for a #GTlsServerConnection.

C - `GTlsAuthenticationMode`

### TLS_AUTHENTICATION_NONE = 0
client authentication not required

### TLS_AUTHENTICATION_REQUESTED = 1
client authentication is requested

### TLS_AUTHENTICATION_REQUIRED = 2
client authentication is required


## `TlsCertificateRequestFlags`

Flags for g_tls_interaction_request_certificate(),
g_tls_interaction_request_certificate_async(), and
g_tls_interaction_invoke_request_certificate().

C - `GTlsCertificateRequestFlags`

### TLS_CERTIFICATE_REQUEST_NONE = 0
No flags


## `TlsDatabaseLookupFlags`

Flags for g_tls_database_lookup_certificate_for_handle(),
g_tls_database_lookup_certificate_issuer(),
and g_tls_database_lookup_certificates_issued_by().

C - `GTlsDatabaseLookupFlags`

### TLS_DATABASE_LOOKUP_NONE = 0
No lookup flags

### TLS_DATABASE_LOOKUP_KEYPAIR = 1
Restrict lookup to certificates that have
    a private key.


## `TlsError`

An error code used with %G_TLS_ERROR in a #GError returned from a
TLS-related routine.

C - `GTlsError`

### TLS_ERROR_UNAVAILABLE = 0
No TLS provider is available

### TLS_ERROR_MISC = 1
Miscellaneous TLS error

### TLS_ERROR_BAD_CERTIFICATE = 2
A certificate could not be parsed

### TLS_ERROR_NOT_TLS = 3
The TLS handshake failed because the
  peer does not seem to be a TLS server.

### TLS_ERROR_HANDSHAKE = 4
The TLS handshake failed because the
  peer's certificate was not acceptable.

### TLS_ERROR_CERTIFICATE_REQUIRED = 5
The TLS handshake failed because
  the server requested a client-side certificate, but none was
  provided. See g_tls_connection_set_certificate().

### TLS_ERROR_EOF = 6
The TLS connection was closed without proper
  notice, which may indicate an attack. See
  g_tls_connection_set_require_close_notify().


## `TlsInteractionResult`

#GTlsInteractionResult is returned by various functions in #GTlsInteraction
when finishing an interaction request.

C - `GTlsInteractionResult`

### TLS_INTERACTION_UNHANDLED = 0
The interaction was unhandled (i.e. not
    implemented).

### TLS_INTERACTION_HANDLED = 1
The interaction completed, and resulting data
    is available.

### TLS_INTERACTION_FAILED = 2
The interaction has failed, or was cancelled.
    and the operation should be aborted.


## `TlsRehandshakeMode`

When to allow rehandshaking. See
g_tls_connection_set_rehandshake_mode().

C - `GTlsRehandshakeMode`

### TLS_REHANDSHAKE_NEVER = 0
Never allow rehandshaking

### TLS_REHANDSHAKE_SAFELY = 1
Allow safe rehandshaking only

### TLS_REHANDSHAKE_UNSAFELY = 2
Allow unsafe rehandshaking


## `UnixSocketAddressType`

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

C - `GUnixSocketAddressType`

### UNIX_SOCKET_ADDRESS_INVALID = 0
invalid

### UNIX_SOCKET_ADDRESS_ANONYMOUS = 1
anonymous

### UNIX_SOCKET_ADDRESS_PATH = 2
a filesystem path

### UNIX_SOCKET_ADDRESS_ABSTRACT = 3
an abstract name

### UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED = 4
an abstract name, 0-padded
  to the full length of a unix socket name


## `ZlibCompressorFormat`

Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.

C - `GZlibCompressorFormat`

### ZLIB_COMPRESSOR_FORMAT_ZLIB = 0
deflate compression with zlib header

### ZLIB_COMPRESSOR_FORMAT_GZIP = 1
gzip file format

### ZLIB_COMPRESSOR_FORMAT_RAW = 2
deflate compression with no header


