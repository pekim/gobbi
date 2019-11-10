// Code generated - DO NOT EDIT.

package gio

type AppInfoCreateFlags uint32

const (
	AppInfoCreateFlags_None                        AppInfoCreateFlags = int64(0)
	AppInfoCreateFlags_NeedsTerminal               AppInfoCreateFlags = int64(1)
	AppInfoCreateFlags_SupportsUris                AppInfoCreateFlags = int64(2)
	AppInfoCreateFlags_SupportsStartupNotification AppInfoCreateFlags = int64(4)
)

type ApplicationFlags uint32

const (
	ApplicationFlags_FlagsNone          ApplicationFlags = int64(0)
	ApplicationFlags_IsService          ApplicationFlags = int64(1)
	ApplicationFlags_IsLauncher         ApplicationFlags = int64(2)
	ApplicationFlags_HandlesOpen        ApplicationFlags = int64(4)
	ApplicationFlags_HandlesCommandLine ApplicationFlags = int64(8)
	ApplicationFlags_SendEnvironment    ApplicationFlags = int64(16)
	ApplicationFlags_NonUnique          ApplicationFlags = int64(32)
	ApplicationFlags_CanOverrideAppId   ApplicationFlags = int64(64)
)

type AskPasswordFlags uint32

const (
	AskPasswordFlags_NeedPassword       AskPasswordFlags = int64(1)
	AskPasswordFlags_NeedUsername       AskPasswordFlags = int64(2)
	AskPasswordFlags_NeedDomain         AskPasswordFlags = int64(4)
	AskPasswordFlags_SavingSupported    AskPasswordFlags = int64(8)
	AskPasswordFlags_AnonymousSupported AskPasswordFlags = int64(16)
)

type BusNameOwnerFlags uint32

const (
	BusNameOwnerFlags_None             BusNameOwnerFlags = int64(0)
	BusNameOwnerFlags_AllowReplacement BusNameOwnerFlags = int64(1)
	BusNameOwnerFlags_Replace          BusNameOwnerFlags = int64(2)
	BusNameOwnerFlags_DoNotQueue       BusNameOwnerFlags = int64(4)
)

type BusNameWatcherFlags uint32

const (
	BusNameWatcherFlags_None      BusNameWatcherFlags = int64(0)
	BusNameWatcherFlags_AutoStart BusNameWatcherFlags = int64(1)
)

type ConverterFlags uint32

const (
	ConverterFlags_None       ConverterFlags = int64(0)
	ConverterFlags_InputAtEnd ConverterFlags = int64(1)
	ConverterFlags_Flush      ConverterFlags = int64(2)
)

type DBusCallFlags uint32

const (
	DBusCallFlags_None                          DBusCallFlags = int64(0)
	DBusCallFlags_NoAutoStart                   DBusCallFlags = int64(1)
	DBusCallFlags_AllowInteractiveAuthorization DBusCallFlags = int64(2)
)

type DBusCapabilityFlags uint32

const (
	DBusCapabilityFlags_None          DBusCapabilityFlags = int64(0)
	DBusCapabilityFlags_UnixFdPassing DBusCapabilityFlags = int64(1)
)

type DBusConnectionFlags uint32

const (
	DBusConnectionFlags_None                         DBusConnectionFlags = int64(0)
	DBusConnectionFlags_AuthenticationClient         DBusConnectionFlags = int64(1)
	DBusConnectionFlags_AuthenticationServer         DBusConnectionFlags = int64(2)
	DBusConnectionFlags_AuthenticationAllowAnonymous DBusConnectionFlags = int64(4)
	DBusConnectionFlags_MessageBusConnection         DBusConnectionFlags = int64(8)
	DBusConnectionFlags_DelayMessageProcessing       DBusConnectionFlags = int64(16)
)

type DBusInterfaceSkeletonFlags uint32

const (
	DBusInterfaceSkeletonFlags_None                            DBusInterfaceSkeletonFlags = int64(0)
	DBusInterfaceSkeletonFlags_HandleMethodInvocationsInThread DBusInterfaceSkeletonFlags = int64(1)
)

type DBusMessageFlags uint32

const (
	DBusMessageFlags_None                          DBusMessageFlags = int64(0)
	DBusMessageFlags_NoReplyExpected               DBusMessageFlags = int64(1)
	DBusMessageFlags_NoAutoStart                   DBusMessageFlags = int64(2)
	DBusMessageFlags_AllowInteractiveAuthorization DBusMessageFlags = int64(4)
)

type DBusObjectManagerClientFlags uint32

const (
	DBusObjectManagerClientFlags_None           DBusObjectManagerClientFlags = int64(0)
	DBusObjectManagerClientFlags_DoNotAutoStart DBusObjectManagerClientFlags = int64(1)
)

type DBusPropertyInfoFlags uint32

const (
	DBusPropertyInfoFlags_None     DBusPropertyInfoFlags = int64(0)
	DBusPropertyInfoFlags_Readable DBusPropertyInfoFlags = int64(1)
	DBusPropertyInfoFlags_Writable DBusPropertyInfoFlags = int64(2)
)

type DBusProxyFlags uint32

const (
	DBusProxyFlags_None                         DBusProxyFlags = int64(0)
	DBusProxyFlags_DoNotLoadProperties          DBusProxyFlags = int64(1)
	DBusProxyFlags_DoNotConnectSignals          DBusProxyFlags = int64(2)
	DBusProxyFlags_DoNotAutoStart               DBusProxyFlags = int64(4)
	DBusProxyFlags_GetInvalidatedProperties     DBusProxyFlags = int64(8)
	DBusProxyFlags_DoNotAutoStartAtConstruction DBusProxyFlags = int64(16)
)

type DBusSendMessageFlags uint32

const (
	DBusSendMessageFlags_None           DBusSendMessageFlags = int64(0)
	DBusSendMessageFlags_PreserveSerial DBusSendMessageFlags = int64(1)
)

type DBusServerFlags uint32

const (
	DBusServerFlags_None                         DBusServerFlags = int64(0)
	DBusServerFlags_RunInThread                  DBusServerFlags = int64(1)
	DBusServerFlags_AuthenticationAllowAnonymous DBusServerFlags = int64(2)
)

type DBusSignalFlags uint32

const (
	DBusSignalFlags_None               DBusSignalFlags = int64(0)
	DBusSignalFlags_NoMatchRule        DBusSignalFlags = int64(1)
	DBusSignalFlags_MatchArg0Namespace DBusSignalFlags = int64(2)
	DBusSignalFlags_MatchArg0Path      DBusSignalFlags = int64(4)
)

type DBusSubtreeFlags uint32

const (
	DBusSubtreeFlags_None                        DBusSubtreeFlags = int64(0)
	DBusSubtreeFlags_DispatchToUnenumeratedNodes DBusSubtreeFlags = int64(1)
)

type DriveStartFlags uint32

const (
	DriveStartFlags_None DriveStartFlags = int64(0)
)

type FileAttributeInfoFlags uint32

const (
	FileAttributeInfoFlags_None          FileAttributeInfoFlags = int64(0)
	FileAttributeInfoFlags_CopyWithFile  FileAttributeInfoFlags = int64(1)
	FileAttributeInfoFlags_CopyWhenMoved FileAttributeInfoFlags = int64(2)
)

type FileCopyFlags uint32

const (
	FileCopyFlags_None               FileCopyFlags = int64(0)
	FileCopyFlags_Overwrite          FileCopyFlags = int64(1)
	FileCopyFlags_Backup             FileCopyFlags = int64(2)
	FileCopyFlags_NofollowSymlinks   FileCopyFlags = int64(4)
	FileCopyFlags_AllMetadata        FileCopyFlags = int64(8)
	FileCopyFlags_NoFallbackForMove  FileCopyFlags = int64(16)
	FileCopyFlags_TargetDefaultPerms FileCopyFlags = int64(32)
)

type FileCreateFlags uint32

const (
	FileCreateFlags_None               FileCreateFlags = int64(0)
	FileCreateFlags_Private            FileCreateFlags = int64(1)
	FileCreateFlags_ReplaceDestination FileCreateFlags = int64(2)
)

type FileMeasureFlags uint32

const (
	FileMeasureFlags_None           FileMeasureFlags = int64(0)
	FileMeasureFlags_ReportAnyError FileMeasureFlags = int64(2)
	FileMeasureFlags_ApparentSize   FileMeasureFlags = int64(4)
	FileMeasureFlags_NoXdev         FileMeasureFlags = int64(8)
)

type FileMonitorFlags uint32

const (
	FileMonitorFlags_None           FileMonitorFlags = int64(0)
	FileMonitorFlags_WatchMounts    FileMonitorFlags = int64(1)
	FileMonitorFlags_SendMoved      FileMonitorFlags = int64(2)
	FileMonitorFlags_WatchHardLinks FileMonitorFlags = int64(4)
	FileMonitorFlags_WatchMoves     FileMonitorFlags = int64(8)
)

type FileQueryInfoFlags uint32

const (
	FileQueryInfoFlags_None             FileQueryInfoFlags = int64(0)
	FileQueryInfoFlags_NofollowSymlinks FileQueryInfoFlags = int64(1)
)

type IOStreamSpliceFlags uint32

const (
	IOStreamSpliceFlags_None         IOStreamSpliceFlags = int64(0)
	IOStreamSpliceFlags_CloseStream1 IOStreamSpliceFlags = int64(1)
	IOStreamSpliceFlags_CloseStream2 IOStreamSpliceFlags = int64(2)
	IOStreamSpliceFlags_WaitForBoth  IOStreamSpliceFlags = int64(4)
)

type MountMountFlags uint32

const (
	MountMountFlags_None MountMountFlags = int64(0)
)

type MountUnmountFlags uint32

const (
	MountUnmountFlags_None  MountUnmountFlags = int64(0)
	MountUnmountFlags_Force MountUnmountFlags = int64(1)
)

type OutputStreamSpliceFlags uint32

const (
	OutputStreamSpliceFlags_None        OutputStreamSpliceFlags = int64(0)
	OutputStreamSpliceFlags_CloseSource OutputStreamSpliceFlags = int64(1)
	OutputStreamSpliceFlags_CloseTarget OutputStreamSpliceFlags = int64(2)
)

type ResourceFlags uint32

const (
	ResourceFlags_None       ResourceFlags = int64(0)
	ResourceFlags_Compressed ResourceFlags = int64(1)
)

type ResourceLookupFlags uint32

const (
	ResourceLookupFlags_None ResourceLookupFlags = int64(0)
)

type SettingsBindFlags uint32

const (
	SettingsBindFlags_Default       SettingsBindFlags = int64(0)
	SettingsBindFlags_Get           SettingsBindFlags = int64(1)
	SettingsBindFlags_Set           SettingsBindFlags = int64(2)
	SettingsBindFlags_NoSensitivity SettingsBindFlags = int64(4)
	SettingsBindFlags_GetNoChanges  SettingsBindFlags = int64(8)
	SettingsBindFlags_InvertBoolean SettingsBindFlags = int64(16)
)

type SocketMsgFlags uint32

const (
	SocketMsgFlags_None      SocketMsgFlags = int64(0)
	SocketMsgFlags_Oob       SocketMsgFlags = int64(1)
	SocketMsgFlags_Peek      SocketMsgFlags = int64(2)
	SocketMsgFlags_Dontroute SocketMsgFlags = int64(4)
)

type SubprocessFlags uint32

const (
	SubprocessFlags_None          SubprocessFlags = int64(0)
	SubprocessFlags_StdinPipe     SubprocessFlags = int64(1)
	SubprocessFlags_StdinInherit  SubprocessFlags = int64(2)
	SubprocessFlags_StdoutPipe    SubprocessFlags = int64(4)
	SubprocessFlags_StdoutSilence SubprocessFlags = int64(8)
	SubprocessFlags_StderrPipe    SubprocessFlags = int64(16)
	SubprocessFlags_StderrSilence SubprocessFlags = int64(32)
	SubprocessFlags_StderrMerge   SubprocessFlags = int64(64)
	SubprocessFlags_InheritFds    SubprocessFlags = int64(128)
)

type TestDBusFlags uint32

const (
	TestDBusFlags_None TestDBusFlags = int64(0)
)

type TlsCertificateFlags uint32

const (
	TlsCertificateFlags_UnknownCa    TlsCertificateFlags = int64(1)
	TlsCertificateFlags_BadIdentity  TlsCertificateFlags = int64(2)
	TlsCertificateFlags_NotActivated TlsCertificateFlags = int64(4)
	TlsCertificateFlags_Expired      TlsCertificateFlags = int64(8)
	TlsCertificateFlags_Revoked      TlsCertificateFlags = int64(16)
	TlsCertificateFlags_Insecure     TlsCertificateFlags = int64(32)
	TlsCertificateFlags_GenericError TlsCertificateFlags = int64(64)
	TlsCertificateFlags_ValidateAll  TlsCertificateFlags = int64(127)
)

type TlsDatabaseVerifyFlags uint32

const (
	TlsDatabaseVerifyFlags_None TlsDatabaseVerifyFlags = int64(0)
)

type TlsPasswordFlags uint32

const (
	TlsPasswordFlags_None      TlsPasswordFlags = int64(0)
	TlsPasswordFlags_Retry     TlsPasswordFlags = int64(2)
	TlsPasswordFlags_ManyTries TlsPasswordFlags = int64(4)
	TlsPasswordFlags_FinalTry  TlsPasswordFlags = int64(8)
)
