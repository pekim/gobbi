// Code generated - DO NOT EDIT.

package gio

// Bustype is a representation of the C type BusType.
//
// since 2.26
type Bustype int

const (
	Bustype_Starter Bustype = -1
	Bustype_None    Bustype = 0
	Bustype_System  Bustype = 1
	Bustype_Session Bustype = 2
)

// Converterresult is a representation of the C type ConverterResult.
//
// since 2.24
type Converterresult int

const (
	Converterresult_Error     Converterresult = 0
	Converterresult_Converted Converterresult = 1
	Converterresult_Finished  Converterresult = 2
	Converterresult_Flushed   Converterresult = 3
)

// Credentialstype is a representation of the C type CredentialsType.
//
// since 2.26
type Credentialstype int

const (
	Credentialstype_Invalid             Credentialstype = 0
	Credentialstype_LinuxUcred          Credentialstype = 1
	Credentialstype_FreebsdCmsgcred     Credentialstype = 2
	Credentialstype_OpenbsdSockpeercred Credentialstype = 3
	Credentialstype_SolarisUcred        Credentialstype = 4
	Credentialstype_NetbsdUnpcbid       Credentialstype = 5
)

// Dbuserror is a representation of the C type DBusError.
//
// since 2.26
type Dbuserror int

const (
	Dbuserror_Failed                        Dbuserror = 0
	Dbuserror_NoMemory                      Dbuserror = 1
	Dbuserror_ServiceUnknown                Dbuserror = 2
	Dbuserror_NameHasNoOwner                Dbuserror = 3
	Dbuserror_NoReply                       Dbuserror = 4
	Dbuserror_IoError                       Dbuserror = 5
	Dbuserror_BadAddress                    Dbuserror = 6
	Dbuserror_NotSupported                  Dbuserror = 7
	Dbuserror_LimitsExceeded                Dbuserror = 8
	Dbuserror_AccessDenied                  Dbuserror = 9
	Dbuserror_AuthFailed                    Dbuserror = 10
	Dbuserror_NoServer                      Dbuserror = 11
	Dbuserror_Timeout                       Dbuserror = 12
	Dbuserror_NoNetwork                     Dbuserror = 13
	Dbuserror_AddressInUse                  Dbuserror = 14
	Dbuserror_Disconnected                  Dbuserror = 15
	Dbuserror_InvalidArgs                   Dbuserror = 16
	Dbuserror_FileNotFound                  Dbuserror = 17
	Dbuserror_FileExists                    Dbuserror = 18
	Dbuserror_UnknownMethod                 Dbuserror = 19
	Dbuserror_TimedOut                      Dbuserror = 20
	Dbuserror_MatchRuleNotFound             Dbuserror = 21
	Dbuserror_MatchRuleInvalid              Dbuserror = 22
	Dbuserror_SpawnExecFailed               Dbuserror = 23
	Dbuserror_SpawnForkFailed               Dbuserror = 24
	Dbuserror_SpawnChildExited              Dbuserror = 25
	Dbuserror_SpawnChildSignaled            Dbuserror = 26
	Dbuserror_SpawnFailed                   Dbuserror = 27
	Dbuserror_SpawnSetupFailed              Dbuserror = 28
	Dbuserror_SpawnConfigInvalid            Dbuserror = 29
	Dbuserror_SpawnServiceInvalid           Dbuserror = 30
	Dbuserror_SpawnServiceNotFound          Dbuserror = 31
	Dbuserror_SpawnPermissionsInvalid       Dbuserror = 32
	Dbuserror_SpawnFileInvalid              Dbuserror = 33
	Dbuserror_SpawnNoMemory                 Dbuserror = 34
	Dbuserror_UnixProcessIdUnknown          Dbuserror = 35
	Dbuserror_InvalidSignature              Dbuserror = 36
	Dbuserror_InvalidFileContent            Dbuserror = 37
	Dbuserror_SelinuxSecurityContextUnknown Dbuserror = 38
	Dbuserror_AdtAuditDataUnknown           Dbuserror = 39
	Dbuserror_ObjectPathInUse               Dbuserror = 40
	Dbuserror_UnknownObject                 Dbuserror = 41
	Dbuserror_UnknownInterface              Dbuserror = 42
	Dbuserror_UnknownProperty               Dbuserror = 43
	Dbuserror_PropertyReadOnly              Dbuserror = 44
)

// Dbusmessagebyteorder is a representation of the C type DBusMessageByteOrder.
//
// since 2.26
type Dbusmessagebyteorder int

const (
	Dbusmessagebyteorder_BigEndian    Dbusmessagebyteorder = 66
	Dbusmessagebyteorder_LittleEndian Dbusmessagebyteorder = 108
)

// Dbusmessageheaderfield is a representation of the C type DBusMessageHeaderField.
//
// since 2.26
type Dbusmessageheaderfield int

const (
	Dbusmessageheaderfield_Invalid     Dbusmessageheaderfield = 0
	Dbusmessageheaderfield_Path        Dbusmessageheaderfield = 1
	Dbusmessageheaderfield_Interface   Dbusmessageheaderfield = 2
	Dbusmessageheaderfield_Member      Dbusmessageheaderfield = 3
	Dbusmessageheaderfield_ErrorName   Dbusmessageheaderfield = 4
	Dbusmessageheaderfield_ReplySerial Dbusmessageheaderfield = 5
	Dbusmessageheaderfield_Destination Dbusmessageheaderfield = 6
	Dbusmessageheaderfield_Sender      Dbusmessageheaderfield = 7
	Dbusmessageheaderfield_Signature   Dbusmessageheaderfield = 8
	Dbusmessageheaderfield_NumUnixFds  Dbusmessageheaderfield = 9
)

// Dbusmessagetype is a representation of the C type DBusMessageType.
//
// since 2.26
type Dbusmessagetype int

const (
	Dbusmessagetype_Invalid      Dbusmessagetype = 0
	Dbusmessagetype_MethodCall   Dbusmessagetype = 1
	Dbusmessagetype_MethodReturn Dbusmessagetype = 2
	Dbusmessagetype_Error        Dbusmessagetype = 3
	Dbusmessagetype_Signal       Dbusmessagetype = 4
)

// Datastreambyteorder is a representation of the C type DataStreamByteOrder.
type Datastreambyteorder int

const (
	Datastreambyteorder_BigEndian    Datastreambyteorder = 0
	Datastreambyteorder_LittleEndian Datastreambyteorder = 1
	Datastreambyteorder_HostEndian   Datastreambyteorder = 2
)

// Datastreamnewlinetype is a representation of the C type DataStreamNewlineType.
type Datastreamnewlinetype int

const (
	Datastreamnewlinetype_Lf   Datastreamnewlinetype = 0
	Datastreamnewlinetype_Cr   Datastreamnewlinetype = 1
	Datastreamnewlinetype_CrLf Datastreamnewlinetype = 2
	Datastreamnewlinetype_Any  Datastreamnewlinetype = 3
)

// Drivestartstoptype is a representation of the C type DriveStartStopType.
//
// since 2.22
type Drivestartstoptype int

const (
	Drivestartstoptype_Unknown   Drivestartstoptype = 0
	Drivestartstoptype_Shutdown  Drivestartstoptype = 1
	Drivestartstoptype_Network   Drivestartstoptype = 2
	Drivestartstoptype_Multidisk Drivestartstoptype = 3
	Drivestartstoptype_Password  Drivestartstoptype = 4
)

// Emblemorigin is a representation of the C type EmblemOrigin.
//
// since 2.18
type Emblemorigin int

const (
	Emblemorigin_Unknown      Emblemorigin = 0
	Emblemorigin_Device       Emblemorigin = 1
	Emblemorigin_Livemetadata Emblemorigin = 2
	Emblemorigin_Tag          Emblemorigin = 3
)

// Fileattributestatus is a representation of the C type FileAttributeStatus.
type Fileattributestatus int

const (
	Fileattributestatus_Unset        Fileattributestatus = 0
	Fileattributestatus_Set          Fileattributestatus = 1
	Fileattributestatus_ErrorSetting Fileattributestatus = 2
)

// Fileattributetype is a representation of the C type FileAttributeType.
type Fileattributetype int

const (
	Fileattributetype_Invalid    Fileattributetype = 0
	Fileattributetype_String     Fileattributetype = 1
	Fileattributetype_ByteString Fileattributetype = 2
	Fileattributetype_Boolean    Fileattributetype = 3
	Fileattributetype_Uint32     Fileattributetype = 4
	Fileattributetype_Int32      Fileattributetype = 5
	Fileattributetype_Uint64     Fileattributetype = 6
	Fileattributetype_Int64      Fileattributetype = 7
	Fileattributetype_Object     Fileattributetype = 8
	Fileattributetype_Stringv    Fileattributetype = 9
)

// Filemonitorevent is a representation of the C type FileMonitorEvent.
type Filemonitorevent int

const (
	Filemonitorevent_Changed          Filemonitorevent = 0
	Filemonitorevent_ChangesDoneHint  Filemonitorevent = 1
	Filemonitorevent_Deleted          Filemonitorevent = 2
	Filemonitorevent_Created          Filemonitorevent = 3
	Filemonitorevent_AttributeChanged Filemonitorevent = 4
	Filemonitorevent_PreUnmount       Filemonitorevent = 5
	Filemonitorevent_Unmounted        Filemonitorevent = 6
	Filemonitorevent_Moved            Filemonitorevent = 7
	Filemonitorevent_Renamed          Filemonitorevent = 8
	Filemonitorevent_MovedIn          Filemonitorevent = 9
	Filemonitorevent_MovedOut         Filemonitorevent = 10
)

// Filetype is a representation of the C type FileType.
type Filetype int

const (
	Filetype_Unknown      Filetype = 0
	Filetype_Regular      Filetype = 1
	Filetype_Directory    Filetype = 2
	Filetype_SymbolicLink Filetype = 3
	Filetype_Special      Filetype = 4
	Filetype_Shortcut     Filetype = 5
	Filetype_Mountable    Filetype = 6
)

// Filesystempreviewtype is a representation of the C type FilesystemPreviewType.
type Filesystempreviewtype int

const (
	Filesystempreviewtype_IfAlways Filesystempreviewtype = 0
	Filesystempreviewtype_IfLocal  Filesystempreviewtype = 1
	Filesystempreviewtype_Never    Filesystempreviewtype = 2
)

// Ioerrorenum is a representation of the C type IOErrorEnum.
type Ioerrorenum int

const (
	Ioerrorenum_Failed             Ioerrorenum = 0
	Ioerrorenum_NotFound           Ioerrorenum = 1
	Ioerrorenum_Exists             Ioerrorenum = 2
	Ioerrorenum_IsDirectory        Ioerrorenum = 3
	Ioerrorenum_NotDirectory       Ioerrorenum = 4
	Ioerrorenum_NotEmpty           Ioerrorenum = 5
	Ioerrorenum_NotRegularFile     Ioerrorenum = 6
	Ioerrorenum_NotSymbolicLink    Ioerrorenum = 7
	Ioerrorenum_NotMountableFile   Ioerrorenum = 8
	Ioerrorenum_FilenameTooLong    Ioerrorenum = 9
	Ioerrorenum_InvalidFilename    Ioerrorenum = 10
	Ioerrorenum_TooManyLinks       Ioerrorenum = 11
	Ioerrorenum_NoSpace            Ioerrorenum = 12
	Ioerrorenum_InvalidArgument    Ioerrorenum = 13
	Ioerrorenum_PermissionDenied   Ioerrorenum = 14
	Ioerrorenum_NotSupported       Ioerrorenum = 15
	Ioerrorenum_NotMounted         Ioerrorenum = 16
	Ioerrorenum_AlreadyMounted     Ioerrorenum = 17
	Ioerrorenum_Closed             Ioerrorenum = 18
	Ioerrorenum_Cancelled          Ioerrorenum = 19
	Ioerrorenum_Pending            Ioerrorenum = 20
	Ioerrorenum_ReadOnly           Ioerrorenum = 21
	Ioerrorenum_CantCreateBackup   Ioerrorenum = 22
	Ioerrorenum_WrongEtag          Ioerrorenum = 23
	Ioerrorenum_TimedOut           Ioerrorenum = 24
	Ioerrorenum_WouldRecurse       Ioerrorenum = 25
	Ioerrorenum_Busy               Ioerrorenum = 26
	Ioerrorenum_WouldBlock         Ioerrorenum = 27
	Ioerrorenum_HostNotFound       Ioerrorenum = 28
	Ioerrorenum_WouldMerge         Ioerrorenum = 29
	Ioerrorenum_FailedHandled      Ioerrorenum = 30
	Ioerrorenum_TooManyOpenFiles   Ioerrorenum = 31
	Ioerrorenum_NotInitialized     Ioerrorenum = 32
	Ioerrorenum_AddressInUse       Ioerrorenum = 33
	Ioerrorenum_PartialInput       Ioerrorenum = 34
	Ioerrorenum_InvalidData        Ioerrorenum = 35
	Ioerrorenum_DbusError          Ioerrorenum = 36
	Ioerrorenum_HostUnreachable    Ioerrorenum = 37
	Ioerrorenum_NetworkUnreachable Ioerrorenum = 38
	Ioerrorenum_ConnectionRefused  Ioerrorenum = 39
	Ioerrorenum_ProxyFailed        Ioerrorenum = 40
	Ioerrorenum_ProxyAuthFailed    Ioerrorenum = 41
	Ioerrorenum_ProxyNeedAuth      Ioerrorenum = 42
	Ioerrorenum_ProxyNotAllowed    Ioerrorenum = 43
	Ioerrorenum_BrokenPipe         Ioerrorenum = 44
	Ioerrorenum_ConnectionClosed   Ioerrorenum = 44
	Ioerrorenum_NotConnected       Ioerrorenum = 45
	Ioerrorenum_MessageTooLarge    Ioerrorenum = 46
)

// Iomodulescopeflags is a representation of the C type IOModuleScopeFlags.
//
// since 2.30
type Iomodulescopeflags int

const (
	Iomodulescopeflags_None            Iomodulescopeflags = 0
	Iomodulescopeflags_BlockDuplicates Iomodulescopeflags = 1
)

// Mountoperationresult is a representation of the C type MountOperationResult.
type Mountoperationresult int

const (
	Mountoperationresult_Handled   Mountoperationresult = 0
	Mountoperationresult_Aborted   Mountoperationresult = 1
	Mountoperationresult_Unhandled Mountoperationresult = 2
)

// Networkconnectivity is a representation of the C type NetworkConnectivity.
//
// since 2.44
type Networkconnectivity int

const (
	Networkconnectivity_Local   Networkconnectivity = 1
	Networkconnectivity_Limited Networkconnectivity = 2
	Networkconnectivity_Portal  Networkconnectivity = 3
	Networkconnectivity_Full    Networkconnectivity = 4
)

// Notificationpriority is a representation of the C type NotificationPriority.
//
// since 2.42
type Notificationpriority int

const (
	Notificationpriority_Normal Notificationpriority = 0
	Notificationpriority_Low    Notificationpriority = 1
	Notificationpriority_High   Notificationpriority = 2
	Notificationpriority_Urgent Notificationpriority = 3
)

// Passwordsave is a representation of the C type PasswordSave.
type Passwordsave int

const (
	Passwordsave_Never       Passwordsave = 0
	Passwordsave_ForSession  Passwordsave = 1
	Passwordsave_Permanently Passwordsave = 2
)

// Resolvererror is a representation of the C type ResolverError.
//
// since 2.22
type Resolvererror int

const (
	Resolvererror_NotFound         Resolvererror = 0
	Resolvererror_TemporaryFailure Resolvererror = 1
	Resolvererror_Internal         Resolvererror = 2
)

// Resolverrecordtype is a representation of the C type ResolverRecordType.
//
// since 2.34
type Resolverrecordtype int

const (
	Resolverrecordtype_Srv Resolverrecordtype = 1
	Resolverrecordtype_Mx  Resolverrecordtype = 2
	Resolverrecordtype_Txt Resolverrecordtype = 3
	Resolverrecordtype_Soa Resolverrecordtype = 4
	Resolverrecordtype_Ns  Resolverrecordtype = 5
)

// Resourceerror is a representation of the C type ResourceError.
//
// since 2.32
type Resourceerror int

const (
	Resourceerror_NotFound Resourceerror = 0
	Resourceerror_Internal Resourceerror = 1
)

// Socketclientevent is a representation of the C type SocketClientEvent.
//
// since 2.32
type Socketclientevent int

const (
	Socketclientevent_Resolving        Socketclientevent = 0
	Socketclientevent_Resolved         Socketclientevent = 1
	Socketclientevent_Connecting       Socketclientevent = 2
	Socketclientevent_Connected        Socketclientevent = 3
	Socketclientevent_ProxyNegotiating Socketclientevent = 4
	Socketclientevent_ProxyNegotiated  Socketclientevent = 5
	Socketclientevent_TlsHandshaking   Socketclientevent = 6
	Socketclientevent_TlsHandshaked    Socketclientevent = 7
	Socketclientevent_Complete         Socketclientevent = 8
)

// Socketfamily is a representation of the C type SocketFamily.
//
// since 2.22
type Socketfamily int

const (
	Socketfamily_Invalid Socketfamily = 0
	Socketfamily_Unix    Socketfamily = 1
	Socketfamily_Ipv4    Socketfamily = 2
	Socketfamily_Ipv6    Socketfamily = 10
)

// Socketlistenerevent is a representation of the C type SocketListenerEvent.
//
// since 2.46
type Socketlistenerevent int

const (
	Socketlistenerevent_Binding   Socketlistenerevent = 0
	Socketlistenerevent_Bound     Socketlistenerevent = 1
	Socketlistenerevent_Listening Socketlistenerevent = 2
	Socketlistenerevent_Listened  Socketlistenerevent = 3
)

// Socketprotocol is a representation of the C type SocketProtocol.
//
// since 2.22
type Socketprotocol int

const (
	Socketprotocol_Unknown Socketprotocol = -1
	Socketprotocol_Default Socketprotocol = 0
	Socketprotocol_Tcp     Socketprotocol = 6
	Socketprotocol_Udp     Socketprotocol = 17
	Socketprotocol_Sctp    Socketprotocol = 132
)

// Sockettype is a representation of the C type SocketType.
//
// since 2.22
type Sockettype int

const (
	Sockettype_Invalid   Sockettype = 0
	Sockettype_Stream    Sockettype = 1
	Sockettype_Datagram  Sockettype = 2
	Sockettype_Seqpacket Sockettype = 3
)

// Tlsauthenticationmode is a representation of the C type TlsAuthenticationMode.
//
// since 2.28
type Tlsauthenticationmode int

const (
	Tlsauthenticationmode_None      Tlsauthenticationmode = 0
	Tlsauthenticationmode_Requested Tlsauthenticationmode = 1
	Tlsauthenticationmode_Required  Tlsauthenticationmode = 2
)

// Tlscertificaterequestflags is a representation of the C type TlsCertificateRequestFlags.
//
// since 2.40
type Tlscertificaterequestflags int

const (
	Tlscertificaterequestflags_None Tlscertificaterequestflags = 0
)

// Tlsdatabaselookupflags is a representation of the C type TlsDatabaseLookupFlags.
//
// since 2.30
type Tlsdatabaselookupflags int

const (
	Tlsdatabaselookupflags_None    Tlsdatabaselookupflags = 0
	Tlsdatabaselookupflags_Keypair Tlsdatabaselookupflags = 1
)

// Tlserror is a representation of the C type TlsError.
//
// since 2.28
type Tlserror int

const (
	Tlserror_Unavailable         Tlserror = 0
	Tlserror_Misc                Tlserror = 1
	Tlserror_BadCertificate      Tlserror = 2
	Tlserror_NotTls              Tlserror = 3
	Tlserror_Handshake           Tlserror = 4
	Tlserror_CertificateRequired Tlserror = 5
	Tlserror_Eof                 Tlserror = 6
)

// Tlsinteractionresult is a representation of the C type TlsInteractionResult.
//
// since 2.30
type Tlsinteractionresult int

const (
	Tlsinteractionresult_Unhandled Tlsinteractionresult = 0
	Tlsinteractionresult_Handled   Tlsinteractionresult = 1
	Tlsinteractionresult_Failed    Tlsinteractionresult = 2
)

// Tlsrehandshakemode is a representation of the C type TlsRehandshakeMode.
//
// since 2.28
type Tlsrehandshakemode int

const (
	Tlsrehandshakemode_Never    Tlsrehandshakemode = 0
	Tlsrehandshakemode_Safely   Tlsrehandshakemode = 1
	Tlsrehandshakemode_Unsafely Tlsrehandshakemode = 2
)

// Unixsocketaddresstype is a representation of the C type UnixSocketAddressType.
//
// since 2.26
type Unixsocketaddresstype int

const (
	Unixsocketaddresstype_Invalid        Unixsocketaddresstype = 0
	Unixsocketaddresstype_Anonymous      Unixsocketaddresstype = 1
	Unixsocketaddresstype_Path           Unixsocketaddresstype = 2
	Unixsocketaddresstype_Abstract       Unixsocketaddresstype = 3
	Unixsocketaddresstype_AbstractPadded Unixsocketaddresstype = 4
)

// Zlibcompressorformat is a representation of the C type ZlibCompressorFormat.
//
// since 2.24
type Zlibcompressorformat int

const (
	Zlibcompressorformat_Zlib Zlibcompressorformat = 0
	Zlibcompressorformat_Gzip Zlibcompressorformat = 1
	Zlibcompressorformat_Raw  Zlibcompressorformat = 2
)
