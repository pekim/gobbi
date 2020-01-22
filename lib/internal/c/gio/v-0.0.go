// Code generated - DO NOT EDIT.
// +build !gio_2.18,!gio_2.20,!gio_2.22,!gio_2.24,!gio_2.26,!gio_2.28,!gio_2.30,!gio_2.32,!gio_2.34,!gio_2.36,!gio_2.38,!gio_2.40,!gio_2.42,!gio_2.44,!gio_2.46,!gio_2.48,!gio_2.50,!gio_2.52,!gio_2.54,!gio_2.56,!gio_2.58,!gio_2.60,!gio_2.62

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
type InputStreamClass C.GInputStreamClass
type InputStreamPrivate C.GInputStreamPrivate
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

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

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

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop : parameter 'func' is callback

// UNSUPPORTED : g_io_scheduler_job_send_to_mainloop_async : parameter 'func' is callback

// UNSUPPORTED : g_resource_enumerate_children : no array length

func Fn_g_settings_schema_get_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GSettingsSchema)(unsafe.Pointer(paramInstance))

	ret := C.g_settings_schema_get_id(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_settings_schema_list_children : no array length

// UNSUPPORTED : g_settings_schema_list_keys : no array length

// UNSUPPORTED : g_settings_schema_source_list_schemas : blacklisted

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

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

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

func Fn_g_content_type_get_description(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_description(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_get_icon(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_icon(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

func Fn_g_content_type_get_mime_type(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_content_type_get_mime_type(cValue0)

	return C.GoString(ret)
}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.guchar)(unsafe.Pointer(&param1[0]))

	cValue2 := (C.gsize)(param2)

	cValue3 := (*C.gboolean)(unsafe.Pointer(param3))

	ret := C.g_content_type_guess(cValue0, cValue1, cValue2, cValue3)

	return C.GoString(ret)
}

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

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

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

func Fn_g_content_types_get_registered() unsafe.Pointer {
	ret := C.g_content_types_get_registered()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

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

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

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

// UNSUPPORTED : g_app_launch_context_get_environment : no array length

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

func Fn_g_application_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GApplicationFlags)(param1)

	ret := C.g_application_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_application_add_main_option_entries : parameter 'entries' is array parameter without length parameter

func Fn_g_application_hold(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_hold(cValueInstance)
}

func Fn_g_application_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_release(cValueInstance)
}

func Fn_g_application_id_is_valid(param0 string) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_application_id_is_valid(cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : g_application_command_line_get_environ : no array length

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

// UNSUPPORTED : g_buffered_input_stream_fill_async : parameter 'callback' is callback

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

func Fn_g_buffered_input_stream_peek(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 uint64) uint64 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (unsafe.Pointer)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (C.gsize)(param2)

	ret := C.g_buffered_input_stream_peek(cValueInstance, cValue0, cValue1, cValue2)

	return (uint64)(ret)
}

func Fn_g_buffered_input_stream_peek_buffer(paramInstance unsafe.Pointer, param0 *uint64) []uint8 {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gsize)(unsafe.Pointer(param0))

	ret := C.g_buffered_input_stream_peek_buffer(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

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

func Fn_g_cancellable_new() unsafe.Pointer {
	ret := C.g_cancellable_new()

	return unsafe.Pointer(ret)
}

func Fn_g_cancellable_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_cancel(cValueInstance)
}

// UNSUPPORTED : g_cancellable_connect : parameter 'callback' is callback

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

func Fn_g_cancellable_pop_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_pop_current(cValueInstance)
}

func Fn_g_cancellable_push_current(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_push_current(cValueInstance)
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

func Fn_g_cancellable_get_current() unsafe.Pointer {
	ret := C.g_cancellable_get_current()

	return unsafe.Pointer(ret)
}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	ret := C.g_converter_input_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))

	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	ret := C.g_converter_output_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_dbus_connection_add_filter : parameter 'filter_function' is callback

// UNSUPPORTED : g_dbus_connection_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_close : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_flush : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_register_object : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_register_subtree : parameter 'user_data_free_func' is callback

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_signal_subscribe : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_connection_new_for_address : parameter 'callback' is callback

func Fn_g_dbus_message_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_byte_order(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : no array length

func Fn_g_dbus_message_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageByteOrder)(param0)

	C.g_dbus_message_set_byte_order(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new : parameter 'get_proxy_type_func' is callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : parameter 'get_proxy_type_func' is callback

func Fn_g_dbus_object_manager_server_set_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_set_connection(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_proxy_call : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : no array length

// UNSUPPORTED : g_dbus_proxy_new : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : parameter 'callback' is callback

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

// UNSUPPORTED : g_data_input_stream_read_line : no array length

// UNSUPPORTED : g_data_input_stream_read_line_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_line_finish : no array length

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

// UNSUPPORTED : g_data_input_stream_read_until_async : parameter 'callback' is callback

// UNSUPPORTED : g_data_input_stream_read_upto_async : parameter 'callback' is callback

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

func Fn_g_desktop_app_info_get_categories(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_desktop_app_info_get_categories(cValueInstance)

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

// UNSUPPORTED : g_desktop_app_info_get_keywords : no array length

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : parameter 'user_setup' is callback

// UNSUPPORTED : g_desktop_app_info_list_actions : no array length

// UNSUPPORTED : g_desktop_app_info_search : array has no type

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_desktop_app_info_set_desktop_env(cValue0)
}

func Fn_g_file_enumerator_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_enumerator_close_async : parameter 'callback' is callback

func Fn_g_file_enumerator_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_close_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
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

func Fn_g_file_enumerator_next_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_enumerator_next_file(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_enumerator_next_files_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_io_stream_query_info_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_info_get_attribute_stringv : no array length

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

// UNSUPPORTED : g_file_info_list_attributes : no array length

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

func Fn_g_file_info_set_attribute_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.g_file_info_set_attribute_string(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_file_info_set_attribute_stringv : parameter 'attr_value' is array parameter without length parameter

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

// UNSUPPORTED : g_file_input_stream_query_info_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_output_stream_query_info_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_filename_completer_get_completions : no array length

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

// UNSUPPORTED : g_io_stream_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_io_stream_splice_async : parameter 'callback' is callback

// UNSUPPORTED : g_inet_address_new_from_bytes : parameter 'bytes' is array parameter without length parameter

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

// UNSUPPORTED : g_input_stream_close_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_input_stream_read : blacklisted

// UNSUPPORTED : g_input_stream_read_all : blacklisted

// UNSUPPORTED : g_input_stream_read_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_async : parameter 'callback' is callback

// UNSUPPORTED : g_input_stream_read_bytes_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_input_stream_skip_async : parameter 'callback' is callback

func Fn_g_input_stream_skip_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_input_stream_skip_finish(cValueInstance, cValue0, cError)

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

func Fn_g_memory_output_stream_get_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_get_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_g_memory_output_stream_get_size(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GMemoryOutputStream)(unsafe.Pointer(paramInstance))

	ret := C.g_memory_output_stream_get_size(cValueInstance)

	return (uint64)(ret)
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

func Fn_g_notification_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GNotification)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GNotificationPriority)(param0)

	C.g_notification_set_priority(cValueInstance, cValue0)
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

// UNSUPPORTED : g_output_stream_close_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_output_stream_flush_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_output_stream_splice_async : parameter 'callback' is callback

func Fn_g_output_stream_splice_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_splice_finish(cValueInstance, cValue0, cError)

	return (uint64)(ret)
}

func Fn_g_output_stream_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (unsafe.Pointer)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

func Fn_g_output_stream_write_all(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (unsafe.Pointer)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	cValue3 := (*C.GCancellable)(unsafe.Pointer(param3))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_all(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_output_stream_write_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_write_async : parameter 'callback' is callback

func Fn_g_output_stream_write_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_output_stream_write_bytes(cValueInstance, cValue0, cValue1, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_output_stream_write_bytes_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_output_stream_writev_all_async : parameter 'callback' is callback

// UNSUPPORTED : g_output_stream_writev_async : parameter 'callback' is callback

// UNSUPPORTED : g_permission_acquire_async : parameter 'callback' is callback

// UNSUPPORTED : g_permission_release_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_address_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_by_name_with_flags_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_records_async : parameter 'callback' is callback

// UNSUPPORTED : g_resolver_lookup_service_async : parameter 'callback' is callback

func Fn_g_settings_apply(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_apply(cValueInstance)
}

// UNSUPPORTED : g_settings_bind_with_mapping : parameter 'get_mapping' is callback

// UNSUPPORTED : g_settings_get_mapped : parameter 'mapping' is callback

// UNSUPPORTED : g_settings_get_strv : no array length

// UNSUPPORTED : g_settings_list_children : no array length

// UNSUPPORTED : g_settings_list_keys : no array length

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

// UNSUPPORTED : g_settings_set_strv : parameter 'value' is array parameter without length parameter

// UNSUPPORTED : g_settings_list_relocatable_schemas : no array length

// UNSUPPORTED : g_settings_list_schemas : no array length

func Fn_g_settings_sync() {
	C.g_settings_sync()
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

// UNSUPPORTED : g_simple_async_result_run_in_thread : parameter 'func' is callback

func Fn_g_simple_async_result_set_error(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.c_g_simple_async_result_set_error(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_simple_async_result_set_error_va(paramInstance unsafe.Pointer, param0 uint32, param1 int, param2 string) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.c_g_simple_async_result_set_error_va(cValueInstance, cValue0, cValue1, cValue2)
}

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

// UNSUPPORTED : g_simple_async_result_set_op_res_gpointer : parameter 'destroy_op_res' is callback

func Fn_g_simple_async_result_set_op_res_gssize(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gssize)(param0)

	C.g_simple_async_result_set_op_res_gssize(cValueInstance, cValue0)
}

func Fn_g_socket_address_enumerator_next(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_socket_address_enumerator_next_async : parameter 'callback' is callback

func Fn_g_socket_address_enumerator_next_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_socket_address_enumerator_next_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_socket_client_add_application_proxy(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_socket_client_add_application_proxy(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_client_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_host_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_socket_client_connect_to_service_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_client_connect_to_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_connection_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_control_message_deserialize : parameter 'data' is array parameter with indirection of 0

// UNSUPPORTED : g_socket_listener_accept_async : parameter 'callback' is callback

// UNSUPPORTED : g_socket_listener_accept_socket_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_newv : parameter 'argv' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_communicate_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_subprocess_communicate_utf8_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_subprocess_wait_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_wait_check_async : parameter 'callback' is callback

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : parameter 'child_setup' is callback

// UNSUPPORTED : g_subprocess_launcher_set_environ : parameter 'env' is array parameter without length parameter

// UNSUPPORTED : g_subprocess_launcher_spawnv : parameter 'argv' is array parameter without length parameter

func Fn_g_subprocess_launcher_take_fd(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.g_subprocess_launcher_take_fd(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_task_new : parameter 'callback' is callback

// UNSUPPORTED : g_task_attach_source : parameter 'callback' is callback

// UNSUPPORTED : g_task_return_pointer : parameter 'result_destroy' is callback

// UNSUPPORTED : g_task_run_in_thread : parameter 'task_func' is callback

// UNSUPPORTED : g_task_run_in_thread_sync : parameter 'task_func' is callback

// UNSUPPORTED : g_task_set_task_data : parameter 'task_data_destroy' is callback

// UNSUPPORTED : g_task_report_error : parameter 'callback' is callback

// UNSUPPORTED : g_task_report_new_error : parameter 'callback' is callback

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

// UNSUPPORTED : g_themed_icon_get_names : no array length

func Fn_g_tls_connection_get_use_system_certdb(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_use_system_certdb(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_tls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

func Fn_g_tls_connection_set_use_system_certdb(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_tls_connection_set_use_system_certdb(cValueInstance, cValue0)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : parameter 'issuer_raw_dn' is array parameter without length parameter

// UNSUPPORTED : g_tls_database_verify_chain_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_ask_password_async : parameter 'callback' is callback

// UNSUPPORTED : g_tls_interaction_request_certificate_async : parameter 'callback' is callback

func Fn_g_tls_password_new(param0 int, param1 string) unsafe.Pointer {
	cValue0 := (C.GTlsPasswordFlags)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_tls_password_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_password_set_value_full : parameter 'destroy' is callback

// UNSUPPORTED : g_unix_connection_receive_credentials_async : parameter 'callback' is callback

// UNSUPPORTED : g_unix_connection_send_credentials_async : parameter 'callback' is callback

func Fn_g_unix_input_stream_new(param0 int, param1 bool) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	ret := C.g_unix_input_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_mount_monitor_new() unsafe.Pointer {
	ret := C.g_unix_mount_monitor_new()

	return unsafe.Pointer(ret)
}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) unsafe.Pointer {
	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	ret := C.g_unix_output_stream_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gint)(param1)

	ret := C.g_unix_socket_address_new_abstract(cValue0, cValue1)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : g_vfs_get_supported_uri_schemes : no array length

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

// UNSUPPORTED : g_vfs_register_uri_scheme : parameter 'uri_func' is callback

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

// UNSUPPORTED : g_action_group_list_actions : no array length

func Fn_g_app_info_add_supports_type(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_app_info_add_supports_type(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_app_info_can_remove_supports_type(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_can_remove_supports_type(cValueInstance)

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

func Fn_g_app_info_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GAppInfo)(unsafe.Pointer(paramInstance))

	ret := C.g_app_info_get_description(cValueInstance)

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

// UNSUPPORTED : g_app_info_get_supported_types : no array length

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

// UNSUPPORTED : g_app_info_launch_uris_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_async_initable_init_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_drive_eject : parameter 'callback' is callback

func Fn_g_drive_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_drive_enumerate_identifiers : no array length

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

// UNSUPPORTED : g_drive_poll_for_media : parameter 'callback' is callback

func Fn_g_drive_poll_for_media_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_drive_poll_for_media_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_drive_start : parameter 'callback' is callback

// UNSUPPORTED : g_drive_stop : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_close_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_handshake_async : parameter 'callback' is callback

// UNSUPPORTED : g_dtls_connection_set_advertised_protocols : parameter 'protocols' is array parameter without length parameter

// UNSUPPORTED : g_dtls_connection_shutdown_async : parameter 'callback' is callback

func Fn_g_file_append_to(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GFileCreateFlags)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_append_to_async : parameter 'callback' is callback

func Fn_g_file_append_to_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_append_to_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_copy : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_copy_async : parameter 'progress_callback' is callback

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

// UNSUPPORTED : g_file_create_async : parameter 'callback' is callback

func Fn_g_file_create_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_create_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_create_readwrite_async : parameter 'callback' is callback

func Fn_g_file_delete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_delete(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_delete_async : parameter 'callback' is callback

func Fn_g_file_dup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	ret := C.g_file_dup(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_eject_mountable : parameter 'callback' is callback

func Fn_g_file_eject_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_eject_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_eject_mountable_with_operation : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_enumerate_children_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_find_enclosing_mount_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_load_bytes_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents : blacklisted

// UNSUPPORTED : g_file_load_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_load_contents_finish : blacklisted

// UNSUPPORTED : g_file_load_partial_contents_async : parameter 'read_more_callback' is callback

// UNSUPPORTED : g_file_load_partial_contents_finish : blacklisted

func Fn_g_file_make_directory(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_make_directory(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_make_directory_async : parameter 'callback' is callback

func Fn_g_file_make_symbolic_link(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_file_make_symbolic_link(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_measure_disk_usage : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_measure_disk_usage_async : parameter 'progress_callback' is callback

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

// UNSUPPORTED : g_file_mount_enclosing_volume : parameter 'callback' is callback

func Fn_g_file_mount_enclosing_volume_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_enclosing_volume_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_mount_mountable : parameter 'callback' is callback

func Fn_g_file_mount_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_mount_mountable_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_move : parameter 'progress_callback' is callback

// UNSUPPORTED : g_file_open_readwrite_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_poll_mountable : parameter 'callback' is callback

func Fn_g_file_query_default_handler(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_query_default_handler(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_query_default_handler_async : parameter 'callback' is callback

func Fn_g_file_query_exists(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	ret := C.g_file_query_exists(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : g_file_query_filesystem_info_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_query_info_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_read_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_replace_async : parameter 'callback' is callback

func Fn_g_file_replace_contents(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 string, param3 bool, param4 int, param5 *string, param6 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := toCBool(param3)

	cValue4 := (C.GFileCreateFlags)(param4)

	var cValue5String *C.gchar
	cValue5 := &cValue5String

	cValue6 := (*C.GCancellable)(unsafe.Pointer(param6))

	cError := (**C.GError)(error)

	ret := C.g_file_replace_contents(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cError)

	param5String := C.GoString(cValue5String)
	*param5 = param5String

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_replace_contents_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_replace_contents_bytes_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_replace_readwrite_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_set_attributes_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_file_set_display_name_async : parameter 'callback' is callback

func Fn_g_file_set_display_name_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_set_display_name_finish(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_file_start_mountable : parameter 'callback' is callback

// UNSUPPORTED : g_file_stop_mountable : parameter 'callback' is callback

func Fn_g_file_trash(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_trash(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_trash_async : parameter 'callback' is callback

// UNSUPPORTED : g_file_unmount_mountable : parameter 'callback' is callback

func Fn_g_file_unmount_mountable_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GFile)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_file_unmount_mountable_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_file_unmount_mountable_with_operation : parameter 'callback' is callback

func Fn_g_icon_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.g_icon_equal(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : g_loadable_icon_load_async : parameter 'callback' is callback

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

// UNSUPPORTED : g_mount_eject : parameter 'callback' is callback

func Fn_g_mount_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_eject_with_operation : parameter 'callback' is callback

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

// UNSUPPORTED : g_mount_guess_content_type : parameter 'callback' is callback

// UNSUPPORTED : g_mount_guess_content_type_finish : no array length

// UNSUPPORTED : g_mount_guess_content_type_sync : no array length

// UNSUPPORTED : g_mount_remount : parameter 'callback' is callback

func Fn_g_mount_remount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_remount_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_unmount : parameter 'callback' is callback

func Fn_g_mount_unmount_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GMount)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_mount_unmount_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_mount_unmount_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_network_monitor_can_reach_async : parameter 'callback' is callback

func Fn_g_network_monitor_can_reach_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_g_pollable_input_stream_read_nonblocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GPollableInputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (unsafe.Pointer)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_pollable_input_stream_read_nonblocking(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

func Fn_g_pollable_output_stream_write_nonblocking(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer, error unsafe.Pointer) uint64 {
	cValueInstance := (*C.GPollableOutputStream)(unsafe.Pointer(paramInstance))

	cValue0 := (unsafe.Pointer)(unsafe.Pointer(&param0[0]))

	cValue1 := (C.gsize)(param1)

	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.g_pollable_output_stream_write_nonblocking(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint64)(ret)
}

// UNSUPPORTED : g_proxy_connect_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup : no array length

// UNSUPPORTED : g_proxy_resolver_lookup_async : parameter 'callback' is callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : no array length

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

// UNSUPPORTED : g_volume_eject : parameter 'callback' is callback

func Fn_g_volume_eject_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_volume_eject_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_volume_eject_with_operation : parameter 'callback' is callback

// UNSUPPORTED : g_volume_enumerate_identifiers : no array length

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

func Fn_g_volume_get_uuid(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GVolume)(unsafe.Pointer(paramInstance))

	ret := C.g_volume_get_uuid(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_volume_mount : parameter 'callback' is callback

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
