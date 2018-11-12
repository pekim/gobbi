# `gio` Enums

## `BusType`

An enumeration for well-known message buses.

C - `GBusType`

## `ConverterResult`

Results returned from g_converter_convert().

C - `GConverterResult`

## `CredentialsType`

Enumeration describing different kinds of native credential types.

C - `GCredentialsType`

## `DBusError`

Error codes for the %G_DBUS_ERROR error domain.

C - `GDBusError`

## `DBusMessageByteOrder`

Enumeration used to describe the byte order of a D-Bus message.

C - `GDBusMessageByteOrder`

## `DBusMessageHeaderField`

Header fields used in #GDBusMessage.

C - `GDBusMessageHeaderField`

## `DBusMessageType`

Message types used in #GDBusMessage.

C - `GDBusMessageType`

## `DataStreamByteOrder`

#GDataStreamByteOrder is used to ensure proper endianness of streaming data sources
across various machine architectures.

C - `GDataStreamByteOrder`

## `DataStreamNewlineType`

#GDataStreamNewlineType is used when checking for or setting the line endings for a given file.

C - `GDataStreamNewlineType`

## `DriveStartStopType`

Enumeration describing how a drive can be started/stopped.

C - `GDriveStartStopType`

## `EmblemOrigin`

GEmblemOrigin is used to add information about the origin of the emblem
to #GEmblem.

C - `GEmblemOrigin`

## `FileAttributeStatus`

Used by g_file_set_attributes_from_info() when setting file attributes.

C - `GFileAttributeStatus`

## `FileAttributeType`

The data types for file attributes.

C - `GFileAttributeType`

## `FileMonitorEvent`

Specifies what type of event a monitor event is.

C - `GFileMonitorEvent`

## `FileType`

Indicates the file's on-disk type.

C - `GFileType`

## `FilesystemPreviewType`

Indicates a hint from the file system whether files should be
previewed in a file manager. Returned as the value of the key
&num;G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.

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

C - `GIOErrorEnum`

## `IOModuleScopeFlags`

Flags for use with g_io_module_scope_new().

C - `GIOModuleScopeFlags`

## `MountOperationResult`

#GMountOperationResult is returned as a result when a request for
information is send by the mounting operation.

C - `GMountOperationResult`

## `NetworkConnectivity`

The host's network connectivity state, as reported by #GNetworkMonitor.

C - `GNetworkConnectivity`

## `NotificationPriority`

Priority levels for #GNotifications.

C - `GNotificationPriority`

## `PasswordSave`

#GPasswordSave is used to indicate the lifespan of a saved password.

&num;Gvfs stores passwords in the Gnome keyring when this flag allows it
to, and later retrieves it again from there.

C - `GPasswordSave`

## `ResolverError`

An error code used with %G_RESOLVER_ERROR in a #GError returned
from a #GResolver routine.

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

C - `GResolverRecordType`

## `ResourceError`

An error code used with %G_RESOURCE_ERROR in a #GError returned
from a #GResource routine.

C - `GResourceError`

## `SocketClientEvent`

Describes an event occurring on a #GSocketClient. See the
&num;GSocketClient::event signal for more details.

Additional values may be added to this type in the future.

C - `GSocketClientEvent`

## `SocketFamily`

The protocol family of a #GSocketAddress. (These values are
identical to the system defines %AF_INET, %AF_INET6 and %AF_UNIX,
if available.)

C - `GSocketFamily`

## `SocketListenerEvent`

Describes an event occurring on a #GSocketListener. See the
&num;GSocketListener::event signal for more details.

Additional values may be added to this type in the future.

C - `GSocketListenerEvent`

## `SocketProtocol`

A protocol identifier is specified when creating a #GSocket, which is a
family/type specific identifier, where 0 means the default protocol for
the particular family/type.

This enum contains a set of commonly available and used protocols. You
can also pass any other identifiers handled by the platform in order to
use protocols not listed here.

C - `GSocketProtocol`

## `SocketType`

Flags used when creating a #GSocket. Some protocols may not implement
all the socket types.

C - `GSocketType`

## `TlsAuthenticationMode`

The client authentication mode for a #GTlsServerConnection.

C - `GTlsAuthenticationMode`

## `TlsCertificateRequestFlags`

Flags for g_tls_interaction_request_certificate(),
g_tls_interaction_request_certificate_async(), and
g_tls_interaction_invoke_request_certificate().

C - `GTlsCertificateRequestFlags`

## `TlsDatabaseLookupFlags`

Flags for g_tls_database_lookup_certificate_for_handle(),
g_tls_database_lookup_certificate_issuer(),
and g_tls_database_lookup_certificates_issued_by().

C - `GTlsDatabaseLookupFlags`

## `TlsError`

An error code used with %G_TLS_ERROR in a #GError returned from a
TLS-related routine.

C - `GTlsError`

## `TlsInteractionResult`

#GTlsInteractionResult is returned by various functions in #GTlsInteraction
when finishing an interaction request.

C - `GTlsInteractionResult`

## `TlsRehandshakeMode`

When to allow rehandshaking. See
g_tls_connection_set_rehandshake_mode().

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

C - `GUnixSocketAddressType`

## `ZlibCompressorFormat`

Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.

C - `GZlibCompressorFormat`

