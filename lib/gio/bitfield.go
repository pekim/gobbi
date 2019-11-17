// Code generated - DO NOT EDIT.

package gio

// AppInfoCreateFlags is a representation of the C type GAppInfoCreateFlags.
type AppInfoCreateFlags int

const (
	// AppInfoCreateFlags_None is a representation of the C type G_APP_INFO_CREATE_NONE.
	AppInfoCreateFlags_None AppInfoCreateFlags = 0
	// AppInfoCreateFlags_NeedsTerminal is a representation of the C type G_APP_INFO_CREATE_NEEDS_TERMINAL.
	AppInfoCreateFlags_NeedsTerminal AppInfoCreateFlags = 1
	// AppInfoCreateFlags_SupportsUris is a representation of the C type G_APP_INFO_CREATE_SUPPORTS_URIS.
	AppInfoCreateFlags_SupportsUris AppInfoCreateFlags = 2
	// AppInfoCreateFlags_SupportsStartupNotification is a representation of the C type G_APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION.
	AppInfoCreateFlags_SupportsStartupNotification AppInfoCreateFlags = 4
)

// ApplicationFlags is a representation of the C type GApplicationFlags.
//
// since 2.28
type ApplicationFlags int

const (
	// ApplicationFlags_FlagsNone is a representation of the C type G_APPLICATION_FLAGS_NONE.
	ApplicationFlags_FlagsNone ApplicationFlags = 0
	// ApplicationFlags_IsService is a representation of the C type G_APPLICATION_IS_SERVICE.
	ApplicationFlags_IsService ApplicationFlags = 1
	// ApplicationFlags_IsLauncher is a representation of the C type G_APPLICATION_IS_LAUNCHER.
	ApplicationFlags_IsLauncher ApplicationFlags = 2
	// ApplicationFlags_HandlesOpen is a representation of the C type G_APPLICATION_HANDLES_OPEN.
	ApplicationFlags_HandlesOpen ApplicationFlags = 4
	// ApplicationFlags_HandlesCommandLine is a representation of the C type G_APPLICATION_HANDLES_COMMAND_LINE.
	ApplicationFlags_HandlesCommandLine ApplicationFlags = 8
	// ApplicationFlags_SendEnvironment is a representation of the C type G_APPLICATION_SEND_ENVIRONMENT.
	ApplicationFlags_SendEnvironment ApplicationFlags = 16
	// ApplicationFlags_NonUnique is a representation of the C type G_APPLICATION_NON_UNIQUE.
	ApplicationFlags_NonUnique ApplicationFlags = 32
	// ApplicationFlags_CanOverrideAppId is a representation of the C type G_APPLICATION_CAN_OVERRIDE_APP_ID.
	ApplicationFlags_CanOverrideAppId ApplicationFlags = 64
	// ApplicationFlags_AllowReplacement is a representation of the C type G_APPLICATION_ALLOW_REPLACEMENT.
	ApplicationFlags_AllowReplacement ApplicationFlags = 128
	// ApplicationFlags_Replace is a representation of the C type G_APPLICATION_REPLACE.
	ApplicationFlags_Replace ApplicationFlags = 256
)

// AskPasswordFlags is a representation of the C type GAskPasswordFlags.
type AskPasswordFlags int

const (
	// AskPasswordFlags_NeedPassword is a representation of the C type G_ASK_PASSWORD_NEED_PASSWORD.
	AskPasswordFlags_NeedPassword AskPasswordFlags = 1
	// AskPasswordFlags_NeedUsername is a representation of the C type G_ASK_PASSWORD_NEED_USERNAME.
	AskPasswordFlags_NeedUsername AskPasswordFlags = 2
	// AskPasswordFlags_NeedDomain is a representation of the C type G_ASK_PASSWORD_NEED_DOMAIN.
	AskPasswordFlags_NeedDomain AskPasswordFlags = 4
	// AskPasswordFlags_SavingSupported is a representation of the C type G_ASK_PASSWORD_SAVING_SUPPORTED.
	AskPasswordFlags_SavingSupported AskPasswordFlags = 8
	// AskPasswordFlags_AnonymousSupported is a representation of the C type G_ASK_PASSWORD_ANONYMOUS_SUPPORTED.
	AskPasswordFlags_AnonymousSupported AskPasswordFlags = 16
	// AskPasswordFlags_Tcrypt is a representation of the C type G_ASK_PASSWORD_TCRYPT.
	AskPasswordFlags_Tcrypt AskPasswordFlags = 32
)

// BusNameOwnerFlags is a representation of the C type GBusNameOwnerFlags.
//
// since 2.26
type BusNameOwnerFlags int

const (
	// BusNameOwnerFlags_None is a representation of the C type G_BUS_NAME_OWNER_FLAGS_NONE.
	BusNameOwnerFlags_None BusNameOwnerFlags = 0
	// BusNameOwnerFlags_AllowReplacement is a representation of the C type G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT.
	BusNameOwnerFlags_AllowReplacement BusNameOwnerFlags = 1
	// BusNameOwnerFlags_Replace is a representation of the C type G_BUS_NAME_OWNER_FLAGS_REPLACE.
	BusNameOwnerFlags_Replace BusNameOwnerFlags = 2
	// BusNameOwnerFlags_DoNotQueue is a representation of the C type G_BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE.
	BusNameOwnerFlags_DoNotQueue BusNameOwnerFlags = 4
)

// BusNameWatcherFlags is a representation of the C type GBusNameWatcherFlags.
//
// since 2.26
type BusNameWatcherFlags int

const (
	// BusNameWatcherFlags_None is a representation of the C type G_BUS_NAME_WATCHER_FLAGS_NONE.
	BusNameWatcherFlags_None BusNameWatcherFlags = 0
	// BusNameWatcherFlags_AutoStart is a representation of the C type G_BUS_NAME_WATCHER_FLAGS_AUTO_START.
	BusNameWatcherFlags_AutoStart BusNameWatcherFlags = 1
)

// ConverterFlags is a representation of the C type GConverterFlags.
//
// since 2.24
type ConverterFlags int

const (
	// ConverterFlags_None is a representation of the C type G_CONVERTER_NO_FLAGS.
	ConverterFlags_None ConverterFlags = 0
	// ConverterFlags_InputAtEnd is a representation of the C type G_CONVERTER_INPUT_AT_END.
	ConverterFlags_InputAtEnd ConverterFlags = 1
	// ConverterFlags_Flush is a representation of the C type G_CONVERTER_FLUSH.
	ConverterFlags_Flush ConverterFlags = 2
)

// DBusCallFlags is a representation of the C type GDBusCallFlags.
//
// since 2.26
type DBusCallFlags int

const (
	// DBusCallFlags_None is a representation of the C type G_DBUS_CALL_FLAGS_NONE.
	DBusCallFlags_None DBusCallFlags = 0
	// DBusCallFlags_NoAutoStart is a representation of the C type G_DBUS_CALL_FLAGS_NO_AUTO_START.
	DBusCallFlags_NoAutoStart DBusCallFlags = 1
	// DBusCallFlags_AllowInteractiveAuthorization is a representation of the C type G_DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
	DBusCallFlags_AllowInteractiveAuthorization DBusCallFlags = 2
)

// DBusCapabilityFlags is a representation of the C type GDBusCapabilityFlags.
//
// since 2.26
type DBusCapabilityFlags int

const (
	// DBusCapabilityFlags_None is a representation of the C type G_DBUS_CAPABILITY_FLAGS_NONE.
	DBusCapabilityFlags_None DBusCapabilityFlags = 0
	// DBusCapabilityFlags_UnixFdPassing is a representation of the C type G_DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING.
	DBusCapabilityFlags_UnixFdPassing DBusCapabilityFlags = 1
)

// DBusConnectionFlags is a representation of the C type GDBusConnectionFlags.
//
// since 2.26
type DBusConnectionFlags int

const (
	// DBusConnectionFlags_None is a representation of the C type G_DBUS_CONNECTION_FLAGS_NONE.
	DBusConnectionFlags_None DBusConnectionFlags = 0
	// DBusConnectionFlags_AuthenticationClient is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT.
	DBusConnectionFlags_AuthenticationClient DBusConnectionFlags = 1
	// DBusConnectionFlags_AuthenticationServer is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER.
	DBusConnectionFlags_AuthenticationServer DBusConnectionFlags = 2
	// DBusConnectionFlags_AuthenticationAllowAnonymous is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
	DBusConnectionFlags_AuthenticationAllowAnonymous DBusConnectionFlags = 4
	// DBusConnectionFlags_MessageBusConnection is a representation of the C type G_DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION.
	DBusConnectionFlags_MessageBusConnection DBusConnectionFlags = 8
	// DBusConnectionFlags_DelayMessageProcessing is a representation of the C type G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING.
	DBusConnectionFlags_DelayMessageProcessing DBusConnectionFlags = 16
)

// DBusInterfaceSkeletonFlags is a representation of the C type GDBusInterfaceSkeletonFlags.
//
// since 2.30
type DBusInterfaceSkeletonFlags int

const (
	// DBusInterfaceSkeletonFlags_None is a representation of the C type G_DBUS_INTERFACE_SKELETON_FLAGS_NONE.
	DBusInterfaceSkeletonFlags_None DBusInterfaceSkeletonFlags = 0
	// DBusInterfaceSkeletonFlags_HandleMethodInvocationsInThread is a representation of the C type G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD.
	DBusInterfaceSkeletonFlags_HandleMethodInvocationsInThread DBusInterfaceSkeletonFlags = 1
)

// DBusMessageFlags is a representation of the C type GDBusMessageFlags.
//
// since 2.26
type DBusMessageFlags int

const (
	// DBusMessageFlags_None is a representation of the C type G_DBUS_MESSAGE_FLAGS_NONE.
	DBusMessageFlags_None DBusMessageFlags = 0
	// DBusMessageFlags_NoReplyExpected is a representation of the C type G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED.
	DBusMessageFlags_NoReplyExpected DBusMessageFlags = 1
	// DBusMessageFlags_NoAutoStart is a representation of the C type G_DBUS_MESSAGE_FLAGS_NO_AUTO_START.
	DBusMessageFlags_NoAutoStart DBusMessageFlags = 2
	// DBusMessageFlags_AllowInteractiveAuthorization is a representation of the C type G_DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
	DBusMessageFlags_AllowInteractiveAuthorization DBusMessageFlags = 4
)

// DBusObjectManagerClientFlags is a representation of the C type GDBusObjectManagerClientFlags.
//
// since 2.30
type DBusObjectManagerClientFlags int

const (
	// DBusObjectManagerClientFlags_None is a representation of the C type G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE.
	DBusObjectManagerClientFlags_None DBusObjectManagerClientFlags = 0
	// DBusObjectManagerClientFlags_DoNotAutoStart is a representation of the C type G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START.
	DBusObjectManagerClientFlags_DoNotAutoStart DBusObjectManagerClientFlags = 1
)

// DBusPropertyInfoFlags is a representation of the C type GDBusPropertyInfoFlags.
//
// since 2.26
type DBusPropertyInfoFlags int

const (
	// DBusPropertyInfoFlags_None is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_NONE.
	DBusPropertyInfoFlags_None DBusPropertyInfoFlags = 0
	// DBusPropertyInfoFlags_Readable is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_READABLE.
	DBusPropertyInfoFlags_Readable DBusPropertyInfoFlags = 1
	// DBusPropertyInfoFlags_Writable is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_WRITABLE.
	DBusPropertyInfoFlags_Writable DBusPropertyInfoFlags = 2
)

// DBusProxyFlags is a representation of the C type GDBusProxyFlags.
//
// since 2.26
type DBusProxyFlags int

const (
	// DBusProxyFlags_None is a representation of the C type G_DBUS_PROXY_FLAGS_NONE.
	DBusProxyFlags_None DBusProxyFlags = 0
	// DBusProxyFlags_DoNotLoadProperties is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES.
	DBusProxyFlags_DoNotLoadProperties DBusProxyFlags = 1
	// DBusProxyFlags_DoNotConnectSignals is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS.
	DBusProxyFlags_DoNotConnectSignals DBusProxyFlags = 2
	// DBusProxyFlags_DoNotAutoStart is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START.
	DBusProxyFlags_DoNotAutoStart DBusProxyFlags = 4
	// DBusProxyFlags_GetInvalidatedProperties is a representation of the C type G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES.
	DBusProxyFlags_GetInvalidatedProperties DBusProxyFlags = 8
	// DBusProxyFlags_DoNotAutoStartAtConstruction is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION.
	DBusProxyFlags_DoNotAutoStartAtConstruction DBusProxyFlags = 16
)

// DBusSendMessageFlags is a representation of the C type GDBusSendMessageFlags.
//
// since 2.26
type DBusSendMessageFlags int

const (
	// DBusSendMessageFlags_None is a representation of the C type G_DBUS_SEND_MESSAGE_FLAGS_NONE.
	DBusSendMessageFlags_None DBusSendMessageFlags = 0
	// DBusSendMessageFlags_PreserveSerial is a representation of the C type G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL.
	DBusSendMessageFlags_PreserveSerial DBusSendMessageFlags = 1
)

// DBusServerFlags is a representation of the C type GDBusServerFlags.
//
// since 2.26
type DBusServerFlags int

const (
	// DBusServerFlags_None is a representation of the C type G_DBUS_SERVER_FLAGS_NONE.
	DBusServerFlags_None DBusServerFlags = 0
	// DBusServerFlags_RunInThread is a representation of the C type G_DBUS_SERVER_FLAGS_RUN_IN_THREAD.
	DBusServerFlags_RunInThread DBusServerFlags = 1
	// DBusServerFlags_AuthenticationAllowAnonymous is a representation of the C type G_DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
	DBusServerFlags_AuthenticationAllowAnonymous DBusServerFlags = 2
)

// DBusSignalFlags is a representation of the C type GDBusSignalFlags.
//
// since 2.26
type DBusSignalFlags int

const (
	// DBusSignalFlags_None is a representation of the C type G_DBUS_SIGNAL_FLAGS_NONE.
	DBusSignalFlags_None DBusSignalFlags = 0
	// DBusSignalFlags_NoMatchRule is a representation of the C type G_DBUS_SIGNAL_FLAGS_NO_MATCH_RULE.
	DBusSignalFlags_NoMatchRule DBusSignalFlags = 1
	// DBusSignalFlags_MatchArg0Namespace is a representation of the C type G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE.
	DBusSignalFlags_MatchArg0Namespace DBusSignalFlags = 2
	// DBusSignalFlags_MatchArg0Path is a representation of the C type G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH.
	DBusSignalFlags_MatchArg0Path DBusSignalFlags = 4
)

// DBusSubtreeFlags is a representation of the C type GDBusSubtreeFlags.
//
// since 2.26
type DBusSubtreeFlags int

const (
	// DBusSubtreeFlags_None is a representation of the C type G_DBUS_SUBTREE_FLAGS_NONE.
	DBusSubtreeFlags_None DBusSubtreeFlags = 0
	// DBusSubtreeFlags_DispatchToUnenumeratedNodes is a representation of the C type G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES.
	DBusSubtreeFlags_DispatchToUnenumeratedNodes DBusSubtreeFlags = 1
)

// DriveStartFlags is a representation of the C type GDriveStartFlags.
//
// since 2.22
type DriveStartFlags int

const (
	// DriveStartFlags_None is a representation of the C type G_DRIVE_START_NONE.
	DriveStartFlags_None DriveStartFlags = 0
)

// FileAttributeInfoFlags is a representation of the C type GFileAttributeInfoFlags.
type FileAttributeInfoFlags int

const (
	// FileAttributeInfoFlags_None is a representation of the C type G_FILE_ATTRIBUTE_INFO_NONE.
	FileAttributeInfoFlags_None FileAttributeInfoFlags = 0
	// FileAttributeInfoFlags_CopyWithFile is a representation of the C type G_FILE_ATTRIBUTE_INFO_COPY_WITH_FILE.
	FileAttributeInfoFlags_CopyWithFile FileAttributeInfoFlags = 1
	// FileAttributeInfoFlags_CopyWhenMoved is a representation of the C type G_FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED.
	FileAttributeInfoFlags_CopyWhenMoved FileAttributeInfoFlags = 2
)

// FileCopyFlags is a representation of the C type GFileCopyFlags.
type FileCopyFlags int

const (
	// FileCopyFlags_None is a representation of the C type G_FILE_COPY_NONE.
	FileCopyFlags_None FileCopyFlags = 0
	// FileCopyFlags_Overwrite is a representation of the C type G_FILE_COPY_OVERWRITE.
	FileCopyFlags_Overwrite FileCopyFlags = 1
	// FileCopyFlags_Backup is a representation of the C type G_FILE_COPY_BACKUP.
	FileCopyFlags_Backup FileCopyFlags = 2
	// FileCopyFlags_NofollowSymlinks is a representation of the C type G_FILE_COPY_NOFOLLOW_SYMLINKS.
	FileCopyFlags_NofollowSymlinks FileCopyFlags = 4
	// FileCopyFlags_AllMetadata is a representation of the C type G_FILE_COPY_ALL_METADATA.
	FileCopyFlags_AllMetadata FileCopyFlags = 8
	// FileCopyFlags_NoFallbackForMove is a representation of the C type G_FILE_COPY_NO_FALLBACK_FOR_MOVE.
	FileCopyFlags_NoFallbackForMove FileCopyFlags = 16
	// FileCopyFlags_TargetDefaultPerms is a representation of the C type G_FILE_COPY_TARGET_DEFAULT_PERMS.
	FileCopyFlags_TargetDefaultPerms FileCopyFlags = 32
)

// FileCreateFlags is a representation of the C type GFileCreateFlags.
type FileCreateFlags int

const (
	// FileCreateFlags_None is a representation of the C type G_FILE_CREATE_NONE.
	FileCreateFlags_None FileCreateFlags = 0
	// FileCreateFlags_Private is a representation of the C type G_FILE_CREATE_PRIVATE.
	FileCreateFlags_Private FileCreateFlags = 1
	// FileCreateFlags_ReplaceDestination is a representation of the C type G_FILE_CREATE_REPLACE_DESTINATION.
	FileCreateFlags_ReplaceDestination FileCreateFlags = 2
)

// FileMeasureFlags is a representation of the C type GFileMeasureFlags.
//
// since 2.38
type FileMeasureFlags int

const (
	// FileMeasureFlags_None is a representation of the C type G_FILE_MEASURE_NONE.
	FileMeasureFlags_None FileMeasureFlags = 0
	// FileMeasureFlags_ReportAnyError is a representation of the C type G_FILE_MEASURE_REPORT_ANY_ERROR.
	FileMeasureFlags_ReportAnyError FileMeasureFlags = 2
	// FileMeasureFlags_ApparentSize is a representation of the C type G_FILE_MEASURE_APPARENT_SIZE.
	FileMeasureFlags_ApparentSize FileMeasureFlags = 4
	// FileMeasureFlags_NoXdev is a representation of the C type G_FILE_MEASURE_NO_XDEV.
	FileMeasureFlags_NoXdev FileMeasureFlags = 8
)

// FileMonitorFlags is a representation of the C type GFileMonitorFlags.
type FileMonitorFlags int

const (
	// FileMonitorFlags_None is a representation of the C type G_FILE_MONITOR_NONE.
	FileMonitorFlags_None FileMonitorFlags = 0
	// FileMonitorFlags_WatchMounts is a representation of the C type G_FILE_MONITOR_WATCH_MOUNTS.
	FileMonitorFlags_WatchMounts FileMonitorFlags = 1
	// FileMonitorFlags_SendMoved is a representation of the C type G_FILE_MONITOR_SEND_MOVED.
	FileMonitorFlags_SendMoved FileMonitorFlags = 2
	// FileMonitorFlags_WatchHardLinks is a representation of the C type G_FILE_MONITOR_WATCH_HARD_LINKS.
	FileMonitorFlags_WatchHardLinks FileMonitorFlags = 4
	// FileMonitorFlags_WatchMoves is a representation of the C type G_FILE_MONITOR_WATCH_MOVES.
	FileMonitorFlags_WatchMoves FileMonitorFlags = 8
)

// FileQueryInfoFlags is a representation of the C type GFileQueryInfoFlags.
type FileQueryInfoFlags int

const (
	// FileQueryInfoFlags_None is a representation of the C type G_FILE_QUERY_INFO_NONE.
	FileQueryInfoFlags_None FileQueryInfoFlags = 0
	// FileQueryInfoFlags_NofollowSymlinks is a representation of the C type G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS.
	FileQueryInfoFlags_NofollowSymlinks FileQueryInfoFlags = 1
)

// IOStreamSpliceFlags is a representation of the C type GIOStreamSpliceFlags.
//
// since 2.28
type IOStreamSpliceFlags int

const (
	// IOStreamSpliceFlags_None is a representation of the C type G_IO_STREAM_SPLICE_NONE.
	IOStreamSpliceFlags_None IOStreamSpliceFlags = 0
	// IOStreamSpliceFlags_CloseStream1 is a representation of the C type G_IO_STREAM_SPLICE_CLOSE_STREAM1.
	IOStreamSpliceFlags_CloseStream1 IOStreamSpliceFlags = 1
	// IOStreamSpliceFlags_CloseStream2 is a representation of the C type G_IO_STREAM_SPLICE_CLOSE_STREAM2.
	IOStreamSpliceFlags_CloseStream2 IOStreamSpliceFlags = 2
	// IOStreamSpliceFlags_WaitForBoth is a representation of the C type G_IO_STREAM_SPLICE_WAIT_FOR_BOTH.
	IOStreamSpliceFlags_WaitForBoth IOStreamSpliceFlags = 4
)

// MountMountFlags is a representation of the C type GMountMountFlags.
type MountMountFlags int

const (
	// MountMountFlags_None is a representation of the C type G_MOUNT_MOUNT_NONE.
	MountMountFlags_None MountMountFlags = 0
)

// MountUnmountFlags is a representation of the C type GMountUnmountFlags.
type MountUnmountFlags int

const (
	// MountUnmountFlags_None is a representation of the C type G_MOUNT_UNMOUNT_NONE.
	MountUnmountFlags_None MountUnmountFlags = 0
	// MountUnmountFlags_Force is a representation of the C type G_MOUNT_UNMOUNT_FORCE.
	MountUnmountFlags_Force MountUnmountFlags = 1
)

// OutputStreamSpliceFlags is a representation of the C type GOutputStreamSpliceFlags.
type OutputStreamSpliceFlags int

const (
	// OutputStreamSpliceFlags_None is a representation of the C type G_OUTPUT_STREAM_SPLICE_NONE.
	OutputStreamSpliceFlags_None OutputStreamSpliceFlags = 0
	// OutputStreamSpliceFlags_CloseSource is a representation of the C type G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE.
	OutputStreamSpliceFlags_CloseSource OutputStreamSpliceFlags = 1
	// OutputStreamSpliceFlags_CloseTarget is a representation of the C type G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET.
	OutputStreamSpliceFlags_CloseTarget OutputStreamSpliceFlags = 2
)

// ResolverNameLookupFlags is a representation of the C type GResolverNameLookupFlags.
//
// since 2.60
type ResolverNameLookupFlags int

const (
	// ResolverNameLookupFlags_Default is a representation of the C type G_RESOLVER_NAME_LOOKUP_FLAGS_DEFAULT.
	ResolverNameLookupFlags_Default ResolverNameLookupFlags = 0
	// ResolverNameLookupFlags_Ipv4Only is a representation of the C type G_RESOLVER_NAME_LOOKUP_FLAGS_IPV4_ONLY.
	ResolverNameLookupFlags_Ipv4Only ResolverNameLookupFlags = 1
	// ResolverNameLookupFlags_Ipv6Only is a representation of the C type G_RESOLVER_NAME_LOOKUP_FLAGS_IPV6_ONLY.
	ResolverNameLookupFlags_Ipv6Only ResolverNameLookupFlags = 2
)

// ResourceFlags is a representation of the C type GResourceFlags.
//
// since 2.32
type ResourceFlags int

const (
	// ResourceFlags_None is a representation of the C type G_RESOURCE_FLAGS_NONE.
	ResourceFlags_None ResourceFlags = 0
	// ResourceFlags_Compressed is a representation of the C type G_RESOURCE_FLAGS_COMPRESSED.
	ResourceFlags_Compressed ResourceFlags = 1
)

// ResourceLookupFlags is a representation of the C type GResourceLookupFlags.
//
// since 2.32
type ResourceLookupFlags int

const (
	// ResourceLookupFlags_None is a representation of the C type G_RESOURCE_LOOKUP_FLAGS_NONE.
	ResourceLookupFlags_None ResourceLookupFlags = 0
)

// SettingsBindFlags is a representation of the C type GSettingsBindFlags.
type SettingsBindFlags int

const (
	// SettingsBindFlags_Default is a representation of the C type G_SETTINGS_BIND_DEFAULT.
	SettingsBindFlags_Default SettingsBindFlags = 0
	// SettingsBindFlags_Get is a representation of the C type G_SETTINGS_BIND_GET.
	SettingsBindFlags_Get SettingsBindFlags = 1
	// SettingsBindFlags_Set is a representation of the C type G_SETTINGS_BIND_SET.
	SettingsBindFlags_Set SettingsBindFlags = 2
	// SettingsBindFlags_NoSensitivity is a representation of the C type G_SETTINGS_BIND_NO_SENSITIVITY.
	SettingsBindFlags_NoSensitivity SettingsBindFlags = 4
	// SettingsBindFlags_GetNoChanges is a representation of the C type G_SETTINGS_BIND_GET_NO_CHANGES.
	SettingsBindFlags_GetNoChanges SettingsBindFlags = 8
	// SettingsBindFlags_InvertBoolean is a representation of the C type G_SETTINGS_BIND_INVERT_BOOLEAN.
	SettingsBindFlags_InvertBoolean SettingsBindFlags = 16
)

// SocketMsgFlags is a representation of the C type GSocketMsgFlags.
//
// since 2.22
type SocketMsgFlags int

const (
	// SocketMsgFlags_None is a representation of the C type G_SOCKET_MSG_NONE.
	SocketMsgFlags_None SocketMsgFlags = 0
	// SocketMsgFlags_Oob is a representation of the C type G_SOCKET_MSG_OOB.
	SocketMsgFlags_Oob SocketMsgFlags = 1
	// SocketMsgFlags_Peek is a representation of the C type G_SOCKET_MSG_PEEK.
	SocketMsgFlags_Peek SocketMsgFlags = 2
	// SocketMsgFlags_Dontroute is a representation of the C type G_SOCKET_MSG_DONTROUTE.
	SocketMsgFlags_Dontroute SocketMsgFlags = 4
)

// SubprocessFlags is a representation of the C type GSubprocessFlags.
//
// since 2.40
type SubprocessFlags int

const (
	// SubprocessFlags_None is a representation of the C type G_SUBPROCESS_FLAGS_NONE.
	SubprocessFlags_None SubprocessFlags = 0
	// SubprocessFlags_StdinPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDIN_PIPE.
	SubprocessFlags_StdinPipe SubprocessFlags = 1
	// SubprocessFlags_StdinInherit is a representation of the C type G_SUBPROCESS_FLAGS_STDIN_INHERIT.
	SubprocessFlags_StdinInherit SubprocessFlags = 2
	// SubprocessFlags_StdoutPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDOUT_PIPE.
	SubprocessFlags_StdoutPipe SubprocessFlags = 4
	// SubprocessFlags_StdoutSilence is a representation of the C type G_SUBPROCESS_FLAGS_STDOUT_SILENCE.
	SubprocessFlags_StdoutSilence SubprocessFlags = 8
	// SubprocessFlags_StderrPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_PIPE.
	SubprocessFlags_StderrPipe SubprocessFlags = 16
	// SubprocessFlags_StderrSilence is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_SILENCE.
	SubprocessFlags_StderrSilence SubprocessFlags = 32
	// SubprocessFlags_StderrMerge is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_MERGE.
	SubprocessFlags_StderrMerge SubprocessFlags = 64
	// SubprocessFlags_InheritFds is a representation of the C type G_SUBPROCESS_FLAGS_INHERIT_FDS.
	SubprocessFlags_InheritFds SubprocessFlags = 128
)

// TestDBusFlags is a representation of the C type GTestDBusFlags.
//
// since 2.34
type TestDBusFlags int

const (
	// TestDBusFlags_None is a representation of the C type G_TEST_DBUS_NONE.
	TestDBusFlags_None TestDBusFlags = 0
)

// TlsCertificateFlags is a representation of the C type GTlsCertificateFlags.
//
// since 2.28
type TlsCertificateFlags int

const (
	// TlsCertificateFlags_UnknownCa is a representation of the C type G_TLS_CERTIFICATE_UNKNOWN_CA.
	TlsCertificateFlags_UnknownCa TlsCertificateFlags = 1
	// TlsCertificateFlags_BadIdentity is a representation of the C type G_TLS_CERTIFICATE_BAD_IDENTITY.
	TlsCertificateFlags_BadIdentity TlsCertificateFlags = 2
	// TlsCertificateFlags_NotActivated is a representation of the C type G_TLS_CERTIFICATE_NOT_ACTIVATED.
	TlsCertificateFlags_NotActivated TlsCertificateFlags = 4
	// TlsCertificateFlags_Expired is a representation of the C type G_TLS_CERTIFICATE_EXPIRED.
	TlsCertificateFlags_Expired TlsCertificateFlags = 8
	// TlsCertificateFlags_Revoked is a representation of the C type G_TLS_CERTIFICATE_REVOKED.
	TlsCertificateFlags_Revoked TlsCertificateFlags = 16
	// TlsCertificateFlags_Insecure is a representation of the C type G_TLS_CERTIFICATE_INSECURE.
	TlsCertificateFlags_Insecure TlsCertificateFlags = 32
	// TlsCertificateFlags_GenericError is a representation of the C type G_TLS_CERTIFICATE_GENERIC_ERROR.
	TlsCertificateFlags_GenericError TlsCertificateFlags = 64
	// TlsCertificateFlags_ValidateAll is a representation of the C type G_TLS_CERTIFICATE_VALIDATE_ALL.
	TlsCertificateFlags_ValidateAll TlsCertificateFlags = 127
)

// TlsDatabaseVerifyFlags is a representation of the C type GTlsDatabaseVerifyFlags.
//
// since 2.30
type TlsDatabaseVerifyFlags int

const (
	// TlsDatabaseVerifyFlags_None is a representation of the C type G_TLS_DATABASE_VERIFY_NONE.
	TlsDatabaseVerifyFlags_None TlsDatabaseVerifyFlags = 0
)

// TlsPasswordFlags is a representation of the C type GTlsPasswordFlags.
//
// since 2.30
type TlsPasswordFlags int

const (
	// TlsPasswordFlags_None is a representation of the C type G_TLS_PASSWORD_NONE.
	TlsPasswordFlags_None TlsPasswordFlags = 0
	// TlsPasswordFlags_Retry is a representation of the C type G_TLS_PASSWORD_RETRY.
	TlsPasswordFlags_Retry TlsPasswordFlags = 2
	// TlsPasswordFlags_ManyTries is a representation of the C type G_TLS_PASSWORD_MANY_TRIES.
	TlsPasswordFlags_ManyTries TlsPasswordFlags = 4
	// TlsPasswordFlags_FinalTry is a representation of the C type G_TLS_PASSWORD_FINAL_TRY.
	TlsPasswordFlags_FinalTry TlsPasswordFlags = 8
)
