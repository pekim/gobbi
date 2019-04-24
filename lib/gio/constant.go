// This is a generated file - DO NOT EDIT

package gio

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
