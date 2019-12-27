// Code generated - DO NOT EDIT.
// +build gio_2.46

package gio

import "unsafe"

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
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
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

func Fn_g_dbus_annotation_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusAnnotationInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_annotation_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_annotation_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusAnnotationInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_annotation_info_unref(cValueInstance)
}

// UNSUPPORTED : g_dbus_annotation_info_lookup : has non-string array param annotations

func Fn_g_dbus_arg_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusArgInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_arg_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_arg_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusArgInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_arg_info_unref(cValueInstance)
}

func Fn_g_dbus_interface_info_cache_build(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_cache_build(cValueInstance)
}

func Fn_g_dbus_interface_info_cache_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_cache_release(cValueInstance)
}

func Fn_g_dbus_interface_info_generate_xml(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	C.g_dbus_interface_info_generate_xml(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_interface_info_lookup_method(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_method(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_signal(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_interface_info_lookup_signal(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_info_unref(cValueInstance)
}

func Fn_g_dbus_method_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMethodInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_method_info_unref(cValueInstance)
}

func Fn_g_dbus_node_info_new_for_xml(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_node_info_new_for_xml(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_generate_xml(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	C.g_dbus_node_info_generate_xml(cValueInstance, cValue0, cValue1)
}

func Fn_g_dbus_node_info_lookup_interface(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_node_info_lookup_interface(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_node_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusNodeInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_node_info_unref(cValueInstance)
}

func Fn_g_dbus_property_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusPropertyInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_property_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_property_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusPropertyInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_property_info_unref(cValueInstance)
}

func Fn_g_dbus_signal_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusSignalInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_signal_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_signal_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusSignalInfo)(unsafe.Pointer(paramInstance))

	C.g_dbus_signal_info_unref(cValueInstance)
}

func Fn_g_file_attribute_info_list_new() unsafe.Pointer {
	ret := C.g_file_attribute_info_list_new()

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_add(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
	cValueInstance := (*C.GFileAttributeInfoList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileAttributeType)(param1)

	cValue2 := (C.GFileAttributeInfoFlags)(param2)

	C.g_file_attribute_info_list_add(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_file_attribute_info_list_dup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileAttributeInfoList)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_info_list_dup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_lookup(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GFileAttributeInfoList)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_attribute_info_list_lookup(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileAttributeInfoList)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_info_list_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileAttributeInfoList)(unsafe.Pointer(paramInstance))

	C.g_file_attribute_info_list_unref(cValueInstance)
}

func Fn_g_file_attribute_matcher_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_attribute_matcher_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_enumerate_namespace(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_attribute_matcher_enumerate_namespace(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_enumerate_next(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_matcher_enumerate_next(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_attribute_matcher_matches(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_attribute_matcher_matches(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_matches_only(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_attribute_matcher_matches_only(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_matcher_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_subtract(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFileAttributeMatcher)(unsafe.Pointer(param0))

	ret := C.g_file_attribute_matcher_subtract(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	ret := C.g_file_attribute_matcher_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_attribute_matcher_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileAttributeMatcher)(unsafe.Pointer(paramInstance))

	C.g_file_attribute_matcher_unref(cValueInstance)
}

func Fn_g_io_extension_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GIOExtension)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_io_extension_get_priority(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GIOExtension)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_get_priority(cValueInstance)

	return (int)(ret)
}

func Fn_g_io_extension_get_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GIOExtension)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_get_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_io_extension_ref_class(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOExtension)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_ref_class(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_extension_by_name(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GIOExtensionPoint)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_io_extension_point_get_extension_by_name(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_extensions(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOExtensionPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_point_get_extensions(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_required_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GIOExtensionPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_io_extension_point_get_required_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_io_extension_point_set_required_type(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GIOExtensionPoint)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.g_io_extension_point_set_required_type(cValueInstance, cValue0)
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gint)(param3)

	ret := C.g_io_extension_point_implement(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_lookup(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_io_extension_point_lookup(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_register(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_io_extension_point_register(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_io_module_scope_block : blacklisted
// UNSUPPORTED : g_io_module_scope_free : blacklisted
// UNSUPPORTED : g_io_module_scope_new : blacklisted
// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : has callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : has callback

func Fn_g_resource_new_from_data(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resource_new_from_data(cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resource_enumerate_children : has array return

func Fn_g_resource_get_info(paramInstance unsafe.Pointer, param0 string, param1 int, param2 *uint64, param3 *uint32, error unsafe.Pointer) bool {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resource_get_info(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_resource_lookup_data(paramInstance unsafe.Pointer, param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resource_lookup_data(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_open_stream(paramInstance unsafe.Pointer, param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resource_open_stream(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	ret := C.g_resource_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_resource_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GResource)(unsafe.Pointer(paramInstance))

	C.g_resource_unref(cValueInstance)
}

func Fn_g_resource_load(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_resource_load(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_get_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_get_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_settings_schema_get_key(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_schema_get_key(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_settings_schema_has_key(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_schema_has_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_schema_list_children : has array return

// UNSUPPORTED : g_settings_schema_list_keys : has array return

func Fn_g_settings_schema_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_unref(cValueInstance)
}

func Fn_g_settings_schema_key_get_default_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_default_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_settings_schema_key_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_settings_schema_key_get_range(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_range(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_get_summary(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_summary(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_settings_schema_key_get_value_type(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_get_value_type(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_range_check(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	ret := C.g_settings_schema_key_range_check(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_settings_schema_key_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_key_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_key_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchemaKey)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_key_unref(cValueInstance)
}

func Fn_g_settings_schema_source_new_from_directory(param0 string, param1 unsafe.Pointer, param2 bool, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsSchemaSource)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cError := (**C.GError)(error)

	ret := C.g_settings_schema_source_new_from_directory(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted
func Fn_g_settings_schema_source_lookup(paramInstance unsafe.Pointer, param0 string, param1 bool) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.g_settings_schema_source_lookup(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_source_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_source_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_schema_source_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettingsSchemaSource)(unsafe.Pointer(paramInstance))

	C.g_settings_schema_source_unref(cValueInstance)
}

func Fn_g_settings_schema_source_get_default() unsafe.Pointer {
	ret := C.g_settings_schema_source_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_new(param0 string, param1 uint16, param2 uint16, param3 uint16) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (C.guint16)(param2)

	cValue3 := (C.guint16)(param3)

	ret := C.g_srv_target_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	C.g_srv_target_free(cValueInstance)
}

func Fn_g_srv_target_get_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_srv_target_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_priority(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_priority(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_weight(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSrvTarget)(unsafe.Pointer(paramInstance))

	ret := C.g_srv_target_get_weight(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	ret := C.g_srv_target_list_sort(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_static_resource_fini(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	C.g_static_resource_fini(cValueInstance)
}

func Fn_g_static_resource_get_resource(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	ret := C.g_static_resource_get_resource(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_static_resource_init(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GStaticResource)(unsafe.Pointer(paramInstance))

	C.g_static_resource_init(cValueInstance)
}

func Fn_g_unix_mount_point_compare(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GUnixMountPoint)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_point_compare(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_unix_mount_point_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	C.g_unix_mount_point_free(cValueInstance)
}

func Fn_g_unix_mount_point_get_device_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_get_device_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_get_fs_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_get_fs_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_get_mount_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_get_mount_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_get_options(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_get_options(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_guess_can_eject(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_guess_can_eject(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_guess_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_guess_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_point_guess_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_guess_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_guess_symbolic_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_guess_symbolic_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_point_is_loopback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_is_loopback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_is_readonly(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_is_readonly(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_is_user_mountable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixMountPoint)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_mount_point_is_user_mountable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : has callback

// UNSUPPORTED : g_async_initable_newv_async : has callback

// UNSUPPORTED : g_bus_get : has callback

func Fn_g_bus_get_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_bus_get_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_sync(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_bus_own_name : has callback

// UNSUPPORTED : g_bus_own_name_on_connection : has callback

func Fn_g_bus_own_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameOwnerFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_own_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

func Fn_g_bus_own_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) uint {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameOwnerFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cValue5 := (*C.GClosure)(unsafe.Pointer(param5))

	ret := C.g_bus_own_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return (uint)(ret)
}

func Fn_g_bus_unown_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unown_name(cValue0)
}

func Fn_g_bus_unwatch_name(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.g_bus_unwatch_name(cValue0)
}

// UNSUPPORTED : g_bus_watch_name : has callback

// UNSUPPORTED : g_bus_watch_name_on_connection : has callback

func Fn_g_bus_watch_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameWatcherFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_watch_name_on_connection_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

func Fn_g_bus_watch_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) uint {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GBusNameWatcherFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	ret := C.g_bus_watch_name_with_closures(cValue0, cValue1, cValue2, cValue3, cValue4)

	return (uint)(ret)
}

func Fn_g_content_type_can_be_executable(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_can_be_executable(cValue0)

	return toGoBool(ret)
}

func Fn_g_content_type_equals(param0 string, param1 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_content_type_equals(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_content_type_from_mime_type(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_from_mime_type(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_get_description(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_description(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_get_generic_icon_name(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_generic_icon_name(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_get_icon(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_icon(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_content_type_get_mime_dirs : has array return

func Fn_g_content_type_get_mime_type(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_mime_type(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_get_symbolic_icon(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_symbolic_icon(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_content_type_guess : has non-string array param data

// UNSUPPORTED : g_content_type_guess_for_tree : has array return

func Fn_g_content_type_is_a(param0 string, param1 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_content_type_is_a(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_content_type_is_unknown(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_is_unknown(cValue0)

	return toGoBool(ret)
}

func Fn_g_content_types_get_registered() unsafe.Pointer {
	ret := C.g_content_types_get_registered()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_address_escape_value(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_address_escape_value(cValue0)

	return C.GoString(ret)
}

func Fn_g_dbus_address_get_for_bus_sync(param0 int, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValue0 := (C.GBusType)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_for_bus_sync(cValue0, cValue1, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_dbus_address_get_stream : has callback

func Fn_g_dbus_address_get_stream_finish(param0 unsafe.Pointer, param1 *string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_stream_finish(cValue0, cValue1, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_address_get_stream_sync(param0 string, param1 *string, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_stream_sync(cValue0, cValue1, cValue2, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_annotation_info_lookup : has non-string array param annotations

// UNSUPPORTED : g_dbus_error_register_error_domain : has non-string array param entries

func Fn_g_dbus_generate_guid() string {
	ret := C.g_dbus_generate_guid()

	return C.GoString(ret)
}

func Fn_g_dbus_gvalue_to_gvariant(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GValue)(unsafe.Pointer(param0))

	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	ret := C.g_dbus_gvalue_to_gvariant(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_gvariant_to_gvalue(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.g_dbus_gvariant_to_gvalue(cValue0, cValue1)
}

func Fn_g_dbus_is_address(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_address(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_guid(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_guid(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_interface_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_interface_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_member_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_member_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_name(cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_is_supported_address(param0 string, error unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_is_supported_address(cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_is_unique_name(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_is_unique_name(cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_initable_newv : has non-string array param parameters

func Fn_g_io_error_from_errno(param0 int) int {
	cValue0 := (C.gint)(param0)

	ret := C.g_io_error_from_errno(cValue0)

	return (int)(ret)
}

func Fn_g_io_error_quark() uint32 {
	ret := C.g_io_error_quark()

	return (uint32)(ret)
}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted
// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted
// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted
// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted
func Fn_g_io_scheduler_cancel_all_jobs() {
	C.g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : g_io_scheduler_push_job : has callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted
// UNSUPPORTED : g_memory_settings_backend_new : blacklisted
func Fn_g_networking_init() {
	C.g_networking_init()
}

// UNSUPPORTED : g_null_settings_backend_new : blacklisted
func Fn_g_pollable_source_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.g_pollable_source_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_pollable_source_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.GSource)(unsafe.Pointer(param1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	ret := C.g_pollable_source_new_full(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_pollable_stream_read : has non-string array param buffer

// UNSUPPORTED : g_pollable_stream_write : has non-string array param buffer

// UNSUPPORTED : g_pollable_stream_write_all : has non-string array param buffer

// UNSUPPORTED : g_resources_enumerate_children : has array return

func Fn_g_resources_get_info(param0 string, param1 int, param2 *uint64, param3 *uint32, error unsafe.Pointer) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resources_get_info(cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_resources_lookup_data(param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resources_lookup_data(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resources_open_stream(param0 string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResourceLookupFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_resources_open_stream(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resources_register(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_register(cValue0)
}

func Fn_g_resources_unregister(param0 unsafe.Pointer) {
	cValue0 := (*C.GResource)(unsafe.Pointer(param0))

	C.g_resources_unregister(cValue0)
}

// UNSUPPORTED : g_simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : has callback

func Fn_g_unix_is_mount_path_system_internal(param0 string) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_unix_is_mount_path_system_internal(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_at(param0 string, param1 *uint64) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.guint64)(unsafe.Pointer(param1))

	ret := C.g_unix_mount_at(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_compare(param0 unsafe.Pointer, param1 unsafe.Pointer) int {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	cValue1 := (*C.GUnixMountEntry)(unsafe.Pointer(param1))

	ret := C.g_unix_mount_compare(cValue0, cValue1)

	return (int)(ret)
}

func Fn_g_unix_mount_free(param0 unsafe.Pointer) {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	C.g_unix_mount_free(cValue0)
}

func Fn_g_unix_mount_get_device_path(param0 unsafe.Pointer) string {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_get_device_path(cValue0)

	return C.GoString(ret)
}

func Fn_g_unix_mount_get_fs_type(param0 unsafe.Pointer) string {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_get_fs_type(cValue0)

	return C.GoString(ret)
}

func Fn_g_unix_mount_get_mount_path(param0 unsafe.Pointer) string {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_get_mount_path(cValue0)

	return C.GoString(ret)
}

func Fn_g_unix_mount_guess_can_eject(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_guess_can_eject(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_guess_icon(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_guess_icon(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_guess_name(param0 unsafe.Pointer) string {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_guess_name(cValue0)

	return C.GoString(ret)
}

func Fn_g_unix_mount_guess_should_display(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_guess_should_display(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_guess_symbolic_icon(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_guess_symbolic_icon(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_is_readonly(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_is_readonly(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_is_system_internal(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GUnixMountEntry)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_is_system_internal(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_points_changed_since(param0 uint64) bool {
	cValue0 := (C.guint64)(param0)

	ret := C.g_unix_mount_points_changed_since(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mount_points_get(param0 *uint64) unsafe.Pointer {
	cValue0 := (*C.guint64)(unsafe.Pointer(param0))

	ret := C.g_unix_mount_points_get(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mounts_changed_since(param0 uint64) bool {
	cValue0 := (C.guint64)(param0)

	ret := C.g_unix_mounts_changed_since(cValue0)

	return toGoBool(ret)
}

func Fn_g_unix_mounts_get(param0 *uint64) unsafe.Pointer {
	cValue0 := (*C.guint64)(unsafe.Pointer(param0))

	ret := C.g_unix_mounts_get(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_monitor_get() unsafe.Pointer {
	ret := C.g_app_info_monitor_get()

	return unsafe.Pointer(ret)
}

func Fn_g_app_launch_context_new() unsafe.Pointer {
	ret := C.g_app_launch_context_new()

	return unsafe.Pointer(ret)
}

func Fn_g_app_launch_context_get_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) string {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))

	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	ret := C.g_app_launch_context_get_display(cValueInstance, cValue0, cValue1)

	return C.GoString(ret)
}

// UNSUPPORTED : g_app_launch_context_get_environment : has array return

func Fn_g_app_launch_context_get_startup_notify_id(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) string {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))

	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	ret := C.g_app_launch_context_get_startup_notify_id(cValueInstance, cValue0, cValue1)

	return C.GoString(ret)
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

func Fn_g_application_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GApplicationFlags)(param1)

	ret := C.g_application_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_application_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_activate(cValueInstance)
}

func Fn_g_application_add_main_option(paramInstance unsafe.Pointer, param0 string, param1 int8, param2 int, param3 int, param4 string, param5 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.char)(param1)

	cValue2 := (C.GOptionFlags)(param2)

	cValue3 := (C.GOptionArg)(param3)

	cValue4 := (*C.char)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (*C.char)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	C.g_application_add_main_option(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : g_application_add_main_option_entries : has non-string array param entries

func Fn_g_application_add_option_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GOptionGroup)(unsafe.Pointer(param0))

	C.g_application_add_option_group(cValueInstance, cValue0)
}

func Fn_g_application_bind_busy_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_application_bind_busy_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_application_get_application_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_application_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_application_get_dbus_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_dbus_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_application_get_dbus_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_dbus_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_application_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_application_get_inactivity_timeout(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_inactivity_timeout(cValueInstance)

	return (uint)(ret)
}

func Fn_g_application_get_is_busy(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_is_busy(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_application_get_is_registered(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_is_registered(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_application_get_is_remote(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_is_remote(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_application_get_resource_base_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	ret := C.g_application_get_resource_base_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_application_hold(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_hold(cValueInstance)
}

func Fn_g_application_mark_busy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_mark_busy(cValueInstance)
}

// UNSUPPORTED : g_application_open : has non-string array param files

func Fn_g_application_quit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_quit(cValueInstance)
}

func Fn_g_application_register(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_application_register(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_application_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_release(cValueInstance)
}

func Fn_g_application_run(paramInstance unsafe.Pointer, param0 int, param1 []string) int {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	ret := C.g_application_run(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_g_application_send_notification(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GNotification)(unsafe.Pointer(param1))

	C.g_application_send_notification(cValueInstance, cValue0, cValue1)
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

func Fn_g_application_set_resource_base_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_set_resource_base_path(cValueInstance, cValue0)
}

func Fn_g_application_unbind_busy_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_application_unbind_busy_property(cValueInstance, cValue0, cValue1)
}

func Fn_g_application_unmark_busy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_unmark_busy(cValueInstance)
}

func Fn_g_application_withdraw_notification(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_application_withdraw_notification(cValueInstance, cValue0)
}

func Fn_g_application_get_default() unsafe.Pointer {
	ret := C.g_application_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_application_id_is_valid(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_application_id_is_valid(cValue0)

	return toGoBool(ret)
}

func Fn_g_application_command_line_create_file_for_arg(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_application_command_line_create_file_for_arg(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_application_command_line_get_arguments : has array return

func Fn_g_application_command_line_get_cwd(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_cwd(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_application_command_line_get_environ : has array return

func Fn_g_application_command_line_get_exit_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_exit_status(cValueInstance)

	return (int)(ret)
}

func Fn_g_application_command_line_get_is_remote(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_is_remote(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_application_command_line_get_options_dict(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_options_dict(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_application_command_line_get_platform_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_platform_data(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_application_command_line_get_stdin(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	ret := C.g_application_command_line_get_stdin(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_application_command_line_getenv(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_application_command_line_getenv(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_application_command_line_print : has varargs

// UNSUPPORTED : g_application_command_line_printerr : has varargs

func Fn_g_application_command_line_set_exit_status(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GApplicationCommandLine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	C.g_application_command_line_set_exit_status(cValueInstance, cValue0)
}

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	ret := C.g_buffered_input_stream_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (C.gsize)(param1)

	ret := C.g_buffered_input_stream_new_sized(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_input_stream_fill(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gssize)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_fill(cValueInstance, cValue0, cValue1, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_buffered_input_stream_fill_async : has callback

func Fn_g_buffered_input_stream_fill_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_fill_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_get_available(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_buffered_input_stream_get_available(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_get_buffer_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_buffered_input_stream_get_buffer_size(cValueInstance)

	return (uint64)(ret)
}

// UNSUPPORTED : g_buffered_input_stream_peek : has non-string array param buffer

// UNSUPPORTED : g_buffered_input_stream_peek_buffer : has array return

func Fn_g_buffered_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_read_byte(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_buffered_input_stream_set_buffer_size(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gsize)(param0)

	C.g_buffered_input_stream_set_buffer_size(cValueInstance, cValue0)
}

func Fn_g_buffered_output_stream_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	ret := C.g_buffered_output_stream_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_output_stream_new_sized(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	cValue1 := (C.gsize)(param1)

	ret := C.g_buffered_output_stream_new_sized(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_output_stream_get_auto_grow(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_buffered_output_stream_get_auto_grow(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_buffered_output_stream_get_buffer_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GBufferedOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_buffered_output_stream_get_buffer_size(cValueInstance)

	return (uint64)(ret)
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

func Fn_g_bytes_icon_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	ret := C.g_bytes_icon_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_bytes_icon_get_bytes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GBytesIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_bytes_icon_get_bytes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_cancellable_new() unsafe.Pointer {
	ret := C.g_cancellable_new()

	return unsafe.Pointer(ret)
}

func Fn_g_cancellable_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_cancel(cValueInstance)
}

// UNSUPPORTED : g_cancellable_connect : has callback

func Fn_g_cancellable_disconnect(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gulong)(param0)

	C.g_cancellable_disconnect(cValueInstance, cValue0)
}

func Fn_g_cancellable_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	ret := C.g_cancellable_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_cancellable_is_cancelled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	ret := C.g_cancellable_is_cancelled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_cancellable_make_pollfd(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GPollFD)(unsafe.Pointer(param0))

	ret := C.g_cancellable_make_pollfd(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_g_cancellable_set_error_if_cancelled(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_cancellable_set_error_if_cancelled(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_cancellable_source_new(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	ret := C.g_cancellable_source_new(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_cancellable_get_current() unsafe.Pointer {
	ret := C.g_cancellable_get_current()

	return unsafe.Pointer(ret)
}

func Fn_g_charset_converter_new(param0 string, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.g_charset_converter_new(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_charset_converter_get_num_fallbacks(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))

	ret := C.g_charset_converter_get_num_fallbacks(cValueInstance)

	return (uint)(ret)
}

func Fn_g_charset_converter_get_use_fallback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))

	ret := C.g_charset_converter_get_use_fallback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_charset_converter_set_use_fallback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GCharsetConverter)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_charset_converter_set_use_fallback(cValueInstance, cValue0)
}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	ret := C.g_converter_input_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_input_stream_get_converter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GConverterInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_converter_input_stream_get_converter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	ret := C.g_converter_output_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_output_stream_get_converter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GConverterOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_converter_output_stream_get_converter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_credentials_new() unsafe.Pointer {
	ret := C.g_credentials_new()

	return unsafe.Pointer(ret)
}

func Fn_g_credentials_get_native(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GCredentialsType)(param0)

	ret := C.g_credentials_get_native(cValueInstance, cValue0)

	return (unsafe.Pointer)(ret)
}

func Fn_g_credentials_get_unix_pid(paramInstance unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_credentials_get_unix_pid(cValueInstance, cError)

	return (int)(ret)
}

func Fn_g_credentials_get_unix_user(paramInstance unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_credentials_get_unix_user(cValueInstance, cError)

	return (uint)(ret)
}

func Fn_g_credentials_is_same_user(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_credentials_is_same_user(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_set_native(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GCredentialsType)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_credentials_set_native(cValueInstance, cValue0, cValue1)
}

func Fn_g_credentials_set_unix_user(paramInstance unsafe.Pointer, param0 uint, error unsafe.Pointer) bool {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	cValue0 := (C.uid_t)(param0)

	cError := (**C.GError)(error)

	ret := C.g_credentials_set_unix_user(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GCredentials)(unsafe.Pointer(paramInstance))

	ret := C.g_credentials_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_action_group_get(param0 unsafe.Pointer, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_action_group_get(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_auth_observer_new() unsafe.Pointer {
	ret := C.g_dbus_auth_observer_new()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_auth_observer_allow_mechanism(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDBusAuthObserver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_auth_observer_allow_mechanism(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_auth_observer_authorize_authenticated_peer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusAuthObserver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GCredentials)(unsafe.Pointer(param1))

	ret := C.g_dbus_auth_observer_authorize_authenticated_peer(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_sync(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GDBusConnectionFlags)(param1)

	cValue2 := (*C.GDBusAuthObserver)(unsafe.Pointer(param2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_sync(cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_sync(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GDBusConnectionFlags)(param2)

	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_add_filter : has callback

// UNSUPPORTED : g_dbus_connection_call : has callback

func Fn_g_dbus_connection_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_call_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
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

	ret := C.g_dbus_connection_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : has callback

func Fn_g_dbus_connection_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, param5 unsafe.Pointer, param6 int, param7 int, param8 unsafe.Pointer, param9 *unsafe.Pointer, param10 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
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

	ret := C.g_dbus_connection_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_close : has callback

func Fn_g_dbus_connection_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_close_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_sync(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_emit_signal(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer, error unsafe.Pointer) bool {
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

	ret := C.g_dbus_connection_emit_signal(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_export_action_group(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GActionGroup)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_export_action_group(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_g_dbus_connection_export_menu_model(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_export_menu_model(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

// UNSUPPORTED : g_dbus_connection_flush : has callback

func Fn_g_dbus_connection_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_flush_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_sync(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_capabilities(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_capabilities(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_connection_get_exit_on_close(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_exit_on_close(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_guid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_guid(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_get_last_serial(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_last_serial(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_connection_get_peer_credentials(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_peer_credentials(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_unique_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_get_unique_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_connection_is_closed(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_register_object : has callback

func Fn_g_dbus_connection_register_object_with_closures(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) uint {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GDBusInterfaceInfo)(unsafe.Pointer(param1))

	cValue2 := (*C.GClosure)(unsafe.Pointer(param2))

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	cValue4 := (*C.GClosure)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_register_object_with_closures(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return (uint)(ret)
}

// UNSUPPORTED : g_dbus_connection_register_subtree : has callback

func Fn_g_dbus_connection_remove_filter(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_dbus_connection_remove_filter(cValueInstance, cValue0)
}

func Fn_g_dbus_connection_send_message(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *uint32, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusSendMessageFlags)(param1)

	cValue2 := (*C.guint32)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : has callback

func Fn_g_dbus_connection_send_message_with_reply_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_send_message_with_reply_sync(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *uint32, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusMessage)(unsafe.Pointer(param0))

	cValue1 := (C.GDBusSendMessageFlags)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.guint32)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_set_exit_on_close(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_dbus_connection_set_exit_on_close(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_connection_signal_subscribe : has callback

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

func Fn_g_dbus_connection_unregister_object(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_connection_unregister_object(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_unregister_subtree(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GDBusConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_connection_unregister_subtree(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_new : has callback

// UNSUPPORTED : g_dbus_connection_new_for_address : has callback

func Fn_g_dbus_interface_skeleton_export(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_interface_skeleton_export(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_interface_skeleton_flush(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	C.g_dbus_interface_skeleton_flush(cValueInstance)
}

func Fn_g_dbus_interface_skeleton_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_connections(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_connections(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_interface_skeleton_get_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_interface_skeleton_get_properties(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_properties(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_get_vtable(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_skeleton_get_vtable(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_skeleton_has_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusInterfaceSkeleton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	ret := C.g_dbus_interface_skeleton_has_connection(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_g_dbus_menu_model_get(param0 unsafe.Pointer, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_menu_model_get(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new() unsafe.Pointer {
	ret := C.g_dbus_message_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_message_new_from_blob : has non-string array param blob

func Fn_g_dbus_message_new_method_call(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.g_dbus_message_new_method_call(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_signal(param0 string, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_dbus_message_new_signal(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_copy(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_copy(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_arg0(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_arg0(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_body(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_body(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_byte_order(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_message_get_destination(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_destination(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_error_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_error_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_message_get_header(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageHeaderField)(param0)

	ret := C.g_dbus_message_get_header(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : has array return

func Fn_g_dbus_message_get_interface(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_interface(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_locked(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_locked(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_message_get_member(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_member(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_message_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_message_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_message_get_num_unix_fds(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_num_unix_fds(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_reply_serial(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_reply_serial(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_sender(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_sender(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_serial(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_serial(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_signature(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_signature(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_unix_fd_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_unix_fd_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_lock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_lock(cValueInstance)
}

// UNSUPPORTED : g_dbus_message_new_method_error : has varargs

func Fn_g_dbus_message_new_method_error_literal(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_message_new_method_error_literal(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_message_new_method_error_valist : has va_list

func Fn_g_dbus_message_new_method_reply(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_new_method_reply(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_print(paramInstance unsafe.Pointer, param0 uint) string {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_dbus_message_print(cValueInstance, cValue0)

	return C.GoString(ret)
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

// UNSUPPORTED : g_dbus_message_to_blob : has array return

func Fn_g_dbus_message_to_gerror(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_to_gerror(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_message_bytes_needed : has non-string array param blob

func Fn_g_dbus_method_invocation_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_interface_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_interface_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_message(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_message(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_method_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_method_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_parameters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_parameters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_property_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_property_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_sender(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_sender(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_user_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_method_invocation_get_user_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_dbus_method_invocation_return_dbus_error(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_dbus_method_invocation_return_dbus_error(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_dbus_method_invocation_return_error : has varargs

func Fn_g_dbus_method_invocation_return_error_literal(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GDBusMethodInvocation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.g_dbus_method_invocation_return_error_literal(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : g_dbus_method_invocation_return_error_valist : has va_list

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

func Fn_g_dbus_object_manager_client_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_object_manager_client_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_client_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_object_manager_client_new_for_bus_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : has callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : has callback

func Fn_g_dbus_object_manager_client_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_client_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_object_manager_client_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_client_get_name_owner(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManagerClient)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_client_get_name_owner(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_dbus_object_manager_client_new : has callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : has callback

func Fn_g_dbus_object_manager_server_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_server_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_g_dbus_object_manager_server_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_server_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_server_is_exported(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusObjectSkeleton)(unsafe.Pointer(param0))

	ret := C.g_dbus_object_manager_server_is_exported(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_object_manager_server_set_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_set_connection(cValueInstance, cValue0)
}

func Fn_g_dbus_object_manager_server_unexport(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_server_unexport(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_dbus_object_proxy_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_object_proxy_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_proxy_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_proxy_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_skeleton_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_skeleton_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_g_dbus_proxy_new_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_finish(param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_for_bus_finish(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_sync(param0 int, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
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

	ret := C.g_dbus_proxy_new_for_bus_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_sync(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
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

	ret := C.g_dbus_proxy_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call : has callback

func Fn_g_dbus_proxy_call_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_call_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (C.GDBusCallFlags)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : has callback

func Fn_g_dbus_proxy_call_with_unix_fd_list_finish(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GUnixFDList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAsyncResult)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_with_unix_fd_list_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_call_with_unix_fd_list_sync(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 *unsafe.Pointer, param6 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
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

	ret := C.g_dbus_proxy_call_with_unix_fd_list_sync(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_cached_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_proxy_get_cached_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : has array return

func Fn_g_dbus_proxy_get_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_default_timeout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_default_timeout(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_interface_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_interface_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_interface_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_interface_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name_owner(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_name_owner(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_proxy_get_object_path(cValueInstance)

	return C.GoString(ret)
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

// UNSUPPORTED : g_dbus_proxy_new : has callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : has callback

func Fn_g_dbus_server_new_sync(param0 string, param1 int, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GDBusServerFlags)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GDBusAuthObserver)(unsafe.Pointer(param3))

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_dbus_server_new_sync(cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_server_get_client_address(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_client_address(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_server_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_dbus_server_get_guid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_get_guid(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_server_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_server_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_dbus_server_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_start(cValueInstance)
}

func Fn_g_dbus_server_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusServer)(unsafe.Pointer(paramInstance))

	C.g_dbus_server_stop(cValueInstance)
}

func Fn_g_data_input_stream_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	ret := C.g_data_input_stream_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_data_input_stream_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_data_input_stream_get_byte_order(cValueInstance)

	return (int)(ret)
}

func Fn_g_data_input_stream_get_newline_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_data_input_stream_get_newline_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_data_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint8 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_byte(cValueInstance, cValue0, cError)

	return (uint8)(ret)
}

func Fn_g_data_input_stream_read_int16(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int16 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int16(cValueInstance, cValue0, cError)

	return (int16)(ret)
}

func Fn_g_data_input_stream_read_int32(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int32 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int32(cValueInstance, cValue0, cError)

	return (int32)(ret)
}

func Fn_g_data_input_stream_read_int64(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int64 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int64(cValueInstance, cValue0, cError)

	return (int64)(ret)
}

// UNSUPPORTED : g_data_input_stream_read_line : has array return

// UNSUPPORTED : g_data_input_stream_read_line_async : has callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : has array return

func Fn_g_data_input_stream_read_line_finish_utf8(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_line_finish_utf8(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_read_line_utf8(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_line_utf8(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_read_uint16(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint16 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint16(cValueInstance, cValue0, cError)

	return (uint16)(ret)
}

func Fn_g_data_input_stream_read_uint32(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint32 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint32(cValueInstance, cValue0, cError)

	return (uint32)(ret)
}

func Fn_g_data_input_stream_read_uint64(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint64(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_data_input_stream_read_until(paramInstance unsafe.Pointer, param0 string, param1 *uint64, param2 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_until(cValueInstance, cValue0, cValue1, cValue2, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_until_async : has callback

func Fn_g_data_input_stream_read_until_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_until_finish(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_read_upto(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_upto(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_upto_async : has callback

func Fn_g_data_input_stream_read_upto_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) string {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_upto_finish(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
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

func Fn_g_data_output_stream_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	ret := C.g_data_output_stream_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_data_output_stream_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_data_output_stream_get_byte_order(cValueInstance)

	return (int)(ret)
}

func Fn_g_data_output_stream_put_byte(paramInstance unsafe.Pointer, param0 uint8, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guchar)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_byte(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int16(paramInstance unsafe.Pointer, param0 int16, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint16)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int16(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int32(paramInstance unsafe.Pointer, param0 int32, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint32)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int32(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int64(paramInstance unsafe.Pointer, param0 int64, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint64)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int64(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_string(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_string(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint16(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint16(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint32(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint32(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint64(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint64)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint64(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDataStreamByteOrder)(param0)

	C.g_data_output_stream_set_byte_order(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_new_from_filename(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_new_from_filename(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_new_from_keyfile(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	ret := C.g_desktop_app_info_new_from_keyfile(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_get_action_name(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_action_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_boolean(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_boolean(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_desktop_app_info_get_categories(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_categories(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_generic_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_generic_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_is_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_is_hidden(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_desktop_app_info_get_keywords : has array return

func Fn_g_desktop_app_info_get_nodisplay(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_nodisplay(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_desktop_app_info_get_show_in(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_show_in(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_desktop_app_info_get_startup_wm_class(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_startup_wm_class(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_desktop_app_info_get_string_list : has array return

func Fn_g_desktop_app_info_has_key(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_has_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_desktop_app_info_launch_action(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))

	C.g_desktop_app_info_launch_action(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : has callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : has callback

// UNSUPPORTED : g_desktop_app_info_list_actions : has array return

func Fn_g_desktop_app_info_get_implementations(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_get_implementations(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_desktop_app_info_search : has array return

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_set_desktop_env(cValue0)
}

func Fn_g_emblem_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.g_emblem_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_new_with_origin(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GEmblemOrigin)(param1)

	ret := C.g_emblem_new_with_origin(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GEmblem)(unsafe.Pointer(paramInstance))

	ret := C.g_emblem_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_get_origin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GEmblem)(unsafe.Pointer(paramInstance))

	ret := C.g_emblem_get_origin(cValueInstance)

	return (int)(ret)
}

func Fn_g_emblemed_icon_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (*C.GEmblem)(unsafe.Pointer(param1))

	ret := C.g_emblemed_icon_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
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

func Fn_g_emblemed_icon_get_emblems(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_emblemed_icon_get_emblems(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_emblemed_icon_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_emblemed_icon_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_enumerator_close_async : has callback

func Fn_g_file_enumerator_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_get_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	ret := C.g_file_enumerator_get_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_get_container(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	ret := C.g_file_enumerator_get_container(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	ret := C.g_file_enumerator_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	ret := C.g_file_enumerator_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_iterate(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GFileInfo)(unsafe.Pointer(param0))

	cValue1 := (**C.GFile)(unsafe.Pointer(param1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_iterate(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_next_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_next_file(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_enumerator_next_files_async : has callback

func Fn_g_file_enumerator_next_files_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_next_files_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_set_pending(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_file_enumerator_set_pending(cValueInstance, cValue0)
}

func Fn_g_file_io_stream_get_etag(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_file_io_stream_get_etag(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_io_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_io_stream_query_info_async : has callback

func Fn_g_file_io_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_icon_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	ret := C.g_file_icon_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_icon_get_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_file_icon_get_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_new() unsafe.Pointer {
	ret := C.g_file_info_new()

	return unsafe.Pointer(ret)
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

func Fn_g_file_info_dup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_dup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_attribute_as_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_as_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_file_info_get_attribute_boolean(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_boolean(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_info_get_attribute_byte_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_byte_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_file_info_get_attribute_data(paramInstance unsafe.Pointer, param0 string, param1 *int, param2 *unsafe.Pointer, param3 *int) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GFileAttributeType)(unsafe.Pointer(param1))

	cValue2 := (*C.gpointer)(unsafe.Pointer(param2))

	cValue3 := (*C.GFileAttributeStatus)(unsafe.Pointer(param3))

	ret := C.g_file_info_get_attribute_data(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_g_file_info_get_attribute_int32(paramInstance unsafe.Pointer, param0 string) int32 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_int32(cValueInstance, cValue0)

	return (int32)(ret)
}

func Fn_g_file_info_get_attribute_int64(paramInstance unsafe.Pointer, param0 string) int64 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_int64(cValueInstance, cValue0)

	return (int64)(ret)
}

func Fn_g_file_info_get_attribute_object(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_attribute_status(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_status(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_file_info_get_attribute_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_file_info_get_attribute_stringv : has array return

func Fn_g_file_info_get_attribute_type(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_file_info_get_attribute_uint32(paramInstance unsafe.Pointer, param0 string) uint32 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_uint32(cValueInstance, cValue0)

	return (uint32)(ret)
}

func Fn_g_file_info_get_attribute_uint64(paramInstance unsafe.Pointer, param0 string) uint64 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_get_attribute_uint64(cValueInstance, cValue0)

	return (uint64)(ret)
}

func Fn_g_file_info_get_content_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_content_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_get_deletion_date(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_deletion_date(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_get_edit_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_edit_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_get_etag(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_etag(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_get_file_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_file_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_file_info_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_is_backup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_is_backup(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_info_get_is_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_is_hidden(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_info_get_is_symlink(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_is_symlink(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_info_get_modification_time(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTimeVal)(unsafe.Pointer(param0))

	C.g_file_info_get_modification_time(cValueInstance, cValue0)
}

func Fn_g_file_info_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_get_size(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_size(cValueInstance)

	return (int64)(ret)
}

func Fn_g_file_info_get_sort_order(paramInstance unsafe.Pointer) int32 {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_sort_order(cValueInstance)

	return (int32)(ret)
}

func Fn_g_file_info_get_symbolic_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_symbolic_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_symlink_target(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_file_info_get_symlink_target(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_info_has_attribute(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_has_attribute(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_info_has_namespace(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_info_has_namespace(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_info_list_attributes : has array return

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

func Fn_g_file_info_set_attribute_status(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileAttributeStatus)(param1)

	ret := C.g_file_info_set_attribute_status(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
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
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	C.g_file_info_set_attribute_stringv(cValueInstance, cValue0, cValue1)
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

func Fn_g_file_input_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_input_stream_query_info(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_input_stream_query_info_async : has callback

func Fn_g_file_input_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_input_stream_query_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_cancel(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_file_monitor_cancel(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_monitor_emit_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cValue1 := (*C.GFile)(unsafe.Pointer(param1))

	cValue2 := (C.GFileMonitorEvent)(param2)

	C.g_file_monitor_emit_event(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_file_monitor_is_cancelled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_file_monitor_is_cancelled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_monitor_set_rate_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GFileMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_file_monitor_set_rate_limit(cValueInstance, cValue0)
}

func Fn_g_file_output_stream_get_etag(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_file_output_stream_get_etag(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_output_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_output_stream_query_info(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_output_stream_query_info_async : has callback

func Fn_g_file_output_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_output_stream_query_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_filename_completer_new() unsafe.Pointer {
	ret := C.g_filename_completer_new()

	return unsafe.Pointer(ret)
}

func Fn_g_filename_completer_get_completion_suffix(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GFilenameCompleter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_filename_completer_get_completion_suffix(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_filename_completer_get_completions : has array return

func Fn_g_filename_completer_set_dirs_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilenameCompleter)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_filename_completer_set_dirs_only(cValueInstance, cValue0)
}

func Fn_g_filter_input_stream_get_base_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_filter_input_stream_get_base_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_filter_input_stream_get_close_base_stream(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_filter_input_stream_get_close_base_stream(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_filter_input_stream_set_close_base_stream(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilterInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_filter_input_stream_set_close_base_stream(cValueInstance, cValue0)
}

func Fn_g_filter_output_stream_get_base_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_filter_output_stream_get_base_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_filter_output_stream_get_close_base_stream(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_filter_output_stream_get_close_base_stream(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_filter_output_stream_set_close_base_stream(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFilterOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_filter_output_stream_set_close_base_stream(cValueInstance, cValue0)
}

// UNSUPPORTED : g_io_module_new : blacklisted
// UNSUPPORTED : g_io_module_load : blacklisted
// UNSUPPORTED : g_io_module_unload : blacklisted
// UNSUPPORTED : g_io_module_query : blacklisted
func Fn_g_io_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	C.g_io_stream_clear_pending(cValueInstance)
}

func Fn_g_io_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_close_async : has callback

func Fn_g_io_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_io_stream_get_input_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_get_input_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_get_output_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_get_output_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_io_stream_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	ret := C.g_io_stream_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_io_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GIOStream)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_set_pending(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_splice_async : has callback

func Fn_g_io_stream_splice_finish(param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_splice_finish(cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_inet_address_new_any(param0 int) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	ret := C.g_inet_address_new_any(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_inet_address_new_from_bytes : has non-string array param bytes

func Fn_g_inet_address_new_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_inet_address_new_from_string(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_new_loopback(param0 int) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	ret := C.g_inet_address_new_loopback(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	ret := C.g_inet_address_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_inet_address_get_is_any(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_any(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_link_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_link_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_loopback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_loopback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_global(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_global(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_link_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_link_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_node_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_node_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_org_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_org_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_site_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_mc_site_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_multicast(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_multicast(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_site_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_is_site_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_native_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_get_native_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_inet_address_to_bytes(paramInstance unsafe.Pointer) *uint8 {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_to_bytes(cValueInstance)

	return (*uint8)(ret)
}

func Fn_g_inet_address_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GInetAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_inet_address_mask_new(param0 unsafe.Pointer, param1 uint, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cError := (**C.GError)(error)

	ret := C.g_inet_address_mask_new(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_new_from_string(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_inet_address_mask_new_from_string(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddressMask)(unsafe.Pointer(param0))

	ret := C.g_inet_address_mask_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_inet_address_mask_get_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_mask_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_inet_address_mask_get_length(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_get_length(cValueInstance)

	return (uint)(ret)
}

func Fn_g_inet_address_mask_matches(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	ret := C.g_inet_address_mask_matches(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_inet_address_mask_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GInetAddressMask)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_address_mask_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_inet_socket_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_new_from_string(param0 string, param1 uint) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	ret := C.g_inet_socket_address_new_from_string(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_flowinfo(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_flowinfo(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_inet_socket_address_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_inet_socket_address_get_scope_id(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_scope_id(cValueInstance)

	return (uint32)(ret)
}

func Fn_g_input_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_clear_pending(cValueInstance)
}

func Fn_g_input_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_input_stream_close_async : has callback

func Fn_g_input_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_input_stream_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_input_stream_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_input_stream_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_input_stream_is_closed(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_input_stream_read : has non-string array param buffer

// UNSUPPORTED : g_input_stream_read_all : has non-string array param buffer

// UNSUPPORTED : g_input_stream_read_all_async : has callback

func Fn_g_input_stream_read_all_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_read_all_finish(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_input_stream_read_async : has callback

func Fn_g_input_stream_read_bytes(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gsize)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_read_bytes(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_input_stream_read_bytes_async : has callback

func Fn_g_input_stream_read_bytes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_read_bytes_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_input_stream_read_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_read_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_input_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_set_pending(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_input_stream_skip(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gsize)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_skip(cValueInstance, cValue0, cValue1, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_input_stream_skip_async : has callback

func Fn_g_input_stream_skip_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_skip_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_list_store_new(param0 uint64) unsafe.Pointer {
	cValue0 := (C.GType)(param0)

	ret := C.g_list_store_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_list_store_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.g_list_store_append(cValueInstance, cValue0)
}

func Fn_g_list_store_insert(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.gpointer)(param1)

	C.g_list_store_insert(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_list_store_insert_sorted : has callback

func Fn_g_list_store_remove(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.g_list_store_remove(cValueInstance, cValue0)
}

func Fn_g_list_store_remove_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GListStore)(unsafe.Pointer(paramInstance))

	C.g_list_store_remove_all(cValueInstance)
}

// UNSUPPORTED : g_list_store_sort : has callback

// UNSUPPORTED : g_list_store_splice : has non-string array param additions

func Fn_g_memory_input_stream_new() unsafe.Pointer {
	ret := C.g_memory_input_stream_new()

	return unsafe.Pointer(ret)
}

func Fn_g_memory_input_stream_new_from_bytes(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	ret := C.g_memory_input_stream_new_from_bytes(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_memory_input_stream_new_from_data : has callback

func Fn_g_memory_input_stream_add_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GMemoryInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	C.g_memory_input_stream_add_bytes(cValueInstance, cValue0)
}

// UNSUPPORTED : g_memory_input_stream_add_data : has callback

// UNSUPPORTED : g_memory_output_stream_new : has callback

func Fn_g_memory_output_stream_new_resizable() unsafe.Pointer {
	ret := C.g_memory_output_stream_new_resizable()

	return unsafe.Pointer(ret)
}

func Fn_g_memory_output_stream_get_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_get_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_memory_output_stream_get_data_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_get_data_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_memory_output_stream_get_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_get_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_memory_output_stream_steal_as_bytes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_steal_as_bytes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_memory_output_stream_steal_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_steal_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_menu_new() unsafe.Pointer {
	ret := C.g_menu_new()

	return unsafe.Pointer(ret)
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

func Fn_g_menu_attribute_iter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_menu_attribute_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (**C.GVariant)(unsafe.Pointer(param1))

	ret := C.g_menu_attribute_iter_get_next(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_g_menu_attribute_iter_get_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_get_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_attribute_iter_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuAttributeIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_attribute_iter_next(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_menu_item_new(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_menu_item_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_new_from_model(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GMenuModel)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.g_menu_item_new_from_model(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_new_section(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_item_new_section(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_new_submenu(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_item_new_submenu(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_menu_item_get_attribute : has varargs

func Fn_g_menu_item_get_attribute_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	ret := C.g_menu_item_get_attribute_value(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_item_get_link(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_menu_item_get_link(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_menu_item_set_action_and_target : has varargs

func Fn_g_menu_item_set_action_and_target_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_menu_item_set_action_and_target_value(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_menu_item_set_attribute : has varargs

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

func Fn_g_menu_link_iter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_menu_link_iter_get_next(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (**C.GMenuModel)(unsafe.Pointer(param1))

	ret := C.g_menu_link_iter_get_next(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
}

func Fn_g_menu_link_iter_get_value(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_get_value(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_link_iter_next(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuLinkIter)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_link_iter_next(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_menu_model_get_item_attribute : has varargs

func Fn_g_menu_model_get_item_attribute_value(paramInstance unsafe.Pointer, param0 int, param1 string, param2 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GVariantType)(unsafe.Pointer(param2))

	ret := C.g_menu_model_get_item_attribute_value(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_get_item_link(paramInstance unsafe.Pointer, param0 int, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_menu_model_get_item_link(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_get_n_items(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_model_get_n_items(cValueInstance)

	return (int)(ret)
}

func Fn_g_menu_model_is_mutable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	ret := C.g_menu_model_is_mutable(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_menu_model_items_changed(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int) {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.g_menu_model_items_changed(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_menu_model_iterate_item_attributes(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.g_menu_model_iterate_item_attributes(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_menu_model_iterate_item_links(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GMenuModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.g_menu_model_iterate_item_links(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_operation_new() unsafe.Pointer {
	ret := C.g_mount_operation_new()

	return unsafe.Pointer(ret)
}

func Fn_g_mount_operation_get_anonymous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_anonymous(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_mount_operation_get_choice(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_choice(cValueInstance)

	return (int)(ret)
}

func Fn_g_mount_operation_get_domain(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_domain(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_mount_operation_get_password(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_password(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_mount_operation_get_password_save(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_password_save(cValueInstance)

	return (int)(ret)
}

func Fn_g_mount_operation_get_username(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_operation_get_username(cValueInstance)

	return C.GoString(ret)
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

func Fn_g_native_socket_address_new(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gsize)(param1)

	ret := C.g_native_socket_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_new(param0 string, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_network_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_new_loopback(param0 uint16) unsafe.Pointer {
	cValue0 := (C.guint16)(param0)

	ret := C.g_network_address_new_loopback(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_get_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_address_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_network_address_get_scheme(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_network_address_get_scheme(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_address_parse(param0 string, param1 uint16, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_parse_uri(param0 string, param1 uint16, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse_uri(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_new(param0 string, param1 string, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_network_service_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_get_domain(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_domain(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_get_protocol(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_protocol(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_get_scheme(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_scheme(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_get_service(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_service(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_network_service_set_scheme(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_network_service_set_scheme(cValueInstance, cValue0)
}

func Fn_g_notification_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_notification_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_notification_add_button(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_notification_add_button(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_notification_add_button_with_target : has varargs

func Fn_g_notification_add_button_with_target_value(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_notification_add_button_with_target_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_notification_set_body(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_body(cValueInstance, cValue0)
}

func Fn_g_notification_set_default_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_default_action(cValueInstance, cValue0)
}

// UNSUPPORTED : g_notification_set_default_action_and_target : has varargs

func Fn_g_notification_set_default_action_and_target_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_notification_set_default_action_and_target_value(cValueInstance, cValue0, cValue1)
}

func Fn_g_notification_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.g_notification_set_icon(cValueInstance, cValue0)
}

func Fn_g_notification_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GNotificationPriority)(param0)

	C.g_notification_set_priority(cValueInstance, cValue0)
}

func Fn_g_notification_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_notification_set_title(cValueInstance, cValue0)
}

func Fn_g_notification_set_urgent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_notification_set_urgent(cValueInstance, cValue0)
}

func Fn_g_output_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_clear_pending(cValueInstance)
}

func Fn_g_output_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_close_async : has callback

func Fn_g_output_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_flush(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_flush(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_flush_async : has callback

func Fn_g_output_stream_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_flush_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_has_pending(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_output_stream_has_pending(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_output_stream_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_output_stream_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_output_stream_is_closing(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_output_stream_is_closing(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_printf : has varargs

func Fn_g_output_stream_set_pending(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_set_pending(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_splice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (C.GOutputStreamSpliceFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_splice(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_splice_async : has callback

func Fn_g_output_stream_splice_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_splice_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_vprintf : has va_list

// UNSUPPORTED : g_output_stream_write : has non-string array param buffer

// UNSUPPORTED : g_output_stream_write_all : has non-string array param buffer

// UNSUPPORTED : g_output_stream_write_all_async : has callback

func Fn_g_output_stream_write_all_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_all_finish(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_write_async : has callback

func Fn_g_output_stream_write_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_bytes(cValueInstance, cValue0, cValue1, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_write_bytes_async : has callback

func Fn_g_output_stream_write_bytes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_bytes_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_output_stream_write_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_writev : has non-string array param vectors

// UNSUPPORTED : g_output_stream_writev_all : has non-string array param vectors

// UNSUPPORTED : g_output_stream_writev_all_async : has callback

// UNSUPPORTED : g_output_stream_writev_async : has callback

func Fn_g_permission_acquire(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_acquire_async : has callback

func Fn_g_permission_acquire_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_permission_get_allowed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_allowed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_acquire(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_can_acquire(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_release(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	ret := C.g_permission_get_can_release(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_permission_impl_update(paramInstance unsafe.Pointer, param0 bool, param1 bool, param2 bool) {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	C.g_permission_impl_update(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_permission_release(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_release(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_release_async : has callback

func Fn_g_permission_release_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GPermission)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_permission_release_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_property_action_new(param0 string, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_property_action_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_address_new(param0 unsafe.Pointer, param1 uint16, param2 string, param3 string, param4 uint16, param5 string, param6 string) unsafe.Pointer {
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

	ret := C.g_proxy_address_new(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_address_get_destination_hostname(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_destination_hostname(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_destination_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_destination_port(cValueInstance)

	return (uint16)(ret)
}

func Fn_g_proxy_address_get_destination_protocol(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_destination_protocol(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_password(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_password(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_protocol(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_protocol(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_username(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GProxyAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_address_get_username(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_resolver_lookup_by_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address(cValueInstance, cValue0, cValue1, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_address_async : has callback

func Fn_g_resolver_lookup_by_address_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) string {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address_finish(cValueInstance, cValue0, cError)

	return C.GoString(ret)
}

func Fn_g_resolver_lookup_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_async : has callback

func Fn_g_resolver_lookup_by_name_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : has callback

func Fn_g_resolver_lookup_records(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GResolverRecordType)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_records(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_records_async : has callback

func Fn_g_resolver_lookup_records_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_records_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resolver_lookup_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_service_async : has callback

func Fn_g_resolver_lookup_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GResolver)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
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

func Fn_g_resolver_get_default() unsafe.Pointer {
	ret := C.g_resolver_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.GSettingsSchema)(unsafe.Pointer(param0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_settings_new_full(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	ret := C.g_settings_new_with_backend(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend_and_path(param0 string, param1 unsafe.Pointer, param2 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GSettingsBackend)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.g_settings_new_with_backend_and_path(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_path(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_settings_new_with_path(cValue0, cValue1)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : g_settings_bind_with_mapping : has callback

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

func Fn_g_settings_create_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_create_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_delay(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_delay(cValueInstance)
}

// UNSUPPORTED : g_settings_get : has varargs

func Fn_g_settings_get_boolean(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_boolean(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_settings_get_child(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_default_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_default_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_double(paramInstance unsafe.Pointer, param0 string) float64 {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_double(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_g_settings_get_enum(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_enum(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_settings_get_flags(paramInstance unsafe.Pointer, param0 string) uint {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_flags(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_g_settings_get_has_unapplied(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_get_has_unapplied(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_settings_get_int(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_int(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : g_settings_get_mapped : has callback

func Fn_g_settings_get_range(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_range(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_get_strv : has array return

func Fn_g_settings_get_uint(paramInstance unsafe.Pointer, param0 string) uint {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_uint(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_g_settings_get_user_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_user_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_value(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_get_value(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_is_writable(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_settings_is_writable(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_children : has array return

// UNSUPPORTED : g_settings_list_keys : has array return

func Fn_g_settings_range_check(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	ret := C.g_settings_range_check(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
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

// UNSUPPORTED : g_settings_set : has varargs

func Fn_g_settings_set_boolean(paramInstance unsafe.Pointer, param0 string, param1 bool) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.g_settings_set_boolean(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	ret := C.g_settings_set_double(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_enum(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.g_settings_set_enum(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_flags(paramInstance unsafe.Pointer, param0 string, param1 uint) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	ret := C.g_settings_set_flags(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.g_settings_set_int(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_string(paramInstance unsafe.Pointer, param0 string, param1 string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_settings_set_string(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_strv(paramInstance unsafe.Pointer, param0 string, param1 []string) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	ret := C.g_settings_set_strv(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_uint(paramInstance unsafe.Pointer, param0 string, param1 uint) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	ret := C.g_settings_set_uint(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_settings_set_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	ret := C.g_settings_set_value(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_relocatable_schemas : has array return

// UNSUPPORTED : g_settings_list_schemas : has array return

func Fn_g_settings_sync() {
	C.g_settings_sync()
}

func Fn_g_settings_unbind(param0 unsafe.Pointer, param1 string) {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_settings_unbind(cValue0, cValue1)
}

// UNSUPPORTED : g_settings_backend_changed : blacklisted
// UNSUPPORTED : g_settings_backend_changed_tree : blacklisted
// UNSUPPORTED : g_settings_backend_keys_changed : blacklisted
// UNSUPPORTED : g_settings_backend_path_changed : blacklisted
// UNSUPPORTED : g_settings_backend_path_writable_changed : blacklisted
// UNSUPPORTED : g_settings_backend_writable_changed : blacklisted
// UNSUPPORTED : g_settings_backend_flatten_tree : blacklisted
// UNSUPPORTED : g_settings_backend_get_default : blacklisted
func Fn_g_simple_action_new(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	ret := C.g_simple_action_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_simple_action_new_stateful(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariantType)(unsafe.Pointer(param1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	ret := C.g_simple_action_new_stateful(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
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

func Fn_g_simple_action_set_state_hint(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_simple_action_set_state_hint(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_new() unsafe.Pointer {
	ret := C.g_simple_action_group_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_simple_action_group_add_entries : has non-string array param entries

func Fn_g_simple_action_group_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAction)(unsafe.Pointer(param0))

	C.g_simple_action_group_insert(cValueInstance, cValue0)
}

func Fn_g_simple_action_group_lookup(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_simple_action_group_lookup(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_simple_action_group_remove(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSimpleActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_simple_action_group_remove(cValueInstance, cValue0)
}

// UNSUPPORTED : g_simple_async_result_new : has callback

// UNSUPPORTED : g_simple_async_result_new_error : has varargs

// UNSUPPORTED : g_simple_async_result_new_from_error : has callback

// UNSUPPORTED : g_simple_async_result_new_take_error : has callback

func Fn_g_simple_async_result_complete(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_complete(cValueInstance)
}

func Fn_g_simple_async_result_complete_in_idle(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_complete_in_idle(cValueInstance)
}

func Fn_g_simple_async_result_get_op_res_gboolean(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_simple_async_result_get_op_res_gboolean(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_simple_async_result_get_op_res_gpointer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_simple_async_result_get_op_res_gpointer(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_simple_async_result_get_op_res_gssize(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_simple_async_result_get_op_res_gssize(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_simple_async_result_get_source_tag(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_simple_async_result_get_source_tag(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_simple_async_result_propagate_error(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_simple_async_result_propagate_error(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_simple_async_result_run_in_thread : has callback

func Fn_g_simple_async_result_set_check_cancellable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_simple_async_result_set_check_cancellable(cValueInstance, cValue0)
}

// UNSUPPORTED : g_simple_async_result_set_error : has varargs

// UNSUPPORTED : g_simple_async_result_set_error_va : has va_list

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

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : has callback

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

func Fn_g_simple_async_result_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (C.gpointer)(param2)

	ret := C.g_simple_async_result_is_valid(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_g_simple_io_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GOutputStream)(unsafe.Pointer(param1))

	ret := C.g_simple_io_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_simple_permission_new(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.g_simple_permission_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_g_simple_proxy_resolver_new(param0 string, param1 *string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	var cValue1String *C.gchar
	cValue1String = (*C.gchar)(C.CString(*param1))
	cValue1 := &cValue1String
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_simple_proxy_resolver_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_new(param0 int, param1 int, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GSocketFamily)(param0)

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.GSocketProtocol)(param2)

	cError := (**C.GError)(error)

	ret := C.g_socket_new(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_new_from_fd(param0 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_socket_new_from_fd(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_accept(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_accept(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_bind(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_bind(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_check_connect_result(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_check_connect_result(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_socket_close(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_close(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_socket_condition_check(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	ret := C.g_socket_condition_check(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_g_socket_condition_timed_wait(paramInstance unsafe.Pointer, param0 int, param1 int64, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (C.gint64)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_condition_timed_wait(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_condition_wait(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_condition_wait(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_connect(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connection_factory_create_connection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_factory_create_connection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_create_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GIOCondition)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	ret := C.g_socket_create_source(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_available_bytes(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_available_bytes(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_get_blocking(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_blocking(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_broadcast(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_broadcast(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_credentials(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_credentials(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_keepalive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_keepalive(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_listen_backlog(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_listen_backlog(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_local_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_multicast_loopback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_multicast_loopback(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_get_multicast_ttl(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_multicast_ttl(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_get_option(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_option(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_get_protocol(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_protocol(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_remote_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_socket_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_socket_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_get_timeout(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_timeout(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_get_ttl(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_ttl(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_is_closed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_is_closed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_is_connected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_is_connected(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_join_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cError := (**C.GError)(error)

	ret := C.g_socket_join_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_leave_multicast_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cError := (**C.GError)(error)

	ret := C.g_socket_leave_multicast_group(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listen(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_listen(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_socket_receive : has non-string array param buffer

// UNSUPPORTED : g_socket_receive_from : has non-string array param buffer

// UNSUPPORTED : g_socket_receive_message : has non-string array param vectors

// UNSUPPORTED : g_socket_receive_messages : has non-string array param messages

// UNSUPPORTED : g_socket_receive_with_blocking : has non-string array param buffer

// UNSUPPORTED : g_socket_send : has non-string array param buffer

// UNSUPPORTED : g_socket_send_message : has non-string array param vectors

// UNSUPPORTED : g_socket_send_message_with_timeout : has non-string array param vectors

// UNSUPPORTED : g_socket_send_messages : has non-string array param messages

// UNSUPPORTED : g_socket_send_to : has non-string array param buffer

// UNSUPPORTED : g_socket_send_with_blocking : has non-string array param buffer

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

func Fn_g_socket_set_option(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cError := (**C.GError)(error)

	ret := C.g_socket_set_option(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
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

func Fn_g_socket_shutdown(paramInstance unsafe.Pointer, param0 bool, param1 bool, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_shutdown(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_speaks_ipv4(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_speaks_ipv4(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_address_new_from_native(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gsize)(param1)

	ret := C.g_socket_address_new_from_native(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_address_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_address_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_address_get_native_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_address_get_native_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_address_to_native(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketAddress)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.g_socket_address_to_native(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_address_enumerator_next(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : has callback

func Fn_g_socket_address_enumerator_next_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_new() unsafe.Pointer {
	ret := C.g_socket_client_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_add_application_proxy(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_socket_client_add_application_proxy(cValueInstance, cValue0)
}

func Fn_g_socket_client_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_async : has callback

func Fn_g_socket_client_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_host(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_host_async : has callback

func Fn_g_socket_client_connect_to_host_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_service(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_service_async : has callback

func Fn_g_socket_client_connect_to_service_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_service_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_uri(paramInstance unsafe.Pointer, param0 string, param1 uint16, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_uri_async : has callback

func Fn_g_socket_client_connect_to_uri_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_enable_proxy(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_enable_proxy(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_client_get_family(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_family(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_get_local_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_local_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_protocol(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_protocol(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_get_proxy_resolver(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_proxy_resolver(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_socket_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_socket_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_client_get_timeout(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_timeout(cValueInstance)

	return (uint)(ret)
}

func Fn_g_socket_client_get_tls(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_tls(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_client_get_tls_validation_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_tls_validation_flags(cValueInstance)

	return (int)(ret)
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

func Fn_g_socket_connection_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_connect(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_socket_connection_connect_async : has callback

func Fn_g_socket_connection_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_connect_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connection_get_local_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_local_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_remote_address(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_remote_address(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_socket(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_get_socket(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_is_connected(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connection_is_connected(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_connection_factory_lookup_type(param0 int, param1 int, param2 int) uint64 {
	cValue0 := (C.GSocketFamily)(param0)

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.gint)(param2)

	ret := C.g_socket_connection_factory_lookup_type(cValue0, cValue1, cValue2)

	return (uint64)(ret)
}

func Fn_g_socket_connection_factory_register_type(param0 uint64, param1 int, param2 int, param3 int) {
	cValue0 := (C.GType)(param0)

	cValue1 := (C.GSocketFamily)(param1)

	cValue2 := (C.GSocketType)(param2)

	cValue3 := (C.gint)(param3)

	C.g_socket_connection_factory_register_type(cValue0, cValue1, cValue2, cValue3)
}

func Fn_g_socket_control_message_get_level(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_level(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_msg_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_msg_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_control_message_get_size(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_socket_control_message_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketControlMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.g_socket_control_message_serialize(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_control_message_deserialize : has non-string array param data

func Fn_g_socket_listener_new() unsafe.Pointer {
	ret := C.g_socket_listener_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_accept(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_listener_accept_async : has callback

func Fn_g_socket_listener_accept_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_accept_socket(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_socket(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_listener_accept_socket_async : has callback

func Fn_g_socket_listener_accept_socket_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_accept_socket_finish(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_listener_add_address(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketAddress)(unsafe.Pointer(param0))

	cValue1 := (C.GSocketType)(param1)

	cValue2 := (C.GSocketProtocol)(param2)

	cValue3 := (*C.GObject)(unsafe.Pointer(param3))

	cValue4 := (**C.GSocketAddress)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_address(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_add_any_inet_port(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint16 {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_any_inet_port(cValueInstance, cValue0, cError)

	return (uint16)(ret)
}

func Fn_g_socket_listener_add_inet_port(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_inet_port(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_add_socket(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketListener)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocket)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_socket(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
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

func Fn_g_socket_service_new() unsafe.Pointer {
	ret := C.g_socket_service_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_service_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_service_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_socket_service_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_start(cValueInstance)
}

func Fn_g_socket_service_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSocketService)(unsafe.Pointer(paramInstance))

	C.g_socket_service_stop(cValueInstance)
}

// UNSUPPORTED : g_subprocess_new : has varargs

func Fn_g_subprocess_newv(param0 []string, param1 int, error unsafe.Pointer) unsafe.Pointer {
	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cValue1 := (C.GSubprocessFlags)(param1)

	cError := (**C.GError)(error)

	ret := C.g_subprocess_newv(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_communicate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer, param3 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cValue2 := (**C.GBytes)(unsafe.Pointer(param2))

	cValue3 := (**C.GBytes)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_communicate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_communicate_async : has callback

func Fn_g_subprocess_communicate_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GBytes)(unsafe.Pointer(param1))

	cValue2 := (**C.GBytes)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_communicate_finish(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_subprocess_communicate_utf8(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 *string, param3 *string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	var cValue2String *C.gchar
	cValue2 := &cValue2String

	var cValue3String *C.gchar
	cValue3 := &cValue3String

	cError := (**C.GError)(error)

	ret := C.g_subprocess_communicate_utf8(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	param2String := C.GoString(cValue2String)
	*param2 = param2String

	param3String := C.GoString(cValue3String)
	*param3 = param3String

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_communicate_utf8_async : has callback

func Fn_g_subprocess_communicate_utf8_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *string, param2 *string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	var cValue2String *C.gchar
	cValue2 := &cValue2String

	cError := (**C.GError)(error)

	ret := C.g_subprocess_communicate_utf8_finish(cValueInstance, cValue0, cValue1, cValue2, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	param2String := C.GoString(cValue2String)
	*param2 = param2String

	return toGoBool(ret)
}

func Fn_g_subprocess_force_exit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	C.g_subprocess_force_exit(cValueInstance)
}

func Fn_g_subprocess_get_exit_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_exit_status(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_get_identifier(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_identifier(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_subprocess_get_if_exited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_if_exited(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_if_signaled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_if_signaled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_status(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_get_stderr_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stderr_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_stdin_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stdin_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_stdout_pipe(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_stdout_pipe(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_get_successful(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_successful(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_subprocess_get_term_sig(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	ret := C.g_subprocess_get_term_sig(cValueInstance)

	return (int)(ret)
}

func Fn_g_subprocess_send_signal(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_send_signal(cValueInstance, cValue0)
}

func Fn_g_subprocess_wait(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_wait_async : has callback

func Fn_g_subprocess_wait_check(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_check(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_subprocess_wait_check_async : has callback

func Fn_g_subprocess_wait_check_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_check_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_subprocess_wait_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_subprocess_wait_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_subprocess_launcher_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GSubprocessFlags)(param0)

	ret := C.g_subprocess_launcher_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_launcher_getenv(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_subprocess_launcher_getenv(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : has callback

func Fn_g_subprocess_launcher_set_cwd(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_cwd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_environ(paramInstance unsafe.Pointer, param0 []string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.g_subprocess_launcher_set_environ(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSubprocessFlags)(param0)

	C.g_subprocess_launcher_set_flags(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stderr_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stderr_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stdin_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stdin_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_set_stdout_file_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_set_stdout_file_path(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_setenv(paramInstance unsafe.Pointer, param0 string, param1 string, param2 bool) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := toCBool(param2)

	C.g_subprocess_launcher_setenv(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : g_subprocess_launcher_spawn : has varargs

func Fn_g_subprocess_launcher_spawnv(paramInstance unsafe.Pointer, param0 []string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cError := (**C.GError)(error)

	ret := C.g_subprocess_launcher_spawnv(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_subprocess_launcher_take_fd(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.g_subprocess_launcher_take_fd(cValueInstance, cValue0, cValue1)
}

func Fn_g_subprocess_launcher_take_stderr_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stderr_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_take_stdin_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stdin_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_take_stdout_fd(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.g_subprocess_launcher_take_stdout_fd(cValueInstance, cValue0)
}

func Fn_g_subprocess_launcher_unsetenv(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_subprocess_launcher_unsetenv(cValueInstance, cValue0)
}

// UNSUPPORTED : g_task_new : has callback

// UNSUPPORTED : g_task_attach_source : has callback

func Fn_g_task_get_cancellable(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_cancellable(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_task_get_check_cancellable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_check_cancellable(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_task_get_completed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_completed(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_task_get_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_task_get_priority(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_priority(cValueInstance)

	return (int)(ret)
}

func Fn_g_task_get_return_on_cancel(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_return_on_cancel(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_task_get_source_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_source_object(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_task_get_source_tag(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_source_tag(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_task_get_task_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_get_task_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_task_had_error(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_had_error(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_task_propagate_boolean(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_task_propagate_boolean(cValueInstance, cError)

	return toGoBool(ret)
}

func Fn_g_task_propagate_int(paramInstance unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_task_propagate_int(cValueInstance, cError)

	return (uint64)(ret)
}

func Fn_g_task_propagate_pointer(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_task_propagate_pointer(cValueInstance, cError)

	return (unsafe.Pointer)(ret)
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

func Fn_g_task_return_error_if_cancelled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	ret := C.g_task_return_error_if_cancelled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_task_return_int(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gssize)(param0)

	C.g_task_return_int(cValueInstance, cValue0)
}

// UNSUPPORTED : g_task_return_new_error : has varargs

// UNSUPPORTED : g_task_return_pointer : has callback

// UNSUPPORTED : g_task_run_in_thread : has callback

// UNSUPPORTED : g_task_run_in_thread_sync : has callback

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

func Fn_g_task_set_return_on_cancel(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.g_task_set_return_on_cancel(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_task_set_source_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTask)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.g_task_set_source_tag(cValueInstance, cValue0)
}

// UNSUPPORTED : g_task_set_task_data : has callback

func Fn_g_task_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (C.gpointer)(param0)

	cValue1 := (C.gpointer)(param1)

	ret := C.g_task_is_valid(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : g_task_report_error : has callback

// UNSUPPORTED : g_task_report_new_error : has varargs

func Fn_g_tcp_connection_get_graceful_disconnect(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tcp_connection_get_graceful_disconnect(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tcp_connection_set_graceful_disconnect(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTcpConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_tcp_connection_set_graceful_disconnect(cValueInstance, cValue0)
}

func Fn_g_tcp_wrapper_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GSocket)(unsafe.Pointer(param1))

	ret := C.g_tcp_wrapper_connection_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_tcp_wrapper_connection_get_base_io_stream(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTcpWrapperConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tcp_wrapper_connection_get_base_io_stream(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_test_dbus_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GTestDBusFlags)(param0)

	ret := C.g_test_dbus_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_g_test_dbus_get_bus_address(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	ret := C.g_test_dbus_get_bus_address(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_test_dbus_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTestDBus)(unsafe.Pointer(paramInstance))

	ret := C.g_test_dbus_get_flags(cValueInstance)

	return (int)(ret)
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

func Fn_g_themed_icon_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_themed_icon_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) unsafe.Pointer {
	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cValue1 := (C.int)(param1)

	ret := C.g_themed_icon_new_from_names(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_themed_icon_new_with_default_fallbacks(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_append_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_append_name(cValueInstance, cValue0)
}

// UNSUPPORTED : g_themed_icon_get_names : has array return

func Fn_g_themed_icon_prepend_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_themed_icon_prepend_name(cValueInstance, cValue0)
}

func Fn_g_threaded_socket_service_new(param0 int) unsafe.Pointer {
	cValue0 := (C.int)(param0)

	ret := C.g_threaded_socket_service_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_certificate_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_tls_certificate_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_certificate_new_from_files(param0 string, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.g_tls_certificate_new_from_files(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_certificate_new_from_pem(param0 string, param1 uint64, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cError := (**C.GError)(error)

	ret := C.g_tls_certificate_new_from_pem(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_certificate_get_issuer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_certificate_get_issuer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_certificate_is_same(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	ret := C.g_tls_certificate_is_same(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_tls_certificate_verify(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) int {
	cValueInstance := (*C.GTlsCertificate)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	cValue1 := (*C.GTlsCertificate)(unsafe.Pointer(param1))

	ret := C.g_tls_certificate_verify(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_g_tls_certificate_list_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_tls_certificate_list_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_emit_accept_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	cValue1 := (C.GTlsCertificateFlags)(param1)

	ret := C.g_tls_connection_emit_accept_certificate(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_g_tls_connection_get_certificate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_certificate(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_database(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_database(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_interaction(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_interaction(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_peer_certificate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_peer_certificate(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_peer_certificate_errors(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_peer_certificate_errors(cValueInstance)

	return (int)(ret)
}

func Fn_g_tls_connection_get_rehandshake_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_rehandshake_mode(cValueInstance)

	return (int)(ret)
}

func Fn_g_tls_connection_get_require_close_notify(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_require_close_notify(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tls_connection_get_use_system_certdb(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_use_system_certdb(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tls_connection_handshake(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_connection_handshake(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_tls_connection_handshake_async : has callback

func Fn_g_tls_connection_handshake_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_connection_handshake_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
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

func Fn_g_tls_database_create_certificate_handle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	ret := C.g_tls_database_create_certificate_handle(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_tls_database_lookup_certificate_for_handle(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))

	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_for_handle(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : has callback

func Fn_g_tls_database_lookup_certificate_for_handle_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_for_handle_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_database_lookup_certificate_issuer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	cValue1 := (*C.GTlsInteraction)(unsafe.Pointer(param1))

	cValue2 := (C.GTlsDatabaseLookupFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_issuer(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : has callback

func Fn_g_tls_database_lookup_certificate_issuer_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificate_issuer_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : has non-string array param issuer_raw_dn

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : has callback

func Fn_g_tls_database_lookup_certificates_issued_by_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_lookup_certificates_issued_by_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_database_verify_chain(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int, param5 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsCertificate)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GSocketConnectable)(unsafe.Pointer(param2))

	cValue3 := (*C.GTlsInteraction)(unsafe.Pointer(param3))

	cValue4 := (C.GTlsDatabaseVerifyFlags)(param4)

	cValue5 := (*C.GCancellable)(unsafe.Pointer(param5))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_verify_chain(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_database_verify_chain_async : has callback

func Fn_g_tls_database_verify_chain_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsDatabase)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_database_verify_chain_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_ask_password(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_interaction_ask_password_async : has callback

func Fn_g_tls_interaction_ask_password_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_ask_password_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_invoke_ask_password(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsPassword)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_invoke_ask_password(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_invoke_request_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsConnection)(unsafe.Pointer(param0))

	cValue1 := (C.GTlsCertificateRequestFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_invoke_request_certificate(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (int)(ret)
}

func Fn_g_tls_interaction_request_certificate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsConnection)(unsafe.Pointer(param0))

	cValue1 := (C.GTlsCertificateRequestFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_request_certificate(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_tls_interaction_request_certificate_async : has callback

func Fn_g_tls_interaction_request_certificate_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GTlsInteraction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_tls_interaction_request_certificate_finish(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_tls_password_new(param0 int, param1 string) unsafe.Pointer {
	cValue0 := (C.GTlsPasswordFlags)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_tls_password_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_password_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_tls_password_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_tls_password_get_value(paramInstance unsafe.Pointer, param0 *uint64) *uint8 {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.g_tls_password_get_value(cValueInstance, cValue0)

	return (*uint8)(ret)
}

func Fn_g_tls_password_get_warning(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_password_get_warning(cValueInstance)

	return C.GoString(ret)
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

// UNSUPPORTED : g_tls_password_set_value : has non-string array param value

// UNSUPPORTED : g_tls_password_set_value_full : has callback

func Fn_g_tls_password_set_warning(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GTlsPassword)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_tls_password_set_warning(cValueInstance, cValue0)
}

func Fn_g_unix_connection_receive_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_credentials(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_connection_receive_credentials_async : has callback

func Fn_g_unix_connection_receive_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_credentials_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_connection_receive_fd(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_fd(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_unix_connection_send_credentials(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_credentials(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : has callback

func Fn_g_unix_connection_send_credentials_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_credentials_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_unix_connection_send_fd(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_fd(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_unix_credentials_message_new() unsafe.Pointer {
	ret := C.g_unix_credentials_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_new_with_credentials(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GCredentials)(unsafe.Pointer(param0))

	ret := C.g_unix_credentials_message_new_with_credentials(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_get_credentials(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixCredentialsMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_credentials_message_get_credentials(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_is_supported() bool {
	ret := C.g_unix_credentials_message_is_supported()

	return toGoBool(ret)
}

func Fn_g_unix_fd_list_new() unsafe.Pointer {
	ret := C.g_unix_fd_list_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_fd_list_new_from_array : has non-string array param fds

func Fn_g_unix_fd_list_append(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) int {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_list_append(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_unix_fd_list_get(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) int {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_list_get(cValueInstance, cValue0, cError)

	return (int)(ret)
}

func Fn_g_unix_fd_list_get_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GUnixFDList)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_fd_list_get_length(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : g_unix_fd_list_peek_fds : has array return

// UNSUPPORTED : g_unix_fd_list_steal_fds : has array return

func Fn_g_unix_fd_message_new() unsafe.Pointer {
	ret := C.g_unix_fd_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_new_with_fd_list(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GUnixFDList)(unsafe.Pointer(param0))

	ret := C.g_unix_fd_message_new_with_fd_list(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_append_fd(paramInstance unsafe.Pointer, param0 int, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_message_append_fd(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_unix_fd_message_get_fd_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GUnixFDMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_fd_message_get_fd_list(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_fd_message_steal_fds : has array return

func Fn_g_unix_input_stream_new(param0 int, param1 bool) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	ret := C.g_unix_input_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_input_stream_get_close_fd(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_input_stream_get_close_fd(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_input_stream_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_input_stream_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_unix_input_stream_set_close_fd(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GUnixInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_unix_input_stream_set_close_fd(cValueInstance, cValue0)
}

func Fn_g_unix_mount_monitor_new() unsafe.Pointer {
	ret := C.g_unix_mount_monitor_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_monitor_set_rate_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GUnixMountMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	C.g_unix_mount_monitor_set_rate_limit(cValueInstance, cValue0)
}

func Fn_g_unix_mount_monitor_get() unsafe.Pointer {
	ret := C.g_unix_mount_monitor_get()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	ret := C.g_unix_output_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_output_stream_get_close_fd(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_output_stream_get_close_fd(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_output_stream_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_output_stream_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_unix_output_stream_set_close_fd(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GUnixOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_unix_output_stream_set_close_fd(cValueInstance, cValue0)
}

func Fn_g_unix_socket_address_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_unix_socket_address_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_socket_address_new_abstract : has non-string array param path

// UNSUPPORTED : g_unix_socket_address_new_with_type : has non-string array param path

func Fn_g_unix_socket_address_get_address_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_address_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_unix_socket_address_get_is_abstract(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_is_abstract(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_unix_socket_address_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_unix_socket_address_get_path_len(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GUnixSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_unix_socket_address_get_path_len(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_unix_socket_address_abstract_names_supported() bool {
	ret := C.g_unix_socket_address_abstract_names_supported()

	return toGoBool(ret)
}

func Fn_g_vfs_get_file_for_path(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_vfs_get_file_for_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_vfs_get_file_for_uri(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_vfs_get_file_for_uri(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : has array return

func Fn_g_vfs_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	ret := C.g_vfs_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_vfs_parse_name(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GVfs)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_vfs_parse_name(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_vfs_register_uri_scheme : has callback

func Fn_g_vfs_get_default() unsafe.Pointer {
	ret := C.g_vfs_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_vfs_get_local() unsafe.Pointer {
	ret := C.g_vfs_get_local()

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_connected_drives(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_monitor_get_connected_drives(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_mount_for_uuid(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_volume_monitor_get_mount_for_uuid(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_mounts(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_monitor_get_mounts(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_volume_for_uuid(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_volume_monitor_get_volume_for_uuid(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_volumes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_monitor_get_volumes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_adopt_orphan_mount(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GMount)(unsafe.Pointer(param0))

	ret := C.g_volume_monitor_adopt_orphan_mount(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get() unsafe.Pointer {
	ret := C.g_volume_monitor_get()

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_new(param0 int, param1 int) unsafe.Pointer {
	cValue0 := (C.GZlibCompressorFormat)(param0)

	cValue1 := (C.int)(param1)

	ret := C.g_zlib_compressor_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_get_file_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))

	ret := C.g_zlib_compressor_get_file_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_set_file_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GZlibCompressor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	C.g_zlib_compressor_set_file_info(cValueInstance, cValue0)
}

func Fn_g_zlib_decompressor_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GZlibCompressorFormat)(param0)

	ret := C.g_zlib_decompressor_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_decompressor_get_file_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GZlibDecompressor)(unsafe.Pointer(paramInstance))

	ret := C.g_zlib_decompressor_get_file_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_action_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_action_activate(cValueInstance, cValue0)
}

func Fn_g_action_change_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GVariant)(unsafe.Pointer(param0))

	C.g_action_change_state(cValueInstance, cValue0)
}

func Fn_g_action_get_enabled(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_enabled(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_action_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_action_get_parameter_type(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_parameter_type(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_action_get_state(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_state(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_action_get_state_hint(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_state_hint(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_action_get_state_type(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAction)(unsafe.Pointer(paramInstance))

	ret := C.g_action_get_state_type(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_action_group_action_added(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_action_group_action_added(cValueInstance, cValue0)
}

func Fn_g_action_group_action_enabled_changed(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	C.g_action_group_action_enabled_changed(cValueInstance, cValue0, cValue1)
}

func Fn_g_action_group_action_removed(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_action_group_action_removed(cValueInstance, cValue0)
}

func Fn_g_action_group_action_state_changed(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_action_group_action_state_changed(cValueInstance, cValue0, cValue1)
}

func Fn_g_action_group_activate_action(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_action_group_activate_action(cValueInstance, cValue0, cValue1)
}

func Fn_g_action_group_change_action_state(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	C.g_action_group_change_action_state(cValueInstance, cValue0, cValue1)
}

func Fn_g_action_group_get_action_enabled(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_get_action_enabled(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_action_group_get_action_parameter_type(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_get_action_parameter_type(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_group_get_action_state(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_get_action_state(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_group_get_action_state_hint(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_get_action_state_hint(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_group_get_action_state_type(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_get_action_state_type(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_group_has_action(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_group_has_action(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_action_group_list_actions : has array return

func Fn_g_action_group_query_action(paramInstance unsafe.Pointer, param0 string, param1 *bool, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 *unsafe.Pointer) bool {
	cValueInstance := (*C.GActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gboolean)(unsafe.Pointer(param1))

	cValue2 := (**C.GVariantType)(unsafe.Pointer(param2))

	cValue3 := (**C.GVariantType)(unsafe.Pointer(param3))

	cValue4 := (**C.GVariant)(unsafe.Pointer(param4))

	cValue5 := (**C.GVariant)(unsafe.Pointer(param5))

	ret := C.g_action_group_query_action(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_g_action_map_add_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAction)(unsafe.Pointer(param0))

	C.g_action_map_add_action(cValueInstance, cValue0)
}

// UNSUPPORTED : g_action_map_add_action_entries : has non-string array param entries

func Fn_g_action_map_lookup_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_action_map_lookup_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_action_map_remove_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GActionMap)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_action_map_remove_action(cValueInstance, cValue0)
}

func Fn_g_app_info_add_supports_type(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_add_supports_type(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_can_delete(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_can_delete(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_app_info_can_remove_supports_type(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_can_remove_supports_type(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_app_info_delete(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_delete(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_app_info_dup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_dup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))

	ret := C.g_app_info_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_app_info_get_commandline(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_commandline(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_app_info_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_app_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_app_info_get_executable(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_executable(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_app_info_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_app_info_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_app_info_get_supported_types : has array return

func Fn_g_app_info_launch(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_app_info_launch(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_launch_uris(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_app_info_launch_uris(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_app_info_launch_uris_async : has callback

func Fn_g_app_info_remove_supports_type(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_remove_supports_type(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_default_for_extension(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_default_for_extension(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_default_for_type(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_default_for_type(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_last_used_for_type(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_last_used_for_type(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_should_show(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_should_show(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_app_info_supports_files(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_supports_files(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_app_info_supports_uris(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_supports_uris(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_async_initable_init_async : has callback

func Fn_g_async_initable_init_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAsyncInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_init_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_async_initable_new_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAsyncInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_new_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_async_result_get_source_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_async_result_get_source_object(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_async_result_get_user_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GAsyncResult)(unsafe.Pointer(paramInstance))

	ret := C.g_async_result_get_user_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_async_result_is_tagged(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	ret := C.g_async_result_is_tagged(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_async_result_legacy_propagate_error(paramInstance unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAsyncResult)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.g_async_result_legacy_propagate_error(cValueInstance, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_converter_convert : has non-string array param inbuf

func Fn_g_converter_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GConverter)(unsafe.Pointer(paramInstance))

	C.g_converter_reset(cValueInstance)
}

func Fn_g_dbus_interface_dup_object(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_dup_object(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_get_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_interface_get_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_set_object(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusInterface)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusObject)(unsafe.Pointer(param0))

	C.g_dbus_interface_set_object(cValueInstance, cValue0)
}

func Fn_g_dbus_object_get_interface(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_get_interface(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_get_interfaces(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_get_interfaces(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObject)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_get_interface(paramInstance unsafe.Pointer, param0 string, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_dbus_object_manager_get_interface(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_get_object(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_dbus_object_manager_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_object_manager_get_object_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_get_object_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_dbus_object_manager_get_objects(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDBusObjectManager)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_object_manager_get_objects(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_datagram_based_receive_messages : has non-string array param messages

// UNSUPPORTED : g_datagram_based_send_messages : has non-string array param messages

func Fn_g_desktop_app_info_lookup_get_default_for_uri_scheme(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GDesktopAppInfoLookup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_desktop_app_info_lookup_get_default_for_uri_scheme(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_can_eject(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_eject(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_poll_for_media(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_poll_for_media(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_start_degraded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_start_degraded(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_can_stop(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_can_stop(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject : has callback

func Fn_g_drive_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject_with_operation : has callback

func Fn_g_drive_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_enumerate_identifiers : has array return

func Fn_g_drive_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_get_identifier(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_drive_get_identifier(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_drive_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_drive_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_drive_get_start_stop_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_start_stop_type(cValueInstance)

	return (int)(ret)
}

func Fn_g_drive_get_symbolic_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_symbolic_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_get_volumes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_volumes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_has_media(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_has_media(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_has_volumes(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_has_volumes(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_is_media_check_automatic(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_is_media_check_automatic(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_drive_is_media_removable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_is_media_removable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_poll_for_media : has callback

func Fn_g_drive_poll_for_media_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_poll_for_media_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_start : has callback

func Fn_g_drive_start_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_start_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_stop : has callback

func Fn_g_drive_stop_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_stop_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dtls_connection_close_async : has callback

// UNSUPPORTED : g_dtls_connection_handshake_async : has callback

// UNSUPPORTED : g_dtls_connection_shutdown_async : has callback

func Fn_g_file_append_to(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileCreateFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_append_to_async : has callback

func Fn_g_file_append_to_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_copy : has callback

// UNSUPPORTED : g_file_copy_async : has callback

func Fn_g_file_copy_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cValue1 := (C.GFileCopyFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_file_copy_attributes(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_file_copy_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_copy_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_create(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileCreateFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_create(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_async : has callback

func Fn_g_file_create_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_create_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_create_readwrite(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileCreateFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_readwrite_async : has callback

func Fn_g_file_create_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_delete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_delete(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_delete_async : has callback

func Fn_g_file_delete_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_delete_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_dup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_dup(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_eject_mountable : has callback

func Fn_g_file_eject_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_eject_mountable_with_operation : has callback

func Fn_g_file_eject_mountable_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_enumerate_children(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileQueryInfoFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerate_children(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_enumerate_children_async : has callback

func Fn_g_file_enumerate_children_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerate_children_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	ret := C.g_file_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_find_enclosing_mount(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_find_enclosing_mount(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_find_enclosing_mount_async : has callback

func Fn_g_file_find_enclosing_mount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_find_enclosing_mount_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_basename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_basename(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_get_child(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_get_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_child_for_display_name(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_file_get_child_for_display_name(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_parse_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_parse_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_get_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_get_relative_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	ret := C.g_file_get_relative_path(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_file_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_get_uri_scheme(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_get_uri_scheme(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_file_has_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	ret := C.g_file_has_parent(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_has_prefix(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	ret := C.g_file_has_prefix(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_has_uri_scheme(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_has_uri_scheme(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_hash(paramInstance unsafe.Pointer) uint {
	cValueInstance := (C.gconstpointer)(paramInstance)

	ret := C.g_file_hash(cValueInstance)

	return (uint)(ret)
}

func Fn_g_file_is_native(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_is_native(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_load_bytes_async : has callback

// UNSUPPORTED : g_file_load_contents : has non-string array param contents

// UNSUPPORTED : g_file_load_contents_async : has callback

// UNSUPPORTED : g_file_load_contents_finish : has non-string array param contents

// UNSUPPORTED : g_file_load_partial_contents_async : has callback

// UNSUPPORTED : g_file_load_partial_contents_finish : has non-string array param contents

func Fn_g_file_make_directory(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_make_directory_async : has callback

func Fn_g_file_make_directory_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_make_directory_with_parents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory_with_parents(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_make_symbolic_link(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_make_symbolic_link(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_measure_disk_usage : has callback

// UNSUPPORTED : g_file_measure_disk_usage_async : has callback

func Fn_g_file_measure_disk_usage_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64, param2 *uint64, param3 *uint64, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.guint64)(unsafe.Pointer(param1))

	cValue2 := (*C.guint64)(unsafe.Pointer(param2))

	cValue3 := (*C.guint64)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_measure_disk_usage_finish(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_monitor(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileMonitorFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_directory(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileMonitorFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor_directory(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_file(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileMonitorFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor_file(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_mount_enclosing_volume : has callback

func Fn_g_file_mount_enclosing_volume_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_enclosing_volume_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_mount_mountable : has callback

func Fn_g_file_mount_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_mountable_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_move : has callback

func Fn_g_file_open_readwrite(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_open_readwrite_async : has callback

func Fn_g_file_open_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_poll_mountable : has callback

func Fn_g_file_poll_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_poll_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_query_default_handler(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_default_handler(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_default_handler_async : has callback

func Fn_g_file_query_exists(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	ret := C.g_file_query_exists(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_file_query_file_type(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) int {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileQueryInfoFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	ret := C.g_file_query_file_type(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_g_file_query_filesystem_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_query_filesystem_info(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_filesystem_info_async : has callback

func Fn_g_file_query_filesystem_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_filesystem_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_info(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileQueryInfoFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_file_query_info(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_info_async : has callback

func Fn_g_file_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_info_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_settable_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_settable_attributes(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_writable_namespaces(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_writable_namespaces(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_read(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_read(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_read_async : has callback

func Fn_g_file_read_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_read_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_replace(paramInstance unsafe.Pointer, param0 string, param1 bool, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	cValue2 := (C.GFileCreateFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_replace(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_replace_async : has callback

// UNSUPPORTED : g_file_replace_contents : has non-string array param contents

// UNSUPPORTED : g_file_replace_contents_async : has callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : has callback

func Fn_g_file_replace_contents_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cError := (**C.GError)(error)

	ret := C.g_file_replace_contents_finish(cValueInstance, cValue0, cValue1, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return toGoBool(ret)
}

func Fn_g_file_replace_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_replace_readwrite(paramInstance unsafe.Pointer, param0 string, param1 bool, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	cValue2 := (C.GFileCreateFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_replace_readwrite_async : has callback

func Fn_g_file_replace_readwrite_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_resolve_relative_path(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_resolve_relative_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_set_attribute(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, param3 int, param4 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GFileAttributeType)(param1)

	cValue2 := (C.gpointer)(param2)

	cValue3 := (C.GFileQueryInfoFlags)(param3)

	cValue4 := (*C.GCancellable)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_byte_string(paramInstance unsafe.Pointer, param0 string, param1 string, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_byte_string(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_int32(paramInstance unsafe.Pointer, param0 string, param1 int32, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint32)(param1)

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_int32(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_int64(paramInstance unsafe.Pointer, param0 string, param1 int64, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint64)(param1)

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_int64(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_string(paramInstance unsafe.Pointer, param0 string, param1 string, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_string(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_uint32(paramInstance unsafe.Pointer, param0 string, param1 uint32, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint32)(param1)

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_uint32(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_uint64(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 int, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint64)(param1)

	cValue2 := (C.GFileQueryInfoFlags)(param2)

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_uint64(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_set_attributes_async : has callback

func Fn_g_file_set_attributes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (**C.GFileInfo)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attributes_finish(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attributes_from_info(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFileInfo)(unsafe.Pointer(param0))

	cValue1 := (C.GFileQueryInfoFlags)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attributes_from_info(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_display_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_set_display_name(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_set_display_name_async : has callback

func Fn_g_file_set_display_name_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_set_display_name_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_start_mountable : has callback

func Fn_g_file_start_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_start_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_stop_mountable : has callback

func Fn_g_file_stop_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_stop_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_supports_thread_contexts(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_supports_thread_contexts(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_file_trash(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_trash(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_trash_async : has callback

func Fn_g_file_trash_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_trash_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_unmount_mountable : has callback

func Fn_g_file_unmount_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_unmount_mountable_with_operation : has callback

func Fn_g_file_unmount_mountable_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_file_descriptor_based_get_fd(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GFileDescriptorBased)(unsafe.Pointer(paramInstance))

	ret := C.g_file_descriptor_based_get_fd(cValueInstance)

	return (int)(ret)
}

func Fn_g_icon_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.g_icon_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_g_icon_serialize(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_icon_serialize(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_icon_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GIcon)(unsafe.Pointer(paramInstance))

	ret := C.g_icon_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_initable_init(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GInitable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_initable_init(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_list_model_get_item_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GListModel)(unsafe.Pointer(paramInstance))

	ret := C.g_list_model_get_item_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_list_model_get_n_items(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GListModel)(unsafe.Pointer(paramInstance))

	ret := C.g_list_model_get_n_items(cValueInstance)

	return (uint)(ret)
}

func Fn_g_list_model_get_object(paramInstance unsafe.Pointer, param0 uint) unsafe.Pointer {
	cValueInstance := (*C.GListModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.g_list_model_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_list_model_items_changed(paramInstance unsafe.Pointer, param0 uint, param1 uint, param2 uint) {
	cValueInstance := (*C.GListModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	C.g_list_model_items_changed(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_loadable_icon_load(paramInstance unsafe.Pointer, param0 int, param1 *string, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GLoadableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (C.int)(param0)

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_loadable_icon_load(cValueInstance, cValue0, cValue1, cValue2, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_loadable_icon_load_async : has callback

func Fn_g_loadable_icon_load_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GLoadableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cError := (**C.GError)(error)

	ret := C.g_loadable_icon_load_finish(cValueInstance, cValue0, cValue1, cError)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return unsafe.Pointer(ret)
}

func Fn_g_mount_can_eject(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_can_eject(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_mount_can_unmount(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_can_unmount(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_eject : has callback

func Fn_g_mount_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_eject_with_operation : has callback

func Fn_g_mount_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_mount_get_default_location(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_default_location(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_drive(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_drive(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_mount_get_root(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_root(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_mount_get_symbolic_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_symbolic_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_uuid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_uuid(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_mount_get_volume(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_get_volume(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_mount_guess_content_type : has callback

// UNSUPPORTED : g_mount_guess_content_type_finish : has array return

// UNSUPPORTED : g_mount_guess_content_type_sync : has array return

func Fn_g_mount_is_shadowed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	ret := C.g_mount_is_shadowed(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_remount : has callback

func Fn_g_mount_remount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_remount_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_mount_shadow(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	C.g_mount_shadow(cValueInstance)
}

// UNSUPPORTED : g_mount_unmount : has callback

func Fn_g_mount_unmount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_unmount_with_operation : has callback

func Fn_g_mount_unmount_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_mount_unshadow(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	C.g_mount_unshadow(cValueInstance)
}

func Fn_g_network_monitor_can_reach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_network_monitor_can_reach_async : has callback

func Fn_g_network_monitor_can_reach_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_network_monitor_get_connectivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_network_monitor_get_connectivity(cValueInstance)

	return (int)(ret)
}

func Fn_g_network_monitor_get_network_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_network_monitor_get_network_available(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_network_monitor_get_network_metered(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	ret := C.g_network_monitor_get_network_metered(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_pollable_input_stream_can_poll(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPollableInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_pollable_input_stream_can_poll(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_pollable_input_stream_create_source(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GPollableInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	ret := C.g_pollable_input_stream_create_source(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_pollable_input_stream_is_readable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPollableInputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_pollable_input_stream_is_readable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_pollable_input_stream_read_nonblocking : has non-string array param buffer

func Fn_g_pollable_output_stream_can_poll(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPollableOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_pollable_output_stream_can_poll(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_pollable_output_stream_create_source(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GPollableOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	ret := C.g_pollable_output_stream_create_source(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_pollable_output_stream_is_writable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GPollableOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_pollable_output_stream_is_writable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_pollable_output_stream_write_nonblocking : has non-string array param buffer

// UNSUPPORTED : g_pollable_output_stream_writev_nonblocking : has non-string array param vectors

func Fn_g_proxy_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIOStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GProxyAddress)(unsafe.Pointer(param1))

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_proxy_connect_async : has callback

func Fn_g_proxy_connect_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_supports_hostname(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GProxy)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_supports_hostname(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_proxy_resolver_is_supported(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GProxyResolver)(unsafe.Pointer(paramInstance))

	ret := C.g_proxy_resolver_is_supported(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_proxy_resolver_lookup : has array return

// UNSUPPORTED : g_proxy_resolver_lookup_async : has callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : has array return

func Fn_g_remote_action_group_activate_action_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GRemoteActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_remote_action_group_activate_action_full(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_remote_action_group_change_action_state_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GRemoteActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GVariant)(unsafe.Pointer(param1))

	cValue2 := (*C.GVariant)(unsafe.Pointer(param2))

	C.g_remote_action_group_change_action_state_full(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_seekable_can_seek(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSeekable)(unsafe.Pointer(paramInstance))

	ret := C.g_seekable_can_seek(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_seekable_can_truncate(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSeekable)(unsafe.Pointer(paramInstance))

	ret := C.g_seekable_can_truncate(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_seekable_seek(paramInstance unsafe.Pointer, param0 int64, param1 int, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSeekable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.goffset)(param0)

	cValue1 := (C.GSeekType)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_seekable_seek(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_g_seekable_tell(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GSeekable)(unsafe.Pointer(paramInstance))

	ret := C.g_seekable_tell(cValueInstance)

	return (int64)(ret)
}

func Fn_g_seekable_truncate(paramInstance unsafe.Pointer, param0 int64, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GSeekable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.goffset)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_seekable_truncate(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connectable_enumerate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnectable)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connectable_enumerate(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connectable_proxy_enumerate(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketConnectable)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_connectable_proxy_enumerate(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_backend_get_certificate_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_certificate_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_tls_backend_get_client_connection_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_client_connection_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_tls_backend_get_default_database(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_default_database(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_backend_get_file_database_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_file_database_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_tls_backend_get_server_connection_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_get_server_connection_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_g_tls_backend_supports_tls(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsBackend)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_backend_supports_tls(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tls_client_connection_copy_session_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GTlsClientConnection)(unsafe.Pointer(param0))

	C.g_tls_client_connection_copy_session_state(cValueInstance, cValue0)
}

func Fn_g_tls_client_connection_get_accepted_cas(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_client_connection_get_accepted_cas(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_client_connection_get_server_identity(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_client_connection_get_server_identity(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_client_connection_get_use_ssl3(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_client_connection_get_use_ssl3(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_tls_client_connection_get_validation_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_client_connection_get_validation_flags(cValueInstance)

	return (int)(ret)
}

func Fn_g_tls_client_connection_set_server_identity(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSocketConnectable)(unsafe.Pointer(param0))

	C.g_tls_client_connection_set_server_identity(cValueInstance, cValue0)
}

func Fn_g_tls_client_connection_set_use_ssl3(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_tls_client_connection_set_use_ssl3(cValueInstance, cValue0)
}

func Fn_g_tls_client_connection_set_validation_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GTlsClientConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GTlsCertificateFlags)(param0)

	C.g_tls_client_connection_set_validation_flags(cValueInstance, cValue0)
}

func Fn_g_volume_can_eject(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_can_eject(cValueInstance)

	return toGoBool(ret)
}

func Fn_g_volume_can_mount(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_can_mount(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_eject : has callback

func Fn_g_volume_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_eject_with_operation : has callback

func Fn_g_volume_eject_with_operation_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_with_operation_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_enumerate_identifiers : has array return

func Fn_g_volume_get_activation_root(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_activation_root(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_drive(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_drive(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_identifier(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_volume_get_identifier(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_g_volume_get_mount(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_mount(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_volume_get_sort_key(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_sort_key(cValueInstance)

	return C.GoString(ret)
}

func Fn_g_volume_get_symbolic_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_symbolic_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_uuid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_uuid(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_volume_mount : has callback

func Fn_g_volume_mount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_volume_mount_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_volume_should_automount(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_should_automount(cValueInstance)

	return toGoBool(ret)
}
