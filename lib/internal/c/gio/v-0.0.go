// Code generated - DO NOT EDIT.
// +build !gio_2.18,!gio_2.20,!gio_2.22,!gio_2.24,!gio_2.26,!gio_2.28,!gio_2.30,!gio_2.32,!gio_2.34,!gio_2.36,!gio_2.38,!gio_2.40,!gio_2.42,!gio_2.44,!gio_2.46,!gio_2.48,!gio_2.52,!gio_2.58,!gio_2.60

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
import "C"

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

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int) {}

func Fn_g_app_info_get_all() {
	C.g_app_info_get_all()
}

func Fn_g_app_info_get_all_for_type(param0 string) {}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) {}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) {}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_g_content_type_can_be_executable(param0 string) {}

func Fn_g_content_type_equals(param0 string, param1 string) {}

func Fn_g_content_type_get_description(param0 string) {}

func Fn_g_content_type_get_icon(param0 string) {}

func Fn_g_content_type_get_mime_type(param0 string) {}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) {}

func Fn_g_content_type_is_a(param0 string, param1 string) {}

func Fn_g_content_type_is_unknown(param0 string) {}

func Fn_g_content_types_get_registered() {
	C.g_content_types_get_registered()
}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_g_dbus_error_quark() {
	C.g_dbus_error_quark()
}

func Fn_g_file_new_for_commandline_arg(param0 string) {}

func Fn_g_file_new_for_path(param0 string) {}

func Fn_g_file_new_for_uri(param0 string) {}

func Fn_g_file_parse_name(param0 string) {}

func Fn_g_icon_hash(param0 unsafe.Pointer) {}

func Fn_g_io_error_from_errno(param0 int) {}

func Fn_g_io_error_quark() {
	C.g_io_error_quark()
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) {}

func Fn_g_io_extension_point_lookup(param0 string) {}

func Fn_g_io_extension_point_register(param0 string) {}

func Fn_g_io_modules_load_all_in_directory(param0 string) {}

func Fn_g_io_scheduler_cancel_all_jobs() {
	C.g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : io_scheduler_push_job : has callback

// UNSUPPORTED : keyfile_settings_backend_new : blacklisted
// UNSUPPORTED : memory_settings_backend_new : blacklisted
// UNSUPPORTED : null_settings_backend_new : blacklisted
// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_g_unix_is_mount_path_system_internal(param0 string) {}

func Fn_g_unix_mount_at(param0 string, param1 *uint64) {}

func Fn_g_unix_mount_compare(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_unix_mount_free(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_device_path(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_fs_type(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_mount_path(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_can_eject(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_icon(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_name(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_should_display(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_is_readonly(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_is_system_internal(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_points_changed_since(param0 uint64) {}

func Fn_g_unix_mount_points_get(param0 *uint64) {}

func Fn_g_unix_mounts_changed_since(param0 uint64) {}

func Fn_g_unix_mounts_get(param0 *uint64) {}

func Fn_g_app_launch_context_new() {
	C.g_app_launch_context_new()
}

func Fn_g_application_new(param0 string, param1 int) {}

func Fn_g_application_id_is_valid(param0 string) {}

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_buffered_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_output_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_cancellable_new() {
	C.g_cancellable_new()
}

func Fn_g_cancellable_get_current() {
	C.g_cancellable_get_current()
}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_address : has callback

// UNSUPPORTED : new_for_bus_sync : has callback

// UNSUPPORTED : new_sync : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_data_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_data_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_desktop_app_info_new(param0 string) {}

func Fn_g_desktop_app_info_new_from_filename(param0 string) {}

func Fn_g_desktop_app_info_search(param0 string) {}

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {}

func Fn_g_file_icon_new(param0 unsafe.Pointer) {}

func Fn_g_file_info_new() {
	C.g_file_info_new()
}

func Fn_g_filename_completer_new() {
	C.g_filename_completer_new()
}

func Fn_g_io_module_new(param0 string) {}

// UNSUPPORTED : query : blacklisted
func Fn_g_memory_input_stream_new() {
	C.g_memory_input_stream_new()
}

// UNSUPPORTED : new_from_data : has callback

// UNSUPPORTED : new : has callback

func Fn_g_mount_operation_new() {
	C.g_mount_operation_new()
}

func Fn_g_settings_sync() {
	C.g_settings_sync()
}

// UNSUPPORTED : flatten_tree : blacklisted
// UNSUPPORTED : get_default : blacklisted
// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_error : has varargs

// UNSUPPORTED : new_from_error : has callback

// UNSUPPORTED : new_take_error : has callback

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new : has callback

// UNSUPPORTED : report_error : has callback

// UNSUPPORTED : report_new_error : has varargs

func Fn_g_test_dbus_new(param0 int) {}

func Fn_g_test_dbus_unset() {
	C.g_test_dbus_unset()
}

func Fn_g_themed_icon_new(param0 string) {}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) {}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) {}

func Fn_g_tls_password_new(param0 int, param1 string) {}

func Fn_g_unix_input_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_mount_monitor_new() {
	C.g_unix_mount_monitor_new()
}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) {}

func Fn_g_vfs_get_default() {
	C.g_vfs_get_default()
}

func Fn_g_vfs_get_local() {
	C.g_vfs_get_local()
}

func Fn_g_volume_monitor_adopt_orphan_mount(param0 unsafe.Pointer) {}

func Fn_g_volume_monitor_get() {
	C.g_volume_monitor_get()
}
