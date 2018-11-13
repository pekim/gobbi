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

// #GDataStreamByteOrder is used to ensure proper endianness of streaming data sources
// across various machine architectures.
type DataStreamByteOrder C.GDataStreamByteOrder

const (
	// Selects Big Endian byte order.
	DATA_STREAM_BYTE_ORDER_BIG_ENDIAN DataStreamByteOrder = 0
	// Selects Little Endian byte order.
	DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN DataStreamByteOrder = 1
	// Selects endianness based on host machine's architecture.
	DATA_STREAM_BYTE_ORDER_HOST_ENDIAN DataStreamByteOrder = 2
)

// #GDataStreamNewlineType is used when checking for or setting the line endings for a given file.
type DataStreamNewlineType C.GDataStreamNewlineType

const (
	// Selects "LF" line endings, common on most modern UNIX platforms.
	DATA_STREAM_NEWLINE_TYPE_LF DataStreamNewlineType = 0
	// Selects "CR" line endings.
	DATA_STREAM_NEWLINE_TYPE_CR DataStreamNewlineType = 1
	// Selects "CR, LF" line ending, common on Microsoft Windows.
	DATA_STREAM_NEWLINE_TYPE_CR_LF DataStreamNewlineType = 2
	// Automatically try to handle any line ending type.
	DATA_STREAM_NEWLINE_TYPE_ANY DataStreamNewlineType = 3
)

// Used by g_file_set_attributes_from_info() when setting file attributes.
type FileAttributeStatus C.GFileAttributeStatus

const (
	// Attribute value is unset (empty).
	FILE_ATTRIBUTE_STATUS_UNSET FileAttributeStatus = 0
	// Attribute value is set.
	FILE_ATTRIBUTE_STATUS_SET FileAttributeStatus = 1
	// Indicates an error in setting the value.
	FILE_ATTRIBUTE_STATUS_ERROR_SETTING FileAttributeStatus = 2
)

// The data types for file attributes.
type FileAttributeType C.GFileAttributeType

const (
	// indicates an invalid or uninitalized type.
	FILE_ATTRIBUTE_TYPE_INVALID FileAttributeType = 0
	// a null terminated UTF8 string.
	FILE_ATTRIBUTE_TYPE_STRING FileAttributeType = 1
	// a zero terminated string of non-zero bytes.
	FILE_ATTRIBUTE_TYPE_BYTE_STRING FileAttributeType = 2
	// a boolean value.
	FILE_ATTRIBUTE_TYPE_BOOLEAN FileAttributeType = 3
	// an unsigned 4-byte/32-bit integer.
	FILE_ATTRIBUTE_TYPE_UINT32 FileAttributeType = 4
	// a signed 4-byte/32-bit integer.
	FILE_ATTRIBUTE_TYPE_INT32 FileAttributeType = 5
	// an unsigned 8-byte/64-bit integer.
	FILE_ATTRIBUTE_TYPE_UINT64 FileAttributeType = 6
	// a signed 8-byte/64-bit integer.
	FILE_ATTRIBUTE_TYPE_INT64 FileAttributeType = 7
	// a #GObject.
	FILE_ATTRIBUTE_TYPE_OBJECT FileAttributeType = 8
	// a %NULL terminated char **. Since 2.22
	FILE_ATTRIBUTE_TYPE_STRINGV FileAttributeType = 9
)

// Specifies what type of event a monitor event is.
type FileMonitorEvent C.GFileMonitorEvent

const (
	// a file changed.
	FILE_MONITOR_EVENT_CHANGED FileMonitorEvent = 0
	// a hint that this was probably the last change in a set of changes.
	FILE_MONITOR_EVENT_CHANGES_DONE_HINT FileMonitorEvent = 1
	// a file was deleted.
	FILE_MONITOR_EVENT_DELETED FileMonitorEvent = 2
	// a file was created.
	FILE_MONITOR_EVENT_CREATED FileMonitorEvent = 3
	// a file attribute was changed.
	FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED FileMonitorEvent = 4
	// the file location will soon be unmounted.
	FILE_MONITOR_EVENT_PRE_UNMOUNT FileMonitorEvent = 5
	// the file location was unmounted.
	FILE_MONITOR_EVENT_UNMOUNTED FileMonitorEvent = 6
	// the file was moved -- only sent if the
	// (deprecated) %G_FILE_MONITOR_SEND_MOVED flag is set
	FILE_MONITOR_EVENT_MOVED FileMonitorEvent = 7
	// the file was renamed within the
	// current directory -- only sent if the %G_FILE_MONITOR_WATCH_MOVES
	// flag is set.  Since: 2.46.
	FILE_MONITOR_EVENT_RENAMED FileMonitorEvent = 8
	// the file was moved into the
	// monitored directory from another location -- only sent if the
	// %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46.
	FILE_MONITOR_EVENT_MOVED_IN FileMonitorEvent = 9
	// the file was moved out of the
	// monitored directory to another location -- only sent if the
	// %G_FILE_MONITOR_WATCH_MOVES flag is set.  Since: 2.46
	FILE_MONITOR_EVENT_MOVED_OUT FileMonitorEvent = 10
)

// Indicates the file's on-disk type.
type FileType C.GFileType

const (
	// File's type is unknown.
	FILE_TYPE_UNKNOWN FileType = 0
	// File handle represents a regular file.
	FILE_TYPE_REGULAR FileType = 1
	// File handle represents a directory.
	FILE_TYPE_DIRECTORY FileType = 2
	// File handle represents a symbolic link
	// (Unix systems).
	FILE_TYPE_SYMBOLIC_LINK FileType = 3
	// File is a "special" file, such as a socket, fifo,
	// block device, or character device.
	FILE_TYPE_SPECIAL FileType = 4
	// File is a shortcut (Windows systems).
	FILE_TYPE_SHORTCUT FileType = 5
	// File is a mountable location.
	FILE_TYPE_MOUNTABLE FileType = 6
)

// Indicates a hint from the file system whether files should be
// previewed in a file manager. Returned as the value of the key
// #G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.
type FilesystemPreviewType C.GFilesystemPreviewType

const (
	// Only preview files if user has explicitly requested it.
	FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS FilesystemPreviewType = 0
	// Preview files if user has requested preview of "local" files.
	FILESYSTEM_PREVIEW_TYPE_IF_LOCAL FilesystemPreviewType = 1
	// Never preview files.
	FILESYSTEM_PREVIEW_TYPE_NEVER FilesystemPreviewType = 2
)

// Error codes returned by GIO functions.
//
// Note that this domain may be extended in future GLib releases. In
// general, new error codes either only apply to new APIs, or else
// replace %G_IO_ERROR_FAILED in cases that were not explicitly
// distinguished before. You should therefore avoid writing code like
// |[<!-- language="C" -->
// if (g_error_matches (error, G_IO_ERROR, G_IO_ERROR_FAILED))
// {
// Assume that this is EPRINTERONFIRE
// ...
// }
// ]|
// but should instead treat all unrecognized error codes the same as
// #G_IO_ERROR_FAILED.
type IOErrorEnum C.GIOErrorEnum

const (
	// Generic error condition for when an operation fails
	// and no more specific #GIOErrorEnum value is defined.
	IO_ERROR_FAILED IOErrorEnum = 0
	// File not found.
	IO_ERROR_NOT_FOUND IOErrorEnum = 1
	// File already exists.
	IO_ERROR_EXISTS IOErrorEnum = 2
	// File is a directory.
	IO_ERROR_IS_DIRECTORY IOErrorEnum = 3
	// File is not a directory.
	IO_ERROR_NOT_DIRECTORY IOErrorEnum = 4
	// File is a directory that isn't empty.
	IO_ERROR_NOT_EMPTY IOErrorEnum = 5
	// File is not a regular file.
	IO_ERROR_NOT_REGULAR_FILE IOErrorEnum = 6
	// File is not a symbolic link.
	IO_ERROR_NOT_SYMBOLIC_LINK IOErrorEnum = 7
	// File cannot be mounted.
	IO_ERROR_NOT_MOUNTABLE_FILE IOErrorEnum = 8
	// Filename is too many characters.
	IO_ERROR_FILENAME_TOO_LONG IOErrorEnum = 9
	// Filename is invalid or contains invalid characters.
	IO_ERROR_INVALID_FILENAME IOErrorEnum = 10
	// File contains too many symbolic links.
	IO_ERROR_TOO_MANY_LINKS IOErrorEnum = 11
	// No space left on drive.
	IO_ERROR_NO_SPACE IOErrorEnum = 12
	// Invalid argument.
	IO_ERROR_INVALID_ARGUMENT IOErrorEnum = 13
	// Permission denied.
	IO_ERROR_PERMISSION_DENIED IOErrorEnum = 14
	// Operation (or one of its parameters) not supported
	IO_ERROR_NOT_SUPPORTED IOErrorEnum = 15
	// File isn't mounted.
	IO_ERROR_NOT_MOUNTED IOErrorEnum = 16
	// File is already mounted.
	IO_ERROR_ALREADY_MOUNTED IOErrorEnum = 17
	// File was closed.
	IO_ERROR_CLOSED IOErrorEnum = 18
	// Operation was cancelled. See #GCancellable.
	IO_ERROR_CANCELLED IOErrorEnum = 19
	// Operations are still pending.
	IO_ERROR_PENDING IOErrorEnum = 20
	// File is read only.
	IO_ERROR_READ_ONLY IOErrorEnum = 21
	// Backup couldn't be created.
	IO_ERROR_CANT_CREATE_BACKUP IOErrorEnum = 22
	// File's Entity Tag was incorrect.
	IO_ERROR_WRONG_ETAG IOErrorEnum = 23
	// Operation timed out.
	IO_ERROR_TIMED_OUT IOErrorEnum = 24
	// Operation would be recursive.
	IO_ERROR_WOULD_RECURSE IOErrorEnum = 25
	// File is busy.
	IO_ERROR_BUSY IOErrorEnum = 26
	// Operation would block.
	IO_ERROR_WOULD_BLOCK IOErrorEnum = 27
	// Host couldn't be found (remote operations).
	IO_ERROR_HOST_NOT_FOUND IOErrorEnum = 28
	// Operation would merge files.
	IO_ERROR_WOULD_MERGE IOErrorEnum = 29
	// Operation failed and a helper program has
	// already interacted with the user. Do not display any error dialog.
	IO_ERROR_FAILED_HANDLED IOErrorEnum = 30
	// The current process has too many files
	// open and can't open any more. Duplicate descriptors do count toward
	// this limit. Since 2.20
	IO_ERROR_TOO_MANY_OPEN_FILES IOErrorEnum = 31
	// The object has not been initialized. Since 2.22
	IO_ERROR_NOT_INITIALIZED IOErrorEnum = 32
	// The requested address is already in use. Since 2.22
	IO_ERROR_ADDRESS_IN_USE IOErrorEnum = 33
	// Need more input to finish operation. Since 2.24
	IO_ERROR_PARTIAL_INPUT IOErrorEnum = 34
	// The input data was invalid. Since 2.24
	IO_ERROR_INVALID_DATA IOErrorEnum = 35
	// A remote object generated an error that
	// doesn't correspond to a locally registered #GError error
	// domain. Use g_dbus_error_get_remote_error() to extract the D-Bus
	// error name and g_dbus_error_strip_remote_error() to fix up the
	// message so it matches what was received on the wire. Since 2.26.
	IO_ERROR_DBUS_ERROR IOErrorEnum = 36
	// Host unreachable. Since 2.26
	IO_ERROR_HOST_UNREACHABLE IOErrorEnum = 37
	// Network unreachable. Since 2.26
	IO_ERROR_NETWORK_UNREACHABLE IOErrorEnum = 38
	// Connection refused. Since 2.26
	IO_ERROR_CONNECTION_REFUSED IOErrorEnum = 39
	// Connection to proxy server failed. Since 2.26
	IO_ERROR_PROXY_FAILED IOErrorEnum = 40
	// Proxy authentication failed. Since 2.26
	IO_ERROR_PROXY_AUTH_FAILED IOErrorEnum = 41
	// Proxy server needs authentication. Since 2.26
	IO_ERROR_PROXY_NEED_AUTH IOErrorEnum = 42
	// Proxy connection is not allowed by ruleset.
	// Since 2.26
	IO_ERROR_PROXY_NOT_ALLOWED IOErrorEnum = 43
	// Broken pipe. Since 2.36
	IO_ERROR_BROKEN_PIPE IOErrorEnum = 44
	// Connection closed by peer. Note that this
	// is the same code as %G_IO_ERROR_BROKEN_PIPE; before 2.44 some
	// "connection closed" errors returned %G_IO_ERROR_BROKEN_PIPE, but others
	// returned %G_IO_ERROR_FAILED. Now they should all return the same
	// value, which has this more logical name. Since 2.44.
	IO_ERROR_CONNECTION_CLOSED IOErrorEnum = 44
	// Transport endpoint is not connected. Since 2.44
	IO_ERROR_NOT_CONNECTED IOErrorEnum = 45
	// Message too large. Since 2.48.
	IO_ERROR_MESSAGE_TOO_LARGE IOErrorEnum = 46
)

// #GMountOperationResult is returned as a result when a request for
// information is send by the mounting operation.
type MountOperationResult C.GMountOperationResult

const (
	// The request was fulfilled and the
	// user specified data is now available
	MOUNT_OPERATION_HANDLED MountOperationResult = 0
	// The user requested the mount operation
	// to be aborted
	MOUNT_OPERATION_ABORTED MountOperationResult = 1
	// The request was unhandled (i.e. not
	// implemented)
	MOUNT_OPERATION_UNHANDLED MountOperationResult = 2
)

// #GPasswordSave is used to indicate the lifespan of a saved password.
//
// #Gvfs stores passwords in the Gnome keyring when this flag allows it
// to, and later retrieves it again from there.
type PasswordSave C.GPasswordSave

const (
	// never save a password.
	PASSWORD_SAVE_NEVER PasswordSave = 0
	// save a password for the session.
	PASSWORD_SAVE_FOR_SESSION PasswordSave = 1
	// save a password permanently.
	PASSWORD_SAVE_PERMANENTLY PasswordSave = 2
)
