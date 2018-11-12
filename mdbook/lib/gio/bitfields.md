# `gio` bitfields

## `AppInfoCreateFlags`

Flags used when creating a #GAppInfo.

- APP_INFO_CREATE_NONE = 0
- APP_INFO_CREATE_NEEDS_TERMINAL = 1
- APP_INFO_CREATE_SUPPORTS_URIS = 2
- APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION = 4

C - `GAppInfoCreateFlags`

## `ApplicationFlags`

Flags used to define the behaviour of a #GApplication.

- APPLICATION_FLAGS_NONE = 0
- APPLICATION_IS_SERVICE = 1
- APPLICATION_IS_LAUNCHER = 2
- APPLICATION_HANDLES_OPEN = 4
- APPLICATION_HANDLES_COMMAND_LINE = 8
- APPLICATION_SEND_ENVIRONMENT = 16
- APPLICATION_NON_UNIQUE = 32
- APPLICATION_CAN_OVERRIDE_APP_ID = 64

C - `GApplicationFlags`

## `AskPasswordFlags`

#GAskPasswordFlags are used to request specific information from the
user, or to notify the user of their choices in an authentication
situation.

- ASK_PASSWORD_NEED_PASSWORD = 1
- ASK_PASSWORD_NEED_USERNAME = 2
- ASK_PASSWORD_NEED_DOMAIN = 4
- ASK_PASSWORD_SAVING_SUPPORTED = 8
- ASK_PASSWORD_ANONYMOUS_SUPPORTED = 16

C - `GAskPasswordFlags`

## `BusNameOwnerFlags`

Flags used in g_bus_own_name().

- BUS_NAME_OWNER_FLAGS_NONE = 0
- BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT = 1
- BUS_NAME_OWNER_FLAGS_REPLACE = 2
- BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE = 4

C - `GBusNameOwnerFlags`

## `BusNameWatcherFlags`

Flags used in g_bus_watch_name().

- BUS_NAME_WATCHER_FLAGS_NONE = 0
- BUS_NAME_WATCHER_FLAGS_AUTO_START = 1

C - `GBusNameWatcherFlags`

## `ConverterFlags`

Flags used when calling a g_converter_convert().

- CONVERTER_NO_FLAGS = 0
- CONVERTER_INPUT_AT_END = 1
- CONVERTER_FLUSH = 2

C - `GConverterFlags`

## `DBusCallFlags`

Flags used in g_dbus_connection_call() and similar APIs.

- DBUS_CALL_FLAGS_NONE = 0
- DBUS_CALL_FLAGS_NO_AUTO_START = 1
- DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION = 2

C - `GDBusCallFlags`

## `DBusCapabilityFlags`

Capabilities negotiated with the remote peer.

- DBUS_CAPABILITY_FLAGS_NONE = 0
- DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING = 1

C - `GDBusCapabilityFlags`

## `DBusConnectionFlags`

Flags used when creating a new #GDBusConnection.

- DBUS_CONNECTION_FLAGS_NONE = 0
- DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT = 1
- DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER = 2
- DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS = 4
- DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION = 8
- DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING = 16

C - `GDBusConnectionFlags`

## `DBusInterfaceSkeletonFlags`

Flags describing the behavior of a #GDBusInterfaceSkeleton instance.

- DBUS_INTERFACE_SKELETON_FLAGS_NONE = 0
- DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD = 1

C - `GDBusInterfaceSkeletonFlags`

## `DBusMessageFlags`

Message flags used in #GDBusMessage.

- DBUS_MESSAGE_FLAGS_NONE = 0
- DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED = 1
- DBUS_MESSAGE_FLAGS_NO_AUTO_START = 2
- DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION = 4

C - `GDBusMessageFlags`

## `DBusObjectManagerClientFlags`

Flags used when constructing a #GDBusObjectManagerClient.

- DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE = 0
- DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START = 1

C - `GDBusObjectManagerClientFlags`

## `DBusPropertyInfoFlags`

Flags describing the access control of a D-Bus property.

- DBUS_PROPERTY_INFO_FLAGS_NONE = 0
- DBUS_PROPERTY_INFO_FLAGS_READABLE = 1
- DBUS_PROPERTY_INFO_FLAGS_WRITABLE = 2

C - `GDBusPropertyInfoFlags`

## `DBusProxyFlags`

Flags used when constructing an instance of a #GDBusProxy derived class.

- DBUS_PROXY_FLAGS_NONE = 0
- DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES = 1
- DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS = 2
- DBUS_PROXY_FLAGS_DO_NOT_AUTO_START = 4
- DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES = 8
- DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION = 16

C - `GDBusProxyFlags`

## `DBusSendMessageFlags`

Flags used when sending #GDBusMessages on a #GDBusConnection.

- DBUS_SEND_MESSAGE_FLAGS_NONE = 0
- DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL = 1

C - `GDBusSendMessageFlags`

## `DBusServerFlags`

Flags used when creating a #GDBusServer.

- DBUS_SERVER_FLAGS_NONE = 0
- DBUS_SERVER_FLAGS_RUN_IN_THREAD = 1
- DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS = 2

C - `GDBusServerFlags`

## `DBusSignalFlags`

Flags used when subscribing to signals via g_dbus_connection_signal_subscribe().

- DBUS_SIGNAL_FLAGS_NONE = 0
- DBUS_SIGNAL_FLAGS_NO_MATCH_RULE = 1
- DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE = 2
- DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH = 4

C - `GDBusSignalFlags`

## `DBusSubtreeFlags`

Flags passed to g_dbus_connection_register_subtree().

- DBUS_SUBTREE_FLAGS_NONE = 0
- DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES = 1

C - `GDBusSubtreeFlags`

## `DriveStartFlags`

Flags used when starting a drive.

- DRIVE_START_NONE = 0

C - `GDriveStartFlags`

## `FileAttributeInfoFlags`

Flags specifying the behaviour of an attribute.

- FILE_ATTRIBUTE_INFO_NONE = 0
- FILE_ATTRIBUTE_INFO_COPY_WITH_FILE = 1
- FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED = 2

C - `GFileAttributeInfoFlags`

## `FileCopyFlags`

Flags used when copying or moving files.

- FILE_COPY_NONE = 0
- FILE_COPY_OVERWRITE = 1
- FILE_COPY_BACKUP = 2
- FILE_COPY_NOFOLLOW_SYMLINKS = 4
- FILE_COPY_ALL_METADATA = 8
- FILE_COPY_NO_FALLBACK_FOR_MOVE = 16
- FILE_COPY_TARGET_DEFAULT_PERMS = 32

C - `GFileCopyFlags`

## `FileCreateFlags`

Flags used when an operation may create a file.

- FILE_CREATE_NONE = 0
- FILE_CREATE_PRIVATE = 1
- FILE_CREATE_REPLACE_DESTINATION = 2

C - `GFileCreateFlags`

## `FileMeasureFlags`

Flags that can be used with g_file_measure_disk_usage().

- FILE_MEASURE_NONE = 0
- FILE_MEASURE_REPORT_ANY_ERROR = 2
- FILE_MEASURE_APPARENT_SIZE = 4
- FILE_MEASURE_NO_XDEV = 8

C - `GFileMeasureFlags`

## `FileMonitorFlags`

Flags used to set what a #GFileMonitor will watch for.

- FILE_MONITOR_NONE = 0
- FILE_MONITOR_WATCH_MOUNTS = 1
- FILE_MONITOR_SEND_MOVED = 2
- FILE_MONITOR_WATCH_HARD_LINKS = 4
- FILE_MONITOR_WATCH_MOVES = 8

C - `GFileMonitorFlags`

## `FileQueryInfoFlags`

Flags used when querying a #GFileInfo.

- FILE_QUERY_INFO_NONE = 0
- FILE_QUERY_INFO_NOFOLLOW_SYMLINKS = 1

C - `GFileQueryInfoFlags`

## `IOStreamSpliceFlags`

GIOStreamSpliceFlags determine how streams should be spliced.

- IO_STREAM_SPLICE_NONE = 0
- IO_STREAM_SPLICE_CLOSE_STREAM1 = 1
- IO_STREAM_SPLICE_CLOSE_STREAM2 = 2
- IO_STREAM_SPLICE_WAIT_FOR_BOTH = 4

C - `GIOStreamSpliceFlags`

## `MountMountFlags`

Flags used when mounting a mount.

- MOUNT_MOUNT_NONE = 0

C - `GMountMountFlags`

## `MountUnmountFlags`

Flags used when an unmounting a mount.

- MOUNT_UNMOUNT_NONE = 0
- MOUNT_UNMOUNT_FORCE = 1

C - `GMountUnmountFlags`

## `OutputStreamSpliceFlags`

GOutputStreamSpliceFlags determine how streams should be spliced.

- OUTPUT_STREAM_SPLICE_NONE = 0
- OUTPUT_STREAM_SPLICE_CLOSE_SOURCE = 1
- OUTPUT_STREAM_SPLICE_CLOSE_TARGET = 2

C - `GOutputStreamSpliceFlags`

## `ResourceFlags`

GResourceFlags give information about a particular file inside a resource
bundle.

- RESOURCE_FLAGS_NONE = 0
- RESOURCE_FLAGS_COMPRESSED = 1

C - `GResourceFlags`

## `ResourceLookupFlags`

GResourceLookupFlags determine how resource path lookups are handled.

- RESOURCE_LOOKUP_FLAGS_NONE = 0

C - `GResourceLookupFlags`

## `SettingsBindFlags`

Flags used when creating a binding. These flags determine in which
direction the binding works. The default is to synchronize in both
directions.

- SETTINGS_BIND_DEFAULT = 0
- SETTINGS_BIND_GET = 1
- SETTINGS_BIND_SET = 2
- SETTINGS_BIND_NO_SENSITIVITY = 4
- SETTINGS_BIND_GET_NO_CHANGES = 8
- SETTINGS_BIND_INVERT_BOOLEAN = 16

C - `GSettingsBindFlags`

## `SocketMsgFlags`

Flags used in g_socket_receive_message() and g_socket_send_message().
The flags listed in the enum are some commonly available flags, but the
values used for them are the same as on the platform, and any other flags
are passed in/out as is. So to use a platform specific flag, just include
the right system header and pass in the flag.

- SOCKET_MSG_NONE = 0
- SOCKET_MSG_OOB = 1
- SOCKET_MSG_PEEK = 2
- SOCKET_MSG_DONTROUTE = 4

C - `GSocketMsgFlags`

## `SubprocessFlags`

Flags to define the behaviour of a #GSubprocess.

Note that the default for stdin is to redirect from `/dev/null`.  For
stdout and stderr the default are for them to inherit the
corresponding descriptor from the calling process.

Note that it is a programmer error to mix 'incompatible' flags.  For
example, you may not request both %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
%G_SUBPROCESS_FLAGS_STDOUT_SILENCE.

- SUBPROCESS_FLAGS_NONE = 0
- SUBPROCESS_FLAGS_STDIN_PIPE = 1
- SUBPROCESS_FLAGS_STDIN_INHERIT = 2
- SUBPROCESS_FLAGS_STDOUT_PIPE = 4
- SUBPROCESS_FLAGS_STDOUT_SILENCE = 8
- SUBPROCESS_FLAGS_STDERR_PIPE = 16
- SUBPROCESS_FLAGS_STDERR_SILENCE = 32
- SUBPROCESS_FLAGS_STDERR_MERGE = 64
- SUBPROCESS_FLAGS_INHERIT_FDS = 128

C - `GSubprocessFlags`

## `TestDBusFlags`

Flags to define future #GTestDBus behaviour.

- TEST_DBUS_NONE = 0

C - `GTestDBusFlags`

## `TlsCertificateFlags`

A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
&num;GTlsConnection::accept-certificate).

- TLS_CERTIFICATE_UNKNOWN_CA = 1
- TLS_CERTIFICATE_BAD_IDENTITY = 2
- TLS_CERTIFICATE_NOT_ACTIVATED = 4
- TLS_CERTIFICATE_EXPIRED = 8
- TLS_CERTIFICATE_REVOKED = 16
- TLS_CERTIFICATE_INSECURE = 32
- TLS_CERTIFICATE_GENERIC_ERROR = 64
- TLS_CERTIFICATE_VALIDATE_ALL = 127

C - `GTlsCertificateFlags`

## `TlsDatabaseVerifyFlags`

Flags for g_tls_database_verify_chain().

- TLS_DATABASE_VERIFY_NONE = 0

C - `GTlsDatabaseVerifyFlags`

## `TlsPasswordFlags`

Various flags for the password.

- TLS_PASSWORD_NONE = 0
- TLS_PASSWORD_RETRY = 2
- TLS_PASSWORD_MANY_TRIES = 4
- TLS_PASSWORD_FINAL_TRY = 8

C - `GTlsPasswordFlags`

RIT_FDS = %!s(int=128)
## `TestDBusFlags`

Flags to define future #GTestDBus behaviour.

C - `GTestDBusFlags`

- TEST_DBUS_NONE = %!s(int=0)
## `TlsCertificateFlags`

A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
&num;GTlsConnection::accept-certificate).

C - `GTlsCertificateFlags`

- TLS_CERTIFICATE_UNKNOWN_CA = %!s(int=1)
- TLS_CERTIFICATE_BAD_IDENTITY = %!s(int=2)
- TLS_CERTIFICATE_NOT_ACTIVATED = %!s(int=4)
- TLS_CERTIFICATE_EXPIRED = %!s(int=8)
- TLS_CERTIFICATE_REVOKED = %!s(int=16)
- TLS_CERTIFICATE_INSECURE = %!s(int=32)
- TLS_CERTIFICATE_GENERIC_ERROR = %!s(int=64)
- TLS_CERTIFICATE_VALIDATE_ALL = %!s(int=127)
## `TlsDatabaseVerifyFlags`

Flags for g_tls_database_verify_chain().

C - `GTlsDatabaseVerifyFlags`

- TLS_DATABASE_VERIFY_NONE = %!s(int=0)
## `TlsPasswordFlags`

Various flags for the password.

C - `GTlsPasswordFlags`

- TLS_PASSWORD_NONE = %!s(int=0)
- TLS_PASSWORD_RETRY = %!s(int=2)
- TLS_PASSWORD_MANY_TRIES = %!s(int=4)
- TLS_PASSWORD_FINAL_TRY = %!s(int=8)
