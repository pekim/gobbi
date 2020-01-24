// Code generated - DO NOT EDIT.
// +build gio_2.50

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// FILE_ATTRIBUTE_FILESYSTEM_USED is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_USED.
//
// since 2.32
const FILE_ATTRIBUTE_FILESYSTEM_USED = "filesystem::used"

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

// FILE_ATTRIBUTE_STANDARD_IS_VOLATILE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_VOLATILE.
//
// since 2.46
const FILE_ATTRIBUTE_STANDARD_IS_VOLATILE = "standard::is-volatile"

// FILE_ATTRIBUTE_STANDARD_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_NAME.
const FILE_ATTRIBUTE_STANDARD_NAME = "standard::name"

// FILE_ATTRIBUTE_STANDARD_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SIZE.
const FILE_ATTRIBUTE_STANDARD_SIZE = "standard::size"

// FILE_ATTRIBUTE_STANDARD_SORT_ORDER is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER = "standard::sort-order"

// FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
//
// since 2.34
const FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON = "standard::symbolic-icon"

// FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET.
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET = "standard::symlink-target"

// FILE_ATTRIBUTE_STANDARD_TARGET_URI is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TARGET_URI.
const FILE_ATTRIBUTE_STANDARD_TARGET_URI = "standard::target-uri"

// FILE_ATTRIBUTE_STANDARD_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TYPE.
const FILE_ATTRIBUTE_STANDARD_TYPE = "standard::type"

// FILE_ATTRIBUTE_THUMBNAILING_FAILED is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAILING_FAILED.
const FILE_ATTRIBUTE_THUMBNAILING_FAILED = "thumbnail::failed"

// FILE_ATTRIBUTE_THUMBNAIL_IS_VALID is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID.
//
// since 2.40
const FILE_ATTRIBUTE_THUMBNAIL_IS_VALID = "thumbnail::is-valid"

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

// MENU_ATTRIBUTE_ACTION is a representation of the C constant G_MENU_ATTRIBUTE_ACTION.
//
// since 2.32
const MENU_ATTRIBUTE_ACTION = "action"

// MENU_ATTRIBUTE_ACTION_NAMESPACE is a representation of the C constant G_MENU_ATTRIBUTE_ACTION_NAMESPACE.
//
// since 2.36
const MENU_ATTRIBUTE_ACTION_NAMESPACE = "action-namespace"

// MENU_ATTRIBUTE_ICON is a representation of the C constant G_MENU_ATTRIBUTE_ICON.
//
// since 2.38
const MENU_ATTRIBUTE_ICON = "icon"

// MENU_ATTRIBUTE_LABEL is a representation of the C constant G_MENU_ATTRIBUTE_LABEL.
//
// since 2.32
const MENU_ATTRIBUTE_LABEL = "label"

// MENU_ATTRIBUTE_TARGET is a representation of the C constant G_MENU_ATTRIBUTE_TARGET.
//
// since 2.32
const MENU_ATTRIBUTE_TARGET = "target"

// MENU_LINK_SECTION is a representation of the C constant G_MENU_LINK_SECTION.
//
// since 2.32
const MENU_LINK_SECTION = "section"

// MENU_LINK_SUBMENU is a representation of the C constant G_MENU_LINK_SUBMENU.
//
// since 2.32
const MENU_LINK_SUBMENU = "submenu"

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

// FileMeasureFlags is a representation of the C bitfield GFileMeasureFlags.
type FileMeasureFlags int

// FileMeasureFlags_none is a representation of the C bitfield member G_FILE_MEASURE_NONE.
const FileMeasureFlags_none = FileMeasureFlags(0)

// FileMeasureFlags_report_any_error is a representation of the C bitfield member G_FILE_MEASURE_REPORT_ANY_ERROR.
const FileMeasureFlags_report_any_error = FileMeasureFlags(2)

// FileMeasureFlags_apparent_size is a representation of the C bitfield member G_FILE_MEASURE_APPARENT_SIZE.
const FileMeasureFlags_apparent_size = FileMeasureFlags(4)

// FileMeasureFlags_no_xdev is a representation of the C bitfield member G_FILE_MEASURE_NO_XDEV.
const FileMeasureFlags_no_xdev = FileMeasureFlags(8)

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

// ResourceFlags is a representation of the C bitfield GResourceFlags.
type ResourceFlags int

// ResourceFlags_none is a representation of the C bitfield member G_RESOURCE_FLAGS_NONE.
const ResourceFlags_none = ResourceFlags(0)

// ResourceFlags_compressed is a representation of the C bitfield member G_RESOURCE_FLAGS_COMPRESSED.
const ResourceFlags_compressed = ResourceFlags(1)

// ResourceLookupFlags is a representation of the C bitfield GResourceLookupFlags.
type ResourceLookupFlags int

// ResourceLookupFlags_none is a representation of the C bitfield member G_RESOURCE_LOOKUP_FLAGS_NONE.
const ResourceLookupFlags_none = ResourceLookupFlags(0)

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

// SubprocessFlags is a representation of the C bitfield GSubprocessFlags.
type SubprocessFlags int

// SubprocessFlags_none is a representation of the C bitfield member G_SUBPROCESS_FLAGS_NONE.
const SubprocessFlags_none = SubprocessFlags(0)

// SubprocessFlags_stdin_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDIN_PIPE.
const SubprocessFlags_stdin_pipe = SubprocessFlags(1)

// SubprocessFlags_stdin_inherit is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDIN_INHERIT.
const SubprocessFlags_stdin_inherit = SubprocessFlags(2)

// SubprocessFlags_stdout_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDOUT_PIPE.
const SubprocessFlags_stdout_pipe = SubprocessFlags(4)

// SubprocessFlags_stdout_silence is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDOUT_SILENCE.
const SubprocessFlags_stdout_silence = SubprocessFlags(8)

// SubprocessFlags_stderr_pipe is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_PIPE.
const SubprocessFlags_stderr_pipe = SubprocessFlags(16)

// SubprocessFlags_stderr_silence is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_SILENCE.
const SubprocessFlags_stderr_silence = SubprocessFlags(32)

// SubprocessFlags_stderr_merge is a representation of the C bitfield member G_SUBPROCESS_FLAGS_STDERR_MERGE.
const SubprocessFlags_stderr_merge = SubprocessFlags(64)

// SubprocessFlags_inherit_fds is a representation of the C bitfield member G_SUBPROCESS_FLAGS_INHERIT_FDS.
const SubprocessFlags_inherit_fds = SubprocessFlags(128)

// TestDBusFlags is a representation of the C bitfield GTestDBusFlags.
type TestDBusFlags int

// TestDBusFlags_none is a representation of the C bitfield member G_TEST_DBUS_NONE.
const TestDBusFlags_none = TestDBusFlags(0)

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

// NetworkConnectivity is a representation of the C enumeration GNetworkConnectivity.
type NetworkConnectivity int

// NetworkConnectivity_local is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_LOCAL.
const NetworkConnectivity_local = NetworkConnectivity(1)

// NetworkConnectivity_limited is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_LIMITED.
const NetworkConnectivity_limited = NetworkConnectivity(2)

// NetworkConnectivity_portal is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_PORTAL.
const NetworkConnectivity_portal = NetworkConnectivity(3)

// NetworkConnectivity_full is a representation of the C enumeration member G_NETWORK_CONNECTIVITY_FULL.
const NetworkConnectivity_full = NetworkConnectivity(4)

// NotificationPriority is a representation of the C enumeration GNotificationPriority.
type NotificationPriority int

// NotificationPriority_normal is a representation of the C enumeration member G_NOTIFICATION_PRIORITY_NORMAL.
const NotificationPriority_normal = NotificationPriority(0)

// NotificationPriority_low is a representation of the C enumeration member G_NOTIFICATION_PRIORITY_LOW.
const NotificationPriority_low = NotificationPriority(1)

// NotificationPriority_high is a representation of the C enumeration member G_NOTIFICATION_PRIORITY_HIGH.
const NotificationPriority_high = NotificationPriority(2)

// NotificationPriority_urgent is a representation of the C enumeration member G_NOTIFICATION_PRIORITY_URGENT.
const NotificationPriority_urgent = NotificationPriority(3)

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

// ResolverRecordType is a representation of the C enumeration GResolverRecordType.
type ResolverRecordType int

// ResolverRecordType_srv is a representation of the C enumeration member G_RESOLVER_RECORD_SRV.
const ResolverRecordType_srv = ResolverRecordType(1)

// ResolverRecordType_mx is a representation of the C enumeration member G_RESOLVER_RECORD_MX.
const ResolverRecordType_mx = ResolverRecordType(2)

// ResolverRecordType_txt is a representation of the C enumeration member G_RESOLVER_RECORD_TXT.
const ResolverRecordType_txt = ResolverRecordType(3)

// ResolverRecordType_soa is a representation of the C enumeration member G_RESOLVER_RECORD_SOA.
const ResolverRecordType_soa = ResolverRecordType(4)

// ResolverRecordType_ns is a representation of the C enumeration member G_RESOLVER_RECORD_NS.
const ResolverRecordType_ns = ResolverRecordType(5)

// ResourceError is a representation of the C enumeration GResourceError.
type ResourceError int

// ResourceError_not_found is a representation of the C enumeration member G_RESOURCE_ERROR_NOT_FOUND.
const ResourceError_not_found = ResourceError(0)

// ResourceError_internal is a representation of the C enumeration member G_RESOURCE_ERROR_INTERNAL.
const ResourceError_internal = ResourceError(1)

// SocketClientEvent is a representation of the C enumeration GSocketClientEvent.
type SocketClientEvent int

// SocketClientEvent_resolving is a representation of the C enumeration member G_SOCKET_CLIENT_RESOLVING.
const SocketClientEvent_resolving = SocketClientEvent(0)

// SocketClientEvent_resolved is a representation of the C enumeration member G_SOCKET_CLIENT_RESOLVED.
const SocketClientEvent_resolved = SocketClientEvent(1)

// SocketClientEvent_connecting is a representation of the C enumeration member G_SOCKET_CLIENT_CONNECTING.
const SocketClientEvent_connecting = SocketClientEvent(2)

// SocketClientEvent_connected is a representation of the C enumeration member G_SOCKET_CLIENT_CONNECTED.
const SocketClientEvent_connected = SocketClientEvent(3)

// SocketClientEvent_proxy_negotiating is a representation of the C enumeration member G_SOCKET_CLIENT_PROXY_NEGOTIATING.
const SocketClientEvent_proxy_negotiating = SocketClientEvent(4)

// SocketClientEvent_proxy_negotiated is a representation of the C enumeration member G_SOCKET_CLIENT_PROXY_NEGOTIATED.
const SocketClientEvent_proxy_negotiated = SocketClientEvent(5)

// SocketClientEvent_tls_handshaking is a representation of the C enumeration member G_SOCKET_CLIENT_TLS_HANDSHAKING.
const SocketClientEvent_tls_handshaking = SocketClientEvent(6)

// SocketClientEvent_tls_handshaked is a representation of the C enumeration member G_SOCKET_CLIENT_TLS_HANDSHAKED.
const SocketClientEvent_tls_handshaked = SocketClientEvent(7)

// SocketClientEvent_complete is a representation of the C enumeration member G_SOCKET_CLIENT_COMPLETE.
const SocketClientEvent_complete = SocketClientEvent(8)

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

// SocketListenerEvent is a representation of the C enumeration GSocketListenerEvent.
type SocketListenerEvent int

// SocketListenerEvent_binding is a representation of the C enumeration member G_SOCKET_LISTENER_BINDING.
const SocketListenerEvent_binding = SocketListenerEvent(0)

// SocketListenerEvent_bound is a representation of the C enumeration member G_SOCKET_LISTENER_BOUND.
const SocketListenerEvent_bound = SocketListenerEvent(1)

// SocketListenerEvent_listening is a representation of the C enumeration member G_SOCKET_LISTENER_LISTENING.
const SocketListenerEvent_listening = SocketListenerEvent(2)

// SocketListenerEvent_listened is a representation of the C enumeration member G_SOCKET_LISTENER_LISTENED.
const SocketListenerEvent_listened = SocketListenerEvent(3)

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

// TlsCertificateRequestFlags is a representation of the C enumeration GTlsCertificateRequestFlags.
type TlsCertificateRequestFlags int

// TlsCertificateRequestFlags_none is a representation of the C enumeration member G_TLS_CERTIFICATE_REQUEST_NONE.
const TlsCertificateRequestFlags_none = TlsCertificateRequestFlags(0)

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

// BusOwnNameOnConnectionWithClosures wraps the C function g_bus_own_name_on_connection_with_closures.
//
// since 2.26
func BusOwnNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameOwnerFlags, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint {
	sys_connection := connection.ToC()
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAcquiredClosure := nameAcquiredClosure.ToC()
	sys_nameLostClosure := nameLostClosure.ToC()
	retSys := gio.Fn_g_bus_own_name_on_connection_with_closures(sys_connection, sys_name, sys_flags, sys_nameAcquiredClosure, sys_nameLostClosure)
	ret := retSys

	return ret
}

// BusOwnNameWithClosures wraps the C function g_bus_own_name_with_closures.
//
// since 2.26
func BusOwnNameWithClosures(busType BusType, name string, flags BusNameOwnerFlags, busAcquiredClosure *gobject.Closure, nameAcquiredClosure *gobject.Closure, nameLostClosure *gobject.Closure) uint {
	sys_busType := (int)(busType)
	sys_name := name
	sys_flags := (int)(flags)
	sys_busAcquiredClosure := busAcquiredClosure.ToC()
	sys_nameAcquiredClosure := nameAcquiredClosure.ToC()
	sys_nameLostClosure := nameLostClosure.ToC()
	retSys := gio.Fn_g_bus_own_name_with_closures(sys_busType, sys_name, sys_flags, sys_busAcquiredClosure, sys_nameAcquiredClosure, sys_nameLostClosure)
	ret := retSys

	return ret
}

// BusUnownName wraps the C function g_bus_unown_name.
//
// since 2.26
func BusUnownName(ownerId uint) {
	sys_ownerId := ownerId
	gio.Fn_g_bus_unown_name(sys_ownerId)
}

// BusUnwatchName wraps the C function g_bus_unwatch_name.
//
// since 2.26
func BusUnwatchName(watcherId uint) {
	sys_watcherId := watcherId
	gio.Fn_g_bus_unwatch_name(sys_watcherId)
}

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// BusWatchNameOnConnectionWithClosures wraps the C function g_bus_watch_name_on_connection_with_closures.
//
// since 2.26
func BusWatchNameOnConnectionWithClosures(connection *DBusConnection, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint {
	sys_connection := connection.ToC()
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAppearedClosure := nameAppearedClosure.ToC()
	sys_nameVanishedClosure := nameVanishedClosure.ToC()
	retSys := gio.Fn_g_bus_watch_name_on_connection_with_closures(sys_connection, sys_name, sys_flags, sys_nameAppearedClosure, sys_nameVanishedClosure)
	ret := retSys

	return ret
}

// BusWatchNameWithClosures wraps the C function g_bus_watch_name_with_closures.
//
// since 2.26
func BusWatchNameWithClosures(busType BusType, name string, flags BusNameWatcherFlags, nameAppearedClosure *gobject.Closure, nameVanishedClosure *gobject.Closure) uint {
	sys_busType := (int)(busType)
	sys_name := name
	sys_flags := (int)(flags)
	sys_nameAppearedClosure := nameAppearedClosure.ToC()
	sys_nameVanishedClosure := nameVanishedClosure.ToC()
	retSys := gio.Fn_g_bus_watch_name_with_closures(sys_busType, sys_name, sys_flags, sys_nameAppearedClosure, sys_nameVanishedClosure)
	ret := retSys

	return ret
}

// ContentTypeCanBeExecutable wraps the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_can_be_executable(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeEquals wraps the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	sys_type1 := type1
	sys_type2 := type2
	retSys := gio.Fn_g_content_type_equals(sys_type1, sys_type2)
	ret := retSys

	return ret
}

// ContentTypeFromMimeType wraps the C function g_content_type_from_mime_type.
//
// since 2.18
func ContentTypeFromMimeType(mimeType string) string {
	sys_mimeType := mimeType
	retSys := gio.Fn_g_content_type_from_mime_type(sys_mimeType)
	ret := retSys

	return ret
}

// ContentTypeGetDescription wraps the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_description(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeGetGenericIconName wraps the C function g_content_type_get_generic_icon_name.
//
// since 2.34
func ContentTypeGetGenericIconName(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_generic_icon_name(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeGetIcon wraps the C function g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) *Icon {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_icon(sys_type_)
	ret := IconNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// ContentTypeGetMimeType wraps the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_mime_type(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeGetSymbolicIcon wraps the C function g_content_type_get_symbolic_icon.
//
// since 2.34
func ContentTypeGetSymbolicIcon(type_ string) *Icon {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_symbolic_icon(sys_type_)
	ret := IconNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// ContentTypeIsA wraps the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	sys_type_ := type_
	sys_supertype := supertype
	retSys := gio.Fn_g_content_type_is_a(sys_type_, sys_supertype)
	ret := retSys

	return ret
}

// ContentTypeIsUnknown wraps the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_is_unknown(sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_content_type_set_mime_dirs : parameter 'dirs' is array parameter without length parameter

// ContentTypesGetRegistered wraps the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retSys := gio.Fn_g_content_types_get_registered()
	ret := glib.ListNewFromC(retSys)

	return ret
}

// DbusAddressEscapeValue wraps the C function g_dbus_address_escape_value.
//
// since 2.36
func DbusAddressEscapeValue(string_ string) string {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_address_escape_value(sys_string_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : throws

// UNSUPPORTED : g_dbus_address_get_stream_sync : throws

// UNSUPPORTED : g_dbus_annotation_info_lookup : parameter 'annotations' is array parameter without length parameter

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// DbusGenerateGuid wraps the C function g_dbus_generate_guid.
//
// since 2.26
func DbusGenerateGuid() string {
	retSys := gio.Fn_g_dbus_generate_guid()
	ret := retSys

	return ret
}

// DbusGvalueToGvariant wraps the C function g_dbus_gvalue_to_gvariant.
//
// since 2.30
func DbusGvalueToGvariant(gvalue *gobject.Value, type_ *glib.VariantType) *glib.Variant {
	sys_gvalue := gvalue.ToC()
	sys_type_ := type_.ToC()
	retSys := gio.Fn_g_dbus_gvalue_to_gvariant(sys_gvalue, sys_type_)
	ret := glib.VariantNewFromC(retSys)

	return ret
}

// DbusGvariantToGvalue wraps the C function g_dbus_gvariant_to_gvalue.
//
// since 2.30
func DbusGvariantToGvalue(value *glib.Variant) *gobject.Value {
	sys_value := value.ToC()
	gio.Fn_g_dbus_gvariant_to_gvalue(sys_value)
}

// DbusIsAddress wraps the C function g_dbus_is_address.
//
// since 2.26
func DbusIsAddress(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_address(sys_string_)
	ret := retSys

	return ret
}

// DbusIsGuid wraps the C function g_dbus_is_guid.
//
// since 2.26
func DbusIsGuid(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_guid(sys_string_)
	ret := retSys

	return ret
}

// DbusIsInterfaceName wraps the C function g_dbus_is_interface_name.
//
// since 2.26
func DbusIsInterfaceName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_interface_name(sys_string_)
	ret := retSys

	return ret
}

// DbusIsMemberName wraps the C function g_dbus_is_member_name.
//
// since 2.26
func DbusIsMemberName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_member_name(sys_string_)
	ret := retSys

	return ret
}

// DbusIsName wraps the C function g_dbus_is_name.
//
// since 2.26
func DbusIsName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_name(sys_string_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dbus_is_supported_address : throws

// DbusIsUniqueName wraps the C function g_dbus_is_unique_name.
//
// since 2.26
func DbusIsUniqueName(string_ string) bool {
	sys_string_ := string_
	retSys := gio.Fn_g_dbus_is_unique_name(sys_string_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : throws

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// IoErrorFromErrno wraps the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int) int {
	sys_errNo := errNo
	retSys := gio.Fn_g_io_error_from_errno(sys_errNo)
	ret := retSys

	return ret
}

// IoErrorQuark wraps the C function g_io_error_quark.
func IoErrorQuark() uint32 {
	retSys := gio.Fn_g_io_error_quark()
	ret := retSys

	return ret
}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// IoSchedulerCancelAllJobs wraps the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	gio.Fn_g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// NetworkingInit wraps the C function g_networking_init.
//
// since 2.36
func NetworkingInit() {
	gio.Fn_g_networking_init()
}

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// PollableSourceNew wraps the C function g_pollable_source_new.
//
// since 2.28
func PollableSourceNew(pollableStream *gobject.Object) *glib.Source {
	sys_pollableStream := pollableStream.ToC()
	retSys := gio.Fn_g_pollable_source_new(sys_pollableStream)
	ret := glib.SourceNewFromC(retSys)

	return ret
}

// PollableSourceNewFull wraps the C function g_pollable_source_new_full.
//
// since 2.34
func PollableSourceNewFull(pollableStream unsafe.Pointer, childSource *glib.Source, cancellable *Cancellable) *glib.Source {
	sys_pollableStream := pollableStream
	sys_childSource := childSource.ToC()
	sys_cancellable := cancellable.ToC()
	retSys := gio.Fn_g_pollable_source_new_full(sys_pollableStream, sys_childSource, sys_cancellable)
	ret := glib.SourceNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// ResourcesRegister wraps the C function g_resources_register.
//
// since 2.32
func ResourcesRegister(resource *Resource) {
	sys_resource := resource.ToC()
	gio.Fn_g_resources_register(sys_resource)
}

// ResourcesUnregister wraps the C function g_resources_unregister.
//
// since 2.32
func ResourcesUnregister(resource *Resource) {
	sys_resource := resource.ToC()
	gio.Fn_g_resources_unregister(sys_resource)
}

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UnixIsMountPathSystemInternal wraps the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	sys_mountPath := mountPath
	retSys := gio.Fn_g_unix_is_mount_path_system_internal(sys_mountPath)
	ret := retSys

	return ret
}

// UnixMountAt wraps the C function g_unix_mount_at.
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	sys_mountPath := mountPath
	retSys := gio.Fn_g_unix_mount_at(sys_mountPath)
	ret := UnixMountEntryNewFromC(retSys)

	return ret
}

// UnixMountCompare wraps the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int {
	sys_mount1 := mount1.ToC()
	sys_mount2 := mount2.ToC()
	retSys := gio.Fn_g_unix_mount_compare(sys_mount1, sys_mount2)
	ret := retSys

	return ret
}

// UnixMountFree wraps the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_free(sys_mountEntry)
}

// UnixMountGetDevicePath wraps the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_device_path(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGetFsType wraps the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_fs_type(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGetMountPath wraps the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_mount_path(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessCanEject wraps the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_can_eject(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessIcon wraps the C function g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_icon(sys_mountEntry)
	ret := IconNewFromC(retSys)

	return ret
}

// UnixMountGuessName wraps the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_name(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessShouldDisplay wraps the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_should_display(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessSymbolicIcon wraps the C function g_unix_mount_guess_symbolic_icon.
//
// since 2.34
func UnixMountGuessSymbolicIcon(mountEntry *UnixMountEntry) *Icon {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_symbolic_icon(sys_mountEntry)
	ret := IconNewFromC(retSys)

	return ret
}

// UnixMountIsReadonly wraps the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_is_readonly(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountIsSystemInternal wraps the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_is_system_internal(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountPointsChangedSince wraps the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	sys_time := time
	retSys := gio.Fn_g_unix_mount_points_changed_since(sys_time)
	ret := retSys

	return ret
}

// UnixMountPointsGet wraps the C function g_unix_mount_points_get.
func UnixMountPointsGet() (*glib.List, uint64) {
	retSys := gio.Fn_g_unix_mount_points_get()
	ret := glib.ListNewFromC(retSys)

	return ret
}

// UnixMountsChangedSince wraps the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	sys_time := time
	retSys := gio.Fn_g_unix_mounts_changed_since(sys_time)
	ret := retSys

	return ret
}

// UnixMountsGet wraps the C function g_unix_mounts_get.
func UnixMountsGet() (*glib.List, uint64) {
	retSys := gio.Fn_g_unix_mounts_get()
	ret := glib.ListNewFromC(retSys)

	return ret
}

// ActionEntry is a representation of the C record GActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionEntry that represents the ActionEntry.
func (recv *ActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ActionEntryNewFromC creates a new ActionEntry from a pointer to the C GActionEntry that represents the ActionEntry.
func ActionEntryNewFromC(native unsafe.Pointer) *ActionEntry {
	return &ActionEntry{native: native}
}

// ActionGroupInterface is a representation of the C record GActionGroupInterface.
//
// since 2.28
type ActionGroupInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionGroupInterface that represents the ActionGroupInterface.
func (recv *ActionGroupInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupInterfaceNewFromC creates a new ActionGroupInterface from a pointer to the C GActionGroupInterface that represents the ActionGroupInterface.
func ActionGroupInterfaceNewFromC(native unsafe.Pointer) *ActionGroupInterface {
	return &ActionGroupInterface{native: native}
}

// ActionInterface is a representation of the C record GActionInterface.
//
// since 2.28
type ActionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionInterface that represents the ActionInterface.
func (recv *ActionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionInterfaceNewFromC creates a new ActionInterface from a pointer to the C GActionInterface that represents the ActionInterface.
func ActionInterfaceNewFromC(native unsafe.Pointer) *ActionInterface {
	return &ActionInterface{native: native}
}

// ActionMapInterface is a representation of the C record GActionMapInterface.
//
// since 2.32
type ActionMapInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionMapInterface that represents the ActionMapInterface.
func (recv *ActionMapInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ActionMapInterfaceNewFromC creates a new ActionMapInterface from a pointer to the C GActionMapInterface that represents the ActionMapInterface.
func ActionMapInterfaceNewFromC(native unsafe.Pointer) *ActionMapInterface {
	return &ActionMapInterface{native: native}
}

// AppInfoIface is a representation of the C record GAppInfoIface.
type AppInfoIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfoIface that represents the AppInfoIface.
func (recv *AppInfoIface) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoIfaceNewFromC creates a new AppInfoIface from a pointer to the C GAppInfoIface that represents the AppInfoIface.
func AppInfoIfaceNewFromC(native unsafe.Pointer) *AppInfoIface {
	return &AppInfoIface{native: native}
}

// AppLaunchContextClass is a representation of the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextClass that represents the AppLaunchContextClass.
func (recv *AppLaunchContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextClassNewFromC creates a new AppLaunchContextClass from a pointer to the C GAppLaunchContextClass that represents the AppLaunchContextClass.
func AppLaunchContextClassNewFromC(native unsafe.Pointer) *AppLaunchContextClass {
	return &AppLaunchContextClass{native: native}
}

// AppLaunchContextPrivate is a representation of the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextPrivate that represents the AppLaunchContextPrivate.
func (recv *AppLaunchContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextPrivateNewFromC creates a new AppLaunchContextPrivate from a pointer to the C GAppLaunchContextPrivate that represents the AppLaunchContextPrivate.
func AppLaunchContextPrivateNewFromC(native unsafe.Pointer) *AppLaunchContextPrivate {
	return &AppLaunchContextPrivate{native: native}
}

// ApplicationClass is a representation of the C record GApplicationClass.
//
// since 2.28
type ApplicationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationClass that represents the ApplicationClass.
func (recv *ApplicationClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationClassNewFromC creates a new ApplicationClass from a pointer to the C GApplicationClass that represents the ApplicationClass.
func ApplicationClassNewFromC(native unsafe.Pointer) *ApplicationClass {
	return &ApplicationClass{native: native}
}

// ApplicationCommandLineClass is a representation of the C record GApplicationCommandLineClass.
//
// since 2.28
type ApplicationCommandLineClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLineClass that represents the ApplicationCommandLineClass.
func (recv *ApplicationCommandLineClass) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLineClassNewFromC creates a new ApplicationCommandLineClass from a pointer to the C GApplicationCommandLineClass that represents the ApplicationCommandLineClass.
func ApplicationCommandLineClassNewFromC(native unsafe.Pointer) *ApplicationCommandLineClass {
	return &ApplicationCommandLineClass{native: native}
}

// ApplicationCommandLinePrivate is a representation of the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLinePrivate that represents the ApplicationCommandLinePrivate.
func (recv *ApplicationCommandLinePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLinePrivateNewFromC creates a new ApplicationCommandLinePrivate from a pointer to the C GApplicationCommandLinePrivate that represents the ApplicationCommandLinePrivate.
func ApplicationCommandLinePrivateNewFromC(native unsafe.Pointer) *ApplicationCommandLinePrivate {
	return &ApplicationCommandLinePrivate{native: native}
}

// ApplicationPrivate is a representation of the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationPrivate that represents the ApplicationPrivate.
func (recv *ApplicationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivateNewFromC creates a new ApplicationPrivate from a pointer to the C GApplicationPrivate that represents the ApplicationPrivate.
func ApplicationPrivateNewFromC(native unsafe.Pointer) *ApplicationPrivate {
	return &ApplicationPrivate{native: native}
}

// AsyncInitableIface is a representation of the C record GAsyncInitableIface.
//
// since 2.22
type AsyncInitableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncInitableIface that represents the AsyncInitableIface.
func (recv *AsyncInitableIface) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncInitableIfaceNewFromC creates a new AsyncInitableIface from a pointer to the C GAsyncInitableIface that represents the AsyncInitableIface.
func AsyncInitableIfaceNewFromC(native unsafe.Pointer) *AsyncInitableIface {
	return &AsyncInitableIface{native: native}
}

// AsyncResultIface is a representation of the C record GAsyncResultIface.
type AsyncResultIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResultIface that represents the AsyncResultIface.
func (recv *AsyncResultIface) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResultIfaceNewFromC creates a new AsyncResultIface from a pointer to the C GAsyncResultIface that represents the AsyncResultIface.
func AsyncResultIfaceNewFromC(native unsafe.Pointer) *AsyncResultIface {
	return &AsyncResultIface{native: native}
}

// BufferedInputStreamClass is a representation of the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamClass that represents the BufferedInputStreamClass.
func (recv *BufferedInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamClassNewFromC creates a new BufferedInputStreamClass from a pointer to the C GBufferedInputStreamClass that represents the BufferedInputStreamClass.
func BufferedInputStreamClassNewFromC(native unsafe.Pointer) *BufferedInputStreamClass {
	return &BufferedInputStreamClass{native: native}
}

// BufferedInputStreamPrivate is a representation of the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamPrivate that represents the BufferedInputStreamPrivate.
func (recv *BufferedInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamPrivateNewFromC creates a new BufferedInputStreamPrivate from a pointer to the C GBufferedInputStreamPrivate that represents the BufferedInputStreamPrivate.
func BufferedInputStreamPrivateNewFromC(native unsafe.Pointer) *BufferedInputStreamPrivate {
	return &BufferedInputStreamPrivate{native: native}
}

// BufferedOutputStreamClass is a representation of the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamClass that represents the BufferedOutputStreamClass.
func (recv *BufferedOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamClassNewFromC creates a new BufferedOutputStreamClass from a pointer to the C GBufferedOutputStreamClass that represents the BufferedOutputStreamClass.
func BufferedOutputStreamClassNewFromC(native unsafe.Pointer) *BufferedOutputStreamClass {
	return &BufferedOutputStreamClass{native: native}
}

// BufferedOutputStreamPrivate is a representation of the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamPrivate that represents the BufferedOutputStreamPrivate.
func (recv *BufferedOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamPrivateNewFromC creates a new BufferedOutputStreamPrivate from a pointer to the C GBufferedOutputStreamPrivate that represents the BufferedOutputStreamPrivate.
func BufferedOutputStreamPrivateNewFromC(native unsafe.Pointer) *BufferedOutputStreamPrivate {
	return &BufferedOutputStreamPrivate{native: native}
}

// CancellableClass is a representation of the C record GCancellableClass.
type CancellableClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellableClass that represents the CancellableClass.
func (recv *CancellableClass) ToC() unsafe.Pointer {
	return recv.native
}

// CancellableClassNewFromC creates a new CancellableClass from a pointer to the C GCancellableClass that represents the CancellableClass.
func CancellableClassNewFromC(native unsafe.Pointer) *CancellableClass {
	return &CancellableClass{native: native}
}

// CancellablePrivate is a representation of the C record GCancellablePrivate.
type CancellablePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellablePrivate that represents the CancellablePrivate.
func (recv *CancellablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CancellablePrivateNewFromC creates a new CancellablePrivate from a pointer to the C GCancellablePrivate that represents the CancellablePrivate.
func CancellablePrivateNewFromC(native unsafe.Pointer) *CancellablePrivate {
	return &CancellablePrivate{native: native}
}

// CharsetConverterClass is a representation of the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverterClass that represents the CharsetConverterClass.
func (recv *CharsetConverterClass) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverterClassNewFromC creates a new CharsetConverterClass from a pointer to the C GCharsetConverterClass that represents the CharsetConverterClass.
func CharsetConverterClassNewFromC(native unsafe.Pointer) *CharsetConverterClass {
	return &CharsetConverterClass{native: native}
}

// ConverterIface is a representation of the C record GConverterIface.
//
// since 2.24
type ConverterIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterIface that represents the ConverterIface.
func (recv *ConverterIface) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterIfaceNewFromC creates a new ConverterIface from a pointer to the C GConverterIface that represents the ConverterIface.
func ConverterIfaceNewFromC(native unsafe.Pointer) *ConverterIface {
	return &ConverterIface{native: native}
}

// ConverterInputStreamClass is a representation of the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamClass that represents the ConverterInputStreamClass.
func (recv *ConverterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamClassNewFromC creates a new ConverterInputStreamClass from a pointer to the C GConverterInputStreamClass that represents the ConverterInputStreamClass.
func ConverterInputStreamClassNewFromC(native unsafe.Pointer) *ConverterInputStreamClass {
	return &ConverterInputStreamClass{native: native}
}

// ConverterInputStreamPrivate is a representation of the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamPrivate that represents the ConverterInputStreamPrivate.
func (recv *ConverterInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamPrivateNewFromC creates a new ConverterInputStreamPrivate from a pointer to the C GConverterInputStreamPrivate that represents the ConverterInputStreamPrivate.
func ConverterInputStreamPrivateNewFromC(native unsafe.Pointer) *ConverterInputStreamPrivate {
	return &ConverterInputStreamPrivate{native: native}
}

// ConverterOutputStreamClass is a representation of the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamClass that represents the ConverterOutputStreamClass.
func (recv *ConverterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamClassNewFromC creates a new ConverterOutputStreamClass from a pointer to the C GConverterOutputStreamClass that represents the ConverterOutputStreamClass.
func ConverterOutputStreamClassNewFromC(native unsafe.Pointer) *ConverterOutputStreamClass {
	return &ConverterOutputStreamClass{native: native}
}

// ConverterOutputStreamPrivate is a representation of the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamPrivate that represents the ConverterOutputStreamPrivate.
func (recv *ConverterOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamPrivateNewFromC creates a new ConverterOutputStreamPrivate from a pointer to the C GConverterOutputStreamPrivate that represents the ConverterOutputStreamPrivate.
func ConverterOutputStreamPrivateNewFromC(native unsafe.Pointer) *ConverterOutputStreamPrivate {
	return &ConverterOutputStreamPrivate{native: native}
}

// CredentialsClass is a representation of the C record GCredentialsClass.
//
// since 2.26
type CredentialsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCredentialsClass that represents the CredentialsClass.
func (recv *CredentialsClass) ToC() unsafe.Pointer {
	return recv.native
}

// CredentialsClassNewFromC creates a new CredentialsClass from a pointer to the C GCredentialsClass that represents the CredentialsClass.
func CredentialsClassNewFromC(native unsafe.Pointer) *CredentialsClass {
	return &CredentialsClass{native: native}
}

// DBusAnnotationInfo is a representation of the C record GDBusAnnotationInfo.
//
// since 2.26
type DBusAnnotationInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusAnnotationInfo that represents the DBusAnnotationInfo.
func (recv *DBusAnnotationInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusAnnotationInfoNewFromC creates a new DBusAnnotationInfo from a pointer to the C GDBusAnnotationInfo that represents the DBusAnnotationInfo.
func DBusAnnotationInfoNewFromC(native unsafe.Pointer) *DBusAnnotationInfo {
	return &DBusAnnotationInfo{native: native}
}

// DBusArgInfo is a representation of the C record GDBusArgInfo.
//
// since 2.26
type DBusArgInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusArgInfo that represents the DBusArgInfo.
func (recv *DBusArgInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusArgInfoNewFromC creates a new DBusArgInfo from a pointer to the C GDBusArgInfo that represents the DBusArgInfo.
func DBusArgInfoNewFromC(native unsafe.Pointer) *DBusArgInfo {
	return &DBusArgInfo{native: native}
}

// DBusErrorEntry is a representation of the C record GDBusErrorEntry.
//
// since 2.26
type DBusErrorEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusErrorEntry that represents the DBusErrorEntry.
func (recv *DBusErrorEntry) ToC() unsafe.Pointer {
	return recv.native
}

// DBusErrorEntryNewFromC creates a new DBusErrorEntry from a pointer to the C GDBusErrorEntry that represents the DBusErrorEntry.
func DBusErrorEntryNewFromC(native unsafe.Pointer) *DBusErrorEntry {
	return &DBusErrorEntry{native: native}
}

// DBusInterfaceIface is a representation of the C record GDBusInterfaceIface.
//
// since 2.30
type DBusInterfaceIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceIface that represents the DBusInterfaceIface.
func (recv *DBusInterfaceIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceIfaceNewFromC creates a new DBusInterfaceIface from a pointer to the C GDBusInterfaceIface that represents the DBusInterfaceIface.
func DBusInterfaceIfaceNewFromC(native unsafe.Pointer) *DBusInterfaceIface {
	return &DBusInterfaceIface{native: native}
}

// DBusInterfaceInfo is a representation of the C record GDBusInterfaceInfo.
//
// since 2.26
type DBusInterfaceInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceInfo that represents the DBusInterfaceInfo.
func (recv *DBusInterfaceInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceInfoNewFromC creates a new DBusInterfaceInfo from a pointer to the C GDBusInterfaceInfo that represents the DBusInterfaceInfo.
func DBusInterfaceInfoNewFromC(native unsafe.Pointer) *DBusInterfaceInfo {
	return &DBusInterfaceInfo{native: native}
}

// DBusInterfaceSkeletonClass is a representation of the C record GDBusInterfaceSkeletonClass.
//
// since 2.30
type DBusInterfaceSkeletonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeletonClass that represents the DBusInterfaceSkeletonClass.
func (recv *DBusInterfaceSkeletonClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonClassNewFromC creates a new DBusInterfaceSkeletonClass from a pointer to the C GDBusInterfaceSkeletonClass that represents the DBusInterfaceSkeletonClass.
func DBusInterfaceSkeletonClassNewFromC(native unsafe.Pointer) *DBusInterfaceSkeletonClass {
	return &DBusInterfaceSkeletonClass{native: native}
}

// DBusInterfaceSkeletonPrivate is a representation of the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeletonPrivate that represents the DBusInterfaceSkeletonPrivate.
func (recv *DBusInterfaceSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonPrivateNewFromC creates a new DBusInterfaceSkeletonPrivate from a pointer to the C GDBusInterfaceSkeletonPrivate that represents the DBusInterfaceSkeletonPrivate.
func DBusInterfaceSkeletonPrivateNewFromC(native unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	return &DBusInterfaceSkeletonPrivate{native: native}
}

// DBusInterfaceVTable is a representation of the C record GDBusInterfaceVTable.
//
// since 2.26
type DBusInterfaceVTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceVTable that represents the DBusInterfaceVTable.
func (recv *DBusInterfaceVTable) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceVTableNewFromC creates a new DBusInterfaceVTable from a pointer to the C GDBusInterfaceVTable that represents the DBusInterfaceVTable.
func DBusInterfaceVTableNewFromC(native unsafe.Pointer) *DBusInterfaceVTable {
	return &DBusInterfaceVTable{native: native}
}

// DBusMethodInfo is a representation of the C record GDBusMethodInfo.
//
// since 2.26
type DBusMethodInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMethodInfo that represents the DBusMethodInfo.
func (recv *DBusMethodInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMethodInfoNewFromC creates a new DBusMethodInfo from a pointer to the C GDBusMethodInfo that represents the DBusMethodInfo.
func DBusMethodInfoNewFromC(native unsafe.Pointer) *DBusMethodInfo {
	return &DBusMethodInfo{native: native}
}

// DBusNodeInfo is a representation of the C record GDBusNodeInfo.
//
// since 2.26
type DBusNodeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusNodeInfo that represents the DBusNodeInfo.
func (recv *DBusNodeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusNodeInfoNewFromC creates a new DBusNodeInfo from a pointer to the C GDBusNodeInfo that represents the DBusNodeInfo.
func DBusNodeInfoNewFromC(native unsafe.Pointer) *DBusNodeInfo {
	return &DBusNodeInfo{native: native}
}

// DBusObjectIface is a representation of the C record GDBusObjectIface.
//
// since 2.30
type DBusObjectIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectIface that represents the DBusObjectIface.
func (recv *DBusObjectIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectIfaceNewFromC creates a new DBusObjectIface from a pointer to the C GDBusObjectIface that represents the DBusObjectIface.
func DBusObjectIfaceNewFromC(native unsafe.Pointer) *DBusObjectIface {
	return &DBusObjectIface{native: native}
}

// DBusObjectManagerClientClass is a representation of the C record GDBusObjectManagerClientClass.
//
// since 2.30
type DBusObjectManagerClientClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClientClass that represents the DBusObjectManagerClientClass.
func (recv *DBusObjectManagerClientClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientClassNewFromC creates a new DBusObjectManagerClientClass from a pointer to the C GDBusObjectManagerClientClass that represents the DBusObjectManagerClientClass.
func DBusObjectManagerClientClassNewFromC(native unsafe.Pointer) *DBusObjectManagerClientClass {
	return &DBusObjectManagerClientClass{native: native}
}

// DBusObjectManagerClientPrivate is a representation of the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClientPrivate that represents the DBusObjectManagerClientPrivate.
func (recv *DBusObjectManagerClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientPrivateNewFromC creates a new DBusObjectManagerClientPrivate from a pointer to the C GDBusObjectManagerClientPrivate that represents the DBusObjectManagerClientPrivate.
func DBusObjectManagerClientPrivateNewFromC(native unsafe.Pointer) *DBusObjectManagerClientPrivate {
	return &DBusObjectManagerClientPrivate{native: native}
}

// DBusObjectManagerIface is a representation of the C record GDBusObjectManagerIface.
//
// since 2.30
type DBusObjectManagerIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerIface that represents the DBusObjectManagerIface.
func (recv *DBusObjectManagerIface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerIfaceNewFromC creates a new DBusObjectManagerIface from a pointer to the C GDBusObjectManagerIface that represents the DBusObjectManagerIface.
func DBusObjectManagerIfaceNewFromC(native unsafe.Pointer) *DBusObjectManagerIface {
	return &DBusObjectManagerIface{native: native}
}

// DBusObjectManagerServerClass is a representation of the C record GDBusObjectManagerServerClass.
//
// since 2.30
type DBusObjectManagerServerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServerClass that represents the DBusObjectManagerServerClass.
func (recv *DBusObjectManagerServerClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerClassNewFromC creates a new DBusObjectManagerServerClass from a pointer to the C GDBusObjectManagerServerClass that represents the DBusObjectManagerServerClass.
func DBusObjectManagerServerClassNewFromC(native unsafe.Pointer) *DBusObjectManagerServerClass {
	return &DBusObjectManagerServerClass{native: native}
}

// DBusObjectManagerServerPrivate is a representation of the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServerPrivate that represents the DBusObjectManagerServerPrivate.
func (recv *DBusObjectManagerServerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerPrivateNewFromC creates a new DBusObjectManagerServerPrivate from a pointer to the C GDBusObjectManagerServerPrivate that represents the DBusObjectManagerServerPrivate.
func DBusObjectManagerServerPrivateNewFromC(native unsafe.Pointer) *DBusObjectManagerServerPrivate {
	return &DBusObjectManagerServerPrivate{native: native}
}

// DBusObjectProxyClass is a representation of the C record GDBusObjectProxyClass.
//
// since 2.30
type DBusObjectProxyClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxyClass that represents the DBusObjectProxyClass.
func (recv *DBusObjectProxyClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyClassNewFromC creates a new DBusObjectProxyClass from a pointer to the C GDBusObjectProxyClass that represents the DBusObjectProxyClass.
func DBusObjectProxyClassNewFromC(native unsafe.Pointer) *DBusObjectProxyClass {
	return &DBusObjectProxyClass{native: native}
}

// DBusObjectProxyPrivate is a representation of the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxyPrivate that represents the DBusObjectProxyPrivate.
func (recv *DBusObjectProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyPrivateNewFromC creates a new DBusObjectProxyPrivate from a pointer to the C GDBusObjectProxyPrivate that represents the DBusObjectProxyPrivate.
func DBusObjectProxyPrivateNewFromC(native unsafe.Pointer) *DBusObjectProxyPrivate {
	return &DBusObjectProxyPrivate{native: native}
}

// DBusObjectSkeletonClass is a representation of the C record GDBusObjectSkeletonClass.
//
// since 2.30
type DBusObjectSkeletonClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeletonClass that represents the DBusObjectSkeletonClass.
func (recv *DBusObjectSkeletonClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonClassNewFromC creates a new DBusObjectSkeletonClass from a pointer to the C GDBusObjectSkeletonClass that represents the DBusObjectSkeletonClass.
func DBusObjectSkeletonClassNewFromC(native unsafe.Pointer) *DBusObjectSkeletonClass {
	return &DBusObjectSkeletonClass{native: native}
}

// DBusObjectSkeletonPrivate is a representation of the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeletonPrivate that represents the DBusObjectSkeletonPrivate.
func (recv *DBusObjectSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonPrivateNewFromC creates a new DBusObjectSkeletonPrivate from a pointer to the C GDBusObjectSkeletonPrivate that represents the DBusObjectSkeletonPrivate.
func DBusObjectSkeletonPrivateNewFromC(native unsafe.Pointer) *DBusObjectSkeletonPrivate {
	return &DBusObjectSkeletonPrivate{native: native}
}

// DBusPropertyInfo is a representation of the C record GDBusPropertyInfo.
//
// since 2.26
type DBusPropertyInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusPropertyInfo that represents the DBusPropertyInfo.
func (recv *DBusPropertyInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusPropertyInfoNewFromC creates a new DBusPropertyInfo from a pointer to the C GDBusPropertyInfo that represents the DBusPropertyInfo.
func DBusPropertyInfoNewFromC(native unsafe.Pointer) *DBusPropertyInfo {
	return &DBusPropertyInfo{native: native}
}

// DBusProxyClass is a representation of the C record GDBusProxyClass.
//
// since 2.26
type DBusProxyClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxyClass that represents the DBusProxyClass.
func (recv *DBusProxyClass) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyClassNewFromC creates a new DBusProxyClass from a pointer to the C GDBusProxyClass that represents the DBusProxyClass.
func DBusProxyClassNewFromC(native unsafe.Pointer) *DBusProxyClass {
	return &DBusProxyClass{native: native}
}

// DBusProxyPrivate is a representation of the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxyPrivate that represents the DBusProxyPrivate.
func (recv *DBusProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyPrivateNewFromC creates a new DBusProxyPrivate from a pointer to the C GDBusProxyPrivate that represents the DBusProxyPrivate.
func DBusProxyPrivateNewFromC(native unsafe.Pointer) *DBusProxyPrivate {
	return &DBusProxyPrivate{native: native}
}

// DBusSignalInfo is a representation of the C record GDBusSignalInfo.
//
// since 2.26
type DBusSignalInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusSignalInfo that represents the DBusSignalInfo.
func (recv *DBusSignalInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DBusSignalInfoNewFromC creates a new DBusSignalInfo from a pointer to the C GDBusSignalInfo that represents the DBusSignalInfo.
func DBusSignalInfoNewFromC(native unsafe.Pointer) *DBusSignalInfo {
	return &DBusSignalInfo{native: native}
}

// DBusSubtreeVTable is a representation of the C record GDBusSubtreeVTable.
//
// since 2.26
type DBusSubtreeVTable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusSubtreeVTable that represents the DBusSubtreeVTable.
func (recv *DBusSubtreeVTable) ToC() unsafe.Pointer {
	return recv.native
}

// DBusSubtreeVTableNewFromC creates a new DBusSubtreeVTable from a pointer to the C GDBusSubtreeVTable that represents the DBusSubtreeVTable.
func DBusSubtreeVTableNewFromC(native unsafe.Pointer) *DBusSubtreeVTable {
	return &DBusSubtreeVTable{native: native}
}

// DataInputStreamClass is a representation of the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamClass that represents the DataInputStreamClass.
func (recv *DataInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamClassNewFromC creates a new DataInputStreamClass from a pointer to the C GDataInputStreamClass that represents the DataInputStreamClass.
func DataInputStreamClassNewFromC(native unsafe.Pointer) *DataInputStreamClass {
	return &DataInputStreamClass{native: native}
}

// DataInputStreamPrivate is a representation of the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamPrivate that represents the DataInputStreamPrivate.
func (recv *DataInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamPrivateNewFromC creates a new DataInputStreamPrivate from a pointer to the C GDataInputStreamPrivate that represents the DataInputStreamPrivate.
func DataInputStreamPrivateNewFromC(native unsafe.Pointer) *DataInputStreamPrivate {
	return &DataInputStreamPrivate{native: native}
}

// DataOutputStreamClass is a representation of the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamClass that represents the DataOutputStreamClass.
func (recv *DataOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamClassNewFromC creates a new DataOutputStreamClass from a pointer to the C GDataOutputStreamClass that represents the DataOutputStreamClass.
func DataOutputStreamClassNewFromC(native unsafe.Pointer) *DataOutputStreamClass {
	return &DataOutputStreamClass{native: native}
}

// DataOutputStreamPrivate is a representation of the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamPrivate that represents the DataOutputStreamPrivate.
func (recv *DataOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamPrivateNewFromC creates a new DataOutputStreamPrivate from a pointer to the C GDataOutputStreamPrivate that represents the DataOutputStreamPrivate.
func DataOutputStreamPrivateNewFromC(native unsafe.Pointer) *DataOutputStreamPrivate {
	return &DataOutputStreamPrivate{native: native}
}

// DatagramBasedInterface is a representation of the C record GDatagramBasedInterface.
//
// since 2.48
type DatagramBasedInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDatagramBasedInterface that represents the DatagramBasedInterface.
func (recv *DatagramBasedInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DatagramBasedInterfaceNewFromC creates a new DatagramBasedInterface from a pointer to the C GDatagramBasedInterface that represents the DatagramBasedInterface.
func DatagramBasedInterfaceNewFromC(native unsafe.Pointer) *DatagramBasedInterface {
	return &DatagramBasedInterface{native: native}
}

// DesktopAppInfoClass is a representation of the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoClass that represents the DesktopAppInfoClass.
func (recv *DesktopAppInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoClassNewFromC creates a new DesktopAppInfoClass from a pointer to the C GDesktopAppInfoClass that represents the DesktopAppInfoClass.
func DesktopAppInfoClassNewFromC(native unsafe.Pointer) *DesktopAppInfoClass {
	return &DesktopAppInfoClass{native: native}
}

// DesktopAppInfoLookupIface is a representation of the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookupIface that represents the DesktopAppInfoLookupIface.
func (recv *DesktopAppInfoLookupIface) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookupIfaceNewFromC creates a new DesktopAppInfoLookupIface from a pointer to the C GDesktopAppInfoLookupIface that represents the DesktopAppInfoLookupIface.
func DesktopAppInfoLookupIfaceNewFromC(native unsafe.Pointer) *DesktopAppInfoLookupIface {
	return &DesktopAppInfoLookupIface{native: native}
}

// DriveIface is a representation of the C record GDriveIface.
type DriveIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDriveIface that represents the DriveIface.
func (recv *DriveIface) ToC() unsafe.Pointer {
	return recv.native
}

// DriveIfaceNewFromC creates a new DriveIface from a pointer to the C GDriveIface that represents the DriveIface.
func DriveIfaceNewFromC(native unsafe.Pointer) *DriveIface {
	return &DriveIface{native: native}
}

// DtlsClientConnectionInterface is a representation of the C record GDtlsClientConnectionInterface.
//
// since 2.48
type DtlsClientConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsClientConnectionInterface that represents the DtlsClientConnectionInterface.
func (recv *DtlsClientConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsClientConnectionInterfaceNewFromC creates a new DtlsClientConnectionInterface from a pointer to the C GDtlsClientConnectionInterface that represents the DtlsClientConnectionInterface.
func DtlsClientConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsClientConnectionInterface {
	return &DtlsClientConnectionInterface{native: native}
}

// DtlsConnectionInterface is a representation of the C record GDtlsConnectionInterface.
//
// since 2.48
type DtlsConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsConnectionInterface that represents the DtlsConnectionInterface.
func (recv *DtlsConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsConnectionInterfaceNewFromC creates a new DtlsConnectionInterface from a pointer to the C GDtlsConnectionInterface that represents the DtlsConnectionInterface.
func DtlsConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsConnectionInterface {
	return &DtlsConnectionInterface{native: native}
}

// DtlsServerConnectionInterface is a representation of the C record GDtlsServerConnectionInterface.
//
// since 2.48
type DtlsServerConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsServerConnectionInterface that represents the DtlsServerConnectionInterface.
func (recv *DtlsServerConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsServerConnectionInterfaceNewFromC creates a new DtlsServerConnectionInterface from a pointer to the C GDtlsServerConnectionInterface that represents the DtlsServerConnectionInterface.
func DtlsServerConnectionInterfaceNewFromC(native unsafe.Pointer) *DtlsServerConnectionInterface {
	return &DtlsServerConnectionInterface{native: native}
}

// EmblemClass is a representation of the C record GEmblemClass.
type EmblemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemClass that represents the EmblemClass.
func (recv *EmblemClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemClassNewFromC creates a new EmblemClass from a pointer to the C GEmblemClass that represents the EmblemClass.
func EmblemClassNewFromC(native unsafe.Pointer) *EmblemClass {
	return &EmblemClass{native: native}
}

// EmblemedIconClass is a representation of the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconClass that represents the EmblemedIconClass.
func (recv *EmblemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconClassNewFromC creates a new EmblemedIconClass from a pointer to the C GEmblemedIconClass that represents the EmblemedIconClass.
func EmblemedIconClassNewFromC(native unsafe.Pointer) *EmblemedIconClass {
	return &EmblemedIconClass{native: native}
}

// EmblemedIconPrivate is a representation of the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconPrivate that represents the EmblemedIconPrivate.
func (recv *EmblemedIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconPrivateNewFromC creates a new EmblemedIconPrivate from a pointer to the C GEmblemedIconPrivate that represents the EmblemedIconPrivate.
func EmblemedIconPrivateNewFromC(native unsafe.Pointer) *EmblemedIconPrivate {
	return &EmblemedIconPrivate{native: native}
}

// FileAttributeInfo is a representation of the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfo that represents the FileAttributeInfo.
func (recv *FileAttributeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfoNewFromC creates a new FileAttributeInfo from a pointer to the C GFileAttributeInfo that represents the FileAttributeInfo.
func FileAttributeInfoNewFromC(native unsafe.Pointer) *FileAttributeInfo {
	return &FileAttributeInfo{native: native}
}

// FileAttributeInfoList is a representation of the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfoList that represents the FileAttributeInfoList.
func (recv *FileAttributeInfoList) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfoListNewFromC creates a new FileAttributeInfoList from a pointer to the C GFileAttributeInfoList that represents the FileAttributeInfoList.
func FileAttributeInfoListNewFromC(native unsafe.Pointer) *FileAttributeInfoList {
	return &FileAttributeInfoList{native: native}
}

// FileAttributeMatcher is a representation of the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeMatcher that represents the FileAttributeMatcher.
func (recv *FileAttributeMatcher) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeMatcherNewFromC creates a new FileAttributeMatcher from a pointer to the C GFileAttributeMatcher that represents the FileAttributeMatcher.
func FileAttributeMatcherNewFromC(native unsafe.Pointer) *FileAttributeMatcher {
	return &FileAttributeMatcher{native: native}
}

// FileDescriptorBasedIface is a representation of the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileDescriptorBasedIface that represents the FileDescriptorBasedIface.
func (recv *FileDescriptorBasedIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileDescriptorBasedIfaceNewFromC creates a new FileDescriptorBasedIface from a pointer to the C GFileDescriptorBasedIface that represents the FileDescriptorBasedIface.
func FileDescriptorBasedIfaceNewFromC(native unsafe.Pointer) *FileDescriptorBasedIface {
	return &FileDescriptorBasedIface{native: native}
}

// FileEnumeratorClass is a representation of the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorClass that represents the FileEnumeratorClass.
func (recv *FileEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorClassNewFromC creates a new FileEnumeratorClass from a pointer to the C GFileEnumeratorClass that represents the FileEnumeratorClass.
func FileEnumeratorClassNewFromC(native unsafe.Pointer) *FileEnumeratorClass {
	return &FileEnumeratorClass{native: native}
}

// FileEnumeratorPrivate is a representation of the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorPrivate that represents the FileEnumeratorPrivate.
func (recv *FileEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorPrivateNewFromC creates a new FileEnumeratorPrivate from a pointer to the C GFileEnumeratorPrivate that represents the FileEnumeratorPrivate.
func FileEnumeratorPrivateNewFromC(native unsafe.Pointer) *FileEnumeratorPrivate {
	return &FileEnumeratorPrivate{native: native}
}

// FileIOStreamClass is a representation of the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamClass that represents the FileIOStreamClass.
func (recv *FileIOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamClassNewFromC creates a new FileIOStreamClass from a pointer to the C GFileIOStreamClass that represents the FileIOStreamClass.
func FileIOStreamClassNewFromC(native unsafe.Pointer) *FileIOStreamClass {
	return &FileIOStreamClass{native: native}
}

// FileIOStreamPrivate is a representation of the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamPrivate that represents the FileIOStreamPrivate.
func (recv *FileIOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamPrivateNewFromC creates a new FileIOStreamPrivate from a pointer to the C GFileIOStreamPrivate that represents the FileIOStreamPrivate.
func FileIOStreamPrivateNewFromC(native unsafe.Pointer) *FileIOStreamPrivate {
	return &FileIOStreamPrivate{native: native}
}

// FileIconClass is a representation of the C record GFileIconClass.
type FileIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIconClass that represents the FileIconClass.
func (recv *FileIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIconClassNewFromC creates a new FileIconClass from a pointer to the C GFileIconClass that represents the FileIconClass.
func FileIconClassNewFromC(native unsafe.Pointer) *FileIconClass {
	return &FileIconClass{native: native}
}

// FileIface is a representation of the C record GFileIface.
type FileIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIface that represents the FileIface.
func (recv *FileIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileIfaceNewFromC creates a new FileIface from a pointer to the C GFileIface that represents the FileIface.
func FileIfaceNewFromC(native unsafe.Pointer) *FileIface {
	return &FileIface{native: native}
}

// FileInfoClass is a representation of the C record GFileInfoClass.
type FileInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfoClass that represents the FileInfoClass.
func (recv *FileInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfoClassNewFromC creates a new FileInfoClass from a pointer to the C GFileInfoClass that represents the FileInfoClass.
func FileInfoClassNewFromC(native unsafe.Pointer) *FileInfoClass {
	return &FileInfoClass{native: native}
}

// FileInputStreamClass is a representation of the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamClass that represents the FileInputStreamClass.
func (recv *FileInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamClassNewFromC creates a new FileInputStreamClass from a pointer to the C GFileInputStreamClass that represents the FileInputStreamClass.
func FileInputStreamClassNewFromC(native unsafe.Pointer) *FileInputStreamClass {
	return &FileInputStreamClass{native: native}
}

// FileInputStreamPrivate is a representation of the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamPrivate that represents the FileInputStreamPrivate.
func (recv *FileInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamPrivateNewFromC creates a new FileInputStreamPrivate from a pointer to the C GFileInputStreamPrivate that represents the FileInputStreamPrivate.
func FileInputStreamPrivateNewFromC(native unsafe.Pointer) *FileInputStreamPrivate {
	return &FileInputStreamPrivate{native: native}
}

// FileMonitorClass is a representation of the C record GFileMonitorClass.
type FileMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorClass that represents the FileMonitorClass.
func (recv *FileMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorClassNewFromC creates a new FileMonitorClass from a pointer to the C GFileMonitorClass that represents the FileMonitorClass.
func FileMonitorClassNewFromC(native unsafe.Pointer) *FileMonitorClass {
	return &FileMonitorClass{native: native}
}

// FileMonitorPrivate is a representation of the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorPrivate that represents the FileMonitorPrivate.
func (recv *FileMonitorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorPrivateNewFromC creates a new FileMonitorPrivate from a pointer to the C GFileMonitorPrivate that represents the FileMonitorPrivate.
func FileMonitorPrivateNewFromC(native unsafe.Pointer) *FileMonitorPrivate {
	return &FileMonitorPrivate{native: native}
}

// FileOutputStreamClass is a representation of the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamClass that represents the FileOutputStreamClass.
func (recv *FileOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamClassNewFromC creates a new FileOutputStreamClass from a pointer to the C GFileOutputStreamClass that represents the FileOutputStreamClass.
func FileOutputStreamClassNewFromC(native unsafe.Pointer) *FileOutputStreamClass {
	return &FileOutputStreamClass{native: native}
}

// FileOutputStreamPrivate is a representation of the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamPrivate that represents the FileOutputStreamPrivate.
func (recv *FileOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamPrivateNewFromC creates a new FileOutputStreamPrivate from a pointer to the C GFileOutputStreamPrivate that represents the FileOutputStreamPrivate.
func FileOutputStreamPrivateNewFromC(native unsafe.Pointer) *FileOutputStreamPrivate {
	return &FileOutputStreamPrivate{native: native}
}

// FilenameCompleterClass is a representation of the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleterClass that represents the FilenameCompleterClass.
func (recv *FilenameCompleterClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleterClassNewFromC creates a new FilenameCompleterClass from a pointer to the C GFilenameCompleterClass that represents the FilenameCompleterClass.
func FilenameCompleterClassNewFromC(native unsafe.Pointer) *FilenameCompleterClass {
	return &FilenameCompleterClass{native: native}
}

// FilterInputStreamClass is a representation of the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStreamClass that represents the FilterInputStreamClass.
func (recv *FilterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStreamClassNewFromC creates a new FilterInputStreamClass from a pointer to the C GFilterInputStreamClass that represents the FilterInputStreamClass.
func FilterInputStreamClassNewFromC(native unsafe.Pointer) *FilterInputStreamClass {
	return &FilterInputStreamClass{native: native}
}

// FilterOutputStreamClass is a representation of the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStreamClass that represents the FilterOutputStreamClass.
func (recv *FilterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStreamClassNewFromC creates a new FilterOutputStreamClass from a pointer to the C GFilterOutputStreamClass that represents the FilterOutputStreamClass.
func FilterOutputStreamClassNewFromC(native unsafe.Pointer) *FilterOutputStreamClass {
	return &FilterOutputStreamClass{native: native}
}

// IOExtension is a representation of the C record GIOExtension.
type IOExtension struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtension that represents the IOExtension.
func (recv *IOExtension) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtensionNewFromC creates a new IOExtension from a pointer to the C GIOExtension that represents the IOExtension.
func IOExtensionNewFromC(native unsafe.Pointer) *IOExtension {
	return &IOExtension{native: native}
}

// IOExtensionPoint is a representation of the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtensionPoint that represents the IOExtensionPoint.
func (recv *IOExtensionPoint) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtensionPointNewFromC creates a new IOExtensionPoint from a pointer to the C GIOExtensionPoint that represents the IOExtensionPoint.
func IOExtensionPointNewFromC(native unsafe.Pointer) *IOExtensionPoint {
	return &IOExtensionPoint{native: native}
}

// IOModuleClass is a representation of the C record GIOModuleClass.
type IOModuleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModuleClass that represents the IOModuleClass.
func (recv *IOModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleClassNewFromC creates a new IOModuleClass from a pointer to the C GIOModuleClass that represents the IOModuleClass.
func IOModuleClassNewFromC(native unsafe.Pointer) *IOModuleClass {
	return &IOModuleClass{native: native}
}

// IOModuleScope is a representation of the C record GIOModuleScope.
//
// since 2.30
type IOModuleScope struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModuleScope that represents the IOModuleScope.
func (recv *IOModuleScope) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleScopeNewFromC creates a new IOModuleScope from a pointer to the C GIOModuleScope that represents the IOModuleScope.
func IOModuleScopeNewFromC(native unsafe.Pointer) *IOModuleScope {
	return &IOModuleScope{native: native}
}

// IOSchedulerJob is a representation of the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOSchedulerJob that represents the IOSchedulerJob.
func (recv *IOSchedulerJob) ToC() unsafe.Pointer {
	return recv.native
}

// IOSchedulerJobNewFromC creates a new IOSchedulerJob from a pointer to the C GIOSchedulerJob that represents the IOSchedulerJob.
func IOSchedulerJobNewFromC(native unsafe.Pointer) *IOSchedulerJob {
	return &IOSchedulerJob{native: native}
}

// IOStreamAdapter is a representation of the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamAdapter that represents the IOStreamAdapter.
func (recv *IOStreamAdapter) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamAdapterNewFromC creates a new IOStreamAdapter from a pointer to the C GIOStreamAdapter that represents the IOStreamAdapter.
func IOStreamAdapterNewFromC(native unsafe.Pointer) *IOStreamAdapter {
	return &IOStreamAdapter{native: native}
}

// IOStreamClass is a representation of the C record GIOStreamClass.
type IOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamClass that represents the IOStreamClass.
func (recv *IOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamClassNewFromC creates a new IOStreamClass from a pointer to the C GIOStreamClass that represents the IOStreamClass.
func IOStreamClassNewFromC(native unsafe.Pointer) *IOStreamClass {
	return &IOStreamClass{native: native}
}

// IOStreamPrivate is a representation of the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamPrivate that represents the IOStreamPrivate.
func (recv *IOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamPrivateNewFromC creates a new IOStreamPrivate from a pointer to the C GIOStreamPrivate that represents the IOStreamPrivate.
func IOStreamPrivateNewFromC(native unsafe.Pointer) *IOStreamPrivate {
	return &IOStreamPrivate{native: native}
}

// IconIface is a representation of the C record GIconIface.
type IconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIconIface that represents the IconIface.
func (recv *IconIface) ToC() unsafe.Pointer {
	return recv.native
}

// IconIfaceNewFromC creates a new IconIface from a pointer to the C GIconIface that represents the IconIface.
func IconIfaceNewFromC(native unsafe.Pointer) *IconIface {
	return &IconIface{native: native}
}

// InetAddressClass is a representation of the C record GInetAddressClass.
type InetAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressClass that represents the InetAddressClass.
func (recv *InetAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressClassNewFromC creates a new InetAddressClass from a pointer to the C GInetAddressClass that represents the InetAddressClass.
func InetAddressClassNewFromC(native unsafe.Pointer) *InetAddressClass {
	return &InetAddressClass{native: native}
}

// InetAddressMaskClass is a representation of the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskClass that represents the InetAddressMaskClass.
func (recv *InetAddressMaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskClassNewFromC creates a new InetAddressMaskClass from a pointer to the C GInetAddressMaskClass that represents the InetAddressMaskClass.
func InetAddressMaskClassNewFromC(native unsafe.Pointer) *InetAddressMaskClass {
	return &InetAddressMaskClass{native: native}
}

// InetAddressMaskPrivate is a representation of the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskPrivate that represents the InetAddressMaskPrivate.
func (recv *InetAddressMaskPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskPrivateNewFromC creates a new InetAddressMaskPrivate from a pointer to the C GInetAddressMaskPrivate that represents the InetAddressMaskPrivate.
func InetAddressMaskPrivateNewFromC(native unsafe.Pointer) *InetAddressMaskPrivate {
	return &InetAddressMaskPrivate{native: native}
}

// InetAddressPrivate is a representation of the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressPrivate that represents the InetAddressPrivate.
func (recv *InetAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressPrivateNewFromC creates a new InetAddressPrivate from a pointer to the C GInetAddressPrivate that represents the InetAddressPrivate.
func InetAddressPrivateNewFromC(native unsafe.Pointer) *InetAddressPrivate {
	return &InetAddressPrivate{native: native}
}

// InetSocketAddressClass is a representation of the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressClass that represents the InetSocketAddressClass.
func (recv *InetSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressClassNewFromC creates a new InetSocketAddressClass from a pointer to the C GInetSocketAddressClass that represents the InetSocketAddressClass.
func InetSocketAddressClassNewFromC(native unsafe.Pointer) *InetSocketAddressClass {
	return &InetSocketAddressClass{native: native}
}

// InetSocketAddressPrivate is a representation of the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressPrivate that represents the InetSocketAddressPrivate.
func (recv *InetSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressPrivateNewFromC creates a new InetSocketAddressPrivate from a pointer to the C GInetSocketAddressPrivate that represents the InetSocketAddressPrivate.
func InetSocketAddressPrivateNewFromC(native unsafe.Pointer) *InetSocketAddressPrivate {
	return &InetSocketAddressPrivate{native: native}
}

// InitableIface is a representation of the C record GInitableIface.
//
// since 2.22
type InitableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitableIface that represents the InitableIface.
func (recv *InitableIface) ToC() unsafe.Pointer {
	return recv.native
}

// InitableIfaceNewFromC creates a new InitableIface from a pointer to the C GInitableIface that represents the InitableIface.
func InitableIfaceNewFromC(native unsafe.Pointer) *InitableIface {
	return &InitableIface{native: native}
}

// InputMessage is a representation of the C record GInputMessage.
//
// since 2.48
type InputMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputMessage that represents the InputMessage.
func (recv *InputMessage) ToC() unsafe.Pointer {
	return recv.native
}

// InputMessageNewFromC creates a new InputMessage from a pointer to the C GInputMessage that represents the InputMessage.
func InputMessageNewFromC(native unsafe.Pointer) *InputMessage {
	return &InputMessage{native: native}
}

// InputStreamClass is a representation of the C record GInputStreamClass.
type InputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamClass that represents the InputStreamClass.
func (recv *InputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamClassNewFromC creates a new InputStreamClass from a pointer to the C GInputStreamClass that represents the InputStreamClass.
func InputStreamClassNewFromC(native unsafe.Pointer) *InputStreamClass {
	return &InputStreamClass{native: native}
}

// InputStreamPrivate is a representation of the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamPrivate that represents the InputStreamPrivate.
func (recv *InputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamPrivateNewFromC creates a new InputStreamPrivate from a pointer to the C GInputStreamPrivate that represents the InputStreamPrivate.
func InputStreamPrivateNewFromC(native unsafe.Pointer) *InputStreamPrivate {
	return &InputStreamPrivate{native: native}
}

// InputVector is a representation of the C record GInputVector.
//
// since 2.22
type InputVector struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputVector that represents the InputVector.
func (recv *InputVector) ToC() unsafe.Pointer {
	return recv.native
}

// InputVectorNewFromC creates a new InputVector from a pointer to the C GInputVector that represents the InputVector.
func InputVectorNewFromC(native unsafe.Pointer) *InputVector {
	return &InputVector{native: native}
}

// ListModelInterface is a representation of the C record GListModelInterface.
//
// since 2.44
type ListModelInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListModelInterface that represents the ListModelInterface.
func (recv *ListModelInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ListModelInterfaceNewFromC creates a new ListModelInterface from a pointer to the C GListModelInterface that represents the ListModelInterface.
func ListModelInterfaceNewFromC(native unsafe.Pointer) *ListModelInterface {
	return &ListModelInterface{native: native}
}

// ListStoreClass is a representation of the C record GListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStoreClass that represents the ListStoreClass.
func (recv *ListStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreClassNewFromC creates a new ListStoreClass from a pointer to the C GListStoreClass that represents the ListStoreClass.
func ListStoreClassNewFromC(native unsafe.Pointer) *ListStoreClass {
	return &ListStoreClass{native: native}
}

// LoadableIconIface is a representation of the C record GLoadableIconIface.
type LoadableIconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIconIface that represents the LoadableIconIface.
func (recv *LoadableIconIface) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIconIfaceNewFromC creates a new LoadableIconIface from a pointer to the C GLoadableIconIface that represents the LoadableIconIface.
func LoadableIconIfaceNewFromC(native unsafe.Pointer) *LoadableIconIface {
	return &LoadableIconIface{native: native}
}

// MemoryInputStreamClass is a representation of the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamClass that represents the MemoryInputStreamClass.
func (recv *MemoryInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamClassNewFromC creates a new MemoryInputStreamClass from a pointer to the C GMemoryInputStreamClass that represents the MemoryInputStreamClass.
func MemoryInputStreamClassNewFromC(native unsafe.Pointer) *MemoryInputStreamClass {
	return &MemoryInputStreamClass{native: native}
}

// MemoryInputStreamPrivate is a representation of the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamPrivate that represents the MemoryInputStreamPrivate.
func (recv *MemoryInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamPrivateNewFromC creates a new MemoryInputStreamPrivate from a pointer to the C GMemoryInputStreamPrivate that represents the MemoryInputStreamPrivate.
func MemoryInputStreamPrivateNewFromC(native unsafe.Pointer) *MemoryInputStreamPrivate {
	return &MemoryInputStreamPrivate{native: native}
}

// MemoryOutputStreamClass is a representation of the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamClass that represents the MemoryOutputStreamClass.
func (recv *MemoryOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamClassNewFromC creates a new MemoryOutputStreamClass from a pointer to the C GMemoryOutputStreamClass that represents the MemoryOutputStreamClass.
func MemoryOutputStreamClassNewFromC(native unsafe.Pointer) *MemoryOutputStreamClass {
	return &MemoryOutputStreamClass{native: native}
}

// MemoryOutputStreamPrivate is a representation of the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamPrivate that represents the MemoryOutputStreamPrivate.
func (recv *MemoryOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamPrivateNewFromC creates a new MemoryOutputStreamPrivate from a pointer to the C GMemoryOutputStreamPrivate that represents the MemoryOutputStreamPrivate.
func MemoryOutputStreamPrivateNewFromC(native unsafe.Pointer) *MemoryOutputStreamPrivate {
	return &MemoryOutputStreamPrivate{native: native}
}

// MenuAttributeIterClass is a representation of the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterClass that represents the MenuAttributeIterClass.
func (recv *MenuAttributeIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterClassNewFromC creates a new MenuAttributeIterClass from a pointer to the C GMenuAttributeIterClass that represents the MenuAttributeIterClass.
func MenuAttributeIterClassNewFromC(native unsafe.Pointer) *MenuAttributeIterClass {
	return &MenuAttributeIterClass{native: native}
}

// MenuAttributeIterPrivate is a representation of the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterPrivate that represents the MenuAttributeIterPrivate.
func (recv *MenuAttributeIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterPrivateNewFromC creates a new MenuAttributeIterPrivate from a pointer to the C GMenuAttributeIterPrivate that represents the MenuAttributeIterPrivate.
func MenuAttributeIterPrivateNewFromC(native unsafe.Pointer) *MenuAttributeIterPrivate {
	return &MenuAttributeIterPrivate{native: native}
}

// MenuLinkIterClass is a representation of the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterClass that represents the MenuLinkIterClass.
func (recv *MenuLinkIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterClassNewFromC creates a new MenuLinkIterClass from a pointer to the C GMenuLinkIterClass that represents the MenuLinkIterClass.
func MenuLinkIterClassNewFromC(native unsafe.Pointer) *MenuLinkIterClass {
	return &MenuLinkIterClass{native: native}
}

// MenuLinkIterPrivate is a representation of the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterPrivate that represents the MenuLinkIterPrivate.
func (recv *MenuLinkIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterPrivateNewFromC creates a new MenuLinkIterPrivate from a pointer to the C GMenuLinkIterPrivate that represents the MenuLinkIterPrivate.
func MenuLinkIterPrivateNewFromC(native unsafe.Pointer) *MenuLinkIterPrivate {
	return &MenuLinkIterPrivate{native: native}
}

// MenuModelClass is a representation of the C record GMenuModelClass.
type MenuModelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelClass that represents the MenuModelClass.
func (recv *MenuModelClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelClassNewFromC creates a new MenuModelClass from a pointer to the C GMenuModelClass that represents the MenuModelClass.
func MenuModelClassNewFromC(native unsafe.Pointer) *MenuModelClass {
	return &MenuModelClass{native: native}
}

// MenuModelPrivate is a representation of the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelPrivate that represents the MenuModelPrivate.
func (recv *MenuModelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelPrivateNewFromC creates a new MenuModelPrivate from a pointer to the C GMenuModelPrivate that represents the MenuModelPrivate.
func MenuModelPrivateNewFromC(native unsafe.Pointer) *MenuModelPrivate {
	return &MenuModelPrivate{native: native}
}

// MountIface is a representation of the C record GMountIface.
type MountIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountIface that represents the MountIface.
func (recv *MountIface) ToC() unsafe.Pointer {
	return recv.native
}

// MountIfaceNewFromC creates a new MountIface from a pointer to the C GMountIface that represents the MountIface.
func MountIfaceNewFromC(native unsafe.Pointer) *MountIface {
	return &MountIface{native: native}
}

// MountOperationClass is a representation of the C record GMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationClass that represents the MountOperationClass.
func (recv *MountOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationClassNewFromC creates a new MountOperationClass from a pointer to the C GMountOperationClass that represents the MountOperationClass.
func MountOperationClassNewFromC(native unsafe.Pointer) *MountOperationClass {
	return &MountOperationClass{native: native}
}

// MountOperationPrivate is a representation of the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationPrivate that represents the MountOperationPrivate.
func (recv *MountOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivateNewFromC creates a new MountOperationPrivate from a pointer to the C GMountOperationPrivate that represents the MountOperationPrivate.
func MountOperationPrivateNewFromC(native unsafe.Pointer) *MountOperationPrivate {
	return &MountOperationPrivate{native: native}
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

// NativeVolumeMonitorClassNewFromC creates a new NativeVolumeMonitorClass from a pointer to the C GNativeVolumeMonitorClass that represents the NativeVolumeMonitorClass.
func NativeVolumeMonitorClassNewFromC(native unsafe.Pointer) *NativeVolumeMonitorClass {
	return &NativeVolumeMonitorClass{native: native}
}

// NetworkAddressClass is a representation of the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressClass that represents the NetworkAddressClass.
func (recv *NetworkAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressClassNewFromC creates a new NetworkAddressClass from a pointer to the C GNetworkAddressClass that represents the NetworkAddressClass.
func NetworkAddressClassNewFromC(native unsafe.Pointer) *NetworkAddressClass {
	return &NetworkAddressClass{native: native}
}

// NetworkAddressPrivate is a representation of the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressPrivate that represents the NetworkAddressPrivate.
func (recv *NetworkAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressPrivateNewFromC creates a new NetworkAddressPrivate from a pointer to the C GNetworkAddressPrivate that represents the NetworkAddressPrivate.
func NetworkAddressPrivateNewFromC(native unsafe.Pointer) *NetworkAddressPrivate {
	return &NetworkAddressPrivate{native: native}
}

// NetworkMonitorInterface is a representation of the C record GNetworkMonitorInterface.
//
// since 2.32
type NetworkMonitorInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkMonitorInterface that represents the NetworkMonitorInterface.
func (recv *NetworkMonitorInterface) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkMonitorInterfaceNewFromC creates a new NetworkMonitorInterface from a pointer to the C GNetworkMonitorInterface that represents the NetworkMonitorInterface.
func NetworkMonitorInterfaceNewFromC(native unsafe.Pointer) *NetworkMonitorInterface {
	return &NetworkMonitorInterface{native: native}
}

// NetworkServiceClass is a representation of the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServiceClass that represents the NetworkServiceClass.
func (recv *NetworkServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServiceClassNewFromC creates a new NetworkServiceClass from a pointer to the C GNetworkServiceClass that represents the NetworkServiceClass.
func NetworkServiceClassNewFromC(native unsafe.Pointer) *NetworkServiceClass {
	return &NetworkServiceClass{native: native}
}

// NetworkServicePrivate is a representation of the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServicePrivate that represents the NetworkServicePrivate.
func (recv *NetworkServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServicePrivateNewFromC creates a new NetworkServicePrivate from a pointer to the C GNetworkServicePrivate that represents the NetworkServicePrivate.
func NetworkServicePrivateNewFromC(native unsafe.Pointer) *NetworkServicePrivate {
	return &NetworkServicePrivate{native: native}
}

// OutputMessage is a representation of the C record GOutputMessage.
//
// since 2.44
type OutputMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputMessage that represents the OutputMessage.
func (recv *OutputMessage) ToC() unsafe.Pointer {
	return recv.native
}

// OutputMessageNewFromC creates a new OutputMessage from a pointer to the C GOutputMessage that represents the OutputMessage.
func OutputMessageNewFromC(native unsafe.Pointer) *OutputMessage {
	return &OutputMessage{native: native}
}

// OutputStreamClass is a representation of the C record GOutputStreamClass.
type OutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamClass that represents the OutputStreamClass.
func (recv *OutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamClassNewFromC creates a new OutputStreamClass from a pointer to the C GOutputStreamClass that represents the OutputStreamClass.
func OutputStreamClassNewFromC(native unsafe.Pointer) *OutputStreamClass {
	return &OutputStreamClass{native: native}
}

// OutputStreamPrivate is a representation of the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamPrivate that represents the OutputStreamPrivate.
func (recv *OutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamPrivateNewFromC creates a new OutputStreamPrivate from a pointer to the C GOutputStreamPrivate that represents the OutputStreamPrivate.
func OutputStreamPrivateNewFromC(native unsafe.Pointer) *OutputStreamPrivate {
	return &OutputStreamPrivate{native: native}
}

// OutputVector is a representation of the C record GOutputVector.
//
// since 2.22
type OutputVector struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputVector that represents the OutputVector.
func (recv *OutputVector) ToC() unsafe.Pointer {
	return recv.native
}

// OutputVectorNewFromC creates a new OutputVector from a pointer to the C GOutputVector that represents the OutputVector.
func OutputVectorNewFromC(native unsafe.Pointer) *OutputVector {
	return &OutputVector{native: native}
}

// PermissionClass is a representation of the C record GPermissionClass.
type PermissionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionClass that represents the PermissionClass.
func (recv *PermissionClass) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionClassNewFromC creates a new PermissionClass from a pointer to the C GPermissionClass that represents the PermissionClass.
func PermissionClassNewFromC(native unsafe.Pointer) *PermissionClass {
	return &PermissionClass{native: native}
}

// PermissionPrivate is a representation of the C record GPermissionPrivate.
type PermissionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionPrivate that represents the PermissionPrivate.
func (recv *PermissionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionPrivateNewFromC creates a new PermissionPrivate from a pointer to the C GPermissionPrivate that represents the PermissionPrivate.
func PermissionPrivateNewFromC(native unsafe.Pointer) *PermissionPrivate {
	return &PermissionPrivate{native: native}
}

// PollableInputStreamInterface is a representation of the C record GPollableInputStreamInterface.
//
// since 2.28
type PollableInputStreamInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableInputStreamInterface that represents the PollableInputStreamInterface.
func (recv *PollableInputStreamInterface) ToC() unsafe.Pointer {
	return recv.native
}

// PollableInputStreamInterfaceNewFromC creates a new PollableInputStreamInterface from a pointer to the C GPollableInputStreamInterface that represents the PollableInputStreamInterface.
func PollableInputStreamInterfaceNewFromC(native unsafe.Pointer) *PollableInputStreamInterface {
	return &PollableInputStreamInterface{native: native}
}

// PollableOutputStreamInterface is a representation of the C record GPollableOutputStreamInterface.
//
// since 2.28
type PollableOutputStreamInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableOutputStreamInterface that represents the PollableOutputStreamInterface.
func (recv *PollableOutputStreamInterface) ToC() unsafe.Pointer {
	return recv.native
}

// PollableOutputStreamInterfaceNewFromC creates a new PollableOutputStreamInterface from a pointer to the C GPollableOutputStreamInterface that represents the PollableOutputStreamInterface.
func PollableOutputStreamInterfaceNewFromC(native unsafe.Pointer) *PollableOutputStreamInterface {
	return &PollableOutputStreamInterface{native: native}
}

// ProxyAddressClass is a representation of the C record GProxyAddressClass.
//
// since 2.26
type ProxyAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressClass that represents the ProxyAddressClass.
func (recv *ProxyAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressClassNewFromC creates a new ProxyAddressClass from a pointer to the C GProxyAddressClass that represents the ProxyAddressClass.
func ProxyAddressClassNewFromC(native unsafe.Pointer) *ProxyAddressClass {
	return &ProxyAddressClass{native: native}
}

// ProxyAddressEnumeratorClass is a representation of the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorClass that represents the ProxyAddressEnumeratorClass.
func (recv *ProxyAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorClassNewFromC creates a new ProxyAddressEnumeratorClass from a pointer to the C GProxyAddressEnumeratorClass that represents the ProxyAddressEnumeratorClass.
func ProxyAddressEnumeratorClassNewFromC(native unsafe.Pointer) *ProxyAddressEnumeratorClass {
	return &ProxyAddressEnumeratorClass{native: native}
}

// ProxyAddressEnumeratorPrivate is a representation of the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorPrivate that represents the ProxyAddressEnumeratorPrivate.
func (recv *ProxyAddressEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorPrivateNewFromC creates a new ProxyAddressEnumeratorPrivate from a pointer to the C GProxyAddressEnumeratorPrivate that represents the ProxyAddressEnumeratorPrivate.
func ProxyAddressEnumeratorPrivateNewFromC(native unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	return &ProxyAddressEnumeratorPrivate{native: native}
}

// ProxyAddressPrivate is a representation of the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressPrivate that represents the ProxyAddressPrivate.
func (recv *ProxyAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressPrivateNewFromC creates a new ProxyAddressPrivate from a pointer to the C GProxyAddressPrivate that represents the ProxyAddressPrivate.
func ProxyAddressPrivateNewFromC(native unsafe.Pointer) *ProxyAddressPrivate {
	return &ProxyAddressPrivate{native: native}
}

// ProxyInterface is a representation of the C record GProxyInterface.
//
// since 2.26
type ProxyInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyInterface that represents the ProxyInterface.
func (recv *ProxyInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyInterfaceNewFromC creates a new ProxyInterface from a pointer to the C GProxyInterface that represents the ProxyInterface.
func ProxyInterfaceNewFromC(native unsafe.Pointer) *ProxyInterface {
	return &ProxyInterface{native: native}
}

// ProxyResolverInterface is a representation of the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyResolverInterface that represents the ProxyResolverInterface.
func (recv *ProxyResolverInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyResolverInterfaceNewFromC creates a new ProxyResolverInterface from a pointer to the C GProxyResolverInterface that represents the ProxyResolverInterface.
func ProxyResolverInterfaceNewFromC(native unsafe.Pointer) *ProxyResolverInterface {
	return &ProxyResolverInterface{native: native}
}

// RemoteActionGroupInterface is a representation of the C record GRemoteActionGroupInterface.
//
// since 2.32
type RemoteActionGroupInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GRemoteActionGroupInterface that represents the RemoteActionGroupInterface.
func (recv *RemoteActionGroupInterface) ToC() unsafe.Pointer {
	return recv.native
}

// RemoteActionGroupInterfaceNewFromC creates a new RemoteActionGroupInterface from a pointer to the C GRemoteActionGroupInterface that represents the RemoteActionGroupInterface.
func RemoteActionGroupInterfaceNewFromC(native unsafe.Pointer) *RemoteActionGroupInterface {
	return &RemoteActionGroupInterface{native: native}
}

// ResolverClass is a representation of the C record GResolverClass.
type ResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverClass that represents the ResolverClass.
func (recv *ResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverClassNewFromC creates a new ResolverClass from a pointer to the C GResolverClass that represents the ResolverClass.
func ResolverClassNewFromC(native unsafe.Pointer) *ResolverClass {
	return &ResolverClass{native: native}
}

// ResolverPrivate is a representation of the C record GResolverPrivate.
type ResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverPrivate that represents the ResolverPrivate.
func (recv *ResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverPrivateNewFromC creates a new ResolverPrivate from a pointer to the C GResolverPrivate that represents the ResolverPrivate.
func ResolverPrivateNewFromC(native unsafe.Pointer) *ResolverPrivate {
	return &ResolverPrivate{native: native}
}

// Resource is a representation of the C record GResource.
//
// since 2.32
type Resource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResource that represents the Resource.
func (recv *Resource) ToC() unsafe.Pointer {
	return recv.native
}

// ResourceNewFromC creates a new Resource from a pointer to the C GResource that represents the Resource.
func ResourceNewFromC(native unsafe.Pointer) *Resource {
	return &Resource{native: native}
}

// SeekableIface is a representation of the C record GSeekableIface.
type SeekableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekableIface that represents the SeekableIface.
func (recv *SeekableIface) ToC() unsafe.Pointer {
	return recv.native
}

// SeekableIfaceNewFromC creates a new SeekableIface from a pointer to the C GSeekableIface that represents the SeekableIface.
func SeekableIfaceNewFromC(native unsafe.Pointer) *SeekableIface {
	return &SeekableIface{native: native}
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

// SettingsClassNewFromC creates a new SettingsClass from a pointer to the C GSettingsClass that represents the SettingsClass.
func SettingsClassNewFromC(native unsafe.Pointer) *SettingsClass {
	return &SettingsClass{native: native}
}

// SettingsPrivate is a representation of the C record GSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsPrivate that represents the SettingsPrivate.
func (recv *SettingsPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsPrivateNewFromC creates a new SettingsPrivate from a pointer to the C GSettingsPrivate that represents the SettingsPrivate.
func SettingsPrivateNewFromC(native unsafe.Pointer) *SettingsPrivate {
	return &SettingsPrivate{native: native}
}

// SettingsSchema is a representation of the C record GSettingsSchema.
//
// since 2.32
type SettingsSchema struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchema that represents the SettingsSchema.
func (recv *SettingsSchema) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaNewFromC creates a new SettingsSchema from a pointer to the C GSettingsSchema that represents the SettingsSchema.
func SettingsSchemaNewFromC(native unsafe.Pointer) *SettingsSchema {
	return &SettingsSchema{native: native}
}

// SettingsSchemaKey is a representation of the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchemaKey that represents the SettingsSchemaKey.
func (recv *SettingsSchemaKey) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaKeyNewFromC creates a new SettingsSchemaKey from a pointer to the C GSettingsSchemaKey that represents the SettingsSchemaKey.
func SettingsSchemaKeyNewFromC(native unsafe.Pointer) *SettingsSchemaKey {
	return &SettingsSchemaKey{native: native}
}

// SettingsSchemaSource is a representation of the C record GSettingsSchemaSource.
//
// since 2.32
type SettingsSchemaSource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchemaSource that represents the SettingsSchemaSource.
func (recv *SettingsSchemaSource) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaSourceNewFromC creates a new SettingsSchemaSource from a pointer to the C GSettingsSchemaSource that represents the SettingsSchemaSource.
func SettingsSchemaSourceNewFromC(native unsafe.Pointer) *SettingsSchemaSource {
	return &SettingsSchemaSource{native: native}
}

// SimpleActionGroupClass is a representation of the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupClass that represents the SimpleActionGroupClass.
func (recv *SimpleActionGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupClassNewFromC creates a new SimpleActionGroupClass from a pointer to the C GSimpleActionGroupClass that represents the SimpleActionGroupClass.
func SimpleActionGroupClassNewFromC(native unsafe.Pointer) *SimpleActionGroupClass {
	return &SimpleActionGroupClass{native: native}
}

// SimpleActionGroupPrivate is a representation of the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupPrivate that represents the SimpleActionGroupPrivate.
func (recv *SimpleActionGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupPrivateNewFromC creates a new SimpleActionGroupPrivate from a pointer to the C GSimpleActionGroupPrivate that represents the SimpleActionGroupPrivate.
func SimpleActionGroupPrivateNewFromC(native unsafe.Pointer) *SimpleActionGroupPrivate {
	return &SimpleActionGroupPrivate{native: native}
}

// SimpleAsyncResultClass is a representation of the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResultClass that represents the SimpleAsyncResultClass.
func (recv *SimpleAsyncResultClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultClassNewFromC creates a new SimpleAsyncResultClass from a pointer to the C GSimpleAsyncResultClass that represents the SimpleAsyncResultClass.
func SimpleAsyncResultClassNewFromC(native unsafe.Pointer) *SimpleAsyncResultClass {
	return &SimpleAsyncResultClass{native: native}
}

// SimpleProxyResolverClass is a representation of the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverClass that represents the SimpleProxyResolverClass.
func (recv *SimpleProxyResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverClassNewFromC creates a new SimpleProxyResolverClass from a pointer to the C GSimpleProxyResolverClass that represents the SimpleProxyResolverClass.
func SimpleProxyResolverClassNewFromC(native unsafe.Pointer) *SimpleProxyResolverClass {
	return &SimpleProxyResolverClass{native: native}
}

// SimpleProxyResolverPrivate is a representation of the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverPrivate that represents the SimpleProxyResolverPrivate.
func (recv *SimpleProxyResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverPrivateNewFromC creates a new SimpleProxyResolverPrivate from a pointer to the C GSimpleProxyResolverPrivate that represents the SimpleProxyResolverPrivate.
func SimpleProxyResolverPrivateNewFromC(native unsafe.Pointer) *SimpleProxyResolverPrivate {
	return &SimpleProxyResolverPrivate{native: native}
}

// SocketAddressClass is a representation of the C record GSocketAddressClass.
type SocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressClass that represents the SocketAddressClass.
func (recv *SocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressClassNewFromC creates a new SocketAddressClass from a pointer to the C GSocketAddressClass that represents the SocketAddressClass.
func SocketAddressClassNewFromC(native unsafe.Pointer) *SocketAddressClass {
	return &SocketAddressClass{native: native}
}

// SocketAddressEnumeratorClass is a representation of the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumeratorClass that represents the SocketAddressEnumeratorClass.
func (recv *SocketAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumeratorClassNewFromC creates a new SocketAddressEnumeratorClass from a pointer to the C GSocketAddressEnumeratorClass that represents the SocketAddressEnumeratorClass.
func SocketAddressEnumeratorClassNewFromC(native unsafe.Pointer) *SocketAddressEnumeratorClass {
	return &SocketAddressEnumeratorClass{native: native}
}

// SocketClass is a representation of the C record GSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClass that represents the SocketClass.
func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClassNewFromC creates a new SocketClass from a pointer to the C GSocketClass that represents the SocketClass.
func SocketClassNewFromC(native unsafe.Pointer) *SocketClass {
	return &SocketClass{native: native}
}

// SocketClientClass is a representation of the C record GSocketClientClass.
type SocketClientClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientClass that represents the SocketClientClass.
func (recv *SocketClientClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientClassNewFromC creates a new SocketClientClass from a pointer to the C GSocketClientClass that represents the SocketClientClass.
func SocketClientClassNewFromC(native unsafe.Pointer) *SocketClientClass {
	return &SocketClientClass{native: native}
}

// SocketClientPrivate is a representation of the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientPrivate that represents the SocketClientPrivate.
func (recv *SocketClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientPrivateNewFromC creates a new SocketClientPrivate from a pointer to the C GSocketClientPrivate that represents the SocketClientPrivate.
func SocketClientPrivateNewFromC(native unsafe.Pointer) *SocketClientPrivate {
	return &SocketClientPrivate{native: native}
}

// SocketConnectableIface is a representation of the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectableIface that represents the SocketConnectableIface.
func (recv *SocketConnectableIface) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectableIfaceNewFromC creates a new SocketConnectableIface from a pointer to the C GSocketConnectableIface that represents the SocketConnectableIface.
func SocketConnectableIfaceNewFromC(native unsafe.Pointer) *SocketConnectableIface {
	return &SocketConnectableIface{native: native}
}

// SocketConnectionClass is a representation of the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionClass that represents the SocketConnectionClass.
func (recv *SocketConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionClassNewFromC creates a new SocketConnectionClass from a pointer to the C GSocketConnectionClass that represents the SocketConnectionClass.
func SocketConnectionClassNewFromC(native unsafe.Pointer) *SocketConnectionClass {
	return &SocketConnectionClass{native: native}
}

// SocketConnectionPrivate is a representation of the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionPrivate that represents the SocketConnectionPrivate.
func (recv *SocketConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionPrivateNewFromC creates a new SocketConnectionPrivate from a pointer to the C GSocketConnectionPrivate that represents the SocketConnectionPrivate.
func SocketConnectionPrivateNewFromC(native unsafe.Pointer) *SocketConnectionPrivate {
	return &SocketConnectionPrivate{native: native}
}

// SocketControlMessageClass is a representation of the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessageClass that represents the SocketControlMessageClass.
func (recv *SocketControlMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessageClassNewFromC creates a new SocketControlMessageClass from a pointer to the C GSocketControlMessageClass that represents the SocketControlMessageClass.
func SocketControlMessageClassNewFromC(native unsafe.Pointer) *SocketControlMessageClass {
	return &SocketControlMessageClass{native: native}
}

// SocketControlMessagePrivate is a representation of the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessagePrivate that represents the SocketControlMessagePrivate.
func (recv *SocketControlMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessagePrivateNewFromC creates a new SocketControlMessagePrivate from a pointer to the C GSocketControlMessagePrivate that represents the SocketControlMessagePrivate.
func SocketControlMessagePrivateNewFromC(native unsafe.Pointer) *SocketControlMessagePrivate {
	return &SocketControlMessagePrivate{native: native}
}

// SocketListenerClass is a representation of the C record GSocketListenerClass.
type SocketListenerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerClass that represents the SocketListenerClass.
func (recv *SocketListenerClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerClassNewFromC creates a new SocketListenerClass from a pointer to the C GSocketListenerClass that represents the SocketListenerClass.
func SocketListenerClassNewFromC(native unsafe.Pointer) *SocketListenerClass {
	return &SocketListenerClass{native: native}
}

// SocketListenerPrivate is a representation of the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerPrivate that represents the SocketListenerPrivate.
func (recv *SocketListenerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerPrivateNewFromC creates a new SocketListenerPrivate from a pointer to the C GSocketListenerPrivate that represents the SocketListenerPrivate.
func SocketListenerPrivateNewFromC(native unsafe.Pointer) *SocketListenerPrivate {
	return &SocketListenerPrivate{native: native}
}

// SocketPrivate is a representation of the C record GSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketPrivate that represents the SocketPrivate.
func (recv *SocketPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketPrivateNewFromC creates a new SocketPrivate from a pointer to the C GSocketPrivate that represents the SocketPrivate.
func SocketPrivateNewFromC(native unsafe.Pointer) *SocketPrivate {
	return &SocketPrivate{native: native}
}

// SocketServiceClass is a representation of the C record GSocketServiceClass.
type SocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServiceClass that represents the SocketServiceClass.
func (recv *SocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServiceClassNewFromC creates a new SocketServiceClass from a pointer to the C GSocketServiceClass that represents the SocketServiceClass.
func SocketServiceClassNewFromC(native unsafe.Pointer) *SocketServiceClass {
	return &SocketServiceClass{native: native}
}

// SocketServicePrivate is a representation of the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServicePrivate that represents the SocketServicePrivate.
func (recv *SocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServicePrivateNewFromC creates a new SocketServicePrivate from a pointer to the C GSocketServicePrivate that represents the SocketServicePrivate.
func SocketServicePrivateNewFromC(native unsafe.Pointer) *SocketServicePrivate {
	return &SocketServicePrivate{native: native}
}

// SrvTarget is a representation of the C record GSrvTarget.
type SrvTarget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSrvTarget that represents the SrvTarget.
func (recv *SrvTarget) ToC() unsafe.Pointer {
	return recv.native
}

// SrvTargetNewFromC creates a new SrvTarget from a pointer to the C GSrvTarget that represents the SrvTarget.
func SrvTargetNewFromC(native unsafe.Pointer) *SrvTarget {
	return &SrvTarget{native: native}
}

// StaticResource is a representation of the C record GStaticResource.
type StaticResource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GStaticResource that represents the StaticResource.
func (recv *StaticResource) ToC() unsafe.Pointer {
	return recv.native
}

// StaticResourceNewFromC creates a new StaticResource from a pointer to the C GStaticResource that represents the StaticResource.
func StaticResourceNewFromC(native unsafe.Pointer) *StaticResource {
	return &StaticResource{native: native}
}

// TaskClass is a representation of the C record GTaskClass.
type TaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTaskClass that represents the TaskClass.
func (recv *TaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// TaskClassNewFromC creates a new TaskClass from a pointer to the C GTaskClass that represents the TaskClass.
func TaskClassNewFromC(native unsafe.Pointer) *TaskClass {
	return &TaskClass{native: native}
}

// TcpConnectionClass is a representation of the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionClass that represents the TcpConnectionClass.
func (recv *TcpConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionClassNewFromC creates a new TcpConnectionClass from a pointer to the C GTcpConnectionClass that represents the TcpConnectionClass.
func TcpConnectionClassNewFromC(native unsafe.Pointer) *TcpConnectionClass {
	return &TcpConnectionClass{native: native}
}

// TcpConnectionPrivate is a representation of the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionPrivate that represents the TcpConnectionPrivate.
func (recv *TcpConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionPrivateNewFromC creates a new TcpConnectionPrivate from a pointer to the C GTcpConnectionPrivate that represents the TcpConnectionPrivate.
func TcpConnectionPrivateNewFromC(native unsafe.Pointer) *TcpConnectionPrivate {
	return &TcpConnectionPrivate{native: native}
}

// TcpWrapperConnectionClass is a representation of the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionClass that represents the TcpWrapperConnectionClass.
func (recv *TcpWrapperConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionClassNewFromC creates a new TcpWrapperConnectionClass from a pointer to the C GTcpWrapperConnectionClass that represents the TcpWrapperConnectionClass.
func TcpWrapperConnectionClassNewFromC(native unsafe.Pointer) *TcpWrapperConnectionClass {
	return &TcpWrapperConnectionClass{native: native}
}

// TcpWrapperConnectionPrivate is a representation of the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionPrivate that represents the TcpWrapperConnectionPrivate.
func (recv *TcpWrapperConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionPrivateNewFromC creates a new TcpWrapperConnectionPrivate from a pointer to the C GTcpWrapperConnectionPrivate that represents the TcpWrapperConnectionPrivate.
func TcpWrapperConnectionPrivateNewFromC(native unsafe.Pointer) *TcpWrapperConnectionPrivate {
	return &TcpWrapperConnectionPrivate{native: native}
}

// ThemedIconClass is a representation of the C record GThemedIconClass.
type ThemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIconClass that represents the ThemedIconClass.
func (recv *ThemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIconClassNewFromC creates a new ThemedIconClass from a pointer to the C GThemedIconClass that represents the ThemedIconClass.
func ThemedIconClassNewFromC(native unsafe.Pointer) *ThemedIconClass {
	return &ThemedIconClass{native: native}
}

// ThreadedSocketServiceClass is a representation of the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServiceClass that represents the ThreadedSocketServiceClass.
func (recv *ThreadedSocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServiceClassNewFromC creates a new ThreadedSocketServiceClass from a pointer to the C GThreadedSocketServiceClass that represents the ThreadedSocketServiceClass.
func ThreadedSocketServiceClassNewFromC(native unsafe.Pointer) *ThreadedSocketServiceClass {
	return &ThreadedSocketServiceClass{native: native}
}

// ThreadedSocketServicePrivate is a representation of the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServicePrivate that represents the ThreadedSocketServicePrivate.
func (recv *ThreadedSocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServicePrivateNewFromC creates a new ThreadedSocketServicePrivate from a pointer to the C GThreadedSocketServicePrivate that represents the ThreadedSocketServicePrivate.
func ThreadedSocketServicePrivateNewFromC(native unsafe.Pointer) *ThreadedSocketServicePrivate {
	return &ThreadedSocketServicePrivate{native: native}
}

// TlsBackendInterface is a representation of the C record GTlsBackendInterface.
//
// since 2.28
type TlsBackendInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsBackendInterface that represents the TlsBackendInterface.
func (recv *TlsBackendInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsBackendInterfaceNewFromC creates a new TlsBackendInterface from a pointer to the C GTlsBackendInterface that represents the TlsBackendInterface.
func TlsBackendInterfaceNewFromC(native unsafe.Pointer) *TlsBackendInterface {
	return &TlsBackendInterface{native: native}
}

// TlsCertificateClass is a representation of the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificateClass that represents the TlsCertificateClass.
func (recv *TlsCertificateClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificateClassNewFromC creates a new TlsCertificateClass from a pointer to the C GTlsCertificateClass that represents the TlsCertificateClass.
func TlsCertificateClassNewFromC(native unsafe.Pointer) *TlsCertificateClass {
	return &TlsCertificateClass{native: native}
}

// TlsCertificatePrivate is a representation of the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificatePrivate that represents the TlsCertificatePrivate.
func (recv *TlsCertificatePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificatePrivateNewFromC creates a new TlsCertificatePrivate from a pointer to the C GTlsCertificatePrivate that represents the TlsCertificatePrivate.
func TlsCertificatePrivateNewFromC(native unsafe.Pointer) *TlsCertificatePrivate {
	return &TlsCertificatePrivate{native: native}
}

// TlsClientConnectionInterface is a representation of the C record GTlsClientConnectionInterface.
//
// since 2.26
type TlsClientConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsClientConnectionInterface that represents the TlsClientConnectionInterface.
func (recv *TlsClientConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsClientConnectionInterfaceNewFromC creates a new TlsClientConnectionInterface from a pointer to the C GTlsClientConnectionInterface that represents the TlsClientConnectionInterface.
func TlsClientConnectionInterfaceNewFromC(native unsafe.Pointer) *TlsClientConnectionInterface {
	return &TlsClientConnectionInterface{native: native}
}

// TlsConnectionClass is a representation of the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionClass that represents the TlsConnectionClass.
func (recv *TlsConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionClassNewFromC creates a new TlsConnectionClass from a pointer to the C GTlsConnectionClass that represents the TlsConnectionClass.
func TlsConnectionClassNewFromC(native unsafe.Pointer) *TlsConnectionClass {
	return &TlsConnectionClass{native: native}
}

// TlsConnectionPrivate is a representation of the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionPrivate that represents the TlsConnectionPrivate.
func (recv *TlsConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionPrivateNewFromC creates a new TlsConnectionPrivate from a pointer to the C GTlsConnectionPrivate that represents the TlsConnectionPrivate.
func TlsConnectionPrivateNewFromC(native unsafe.Pointer) *TlsConnectionPrivate {
	return &TlsConnectionPrivate{native: native}
}

// TlsDatabaseClass is a representation of the C record GTlsDatabaseClass.
//
// since 2.30
type TlsDatabaseClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabaseClass that represents the TlsDatabaseClass.
func (recv *TlsDatabaseClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabaseClassNewFromC creates a new TlsDatabaseClass from a pointer to the C GTlsDatabaseClass that represents the TlsDatabaseClass.
func TlsDatabaseClassNewFromC(native unsafe.Pointer) *TlsDatabaseClass {
	return &TlsDatabaseClass{native: native}
}

// TlsDatabasePrivate is a representation of the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabasePrivate that represents the TlsDatabasePrivate.
func (recv *TlsDatabasePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabasePrivateNewFromC creates a new TlsDatabasePrivate from a pointer to the C GTlsDatabasePrivate that represents the TlsDatabasePrivate.
func TlsDatabasePrivateNewFromC(native unsafe.Pointer) *TlsDatabasePrivate {
	return &TlsDatabasePrivate{native: native}
}

// TlsFileDatabaseInterface is a representation of the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsFileDatabaseInterface that represents the TlsFileDatabaseInterface.
func (recv *TlsFileDatabaseInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsFileDatabaseInterfaceNewFromC creates a new TlsFileDatabaseInterface from a pointer to the C GTlsFileDatabaseInterface that represents the TlsFileDatabaseInterface.
func TlsFileDatabaseInterfaceNewFromC(native unsafe.Pointer) *TlsFileDatabaseInterface {
	return &TlsFileDatabaseInterface{native: native}
}

// TlsInteractionClass is a representation of the C record GTlsInteractionClass.
//
// since 2.30
type TlsInteractionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteractionClass that represents the TlsInteractionClass.
func (recv *TlsInteractionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionClassNewFromC creates a new TlsInteractionClass from a pointer to the C GTlsInteractionClass that represents the TlsInteractionClass.
func TlsInteractionClassNewFromC(native unsafe.Pointer) *TlsInteractionClass {
	return &TlsInteractionClass{native: native}
}

// TlsInteractionPrivate is a representation of the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteractionPrivate that represents the TlsInteractionPrivate.
func (recv *TlsInteractionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionPrivateNewFromC creates a new TlsInteractionPrivate from a pointer to the C GTlsInteractionPrivate that represents the TlsInteractionPrivate.
func TlsInteractionPrivateNewFromC(native unsafe.Pointer) *TlsInteractionPrivate {
	return &TlsInteractionPrivate{native: native}
}

// TlsPasswordClass is a representation of the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordClass that represents the TlsPasswordClass.
func (recv *TlsPasswordClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordClassNewFromC creates a new TlsPasswordClass from a pointer to the C GTlsPasswordClass that represents the TlsPasswordClass.
func TlsPasswordClassNewFromC(native unsafe.Pointer) *TlsPasswordClass {
	return &TlsPasswordClass{native: native}
}

// TlsPasswordPrivate is a representation of the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordPrivate that represents the TlsPasswordPrivate.
func (recv *TlsPasswordPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordPrivateNewFromC creates a new TlsPasswordPrivate from a pointer to the C GTlsPasswordPrivate that represents the TlsPasswordPrivate.
func TlsPasswordPrivateNewFromC(native unsafe.Pointer) *TlsPasswordPrivate {
	return &TlsPasswordPrivate{native: native}
}

// TlsServerConnectionInterface is a representation of the C record GTlsServerConnectionInterface.
//
// since 2.26
type TlsServerConnectionInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsServerConnectionInterface that represents the TlsServerConnectionInterface.
func (recv *TlsServerConnectionInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsServerConnectionInterfaceNewFromC creates a new TlsServerConnectionInterface from a pointer to the C GTlsServerConnectionInterface that represents the TlsServerConnectionInterface.
func TlsServerConnectionInterfaceNewFromC(native unsafe.Pointer) *TlsServerConnectionInterface {
	return &TlsServerConnectionInterface{native: native}
}

// UnixConnectionClass is a representation of the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionClass that represents the UnixConnectionClass.
func (recv *UnixConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionClassNewFromC creates a new UnixConnectionClass from a pointer to the C GUnixConnectionClass that represents the UnixConnectionClass.
func UnixConnectionClassNewFromC(native unsafe.Pointer) *UnixConnectionClass {
	return &UnixConnectionClass{native: native}
}

// UnixConnectionPrivate is a representation of the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionPrivate that represents the UnixConnectionPrivate.
func (recv *UnixConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionPrivateNewFromC creates a new UnixConnectionPrivate from a pointer to the C GUnixConnectionPrivate that represents the UnixConnectionPrivate.
func UnixConnectionPrivateNewFromC(native unsafe.Pointer) *UnixConnectionPrivate {
	return &UnixConnectionPrivate{native: native}
}

// UnixCredentialsMessageClass is a representation of the C record GUnixCredentialsMessageClass.
//
// since 2.26
type UnixCredentialsMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessageClass that represents the UnixCredentialsMessageClass.
func (recv *UnixCredentialsMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessageClassNewFromC creates a new UnixCredentialsMessageClass from a pointer to the C GUnixCredentialsMessageClass that represents the UnixCredentialsMessageClass.
func UnixCredentialsMessageClassNewFromC(native unsafe.Pointer) *UnixCredentialsMessageClass {
	return &UnixCredentialsMessageClass{native: native}
}

// UnixCredentialsMessagePrivate is a representation of the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessagePrivate that represents the UnixCredentialsMessagePrivate.
func (recv *UnixCredentialsMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessagePrivateNewFromC creates a new UnixCredentialsMessagePrivate from a pointer to the C GUnixCredentialsMessagePrivate that represents the UnixCredentialsMessagePrivate.
func UnixCredentialsMessagePrivateNewFromC(native unsafe.Pointer) *UnixCredentialsMessagePrivate {
	return &UnixCredentialsMessagePrivate{native: native}
}

// UnixFDListClass is a representation of the C record GUnixFDListClass.
type UnixFDListClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListClass that represents the UnixFDListClass.
func (recv *UnixFDListClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListClassNewFromC creates a new UnixFDListClass from a pointer to the C GUnixFDListClass that represents the UnixFDListClass.
func UnixFDListClassNewFromC(native unsafe.Pointer) *UnixFDListClass {
	return &UnixFDListClass{native: native}
}

// UnixFDListPrivate is a representation of the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListPrivate that represents the UnixFDListPrivate.
func (recv *UnixFDListPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListPrivateNewFromC creates a new UnixFDListPrivate from a pointer to the C GUnixFDListPrivate that represents the UnixFDListPrivate.
func UnixFDListPrivateNewFromC(native unsafe.Pointer) *UnixFDListPrivate {
	return &UnixFDListPrivate{native: native}
}

// UnixFDMessageClass is a representation of the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessageClass that represents the UnixFDMessageClass.
func (recv *UnixFDMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessageClassNewFromC creates a new UnixFDMessageClass from a pointer to the C GUnixFDMessageClass that represents the UnixFDMessageClass.
func UnixFDMessageClassNewFromC(native unsafe.Pointer) *UnixFDMessageClass {
	return &UnixFDMessageClass{native: native}
}

// UnixFDMessagePrivate is a representation of the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessagePrivate that represents the UnixFDMessagePrivate.
func (recv *UnixFDMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessagePrivateNewFromC creates a new UnixFDMessagePrivate from a pointer to the C GUnixFDMessagePrivate that represents the UnixFDMessagePrivate.
func UnixFDMessagePrivateNewFromC(native unsafe.Pointer) *UnixFDMessagePrivate {
	return &UnixFDMessagePrivate{native: native}
}

// UnixInputStreamClass is a representation of the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamClass that represents the UnixInputStreamClass.
func (recv *UnixInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamClassNewFromC creates a new UnixInputStreamClass from a pointer to the C GUnixInputStreamClass that represents the UnixInputStreamClass.
func UnixInputStreamClassNewFromC(native unsafe.Pointer) *UnixInputStreamClass {
	return &UnixInputStreamClass{native: native}
}

// UnixInputStreamPrivate is a representation of the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamPrivate that represents the UnixInputStreamPrivate.
func (recv *UnixInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamPrivateNewFromC creates a new UnixInputStreamPrivate from a pointer to the C GUnixInputStreamPrivate that represents the UnixInputStreamPrivate.
func UnixInputStreamPrivateNewFromC(native unsafe.Pointer) *UnixInputStreamPrivate {
	return &UnixInputStreamPrivate{native: native}
}

// UnixMountEntry is a representation of the C record GUnixMountEntry.
type UnixMountEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountEntry that represents the UnixMountEntry.
func (recv *UnixMountEntry) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountEntryNewFromC creates a new UnixMountEntry from a pointer to the C GUnixMountEntry that represents the UnixMountEntry.
func UnixMountEntryNewFromC(native unsafe.Pointer) *UnixMountEntry {
	return &UnixMountEntry{native: native}
}

// UnixMountMonitorClass is a representation of the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitorClass that represents the UnixMountMonitorClass.
func (recv *UnixMountMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorClassNewFromC creates a new UnixMountMonitorClass from a pointer to the C GUnixMountMonitorClass that represents the UnixMountMonitorClass.
func UnixMountMonitorClassNewFromC(native unsafe.Pointer) *UnixMountMonitorClass {
	return &UnixMountMonitorClass{native: native}
}

// UnixMountPoint is a representation of the C record GUnixMountPoint.
type UnixMountPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountPoint that represents the UnixMountPoint.
func (recv *UnixMountPoint) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountPointNewFromC creates a new UnixMountPoint from a pointer to the C GUnixMountPoint that represents the UnixMountPoint.
func UnixMountPointNewFromC(native unsafe.Pointer) *UnixMountPoint {
	return &UnixMountPoint{native: native}
}

// UnixOutputStreamClass is a representation of the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamClass that represents the UnixOutputStreamClass.
func (recv *UnixOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamClassNewFromC creates a new UnixOutputStreamClass from a pointer to the C GUnixOutputStreamClass that represents the UnixOutputStreamClass.
func UnixOutputStreamClassNewFromC(native unsafe.Pointer) *UnixOutputStreamClass {
	return &UnixOutputStreamClass{native: native}
}

// UnixOutputStreamPrivate is a representation of the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamPrivate that represents the UnixOutputStreamPrivate.
func (recv *UnixOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamPrivateNewFromC creates a new UnixOutputStreamPrivate from a pointer to the C GUnixOutputStreamPrivate that represents the UnixOutputStreamPrivate.
func UnixOutputStreamPrivateNewFromC(native unsafe.Pointer) *UnixOutputStreamPrivate {
	return &UnixOutputStreamPrivate{native: native}
}

// UnixSocketAddressClass is a representation of the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressClass that represents the UnixSocketAddressClass.
func (recv *UnixSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressClassNewFromC creates a new UnixSocketAddressClass from a pointer to the C GUnixSocketAddressClass that represents the UnixSocketAddressClass.
func UnixSocketAddressClassNewFromC(native unsafe.Pointer) *UnixSocketAddressClass {
	return &UnixSocketAddressClass{native: native}
}

// UnixSocketAddressPrivate is a representation of the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressPrivate that represents the UnixSocketAddressPrivate.
func (recv *UnixSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressPrivateNewFromC creates a new UnixSocketAddressPrivate from a pointer to the C GUnixSocketAddressPrivate that represents the UnixSocketAddressPrivate.
func UnixSocketAddressPrivateNewFromC(native unsafe.Pointer) *UnixSocketAddressPrivate {
	return &UnixSocketAddressPrivate{native: native}
}

// VfsClass is a representation of the C record GVfsClass.
type VfsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfsClass that represents the VfsClass.
func (recv *VfsClass) ToC() unsafe.Pointer {
	return recv.native
}

// VfsClassNewFromC creates a new VfsClass from a pointer to the C GVfsClass that represents the VfsClass.
func VfsClassNewFromC(native unsafe.Pointer) *VfsClass {
	return &VfsClass{native: native}
}

// VolumeIface is a representation of the C record GVolumeIface.
type VolumeIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeIface that represents the VolumeIface.
func (recv *VolumeIface) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeIfaceNewFromC creates a new VolumeIface from a pointer to the C GVolumeIface that represents the VolumeIface.
func VolumeIfaceNewFromC(native unsafe.Pointer) *VolumeIface {
	return &VolumeIface{native: native}
}

// VolumeMonitorClass is a representation of the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitorClass that represents the VolumeMonitorClass.
func (recv *VolumeMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitorClassNewFromC creates a new VolumeMonitorClass from a pointer to the C GVolumeMonitorClass that represents the VolumeMonitorClass.
func VolumeMonitorClassNewFromC(native unsafe.Pointer) *VolumeMonitorClass {
	return &VolumeMonitorClass{native: native}
}

// ZlibCompressorClass is a representation of the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressorClass that represents the ZlibCompressorClass.
func (recv *ZlibCompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressorClassNewFromC creates a new ZlibCompressorClass from a pointer to the C GZlibCompressorClass that represents the ZlibCompressorClass.
func ZlibCompressorClassNewFromC(native unsafe.Pointer) *ZlibCompressorClass {
	return &ZlibCompressorClass{native: native}
}

// ZlibDecompressorClass is a representation of the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressorClass that represents the ZlibDecompressorClass.
func (recv *ZlibDecompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressorClassNewFromC creates a new ZlibDecompressorClass from a pointer to the C GZlibDecompressorClass that represents the ZlibDecompressorClass.
func ZlibDecompressorClassNewFromC(native unsafe.Pointer) *ZlibDecompressorClass {
	return &ZlibDecompressorClass{native: native}
}

// AppInfoMonitor is a representation of the C record GAppInfoMonitor.
//
// since 2.40
type AppInfoMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfoMonitor that represents the AppInfoMonitor.
func (recv *AppInfoMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoMonitorNewFromC creates a new AppInfoMonitor from a pointer to the C GAppInfoMonitor that represents the AppInfoMonitor.
func AppInfoMonitorNewFromC(native unsafe.Pointer) *AppInfoMonitor {
	return &AppInfoMonitor{native: native}
}

// AppLaunchContext is a representation of the C record GAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContext that represents the AppLaunchContext.
func (recv *AppLaunchContext) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextNewFromC creates a new AppLaunchContext from a pointer to the C GAppLaunchContext that represents the AppLaunchContext.
func AppLaunchContextNewFromC(native unsafe.Pointer) *AppLaunchContext {
	return &AppLaunchContext{native: native}
}

// Application is a representation of the C record GApplication.
//
// since 2.28
type Application struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplication that represents the Application.
func (recv *Application) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationNewFromC creates a new Application from a pointer to the C GApplication that represents the Application.
func ApplicationNewFromC(native unsafe.Pointer) *Application {
	return &Application{native: native}
}

// ApplicationCommandLine is a representation of the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLine that represents the ApplicationCommandLine.
func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLineNewFromC creates a new ApplicationCommandLine from a pointer to the C GApplicationCommandLine that represents the ApplicationCommandLine.
func ApplicationCommandLineNewFromC(native unsafe.Pointer) *ApplicationCommandLine {
	return &ApplicationCommandLine{native: native}
}

// BufferedInputStream is a representation of the C record GBufferedInputStream.
type BufferedInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStream that represents the BufferedInputStream.
func (recv *BufferedInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamNewFromC creates a new BufferedInputStream from a pointer to the C GBufferedInputStream that represents the BufferedInputStream.
func BufferedInputStreamNewFromC(native unsafe.Pointer) *BufferedInputStream {
	return &BufferedInputStream{native: native}
}

// BufferedOutputStream is a representation of the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStream that represents the BufferedOutputStream.
func (recv *BufferedOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamNewFromC creates a new BufferedOutputStream from a pointer to the C GBufferedOutputStream that represents the BufferedOutputStream.
func BufferedOutputStreamNewFromC(native unsafe.Pointer) *BufferedOutputStream {
	return &BufferedOutputStream{native: native}
}

// BytesIcon is a representation of the C record GBytesIcon.
//
// since 2.38
type BytesIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBytesIcon that represents the BytesIcon.
func (recv *BytesIcon) ToC() unsafe.Pointer {
	return recv.native
}

// BytesIconNewFromC creates a new BytesIcon from a pointer to the C GBytesIcon that represents the BytesIcon.
func BytesIconNewFromC(native unsafe.Pointer) *BytesIcon {
	return &BytesIcon{native: native}
}

// Cancellable is a representation of the C record GCancellable.
type Cancellable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellable that represents the Cancellable.
func (recv *Cancellable) ToC() unsafe.Pointer {
	return recv.native
}

// CancellableNewFromC creates a new Cancellable from a pointer to the C GCancellable that represents the Cancellable.
func CancellableNewFromC(native unsafe.Pointer) *Cancellable {
	return &Cancellable{native: native}
}

// CharsetConverter is a representation of the C record GCharsetConverter.
type CharsetConverter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverter that represents the CharsetConverter.
func (recv *CharsetConverter) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverterNewFromC creates a new CharsetConverter from a pointer to the C GCharsetConverter that represents the CharsetConverter.
func CharsetConverterNewFromC(native unsafe.Pointer) *CharsetConverter {
	return &CharsetConverter{native: native}
}

// ConverterInputStream is a representation of the C record GConverterInputStream.
type ConverterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStream that represents the ConverterInputStream.
func (recv *ConverterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamNewFromC creates a new ConverterInputStream from a pointer to the C GConverterInputStream that represents the ConverterInputStream.
func ConverterInputStreamNewFromC(native unsafe.Pointer) *ConverterInputStream {
	return &ConverterInputStream{native: native}
}

// ConverterOutputStream is a representation of the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStream that represents the ConverterOutputStream.
func (recv *ConverterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamNewFromC creates a new ConverterOutputStream from a pointer to the C GConverterOutputStream that represents the ConverterOutputStream.
func ConverterOutputStreamNewFromC(native unsafe.Pointer) *ConverterOutputStream {
	return &ConverterOutputStream{native: native}
}

// Credentials is a representation of the C record GCredentials.
//
// since 2.26
type Credentials struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCredentials that represents the Credentials.
func (recv *Credentials) ToC() unsafe.Pointer {
	return recv.native
}

// CredentialsNewFromC creates a new Credentials from a pointer to the C GCredentials that represents the Credentials.
func CredentialsNewFromC(native unsafe.Pointer) *Credentials {
	return &Credentials{native: native}
}

// DBusActionGroup is a representation of the C record GDBusActionGroup.
type DBusActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusActionGroup that represents the DBusActionGroup.
func (recv *DBusActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// DBusActionGroupNewFromC creates a new DBusActionGroup from a pointer to the C GDBusActionGroup that represents the DBusActionGroup.
func DBusActionGroupNewFromC(native unsafe.Pointer) *DBusActionGroup {
	return &DBusActionGroup{native: native}
}

// DBusAuthObserver is a representation of the C record GDBusAuthObserver.
//
// since 2.26
type DBusAuthObserver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusAuthObserver that represents the DBusAuthObserver.
func (recv *DBusAuthObserver) ToC() unsafe.Pointer {
	return recv.native
}

// DBusAuthObserverNewFromC creates a new DBusAuthObserver from a pointer to the C GDBusAuthObserver that represents the DBusAuthObserver.
func DBusAuthObserverNewFromC(native unsafe.Pointer) *DBusAuthObserver {
	return &DBusAuthObserver{native: native}
}

// DBusConnection is a representation of the C record GDBusConnection.
//
// since 2.26
type DBusConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusConnection that represents the DBusConnection.
func (recv *DBusConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DBusConnectionNewFromC creates a new DBusConnection from a pointer to the C GDBusConnection that represents the DBusConnection.
func DBusConnectionNewFromC(native unsafe.Pointer) *DBusConnection {
	return &DBusConnection{native: native}
}

// DBusInterfaceSkeleton is a representation of the C record GDBusInterfaceSkeleton.
//
// since 2.30
type DBusInterfaceSkeleton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeleton that represents the DBusInterfaceSkeleton.
func (recv *DBusInterfaceSkeleton) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonNewFromC creates a new DBusInterfaceSkeleton from a pointer to the C GDBusInterfaceSkeleton that represents the DBusInterfaceSkeleton.
func DBusInterfaceSkeletonNewFromC(native unsafe.Pointer) *DBusInterfaceSkeleton {
	return &DBusInterfaceSkeleton{native: native}
}

// DBusMenuModel is a representation of the C record GDBusMenuModel.
type DBusMenuModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMenuModel that represents the DBusMenuModel.
func (recv *DBusMenuModel) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMenuModelNewFromC creates a new DBusMenuModel from a pointer to the C GDBusMenuModel that represents the DBusMenuModel.
func DBusMenuModelNewFromC(native unsafe.Pointer) *DBusMenuModel {
	return &DBusMenuModel{native: native}
}

// DBusMessage is a representation of the C record GDBusMessage.
//
// since 2.26
type DBusMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMessage that represents the DBusMessage.
func (recv *DBusMessage) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMessageNewFromC creates a new DBusMessage from a pointer to the C GDBusMessage that represents the DBusMessage.
func DBusMessageNewFromC(native unsafe.Pointer) *DBusMessage {
	return &DBusMessage{native: native}
}

// DBusMethodInvocation is a representation of the C record GDBusMethodInvocation.
//
// since 2.26
type DBusMethodInvocation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMethodInvocation that represents the DBusMethodInvocation.
func (recv *DBusMethodInvocation) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMethodInvocationNewFromC creates a new DBusMethodInvocation from a pointer to the C GDBusMethodInvocation that represents the DBusMethodInvocation.
func DBusMethodInvocationNewFromC(native unsafe.Pointer) *DBusMethodInvocation {
	return &DBusMethodInvocation{native: native}
}

// DBusObjectManagerClient is a representation of the C record GDBusObjectManagerClient.
//
// since 2.30
type DBusObjectManagerClient struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClient that represents the DBusObjectManagerClient.
func (recv *DBusObjectManagerClient) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientNewFromC creates a new DBusObjectManagerClient from a pointer to the C GDBusObjectManagerClient that represents the DBusObjectManagerClient.
func DBusObjectManagerClientNewFromC(native unsafe.Pointer) *DBusObjectManagerClient {
	return &DBusObjectManagerClient{native: native}
}

// DBusObjectManagerServer is a representation of the C record GDBusObjectManagerServer.
//
// since 2.30
type DBusObjectManagerServer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServer that represents the DBusObjectManagerServer.
func (recv *DBusObjectManagerServer) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerNewFromC creates a new DBusObjectManagerServer from a pointer to the C GDBusObjectManagerServer that represents the DBusObjectManagerServer.
func DBusObjectManagerServerNewFromC(native unsafe.Pointer) *DBusObjectManagerServer {
	return &DBusObjectManagerServer{native: native}
}

// DBusObjectProxy is a representation of the C record GDBusObjectProxy.
//
// since 2.30
type DBusObjectProxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxy that represents the DBusObjectProxy.
func (recv *DBusObjectProxy) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyNewFromC creates a new DBusObjectProxy from a pointer to the C GDBusObjectProxy that represents the DBusObjectProxy.
func DBusObjectProxyNewFromC(native unsafe.Pointer) *DBusObjectProxy {
	return &DBusObjectProxy{native: native}
}

// DBusObjectSkeleton is a representation of the C record GDBusObjectSkeleton.
//
// since 2.30
type DBusObjectSkeleton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeleton that represents the DBusObjectSkeleton.
func (recv *DBusObjectSkeleton) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonNewFromC creates a new DBusObjectSkeleton from a pointer to the C GDBusObjectSkeleton that represents the DBusObjectSkeleton.
func DBusObjectSkeletonNewFromC(native unsafe.Pointer) *DBusObjectSkeleton {
	return &DBusObjectSkeleton{native: native}
}

// DBusProxy is a representation of the C record GDBusProxy.
//
// since 2.26
type DBusProxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxy that represents the DBusProxy.
func (recv *DBusProxy) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyNewFromC creates a new DBusProxy from a pointer to the C GDBusProxy that represents the DBusProxy.
func DBusProxyNewFromC(native unsafe.Pointer) *DBusProxy {
	return &DBusProxy{native: native}
}

// DBusServer is a representation of the C record GDBusServer.
//
// since 2.26
type DBusServer struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusServer that represents the DBusServer.
func (recv *DBusServer) ToC() unsafe.Pointer {
	return recv.native
}

// DBusServerNewFromC creates a new DBusServer from a pointer to the C GDBusServer that represents the DBusServer.
func DBusServerNewFromC(native unsafe.Pointer) *DBusServer {
	return &DBusServer{native: native}
}

// DataInputStream is a representation of the C record GDataInputStream.
type DataInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStream that represents the DataInputStream.
func (recv *DataInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamNewFromC creates a new DataInputStream from a pointer to the C GDataInputStream that represents the DataInputStream.
func DataInputStreamNewFromC(native unsafe.Pointer) *DataInputStream {
	return &DataInputStream{native: native}
}

// DataOutputStream is a representation of the C record GDataOutputStream.
type DataOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStream that represents the DataOutputStream.
func (recv *DataOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamNewFromC creates a new DataOutputStream from a pointer to the C GDataOutputStream that represents the DataOutputStream.
func DataOutputStreamNewFromC(native unsafe.Pointer) *DataOutputStream {
	return &DataOutputStream{native: native}
}

// DesktopAppInfo is a representation of the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfo that represents the DesktopAppInfo.
func (recv *DesktopAppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoNewFromC creates a new DesktopAppInfo from a pointer to the C GDesktopAppInfo that represents the DesktopAppInfo.
func DesktopAppInfoNewFromC(native unsafe.Pointer) *DesktopAppInfo {
	return &DesktopAppInfo{native: native}
}

// Emblem is a representation of the C record GEmblem.
type Emblem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblem that represents the Emblem.
func (recv *Emblem) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemNewFromC creates a new Emblem from a pointer to the C GEmblem that represents the Emblem.
func EmblemNewFromC(native unsafe.Pointer) *Emblem {
	return &Emblem{native: native}
}

// EmblemedIcon is a representation of the C record GEmblemedIcon.
type EmblemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIcon that represents the EmblemedIcon.
func (recv *EmblemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconNewFromC creates a new EmblemedIcon from a pointer to the C GEmblemedIcon that represents the EmblemedIcon.
func EmblemedIconNewFromC(native unsafe.Pointer) *EmblemedIcon {
	return &EmblemedIcon{native: native}
}

// FileEnumerator is a representation of the C record GFileEnumerator.
type FileEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumerator that represents the FileEnumerator.
func (recv *FileEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorNewFromC creates a new FileEnumerator from a pointer to the C GFileEnumerator that represents the FileEnumerator.
func FileEnumeratorNewFromC(native unsafe.Pointer) *FileEnumerator {
	return &FileEnumerator{native: native}
}

// FileIOStream is a representation of the C record GFileIOStream.
//
// since 2.22
type FileIOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStream that represents the FileIOStream.
func (recv *FileIOStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamNewFromC creates a new FileIOStream from a pointer to the C GFileIOStream that represents the FileIOStream.
func FileIOStreamNewFromC(native unsafe.Pointer) *FileIOStream {
	return &FileIOStream{native: native}
}

// FileIcon is a representation of the C record GFileIcon.
type FileIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIcon that represents the FileIcon.
func (recv *FileIcon) ToC() unsafe.Pointer {
	return recv.native
}

// FileIconNewFromC creates a new FileIcon from a pointer to the C GFileIcon that represents the FileIcon.
func FileIconNewFromC(native unsafe.Pointer) *FileIcon {
	return &FileIcon{native: native}
}

// FileInfo is a representation of the C record GFileInfo.
type FileInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfo that represents the FileInfo.
func (recv *FileInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfoNewFromC creates a new FileInfo from a pointer to the C GFileInfo that represents the FileInfo.
func FileInfoNewFromC(native unsafe.Pointer) *FileInfo {
	return &FileInfo{native: native}
}

// FileInputStream is a representation of the C record GFileInputStream.
type FileInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStream that represents the FileInputStream.
func (recv *FileInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamNewFromC creates a new FileInputStream from a pointer to the C GFileInputStream that represents the FileInputStream.
func FileInputStreamNewFromC(native unsafe.Pointer) *FileInputStream {
	return &FileInputStream{native: native}
}

// FileMonitor is a representation of the C record GFileMonitor.
type FileMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitor that represents the FileMonitor.
func (recv *FileMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorNewFromC creates a new FileMonitor from a pointer to the C GFileMonitor that represents the FileMonitor.
func FileMonitorNewFromC(native unsafe.Pointer) *FileMonitor {
	return &FileMonitor{native: native}
}

// FileOutputStream is a representation of the C record GFileOutputStream.
type FileOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStream that represents the FileOutputStream.
func (recv *FileOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamNewFromC creates a new FileOutputStream from a pointer to the C GFileOutputStream that represents the FileOutputStream.
func FileOutputStreamNewFromC(native unsafe.Pointer) *FileOutputStream {
	return &FileOutputStream{native: native}
}

// FilenameCompleter is a representation of the C record GFilenameCompleter.
type FilenameCompleter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleter that represents the FilenameCompleter.
func (recv *FilenameCompleter) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleterNewFromC creates a new FilenameCompleter from a pointer to the C GFilenameCompleter that represents the FilenameCompleter.
func FilenameCompleterNewFromC(native unsafe.Pointer) *FilenameCompleter {
	return &FilenameCompleter{native: native}
}

// FilterInputStream is a representation of the C record GFilterInputStream.
type FilterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStream that represents the FilterInputStream.
func (recv *FilterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStreamNewFromC creates a new FilterInputStream from a pointer to the C GFilterInputStream that represents the FilterInputStream.
func FilterInputStreamNewFromC(native unsafe.Pointer) *FilterInputStream {
	return &FilterInputStream{native: native}
}

// FilterOutputStream is a representation of the C record GFilterOutputStream.
type FilterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStream that represents the FilterOutputStream.
func (recv *FilterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStreamNewFromC creates a new FilterOutputStream from a pointer to the C GFilterOutputStream that represents the FilterOutputStream.
func FilterOutputStreamNewFromC(native unsafe.Pointer) *FilterOutputStream {
	return &FilterOutputStream{native: native}
}

// IOModule is a representation of the C record GIOModule.
type IOModule struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModule that represents the IOModule.
func (recv *IOModule) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleNewFromC creates a new IOModule from a pointer to the C GIOModule that represents the IOModule.
func IOModuleNewFromC(native unsafe.Pointer) *IOModule {
	return &IOModule{native: native}
}

// IOStream is a representation of the C record GIOStream.
//
// since 2.22
type IOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStream that represents the IOStream.
func (recv *IOStream) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamNewFromC creates a new IOStream from a pointer to the C GIOStream that represents the IOStream.
func IOStreamNewFromC(native unsafe.Pointer) *IOStream {
	return &IOStream{native: native}
}

// InetAddress is a representation of the C record GInetAddress.
type InetAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddress that represents the InetAddress.
func (recv *InetAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressNewFromC creates a new InetAddress from a pointer to the C GInetAddress that represents the InetAddress.
func InetAddressNewFromC(native unsafe.Pointer) *InetAddress {
	return &InetAddress{native: native}
}

// InetAddressMask is a representation of the C record GInetAddressMask.
//
// since 2.32
type InetAddressMask struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMask that represents the InetAddressMask.
func (recv *InetAddressMask) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskNewFromC creates a new InetAddressMask from a pointer to the C GInetAddressMask that represents the InetAddressMask.
func InetAddressMaskNewFromC(native unsafe.Pointer) *InetAddressMask {
	return &InetAddressMask{native: native}
}

// InetSocketAddress is a representation of the C record GInetSocketAddress.
type InetSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddress that represents the InetSocketAddress.
func (recv *InetSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressNewFromC creates a new InetSocketAddress from a pointer to the C GInetSocketAddress that represents the InetSocketAddress.
func InetSocketAddressNewFromC(native unsafe.Pointer) *InetSocketAddress {
	return &InetSocketAddress{native: native}
}

// InputStream is a representation of the C record GInputStream.
type InputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStream that represents the InputStream.
func (recv *InputStream) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamNewFromC creates a new InputStream from a pointer to the C GInputStream that represents the InputStream.
func InputStreamNewFromC(native unsafe.Pointer) *InputStream {
	return &InputStream{native: native}
}

// ListStore is a representation of the C record GListStore.
type ListStore struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStore that represents the ListStore.
func (recv *ListStore) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreNewFromC creates a new ListStore from a pointer to the C GListStore that represents the ListStore.
func ListStoreNewFromC(native unsafe.Pointer) *ListStore {
	return &ListStore{native: native}
}

// MemoryInputStream is a representation of the C record GMemoryInputStream.
type MemoryInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStream that represents the MemoryInputStream.
func (recv *MemoryInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamNewFromC creates a new MemoryInputStream from a pointer to the C GMemoryInputStream that represents the MemoryInputStream.
func MemoryInputStreamNewFromC(native unsafe.Pointer) *MemoryInputStream {
	return &MemoryInputStream{native: native}
}

// MemoryOutputStream is a representation of the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStream that represents the MemoryOutputStream.
func (recv *MemoryOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamNewFromC creates a new MemoryOutputStream from a pointer to the C GMemoryOutputStream that represents the MemoryOutputStream.
func MemoryOutputStreamNewFromC(native unsafe.Pointer) *MemoryOutputStream {
	return &MemoryOutputStream{native: native}
}

// Menu is a representation of the C record GMenu.
//
// since 2.32
type Menu struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenu that represents the Menu.
func (recv *Menu) ToC() unsafe.Pointer {
	return recv.native
}

// MenuNewFromC creates a new Menu from a pointer to the C GMenu that represents the Menu.
func MenuNewFromC(native unsafe.Pointer) *Menu {
	return &Menu{native: native}
}

// MenuAttributeIter is a representation of the C record GMenuAttributeIter.
//
// since 2.32
type MenuAttributeIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIter that represents the MenuAttributeIter.
func (recv *MenuAttributeIter) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterNewFromC creates a new MenuAttributeIter from a pointer to the C GMenuAttributeIter that represents the MenuAttributeIter.
func MenuAttributeIterNewFromC(native unsafe.Pointer) *MenuAttributeIter {
	return &MenuAttributeIter{native: native}
}

// MenuItem is a representation of the C record GMenuItem.
//
// since 2.32
type MenuItem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuItem that represents the MenuItem.
func (recv *MenuItem) ToC() unsafe.Pointer {
	return recv.native
}

// MenuItemNewFromC creates a new MenuItem from a pointer to the C GMenuItem that represents the MenuItem.
func MenuItemNewFromC(native unsafe.Pointer) *MenuItem {
	return &MenuItem{native: native}
}

// MenuLinkIter is a representation of the C record GMenuLinkIter.
//
// since 2.32
type MenuLinkIter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIter that represents the MenuLinkIter.
func (recv *MenuLinkIter) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterNewFromC creates a new MenuLinkIter from a pointer to the C GMenuLinkIter that represents the MenuLinkIter.
func MenuLinkIterNewFromC(native unsafe.Pointer) *MenuLinkIter {
	return &MenuLinkIter{native: native}
}

// MenuModel is a representation of the C record GMenuModel.
//
// since 2.32
type MenuModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModel that represents the MenuModel.
func (recv *MenuModel) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelNewFromC creates a new MenuModel from a pointer to the C GMenuModel that represents the MenuModel.
func MenuModelNewFromC(native unsafe.Pointer) *MenuModel {
	return &MenuModel{native: native}
}

// MountOperation is a representation of the C record GMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperation that represents the MountOperation.
func (recv *MountOperation) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationNewFromC creates a new MountOperation from a pointer to the C GMountOperation that represents the MountOperation.
func MountOperationNewFromC(native unsafe.Pointer) *MountOperation {
	return &MountOperation{native: native}
}

// NativeSocketAddress is a representation of the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeSocketAddress that represents the NativeSocketAddress.
func (recv *NativeSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NativeSocketAddressNewFromC creates a new NativeSocketAddress from a pointer to the C GNativeSocketAddress that represents the NativeSocketAddress.
func NativeSocketAddressNewFromC(native unsafe.Pointer) *NativeSocketAddress {
	return &NativeSocketAddress{native: native}
}

// NativeVolumeMonitor is a representation of the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeVolumeMonitor that represents the NativeVolumeMonitor.
func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// NativeVolumeMonitorNewFromC creates a new NativeVolumeMonitor from a pointer to the C GNativeVolumeMonitor that represents the NativeVolumeMonitor.
func NativeVolumeMonitorNewFromC(native unsafe.Pointer) *NativeVolumeMonitor {
	return &NativeVolumeMonitor{native: native}
}

// NetworkAddress is a representation of the C record GNetworkAddress.
type NetworkAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddress that represents the NetworkAddress.
func (recv *NetworkAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressNewFromC creates a new NetworkAddress from a pointer to the C GNetworkAddress that represents the NetworkAddress.
func NetworkAddressNewFromC(native unsafe.Pointer) *NetworkAddress {
	return &NetworkAddress{native: native}
}

// NetworkService is a representation of the C record GNetworkService.
type NetworkService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkService that represents the NetworkService.
func (recv *NetworkService) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServiceNewFromC creates a new NetworkService from a pointer to the C GNetworkService that represents the NetworkService.
func NetworkServiceNewFromC(native unsafe.Pointer) *NetworkService {
	return &NetworkService{native: native}
}

// Notification is a representation of the C record GNotification.
//
// since 2.40
type Notification struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNotification that represents the Notification.
func (recv *Notification) ToC() unsafe.Pointer {
	return recv.native
}

// NotificationNewFromC creates a new Notification from a pointer to the C GNotification that represents the Notification.
func NotificationNewFromC(native unsafe.Pointer) *Notification {
	return &Notification{native: native}
}

// OutputStream is a representation of the C record GOutputStream.
type OutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStream that represents the OutputStream.
func (recv *OutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamNewFromC creates a new OutputStream from a pointer to the C GOutputStream that represents the OutputStream.
func OutputStreamNewFromC(native unsafe.Pointer) *OutputStream {
	return &OutputStream{native: native}
}

// Permission is a representation of the C record GPermission.
type Permission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermission that represents the Permission.
func (recv *Permission) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionNewFromC creates a new Permission from a pointer to the C GPermission that represents the Permission.
func PermissionNewFromC(native unsafe.Pointer) *Permission {
	return &Permission{native: native}
}

// PropertyAction is a representation of the C record GPropertyAction.
//
// since 2.38
type PropertyAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPropertyAction that represents the PropertyAction.
func (recv *PropertyAction) ToC() unsafe.Pointer {
	return recv.native
}

// PropertyActionNewFromC creates a new PropertyAction from a pointer to the C GPropertyAction that represents the PropertyAction.
func PropertyActionNewFromC(native unsafe.Pointer) *PropertyAction {
	return &PropertyAction{native: native}
}

// ProxyAddress is a representation of the C record GProxyAddress.
//
// since 2.26
type ProxyAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddress that represents the ProxyAddress.
func (recv *ProxyAddress) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressNewFromC creates a new ProxyAddress from a pointer to the C GProxyAddress that represents the ProxyAddress.
func ProxyAddressNewFromC(native unsafe.Pointer) *ProxyAddress {
	return &ProxyAddress{native: native}
}

// ProxyAddressEnumerator is a representation of the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumerator that represents the ProxyAddressEnumerator.
func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorNewFromC creates a new ProxyAddressEnumerator from a pointer to the C GProxyAddressEnumerator that represents the ProxyAddressEnumerator.
func ProxyAddressEnumeratorNewFromC(native unsafe.Pointer) *ProxyAddressEnumerator {
	return &ProxyAddressEnumerator{native: native}
}

// Resolver is a representation of the C record GResolver.
type Resolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolver that represents the Resolver.
func (recv *Resolver) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverNewFromC creates a new Resolver from a pointer to the C GResolver that represents the Resolver.
func ResolverNewFromC(native unsafe.Pointer) *Resolver {
	return &Resolver{native: native}
}

// Settings is a representation of the C record GSettings.
type Settings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettings that represents the Settings.
func (recv *Settings) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsNewFromC creates a new Settings from a pointer to the C GSettings that represents the Settings.
func SettingsNewFromC(native unsafe.Pointer) *Settings {
	return &Settings{native: native}
}

// SettingsBackend is a representation of the C record GSettingsBackend.
type SettingsBackend struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsBackend that represents the SettingsBackend.
func (recv *SettingsBackend) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsBackendNewFromC creates a new SettingsBackend from a pointer to the C GSettingsBackend that represents the SettingsBackend.
func SettingsBackendNewFromC(native unsafe.Pointer) *SettingsBackend {
	return &SettingsBackend{native: native}
}

// SimpleAction is a representation of the C record GSimpleAction.
type SimpleAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAction that represents the SimpleAction.
func (recv *SimpleAction) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionNewFromC creates a new SimpleAction from a pointer to the C GSimpleAction that represents the SimpleAction.
func SimpleActionNewFromC(native unsafe.Pointer) *SimpleAction {
	return &SimpleAction{native: native}
}

// SimpleActionGroup is a representation of the C record GSimpleActionGroup.
//
// since 2.28
type SimpleActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroup that represents the SimpleActionGroup.
func (recv *SimpleActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupNewFromC creates a new SimpleActionGroup from a pointer to the C GSimpleActionGroup that represents the SimpleActionGroup.
func SimpleActionGroupNewFromC(native unsafe.Pointer) *SimpleActionGroup {
	return &SimpleActionGroup{native: native}
}

// SimpleAsyncResult is a representation of the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResult that represents the SimpleAsyncResult.
func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultNewFromC creates a new SimpleAsyncResult from a pointer to the C GSimpleAsyncResult that represents the SimpleAsyncResult.
func SimpleAsyncResultNewFromC(native unsafe.Pointer) *SimpleAsyncResult {
	return &SimpleAsyncResult{native: native}
}

// SimpleIOStream is a representation of the C record GSimpleIOStream.
//
// since 2.44
type SimpleIOStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleIOStream that represents the SimpleIOStream.
func (recv *SimpleIOStream) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleIOStreamNewFromC creates a new SimpleIOStream from a pointer to the C GSimpleIOStream that represents the SimpleIOStream.
func SimpleIOStreamNewFromC(native unsafe.Pointer) *SimpleIOStream {
	return &SimpleIOStream{native: native}
}

// SimplePermission is a representation of the C record GSimplePermission.
type SimplePermission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimplePermission that represents the SimplePermission.
func (recv *SimplePermission) ToC() unsafe.Pointer {
	return recv.native
}

// SimplePermissionNewFromC creates a new SimplePermission from a pointer to the C GSimplePermission that represents the SimplePermission.
func SimplePermissionNewFromC(native unsafe.Pointer) *SimplePermission {
	return &SimplePermission{native: native}
}

// SimpleProxyResolver is a representation of the C record GSimpleProxyResolver.
//
// since 2.36
type SimpleProxyResolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolver that represents the SimpleProxyResolver.
func (recv *SimpleProxyResolver) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverNewFromC creates a new SimpleProxyResolver from a pointer to the C GSimpleProxyResolver that represents the SimpleProxyResolver.
func SimpleProxyResolverNewFromC(native unsafe.Pointer) *SimpleProxyResolver {
	return &SimpleProxyResolver{native: native}
}

// Socket is a representation of the C record GSocket.
//
// since 2.22
type Socket struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocket that represents the Socket.
func (recv *Socket) ToC() unsafe.Pointer {
	return recv.native
}

// SocketNewFromC creates a new Socket from a pointer to the C GSocket that represents the Socket.
func SocketNewFromC(native unsafe.Pointer) *Socket {
	return &Socket{native: native}
}

// SocketAddress is a representation of the C record GSocketAddress.
type SocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddress that represents the SocketAddress.
func (recv *SocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressNewFromC creates a new SocketAddress from a pointer to the C GSocketAddress that represents the SocketAddress.
func SocketAddressNewFromC(native unsafe.Pointer) *SocketAddress {
	return &SocketAddress{native: native}
}

// SocketAddressEnumerator is a representation of the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumerator that represents the SocketAddressEnumerator.
func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumeratorNewFromC creates a new SocketAddressEnumerator from a pointer to the C GSocketAddressEnumerator that represents the SocketAddressEnumerator.
func SocketAddressEnumeratorNewFromC(native unsafe.Pointer) *SocketAddressEnumerator {
	return &SocketAddressEnumerator{native: native}
}

// SocketClient is a representation of the C record GSocketClient.
//
// since 2.22
type SocketClient struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClient that represents the SocketClient.
func (recv *SocketClient) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientNewFromC creates a new SocketClient from a pointer to the C GSocketClient that represents the SocketClient.
func SocketClientNewFromC(native unsafe.Pointer) *SocketClient {
	return &SocketClient{native: native}
}

// SocketConnection is a representation of the C record GSocketConnection.
//
// since 2.22
type SocketConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnection that represents the SocketConnection.
func (recv *SocketConnection) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionNewFromC creates a new SocketConnection from a pointer to the C GSocketConnection that represents the SocketConnection.
func SocketConnectionNewFromC(native unsafe.Pointer) *SocketConnection {
	return &SocketConnection{native: native}
}

// SocketControlMessage is a representation of the C record GSocketControlMessage.
//
// since 2.22
type SocketControlMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessage that represents the SocketControlMessage.
func (recv *SocketControlMessage) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessageNewFromC creates a new SocketControlMessage from a pointer to the C GSocketControlMessage that represents the SocketControlMessage.
func SocketControlMessageNewFromC(native unsafe.Pointer) *SocketControlMessage {
	return &SocketControlMessage{native: native}
}

// SocketListener is a representation of the C record GSocketListener.
//
// since 2.22
type SocketListener struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListener that represents the SocketListener.
func (recv *SocketListener) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerNewFromC creates a new SocketListener from a pointer to the C GSocketListener that represents the SocketListener.
func SocketListenerNewFromC(native unsafe.Pointer) *SocketListener {
	return &SocketListener{native: native}
}

// SocketService is a representation of the C record GSocketService.
//
// since 2.22
type SocketService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketService that represents the SocketService.
func (recv *SocketService) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServiceNewFromC creates a new SocketService from a pointer to the C GSocketService that represents the SocketService.
func SocketServiceNewFromC(native unsafe.Pointer) *SocketService {
	return &SocketService{native: native}
}

// Subprocess is a representation of the C record GSubprocess.
//
// since 2.40
type Subprocess struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSubprocess that represents the Subprocess.
func (recv *Subprocess) ToC() unsafe.Pointer {
	return recv.native
}

// SubprocessNewFromC creates a new Subprocess from a pointer to the C GSubprocess that represents the Subprocess.
func SubprocessNewFromC(native unsafe.Pointer) *Subprocess {
	return &Subprocess{native: native}
}

// SubprocessLauncher is a representation of the C record GSubprocessLauncher.
//
// since 2.40
type SubprocessLauncher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSubprocessLauncher that represents the SubprocessLauncher.
func (recv *SubprocessLauncher) ToC() unsafe.Pointer {
	return recv.native
}

// SubprocessLauncherNewFromC creates a new SubprocessLauncher from a pointer to the C GSubprocessLauncher that represents the SubprocessLauncher.
func SubprocessLauncherNewFromC(native unsafe.Pointer) *SubprocessLauncher {
	return &SubprocessLauncher{native: native}
}

// Task is a representation of the C record GTask.
type Task struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTask that represents the Task.
func (recv *Task) ToC() unsafe.Pointer {
	return recv.native
}

// TaskNewFromC creates a new Task from a pointer to the C GTask that represents the Task.
func TaskNewFromC(native unsafe.Pointer) *Task {
	return &Task{native: native}
}

// TcpConnection is a representation of the C record GTcpConnection.
//
// since 2.22
type TcpConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnection that represents the TcpConnection.
func (recv *TcpConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionNewFromC creates a new TcpConnection from a pointer to the C GTcpConnection that represents the TcpConnection.
func TcpConnectionNewFromC(native unsafe.Pointer) *TcpConnection {
	return &TcpConnection{native: native}
}

// TcpWrapperConnection is a representation of the C record GTcpWrapperConnection.
//
// since 2.28
type TcpWrapperConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnection that represents the TcpWrapperConnection.
func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionNewFromC creates a new TcpWrapperConnection from a pointer to the C GTcpWrapperConnection that represents the TcpWrapperConnection.
func TcpWrapperConnectionNewFromC(native unsafe.Pointer) *TcpWrapperConnection {
	return &TcpWrapperConnection{native: native}
}

// TestDBus is a representation of the C record GTestDBus.
//
// since 2.34
type TestDBus struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTestDBus that represents the TestDBus.
func (recv *TestDBus) ToC() unsafe.Pointer {
	return recv.native
}

// TestDBusNewFromC creates a new TestDBus from a pointer to the C GTestDBus that represents the TestDBus.
func TestDBusNewFromC(native unsafe.Pointer) *TestDBus {
	return &TestDBus{native: native}
}

// ThemedIcon is a representation of the C record GThemedIcon.
type ThemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIcon that represents the ThemedIcon.
func (recv *ThemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIconNewFromC creates a new ThemedIcon from a pointer to the C GThemedIcon that represents the ThemedIcon.
func ThemedIconNewFromC(native unsafe.Pointer) *ThemedIcon {
	return &ThemedIcon{native: native}
}

// ThreadedSocketService is a representation of the C record GThreadedSocketService.
//
// since 2.22
type ThreadedSocketService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketService that represents the ThreadedSocketService.
func (recv *ThreadedSocketService) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServiceNewFromC creates a new ThreadedSocketService from a pointer to the C GThreadedSocketService that represents the ThreadedSocketService.
func ThreadedSocketServiceNewFromC(native unsafe.Pointer) *ThreadedSocketService {
	return &ThreadedSocketService{native: native}
}

// TlsCertificate is a representation of the C record GTlsCertificate.
//
// since 2.28
type TlsCertificate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificate that represents the TlsCertificate.
func (recv *TlsCertificate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificateNewFromC creates a new TlsCertificate from a pointer to the C GTlsCertificate that represents the TlsCertificate.
func TlsCertificateNewFromC(native unsafe.Pointer) *TlsCertificate {
	return &TlsCertificate{native: native}
}

// TlsConnection is a representation of the C record GTlsConnection.
//
// since 2.28
type TlsConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnection that represents the TlsConnection.
func (recv *TlsConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionNewFromC creates a new TlsConnection from a pointer to the C GTlsConnection that represents the TlsConnection.
func TlsConnectionNewFromC(native unsafe.Pointer) *TlsConnection {
	return &TlsConnection{native: native}
}

// TlsDatabase is a representation of the C record GTlsDatabase.
//
// since 2.30
type TlsDatabase struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabase that represents the TlsDatabase.
func (recv *TlsDatabase) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabaseNewFromC creates a new TlsDatabase from a pointer to the C GTlsDatabase that represents the TlsDatabase.
func TlsDatabaseNewFromC(native unsafe.Pointer) *TlsDatabase {
	return &TlsDatabase{native: native}
}

// TlsInteraction is a representation of the C record GTlsInteraction.
//
// since 2.30
type TlsInteraction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteraction that represents the TlsInteraction.
func (recv *TlsInteraction) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionNewFromC creates a new TlsInteraction from a pointer to the C GTlsInteraction that represents the TlsInteraction.
func TlsInteractionNewFromC(native unsafe.Pointer) *TlsInteraction {
	return &TlsInteraction{native: native}
}

// TlsPassword is a representation of the C record GTlsPassword.
//
// since 2.30
type TlsPassword struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPassword that represents the TlsPassword.
func (recv *TlsPassword) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordNewFromC creates a new TlsPassword from a pointer to the C GTlsPassword that represents the TlsPassword.
func TlsPasswordNewFromC(native unsafe.Pointer) *TlsPassword {
	return &TlsPassword{native: native}
}

// UnixConnection is a representation of the C record GUnixConnection.
//
// since 2.22
type UnixConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnection that represents the UnixConnection.
func (recv *UnixConnection) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionNewFromC creates a new UnixConnection from a pointer to the C GUnixConnection that represents the UnixConnection.
func UnixConnectionNewFromC(native unsafe.Pointer) *UnixConnection {
	return &UnixConnection{native: native}
}

// UnixCredentialsMessage is a representation of the C record GUnixCredentialsMessage.
//
// since 2.26
type UnixCredentialsMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessage that represents the UnixCredentialsMessage.
func (recv *UnixCredentialsMessage) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessageNewFromC creates a new UnixCredentialsMessage from a pointer to the C GUnixCredentialsMessage that represents the UnixCredentialsMessage.
func UnixCredentialsMessageNewFromC(native unsafe.Pointer) *UnixCredentialsMessage {
	return &UnixCredentialsMessage{native: native}
}

// UnixFDList is a representation of the C record GUnixFDList.
type UnixFDList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDList that represents the UnixFDList.
func (recv *UnixFDList) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListNewFromC creates a new UnixFDList from a pointer to the C GUnixFDList that represents the UnixFDList.
func UnixFDListNewFromC(native unsafe.Pointer) *UnixFDList {
	return &UnixFDList{native: native}
}

// UnixFDMessage is a representation of the C record GUnixFDMessage.
type UnixFDMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessage that represents the UnixFDMessage.
func (recv *UnixFDMessage) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessageNewFromC creates a new UnixFDMessage from a pointer to the C GUnixFDMessage that represents the UnixFDMessage.
func UnixFDMessageNewFromC(native unsafe.Pointer) *UnixFDMessage {
	return &UnixFDMessage{native: native}
}

// UnixInputStream is a representation of the C record GUnixInputStream.
type UnixInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStream that represents the UnixInputStream.
func (recv *UnixInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamNewFromC creates a new UnixInputStream from a pointer to the C GUnixInputStream that represents the UnixInputStream.
func UnixInputStreamNewFromC(native unsafe.Pointer) *UnixInputStream {
	return &UnixInputStream{native: native}
}

// UnixMountMonitor is a representation of the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitor that represents the UnixMountMonitor.
func (recv *UnixMountMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorNewFromC creates a new UnixMountMonitor from a pointer to the C GUnixMountMonitor that represents the UnixMountMonitor.
func UnixMountMonitorNewFromC(native unsafe.Pointer) *UnixMountMonitor {
	return &UnixMountMonitor{native: native}
}

// UnixOutputStream is a representation of the C record GUnixOutputStream.
type UnixOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStream that represents the UnixOutputStream.
func (recv *UnixOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamNewFromC creates a new UnixOutputStream from a pointer to the C GUnixOutputStream that represents the UnixOutputStream.
func UnixOutputStreamNewFromC(native unsafe.Pointer) *UnixOutputStream {
	return &UnixOutputStream{native: native}
}

// UnixSocketAddress is a representation of the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddress that represents the UnixSocketAddress.
func (recv *UnixSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressNewFromC creates a new UnixSocketAddress from a pointer to the C GUnixSocketAddress that represents the UnixSocketAddress.
func UnixSocketAddressNewFromC(native unsafe.Pointer) *UnixSocketAddress {
	return &UnixSocketAddress{native: native}
}

// Vfs is a representation of the C record GVfs.
type Vfs struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfs that represents the Vfs.
func (recv *Vfs) ToC() unsafe.Pointer {
	return recv.native
}

// VfsNewFromC creates a new Vfs from a pointer to the C GVfs that represents the Vfs.
func VfsNewFromC(native unsafe.Pointer) *Vfs {
	return &Vfs{native: native}
}

// VolumeMonitor is a representation of the C record GVolumeMonitor.
type VolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitor that represents the VolumeMonitor.
func (recv *VolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitorNewFromC creates a new VolumeMonitor from a pointer to the C GVolumeMonitor that represents the VolumeMonitor.
func VolumeMonitorNewFromC(native unsafe.Pointer) *VolumeMonitor {
	return &VolumeMonitor{native: native}
}

// ZlibCompressor is a representation of the C record GZlibCompressor.
type ZlibCompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressor that represents the ZlibCompressor.
func (recv *ZlibCompressor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressorNewFromC creates a new ZlibCompressor from a pointer to the C GZlibCompressor that represents the ZlibCompressor.
func ZlibCompressorNewFromC(native unsafe.Pointer) *ZlibCompressor {
	return &ZlibCompressor{native: native}
}

// ZlibDecompressor is a representation of the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressor that represents the ZlibDecompressor.
func (recv *ZlibDecompressor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressorNewFromC creates a new ZlibDecompressor from a pointer to the C GZlibDecompressor that represents the ZlibDecompressor.
func ZlibDecompressorNewFromC(native unsafe.Pointer) *ZlibDecompressor {
	return &ZlibDecompressor{native: native}
}

// Action is a representation of the C interface GAction.
type Action struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAction that represents the Action.
func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionNewFromC creates a new Action from a pointer to the C GAction that represents the Action.
func ActionNewFromC(native unsafe.Pointer) *Action {
	return &Action{native: native}
}

// ActionGroup is a representation of the C interface GActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionGroup that represents the ActionGroup.
func (recv *ActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupNewFromC creates a new ActionGroup from a pointer to the C GActionGroup that represents the ActionGroup.
func ActionGroupNewFromC(native unsafe.Pointer) *ActionGroup {
	return &ActionGroup{native: native}
}

// ActionMap is a representation of the C interface GActionMap.
//
// since 2.32
type ActionMap struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionMap that represents the ActionMap.
func (recv *ActionMap) ToC() unsafe.Pointer {
	return recv.native
}

// ActionMapNewFromC creates a new ActionMap from a pointer to the C GActionMap that represents the ActionMap.
func ActionMapNewFromC(native unsafe.Pointer) *ActionMap {
	return &ActionMap{native: native}
}

// AppInfo is a representation of the C interface GAppInfo.
type AppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfo that represents the AppInfo.
func (recv *AppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoNewFromC creates a new AppInfo from a pointer to the C GAppInfo that represents the AppInfo.
func AppInfoNewFromC(native unsafe.Pointer) *AppInfo {
	return &AppInfo{native: native}
}

// AsyncInitable is a representation of the C interface GAsyncInitable.
//
// since 2.22
type AsyncInitable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncInitable that represents the AsyncInitable.
func (recv *AsyncInitable) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncInitableNewFromC creates a new AsyncInitable from a pointer to the C GAsyncInitable that represents the AsyncInitable.
func AsyncInitableNewFromC(native unsafe.Pointer) *AsyncInitable {
	return &AsyncInitable{native: native}
}

// AsyncResult is a representation of the C interface GAsyncResult.
type AsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResult that represents the AsyncResult.
func (recv *AsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResultNewFromC creates a new AsyncResult from a pointer to the C GAsyncResult that represents the AsyncResult.
func AsyncResultNewFromC(native unsafe.Pointer) *AsyncResult {
	return &AsyncResult{native: native}
}

// Converter is a representation of the C interface GConverter.
//
// since 2.24
type Converter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverter that represents the Converter.
func (recv *Converter) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterNewFromC creates a new Converter from a pointer to the C GConverter that represents the Converter.
func ConverterNewFromC(native unsafe.Pointer) *Converter {
	return &Converter{native: native}
}

// DBusInterface is a representation of the C interface GDBusInterface.
//
// since 2.30
type DBusInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterface that represents the DBusInterface.
func (recv *DBusInterface) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceNewFromC creates a new DBusInterface from a pointer to the C GDBusInterface that represents the DBusInterface.
func DBusInterfaceNewFromC(native unsafe.Pointer) *DBusInterface {
	return &DBusInterface{native: native}
}

// DBusObject is a representation of the C interface GDBusObject.
type DBusObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObject that represents the DBusObject.
func (recv *DBusObject) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectNewFromC creates a new DBusObject from a pointer to the C GDBusObject that represents the DBusObject.
func DBusObjectNewFromC(native unsafe.Pointer) *DBusObject {
	return &DBusObject{native: native}
}

// DBusObjectManager is a representation of the C interface GDBusObjectManager.
type DBusObjectManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManager that represents the DBusObjectManager.
func (recv *DBusObjectManager) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerNewFromC creates a new DBusObjectManager from a pointer to the C GDBusObjectManager that represents the DBusObjectManager.
func DBusObjectManagerNewFromC(native unsafe.Pointer) *DBusObjectManager {
	return &DBusObjectManager{native: native}
}

// DatagramBased is a representation of the C interface GDatagramBased.
//
// since 2.48
type DatagramBased struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDatagramBased that represents the DatagramBased.
func (recv *DatagramBased) ToC() unsafe.Pointer {
	return recv.native
}

// DatagramBasedNewFromC creates a new DatagramBased from a pointer to the C GDatagramBased that represents the DatagramBased.
func DatagramBasedNewFromC(native unsafe.Pointer) *DatagramBased {
	return &DatagramBased{native: native}
}

// DesktopAppInfoLookup is a representation of the C interface GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookup that represents the DesktopAppInfoLookup.
func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookupNewFromC creates a new DesktopAppInfoLookup from a pointer to the C GDesktopAppInfoLookup that represents the DesktopAppInfoLookup.
func DesktopAppInfoLookupNewFromC(native unsafe.Pointer) *DesktopAppInfoLookup {
	return &DesktopAppInfoLookup{native: native}
}

// Drive is a representation of the C interface GDrive.
type Drive struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDrive that represents the Drive.
func (recv *Drive) ToC() unsafe.Pointer {
	return recv.native
}

// DriveNewFromC creates a new Drive from a pointer to the C GDrive that represents the Drive.
func DriveNewFromC(native unsafe.Pointer) *Drive {
	return &Drive{native: native}
}

// DtlsClientConnection is a representation of the C interface GDtlsClientConnection.
//
// since 2.48
type DtlsClientConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsClientConnection that represents the DtlsClientConnection.
func (recv *DtlsClientConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsClientConnectionNewFromC creates a new DtlsClientConnection from a pointer to the C GDtlsClientConnection that represents the DtlsClientConnection.
func DtlsClientConnectionNewFromC(native unsafe.Pointer) *DtlsClientConnection {
	return &DtlsClientConnection{native: native}
}

// DtlsConnection is a representation of the C interface GDtlsConnection.
//
// since 2.48
type DtlsConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsConnection that represents the DtlsConnection.
func (recv *DtlsConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsConnectionNewFromC creates a new DtlsConnection from a pointer to the C GDtlsConnection that represents the DtlsConnection.
func DtlsConnectionNewFromC(native unsafe.Pointer) *DtlsConnection {
	return &DtlsConnection{native: native}
}

// DtlsServerConnection is a representation of the C interface GDtlsServerConnection.
//
// since 2.48
type DtlsServerConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDtlsServerConnection that represents the DtlsServerConnection.
func (recv *DtlsServerConnection) ToC() unsafe.Pointer {
	return recv.native
}

// DtlsServerConnectionNewFromC creates a new DtlsServerConnection from a pointer to the C GDtlsServerConnection that represents the DtlsServerConnection.
func DtlsServerConnectionNewFromC(native unsafe.Pointer) *DtlsServerConnection {
	return &DtlsServerConnection{native: native}
}

// File is a representation of the C interface GFile.
type File struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFile that represents the File.
func (recv *File) ToC() unsafe.Pointer {
	return recv.native
}

// FileNewFromC creates a new File from a pointer to the C GFile that represents the File.
func FileNewFromC(native unsafe.Pointer) *File {
	return &File{native: native}
}

// FileDescriptorBased is a representation of the C interface GFileDescriptorBased.
//
// since 2.24
type FileDescriptorBased struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileDescriptorBased that represents the FileDescriptorBased.
func (recv *FileDescriptorBased) ToC() unsafe.Pointer {
	return recv.native
}

// FileDescriptorBasedNewFromC creates a new FileDescriptorBased from a pointer to the C GFileDescriptorBased that represents the FileDescriptorBased.
func FileDescriptorBasedNewFromC(native unsafe.Pointer) *FileDescriptorBased {
	return &FileDescriptorBased{native: native}
}

// Icon is a representation of the C interface GIcon.
type Icon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIcon that represents the Icon.
func (recv *Icon) ToC() unsafe.Pointer {
	return recv.native
}

// IconNewFromC creates a new Icon from a pointer to the C GIcon that represents the Icon.
func IconNewFromC(native unsafe.Pointer) *Icon {
	return &Icon{native: native}
}

// Initable is a representation of the C interface GInitable.
//
// since 2.22
type Initable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInitable that represents the Initable.
func (recv *Initable) ToC() unsafe.Pointer {
	return recv.native
}

// InitableNewFromC creates a new Initable from a pointer to the C GInitable that represents the Initable.
func InitableNewFromC(native unsafe.Pointer) *Initable {
	return &Initable{native: native}
}

// ListModel is a representation of the C interface GListModel.
type ListModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListModel that represents the ListModel.
func (recv *ListModel) ToC() unsafe.Pointer {
	return recv.native
}

// ListModelNewFromC creates a new ListModel from a pointer to the C GListModel that represents the ListModel.
func ListModelNewFromC(native unsafe.Pointer) *ListModel {
	return &ListModel{native: native}
}

// LoadableIcon is a representation of the C interface GLoadableIcon.
type LoadableIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIcon that represents the LoadableIcon.
func (recv *LoadableIcon) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIconNewFromC creates a new LoadableIcon from a pointer to the C GLoadableIcon that represents the LoadableIcon.
func LoadableIconNewFromC(native unsafe.Pointer) *LoadableIcon {
	return &LoadableIcon{native: native}
}

// Mount is a representation of the C interface GMount.
type Mount struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMount that represents the Mount.
func (recv *Mount) ToC() unsafe.Pointer {
	return recv.native
}

// MountNewFromC creates a new Mount from a pointer to the C GMount that represents the Mount.
func MountNewFromC(native unsafe.Pointer) *Mount {
	return &Mount{native: native}
}

// NetworkMonitor is a representation of the C interface GNetworkMonitor.
//
// since 2.32
type NetworkMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkMonitor that represents the NetworkMonitor.
func (recv *NetworkMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkMonitorNewFromC creates a new NetworkMonitor from a pointer to the C GNetworkMonitor that represents the NetworkMonitor.
func NetworkMonitorNewFromC(native unsafe.Pointer) *NetworkMonitor {
	return &NetworkMonitor{native: native}
}

// PollableInputStream is a representation of the C interface GPollableInputStream.
//
// since 2.28
type PollableInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableInputStream that represents the PollableInputStream.
func (recv *PollableInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// PollableInputStreamNewFromC creates a new PollableInputStream from a pointer to the C GPollableInputStream that represents the PollableInputStream.
func PollableInputStreamNewFromC(native unsafe.Pointer) *PollableInputStream {
	return &PollableInputStream{native: native}
}

// PollableOutputStream is a representation of the C interface GPollableOutputStream.
//
// since 2.28
type PollableOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPollableOutputStream that represents the PollableOutputStream.
func (recv *PollableOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// PollableOutputStreamNewFromC creates a new PollableOutputStream from a pointer to the C GPollableOutputStream that represents the PollableOutputStream.
func PollableOutputStreamNewFromC(native unsafe.Pointer) *PollableOutputStream {
	return &PollableOutputStream{native: native}
}

// Proxy is a representation of the C interface GProxy.
//
// since 2.26
type Proxy struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxy that represents the Proxy.
func (recv *Proxy) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyNewFromC creates a new Proxy from a pointer to the C GProxy that represents the Proxy.
func ProxyNewFromC(native unsafe.Pointer) *Proxy {
	return &Proxy{native: native}
}

// ProxyResolver is a representation of the C interface GProxyResolver.
//
// since 2.26
type ProxyResolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyResolver that represents the ProxyResolver.
func (recv *ProxyResolver) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyResolverNewFromC creates a new ProxyResolver from a pointer to the C GProxyResolver that represents the ProxyResolver.
func ProxyResolverNewFromC(native unsafe.Pointer) *ProxyResolver {
	return &ProxyResolver{native: native}
}

// RemoteActionGroup is a representation of the C interface GRemoteActionGroup.
//
// since 2.32
type RemoteActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GRemoteActionGroup that represents the RemoteActionGroup.
func (recv *RemoteActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// RemoteActionGroupNewFromC creates a new RemoteActionGroup from a pointer to the C GRemoteActionGroup that represents the RemoteActionGroup.
func RemoteActionGroupNewFromC(native unsafe.Pointer) *RemoteActionGroup {
	return &RemoteActionGroup{native: native}
}

// Seekable is a representation of the C interface GSeekable.
type Seekable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekable that represents the Seekable.
func (recv *Seekable) ToC() unsafe.Pointer {
	return recv.native
}

// SeekableNewFromC creates a new Seekable from a pointer to the C GSeekable that represents the Seekable.
func SeekableNewFromC(native unsafe.Pointer) *Seekable {
	return &Seekable{native: native}
}

// SocketConnectable is a representation of the C interface GSocketConnectable.
type SocketConnectable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectable that represents the SocketConnectable.
func (recv *SocketConnectable) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectableNewFromC creates a new SocketConnectable from a pointer to the C GSocketConnectable that represents the SocketConnectable.
func SocketConnectableNewFromC(native unsafe.Pointer) *SocketConnectable {
	return &SocketConnectable{native: native}
}

// TlsBackend is a representation of the C interface GTlsBackend.
//
// since 2.28
type TlsBackend struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsBackend that represents the TlsBackend.
func (recv *TlsBackend) ToC() unsafe.Pointer {
	return recv.native
}

// TlsBackendNewFromC creates a new TlsBackend from a pointer to the C GTlsBackend that represents the TlsBackend.
func TlsBackendNewFromC(native unsafe.Pointer) *TlsBackend {
	return &TlsBackend{native: native}
}

// TlsClientConnection is a representation of the C interface GTlsClientConnection.
//
// since 2.28
type TlsClientConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsClientConnection that represents the TlsClientConnection.
func (recv *TlsClientConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsClientConnectionNewFromC creates a new TlsClientConnection from a pointer to the C GTlsClientConnection that represents the TlsClientConnection.
func TlsClientConnectionNewFromC(native unsafe.Pointer) *TlsClientConnection {
	return &TlsClientConnection{native: native}
}

// TlsFileDatabase is a representation of the C interface GTlsFileDatabase.
//
// since 2.30
type TlsFileDatabase struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsFileDatabase that represents the TlsFileDatabase.
func (recv *TlsFileDatabase) ToC() unsafe.Pointer {
	return recv.native
}

// TlsFileDatabaseNewFromC creates a new TlsFileDatabase from a pointer to the C GTlsFileDatabase that represents the TlsFileDatabase.
func TlsFileDatabaseNewFromC(native unsafe.Pointer) *TlsFileDatabase {
	return &TlsFileDatabase{native: native}
}

// TlsServerConnection is a representation of the C interface GTlsServerConnection.
//
// since 2.28
type TlsServerConnection struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsServerConnection that represents the TlsServerConnection.
func (recv *TlsServerConnection) ToC() unsafe.Pointer {
	return recv.native
}

// TlsServerConnectionNewFromC creates a new TlsServerConnection from a pointer to the C GTlsServerConnection that represents the TlsServerConnection.
func TlsServerConnectionNewFromC(native unsafe.Pointer) *TlsServerConnection {
	return &TlsServerConnection{native: native}
}

// Volume is a representation of the C interface GVolume.
type Volume struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolume that represents the Volume.
func (recv *Volume) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeNewFromC creates a new Volume from a pointer to the C GVolume that represents the Volume.
func VolumeNewFromC(native unsafe.Pointer) *Volume {
	return &Volume{native: native}
}
