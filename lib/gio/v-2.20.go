// Code generated - DO NOT EDIT.
// +build gio_2.20

package gio

import (
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME is a representation of the C constant G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME.
const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME = "gio-desktop-app-info-lookup"

// FILE_ATTRIBUTE_ACCESS_CAN_DELETE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE.
const FILE_ATTRIBUTE_ACCESS_CAN_DELETE = "access::can-delete"

// FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE.
const FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE = "access::can-execute"

// FILE_ATTRIBUTE_ACCESS_CAN_READ is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_READ.
const FILE_ATTRIBUTE_ACCESS_CAN_READ = "access::can-read"

// FILE_ATTRIBUTE_ACCESS_CAN_RENAME is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME.
const FILE_ATTRIBUTE_ACCESS_CAN_RENAME = "access::can-rename"

// FILE_ATTRIBUTE_ACCESS_CAN_TRASH is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH.
const FILE_ATTRIBUTE_ACCESS_CAN_TRASH = "access::can-trash"

// FILE_ATTRIBUTE_ACCESS_CAN_WRITE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE.
const FILE_ATTRIBUTE_ACCESS_CAN_WRITE = "access::can-write"

// FILE_ATTRIBUTE_DOS_IS_ARCHIVE is a representation of the C constant G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE.
const FILE_ATTRIBUTE_DOS_IS_ARCHIVE = "dos::is-archive"

// FILE_ATTRIBUTE_DOS_IS_SYSTEM is a representation of the C constant G_FILE_ATTRIBUTE_DOS_IS_SYSTEM.
const FILE_ATTRIBUTE_DOS_IS_SYSTEM = "dos::is-system"

// FILE_ATTRIBUTE_ETAG_VALUE is a representation of the C constant G_FILE_ATTRIBUTE_ETAG_VALUE.
const FILE_ATTRIBUTE_ETAG_VALUE = "etag::value"

// FILE_ATTRIBUTE_FILESYSTEM_FREE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_FREE.
const FILE_ATTRIBUTE_FILESYSTEM_FREE = "filesystem::free"

// FILE_ATTRIBUTE_FILESYSTEM_READONLY is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_READONLY.
const FILE_ATTRIBUTE_FILESYSTEM_READONLY = "filesystem::readonly"

// FILE_ATTRIBUTE_FILESYSTEM_REMOTE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_REMOTE.
const FILE_ATTRIBUTE_FILESYSTEM_REMOTE = "filesystem::remote"

// FILE_ATTRIBUTE_FILESYSTEM_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_SIZE.
const FILE_ATTRIBUTE_FILESYSTEM_SIZE = "filesystem::size"

// FILE_ATTRIBUTE_FILESYSTEM_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_TYPE.
const FILE_ATTRIBUTE_FILESYSTEM_TYPE = "filesystem::type"

// FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.
const FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW = "filesystem::use-preview"

// FILE_ATTRIBUTE_GVFS_BACKEND is a representation of the C constant G_FILE_ATTRIBUTE_GVFS_BACKEND.
const FILE_ATTRIBUTE_GVFS_BACKEND = "gvfs::backend"

// FILE_ATTRIBUTE_ID_FILE is a representation of the C constant G_FILE_ATTRIBUTE_ID_FILE.
const FILE_ATTRIBUTE_ID_FILE = "id::file"

// FILE_ATTRIBUTE_ID_FILESYSTEM is a representation of the C constant G_FILE_ATTRIBUTE_ID_FILESYSTEM.
const FILE_ATTRIBUTE_ID_FILESYSTEM = "id::filesystem"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT = "mountable::can-eject"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT = "mountable::can-mount"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT = "mountable::can-unmount"

// FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI.
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI = "mountable::hal-udi"

// FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE.
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE = "mountable::unix-device"

// FILE_ATTRIBUTE_OWNER_GROUP is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_GROUP.
const FILE_ATTRIBUTE_OWNER_GROUP = "owner::group"

// FILE_ATTRIBUTE_OWNER_USER is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_USER.
const FILE_ATTRIBUTE_OWNER_USER = "owner::user"

// FILE_ATTRIBUTE_OWNER_USER_REAL is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_USER_REAL.
const FILE_ATTRIBUTE_OWNER_USER_REAL = "owner::user-real"

// FILE_ATTRIBUTE_PREVIEW_ICON is a representation of the C constant G_FILE_ATTRIBUTE_PREVIEW_ICON.
//
// since 2.20
const FILE_ATTRIBUTE_PREVIEW_ICON = "preview::icon"

// FILE_ATTRIBUTE_SELINUX_CONTEXT is a representation of the C constant G_FILE_ATTRIBUTE_SELINUX_CONTEXT.
const FILE_ATTRIBUTE_SELINUX_CONTEXT = "selinux::context"

// FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE.
//
// since 2.20
const FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE = "standard::allocated-size"

// FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
const FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE = "standard::content-type"

// FILE_ATTRIBUTE_STANDARD_COPY_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_COPY_NAME.
const FILE_ATTRIBUTE_STANDARD_COPY_NAME = "standard::copy-name"

// FILE_ATTRIBUTE_STANDARD_DESCRIPTION is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION.
const FILE_ATTRIBUTE_STANDARD_DESCRIPTION = "standard::description"

// FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
const FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME = "standard::display-name"

// FILE_ATTRIBUTE_STANDARD_EDIT_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
const FILE_ATTRIBUTE_STANDARD_EDIT_NAME = "standard::edit-name"

// FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE.
const FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE = "standard::fast-content-type"

// FILE_ATTRIBUTE_STANDARD_ICON is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_ICON.
const FILE_ATTRIBUTE_STANDARD_ICON = "standard::icon"

// FILE_ATTRIBUTE_STANDARD_IS_BACKUP is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP.
const FILE_ATTRIBUTE_STANDARD_IS_BACKUP = "standard::is-backup"

// FILE_ATTRIBUTE_STANDARD_IS_HIDDEN is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
const FILE_ATTRIBUTE_STANDARD_IS_HIDDEN = "standard::is-hidden"

// FILE_ATTRIBUTE_STANDARD_IS_SYMLINK is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
const FILE_ATTRIBUTE_STANDARD_IS_SYMLINK = "standard::is-symlink"

// FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL.
const FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL = "standard::is-virtual"

// FILE_ATTRIBUTE_STANDARD_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_NAME.
const FILE_ATTRIBUTE_STANDARD_NAME = "standard::name"

// FILE_ATTRIBUTE_STANDARD_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SIZE.
const FILE_ATTRIBUTE_STANDARD_SIZE = "standard::size"

// FILE_ATTRIBUTE_STANDARD_SORT_ORDER is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER = "standard::sort-order"

// FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET.
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET = "standard::symlink-target"

// FILE_ATTRIBUTE_STANDARD_TARGET_URI is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TARGET_URI.
const FILE_ATTRIBUTE_STANDARD_TARGET_URI = "standard::target-uri"

// FILE_ATTRIBUTE_STANDARD_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TYPE.
const FILE_ATTRIBUTE_STANDARD_TYPE = "standard::type"

// FILE_ATTRIBUTE_THUMBNAILING_FAILED is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAILING_FAILED.
const FILE_ATTRIBUTE_THUMBNAILING_FAILED = "thumbnail::failed"

// FILE_ATTRIBUTE_THUMBNAIL_PATH is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAIL_PATH.
const FILE_ATTRIBUTE_THUMBNAIL_PATH = "thumbnail::path"

// FILE_ATTRIBUTE_TIME_ACCESS is a representation of the C constant G_FILE_ATTRIBUTE_TIME_ACCESS.
const FILE_ATTRIBUTE_TIME_ACCESS = "time::access"

// FILE_ATTRIBUTE_TIME_ACCESS_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_ACCESS_USEC.
const FILE_ATTRIBUTE_TIME_ACCESS_USEC = "time::access-usec"

// FILE_ATTRIBUTE_TIME_CHANGED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CHANGED.
const FILE_ATTRIBUTE_TIME_CHANGED = "time::changed"

// FILE_ATTRIBUTE_TIME_CHANGED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CHANGED_USEC.
const FILE_ATTRIBUTE_TIME_CHANGED_USEC = "time::changed-usec"

// FILE_ATTRIBUTE_TIME_CREATED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CREATED.
const FILE_ATTRIBUTE_TIME_CREATED = "time::created"

// FILE_ATTRIBUTE_TIME_CREATED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CREATED_USEC.
const FILE_ATTRIBUTE_TIME_CREATED_USEC = "time::created-usec"

// FILE_ATTRIBUTE_TIME_MODIFIED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_MODIFIED.
const FILE_ATTRIBUTE_TIME_MODIFIED = "time::modified"

// FILE_ATTRIBUTE_TIME_MODIFIED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC.
const FILE_ATTRIBUTE_TIME_MODIFIED_USEC = "time::modified-usec"

// FILE_ATTRIBUTE_TRASH_ITEM_COUNT is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT.
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT = "trash::item-count"

// FILE_ATTRIBUTE_UNIX_BLOCKS is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_BLOCKS.
const FILE_ATTRIBUTE_UNIX_BLOCKS = "unix::blocks"

// FILE_ATTRIBUTE_UNIX_BLOCK_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE.
const FILE_ATTRIBUTE_UNIX_BLOCK_SIZE = "unix::block-size"

// FILE_ATTRIBUTE_UNIX_DEVICE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_DEVICE.
const FILE_ATTRIBUTE_UNIX_DEVICE = "unix::device"

// FILE_ATTRIBUTE_UNIX_GID is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_GID.
const FILE_ATTRIBUTE_UNIX_GID = "unix::gid"

// FILE_ATTRIBUTE_UNIX_INODE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_INODE.
const FILE_ATTRIBUTE_UNIX_INODE = "unix::inode"

// FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT.
const FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT = "unix::is-mountpoint"

// FILE_ATTRIBUTE_UNIX_MODE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_MODE.
const FILE_ATTRIBUTE_UNIX_MODE = "unix::mode"

// FILE_ATTRIBUTE_UNIX_NLINK is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_NLINK.
const FILE_ATTRIBUTE_UNIX_NLINK = "unix::nlink"

// FILE_ATTRIBUTE_UNIX_RDEV is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_RDEV.
const FILE_ATTRIBUTE_UNIX_RDEV = "unix::rdev"

// FILE_ATTRIBUTE_UNIX_UID is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_UID.
const FILE_ATTRIBUTE_UNIX_UID = "unix::uid"

// NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME.
const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-native-volume-monitor"

// PROXY_RESOLVER_EXTENSION_POINT_NAME is a representation of the C constant G_PROXY_RESOLVER_EXTENSION_POINT_NAME.
const PROXY_RESOLVER_EXTENSION_POINT_NAME = "gio-proxy-resolver"

// SETTINGS_BACKEND_EXTENSION_POINT_NAME is a representation of the C constant G_SETTINGS_BACKEND_EXTENSION_POINT_NAME.
const SETTINGS_BACKEND_EXTENSION_POINT_NAME = "gsettings-backend"

// TLS_BACKEND_EXTENSION_POINT_NAME is a representation of the C constant G_TLS_BACKEND_EXTENSION_POINT_NAME.
const TLS_BACKEND_EXTENSION_POINT_NAME = "gio-tls-backend"

// TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT is a representation of the C constant G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT = "1.3.6.1.5.5.7.3.2"

// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER is a representation of the C constant G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER = "1.3.6.1.5.5.7.3.1"

// VFS_EXTENSION_POINT_NAME is a representation of the C constant G_VFS_EXTENSION_POINT_NAME.
const VFS_EXTENSION_POINT_NAME = "gio-vfs"

// VOLUME_IDENTIFIER_KIND_CLASS is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_CLASS.
const VOLUME_IDENTIFIER_KIND_CLASS = "class"

// VOLUME_IDENTIFIER_KIND_HAL_UDI is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_HAL_UDI.
const VOLUME_IDENTIFIER_KIND_HAL_UDI = "hal-udi"

// VOLUME_IDENTIFIER_KIND_LABEL is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_LABEL.
const VOLUME_IDENTIFIER_KIND_LABEL = "label"

// VOLUME_IDENTIFIER_KIND_NFS_MOUNT is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT.
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT = "nfs-mount"

// VOLUME_IDENTIFIER_KIND_UNIX_DEVICE is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE.
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE = "unix-device"

// VOLUME_IDENTIFIER_KIND_UUID is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_UUID.
const VOLUME_IDENTIFIER_KIND_UUID = "uuid"

// VOLUME_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_VOLUME_MONITOR_EXTENSION_POINT_NAME.
const VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-volume-monitor"

// AppInfoCreateFlags is a representation of the C bitfield GAppInfoCreateFlags.
type AppInfoCreateFlags int

// AppInfoCreateFlags_none is a representation of the C bitfield member G_APP_INFO_CREATE_NONE.
const AppInfoCreateFlags_none = AppInfoCreateFlags(0)

// AppInfoCreateFlags_needs_terminal is a representation of the C bitfield member G_APP_INFO_CREATE_NEEDS_TERMINAL.
const AppInfoCreateFlags_needs_terminal = AppInfoCreateFlags(1)

// AppInfoCreateFlags_supports_uris is a representation of the C bitfield member G_APP_INFO_CREATE_SUPPORTS_URIS.
const AppInfoCreateFlags_supports_uris = AppInfoCreateFlags(2)

// AppInfoCreateFlags_supports_startup_notification is a representation of the C bitfield member G_APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION.
const AppInfoCreateFlags_supports_startup_notification = AppInfoCreateFlags(4)

// AskPasswordFlags is a representation of the C bitfield GAskPasswordFlags.
type AskPasswordFlags int

// AskPasswordFlags_need_password is a representation of the C bitfield member G_ASK_PASSWORD_NEED_PASSWORD.
const AskPasswordFlags_need_password = AskPasswordFlags(1)

// AskPasswordFlags_need_username is a representation of the C bitfield member G_ASK_PASSWORD_NEED_USERNAME.
const AskPasswordFlags_need_username = AskPasswordFlags(2)

// AskPasswordFlags_need_domain is a representation of the C bitfield member G_ASK_PASSWORD_NEED_DOMAIN.
const AskPasswordFlags_need_domain = AskPasswordFlags(4)

// AskPasswordFlags_saving_supported is a representation of the C bitfield member G_ASK_PASSWORD_SAVING_SUPPORTED.
const AskPasswordFlags_saving_supported = AskPasswordFlags(8)

// AskPasswordFlags_anonymous_supported is a representation of the C bitfield member G_ASK_PASSWORD_ANONYMOUS_SUPPORTED.
const AskPasswordFlags_anonymous_supported = AskPasswordFlags(16)

// AskPasswordFlags_tcrypt is a representation of the C bitfield member G_ASK_PASSWORD_TCRYPT.
const AskPasswordFlags_tcrypt = AskPasswordFlags(32)

// FileAttributeInfoFlags is a representation of the C bitfield GFileAttributeInfoFlags.
type FileAttributeInfoFlags int

// FileAttributeInfoFlags_none is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_NONE.
const FileAttributeInfoFlags_none = FileAttributeInfoFlags(0)

// FileAttributeInfoFlags_copy_with_file is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_COPY_WITH_FILE.
const FileAttributeInfoFlags_copy_with_file = FileAttributeInfoFlags(1)

// FileAttributeInfoFlags_copy_when_moved is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED.
const FileAttributeInfoFlags_copy_when_moved = FileAttributeInfoFlags(2)

// FileCopyFlags is a representation of the C bitfield GFileCopyFlags.
type FileCopyFlags int

// FileCopyFlags_none is a representation of the C bitfield member G_FILE_COPY_NONE.
const FileCopyFlags_none = FileCopyFlags(0)

// FileCopyFlags_overwrite is a representation of the C bitfield member G_FILE_COPY_OVERWRITE.
const FileCopyFlags_overwrite = FileCopyFlags(1)

// FileCopyFlags_backup is a representation of the C bitfield member G_FILE_COPY_BACKUP.
const FileCopyFlags_backup = FileCopyFlags(2)

// FileCopyFlags_nofollow_symlinks is a representation of the C bitfield member G_FILE_COPY_NOFOLLOW_SYMLINKS.
const FileCopyFlags_nofollow_symlinks = FileCopyFlags(4)

// FileCopyFlags_all_metadata is a representation of the C bitfield member G_FILE_COPY_ALL_METADATA.
const FileCopyFlags_all_metadata = FileCopyFlags(8)

// FileCopyFlags_no_fallback_for_move is a representation of the C bitfield member G_FILE_COPY_NO_FALLBACK_FOR_MOVE.
const FileCopyFlags_no_fallback_for_move = FileCopyFlags(16)

// FileCopyFlags_target_default_perms is a representation of the C bitfield member G_FILE_COPY_TARGET_DEFAULT_PERMS.
const FileCopyFlags_target_default_perms = FileCopyFlags(32)

// FileCreateFlags is a representation of the C bitfield GFileCreateFlags.
type FileCreateFlags int

// FileCreateFlags_none is a representation of the C bitfield member G_FILE_CREATE_NONE.
const FileCreateFlags_none = FileCreateFlags(0)

// FileCreateFlags_private is a representation of the C bitfield member G_FILE_CREATE_PRIVATE.
const FileCreateFlags_private = FileCreateFlags(1)

// FileCreateFlags_replace_destination is a representation of the C bitfield member G_FILE_CREATE_REPLACE_DESTINATION.
const FileCreateFlags_replace_destination = FileCreateFlags(2)

// FileMonitorFlags is a representation of the C bitfield GFileMonitorFlags.
type FileMonitorFlags int

// FileMonitorFlags_none is a representation of the C bitfield member G_FILE_MONITOR_NONE.
const FileMonitorFlags_none = FileMonitorFlags(0)

// FileMonitorFlags_watch_mounts is a representation of the C bitfield member G_FILE_MONITOR_WATCH_MOUNTS.
const FileMonitorFlags_watch_mounts = FileMonitorFlags(1)

// FileMonitorFlags_send_moved is a representation of the C bitfield member G_FILE_MONITOR_SEND_MOVED.
const FileMonitorFlags_send_moved = FileMonitorFlags(2)

// FileMonitorFlags_watch_hard_links is a representation of the C bitfield member G_FILE_MONITOR_WATCH_HARD_LINKS.
const FileMonitorFlags_watch_hard_links = FileMonitorFlags(4)

// FileMonitorFlags_watch_moves is a representation of the C bitfield member G_FILE_MONITOR_WATCH_MOVES.
const FileMonitorFlags_watch_moves = FileMonitorFlags(8)

// FileQueryInfoFlags is a representation of the C bitfield GFileQueryInfoFlags.
type FileQueryInfoFlags int

// FileQueryInfoFlags_none is a representation of the C bitfield member G_FILE_QUERY_INFO_NONE.
const FileQueryInfoFlags_none = FileQueryInfoFlags(0)

// FileQueryInfoFlags_nofollow_symlinks is a representation of the C bitfield member G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS.
const FileQueryInfoFlags_nofollow_symlinks = FileQueryInfoFlags(1)

// MountMountFlags is a representation of the C bitfield GMountMountFlags.
type MountMountFlags int

// MountMountFlags_none is a representation of the C bitfield member G_MOUNT_MOUNT_NONE.
const MountMountFlags_none = MountMountFlags(0)

// MountUnmountFlags is a representation of the C bitfield GMountUnmountFlags.
type MountUnmountFlags int

// MountUnmountFlags_none is a representation of the C bitfield member G_MOUNT_UNMOUNT_NONE.
const MountUnmountFlags_none = MountUnmountFlags(0)

// MountUnmountFlags_force is a representation of the C bitfield member G_MOUNT_UNMOUNT_FORCE.
const MountUnmountFlags_force = MountUnmountFlags(1)

// OutputStreamSpliceFlags is a representation of the C bitfield GOutputStreamSpliceFlags.
type OutputStreamSpliceFlags int

// OutputStreamSpliceFlags_none is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_NONE.
const OutputStreamSpliceFlags_none = OutputStreamSpliceFlags(0)

// OutputStreamSpliceFlags_close_source is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE.
const OutputStreamSpliceFlags_close_source = OutputStreamSpliceFlags(1)

// OutputStreamSpliceFlags_close_target is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET.
const OutputStreamSpliceFlags_close_target = OutputStreamSpliceFlags(2)

// SettingsBindFlags is a representation of the C bitfield GSettingsBindFlags.
type SettingsBindFlags int

// SettingsBindFlags_default is a representation of the C bitfield member G_SETTINGS_BIND_DEFAULT.
const SettingsBindFlags_default = SettingsBindFlags(0)

// SettingsBindFlags_get is a representation of the C bitfield member G_SETTINGS_BIND_GET.
const SettingsBindFlags_get = SettingsBindFlags(1)

// SettingsBindFlags_set is a representation of the C bitfield member G_SETTINGS_BIND_SET.
const SettingsBindFlags_set = SettingsBindFlags(2)

// SettingsBindFlags_no_sensitivity is a representation of the C bitfield member G_SETTINGS_BIND_NO_SENSITIVITY.
const SettingsBindFlags_no_sensitivity = SettingsBindFlags(4)

// SettingsBindFlags_get_no_changes is a representation of the C bitfield member G_SETTINGS_BIND_GET_NO_CHANGES.
const SettingsBindFlags_get_no_changes = SettingsBindFlags(8)

// SettingsBindFlags_invert_boolean is a representation of the C bitfield member G_SETTINGS_BIND_INVERT_BOOLEAN.
const SettingsBindFlags_invert_boolean = SettingsBindFlags(16)

// DataStreamByteOrder is a representation of the C enumeration GDataStreamByteOrder.
type DataStreamByteOrder int

// DataStreamByteOrder_big_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN.
const DataStreamByteOrder_big_endian = DataStreamByteOrder(0)

// DataStreamByteOrder_little_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN.
const DataStreamByteOrder_little_endian = DataStreamByteOrder(1)

// DataStreamByteOrder_host_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN.
const DataStreamByteOrder_host_endian = DataStreamByteOrder(2)

// DataStreamNewlineType is a representation of the C enumeration GDataStreamNewlineType.
type DataStreamNewlineType int

// DataStreamNewlineType_lf is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_LF.
const DataStreamNewlineType_lf = DataStreamNewlineType(0)

// DataStreamNewlineType_cr is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_CR.
const DataStreamNewlineType_cr = DataStreamNewlineType(1)

// DataStreamNewlineType_cr_lf is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_CR_LF.
const DataStreamNewlineType_cr_lf = DataStreamNewlineType(2)

// DataStreamNewlineType_any is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_ANY.
const DataStreamNewlineType_any = DataStreamNewlineType(3)

// EmblemOrigin is a representation of the C enumeration GEmblemOrigin.
type EmblemOrigin int

// EmblemOrigin_unknown is a representation of the C enumeration member G_EMBLEM_ORIGIN_UNKNOWN.
const EmblemOrigin_unknown = EmblemOrigin(0)

// EmblemOrigin_device is a representation of the C enumeration member G_EMBLEM_ORIGIN_DEVICE.
const EmblemOrigin_device = EmblemOrigin(1)

// EmblemOrigin_livemetadata is a representation of the C enumeration member G_EMBLEM_ORIGIN_LIVEMETADATA.
const EmblemOrigin_livemetadata = EmblemOrigin(2)

// EmblemOrigin_tag is a representation of the C enumeration member G_EMBLEM_ORIGIN_TAG.
const EmblemOrigin_tag = EmblemOrigin(3)

// FileAttributeStatus is a representation of the C enumeration GFileAttributeStatus.
type FileAttributeStatus int

// FileAttributeStatus_unset is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_UNSET.
const FileAttributeStatus_unset = FileAttributeStatus(0)

// FileAttributeStatus_set is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_SET.
const FileAttributeStatus_set = FileAttributeStatus(1)

// FileAttributeStatus_error_setting is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING.
const FileAttributeStatus_error_setting = FileAttributeStatus(2)

// FileAttributeType is a representation of the C enumeration GFileAttributeType.
type FileAttributeType int

// FileAttributeType_invalid is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INVALID.
const FileAttributeType_invalid = FileAttributeType(0)

// FileAttributeType_string is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_STRING.
const FileAttributeType_string = FileAttributeType(1)

// FileAttributeType_byte_string is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
const FileAttributeType_byte_string = FileAttributeType(2)

// FileAttributeType_boolean is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
const FileAttributeType_boolean = FileAttributeType(3)

// FileAttributeType_uint32 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_UINT32.
const FileAttributeType_uint32 = FileAttributeType(4)

// FileAttributeType_int32 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INT32.
const FileAttributeType_int32 = FileAttributeType(5)

// FileAttributeType_uint64 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_UINT64.
const FileAttributeType_uint64 = FileAttributeType(6)

// FileAttributeType_int64 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INT64.
const FileAttributeType_int64 = FileAttributeType(7)

// FileAttributeType_object is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_OBJECT.
const FileAttributeType_object = FileAttributeType(8)

// FileAttributeType_stringv is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_STRINGV.
const FileAttributeType_stringv = FileAttributeType(9)

// FileMonitorEvent is a representation of the C enumeration GFileMonitorEvent.
type FileMonitorEvent int

// FileMonitorEvent_changed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CHANGED.
const FileMonitorEvent_changed = FileMonitorEvent(0)

// FileMonitorEvent_changes_done_hint is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT.
const FileMonitorEvent_changes_done_hint = FileMonitorEvent(1)

// FileMonitorEvent_deleted is a representation of the C enumeration member G_FILE_MONITOR_EVENT_DELETED.
const FileMonitorEvent_deleted = FileMonitorEvent(2)

// FileMonitorEvent_created is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CREATED.
const FileMonitorEvent_created = FileMonitorEvent(3)

// FileMonitorEvent_attribute_changed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED.
const FileMonitorEvent_attribute_changed = FileMonitorEvent(4)

// FileMonitorEvent_pre_unmount is a representation of the C enumeration member G_FILE_MONITOR_EVENT_PRE_UNMOUNT.
const FileMonitorEvent_pre_unmount = FileMonitorEvent(5)

// FileMonitorEvent_unmounted is a representation of the C enumeration member G_FILE_MONITOR_EVENT_UNMOUNTED.
const FileMonitorEvent_unmounted = FileMonitorEvent(6)

// FileMonitorEvent_moved is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED.
const FileMonitorEvent_moved = FileMonitorEvent(7)

// FileMonitorEvent_renamed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_RENAMED.
const FileMonitorEvent_renamed = FileMonitorEvent(8)

// FileMonitorEvent_moved_in is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED_IN.
const FileMonitorEvent_moved_in = FileMonitorEvent(9)

// FileMonitorEvent_moved_out is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED_OUT.
const FileMonitorEvent_moved_out = FileMonitorEvent(10)

// FileType is a representation of the C enumeration GFileType.
type FileType int

// FileType_unknown is a representation of the C enumeration member G_FILE_TYPE_UNKNOWN.
const FileType_unknown = FileType(0)

// FileType_regular is a representation of the C enumeration member G_FILE_TYPE_REGULAR.
const FileType_regular = FileType(1)

// FileType_directory is a representation of the C enumeration member G_FILE_TYPE_DIRECTORY.
const FileType_directory = FileType(2)

// FileType_symbolic_link is a representation of the C enumeration member G_FILE_TYPE_SYMBOLIC_LINK.
const FileType_symbolic_link = FileType(3)

// FileType_special is a representation of the C enumeration member G_FILE_TYPE_SPECIAL.
const FileType_special = FileType(4)

// FileType_shortcut is a representation of the C enumeration member G_FILE_TYPE_SHORTCUT.
const FileType_shortcut = FileType(5)

// FileType_mountable is a representation of the C enumeration member G_FILE_TYPE_MOUNTABLE.
const FileType_mountable = FileType(6)

// FilesystemPreviewType is a representation of the C enumeration GFilesystemPreviewType.
type FilesystemPreviewType int

// FilesystemPreviewType_if_always is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS.
const FilesystemPreviewType_if_always = FilesystemPreviewType(0)

// FilesystemPreviewType_if_local is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL.
const FilesystemPreviewType_if_local = FilesystemPreviewType(1)

// FilesystemPreviewType_never is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_NEVER.
const FilesystemPreviewType_never = FilesystemPreviewType(2)

// IOErrorEnum is a representation of the C enumeration GIOErrorEnum.
type IOErrorEnum int

// IOErrorEnum_failed is a representation of the C enumeration member G_IO_ERROR_FAILED.
const IOErrorEnum_failed = IOErrorEnum(0)

// IOErrorEnum_not_found is a representation of the C enumeration member G_IO_ERROR_NOT_FOUND.
const IOErrorEnum_not_found = IOErrorEnum(1)

// IOErrorEnum_exists is a representation of the C enumeration member G_IO_ERROR_EXISTS.
const IOErrorEnum_exists = IOErrorEnum(2)

// IOErrorEnum_is_directory is a representation of the C enumeration member G_IO_ERROR_IS_DIRECTORY.
const IOErrorEnum_is_directory = IOErrorEnum(3)

// IOErrorEnum_not_directory is a representation of the C enumeration member G_IO_ERROR_NOT_DIRECTORY.
const IOErrorEnum_not_directory = IOErrorEnum(4)

// IOErrorEnum_not_empty is a representation of the C enumeration member G_IO_ERROR_NOT_EMPTY.
const IOErrorEnum_not_empty = IOErrorEnum(5)

// IOErrorEnum_not_regular_file is a representation of the C enumeration member G_IO_ERROR_NOT_REGULAR_FILE.
const IOErrorEnum_not_regular_file = IOErrorEnum(6)

// IOErrorEnum_not_symbolic_link is a representation of the C enumeration member G_IO_ERROR_NOT_SYMBOLIC_LINK.
const IOErrorEnum_not_symbolic_link = IOErrorEnum(7)

// IOErrorEnum_not_mountable_file is a representation of the C enumeration member G_IO_ERROR_NOT_MOUNTABLE_FILE.
const IOErrorEnum_not_mountable_file = IOErrorEnum(8)

// IOErrorEnum_filename_too_long is a representation of the C enumeration member G_IO_ERROR_FILENAME_TOO_LONG.
const IOErrorEnum_filename_too_long = IOErrorEnum(9)

// IOErrorEnum_invalid_filename is a representation of the C enumeration member G_IO_ERROR_INVALID_FILENAME.
const IOErrorEnum_invalid_filename = IOErrorEnum(10)

// IOErrorEnum_too_many_links is a representation of the C enumeration member G_IO_ERROR_TOO_MANY_LINKS.
const IOErrorEnum_too_many_links = IOErrorEnum(11)

// IOErrorEnum_no_space is a representation of the C enumeration member G_IO_ERROR_NO_SPACE.
const IOErrorEnum_no_space = IOErrorEnum(12)

// IOErrorEnum_invalid_argument is a representation of the C enumeration member G_IO_ERROR_INVALID_ARGUMENT.
const IOErrorEnum_invalid_argument = IOErrorEnum(13)

// IOErrorEnum_permission_denied is a representation of the C enumeration member G_IO_ERROR_PERMISSION_DENIED.
const IOErrorEnum_permission_denied = IOErrorEnum(14)

// IOErrorEnum_not_supported is a representation of the C enumeration member G_IO_ERROR_NOT_SUPPORTED.
const IOErrorEnum_not_supported = IOErrorEnum(15)

// IOErrorEnum_not_mounted is a representation of the C enumeration member G_IO_ERROR_NOT_MOUNTED.
const IOErrorEnum_not_mounted = IOErrorEnum(16)

// IOErrorEnum_already_mounted is a representation of the C enumeration member G_IO_ERROR_ALREADY_MOUNTED.
const IOErrorEnum_already_mounted = IOErrorEnum(17)

// IOErrorEnum_closed is a representation of the C enumeration member G_IO_ERROR_CLOSED.
const IOErrorEnum_closed = IOErrorEnum(18)

// IOErrorEnum_cancelled is a representation of the C enumeration member G_IO_ERROR_CANCELLED.
const IOErrorEnum_cancelled = IOErrorEnum(19)

// IOErrorEnum_pending is a representation of the C enumeration member G_IO_ERROR_PENDING.
const IOErrorEnum_pending = IOErrorEnum(20)

// IOErrorEnum_read_only is a representation of the C enumeration member G_IO_ERROR_READ_ONLY.
const IOErrorEnum_read_only = IOErrorEnum(21)

// IOErrorEnum_cant_create_backup is a representation of the C enumeration member G_IO_ERROR_CANT_CREATE_BACKUP.
const IOErrorEnum_cant_create_backup = IOErrorEnum(22)

// IOErrorEnum_wrong_etag is a representation of the C enumeration member G_IO_ERROR_WRONG_ETAG.
const IOErrorEnum_wrong_etag = IOErrorEnum(23)

// IOErrorEnum_timed_out is a representation of the C enumeration member G_IO_ERROR_TIMED_OUT.
const IOErrorEnum_timed_out = IOErrorEnum(24)

// IOErrorEnum_would_recurse is a representation of the C enumeration member G_IO_ERROR_WOULD_RECURSE.
const IOErrorEnum_would_recurse = IOErrorEnum(25)

// IOErrorEnum_busy is a representation of the C enumeration member G_IO_ERROR_BUSY.
const IOErrorEnum_busy = IOErrorEnum(26)

// IOErrorEnum_would_block is a representation of the C enumeration member G_IO_ERROR_WOULD_BLOCK.
const IOErrorEnum_would_block = IOErrorEnum(27)

// IOErrorEnum_host_not_found is a representation of the C enumeration member G_IO_ERROR_HOST_NOT_FOUND.
const IOErrorEnum_host_not_found = IOErrorEnum(28)

// IOErrorEnum_would_merge is a representation of the C enumeration member G_IO_ERROR_WOULD_MERGE.
const IOErrorEnum_would_merge = IOErrorEnum(29)

// IOErrorEnum_failed_handled is a representation of the C enumeration member G_IO_ERROR_FAILED_HANDLED.
const IOErrorEnum_failed_handled = IOErrorEnum(30)

// IOErrorEnum_too_many_open_files is a representation of the C enumeration member G_IO_ERROR_TOO_MANY_OPEN_FILES.
const IOErrorEnum_too_many_open_files = IOErrorEnum(31)

// IOErrorEnum_not_initialized is a representation of the C enumeration member G_IO_ERROR_NOT_INITIALIZED.
const IOErrorEnum_not_initialized = IOErrorEnum(32)

// IOErrorEnum_address_in_use is a representation of the C enumeration member G_IO_ERROR_ADDRESS_IN_USE.
const IOErrorEnum_address_in_use = IOErrorEnum(33)

// IOErrorEnum_partial_input is a representation of the C enumeration member G_IO_ERROR_PARTIAL_INPUT.
const IOErrorEnum_partial_input = IOErrorEnum(34)

// IOErrorEnum_invalid_data is a representation of the C enumeration member G_IO_ERROR_INVALID_DATA.
const IOErrorEnum_invalid_data = IOErrorEnum(35)

// IOErrorEnum_dbus_error is a representation of the C enumeration member G_IO_ERROR_DBUS_ERROR.
const IOErrorEnum_dbus_error = IOErrorEnum(36)

// IOErrorEnum_host_unreachable is a representation of the C enumeration member G_IO_ERROR_HOST_UNREACHABLE.
const IOErrorEnum_host_unreachable = IOErrorEnum(37)

// IOErrorEnum_network_unreachable is a representation of the C enumeration member G_IO_ERROR_NETWORK_UNREACHABLE.
const IOErrorEnum_network_unreachable = IOErrorEnum(38)

// IOErrorEnum_connection_refused is a representation of the C enumeration member G_IO_ERROR_CONNECTION_REFUSED.
const IOErrorEnum_connection_refused = IOErrorEnum(39)

// IOErrorEnum_proxy_failed is a representation of the C enumeration member G_IO_ERROR_PROXY_FAILED.
const IOErrorEnum_proxy_failed = IOErrorEnum(40)

// IOErrorEnum_proxy_auth_failed is a representation of the C enumeration member G_IO_ERROR_PROXY_AUTH_FAILED.
const IOErrorEnum_proxy_auth_failed = IOErrorEnum(41)

// IOErrorEnum_proxy_need_auth is a representation of the C enumeration member G_IO_ERROR_PROXY_NEED_AUTH.
const IOErrorEnum_proxy_need_auth = IOErrorEnum(42)

// IOErrorEnum_proxy_not_allowed is a representation of the C enumeration member G_IO_ERROR_PROXY_NOT_ALLOWED.
const IOErrorEnum_proxy_not_allowed = IOErrorEnum(43)

// IOErrorEnum_broken_pipe is a representation of the C enumeration member G_IO_ERROR_BROKEN_PIPE.
const IOErrorEnum_broken_pipe = IOErrorEnum(44)

// IOErrorEnum_connection_closed is a representation of the C enumeration member G_IO_ERROR_CONNECTION_CLOSED.
const IOErrorEnum_connection_closed = IOErrorEnum(44)

// IOErrorEnum_not_connected is a representation of the C enumeration member G_IO_ERROR_NOT_CONNECTED.
const IOErrorEnum_not_connected = IOErrorEnum(45)

// IOErrorEnum_message_too_large is a representation of the C enumeration member G_IO_ERROR_MESSAGE_TOO_LARGE.
const IOErrorEnum_message_too_large = IOErrorEnum(46)

// MountOperationResult is a representation of the C enumeration GMountOperationResult.
type MountOperationResult int

// MountOperationResult_handled is a representation of the C enumeration member G_MOUNT_OPERATION_HANDLED.
const MountOperationResult_handled = MountOperationResult(0)

// MountOperationResult_aborted is a representation of the C enumeration member G_MOUNT_OPERATION_ABORTED.
const MountOperationResult_aborted = MountOperationResult(1)

// MountOperationResult_unhandled is a representation of the C enumeration member G_MOUNT_OPERATION_UNHANDLED.
const MountOperationResult_unhandled = MountOperationResult(2)

// PasswordSave is a representation of the C enumeration GPasswordSave.
type PasswordSave int

// PasswordSave_never is a representation of the C enumeration member G_PASSWORD_SAVE_NEVER.
const PasswordSave_never = PasswordSave(0)

// PasswordSave_for_session is a representation of the C enumeration member G_PASSWORD_SAVE_FOR_SESSION.
const PasswordSave_for_session = PasswordSave(1)

// PasswordSave_permanently is a representation of the C enumeration member G_PASSWORD_SAVE_PERMANENTLY.
const PasswordSave_permanently = PasswordSave(2)

// UNSUPPORTED : g_action_parse_detailed_name : throws

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// ContentTypeCanBeExecutable is analogous to the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) {
	sys_type_ := type_
	gio.Fn_g_content_type_can_be_executable(sys_type_)
}

// ContentTypeEquals is analogous to the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) {
	sys_type1 := type1
	sys_type2 := type2
	gio.Fn_g_content_type_equals(sys_type1, sys_type2)
}

// ContentTypeFromMimeType is analogous to the C function g_content_type_from_mime_type.
func ContentTypeFromMimeType(mimeType string) {
	sys_mimeType := mimeType
	gio.Fn_g_content_type_from_mime_type(sys_mimeType)
}

// ContentTypeGetDescription is analogous to the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) {
	sys_type_ := type_
	gio.Fn_g_content_type_get_description(sys_type_)
}

// ContentTypeGetIcon is analogous to the C function g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) {
	sys_type_ := type_
	gio.Fn_g_content_type_get_icon(sys_type_)
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// ContentTypeGetMimeType is analogous to the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) {
	sys_type_ := type_
	gio.Fn_g_content_type_get_mime_type(sys_type_)
}

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// ContentTypeIsA is analogous to the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) {
	sys_type_ := type_
	sys_supertype := supertype
	gio.Fn_g_content_type_is_a(sys_type_, sys_supertype)
}

// ContentTypeIsUnknown is analogous to the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) {
	sys_type_ := type_
	gio.Fn_g_content_type_is_unknown(sys_type_)
}

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// ContentTypesGetRegistered is analogous to the C function g_content_types_get_registered.
func ContentTypesGetRegistered() {
	gio.Fn_g_content_types_get_registered()
}

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has array [in]out, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// IoErrorFromErrno is analogous to the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int) {
	sys_errNo := errNo
	gio.Fn_g_io_error_from_errno(sys_errNo)
}

// IoErrorQuark is analogous to the C function g_io_error_quark.
func IoErrorQuark() {
	gio.Fn_g_io_error_quark()
}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// IoSchedulerCancelAllJobs is analogous to the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	gio.Fn_g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UnixIsMountPathSystemInternal is analogous to the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) {
	sys_mountPath := mountPath
	gio.Fn_g_unix_is_mount_path_system_internal(sys_mountPath)
}

// UNSUPPORTED : g_unix_mount_at : has array [in]out, time_read

// UnixMountCompare is analogous to the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) {
	sys_mount1 := mount1.ToC()
	sys_mount2 := mount2.ToC()
	gio.Fn_g_unix_mount_compare(sys_mount1, sys_mount2)
}

// UNSUPPORTED : g_unix_mount_for : has array [in]out, time_read

// UnixMountFree is analogous to the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_free(sys_mountEntry)
}

// UnixMountGetDevicePath is analogous to the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_get_device_path(sys_mountEntry)
}

// UnixMountGetFsType is analogous to the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_get_fs_type(sys_mountEntry)
}

// UnixMountGetMountPath is analogous to the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_get_mount_path(sys_mountEntry)
}

// UnixMountGuessCanEject is analogous to the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_guess_can_eject(sys_mountEntry)
}

// UnixMountGuessIcon is analogous to the C function g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_guess_icon(sys_mountEntry)
}

// UnixMountGuessName is analogous to the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_guess_name(sys_mountEntry)
}

// UnixMountGuessShouldDisplay is analogous to the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_guess_should_display(sys_mountEntry)
}

// UnixMountIsReadonly is analogous to the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_is_readonly(sys_mountEntry)
}

// UnixMountIsSystemInternal is analogous to the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_is_system_internal(sys_mountEntry)
}

// UnixMountPointsChangedSince is analogous to the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) {
	sys_time := time
	gio.Fn_g_unix_mount_points_changed_since(sys_time)
}

// UNSUPPORTED : g_unix_mount_points_get : has array [in]out, time_read

// UnixMountsChangedSince is analogous to the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) {
	sys_time := time
	gio.Fn_g_unix_mounts_changed_since(sys_time)
}

// UNSUPPORTED : g_unix_mounts_get : has array [in]out, time_read

// ActionEntry is a representation of the C record GActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionEntry that represents the ActionEntry.
func (recv *ActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoIface is a representation of the C record GAppInfoIface.
type AppInfoIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfoIface that represents the AppInfoIface.
func (recv *AppInfoIface) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextClass is a representation of the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextClass that represents the AppLaunchContextClass.
func (recv *AppLaunchContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextPrivate is a representation of the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextPrivate that represents the AppLaunchContextPrivate.
func (recv *AppLaunchContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLinePrivate is a representation of the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLinePrivate that represents the ApplicationCommandLinePrivate.
func (recv *ApplicationCommandLinePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivate is a representation of the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationPrivate that represents the ApplicationPrivate.
func (recv *ApplicationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResultIface is a representation of the C record GAsyncResultIface.
type AsyncResultIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResultIface that represents the AsyncResultIface.
func (recv *AsyncResultIface) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamClass is a representation of the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamClass that represents the BufferedInputStreamClass.
func (recv *BufferedInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamPrivate is a representation of the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamPrivate that represents the BufferedInputStreamPrivate.
func (recv *BufferedInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamClass is a representation of the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamClass that represents the BufferedOutputStreamClass.
func (recv *BufferedOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamPrivate is a representation of the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamPrivate that represents the BufferedOutputStreamPrivate.
func (recv *BufferedOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CancellableClass is a representation of the C record GCancellableClass.
type CancellableClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellableClass that represents the CancellableClass.
func (recv *CancellableClass) ToC() unsafe.Pointer {
	return recv.native
}

// CancellablePrivate is a representation of the C record GCancellablePrivate.
type CancellablePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellablePrivate that represents the CancellablePrivate.
func (recv *CancellablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverterClass is a representation of the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverterClass that represents the CharsetConverterClass.
func (recv *CharsetConverterClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamClass is a representation of the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamClass that represents the ConverterInputStreamClass.
func (recv *ConverterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamPrivate is a representation of the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamPrivate that represents the ConverterInputStreamPrivate.
func (recv *ConverterInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamClass is a representation of the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamClass that represents the ConverterOutputStreamClass.
func (recv *ConverterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamPrivate is a representation of the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamPrivate that represents the ConverterOutputStreamPrivate.
func (recv *ConverterOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonPrivate is a representation of the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeletonPrivate that represents the DBusInterfaceSkeletonPrivate.
func (recv *DBusInterfaceSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientPrivate is a representation of the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClientPrivate that represents the DBusObjectManagerClientPrivate.
func (recv *DBusObjectManagerClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerPrivate is a representation of the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServerPrivate that represents the DBusObjectManagerServerPrivate.
func (recv *DBusObjectManagerServerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyPrivate is a representation of the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxyPrivate that represents the DBusObjectProxyPrivate.
func (recv *DBusObjectProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonPrivate is a representation of the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeletonPrivate that represents the DBusObjectSkeletonPrivate.
func (recv *DBusObjectSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyPrivate is a representation of the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxyPrivate that represents the DBusProxyPrivate.
func (recv *DBusProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamClass is a representation of the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamClass that represents the DataInputStreamClass.
func (recv *DataInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamPrivate is a representation of the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamPrivate that represents the DataInputStreamPrivate.
func (recv *DataInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamClass is a representation of the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamClass that represents the DataOutputStreamClass.
func (recv *DataOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamPrivate is a representation of the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamPrivate that represents the DataOutputStreamPrivate.
func (recv *DataOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoClass is a representation of the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoClass that represents the DesktopAppInfoClass.
func (recv *DesktopAppInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookupIface is a representation of the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookupIface that represents the DesktopAppInfoLookupIface.
func (recv *DesktopAppInfoLookupIface) ToC() unsafe.Pointer {
	return recv.native
}

// DriveIface is a representation of the C record GDriveIface.
type DriveIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDriveIface that represents the DriveIface.
func (recv *DriveIface) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemClass is a representation of the C record GEmblemClass.
type EmblemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemClass that represents the EmblemClass.
func (recv *EmblemClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconClass is a representation of the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconClass that represents the EmblemedIconClass.
func (recv *EmblemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconPrivate is a representation of the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconPrivate that represents the EmblemedIconPrivate.
func (recv *EmblemedIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfo is a representation of the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfo that represents the FileAttributeInfo.
func (recv *FileAttributeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfoList is a representation of the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfoList that represents the FileAttributeInfoList.
func (recv *FileAttributeInfoList) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeMatcher is a representation of the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeMatcher that represents the FileAttributeMatcher.
func (recv *FileAttributeMatcher) ToC() unsafe.Pointer {
	return recv.native
}

// FileDescriptorBasedIface is a representation of the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileDescriptorBasedIface that represents the FileDescriptorBasedIface.
func (recv *FileDescriptorBasedIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorClass is a representation of the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorClass that represents the FileEnumeratorClass.
func (recv *FileEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorPrivate is a representation of the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorPrivate that represents the FileEnumeratorPrivate.
func (recv *FileEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamClass is a representation of the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamClass that represents the FileIOStreamClass.
func (recv *FileIOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamPrivate is a representation of the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamPrivate that represents the FileIOStreamPrivate.
func (recv *FileIOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileIconClass is a representation of the C record GFileIconClass.
type FileIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIconClass that represents the FileIconClass.
func (recv *FileIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIface is a representation of the C record GFileIface.
type FileIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIface that represents the FileIface.
func (recv *FileIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfoClass is a representation of the C record GFileInfoClass.
type FileInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfoClass that represents the FileInfoClass.
func (recv *FileInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamClass is a representation of the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamClass that represents the FileInputStreamClass.
func (recv *FileInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamPrivate is a representation of the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamPrivate that represents the FileInputStreamPrivate.
func (recv *FileInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorClass is a representation of the C record GFileMonitorClass.
type FileMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorClass that represents the FileMonitorClass.
func (recv *FileMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorPrivate is a representation of the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorPrivate that represents the FileMonitorPrivate.
func (recv *FileMonitorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamClass is a representation of the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamClass that represents the FileOutputStreamClass.
func (recv *FileOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamPrivate is a representation of the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamPrivate that represents the FileOutputStreamPrivate.
func (recv *FileOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleterClass is a representation of the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleterClass that represents the FilenameCompleterClass.
func (recv *FilenameCompleterClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStreamClass is a representation of the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStreamClass that represents the FilterInputStreamClass.
func (recv *FilterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStreamClass is a representation of the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStreamClass that represents the FilterOutputStreamClass.
func (recv *FilterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtension is a representation of the C record GIOExtension.
type IOExtension struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtension that represents the IOExtension.
func (recv *IOExtension) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtensionPoint is a representation of the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtensionPoint that represents the IOExtensionPoint.
func (recv *IOExtensionPoint) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleClass is a representation of the C record GIOModuleClass.
type IOModuleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModuleClass that represents the IOModuleClass.
func (recv *IOModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOSchedulerJob is a representation of the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOSchedulerJob that represents the IOSchedulerJob.
func (recv *IOSchedulerJob) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamAdapter is a representation of the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamAdapter that represents the IOStreamAdapter.
func (recv *IOStreamAdapter) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamClass is a representation of the C record GIOStreamClass.
type IOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamClass that represents the IOStreamClass.
func (recv *IOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamPrivate is a representation of the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamPrivate that represents the IOStreamPrivate.
func (recv *IOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IconIface is a representation of the C record GIconIface.
type IconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIconIface that represents the IconIface.
func (recv *IconIface) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressClass is a representation of the C record GInetAddressClass.
type InetAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressClass that represents the InetAddressClass.
func (recv *InetAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskClass is a representation of the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskClass that represents the InetAddressMaskClass.
func (recv *InetAddressMaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskPrivate is a representation of the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskPrivate that represents the InetAddressMaskPrivate.
func (recv *InetAddressMaskPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressPrivate is a representation of the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressPrivate that represents the InetAddressPrivate.
func (recv *InetAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressClass is a representation of the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressClass that represents the InetSocketAddressClass.
func (recv *InetSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressPrivate is a representation of the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressPrivate that represents the InetSocketAddressPrivate.
func (recv *InetSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamClass is a representation of the C record GInputStreamClass.
type InputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamClass that represents the InputStreamClass.
func (recv *InputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamPrivate is a representation of the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamPrivate that represents the InputStreamPrivate.
func (recv *InputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreClass is a representation of the C record GListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStoreClass that represents the ListStoreClass.
func (recv *ListStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIconIface is a representation of the C record GLoadableIconIface.
type LoadableIconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIconIface that represents the LoadableIconIface.
func (recv *LoadableIconIface) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamClass is a representation of the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamClass that represents the MemoryInputStreamClass.
func (recv *MemoryInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamPrivate is a representation of the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamPrivate that represents the MemoryInputStreamPrivate.
func (recv *MemoryInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamClass is a representation of the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamClass that represents the MemoryOutputStreamClass.
func (recv *MemoryOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamPrivate is a representation of the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamPrivate that represents the MemoryOutputStreamPrivate.
func (recv *MemoryOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterClass is a representation of the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterClass that represents the MenuAttributeIterClass.
func (recv *MenuAttributeIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterPrivate is a representation of the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterPrivate that represents the MenuAttributeIterPrivate.
func (recv *MenuAttributeIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterClass is a representation of the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterClass that represents the MenuLinkIterClass.
func (recv *MenuLinkIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterPrivate is a representation of the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterPrivate that represents the MenuLinkIterPrivate.
func (recv *MenuLinkIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelClass is a representation of the C record GMenuModelClass.
type MenuModelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelClass that represents the MenuModelClass.
func (recv *MenuModelClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelPrivate is a representation of the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelPrivate that represents the MenuModelPrivate.
func (recv *MenuModelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MountIface is a representation of the C record GMountIface.
type MountIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountIface that represents the MountIface.
func (recv *MountIface) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationClass is a representation of the C record GMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationClass that represents the MountOperationClass.
func (recv *MountOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivate is a representation of the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationPrivate that represents the MountOperationPrivate.
func (recv *MountOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// NativeVolumeMonitorClass is a representation of the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeVolumeMonitorClass that represents the NativeVolumeMonitorClass.
func (recv *NativeVolumeMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressClass is a representation of the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressClass that represents the NetworkAddressClass.
func (recv *NetworkAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressPrivate is a representation of the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressPrivate that represents the NetworkAddressPrivate.
func (recv *NetworkAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServiceClass is a representation of the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServiceClass that represents the NetworkServiceClass.
func (recv *NetworkServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServicePrivate is a representation of the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServicePrivate that represents the NetworkServicePrivate.
func (recv *NetworkServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamClass is a representation of the C record GOutputStreamClass.
type OutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamClass that represents the OutputStreamClass.
func (recv *OutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamPrivate is a representation of the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamPrivate that represents the OutputStreamPrivate.
func (recv *OutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionClass is a representation of the C record GPermissionClass.
type PermissionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionClass that represents the PermissionClass.
func (recv *PermissionClass) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionPrivate is a representation of the C record GPermissionPrivate.
type PermissionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionPrivate that represents the PermissionPrivate.
func (recv *PermissionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorClass is a representation of the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorClass that represents the ProxyAddressEnumeratorClass.
func (recv *ProxyAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorPrivate is a representation of the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorPrivate that represents the ProxyAddressEnumeratorPrivate.
func (recv *ProxyAddressEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressPrivate is a representation of the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressPrivate that represents the ProxyAddressPrivate.
func (recv *ProxyAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyResolverInterface is a representation of the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyResolverInterface that represents the ProxyResolverInterface.
func (recv *ProxyResolverInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverClass is a representation of the C record GResolverClass.
type ResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverClass that represents the ResolverClass.
func (recv *ResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverPrivate is a representation of the C record GResolverPrivate.
type ResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverPrivate that represents the ResolverPrivate.
func (recv *ResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SeekableIface is a representation of the C record GSeekableIface.
type SeekableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekableIface that represents the SeekableIface.
func (recv *SeekableIface) ToC() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// SettingsClass is a representation of the C record GSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsClass that represents the SettingsClass.
func (recv *SettingsClass) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsPrivate is a representation of the C record GSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsPrivate that represents the SettingsPrivate.
func (recv *SettingsPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaKey is a representation of the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchemaKey that represents the SettingsSchemaKey.
func (recv *SettingsSchemaKey) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupClass is a representation of the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupClass that represents the SimpleActionGroupClass.
func (recv *SimpleActionGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupPrivate is a representation of the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupPrivate that represents the SimpleActionGroupPrivate.
func (recv *SimpleActionGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultClass is a representation of the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResultClass that represents the SimpleAsyncResultClass.
func (recv *SimpleAsyncResultClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverClass is a representation of the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverClass that represents the SimpleProxyResolverClass.
func (recv *SimpleProxyResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverPrivate is a representation of the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverPrivate that represents the SimpleProxyResolverPrivate.
func (recv *SimpleProxyResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressClass is a representation of the C record GSocketAddressClass.
type SocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressClass that represents the SocketAddressClass.
func (recv *SocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumeratorClass is a representation of the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumeratorClass that represents the SocketAddressEnumeratorClass.
func (recv *SocketAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClass is a representation of the C record GSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClass that represents the SocketClass.
func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientClass is a representation of the C record GSocketClientClass.
type SocketClientClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientClass that represents the SocketClientClass.
func (recv *SocketClientClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientPrivate is a representation of the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientPrivate that represents the SocketClientPrivate.
func (recv *SocketClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectableIface is a representation of the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectableIface that represents the SocketConnectableIface.
func (recv *SocketConnectableIface) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionClass is a representation of the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionClass that represents the SocketConnectionClass.
func (recv *SocketConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionPrivate is a representation of the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionPrivate that represents the SocketConnectionPrivate.
func (recv *SocketConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessageClass is a representation of the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessageClass that represents the SocketControlMessageClass.
func (recv *SocketControlMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessagePrivate is a representation of the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessagePrivate that represents the SocketControlMessagePrivate.
func (recv *SocketControlMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerClass is a representation of the C record GSocketListenerClass.
type SocketListenerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerClass that represents the SocketListenerClass.
func (recv *SocketListenerClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerPrivate is a representation of the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerPrivate that represents the SocketListenerPrivate.
func (recv *SocketListenerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketPrivate is a representation of the C record GSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketPrivate that represents the SocketPrivate.
func (recv *SocketPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServiceClass is a representation of the C record GSocketServiceClass.
type SocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServiceClass that represents the SocketServiceClass.
func (recv *SocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServicePrivate is a representation of the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServicePrivate that represents the SocketServicePrivate.
func (recv *SocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SrvTarget is a representation of the C record GSrvTarget.
type SrvTarget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSrvTarget that represents the SrvTarget.
func (recv *SrvTarget) ToC() unsafe.Pointer {
	return recv.native
}

// StaticResource is a representation of the C record GStaticResource.
type StaticResource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GStaticResource that represents the StaticResource.
func (recv *StaticResource) ToC() unsafe.Pointer {
	return recv.native
}

// TaskClass is a representation of the C record GTaskClass.
type TaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTaskClass that represents the TaskClass.
func (recv *TaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionClass is a representation of the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionClass that represents the TcpConnectionClass.
func (recv *TcpConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionPrivate is a representation of the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionPrivate that represents the TcpConnectionPrivate.
func (recv *TcpConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionClass is a representation of the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionClass that represents the TcpWrapperConnectionClass.
func (recv *TcpWrapperConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionPrivate is a representation of the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionPrivate that represents the TcpWrapperConnectionPrivate.
func (recv *TcpWrapperConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIconClass is a representation of the C record GThemedIconClass.
type ThemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIconClass that represents the ThemedIconClass.
func (recv *ThemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServiceClass is a representation of the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServiceClass that represents the ThreadedSocketServiceClass.
func (recv *ThreadedSocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServicePrivate is a representation of the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServicePrivate that represents the ThreadedSocketServicePrivate.
func (recv *ThreadedSocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificateClass is a representation of the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificateClass that represents the TlsCertificateClass.
func (recv *TlsCertificateClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificatePrivate is a representation of the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificatePrivate that represents the TlsCertificatePrivate.
func (recv *TlsCertificatePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionClass is a representation of the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionClass that represents the TlsConnectionClass.
func (recv *TlsConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionPrivate is a representation of the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionPrivate that represents the TlsConnectionPrivate.
func (recv *TlsConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabasePrivate is a representation of the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabasePrivate that represents the TlsDatabasePrivate.
func (recv *TlsDatabasePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsFileDatabaseInterface is a representation of the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsFileDatabaseInterface that represents the TlsFileDatabaseInterface.
func (recv *TlsFileDatabaseInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionPrivate is a representation of the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteractionPrivate that represents the TlsInteractionPrivate.
func (recv *TlsInteractionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordClass is a representation of the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordClass that represents the TlsPasswordClass.
func (recv *TlsPasswordClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordPrivate is a representation of the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordPrivate that represents the TlsPasswordPrivate.
func (recv *TlsPasswordPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionClass is a representation of the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionClass that represents the UnixConnectionClass.
func (recv *UnixConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionPrivate is a representation of the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionPrivate that represents the UnixConnectionPrivate.
func (recv *UnixConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessagePrivate is a representation of the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessagePrivate that represents the UnixCredentialsMessagePrivate.
func (recv *UnixCredentialsMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListClass is a representation of the C record GUnixFDListClass.
type UnixFDListClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListClass that represents the UnixFDListClass.
func (recv *UnixFDListClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListPrivate is a representation of the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListPrivate that represents the UnixFDListPrivate.
func (recv *UnixFDListPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessageClass is a representation of the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessageClass that represents the UnixFDMessageClass.
func (recv *UnixFDMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessagePrivate is a representation of the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessagePrivate that represents the UnixFDMessagePrivate.
func (recv *UnixFDMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamClass is a representation of the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamClass that represents the UnixInputStreamClass.
func (recv *UnixInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamPrivate is a representation of the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamPrivate that represents the UnixInputStreamPrivate.
func (recv *UnixInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountEntry is a representation of the C record GUnixMountEntry.
type UnixMountEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountEntry that represents the UnixMountEntry.
func (recv *UnixMountEntry) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorClass is a representation of the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitorClass that represents the UnixMountMonitorClass.
func (recv *UnixMountMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountPoint is a representation of the C record GUnixMountPoint.
type UnixMountPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountPoint that represents the UnixMountPoint.
func (recv *UnixMountPoint) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamClass is a representation of the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamClass that represents the UnixOutputStreamClass.
func (recv *UnixOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamPrivate is a representation of the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamPrivate that represents the UnixOutputStreamPrivate.
func (recv *UnixOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressClass is a representation of the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressClass that represents the UnixSocketAddressClass.
func (recv *UnixSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressPrivate is a representation of the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressPrivate that represents the UnixSocketAddressPrivate.
func (recv *UnixSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// VfsClass is a representation of the C record GVfsClass.
type VfsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfsClass that represents the VfsClass.
func (recv *VfsClass) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeIface is a representation of the C record GVolumeIface.
type VolumeIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeIface that represents the VolumeIface.
func (recv *VolumeIface) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitorClass is a representation of the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitorClass that represents the VolumeMonitorClass.
func (recv *VolumeMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressorClass is a representation of the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressorClass that represents the ZlibCompressorClass.
func (recv *ZlibCompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressorClass is a representation of the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressorClass that represents the ZlibDecompressorClass.
func (recv *ZlibDecompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContext is a representation of the C record GAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContext that represents the AppLaunchContext.
func (recv *AppLaunchContext) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLine is a representation of the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLine that represents the ApplicationCommandLine.
func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStream is a representation of the C record GBufferedInputStream.
type BufferedInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStream that represents the BufferedInputStream.
func (recv *BufferedInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStream is a representation of the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStream that represents the BufferedOutputStream.
func (recv *BufferedOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// Cancellable is a representation of the C record GCancellable.
type Cancellable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellable that represents the Cancellable.
func (recv *Cancellable) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverter is a representation of the C record GCharsetConverter.
type CharsetConverter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverter that represents the CharsetConverter.
func (recv *CharsetConverter) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStream is a representation of the C record GConverterInputStream.
type ConverterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStream that represents the ConverterInputStream.
func (recv *ConverterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStream is a representation of the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStream that represents the ConverterOutputStream.
func (recv *ConverterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DBusActionGroup is a representation of the C record GDBusActionGroup.
type DBusActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusActionGroup that represents the DBusActionGroup.
func (recv *DBusActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMenuModel is a representation of the C record GDBusMenuModel.
type DBusMenuModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMenuModel that represents the DBusMenuModel.
func (recv *DBusMenuModel) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStream is a representation of the C record GDataInputStream.
type DataInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStream that represents the DataInputStream.
func (recv *DataInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStream is a representation of the C record GDataOutputStream.
type DataOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStream that represents the DataOutputStream.
func (recv *DataOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfo is a representation of the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfo that represents the DesktopAppInfo.
func (recv *DesktopAppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// Emblem is a representation of the C record GEmblem.
type Emblem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblem that represents the Emblem.
func (recv *Emblem) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIcon is a representation of the C record GEmblemedIcon.
type EmblemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIcon that represents the EmblemedIcon.
func (recv *EmblemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumerator is a representation of the C record GFileEnumerator.
type FileEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumerator that represents the FileEnumerator.
func (recv *FileEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// FileIcon is a representation of the C record GFileIcon.
type FileIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIcon that represents the FileIcon.
func (recv *FileIcon) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfo is a representation of the C record GFileInfo.
type FileInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfo that represents the FileInfo.
func (recv *FileInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStream is a representation of the C record GFileInputStream.
type FileInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStream that represents the FileInputStream.
func (recv *FileInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitor is a representation of the C record GFileMonitor.
type FileMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitor that represents the FileMonitor.
func (recv *FileMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStream is a representation of the C record GFileOutputStream.
type FileOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStream that represents the FileOutputStream.
func (recv *FileOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleter is a representation of the C record GFilenameCompleter.
type FilenameCompleter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleter that represents the FilenameCompleter.
func (recv *FilenameCompleter) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStream is a representation of the C record GFilterInputStream.
type FilterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStream that represents the FilterInputStream.
func (recv *FilterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStream is a representation of the C record GFilterOutputStream.
type FilterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStream that represents the FilterOutputStream.
func (recv *FilterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// IOModule is a representation of the C record GIOModule.
type IOModule struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModule that represents the IOModule.
func (recv *IOModule) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddress is a representation of the C record GInetAddress.
type InetAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddress that represents the InetAddress.
func (recv *InetAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddress is a representation of the C record GInetSocketAddress.
type InetSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddress that represents the InetSocketAddress.
func (recv *InetSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InputStream is a representation of the C record GInputStream.
type InputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStream that represents the InputStream.
func (recv *InputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ListStore is a representation of the C record GListStore.
type ListStore struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStore that represents the ListStore.
func (recv *ListStore) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStream is a representation of the C record GMemoryInputStream.
type MemoryInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStream that represents the MemoryInputStream.
func (recv *MemoryInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStream is a representation of the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStream that represents the MemoryOutputStream.
func (recv *MemoryOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperation is a representation of the C record GMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperation that represents the MountOperation.
func (recv *MountOperation) ToC() unsafe.Pointer {
	return recv.native
}

// NativeSocketAddress is a representation of the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeSocketAddress that represents the NativeSocketAddress.
func (recv *NativeSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NativeVolumeMonitor is a representation of the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeVolumeMonitor that represents the NativeVolumeMonitor.
func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddress is a representation of the C record GNetworkAddress.
type NetworkAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddress that represents the NetworkAddress.
func (recv *NetworkAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkService is a representation of the C record GNetworkService.
type NetworkService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkService that represents the NetworkService.
func (recv *NetworkService) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStream is a representation of the C record GOutputStream.
type OutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStream that represents the OutputStream.
func (recv *OutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// Permission is a representation of the C record GPermission.
type Permission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermission that represents the Permission.
func (recv *Permission) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumerator is a representation of the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumerator that represents the ProxyAddressEnumerator.
func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// Resolver is a representation of the C record GResolver.
type Resolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolver that represents the Resolver.
func (recv *Resolver) ToC() unsafe.Pointer {
	return recv.native
}

// Settings is a representation of the C record GSettings.
type Settings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettings that represents the Settings.
func (recv *Settings) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsBackend is a representation of the C record GSettingsBackend.
type SettingsBackend struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsBackend that represents the SettingsBackend.
func (recv *SettingsBackend) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAction is a representation of the C record GSimpleAction.
type SimpleAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAction that represents the SimpleAction.
func (recv *SimpleAction) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResult is a representation of the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResult that represents the SimpleAsyncResult.
func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// SimplePermission is a representation of the C record GSimplePermission.
type SimplePermission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimplePermission that represents the SimplePermission.
func (recv *SimplePermission) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddress is a representation of the C record GSocketAddress.
type SocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddress that represents the SocketAddress.
func (recv *SocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumerator is a representation of the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumerator that represents the SocketAddressEnumerator.
func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// Task is a representation of the C record GTask.
type Task struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTask that represents the Task.
func (recv *Task) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIcon is a representation of the C record GThemedIcon.
type ThemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIcon that represents the ThemedIcon.
func (recv *ThemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDList is a representation of the C record GUnixFDList.
type UnixFDList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDList that represents the UnixFDList.
func (recv *UnixFDList) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessage is a representation of the C record GUnixFDMessage.
type UnixFDMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessage that represents the UnixFDMessage.
func (recv *UnixFDMessage) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStream is a representation of the C record GUnixInputStream.
type UnixInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStream that represents the UnixInputStream.
func (recv *UnixInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitor is a representation of the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitor that represents the UnixMountMonitor.
func (recv *UnixMountMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStream is a representation of the C record GUnixOutputStream.
type UnixOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStream that represents the UnixOutputStream.
func (recv *UnixOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddress is a representation of the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddress that represents the UnixSocketAddress.
func (recv *UnixSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// Vfs is a representation of the C record GVfs.
type Vfs struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfs that represents the Vfs.
func (recv *Vfs) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitor is a representation of the C record GVolumeMonitor.
type VolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitor that represents the VolumeMonitor.
func (recv *VolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressor is a representation of the C record GZlibCompressor.
type ZlibCompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressor that represents the ZlibCompressor.
func (recv *ZlibCompressor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressor is a representation of the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressor that represents the ZlibDecompressor.
func (recv *ZlibDecompressor) ToC() unsafe.Pointer {
	return recv.native
}

// Action is a representation of the C interface GAction.
type Action struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAction that represents the Action.
func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroup is a representation of the C interface GActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionGroup that represents the ActionGroup.
func (recv *ActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfo is a representation of the C interface GAppInfo.
type AppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfo that represents the AppInfo.
func (recv *AppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResult is a representation of the C interface GAsyncResult.
type AsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResult that represents the AsyncResult.
func (recv *AsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObject is a representation of the C interface GDBusObject.
type DBusObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObject that represents the DBusObject.
func (recv *DBusObject) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManager is a representation of the C interface GDBusObjectManager.
type DBusObjectManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManager that represents the DBusObjectManager.
func (recv *DBusObjectManager) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookup is a representation of the C interface GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookup that represents the DesktopAppInfoLookup.
func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {
	return recv.native
}

// Drive is a representation of the C interface GDrive.
type Drive struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDrive that represents the Drive.
func (recv *Drive) ToC() unsafe.Pointer {
	return recv.native
}

// File is a representation of the C interface GFile.
type File struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFile that represents the File.
func (recv *File) ToC() unsafe.Pointer {
	return recv.native
}

// Icon is a representation of the C interface GIcon.
type Icon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIcon that represents the Icon.
func (recv *Icon) ToC() unsafe.Pointer {
	return recv.native
}

// ListModel is a representation of the C interface GListModel.
type ListModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListModel that represents the ListModel.
func (recv *ListModel) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIcon is a representation of the C interface GLoadableIcon.
type LoadableIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIcon that represents the LoadableIcon.
func (recv *LoadableIcon) ToC() unsafe.Pointer {
	return recv.native
}

// Mount is a representation of the C interface GMount.
type Mount struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMount that represents the Mount.
func (recv *Mount) ToC() unsafe.Pointer {
	return recv.native
}

// Seekable is a representation of the C interface GSeekable.
type Seekable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekable that represents the Seekable.
func (recv *Seekable) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectable is a representation of the C interface GSocketConnectable.
type SocketConnectable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectable that represents the SocketConnectable.
func (recv *SocketConnectable) ToC() unsafe.Pointer {
	return recv.native
}

// Volume is a representation of the C interface GVolume.
type Volume struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolume that represents the Volume.
func (recv *Volume) ToC() unsafe.Pointer {
	return recv.native
}
