// Code generated - DO NOT EDIT.

package gio

// Appinfocreateflags is a representation of the C type AppInfoCreateFlags.
type Appinfocreateflags int

const (
	Appinfocreateflags_None                        Appinfocreateflags = 0
	Appinfocreateflags_NeedsTerminal               Appinfocreateflags = 1
	Appinfocreateflags_SupportsUris                Appinfocreateflags = 2
	Appinfocreateflags_SupportsStartupNotification Appinfocreateflags = 4
)

// Applicationflags is a representation of the C type ApplicationFlags.
//
// since 2.28
type Applicationflags int

const (
	Applicationflags_FlagsNone          Applicationflags = 0
	Applicationflags_IsService          Applicationflags = 1
	Applicationflags_IsLauncher         Applicationflags = 2
	Applicationflags_HandlesOpen        Applicationflags = 4
	Applicationflags_HandlesCommandLine Applicationflags = 8
	Applicationflags_SendEnvironment    Applicationflags = 16
	Applicationflags_NonUnique          Applicationflags = 32
	Applicationflags_CanOverrideAppId   Applicationflags = 64
)

// Askpasswordflags is a representation of the C type AskPasswordFlags.
type Askpasswordflags int

const (
	Askpasswordflags_NeedPassword       Askpasswordflags = 1
	Askpasswordflags_NeedUsername       Askpasswordflags = 2
	Askpasswordflags_NeedDomain         Askpasswordflags = 4
	Askpasswordflags_SavingSupported    Askpasswordflags = 8
	Askpasswordflags_AnonymousSupported Askpasswordflags = 16
)

// Busnameownerflags is a representation of the C type BusNameOwnerFlags.
//
// since 2.26
type Busnameownerflags int

const (
	Busnameownerflags_None             Busnameownerflags = 0
	Busnameownerflags_AllowReplacement Busnameownerflags = 1
	Busnameownerflags_Replace          Busnameownerflags = 2
	Busnameownerflags_DoNotQueue       Busnameownerflags = 4
)

// Busnamewatcherflags is a representation of the C type BusNameWatcherFlags.
//
// since 2.26
type Busnamewatcherflags int

const (
	Busnamewatcherflags_None      Busnamewatcherflags = 0
	Busnamewatcherflags_AutoStart Busnamewatcherflags = 1
)

// Converterflags is a representation of the C type ConverterFlags.
//
// since 2.24
type Converterflags int

const (
	Converterflags_None       Converterflags = 0
	Converterflags_InputAtEnd Converterflags = 1
	Converterflags_Flush      Converterflags = 2
)

// Dbuscallflags is a representation of the C type DBusCallFlags.
//
// since 2.26
type Dbuscallflags int

const (
	Dbuscallflags_None                          Dbuscallflags = 0
	Dbuscallflags_NoAutoStart                   Dbuscallflags = 1
	Dbuscallflags_AllowInteractiveAuthorization Dbuscallflags = 2
)

// Dbuscapabilityflags is a representation of the C type DBusCapabilityFlags.
//
// since 2.26
type Dbuscapabilityflags int

const (
	Dbuscapabilityflags_None          Dbuscapabilityflags = 0
	Dbuscapabilityflags_UnixFdPassing Dbuscapabilityflags = 1
)

// Dbusconnectionflags is a representation of the C type DBusConnectionFlags.
//
// since 2.26
type Dbusconnectionflags int

const (
	Dbusconnectionflags_None                         Dbusconnectionflags = 0
	Dbusconnectionflags_AuthenticationClient         Dbusconnectionflags = 1
	Dbusconnectionflags_AuthenticationServer         Dbusconnectionflags = 2
	Dbusconnectionflags_AuthenticationAllowAnonymous Dbusconnectionflags = 4
	Dbusconnectionflags_MessageBusConnection         Dbusconnectionflags = 8
	Dbusconnectionflags_DelayMessageProcessing       Dbusconnectionflags = 16
)

// Dbusinterfaceskeletonflags is a representation of the C type DBusInterfaceSkeletonFlags.
//
// since 2.30
type Dbusinterfaceskeletonflags int

const (
	Dbusinterfaceskeletonflags_None                            Dbusinterfaceskeletonflags = 0
	Dbusinterfaceskeletonflags_HandleMethodInvocationsInThread Dbusinterfaceskeletonflags = 1
)

// Dbusmessageflags is a representation of the C type DBusMessageFlags.
//
// since 2.26
type Dbusmessageflags int

const (
	Dbusmessageflags_None                          Dbusmessageflags = 0
	Dbusmessageflags_NoReplyExpected               Dbusmessageflags = 1
	Dbusmessageflags_NoAutoStart                   Dbusmessageflags = 2
	Dbusmessageflags_AllowInteractiveAuthorization Dbusmessageflags = 4
)

// Dbusobjectmanagerclientflags is a representation of the C type DBusObjectManagerClientFlags.
//
// since 2.30
type Dbusobjectmanagerclientflags int

const (
	Dbusobjectmanagerclientflags_None           Dbusobjectmanagerclientflags = 0
	Dbusobjectmanagerclientflags_DoNotAutoStart Dbusobjectmanagerclientflags = 1
)

// Dbuspropertyinfoflags is a representation of the C type DBusPropertyInfoFlags.
//
// since 2.26
type Dbuspropertyinfoflags int

const (
	Dbuspropertyinfoflags_None     Dbuspropertyinfoflags = 0
	Dbuspropertyinfoflags_Readable Dbuspropertyinfoflags = 1
	Dbuspropertyinfoflags_Writable Dbuspropertyinfoflags = 2
)

// Dbusproxyflags is a representation of the C type DBusProxyFlags.
//
// since 2.26
type Dbusproxyflags int

const (
	Dbusproxyflags_None                         Dbusproxyflags = 0
	Dbusproxyflags_DoNotLoadProperties          Dbusproxyflags = 1
	Dbusproxyflags_DoNotConnectSignals          Dbusproxyflags = 2
	Dbusproxyflags_DoNotAutoStart               Dbusproxyflags = 4
	Dbusproxyflags_GetInvalidatedProperties     Dbusproxyflags = 8
	Dbusproxyflags_DoNotAutoStartAtConstruction Dbusproxyflags = 16
)

// Dbussendmessageflags is a representation of the C type DBusSendMessageFlags.
//
// since 2.26
type Dbussendmessageflags int

const (
	Dbussendmessageflags_None           Dbussendmessageflags = 0
	Dbussendmessageflags_PreserveSerial Dbussendmessageflags = 1
)

// Dbusserverflags is a representation of the C type DBusServerFlags.
//
// since 2.26
type Dbusserverflags int

const (
	Dbusserverflags_None                         Dbusserverflags = 0
	Dbusserverflags_RunInThread                  Dbusserverflags = 1
	Dbusserverflags_AuthenticationAllowAnonymous Dbusserverflags = 2
)

// Dbussignalflags is a representation of the C type DBusSignalFlags.
//
// since 2.26
type Dbussignalflags int

const (
	Dbussignalflags_None               Dbussignalflags = 0
	Dbussignalflags_NoMatchRule        Dbussignalflags = 1
	Dbussignalflags_MatchArg0Namespace Dbussignalflags = 2
	Dbussignalflags_MatchArg0Path      Dbussignalflags = 4
)

// Dbussubtreeflags is a representation of the C type DBusSubtreeFlags.
//
// since 2.26
type Dbussubtreeflags int

const (
	Dbussubtreeflags_None                        Dbussubtreeflags = 0
	Dbussubtreeflags_DispatchToUnenumeratedNodes Dbussubtreeflags = 1
)

// Drivestartflags is a representation of the C type DriveStartFlags.
//
// since 2.22
type Drivestartflags int

const (
	Drivestartflags_None Drivestartflags = 0
)

// Fileattributeinfoflags is a representation of the C type FileAttributeInfoFlags.
type Fileattributeinfoflags int

const (
	Fileattributeinfoflags_None          Fileattributeinfoflags = 0
	Fileattributeinfoflags_CopyWithFile  Fileattributeinfoflags = 1
	Fileattributeinfoflags_CopyWhenMoved Fileattributeinfoflags = 2
)

// Filecopyflags is a representation of the C type FileCopyFlags.
type Filecopyflags int

const (
	Filecopyflags_None               Filecopyflags = 0
	Filecopyflags_Overwrite          Filecopyflags = 1
	Filecopyflags_Backup             Filecopyflags = 2
	Filecopyflags_NofollowSymlinks   Filecopyflags = 4
	Filecopyflags_AllMetadata        Filecopyflags = 8
	Filecopyflags_NoFallbackForMove  Filecopyflags = 16
	Filecopyflags_TargetDefaultPerms Filecopyflags = 32
)

// Filecreateflags is a representation of the C type FileCreateFlags.
type Filecreateflags int

const (
	Filecreateflags_None               Filecreateflags = 0
	Filecreateflags_Private            Filecreateflags = 1
	Filecreateflags_ReplaceDestination Filecreateflags = 2
)

// Filemeasureflags is a representation of the C type FileMeasureFlags.
//
// since 2.38
type Filemeasureflags int

const (
	Filemeasureflags_None           Filemeasureflags = 0
	Filemeasureflags_ReportAnyError Filemeasureflags = 2
	Filemeasureflags_ApparentSize   Filemeasureflags = 4
	Filemeasureflags_NoXdev         Filemeasureflags = 8
)

// Filemonitorflags is a representation of the C type FileMonitorFlags.
type Filemonitorflags int

const (
	Filemonitorflags_None           Filemonitorflags = 0
	Filemonitorflags_WatchMounts    Filemonitorflags = 1
	Filemonitorflags_SendMoved      Filemonitorflags = 2
	Filemonitorflags_WatchHardLinks Filemonitorflags = 4
	Filemonitorflags_WatchMoves     Filemonitorflags = 8
)

// Filequeryinfoflags is a representation of the C type FileQueryInfoFlags.
type Filequeryinfoflags int

const (
	Filequeryinfoflags_None             Filequeryinfoflags = 0
	Filequeryinfoflags_NofollowSymlinks Filequeryinfoflags = 1
)

// Iostreamspliceflags is a representation of the C type IOStreamSpliceFlags.
//
// since 2.28
type Iostreamspliceflags int

const (
	Iostreamspliceflags_None         Iostreamspliceflags = 0
	Iostreamspliceflags_CloseStream1 Iostreamspliceflags = 1
	Iostreamspliceflags_CloseStream2 Iostreamspliceflags = 2
	Iostreamspliceflags_WaitForBoth  Iostreamspliceflags = 4
)

// Mountmountflags is a representation of the C type MountMountFlags.
type Mountmountflags int

const (
	Mountmountflags_None Mountmountflags = 0
)

// Mountunmountflags is a representation of the C type MountUnmountFlags.
type Mountunmountflags int

const (
	Mountunmountflags_None  Mountunmountflags = 0
	Mountunmountflags_Force Mountunmountflags = 1
)

// Outputstreamspliceflags is a representation of the C type OutputStreamSpliceFlags.
type Outputstreamspliceflags int

const (
	Outputstreamspliceflags_None        Outputstreamspliceflags = 0
	Outputstreamspliceflags_CloseSource Outputstreamspliceflags = 1
	Outputstreamspliceflags_CloseTarget Outputstreamspliceflags = 2
)

// Resourceflags is a representation of the C type ResourceFlags.
//
// since 2.32
type Resourceflags int

const (
	Resourceflags_None       Resourceflags = 0
	Resourceflags_Compressed Resourceflags = 1
)

// Resourcelookupflags is a representation of the C type ResourceLookupFlags.
//
// since 2.32
type Resourcelookupflags int

const (
	Resourcelookupflags_None Resourcelookupflags = 0
)

// Settingsbindflags is a representation of the C type SettingsBindFlags.
type Settingsbindflags int

const (
	Settingsbindflags_Default       Settingsbindflags = 0
	Settingsbindflags_Get           Settingsbindflags = 1
	Settingsbindflags_Set           Settingsbindflags = 2
	Settingsbindflags_NoSensitivity Settingsbindflags = 4
	Settingsbindflags_GetNoChanges  Settingsbindflags = 8
	Settingsbindflags_InvertBoolean Settingsbindflags = 16
)

// Socketmsgflags is a representation of the C type SocketMsgFlags.
//
// since 2.22
type Socketmsgflags int

const (
	Socketmsgflags_None      Socketmsgflags = 0
	Socketmsgflags_Oob       Socketmsgflags = 1
	Socketmsgflags_Peek      Socketmsgflags = 2
	Socketmsgflags_Dontroute Socketmsgflags = 4
)

// Subprocessflags is a representation of the C type SubprocessFlags.
//
// since 2.40
type Subprocessflags int

const (
	Subprocessflags_None          Subprocessflags = 0
	Subprocessflags_StdinPipe     Subprocessflags = 1
	Subprocessflags_StdinInherit  Subprocessflags = 2
	Subprocessflags_StdoutPipe    Subprocessflags = 4
	Subprocessflags_StdoutSilence Subprocessflags = 8
	Subprocessflags_StderrPipe    Subprocessflags = 16
	Subprocessflags_StderrSilence Subprocessflags = 32
	Subprocessflags_StderrMerge   Subprocessflags = 64
	Subprocessflags_InheritFds    Subprocessflags = 128
)

// Testdbusflags is a representation of the C type TestDBusFlags.
//
// since 2.34
type Testdbusflags int

const (
	Testdbusflags_None Testdbusflags = 0
)

// Tlscertificateflags is a representation of the C type TlsCertificateFlags.
//
// since 2.28
type Tlscertificateflags int

const (
	Tlscertificateflags_UnknownCa    Tlscertificateflags = 1
	Tlscertificateflags_BadIdentity  Tlscertificateflags = 2
	Tlscertificateflags_NotActivated Tlscertificateflags = 4
	Tlscertificateflags_Expired      Tlscertificateflags = 8
	Tlscertificateflags_Revoked      Tlscertificateflags = 16
	Tlscertificateflags_Insecure     Tlscertificateflags = 32
	Tlscertificateflags_GenericError Tlscertificateflags = 64
	Tlscertificateflags_ValidateAll  Tlscertificateflags = 127
)

// Tlsdatabaseverifyflags is a representation of the C type TlsDatabaseVerifyFlags.
//
// since 2.30
type Tlsdatabaseverifyflags int

const (
	Tlsdatabaseverifyflags_None Tlsdatabaseverifyflags = 0
)

// Tlspasswordflags is a representation of the C type TlsPasswordFlags.
//
// since 2.30
type Tlspasswordflags int

const (
	Tlspasswordflags_None      Tlspasswordflags = 0
	Tlspasswordflags_Retry     Tlspasswordflags = 2
	Tlspasswordflags_ManyTries Tlspasswordflags = 4
	Tlspasswordflags_FinalTry  Tlspasswordflags = 8
)
