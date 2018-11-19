+++
title = "constants"
+++
<p class="api-heading">DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for default handler to URI association. See
[Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_DELETE</p>
<p class="api-doc">A key in the "access" namespace for checking deletion privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to delete the file.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE</p>
<p class="api-doc">A key in the "access" namespace for getting execution privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to execute the file.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_READ</p>
<p class="api-doc">A key in the "access" namespace for getting read privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to read the file.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_READ</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_RENAME</p>
<p class="api-doc">A key in the "access" namespace for checking renaming privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to rename the file.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_TRASH</p>
<p class="api-doc">A key in the "access" namespace for checking trashing privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to move the file to
the trash.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ACCESS_CAN_WRITE</p>
<p class="api-doc">A key in the "access" namespace for getting write privileges.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
This attribute will be %TRUE if the user is able to write to the file.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_DOS_IS_ARCHIVE</p>
<p class="api-doc">A key in the "dos" namespace for checking if the file's archive flag
is set. This attribute is %TRUE if the archive flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_DOS_IS_SYSTEM</p>
<p class="api-doc">A key in the "dos" namespace for checking if the file's backup flag
is set. This attribute is %TRUE if the backup flag is set. This attribute
is only available for DOS file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_DOS_IS_SYSTEM</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ETAG_VALUE</p>
<p class="api-doc">A key in the "etag" namespace for getting the value of the file's
entity tag. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ETAG_VALUE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_FREE</p>
<p class="api-doc">A key in the "filesystem" namespace for getting the number of bytes of free space left on the
file system. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_FREE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_READONLY</p>
<p class="api-doc">A key in the "filesystem" namespace for checking if the file system
is read only. Is set to %TRUE if the file system is read only.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_READONLY</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_SIZE</p>
<p class="api-doc">A key in the "filesystem" namespace for getting the total size (in bytes) of the file system,
used in g_file_query_filesystem_info(). Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_SIZE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_TYPE</p>
<p class="api-doc">A key in the "filesystem" namespace for getting the file system's type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_TYPE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_USED</p>
<p class="api-doc">A key in the "filesystem" namespace for getting the number of bytes of used on the
file system. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_USED</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW</p>
<p class="api-doc">A key in the "filesystem" namespace for hinting a file manager
application whether it should preview (e.g. thumbnail) files on the
file system. The value for this key contain a
#GFilesystemPreviewType.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_GVFS_BACKEND</p>
<p class="api-doc">A key in the "gvfs" namespace that gets the name of the current
GVFS backend in use. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_GVFS_BACKEND</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ID_FILE</p>
<p class="api-doc">A key in the "id" namespace for getting a file identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during listing files, to avoid recursive
directory scanning.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ID_FILE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_ID_FILESYSTEM</p>
<p class="api-doc">A key in the "id" namespace for getting the file system identifier.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
An example use would be during drag and drop to see if the source
and target are on the same filesystem (default to move) or not (default
to copy).</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_ID_FILESYSTEM</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be ejected.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) is mountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be polled.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_START</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be started
degraded.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE) can be stopped.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)  is unmountable.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI</p>
<p class="api-doc">A key in the "mountable" namespace for getting the HAL UDI for the mountable
file. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC</p>
<p class="api-doc">A key in the "mountable" namespace for checking if a file (of type G_FILE_TYPE_MOUNTABLE)
is automatically polled for media.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE</p>
<p class="api-doc">A key in the "mountable" namespace for getting the #GDriveStartStopType.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE</p>
<p class="api-doc">A key in the "mountable" namespace for getting the unix device.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE</p>
<p class="api-doc">A key in the "mountable" namespace for getting the unix device file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE</p>
  <p class="api-since">since 2.22</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_OWNER_GROUP</p>
<p class="api-doc">A key in the "owner" namespace for getting the file owner's group.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_OWNER_GROUP</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_OWNER_USER</p>
<p class="api-doc">A key in the "owner" namespace for getting the user name of the
file's owner. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_OWNER_USER</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_OWNER_USER_REAL</p>
<p class="api-doc">A key in the "owner" namespace for getting the real name of the
user that owns the file. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_OWNER_USER_REAL</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_PREVIEW_ICON</p>
<p class="api-doc">A key in the "preview" namespace for getting a #GIcon that can be
used to get preview of the file. For example, it may be a low
resolution thumbnail without metadata. Corresponding
#GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.  The value
for this key should contain a #GIcon.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_PREVIEW_ICON</p>
  <p class="api-since">since 2.20</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_RECENT_MODIFIED</p>
<p class="api-doc">A key in the "recent" namespace for getting time, when the metadata for the
file in `recent:///` was last changed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_INT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_RECENT_MODIFIED</p>
  <p class="api-since">since 2.52</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_SELINUX_CONTEXT</p>
<p class="api-doc">A key in the "selinux" namespace for getting the file's SELinux
context. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_STRING. Note that this attribute is only
available if GLib has been built with SELinux support.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_SELINUX_CONTEXT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE</p>
<p class="api-doc">A key in the "standard" namespace for getting the amount of disk space
that is consumed by the file (in bytes).  This will generally be larger
than the file size (due to block size overhead) but can occasionally be
smaller (for example, for sparse files).
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_ALLOCATED_SIZE</p>
  <p class="api-since">since 2.20</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE</p>
<p class="api-doc">A key in the "standard" namespace for getting the content type of the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.
The value for this key should contain a valid content type.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_COPY_NAME</p>
<p class="api-doc">A key in the "standard" namespace for getting the copy name of the file.
The copy name is an optional version of the name. If available it's always
in UTF8, and corresponds directly to the original filename (only transcoded to
UTF8). This is useful if you want to copy the file to another filesystem that
might have a different encoding. If the filename is not a valid string in the
encoding selected for the filesystem it is in then the copy name will not be set.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_COPY_NAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_DESCRIPTION</p>
<p class="api-doc">A key in the "standard" namespace for getting the description of the file.
The description is a utf8 string that describes the file, generally containing
the filename, but can also contain furter information. Example descriptions
could be "filename (on hostname)" for a remote file or "filename (in trash)"
for a file in the trash. This is useful for instance as the window title
when displaying a directory or for a bookmarks menu.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME</p>
<p class="api-doc">A key in the "standard" namespace for getting the display name of the file.
A display name is guaranteed to be in UTF8 and can thus be displayed in
the UI.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_EDIT_NAME</p>
<p class="api-doc">A key in the "standard" namespace for edit name of the file.
An edit name is similar to the display name, but it is meant to be
used when you want to rename the file in the UI. The display name
might contain information you don't want in the new filename (such as
"(invalid unicode)" if the filename was in an invalid encoding).

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE</p>
<p class="api-doc">A key in the "standard" namespace for getting the fast content type.
The fast content type isn't as reliable as the regular one, as it
only uses the filename to guess it, but it is faster to calculate than the
regular content type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_ICON</p>
<p class="api-doc">A key in the "standard" namespace for getting the icon for the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.
The value for this key should contain a #GIcon.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_ICON</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_IS_BACKUP</p>
<p class="api-doc">A key in the "standard" namespace for checking if a file is a backup file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_IS_HIDDEN</p>
<p class="api-doc">A key in the "standard" namespace for checking if a file is hidden.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_IS_SYMLINK</p>
<p class="api-doc">A key in the "standard" namespace for checking if the file is a symlink.
Typically the actual type is something else, if we followed the symlink
to get the type.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL</p>
<p class="api-doc">A key in the "standard" namespace for checking if a file is virtual.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_IS_VOLATILE</p>
<p class="api-doc">A key in the "standard" namespace for checking if a file is
volatile. This is meant for opaque, non-POSIX-like backends to
indicate that the URI is not persistent. Applications should look
at #G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET for the persistent URI.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_IS_VOLATILE</p>
  <p class="api-since">since 2.46</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_NAME</p>
<p class="api-doc">A key in the "standard" namespace for getting the name of the file.
The name is the on-disk filename which may not be in any known encoding,
and can thus not be generally displayed as is.
Use #G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME if you need to display the
name in a user interface.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_NAME</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_SIZE</p>
<p class="api-doc">A key in the "standard" namespace for getting the file's size (in bytes).
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_SIZE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_SORT_ORDER</p>
<p class="api-doc">A key in the "standard" namespace for setting the sort order of a file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_INT32.
An example use would be in file managers, which would use this key
to set the order files are displayed. Files with smaller sort order
should be sorted first, and files without sort order as if sort order
was zero.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON</p>
<p class="api-doc">A key in the "standard" namespace for getting the symbolic icon for the file.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_OBJECT.
The value for this key should contain a #GIcon.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON</p>
  <p class="api-since">since 2.34</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET</p>
<p class="api-doc">A key in the "standard" namespace for getting the symlink target, if the file
is a symlink. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_TARGET_URI</p>
<p class="api-doc">A key in the "standard" namespace for getting the target URI for the file, in
the case of %G_FILE_TYPE_SHORTCUT or %G_FILE_TYPE_MOUNTABLE files.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_TARGET_URI</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_STANDARD_TYPE</p>
<p class="api-doc">A key in the "standard" namespace for storing file types.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.
The value for this key should contain a #GFileType.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_STANDARD_TYPE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_THUMBNAILING_FAILED</p>
<p class="api-doc">A key in the "thumbnail" namespace for checking if thumbnailing failed.
This attribute is %TRUE if thumbnailing failed. Corresponding
#GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_THUMBNAILING_FAILED</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_THUMBNAIL_IS_VALID</p>
<p class="api-doc">A key in the "thumbnail" namespace for checking whether the thumbnail is outdated.
This attribute is %TRUE if the thumbnail is up-to-date with the file it represents,
and %FALSE if the file has been modified since the thumbnail was generated.

If %G_FILE_ATTRIBUTE_THUMBNAILING_FAILED is %TRUE and this attribute is %FALSE,
it indicates that thumbnailing may be attempted again and may succeed.

Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID</p>
  <p class="api-since">since 2.40</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_THUMBNAIL_PATH</p>
<p class="api-doc">A key in the "thumbnail" namespace for getting the path to the thumbnail
image. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_THUMBNAIL_PATH</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_ACCESS</p>
<p class="api-doc">A key in the "time" namespace for getting the time the file was last
accessed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was last accessed, in seconds since the UNIX epoch.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_ACCESS</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_ACCESS_USEC</p>
<p class="api-doc">A key in the "time" namespace for getting the microseconds of the time
the file was last accessed. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_ACCESS. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_ACCESS_USEC</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_CHANGED</p>
<p class="api-doc">A key in the "time" namespace for getting the time the file was last
changed. Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was last changed, in seconds since the
UNIX epoch.

This corresponds to the traditional UNIX ctime.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_CHANGED</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_CHANGED_USEC</p>
<p class="api-doc">A key in the "time" namespace for getting the microseconds of the time
the file was last changed. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_CHANGED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_CHANGED_USEC</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_CREATED</p>
<p class="api-doc">A key in the "time" namespace for getting the time the file was created.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64,
and contains the time since the file was created, in seconds since the UNIX
epoch.

This corresponds to the NTFS ctime.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_CREATED</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_CREATED_USEC</p>
<p class="api-doc">A key in the "time" namespace for getting the microseconds of the time
the file was created. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_CREATED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_CREATED_USEC</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_MODIFIED</p>
<p class="api-doc">A key in the "time" namespace for getting the time the file was last
modified. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT64, and contains the time since the
file was modified, in seconds since the UNIX epoch.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_MODIFIED</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TIME_MODIFIED_USEC</p>
<p class="api-doc">A key in the "time" namespace for getting the microseconds of the time
the file was last modified. This should be used in conjunction with
#G_FILE_ATTRIBUTE_TIME_MODIFIED. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TRASH_DELETION_DATE</p>
<p class="api-doc">A key in the "trash" namespace.  When requested against
items in `trash:///`, will return the date and time when the file
was trashed. The format of the returned string is YYYY-MM-DDThh:mm:ss.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TRASH_DELETION_DATE</p>
  <p class="api-since">since 2.24</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TRASH_ITEM_COUNT</p>
<p class="api-doc">A key in the "trash" namespace.  When requested against
`trash:///` returns the number of (toplevel) items in the trash folder.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_TRASH_ORIG_PATH</p>
<p class="api-doc">A key in the "trash" namespace.  When requested against
items in `trash:///`, will return the original path to the file before it
was trashed. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_TRASH_ORIG_PATH</p>
  <p class="api-since">since 2.24</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_BLOCKS</p>
<p class="api-doc">A key in the "unix" namespace for getting the number of blocks allocated
for the file. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_BLOCKS</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_BLOCK_SIZE</p>
<p class="api-doc">A key in the "unix" namespace for getting the block size for the file
system. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_DEVICE</p>
<p class="api-doc">A key in the "unix" namespace for getting the device id of the device the
file is located on (see stat() documentation). This attribute is only
available for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_DEVICE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_GID</p>
<p class="api-doc">A key in the "unix" namespace for getting the group ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_GID</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_INODE</p>
<p class="api-doc">A key in the "unix" namespace for getting the inode of the file.
This attribute is only available for UNIX file systems. Corresponding
#GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT64.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_INODE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT</p>
<p class="api-doc">A key in the "unix" namespace for checking if the file represents a
UNIX mount point. This attribute is %TRUE if the file is a UNIX mount
point. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_BOOLEAN.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_MODE</p>
<p class="api-doc">A key in the "unix" namespace for getting the mode of the file
(e.g. whether the file is a regular file, symlink, etc). See lstat()
documentation. This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_MODE</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_NLINK</p>
<p class="api-doc">A key in the "unix" namespace for getting the number of hard links
for a file. See lstat() documentation. This attribute is only available
for UNIX file systems. Corresponding #GFileAttributeType is
%G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_NLINK</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_RDEV</p>
<p class="api-doc">A key in the "unix" namespace for getting the device ID for the file
(if it is a special file). See lstat() documentation. This attribute
is only available for UNIX file systems. Corresponding #GFileAttributeType
is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_RDEV</p>
</div>
<p class="api-heading">FILE_ATTRIBUTE_UNIX_UID</p>
<p class="api-doc">A key in the "unix" namespace for getting the user ID for the file.
This attribute is only available for UNIX file systems.
Corresponding #GFileAttributeType is %G_FILE_ATTRIBUTE_TYPE_UINT32.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_FILE_ATTRIBUTE_UNIX_UID</p>
</div>
<p class="api-heading">MENU_ATTRIBUTE_ACTION</p>
<p class="api-doc">The menu item attribute which holds the action name of the item.  Action
names are namespaced with an identifier for the action group in which the
action resides. For example, "win." for window-specific actions and "app."
for application-wide actions.

See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_ATTRIBUTE_ACTION</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">MENU_ATTRIBUTE_ACTION_NAMESPACE</p>
<p class="api-doc">The menu item attribute that holds the namespace for all action names in
menus that are linked from this item.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_ATTRIBUTE_ACTION_NAMESPACE</p>
  <p class="api-since">since 2.36</p>
</div>
<p class="api-heading">MENU_ATTRIBUTE_ICON</p>
<p class="api-doc">The menu item attribute which holds the icon of the item.

The icon is stored in the format returned by g_icon_serialize().

This attribute is intended only to represent 'noun' icons such as
favicons for a webpage, or application icons.  It should not be used
for 'verbs' (ie: stock icons).</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_ATTRIBUTE_ICON</p>
  <p class="api-since">since 2.38</p>
</div>
<p class="api-heading">MENU_ATTRIBUTE_LABEL</p>
<p class="api-doc">The menu item attribute which holds the label of the item.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_ATTRIBUTE_LABEL</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">MENU_ATTRIBUTE_TARGET</p>
<p class="api-doc">The menu item attribute which holds the target with which the item's action
will be activated.

See also g_menu_item_set_action_and_target()</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_ATTRIBUTE_TARGET</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">MENU_LINK_SECTION</p>
<p class="api-doc">The name of the link that associates a menu item with a section.  The linked
menu will usually be shown in place of the menu item, using the item's label
as a header.

See also g_menu_item_set_link().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_LINK_SECTION</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">MENU_LINK_SUBMENU</p>
<p class="api-doc">The name of the link that associates a menu item with a submenu.

See also g_menu_item_set_link().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_MENU_LINK_SUBMENU</p>
  <p class="api-since">since 2.32</p>
</div>
<p class="api-heading">NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME</p>
</div>
<p class="api-heading">NETWORK_MONITOR_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for network status monitoring functionality.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_NETWORK_MONITOR_EXTENSION_POINT_NAME</p>
  <p class="api-since">since 2.30</p>
</div>
<p class="api-heading">PROXY_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for proxy functionality.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_PROXY_EXTENSION_POINT_NAME</p>
  <p class="api-since">since 2.26</p>
</div>
<p class="api-heading">PROXY_RESOLVER_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for proxy resolving functionality.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_PROXY_RESOLVER_EXTENSION_POINT_NAME</p>
</div>
<p class="api-heading">TLS_BACKEND_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for TLS functionality via #GTlsBackend.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TLS_BACKEND_EXTENSION_POINT_NAME</p>
</div>
<p class="api-heading">TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT</p>
<p class="api-doc">The purpose used to verify the client certificate in a TLS connection.
Used by TLS servers.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT</p>
</div>
<p class="api-heading">TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER</p>
<p class="api-doc">The purpose used to verify the server certificate in a TLS connection. This
is the most common purpose in use. Used by TLS clients.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER</p>
</div>
<p class="api-heading">VFS_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for #GVfs functionality.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VFS_EXTENSION_POINT_NAME</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_CLASS</p>
<p class="api-doc">The string used to obtain the volume class with g_volume_get_identifier().

Known volume classes include `device` and `network`. Other classes may
be added in the future.

This is intended to be used by applications to classify #GVolume
instances into different sections - for example a file manager or
file chooser can use this information to show `network` volumes under
a "Network" heading and `device` volumes under a "Devices" heading.</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_CLASS</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_HAL_UDI</p>
<p class="api-doc">The string used to obtain a Hal UDI with g_volume_get_identifier().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_HAL_UDI</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_LABEL</p>
<p class="api-doc">The string used to obtain a filesystem label with g_volume_get_identifier().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_LABEL</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_NFS_MOUNT</p>
<p class="api-doc">The string used to obtain a NFS mount with g_volume_get_identifier().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_UNIX_DEVICE</p>
<p class="api-doc">The string used to obtain a Unix device path with g_volume_get_identifier().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE</p>
</div>
<p class="api-heading">VOLUME_IDENTIFIER_KIND_UUID</p>
<p class="api-doc">The string used to obtain a UUID with g_volume_get_identifier().</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_IDENTIFIER_KIND_UUID</p>
</div>
<p class="api-heading">VOLUME_MONITOR_EXTENSION_POINT_NAME</p>
<p class="api-doc">Extension point for volume monitor functionality.
See [Extending GIO][extending-gio].</p>
<div class="api-notes">
  <p class="api-ctype">C type: G_VOLUME_MONITOR_EXTENSION_POINT_NAME</p>
</div>
