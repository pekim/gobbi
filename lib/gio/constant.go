// This is a generated file - DO NOT EDIT

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
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

/*
Extension point for default handler to URI association. See
[Extending GIO][extending-gio].
*/
const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME string = C.G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME

/*
A key in the "access" namespace for checking deletion privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to delete the file.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_DELETE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE

/*
A key in the "access" namespace for getting execution privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to execute the file.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE

/*
A key in the "access" namespace for getting read privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to read the file.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_READ string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_READ

/*
A key in the "access" namespace for checking renaming privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to rename the file.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_RENAME string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME

/*
A key in the "access" namespace for checking trashing privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to move the file to
the trash.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_TRASH string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH

/*
A key in the "access" namespace for getting write privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to write to the file.
*/
const FILE_ATTRIBUTE_ACCESS_CAN_WRITE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE

/*
A key in the "dos" namespace for checking if the file's archive flag
is set. This attribute is %TRUE if the archive flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_DOS_IS_ARCHIVE string = C.G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE

/*
A key in the "dos" namespace for checking if the file's backup flag
is set. This attribute is %TRUE if the backup flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_DOS_IS_SYSTEM string = C.G_FILE_ATTRIBUTE_DOS_IS_SYSTEM

/*
A key in the "etag" namespace for getting the value of the file's
entity tag. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_ETAG_VALUE string = C.G_FILE_ATTRIBUTE_ETAG_VALUE

/*
A key in the "filesystem" namespace for getting the number of bytes of free space left on the
file system. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64.
*/
const FILE_ATTRIBUTE_FILESYSTEM_FREE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_FREE

/*
A key in the "filesystem" namespace for checking if the file system
is read only. Is set to %TRUE if the file system is read only.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_FILESYSTEM_READONLY string = C.G_FILE_ATTRIBUTE_FILESYSTEM_READONLY

// Blacklisted : FILE_ATTRIBUTE_FILESYSTEM_REMOTE

/*
A key in the "filesystem" namespace for getting the total size (in bytes) of the file system,
used in g_file_query_filesystem_info(). Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT64.
*/
const FILE_ATTRIBUTE_FILESYSTEM_SIZE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_SIZE

/*
A key in the "filesystem" namespace for getting the file system's type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_FILESYSTEM_TYPE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_TYPE

/*
A key in the "filesystem" namespace for hinting a file manager
application whether it should preview (e.g. thumbnail) files on the
file system. The value for this key contain a
#GFilesystemPreviewType.
*/
const FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW string = C.G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW

/*
A key in the "gvfs" namespace that gets the name of the current
GVFS backend in use. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_GVFS_BACKEND string = C.G_FILE_ATTRIBUTE_GVFS_BACKEND

/*
A key in the "id" namespace for getting a file identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during listing files, to avoid recursive
directory scanning.
*/
const FILE_ATTRIBUTE_ID_FILE string = C.G_FILE_ATTRIBUTE_ID_FILE

/*
A key in the "id" namespace for getting the file system identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during drag and drop to see if the source
and target are on the same filesystem (default to move) or not (default
to copy).
*/
const FILE_ATTRIBUTE_ID_FILESYSTEM string = C.G_FILE_ATTRIBUTE_ID_FILESYSTEM

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be ejected.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) is mountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT

/*
A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)  is unmountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT

/*
A key in the "mountable" namespace for getting the HAL UDI for the mountable
file. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI string = C.G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI

/*
A key in the "mountable" namespace for getting the unix device.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE

/*
A key in the "owner" namespace for getting the file owner's group.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_OWNER_GROUP string = C.G_FILE_ATTRIBUTE_OWNER_GROUP

/*
A key in the "owner" namespace for getting the user name of the
file's owner. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_OWNER_USER string = C.G_FILE_ATTRIBUTE_OWNER_USER

/*
A key in the "owner" namespace for getting the real name of the
user that owns the file. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_OWNER_USER_REAL string = C.G_FILE_ATTRIBUTE_OWNER_USER_REAL

/*
A key in the "selinux" namespace for getting the file's SELinux
context. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING. Note that this attribute is only
available if GLib has been built with SELinux support.
*/
const FILE_ATTRIBUTE_SELINUX_CONTEXT string = C.G_FILE_ATTRIBUTE_SELINUX_CONTEXT

/*
A key in the "standard" namespace for getting the content type of the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
The value for this key should contain a valid content type.
*/
const FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE

/*
A key in the "standard" namespace for getting the copy name of the file.
The copy name is an optional version of the name. If available it's always
in UTF8, and corresponds directly to the original filename (only transcoded to
UTF8). This is useful if you want to copy the file to another filesystem that
might have a different encoding. If the filename is not a valid string in the
encoding selected for the filesystem it is in then the copy name will not be set.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_COPY_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_COPY_NAME

/*
A key in the "standard" namespace for getting the description of the file.
The description is a utf8 string that describes the file, generally containing
the filename, but can also contain furter information. Example descriptions
could be "filename (on hostname)" for a remote file or "filename (in trash)"
for a file in the trash. This is useful for instance as the window title
when displaying a directory or for a bookmarks menu.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_DESCRIPTION string = C.G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION

/*
A key in the "standard" namespace for getting the display name of the file.
A display name is guaranteed to be in UTF8 and can thus be displayed in
the UI.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME

/*
A key in the "standard" namespace for edit name of the file.
An edit name is similar to the display name, but it is meant to be
used when you want to rename the file in the UI. The display name
might contain information you don't want in the new filename (such as
"(invalid unicode)" if the filename was in an invalid encoding).

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_EDIT_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME

/*
A key in the "standard" namespace for getting the fast content type.
The fast content type isn't as reliable as the regular one, as it
only uses the filename to guess it, but it is faster to calculate than the
regular content type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE

/*
A key in the "standard" namespace for getting the icon for the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.
The value for this key should contain a #GIcon.
*/
const FILE_ATTRIBUTE_STANDARD_ICON string = C.G_FILE_ATTRIBUTE_STANDARD_ICON

/*
A key in the "standard" namespace for checking if a file is a backup file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_STANDARD_IS_BACKUP string = C.G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP

/*
A key in the "standard" namespace for checking if a file is hidden.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_STANDARD_IS_HIDDEN string = C.G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN

/*
A key in the "standard" namespace for checking if the file is a symlink.
Typically the actual type is something else, if we followed the symlink
to get the type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_STANDARD_IS_SYMLINK string = C.G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK

/*
A key in the "standard" namespace for checking if a file is virtual.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL string = C.G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL

/*
A key in the "standard" namespace for getting the name of the file.
The name is the on-disk filename which may not be in any known encoding,
and can thus not be generally displayed as is.
Use #G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME if you need to display the
name in a user interface.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_NAME

/*
A key in the "standard" namespace for getting the file's size (in bytes).
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.
*/
const FILE_ATTRIBUTE_STANDARD_SIZE string = C.G_FILE_ATTRIBUTE_STANDARD_SIZE

/*
A key in the "standard" namespace for setting the sort order of a file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_INT32.
An example use would be in file managers, which would use this key
to set the order files are displayed. Files with smaller sort order
should be sorted first, and files without sort order as if sort order
was zero.
*/
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER string = C.G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER

/*
A key in the "standard" namespace for getting the symlink target, if the file
is a symlink. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET string = C.G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET

/*
A key in the "standard" namespace for getting the target URI for the file, in
the case of %G_FILE_TYPE_SHORTCUT or %G_FILE_TYPE_MOUNTABLE files.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
*/
const FILE_ATTRIBUTE_STANDARD_TARGET_URI string = C.G_FILE_ATTRIBUTE_STANDARD_TARGET_URI

/*
A key in the "standard" namespace for storing file types.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
The value for this key should contain a #GFileType.
*/
const FILE_ATTRIBUTE_STANDARD_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_TYPE

/*
A key in the "thumbnail" namespace for checking if thumbnailing failed.
This attribute is %TRUE if thumbnailing failed. Corresponding
#GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_THUMBNAILING_FAILED string = C.G_FILE_ATTRIBUTE_THUMBNAILING_FAILED

/*
A key in the "thumbnail" namespace for getting the path to the thumbnail
image. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
*/
const FILE_ATTRIBUTE_THUMBNAIL_PATH string = C.G_FILE_ATTRIBUTE_THUMBNAIL_PATH

/*
A key in the "time" namespace for getting the time the file was last
accessed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was last accessed, in seconds since the UNIX epoch.
*/
const FILE_ATTRIBUTE_TIME_ACCESS string = C.G_FILE_ATTRIBUTE_TIME_ACCESS

/*
A key in the "time" namespace for getting the microseconds of the time
the file was last accessed. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_ACCESS. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_TIME_ACCESS_USEC string = C.G_FILE_ATTRIBUTE_TIME_ACCESS_USEC

/*
A key in the "time" namespace for getting the time the file was last
changed. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was last changed, in seconds since the
UNIX epoch.

This corresponds to the traditional UNIX ctime.
*/
const FILE_ATTRIBUTE_TIME_CHANGED string = C.G_FILE_ATTRIBUTE_TIME_CHANGED

/*
A key in the "time" namespace for getting the microseconds of the time
the file was last changed. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_CHANGED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_TIME_CHANGED_USEC string = C.G_FILE_ATTRIBUTE_TIME_CHANGED_USEC

/*
A key in the "time" namespace for getting the time the file was created.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was created, in seconds since the UNIX
epoch.

This corresponds to the NTFS ctime.
*/
const FILE_ATTRIBUTE_TIME_CREATED string = C.G_FILE_ATTRIBUTE_TIME_CREATED

/*
A key in the "time" namespace for getting the microseconds of the time
the file was created. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_CREATED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_TIME_CREATED_USEC string = C.G_FILE_ATTRIBUTE_TIME_CREATED_USEC

/*
A key in the "time" namespace for getting the time the file was last
modified. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was modified, in seconds since the UNIX epoch.
*/
const FILE_ATTRIBUTE_TIME_MODIFIED string = C.G_FILE_ATTRIBUTE_TIME_MODIFIED

/*
A key in the "time" namespace for getting the microseconds of the time
the file was last modified. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_MODIFIED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_TIME_MODIFIED_USEC string = C.G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC

/*
A key in the "trash" namespace.  When requested against
`trash:///` returns the number of (toplevel) items in the trash folder.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT string = C.G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT

/*
A key in the "unix" namespace for getting the number of blocks allocated
for the file. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.
*/
const FILE_ATTRIBUTE_UNIX_BLOCKS string = C.G_FILE_ATTRIBUTE_UNIX_BLOCKS

/*
A key in the "unix" namespace for getting the block size for the file
system. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_BLOCK_SIZE string = C.G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE

/*
A key in the "unix" namespace for getting the device id of the device the
file is located on (see stat() documentation). This attribute is only
available for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_DEVICE string = C.G_FILE_ATTRIBUTE_UNIX_DEVICE

/*
A key in the "unix" namespace for getting the group ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_GID string = C.G_FILE_ATTRIBUTE_UNIX_GID

/*
A key in the "unix" namespace for getting the inode of the file.
This attribute is only available for UNIX file systems. Corresponding
#GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.
*/
const FILE_ATTRIBUTE_UNIX_INODE string = C.G_FILE_ATTRIBUTE_UNIX_INODE

/*
A key in the "unix" namespace for checking if the file represents a
UNIX mount point. This attribute is %TRUE if the file is a UNIX mount
point. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
*/
const FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT string = C.G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT

/*
A key in the "unix" namespace for getting the mode of the file
(e.g. whether the file is a regular file, symlink, etc). See lstat()
documentation. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_MODE string = C.G_FILE_ATTRIBUTE_UNIX_MODE

/*
A key in the "unix" namespace for getting the number of hard links
for a file. See lstat() documentation. This attribute is only available
for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_NLINK string = C.G_FILE_ATTRIBUTE_UNIX_NLINK

/*
A key in the "unix" namespace for getting the device ID for the file
(if it is a special file). See lstat() documentation. This attribute
is only available for UNIX file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_RDEV string = C.G_FILE_ATTRIBUTE_UNIX_RDEV

/*
A key in the "unix" namespace for getting the user ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
*/
const FILE_ATTRIBUTE_UNIX_UID string = C.G_FILE_ATTRIBUTE_UNIX_UID
const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME string = C.G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME

/*
Extension point for proxy resolving functionality.
See [Extending GIO][extending-gio].
*/
const PROXY_RESOLVER_EXTENSION_POINT_NAME string = C.G_PROXY_RESOLVER_EXTENSION_POINT_NAME

// Blacklisted : SETTINGS_BACKEND_EXTENSION_POINT_NAME

/*
Extension point for TLS functionality via #GTlsBackend.
See [Extending GIO][extending-gio].
*/
const TLS_BACKEND_EXTENSION_POINT_NAME string = C.G_TLS_BACKEND_EXTENSION_POINT_NAME

/*
The purpose used to verify the client certificate in a TLS connection.
Used by TLS servers.
*/
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT string = C.G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT

/*
The purpose used to verify the server certificate in a TLS connection. This
is the most common purpose in use. Used by TLS clients.
*/
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER string = C.G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER

/*
Extension point for #GVfs functionality.
See [Extending GIO][extending-gio].
*/
const VFS_EXTENSION_POINT_NAME string = C.G_VFS_EXTENSION_POINT_NAME

/*
The string used to obtain the volume class with g_volume_get_identifier().

Known volume classes include `device` and `network`. Other classes may
be added in the future.

This is intended to be used by applications to classify #GVolume
instances into different sections - for example a file manager or
file chooser can use this information to show `network` volumes under
a "Network" heading and `device` volumes under a "Devices" heading.
*/
const VOLUME_IDENTIFIER_KIND_CLASS string = C.G_VOLUME_IDENTIFIER_KIND_CLASS

// The string used to obtain a Hal UDI with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_HAL_UDI string = C.G_VOLUME_IDENTIFIER_KIND_HAL_UDI

// The string used to obtain a filesystem label with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_LABEL string = C.G_VOLUME_IDENTIFIER_KIND_LABEL

// The string used to obtain a NFS mount with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT string = C.G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT

// The string used to obtain a Unix device path with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE string = C.G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE

// The string used to obtain a UUID with g_volume_get_identifier().
const VOLUME_IDENTIFIER_KIND_UUID string = C.G_VOLUME_IDENTIFIER_KIND_UUID

/*
Extension point for volume monitor functionality.
See [Extending GIO][extending-gio].
*/
const VOLUME_MONITOR_EXTENSION_POINT_NAME string = C.G_VOLUME_MONITOR_EXTENSION_POINT_NAME
