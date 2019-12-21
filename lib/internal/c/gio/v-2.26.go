// Code generated - DO NOT EDIT.
// +build gio_2.26

package gio

import (
	c "github.com/pekim/gobbi/lib/internal/c"
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
import "C"

// bitfields
type AppInfoCreateFlags C.GAppInfoCreateFlags
type AskPasswordFlags C.GAskPasswordFlags
type BusNameOwnerFlags C.GBusNameOwnerFlags
type BusNameWatcherFlags C.GBusNameWatcherFlags
type ConverterFlags C.GConverterFlags
type DBusCallFlags C.GDBusCallFlags
type DBusCapabilityFlags C.GDBusCapabilityFlags
type DBusConnectionFlags C.GDBusConnectionFlags
type DBusMessageFlags C.GDBusMessageFlags
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
type FileMonitorFlags C.GFileMonitorFlags
type FileQueryInfoFlags C.GFileQueryInfoFlags
type MountMountFlags C.GMountMountFlags
type MountUnmountFlags C.GMountUnmountFlags
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags
type SettingsBindFlags C.GSettingsBindFlags
type SocketMsgFlags C.GSocketMsgFlags

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
type MountOperationResult C.GMountOperationResult
type PasswordSave C.GPasswordSave
type ResolverError C.GResolverError
type SocketFamily C.GSocketFamily
type SocketProtocol C.GSocketProtocol
type SocketType C.GSocketType
type UnixSocketAddressType C.GUnixSocketAddressType
type ZlibCompressorFormat C.GZlibCompressorFormat

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
type ConverterIface C.GConverterIface
type ConverterInputStreamClass C.GConverterInputStreamClass
type ConverterInputStreamPrivate C.GConverterInputStreamPrivate
type ConverterOutputStreamClass C.GConverterOutputStreamClass
type ConverterOutputStreamPrivate C.GConverterOutputStreamPrivate
type CredentialsClass C.GCredentialsClass
type DBusAnnotationInfo C.GDBusAnnotationInfo
type DBusArgInfo C.GDBusArgInfo
type DBusErrorEntry C.GDBusErrorEntry
type DBusInterfaceInfo C.GDBusInterfaceInfo
type DBusInterfaceSkeletonPrivate C.GDBusInterfaceSkeletonPrivate
type DBusInterfaceVTable C.GDBusInterfaceVTable
type DBusMethodInfo C.GDBusMethodInfo
type DBusNodeInfo C.GDBusNodeInfo
type DBusObjectManagerClientPrivate C.GDBusObjectManagerClientPrivate
type DBusObjectManagerServerPrivate C.GDBusObjectManagerServerPrivate
type DBusObjectProxyPrivate C.GDBusObjectProxyPrivate
type DBusObjectSkeletonPrivate C.GDBusObjectSkeletonPrivate
type DBusPropertyInfo C.GDBusPropertyInfo
type DBusProxyClass C.GDBusProxyClass
type DBusProxyPrivate C.GDBusProxyPrivate
type DBusSignalInfo C.GDBusSignalInfo
type DBusSubtreeVTable C.GDBusSubtreeVTable
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
type ProxyAddressClass C.GProxyAddressClass
type ProxyAddressEnumeratorClass C.GProxyAddressEnumeratorClass
type ProxyAddressEnumeratorPrivate C.GProxyAddressEnumeratorPrivate
type ProxyAddressPrivate C.GProxyAddressPrivate
type ProxyInterface C.GProxyInterface
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
type TlsClientConnectionInterface C.GTlsClientConnectionInterface
type TlsConnectionClass C.GTlsConnectionClass
type TlsConnectionPrivate C.GTlsConnectionPrivate
type TlsDatabasePrivate C.GTlsDatabasePrivate
type TlsFileDatabaseInterface C.GTlsFileDatabaseInterface
type TlsInteractionPrivate C.GTlsInteractionPrivate
type TlsPasswordClass C.GTlsPasswordClass
type TlsPasswordPrivate C.GTlsPasswordPrivate
type TlsServerConnectionInterface C.GTlsServerConnectionInterface
type UnixConnectionClass C.GUnixConnectionClass
type UnixConnectionPrivate C.GUnixConnectionPrivate
type UnixCredentialsMessageClass C.GUnixCredentialsMessageClass
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
type Credentials C.GCredentials
type DBusActionGroup C.GDBusActionGroup
type DBusAuthObserver C.GDBusAuthObserver
type DBusConnection C.GDBusConnection
type DBusMenuModel C.GDBusMenuModel
type DBusMessage C.GDBusMessage
type DBusMethodInvocation C.GDBusMethodInvocation
type DBusProxy C.GDBusProxy
type DBusServer C.GDBusServer
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
type ProxyAddress C.GProxyAddress
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
type UnixCredentialsMessage C.GUnixCredentialsMessage
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
type Converter C.GConverter
type DBusObject C.GDBusObject
type DBusObjectManager C.GDBusObjectManager
type DesktopAppInfoLookup C.GDesktopAppInfoLookup
type Drive C.GDrive
type File C.GFile
type FileDescriptorBased C.GFileDescriptorBased
type Icon C.GIcon
type Initable C.GInitable
type ListModel C.GListModel
type LoadableIcon C.GLoadableIcon
type Mount C.GMount
type Proxy C.GProxy
type ProxyResolver C.GProxyResolver
type Seekable C.GSeekable
type SocketConnectable C.GSocketConnectable
type Volume C.GVolume

func Fn_app_info_create_from_commandline(commandline c.UndefinedParamType, applicationName c.UndefinedParamType, flags c.UndefinedParamType) {
}

func Fn_app_info_get_all() {}

func Fn_app_info_get_all_for_type(contentType c.UndefinedParamType) {}

func Fn_app_info_get_default_for_type(contentType c.UndefinedParamType, mustSupportUris bool) {}

func Fn_app_info_get_default_for_uri_scheme(uriScheme c.UndefinedParamType) {}

func Fn_app_info_launch_default_for_uri(uri c.UndefinedParamType, context c.UndefinedParamType) {}

func Fn_app_info_reset_type_associations(contentType c.UndefinedParamType) {}

func Fn_async_initable_newv_async(objectType c.UndefinedParamType, nParameters uint, parameters c.UndefinedParamType, ioPriority int, cancellable c.UndefinedParamType, callback c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_bus_get(busType c.UndefinedParamType, cancellable c.UndefinedParamType, callback c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_bus_get_finish(res c.UndefinedParamType) {}

func Fn_bus_get_sync(busType c.UndefinedParamType, cancellable c.UndefinedParamType) {}

func Fn_bus_own_name(busType c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, busAcquiredHandler c.UndefinedParamType, nameAcquiredHandler c.UndefinedParamType, nameLostHandler c.UndefinedParamType, userData unsafe.Pointer, userDataFreeFunc c.UndefinedParamType) {
}

func Fn_bus_own_name_on_connection(connection c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAcquiredHandler c.UndefinedParamType, nameLostHandler c.UndefinedParamType, userData unsafe.Pointer, userDataFreeFunc c.UndefinedParamType) {
}

func Fn_bus_own_name_on_connection_with_closures(connection c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAcquiredClosure c.UndefinedParamType, nameLostClosure c.UndefinedParamType) {
}

func Fn_bus_own_name_with_closures(busType c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, busAcquiredClosure c.UndefinedParamType, nameAcquiredClosure c.UndefinedParamType, nameLostClosure c.UndefinedParamType) {
}

func Fn_bus_unown_name(ownerId uint) {}

func Fn_bus_unwatch_name(watcherId uint) {}

func Fn_bus_watch_name(busType c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAppearedHandler c.UndefinedParamType, nameVanishedHandler c.UndefinedParamType, userData unsafe.Pointer, userDataFreeFunc c.UndefinedParamType) {
}

func Fn_bus_watch_name_on_connection(connection c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAppearedHandler c.UndefinedParamType, nameVanishedHandler c.UndefinedParamType, userData unsafe.Pointer, userDataFreeFunc c.UndefinedParamType) {
}

func Fn_bus_watch_name_on_connection_with_closures(connection c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAppearedClosure c.UndefinedParamType, nameVanishedClosure c.UndefinedParamType) {
}

func Fn_bus_watch_name_with_closures(busType c.UndefinedParamType, name c.UndefinedParamType, flags c.UndefinedParamType, nameAppearedClosure c.UndefinedParamType, nameVanishedClosure c.UndefinedParamType) {
}

func Fn_content_type_can_be_executable(type_ c.UndefinedParamType) {}

func Fn_content_type_equals(type1 c.UndefinedParamType, type2 c.UndefinedParamType) {}

func Fn_content_type_from_mime_type(mimeType c.UndefinedParamType) {}

func Fn_content_type_get_description(type_ c.UndefinedParamType) {}

func Fn_content_type_get_icon(type_ c.UndefinedParamType) {}

func Fn_content_type_get_mime_type(type_ c.UndefinedParamType) {}

func Fn_content_type_guess(filename c.UndefinedParamType, data c.UndefinedParamType, dataSize uint64) {
}

func Fn_content_type_guess_for_tree(root c.UndefinedParamType) {}

func Fn_content_type_is_a(type_ c.UndefinedParamType, supertype c.UndefinedParamType) {}

func Fn_content_type_is_unknown(type_ c.UndefinedParamType) {}

func Fn_content_types_get_registered() {}

func Fn_dbus_address_get_for_bus_sync(busType c.UndefinedParamType, cancellable c.UndefinedParamType) {
}

func Fn_dbus_address_get_stream(address c.UndefinedParamType, cancellable c.UndefinedParamType, callback c.UndefinedParamType, userData unsafe.Pointer) {
}

func Fn_dbus_address_get_stream_finish(res c.UndefinedParamType) {}

func Fn_dbus_address_get_stream_sync(address c.UndefinedParamType, cancellable c.UndefinedParamType) {}

func Fn_dbus_annotation_info_lookup(annotations c.UndefinedParamType, name c.UndefinedParamType) {}

func Fn_dbus_error_encode_gerror(error c.UndefinedParamType) {}

func Fn_dbus_error_get_remote_error(error c.UndefinedParamType) {}

func Fn_dbus_error_is_remote_error(error c.UndefinedParamType) {}

func Fn_dbus_error_new_for_dbus_error(dbusErrorName c.UndefinedParamType, dbusErrorMessage c.UndefinedParamType) {
}

func Fn_dbus_error_quark() {}

func Fn_dbus_error_register_error(errorDomain c.UndefinedParamType, errorCode int, dbusErrorName c.UndefinedParamType) {
}

func Fn_dbus_error_register_error_domain(errorDomainQuarkName c.UndefinedParamType, quarkVolatile c.UndefinedParamType, entries c.UndefinedParamType, numEntries uint) {
}

func Fn_dbus_error_strip_remote_error(error c.UndefinedParamType) {}

func Fn_dbus_error_unregister_error(errorDomain c.UndefinedParamType, errorCode int, dbusErrorName c.UndefinedParamType) {
}

func Fn_dbus_generate_guid() {}

func Fn_dbus_is_address(string_ c.UndefinedParamType) {}

func Fn_dbus_is_guid(string_ c.UndefinedParamType) {}

func Fn_dbus_is_interface_name(string_ c.UndefinedParamType) {}

func Fn_dbus_is_member_name(string_ c.UndefinedParamType) {}

func Fn_dbus_is_name(string_ c.UndefinedParamType) {}

func Fn_dbus_is_supported_address(string_ c.UndefinedParamType) {}

func Fn_dbus_is_unique_name(string_ c.UndefinedParamType) {}

func Fn_file_new_for_commandline_arg(arg c.UndefinedParamType) {}

func Fn_file_new_for_path(path c.UndefinedParamType) {}

func Fn_file_new_for_uri(uri c.UndefinedParamType) {}

func Fn_file_parse_name(parseName c.UndefinedParamType) {}

func Fn_icon_hash(icon c.UndefinedParamType) {}

func Fn_icon_new_for_string(str c.UndefinedParamType) {}

func Fn_initable_newv(objectType c.UndefinedParamType, nParameters uint, parameters c.UndefinedParamType, cancellable c.UndefinedParamType) {
}

func Fn_io_error_from_errno(errNo int) {}

func Fn_io_error_quark() {}

func Fn_io_extension_point_implement(extensionPointName c.UndefinedParamType, type_ c.UndefinedParamType, extensionName c.UndefinedParamType, priority int) {
}

func Fn_io_extension_point_lookup(name c.UndefinedParamType) {}

func Fn_io_extension_point_register(name c.UndefinedParamType) {}

func Fn_io_modules_load_all_in_directory(dirname c.UndefinedParamType) {}

func Fn_io_modules_scan_all_in_directory(dirname c.UndefinedParamType) {}

func Fn_io_scheduler_cancel_all_jobs() {}

func Fn_io_scheduler_push_job(jobFunc c.UndefinedParamType, userData unsafe.Pointer, notify c.UndefinedParamType, ioPriority int, cancellable c.UndefinedParamType) {
}

func Fn_keyfile_settings_backend_new(filename c.UndefinedParamType, rootPath c.UndefinedParamType, rootGroup c.UndefinedParamType) {
}

func Fn_proxy_get_default_for_protocol(protocol c.UndefinedParamType) {}

func Fn_proxy_resolver_get_default() {}

func Fn_resolver_error_quark() {}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

func Fn_simple_async_report_gerror_in_idle(object c.UndefinedParamType, callback c.UndefinedParamType, userData unsafe.Pointer, error c.UndefinedParamType) {
}

func Fn_srv_target_list_sort(targets c.UndefinedParamType) {}

func Fn_unix_is_mount_path_system_internal(mountPath c.UndefinedParamType) {}

func Fn_unix_mount_at(mountPath c.UndefinedParamType) {}

func Fn_unix_mount_compare(mount1 c.UndefinedParamType, mount2 c.UndefinedParamType) {}

func Fn_unix_mount_free(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_get_device_path(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_get_fs_type(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_get_mount_path(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_guess_can_eject(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_guess_icon(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_guess_name(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_guess_should_display(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_is_readonly(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_is_system_internal(mountEntry c.UndefinedParamType) {}

func Fn_unix_mount_points_changed_since(time uint64) {}

func Fn_unix_mount_points_get() {}

func Fn_unix_mounts_changed_since(time uint64) {}

func Fn_unix_mounts_get() {}
