// This is a generated file - DO NOT EDIT

package gio

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

type AppInfoCreateFlags C.GAppInfoCreateFlags

const (
	APP_INFO_CREATE_NONE                          AppInfoCreateFlags = 0
	APP_INFO_CREATE_NEEDS_TERMINAL                AppInfoCreateFlags = 1
	APP_INFO_CREATE_SUPPORTS_URIS                 AppInfoCreateFlags = 2
	APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION AppInfoCreateFlags = 4
)

type AskPasswordFlags C.GAskPasswordFlags

const (
	ASK_PASSWORD_NEED_PASSWORD       AskPasswordFlags = 1
	ASK_PASSWORD_NEED_USERNAME       AskPasswordFlags = 2
	ASK_PASSWORD_NEED_DOMAIN         AskPasswordFlags = 4
	ASK_PASSWORD_SAVING_SUPPORTED    AskPasswordFlags = 8
	ASK_PASSWORD_ANONYMOUS_SUPPORTED AskPasswordFlags = 16
)

type FileAttributeInfoFlags C.GFileAttributeInfoFlags

const (
	FILE_ATTRIBUTE_INFO_NONE            FileAttributeInfoFlags = 0
	FILE_ATTRIBUTE_INFO_COPY_WITH_FILE  FileAttributeInfoFlags = 1
	FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED FileAttributeInfoFlags = 2
)

type FileCopyFlags C.GFileCopyFlags

const (
	FILE_COPY_NONE                 FileCopyFlags = 0
	FILE_COPY_OVERWRITE            FileCopyFlags = 1
	FILE_COPY_BACKUP               FileCopyFlags = 2
	FILE_COPY_NOFOLLOW_SYMLINKS    FileCopyFlags = 4
	FILE_COPY_ALL_METADATA         FileCopyFlags = 8
	FILE_COPY_NO_FALLBACK_FOR_MOVE FileCopyFlags = 16
	FILE_COPY_TARGET_DEFAULT_PERMS FileCopyFlags = 32
)

type FileCreateFlags C.GFileCreateFlags

const (
	FILE_CREATE_NONE                FileCreateFlags = 0
	FILE_CREATE_PRIVATE             FileCreateFlags = 1
	FILE_CREATE_REPLACE_DESTINATION FileCreateFlags = 2
)

type FileMonitorFlags C.GFileMonitorFlags

const (
	FILE_MONITOR_NONE             FileMonitorFlags = 0
	FILE_MONITOR_WATCH_MOUNTS     FileMonitorFlags = 1
	FILE_MONITOR_SEND_MOVED       FileMonitorFlags = 2
	FILE_MONITOR_WATCH_HARD_LINKS FileMonitorFlags = 4
	FILE_MONITOR_WATCH_MOVES      FileMonitorFlags = 8
)

type FileQueryInfoFlags C.GFileQueryInfoFlags

const (
	FILE_QUERY_INFO_NONE              FileQueryInfoFlags = 0
	FILE_QUERY_INFO_NOFOLLOW_SYMLINKS FileQueryInfoFlags = 1
)

type MountMountFlags C.GMountMountFlags

const (
	MOUNT_MOUNT_NONE MountMountFlags = 0
)

type MountUnmountFlags C.GMountUnmountFlags

const (
	MOUNT_UNMOUNT_NONE  MountUnmountFlags = 0
	MOUNT_UNMOUNT_FORCE MountUnmountFlags = 1
)

type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags

const (
	OUTPUT_STREAM_SPLICE_NONE         OutputStreamSpliceFlags = 0
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE OutputStreamSpliceFlags = 1
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET OutputStreamSpliceFlags = 2
)

type SettingsBindFlags C.GSettingsBindFlags

const (
	SETTINGS_BIND_DEFAULT        SettingsBindFlags = 0
	SETTINGS_BIND_GET            SettingsBindFlags = 1
	SETTINGS_BIND_SET            SettingsBindFlags = 2
	SETTINGS_BIND_NO_SENSITIVITY SettingsBindFlags = 4
	SETTINGS_BIND_GET_NO_CHANGES SettingsBindFlags = 8
	SETTINGS_BIND_INVERT_BOOLEAN SettingsBindFlags = 16
)
