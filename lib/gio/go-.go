// This is a generated file - DO NOT EDIT

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	call "github.com/pekim/gobbi/lib/internal/call"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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
// #include <stdlib.h>
import "C"

type AppInfoCreateFlags int

const (
	APP_INFO_CREATE_NONE                          AppInfoCreateFlags = 0
	APP_INFO_CREATE_NEEDS_TERMINAL                AppInfoCreateFlags = 1
	APP_INFO_CREATE_SUPPORTS_URIS                 AppInfoCreateFlags = 2
	APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION AppInfoCreateFlags = 4
)

type AskPasswordFlags int

const (
	ASK_PASSWORD_NEED_PASSWORD       AskPasswordFlags = 1
	ASK_PASSWORD_NEED_USERNAME       AskPasswordFlags = 2
	ASK_PASSWORD_NEED_DOMAIN         AskPasswordFlags = 4
	ASK_PASSWORD_SAVING_SUPPORTED    AskPasswordFlags = 8
	ASK_PASSWORD_ANONYMOUS_SUPPORTED AskPasswordFlags = 16
)

type FileAttributeInfoFlags int

const (
	FILE_ATTRIBUTE_INFO_NONE            FileAttributeInfoFlags = 0
	FILE_ATTRIBUTE_INFO_COPY_WITH_FILE  FileAttributeInfoFlags = 1
	FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED FileAttributeInfoFlags = 2
)

type FileCopyFlags int

const (
	FILE_COPY_NONE                 FileCopyFlags = 0
	FILE_COPY_OVERWRITE            FileCopyFlags = 1
	FILE_COPY_BACKUP               FileCopyFlags = 2
	FILE_COPY_NOFOLLOW_SYMLINKS    FileCopyFlags = 4
	FILE_COPY_ALL_METADATA         FileCopyFlags = 8
	FILE_COPY_NO_FALLBACK_FOR_MOVE FileCopyFlags = 16
	FILE_COPY_TARGET_DEFAULT_PERMS FileCopyFlags = 32
)

type FileCreateFlags int

const (
	FILE_CREATE_NONE                FileCreateFlags = 0
	FILE_CREATE_PRIVATE             FileCreateFlags = 1
	FILE_CREATE_REPLACE_DESTINATION FileCreateFlags = 2
)

type FileMonitorFlags int

const (
	FILE_MONITOR_NONE             FileMonitorFlags = 0
	FILE_MONITOR_WATCH_MOUNTS     FileMonitorFlags = 1
	FILE_MONITOR_SEND_MOVED       FileMonitorFlags = 2
	FILE_MONITOR_WATCH_HARD_LINKS FileMonitorFlags = 4
	FILE_MONITOR_WATCH_MOVES      FileMonitorFlags = 8
)

type FileQueryInfoFlags int

const (
	FILE_QUERY_INFO_NONE              FileQueryInfoFlags = 0
	FILE_QUERY_INFO_NOFOLLOW_SYMLINKS FileQueryInfoFlags = 1
)

type MountMountFlags int

const (
	MOUNT_MOUNT_NONE MountMountFlags = 0
)

type MountUnmountFlags int

const (
	MOUNT_UNMOUNT_NONE  MountUnmountFlags = 0
	MOUNT_UNMOUNT_FORCE MountUnmountFlags = 1
)

type OutputStreamSpliceFlags int

const (
	OUTPUT_STREAM_SPLICE_NONE         OutputStreamSpliceFlags = 0
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = 1
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET OutputStreamSpliceFlags = 2
)

type SettingsBindFlags int

const (
	SETTINGS_BIND_DEFAULT        SettingsBindFlags = 0
	SETTINGS_BIND_GET            SettingsBindFlags = 1
	SETTINGS_BIND_SET            SettingsBindFlags = 2
	SETTINGS_BIND_NO_SENSITIVITY SettingsBindFlags = 4
	SETTINGS_BIND_GET_NO_CHANGES SettingsBindFlags = 8
	SETTINGS_BIND_INVERT_BOOLEAN SettingsBindFlags = 16
)

// Unsupported : g_app_launch_context_new : return type :

// Unsupported : g_buffered_input_stream_new : return type :

// Unsupported : g_buffered_input_stream_new_sized : return type :

// Seekable returns the Seekable interface implemented by BufferedInputStream
func (recv *BufferedInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_buffered_output_stream_new : return type :

// Unsupported : g_buffered_output_stream_new_sized : return type :

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by BytesIcon
func (recv *BytesIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by BytesIcon
func (recv *BytesIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Unsupported : g_cancellable_new : return type :

// Converter returns the Converter interface implemented by CharsetConverter
func (recv *CharsetConverter) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Initable returns the Initable interface implemented by CharsetConverter
func (recv *CharsetConverter) Initable() *Initable {
	return InitableNewFromC(recv.ToC())
}

// Unsupported : g_converter_input_stream_new : return type :

// PollableInputStream returns the PollableInputStream interface implemented by ConverterInputStream
func (recv *ConverterInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Unsupported : g_converter_output_stream_new : return type :

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) ActionGroup() *ActionGroup {
	return ActionGroupNewFromC(recv.ToC())
}

// RemoteActionGroup returns the RemoteActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) RemoteActionGroup() *RemoteActionGroup {
	return RemoteActionGroupNewFromC(recv.ToC())
}

// Unsupported : g_data_input_stream_new : return type :

// Seekable returns the Seekable interface implemented by DataInputStream
func (recv *DataInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_data_output_stream_new : return type :

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_desktop_app_info_new : return type :

// Unsupported : g_desktop_app_info_new_from_filename : return type :

// AppInfo returns the AppInfo interface implemented by DesktopAppInfo
func (recv *DesktopAppInfo) AppInfo() *AppInfo {
	return AppInfoNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by Emblem
func (recv *Emblem) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by EmblemedIcon
func (recv *EmblemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_file_icon_new : return type :

// Icon returns the Icon interface implemented by FileIcon
func (recv *FileIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by FileIcon
func (recv *FileIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Unsupported : g_file_info_new : return type :

// Seekable returns the Seekable interface implemented by FileInputStream
func (recv *FileInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_filename_completer_new : return type :

// Unsupported : g_io_module_new : return type :

// TypePlugin returns the TypePlugin interface implemented by IOModule
func (recv *IOModule) TypePlugin() *gobject.TypePlugin {
	return gobject.TypePluginNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by InetSocketAddress
func (recv *InetSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// ListModel returns the ListModel interface implemented by ListStore
func (recv *ListStore) ListModel() *ListModel {
	return ListModelNewFromC(recv.ToC())
}

// Unsupported : g_memory_input_stream_new : return type :

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// PollableInputStream returns the PollableInputStream interface implemented by MemoryInputStream
func (recv *MemoryInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryInputStream
func (recv *MemoryInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_memory_output_stream_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PollableOutputStream returns the PollableOutputStream interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_mount_operation_new : return type :

// SocketConnectable returns the SocketConnectable interface implemented by NetworkAddress
func (recv *NetworkAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkService
func (recv *NetworkService) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// ProxyResolver returns the ProxyResolver interface implemented by SimpleProxyResolver
func (recv *SimpleProxyResolver) ProxyResolver() *ProxyResolver {
	return ProxyResolverNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by SocketAddress
func (recv *SocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// AsyncResult returns the AsyncResult interface implemented by Task
func (recv *Task) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// Unsupported : g_themed_icon_new : return type :

// Unsupported : g_themed_icon_new_from_names : return type :

// Unsupported : g_themed_icon_new_with_default_fallbacks : return type :

// Icon returns the Icon interface implemented by ThemedIcon
func (recv *ThemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Unsupported : g_unix_input_stream_new : return type :

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixInputStream
func (recv *UnixInputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableInputStream returns the PollableInputStream interface implemented by UnixInputStream
func (recv *UnixInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Unsupported : g_unix_mount_monitor_new : return type :

// Unsupported : g_unix_output_stream_new : return type :

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixOutputStream
func (recv *UnixOutputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableOutputStream returns the PollableOutputStream interface implemented by UnixOutputStream
func (recv *UnixOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Unsupported : g_unix_socket_address_new_abstract : return type :

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME string = "gio-desktop-app-info-lookup"
const FILE_ATTRIBUTE_ACCESS_CAN_DELETE string = "access::can-delete"
const FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE string = "access::can-execute"
const FILE_ATTRIBUTE_ACCESS_CAN_READ string = "access::can-read"
const FILE_ATTRIBUTE_ACCESS_CAN_RENAME string = "access::can-rename"
const FILE_ATTRIBUTE_ACCESS_CAN_TRASH string = "access::can-trash"
const FILE_ATTRIBUTE_ACCESS_CAN_WRITE string = "access::can-write"
const FILE_ATTRIBUTE_DOS_IS_ARCHIVE string = "dos::is-archive"
const FILE_ATTRIBUTE_DOS_IS_SYSTEM string = "dos::is-system"
const FILE_ATTRIBUTE_ETAG_VALUE string = "etag::value"
const FILE_ATTRIBUTE_FILESYSTEM_FREE string = "filesystem::free"
const FILE_ATTRIBUTE_FILESYSTEM_READONLY string = "filesystem::readonly"

// Blacklisted : FILE_ATTRIBUTE_FILESYSTEM_REMOTE

const FILE_ATTRIBUTE_FILESYSTEM_SIZE string = "filesystem::size"
const FILE_ATTRIBUTE_FILESYSTEM_TYPE string = "filesystem::type"
const FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW string = "filesystem::use-preview"
const FILE_ATTRIBUTE_GVFS_BACKEND string = "gvfs::backend"
const FILE_ATTRIBUTE_ID_FILE string = "id::file"
const FILE_ATTRIBUTE_ID_FILESYSTEM string = "id::filesystem"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT string = "mountable::can-eject"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT string = "mountable::can-mount"
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT string = "mountable::can-unmount"
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI string = "mountable::hal-udi"
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE string = "mountable::unix-device"
const FILE_ATTRIBUTE_OWNER_GROUP string = "owner::group"
const FILE_ATTRIBUTE_OWNER_USER string = "owner::user"
const FILE_ATTRIBUTE_OWNER_USER_REAL string = "owner::user-real"
const FILE_ATTRIBUTE_SELINUX_CONTEXT string = "selinux::context"
const FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE string = "standard::content-type"
const FILE_ATTRIBUTE_STANDARD_COPY_NAME string = "standard::copy-name"
const FILE_ATTRIBUTE_STANDARD_DESCRIPTION string = "standard::description"
const FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME string = "standard::display-name"
const FILE_ATTRIBUTE_STANDARD_EDIT_NAME string = "standard::edit-name"
const FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE string = "standard::fast-content-type"
const FILE_ATTRIBUTE_STANDARD_ICON string = "standard::icon"
const FILE_ATTRIBUTE_STANDARD_IS_BACKUP string = "standard::is-backup"
const FILE_ATTRIBUTE_STANDARD_IS_HIDDEN string = "standard::is-hidden"
const FILE_ATTRIBUTE_STANDARD_IS_SYMLINK string = "standard::is-symlink"
const FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL string = "standard::is-virtual"
const FILE_ATTRIBUTE_STANDARD_NAME string = "standard::name"
const FILE_ATTRIBUTE_STANDARD_SIZE string = "standard::size"
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER string = "standard::sort-order"
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET string = "standard::symlink-target"
const FILE_ATTRIBUTE_STANDARD_TARGET_URI string = "standard::target-uri"
const FILE_ATTRIBUTE_STANDARD_TYPE string = "standard::type"
const FILE_ATTRIBUTE_THUMBNAILING_FAILED string = "thumbnail::failed"
const FILE_ATTRIBUTE_THUMBNAIL_PATH string = "thumbnail::path"
const FILE_ATTRIBUTE_TIME_ACCESS string = "time::access"
const FILE_ATTRIBUTE_TIME_ACCESS_USEC string = "time::access-usec"
const FILE_ATTRIBUTE_TIME_CHANGED string = "time::changed"
const FILE_ATTRIBUTE_TIME_CHANGED_USEC string = "time::changed-usec"
const FILE_ATTRIBUTE_TIME_CREATED string = "time::created"
const FILE_ATTRIBUTE_TIME_CREATED_USEC string = "time::created-usec"
const FILE_ATTRIBUTE_TIME_MODIFIED string = "time::modified"
const FILE_ATTRIBUTE_TIME_MODIFIED_USEC string = "time::modified-usec"
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT string = "trash::item-count"
const FILE_ATTRIBUTE_UNIX_BLOCKS string = "unix::blocks"
const FILE_ATTRIBUTE_UNIX_BLOCK_SIZE string = "unix::block-size"
const FILE_ATTRIBUTE_UNIX_DEVICE string = "unix::device"
const FILE_ATTRIBUTE_UNIX_GID string = "unix::gid"
const FILE_ATTRIBUTE_UNIX_INODE string = "unix::inode"
const FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT string = "unix::is-mountpoint"
const FILE_ATTRIBUTE_UNIX_MODE string = "unix::mode"
const FILE_ATTRIBUTE_UNIX_NLINK string = "unix::nlink"
const FILE_ATTRIBUTE_UNIX_RDEV string = "unix::rdev"
const FILE_ATTRIBUTE_UNIX_UID string = "unix::uid"
const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME string = "gio-native-volume-monitor"
const PROXY_RESOLVER_EXTENSION_POINT_NAME string = "gio-proxy-resolver"

// Blacklisted : SETTINGS_BACKEND_EXTENSION_POINT_NAME

const TLS_BACKEND_EXTENSION_POINT_NAME string = "gio-tls-backend"
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT string = "1.3.6.1.5.5.7.3.2"
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER string = "1.3.6.1.5.5.7.3.1"
const VFS_EXTENSION_POINT_NAME string = "gio-vfs"
const VOLUME_IDENTIFIER_KIND_CLASS string = "class"
const VOLUME_IDENTIFIER_KIND_HAL_UDI string = "hal-udi"
const VOLUME_IDENTIFIER_KIND_LABEL string = "label"
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT string = "nfs-mount"
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE string = "unix-device"
const VOLUME_IDENTIFIER_KIND_UUID string = "uuid"
const VOLUME_MONITOR_EXTENSION_POINT_NAME string = "gio-volume-monitor"

type DataStreamByteOrder int

const (
	DATA_STREAM_BYTE_ORDER_BIG_ENDIAN    DataStreamByteOrder = 0
	DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN DataStreamByteOrder = 1
	DATA_STREAM_BYTE_ORDER_HOST_ENDIAN   DataStreamByteOrder = 2
)

type DataStreamNewlineType int

const (
	DATA_STREAM_NEWLINE_TYPE_LF    DataStreamNewlineType = 0
	DATA_STREAM_NEWLINE_TYPE_CR    DataStreamNewlineType = 1
	DATA_STREAM_NEWLINE_TYPE_CR_LF DataStreamNewlineType = 2
	DATA_STREAM_NEWLINE_TYPE_ANY   DataStreamNewlineType = 3
)

type FileAttributeStatus int

const (
	FILE_ATTRIBUTE_STATUS_UNSET         FileAttributeStatus = 0
	FILE_ATTRIBUTE_STATUS_SET           FileAttributeStatus = 1
	FILE_ATTRIBUTE_STATUS_ERROR_SETTING FileAttributeStatus = 2
)

type FileAttributeType int

const (
	FILE_ATTRIBUTE_TYPE_INVALID     FileAttributeType = 0
	FILE_ATTRIBUTE_TYPE_STRING      FileAttributeType = 1
	FILE_ATTRIBUTE_TYPE_BYTE_STRING FileAttributeType = 2
	FILE_ATTRIBUTE_TYPE_BOOLEAN     FileAttributeType = 3
	FILE_ATTRIBUTE_TYPE_UINT32      FileAttributeType = 4
	FILE_ATTRIBUTE_TYPE_INT32       FileAttributeType = 5
	FILE_ATTRIBUTE_TYPE_UINT64      FileAttributeType = 6
	FILE_ATTRIBUTE_TYPE_INT64       FileAttributeType = 7
	FILE_ATTRIBUTE_TYPE_OBJECT      FileAttributeType = 8
	FILE_ATTRIBUTE_TYPE_STRINGV     FileAttributeType = 9
)

type FileMonitorEvent int

const (
	FILE_MONITOR_EVENT_CHANGED           FileMonitorEvent = 0
	FILE_MONITOR_EVENT_CHANGES_DONE_HINT FileMonitorEvent = 1
	FILE_MONITOR_EVENT_DELETED           FileMonitorEvent = 2
	FILE_MONITOR_EVENT_CREATED           FileMonitorEvent = 3
	FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED FileMonitorEvent = 4
	FILE_MONITOR_EVENT_PRE_UNMOUNT       FileMonitorEvent = 5
	FILE_MONITOR_EVENT_UNMOUNTED         FileMonitorEvent = 6
	FILE_MONITOR_EVENT_MOVED             FileMonitorEvent = 7
	FILE_MONITOR_EVENT_RENAMED           FileMonitorEvent = 8
	FILE_MONITOR_EVENT_MOVED_IN          FileMonitorEvent = 9
	FILE_MONITOR_EVENT_MOVED_OUT         FileMonitorEvent = 10
)

type FileType int

const (
	FILE_TYPE_UNKNOWN       FileType = 0
	FILE_TYPE_REGULAR       FileType = 1
	FILE_TYPE_DIRECTORY     FileType = 2
	FILE_TYPE_SYMBOLIC_LINK FileType = 3
	FILE_TYPE_SPECIAL       FileType = 4
	FILE_TYPE_SHORTCUT      FileType = 5
	FILE_TYPE_MOUNTABLE     FileType = 6
)

type FilesystemPreviewType int

const (
	FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS FilesystemPreviewType = 0
	FILESYSTEM_PREVIEW_TYPE_IF_LOCAL  FilesystemPreviewType = 1
	FILESYSTEM_PREVIEW_TYPE_NEVER     FilesystemPreviewType = 2
)

type IOErrorEnum int

const (
	IO_ERROR_FAILED              IOErrorEnum = 0
	IO_ERROR_NOT_FOUND           IOErrorEnum = 1
	IO_ERROR_EXISTS              IOErrorEnum = 2
	IO_ERROR_IS_DIRECTORY        IOErrorEnum = 3
	IO_ERROR_NOT_DIRECTORY       IOErrorEnum = 4
	IO_ERROR_NOT_EMPTY           IOErrorEnum = 5
	IO_ERROR_NOT_REGULAR_FILE    IOErrorEnum = 6
	IO_ERROR_NOT_SYMBOLIC_LINK   IOErrorEnum = 7
	IO_ERROR_NOT_MOUNTABLE_FILE  IOErrorEnum = 8
	IO_ERROR_FILENAME_TOO_LONG   IOErrorEnum = 9
	IO_ERROR_INVALID_FILENAME    IOErrorEnum = 10
	IO_ERROR_TOO_MANY_LINKS      IOErrorEnum = 11
	IO_ERROR_NO_SPACE            IOErrorEnum = 12
	IO_ERROR_INVALID_ARGUMENT    IOErrorEnum = 13
	IO_ERROR_PERMISSION_DENIED   IOErrorEnum = 14
	IO_ERROR_NOT_SUPPORTED       IOErrorEnum = 15
	IO_ERROR_NOT_MOUNTED         IOErrorEnum = 16
	IO_ERROR_ALREADY_MOUNTED     IOErrorEnum = 17
	IO_ERROR_CLOSED              IOErrorEnum = 18
	IO_ERROR_CANCELLED           IOErrorEnum = 19
	IO_ERROR_PENDING             IOErrorEnum = 20
	IO_ERROR_READ_ONLY           IOErrorEnum = 21
	IO_ERROR_CANT_CREATE_BACKUP  IOErrorEnum = 22
	IO_ERROR_WRONG_ETAG          IOErrorEnum = 23
	IO_ERROR_TIMED_OUT           IOErrorEnum = 24
	IO_ERROR_WOULD_RECURSE       IOErrorEnum = 25
	IO_ERROR_BUSY                IOErrorEnum = 26
	IO_ERROR_WOULD_BLOCK         IOErrorEnum = 27
	IO_ERROR_HOST_NOT_FOUND      IOErrorEnum = 28
	IO_ERROR_WOULD_MERGE         IOErrorEnum = 29
	IO_ERROR_FAILED_HANDLED      IOErrorEnum = 30
	IO_ERROR_TOO_MANY_OPEN_FILES IOErrorEnum = 31
	IO_ERROR_NOT_INITIALIZED     IOErrorEnum = 32
	IO_ERROR_ADDRESS_IN_USE      IOErrorEnum = 33
	IO_ERROR_PARTIAL_INPUT       IOErrorEnum = 34
	IO_ERROR_INVALID_DATA        IOErrorEnum = 35
	IO_ERROR_DBUS_ERROR          IOErrorEnum = 36
	IO_ERROR_HOST_UNREACHABLE    IOErrorEnum = 37
	IO_ERROR_NETWORK_UNREACHABLE IOErrorEnum = 38
	IO_ERROR_CONNECTION_REFUSED  IOErrorEnum = 39
	IO_ERROR_PROXY_FAILED        IOErrorEnum = 40
	IO_ERROR_PROXY_AUTH_FAILED   IOErrorEnum = 41
	IO_ERROR_PROXY_NEED_AUTH     IOErrorEnum = 42
	IO_ERROR_PROXY_NOT_ALLOWED   IOErrorEnum = 43
	IO_ERROR_BROKEN_PIPE         IOErrorEnum = 44
	IO_ERROR_CONNECTION_CLOSED   IOErrorEnum = 44
	IO_ERROR_NOT_CONNECTED       IOErrorEnum = 45
	IO_ERROR_MESSAGE_TOO_LARGE   IOErrorEnum = 46
)

type MountOperationResult int

const (
	MOUNT_OPERATION_HANDLED   MountOperationResult = 0
	MOUNT_OPERATION_ABORTED   MountOperationResult = 1
	MOUNT_OPERATION_UNHANDLED MountOperationResult = 2
)

type PasswordSave int

const (
	PASSWORD_SAVE_NEVER       PasswordSave = 0
	PASSWORD_SAVE_FOR_SESSION PasswordSave = 1
	PASSWORD_SAVE_PERMANENTLY PasswordSave = 2
)

// Unsupported : g_app_info_create_from_commandline : return type :

// Unsupported : g_app_info_get_all : return type :

// Unsupported : g_app_info_get_all_for_type : return type :

// Unsupported : g_app_info_get_default_for_type : return type :

// Unsupported : g_app_info_get_default_for_uri_scheme : return type :

// Unsupported : g_app_info_launch_default_for_uri : return type :

// Unsupported : g_content_type_can_be_executable : return type :

// Unsupported : g_content_type_equals : return type :

// Unsupported : g_content_type_get_description : return type :

// Unsupported : g_content_type_get_icon : return type :

// Unsupported : g_content_type_get_mime_type : return type :

// Unsupported : g_content_type_guess : return type :

// Unsupported : g_content_type_is_a : return type :

// Unsupported : g_content_type_is_unknown : return type :

// Unsupported : g_content_types_get_registered : return type :

// Unsupported : g_dbus_error_quark : return type :

// Unsupported : g_file_new_for_commandline_arg : return type :

// Unsupported : g_file_new_for_path : return type :

// Unsupported : g_file_new_for_uri : return type :

// Unsupported : g_file_parse_name : return type :

// Unsupported : g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon

// Unsupported : g_io_error_from_errno : return type :

// Unsupported : g_io_error_quark : return type :

// Unsupported : g_io_extension_point_implement : return type :

// Unsupported : g_io_extension_point_lookup : return type :

// Unsupported : g_io_extension_point_register : return type :

// Unsupported : g_io_modules_load_all_in_directory : return type :

// IoSchedulerCancelAllJobs is a wrapper around the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(1853, &data)
	return
}

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Unsupported : g_keyfile_settings_backend_new : return type :

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_unix_is_mount_path_system_internal : return type :

// Unsupported : g_unix_mount_at : return type :

// Unsupported : g_unix_mount_get_device_path : return type :

// Unsupported : g_unix_mount_get_fs_type : return type :

// Unsupported : g_unix_mount_get_mount_path : return type :

// Unsupported : g_unix_mount_guess_can_eject : return type :

// Unsupported : g_unix_mount_guess_icon : return type :

// Unsupported : g_unix_mount_guess_name : return type :

// Unsupported : g_unix_mount_guess_should_display : return type :

// Unsupported : g_unix_mount_is_readonly : return type :

// Unsupported : g_unix_mount_is_system_internal : return type :

// Unsupported : g_unix_mount_points_changed_since : return type :

// Unsupported : g_unix_mount_points_get : return type :

// Unsupported : g_unix_mounts_changed_since : return type :

// Unsupported : g_unix_mounts_get : return type :

// Unsupported : g_file_attribute_info_list_new : return type :

// Unsupported : g_file_attribute_matcher_new : return type :

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate
