// Code generated - DO NOT EDIT.

package gio

// Appinfocreateflags is a representation of the C type GAppInfoCreateFlags.
type Appinfocreateflags int

const (
	// Appinfocreateflags_None is a representation of the C type G_APP_INFO_CREATE_NONE.
	Appinfocreateflags_None Appinfocreateflags = 0
	// Appinfocreateflags_NeedsTerminal is a representation of the C type G_APP_INFO_CREATE_NEEDS_TERMINAL.
	Appinfocreateflags_NeedsTerminal Appinfocreateflags = 1
	// Appinfocreateflags_SupportsUris is a representation of the C type G_APP_INFO_CREATE_SUPPORTS_URIS.
	Appinfocreateflags_SupportsUris Appinfocreateflags = 2
	// Appinfocreateflags_SupportsStartupNotification is a representation of the C type G_APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION.
	Appinfocreateflags_SupportsStartupNotification Appinfocreateflags = 4
)

// Applicationflags is a representation of the C type GApplicationFlags.
//
// since 2.28
type Applicationflags int

const (
	// Applicationflags_FlagsNone is a representation of the C type G_APPLICATION_FLAGS_NONE.
	Applicationflags_FlagsNone Applicationflags = 0
	// Applicationflags_IsService is a representation of the C type G_APPLICATION_IS_SERVICE.
	Applicationflags_IsService Applicationflags = 1
	// Applicationflags_IsLauncher is a representation of the C type G_APPLICATION_IS_LAUNCHER.
	Applicationflags_IsLauncher Applicationflags = 2
	// Applicationflags_HandlesOpen is a representation of the C type G_APPLICATION_HANDLES_OPEN.
	Applicationflags_HandlesOpen Applicationflags = 4
	// Applicationflags_HandlesCommandLine is a representation of the C type G_APPLICATION_HANDLES_COMMAND_LINE.
	Applicationflags_HandlesCommandLine Applicationflags = 8
	// Applicationflags_SendEnvironment is a representation of the C type G_APPLICATION_SEND_ENVIRONMENT.
	Applicationflags_SendEnvironment Applicationflags = 16
	// Applicationflags_NonUnique is a representation of the C type G_APPLICATION_NON_UNIQUE.
	Applicationflags_NonUnique Applicationflags = 32
	// Applicationflags_CanOverrideAppId is a representation of the C type G_APPLICATION_CAN_OVERRIDE_APP_ID.
	Applicationflags_CanOverrideAppId Applicationflags = 64
)

// Askpasswordflags is a representation of the C type GAskPasswordFlags.
type Askpasswordflags int

const (
	// Askpasswordflags_NeedPassword is a representation of the C type G_ASK_PASSWORD_NEED_PASSWORD.
	Askpasswordflags_NeedPassword Askpasswordflags = 1
	// Askpasswordflags_NeedUsername is a representation of the C type G_ASK_PASSWORD_NEED_USERNAME.
	Askpasswordflags_NeedUsername Askpasswordflags = 2
	// Askpasswordflags_NeedDomain is a representation of the C type G_ASK_PASSWORD_NEED_DOMAIN.
	Askpasswordflags_NeedDomain Askpasswordflags = 4
	// Askpasswordflags_SavingSupported is a representation of the C type G_ASK_PASSWORD_SAVING_SUPPORTED.
	Askpasswordflags_SavingSupported Askpasswordflags = 8
	// Askpasswordflags_AnonymousSupported is a representation of the C type G_ASK_PASSWORD_ANONYMOUS_SUPPORTED.
	Askpasswordflags_AnonymousSupported Askpasswordflags = 16
)

// Busnameownerflags is a representation of the C type GBusNameOwnerFlags.
//
// since 2.26
type Busnameownerflags int

const (
	// Busnameownerflags_None is a representation of the C type G_BUS_NAME_OWNER_FLAGS_NONE.
	Busnameownerflags_None Busnameownerflags = 0
	// Busnameownerflags_AllowReplacement is a representation of the C type G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT.
	Busnameownerflags_AllowReplacement Busnameownerflags = 1
	// Busnameownerflags_Replace is a representation of the C type G_BUS_NAME_OWNER_FLAGS_REPLACE.
	Busnameownerflags_Replace Busnameownerflags = 2
	// Busnameownerflags_DoNotQueue is a representation of the C type G_BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE.
	Busnameownerflags_DoNotQueue Busnameownerflags = 4
)

// Busnamewatcherflags is a representation of the C type GBusNameWatcherFlags.
//
// since 2.26
type Busnamewatcherflags int

const (
	// Busnamewatcherflags_None is a representation of the C type G_BUS_NAME_WATCHER_FLAGS_NONE.
	Busnamewatcherflags_None Busnamewatcherflags = 0
	// Busnamewatcherflags_AutoStart is a representation of the C type G_BUS_NAME_WATCHER_FLAGS_AUTO_START.
	Busnamewatcherflags_AutoStart Busnamewatcherflags = 1
)

// Converterflags is a representation of the C type GConverterFlags.
//
// since 2.24
type Converterflags int

const (
	// Converterflags_None is a representation of the C type G_CONVERTER_NO_FLAGS.
	Converterflags_None Converterflags = 0
	// Converterflags_InputAtEnd is a representation of the C type G_CONVERTER_INPUT_AT_END.
	Converterflags_InputAtEnd Converterflags = 1
	// Converterflags_Flush is a representation of the C type G_CONVERTER_FLUSH.
	Converterflags_Flush Converterflags = 2
)

// Dbuscallflags is a representation of the C type GDBusCallFlags.
//
// since 2.26
type Dbuscallflags int

const (
	// Dbuscallflags_None is a representation of the C type G_DBUS_CALL_FLAGS_NONE.
	Dbuscallflags_None Dbuscallflags = 0
	// Dbuscallflags_NoAutoStart is a representation of the C type G_DBUS_CALL_FLAGS_NO_AUTO_START.
	Dbuscallflags_NoAutoStart Dbuscallflags = 1
	// Dbuscallflags_AllowInteractiveAuthorization is a representation of the C type G_DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
	Dbuscallflags_AllowInteractiveAuthorization Dbuscallflags = 2
)

// Dbuscapabilityflags is a representation of the C type GDBusCapabilityFlags.
//
// since 2.26
type Dbuscapabilityflags int

const (
	// Dbuscapabilityflags_None is a representation of the C type G_DBUS_CAPABILITY_FLAGS_NONE.
	Dbuscapabilityflags_None Dbuscapabilityflags = 0
	// Dbuscapabilityflags_UnixFdPassing is a representation of the C type G_DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING.
	Dbuscapabilityflags_UnixFdPassing Dbuscapabilityflags = 1
)

// Dbusconnectionflags is a representation of the C type GDBusConnectionFlags.
//
// since 2.26
type Dbusconnectionflags int

const (
	// Dbusconnectionflags_None is a representation of the C type G_DBUS_CONNECTION_FLAGS_NONE.
	Dbusconnectionflags_None Dbusconnectionflags = 0
	// Dbusconnectionflags_AuthenticationClient is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT.
	Dbusconnectionflags_AuthenticationClient Dbusconnectionflags = 1
	// Dbusconnectionflags_AuthenticationServer is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER.
	Dbusconnectionflags_AuthenticationServer Dbusconnectionflags = 2
	// Dbusconnectionflags_AuthenticationAllowAnonymous is a representation of the C type G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
	Dbusconnectionflags_AuthenticationAllowAnonymous Dbusconnectionflags = 4
	// Dbusconnectionflags_MessageBusConnection is a representation of the C type G_DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION.
	Dbusconnectionflags_MessageBusConnection Dbusconnectionflags = 8
	// Dbusconnectionflags_DelayMessageProcessing is a representation of the C type G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING.
	Dbusconnectionflags_DelayMessageProcessing Dbusconnectionflags = 16
)

// Dbusinterfaceskeletonflags is a representation of the C type GDBusInterfaceSkeletonFlags.
//
// since 2.30
type Dbusinterfaceskeletonflags int

const (
	// Dbusinterfaceskeletonflags_None is a representation of the C type G_DBUS_INTERFACE_SKELETON_FLAGS_NONE.
	Dbusinterfaceskeletonflags_None Dbusinterfaceskeletonflags = 0
	// Dbusinterfaceskeletonflags_HandleMethodInvocationsInThread is a representation of the C type G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD.
	Dbusinterfaceskeletonflags_HandleMethodInvocationsInThread Dbusinterfaceskeletonflags = 1
)

// Dbusmessageflags is a representation of the C type GDBusMessageFlags.
//
// since 2.26
type Dbusmessageflags int

const (
	// Dbusmessageflags_None is a representation of the C type G_DBUS_MESSAGE_FLAGS_NONE.
	Dbusmessageflags_None Dbusmessageflags = 0
	// Dbusmessageflags_NoReplyExpected is a representation of the C type G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED.
	Dbusmessageflags_NoReplyExpected Dbusmessageflags = 1
	// Dbusmessageflags_NoAutoStart is a representation of the C type G_DBUS_MESSAGE_FLAGS_NO_AUTO_START.
	Dbusmessageflags_NoAutoStart Dbusmessageflags = 2
	// Dbusmessageflags_AllowInteractiveAuthorization is a representation of the C type G_DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
	Dbusmessageflags_AllowInteractiveAuthorization Dbusmessageflags = 4
)

// Dbusobjectmanagerclientflags is a representation of the C type GDBusObjectManagerClientFlags.
//
// since 2.30
type Dbusobjectmanagerclientflags int

const (
	// Dbusobjectmanagerclientflags_None is a representation of the C type G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE.
	Dbusobjectmanagerclientflags_None Dbusobjectmanagerclientflags = 0
	// Dbusobjectmanagerclientflags_DoNotAutoStart is a representation of the C type G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START.
	Dbusobjectmanagerclientflags_DoNotAutoStart Dbusobjectmanagerclientflags = 1
)

// Dbuspropertyinfoflags is a representation of the C type GDBusPropertyInfoFlags.
//
// since 2.26
type Dbuspropertyinfoflags int

const (
	// Dbuspropertyinfoflags_None is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_NONE.
	Dbuspropertyinfoflags_None Dbuspropertyinfoflags = 0
	// Dbuspropertyinfoflags_Readable is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_READABLE.
	Dbuspropertyinfoflags_Readable Dbuspropertyinfoflags = 1
	// Dbuspropertyinfoflags_Writable is a representation of the C type G_DBUS_PROPERTY_INFO_FLAGS_WRITABLE.
	Dbuspropertyinfoflags_Writable Dbuspropertyinfoflags = 2
)

// Dbusproxyflags is a representation of the C type GDBusProxyFlags.
//
// since 2.26
type Dbusproxyflags int

const (
	// Dbusproxyflags_None is a representation of the C type G_DBUS_PROXY_FLAGS_NONE.
	Dbusproxyflags_None Dbusproxyflags = 0
	// Dbusproxyflags_DoNotLoadProperties is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES.
	Dbusproxyflags_DoNotLoadProperties Dbusproxyflags = 1
	// Dbusproxyflags_DoNotConnectSignals is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS.
	Dbusproxyflags_DoNotConnectSignals Dbusproxyflags = 2
	// Dbusproxyflags_DoNotAutoStart is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START.
	Dbusproxyflags_DoNotAutoStart Dbusproxyflags = 4
	// Dbusproxyflags_GetInvalidatedProperties is a representation of the C type G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES.
	Dbusproxyflags_GetInvalidatedProperties Dbusproxyflags = 8
	// Dbusproxyflags_DoNotAutoStartAtConstruction is a representation of the C type G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION.
	Dbusproxyflags_DoNotAutoStartAtConstruction Dbusproxyflags = 16
)

// Dbussendmessageflags is a representation of the C type GDBusSendMessageFlags.
//
// since 2.26
type Dbussendmessageflags int

const (
	// Dbussendmessageflags_None is a representation of the C type G_DBUS_SEND_MESSAGE_FLAGS_NONE.
	Dbussendmessageflags_None Dbussendmessageflags = 0
	// Dbussendmessageflags_PreserveSerial is a representation of the C type G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL.
	Dbussendmessageflags_PreserveSerial Dbussendmessageflags = 1
)

// Dbusserverflags is a representation of the C type GDBusServerFlags.
//
// since 2.26
type Dbusserverflags int

const (
	// Dbusserverflags_None is a representation of the C type G_DBUS_SERVER_FLAGS_NONE.
	Dbusserverflags_None Dbusserverflags = 0
	// Dbusserverflags_RunInThread is a representation of the C type G_DBUS_SERVER_FLAGS_RUN_IN_THREAD.
	Dbusserverflags_RunInThread Dbusserverflags = 1
	// Dbusserverflags_AuthenticationAllowAnonymous is a representation of the C type G_DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
	Dbusserverflags_AuthenticationAllowAnonymous Dbusserverflags = 2
)

// Dbussignalflags is a representation of the C type GDBusSignalFlags.
//
// since 2.26
type Dbussignalflags int

const (
	// Dbussignalflags_None is a representation of the C type G_DBUS_SIGNAL_FLAGS_NONE.
	Dbussignalflags_None Dbussignalflags = 0
	// Dbussignalflags_NoMatchRule is a representation of the C type G_DBUS_SIGNAL_FLAGS_NO_MATCH_RULE.
	Dbussignalflags_NoMatchRule Dbussignalflags = 1
	// Dbussignalflags_MatchArg0Namespace is a representation of the C type G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE.
	Dbussignalflags_MatchArg0Namespace Dbussignalflags = 2
	// Dbussignalflags_MatchArg0Path is a representation of the C type G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH.
	Dbussignalflags_MatchArg0Path Dbussignalflags = 4
)

// Dbussubtreeflags is a representation of the C type GDBusSubtreeFlags.
//
// since 2.26
type Dbussubtreeflags int

const (
	// Dbussubtreeflags_None is a representation of the C type G_DBUS_SUBTREE_FLAGS_NONE.
	Dbussubtreeflags_None Dbussubtreeflags = 0
	// Dbussubtreeflags_DispatchToUnenumeratedNodes is a representation of the C type G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES.
	Dbussubtreeflags_DispatchToUnenumeratedNodes Dbussubtreeflags = 1
)

// Drivestartflags is a representation of the C type GDriveStartFlags.
//
// since 2.22
type Drivestartflags int

const (
	// Drivestartflags_None is a representation of the C type G_DRIVE_START_NONE.
	Drivestartflags_None Drivestartflags = 0
)

// Fileattributeinfoflags is a representation of the C type GFileAttributeInfoFlags.
type Fileattributeinfoflags int

const (
	// Fileattributeinfoflags_None is a representation of the C type G_FILE_ATTRIBUTE_INFO_NONE.
	Fileattributeinfoflags_None Fileattributeinfoflags = 0
	// Fileattributeinfoflags_CopyWithFile is a representation of the C type G_FILE_ATTRIBUTE_INFO_COPY_WITH_FILE.
	Fileattributeinfoflags_CopyWithFile Fileattributeinfoflags = 1
	// Fileattributeinfoflags_CopyWhenMoved is a representation of the C type G_FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED.
	Fileattributeinfoflags_CopyWhenMoved Fileattributeinfoflags = 2
)

// Filecopyflags is a representation of the C type GFileCopyFlags.
type Filecopyflags int

const (
	// Filecopyflags_None is a representation of the C type G_FILE_COPY_NONE.
	Filecopyflags_None Filecopyflags = 0
	// Filecopyflags_Overwrite is a representation of the C type G_FILE_COPY_OVERWRITE.
	Filecopyflags_Overwrite Filecopyflags = 1
	// Filecopyflags_Backup is a representation of the C type G_FILE_COPY_BACKUP.
	Filecopyflags_Backup Filecopyflags = 2
	// Filecopyflags_NofollowSymlinks is a representation of the C type G_FILE_COPY_NOFOLLOW_SYMLINKS.
	Filecopyflags_NofollowSymlinks Filecopyflags = 4
	// Filecopyflags_AllMetadata is a representation of the C type G_FILE_COPY_ALL_METADATA.
	Filecopyflags_AllMetadata Filecopyflags = 8
	// Filecopyflags_NoFallbackForMove is a representation of the C type G_FILE_COPY_NO_FALLBACK_FOR_MOVE.
	Filecopyflags_NoFallbackForMove Filecopyflags = 16
	// Filecopyflags_TargetDefaultPerms is a representation of the C type G_FILE_COPY_TARGET_DEFAULT_PERMS.
	Filecopyflags_TargetDefaultPerms Filecopyflags = 32
)

// Filecreateflags is a representation of the C type GFileCreateFlags.
type Filecreateflags int

const (
	// Filecreateflags_None is a representation of the C type G_FILE_CREATE_NONE.
	Filecreateflags_None Filecreateflags = 0
	// Filecreateflags_Private is a representation of the C type G_FILE_CREATE_PRIVATE.
	Filecreateflags_Private Filecreateflags = 1
	// Filecreateflags_ReplaceDestination is a representation of the C type G_FILE_CREATE_REPLACE_DESTINATION.
	Filecreateflags_ReplaceDestination Filecreateflags = 2
)

// Filemeasureflags is a representation of the C type GFileMeasureFlags.
//
// since 2.38
type Filemeasureflags int

const (
	// Filemeasureflags_None is a representation of the C type G_FILE_MEASURE_NONE.
	Filemeasureflags_None Filemeasureflags = 0
	// Filemeasureflags_ReportAnyError is a representation of the C type G_FILE_MEASURE_REPORT_ANY_ERROR.
	Filemeasureflags_ReportAnyError Filemeasureflags = 2
	// Filemeasureflags_ApparentSize is a representation of the C type G_FILE_MEASURE_APPARENT_SIZE.
	Filemeasureflags_ApparentSize Filemeasureflags = 4
	// Filemeasureflags_NoXdev is a representation of the C type G_FILE_MEASURE_NO_XDEV.
	Filemeasureflags_NoXdev Filemeasureflags = 8
)

// Filemonitorflags is a representation of the C type GFileMonitorFlags.
type Filemonitorflags int

const (
	// Filemonitorflags_None is a representation of the C type G_FILE_MONITOR_NONE.
	Filemonitorflags_None Filemonitorflags = 0
	// Filemonitorflags_WatchMounts is a representation of the C type G_FILE_MONITOR_WATCH_MOUNTS.
	Filemonitorflags_WatchMounts Filemonitorflags = 1
	// Filemonitorflags_SendMoved is a representation of the C type G_FILE_MONITOR_SEND_MOVED.
	Filemonitorflags_SendMoved Filemonitorflags = 2
	// Filemonitorflags_WatchHardLinks is a representation of the C type G_FILE_MONITOR_WATCH_HARD_LINKS.
	Filemonitorflags_WatchHardLinks Filemonitorflags = 4
	// Filemonitorflags_WatchMoves is a representation of the C type G_FILE_MONITOR_WATCH_MOVES.
	Filemonitorflags_WatchMoves Filemonitorflags = 8
)

// Filequeryinfoflags is a representation of the C type GFileQueryInfoFlags.
type Filequeryinfoflags int

const (
	// Filequeryinfoflags_None is a representation of the C type G_FILE_QUERY_INFO_NONE.
	Filequeryinfoflags_None Filequeryinfoflags = 0
	// Filequeryinfoflags_NofollowSymlinks is a representation of the C type G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS.
	Filequeryinfoflags_NofollowSymlinks Filequeryinfoflags = 1
)

// Iostreamspliceflags is a representation of the C type GIOStreamSpliceFlags.
//
// since 2.28
type Iostreamspliceflags int

const (
	// Iostreamspliceflags_None is a representation of the C type G_IO_STREAM_SPLICE_NONE.
	Iostreamspliceflags_None Iostreamspliceflags = 0
	// Iostreamspliceflags_CloseStream1 is a representation of the C type G_IO_STREAM_SPLICE_CLOSE_STREAM1.
	Iostreamspliceflags_CloseStream1 Iostreamspliceflags = 1
	// Iostreamspliceflags_CloseStream2 is a representation of the C type G_IO_STREAM_SPLICE_CLOSE_STREAM2.
	Iostreamspliceflags_CloseStream2 Iostreamspliceflags = 2
	// Iostreamspliceflags_WaitForBoth is a representation of the C type G_IO_STREAM_SPLICE_WAIT_FOR_BOTH.
	Iostreamspliceflags_WaitForBoth Iostreamspliceflags = 4
)

// Mountmountflags is a representation of the C type GMountMountFlags.
type Mountmountflags int

const (
	// Mountmountflags_None is a representation of the C type G_MOUNT_MOUNT_NONE.
	Mountmountflags_None Mountmountflags = 0
)

// Mountunmountflags is a representation of the C type GMountUnmountFlags.
type Mountunmountflags int

const (
	// Mountunmountflags_None is a representation of the C type G_MOUNT_UNMOUNT_NONE.
	Mountunmountflags_None Mountunmountflags = 0
	// Mountunmountflags_Force is a representation of the C type G_MOUNT_UNMOUNT_FORCE.
	Mountunmountflags_Force Mountunmountflags = 1
)

// Outputstreamspliceflags is a representation of the C type GOutputStreamSpliceFlags.
type Outputstreamspliceflags int

const (
	// Outputstreamspliceflags_None is a representation of the C type G_OUTPUT_STREAM_SPLICE_NONE.
	Outputstreamspliceflags_None Outputstreamspliceflags = 0
	// Outputstreamspliceflags_CloseSource is a representation of the C type G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE.
	Outputstreamspliceflags_CloseSource Outputstreamspliceflags = 1
	// Outputstreamspliceflags_CloseTarget is a representation of the C type G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET.
	Outputstreamspliceflags_CloseTarget Outputstreamspliceflags = 2
)

// Resourceflags is a representation of the C type GResourceFlags.
//
// since 2.32
type Resourceflags int

const (
	// Resourceflags_None is a representation of the C type G_RESOURCE_FLAGS_NONE.
	Resourceflags_None Resourceflags = 0
	// Resourceflags_Compressed is a representation of the C type G_RESOURCE_FLAGS_COMPRESSED.
	Resourceflags_Compressed Resourceflags = 1
)

// Resourcelookupflags is a representation of the C type GResourceLookupFlags.
//
// since 2.32
type Resourcelookupflags int

const (
	// Resourcelookupflags_None is a representation of the C type G_RESOURCE_LOOKUP_FLAGS_NONE.
	Resourcelookupflags_None Resourcelookupflags = 0
)

// Settingsbindflags is a representation of the C type GSettingsBindFlags.
type Settingsbindflags int

const (
	// Settingsbindflags_Default is a representation of the C type G_SETTINGS_BIND_DEFAULT.
	Settingsbindflags_Default Settingsbindflags = 0
	// Settingsbindflags_Get is a representation of the C type G_SETTINGS_BIND_GET.
	Settingsbindflags_Get Settingsbindflags = 1
	// Settingsbindflags_Set is a representation of the C type G_SETTINGS_BIND_SET.
	Settingsbindflags_Set Settingsbindflags = 2
	// Settingsbindflags_NoSensitivity is a representation of the C type G_SETTINGS_BIND_NO_SENSITIVITY.
	Settingsbindflags_NoSensitivity Settingsbindflags = 4
	// Settingsbindflags_GetNoChanges is a representation of the C type G_SETTINGS_BIND_GET_NO_CHANGES.
	Settingsbindflags_GetNoChanges Settingsbindflags = 8
	// Settingsbindflags_InvertBoolean is a representation of the C type G_SETTINGS_BIND_INVERT_BOOLEAN.
	Settingsbindflags_InvertBoolean Settingsbindflags = 16
)

// Socketmsgflags is a representation of the C type GSocketMsgFlags.
//
// since 2.22
type Socketmsgflags int

const (
	// Socketmsgflags_None is a representation of the C type G_SOCKET_MSG_NONE.
	Socketmsgflags_None Socketmsgflags = 0
	// Socketmsgflags_Oob is a representation of the C type G_SOCKET_MSG_OOB.
	Socketmsgflags_Oob Socketmsgflags = 1
	// Socketmsgflags_Peek is a representation of the C type G_SOCKET_MSG_PEEK.
	Socketmsgflags_Peek Socketmsgflags = 2
	// Socketmsgflags_Dontroute is a representation of the C type G_SOCKET_MSG_DONTROUTE.
	Socketmsgflags_Dontroute Socketmsgflags = 4
)

// Subprocessflags is a representation of the C type GSubprocessFlags.
//
// since 2.40
type Subprocessflags int

const (
	// Subprocessflags_None is a representation of the C type G_SUBPROCESS_FLAGS_NONE.
	Subprocessflags_None Subprocessflags = 0
	// Subprocessflags_StdinPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDIN_PIPE.
	Subprocessflags_StdinPipe Subprocessflags = 1
	// Subprocessflags_StdinInherit is a representation of the C type G_SUBPROCESS_FLAGS_STDIN_INHERIT.
	Subprocessflags_StdinInherit Subprocessflags = 2
	// Subprocessflags_StdoutPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDOUT_PIPE.
	Subprocessflags_StdoutPipe Subprocessflags = 4
	// Subprocessflags_StdoutSilence is a representation of the C type G_SUBPROCESS_FLAGS_STDOUT_SILENCE.
	Subprocessflags_StdoutSilence Subprocessflags = 8
	// Subprocessflags_StderrPipe is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_PIPE.
	Subprocessflags_StderrPipe Subprocessflags = 16
	// Subprocessflags_StderrSilence is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_SILENCE.
	Subprocessflags_StderrSilence Subprocessflags = 32
	// Subprocessflags_StderrMerge is a representation of the C type G_SUBPROCESS_FLAGS_STDERR_MERGE.
	Subprocessflags_StderrMerge Subprocessflags = 64
	// Subprocessflags_InheritFds is a representation of the C type G_SUBPROCESS_FLAGS_INHERIT_FDS.
	Subprocessflags_InheritFds Subprocessflags = 128
)

// Testdbusflags is a representation of the C type GTestDBusFlags.
//
// since 2.34
type Testdbusflags int

const (
	// Testdbusflags_None is a representation of the C type G_TEST_DBUS_NONE.
	Testdbusflags_None Testdbusflags = 0
)

// Tlscertificateflags is a representation of the C type GTlsCertificateFlags.
//
// since 2.28
type Tlscertificateflags int

const (
	// Tlscertificateflags_UnknownCa is a representation of the C type G_TLS_CERTIFICATE_UNKNOWN_CA.
	Tlscertificateflags_UnknownCa Tlscertificateflags = 1
	// Tlscertificateflags_BadIdentity is a representation of the C type G_TLS_CERTIFICATE_BAD_IDENTITY.
	Tlscertificateflags_BadIdentity Tlscertificateflags = 2
	// Tlscertificateflags_NotActivated is a representation of the C type G_TLS_CERTIFICATE_NOT_ACTIVATED.
	Tlscertificateflags_NotActivated Tlscertificateflags = 4
	// Tlscertificateflags_Expired is a representation of the C type G_TLS_CERTIFICATE_EXPIRED.
	Tlscertificateflags_Expired Tlscertificateflags = 8
	// Tlscertificateflags_Revoked is a representation of the C type G_TLS_CERTIFICATE_REVOKED.
	Tlscertificateflags_Revoked Tlscertificateflags = 16
	// Tlscertificateflags_Insecure is a representation of the C type G_TLS_CERTIFICATE_INSECURE.
	Tlscertificateflags_Insecure Tlscertificateflags = 32
	// Tlscertificateflags_GenericError is a representation of the C type G_TLS_CERTIFICATE_GENERIC_ERROR.
	Tlscertificateflags_GenericError Tlscertificateflags = 64
	// Tlscertificateflags_ValidateAll is a representation of the C type G_TLS_CERTIFICATE_VALIDATE_ALL.
	Tlscertificateflags_ValidateAll Tlscertificateflags = 127
)

// Tlsdatabaseverifyflags is a representation of the C type GTlsDatabaseVerifyFlags.
//
// since 2.30
type Tlsdatabaseverifyflags int

const (
	// Tlsdatabaseverifyflags_None is a representation of the C type G_TLS_DATABASE_VERIFY_NONE.
	Tlsdatabaseverifyflags_None Tlsdatabaseverifyflags = 0
)

// Tlspasswordflags is a representation of the C type GTlsPasswordFlags.
//
// since 2.30
type Tlspasswordflags int

const (
	// Tlspasswordflags_None is a representation of the C type G_TLS_PASSWORD_NONE.
	Tlspasswordflags_None Tlspasswordflags = 0
	// Tlspasswordflags_Retry is a representation of the C type G_TLS_PASSWORD_RETRY.
	Tlspasswordflags_Retry Tlspasswordflags = 2
	// Tlspasswordflags_ManyTries is a representation of the C type G_TLS_PASSWORD_MANY_TRIES.
	Tlspasswordflags_ManyTries Tlspasswordflags = 4
	// Tlspasswordflags_FinalTry is a representation of the C type G_TLS_PASSWORD_FINAL_TRY.
	Tlspasswordflags_FinalTry Tlspasswordflags = 8
)
