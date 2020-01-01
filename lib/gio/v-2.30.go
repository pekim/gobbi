// Code generated - DO NOT EDIT.
// +build gio_2.30

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL = "mountable::can-poll"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_START is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START = "mountable::can-start"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED = "mountable::can-start-degraded"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP = "mountable::can-stop"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT = "mountable::can-unmount"

// FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI.
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI = "mountable::hal-udi"

// FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC = "mountable::is-media-check-automatic"

// FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE = "mountable::start-stop-type"

// FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE.
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE = "mountable::unix-device"

// FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE.
//
// since 2.22
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE = "mountable::unix-device-file"

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

// FILE_ATTRIBUTE_TRASH_DELETION_DATE is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_DELETION_DATE.
//
// since 2.24
const FILE_ATTRIBUTE_TRASH_DELETION_DATE = "trash::deletion-date"

// FILE_ATTRIBUTE_TRASH_ITEM_COUNT is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT.
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT = "trash::item-count"

// FILE_ATTRIBUTE_TRASH_ORIG_PATH is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_ORIG_PATH.
//
// since 2.24
const FILE_ATTRIBUTE_TRASH_ORIG_PATH = "trash::orig-path"

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

// NETWORK_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_NETWORK_MONITOR_EXTENSION_POINT_NAME.
//
// since 2.30
const NETWORK_MONITOR_EXTENSION_POINT_NAME = "gio-network-monitor"

// PROXY_EXTENSION_POINT_NAME is a representation of the C constant G_PROXY_EXTENSION_POINT_NAME.
//
// since 2.26
const PROXY_EXTENSION_POINT_NAME = "gio-proxy"

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

// ApplicationFlags is a representation of the C bitfield GApplicationFlags.
type ApplicationFlags int

// ApplicationFlags_flags_none is a representation of the C bitfield member G_APPLICATION_FLAGS_NONE.
const ApplicationFlags_flags_none = ApplicationFlags(0)

// ApplicationFlags_is_service is a representation of the C bitfield member G_APPLICATION_IS_SERVICE.
const ApplicationFlags_is_service = ApplicationFlags(1)

// ApplicationFlags_is_launcher is a representation of the C bitfield member G_APPLICATION_IS_LAUNCHER.
const ApplicationFlags_is_launcher = ApplicationFlags(2)

// ApplicationFlags_handles_open is a representation of the C bitfield member G_APPLICATION_HANDLES_OPEN.
const ApplicationFlags_handles_open = ApplicationFlags(4)

// ApplicationFlags_handles_command_line is a representation of the C bitfield member G_APPLICATION_HANDLES_COMMAND_LINE.
const ApplicationFlags_handles_command_line = ApplicationFlags(8)

// ApplicationFlags_send_environment is a representation of the C bitfield member G_APPLICATION_SEND_ENVIRONMENT.
const ApplicationFlags_send_environment = ApplicationFlags(16)

// ApplicationFlags_non_unique is a representation of the C bitfield member G_APPLICATION_NON_UNIQUE.
const ApplicationFlags_non_unique = ApplicationFlags(32)

// ApplicationFlags_can_override_app_id is a representation of the C bitfield member G_APPLICATION_CAN_OVERRIDE_APP_ID.
const ApplicationFlags_can_override_app_id = ApplicationFlags(64)

// ApplicationFlags_allow_replacement is a representation of the C bitfield member G_APPLICATION_ALLOW_REPLACEMENT.
const ApplicationFlags_allow_replacement = ApplicationFlags(128)

// ApplicationFlags_replace is a representation of the C bitfield member G_APPLICATION_REPLACE.
const ApplicationFlags_replace = ApplicationFlags(256)

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

// BusNameOwnerFlags is a representation of the C bitfield GBusNameOwnerFlags.
type BusNameOwnerFlags int

// BusNameOwnerFlags_none is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_NONE.
const BusNameOwnerFlags_none = BusNameOwnerFlags(0)

// BusNameOwnerFlags_allow_replacement is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_ALLOW_REPLACEMENT.
const BusNameOwnerFlags_allow_replacement = BusNameOwnerFlags(1)

// BusNameOwnerFlags_replace is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_REPLACE.
const BusNameOwnerFlags_replace = BusNameOwnerFlags(2)

// BusNameOwnerFlags_do_not_queue is a representation of the C bitfield member G_BUS_NAME_OWNER_FLAGS_DO_NOT_QUEUE.
const BusNameOwnerFlags_do_not_queue = BusNameOwnerFlags(4)

// BusNameWatcherFlags is a representation of the C bitfield GBusNameWatcherFlags.
type BusNameWatcherFlags int

// BusNameWatcherFlags_none is a representation of the C bitfield member G_BUS_NAME_WATCHER_FLAGS_NONE.
const BusNameWatcherFlags_none = BusNameWatcherFlags(0)

// BusNameWatcherFlags_auto_start is a representation of the C bitfield member G_BUS_NAME_WATCHER_FLAGS_AUTO_START.
const BusNameWatcherFlags_auto_start = BusNameWatcherFlags(1)

// ConverterFlags is a representation of the C bitfield GConverterFlags.
type ConverterFlags int

// ConverterFlags_none is a representation of the C bitfield member G_CONVERTER_NO_FLAGS.
const ConverterFlags_none = ConverterFlags(0)

// ConverterFlags_input_at_end is a representation of the C bitfield member G_CONVERTER_INPUT_AT_END.
const ConverterFlags_input_at_end = ConverterFlags(1)

// ConverterFlags_flush is a representation of the C bitfield member G_CONVERTER_FLUSH.
const ConverterFlags_flush = ConverterFlags(2)

// DBusCallFlags is a representation of the C bitfield GDBusCallFlags.
type DBusCallFlags int

// DBusCallFlags_none is a representation of the C bitfield member G_DBUS_CALL_FLAGS_NONE.
const DBusCallFlags_none = DBusCallFlags(0)

// DBusCallFlags_no_auto_start is a representation of the C bitfield member G_DBUS_CALL_FLAGS_NO_AUTO_START.
const DBusCallFlags_no_auto_start = DBusCallFlags(1)

// DBusCallFlags_allow_interactive_authorization is a representation of the C bitfield member G_DBUS_CALL_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
const DBusCallFlags_allow_interactive_authorization = DBusCallFlags(2)

// DBusCapabilityFlags is a representation of the C bitfield GDBusCapabilityFlags.
type DBusCapabilityFlags int

// DBusCapabilityFlags_none is a representation of the C bitfield member G_DBUS_CAPABILITY_FLAGS_NONE.
const DBusCapabilityFlags_none = DBusCapabilityFlags(0)

// DBusCapabilityFlags_unix_fd_passing is a representation of the C bitfield member G_DBUS_CAPABILITY_FLAGS_UNIX_FD_PASSING.
const DBusCapabilityFlags_unix_fd_passing = DBusCapabilityFlags(1)

// DBusConnectionFlags is a representation of the C bitfield GDBusConnectionFlags.
type DBusConnectionFlags int

// DBusConnectionFlags_none is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_NONE.
const DBusConnectionFlags_none = DBusConnectionFlags(0)

// DBusConnectionFlags_authentication_client is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_CLIENT.
const DBusConnectionFlags_authentication_client = DBusConnectionFlags(1)

// DBusConnectionFlags_authentication_server is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_SERVER.
const DBusConnectionFlags_authentication_server = DBusConnectionFlags(2)

// DBusConnectionFlags_authentication_allow_anonymous is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
const DBusConnectionFlags_authentication_allow_anonymous = DBusConnectionFlags(4)

// DBusConnectionFlags_message_bus_connection is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_MESSAGE_BUS_CONNECTION.
const DBusConnectionFlags_message_bus_connection = DBusConnectionFlags(8)

// DBusConnectionFlags_delay_message_processing is a representation of the C bitfield member G_DBUS_CONNECTION_FLAGS_DELAY_MESSAGE_PROCESSING.
const DBusConnectionFlags_delay_message_processing = DBusConnectionFlags(16)

// DBusInterfaceSkeletonFlags is a representation of the C bitfield GDBusInterfaceSkeletonFlags.
type DBusInterfaceSkeletonFlags int

// DBusInterfaceSkeletonFlags_none is a representation of the C bitfield member G_DBUS_INTERFACE_SKELETON_FLAGS_NONE.
const DBusInterfaceSkeletonFlags_none = DBusInterfaceSkeletonFlags(0)

// DBusInterfaceSkeletonFlags_handle_method_invocations_in_thread is a representation of the C bitfield member G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD.
const DBusInterfaceSkeletonFlags_handle_method_invocations_in_thread = DBusInterfaceSkeletonFlags(1)

// DBusMessageFlags is a representation of the C bitfield GDBusMessageFlags.
type DBusMessageFlags int

// DBusMessageFlags_none is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NONE.
const DBusMessageFlags_none = DBusMessageFlags(0)

// DBusMessageFlags_no_reply_expected is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED.
const DBusMessageFlags_no_reply_expected = DBusMessageFlags(1)

// DBusMessageFlags_no_auto_start is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_NO_AUTO_START.
const DBusMessageFlags_no_auto_start = DBusMessageFlags(2)

// DBusMessageFlags_allow_interactive_authorization is a representation of the C bitfield member G_DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION.
const DBusMessageFlags_allow_interactive_authorization = DBusMessageFlags(4)

// DBusObjectManagerClientFlags is a representation of the C bitfield GDBusObjectManagerClientFlags.
type DBusObjectManagerClientFlags int

// DBusObjectManagerClientFlags_none is a representation of the C bitfield member G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_NONE.
const DBusObjectManagerClientFlags_none = DBusObjectManagerClientFlags(0)

// DBusObjectManagerClientFlags_do_not_auto_start is a representation of the C bitfield member G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START.
const DBusObjectManagerClientFlags_do_not_auto_start = DBusObjectManagerClientFlags(1)

// DBusPropertyInfoFlags is a representation of the C bitfield GDBusPropertyInfoFlags.
type DBusPropertyInfoFlags int

// DBusPropertyInfoFlags_none is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_NONE.
const DBusPropertyInfoFlags_none = DBusPropertyInfoFlags(0)

// DBusPropertyInfoFlags_readable is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_READABLE.
const DBusPropertyInfoFlags_readable = DBusPropertyInfoFlags(1)

// DBusPropertyInfoFlags_writable is a representation of the C bitfield member G_DBUS_PROPERTY_INFO_FLAGS_WRITABLE.
const DBusPropertyInfoFlags_writable = DBusPropertyInfoFlags(2)

// DBusProxyFlags is a representation of the C bitfield GDBusProxyFlags.
type DBusProxyFlags int

// DBusProxyFlags_none is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_NONE.
const DBusProxyFlags_none = DBusProxyFlags(0)

// DBusProxyFlags_do_not_load_properties is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES.
const DBusProxyFlags_do_not_load_properties = DBusProxyFlags(1)

// DBusProxyFlags_do_not_connect_signals is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS.
const DBusProxyFlags_do_not_connect_signals = DBusProxyFlags(2)

// DBusProxyFlags_do_not_auto_start is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START.
const DBusProxyFlags_do_not_auto_start = DBusProxyFlags(4)

// DBusProxyFlags_get_invalidated_properties is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES.
const DBusProxyFlags_get_invalidated_properties = DBusProxyFlags(8)

// DBusProxyFlags_do_not_auto_start_at_construction is a representation of the C bitfield member G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION.
const DBusProxyFlags_do_not_auto_start_at_construction = DBusProxyFlags(16)

// DBusSendMessageFlags is a representation of the C bitfield GDBusSendMessageFlags.
type DBusSendMessageFlags int

// DBusSendMessageFlags_none is a representation of the C bitfield member G_DBUS_SEND_MESSAGE_FLAGS_NONE.
const DBusSendMessageFlags_none = DBusSendMessageFlags(0)

// DBusSendMessageFlags_preserve_serial is a representation of the C bitfield member G_DBUS_SEND_MESSAGE_FLAGS_PRESERVE_SERIAL.
const DBusSendMessageFlags_preserve_serial = DBusSendMessageFlags(1)

// DBusServerFlags is a representation of the C bitfield GDBusServerFlags.
type DBusServerFlags int

// DBusServerFlags_none is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_NONE.
const DBusServerFlags_none = DBusServerFlags(0)

// DBusServerFlags_run_in_thread is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_RUN_IN_THREAD.
const DBusServerFlags_run_in_thread = DBusServerFlags(1)

// DBusServerFlags_authentication_allow_anonymous is a representation of the C bitfield member G_DBUS_SERVER_FLAGS_AUTHENTICATION_ALLOW_ANONYMOUS.
const DBusServerFlags_authentication_allow_anonymous = DBusServerFlags(2)

// DBusSignalFlags is a representation of the C bitfield GDBusSignalFlags.
type DBusSignalFlags int

// DBusSignalFlags_none is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_NONE.
const DBusSignalFlags_none = DBusSignalFlags(0)

// DBusSignalFlags_no_match_rule is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_NO_MATCH_RULE.
const DBusSignalFlags_no_match_rule = DBusSignalFlags(1)

// DBusSignalFlags_match_arg0_namespace is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_NAMESPACE.
const DBusSignalFlags_match_arg0_namespace = DBusSignalFlags(2)

// DBusSignalFlags_match_arg0_path is a representation of the C bitfield member G_DBUS_SIGNAL_FLAGS_MATCH_ARG0_PATH.
const DBusSignalFlags_match_arg0_path = DBusSignalFlags(4)

// DBusSubtreeFlags is a representation of the C bitfield GDBusSubtreeFlags.
type DBusSubtreeFlags int

// DBusSubtreeFlags_none is a representation of the C bitfield member G_DBUS_SUBTREE_FLAGS_NONE.
const DBusSubtreeFlags_none = DBusSubtreeFlags(0)

// DBusSubtreeFlags_dispatch_to_unenumerated_nodes is a representation of the C bitfield member G_DBUS_SUBTREE_FLAGS_DISPATCH_TO_UNENUMERATED_NODES.
const DBusSubtreeFlags_dispatch_to_unenumerated_nodes = DBusSubtreeFlags(1)

// DriveStartFlags is a representation of the C bitfield GDriveStartFlags.
type DriveStartFlags int

// DriveStartFlags_none is a representation of the C bitfield member G_DRIVE_START_NONE.
const DriveStartFlags_none = DriveStartFlags(0)

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

// IOStreamSpliceFlags is a representation of the C bitfield GIOStreamSpliceFlags.
type IOStreamSpliceFlags int

// IOStreamSpliceFlags_none is a representation of the C bitfield member G_IO_STREAM_SPLICE_NONE.
const IOStreamSpliceFlags_none = IOStreamSpliceFlags(0)

// IOStreamSpliceFlags_close_stream1 is a representation of the C bitfield member G_IO_STREAM_SPLICE_CLOSE_STREAM1.
const IOStreamSpliceFlags_close_stream1 = IOStreamSpliceFlags(1)

// IOStreamSpliceFlags_close_stream2 is a representation of the C bitfield member G_IO_STREAM_SPLICE_CLOSE_STREAM2.
const IOStreamSpliceFlags_close_stream2 = IOStreamSpliceFlags(2)

// IOStreamSpliceFlags_wait_for_both is a representation of the C bitfield member G_IO_STREAM_SPLICE_WAIT_FOR_BOTH.
const IOStreamSpliceFlags_wait_for_both = IOStreamSpliceFlags(4)

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

// SocketMsgFlags is a representation of the C bitfield GSocketMsgFlags.
type SocketMsgFlags int

// SocketMsgFlags_none is a representation of the C bitfield member G_SOCKET_MSG_NONE.
const SocketMsgFlags_none = SocketMsgFlags(0)

// SocketMsgFlags_oob is a representation of the C bitfield member G_SOCKET_MSG_OOB.
const SocketMsgFlags_oob = SocketMsgFlags(1)

// SocketMsgFlags_peek is a representation of the C bitfield member G_SOCKET_MSG_PEEK.
const SocketMsgFlags_peek = SocketMsgFlags(2)

// SocketMsgFlags_dontroute is a representation of the C bitfield member G_SOCKET_MSG_DONTROUTE.
const SocketMsgFlags_dontroute = SocketMsgFlags(4)

// TlsCertificateFlags is a representation of the C bitfield GTlsCertificateFlags.
type TlsCertificateFlags int

// TlsCertificateFlags_unknown_ca is a representation of the C bitfield member G_TLS_CERTIFICATE_UNKNOWN_CA.
const TlsCertificateFlags_unknown_ca = TlsCertificateFlags(1)

// TlsCertificateFlags_bad_identity is a representation of the C bitfield member G_TLS_CERTIFICATE_BAD_IDENTITY.
const TlsCertificateFlags_bad_identity = TlsCertificateFlags(2)

// TlsCertificateFlags_not_activated is a representation of the C bitfield member G_TLS_CERTIFICATE_NOT_ACTIVATED.
const TlsCertificateFlags_not_activated = TlsCertificateFlags(4)

// TlsCertificateFlags_expired is a representation of the C bitfield member G_TLS_CERTIFICATE_EXPIRED.
const TlsCertificateFlags_expired = TlsCertificateFlags(8)

// TlsCertificateFlags_revoked is a representation of the C bitfield member G_TLS_CERTIFICATE_REVOKED.
const TlsCertificateFlags_revoked = TlsCertificateFlags(16)

// TlsCertificateFlags_insecure is a representation of the C bitfield member G_TLS_CERTIFICATE_INSECURE.
const TlsCertificateFlags_insecure = TlsCertificateFlags(32)

// TlsCertificateFlags_generic_error is a representation of the C bitfield member G_TLS_CERTIFICATE_GENERIC_ERROR.
const TlsCertificateFlags_generic_error = TlsCertificateFlags(64)

// TlsCertificateFlags_validate_all is a representation of the C bitfield member G_TLS_CERTIFICATE_VALIDATE_ALL.
const TlsCertificateFlags_validate_all = TlsCertificateFlags(127)

// TlsDatabaseVerifyFlags is a representation of the C bitfield GTlsDatabaseVerifyFlags.
type TlsDatabaseVerifyFlags int

// TlsDatabaseVerifyFlags_none is a representation of the C bitfield member G_TLS_DATABASE_VERIFY_NONE.
const TlsDatabaseVerifyFlags_none = TlsDatabaseVerifyFlags(0)

// TlsPasswordFlags is a representation of the C bitfield GTlsPasswordFlags.
type TlsPasswordFlags int

// TlsPasswordFlags_none is a representation of the C bitfield member G_TLS_PASSWORD_NONE.
const TlsPasswordFlags_none = TlsPasswordFlags(0)

// TlsPasswordFlags_retry is a representation of the C bitfield member G_TLS_PASSWORD_RETRY.
const TlsPasswordFlags_retry = TlsPasswordFlags(2)

// TlsPasswordFlags_many_tries is a representation of the C bitfield member G_TLS_PASSWORD_MANY_TRIES.
const TlsPasswordFlags_many_tries = TlsPasswordFlags(4)

// TlsPasswordFlags_final_try is a representation of the C bitfield member G_TLS_PASSWORD_FINAL_TRY.
const TlsPasswordFlags_final_try = TlsPasswordFlags(8)

// BusType is a representation of the C enumeration GBusType.
type BusType int

// BusType_starter is a representation of the C enumeration member G_BUS_TYPE_STARTER.
const BusType_starter = BusType(-1)

// BusType_none is a representation of the C enumeration member G_BUS_TYPE_NONE.
const BusType_none = BusType(0)

// BusType_system is a representation of the C enumeration member G_BUS_TYPE_SYSTEM.
const BusType_system = BusType(1)

// BusType_session is a representation of the C enumeration member G_BUS_TYPE_SESSION.
const BusType_session = BusType(2)

// ConverterResult is a representation of the C enumeration GConverterResult.
type ConverterResult int

// ConverterResult_error is a representation of the C enumeration member G_CONVERTER_ERROR.
const ConverterResult_error = ConverterResult(0)

// ConverterResult_converted is a representation of the C enumeration member G_CONVERTER_CONVERTED.
const ConverterResult_converted = ConverterResult(1)

// ConverterResult_finished is a representation of the C enumeration member G_CONVERTER_FINISHED.
const ConverterResult_finished = ConverterResult(2)

// ConverterResult_flushed is a representation of the C enumeration member G_CONVERTER_FLUSHED.
const ConverterResult_flushed = ConverterResult(3)

// CredentialsType is a representation of the C enumeration GCredentialsType.
type CredentialsType int

// CredentialsType_invalid is a representation of the C enumeration member G_CREDENTIALS_TYPE_INVALID.
const CredentialsType_invalid = CredentialsType(0)

// CredentialsType_linux_ucred is a representation of the C enumeration member G_CREDENTIALS_TYPE_LINUX_UCRED.
const CredentialsType_linux_ucred = CredentialsType(1)

// CredentialsType_freebsd_cmsgcred is a representation of the C enumeration member G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
const CredentialsType_freebsd_cmsgcred = CredentialsType(2)

// CredentialsType_openbsd_sockpeercred is a representation of the C enumeration member G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
const CredentialsType_openbsd_sockpeercred = CredentialsType(3)

// CredentialsType_solaris_ucred is a representation of the C enumeration member G_CREDENTIALS_TYPE_SOLARIS_UCRED.
const CredentialsType_solaris_ucred = CredentialsType(4)

// CredentialsType_netbsd_unpcbid is a representation of the C enumeration member G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
const CredentialsType_netbsd_unpcbid = CredentialsType(5)

// DBusError is a representation of the C enumeration GDBusError.
type DBusError int

// DBusError_failed is a representation of the C enumeration member G_DBUS_ERROR_FAILED.
const DBusError_failed = DBusError(0)

// DBusError_no_memory is a representation of the C enumeration member G_DBUS_ERROR_NO_MEMORY.
const DBusError_no_memory = DBusError(1)

// DBusError_service_unknown is a representation of the C enumeration member G_DBUS_ERROR_SERVICE_UNKNOWN.
const DBusError_service_unknown = DBusError(2)

// DBusError_name_has_no_owner is a representation of the C enumeration member G_DBUS_ERROR_NAME_HAS_NO_OWNER.
const DBusError_name_has_no_owner = DBusError(3)

// DBusError_no_reply is a representation of the C enumeration member G_DBUS_ERROR_NO_REPLY.
const DBusError_no_reply = DBusError(4)

// DBusError_io_error is a representation of the C enumeration member G_DBUS_ERROR_IO_ERROR.
const DBusError_io_error = DBusError(5)

// DBusError_bad_address is a representation of the C enumeration member G_DBUS_ERROR_BAD_ADDRESS.
const DBusError_bad_address = DBusError(6)

// DBusError_not_supported is a representation of the C enumeration member G_DBUS_ERROR_NOT_SUPPORTED.
const DBusError_not_supported = DBusError(7)

// DBusError_limits_exceeded is a representation of the C enumeration member G_DBUS_ERROR_LIMITS_EXCEEDED.
const DBusError_limits_exceeded = DBusError(8)

// DBusError_access_denied is a representation of the C enumeration member G_DBUS_ERROR_ACCESS_DENIED.
const DBusError_access_denied = DBusError(9)

// DBusError_auth_failed is a representation of the C enumeration member G_DBUS_ERROR_AUTH_FAILED.
const DBusError_auth_failed = DBusError(10)

// DBusError_no_server is a representation of the C enumeration member G_DBUS_ERROR_NO_SERVER.
const DBusError_no_server = DBusError(11)

// DBusError_timeout is a representation of the C enumeration member G_DBUS_ERROR_TIMEOUT.
const DBusError_timeout = DBusError(12)

// DBusError_no_network is a representation of the C enumeration member G_DBUS_ERROR_NO_NETWORK.
const DBusError_no_network = DBusError(13)

// DBusError_address_in_use is a representation of the C enumeration member G_DBUS_ERROR_ADDRESS_IN_USE.
const DBusError_address_in_use = DBusError(14)

// DBusError_disconnected is a representation of the C enumeration member G_DBUS_ERROR_DISCONNECTED.
const DBusError_disconnected = DBusError(15)

// DBusError_invalid_args is a representation of the C enumeration member G_DBUS_ERROR_INVALID_ARGS.
const DBusError_invalid_args = DBusError(16)

// DBusError_file_not_found is a representation of the C enumeration member G_DBUS_ERROR_FILE_NOT_FOUND.
const DBusError_file_not_found = DBusError(17)

// DBusError_file_exists is a representation of the C enumeration member G_DBUS_ERROR_FILE_EXISTS.
const DBusError_file_exists = DBusError(18)

// DBusError_unknown_method is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_METHOD.
const DBusError_unknown_method = DBusError(19)

// DBusError_timed_out is a representation of the C enumeration member G_DBUS_ERROR_TIMED_OUT.
const DBusError_timed_out = DBusError(20)

// DBusError_match_rule_not_found is a representation of the C enumeration member G_DBUS_ERROR_MATCH_RULE_NOT_FOUND.
const DBusError_match_rule_not_found = DBusError(21)

// DBusError_match_rule_invalid is a representation of the C enumeration member G_DBUS_ERROR_MATCH_RULE_INVALID.
const DBusError_match_rule_invalid = DBusError(22)

// DBusError_spawn_exec_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_EXEC_FAILED.
const DBusError_spawn_exec_failed = DBusError(23)

// DBusError_spawn_fork_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FORK_FAILED.
const DBusError_spawn_fork_failed = DBusError(24)

// DBusError_spawn_child_exited is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CHILD_EXITED.
const DBusError_spawn_child_exited = DBusError(25)

// DBusError_spawn_child_signaled is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CHILD_SIGNALED.
const DBusError_spawn_child_signaled = DBusError(26)

// DBusError_spawn_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FAILED.
const DBusError_spawn_failed = DBusError(27)

// DBusError_spawn_setup_failed is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SETUP_FAILED.
const DBusError_spawn_setup_failed = DBusError(28)

// DBusError_spawn_config_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_CONFIG_INVALID.
const DBusError_spawn_config_invalid = DBusError(29)

// DBusError_spawn_service_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SERVICE_INVALID.
const DBusError_spawn_service_invalid = DBusError(30)

// DBusError_spawn_service_not_found is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_SERVICE_NOT_FOUND.
const DBusError_spawn_service_not_found = DBusError(31)

// DBusError_spawn_permissions_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_PERMISSIONS_INVALID.
const DBusError_spawn_permissions_invalid = DBusError(32)

// DBusError_spawn_file_invalid is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_FILE_INVALID.
const DBusError_spawn_file_invalid = DBusError(33)

// DBusError_spawn_no_memory is a representation of the C enumeration member G_DBUS_ERROR_SPAWN_NO_MEMORY.
const DBusError_spawn_no_memory = DBusError(34)

// DBusError_unix_process_id_unknown is a representation of the C enumeration member G_DBUS_ERROR_UNIX_PROCESS_ID_UNKNOWN.
const DBusError_unix_process_id_unknown = DBusError(35)

// DBusError_invalid_signature is a representation of the C enumeration member G_DBUS_ERROR_INVALID_SIGNATURE.
const DBusError_invalid_signature = DBusError(36)

// DBusError_invalid_file_content is a representation of the C enumeration member G_DBUS_ERROR_INVALID_FILE_CONTENT.
const DBusError_invalid_file_content = DBusError(37)

// DBusError_selinux_security_context_unknown is a representation of the C enumeration member G_DBUS_ERROR_SELINUX_SECURITY_CONTEXT_UNKNOWN.
const DBusError_selinux_security_context_unknown = DBusError(38)

// DBusError_adt_audit_data_unknown is a representation of the C enumeration member G_DBUS_ERROR_ADT_AUDIT_DATA_UNKNOWN.
const DBusError_adt_audit_data_unknown = DBusError(39)

// DBusError_object_path_in_use is a representation of the C enumeration member G_DBUS_ERROR_OBJECT_PATH_IN_USE.
const DBusError_object_path_in_use = DBusError(40)

// DBusError_unknown_object is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_OBJECT.
const DBusError_unknown_object = DBusError(41)

// DBusError_unknown_interface is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_INTERFACE.
const DBusError_unknown_interface = DBusError(42)

// DBusError_unknown_property is a representation of the C enumeration member G_DBUS_ERROR_UNKNOWN_PROPERTY.
const DBusError_unknown_property = DBusError(43)

// DBusError_property_read_only is a representation of the C enumeration member G_DBUS_ERROR_PROPERTY_READ_ONLY.
const DBusError_property_read_only = DBusError(44)

// DBusMessageByteOrder is a representation of the C enumeration GDBusMessageByteOrder.
type DBusMessageByteOrder int

// DBusMessageByteOrder_big_endian is a representation of the C enumeration member G_DBUS_MESSAGE_BYTE_ORDER_BIG_ENDIAN.
const DBusMessageByteOrder_big_endian = DBusMessageByteOrder(66)

// DBusMessageByteOrder_little_endian is a representation of the C enumeration member G_DBUS_MESSAGE_BYTE_ORDER_LITTLE_ENDIAN.
const DBusMessageByteOrder_little_endian = DBusMessageByteOrder(108)

// DBusMessageHeaderField is a representation of the C enumeration GDBusMessageHeaderField.
type DBusMessageHeaderField int

// DBusMessageHeaderField_invalid is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_INVALID.
const DBusMessageHeaderField_invalid = DBusMessageHeaderField(0)

// DBusMessageHeaderField_path is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_PATH.
const DBusMessageHeaderField_path = DBusMessageHeaderField(1)

// DBusMessageHeaderField_interface is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_INTERFACE.
const DBusMessageHeaderField_interface = DBusMessageHeaderField(2)

// DBusMessageHeaderField_member is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_MEMBER.
const DBusMessageHeaderField_member = DBusMessageHeaderField(3)

// DBusMessageHeaderField_error_name is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_ERROR_NAME.
const DBusMessageHeaderField_error_name = DBusMessageHeaderField(4)

// DBusMessageHeaderField_reply_serial is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_REPLY_SERIAL.
const DBusMessageHeaderField_reply_serial = DBusMessageHeaderField(5)

// DBusMessageHeaderField_destination is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_DESTINATION.
const DBusMessageHeaderField_destination = DBusMessageHeaderField(6)

// DBusMessageHeaderField_sender is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_SENDER.
const DBusMessageHeaderField_sender = DBusMessageHeaderField(7)

// DBusMessageHeaderField_signature is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_SIGNATURE.
const DBusMessageHeaderField_signature = DBusMessageHeaderField(8)

// DBusMessageHeaderField_num_unix_fds is a representation of the C enumeration member G_DBUS_MESSAGE_HEADER_FIELD_NUM_UNIX_FDS.
const DBusMessageHeaderField_num_unix_fds = DBusMessageHeaderField(9)

// DBusMessageType is a representation of the C enumeration GDBusMessageType.
type DBusMessageType int

// DBusMessageType_invalid is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_INVALID.
const DBusMessageType_invalid = DBusMessageType(0)

// DBusMessageType_method_call is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_METHOD_CALL.
const DBusMessageType_method_call = DBusMessageType(1)

// DBusMessageType_method_return is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_METHOD_RETURN.
const DBusMessageType_method_return = DBusMessageType(2)

// DBusMessageType_error is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_ERROR.
const DBusMessageType_error = DBusMessageType(3)

// DBusMessageType_signal is a representation of the C enumeration member G_DBUS_MESSAGE_TYPE_SIGNAL.
const DBusMessageType_signal = DBusMessageType(4)

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

// DriveStartStopType is a representation of the C enumeration GDriveStartStopType.
type DriveStartStopType int

// DriveStartStopType_unknown is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_UNKNOWN.
const DriveStartStopType_unknown = DriveStartStopType(0)

// DriveStartStopType_shutdown is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_SHUTDOWN.
const DriveStartStopType_shutdown = DriveStartStopType(1)

// DriveStartStopType_network is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_NETWORK.
const DriveStartStopType_network = DriveStartStopType(2)

// DriveStartStopType_multidisk is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_MULTIDISK.
const DriveStartStopType_multidisk = DriveStartStopType(3)

// DriveStartStopType_password is a representation of the C enumeration member G_DRIVE_START_STOP_TYPE_PASSWORD.
const DriveStartStopType_password = DriveStartStopType(4)

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

// IOModuleScopeFlags is a representation of the C enumeration GIOModuleScopeFlags.
type IOModuleScopeFlags int

// IOModuleScopeFlags_none is a representation of the C enumeration member G_IO_MODULE_SCOPE_NONE.
const IOModuleScopeFlags_none = IOModuleScopeFlags(0)

// IOModuleScopeFlags_block_duplicates is a representation of the C enumeration member G_IO_MODULE_SCOPE_BLOCK_DUPLICATES.
const IOModuleScopeFlags_block_duplicates = IOModuleScopeFlags(1)

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

// ResolverError is a representation of the C enumeration GResolverError.
type ResolverError int

// ResolverError_not_found is a representation of the C enumeration member G_RESOLVER_ERROR_NOT_FOUND.
const ResolverError_not_found = ResolverError(0)

// ResolverError_temporary_failure is a representation of the C enumeration member G_RESOLVER_ERROR_TEMPORARY_FAILURE.
const ResolverError_temporary_failure = ResolverError(1)

// ResolverError_internal is a representation of the C enumeration member G_RESOLVER_ERROR_INTERNAL.
const ResolverError_internal = ResolverError(2)

// SocketFamily is a representation of the C enumeration GSocketFamily.
type SocketFamily int

// SocketFamily_invalid is a representation of the C enumeration member G_SOCKET_FAMILY_INVALID.
const SocketFamily_invalid = SocketFamily(0)

// SocketFamily_unix is a representation of the C enumeration member G_SOCKET_FAMILY_UNIX.
const SocketFamily_unix = SocketFamily(1)

// SocketFamily_ipv4 is a representation of the C enumeration member G_SOCKET_FAMILY_IPV4.
const SocketFamily_ipv4 = SocketFamily(2)

// SocketFamily_ipv6 is a representation of the C enumeration member G_SOCKET_FAMILY_IPV6.
const SocketFamily_ipv6 = SocketFamily(10)

// SocketProtocol is a representation of the C enumeration GSocketProtocol.
type SocketProtocol int

// SocketProtocol_unknown is a representation of the C enumeration member G_SOCKET_PROTOCOL_UNKNOWN.
const SocketProtocol_unknown = SocketProtocol(-1)

// SocketProtocol_default is a representation of the C enumeration member G_SOCKET_PROTOCOL_DEFAULT.
const SocketProtocol_default = SocketProtocol(0)

// SocketProtocol_tcp is a representation of the C enumeration member G_SOCKET_PROTOCOL_TCP.
const SocketProtocol_tcp = SocketProtocol(6)

// SocketProtocol_udp is a representation of the C enumeration member G_SOCKET_PROTOCOL_UDP.
const SocketProtocol_udp = SocketProtocol(17)

// SocketProtocol_sctp is a representation of the C enumeration member G_SOCKET_PROTOCOL_SCTP.
const SocketProtocol_sctp = SocketProtocol(132)

// SocketType is a representation of the C enumeration GSocketType.
type SocketType int

// SocketType_invalid is a representation of the C enumeration member G_SOCKET_TYPE_INVALID.
const SocketType_invalid = SocketType(0)

// SocketType_stream is a representation of the C enumeration member G_SOCKET_TYPE_STREAM.
const SocketType_stream = SocketType(1)

// SocketType_datagram is a representation of the C enumeration member G_SOCKET_TYPE_DATAGRAM.
const SocketType_datagram = SocketType(2)

// SocketType_seqpacket is a representation of the C enumeration member G_SOCKET_TYPE_SEQPACKET.
const SocketType_seqpacket = SocketType(3)

// TlsAuthenticationMode is a representation of the C enumeration GTlsAuthenticationMode.
type TlsAuthenticationMode int

// TlsAuthenticationMode_none is a representation of the C enumeration member G_TLS_AUTHENTICATION_NONE.
const TlsAuthenticationMode_none = TlsAuthenticationMode(0)

// TlsAuthenticationMode_requested is a representation of the C enumeration member G_TLS_AUTHENTICATION_REQUESTED.
const TlsAuthenticationMode_requested = TlsAuthenticationMode(1)

// TlsAuthenticationMode_required is a representation of the C enumeration member G_TLS_AUTHENTICATION_REQUIRED.
const TlsAuthenticationMode_required = TlsAuthenticationMode(2)

// TlsDatabaseLookupFlags is a representation of the C enumeration GTlsDatabaseLookupFlags.
type TlsDatabaseLookupFlags int

// TlsDatabaseLookupFlags_none is a representation of the C enumeration member G_TLS_DATABASE_LOOKUP_NONE.
const TlsDatabaseLookupFlags_none = TlsDatabaseLookupFlags(0)

// TlsDatabaseLookupFlags_keypair is a representation of the C enumeration member G_TLS_DATABASE_LOOKUP_KEYPAIR.
const TlsDatabaseLookupFlags_keypair = TlsDatabaseLookupFlags(1)

// TlsError is a representation of the C enumeration GTlsError.
type TlsError int

// TlsError_unavailable is a representation of the C enumeration member G_TLS_ERROR_UNAVAILABLE.
const TlsError_unavailable = TlsError(0)

// TlsError_misc is a representation of the C enumeration member G_TLS_ERROR_MISC.
const TlsError_misc = TlsError(1)

// TlsError_bad_certificate is a representation of the C enumeration member G_TLS_ERROR_BAD_CERTIFICATE.
const TlsError_bad_certificate = TlsError(2)

// TlsError_not_tls is a representation of the C enumeration member G_TLS_ERROR_NOT_TLS.
const TlsError_not_tls = TlsError(3)

// TlsError_handshake is a representation of the C enumeration member G_TLS_ERROR_HANDSHAKE.
const TlsError_handshake = TlsError(4)

// TlsError_certificate_required is a representation of the C enumeration member G_TLS_ERROR_CERTIFICATE_REQUIRED.
const TlsError_certificate_required = TlsError(5)

// TlsError_eof is a representation of the C enumeration member G_TLS_ERROR_EOF.
const TlsError_eof = TlsError(6)

// TlsError_inappropriate_fallback is a representation of the C enumeration member G_TLS_ERROR_INAPPROPRIATE_FALLBACK.
const TlsError_inappropriate_fallback = TlsError(7)

// TlsInteractionResult is a representation of the C enumeration GTlsInteractionResult.
type TlsInteractionResult int

// TlsInteractionResult_unhandled is a representation of the C enumeration member G_TLS_INTERACTION_UNHANDLED.
const TlsInteractionResult_unhandled = TlsInteractionResult(0)

// TlsInteractionResult_handled is a representation of the C enumeration member G_TLS_INTERACTION_HANDLED.
const TlsInteractionResult_handled = TlsInteractionResult(1)

// TlsInteractionResult_failed is a representation of the C enumeration member G_TLS_INTERACTION_FAILED.
const TlsInteractionResult_failed = TlsInteractionResult(2)

// TlsRehandshakeMode is a representation of the C enumeration GTlsRehandshakeMode.
type TlsRehandshakeMode int

// TlsRehandshakeMode_never is a representation of the C enumeration member G_TLS_REHANDSHAKE_NEVER.
const TlsRehandshakeMode_never = TlsRehandshakeMode(0)

// TlsRehandshakeMode_safely is a representation of the C enumeration member G_TLS_REHANDSHAKE_SAFELY.
const TlsRehandshakeMode_safely = TlsRehandshakeMode(1)

// TlsRehandshakeMode_unsafely is a representation of the C enumeration member G_TLS_REHANDSHAKE_UNSAFELY.
const TlsRehandshakeMode_unsafely = TlsRehandshakeMode(2)

// UnixSocketAddressType is a representation of the C enumeration GUnixSocketAddressType.
type UnixSocketAddressType int

// UnixSocketAddressType_invalid is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_INVALID.
const UnixSocketAddressType_invalid = UnixSocketAddressType(0)

// UnixSocketAddressType_anonymous is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ANONYMOUS.
const UnixSocketAddressType_anonymous = UnixSocketAddressType(1)

// UnixSocketAddressType_path is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_PATH.
const UnixSocketAddressType_path = UnixSocketAddressType(2)

// UnixSocketAddressType_abstract is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ABSTRACT.
const UnixSocketAddressType_abstract = UnixSocketAddressType(3)

// UnixSocketAddressType_abstract_padded is a representation of the C enumeration member G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED.
const UnixSocketAddressType_abstract_padded = UnixSocketAddressType(4)

// ZlibCompressorFormat is a representation of the C enumeration GZlibCompressorFormat.
type ZlibCompressorFormat int

// ZlibCompressorFormat_zlib is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_ZLIB.
const ZlibCompressorFormat_zlib = ZlibCompressorFormat(0)

// ZlibCompressorFormat_gzip is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_GZIP.
const ZlibCompressorFormat_gzip = ZlibCompressorFormat(1)

// ZlibCompressorFormat_raw is a representation of the C enumeration member G_ZLIB_COMPRESSOR_FORMAT_RAW.
const ZlibCompressorFormat_raw = ZlibCompressorFormat(2)

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

func Fn_g_bus_get_finish(res *AsyncResult) {}

func Fn_g_bus_get_sync(busType int, cancellable *Cancellable) {}

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

func Fn_g_bus_own_name_on_connection_with_closures(connection *DBusConnection, name string, flags int, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) {
}

func Fn_g_bus_own_name_with_closures(busType int, name string, flags int, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) {
}

func Fn_g_bus_unown_name(ownerId uint) {}

func Fn_g_bus_unwatch_name(watcherId uint) {}

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

func Fn_g_bus_watch_name_on_connection_with_closures(connection *DBusConnection, name string, flags int, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) {
}

func Fn_g_bus_watch_name_with_closures(busType int, name string, flags int, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) {
}

func Fn_g_content_type_can_be_executable(type_ string) {}

func Fn_g_content_type_equals(type1 string, type2 string) {}

func Fn_g_content_type_from_mime_type(mimeType string) {}

func Fn_g_content_type_get_description(type_ string) {}

func Fn_g_content_type_get_icon(type_ string) {}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

func Fn_g_content_type_get_mime_type(type_ string) {}

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

func Fn_g_content_type_is_a(type_ string, supertype string) {}

func Fn_g_content_type_is_unknown(type_ string) {}

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

func Fn_g_content_types_get_registered() {}

func Fn_g_dbus_address_get_for_bus_sync(busType int, cancellable *Cancellable) {}

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

func Fn_g_dbus_address_get_stream_finish(res *AsyncResult) {}

func Fn_g_dbus_address_get_stream_sync(address string, cancellable *Cancellable) {}

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

func Fn_g_dbus_generate_guid() {}

func Fn_g_dbus_gvalue_to_gvariant(gvalue *gobject.Value, type_ *glib.VariantType) {}

func Fn_g_dbus_gvariant_to_gvalue(value *glib.Variant) {}

func Fn_g_dbus_is_address(string_ string) {}

func Fn_g_dbus_is_guid(string_ string) {}

func Fn_g_dbus_is_interface_name(string_ string) {}

func Fn_g_dbus_is_member_name(string_ string) {}

func Fn_g_dbus_is_name(string_ string) {}

func Fn_g_dbus_is_supported_address(string_ string) {}

func Fn_g_dbus_is_unique_name(string_ string) {}

// UNSUPPORTED : g_initable_newv : has array param, parameters

func Fn_g_io_error_from_errno(errNo int) {}

func Fn_g_io_error_quark() {}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

func Fn_g_io_scheduler_cancel_all_jobs() {}

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

func Fn_g_pollable_source_new(pollableStream *gobject.Object) {}

// UNSUPPORTED : g_pollable_stream_read : has array param, buffer

// UNSUPPORTED : g_pollable_stream_write : has array param, buffer

// UNSUPPORTED : g_pollable_stream_write_all : has array param, buffer

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

func Fn_g_unix_is_mount_path_system_internal(mountPath string) {}

func Fn_g_unix_mount_at(mountPath string) {}

func Fn_g_unix_mount_compare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) {}

func Fn_g_unix_mount_free(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_get_device_path(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_get_fs_type(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_get_mount_path(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_guess_can_eject(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_guess_icon(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_guess_name(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_guess_should_display(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_is_readonly(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_is_system_internal(mountEntry *UnixMountEntry) {}

func Fn_g_unix_mount_points_changed_since(time uint64) {}

func Fn_g_unix_mount_points_get() {}

func Fn_g_unix_mounts_changed_since(time uint64) {}

func Fn_g_unix_mounts_get() {}

// ActionEntry is a representation of the C record GActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ActionGroupInterface is a representation of the C record GActionGroupInterface.
//
// since 2.28
type ActionGroupInterface struct {
	native unsafe.Pointer
}

// ActionInterface is a representation of the C record GActionInterface.
//
// since 2.28
type ActionInterface struct {
	native unsafe.Pointer
}

// AppInfoIface is a representation of the C record GAppInfoIface.
type AppInfoIface struct {
	native unsafe.Pointer
}

// AppLaunchContextClass is a representation of the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native unsafe.Pointer
}

// AppLaunchContextPrivate is a representation of the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

// ApplicationClass is a representation of the C record GApplicationClass.
//
// since 2.28
type ApplicationClass struct {
	native unsafe.Pointer
}

// ApplicationCommandLineClass is a representation of the C record GApplicationCommandLineClass.
//
// since 2.28
type ApplicationCommandLineClass struct {
	native unsafe.Pointer
}

// ApplicationCommandLinePrivate is a representation of the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

// ApplicationPrivate is a representation of the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// AsyncInitableIface is a representation of the C record GAsyncInitableIface.
//
// since 2.22
type AsyncInitableIface struct {
	native unsafe.Pointer
}

// AsyncResultIface is a representation of the C record GAsyncResultIface.
type AsyncResultIface struct {
	native unsafe.Pointer
}

// BufferedInputStreamClass is a representation of the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native unsafe.Pointer
}

// BufferedInputStreamPrivate is a representation of the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

// BufferedOutputStreamClass is a representation of the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native unsafe.Pointer
}

// BufferedOutputStreamPrivate is a representation of the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

// CancellableClass is a representation of the C record GCancellableClass.
type CancellableClass struct {
	native unsafe.Pointer
}

// CancellablePrivate is a representation of the C record GCancellablePrivate.
type CancellablePrivate struct {
	native unsafe.Pointer
}

// CharsetConverterClass is a representation of the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native unsafe.Pointer
}

// ConverterIface is a representation of the C record GConverterIface.
//
// since 2.24
type ConverterIface struct {
	native unsafe.Pointer
}

// ConverterInputStreamClass is a representation of the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native unsafe.Pointer
}

// ConverterInputStreamPrivate is a representation of the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

// ConverterOutputStreamClass is a representation of the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native unsafe.Pointer
}

// ConverterOutputStreamPrivate is a representation of the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

// CredentialsClass is a representation of the C record GCredentialsClass.
//
// since 2.26
type CredentialsClass struct {
	native unsafe.Pointer
}

// DBusAnnotationInfo is a representation of the C record GDBusAnnotationInfo.
//
// since 2.26
type DBusAnnotationInfo struct {
	native unsafe.Pointer
}

// DBusArgInfo is a representation of the C record GDBusArgInfo.
//
// since 2.26
type DBusArgInfo struct {
	native unsafe.Pointer
}

// DBusErrorEntry is a representation of the C record GDBusErrorEntry.
//
// since 2.26
type DBusErrorEntry struct {
	native unsafe.Pointer
}

// DBusInterfaceIface is a representation of the C record GDBusInterfaceIface.
//
// since 2.30
type DBusInterfaceIface struct {
	native unsafe.Pointer
}

// DBusInterfaceInfo is a representation of the C record GDBusInterfaceInfo.
//
// since 2.26
type DBusInterfaceInfo struct {
	native unsafe.Pointer
}

// DBusInterfaceSkeletonClass is a representation of the C record GDBusInterfaceSkeletonClass.
//
// since 2.30
type DBusInterfaceSkeletonClass struct {
	native unsafe.Pointer
}

// DBusInterfaceSkeletonPrivate is a representation of the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

// DBusInterfaceVTable is a representation of the C record GDBusInterfaceVTable.
//
// since 2.26
type DBusInterfaceVTable struct {
	native unsafe.Pointer
}

// DBusMethodInfo is a representation of the C record GDBusMethodInfo.
//
// since 2.26
type DBusMethodInfo struct {
	native unsafe.Pointer
}

// DBusNodeInfo is a representation of the C record GDBusNodeInfo.
//
// since 2.26
type DBusNodeInfo struct {
	native unsafe.Pointer
}

// DBusObjectIface is a representation of the C record GDBusObjectIface.
//
// since 2.30
type DBusObjectIface struct {
	native unsafe.Pointer
}

// DBusObjectManagerClientClass is a representation of the C record GDBusObjectManagerClientClass.
//
// since 2.30
type DBusObjectManagerClientClass struct {
	native unsafe.Pointer
}

// DBusObjectManagerClientPrivate is a representation of the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

// DBusObjectManagerIface is a representation of the C record GDBusObjectManagerIface.
//
// since 2.30
type DBusObjectManagerIface struct {
	native unsafe.Pointer
}

// DBusObjectManagerServerClass is a representation of the C record GDBusObjectManagerServerClass.
//
// since 2.30
type DBusObjectManagerServerClass struct {
	native unsafe.Pointer
}

// DBusObjectManagerServerPrivate is a representation of the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

// DBusObjectProxyClass is a representation of the C record GDBusObjectProxyClass.
//
// since 2.30
type DBusObjectProxyClass struct {
	native unsafe.Pointer
}

// DBusObjectProxyPrivate is a representation of the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

// DBusObjectSkeletonClass is a representation of the C record GDBusObjectSkeletonClass.
//
// since 2.30
type DBusObjectSkeletonClass struct {
	native unsafe.Pointer
}

// DBusObjectSkeletonPrivate is a representation of the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

// DBusPropertyInfo is a representation of the C record GDBusPropertyInfo.
//
// since 2.26
type DBusPropertyInfo struct {
	native unsafe.Pointer
}

// DBusProxyClass is a representation of the C record GDBusProxyClass.
//
// since 2.26
type DBusProxyClass struct {
	native unsafe.Pointer
}

// DBusProxyPrivate is a representation of the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native unsafe.Pointer
}

// DBusSignalInfo is a representation of the C record GDBusSignalInfo.
//
// since 2.26
type DBusSignalInfo struct {
	native unsafe.Pointer
}

// DBusSubtreeVTable is a representation of the C record GDBusSubtreeVTable.
//
// since 2.26
type DBusSubtreeVTable struct {
	native unsafe.Pointer
}

// DataInputStreamClass is a representation of the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native unsafe.Pointer
}

// DataInputStreamPrivate is a representation of the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

// DataOutputStreamClass is a representation of the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native unsafe.Pointer
}

// DataOutputStreamPrivate is a representation of the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

// DesktopAppInfoClass is a representation of the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native unsafe.Pointer
}

// DesktopAppInfoLookupIface is a representation of the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
}

// DriveIface is a representation of the C record GDriveIface.
type DriveIface struct {
	native unsafe.Pointer
}

// EmblemClass is a representation of the C record GEmblemClass.
type EmblemClass struct {
	native unsafe.Pointer
}

// EmblemedIconClass is a representation of the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native unsafe.Pointer
}

// EmblemedIconPrivate is a representation of the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

// FileAttributeInfo is a representation of the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native unsafe.Pointer
}

// FileAttributeInfoList is a representation of the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native unsafe.Pointer
}

// FileAttributeMatcher is a representation of the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native unsafe.Pointer
}

// FileDescriptorBasedIface is a representation of the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native unsafe.Pointer
}

// FileEnumeratorClass is a representation of the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native unsafe.Pointer
}

// FileEnumeratorPrivate is a representation of the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

// FileIOStreamClass is a representation of the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native unsafe.Pointer
}

// FileIOStreamPrivate is a representation of the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

// FileIconClass is a representation of the C record GFileIconClass.
type FileIconClass struct {
	native unsafe.Pointer
}

// FileIface is a representation of the C record GFileIface.
type FileIface struct {
	native unsafe.Pointer
}

// FileInfoClass is a representation of the C record GFileInfoClass.
type FileInfoClass struct {
	native unsafe.Pointer
}

// FileInputStreamClass is a representation of the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native unsafe.Pointer
}

// FileInputStreamPrivate is a representation of the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

// FileMonitorClass is a representation of the C record GFileMonitorClass.
type FileMonitorClass struct {
	native unsafe.Pointer
}

// FileMonitorPrivate is a representation of the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native unsafe.Pointer
}

// FileOutputStreamClass is a representation of the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native unsafe.Pointer
}

// FileOutputStreamPrivate is a representation of the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

// FilenameCompleterClass is a representation of the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native unsafe.Pointer
}

// FilterInputStreamClass is a representation of the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native unsafe.Pointer
}

// FilterOutputStreamClass is a representation of the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native unsafe.Pointer
}

// IOExtension is a representation of the C record GIOExtension.
type IOExtension struct {
	native unsafe.Pointer
}

// IOExtensionPoint is a representation of the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native unsafe.Pointer
}

// IOModuleClass is a representation of the C record GIOModuleClass.
type IOModuleClass struct {
	native unsafe.Pointer
}

// IOModuleScope is a representation of the C record GIOModuleScope.
//
// since 2.30
type IOModuleScope struct {
	native unsafe.Pointer
}

// IOSchedulerJob is a representation of the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native unsafe.Pointer
}

// IOStreamAdapter is a representation of the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native unsafe.Pointer
}

// IOStreamClass is a representation of the C record GIOStreamClass.
type IOStreamClass struct {
	native unsafe.Pointer
}

// IOStreamPrivate is a representation of the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native unsafe.Pointer
}

// IconIface is a representation of the C record GIconIface.
type IconIface struct {
	native unsafe.Pointer
}

// InetAddressClass is a representation of the C record GInetAddressClass.
type InetAddressClass struct {
	native unsafe.Pointer
}

// InetAddressMaskClass is a representation of the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native unsafe.Pointer
}

// InetAddressMaskPrivate is a representation of the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

// InetAddressPrivate is a representation of the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native unsafe.Pointer
}

// InetSocketAddressClass is a representation of the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native unsafe.Pointer
}

// InetSocketAddressPrivate is a representation of the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

// InitableIface is a representation of the C record GInitableIface.
//
// since 2.22
type InitableIface struct {
	native unsafe.Pointer
}

// InputStreamClass is a representation of the C record GInputStreamClass.
type InputStreamClass struct {
	native unsafe.Pointer
}

// InputStreamPrivate is a representation of the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native unsafe.Pointer
}

// InputVector is a representation of the C record GInputVector.
//
// since 2.22
type InputVector struct {
	native unsafe.Pointer
}

// ListStoreClass is a representation of the C record GListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// LoadableIconIface is a representation of the C record GLoadableIconIface.
type LoadableIconIface struct {
	native unsafe.Pointer
}

// MemoryInputStreamClass is a representation of the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native unsafe.Pointer
}

// MemoryInputStreamPrivate is a representation of the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

// MemoryOutputStreamClass is a representation of the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native unsafe.Pointer
}

// MemoryOutputStreamPrivate is a representation of the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

// MenuAttributeIterClass is a representation of the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native unsafe.Pointer
}

// MenuAttributeIterPrivate is a representation of the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

// MenuLinkIterClass is a representation of the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native unsafe.Pointer
}

// MenuLinkIterPrivate is a representation of the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

// MenuModelClass is a representation of the C record GMenuModelClass.
type MenuModelClass struct {
	native unsafe.Pointer
}

// MenuModelPrivate is a representation of the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native unsafe.Pointer
}

// MountIface is a representation of the C record GMountIface.
type MountIface struct {
	native unsafe.Pointer
}

// MountOperationClass is a representation of the C record GMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// MountOperationPrivate is a representation of the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// NativeVolumeMonitorClass is a representation of the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native unsafe.Pointer
}

// NetworkAddressClass is a representation of the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native unsafe.Pointer
}

// NetworkAddressPrivate is a representation of the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

// NetworkServiceClass is a representation of the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native unsafe.Pointer
}

// NetworkServicePrivate is a representation of the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native unsafe.Pointer
}

// OutputStreamClass is a representation of the C record GOutputStreamClass.
type OutputStreamClass struct {
	native unsafe.Pointer
}

// OutputStreamPrivate is a representation of the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native unsafe.Pointer
}

// OutputVector is a representation of the C record GOutputVector.
//
// since 2.22
type OutputVector struct {
	native unsafe.Pointer
}

// PermissionClass is a representation of the C record GPermissionClass.
type PermissionClass struct {
	native unsafe.Pointer
}

// PermissionPrivate is a representation of the C record GPermissionPrivate.
type PermissionPrivate struct {
	native unsafe.Pointer
}

// PollableInputStreamInterface is a representation of the C record GPollableInputStreamInterface.
//
// since 2.28
type PollableInputStreamInterface struct {
	native unsafe.Pointer
}

// PollableOutputStreamInterface is a representation of the C record GPollableOutputStreamInterface.
//
// since 2.28
type PollableOutputStreamInterface struct {
	native unsafe.Pointer
}

// ProxyAddressClass is a representation of the C record GProxyAddressClass.
//
// since 2.26
type ProxyAddressClass struct {
	native unsafe.Pointer
}

// ProxyAddressEnumeratorClass is a representation of the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ProxyAddressEnumeratorPrivate is a representation of the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ProxyAddressPrivate is a representation of the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

// ProxyInterface is a representation of the C record GProxyInterface.
//
// since 2.26
type ProxyInterface struct {
	native unsafe.Pointer
}

// ProxyResolverInterface is a representation of the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native unsafe.Pointer
}

// ResolverClass is a representation of the C record GResolverClass.
type ResolverClass struct {
	native unsafe.Pointer
}

// ResolverPrivate is a representation of the C record GResolverPrivate.
type ResolverPrivate struct {
	native unsafe.Pointer
}

// SeekableIface is a representation of the C record GSeekableIface.
type SeekableIface struct {
	native unsafe.Pointer
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// SettingsClass is a representation of the C record GSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

// SettingsPrivate is a representation of the C record GSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// SettingsSchemaKey is a representation of the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native unsafe.Pointer
}

// SimpleActionGroupClass is a representation of the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native unsafe.Pointer
}

// SimpleActionGroupPrivate is a representation of the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

// SimpleAsyncResultClass is a representation of the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

// SimpleProxyResolverClass is a representation of the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native unsafe.Pointer
}

// SimpleProxyResolverPrivate is a representation of the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

// SocketAddressClass is a representation of the C record GSocketAddressClass.
type SocketAddressClass struct {
	native unsafe.Pointer
}

// SocketAddressEnumeratorClass is a representation of the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// SocketClass is a representation of the C record GSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// SocketClientClass is a representation of the C record GSocketClientClass.
type SocketClientClass struct {
	native unsafe.Pointer
}

// SocketClientPrivate is a representation of the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native unsafe.Pointer
}

// SocketConnectableIface is a representation of the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native unsafe.Pointer
}

// SocketConnectionClass is a representation of the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native unsafe.Pointer
}

// SocketConnectionPrivate is a representation of the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

// SocketControlMessageClass is a representation of the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native unsafe.Pointer
}

// SocketControlMessagePrivate is a representation of the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

// SocketListenerClass is a representation of the C record GSocketListenerClass.
type SocketListenerClass struct {
	native unsafe.Pointer
}

// SocketListenerPrivate is a representation of the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native unsafe.Pointer
}

// SocketPrivate is a representation of the C record GSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// SocketServiceClass is a representation of the C record GSocketServiceClass.
type SocketServiceClass struct {
	native unsafe.Pointer
}

// SocketServicePrivate is a representation of the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native unsafe.Pointer
}

// SrvTarget is a representation of the C record GSrvTarget.
type SrvTarget struct {
	native unsafe.Pointer
}

// StaticResource is a representation of the C record GStaticResource.
type StaticResource struct {
	native unsafe.Pointer
}

// TaskClass is a representation of the C record GTaskClass.
type TaskClass struct {
	native unsafe.Pointer
}

// TcpConnectionClass is a representation of the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native unsafe.Pointer
}

// TcpConnectionPrivate is a representation of the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

// TcpWrapperConnectionClass is a representation of the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

// TcpWrapperConnectionPrivate is a representation of the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

// ThemedIconClass is a representation of the C record GThemedIconClass.
type ThemedIconClass struct {
	native unsafe.Pointer
}

// ThreadedSocketServiceClass is a representation of the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
}

// ThreadedSocketServicePrivate is a representation of the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

// TlsBackendInterface is a representation of the C record GTlsBackendInterface.
//
// since 2.28
type TlsBackendInterface struct {
	native unsafe.Pointer
}

// TlsCertificateClass is a representation of the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native unsafe.Pointer
}

// TlsCertificatePrivate is a representation of the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

// TlsClientConnectionInterface is a representation of the C record GTlsClientConnectionInterface.
//
// since 2.26
type TlsClientConnectionInterface struct {
	native unsafe.Pointer
}

// TlsConnectionClass is a representation of the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native unsafe.Pointer
}

// TlsConnectionPrivate is a representation of the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

// TlsDatabaseClass is a representation of the C record GTlsDatabaseClass.
//
// since 2.30
type TlsDatabaseClass struct {
	native unsafe.Pointer
}

// TlsDatabasePrivate is a representation of the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

// TlsFileDatabaseInterface is a representation of the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
}

// TlsInteractionClass is a representation of the C record GTlsInteractionClass.
//
// since 2.30
type TlsInteractionClass struct {
	native unsafe.Pointer
}

// TlsInteractionPrivate is a representation of the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

// TlsPasswordClass is a representation of the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native unsafe.Pointer
}

// TlsPasswordPrivate is a representation of the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

// TlsServerConnectionInterface is a representation of the C record GTlsServerConnectionInterface.
//
// since 2.26
type TlsServerConnectionInterface struct {
	native unsafe.Pointer
}

// UnixConnectionClass is a representation of the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native unsafe.Pointer
}

// UnixConnectionPrivate is a representation of the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

// UnixCredentialsMessageClass is a representation of the C record GUnixCredentialsMessageClass.
//
// since 2.26
type UnixCredentialsMessageClass struct {
	native unsafe.Pointer
}

// UnixCredentialsMessagePrivate is a representation of the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

// UnixFDListClass is a representation of the C record GUnixFDListClass.
type UnixFDListClass struct {
	native unsafe.Pointer
}

// UnixFDListPrivate is a representation of the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native unsafe.Pointer
}

// UnixFDMessageClass is a representation of the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native unsafe.Pointer
}

// UnixFDMessagePrivate is a representation of the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

// UnixInputStreamClass is a representation of the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native unsafe.Pointer
}

// UnixInputStreamPrivate is a representation of the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

// UnixMountEntry is a representation of the C record GUnixMountEntry.
type UnixMountEntry struct {
	native unsafe.Pointer
}

// UnixMountMonitorClass is a representation of the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

// UnixMountPoint is a representation of the C record GUnixMountPoint.
type UnixMountPoint struct {
	native unsafe.Pointer
}

// UnixOutputStreamClass is a representation of the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native unsafe.Pointer
}

// UnixOutputStreamPrivate is a representation of the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

// UnixSocketAddressClass is a representation of the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native unsafe.Pointer
}

// UnixSocketAddressPrivate is a representation of the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

// VfsClass is a representation of the C record GVfsClass.
type VfsClass struct {
	native unsafe.Pointer
}

// VolumeIface is a representation of the C record GVolumeIface.
type VolumeIface struct {
	native unsafe.Pointer
}

// VolumeMonitorClass is a representation of the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native unsafe.Pointer
}

// ZlibCompressorClass is a representation of the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native unsafe.Pointer
}

// ZlibDecompressorClass is a representation of the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native unsafe.Pointer
}
