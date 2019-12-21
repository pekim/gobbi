// Code generated - DO NOT EDIT.
// +build gio_2.60

package gio

import (
	c "github.com/pekim/gobbi/lib/internal/c"
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
import "C"

// aliases

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
type ResolverNameLookupFlags C.GResolverNameLookupFlags
type ResourceFlags C.GResourceFlags
type ResourceLookupFlags C.GResourceLookupFlags
type SettingsBindFlags C.GSettingsBindFlags
type SocketMsgFlags C.GSocketMsgFlags
type SubprocessFlags C.GSubprocessFlags
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
type NetworkConnectivity C.GNetworkConnectivity
type NotificationPriority C.GNotificationPriority
type PasswordSave C.GPasswordSave
type PollableReturn C.GPollableReturn
type ResolverError C.GResolverError
type ResolverRecordType C.GResolverRecordType
type ResourceError C.GResourceError
type SocketClientEvent C.GSocketClientEvent
type SocketFamily C.GSocketFamily
type SocketListenerEvent C.GSocketListenerEvent
type SocketProtocol C.GSocketProtocol
type SocketType C.GSocketType
type TlsAuthenticationMode C.GTlsAuthenticationMode
type TlsCertificateRequestFlags C.GTlsCertificateRequestFlags
type TlsDatabaseLookupFlags C.GTlsDatabaseLookupFlags
type TlsError C.GTlsError
type TlsInteractionResult C.GTlsInteractionResult
type TlsRehandshakeMode C.GTlsRehandshakeMode
type UnixSocketAddressType C.GUnixSocketAddressType
type ZlibCompressorFormat C.GZlibCompressorFormat

// records
type ActionEntry C.GActionEntry
type ActionGroupInterface C.GActionGroupInterface
type ActionInterface C.GActionInterface
type ActionMapInterface C.GActionMapInterface
type AppInfoIface C.GAppInfoIface
type AppLaunchContextClass C.GAppLaunchContextClass
type AppLaunchContextPrivate C.GAppLaunchContextPrivate
type ApplicationClass C.GApplicationClass
type ApplicationCommandLineClass C.GApplicationCommandLineClass
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
type DBusInterfaceIface C.GDBusInterfaceIface
type DBusInterfaceInfo C.GDBusInterfaceInfo
type DBusInterfaceSkeletonClass C.GDBusInterfaceSkeletonClass
type DBusInterfaceSkeletonPrivate C.GDBusInterfaceSkeletonPrivate
type DBusInterfaceVTable C.GDBusInterfaceVTable
type DBusMethodInfo C.GDBusMethodInfo
type DBusNodeInfo C.GDBusNodeInfo
type DBusObjectIface C.GDBusObjectIface
type DBusObjectManagerClientClass C.GDBusObjectManagerClientClass
type DBusObjectManagerClientPrivate C.GDBusObjectManagerClientPrivate
type DBusObjectManagerIface C.GDBusObjectManagerIface
type DBusObjectManagerServerClass C.GDBusObjectManagerServerClass
type DBusObjectManagerServerPrivate C.GDBusObjectManagerServerPrivate
type DBusObjectProxyClass C.GDBusObjectProxyClass
type DBusObjectProxyPrivate C.GDBusObjectProxyPrivate
type DBusObjectSkeletonClass C.GDBusObjectSkeletonClass
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
type DatagramBasedInterface C.GDatagramBasedInterface
type DesktopAppInfoClass C.GDesktopAppInfoClass
type DesktopAppInfoLookupIface C.GDesktopAppInfoLookupIface
type DriveIface C.GDriveIface
type DtlsClientConnectionInterface C.GDtlsClientConnectionInterface
type DtlsConnectionInterface C.GDtlsConnectionInterface
type DtlsServerConnectionInterface C.GDtlsServerConnectionInterface
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
type IOModuleScope C.GIOModuleScope
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
type InputMessage C.GInputMessage
type InputStreamClass C.GInputStreamClass
type InputStreamPrivate C.GInputStreamPrivate
type InputVector C.GInputVector
type ListModelInterface C.GListModelInterface
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
type NetworkMonitorInterface C.GNetworkMonitorInterface
type NetworkServiceClass C.GNetworkServiceClass
type NetworkServicePrivate C.GNetworkServicePrivate
type OutputMessage C.GOutputMessage
type OutputStreamClass C.GOutputStreamClass
type OutputStreamPrivate C.GOutputStreamPrivate
type OutputVector C.GOutputVector
type PermissionClass C.GPermissionClass
type PermissionPrivate C.GPermissionPrivate
type PollableInputStreamInterface C.GPollableInputStreamInterface
type PollableOutputStreamInterface C.GPollableOutputStreamInterface
type ProxyAddressClass C.GProxyAddressClass
type ProxyAddressEnumeratorClass C.GProxyAddressEnumeratorClass
type ProxyAddressEnumeratorPrivate C.GProxyAddressEnumeratorPrivate
type ProxyAddressPrivate C.GProxyAddressPrivate
type ProxyInterface C.GProxyInterface
type ProxyResolverInterface C.GProxyResolverInterface
type RemoteActionGroupInterface C.GRemoteActionGroupInterface
type ResolverClass C.GResolverClass
type ResolverPrivate C.GResolverPrivate
type Resource C.GResource
type SeekableIface C.GSeekableIface

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type SettingsClass C.GSettingsClass
type SettingsPrivate C.GSettingsPrivate
type SettingsSchema C.GSettingsSchema
type SettingsSchemaKey C.GSettingsSchemaKey
type SettingsSchemaSource C.GSettingsSchemaSource
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
type TlsBackendInterface C.GTlsBackendInterface
type TlsCertificateClass C.GTlsCertificateClass
type TlsCertificatePrivate C.GTlsCertificatePrivate
type TlsClientConnectionInterface C.GTlsClientConnectionInterface
type TlsConnectionClass C.GTlsConnectionClass
type TlsConnectionPrivate C.GTlsConnectionPrivate
type TlsDatabaseClass C.GTlsDatabaseClass
type TlsDatabasePrivate C.GTlsDatabasePrivate
type TlsFileDatabaseInterface C.GTlsFileDatabaseInterface
type TlsInteractionClass C.GTlsInteractionClass
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
type AppInfoMonitor C.GAppInfoMonitor
type AppLaunchContext C.GAppLaunchContext
type Application C.GApplication
type ApplicationCommandLine C.GApplicationCommandLine
type BufferedInputStream C.GBufferedInputStream
type BufferedOutputStream C.GBufferedOutputStream
type BytesIcon C.GBytesIcon
type Cancellable C.GCancellable
type CharsetConverter C.GCharsetConverter
type ConverterInputStream C.GConverterInputStream
type ConverterOutputStream C.GConverterOutputStream
type Credentials C.GCredentials
type DBusActionGroup C.GDBusActionGroup
type DBusAuthObserver C.GDBusAuthObserver
type DBusConnection C.GDBusConnection
type DBusInterfaceSkeleton C.GDBusInterfaceSkeleton
type DBusMenuModel C.GDBusMenuModel
type DBusMessage C.GDBusMessage
type DBusMethodInvocation C.GDBusMethodInvocation
type DBusObjectManagerClient C.GDBusObjectManagerClient
type DBusObjectManagerServer C.GDBusObjectManagerServer
type DBusObjectProxy C.GDBusObjectProxy
type DBusObjectSkeleton C.GDBusObjectSkeleton
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
type InetAddressMask C.GInetAddressMask
type InetSocketAddress C.GInetSocketAddress
type InputStream C.GInputStream
type ListStore C.GListStore
type MemoryInputStream C.GMemoryInputStream
type MemoryOutputStream C.GMemoryOutputStream
type Menu C.GMenu
type MenuAttributeIter C.GMenuAttributeIter
type MenuItem C.GMenuItem
type MenuLinkIter C.GMenuLinkIter
type MenuModel C.GMenuModel
type MountOperation C.GMountOperation
type NativeSocketAddress C.GNativeSocketAddress
type NativeVolumeMonitor C.GNativeVolumeMonitor
type NetworkAddress C.GNetworkAddress
type NetworkService C.GNetworkService
type Notification C.GNotification
type OutputStream C.GOutputStream
type Permission C.GPermission
type PropertyAction C.GPropertyAction
type ProxyAddress C.GProxyAddress
type ProxyAddressEnumerator C.GProxyAddressEnumerator
type Resolver C.GResolver
type Settings C.GSettings
type SettingsBackend C.GSettingsBackend
type SimpleAction C.GSimpleAction
type SimpleActionGroup C.GSimpleActionGroup
type SimpleAsyncResult C.GSimpleAsyncResult
type SimpleIOStream C.GSimpleIOStream
type SimplePermission C.GSimplePermission
type SimpleProxyResolver C.GSimpleProxyResolver
type Socket C.GSocket
type SocketAddress C.GSocketAddress
type SocketAddressEnumerator C.GSocketAddressEnumerator
type SocketClient C.GSocketClient
type SocketConnection C.GSocketConnection
type SocketControlMessage C.GSocketControlMessage
type SocketListener C.GSocketListener
type SocketService C.GSocketService
type Subprocess C.GSubprocess
type SubprocessLauncher C.GSubprocessLauncher
type Task C.GTask
type TcpConnection C.GTcpConnection
type TcpWrapperConnection C.GTcpWrapperConnection
type TestDBus C.GTestDBus
type ThemedIcon C.GThemedIcon
type ThreadedSocketService C.GThreadedSocketService
type TlsCertificate C.GTlsCertificate
type TlsConnection C.GTlsConnection
type TlsDatabase C.GTlsDatabase
type TlsInteraction C.GTlsInteraction
type TlsPassword C.GTlsPassword
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
type ActionMap C.GActionMap
type AppInfo C.GAppInfo
type AsyncInitable C.GAsyncInitable
type AsyncResult C.GAsyncResult
type Converter C.GConverter
type DBusInterface C.GDBusInterface
type DBusObject C.GDBusObject
type DBusObjectManager C.GDBusObjectManager
type DatagramBased C.GDatagramBased
type DesktopAppInfoLookup C.GDesktopAppInfoLookup
type Drive C.GDrive
type DtlsClientConnection C.GDtlsClientConnection
type DtlsConnection C.GDtlsConnection
type DtlsServerConnection C.GDtlsServerConnection
type File C.GFile
type FileDescriptorBased C.GFileDescriptorBased
type Icon C.GIcon
type Initable C.GInitable
type ListModel C.GListModel
type LoadableIcon C.GLoadableIcon
type Mount C.GMount
type NetworkMonitor C.GNetworkMonitor
type PollableInputStream C.GPollableInputStream
type PollableOutputStream C.GPollableOutputStream
type Proxy C.GProxy
type ProxyResolver C.GProxyResolver
type RemoteActionGroup C.GRemoteActionGroup
type Seekable C.GSeekable
type SocketConnectable C.GSocketConnectable
type TlsBackend C.GTlsBackend
type TlsClientConnection C.GTlsClientConnection
type TlsFileDatabase C.GTlsFileDatabase
type TlsServerConnection C.GTlsServerConnection
type Volume C.GVolume

func Fn_action_name_is_valid(actionName string) {}

func Fn_action_parse_detailed_name(detailedName string) {}

func Fn_action_print_detailed_name(actionName string, targetValue *glib.Variant) {}

func Fn_app_info_create_from_commandline(commandline string, applicationName string, flags AppInfoCreateFlags) {
}

func Fn_app_info_get_all() {}

func Fn_app_info_get_all_for_type(contentType string) {}

func Fn_app_info_get_default_for_type(contentType string, mustSupportUris bool) {}

func Fn_app_info_get_default_for_uri_scheme(uriScheme string) {}

func Fn_app_info_get_fallback_for_type(contentType string) {}

func Fn_app_info_get_recommended_for_type(contentType string) {}

func Fn_app_info_launch_default_for_uri(uri string, context *AppLaunchContext) {}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_app_info_launch_default_for_uri_finish(result *AsyncResult) {}

func Fn_app_info_reset_type_associations(contentType string) {}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

func Fn_bus_get_finish(res *AsyncResult) {}

func Fn_bus_get_sync(busType BusType, cancellable *Cancellable) {}

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

func Fn_bus_own_name_on_connection_with_closures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) {
}

func Fn_bus_own_name_with_closures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) {
}

func Fn_bus_unown_name(ownerId uint) {}

func Fn_bus_unwatch_name(watcherId uint) {}

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_bus_watch_name_on_connection_with_closures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) {
}

func Fn_bus_watch_name_with_closures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) {
}

func Fn_content_type_can_be_executable(type_ string) {}

func Fn_content_type_equals(type1 string, type2 string) {}

func Fn_content_type_from_mime_type(mimeType string) {}

func Fn_content_type_get_description(type_ string) {}

func Fn_content_type_get_generic_icon_name(type_ string) {}

func Fn_content_type_get_icon(type_ string) {}

func Fn_content_type_get_mime_dirs() {}

func Fn_content_type_get_mime_type(type_ string) {}

func Fn_content_type_get_symbolic_icon(type_ string) {}

func Fn_content_type_guess(filename string, data c.UndefinedParamType, dataSize uint64) {}

func Fn_content_type_guess_for_tree(root *File) {}

func Fn_content_type_is_a(type_ string, supertype string) {}

func Fn_content_type_is_mime_type(type_ string, mimeType string) {}

func Fn_content_type_is_unknown(type_ string) {}

func Fn_content_type_set_mime_dirs(dirs c.UndefinedParamType) {}

func Fn_content_types_get_registered() {}

func Fn_dbus_address_escape_value(string_ string) {}

func Fn_dbus_address_get_for_bus_sync(busType BusType, cancellable *Cancellable) {}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_dbus_address_get_stream_finish(res *AsyncResult) {}

func Fn_dbus_address_get_stream_sync(address string, cancellable *Cancellable) {}

func Fn_dbus_annotation_info_lookup(annotations c.UndefinedParamType, name string) {}

func Fn_dbus_error_encode_gerror(error *glib.Error) {}

func Fn_dbus_error_get_remote_error(error *glib.Error) {}

func Fn_dbus_error_is_remote_error(error *glib.Error) {}

func Fn_dbus_error_new_for_dbus_error(dbusErrorName string, dbusErrorMessage string) {}

func Fn_dbus_error_quark() {}

func Fn_dbus_error_register_error(errorDomain glib.Quark, errorCode int, dbusErrorName string) {}

func Fn_dbus_error_register_error_domain(errorDomainQuarkName string, quarkVolatile *uint64, entries c.UndefinedParamType, numEntries uint) {
}

func Fn_dbus_error_strip_remote_error(error *glib.Error) {}

func Fn_dbus_error_unregister_error(errorDomain glib.Quark, errorCode int, dbusErrorName string) {}

func Fn_dbus_generate_guid() {}

func Fn_dbus_gvalue_to_gvariant(gvalue *gobject.Value, type_ *glib.VariantType) {}

func Fn_dbus_gvariant_to_gvalue(value *glib.Variant) {}

func Fn_dbus_is_address(string_ string) {}

func Fn_dbus_is_guid(string_ string) {}

func Fn_dbus_is_interface_name(string_ string) {}

func Fn_dbus_is_member_name(string_ string) {}

func Fn_dbus_is_name(string_ string) {}

func Fn_dbus_is_supported_address(string_ string) {}

func Fn_dbus_is_unique_name(string_ string) {}

func Fn_dtls_client_connection_new(baseSocket *DatagramBased, serverIdentity *SocketConnectable) {}

func Fn_dtls_server_connection_new(baseSocket *DatagramBased, certificate *TlsCertificate) {}

func Fn_file_new_for_commandline_arg(arg string) {}

func Fn_file_new_for_commandline_arg_and_cwd(arg string, cwd string) {}

func Fn_file_new_for_path(path string) {}

func Fn_file_new_for_uri(uri string) {}

func Fn_file_new_tmp(tmpl string) {}

func Fn_file_parse_name(parseName string) {}

func Fn_icon_deserialize(value *glib.Variant) {}

func Fn_icon_hash(icon unsafe.Pointer) {}

func Fn_icon_new_for_string(str string) {}

func Fn_initable_newv(objectType glib.Type, nParameters uint, parameters c.UndefinedParamType, cancellable *Cancellable) {
}

func Fn_io_error_from_errno(errNo int) {}

func Fn_io_error_quark() {}

func Fn_io_extension_point_implement(extensionPointName string, type_ glib.Type, extensionName string, priority int) {
}

func Fn_io_extension_point_lookup(name string) {}

func Fn_io_extension_point_register(name string) {}

func Fn_io_modules_load_all_in_directory(dirname string) {}

func Fn_io_modules_load_all_in_directory_with_scope(dirname string, scope *IOModuleScope) {}

func Fn_io_modules_scan_all_in_directory(dirname string) {}

func Fn_io_modules_scan_all_in_directory_with_scope(dirname string, scope *IOModuleScope) {}

func Fn_io_scheduler_cancel_all_jobs() {}

// UNSUPPORTED : io_scheduler_push_job : has callback

func Fn_keyfile_settings_backend_new(filename string, rootPath string, rootGroup string) {}

func Fn_memory_settings_backend_new() {}

func Fn_network_monitor_get_default() {}

func Fn_networking_init() {}

func Fn_null_settings_backend_new() {}

func Fn_pollable_source_new(pollableStream *gobject.Object) {}

func Fn_pollable_source_new_full(pollableStream unsafe.Pointer, childSource *glib.Source, cancellable *Cancellable) {
}

func Fn_pollable_stream_read(stream *InputStream, buffer c.UndefinedParamType, count uint64, blocking bool, cancellable *Cancellable) {
}

func Fn_pollable_stream_write(stream *OutputStream, buffer c.UndefinedParamType, count uint64, blocking bool, cancellable *Cancellable) {
}

func Fn_pollable_stream_write_all(stream *OutputStream, buffer c.UndefinedParamType, count uint64, blocking bool, cancellable *Cancellable) {
}

func Fn_proxy_get_default_for_protocol(protocol string) {}

func Fn_proxy_resolver_get_default() {}

func Fn_resolver_error_quark() {}

func Fn_resource_error_quark() {}

func Fn_resource_load(filename string) {}

func Fn_resources_enumerate_children(path string, lookupFlags ResourceLookupFlags) {}

func Fn_resources_get_info(path string, lookupFlags ResourceLookupFlags) {}

func Fn_resources_lookup_data(path string, lookupFlags ResourceLookupFlags) {}

func Fn_resources_open_stream(path string, lookupFlags ResourceLookupFlags) {}

func Fn_resources_register(resource *Resource) {}

func Fn_resources_unregister(resource *Resource) {}

func Fn_settings_schema_source_get_default() {}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_srv_target_list_sort(targets *glib.List) {}

func Fn_tls_backend_get_default() {}

func Fn_tls_client_connection_new(baseIoStream *IOStream, serverIdentity *SocketConnectable) {}

func Fn_tls_error_quark() {}

func Fn_tls_file_database_new(anchors string) {}

func Fn_tls_server_connection_new(baseIoStream *IOStream, certificate *TlsCertificate) {}

func Fn_unix_is_mount_path_system_internal(mountPath string) {}

func Fn_unix_is_system_device_path(devicePath string) {}

func Fn_unix_is_system_fs_type(fsType string) {}

func Fn_unix_mount_at(mountPath string) {}

func Fn_unix_mount_compare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) {}

func Fn_unix_mount_copy(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_for(filePath string) {}

func Fn_unix_mount_free(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_device_path(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_fs_type(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_mount_path(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_options(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_get_root_path(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_can_eject(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_icon(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_name(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_should_display(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_guess_symbolic_icon(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_is_readonly(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_is_system_internal(mountEntry *UnixMountEntry) {}

func Fn_unix_mount_points_changed_since(time uint64) {}

func Fn_unix_mount_points_get() {}

func Fn_unix_mounts_changed_since(time uint64) {}

func Fn_unix_mounts_get() {}
