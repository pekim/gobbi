// Code generated - DO NOT EDIT.
// +build gio_2.20

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

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
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

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int) {
	// has string param
}

func Fn_g_app_info_get_all() {

	C.g_app_info_get_all()
}

func Fn_g_app_info_get_all_for_type(param0 string) {
	// has string param
}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) {
	// has string param
}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) {
	// has string param
}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer) {
	// has string param
}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_g_app_info_reset_type_associations(param0 string) {
	// has string param
}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_g_content_type_can_be_executable(param0 string) {
	// has string param
}

func Fn_g_content_type_equals(param0 string, param1 string) {
	// has string param
}

func Fn_g_content_type_from_mime_type(param0 string) {
	// has string param
}

func Fn_g_content_type_get_description(param0 string) {
	// has string param
}

func Fn_g_content_type_get_icon(param0 string) {
	// has string param
}

func Fn_g_content_type_get_mime_type(param0 string) {
	// has string param
}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) {
	// has array param
}

func Fn_g_content_type_guess_for_tree(param0 unsafe.Pointer) {
	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.g_content_type_guess_for_tree(cValue0)
}

func Fn_g_content_type_is_a(param0 string, param1 string) {
	// has string param
}

func Fn_g_content_type_is_unknown(param0 string) {
	// has string param
}

func Fn_g_content_types_get_registered() {

	C.g_content_types_get_registered()
}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_g_dbus_error_quark() {

	C.g_dbus_error_quark()
}

func Fn_g_file_new_for_commandline_arg(param0 string) {
	// has string param
}

func Fn_g_file_new_for_path(param0 string) {
	// has string param
}

func Fn_g_file_new_for_uri(param0 string) {
	// has string param
}

func Fn_g_file_parse_name(param0 string) {
	// has string param
}

func Fn_g_icon_hash(param0 unsafe.Pointer) {
	cValue0 := (C.gconstpointer)(param0)

	C.g_icon_hash(cValue0)
}

func Fn_g_icon_new_for_string(param0 string) {
	// has string param
}

func Fn_g_io_error_from_errno(param0 int) {
	cValue0 := (C.gint)(param0)

	C.g_io_error_from_errno(cValue0)
}

func Fn_g_io_error_quark() {

	C.g_io_error_quark()
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) {
	// has string param
}

func Fn_g_io_extension_point_lookup(param0 string) {
	// has string param
}

func Fn_g_io_extension_point_register(param0 string) {
	// has string param
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
// UNSUPPORTED : null_settings_backend_new : blacklisted
// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_g_unix_is_mount_path_system_internal(param0 string) {
	// has string param
}

func Fn_g_unix_mount_at(param0 string, param1 *uint64) {
	// has string param
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

func Fn_g_app_launch_context_get_startup_notify_id(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GAppLaunchContext)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAppInfo)(unsafe.Pointer(param0))
	cValue1 := (*C.GList)(unsafe.Pointer(param1))

	C.g_app_launch_context_get_startup_notify_id(cValueInstance, cValue0, cValue1)
}

func Fn_g_app_launch_context_launch_failed(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_application_new(param0 string, param1 int) {
	// has string param
}

func Fn_g_application_hold(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_hold(cValueInstance)
}

func Fn_g_application_release(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GApplication)(unsafe.Pointer(paramInstance))

	C.g_application_release(cValueInstance)
}

func Fn_g_application_id_is_valid(param0 string) {
	// has string param
}

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))

	C.g_buffered_input_stream_new(cValue0)
}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (C.gsize)(param1)

	C.g_buffered_input_stream_new_sized(cValue0, cValue1)
}

func Fn_g_buffered_input_stream_fill(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gssize)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_buffered_input_stream_fill(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : fill_async : has callback

func Fn_g_buffered_input_stream_fill_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_buffered_input_stream_fill_finish(cValueInstance, cValue0)
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

func Fn_g_buffered_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GBufferedInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_buffered_input_stream_read_byte(cValueInstance, cValue0)
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

func Fn_g_cancellable_new() {

	C.g_cancellable_new()
}

func Fn_g_cancellable_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_cancel(cValueInstance)
}

// UNSUPPORTED : connect : has callback

func Fn_g_cancellable_get_fd(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_get_fd(cValueInstance)
}

func Fn_g_cancellable_is_cancelled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_is_cancelled(cValueInstance)
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

func Fn_g_cancellable_set_error_if_cancelled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GCancellable)(unsafe.Pointer(paramInstance))

	C.g_cancellable_set_error_if_cancelled(cValueInstance)
}

func Fn_g_cancellable_get_current() {

	C.g_cancellable_get_current()
}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	C.g_converter_input_stream_new(cValue0, cValue1)
}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GOutputStream)(unsafe.Pointer(param0))
	cValue1 := (*C.GConverter)(unsafe.Pointer(param1))

	C.g_converter_output_stream_new(cValue0, cValue1)
}

// UNSUPPORTED : add_filter : has callback

// UNSUPPORTED : call : has callback

// UNSUPPORTED : call_with_unix_fd_list : has callback

// UNSUPPORTED : close : has callback

// UNSUPPORTED : flush : has callback

// UNSUPPORTED : register_object : has callback

// UNSUPPORTED : register_subtree : has callback

// UNSUPPORTED : send_message_with_reply : has callback

// UNSUPPORTED : signal_subscribe : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_address : has callback

func Fn_g_dbus_message_get_byte_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))

	C.g_dbus_message_get_byte_order(cValueInstance)
}

// UNSUPPORTED : new_method_error : has varargs

// UNSUPPORTED : new_method_error_valist : has va_list

func Fn_g_dbus_message_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDBusMessage)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDBusMessageByteOrder)(param0)

	C.g_dbus_message_set_byte_order(cValueInstance, cValue0)
}

// UNSUPPORTED : return_error : has varargs

// UNSUPPORTED : return_error_valist : has va_list

// UNSUPPORTED : new_for_bus_sync : has callback

// UNSUPPORTED : new_sync : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_dbus_object_manager_server_set_connection(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDBusObjectManagerServer)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GDBusConnection)(unsafe.Pointer(param0))

	C.g_dbus_object_manager_server_set_connection(cValueInstance, cValue0)
}

// UNSUPPORTED : call : has callback

// UNSUPPORTED : call_with_unix_fd_list : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

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

func Fn_g_data_input_stream_read_byte(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_byte(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_int16(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_int16(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_int32(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_int32(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_int64(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_int64(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_line(paramInstance unsafe.Pointer, param0 *uint64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.gsize)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_input_stream_read_line(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : read_line_async : has callback

func Fn_g_data_input_stream_read_line_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	C.g_data_input_stream_read_line_finish(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_input_stream_read_uint16(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_uint16(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_uint32(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_uint32(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_uint64(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_data_input_stream_read_uint64(cValueInstance, cValue0)
}

func Fn_g_data_input_stream_read_until(paramInstance unsafe.Pointer, param0 string, param1 *uint64, param2 unsafe.Pointer) {
	// has string param
}

// UNSUPPORTED : read_until_async : has callback

func Fn_g_data_input_stream_read_until_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *uint64) {
	cValueInstance := (*C.GDataInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.gsize)(unsafe.Pointer(param1))

	C.g_data_input_stream_read_until_finish(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : read_upto_async : has callback

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

func Fn_g_data_output_stream_put_byte(paramInstance unsafe.Pointer, param0 uint8, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guchar)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_byte(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_int16(paramInstance unsafe.Pointer, param0 int16, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint16)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_int16(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_int32(paramInstance unsafe.Pointer, param0 int32, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint32)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_int32(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_int64(paramInstance unsafe.Pointer, param0 int64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gint64)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_int64(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_string(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	// has string param
}

func Fn_g_data_output_stream_put_uint16(paramInstance unsafe.Pointer, param0 uint16, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint16)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_uint16(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_uint32(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint32)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_uint32(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_put_uint64(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.guint64)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_data_output_stream_put_uint64(cValueInstance, cValue0, cValue1)
}

func Fn_g_data_output_stream_set_byte_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GDataOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GDataStreamByteOrder)(param0)

	C.g_data_output_stream_set_byte_order(cValueInstance, cValue0)
}

func Fn_g_desktop_app_info_new(param0 string) {
	// has string param
}

func Fn_g_desktop_app_info_new_from_filename(param0 string) {
	// has string param
}

func Fn_g_desktop_app_info_new_from_keyfile(param0 unsafe.Pointer) {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	C.g_desktop_app_info_new_from_keyfile(cValue0)
}

func Fn_g_desktop_app_info_get_categories(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_categories(cValueInstance)
}

func Fn_g_desktop_app_info_get_generic_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_generic_name(cValueInstance)
}

func Fn_g_desktop_app_info_get_is_hidden(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GDesktopAppInfo)(unsafe.Pointer(paramInstance))

	C.g_desktop_app_info_get_is_hidden(cValueInstance)
}

// UNSUPPORTED : launch_uris_as_manager : has callback

// UNSUPPORTED : launch_uris_as_manager_with_fds : has callback

func Fn_g_desktop_app_info_search(param0 string) {
	// has string param
}

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {
	// has string param
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

func Fn_g_emblemed_icon_get_emblems(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	C.g_emblemed_icon_get_emblems(cValueInstance)
}

func Fn_g_emblemed_icon_get_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GEmblemedIcon)(unsafe.Pointer(paramInstance))

	C.g_emblemed_icon_get_icon(cValueInstance)
}

func Fn_g_file_enumerator_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_file_enumerator_close(cValueInstance, cValue0)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_file_enumerator_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_file_enumerator_close_finish(cValueInstance, cValue0)
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

func Fn_g_file_enumerator_next_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_file_enumerator_next_file(cValueInstance, cValue0)
}

// UNSUPPORTED : next_files_async : has callback

func Fn_g_file_enumerator_next_files_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_file_enumerator_next_files_finish(cValueInstance, cValue0)
}

func Fn_g_file_enumerator_set_pending(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GFileEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_file_enumerator_set_pending(cValueInstance, cValue0)
}

// UNSUPPORTED : query_info_async : has callback

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
	// has string param
}

func Fn_g_file_info_get_attribute_boolean(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_byte_string(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_data(paramInstance unsafe.Pointer, param0 string, param1 *int, param2 *unsafe.Pointer, param3 *int) {
	// has string param
}

func Fn_g_file_info_get_attribute_int32(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_int64(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_object(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_status(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_string(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_type(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_uint32(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_attribute_uint64(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_get_content_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_content_type(cValueInstance)
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

func Fn_g_file_info_get_symlink_target(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_get_symlink_target(cValueInstance)
}

func Fn_g_file_info_has_attribute(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_list_attributes(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_remove_attribute(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_set_attribute(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	// has string param
}

func Fn_g_file_info_set_attribute_boolean(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	// has string param
}

func Fn_g_file_info_set_attribute_byte_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	// has string param
}

func Fn_g_file_info_set_attribute_int32(paramInstance unsafe.Pointer, param0 string, param1 int32) {
	// has string param
}

func Fn_g_file_info_set_attribute_int64(paramInstance unsafe.Pointer, param0 string, param1 int64) {
	// has string param
}

func Fn_g_file_info_set_attribute_mask(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GFileAttributeMatcher)(unsafe.Pointer(param0))

	C.g_file_info_set_attribute_mask(cValueInstance, cValue0)
}

func Fn_g_file_info_set_attribute_object(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	// has string param
}

func Fn_g_file_info_set_attribute_string(paramInstance unsafe.Pointer, param0 string, param1 string) {
	// has string param
}

func Fn_g_file_info_set_attribute_stringv(paramInstance unsafe.Pointer, param0 string, param1 []string) {
	// has array param
}

func Fn_g_file_info_set_attribute_uint32(paramInstance unsafe.Pointer, param0 string, param1 uint32) {
	// has string param
}

func Fn_g_file_info_set_attribute_uint64(paramInstance unsafe.Pointer, param0 string, param1 uint64) {
	// has string param
}

func Fn_g_file_info_set_content_type(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_set_display_name(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_file_info_set_edit_name(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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
	// has string param
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
	// has string param
}

func Fn_g_file_info_unset_attribute_mask(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GFileInfo)(unsafe.Pointer(paramInstance))

	C.g_file_info_unset_attribute_mask(cValueInstance)
}

func Fn_g_file_input_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	// has string param
}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_input_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_file_input_stream_query_info_finish(cValueInstance, cValue0)
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

func Fn_g_file_output_stream_query_info(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	// has string param
}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_output_stream_query_info_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GFileOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_file_output_stream_query_info_finish(cValueInstance, cValue0)
}

func Fn_g_filename_completer_new() {

	C.g_filename_completer_new()
}

func Fn_g_filename_completer_get_completion_suffix(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_filename_completer_get_completions(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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
// UNSUPPORTED : close_async : has callback

// UNSUPPORTED : splice_async : has callback

func Fn_g_input_stream_clear_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_clear_pending(cValueInstance)
}

func Fn_g_input_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_input_stream_close(cValueInstance, cValue0)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_input_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_input_stream_close_finish(cValueInstance, cValue0)
}

func Fn_g_input_stream_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_has_pending(cValueInstance)
}

func Fn_g_input_stream_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_is_closed(cValueInstance)
}

func Fn_g_input_stream_read(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer) {
	// has array param
}

func Fn_g_input_stream_read_all(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : read_all_async : has callback

// UNSUPPORTED : read_async : has callback

// UNSUPPORTED : read_bytes_async : has callback

func Fn_g_input_stream_read_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_input_stream_read_finish(cValueInstance, cValue0)
}

func Fn_g_input_stream_set_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))

	C.g_input_stream_set_pending(cValueInstance)
}

func Fn_g_input_stream_skip(paramInstance unsafe.Pointer, param0 uint64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (C.gsize)(param0)
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_input_stream_skip(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : skip_async : has callback

func Fn_g_input_stream_skip_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GInputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_input_stream_skip_finish(cValueInstance, cValue0)
}

// UNSUPPORTED : insert_sorted : has callback

// UNSUPPORTED : sort : has callback

func Fn_g_memory_input_stream_new() {

	C.g_memory_input_stream_new()
}

// UNSUPPORTED : new_from_data : has callback

// UNSUPPORTED : add_data : has callback

// UNSUPPORTED : new : has callback

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

// UNSUPPORTED : get_attribute : has varargs

// UNSUPPORTED : set_action_and_target : has varargs

// UNSUPPORTED : set_attribute : has varargs

// UNSUPPORTED : get_item_attribute : has varargs

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
	// has string param
}

func Fn_g_mount_operation_set_password(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_mount_operation_set_password_save(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GMountOperation)(unsafe.Pointer(paramInstance))
	cValue0 := (C.GPasswordSave)(param0)

	C.g_mount_operation_set_password_save(cValueInstance, cValue0)
}

func Fn_g_mount_operation_set_username(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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

func Fn_g_output_stream_close(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_output_stream_close(cValueInstance, cValue0)
}

// UNSUPPORTED : close_async : has callback

func Fn_g_output_stream_close_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_output_stream_close_finish(cValueInstance, cValue0)
}

func Fn_g_output_stream_flush(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_output_stream_flush(cValueInstance, cValue0)
}

// UNSUPPORTED : flush_async : has callback

func Fn_g_output_stream_flush_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_output_stream_flush_finish(cValueInstance, cValue0)
}

func Fn_g_output_stream_has_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_has_pending(cValueInstance)
}

func Fn_g_output_stream_is_closed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_is_closed(cValueInstance)
}

// UNSUPPORTED : printf : has varargs

func Fn_g_output_stream_set_pending(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))

	C.g_output_stream_set_pending(cValueInstance)
}

func Fn_g_output_stream_splice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GInputStream)(unsafe.Pointer(param0))
	cValue1 := (C.GOutputStreamSpliceFlags)(param1)
	cValue2 := (*C.GCancellable)(unsafe.Pointer(param2))

	C.g_output_stream_splice(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : splice_async : has callback

func Fn_g_output_stream_splice_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_output_stream_splice_finish(cValueInstance, cValue0)
}

// UNSUPPORTED : vprintf : has va_list

func Fn_g_output_stream_write(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 unsafe.Pointer) {
	// has array param
}

func Fn_g_output_stream_write_all(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer) {
	// has array param
}

// UNSUPPORTED : write_all_async : has callback

// UNSUPPORTED : write_async : has callback

func Fn_g_output_stream_write_bytes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GBytes)(unsafe.Pointer(param0))
	cValue1 := (*C.GCancellable)(unsafe.Pointer(param1))

	C.g_output_stream_write_bytes(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : write_bytes_async : has callback

func Fn_g_output_stream_write_bytes_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_output_stream_write_bytes_finish(cValueInstance, cValue0)
}

func Fn_g_output_stream_write_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GOutputStream)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_output_stream_write_finish(cValueInstance, cValue0)
}

// UNSUPPORTED : writev_all_async : has callback

// UNSUPPORTED : writev_async : has callback

// UNSUPPORTED : acquire_async : has callback

// UNSUPPORTED : release_async : has callback

// UNSUPPORTED : lookup_by_address_async : has callback

// UNSUPPORTED : lookup_by_name_async : has callback

// UNSUPPORTED : lookup_by_name_with_flags_async : has callback

// UNSUPPORTED : lookup_records_async : has callback

// UNSUPPORTED : lookup_service_async : has callback

func Fn_g_settings_apply(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_apply(cValueInstance)
}

// UNSUPPORTED : bind_with_mapping : has callback

// UNSUPPORTED : get : has varargs

// UNSUPPORTED : get_mapped : has callback

func Fn_g_settings_list_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_list_children(cValueInstance)
}

func Fn_g_settings_list_keys(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_list_keys(cValueInstance)
}

func Fn_g_settings_reset(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_settings_revert(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSettings)(unsafe.Pointer(paramInstance))

	C.g_settings_revert(cValueInstance)
}

// UNSUPPORTED : set : has varargs

func Fn_g_settings_set_enum(paramInstance unsafe.Pointer, param0 string, param1 int) {
	// has string param
}

func Fn_g_settings_set_flags(paramInstance unsafe.Pointer, param0 string, param1 uint) {
	// has string param
}

func Fn_g_settings_sync() {

	C.g_settings_sync()
}

// UNSUPPORTED : changed : blacklisted
// UNSUPPORTED : changed_tree : blacklisted
// UNSUPPORTED : keys_changed : blacklisted
// UNSUPPORTED : path_changed : blacklisted
// UNSUPPORTED : path_writable_changed : blacklisted
// UNSUPPORTED : writable_changed : blacklisted
// UNSUPPORTED : flatten_tree : blacklisted
// UNSUPPORTED : get_default : blacklisted
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

func Fn_g_simple_async_result_propagate_error(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GSimpleAsyncResult)(unsafe.Pointer(paramInstance))

	C.g_simple_async_result_propagate_error(cValueInstance)
}

// UNSUPPORTED : run_in_thread : has callback

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

func Fn_g_simple_async_result_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (*C.GObject)(unsafe.Pointer(param1))
	cValue2 := (C.gpointer)(param2)

	C.g_simple_async_result_is_valid(cValue0, cValue1, cValue2)
}

func Fn_g_socket_address_enumerator_next(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GCancellable)(unsafe.Pointer(param0))

	C.g_socket_address_enumerator_next(cValueInstance, cValue0)
}

// UNSUPPORTED : next_async : has callback

func Fn_g_socket_address_enumerator_next_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GSocketAddressEnumerator)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))

	C.g_socket_address_enumerator_next_finish(cValueInstance, cValue0)
}

func Fn_g_socket_client_add_application_proxy(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

// UNSUPPORTED : connect_async : has callback

// UNSUPPORTED : connect_to_host_async : has callback

func Fn_g_socket_client_connect_to_service(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	// has string param
}

// UNSUPPORTED : connect_to_service_async : has callback

// UNSUPPORTED : connect_to_uri_async : has callback

// UNSUPPORTED : connect_async : has callback

// UNSUPPORTED : accept_async : has callback

// UNSUPPORTED : accept_socket_async : has callback

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : communicate_async : has callback

func Fn_g_subprocess_communicate_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
	cValueInstance := (*C.GSubprocess)(unsafe.Pointer(paramInstance))
	cValue0 := (*C.GAsyncResult)(unsafe.Pointer(param0))
	cValue1 := (**C.GBytes)(unsafe.Pointer(param1))
	cValue2 := (**C.GBytes)(unsafe.Pointer(param2))

	C.g_subprocess_communicate_finish(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_g_subprocess_communicate_utf8(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 string, param3 string) {
	// has string param
}

// UNSUPPORTED : communicate_utf8_async : has callback

func Fn_g_subprocess_communicate_utf8_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 string) {
	// has string param
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

// UNSUPPORTED : return_new_error : has varargs

// UNSUPPORTED : return_pointer : has callback

// UNSUPPORTED : run_in_thread : has callback

// UNSUPPORTED : run_in_thread_sync : has callback

// UNSUPPORTED : set_task_data : has callback

// UNSUPPORTED : report_error : has callback

// UNSUPPORTED : report_new_error : has varargs

func Fn_g_tcp_wrapper_connection_get_base_io_stream(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTcpWrapperConnection)(unsafe.Pointer(paramInstance))

	C.g_tcp_wrapper_connection_get_base_io_stream(cValueInstance)
}

func Fn_g_test_dbus_new(param0 int) {
	cValue0 := (C.GTestDBusFlags)(param0)

	C.g_test_dbus_new(cValue0)
}

func Fn_g_test_dbus_add_service_dir(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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
	// has string param
}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) {
	// has array param
}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) {
	// has string param
}

func Fn_g_themed_icon_append_name(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_themed_icon_get_names(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GThemedIcon)(unsafe.Pointer(paramInstance))

	C.g_themed_icon_get_names(cValueInstance)
}

func Fn_g_themed_icon_prepend_name(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_tls_connection_get_use_system_certdb(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))

	C.g_tls_connection_get_use_system_certdb(cValueInstance)
}

// UNSUPPORTED : handshake_async : has callback

func Fn_g_tls_connection_set_use_system_certdb(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GTlsConnection)(unsafe.Pointer(paramInstance))
	cValue0 := toCBool(param0)

	C.g_tls_connection_set_use_system_certdb(cValueInstance, cValue0)
}

// UNSUPPORTED : lookup_certificate_for_handle_async : has callback

// UNSUPPORTED : lookup_certificate_issuer_async : has callback

// UNSUPPORTED : lookup_certificates_issued_by_async : has callback

// UNSUPPORTED : verify_chain_async : has callback

// UNSUPPORTED : ask_password_async : has callback

// UNSUPPORTED : request_certificate_async : has callback

func Fn_g_tls_password_new(param0 int, param1 string) {
	// has string param
}

// UNSUPPORTED : set_value_full : has callback

// UNSUPPORTED : receive_credentials_async : has callback

// UNSUPPORTED : send_credentials_async : has callback

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

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) {
	// has array param
}

func Fn_g_vfs_get_file_for_path(paramInstance unsafe.Pointer, param0 string) {
	// has string param
}

func Fn_g_vfs_get_file_for_uri(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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
	// has string param
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
	// has string param
}

func Fn_g_volume_monitor_get_mounts(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GVolumeMonitor)(unsafe.Pointer(paramInstance))

	C.g_volume_monitor_get_mounts(cValueInstance)
}

func Fn_g_volume_monitor_get_volume_for_uuid(paramInstance unsafe.Pointer, param0 string) {
	// has string param
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
