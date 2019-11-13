// Code generated - DO NOT EDIT.

package gio

// Bustype is a representation of the C type BusType.
//
// since 2.26
type Bustype int

const (
	// starter
	GBusTypeStarter Bustype = -1
	// none
	GBusTypeNone Bustype = 0
	// system
	GBusTypeSystem Bustype = 1
	// session
	GBusTypeSession Bustype = 2
)

// Converterresult is a representation of the C type ConverterResult.
//
// since 2.24
type Converterresult int

const (
	// error
	GConverterError Converterresult = 0
	// converted
	GConverterConverted Converterresult = 1
	// finished
	GConverterFinished Converterresult = 2
	// flushed
	GConverterFlushed Converterresult = 3
)

// Credentialstype is a representation of the C type CredentialsType.
//
// since 2.26
type Credentialstype int

const (
	// invalid
	GCredentialsTypeInvalid Credentialstype = 0
	// linux_ucred
	GCredentialsTypeLinuxUcred Credentialstype = 1
	// freebsd_cmsgcred
	GCredentialsTypeFreebsdCmsgcred Credentialstype = 2
	// openbsd_sockpeercred
	GCredentialsTypeOpenbsdSockpeercred Credentialstype = 3
	// solaris_ucred
	GCredentialsTypeSolarisUcred Credentialstype = 4
	// netbsd_unpcbid
	GCredentialsTypeNetbsdUnpcbid Credentialstype = 5
)

// Dbuserror is a representation of the C type DBusError.
//
// since 2.26
type Dbuserror int

const (
	// failed
	GDbusErrorFailed Dbuserror = 0
	// no_memory
	GDbusErrorNoMemory Dbuserror = 1
	// service_unknown
	GDbusErrorServiceUnknown Dbuserror = 2
	// name_has_no_owner
	GDbusErrorNameHasNoOwner Dbuserror = 3
	// no_reply
	GDbusErrorNoReply Dbuserror = 4
	// io_error
	GDbusErrorIoError Dbuserror = 5
	// bad_address
	GDbusErrorBadAddress Dbuserror = 6
	// not_supported
	GDbusErrorNotSupported Dbuserror = 7
	// limits_exceeded
	GDbusErrorLimitsExceeded Dbuserror = 8
	// access_denied
	GDbusErrorAccessDenied Dbuserror = 9
	// auth_failed
	GDbusErrorAuthFailed Dbuserror = 10
	// no_server
	GDbusErrorNoServer Dbuserror = 11
	// timeout
	GDbusErrorTimeout Dbuserror = 12
	// no_network
	GDbusErrorNoNetwork Dbuserror = 13
	// address_in_use
	GDbusErrorAddressInUse Dbuserror = 14
	// disconnected
	GDbusErrorDisconnected Dbuserror = 15
	// invalid_args
	GDbusErrorInvalidArgs Dbuserror = 16
	// file_not_found
	GDbusErrorFileNotFound Dbuserror = 17
	// file_exists
	GDbusErrorFileExists Dbuserror = 18
	// unknown_method
	GDbusErrorUnknownMethod Dbuserror = 19
	// timed_out
	GDbusErrorTimedOut Dbuserror = 20
	// match_rule_not_found
	GDbusErrorMatchRuleNotFound Dbuserror = 21
	// match_rule_invalid
	GDbusErrorMatchRuleInvalid Dbuserror = 22
	// spawn_exec_failed
	GDbusErrorSpawnExecFailed Dbuserror = 23
	// spawn_fork_failed
	GDbusErrorSpawnForkFailed Dbuserror = 24
	// spawn_child_exited
	GDbusErrorSpawnChildExited Dbuserror = 25
	// spawn_child_signaled
	GDbusErrorSpawnChildSignaled Dbuserror = 26
	// spawn_failed
	GDbusErrorSpawnFailed Dbuserror = 27
	// spawn_setup_failed
	GDbusErrorSpawnSetupFailed Dbuserror = 28
	// spawn_config_invalid
	GDbusErrorSpawnConfigInvalid Dbuserror = 29
	// spawn_service_invalid
	GDbusErrorSpawnServiceInvalid Dbuserror = 30
	// spawn_service_not_found
	GDbusErrorSpawnServiceNotFound Dbuserror = 31
	// spawn_permissions_invalid
	GDbusErrorSpawnPermissionsInvalid Dbuserror = 32
	// spawn_file_invalid
	GDbusErrorSpawnFileInvalid Dbuserror = 33
	// spawn_no_memory
	GDbusErrorSpawnNoMemory Dbuserror = 34
	// unix_process_id_unknown
	GDbusErrorUnixProcessIdUnknown Dbuserror = 35
	// invalid_signature
	GDbusErrorInvalidSignature Dbuserror = 36
	// invalid_file_content
	GDbusErrorInvalidFileContent Dbuserror = 37
	// selinux_security_context_unknown
	GDbusErrorSelinuxSecurityContextUnknown Dbuserror = 38
	// adt_audit_data_unknown
	GDbusErrorAdtAuditDataUnknown Dbuserror = 39
	// object_path_in_use
	GDbusErrorObjectPathInUse Dbuserror = 40
	// unknown_object
	GDbusErrorUnknownObject Dbuserror = 41
	// unknown_interface
	GDbusErrorUnknownInterface Dbuserror = 42
	// unknown_property
	GDbusErrorUnknownProperty Dbuserror = 43
	// property_read_only
	GDbusErrorPropertyReadOnly Dbuserror = 44
)

// Dbusmessagebyteorder is a representation of the C type DBusMessageByteOrder.
//
// since 2.26
type Dbusmessagebyteorder int

const (
	// big_endian
	GDbusMessageByteOrderBigEndian Dbusmessagebyteorder = 66
	// little_endian
	GDbusMessageByteOrderLittleEndian Dbusmessagebyteorder = 108
)

// Dbusmessageheaderfield is a representation of the C type DBusMessageHeaderField.
//
// since 2.26
type Dbusmessageheaderfield int

const (
	// invalid
	GDbusMessageHeaderFieldInvalid Dbusmessageheaderfield = 0
	// path
	GDbusMessageHeaderFieldPath Dbusmessageheaderfield = 1
	// interface
	GDbusMessageHeaderFieldInterface Dbusmessageheaderfield = 2
	// member
	GDbusMessageHeaderFieldMember Dbusmessageheaderfield = 3
	// error_name
	GDbusMessageHeaderFieldErrorName Dbusmessageheaderfield = 4
	// reply_serial
	GDbusMessageHeaderFieldReplySerial Dbusmessageheaderfield = 5
	// destination
	GDbusMessageHeaderFieldDestination Dbusmessageheaderfield = 6
	// sender
	GDbusMessageHeaderFieldSender Dbusmessageheaderfield = 7
	// signature
	GDbusMessageHeaderFieldSignature Dbusmessageheaderfield = 8
	// num_unix_fds
	GDbusMessageHeaderFieldNumUnixFds Dbusmessageheaderfield = 9
)

// Dbusmessagetype is a representation of the C type DBusMessageType.
//
// since 2.26
type Dbusmessagetype int

const (
	// invalid
	GDbusMessageTypeInvalid Dbusmessagetype = 0
	// method_call
	GDbusMessageTypeMethodCall Dbusmessagetype = 1
	// method_return
	GDbusMessageTypeMethodReturn Dbusmessagetype = 2
	// error
	GDbusMessageTypeError Dbusmessagetype = 3
	// signal
	GDbusMessageTypeSignal Dbusmessagetype = 4
)

// Datastreambyteorder is a representation of the C type DataStreamByteOrder.
type Datastreambyteorder int

const (
	// big_endian
	GDataStreamByteOrderBigEndian Datastreambyteorder = 0
	// little_endian
	GDataStreamByteOrderLittleEndian Datastreambyteorder = 1
	// host_endian
	GDataStreamByteOrderHostEndian Datastreambyteorder = 2
)

// Datastreamnewlinetype is a representation of the C type DataStreamNewlineType.
type Datastreamnewlinetype int

const (
	// lf
	GDataStreamNewlineTypeLf Datastreamnewlinetype = 0
	// cr
	GDataStreamNewlineTypeCr Datastreamnewlinetype = 1
	// cr_lf
	GDataStreamNewlineTypeCrLf Datastreamnewlinetype = 2
	// any
	GDataStreamNewlineTypeAny Datastreamnewlinetype = 3
)

// Drivestartstoptype is a representation of the C type DriveStartStopType.
//
// since 2.22
type Drivestartstoptype int

const (
	// unknown
	GDriveStartStopTypeUnknown Drivestartstoptype = 0
	// shutdown
	GDriveStartStopTypeShutdown Drivestartstoptype = 1
	// network
	GDriveStartStopTypeNetwork Drivestartstoptype = 2
	// multidisk
	GDriveStartStopTypeMultidisk Drivestartstoptype = 3
	// password
	GDriveStartStopTypePassword Drivestartstoptype = 4
)

// Emblemorigin is a representation of the C type EmblemOrigin.
//
// since 2.18
type Emblemorigin int

const (
	// unknown
	GEmblemOriginUnknown Emblemorigin = 0
	// device
	GEmblemOriginDevice Emblemorigin = 1
	// livemetadata
	GEmblemOriginLivemetadata Emblemorigin = 2
	// tag
	GEmblemOriginTag Emblemorigin = 3
)

// Fileattributestatus is a representation of the C type FileAttributeStatus.
type Fileattributestatus int

const (
	// unset
	GFileAttributeStatusUnset Fileattributestatus = 0
	// set
	GFileAttributeStatusSet Fileattributestatus = 1
	// error_setting
	GFileAttributeStatusErrorSetting Fileattributestatus = 2
)

// Fileattributetype is a representation of the C type FileAttributeType.
type Fileattributetype int

const (
	// invalid
	GFileAttributeTypeInvalid Fileattributetype = 0
	// string
	GFileAttributeTypeString Fileattributetype = 1
	// byte_string
	GFileAttributeTypeByteString Fileattributetype = 2
	// boolean
	GFileAttributeTypeBoolean Fileattributetype = 3
	// uint32
	GFileAttributeTypeUint32 Fileattributetype = 4
	// int32
	GFileAttributeTypeInt32 Fileattributetype = 5
	// uint64
	GFileAttributeTypeUint64 Fileattributetype = 6
	// int64
	GFileAttributeTypeInt64 Fileattributetype = 7
	// object
	GFileAttributeTypeObject Fileattributetype = 8
	// stringv
	GFileAttributeTypeStringv Fileattributetype = 9
)

// Filemonitorevent is a representation of the C type FileMonitorEvent.
type Filemonitorevent int

const (
	// changed
	GFileMonitorEventChanged Filemonitorevent = 0
	// changes_done_hint
	GFileMonitorEventChangesDoneHint Filemonitorevent = 1
	// deleted
	GFileMonitorEventDeleted Filemonitorevent = 2
	// created
	GFileMonitorEventCreated Filemonitorevent = 3
	// attribute_changed
	GFileMonitorEventAttributeChanged Filemonitorevent = 4
	// pre_unmount
	GFileMonitorEventPreUnmount Filemonitorevent = 5
	// unmounted
	GFileMonitorEventUnmounted Filemonitorevent = 6
	// moved
	GFileMonitorEventMoved Filemonitorevent = 7
	// renamed
	GFileMonitorEventRenamed Filemonitorevent = 8
	// moved_in
	GFileMonitorEventMovedIn Filemonitorevent = 9
	// moved_out
	GFileMonitorEventMovedOut Filemonitorevent = 10
)

// Filetype is a representation of the C type FileType.
type Filetype int

const (
	// unknown
	GFileTypeUnknown Filetype = 0
	// regular
	GFileTypeRegular Filetype = 1
	// directory
	GFileTypeDirectory Filetype = 2
	// symbolic_link
	GFileTypeSymbolicLink Filetype = 3
	// special
	GFileTypeSpecial Filetype = 4
	// shortcut
	GFileTypeShortcut Filetype = 5
	// mountable
	GFileTypeMountable Filetype = 6
)

// Filesystempreviewtype is a representation of the C type FilesystemPreviewType.
type Filesystempreviewtype int

const (
	// if_always
	GFilesystemPreviewTypeIfAlways Filesystempreviewtype = 0
	// if_local
	GFilesystemPreviewTypeIfLocal Filesystempreviewtype = 1
	// never
	GFilesystemPreviewTypeNever Filesystempreviewtype = 2
)

// Ioerrorenum is a representation of the C type IOErrorEnum.
type Ioerrorenum int

const (
	// failed
	GIoErrorFailed Ioerrorenum = 0
	// not_found
	GIoErrorNotFound Ioerrorenum = 1
	// exists
	GIoErrorExists Ioerrorenum = 2
	// is_directory
	GIoErrorIsDirectory Ioerrorenum = 3
	// not_directory
	GIoErrorNotDirectory Ioerrorenum = 4
	// not_empty
	GIoErrorNotEmpty Ioerrorenum = 5
	// not_regular_file
	GIoErrorNotRegularFile Ioerrorenum = 6
	// not_symbolic_link
	GIoErrorNotSymbolicLink Ioerrorenum = 7
	// not_mountable_file
	GIoErrorNotMountableFile Ioerrorenum = 8
	// filename_too_long
	GIoErrorFilenameTooLong Ioerrorenum = 9
	// invalid_filename
	GIoErrorInvalidFilename Ioerrorenum = 10
	// too_many_links
	GIoErrorTooManyLinks Ioerrorenum = 11
	// no_space
	GIoErrorNoSpace Ioerrorenum = 12
	// invalid_argument
	GIoErrorInvalidArgument Ioerrorenum = 13
	// permission_denied
	GIoErrorPermissionDenied Ioerrorenum = 14
	// not_supported
	GIoErrorNotSupported Ioerrorenum = 15
	// not_mounted
	GIoErrorNotMounted Ioerrorenum = 16
	// already_mounted
	GIoErrorAlreadyMounted Ioerrorenum = 17
	// closed
	GIoErrorClosed Ioerrorenum = 18
	// cancelled
	GIoErrorCancelled Ioerrorenum = 19
	// pending
	GIoErrorPending Ioerrorenum = 20
	// read_only
	GIoErrorReadOnly Ioerrorenum = 21
	// cant_create_backup
	GIoErrorCantCreateBackup Ioerrorenum = 22
	// wrong_etag
	GIoErrorWrongEtag Ioerrorenum = 23
	// timed_out
	GIoErrorTimedOut Ioerrorenum = 24
	// would_recurse
	GIoErrorWouldRecurse Ioerrorenum = 25
	// busy
	GIoErrorBusy Ioerrorenum = 26
	// would_block
	GIoErrorWouldBlock Ioerrorenum = 27
	// host_not_found
	GIoErrorHostNotFound Ioerrorenum = 28
	// would_merge
	GIoErrorWouldMerge Ioerrorenum = 29
	// failed_handled
	GIoErrorFailedHandled Ioerrorenum = 30
	// too_many_open_files
	GIoErrorTooManyOpenFiles Ioerrorenum = 31
	// not_initialized
	GIoErrorNotInitialized Ioerrorenum = 32
	// address_in_use
	GIoErrorAddressInUse Ioerrorenum = 33
	// partial_input
	GIoErrorPartialInput Ioerrorenum = 34
	// invalid_data
	GIoErrorInvalidData Ioerrorenum = 35
	// dbus_error
	GIoErrorDbusError Ioerrorenum = 36
	// host_unreachable
	GIoErrorHostUnreachable Ioerrorenum = 37
	// network_unreachable
	GIoErrorNetworkUnreachable Ioerrorenum = 38
	// connection_refused
	GIoErrorConnectionRefused Ioerrorenum = 39
	// proxy_failed
	GIoErrorProxyFailed Ioerrorenum = 40
	// proxy_auth_failed
	GIoErrorProxyAuthFailed Ioerrorenum = 41
	// proxy_need_auth
	GIoErrorProxyNeedAuth Ioerrorenum = 42
	// proxy_not_allowed
	GIoErrorProxyNotAllowed Ioerrorenum = 43
	// broken_pipe
	GIoErrorBrokenPipe Ioerrorenum = 44
	// connection_closed
	GIoErrorConnectionClosed Ioerrorenum = 44
	// not_connected
	GIoErrorNotConnected Ioerrorenum = 45
	// message_too_large
	GIoErrorMessageTooLarge Ioerrorenum = 46
)

// Iomodulescopeflags is a representation of the C type IOModuleScopeFlags.
//
// since 2.30
type Iomodulescopeflags int

const (
	// none
	GIoModuleScopeNone Iomodulescopeflags = 0
	// block_duplicates
	GIoModuleScopeBlockDuplicates Iomodulescopeflags = 1
)

// Mountoperationresult is a representation of the C type MountOperationResult.
type Mountoperationresult int

const (
	// handled
	GMountOperationHandled Mountoperationresult = 0
	// aborted
	GMountOperationAborted Mountoperationresult = 1
	// unhandled
	GMountOperationUnhandled Mountoperationresult = 2
)

// Networkconnectivity is a representation of the C type NetworkConnectivity.
//
// since 2.44
type Networkconnectivity int

const (
	// local
	GNetworkConnectivityLocal Networkconnectivity = 1
	// limited
	GNetworkConnectivityLimited Networkconnectivity = 2
	// portal
	GNetworkConnectivityPortal Networkconnectivity = 3
	// full
	GNetworkConnectivityFull Networkconnectivity = 4
)

// Notificationpriority is a representation of the C type NotificationPriority.
//
// since 2.42
type Notificationpriority int

const (
	// normal
	GNotificationPriorityNormal Notificationpriority = 0
	// low
	GNotificationPriorityLow Notificationpriority = 1
	// high
	GNotificationPriorityHigh Notificationpriority = 2
	// urgent
	GNotificationPriorityUrgent Notificationpriority = 3
)

// Passwordsave is a representation of the C type PasswordSave.
type Passwordsave int

const (
	// never
	GPasswordSaveNever Passwordsave = 0
	// for_session
	GPasswordSaveForSession Passwordsave = 1
	// permanently
	GPasswordSavePermanently Passwordsave = 2
)

// Resolvererror is a representation of the C type ResolverError.
//
// since 2.22
type Resolvererror int

const (
	// not_found
	GResolverErrorNotFound Resolvererror = 0
	// temporary_failure
	GResolverErrorTemporaryFailure Resolvererror = 1
	// internal
	GResolverErrorInternal Resolvererror = 2
)

// Resolverrecordtype is a representation of the C type ResolverRecordType.
//
// since 2.34
type Resolverrecordtype int

const (
	// srv
	GResolverRecordSrv Resolverrecordtype = 1
	// mx
	GResolverRecordMx Resolverrecordtype = 2
	// txt
	GResolverRecordTxt Resolverrecordtype = 3
	// soa
	GResolverRecordSoa Resolverrecordtype = 4
	// ns
	GResolverRecordNs Resolverrecordtype = 5
)

// Resourceerror is a representation of the C type ResourceError.
//
// since 2.32
type Resourceerror int

const (
	// not_found
	GResourceErrorNotFound Resourceerror = 0
	// internal
	GResourceErrorInternal Resourceerror = 1
)

// Socketclientevent is a representation of the C type SocketClientEvent.
//
// since 2.32
type Socketclientevent int

const (
	// resolving
	GSocketClientResolving Socketclientevent = 0
	// resolved
	GSocketClientResolved Socketclientevent = 1
	// connecting
	GSocketClientConnecting Socketclientevent = 2
	// connected
	GSocketClientConnected Socketclientevent = 3
	// proxy_negotiating
	GSocketClientProxyNegotiating Socketclientevent = 4
	// proxy_negotiated
	GSocketClientProxyNegotiated Socketclientevent = 5
	// tls_handshaking
	GSocketClientTlsHandshaking Socketclientevent = 6
	// tls_handshaked
	GSocketClientTlsHandshaked Socketclientevent = 7
	// complete
	GSocketClientComplete Socketclientevent = 8
)

// Socketfamily is a representation of the C type SocketFamily.
//
// since 2.22
type Socketfamily int

const (
	// invalid
	GSocketFamilyInvalid Socketfamily = 0
	// unix
	GSocketFamilyUnix Socketfamily = 1
	// ipv4
	GSocketFamilyIpv4 Socketfamily = 2
	// ipv6
	GSocketFamilyIpv6 Socketfamily = 10
)

// Socketlistenerevent is a representation of the C type SocketListenerEvent.
//
// since 2.46
type Socketlistenerevent int

const (
	// binding
	GSocketListenerBinding Socketlistenerevent = 0
	// bound
	GSocketListenerBound Socketlistenerevent = 1
	// listening
	GSocketListenerListening Socketlistenerevent = 2
	// listened
	GSocketListenerListened Socketlistenerevent = 3
)

// Socketprotocol is a representation of the C type SocketProtocol.
//
// since 2.22
type Socketprotocol int

const (
	// unknown
	GSocketProtocolUnknown Socketprotocol = -1
	// default
	GSocketProtocolDefault Socketprotocol = 0
	// tcp
	GSocketProtocolTcp Socketprotocol = 6
	// udp
	GSocketProtocolUdp Socketprotocol = 17
	// sctp
	GSocketProtocolSctp Socketprotocol = 132
)

// Sockettype is a representation of the C type SocketType.
//
// since 2.22
type Sockettype int

const (
	// invalid
	GSocketTypeInvalid Sockettype = 0
	// stream
	GSocketTypeStream Sockettype = 1
	// datagram
	GSocketTypeDatagram Sockettype = 2
	// seqpacket
	GSocketTypeSeqpacket Sockettype = 3
)

// Tlsauthenticationmode is a representation of the C type TlsAuthenticationMode.
//
// since 2.28
type Tlsauthenticationmode int

const (
	// none
	GTlsAuthenticationNone Tlsauthenticationmode = 0
	// requested
	GTlsAuthenticationRequested Tlsauthenticationmode = 1
	// required
	GTlsAuthenticationRequired Tlsauthenticationmode = 2
)

// Tlscertificaterequestflags is a representation of the C type TlsCertificateRequestFlags.
//
// since 2.40
type Tlscertificaterequestflags int

const (
	// none
	GTlsCertificateRequestNone Tlscertificaterequestflags = 0
)

// Tlsdatabaselookupflags is a representation of the C type TlsDatabaseLookupFlags.
//
// since 2.30
type Tlsdatabaselookupflags int

const (
	// none
	GTlsDatabaseLookupNone Tlsdatabaselookupflags = 0
	// keypair
	GTlsDatabaseLookupKeypair Tlsdatabaselookupflags = 1
)

// Tlserror is a representation of the C type TlsError.
//
// since 2.28
type Tlserror int

const (
	// unavailable
	GTlsErrorUnavailable Tlserror = 0
	// misc
	GTlsErrorMisc Tlserror = 1
	// bad_certificate
	GTlsErrorBadCertificate Tlserror = 2
	// not_tls
	GTlsErrorNotTls Tlserror = 3
	// handshake
	GTlsErrorHandshake Tlserror = 4
	// certificate_required
	GTlsErrorCertificateRequired Tlserror = 5
	// eof
	GTlsErrorEof Tlserror = 6
)

// Tlsinteractionresult is a representation of the C type TlsInteractionResult.
//
// since 2.30
type Tlsinteractionresult int

const (
	// unhandled
	GTlsInteractionUnhandled Tlsinteractionresult = 0
	// handled
	GTlsInteractionHandled Tlsinteractionresult = 1
	// failed
	GTlsInteractionFailed Tlsinteractionresult = 2
)

// Tlsrehandshakemode is a representation of the C type TlsRehandshakeMode.
//
// since 2.28
type Tlsrehandshakemode int

const (
	// never
	GTlsRehandshakeNever Tlsrehandshakemode = 0
	// safely
	GTlsRehandshakeSafely Tlsrehandshakemode = 1
	// unsafely
	GTlsRehandshakeUnsafely Tlsrehandshakemode = 2
)

// Unixsocketaddresstype is a representation of the C type UnixSocketAddressType.
//
// since 2.26
type Unixsocketaddresstype int

const (
	// invalid
	GUnixSocketAddressInvalid Unixsocketaddresstype = 0
	// anonymous
	GUnixSocketAddressAnonymous Unixsocketaddresstype = 1
	// path
	GUnixSocketAddressPath Unixsocketaddresstype = 2
	// abstract
	GUnixSocketAddressAbstract Unixsocketaddresstype = 3
	// abstract_padded
	GUnixSocketAddressAbstractPadded Unixsocketaddresstype = 4
)

// Zlibcompressorformat is a representation of the C type ZlibCompressorFormat.
//
// since 2.24
type Zlibcompressorformat int

const (
	// zlib
	GZlibCompressorFormatZlib Zlibcompressorformat = 0
	// gzip
	GZlibCompressorFormatGzip Zlibcompressorformat = 1
	// raw
	GZlibCompressorFormatRaw Zlibcompressorformat = 2
)
