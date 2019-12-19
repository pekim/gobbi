// Code generated - DO NOT EDIT.
// +build gio_2.24

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
type BusNameOwnerFlags c.UnusupportedByVersion
type BusNameWatcherFlags c.UnusupportedByVersion
type ConverterFlags C.GConverterFlags
type DBusCallFlags c.UnusupportedByVersion
type DBusCapabilityFlags c.UnusupportedByVersion
type DBusConnectionFlags c.UnusupportedByVersion
type DBusInterfaceSkeletonFlags c.UnusupportedByVersion
type DBusMessageFlags c.UnusupportedByVersion
type DBusObjectManagerClientFlags c.UnusupportedByVersion
type DBusPropertyInfoFlags c.UnusupportedByVersion
type DBusProxyFlags c.UnusupportedByVersion
type DBusSendMessageFlags c.UnusupportedByVersion
type DBusServerFlags c.UnusupportedByVersion
type DBusSignalFlags c.UnusupportedByVersion
type DBusSubtreeFlags c.UnusupportedByVersion
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
type BusType c.UnusupportedByVersion
type ConverterResult C.GConverterResult
type CredentialsType c.UnusupportedByVersion
type DBusError c.UnusupportedByVersion
type DBusMessageByteOrder c.UnusupportedByVersion
type DBusMessageHeaderField c.UnusupportedByVersion
type DBusMessageType c.UnusupportedByVersion
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
type UnixSocketAddressType c.UnusupportedByVersion
type ZlibCompressorFormat C.GZlibCompressorFormat
