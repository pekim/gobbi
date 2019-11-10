// Code generated - DO NOT EDIT.

package gio

type BusType int32

const (
	BusType_Starter BusType = int64(-1)
	BusType_None    BusType = int64(0)
	BusType_System  BusType = int64(1)
	BusType_Session BusType = int64(2)
)

type ConverterResult uint32

const (
	ConverterResult_Error     ConverterResult = int64(0)
	ConverterResult_Converted ConverterResult = int64(1)
	ConverterResult_Finished  ConverterResult = int64(2)
	ConverterResult_Flushed   ConverterResult = int64(3)
)

type CredentialsType uint32

const (
	CredentialsType_Invalid             CredentialsType = int64(0)
	CredentialsType_LinuxUcred          CredentialsType = int64(1)
	CredentialsType_FreebsdCmsgcred     CredentialsType = int64(2)
	CredentialsType_OpenbsdSockpeercred CredentialsType = int64(3)
	CredentialsType_SolarisUcred        CredentialsType = int64(4)
	CredentialsType_NetbsdUnpcbid       CredentialsType = int64(5)
)

type DBusError uint32

const (
	DBusError_Failed                        DBusError = int64(0)
	DBusError_NoMemory                      DBusError = int64(1)
	DBusError_ServiceUnknown                DBusError = int64(2)
	DBusError_NameHasNoOwner                DBusError = int64(3)
	DBusError_NoReply                       DBusError = int64(4)
	DBusError_IoError                       DBusError = int64(5)
	DBusError_BadAddress                    DBusError = int64(6)
	DBusError_NotSupported                  DBusError = int64(7)
	DBusError_LimitsExceeded                DBusError = int64(8)
	DBusError_AccessDenied                  DBusError = int64(9)
	DBusError_AuthFailed                    DBusError = int64(10)
	DBusError_NoServer                      DBusError = int64(11)
	DBusError_Timeout                       DBusError = int64(12)
	DBusError_NoNetwork                     DBusError = int64(13)
	DBusError_AddressInUse                  DBusError = int64(14)
	DBusError_Disconnected                  DBusError = int64(15)
	DBusError_InvalidArgs                   DBusError = int64(16)
	DBusError_FileNotFound                  DBusError = int64(17)
	DBusError_FileExists                    DBusError = int64(18)
	DBusError_UnknownMethod                 DBusError = int64(19)
	DBusError_TimedOut                      DBusError = int64(20)
	DBusError_MatchRuleNotFound             DBusError = int64(21)
	DBusError_MatchRuleInvalid              DBusError = int64(22)
	DBusError_SpawnExecFailed               DBusError = int64(23)
	DBusError_SpawnForkFailed               DBusError = int64(24)
	DBusError_SpawnChildExited              DBusError = int64(25)
	DBusError_SpawnChildSignaled            DBusError = int64(26)
	DBusError_SpawnFailed                   DBusError = int64(27)
	DBusError_SpawnSetupFailed              DBusError = int64(28)
	DBusError_SpawnConfigInvalid            DBusError = int64(29)
	DBusError_SpawnServiceInvalid           DBusError = int64(30)
	DBusError_SpawnServiceNotFound          DBusError = int64(31)
	DBusError_SpawnPermissionsInvalid       DBusError = int64(32)
	DBusError_SpawnFileInvalid              DBusError = int64(33)
	DBusError_SpawnNoMemory                 DBusError = int64(34)
	DBusError_UnixProcessIdUnknown          DBusError = int64(35)
	DBusError_InvalidSignature              DBusError = int64(36)
	DBusError_InvalidFileContent            DBusError = int64(37)
	DBusError_SelinuxSecurityContextUnknown DBusError = int64(38)
	DBusError_AdtAuditDataUnknown           DBusError = int64(39)
	DBusError_ObjectPathInUse               DBusError = int64(40)
	DBusError_UnknownObject                 DBusError = int64(41)
	DBusError_UnknownInterface              DBusError = int64(42)
	DBusError_UnknownProperty               DBusError = int64(43)
	DBusError_PropertyReadOnly              DBusError = int64(44)
)

type DBusMessageByteOrder uint32

const (
	DBusMessageByteOrder_BigEndian    DBusMessageByteOrder = int64(66)
	DBusMessageByteOrder_LittleEndian DBusMessageByteOrder = int64(108)
)

type DBusMessageHeaderField uint32

const (
	DBusMessageHeaderField_Invalid     DBusMessageHeaderField = int64(0)
	DBusMessageHeaderField_Path        DBusMessageHeaderField = int64(1)
	DBusMessageHeaderField_Interface   DBusMessageHeaderField = int64(2)
	DBusMessageHeaderField_Member      DBusMessageHeaderField = int64(3)
	DBusMessageHeaderField_ErrorName   DBusMessageHeaderField = int64(4)
	DBusMessageHeaderField_ReplySerial DBusMessageHeaderField = int64(5)
	DBusMessageHeaderField_Destination DBusMessageHeaderField = int64(6)
	DBusMessageHeaderField_Sender      DBusMessageHeaderField = int64(7)
	DBusMessageHeaderField_Signature   DBusMessageHeaderField = int64(8)
	DBusMessageHeaderField_NumUnixFds  DBusMessageHeaderField = int64(9)
)

type DBusMessageType uint32

const (
	DBusMessageType_Invalid      DBusMessageType = int64(0)
	DBusMessageType_MethodCall   DBusMessageType = int64(1)
	DBusMessageType_MethodReturn DBusMessageType = int64(2)
	DBusMessageType_Error        DBusMessageType = int64(3)
	DBusMessageType_Signal       DBusMessageType = int64(4)
)

type DataStreamByteOrder uint32

const (
	DataStreamByteOrder_BigEndian    DataStreamByteOrder = int64(0)
	DataStreamByteOrder_LittleEndian DataStreamByteOrder = int64(1)
	DataStreamByteOrder_HostEndian   DataStreamByteOrder = int64(2)
)

type DataStreamNewlineType uint32

const (
	DataStreamNewlineType_Lf   DataStreamNewlineType = int64(0)
	DataStreamNewlineType_Cr   DataStreamNewlineType = int64(1)
	DataStreamNewlineType_CrLf DataStreamNewlineType = int64(2)
	DataStreamNewlineType_Any  DataStreamNewlineType = int64(3)
)

type DriveStartStopType uint32

const (
	DriveStartStopType_Unknown   DriveStartStopType = int64(0)
	DriveStartStopType_Shutdown  DriveStartStopType = int64(1)
	DriveStartStopType_Network   DriveStartStopType = int64(2)
	DriveStartStopType_Multidisk DriveStartStopType = int64(3)
	DriveStartStopType_Password  DriveStartStopType = int64(4)
)

type EmblemOrigin uint32

const (
	EmblemOrigin_Unknown      EmblemOrigin = int64(0)
	EmblemOrigin_Device       EmblemOrigin = int64(1)
	EmblemOrigin_Livemetadata EmblemOrigin = int64(2)
	EmblemOrigin_Tag          EmblemOrigin = int64(3)
)

type FileAttributeStatus uint32

const (
	FileAttributeStatus_Unset        FileAttributeStatus = int64(0)
	FileAttributeStatus_Set          FileAttributeStatus = int64(1)
	FileAttributeStatus_ErrorSetting FileAttributeStatus = int64(2)
)

type FileAttributeType uint32

const (
	FileAttributeType_Invalid    FileAttributeType = int64(0)
	FileAttributeType_String     FileAttributeType = int64(1)
	FileAttributeType_ByteString FileAttributeType = int64(2)
	FileAttributeType_Boolean    FileAttributeType = int64(3)
	FileAttributeType_Uint32     FileAttributeType = int64(4)
	FileAttributeType_Int32      FileAttributeType = int64(5)
	FileAttributeType_Uint64     FileAttributeType = int64(6)
	FileAttributeType_Int64      FileAttributeType = int64(7)
	FileAttributeType_Object     FileAttributeType = int64(8)
	FileAttributeType_Stringv    FileAttributeType = int64(9)
)

type FileMonitorEvent uint32

const (
	FileMonitorEvent_Changed          FileMonitorEvent = int64(0)
	FileMonitorEvent_ChangesDoneHint  FileMonitorEvent = int64(1)
	FileMonitorEvent_Deleted          FileMonitorEvent = int64(2)
	FileMonitorEvent_Created          FileMonitorEvent = int64(3)
	FileMonitorEvent_AttributeChanged FileMonitorEvent = int64(4)
	FileMonitorEvent_PreUnmount       FileMonitorEvent = int64(5)
	FileMonitorEvent_Unmounted        FileMonitorEvent = int64(6)
	FileMonitorEvent_Moved            FileMonitorEvent = int64(7)
	FileMonitorEvent_Renamed          FileMonitorEvent = int64(8)
	FileMonitorEvent_MovedIn          FileMonitorEvent = int64(9)
	FileMonitorEvent_MovedOut         FileMonitorEvent = int64(10)
)

type FileType uint32

const (
	FileType_Unknown      FileType = int64(0)
	FileType_Regular      FileType = int64(1)
	FileType_Directory    FileType = int64(2)
	FileType_SymbolicLink FileType = int64(3)
	FileType_Special      FileType = int64(4)
	FileType_Shortcut     FileType = int64(5)
	FileType_Mountable    FileType = int64(6)
)

type FilesystemPreviewType uint32

const (
	FilesystemPreviewType_IfAlways FilesystemPreviewType = int64(0)
	FilesystemPreviewType_IfLocal  FilesystemPreviewType = int64(1)
	FilesystemPreviewType_Never    FilesystemPreviewType = int64(2)
)

type IOErrorEnum uint32

const (
	IOErrorEnum_Failed             IOErrorEnum = int64(0)
	IOErrorEnum_NotFound           IOErrorEnum = int64(1)
	IOErrorEnum_Exists             IOErrorEnum = int64(2)
	IOErrorEnum_IsDirectory        IOErrorEnum = int64(3)
	IOErrorEnum_NotDirectory       IOErrorEnum = int64(4)
	IOErrorEnum_NotEmpty           IOErrorEnum = int64(5)
	IOErrorEnum_NotRegularFile     IOErrorEnum = int64(6)
	IOErrorEnum_NotSymbolicLink    IOErrorEnum = int64(7)
	IOErrorEnum_NotMountableFile   IOErrorEnum = int64(8)
	IOErrorEnum_FilenameTooLong    IOErrorEnum = int64(9)
	IOErrorEnum_InvalidFilename    IOErrorEnum = int64(10)
	IOErrorEnum_TooManyLinks       IOErrorEnum = int64(11)
	IOErrorEnum_NoSpace            IOErrorEnum = int64(12)
	IOErrorEnum_InvalidArgument    IOErrorEnum = int64(13)
	IOErrorEnum_PermissionDenied   IOErrorEnum = int64(14)
	IOErrorEnum_NotSupported       IOErrorEnum = int64(15)
	IOErrorEnum_NotMounted         IOErrorEnum = int64(16)
	IOErrorEnum_AlreadyMounted     IOErrorEnum = int64(17)
	IOErrorEnum_Closed             IOErrorEnum = int64(18)
	IOErrorEnum_Cancelled          IOErrorEnum = int64(19)
	IOErrorEnum_Pending            IOErrorEnum = int64(20)
	IOErrorEnum_ReadOnly           IOErrorEnum = int64(21)
	IOErrorEnum_CantCreateBackup   IOErrorEnum = int64(22)
	IOErrorEnum_WrongEtag          IOErrorEnum = int64(23)
	IOErrorEnum_TimedOut           IOErrorEnum = int64(24)
	IOErrorEnum_WouldRecurse       IOErrorEnum = int64(25)
	IOErrorEnum_Busy               IOErrorEnum = int64(26)
	IOErrorEnum_WouldBlock         IOErrorEnum = int64(27)
	IOErrorEnum_HostNotFound       IOErrorEnum = int64(28)
	IOErrorEnum_WouldMerge         IOErrorEnum = int64(29)
	IOErrorEnum_FailedHandled      IOErrorEnum = int64(30)
	IOErrorEnum_TooManyOpenFiles   IOErrorEnum = int64(31)
	IOErrorEnum_NotInitialized     IOErrorEnum = int64(32)
	IOErrorEnum_AddressInUse       IOErrorEnum = int64(33)
	IOErrorEnum_PartialInput       IOErrorEnum = int64(34)
	IOErrorEnum_InvalidData        IOErrorEnum = int64(35)
	IOErrorEnum_DbusError          IOErrorEnum = int64(36)
	IOErrorEnum_HostUnreachable    IOErrorEnum = int64(37)
	IOErrorEnum_NetworkUnreachable IOErrorEnum = int64(38)
	IOErrorEnum_ConnectionRefused  IOErrorEnum = int64(39)
	IOErrorEnum_ProxyFailed        IOErrorEnum = int64(40)
	IOErrorEnum_ProxyAuthFailed    IOErrorEnum = int64(41)
	IOErrorEnum_ProxyNeedAuth      IOErrorEnum = int64(42)
	IOErrorEnum_ProxyNotAllowed    IOErrorEnum = int64(43)
	IOErrorEnum_BrokenPipe         IOErrorEnum = int64(44)
	IOErrorEnum_ConnectionClosed   IOErrorEnum = int64(44)
	IOErrorEnum_NotConnected       IOErrorEnum = int64(45)
	IOErrorEnum_MessageTooLarge    IOErrorEnum = int64(46)
)

type IOModuleScopeFlags uint32

const (
	IOModuleScopeFlags_None            IOModuleScopeFlags = int64(0)
	IOModuleScopeFlags_BlockDuplicates IOModuleScopeFlags = int64(1)
)

type MountOperationResult uint32

const (
	MountOperationResult_Handled   MountOperationResult = int64(0)
	MountOperationResult_Aborted   MountOperationResult = int64(1)
	MountOperationResult_Unhandled MountOperationResult = int64(2)
)

type NetworkConnectivity uint32

const (
	NetworkConnectivity_Local   NetworkConnectivity = int64(1)
	NetworkConnectivity_Limited NetworkConnectivity = int64(2)
	NetworkConnectivity_Portal  NetworkConnectivity = int64(3)
	NetworkConnectivity_Full    NetworkConnectivity = int64(4)
)

type NotificationPriority uint32

const (
	NotificationPriority_Normal NotificationPriority = int64(0)
	NotificationPriority_Low    NotificationPriority = int64(1)
	NotificationPriority_High   NotificationPriority = int64(2)
	NotificationPriority_Urgent NotificationPriority = int64(3)
)

type PasswordSave uint32

const (
	PasswordSave_Never       PasswordSave = int64(0)
	PasswordSave_ForSession  PasswordSave = int64(1)
	PasswordSave_Permanently PasswordSave = int64(2)
)

type ResolverError uint32

const (
	ResolverError_NotFound         ResolverError = int64(0)
	ResolverError_TemporaryFailure ResolverError = int64(1)
	ResolverError_Internal         ResolverError = int64(2)
)

type ResolverRecordType uint32

const (
	ResolverRecordType_Srv ResolverRecordType = int64(1)
	ResolverRecordType_Mx  ResolverRecordType = int64(2)
	ResolverRecordType_Txt ResolverRecordType = int64(3)
	ResolverRecordType_Soa ResolverRecordType = int64(4)
	ResolverRecordType_Ns  ResolverRecordType = int64(5)
)

type ResourceError uint32

const (
	ResourceError_NotFound ResourceError = int64(0)
	ResourceError_Internal ResourceError = int64(1)
)

type SocketClientEvent uint32

const (
	SocketClientEvent_Resolving        SocketClientEvent = int64(0)
	SocketClientEvent_Resolved         SocketClientEvent = int64(1)
	SocketClientEvent_Connecting       SocketClientEvent = int64(2)
	SocketClientEvent_Connected        SocketClientEvent = int64(3)
	SocketClientEvent_ProxyNegotiating SocketClientEvent = int64(4)
	SocketClientEvent_ProxyNegotiated  SocketClientEvent = int64(5)
	SocketClientEvent_TlsHandshaking   SocketClientEvent = int64(6)
	SocketClientEvent_TlsHandshaked    SocketClientEvent = int64(7)
	SocketClientEvent_Complete         SocketClientEvent = int64(8)
)

type SocketFamily uint32

const (
	SocketFamily_Invalid SocketFamily = int64(0)
	SocketFamily_Unix    SocketFamily = int64(1)
	SocketFamily_Ipv4    SocketFamily = int64(2)
	SocketFamily_Ipv6    SocketFamily = int64(10)
)

type SocketListenerEvent uint32

const (
	SocketListenerEvent_Binding   SocketListenerEvent = int64(0)
	SocketListenerEvent_Bound     SocketListenerEvent = int64(1)
	SocketListenerEvent_Listening SocketListenerEvent = int64(2)
	SocketListenerEvent_Listened  SocketListenerEvent = int64(3)
)

type SocketProtocol int32

const (
	SocketProtocol_Unknown SocketProtocol = int64(-1)
	SocketProtocol_Default SocketProtocol = int64(0)
	SocketProtocol_Tcp     SocketProtocol = int64(6)
	SocketProtocol_Udp     SocketProtocol = int64(17)
	SocketProtocol_Sctp    SocketProtocol = int64(132)
)

type SocketType uint32

const (
	SocketType_Invalid   SocketType = int64(0)
	SocketType_Stream    SocketType = int64(1)
	SocketType_Datagram  SocketType = int64(2)
	SocketType_Seqpacket SocketType = int64(3)
)

type TlsAuthenticationMode uint32

const (
	TlsAuthenticationMode_None      TlsAuthenticationMode = int64(0)
	TlsAuthenticationMode_Requested TlsAuthenticationMode = int64(1)
	TlsAuthenticationMode_Required  TlsAuthenticationMode = int64(2)
)

type TlsCertificateRequestFlags uint32

const (
	TlsCertificateRequestFlags_None TlsCertificateRequestFlags = int64(0)
)

type TlsDatabaseLookupFlags uint32

const (
	TlsDatabaseLookupFlags_None    TlsDatabaseLookupFlags = int64(0)
	TlsDatabaseLookupFlags_Keypair TlsDatabaseLookupFlags = int64(1)
)

type TlsError uint32

const (
	TlsError_Unavailable         TlsError = int64(0)
	TlsError_Misc                TlsError = int64(1)
	TlsError_BadCertificate      TlsError = int64(2)
	TlsError_NotTls              TlsError = int64(3)
	TlsError_Handshake           TlsError = int64(4)
	TlsError_CertificateRequired TlsError = int64(5)
	TlsError_Eof                 TlsError = int64(6)
)

type TlsInteractionResult uint32

const (
	TlsInteractionResult_Unhandled TlsInteractionResult = int64(0)
	TlsInteractionResult_Handled   TlsInteractionResult = int64(1)
	TlsInteractionResult_Failed    TlsInteractionResult = int64(2)
)

type TlsRehandshakeMode uint32

const (
	TlsRehandshakeMode_Never    TlsRehandshakeMode = int64(0)
	TlsRehandshakeMode_Safely   TlsRehandshakeMode = int64(1)
	TlsRehandshakeMode_Unsafely TlsRehandshakeMode = int64(2)
)

type UnixSocketAddressType uint32

const (
	UnixSocketAddressType_Invalid        UnixSocketAddressType = int64(0)
	UnixSocketAddressType_Anonymous      UnixSocketAddressType = int64(1)
	UnixSocketAddressType_Path           UnixSocketAddressType = int64(2)
	UnixSocketAddressType_Abstract       UnixSocketAddressType = int64(3)
	UnixSocketAddressType_AbstractPadded UnixSocketAddressType = int64(4)
)

type ZlibCompressorFormat uint32

const (
	ZlibCompressorFormat_Zlib ZlibCompressorFormat = int64(0)
	ZlibCompressorFormat_Gzip ZlibCompressorFormat = int64(1)
	ZlibCompressorFormat_Raw  ZlibCompressorFormat = int64(2)
)
