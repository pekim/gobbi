// Code generated - DO NOT EDIT.
// +build gio_2.26

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
/*

static GDBusMessage* c_g_dbus_message_new_method_error(GDBusMessage* method_call_message, const gchar* error_name, const gchar* error_message_format) {
    return g_dbus_message_new_method_error(method_call_message, error_name, error_message_format, NULL);
}
*/
/*

static GDBusMessage* c_g_dbus_message_new_method_error_valist(GDBusMessage* method_call_message, const gchar* error_name, const gchar* error_message_format) {
    return g_dbus_message_new_method_error_valist(method_call_message, error_name, error_message_format, NULL);
}
*/
/*

static void c_g_dbus_method_invocation_return_error(GDBusMethodInvocation* invocation, GQuark domain, gint code, const gchar* format) {
    return g_dbus_method_invocation_return_error(invocation, domain, code, format, NULL);
}
*/
/*

static void c_g_dbus_method_invocation_return_error_valist(GDBusMethodInvocation* invocation, GQuark domain, gint code, const gchar* format) {
    return g_dbus_method_invocation_return_error_valist(invocation, domain, code, format, NULL);
}
*/
/*

static void c_g_settings_get(GSettings* settings, const gchar* key, const gchar* format) {
    return g_settings_get(settings, key, format, NULL);
}
*/
/*

static gboolean c_g_settings_set(GSettings* settings, const gchar* key, const gchar* format) {
    return g_settings_set(settings, key, format, NULL);
}
*/
/*

static void c_g_simple_async_result_set_error(GSimpleAsyncResult* simple, GQuark domain, gint code, const char* format) {
    return g_simple_async_result_set_error(simple, domain, code, format, NULL);
}
*/
/*

static void c_g_simple_async_result_set_error_va(GSimpleAsyncResult* simple, GQuark domain, gint code, const char* format) {
    return g_simple_async_result_set_error_va(simple, domain, code, format, NULL);
}
*/
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

func Fn_g_dbus_annotation_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusAnnotationInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_annotation_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_annotation_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusAnnotationInfo)(unsafe.Pointer(info))

	C.g_dbus_annotation_info_unref(c_info)
}

func Fn_g_dbus_annotation_info_lookup(annotations []unsafe.Pointer, name string) string {
	c_annotations := (**C.GDBusAnnotationInfo)(unsafe.Pointer(&annotations[0]))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_dbus_annotation_info_lookup(c_annotations, c_name)

	return C.GoString(ret)
}

func Fn_g_dbus_arg_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusArgInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_arg_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_arg_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusArgInfo)(unsafe.Pointer(info))

	C.g_dbus_arg_info_unref(c_info)
}

func Fn_g_dbus_interface_info_generate_xml(info unsafe.Pointer, indent uint, stringBuilder unsafe.Pointer) {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	c_indent := (C.guint)(indent)

	c_stringBuilder := (*C.GString)(unsafe.Pointer(stringBuilder))

	C.g_dbus_interface_info_generate_xml(c_info, c_indent, c_stringBuilder)
}

func Fn_g_dbus_interface_info_lookup_method(info unsafe.Pointer, name string) unsafe.Pointer {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_dbus_interface_info_lookup_method(c_info, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_property(info unsafe.Pointer, name string) unsafe.Pointer {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_dbus_interface_info_lookup_property(c_info, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_lookup_signal(info unsafe.Pointer, name string) unsafe.Pointer {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_dbus_interface_info_lookup_signal(c_info, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_interface_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_interface_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_interface_info_unref(c_info)
}

func Fn_g_dbus_method_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusMethodInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_method_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusMethodInfo)(unsafe.Pointer(info))

	C.g_dbus_method_info_unref(c_info)
}

func Fn_g_dbus_node_info_new_for_xml(xmlData string, error unsafe.Pointer) unsafe.Pointer {
	c_xmlData := (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(c_xmlData))

	cError := (**C.GError)(error)

	ret := C.g_dbus_node_info_new_for_xml(c_xmlData, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_generate_xml(info unsafe.Pointer, indent uint, stringBuilder unsafe.Pointer) {
	c_info := (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	c_indent := (C.guint)(indent)

	c_stringBuilder := (*C.GString)(unsafe.Pointer(stringBuilder))

	C.g_dbus_node_info_generate_xml(c_info, c_indent, c_stringBuilder)
}

func Fn_g_dbus_node_info_lookup_interface(info unsafe.Pointer, name string) unsafe.Pointer {
	c_info := (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_dbus_node_info_lookup_interface(c_info, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_node_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_node_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	C.g_dbus_node_info_unref(c_info)
}

func Fn_g_dbus_property_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusPropertyInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_property_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_property_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusPropertyInfo)(unsafe.Pointer(info))

	C.g_dbus_property_info_unref(c_info)
}

func Fn_g_dbus_signal_info_ref(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GDBusSignalInfo)(unsafe.Pointer(info))

	ret := C.g_dbus_signal_info_ref(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_signal_info_unref(info unsafe.Pointer) {
	c_info := (*C.GDBusSignalInfo)(unsafe.Pointer(info))

	C.g_dbus_signal_info_unref(c_info)
}

func Fn_g_file_attribute_info_list_new() unsafe.Pointer {
	ret := C.g_file_attribute_info_list_new()

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_add(list unsafe.Pointer, name string, type_ int, flags int) {
	c_list := (*C.GFileAttributeInfoList)(unsafe.Pointer(list))

	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_type_ := (C.GFileAttributeType)(type_)

	c_flags := (C.GFileAttributeInfoFlags)(flags)

	C.g_file_attribute_info_list_add(c_list, c_name, c_type_, c_flags)
}

func Fn_g_file_attribute_info_list_dup(list unsafe.Pointer) unsafe.Pointer {
	c_list := (*C.GFileAttributeInfoList)(unsafe.Pointer(list))

	ret := C.g_file_attribute_info_list_dup(c_list)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_lookup(list unsafe.Pointer, name string) unsafe.Pointer {
	c_list := (*C.GFileAttributeInfoList)(unsafe.Pointer(list))

	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_file_attribute_info_list_lookup(c_list, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_ref(list unsafe.Pointer) unsafe.Pointer {
	c_list := (*C.GFileAttributeInfoList)(unsafe.Pointer(list))

	ret := C.g_file_attribute_info_list_ref(c_list)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_info_list_unref(list unsafe.Pointer) {
	c_list := (*C.GFileAttributeInfoList)(unsafe.Pointer(list))

	C.g_file_attribute_info_list_unref(c_list)
}

func Fn_g_file_attribute_matcher_new(attributes string) unsafe.Pointer {
	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	ret := C.g_file_attribute_matcher_new(c_attributes)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_enumerate_namespace(matcher unsafe.Pointer, ns string) bool {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	c_ns := (*C.char)(C.CString(ns))
	defer C.free(unsafe.Pointer(c_ns))

	ret := C.g_file_attribute_matcher_enumerate_namespace(c_matcher, c_ns)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_enumerate_next(matcher unsafe.Pointer) string {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	ret := C.g_file_attribute_matcher_enumerate_next(c_matcher)

	return C.GoString(ret)
}

func Fn_g_file_attribute_matcher_matches(matcher unsafe.Pointer, attribute string) bool {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_attribute_matcher_matches(c_matcher, c_attribute)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_matches_only(matcher unsafe.Pointer, attribute string) bool {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_attribute_matcher_matches_only(c_matcher, c_attribute)

	return toGoBool(ret)
}

func Fn_g_file_attribute_matcher_ref(matcher unsafe.Pointer) unsafe.Pointer {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	ret := C.g_file_attribute_matcher_ref(c_matcher)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_subtract(matcher unsafe.Pointer, subtract unsafe.Pointer) unsafe.Pointer {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	c_subtract := (*C.GFileAttributeMatcher)(unsafe.Pointer(subtract))

	ret := C.g_file_attribute_matcher_subtract(c_matcher, c_subtract)

	return unsafe.Pointer(ret)
}

func Fn_g_file_attribute_matcher_unref(matcher unsafe.Pointer) {
	c_matcher := (*C.GFileAttributeMatcher)(unsafe.Pointer(matcher))

	C.g_file_attribute_matcher_unref(c_matcher)
}

func Fn_g_io_extension_get_name(extension unsafe.Pointer) string {
	c_extension := (*C.GIOExtension)(unsafe.Pointer(extension))

	ret := C.g_io_extension_get_name(c_extension)

	return C.GoString(ret)
}

func Fn_g_io_extension_get_priority(extension unsafe.Pointer) int {
	c_extension := (*C.GIOExtension)(unsafe.Pointer(extension))

	ret := C.g_io_extension_get_priority(c_extension)

	return (int)(ret)
}

func Fn_g_io_extension_get_type(extension unsafe.Pointer) uint64 {
	c_extension := (*C.GIOExtension)(unsafe.Pointer(extension))

	ret := C.g_io_extension_get_type(c_extension)

	return (uint64)(ret)
}

func Fn_g_io_extension_ref_class(extension unsafe.Pointer) unsafe.Pointer {
	c_extension := (*C.GIOExtension)(unsafe.Pointer(extension))

	ret := C.g_io_extension_ref_class(c_extension)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_extension_by_name(extensionPoint unsafe.Pointer, name string) unsafe.Pointer {
	c_extensionPoint := (*C.GIOExtensionPoint)(unsafe.Pointer(extensionPoint))

	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_io_extension_point_get_extension_by_name(c_extensionPoint, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_extensions(extensionPoint unsafe.Pointer) unsafe.Pointer {
	c_extensionPoint := (*C.GIOExtensionPoint)(unsafe.Pointer(extensionPoint))

	ret := C.g_io_extension_point_get_extensions(c_extensionPoint)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_get_required_type(extensionPoint unsafe.Pointer) uint64 {
	c_extensionPoint := (*C.GIOExtensionPoint)(unsafe.Pointer(extensionPoint))

	ret := C.g_io_extension_point_get_required_type(c_extensionPoint)

	return (uint64)(ret)
}

func Fn_g_io_extension_point_set_required_type(extensionPoint unsafe.Pointer, type_ uint64) {
	c_extensionPoint := (*C.GIOExtensionPoint)(unsafe.Pointer(extensionPoint))

	c_type_ := (C.GType)(type_)

	C.g_io_extension_point_set_required_type(c_extensionPoint, c_type_)
}

func Fn_g_io_extension_point_implement(extensionPointName string, type_ uint64, extensionName string, priority int) unsafe.Pointer {
	c_extensionPointName := (*C.char)(C.CString(extensionPointName))
	defer C.free(unsafe.Pointer(c_extensionPointName))

	c_type_ := (C.GType)(type_)

	c_extensionName := (*C.char)(C.CString(extensionName))
	defer C.free(unsafe.Pointer(c_extensionName))

	c_priority := (C.gint)(priority)

	ret := C.g_io_extension_point_implement(c_extensionPointName, c_type_, c_extensionName, c_priority)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_lookup(name string) unsafe.Pointer {
	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_io_extension_point_lookup(c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_io_extension_point_register(name string) unsafe.Pointer {
	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_io_extension_point_register(c_name)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_io_module_scope_block : blacklisted

// UNSUPPORTED : g_io_module_scope_free : blacklisted

// UNSUPPORTED : g_io_module_scope_new : blacklisted

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

func Fn_g_settings_schema_get_id(schema unsafe.Pointer) string {
	c_schema := (*C.GSettingsSchema)(unsafe.Pointer(schema))

	ret := C.g_settings_schema_get_id(c_schema)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

func Fn_g_srv_target_new(hostname string, port uint16, priority uint16, weight uint16) unsafe.Pointer {
	c_hostname := (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	c_priority := (C.guint16)(priority)

	c_weight := (C.guint16)(weight)

	ret := C.g_srv_target_new(c_hostname, c_port, c_priority, c_weight)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_copy(target unsafe.Pointer) unsafe.Pointer {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	ret := C.g_srv_target_copy(c_target)

	return unsafe.Pointer(ret)
}

func Fn_g_srv_target_free(target unsafe.Pointer) {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	C.g_srv_target_free(c_target)
}

func Fn_g_srv_target_get_hostname(target unsafe.Pointer) string {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	ret := C.g_srv_target_get_hostname(c_target)

	return C.GoString(ret)
}

func Fn_g_srv_target_get_port(target unsafe.Pointer) uint16 {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	ret := C.g_srv_target_get_port(c_target)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_priority(target unsafe.Pointer) uint16 {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	ret := C.g_srv_target_get_priority(c_target)

	return (uint16)(ret)
}

func Fn_g_srv_target_get_weight(target unsafe.Pointer) uint16 {
	c_target := (*C.GSrvTarget)(unsafe.Pointer(target))

	ret := C.g_srv_target_get_weight(c_target)

	return (uint16)(ret)
}

func Fn_g_srv_target_list_sort(targets unsafe.Pointer) unsafe.Pointer {
	c_targets := (*C.GList)(unsafe.Pointer(targets))

	ret := C.g_srv_target_list_sort(c_targets)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_point_compare(mount1 unsafe.Pointer, mount2 unsafe.Pointer) int {
	c_mount1 := (*C.GUnixMountPoint)(unsafe.Pointer(mount1))

	c_mount2 := (*C.GUnixMountPoint)(unsafe.Pointer(mount2))

	ret := C.g_unix_mount_point_compare(c_mount1, c_mount2)

	return (int)(ret)
}

func Fn_g_unix_mount_point_free(mountPoint unsafe.Pointer) {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	C.g_unix_mount_point_free(c_mountPoint)
}

func Fn_g_unix_mount_point_get_device_path(mountPoint unsafe.Pointer) string {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_get_device_path(c_mountPoint)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_get_fs_type(mountPoint unsafe.Pointer) string {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_get_fs_type(c_mountPoint)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_get_mount_path(mountPoint unsafe.Pointer) string {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_get_mount_path(c_mountPoint)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_guess_can_eject(mountPoint unsafe.Pointer) bool {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_guess_can_eject(c_mountPoint)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_guess_icon(mountPoint unsafe.Pointer) unsafe.Pointer {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_guess_icon(c_mountPoint)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_point_guess_name(mountPoint unsafe.Pointer) string {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_guess_name(c_mountPoint)

	return C.GoString(ret)
}

func Fn_g_unix_mount_point_is_loopback(mountPoint unsafe.Pointer) bool {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_is_loopback(c_mountPoint)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_is_readonly(mountPoint unsafe.Pointer) bool {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_is_readonly(c_mountPoint)

	return toGoBool(ret)
}

func Fn_g_unix_mount_point_is_user_mountable(mountPoint unsafe.Pointer) bool {
	c_mountPoint := (*C.GUnixMountPoint)(unsafe.Pointer(mountPoint))

	ret := C.g_unix_mount_point_is_user_mountable(c_mountPoint)

	return toGoBool(ret)
}

// UNSUPPORTED : g_action_parse_detailed_name : parameter 'action_name' is non array with indirect count > 1

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

func Fn_g_bus_get_finish(res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_finish(c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_bus_get_sync(busType int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_busType := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_bus_get_sync(c_busType, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

func Fn_g_bus_own_name_on_connection_with_closures(connection unsafe.Pointer, name string, flags int, nameAcquiredClosure unsafe.Pointer, nameLostClosure unsafe.Pointer) uint {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_nameAcquiredClosure := (*C.GClosure)(unsafe.Pointer(nameAcquiredClosure))

	c_nameLostClosure := (*C.GClosure)(unsafe.Pointer(nameLostClosure))

	ret := C.g_bus_own_name_on_connection_with_closures(c_connection, c_name, c_flags, c_nameAcquiredClosure, c_nameLostClosure)

	return (uint)(ret)
}

func Fn_g_bus_own_name_with_closures(busType int, name string, flags int, busAcquiredClosure unsafe.Pointer, nameAcquiredClosure unsafe.Pointer, nameLostClosure unsafe.Pointer) uint {
	c_busType := (C.GBusType)(busType)

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameOwnerFlags)(flags)

	c_busAcquiredClosure := (*C.GClosure)(unsafe.Pointer(busAcquiredClosure))

	c_nameAcquiredClosure := (*C.GClosure)(unsafe.Pointer(nameAcquiredClosure))

	c_nameLostClosure := (*C.GClosure)(unsafe.Pointer(nameLostClosure))

	ret := C.g_bus_own_name_with_closures(c_busType, c_name, c_flags, c_busAcquiredClosure, c_nameAcquiredClosure, c_nameLostClosure)

	return (uint)(ret)
}

func Fn_g_bus_unown_name(ownerId uint) {
	c_ownerId := (C.guint)(ownerId)

	C.g_bus_unown_name(c_ownerId)
}

func Fn_g_bus_unwatch_name(watcherId uint) {
	c_watcherId := (C.guint)(watcherId)

	C.g_bus_unwatch_name(c_watcherId)
}

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

func Fn_g_bus_watch_name_on_connection_with_closures(connection unsafe.Pointer, name string, flags int, nameAppearedClosure unsafe.Pointer, nameVanishedClosure unsafe.Pointer) uint {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_nameAppearedClosure := (*C.GClosure)(unsafe.Pointer(nameAppearedClosure))

	c_nameVanishedClosure := (*C.GClosure)(unsafe.Pointer(nameVanishedClosure))

	ret := C.g_bus_watch_name_on_connection_with_closures(c_connection, c_name, c_flags, c_nameAppearedClosure, c_nameVanishedClosure)

	return (uint)(ret)
}

func Fn_g_bus_watch_name_with_closures(busType int, name string, flags int, nameAppearedClosure unsafe.Pointer, nameVanishedClosure unsafe.Pointer) uint {
	c_busType := (C.GBusType)(busType)

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GBusNameWatcherFlags)(flags)

	c_nameAppearedClosure := (*C.GClosure)(unsafe.Pointer(nameAppearedClosure))

	c_nameVanishedClosure := (*C.GClosure)(unsafe.Pointer(nameVanishedClosure))

	ret := C.g_bus_watch_name_with_closures(c_busType, c_name, c_flags, c_nameAppearedClosure, c_nameVanishedClosure)

	return (uint)(ret)
}

func Fn_g_content_type_can_be_executable(type_ string) bool {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	ret := C.g_content_type_can_be_executable(c_type_)

	return toGoBool(ret)
}

func Fn_g_content_type_equals(type1 string, type2 string) bool {
	c_type1 := (*C.gchar)(C.CString(type1))
	defer C.free(unsafe.Pointer(c_type1))

	c_type2 := (*C.gchar)(C.CString(type2))
	defer C.free(unsafe.Pointer(c_type2))

	ret := C.g_content_type_equals(c_type1, c_type2)

	return toGoBool(ret)
}

func Fn_g_content_type_from_mime_type(mimeType string) string {
	c_mimeType := (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(c_mimeType))

	ret := C.g_content_type_from_mime_type(c_mimeType)

	return C.GoString(ret)
}

func Fn_g_content_type_get_description(type_ string) string {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	ret := C.g_content_type_get_description(c_type_)

	return C.GoString(ret)
}

func Fn_g_content_type_get_icon(type_ string) unsafe.Pointer {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	ret := C.g_content_type_get_icon(c_type_)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

func Fn_g_content_type_get_mime_type(type_ string) string {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	ret := C.g_content_type_get_mime_type(c_type_)

	return C.GoString(ret)
}

func Fn_g_content_type_guess(filename *string, data []uint8, dataSize uint64, resultUncertain *bool) string {
	var c_filenameValue (*C.gchar)
	if filename != nil {
		c_filenameValue = (*C.gchar)(C.CString(*filename))
		defer C.free(unsafe.Pointer(c_filenameValue))
	}
	c_filename := c_filenameValue

	c_data := (*C.guchar)(unsafe.Pointer(&data[0]))

	c_dataSize := (C.gsize)(dataSize)

	c_resultUncertain := (*C.gboolean)(unsafe.Pointer(resultUncertain))

	ret := C.g_content_type_guess(c_filename, c_data, c_dataSize, c_resultUncertain)

	return C.GoString(ret)
}

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

func Fn_g_content_type_is_a(type_ string, supertype string) bool {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	c_supertype := (*C.gchar)(C.CString(supertype))
	defer C.free(unsafe.Pointer(c_supertype))

	ret := C.g_content_type_is_a(c_type_, c_supertype)

	return toGoBool(ret)
}

func Fn_g_content_type_is_unknown(type_ string) bool {
	c_type_ := (*C.gchar)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	ret := C.g_content_type_is_unknown(c_type_)

	return toGoBool(ret)
}

func Fn_g_content_types_get_registered() unsafe.Pointer {
	ret := C.g_content_types_get_registered()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_address_get_for_bus_sync(busType int, cancellable unsafe.Pointer, error unsafe.Pointer) string {
	c_busType := (C.GBusType)(busType)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_address_get_for_bus_sync(c_busType, c_cancellable, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : parameter 'out_guid' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_address_get_stream_sync : parameter 'out_guid' is non array with indirect count > 1

func Fn_g_dbus_generate_guid() string {
	ret := C.g_dbus_generate_guid()

	return C.GoString(ret)
}

func Fn_g_dbus_is_address(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_address(c_string_)

	return toGoBool(ret)
}

func Fn_g_dbus_is_guid(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_guid(c_string_)

	return toGoBool(ret)
}

func Fn_g_dbus_is_interface_name(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_interface_name(c_string_)

	return toGoBool(ret)
}

func Fn_g_dbus_is_member_name(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_member_name(c_string_)

	return toGoBool(ret)
}

func Fn_g_dbus_is_name(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_name(c_string_)

	return toGoBool(ret)
}

func Fn_g_dbus_is_supported_address(string_ string, error unsafe.Pointer) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	cError := (**C.GError)(error)

	ret := C.g_dbus_is_supported_address(c_string_, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_is_unique_name(string_ string) bool {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_dbus_is_unique_name(c_string_)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_new_tmp : parameter 'iostream' is non array with indirect count > 1

func Fn_g_io_error_from_errno(errNo int) int {
	c_errNo := (C.gint)(errNo)

	ret := C.g_io_error_from_errno(c_errNo)

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

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

func Fn_g_unix_is_mount_path_system_internal(mountPath string) bool {
	c_mountPath := (*C.char)(C.CString(mountPath))
	defer C.free(unsafe.Pointer(c_mountPath))

	ret := C.g_unix_is_mount_path_system_internal(c_mountPath)

	return toGoBool(ret)
}

func Fn_g_unix_mount_at(mountPath string, timeRead *uint64) unsafe.Pointer {
	c_mountPath := (*C.char)(C.CString(mountPath))
	defer C.free(unsafe.Pointer(c_mountPath))

	c_timeRead := (*C.guint64)(unsafe.Pointer(timeRead))

	ret := C.g_unix_mount_at(c_mountPath, c_timeRead)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_compare(mount1 unsafe.Pointer, mount2 unsafe.Pointer) int {
	c_mount1 := (*C.GUnixMountEntry)(unsafe.Pointer(mount1))

	c_mount2 := (*C.GUnixMountEntry)(unsafe.Pointer(mount2))

	ret := C.g_unix_mount_compare(c_mount1, c_mount2)

	return (int)(ret)
}

func Fn_g_unix_mount_free(mountEntry unsafe.Pointer) {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	C.g_unix_mount_free(c_mountEntry)
}

func Fn_g_unix_mount_get_device_path(mountEntry unsafe.Pointer) string {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_get_device_path(c_mountEntry)

	return C.GoString(ret)
}

func Fn_g_unix_mount_get_fs_type(mountEntry unsafe.Pointer) string {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_get_fs_type(c_mountEntry)

	return C.GoString(ret)
}

func Fn_g_unix_mount_get_mount_path(mountEntry unsafe.Pointer) string {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_get_mount_path(c_mountEntry)

	return C.GoString(ret)
}

func Fn_g_unix_mount_guess_can_eject(mountEntry unsafe.Pointer) bool {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_guess_can_eject(c_mountEntry)

	return toGoBool(ret)
}

func Fn_g_unix_mount_guess_icon(mountEntry unsafe.Pointer) unsafe.Pointer {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_guess_icon(c_mountEntry)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_guess_name(mountEntry unsafe.Pointer) string {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_guess_name(c_mountEntry)

	return C.GoString(ret)
}

func Fn_g_unix_mount_guess_should_display(mountEntry unsafe.Pointer) bool {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_guess_should_display(c_mountEntry)

	return toGoBool(ret)
}

func Fn_g_unix_mount_is_readonly(mountEntry unsafe.Pointer) bool {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_is_readonly(c_mountEntry)

	return toGoBool(ret)
}

func Fn_g_unix_mount_is_system_internal(mountEntry unsafe.Pointer) bool {
	c_mountEntry := (*C.GUnixMountEntry)(unsafe.Pointer(mountEntry))

	ret := C.g_unix_mount_is_system_internal(c_mountEntry)

	return toGoBool(ret)
}

func Fn_g_unix_mount_points_changed_since(time uint64) bool {
	c_time := (C.guint64)(time)

	ret := C.g_unix_mount_points_changed_since(c_time)

	return toGoBool(ret)
}

func Fn_g_unix_mount_points_get(timeRead *uint64) unsafe.Pointer {
	c_timeRead := (*C.guint64)(unsafe.Pointer(timeRead))

	ret := C.g_unix_mount_points_get(c_timeRead)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mounts_changed_since(time uint64) bool {
	c_time := (C.guint64)(time)

	ret := C.g_unix_mounts_changed_since(c_time)

	return toGoBool(ret)
}

func Fn_g_unix_mounts_get(timeRead *uint64) unsafe.Pointer {
	c_timeRead := (*C.guint64)(unsafe.Pointer(timeRead))

	ret := C.g_unix_mounts_get(c_timeRead)

	return unsafe.Pointer(ret)
}

func Fn_g_app_launch_context_new() unsafe.Pointer {
	ret := C.g_app_launch_context_new()

	return unsafe.Pointer(ret)
}

func Fn_g_app_launch_context_get_display(context unsafe.Pointer, info unsafe.Pointer, files unsafe.Pointer) string {
	c_context := (*C.GAppLaunchContext)(unsafe.Pointer(context))

	c_info := (*C.GAppInfo)(unsafe.Pointer(info))

	c_files := (*C.GList)(unsafe.Pointer(files))

	ret := C.g_app_launch_context_get_display(c_context, c_info, c_files)

	return C.GoString(ret)
}

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

func Fn_g_app_launch_context_get_startup_notify_id(context unsafe.Pointer, info unsafe.Pointer, files unsafe.Pointer) string {
	c_context := (*C.GAppLaunchContext)(unsafe.Pointer(context))

	c_info := (*C.GAppInfo)(unsafe.Pointer(info))

	c_files := (*C.GList)(unsafe.Pointer(files))

	ret := C.g_app_launch_context_get_startup_notify_id(c_context, c_info, c_files)

	return C.GoString(ret)
}

func Fn_g_app_launch_context_launch_failed(context unsafe.Pointer, startupNotifyId string) {
	c_context := (*C.GAppLaunchContext)(unsafe.Pointer(context))

	c_startupNotifyId := (*C.char)(C.CString(startupNotifyId))
	defer C.free(unsafe.Pointer(c_startupNotifyId))

	C.g_app_launch_context_launch_failed(c_context, c_startupNotifyId)
}

func Fn_g_application_new(applicationId *string, flags int) unsafe.Pointer {
	var c_applicationIdValue (*C.gchar)
	if applicationId != nil {
		c_applicationIdValue = (*C.gchar)(C.CString(*applicationId))
		defer C.free(unsafe.Pointer(c_applicationIdValue))
	}
	c_applicationId := c_applicationIdValue

	c_flags := (C.GApplicationFlags)(flags)

	ret := C.g_application_new(c_applicationId, c_flags)

	return unsafe.Pointer(ret)
}

func Fn_g_application_hold(application unsafe.Pointer) {
	c_application := (*C.GApplication)(unsafe.Pointer(application))

	C.g_application_hold(c_application)
}

func Fn_g_application_release(application unsafe.Pointer) {
	c_application := (*C.GApplication)(unsafe.Pointer(application))

	C.g_application_release(c_application)
}

func Fn_g_application_id_is_valid(applicationId string) bool {
	c_applicationId := (*C.gchar)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(c_applicationId))

	ret := C.g_application_id_is_valid(c_applicationId)

	return toGoBool(ret)
}

// UNSUPPORTED : g_application_command_line_get_environ : no array length

func Fn_g_buffered_input_stream_new(baseStream unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GInputStream)(unsafe.Pointer(baseStream))

	ret := C.g_buffered_input_stream_new(c_baseStream)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_input_stream_new_sized(baseStream unsafe.Pointer, size uint64) unsafe.Pointer {
	c_baseStream := (*C.GInputStream)(unsafe.Pointer(baseStream))

	c_size := (C.gsize)(size)

	ret := C.g_buffered_input_stream_new_sized(c_baseStream, c_size)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_input_stream_fill(stream unsafe.Pointer, count uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_count := (C.gssize)(count)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_fill(c_stream, c_count, c_cancellable, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

func Fn_g_buffered_input_stream_fill_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_fill_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_get_available(stream unsafe.Pointer) uint64 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	ret := C.g_buffered_input_stream_get_available(c_stream)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_get_buffer_size(stream unsafe.Pointer) uint64 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	ret := C.g_buffered_input_stream_get_buffer_size(c_stream)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_peek(stream unsafe.Pointer, buffer []uint8, offset uint64, count uint64) uint64 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_buffer := (unsafe.Pointer)(unsafe.Pointer(&buffer[0]))

	c_offset := (C.gsize)(offset)

	c_count := (C.gsize)(count)

	ret := C.g_buffered_input_stream_peek(c_stream, c_buffer, c_offset, c_count)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_peek_buffer(stream unsafe.Pointer, count *uint64) []uint8 {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_count := (*C.gsize)(unsafe.Pointer(count))

	ret := C.g_buffered_input_stream_peek_buffer(c_stream, c_count)

	retLen := int(*c_count)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_buffered_input_stream_read_byte(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) int {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_buffered_input_stream_read_byte(c_stream, c_cancellable, cError)

	return (int)(ret)
}

func Fn_g_buffered_input_stream_set_buffer_size(stream unsafe.Pointer, size uint64) {
	c_stream := (*C.GBufferedInputStream)(unsafe.Pointer(stream))

	c_size := (C.gsize)(size)

	C.g_buffered_input_stream_set_buffer_size(c_stream, c_size)
}

func Fn_g_buffered_output_stream_new(baseStream unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GOutputStream)(unsafe.Pointer(baseStream))

	ret := C.g_buffered_output_stream_new(c_baseStream)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_output_stream_new_sized(baseStream unsafe.Pointer, size uint64) unsafe.Pointer {
	c_baseStream := (*C.GOutputStream)(unsafe.Pointer(baseStream))

	c_size := (C.gsize)(size)

	ret := C.g_buffered_output_stream_new_sized(c_baseStream, c_size)

	return unsafe.Pointer(ret)
}

func Fn_g_buffered_output_stream_get_auto_grow(stream unsafe.Pointer) bool {
	c_stream := (*C.GBufferedOutputStream)(unsafe.Pointer(stream))

	ret := C.g_buffered_output_stream_get_auto_grow(c_stream)

	return toGoBool(ret)
}

func Fn_g_buffered_output_stream_get_buffer_size(stream unsafe.Pointer) uint64 {
	c_stream := (*C.GBufferedOutputStream)(unsafe.Pointer(stream))

	ret := C.g_buffered_output_stream_get_buffer_size(c_stream)

	return (uint64)(ret)
}

func Fn_g_buffered_output_stream_set_auto_grow(stream unsafe.Pointer, autoGrow bool) {
	c_stream := (*C.GBufferedOutputStream)(unsafe.Pointer(stream))

	c_autoGrow := toCBool(autoGrow)

	C.g_buffered_output_stream_set_auto_grow(c_stream, c_autoGrow)
}

func Fn_g_buffered_output_stream_set_buffer_size(stream unsafe.Pointer, size uint64) {
	c_stream := (*C.GBufferedOutputStream)(unsafe.Pointer(stream))

	c_size := (C.gsize)(size)

	C.g_buffered_output_stream_set_buffer_size(c_stream, c_size)
}

func Fn_g_cancellable_new() unsafe.Pointer {
	ret := C.g_cancellable_new()

	return unsafe.Pointer(ret)
}

func Fn_g_cancellable_cancel(cancellable unsafe.Pointer) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	C.g_cancellable_cancel(c_cancellable)
}

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

func Fn_g_cancellable_disconnect(cancellable unsafe.Pointer, handlerId uint64) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	c_handlerId := (C.gulong)(handlerId)

	C.g_cancellable_disconnect(c_cancellable, c_handlerId)
}

func Fn_g_cancellable_get_fd(cancellable unsafe.Pointer) int {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	ret := C.g_cancellable_get_fd(c_cancellable)

	return (int)(ret)
}

func Fn_g_cancellable_is_cancelled(cancellable unsafe.Pointer) bool {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	ret := C.g_cancellable_is_cancelled(c_cancellable)

	return toGoBool(ret)
}

func Fn_g_cancellable_make_pollfd(cancellable unsafe.Pointer, pollfd unsafe.Pointer) bool {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	c_pollfd := (*C.GPollFD)(unsafe.Pointer(pollfd))

	ret := C.g_cancellable_make_pollfd(c_cancellable, c_pollfd)

	return toGoBool(ret)
}

func Fn_g_cancellable_pop_current(cancellable unsafe.Pointer) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	C.g_cancellable_pop_current(c_cancellable)
}

func Fn_g_cancellable_push_current(cancellable unsafe.Pointer) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	C.g_cancellable_push_current(c_cancellable)
}

func Fn_g_cancellable_release_fd(cancellable unsafe.Pointer) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	C.g_cancellable_release_fd(c_cancellable)
}

func Fn_g_cancellable_reset(cancellable unsafe.Pointer) {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	C.g_cancellable_reset(c_cancellable)
}

func Fn_g_cancellable_set_error_if_cancelled(cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_cancellable_set_error_if_cancelled(c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_cancellable_get_current() unsafe.Pointer {
	ret := C.g_cancellable_get_current()

	return unsafe.Pointer(ret)
}

func Fn_g_charset_converter_new(toCharset string, fromCharset string, error unsafe.Pointer) unsafe.Pointer {
	c_toCharset := (*C.gchar)(C.CString(toCharset))
	defer C.free(unsafe.Pointer(c_toCharset))

	c_fromCharset := (*C.gchar)(C.CString(fromCharset))
	defer C.free(unsafe.Pointer(c_fromCharset))

	cError := (**C.GError)(error)

	ret := C.g_charset_converter_new(c_toCharset, c_fromCharset, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_charset_converter_get_num_fallbacks(converter unsafe.Pointer) uint {
	c_converter := (*C.GCharsetConverter)(unsafe.Pointer(converter))

	ret := C.g_charset_converter_get_num_fallbacks(c_converter)

	return (uint)(ret)
}

func Fn_g_charset_converter_get_use_fallback(converter unsafe.Pointer) bool {
	c_converter := (*C.GCharsetConverter)(unsafe.Pointer(converter))

	ret := C.g_charset_converter_get_use_fallback(c_converter)

	return toGoBool(ret)
}

func Fn_g_charset_converter_set_use_fallback(converter unsafe.Pointer, useFallback bool) {
	c_converter := (*C.GCharsetConverter)(unsafe.Pointer(converter))

	c_useFallback := toCBool(useFallback)

	C.g_charset_converter_set_use_fallback(c_converter, c_useFallback)
}

func Fn_g_converter_input_stream_new(baseStream unsafe.Pointer, converter unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GInputStream)(unsafe.Pointer(baseStream))

	c_converter := (*C.GConverter)(unsafe.Pointer(converter))

	ret := C.g_converter_input_stream_new(c_baseStream, c_converter)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_input_stream_get_converter(converterStream unsafe.Pointer) unsafe.Pointer {
	c_converterStream := (*C.GConverterInputStream)(unsafe.Pointer(converterStream))

	ret := C.g_converter_input_stream_get_converter(c_converterStream)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_output_stream_new(baseStream unsafe.Pointer, converter unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GOutputStream)(unsafe.Pointer(baseStream))

	c_converter := (*C.GConverter)(unsafe.Pointer(converter))

	ret := C.g_converter_output_stream_new(c_baseStream, c_converter)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_output_stream_get_converter(converterStream unsafe.Pointer) unsafe.Pointer {
	c_converterStream := (*C.GConverterOutputStream)(unsafe.Pointer(converterStream))

	ret := C.g_converter_output_stream_get_converter(c_converterStream)

	return unsafe.Pointer(ret)
}

func Fn_g_credentials_new() unsafe.Pointer {
	ret := C.g_credentials_new()

	return unsafe.Pointer(ret)
}

func Fn_g_credentials_get_native(credentials unsafe.Pointer, nativeType int) unsafe.Pointer {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	c_nativeType := (C.GCredentialsType)(nativeType)

	ret := C.g_credentials_get_native(c_credentials, c_nativeType)

	return (unsafe.Pointer)(ret)
}

func Fn_g_credentials_get_unix_user(credentials unsafe.Pointer, error unsafe.Pointer) uint {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	cError := (**C.GError)(error)

	ret := C.g_credentials_get_unix_user(c_credentials, cError)

	return (uint)(ret)
}

func Fn_g_credentials_is_same_user(credentials unsafe.Pointer, otherCredentials unsafe.Pointer, error unsafe.Pointer) bool {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	c_otherCredentials := (*C.GCredentials)(unsafe.Pointer(otherCredentials))

	cError := (**C.GError)(error)

	ret := C.g_credentials_is_same_user(c_credentials, c_otherCredentials, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_set_native(credentials unsafe.Pointer, nativeType int, native unsafe.Pointer) {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	c_nativeType := (C.GCredentialsType)(nativeType)

	c_native := (C.gpointer)(native)

	C.g_credentials_set_native(c_credentials, c_nativeType, c_native)
}

func Fn_g_credentials_set_unix_user(credentials unsafe.Pointer, uid uint, error unsafe.Pointer) bool {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	c_uid := (C.uid_t)(uid)

	cError := (**C.GError)(error)

	ret := C.g_credentials_set_unix_user(c_credentials, c_uid, cError)

	return toGoBool(ret)
}

func Fn_g_credentials_to_string(credentials unsafe.Pointer) string {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	ret := C.g_credentials_to_string(c_credentials)

	return C.GoString(ret)
}

func Fn_g_dbus_auth_observer_new() unsafe.Pointer {
	ret := C.g_dbus_auth_observer_new()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_auth_observer_authorize_authenticated_peer(observer unsafe.Pointer, stream unsafe.Pointer, credentials unsafe.Pointer) bool {
	c_observer := (*C.GDBusAuthObserver)(unsafe.Pointer(observer))

	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	ret := C.g_dbus_auth_observer_authorize_authenticated_peer(c_observer, c_stream, c_credentials)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_new_finish(res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_finish(c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_finish(res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_finish(c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_for_address_sync(address string, flags int, observer unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_address := (*C.gchar)(C.CString(address))
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(unsafe.Pointer(observer))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_for_address_sync(c_address, c_flags, c_observer, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_new_sync(stream unsafe.Pointer, guid *string, flags int, observer unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	var c_guidValue (*C.gchar)
	if guid != nil {
		c_guidValue = (*C.gchar)(C.CString(*guid))
		defer C.free(unsafe.Pointer(c_guidValue))
	}
	c_guid := c_guidValue

	c_flags := (C.GDBusConnectionFlags)(flags)

	c_observer := (*C.GDBusAuthObserver)(unsafe.Pointer(observer))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_new_sync(c_stream, c_guid, c_flags, c_observer, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

func Fn_g_dbus_connection_call_finish(connection unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_finish(c_connection, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_call_sync(connection unsafe.Pointer, busName *string, objectPath string, interfaceName string, methodName string, parameters unsafe.Pointer, replyType unsafe.Pointer, flags int, timeoutMsec int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	var c_busNameValue (*C.gchar)
	if busName != nil {
		c_busNameValue = (*C.gchar)(C.CString(*busName))
		defer C.free(unsafe.Pointer(c_busNameValue))
	}
	c_busName := c_busNameValue

	c_objectPath := (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(c_objectPath))

	c_interfaceName := (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(c_interfaceName))

	c_methodName := (*C.gchar)(C.CString(methodName))
	defer C.free(unsafe.Pointer(c_methodName))

	c_parameters := (*C.GVariant)(unsafe.Pointer(parameters))

	c_replyType := (*C.GVariantType)(unsafe.Pointer(replyType))

	c_flags := (C.GDBusCallFlags)(flags)

	c_timeoutMsec := (C.gint)(timeoutMsec)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_call_sync(c_connection, c_busName, c_objectPath, c_interfaceName, c_methodName, c_parameters, c_replyType, c_flags, c_timeoutMsec, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list_finish : parameter 'out_fd_list' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list_sync : parameter 'out_fd_list' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

func Fn_g_dbus_connection_close_finish(connection unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_finish(c_connection, c_res, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_close_sync(connection unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_close_sync(c_connection, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_emit_signal(connection unsafe.Pointer, destinationBusName *string, objectPath string, interfaceName string, signalName string, parameters unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	var c_destinationBusNameValue (*C.gchar)
	if destinationBusName != nil {
		c_destinationBusNameValue = (*C.gchar)(C.CString(*destinationBusName))
		defer C.free(unsafe.Pointer(c_destinationBusNameValue))
	}
	c_destinationBusName := c_destinationBusNameValue

	c_objectPath := (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(c_objectPath))

	c_interfaceName := (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(c_interfaceName))

	c_signalName := (*C.gchar)(C.CString(signalName))
	defer C.free(unsafe.Pointer(c_signalName))

	c_parameters := (*C.GVariant)(unsafe.Pointer(parameters))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_emit_signal(c_connection, c_destinationBusName, c_objectPath, c_interfaceName, c_signalName, c_parameters, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

func Fn_g_dbus_connection_flush_finish(connection unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_finish(c_connection, c_res, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_flush_sync(connection unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_flush_sync(c_connection, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_capabilities(connection unsafe.Pointer) int {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_capabilities(c_connection)

	return (int)(ret)
}

func Fn_g_dbus_connection_get_exit_on_close(connection unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_exit_on_close(c_connection)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_get_guid(connection unsafe.Pointer) string {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_guid(c_connection)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_get_peer_credentials(connection unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_peer_credentials(c_connection)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_stream(connection unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_stream(c_connection)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_get_unique_name(connection unsafe.Pointer) string {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_get_unique_name(c_connection)

	return C.GoString(ret)
}

func Fn_g_dbus_connection_is_closed(connection unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	ret := C.g_dbus_connection_is_closed(c_connection)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

func Fn_g_dbus_connection_remove_filter(connection unsafe.Pointer, filterId uint) {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_filterId := (C.guint)(filterId)

	C.g_dbus_connection_remove_filter(c_connection, c_filterId)
}

func Fn_g_dbus_connection_send_message(connection unsafe.Pointer, message unsafe.Pointer, flags int, outSerial *uint32, error unsafe.Pointer) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_flags := (C.GDBusSendMessageFlags)(flags)

	c_outSerial := (*C.guint32)(unsafe.Pointer(outSerial))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message(c_connection, c_message, c_flags, c_outSerial, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

func Fn_g_dbus_connection_send_message_with_reply_finish(connection unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_finish(c_connection, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_send_message_with_reply_sync(connection unsafe.Pointer, message unsafe.Pointer, flags int, timeoutMsec int, outSerial *uint32, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_flags := (C.GDBusSendMessageFlags)(flags)

	c_timeoutMsec := (C.gint)(timeoutMsec)

	c_outSerial := (*C.guint32)(unsafe.Pointer(outSerial))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_connection_send_message_with_reply_sync(c_connection, c_message, c_flags, c_timeoutMsec, c_outSerial, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_connection_set_exit_on_close(connection unsafe.Pointer, exitOnClose bool) {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_exitOnClose := toCBool(exitOnClose)

	C.g_dbus_connection_set_exit_on_close(c_connection, c_exitOnClose)
}

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

func Fn_g_dbus_connection_signal_unsubscribe(connection unsafe.Pointer, subscriptionId uint) {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_subscriptionId := (C.guint)(subscriptionId)

	C.g_dbus_connection_signal_unsubscribe(c_connection, c_subscriptionId)
}

func Fn_g_dbus_connection_start_message_processing(connection unsafe.Pointer) {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	C.g_dbus_connection_start_message_processing(c_connection)
}

func Fn_g_dbus_connection_unregister_object(connection unsafe.Pointer, registrationId uint) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_registrationId := (C.guint)(registrationId)

	ret := C.g_dbus_connection_unregister_object(c_connection, c_registrationId)

	return toGoBool(ret)
}

func Fn_g_dbus_connection_unregister_subtree(connection unsafe.Pointer, registrationId uint) bool {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_registrationId := (C.guint)(registrationId)

	ret := C.g_dbus_connection_unregister_subtree(c_connection, c_registrationId)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

func Fn_g_dbus_message_new() unsafe.Pointer {
	ret := C.g_dbus_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_from_blob(blob []uint8, blobLen uint64, capabilities int, error unsafe.Pointer) unsafe.Pointer {
	c_blob := (*C.guchar)(unsafe.Pointer(&blob[0]))

	c_blobLen := (C.gsize)(blobLen)

	c_capabilities := (C.GDBusCapabilityFlags)(capabilities)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_new_from_blob(c_blob, c_blobLen, c_capabilities, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_call(name *string, path string, interface_ *string, method string) unsafe.Pointer {
	var c_nameValue (*C.gchar)
	if name != nil {
		c_nameValue = (*C.gchar)(C.CString(*name))
		defer C.free(unsafe.Pointer(c_nameValue))
	}
	c_name := c_nameValue

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	var c_interface_Value (*C.gchar)
	if interface_ != nil {
		c_interface_Value = (*C.gchar)(C.CString(*interface_))
		defer C.free(unsafe.Pointer(c_interface_Value))
	}
	c_interface_ := c_interface_Value

	c_method := (*C.gchar)(C.CString(method))
	defer C.free(unsafe.Pointer(c_method))

	ret := C.g_dbus_message_new_method_call(c_name, c_path, c_interface_, c_method)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_signal(path string, interface_ string, signal string) unsafe.Pointer {
	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	c_interface_ := (*C.gchar)(C.CString(interface_))
	defer C.free(unsafe.Pointer(c_interface_))

	c_signal := (*C.gchar)(C.CString(signal))
	defer C.free(unsafe.Pointer(c_signal))

	ret := C.g_dbus_message_new_signal(c_path, c_interface_, c_signal)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_copy(message unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_copy(c_message, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_arg0(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_arg0(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_body(message unsafe.Pointer) unsafe.Pointer {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_body(c_message)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_get_byte_order(message unsafe.Pointer) int {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_byte_order(c_message)

	return (int)(ret)
}

func Fn_g_dbus_message_get_destination(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_destination(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_error_name(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_error_name(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_flags(message unsafe.Pointer) int {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_flags(c_message)

	return (int)(ret)
}

func Fn_g_dbus_message_get_header(message unsafe.Pointer, headerField int) unsafe.Pointer {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_headerField := (C.GDBusMessageHeaderField)(headerField)

	ret := C.g_dbus_message_get_header(c_message, c_headerField)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : no array length

func Fn_g_dbus_message_get_interface(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_interface(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_locked(message unsafe.Pointer) bool {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_locked(c_message)

	return toGoBool(ret)
}

func Fn_g_dbus_message_get_member(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_member(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_message_type(message unsafe.Pointer) int {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_message_type(c_message)

	return (int)(ret)
}

func Fn_g_dbus_message_get_num_unix_fds(message unsafe.Pointer) uint32 {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_num_unix_fds(c_message)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_path(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_path(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_reply_serial(message unsafe.Pointer) uint32 {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_reply_serial(c_message)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_sender(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_sender(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_serial(message unsafe.Pointer) uint32 {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_serial(c_message)

	return (uint32)(ret)
}

func Fn_g_dbus_message_get_signature(message unsafe.Pointer) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_signature(c_message)

	return C.GoString(ret)
}

func Fn_g_dbus_message_get_unix_fd_list(message unsafe.Pointer) unsafe.Pointer {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	ret := C.g_dbus_message_get_unix_fd_list(c_message)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_lock(message unsafe.Pointer) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	C.g_dbus_message_lock(c_message)
}

func Fn_g_dbus_message_new_method_error(methodCallMessage unsafe.Pointer, errorName string, errorMessageFormat string) unsafe.Pointer {
	c_methodCallMessage := (*C.GDBusMessage)(unsafe.Pointer(methodCallMessage))

	c_errorName := (*C.gchar)(C.CString(errorName))
	defer C.free(unsafe.Pointer(c_errorName))

	c_errorMessageFormat := (*C.gchar)(C.CString(errorMessageFormat))
	defer C.free(unsafe.Pointer(c_errorMessageFormat))

	ret := C.c_g_dbus_message_new_method_error(c_methodCallMessage, c_errorName, c_errorMessageFormat)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_error_literal(methodCallMessage unsafe.Pointer, errorName string, errorMessage string) unsafe.Pointer {
	c_methodCallMessage := (*C.GDBusMessage)(unsafe.Pointer(methodCallMessage))

	c_errorName := (*C.gchar)(C.CString(errorName))
	defer C.free(unsafe.Pointer(c_errorName))

	c_errorMessage := (*C.gchar)(C.CString(errorMessage))
	defer C.free(unsafe.Pointer(c_errorMessage))

	ret := C.g_dbus_message_new_method_error_literal(c_methodCallMessage, c_errorName, c_errorMessage)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_error_valist(methodCallMessage unsafe.Pointer, errorName string, errorMessageFormat string) unsafe.Pointer {
	c_methodCallMessage := (*C.GDBusMessage)(unsafe.Pointer(methodCallMessage))

	c_errorName := (*C.gchar)(C.CString(errorName))
	defer C.free(unsafe.Pointer(c_errorName))

	c_errorMessageFormat := (*C.gchar)(C.CString(errorMessageFormat))
	defer C.free(unsafe.Pointer(c_errorMessageFormat))

	ret := C.c_g_dbus_message_new_method_error_valist(c_methodCallMessage, c_errorName, c_errorMessageFormat)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_new_method_reply(methodCallMessage unsafe.Pointer) unsafe.Pointer {
	c_methodCallMessage := (*C.GDBusMessage)(unsafe.Pointer(methodCallMessage))

	ret := C.g_dbus_message_new_method_reply(c_methodCallMessage)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_message_print(message unsafe.Pointer, indent uint) string {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_indent := (C.guint)(indent)

	ret := C.g_dbus_message_print(c_message, c_indent)

	return C.GoString(ret)
}

func Fn_g_dbus_message_set_body(message unsafe.Pointer, body unsafe.Pointer) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_body := (*C.GVariant)(unsafe.Pointer(body))

	C.g_dbus_message_set_body(c_message, c_body)
}

func Fn_g_dbus_message_set_byte_order(message unsafe.Pointer, byteOrder int) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_byteOrder := (C.GDBusMessageByteOrder)(byteOrder)

	C.g_dbus_message_set_byte_order(c_message, c_byteOrder)
}

func Fn_g_dbus_message_set_destination(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_destination(c_message, c_value)
}

func Fn_g_dbus_message_set_error_name(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_error_name(c_message, c_value)
}

func Fn_g_dbus_message_set_flags(message unsafe.Pointer, flags int) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_flags := (C.GDBusMessageFlags)(flags)

	C.g_dbus_message_set_flags(c_message, c_flags)
}

func Fn_g_dbus_message_set_header(message unsafe.Pointer, headerField int, value unsafe.Pointer) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_headerField := (C.GDBusMessageHeaderField)(headerField)

	c_value := (*C.GVariant)(unsafe.Pointer(value))

	C.g_dbus_message_set_header(c_message, c_headerField, c_value)
}

func Fn_g_dbus_message_set_interface(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_interface(c_message, c_value)
}

func Fn_g_dbus_message_set_member(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_member(c_message, c_value)
}

func Fn_g_dbus_message_set_message_type(message unsafe.Pointer, type_ int) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_type_ := (C.GDBusMessageType)(type_)

	C.g_dbus_message_set_message_type(c_message, c_type_)
}

func Fn_g_dbus_message_set_num_unix_fds(message unsafe.Pointer, value uint32) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (C.guint32)(value)

	C.g_dbus_message_set_num_unix_fds(c_message, c_value)
}

func Fn_g_dbus_message_set_path(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_path(c_message, c_value)
}

func Fn_g_dbus_message_set_reply_serial(message unsafe.Pointer, value uint32) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (C.guint32)(value)

	C.g_dbus_message_set_reply_serial(c_message, c_value)
}

func Fn_g_dbus_message_set_sender(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_sender(c_message, c_value)
}

func Fn_g_dbus_message_set_serial(message unsafe.Pointer, serial uint32) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_serial := (C.guint32)(serial)

	C.g_dbus_message_set_serial(c_message, c_serial)
}

func Fn_g_dbus_message_set_signature(message unsafe.Pointer, value string) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	C.g_dbus_message_set_signature(c_message, c_value)
}

func Fn_g_dbus_message_set_unix_fd_list(message unsafe.Pointer, fdList unsafe.Pointer) {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_fdList := (*C.GUnixFDList)(unsafe.Pointer(fdList))

	C.g_dbus_message_set_unix_fd_list(c_message, c_fdList)
}

func Fn_g_dbus_message_to_blob(message unsafe.Pointer, outSize *uint64, capabilities int, error unsafe.Pointer) []uint8 {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	c_outSize := (*C.gsize)(unsafe.Pointer(outSize))

	c_capabilities := (C.GDBusCapabilityFlags)(capabilities)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_to_blob(c_message, c_outSize, c_capabilities, cError)

	retLen := int(*c_outSize)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_dbus_message_to_gerror(message unsafe.Pointer, error unsafe.Pointer) bool {
	c_message := (*C.GDBusMessage)(unsafe.Pointer(message))

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_to_gerror(c_message, cError)

	return toGoBool(ret)
}

func Fn_g_dbus_message_bytes_needed(blob []uint8, blobLen uint64, error unsafe.Pointer) uint64 {
	c_blob := (*C.guchar)(unsafe.Pointer(&blob[0]))

	c_blobLen := (C.gsize)(blobLen)

	cError := (**C.GError)(error)

	ret := C.g_dbus_message_bytes_needed(c_blob, c_blobLen, cError)

	return (uint64)(ret)
}

func Fn_g_dbus_method_invocation_get_connection(invocation unsafe.Pointer) unsafe.Pointer {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_connection(c_invocation)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_interface_name(invocation unsafe.Pointer) string {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_interface_name(c_invocation)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_message(invocation unsafe.Pointer) unsafe.Pointer {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_message(c_invocation)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_info(invocation unsafe.Pointer) unsafe.Pointer {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_method_info(c_invocation)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_method_name(invocation unsafe.Pointer) string {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_method_name(c_invocation)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_object_path(invocation unsafe.Pointer) string {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_object_path(c_invocation)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_parameters(invocation unsafe.Pointer) unsafe.Pointer {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_parameters(c_invocation)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_method_invocation_get_sender(invocation unsafe.Pointer) string {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_sender(c_invocation)

	return C.GoString(ret)
}

func Fn_g_dbus_method_invocation_get_user_data(invocation unsafe.Pointer) unsafe.Pointer {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	ret := C.g_dbus_method_invocation_get_user_data(c_invocation)

	return (unsafe.Pointer)(ret)
}

func Fn_g_dbus_method_invocation_return_dbus_error(invocation unsafe.Pointer, errorName string, errorMessage string) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_errorName := (*C.gchar)(C.CString(errorName))
	defer C.free(unsafe.Pointer(c_errorName))

	c_errorMessage := (*C.gchar)(C.CString(errorMessage))
	defer C.free(unsafe.Pointer(c_errorMessage))

	C.g_dbus_method_invocation_return_dbus_error(c_invocation, c_errorName, c_errorMessage)
}

func Fn_g_dbus_method_invocation_return_error(invocation unsafe.Pointer, domain uint32, code int, format string) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_format := (*C.gchar)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	C.c_g_dbus_method_invocation_return_error(c_invocation, c_domain, c_code, c_format)
}

func Fn_g_dbus_method_invocation_return_error_literal(invocation unsafe.Pointer, domain uint32, code int, message string) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := (*C.gchar)(C.CString(message))
	defer C.free(unsafe.Pointer(c_message))

	C.g_dbus_method_invocation_return_error_literal(c_invocation, c_domain, c_code, c_message)
}

func Fn_g_dbus_method_invocation_return_error_valist(invocation unsafe.Pointer, domain uint32, code int, format string) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_format := (*C.gchar)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	C.c_g_dbus_method_invocation_return_error_valist(c_invocation, c_domain, c_code, c_format)
}

func Fn_g_dbus_method_invocation_return_gerror(invocation unsafe.Pointer, error unsafe.Pointer) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_error := (*C.GError)(unsafe.Pointer(error))

	C.g_dbus_method_invocation_return_gerror(c_invocation, c_error)
}

func Fn_g_dbus_method_invocation_return_value(invocation unsafe.Pointer, parameters unsafe.Pointer) {
	c_invocation := (*C.GDBusMethodInvocation)(unsafe.Pointer(invocation))

	c_parameters := (*C.GVariant)(unsafe.Pointer(parameters))

	C.g_dbus_method_invocation_return_value(c_invocation, c_parameters)
}

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : parameter 'get_proxy_type_func' is callback

func Fn_g_dbus_object_manager_server_set_connection(manager unsafe.Pointer, connection unsafe.Pointer) {
	c_manager := (*C.GDBusObjectManagerServer)(unsafe.Pointer(manager))

	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	C.g_dbus_object_manager_server_set_connection(c_manager, c_connection)
}

func Fn_g_dbus_proxy_new_finish(res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_finish(c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_finish(res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_for_bus_finish(c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_for_bus_sync(busType int, flags int, info unsafe.Pointer, name string, objectPath string, interfaceName string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_busType := (C.GBusType)(busType)

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_objectPath := (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(c_objectPath))

	c_interfaceName := (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(c_interfaceName))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_for_bus_sync(c_busType, c_flags, c_info, c_name, c_objectPath, c_interfaceName, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_new_sync(connection unsafe.Pointer, flags int, info unsafe.Pointer, name *string, objectPath string, interfaceName string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GDBusConnection)(unsafe.Pointer(connection))

	c_flags := (C.GDBusProxyFlags)(flags)

	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	var c_nameValue (*C.gchar)
	if name != nil {
		c_nameValue = (*C.gchar)(C.CString(*name))
		defer C.free(unsafe.Pointer(c_nameValue))
	}
	c_name := c_nameValue

	c_objectPath := (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(c_objectPath))

	c_interfaceName := (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(c_interfaceName))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_new_sync(c_connection, c_flags, c_info, c_name, c_objectPath, c_interfaceName, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call : parameter 'callback' is callback

func Fn_g_dbus_proxy_call_finish(proxy unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_finish(c_proxy, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_call_sync(proxy unsafe.Pointer, methodName string, parameters unsafe.Pointer, flags int, timeoutMsec int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_methodName := (*C.gchar)(C.CString(methodName))
	defer C.free(unsafe.Pointer(c_methodName))

	c_parameters := (*C.GVariant)(unsafe.Pointer(parameters))

	c_flags := (C.GDBusCallFlags)(flags)

	c_timeoutMsec := (C.gint)(timeoutMsec)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_proxy_call_sync(c_proxy, c_methodName, c_parameters, c_flags, c_timeoutMsec, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list_finish : parameter 'out_fd_list' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list_sync : parameter 'out_fd_list' is non array with indirect count > 1

func Fn_g_dbus_proxy_get_cached_property(proxy unsafe.Pointer, propertyName string) unsafe.Pointer {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	ret := C.g_dbus_proxy_get_cached_property(c_proxy, c_propertyName)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : no array length

func Fn_g_dbus_proxy_get_connection(proxy unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_connection(c_proxy)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_default_timeout(proxy unsafe.Pointer) int {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_default_timeout(c_proxy)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_flags(proxy unsafe.Pointer) int {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_flags(c_proxy)

	return (int)(ret)
}

func Fn_g_dbus_proxy_get_interface_info(proxy unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_interface_info(c_proxy)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_proxy_get_interface_name(proxy unsafe.Pointer) string {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_interface_name(c_proxy)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name(proxy unsafe.Pointer) string {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_name(c_proxy)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_name_owner(proxy unsafe.Pointer) string {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_name_owner(c_proxy)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_get_object_path(proxy unsafe.Pointer) string {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	ret := C.g_dbus_proxy_get_object_path(c_proxy)

	return C.GoString(ret)
}

func Fn_g_dbus_proxy_set_cached_property(proxy unsafe.Pointer, propertyName string, value unsafe.Pointer) {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	c_value := (*C.GVariant)(unsafe.Pointer(value))

	C.g_dbus_proxy_set_cached_property(c_proxy, c_propertyName, c_value)
}

func Fn_g_dbus_proxy_set_default_timeout(proxy unsafe.Pointer, timeoutMsec int) {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_timeoutMsec := (C.gint)(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout(c_proxy, c_timeoutMsec)
}

func Fn_g_dbus_proxy_set_interface_info(proxy unsafe.Pointer, info unsafe.Pointer) {
	c_proxy := (*C.GDBusProxy)(unsafe.Pointer(proxy))

	c_info := (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_proxy_set_interface_info(c_proxy, c_info)
}

// UNSUPPORTED : g_dbus_proxy_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : parameter 'callback' is callback

func Fn_g_dbus_server_new_sync(address string, flags int, guid string, observer unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_address := (*C.gchar)(C.CString(address))
	defer C.free(unsafe.Pointer(c_address))

	c_flags := (C.GDBusServerFlags)(flags)

	c_guid := (*C.gchar)(C.CString(guid))
	defer C.free(unsafe.Pointer(c_guid))

	c_observer := (*C.GDBusAuthObserver)(unsafe.Pointer(observer))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_dbus_server_new_sync(c_address, c_flags, c_guid, c_observer, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_dbus_server_get_client_address(server unsafe.Pointer) string {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	ret := C.g_dbus_server_get_client_address(c_server)

	return C.GoString(ret)
}

func Fn_g_dbus_server_get_flags(server unsafe.Pointer) int {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	ret := C.g_dbus_server_get_flags(c_server)

	return (int)(ret)
}

func Fn_g_dbus_server_get_guid(server unsafe.Pointer) string {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	ret := C.g_dbus_server_get_guid(c_server)

	return C.GoString(ret)
}

func Fn_g_dbus_server_is_active(server unsafe.Pointer) bool {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	ret := C.g_dbus_server_is_active(c_server)

	return toGoBool(ret)
}

func Fn_g_dbus_server_start(server unsafe.Pointer) {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	C.g_dbus_server_start(c_server)
}

func Fn_g_dbus_server_stop(server unsafe.Pointer) {
	c_server := (*C.GDBusServer)(unsafe.Pointer(server))

	C.g_dbus_server_stop(c_server)
}

func Fn_g_data_input_stream_new(baseStream unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GInputStream)(unsafe.Pointer(baseStream))

	ret := C.g_data_input_stream_new(c_baseStream)

	return unsafe.Pointer(ret)
}

func Fn_g_data_input_stream_get_byte_order(stream unsafe.Pointer) int {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	ret := C.g_data_input_stream_get_byte_order(c_stream)

	return (int)(ret)
}

func Fn_g_data_input_stream_get_newline_type(stream unsafe.Pointer) int {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	ret := C.g_data_input_stream_get_newline_type(c_stream)

	return (int)(ret)
}

func Fn_g_data_input_stream_read_byte(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) uint8 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_byte(c_stream, c_cancellable, cError)

	return (uint8)(ret)
}

func Fn_g_data_input_stream_read_int16(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) int16 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int16(c_stream, c_cancellable, cError)

	return (int16)(ret)
}

func Fn_g_data_input_stream_read_int32(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) int32 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int32(c_stream, c_cancellable, cError)

	return (int32)(ret)
}

func Fn_g_data_input_stream_read_int64(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) int64 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_int64(c_stream, c_cancellable, cError)

	return (int64)(ret)
}

// UNSUPPORTED : g_data_input_stream_read_line : no array length

// UNSUPPORTED : g_data_input_stream_read_line_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : no array length

func Fn_g_data_input_stream_read_uint16(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) uint16 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint16(c_stream, c_cancellable, cError)

	return (uint16)(ret)
}

func Fn_g_data_input_stream_read_uint32(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) uint32 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint32(c_stream, c_cancellable, cError)

	return (uint32)(ret)
}

func Fn_g_data_input_stream_read_uint64(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_uint64(c_stream, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_data_input_stream_read_until(stream unsafe.Pointer, stopChars string, length *uint64, cancellable unsafe.Pointer, error unsafe.Pointer) string {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_stopChars := (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(c_stopChars))

	c_length := (*C.gsize)(unsafe.Pointer(length))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_until(c_stream, c_stopChars, c_length, c_cancellable, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_until_async : parameter 'callback' is callback

func Fn_g_data_input_stream_read_until_finish(stream unsafe.Pointer, result unsafe.Pointer, length *uint64, error unsafe.Pointer) string {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	c_length := (*C.gsize)(unsafe.Pointer(length))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_until_finish(c_stream, c_result, c_length, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_read_upto(stream unsafe.Pointer, stopChars string, stopCharsLen uint64, length *uint64, cancellable unsafe.Pointer, error unsafe.Pointer) string {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_stopChars := (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(c_stopChars))

	c_stopCharsLen := (C.gssize)(stopCharsLen)

	c_length := (*C.gsize)(unsafe.Pointer(length))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_upto(c_stream, c_stopChars, c_stopCharsLen, c_length, c_cancellable, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_data_input_stream_read_upto_async : parameter 'callback' is callback

func Fn_g_data_input_stream_read_upto_finish(stream unsafe.Pointer, result unsafe.Pointer, length *uint64, error unsafe.Pointer) string {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	c_length := (*C.gsize)(unsafe.Pointer(length))

	cError := (**C.GError)(error)

	ret := C.g_data_input_stream_read_upto_finish(c_stream, c_result, c_length, cError)

	return C.GoString(ret)
}

func Fn_g_data_input_stream_set_byte_order(stream unsafe.Pointer, order int) {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_input_stream_set_byte_order(c_stream, c_order)
}

func Fn_g_data_input_stream_set_newline_type(stream unsafe.Pointer, type_ int) {
	c_stream := (*C.GDataInputStream)(unsafe.Pointer(stream))

	c_type_ := (C.GDataStreamNewlineType)(type_)

	C.g_data_input_stream_set_newline_type(c_stream, c_type_)
}

func Fn_g_data_output_stream_new(baseStream unsafe.Pointer) unsafe.Pointer {
	c_baseStream := (*C.GOutputStream)(unsafe.Pointer(baseStream))

	ret := C.g_data_output_stream_new(c_baseStream)

	return unsafe.Pointer(ret)
}

func Fn_g_data_output_stream_get_byte_order(stream unsafe.Pointer) int {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	ret := C.g_data_output_stream_get_byte_order(c_stream)

	return (int)(ret)
}

func Fn_g_data_output_stream_put_byte(stream unsafe.Pointer, data uint8, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.guchar)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_byte(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int16(stream unsafe.Pointer, data int16, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.gint16)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int16(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int32(stream unsafe.Pointer, data int32, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.gint32)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int32(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_int64(stream unsafe.Pointer, data int64, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.gint64)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_int64(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_string(stream unsafe.Pointer, str string, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_str := (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_string(c_stream, c_str, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint16(stream unsafe.Pointer, data uint16, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.guint16)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint16(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint32(stream unsafe.Pointer, data uint32, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.guint32)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint32(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_put_uint64(stream unsafe.Pointer, data uint64, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_data := (C.guint64)(data)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_data_output_stream_put_uint64(c_stream, c_data, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_data_output_stream_set_byte_order(stream unsafe.Pointer, order int) {
	c_stream := (*C.GDataOutputStream)(unsafe.Pointer(stream))

	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_output_stream_set_byte_order(c_stream, c_order)
}

func Fn_g_desktop_app_info_new(desktopId string) unsafe.Pointer {
	c_desktopId := (*C.char)(C.CString(desktopId))
	defer C.free(unsafe.Pointer(c_desktopId))

	ret := C.g_desktop_app_info_new(c_desktopId)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_new_from_filename(filename string) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	ret := C.g_desktop_app_info_new_from_filename(c_filename)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_new_from_keyfile(keyFile unsafe.Pointer) unsafe.Pointer {
	c_keyFile := (*C.GKeyFile)(unsafe.Pointer(keyFile))

	ret := C.g_desktop_app_info_new_from_keyfile(c_keyFile)

	return unsafe.Pointer(ret)
}

func Fn_g_desktop_app_info_get_categories(info unsafe.Pointer) string {
	c_info := (*C.GDesktopAppInfo)(unsafe.Pointer(info))

	ret := C.g_desktop_app_info_get_categories(c_info)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_filename(info unsafe.Pointer) string {
	c_info := (*C.GDesktopAppInfo)(unsafe.Pointer(info))

	ret := C.g_desktop_app_info_get_filename(c_info)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_generic_name(info unsafe.Pointer) string {
	c_info := (*C.GDesktopAppInfo)(unsafe.Pointer(info))

	ret := C.g_desktop_app_info_get_generic_name(c_info)

	return C.GoString(ret)
}

func Fn_g_desktop_app_info_get_is_hidden(info unsafe.Pointer) bool {
	c_info := (*C.GDesktopAppInfo)(unsafe.Pointer(info))

	ret := C.g_desktop_app_info_get_is_hidden(c_info)

	return toGoBool(ret)
}

// UNSUPPORTED : g_desktop_app_info_get_keywords : no array length

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_list_actions : no array length

// UNSUPPORTED : g_desktop_app_info_search : array has no type

func Fn_g_desktop_app_info_set_desktop_env(desktopEnv string) {
	c_desktopEnv := (*C.char)(C.CString(desktopEnv))
	defer C.free(unsafe.Pointer(c_desktopEnv))

	C.g_desktop_app_info_set_desktop_env(c_desktopEnv)
}

func Fn_g_emblem_new(icon unsafe.Pointer) unsafe.Pointer {
	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	ret := C.g_emblem_new(c_icon)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_new_with_origin(icon unsafe.Pointer, origin int) unsafe.Pointer {
	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	c_origin := (C.GEmblemOrigin)(origin)

	ret := C.g_emblem_new_with_origin(c_icon, c_origin)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_get_icon(emblem unsafe.Pointer) unsafe.Pointer {
	c_emblem := (*C.GEmblem)(unsafe.Pointer(emblem))

	ret := C.g_emblem_get_icon(c_emblem)

	return unsafe.Pointer(ret)
}

func Fn_g_emblem_get_origin(emblem unsafe.Pointer) int {
	c_emblem := (*C.GEmblem)(unsafe.Pointer(emblem))

	ret := C.g_emblem_get_origin(c_emblem)

	return (int)(ret)
}

func Fn_g_emblemed_icon_new(icon unsafe.Pointer, emblem unsafe.Pointer) unsafe.Pointer {
	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	c_emblem := (*C.GEmblem)(unsafe.Pointer(emblem))

	ret := C.g_emblemed_icon_new(c_icon, c_emblem)

	return unsafe.Pointer(ret)
}

func Fn_g_emblemed_icon_add_emblem(emblemed unsafe.Pointer, emblem unsafe.Pointer) {
	c_emblemed := (*C.GEmblemedIcon)(unsafe.Pointer(emblemed))

	c_emblem := (*C.GEmblem)(unsafe.Pointer(emblem))

	C.g_emblemed_icon_add_emblem(c_emblemed, c_emblem)
}

func Fn_g_emblemed_icon_get_emblems(emblemed unsafe.Pointer) unsafe.Pointer {
	c_emblemed := (*C.GEmblemedIcon)(unsafe.Pointer(emblemed))

	ret := C.g_emblemed_icon_get_emblems(c_emblemed)

	return unsafe.Pointer(ret)
}

func Fn_g_emblemed_icon_get_icon(emblemed unsafe.Pointer) unsafe.Pointer {
	c_emblemed := (*C.GEmblemedIcon)(unsafe.Pointer(emblemed))

	ret := C.g_emblemed_icon_get_icon(c_emblemed)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_close(enumerator unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close(c_enumerator, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_enumerator_close_async : parameter 'callback' is callback

func Fn_g_file_enumerator_close_finish(enumerator unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close_finish(c_enumerator, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_get_container(enumerator unsafe.Pointer) unsafe.Pointer {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	ret := C.g_file_enumerator_get_container(c_enumerator)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_has_pending(enumerator unsafe.Pointer) bool {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	ret := C.g_file_enumerator_has_pending(c_enumerator)

	return toGoBool(ret)
}

func Fn_g_file_enumerator_is_closed(enumerator unsafe.Pointer) bool {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	ret := C.g_file_enumerator_is_closed(c_enumerator)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_enumerator_iterate : parameter 'out_info' is non array with indirect count > 1

func Fn_g_file_enumerator_next_file(enumerator unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_next_file(c_enumerator, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_enumerator_next_files_async : parameter 'callback' is callback

func Fn_g_file_enumerator_next_files_finish(enumerator unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_next_files_finish(c_enumerator, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_enumerator_set_pending(enumerator unsafe.Pointer, pending bool) {
	c_enumerator := (*C.GFileEnumerator)(unsafe.Pointer(enumerator))

	c_pending := toCBool(pending)

	C.g_file_enumerator_set_pending(c_enumerator, c_pending)
}

func Fn_g_file_io_stream_get_etag(stream unsafe.Pointer) string {
	c_stream := (*C.GFileIOStream)(unsafe.Pointer(stream))

	ret := C.g_file_io_stream_get_etag(c_stream)

	return C.GoString(ret)
}

func Fn_g_file_io_stream_query_info(stream unsafe.Pointer, attributes string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileIOStream)(unsafe.Pointer(stream))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info(c_stream, c_attributes, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_io_stream_query_info_async : parameter 'callback' is callback

func Fn_g_file_io_stream_query_info_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileIOStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_io_stream_query_info_finish(c_stream, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_icon_new(file unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_icon_new(c_file)

	return unsafe.Pointer(ret)
}

func Fn_g_file_icon_get_file(icon unsafe.Pointer) unsafe.Pointer {
	c_icon := (*C.GFileIcon)(unsafe.Pointer(icon))

	ret := C.g_file_icon_get_file(c_icon)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_new() unsafe.Pointer {
	ret := C.g_file_info_new()

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_clear_status(info unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	C.g_file_info_clear_status(c_info)
}

func Fn_g_file_info_copy_into(srcInfo unsafe.Pointer, destInfo unsafe.Pointer) {
	c_srcInfo := (*C.GFileInfo)(unsafe.Pointer(srcInfo))

	c_destInfo := (*C.GFileInfo)(unsafe.Pointer(destInfo))

	C.g_file_info_copy_into(c_srcInfo, c_destInfo)
}

func Fn_g_file_info_dup(other unsafe.Pointer) unsafe.Pointer {
	c_other := (*C.GFileInfo)(unsafe.Pointer(other))

	ret := C.g_file_info_dup(c_other)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_attribute_as_string(info unsafe.Pointer, attribute string) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_as_string(c_info, c_attribute)

	return C.GoString(ret)
}

func Fn_g_file_info_get_attribute_boolean(info unsafe.Pointer, attribute string) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_boolean(c_info, c_attribute)

	return toGoBool(ret)
}

func Fn_g_file_info_get_attribute_byte_string(info unsafe.Pointer, attribute string) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_byte_string(c_info, c_attribute)

	return C.GoString(ret)
}

func Fn_g_file_info_get_attribute_data(info unsafe.Pointer, attribute string, type_ *int, valuePp *unsafe.Pointer, status *int) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_type_ := (*C.GFileAttributeType)(unsafe.Pointer(type_))

	c_valuePp := (*C.gpointer)(unsafe.Pointer(valuePp))

	c_status := (*C.GFileAttributeStatus)(unsafe.Pointer(status))

	ret := C.g_file_info_get_attribute_data(c_info, c_attribute, c_type_, c_valuePp, c_status)

	return toGoBool(ret)
}

func Fn_g_file_info_get_attribute_int32(info unsafe.Pointer, attribute string) int32 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_int32(c_info, c_attribute)

	return (int32)(ret)
}

func Fn_g_file_info_get_attribute_int64(info unsafe.Pointer, attribute string) int64 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_int64(c_info, c_attribute)

	return (int64)(ret)
}

func Fn_g_file_info_get_attribute_object(info unsafe.Pointer, attribute string) unsafe.Pointer {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_object(c_info, c_attribute)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_attribute_status(info unsafe.Pointer, attribute string) int {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_status(c_info, c_attribute)

	return (int)(ret)
}

func Fn_g_file_info_get_attribute_string(info unsafe.Pointer, attribute string) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_string(c_info, c_attribute)

	return C.GoString(ret)
}

// UNSUPPORTED : g_file_info_get_attribute_stringv : no array length

func Fn_g_file_info_get_attribute_type(info unsafe.Pointer, attribute string) int {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_type(c_info, c_attribute)

	return (int)(ret)
}

func Fn_g_file_info_get_attribute_uint32(info unsafe.Pointer, attribute string) uint32 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_uint32(c_info, c_attribute)

	return (uint32)(ret)
}

func Fn_g_file_info_get_attribute_uint64(info unsafe.Pointer, attribute string) uint64 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_get_attribute_uint64(c_info, c_attribute)

	return (uint64)(ret)
}

func Fn_g_file_info_get_content_type(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_content_type(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_get_display_name(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_display_name(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_get_edit_name(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_edit_name(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_get_etag(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_etag(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_get_file_type(info unsafe.Pointer) int {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_file_type(c_info)

	return (int)(ret)
}

func Fn_g_file_info_get_icon(info unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_icon(c_info)

	return unsafe.Pointer(ret)
}

func Fn_g_file_info_get_is_backup(info unsafe.Pointer) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_is_backup(c_info)

	return toGoBool(ret)
}

func Fn_g_file_info_get_is_hidden(info unsafe.Pointer) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_is_hidden(c_info)

	return toGoBool(ret)
}

func Fn_g_file_info_get_is_symlink(info unsafe.Pointer) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_is_symlink(c_info)

	return toGoBool(ret)
}

func Fn_g_file_info_get_modification_time(info unsafe.Pointer, result unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_result := (*C.GTimeVal)(unsafe.Pointer(result))

	C.g_file_info_get_modification_time(c_info, c_result)
}

func Fn_g_file_info_get_name(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_name(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_get_size(info unsafe.Pointer) int64 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_size(c_info)

	return (int64)(ret)
}

func Fn_g_file_info_get_sort_order(info unsafe.Pointer) int32 {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_sort_order(c_info)

	return (int32)(ret)
}

func Fn_g_file_info_get_symlink_target(info unsafe.Pointer) string {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	ret := C.g_file_info_get_symlink_target(c_info)

	return C.GoString(ret)
}

func Fn_g_file_info_has_attribute(info unsafe.Pointer, attribute string) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	ret := C.g_file_info_has_attribute(c_info, c_attribute)

	return toGoBool(ret)
}

func Fn_g_file_info_has_namespace(info unsafe.Pointer, nameSpace string) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_nameSpace := (*C.char)(C.CString(nameSpace))
	defer C.free(unsafe.Pointer(c_nameSpace))

	ret := C.g_file_info_has_namespace(c_info, c_nameSpace)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_info_list_attributes : no array length

func Fn_g_file_info_remove_attribute(info unsafe.Pointer, attribute string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	C.g_file_info_remove_attribute(c_info, c_attribute)
}

func Fn_g_file_info_set_attribute(info unsafe.Pointer, attribute string, type_ int, valueP unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_type_ := (C.GFileAttributeType)(type_)

	c_valueP := (C.gpointer)(valueP)

	C.g_file_info_set_attribute(c_info, c_attribute, c_type_, c_valueP)
}

func Fn_g_file_info_set_attribute_boolean(info unsafe.Pointer, attribute string, attrValue bool) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := toCBool(attrValue)

	C.g_file_info_set_attribute_boolean(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_byte_string(info unsafe.Pointer, attribute string, attrValue string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(c_attrValue))

	C.g_file_info_set_attribute_byte_string(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_int32(info unsafe.Pointer, attribute string, attrValue int32) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (C.gint32)(attrValue)

	C.g_file_info_set_attribute_int32(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_int64(info unsafe.Pointer, attribute string, attrValue int64) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (C.gint64)(attrValue)

	C.g_file_info_set_attribute_int64(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_mask(info unsafe.Pointer, mask unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_mask := (*C.GFileAttributeMatcher)(unsafe.Pointer(mask))

	C.g_file_info_set_attribute_mask(c_info, c_mask)
}

func Fn_g_file_info_set_attribute_object(info unsafe.Pointer, attribute string, attrValue unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (*C.GObject)(unsafe.Pointer(attrValue))

	C.g_file_info_set_attribute_object(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_status(info unsafe.Pointer, attribute string, status int) bool {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_status := (C.GFileAttributeStatus)(status)

	ret := C.g_file_info_set_attribute_status(c_info, c_attribute, c_status)

	return toGoBool(ret)
}

func Fn_g_file_info_set_attribute_string(info unsafe.Pointer, attribute string, attrValue string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(c_attrValue))

	C.g_file_info_set_attribute_string(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_stringv(info unsafe.Pointer, attribute string, attrValue []string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	attrValueLen := len(attrValue)
	c_attrValueArray := C.malloc((C.ulong)(attrValueLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_attrValueArray))
	attrValueSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_attrValueArray))[:attrValueLen:attrValueLen]
	for attrValuei, attrValueString := range attrValue {
		attrValueSlice[attrValuei] = (*C.gchar)(C.CString(attrValueString))
		defer C.free(unsafe.Pointer(attrValueSlice[attrValuei]))
	}
	c_attrValue := &attrValueSlice[0]

	C.g_file_info_set_attribute_stringv(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_uint32(info unsafe.Pointer, attribute string, attrValue uint32) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (C.guint32)(attrValue)

	C.g_file_info_set_attribute_uint32(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_attribute_uint64(info unsafe.Pointer, attribute string, attrValue uint64) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_attrValue := (C.guint64)(attrValue)

	C.g_file_info_set_attribute_uint64(c_info, c_attribute, c_attrValue)
}

func Fn_g_file_info_set_content_type(info unsafe.Pointer, contentType string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_contentType := (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(c_contentType))

	C.g_file_info_set_content_type(c_info, c_contentType)
}

func Fn_g_file_info_set_display_name(info unsafe.Pointer, displayName string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_displayName := (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(c_displayName))

	C.g_file_info_set_display_name(c_info, c_displayName)
}

func Fn_g_file_info_set_edit_name(info unsafe.Pointer, editName string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_editName := (*C.char)(C.CString(editName))
	defer C.free(unsafe.Pointer(c_editName))

	C.g_file_info_set_edit_name(c_info, c_editName)
}

func Fn_g_file_info_set_file_type(info unsafe.Pointer, type_ int) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_type_ := (C.GFileType)(type_)

	C.g_file_info_set_file_type(c_info, c_type_)
}

func Fn_g_file_info_set_icon(info unsafe.Pointer, icon unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	C.g_file_info_set_icon(c_info, c_icon)
}

func Fn_g_file_info_set_is_hidden(info unsafe.Pointer, isHidden bool) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_isHidden := toCBool(isHidden)

	C.g_file_info_set_is_hidden(c_info, c_isHidden)
}

func Fn_g_file_info_set_is_symlink(info unsafe.Pointer, isSymlink bool) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_isSymlink := toCBool(isSymlink)

	C.g_file_info_set_is_symlink(c_info, c_isSymlink)
}

func Fn_g_file_info_set_modification_time(info unsafe.Pointer, mtime unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_mtime := (*C.GTimeVal)(unsafe.Pointer(mtime))

	C.g_file_info_set_modification_time(c_info, c_mtime)
}

func Fn_g_file_info_set_name(info unsafe.Pointer, name string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	C.g_file_info_set_name(c_info, c_name)
}

func Fn_g_file_info_set_size(info unsafe.Pointer, size int64) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_size := (C.goffset)(size)

	C.g_file_info_set_size(c_info, c_size)
}

func Fn_g_file_info_set_sort_order(info unsafe.Pointer, sortOrder int32) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_sortOrder := (C.gint32)(sortOrder)

	C.g_file_info_set_sort_order(c_info, c_sortOrder)
}

func Fn_g_file_info_set_symlink_target(info unsafe.Pointer, symlinkTarget string) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_symlinkTarget := (*C.char)(C.CString(symlinkTarget))
	defer C.free(unsafe.Pointer(c_symlinkTarget))

	C.g_file_info_set_symlink_target(c_info, c_symlinkTarget)
}

func Fn_g_file_info_unset_attribute_mask(info unsafe.Pointer) {
	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	C.g_file_info_unset_attribute_mask(c_info)
}

func Fn_g_file_input_stream_query_info(stream unsafe.Pointer, attributes string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileInputStream)(unsafe.Pointer(stream))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_input_stream_query_info(c_stream, c_attributes, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_input_stream_query_info_async : parameter 'callback' is callback

func Fn_g_file_input_stream_query_info_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_input_stream_query_info_finish(c_stream, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_cancel(monitor unsafe.Pointer) bool {
	c_monitor := (*C.GFileMonitor)(unsafe.Pointer(monitor))

	ret := C.g_file_monitor_cancel(c_monitor)

	return toGoBool(ret)
}

func Fn_g_file_monitor_emit_event(monitor unsafe.Pointer, child unsafe.Pointer, otherFile unsafe.Pointer, eventType int) {
	c_monitor := (*C.GFileMonitor)(unsafe.Pointer(monitor))

	c_child := (*C.GFile)(unsafe.Pointer(child))

	c_otherFile := (*C.GFile)(unsafe.Pointer(otherFile))

	c_eventType := (C.GFileMonitorEvent)(eventType)

	C.g_file_monitor_emit_event(c_monitor, c_child, c_otherFile, c_eventType)
}

func Fn_g_file_monitor_is_cancelled(monitor unsafe.Pointer) bool {
	c_monitor := (*C.GFileMonitor)(unsafe.Pointer(monitor))

	ret := C.g_file_monitor_is_cancelled(c_monitor)

	return toGoBool(ret)
}

func Fn_g_file_monitor_set_rate_limit(monitor unsafe.Pointer, limitMsecs int) {
	c_monitor := (*C.GFileMonitor)(unsafe.Pointer(monitor))

	c_limitMsecs := (C.gint)(limitMsecs)

	C.g_file_monitor_set_rate_limit(c_monitor, c_limitMsecs)
}

func Fn_g_file_output_stream_get_etag(stream unsafe.Pointer) string {
	c_stream := (*C.GFileOutputStream)(unsafe.Pointer(stream))

	ret := C.g_file_output_stream_get_etag(c_stream)

	return C.GoString(ret)
}

func Fn_g_file_output_stream_query_info(stream unsafe.Pointer, attributes string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileOutputStream)(unsafe.Pointer(stream))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_output_stream_query_info(c_stream, c_attributes, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_output_stream_query_info_async : parameter 'callback' is callback

func Fn_g_file_output_stream_query_info_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFileOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_output_stream_query_info_finish(c_stream, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_filename_completer_new() unsafe.Pointer {
	ret := C.g_filename_completer_new()

	return unsafe.Pointer(ret)
}

func Fn_g_filename_completer_get_completion_suffix(completer unsafe.Pointer, initialText string) string {
	c_completer := (*C.GFilenameCompleter)(unsafe.Pointer(completer))

	c_initialText := (*C.char)(C.CString(initialText))
	defer C.free(unsafe.Pointer(c_initialText))

	ret := C.g_filename_completer_get_completion_suffix(c_completer, c_initialText)

	return C.GoString(ret)
}

// UNSUPPORTED : g_filename_completer_get_completions : no array length

func Fn_g_filename_completer_set_dirs_only(completer unsafe.Pointer, dirsOnly bool) {
	c_completer := (*C.GFilenameCompleter)(unsafe.Pointer(completer))

	c_dirsOnly := toCBool(dirsOnly)

	C.g_filename_completer_set_dirs_only(c_completer, c_dirsOnly)
}

func Fn_g_filter_input_stream_get_base_stream(stream unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFilterInputStream)(unsafe.Pointer(stream))

	ret := C.g_filter_input_stream_get_base_stream(c_stream)

	return unsafe.Pointer(ret)
}

func Fn_g_filter_input_stream_get_close_base_stream(stream unsafe.Pointer) bool {
	c_stream := (*C.GFilterInputStream)(unsafe.Pointer(stream))

	ret := C.g_filter_input_stream_get_close_base_stream(c_stream)

	return toGoBool(ret)
}

func Fn_g_filter_input_stream_set_close_base_stream(stream unsafe.Pointer, closeBase bool) {
	c_stream := (*C.GFilterInputStream)(unsafe.Pointer(stream))

	c_closeBase := toCBool(closeBase)

	C.g_filter_input_stream_set_close_base_stream(c_stream, c_closeBase)
}

func Fn_g_filter_output_stream_get_base_stream(stream unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GFilterOutputStream)(unsafe.Pointer(stream))

	ret := C.g_filter_output_stream_get_base_stream(c_stream)

	return unsafe.Pointer(ret)
}

func Fn_g_filter_output_stream_get_close_base_stream(stream unsafe.Pointer) bool {
	c_stream := (*C.GFilterOutputStream)(unsafe.Pointer(stream))

	ret := C.g_filter_output_stream_get_close_base_stream(c_stream)

	return toGoBool(ret)
}

func Fn_g_filter_output_stream_set_close_base_stream(stream unsafe.Pointer, closeBase bool) {
	c_stream := (*C.GFilterOutputStream)(unsafe.Pointer(stream))

	c_closeBase := toCBool(closeBase)

	C.g_filter_output_stream_set_close_base_stream(c_stream, c_closeBase)
}

// UNSUPPORTED : g_io_module_new : blacklisted

// UNSUPPORTED : g_io_module_load : blacklisted

// UNSUPPORTED : g_io_module_unload : blacklisted

// UNSUPPORTED : g_io_module_query : blacklisted

func Fn_g_io_stream_clear_pending(stream unsafe.Pointer) {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	C.g_io_stream_clear_pending(c_stream)
}

func Fn_g_io_stream_close(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close(c_stream, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_close_async : parameter 'callback' is callback

func Fn_g_io_stream_close_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_close_finish(c_stream, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_io_stream_get_input_stream(stream unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	ret := C.g_io_stream_get_input_stream(c_stream)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_get_output_stream(stream unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	ret := C.g_io_stream_get_output_stream(c_stream)

	return unsafe.Pointer(ret)
}

func Fn_g_io_stream_has_pending(stream unsafe.Pointer) bool {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	ret := C.g_io_stream_has_pending(c_stream)

	return toGoBool(ret)
}

func Fn_g_io_stream_is_closed(stream unsafe.Pointer) bool {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	ret := C.g_io_stream_is_closed(c_stream)

	return toGoBool(ret)
}

func Fn_g_io_stream_set_pending(stream unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GIOStream)(unsafe.Pointer(stream))

	cError := (**C.GError)(error)

	ret := C.g_io_stream_set_pending(c_stream, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_io_stream_splice_async : parameter 'callback' is callback

func Fn_g_inet_address_new_any(family int) unsafe.Pointer {
	c_family := (C.GSocketFamily)(family)

	ret := C.g_inet_address_new_any(c_family)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_new_from_bytes(bytes []uint8, family int) unsafe.Pointer {
	c_bytes := (*C.guint8)(unsafe.Pointer(&bytes[0]))

	c_family := (C.GSocketFamily)(family)

	ret := C.g_inet_address_new_from_bytes(c_bytes, c_family)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_new_from_string(string_ string) unsafe.Pointer {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.g_inet_address_new_from_string(c_string_)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_new_loopback(family int) unsafe.Pointer {
	c_family := (C.GSocketFamily)(family)

	ret := C.g_inet_address_new_loopback(c_family)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_address_get_family(address unsafe.Pointer) int {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_family(c_address)

	return (int)(ret)
}

func Fn_g_inet_address_get_is_any(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_any(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_link_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_link_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_loopback(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_loopback(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_global(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_mc_global(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_link_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_mc_link_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_node_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_mc_node_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_org_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_mc_org_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_mc_site_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_mc_site_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_multicast(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_multicast(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_is_site_local(address unsafe.Pointer) bool {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_is_site_local(c_address)

	return toGoBool(ret)
}

func Fn_g_inet_address_get_native_size(address unsafe.Pointer) uint64 {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_get_native_size(c_address)

	return (uint64)(ret)
}

func Fn_g_inet_address_to_bytes(address unsafe.Pointer) *uint8 {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_to_bytes(c_address)

	return (*uint8)(ret)
}

func Fn_g_inet_address_to_string(address unsafe.Pointer) string {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	ret := C.g_inet_address_to_string(c_address)

	return C.GoString(ret)
}

func Fn_g_inet_socket_address_new(address unsafe.Pointer, port uint16) unsafe.Pointer {
	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	c_port := (C.guint16)(port)

	ret := C.g_inet_socket_address_new(c_address, c_port)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_address(address unsafe.Pointer) unsafe.Pointer {
	c_address := (*C.GInetSocketAddress)(unsafe.Pointer(address))

	ret := C.g_inet_socket_address_get_address(c_address)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_port(address unsafe.Pointer) uint16 {
	c_address := (*C.GInetSocketAddress)(unsafe.Pointer(address))

	ret := C.g_inet_socket_address_get_port(c_address)

	return (uint16)(ret)
}

func Fn_g_input_stream_clear_pending(stream unsafe.Pointer) {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	C.g_input_stream_clear_pending(c_stream)
}

func Fn_g_input_stream_close(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_close(c_stream, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_input_stream_close_async : parameter 'callback' is callback

func Fn_g_input_stream_close_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_close_finish(c_stream, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_input_stream_has_pending(stream unsafe.Pointer) bool {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	ret := C.g_input_stream_has_pending(c_stream)

	return toGoBool(ret)
}

func Fn_g_input_stream_is_closed(stream unsafe.Pointer) bool {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	ret := C.g_input_stream_is_closed(c_stream)

	return toGoBool(ret)
}

// UNSUPPORTED : g_input_stream_read : blacklisted

// UNSUPPORTED : g_input_stream_read_all : blacklisted

// UNSUPPORTED : g_input_stream_read_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_bytes_async : parameter 'callback' is callback

func Fn_g_input_stream_read_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_read_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

func Fn_g_input_stream_set_pending(stream unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_set_pending(c_stream, cError)

	return toGoBool(ret)
}

func Fn_g_input_stream_skip(stream unsafe.Pointer, count uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_skip(c_stream, c_count, c_cancellable, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_input_stream_skip_async : parameter 'callback' is callback

func Fn_g_input_stream_skip_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_skip_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_list_store_insert_sorted : parameter 'compare_func' is callback

// UNSUPPORTED : g_list_store_sort : parameter 'compare_func' is callback

func Fn_g_memory_input_stream_new() unsafe.Pointer {
	ret := C.g_memory_input_stream_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_memory_input_stream_new_from_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_input_stream_add_data : parameter 'destroy' is callback

// UNSUPPORTED : g_memory_output_stream_new : parameter 'realloc_function' is callback

func Fn_g_memory_output_stream_get_data(ostream unsafe.Pointer) unsafe.Pointer {
	c_ostream := (*C.GMemoryOutputStream)(unsafe.Pointer(ostream))

	ret := C.g_memory_output_stream_get_data(c_ostream)

	return (unsafe.Pointer)(ret)
}

func Fn_g_memory_output_stream_get_data_size(ostream unsafe.Pointer) uint64 {
	c_ostream := (*C.GMemoryOutputStream)(unsafe.Pointer(ostream))

	ret := C.g_memory_output_stream_get_data_size(c_ostream)

	return (uint64)(ret)
}

func Fn_g_memory_output_stream_get_size(ostream unsafe.Pointer) uint64 {
	c_ostream := (*C.GMemoryOutputStream)(unsafe.Pointer(ostream))

	ret := C.g_memory_output_stream_get_size(c_ostream)

	return (uint64)(ret)
}

func Fn_g_memory_output_stream_steal_data(ostream unsafe.Pointer) unsafe.Pointer {
	c_ostream := (*C.GMemoryOutputStream)(unsafe.Pointer(ostream))

	ret := C.g_memory_output_stream_steal_data(c_ostream)

	return (unsafe.Pointer)(ret)
}

// UNSUPPORTED : g_menu_attribute_iter_get_next : parameter 'out_name' is non array with indirect count > 1

// UNSUPPORTED : g_menu_link_iter_get_next : parameter 'out_link' is non array with indirect count > 1

func Fn_g_mount_operation_new() unsafe.Pointer {
	ret := C.g_mount_operation_new()

	return unsafe.Pointer(ret)
}

func Fn_g_mount_operation_get_anonymous(op unsafe.Pointer) bool {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_anonymous(c_op)

	return toGoBool(ret)
}

func Fn_g_mount_operation_get_choice(op unsafe.Pointer) int {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_choice(c_op)

	return (int)(ret)
}

func Fn_g_mount_operation_get_domain(op unsafe.Pointer) string {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_domain(c_op)

	return C.GoString(ret)
}

func Fn_g_mount_operation_get_password(op unsafe.Pointer) string {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_password(c_op)

	return C.GoString(ret)
}

func Fn_g_mount_operation_get_password_save(op unsafe.Pointer) int {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_password_save(c_op)

	return (int)(ret)
}

func Fn_g_mount_operation_get_username(op unsafe.Pointer) string {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	ret := C.g_mount_operation_get_username(c_op)

	return C.GoString(ret)
}

func Fn_g_mount_operation_reply(op unsafe.Pointer, result int) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_result := (C.GMountOperationResult)(result)

	C.g_mount_operation_reply(c_op, c_result)
}

func Fn_g_mount_operation_set_anonymous(op unsafe.Pointer, anonymous bool) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_anonymous := toCBool(anonymous)

	C.g_mount_operation_set_anonymous(c_op, c_anonymous)
}

func Fn_g_mount_operation_set_choice(op unsafe.Pointer, choice int) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_choice := (C.int)(choice)

	C.g_mount_operation_set_choice(c_op, c_choice)
}

func Fn_g_mount_operation_set_domain(op unsafe.Pointer, domain string) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_domain := (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(c_domain))

	C.g_mount_operation_set_domain(c_op, c_domain)
}

func Fn_g_mount_operation_set_password(op unsafe.Pointer, password string) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_password := (*C.char)(C.CString(password))
	defer C.free(unsafe.Pointer(c_password))

	C.g_mount_operation_set_password(c_op, c_password)
}

func Fn_g_mount_operation_set_password_save(op unsafe.Pointer, save int) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_save := (C.GPasswordSave)(save)

	C.g_mount_operation_set_password_save(c_op, c_save)
}

func Fn_g_mount_operation_set_username(op unsafe.Pointer, username string) {
	c_op := (*C.GMountOperation)(unsafe.Pointer(op))

	c_username := (*C.char)(C.CString(username))
	defer C.free(unsafe.Pointer(c_username))

	C.g_mount_operation_set_username(c_op, c_username)
}

func Fn_g_network_address_new(hostname string, port uint16) unsafe.Pointer {
	c_hostname := (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	ret := C.g_network_address_new(c_hostname, c_port)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_get_hostname(addr unsafe.Pointer) string {
	c_addr := (*C.GNetworkAddress)(unsafe.Pointer(addr))

	ret := C.g_network_address_get_hostname(c_addr)

	return C.GoString(ret)
}

func Fn_g_network_address_get_port(addr unsafe.Pointer) uint16 {
	c_addr := (*C.GNetworkAddress)(unsafe.Pointer(addr))

	ret := C.g_network_address_get_port(c_addr)

	return (uint16)(ret)
}

func Fn_g_network_address_get_scheme(addr unsafe.Pointer) string {
	c_addr := (*C.GNetworkAddress)(unsafe.Pointer(addr))

	ret := C.g_network_address_get_scheme(c_addr)

	return C.GoString(ret)
}

func Fn_g_network_address_parse(hostAndPort string, defaultPort uint16, error unsafe.Pointer) unsafe.Pointer {
	c_hostAndPort := (*C.gchar)(C.CString(hostAndPort))
	defer C.free(unsafe.Pointer(c_hostAndPort))

	c_defaultPort := (C.guint16)(defaultPort)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse(c_hostAndPort, c_defaultPort, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_address_parse_uri(uri string, defaultPort uint16, error unsafe.Pointer) unsafe.Pointer {
	c_uri := (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(c_uri))

	c_defaultPort := (C.guint16)(defaultPort)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse_uri(c_uri, c_defaultPort, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_new(service string, protocol string, domain string) unsafe.Pointer {
	c_service := (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(c_domain))

	ret := C.g_network_service_new(c_service, c_protocol, c_domain)

	return unsafe.Pointer(ret)
}

func Fn_g_network_service_get_domain(srv unsafe.Pointer) string {
	c_srv := (*C.GNetworkService)(unsafe.Pointer(srv))

	ret := C.g_network_service_get_domain(c_srv)

	return C.GoString(ret)
}

func Fn_g_network_service_get_protocol(srv unsafe.Pointer) string {
	c_srv := (*C.GNetworkService)(unsafe.Pointer(srv))

	ret := C.g_network_service_get_protocol(c_srv)

	return C.GoString(ret)
}

func Fn_g_network_service_get_scheme(srv unsafe.Pointer) string {
	c_srv := (*C.GNetworkService)(unsafe.Pointer(srv))

	ret := C.g_network_service_get_scheme(c_srv)

	return C.GoString(ret)
}

func Fn_g_network_service_get_service(srv unsafe.Pointer) string {
	c_srv := (*C.GNetworkService)(unsafe.Pointer(srv))

	ret := C.g_network_service_get_service(c_srv)

	return C.GoString(ret)
}

func Fn_g_network_service_set_scheme(srv unsafe.Pointer, scheme string) {
	c_srv := (*C.GNetworkService)(unsafe.Pointer(srv))

	c_scheme := (*C.gchar)(C.CString(scheme))
	defer C.free(unsafe.Pointer(c_scheme))

	C.g_network_service_set_scheme(c_srv, c_scheme)
}

func Fn_g_notification_set_priority(notification unsafe.Pointer, priority int) {
	c_notification := (*C.GNotification)(unsafe.Pointer(notification))

	c_priority := (C.GNotificationPriority)(priority)

	C.g_notification_set_priority(c_notification, c_priority)
}

func Fn_g_output_stream_clear_pending(stream unsafe.Pointer) {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	C.g_output_stream_clear_pending(c_stream)
}

func Fn_g_output_stream_close(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_close(c_stream, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_close_async : parameter 'callback' is callback

func Fn_g_output_stream_close_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_close_finish(c_stream, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_flush(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_flush(c_stream, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_flush_async : parameter 'callback' is callback

func Fn_g_output_stream_flush_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_flush_finish(c_stream, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_has_pending(stream unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	ret := C.g_output_stream_has_pending(c_stream)

	return toGoBool(ret)
}

func Fn_g_output_stream_is_closed(stream unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	ret := C.g_output_stream_is_closed(c_stream)

	return toGoBool(ret)
}

func Fn_g_output_stream_is_closing(stream unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	ret := C.g_output_stream_is_closing(c_stream)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_printf : parameter 'error' is non array with indirect count > 1

func Fn_g_output_stream_set_pending(stream unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_set_pending(c_stream, cError)

	return toGoBool(ret)
}

func Fn_g_output_stream_splice(stream unsafe.Pointer, source unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_source := (*C.GInputStream)(unsafe.Pointer(source))

	c_flags := (C.GOutputStreamSpliceFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_splice(c_stream, c_source, c_flags, c_cancellable, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_splice_async : parameter 'callback' is callback

func Fn_g_output_stream_splice_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_splice_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_vprintf : parameter 'error' is non array with indirect count > 1

func Fn_g_output_stream_write(stream unsafe.Pointer, buffer []uint8, count uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_buffer := (unsafe.Pointer)(unsafe.Pointer(&buffer[0]))

	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write(c_stream, c_buffer, c_count, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_output_stream_write_all(stream unsafe.Pointer, buffer []uint8, count uint64, bytesWritten *uint64, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_buffer := (unsafe.Pointer)(unsafe.Pointer(&buffer[0]))

	c_count := (C.gsize)(count)

	c_bytesWritten := (*C.gsize)(unsafe.Pointer(bytesWritten))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_all(c_stream, c_buffer, c_count, c_bytesWritten, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_write_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_async : parameter 'callback' is callback

func Fn_g_output_stream_write_bytes(stream unsafe.Pointer, bytes unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_bytes := (*C.GBytes)(unsafe.Pointer(bytes))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_bytes(c_stream, c_bytes, c_cancellable, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_write_bytes_async : parameter 'callback' is callback

func Fn_g_output_stream_write_bytes_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_bytes_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

func Fn_g_output_stream_write_finish(stream unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_finish(c_stream, c_result, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_writev_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_async : parameter 'callback' is callback

func Fn_g_permission_acquire(permission unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire(c_permission, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_acquire_async : parameter 'callback' is callback

func Fn_g_permission_acquire_finish(permission unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_permission_acquire_finish(c_permission, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_permission_get_allowed(permission unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	ret := C.g_permission_get_allowed(c_permission)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_acquire(permission unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	ret := C.g_permission_get_can_acquire(c_permission)

	return toGoBool(ret)
}

func Fn_g_permission_get_can_release(permission unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	ret := C.g_permission_get_can_release(c_permission)

	return toGoBool(ret)
}

func Fn_g_permission_impl_update(permission unsafe.Pointer, allowed bool, canAcquire bool, canRelease bool) {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	c_allowed := toCBool(allowed)

	c_canAcquire := toCBool(canAcquire)

	c_canRelease := toCBool(canRelease)

	C.g_permission_impl_update(c_permission, c_allowed, c_canAcquire, c_canRelease)
}

func Fn_g_permission_release(permission unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_permission_release(c_permission, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_permission_release_async : parameter 'callback' is callback

func Fn_g_permission_release_finish(permission unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_permission := (*C.GPermission)(unsafe.Pointer(permission))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_permission_release_finish(c_permission, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_proxy_address_new(inetaddr unsafe.Pointer, port uint16, protocol string, destHostname string, destPort uint16, username *string, password *string) unsafe.Pointer {
	c_inetaddr := (*C.GInetAddress)(unsafe.Pointer(inetaddr))

	c_port := (C.guint16)(port)

	c_protocol := (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(c_protocol))

	c_destHostname := (*C.gchar)(C.CString(destHostname))
	defer C.free(unsafe.Pointer(c_destHostname))

	c_destPort := (C.guint16)(destPort)

	var c_usernameValue (*C.gchar)
	if username != nil {
		c_usernameValue = (*C.gchar)(C.CString(*username))
		defer C.free(unsafe.Pointer(c_usernameValue))
	}
	c_username := c_usernameValue

	var c_passwordValue (*C.gchar)
	if password != nil {
		c_passwordValue = (*C.gchar)(C.CString(*password))
		defer C.free(unsafe.Pointer(c_passwordValue))
	}
	c_password := c_passwordValue

	ret := C.g_proxy_address_new(c_inetaddr, c_port, c_protocol, c_destHostname, c_destPort, c_username, c_password)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_address_get_destination_hostname(proxy unsafe.Pointer) string {
	c_proxy := (*C.GProxyAddress)(unsafe.Pointer(proxy))

	ret := C.g_proxy_address_get_destination_hostname(c_proxy)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_destination_port(proxy unsafe.Pointer) uint16 {
	c_proxy := (*C.GProxyAddress)(unsafe.Pointer(proxy))

	ret := C.g_proxy_address_get_destination_port(c_proxy)

	return (uint16)(ret)
}

func Fn_g_proxy_address_get_password(proxy unsafe.Pointer) string {
	c_proxy := (*C.GProxyAddress)(unsafe.Pointer(proxy))

	ret := C.g_proxy_address_get_password(c_proxy)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_protocol(proxy unsafe.Pointer) string {
	c_proxy := (*C.GProxyAddress)(unsafe.Pointer(proxy))

	ret := C.g_proxy_address_get_protocol(c_proxy)

	return C.GoString(ret)
}

func Fn_g_proxy_address_get_username(proxy unsafe.Pointer) string {
	c_proxy := (*C.GProxyAddress)(unsafe.Pointer(proxy))

	ret := C.g_proxy_address_get_username(c_proxy)

	return C.GoString(ret)
}

func Fn_g_resolver_lookup_by_address(resolver unsafe.Pointer, address unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) string {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_address := (*C.GInetAddress)(unsafe.Pointer(address))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address(c_resolver, c_address, c_cancellable, cError)

	return C.GoString(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_address_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_by_address_finish(resolver unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) string {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_address_finish(c_resolver, c_result, cError)

	return C.GoString(ret)
}

func Fn_g_resolver_lookup_by_name(resolver unsafe.Pointer, hostname string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_hostname := (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(c_hostname))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name(c_resolver, c_hostname, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_by_name_finish(resolver unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_by_name_finish(c_resolver, c_result, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_records_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_service(resolver unsafe.Pointer, service string, protocol string, domain string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_service := (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(c_domain))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service(c_resolver, c_service, c_protocol, c_domain, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_resolver_lookup_service_async : parameter 'callback' is callback

func Fn_g_resolver_lookup_service_finish(resolver unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_resolver_lookup_service_finish(c_resolver, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_resolver_set_default(resolver unsafe.Pointer) {
	c_resolver := (*C.GResolver)(unsafe.Pointer(resolver))

	C.g_resolver_set_default(c_resolver)
}

func Fn_g_resolver_free_addresses(addresses unsafe.Pointer) {
	c_addresses := (*C.GList)(unsafe.Pointer(addresses))

	C.g_resolver_free_addresses(c_addresses)
}

func Fn_g_resolver_free_targets(targets unsafe.Pointer) {
	c_targets := (*C.GList)(unsafe.Pointer(targets))

	C.g_resolver_free_targets(c_targets)
}

func Fn_g_resolver_get_default() unsafe.Pointer {
	ret := C.g_resolver_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new(schemaId string) unsafe.Pointer {
	c_schemaId := (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(c_schemaId))

	ret := C.g_settings_new(c_schemaId)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend(schemaId string, backend unsafe.Pointer) unsafe.Pointer {
	c_schemaId := (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(c_schemaId))

	c_backend := (*C.GSettingsBackend)(unsafe.Pointer(backend))

	ret := C.g_settings_new_with_backend(c_schemaId, c_backend)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_backend_and_path(schemaId string, backend unsafe.Pointer, path string) unsafe.Pointer {
	c_schemaId := (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(c_schemaId))

	c_backend := (*C.GSettingsBackend)(unsafe.Pointer(backend))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	ret := C.g_settings_new_with_backend_and_path(c_schemaId, c_backend, c_path)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_new_with_path(schemaId string, path string) unsafe.Pointer {
	c_schemaId := (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(c_schemaId))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	ret := C.g_settings_new_with_path(c_schemaId, c_path)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_apply(settings unsafe.Pointer) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	C.g_settings_apply(c_settings)
}

func Fn_g_settings_bind(settings unsafe.Pointer, key string, object unsafe.Pointer, property string, flags int) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(c_property))

	c_flags := (C.GSettingsBindFlags)(flags)

	C.g_settings_bind(c_settings, c_key, c_object, c_property, c_flags)
}

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

func Fn_g_settings_bind_writable(settings unsafe.Pointer, key string, object unsafe.Pointer, property string, inverted bool) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_object := (C.gpointer)(object)

	c_property := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(c_property))

	c_inverted := toCBool(inverted)

	C.g_settings_bind_writable(c_settings, c_key, c_object, c_property, c_inverted)
}

func Fn_g_settings_delay(settings unsafe.Pointer) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	C.g_settings_delay(c_settings)
}

func Fn_g_settings_get(settings unsafe.Pointer, key string, format string) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_format := (*C.gchar)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	C.c_g_settings_get(c_settings, c_key, c_format)
}

func Fn_g_settings_get_boolean(settings unsafe.Pointer, key string) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_boolean(c_settings, c_key)

	return toGoBool(ret)
}

func Fn_g_settings_get_child(settings unsafe.Pointer, name string) unsafe.Pointer {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_settings_get_child(c_settings, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_get_double(settings unsafe.Pointer, key string) float64 {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_double(c_settings, c_key)

	return (float64)(ret)
}

func Fn_g_settings_get_enum(settings unsafe.Pointer, key string) int {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_enum(c_settings, c_key)

	return (int)(ret)
}

func Fn_g_settings_get_flags(settings unsafe.Pointer, key string) uint {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_flags(c_settings, c_key)

	return (uint)(ret)
}

func Fn_g_settings_get_has_unapplied(settings unsafe.Pointer) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	ret := C.g_settings_get_has_unapplied(c_settings)

	return toGoBool(ret)
}

func Fn_g_settings_get_int(settings unsafe.Pointer, key string) int {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_int(c_settings, c_key)

	return (int)(ret)
}

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

func Fn_g_settings_get_string(settings unsafe.Pointer, key string) string {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_string(c_settings, c_key)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_get_strv : no array length

func Fn_g_settings_get_value(settings unsafe.Pointer, key string) unsafe.Pointer {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.g_settings_get_value(c_settings, c_key)

	return unsafe.Pointer(ret)
}

func Fn_g_settings_is_writable(settings unsafe.Pointer, name string) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_settings_is_writable(c_settings, c_name)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_children : no array length

// UNSUPPORTED : g_settings_list_keys : no array length

func Fn_g_settings_reset(settings unsafe.Pointer, key string) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	C.g_settings_reset(c_settings, c_key)
}

func Fn_g_settings_revert(settings unsafe.Pointer) {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	C.g_settings_revert(c_settings)
}

func Fn_g_settings_set(settings unsafe.Pointer, key string, format string) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_format := (*C.gchar)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	ret := C.c_g_settings_set(c_settings, c_key, c_format)

	return toGoBool(ret)
}

func Fn_g_settings_set_boolean(settings unsafe.Pointer, key string, value bool) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := toCBool(value)

	ret := C.g_settings_set_boolean(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_double(settings unsafe.Pointer, key string, value float64) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	ret := C.g_settings_set_double(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_enum(settings unsafe.Pointer, key string, value int) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	ret := C.g_settings_set_enum(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_flags(settings unsafe.Pointer, key string, value uint) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint)(value)

	ret := C.g_settings_set_flags(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_int(settings unsafe.Pointer, key string, value int) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	ret := C.g_settings_set_int(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_string(settings unsafe.Pointer, key string, value string) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	ret := C.g_settings_set_string(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_strv(settings unsafe.Pointer, key string, value []string) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	valueLen := len(value)
	c_valueArray := C.malloc((C.ulong)(valueLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_valueArray))
	valueSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_valueArray))[:valueLen:valueLen]
	for valuei, valueString := range value {
		valueSlice[valuei] = (*C.gchar)(C.CString(valueString))
		defer C.free(unsafe.Pointer(valueSlice[valuei]))
	}
	c_value := &valueSlice[0]

	ret := C.g_settings_set_strv(c_settings, c_key, c_value)

	return toGoBool(ret)
}

func Fn_g_settings_set_value(settings unsafe.Pointer, key string, value unsafe.Pointer) bool {
	c_settings := (*C.GSettings)(unsafe.Pointer(settings))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.GVariant)(unsafe.Pointer(value))

	ret := C.g_settings_set_value(c_settings, c_key, c_value)

	return toGoBool(ret)
}

// UNSUPPORTED : g_settings_list_relocatable_schemas : no array length

// UNSUPPORTED : g_settings_list_schemas : no array length

func Fn_g_settings_sync() {
	C.g_settings_sync()
}

func Fn_g_settings_unbind(object unsafe.Pointer, property string) {
	c_object := (C.gpointer)(object)

	c_property := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(c_property))

	C.g_settings_unbind(c_object, c_property)
}

// UNSUPPORTED : g_settings_backend_changed : blacklisted

// UNSUPPORTED : g_settings_backend_changed_tree : blacklisted

// UNSUPPORTED : g_settings_backend_keys_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_changed : blacklisted

// UNSUPPORTED : g_settings_backend_path_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_writable_changed : blacklisted

// UNSUPPORTED : g_settings_backend_flatten_tree : blacklisted

// UNSUPPORTED : g_settings_backend_get_default : blacklisted

// UNSUPPORTED : g_simple_async_result_new : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_from_error : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_result_new_take_error : parameter 'callback' is callback

func Fn_g_simple_async_result_complete(simple unsafe.Pointer) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	C.g_simple_async_result_complete(c_simple)
}

func Fn_g_simple_async_result_complete_in_idle(simple unsafe.Pointer) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	C.g_simple_async_result_complete_in_idle(c_simple)
}

func Fn_g_simple_async_result_get_op_res_gboolean(simple unsafe.Pointer) bool {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	ret := C.g_simple_async_result_get_op_res_gboolean(c_simple)

	return toGoBool(ret)
}

func Fn_g_simple_async_result_get_op_res_gpointer(simple unsafe.Pointer) unsafe.Pointer {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	ret := C.g_simple_async_result_get_op_res_gpointer(c_simple)

	return (unsafe.Pointer)(ret)
}

func Fn_g_simple_async_result_get_op_res_gssize(simple unsafe.Pointer) uint64 {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	ret := C.g_simple_async_result_get_op_res_gssize(c_simple)

	return (uint64)(ret)
}

func Fn_g_simple_async_result_get_source_tag(simple unsafe.Pointer) unsafe.Pointer {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	ret := C.g_simple_async_result_get_source_tag(c_simple)

	return (unsafe.Pointer)(ret)
}

func Fn_g_simple_async_result_propagate_error(simple unsafe.Pointer, error unsafe.Pointer) bool {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	cError := (**C.GError)(error)

	ret := C.g_simple_async_result_propagate_error(c_simple, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_simple_async_result_run_in_thread : parameter 'func' is callback

func Fn_g_simple_async_result_set_error(simple unsafe.Pointer, domain uint32, code int, format string) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_format := (*C.char)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	C.c_g_simple_async_result_set_error(c_simple, c_domain, c_code, c_format)
}

func Fn_g_simple_async_result_set_error_va(simple unsafe.Pointer, domain uint32, code int, format string) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_format := (*C.char)(C.CString(format))
	defer C.free(unsafe.Pointer(c_format))

	C.c_g_simple_async_result_set_error_va(c_simple, c_domain, c_code, c_format)
}

func Fn_g_simple_async_result_set_from_error(simple unsafe.Pointer, error unsafe.Pointer) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_error := (*C.GError)(unsafe.Pointer(error))

	C.g_simple_async_result_set_from_error(c_simple, c_error)
}

func Fn_g_simple_async_result_set_handle_cancellation(simple unsafe.Pointer, handleCancellation bool) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_handleCancellation := toCBool(handleCancellation)

	C.g_simple_async_result_set_handle_cancellation(c_simple, c_handleCancellation)
}

func Fn_g_simple_async_result_set_op_res_gboolean(simple unsafe.Pointer, opRes bool) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_opRes := toCBool(opRes)

	C.g_simple_async_result_set_op_res_gboolean(c_simple, c_opRes)
}

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

func Fn_g_simple_async_result_set_op_res_gssize(simple unsafe.Pointer, opRes uint64) {
	c_simple := (*C.GSimpleAsyncResult)(unsafe.Pointer(simple))

	c_opRes := (C.gssize)(opRes)

	C.g_simple_async_result_set_op_res_gssize(c_simple, c_opRes)
}

func Fn_g_simple_async_result_is_valid(result unsafe.Pointer, source unsafe.Pointer, sourceTag unsafe.Pointer) bool {
	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	c_source := (*C.GObject)(unsafe.Pointer(source))

	c_sourceTag := (C.gpointer)(sourceTag)

	ret := C.g_simple_async_result_is_valid(c_result, c_source, c_sourceTag)

	return toGoBool(ret)
}

func Fn_g_simple_permission_new(allowed bool) unsafe.Pointer {
	c_allowed := toCBool(allowed)

	ret := C.g_simple_permission_new(c_allowed)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_simple_proxy_resolver_set_ignore_hosts : parameter 'ignore_hosts' is non array with indirect count > 1

// UNSUPPORTED : g_simple_proxy_resolver_new : parameter 'ignore_hosts' is non array with indirect count > 1

func Fn_g_socket_new(family int, type_ int, protocol int, error unsafe.Pointer) unsafe.Pointer {
	c_family := (C.GSocketFamily)(family)

	c_type_ := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	cError := (**C.GError)(error)

	ret := C.g_socket_new(c_family, c_type_, c_protocol, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_new_from_fd(fd int, error unsafe.Pointer) unsafe.Pointer {
	c_fd := (C.gint)(fd)

	cError := (**C.GError)(error)

	ret := C.g_socket_new_from_fd(c_fd, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_accept(socket unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_accept(c_socket, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_bind(socket unsafe.Pointer, address unsafe.Pointer, allowReuse bool, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	c_allowReuse := toCBool(allowReuse)

	cError := (**C.GError)(error)

	ret := C.g_socket_bind(c_socket, c_address, c_allowReuse, cError)

	return toGoBool(ret)
}

func Fn_g_socket_check_connect_result(socket unsafe.Pointer, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_check_connect_result(c_socket, cError)

	return toGoBool(ret)
}

func Fn_g_socket_close(socket unsafe.Pointer, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_close(c_socket, cError)

	return toGoBool(ret)
}

func Fn_g_socket_condition_check(socket unsafe.Pointer, condition int) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_condition := (C.GIOCondition)(condition)

	ret := C.g_socket_condition_check(c_socket, c_condition)

	return (int)(ret)
}

func Fn_g_socket_condition_wait(socket unsafe.Pointer, condition int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_condition_wait(c_socket, c_condition, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connect(socket unsafe.Pointer, address unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_connect(c_socket, c_address, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connection_factory_create_connection(socket unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_connection_factory_create_connection(c_socket)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_create_source(socket unsafe.Pointer, condition int, cancellable unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	ret := C.g_socket_create_source(c_socket, c_condition, c_cancellable)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_blocking(socket unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_blocking(c_socket)

	return toGoBool(ret)
}

func Fn_g_socket_get_credentials(socket unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_credentials(c_socket, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_family(socket unsafe.Pointer) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_family(c_socket)

	return (int)(ret)
}

func Fn_g_socket_get_fd(socket unsafe.Pointer) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_fd(c_socket)

	return (int)(ret)
}

func Fn_g_socket_get_keepalive(socket unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_keepalive(c_socket)

	return toGoBool(ret)
}

func Fn_g_socket_get_listen_backlog(socket unsafe.Pointer) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_listen_backlog(c_socket)

	return (int)(ret)
}

func Fn_g_socket_get_local_address(socket unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_local_address(c_socket, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_protocol(socket unsafe.Pointer) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_protocol(c_socket)

	return (int)(ret)
}

func Fn_g_socket_get_remote_address(socket unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_get_remote_address(c_socket, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_get_socket_type(socket unsafe.Pointer) int {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_socket_type(c_socket)

	return (int)(ret)
}

func Fn_g_socket_get_timeout(socket unsafe.Pointer) uint {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_get_timeout(c_socket)

	return (uint)(ret)
}

func Fn_g_socket_is_closed(socket unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_is_closed(c_socket)

	return toGoBool(ret)
}

func Fn_g_socket_is_connected(socket unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_is_connected(c_socket)

	return toGoBool(ret)
}

func Fn_g_socket_listen(socket unsafe.Pointer, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	cError := (**C.GError)(error)

	ret := C.g_socket_listen(c_socket, cError)

	return toGoBool(ret)
}

func Fn_g_socket_receive(socket unsafe.Pointer, buffer []uint8, size uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_buffer := (*C.gchar)(unsafe.Pointer(&buffer[0]))

	c_size := (C.gsize)(size)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive(c_socket, c_buffer, c_size, c_cancellable, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_socket_receive_from : parameter 'address' is non array with indirect count > 1

// UNSUPPORTED : g_socket_receive_message : parameter 'address' is non array with indirect count > 1

func Fn_g_socket_receive_with_blocking(socket unsafe.Pointer, buffer []uint8, size uint64, blocking bool, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_buffer := (*C.gchar)(unsafe.Pointer(&buffer[0]))

	c_size := (C.gsize)(size)

	c_blocking := toCBool(blocking)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_receive_with_blocking(c_socket, c_buffer, c_size, c_blocking, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send(socket unsafe.Pointer, buffer []uint8, size uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_buffer := (*C.gchar)(unsafe.Pointer(&buffer[0]))

	c_size := (C.gsize)(size)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_send(c_socket, c_buffer, c_size, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_message(socket unsafe.Pointer, address unsafe.Pointer, vectors []OutputVector, numVectors int, messages []unsafe.Pointer, numMessages int, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	c_vectors := (*C.GOutputVector)(unsafe.Pointer(&vectors[0]))

	c_numVectors := (C.gint)(numVectors)

	c_messages := (**C.GSocketControlMessage)(unsafe.Pointer(&messages[0]))

	c_numMessages := (C.gint)(numMessages)

	c_flags := (C.gint)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_message(c_socket, c_address, c_vectors, c_numVectors, c_messages, c_numMessages, c_flags, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_to(socket unsafe.Pointer, address unsafe.Pointer, buffer []uint8, size uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	c_buffer := (*C.gchar)(unsafe.Pointer(&buffer[0]))

	c_size := (C.gsize)(size)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_to(c_socket, c_address, c_buffer, c_size, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_socket_send_with_blocking(socket unsafe.Pointer, buffer []uint8, size uint64, blocking bool, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_buffer := (*C.gchar)(unsafe.Pointer(&buffer[0]))

	c_size := (C.gsize)(size)

	c_blocking := toCBool(blocking)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_send_with_blocking(c_socket, c_buffer, c_size, c_blocking, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_socket_set_blocking(socket unsafe.Pointer, blocking bool) {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_blocking := toCBool(blocking)

	C.g_socket_set_blocking(c_socket, c_blocking)
}

func Fn_g_socket_set_keepalive(socket unsafe.Pointer, keepalive bool) {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_keepalive := toCBool(keepalive)

	C.g_socket_set_keepalive(c_socket, c_keepalive)
}

func Fn_g_socket_set_listen_backlog(socket unsafe.Pointer, backlog int) {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_backlog := (C.gint)(backlog)

	C.g_socket_set_listen_backlog(c_socket, c_backlog)
}

func Fn_g_socket_set_timeout(socket unsafe.Pointer, timeout uint) {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_timeout := (C.guint)(timeout)

	C.g_socket_set_timeout(c_socket, c_timeout)
}

func Fn_g_socket_shutdown(socket unsafe.Pointer, shutdownRead bool, shutdownWrite bool, error unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_shutdownRead := toCBool(shutdownRead)

	c_shutdownWrite := toCBool(shutdownWrite)

	cError := (**C.GError)(error)

	ret := C.g_socket_shutdown(c_socket, c_shutdownRead, c_shutdownWrite, cError)

	return toGoBool(ret)
}

func Fn_g_socket_speaks_ipv4(socket unsafe.Pointer) bool {
	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	ret := C.g_socket_speaks_ipv4(c_socket)

	return toGoBool(ret)
}

func Fn_g_socket_address_new_from_native(native unsafe.Pointer, len_ uint64) unsafe.Pointer {
	c_native := (C.gpointer)(native)

	c_len_ := (C.gsize)(len_)

	ret := C.g_socket_address_new_from_native(c_native, c_len_)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_address_get_family(address unsafe.Pointer) int {
	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	ret := C.g_socket_address_get_family(c_address)

	return (int)(ret)
}

func Fn_g_socket_address_get_native_size(address unsafe.Pointer) uint64 {
	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	ret := C.g_socket_address_get_native_size(c_address)

	return (uint64)(ret)
}

func Fn_g_socket_address_to_native(address unsafe.Pointer, dest unsafe.Pointer, destlen uint64, error unsafe.Pointer) bool {
	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	c_dest := (C.gpointer)(dest)

	c_destlen := (C.gsize)(destlen)

	cError := (**C.GError)(error)

	ret := C.g_socket_address_to_native(c_address, c_dest, c_destlen, cError)

	return toGoBool(ret)
}

func Fn_g_socket_address_enumerator_next(enumerator unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_enumerator := (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next(c_enumerator, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

func Fn_g_socket_address_enumerator_next_finish(enumerator unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_enumerator := (*C.GSocketAddressEnumerator)(unsafe.Pointer(enumerator))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next_finish(c_enumerator, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_new() unsafe.Pointer {
	ret := C.g_socket_client_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_add_application_proxy(client unsafe.Pointer, protocol string) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_protocol := (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(c_protocol))

	C.g_socket_client_add_application_proxy(c_client, c_protocol)
}

func Fn_g_socket_client_connect(client unsafe.Pointer, connectable unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_connectable := (*C.GSocketConnectable)(unsafe.Pointer(connectable))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect(c_client, c_connectable, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_finish(client unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_finish(c_client, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_host(client unsafe.Pointer, hostAndPort string, defaultPort uint16, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_hostAndPort := (*C.gchar)(C.CString(hostAndPort))
	defer C.free(unsafe.Pointer(c_hostAndPort))

	c_defaultPort := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host(c_client, c_hostAndPort, c_defaultPort, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_host_finish(client unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_host_finish(c_client, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_service(client unsafe.Pointer, domain string, service string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_domain := (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(c_domain))

	c_service := (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(c_service))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_service(c_client, c_domain, c_service, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_service_finish(client unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_service_finish(c_client, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_connect_to_uri(client unsafe.Pointer, uri string, defaultPort uint16, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_uri := (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(c_uri))

	c_defaultPort := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri(c_client, c_uri, c_defaultPort, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

func Fn_g_socket_client_connect_to_uri_finish(client unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_socket_client_connect_to_uri_finish(c_client, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_enable_proxy(client unsafe.Pointer) bool {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_enable_proxy(c_client)

	return toGoBool(ret)
}

func Fn_g_socket_client_get_family(client unsafe.Pointer) int {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_family(c_client)

	return (int)(ret)
}

func Fn_g_socket_client_get_local_address(client unsafe.Pointer) unsafe.Pointer {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_local_address(c_client)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_get_protocol(client unsafe.Pointer) int {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_protocol(c_client)

	return (int)(ret)
}

func Fn_g_socket_client_get_socket_type(client unsafe.Pointer) int {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_socket_type(c_client)

	return (int)(ret)
}

func Fn_g_socket_client_get_timeout(client unsafe.Pointer) uint {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	ret := C.g_socket_client_get_timeout(c_client)

	return (uint)(ret)
}

func Fn_g_socket_client_set_enable_proxy(client unsafe.Pointer, enable bool) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_enable := toCBool(enable)

	C.g_socket_client_set_enable_proxy(c_client, c_enable)
}

func Fn_g_socket_client_set_family(client unsafe.Pointer, family int) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_family := (C.GSocketFamily)(family)

	C.g_socket_client_set_family(c_client, c_family)
}

func Fn_g_socket_client_set_local_address(client unsafe.Pointer, address unsafe.Pointer) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_address := (*C.GSocketAddress)(unsafe.Pointer(address))

	C.g_socket_client_set_local_address(c_client, c_address)
}

func Fn_g_socket_client_set_protocol(client unsafe.Pointer, protocol int) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_protocol := (C.GSocketProtocol)(protocol)

	C.g_socket_client_set_protocol(c_client, c_protocol)
}

func Fn_g_socket_client_set_socket_type(client unsafe.Pointer, type_ int) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_type_ := (C.GSocketType)(type_)

	C.g_socket_client_set_socket_type(c_client, c_type_)
}

func Fn_g_socket_client_set_timeout(client unsafe.Pointer, timeout uint) {
	c_client := (*C.GSocketClient)(unsafe.Pointer(client))

	c_timeout := (C.guint)(timeout)

	C.g_socket_client_set_timeout(c_client, c_timeout)
}

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

func Fn_g_socket_connection_get_local_address(connection unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GSocketConnection)(unsafe.Pointer(connection))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_local_address(c_connection, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_remote_address(connection unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GSocketConnection)(unsafe.Pointer(connection))

	cError := (**C.GError)(error)

	ret := C.g_socket_connection_get_remote_address(c_connection, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_get_socket(connection unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GSocketConnection)(unsafe.Pointer(connection))

	ret := C.g_socket_connection_get_socket(c_connection)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connection_factory_lookup_type(family int, type_ int, protocolId int) uint64 {
	c_family := (C.GSocketFamily)(family)

	c_type_ := (C.GSocketType)(type_)

	c_protocolId := (C.gint)(protocolId)

	ret := C.g_socket_connection_factory_lookup_type(c_family, c_type_, c_protocolId)

	return (uint64)(ret)
}

func Fn_g_socket_connection_factory_register_type(gType uint64, family int, type_ int, protocol int) {
	c_gType := (C.GType)(gType)

	c_family := (C.GSocketFamily)(family)

	c_type_ := (C.GSocketType)(type_)

	c_protocol := (C.gint)(protocol)

	C.g_socket_connection_factory_register_type(c_gType, c_family, c_type_, c_protocol)
}

func Fn_g_socket_control_message_get_level(message unsafe.Pointer) int {
	c_message := (*C.GSocketControlMessage)(unsafe.Pointer(message))

	ret := C.g_socket_control_message_get_level(c_message)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_msg_type(message unsafe.Pointer) int {
	c_message := (*C.GSocketControlMessage)(unsafe.Pointer(message))

	ret := C.g_socket_control_message_get_msg_type(c_message)

	return (int)(ret)
}

func Fn_g_socket_control_message_get_size(message unsafe.Pointer) uint64 {
	c_message := (*C.GSocketControlMessage)(unsafe.Pointer(message))

	ret := C.g_socket_control_message_get_size(c_message)

	return (uint64)(ret)
}

func Fn_g_socket_control_message_serialize(message unsafe.Pointer, data unsafe.Pointer) {
	c_message := (*C.GSocketControlMessage)(unsafe.Pointer(message))

	c_data := (C.gpointer)(data)

	C.g_socket_control_message_serialize(c_message, c_data)
}

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

func Fn_g_socket_listener_new() unsafe.Pointer {
	ret := C.g_socket_listener_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_listener_accept : parameter 'source_object' is non array with indirect count > 1

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_finish : parameter 'source_object' is non array with indirect count > 1

// UNSUPPORTED : g_socket_listener_accept_socket : parameter 'source_object' is non array with indirect count > 1

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_socket_finish : parameter 'source_object' is non array with indirect count > 1

// UNSUPPORTED : g_socket_listener_add_address : parameter 'effective_address' is non array with indirect count > 1

func Fn_g_socket_listener_add_any_inet_port(listener unsafe.Pointer, sourceObject unsafe.Pointer, error unsafe.Pointer) uint16 {
	c_listener := (*C.GSocketListener)(unsafe.Pointer(listener))

	c_sourceObject := (*C.GObject)(unsafe.Pointer(sourceObject))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_any_inet_port(c_listener, c_sourceObject, cError)

	return (uint16)(ret)
}

func Fn_g_socket_listener_add_inet_port(listener unsafe.Pointer, port uint16, sourceObject unsafe.Pointer, error unsafe.Pointer) bool {
	c_listener := (*C.GSocketListener)(unsafe.Pointer(listener))

	c_port := (C.guint16)(port)

	c_sourceObject := (*C.GObject)(unsafe.Pointer(sourceObject))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_inet_port(c_listener, c_port, c_sourceObject, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_add_socket(listener unsafe.Pointer, socket unsafe.Pointer, sourceObject unsafe.Pointer, error unsafe.Pointer) bool {
	c_listener := (*C.GSocketListener)(unsafe.Pointer(listener))

	c_socket := (*C.GSocket)(unsafe.Pointer(socket))

	c_sourceObject := (*C.GObject)(unsafe.Pointer(sourceObject))

	cError := (**C.GError)(error)

	ret := C.g_socket_listener_add_socket(c_listener, c_socket, c_sourceObject, cError)

	return toGoBool(ret)
}

func Fn_g_socket_listener_close(listener unsafe.Pointer) {
	c_listener := (*C.GSocketListener)(unsafe.Pointer(listener))

	C.g_socket_listener_close(c_listener)
}

func Fn_g_socket_listener_set_backlog(listener unsafe.Pointer, listenBacklog int) {
	c_listener := (*C.GSocketListener)(unsafe.Pointer(listener))

	c_listenBacklog := (C.int)(listenBacklog)

	C.g_socket_listener_set_backlog(c_listener, c_listenBacklog)
}

func Fn_g_socket_service_new() unsafe.Pointer {
	ret := C.g_socket_service_new()

	return unsafe.Pointer(ret)
}

func Fn_g_socket_service_is_active(service unsafe.Pointer) bool {
	c_service := (*C.GSocketService)(unsafe.Pointer(service))

	ret := C.g_socket_service_is_active(c_service)

	return toGoBool(ret)
}

func Fn_g_socket_service_start(service unsafe.Pointer) {
	c_service := (*C.GSocketService)(unsafe.Pointer(service))

	C.g_socket_service_start(c_service)
}

func Fn_g_socket_service_stop(service unsafe.Pointer) {
	c_service := (*C.GSocketService)(unsafe.Pointer(service))

	C.g_socket_service_stop(c_service)
}

// UNSUPPORTED : g_subprocess_new : parameter 'error' is non array with indirect count > 1

// UNSUPPORTED : g_subprocess_communicate : parameter 'stdout_buf' is non array with indirect count > 1

// UNSUPPORTED : g_subprocess_communicate_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_communicate_finish : parameter 'stdout_buf' is non array with indirect count > 1

// UNSUPPORTED : g_subprocess_communicate_utf8 : parameter 'stdout_buf' is non array with indirect count > 1

// UNSUPPORTED : g_subprocess_communicate_utf8_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_communicate_utf8_finish : parameter 'stdout_buf' is non array with indirect count > 1

// UNSUPPORTED : g_subprocess_wait_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_check_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : parameter 'child_setup' is callback

// UNSUPPORTED : g_subprocess_launcher_spawn : parameter 'error' is non array with indirect count > 1

func Fn_g_subprocess_launcher_take_fd(self unsafe.Pointer, sourceFd int, targetFd int) {
	c_self := (*C.GSubprocessLauncher)(unsafe.Pointer(self))

	c_sourceFd := (C.gint)(sourceFd)

	c_targetFd := (C.gint)(targetFd)

	C.g_subprocess_launcher_take_fd(c_self, c_sourceFd, c_targetFd)
}

// UNSUPPORTED : g_task_new : parameter 'callback' is callback

// UNSUPPORTED : g_task_attach_source : parameter 'callback' is callback

// UNSUPPORTED : g_task_return_pointer : parameter 'result_destroy' is callback

// UNSUPPORTED : g_task_run_in_thread : parameter 'task_func' is callback

// UNSUPPORTED : g_task_run_in_thread_sync : parameter 'task_func' is callback

// UNSUPPORTED : g_task_set_task_data : parameter 'task_data_destroy' is callback

// UNSUPPORTED : g_task_report_error : parameter 'callback' is callback

// UNSUPPORTED : g_task_report_new_error : parameter 'callback' is callback

func Fn_g_tcp_connection_get_graceful_disconnect(connection unsafe.Pointer) bool {
	c_connection := (*C.GTcpConnection)(unsafe.Pointer(connection))

	ret := C.g_tcp_connection_get_graceful_disconnect(c_connection)

	return toGoBool(ret)
}

func Fn_g_tcp_connection_set_graceful_disconnect(connection unsafe.Pointer, gracefulDisconnect bool) {
	c_connection := (*C.GTcpConnection)(unsafe.Pointer(connection))

	c_gracefulDisconnect := toCBool(gracefulDisconnect)

	C.g_tcp_connection_set_graceful_disconnect(c_connection, c_gracefulDisconnect)
}

func Fn_g_tcp_wrapper_connection_get_base_io_stream(conn unsafe.Pointer) unsafe.Pointer {
	c_conn := (*C.GTcpWrapperConnection)(unsafe.Pointer(conn))

	ret := C.g_tcp_wrapper_connection_get_base_io_stream(c_conn)

	return unsafe.Pointer(ret)
}

func Fn_g_test_dbus_new(flags int) unsafe.Pointer {
	c_flags := (C.GTestDBusFlags)(flags)

	ret := C.g_test_dbus_new(c_flags)

	return unsafe.Pointer(ret)
}

func Fn_g_test_dbus_add_service_dir(self unsafe.Pointer, path string) {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	C.g_test_dbus_add_service_dir(c_self, c_path)
}

func Fn_g_test_dbus_down(self unsafe.Pointer) {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	C.g_test_dbus_down(c_self)
}

func Fn_g_test_dbus_get_bus_address(self unsafe.Pointer) string {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	ret := C.g_test_dbus_get_bus_address(c_self)

	return C.GoString(ret)
}

func Fn_g_test_dbus_get_flags(self unsafe.Pointer) int {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	ret := C.g_test_dbus_get_flags(c_self)

	return (int)(ret)
}

func Fn_g_test_dbus_stop(self unsafe.Pointer) {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	C.g_test_dbus_stop(c_self)
}

func Fn_g_test_dbus_up(self unsafe.Pointer) {
	c_self := (*C.GTestDBus)(unsafe.Pointer(self))

	C.g_test_dbus_up(c_self)
}

func Fn_g_test_dbus_unset() {
	C.g_test_dbus_unset()
}

func Fn_g_themed_icon_new(iconname string) unsafe.Pointer {
	c_iconname := (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(c_iconname))

	ret := C.g_themed_icon_new(c_iconname)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_new_from_names(iconnames []string, len_ int) unsafe.Pointer {
	iconnamesLen := len(iconnames)
	c_iconnamesArray := C.malloc((C.ulong)(iconnamesLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_iconnamesArray))
	iconnamesSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_iconnamesArray))[:iconnamesLen:iconnamesLen]
	for iconnamesi, iconnamesString := range iconnames {
		iconnamesSlice[iconnamesi] = (*C.gchar)(C.CString(iconnamesString))
		defer C.free(unsafe.Pointer(iconnamesSlice[iconnamesi]))
	}
	c_iconnames := &iconnamesSlice[0]

	c_len_ := (C.int)(len_)

	ret := C.g_themed_icon_new_from_names(c_iconnames, c_len_)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_new_with_default_fallbacks(iconname string) unsafe.Pointer {
	c_iconname := (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(c_iconname))

	ret := C.g_themed_icon_new_with_default_fallbacks(c_iconname)

	return unsafe.Pointer(ret)
}

func Fn_g_themed_icon_append_name(icon unsafe.Pointer, iconname string) {
	c_icon := (*C.GThemedIcon)(unsafe.Pointer(icon))

	c_iconname := (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_append_name(c_icon, c_iconname)
}

// UNSUPPORTED : g_themed_icon_get_names : no array length

func Fn_g_themed_icon_prepend_name(icon unsafe.Pointer, iconname string) {
	c_icon := (*C.GThemedIcon)(unsafe.Pointer(icon))

	c_iconname := (*C.char)(C.CString(iconname))
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_prepend_name(c_icon, c_iconname)
}

func Fn_g_threaded_socket_service_new(maxThreads int) unsafe.Pointer {
	c_maxThreads := (C.int)(maxThreads)

	ret := C.g_threaded_socket_service_new(c_maxThreads)

	return unsafe.Pointer(ret)
}

func Fn_g_tls_connection_get_use_system_certdb(conn unsafe.Pointer) bool {
	c_conn := (*C.GTlsConnection)(unsafe.Pointer(conn))

	ret := C.g_tls_connection_get_use_system_certdb(c_conn)

	return toGoBool(ret)
}

// UNSUPPORTED : g_tls_connection_handshake_async : parameter 'callback' is callback

func Fn_g_tls_connection_set_use_system_certdb(conn unsafe.Pointer, useSystemCertdb bool) {
	c_conn := (*C.GTlsConnection)(unsafe.Pointer(conn))

	c_useSystemCertdb := toCBool(useSystemCertdb)

	C.g_tls_connection_set_use_system_certdb(c_conn, c_useSystemCertdb)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_verify_chain_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_ask_password_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

func Fn_g_tls_password_new(flags int, description string) unsafe.Pointer {
	c_flags := (C.GTlsPasswordFlags)(flags)

	c_description := (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(c_description))

	ret := C.g_tls_password_new(c_flags, c_description)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_password_get_value : blacklisted

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

func Fn_g_unix_connection_receive_credentials(connection unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_connection := (*C.GUnixConnection)(unsafe.Pointer(connection))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_credentials(c_connection, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_receive_fd(connection unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) int {
	c_connection := (*C.GUnixConnection)(unsafe.Pointer(connection))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_fd(c_connection, c_cancellable, cError)

	return (int)(ret)
}

func Fn_g_unix_connection_send_credentials(connection unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GUnixConnection)(unsafe.Pointer(connection))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_credentials(c_connection, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

func Fn_g_unix_connection_send_fd(connection unsafe.Pointer, fd int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_connection := (*C.GUnixConnection)(unsafe.Pointer(connection))

	c_fd := (C.gint)(fd)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_fd(c_connection, c_fd, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_unix_credentials_message_new() unsafe.Pointer {
	ret := C.g_unix_credentials_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_new_with_credentials(credentials unsafe.Pointer) unsafe.Pointer {
	c_credentials := (*C.GCredentials)(unsafe.Pointer(credentials))

	ret := C.g_unix_credentials_message_new_with_credentials(c_credentials)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_credentials_message_get_credentials(message unsafe.Pointer) unsafe.Pointer {
	c_message := (*C.GUnixCredentialsMessage)(unsafe.Pointer(message))

	ret := C.g_unix_credentials_message_get_credentials(c_message)

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

func Fn_g_unix_fd_list_new_from_array(fds []int, nFds int) unsafe.Pointer {
	c_fds := (*C.gint)(unsafe.Pointer(&fds[0]))

	c_nFds := (C.gint)(nFds)

	ret := C.g_unix_fd_list_new_from_array(c_fds, c_nFds)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_list_append(list unsafe.Pointer, fd int, error unsafe.Pointer) int {
	c_list := (*C.GUnixFDList)(unsafe.Pointer(list))

	c_fd := (C.gint)(fd)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_list_append(c_list, c_fd, cError)

	return (int)(ret)
}

func Fn_g_unix_fd_list_get(list unsafe.Pointer, index int, error unsafe.Pointer) int {
	c_list := (*C.GUnixFDList)(unsafe.Pointer(list))

	c_index := (C.gint)(index)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_list_get(c_list, c_index, cError)

	return (int)(ret)
}

func Fn_g_unix_fd_list_get_length(list unsafe.Pointer) int {
	c_list := (*C.GUnixFDList)(unsafe.Pointer(list))

	ret := C.g_unix_fd_list_get_length(c_list)

	return (int)(ret)
}

func Fn_g_unix_fd_list_peek_fds(list unsafe.Pointer, length *int) []int {
	c_list := (*C.GUnixFDList)(unsafe.Pointer(list))

	c_length := (*C.gint)(unsafe.Pointer(length))

	ret := C.g_unix_fd_list_peek_fds(c_list, c_length)

	retLen := int(*c_length)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_unix_fd_list_steal_fds(list unsafe.Pointer, length *int) []int {
	c_list := (*C.GUnixFDList)(unsafe.Pointer(list))

	c_length := (*C.gint)(unsafe.Pointer(length))

	ret := C.g_unix_fd_list_steal_fds(c_list, c_length)

	retLen := int(*c_length)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_unix_fd_message_new() unsafe.Pointer {
	ret := C.g_unix_fd_message_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_new_with_fd_list(fdList unsafe.Pointer) unsafe.Pointer {
	c_fdList := (*C.GUnixFDList)(unsafe.Pointer(fdList))

	ret := C.g_unix_fd_message_new_with_fd_list(c_fdList)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_append_fd(message unsafe.Pointer, fd int, error unsafe.Pointer) bool {
	c_message := (*C.GUnixFDMessage)(unsafe.Pointer(message))

	c_fd := (C.gint)(fd)

	cError := (**C.GError)(error)

	ret := C.g_unix_fd_message_append_fd(c_message, c_fd, cError)

	return toGoBool(ret)
}

func Fn_g_unix_fd_message_get_fd_list(message unsafe.Pointer) unsafe.Pointer {
	c_message := (*C.GUnixFDMessage)(unsafe.Pointer(message))

	ret := C.g_unix_fd_message_get_fd_list(c_message)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_fd_message_steal_fds(message unsafe.Pointer, length *int) []int {
	c_message := (*C.GUnixFDMessage)(unsafe.Pointer(message))

	c_length := (*C.gint)(unsafe.Pointer(length))

	ret := C.g_unix_fd_message_steal_fds(c_message, c_length)

	retLen := int(*c_length)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_g_unix_input_stream_new(fd int, closeFd bool) unsafe.Pointer {
	c_fd := (C.gint)(fd)

	c_closeFd := toCBool(closeFd)

	ret := C.g_unix_input_stream_new(c_fd, c_closeFd)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_input_stream_get_close_fd(stream unsafe.Pointer) bool {
	c_stream := (*C.GUnixInputStream)(unsafe.Pointer(stream))

	ret := C.g_unix_input_stream_get_close_fd(c_stream)

	return toGoBool(ret)
}

func Fn_g_unix_input_stream_get_fd(stream unsafe.Pointer) int {
	c_stream := (*C.GUnixInputStream)(unsafe.Pointer(stream))

	ret := C.g_unix_input_stream_get_fd(c_stream)

	return (int)(ret)
}

func Fn_g_unix_input_stream_set_close_fd(stream unsafe.Pointer, closeFd bool) {
	c_stream := (*C.GUnixInputStream)(unsafe.Pointer(stream))

	c_closeFd := toCBool(closeFd)

	C.g_unix_input_stream_set_close_fd(c_stream, c_closeFd)
}

func Fn_g_unix_mount_monitor_new() unsafe.Pointer {
	ret := C.g_unix_mount_monitor_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_monitor_set_rate_limit(mountMonitor unsafe.Pointer, limitMsec int) {
	c_mountMonitor := (*C.GUnixMountMonitor)(unsafe.Pointer(mountMonitor))

	c_limitMsec := (C.int)(limitMsec)

	C.g_unix_mount_monitor_set_rate_limit(c_mountMonitor, c_limitMsec)
}

func Fn_g_unix_output_stream_new(fd int, closeFd bool) unsafe.Pointer {
	c_fd := (C.gint)(fd)

	c_closeFd := toCBool(closeFd)

	ret := C.g_unix_output_stream_new(c_fd, c_closeFd)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_output_stream_get_close_fd(stream unsafe.Pointer) bool {
	c_stream := (*C.GUnixOutputStream)(unsafe.Pointer(stream))

	ret := C.g_unix_output_stream_get_close_fd(c_stream)

	return toGoBool(ret)
}

func Fn_g_unix_output_stream_get_fd(stream unsafe.Pointer) int {
	c_stream := (*C.GUnixOutputStream)(unsafe.Pointer(stream))

	ret := C.g_unix_output_stream_get_fd(c_stream)

	return (int)(ret)
}

func Fn_g_unix_output_stream_set_close_fd(stream unsafe.Pointer, closeFd bool) {
	c_stream := (*C.GUnixOutputStream)(unsafe.Pointer(stream))

	c_closeFd := toCBool(closeFd)

	C.g_unix_output_stream_set_close_fd(c_stream, c_closeFd)
}

func Fn_g_unix_socket_address_new(path string) unsafe.Pointer {
	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	ret := C.g_unix_socket_address_new(c_path)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_new_abstract(path []int8, pathLen int) unsafe.Pointer {
	c_path := (*C.gchar)(unsafe.Pointer(&path[0]))

	c_pathLen := (C.gint)(pathLen)

	ret := C.g_unix_socket_address_new_abstract(c_path, c_pathLen)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_new_with_type(path []int8, pathLen int, type_ int) unsafe.Pointer {
	c_path := (*C.gchar)(unsafe.Pointer(&path[0]))

	c_pathLen := (C.gint)(pathLen)

	c_type_ := (C.GUnixSocketAddressType)(type_)

	ret := C.g_unix_socket_address_new_with_type(c_path, c_pathLen, c_type_)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_get_address_type(address unsafe.Pointer) int {
	c_address := (*C.GUnixSocketAddress)(unsafe.Pointer(address))

	ret := C.g_unix_socket_address_get_address_type(c_address)

	return (int)(ret)
}

func Fn_g_unix_socket_address_get_is_abstract(address unsafe.Pointer) bool {
	c_address := (*C.GUnixSocketAddress)(unsafe.Pointer(address))

	ret := C.g_unix_socket_address_get_is_abstract(c_address)

	return toGoBool(ret)
}

func Fn_g_unix_socket_address_get_path(address unsafe.Pointer) string {
	c_address := (*C.GUnixSocketAddress)(unsafe.Pointer(address))

	ret := C.g_unix_socket_address_get_path(c_address)

	return C.GoString(ret)
}

func Fn_g_unix_socket_address_get_path_len(address unsafe.Pointer) uint64 {
	c_address := (*C.GUnixSocketAddress)(unsafe.Pointer(address))

	ret := C.g_unix_socket_address_get_path_len(c_address)

	return (uint64)(ret)
}

func Fn_g_unix_socket_address_abstract_names_supported() bool {
	ret := C.g_unix_socket_address_abstract_names_supported()

	return toGoBool(ret)
}

func Fn_g_vfs_get_file_for_path(vfs unsafe.Pointer, path string) unsafe.Pointer {
	c_vfs := (*C.GVfs)(unsafe.Pointer(vfs))

	c_path := (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	ret := C.g_vfs_get_file_for_path(c_vfs, c_path)

	return unsafe.Pointer(ret)
}

func Fn_g_vfs_get_file_for_uri(vfs unsafe.Pointer, uri string) unsafe.Pointer {
	c_vfs := (*C.GVfs)(unsafe.Pointer(vfs))

	c_uri := (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(c_uri))

	ret := C.g_vfs_get_file_for_uri(c_vfs, c_uri)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

func Fn_g_vfs_is_active(vfs unsafe.Pointer) bool {
	c_vfs := (*C.GVfs)(unsafe.Pointer(vfs))

	ret := C.g_vfs_is_active(c_vfs)

	return toGoBool(ret)
}

func Fn_g_vfs_parse_name(vfs unsafe.Pointer, parseName string) unsafe.Pointer {
	c_vfs := (*C.GVfs)(unsafe.Pointer(vfs))

	c_parseName := (*C.char)(C.CString(parseName))
	defer C.free(unsafe.Pointer(c_parseName))

	ret := C.g_vfs_parse_name(c_vfs, c_parseName)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

func Fn_g_vfs_get_default() unsafe.Pointer {
	ret := C.g_vfs_get_default()

	return unsafe.Pointer(ret)
}

func Fn_g_vfs_get_local() unsafe.Pointer {
	ret := C.g_vfs_get_local()

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_connected_drives(volumeMonitor unsafe.Pointer) unsafe.Pointer {
	c_volumeMonitor := (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor))

	ret := C.g_volume_monitor_get_connected_drives(c_volumeMonitor)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_mount_for_uuid(volumeMonitor unsafe.Pointer, uuid string) unsafe.Pointer {
	c_volumeMonitor := (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor))

	c_uuid := (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(c_uuid))

	ret := C.g_volume_monitor_get_mount_for_uuid(c_volumeMonitor, c_uuid)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_mounts(volumeMonitor unsafe.Pointer) unsafe.Pointer {
	c_volumeMonitor := (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor))

	ret := C.g_volume_monitor_get_mounts(c_volumeMonitor)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_volume_for_uuid(volumeMonitor unsafe.Pointer, uuid string) unsafe.Pointer {
	c_volumeMonitor := (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor))

	c_uuid := (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(c_uuid))

	ret := C.g_volume_monitor_get_volume_for_uuid(c_volumeMonitor, c_uuid)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get_volumes(volumeMonitor unsafe.Pointer) unsafe.Pointer {
	c_volumeMonitor := (*C.GVolumeMonitor)(unsafe.Pointer(volumeMonitor))

	ret := C.g_volume_monitor_get_volumes(c_volumeMonitor)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_adopt_orphan_mount(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_volume_monitor_adopt_orphan_mount(c_mount)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_monitor_get() unsafe.Pointer {
	ret := C.g_volume_monitor_get()

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_new(format int, level int) unsafe.Pointer {
	c_format := (C.GZlibCompressorFormat)(format)

	c_level := (C.int)(level)

	ret := C.g_zlib_compressor_new(c_format, c_level)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_get_file_info(compressor unsafe.Pointer) unsafe.Pointer {
	c_compressor := (*C.GZlibCompressor)(unsafe.Pointer(compressor))

	ret := C.g_zlib_compressor_get_file_info(c_compressor)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_compressor_set_file_info(compressor unsafe.Pointer, fileInfo unsafe.Pointer) {
	c_compressor := (*C.GZlibCompressor)(unsafe.Pointer(compressor))

	c_fileInfo := (*C.GFileInfo)(unsafe.Pointer(fileInfo))

	C.g_zlib_compressor_set_file_info(c_compressor, c_fileInfo)
}

func Fn_g_zlib_decompressor_new(format int) unsafe.Pointer {
	c_format := (C.GZlibCompressorFormat)(format)

	ret := C.g_zlib_decompressor_new(c_format)

	return unsafe.Pointer(ret)
}

func Fn_g_zlib_decompressor_get_file_info(decompressor unsafe.Pointer) unsafe.Pointer {
	c_decompressor := (*C.GZlibDecompressor)(unsafe.Pointer(decompressor))

	ret := C.g_zlib_decompressor_get_file_info(c_decompressor)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_action_group_list_actions : no array length

// UNSUPPORTED : g_action_group_query_action : parameter 'parameter_type' is non array with indirect count > 1

func Fn_g_app_info_add_supports_type(appinfo unsafe.Pointer, contentType string, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_contentType := (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(c_contentType))

	cError := (**C.GError)(error)

	ret := C.g_app_info_add_supports_type(c_appinfo, c_contentType, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_can_delete(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_can_delete(c_appinfo)

	return toGoBool(ret)
}

func Fn_g_app_info_can_remove_supports_type(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_can_remove_supports_type(c_appinfo)

	return toGoBool(ret)
}

func Fn_g_app_info_delete(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_delete(c_appinfo)

	return toGoBool(ret)
}

func Fn_g_app_info_dup(appinfo unsafe.Pointer) unsafe.Pointer {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_dup(c_appinfo)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_equal(appinfo1 unsafe.Pointer, appinfo2 unsafe.Pointer) bool {
	c_appinfo1 := (*C.GAppInfo)(unsafe.Pointer(appinfo1))

	c_appinfo2 := (*C.GAppInfo)(unsafe.Pointer(appinfo2))

	ret := C.g_app_info_equal(c_appinfo1, c_appinfo2)

	return toGoBool(ret)
}

func Fn_g_app_info_get_commandline(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_commandline(c_appinfo)

	return C.GoString(ret)
}

func Fn_g_app_info_get_description(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_description(c_appinfo)

	return C.GoString(ret)
}

func Fn_g_app_info_get_display_name(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_display_name(c_appinfo)

	return C.GoString(ret)
}

func Fn_g_app_info_get_executable(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_executable(c_appinfo)

	return C.GoString(ret)
}

func Fn_g_app_info_get_icon(appinfo unsafe.Pointer) unsafe.Pointer {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_icon(c_appinfo)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_id(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_id(c_appinfo)

	return C.GoString(ret)
}

func Fn_g_app_info_get_name(appinfo unsafe.Pointer) string {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_get_name(c_appinfo)

	return C.GoString(ret)
}

// UNSUPPORTED : g_app_info_get_supported_types : no array length

func Fn_g_app_info_launch(appinfo unsafe.Pointer, files unsafe.Pointer, context unsafe.Pointer, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_files := (*C.GList)(unsafe.Pointer(files))

	c_context := (*C.GAppLaunchContext)(unsafe.Pointer(context))

	cError := (**C.GError)(error)

	ret := C.g_app_info_launch(c_appinfo, c_files, c_context, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_launch_uris(appinfo unsafe.Pointer, uris unsafe.Pointer, context unsafe.Pointer, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_uris := (*C.GList)(unsafe.Pointer(uris))

	c_context := (*C.GAppLaunchContext)(unsafe.Pointer(context))

	cError := (**C.GError)(error)

	ret := C.g_app_info_launch_uris(c_appinfo, c_uris, c_context, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

func Fn_g_app_info_remove_supports_type(appinfo unsafe.Pointer, contentType string, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_contentType := (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(c_contentType))

	cError := (**C.GError)(error)

	ret := C.g_app_info_remove_supports_type(c_appinfo, c_contentType, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_default_for_extension(appinfo unsafe.Pointer, extension string, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_extension := (*C.char)(C.CString(extension))
	defer C.free(unsafe.Pointer(c_extension))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_default_for_extension(c_appinfo, c_extension, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_default_for_type(appinfo unsafe.Pointer, contentType string, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_contentType := (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(c_contentType))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_default_for_type(c_appinfo, c_contentType, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_set_as_last_used_for_type(appinfo unsafe.Pointer, contentType string, error unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	c_contentType := (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(c_contentType))

	cError := (**C.GError)(error)

	ret := C.g_app_info_set_as_last_used_for_type(c_appinfo, c_contentType, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_should_show(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_should_show(c_appinfo)

	return toGoBool(ret)
}

func Fn_g_app_info_supports_files(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_supports_files(c_appinfo)

	return toGoBool(ret)
}

func Fn_g_app_info_supports_uris(appinfo unsafe.Pointer) bool {
	c_appinfo := (*C.GAppInfo)(unsafe.Pointer(appinfo))

	ret := C.g_app_info_supports_uris(c_appinfo)

	return toGoBool(ret)
}

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

func Fn_g_async_initable_init_finish(initable unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) bool {
	c_initable := (*C.GAsyncInitable)(unsafe.Pointer(initable))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_init_finish(c_initable, c_res, cError)

	return toGoBool(ret)
}

func Fn_g_async_initable_new_finish(initable unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_initable := (*C.GAsyncInitable)(unsafe.Pointer(initable))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_async_initable_new_finish(c_initable, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_async_result_get_source_object(res unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	ret := C.g_async_result_get_source_object(c_res)

	return unsafe.Pointer(ret)
}

func Fn_g_async_result_get_user_data(res unsafe.Pointer) unsafe.Pointer {
	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	ret := C.g_async_result_get_user_data(c_res)

	return (unsafe.Pointer)(ret)
}

func Fn_g_converter_convert(converter unsafe.Pointer, inbuf []uint8, inbufSize uint64, outbuf []uint8, outbufSize uint64, flags int, bytesRead *uint64, bytesWritten *uint64, error unsafe.Pointer) int {
	c_converter := (*C.GConverter)(unsafe.Pointer(converter))

	c_inbuf := (unsafe.Pointer)(unsafe.Pointer(&inbuf[0]))

	c_inbufSize := (C.gsize)(inbufSize)

	c_outbuf := (unsafe.Pointer)(unsafe.Pointer(&outbuf[0]))

	c_outbufSize := (C.gsize)(outbufSize)

	c_flags := (C.GConverterFlags)(flags)

	c_bytesRead := (*C.gsize)(unsafe.Pointer(bytesRead))

	c_bytesWritten := (*C.gsize)(unsafe.Pointer(bytesWritten))

	cError := (**C.GError)(error)

	ret := C.g_converter_convert(c_converter, c_inbuf, c_inbufSize, c_outbuf, c_outbufSize, c_flags, c_bytesRead, c_bytesWritten, cError)

	return (int)(ret)
}

func Fn_g_converter_reset(converter unsafe.Pointer) {
	c_converter := (*C.GConverter)(unsafe.Pointer(converter))

	C.g_converter_reset(c_converter)
}

func Fn_g_desktop_app_info_lookup_get_default_for_uri_scheme(lookup unsafe.Pointer, uriScheme string) unsafe.Pointer {
	c_lookup := (*C.GDesktopAppInfoLookup)(unsafe.Pointer(lookup))

	c_uriScheme := (*C.char)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(c_uriScheme))

	ret := C.g_desktop_app_info_lookup_get_default_for_uri_scheme(c_lookup, c_uriScheme)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_can_eject(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_can_eject(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_can_poll_for_media(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_can_poll_for_media(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_can_start(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_can_start(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_can_start_degraded(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_can_start_degraded(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_can_stop(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_can_stop(c_drive)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

func Fn_g_drive_eject_finish(drive unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_finish(c_drive, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

func Fn_g_drive_eject_with_operation_finish(drive unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_with_operation_finish(c_drive, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

func Fn_g_drive_get_icon(drive unsafe.Pointer) unsafe.Pointer {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_get_icon(c_drive)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_get_identifier(drive unsafe.Pointer, kind string) string {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_kind := (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(c_kind))

	ret := C.g_drive_get_identifier(c_drive, c_kind)

	return C.GoString(ret)
}

func Fn_g_drive_get_name(drive unsafe.Pointer) string {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_get_name(c_drive)

	return C.GoString(ret)
}

func Fn_g_drive_get_start_stop_type(drive unsafe.Pointer) int {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_get_start_stop_type(c_drive)

	return (int)(ret)
}

func Fn_g_drive_get_volumes(drive unsafe.Pointer) unsafe.Pointer {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_get_volumes(c_drive)

	return unsafe.Pointer(ret)
}

func Fn_g_drive_has_media(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_has_media(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_has_volumes(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_has_volumes(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_is_media_check_automatic(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_is_media_check_automatic(c_drive)

	return toGoBool(ret)
}

func Fn_g_drive_is_media_removable(drive unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	ret := C.g_drive_is_media_removable(c_drive)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_poll_for_media : parameter 'callback' is callback

func Fn_g_drive_poll_for_media_finish(drive unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_drive_poll_for_media_finish(c_drive, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_start : parameter 'callback' is callback

func Fn_g_drive_start_finish(drive unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_drive_start_finish(c_drive, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_stop : parameter 'callback' is callback

func Fn_g_drive_stop_finish(drive unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_drive := (*C.GDrive)(unsafe.Pointer(drive))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_drive_stop_finish(c_drive, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_dtls_connection_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_shutdown_async : parameter 'callback' is callback

func Fn_g_file_append_to(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_append_to_async : parameter 'callback' is callback

func Fn_g_file_append_to_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_copy : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_copy_async : parameter 'progress_callback' is callback

func Fn_g_file_copy_attributes(source unsafe.Pointer, destination unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_source := (*C.GFile)(unsafe.Pointer(source))

	c_destination := (*C.GFile)(unsafe.Pointer(destination))

	c_flags := (C.GFileCopyFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_copy_attributes(c_source, c_destination, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_copy_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_copy_finish(c_file, c_res, cError)

	return toGoBool(ret)
}

func Fn_g_file_create(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_create(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_async : parameter 'callback' is callback

func Fn_g_file_create_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_create_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_create_readwrite(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_readwrite_async : parameter 'callback' is callback

func Fn_g_file_create_readwrite_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_create_readwrite_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_delete(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_delete(c_file, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_delete_async : parameter 'callback' is callback

func Fn_g_file_dup(file unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_dup(c_file)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_eject_mountable : parameter 'callback' is callback

func Fn_g_file_eject_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_eject_mountable_with_operation : parameter 'callback' is callback

func Fn_g_file_eject_mountable_with_operation_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_with_operation_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_file_enumerate_children(file unsafe.Pointer, attributes string, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerate_children(c_file, c_attributes, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_enumerate_children_async : parameter 'callback' is callback

func Fn_g_file_enumerate_children_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerate_children_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_equal(file1 unsafe.Pointer, file2 unsafe.Pointer) bool {
	c_file1 := (*C.GFile)(unsafe.Pointer(file1))

	c_file2 := (*C.GFile)(unsafe.Pointer(file2))

	ret := C.g_file_equal(c_file1, c_file2)

	return toGoBool(ret)
}

func Fn_g_file_find_enclosing_mount(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_find_enclosing_mount(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_find_enclosing_mount_async : parameter 'callback' is callback

func Fn_g_file_find_enclosing_mount_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_find_enclosing_mount_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_basename(file unsafe.Pointer) string {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_basename(c_file)

	return C.GoString(ret)
}

func Fn_g_file_get_child(file unsafe.Pointer, name string) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_name := (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.g_file_get_child(c_file, c_name)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_child_for_display_name(file unsafe.Pointer, displayName string, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_displayName := (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(c_displayName))

	cError := (**C.GError)(error)

	ret := C.g_file_get_child_for_display_name(c_file, c_displayName, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_parent(file unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_parent(c_file)

	return unsafe.Pointer(ret)
}

func Fn_g_file_get_parse_name(file unsafe.Pointer) string {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_parse_name(c_file)

	return C.GoString(ret)
}

func Fn_g_file_get_path(file unsafe.Pointer) string {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_path(c_file)

	return C.GoString(ret)
}

func Fn_g_file_get_relative_path(parent unsafe.Pointer, descendant unsafe.Pointer) string {
	c_parent := (*C.GFile)(unsafe.Pointer(parent))

	c_descendant := (*C.GFile)(unsafe.Pointer(descendant))

	ret := C.g_file_get_relative_path(c_parent, c_descendant)

	return C.GoString(ret)
}

func Fn_g_file_get_uri(file unsafe.Pointer) string {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_uri(c_file)

	return C.GoString(ret)
}

func Fn_g_file_get_uri_scheme(file unsafe.Pointer) string {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_get_uri_scheme(c_file)

	return C.GoString(ret)
}

func Fn_g_file_has_parent(file unsafe.Pointer, parent unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_parent := (*C.GFile)(unsafe.Pointer(parent))

	ret := C.g_file_has_parent(c_file, c_parent)

	return toGoBool(ret)
}

func Fn_g_file_has_prefix(file unsafe.Pointer, prefix unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_prefix := (*C.GFile)(unsafe.Pointer(prefix))

	ret := C.g_file_has_prefix(c_file, c_prefix)

	return toGoBool(ret)
}

func Fn_g_file_has_uri_scheme(file unsafe.Pointer, uriScheme string) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_uriScheme := (*C.char)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(c_uriScheme))

	ret := C.g_file_has_uri_scheme(c_file, c_uriScheme)

	return toGoBool(ret)
}

func Fn_g_file_hash(file unsafe.Pointer) uint {
	c_file := (C.gconstpointer)(file)

	ret := C.g_file_hash(c_file)

	return (uint)(ret)
}

func Fn_g_file_is_native(file unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_is_native(c_file)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_load_bytes : parameter 'etag_out' is non array with indirect count > 1

// UNSUPPORTED : g_file_load_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_bytes_finish : parameter 'etag_out' is non array with indirect count > 1

// UNSUPPORTED : g_file_load_contents : blacklisted

// UNSUPPORTED : g_file_load_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents_finish : blacklisted

// UNSUPPORTED : g_file_load_partial_contents_async : parameter 'read_more_callback' is callback

// UNSUPPORTED : g_file_load_partial_contents_finish : blacklisted

func Fn_g_file_make_directory(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory(c_file, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_make_directory_async : parameter 'callback' is callback

func Fn_g_file_make_directory_with_parents(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory_with_parents(c_file, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_make_symbolic_link(file unsafe.Pointer, symlinkValue string, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_symlinkValue := (*C.char)(C.CString(symlinkValue))
	defer C.free(unsafe.Pointer(c_symlinkValue))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_make_symbolic_link(c_file, c_symlinkValue, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_measure_disk_usage : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage_async : parameter 'progress_callback' is callback

func Fn_g_file_monitor(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_directory(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor_directory(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_monitor_file(file unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_monitor_file(c_file, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_mount_enclosing_volume : parameter 'callback' is callback

func Fn_g_file_mount_enclosing_volume_finish(location unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_location := (*C.GFile)(unsafe.Pointer(location))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_enclosing_volume_finish(c_location, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_mount_mountable : parameter 'callback' is callback

func Fn_g_file_mount_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_mountable_finish(c_file, c_result, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_move : parameter 'progress_callback' is callback

func Fn_g_file_open_readwrite(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_open_readwrite_async : parameter 'callback' is callback

func Fn_g_file_open_readwrite_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_open_readwrite_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_poll_mountable : parameter 'callback' is callback

func Fn_g_file_poll_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_poll_mountable_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_file_query_default_handler(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_query_default_handler(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_default_handler_async : parameter 'callback' is callback

func Fn_g_file_query_exists(file unsafe.Pointer, cancellable unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	ret := C.g_file_query_exists(c_file, c_cancellable)

	return toGoBool(ret)
}

func Fn_g_file_query_file_type(file unsafe.Pointer, flags int, cancellable unsafe.Pointer) int {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	ret := C.g_file_query_file_type(c_file, c_flags, c_cancellable)

	return (int)(ret)
}

func Fn_g_file_query_filesystem_info(file unsafe.Pointer, attributes string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_query_filesystem_info(c_file, c_attributes, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_filesystem_info_async : parameter 'callback' is callback

func Fn_g_file_query_filesystem_info_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_query_filesystem_info_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_info(file unsafe.Pointer, attributes string, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attributes := (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_query_info(c_file, c_attributes, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_info_async : parameter 'callback' is callback

func Fn_g_file_query_info_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_query_info_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_settable_attributes(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_query_settable_attributes(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_query_writable_namespaces(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_query_writable_namespaces(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_read(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_read(c_file, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_read_async : parameter 'callback' is callback

func Fn_g_file_read_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_read_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_replace(file unsafe.Pointer, etag *string, makeBackup bool, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	var c_etagValue (*C.char)
	if etag != nil {
		c_etagValue = (*C.char)(C.CString(*etag))
		defer C.free(unsafe.Pointer(c_etagValue))
	}
	c_etag := c_etagValue

	c_makeBackup := toCBool(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_replace(c_file, c_etag, c_makeBackup, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_replace_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents : parameter 'new_etag' is non array with indirect count > 1

// UNSUPPORTED : g_file_replace_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_finish : parameter 'new_etag' is non array with indirect count > 1

func Fn_g_file_replace_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_replace_readwrite(file unsafe.Pointer, etag *string, makeBackup bool, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	var c_etagValue (*C.char)
	if etag != nil {
		c_etagValue = (*C.char)(C.CString(*etag))
		defer C.free(unsafe.Pointer(c_etagValue))
	}
	c_etag := c_etagValue

	c_makeBackup := toCBool(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite(c_file, c_etag, c_makeBackup, c_flags, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_replace_readwrite_async : parameter 'callback' is callback

func Fn_g_file_replace_readwrite_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_readwrite_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_file_resolve_relative_path(file unsafe.Pointer, relativePath string) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_relativePath := (*C.char)(C.CString(relativePath))
	defer C.free(unsafe.Pointer(c_relativePath))

	ret := C.g_file_resolve_relative_path(c_file, c_relativePath)

	return unsafe.Pointer(ret)
}

func Fn_g_file_set_attribute(file unsafe.Pointer, attribute string, type_ int, valueP unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_type_ := (C.GFileAttributeType)(type_)

	c_valueP := (C.gpointer)(valueP)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute(c_file, c_attribute, c_type_, c_valueP, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_byte_string(file unsafe.Pointer, attribute string, value string, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_byte_string(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_int32(file unsafe.Pointer, attribute string, value int32, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_int32(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_int64(file unsafe.Pointer, attribute string, value int64, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_int64(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_string(file unsafe.Pointer, attribute string, value string, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_string(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_uint32(file unsafe.Pointer, attribute string, value uint32, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_uint32(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_attribute_uint64(file unsafe.Pointer, attribute string, value uint64, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_attribute := (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attribute_uint64(c_file, c_attribute, c_value, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_set_attributes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_set_attributes_finish : parameter 'info' is non array with indirect count > 1

func Fn_g_file_set_attributes_from_info(file unsafe.Pointer, info unsafe.Pointer, flags int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_info := (*C.GFileInfo)(unsafe.Pointer(info))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_attributes_from_info(c_file, c_info, c_flags, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_file_set_display_name(file unsafe.Pointer, displayName string, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_displayName := (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(c_displayName))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_set_display_name(c_file, c_displayName, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_set_display_name_async : parameter 'callback' is callback

func Fn_g_file_set_display_name_finish(file unsafe.Pointer, res unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_res := (*C.GAsyncResult)(unsafe.Pointer(res))

	cError := (**C.GError)(error)

	ret := C.g_file_set_display_name_finish(c_file, c_res, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_start_mountable : parameter 'callback' is callback

func Fn_g_file_start_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_start_mountable_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_stop_mountable : parameter 'callback' is callback

func Fn_g_file_stop_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_stop_mountable_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_file_supports_thread_contexts(file unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	ret := C.g_file_supports_thread_contexts(c_file)

	return toGoBool(ret)
}

func Fn_g_file_trash(file unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_file_trash(c_file, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_trash_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable : parameter 'callback' is callback

func Fn_g_file_unmount_mountable_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_unmount_mountable_with_operation : parameter 'callback' is callback

func Fn_g_file_unmount_mountable_with_operation_finish(file unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_file := (*C.GFile)(unsafe.Pointer(file))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_with_operation_finish(c_file, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_file_descriptor_based_get_fd(fdBased unsafe.Pointer) int {
	c_fdBased := (*C.GFileDescriptorBased)(unsafe.Pointer(fdBased))

	ret := C.g_file_descriptor_based_get_fd(c_fdBased)

	return (int)(ret)
}

func Fn_g_icon_equal(icon1 unsafe.Pointer, icon2 unsafe.Pointer) bool {
	c_icon1 := (*C.GIcon)(unsafe.Pointer(icon1))

	c_icon2 := (*C.GIcon)(unsafe.Pointer(icon2))

	ret := C.g_icon_equal(c_icon1, c_icon2)

	return toGoBool(ret)
}

func Fn_g_icon_to_string(icon unsafe.Pointer) string {
	c_icon := (*C.GIcon)(unsafe.Pointer(icon))

	ret := C.g_icon_to_string(c_icon)

	return C.GoString(ret)
}

func Fn_g_initable_init(initable unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_initable := (*C.GInitable)(unsafe.Pointer(initable))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_initable_init(c_initable, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_loadable_icon_load : parameter 'type' is non array with indirect count > 1

// UNSUPPORTED : g_loadable_icon_load_async : parameter 'callback' is callback

// UNSUPPORTED : g_loadable_icon_load_finish : parameter 'type' is non array with indirect count > 1

func Fn_g_mount_can_eject(mount unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_can_eject(c_mount)

	return toGoBool(ret)
}

func Fn_g_mount_can_unmount(mount unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_can_unmount(c_mount)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_eject : parameter 'callback' is callback

func Fn_g_mount_eject_finish(mount unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_finish(c_mount, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_eject_with_operation : parameter 'callback' is callback

func Fn_g_mount_eject_with_operation_finish(mount unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_with_operation_finish(c_mount, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_mount_get_default_location(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_default_location(c_mount)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_drive(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_drive(c_mount)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_icon(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_icon(c_mount)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_name(mount unsafe.Pointer) string {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_name(c_mount)

	return C.GoString(ret)
}

func Fn_g_mount_get_root(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_root(c_mount)

	return unsafe.Pointer(ret)
}

func Fn_g_mount_get_uuid(mount unsafe.Pointer) string {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_uuid(c_mount)

	return C.GoString(ret)
}

func Fn_g_mount_get_volume(mount unsafe.Pointer) unsafe.Pointer {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_get_volume(c_mount)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

func Fn_g_mount_is_shadowed(mount unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	ret := C.g_mount_is_shadowed(c_mount)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

func Fn_g_mount_remount_finish(mount unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_mount_remount_finish(c_mount, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_mount_shadow(mount unsafe.Pointer) {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	C.g_mount_shadow(c_mount)
}

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

func Fn_g_mount_unmount_finish(mount unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_finish(c_mount, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

func Fn_g_mount_unmount_with_operation_finish(mount unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_with_operation_finish(c_mount, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_mount_unshadow(mount unsafe.Pointer) {
	c_mount := (*C.GMount)(unsafe.Pointer(mount))

	C.g_mount_unshadow(c_mount)
}

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

func Fn_g_network_monitor_can_reach_finish(monitor unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_monitor := (*C.GNetworkMonitor)(unsafe.Pointer(monitor))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach_finish(c_monitor, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_pollable_input_stream_read_nonblocking(stream unsafe.Pointer, buffer []uint8, count uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GPollableInputStream)(unsafe.Pointer(stream))

	c_buffer := (unsafe.Pointer)(unsafe.Pointer(&buffer[0]))

	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_pollable_input_stream_read_nonblocking(c_stream, c_buffer, c_count, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_pollable_output_stream_write_nonblocking(stream unsafe.Pointer, buffer []uint8, count uint64, cancellable unsafe.Pointer, error unsafe.Pointer) uint64 {
	c_stream := (*C.GPollableOutputStream)(unsafe.Pointer(stream))

	c_buffer := (unsafe.Pointer)(unsafe.Pointer(&buffer[0]))

	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_pollable_output_stream_write_nonblocking(c_stream, c_buffer, c_count, c_cancellable, cError)

	return (uint64)(ret)
}

func Fn_g_proxy_connect(proxy unsafe.Pointer, connection unsafe.Pointer, proxyAddress unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GProxy)(unsafe.Pointer(proxy))

	c_connection := (*C.GIOStream)(unsafe.Pointer(connection))

	c_proxyAddress := (*C.GProxyAddress)(unsafe.Pointer(proxyAddress))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect(c_proxy, c_connection, c_proxyAddress, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

func Fn_g_proxy_connect_finish(proxy unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GProxy)(unsafe.Pointer(proxy))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_proxy_connect_finish(c_proxy, c_result, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_proxy_supports_hostname(proxy unsafe.Pointer) bool {
	c_proxy := (*C.GProxy)(unsafe.Pointer(proxy))

	ret := C.g_proxy_supports_hostname(c_proxy)

	return toGoBool(ret)
}

func Fn_g_proxy_resolver_is_supported(resolver unsafe.Pointer) bool {
	c_resolver := (*C.GProxyResolver)(unsafe.Pointer(resolver))

	ret := C.g_proxy_resolver_is_supported(c_resolver)

	return toGoBool(ret)
}

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

func Fn_g_seekable_can_seek(seekable unsafe.Pointer) bool {
	c_seekable := (*C.GSeekable)(unsafe.Pointer(seekable))

	ret := C.g_seekable_can_seek(c_seekable)

	return toGoBool(ret)
}

func Fn_g_seekable_can_truncate(seekable unsafe.Pointer) bool {
	c_seekable := (*C.GSeekable)(unsafe.Pointer(seekable))

	ret := C.g_seekable_can_truncate(c_seekable)

	return toGoBool(ret)
}

func Fn_g_seekable_seek(seekable unsafe.Pointer, offset int64, type_ int, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_seekable := (*C.GSeekable)(unsafe.Pointer(seekable))

	c_offset := (C.goffset)(offset)

	c_type_ := (C.GSeekType)(type_)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_seekable_seek(c_seekable, c_offset, c_type_, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_seekable_tell(seekable unsafe.Pointer) int64 {
	c_seekable := (*C.GSeekable)(unsafe.Pointer(seekable))

	ret := C.g_seekable_tell(c_seekable)

	return (int64)(ret)
}

func Fn_g_seekable_truncate(seekable unsafe.Pointer, offset int64, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_seekable := (*C.GSeekable)(unsafe.Pointer(seekable))

	c_offset := (C.goffset)(offset)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.g_seekable_truncate(c_seekable, c_offset, c_cancellable, cError)

	return toGoBool(ret)
}

func Fn_g_socket_connectable_enumerate(connectable unsafe.Pointer) unsafe.Pointer {
	c_connectable := (*C.GSocketConnectable)(unsafe.Pointer(connectable))

	ret := C.g_socket_connectable_enumerate(c_connectable)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_connectable_proxy_enumerate(connectable unsafe.Pointer) unsafe.Pointer {
	c_connectable := (*C.GSocketConnectable)(unsafe.Pointer(connectable))

	ret := C.g_socket_connectable_proxy_enumerate(c_connectable)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_can_eject(volume unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_can_eject(c_volume)

	return toGoBool(ret)
}

func Fn_g_volume_can_mount(volume unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_can_mount(c_volume)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

func Fn_g_volume_eject_finish(volume unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_finish(c_volume, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

func Fn_g_volume_eject_with_operation_finish(volume unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_with_operation_finish(c_volume, c_result, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

func Fn_g_volume_get_activation_root(volume unsafe.Pointer) unsafe.Pointer {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_activation_root(c_volume)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_drive(volume unsafe.Pointer) unsafe.Pointer {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_drive(c_volume)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_icon(volume unsafe.Pointer) unsafe.Pointer {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_icon(c_volume)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_identifier(volume unsafe.Pointer, kind string) string {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	c_kind := (*C.char)(C.CString(kind))
	defer C.free(unsafe.Pointer(c_kind))

	ret := C.g_volume_get_identifier(c_volume, c_kind)

	return C.GoString(ret)
}

func Fn_g_volume_get_mount(volume unsafe.Pointer) unsafe.Pointer {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_mount(c_volume)

	return unsafe.Pointer(ret)
}

func Fn_g_volume_get_name(volume unsafe.Pointer) string {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_name(c_volume)

	return C.GoString(ret)
}

func Fn_g_volume_get_uuid(volume unsafe.Pointer) string {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_get_uuid(c_volume)

	return C.GoString(ret)
}

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback

func Fn_g_volume_mount_finish(volume unsafe.Pointer, result unsafe.Pointer, error unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	c_result := (*C.GAsyncResult)(unsafe.Pointer(result))

	cError := (**C.GError)(error)

	ret := C.g_volume_mount_finish(c_volume, c_result, cError)

	return toGoBool(ret)
}

func Fn_g_volume_should_automount(volume unsafe.Pointer) bool {
	c_volume := (*C.GVolume)(unsafe.Pointer(volume))

	ret := C.g_volume_should_automount(c_volume)

	return toGoBool(ret)
}
