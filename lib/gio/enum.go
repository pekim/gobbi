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

type DataStreamByteOrder C.GDataStreamByteOrder

const (
	DATA_STREAM_BYTE_ORDER_BIG_ENDIAN    DataStreamByteOrder = 0
	DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN DataStreamByteOrder = 1
	DATA_STREAM_BYTE_ORDER_HOST_ENDIAN   DataStreamByteOrder = 2
)

type DataStreamNewlineType C.GDataStreamNewlineType

const (
	DATA_STREAM_NEWLINE_TYPE_LF    DataStreamNewlineType = 0
	DATA_STREAM_NEWLINE_TYPE_CR    DataStreamNewlineType = 1
	DATA_STREAM_NEWLINE_TYPE_CR_LF DataStreamNewlineType = 2
	DATA_STREAM_NEWLINE_TYPE_ANY   DataStreamNewlineType = 3
)

type FileAttributeStatus C.GFileAttributeStatus

const (
	FILE_ATTRIBUTE_STATUS_UNSET         FileAttributeStatus = 0
	FILE_ATTRIBUTE_STATUS_SET           FileAttributeStatus = 1
	FILE_ATTRIBUTE_STATUS_ERROR_SETTING FileAttributeStatus = 2
)

type FileAttributeType C.GFileAttributeType

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

type FileMonitorEvent C.GFileMonitorEvent

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

type FileType C.GFileType

const (
	FILE_TYPE_UNKNOWN       FileType = 0
	FILE_TYPE_REGULAR       FileType = 1
	FILE_TYPE_DIRECTORY     FileType = 2
	FILE_TYPE_SYMBOLIC_LINK FileType = 3
	FILE_TYPE_SPECIAL       FileType = 4
	FILE_TYPE_SHORTCUT      FileType = 5
	FILE_TYPE_MOUNTABLE     FileType = 6
)

type FilesystemPreviewType C.GFilesystemPreviewType

const (
	FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS FilesystemPreviewType = 0
	FILESYSTEM_PREVIEW_TYPE_IF_LOCAL  FilesystemPreviewType = 1
	FILESYSTEM_PREVIEW_TYPE_NEVER     FilesystemPreviewType = 2
)

type IOErrorEnum C.GIOErrorEnum

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

type MountOperationResult C.GMountOperationResult

const (
	MOUNT_OPERATION_HANDLED   MountOperationResult = 0
	MOUNT_OPERATION_ABORTED   MountOperationResult = 1
	MOUNT_OPERATION_UNHANDLED MountOperationResult = 2
)

type PasswordSave C.GPasswordSave

const (
	PASSWORD_SAVE_NEVER       PasswordSave = 0
	PASSWORD_SAVE_FOR_SESSION PasswordSave = 1
	PASSWORD_SAVE_PERMANENTLY PasswordSave = 2
)
