// Code generated - DO NOT EDIT.
// +build gio_2.22

package gio

import (
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

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
// #include <gio/gnetworking.h>
import "C"

// aliases

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

// unions

// records
type ActionEntry C.GActionEntry
type AppInfoIface C.GAppInfoIface
type AppLaunchContextClass C.GAppLaunchContextClass
type AppLaunchContextPrivate C.GAppLaunchContextPrivate
type ApplicationCommandLinePrivate C.GApplicationCommandLinePrivate
type ApplicationPrivate C.GApplicationPrivate
type AsyncInitableIface C.GAsyncInitableIface
type AsyncResultIface C.GAsyncResultIface
type BufferedInputStreamClass C.GBufferedInputStreamClass
type BufferedInputStreamPrivate C.GBufferedInputStreamPrivate
type BufferedOutputStreamClass C.GBufferedOutputStreamClass
type BufferedOutputStreamPrivate C.GBufferedOutputStreamPrivate
type CancellableClass C.GCancellableClass
type CancellablePrivate C.GCancellablePrivate
type CharsetConverterClass C.GCharsetConverterClass
type ConverterInputStreamClass C.GConverterInputStreamClass
type ConverterInputStreamPrivate C.GConverterInputStreamPrivate
type ConverterOutputStreamClass C.GConverterOutputStreamClass
type ConverterOutputStreamPrivate C.GConverterOutputStreamPrivate
type DBusInterfaceSkeletonPrivate C.GDBusInterfaceSkeletonPrivate
type DBusObjectManagerClientPrivate C.GDBusObjectManagerClientPrivate
type DBusObjectManagerServerPrivate C.GDBusObjectManagerServerPrivate
type DBusObjectProxyPrivate C.GDBusObjectProxyPrivate
type DBusObjectSkeletonPrivate C.GDBusObjectSkeletonPrivate
type DBusProxyPrivate C.GDBusProxyPrivate
type DataInputStreamClass C.GDataInputStreamClass
type DataInputStreamPrivate C.GDataInputStreamPrivate
type DataOutputStreamClass C.GDataOutputStreamClass
type DataOutputStreamPrivate C.GDataOutputStreamPrivate
type DesktopAppInfoClass C.GDesktopAppInfoClass
type DesktopAppInfoLookupIface C.GDesktopAppInfoLookupIface
type DriveIface C.GDriveIface
type EmblemClass C.GEmblemClass
type EmblemedIconClass C.GEmblemedIconClass
type EmblemedIconPrivate C.GEmblemedIconPrivate
type FileAttributeInfo C.GFileAttributeInfo
type FileAttributeInfoList C.GFileAttributeInfoList
type FileAttributeMatcher C.GFileAttributeMatcher
type FileDescriptorBasedIface C.GFileDescriptorBasedIface
type FileEnumeratorClass C.GFileEnumeratorClass
type FileEnumeratorPrivate C.GFileEnumeratorPrivate
type FileIOStreamClass C.GFileIOStreamClass
type FileIOStreamPrivate C.GFileIOStreamPrivate
type FileIconClass C.GFileIconClass
type FileIface C.GFileIface
type FileInfoClass C.GFileInfoClass
type FileInputStreamClass C.GFileInputStreamClass
type FileInputStreamPrivate C.GFileInputStreamPrivate
type FileMonitorClass C.GFileMonitorClass
type FileMonitorPrivate C.GFileMonitorPrivate
type FileOutputStreamClass C.GFileOutputStreamClass
type FileOutputStreamPrivate C.GFileOutputStreamPrivate
type FilenameCompleterClass C.GFilenameCompleterClass
type FilterInputStreamClass C.GFilterInputStreamClass
type FilterOutputStreamClass C.GFilterOutputStreamClass
type IOExtension C.GIOExtension
type IOExtensionPoint C.GIOExtensionPoint
type IOModuleClass C.GIOModuleClass
type IOSchedulerJob C.GIOSchedulerJob
type IOStreamAdapter C.GIOStreamAdapter
type IOStreamClass C.GIOStreamClass
type IOStreamPrivate C.GIOStreamPrivate
type IconIface C.GIconIface
type InetAddressClass C.GInetAddressClass
type InetAddressMaskClass C.GInetAddressMaskClass
type InetAddressMaskPrivate C.GInetAddressMaskPrivate
type InetAddressPrivate C.GInetAddressPrivate
type InetSocketAddressClass C.GInetSocketAddressClass
type InetSocketAddressPrivate C.GInetSocketAddressPrivate
type InitableIface C.GInitableIface
type InputStreamClass C.GInputStreamClass
type InputStreamPrivate C.GInputStreamPrivate
type InputVector C.GInputVector
type ListStoreClass C.GListStoreClass
type LoadableIconIface C.GLoadableIconIface
type MemoryInputStreamClass C.GMemoryInputStreamClass
type MemoryInputStreamPrivate C.GMemoryInputStreamPrivate
type MemoryOutputStreamClass C.GMemoryOutputStreamClass
type MemoryOutputStreamPrivate C.GMemoryOutputStreamPrivate
type MenuAttributeIterClass C.GMenuAttributeIterClass
type MenuAttributeIterPrivate C.GMenuAttributeIterPrivate
type MenuLinkIterClass C.GMenuLinkIterClass
type MenuLinkIterPrivate C.GMenuLinkIterPrivate
type MenuModelClass C.GMenuModelClass
type MenuModelPrivate C.GMenuModelPrivate
type MountIface C.GMountIface
type MountOperationClass C.GMountOperationClass
type MountOperationPrivate C.GMountOperationPrivate

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
type NativeVolumeMonitorClass C.GNativeVolumeMonitorClass
type NetworkAddressClass C.GNetworkAddressClass
type NetworkAddressPrivate C.GNetworkAddressPrivate
type NetworkServiceClass C.GNetworkServiceClass
type NetworkServicePrivate C.GNetworkServicePrivate
type OutputStreamClass C.GOutputStreamClass
type OutputStreamPrivate C.GOutputStreamPrivate
type OutputVector C.GOutputVector
type PermissionClass C.GPermissionClass
type PermissionPrivate C.GPermissionPrivate
type ProxyAddressEnumeratorClass C.GProxyAddressEnumeratorClass
type ProxyAddressEnumeratorPrivate C.GProxyAddressEnumeratorPrivate
type ProxyAddressPrivate C.GProxyAddressPrivate
type ProxyResolverInterface C.GProxyResolverInterface
type ResolverClass C.GResolverClass
type ResolverPrivate C.GResolverPrivate
type SeekableIface C.GSeekableIface

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type SettingsClass C.GSettingsClass
type SettingsPrivate C.GSettingsPrivate
type SettingsSchemaKey C.GSettingsSchemaKey
type SimpleActionGroupClass C.GSimpleActionGroupClass
type SimpleActionGroupPrivate C.GSimpleActionGroupPrivate
type SimpleAsyncResultClass C.GSimpleAsyncResultClass
type SimpleProxyResolverClass C.GSimpleProxyResolverClass
type SimpleProxyResolverPrivate C.GSimpleProxyResolverPrivate
type SocketAddressClass C.GSocketAddressClass
type SocketAddressEnumeratorClass C.GSocketAddressEnumeratorClass
type SocketClass C.GSocketClass
type SocketClientClass C.GSocketClientClass
type SocketClientPrivate C.GSocketClientPrivate
type SocketConnectableIface C.GSocketConnectableIface
type SocketConnectionClass C.GSocketConnectionClass
type SocketConnectionPrivate C.GSocketConnectionPrivate
type SocketControlMessageClass C.GSocketControlMessageClass
type SocketControlMessagePrivate C.GSocketControlMessagePrivate
type SocketListenerClass C.GSocketListenerClass
type SocketListenerPrivate C.GSocketListenerPrivate
type SocketPrivate C.GSocketPrivate
type SocketServiceClass C.GSocketServiceClass
type SocketServicePrivate C.GSocketServicePrivate
type SrvTarget C.GSrvTarget
type StaticResource C.GStaticResource
type TaskClass C.GTaskClass
type TcpConnectionClass C.GTcpConnectionClass
type TcpConnectionPrivate C.GTcpConnectionPrivate
type TcpWrapperConnectionClass C.GTcpWrapperConnectionClass
type TcpWrapperConnectionPrivate C.GTcpWrapperConnectionPrivate
type ThemedIconClass C.GThemedIconClass
type ThreadedSocketServiceClass C.GThreadedSocketServiceClass
type ThreadedSocketServicePrivate C.GThreadedSocketServicePrivate
type TlsCertificateClass C.GTlsCertificateClass
type TlsCertificatePrivate C.GTlsCertificatePrivate
type TlsConnectionClass C.GTlsConnectionClass
type TlsConnectionPrivate C.GTlsConnectionPrivate
type TlsDatabasePrivate C.GTlsDatabasePrivate
type TlsFileDatabaseInterface C.GTlsFileDatabaseInterface
type TlsInteractionPrivate C.GTlsInteractionPrivate
type TlsPasswordClass C.GTlsPasswordClass
type TlsPasswordPrivate C.GTlsPasswordPrivate
type UnixConnectionClass C.GUnixConnectionClass
type UnixConnectionPrivate C.GUnixConnectionPrivate
type UnixCredentialsMessagePrivate C.GUnixCredentialsMessagePrivate
type UnixFDListClass C.GUnixFDListClass
type UnixFDListPrivate C.GUnixFDListPrivate
type UnixFDMessageClass C.GUnixFDMessageClass
type UnixFDMessagePrivate C.GUnixFDMessagePrivate
type UnixInputStreamClass C.GUnixInputStreamClass
type UnixInputStreamPrivate C.GUnixInputStreamPrivate
type UnixMountEntry C.GUnixMountEntry
type UnixMountMonitorClass C.GUnixMountMonitorClass
type UnixMountPoint C.GUnixMountPoint
type UnixOutputStreamClass C.GUnixOutputStreamClass
type UnixOutputStreamPrivate C.GUnixOutputStreamPrivate
type UnixSocketAddressClass C.GUnixSocketAddressClass
type UnixSocketAddressPrivate C.GUnixSocketAddressPrivate
type VfsClass C.GVfsClass
type VolumeIface C.GVolumeIface
type VolumeMonitorClass C.GVolumeMonitorClass
type ZlibCompressorClass C.GZlibCompressorClass
type ZlibDecompressorClass C.GZlibDecompressorClass

// classes
type AppLaunchContext C.GAppLaunchContext
type ApplicationCommandLine C.GApplicationCommandLine
type BufferedInputStream C.GBufferedInputStream
type BufferedOutputStream C.GBufferedOutputStream
type Cancellable C.GCancellable
type CharsetConverter C.GCharsetConverter
type ConverterInputStream C.GConverterInputStream
type ConverterOutputStream C.GConverterOutputStream
type DBusActionGroup C.GDBusActionGroup
type DBusMenuModel C.GDBusMenuModel
type DataInputStream C.GDataInputStream
type DataOutputStream C.GDataOutputStream
type DesktopAppInfo C.GDesktopAppInfo
type Emblem C.GEmblem
type EmblemedIcon C.GEmblemedIcon
type FileEnumerator C.GFileEnumerator
type FileIOStream C.GFileIOStream
type FileIcon C.GFileIcon
type FileInfo C.GFileInfo
type FileInputStream C.GFileInputStream
type FileMonitor C.GFileMonitor
type FileOutputStream C.GFileOutputStream
type FilenameCompleter C.GFilenameCompleter
type FilterInputStream C.GFilterInputStream
type FilterOutputStream C.GFilterOutputStream
type IOModule C.GIOModule
type IOStream C.GIOStream
type InetAddress C.GInetAddress
type InetSocketAddress C.GInetSocketAddress
type InputStream C.GInputStream
type ListStore C.GListStore
type MemoryInputStream C.GMemoryInputStream
type MemoryOutputStream C.GMemoryOutputStream
type MountOperation C.GMountOperation
type NativeSocketAddress C.GNativeSocketAddress
type NativeVolumeMonitor C.GNativeVolumeMonitor
type NetworkAddress C.GNetworkAddress
type NetworkService C.GNetworkService
type OutputStream C.GOutputStream
type Permission C.GPermission
type ProxyAddressEnumerator C.GProxyAddressEnumerator
type Resolver C.GResolver
type Settings C.GSettings
type SettingsBackend C.GSettingsBackend
type SimpleAction C.GSimpleAction
type SimpleAsyncResult C.GSimpleAsyncResult
type SimplePermission C.GSimplePermission
type Socket C.GSocket
type SocketAddress C.GSocketAddress
type SocketAddressEnumerator C.GSocketAddressEnumerator
type SocketClient C.GSocketClient
type SocketConnection C.GSocketConnection
type SocketControlMessage C.GSocketControlMessage
type SocketListener C.GSocketListener
type SocketService C.GSocketService
type Task C.GTask
type TcpConnection C.GTcpConnection
type ThemedIcon C.GThemedIcon
type ThreadedSocketService C.GThreadedSocketService
type UnixConnection C.GUnixConnection
type UnixFDList C.GUnixFDList
type UnixFDMessage C.GUnixFDMessage
type UnixInputStream C.GUnixInputStream
type UnixMountMonitor C.GUnixMountMonitor
type UnixOutputStream C.GUnixOutputStream
type UnixSocketAddress C.GUnixSocketAddress
type Vfs C.GVfs
type VolumeMonitor C.GVolumeMonitor
type ZlibCompressor C.GZlibCompressor
type ZlibDecompressor C.GZlibDecompressor

// interfaces
type Action C.GAction
type ActionGroup C.GActionGroup
type AppInfo C.GAppInfo
type AsyncInitable C.GAsyncInitable
type AsyncResult C.GAsyncResult
type DBusObject C.GDBusObject
type DBusObjectManager C.GDBusObjectManager
type DesktopAppInfoLookup C.GDesktopAppInfoLookup
type Drive C.GDrive
type File C.GFile
type Icon C.GIcon
type Initable C.GInitable
type ListModel C.GListModel
type LoadableIcon C.GLoadableIcon
type Mount C.GMount
type Seekable C.GSeekable
type SocketConnectable C.GSocketConnectable
type Volume C.GVolume

func Fn_app_info_create_from_commandline(commandline string, applicationName string, flags AppInfoCreateFlags) {
}

func Fn_app_info_get_all() {
	C.g_app_info_get_all()
}

func Fn_app_info_get_all_for_type(contentType string) {}

func Fn_app_info_get_default_for_type(contentType string, mustSupportUris bool) {}

func Fn_app_info_get_default_for_uri_scheme(uriScheme string) {}

func Fn_app_info_launch_default_for_uri(uri string, context *AppLaunchContext) {}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_app_info_reset_type_associations(contentType string) {}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_content_type_can_be_executable(type_ string) {}

func Fn_content_type_equals(type1 string, type2 string) {}

func Fn_content_type_from_mime_type(mimeType string) {}

func Fn_content_type_get_description(type_ string) {}

func Fn_content_type_get_icon(type_ string) {}

func Fn_content_type_get_mime_type(type_ string) {}

func Fn_content_type_guess(filename string, data *uint8, dataSize uint64) {}

func Fn_content_type_guess_for_tree(root *File) {}

func Fn_content_type_is_a(type_ string, supertype string) {}

func Fn_content_type_is_unknown(type_ string) {}

func Fn_content_types_get_registered() {
	C.g_content_types_get_registered()
}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_dbus_error_quark() {
	C.g_dbus_error_quark()
}

func Fn_file_new_for_commandline_arg(arg string) {}

func Fn_file_new_for_path(path string) {}

func Fn_file_new_for_uri(uri string) {}

func Fn_file_parse_name(parseName string) {}

func Fn_icon_hash(icon unsafe.Pointer) {}

func Fn_icon_new_for_string(str string) {}

func Fn_initable_newv(objectType glib.Type, nParameters uint, parameters *gobject.Parameter, cancellable *Cancellable) {
}

func Fn_io_error_from_errno(errNo int) {}

func Fn_io_error_quark() {
	C.g_io_error_quark()
}

func Fn_io_extension_point_implement(extensionPointName string, type_ glib.Type, extensionName string, priority int) {
}

func Fn_io_extension_point_lookup(name string) {}

func Fn_io_extension_point_register(name string) {}

func Fn_io_modules_load_all_in_directory(dirname string) {}

func Fn_io_scheduler_cancel_all_jobs() {
	C.g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : io_scheduler_push_job : has callback

// UNSUPPORTED : keyfile_settings_backend_new : blacklisted
// UNSUPPORTED : memory_settings_backend_new : blacklisted
// UNSUPPORTED : null_settings_backend_new : blacklisted
func Fn_resolver_error_quark() {
	C.g_resolver_error_quark()
}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_srv_target_list_sort(targets *glib.List) {}

func Fn_unix_is_mount_path_system_internal(mountPath string) {}

func Fn_unix_mount_at(mountPath string) {}

func Fn_unix_mount_compare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) {}

func Fn_unix_mount_free(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_device_path(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_fs_type(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_mount_path(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_can_eject(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_icon(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_name(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_should_display(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_is_readonly(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_is_system_internal(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_points_changed_since(time uint64) {}

func Fn_unix_mount_points_get() {}

func Fn_unix_mounts_changed_since(time uint64) {}

func Fn_unix_mounts_get() {}
