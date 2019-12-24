// Code generated - DO NOT EDIT.
// +build gio_2.38

package gio

import (
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
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

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
type NetworkMonitorInterface C.GNetworkMonitorInterface
type NetworkServiceClass C.GNetworkServiceClass
type NetworkServicePrivate C.GNetworkServicePrivate
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

func Fn_g_action_name_is_valid(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_action_name_is_valid(cValue0)
}

func Fn_g_action_parse_detailed_name(param0 string, param1 *string, param2 *unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	var cValue1String *C.gchar
	cValue1 := &cValue1String
	cValue2 := (**C.GVariant)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_action_parse_detailed_name(cValue0, cValue1, cValue2, cError)
	param1String := C.GoString(cValue1String)
	*param1 = param1String
}

func Fn_g_action_print_detailed_name(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_action_print_detailed_name(cValue0, cValue1)
}

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GAppInfoCreateFlags)(param2)
	cError := (**C.GError)(error)

	C.g_app_info_create_from_commandline(cValue0, cValue1, cValue2, cError)
}

func Fn_g_app_info_get_all() {

	C.g_app_info_get_all()
}

func Fn_g_app_info_get_all_for_type(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_get_all_for_type(cValue0)
}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := toCBool(param1)

	C.g_app_info_get_default_for_type(cValue0, cValue1)
}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_get_default_for_uri_scheme(cValue0)
}

func Fn_g_app_info_get_fallback_for_type(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_get_fallback_for_type(cValue0)
}

func Fn_g_app_info_get_recommended_for_type(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_get_recommended_for_type(cValue0)
}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_app_info_launch_default_for_uri(cValue0, cValue1, cError)
}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_g_app_info_reset_type_associations(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_reset_type_associations(cValue0)
}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

func Fn_g_bus_get_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_bus_get_finish(cValue0, cError)
}

func Fn_g_bus_get_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (C.GBusType)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_bus_get_sync(cValue0, cValue1, cError)
}

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

func Fn_g_bus_own_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GBusNameOwnerFlags)(param2)
	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	C.g_bus_own_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_g_bus_own_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
	cValue0 := (C.GBusType)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GBusNameOwnerFlags)(param2)
	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))
	cValue5 := (*C.GClosure)(unsafe.Pointer(param5))

	C.g_bus_own_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_g_bus_unown_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unown_name(cValue0)
}

func Fn_g_bus_unwatch_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unwatch_name(cValue0)
}

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_g_bus_watch_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GBusNameWatcherFlags)(param2)
	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	C.g_bus_watch_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_g_bus_watch_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValue0 := (C.GBusType)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GBusNameWatcherFlags)(param2)
	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))
	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	C.g_bus_watch_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_g_content_type_can_be_executable(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_can_be_executable(cValue0)
}

func Fn_g_content_type_equals(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_content_type_equals(cValue0, cValue1)
}

func Fn_g_content_type_from_mime_type(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_from_mime_type(cValue0)
}

func Fn_g_content_type_get_description(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_get_description(cValue0)
}

func Fn_g_content_type_get_generic_icon_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_get_generic_icon_name(cValue0)
}

func Fn_g_content_type_get_icon(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_get_icon(cValue0)
}

func Fn_g_content_type_get_mime_type(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_get_mime_type(cValue0)
}

func Fn_g_content_type_get_symbolic_icon(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_get_symbolic_icon(cValue0)
}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) {
	// has array param
}

func Fn_g_content_type_guess_for_tree(param0 unsafe.Pointer) {
	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.g_content_type_guess_for_tree(cValue0)
}

func Fn_g_content_type_is_a(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_content_type_is_a(cValue0, cValue1)
}

func Fn_g_content_type_is_unknown(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_content_type_is_unknown(cValue0)
}

func Fn_g_content_types_get_registered() {

	C.g_content_types_get_registered()
}

func Fn_g_dbus_address_escape_value(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_address_escape_value(cValue0)
}

func Fn_g_dbus_address_get_for_bus_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (C.GBusType)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_dbus_address_get_for_bus_sync(cValue0, cValue1, cError)
}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_g_dbus_address_get_stream_finish(param0 unsafe.Pointer, param1 *string, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	var cValue1String *C.gchar
	cValue1 := &cValue1String
	cError := (**C.GError)(error)

	C.g_dbus_address_get_stream_finish(cValue0, cValue1, cError)
	param1String := C.GoString(cValue1String)
	*param1 = param1String
}

func Fn_g_dbus_address_get_stream_sync(param0 string, param1 *string, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	var cValue1String *C.gchar
	cValue1 := &cValue1String
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_dbus_address_get_stream_sync(cValue0, cValue1, cValue2, cError)
	param1String := C.GoString(cValue1String)
	*param1 = param1String
}

func Fn_g_dbus_annotation_info_lookup(param0 []unsafe.Pointer, param1 string) {
	// has array param
}

func Fn_g_dbus_error_encode_gerror(param0 unsafe.Pointer) {
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_error_encode_gerror(cValue0)
}

func Fn_g_dbus_error_get_remote_error(param0 unsafe.Pointer) {
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_error_get_remote_error(cValue0)
}

func Fn_g_dbus_error_is_remote_error(param0 unsafe.Pointer) {
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_error_is_remote_error(cValue0)
}

func Fn_g_dbus_error_new_for_dbus_error(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_error_new_for_dbus_error(cValue0, cValue1)
}

func Fn_g_dbus_error_quark() {

	C.g_dbus_error_quark()
}

func Fn_g_dbus_error_register_error(param0 uint32, param1 int, param2 string) {
	cValue0 := (C.GQuark)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_error_register_error(cValue0, cValue1, cValue2)
}

func Fn_g_dbus_error_register_error_domain(param0 string, param1 *uint64, param2 []DBusErrorEntry, param3 uint) {
	// has array param
}

func Fn_g_dbus_error_strip_remote_error(param0 unsafe.Pointer) {
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_error_strip_remote_error(cValue0)
}

func Fn_g_dbus_error_unregister_error(param0 uint32, param1 int, param2 string) {
	cValue0 := (C.GQuark)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_error_unregister_error(cValue0, cValue1, cValue2)
}

func Fn_g_dbus_generate_guid() {

	C.g_dbus_generate_guid()
}

func Fn_g_dbus_gvalue_to_gvariant(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))
	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	C.g_dbus_gvalue_to_gvariant(cValue0, cValue1)
}

func Fn_g_dbus_gvariant_to_gvalue(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))
	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_dbus_gvariant_to_gvalue(cValue0, cValue1)
}

func Fn_g_dbus_is_address(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_address(cValue0)
}

func Fn_g_dbus_is_guid(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_guid(cValue0)
}

func Fn_g_dbus_is_interface_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_interface_name(cValue0)
}

func Fn_g_dbus_is_member_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_member_name(cValue0)
}

func Fn_g_dbus_is_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_name(cValue0)
}

func Fn_g_dbus_is_supported_address(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_dbus_is_supported_address(cValue0, cError)
}

func Fn_g_dbus_is_unique_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_is_unique_name(cValue0)
}

func Fn_g_file_new_for_commandline_arg(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_new_for_commandline_arg(cValue0)
}

func Fn_g_file_new_for_commandline_arg_and_cwd(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_file_new_for_commandline_arg_and_cwd(cValue0, cValue1)
}

func Fn_g_file_new_for_path(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_new_for_path(cValue0)
}

func Fn_g_file_new_for_uri(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_new_for_uri(cValue0)
}

func Fn_g_file_new_tmp(param0 string, param1 *unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (**C.GFileIOStream)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_file_new_tmp(cValue0, cValue1, cError)
}

func Fn_g_file_parse_name(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_parse_name(cValue0)
}

func Fn_g_icon_deserialize(param0 unsafe.Pointer) {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_icon_deserialize(cValue0)
}

func Fn_g_icon_hash(param0 unsafe.Pointer) {
	cValue0 := (C.gconstpointer)(param0)

	C.g_icon_hash(cValue0)
}

func Fn_g_icon_new_for_string(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_icon_new_for_string(cValue0, cError)
}

func Fn_g_initable_newv(param0 uint64, param1 uint, param2 []gobject.Parameter, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_io_error_from_errno(param0 int) {
	cValue0 := (C.gint)(param0)

	C.g_io_error_from_errno(cValue0)
}

func Fn_g_io_error_quark() {

	C.g_io_error_quark()
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GType)(param1)
	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (C.gint)(param3)

	C.g_io_extension_point_implement(cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_io_extension_point_lookup(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_io_extension_point_lookup(cValue0)
}

func Fn_g_io_extension_point_register(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_io_extension_point_register(cValue0)
}

// UNSUPPORTED : io_modules_load_all_in_directory : blacklisted
// UNSUPPORTED : io_modules_load_all_in_directory_with_scope : blacklisted
// UNSUPPORTED : io_modules_scan_all_in_directory : blacklisted
// UNSUPPORTED : io_modules_scan_all_in_directory_with_scope : blacklisted
func Fn_g_io_scheduler_cancel_all_jobs() {

	C.g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : io_scheduler_push_job : has callback

// UNSUPPORTED : keyfile_settings_backend_new : blacklisted
// UNSUPPORTED : memory_settings_backend_new : blacklisted
func Fn_g_network_monitor_get_default() {

	C.g_network_monitor_get_default()
}

func Fn_g_networking_init() {

	C.g_networking_init()
}

// UNSUPPORTED : null_settings_backend_new : blacklisted
func Fn_g_pollable_source_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	C.g_pollable_source_new(cValue0)
}

func Fn_g_pollable_source_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)
	cValue1 := (*C.GSource)(unsafe.Pointer(param1))
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	C.g_pollable_source_new_full(cValue0, cValue1, cValue2)
}

func Fn_g_pollable_stream_read(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_pollable_stream_write(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_pollable_stream_write_all(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 *uint64, param5 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_proxy_get_default_for_protocol(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_proxy_get_default_for_protocol(cValue0)
}

func Fn_g_proxy_resolver_get_default() {

	C.g_proxy_resolver_get_default()
}

func Fn_g_resolver_error_quark() {

	C.g_resolver_error_quark()
}

func Fn_g_resource_error_quark() {

	C.g_resource_error_quark()
}

func Fn_g_resource_load(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_resource_load(cValue0, cError)
}

func Fn_g_resources_enumerate_children(param0 string, param1 int, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GResourceLookupFlags)(param1)
	cError := (**C.GError)(error)

	C.g_resources_enumerate_children(cValue0, cValue1, cError)
}

func Fn_g_resources_get_info(param0 string, param1 int, param2 *uint64, param3 *uint32, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GResourceLookupFlags)(param1)
	cValue2 := (*C.gsize)(unsafe.Pointer(param2))
	cValue3 := (*C.guint32)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_resources_get_info(cValue0, cValue1, cValue2, cValue3, cError)
}

func Fn_g_resources_lookup_data(param0 string, param1 int, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GResourceLookupFlags)(param1)
	cError := (**C.GError)(error)

	C.g_resources_lookup_data(cValue0, cValue1, cError)
}

func Fn_g_resources_open_stream(param0 string, param1 int, error unsafe.Pointer) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GResourceLookupFlags)(param1)
	cError := (**C.GError)(error)

	C.g_resources_open_stream(cValue0, cValue1, cError)
}

func Fn_g_resources_register(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_register(cValue0)
}

func Fn_g_resources_unregister(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_unregister(cValue0)
}

func Fn_g_settings_schema_source_get_default() {

	C.g_settings_schema_source_get_default()
}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.g_srv_target_list_sort(cValue0)
}

func Fn_g_tls_backend_get_default() {

	C.g_tls_backend_get_default()
}

func Fn_g_tls_client_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GSocketConnectable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_tls_client_connection_new(cValue0, cValue1, cError)
}

func Fn_g_tls_error_quark() {

	C.g_tls_error_quark()
}

func Fn_g_tls_file_database_new(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_tls_file_database_new(cValue0, cError)
}

func Fn_g_tls_server_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GTlsCertificate)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_tls_server_connection_new(cValue0, cValue1, cError)
}

func Fn_g_unix_is_mount_path_system_internal(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_unix_is_mount_path_system_internal(cValue0)
}

func Fn_g_unix_mount_at(param0 string, param1 *uint64) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.guint64)(unsafe.Pointer(param1))

	C.g_unix_mount_at(cValue0, cValue1)
}

func Fn_g_unix_mount_compare(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))
	cValue1 := (*C.GUnixMountEntry)(unsafe.Pointer(param1))

	C.g_unix_mount_compare(cValue0, cValue1)
}

func Fn_g_unix_mount_free(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_free(cValue0)
}

func Fn_g_unix_mount_get_device_path(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_get_device_path(cValue0)
}

func Fn_g_unix_mount_get_fs_type(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_get_fs_type(cValue0)
}

func Fn_g_unix_mount_get_mount_path(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_get_mount_path(cValue0)
}

func Fn_g_unix_mount_guess_can_eject(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_guess_can_eject(cValue0)
}

func Fn_g_unix_mount_guess_icon(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_guess_icon(cValue0)
}

func Fn_g_unix_mount_guess_name(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_guess_name(cValue0)
}

func Fn_g_unix_mount_guess_should_display(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_guess_should_display(cValue0)
}

func Fn_g_unix_mount_guess_symbolic_icon(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_guess_symbolic_icon(cValue0)
}

func Fn_g_unix_mount_is_readonly(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_is_readonly(cValue0)
}

func Fn_g_unix_mount_is_system_internal(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_is_system_internal(cValue0)
}

func Fn_g_unix_mount_points_changed_since(param0 uint64) {
	cValue0 := (C.guint64)(param0)

	C.g_unix_mount_points_changed_since(cValue0)
}

func Fn_g_unix_mount_points_get(param0 *uint64) {
	cValue0 := (*C.guint64)(unsafe.Pointer(param0))

	C.g_unix_mount_points_get(cValue0)
}

func Fn_g_unix_mounts_changed_since(param0 uint64) {
	cValue0 := (C.guint64)(param0)

	C.g_unix_mounts_changed_since(cValue0)
}

func Fn_g_unix_mounts_get(param0 *uint64) {
	cValue0 := (*C.guint64)(unsafe.Pointer(param0))

	C.g_unix_mounts_get(cValue0)
}

func Fn_g_app_launch_context_new() {

	C.g_app_launch_context_new()
}

func Fn_g_app_launch_context_get_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))
	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	C.g_app_launch_context_get_display(cValueInstance, cValue0, cValue1)
}

func Fn_g_app_launch_context_get_environment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))

	C.g_app_launch_context_get_environment(cValueInstance)
}

func Fn_g_app_launch_context_get_startup_notify_id(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))
	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	C.g_app_launch_context_get_startup_notify_id(cValueInstance, cValue0, cValue1)
}

func Fn_g_app_launch_context_launch_failed(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_launch_context_launch_failed(cValueInstance, cValue0)
}

func Fn_g_app_launch_context_setenv(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_app_launch_context_setenv(cValueInstance, cValue0, cValue1)
}

func Fn_g_app_launch_context_unsetenv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_launch_context_unsetenv(cValueInstance, cValue0)
}

func Fn_g_application_new(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GApplicationFlags)(param1)

	C.g_application_new(cValue0, cValue1)
}

func Fn_g_application_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_activate(cValueInstance)
}

func Fn_g_application_get_application_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_application_id(cValueInstance)
}

func Fn_g_application_get_dbus_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_dbus_connection(cValueInstance)
}

func Fn_g_application_get_dbus_object_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_dbus_object_path(cValueInstance)
}

func Fn_g_application_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_flags(cValueInstance)
}

func Fn_g_application_get_inactivity_timeout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_inactivity_timeout(cValueInstance)
}

func Fn_g_application_get_is_registered(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_is_registered(cValueInstance)
}

func Fn_g_application_get_is_remote(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_get_is_remote(cValueInstance)
}

func Fn_g_application_hold(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_hold(cValueInstance)
}

func Fn_g_application_mark_busy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_mark_busy(cValueInstance)
}

func Fn_g_application_open(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 int, param2 string) {
	// has array param
}

func Fn_g_application_quit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_quit(cValueInstance)
}

func Fn_g_application_register(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_application_register(cValueInstance, cValue0, cError)
}

func Fn_g_application_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_release(cValueInstance)
}

func Fn_g_application_run(paramInstance unsafe.Pointer, param0 int, param1 []string) {
	// has array param
}

func Fn_g_application_set_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GActionGroup)(unsafe.Pointer(param0))

	C.g_application_set_action_group(cValueInstance, cValue0)
}

func Fn_g_application_set_application_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_set_application_id(cValueInstance, cValue0)
}

func Fn_g_application_set_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_set_default(cValueInstance)
}

func Fn_g_application_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GApplicationFlags)(param0)

	C.g_application_set_flags(cValueInstance, cValue0)
}

func Fn_g_application_set_inactivity_timeout(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_application_set_inactivity_timeout(cValueInstance, cValue0)
}

func Fn_g_application_unmark_busy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_unmark_busy(cValueInstance)
}

func Fn_g_application_get_default() {

	C.g_application_get_default()
}

func Fn_g_application_id_is_valid(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_id_is_valid(cValue0)
}

func Fn_g_application_command_line_create_file_for_arg(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_command_line_create_file_for_arg(cValueInstance, cValue0)
}

func Fn_g_application_command_line_get_arguments(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	C.g_application_command_line_get_arguments(cValueInstance, cValue0)
}

func Fn_g_application_command_line_get_cwd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_cwd(cValueInstance)
}

func Fn_g_application_command_line_get_environ(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_environ(cValueInstance)
}

func Fn_g_application_command_line_get_exit_status(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_exit_status(cValueInstance)
}

func Fn_g_application_command_line_get_is_remote(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_is_remote(cValueInstance)
}

func Fn_g_application_command_line_get_platform_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_platform_data(cValueInstance)
}

func Fn_g_application_command_line_get_stdin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	C.g_application_command_line_get_stdin(cValueInstance)
}

func Fn_g_application_command_line_getenv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_command_line_getenv(cValueInstance, cValue0)
}

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

func Fn_g_application_command_line_set_exit_status(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.g_application_command_line_set_exit_status(cValueInstance, cValue0)
}

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	C.g_buffered_input_stream_new(cValue0)
}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (C.gsize)(param1)

	C.g_buffered_input_stream_new_sized(cValue0, cValue1)
}

func Fn_g_buffered_input_stream_fill(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gssize)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_buffered_input_stream_fill(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : fill_async : has callback

func Fn_g_buffered_input_stream_fill_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_buffered_input_stream_fill_finish(cValueInstance, cValue0, cError)
}

func Fn_g_buffered_input_stream_get_available(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	C.g_buffered_input_stream_get_available(cValueInstance)
}

func Fn_g_buffered_input_stream_get_buffer_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	C.g_buffered_input_stream_get_buffer_size(cValueInstance)
}

func Fn_g_buffered_input_stream_peek(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 uint64) {
	// has array param
}

func Fn_g_buffered_input_stream_peek_buffer(paramInstance unsafe.Pointer, param0 *uint64) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	C.g_buffered_input_stream_peek_buffer(cValueInstance, cValue0)
}

func Fn_g_buffered_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_buffered_input_stream_read_byte(cValueInstance, cValue0, cError)
}

func Fn_g_buffered_input_stream_set_buffer_size(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gsize)(param0)

	C.g_buffered_input_stream_set_buffer_size(cValueInstance, cValue0)
}

func Fn_g_buffered_output_stream_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	C.g_buffered_output_stream_new(cValue0)
}

func Fn_g_buffered_output_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))
	cValue1 := (C.gsize)(param1)

	C.g_buffered_output_stream_new_sized(cValue0, cValue1)
}

func Fn_g_buffered_output_stream_get_auto_grow(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))

	C.g_buffered_output_stream_get_auto_grow(cValueInstance)
}

func Fn_g_buffered_output_stream_get_buffer_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))

	C.g_buffered_output_stream_get_buffer_size(cValueInstance)
}

func Fn_g_buffered_output_stream_set_auto_grow(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_buffered_output_stream_set_auto_grow(cValueInstance, cValue0)
}

func Fn_g_buffered_output_stream_set_buffer_size(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gsize)(param0)

	C.g_buffered_output_stream_set_buffer_size(cValueInstance, cValue0)
}

func Fn_g_bytes_icon_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	C.g_bytes_icon_new(cValue0)
}

func Fn_g_bytes_icon_get_bytes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GBytesIcon)(unsafe.Pointer(paramInstance))

	C.g_bytes_icon_get_bytes(cValueInstance)
}

func Fn_g_cancellable_new() {

	C.g_cancellable_new()
}

func Fn_g_cancellable_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_cancel(cValueInstance)
}

// UNSUPPORTED : connect : has callback

func Fn_g_cancellable_disconnect(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gulong)(param0)

	C.g_cancellable_disconnect(cValueInstance, cValue0)
}

func Fn_g_cancellable_get_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_get_fd(cValueInstance)
}

func Fn_g_cancellable_is_cancelled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_is_cancelled(cValueInstance)
}

func Fn_g_cancellable_make_pollfd(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GPollFD)(unsafe.Pointer(param0))

	C.g_cancellable_make_pollfd(cValueInstance, cValue0)
}

func Fn_g_cancellable_pop_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_pop_current(cValueInstance)
}

func Fn_g_cancellable_push_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_push_current(cValueInstance)
}

func Fn_g_cancellable_release_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_release_fd(cValueInstance)
}

func Fn_g_cancellable_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_reset(cValueInstance)
}

func Fn_g_cancellable_set_error_if_cancelled(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_cancellable_set_error_if_cancelled(cValueInstance, cError)
}

func Fn_g_cancellable_source_new(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_source_new(cValueInstance)
}

func Fn_g_cancellable_get_current() {

	C.g_cancellable_get_current()
}

func Fn_g_charset_converter_new(param0 string, param1 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cError := (**C.GError)(error)

	C.g_charset_converter_new(cValue0, cValue1, cError)
}

func Fn_g_charset_converter_get_num_fallbacks(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))

	C.g_charset_converter_get_num_fallbacks(cValueInstance)
}

func Fn_g_charset_converter_get_use_fallback(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))

	C.g_charset_converter_get_use_fallback(cValueInstance)
}

func Fn_g_charset_converter_set_use_fallback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_charset_converter_set_use_fallback(cValueInstance, cValue0)
}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	C.g_converter_input_stream_new(cValue0, cValue1)
}

func Fn_g_converter_input_stream_get_converter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GConverterInputStream)(unsafe.Pointer(paramInstance))

	C.g_converter_input_stream_get_converter(cValueInstance)
}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	C.g_converter_output_stream_new(cValue0, cValue1)
}

func Fn_g_converter_output_stream_get_converter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GConverterOutputStream)(unsafe.Pointer(paramInstance))

	C.g_converter_output_stream_get_converter(cValueInstance)
}

func Fn_g_credentials_new() {

	C.g_credentials_new()
}

func Fn_g_credentials_get_native(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GCredentialsType)(param0)

	C.g_credentials_get_native(cValueInstance, cValue0)
}

func Fn_g_credentials_get_unix_pid(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_credentials_get_unix_pid(cValueInstance, cError)
}

func Fn_g_credentials_get_unix_user(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_credentials_get_unix_user(cValueInstance, cError)
}

func Fn_g_credentials_is_same_user(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_credentials_is_same_user(cValueInstance, cValue0, cError)
}

func Fn_g_credentials_set_native(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GCredentialsType)(param0)
	cValue1 := (C.gpointer)(param1)

	C.g_credentials_set_native(cValueInstance, cValue0, cValue1)
}

func Fn_g_credentials_set_unix_user(paramInstance unsafe.Pointer, param0 uint, error unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))
	cValue0 := (C.uid_t)(param0)
	cError := (**C.GError)(error)

	C.g_credentials_set_unix_user(cValueInstance, cValue0, cError)
}

func Fn_g_credentials_to_string(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	C.g_credentials_to_string(cValueInstance)
}

func Fn_g_dbus_action_group_get(param0 unsafe.Pointer, param1 string, param2 string) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_action_group_get(cValue0, cValue1, cValue2)
}

func Fn_g_dbus_auth_observer_new() {

	C.g_dbus_auth_observer_new()
}

func Fn_g_dbus_auth_observer_allow_mechanism(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusAuthObserver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_auth_observer_allow_mechanism(cValueInstance, cValue0)
}

func Fn_g_dbus_auth_observer_authorize_authenticated_peer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusAuthObserver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GCredentials)(unsafe.Pointer(param1))

	C.g_dbus_auth_observer_authorize_authenticated_peer(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_connection_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_new_finish(cValue0, cError)
}

func Fn_g_dbus_connection_new_for_address_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_new_for_address_finish(cValue0, cError)
}

func Fn_g_dbus_connection_new_for_address_sync(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GDBusConnectionFlags)(param1)
	cValue2 := (*C.GDBusAuthObserver)(unsafe.Pointer(param2))
	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_dbus_connection_new_for_address_sync(cValue0, cValue1, cValue2, cValue3, cError)
}

func Fn_g_dbus_connection_new_sync(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (C.GDBusConnectionFlags)(param2)
	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))
	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_dbus_connection_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

// UNSUPPORTED : add_filter : has callback

// UNSUPPORTED : call : has callback

func Fn_g_dbus_connection_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_call_finish(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_call_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))
	cValue5 := (*C.GVariantType)(unsafe.Pointer(param5))
	cValue6 := (C.GDBusCallFlags)(param6)
	cValue7 := (C.gint)(param7)
	cValue8 := (*C.GCancellable)(unsafe.Pointer(param8))
	cError := (**C.GError)(error)

	C.g_dbus_connection_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cError)
}

// UNSUPPORTED : call_with_unix_fd_list : has callback

func Fn_g_dbus_connection_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))
	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_dbus_connection_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_dbus_connection_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, param9 *unsafe.Pointer, param10 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))
	cValue5 := (*C.GVariantType)(unsafe.Pointer(param5))
	cValue6 := (C.GDBusCallFlags)(param6)
	cValue7 := (C.gint)(param7)
	cValue8 := (*C.GUnixFDList)(unsafe.Pointer(param8))
	cValue9 := (**C.GUnixFDList)(unsafe.Pointer(param9))
	cValue10 := (*C.GCancellable)(unsafe.Pointer(param10))
	cError := (**C.GError)(error)

	C.g_dbus_connection_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cError)
}

// UNSUPPORTED : close : has callback

func Fn_g_dbus_connection_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_close_finish(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_close_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_close_sync(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_emit_signal(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (*C.GVariant)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_dbus_connection_emit_signal(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

func Fn_g_dbus_connection_export_action_group(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GActionGroup)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_dbus_connection_export_action_group(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_dbus_connection_export_menu_model(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_dbus_connection_export_menu_model(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : flush : has callback

func Fn_g_dbus_connection_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_flush_finish(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_flush_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_flush_sync(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_get_capabilities(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_capabilities(cValueInstance)
}

func Fn_g_dbus_connection_get_exit_on_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_exit_on_close(cValueInstance)
}

func Fn_g_dbus_connection_get_guid(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_guid(cValueInstance)
}

func Fn_g_dbus_connection_get_last_serial(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_last_serial(cValueInstance)
}

func Fn_g_dbus_connection_get_peer_credentials(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_peer_credentials(cValueInstance)
}

func Fn_g_dbus_connection_get_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_stream(cValueInstance)
}

func Fn_g_dbus_connection_get_unique_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_get_unique_name(cValueInstance)
}

func Fn_g_dbus_connection_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_is_closed(cValueInstance)
}

// UNSUPPORTED : register_object : has callback

// UNSUPPORTED : register_subtree : has callback

func Fn_g_dbus_connection_remove_filter(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_remove_filter(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_send_message(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *uint32, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))
	cValue1 := (C.GDBusSendMessageFlags)(param1)
	cValue2 := (*C.guint32)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_dbus_connection_send_message(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : send_message_with_reply : has callback

func Fn_g_dbus_connection_send_message_with_reply_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_connection_send_message_with_reply_finish(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_connection_send_message_with_reply_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *uint32, param4 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))
	cValue1 := (C.GDBusSendMessageFlags)(param1)
	cValue2 := (C.gint)(param2)
	cValue3 := (*C.guint32)(unsafe.Pointer(param3))
	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_dbus_connection_send_message_with_reply_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

func Fn_g_dbus_connection_set_exit_on_close(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_dbus_connection_set_exit_on_close(cValueInstance, cValue0)
}

// UNSUPPORTED : signal_subscribe : has callback

func Fn_g_dbus_connection_signal_unsubscribe(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_signal_unsubscribe(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_start_message_processing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	C.g_dbus_connection_start_message_processing(cValueInstance)
}

func Fn_g_dbus_connection_unexport_action_group(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unexport_action_group(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_unexport_menu_model(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unexport_menu_model(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_unregister_object(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unregister_object(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_unregister_subtree(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_unregister_subtree(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_address : has callback

func Fn_g_dbus_interface_skeleton_export(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cError := (**C.GError)(error)

	C.g_dbus_interface_skeleton_export(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_dbus_interface_skeleton_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_flush(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_connection(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_connections(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_connections(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_flags(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_info(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_object_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_object_path(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_properties(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_properties(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_vtable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_get_vtable(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_has_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_interface_skeleton_has_connection(cValueInstance, cValue0)
}

func Fn_g_dbus_interface_skeleton_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusInterfaceSkeletonFlags)(param0)

	C.g_dbus_interface_skeleton_set_flags(cValueInstance, cValue0)
}

func Fn_g_dbus_interface_skeleton_unexport(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_unexport(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_unexport_from_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_interface_skeleton_unexport_from_connection(cValueInstance, cValue0)
}

func Fn_g_dbus_menu_model_get(param0 unsafe.Pointer, param1 string, param2 string) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_menu_model_get(cValue0, cValue1, cValue2)
}

func Fn_g_dbus_message_new() {

	C.g_dbus_message_new()
}

func Fn_g_dbus_message_new_from_blob(param0 []uint8, param1 uint64, param2 int, error unsafe.Pointer) {
	// has array param
}

func Fn_g_dbus_message_new_method_call(param0 string, param1 string, param2 string, param3 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	C.g_dbus_message_new_method_call(cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_dbus_message_new_signal(param0 string, param1 string, param2 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_message_new_signal(cValue0, cValue1, cValue2)
}

func Fn_g_dbus_message_copy(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_dbus_message_copy(cValueInstance, cError)
}

func Fn_g_dbus_message_get_arg0(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_arg0(cValueInstance)
}

func Fn_g_dbus_message_get_body(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_body(cValueInstance)
}

func Fn_g_dbus_message_get_byte_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_byte_order(cValueInstance)
}

func Fn_g_dbus_message_get_destination(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_destination(cValueInstance)
}

func Fn_g_dbus_message_get_error_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_error_name(cValueInstance)
}

func Fn_g_dbus_message_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_flags(cValueInstance)
}

func Fn_g_dbus_message_get_header(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageHeaderField)(param0)

	C.g_dbus_message_get_header(cValueInstance, cValue0)
}

func Fn_g_dbus_message_get_header_fields(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_header_fields(cValueInstance)
}

func Fn_g_dbus_message_get_interface(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_interface(cValueInstance)
}

func Fn_g_dbus_message_get_locked(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_locked(cValueInstance)
}

func Fn_g_dbus_message_get_member(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_member(cValueInstance)
}

func Fn_g_dbus_message_get_message_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_message_type(cValueInstance)
}

func Fn_g_dbus_message_get_num_unix_fds(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_num_unix_fds(cValueInstance)
}

func Fn_g_dbus_message_get_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_path(cValueInstance)
}

func Fn_g_dbus_message_get_reply_serial(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_reply_serial(cValueInstance)
}

func Fn_g_dbus_message_get_sender(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_sender(cValueInstance)
}

func Fn_g_dbus_message_get_serial(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_serial(cValueInstance)
}

func Fn_g_dbus_message_get_signature(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_signature(cValueInstance)
}

func Fn_g_dbus_message_get_unix_fd_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_unix_fd_list(cValueInstance)
}

func Fn_g_dbus_message_lock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_lock(cValueInstance)
}

// UNSUPPORTED : new_method_error : has varargs

func Fn_g_dbus_message_new_method_error_literal(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_message_new_method_error_literal(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : new_method_error_valist : has va_list

func Fn_g_dbus_message_new_method_reply(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_new_method_reply(cValueInstance)
}

func Fn_g_dbus_message_print(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_dbus_message_print(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_body(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_dbus_message_set_body(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageByteOrder)(param0)

	C.g_dbus_message_set_byte_order(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_destination(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_destination(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_error_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_error_name(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageFlags)(param0)

	C.g_dbus_message_set_flags(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_header(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageHeaderField)(param0)
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_dbus_message_set_header(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_message_set_interface(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_member(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_member(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_message_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageType)(param0)

	C.g_dbus_message_set_message_type(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_num_unix_fds(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_num_unix_fds(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_path(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_reply_serial(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_reply_serial(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_sender(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_sender(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_serial(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)

	C.g_dbus_message_set_serial(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_signature(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_message_set_signature(cValueInstance, cValue0)
}

func Fn_g_dbus_message_set_unix_fd_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GUnixFDList)(unsafe.Pointer(param0))

	C.g_dbus_message_set_unix_fd_list(cValueInstance, cValue0)
}

func Fn_g_dbus_message_to_blob(paramInstance unsafe.Pointer, param0 *uint64, param1 int, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))
	cValue1 := (C.GDBusCapabilityFlags)(param1)
	cError := (**C.GError)(error)

	C.g_dbus_message_to_blob(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_dbus_message_to_gerror(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_dbus_message_to_gerror(cValueInstance, cError)
}

func Fn_g_dbus_message_bytes_needed(param0 []uint8, param1 uint64, error unsafe.Pointer) {
	// has array param
}

func Fn_g_dbus_method_invocation_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_connection(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_interface_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_interface_name(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_message(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_message(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_method_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_method_info(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_method_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_method_name(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_object_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_object_path(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_parameters(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_parameters(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_property_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_property_info(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_sender(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_sender(cValueInstance)
}

func Fn_g_dbus_method_invocation_get_user_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_invocation_get_user_data(cValueInstance)
}

func Fn_g_dbus_method_invocation_return_dbus_error(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_method_invocation_return_dbus_error(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : return_error : has varargs

func Fn_g_dbus_method_invocation_return_error_literal(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GQuark)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_method_invocation_return_error_literal(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : return_error_valist : has va_list

func Fn_g_dbus_method_invocation_return_gerror(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_return_gerror(cValueInstance, cValue0)
}

func Fn_g_dbus_method_invocation_return_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_return_value(cValueInstance, cValue0)
}

func Fn_g_dbus_method_invocation_return_value_with_unix_fd_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))
	cValue1 := (*C.GUnixFDList)(unsafe.Pointer(param1))

	C.g_dbus_method_invocation_return_value_with_unix_fd_list(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_method_invocation_take_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_dbus_method_invocation_take_error(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_client_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_object_manager_client_new_finish(cValue0, cError)
}

func Fn_g_dbus_object_manager_client_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_object_manager_client_new_for_bus_finish(cValue0, cError)
}

// UNSUPPORTED : new_for_bus_sync : has callback

// UNSUPPORTED : new_sync : has callback

func Fn_g_dbus_object_manager_client_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_manager_client_get_connection(cValueInstance)
}

func Fn_g_dbus_object_manager_client_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_manager_client_get_flags(cValueInstance)
}

func Fn_g_dbus_object_manager_client_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_manager_client_get_name(cValueInstance)
}

func Fn_g_dbus_object_manager_client_get_name_owner(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_manager_client_get_name_owner(cValueInstance)
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_dbus_object_manager_server_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_manager_server_new(cValue0)
}

func Fn_g_dbus_object_manager_server_export(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_export(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_export_uniquely(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_export_uniquely(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_manager_server_get_connection(cValueInstance)
}

func Fn_g_dbus_object_manager_server_is_exported(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_is_exported(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_set_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_set_connection(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_unexport(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_manager_server_unexport(cValueInstance, cValue0)
}

func Fn_g_dbus_object_proxy_new(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_object_proxy_new(cValue0, cValue1)
}

func Fn_g_dbus_object_proxy_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_proxy_get_connection(cValueInstance)
}

func Fn_g_dbus_object_skeleton_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_skeleton_new(cValue0)
}

func Fn_g_dbus_object_skeleton_add_interface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_skeleton_add_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_object_skeleton_flush(cValueInstance)
}

func Fn_g_dbus_object_skeleton_remove_interface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(param0))

	C.g_dbus_object_skeleton_remove_interface(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_remove_interface_by_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_skeleton_remove_interface_by_name(cValueInstance, cValue0)
}

func Fn_g_dbus_object_skeleton_set_object_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusObjectSkeleton)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_object_skeleton_set_object_path(cValueInstance, cValue0)
}

func Fn_g_dbus_proxy_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_new_finish(cValue0, cError)
}

func Fn_g_dbus_proxy_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_new_for_bus_finish(cValue0, cError)
}

func Fn_g_dbus_proxy_new_for_bus_sync(param0 int, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (C.GBusType)(param0)
	cValue1 := (C.GDBusProxyFlags)(param1)
	cValue2 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))
	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))
	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_new_for_bus_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)
}

func Fn_g_dbus_proxy_new_sync(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))
	cValue1 := (C.GDBusProxyFlags)(param1)
	cValue2 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))
	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))
	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)
}

// UNSUPPORTED : call : has callback

func Fn_g_dbus_proxy_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_call_finish(cValueInstance, cValue0, cError)
}

func Fn_g_dbus_proxy_call_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))
	cValue2 := (C.GDBusCallFlags)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

// UNSUPPORTED : call_with_unix_fd_list : has callback

func Fn_g_dbus_proxy_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))
	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_dbus_proxy_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))
	cValue2 := (C.GDBusCallFlags)(param2)
	cValue3 := (C.gint)(param3)
	cValue4 := (*C.GUnixFDList)(unsafe.Pointer(param4))
	cValue5 := (**C.GUnixFDList)(unsafe.Pointer(param5))
	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))
	cError := (**C.GError)(error)

	C.g_dbus_proxy_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)
}

func Fn_g_dbus_proxy_get_cached_property(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_dbus_proxy_get_cached_property(cValueInstance, cValue0)
}

func Fn_g_dbus_proxy_get_cached_property_names(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_cached_property_names(cValueInstance)
}

func Fn_g_dbus_proxy_get_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_connection(cValueInstance)
}

func Fn_g_dbus_proxy_get_default_timeout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_default_timeout(cValueInstance)
}

func Fn_g_dbus_proxy_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_flags(cValueInstance)
}

func Fn_g_dbus_proxy_get_interface_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_interface_info(cValueInstance)
}

func Fn_g_dbus_proxy_get_interface_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_interface_name(cValueInstance)
}

func Fn_g_dbus_proxy_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_name(cValueInstance)
}

func Fn_g_dbus_proxy_get_name_owner(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_name_owner(cValueInstance)
}

func Fn_g_dbus_proxy_get_object_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	C.g_dbus_proxy_get_object_path(cValueInstance)
}

func Fn_g_dbus_proxy_set_cached_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_dbus_proxy_set_cached_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_proxy_set_default_timeout(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_dbus_proxy_set_default_timeout(cValueInstance, cValue0)
}

func Fn_g_dbus_proxy_set_interface_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param0))

	C.g_dbus_proxy_set_interface_info(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_dbus_server_new_sync(param0 string, param1 int, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GDBusServerFlags)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))
	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_dbus_server_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

func Fn_g_dbus_server_get_client_address(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_get_client_address(cValueInstance)
}

func Fn_g_dbus_server_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_get_flags(cValueInstance)
}

func Fn_g_dbus_server_get_guid(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_get_guid(cValueInstance)
}

func Fn_g_dbus_server_is_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_is_active(cValueInstance)
}

func Fn_g_dbus_server_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_start(cValueInstance)
}

func Fn_g_dbus_server_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_stop(cValueInstance)
}

func Fn_g_data_input_stream_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	C.g_data_input_stream_new(cValue0)
}

func Fn_g_data_input_stream_get_byte_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	C.g_data_input_stream_get_byte_order(cValueInstance)
}

func Fn_g_data_input_stream_get_newline_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	C.g_data_input_stream_get_newline_type(cValueInstance)
}

func Fn_g_data_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_byte(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_int16(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_int16(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_int32(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_int32(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_int64(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_int64(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_line(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_line(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : read_line_async : has callback

func Fn_g_data_input_stream_read_line_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_line_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_input_stream_read_line_finish_utf8(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_line_finish_utf8(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_input_stream_read_line_utf8(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_line_utf8(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_input_stream_read_uint16(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_uint16(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_uint32(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_uint32(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_uint64(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_uint64(cValueInstance, cValue0, cError)
}

func Fn_g_data_input_stream_read_until(paramInstance unsafe.Pointer, param0 string, param1 *uint64, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_until(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : read_until_async : has callback

func Fn_g_data_input_stream_read_until_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_until_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_input_stream_read_upto(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gssize)(param1)
	cValue2 := (*C.gsize)(unsafe.Pointer(param2))
	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_upto(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
}

// UNSUPPORTED : read_upto_async : has callback

func Fn_g_data_input_stream_read_upto_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_input_stream_read_upto_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_input_stream_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDataStreamByteOrder)(param0)

	C.g_data_input_stream_set_byte_order(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_set_newline_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDataStreamNewlineType)(param0)

	C.g_data_input_stream_set_newline_type(cValueInstance, cValue0)
}

func Fn_g_data_output_stream_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	C.g_data_output_stream_new(cValue0)
}

func Fn_g_data_output_stream_get_byte_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	C.g_data_output_stream_get_byte_order(cValueInstance)
}

func Fn_g_data_output_stream_put_byte(paramInstance unsafe.Pointer, param0 uint8, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guchar)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_byte(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_int16(paramInstance unsafe.Pointer, param0 int16, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint16)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_int16(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_int32(paramInstance unsafe.Pointer, param0 int32, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint32)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_int32(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_int64(paramInstance unsafe.Pointer, param0 int64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint64)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_int64(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_string(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_string(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_uint16(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint16)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_uint16(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_uint32(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_uint32(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_put_uint64(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint64)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_data_output_stream_put_uint64(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_data_output_stream_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDataStreamByteOrder)(param0)

	C.g_data_output_stream_set_byte_order(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_new(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_new(cValue0)
}

func Fn_g_desktop_app_info_new_from_filename(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_new_from_filename(cValue0)
}

func Fn_g_desktop_app_info_new_from_keyfile(param0 unsafe.Pointer) {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	C.g_desktop_app_info_new_from_keyfile(cValue0)
}

func Fn_g_desktop_app_info_get_action_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_get_action_name(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_get_boolean(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_get_boolean(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_get_categories(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_categories(cValueInstance)
}

func Fn_g_desktop_app_info_get_filename(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_filename(cValueInstance)
}

func Fn_g_desktop_app_info_get_generic_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_generic_name(cValueInstance)
}

func Fn_g_desktop_app_info_get_is_hidden(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_is_hidden(cValueInstance)
}

func Fn_g_desktop_app_info_get_keywords(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_keywords(cValueInstance)
}

func Fn_g_desktop_app_info_get_nodisplay(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_nodisplay(cValueInstance)
}

func Fn_g_desktop_app_info_get_show_in(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_get_show_in(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_get_startup_wm_class(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_startup_wm_class(cValueInstance)
}

func Fn_g_desktop_app_info_get_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_get_string(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_has_key(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_has_key(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_launch_action(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))

	C.g_desktop_app_info_launch_action(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : launch_uris_as_manager : has callback

// UNSUPPORTED : launch_uris_as_manager_with_fds : has callback

func Fn_g_desktop_app_info_list_actions(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_list_actions(cValueInstance)
}

func Fn_g_desktop_app_info_search(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_search(cValue0)
}

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_set_desktop_env(cValue0)
}

func Fn_g_emblem_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_emblem_new(cValue0)
}

func Fn_g_emblem_new_with_origin(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))
	cValue1 := (C.GEmblemOrigin)(param1)

	C.g_emblem_new_with_origin(cValue0, cValue1)
}

func Fn_g_emblem_get_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblem)(unsafe.Pointer(paramInstance))

	C.g_emblem_get_icon(cValueInstance)
}

func Fn_g_emblem_get_origin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblem)(unsafe.Pointer(paramInstance))

	C.g_emblem_get_origin(cValueInstance)
}

func Fn_g_emblemed_icon_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))
	cValue1 := (*C.GEmblem)(unsafe.Pointer(param1))

	C.g_emblemed_icon_new(cValue0, cValue1)
}

func Fn_g_emblemed_icon_add_emblem(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GEmblem)(unsafe.Pointer(param0))

	C.g_emblemed_icon_add_emblem(cValueInstance, cValue0)
}

func Fn_g_emblemed_icon_clear_emblems(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	C.g_emblemed_icon_clear_emblems(cValueInstance)
}

func Fn_g_emblemed_icon_get_emblems(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	C.g_emblemed_icon_get_emblems(cValueInstance)
}

func Fn_g_emblemed_icon_get_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	C.g_emblemed_icon_get_icon(cValueInstance)
}

func Fn_g_file_enumerator_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_enumerator_close(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_file_enumerator_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_enumerator_close_finish(cValueInstance, cValue0, cError)
}

func Fn_g_file_enumerator_get_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	C.g_file_enumerator_get_child(cValueInstance, cValue0)
}

func Fn_g_file_enumerator_get_container(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	C.g_file_enumerator_get_container(cValueInstance)
}

func Fn_g_file_enumerator_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	C.g_file_enumerator_has_pending(cValueInstance)
}

func Fn_g_file_enumerator_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	C.g_file_enumerator_is_closed(cValueInstance)
}

func Fn_g_file_enumerator_next_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_enumerator_next_file(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : next_files_async : has callback

func Fn_g_file_enumerator_next_files_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_enumerator_next_files_finish(cValueInstance, cValue0, cError)
}

func Fn_g_file_enumerator_set_pending(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_file_enumerator_set_pending(cValueInstance, cValue0)
}

func Fn_g_file_io_stream_get_etag(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	C.g_file_io_stream_get_etag(cValueInstance)
}

func Fn_g_file_io_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_file_io_stream_query_info(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_io_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_io_stream_query_info_finish(cValueInstance, cValue0, cError)
}

func Fn_g_file_icon_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.g_file_icon_new(cValue0)
}

func Fn_g_file_icon_get_file(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileIcon)(unsafe.Pointer(paramInstance))

	C.g_file_icon_get_file(cValueInstance)
}

func Fn_g_file_info_new() {

	C.g_file_info_new()
}

func Fn_g_file_info_clear_status(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_clear_status(cValueInstance)
}

func Fn_g_file_info_copy_into(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	C.g_file_info_copy_into(cValueInstance, cValue0)
}

func Fn_g_file_info_dup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_dup(cValueInstance)
}

func Fn_g_file_info_get_attribute_as_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_as_string(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_boolean(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_boolean(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_byte_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_byte_string(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_data(paramInstance unsafe.Pointer, param0 string, param1 *int, param2 *unsafe.Pointer, param3 *int) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GFileAttributeType)(unsafe.Pointer(param1))
	cValue2 := (*C.gpointer)(unsafe.Pointer(param2))
	cValue3 := (*C.GFileAttributeStatus)(unsafe.Pointer(param3))

	C.g_file_info_get_attribute_data(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_file_info_get_attribute_int32(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_int32(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_int64(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_int64(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_object(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_object(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_status(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_status(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_string(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_stringv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_stringv(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_type(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_uint32(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_uint32(cValueInstance, cValue0)
}

func Fn_g_file_info_get_attribute_uint64(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_get_attribute_uint64(cValueInstance, cValue0)
}

func Fn_g_file_info_get_content_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_content_type(cValueInstance)
}

func Fn_g_file_info_get_deletion_date(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_deletion_date(cValueInstance)
}

func Fn_g_file_info_get_display_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_display_name(cValueInstance)
}

func Fn_g_file_info_get_edit_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_edit_name(cValueInstance)
}

func Fn_g_file_info_get_etag(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_etag(cValueInstance)
}

func Fn_g_file_info_get_file_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_file_type(cValueInstance)
}

func Fn_g_file_info_get_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_icon(cValueInstance)
}

func Fn_g_file_info_get_is_backup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_is_backup(cValueInstance)
}

func Fn_g_file_info_get_is_hidden(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_is_hidden(cValueInstance)
}

func Fn_g_file_info_get_is_symlink(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_is_symlink(cValueInstance)
}

func Fn_g_file_info_get_modification_time(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	C.g_file_info_get_modification_time(cValueInstance, cValue0)
}

func Fn_g_file_info_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_name(cValueInstance)
}

func Fn_g_file_info_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_size(cValueInstance)
}

func Fn_g_file_info_get_sort_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_sort_order(cValueInstance)
}

func Fn_g_file_info_get_symbolic_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_symbolic_icon(cValueInstance)
}

func Fn_g_file_info_get_symlink_target(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_symlink_target(cValueInstance)
}

func Fn_g_file_info_has_attribute(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_has_attribute(cValueInstance, cValue0)
}

func Fn_g_file_info_has_namespace(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_has_namespace(cValueInstance, cValue0)
}

func Fn_g_file_info_list_attributes(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_list_attributes(cValueInstance, cValue0)
}

func Fn_g_file_info_remove_attribute(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_remove_attribute(cValueInstance, cValue0)
}

func Fn_g_file_info_set_attribute(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GFileAttributeType)(param1)
	cValue2 := (C.gpointer)(param2)

	C.g_file_info_set_attribute(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_file_info_set_attribute_boolean(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := toCBool(param1)

	C.g_file_info_set_attribute_boolean(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_byte_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_file_info_set_attribute_byte_string(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_int32(paramInstance unsafe.Pointer, param0 string, param1 int32) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gint32)(param1)

	C.g_file_info_set_attribute_int32(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_int64(paramInstance unsafe.Pointer, param0 string, param1 int64) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gint64)(param1)

	C.g_file_info_set_attribute_int64(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_mask(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFileAttributeMatcher)(unsafe.Pointer(param0))

	C.g_file_info_set_attribute_mask(cValueInstance, cValue0)
}

func Fn_g_file_info_set_attribute_object(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	C.g_file_info_set_attribute_object(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_status(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GFileAttributeStatus)(param1)

	C.g_file_info_set_attribute_status(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_file_info_set_attribute_string(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_stringv(paramInstance unsafe.Pointer, param0 string, param1 []string) {
	// has array param
}

func Fn_g_file_info_set_attribute_uint32(paramInstance unsafe.Pointer, param0 string, param1 uint32) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint32)(param1)

	C.g_file_info_set_attribute_uint32(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_attribute_uint64(paramInstance unsafe.Pointer, param0 string, param1 uint64) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint64)(param1)

	C.g_file_info_set_attribute_uint64(cValueInstance, cValue0, cValue1)
}

func Fn_g_file_info_set_content_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_set_content_type(cValueInstance, cValue0)
}

func Fn_g_file_info_set_display_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_set_display_name(cValueInstance, cValue0)
}

func Fn_g_file_info_set_edit_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_set_edit_name(cValueInstance, cValue0)
}

func Fn_g_file_info_set_file_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GFileType)(param0)

	C.g_file_info_set_file_type(cValueInstance, cValue0)
}

func Fn_g_file_info_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_file_info_set_icon(cValueInstance, cValue0)
}

func Fn_g_file_info_set_is_hidden(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_file_info_set_is_hidden(cValueInstance, cValue0)
}

func Fn_g_file_info_set_is_symlink(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_file_info_set_is_symlink(cValueInstance, cValue0)
}

func Fn_g_file_info_set_modification_time(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	C.g_file_info_set_modification_time(cValueInstance, cValue0)
}

func Fn_g_file_info_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_set_name(cValueInstance, cValue0)
}

func Fn_g_file_info_set_size(paramInstance unsafe.Pointer, param0 int64) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (C.goffset)(param0)

	C.g_file_info_set_size(cValueInstance, cValue0)
}

func Fn_g_file_info_set_sort_order(paramInstance unsafe.Pointer, param0 int32) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint32)(param0)

	C.g_file_info_set_sort_order(cValueInstance, cValue0)
}

func Fn_g_file_info_set_symbolic_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_file_info_set_symbolic_icon(cValueInstance, cValue0)
}

func Fn_g_file_info_set_symlink_target(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_file_info_set_symlink_target(cValueInstance, cValue0)
}

func Fn_g_file_info_unset_attribute_mask(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_unset_attribute_mask(cValueInstance)
}

func Fn_g_file_input_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_file_input_stream_query_info(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_input_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_input_stream_query_info_finish(cValueInstance, cValue0, cError)
}

func Fn_g_file_monitor_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	C.g_file_monitor_cancel(cValueInstance)
}

func Fn_g_file_monitor_emit_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFile)(unsafe.Pointer(param0))
	cValue1 := (*C.GFile)(unsafe.Pointer(param1))
	cValue2 := (C.GFileMonitorEvent)(param2)

	C.g_file_monitor_emit_event(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_file_monitor_is_cancelled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	C.g_file_monitor_is_cancelled(cValueInstance)
}

func Fn_g_file_monitor_set_rate_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_file_monitor_set_rate_limit(cValueInstance, cValue0)
}

func Fn_g_file_output_stream_get_etag(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))

	C.g_file_output_stream_get_etag(cValueInstance)
}

func Fn_g_file_output_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_file_output_stream_query_info(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_output_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_file_output_stream_query_info_finish(cValueInstance, cValue0, cError)
}

func Fn_g_filename_completer_new() {

	C.g_filename_completer_new()
}

func Fn_g_filename_completer_get_completion_suffix(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFilenameCompleter)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_filename_completer_get_completion_suffix(cValueInstance, cValue0)
}

func Fn_g_filename_completer_get_completions(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GFilenameCompleter)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_filename_completer_get_completions(cValueInstance, cValue0)
}

func Fn_g_filename_completer_set_dirs_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilenameCompleter)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_filename_completer_set_dirs_only(cValueInstance, cValue0)
}

func Fn_g_filter_input_stream_get_base_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))

	C.g_filter_input_stream_get_base_stream(cValueInstance)
}

func Fn_g_filter_input_stream_get_close_base_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))

	C.g_filter_input_stream_get_close_base_stream(cValueInstance)
}

func Fn_g_filter_input_stream_set_close_base_stream(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_filter_input_stream_set_close_base_stream(cValueInstance, cValue0)
}

func Fn_g_filter_output_stream_get_base_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))

	C.g_filter_output_stream_get_base_stream(cValueInstance)
}

func Fn_g_filter_output_stream_get_close_base_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))

	C.g_filter_output_stream_get_close_base_stream(cValueInstance)
}

func Fn_g_filter_output_stream_set_close_base_stream(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_filter_output_stream_set_close_base_stream(cValueInstance, cValue0)
}

// UNSUPPORTED : new : blacklisted
// UNSUPPORTED : load : blacklisted
// UNSUPPORTED : unload : blacklisted
// UNSUPPORTED : query : blacklisted
func Fn_g_io_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_clear_pending(cValueInstance)
}

func Fn_g_io_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_io_stream_close(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_io_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_io_stream_close_finish(cValueInstance, cValue0, cError)
}

func Fn_g_io_stream_get_input_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_get_input_stream(cValueInstance)
}

func Fn_g_io_stream_get_output_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_get_output_stream(cValueInstance)
}

func Fn_g_io_stream_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_has_pending(cValueInstance)
}

func Fn_g_io_stream_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_is_closed(cValueInstance)
}

func Fn_g_io_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_io_stream_set_pending(cValueInstance, cError)
}

// UNSUPPORTED : splice_async : has callback

func Fn_g_io_stream_splice_finish(param0 unsafe.Pointer, error unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_io_stream_splice_finish(cValue0, cError)
}

func Fn_g_inet_address_new_any(param0 int) {
	cValue0 := (C.GSocketFamily)(param0)

	C.g_inet_address_new_any(cValue0)
}

func Fn_g_inet_address_new_from_bytes(param0 []uint8, param1 int) {
	// has array param
}

func Fn_g_inet_address_new_from_string(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_inet_address_new_from_string(cValue0)
}

func Fn_g_inet_address_new_loopback(param0 int) {
	cValue0 := (C.GSocketFamily)(param0)

	C.g_inet_address_new_loopback(cValue0)
}

func Fn_g_inet_address_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	C.g_inet_address_equal(cValueInstance, cValue0)
}

func Fn_g_inet_address_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_family(cValueInstance)
}

func Fn_g_inet_address_get_is_any(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_any(cValueInstance)
}

func Fn_g_inet_address_get_is_link_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_link_local(cValueInstance)
}

func Fn_g_inet_address_get_is_loopback(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_loopback(cValueInstance)
}

func Fn_g_inet_address_get_is_mc_global(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_mc_global(cValueInstance)
}

func Fn_g_inet_address_get_is_mc_link_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_mc_link_local(cValueInstance)
}

func Fn_g_inet_address_get_is_mc_node_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_mc_node_local(cValueInstance)
}

func Fn_g_inet_address_get_is_mc_org_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_mc_org_local(cValueInstance)
}

func Fn_g_inet_address_get_is_mc_site_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_mc_site_local(cValueInstance)
}

func Fn_g_inet_address_get_is_multicast(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_multicast(cValueInstance)
}

func Fn_g_inet_address_get_is_site_local(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_is_site_local(cValueInstance)
}

func Fn_g_inet_address_get_native_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_get_native_size(cValueInstance)
}

func Fn_g_inet_address_to_bytes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_to_bytes(cValueInstance)
}

func Fn_g_inet_address_to_string(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_address_to_string(cValueInstance)
}

func Fn_g_inet_address_mask_new(param0 unsafe.Pointer, param1 uint, error unsafe.Pointer) {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := (C.guint)(param1)
	cError := (**C.GError)(error)

	C.g_inet_address_mask_new(cValue0, cValue1, cError)
}

func Fn_g_inet_address_mask_new_from_string(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_inet_address_mask_new_from_string(cValue0, cError)
}

func Fn_g_inet_address_mask_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddressMask)(unsafe.Pointer(param0))

	C.g_inet_address_mask_equal(cValueInstance, cValue0)
}

func Fn_g_inet_address_mask_get_address(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	C.g_inet_address_mask_get_address(cValueInstance)
}

func Fn_g_inet_address_mask_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	C.g_inet_address_mask_get_family(cValueInstance)
}

func Fn_g_inet_address_mask_get_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	C.g_inet_address_mask_get_length(cValueInstance)
}

func Fn_g_inet_address_mask_matches(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	C.g_inet_address_mask_matches(cValueInstance, cValue0)
}

func Fn_g_inet_address_mask_to_string(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	C.g_inet_address_mask_to_string(cValueInstance)
}

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := (C.guint16)(param1)

	C.g_inet_socket_address_new(cValue0, cValue1)
}

func Fn_g_inet_socket_address_get_address(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_socket_address_get_address(cValueInstance)
}

func Fn_g_inet_socket_address_get_flowinfo(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_socket_address_get_flowinfo(cValueInstance)
}

func Fn_g_inet_socket_address_get_port(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_socket_address_get_port(cValueInstance)
}

func Fn_g_inet_socket_address_get_scope_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_inet_socket_address_get_scope_id(cValueInstance)
}

func Fn_g_input_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_clear_pending(cValueInstance)
}

func Fn_g_input_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_input_stream_close(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_input_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_input_stream_close_finish(cValueInstance, cValue0, cError)
}

func Fn_g_input_stream_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_has_pending(cValueInstance)
}

func Fn_g_input_stream_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_is_closed(cValueInstance)
}

func Fn_g_input_stream_read(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_input_stream_read_all(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : read_all_async : has callback

// UNSUPPORTED : read_async : has callback

func Fn_g_input_stream_read_bytes(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gsize)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_input_stream_read_bytes(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : read_bytes_async : has callback

func Fn_g_input_stream_read_bytes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_input_stream_read_bytes_finish(cValueInstance, cValue0, cError)
}

func Fn_g_input_stream_read_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_input_stream_read_finish(cValueInstance, cValue0, cError)
}

func Fn_g_input_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_input_stream_set_pending(cValueInstance, cError)
}

func Fn_g_input_stream_skip(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gsize)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_input_stream_skip(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : skip_async : has callback

func Fn_g_input_stream_skip_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_input_stream_skip_finish(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : insert_sorted : has callback

// UNSUPPORTED : sort : has callback

func Fn_g_memory_input_stream_new() {

	C.g_memory_input_stream_new()
}

func Fn_g_memory_input_stream_new_from_bytes(param0 unsafe.Pointer) {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	C.g_memory_input_stream_new_from_bytes(cValue0)
}

// UNSUPPORTED : new_from_data : has callback

func Fn_g_memory_input_stream_add_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMemoryInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	C.g_memory_input_stream_add_bytes(cValueInstance, cValue0)
}

// UNSUPPORTED : add_data : has callback

// UNSUPPORTED : new : has callback

func Fn_g_memory_output_stream_new_resizable() {

	C.g_memory_output_stream_new_resizable()
}

func Fn_g_memory_output_stream_get_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	C.g_memory_output_stream_get_data(cValueInstance)
}

func Fn_g_memory_output_stream_get_data_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	C.g_memory_output_stream_get_data_size(cValueInstance)
}

func Fn_g_memory_output_stream_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	C.g_memory_output_stream_get_size(cValueInstance)
}

func Fn_g_memory_output_stream_steal_as_bytes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	C.g_memory_output_stream_steal_as_bytes(cValueInstance)
}

func Fn_g_memory_output_stream_steal_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	C.g_memory_output_stream_steal_data(cValueInstance)
}

func Fn_g_menu_new() {

	C.g_menu_new()
}

func Fn_g_menu_append(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_append(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_append_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GMenuItem)(unsafe.Pointer(param0))

	C.g_menu_append_item(cValueInstance, cValue0)
}

func Fn_g_menu_append_section(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_append_section(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_append_submenu(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_append_submenu(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_freeze(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	C.g_menu_freeze(cValueInstance)
}

func Fn_g_menu_insert(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_menu_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_insert_item(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.GMenuItem)(unsafe.Pointer(param1))

	C.g_menu_insert_item(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_insert_section(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.GMenuModel)(unsafe.Pointer(param2))

	C.g_menu_insert_section(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_insert_submenu(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.GMenuModel)(unsafe.Pointer(param2))

	C.g_menu_insert_submenu(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_prepend(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_prepend(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_prepend_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GMenuItem)(unsafe.Pointer(param0))

	C.g_menu_prepend_item(cValueInstance, cValue0)
}

func Fn_g_menu_prepend_section(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_prepend_section(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_prepend_submenu(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_prepend_submenu(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_remove(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_menu_remove(cValueInstance, cValue0)
}

func Fn_g_menu_remove_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenu)(unsafe.Pointer(paramInstance))

	C.g_menu_remove_all(cValueInstance)
}

func Fn_g_menu_attribute_iter_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	C.g_menu_attribute_iter_get_name(cValueInstance)
}

func Fn_g_menu_attribute_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))
	var cValue0String *C.gchar
	cValue0 := &cValue0String
	cValue1 := (**C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_attribute_iter_get_next(cValueInstance, cValue0, cValue1)
	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_g_menu_attribute_iter_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	C.g_menu_attribute_iter_get_value(cValueInstance)
}

func Fn_g_menu_attribute_iter_next(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	C.g_menu_attribute_iter_next(cValueInstance)
}

func Fn_g_menu_item_new(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_item_new(cValue0, cValue1)
}

func Fn_g_menu_item_new_from_model(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))
	cValue1 := (C.gint)(param1)

	C.g_menu_item_new_from_model(cValue0, cValue1)
}

func Fn_g_menu_item_new_section(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_item_new_section(cValue0, cValue1)
}

func Fn_g_menu_item_new_submenu(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_item_new_submenu(cValue0, cValue1)
}

// UNSUPPORTED : get_attribute : has varargs

func Fn_g_menu_item_get_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	C.g_menu_item_get_attribute_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_get_link(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_menu_item_get_link(cValueInstance, cValue0)
}

// UNSUPPORTED : set_action_and_target : has varargs

func Fn_g_menu_item_set_action_and_target_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_item_set_action_and_target_value(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_attribute : has varargs

func Fn_g_menu_item_set_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_item_set_attribute_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_detailed_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_menu_item_set_detailed_action(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_menu_item_set_icon(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_menu_item_set_label(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_link(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_item_set_link(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_item_set_section(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.g_menu_item_set_section(cValueInstance, cValue0)
}

func Fn_g_menu_item_set_submenu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	C.g_menu_item_set_submenu(cValueInstance, cValue0)
}

func Fn_g_menu_link_iter_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	C.g_menu_link_iter_get_name(cValueInstance)
}

func Fn_g_menu_link_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))
	var cValue0String *C.gchar
	cValue0 := &cValue0String
	cValue1 := (**C.GMenuModel)(unsafe.Pointer(param1))

	C.g_menu_link_iter_get_next(cValueInstance, cValue0, cValue1)
	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_g_menu_link_iter_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	C.g_menu_link_iter_get_value(cValueInstance)
}

func Fn_g_menu_link_iter_next(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	C.g_menu_link_iter_next(cValueInstance)
}

// UNSUPPORTED : get_item_attribute : has varargs

func Fn_g_menu_model_get_item_attribute_value(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.GVariantType)(unsafe.Pointer(param2))

	C.g_menu_model_get_item_attribute_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_model_get_item_link(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_menu_model_get_item_link(cValueInstance, cValue0, cValue1)
}

func Fn_g_menu_model_get_n_items(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	C.g_menu_model_get_n_items(cValueInstance)
}

func Fn_g_menu_model_is_mutable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	C.g_menu_model_is_mutable(cValueInstance)
}

func Fn_g_menu_model_items_changed(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)

	C.g_menu_model_items_changed(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_model_iterate_item_attributes(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_menu_model_iterate_item_attributes(cValueInstance, cValue0)
}

func Fn_g_menu_model_iterate_item_links(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_menu_model_iterate_item_links(cValueInstance, cValue0)
}

func Fn_g_mount_operation_new() {

	C.g_mount_operation_new()
}

func Fn_g_mount_operation_get_anonymous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_anonymous(cValueInstance)
}

func Fn_g_mount_operation_get_choice(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_choice(cValueInstance)
}

func Fn_g_mount_operation_get_domain(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_domain(cValueInstance)
}

func Fn_g_mount_operation_get_password(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_password(cValueInstance)
}

func Fn_g_mount_operation_get_password_save(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_password_save(cValueInstance)
}

func Fn_g_mount_operation_get_username(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	C.g_mount_operation_get_username(cValueInstance)
}

func Fn_g_mount_operation_reply(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GMountOperationResult)(param0)

	C.g_mount_operation_reply(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_anonymous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_mount_operation_set_anonymous(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_choice(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.g_mount_operation_set_choice(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_mount_operation_set_domain(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_password(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_mount_operation_set_password(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_password_save(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GPasswordSave)(param0)

	C.g_mount_operation_set_password_save(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_username(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_mount_operation_set_username(cValueInstance, cValue0)
}

func Fn_g_network_address_new(param0 string, param1 uint16) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint16)(param1)

	C.g_network_address_new(cValue0, cValue1)
}

func Fn_g_network_address_get_hostname(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	C.g_network_address_get_hostname(cValueInstance)
}

func Fn_g_network_address_get_port(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	C.g_network_address_get_port(cValueInstance)
}

func Fn_g_network_address_get_scheme(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	C.g_network_address_get_scheme(cValueInstance)
}

func Fn_g_network_address_parse(param0 string, param1 uint16, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint16)(param1)
	cError := (**C.GError)(error)

	C.g_network_address_parse(cValue0, cValue1, cError)
}

func Fn_g_network_address_parse_uri(param0 string, param1 uint16, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint16)(param1)
	cError := (**C.GError)(error)

	C.g_network_address_parse_uri(cValue0, cValue1, cError)
}

func Fn_g_network_service_new(param0 string, param1 string, param2 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_network_service_new(cValue0, cValue1, cValue2)
}

func Fn_g_network_service_get_domain(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	C.g_network_service_get_domain(cValueInstance)
}

func Fn_g_network_service_get_protocol(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	C.g_network_service_get_protocol(cValueInstance)
}

func Fn_g_network_service_get_scheme(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	C.g_network_service_get_scheme(cValueInstance)
}

func Fn_g_network_service_get_service(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	C.g_network_service_get_service(cValueInstance)
}

func Fn_g_network_service_set_scheme(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_network_service_set_scheme(cValueInstance, cValue0)
}

// UNSUPPORTED : add_button_with_target : has varargs

// UNSUPPORTED : set_default_action_and_target : has varargs

func Fn_g_notification_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GNotificationPriority)(param0)

	C.g_notification_set_priority(cValueInstance, cValue0)
}

func Fn_g_output_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_clear_pending(cValueInstance)
}

func Fn_g_output_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_close(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_output_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_close_finish(cValueInstance, cValue0, cError)
}

func Fn_g_output_stream_flush(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_flush(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : flush_async : has callback

func Fn_g_output_stream_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_flush_finish(cValueInstance, cValue0, cError)
}

func Fn_g_output_stream_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_has_pending(cValueInstance)
}

func Fn_g_output_stream_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_is_closed(cValueInstance)
}

func Fn_g_output_stream_is_closing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_is_closing(cValueInstance)
}

// UNSUPPORTED : printf : has varargs

func Fn_g_output_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_output_stream_set_pending(cValueInstance, cError)
}

func Fn_g_output_stream_splice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (C.GOutputStreamSpliceFlags)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_output_stream_splice(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : splice_async : has callback

func Fn_g_output_stream_splice_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_splice_finish(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : vprintf : has va_list

func Fn_g_output_stream_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_output_stream_write_all(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : write_all_async : has callback

// UNSUPPORTED : write_async : has callback

func Fn_g_output_stream_write_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_output_stream_write_bytes(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : write_bytes_async : has callback

func Fn_g_output_stream_write_bytes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_write_bytes_finish(cValueInstance, cValue0, cError)
}

func Fn_g_output_stream_write_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_output_stream_write_finish(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : writev_all_async : has callback

// UNSUPPORTED : writev_async : has callback

func Fn_g_permission_acquire(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_permission_acquire(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : acquire_async : has callback

func Fn_g_permission_acquire_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_permission_acquire_finish(cValueInstance, cValue0, cError)
}

func Fn_g_permission_get_allowed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	C.g_permission_get_allowed(cValueInstance)
}

func Fn_g_permission_get_can_acquire(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	C.g_permission_get_can_acquire(cValueInstance)
}

func Fn_g_permission_get_can_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	C.g_permission_get_can_release(cValueInstance)
}

func Fn_g_permission_impl_update(paramInstance unsafe.Pointer, param0 bool, param1 bool, param2 bool) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)
	cValue1 := toCBool(param1)
	cValue2 := toCBool(param2)

	C.g_permission_impl_update(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_permission_release(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_permission_release(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : release_async : has callback

func Fn_g_permission_release_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_permission_release_finish(cValueInstance, cValue0, cError)
}

func Fn_g_property_action_new(param0 string, param1 unsafe.Pointer, param2 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gpointer)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_property_action_new(cValue0, cValue1, cValue2)
}

func Fn_g_proxy_address_new(param0 unsafe.Pointer, param1 uint16, param2 string, param3 string, param4 uint16, param5 string, param6 string) {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := (C.guint16)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))
	cValue4 := (C.guint16)(param4)
	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))
	cValue6 := (*C.gchar)(C.CString(param6))
	defer C.free(unsafe.Pointer(cValue6))

	C.g_proxy_address_new(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_g_proxy_address_get_destination_hostname(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_destination_hostname(cValueInstance)
}

func Fn_g_proxy_address_get_destination_port(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_destination_port(cValueInstance)
}

func Fn_g_proxy_address_get_destination_protocol(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_destination_protocol(cValueInstance)
}

func Fn_g_proxy_address_get_password(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_password(cValueInstance)
}

func Fn_g_proxy_address_get_protocol(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_protocol(cValueInstance)
}

func Fn_g_proxy_address_get_uri(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_uri(cValueInstance)
}

func Fn_g_proxy_address_get_username(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	C.g_proxy_address_get_username(cValueInstance)
}

func Fn_g_resolver_lookup_by_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_by_address(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : lookup_by_address_async : has callback

func Fn_g_resolver_lookup_by_address_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_by_address_finish(cValueInstance, cValue0, cError)
}

func Fn_g_resolver_lookup_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_by_name(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : lookup_by_name_async : has callback

func Fn_g_resolver_lookup_by_name_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_by_name_finish(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : lookup_by_name_with_flags_async : has callback

func Fn_g_resolver_lookup_records(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.GResolverRecordType)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_records(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : lookup_records_async : has callback

func Fn_g_resolver_lookup_records_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_records_finish(cValueInstance, cValue0, cError)
}

func Fn_g_resolver_lookup_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_service(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
}

// UNSUPPORTED : lookup_service_async : has callback

func Fn_g_resolver_lookup_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_resolver_lookup_service_finish(cValueInstance, cValue0, cError)
}

func Fn_g_resolver_set_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	C.g_resolver_set_default(cValueInstance)
}

func Fn_g_resolver_free_addresses(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.g_resolver_free_addresses(cValue0)
}

func Fn_g_resolver_free_targets(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.g_resolver_free_targets(cValue0)
}

func Fn_g_resolver_get_default() {

	C.g_resolver_get_default()
}

func Fn_g_settings_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_new(cValue0)
}

func Fn_g_settings_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {
	cValue0 := (*C.GSettingsSchema)(unsafe.Pointer(param0))
	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_settings_new_full(cValue0, cValue1, cValue2)
}

func Fn_g_settings_new_with_backend(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	C.g_settings_new_with_backend(cValue0, cValue1)
}

func Fn_g_settings_new_with_backend_and_path(param0 string, param1 unsafe.Pointer, param2 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_settings_new_with_backend_and_path(cValue0, cValue1, cValue2)
}

func Fn_g_settings_new_with_path(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_settings_new_with_path(cValue0, cValue1)
}

func Fn_g_settings_apply(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_apply(cValueInstance)
}

func Fn_g_settings_bind(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 int) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gpointer)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := (C.GSettingsBindFlags)(param3)

	C.g_settings_bind(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : bind_with_mapping : has callback

func Fn_g_settings_bind_writable(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 bool) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gpointer)(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cValue3 := toCBool(param3)

	C.g_settings_bind_writable(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_settings_create_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_create_action(cValueInstance, cValue0)
}

func Fn_g_settings_delay(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_delay(cValueInstance)
}

// UNSUPPORTED : get : has varargs

func Fn_g_settings_get_boolean(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_boolean(cValueInstance, cValue0)
}

func Fn_g_settings_get_child(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_child(cValueInstance, cValue0)
}

func Fn_g_settings_get_double(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_double(cValueInstance, cValue0)
}

func Fn_g_settings_get_enum(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_enum(cValueInstance, cValue0)
}

func Fn_g_settings_get_flags(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_flags(cValueInstance, cValue0)
}

func Fn_g_settings_get_has_unapplied(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_get_has_unapplied(cValueInstance)
}

func Fn_g_settings_get_int(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_int(cValueInstance, cValue0)
}

// UNSUPPORTED : get_mapped : has callback

func Fn_g_settings_get_range(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_range(cValueInstance, cValue0)
}

func Fn_g_settings_get_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_string(cValueInstance, cValue0)
}

func Fn_g_settings_get_strv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_strv(cValueInstance, cValue0)
}

func Fn_g_settings_get_uint(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_uint(cValueInstance, cValue0)
}

func Fn_g_settings_get_value(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_get_value(cValueInstance, cValue0)
}

func Fn_g_settings_is_writable(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_is_writable(cValueInstance, cValue0)
}

func Fn_g_settings_list_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_list_children(cValueInstance)
}

func Fn_g_settings_list_keys(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_list_keys(cValueInstance)
}

func Fn_g_settings_range_check(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_settings_range_check(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_reset(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_settings_reset(cValueInstance, cValue0)
}

func Fn_g_settings_revert(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_revert(cValueInstance)
}

// UNSUPPORTED : set : has varargs

func Fn_g_settings_set_boolean(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := toCBool(param1)

	C.g_settings_set_boolean(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gdouble)(param1)

	C.g_settings_set_double(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_enum(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gint)(param1)

	C.g_settings_set_enum(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_flags(paramInstance unsafe.Pointer, param0 string, param1 uint) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint)(param1)

	C.g_settings_set_flags(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gint)(param1)

	C.g_settings_set_int(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_settings_set_string(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_strv(paramInstance unsafe.Pointer, param0 string, param1 []string) {
	// has array param
}

func Fn_g_settings_set_uint(paramInstance unsafe.Pointer, param0 string, param1 uint) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint)(param1)

	C.g_settings_set_uint(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_set_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_settings_set_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_settings_list_relocatable_schemas() {

	C.g_settings_list_relocatable_schemas()
}

func Fn_g_settings_list_schemas() {

	C.g_settings_list_schemas()
}

func Fn_g_settings_sync() {

	C.g_settings_sync()
}

func Fn_g_settings_unbind(param0 unsafe.Pointer, param1 string) {
	cValue0 := (C.gpointer)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_settings_unbind(cValue0, cValue1)
}

// UNSUPPORTED : changed : blacklisted
// UNSUPPORTED : changed_tree : blacklisted
// UNSUPPORTED : keys_changed : blacklisted
// UNSUPPORTED : path_changed : blacklisted
// UNSUPPORTED : path_writable_changed : blacklisted
// UNSUPPORTED : writable_changed : blacklisted
// UNSUPPORTED : flatten_tree : blacklisted
// UNSUPPORTED : get_default : blacklisted
func Fn_g_simple_action_new(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	C.g_simple_action_new(cValue0, cValue1)
}

func Fn_g_simple_action_new_stateful(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))
	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_simple_action_new_stateful(cValue0, cValue1, cValue2)
}

func Fn_g_simple_action_set_enabled(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSimpleAction)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_simple_action_set_enabled(cValueInstance, cValue0)
}

func Fn_g_simple_action_set_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAction)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_simple_action_set_state(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_new() {

	C.g_simple_action_group_new()
}

func Fn_g_simple_action_group_add_entries(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 int, param2 unsafe.Pointer) {
	// has array param
}

func Fn_g_simple_action_group_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAction)(unsafe.Pointer(param0))

	C.g_simple_action_group_insert(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_lookup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_simple_action_group_lookup(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_remove(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_simple_action_group_remove(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_error : has varargs

// UNSUPPORTED : new_from_error : has callback

// UNSUPPORTED : new_take_error : has callback

func Fn_g_simple_async_result_complete(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_complete(cValueInstance)
}

func Fn_g_simple_async_result_complete_in_idle(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_complete_in_idle(cValueInstance)
}

func Fn_g_simple_async_result_get_op_res_gboolean(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_get_op_res_gboolean(cValueInstance)
}

func Fn_g_simple_async_result_get_op_res_gpointer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_get_op_res_gpointer(cValueInstance)
}

func Fn_g_simple_async_result_get_op_res_gssize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_get_op_res_gssize(cValueInstance)
}

func Fn_g_simple_async_result_get_source_tag(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_get_source_tag(cValueInstance)
}

func Fn_g_simple_async_result_propagate_error(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_simple_async_result_propagate_error(cValueInstance, cError)
}

// UNSUPPORTED : run_in_thread : has callback

func Fn_g_simple_async_result_set_check_cancellable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_simple_async_result_set_check_cancellable(cValueInstance, cValue0)
}

// UNSUPPORTED : set_error : has varargs

// UNSUPPORTED : set_error_va : has va_list

func Fn_g_simple_async_result_set_from_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_simple_async_result_set_from_error(cValueInstance, cValue0)
}

func Fn_g_simple_async_result_set_handle_cancellation(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_simple_async_result_set_handle_cancellation(cValueInstance, cValue0)
}

func Fn_g_simple_async_result_set_op_res_gboolean(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_simple_async_result_set_op_res_gboolean(cValueInstance, cValue0)
}

// UNSUPPORTED : set_op_res_gpointer : has callback

func Fn_g_simple_async_result_set_op_res_gssize(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gssize)(param0)

	C.g_simple_async_result_set_op_res_gssize(cValueInstance, cValue0)
}

func Fn_g_simple_async_result_take_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_simple_async_result_take_error(cValueInstance, cValue0)
}

func Fn_g_simple_async_result_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.GObject)(unsafe.Pointer(param1))
	cValue2 := (C.gpointer)(param2)

	C.g_simple_async_result_is_valid(cValue0, cValue1, cValue2)
}

func Fn_g_simple_permission_new(param0 bool) {
	cValue0 := toCBool(param0)

	C.g_simple_permission_new(cValue0)
}

func Fn_g_simple_proxy_resolver_set_default_proxy(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSimpleProxyResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_simple_proxy_resolver_set_default_proxy(cValueInstance, cValue0)
}

func Fn_g_simple_proxy_resolver_set_ignore_hosts(paramInstance unsafe.Pointer, param0 *string) {
	cValueInstance := (*C.GSimpleProxyResolver)(unsafe.Pointer(paramInstance))
	var cValue0String *C.gchar
	cValue0String = (*C.gchar)(C.CString(*param0))
	cValue0 := &cValue0String
	defer C.free(unsafe.Pointer(cValue0))

	C.g_simple_proxy_resolver_set_ignore_hosts(cValueInstance, cValue0)
}

func Fn_g_simple_proxy_resolver_set_uri_proxy(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GSimpleProxyResolver)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_simple_proxy_resolver_set_uri_proxy(cValueInstance, cValue0, cValue1)
}

func Fn_g_simple_proxy_resolver_new(param0 string, param1 *string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	var cValue1String *C.gchar
	cValue1String = (*C.gchar)(C.CString(*param1))
	cValue1 := &cValue1String
	defer C.free(unsafe.Pointer(cValue1))

	C.g_simple_proxy_resolver_new(cValue0, cValue1)
}

func Fn_g_socket_new(param0 int, param1 int, param2 int, error unsafe.Pointer) {
	cValue0 := (C.GSocketFamily)(param0)
	cValue1 := (C.GSocketType)(param1)
	cValue2 := (C.GSocketProtocol)(param2)
	cError := (**C.GError)(error)

	C.g_socket_new(cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_new_from_fd(param0 int, error unsafe.Pointer) {
	cValue0 := (C.gint)(param0)
	cError := (**C.GError)(error)

	C.g_socket_new_from_fd(cValue0, cError)
}

func Fn_g_socket_accept(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_accept(cValueInstance, cValue0, cError)
}

func Fn_g_socket_bind(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cError := (**C.GError)(error)

	C.g_socket_bind(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_check_connect_result(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_check_connect_result(cValueInstance, cError)
}

func Fn_g_socket_close(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_close(cValueInstance, cError)
}

func Fn_g_socket_condition_check(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GIOCondition)(param0)

	C.g_socket_condition_check(cValueInstance, cValue0)
}

func Fn_g_socket_condition_timed_wait(paramInstance unsafe.Pointer, param0 int, param1 int64, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GIOCondition)(param0)
	cValue1 := (C.gint64)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_socket_condition_timed_wait(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_condition_wait(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GIOCondition)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_condition_wait(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_connect(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_connection_factory_create_connection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_connection_factory_create_connection(cValueInstance)
}

func Fn_g_socket_create_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GIOCondition)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_socket_create_source(cValueInstance, cValue0, cValue1)
}

func Fn_g_socket_get_available_bytes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_available_bytes(cValueInstance)
}

func Fn_g_socket_get_blocking(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_blocking(cValueInstance)
}

func Fn_g_socket_get_broadcast(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_broadcast(cValueInstance)
}

func Fn_g_socket_get_credentials(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_get_credentials(cValueInstance, cError)
}

func Fn_g_socket_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_family(cValueInstance)
}

func Fn_g_socket_get_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_fd(cValueInstance)
}

func Fn_g_socket_get_keepalive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_keepalive(cValueInstance)
}

func Fn_g_socket_get_listen_backlog(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_listen_backlog(cValueInstance)
}

func Fn_g_socket_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_get_local_address(cValueInstance, cError)
}

func Fn_g_socket_get_multicast_loopback(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_multicast_loopback(cValueInstance)
}

func Fn_g_socket_get_multicast_ttl(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_multicast_ttl(cValueInstance)
}

func Fn_g_socket_get_option(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (*C.gint)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_socket_get_option(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_get_protocol(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_protocol(cValueInstance)
}

func Fn_g_socket_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_get_remote_address(cValueInstance, cError)
}

func Fn_g_socket_get_socket_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_socket_type(cValueInstance)
}

func Fn_g_socket_get_timeout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_timeout(cValueInstance)
}

func Fn_g_socket_get_ttl(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_get_ttl(cValueInstance)
}

func Fn_g_socket_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_is_closed(cValueInstance)
}

func Fn_g_socket_is_connected(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_is_connected(cValueInstance)
}

func Fn_g_socket_join_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cError := (**C.GError)(error)

	C.g_socket_join_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_leave_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))
	cValue1 := toCBool(param1)
	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))
	cError := (**C.GError)(error)

	C.g_socket_leave_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_listen(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_listen(cValueInstance, cError)
}

func Fn_g_socket_receive(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_receive_from(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_receive_message(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 []InputVector, param2 int, param3 []*unsafe.Pointer, param4 *int, param5 *int, param6 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_receive_with_blocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 bool, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_send(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_send_message(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []OutputVector, param2 int, param3 []unsafe.Pointer, param4 int, param5 int, param6 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_send_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_send_with_blocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 bool, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

func Fn_g_socket_set_blocking(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_set_blocking(cValueInstance, cValue0)
}

func Fn_g_socket_set_broadcast(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_set_broadcast(cValueInstance, cValue0)
}

func Fn_g_socket_set_keepalive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_set_keepalive(cValueInstance, cValue0)
}

func Fn_g_socket_set_listen_backlog(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_socket_set_listen_backlog(cValueInstance, cValue0)
}

func Fn_g_socket_set_multicast_loopback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_set_multicast_loopback(cValueInstance, cValue0)
}

func Fn_g_socket_set_multicast_ttl(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_socket_set_multicast_ttl(cValueInstance, cValue0)
}

func Fn_g_socket_set_option(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)
	cValue2 := (C.gint)(param2)
	cError := (**C.GError)(error)

	C.g_socket_set_option(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_socket_set_timeout(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_socket_set_timeout(cValueInstance, cValue0)
}

func Fn_g_socket_set_ttl(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_socket_set_ttl(cValueInstance, cValue0)
}

func Fn_g_socket_shutdown(paramInstance unsafe.Pointer, param0 bool, param1 bool, error unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)
	cValue1 := toCBool(param1)
	cError := (**C.GError)(error)

	C.g_socket_shutdown(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_speaks_ipv4(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	C.g_socket_speaks_ipv4(cValueInstance)
}

func Fn_g_socket_address_new_from_native(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (C.gpointer)(param0)
	cValue1 := (C.gsize)(param1)

	C.g_socket_address_new_from_native(cValue0, cValue1)
}

func Fn_g_socket_address_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_socket_address_get_family(cValueInstance)
}

func Fn_g_socket_address_get_native_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_socket_address_get_native_size(cValueInstance)
}

func Fn_g_socket_address_to_native(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gpointer)(param0)
	cValue1 := (C.gsize)(param1)
	cError := (**C.GError)(error)

	C.g_socket_address_to_native(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_address_enumerator_next(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_address_enumerator_next(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : next_async : has callback

func Fn_g_socket_address_enumerator_next_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_address_enumerator_next_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_client_new() {

	C.g_socket_client_new()
}

func Fn_g_socket_client_add_application_proxy(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_socket_client_add_application_proxy(cValueInstance, cValue0)
}

func Fn_g_socket_client_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_client_connect(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : connect_async : has callback

func Fn_g_socket_client_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_client_connect_to_host(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint16)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_host(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : connect_to_host_async : has callback

func Fn_g_socket_client_connect_to_host_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_host_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_client_connect_to_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_service(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : connect_to_service_async : has callback

func Fn_g_socket_client_connect_to_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_service_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_client_connect_to_uri(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.guint16)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_uri(cValueInstance, cValue0, cValue1, cValue2, cError)
}

// UNSUPPORTED : connect_to_uri_async : has callback

func Fn_g_socket_client_connect_to_uri_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_client_connect_to_uri_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_client_get_enable_proxy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_enable_proxy(cValueInstance)
}

func Fn_g_socket_client_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_family(cValueInstance)
}

func Fn_g_socket_client_get_local_address(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_local_address(cValueInstance)
}

func Fn_g_socket_client_get_protocol(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_protocol(cValueInstance)
}

func Fn_g_socket_client_get_proxy_resolver(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_proxy_resolver(cValueInstance)
}

func Fn_g_socket_client_get_socket_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_socket_type(cValueInstance)
}

func Fn_g_socket_client_get_timeout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_timeout(cValueInstance)
}

func Fn_g_socket_client_get_tls(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_tls(cValueInstance)
}

func Fn_g_socket_client_get_tls_validation_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	C.g_socket_client_get_tls_validation_flags(cValueInstance)
}

func Fn_g_socket_client_set_enable_proxy(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_client_set_enable_proxy(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_family(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GSocketFamily)(param0)

	C.g_socket_client_set_family(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_local_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	C.g_socket_client_set_local_address(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_protocol(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GSocketProtocol)(param0)

	C.g_socket_client_set_protocol(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_proxy_resolver(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GProxyResolver)(unsafe.Pointer(param0))

	C.g_socket_client_set_proxy_resolver(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_socket_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GSocketType)(param0)

	C.g_socket_client_set_socket_type(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_timeout(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint)(param0)

	C.g_socket_client_set_timeout(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_tls(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_socket_client_set_tls(cValueInstance, cValue0)
}

func Fn_g_socket_client_set_tls_validation_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GTlsCertificateFlags)(param0)

	C.g_socket_client_set_tls_validation_flags(cValueInstance, cValue0)
}

func Fn_g_socket_connection_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_connection_connect(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : connect_async : has callback

func Fn_g_socket_connection_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_connection_connect_finish(cValueInstance, cValue0, cError)
}

func Fn_g_socket_connection_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_connection_get_local_address(cValueInstance, cError)
}

func Fn_g_socket_connection_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_socket_connection_get_remote_address(cValueInstance, cError)
}

func Fn_g_socket_connection_get_socket(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	C.g_socket_connection_get_socket(cValueInstance)
}

func Fn_g_socket_connection_is_connected(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	C.g_socket_connection_is_connected(cValueInstance)
}

func Fn_g_socket_connection_factory_lookup_type(param0 int, param1 int, param2 int) {
	cValue0 := (C.GSocketFamily)(param0)
	cValue1 := (C.GSocketType)(param1)
	cValue2 := (C.gint)(param2)

	C.g_socket_connection_factory_lookup_type(cValue0, cValue1, cValue2)
}

func Fn_g_socket_connection_factory_register_type(param0 uint64, param1 int, param2 int, param3 int) {
	cValue0 := (C.GType)(param0)
	cValue1 := (C.GSocketFamily)(param1)
	cValue2 := (C.GSocketType)(param2)
	cValue3 := (C.gint)(param3)

	C.g_socket_connection_factory_register_type(cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_socket_control_message_get_level(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	C.g_socket_control_message_get_level(cValueInstance)
}

func Fn_g_socket_control_message_get_msg_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	C.g_socket_control_message_get_msg_type(cValueInstance)
}

func Fn_g_socket_control_message_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	C.g_socket_control_message_get_size(cValueInstance)
}

func Fn_g_socket_control_message_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gpointer)(param0)

	C.g_socket_control_message_serialize(cValueInstance, cValue0)
}

func Fn_g_socket_control_message_deserialize(param0 int, param1 int, param2 uint64, param3 []uint8) {
	// has array param
}

func Fn_g_socket_listener_new() {

	C.g_socket_listener_new()
}

func Fn_g_socket_listener_accept(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (**C.GObject)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_accept(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : accept_async : has callback

func Fn_g_socket_listener_accept_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (**C.GObject)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_accept_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_listener_accept_socket(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (**C.GObject)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_accept_socket(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : accept_socket_async : has callback

func Fn_g_socket_listener_accept_socket_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (**C.GObject)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_accept_socket_finish(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_listener_add_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 *unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))
	cValue1 := (C.GSocketType)(param1)
	cValue2 := (C.GSocketProtocol)(param2)
	cValue3 := (*C.GObject)(unsafe.Pointer(param3))
	cValue4 := (**C.GSocketAddress)(unsafe.Pointer(param4))
	cError := (**C.GError)(error)

	C.g_socket_listener_add_address(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)
}

func Fn_g_socket_listener_add_any_inet_port(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_socket_listener_add_any_inet_port(cValueInstance, cValue0, cError)
}

func Fn_g_socket_listener_add_inet_port(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint16)(param0)
	cValue1 := (*C.GObject)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_add_inet_port(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_listener_add_socket(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocket)(unsafe.Pointer(param0))
	cValue1 := (*C.GObject)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_socket_listener_add_socket(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_socket_listener_close(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	C.g_socket_listener_close(cValueInstance)
}

func Fn_g_socket_listener_set_backlog(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.g_socket_listener_set_backlog(cValueInstance, cValue0)
}

func Fn_g_socket_service_new() {

	C.g_socket_service_new()
}

func Fn_g_socket_service_is_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_is_active(cValueInstance)
}

func Fn_g_socket_service_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_start(cValueInstance)
}

func Fn_g_socket_service_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_stop(cValueInstance)
}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : communicate_async : has callback

func Fn_g_subprocess_communicate_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (**C.GBytes)(unsafe.Pointer(param1))
	cValue2 := (**C.GBytes)(unsafe.Pointer(param2))
	cError := (**C.GError)(error)

	C.g_subprocess_communicate_finish(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_g_subprocess_communicate_utf8(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 *string, param3 *string, error unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	var cValue2String *C.gchar
	cValue2 := &cValue2String
	var cValue3String *C.gchar
	cValue3 := &cValue3String
	cError := (**C.GError)(error)

	C.g_subprocess_communicate_utf8(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
	param2String := C.GoString(cValue2String)
	*param2 = param2String
	param3String := C.GoString(cValue3String)
	*param3 = param3String
}

// UNSUPPORTED : communicate_utf8_async : has callback

func Fn_g_subprocess_communicate_utf8_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *string, param2 *string, error unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	var cValue1String *C.gchar
	cValue1 := &cValue1String
	var cValue2String *C.gchar
	cValue2 := &cValue2String
	cError := (**C.GError)(error)

	C.g_subprocess_communicate_utf8_finish(cValueInstance, cValue0, cValue1, cValue2, cError)
	param1String := C.GoString(cValue1String)
	*param1 = param1String
	param2String := C.GoString(cValue2String)
	*param2 = param2String
}

// UNSUPPORTED : wait_async : has callback

// UNSUPPORTED : wait_check_async : has callback

// UNSUPPORTED : set_child_setup : has callback

// UNSUPPORTED : spawn : has varargs

func Fn_g_subprocess_launcher_take_fd(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (C.gint)(param1)

	C.g_subprocess_launcher_take_fd(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : attach_source : has callback

func Fn_g_task_get_cancellable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_cancellable(cValueInstance)
}

func Fn_g_task_get_check_cancellable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_check_cancellable(cValueInstance)
}

func Fn_g_task_get_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_context(cValueInstance)
}

func Fn_g_task_get_priority(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_priority(cValueInstance)
}

func Fn_g_task_get_return_on_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_return_on_cancel(cValueInstance)
}

func Fn_g_task_get_source_object(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_source_object(cValueInstance)
}

func Fn_g_task_get_source_tag(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_source_tag(cValueInstance)
}

func Fn_g_task_get_task_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_get_task_data(cValueInstance)
}

func Fn_g_task_had_error(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_had_error(cValueInstance)
}

func Fn_g_task_propagate_boolean(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_task_propagate_boolean(cValueInstance, cError)
}

func Fn_g_task_propagate_int(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_task_propagate_int(cValueInstance, cError)
}

func Fn_g_task_propagate_pointer(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cError := (**C.GError)(error)

	C.g_task_propagate_pointer(cValueInstance, cError)
}

func Fn_g_task_return_boolean(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_task_return_boolean(cValueInstance, cValue0)
}

func Fn_g_task_return_error(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GError)(unsafe.Pointer(param0))

	C.g_task_return_error(cValueInstance, cValue0)
}

func Fn_g_task_return_error_if_cancelled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	C.g_task_return_error_if_cancelled(cValueInstance)
}

func Fn_g_task_return_int(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gssize)(param0)

	C.g_task_return_int(cValueInstance, cValue0)
}

// UNSUPPORTED : return_new_error : has varargs

// UNSUPPORTED : return_pointer : has callback

// UNSUPPORTED : run_in_thread : has callback

// UNSUPPORTED : run_in_thread_sync : has callback

func Fn_g_task_set_check_cancellable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_task_set_check_cancellable(cValueInstance, cValue0)
}

func Fn_g_task_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)

	C.g_task_set_priority(cValueInstance, cValue0)
}

func Fn_g_task_set_return_on_cancel(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_task_set_return_on_cancel(cValueInstance, cValue0)
}

func Fn_g_task_set_source_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gpointer)(param0)

	C.g_task_set_source_tag(cValueInstance, cValue0)
}

// UNSUPPORTED : set_task_data : has callback

func Fn_g_task_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)
	cValue1 := (C.gpointer)(param1)

	C.g_task_is_valid(cValue0, cValue1)
}

// UNSUPPORTED : report_error : has callback

// UNSUPPORTED : report_new_error : has varargs

func Fn_g_tcp_connection_get_graceful_disconnect(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))

	C.g_tcp_connection_get_graceful_disconnect(cValueInstance)
}

func Fn_g_tcp_connection_set_graceful_disconnect(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_tcp_connection_set_graceful_disconnect(cValueInstance, cValue0)
}

func Fn_g_tcp_wrapper_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GSocket)(unsafe.Pointer(param1))

	C.g_tcp_wrapper_connection_new(cValue0, cValue1)
}

func Fn_g_tcp_wrapper_connection_get_base_io_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTcpWrapperConnection)(unsafe.Pointer(paramInstance))

	C.g_tcp_wrapper_connection_get_base_io_stream(cValueInstance)
}

func Fn_g_test_dbus_new(param0 int) {
	cValue0 := (C.GTestDBusFlags)(param0)

	C.g_test_dbus_new(cValue0)
}

func Fn_g_test_dbus_add_service_dir(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_test_dbus_add_service_dir(cValueInstance, cValue0)
}

func Fn_g_test_dbus_down(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	C.g_test_dbus_down(cValueInstance)
}

func Fn_g_test_dbus_get_bus_address(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	C.g_test_dbus_get_bus_address(cValueInstance)
}

func Fn_g_test_dbus_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	C.g_test_dbus_get_flags(cValueInstance)
}

func Fn_g_test_dbus_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	C.g_test_dbus_stop(cValueInstance)
}

func Fn_g_test_dbus_up(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	C.g_test_dbus_up(cValueInstance)
}

func Fn_g_test_dbus_unset() {

	C.g_test_dbus_unset()
}

func Fn_g_themed_icon_new(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_new(cValue0)
}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) {
	// has array param
}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_new_with_default_fallbacks(cValue0)
}

func Fn_g_themed_icon_append_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_append_name(cValueInstance, cValue0)
}

func Fn_g_themed_icon_get_names(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))

	C.g_themed_icon_get_names(cValueInstance)
}

func Fn_g_themed_icon_prepend_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_prepend_name(cValueInstance, cValue0)
}

func Fn_g_threaded_socket_service_new(param0 int) {
	cValue0 := (C.int)(param0)

	C.g_threaded_socket_service_new(cValue0)
}

func Fn_g_tls_certificate_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_tls_certificate_new_from_file(cValue0, cError)
}

func Fn_g_tls_certificate_new_from_files(param0 string, param1 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cError := (**C.GError)(error)

	C.g_tls_certificate_new_from_files(cValue0, cValue1, cError)
}

func Fn_g_tls_certificate_new_from_pem(param0 string, param1 uint64, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (C.gssize)(param1)
	cError := (**C.GError)(error)

	C.g_tls_certificate_new_from_pem(cValue0, cValue1, cError)
}

func Fn_g_tls_certificate_get_issuer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))

	C.g_tls_certificate_get_issuer(cValueInstance)
}

func Fn_g_tls_certificate_is_same(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	C.g_tls_certificate_is_same(cValueInstance, cValue0)
}

func Fn_g_tls_certificate_verify(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))
	cValue1 := (*C.GTlsCertificate)(unsafe.Pointer(param1))

	C.g_tls_certificate_verify(cValueInstance, cValue0, cValue1)
}

func Fn_g_tls_certificate_list_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cError := (**C.GError)(error)

	C.g_tls_certificate_list_new_from_file(cValue0, cError)
}

func Fn_g_tls_connection_emit_accept_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))
	cValue1 := (C.GTlsCertificateFlags)(param1)

	C.g_tls_connection_emit_accept_certificate(cValueInstance, cValue0, cValue1)
}

func Fn_g_tls_connection_get_certificate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_certificate(cValueInstance)
}

func Fn_g_tls_connection_get_database(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_database(cValueInstance)
}

func Fn_g_tls_connection_get_interaction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_interaction(cValueInstance)
}

func Fn_g_tls_connection_get_peer_certificate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_peer_certificate(cValueInstance)
}

func Fn_g_tls_connection_get_peer_certificate_errors(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_peer_certificate_errors(cValueInstance)
}

func Fn_g_tls_connection_get_rehandshake_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_rehandshake_mode(cValueInstance)
}

func Fn_g_tls_connection_get_require_close_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_require_close_notify(cValueInstance)
}

func Fn_g_tls_connection_get_use_system_certdb(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_use_system_certdb(cValueInstance)
}

func Fn_g_tls_connection_handshake(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_connection_handshake(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : handshake_async : has callback

func Fn_g_tls_connection_handshake_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_connection_handshake_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_connection_set_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	C.g_tls_connection_set_certificate(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_database(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsDatabase)(unsafe.Pointer(param0))

	C.g_tls_connection_set_database(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_interaction(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsInteraction)(unsafe.Pointer(param0))

	C.g_tls_connection_set_interaction(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_rehandshake_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GTlsRehandshakeMode)(param0)

	C.g_tls_connection_set_rehandshake_mode(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_require_close_notify(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_tls_connection_set_require_close_notify(cValueInstance, cValue0)
}

func Fn_g_tls_connection_set_use_system_certdb(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_tls_connection_set_use_system_certdb(cValueInstance, cValue0)
}

func Fn_g_tls_database_create_certificate_handle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	C.g_tls_database_create_certificate_handle(cValueInstance, cValue0)
}

func Fn_g_tls_database_lookup_certificate_for_handle(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))
	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))
	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)
	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_tls_database_lookup_certificate_for_handle(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
}

// UNSUPPORTED : lookup_certificate_for_handle_async : has callback

func Fn_g_tls_database_lookup_certificate_for_handle_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_database_lookup_certificate_for_handle_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_database_lookup_certificate_issuer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))
	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))
	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)
	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))
	cError := (**C.GError)(error)

	C.g_tls_database_lookup_certificate_issuer(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
}

// UNSUPPORTED : lookup_certificate_issuer_async : has callback

func Fn_g_tls_database_lookup_certificate_issuer_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_database_lookup_certificate_issuer_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_database_lookup_certificates_issued_by(paramInstance unsafe.Pointer, param0 []uint8, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : lookup_certificates_issued_by_async : has callback

func Fn_g_tls_database_lookup_certificates_issued_by_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_database_lookup_certificates_issued_by_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_database_verify_chain(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int, param5 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))
	cValue2 := (*C.GSocketConnectable)(unsafe.Pointer(param2))
	cValue3 := (*C.GTlsInteraction)(unsafe.Pointer(param3))
	cValue4 := (C.GTlsDatabaseVerifyFlags)(param4)
	cValue5 := (*C.GCancellable)(unsafe.Pointer(param5))
	cError := (**C.GError)(error)

	C.g_tls_database_verify_chain(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cError)
}

// UNSUPPORTED : verify_chain_async : has callback

func Fn_g_tls_database_verify_chain_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_database_verify_chain_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_interaction_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_tls_interaction_ask_password(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : ask_password_async : has callback

func Fn_g_tls_interaction_ask_password_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_tls_interaction_ask_password_finish(cValueInstance, cValue0, cError)
}

func Fn_g_tls_interaction_invoke_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_tls_interaction_invoke_ask_password(cValueInstance, cValue0, cValue1, cError)
}

// UNSUPPORTED : request_certificate_async : has callback

func Fn_g_tls_password_new(param0 int, param1 string) {
	cValue0 := (C.GTlsPasswordFlags)(param0)
	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_tls_password_new(cValue0, cValue1)
}

func Fn_g_tls_password_get_description(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	C.g_tls_password_get_description(cValueInstance)
}

func Fn_g_tls_password_get_flags(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	C.g_tls_password_get_flags(cValueInstance)
}

func Fn_g_tls_password_get_value(paramInstance unsafe.Pointer, param0 *uint64) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	C.g_tls_password_get_value(cValueInstance, cValue0)
}

func Fn_g_tls_password_get_warning(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	C.g_tls_password_get_warning(cValueInstance)
}

func Fn_g_tls_password_set_description(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_tls_password_set_description(cValueInstance, cValue0)
}

func Fn_g_tls_password_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GTlsPasswordFlags)(param0)

	C.g_tls_password_set_flags(cValueInstance, cValue0)
}

func Fn_g_tls_password_set_value(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64) {
	// has array param
}

// UNSUPPORTED : set_value_full : has callback

func Fn_g_tls_password_set_warning(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_tls_password_set_warning(cValueInstance, cValue0)
}

func Fn_g_unix_connection_receive_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_unix_connection_receive_credentials(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : receive_credentials_async : has callback

func Fn_g_unix_connection_receive_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_unix_connection_receive_credentials_finish(cValueInstance, cValue0, cError)
}

func Fn_g_unix_connection_receive_fd(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_unix_connection_receive_fd(cValueInstance, cValue0, cError)
}

func Fn_g_unix_connection_send_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_unix_connection_send_credentials(cValueInstance, cValue0, cError)
}

// UNSUPPORTED : send_credentials_async : has callback

func Fn_g_unix_connection_send_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cError := (**C.GError)(error)

	C.g_unix_connection_send_credentials_finish(cValueInstance, cValue0, cError)
}

func Fn_g_unix_connection_send_fd(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))
	cError := (**C.GError)(error)

	C.g_unix_connection_send_fd(cValueInstance, cValue0, cValue1, cError)
}

func Fn_g_unix_credentials_message_new() {

	C.g_unix_credentials_message_new()
}

func Fn_g_unix_credentials_message_new_with_credentials(param0 unsafe.Pointer) {
	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))

	C.g_unix_credentials_message_new_with_credentials(cValue0)
}

func Fn_g_unix_credentials_message_get_credentials(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixCredentialsMessage)(unsafe.Pointer(paramInstance))

	C.g_unix_credentials_message_get_credentials(cValueInstance)
}

func Fn_g_unix_credentials_message_is_supported() {

	C.g_unix_credentials_message_is_supported()
}

func Fn_g_unix_fd_list_new() {

	C.g_unix_fd_list_new()
}

func Fn_g_unix_fd_list_new_from_array(param0 []int, param1 int) {
	// has array param
}

func Fn_g_unix_fd_list_append(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cError := (**C.GError)(error)

	C.g_unix_fd_list_append(cValueInstance, cValue0, cError)
}

func Fn_g_unix_fd_list_get(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cError := (**C.GError)(error)

	C.g_unix_fd_list_get(cValueInstance, cValue0, cError)
}

func Fn_g_unix_fd_list_get_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))

	C.g_unix_fd_list_get_length(cValueInstance)
}

func Fn_g_unix_fd_list_peek_fds(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.g_unix_fd_list_peek_fds(cValueInstance, cValue0)
}

func Fn_g_unix_fd_list_steal_fds(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.g_unix_fd_list_steal_fds(cValueInstance, cValue0)
}

func Fn_g_unix_fd_message_new() {

	C.g_unix_fd_message_new()
}

func Fn_g_unix_fd_message_new_with_fd_list(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixFDList)(unsafe.Pointer(param0))

	C.g_unix_fd_message_new_with_fd_list(cValue0)
}

func Fn_g_unix_fd_message_append_fd(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint)(param0)
	cError := (**C.GError)(error)

	C.g_unix_fd_message_append_fd(cValueInstance, cValue0, cError)
}

func Fn_g_unix_fd_message_get_fd_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))

	C.g_unix_fd_message_get_fd_list(cValueInstance)
}

func Fn_g_unix_fd_message_steal_fds(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.g_unix_fd_message_steal_fds(cValueInstance, cValue0)
}

func Fn_g_unix_input_stream_new(param0 int, param1 bool) {
	cValue0 := (C.gint)(param0)
	cValue1 := toCBool(param1)

	C.g_unix_input_stream_new(cValue0, cValue1)
}

func Fn_g_unix_input_stream_get_close_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))

	C.g_unix_input_stream_get_close_fd(cValueInstance)
}

func Fn_g_unix_input_stream_get_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))

	C.g_unix_input_stream_get_fd(cValueInstance)
}

func Fn_g_unix_input_stream_set_close_fd(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_unix_input_stream_set_close_fd(cValueInstance, cValue0)
}

func Fn_g_unix_mount_monitor_new() {

	C.g_unix_mount_monitor_new()
}

func Fn_g_unix_mount_monitor_set_rate_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GUnixMountMonitor)(unsafe.Pointer(paramInstance))
	cValue0 := (C.int)(param0)

	C.g_unix_mount_monitor_set_rate_limit(cValueInstance, cValue0)
}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) {
	cValue0 := (C.gint)(param0)
	cValue1 := toCBool(param1)

	C.g_unix_output_stream_new(cValue0, cValue1)
}

func Fn_g_unix_output_stream_get_close_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))

	C.g_unix_output_stream_get_close_fd(cValueInstance)
}

func Fn_g_unix_output_stream_get_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))

	C.g_unix_output_stream_get_fd(cValueInstance)
}

func Fn_g_unix_output_stream_set_close_fd(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_unix_output_stream_set_close_fd(cValueInstance, cValue0)
}

func Fn_g_unix_socket_address_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_unix_socket_address_new(cValue0)
}

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) {
	// has array param
}

func Fn_g_unix_socket_address_new_with_type(param0 []int8, param1 int, param2 int) {
	// has array param
}

func Fn_g_unix_socket_address_get_address_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_unix_socket_address_get_address_type(cValueInstance)
}

func Fn_g_unix_socket_address_get_is_abstract(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_unix_socket_address_get_is_abstract(cValueInstance)
}

func Fn_g_unix_socket_address_get_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_unix_socket_address_get_path(cValueInstance)
}

func Fn_g_unix_socket_address_get_path_len(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	C.g_unix_socket_address_get_path_len(cValueInstance)
}

func Fn_g_unix_socket_address_abstract_names_supported() {

	C.g_unix_socket_address_abstract_names_supported()
}

func Fn_g_vfs_get_file_for_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_vfs_get_file_for_path(cValueInstance, cValue0)
}

func Fn_g_vfs_get_file_for_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_vfs_get_file_for_uri(cValueInstance, cValue0)
}

func Fn_g_vfs_get_supported_uri_schemes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	C.g_vfs_get_supported_uri_schemes(cValueInstance)
}

func Fn_g_vfs_is_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	C.g_vfs_is_active(cValueInstance)
}

func Fn_g_vfs_parse_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_vfs_parse_name(cValueInstance, cValue0)
}

// UNSUPPORTED : register_uri_scheme : has callback

func Fn_g_vfs_get_default() {

	C.g_vfs_get_default()
}

func Fn_g_vfs_get_local() {

	C.g_vfs_get_local()
}

func Fn_g_volume_monitor_get_connected_drives(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	C.g_volume_monitor_get_connected_drives(cValueInstance)
}

func Fn_g_volume_monitor_get_mount_for_uuid(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_volume_monitor_get_mount_for_uuid(cValueInstance, cValue0)
}

func Fn_g_volume_monitor_get_mounts(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	C.g_volume_monitor_get_mounts(cValueInstance)
}

func Fn_g_volume_monitor_get_volume_for_uuid(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_volume_monitor_get_volume_for_uuid(cValueInstance, cValue0)
}

func Fn_g_volume_monitor_get_volumes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	C.g_volume_monitor_get_volumes(cValueInstance)
}

func Fn_g_volume_monitor_adopt_orphan_mount(param0 unsafe.Pointer) {
	cValue0 := (*C.GMount)(unsafe.Pointer(param0))

	C.g_volume_monitor_adopt_orphan_mount(cValue0)
}

func Fn_g_volume_monitor_get() {

	C.g_volume_monitor_get()
}

func Fn_g_zlib_compressor_new(param0 int, param1 int) {
	cValue0 := (C.GZlibCompressorFormat)(param0)
	cValue1 := (C.int)(param1)

	C.g_zlib_compressor_new(cValue0, cValue1)
}

func Fn_g_zlib_compressor_get_file_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))

	C.g_zlib_compressor_get_file_info(cValueInstance)
}

func Fn_g_zlib_compressor_set_file_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	C.g_zlib_compressor_set_file_info(cValueInstance, cValue0)
}

func Fn_g_zlib_decompressor_new(param0 int) {
	cValue0 := (C.GZlibCompressorFormat)(param0)

	C.g_zlib_decompressor_new(cValue0)
}

func Fn_g_zlib_decompressor_get_file_info(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GZlibDecompressor)(unsafe.Pointer(paramInstance))

	C.g_zlib_decompressor_get_file_info(cValueInstance)
}
