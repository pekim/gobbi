// This is a generated file - DO NOT EDIT

package gio

import "unsafe"

// ActionEntry is a wrapper around the C record GActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
	Name   string
	// no type for activate
	ParameterType string
	State         string
	// no type for change_state
	// Private : padding
}

// AppInfoIface is a wrapper around the C record GAppInfoIface.
type AppInfoIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for dup
	// no type for equal
	// no type for get_id
	// no type for get_name
	// no type for get_description
	// no type for get_executable
	// no type for get_icon
	// no type for launch
	// no type for supports_uris
	// no type for supports_files
	// no type for launch_uris
	// no type for should_show
	// no type for set_as_default_for_type
	// no type for set_as_default_for_extension
	// no type for add_supports_type
	// no type for can_remove_supports_type
	// no type for remove_supports_type
	// no type for can_delete
	// no type for do_delete
	// no type for get_commandline
	// no type for get_display_name
	// no type for set_as_last_used_for_type
	// no type for get_supported_types
}

// AppLaunchContextClass is a wrapper around the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_display
	// no type for get_startup_notify_id
	// no type for launch_failed
	// no type for launched
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

// AppLaunchContextPrivate is a wrapper around the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

// ApplicationCommandLinePrivate is a wrapper around the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

// ApplicationPrivate is a wrapper around the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// AsyncResultIface is a wrapper around the C record GAsyncResultIface.
type AsyncResultIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for get_user_data
	// no type for get_source_object
	// no type for is_tagged
}

// BufferedInputStreamClass is a wrapper around the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for fill
	// no type for fill_async
	// no type for fill_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// BufferedInputStreamPrivate is a wrapper around the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

// BufferedOutputStreamClass is a wrapper around the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

// BufferedOutputStreamPrivate is a wrapper around the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

// CancellableClass is a wrapper around the C record GCancellableClass.
type CancellableClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for cancelled
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// CancellablePrivate is a wrapper around the C record GCancellablePrivate.
type CancellablePrivate struct {
	native unsafe.Pointer
}

// CharsetConverterClass is a wrapper around the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ConverterInputStreamClass is a wrapper around the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// ConverterInputStreamPrivate is a wrapper around the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

// ConverterOutputStreamClass is a wrapper around the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// ConverterOutputStreamPrivate is a wrapper around the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

// DBusInterfaceSkeletonPrivate is a wrapper around the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

// DBusObjectManagerClientPrivate is a wrapper around the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

// DBusObjectManagerServerPrivate is a wrapper around the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

// DBusObjectProxyPrivate is a wrapper around the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

// DBusObjectSkeletonPrivate is a wrapper around the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

// DBusProxyPrivate is a wrapper around the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native unsafe.Pointer
}

// DataInputStreamClass is a wrapper around the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// DataInputStreamPrivate is a wrapper around the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

// DataOutputStreamClass is a wrapper around the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// DataOutputStreamPrivate is a wrapper around the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

// DesktopAppInfoClass is a wrapper around the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// DesktopAppInfoLookupIface is a wrapper around the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for get_default_for_uri_scheme
}

// DriveIface is a wrapper around the C record GDriveIface.
type DriveIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for changed
	// no type for disconnected
	// no type for eject_button
	// no type for get_name
	// no type for get_icon
	// no type for has_volumes
	// no type for get_volumes
	// no type for is_media_removable
	// no type for has_media
	// no type for is_media_check_automatic
	// no type for can_eject
	// no type for can_poll_for_media
	// no type for eject
	// no type for eject_finish
	// no type for poll_for_media
	// no type for poll_for_media_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for get_start_stop_type
	// no type for can_start
	// no type for can_start_degraded
	// no type for start
	// no type for start_finish
	// no type for can_stop
	// no type for stop
	// no type for stop_finish
	// no type for stop_button
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
	// no type for is_removable
}

// EmblemClass is a wrapper around the C record GEmblemClass.
type EmblemClass struct {
	native unsafe.Pointer
}

// EmblemedIconClass is a wrapper around the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// EmblemedIconPrivate is a wrapper around the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

// FileAttributeInfo is a wrapper around the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native unsafe.Pointer
	Name   string
	Type   FileAttributeType
	Flags  FileAttributeInfoFlags
}

// FileAttributeInfoList is a wrapper around the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native unsafe.Pointer
	// infos : record
	NInfos int32
}

// FileAttributeMatcher is a wrapper around the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native unsafe.Pointer
}

// FileDescriptorBasedIface is a wrapper around the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for get_fd
}

// FileEnumeratorClass is a wrapper around the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for next_file
	// no type for close_fn
	// no type for next_files_async
	// no type for next_files_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

// FileEnumeratorPrivate is a wrapper around the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

// FileIOStreamClass is a wrapper around the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// FileIOStreamPrivate is a wrapper around the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

// FileIconClass is a wrapper around the C record GFileIconClass.
type FileIconClass struct {
	native unsafe.Pointer
}

// FileIface is a wrapper around the C record GFileIface.
type FileIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for dup
	// no type for hash
	// no type for equal
	// no type for is_native
	// no type for has_uri_scheme
	// no type for get_uri_scheme
	// no type for get_basename
	// no type for get_path
	// no type for get_uri
	// no type for get_parse_name
	// no type for get_parent
	// no type for prefix_matches
	// no type for get_relative_path
	// no type for resolve_relative_path
	// no type for get_child_for_display_name
	// no type for enumerate_children
	// no type for enumerate_children_async
	// no type for enumerate_children_finish
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for query_filesystem_info
	// no type for query_filesystem_info_async
	// no type for query_filesystem_info_finish
	// no type for find_enclosing_mount
	// no type for find_enclosing_mount_async
	// no type for find_enclosing_mount_finish
	// no type for set_display_name
	// no type for set_display_name_async
	// no type for set_display_name_finish
	// no type for query_settable_attributes
	// no type for _query_settable_attributes_async
	// no type for _query_settable_attributes_finish
	// no type for query_writable_namespaces
	// no type for _query_writable_namespaces_async
	// no type for _query_writable_namespaces_finish
	// no type for set_attribute
	// no type for set_attributes_from_info
	// no type for set_attributes_async
	// no type for set_attributes_finish
	// no type for read_fn
	// no type for read_async
	// no type for read_finish
	// no type for append_to
	// no type for append_to_async
	// no type for append_to_finish
	// no type for create
	// no type for create_async
	// no type for create_finish
	// no type for replace
	// no type for replace_async
	// no type for replace_finish
	// no type for delete_file
	// no type for delete_file_async
	// no type for delete_file_finish
	// no type for trash
	// no type for trash_async
	// no type for trash_finish
	// no type for make_directory
	// no type for make_directory_async
	// no type for make_directory_finish
	// no type for make_symbolic_link
	// no type for _make_symbolic_link_async
	// no type for _make_symbolic_link_finish
	// no type for copy
	// no type for copy_async
	// no type for copy_finish
	// no type for move
	// no type for _move_async
	// no type for _move_finish
	// no type for mount_mountable
	// no type for mount_mountable_finish
	// no type for unmount_mountable
	// no type for unmount_mountable_finish
	// no type for eject_mountable
	// no type for eject_mountable_finish
	// no type for mount_enclosing_volume
	// no type for mount_enclosing_volume_finish
	// no type for monitor_dir
	// no type for monitor_file
	// no type for open_readwrite
	// no type for open_readwrite_async
	// no type for open_readwrite_finish
	// no type for create_readwrite
	// no type for create_readwrite_async
	// no type for create_readwrite_finish
	// no type for replace_readwrite
	// no type for replace_readwrite_async
	// no type for replace_readwrite_finish
	// no type for start_mountable
	// no type for start_mountable_finish
	// no type for stop_mountable
	// no type for stop_mountable_finish
	SupportsThreadContexts bool
	// no type for unmount_mountable_with_operation
	// no type for unmount_mountable_with_operation_finish
	// no type for eject_mountable_with_operation
	// no type for eject_mountable_with_operation_finish
	// no type for poll_mountable
	// no type for poll_mountable_finish
	// no type for measure_disk_usage
	// no type for measure_disk_usage_async
	// no type for measure_disk_usage_finish
}

// FileInfoClass is a wrapper around the C record GFileInfoClass.
type FileInfoClass struct {
	native unsafe.Pointer
}

// FileInputStreamClass is a wrapper around the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// FileInputStreamPrivate is a wrapper around the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

// FileMonitorClass is a wrapper around the C record GFileMonitorClass.
type FileMonitorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for cancel
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// FileMonitorPrivate is a wrapper around the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native unsafe.Pointer
}

// FileOutputStreamClass is a wrapper around the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// FileOutputStreamPrivate is a wrapper around the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

// FilenameCompleterClass is a wrapper around the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for got_completion_data
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

// FilterInputStreamClass is a wrapper around the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

// FilterOutputStreamClass is a wrapper around the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

// IOExtension is a wrapper around the C record GIOExtension.
type IOExtension struct {
	native unsafe.Pointer
}

// IOExtensionPoint is a wrapper around the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native unsafe.Pointer
}

// IOModuleClass is a wrapper around the C record GIOModuleClass.
type IOModuleClass struct {
	native unsafe.Pointer
}

// IOSchedulerJob is a wrapper around the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native unsafe.Pointer
}

// IOStreamAdapter is a wrapper around the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native unsafe.Pointer
}

// IOStreamClass is a wrapper around the C record GIOStreamClass.
type IOStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_input_stream
	// no type for get_output_stream
	// no type for close_fn
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

// IOStreamPrivate is a wrapper around the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native unsafe.Pointer
}

// IconIface is a wrapper around the C record GIconIface.
type IconIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for hash
	// no type for equal
	// no type for to_tokens
	// no type for from_tokens
	// no type for serialize
}

// InetAddressClass is a wrapper around the C record GInetAddressClass.
type InetAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for to_string
	// no type for to_bytes
}

// InetAddressMaskClass is a wrapper around the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// InetAddressMaskPrivate is a wrapper around the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

// InetAddressPrivate is a wrapper around the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native unsafe.Pointer
}

// InetSocketAddressClass is a wrapper around the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// InetSocketAddressPrivate is a wrapper around the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

// InputStreamClass is a wrapper around the C record GInputStreamClass.
type InputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for read_fn
	// no type for skip
	// no type for close_fn
	// no type for read_async
	// no type for read_finish
	// no type for skip_async
	// no type for skip_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// InputStreamPrivate is a wrapper around the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native unsafe.Pointer
}

// ListStoreClass is a wrapper around the C record GListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// LoadableIconIface is a wrapper around the C record GLoadableIconIface.
type LoadableIconIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for load
	// no type for load_async
	// no type for load_finish
}

// MemoryInputStreamClass is a wrapper around the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// MemoryInputStreamPrivate is a wrapper around the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

// MemoryOutputStreamClass is a wrapper around the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// MemoryOutputStreamPrivate is a wrapper around the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

// MenuAttributeIterClass is a wrapper around the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_next
}

// MenuAttributeIterPrivate is a wrapper around the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

// MenuLinkIterClass is a wrapper around the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_next
}

// MenuLinkIterPrivate is a wrapper around the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

// MenuModelClass is a wrapper around the C record GMenuModelClass.
type MenuModelClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for is_mutable
	// no type for get_n_items
	// no type for get_item_attributes
	// no type for iterate_item_attributes
	// no type for get_item_attribute_value
	// no type for get_item_links
	// no type for iterate_item_links
	// no type for get_item_link
}

// MenuModelPrivate is a wrapper around the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native unsafe.Pointer
}

// MountIface is a wrapper around the C record GMountIface.
type MountIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for changed
	// no type for unmounted
	// no type for get_root
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_volume
	// no type for get_drive
	// no type for can_unmount
	// no type for can_eject
	// no type for unmount
	// no type for unmount_finish
	// no type for eject
	// no type for eject_finish
	// no type for remount
	// no type for remount_finish
	// no type for guess_content_type
	// no type for guess_content_type_finish
	// no type for guess_content_type_sync
	// no type for pre_unmount
	// no type for unmount_with_operation
	// no type for unmount_with_operation_finish
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_default_location
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

// MountOperationClass is a wrapper around the C record GMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for ask_password
	// no type for ask_question
	// no type for reply
	// no type for aborted
	// no type for show_processes
	// no type for show_unmount_progress
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
}

// MountOperationPrivate is a wrapper around the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// NativeSocketAddress is a wrapper around the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native unsafe.Pointer
}

// NativeVolumeMonitorClass is a wrapper around the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_mount_for_mount_path
}

// NetworkAddressClass is a wrapper around the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NetworkAddressPrivate is a wrapper around the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

// NetworkServiceClass is a wrapper around the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NetworkServicePrivate is a wrapper around the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native unsafe.Pointer
}

// OutputStreamClass is a wrapper around the C record GOutputStreamClass.
type OutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for write_fn
	// no type for splice
	// no type for flush
	// no type for close_fn
	// no type for write_async
	// no type for write_finish
	// no type for splice_async
	// no type for splice_finish
	// no type for flush_async
	// no type for flush_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
}

// OutputStreamPrivate is a wrapper around the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native unsafe.Pointer
}

// PermissionClass is a wrapper around the C record GPermissionClass.
type PermissionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for acquire
	// no type for acquire_async
	// no type for acquire_finish
	// no type for release
	// no type for release_async
	// no type for release_finish
	// no type for reserved
}

// PermissionPrivate is a wrapper around the C record GPermissionPrivate.
type PermissionPrivate struct {
	native unsafe.Pointer
}

// ProxyAddressEnumeratorClass is a wrapper around the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

// ProxyAddressEnumeratorPrivate is a wrapper around the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ProxyAddressPrivate is a wrapper around the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

// ProxyResolverInterface is a wrapper around the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for is_supported
	// no type for lookup
	// no type for lookup_async
	// no type for lookup_finish
}

// ResolverClass is a wrapper around the C record GResolverClass.
type ResolverClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for reload
	// no type for lookup_by_name
	// no type for lookup_by_name_async
	// no type for lookup_by_name_finish
	// no type for lookup_by_address
	// no type for lookup_by_address_async
	// no type for lookup_by_address_finish
	// no type for lookup_service
	// no type for lookup_service_async
	// no type for lookup_service_finish
	// no type for lookup_records
	// no type for lookup_records_async
	// no type for lookup_records_finish
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// ResolverPrivate is a wrapper around the C record GResolverPrivate.
type ResolverPrivate struct {
	native unsafe.Pointer
}

// SeekableIface is a wrapper around the C record GSeekableIface.
type SeekableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
}

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// SettingsClass is a wrapper around the C record GSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for writable_changed
	// no type for changed
	// no type for writable_change_event
	// no type for change_event
	// no type for padding
}

// SettingsPrivate is a wrapper around the C record GSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// SettingsSchemaKey is a wrapper around the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native unsafe.Pointer
}

// SimpleActionGroupClass is a wrapper around the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// Private : padding
}

// SimpleActionGroupPrivate is a wrapper around the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

// SimpleAsyncResultClass is a wrapper around the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

// SimpleProxyResolverClass is a wrapper around the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// SimpleProxyResolverPrivate is a wrapper around the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

// SocketAddressClass is a wrapper around the C record GSocketAddressClass.
type SocketAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_family
	// no type for get_native_size
	// no type for to_native
}

// SocketAddressEnumeratorClass is a wrapper around the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for next
	// no type for next_async
	// no type for next_finish
}

// SocketClass is a wrapper around the C record GSocketClass.
type SocketClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

// SocketClientClass is a wrapper around the C record GSocketClientClass.
type SocketClientClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for event
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

// SocketClientPrivate is a wrapper around the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native unsafe.Pointer
}

// SocketConnectableIface is a wrapper around the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for enumerate
	// no type for proxy_enumerate
	// no type for to_string
}

// SocketConnectionClass is a wrapper around the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// SocketConnectionPrivate is a wrapper around the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

// SocketControlMessageClass is a wrapper around the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_size
	// no type for get_level
	// no type for get_type
	// no type for serialize
	// no type for deserialize
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// SocketControlMessagePrivate is a wrapper around the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

// SocketListenerClass is a wrapper around the C record GSocketListenerClass.
type SocketListenerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for event
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// SocketListenerPrivate is a wrapper around the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native unsafe.Pointer
}

// SocketPrivate is a wrapper around the C record GSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// SocketServiceClass is a wrapper around the C record GSocketServiceClass.
type SocketServiceClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for incoming
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// SocketServicePrivate is a wrapper around the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native unsafe.Pointer
}

// SrvTarget is a wrapper around the C record GSrvTarget.
type SrvTarget struct {
	native unsafe.Pointer
}

// StaticResource is a wrapper around the C record GStaticResource.
type StaticResource struct {
	native unsafe.Pointer
	// Private : data
	// Private : data_len
	// Private : resource
	// Private : next
	// Private : padding
}

// TaskClass is a wrapper around the C record GTaskClass.
type TaskClass struct {
	native unsafe.Pointer
}

// TcpConnectionClass is a wrapper around the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// TcpConnectionPrivate is a wrapper around the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

// TcpWrapperConnectionClass is a wrapper around the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// TcpWrapperConnectionPrivate is a wrapper around the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

// ThemedIconClass is a wrapper around the C record GThemedIconClass.
type ThemedIconClass struct {
	native unsafe.Pointer
}

// ThreadedSocketServiceClass is a wrapper around the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for run
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// ThreadedSocketServicePrivate is a wrapper around the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

// TlsCertificateClass is a wrapper around the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for verify
	// Private : padding
}

// TlsCertificatePrivate is a wrapper around the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

// TlsConnectionClass is a wrapper around the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// Private : padding
}

// TlsConnectionPrivate is a wrapper around the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

// TlsDatabasePrivate is a wrapper around the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

// TlsFileDatabaseInterface is a wrapper around the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
	// g_iface : record
	// Private : padding
}

// TlsInteractionPrivate is a wrapper around the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

// TlsPasswordClass is a wrapper around the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_value
	// no type for set_value
	// no type for get_default_warning
	// Private : padding
}

// TlsPasswordPrivate is a wrapper around the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

// UnixConnectionClass is a wrapper around the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// UnixConnectionPrivate is a wrapper around the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

// UnixCredentialsMessagePrivate is a wrapper around the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

// UnixFDListClass is a wrapper around the C record GUnixFDListClass.
type UnixFDListClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// UnixFDListPrivate is a wrapper around the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native unsafe.Pointer
}

// UnixFDMessageClass is a wrapper around the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

// UnixFDMessagePrivate is a wrapper around the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

// UnixInputStreamClass is a wrapper around the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// UnixInputStreamPrivate is a wrapper around the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

// UnixMountEntry is a wrapper around the C record GUnixMountEntry.
type UnixMountEntry struct {
	native unsafe.Pointer
}

// UnixMountMonitorClass is a wrapper around the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

// UnixMountPoint is a wrapper around the C record GUnixMountPoint.
type UnixMountPoint struct {
	native unsafe.Pointer
}

// UnixOutputStreamClass is a wrapper around the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

// UnixOutputStreamPrivate is a wrapper around the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

// UnixSocketAddressClass is a wrapper around the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// UnixSocketAddressPrivate is a wrapper around the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

// VfsClass is a wrapper around the C record GVfsClass.
type VfsClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for is_active
	// no type for get_file_for_path
	// no type for get_file_for_uri
	// no type for get_supported_uri_schemes
	// no type for parse_name
	// no type for local_file_add_info
	// no type for add_writable_namespaces
	// no type for local_file_set_attributes
	// no type for local_file_removed
	// no type for local_file_moved
	// no type for deserialize_icon
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// VolumeIface is a wrapper around the C record GVolumeIface.
type VolumeIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for changed
	// no type for removed
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_drive
	// no type for get_mount
	// no type for can_mount
	// no type for can_eject
	// no type for mount_fn
	// no type for mount_finish
	// no type for eject
	// no type for eject_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for should_automount
	// no type for get_activation_root
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

// VolumeMonitorClass is a wrapper around the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for volume_added
	// no type for volume_removed
	// no type for volume_changed
	// no type for mount_added
	// no type for mount_removed
	// no type for mount_pre_unmount
	// no type for mount_changed
	// no type for drive_connected
	// no type for drive_disconnected
	// no type for drive_changed
	// no type for is_supported
	// no type for get_connected_drives
	// no type for get_volumes
	// no type for get_mounts
	// no type for get_volume_for_uuid
	// no type for get_mount_for_uuid
	// no type for adopt_orphan_mount
	// no type for drive_eject_button
	// no type for drive_stop_button
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

// ZlibCompressorClass is a wrapper around the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ZlibDecompressorClass is a wrapper around the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native unsafe.Pointer
	// parent_class : record
}
