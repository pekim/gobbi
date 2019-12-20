// Code generated - DO NOT EDIT.
// +build !gio_2.18,!gio_2.20,!gio_2.22,!gio_2.24,!gio_2.26,!gio_2.28,!gio_2.30,!gio_2.32,!gio_2.34,!gio_2.36,!gio_2.38,!gio_2.40,!gio_2.42,!gio_2.44,!gio_2.46,!gio_2.52,!gio_2.58,!gio_2.60

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
type FileAttributeInfoFlags C.GFileAttributeInfoFlags
type FileCopyFlags C.GFileCopyFlags
type FileCreateFlags C.GFileCreateFlags
type FileMonitorFlags C.GFileMonitorFlags
type FileQueryInfoFlags C.GFileQueryInfoFlags
type MountMountFlags C.GMountMountFlags
type MountUnmountFlags C.GMountUnmountFlags
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags
type SettingsBindFlags C.GSettingsBindFlags

// enumerations
type DataStreamByteOrder C.GDataStreamByteOrder
type DataStreamNewlineType C.GDataStreamNewlineType
type FileAttributeStatus C.GFileAttributeStatus
type FileAttributeType C.GFileAttributeType
type FileMonitorEvent C.GFileMonitorEvent
type FileType C.GFileType
type FilesystemPreviewType C.GFilesystemPreviewType
type IOErrorEnum C.GIOErrorEnum
type MountOperationResult C.GMountOperationResult
type PasswordSave C.GPasswordSave
