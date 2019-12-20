// Code generated - DO NOT EDIT.
// +build gio_2.22

package gio

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
type AskPasswordFlags C.GAskPasswordFlags
type DriveStartFlags C.GDriveStartFlags
type FileAttributeInfoFlags C.GFileAttributeInfoFlags
type FileCopyFlags C.GFileCopyFlags
type FileCreateFlags C.GFileCreateFlags
type FileMonitorFlags C.GFileMonitorFlags
type FileQueryInfoFlags C.GFileQueryInfoFlags
type MountMountFlags C.GMountMountFlags
type MountUnmountFlags C.GMountUnmountFlags
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags
type SettingsBindFlags C.GSettingsBindFlags
type SocketMsgFlags C.GSocketMsgFlags

// enumerations
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
type MountOperationResult C.GMountOperationResult
type PasswordSave C.GPasswordSave
type ResolverError C.GResolverError
type SocketFamily C.GSocketFamily
type SocketProtocol C.GSocketProtocol
type SocketType C.GSocketType
