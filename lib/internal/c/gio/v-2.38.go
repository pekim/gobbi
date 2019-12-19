// Code generated - DO NOT EDIT.
// +build gio_2.38

package gio

import c "github.com/pekim/gobbi/lib/internal/c"

// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

// bitfields
type AppInfoCreateFlags C.GAppInfoCreateFlags
type ApplicationFlags C.GApplicationFlags
type AskPasswordFlags C.GAskPasswordFlags
type BusNameOwnerFlags C.GBusNameOwnerFlags
type BusNameWatcherFlags C.GBusNameWatcherFlags
type ConverterFlags C.GConverterFlags
type DBusCallFlags C.GDBusCallFlags
type DBusCapabilityFlags C.GDBusCapabilityFlags
type DBusConnectionFlags C.GDBusConnectionFlags
type DBusInterfaceSkeletonFlags C.GDBusInterfaceSkeletonFlags
type DBusMessageFlags C.GDBusMessageFlags
type DBusObjectManagerClientFlags C.GDBusObjectManagerClientFlags
type DBusPropertyInfoFlags C.GDBusPropertyInfoFlags
type DBusProxyFlags C.GDBusProxyFlags
type DBusSendMessageFlags C.GDBusSendMessageFlags
type DBusServerFlags C.GDBusServerFlags
type DBusSignalFlags C.GDBusSignalFlags
type DBusSubtreeFlags C.GDBusSubtreeFlags
type DriveStartFlags C.GDriveStartFlags
type FileAttributeInfoFlags C.GFileAttributeInfoFlags
type FileCopyFlags C.GFileCopyFlags
type FileCreateFlags C.GFileCreateFlags
type FileMeasureFlags C.GFileMeasureFlags
type FileMonitorFlags C.GFileMonitorFlags
type FileQueryInfoFlags C.GFileQueryInfoFlags
type IOStreamSpliceFlags C.GIOStreamSpliceFlags
type MountMountFlags C.GMountMountFlags
type MountUnmountFlags C.GMountUnmountFlags
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags
type ResolverNameLookupFlags c.UnusupportedByVersion
type ResourceFlags C.GResourceFlags
type ResourceLookupFlags C.GResourceLookupFlags
type SettingsBindFlags C.GSettingsBindFlags
type SocketMsgFlags C.GSocketMsgFlags
type SubprocessFlags c.UnusupportedByVersion
type TestDBusFlags C.GTestDBusFlags
type TlsCertificateFlags C.GTlsCertificateFlags
type TlsDatabaseVerifyFlags C.GTlsDatabaseVerifyFlags
type TlsPasswordFlags C.GTlsPasswordFlags

// enumerations
type BusType C.GBusType
type ConverterResult C.GConverterResult
type CredentialsType C.GCredentialsType
type DBusError C.GDBusError
type DBusMessageByteOrder C.GDBusMessageByteOrder
type DBusMessageHeaderField C.GDBusMessageHeaderField
type DBusMessageType C.GDBusMessageType
type DataStreamByteOrder C.GDataStreamByteOrder
type DataStreamNewlineType C.GDataStreamNewlineType
type DriveStartStopType C.GDriveStartStopType
type EmblemOrigin C.GEmblemOrigin
type FileAttributeStatus C.GFileAttributeStatus
type FileAttributeType C.GFileAttributeType
type FileMonitorEvent C.GFileMonitorEvent
type FileType C.GFileType
type FilesystemPreviewType C.GFilesystemPreviewType
type IOErrorEnum C.GIOErrorEnum
type IOModuleScopeFlags C.GIOModuleScopeFlags
type MountOperationResult C.GMountOperationResult
type NetworkConnectivity c.UnusupportedByVersion
type NotificationPriority c.UnusupportedByVersion
type PasswordSave C.GPasswordSave
type PollableReturn c.UnusupportedByVersion
type ResolverError C.GResolverError
type ResolverRecordType C.GResolverRecordType
type ResourceError C.GResourceError
type SocketClientEvent C.GSocketClientEvent
type SocketFamily C.GSocketFamily
type SocketListenerEvent c.UnusupportedByVersion
type SocketProtocol C.GSocketProtocol
type SocketType C.GSocketType
type TlsAuthenticationMode C.GTlsAuthenticationMode
type TlsCertificateRequestFlags c.UnusupportedByVersion
type TlsDatabaseLookupFlags C.GTlsDatabaseLookupFlags
type TlsError C.GTlsError
type TlsInteractionResult C.GTlsInteractionResult
type TlsRehandshakeMode C.GTlsRehandshakeMode
type UnixSocketAddressType C.GUnixSocketAddressType
type ZlibCompressorFormat C.GZlibCompressorFormat
