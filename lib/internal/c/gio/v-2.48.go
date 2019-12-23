// Code generated - DO NOT EDIT.
// +build gio_2.48

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
import "C"

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

func Fn_g_action_name_is_valid(param0 string) {}

func Fn_g_action_parse_detailed_name(param0 string, param1 string, param2 *unsafe.Pointer) {}

func Fn_g_action_print_detailed_name(param0 string, param1 unsafe.Pointer) {}

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int) {}

func Fn_g_app_info_get_all() {
	C.g_app_info_get_all()
}

func Fn_g_app_info_get_all_for_type(param0 string) {}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) {}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) {}

func Fn_g_app_info_get_fallback_for_type(param0 string) {}

func Fn_g_app_info_get_recommended_for_type(param0 string) {}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_g_app_info_reset_type_associations(param0 string) {}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

func Fn_g_bus_get_finish(param0 unsafe.Pointer) {}

func Fn_g_bus_get_sync(param0 int, param1 unsafe.Pointer) {}

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

func Fn_g_bus_own_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_g_bus_own_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_g_bus_unown_name(param0 uint) {}

func Fn_g_bus_unwatch_name(param0 uint) {}

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_g_bus_watch_name_on_connection_with_closures(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_g_bus_watch_name_with_closures(param0 int, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_g_content_type_can_be_executable(param0 string) {}

func Fn_g_content_type_equals(param0 string, param1 string) {}

func Fn_g_content_type_from_mime_type(param0 string) {}

func Fn_g_content_type_get_description(param0 string) {}

func Fn_g_content_type_get_generic_icon_name(param0 string) {}

func Fn_g_content_type_get_icon(param0 string) {}

func Fn_g_content_type_get_mime_type(param0 string) {}

func Fn_g_content_type_get_symbolic_icon(param0 string) {}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) {}

func Fn_g_content_type_guess_for_tree(param0 unsafe.Pointer) {}

func Fn_g_content_type_is_a(param0 string, param1 string) {}

func Fn_g_content_type_is_unknown(param0 string) {}

func Fn_g_content_types_get_registered() {
	C.g_content_types_get_registered()
}

func Fn_g_dbus_address_escape_value(param0 string) {}

func Fn_g_dbus_address_get_for_bus_sync(param0 int, param1 unsafe.Pointer) {}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_g_dbus_address_get_stream_finish(param0 unsafe.Pointer, param1 string) {}

func Fn_g_dbus_address_get_stream_sync(param0 string, param1 string, param2 unsafe.Pointer) {}

func Fn_g_dbus_annotation_info_lookup(param0 []unsafe.Pointer, param1 string) {}

func Fn_g_dbus_error_encode_gerror(param0 unsafe.Pointer) {}

func Fn_g_dbus_error_get_remote_error(param0 unsafe.Pointer) {}

func Fn_g_dbus_error_is_remote_error(param0 unsafe.Pointer) {}

func Fn_g_dbus_error_new_for_dbus_error(param0 string, param1 string) {}

func Fn_g_dbus_error_quark() {
	C.g_dbus_error_quark()
}

func Fn_g_dbus_error_register_error(param0 uint32, param1 int, param2 string) {}

func Fn_g_dbus_error_register_error_domain(param0 string, param1 *uint64, param2 []DBusErrorEntry, param3 uint) {
}

func Fn_g_dbus_error_strip_remote_error(param0 unsafe.Pointer) {}

func Fn_g_dbus_error_unregister_error(param0 uint32, param1 int, param2 string) {}

func Fn_g_dbus_generate_guid() {
	C.g_dbus_generate_guid()
}

func Fn_g_dbus_gvalue_to_gvariant(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_dbus_gvariant_to_gvalue(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_dbus_is_address(param0 string) {}

func Fn_g_dbus_is_guid(param0 string) {}

func Fn_g_dbus_is_interface_name(param0 string) {}

func Fn_g_dbus_is_member_name(param0 string) {}

func Fn_g_dbus_is_name(param0 string) {}

func Fn_g_dbus_is_supported_address(param0 string) {}

func Fn_g_dbus_is_unique_name(param0 string) {}

func Fn_g_dtls_client_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_dtls_server_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_file_new_for_commandline_arg(param0 string) {}

func Fn_g_file_new_for_commandline_arg_and_cwd(param0 string, param1 string) {}

func Fn_g_file_new_for_path(param0 string) {}

func Fn_g_file_new_for_uri(param0 string) {}

func Fn_g_file_new_tmp(param0 string, param1 *unsafe.Pointer) {}

func Fn_g_file_parse_name(param0 string) {}

func Fn_g_icon_deserialize(param0 unsafe.Pointer) {}

func Fn_g_icon_hash(param0 unsafe.Pointer) {}

func Fn_g_icon_new_for_string(param0 string) {}

func Fn_g_initable_newv(param0 uint64, param1 uint, param2 []gobject.Parameter, param3 unsafe.Pointer) {
}

func Fn_g_io_error_from_errno(param0 int) {}

func Fn_g_io_error_quark() {
	C.g_io_error_quark()
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) {}

func Fn_g_io_extension_point_lookup(param0 string) {}

func Fn_g_io_extension_point_register(param0 string) {}

func Fn_g_io_modules_load_all_in_directory(param0 string) {}

func Fn_g_io_modules_load_all_in_directory_with_scope(param0 string, param1 unsafe.Pointer) {}

func Fn_g_io_modules_scan_all_in_directory(param0 string) {}

func Fn_g_io_modules_scan_all_in_directory_with_scope(param0 string, param1 unsafe.Pointer) {}

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
func Fn_g_pollable_source_new(param0 unsafe.Pointer) {}

func Fn_g_pollable_source_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_g_pollable_stream_read(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 unsafe.Pointer) {
}

func Fn_g_pollable_stream_write(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 unsafe.Pointer) {
}

func Fn_g_pollable_stream_write_all(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 bool, param4 *uint64, param5 unsafe.Pointer) {
}

func Fn_g_proxy_get_default_for_protocol(param0 string) {}

func Fn_g_proxy_resolver_get_default() {
	C.g_proxy_resolver_get_default()
}

func Fn_g_resolver_error_quark() {
	C.g_resolver_error_quark()
}

func Fn_g_resource_error_quark() {
	C.g_resource_error_quark()
}

func Fn_g_resource_load(param0 string) {}

func Fn_g_resources_enumerate_children(param0 string, param1 int) {}

func Fn_g_resources_get_info(param0 string, param1 int, param2 *uint64, param3 *uint32) {}

func Fn_g_resources_lookup_data(param0 string, param1 int) {}

func Fn_g_resources_open_stream(param0 string, param1 int) {}

func Fn_g_resources_register(param0 unsafe.Pointer) {}

func Fn_g_resources_unregister(param0 unsafe.Pointer) {}

func Fn_g_settings_schema_source_get_default() {
	C.g_settings_schema_source_get_default()
}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) {}

func Fn_g_tls_backend_get_default() {
	C.g_tls_backend_get_default()
}

func Fn_g_tls_client_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_tls_error_quark() {
	C.g_tls_error_quark()
}

func Fn_g_tls_file_database_new(param0 string) {}

func Fn_g_tls_server_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

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

func Fn_g_unix_mount_guess_symbolic_icon(param0 unsafe.Pointer) {}

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

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_buffered_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_output_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_bytes_icon_new(param0 unsafe.Pointer) {}

func Fn_g_cancellable_new() {
	C.g_cancellable_new()
}

func Fn_g_charset_converter_new(param0 string, param1 string) {}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_credentials_new() {
	C.g_credentials_new()
}

func Fn_g_dbus_auth_observer_new() {
	C.g_dbus_auth_observer_new()
}

func Fn_g_dbus_connection_new_finish(param0 unsafe.Pointer) {}

func Fn_g_dbus_connection_new_for_address_finish(param0 unsafe.Pointer) {}

func Fn_g_dbus_connection_new_for_address_sync(param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

func Fn_g_dbus_connection_new_sync(param0 unsafe.Pointer, param1 string, param2 int, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_g_dbus_message_new() {
	C.g_dbus_message_new()
}

func Fn_g_dbus_message_new_from_blob(param0 []uint8, param1 uint64, param2 int) {}

func Fn_g_dbus_message_new_method_call(param0 string, param1 string, param2 string, param3 string) {}

func Fn_g_dbus_message_new_signal(param0 string, param1 string, param2 string) {}

func Fn_g_dbus_object_manager_client_new_finish(param0 unsafe.Pointer) {}

func Fn_g_dbus_object_manager_client_new_for_bus_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : new_for_bus_sync : has callback

// UNSUPPORTED : new_sync : has callback

func Fn_g_dbus_object_manager_server_new(param0 string) {}

func Fn_g_dbus_object_proxy_new(param0 unsafe.Pointer, param1 string) {}

func Fn_g_dbus_object_skeleton_new(param0 string) {}

func Fn_g_dbus_proxy_new_finish(param0 unsafe.Pointer) {}

func Fn_g_dbus_proxy_new_for_bus_finish(param0 unsafe.Pointer) {}

func Fn_g_dbus_proxy_new_for_bus_sync(param0 int, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer) {
}

func Fn_g_dbus_proxy_new_sync(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 string, param4 string, param5 string, param6 unsafe.Pointer) {
}

func Fn_g_dbus_server_new_sync(param0 string, param1 int, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_g_data_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_data_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_desktop_app_info_new(param0 string) {}

func Fn_g_desktop_app_info_new_from_filename(param0 string) {}

func Fn_g_desktop_app_info_new_from_keyfile(param0 unsafe.Pointer) {}

func Fn_g_emblem_new(param0 unsafe.Pointer) {}

func Fn_g_emblem_new_with_origin(param0 unsafe.Pointer, param1 int) {}

func Fn_g_emblemed_icon_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_file_icon_new(param0 unsafe.Pointer) {}

func Fn_g_file_info_new() {
	C.g_file_info_new()
}

func Fn_g_filename_completer_new() {
	C.g_filename_completer_new()
}

func Fn_g_io_module_new(param0 string) {}

func Fn_g_inet_address_new_any(param0 int) {}

func Fn_g_inet_address_new_from_bytes(param0 []uint8, param1 int) {}

func Fn_g_inet_address_new_from_string(param0 string) {}

func Fn_g_inet_address_new_loopback(param0 int) {}

func Fn_g_inet_address_mask_new(param0 unsafe.Pointer, param1 uint) {}

func Fn_g_inet_address_mask_new_from_string(param0 string) {}

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) {}

func Fn_g_inet_socket_address_new_from_string(param0 string, param1 uint) {}

func Fn_g_list_store_new(param0 uint64) {}

func Fn_g_memory_input_stream_new() {
	C.g_memory_input_stream_new()
}

func Fn_g_memory_input_stream_new_from_bytes(param0 unsafe.Pointer) {}

// UNSUPPORTED : new_from_data : has callback

// UNSUPPORTED : new : has callback

func Fn_g_memory_output_stream_new_resizable() {
	C.g_memory_output_stream_new_resizable()
}

func Fn_g_menu_new() {
	C.g_menu_new()
}

func Fn_g_menu_item_new(param0 string, param1 string) {}

func Fn_g_menu_item_new_from_model(param0 unsafe.Pointer, param1 int) {}

func Fn_g_menu_item_new_section(param0 string, param1 unsafe.Pointer) {}

func Fn_g_menu_item_new_submenu(param0 string, param1 unsafe.Pointer) {}

func Fn_g_mount_operation_new() {
	C.g_mount_operation_new()
}

func Fn_g_native_socket_address_new(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_network_address_new(param0 string, param1 uint16) {}

func Fn_g_network_address_new_loopback(param0 uint16) {}

func Fn_g_network_service_new(param0 string, param1 string, param2 string) {}

func Fn_g_notification_new(param0 string) {}

func Fn_g_property_action_new(param0 string, param1 unsafe.Pointer, param2 string) {}

func Fn_g_proxy_address_new(param0 unsafe.Pointer, param1 uint16, param2 string, param3 string, param4 uint16, param5 string, param6 string) {
}

func Fn_g_settings_new(param0 string) {}

func Fn_g_settings_new_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {}

func Fn_g_settings_new_with_backend(param0 string, param1 unsafe.Pointer) {}

func Fn_g_settings_new_with_backend_and_path(param0 string, param1 unsafe.Pointer, param2 string) {}

func Fn_g_settings_new_with_path(param0 string, param1 string) {}

func Fn_g_simple_action_new(param0 string, param1 unsafe.Pointer) {}

func Fn_g_simple_action_new_stateful(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_g_simple_action_group_new() {
	C.g_simple_action_group_new()
}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_error : has varargs

// UNSUPPORTED : new_from_error : has callback

// UNSUPPORTED : new_take_error : has callback

func Fn_g_simple_io_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_simple_permission_new(param0 bool) {}

func Fn_g_socket_new(param0 int, param1 int, param2 int) {}

func Fn_g_socket_new_from_fd(param0 int) {}

func Fn_g_socket_address_new_from_native(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_socket_client_new() {
	C.g_socket_client_new()
}

func Fn_g_socket_listener_new() {
	C.g_socket_listener_new()
}

func Fn_g_socket_service_new() {
	C.g_socket_service_new()
}

// UNSUPPORTED : new : has varargs

func Fn_g_subprocess_newv(param0 []string, param1 int) {}

func Fn_g_subprocess_launcher_new(param0 int) {}

// UNSUPPORTED : new : has callback

func Fn_g_tcp_wrapper_connection_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_test_dbus_new(param0 int) {}

func Fn_g_themed_icon_new(param0 string) {}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) {}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) {}

func Fn_g_threaded_socket_service_new(param0 int) {}

func Fn_g_tls_certificate_new_from_file(param0 string) {}

func Fn_g_tls_certificate_new_from_files(param0 string, param1 string) {}

func Fn_g_tls_certificate_new_from_pem(param0 string, param1 uint64) {}

func Fn_g_tls_password_new(param0 int, param1 string) {}

func Fn_g_unix_credentials_message_new() {
	C.g_unix_credentials_message_new()
}

func Fn_g_unix_credentials_message_new_with_credentials(param0 unsafe.Pointer) {}

func Fn_g_unix_fd_list_new() {
	C.g_unix_fd_list_new()
}

func Fn_g_unix_fd_list_new_from_array(param0 []int, param1 int) {}

func Fn_g_unix_fd_message_new() {
	C.g_unix_fd_message_new()
}

func Fn_g_unix_fd_message_new_with_fd_list(param0 unsafe.Pointer) {}

func Fn_g_unix_input_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_mount_monitor_new() {
	C.g_unix_mount_monitor_new()
}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_socket_address_new(param0 string) {}

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) {}

func Fn_g_unix_socket_address_new_with_type(param0 []int8, param1 int, param2 int) {}

func Fn_g_zlib_compressor_new(param0 int, param1 int) {}

func Fn_g_zlib_decompressor_new(param0 int) {}
