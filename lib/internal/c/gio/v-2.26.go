// Code generated - DO NOT EDIT.
// +build gio_2.26

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
type ApplicationFlags c.UnusupportedByVersion
type AskPasswordFlags C.GAskPasswordFlags
type BusNameOwnerFlags C.GBusNameOwnerFlags
type BusNameWatcherFlags C.GBusNameWatcherFlags
type ConverterFlags C.GConverterFlags
type DBusCallFlags C.GDBusCallFlags
type DBusCapabilityFlags C.GDBusCapabilityFlags
type DBusConnectionFlags C.GDBusConnectionFlags
type DBusInterfaceSkeletonFlags c.UnusupportedByVersion
type DBusMessageFlags C.GDBusMessageFlags
type DBusObjectManagerClientFlags c.UnusupportedByVersion
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
type FileMeasureFlags c.UnusupportedByVersion
type FileMonitorFlags C.GFileMonitorFlags
type FileQueryInfoFlags C.GFileQueryInfoFlags
type IOStreamSpliceFlags c.UnusupportedByVersion
type MountMountFlags C.GMountMountFlags
type MountUnmountFlags C.GMountUnmountFlags
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags
type ResolverNameLookupFlags c.UnusupportedByVersion
type ResourceFlags c.UnusupportedByVersion
type ResourceLookupFlags c.UnusupportedByVersion
type SettingsBindFlags C.GSettingsBindFlags
type SocketMsgFlags C.GSocketMsgFlags
type SubprocessFlags c.UnusupportedByVersion
type TestDBusFlags c.UnusupportedByVersion
type TlsCertificateFlags c.UnusupportedByVersion
type TlsDatabaseVerifyFlags c.UnusupportedByVersion
type TlsPasswordFlags c.UnusupportedByVersion

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
type IOModuleScopeFlags c.UnusupportedByVersion
type MountOperationResult C.GMountOperationResult
type NetworkConnectivity c.UnusupportedByVersion
type NotificationPriority c.UnusupportedByVersion
type PasswordSave C.GPasswordSave
type PollableReturn c.UnusupportedByVersion
type ResolverError C.GResolverError
type ResolverRecordType c.UnusupportedByVersion
type ResourceError c.UnusupportedByVersion
type SocketClientEvent c.UnusupportedByVersion
type SocketFamily C.GSocketFamily
type SocketListenerEvent c.UnusupportedByVersion
type SocketProtocol C.GSocketProtocol
type SocketType C.GSocketType
type TlsAuthenticationMode c.UnusupportedByVersion
type TlsCertificateRequestFlags c.UnusupportedByVersion
type TlsDatabaseLookupFlags c.UnusupportedByVersion
type TlsError c.UnusupportedByVersion
type TlsInteractionResult c.UnusupportedByVersion
type TlsRehandshakeMode c.UnusupportedByVersion
type UnixSocketAddressType C.GUnixSocketAddressType
type ZlibCompressorFormat C.GZlibCompressorFormat
