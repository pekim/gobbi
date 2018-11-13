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

// Flags used when creating a #GAppInfo.
type AppInfoCreateFlags C.GAppInfoCreateFlags

const (
	// No flags.
	APP_INFO_CREATE_NONE AppInfoCreateFlags = 0

	// Application opens in a terminal window.
	APP_INFO_CREATE_NEEDS_TERMINAL AppInfoCreateFlags = 1

	// Application supports URI arguments.
	APP_INFO_CREATE_SUPPORTS_URIS AppInfoCreateFlags = 2

	// Application supports startup notification. Since 2.26
	APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION AppInfoCreateFlags = 4
)

// #GAskPasswordFlags are used to request specific information from the
// user, or to notify the user of their choices in an authentication
// situation.
type AskPasswordFlags C.GAskPasswordFlags

const (
	// operation requires a password.
	ASK_PASSWORD_NEED_PASSWORD AskPasswordFlags = 1

	// operation requires a username.
	ASK_PASSWORD_NEED_USERNAME AskPasswordFlags = 2

	// operation requires a domain.
	ASK_PASSWORD_NEED_DOMAIN AskPasswordFlags = 4

	// operation supports saving settings.
	ASK_PASSWORD_SAVING_SUPPORTED AskPasswordFlags = 8

	// operation supports anonymous users.
	ASK_PASSWORD_ANONYMOUS_SUPPORTED AskPasswordFlags = 16
)

// Flags specifying the behaviour of an attribute.
type FileAttributeInfoFlags C.GFileAttributeInfoFlags

const (
	// no flags set.
	FILE_ATTRIBUTE_INFO_NONE FileAttributeInfoFlags = 0

	// copy the attribute values when the file is copied.
	FILE_ATTRIBUTE_INFO_COPY_WITH_FILE FileAttributeInfoFlags = 1

	// copy the attribute values when the file is moved.
	FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED FileAttributeInfoFlags = 2
)

// Flags used when copying or moving files.
type FileCopyFlags C.GFileCopyFlags

const (
	// No flags set.
	FILE_COPY_NONE FileCopyFlags = 0

	// Overwrite any existing files
	FILE_COPY_OVERWRITE FileCopyFlags = 1

	// Make a backup of any existing files.
	FILE_COPY_BACKUP FileCopyFlags = 2

	// Don't follow symlinks.
	FILE_COPY_NOFOLLOW_SYMLINKS FileCopyFlags = 4

	// Copy all file metadata instead of just default set used for copy (see #GFileInfo).
	FILE_COPY_ALL_METADATA FileCopyFlags = 8

	// Don't use copy and delete fallback if native move not supported.
	FILE_COPY_NO_FALLBACK_FOR_MOVE FileCopyFlags = 16

	// Leaves target file with default perms, instead of setting the source file perms.
	FILE_COPY_TARGET_DEFAULT_PERMS FileCopyFlags = 32
)

// Flags used when an operation may create a file.
type FileCreateFlags C.GFileCreateFlags

const (
	// No flags set.
	FILE_CREATE_NONE FileCreateFlags = 0

	// Create a file that can only be
	// accessed by the current user.
	FILE_CREATE_PRIVATE FileCreateFlags = 1

	// Replace the destination
	// as if it didn't exist before. Don't try to keep any old
	// permissions, replace instead of following links. This
	// is generally useful if you're doing a "copy over"
	// rather than a "save new version of" replace operation.
	// You can think of it as "unlink destination" before
	// writing to it, although the implementation may not
	// be exactly like that. Since 2.20
	FILE_CREATE_REPLACE_DESTINATION FileCreateFlags = 2
)

// Flags used to set what a #GFileMonitor will watch for.
type FileMonitorFlags C.GFileMonitorFlags

const (
	// No flags set.
	FILE_MONITOR_NONE FileMonitorFlags = 0

	// Watch for mount events.
	FILE_MONITOR_WATCH_MOUNTS FileMonitorFlags = 1

	// Pair DELETED and CREATED events caused
	// by file renames (moves) and send a single G_FILE_MONITOR_EVENT_MOVED
	// event instead (NB: not supported on all backends; the default
	// behaviour -without specifying this flag- is to send single DELETED
	// and CREATED events).  Deprecated since 2.46: use
	// %G_FILE_MONITOR_WATCH_MOVES instead.
	FILE_MONITOR_SEND_MOVED FileMonitorFlags = 2

	// Watch for changes to the file made
	// via another hard link. Since 2.36.
	FILE_MONITOR_WATCH_HARD_LINKS FileMonitorFlags = 4

	// Watch for rename operations on a
	// monitored directory.  This causes %G_FILE_MONITOR_EVENT_RENAMED,
	// %G_FILE_MONITOR_EVENT_MOVED_IN and %G_FILE_MONITOR_EVENT_MOVED_OUT
	// events to be emitted when possible.  Since: 2.46.
	FILE_MONITOR_WATCH_MOVES FileMonitorFlags = 8
)

// Flags used when querying a #GFileInfo.
type FileQueryInfoFlags C.GFileQueryInfoFlags

const (
	// No flags set.
	FILE_QUERY_INFO_NONE FileQueryInfoFlags = 0

	// Don't follow symlinks.
	FILE_QUERY_INFO_NOFOLLOW_SYMLINKS FileQueryInfoFlags = 1
)

// Flags used when mounting a mount.
type MountMountFlags C.GMountMountFlags

const (
	// No flags set.
	MOUNT_MOUNT_NONE MountMountFlags = 0
)

// Flags used when an unmounting a mount.
type MountUnmountFlags C.GMountUnmountFlags

const (
	// No flags set.
	MOUNT_UNMOUNT_NONE MountUnmountFlags = 0

	// Unmount even if there are outstanding
	// file operations on the mount.
	MOUNT_UNMOUNT_FORCE MountUnmountFlags = 1
)

// GOutputStreamSpliceFlags determine how streams should be spliced.
type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags

const (
	// Do not close either stream.
	OUTPUT_STREAM_SPLICE_NONE OutputStreamSpliceFlags = 0

	// Close the source stream after
	// the splice.
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = 1

	// Close the target stream after
	// the splice.
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET OutputStreamSpliceFlags = 2
)

// Flags used when creating a binding. These flags determine in which
// direction the binding works. The default is to synchronize in both
// directions.
type SettingsBindFlags C.GSettingsBindFlags

const (
	// Equivalent to `G_SETTINGS_BIND_GET|G_SETTINGS_BIND_SET`
	SETTINGS_BIND_DEFAULT SettingsBindFlags = 0

	// Update the #GObject property when the setting changes.
	// It is an error to use this flag if the property is not writable.
	SETTINGS_BIND_GET SettingsBindFlags = 1

	// Update the setting when the #GObject property changes.
	// It is an error to use this flag if the property is not readable.
	SETTINGS_BIND_SET SettingsBindFlags = 2

	// Do not try to bind a "sensitivity" property to the writability of the setting
	SETTINGS_BIND_NO_SENSITIVITY SettingsBindFlags = 4

	// When set in addition to #G_SETTINGS_BIND_GET, set the #GObject property
	// value initially from the setting, but do not listen for changes of the setting
	SETTINGS_BIND_GET_NO_CHANGES SettingsBindFlags = 8

	// When passed to g_settings_bind(), uses a pair of mapping functions that invert
	// the boolean value when mapping between the setting and the property.  The setting and property must both
	// be booleans.  You cannot pass this flag to g_settings_bind_with_mapping().
	SETTINGS_BIND_INVERT_BOOLEAN SettingsBindFlags = 16
)
