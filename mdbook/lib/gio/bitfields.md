# `gio` bitfields

## `AppInfoCreateFlags`

Flags used when creating a #GAppInfo.

C - `GAppInfoCreateFlags`

## `ApplicationFlags`

Flags used to define the behaviour of a #GApplication.

C - `GApplicationFlags`

## `AskPasswordFlags`

#GAskPasswordFlags are used to request specific information from the
user, or to notify the user of their choices in an authentication
situation.

C - `GAskPasswordFlags`

## `BusNameOwnerFlags`

Flags used in g_bus_own_name().

C - `GBusNameOwnerFlags`

## `BusNameWatcherFlags`

Flags used in g_bus_watch_name().

C - `GBusNameWatcherFlags`

## `ConverterFlags`

Flags used when calling a g_converter_convert().

C - `GConverterFlags`

## `DBusCallFlags`

Flags used in g_dbus_connection_call() and similar APIs.

C - `GDBusCallFlags`

## `DBusCapabilityFlags`

Capabilities negotiated with the remote peer.

C - `GDBusCapabilityFlags`

## `DBusConnectionFlags`

Flags used when creating a new #GDBusConnection.

C - `GDBusConnectionFlags`

## `DBusInterfaceSkeletonFlags`

Flags describing the behavior of a #GDBusInterfaceSkeleton instance.

C - `GDBusInterfaceSkeletonFlags`

## `DBusMessageFlags`

Message flags used in #GDBusMessage.

C - `GDBusMessageFlags`

## `DBusObjectManagerClientFlags`

Flags used when constructing a #GDBusObjectManagerClient.

C - `GDBusObjectManagerClientFlags`

## `DBusPropertyInfoFlags`

Flags describing the access control of a D-Bus property.

C - `GDBusPropertyInfoFlags`

## `DBusProxyFlags`

Flags used when constructing an instance of a #GDBusProxy derived class.

C - `GDBusProxyFlags`

## `DBusSendMessageFlags`

Flags used when sending #GDBusMessages on a #GDBusConnection.

C - `GDBusSendMessageFlags`

## `DBusServerFlags`

Flags used when creating a #GDBusServer.

C - `GDBusServerFlags`

## `DBusSignalFlags`

Flags used when subscribing to signals via g_dbus_connection_signal_subscribe().

C - `GDBusSignalFlags`

## `DBusSubtreeFlags`

Flags passed to g_dbus_connection_register_subtree().

C - `GDBusSubtreeFlags`

## `DriveStartFlags`

Flags used when starting a drive.

C - `GDriveStartFlags`

## `FileAttributeInfoFlags`

Flags specifying the behaviour of an attribute.

C - `GFileAttributeInfoFlags`

## `FileCopyFlags`

Flags used when copying or moving files.

C - `GFileCopyFlags`

## `FileCreateFlags`

Flags used when an operation may create a file.

C - `GFileCreateFlags`

## `FileMeasureFlags`

Flags that can be used with g_file_measure_disk_usage().

C - `GFileMeasureFlags`

## `FileMonitorFlags`

Flags used to set what a #GFileMonitor will watch for.

C - `GFileMonitorFlags`

## `FileQueryInfoFlags`

Flags used when querying a #GFileInfo.

C - `GFileQueryInfoFlags`

## `IOStreamSpliceFlags`

GIOStreamSpliceFlags determine how streams should be spliced.

C - `GIOStreamSpliceFlags`

## `MountMountFlags`

Flags used when mounting a mount.

C - `GMountMountFlags`

## `MountUnmountFlags`

Flags used when an unmounting a mount.

C - `GMountUnmountFlags`

## `OutputStreamSpliceFlags`

GOutputStreamSpliceFlags determine how streams should be spliced.

C - `GOutputStreamSpliceFlags`

## `ResourceFlags`

GResourceFlags give information about a particular file inside a resource
bundle.

C - `GResourceFlags`

## `ResourceLookupFlags`

GResourceLookupFlags determine how resource path lookups are handled.

C - `GResourceLookupFlags`

## `SettingsBindFlags`

Flags used when creating a binding. These flags determine in which
direction the binding works. The default is to synchronize in both
directions.

C - `GSettingsBindFlags`

## `SocketMsgFlags`

Flags used in g_socket_receive_message() and g_socket_send_message().
The flags listed in the enum are some commonly available flags, but the
values used for them are the same as on the platform, and any other flags
are passed in/out as is. So to use a platform specific flag, just include
the right system header and pass in the flag.

C - `GSocketMsgFlags`

## `SubprocessFlags`

Flags to define the behaviour of a #GSubprocess.

Note that the default for stdin is to redirect from `/dev/null`.  For
stdout and stderr the default are for them to inherit the
corresponding descriptor from the calling process.

Note that it is a programmer error to mix 'incompatible' flags.  For
example, you may not request both %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
%G_SUBPROCESS_FLAGS_STDOUT_SILENCE.

C - `GSubprocessFlags`

## `TestDBusFlags`

Flags to define future #GTestDBus behaviour.

C - `GTestDBusFlags`

## `TlsCertificateFlags`

A set of flags describing TLS certification validation. This can be
used to set which validation steps to perform (eg, with
g_tls_client_connection_set_validation_flags()), or to describe why
a particular certificate was rejected (eg, in
&num;GTlsConnection::accept-certificate).

C - `GTlsCertificateFlags`

## `TlsDatabaseVerifyFlags`

Flags for g_tls_database_verify_chain().

C - `GTlsDatabaseVerifyFlags`

## `TlsPasswordFlags`

Various flags for the password.

C - `GTlsPasswordFlags`

