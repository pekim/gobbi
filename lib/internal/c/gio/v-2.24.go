// Code generated - DO NOT EDIT.
// +build gio_2.24

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

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GAppInfoCreateFlags)(param2)

	cError := (**C.GError)(error)

	ret := C.g_app_info_create_from_commandline(cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_all() unsafe.Pointer {
	ret := C.g_app_info_get_all()

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_all_for_type(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_app_info_get_all_for_type(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.g_app_info_get_default_for_type(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_app_info_get_default_for_uri_scheme(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GAppLaunchContext)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_app_info_launch_default_for_uri(cValue0, cValue1, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : has callback

func Fn_g_app_info_reset_type_associations(param0 string) {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.g_app_info_reset_type_associations(cValue0)
}

// UNSUPPORTED : g_async_initable_newv_async : has callback

// UNSUPPORTED : g_bus_get : has callback

// UNSUPPORTED : g_bus_own_name : has callback

// UNSUPPORTED : g_bus_own_name_on_connection : has callback

// UNSUPPORTED : g_bus_watch_name : has callback

// UNSUPPORTED : g_bus_watch_name_on_connection : has callback

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

// UNSUPPORTED : g_dbus_address_get_stream : has callback

// UNSUPPORTED : g_dbus_annotation_info_lookup : has non-string array param annotations

func Fn_g_dbus_error_quark() uint32 {
	ret := C.g_dbus_error_quark()

	return (uint32)(ret)
}

// UNSUPPORTED : g_dbus_error_register_error_domain : has non-string array param entries

func Fn_g_file_new_for_commandline_arg(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_new_for_commandline_arg(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_new_for_path(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_new_for_path(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_new_for_uri(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_new_for_uri(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_file_parse_name(param0 string) unsafe.Pointer {
	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.g_file_parse_name(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_g_icon_hash(param0 unsafe.Pointer) uint {
	cValue0 := (C.gconstpointer)(param0)

	ret := C.g_icon_hash(cValue0)

	return (uint)(ret)
}

func Fn_g_icon_new_for_string(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.g_icon_new_for_string(cValue0, cError)

	return unsafe.Pointer(ret)
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
// UNSUPPORTED : g_null_settings_backend_new : blacklisted
// UNSUPPORTED : g_pollable_stream_read : has non-string array param buffer

// UNSUPPORTED : g_pollable_stream_write : has non-string array param buffer

// UNSUPPORTED : g_pollable_stream_write_all : has non-string array param buffer

func Fn_g_resolver_error_quark() uint32 {
	ret := C.g_resolver_error_quark()

	return (uint32)(ret)
}

// UNSUPPORTED : g_resources_enumerate_children : has array return

// UNSUPPORTED : g_simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : has callback

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	ret := C.g_srv_target_list_sort(cValue0)

	return unsafe.Pointer(ret)
}

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

func Fn_g_application_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GApplicationFlags)(param1)

	ret := C.g_application_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_application_add_main_option_entries : has non-string array param entries

func Fn_g_application_hold(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_hold(cValueInstance)
}

// UNSUPPORTED : g_application_open : has non-string array param files

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

// UNSUPPORTED : g_application_command_line_get_arguments : has array return

// UNSUPPORTED : g_application_command_line_get_environ : has array return

// UNSUPPORTED : g_application_command_line_print : has varargs

// UNSUPPORTED : g_application_command_line_printerr : has varargs

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

// UNSUPPORTED : g_dbus_connection_add_filter : has callback

// UNSUPPORTED : g_dbus_connection_call : has callback

// UNSUPPORTED : g_dbus_connection_call_with_unix_fd_list : has callback

// UNSUPPORTED : g_dbus_connection_close : has callback

// UNSUPPORTED : g_dbus_connection_flush : has callback

// UNSUPPORTED : g_dbus_connection_register_object : has callback

// UNSUPPORTED : g_dbus_connection_register_subtree : has callback

// UNSUPPORTED : g_dbus_connection_send_message_with_reply : has callback

// UNSUPPORTED : g_dbus_connection_signal_subscribe : has callback

// UNSUPPORTED : g_dbus_connection_new : has callback

// UNSUPPORTED : g_dbus_connection_new_for_address : has callback

// UNSUPPORTED : g_dbus_message_new_from_blob : has non-string array param blob

func Fn_g_dbus_message_get_byte_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	ret := C.g_dbus_message_get_byte_order(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : g_dbus_message_get_header_fields : has array return

// UNSUPPORTED : g_dbus_message_new_method_error : has varargs

// UNSUPPORTED : g_dbus_message_new_method_error_valist : has va_list

func Fn_g_dbus_message_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GDBusMessageByteOrder)(param0)

	C.g_dbus_message_set_byte_order(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_message_to_blob : has array return

// UNSUPPORTED : g_dbus_message_bytes_needed : has non-string array param blob

// UNSUPPORTED : g_dbus_method_invocation_return_error : has varargs

// UNSUPPORTED : g_dbus_method_invocation_return_error_valist : has va_list

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus_sync : has callback

// UNSUPPORTED : g_dbus_object_manager_client_new_sync : has callback

// UNSUPPORTED : g_dbus_object_manager_client_new : has callback

// UNSUPPORTED : g_dbus_object_manager_client_new_for_bus : has callback

func Fn_g_dbus_object_manager_server_set_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_set_connection(cValueInstance, cValue0)
}

// UNSUPPORTED : g_dbus_proxy_call : has callback

// UNSUPPORTED : g_dbus_proxy_call_with_unix_fd_list : has callback

// UNSUPPORTED : g_dbus_proxy_get_cached_property_names : has array return

// UNSUPPORTED : g_dbus_proxy_new : has callback

// UNSUPPORTED : g_dbus_proxy_new_for_bus : has callback

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

// UNSUPPORTED : g_desktop_app_info_get_string_list : has array return

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager : has callback

// UNSUPPORTED : g_desktop_app_info_launch_uris_as_manager_with_fds : has callback

// UNSUPPORTED : g_desktop_app_info_list_actions : has array return

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

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.GInetAddress)(unsafe.Pointer(param0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_inet_socket_address_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_address(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_address(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_g_inet_socket_address_get_port(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GInetSocketAddress)(unsafe.Pointer(paramInstance))

	ret := C.g_inet_socket_address_get_port(cValueInstance)

	return (uint16)(ret)
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

// UNSUPPORTED : g_input_stream_read_async : has callback

// UNSUPPORTED : g_input_stream_read_bytes_async : has callback

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

// UNSUPPORTED : g_list_store_insert_sorted : has callback

// UNSUPPORTED : g_list_store_sort : has callback

// UNSUPPORTED : g_list_store_splice : has non-string array param additions

func Fn_g_memory_input_stream_new() unsafe.Pointer {
	ret := C.g_memory_input_stream_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_memory_input_stream_new_from_data : has callback

// UNSUPPORTED : g_memory_input_stream_add_data : has callback

// UNSUPPORTED : g_memory_output_stream_new : has callback

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

// UNSUPPORTED : g_menu_item_get_attribute : has varargs

// UNSUPPORTED : g_menu_item_set_action_and_target : has varargs

// UNSUPPORTED : g_menu_item_set_attribute : has varargs

// UNSUPPORTED : g_menu_model_get_item_attribute : has varargs

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

func Fn_g_network_address_new(param0 string, param1 uint16) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	ret := C.g_network_address_new(cValue0, cValue1)

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

func Fn_g_network_address_parse(param0 string, param1 uint16, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint16)(param1)

	cError := (**C.GError)(error)

	ret := C.g_network_address_parse(cValue0, cValue1, cError)

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

func Fn_g_network_service_get_service(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GNetworkService)(unsafe.Pointer(paramInstance))

	ret := C.g_network_service_get_service(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : g_notification_add_button_with_target : has varargs

// UNSUPPORTED : g_notification_set_default_action_and_target : has varargs

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

// UNSUPPORTED : g_permission_acquire_async : has callback

// UNSUPPORTED : g_permission_release_async : has callback

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

// UNSUPPORTED : g_resolver_lookup_records_async : has callback

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

func Fn_g_settings_apply(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_apply(cValueInstance)
}

// UNSUPPORTED : g_settings_bind_with_mapping : has callback

// UNSUPPORTED : g_settings_get : has varargs

// UNSUPPORTED : g_settings_get_mapped : has callback

// UNSUPPORTED : g_settings_get_strv : has array return

// UNSUPPORTED : g_settings_list_children : has array return

// UNSUPPORTED : g_settings_list_keys : has array return

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

// UNSUPPORTED : g_settings_list_relocatable_schemas : has array return

// UNSUPPORTED : g_settings_list_schemas : has array return

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
// UNSUPPORTED : g_simple_action_group_add_entries : has non-string array param entries

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

func Fn_g_simple_async_result_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (C.gpointer)(param2)

	ret := C.g_simple_async_result_is_valid(cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_g_socket_get_blocking(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GSocket)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_get_blocking(cValueInstance)

	return toGoBool(ret)
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

// UNSUPPORTED : g_socket_client_connect_to_uri_async : has callback

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

func Fn_g_socket_client_get_socket_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	ret := C.g_socket_client_get_socket_type(cValueInstance)

	return (int)(ret)
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

func Fn_g_socket_client_set_socket_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GSocketClient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GSocketType)(param0)

	C.g_socket_client_set_socket_type(cValueInstance, cValue0)
}

// UNSUPPORTED : g_socket_connection_connect_async : has callback

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

// UNSUPPORTED : g_subprocess_wait_async : has callback

// UNSUPPORTED : g_subprocess_wait_check_async : has callback

// UNSUPPORTED : g_subprocess_launcher_set_child_setup : has callback

// UNSUPPORTED : g_subprocess_launcher_spawn : has varargs

func Fn_g_subprocess_launcher_take_fd(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GSubprocessLauncher)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.g_subprocess_launcher_take_fd(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : g_task_new : has callback

// UNSUPPORTED : g_task_attach_source : has callback

// UNSUPPORTED : g_task_return_new_error : has varargs

// UNSUPPORTED : g_task_return_pointer : has callback

// UNSUPPORTED : g_task_run_in_thread : has callback

// UNSUPPORTED : g_task_run_in_thread_sync : has callback

// UNSUPPORTED : g_task_set_task_data : has callback

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

func Fn_g_tls_connection_get_use_system_certdb(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	ret := C.g_tls_connection_get_use_system_certdb(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : g_tls_connection_handshake_async : has callback

func Fn_g_tls_connection_set_use_system_certdb(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.g_tls_connection_set_use_system_certdb(cValueInstance, cValue0)
}

// UNSUPPORTED : g_tls_database_lookup_certificate_for_handle_async : has callback

// UNSUPPORTED : g_tls_database_lookup_certificate_issuer_async : has callback

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by : has non-string array param issuer_raw_dn

// UNSUPPORTED : g_tls_database_lookup_certificates_issued_by_async : has callback

// UNSUPPORTED : g_tls_database_verify_chain_async : has callback

// UNSUPPORTED : g_tls_interaction_ask_password_async : has callback

// UNSUPPORTED : g_tls_interaction_request_certificate_async : has callback

func Fn_g_tls_password_new(param0 int, param1 string) unsafe.Pointer {
	cValue0 := (C.GTlsPasswordFlags)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.g_tls_password_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_tls_password_set_value : has non-string array param value

// UNSUPPORTED : g_tls_password_set_value_full : has callback

// UNSUPPORTED : g_unix_connection_receive_credentials_async : has callback

func Fn_g_unix_connection_receive_fd(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_receive_fd(cValueInstance, cValue0, cError)

	return (int)(ret)
}

// UNSUPPORTED : g_unix_connection_send_credentials_async : has callback

func Fn_g_unix_connection_send_fd(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GUnixConnection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.g_unix_connection_send_fd(cValueInstance, cValue0, cValue1, cError)

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

func Fn_g_zlib_decompressor_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GZlibCompressorFormat)(param0)

	ret := C.g_zlib_decompressor_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : g_action_group_list_actions : has array return

// UNSUPPORTED : g_action_map_add_action_entries : has non-string array param entries

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

// UNSUPPORTED : g_converter_convert : has non-string array param inbuf

func Fn_g_converter_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GConverter)(unsafe.Pointer(paramInstance))

	C.g_converter_reset(cValueInstance)
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

func Fn_g_drive_get_start_stop_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GDrive)(unsafe.Pointer(paramInstance))

	ret := C.g_drive_get_start_stop_type(cValueInstance)

	return (int)(ret)
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

// UNSUPPORTED : g_network_monitor_can_reach_async : has callback

func Fn_g_network_monitor_can_reach_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GNetworkMonitor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.g_network_monitor_can_reach_finish(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : g_pollable_input_stream_read_nonblocking : has non-string array param buffer

// UNSUPPORTED : g_pollable_output_stream_write_nonblocking : has non-string array param buffer

// UNSUPPORTED : g_pollable_output_stream_writev_nonblocking : has non-string array param vectors

// UNSUPPORTED : g_proxy_connect_async : has callback

// UNSUPPORTED : g_proxy_resolver_lookup : has array return

// UNSUPPORTED : g_proxy_resolver_lookup_async : has callback

// UNSUPPORTED : g_proxy_resolver_lookup_finish : has array return

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
