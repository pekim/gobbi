// Code generated - DO NOT EDIT.

package gio

// BusType is a representation of the C type GBusType.
//
// since 2.26
type BusType int

const (
	// BusType_Starter is a representation of the C type G_BUS_TYPE_STARTER.
	BusType_Starter BusType = -1
	// BusType_None is a representation of the C type G_BUS_TYPE_NONE.
	BusType_None BusType = 0
	// BusType_System is a representation of the C type G_BUS_TYPE_SYSTEM.
	BusType_System BusType = 1
	// BusType_Session is a representation of the C type G_BUS_TYPE_SESSION.
	BusType_Session BusType = 2
)

// ConverterResult is a representation of the C type GConverterResult.
//
// since 2.24
type ConverterResult int

const (
	// ConverterResult_Error is a representation of the C type G_CONVERTER_ERROR.
	ConverterResult_Error ConverterResult = 0
	// ConverterResult_Converted is a representation of the C type G_CONVERTER_CONVERTED.
	ConverterResult_Converted ConverterResult = 1
	// ConverterResult_Finished is a representation of the C type G_CONVERTER_FINISHED.
	ConverterResult_Finished ConverterResult = 2
	// ConverterResult_Flushed is a representation of the C type G_CONVERTER_FLUSHED.
	ConverterResult_Flushed ConverterResult = 3
)

// CredentialsType is a representation of the C type GCredentialsType.
//
// since 2.26
type CredentialsType int

const (
	// CredentialsType_Invalid is a representation of the C type G_CREDENTIALS_TYPE_INVALID.
	CredentialsType_Invalid CredentialsType = 0
	// CredentialsType_LinuxUcred is a representation of the C type G_CREDENTIALS_TYPE_LINUX_UCRED.
	CredentialsType_LinuxUcred CredentialsType = 1
	// CredentialsType_FreebsdCmsgcred is a representation of the C type G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
	CredentialsType_FreebsdCmsgcred CredentialsType = 2
	// CredentialsType_OpenbsdSockpeercred is a representation of the C type G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
	CredentialsType_OpenbsdSockpeercred CredentialsType = 3
	// CredentialsType_SolarisUcred is a representation of the C type G_CREDENTIALS_TYPE_SOLARIS_UCRED.
	CredentialsType_SolarisUcred CredentialsType = 4
	// CredentialsType_NetbsdUnpcbid is a representation of the C type G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
	CredentialsType_NetbsdUnpcbid CredentialsType = 5
)

// DBusError is a representation of the C type GDBusError.
//
// since 2.26
type DBusError int

const (
	// DBusError_Failed is a representation of the C type G_DBUS_ERROR_FAILED.
	DBusError_Failed DBusError = 0
	// DBusError_NoMemory is a representation of the C type G_DBUS_ERROR_NO_MEMORY.
	DBusError_NoMemory DBusError = 1
	// DBusError_ServiceUnknown is a representation of the C type G_DBUS_ERROR_SERVICE_UNKNOWN.
	DBusError_ServiceUnknown DBusError = 2
	// DBusError_NameHasNoOwner is a representation of the C type G_DBUS_ERROR_NAME_HAS_NO_OWNER.
	DBusError_NameHasNoOwner DBusError = 3
	// DBusError_NoReply is a representation of the C type G_DBUS_ERROR_NO_REPLY.
	DBusError_NoReply DBusError = 4
	// DBusError_IoError is a representation of the C type G_DBUS_ERROR_IO_ERROR.
	DBusError_IoError DBusError = 5
	// DBusError_BadAddress is a representation of the C type G_DBUS_ERROR_BAD_ADDRESS.
	DBusError_BadAddress DBusError = 6
	// DBusError_NotSupported is a representation of the C type G_DBUS_ERROR_NOT_SUPPORTED.
	DBusError_NotSupported DBusError = 7
	// DBusError_LimitsExceeded is a representation of the C type G_DBUS_ERROR_LIMITS_EXCEEDED.
	DBusError_LimitsExceeded DBusError = 8
	// DBusError_AccessDenied is a representation of the C type G_DBUS_ERROR_ACCESS_DENIED.
	DBusError_AccessDenied DBusError = 9
	// DBusError_AuthFailed is a representation of the C type G_DBUS_ERROR_AUTH_FAILED.
	DBusError_AuthFailed DBusError = 10
	// DBusError_NoServer is a representation of the C type G_DBUS_ERROR_NO_SERVER.
	DBusError_NoServer DBusError = 11
	// DBusError_Timeout is a representation of the C type G_DBUS_ERROR_TIMEOUT.
	DBusError_Timeout DBusError = 12
	// DBusError_NoNetwork is a representation of the C type G_DBUS_ERROR_NO_NETWORK.
	DBusError_NoNetwork DBusError = 13
	// DBusError_AddressInUse is a representation of the C type G_DBUS_ERROR_ADDRESS_IN_USE.
	DBusError_AddressInUse DBusError = 14
	// DBusError_Disconnected is a representation of the C type G_DBUS_ERROR_DISCONNECTED.
	DBusError_Disconnected DBusError = 15
	// DBusError_InvalidArgs is a representation of the C type G_DBUS_ERROR_INVALID_ARGS.
	DBusError_InvalidArgs DBusError = 16
	// DBusError_FileNotFound is a representation of the C type G_DBUS_ERROR_FILE_NOT_FOUND.
	DBusError_FileNotFound DBusError = 17
	// DBusError_FileExists is a representation of the C type G_DBUS_ERROR_FILE_EXISTS.
	DBusError_FileExists DBusError = 18
	// DBusError_UnknownMethod is a representation of the C type G_DBUS_ERROR_UNKNOWN_METHOD.
	DBusError_UnknownMethod DBusError = 19
	// DBusError_TimedOut is a representation of the C type G_DBUS_ERROR_TIMED_OUT.
	DBusError_TimedOut DBusError = 20
	// DBusError_MatchRuleNotFound is a representation of the C type G_DBUS_ERROR_MATCH_RULE_NOT_FOUND.
	DBusError_MatchRuleNotFound DBusError = 21
	// DBusError_MatchRuleInvalid is a representation of the C type G_DBUS_ERROR_MATCH_RULE_INVALID.
	DBusError_MatchRuleInvalid DBusError = 22
	// DBusError_SpawnExecFailed is a representation of the C type G_DBUS_ERROR_SPAWN_EXEC_FAILED.
	DBusError_SpawnExecFailed DBusError = 23
	// DBusError_SpawnForkFailed is a representation of the C type G_DBUS_ERROR_SPAWN_FORK_FAILED.
	DBusError_SpawnForkFailed DBusError = 24
	// DBusError_SpawnChildExited is a representation of the C type G_DBUS_ERROR_SPAWN_CHILD_EXITED.
	DBusError_SpawnChildExited DBusError = 25
	// DBusError_SpawnChildSignaled is a representation of the C type G_DBUS_ERROR_SPAWN_CHILD_SIGNALED.
	DBusError_SpawnChildSignaled DBusError = 26
	// DBusError_SpawnFailed is a representation of the C type G_DBUS_ERROR_SPAWN_FAILED.
	DBusError_SpawnFailed DBusError = 27
	// DBusError_SpawnSetupFailed is a representation of the C type G_DBUS_ERROR_SPAWN_SETUP_FAILED.
	DBusError_SpawnSetupFailed DBusError = 28
	// DBusError_SpawnConfigInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_CONFIG_INVALID.
	DBusError_SpawnConfigInvalid DBusError = 29
	// DBusError_SpawnServiceInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_SERVICE_INVALID.
	DBusError_SpawnServiceInvalid DBusError = 30
	// DBusError_SpawnServiceNotFound is a representation of the C type G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND.
	DBusError_SpawnServiceNotFound DBusError = 31
	// DBusError_SpawnPermissionsInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID.
	DBusError_SpawnPermissionsInvalid DBusError = 32
	// DBusError_SpawnFileInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_FILE_INVALID.
	DBusError_SpawnFileInvalid DBusError = 33
	// DBusError_SpawnNoMemory is a representation of the C type G_DBUS_ERROR_SPAWN_NO_MEMORY.
	DBusError_SpawnNoMemory DBusError = 34
	// DBusError_UnixProcessIdUnknown is a representation of the C type G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN.
	DBusError_UnixProcessIdUnknown DBusError = 35
	// DBusError_InvalidSignature is a representation of the C type G_DBUS_ERROR_INVALID_SIGNATURE.
	DBusError_InvalidSignature DBusError = 36
	// DBusError_InvalidFileContent is a representation of the C type G_DBUS_ERROR_INVALID_FILE_CONTENT.
	DBusError_InvalidFileContent DBusError = 37
	// DBusError_SelinuxSecurityContextUnknown is a representation of the C type G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN.
	DBusError_SelinuxSecurityContextUnknown DBusError = 38
	// DBusError_AdtAuditDataUnknown is a representation of the C type G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN.
	DBusError_AdtAuditDataUnknown DBusError = 39
	// DBusError_ObjectPathInUse is a representation of the C type G_DBUS_ERROR_OBJECT_PATH_IN_USE.
	DBusError_ObjectPathInUse DBusError = 40
	// DBusError_UnknownObject is a representation of the C type G_DBUS_ERROR_UNKNOWN_OBJECT.
	DBusError_UnknownObject DBusError = 41
	// DBusError_UnknownInterface is a representation of the C type G_DBUS_ERROR_UNKNOWN_INTERFACE.
	DBusError_UnknownInterface DBusError = 42
	// DBusError_UnknownProperty is a representation of the C type G_DBUS_ERROR_UNKNOWN_PROPERTY.
	DBusError_UnknownProperty DBusError = 43
	// DBusError_PropertyReadOnly is a representation of the C type G_DBUS_ERROR_PROPERTY_READ_ONLY.
	DBusError_PropertyReadOnly DBusError = 44
)

// DBusMessageByteOrder is a representation of the C type GDBusMessageByteOrder.
//
// since 2.26
type DBusMessageByteOrder int

const (
	// DBusMessageByteOrder_BigEndian is a representation of the C type G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN.
	DBusMessageByteOrder_BigEndian DBusMessageByteOrder = 66
	// DBusMessageByteOrder_LittleEndian is a representation of the C type G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN.
	DBusMessageByteOrder_LittleEndian DBusMessageByteOrder = 108
)

// DBusMessageHeaderField is a representation of the C type GDBusMessageHeaderField.
//
// since 2.26
type DBusMessageHeaderField int

const (
	// DBusMessageHeaderField_Invalid is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_INVALID.
	DBusMessageHeaderField_Invalid DBusMessageHeaderField = 0
	// DBusMessageHeaderField_Path is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_PATH.
	DBusMessageHeaderField_Path DBusMessageHeaderField = 1
	// DBusMessageHeaderField_Interface is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE.
	DBusMessageHeaderField_Interface DBusMessageHeaderField = 2
	// DBusMessageHeaderField_Member is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_MEMBER.
	DBusMessageHeaderField_Member DBusMessageHeaderField = 3
	// DBusMessageHeaderField_ErrorName is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME.
	DBusMessageHeaderField_ErrorName DBusMessageHeaderField = 4
	// DBusMessageHeaderField_ReplySerial is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL.
	DBusMessageHeaderField_ReplySerial DBusMessageHeaderField = 5
	// DBusMessageHeaderField_Destination is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION.
	DBusMessageHeaderField_Destination DBusMessageHeaderField = 6
	// DBusMessageHeaderField_Sender is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_SENDER.
	DBusMessageHeaderField_Sender DBusMessageHeaderField = 7
	// DBusMessageHeaderField_Signature is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE.
	DBusMessageHeaderField_Signature DBusMessageHeaderField = 8
	// DBusMessageHeaderField_NumUnixFds is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS.
	DBusMessageHeaderField_NumUnixFds DBusMessageHeaderField = 9
)

// DBusMessageType is a representation of the C type GDBusMessageType.
//
// since 2.26
type DBusMessageType int

const (
	// DBusMessageType_Invalid is a representation of the C type G_DBUS_MESSAGE_TYPE_INVALID.
	DBusMessageType_Invalid DBusMessageType = 0
	// DBusMessageType_MethodCall is a representation of the C type G_DBUS_MESSAGE_TYPE_METHOD_CALL.
	DBusMessageType_MethodCall DBusMessageType = 1
	// DBusMessageType_MethodReturn is a representation of the C type G_DBUS_MESSAGE_TYPE_METHOD_RETURN.
	DBusMessageType_MethodReturn DBusMessageType = 2
	// DBusMessageType_Error is a representation of the C type G_DBUS_MESSAGE_TYPE_ERROR.
	DBusMessageType_Error DBusMessageType = 3
	// DBusMessageType_Signal is a representation of the C type G_DBUS_MESSAGE_TYPE_SIGNAL.
	DBusMessageType_Signal DBusMessageType = 4
)

// DataStreamByteOrder is a representation of the C type GDataStreamByteOrder.
//
type DataStreamByteOrder int

const (
	// DataStreamByteOrder_BigEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN.
	DataStreamByteOrder_BigEndian DataStreamByteOrder = 0
	// DataStreamByteOrder_LittleEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN.
	DataStreamByteOrder_LittleEndian DataStreamByteOrder = 1
	// DataStreamByteOrder_HostEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN.
	DataStreamByteOrder_HostEndian DataStreamByteOrder = 2
)

// DataStreamNewlineType is a representation of the C type GDataStreamNewlineType.
//
type DataStreamNewlineType int

const (
	// DataStreamNewlineType_Lf is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_LF.
	DataStreamNewlineType_Lf DataStreamNewlineType = 0
	// DataStreamNewlineType_Cr is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_CR.
	DataStreamNewlineType_Cr DataStreamNewlineType = 1
	// DataStreamNewlineType_CrLf is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_CR_LF.
	DataStreamNewlineType_CrLf DataStreamNewlineType = 2
	// DataStreamNewlineType_Any is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_ANY.
	DataStreamNewlineType_Any DataStreamNewlineType = 3
)

// DriveStartStopType is a representation of the C type GDriveStartStopType.
//
// since 2.22
type DriveStartStopType int

const (
	// DriveStartStopType_Unknown is a representation of the C type G_DRIVE_START_STOP_TYPE_UNKNOWN.
	DriveStartStopType_Unknown DriveStartStopType = 0
	// DriveStartStopType_Shutdown is a representation of the C type G_DRIVE_START_STOP_TYPE_SHUTDOWN.
	DriveStartStopType_Shutdown DriveStartStopType = 1
	// DriveStartStopType_Network is a representation of the C type G_DRIVE_START_STOP_TYPE_NETWORK.
	DriveStartStopType_Network DriveStartStopType = 2
	// DriveStartStopType_Multidisk is a representation of the C type G_DRIVE_START_STOP_TYPE_MULTIDISK.
	DriveStartStopType_Multidisk DriveStartStopType = 3
	// DriveStartStopType_Password is a representation of the C type G_DRIVE_START_STOP_TYPE_PASSWORD.
	DriveStartStopType_Password DriveStartStopType = 4
)

// EmblemOrigin is a representation of the C type GEmblemOrigin.
//
// since 2.18
type EmblemOrigin int

const (
	// EmblemOrigin_Unknown is a representation of the C type G_EMBLEM_ORIGIN_UNKNOWN.
	EmblemOrigin_Unknown EmblemOrigin = 0
	// EmblemOrigin_Device is a representation of the C type G_EMBLEM_ORIGIN_DEVICE.
	EmblemOrigin_Device EmblemOrigin = 1
	// EmblemOrigin_Livemetadata is a representation of the C type G_EMBLEM_ORIGIN_LIVEMETADATA.
	EmblemOrigin_Livemetadata EmblemOrigin = 2
	// EmblemOrigin_Tag is a representation of the C type G_EMBLEM_ORIGIN_TAG.
	EmblemOrigin_Tag EmblemOrigin = 3
)

// FileAttributeStatus is a representation of the C type GFileAttributeStatus.
//
type FileAttributeStatus int

const (
	// FileAttributeStatus_Unset is a representation of the C type G_FILE_ATTRIBUTE_STATUS_UNSET.
	FileAttributeStatus_Unset FileAttributeStatus = 0
	// FileAttributeStatus_Set is a representation of the C type G_FILE_ATTRIBUTE_STATUS_SET.
	FileAttributeStatus_Set FileAttributeStatus = 1
	// FileAttributeStatus_ErrorSetting is a representation of the C type G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING.
	FileAttributeStatus_ErrorSetting FileAttributeStatus = 2
)

// FileAttributeType is a representation of the C type GFileAttributeType.
//
type FileAttributeType int

const (
	// FileAttributeType_Invalid is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INVALID.
	FileAttributeType_Invalid FileAttributeType = 0
	// FileAttributeType_String is a representation of the C type G_FILE_ATTRIBUTE_TYPE_STRING.
	FileAttributeType_String FileAttributeType = 1
	// FileAttributeType_ByteString is a representation of the C type G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
	FileAttributeType_ByteString FileAttributeType = 2
	// FileAttributeType_Boolean is a representation of the C type G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
	FileAttributeType_Boolean FileAttributeType = 3
	// FileAttributeType_Uint32 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_UINT32.
	FileAttributeType_Uint32 FileAttributeType = 4
	// FileAttributeType_Int32 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INT32.
	FileAttributeType_Int32 FileAttributeType = 5
	// FileAttributeType_Uint64 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_UINT64.
	FileAttributeType_Uint64 FileAttributeType = 6
	// FileAttributeType_Int64 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INT64.
	FileAttributeType_Int64 FileAttributeType = 7
	// FileAttributeType_Object is a representation of the C type G_FILE_ATTRIBUTE_TYPE_OBJECT.
	FileAttributeType_Object FileAttributeType = 8
	// FileAttributeType_Stringv is a representation of the C type G_FILE_ATTRIBUTE_TYPE_STRINGV.
	FileAttributeType_Stringv FileAttributeType = 9
)

// FileMonitorEvent is a representation of the C type GFileMonitorEvent.
//
type FileMonitorEvent int

const (
	// FileMonitorEvent_Changed is a representation of the C type G_FILE_MONITOR_EVENT_CHANGED.
	FileMonitorEvent_Changed FileMonitorEvent = 0
	// FileMonitorEvent_ChangesDoneHint is a representation of the C type G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT.
	FileMonitorEvent_ChangesDoneHint FileMonitorEvent = 1
	// FileMonitorEvent_Deleted is a representation of the C type G_FILE_MONITOR_EVENT_DELETED.
	FileMonitorEvent_Deleted FileMonitorEvent = 2
	// FileMonitorEvent_Created is a representation of the C type G_FILE_MONITOR_EVENT_CREATED.
	FileMonitorEvent_Created FileMonitorEvent = 3
	// FileMonitorEvent_AttributeChanged is a representation of the C type G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED.
	FileMonitorEvent_AttributeChanged FileMonitorEvent = 4
	// FileMonitorEvent_PreUnmount is a representation of the C type G_FILE_MONITOR_EVENT_PRE_UNMOUNT.
	FileMonitorEvent_PreUnmount FileMonitorEvent = 5
	// FileMonitorEvent_Unmounted is a representation of the C type G_FILE_MONITOR_EVENT_UNMOUNTED.
	FileMonitorEvent_Unmounted FileMonitorEvent = 6
	// FileMonitorEvent_Moved is a representation of the C type G_FILE_MONITOR_EVENT_MOVED.
	FileMonitorEvent_Moved FileMonitorEvent = 7
	// FileMonitorEvent_Renamed is a representation of the C type G_FILE_MONITOR_EVENT_RENAMED.
	FileMonitorEvent_Renamed FileMonitorEvent = 8
	// FileMonitorEvent_MovedIn is a representation of the C type G_FILE_MONITOR_EVENT_MOVED_IN.
	FileMonitorEvent_MovedIn FileMonitorEvent = 9
	// FileMonitorEvent_MovedOut is a representation of the C type G_FILE_MONITOR_EVENT_MOVED_OUT.
	FileMonitorEvent_MovedOut FileMonitorEvent = 10
)

// FileType is a representation of the C type GFileType.
//
type FileType int

const (
	// FileType_Unknown is a representation of the C type G_FILE_TYPE_UNKNOWN.
	FileType_Unknown FileType = 0
	// FileType_Regular is a representation of the C type G_FILE_TYPE_REGULAR.
	FileType_Regular FileType = 1
	// FileType_Directory is a representation of the C type G_FILE_TYPE_DIRECTORY.
	FileType_Directory FileType = 2
	// FileType_SymbolicLink is a representation of the C type G_FILE_TYPE_SYMBOLIC_LINK.
	FileType_SymbolicLink FileType = 3
	// FileType_Special is a representation of the C type G_FILE_TYPE_SPECIAL.
	FileType_Special FileType = 4
	// FileType_Shortcut is a representation of the C type G_FILE_TYPE_SHORTCUT.
	FileType_Shortcut FileType = 5
	// FileType_Mountable is a representation of the C type G_FILE_TYPE_MOUNTABLE.
	FileType_Mountable FileType = 6
)

// FilesystemPreviewType is a representation of the C type GFilesystemPreviewType.
//
type FilesystemPreviewType int

const (
	// FilesystemPreviewType_IfAlways is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS.
	FilesystemPreviewType_IfAlways FilesystemPreviewType = 0
	// FilesystemPreviewType_IfLocal is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL.
	FilesystemPreviewType_IfLocal FilesystemPreviewType = 1
	// FilesystemPreviewType_Never is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_NEVER.
	FilesystemPreviewType_Never FilesystemPreviewType = 2
)

// IOErrorEnum is a representation of the C type GIOErrorEnum.
//
type IOErrorEnum int

const (
	// IOErrorEnum_Failed is a representation of the C type G_IO_ERROR_FAILED.
	IOErrorEnum_Failed IOErrorEnum = 0
	// IOErrorEnum_NotFound is a representation of the C type G_IO_ERROR_NOT_FOUND.
	IOErrorEnum_NotFound IOErrorEnum = 1
	// IOErrorEnum_Exists is a representation of the C type G_IO_ERROR_EXISTS.
	IOErrorEnum_Exists IOErrorEnum = 2
	// IOErrorEnum_IsDirectory is a representation of the C type G_IO_ERROR_IS_DIRECTORY.
	IOErrorEnum_IsDirectory IOErrorEnum = 3
	// IOErrorEnum_NotDirectory is a representation of the C type G_IO_ERROR_NOT_DIRECTORY.
	IOErrorEnum_NotDirectory IOErrorEnum = 4
	// IOErrorEnum_NotEmpty is a representation of the C type G_IO_ERROR_NOT_EMPTY.
	IOErrorEnum_NotEmpty IOErrorEnum = 5
	// IOErrorEnum_NotRegularFile is a representation of the C type G_IO_ERROR_NOT_REGULAR_FILE.
	IOErrorEnum_NotRegularFile IOErrorEnum = 6
	// IOErrorEnum_NotSymbolicLink is a representation of the C type G_IO_ERROR_NOT_SYMBOLIC_LINK.
	IOErrorEnum_NotSymbolicLink IOErrorEnum = 7
	// IOErrorEnum_NotMountableFile is a representation of the C type G_IO_ERROR_NOT_MOUNTABLE_FILE.
	IOErrorEnum_NotMountableFile IOErrorEnum = 8
	// IOErrorEnum_FilenameTooLong is a representation of the C type G_IO_ERROR_FILENAME_TOO_LONG.
	IOErrorEnum_FilenameTooLong IOErrorEnum = 9
	// IOErrorEnum_InvalidFilename is a representation of the C type G_IO_ERROR_INVALID_FILENAME.
	IOErrorEnum_InvalidFilename IOErrorEnum = 10
	// IOErrorEnum_TooManyLinks is a representation of the C type G_IO_ERROR_TOO_MANY_LINKS.
	IOErrorEnum_TooManyLinks IOErrorEnum = 11
	// IOErrorEnum_NoSpace is a representation of the C type G_IO_ERROR_NO_SPACE.
	IOErrorEnum_NoSpace IOErrorEnum = 12
	// IOErrorEnum_InvalidArgument is a representation of the C type G_IO_ERROR_INVALID_ARGUMENT.
	IOErrorEnum_InvalidArgument IOErrorEnum = 13
	// IOErrorEnum_PermissionDenied is a representation of the C type G_IO_ERROR_PERMISSION_DENIED.
	IOErrorEnum_PermissionDenied IOErrorEnum = 14
	// IOErrorEnum_NotSupported is a representation of the C type G_IO_ERROR_NOT_SUPPORTED.
	IOErrorEnum_NotSupported IOErrorEnum = 15
	// IOErrorEnum_NotMounted is a representation of the C type G_IO_ERROR_NOT_MOUNTED.
	IOErrorEnum_NotMounted IOErrorEnum = 16
	// IOErrorEnum_AlreadyMounted is a representation of the C type G_IO_ERROR_ALREADY_MOUNTED.
	IOErrorEnum_AlreadyMounted IOErrorEnum = 17
	// IOErrorEnum_Closed is a representation of the C type G_IO_ERROR_CLOSED.
	IOErrorEnum_Closed IOErrorEnum = 18
	// IOErrorEnum_Cancelled is a representation of the C type G_IO_ERROR_CANCELLED.
	IOErrorEnum_Cancelled IOErrorEnum = 19
	// IOErrorEnum_Pending is a representation of the C type G_IO_ERROR_PENDING.
	IOErrorEnum_Pending IOErrorEnum = 20
	// IOErrorEnum_ReadOnly is a representation of the C type G_IO_ERROR_READ_ONLY.
	IOErrorEnum_ReadOnly IOErrorEnum = 21
	// IOErrorEnum_CantCreateBackup is a representation of the C type G_IO_ERROR_CANT_CREATE_BACKUP.
	IOErrorEnum_CantCreateBackup IOErrorEnum = 22
	// IOErrorEnum_WrongEtag is a representation of the C type G_IO_ERROR_WRONG_ETAG.
	IOErrorEnum_WrongEtag IOErrorEnum = 23
	// IOErrorEnum_TimedOut is a representation of the C type G_IO_ERROR_TIMED_OUT.
	IOErrorEnum_TimedOut IOErrorEnum = 24
	// IOErrorEnum_WouldRecurse is a representation of the C type G_IO_ERROR_WOULD_RECURSE.
	IOErrorEnum_WouldRecurse IOErrorEnum = 25
	// IOErrorEnum_Busy is a representation of the C type G_IO_ERROR_BUSY.
	IOErrorEnum_Busy IOErrorEnum = 26
	// IOErrorEnum_WouldBlock is a representation of the C type G_IO_ERROR_WOULD_BLOCK.
	IOErrorEnum_WouldBlock IOErrorEnum = 27
	// IOErrorEnum_HostNotFound is a representation of the C type G_IO_ERROR_HOST_NOT_FOUND.
	IOErrorEnum_HostNotFound IOErrorEnum = 28
	// IOErrorEnum_WouldMerge is a representation of the C type G_IO_ERROR_WOULD_MERGE.
	IOErrorEnum_WouldMerge IOErrorEnum = 29
	// IOErrorEnum_FailedHandled is a representation of the C type G_IO_ERROR_FAILED_HANDLED.
	IOErrorEnum_FailedHandled IOErrorEnum = 30
	// IOErrorEnum_TooManyOpenFiles is a representation of the C type G_IO_ERROR_TOO_MANY_OPEN_FILES.
	IOErrorEnum_TooManyOpenFiles IOErrorEnum = 31
	// IOErrorEnum_NotInitialized is a representation of the C type G_IO_ERROR_NOT_INITIALIZED.
	IOErrorEnum_NotInitialized IOErrorEnum = 32
	// IOErrorEnum_AddressInUse is a representation of the C type G_IO_ERROR_ADDRESS_IN_USE.
	IOErrorEnum_AddressInUse IOErrorEnum = 33
	// IOErrorEnum_PartialInput is a representation of the C type G_IO_ERROR_PARTIAL_INPUT.
	IOErrorEnum_PartialInput IOErrorEnum = 34
	// IOErrorEnum_InvalidData is a representation of the C type G_IO_ERROR_INVALID_DATA.
	IOErrorEnum_InvalidData IOErrorEnum = 35
	// IOErrorEnum_DbusError is a representation of the C type G_IO_ERROR_DBUS_ERROR.
	IOErrorEnum_DbusError IOErrorEnum = 36
	// IOErrorEnum_HostUnreachable is a representation of the C type G_IO_ERROR_HOST_UNREACHABLE.
	IOErrorEnum_HostUnreachable IOErrorEnum = 37
	// IOErrorEnum_NetworkUnreachable is a representation of the C type G_IO_ERROR_NETWORK_UNREACHABLE.
	IOErrorEnum_NetworkUnreachable IOErrorEnum = 38
	// IOErrorEnum_ConnectionRefused is a representation of the C type G_IO_ERROR_CONNECTION_REFUSED.
	IOErrorEnum_ConnectionRefused IOErrorEnum = 39
	// IOErrorEnum_ProxyFailed is a representation of the C type G_IO_ERROR_PROXY_FAILED.
	IOErrorEnum_ProxyFailed IOErrorEnum = 40
	// IOErrorEnum_ProxyAuthFailed is a representation of the C type G_IO_ERROR_PROXY_AUTH_FAILED.
	IOErrorEnum_ProxyAuthFailed IOErrorEnum = 41
	// IOErrorEnum_ProxyNeedAuth is a representation of the C type G_IO_ERROR_PROXY_NEED_AUTH.
	IOErrorEnum_ProxyNeedAuth IOErrorEnum = 42
	// IOErrorEnum_ProxyNotAllowed is a representation of the C type G_IO_ERROR_PROXY_NOT_ALLOWED.
	IOErrorEnum_ProxyNotAllowed IOErrorEnum = 43
	// IOErrorEnum_BrokenPipe is a representation of the C type G_IO_ERROR_BROKEN_PIPE.
	IOErrorEnum_BrokenPipe IOErrorEnum = 44
	// IOErrorEnum_ConnectionClosed is a representation of the C type G_IO_ERROR_CONNECTION_CLOSED.
	IOErrorEnum_ConnectionClosed IOErrorEnum = 44
	// IOErrorEnum_NotConnected is a representation of the C type G_IO_ERROR_NOT_CONNECTED.
	IOErrorEnum_NotConnected IOErrorEnum = 45
	// IOErrorEnum_MessageTooLarge is a representation of the C type G_IO_ERROR_MESSAGE_TOO_LARGE.
	IOErrorEnum_MessageTooLarge IOErrorEnum = 46
)

// IOModuleScopeFlags is a representation of the C type GIOModuleScopeFlags.
//
// since 2.30
type IOModuleScopeFlags int

const (
	// IOModuleScopeFlags_None is a representation of the C type G_IO_MODULE_SCOPE_NONE.
	IOModuleScopeFlags_None IOModuleScopeFlags = 0
	// IOModuleScopeFlags_BlockDuplicates is a representation of the C type G_IO_MODULE_SCOPE_BLOCK_DUPLICATES.
	IOModuleScopeFlags_BlockDuplicates IOModuleScopeFlags = 1
)

// MountOperationResult is a representation of the C type GMountOperationResult.
//
type MountOperationResult int

const (
	// MountOperationResult_Handled is a representation of the C type G_MOUNT_OPERATION_HANDLED.
	MountOperationResult_Handled MountOperationResult = 0
	// MountOperationResult_Aborted is a representation of the C type G_MOUNT_OPERATION_ABORTED.
	MountOperationResult_Aborted MountOperationResult = 1
	// MountOperationResult_Unhandled is a representation of the C type G_MOUNT_OPERATION_UNHANDLED.
	MountOperationResult_Unhandled MountOperationResult = 2
)

// NetworkConnectivity is a representation of the C type GNetworkConnectivity.
//
// since 2.44
type NetworkConnectivity int

const (
	// NetworkConnectivity_Local is a representation of the C type G_NETWORK_CONNECTIVITY_LOCAL.
	NetworkConnectivity_Local NetworkConnectivity = 1
	// NetworkConnectivity_Limited is a representation of the C type G_NETWORK_CONNECTIVITY_LIMITED.
	NetworkConnectivity_Limited NetworkConnectivity = 2
	// NetworkConnectivity_Portal is a representation of the C type G_NETWORK_CONNECTIVITY_PORTAL.
	NetworkConnectivity_Portal NetworkConnectivity = 3
	// NetworkConnectivity_Full is a representation of the C type G_NETWORK_CONNECTIVITY_FULL.
	NetworkConnectivity_Full NetworkConnectivity = 4
)

// NotificationPriority is a representation of the C type GNotificationPriority.
//
// since 2.42
type NotificationPriority int

const (
	// NotificationPriority_Normal is a representation of the C type G_NOTIFICATION_PRIORITY_NORMAL.
	NotificationPriority_Normal NotificationPriority = 0
	// NotificationPriority_Low is a representation of the C type G_NOTIFICATION_PRIORITY_LOW.
	NotificationPriority_Low NotificationPriority = 1
	// NotificationPriority_High is a representation of the C type G_NOTIFICATION_PRIORITY_HIGH.
	NotificationPriority_High NotificationPriority = 2
	// NotificationPriority_Urgent is a representation of the C type G_NOTIFICATION_PRIORITY_URGENT.
	NotificationPriority_Urgent NotificationPriority = 3
)

// PasswordSave is a representation of the C type GPasswordSave.
//
type PasswordSave int

const (
	// PasswordSave_Never is a representation of the C type G_PASSWORD_SAVE_NEVER.
	PasswordSave_Never PasswordSave = 0
	// PasswordSave_ForSession is a representation of the C type G_PASSWORD_SAVE_FOR_SESSION.
	PasswordSave_ForSession PasswordSave = 1
	// PasswordSave_Permanently is a representation of the C type G_PASSWORD_SAVE_PERMANENTLY.
	PasswordSave_Permanently PasswordSave = 2
)

// PollableReturn is a representation of the C type GPollableReturn.
//
// since 2.60
type PollableReturn int

const (
	// PollableReturn_Failed is a representation of the C type G_POLLABLE_RETURN_FAILED.
	PollableReturn_Failed PollableReturn = 0
	// PollableReturn_Ok is a representation of the C type G_POLLABLE_RETURN_OK.
	PollableReturn_Ok PollableReturn = 1
	// PollableReturn_WouldBlock is a representation of the C type G_POLLABLE_RETURN_WOULD_BLOCK.
	PollableReturn_WouldBlock PollableReturn = -27
)

// ResolverError is a representation of the C type GResolverError.
//
// since 2.22
type ResolverError int

const (
	// ResolverError_NotFound is a representation of the C type G_RESOLVER_ERROR_NOT_FOUND.
	ResolverError_NotFound ResolverError = 0
	// ResolverError_TemporaryFailure is a representation of the C type G_RESOLVER_ERROR_TEMPORARY_FAILURE.
	ResolverError_TemporaryFailure ResolverError = 1
	// ResolverError_Internal is a representation of the C type G_RESOLVER_ERROR_INTERNAL.
	ResolverError_Internal ResolverError = 2
)

// ResolverRecordType is a representation of the C type GResolverRecordType.
//
// since 2.34
type ResolverRecordType int

const (
	// ResolverRecordType_Srv is a representation of the C type G_RESOLVER_RECORD_SRV.
	ResolverRecordType_Srv ResolverRecordType = 1
	// ResolverRecordType_Mx is a representation of the C type G_RESOLVER_RECORD_MX.
	ResolverRecordType_Mx ResolverRecordType = 2
	// ResolverRecordType_Txt is a representation of the C type G_RESOLVER_RECORD_TXT.
	ResolverRecordType_Txt ResolverRecordType = 3
	// ResolverRecordType_Soa is a representation of the C type G_RESOLVER_RECORD_SOA.
	ResolverRecordType_Soa ResolverRecordType = 4
	// ResolverRecordType_Ns is a representation of the C type G_RESOLVER_RECORD_NS.
	ResolverRecordType_Ns ResolverRecordType = 5
)

// ResourceError is a representation of the C type GResourceError.
//
// since 2.32
type ResourceError int

const (
	// ResourceError_NotFound is a representation of the C type G_RESOURCE_ERROR_NOT_FOUND.
	ResourceError_NotFound ResourceError = 0
	// ResourceError_Internal is a representation of the C type G_RESOURCE_ERROR_INTERNAL.
	ResourceError_Internal ResourceError = 1
)

// SocketClientEvent is a representation of the C type GSocketClientEvent.
//
// since 2.32
type SocketClientEvent int

const (
	// SocketClientEvent_Resolving is a representation of the C type G_SOCKET_CLIENT_RESOLVING.
	SocketClientEvent_Resolving SocketClientEvent = 0
	// SocketClientEvent_Resolved is a representation of the C type G_SOCKET_CLIENT_RESOLVED.
	SocketClientEvent_Resolved SocketClientEvent = 1
	// SocketClientEvent_Connecting is a representation of the C type G_SOCKET_CLIENT_CONNECTING.
	SocketClientEvent_Connecting SocketClientEvent = 2
	// SocketClientEvent_Connected is a representation of the C type G_SOCKET_CLIENT_CONNECTED.
	SocketClientEvent_Connected SocketClientEvent = 3
	// SocketClientEvent_ProxyNegotiating is a representation of the C type G_SOCKET_CLIENT_PROXY_NEGOTIATING.
	SocketClientEvent_ProxyNegotiating SocketClientEvent = 4
	// SocketClientEvent_ProxyNegotiated is a representation of the C type G_SOCKET_CLIENT_PROXY_NEGOTIATED.
	SocketClientEvent_ProxyNegotiated SocketClientEvent = 5
	// SocketClientEvent_TlsHandshaking is a representation of the C type G_SOCKET_CLIENT_TLS_HANDSHAKING.
	SocketClientEvent_TlsHandshaking SocketClientEvent = 6
	// SocketClientEvent_TlsHandshaked is a representation of the C type G_SOCKET_CLIENT_TLS_HANDSHAKED.
	SocketClientEvent_TlsHandshaked SocketClientEvent = 7
	// SocketClientEvent_Complete is a representation of the C type G_SOCKET_CLIENT_COMPLETE.
	SocketClientEvent_Complete SocketClientEvent = 8
)

// SocketFamily is a representation of the C type GSocketFamily.
//
// since 2.22
type SocketFamily int

const (
	// SocketFamily_Invalid is a representation of the C type G_SOCKET_FAMILY_INVALID.
	SocketFamily_Invalid SocketFamily = 0
	// SocketFamily_Unix is a representation of the C type G_SOCKET_FAMILY_UNIX.
	SocketFamily_Unix SocketFamily = 1
	// SocketFamily_Ipv4 is a representation of the C type G_SOCKET_FAMILY_IPV4.
	SocketFamily_Ipv4 SocketFamily = 2
	// SocketFamily_Ipv6 is a representation of the C type G_SOCKET_FAMILY_IPV6.
	SocketFamily_Ipv6 SocketFamily = 10
)

// SocketListenerEvent is a representation of the C type GSocketListenerEvent.
//
// since 2.46
type SocketListenerEvent int

const (
	// SocketListenerEvent_Binding is a representation of the C type G_SOCKET_LISTENER_BINDING.
	SocketListenerEvent_Binding SocketListenerEvent = 0
	// SocketListenerEvent_Bound is a representation of the C type G_SOCKET_LISTENER_BOUND.
	SocketListenerEvent_Bound SocketListenerEvent = 1
	// SocketListenerEvent_Listening is a representation of the C type G_SOCKET_LISTENER_LISTENING.
	SocketListenerEvent_Listening SocketListenerEvent = 2
	// SocketListenerEvent_Listened is a representation of the C type G_SOCKET_LISTENER_LISTENED.
	SocketListenerEvent_Listened SocketListenerEvent = 3
)

// SocketProtocol is a representation of the C type GSocketProtocol.
//
// since 2.22
type SocketProtocol int

const (
	// SocketProtocol_Unknown is a representation of the C type G_SOCKET_PROTOCOL_UNKNOWN.
	SocketProtocol_Unknown SocketProtocol = -1
	// SocketProtocol_Default is a representation of the C type G_SOCKET_PROTOCOL_DEFAULT.
	SocketProtocol_Default SocketProtocol = 0
	// SocketProtocol_Tcp is a representation of the C type G_SOCKET_PROTOCOL_TCP.
	SocketProtocol_Tcp SocketProtocol = 6
	// SocketProtocol_Udp is a representation of the C type G_SOCKET_PROTOCOL_UDP.
	SocketProtocol_Udp SocketProtocol = 17
	// SocketProtocol_Sctp is a representation of the C type G_SOCKET_PROTOCOL_SCTP.
	SocketProtocol_Sctp SocketProtocol = 132
)

// SocketType is a representation of the C type GSocketType.
//
// since 2.22
type SocketType int

const (
	// SocketType_Invalid is a representation of the C type G_SOCKET_TYPE_INVALID.
	SocketType_Invalid SocketType = 0
	// SocketType_Stream is a representation of the C type G_SOCKET_TYPE_STREAM.
	SocketType_Stream SocketType = 1
	// SocketType_Datagram is a representation of the C type G_SOCKET_TYPE_DATAGRAM.
	SocketType_Datagram SocketType = 2
	// SocketType_Seqpacket is a representation of the C type G_SOCKET_TYPE_SEQPACKET.
	SocketType_Seqpacket SocketType = 3
)

// TlsAuthenticationMode is a representation of the C type GTlsAuthenticationMode.
//
// since 2.28
type TlsAuthenticationMode int

const (
	// TlsAuthenticationMode_None is a representation of the C type G_TLS_AUTHENTICATION_NONE.
	TlsAuthenticationMode_None TlsAuthenticationMode = 0
	// TlsAuthenticationMode_Requested is a representation of the C type G_TLS_AUTHENTICATION_REQUESTED.
	TlsAuthenticationMode_Requested TlsAuthenticationMode = 1
	// TlsAuthenticationMode_Required is a representation of the C type G_TLS_AUTHENTICATION_REQUIRED.
	TlsAuthenticationMode_Required TlsAuthenticationMode = 2
)

// TlsCertificateRequestFlags is a representation of the C type GTlsCertificateRequestFlags.
//
// since 2.40
type TlsCertificateRequestFlags int

const (
	// TlsCertificateRequestFlags_None is a representation of the C type G_TLS_CERTIFICATE_REQUEST_NONE.
	TlsCertificateRequestFlags_None TlsCertificateRequestFlags = 0
)

// TlsDatabaseLookupFlags is a representation of the C type GTlsDatabaseLookupFlags.
//
// since 2.30
type TlsDatabaseLookupFlags int

const (
	// TlsDatabaseLookupFlags_None is a representation of the C type G_TLS_DATABASE_LOOKUP_NONE.
	TlsDatabaseLookupFlags_None TlsDatabaseLookupFlags = 0
	// TlsDatabaseLookupFlags_Keypair is a representation of the C type G_TLS_DATABASE_LOOKUP_KEYPAIR.
	TlsDatabaseLookupFlags_Keypair TlsDatabaseLookupFlags = 1
)

// TlsError is a representation of the C type GTlsError.
//
// since 2.28
type TlsError int

const (
	// TlsError_Unavailable is a representation of the C type G_TLS_ERROR_UNAVAILABLE.
	TlsError_Unavailable TlsError = 0
	// TlsError_Misc is a representation of the C type G_TLS_ERROR_MISC.
	TlsError_Misc TlsError = 1
	// TlsError_BadCertificate is a representation of the C type G_TLS_ERROR_BAD_CERTIFICATE.
	TlsError_BadCertificate TlsError = 2
	// TlsError_NotTls is a representation of the C type G_TLS_ERROR_NOT_TLS.
	TlsError_NotTls TlsError = 3
	// TlsError_Handshake is a representation of the C type G_TLS_ERROR_HANDSHAKE.
	TlsError_Handshake TlsError = 4
	// TlsError_CertificateRequired is a representation of the C type G_TLS_ERROR_CERTIFICATE_REQUIRED.
	TlsError_CertificateRequired TlsError = 5
	// TlsError_Eof is a representation of the C type G_TLS_ERROR_EOF.
	TlsError_Eof TlsError = 6
	// TlsError_InappropriateFallback is a representation of the C type G_TLS_ERROR_INAPPROPRIATE_FALLBACK.
	TlsError_InappropriateFallback TlsError = 7
)

// TlsInteractionResult is a representation of the C type GTlsInteractionResult.
//
// since 2.30
type TlsInteractionResult int

const (
	// TlsInteractionResult_Unhandled is a representation of the C type G_TLS_INTERACTION_UNHANDLED.
	TlsInteractionResult_Unhandled TlsInteractionResult = 0
	// TlsInteractionResult_Handled is a representation of the C type G_TLS_INTERACTION_HANDLED.
	TlsInteractionResult_Handled TlsInteractionResult = 1
	// TlsInteractionResult_Failed is a representation of the C type G_TLS_INTERACTION_FAILED.
	TlsInteractionResult_Failed TlsInteractionResult = 2
)

// TlsRehandshakeMode is a representation of the C type GTlsRehandshakeMode.
//
// since 2.28
type TlsRehandshakeMode int

const (
	// TlsRehandshakeMode_Never is a representation of the C type G_TLS_REHANDSHAKE_NEVER.
	TlsRehandshakeMode_Never TlsRehandshakeMode = 0
	// TlsRehandshakeMode_Safely is a representation of the C type G_TLS_REHANDSHAKE_SAFELY.
	TlsRehandshakeMode_Safely TlsRehandshakeMode = 1
	// TlsRehandshakeMode_Unsafely is a representation of the C type G_TLS_REHANDSHAKE_UNSAFELY.
	TlsRehandshakeMode_Unsafely TlsRehandshakeMode = 2
)

// UnixSocketAddressType is a representation of the C type GUnixSocketAddressType.
//
// since 2.26
type UnixSocketAddressType int

const (
	// UnixSocketAddressType_Invalid is a representation of the C type G_UNIX_SOCKET_ADDRESS_INVALID.
	UnixSocketAddressType_Invalid UnixSocketAddressType = 0
	// UnixSocketAddressType_Anonymous is a representation of the C type G_UNIX_SOCKET_ADDRESS_ANONYMOUS.
	UnixSocketAddressType_Anonymous UnixSocketAddressType = 1
	// UnixSocketAddressType_Path is a representation of the C type G_UNIX_SOCKET_ADDRESS_PATH.
	UnixSocketAddressType_Path UnixSocketAddressType = 2
	// UnixSocketAddressType_Abstract is a representation of the C type G_UNIX_SOCKET_ADDRESS_ABSTRACT.
	UnixSocketAddressType_Abstract UnixSocketAddressType = 3
	// UnixSocketAddressType_AbstractPadded is a representation of the C type G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
	UnixSocketAddressType_AbstractPadded UnixSocketAddressType = 4
)

// ZlibCompressorFormat is a representation of the C type GZlibCompressorFormat.
//
// since 2.24
type ZlibCompressorFormat int

const (
	// ZlibCompressorFormat_Zlib is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_ZLIB.
	ZlibCompressorFormat_Zlib ZlibCompressorFormat = 0
	// ZlibCompressorFormat_Gzip is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_GZIP.
	ZlibCompressorFormat_Gzip ZlibCompressorFormat = 1
	// ZlibCompressorFormat_Raw is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_RAW.
	ZlibCompressorFormat_Raw ZlibCompressorFormat = 2
)
