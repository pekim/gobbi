# `gio` bitfields

## `AppInfoCreateFlags`

Flags used when creating a #GAppInfo.

C - `GAppInfoCreateFlags`

### APP_INFO_CREATE_NONE = 0
No flags.

### APP_INFO_CREATE_NEEDS_TERMINAL = 1
Application opens in a terminal window.

### APP_INFO_CREATE_SUPPORTS_URIS = 2
Application supports URI arguments.

### APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION = 4
Application supports startup notification. Since 2.26


## `ApplicationFlags`

Flags used to define the behaviour of a #GApplication.

C - `GApplicationFlags`

### APPLICATION_FLAGS_NONE = 0
Default

### APPLICATION_IS_SERVICE = 1
Run as a service. In this mode, registration
     fails if the service is already running, and the application
     will initially wait up to 10 seconds for an initial activation
     message to arrive.

### APPLICATION_IS_LAUNCHER = 2
Don't try to become the primary instance.

### APPLICATION_HANDLES_OPEN = 4
This application handles opening files (in
    the primary instance). Note that this flag only affects the default
    implementation of local_command_line(), and has no effect if
    %G_APPLICATION_HANDLES_COMMAND_LINE is given.
    See g_application_run() for details.

### APPLICATION_HANDLES_COMMAND_LINE = 8
This application handles command line
    arguments (in the primary instance). Note that this flag only affect
    the default implementation of local_command_line().
    See g_application_run() for details.

### APPLICATION_SEND_ENVIRONMENT = 16
Send the environment of the
    launching process to the primary instance. Set this flag if your
    application is expected to behave differently depending on certain
    environment variables. For instance, an editor might be expected
    to use the `GIT_COMMITTER_NAME` environment variable
    when editing a git commit message. The environment is available
    to the #GApplication::command-line signal handler, via
    g_application_command_line_getenv().

### APPLICATION_NON_UNIQUE = 32
Make no attempts to do any of the typical
    single-instance application negotiation, even if the application
    ID is given.  The application neither attempts to become the
    owner of the application ID nor does it check if an existing
    owner already exists.  Everything occurs in the local process.
    Since: 2.30.

### APPLICATION_CAN_OVERRIDE_APP_ID = 64
Allow users to override the
    application ID from the command line with `--gapplication-app-id`.
    Since: 2.48


## `AskPasswordFlags`

#GAskPasswordFlags are used to request specific information from the
user, or to notify the user of their choices in an authentication
situation.

C - `GAskPasswordFlags`

### ASK_PASSWORD_NEED_PASSWORD = 1
operation requires a password.

### ASK_PASSWORD_NEED_USERNAME = 2
operation requires a username.

### ASK_PASSWORD_NEED_DOMAIN = 4
operation requires a domain.

### ASK_PASSWORD_SAVING_SUPPORTED = 8
operation supports saving settings.

### ASK_PASSWORD_ANONYMOUS_SUPPORTED = 16
operation supports anonymous users.


## `BusNameOwnerFlags`

Flags used in g_bus_own_name().

C - `GBusNameOwnerFlags`

### BUS_NAME_OWNER_FLAGS_NONE = 0
No flags set.

### BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT = 1
Allow another message bus connection to claim the name.

### BUS_NAME_OWNER_FLAGS_REPLACE = 2
If another message bus connection owns the name and have
specified #G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT, then take the name from the other connection.

### BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE = 4
If another message bus connection owns the name, immediately
return an error from g_bus_own_name() rather than entering the waiting queue for that name. (Since 2.54)


## `BusNameWatcherFlags`

Flags used in g_bus_watch_name().

C - `GBusNameWatcherFlags`

### BUS_NAME_WATCHER_FLAGS_NONE = 0
No flags set.

### BUS_NAME_WATCHER_FLAGS_AUTO_START = 1
If no-one owns the name when
beginning to watch the name, ask the bus to launch an owner for the
name.


## `ConverterFlags`

Flags used when calling a g_converter_convert().

C - `GConverterFlags`

### CONVERTER_NO_FLAGS = 0
No flags.

### CONVERTER_INPUT_AT_END = 1
At end of input data

### CONVERTER_FLUSH = 2
Flush data


## `DBusCallFlags`

Flags used in g_dbus_connection_call() and similar APIs.

C - `GDBusCallFlags`

### DBUS_CALL_FLAGS_NONE = 0
No flags set.

### DBUS_CALL_FLAGS_NO_AUTO_START = 1
The bus must not launch
an owner for the destination name in response to this method
invocation.

### DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION = 2
the caller is prepared to
wait for interactive authorization. Since 2.46.


## `DBusCapabilityFlags`

Capabilities negotiated with the remote peer.

C - `GDBusCapabilityFlags`

### DBUS_CAPABILITY_FLAGS_NONE = 0
No flags set.

### DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING = 1
The connection
supports exchanging UNIX file descriptors with the remote peer.


## `DBusConnectionFlags`

Flags used when creating a new #GDBusConnection.

C - `GDBusConnectionFlags`

### DBUS_CONNECTION_FLAGS_NONE = 0
No flags set.

### DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT = 1
Perform authentication against server.

### DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER = 2
Perform authentication against client.

### DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS = 4
When
authenticating as a server, allow the anonymous authentication
method.

### DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION = 8
Pass this flag if connecting to a peer that is a
message bus. This means that the Hello() method will be invoked as part of the connection setup.

### DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING = 16
If set, processing of D-Bus messages is
delayed until g_dbus_connection_start_message_processing() is called.


## `DBusInterfaceSkeletonFlags`

Flags describing the behavior of a #GDBusInterfaceSkeleton instance.

C - `GDBusInterfaceSkeletonFlags`

### DBUS_INTERFACE_SKELETON_FLAGS_NONE = 0
No flags set.

### DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD = 1
Each method invocation is handled in
  a thread dedicated to the invocation. This means that the method implementation can use blocking IO
  without blocking any other part of the process. It also means that the method implementation must
  use locking to access data structures used by other threads.


## `DBusMessageFlags`

Message flags used in #GDBusMessage.

C - `GDBusMessageFlags`

### DBUS_MESSAGE_FLAGS_NONE = 0
No flags set.

### DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED = 1
A reply is not expected.

### DBUS_MESSAGE_FLAGS_NO_AUTO_START = 2
The bus must not launch an
owner for the destination name in response to this message.

### DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION = 4
If set on a method
call, this flag means that the caller is prepared to wait for interactive
authorization. Since 2.46.


## `DBusObjectManagerClientFlags`

Flags used when constructing a #GDBusObjectManagerClient.

C - `GDBusObjectManagerClientFlags`

### DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE = 0
No flags set.

### DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START = 1
If not set and the
  manager is for a well-known name, then request the bus to launch
  an owner for the name if no-one owns the name. This flag can only
  be used in managers for well-known names.


## `DBusPropertyInfoFlags`

Flags describing the access control of a D-Bus property.

C - `GDBusPropertyInfoFlags`

### DBUS_PROPERTY_INFO_FLAGS_NONE = 0
No flags set.

### DBUS_PROPERTY_INFO_FLAGS_READABLE = 1
Property is readable.

### DBUS_PROPERTY_INFO_FLAGS_WRITABLE = 2
Property is writable.


## `DBusProxyFlags`

Flags used when constructing an instance of a #GDBusProxy derived class.

C - `GDBusProxyFlags`

### DBUS_PROXY_FLAGS_NONE = 0
No flags set.

### DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES = 1
Don't load properties.

### DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS = 2
Don't connect to signals on the remote object.

### DBUS_PROXY_FLAGS_DO_NOT_AUTO_START = 4
If the proxy is for a well-known name,
do not ask the bus to launch an owner during proxy initialization or a method call.
This flag is only meaningful in proxies for well-known names.

### DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES = 8
If set, the property value for any __invalidated property__ will be (asynchronously) retrieved upon receiving the [`PropertiesChanged`](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-properties) D-Bus signal and the property will not cause emission of the #GDBusProxy::g-properties-changed signal. When the value is received the #GDBusProxy::g-properties-changed signal is emitted for the property along with the retrieved value. Since 2.32.

### DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION = 16
If the proxy is for a well-known name,
do not ask the bus to launch an owner during proxy initialization, but allow it to be
autostarted by a method call. This flag is only meaningful in proxies for well-known names,
and only if %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START is not also specified.


## `DBusSendMessageFlags`

Flags used when sending #GDBusMessages on a #GDBusConnection.

C - `GDBusSendMessageFlags`

### DBUS_SEND_MESSAGE_FLAGS_NONE = 0
No flags set.

### DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL = 1
Do not automatically
assign a serial number from the #GDBusConnection object when
sending a message.


## `DBusServerFlags`

Flags used when creating a #GDBusServer.

C - `GDBusServerFlags`

### DBUS_SERVER_FLAGS_NONE = 0
No flags set.

### DBUS_SERVER_FLAGS_RUN_IN_THREAD = 1
All #GDBusServer::new-connection
signals will run in separated dedicated threads (see signal for
details).

### DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS = 2
Allow the anonymous
authentication method.


## `DBusSignalFlags`

Flags used when subscribing to signals via g_dbus_connection_signal_subscribe().

C - `GDBusSignalFlags`

### DBUS_SIGNAL_FLAGS_NONE = 0
No flags set.

### DBUS_SIGNAL_FLAGS_NO_MATCH_RULE = 1
Don't actually send the AddMatch
D-Bus call for this signal subscription.  This gives you more control
over which match rules you add (but you must add them manually).

### DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE = 2
Match first arguments that
contain a bus or interface name with the given namespace.

### DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH = 4
Match first arguments that
contain an object path that is either equivalent to the given path,
or one of the paths is a subpath of the other.


## `DBusSubtreeFlags`

Flags passed to g_dbus_connection_register_subtree().

C - `GDBusSubtreeFlags`

### DBUS_SUBTREE_FLAGS_NONE = 0
No flags set.

### DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES = 1
Method calls to objects not in the enumerated range
                                                      will still be dispatched. This is useful if you want
                                                      to dynamically spawn objects in the subtree.


## `DriveStartFlags`

Flags used when starting a drive.

C - `GDriveStartFlags`

### DRIVE_START_NONE = 0
No flags set.


## `FileAttributeInfoFlags`

Flags specifying the behaviour of an attribute.

C - `GFileAttributeInfoFlags`

### FILE_ATTRIBUTE_INFO_NONE = 0
no flags set.

### FILE_ATTRIBUTE_INFO_COPY_WITH_FILE = 1
copy the attribute values when the file is copied.

### FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED = 2
copy the attribute values when the file is moved.


## `FileCopyFlags`

Flags used when copying or moving files.

C - `GFileCopyFlags`

### FILE_COPY_NONE = 0
No flags set.

### FILE_COPY_OVERWRITE = 1
Overwrite any existing files

### FILE_COPY_BACKUP = 2
Make a backup of any existing files.

### FILE_COPY_NOFOLLOW_SYMLINKS = 4
Don't follow symlinks.

### FILE_COPY_ALL_METADATA = 8
Copy all file metadata instead of just default set used for copy (see #GFileInfo).

### FILE_COPY_NO_FALLBACK_FOR_MOVE = 16
Don't use copy and delete fallback if native move not supported.

### FILE_COPY_TARGET_DEFAULT_PERMS = 32
Leaves target file with default perms, instead of setting the source file perms.


## `FileCreateFlags`

Flags used when an operation may create a file.

C - `GFileCreateFlags`

### FILE_CREATE_NONE = 0
No flags set.

### FILE_CREATE_PRIVATE = 1
Create a file that can only be
   accessed by the current user.

### FILE_CREATE_REPLACE_DESTINATION = 2
Replace the destination
   as if it didn't exist before. Don't try to keep any old
   permissions, replace instead of following links. This
   is generally useful if you're doing a "copy over"
   rather than a "save new version of" replace operation.
   You can think of it as "unlink destination" before
   writing to it, although the implementation may not
   be exactly like that. Since 2.20


## `FileMeasureFlags`

Flags that can be used with g_file_measure_disk_usage().

C - `GFileMeasureFlags`

### FILE_MEASURE_NONE = 0
No flags set.

### FILE_MEASURE_REPORT_ANY_ERROR = 2
Report any error encountered
  while traversing the directory tree.  Normally errors are only
  reported for the toplevel file.

### FILE_MEASURE_APPARENT_SIZE = 4
Tally usage based on apparent file
  sizes.  Normally, the block-size is used, if available, as this is a
  more accurate representation of disk space used.
  Compare with `du --apparent-size`.

### FILE_MEASURE_NO_XDEV = 8
Do not cross mount point boundaries.
  Compare with `du -x`.


## `FileMonitorFlags`

Flags used to set what a #GFileMonitor will watch for.

C - `GFileMonitorFlags`

### FILE_MONITOR_NONE = 0
No flags set.

### FILE_MONITOR_WATCH_MOUNTS = 1
Watch for mount events.

### FILE_MONITOR_SEND_MOVED = 2
Pair DELETED and CREATED events caused
  by file renames (moves) and send a single G_FILE_MONITOR_EVENT_MOVED
  event instead (NB: not supported on all backends; the default
  behaviour -without specifying this flag- is to send single DELETED
  and CREATED events).  Deprecated since 2.46: use
  %G_FILE_MONITOR_WATCH_MOVES instead.

### FILE_MONITOR_WATCH_HARD_LINKS = 4
Watch for changes to the file made
  via another hard link. Since 2.36.

### FILE_MONITOR_WATCH_MOVES = 8
Watch for rename operations on a
  monitored directory.  This causes %G_FILE_MONITOR_EVENT_RENAMED,
  %G_FILE_MONITOR_EVENT_MOVED_IN and %G_FILE_MONITOR_EVENT_MOVED_OUT
  events to be emitted when possible.  Since: 2.46.


## `FileQueryInfoFlags`

Flags used when querying a #GFileInfo.

C - `GFileQueryInfoFlags`

### FILE_QUERY_INFO_NONE = 0
No flags set.

### FILE_QUERY_INFO_NOFOLLOW_SYMLINKS = 1
Don't follow symlinks.


## `IOStreamSpliceFlags`

GIOStreamSpliceFlags determine how streams should be spliced.

C - `GIOStreamSpliceFlags`

### IO_STREAM_SPLICE_NONE = 0
Do not close either stream.

### IO_STREAM_SPLICE_CLOSE_STREAM1 = 1
Close the first stream after
    the splice.

### IO_STREAM_SPLICE_CLOSE_STREAM2 = 2
Close the second stream after
    the splice.

### IO_STREAM_SPLICE_WAIT_FOR_BOTH = 4
Wait for both splice operations to finish
    before calling the callback.


## `MountMountFlags`

Flags used when mounting a mount.

C - `GMountMountFlags`

### MOUNT_MOUNT_NONE = 0
No flags set.


## `MountUnmountFlags`

Flags used when an unmounting a mount.

C - `GMountUnmountFlags`

### MOUNT_UNMOUNT_NONE = 0
No flags set.

### MOUNT_UNMOUNT_FORCE = 1
Unmount even if there are outstanding
 file operations on the mount.


## `OutputStreamSpliceFlags`

GOutputStreamSpliceFlags determine how streams should be spliced.

C - `GOutputStreamSpliceFlags`

### OUTPUT_STREAM_SPLICE_NONE = 0
Do not close either stream.

### OUTPUT_STREAM_SPLICE_CLOSE_SOURCE = 1
Close the source stream after
    the splice.

### OUTPUT_STREAM_SPLICE_CLOSE_TARGET = 2
Close the target stream after
    the splice.


## `ResourceFlags`

GResourceFlags give information about a particular file inside a resource
bundle.

C - `GResourceFlags`

### RESOURCE_FLAGS_NONE = 0
No flags set.

### RESOURCE_FLAGS_COMPRESSED = 1
The file is compressed.


## `ResourceLookupFlags`

GResourceLookupFlags determine how resource path lookups are handled.

C - `GResourceLookupFlags`

### RESOURCE_LOOKUP_FLAGS_NONE = 0
No flags set.


## `SettingsBindFlags`

Flags used when creating a binding. These flags determine in which
direction the binding works. The default is to synchronize in both
directions.

C - `GSettingsBindFlags`

### SETTINGS_BIND_DEFAULT = 0
Equivalent to `G_SETTINGS_BIND_GET|G_SETTINGS_BIND_SET`

### SETTINGS_BIND_GET = 1
Update the #GObject property when the setting changes.
    It is an error to use this flag if the property is not writable.

### SETTINGS_BIND_SET = 2
Update the setting when the #GObject property changes.
    It is an error to use this flag if the property is not readable.

### SETTINGS_BIND_NO_SENSITIVITY = 4
Do not try to bind a "sensitivity" property to the writability of the setting

### SETTINGS_BIND_GET_NO_CHANGES = 8
When set in addition to #G_SETTINGS_BIND_GET, set the #GObject property
    value initially from the setting, but do not listen for changes of the setting

### SETTINGS_BIND_INVERT_BOOLEAN = 16
When passed to g_settings_bind(), uses a pair of mapping functions that invert
    the boolean value when mapping between the setting and the property.  The setting and property must both
    be booleans.  You cannot pass this flag to g_settings_bind_with_mapping().


## `SocketMsgFlags`

Flags used in g_socket_receive_message() and g_socket_send_message().
The flags listed in the enum are some commonly available flags, but the
values used for them are the same as on the platform, and any other flags
are passed in/out as is. So to use a platform specific flag, just include
the right system header and pass in the flag.

C - `GSocketMsgFlags`

### SOCKET_MSG_NONE = 0
No flags.

### SOCKET_MSG_OOB = 1
Request to send/receive out of band data.

### SOCKET_MSG_PEEK = 2
Read data from the socket without removing it from
    the queue.

### SOCKET_MSG_DONTROUTE = 4
Don't use a gateway to send out the packet,
    only send to hosts on directly connected networks.


## `SubprocessFlags`

Flags to define the behaviour of a #GSubprocess.

Note that the default for stdin is to redirect from `/dev/null`.  For
stdout and stderr the default are for them to inherit the
corresponding descriptor from the calling process.

Note that it is a programmer error to mix 'incompatible' flags.  For
example, you may not request both %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
%G_SUBPROCESS_FLAGS_STDOUT_SILENCE.

C - `GSubprocessFlags`

### SUBPROCESS_FLAGS_NONE = 0
No flags.

### SUBPROCESS_FLAGS_STDIN_PIPE = 1
create a pipe for the stdin of the
  spawned process that can be accessed with
  g_subprocess_get_stdin_pipe().

### SUBPROCESS_FLAGS_STDIN_INHERIT = 2
stdin is inherited from the
  calling process.

### SUBPROCESS_FLAGS_STDOUT_PIPE = 4
create a pipe for the stdout of the
  spawned process that can be accessed with
  g_subprocess_get_stdout_pipe().

### SUBPROCESS_FLAGS_STDOUT_SILENCE = 8
silence the stdout of the spawned
  process (ie: redirect to `/dev/null`).

### SUBPROCESS_FLAGS_STDERR_PIPE = 16
create a pipe for the stderr of the
  spawned process that can be accessed with
  g_subprocess_get_stderr_pipe().

### SUBPROCESS_FLAGS_STDERR_SILENCE = 32
silence the stderr of the spawned
  process (ie: redirect to `/dev/null`).

### SUBPROCESS_FLAGS_STDERR_MERGE = 64
merge the stderr of the spawned
  process with whatever the stdout happens to be.  This is a good way
  of directing both streams to a common log file, for example.

### SUBPROCESS_FLAGS_INHERIT_FDS = 128
spawned processes will inherit the
  file descriptors of their parent, unless those descriptors have
  been explicitly marked as close-on-exec.  This flag has no effect
  over the "standard" file descriptors (stdin, stdout, stderr).


## `TestDBusFlags`

Flags to define future #GTestDBus behaviour.

C - `GTestDBusFlags`

### TEST_DBUS_NONE = 0
No flags.


## `TlsCertificateFlags`

A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
&num;GTlsConnection::accept-certificate).

C - `GTlsCertificateFlags`

### TLS_CERTIFICATE_UNKNOWN_CA = 1
The signing certificate authority is
  not known.

### TLS_CERTIFICATE_BAD_IDENTITY = 2
The certificate does not match the
  expected identity of the site that it was retrieved from.

### TLS_CERTIFICATE_NOT_ACTIVATED = 4
The certificate's activation time
  is still in the future

### TLS_CERTIFICATE_EXPIRED = 8
The certificate has expired

### TLS_CERTIFICATE_REVOKED = 16
The certificate has been revoked
  according to the #GTlsConnection's certificate revocation list.

### TLS_CERTIFICATE_INSECURE = 32
The certificate's algorithm is
  considered insecure.

### TLS_CERTIFICATE_GENERIC_ERROR = 64
Some other error occurred validating
  the certificate

### TLS_CERTIFICATE_VALIDATE_ALL = 127
the combination of all of the above
  flags


## `TlsDatabaseVerifyFlags`

Flags for g_tls_database_verify_chain().

C - `GTlsDatabaseVerifyFlags`

### TLS_DATABASE_VERIFY_NONE = 0
No verification flags


## `TlsPasswordFlags`

Various flags for the password.

C - `GTlsPasswordFlags`

### TLS_PASSWORD_NONE = 0
No flags

### TLS_PASSWORD_RETRY = 2
The password was wrong, and the user should retry.

### TLS_PASSWORD_MANY_TRIES = 4
Hint to the user that the password has been
   wrong many times, and the user may not have many chances left.

### TLS_PASSWORD_FINAL_TRY = 8
Hint to the user that this is the last try to get
   this password right.


