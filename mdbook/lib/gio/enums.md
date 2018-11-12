# `gio` enums

## `BusType`

An enumeration for well-known message buses.

- BUS_TYPE_STARTER = -1
- BUS_TYPE_NONE = 0
- BUS_TYPE_SYSTEM = 1
- BUS_TYPE_SESSION = 2

C - `GBusType`

## `ConverterResult`

Results returned from g_converter_convert().

- CONVERTER_ERROR = 0
- CONVERTER_CONVERTED = 1
- CONVERTER_FINISHED = 2
- CONVERTER_FLUSHED = 3

C - `GConverterResult`

## `CredentialsType`

Enumeration describing different kinds of native credential types.

- CREDENTIALS_TYPE_INVALID = 0
- CREDENTIALS_TYPE_LINUX_UCRED = 1
- CREDENTIALS_TYPE_FREEBSD_CMSGCRED = 2
- CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED = 3
- CREDENTIALS_TYPE_SOLARIS_UCRED = 4
- CREDENTIALS_TYPE_NETBSD_UNPCBID = 5

C - `GCredentialsType`

## `DBusError`

Error codes for the %G_DBUS_ERROR error domain.

- DBUS_ERROR_FAILED = 0
- DBUS_ERROR_NO_MEMORY = 1
- DBUS_ERROR_SERVICE_UNKNOWN = 2
- DBUS_ERROR_NAME_HAS_NO_OWNER = 3
- DBUS_ERROR_NO_REPLY = 4
- DBUS_ERROR_IO_ERROR = 5
- DBUS_ERROR_BAD_ADDRESS = 6
- DBUS_ERROR_NOT_SUPPORTED = 7
- DBUS_ERROR_LIMITS_EXCEEDED = 8
- DBUS_ERROR_ACCESS_DENIED = 9
- DBUS_ERROR_AUTH_FAILED = 10
- DBUS_ERROR_NO_SERVER = 11
- DBUS_ERROR_TIMEOUT = 12
- DBUS_ERROR_NO_NETWORK = 13
- DBUS_ERROR_ADDRESS_IN_USE = 14
- DBUS_ERROR_DISCONNECTED = 15
- DBUS_ERROR_INVALID_ARGS = 16
- DBUS_ERROR_FILE_NOT_FOUND = 17
- DBUS_ERROR_FILE_EXISTS = 18
- DBUS_ERROR_UNKNOWN_METHOD = 19
- DBUS_ERROR_TIMED_OUT = 20
- DBUS_ERROR_MATCH_RULE_NOT_FOUND = 21
- DBUS_ERROR_MATCH_RULE_INVALID = 22
- DBUS_ERROR_SPAWN_EXEC_FAILED = 23
- DBUS_ERROR_SPAWN_FORK_FAILED = 24
- DBUS_ERROR_SPAWN_CHILD_EXITED = 25
- DBUS_ERROR_SPAWN_CHILD_SIGNALED = 26
- DBUS_ERROR_SPAWN_FAILED = 27
- DBUS_ERROR_SPAWN_SETUP_FAILED = 28
- DBUS_ERROR_SPAWN_CONFIG_INVALID = 29
- DBUS_ERROR_SPAWN_SERVICE_INVALID = 30
- DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND = 31
- DBUS_ERROR_SPAWN_PERMISSIONS_INVALID = 32
- DBUS_ERROR_SPAWN_FILE_INVALID = 33
- DBUS_ERROR_SPAWN_NO_MEMORY = 34
- DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN = 35
- DBUS_ERROR_INVALID_SIGNATURE = 36
- DBUS_ERROR_INVALID_FILE_CONTENT = 37
- DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN = 38
- DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN = 39
- DBUS_ERROR_OBJECT_PATH_IN_USE = 40
- DBUS_ERROR_UNKNOWN_OBJECT = 41
- DBUS_ERROR_UNKNOWN_INTERFACE = 42
- DBUS_ERROR_UNKNOWN_PROPERTY = 43
- DBUS_ERROR_PROPERTY_READ_ONLY = 44

C - `GDBusError`

## `DBusMessageByteOrder`

Enumeration used to describe the byte order of a D-Bus message.

- DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN = 66
- DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN = 108

C - `GDBusMessageByteOrder`

## `DBusMessageHeaderField`

Header fields used in #GDBusMessage.

- DBUS_MESSAGE_HEADER_FIELD_INVALID = 0
- DBUS_MESSAGE_HEADER_FIELD_PATH = 1
- DBUS_MESSAGE_HEADER_FIELD_INTERFACE = 2
- DBUS_MESSAGE_HEADER_FIELD_MEMBER = 3
- DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME = 4
- DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL = 5
- DBUS_MESSAGE_HEADER_FIELD_DESTINATION = 6
- DBUS_MESSAGE_HEADER_FIELD_SENDER = 7
- DBUS_MESSAGE_HEADER_FIELD_SIGNATURE = 8
- DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS = 9

C - `GDBusMessageHeaderField`

## `DBusMessageType`

Message types used in #GDBusMessage.

- DBUS_MESSAGE_TYPE_INVALID = 0
- DBUS_MESSAGE_TYPE_METHOD_CALL = 1
- DBUS_MESSAGE_TYPE_METHOD_RETURN = 2
- DBUS_MESSAGE_TYPE_ERROR = 3
- DBUS_MESSAGE_TYPE_SIGNAL = 4

C - `GDBusMessageType`

## `DataStreamByteOrder`

#GDataStreamByteOrder is used to ensure proper endianness of streaming data sources
across various machine architectures.

- DATA_STREAM_BYTE_ORDER_BIG_ENDIAN = 0
- DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN = 1
- DATA_STREAM_BYTE_ORDER_HOST_ENDIAN = 2

C - `GDataStreamByteOrder`

## `DataStreamNewlineType`

#GDataStreamNewlineType is used when checking for or setting the line endings for a given file.

- DATA_STREAM_NEWLINE_TYPE_LF = 0
- DATA_STREAM_NEWLINE_TYPE_CR = 1
- DATA_STREAM_NEWLINE_TYPE_CR_LF = 2
- DATA_STREAM_NEWLINE_TYPE_ANY = 3

C - `GDataStreamNewlineType`

## `DriveStartStopType`

Enumeration describing how a drive can be started/stopped.

- DRIVE_START_STOP_TYPE_UNKNOWN = 0
- DRIVE_START_STOP_TYPE_SHUTDOWN = 1
- DRIVE_START_STOP_TYPE_NETWORK = 2
- DRIVE_START_STOP_TYPE_MULTIDISK = 3
- DRIVE_START_STOP_TYPE_PASSWORD = 4

C - `GDriveStartStopType`

## `EmblemOrigin`

GEmblemOrigin is used to add information about the origin of the emblem
to #GEmblem.

- EMBLEM_ORIGIN_UNKNOWN = 0
- EMBLEM_ORIGIN_DEVICE = 1
- EMBLEM_ORIGIN_LIVEMETADATA = 2
- EMBLEM_ORIGIN_TAG = 3

C - `GEmblemOrigin`

## `FileAttributeStatus`

Used by g_file_set_attributes_from_info() when setting file attributes.

- FILE_ATTRIBUTE_STATUS_UNSET = 0
- FILE_ATTRIBUTE_STATUS_SET = 1
- FILE_ATTRIBUTE_STATUS_ERROR_SETTING = 2

C - `GFileAttributeStatus`

## `FileAttributeType`

The data types for file attributes.

- FILE_ATTRIBUTE_TYPE_INVALID = 0
- FILE_ATTRIBUTE_TYPE_STRING = 1
- FILE_ATTRIBUTE_TYPE_BYTE_STRING = 2
- FILE_ATTRIBUTE_TYPE_BOOLEAN = 3
- FILE_ATTRIBUTE_TYPE_UINT32 = 4
- FILE_ATTRIBUTE_TYPE_INT32 = 5
- FILE_ATTRIBUTE_TYPE_UINT64 = 6
- FILE_ATTRIBUTE_TYPE_INT64 = 7
- FILE_ATTRIBUTE_TYPE_OBJECT = 8
- FILE_ATTRIBUTE_TYPE_STRINGV = 9

C - `GFileAttributeType`

## `FileMonitorEvent`

Specifies what type of event a monitor event is.

- FILE_MONITOR_EVENT_CHANGED = 0
- FILE_MONITOR_EVENT_CHANGES_DONE_HINT = 1
- FILE_MONITOR_EVENT_DELETED = 2
- FILE_MONITOR_EVENT_CREATED = 3
- FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED = 4
- FILE_MONITOR_EVENT_PRE_UNMOUNT = 5
- FILE_MONITOR_EVENT_UNMOUNTED = 6
- FILE_MONITOR_EVENT_MOVED = 7
- FILE_MONITOR_EVENT_RENAMED = 8
- FILE_MONITOR_EVENT_MOVED_IN = 9
- FILE_MONITOR_EVENT_MOVED_OUT = 10

C - `GFileMonitorEvent`

## `FileType`

Indicates the file's on-disk type.

- FILE_TYPE_UNKNOWN = 0
- FILE_TYPE_REGULAR = 1
- FILE_TYPE_DIRECTORY = 2
- FILE_TYPE_SYMBOLIC_LINK = 3
- FILE_TYPE_SPECIAL = 4
- FILE_TYPE_SHORTCUT = 5
- FILE_TYPE_MOUNTABLE = 6

C - `GFileType`

## `FilesystemPreviewType`

Indicates a hint from the file system whether files should be
previewed in a file manager. Returned as the value of the key
&num;G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.

- FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS = 0
- FILESYSTEM_PREVIEW_TYPE_IF_LOCAL = 1
- FILESYSTEM_PREVIEW_TYPE_NEVER = 2

C - `GFilesystemPreviewType`

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

- IO_ERROR_FAILED = 0
- IO_ERROR_NOT_FOUND = 1
- IO_ERROR_EXISTS = 2
- IO_ERROR_IS_DIRECTORY = 3
- IO_ERROR_NOT_DIRECTORY = 4
- IO_ERROR_NOT_EMPTY = 5
- IO_ERROR_NOT_REGULAR_FILE = 6
- IO_ERROR_NOT_SYMBOLIC_LINK = 7
- IO_ERROR_NOT_MOUNTABLE_FILE = 8
- IO_ERROR_FILENAME_TOO_LONG = 9
- IO_ERROR_INVALID_FILENAME = 10
- IO_ERROR_TOO_MANY_LINKS = 11
- IO_ERROR_NO_SPACE = 12
- IO_ERROR_INVALID_ARGUMENT = 13
- IO_ERROR_PERMISSION_DENIED = 14
- IO_ERROR_NOT_SUPPORTED = 15
- IO_ERROR_NOT_MOUNTED = 16
- IO_ERROR_ALREADY_MOUNTED = 17
- IO_ERROR_CLOSED = 18
- IO_ERROR_CANCELLED = 19
- IO_ERROR_PENDING = 20
- IO_ERROR_READ_ONLY = 21
- IO_ERROR_CANT_CREATE_BACKUP = 22
- IO_ERROR_WRONG_ETAG = 23
- IO_ERROR_TIMED_OUT = 24
- IO_ERROR_WOULD_RECURSE = 25
- IO_ERROR_BUSY = 26
- IO_ERROR_WOULD_BLOCK = 27
- IO_ERROR_HOST_NOT_FOUND = 28
- IO_ERROR_WOULD_MERGE = 29
- IO_ERROR_FAILED_HANDLED = 30
- IO_ERROR_TOO_MANY_OPEN_FILES = 31
- IO_ERROR_NOT_INITIALIZED = 32
- IO_ERROR_ADDRESS_IN_USE = 33
- IO_ERROR_PARTIAL_INPUT = 34
- IO_ERROR_INVALID_DATA = 35
- IO_ERROR_DBUS_ERROR = 36
- IO_ERROR_HOST_UNREACHABLE = 37
- IO_ERROR_NETWORK_UNREACHABLE = 38
- IO_ERROR_CONNECTION_REFUSED = 39
- IO_ERROR_PROXY_FAILED = 40
- IO_ERROR_PROXY_AUTH_FAILED = 41
- IO_ERROR_PROXY_NEED_AUTH = 42
- IO_ERROR_PROXY_NOT_ALLOWED = 43
- IO_ERROR_BROKEN_PIPE = 44
- IO_ERROR_CONNECTION_CLOSED = 44
- IO_ERROR_NOT_CONNECTED = 45
- IO_ERROR_MESSAGE_TOO_LARGE = 46

C - `GIOErrorEnum`

## `IOModuleScopeFlags`

Flags for use with g_io_module_scope_new().

- IO_MODULE_SCOPE_NONE = 0
- IO_MODULE_SCOPE_BLOCK_DUPLICATES = 1

C - `GIOModuleScopeFlags`

## `MountOperationResult`

#GMountOperationResult is returned as a result when a request for
information is send by the mounting operation.

- MOUNT_OPERATION_HANDLED = 0
- MOUNT_OPERATION_ABORTED = 1
- MOUNT_OPERATION_UNHANDLED = 2

C - `GMountOperationResult`

## `NetworkConnectivity`

The host's network connectivity state, as reported by #GNetworkMonitor.

- NETWORK_CONNECTIVITY_LOCAL = 1
- NETWORK_CONNECTIVITY_LIMITED = 2
- NETWORK_CONNECTIVITY_PORTAL = 3
- NETWORK_CONNECTIVITY_FULL = 4

C - `GNetworkConnectivity`

## `NotificationPriority`

Priority levels for #GNotifications.

- NOTIFICATION_PRIORITY_NORMAL = 0
- NOTIFICATION_PRIORITY_LOW = 1
- NOTIFICATION_PRIORITY_HIGH = 2
- NOTIFICATION_PRIORITY_URGENT = 3

C - `GNotificationPriority`

## `PasswordSave`

#GPasswordSave is used to indicate the lifespan of a saved password.

&num;Gvfs stores passwords in the Gnome keyring when this flag allows it
to, and later retrieves it again from there.

- PASSWORD_SAVE_NEVER = 0
- PASSWORD_SAVE_FOR_SESSION = 1
- PASSWORD_SAVE_PERMANENTLY = 2

C - `GPasswordSave`

## `ResolverError`

An error code used with %G_RESOLVER_ERROR in a #GError returned
from a #GResolver routine.

- RESOLVER_ERROR_NOT_FOUND = 0
- RESOLVER_ERROR_TEMPORARY_FAILURE = 1
- RESOLVER_ERROR_INTERNAL = 2

C - `GResolverError`

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

- RESOLVER_RECORD_SRV = 1
- RESOLVER_RECORD_MX = 2
- RESOLVER_RECORD_TXT = 3
- RESOLVER_RECORD_SOA = 4
- RESOLVER_RECORD_NS = 5

C - `GResolverRecordType`

## `ResourceError`

An error code used with %G_RESOURCE_ERROR in a #GError returned
from a #GResource routine.

- RESOURCE_ERROR_NOT_FOUND = 0
- RESOURCE_ERROR_INTERNAL = 1

C - `GResourceError`

## `SocketClientEvent`

Describes an event occurring on a #GSocketClient. See the
&num;GSocketClient::event signal for more details.

Additional values may be added to this type in the future.

- SOCKET_CLIENT_RESOLVING = 0
- SOCKET_CLIENT_RESOLVED = 1
- SOCKET_CLIENT_CONNECTING = 2
- SOCKET_CLIENT_CONNECTED = 3
- SOCKET_CLIENT_PROXY_NEGOTIATING = 4
- SOCKET_CLIENT_PROXY_NEGOTIATED = 5
- SOCKET_CLIENT_TLS_HANDSHAKING = 6
- SOCKET_CLIENT_TLS_HANDSHAKED = 7
- SOCKET_CLIENT_COMPLETE = 8

C - `GSocketClientEvent`

## `SocketFamily`

The protocol family of a #GSocketAddress. (These values are
identical to the system defines %AF_INET, %AF_INET6 and %AF_UNIX,
if available.)

- SOCKET_FAMILY_INVALID = 0
- SOCKET_FAMILY_UNIX = 1
- SOCKET_FAMILY_IPV4 = 2
- SOCKET_FAMILY_IPV6 = 10

C - `GSocketFamily`

## `SocketListenerEvent`

Describes an event occurring on a #GSocketListener. See the
&num;GSocketListener::event signal for more details.

Additional values may be added to this type in the future.

- SOCKET_LISTENER_BINDING = 0
- SOCKET_LISTENER_BOUND = 1
- SOCKET_LISTENER_LISTENING = 2
- SOCKET_LISTENER_LISTENED = 3

C - `GSocketListenerEvent`

## `SocketProtocol`

A protocol identifier is specified when creating a #GSocket, which is a
family/type specific identifier, where 0 means the default protocol for
the particular family/type.

This enum contains a set of commonly available and used protocols. You
can also pass any other identifiers handled by the platform in order to
use protocols not listed here.

- SOCKET_PROTOCOL_UNKNOWN = -1
- SOCKET_PROTOCOL_DEFAULT = 0
- SOCKET_PROTOCOL_TCP = 6
- SOCKET_PROTOCOL_UDP = 17
- SOCKET_PROTOCOL_SCTP = 132

C - `GSocketProtocol`

## `SocketType`

Flags used when creating a #GSocket. Some protocols may not implement
all the socket types.

- SOCKET_TYPE_INVALID = 0
- SOCKET_TYPE_STREAM = 1
- SOCKET_TYPE_DATAGRAM = 2
- SOCKET_TYPE_SEQPACKET = 3

C - `GSocketType`

## `TlsAuthenticationMode`

The client authentication mode for a #GTlsServerConnection.

- TLS_AUTHENTICATION_NONE = 0
- TLS_AUTHENTICATION_REQUESTED = 1
- TLS_AUTHENTICATION_REQUIRED = 2

C - `GTlsAuthenticationMode`

## `TlsCertificateRequestFlags`

Flags for g_tls_interaction_request_certificate(),
g_tls_interaction_request_certificate_async(), and
g_tls_interaction_invoke_request_certificate().

- TLS_CERTIFICATE_REQUEST_NONE = 0

C - `GTlsCertificateRequestFlags`

## `TlsDatabaseLookupFlags`

Flags for g_tls_database_lookup_certificate_for_handle(),
g_tls_database_lookup_certificate_issuer(),
and g_tls_database_lookup_certificates_issued_by().

- TLS_DATABASE_LOOKUP_NONE = 0
- TLS_DATABASE_LOOKUP_KEYPAIR = 1

C - `GTlsDatabaseLookupFlags`

## `TlsError`

An error code used with %G_TLS_ERROR in a #GError returned from a
TLS-related routine.

- TLS_ERROR_UNAVAILABLE = 0
- TLS_ERROR_MISC = 1
- TLS_ERROR_BAD_CERTIFICATE = 2
- TLS_ERROR_NOT_TLS = 3
- TLS_ERROR_HANDSHAKE = 4
- TLS_ERROR_CERTIFICATE_REQUIRED = 5
- TLS_ERROR_EOF = 6

C - `GTlsError`

## `TlsInteractionResult`

#GTlsInteractionResult is returned by various functions in #GTlsInteraction
when finishing an interaction request.

- TLS_INTERACTION_UNHANDLED = 0
- TLS_INTERACTION_HANDLED = 1
- TLS_INTERACTION_FAILED = 2

C - `GTlsInteractionResult`

## `TlsRehandshakeMode`

When to allow rehandshaking. See
g_tls_connection_set_rehandshake_mode().

- TLS_REHANDSHAKE_NEVER = 0
- TLS_REHANDSHAKE_SAFELY = 1
- TLS_REHANDSHAKE_UNSAFELY = 2

C - `GTlsRehandshakeMode`

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

- UNIX_SOCKET_ADDRESS_INVALID = 0
- UNIX_SOCKET_ADDRESS_ANONYMOUS = 1
- UNIX_SOCKET_ADDRESS_PATH = 2
- UNIX_SOCKET_ADDRESS_ABSTRACT = 3
- UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED = 4

C - `GUnixSocketAddressType`

## `ZlibCompressorFormat`

Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.

- ZLIB_COMPRESSOR_FORMAT_ZLIB = 0
- ZLIB_COMPRESSOR_FORMAT_GZIP = 1
- ZLIB_COMPRESSOR_FORMAT_RAW = 2

C - `GZlibCompressorFormat`

 = %!s(int=0)
- TLS_DATABASE_LOOKUP_KEYPAIR = %!s(int=1)
## `TlsError`

An error code used with %G_TLS_ERROR in a #GError returned from a
TLS-related routine.

C - `GTlsError`

- TLS_ERROR_UNAVAILABLE = %!s(int=0)
- TLS_ERROR_MISC = %!s(int=1)
- TLS_ERROR_BAD_CERTIFICATE = %!s(int=2)
- TLS_ERROR_NOT_TLS = %!s(int=3)
- TLS_ERROR_HANDSHAKE = %!s(int=4)
- TLS_ERROR_CERTIFICATE_REQUIRED = %!s(int=5)
- TLS_ERROR_EOF = %!s(int=6)
## `TlsInteractionResult`

#GTlsInteractionResult is returned by various functions in #GTlsInteraction
when finishing an interaction request.

C - `GTlsInteractionResult`

- TLS_INTERACTION_UNHANDLED = %!s(int=0)
- TLS_INTERACTION_HANDLED = %!s(int=1)
- TLS_INTERACTION_FAILED = %!s(int=2)
## `TlsRehandshakeMode`

When to allow rehandshaking. See
g_tls_connection_set_rehandshake_mode().

C - `GTlsRehandshakeMode`

- TLS_REHANDSHAKE_NEVER = %!s(int=0)
- TLS_REHANDSHAKE_SAFELY = %!s(int=1)
- TLS_REHANDSHAKE_UNSAFELY = %!s(int=2)
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

- UNIX_SOCKET_ADDRESS_INVALID = %!s(int=0)
- UNIX_SOCKET_ADDRESS_ANONYMOUS = %!s(int=1)
- UNIX_SOCKET_ADDRESS_PATH = %!s(int=2)
- UNIX_SOCKET_ADDRESS_ABSTRACT = %!s(int=3)
- UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED = %!s(int=4)
## `ZlibCompressorFormat`

Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.

C - `GZlibCompressorFormat`

- ZLIB_COMPRESSOR_FORMAT_ZLIB = %!s(int=0)
- ZLIB_COMPRESSOR_FORMAT_GZIP = %!s(int=1)
- ZLIB_COMPRESSOR_FORMAT_RAW = %!s(int=2)
