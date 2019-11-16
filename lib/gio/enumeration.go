// Code generated - DO NOT EDIT.

package gio

// Bustype is a representation of the C type GBusType.
//
// since 2.26
type Bustype int

const (
	// Bustype_Starter is a representation of the C type G_BUS_TYPE_STARTER.
	Bustype_Starter Bustype = -1
	// Bustype_None is a representation of the C type G_BUS_TYPE_NONE.
	Bustype_None Bustype = 0
	// Bustype_System is a representation of the C type G_BUS_TYPE_SYSTEM.
	Bustype_System Bustype = 1
	// Bustype_Session is a representation of the C type G_BUS_TYPE_SESSION.
	Bustype_Session Bustype = 2
)

// Converterresult is a representation of the C type GConverterResult.
//
// since 2.24
type Converterresult int

const (
	// Converterresult_Error is a representation of the C type G_CONVERTER_ERROR.
	Converterresult_Error Converterresult = 0
	// Converterresult_Converted is a representation of the C type G_CONVERTER_CONVERTED.
	Converterresult_Converted Converterresult = 1
	// Converterresult_Finished is a representation of the C type G_CONVERTER_FINISHED.
	Converterresult_Finished Converterresult = 2
	// Converterresult_Flushed is a representation of the C type G_CONVERTER_FLUSHED.
	Converterresult_Flushed Converterresult = 3
)

// Credentialstype is a representation of the C type GCredentialsType.
//
// since 2.26
type Credentialstype int

const (
	// Credentialstype_Invalid is a representation of the C type G_CREDENTIALS_TYPE_INVALID.
	Credentialstype_Invalid Credentialstype = 0
	// Credentialstype_LinuxUcred is a representation of the C type G_CREDENTIALS_TYPE_LINUX_UCRED.
	Credentialstype_LinuxUcred Credentialstype = 1
	// Credentialstype_FreebsdCmsgcred is a representation of the C type G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
	Credentialstype_FreebsdCmsgcred Credentialstype = 2
	// Credentialstype_OpenbsdSockpeercred is a representation of the C type G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
	Credentialstype_OpenbsdSockpeercred Credentialstype = 3
	// Credentialstype_SolarisUcred is a representation of the C type G_CREDENTIALS_TYPE_SOLARIS_UCRED.
	Credentialstype_SolarisUcred Credentialstype = 4
	// Credentialstype_NetbsdUnpcbid is a representation of the C type G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
	Credentialstype_NetbsdUnpcbid Credentialstype = 5
)

// Dbuserror is a representation of the C type GDBusError.
//
// since 2.26
type Dbuserror int

const (
	// Dbuserror_Failed is a representation of the C type G_DBUS_ERROR_FAILED.
	Dbuserror_Failed Dbuserror = 0
	// Dbuserror_NoMemory is a representation of the C type G_DBUS_ERROR_NO_MEMORY.
	Dbuserror_NoMemory Dbuserror = 1
	// Dbuserror_ServiceUnknown is a representation of the C type G_DBUS_ERROR_SERVICE_UNKNOWN.
	Dbuserror_ServiceUnknown Dbuserror = 2
	// Dbuserror_NameHasNoOwner is a representation of the C type G_DBUS_ERROR_NAME_HAS_NO_OWNER.
	Dbuserror_NameHasNoOwner Dbuserror = 3
	// Dbuserror_NoReply is a representation of the C type G_DBUS_ERROR_NO_REPLY.
	Dbuserror_NoReply Dbuserror = 4
	// Dbuserror_IoError is a representation of the C type G_DBUS_ERROR_IO_ERROR.
	Dbuserror_IoError Dbuserror = 5
	// Dbuserror_BadAddress is a representation of the C type G_DBUS_ERROR_BAD_ADDRESS.
	Dbuserror_BadAddress Dbuserror = 6
	// Dbuserror_NotSupported is a representation of the C type G_DBUS_ERROR_NOT_SUPPORTED.
	Dbuserror_NotSupported Dbuserror = 7
	// Dbuserror_LimitsExceeded is a representation of the C type G_DBUS_ERROR_LIMITS_EXCEEDED.
	Dbuserror_LimitsExceeded Dbuserror = 8
	// Dbuserror_AccessDenied is a representation of the C type G_DBUS_ERROR_ACCESS_DENIED.
	Dbuserror_AccessDenied Dbuserror = 9
	// Dbuserror_AuthFailed is a representation of the C type G_DBUS_ERROR_AUTH_FAILED.
	Dbuserror_AuthFailed Dbuserror = 10
	// Dbuserror_NoServer is a representation of the C type G_DBUS_ERROR_NO_SERVER.
	Dbuserror_NoServer Dbuserror = 11
	// Dbuserror_Timeout is a representation of the C type G_DBUS_ERROR_TIMEOUT.
	Dbuserror_Timeout Dbuserror = 12
	// Dbuserror_NoNetwork is a representation of the C type G_DBUS_ERROR_NO_NETWORK.
	Dbuserror_NoNetwork Dbuserror = 13
	// Dbuserror_AddressInUse is a representation of the C type G_DBUS_ERROR_ADDRESS_IN_USE.
	Dbuserror_AddressInUse Dbuserror = 14
	// Dbuserror_Disconnected is a representation of the C type G_DBUS_ERROR_DISCONNECTED.
	Dbuserror_Disconnected Dbuserror = 15
	// Dbuserror_InvalidArgs is a representation of the C type G_DBUS_ERROR_INVALID_ARGS.
	Dbuserror_InvalidArgs Dbuserror = 16
	// Dbuserror_FileNotFound is a representation of the C type G_DBUS_ERROR_FILE_NOT_FOUND.
	Dbuserror_FileNotFound Dbuserror = 17
	// Dbuserror_FileExists is a representation of the C type G_DBUS_ERROR_FILE_EXISTS.
	Dbuserror_FileExists Dbuserror = 18
	// Dbuserror_UnknownMethod is a representation of the C type G_DBUS_ERROR_UNKNOWN_METHOD.
	Dbuserror_UnknownMethod Dbuserror = 19
	// Dbuserror_TimedOut is a representation of the C type G_DBUS_ERROR_TIMED_OUT.
	Dbuserror_TimedOut Dbuserror = 20
	// Dbuserror_MatchRuleNotFound is a representation of the C type G_DBUS_ERROR_MATCH_RULE_NOT_FOUND.
	Dbuserror_MatchRuleNotFound Dbuserror = 21
	// Dbuserror_MatchRuleInvalid is a representation of the C type G_DBUS_ERROR_MATCH_RULE_INVALID.
	Dbuserror_MatchRuleInvalid Dbuserror = 22
	// Dbuserror_SpawnExecFailed is a representation of the C type G_DBUS_ERROR_SPAWN_EXEC_FAILED.
	Dbuserror_SpawnExecFailed Dbuserror = 23
	// Dbuserror_SpawnForkFailed is a representation of the C type G_DBUS_ERROR_SPAWN_FORK_FAILED.
	Dbuserror_SpawnForkFailed Dbuserror = 24
	// Dbuserror_SpawnChildExited is a representation of the C type G_DBUS_ERROR_SPAWN_CHILD_EXITED.
	Dbuserror_SpawnChildExited Dbuserror = 25
	// Dbuserror_SpawnChildSignaled is a representation of the C type G_DBUS_ERROR_SPAWN_CHILD_SIGNALED.
	Dbuserror_SpawnChildSignaled Dbuserror = 26
	// Dbuserror_SpawnFailed is a representation of the C type G_DBUS_ERROR_SPAWN_FAILED.
	Dbuserror_SpawnFailed Dbuserror = 27
	// Dbuserror_SpawnSetupFailed is a representation of the C type G_DBUS_ERROR_SPAWN_SETUP_FAILED.
	Dbuserror_SpawnSetupFailed Dbuserror = 28
	// Dbuserror_SpawnConfigInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_CONFIG_INVALID.
	Dbuserror_SpawnConfigInvalid Dbuserror = 29
	// Dbuserror_SpawnServiceInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_SERVICE_INVALID.
	Dbuserror_SpawnServiceInvalid Dbuserror = 30
	// Dbuserror_SpawnServiceNotFound is a representation of the C type G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND.
	Dbuserror_SpawnServiceNotFound Dbuserror = 31
	// Dbuserror_SpawnPermissionsInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID.
	Dbuserror_SpawnPermissionsInvalid Dbuserror = 32
	// Dbuserror_SpawnFileInvalid is a representation of the C type G_DBUS_ERROR_SPAWN_FILE_INVALID.
	Dbuserror_SpawnFileInvalid Dbuserror = 33
	// Dbuserror_SpawnNoMemory is a representation of the C type G_DBUS_ERROR_SPAWN_NO_MEMORY.
	Dbuserror_SpawnNoMemory Dbuserror = 34
	// Dbuserror_UnixProcessIdUnknown is a representation of the C type G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN.
	Dbuserror_UnixProcessIdUnknown Dbuserror = 35
	// Dbuserror_InvalidSignature is a representation of the C type G_DBUS_ERROR_INVALID_SIGNATURE.
	Dbuserror_InvalidSignature Dbuserror = 36
	// Dbuserror_InvalidFileContent is a representation of the C type G_DBUS_ERROR_INVALID_FILE_CONTENT.
	Dbuserror_InvalidFileContent Dbuserror = 37
	// Dbuserror_SelinuxSecurityContextUnknown is a representation of the C type G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN.
	Dbuserror_SelinuxSecurityContextUnknown Dbuserror = 38
	// Dbuserror_AdtAuditDataUnknown is a representation of the C type G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN.
	Dbuserror_AdtAuditDataUnknown Dbuserror = 39
	// Dbuserror_ObjectPathInUse is a representation of the C type G_DBUS_ERROR_OBJECT_PATH_IN_USE.
	Dbuserror_ObjectPathInUse Dbuserror = 40
	// Dbuserror_UnknownObject is a representation of the C type G_DBUS_ERROR_UNKNOWN_OBJECT.
	Dbuserror_UnknownObject Dbuserror = 41
	// Dbuserror_UnknownInterface is a representation of the C type G_DBUS_ERROR_UNKNOWN_INTERFACE.
	Dbuserror_UnknownInterface Dbuserror = 42
	// Dbuserror_UnknownProperty is a representation of the C type G_DBUS_ERROR_UNKNOWN_PROPERTY.
	Dbuserror_UnknownProperty Dbuserror = 43
	// Dbuserror_PropertyReadOnly is a representation of the C type G_DBUS_ERROR_PROPERTY_READ_ONLY.
	Dbuserror_PropertyReadOnly Dbuserror = 44
)

// Dbusmessagebyteorder is a representation of the C type GDBusMessageByteOrder.
//
// since 2.26
type Dbusmessagebyteorder int

const (
	// Dbusmessagebyteorder_BigEndian is a representation of the C type G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN.
	Dbusmessagebyteorder_BigEndian Dbusmessagebyteorder = 66
	// Dbusmessagebyteorder_LittleEndian is a representation of the C type G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN.
	Dbusmessagebyteorder_LittleEndian Dbusmessagebyteorder = 108
)

// Dbusmessageheaderfield is a representation of the C type GDBusMessageHeaderField.
//
// since 2.26
type Dbusmessageheaderfield int

const (
	// Dbusmessageheaderfield_Invalid is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_INVALID.
	Dbusmessageheaderfield_Invalid Dbusmessageheaderfield = 0
	// Dbusmessageheaderfield_Path is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_PATH.
	Dbusmessageheaderfield_Path Dbusmessageheaderfield = 1
	// Dbusmessageheaderfield_Interface is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE.
	Dbusmessageheaderfield_Interface Dbusmessageheaderfield = 2
	// Dbusmessageheaderfield_Member is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_MEMBER.
	Dbusmessageheaderfield_Member Dbusmessageheaderfield = 3
	// Dbusmessageheaderfield_ErrorName is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME.
	Dbusmessageheaderfield_ErrorName Dbusmessageheaderfield = 4
	// Dbusmessageheaderfield_ReplySerial is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL.
	Dbusmessageheaderfield_ReplySerial Dbusmessageheaderfield = 5
	// Dbusmessageheaderfield_Destination is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION.
	Dbusmessageheaderfield_Destination Dbusmessageheaderfield = 6
	// Dbusmessageheaderfield_Sender is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_SENDER.
	Dbusmessageheaderfield_Sender Dbusmessageheaderfield = 7
	// Dbusmessageheaderfield_Signature is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE.
	Dbusmessageheaderfield_Signature Dbusmessageheaderfield = 8
	// Dbusmessageheaderfield_NumUnixFds is a representation of the C type G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS.
	Dbusmessageheaderfield_NumUnixFds Dbusmessageheaderfield = 9
)

// Dbusmessagetype is a representation of the C type GDBusMessageType.
//
// since 2.26
type Dbusmessagetype int

const (
	// Dbusmessagetype_Invalid is a representation of the C type G_DBUS_MESSAGE_TYPE_INVALID.
	Dbusmessagetype_Invalid Dbusmessagetype = 0
	// Dbusmessagetype_MethodCall is a representation of the C type G_DBUS_MESSAGE_TYPE_METHOD_CALL.
	Dbusmessagetype_MethodCall Dbusmessagetype = 1
	// Dbusmessagetype_MethodReturn is a representation of the C type G_DBUS_MESSAGE_TYPE_METHOD_RETURN.
	Dbusmessagetype_MethodReturn Dbusmessagetype = 2
	// Dbusmessagetype_Error is a representation of the C type G_DBUS_MESSAGE_TYPE_ERROR.
	Dbusmessagetype_Error Dbusmessagetype = 3
	// Dbusmessagetype_Signal is a representation of the C type G_DBUS_MESSAGE_TYPE_SIGNAL.
	Dbusmessagetype_Signal Dbusmessagetype = 4
)

// Datastreambyteorder is a representation of the C type GDataStreamByteOrder.
type Datastreambyteorder int

const (
	// Datastreambyteorder_BigEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN.
	Datastreambyteorder_BigEndian Datastreambyteorder = 0
	// Datastreambyteorder_LittleEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN.
	Datastreambyteorder_LittleEndian Datastreambyteorder = 1
	// Datastreambyteorder_HostEndian is a representation of the C type G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN.
	Datastreambyteorder_HostEndian Datastreambyteorder = 2
)

// Datastreamnewlinetype is a representation of the C type GDataStreamNewlineType.
type Datastreamnewlinetype int

const (
	// Datastreamnewlinetype_Lf is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_LF.
	Datastreamnewlinetype_Lf Datastreamnewlinetype = 0
	// Datastreamnewlinetype_Cr is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_CR.
	Datastreamnewlinetype_Cr Datastreamnewlinetype = 1
	// Datastreamnewlinetype_CrLf is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_CR_LF.
	Datastreamnewlinetype_CrLf Datastreamnewlinetype = 2
	// Datastreamnewlinetype_Any is a representation of the C type G_DATA_STREAM_NEWLINE_TYPE_ANY.
	Datastreamnewlinetype_Any Datastreamnewlinetype = 3
)

// Drivestartstoptype is a representation of the C type GDriveStartStopType.
//
// since 2.22
type Drivestartstoptype int

const (
	// Drivestartstoptype_Unknown is a representation of the C type G_DRIVE_START_STOP_TYPE_UNKNOWN.
	Drivestartstoptype_Unknown Drivestartstoptype = 0
	// Drivestartstoptype_Shutdown is a representation of the C type G_DRIVE_START_STOP_TYPE_SHUTDOWN.
	Drivestartstoptype_Shutdown Drivestartstoptype = 1
	// Drivestartstoptype_Network is a representation of the C type G_DRIVE_START_STOP_TYPE_NETWORK.
	Drivestartstoptype_Network Drivestartstoptype = 2
	// Drivestartstoptype_Multidisk is a representation of the C type G_DRIVE_START_STOP_TYPE_MULTIDISK.
	Drivestartstoptype_Multidisk Drivestartstoptype = 3
	// Drivestartstoptype_Password is a representation of the C type G_DRIVE_START_STOP_TYPE_PASSWORD.
	Drivestartstoptype_Password Drivestartstoptype = 4
)

// Emblemorigin is a representation of the C type GEmblemOrigin.
//
// since 2.18
type Emblemorigin int

const (
	// Emblemorigin_Unknown is a representation of the C type G_EMBLEM_ORIGIN_UNKNOWN.
	Emblemorigin_Unknown Emblemorigin = 0
	// Emblemorigin_Device is a representation of the C type G_EMBLEM_ORIGIN_DEVICE.
	Emblemorigin_Device Emblemorigin = 1
	// Emblemorigin_Livemetadata is a representation of the C type G_EMBLEM_ORIGIN_LIVEMETADATA.
	Emblemorigin_Livemetadata Emblemorigin = 2
	// Emblemorigin_Tag is a representation of the C type G_EMBLEM_ORIGIN_TAG.
	Emblemorigin_Tag Emblemorigin = 3
)

// Fileattributestatus is a representation of the C type GFileAttributeStatus.
type Fileattributestatus int

const (
	// Fileattributestatus_Unset is a representation of the C type G_FILE_ATTRIBUTE_STATUS_UNSET.
	Fileattributestatus_Unset Fileattributestatus = 0
	// Fileattributestatus_Set is a representation of the C type G_FILE_ATTRIBUTE_STATUS_SET.
	Fileattributestatus_Set Fileattributestatus = 1
	// Fileattributestatus_ErrorSetting is a representation of the C type G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING.
	Fileattributestatus_ErrorSetting Fileattributestatus = 2
)

// Fileattributetype is a representation of the C type GFileAttributeType.
type Fileattributetype int

const (
	// Fileattributetype_Invalid is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INVALID.
	Fileattributetype_Invalid Fileattributetype = 0
	// Fileattributetype_String is a representation of the C type G_FILE_ATTRIBUTE_TYPE_STRING.
	Fileattributetype_String Fileattributetype = 1
	// Fileattributetype_ByteString is a representation of the C type G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
	Fileattributetype_ByteString Fileattributetype = 2
	// Fileattributetype_Boolean is a representation of the C type G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
	Fileattributetype_Boolean Fileattributetype = 3
	// Fileattributetype_Uint32 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_UINT32.
	Fileattributetype_Uint32 Fileattributetype = 4
	// Fileattributetype_Int32 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INT32.
	Fileattributetype_Int32 Fileattributetype = 5
	// Fileattributetype_Uint64 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_UINT64.
	Fileattributetype_Uint64 Fileattributetype = 6
	// Fileattributetype_Int64 is a representation of the C type G_FILE_ATTRIBUTE_TYPE_INT64.
	Fileattributetype_Int64 Fileattributetype = 7
	// Fileattributetype_Object is a representation of the C type G_FILE_ATTRIBUTE_TYPE_OBJECT.
	Fileattributetype_Object Fileattributetype = 8
	// Fileattributetype_Stringv is a representation of the C type G_FILE_ATTRIBUTE_TYPE_STRINGV.
	Fileattributetype_Stringv Fileattributetype = 9
)

// Filemonitorevent is a representation of the C type GFileMonitorEvent.
type Filemonitorevent int

const (
	// Filemonitorevent_Changed is a representation of the C type G_FILE_MONITOR_EVENT_CHANGED.
	Filemonitorevent_Changed Filemonitorevent = 0
	// Filemonitorevent_ChangesDoneHint is a representation of the C type G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT.
	Filemonitorevent_ChangesDoneHint Filemonitorevent = 1
	// Filemonitorevent_Deleted is a representation of the C type G_FILE_MONITOR_EVENT_DELETED.
	Filemonitorevent_Deleted Filemonitorevent = 2
	// Filemonitorevent_Created is a representation of the C type G_FILE_MONITOR_EVENT_CREATED.
	Filemonitorevent_Created Filemonitorevent = 3
	// Filemonitorevent_AttributeChanged is a representation of the C type G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED.
	Filemonitorevent_AttributeChanged Filemonitorevent = 4
	// Filemonitorevent_PreUnmount is a representation of the C type G_FILE_MONITOR_EVENT_PRE_UNMOUNT.
	Filemonitorevent_PreUnmount Filemonitorevent = 5
	// Filemonitorevent_Unmounted is a representation of the C type G_FILE_MONITOR_EVENT_UNMOUNTED.
	Filemonitorevent_Unmounted Filemonitorevent = 6
	// Filemonitorevent_Moved is a representation of the C type G_FILE_MONITOR_EVENT_MOVED.
	Filemonitorevent_Moved Filemonitorevent = 7
	// Filemonitorevent_Renamed is a representation of the C type G_FILE_MONITOR_EVENT_RENAMED.
	Filemonitorevent_Renamed Filemonitorevent = 8
	// Filemonitorevent_MovedIn is a representation of the C type G_FILE_MONITOR_EVENT_MOVED_IN.
	Filemonitorevent_MovedIn Filemonitorevent = 9
	// Filemonitorevent_MovedOut is a representation of the C type G_FILE_MONITOR_EVENT_MOVED_OUT.
	Filemonitorevent_MovedOut Filemonitorevent = 10
)

// Filetype is a representation of the C type GFileType.
type Filetype int

const (
	// Filetype_Unknown is a representation of the C type G_FILE_TYPE_UNKNOWN.
	Filetype_Unknown Filetype = 0
	// Filetype_Regular is a representation of the C type G_FILE_TYPE_REGULAR.
	Filetype_Regular Filetype = 1
	// Filetype_Directory is a representation of the C type G_FILE_TYPE_DIRECTORY.
	Filetype_Directory Filetype = 2
	// Filetype_SymbolicLink is a representation of the C type G_FILE_TYPE_SYMBOLIC_LINK.
	Filetype_SymbolicLink Filetype = 3
	// Filetype_Special is a representation of the C type G_FILE_TYPE_SPECIAL.
	Filetype_Special Filetype = 4
	// Filetype_Shortcut is a representation of the C type G_FILE_TYPE_SHORTCUT.
	Filetype_Shortcut Filetype = 5
	// Filetype_Mountable is a representation of the C type G_FILE_TYPE_MOUNTABLE.
	Filetype_Mountable Filetype = 6
)

// Filesystempreviewtype is a representation of the C type GFilesystemPreviewType.
type Filesystempreviewtype int

const (
	// Filesystempreviewtype_IfAlways is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS.
	Filesystempreviewtype_IfAlways Filesystempreviewtype = 0
	// Filesystempreviewtype_IfLocal is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL.
	Filesystempreviewtype_IfLocal Filesystempreviewtype = 1
	// Filesystempreviewtype_Never is a representation of the C type G_FILESYSTEM_PREVIEW_TYPE_NEVER.
	Filesystempreviewtype_Never Filesystempreviewtype = 2
)

// Ioerrorenum is a representation of the C type GIOErrorEnum.
type Ioerrorenum int

const (
	// Ioerrorenum_Failed is a representation of the C type G_IO_ERROR_FAILED.
	Ioerrorenum_Failed Ioerrorenum = 0
	// Ioerrorenum_NotFound is a representation of the C type G_IO_ERROR_NOT_FOUND.
	Ioerrorenum_NotFound Ioerrorenum = 1
	// Ioerrorenum_Exists is a representation of the C type G_IO_ERROR_EXISTS.
	Ioerrorenum_Exists Ioerrorenum = 2
	// Ioerrorenum_IsDirectory is a representation of the C type G_IO_ERROR_IS_DIRECTORY.
	Ioerrorenum_IsDirectory Ioerrorenum = 3
	// Ioerrorenum_NotDirectory is a representation of the C type G_IO_ERROR_NOT_DIRECTORY.
	Ioerrorenum_NotDirectory Ioerrorenum = 4
	// Ioerrorenum_NotEmpty is a representation of the C type G_IO_ERROR_NOT_EMPTY.
	Ioerrorenum_NotEmpty Ioerrorenum = 5
	// Ioerrorenum_NotRegularFile is a representation of the C type G_IO_ERROR_NOT_REGULAR_FILE.
	Ioerrorenum_NotRegularFile Ioerrorenum = 6
	// Ioerrorenum_NotSymbolicLink is a representation of the C type G_IO_ERROR_NOT_SYMBOLIC_LINK.
	Ioerrorenum_NotSymbolicLink Ioerrorenum = 7
	// Ioerrorenum_NotMountableFile is a representation of the C type G_IO_ERROR_NOT_MOUNTABLE_FILE.
	Ioerrorenum_NotMountableFile Ioerrorenum = 8
	// Ioerrorenum_FilenameTooLong is a representation of the C type G_IO_ERROR_FILENAME_TOO_LONG.
	Ioerrorenum_FilenameTooLong Ioerrorenum = 9
	// Ioerrorenum_InvalidFilename is a representation of the C type G_IO_ERROR_INVALID_FILENAME.
	Ioerrorenum_InvalidFilename Ioerrorenum = 10
	// Ioerrorenum_TooManyLinks is a representation of the C type G_IO_ERROR_TOO_MANY_LINKS.
	Ioerrorenum_TooManyLinks Ioerrorenum = 11
	// Ioerrorenum_NoSpace is a representation of the C type G_IO_ERROR_NO_SPACE.
	Ioerrorenum_NoSpace Ioerrorenum = 12
	// Ioerrorenum_InvalidArgument is a representation of the C type G_IO_ERROR_INVALID_ARGUMENT.
	Ioerrorenum_InvalidArgument Ioerrorenum = 13
	// Ioerrorenum_PermissionDenied is a representation of the C type G_IO_ERROR_PERMISSION_DENIED.
	Ioerrorenum_PermissionDenied Ioerrorenum = 14
	// Ioerrorenum_NotSupported is a representation of the C type G_IO_ERROR_NOT_SUPPORTED.
	Ioerrorenum_NotSupported Ioerrorenum = 15
	// Ioerrorenum_NotMounted is a representation of the C type G_IO_ERROR_NOT_MOUNTED.
	Ioerrorenum_NotMounted Ioerrorenum = 16
	// Ioerrorenum_AlreadyMounted is a representation of the C type G_IO_ERROR_ALREADY_MOUNTED.
	Ioerrorenum_AlreadyMounted Ioerrorenum = 17
	// Ioerrorenum_Closed is a representation of the C type G_IO_ERROR_CLOSED.
	Ioerrorenum_Closed Ioerrorenum = 18
	// Ioerrorenum_Cancelled is a representation of the C type G_IO_ERROR_CANCELLED.
	Ioerrorenum_Cancelled Ioerrorenum = 19
	// Ioerrorenum_Pending is a representation of the C type G_IO_ERROR_PENDING.
	Ioerrorenum_Pending Ioerrorenum = 20
	// Ioerrorenum_ReadOnly is a representation of the C type G_IO_ERROR_READ_ONLY.
	Ioerrorenum_ReadOnly Ioerrorenum = 21
	// Ioerrorenum_CantCreateBackup is a representation of the C type G_IO_ERROR_CANT_CREATE_BACKUP.
	Ioerrorenum_CantCreateBackup Ioerrorenum = 22
	// Ioerrorenum_WrongEtag is a representation of the C type G_IO_ERROR_WRONG_ETAG.
	Ioerrorenum_WrongEtag Ioerrorenum = 23
	// Ioerrorenum_TimedOut is a representation of the C type G_IO_ERROR_TIMED_OUT.
	Ioerrorenum_TimedOut Ioerrorenum = 24
	// Ioerrorenum_WouldRecurse is a representation of the C type G_IO_ERROR_WOULD_RECURSE.
	Ioerrorenum_WouldRecurse Ioerrorenum = 25
	// Ioerrorenum_Busy is a representation of the C type G_IO_ERROR_BUSY.
	Ioerrorenum_Busy Ioerrorenum = 26
	// Ioerrorenum_WouldBlock is a representation of the C type G_IO_ERROR_WOULD_BLOCK.
	Ioerrorenum_WouldBlock Ioerrorenum = 27
	// Ioerrorenum_HostNotFound is a representation of the C type G_IO_ERROR_HOST_NOT_FOUND.
	Ioerrorenum_HostNotFound Ioerrorenum = 28
	// Ioerrorenum_WouldMerge is a representation of the C type G_IO_ERROR_WOULD_MERGE.
	Ioerrorenum_WouldMerge Ioerrorenum = 29
	// Ioerrorenum_FailedHandled is a representation of the C type G_IO_ERROR_FAILED_HANDLED.
	Ioerrorenum_FailedHandled Ioerrorenum = 30
	// Ioerrorenum_TooManyOpenFiles is a representation of the C type G_IO_ERROR_TOO_MANY_OPEN_FILES.
	Ioerrorenum_TooManyOpenFiles Ioerrorenum = 31
	// Ioerrorenum_NotInitialized is a representation of the C type G_IO_ERROR_NOT_INITIALIZED.
	Ioerrorenum_NotInitialized Ioerrorenum = 32
	// Ioerrorenum_AddressInUse is a representation of the C type G_IO_ERROR_ADDRESS_IN_USE.
	Ioerrorenum_AddressInUse Ioerrorenum = 33
	// Ioerrorenum_PartialInput is a representation of the C type G_IO_ERROR_PARTIAL_INPUT.
	Ioerrorenum_PartialInput Ioerrorenum = 34
	// Ioerrorenum_InvalidData is a representation of the C type G_IO_ERROR_INVALID_DATA.
	Ioerrorenum_InvalidData Ioerrorenum = 35
	// Ioerrorenum_DbusError is a representation of the C type G_IO_ERROR_DBUS_ERROR.
	Ioerrorenum_DbusError Ioerrorenum = 36
	// Ioerrorenum_HostUnreachable is a representation of the C type G_IO_ERROR_HOST_UNREACHABLE.
	Ioerrorenum_HostUnreachable Ioerrorenum = 37
	// Ioerrorenum_NetworkUnreachable is a representation of the C type G_IO_ERROR_NETWORK_UNREACHABLE.
	Ioerrorenum_NetworkUnreachable Ioerrorenum = 38
	// Ioerrorenum_ConnectionRefused is a representation of the C type G_IO_ERROR_CONNECTION_REFUSED.
	Ioerrorenum_ConnectionRefused Ioerrorenum = 39
	// Ioerrorenum_ProxyFailed is a representation of the C type G_IO_ERROR_PROXY_FAILED.
	Ioerrorenum_ProxyFailed Ioerrorenum = 40
	// Ioerrorenum_ProxyAuthFailed is a representation of the C type G_IO_ERROR_PROXY_AUTH_FAILED.
	Ioerrorenum_ProxyAuthFailed Ioerrorenum = 41
	// Ioerrorenum_ProxyNeedAuth is a representation of the C type G_IO_ERROR_PROXY_NEED_AUTH.
	Ioerrorenum_ProxyNeedAuth Ioerrorenum = 42
	// Ioerrorenum_ProxyNotAllowed is a representation of the C type G_IO_ERROR_PROXY_NOT_ALLOWED.
	Ioerrorenum_ProxyNotAllowed Ioerrorenum = 43
	// Ioerrorenum_BrokenPipe is a representation of the C type G_IO_ERROR_BROKEN_PIPE.
	Ioerrorenum_BrokenPipe Ioerrorenum = 44
	// Ioerrorenum_ConnectionClosed is a representation of the C type G_IO_ERROR_CONNECTION_CLOSED.
	Ioerrorenum_ConnectionClosed Ioerrorenum = 44
	// Ioerrorenum_NotConnected is a representation of the C type G_IO_ERROR_NOT_CONNECTED.
	Ioerrorenum_NotConnected Ioerrorenum = 45
	// Ioerrorenum_MessageTooLarge is a representation of the C type G_IO_ERROR_MESSAGE_TOO_LARGE.
	Ioerrorenum_MessageTooLarge Ioerrorenum = 46
)

// Iomodulescopeflags is a representation of the C type GIOModuleScopeFlags.
//
// since 2.30
type Iomodulescopeflags int

const (
	// Iomodulescopeflags_None is a representation of the C type G_IO_MODULE_SCOPE_NONE.
	Iomodulescopeflags_None Iomodulescopeflags = 0
	// Iomodulescopeflags_BlockDuplicates is a representation of the C type G_IO_MODULE_SCOPE_BLOCK_DUPLICATES.
	Iomodulescopeflags_BlockDuplicates Iomodulescopeflags = 1
)

// Mountoperationresult is a representation of the C type GMountOperationResult.
type Mountoperationresult int

const (
	// Mountoperationresult_Handled is a representation of the C type G_MOUNT_OPERATION_HANDLED.
	Mountoperationresult_Handled Mountoperationresult = 0
	// Mountoperationresult_Aborted is a representation of the C type G_MOUNT_OPERATION_ABORTED.
	Mountoperationresult_Aborted Mountoperationresult = 1
	// Mountoperationresult_Unhandled is a representation of the C type G_MOUNT_OPERATION_UNHANDLED.
	Mountoperationresult_Unhandled Mountoperationresult = 2
)

// Networkconnectivity is a representation of the C type GNetworkConnectivity.
//
// since 2.44
type Networkconnectivity int

const (
	// Networkconnectivity_Local is a representation of the C type G_NETWORK_CONNECTIVITY_LOCAL.
	Networkconnectivity_Local Networkconnectivity = 1
	// Networkconnectivity_Limited is a representation of the C type G_NETWORK_CONNECTIVITY_LIMITED.
	Networkconnectivity_Limited Networkconnectivity = 2
	// Networkconnectivity_Portal is a representation of the C type G_NETWORK_CONNECTIVITY_PORTAL.
	Networkconnectivity_Portal Networkconnectivity = 3
	// Networkconnectivity_Full is a representation of the C type G_NETWORK_CONNECTIVITY_FULL.
	Networkconnectivity_Full Networkconnectivity = 4
)

// Notificationpriority is a representation of the C type GNotificationPriority.
//
// since 2.42
type Notificationpriority int

const (
	// Notificationpriority_Normal is a representation of the C type G_NOTIFICATION_PRIORITY_NORMAL.
	Notificationpriority_Normal Notificationpriority = 0
	// Notificationpriority_Low is a representation of the C type G_NOTIFICATION_PRIORITY_LOW.
	Notificationpriority_Low Notificationpriority = 1
	// Notificationpriority_High is a representation of the C type G_NOTIFICATION_PRIORITY_HIGH.
	Notificationpriority_High Notificationpriority = 2
	// Notificationpriority_Urgent is a representation of the C type G_NOTIFICATION_PRIORITY_URGENT.
	Notificationpriority_Urgent Notificationpriority = 3
)

// Passwordsave is a representation of the C type GPasswordSave.
type Passwordsave int

const (
	// Passwordsave_Never is a representation of the C type G_PASSWORD_SAVE_NEVER.
	Passwordsave_Never Passwordsave = 0
	// Passwordsave_ForSession is a representation of the C type G_PASSWORD_SAVE_FOR_SESSION.
	Passwordsave_ForSession Passwordsave = 1
	// Passwordsave_Permanently is a representation of the C type G_PASSWORD_SAVE_PERMANENTLY.
	Passwordsave_Permanently Passwordsave = 2
)

// Pollablereturn is a representation of the C type GPollableReturn.
//
// since 2.60
type Pollablereturn int

const (
	// Pollablereturn_Failed is a representation of the C type G_POLLABLE_RETURN_FAILED.
	Pollablereturn_Failed Pollablereturn = 0
	// Pollablereturn_Ok is a representation of the C type G_POLLABLE_RETURN_OK.
	Pollablereturn_Ok Pollablereturn = 1
	// Pollablereturn_WouldBlock is a representation of the C type G_POLLABLE_RETURN_WOULD_BLOCK.
	Pollablereturn_WouldBlock Pollablereturn = -27
)

// Resolvererror is a representation of the C type GResolverError.
//
// since 2.22
type Resolvererror int

const (
	// Resolvererror_NotFound is a representation of the C type G_RESOLVER_ERROR_NOT_FOUND.
	Resolvererror_NotFound Resolvererror = 0
	// Resolvererror_TemporaryFailure is a representation of the C type G_RESOLVER_ERROR_TEMPORARY_FAILURE.
	Resolvererror_TemporaryFailure Resolvererror = 1
	// Resolvererror_Internal is a representation of the C type G_RESOLVER_ERROR_INTERNAL.
	Resolvererror_Internal Resolvererror = 2
)

// Resolverrecordtype is a representation of the C type GResolverRecordType.
//
// since 2.34
type Resolverrecordtype int

const (
	// Resolverrecordtype_Srv is a representation of the C type G_RESOLVER_RECORD_SRV.
	Resolverrecordtype_Srv Resolverrecordtype = 1
	// Resolverrecordtype_Mx is a representation of the C type G_RESOLVER_RECORD_MX.
	Resolverrecordtype_Mx Resolverrecordtype = 2
	// Resolverrecordtype_Txt is a representation of the C type G_RESOLVER_RECORD_TXT.
	Resolverrecordtype_Txt Resolverrecordtype = 3
	// Resolverrecordtype_Soa is a representation of the C type G_RESOLVER_RECORD_SOA.
	Resolverrecordtype_Soa Resolverrecordtype = 4
	// Resolverrecordtype_Ns is a representation of the C type G_RESOLVER_RECORD_NS.
	Resolverrecordtype_Ns Resolverrecordtype = 5
)

// Resourceerror is a representation of the C type GResourceError.
//
// since 2.32
type Resourceerror int

const (
	// Resourceerror_NotFound is a representation of the C type G_RESOURCE_ERROR_NOT_FOUND.
	Resourceerror_NotFound Resourceerror = 0
	// Resourceerror_Internal is a representation of the C type G_RESOURCE_ERROR_INTERNAL.
	Resourceerror_Internal Resourceerror = 1
)

// Socketclientevent is a representation of the C type GSocketClientEvent.
//
// since 2.32
type Socketclientevent int

const (
	// Socketclientevent_Resolving is a representation of the C type G_SOCKET_CLIENT_RESOLVING.
	Socketclientevent_Resolving Socketclientevent = 0
	// Socketclientevent_Resolved is a representation of the C type G_SOCKET_CLIENT_RESOLVED.
	Socketclientevent_Resolved Socketclientevent = 1
	// Socketclientevent_Connecting is a representation of the C type G_SOCKET_CLIENT_CONNECTING.
	Socketclientevent_Connecting Socketclientevent = 2
	// Socketclientevent_Connected is a representation of the C type G_SOCKET_CLIENT_CONNECTED.
	Socketclientevent_Connected Socketclientevent = 3
	// Socketclientevent_ProxyNegotiating is a representation of the C type G_SOCKET_CLIENT_PROXY_NEGOTIATING.
	Socketclientevent_ProxyNegotiating Socketclientevent = 4
	// Socketclientevent_ProxyNegotiated is a representation of the C type G_SOCKET_CLIENT_PROXY_NEGOTIATED.
	Socketclientevent_ProxyNegotiated Socketclientevent = 5
	// Socketclientevent_TlsHandshaking is a representation of the C type G_SOCKET_CLIENT_TLS_HANDSHAKING.
	Socketclientevent_TlsHandshaking Socketclientevent = 6
	// Socketclientevent_TlsHandshaked is a representation of the C type G_SOCKET_CLIENT_TLS_HANDSHAKED.
	Socketclientevent_TlsHandshaked Socketclientevent = 7
	// Socketclientevent_Complete is a representation of the C type G_SOCKET_CLIENT_COMPLETE.
	Socketclientevent_Complete Socketclientevent = 8
)

// Socketfamily is a representation of the C type GSocketFamily.
//
// since 2.22
type Socketfamily int

const (
	// Socketfamily_Invalid is a representation of the C type G_SOCKET_FAMILY_INVALID.
	Socketfamily_Invalid Socketfamily = 0
	// Socketfamily_Unix is a representation of the C type G_SOCKET_FAMILY_UNIX.
	Socketfamily_Unix Socketfamily = 1
	// Socketfamily_Ipv4 is a representation of the C type G_SOCKET_FAMILY_IPV4.
	Socketfamily_Ipv4 Socketfamily = 2
	// Socketfamily_Ipv6 is a representation of the C type G_SOCKET_FAMILY_IPV6.
	Socketfamily_Ipv6 Socketfamily = 10
)

// Socketlistenerevent is a representation of the C type GSocketListenerEvent.
//
// since 2.46
type Socketlistenerevent int

const (
	// Socketlistenerevent_Binding is a representation of the C type G_SOCKET_LISTENER_BINDING.
	Socketlistenerevent_Binding Socketlistenerevent = 0
	// Socketlistenerevent_Bound is a representation of the C type G_SOCKET_LISTENER_BOUND.
	Socketlistenerevent_Bound Socketlistenerevent = 1
	// Socketlistenerevent_Listening is a representation of the C type G_SOCKET_LISTENER_LISTENING.
	Socketlistenerevent_Listening Socketlistenerevent = 2
	// Socketlistenerevent_Listened is a representation of the C type G_SOCKET_LISTENER_LISTENED.
	Socketlistenerevent_Listened Socketlistenerevent = 3
)

// Socketprotocol is a representation of the C type GSocketProtocol.
//
// since 2.22
type Socketprotocol int

const (
	// Socketprotocol_Unknown is a representation of the C type G_SOCKET_PROTOCOL_UNKNOWN.
	Socketprotocol_Unknown Socketprotocol = -1
	// Socketprotocol_Default is a representation of the C type G_SOCKET_PROTOCOL_DEFAULT.
	Socketprotocol_Default Socketprotocol = 0
	// Socketprotocol_Tcp is a representation of the C type G_SOCKET_PROTOCOL_TCP.
	Socketprotocol_Tcp Socketprotocol = 6
	// Socketprotocol_Udp is a representation of the C type G_SOCKET_PROTOCOL_UDP.
	Socketprotocol_Udp Socketprotocol = 17
	// Socketprotocol_Sctp is a representation of the C type G_SOCKET_PROTOCOL_SCTP.
	Socketprotocol_Sctp Socketprotocol = 132
)

// Sockettype is a representation of the C type GSocketType.
//
// since 2.22
type Sockettype int

const (
	// Sockettype_Invalid is a representation of the C type G_SOCKET_TYPE_INVALID.
	Sockettype_Invalid Sockettype = 0
	// Sockettype_Stream is a representation of the C type G_SOCKET_TYPE_STREAM.
	Sockettype_Stream Sockettype = 1
	// Sockettype_Datagram is a representation of the C type G_SOCKET_TYPE_DATAGRAM.
	Sockettype_Datagram Sockettype = 2
	// Sockettype_Seqpacket is a representation of the C type G_SOCKET_TYPE_SEQPACKET.
	Sockettype_Seqpacket Sockettype = 3
)

// Tlsauthenticationmode is a representation of the C type GTlsAuthenticationMode.
//
// since 2.28
type Tlsauthenticationmode int

const (
	// Tlsauthenticationmode_None is a representation of the C type G_TLS_AUTHENTICATION_NONE.
	Tlsauthenticationmode_None Tlsauthenticationmode = 0
	// Tlsauthenticationmode_Requested is a representation of the C type G_TLS_AUTHENTICATION_REQUESTED.
	Tlsauthenticationmode_Requested Tlsauthenticationmode = 1
	// Tlsauthenticationmode_Required is a representation of the C type G_TLS_AUTHENTICATION_REQUIRED.
	Tlsauthenticationmode_Required Tlsauthenticationmode = 2
)

// Tlscertificaterequestflags is a representation of the C type GTlsCertificateRequestFlags.
//
// since 2.40
type Tlscertificaterequestflags int

const (
	// Tlscertificaterequestflags_None is a representation of the C type G_TLS_CERTIFICATE_REQUEST_NONE.
	Tlscertificaterequestflags_None Tlscertificaterequestflags = 0
)

// Tlsdatabaselookupflags is a representation of the C type GTlsDatabaseLookupFlags.
//
// since 2.30
type Tlsdatabaselookupflags int

const (
	// Tlsdatabaselookupflags_None is a representation of the C type G_TLS_DATABASE_LOOKUP_NONE.
	Tlsdatabaselookupflags_None Tlsdatabaselookupflags = 0
	// Tlsdatabaselookupflags_Keypair is a representation of the C type G_TLS_DATABASE_LOOKUP_KEYPAIR.
	Tlsdatabaselookupflags_Keypair Tlsdatabaselookupflags = 1
)

// Tlserror is a representation of the C type GTlsError.
//
// since 2.28
type Tlserror int

const (
	// Tlserror_Unavailable is a representation of the C type G_TLS_ERROR_UNAVAILABLE.
	Tlserror_Unavailable Tlserror = 0
	// Tlserror_Misc is a representation of the C type G_TLS_ERROR_MISC.
	Tlserror_Misc Tlserror = 1
	// Tlserror_BadCertificate is a representation of the C type G_TLS_ERROR_BAD_CERTIFICATE.
	Tlserror_BadCertificate Tlserror = 2
	// Tlserror_NotTls is a representation of the C type G_TLS_ERROR_NOT_TLS.
	Tlserror_NotTls Tlserror = 3
	// Tlserror_Handshake is a representation of the C type G_TLS_ERROR_HANDSHAKE.
	Tlserror_Handshake Tlserror = 4
	// Tlserror_CertificateRequired is a representation of the C type G_TLS_ERROR_CERTIFICATE_REQUIRED.
	Tlserror_CertificateRequired Tlserror = 5
	// Tlserror_Eof is a representation of the C type G_TLS_ERROR_EOF.
	Tlserror_Eof Tlserror = 6
	// Tlserror_InappropriateFallback is a representation of the C type G_TLS_ERROR_INAPPROPRIATE_FALLBACK.
	Tlserror_InappropriateFallback Tlserror = 7
)

// Tlsinteractionresult is a representation of the C type GTlsInteractionResult.
//
// since 2.30
type Tlsinteractionresult int

const (
	// Tlsinteractionresult_Unhandled is a representation of the C type G_TLS_INTERACTION_UNHANDLED.
	Tlsinteractionresult_Unhandled Tlsinteractionresult = 0
	// Tlsinteractionresult_Handled is a representation of the C type G_TLS_INTERACTION_HANDLED.
	Tlsinteractionresult_Handled Tlsinteractionresult = 1
	// Tlsinteractionresult_Failed is a representation of the C type G_TLS_INTERACTION_FAILED.
	Tlsinteractionresult_Failed Tlsinteractionresult = 2
)

// Tlsrehandshakemode is a representation of the C type GTlsRehandshakeMode.
//
// since 2.28
type Tlsrehandshakemode int

const (
	// Tlsrehandshakemode_Never is a representation of the C type G_TLS_REHANDSHAKE_NEVER.
	Tlsrehandshakemode_Never Tlsrehandshakemode = 0
	// Tlsrehandshakemode_Safely is a representation of the C type G_TLS_REHANDSHAKE_SAFELY.
	Tlsrehandshakemode_Safely Tlsrehandshakemode = 1
	// Tlsrehandshakemode_Unsafely is a representation of the C type G_TLS_REHANDSHAKE_UNSAFELY.
	Tlsrehandshakemode_Unsafely Tlsrehandshakemode = 2
)

// Unixsocketaddresstype is a representation of the C type GUnixSocketAddressType.
//
// since 2.26
type Unixsocketaddresstype int

const (
	// Unixsocketaddresstype_Invalid is a representation of the C type G_UNIX_SOCKET_ADDRESS_INVALID.
	Unixsocketaddresstype_Invalid Unixsocketaddresstype = 0
	// Unixsocketaddresstype_Anonymous is a representation of the C type G_UNIX_SOCKET_ADDRESS_ANONYMOUS.
	Unixsocketaddresstype_Anonymous Unixsocketaddresstype = 1
	// Unixsocketaddresstype_Path is a representation of the C type G_UNIX_SOCKET_ADDRESS_PATH.
	Unixsocketaddresstype_Path Unixsocketaddresstype = 2
	// Unixsocketaddresstype_Abstract is a representation of the C type G_UNIX_SOCKET_ADDRESS_ABSTRACT.
	Unixsocketaddresstype_Abstract Unixsocketaddresstype = 3
	// Unixsocketaddresstype_AbstractPadded is a representation of the C type G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
	Unixsocketaddresstype_AbstractPadded Unixsocketaddresstype = 4
)

// Zlibcompressorformat is a representation of the C type GZlibCompressorFormat.
//
// since 2.24
type Zlibcompressorformat int

const (
	// Zlibcompressorformat_Zlib is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_ZLIB.
	Zlibcompressorformat_Zlib Zlibcompressorformat = 0
	// Zlibcompressorformat_Gzip is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_GZIP.
	Zlibcompressorformat_Gzip Zlibcompressorformat = 1
	// Zlibcompressorformat_Raw is a representation of the C type G_ZLIB_COMPRESSOR_FORMAT_RAW.
	Zlibcompressorformat_Raw Zlibcompressorformat = 2
)
