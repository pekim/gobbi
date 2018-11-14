# `gio` Constants

## `DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME`

Extension point for default handler to URI association. See
[Extending GIO][extending-gio].

C - `G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME`

## `FILE_ATTRIBUTE_ACCESS_CAN_DELETE`

A key in the "access" namespace for checking deletion privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to delete the file.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE`

## `FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE`

A key in the "access" namespace for getting execution privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to execute the file.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE`

## `FILE_ATTRIBUTE_ACCESS_CAN_READ`

A key in the "access" namespace for getting read privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to read the file.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_READ`

## `FILE_ATTRIBUTE_ACCESS_CAN_RENAME`

A key in the "access" namespace for checking renaming privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to rename the file.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME`

## `FILE_ATTRIBUTE_ACCESS_CAN_TRASH`

A key in the "access" namespace for checking trashing privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to move the file to
the trash.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH`

## `FILE_ATTRIBUTE_ACCESS_CAN_WRITE`

A key in the "access" namespace for getting write privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to write to the file.

C - `G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE`

## `FILE_ATTRIBUTE_DOS_IS_ARCHIVE`

A key in the "dos" namespace for checking if the file's archive flag
is set. This attribute is %TRUE if the archive flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE`

## `FILE_ATTRIBUTE_DOS_IS_SYSTEM`

A key in the "dos" namespace for checking if the file's backup flag
is set. This attribute is %TRUE if the backup flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_DOS_IS_SYSTEM`

## `FILE_ATTRIBUTE_ETAG_VALUE`

A key in the "etag" namespace for getting the value of the file's
entity tag. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_ETAG_VALUE`

## `FILE_ATTRIBUTE_FILESYSTEM_FREE`

A key in the "filesystem" namespace for getting the number of bytes of free space left on the
file system. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_FREE`

## `FILE_ATTRIBUTE_FILESYSTEM_READONLY`

A key in the "filesystem" namespace for checking if the file system
is read only. Is set to %TRUE if the file system is read only.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_READONLY`

## `FILE_ATTRIBUTE_FILESYSTEM_SIZE`

A key in the "filesystem" namespace for getting the total size (in bytes) of the file system,
used in g_file_query_filesystem_info(). Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_SIZE`

## `FILE_ATTRIBUTE_FILESYSTEM_TYPE`

A key in the "filesystem" namespace for getting the file system's type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_TYPE`

## `FILE_ATTRIBUTE_FILESYSTEM_USED`

A key in the "filesystem" namespace for getting the number of bytes of used on the
file system. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_USED`

## `FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW`

A key in the "filesystem" namespace for hinting a file manager
application whether it should preview (e.g. thumbnail) files on the
file system. The value for this key contain a
&num;GFilesystemPreviewType.

C - `G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW`

## `FILE_ATTRIBUTE_GVFS_BACKEND`

A key in the "gvfs" namespace that gets the name of the current
GVFS backend in use. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_GVFS_BACKEND`

## `FILE_ATTRIBUTE_ID_FILE`

A key in the "id" namespace for getting a file identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during listing files, to avoid recursive
directory scanning.

C - `G_FILE_ATTRIBUTE_ID_FILE`

## `FILE_ATTRIBUTE_ID_FILESYSTEM`

A key in the "id" namespace for getting the file system identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during drag and drop to see if the source
and target are on the same filesystem (default to move) or not (default
to copy).

C - `G_FILE_ATTRIBUTE_ID_FILESYSTEM`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be ejected.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) is mountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be polled.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_START`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started
degraded.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be stopped.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP`

## `FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)  is unmountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT`

## `FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI`

A key in the "mountable" namespace for getting the HAL UDI for the mountable
file. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI`

## `FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC`

A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)
is automatically polled for media.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC`

## `FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE`

A key in the "mountable" namespace for getting the #GDriveStartStopType.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE`

## `FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE`

A key in the "mountable" namespace for getting the unix device.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE`

## `FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE`

A key in the "mountable" namespace for getting the unix device file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE`

## `FILE_ATTRIBUTE_OWNER_GROUP`

A key in the "owner" namespace for getting the file owner's group.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_OWNER_GROUP`

## `FILE_ATTRIBUTE_OWNER_USER`

A key in the "owner" namespace for getting the user name of the
file's owner. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_OWNER_USER`

## `FILE_ATTRIBUTE_OWNER_USER_REAL`

A key in the "owner" namespace for getting the real name of the
user that owns the file. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_OWNER_USER_REAL`

## `FILE_ATTRIBUTE_PREVIEW_ICON`

A key in the "preview" namespace for getting a #GIcon that can be
used to get preview of the file. For example, it may be a low
resolution thumbnail without metadata. Corresponding
&num;GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.  The value
for this key should contain a #GIcon.

C - `G_FILE_ATTRIBUTE_PREVIEW_ICON`

## `FILE_ATTRIBUTE_RECENT_MODIFIED`

A key in the "recent" namespace for getting time, when the metadata for the
file in `recent:///` was last changed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_INT64.

C - `G_FILE_ATTRIBUTE_RECENT_MODIFIED`

## `FILE_ATTRIBUTE_SELINUX_CONTEXT`

A key in the "selinux" namespace for getting the file's SELinux
context. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING. Note that this attribute is only
available if GLib has been built with SELinux support.

C - `G_FILE_ATTRIBUTE_SELINUX_CONTEXT`

## `FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE`

A key in the "standard" namespace for getting the amount of disk space
that is consumed by the file (in bytes).  This will generally be larger
than the file size (due to block size overhead) but can occasionally be
smaller (for example, for sparse files).
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE`

## `FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE`

A key in the "standard" namespace for getting the content type of the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
The value for this key should contain a valid content type.

C - `G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE`

## `FILE_ATTRIBUTE_STANDARD_COPY_NAME`

A key in the "standard" namespace for getting the copy name of the file.
The copy name is an optional version of the name. If available it's always
in UTF8, and corresponds directly to the original filename (only transcoded to
UTF8). This is useful if you want to copy the file to another filesystem that
might have a different encoding. If the filename is not a valid string in the
encoding selected for the filesystem it is in then the copy name will not be set.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_COPY_NAME`

## `FILE_ATTRIBUTE_STANDARD_DESCRIPTION`

A key in the "standard" namespace for getting the description of the file.
The description is a utf8 string that describes the file, generally containing
the filename, but can also contain furter information. Example descriptions
could be "filename (on hostname)" for a remote file or "filename (in trash)"
for a file in the trash. This is useful for instance as the window title
when displaying a directory or for a bookmarks menu.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION`

## `FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME`

A key in the "standard" namespace for getting the display name of the file.
A display name is guaranteed to be in UTF8 and can thus be displayed in
the UI.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME`

## `FILE_ATTRIBUTE_STANDARD_EDIT_NAME`

A key in the "standard" namespace for edit name of the file.
An edit name is similar to the display name, but it is meant to be
used when you want to rename the file in the UI. The display name
might contain information you don't want in the new filename (such as
"(invalid unicode)" if the filename was in an invalid encoding).

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME`

## `FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE`

A key in the "standard" namespace for getting the fast content type.
The fast content type isn't as reliable as the regular one, as it
only uses the filename to guess it, but it is faster to calculate than the
regular content type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE`

## `FILE_ATTRIBUTE_STANDARD_ICON`

A key in the "standard" namespace for getting the icon for the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.
The value for this key should contain a #GIcon.

C - `G_FILE_ATTRIBUTE_STANDARD_ICON`

## `FILE_ATTRIBUTE_STANDARD_IS_BACKUP`

A key in the "standard" namespace for checking if a file is a backup file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP`

## `FILE_ATTRIBUTE_STANDARD_IS_HIDDEN`

A key in the "standard" namespace for checking if a file is hidden.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN`

## `FILE_ATTRIBUTE_STANDARD_IS_SYMLINK`

A key in the "standard" namespace for checking if the file is a symlink.
Typically the actual type is something else, if we followed the symlink
to get the type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK`

## `FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL`

A key in the "standard" namespace for checking if a file is virtual.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL`

## `FILE_ATTRIBUTE_STANDARD_IS_VOLATILE`

A key in the "standard" namespace for checking if a file is
volatile. This is meant for opaque, non-POSIX-like backends to
indicate that the URI is not persistent. Applications should look
at #G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET for the persistent URI.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_STANDARD_IS_VOLATILE`

## `FILE_ATTRIBUTE_STANDARD_NAME`

A key in the "standard" namespace for getting the name of the file.
The name is the on-disk filename which may not be in any known encoding,
and can thus not be generally displayed as is.
Use #G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME if you need to display the
name in a user interface.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_NAME`

## `FILE_ATTRIBUTE_STANDARD_SIZE`

A key in the "standard" namespace for getting the file's size (in bytes).
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_STANDARD_SIZE`

## `FILE_ATTRIBUTE_STANDARD_SORT_ORDER`

A key in the "standard" namespace for setting the sort order of a file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_INT32.
An example use would be in file managers, which would use this key
to set the order files are displayed. Files with smaller sort order
should be sorted first, and files without sort order as if sort order
was zero.

C - `G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER`

## `FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON`

A key in the "standard" namespace for getting the symbolic icon for the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.
The value for this key should contain a #GIcon.

C - `G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON`

## `FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET`

A key in the "standard" namespace for getting the symlink target, if the file
is a symlink. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET`

## `FILE_ATTRIBUTE_STANDARD_TARGET_URI`

A key in the "standard" namespace for getting the target URI for the file, in
the case of %G_FILE_TYPE_SHORTCUT or %G_FILE_TYPE_MOUNTABLE files.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_STANDARD_TARGET_URI`

## `FILE_ATTRIBUTE_STANDARD_TYPE`

A key in the "standard" namespace for storing file types.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
The value for this key should contain a #GFileType.

C - `G_FILE_ATTRIBUTE_STANDARD_TYPE`

## `FILE_ATTRIBUTE_THUMBNAILING_FAILED`

A key in the "thumbnail" namespace for checking if thumbnailing failed.
This attribute is %TRUE if thumbnailing failed. Corresponding
&num;GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_THUMBNAILING_FAILED`

## `FILE_ATTRIBUTE_THUMBNAIL_IS_VALID`

A key in the "thumbnail" namespace for checking whether the thumbnail is outdated.
This attribute is %TRUE if the thumbnail is up-to-date with the file it represents,
and %FALSE if the file has been modified since the thumbnail was generated.

If %G_FILE_ATTRIBUTE_THUMBNAILING_FAILED is %TRUE and this attribute is %FALSE,
it indicates that thumbnailing may be attempted again and may succeed.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID`

## `FILE_ATTRIBUTE_THUMBNAIL_PATH`

A key in the "thumbnail" namespace for getting the path to the thumbnail
image. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.

C - `G_FILE_ATTRIBUTE_THUMBNAIL_PATH`

## `FILE_ATTRIBUTE_TIME_ACCESS`

A key in the "time" namespace for getting the time the file was last
accessed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was last accessed, in seconds since the UNIX epoch.

C - `G_FILE_ATTRIBUTE_TIME_ACCESS`

## `FILE_ATTRIBUTE_TIME_ACCESS_USEC`

A key in the "time" namespace for getting the microseconds of the time
the file was last accessed. This should be used in conjunction with
&num;G_FILE_ATTRIBUTE_TIME_ACCESS. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_TIME_ACCESS_USEC`

## `FILE_ATTRIBUTE_TIME_CHANGED`

A key in the "time" namespace for getting the time the file was last
changed. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was last changed, in seconds since the
UNIX epoch.

This corresponds to the traditional UNIX ctime.

C - `G_FILE_ATTRIBUTE_TIME_CHANGED`

## `FILE_ATTRIBUTE_TIME_CHANGED_USEC`

A key in the "time" namespace for getting the microseconds of the time
the file was last changed. This should be used in conjunction with
&num;G_FILE_ATTRIBUTE_TIME_CHANGED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_TIME_CHANGED_USEC`

## `FILE_ATTRIBUTE_TIME_CREATED`

A key in the "time" namespace for getting the time the file was created.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was created, in seconds since the UNIX
epoch.

This corresponds to the NTFS ctime.

C - `G_FILE_ATTRIBUTE_TIME_CREATED`

## `FILE_ATTRIBUTE_TIME_CREATED_USEC`

A key in the "time" namespace for getting the microseconds of the time
the file was created. This should be used in conjunction with
&num;G_FILE_ATTRIBUTE_TIME_CREATED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_TIME_CREATED_USEC`

## `FILE_ATTRIBUTE_TIME_MODIFIED`

A key in the "time" namespace for getting the time the file was last
modified. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was modified, in seconds since the UNIX epoch.

C - `G_FILE_ATTRIBUTE_TIME_MODIFIED`

## `FILE_ATTRIBUTE_TIME_MODIFIED_USEC`

A key in the "time" namespace for getting the microseconds of the time
the file was last modified. This should be used in conjunction with
&num;G_FILE_ATTRIBUTE_TIME_MODIFIED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC`

## `FILE_ATTRIBUTE_TRASH_DELETION_DATE`

A key in the "trash" namespace.  When requested against
items in `trash:///`, will return the date and time when the file
was trashed. The format of the returned string is YYYY-MM-DDThh:mm:ss.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.

C - `G_FILE_ATTRIBUTE_TRASH_DELETION_DATE`

## `FILE_ATTRIBUTE_TRASH_ITEM_COUNT`

A key in the "trash" namespace.  When requested against
`trash:///` returns the number of (toplevel) items in the trash folder.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT`

## `FILE_ATTRIBUTE_TRASH_ORIG_PATH`

A key in the "trash" namespace.  When requested against
items in `trash:///`, will return the original path to the file before it
was trashed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.

C - `G_FILE_ATTRIBUTE_TRASH_ORIG_PATH`

## `FILE_ATTRIBUTE_UNIX_BLOCKS`

A key in the "unix" namespace for getting the number of blocks allocated
for the file. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_UNIX_BLOCKS`

## `FILE_ATTRIBUTE_UNIX_BLOCK_SIZE`

A key in the "unix" namespace for getting the block size for the file
system. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE`

## `FILE_ATTRIBUTE_UNIX_DEVICE`

A key in the "unix" namespace for getting the device id of the device the
file is located on (see stat() documentation). This attribute is only
available for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_DEVICE`

## `FILE_ATTRIBUTE_UNIX_GID`

A key in the "unix" namespace for getting the group ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_GID`

## `FILE_ATTRIBUTE_UNIX_INODE`

A key in the "unix" namespace for getting the inode of the file.
This attribute is only available for UNIX file systems. Corresponding
&num;GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.

C - `G_FILE_ATTRIBUTE_UNIX_INODE`

## `FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT`

A key in the "unix" namespace for checking if the file represents a
UNIX mount point. This attribute is %TRUE if the file is a UNIX mount
point. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.

C - `G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT`

## `FILE_ATTRIBUTE_UNIX_MODE`

A key in the "unix" namespace for getting the mode of the file
(e.g. whether the file is a regular file, symlink, etc). See lstat()
documentation. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_MODE`

## `FILE_ATTRIBUTE_UNIX_NLINK`

A key in the "unix" namespace for getting the number of hard links
for a file. See lstat() documentation. This attribute is only available
for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_NLINK`

## `FILE_ATTRIBUTE_UNIX_RDEV`

A key in the "unix" namespace for getting the device ID for the file
(if it is a special file). See lstat() documentation. This attribute
is only available for UNIX file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_RDEV`

## `FILE_ATTRIBUTE_UNIX_UID`

A key in the "unix" namespace for getting the user ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.

C - `G_FILE_ATTRIBUTE_UNIX_UID`

## `MENU_ATTRIBUTE_ACTION`

The menu item attribute which holds the action name of the item.  Action
names are namespaced with an identifier for the action group in which the
action resides. For example, "win." for window-specific actions and "app."
for application-wide actions.

See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().

C - `G_MENU_ATTRIBUTE_ACTION`

## `MENU_ATTRIBUTE_ACTION_NAMESPACE`

The menu item attribute that holds the namespace for all action names in
menus that are linked from this item.

C - `G_MENU_ATTRIBUTE_ACTION_NAMESPACE`

## `MENU_ATTRIBUTE_ICON`

The menu item attribute which holds the icon of the item.

The icon is stored in the format returned by g_icon_serialize().

This attribute is intended only to represent 'noun' icons such as
favicons for a webpage, or application icons.  It should not be used
for 'verbs' (ie: stock icons).

C - `G_MENU_ATTRIBUTE_ICON`

## `MENU_ATTRIBUTE_LABEL`

The menu item attribute which holds the label of the item.

C - `G_MENU_ATTRIBUTE_LABEL`

## `MENU_ATTRIBUTE_TARGET`

The menu item attribute which holds the target with which the item's action
will be activated.

See also g_menu_item_set_action_and_target()

C - `G_MENU_ATTRIBUTE_TARGET`

## `MENU_LINK_SECTION`

The name of the link that associates a menu item with a section.  The linked
menu will usually be shown in place of the menu item, using the item's label
as a header.

See also g_menu_item_set_link().

C - `G_MENU_LINK_SECTION`

## `MENU_LINK_SUBMENU`

The name of the link that associates a menu item with a submenu.

See also g_menu_item_set_link().

C - `G_MENU_LINK_SUBMENU`

## `NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME`



C - `G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME`

## `NETWORK_MONITOR_EXTENSION_POINT_NAME`

Extension point for network status monitoring functionality.
See [Extending GIO][extending-gio].

C - `G_NETWORK_MONITOR_EXTENSION_POINT_NAME`

## `PROXY_EXTENSION_POINT_NAME`

Extension point for proxy functionality.
See [Extending GIO][extending-gio].

C - `G_PROXY_EXTENSION_POINT_NAME`

## `PROXY_RESOLVER_EXTENSION_POINT_NAME`

Extension point for proxy resolving functionality.
See [Extending GIO][extending-gio].

C - `G_PROXY_RESOLVER_EXTENSION_POINT_NAME`

## `TLS_BACKEND_EXTENSION_POINT_NAME`

Extension point for TLS functionality via #GTlsBackend.
See [Extending GIO][extending-gio].

C - `G_TLS_BACKEND_EXTENSION_POINT_NAME`

## `TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT`

The purpose used to verify the client certificate in a TLS connection.
Used by TLS servers.

C - `G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT`

## `TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER`

The purpose used to verify the server certificate in a TLS connection. This
is the most common purpose in use. Used by TLS clients.

C - `G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER`

## `VFS_EXTENSION_POINT_NAME`

Extension point for #GVfs functionality.
See [Extending GIO][extending-gio].

C - `G_VFS_EXTENSION_POINT_NAME`

## `VOLUME_IDENTIFIER_KIND_CLASS`

The string used to obtain the volume class with g_volume_get_identifier().

Known volume classes include `device` and `network`. Other classes may
be added in the future.

This is intended to be used by applications to classify #GVolume
instances into different sections - for example a file manager or
file chooser can use this information to show `network` volumes under
a "Network" heading and `device` volumes under a "Devices" heading.

C - `G_VOLUME_IDENTIFIER_KIND_CLASS`

## `VOLUME_IDENTIFIER_KIND_HAL_UDI`

The string used to obtain a Hal UDI with g_volume_get_identifier().

C - `G_VOLUME_IDENTIFIER_KIND_HAL_UDI`

## `VOLUME_IDENTIFIER_KIND_LABEL`

The string used to obtain a filesystem label with g_volume_get_identifier().

C - `G_VOLUME_IDENTIFIER_KIND_LABEL`

## `VOLUME_IDENTIFIER_KIND_NFS_MOUNT`

The string used to obtain a NFS mount with g_volume_get_identifier().

C - `G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT`

## `VOLUME_IDENTIFIER_KIND_UNIX_DEVICE`

The string used to obtain a Unix device path with g_volume_get_identifier().

C - `G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE`

## `VOLUME_IDENTIFIER_KIND_UUID`

The string used to obtain a UUID with g_volume_get_identifier().

C - `G_VOLUME_IDENTIFIER_KIND_UUID`

## `VOLUME_MONITOR_EXTENSION_POINT_NAME`

Extension point for volume monitor functionality.
See [Extending GIO][extending-gio].

C - `G_VOLUME_MONITOR_EXTENSION_POINT_NAME`

