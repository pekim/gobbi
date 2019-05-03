// Code generated - DO NOT EDIT.
// +build !gio_2.18,!gio_2.20,!gio_2.22,!gio_2.24,!gio_2.26,!gio_2.28,!gio_2.30,!gio_2.32,!gio_2.34,!gio_2.36,!gio_2.38,!gio_2.40,!gio_2.42,!gio_2.44,!gio_2.46,!gio_2.48,!gio_2.50,!gio_2.52,!gio_2.54,!gio_2.56

package gio

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

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
/*

	void cancellable_cancelledHandler(GObject *, gpointer);

	static gulong Cancellable_signal_connect_cancelled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancelled", G_CALLBACK(cancellable_cancelledHandler), data);
	}

*/
/*

	void filemonitor_changedHandler(GObject *, GFile *, GFile *, GFileMonitorEvent, gpointer);

	static gulong FileMonitor_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(filemonitor_changedHandler), data);
	}

*/
/*

	void filenamecompleter_gotCompletionDataHandler(GObject *, gpointer);

	static gulong FilenameCompleter_signal_connect_got_completion_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-completion-data", G_CALLBACK(filenamecompleter_gotCompletionDataHandler), data);
	}

*/
/*

	void mountoperation_askPasswordHandler(GObject *, gchar*, gchar*, gchar*, GAskPasswordFlags, gpointer);

	static gulong MountOperation_signal_connect_ask_password(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "ask-password", G_CALLBACK(mountoperation_askPasswordHandler), data);
	}

*/
/*

	void mountoperation_replyHandler(GObject *, GMountOperationResult, gpointer);

	static gulong MountOperation_signal_connect_reply(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "reply", G_CALLBACK(mountoperation_replyHandler), data);
	}

*/
/*

	void resolver_reloadHandler(GObject *, gpointer);

	static gulong Resolver_signal_connect_reload(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "reload", G_CALLBACK(resolver_reloadHandler), data);
	}

*/
/*

	void settings_changedHandler(GObject *, gchar*, gpointer);

	static gulong Settings_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(settings_changedHandler), data);
	}

*/
/*

	gboolean settings_writableChangeEventHandler(GObject *, guint, gpointer);

	static gulong Settings_signal_connect_writable_change_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "writable-change-event", G_CALLBACK(settings_writableChangeEventHandler), data);
	}

*/
/*

	void settings_writableChangedHandler(GObject *, gchar*, gpointer);

	static gulong Settings_signal_connect_writable_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "writable-changed", G_CALLBACK(settings_writableChangedHandler), data);
	}

*/
/*

	static void _g_simple_async_result_set_error(GSimpleAsyncResult* simple, GQuark domain, gint code, const char* format) {
		return g_simple_async_result_set_error(simple, domain, code, format);
    }
*/
/*

	void unixmountmonitor_mountpointsChangedHandler(GObject *, gpointer);

	static gulong UnixMountMonitor_signal_connect_mountpoints_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mountpoints-changed", G_CALLBACK(unixmountmonitor_mountpointsChangedHandler), data);
	}

*/
/*

	void unixmountmonitor_mountsChangedHandler(GObject *, gpointer);

	static gulong UnixMountMonitor_signal_connect_mounts_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mounts-changed", G_CALLBACK(unixmountmonitor_mountsChangedHandler), data);
	}

*/
/*

	void volumemonitor_driveChangedHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-changed", G_CALLBACK(volumemonitor_driveChangedHandler), data);
	}

*/
/*

	void volumemonitor_driveConnectedHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_connected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-connected", G_CALLBACK(volumemonitor_driveConnectedHandler), data);
	}

*/
/*

	void volumemonitor_driveDisconnectedHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_disconnected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-disconnected", G_CALLBACK(volumemonitor_driveDisconnectedHandler), data);
	}

*/
/*

	void volumemonitor_mountAddedHandler(GObject *, GMount *, gpointer);

	static gulong VolumeMonitor_signal_connect_mount_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount-added", G_CALLBACK(volumemonitor_mountAddedHandler), data);
	}

*/
/*

	void volumemonitor_mountChangedHandler(GObject *, GMount *, gpointer);

	static gulong VolumeMonitor_signal_connect_mount_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount-changed", G_CALLBACK(volumemonitor_mountChangedHandler), data);
	}

*/
/*

	void volumemonitor_mountPreUnmountHandler(GObject *, GMount *, gpointer);

	static gulong VolumeMonitor_signal_connect_mount_pre_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount-pre-unmount", G_CALLBACK(volumemonitor_mountPreUnmountHandler), data);
	}

*/
/*

	void volumemonitor_mountRemovedHandler(GObject *, GMount *, gpointer);

	static gulong VolumeMonitor_signal_connect_mount_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount-removed", G_CALLBACK(volumemonitor_mountRemovedHandler), data);
	}

*/
/*

	void volumemonitor_volumeAddedHandler(GObject *, GVolume *, gpointer);

	static gulong VolumeMonitor_signal_connect_volume_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "volume-added", G_CALLBACK(volumemonitor_volumeAddedHandler), data);
	}

*/
/*

	void volumemonitor_volumeChangedHandler(GObject *, GVolume *, gpointer);

	static gulong VolumeMonitor_signal_connect_volume_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "volume-changed", G_CALLBACK(volumemonitor_volumeChangedHandler), data);
	}

*/
/*

	void volumemonitor_volumeRemovedHandler(GObject *, GVolume *, gpointer);

	static gulong VolumeMonitor_signal_connect_volume_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "volume-removed", G_CALLBACK(volumemonitor_volumeRemovedHandler), data);
	}

*/
/*

	void drive_changedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(drive_changedHandler), data);
	}

*/
/*

	void drive_disconnectedHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_disconnected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnected", G_CALLBACK(drive_disconnectedHandler), data);
	}

*/
/*

	void drive_ejectButtonHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_eject_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "eject-button", G_CALLBACK(drive_ejectButtonHandler), data);
	}

*/
/*

	void mount_changedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(mount_changedHandler), data);
	}

*/
/*

	void mount_unmountedHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_unmounted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmounted", G_CALLBACK(mount_unmountedHandler), data);
	}

*/
/*

	void volume_changedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(volume_changedHandler), data);
	}

*/
/*

	void volume_removedHandler(GObject *, gpointer);

	static gulong Volume_signal_connect_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "removed", G_CALLBACK(volume_removedHandler), data);
	}

*/
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

// AppLaunchContext is a wrapper around the C record GAppLaunchContext.
type AppLaunchContext struct {
	native *C.GAppLaunchContext
	// parent_instance : record
	// Private : priv
}

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	c := (*C.GAppLaunchContext)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppLaunchContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppLaunchContext with another AppLaunchContext, and returns true if they represent the same GObject.
func (recv *AppLaunchContext) Equals(other *AppLaunchContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AppLaunchContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AppLaunchContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppLaunchContext.
func CastToAppLaunchContext(object *gobject.Object) *AppLaunchContext {
	return AppLaunchContextNewFromC(object.ToC())
}

// AppLaunchContextNew is a wrapper around the C function g_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.g_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetDisplay is a wrapper around the C function g_app_launch_context_get_display.
func (recv *AppLaunchContext) GetDisplay(info *AppInfo, files *glib.List) string {
	c_info := (*C.GAppInfo)(info.ToC())

	c_files := (*C.GList)(C.NULL)
	if files != nil {
		c_files = (*C.GList)(files.ToC())
	}

	retC := C.g_app_launch_context_get_display((*C.GAppLaunchContext)(recv.native), c_info, c_files)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetStartupNotifyId is a wrapper around the C function g_app_launch_context_get_startup_notify_id.
func (recv *AppLaunchContext) GetStartupNotifyId(info *AppInfo, files *glib.List) string {
	c_info := (*C.GAppInfo)(info.ToC())

	c_files := (*C.GList)(C.NULL)
	if files != nil {
		c_files = (*C.GList)(files.ToC())
	}

	retC := C.g_app_launch_context_get_startup_notify_id((*C.GAppLaunchContext)(recv.native), c_info, c_files)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// LaunchFailed is a wrapper around the C function g_app_launch_context_launch_failed.
func (recv *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	c_startup_notify_id := C.CString(startupNotifyId)
	defer C.free(unsafe.Pointer(c_startup_notify_id))

	C.g_app_launch_context_launch_failed((*C.GAppLaunchContext)(recv.native), c_startup_notify_id)

	return
}

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native *C.GApplicationCommandLine
	// Private : parent_instance
	// Private : priv
}

func ApplicationCommandLineNewFromC(u unsafe.Pointer) *ApplicationCommandLine {
	c := (*C.GApplicationCommandLine)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLine{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ApplicationCommandLine) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationCommandLine with another ApplicationCommandLine, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLine) Equals(other *ApplicationCommandLine) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ApplicationCommandLine) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ApplicationCommandLine.
// Exercise care, as this is a potentially dangerous function if the Object is not a ApplicationCommandLine.
func CastToApplicationCommandLine(object *gobject.Object) *ApplicationCommandLine {
	return ApplicationCommandLineNewFromC(object.ToC())
}

// BufferedInputStream is a wrapper around the C record GBufferedInputStream.
type BufferedInputStream struct {
	native *C.GBufferedInputStream
	// parent_instance : record
	// Private : priv
}

func BufferedInputStreamNewFromC(u unsafe.Pointer) *BufferedInputStream {
	c := (*C.GBufferedInputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BufferedInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BufferedInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedInputStream with another BufferedInputStream, and returns true if they represent the same GObject.
func (recv *BufferedInputStream) Equals(other *BufferedInputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *BufferedInputStream) FilterInputStream() *FilterInputStream {
	return FilterInputStreamNewFromC(unsafe.Pointer(recv.native))
}

// InputStream upcasts to *InputStream
func (recv *BufferedInputStream) InputStream() *InputStream {
	return recv.FilterInputStream().InputStream()
}

// Object upcasts to *Object
func (recv *BufferedInputStream) Object() *gobject.Object {
	return recv.FilterInputStream().Object()
}

// CastToWidget down casts any arbitrary Object to BufferedInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a BufferedInputStream.
func CastToBufferedInputStream(object *gobject.Object) *BufferedInputStream {
	return BufferedInputStreamNewFromC(object.ToC())
}

// BufferedInputStreamNew is a wrapper around the C function g_buffered_input_stream_new.
func BufferedInputStreamNew(baseStream *InputStream) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_input_stream_new(c_base_stream)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferedInputStreamNewSized is a wrapper around the C function g_buffered_input_stream_new_sized.
func BufferedInputStreamNewSized(baseStream *InputStream, size uint64) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_input_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Fill is a wrapper around the C function g_buffered_input_stream_fill.
func (recv *BufferedInputStream) Fill(count int64, cancellable *Cancellable) (int64, error) {
	c_count := (C.gssize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_fill((*C.GBufferedInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_buffered_input_stream_fill_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// FillFinish is a wrapper around the C function g_buffered_input_stream_fill_finish.
func (recv *BufferedInputStream) FillFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_fill_finish((*C.GBufferedInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAvailable is a wrapper around the C function g_buffered_input_stream_get_available.
func (recv *BufferedInputStream) GetAvailable() uint64 {
	retC := C.g_buffered_input_stream_get_available((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetBufferSize is a wrapper around the C function g_buffered_input_stream_get_buffer_size.
func (recv *BufferedInputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_input_stream_get_buffer_size((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Peek is a wrapper around the C function g_buffered_input_stream_peek.
func (recv *BufferedInputStream) Peek(buffer []uint8, offset uint64) uint64 {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (unsafe.Pointer(c_buffer_arrayPtr))

	c_offset := (C.gsize)(offset)

	c_count := (C.gsize)(len(buffer))

	retC := C.g_buffered_input_stream_peek((*C.GBufferedInputStream)(recv.native), c_buffer, c_offset, c_count)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_buffered_input_stream_peek_buffer : array return type :

// ReadByte is a wrapper around the C function g_buffered_input_stream_read_byte.
func (recv *BufferedInputStream) ReadByte(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_read_byte((*C.GBufferedInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBufferSize is a wrapper around the C function g_buffered_input_stream_set_buffer_size.
func (recv *BufferedInputStream) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_buffered_input_stream_set_buffer_size((*C.GBufferedInputStream)(recv.native), c_size)

	return
}

// Seekable returns the Seekable interface implemented by BufferedInputStream
func (recv *BufferedInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// BufferedOutputStream is a wrapper around the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native *C.GBufferedOutputStream
	// parent_instance : record
	// priv : record
}

func BufferedOutputStreamNewFromC(u unsafe.Pointer) *BufferedOutputStream {
	c := (*C.GBufferedOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BufferedOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BufferedOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedOutputStream with another BufferedOutputStream, and returns true if they represent the same GObject.
func (recv *BufferedOutputStream) Equals(other *BufferedOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *BufferedOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// OutputStream upcasts to *OutputStream
func (recv *BufferedOutputStream) OutputStream() *OutputStream {
	return recv.FilterOutputStream().OutputStream()
}

// Object upcasts to *Object
func (recv *BufferedOutputStream) Object() *gobject.Object {
	return recv.FilterOutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to BufferedOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a BufferedOutputStream.
func CastToBufferedOutputStream(object *gobject.Object) *BufferedOutputStream {
	return BufferedOutputStreamNewFromC(object.ToC())
}

// BufferedOutputStreamNew is a wrapper around the C function g_buffered_output_stream_new.
func BufferedOutputStreamNew(baseStream *OutputStream) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_output_stream_new(c_base_stream)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferedOutputStreamNewSized is a wrapper around the C function g_buffered_output_stream_new_sized.
func BufferedOutputStreamNewSized(baseStream *OutputStream, size uint64) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_output_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetAutoGrow is a wrapper around the C function g_buffered_output_stream_get_auto_grow.
func (recv *BufferedOutputStream) GetAutoGrow() bool {
	retC := C.g_buffered_output_stream_get_auto_grow((*C.GBufferedOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBufferSize is a wrapper around the C function g_buffered_output_stream_get_buffer_size.
func (recv *BufferedOutputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_output_stream_get_buffer_size((*C.GBufferedOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// SetAutoGrow is a wrapper around the C function g_buffered_output_stream_set_auto_grow.
func (recv *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	c_auto_grow :=
		boolToGboolean(autoGrow)

	C.g_buffered_output_stream_set_auto_grow((*C.GBufferedOutputStream)(recv.native), c_auto_grow)

	return
}

// SetBufferSize is a wrapper around the C function g_buffered_output_stream_set_buffer_size.
func (recv *BufferedOutputStream) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_buffered_output_stream_set_buffer_size((*C.GBufferedOutputStream)(recv.native), c_size)

	return
}

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// BytesIcon is a wrapper around the C record GBytesIcon.
type BytesIcon struct {
	native *C.GBytesIcon
}

func BytesIconNewFromC(u unsafe.Pointer) *BytesIcon {
	c := (*C.GBytesIcon)(u)
	if c == nil {
		return nil
	}

	g := &BytesIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BytesIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BytesIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BytesIcon with another BytesIcon, and returns true if they represent the same GObject.
func (recv *BytesIcon) Equals(other *BytesIcon) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *BytesIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to BytesIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a BytesIcon.
func CastToBytesIcon(object *gobject.Object) *BytesIcon {
	return BytesIconNewFromC(object.ToC())
}

// Icon returns the Icon interface implemented by BytesIcon
func (recv *BytesIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by BytesIcon
func (recv *BytesIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Cancellable is a wrapper around the C record GCancellable.
type Cancellable struct {
	native *C.GCancellable
	// parent_instance : record
	// Private : priv
}

func CancellableNewFromC(u unsafe.Pointer) *Cancellable {
	c := (*C.GCancellable)(u)
	if c == nil {
		return nil
	}

	g := &Cancellable{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Cancellable) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Cancellable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Cancellable with another Cancellable, and returns true if they represent the same GObject.
func (recv *Cancellable) Equals(other *Cancellable) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Cancellable) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Cancellable.
// Exercise care, as this is a potentially dangerous function if the Object is not a Cancellable.
func CastToCancellable(object *gobject.Object) *Cancellable {
	return CancellableNewFromC(object.ToC())
}

type signalCancellableCancelledDetail struct {
	callback  CancellableSignalCancelledCallback
	handlerID C.gulong
}

var signalCancellableCancelledId int
var signalCancellableCancelledMap = make(map[int]signalCancellableCancelledDetail)
var signalCancellableCancelledLock sync.RWMutex

// CancellableSignalCancelledCallback is a callback function for a 'cancelled' signal emitted from a Cancellable.
type CancellableSignalCancelledCallback func()

/*
ConnectCancelled connects the callback to the 'cancelled' signal for the Cancellable.

The returned value represents the connection, and may be passed to DisconnectCancelled to remove it.
*/
func (recv *Cancellable) ConnectCancelled(callback CancellableSignalCancelledCallback) int {
	signalCancellableCancelledLock.Lock()
	defer signalCancellableCancelledLock.Unlock()

	signalCancellableCancelledId++
	instance := C.gpointer(recv.native)
	handlerID := C.Cancellable_signal_connect_cancelled(instance, C.gpointer(uintptr(signalCancellableCancelledId)))

	detail := signalCancellableCancelledDetail{callback, handlerID}
	signalCancellableCancelledMap[signalCancellableCancelledId] = detail

	return signalCancellableCancelledId
}

/*
DisconnectCancelled disconnects a callback from the 'cancelled' signal for the Cancellable.

The connectionID should be a value returned from a call to ConnectCancelled.
*/
func (recv *Cancellable) DisconnectCancelled(connectionID int) {
	signalCancellableCancelledLock.Lock()
	defer signalCancellableCancelledLock.Unlock()

	detail, exists := signalCancellableCancelledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCancellableCancelledMap, connectionID)
}

//export cancellable_cancelledHandler
func cancellable_cancelledHandler(_ *C.GObject, data C.gpointer) {
	signalCancellableCancelledLock.RLock()
	defer signalCancellableCancelledLock.RUnlock()

	index := int(uintptr(data))
	callback := signalCancellableCancelledMap[index].callback
	callback()
}

// CancellableNew is a wrapper around the C function g_cancellable_new.
func CancellableNew() *Cancellable {
	retC := C.g_cancellable_new()
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CancellableGetCurrent is a wrapper around the C function g_cancellable_get_current.
func CancellableGetCurrent() *Cancellable {
	retC := C.g_cancellable_get_current()
	var retGo (*Cancellable)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CancellableNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Cancel is a wrapper around the C function g_cancellable_cancel.
func (recv *Cancellable) Cancel() {
	C.g_cancellable_cancel((*C.GCancellable)(recv.native))

	return
}

// GetFd is a wrapper around the C function g_cancellable_get_fd.
func (recv *Cancellable) GetFd() int32 {
	retC := C.g_cancellable_get_fd((*C.GCancellable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsCancelled is a wrapper around the C function g_cancellable_is_cancelled.
func (recv *Cancellable) IsCancelled() bool {
	retC := C.g_cancellable_is_cancelled((*C.GCancellable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PopCurrent is a wrapper around the C function g_cancellable_pop_current.
func (recv *Cancellable) PopCurrent() {
	C.g_cancellable_pop_current((*C.GCancellable)(recv.native))

	return
}

// PushCurrent is a wrapper around the C function g_cancellable_push_current.
func (recv *Cancellable) PushCurrent() {
	C.g_cancellable_push_current((*C.GCancellable)(recv.native))

	return
}

// Reset is a wrapper around the C function g_cancellable_reset.
func (recv *Cancellable) Reset() {
	C.g_cancellable_reset((*C.GCancellable)(recv.native))

	return
}

// SetErrorIfCancelled is a wrapper around the C function g_cancellable_set_error_if_cancelled.
func (recv *Cancellable) SetErrorIfCancelled() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_cancellable_set_error_if_cancelled((*C.GCancellable)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CharsetConverter is a wrapper around the C record GCharsetConverter.
type CharsetConverter struct {
	native *C.GCharsetConverter
}

func CharsetConverterNewFromC(u unsafe.Pointer) *CharsetConverter {
	c := (*C.GCharsetConverter)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CharsetConverter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CharsetConverter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CharsetConverter with another CharsetConverter, and returns true if they represent the same GObject.
func (recv *CharsetConverter) Equals(other *CharsetConverter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *CharsetConverter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to CharsetConverter.
// Exercise care, as this is a potentially dangerous function if the Object is not a CharsetConverter.
func CastToCharsetConverter(object *gobject.Object) *CharsetConverter {
	return CharsetConverterNewFromC(object.ToC())
}

// Converter returns the Converter interface implemented by CharsetConverter
func (recv *CharsetConverter) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Initable returns the Initable interface implemented by CharsetConverter
func (recv *CharsetConverter) Initable() *Initable {
	return InitableNewFromC(recv.ToC())
}

// ConverterInputStream is a wrapper around the C record GConverterInputStream.
type ConverterInputStream struct {
	native *C.GConverterInputStream
	// parent_instance : record
	// Private : priv
}

func ConverterInputStreamNewFromC(u unsafe.Pointer) *ConverterInputStream {
	c := (*C.GConverterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ConverterInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ConverterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterInputStream with another ConverterInputStream, and returns true if they represent the same GObject.
func (recv *ConverterInputStream) Equals(other *ConverterInputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *ConverterInputStream) FilterInputStream() *FilterInputStream {
	return FilterInputStreamNewFromC(unsafe.Pointer(recv.native))
}

// InputStream upcasts to *InputStream
func (recv *ConverterInputStream) InputStream() *InputStream {
	return recv.FilterInputStream().InputStream()
}

// Object upcasts to *Object
func (recv *ConverterInputStream) Object() *gobject.Object {
	return recv.FilterInputStream().Object()
}

// CastToWidget down casts any arbitrary Object to ConverterInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a ConverterInputStream.
func CastToConverterInputStream(object *gobject.Object) *ConverterInputStream {
	return ConverterInputStreamNewFromC(object.ToC())
}

// ConverterInputStreamNew is a wrapper around the C function g_converter_input_stream_new.
func ConverterInputStreamNew(baseStream *InputStream, converter *Converter) *ConverterInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_input_stream_new(c_base_stream, c_converter)
	retGo := ConverterInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PollableInputStream returns the PollableInputStream interface implemented by ConverterInputStream
func (recv *ConverterInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// ConverterOutputStream is a wrapper around the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native *C.GConverterOutputStream
	// parent_instance : record
	// Private : priv
}

func ConverterOutputStreamNewFromC(u unsafe.Pointer) *ConverterOutputStream {
	c := (*C.GConverterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ConverterOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ConverterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterOutputStream with another ConverterOutputStream, and returns true if they represent the same GObject.
func (recv *ConverterOutputStream) Equals(other *ConverterOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *ConverterOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// OutputStream upcasts to *OutputStream
func (recv *ConverterOutputStream) OutputStream() *OutputStream {
	return recv.FilterOutputStream().OutputStream()
}

// Object upcasts to *Object
func (recv *ConverterOutputStream) Object() *gobject.Object {
	return recv.FilterOutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to ConverterOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a ConverterOutputStream.
func CastToConverterOutputStream(object *gobject.Object) *ConverterOutputStream {
	return ConverterOutputStreamNewFromC(object.ToC())
}

// ConverterOutputStreamNew is a wrapper around the C function g_converter_output_stream_new.
func ConverterOutputStreamNew(baseStream *OutputStream, converter *Converter) *ConverterOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_output_stream_new(c_base_stream, c_converter)
	retGo := ConverterOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// DBusActionGroup is a wrapper around the C record GDBusActionGroup.
type DBusActionGroup struct {
	native *C.GDBusActionGroup
}

func DBusActionGroupNewFromC(u unsafe.Pointer) *DBusActionGroup {
	c := (*C.GDBusActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &DBusActionGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusActionGroup with another DBusActionGroup, and returns true if they represent the same GObject.
func (recv *DBusActionGroup) Equals(other *DBusActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DBusActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DBusActionGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusActionGroup.
func CastToDBusActionGroup(object *gobject.Object) *DBusActionGroup {
	return DBusActionGroupNewFromC(object.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) ActionGroup() *ActionGroup {
	return ActionGroupNewFromC(recv.ToC())
}

// RemoteActionGroup returns the RemoteActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) RemoteActionGroup() *RemoteActionGroup {
	return RemoteActionGroupNewFromC(recv.ToC())
}

// DBusMenuModel is a wrapper around the C record GDBusMenuModel.
type DBusMenuModel struct {
	native *C.GDBusMenuModel
}

func DBusMenuModelNewFromC(u unsafe.Pointer) *DBusMenuModel {
	c := (*C.GDBusMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &DBusMenuModel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMenuModel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusMenuModel with another DBusMenuModel, and returns true if they represent the same GObject.
func (recv *DBusMenuModel) Equals(other *DBusMenuModel) bool {
	return other.ToC() == recv.ToC()
}

// MenuModel upcasts to *MenuModel
func (recv *DBusMenuModel) MenuModel() *MenuModel {
	return MenuModelNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *DBusMenuModel) Object() *gobject.Object {
	return recv.MenuModel().Object()
}

// CastToWidget down casts any arbitrary Object to DBusMenuModel.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMenuModel.
func CastToDBusMenuModel(object *gobject.Object) *DBusMenuModel {
	return DBusMenuModelNewFromC(object.ToC())
}

// DataInputStream is a wrapper around the C record GDataInputStream.
type DataInputStream struct {
	native *C.GDataInputStream
	// parent_instance : record
	// Private : priv
}

func DataInputStreamNewFromC(u unsafe.Pointer) *DataInputStream {
	c := (*C.GDataInputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DataInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DataInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataInputStream with another DataInputStream, and returns true if they represent the same GObject.
func (recv *DataInputStream) Equals(other *DataInputStream) bool {
	return other.ToC() == recv.ToC()
}

// BufferedInputStream upcasts to *BufferedInputStream
func (recv *DataInputStream) BufferedInputStream() *BufferedInputStream {
	return BufferedInputStreamNewFromC(unsafe.Pointer(recv.native))
}

// FilterInputStream upcasts to *FilterInputStream
func (recv *DataInputStream) FilterInputStream() *FilterInputStream {
	return recv.BufferedInputStream().FilterInputStream()
}

// InputStream upcasts to *InputStream
func (recv *DataInputStream) InputStream() *InputStream {
	return recv.BufferedInputStream().InputStream()
}

// Object upcasts to *Object
func (recv *DataInputStream) Object() *gobject.Object {
	return recv.BufferedInputStream().Object()
}

// CastToWidget down casts any arbitrary Object to DataInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a DataInputStream.
func CastToDataInputStream(object *gobject.Object) *DataInputStream {
	return DataInputStreamNewFromC(object.ToC())
}

// DataInputStreamNew is a wrapper around the C function g_data_input_stream_new.
func DataInputStreamNew(baseStream *InputStream) *DataInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_data_input_stream_new(c_base_stream)
	retGo := DataInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetByteOrder is a wrapper around the C function g_data_input_stream_get_byte_order.
func (recv *DataInputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_input_stream_get_byte_order((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// GetNewlineType is a wrapper around the C function g_data_input_stream_get_newline_type.
func (recv *DataInputStream) GetNewlineType() DataStreamNewlineType {
	retC := C.g_data_input_stream_get_newline_type((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamNewlineType)(retC)

	return retGo
}

// ReadByte is a wrapper around the C function g_data_input_stream_read_byte.
func (recv *DataInputStream) ReadByte(cancellable *Cancellable) (uint8, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_byte((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint8)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadInt16 is a wrapper around the C function g_data_input_stream_read_int16.
func (recv *DataInputStream) ReadInt16(cancellable *Cancellable) (int16, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int16((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int16)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadInt32 is a wrapper around the C function g_data_input_stream_read_int32.
func (recv *DataInputStream) ReadInt32(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int32((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadInt64 is a wrapper around the C function g_data_input_stream_read_int64.
func (recv *DataInputStream) ReadInt64(cancellable *Cancellable) (int64, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int64((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_data_input_stream_read_line : array return type :

// ReadUint16 is a wrapper around the C function g_data_input_stream_read_uint16.
func (recv *DataInputStream) ReadUint16(cancellable *Cancellable) (uint16, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint16((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint16)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadUint32 is a wrapper around the C function g_data_input_stream_read_uint32.
func (recv *DataInputStream) ReadUint32(cancellable *Cancellable) (uint32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint32((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadUint64 is a wrapper around the C function g_data_input_stream_read_uint64.
func (recv *DataInputStream) ReadUint64(cancellable *Cancellable) (uint64, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint64((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadUntil is a wrapper around the C function g_data_input_stream_read_until.
func (recv *DataInputStream) ReadUntil(stopChars string, cancellable *Cancellable) (string, uint64, error) {
	c_stop_chars := C.CString(stopChars)
	defer C.free(unsafe.Pointer(c_stop_chars))

	var c_length C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_until((*C.GDataInputStream)(recv.native), c_stop_chars, &c_length, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// SetByteOrder is a wrapper around the C function g_data_input_stream_set_byte_order.
func (recv *DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_input_stream_set_byte_order((*C.GDataInputStream)(recv.native), c_order)

	return
}

// SetNewlineType is a wrapper around the C function g_data_input_stream_set_newline_type.
func (recv *DataInputStream) SetNewlineType(type_ DataStreamNewlineType) {
	c_type := (C.GDataStreamNewlineType)(type_)

	C.g_data_input_stream_set_newline_type((*C.GDataInputStream)(recv.native), c_type)

	return
}

// Seekable returns the Seekable interface implemented by DataInputStream
func (recv *DataInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// DataOutputStream is a wrapper around the C record GDataOutputStream.
type DataOutputStream struct {
	native *C.GDataOutputStream
	// parent_instance : record
	// Private : priv
}

func DataOutputStreamNewFromC(u unsafe.Pointer) *DataOutputStream {
	c := (*C.GDataOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DataOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DataOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataOutputStream with another DataOutputStream, and returns true if they represent the same GObject.
func (recv *DataOutputStream) Equals(other *DataOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// FilterOutputStream upcasts to *FilterOutputStream
func (recv *DataOutputStream) FilterOutputStream() *FilterOutputStream {
	return FilterOutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// OutputStream upcasts to *OutputStream
func (recv *DataOutputStream) OutputStream() *OutputStream {
	return recv.FilterOutputStream().OutputStream()
}

// Object upcasts to *Object
func (recv *DataOutputStream) Object() *gobject.Object {
	return recv.FilterOutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to DataOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a DataOutputStream.
func CastToDataOutputStream(object *gobject.Object) *DataOutputStream {
	return DataOutputStreamNewFromC(object.ToC())
}

// DataOutputStreamNew is a wrapper around the C function g_data_output_stream_new.
func DataOutputStreamNew(baseStream *OutputStream) *DataOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_data_output_stream_new(c_base_stream)
	retGo := DataOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetByteOrder is a wrapper around the C function g_data_output_stream_get_byte_order.
func (recv *DataOutputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_output_stream_get_byte_order((*C.GDataOutputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// PutByte is a wrapper around the C function g_data_output_stream_put_byte.
func (recv *DataOutputStream) PutByte(data uint8, cancellable *Cancellable) (bool, error) {
	c_data := (C.guchar)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_byte((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutInt16 is a wrapper around the C function g_data_output_stream_put_int16.
func (recv *DataOutputStream) PutInt16(data int16, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint16)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int16((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutInt32 is a wrapper around the C function g_data_output_stream_put_int32.
func (recv *DataOutputStream) PutInt32(data int32, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint32)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int32((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutInt64 is a wrapper around the C function g_data_output_stream_put_int64.
func (recv *DataOutputStream) PutInt64(data int64, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint64)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int64((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutString is a wrapper around the C function g_data_output_stream_put_string.
func (recv *DataOutputStream) PutString(str string, cancellable *Cancellable) (bool, error) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_string((*C.GDataOutputStream)(recv.native), c_str, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutUint16 is a wrapper around the C function g_data_output_stream_put_uint16.
func (recv *DataOutputStream) PutUint16(data uint16, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint16)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint16((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutUint32 is a wrapper around the C function g_data_output_stream_put_uint32.
func (recv *DataOutputStream) PutUint32(data uint32, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint32)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint32((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PutUint64 is a wrapper around the C function g_data_output_stream_put_uint64.
func (recv *DataOutputStream) PutUint64(data uint64, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint64)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint64((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetByteOrder is a wrapper around the C function g_data_output_stream_set_byte_order.
func (recv *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_output_stream_set_byte_order((*C.GDataOutputStream)(recv.native), c_order)

	return
}

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// DesktopAppInfo is a wrapper around the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native *C.GDesktopAppInfo
}

func DesktopAppInfoNewFromC(u unsafe.Pointer) *DesktopAppInfo {
	c := (*C.GDesktopAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DesktopAppInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DesktopAppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DesktopAppInfo with another DesktopAppInfo, and returns true if they represent the same GObject.
func (recv *DesktopAppInfo) Equals(other *DesktopAppInfo) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DesktopAppInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DesktopAppInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a DesktopAppInfo.
func CastToDesktopAppInfo(object *gobject.Object) *DesktopAppInfo {
	return DesktopAppInfoNewFromC(object.ToC())
}

// DesktopAppInfoNew is a wrapper around the C function g_desktop_app_info_new.
func DesktopAppInfoNew(desktopId string) *DesktopAppInfo {
	c_desktop_id := C.CString(desktopId)
	defer C.free(unsafe.Pointer(c_desktop_id))

	retC := C.g_desktop_app_info_new(c_desktop_id)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DesktopAppInfoNewFromFilename is a wrapper around the C function g_desktop_app_info_new_from_filename.
func DesktopAppInfoNewFromFilename(filename string) *DesktopAppInfo {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_desktop_app_info_new_from_filename(c_filename)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// g_desktop_app_info_search : no type for array return
// DesktopAppInfoSetDesktopEnv is a wrapper around the C function g_desktop_app_info_set_desktop_env.
func DesktopAppInfoSetDesktopEnv(desktopEnv string) {
	c_desktop_env := C.CString(desktopEnv)
	defer C.free(unsafe.Pointer(c_desktop_env))

	C.g_desktop_app_info_set_desktop_env(c_desktop_env)

	return
}

// GetCategories is a wrapper around the C function g_desktop_app_info_get_categories.
func (recv *DesktopAppInfo) GetCategories() string {
	retC := C.g_desktop_app_info_get_categories((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetGenericName is a wrapper around the C function g_desktop_app_info_get_generic_name.
func (recv *DesktopAppInfo) GetGenericName() string {
	retC := C.g_desktop_app_info_get_generic_name((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIsHidden is a wrapper around the C function g_desktop_app_info_get_is_hidden.
func (recv *DesktopAppInfo) GetIsHidden() bool {
	retC := C.g_desktop_app_info_get_is_hidden((*C.GDesktopAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_desktop_app_info_launch_uris_as_manager : unsupported parameter user_setup : no type generator for GLib.SpawnChildSetupFunc (GSpawnChildSetupFunc) for param user_setup

// AppInfo returns the AppInfo interface implemented by DesktopAppInfo
func (recv *DesktopAppInfo) AppInfo() *AppInfo {
	return AppInfoNewFromC(recv.ToC())
}

// Emblem is a wrapper around the C record GEmblem.
type Emblem struct {
	native *C.GEmblem
}

func EmblemNewFromC(u unsafe.Pointer) *Emblem {
	c := (*C.GEmblem)(u)
	if c == nil {
		return nil
	}

	g := &Emblem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Emblem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Emblem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Emblem with another Emblem, and returns true if they represent the same GObject.
func (recv *Emblem) Equals(other *Emblem) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Emblem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Emblem.
// Exercise care, as this is a potentially dangerous function if the Object is not a Emblem.
func CastToEmblem(object *gobject.Object) *Emblem {
	return EmblemNewFromC(object.ToC())
}

// Icon returns the Icon interface implemented by Emblem
func (recv *Emblem) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// EmblemedIcon is a wrapper around the C record GEmblemedIcon.
type EmblemedIcon struct {
	native *C.GEmblemedIcon
	// parent_instance : record
	// Private : priv
}

func EmblemedIconNewFromC(u unsafe.Pointer) *EmblemedIcon {
	c := (*C.GEmblemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EmblemedIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EmblemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EmblemedIcon with another EmblemedIcon, and returns true if they represent the same GObject.
func (recv *EmblemedIcon) Equals(other *EmblemedIcon) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *EmblemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to EmblemedIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a EmblemedIcon.
func CastToEmblemedIcon(object *gobject.Object) *EmblemedIcon {
	return EmblemedIconNewFromC(object.ToC())
}

// Icon returns the Icon interface implemented by EmblemedIcon
func (recv *EmblemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// FileEnumerator is a wrapper around the C record GFileEnumerator.
type FileEnumerator struct {
	native *C.GFileEnumerator
	// parent_instance : record
	// Private : priv
}

func FileEnumeratorNewFromC(u unsafe.Pointer) *FileEnumerator {
	c := (*C.GFileEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileEnumerator with another FileEnumerator, and returns true if they represent the same GObject.
func (recv *FileEnumerator) Equals(other *FileEnumerator) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileEnumerator.
func CastToFileEnumerator(object *gobject.Object) *FileEnumerator {
	return FileEnumeratorNewFromC(object.ToC())
}

// Close is a wrapper around the C function g_file_enumerator_close.
func (recv *FileEnumerator) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_close((*C.GFileEnumerator)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_enumerator_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_file_enumerator_close_finish.
func (recv *FileEnumerator) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_close_finish((*C.GFileEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasPending is a wrapper around the C function g_file_enumerator_has_pending.
func (recv *FileEnumerator) HasPending() bool {
	retC := C.g_file_enumerator_has_pending((*C.GFileEnumerator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsClosed is a wrapper around the C function g_file_enumerator_is_closed.
func (recv *FileEnumerator) IsClosed() bool {
	retC := C.g_file_enumerator_is_closed((*C.GFileEnumerator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// NextFile is a wrapper around the C function g_file_enumerator_next_file.
func (recv *FileEnumerator) NextFile(cancellable *Cancellable) (*FileInfo, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_next_file((*C.GFileEnumerator)(recv.native), c_cancellable, &cThrowableError)
	var retGo (*FileInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileInfoNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_enumerator_next_files_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// NextFilesFinish is a wrapper around the C function g_file_enumerator_next_files_finish.
func (recv *FileEnumerator) NextFilesFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_next_files_finish((*C.GFileEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetPending is a wrapper around the C function g_file_enumerator_set_pending.
func (recv *FileEnumerator) SetPending(pending bool) {
	c_pending :=
		boolToGboolean(pending)

	C.g_file_enumerator_set_pending((*C.GFileEnumerator)(recv.native), c_pending)

	return
}

// FileIOStream is a wrapper around the C record GFileIOStream.
type FileIOStream struct {
	native *C.GFileIOStream
	// parent_instance : record
	// Private : priv
}

func FileIOStreamNewFromC(u unsafe.Pointer) *FileIOStream {
	c := (*C.GFileIOStream)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileIOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIOStream with another FileIOStream, and returns true if they represent the same GObject.
func (recv *FileIOStream) Equals(other *FileIOStream) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *FileIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to FileIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileIOStream.
func CastToFileIOStream(object *gobject.Object) *FileIOStream {
	return FileIOStreamNewFromC(object.ToC())
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FileIcon is a wrapper around the C record GFileIcon.
type FileIcon struct {
	native *C.GFileIcon
}

func FileIconNewFromC(u unsafe.Pointer) *FileIcon {
	c := (*C.GFileIcon)(u)
	if c == nil {
		return nil
	}

	g := &FileIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIcon with another FileIcon, and returns true if they represent the same GObject.
func (recv *FileIcon) Equals(other *FileIcon) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileIcon.
func CastToFileIcon(object *gobject.Object) *FileIcon {
	return FileIconNewFromC(object.ToC())
}

// FileIconNew is a wrapper around the C function g_file_icon_new.
func FileIconNew(file *File) *FileIcon {
	c_file := (*C.GFile)(file.ToC())

	retC := C.g_file_icon_new(c_file)
	retGo := FileIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetFile is a wrapper around the C function g_file_icon_get_file.
func (recv *FileIcon) GetFile() *File {
	retC := C.g_file_icon_get_file((*C.GFileIcon)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Icon returns the Icon interface implemented by FileIcon
func (recv *FileIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by FileIcon
func (recv *FileIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// FileInfo is a wrapper around the C record GFileInfo.
type FileInfo struct {
	native *C.GFileInfo
}

func FileInfoNewFromC(u unsafe.Pointer) *FileInfo {
	c := (*C.GFileInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileInfo with another FileInfo, and returns true if they represent the same GObject.
func (recv *FileInfo) Equals(other *FileInfo) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileInfo.
func CastToFileInfo(object *gobject.Object) *FileInfo {
	return FileInfoNewFromC(object.ToC())
}

// FileInfoNew is a wrapper around the C function g_file_info_new.
func FileInfoNew() *FileInfo {
	retC := C.g_file_info_new()
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ClearStatus is a wrapper around the C function g_file_info_clear_status.
func (recv *FileInfo) ClearStatus() {
	C.g_file_info_clear_status((*C.GFileInfo)(recv.native))

	return
}

// CopyInto is a wrapper around the C function g_file_info_copy_into.
func (recv *FileInfo) CopyInto(destInfo *FileInfo) {
	c_dest_info := (*C.GFileInfo)(C.NULL)
	if destInfo != nil {
		c_dest_info = (*C.GFileInfo)(destInfo.ToC())
	}

	C.g_file_info_copy_into((*C.GFileInfo)(recv.native), c_dest_info)

	return
}

// Dup is a wrapper around the C function g_file_info_dup.
func (recv *FileInfo) Dup() *FileInfo {
	retC := C.g_file_info_dup((*C.GFileInfo)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttributeAsString is a wrapper around the C function g_file_info_get_attribute_as_string.
func (recv *FileInfo) GetAttributeAsString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_as_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetAttributeBoolean is a wrapper around the C function g_file_info_get_attribute_boolean.
func (recv *FileInfo) GetAttributeBoolean(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_boolean((*C.GFileInfo)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// GetAttributeByteString is a wrapper around the C function g_file_info_get_attribute_byte_string.
func (recv *FileInfo) GetAttributeByteString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_byte_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_file_info_get_attribute_data : unsupported parameter type : GFileAttributeType* with indirection level of 1

// GetAttributeInt32 is a wrapper around the C function g_file_info_get_attribute_int32.
func (recv *FileInfo) GetAttributeInt32(attribute string) int32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int32)(retC)

	return retGo
}

// GetAttributeInt64 is a wrapper around the C function g_file_info_get_attribute_int64.
func (recv *FileInfo) GetAttributeInt64(attribute string) int64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int64)(retC)

	return retGo
}

// GetAttributeObject is a wrapper around the C function g_file_info_get_attribute_object.
func (recv *FileInfo) GetAttributeObject(attribute string) *gobject.Object {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_object((*C.GFileInfo)(recv.native), c_attribute)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttributeStatus is a wrapper around the C function g_file_info_get_attribute_status.
func (recv *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_status((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeStatus)(retC)

	return retGo
}

// GetAttributeString is a wrapper around the C function g_file_info_get_attribute_string.
func (recv *FileInfo) GetAttributeString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// GetAttributeType is a wrapper around the C function g_file_info_get_attribute_type.
func (recv *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_type((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeType)(retC)

	return retGo
}

// GetAttributeUint32 is a wrapper around the C function g_file_info_get_attribute_uint32.
func (recv *FileInfo) GetAttributeUint32(attribute string) uint32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint32)(retC)

	return retGo
}

// GetAttributeUint64 is a wrapper around the C function g_file_info_get_attribute_uint64.
func (recv *FileInfo) GetAttributeUint64(attribute string) uint64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint64)(retC)

	return retGo
}

// GetContentType is a wrapper around the C function g_file_info_get_content_type.
func (recv *FileInfo) GetContentType() string {
	retC := C.g_file_info_get_content_type((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function g_file_info_get_display_name.
func (recv *FileInfo) GetDisplayName() string {
	retC := C.g_file_info_get_display_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEditName is a wrapper around the C function g_file_info_get_edit_name.
func (recv *FileInfo) GetEditName() string {
	retC := C.g_file_info_get_edit_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEtag is a wrapper around the C function g_file_info_get_etag.
func (recv *FileInfo) GetEtag() string {
	retC := C.g_file_info_get_etag((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFileType is a wrapper around the C function g_file_info_get_file_type.
func (recv *FileInfo) GetFileType() FileType {
	retC := C.g_file_info_get_file_type((*C.GFileInfo)(recv.native))
	retGo := (FileType)(retC)

	return retGo
}

// GetIcon is a wrapper around the C function g_file_info_get_icon.
func (recv *FileInfo) GetIcon() *Icon {
	retC := C.g_file_info_get_icon((*C.GFileInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIsBackup is a wrapper around the C function g_file_info_get_is_backup.
func (recv *FileInfo) GetIsBackup() bool {
	retC := C.g_file_info_get_is_backup((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsHidden is a wrapper around the C function g_file_info_get_is_hidden.
func (recv *FileInfo) GetIsHidden() bool {
	retC := C.g_file_info_get_is_hidden((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsSymlink is a wrapper around the C function g_file_info_get_is_symlink.
func (recv *FileInfo) GetIsSymlink() bool {
	retC := C.g_file_info_get_is_symlink((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModificationTime is a wrapper around the C function g_file_info_get_modification_time.
func (recv *FileInfo) GetModificationTime() *glib.TimeVal {
	var c_result C.GTimeVal

	C.g_file_info_get_modification_time((*C.GFileInfo)(recv.native), &c_result)

	result := glib.TimeValNewFromC(unsafe.Pointer(&c_result))

	return result
}

// GetName is a wrapper around the C function g_file_info_get_name.
func (recv *FileInfo) GetName() string {
	retC := C.g_file_info_get_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSize is a wrapper around the C function g_file_info_get_size.
func (recv *FileInfo) GetSize() int64 {
	retC := C.g_file_info_get_size((*C.GFileInfo)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetSortOrder is a wrapper around the C function g_file_info_get_sort_order.
func (recv *FileInfo) GetSortOrder() int32 {
	retC := C.g_file_info_get_sort_order((*C.GFileInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSymlinkTarget is a wrapper around the C function g_file_info_get_symlink_target.
func (recv *FileInfo) GetSymlinkTarget() string {
	retC := C.g_file_info_get_symlink_target((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// HasAttribute is a wrapper around the C function g_file_info_has_attribute.
func (recv *FileInfo) HasAttribute(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_has_attribute((*C.GFileInfo)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// ListAttributes is a wrapper around the C function g_file_info_list_attributes.
func (recv *FileInfo) ListAttributes(nameSpace string) []string {
	c_name_space := C.CString(nameSpace)
	defer C.free(unsafe.Pointer(c_name_space))

	retC := C.g_file_info_list_attributes((*C.GFileInfo)(recv.native), c_name_space)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// RemoveAttribute is a wrapper around the C function g_file_info_remove_attribute.
func (recv *FileInfo) RemoveAttribute(attribute string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	C.g_file_info_remove_attribute((*C.GFileInfo)(recv.native), c_attribute)

	return
}

// Unsupported : g_file_info_set_attribute : unsupported parameter value_p : no type generator for gpointer (gpointer) for param value_p

// SetAttributeBoolean is a wrapper around the C function g_file_info_set_attribute_boolean.
func (recv *FileInfo) SetAttributeBoolean(attribute string, attrValue bool) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value :=
		boolToGboolean(attrValue)

	C.g_file_info_set_attribute_boolean((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeByteString is a wrapper around the C function g_file_info_set_attribute_byte_string.
func (recv *FileInfo) SetAttributeByteString(attribute string, attrValue string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := C.CString(attrValue)
	defer C.free(unsafe.Pointer(c_attr_value))

	C.g_file_info_set_attribute_byte_string((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeInt32 is a wrapper around the C function g_file_info_set_attribute_int32.
func (recv *FileInfo) SetAttributeInt32(attribute string, attrValue int32) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.gint32)(attrValue)

	C.g_file_info_set_attribute_int32((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeInt64 is a wrapper around the C function g_file_info_set_attribute_int64.
func (recv *FileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.gint64)(attrValue)

	C.g_file_info_set_attribute_int64((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeMask is a wrapper around the C function g_file_info_set_attribute_mask.
func (recv *FileInfo) SetAttributeMask(mask *FileAttributeMatcher) {
	c_mask := (*C.GFileAttributeMatcher)(C.NULL)
	if mask != nil {
		c_mask = (*C.GFileAttributeMatcher)(mask.ToC())
	}

	C.g_file_info_set_attribute_mask((*C.GFileInfo)(recv.native), c_mask)

	return
}

// SetAttributeObject is a wrapper around the C function g_file_info_set_attribute_object.
func (recv *FileInfo) SetAttributeObject(attribute string, attrValue *gobject.Object) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (*C.GObject)(C.NULL)
	if attrValue != nil {
		c_attr_value = (*C.GObject)(attrValue.ToC())
	}

	C.g_file_info_set_attribute_object((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeString is a wrapper around the C function g_file_info_set_attribute_string.
func (recv *FileInfo) SetAttributeString(attribute string, attrValue string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := C.CString(attrValue)
	defer C.free(unsafe.Pointer(c_attr_value))

	C.g_file_info_set_attribute_string((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeStringv is a wrapper around the C function g_file_info_set_attribute_stringv.
func (recv *FileInfo) SetAttributeStringv(attribute string, attrValue []string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value_array := make([]*C.gchar, len(attrValue)+1, len(attrValue)+1)
	for i, item := range attrValue {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_attr_value_array[i] = c
	}
	c_attr_value_array[len(attrValue)] = nil
	c_attr_value_arrayPtr := &c_attr_value_array[0]
	c_attr_value := (**C.char)(unsafe.Pointer(c_attr_value_arrayPtr))

	C.g_file_info_set_attribute_stringv((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeUint32 is a wrapper around the C function g_file_info_set_attribute_uint32.
func (recv *FileInfo) SetAttributeUint32(attribute string, attrValue uint32) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.guint32)(attrValue)

	C.g_file_info_set_attribute_uint32((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetAttributeUint64 is a wrapper around the C function g_file_info_set_attribute_uint64.
func (recv *FileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.guint64)(attrValue)

	C.g_file_info_set_attribute_uint64((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// SetContentType is a wrapper around the C function g_file_info_set_content_type.
func (recv *FileInfo) SetContentType(contentType string) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	C.g_file_info_set_content_type((*C.GFileInfo)(recv.native), c_content_type)

	return
}

// SetDisplayName is a wrapper around the C function g_file_info_set_display_name.
func (recv *FileInfo) SetDisplayName(displayName string) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	C.g_file_info_set_display_name((*C.GFileInfo)(recv.native), c_display_name)

	return
}

// SetEditName is a wrapper around the C function g_file_info_set_edit_name.
func (recv *FileInfo) SetEditName(editName string) {
	c_edit_name := C.CString(editName)
	defer C.free(unsafe.Pointer(c_edit_name))

	C.g_file_info_set_edit_name((*C.GFileInfo)(recv.native), c_edit_name)

	return
}

// SetFileType is a wrapper around the C function g_file_info_set_file_type.
func (recv *FileInfo) SetFileType(type_ FileType) {
	c_type := (C.GFileType)(type_)

	C.g_file_info_set_file_type((*C.GFileInfo)(recv.native), c_type)

	return
}

// SetIcon is a wrapper around the C function g_file_info_set_icon.
func (recv *FileInfo) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_file_info_set_icon((*C.GFileInfo)(recv.native), c_icon)

	return
}

// SetIsHidden is a wrapper around the C function g_file_info_set_is_hidden.
func (recv *FileInfo) SetIsHidden(isHidden bool) {
	c_is_hidden :=
		boolToGboolean(isHidden)

	C.g_file_info_set_is_hidden((*C.GFileInfo)(recv.native), c_is_hidden)

	return
}

// SetIsSymlink is a wrapper around the C function g_file_info_set_is_symlink.
func (recv *FileInfo) SetIsSymlink(isSymlink bool) {
	c_is_symlink :=
		boolToGboolean(isSymlink)

	C.g_file_info_set_is_symlink((*C.GFileInfo)(recv.native), c_is_symlink)

	return
}

// SetModificationTime is a wrapper around the C function g_file_info_set_modification_time.
func (recv *FileInfo) SetModificationTime(mtime *glib.TimeVal) {
	c_mtime := (*C.GTimeVal)(C.NULL)
	if mtime != nil {
		c_mtime = (*C.GTimeVal)(mtime.ToC())
	}

	C.g_file_info_set_modification_time((*C.GFileInfo)(recv.native), c_mtime)

	return
}

// SetName is a wrapper around the C function g_file_info_set_name.
func (recv *FileInfo) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_file_info_set_name((*C.GFileInfo)(recv.native), c_name)

	return
}

// SetSize is a wrapper around the C function g_file_info_set_size.
func (recv *FileInfo) SetSize(size int64) {
	c_size := (C.goffset)(size)

	C.g_file_info_set_size((*C.GFileInfo)(recv.native), c_size)

	return
}

// SetSortOrder is a wrapper around the C function g_file_info_set_sort_order.
func (recv *FileInfo) SetSortOrder(sortOrder int32) {
	c_sort_order := (C.gint32)(sortOrder)

	C.g_file_info_set_sort_order((*C.GFileInfo)(recv.native), c_sort_order)

	return
}

// SetSymlinkTarget is a wrapper around the C function g_file_info_set_symlink_target.
func (recv *FileInfo) SetSymlinkTarget(symlinkTarget string) {
	c_symlink_target := C.CString(symlinkTarget)
	defer C.free(unsafe.Pointer(c_symlink_target))

	C.g_file_info_set_symlink_target((*C.GFileInfo)(recv.native), c_symlink_target)

	return
}

// UnsetAttributeMask is a wrapper around the C function g_file_info_unset_attribute_mask.
func (recv *FileInfo) UnsetAttributeMask() {
	C.g_file_info_unset_attribute_mask((*C.GFileInfo)(recv.native))

	return
}

// FileInputStream is a wrapper around the C record GFileInputStream.
type FileInputStream struct {
	native *C.GFileInputStream
	// parent_instance : record
	// Private : priv
}

func FileInputStreamNewFromC(u unsafe.Pointer) *FileInputStream {
	c := (*C.GFileInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileInputStream with another FileInputStream, and returns true if they represent the same GObject.
func (recv *FileInputStream) Equals(other *FileInputStream) bool {
	return other.ToC() == recv.ToC()
}

// InputStream upcasts to *InputStream
func (recv *FileInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitrary Object to FileInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileInputStream.
func CastToFileInputStream(object *gobject.Object) *FileInputStream {
	return FileInputStreamNewFromC(object.ToC())
}

// QueryInfo is a wrapper around the C function g_file_input_stream_query_info.
func (recv *FileInputStream) QueryInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_input_stream_query_info((*C.GFileInputStream)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_input_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryInfoFinish is a wrapper around the C function g_file_input_stream_query_info_finish.
func (recv *FileInputStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_input_stream_query_info_finish((*C.GFileInputStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Seekable returns the Seekable interface implemented by FileInputStream
func (recv *FileInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FileMonitor is a wrapper around the C record GFileMonitor.
type FileMonitor struct {
	native *C.GFileMonitor
	// parent_instance : record
	// Private : priv
}

func FileMonitorNewFromC(u unsafe.Pointer) *FileMonitor {
	c := (*C.GFileMonitor)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileMonitor with another FileMonitor, and returns true if they represent the same GObject.
func (recv *FileMonitor) Equals(other *FileMonitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileMonitor.
func CastToFileMonitor(object *gobject.Object) *FileMonitor {
	return FileMonitorNewFromC(object.ToC())
}

type signalFileMonitorChangedDetail struct {
	callback  FileMonitorSignalChangedCallback
	handlerID C.gulong
}

var signalFileMonitorChangedId int
var signalFileMonitorChangedMap = make(map[int]signalFileMonitorChangedDetail)
var signalFileMonitorChangedLock sync.RWMutex

// FileMonitorSignalChangedCallback is a callback function for a 'changed' signal emitted from a FileMonitor.
type FileMonitorSignalChangedCallback func(file *File, otherFile *File, eventType FileMonitorEvent)

/*
ConnectChanged connects the callback to the 'changed' signal for the FileMonitor.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *FileMonitor) ConnectChanged(callback FileMonitorSignalChangedCallback) int {
	signalFileMonitorChangedLock.Lock()
	defer signalFileMonitorChangedLock.Unlock()

	signalFileMonitorChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileMonitor_signal_connect_changed(instance, C.gpointer(uintptr(signalFileMonitorChangedId)))

	detail := signalFileMonitorChangedDetail{callback, handlerID}
	signalFileMonitorChangedMap[signalFileMonitorChangedId] = detail

	return signalFileMonitorChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the FileMonitor.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *FileMonitor) DisconnectChanged(connectionID int) {
	signalFileMonitorChangedLock.Lock()
	defer signalFileMonitorChangedLock.Unlock()

	detail, exists := signalFileMonitorChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileMonitorChangedMap, connectionID)
}

//export filemonitor_changedHandler
func filemonitor_changedHandler(_ *C.GObject, c_file *C.GFile, c_other_file *C.GFile, c_event_type C.GFileMonitorEvent, data C.gpointer) {
	signalFileMonitorChangedLock.RLock()
	defer signalFileMonitorChangedLock.RUnlock()

	file := FileNewFromC(unsafe.Pointer(c_file))

	otherFile := FileNewFromC(unsafe.Pointer(c_other_file))

	eventType := FileMonitorEvent(c_event_type)

	index := int(uintptr(data))
	callback := signalFileMonitorChangedMap[index].callback
	callback(file, otherFile, eventType)
}

// Cancel is a wrapper around the C function g_file_monitor_cancel.
func (recv *FileMonitor) Cancel() bool {
	retC := C.g_file_monitor_cancel((*C.GFileMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// EmitEvent is a wrapper around the C function g_file_monitor_emit_event.
func (recv *FileMonitor) EmitEvent(child *File, otherFile *File, eventType FileMonitorEvent) {
	c_child := (*C.GFile)(child.ToC())

	c_other_file := (*C.GFile)(otherFile.ToC())

	c_event_type := (C.GFileMonitorEvent)(eventType)

	C.g_file_monitor_emit_event((*C.GFileMonitor)(recv.native), c_child, c_other_file, c_event_type)

	return
}

// IsCancelled is a wrapper around the C function g_file_monitor_is_cancelled.
func (recv *FileMonitor) IsCancelled() bool {
	retC := C.g_file_monitor_is_cancelled((*C.GFileMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetRateLimit is a wrapper around the C function g_file_monitor_set_rate_limit.
func (recv *FileMonitor) SetRateLimit(limitMsecs int32) {
	c_limit_msecs := (C.gint)(limitMsecs)

	C.g_file_monitor_set_rate_limit((*C.GFileMonitor)(recv.native), c_limit_msecs)

	return
}

// FileOutputStream is a wrapper around the C record GFileOutputStream.
type FileOutputStream struct {
	native *C.GFileOutputStream
	// parent_instance : record
	// Private : priv
}

func FileOutputStreamNewFromC(u unsafe.Pointer) *FileOutputStream {
	c := (*C.GFileOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileOutputStream with another FileOutputStream, and returns true if they represent the same GObject.
func (recv *FileOutputStream) Equals(other *FileOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// OutputStream upcasts to *OutputStream
func (recv *FileOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to FileOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileOutputStream.
func CastToFileOutputStream(object *gobject.Object) *FileOutputStream {
	return FileOutputStreamNewFromC(object.ToC())
}

// GetEtag is a wrapper around the C function g_file_output_stream_get_etag.
func (recv *FileOutputStream) GetEtag() string {
	retC := C.g_file_output_stream_get_etag((*C.GFileOutputStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// QueryInfo is a wrapper around the C function g_file_output_stream_query_info.
func (recv *FileOutputStream) QueryInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_output_stream_query_info((*C.GFileOutputStream)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_output_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryInfoFinish is a wrapper around the C function g_file_output_stream_query_info_finish.
func (recv *FileOutputStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_output_stream_query_info_finish((*C.GFileOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FilenameCompleter is a wrapper around the C record GFilenameCompleter.
type FilenameCompleter struct {
	native *C.GFilenameCompleter
}

func FilenameCompleterNewFromC(u unsafe.Pointer) *FilenameCompleter {
	c := (*C.GFilenameCompleter)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilenameCompleter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilenameCompleter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilenameCompleter with another FilenameCompleter, and returns true if they represent the same GObject.
func (recv *FilenameCompleter) Equals(other *FilenameCompleter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FilenameCompleter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FilenameCompleter.
// Exercise care, as this is a potentially dangerous function if the Object is not a FilenameCompleter.
func CastToFilenameCompleter(object *gobject.Object) *FilenameCompleter {
	return FilenameCompleterNewFromC(object.ToC())
}

type signalFilenameCompleterGotCompletionDataDetail struct {
	callback  FilenameCompleterSignalGotCompletionDataCallback
	handlerID C.gulong
}

var signalFilenameCompleterGotCompletionDataId int
var signalFilenameCompleterGotCompletionDataMap = make(map[int]signalFilenameCompleterGotCompletionDataDetail)
var signalFilenameCompleterGotCompletionDataLock sync.RWMutex

// FilenameCompleterSignalGotCompletionDataCallback is a callback function for a 'got-completion-data' signal emitted from a FilenameCompleter.
type FilenameCompleterSignalGotCompletionDataCallback func()

/*
ConnectGotCompletionData connects the callback to the 'got-completion-data' signal for the FilenameCompleter.

The returned value represents the connection, and may be passed to DisconnectGotCompletionData to remove it.
*/
func (recv *FilenameCompleter) ConnectGotCompletionData(callback FilenameCompleterSignalGotCompletionDataCallback) int {
	signalFilenameCompleterGotCompletionDataLock.Lock()
	defer signalFilenameCompleterGotCompletionDataLock.Unlock()

	signalFilenameCompleterGotCompletionDataId++
	instance := C.gpointer(recv.native)
	handlerID := C.FilenameCompleter_signal_connect_got_completion_data(instance, C.gpointer(uintptr(signalFilenameCompleterGotCompletionDataId)))

	detail := signalFilenameCompleterGotCompletionDataDetail{callback, handlerID}
	signalFilenameCompleterGotCompletionDataMap[signalFilenameCompleterGotCompletionDataId] = detail

	return signalFilenameCompleterGotCompletionDataId
}

/*
DisconnectGotCompletionData disconnects a callback from the 'got-completion-data' signal for the FilenameCompleter.

The connectionID should be a value returned from a call to ConnectGotCompletionData.
*/
func (recv *FilenameCompleter) DisconnectGotCompletionData(connectionID int) {
	signalFilenameCompleterGotCompletionDataLock.Lock()
	defer signalFilenameCompleterGotCompletionDataLock.Unlock()

	detail, exists := signalFilenameCompleterGotCompletionDataMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFilenameCompleterGotCompletionDataMap, connectionID)
}

//export filenamecompleter_gotCompletionDataHandler
func filenamecompleter_gotCompletionDataHandler(_ *C.GObject, data C.gpointer) {
	signalFilenameCompleterGotCompletionDataLock.RLock()
	defer signalFilenameCompleterGotCompletionDataLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFilenameCompleterGotCompletionDataMap[index].callback
	callback()
}

// FilenameCompleterNew is a wrapper around the C function g_filename_completer_new.
func FilenameCompleterNew() *FilenameCompleter {
	retC := C.g_filename_completer_new()
	retGo := FilenameCompleterNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetCompletionSuffix is a wrapper around the C function g_filename_completer_get_completion_suffix.
func (recv *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	c_initial_text := C.CString(initialText)
	defer C.free(unsafe.Pointer(c_initial_text))

	retC := C.g_filename_completer_get_completion_suffix((*C.GFilenameCompleter)(recv.native), c_initial_text)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCompletions is a wrapper around the C function g_filename_completer_get_completions.
func (recv *FilenameCompleter) GetCompletions(initialText string) []string {
	c_initial_text := C.CString(initialText)
	defer C.free(unsafe.Pointer(c_initial_text))

	retC := C.g_filename_completer_get_completions((*C.GFilenameCompleter)(recv.native), c_initial_text)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// SetDirsOnly is a wrapper around the C function g_filename_completer_set_dirs_only.
func (recv *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	c_dirs_only :=
		boolToGboolean(dirsOnly)

	C.g_filename_completer_set_dirs_only((*C.GFilenameCompleter)(recv.native), c_dirs_only)

	return
}

// FilterInputStream is a wrapper around the C record GFilterInputStream.
type FilterInputStream struct {
	native *C.GFilterInputStream
	// parent_instance : record
	// base_stream : record
}

func FilterInputStreamNewFromC(u unsafe.Pointer) *FilterInputStream {
	c := (*C.GFilterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilterInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilterInputStream with another FilterInputStream, and returns true if they represent the same GObject.
func (recv *FilterInputStream) Equals(other *FilterInputStream) bool {
	return other.ToC() == recv.ToC()
}

// InputStream upcasts to *InputStream
func (recv *FilterInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FilterInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitrary Object to FilterInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FilterInputStream.
func CastToFilterInputStream(object *gobject.Object) *FilterInputStream {
	return FilterInputStreamNewFromC(object.ToC())
}

// GetBaseStream is a wrapper around the C function g_filter_input_stream_get_base_stream.
func (recv *FilterInputStream) GetBaseStream() *InputStream {
	retC := C.g_filter_input_stream_get_base_stream((*C.GFilterInputStream)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCloseBaseStream is a wrapper around the C function g_filter_input_stream_get_close_base_stream.
func (recv *FilterInputStream) GetCloseBaseStream() bool {
	retC := C.g_filter_input_stream_get_close_base_stream((*C.GFilterInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCloseBaseStream is a wrapper around the C function g_filter_input_stream_set_close_base_stream.
func (recv *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	c_close_base :=
		boolToGboolean(closeBase)

	C.g_filter_input_stream_set_close_base_stream((*C.GFilterInputStream)(recv.native), c_close_base)

	return
}

// FilterOutputStream is a wrapper around the C record GFilterOutputStream.
type FilterOutputStream struct {
	native *C.GFilterOutputStream
	// parent_instance : record
	// base_stream : record
}

func FilterOutputStreamNewFromC(u unsafe.Pointer) *FilterOutputStream {
	c := (*C.GFilterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilterOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilterOutputStream with another FilterOutputStream, and returns true if they represent the same GObject.
func (recv *FilterOutputStream) Equals(other *FilterOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// OutputStream upcasts to *OutputStream
func (recv *FilterOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FilterOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to FilterOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FilterOutputStream.
func CastToFilterOutputStream(object *gobject.Object) *FilterOutputStream {
	return FilterOutputStreamNewFromC(object.ToC())
}

// GetBaseStream is a wrapper around the C function g_filter_output_stream_get_base_stream.
func (recv *FilterOutputStream) GetBaseStream() *OutputStream {
	retC := C.g_filter_output_stream_get_base_stream((*C.GFilterOutputStream)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCloseBaseStream is a wrapper around the C function g_filter_output_stream_get_close_base_stream.
func (recv *FilterOutputStream) GetCloseBaseStream() bool {
	retC := C.g_filter_output_stream_get_close_base_stream((*C.GFilterOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCloseBaseStream is a wrapper around the C function g_filter_output_stream_set_close_base_stream.
func (recv *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	c_close_base :=
		boolToGboolean(closeBase)

	C.g_filter_output_stream_set_close_base_stream((*C.GFilterOutputStream)(recv.native), c_close_base)

	return
}

// IOModule is a wrapper around the C record GIOModule.
type IOModule struct {
	native *C.GIOModule
}

func IOModuleNewFromC(u unsafe.Pointer) *IOModule {
	c := (*C.GIOModule)(u)
	if c == nil {
		return nil
	}

	g := &IOModule{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IOModule) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IOModule) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOModule with another IOModule, and returns true if they represent the same GObject.
func (recv *IOModule) Equals(other *IOModule) bool {
	return other.ToC() == recv.ToC()
}

// CastToWidget down casts any arbitrary Object to IOModule.
// Exercise care, as this is a potentially dangerous function if the Object is not a IOModule.
func CastToIOModule(object *gobject.Object) *IOModule {
	return IOModuleNewFromC(object.ToC())
}

// IOModuleNew is a wrapper around the C function g_io_module_new.
func IOModuleNew(filename string) *IOModule {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_io_module_new(c_filename)
	retGo := IOModuleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : g_io_module_load

// Blacklisted : g_io_module_unload

// TypePlugin returns the TypePlugin interface implemented by IOModule
func (recv *IOModule) TypePlugin() *gobject.TypePlugin {
	return gobject.TypePluginNewFromC(recv.ToC())
}

// IOStream is a wrapper around the C record GIOStream.
type IOStream struct {
	native *C.GIOStream
	// parent_instance : record
	// Private : priv
}

func IOStreamNewFromC(u unsafe.Pointer) *IOStream {
	c := (*C.GIOStream)(u)
	if c == nil {
		return nil
	}

	g := &IOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOStream with another IOStream, and returns true if they represent the same GObject.
func (recv *IOStream) Equals(other *IOStream) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *IOStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to IOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a IOStream.
func CastToIOStream(object *gobject.Object) *IOStream {
	return IOStreamNewFromC(object.ToC())
}

// InetAddress is a wrapper around the C record GInetAddress.
type InetAddress struct {
	native *C.GInetAddress
	// parent_instance : record
	// Private : priv
}

func InetAddressNewFromC(u unsafe.Pointer) *InetAddress {
	c := (*C.GInetAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddress with another InetAddress, and returns true if they represent the same GObject.
func (recv *InetAddress) Equals(other *InetAddress) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *InetAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to InetAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetAddress.
func CastToInetAddress(object *gobject.Object) *InetAddress {
	return InetAddressNewFromC(object.ToC())
}

// InetSocketAddress is a wrapper around the C record GInetSocketAddress.
type InetSocketAddress struct {
	native *C.GInetSocketAddress
	// parent_instance : record
	// Private : priv
}

func InetSocketAddressNewFromC(u unsafe.Pointer) *InetSocketAddress {
	c := (*C.GInetSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetSocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetSocketAddress with another InetSocketAddress, and returns true if they represent the same GObject.
func (recv *InetSocketAddress) Equals(other *InetSocketAddress) bool {
	return other.ToC() == recv.ToC()
}

// SocketAddress upcasts to *SocketAddress
func (recv *InetSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *InetSocketAddress) Object() *gobject.Object {
	return recv.SocketAddress().Object()
}

// CastToWidget down casts any arbitrary Object to InetSocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetSocketAddress.
func CastToInetSocketAddress(object *gobject.Object) *InetSocketAddress {
	return InetSocketAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by InetSocketAddress
func (recv *InetSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// InputStream is a wrapper around the C record GInputStream.
type InputStream struct {
	native *C.GInputStream
	// parent_instance : record
	// Private : priv
}

func InputStreamNewFromC(u unsafe.Pointer) *InputStream {
	c := (*C.GInputStream)(u)
	if c == nil {
		return nil
	}

	g := &InputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InputStream with another InputStream, and returns true if they represent the same GObject.
func (recv *InputStream) Equals(other *InputStream) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *InputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to InputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a InputStream.
func CastToInputStream(object *gobject.Object) *InputStream {
	return InputStreamNewFromC(object.ToC())
}

// ClearPending is a wrapper around the C function g_input_stream_clear_pending.
func (recv *InputStream) ClearPending() {
	C.g_input_stream_clear_pending((*C.GInputStream)(recv.native))

	return
}

// Close is a wrapper around the C function g_input_stream_close.
func (recv *InputStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_close((*C.GInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_input_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_input_stream_close_finish.
func (recv *InputStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_close_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasPending is a wrapper around the C function g_input_stream_has_pending.
func (recv *InputStream) HasPending() bool {
	retC := C.g_input_stream_has_pending((*C.GInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsClosed is a wrapper around the C function g_input_stream_is_closed.
func (recv *InputStream) IsClosed() bool {
	retC := C.g_input_stream_is_closed((*C.GInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Read is a wrapper around the C function g_input_stream_read.
func (recv *InputStream) Read(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (unsafe.Pointer(c_buffer_arrayPtr))

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read((*C.GInputStream)(recv.native), c_buffer, c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReadAll is a wrapper around the C function g_input_stream_read_all.
func (recv *InputStream) ReadAll(buffer []uint8, cancellable *Cancellable) (bool, uint64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (unsafe.Pointer(c_buffer_arrayPtr))

	c_count := (C.gsize)(len(buffer))

	var c_bytes_read C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_all((*C.GInputStream)(recv.native), c_buffer, c_count, &c_bytes_read, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	return retGo, bytesRead, goError
}

// Unsupported : g_input_stream_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadFinish is a wrapper around the C function g_input_stream_read_finish.
func (recv *InputStream) ReadFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetPending is a wrapper around the C function g_input_stream_set_pending.
func (recv *InputStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_input_stream_set_pending((*C.GInputStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Skip is a wrapper around the C function g_input_stream_skip.
func (recv *InputStream) Skip(count uint64, cancellable *Cancellable) (int64, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_skip((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_input_stream_skip_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SkipFinish is a wrapper around the C function g_input_stream_skip_finish.
func (recv *InputStream) SkipFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_skip_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ListStore is a wrapper around the C record GListStore.
type ListStore struct {
	native *C.GListStore
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListStore) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ListStore with another ListStore, and returns true if they represent the same GObject.
func (recv *ListStore) Equals(other *ListStore) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ListStore) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ListStore.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListStore.
func CastToListStore(object *gobject.Object) *ListStore {
	return ListStoreNewFromC(object.ToC())
}

// ListModel returns the ListModel interface implemented by ListStore
func (recv *ListStore) ListModel() *ListModel {
	return ListModelNewFromC(recv.ToC())
}

// MemoryInputStream is a wrapper around the C record GMemoryInputStream.
type MemoryInputStream struct {
	native *C.GMemoryInputStream
	// parent_instance : record
	// Private : priv
}

func MemoryInputStreamNewFromC(u unsafe.Pointer) *MemoryInputStream {
	c := (*C.GMemoryInputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MemoryInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MemoryInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryInputStream with another MemoryInputStream, and returns true if they represent the same GObject.
func (recv *MemoryInputStream) Equals(other *MemoryInputStream) bool {
	return other.ToC() == recv.ToC()
}

// InputStream upcasts to *InputStream
func (recv *MemoryInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *MemoryInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitrary Object to MemoryInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a MemoryInputStream.
func CastToMemoryInputStream(object *gobject.Object) *MemoryInputStream {
	return MemoryInputStreamNewFromC(object.ToC())
}

// MemoryInputStreamNew is a wrapper around the C function g_memory_input_stream_new.
func MemoryInputStreamNew() *MemoryInputStream {
	retC := C.g_memory_input_stream_new()
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_memory_input_stream_add_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// PollableInputStream returns the PollableInputStream interface implemented by MemoryInputStream
func (recv *MemoryInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryInputStream
func (recv *MemoryInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// MemoryOutputStream is a wrapper around the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native *C.GMemoryOutputStream
	// parent_instance : record
	// Private : priv
}

func MemoryOutputStreamNewFromC(u unsafe.Pointer) *MemoryOutputStream {
	c := (*C.GMemoryOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MemoryOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MemoryOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryOutputStream with another MemoryOutputStream, and returns true if they represent the same GObject.
func (recv *MemoryOutputStream) Equals(other *MemoryOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// OutputStream upcasts to *OutputStream
func (recv *MemoryOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *MemoryOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to MemoryOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a MemoryOutputStream.
func CastToMemoryOutputStream(object *gobject.Object) *MemoryOutputStream {
	return MemoryOutputStreamNewFromC(object.ToC())
}

// Unsupported : g_memory_output_stream_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_memory_output_stream_get_data : no return generator

// GetSize is a wrapper around the C function g_memory_output_stream_get_size.
func (recv *MemoryOutputStream) GetSize() uint64 {
	retC := C.g_memory_output_stream_get_size((*C.GMemoryOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// PollableOutputStream returns the PollableOutputStream interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// MountOperation is a wrapper around the C record GMountOperation.
type MountOperation struct {
	native *C.GMountOperation
	// parent_instance : record
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	c := (*C.GMountOperation)(u)
	if c == nil {
		return nil
	}

	g := &MountOperation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MountOperation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MountOperation with another MountOperation, and returns true if they represent the same GObject.
func (recv *MountOperation) Equals(other *MountOperation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MountOperation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MountOperation.
// Exercise care, as this is a potentially dangerous function if the Object is not a MountOperation.
func CastToMountOperation(object *gobject.Object) *MountOperation {
	return MountOperationNewFromC(object.ToC())
}

type signalMountOperationAskPasswordDetail struct {
	callback  MountOperationSignalAskPasswordCallback
	handlerID C.gulong
}

var signalMountOperationAskPasswordId int
var signalMountOperationAskPasswordMap = make(map[int]signalMountOperationAskPasswordDetail)
var signalMountOperationAskPasswordLock sync.RWMutex

// MountOperationSignalAskPasswordCallback is a callback function for a 'ask-password' signal emitted from a MountOperation.
type MountOperationSignalAskPasswordCallback func(message string, defaultUser string, defaultDomain string, flags AskPasswordFlags)

/*
ConnectAskPassword connects the callback to the 'ask-password' signal for the MountOperation.

The returned value represents the connection, and may be passed to DisconnectAskPassword to remove it.
*/
func (recv *MountOperation) ConnectAskPassword(callback MountOperationSignalAskPasswordCallback) int {
	signalMountOperationAskPasswordLock.Lock()
	defer signalMountOperationAskPasswordLock.Unlock()

	signalMountOperationAskPasswordId++
	instance := C.gpointer(recv.native)
	handlerID := C.MountOperation_signal_connect_ask_password(instance, C.gpointer(uintptr(signalMountOperationAskPasswordId)))

	detail := signalMountOperationAskPasswordDetail{callback, handlerID}
	signalMountOperationAskPasswordMap[signalMountOperationAskPasswordId] = detail

	return signalMountOperationAskPasswordId
}

/*
DisconnectAskPassword disconnects a callback from the 'ask-password' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectAskPassword.
*/
func (recv *MountOperation) DisconnectAskPassword(connectionID int) {
	signalMountOperationAskPasswordLock.Lock()
	defer signalMountOperationAskPasswordLock.Unlock()

	detail, exists := signalMountOperationAskPasswordMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountOperationAskPasswordMap, connectionID)
}

//export mountoperation_askPasswordHandler
func mountoperation_askPasswordHandler(_ *C.GObject, c_message *C.gchar, c_default_user *C.gchar, c_default_domain *C.gchar, c_flags C.GAskPasswordFlags, data C.gpointer) {
	signalMountOperationAskPasswordLock.RLock()
	defer signalMountOperationAskPasswordLock.RUnlock()

	message := C.GoString(c_message)

	defaultUser := C.GoString(c_default_user)

	defaultDomain := C.GoString(c_default_domain)

	flags := AskPasswordFlags(c_flags)

	index := int(uintptr(data))
	callback := signalMountOperationAskPasswordMap[index].callback
	callback(message, defaultUser, defaultDomain, flags)
}

// Unsupported signal 'ask-question' for MountOperation : unsupported parameter choices :

type signalMountOperationReplyDetail struct {
	callback  MountOperationSignalReplyCallback
	handlerID C.gulong
}

var signalMountOperationReplyId int
var signalMountOperationReplyMap = make(map[int]signalMountOperationReplyDetail)
var signalMountOperationReplyLock sync.RWMutex

// MountOperationSignalReplyCallback is a callback function for a 'reply' signal emitted from a MountOperation.
type MountOperationSignalReplyCallback func(result MountOperationResult)

/*
ConnectReply connects the callback to the 'reply' signal for the MountOperation.

The returned value represents the connection, and may be passed to DisconnectReply to remove it.
*/
func (recv *MountOperation) ConnectReply(callback MountOperationSignalReplyCallback) int {
	signalMountOperationReplyLock.Lock()
	defer signalMountOperationReplyLock.Unlock()

	signalMountOperationReplyId++
	instance := C.gpointer(recv.native)
	handlerID := C.MountOperation_signal_connect_reply(instance, C.gpointer(uintptr(signalMountOperationReplyId)))

	detail := signalMountOperationReplyDetail{callback, handlerID}
	signalMountOperationReplyMap[signalMountOperationReplyId] = detail

	return signalMountOperationReplyId
}

/*
DisconnectReply disconnects a callback from the 'reply' signal for the MountOperation.

The connectionID should be a value returned from a call to ConnectReply.
*/
func (recv *MountOperation) DisconnectReply(connectionID int) {
	signalMountOperationReplyLock.Lock()
	defer signalMountOperationReplyLock.Unlock()

	detail, exists := signalMountOperationReplyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountOperationReplyMap, connectionID)
}

//export mountoperation_replyHandler
func mountoperation_replyHandler(_ *C.GObject, c_result C.GMountOperationResult, data C.gpointer) {
	signalMountOperationReplyLock.RLock()
	defer signalMountOperationReplyLock.RUnlock()

	result := MountOperationResult(c_result)

	index := int(uintptr(data))
	callback := signalMountOperationReplyMap[index].callback
	callback(result)
}

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter processes :

// MountOperationNew is a wrapper around the C function g_mount_operation_new.
func MountOperationNew() *MountOperation {
	retC := C.g_mount_operation_new()
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetAnonymous is a wrapper around the C function g_mount_operation_get_anonymous.
func (recv *MountOperation) GetAnonymous() bool {
	retC := C.g_mount_operation_get_anonymous((*C.GMountOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetChoice is a wrapper around the C function g_mount_operation_get_choice.
func (recv *MountOperation) GetChoice() int32 {
	retC := C.g_mount_operation_get_choice((*C.GMountOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDomain is a wrapper around the C function g_mount_operation_get_domain.
func (recv *MountOperation) GetDomain() string {
	retC := C.g_mount_operation_get_domain((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPassword is a wrapper around the C function g_mount_operation_get_password.
func (recv *MountOperation) GetPassword() string {
	retC := C.g_mount_operation_get_password((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPasswordSave is a wrapper around the C function g_mount_operation_get_password_save.
func (recv *MountOperation) GetPasswordSave() PasswordSave {
	retC := C.g_mount_operation_get_password_save((*C.GMountOperation)(recv.native))
	retGo := (PasswordSave)(retC)

	return retGo
}

// GetUsername is a wrapper around the C function g_mount_operation_get_username.
func (recv *MountOperation) GetUsername() string {
	retC := C.g_mount_operation_get_username((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Reply is a wrapper around the C function g_mount_operation_reply.
func (recv *MountOperation) Reply(result MountOperationResult) {
	c_result := (C.GMountOperationResult)(result)

	C.g_mount_operation_reply((*C.GMountOperation)(recv.native), c_result)

	return
}

// SetAnonymous is a wrapper around the C function g_mount_operation_set_anonymous.
func (recv *MountOperation) SetAnonymous(anonymous bool) {
	c_anonymous :=
		boolToGboolean(anonymous)

	C.g_mount_operation_set_anonymous((*C.GMountOperation)(recv.native), c_anonymous)

	return
}

// SetChoice is a wrapper around the C function g_mount_operation_set_choice.
func (recv *MountOperation) SetChoice(choice int32) {
	c_choice := (C.int)(choice)

	C.g_mount_operation_set_choice((*C.GMountOperation)(recv.native), c_choice)

	return
}

// SetDomain is a wrapper around the C function g_mount_operation_set_domain.
func (recv *MountOperation) SetDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_mount_operation_set_domain((*C.GMountOperation)(recv.native), c_domain)

	return
}

// SetPassword is a wrapper around the C function g_mount_operation_set_password.
func (recv *MountOperation) SetPassword(password string) {
	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.g_mount_operation_set_password((*C.GMountOperation)(recv.native), c_password)

	return
}

// SetPasswordSave is a wrapper around the C function g_mount_operation_set_password_save.
func (recv *MountOperation) SetPasswordSave(save PasswordSave) {
	c_save := (C.GPasswordSave)(save)

	C.g_mount_operation_set_password_save((*C.GMountOperation)(recv.native), c_save)

	return
}

// SetUsername is a wrapper around the C function g_mount_operation_set_username.
func (recv *MountOperation) SetUsername(username string) {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	C.g_mount_operation_set_username((*C.GMountOperation)(recv.native), c_username)

	return
}

// NativeVolumeMonitor is a wrapper around the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native *C.GNativeVolumeMonitor
	// parent_instance : record
}

func NativeVolumeMonitorNewFromC(u unsafe.Pointer) *NativeVolumeMonitor {
	c := (*C.GNativeVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NativeVolumeMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeVolumeMonitor with another NativeVolumeMonitor, and returns true if they represent the same GObject.
func (recv *NativeVolumeMonitor) Equals(other *NativeVolumeMonitor) bool {
	return other.ToC() == recv.ToC()
}

// VolumeMonitor upcasts to *VolumeMonitor
func (recv *NativeVolumeMonitor) VolumeMonitor() *VolumeMonitor {
	return VolumeMonitorNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *NativeVolumeMonitor) Object() *gobject.Object {
	return recv.VolumeMonitor().Object()
}

// CastToWidget down casts any arbitrary Object to NativeVolumeMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a NativeVolumeMonitor.
func CastToNativeVolumeMonitor(object *gobject.Object) *NativeVolumeMonitor {
	return NativeVolumeMonitorNewFromC(object.ToC())
}

// NetworkAddress is a wrapper around the C record GNetworkAddress.
type NetworkAddress struct {
	native *C.GNetworkAddress
	// parent_instance : record
	// Private : priv
}

func NetworkAddressNewFromC(u unsafe.Pointer) *NetworkAddress {
	c := (*C.GNetworkAddress)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NetworkAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NetworkAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkAddress with another NetworkAddress, and returns true if they represent the same GObject.
func (recv *NetworkAddress) Equals(other *NetworkAddress) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *NetworkAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to NetworkAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a NetworkAddress.
func CastToNetworkAddress(object *gobject.Object) *NetworkAddress {
	return NetworkAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkAddress
func (recv *NetworkAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// NetworkService is a wrapper around the C record GNetworkService.
type NetworkService struct {
	native *C.GNetworkService
	// parent_instance : record
	// Private : priv
}

func NetworkServiceNewFromC(u unsafe.Pointer) *NetworkService {
	c := (*C.GNetworkService)(u)
	if c == nil {
		return nil
	}

	g := &NetworkService{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NetworkService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NetworkService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkService with another NetworkService, and returns true if they represent the same GObject.
func (recv *NetworkService) Equals(other *NetworkService) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *NetworkService) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to NetworkService.
// Exercise care, as this is a potentially dangerous function if the Object is not a NetworkService.
func CastToNetworkService(object *gobject.Object) *NetworkService {
	return NetworkServiceNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkService
func (recv *NetworkService) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// OutputStream is a wrapper around the C record GOutputStream.
type OutputStream struct {
	native *C.GOutputStream
	// parent_instance : record
	// Private : priv
}

func OutputStreamNewFromC(u unsafe.Pointer) *OutputStream {
	c := (*C.GOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &OutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *OutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *OutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OutputStream with another OutputStream, and returns true if they represent the same GObject.
func (recv *OutputStream) Equals(other *OutputStream) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *OutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to OutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a OutputStream.
func CastToOutputStream(object *gobject.Object) *OutputStream {
	return OutputStreamNewFromC(object.ToC())
}

// ClearPending is a wrapper around the C function g_output_stream_clear_pending.
func (recv *OutputStream) ClearPending() {
	C.g_output_stream_clear_pending((*C.GOutputStream)(recv.native))

	return
}

// Close is a wrapper around the C function g_output_stream_close.
func (recv *OutputStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_close((*C.GOutputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_output_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_output_stream_close_finish.
func (recv *OutputStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_close_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Flush is a wrapper around the C function g_output_stream_flush.
func (recv *OutputStream) Flush(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_flush((*C.GOutputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_output_stream_flush_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// FlushFinish is a wrapper around the C function g_output_stream_flush_finish.
func (recv *OutputStream) FlushFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_flush_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasPending is a wrapper around the C function g_output_stream_has_pending.
func (recv *OutputStream) HasPending() bool {
	retC := C.g_output_stream_has_pending((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsClosed is a wrapper around the C function g_output_stream_is_closed.
func (recv *OutputStream) IsClosed() bool {
	retC := C.g_output_stream_is_closed((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPending is a wrapper around the C function g_output_stream_set_pending.
func (recv *OutputStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_output_stream_set_pending((*C.GOutputStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Splice is a wrapper around the C function g_output_stream_splice.
func (recv *OutputStream) Splice(source *InputStream, flags OutputStreamSpliceFlags, cancellable *Cancellable) (int64, error) {
	c_source := (*C.GInputStream)(C.NULL)
	if source != nil {
		c_source = (*C.GInputStream)(source.ToC())
	}

	c_flags := (C.GOutputStreamSpliceFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_splice((*C.GOutputStream)(recv.native), c_source, c_flags, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_output_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SpliceFinish is a wrapper around the C function g_output_stream_splice_finish.
func (recv *OutputStream) SpliceFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_splice_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Write is a wrapper around the C function g_output_stream_write.
func (recv *OutputStream) Write(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (unsafe.Pointer(c_buffer_arrayPtr))

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_write((*C.GOutputStream)(recv.native), c_buffer, c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WriteAll is a wrapper around the C function g_output_stream_write_all.
func (recv *OutputStream) WriteAll(buffer []uint8, cancellable *Cancellable) (bool, uint64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (unsafe.Pointer(c_buffer_arrayPtr))

	c_count := (C.gsize)(len(buffer))

	var c_bytes_written C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_all((*C.GOutputStream)(recv.native), c_buffer, c_count, &c_bytes_written, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goError
}

// Unsupported : g_output_stream_write_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_output_stream_write_bytes

// Unsupported : g_output_stream_write_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WriteBytesFinish is a wrapper around the C function g_output_stream_write_bytes_finish.
func (recv *OutputStream) WriteBytesFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_bytes_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WriteFinish is a wrapper around the C function g_output_stream_write_finish.
func (recv *OutputStream) WriteFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Permission is a wrapper around the C record GPermission.
type Permission struct {
	native *C.GPermission
	// parent_instance : record
	// Private : priv
}

func PermissionNewFromC(u unsafe.Pointer) *Permission {
	c := (*C.GPermission)(u)
	if c == nil {
		return nil
	}

	g := &Permission{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Permission) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Permission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Permission with another Permission, and returns true if they represent the same GObject.
func (recv *Permission) Equals(other *Permission) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Permission) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Permission.
// Exercise care, as this is a potentially dangerous function if the Object is not a Permission.
func CastToPermission(object *gobject.Object) *Permission {
	return PermissionNewFromC(object.ToC())
}

// ProxyAddressEnumerator is a wrapper around the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native *C.GProxyAddressEnumerator
	// parent_instance : record
	// priv : record
}

func ProxyAddressEnumeratorNewFromC(u unsafe.Pointer) *ProxyAddressEnumerator {
	c := (*C.GProxyAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyAddressEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyAddressEnumerator with another ProxyAddressEnumerator, and returns true if they represent the same GObject.
func (recv *ProxyAddressEnumerator) Equals(other *ProxyAddressEnumerator) bool {
	return other.ToC() == recv.ToC()
}

// SocketAddressEnumerator upcasts to *SocketAddressEnumerator
func (recv *ProxyAddressEnumerator) SocketAddressEnumerator() *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *ProxyAddressEnumerator) Object() *gobject.Object {
	return recv.SocketAddressEnumerator().Object()
}

// CastToWidget down casts any arbitrary Object to ProxyAddressEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProxyAddressEnumerator.
func CastToProxyAddressEnumerator(object *gobject.Object) *ProxyAddressEnumerator {
	return ProxyAddressEnumeratorNewFromC(object.ToC())
}

// Resolver is a wrapper around the C record GResolver.
type Resolver struct {
	native *C.GResolver
	// parent_instance : record
	// priv : record
}

func ResolverNewFromC(u unsafe.Pointer) *Resolver {
	c := (*C.GResolver)(u)
	if c == nil {
		return nil
	}

	g := &Resolver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Resolver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Resolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Resolver with another Resolver, and returns true if they represent the same GObject.
func (recv *Resolver) Equals(other *Resolver) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Resolver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Resolver.
// Exercise care, as this is a potentially dangerous function if the Object is not a Resolver.
func CastToResolver(object *gobject.Object) *Resolver {
	return ResolverNewFromC(object.ToC())
}

type signalResolverReloadDetail struct {
	callback  ResolverSignalReloadCallback
	handlerID C.gulong
}

var signalResolverReloadId int
var signalResolverReloadMap = make(map[int]signalResolverReloadDetail)
var signalResolverReloadLock sync.RWMutex

// ResolverSignalReloadCallback is a callback function for a 'reload' signal emitted from a Resolver.
type ResolverSignalReloadCallback func()

/*
ConnectReload connects the callback to the 'reload' signal for the Resolver.

The returned value represents the connection, and may be passed to DisconnectReload to remove it.
*/
func (recv *Resolver) ConnectReload(callback ResolverSignalReloadCallback) int {
	signalResolverReloadLock.Lock()
	defer signalResolverReloadLock.Unlock()

	signalResolverReloadId++
	instance := C.gpointer(recv.native)
	handlerID := C.Resolver_signal_connect_reload(instance, C.gpointer(uintptr(signalResolverReloadId)))

	detail := signalResolverReloadDetail{callback, handlerID}
	signalResolverReloadMap[signalResolverReloadId] = detail

	return signalResolverReloadId
}

/*
DisconnectReload disconnects a callback from the 'reload' signal for the Resolver.

The connectionID should be a value returned from a call to ConnectReload.
*/
func (recv *Resolver) DisconnectReload(connectionID int) {
	signalResolverReloadLock.Lock()
	defer signalResolverReloadLock.Unlock()

	detail, exists := signalResolverReloadMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalResolverReloadMap, connectionID)
}

//export resolver_reloadHandler
func resolver_reloadHandler(_ *C.GObject, data C.gpointer) {
	signalResolverReloadLock.RLock()
	defer signalResolverReloadLock.RUnlock()

	index := int(uintptr(data))
	callback := signalResolverReloadMap[index].callback
	callback()
}

// Settings is a wrapper around the C record GSettings.
type Settings struct {
	native *C.GSettings
	// parent_instance : record
	// priv : record
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.GSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Settings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Settings with another Settings, and returns true if they represent the same GObject.
func (recv *Settings) Equals(other *Settings) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Settings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Settings.
// Exercise care, as this is a potentially dangerous function if the Object is not a Settings.
func CastToSettings(object *gobject.Object) *Settings {
	return SettingsNewFromC(object.ToC())
}

// Unsupported signal 'change-event' for Settings : unsupported parameter keys :

type signalSettingsChangedDetail struct {
	callback  SettingsSignalChangedCallback
	handlerID C.gulong
}

var signalSettingsChangedId int
var signalSettingsChangedMap = make(map[int]signalSettingsChangedDetail)
var signalSettingsChangedLock sync.RWMutex

// SettingsSignalChangedCallback is a callback function for a 'changed' signal emitted from a Settings.
type SettingsSignalChangedCallback func(key string)

/*
ConnectChanged connects the callback to the 'changed' signal for the Settings.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Settings) ConnectChanged(callback SettingsSignalChangedCallback) int {
	signalSettingsChangedLock.Lock()
	defer signalSettingsChangedLock.Unlock()

	signalSettingsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Settings_signal_connect_changed(instance, C.gpointer(uintptr(signalSettingsChangedId)))

	detail := signalSettingsChangedDetail{callback, handlerID}
	signalSettingsChangedMap[signalSettingsChangedId] = detail

	return signalSettingsChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Settings.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Settings) DisconnectChanged(connectionID int) {
	signalSettingsChangedLock.Lock()
	defer signalSettingsChangedLock.Unlock()

	detail, exists := signalSettingsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSettingsChangedMap, connectionID)
}

//export settings_changedHandler
func settings_changedHandler(_ *C.GObject, c_key *C.gchar, data C.gpointer) {
	signalSettingsChangedLock.RLock()
	defer signalSettingsChangedLock.RUnlock()

	key := C.GoString(c_key)

	index := int(uintptr(data))
	callback := signalSettingsChangedMap[index].callback
	callback(key)
}

type signalSettingsWritableChangeEventDetail struct {
	callback  SettingsSignalWritableChangeEventCallback
	handlerID C.gulong
}

var signalSettingsWritableChangeEventId int
var signalSettingsWritableChangeEventMap = make(map[int]signalSettingsWritableChangeEventDetail)
var signalSettingsWritableChangeEventLock sync.RWMutex

// SettingsSignalWritableChangeEventCallback is a callback function for a 'writable-change-event' signal emitted from a Settings.
type SettingsSignalWritableChangeEventCallback func(key uint32) bool

/*
ConnectWritableChangeEvent connects the callback to the 'writable-change-event' signal for the Settings.

The returned value represents the connection, and may be passed to DisconnectWritableChangeEvent to remove it.
*/
func (recv *Settings) ConnectWritableChangeEvent(callback SettingsSignalWritableChangeEventCallback) int {
	signalSettingsWritableChangeEventLock.Lock()
	defer signalSettingsWritableChangeEventLock.Unlock()

	signalSettingsWritableChangeEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Settings_signal_connect_writable_change_event(instance, C.gpointer(uintptr(signalSettingsWritableChangeEventId)))

	detail := signalSettingsWritableChangeEventDetail{callback, handlerID}
	signalSettingsWritableChangeEventMap[signalSettingsWritableChangeEventId] = detail

	return signalSettingsWritableChangeEventId
}

/*
DisconnectWritableChangeEvent disconnects a callback from the 'writable-change-event' signal for the Settings.

The connectionID should be a value returned from a call to ConnectWritableChangeEvent.
*/
func (recv *Settings) DisconnectWritableChangeEvent(connectionID int) {
	signalSettingsWritableChangeEventLock.Lock()
	defer signalSettingsWritableChangeEventLock.Unlock()

	detail, exists := signalSettingsWritableChangeEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSettingsWritableChangeEventMap, connectionID)
}

//export settings_writableChangeEventHandler
func settings_writableChangeEventHandler(_ *C.GObject, c_key C.guint, data C.gpointer) C.gboolean {
	signalSettingsWritableChangeEventLock.RLock()
	defer signalSettingsWritableChangeEventLock.RUnlock()

	key := uint32(c_key)

	index := int(uintptr(data))
	callback := signalSettingsWritableChangeEventMap[index].callback
	retGo := callback(key)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalSettingsWritableChangedDetail struct {
	callback  SettingsSignalWritableChangedCallback
	handlerID C.gulong
}

var signalSettingsWritableChangedId int
var signalSettingsWritableChangedMap = make(map[int]signalSettingsWritableChangedDetail)
var signalSettingsWritableChangedLock sync.RWMutex

// SettingsSignalWritableChangedCallback is a callback function for a 'writable-changed' signal emitted from a Settings.
type SettingsSignalWritableChangedCallback func(key string)

/*
ConnectWritableChanged connects the callback to the 'writable-changed' signal for the Settings.

The returned value represents the connection, and may be passed to DisconnectWritableChanged to remove it.
*/
func (recv *Settings) ConnectWritableChanged(callback SettingsSignalWritableChangedCallback) int {
	signalSettingsWritableChangedLock.Lock()
	defer signalSettingsWritableChangedLock.Unlock()

	signalSettingsWritableChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Settings_signal_connect_writable_changed(instance, C.gpointer(uintptr(signalSettingsWritableChangedId)))

	detail := signalSettingsWritableChangedDetail{callback, handlerID}
	signalSettingsWritableChangedMap[signalSettingsWritableChangedId] = detail

	return signalSettingsWritableChangedId
}

/*
DisconnectWritableChanged disconnects a callback from the 'writable-changed' signal for the Settings.

The connectionID should be a value returned from a call to ConnectWritableChanged.
*/
func (recv *Settings) DisconnectWritableChanged(connectionID int) {
	signalSettingsWritableChangedLock.Lock()
	defer signalSettingsWritableChangedLock.Unlock()

	detail, exists := signalSettingsWritableChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSettingsWritableChangedMap, connectionID)
}

//export settings_writableChangedHandler
func settings_writableChangedHandler(_ *C.GObject, c_key *C.gchar, data C.gpointer) {
	signalSettingsWritableChangedLock.RLock()
	defer signalSettingsWritableChangedLock.RUnlock()

	key := C.GoString(c_key)

	index := int(uintptr(data))
	callback := signalSettingsWritableChangedMap[index].callback
	callback(key)
}

// SettingsSync is a wrapper around the C function g_settings_sync.
func SettingsSync() {
	C.g_settings_sync()

	return
}

// Apply is a wrapper around the C function g_settings_apply.
func (recv *Settings) Apply() {
	C.g_settings_apply((*C.GSettings)(recv.native))

	return
}

// Unsupported : g_settings_get_mapped : unsupported parameter mapping : no type generator for SettingsGetMapping (GSettingsGetMapping) for param mapping

// ListChildren is a wrapper around the C function g_settings_list_children.
func (recv *Settings) ListChildren() []string {
	retC := C.g_settings_list_children((*C.GSettings)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// ListKeys is a wrapper around the C function g_settings_list_keys.
func (recv *Settings) ListKeys() []string {
	retC := C.g_settings_list_keys((*C.GSettings)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// Reset is a wrapper around the C function g_settings_reset.
func (recv *Settings) Reset(key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	C.g_settings_reset((*C.GSettings)(recv.native), c_key)

	return
}

// Revert is a wrapper around the C function g_settings_revert.
func (recv *Settings) Revert() {
	C.g_settings_revert((*C.GSettings)(recv.native))

	return
}

// SetEnum is a wrapper around the C function g_settings_set_enum.
func (recv *Settings) SetEnum(key string, value int32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	retC := C.g_settings_set_enum((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SetFlags is a wrapper around the C function g_settings_set_flags.
func (recv *Settings) SetFlags(key string, value uint32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint)(value)

	retC := C.g_settings_set_flags((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// SettingsBackend is a wrapper around the C record GSettingsBackend.
type SettingsBackend struct {
	native *C.GSettingsBackend
	// parent_instance : record
	// Private : priv
}

func SettingsBackendNewFromC(u unsafe.Pointer) *SettingsBackend {
	c := (*C.GSettingsBackend)(u)
	if c == nil {
		return nil
	}

	g := &SettingsBackend{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SettingsBackend) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SettingsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsBackend with another SettingsBackend, and returns true if they represent the same GObject.
func (recv *SettingsBackend) Equals(other *SettingsBackend) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SettingsBackend) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SettingsBackend.
// Exercise care, as this is a potentially dangerous function if the Object is not a SettingsBackend.
func CastToSettingsBackend(object *gobject.Object) *SettingsBackend {
	return SettingsBackendNewFromC(object.ToC())
}

// SimpleAction is a wrapper around the C record GSimpleAction.
type SimpleAction struct {
	native *C.GSimpleAction
}

func SimpleActionNewFromC(u unsafe.Pointer) *SimpleAction {
	c := (*C.GSimpleAction)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleAction with another SimpleAction, and returns true if they represent the same GObject.
func (recv *SimpleAction) Equals(other *SimpleAction) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SimpleAction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SimpleAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleAction.
func CastToSimpleAction(object *gobject.Object) *SimpleAction {
	return SimpleActionNewFromC(object.ToC())
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// SimpleAsyncResult is a wrapper around the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native *C.GSimpleAsyncResult
}

func SimpleAsyncResultNewFromC(u unsafe.Pointer) *SimpleAsyncResult {
	c := (*C.GSimpleAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResult{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleAsyncResult) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleAsyncResult with another SimpleAsyncResult, and returns true if they represent the same GObject.
func (recv *SimpleAsyncResult) Equals(other *SimpleAsyncResult) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SimpleAsyncResult) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SimpleAsyncResult.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleAsyncResult.
func CastToSimpleAsyncResult(object *gobject.Object) *SimpleAsyncResult {
	return SimpleAsyncResultNewFromC(object.ToC())
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Complete is a wrapper around the C function g_simple_async_result_complete.
func (recv *SimpleAsyncResult) Complete() {
	C.g_simple_async_result_complete((*C.GSimpleAsyncResult)(recv.native))

	return
}

// CompleteInIdle is a wrapper around the C function g_simple_async_result_complete_in_idle.
func (recv *SimpleAsyncResult) CompleteInIdle() {
	C.g_simple_async_result_complete_in_idle((*C.GSimpleAsyncResult)(recv.native))

	return
}

// GetOpResGboolean is a wrapper around the C function g_simple_async_result_get_op_res_gboolean.
func (recv *SimpleAsyncResult) GetOpResGboolean() bool {
	retC := C.g_simple_async_result_get_op_res_gboolean((*C.GSimpleAsyncResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_simple_async_result_get_op_res_gpointer : no return generator

// GetOpResGssize is a wrapper around the C function g_simple_async_result_get_op_res_gssize.
func (recv *SimpleAsyncResult) GetOpResGssize() int64 {
	retC := C.g_simple_async_result_get_op_res_gssize((*C.GSimpleAsyncResult)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Unsupported : g_simple_async_result_get_source_tag : no return generator

// PropagateError is a wrapper around the C function g_simple_async_result_propagate_error.
func (recv *SimpleAsyncResult) PropagateError() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_simple_async_result_propagate_error((*C.GSimpleAsyncResult)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_simple_async_result_run_in_thread : unsupported parameter func : no type generator for SimpleAsyncThreadFunc (GSimpleAsyncThreadFunc) for param func

// SetError is a wrapper around the C function g_simple_async_result_set_error.
func (recv *SimpleAsyncResult) SetError(domain glib.Quark, code int32, format string, args ...interface{}) {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_simple_async_result_set_error((*C.GSimpleAsyncResult)(recv.native), c_domain, c_code, c_format)

	return
}

// Unsupported : g_simple_async_result_set_error_va : unsupported parameter args : no type generator for va_list (va_list) for param args

// SetFromError is a wrapper around the C function g_simple_async_result_set_from_error.
func (recv *SimpleAsyncResult) SetFromError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_simple_async_result_set_from_error((*C.GSimpleAsyncResult)(recv.native), c_error)

	return
}

// SetHandleCancellation is a wrapper around the C function g_simple_async_result_set_handle_cancellation.
func (recv *SimpleAsyncResult) SetHandleCancellation(handleCancellation bool) {
	c_handle_cancellation :=
		boolToGboolean(handleCancellation)

	C.g_simple_async_result_set_handle_cancellation((*C.GSimpleAsyncResult)(recv.native), c_handle_cancellation)

	return
}

// SetOpResGboolean is a wrapper around the C function g_simple_async_result_set_op_res_gboolean.
func (recv *SimpleAsyncResult) SetOpResGboolean(opRes bool) {
	c_op_res :=
		boolToGboolean(opRes)

	C.g_simple_async_result_set_op_res_gboolean((*C.GSimpleAsyncResult)(recv.native), c_op_res)

	return
}

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter op_res : no type generator for gpointer (gpointer) for param op_res

// SetOpResGssize is a wrapper around the C function g_simple_async_result_set_op_res_gssize.
func (recv *SimpleAsyncResult) SetOpResGssize(opRes int64) {
	c_op_res := (C.gssize)(opRes)

	C.g_simple_async_result_set_op_res_gssize((*C.GSimpleAsyncResult)(recv.native), c_op_res)

	return
}

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// SimplePermission is a wrapper around the C record GSimplePermission.
type SimplePermission struct {
	native *C.GSimplePermission
}

func SimplePermissionNewFromC(u unsafe.Pointer) *SimplePermission {
	c := (*C.GSimplePermission)(u)
	if c == nil {
		return nil
	}

	g := &SimplePermission{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimplePermission) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimplePermission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimplePermission with another SimplePermission, and returns true if they represent the same GObject.
func (recv *SimplePermission) Equals(other *SimplePermission) bool {
	return other.ToC() == recv.ToC()
}

// Permission upcasts to *Permission
func (recv *SimplePermission) Permission() *Permission {
	return PermissionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimplePermission) Object() *gobject.Object {
	return recv.Permission().Object()
}

// CastToWidget down casts any arbitrary Object to SimplePermission.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimplePermission.
func CastToSimplePermission(object *gobject.Object) *SimplePermission {
	return SimplePermissionNewFromC(object.ToC())
}

// SimpleProxyResolver is a wrapper around the C record GSimpleProxyResolver.
type SimpleProxyResolver struct {
	native *C.GSimpleProxyResolver
	// parent_instance : record
	// Private : priv
}

func SimpleProxyResolverNewFromC(u unsafe.Pointer) *SimpleProxyResolver {
	c := (*C.GSimpleProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleProxyResolver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleProxyResolver with another SimpleProxyResolver, and returns true if they represent the same GObject.
func (recv *SimpleProxyResolver) Equals(other *SimpleProxyResolver) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SimpleProxyResolver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SimpleProxyResolver.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleProxyResolver.
func CastToSimpleProxyResolver(object *gobject.Object) *SimpleProxyResolver {
	return SimpleProxyResolverNewFromC(object.ToC())
}

// ProxyResolver returns the ProxyResolver interface implemented by SimpleProxyResolver
func (recv *SimpleProxyResolver) ProxyResolver() *ProxyResolver {
	return ProxyResolverNewFromC(recv.ToC())
}

// SocketAddress is a wrapper around the C record GSocketAddress.
type SocketAddress struct {
	native *C.GSocketAddress
	// parent_instance : record
}

func SocketAddressNewFromC(u unsafe.Pointer) *SocketAddress {
	c := (*C.GSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketAddress with another SocketAddress, and returns true if they represent the same GObject.
func (recv *SocketAddress) Equals(other *SocketAddress) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketAddress.
func CastToSocketAddress(object *gobject.Object) *SocketAddress {
	return SocketAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by SocketAddress
func (recv *SocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// SocketAddressEnumerator is a wrapper around the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native *C.GSocketAddressEnumerator
	// parent_instance : record
}

func SocketAddressEnumeratorNewFromC(u unsafe.Pointer) *SocketAddressEnumerator {
	c := (*C.GSocketAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketAddressEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketAddressEnumerator with another SocketAddressEnumerator, and returns true if they represent the same GObject.
func (recv *SocketAddressEnumerator) Equals(other *SocketAddressEnumerator) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketAddressEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketAddressEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketAddressEnumerator.
func CastToSocketAddressEnumerator(object *gobject.Object) *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromC(object.ToC())
}

// Next is a wrapper around the C function g_socket_address_enumerator_next.
func (recv *SocketAddressEnumerator) Next(cancellable *Cancellable) (*SocketAddress, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_address_enumerator_next((*C.GSocketAddressEnumerator)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_address_enumerator_next_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// NextFinish is a wrapper around the C function g_socket_address_enumerator_next_finish.
func (recv *SocketAddressEnumerator) NextFinish(result *AsyncResult) (*SocketAddress, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_address_enumerator_next_finish((*C.GSocketAddressEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SocketControlMessage is a wrapper around the C record GSocketControlMessage.
type SocketControlMessage struct {
	native *C.GSocketControlMessage
	// parent_instance : record
	// priv : record
}

func SocketControlMessageNewFromC(u unsafe.Pointer) *SocketControlMessage {
	c := (*C.GSocketControlMessage)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketControlMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketControlMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketControlMessage with another SocketControlMessage, and returns true if they represent the same GObject.
func (recv *SocketControlMessage) Equals(other *SocketControlMessage) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketControlMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketControlMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketControlMessage.
func CastToSocketControlMessage(object *gobject.Object) *SocketControlMessage {
	return SocketControlMessageNewFromC(object.ToC())
}

// Task is a wrapper around the C record GTask.
type Task struct {
	native *C.GTask
}

func TaskNewFromC(u unsafe.Pointer) *Task {
	c := (*C.GTask)(u)
	if c == nil {
		return nil
	}

	g := &Task{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Task) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Task) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Task with another Task, and returns true if they represent the same GObject.
func (recv *Task) Equals(other *Task) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Task) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Task.
// Exercise care, as this is a potentially dangerous function if the Object is not a Task.
func CastToTask(object *gobject.Object) *Task {
	return TaskNewFromC(object.ToC())
}

// AsyncResult returns the AsyncResult interface implemented by Task
func (recv *Task) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// TcpWrapperConnection is a wrapper around the C record GTcpWrapperConnection.
type TcpWrapperConnection struct {
	native *C.GTcpWrapperConnection
	// parent_instance : record
	// priv : record
}

func TcpWrapperConnectionNewFromC(u unsafe.Pointer) *TcpWrapperConnection {
	c := (*C.GTcpWrapperConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TcpWrapperConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpWrapperConnection with another TcpWrapperConnection, and returns true if they represent the same GObject.
func (recv *TcpWrapperConnection) Equals(other *TcpWrapperConnection) bool {
	return other.ToC() == recv.ToC()
}

// TcpConnection upcasts to *TcpConnection
func (recv *TcpWrapperConnection) TcpConnection() *TcpConnection {
	return TcpConnectionNewFromC(unsafe.Pointer(recv.native))
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpWrapperConnection) SocketConnection() *SocketConnection {
	return recv.TcpConnection().SocketConnection()
}

// IOStream upcasts to *IOStream
func (recv *TcpWrapperConnection) IOStream() *IOStream {
	return recv.TcpConnection().IOStream()
}

// Object upcasts to *Object
func (recv *TcpWrapperConnection) Object() *gobject.Object {
	return recv.TcpConnection().Object()
}

// CastToWidget down casts any arbitrary Object to TcpWrapperConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TcpWrapperConnection.
func CastToTcpWrapperConnection(object *gobject.Object) *TcpWrapperConnection {
	return TcpWrapperConnectionNewFromC(object.ToC())
}

// GetBaseIoStream is a wrapper around the C function g_tcp_wrapper_connection_get_base_io_stream.
func (recv *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	retC := C.g_tcp_wrapper_connection_get_base_io_stream((*C.GTcpWrapperConnection)(recv.native))
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ThemedIcon is a wrapper around the C record GThemedIcon.
type ThemedIcon struct {
	native *C.GThemedIcon
}

func ThemedIconNewFromC(u unsafe.Pointer) *ThemedIcon {
	c := (*C.GThemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ThemedIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ThemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThemedIcon with another ThemedIcon, and returns true if they represent the same GObject.
func (recv *ThemedIcon) Equals(other *ThemedIcon) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ThemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ThemedIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThemedIcon.
func CastToThemedIcon(object *gobject.Object) *ThemedIcon {
	return ThemedIconNewFromC(object.ToC())
}

// ThemedIconNew is a wrapper around the C function g_themed_icon_new.
func ThemedIconNew(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ThemedIconNewFromNames is a wrapper around the C function g_themed_icon_new_from_names.
func ThemedIconNewFromNames(iconnames []string) *ThemedIcon {
	c_iconnames_array := make([]*C.char, len(iconnames)+1, len(iconnames)+1)
	for i, item := range iconnames {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_iconnames_array[i] = c
	}
	c_iconnames_array[len(iconnames)] = nil
	c_iconnames_arrayPtr := &c_iconnames_array[0]
	c_iconnames := (**C.char)(unsafe.Pointer(c_iconnames_arrayPtr))

	c_len := (C.int)(len(iconnames))

	retC := C.g_themed_icon_new_from_names(c_iconnames, c_len)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ThemedIconNewWithDefaultFallbacks is a wrapper around the C function g_themed_icon_new_with_default_fallbacks.
func ThemedIconNewWithDefaultFallbacks(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new_with_default_fallbacks(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AppendName is a wrapper around the C function g_themed_icon_append_name.
func (recv *ThemedIcon) AppendName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_append_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
}

// GetNames is a wrapper around the C function g_themed_icon_get_names.
func (recv *ThemedIcon) GetNames() []string {
	retC := C.g_themed_icon_get_names((*C.GThemedIcon)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// Icon returns the Icon interface implemented by ThemedIcon
func (recv *ThemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// UnixConnection is a wrapper around the C record GUnixConnection.
type UnixConnection struct {
	native *C.GUnixConnection
	// parent_instance : record
	// priv : record
}

func UnixConnectionNewFromC(u unsafe.Pointer) *UnixConnection {
	c := (*C.GUnixConnection)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixConnection with another UnixConnection, and returns true if they represent the same GObject.
func (recv *UnixConnection) Equals(other *UnixConnection) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnection upcasts to *SocketConnection
func (recv *UnixConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromC(unsafe.Pointer(recv.native))
}

// IOStream upcasts to *IOStream
func (recv *UnixConnection) IOStream() *IOStream {
	return recv.SocketConnection().IOStream()
}

// Object upcasts to *Object
func (recv *UnixConnection) Object() *gobject.Object {
	return recv.SocketConnection().Object()
}

// CastToWidget down casts any arbitrary Object to UnixConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixConnection.
func CastToUnixConnection(object *gobject.Object) *UnixConnection {
	return UnixConnectionNewFromC(object.ToC())
}

// UnixFDList is a wrapper around the C record GUnixFDList.
type UnixFDList struct {
	native *C.GUnixFDList
	// parent_instance : record
	// priv : record
}

func UnixFDListNewFromC(u unsafe.Pointer) *UnixFDList {
	c := (*C.GUnixFDList)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDList{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixFDList) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixFDList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDList with another UnixFDList, and returns true if they represent the same GObject.
func (recv *UnixFDList) Equals(other *UnixFDList) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *UnixFDList) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to UnixFDList.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixFDList.
func CastToUnixFDList(object *gobject.Object) *UnixFDList {
	return UnixFDListNewFromC(object.ToC())
}

// UnixFDMessage is a wrapper around the C record GUnixFDMessage.
type UnixFDMessage struct {
	native *C.GUnixFDMessage
	// parent_instance : record
	// priv : record
}

func UnixFDMessageNewFromC(u unsafe.Pointer) *UnixFDMessage {
	c := (*C.GUnixFDMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixFDMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixFDMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDMessage with another UnixFDMessage, and returns true if they represent the same GObject.
func (recv *UnixFDMessage) Equals(other *UnixFDMessage) bool {
	return other.ToC() == recv.ToC()
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixFDMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixFDMessage) Object() *gobject.Object {
	return recv.SocketControlMessage().Object()
}

// CastToWidget down casts any arbitrary Object to UnixFDMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixFDMessage.
func CastToUnixFDMessage(object *gobject.Object) *UnixFDMessage {
	return UnixFDMessageNewFromC(object.ToC())
}

// UnixInputStream is a wrapper around the C record GUnixInputStream.
type UnixInputStream struct {
	native *C.GUnixInputStream
	// parent_instance : record
	// Private : priv
}

func UnixInputStreamNewFromC(u unsafe.Pointer) *UnixInputStream {
	c := (*C.GUnixInputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixInputStream with another UnixInputStream, and returns true if they represent the same GObject.
func (recv *UnixInputStream) Equals(other *UnixInputStream) bool {
	return other.ToC() == recv.ToC()
}

// InputStream upcasts to *InputStream
func (recv *UnixInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitrary Object to UnixInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixInputStream.
func CastToUnixInputStream(object *gobject.Object) *UnixInputStream {
	return UnixInputStreamNewFromC(object.ToC())
}

// UnixInputStreamNew is a wrapper around the C function g_unix_input_stream_new.
func UnixInputStreamNew(fd int32, closeFd bool) *UnixInputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_input_stream_new(c_fd, c_close_fd)
	retGo := UnixInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixInputStream
func (recv *UnixInputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableInputStream returns the PollableInputStream interface implemented by UnixInputStream
func (recv *UnixInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// UnixMountMonitor is a wrapper around the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native *C.GUnixMountMonitor
}

func UnixMountMonitorNewFromC(u unsafe.Pointer) *UnixMountMonitor {
	c := (*C.GUnixMountMonitor)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixMountMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixMountMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixMountMonitor with another UnixMountMonitor, and returns true if they represent the same GObject.
func (recv *UnixMountMonitor) Equals(other *UnixMountMonitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *UnixMountMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to UnixMountMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixMountMonitor.
func CastToUnixMountMonitor(object *gobject.Object) *UnixMountMonitor {
	return UnixMountMonitorNewFromC(object.ToC())
}

type signalUnixMountMonitorMountpointsChangedDetail struct {
	callback  UnixMountMonitorSignalMountpointsChangedCallback
	handlerID C.gulong
}

var signalUnixMountMonitorMountpointsChangedId int
var signalUnixMountMonitorMountpointsChangedMap = make(map[int]signalUnixMountMonitorMountpointsChangedDetail)
var signalUnixMountMonitorMountpointsChangedLock sync.RWMutex

// UnixMountMonitorSignalMountpointsChangedCallback is a callback function for a 'mountpoints-changed' signal emitted from a UnixMountMonitor.
type UnixMountMonitorSignalMountpointsChangedCallback func()

/*
ConnectMountpointsChanged connects the callback to the 'mountpoints-changed' signal for the UnixMountMonitor.

The returned value represents the connection, and may be passed to DisconnectMountpointsChanged to remove it.
*/
func (recv *UnixMountMonitor) ConnectMountpointsChanged(callback UnixMountMonitorSignalMountpointsChangedCallback) int {
	signalUnixMountMonitorMountpointsChangedLock.Lock()
	defer signalUnixMountMonitorMountpointsChangedLock.Unlock()

	signalUnixMountMonitorMountpointsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UnixMountMonitor_signal_connect_mountpoints_changed(instance, C.gpointer(uintptr(signalUnixMountMonitorMountpointsChangedId)))

	detail := signalUnixMountMonitorMountpointsChangedDetail{callback, handlerID}
	signalUnixMountMonitorMountpointsChangedMap[signalUnixMountMonitorMountpointsChangedId] = detail

	return signalUnixMountMonitorMountpointsChangedId
}

/*
DisconnectMountpointsChanged disconnects a callback from the 'mountpoints-changed' signal for the UnixMountMonitor.

The connectionID should be a value returned from a call to ConnectMountpointsChanged.
*/
func (recv *UnixMountMonitor) DisconnectMountpointsChanged(connectionID int) {
	signalUnixMountMonitorMountpointsChangedLock.Lock()
	defer signalUnixMountMonitorMountpointsChangedLock.Unlock()

	detail, exists := signalUnixMountMonitorMountpointsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUnixMountMonitorMountpointsChangedMap, connectionID)
}

//export unixmountmonitor_mountpointsChangedHandler
func unixmountmonitor_mountpointsChangedHandler(_ *C.GObject, data C.gpointer) {
	signalUnixMountMonitorMountpointsChangedLock.RLock()
	defer signalUnixMountMonitorMountpointsChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalUnixMountMonitorMountpointsChangedMap[index].callback
	callback()
}

type signalUnixMountMonitorMountsChangedDetail struct {
	callback  UnixMountMonitorSignalMountsChangedCallback
	handlerID C.gulong
}

var signalUnixMountMonitorMountsChangedId int
var signalUnixMountMonitorMountsChangedMap = make(map[int]signalUnixMountMonitorMountsChangedDetail)
var signalUnixMountMonitorMountsChangedLock sync.RWMutex

// UnixMountMonitorSignalMountsChangedCallback is a callback function for a 'mounts-changed' signal emitted from a UnixMountMonitor.
type UnixMountMonitorSignalMountsChangedCallback func()

/*
ConnectMountsChanged connects the callback to the 'mounts-changed' signal for the UnixMountMonitor.

The returned value represents the connection, and may be passed to DisconnectMountsChanged to remove it.
*/
func (recv *UnixMountMonitor) ConnectMountsChanged(callback UnixMountMonitorSignalMountsChangedCallback) int {
	signalUnixMountMonitorMountsChangedLock.Lock()
	defer signalUnixMountMonitorMountsChangedLock.Unlock()

	signalUnixMountMonitorMountsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UnixMountMonitor_signal_connect_mounts_changed(instance, C.gpointer(uintptr(signalUnixMountMonitorMountsChangedId)))

	detail := signalUnixMountMonitorMountsChangedDetail{callback, handlerID}
	signalUnixMountMonitorMountsChangedMap[signalUnixMountMonitorMountsChangedId] = detail

	return signalUnixMountMonitorMountsChangedId
}

/*
DisconnectMountsChanged disconnects a callback from the 'mounts-changed' signal for the UnixMountMonitor.

The connectionID should be a value returned from a call to ConnectMountsChanged.
*/
func (recv *UnixMountMonitor) DisconnectMountsChanged(connectionID int) {
	signalUnixMountMonitorMountsChangedLock.Lock()
	defer signalUnixMountMonitorMountsChangedLock.Unlock()

	detail, exists := signalUnixMountMonitorMountsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUnixMountMonitorMountsChangedMap, connectionID)
}

//export unixmountmonitor_mountsChangedHandler
func unixmountmonitor_mountsChangedHandler(_ *C.GObject, data C.gpointer) {
	signalUnixMountMonitorMountsChangedLock.RLock()
	defer signalUnixMountMonitorMountsChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalUnixMountMonitorMountsChangedMap[index].callback
	callback()
}

// UnixMountMonitorNew is a wrapper around the C function g_unix_mount_monitor_new.
func UnixMountMonitorNew() *UnixMountMonitor {
	retC := C.g_unix_mount_monitor_new()
	retGo := UnixMountMonitorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixOutputStream is a wrapper around the C record GUnixOutputStream.
type UnixOutputStream struct {
	native *C.GUnixOutputStream
	// parent_instance : record
	// Private : priv
}

func UnixOutputStreamNewFromC(u unsafe.Pointer) *UnixOutputStream {
	c := (*C.GUnixOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixOutputStream with another UnixOutputStream, and returns true if they represent the same GObject.
func (recv *UnixOutputStream) Equals(other *UnixOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// OutputStream upcasts to *OutputStream
func (recv *UnixOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitrary Object to UnixOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixOutputStream.
func CastToUnixOutputStream(object *gobject.Object) *UnixOutputStream {
	return UnixOutputStreamNewFromC(object.ToC())
}

// UnixOutputStreamNew is a wrapper around the C function g_unix_output_stream_new.
func UnixOutputStreamNew(fd int32, closeFd bool) *UnixOutputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_output_stream_new(c_fd, c_close_fd)
	retGo := UnixOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixOutputStream
func (recv *UnixOutputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableOutputStream returns the PollableOutputStream interface implemented by UnixOutputStream
func (recv *UnixOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// UnixSocketAddress is a wrapper around the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native *C.GUnixSocketAddress
	// parent_instance : record
	// Private : priv
}

func UnixSocketAddressNewFromC(u unsafe.Pointer) *UnixSocketAddress {
	c := (*C.GUnixSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixSocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixSocketAddress with another UnixSocketAddress, and returns true if they represent the same GObject.
func (recv *UnixSocketAddress) Equals(other *UnixSocketAddress) bool {
	return other.ToC() == recv.ToC()
}

// SocketAddress upcasts to *SocketAddress
func (recv *UnixSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixSocketAddress) Object() *gobject.Object {
	return recv.SocketAddress().Object()
}

// CastToWidget down casts any arbitrary Object to UnixSocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixSocketAddress.
func CastToUnixSocketAddress(object *gobject.Object) *UnixSocketAddress {
	return UnixSocketAddressNewFromC(object.ToC())
}

// UnixSocketAddressNewAbstract is a wrapper around the C function g_unix_socket_address_new_abstract.
func UnixSocketAddressNewAbstract(path []rune) *UnixSocketAddress {
	c_path_array := make([]C.gchar, len(path)+1, len(path)+1)
	for i, item := range path {
		c := (C.gchar)(item)
		c_path_array[i] = c
	}
	c_path_array[len(path)] = 0
	c_path_arrayPtr := &c_path_array[0]
	c_path := (*C.gchar)(unsafe.Pointer(c_path_arrayPtr))

	c_path_len := (C.gint)(len(path))

	retC := C.g_unix_socket_address_new_abstract(c_path, c_path_len)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Vfs is a wrapper around the C record GVfs.
type Vfs struct {
	native *C.GVfs
	// parent_instance : record
}

func VfsNewFromC(u unsafe.Pointer) *Vfs {
	c := (*C.GVfs)(u)
	if c == nil {
		return nil
	}

	g := &Vfs{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Vfs) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Vfs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Vfs with another Vfs, and returns true if they represent the same GObject.
func (recv *Vfs) Equals(other *Vfs) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Vfs) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Vfs.
// Exercise care, as this is a potentially dangerous function if the Object is not a Vfs.
func CastToVfs(object *gobject.Object) *Vfs {
	return VfsNewFromC(object.ToC())
}

// VfsGetDefault is a wrapper around the C function g_vfs_get_default.
func VfsGetDefault() *Vfs {
	retC := C.g_vfs_get_default()
	retGo := VfsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VfsGetLocal is a wrapper around the C function g_vfs_get_local.
func VfsGetLocal() *Vfs {
	retC := C.g_vfs_get_local()
	retGo := VfsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFileForPath is a wrapper around the C function g_vfs_get_file_for_path.
func (recv *Vfs) GetFileForPath(path string) *File {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_vfs_get_file_for_path((*C.GVfs)(recv.native), c_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFileForUri is a wrapper around the C function g_vfs_get_file_for_uri.
func (recv *Vfs) GetFileForUri(uri string) *File {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_vfs_get_file_for_uri((*C.GVfs)(recv.native), c_uri)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSupportedUriSchemes is a wrapper around the C function g_vfs_get_supported_uri_schemes.
func (recv *Vfs) GetSupportedUriSchemes() []string {
	retC := C.g_vfs_get_supported_uri_schemes((*C.GVfs)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// IsActive is a wrapper around the C function g_vfs_is_active.
func (recv *Vfs) IsActive() bool {
	retC := C.g_vfs_is_active((*C.GVfs)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ParseName is a wrapper around the C function g_vfs_parse_name.
func (recv *Vfs) ParseName(parseName string) *File {
	c_parse_name := C.CString(parseName)
	defer C.free(unsafe.Pointer(c_parse_name))

	retC := C.g_vfs_parse_name((*C.GVfs)(recv.native), c_parse_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VolumeMonitor is a wrapper around the C record GVolumeMonitor.
type VolumeMonitor struct {
	native *C.GVolumeMonitor
	// parent_instance : record
	// Private : priv
}

func VolumeMonitorNewFromC(u unsafe.Pointer) *VolumeMonitor {
	c := (*C.GVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VolumeMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VolumeMonitor with another VolumeMonitor, and returns true if they represent the same GObject.
func (recv *VolumeMonitor) Equals(other *VolumeMonitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *VolumeMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to VolumeMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a VolumeMonitor.
func CastToVolumeMonitor(object *gobject.Object) *VolumeMonitor {
	return VolumeMonitorNewFromC(object.ToC())
}

type signalVolumeMonitorDriveChangedDetail struct {
	callback  VolumeMonitorSignalDriveChangedCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveChangedId int
var signalVolumeMonitorDriveChangedMap = make(map[int]signalVolumeMonitorDriveChangedDetail)
var signalVolumeMonitorDriveChangedLock sync.RWMutex

// VolumeMonitorSignalDriveChangedCallback is a callback function for a 'drive-changed' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveChangedCallback func(drive *Drive)

/*
ConnectDriveChanged connects the callback to the 'drive-changed' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveChanged to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveChanged(callback VolumeMonitorSignalDriveChangedCallback) int {
	signalVolumeMonitorDriveChangedLock.Lock()
	defer signalVolumeMonitorDriveChangedLock.Unlock()

	signalVolumeMonitorDriveChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_changed(instance, C.gpointer(uintptr(signalVolumeMonitorDriveChangedId)))

	detail := signalVolumeMonitorDriveChangedDetail{callback, handlerID}
	signalVolumeMonitorDriveChangedMap[signalVolumeMonitorDriveChangedId] = detail

	return signalVolumeMonitorDriveChangedId
}

/*
DisconnectDriveChanged disconnects a callback from the 'drive-changed' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveChanged.
*/
func (recv *VolumeMonitor) DisconnectDriveChanged(connectionID int) {
	signalVolumeMonitorDriveChangedLock.Lock()
	defer signalVolumeMonitorDriveChangedLock.Unlock()

	detail, exists := signalVolumeMonitorDriveChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveChangedMap, connectionID)
}

//export volumemonitor_driveChangedHandler
func volumemonitor_driveChangedHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveChangedLock.RLock()
	defer signalVolumeMonitorDriveChangedLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveChangedMap[index].callback
	callback(drive)
}

type signalVolumeMonitorDriveConnectedDetail struct {
	callback  VolumeMonitorSignalDriveConnectedCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveConnectedId int
var signalVolumeMonitorDriveConnectedMap = make(map[int]signalVolumeMonitorDriveConnectedDetail)
var signalVolumeMonitorDriveConnectedLock sync.RWMutex

// VolumeMonitorSignalDriveConnectedCallback is a callback function for a 'drive-connected' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveConnectedCallback func(drive *Drive)

/*
ConnectDriveConnected connects the callback to the 'drive-connected' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveConnected to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveConnected(callback VolumeMonitorSignalDriveConnectedCallback) int {
	signalVolumeMonitorDriveConnectedLock.Lock()
	defer signalVolumeMonitorDriveConnectedLock.Unlock()

	signalVolumeMonitorDriveConnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_connected(instance, C.gpointer(uintptr(signalVolumeMonitorDriveConnectedId)))

	detail := signalVolumeMonitorDriveConnectedDetail{callback, handlerID}
	signalVolumeMonitorDriveConnectedMap[signalVolumeMonitorDriveConnectedId] = detail

	return signalVolumeMonitorDriveConnectedId
}

/*
DisconnectDriveConnected disconnects a callback from the 'drive-connected' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveConnected.
*/
func (recv *VolumeMonitor) DisconnectDriveConnected(connectionID int) {
	signalVolumeMonitorDriveConnectedLock.Lock()
	defer signalVolumeMonitorDriveConnectedLock.Unlock()

	detail, exists := signalVolumeMonitorDriveConnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveConnectedMap, connectionID)
}

//export volumemonitor_driveConnectedHandler
func volumemonitor_driveConnectedHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveConnectedLock.RLock()
	defer signalVolumeMonitorDriveConnectedLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveConnectedMap[index].callback
	callback(drive)
}

type signalVolumeMonitorDriveDisconnectedDetail struct {
	callback  VolumeMonitorSignalDriveDisconnectedCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveDisconnectedId int
var signalVolumeMonitorDriveDisconnectedMap = make(map[int]signalVolumeMonitorDriveDisconnectedDetail)
var signalVolumeMonitorDriveDisconnectedLock sync.RWMutex

// VolumeMonitorSignalDriveDisconnectedCallback is a callback function for a 'drive-disconnected' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveDisconnectedCallback func(drive *Drive)

/*
ConnectDriveDisconnected connects the callback to the 'drive-disconnected' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveDisconnected to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveDisconnected(callback VolumeMonitorSignalDriveDisconnectedCallback) int {
	signalVolumeMonitorDriveDisconnectedLock.Lock()
	defer signalVolumeMonitorDriveDisconnectedLock.Unlock()

	signalVolumeMonitorDriveDisconnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_disconnected(instance, C.gpointer(uintptr(signalVolumeMonitorDriveDisconnectedId)))

	detail := signalVolumeMonitorDriveDisconnectedDetail{callback, handlerID}
	signalVolumeMonitorDriveDisconnectedMap[signalVolumeMonitorDriveDisconnectedId] = detail

	return signalVolumeMonitorDriveDisconnectedId
}

/*
DisconnectDriveDisconnected disconnects a callback from the 'drive-disconnected' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveDisconnected.
*/
func (recv *VolumeMonitor) DisconnectDriveDisconnected(connectionID int) {
	signalVolumeMonitorDriveDisconnectedLock.Lock()
	defer signalVolumeMonitorDriveDisconnectedLock.Unlock()

	detail, exists := signalVolumeMonitorDriveDisconnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveDisconnectedMap, connectionID)
}

//export volumemonitor_driveDisconnectedHandler
func volumemonitor_driveDisconnectedHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveDisconnectedLock.RLock()
	defer signalVolumeMonitorDriveDisconnectedLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveDisconnectedMap[index].callback
	callback(drive)
}

type signalVolumeMonitorMountAddedDetail struct {
	callback  VolumeMonitorSignalMountAddedCallback
	handlerID C.gulong
}

var signalVolumeMonitorMountAddedId int
var signalVolumeMonitorMountAddedMap = make(map[int]signalVolumeMonitorMountAddedDetail)
var signalVolumeMonitorMountAddedLock sync.RWMutex

// VolumeMonitorSignalMountAddedCallback is a callback function for a 'mount-added' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalMountAddedCallback func(mount *Mount)

/*
ConnectMountAdded connects the callback to the 'mount-added' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectMountAdded to remove it.
*/
func (recv *VolumeMonitor) ConnectMountAdded(callback VolumeMonitorSignalMountAddedCallback) int {
	signalVolumeMonitorMountAddedLock.Lock()
	defer signalVolumeMonitorMountAddedLock.Unlock()

	signalVolumeMonitorMountAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_mount_added(instance, C.gpointer(uintptr(signalVolumeMonitorMountAddedId)))

	detail := signalVolumeMonitorMountAddedDetail{callback, handlerID}
	signalVolumeMonitorMountAddedMap[signalVolumeMonitorMountAddedId] = detail

	return signalVolumeMonitorMountAddedId
}

/*
DisconnectMountAdded disconnects a callback from the 'mount-added' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectMountAdded.
*/
func (recv *VolumeMonitor) DisconnectMountAdded(connectionID int) {
	signalVolumeMonitorMountAddedLock.Lock()
	defer signalVolumeMonitorMountAddedLock.Unlock()

	detail, exists := signalVolumeMonitorMountAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorMountAddedMap, connectionID)
}

//export volumemonitor_mountAddedHandler
func volumemonitor_mountAddedHandler(_ *C.GObject, c_mount *C.GMount, data C.gpointer) {
	signalVolumeMonitorMountAddedLock.RLock()
	defer signalVolumeMonitorMountAddedLock.RUnlock()

	mount := MountNewFromC(unsafe.Pointer(c_mount))

	index := int(uintptr(data))
	callback := signalVolumeMonitorMountAddedMap[index].callback
	callback(mount)
}

type signalVolumeMonitorMountChangedDetail struct {
	callback  VolumeMonitorSignalMountChangedCallback
	handlerID C.gulong
}

var signalVolumeMonitorMountChangedId int
var signalVolumeMonitorMountChangedMap = make(map[int]signalVolumeMonitorMountChangedDetail)
var signalVolumeMonitorMountChangedLock sync.RWMutex

// VolumeMonitorSignalMountChangedCallback is a callback function for a 'mount-changed' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalMountChangedCallback func(mount *Mount)

/*
ConnectMountChanged connects the callback to the 'mount-changed' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectMountChanged to remove it.
*/
func (recv *VolumeMonitor) ConnectMountChanged(callback VolumeMonitorSignalMountChangedCallback) int {
	signalVolumeMonitorMountChangedLock.Lock()
	defer signalVolumeMonitorMountChangedLock.Unlock()

	signalVolumeMonitorMountChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_mount_changed(instance, C.gpointer(uintptr(signalVolumeMonitorMountChangedId)))

	detail := signalVolumeMonitorMountChangedDetail{callback, handlerID}
	signalVolumeMonitorMountChangedMap[signalVolumeMonitorMountChangedId] = detail

	return signalVolumeMonitorMountChangedId
}

/*
DisconnectMountChanged disconnects a callback from the 'mount-changed' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectMountChanged.
*/
func (recv *VolumeMonitor) DisconnectMountChanged(connectionID int) {
	signalVolumeMonitorMountChangedLock.Lock()
	defer signalVolumeMonitorMountChangedLock.Unlock()

	detail, exists := signalVolumeMonitorMountChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorMountChangedMap, connectionID)
}

//export volumemonitor_mountChangedHandler
func volumemonitor_mountChangedHandler(_ *C.GObject, c_mount *C.GMount, data C.gpointer) {
	signalVolumeMonitorMountChangedLock.RLock()
	defer signalVolumeMonitorMountChangedLock.RUnlock()

	mount := MountNewFromC(unsafe.Pointer(c_mount))

	index := int(uintptr(data))
	callback := signalVolumeMonitorMountChangedMap[index].callback
	callback(mount)
}

type signalVolumeMonitorMountPreUnmountDetail struct {
	callback  VolumeMonitorSignalMountPreUnmountCallback
	handlerID C.gulong
}

var signalVolumeMonitorMountPreUnmountId int
var signalVolumeMonitorMountPreUnmountMap = make(map[int]signalVolumeMonitorMountPreUnmountDetail)
var signalVolumeMonitorMountPreUnmountLock sync.RWMutex

// VolumeMonitorSignalMountPreUnmountCallback is a callback function for a 'mount-pre-unmount' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalMountPreUnmountCallback func(mount *Mount)

/*
ConnectMountPreUnmount connects the callback to the 'mount-pre-unmount' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectMountPreUnmount to remove it.
*/
func (recv *VolumeMonitor) ConnectMountPreUnmount(callback VolumeMonitorSignalMountPreUnmountCallback) int {
	signalVolumeMonitorMountPreUnmountLock.Lock()
	defer signalVolumeMonitorMountPreUnmountLock.Unlock()

	signalVolumeMonitorMountPreUnmountId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_mount_pre_unmount(instance, C.gpointer(uintptr(signalVolumeMonitorMountPreUnmountId)))

	detail := signalVolumeMonitorMountPreUnmountDetail{callback, handlerID}
	signalVolumeMonitorMountPreUnmountMap[signalVolumeMonitorMountPreUnmountId] = detail

	return signalVolumeMonitorMountPreUnmountId
}

/*
DisconnectMountPreUnmount disconnects a callback from the 'mount-pre-unmount' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectMountPreUnmount.
*/
func (recv *VolumeMonitor) DisconnectMountPreUnmount(connectionID int) {
	signalVolumeMonitorMountPreUnmountLock.Lock()
	defer signalVolumeMonitorMountPreUnmountLock.Unlock()

	detail, exists := signalVolumeMonitorMountPreUnmountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorMountPreUnmountMap, connectionID)
}

//export volumemonitor_mountPreUnmountHandler
func volumemonitor_mountPreUnmountHandler(_ *C.GObject, c_mount *C.GMount, data C.gpointer) {
	signalVolumeMonitorMountPreUnmountLock.RLock()
	defer signalVolumeMonitorMountPreUnmountLock.RUnlock()

	mount := MountNewFromC(unsafe.Pointer(c_mount))

	index := int(uintptr(data))
	callback := signalVolumeMonitorMountPreUnmountMap[index].callback
	callback(mount)
}

type signalVolumeMonitorMountRemovedDetail struct {
	callback  VolumeMonitorSignalMountRemovedCallback
	handlerID C.gulong
}

var signalVolumeMonitorMountRemovedId int
var signalVolumeMonitorMountRemovedMap = make(map[int]signalVolumeMonitorMountRemovedDetail)
var signalVolumeMonitorMountRemovedLock sync.RWMutex

// VolumeMonitorSignalMountRemovedCallback is a callback function for a 'mount-removed' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalMountRemovedCallback func(mount *Mount)

/*
ConnectMountRemoved connects the callback to the 'mount-removed' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectMountRemoved to remove it.
*/
func (recv *VolumeMonitor) ConnectMountRemoved(callback VolumeMonitorSignalMountRemovedCallback) int {
	signalVolumeMonitorMountRemovedLock.Lock()
	defer signalVolumeMonitorMountRemovedLock.Unlock()

	signalVolumeMonitorMountRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_mount_removed(instance, C.gpointer(uintptr(signalVolumeMonitorMountRemovedId)))

	detail := signalVolumeMonitorMountRemovedDetail{callback, handlerID}
	signalVolumeMonitorMountRemovedMap[signalVolumeMonitorMountRemovedId] = detail

	return signalVolumeMonitorMountRemovedId
}

/*
DisconnectMountRemoved disconnects a callback from the 'mount-removed' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectMountRemoved.
*/
func (recv *VolumeMonitor) DisconnectMountRemoved(connectionID int) {
	signalVolumeMonitorMountRemovedLock.Lock()
	defer signalVolumeMonitorMountRemovedLock.Unlock()

	detail, exists := signalVolumeMonitorMountRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorMountRemovedMap, connectionID)
}

//export volumemonitor_mountRemovedHandler
func volumemonitor_mountRemovedHandler(_ *C.GObject, c_mount *C.GMount, data C.gpointer) {
	signalVolumeMonitorMountRemovedLock.RLock()
	defer signalVolumeMonitorMountRemovedLock.RUnlock()

	mount := MountNewFromC(unsafe.Pointer(c_mount))

	index := int(uintptr(data))
	callback := signalVolumeMonitorMountRemovedMap[index].callback
	callback(mount)
}

type signalVolumeMonitorVolumeAddedDetail struct {
	callback  VolumeMonitorSignalVolumeAddedCallback
	handlerID C.gulong
}

var signalVolumeMonitorVolumeAddedId int
var signalVolumeMonitorVolumeAddedMap = make(map[int]signalVolumeMonitorVolumeAddedDetail)
var signalVolumeMonitorVolumeAddedLock sync.RWMutex

// VolumeMonitorSignalVolumeAddedCallback is a callback function for a 'volume-added' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalVolumeAddedCallback func(volume *Volume)

/*
ConnectVolumeAdded connects the callback to the 'volume-added' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectVolumeAdded to remove it.
*/
func (recv *VolumeMonitor) ConnectVolumeAdded(callback VolumeMonitorSignalVolumeAddedCallback) int {
	signalVolumeMonitorVolumeAddedLock.Lock()
	defer signalVolumeMonitorVolumeAddedLock.Unlock()

	signalVolumeMonitorVolumeAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_volume_added(instance, C.gpointer(uintptr(signalVolumeMonitorVolumeAddedId)))

	detail := signalVolumeMonitorVolumeAddedDetail{callback, handlerID}
	signalVolumeMonitorVolumeAddedMap[signalVolumeMonitorVolumeAddedId] = detail

	return signalVolumeMonitorVolumeAddedId
}

/*
DisconnectVolumeAdded disconnects a callback from the 'volume-added' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectVolumeAdded.
*/
func (recv *VolumeMonitor) DisconnectVolumeAdded(connectionID int) {
	signalVolumeMonitorVolumeAddedLock.Lock()
	defer signalVolumeMonitorVolumeAddedLock.Unlock()

	detail, exists := signalVolumeMonitorVolumeAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorVolumeAddedMap, connectionID)
}

//export volumemonitor_volumeAddedHandler
func volumemonitor_volumeAddedHandler(_ *C.GObject, c_volume *C.GVolume, data C.gpointer) {
	signalVolumeMonitorVolumeAddedLock.RLock()
	defer signalVolumeMonitorVolumeAddedLock.RUnlock()

	volume := VolumeNewFromC(unsafe.Pointer(c_volume))

	index := int(uintptr(data))
	callback := signalVolumeMonitorVolumeAddedMap[index].callback
	callback(volume)
}

type signalVolumeMonitorVolumeChangedDetail struct {
	callback  VolumeMonitorSignalVolumeChangedCallback
	handlerID C.gulong
}

var signalVolumeMonitorVolumeChangedId int
var signalVolumeMonitorVolumeChangedMap = make(map[int]signalVolumeMonitorVolumeChangedDetail)
var signalVolumeMonitorVolumeChangedLock sync.RWMutex

// VolumeMonitorSignalVolumeChangedCallback is a callback function for a 'volume-changed' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalVolumeChangedCallback func(volume *Volume)

/*
ConnectVolumeChanged connects the callback to the 'volume-changed' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectVolumeChanged to remove it.
*/
func (recv *VolumeMonitor) ConnectVolumeChanged(callback VolumeMonitorSignalVolumeChangedCallback) int {
	signalVolumeMonitorVolumeChangedLock.Lock()
	defer signalVolumeMonitorVolumeChangedLock.Unlock()

	signalVolumeMonitorVolumeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_volume_changed(instance, C.gpointer(uintptr(signalVolumeMonitorVolumeChangedId)))

	detail := signalVolumeMonitorVolumeChangedDetail{callback, handlerID}
	signalVolumeMonitorVolumeChangedMap[signalVolumeMonitorVolumeChangedId] = detail

	return signalVolumeMonitorVolumeChangedId
}

/*
DisconnectVolumeChanged disconnects a callback from the 'volume-changed' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectVolumeChanged.
*/
func (recv *VolumeMonitor) DisconnectVolumeChanged(connectionID int) {
	signalVolumeMonitorVolumeChangedLock.Lock()
	defer signalVolumeMonitorVolumeChangedLock.Unlock()

	detail, exists := signalVolumeMonitorVolumeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorVolumeChangedMap, connectionID)
}

//export volumemonitor_volumeChangedHandler
func volumemonitor_volumeChangedHandler(_ *C.GObject, c_volume *C.GVolume, data C.gpointer) {
	signalVolumeMonitorVolumeChangedLock.RLock()
	defer signalVolumeMonitorVolumeChangedLock.RUnlock()

	volume := VolumeNewFromC(unsafe.Pointer(c_volume))

	index := int(uintptr(data))
	callback := signalVolumeMonitorVolumeChangedMap[index].callback
	callback(volume)
}

type signalVolumeMonitorVolumeRemovedDetail struct {
	callback  VolumeMonitorSignalVolumeRemovedCallback
	handlerID C.gulong
}

var signalVolumeMonitorVolumeRemovedId int
var signalVolumeMonitorVolumeRemovedMap = make(map[int]signalVolumeMonitorVolumeRemovedDetail)
var signalVolumeMonitorVolumeRemovedLock sync.RWMutex

// VolumeMonitorSignalVolumeRemovedCallback is a callback function for a 'volume-removed' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalVolumeRemovedCallback func(volume *Volume)

/*
ConnectVolumeRemoved connects the callback to the 'volume-removed' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectVolumeRemoved to remove it.
*/
func (recv *VolumeMonitor) ConnectVolumeRemoved(callback VolumeMonitorSignalVolumeRemovedCallback) int {
	signalVolumeMonitorVolumeRemovedLock.Lock()
	defer signalVolumeMonitorVolumeRemovedLock.Unlock()

	signalVolumeMonitorVolumeRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_volume_removed(instance, C.gpointer(uintptr(signalVolumeMonitorVolumeRemovedId)))

	detail := signalVolumeMonitorVolumeRemovedDetail{callback, handlerID}
	signalVolumeMonitorVolumeRemovedMap[signalVolumeMonitorVolumeRemovedId] = detail

	return signalVolumeMonitorVolumeRemovedId
}

/*
DisconnectVolumeRemoved disconnects a callback from the 'volume-removed' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectVolumeRemoved.
*/
func (recv *VolumeMonitor) DisconnectVolumeRemoved(connectionID int) {
	signalVolumeMonitorVolumeRemovedLock.Lock()
	defer signalVolumeMonitorVolumeRemovedLock.Unlock()

	detail, exists := signalVolumeMonitorVolumeRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorVolumeRemovedMap, connectionID)
}

//export volumemonitor_volumeRemovedHandler
func volumemonitor_volumeRemovedHandler(_ *C.GObject, c_volume *C.GVolume, data C.gpointer) {
	signalVolumeMonitorVolumeRemovedLock.RLock()
	defer signalVolumeMonitorVolumeRemovedLock.RUnlock()

	volume := VolumeNewFromC(unsafe.Pointer(c_volume))

	index := int(uintptr(data))
	callback := signalVolumeMonitorVolumeRemovedMap[index].callback
	callback(volume)
}

// VolumeMonitorAdoptOrphanMount is a wrapper around the C function g_volume_monitor_adopt_orphan_mount.
func VolumeMonitorAdoptOrphanMount(mount *Mount) *Volume {
	c_mount := (*C.GMount)(mount.ToC())

	retC := C.g_volume_monitor_adopt_orphan_mount(c_mount)
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VolumeMonitorGet is a wrapper around the C function g_volume_monitor_get.
func VolumeMonitorGet() *VolumeMonitor {
	retC := C.g_volume_monitor_get()
	retGo := VolumeMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetConnectedDrives is a wrapper around the C function g_volume_monitor_get_connected_drives.
func (recv *VolumeMonitor) GetConnectedDrives() *glib.List {
	retC := C.g_volume_monitor_get_connected_drives((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMountForUuid is a wrapper around the C function g_volume_monitor_get_mount_for_uuid.
func (recv *VolumeMonitor) GetMountForUuid(uuid string) *Mount {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	retC := C.g_volume_monitor_get_mount_for_uuid((*C.GVolumeMonitor)(recv.native), c_uuid)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMounts is a wrapper around the C function g_volume_monitor_get_mounts.
func (recv *VolumeMonitor) GetMounts() *glib.List {
	retC := C.g_volume_monitor_get_mounts((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVolumeForUuid is a wrapper around the C function g_volume_monitor_get_volume_for_uuid.
func (recv *VolumeMonitor) GetVolumeForUuid(uuid string) *Volume {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	retC := C.g_volume_monitor_get_volume_for_uuid((*C.GVolumeMonitor)(recv.native), c_uuid)
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVolumes is a wrapper around the C function g_volume_monitor_get_volumes.
func (recv *VolumeMonitor) GetVolumes() *glib.List {
	retC := C.g_volume_monitor_get_volumes((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ZlibCompressor is a wrapper around the C record GZlibCompressor.
type ZlibCompressor struct {
	native *C.GZlibCompressor
}

func ZlibCompressorNewFromC(u unsafe.Pointer) *ZlibCompressor {
	c := (*C.GZlibCompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ZlibCompressor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ZlibCompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ZlibCompressor with another ZlibCompressor, and returns true if they represent the same GObject.
func (recv *ZlibCompressor) Equals(other *ZlibCompressor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ZlibCompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ZlibCompressor.
// Exercise care, as this is a potentially dangerous function if the Object is not a ZlibCompressor.
func CastToZlibCompressor(object *gobject.Object) *ZlibCompressor {
	return ZlibCompressorNewFromC(object.ToC())
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// ZlibDecompressor is a wrapper around the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native *C.GZlibDecompressor
}

func ZlibDecompressorNewFromC(u unsafe.Pointer) *ZlibDecompressor {
	c := (*C.GZlibDecompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ZlibDecompressor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ZlibDecompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ZlibDecompressor with another ZlibDecompressor, and returns true if they represent the same GObject.
func (recv *ZlibDecompressor) Equals(other *ZlibDecompressor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ZlibDecompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ZlibDecompressor.
// Exercise care, as this is a potentially dangerous function if the Object is not a ZlibDecompressor.
func CastToZlibDecompressor(object *gobject.Object) *ZlibDecompressor {
	return ZlibDecompressorNewFromC(object.ToC())
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME string = C.G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME
const FILE_ATTRIBUTE_ACCESS_CAN_DELETE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE
const FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE
const FILE_ATTRIBUTE_ACCESS_CAN_READ string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_READ
const FILE_ATTRIBUTE_ACCESS_CAN_RENAME string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME
const FILE_ATTRIBUTE_ACCESS_CAN_TRASH string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH
const FILE_ATTRIBUTE_ACCESS_CAN_WRITE string = C.G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE
const FILE_ATTRIBUTE_DOS_IS_ARCHIVE string = C.G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE
const FILE_ATTRIBUTE_DOS_IS_SYSTEM string = C.G_FILE_ATTRIBUTE_DOS_IS_SYSTEM
const FILE_ATTRIBUTE_ETAG_VALUE string = C.G_FILE_ATTRIBUTE_ETAG_VALUE
const FILE_ATTRIBUTE_FILESYSTEM_FREE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_FREE
const FILE_ATTRIBUTE_FILESYSTEM_READONLY string = C.G_FILE_ATTRIBUTE_FILESYSTEM_READONLY

// Blacklisted : FILE_ATTRIBUTE_FILESYSTEM_REMOTE

const FILE_ATTRIBUTE_FILESYSTEM_SIZE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_SIZE
const FILE_ATTRIBUTE_FILESYSTEM_TYPE string = C.G_FILE_ATTRIBUTE_FILESYSTEM_TYPE
const FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW string = C.G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW
const FILE_ATTRIBUTE_GVFS_BACKEND string = C.G_FILE_ATTRIBUTE_GVFS_BACKEND
const FILE_ATTRIBUTE_ID_FILE string = C.G_FILE_ATTRIBUTE_ID_FILE
const FILE_ATTRIBUTE_ID_FILESYSTEM string = C.G_FILE_ATTRIBUTE_ID_FILESYSTEM
const FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT
const FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI string = C.G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE
const FILE_ATTRIBUTE_OWNER_GROUP string = C.G_FILE_ATTRIBUTE_OWNER_GROUP
const FILE_ATTRIBUTE_OWNER_USER string = C.G_FILE_ATTRIBUTE_OWNER_USER
const FILE_ATTRIBUTE_OWNER_USER_REAL string = C.G_FILE_ATTRIBUTE_OWNER_USER_REAL
const FILE_ATTRIBUTE_SELINUX_CONTEXT string = C.G_FILE_ATTRIBUTE_SELINUX_CONTEXT
const FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE
const FILE_ATTRIBUTE_STANDARD_COPY_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_COPY_NAME
const FILE_ATTRIBUTE_STANDARD_DESCRIPTION string = C.G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION
const FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME
const FILE_ATTRIBUTE_STANDARD_EDIT_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME
const FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE
const FILE_ATTRIBUTE_STANDARD_ICON string = C.G_FILE_ATTRIBUTE_STANDARD_ICON
const FILE_ATTRIBUTE_STANDARD_IS_BACKUP string = C.G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP
const FILE_ATTRIBUTE_STANDARD_IS_HIDDEN string = C.G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN
const FILE_ATTRIBUTE_STANDARD_IS_SYMLINK string = C.G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK
const FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL string = C.G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL
const FILE_ATTRIBUTE_STANDARD_NAME string = C.G_FILE_ATTRIBUTE_STANDARD_NAME
const FILE_ATTRIBUTE_STANDARD_SIZE string = C.G_FILE_ATTRIBUTE_STANDARD_SIZE
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER string = C.G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET string = C.G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET
const FILE_ATTRIBUTE_STANDARD_TARGET_URI string = C.G_FILE_ATTRIBUTE_STANDARD_TARGET_URI
const FILE_ATTRIBUTE_STANDARD_TYPE string = C.G_FILE_ATTRIBUTE_STANDARD_TYPE
const FILE_ATTRIBUTE_THUMBNAILING_FAILED string = C.G_FILE_ATTRIBUTE_THUMBNAILING_FAILED
const FILE_ATTRIBUTE_THUMBNAIL_PATH string = C.G_FILE_ATTRIBUTE_THUMBNAIL_PATH
const FILE_ATTRIBUTE_TIME_ACCESS string = C.G_FILE_ATTRIBUTE_TIME_ACCESS
const FILE_ATTRIBUTE_TIME_ACCESS_USEC string = C.G_FILE_ATTRIBUTE_TIME_ACCESS_USEC
const FILE_ATTRIBUTE_TIME_CHANGED string = C.G_FILE_ATTRIBUTE_TIME_CHANGED
const FILE_ATTRIBUTE_TIME_CHANGED_USEC string = C.G_FILE_ATTRIBUTE_TIME_CHANGED_USEC
const FILE_ATTRIBUTE_TIME_CREATED string = C.G_FILE_ATTRIBUTE_TIME_CREATED
const FILE_ATTRIBUTE_TIME_CREATED_USEC string = C.G_FILE_ATTRIBUTE_TIME_CREATED_USEC
const FILE_ATTRIBUTE_TIME_MODIFIED string = C.G_FILE_ATTRIBUTE_TIME_MODIFIED
const FILE_ATTRIBUTE_TIME_MODIFIED_USEC string = C.G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT string = C.G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT
const FILE_ATTRIBUTE_UNIX_BLOCKS string = C.G_FILE_ATTRIBUTE_UNIX_BLOCKS
const FILE_ATTRIBUTE_UNIX_BLOCK_SIZE string = C.G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE
const FILE_ATTRIBUTE_UNIX_DEVICE string = C.G_FILE_ATTRIBUTE_UNIX_DEVICE
const FILE_ATTRIBUTE_UNIX_GID string = C.G_FILE_ATTRIBUTE_UNIX_GID
const FILE_ATTRIBUTE_UNIX_INODE string = C.G_FILE_ATTRIBUTE_UNIX_INODE
const FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT string = C.G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT
const FILE_ATTRIBUTE_UNIX_MODE string = C.G_FILE_ATTRIBUTE_UNIX_MODE
const FILE_ATTRIBUTE_UNIX_NLINK string = C.G_FILE_ATTRIBUTE_UNIX_NLINK
const FILE_ATTRIBUTE_UNIX_RDEV string = C.G_FILE_ATTRIBUTE_UNIX_RDEV
const FILE_ATTRIBUTE_UNIX_UID string = C.G_FILE_ATTRIBUTE_UNIX_UID
const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME string = C.G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME
const PROXY_RESOLVER_EXTENSION_POINT_NAME string = C.G_PROXY_RESOLVER_EXTENSION_POINT_NAME

// Blacklisted : SETTINGS_BACKEND_EXTENSION_POINT_NAME

const TLS_BACKEND_EXTENSION_POINT_NAME string = C.G_TLS_BACKEND_EXTENSION_POINT_NAME
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT string = C.G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER string = C.G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER
const VFS_EXTENSION_POINT_NAME string = C.G_VFS_EXTENSION_POINT_NAME
const VOLUME_IDENTIFIER_KIND_CLASS string = C.G_VOLUME_IDENTIFIER_KIND_CLASS
const VOLUME_IDENTIFIER_KIND_HAL_UDI string = C.G_VOLUME_IDENTIFIER_KIND_HAL_UDI
const VOLUME_IDENTIFIER_KIND_LABEL string = C.G_VOLUME_IDENTIFIER_KIND_LABEL
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT string = C.G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE string = C.G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE
const VOLUME_IDENTIFIER_KIND_UUID string = C.G_VOLUME_IDENTIFIER_KIND_UUID
const VOLUME_MONITOR_EXTENSION_POINT_NAME string = C.G_VOLUME_MONITOR_EXTENSION_POINT_NAME

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

// ContentTypeCanBeExecutable is a wrapper around the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_can_be_executable(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeEquals is a wrapper around the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	c_type1 := C.CString(type1)
	defer C.free(unsafe.Pointer(c_type1))

	c_type2 := C.CString(type2)
	defer C.free(unsafe.Pointer(c_type2))

	retC := C.g_content_type_equals(c_type1, c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeGetDescription is a wrapper around the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_description(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGetIcon is a wrapper around the C function g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) *Icon {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_icon(c_type)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGetMimeType is a wrapper around the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_mime_type(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGuess is a wrapper around the C function g_content_type_guess.
func ContentTypeGuess(filename string, data []uint8) (string, bool) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_data_size := (C.gsize)(len(data))

	var c_result_uncertain C.gboolean

	retC := C.g_content_type_guess(c_filename, c_data, c_data_size, &c_result_uncertain)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	resultUncertain := c_result_uncertain == C.TRUE

	return retGo, resultUncertain
}

// ContentTypeIsA is a wrapper around the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_supertype := C.CString(supertype)
	defer C.free(unsafe.Pointer(c_supertype))

	retC := C.g_content_type_is_a(c_type, c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypeIsUnknown is a wrapper around the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_is_unknown(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// ContentTypesGetRegistered is a wrapper around the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retC := C.g_content_types_get_registered()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon

// IoErrorFromErrno is a wrapper around the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	c_err_no := (C.gint)(errNo)

	retC := C.g_io_error_from_errno(c_err_no)
	retGo := (IOErrorEnum)(retC)

	return retGo
}

// IoErrorQuark is a wrapper around the C function g_io_error_quark.
func IoErrorQuark() glib.Quark {
	retC := C.g_io_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// IoModulesLoadAllInDirectory is a wrapper around the C function g_io_modules_load_all_in_directory.
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	retC := C.g_io_modules_load_all_in_directory(c_dirname)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IoSchedulerCancelAllJobs is a wrapper around the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()

	return
}

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Blacklisted : g_keyfile_settings_backend_new

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnixIsMountPathSystemInternal is a wrapper around the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	retC := C.g_unix_is_mount_path_system_internal(c_mount_path)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountAt is a wrapper around the C function g_unix_mount_at.
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_at(c_mount_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountCompare is a wrapper around the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	c_mount1 := (*C.GUnixMountEntry)(C.NULL)
	if mount1 != nil {
		c_mount1 = (*C.GUnixMountEntry)(mount1.ToC())
	}

	c_mount2 := (*C.GUnixMountEntry)(C.NULL)
	if mount2 != nil {
		c_mount2 = (*C.GUnixMountEntry)(mount2.ToC())
	}

	retC := C.g_unix_mount_compare(c_mount1, c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// UnixMountFree is a wrapper around the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	C.g_unix_mount_free(c_mount_entry)

	return
}

// UnixMountGetDevicePath is a wrapper around the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_device_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetFsType is a wrapper around the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_fs_type(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGetMountPath is a wrapper around the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_mount_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// UnixMountGuessCanEject is a wrapper around the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_can_eject(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountGuessIcon is a wrapper around the C function g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_icon(c_mount_entry)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnixMountGuessName is a wrapper around the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_name(c_mount_entry)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UnixMountGuessShouldDisplay is a wrapper around the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_should_display(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountIsReadonly is a wrapper around the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_readonly(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountIsSystemInternal is a wrapper around the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_system_internal(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsChangedSince is a wrapper around the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mount_points_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountPointsGet is a wrapper around the C function g_unix_mount_points_get.
func UnixMountPointsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mount_points_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// UnixMountsChangedSince is a wrapper around the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mounts_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// UnixMountsGet is a wrapper around the C function g_unix_mounts_get.
func UnixMountsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mounts_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// Action is a wrapper around the C record GAction.
type Action struct {
	native *C.GAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// ActionGroup is a wrapper around the C record GActionGroup.
type ActionGroup struct {
	native *C.GActionGroup
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionGroup with another ActionGroup, and returns true if they represent the same GObject.
func (recv *ActionGroup) Equals(other *ActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// ActionMap is a wrapper around the C record GActionMap.
type ActionMap struct {
	native *C.GActionMap
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	c := (*C.GActionMap)(u)
	if c == nil {
		return nil
	}

	g := &ActionMap{native: c}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionMap with another ActionMap, and returns true if they represent the same GObject.
func (recv *ActionMap) Equals(other *ActionMap) bool {
	return other.ToC() == recv.ToC()
}

// AppInfo is a wrapper around the C record GAppInfo.
type AppInfo struct {
	native *C.GAppInfo
}

func AppInfoNewFromC(u unsafe.Pointer) *AppInfo {
	c := (*C.GAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &AppInfo{native: c}

	return g
}

func (recv *AppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppInfo with another AppInfo, and returns true if they represent the same GObject.
func (recv *AppInfo) Equals(other *AppInfo) bool {
	return other.ToC() == recv.ToC()
}

// AppInfoCreateFromCommandline is a wrapper around the C function g_app_info_create_from_commandline.
func AppInfoCreateFromCommandline(commandline string, applicationName string, flags AppInfoCreateFlags) (*AppInfo, error) {
	c_commandline := C.CString(commandline)
	defer C.free(unsafe.Pointer(c_commandline))

	c_application_name := C.CString(applicationName)
	defer C.free(unsafe.Pointer(c_application_name))

	c_flags := (C.GAppInfoCreateFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_app_info_create_from_commandline(c_commandline, c_application_name, c_flags, &cThrowableError)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AppInfoGetAll is a wrapper around the C function g_app_info_get_all.
func AppInfoGetAll() *glib.List {
	retC := C.g_app_info_get_all()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetAllForType is a wrapper around the C function g_app_info_get_all_for_type.
func AppInfoGetAllForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_all_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetDefaultForType is a wrapper around the C function g_app_info_get_default_for_type.
func AppInfoGetDefaultForType(contentType string, mustSupportUris bool) *AppInfo {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_must_support_uris :=
		boolToGboolean(mustSupportUris)

	retC := C.g_app_info_get_default_for_type(c_content_type, c_must_support_uris)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetDefaultForUriScheme is a wrapper around the C function g_app_info_get_default_for_uri_scheme.
func AppInfoGetDefaultForUriScheme(uriScheme string) *AppInfo {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_app_info_get_default_for_uri_scheme(c_uri_scheme)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoLaunchDefaultForUri is a wrapper around the C function g_app_info_launch_default_for_uri.
func AppInfoLaunchDefaultForUri(uri string, context *AppLaunchContext) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_default_for_uri(c_uri, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddSupportsType is a wrapper around the C function g_app_info_add_supports_type.
func (recv *AppInfo) AddSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_add_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CanRemoveSupportsType is a wrapper around the C function g_app_info_can_remove_supports_type.
func (recv *AppInfo) CanRemoveSupportsType() bool {
	retC := C.g_app_info_can_remove_supports_type((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Dup is a wrapper around the C function g_app_info_dup.
func (recv *AppInfo) Dup() *AppInfo {
	retC := C.g_app_info_dup((*C.GAppInfo)(recv.native))
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function g_app_info_equal.
func (recv *AppInfo) Equal(appinfo2 *AppInfo) bool {
	c_appinfo2 := (*C.GAppInfo)(appinfo2.ToC())

	retC := C.g_app_info_equal((*C.GAppInfo)(recv.native), c_appinfo2)
	retGo := retC == C.TRUE

	return retGo
}

// GetDescription is a wrapper around the C function g_app_info_get_description.
func (recv *AppInfo) GetDescription() string {
	retC := C.g_app_info_get_description((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetExecutable is a wrapper around the C function g_app_info_get_executable.
func (recv *AppInfo) GetExecutable() string {
	retC := C.g_app_info_get_executable((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIcon is a wrapper around the C function g_app_info_get_icon.
func (recv *AppInfo) GetIcon() *Icon {
	retC := C.g_app_info_get_icon((*C.GAppInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetId is a wrapper around the C function g_app_info_get_id.
func (recv *AppInfo) GetId() string {
	retC := C.g_app_info_get_id((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function g_app_info_get_name.
func (recv *AppInfo) GetName() string {
	retC := C.g_app_info_get_name((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Launch is a wrapper around the C function g_app_info_launch.
func (recv *AppInfo) Launch(files *glib.List, context *AppLaunchContext) (bool, error) {
	c_files := (*C.GList)(C.NULL)
	if files != nil {
		c_files = (*C.GList)(files.ToC())
	}

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch((*C.GAppInfo)(recv.native), c_files, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LaunchUris is a wrapper around the C function g_app_info_launch_uris.
func (recv *AppInfo) LaunchUris(uris *glib.List, context *AppLaunchContext) (bool, error) {
	c_uris := (*C.GList)(C.NULL)
	if uris != nil {
		c_uris = (*C.GList)(uris.ToC())
	}

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_uris((*C.GAppInfo)(recv.native), c_uris, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveSupportsType is a wrapper around the C function g_app_info_remove_supports_type.
func (recv *AppInfo) RemoveSupportsType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_remove_supports_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAsDefaultForExtension is a wrapper around the C function g_app_info_set_as_default_for_extension.
func (recv *AppInfo) SetAsDefaultForExtension(extension string) (bool, error) {
	c_extension := C.CString(extension)
	defer C.free(unsafe.Pointer(c_extension))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_extension((*C.GAppInfo)(recv.native), c_extension, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAsDefaultForType is a wrapper around the C function g_app_info_set_as_default_for_type.
func (recv *AppInfo) SetAsDefaultForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_default_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAsLastUsedForType is a wrapper around the C function g_app_info_set_as_last_used_for_type.
func (recv *AppInfo) SetAsLastUsedForType(contentType string) (bool, error) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	var cThrowableError *C.GError

	retC := C.g_app_info_set_as_last_used_for_type((*C.GAppInfo)(recv.native), c_content_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ShouldShow is a wrapper around the C function g_app_info_should_show.
func (recv *AppInfo) ShouldShow() bool {
	retC := C.g_app_info_should_show((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsFiles is a wrapper around the C function g_app_info_supports_files.
func (recv *AppInfo) SupportsFiles() bool {
	retC := C.g_app_info_supports_files((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsUris is a wrapper around the C function g_app_info_supports_uris.
func (recv *AppInfo) SupportsUris() bool {
	retC := C.g_app_info_supports_uris((*C.GAppInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AsyncResult is a wrapper around the C record GAsyncResult.
type AsyncResult struct {
	native *C.GAsyncResult
}

func AsyncResultNewFromC(u unsafe.Pointer) *AsyncResult {
	c := (*C.GAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResult{native: c}

	return g
}

func (recv *AsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncResult with another AsyncResult, and returns true if they represent the same GObject.
func (recv *AsyncResult) Equals(other *AsyncResult) bool {
	return other.ToC() == recv.ToC()
}

// GetSourceObject is a wrapper around the C function g_async_result_get_source_object.
func (recv *AsyncResult) GetSourceObject() *gobject.Object {
	retC := C.g_async_result_get_source_object((*C.GAsyncResult)(recv.native))
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : g_async_result_get_user_data : no return generator

// DBusObject is a wrapper around the C record GDBusObject.
type DBusObject struct {
	native *C.GDBusObject
}

func DBusObjectNewFromC(u unsafe.Pointer) *DBusObject {
	c := (*C.GDBusObject)(u)
	if c == nil {
		return nil
	}

	g := &DBusObject{native: c}

	return g
}

func (recv *DBusObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObject with another DBusObject, and returns true if they represent the same GObject.
func (recv *DBusObject) Equals(other *DBusObject) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectManager is a wrapper around the C record GDBusObjectManager.
type DBusObjectManager struct {
	native *C.GDBusObjectManager
}

func DBusObjectManagerNewFromC(u unsafe.Pointer) *DBusObjectManager {
	c := (*C.GDBusObjectManager)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManager{native: c}

	return g
}

func (recv *DBusObjectManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManager with another DBusObjectManager, and returns true if they represent the same GObject.
func (recv *DBusObjectManager) Equals(other *DBusObjectManager) bool {
	return other.ToC() == recv.ToC()
}

// DesktopAppInfoLookup is a wrapper around the C record GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native *C.GDesktopAppInfoLookup
}

func DesktopAppInfoLookupNewFromC(u unsafe.Pointer) *DesktopAppInfoLookup {
	c := (*C.GDesktopAppInfoLookup)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookup{native: c}

	return g
}

func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DesktopAppInfoLookup with another DesktopAppInfoLookup, and returns true if they represent the same GObject.
func (recv *DesktopAppInfoLookup) Equals(other *DesktopAppInfoLookup) bool {
	return other.ToC() == recv.ToC()
}

// GetDefaultForUriScheme is a wrapper around the C function g_desktop_app_info_lookup_get_default_for_uri_scheme.
func (recv *DesktopAppInfoLookup) GetDefaultForUriScheme(uriScheme string) *AppInfo {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_desktop_app_info_lookup_get_default_for_uri_scheme((*C.GDesktopAppInfoLookup)(recv.native), c_uri_scheme)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Drive is a wrapper around the C record GDrive.
type Drive struct {
	native *C.GDrive
}

func DriveNewFromC(u unsafe.Pointer) *Drive {
	c := (*C.GDrive)(u)
	if c == nil {
		return nil
	}

	g := &Drive{native: c}

	return g
}

func (recv *Drive) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Drive with another Drive, and returns true if they represent the same GObject.
func (recv *Drive) Equals(other *Drive) bool {
	return other.ToC() == recv.ToC()
}

type signalDriveChangedDetail struct {
	callback  DriveSignalChangedCallback
	handlerID C.gulong
}

var signalDriveChangedId int
var signalDriveChangedMap = make(map[int]signalDriveChangedDetail)
var signalDriveChangedLock sync.RWMutex

// DriveSignalChangedCallback is a callback function for a 'changed' signal emitted from a Drive.
type DriveSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Drive) ConnectChanged(callback DriveSignalChangedCallback) int {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	signalDriveChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_changed(instance, C.gpointer(uintptr(signalDriveChangedId)))

	detail := signalDriveChangedDetail{callback, handlerID}
	signalDriveChangedMap[signalDriveChangedId] = detail

	return signalDriveChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Drive.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Drive) DisconnectChanged(connectionID int) {
	signalDriveChangedLock.Lock()
	defer signalDriveChangedLock.Unlock()

	detail, exists := signalDriveChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveChangedMap, connectionID)
}

//export drive_changedHandler
func drive_changedHandler(_ *C.GObject, data C.gpointer) {
	signalDriveChangedLock.RLock()
	defer signalDriveChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveChangedMap[index].callback
	callback()
}

type signalDriveDisconnectedDetail struct {
	callback  DriveSignalDisconnectedCallback
	handlerID C.gulong
}

var signalDriveDisconnectedId int
var signalDriveDisconnectedMap = make(map[int]signalDriveDisconnectedDetail)
var signalDriveDisconnectedLock sync.RWMutex

// DriveSignalDisconnectedCallback is a callback function for a 'disconnected' signal emitted from a Drive.
type DriveSignalDisconnectedCallback func()

/*
ConnectDisconnected connects the callback to the 'disconnected' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectDisconnected to remove it.
*/
func (recv *Drive) ConnectDisconnected(callback DriveSignalDisconnectedCallback) int {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	signalDriveDisconnectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_disconnected(instance, C.gpointer(uintptr(signalDriveDisconnectedId)))

	detail := signalDriveDisconnectedDetail{callback, handlerID}
	signalDriveDisconnectedMap[signalDriveDisconnectedId] = detail

	return signalDriveDisconnectedId
}

/*
DisconnectDisconnected disconnects a callback from the 'disconnected' signal for the Drive.

The connectionID should be a value returned from a call to ConnectDisconnected.
*/
func (recv *Drive) DisconnectDisconnected(connectionID int) {
	signalDriveDisconnectedLock.Lock()
	defer signalDriveDisconnectedLock.Unlock()

	detail, exists := signalDriveDisconnectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveDisconnectedMap, connectionID)
}

//export drive_disconnectedHandler
func drive_disconnectedHandler(_ *C.GObject, data C.gpointer) {
	signalDriveDisconnectedLock.RLock()
	defer signalDriveDisconnectedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveDisconnectedMap[index].callback
	callback()
}

type signalDriveEjectButtonDetail struct {
	callback  DriveSignalEjectButtonCallback
	handlerID C.gulong
}

var signalDriveEjectButtonId int
var signalDriveEjectButtonMap = make(map[int]signalDriveEjectButtonDetail)
var signalDriveEjectButtonLock sync.RWMutex

// DriveSignalEjectButtonCallback is a callback function for a 'eject-button' signal emitted from a Drive.
type DriveSignalEjectButtonCallback func()

/*
ConnectEjectButton connects the callback to the 'eject-button' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectEjectButton to remove it.
*/
func (recv *Drive) ConnectEjectButton(callback DriveSignalEjectButtonCallback) int {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	signalDriveEjectButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_eject_button(instance, C.gpointer(uintptr(signalDriveEjectButtonId)))

	detail := signalDriveEjectButtonDetail{callback, handlerID}
	signalDriveEjectButtonMap[signalDriveEjectButtonId] = detail

	return signalDriveEjectButtonId
}

/*
DisconnectEjectButton disconnects a callback from the 'eject-button' signal for the Drive.

The connectionID should be a value returned from a call to ConnectEjectButton.
*/
func (recv *Drive) DisconnectEjectButton(connectionID int) {
	signalDriveEjectButtonLock.Lock()
	defer signalDriveEjectButtonLock.Unlock()

	detail, exists := signalDriveEjectButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveEjectButtonMap, connectionID)
}

//export drive_ejectButtonHandler
func drive_ejectButtonHandler(_ *C.GObject, data C.gpointer) {
	signalDriveEjectButtonLock.RLock()
	defer signalDriveEjectButtonLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveEjectButtonMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_drive_can_eject.
func (recv *Drive) CanEject() bool {
	retC := C.g_drive_can_eject((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanPollForMedia is a wrapper around the C function g_drive_can_poll_for_media.
func (recv *Drive) CanPollForMedia() bool {
	retC := C.g_drive_can_poll_for_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_drive_eject_finish.
func (recv *Drive) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EnumerateIdentifiers is a wrapper around the C function g_drive_enumerate_identifiers.
func (recv *Drive) EnumerateIdentifiers() []string {
	retC := C.g_drive_enumerate_identifiers((*C.GDrive)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetIcon is a wrapper around the C function g_drive_get_icon.
func (recv *Drive) GetIcon() *Icon {
	retC := C.g_drive_get_icon((*C.GDrive)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIdentifier is a wrapper around the C function g_drive_get_identifier.
func (recv *Drive) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_drive_get_identifier((*C.GDrive)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_drive_get_name.
func (recv *Drive) GetName() string {
	retC := C.g_drive_get_name((*C.GDrive)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetVolumes is a wrapper around the C function g_drive_get_volumes.
func (recv *Drive) GetVolumes() *glib.List {
	retC := C.g_drive_get_volumes((*C.GDrive)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasMedia is a wrapper around the C function g_drive_has_media.
func (recv *Drive) HasMedia() bool {
	retC := C.g_drive_has_media((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasVolumes is a wrapper around the C function g_drive_has_volumes.
func (recv *Drive) HasVolumes() bool {
	retC := C.g_drive_has_volumes((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMediaCheckAutomatic is a wrapper around the C function g_drive_is_media_check_automatic.
func (recv *Drive) IsMediaCheckAutomatic() bool {
	retC := C.g_drive_is_media_check_automatic((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsMediaRemovable is a wrapper around the C function g_drive_is_media_removable.
func (recv *Drive) IsMediaRemovable() bool {
	retC := C.g_drive_is_media_removable((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_poll_for_media : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// PollForMediaFinish is a wrapper around the C function g_drive_poll_for_media_finish.
func (recv *Drive) PollForMediaFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_poll_for_media_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// File is a wrapper around the C record GFile.
type File struct {
	native *C.GFile
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this File with another File, and returns true if they represent the same GObject.
func (recv *File) Equals(other *File) bool {
	return other.ToC() == recv.ToC()
}

// FileNewForCommandlineArg is a wrapper around the C function g_file_new_for_commandline_arg.
func FileNewForCommandlineArg(arg string) *File {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	retC := C.g_file_new_for_commandline_arg(c_arg)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileNewForPath is a wrapper around the C function g_file_new_for_path.
func FileNewForPath(path string) *File {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_file_new_for_path(c_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileNewForUri is a wrapper around the C function g_file_new_for_uri.
func FileNewForUri(uri string) *File {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_file_new_for_uri(c_uri)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileParseName is a wrapper around the C function g_file_parse_name.
func FileParseName(parseName string) *File {
	c_parse_name := C.CString(parseName)
	defer C.free(unsafe.Pointer(c_parse_name))

	retC := C.g_file_parse_name(c_parse_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendTo is a wrapper around the C function g_file_append_to.
func (recv *File) AppendTo(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_append_to((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_append_to_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AppendToFinish is a wrapper around the C function g_file_append_to_finish.
func (recv *File) AppendToFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_append_to_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_copy : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Unsupported : g_file_copy_async : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// CopyAttributes is a wrapper around the C function g_file_copy_attributes.
func (recv *File) CopyAttributes(destination *File, flags FileCopyFlags, cancellable *Cancellable) (bool, error) {
	c_destination := (*C.GFile)(destination.ToC())

	c_flags := (C.GFileCopyFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_copy_attributes((*C.GFile)(recv.native), c_destination, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CopyFinish is a wrapper around the C function g_file_copy_finish.
func (recv *File) CopyFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_copy_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Create is a wrapper around the C function g_file_create.
func (recv *File) Create(flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_create((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_create_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CreateFinish is a wrapper around the C function g_file_create_finish.
func (recv *File) CreateFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Delete is a wrapper around the C function g_file_delete.
func (recv *File) Delete(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_delete((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Dup is a wrapper around the C function g_file_dup.
func (recv *File) Dup() *File {
	retC := C.g_file_dup((*C.GFile)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_eject_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectMountableFinish is a wrapper around the C function g_file_eject_mountable_finish.
func (recv *File) EjectMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EnumerateChildren is a wrapper around the C function g_file_enumerate_children.
func (recv *File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileEnumerator, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_enumerate_children_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EnumerateChildrenFinish is a wrapper around the C function g_file_enumerate_children_finish.
func (recv *File) EnumerateChildrenFinish(res *AsyncResult) (*FileEnumerator, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerate_children_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileEnumeratorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Equal is a wrapper around the C function g_file_equal.
func (recv *File) Equal(file2 *File) bool {
	c_file2 := (*C.GFile)(file2.ToC())

	retC := C.g_file_equal((*C.GFile)(recv.native), c_file2)
	retGo := retC == C.TRUE

	return retGo
}

// FindEnclosingMount is a wrapper around the C function g_file_find_enclosing_mount.
func (recv *File) FindEnclosingMount(cancellable *Cancellable) (*Mount, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_find_enclosing_mount_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// FindEnclosingMountFinish is a wrapper around the C function g_file_find_enclosing_mount_finish.
func (recv *File) FindEnclosingMountFinish(res *AsyncResult) (*Mount, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_find_enclosing_mount_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetBasename is a wrapper around the C function g_file_get_basename.
func (recv *File) GetBasename() string {
	retC := C.g_file_get_basename((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetChild is a wrapper around the C function g_file_get_child.
func (recv *File) GetChild(name string) *File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_file_get_child((*C.GFile)(recv.native), c_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildForDisplayName is a wrapper around the C function g_file_get_child_for_display_name.
func (recv *File) GetChildForDisplayName(displayName string) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	var cThrowableError *C.GError

	retC := C.g_file_get_child_for_display_name((*C.GFile)(recv.native), c_display_name, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetParent is a wrapper around the C function g_file_get_parent.
func (recv *File) GetParent() *File {
	retC := C.g_file_get_parent((*C.GFile)(recv.native))
	var retGo (*File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetParseName is a wrapper around the C function g_file_get_parse_name.
func (recv *File) GetParseName() string {
	retC := C.g_file_get_parse_name((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPath is a wrapper around the C function g_file_get_path.
func (recv *File) GetPath() string {
	retC := C.g_file_get_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetRelativePath is a wrapper around the C function g_file_get_relative_path.
func (recv *File) GetRelativePath(descendant *File) string {
	c_descendant := (*C.GFile)(descendant.ToC())

	retC := C.g_file_get_relative_path((*C.GFile)(recv.native), c_descendant)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function g_file_get_uri.
func (recv *File) GetUri() string {
	retC := C.g_file_get_uri((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUriScheme is a wrapper around the C function g_file_get_uri_scheme.
func (recv *File) GetUriScheme() string {
	retC := C.g_file_get_uri_scheme((*C.GFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// HasPrefix is a wrapper around the C function g_file_has_prefix.
func (recv *File) HasPrefix(prefix *File) bool {
	c_prefix := (*C.GFile)(prefix.ToC())

	retC := C.g_file_has_prefix((*C.GFile)(recv.native), c_prefix)
	retGo := retC == C.TRUE

	return retGo
}

// HasUriScheme is a wrapper around the C function g_file_has_uri_scheme.
func (recv *File) HasUriScheme(uriScheme string) bool {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_file_has_uri_scheme((*C.GFile)(recv.native), c_uri_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// Hash is a wrapper around the C function g_file_hash.
func (recv *File) Hash() uint32 {
	retC := C.g_file_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// IsNative is a wrapper around the C function g_file_is_native.
func (recv *File) IsNative() bool {
	retC := C.g_file_is_native((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_load_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_load_contents_finish : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_partial_contents_async : unsupported parameter read_more_callback : no type generator for FileReadMoreCallback (GFileReadMoreCallback) for param read_more_callback

// Unsupported : g_file_load_partial_contents_finish : unsupported parameter contents : output array param contents

// MakeDirectory is a wrapper around the C function g_file_make_directory.
func (recv *File) MakeDirectory(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_directory((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MakeSymbolicLink is a wrapper around the C function g_file_make_symbolic_link.
func (recv *File) MakeSymbolicLink(symlinkValue string, cancellable *Cancellable) (bool, error) {
	c_symlink_value := C.CString(symlinkValue)
	defer C.free(unsafe.Pointer(c_symlink_value))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_symbolic_link((*C.GFile)(recv.native), c_symlink_value, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MonitorDirectory is a wrapper around the C function g_file_monitor_directory.
func (recv *File) MonitorDirectory(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor_directory((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MonitorFile is a wrapper around the C function g_file_monitor_file.
func (recv *File) MonitorFile(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor_file((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_mount_enclosing_volume : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountEnclosingVolumeFinish is a wrapper around the C function g_file_mount_enclosing_volume_finish.
func (recv *File) MountEnclosingVolumeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_enclosing_volume_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_mount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountMountableFinish is a wrapper around the C function g_file_mount_mountable_finish.
func (recv *File) MountMountableFinish(result *AsyncResult) (*File, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_mount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_move : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// QueryDefaultHandler is a wrapper around the C function g_file_query_default_handler.
func (recv *File) QueryDefaultHandler(cancellable *Cancellable) (*AppInfo, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_default_handler((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QueryExists is a wrapper around the C function g_file_query_exists.
func (recv *File) QueryExists(cancellable *Cancellable) bool {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_file_query_exists((*C.GFile)(recv.native), c_cancellable)
	retGo := retC == C.TRUE

	return retGo
}

// QueryFilesystemInfo is a wrapper around the C function g_file_query_filesystem_info.
func (recv *File) QueryFilesystemInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info((*C.GFile)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_query_filesystem_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryFilesystemInfoFinish is a wrapper around the C function g_file_query_filesystem_info_finish.
func (recv *File) QueryFilesystemInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_filesystem_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QueryInfo is a wrapper around the C function g_file_query_info.
func (recv *File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_info((*C.GFile)(recv.native), c_attributes, c_flags, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryInfoFinish is a wrapper around the C function g_file_query_info_finish.
func (recv *File) QueryInfoFinish(res *AsyncResult) (*FileInfo, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_query_info_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QuerySettableAttributes is a wrapper around the C function g_file_query_settable_attributes.
func (recv *File) QuerySettableAttributes(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_settable_attributes((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QueryWritableNamespaces is a wrapper around the C function g_file_query_writable_namespaces.
func (recv *File) QueryWritableNamespaces(cancellable *Cancellable) (*FileAttributeInfoList, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_query_writable_namespaces((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Read is a wrapper around the C function g_file_read.
func (recv *File) Read(cancellable *Cancellable) (*FileInputStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_read((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReadFinish is a wrapper around the C function g_file_read_finish.
func (recv *File) ReadFinish(res *AsyncResult) (*FileInputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_read_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileInputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Replace is a wrapper around the C function g_file_replace.
func (recv *File) Replace(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (*FileOutputStream, error) {
	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace((*C.GFile)(recv.native), c_etag, c_make_backup, c_flags, c_cancellable, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_replace_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceContents is a wrapper around the C function g_file_replace_contents.
func (recv *File) ReplaceContents(contents []uint8, etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (bool, string, error) {
	c_contents_array := make([]C.guint8, len(contents)+1, len(contents)+1)
	for i, item := range contents {
		c := (C.guint8)(item)
		c_contents_array[i] = c
	}
	c_contents_array[len(contents)] = 0
	c_contents_arrayPtr := &c_contents_array[0]
	c_contents := (*C.char)(unsafe.Pointer(c_contents_arrayPtr))

	c_length := (C.gsize)(len(contents))

	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	var c_new_etag *C.char

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents((*C.GFile)(recv.native), c_contents, c_length, c_etag, c_make_backup, c_flags, &c_new_etag, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goError
}

// Unsupported : g_file_replace_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceContentsFinish is a wrapper around the C function g_file_replace_contents_finish.
func (recv *File) ReplaceContentsFinish(res *AsyncResult) (bool, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_new_etag *C.char

	var cThrowableError *C.GError

	retC := C.g_file_replace_contents_finish((*C.GFile)(recv.native), c_res, &c_new_etag, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	newEtag := C.GoString(c_new_etag)
	defer C.free(unsafe.Pointer(c_new_etag))

	return retGo, newEtag, goError
}

// ReplaceFinish is a wrapper around the C function g_file_replace_finish.
func (recv *File) ReplaceFinish(res *AsyncResult) (*FileOutputStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileOutputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ResolveRelativePath is a wrapper around the C function g_file_resolve_relative_path.
func (recv *File) ResolveRelativePath(relativePath string) *File {
	c_relative_path := C.CString(relativePath)
	defer C.free(unsafe.Pointer(c_relative_path))

	retC := C.g_file_resolve_relative_path((*C.GFile)(recv.native), c_relative_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_file_set_attribute : unsupported parameter value_p : no type generator for gpointer (gpointer) for param value_p

// SetAttributeByteString is a wrapper around the C function g_file_set_attribute_byte_string.
func (recv *File) SetAttributeByteString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_byte_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAttributeInt32 is a wrapper around the C function g_file_set_attribute_int32.
func (recv *File) SetAttributeInt32(attribute string, value int32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAttributeInt64 is a wrapper around the C function g_file_set_attribute_int64.
func (recv *File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.gint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_int64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAttributeString is a wrapper around the C function g_file_set_attribute_string.
func (recv *File) SetAttributeString(attribute string, value string, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_string((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAttributeUint32 is a wrapper around the C function g_file_set_attribute_uint32.
func (recv *File) SetAttributeUint32(attribute string, value uint32, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint32)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint32((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAttributeUint64 is a wrapper around the C function g_file_set_attribute_uint64.
func (recv *File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_value := (C.guint64)(value)

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attribute_uint64((*C.GFile)(recv.native), c_attribute, c_value, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_set_attributes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SetAttributesFinish is a wrapper around the C function g_file_set_attributes_finish.
func (recv *File) SetAttributesFinish(result *AsyncResult) (bool, *FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_info *C.GFileInfo

	var cThrowableError *C.GError

	retC := C.g_file_set_attributes_finish((*C.GFile)(recv.native), c_result, &c_info, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	info := FileInfoNewFromC(unsafe.Pointer(c_info))

	return retGo, info, goError
}

// SetAttributesFromInfo is a wrapper around the C function g_file_set_attributes_from_info.
func (recv *File) SetAttributesFromInfo(info *FileInfo, flags FileQueryInfoFlags, cancellable *Cancellable) (bool, error) {
	c_info := (*C.GFileInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GFileInfo)(info.ToC())
	}

	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_attributes_from_info((*C.GFile)(recv.native), c_info, c_flags, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetDisplayName is a wrapper around the C function g_file_set_display_name.
func (recv *File) SetDisplayName(displayName string, cancellable *Cancellable) (*File, error) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name((*C.GFile)(recv.native), c_display_name, c_cancellable, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_set_display_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SetDisplayNameFinish is a wrapper around the C function g_file_set_display_name_finish.
func (recv *File) SetDisplayNameFinish(res *AsyncResult) (*File, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_set_display_name_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Trash is a wrapper around the C function g_file_trash.
func (recv *File) Trash(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_trash((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_unmount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountMountableFinish is a wrapper around the C function g_file_unmount_mountable_finish.
func (recv *File) UnmountMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// FileDescriptorBased is a wrapper around the C record GFileDescriptorBased.
type FileDescriptorBased struct {
	native *C.GFileDescriptorBased
}

func FileDescriptorBasedNewFromC(u unsafe.Pointer) *FileDescriptorBased {
	c := (*C.GFileDescriptorBased)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBased{native: c}

	return g
}

func (recv *FileDescriptorBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileDescriptorBased with another FileDescriptorBased, and returns true if they represent the same GObject.
func (recv *FileDescriptorBased) Equals(other *FileDescriptorBased) bool {
	return other.ToC() == recv.ToC()
}

// Icon is a wrapper around the C record GIcon.
type Icon struct {
	native *C.GIcon
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	c := (*C.GIcon)(u)
	if c == nil {
		return nil
	}

	g := &Icon{native: c}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Icon with another Icon, and returns true if they represent the same GObject.
func (recv *Icon) Equals(other *Icon) bool {
	return other.ToC() == recv.ToC()
}

// g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon
// Equal is a wrapper around the C function g_icon_equal.
func (recv *Icon) Equal(icon2 *Icon) bool {
	c_icon2 := (*C.GIcon)(icon2.ToC())

	retC := C.g_icon_equal((*C.GIcon)(recv.native), c_icon2)
	retGo := retC == C.TRUE

	return retGo
}

// ListModel is a wrapper around the C record GListModel.
type ListModel struct {
	native *C.GListModel
}

func ListModelNewFromC(u unsafe.Pointer) *ListModel {
	c := (*C.GListModel)(u)
	if c == nil {
		return nil
	}

	g := &ListModel{native: c}

	return g
}

func (recv *ListModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ListModel with another ListModel, and returns true if they represent the same GObject.
func (recv *ListModel) Equals(other *ListModel) bool {
	return other.ToC() == recv.ToC()
}

// LoadableIcon is a wrapper around the C record GLoadableIcon.
type LoadableIcon struct {
	native *C.GLoadableIcon
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	c := (*C.GLoadableIcon)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIcon{native: c}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LoadableIcon with another LoadableIcon, and returns true if they represent the same GObject.
func (recv *LoadableIcon) Equals(other *LoadableIcon) bool {
	return other.ToC() == recv.ToC()
}

// Load is a wrapper around the C function g_loadable_icon_load.
func (recv *LoadableIcon) Load(size int32, cancellable *Cancellable) (*InputStream, string, error) {
	c_size := (C.int)(size)

	var c_type *C.char

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load((*C.GLoadableIcon)(recv.native), c_size, &c_type, c_cancellable, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goError
}

// Unsupported : g_loadable_icon_load_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadFinish is a wrapper around the C function g_loadable_icon_load_finish.
func (recv *LoadableIcon) LoadFinish(res *AsyncResult) (*InputStream, string, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var c_type *C.char

	var cThrowableError *C.GError

	retC := C.g_loadable_icon_load_finish((*C.GLoadableIcon)(recv.native), c_res, &c_type, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	type_ := C.GoString(c_type)
	defer C.free(unsafe.Pointer(c_type))

	return retGo, type_, goError
}

// Mount is a wrapper around the C record GMount.
type Mount struct {
	native *C.GMount
}

func MountNewFromC(u unsafe.Pointer) *Mount {
	c := (*C.GMount)(u)
	if c == nil {
		return nil
	}

	g := &Mount{native: c}

	return g
}

func (recv *Mount) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Mount with another Mount, and returns true if they represent the same GObject.
func (recv *Mount) Equals(other *Mount) bool {
	return other.ToC() == recv.ToC()
}

type signalMountChangedDetail struct {
	callback  MountSignalChangedCallback
	handlerID C.gulong
}

var signalMountChangedId int
var signalMountChangedMap = make(map[int]signalMountChangedDetail)
var signalMountChangedLock sync.RWMutex

// MountSignalChangedCallback is a callback function for a 'changed' signal emitted from a Mount.
type MountSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Mount) ConnectChanged(callback MountSignalChangedCallback) int {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	signalMountChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_changed(instance, C.gpointer(uintptr(signalMountChangedId)))

	detail := signalMountChangedDetail{callback, handlerID}
	signalMountChangedMap[signalMountChangedId] = detail

	return signalMountChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Mount.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Mount) DisconnectChanged(connectionID int) {
	signalMountChangedLock.Lock()
	defer signalMountChangedLock.Unlock()

	detail, exists := signalMountChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountChangedMap, connectionID)
}

//export mount_changedHandler
func mount_changedHandler(_ *C.GObject, data C.gpointer) {
	signalMountChangedLock.RLock()
	defer signalMountChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountChangedMap[index].callback
	callback()
}

type signalMountUnmountedDetail struct {
	callback  MountSignalUnmountedCallback
	handlerID C.gulong
}

var signalMountUnmountedId int
var signalMountUnmountedMap = make(map[int]signalMountUnmountedDetail)
var signalMountUnmountedLock sync.RWMutex

// MountSignalUnmountedCallback is a callback function for a 'unmounted' signal emitted from a Mount.
type MountSignalUnmountedCallback func()

/*
ConnectUnmounted connects the callback to the 'unmounted' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectUnmounted to remove it.
*/
func (recv *Mount) ConnectUnmounted(callback MountSignalUnmountedCallback) int {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	signalMountUnmountedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_unmounted(instance, C.gpointer(uintptr(signalMountUnmountedId)))

	detail := signalMountUnmountedDetail{callback, handlerID}
	signalMountUnmountedMap[signalMountUnmountedId] = detail

	return signalMountUnmountedId
}

/*
DisconnectUnmounted disconnects a callback from the 'unmounted' signal for the Mount.

The connectionID should be a value returned from a call to ConnectUnmounted.
*/
func (recv *Mount) DisconnectUnmounted(connectionID int) {
	signalMountUnmountedLock.Lock()
	defer signalMountUnmountedLock.Unlock()

	detail, exists := signalMountUnmountedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountUnmountedMap, connectionID)
}

//export mount_unmountedHandler
func mount_unmountedHandler(_ *C.GObject, data C.gpointer) {
	signalMountUnmountedLock.RLock()
	defer signalMountUnmountedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountUnmountedMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_mount_can_eject.
func (recv *Mount) CanEject() bool {
	retC := C.g_mount_can_eject((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanUnmount is a wrapper around the C function g_mount_can_unmount.
func (recv *Mount) CanUnmount() bool {
	retC := C.g_mount_can_unmount((*C.GMount)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_mount_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_mount_eject_finish.
func (recv *Mount) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetDefaultLocation is a wrapper around the C function g_mount_get_default_location.
func (recv *Mount) GetDefaultLocation() *File {
	retC := C.g_mount_get_default_location((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDrive is a wrapper around the C function g_mount_get_drive.
func (recv *Mount) GetDrive() *Drive {
	retC := C.g_mount_get_drive((*C.GMount)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIcon is a wrapper around the C function g_mount_get_icon.
func (recv *Mount) GetIcon() *Icon {
	retC := C.g_mount_get_icon((*C.GMount)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_mount_get_name.
func (recv *Mount) GetName() string {
	retC := C.g_mount_get_name((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetRoot is a wrapper around the C function g_mount_get_root.
func (recv *Mount) GetRoot() *File {
	retC := C.g_mount_get_root((*C.GMount)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUuid is a wrapper around the C function g_mount_get_uuid.
func (recv *Mount) GetUuid() string {
	retC := C.g_mount_get_uuid((*C.GMount)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetVolume is a wrapper around the C function g_mount_get_volume.
func (recv *Mount) GetVolume() *Volume {
	retC := C.g_mount_get_volume((*C.GMount)(recv.native))
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_mount_remount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// RemountFinish is a wrapper around the C function g_mount_remount_finish.
func (recv *Mount) RemountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_remount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_mount_unmount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountFinish is a wrapper around the C function g_mount_unmount_finish.
func (recv *Mount) UnmountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoteActionGroup is a wrapper around the C record GRemoteActionGroup.
type RemoteActionGroup struct {
	native *C.GRemoteActionGroup
}

func RemoteActionGroupNewFromC(u unsafe.Pointer) *RemoteActionGroup {
	c := (*C.GRemoteActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroup{native: c}

	return g
}

func (recv *RemoteActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RemoteActionGroup with another RemoteActionGroup, and returns true if they represent the same GObject.
func (recv *RemoteActionGroup) Equals(other *RemoteActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Seekable is a wrapper around the C record GSeekable.
type Seekable struct {
	native *C.GSeekable
}

func SeekableNewFromC(u unsafe.Pointer) *Seekable {
	c := (*C.GSeekable)(u)
	if c == nil {
		return nil
	}

	g := &Seekable{native: c}

	return g
}

func (recv *Seekable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Seekable with another Seekable, and returns true if they represent the same GObject.
func (recv *Seekable) Equals(other *Seekable) bool {
	return other.ToC() == recv.ToC()
}

// CanSeek is a wrapper around the C function g_seekable_can_seek.
func (recv *Seekable) CanSeek() bool {
	retC := C.g_seekable_can_seek((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanTruncate is a wrapper around the C function g_seekable_can_truncate.
func (recv *Seekable) CanTruncate() bool {
	retC := C.g_seekable_can_truncate((*C.GSeekable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Seek is a wrapper around the C function g_seekable_seek.
func (recv *Seekable) Seek(offset int64, type_ glib.SeekType, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_type := (C.GSeekType)(type_)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_seekable_seek((*C.GSeekable)(recv.native), c_offset, c_type, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Tell is a wrapper around the C function g_seekable_tell.
func (recv *Seekable) Tell() int64 {
	retC := C.g_seekable_tell((*C.GSeekable)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Truncate is a wrapper around the C function g_seekable_truncate.
func (recv *Seekable) Truncate(offset int64, cancellable *Cancellable) (bool, error) {
	c_offset := (C.goffset)(offset)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_seekable_truncate((*C.GSeekable)(recv.native), c_offset, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SocketConnectable is a wrapper around the C record GSocketConnectable.
type SocketConnectable struct {
	native *C.GSocketConnectable
}

func SocketConnectableNewFromC(u unsafe.Pointer) *SocketConnectable {
	c := (*C.GSocketConnectable)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectable{native: c}

	return g
}

func (recv *SocketConnectable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnectable with another SocketConnectable, and returns true if they represent the same GObject.
func (recv *SocketConnectable) Equals(other *SocketConnectable) bool {
	return other.ToC() == recv.ToC()
}

// Volume is a wrapper around the C record GVolume.
type Volume struct {
	native *C.GVolume
}

func VolumeNewFromC(u unsafe.Pointer) *Volume {
	c := (*C.GVolume)(u)
	if c == nil {
		return nil
	}

	g := &Volume{native: c}

	return g
}

func (recv *Volume) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Volume with another Volume, and returns true if they represent the same GObject.
func (recv *Volume) Equals(other *Volume) bool {
	return other.ToC() == recv.ToC()
}

type signalVolumeChangedDetail struct {
	callback  VolumeSignalChangedCallback
	handlerID C.gulong
}

var signalVolumeChangedId int
var signalVolumeChangedMap = make(map[int]signalVolumeChangedDetail)
var signalVolumeChangedLock sync.RWMutex

// VolumeSignalChangedCallback is a callback function for a 'changed' signal emitted from a Volume.
type VolumeSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *Volume) ConnectChanged(callback VolumeSignalChangedCallback) int {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	signalVolumeChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_changed(instance, C.gpointer(uintptr(signalVolumeChangedId)))

	detail := signalVolumeChangedDetail{callback, handlerID}
	signalVolumeChangedMap[signalVolumeChangedId] = detail

	return signalVolumeChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *Volume) DisconnectChanged(connectionID int) {
	signalVolumeChangedLock.Lock()
	defer signalVolumeChangedLock.Unlock()

	detail, exists := signalVolumeChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeChangedMap, connectionID)
}

//export volume_changedHandler
func volume_changedHandler(_ *C.GObject, data C.gpointer) {
	signalVolumeChangedLock.RLock()
	defer signalVolumeChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalVolumeChangedMap[index].callback
	callback()
}

type signalVolumeRemovedDetail struct {
	callback  VolumeSignalRemovedCallback
	handlerID C.gulong
}

var signalVolumeRemovedId int
var signalVolumeRemovedMap = make(map[int]signalVolumeRemovedDetail)
var signalVolumeRemovedLock sync.RWMutex

// VolumeSignalRemovedCallback is a callback function for a 'removed' signal emitted from a Volume.
type VolumeSignalRemovedCallback func()

/*
ConnectRemoved connects the callback to the 'removed' signal for the Volume.

The returned value represents the connection, and may be passed to DisconnectRemoved to remove it.
*/
func (recv *Volume) ConnectRemoved(callback VolumeSignalRemovedCallback) int {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	signalVolumeRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Volume_signal_connect_removed(instance, C.gpointer(uintptr(signalVolumeRemovedId)))

	detail := signalVolumeRemovedDetail{callback, handlerID}
	signalVolumeRemovedMap[signalVolumeRemovedId] = detail

	return signalVolumeRemovedId
}

/*
DisconnectRemoved disconnects a callback from the 'removed' signal for the Volume.

The connectionID should be a value returned from a call to ConnectRemoved.
*/
func (recv *Volume) DisconnectRemoved(connectionID int) {
	signalVolumeRemovedLock.Lock()
	defer signalVolumeRemovedLock.Unlock()

	detail, exists := signalVolumeRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeRemovedMap, connectionID)
}

//export volume_removedHandler
func volume_removedHandler(_ *C.GObject, data C.gpointer) {
	signalVolumeRemovedLock.RLock()
	defer signalVolumeRemovedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalVolumeRemovedMap[index].callback
	callback()
}

// CanEject is a wrapper around the C function g_volume_can_eject.
func (recv *Volume) CanEject() bool {
	retC := C.g_volume_can_eject((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanMount is a wrapper around the C function g_volume_can_mount.
func (recv *Volume) CanMount() bool {
	retC := C.g_volume_can_mount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_volume_eject : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectFinish is a wrapper around the C function g_volume_eject_finish.
func (recv *Volume) EjectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EnumerateIdentifiers is a wrapper around the C function g_volume_enumerate_identifiers.
func (recv *Volume) EnumerateIdentifiers() []string {
	retC := C.g_volume_enumerate_identifiers((*C.GVolume)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetDrive is a wrapper around the C function g_volume_get_drive.
func (recv *Volume) GetDrive() *Drive {
	retC := C.g_volume_get_drive((*C.GVolume)(recv.native))
	retGo := DriveNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIcon is a wrapper around the C function g_volume_get_icon.
func (recv *Volume) GetIcon() *Icon {
	retC := C.g_volume_get_icon((*C.GVolume)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIdentifier is a wrapper around the C function g_volume_get_identifier.
func (recv *Volume) GetIdentifier(kind string) string {
	c_kind := C.CString(kind)
	defer C.free(unsafe.Pointer(c_kind))

	retC := C.g_volume_get_identifier((*C.GVolume)(recv.native), c_kind)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetMount is a wrapper around the C function g_volume_get_mount.
func (recv *Volume) GetMount() *Mount {
	retC := C.g_volume_get_mount((*C.GVolume)(recv.native))
	retGo := MountNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function g_volume_get_name.
func (recv *Volume) GetName() string {
	retC := C.g_volume_get_name((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUuid is a wrapper around the C function g_volume_get_uuid.
func (recv *Volume) GetUuid() string {
	retC := C.g_volume_get_uuid((*C.GVolume)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_mount : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// MountFinish is a wrapper around the C function g_volume_mount_finish.
func (recv *Volume) MountFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_mount_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ShouldAutomount is a wrapper around the C function g_volume_should_automount.
func (recv *Volume) ShouldAutomount() bool {
	retC := C.g_volume_should_automount((*C.GVolume)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ActionEntry is a wrapper around the C record GActionEntry.
type ActionEntry struct {
	native *C.GActionEntry
	Name   string
	// no type for activate
	ParameterType string
	State         string
	// no type for change_state
	// Private : padding
}

func ActionEntryNewFromC(u unsafe.Pointer) *ActionEntry {
	c := (*C.GActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &ActionEntry{
		Name:          C.GoString(c.name),
		ParameterType: C.GoString(c.parameter_type),
		State:         C.GoString(c.state),
		native:        c,
	}

	return g
}

func (recv *ActionEntry) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.parameter_type =
		C.CString(recv.ParameterType)
	recv.native.state =
		C.CString(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionEntry with another ActionEntry, and returns true if they represent the same GObject.
func (recv *ActionEntry) Equals(other *ActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// AppInfoIface is a wrapper around the C record GAppInfoIface.
type AppInfoIface struct {
	native *C.GAppInfoIface
	// g_iface : record
	// no type for dup
	// no type for equal
	// no type for get_id
	// no type for get_name
	// no type for get_description
	// no type for get_executable
	// no type for get_icon
	// no type for launch
	// no type for supports_uris
	// no type for supports_files
	// no type for launch_uris
	// no type for should_show
	// no type for set_as_default_for_type
	// no type for set_as_default_for_extension
	// no type for add_supports_type
	// no type for can_remove_supports_type
	// no type for remove_supports_type
	// no type for can_delete
	// no type for do_delete
	// no type for get_commandline
	// no type for get_display_name
	// no type for set_as_last_used_for_type
	// no type for get_supported_types
}

func AppInfoIfaceNewFromC(u unsafe.Pointer) *AppInfoIface {
	c := (*C.GAppInfoIface)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoIface{native: c}

	return g
}

func (recv *AppInfoIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppInfoIface with another AppInfoIface, and returns true if they represent the same GObject.
func (recv *AppInfoIface) Equals(other *AppInfoIface) bool {
	return other.ToC() == recv.ToC()
}

// AppLaunchContextClass is a wrapper around the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native *C.GAppLaunchContextClass
	// parent_class : record
	// no type for get_display
	// no type for get_startup_notify_id
	// no type for launch_failed
	// no type for launched
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func AppLaunchContextClassNewFromC(u unsafe.Pointer) *AppLaunchContextClass {
	c := (*C.GAppLaunchContextClass)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextClass{native: c}

	return g
}

func (recv *AppLaunchContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppLaunchContextClass with another AppLaunchContextClass, and returns true if they represent the same GObject.
func (recv *AppLaunchContextClass) Equals(other *AppLaunchContextClass) bool {
	return other.ToC() == recv.ToC()
}

// AppLaunchContextPrivate is a wrapper around the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native *C.GAppLaunchContextPrivate
}

func AppLaunchContextPrivateNewFromC(u unsafe.Pointer) *AppLaunchContextPrivate {
	c := (*C.GAppLaunchContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextPrivate{native: c}

	return g
}

func (recv *AppLaunchContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppLaunchContextPrivate with another AppLaunchContextPrivate, and returns true if they represent the same GObject.
func (recv *AppLaunchContextPrivate) Equals(other *AppLaunchContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationCommandLinePrivate is a wrapper around the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native *C.GApplicationCommandLinePrivate
}

func ApplicationCommandLinePrivateNewFromC(u unsafe.Pointer) *ApplicationCommandLinePrivate {
	c := (*C.GApplicationCommandLinePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLinePrivate{native: c}

	return g
}

func (recv *ApplicationCommandLinePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationCommandLinePrivate with another ApplicationCommandLinePrivate, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLinePrivate) Equals(other *ApplicationCommandLinePrivate) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationPrivate is a wrapper around the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native *C.GApplicationPrivate
}

func ApplicationPrivateNewFromC(u unsafe.Pointer) *ApplicationPrivate {
	c := (*C.GApplicationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationPrivate{native: c}

	return g
}

func (recv *ApplicationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationPrivate with another ApplicationPrivate, and returns true if they represent the same GObject.
func (recv *ApplicationPrivate) Equals(other *ApplicationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AsyncResultIface is a wrapper around the C record GAsyncResultIface.
type AsyncResultIface struct {
	native *C.GAsyncResultIface
	// g_iface : record
	// no type for get_user_data
	// no type for get_source_object
	// no type for is_tagged
}

func AsyncResultIfaceNewFromC(u unsafe.Pointer) *AsyncResultIface {
	c := (*C.GAsyncResultIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResultIface{native: c}

	return g
}

func (recv *AsyncResultIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncResultIface with another AsyncResultIface, and returns true if they represent the same GObject.
func (recv *AsyncResultIface) Equals(other *AsyncResultIface) bool {
	return other.ToC() == recv.ToC()
}

// BufferedInputStreamClass is a wrapper around the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native *C.GBufferedInputStreamClass
	// parent_class : record
	// no type for fill
	// no type for fill_async
	// no type for fill_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func BufferedInputStreamClassNewFromC(u unsafe.Pointer) *BufferedInputStreamClass {
	c := (*C.GBufferedInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamClass{native: c}

	return g
}

func (recv *BufferedInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedInputStreamClass with another BufferedInputStreamClass, and returns true if they represent the same GObject.
func (recv *BufferedInputStreamClass) Equals(other *BufferedInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// BufferedInputStreamPrivate is a wrapper around the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native *C.GBufferedInputStreamPrivate
}

func BufferedInputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedInputStreamPrivate {
	c := (*C.GBufferedInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamPrivate{native: c}

	return g
}

func (recv *BufferedInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedInputStreamPrivate with another BufferedInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *BufferedInputStreamPrivate) Equals(other *BufferedInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BufferedOutputStreamClass is a wrapper around the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native *C.GBufferedOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func BufferedOutputStreamClassNewFromC(u unsafe.Pointer) *BufferedOutputStreamClass {
	c := (*C.GBufferedOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamClass{native: c}

	return g
}

func (recv *BufferedOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedOutputStreamClass with another BufferedOutputStreamClass, and returns true if they represent the same GObject.
func (recv *BufferedOutputStreamClass) Equals(other *BufferedOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// BufferedOutputStreamPrivate is a wrapper around the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native *C.GBufferedOutputStreamPrivate
}

func BufferedOutputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedOutputStreamPrivate {
	c := (*C.GBufferedOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamPrivate{native: c}

	return g
}

func (recv *BufferedOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferedOutputStreamPrivate with another BufferedOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *BufferedOutputStreamPrivate) Equals(other *BufferedOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CancellableClass is a wrapper around the C record GCancellableClass.
type CancellableClass struct {
	native *C.GCancellableClass
	// parent_class : record
	// no type for cancelled
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func CancellableClassNewFromC(u unsafe.Pointer) *CancellableClass {
	c := (*C.GCancellableClass)(u)
	if c == nil {
		return nil
	}

	g := &CancellableClass{native: c}

	return g
}

func (recv *CancellableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CancellableClass with another CancellableClass, and returns true if they represent the same GObject.
func (recv *CancellableClass) Equals(other *CancellableClass) bool {
	return other.ToC() == recv.ToC()
}

// CancellablePrivate is a wrapper around the C record GCancellablePrivate.
type CancellablePrivate struct {
	native *C.GCancellablePrivate
}

func CancellablePrivateNewFromC(u unsafe.Pointer) *CancellablePrivate {
	c := (*C.GCancellablePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CancellablePrivate{native: c}

	return g
}

func (recv *CancellablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CancellablePrivate with another CancellablePrivate, and returns true if they represent the same GObject.
func (recv *CancellablePrivate) Equals(other *CancellablePrivate) bool {
	return other.ToC() == recv.ToC()
}

// CharsetConverterClass is a wrapper around the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native *C.GCharsetConverterClass
	// parent_class : record
}

func CharsetConverterClassNewFromC(u unsafe.Pointer) *CharsetConverterClass {
	c := (*C.GCharsetConverterClass)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverterClass{native: c}

	return g
}

func (recv *CharsetConverterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CharsetConverterClass with another CharsetConverterClass, and returns true if they represent the same GObject.
func (recv *CharsetConverterClass) Equals(other *CharsetConverterClass) bool {
	return other.ToC() == recv.ToC()
}

// ConverterInputStreamClass is a wrapper around the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native *C.GConverterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterInputStreamClassNewFromC(u unsafe.Pointer) *ConverterInputStreamClass {
	c := (*C.GConverterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamClass{native: c}

	return g
}

func (recv *ConverterInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterInputStreamClass with another ConverterInputStreamClass, and returns true if they represent the same GObject.
func (recv *ConverterInputStreamClass) Equals(other *ConverterInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// ConverterInputStreamPrivate is a wrapper around the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native *C.GConverterInputStreamPrivate
}

func ConverterInputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterInputStreamPrivate {
	c := (*C.GConverterInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamPrivate{native: c}

	return g
}

func (recv *ConverterInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterInputStreamPrivate with another ConverterInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *ConverterInputStreamPrivate) Equals(other *ConverterInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ConverterOutputStreamClass is a wrapper around the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native *C.GConverterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterOutputStreamClassNewFromC(u unsafe.Pointer) *ConverterOutputStreamClass {
	c := (*C.GConverterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamClass{native: c}

	return g
}

func (recv *ConverterOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterOutputStreamClass with another ConverterOutputStreamClass, and returns true if they represent the same GObject.
func (recv *ConverterOutputStreamClass) Equals(other *ConverterOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// ConverterOutputStreamPrivate is a wrapper around the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native *C.GConverterOutputStreamPrivate
}

func ConverterOutputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterOutputStreamPrivate {
	c := (*C.GConverterOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamPrivate{native: c}

	return g
}

func (recv *ConverterOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ConverterOutputStreamPrivate with another ConverterOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *ConverterOutputStreamPrivate) Equals(other *ConverterOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusInterfaceSkeletonPrivate is a wrapper around the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native *C.GDBusInterfaceSkeletonPrivate
}

func DBusInterfaceSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	c := (*C.GDBusInterfaceSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeletonPrivate{native: c}

	return g
}

func (recv *DBusInterfaceSkeletonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusInterfaceSkeletonPrivate with another DBusInterfaceSkeletonPrivate, and returns true if they represent the same GObject.
func (recv *DBusInterfaceSkeletonPrivate) Equals(other *DBusInterfaceSkeletonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectManagerClientPrivate is a wrapper around the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native *C.GDBusObjectManagerClientPrivate
}

func DBusObjectManagerClientPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerClientPrivate {
	c := (*C.GDBusObjectManagerClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClientPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerClientPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManagerClientPrivate with another DBusObjectManagerClientPrivate, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerClientPrivate) Equals(other *DBusObjectManagerClientPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectManagerServerPrivate is a wrapper around the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native *C.GDBusObjectManagerServerPrivate
}

func DBusObjectManagerServerPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerServerPrivate {
	c := (*C.GDBusObjectManagerServerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServerPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerServerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectManagerServerPrivate with another DBusObjectManagerServerPrivate, and returns true if they represent the same GObject.
func (recv *DBusObjectManagerServerPrivate) Equals(other *DBusObjectManagerServerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectProxyPrivate is a wrapper around the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native *C.GDBusObjectProxyPrivate
}

func DBusObjectProxyPrivateNewFromC(u unsafe.Pointer) *DBusObjectProxyPrivate {
	c := (*C.GDBusObjectProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxyPrivate{native: c}

	return g
}

func (recv *DBusObjectProxyPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectProxyPrivate with another DBusObjectProxyPrivate, and returns true if they represent the same GObject.
func (recv *DBusObjectProxyPrivate) Equals(other *DBusObjectProxyPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusObjectSkeletonPrivate is a wrapper around the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native *C.GDBusObjectSkeletonPrivate
}

func DBusObjectSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusObjectSkeletonPrivate {
	c := (*C.GDBusObjectSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeletonPrivate{native: c}

	return g
}

func (recv *DBusObjectSkeletonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusObjectSkeletonPrivate with another DBusObjectSkeletonPrivate, and returns true if they represent the same GObject.
func (recv *DBusObjectSkeletonPrivate) Equals(other *DBusObjectSkeletonPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DBusProxyPrivate is a wrapper around the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native *C.GDBusProxyPrivate
}

func DBusProxyPrivateNewFromC(u unsafe.Pointer) *DBusProxyPrivate {
	c := (*C.GDBusProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxyPrivate{native: c}

	return g
}

func (recv *DBusProxyPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DBusProxyPrivate with another DBusProxyPrivate, and returns true if they represent the same GObject.
func (recv *DBusProxyPrivate) Equals(other *DBusProxyPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DataInputStreamClass is a wrapper around the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native *C.GDataInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataInputStreamClassNewFromC(u unsafe.Pointer) *DataInputStreamClass {
	c := (*C.GDataInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamClass{native: c}

	return g
}

func (recv *DataInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataInputStreamClass with another DataInputStreamClass, and returns true if they represent the same GObject.
func (recv *DataInputStreamClass) Equals(other *DataInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// DataInputStreamPrivate is a wrapper around the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native *C.GDataInputStreamPrivate
}

func DataInputStreamPrivateNewFromC(u unsafe.Pointer) *DataInputStreamPrivate {
	c := (*C.GDataInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamPrivate{native: c}

	return g
}

func (recv *DataInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataInputStreamPrivate with another DataInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *DataInputStreamPrivate) Equals(other *DataInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DataOutputStreamClass is a wrapper around the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native *C.GDataOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataOutputStreamClassNewFromC(u unsafe.Pointer) *DataOutputStreamClass {
	c := (*C.GDataOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamClass{native: c}

	return g
}

func (recv *DataOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataOutputStreamClass with another DataOutputStreamClass, and returns true if they represent the same GObject.
func (recv *DataOutputStreamClass) Equals(other *DataOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// DataOutputStreamPrivate is a wrapper around the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native *C.GDataOutputStreamPrivate
}

func DataOutputStreamPrivateNewFromC(u unsafe.Pointer) *DataOutputStreamPrivate {
	c := (*C.GDataOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamPrivate{native: c}

	return g
}

func (recv *DataOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DataOutputStreamPrivate with another DataOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *DataOutputStreamPrivate) Equals(other *DataOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DesktopAppInfoClass is a wrapper around the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native *C.GDesktopAppInfoClass
	// parent_class : record
}

func DesktopAppInfoClassNewFromC(u unsafe.Pointer) *DesktopAppInfoClass {
	c := (*C.GDesktopAppInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoClass{native: c}

	return g
}

func (recv *DesktopAppInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DesktopAppInfoClass with another DesktopAppInfoClass, and returns true if they represent the same GObject.
func (recv *DesktopAppInfoClass) Equals(other *DesktopAppInfoClass) bool {
	return other.ToC() == recv.ToC()
}

// DesktopAppInfoLookupIface is a wrapper around the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native *C.GDesktopAppInfoLookupIface
	// g_iface : record
	// no type for get_default_for_uri_scheme
}

func DesktopAppInfoLookupIfaceNewFromC(u unsafe.Pointer) *DesktopAppInfoLookupIface {
	c := (*C.GDesktopAppInfoLookupIface)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookupIface{native: c}

	return g
}

func (recv *DesktopAppInfoLookupIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DesktopAppInfoLookupIface with another DesktopAppInfoLookupIface, and returns true if they represent the same GObject.
func (recv *DesktopAppInfoLookupIface) Equals(other *DesktopAppInfoLookupIface) bool {
	return other.ToC() == recv.ToC()
}

// DriveIface is a wrapper around the C record GDriveIface.
type DriveIface struct {
	native *C.GDriveIface
	// g_iface : record
	// no type for changed
	// no type for disconnected
	// no type for eject_button
	// no type for get_name
	// no type for get_icon
	// no type for has_volumes
	// no type for get_volumes
	// no type for is_media_removable
	// no type for has_media
	// no type for is_media_check_automatic
	// no type for can_eject
	// no type for can_poll_for_media
	// no type for eject
	// no type for eject_finish
	// no type for poll_for_media
	// no type for poll_for_media_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for get_start_stop_type
	// no type for can_start
	// no type for can_start_degraded
	// no type for start
	// no type for start_finish
	// no type for can_stop
	// no type for stop
	// no type for stop_finish
	// no type for stop_button
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
	// no type for is_removable
}

func DriveIfaceNewFromC(u unsafe.Pointer) *DriveIface {
	c := (*C.GDriveIface)(u)
	if c == nil {
		return nil
	}

	g := &DriveIface{native: c}

	return g
}

func (recv *DriveIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DriveIface with another DriveIface, and returns true if they represent the same GObject.
func (recv *DriveIface) Equals(other *DriveIface) bool {
	return other.ToC() == recv.ToC()
}

// EmblemClass is a wrapper around the C record GEmblemClass.
type EmblemClass struct {
	native *C.GEmblemClass
}

func EmblemClassNewFromC(u unsafe.Pointer) *EmblemClass {
	c := (*C.GEmblemClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemClass{native: c}

	return g
}

func (recv *EmblemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EmblemClass with another EmblemClass, and returns true if they represent the same GObject.
func (recv *EmblemClass) Equals(other *EmblemClass) bool {
	return other.ToC() == recv.ToC()
}

// EmblemedIconClass is a wrapper around the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native *C.GEmblemedIconClass
	// parent_class : record
}

func EmblemedIconClassNewFromC(u unsafe.Pointer) *EmblemedIconClass {
	c := (*C.GEmblemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconClass{native: c}

	return g
}

func (recv *EmblemedIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EmblemedIconClass with another EmblemedIconClass, and returns true if they represent the same GObject.
func (recv *EmblemedIconClass) Equals(other *EmblemedIconClass) bool {
	return other.ToC() == recv.ToC()
}

// EmblemedIconPrivate is a wrapper around the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native *C.GEmblemedIconPrivate
}

func EmblemedIconPrivateNewFromC(u unsafe.Pointer) *EmblemedIconPrivate {
	c := (*C.GEmblemedIconPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconPrivate{native: c}

	return g
}

func (recv *EmblemedIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EmblemedIconPrivate with another EmblemedIconPrivate, and returns true if they represent the same GObject.
func (recv *EmblemedIconPrivate) Equals(other *EmblemedIconPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileAttributeInfo is a wrapper around the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native *C.GFileAttributeInfo
	Name   string
	Type   FileAttributeType
	Flags  FileAttributeInfoFlags
}

func FileAttributeInfoNewFromC(u unsafe.Pointer) *FileAttributeInfo {
	c := (*C.GFileAttributeInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfo{
		Flags:  (FileAttributeInfoFlags)(c.flags),
		Name:   C.GoString(c.name),
		Type:   (FileAttributeType)(c._type),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native._type =
		(C.GFileAttributeType)(recv.Type)
	recv.native.flags =
		(C.GFileAttributeInfoFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileAttributeInfo with another FileAttributeInfo, and returns true if they represent the same GObject.
func (recv *FileAttributeInfo) Equals(other *FileAttributeInfo) bool {
	return other.ToC() == recv.ToC()
}

// FileAttributeInfoList is a wrapper around the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native *C.GFileAttributeInfoList
	// infos : record
	NInfos int32
}

func FileAttributeInfoListNewFromC(u unsafe.Pointer) *FileAttributeInfoList {
	c := (*C.GFileAttributeInfoList)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfoList{
		NInfos: (int32)(c.n_infos),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfoList) ToC() unsafe.Pointer {
	recv.native.n_infos =
		(C.int)(recv.NInfos)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileAttributeInfoList with another FileAttributeInfoList, and returns true if they represent the same GObject.
func (recv *FileAttributeInfoList) Equals(other *FileAttributeInfoList) bool {
	return other.ToC() == recv.ToC()
}

// FileAttributeInfoListNew is a wrapper around the C function g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_new()
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function g_file_attribute_info_list_add.
func (recv *FileAttributeInfoList) Add(name string, type_ FileAttributeType, flags FileAttributeInfoFlags) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_type := (C.GFileAttributeType)(type_)

	c_flags := (C.GFileAttributeInfoFlags)(flags)

	C.g_file_attribute_info_list_add((*C.GFileAttributeInfoList)(recv.native), c_name, c_type, c_flags)

	return
}

// Dup is a wrapper around the C function g_file_attribute_info_list_dup.
func (recv *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_dup((*C.GFileAttributeInfoList)(recv.native))
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Lookup is a wrapper around the C function g_file_attribute_info_list_lookup.
func (recv *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_file_attribute_info_list_lookup((*C.GFileAttributeInfoList)(recv.native), c_name)
	retGo := FileAttributeInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_file_attribute_info_list_ref.
func (recv *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_ref((*C.GFileAttributeInfoList)(recv.native))
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_file_attribute_info_list_unref.
func (recv *FileAttributeInfoList) Unref() {
	C.g_file_attribute_info_list_unref((*C.GFileAttributeInfoList)(recv.native))

	return
}

// FileAttributeMatcher is a wrapper around the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native *C.GFileAttributeMatcher
}

func FileAttributeMatcherNewFromC(u unsafe.Pointer) *FileAttributeMatcher {
	c := (*C.GFileAttributeMatcher)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeMatcher{native: c}

	return g
}

func (recv *FileAttributeMatcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileAttributeMatcher with another FileAttributeMatcher, and returns true if they represent the same GObject.
func (recv *FileAttributeMatcher) Equals(other *FileAttributeMatcher) bool {
	return other.ToC() == recv.ToC()
}

// FileAttributeMatcherNew is a wrapper around the C function g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	retC := C.g_file_attribute_matcher_new(c_attributes)
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EnumerateNamespace is a wrapper around the C function g_file_attribute_matcher_enumerate_namespace.
func (recv *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	c_ns := C.CString(ns)
	defer C.free(unsafe.Pointer(c_ns))

	retC := C.g_file_attribute_matcher_enumerate_namespace((*C.GFileAttributeMatcher)(recv.native), c_ns)
	retGo := retC == C.TRUE

	return retGo
}

// EnumerateNext is a wrapper around the C function g_file_attribute_matcher_enumerate_next.
func (recv *FileAttributeMatcher) EnumerateNext() string {
	retC := C.g_file_attribute_matcher_enumerate_next((*C.GFileAttributeMatcher)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Matches is a wrapper around the C function g_file_attribute_matcher_matches.
func (recv *FileAttributeMatcher) Matches(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_attribute_matcher_matches((*C.GFileAttributeMatcher)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// MatchesOnly is a wrapper around the C function g_file_attribute_matcher_matches_only.
func (recv *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_attribute_matcher_matches_only((*C.GFileAttributeMatcher)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_file_attribute_matcher_ref.
func (recv *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	retC := C.g_file_attribute_matcher_ref((*C.GFileAttributeMatcher)(recv.native))
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Subtract is a wrapper around the C function g_file_attribute_matcher_subtract.
func (recv *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) *FileAttributeMatcher {
	c_subtract := (*C.GFileAttributeMatcher)(C.NULL)
	if subtract != nil {
		c_subtract = (*C.GFileAttributeMatcher)(subtract.ToC())
	}

	retC := C.g_file_attribute_matcher_subtract((*C.GFileAttributeMatcher)(recv.native), c_subtract)
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_file_attribute_matcher_unref.
func (recv *FileAttributeMatcher) Unref() {
	C.g_file_attribute_matcher_unref((*C.GFileAttributeMatcher)(recv.native))

	return
}

// FileDescriptorBasedIface is a wrapper around the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native *C.GFileDescriptorBasedIface
	// g_iface : record
	// no type for get_fd
}

func FileDescriptorBasedIfaceNewFromC(u unsafe.Pointer) *FileDescriptorBasedIface {
	c := (*C.GFileDescriptorBasedIface)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBasedIface{native: c}

	return g
}

func (recv *FileDescriptorBasedIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileDescriptorBasedIface with another FileDescriptorBasedIface, and returns true if they represent the same GObject.
func (recv *FileDescriptorBasedIface) Equals(other *FileDescriptorBasedIface) bool {
	return other.ToC() == recv.ToC()
}

// FileEnumeratorClass is a wrapper around the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native *C.GFileEnumeratorClass
	// parent_class : record
	// no type for next_file
	// no type for close_fn
	// no type for next_files_async
	// no type for next_files_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func FileEnumeratorClassNewFromC(u unsafe.Pointer) *FileEnumeratorClass {
	c := (*C.GFileEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorClass{native: c}

	return g
}

func (recv *FileEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileEnumeratorClass with another FileEnumeratorClass, and returns true if they represent the same GObject.
func (recv *FileEnumeratorClass) Equals(other *FileEnumeratorClass) bool {
	return other.ToC() == recv.ToC()
}

// FileEnumeratorPrivate is a wrapper around the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native *C.GFileEnumeratorPrivate
}

func FileEnumeratorPrivateNewFromC(u unsafe.Pointer) *FileEnumeratorPrivate {
	c := (*C.GFileEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorPrivate{native: c}

	return g
}

func (recv *FileEnumeratorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileEnumeratorPrivate with another FileEnumeratorPrivate, and returns true if they represent the same GObject.
func (recv *FileEnumeratorPrivate) Equals(other *FileEnumeratorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileIOStreamClass is a wrapper around the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native *C.GFileIOStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileIOStreamClassNewFromC(u unsafe.Pointer) *FileIOStreamClass {
	c := (*C.GFileIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamClass{native: c}

	return g
}

func (recv *FileIOStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIOStreamClass with another FileIOStreamClass, and returns true if they represent the same GObject.
func (recv *FileIOStreamClass) Equals(other *FileIOStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// FileIOStreamPrivate is a wrapper around the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native *C.GFileIOStreamPrivate
}

func FileIOStreamPrivateNewFromC(u unsafe.Pointer) *FileIOStreamPrivate {
	c := (*C.GFileIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamPrivate{native: c}

	return g
}

func (recv *FileIOStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIOStreamPrivate with another FileIOStreamPrivate, and returns true if they represent the same GObject.
func (recv *FileIOStreamPrivate) Equals(other *FileIOStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileIconClass is a wrapper around the C record GFileIconClass.
type FileIconClass struct {
	native *C.GFileIconClass
}

func FileIconClassNewFromC(u unsafe.Pointer) *FileIconClass {
	c := (*C.GFileIconClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIconClass{native: c}

	return g
}

func (recv *FileIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIconClass with another FileIconClass, and returns true if they represent the same GObject.
func (recv *FileIconClass) Equals(other *FileIconClass) bool {
	return other.ToC() == recv.ToC()
}

// FileIface is a wrapper around the C record GFileIface.
type FileIface struct {
	native *C.GFileIface
	// g_iface : record
	// no type for dup
	// no type for hash
	// no type for equal
	// no type for is_native
	// no type for has_uri_scheme
	// no type for get_uri_scheme
	// no type for get_basename
	// no type for get_path
	// no type for get_uri
	// no type for get_parse_name
	// no type for get_parent
	// no type for prefix_matches
	// no type for get_relative_path
	// no type for resolve_relative_path
	// no type for get_child_for_display_name
	// no type for enumerate_children
	// no type for enumerate_children_async
	// no type for enumerate_children_finish
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for query_filesystem_info
	// no type for query_filesystem_info_async
	// no type for query_filesystem_info_finish
	// no type for find_enclosing_mount
	// no type for find_enclosing_mount_async
	// no type for find_enclosing_mount_finish
	// no type for set_display_name
	// no type for set_display_name_async
	// no type for set_display_name_finish
	// no type for query_settable_attributes
	// no type for _query_settable_attributes_async
	// no type for _query_settable_attributes_finish
	// no type for query_writable_namespaces
	// no type for _query_writable_namespaces_async
	// no type for _query_writable_namespaces_finish
	// no type for set_attribute
	// no type for set_attributes_from_info
	// no type for set_attributes_async
	// no type for set_attributes_finish
	// no type for read_fn
	// no type for read_async
	// no type for read_finish
	// no type for append_to
	// no type for append_to_async
	// no type for append_to_finish
	// no type for create
	// no type for create_async
	// no type for create_finish
	// no type for replace
	// no type for replace_async
	// no type for replace_finish
	// no type for delete_file
	// no type for delete_file_async
	// no type for delete_file_finish
	// no type for trash
	// no type for trash_async
	// no type for trash_finish
	// no type for make_directory
	// no type for make_directory_async
	// no type for make_directory_finish
	// no type for make_symbolic_link
	// no type for _make_symbolic_link_async
	// no type for _make_symbolic_link_finish
	// no type for copy
	// no type for copy_async
	// no type for copy_finish
	// no type for move
	// no type for _move_async
	// no type for _move_finish
	// no type for mount_mountable
	// no type for mount_mountable_finish
	// no type for unmount_mountable
	// no type for unmount_mountable_finish
	// no type for eject_mountable
	// no type for eject_mountable_finish
	// no type for mount_enclosing_volume
	// no type for mount_enclosing_volume_finish
	// no type for monitor_dir
	// no type for monitor_file
	// no type for open_readwrite
	// no type for open_readwrite_async
	// no type for open_readwrite_finish
	// no type for create_readwrite
	// no type for create_readwrite_async
	// no type for create_readwrite_finish
	// no type for replace_readwrite
	// no type for replace_readwrite_async
	// no type for replace_readwrite_finish
	// no type for start_mountable
	// no type for start_mountable_finish
	// no type for stop_mountable
	// no type for stop_mountable_finish
	SupportsThreadContexts bool
	// no type for unmount_mountable_with_operation
	// no type for unmount_mountable_with_operation_finish
	// no type for eject_mountable_with_operation
	// no type for eject_mountable_with_operation_finish
	// no type for poll_mountable
	// no type for poll_mountable_finish
	// no type for measure_disk_usage
	// no type for measure_disk_usage_async
	// no type for measure_disk_usage_finish
}

func FileIfaceNewFromC(u unsafe.Pointer) *FileIface {
	c := (*C.GFileIface)(u)
	if c == nil {
		return nil
	}

	g := &FileIface{
		SupportsThreadContexts: c.supports_thread_contexts == C.TRUE,
		native:                 c,
	}

	return g
}

func (recv *FileIface) ToC() unsafe.Pointer {
	recv.native.supports_thread_contexts =
		boolToGboolean(recv.SupportsThreadContexts)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileIface with another FileIface, and returns true if they represent the same GObject.
func (recv *FileIface) Equals(other *FileIface) bool {
	return other.ToC() == recv.ToC()
}

// FileInfoClass is a wrapper around the C record GFileInfoClass.
type FileInfoClass struct {
	native *C.GFileInfoClass
}

func FileInfoClassNewFromC(u unsafe.Pointer) *FileInfoClass {
	c := (*C.GFileInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInfoClass{native: c}

	return g
}

func (recv *FileInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileInfoClass with another FileInfoClass, and returns true if they represent the same GObject.
func (recv *FileInfoClass) Equals(other *FileInfoClass) bool {
	return other.ToC() == recv.ToC()
}

// FileInputStreamClass is a wrapper around the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native *C.GFileInputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileInputStreamClassNewFromC(u unsafe.Pointer) *FileInputStreamClass {
	c := (*C.GFileInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamClass{native: c}

	return g
}

func (recv *FileInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileInputStreamClass with another FileInputStreamClass, and returns true if they represent the same GObject.
func (recv *FileInputStreamClass) Equals(other *FileInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// FileInputStreamPrivate is a wrapper around the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native *C.GFileInputStreamPrivate
}

func FileInputStreamPrivateNewFromC(u unsafe.Pointer) *FileInputStreamPrivate {
	c := (*C.GFileInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamPrivate{native: c}

	return g
}

func (recv *FileInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileInputStreamPrivate with another FileInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *FileInputStreamPrivate) Equals(other *FileInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileMonitorClass is a wrapper around the C record GFileMonitorClass.
type FileMonitorClass struct {
	native *C.GFileMonitorClass
	// parent_class : record
	// no type for changed
	// no type for cancel
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileMonitorClassNewFromC(u unsafe.Pointer) *FileMonitorClass {
	c := (*C.GFileMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorClass{native: c}

	return g
}

func (recv *FileMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileMonitorClass with another FileMonitorClass, and returns true if they represent the same GObject.
func (recv *FileMonitorClass) Equals(other *FileMonitorClass) bool {
	return other.ToC() == recv.ToC()
}

// FileMonitorPrivate is a wrapper around the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native *C.GFileMonitorPrivate
}

func FileMonitorPrivateNewFromC(u unsafe.Pointer) *FileMonitorPrivate {
	c := (*C.GFileMonitorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorPrivate{native: c}

	return g
}

func (recv *FileMonitorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileMonitorPrivate with another FileMonitorPrivate, and returns true if they represent the same GObject.
func (recv *FileMonitorPrivate) Equals(other *FileMonitorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileOutputStreamClass is a wrapper around the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native *C.GFileOutputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileOutputStreamClassNewFromC(u unsafe.Pointer) *FileOutputStreamClass {
	c := (*C.GFileOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamClass{native: c}

	return g
}

func (recv *FileOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileOutputStreamClass with another FileOutputStreamClass, and returns true if they represent the same GObject.
func (recv *FileOutputStreamClass) Equals(other *FileOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// FileOutputStreamPrivate is a wrapper around the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native *C.GFileOutputStreamPrivate
}

func FileOutputStreamPrivateNewFromC(u unsafe.Pointer) *FileOutputStreamPrivate {
	c := (*C.GFileOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamPrivate{native: c}

	return g
}

func (recv *FileOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileOutputStreamPrivate with another FileOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *FileOutputStreamPrivate) Equals(other *FileOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FilenameCompleterClass is a wrapper around the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native *C.GFilenameCompleterClass
	// parent_class : record
	// no type for got_completion_data
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilenameCompleterClassNewFromC(u unsafe.Pointer) *FilenameCompleterClass {
	c := (*C.GFilenameCompleterClass)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleterClass{native: c}

	return g
}

func (recv *FilenameCompleterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilenameCompleterClass with another FilenameCompleterClass, and returns true if they represent the same GObject.
func (recv *FilenameCompleterClass) Equals(other *FilenameCompleterClass) bool {
	return other.ToC() == recv.ToC()
}

// FilterInputStreamClass is a wrapper around the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native *C.GFilterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterInputStreamClassNewFromC(u unsafe.Pointer) *FilterInputStreamClass {
	c := (*C.GFilterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStreamClass{native: c}

	return g
}

func (recv *FilterInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilterInputStreamClass with another FilterInputStreamClass, and returns true if they represent the same GObject.
func (recv *FilterInputStreamClass) Equals(other *FilterInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// FilterOutputStreamClass is a wrapper around the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native *C.GFilterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterOutputStreamClassNewFromC(u unsafe.Pointer) *FilterOutputStreamClass {
	c := (*C.GFilterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStreamClass{native: c}

	return g
}

func (recv *FilterOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilterOutputStreamClass with another FilterOutputStreamClass, and returns true if they represent the same GObject.
func (recv *FilterOutputStreamClass) Equals(other *FilterOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// IOExtension is a wrapper around the C record GIOExtension.
type IOExtension struct {
	native *C.GIOExtension
}

func IOExtensionNewFromC(u unsafe.Pointer) *IOExtension {
	c := (*C.GIOExtension)(u)
	if c == nil {
		return nil
	}

	g := &IOExtension{native: c}

	return g
}

func (recv *IOExtension) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOExtension with another IOExtension, and returns true if they represent the same GObject.
func (recv *IOExtension) Equals(other *IOExtension) bool {
	return other.ToC() == recv.ToC()
}

// GetName is a wrapper around the C function g_io_extension_get_name.
func (recv *IOExtension) GetName() string {
	retC := C.g_io_extension_get_name((*C.GIOExtension)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_io_extension_get_priority.
func (recv *IOExtension) GetPriority() int32 {
	retC := C.g_io_extension_get_priority((*C.GIOExtension)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetType is a wrapper around the C function g_io_extension_get_type.
func (recv *IOExtension) GetType() gobject.Type {
	retC := C.g_io_extension_get_type((*C.GIOExtension)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// RefClass is a wrapper around the C function g_io_extension_ref_class.
func (recv *IOExtension) RefClass() *gobject.TypeClass {
	retC := C.g_io_extension_ref_class((*C.GIOExtension)(recv.native))
	retGo := gobject.TypeClassNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOExtensionPoint is a wrapper around the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native *C.GIOExtensionPoint
}

func IOExtensionPointNewFromC(u unsafe.Pointer) *IOExtensionPoint {
	c := (*C.GIOExtensionPoint)(u)
	if c == nil {
		return nil
	}

	g := &IOExtensionPoint{native: c}

	return g
}

func (recv *IOExtensionPoint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOExtensionPoint with another IOExtensionPoint, and returns true if they represent the same GObject.
func (recv *IOExtensionPoint) Equals(other *IOExtensionPoint) bool {
	return other.ToC() == recv.ToC()
}

// IOExtensionPointImplement is a wrapper around the C function g_io_extension_point_implement.
func IOExtensionPointImplement(extensionPointName string, type_ gobject.Type, extensionName string, priority int32) *IOExtension {
	c_extension_point_name := C.CString(extensionPointName)
	defer C.free(unsafe.Pointer(c_extension_point_name))

	c_type := (C.GType)(type_)

	c_extension_name := C.CString(extensionName)
	defer C.free(unsafe.Pointer(c_extension_name))

	c_priority := (C.gint)(priority)

	retC := C.g_io_extension_point_implement(c_extension_point_name, c_type, c_extension_name, c_priority)
	retGo := IOExtensionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOExtensionPointLookup is a wrapper around the C function g_io_extension_point_lookup.
func IOExtensionPointLookup(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_lookup(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IOExtensionPointRegister is a wrapper around the C function g_io_extension_point_register.
func IOExtensionPointRegister(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_register(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExtensionByName is a wrapper around the C function g_io_extension_point_get_extension_by_name.
func (recv *IOExtensionPoint) GetExtensionByName(name string) *IOExtension {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_get_extension_by_name((*C.GIOExtensionPoint)(recv.native), c_name)
	retGo := IOExtensionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExtensions is a wrapper around the C function g_io_extension_point_get_extensions.
func (recv *IOExtensionPoint) GetExtensions() *glib.List {
	retC := C.g_io_extension_point_get_extensions((*C.GIOExtensionPoint)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRequiredType is a wrapper around the C function g_io_extension_point_get_required_type.
func (recv *IOExtensionPoint) GetRequiredType() gobject.Type {
	retC := C.g_io_extension_point_get_required_type((*C.GIOExtensionPoint)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// SetRequiredType is a wrapper around the C function g_io_extension_point_set_required_type.
func (recv *IOExtensionPoint) SetRequiredType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.g_io_extension_point_set_required_type((*C.GIOExtensionPoint)(recv.native), c_type)

	return
}

// IOModuleClass is a wrapper around the C record GIOModuleClass.
type IOModuleClass struct {
	native *C.GIOModuleClass
}

func IOModuleClassNewFromC(u unsafe.Pointer) *IOModuleClass {
	c := (*C.GIOModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &IOModuleClass{native: c}

	return g
}

func (recv *IOModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOModuleClass with another IOModuleClass, and returns true if they represent the same GObject.
func (recv *IOModuleClass) Equals(other *IOModuleClass) bool {
	return other.ToC() == recv.ToC()
}

// IOSchedulerJob is a wrapper around the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native *C.GIOSchedulerJob
}

func IOSchedulerJobNewFromC(u unsafe.Pointer) *IOSchedulerJob {
	c := (*C.GIOSchedulerJob)(u)
	if c == nil {
		return nil
	}

	g := &IOSchedulerJob{native: c}

	return g
}

func (recv *IOSchedulerJob) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOSchedulerJob with another IOSchedulerJob, and returns true if they represent the same GObject.
func (recv *IOSchedulerJob) Equals(other *IOSchedulerJob) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_io_scheduler_job_send_to_mainloop : unsupported parameter func : no type generator for GLib.SourceFunc (GSourceFunc) for param func

// Unsupported : g_io_scheduler_job_send_to_mainloop_async : unsupported parameter func : no type generator for GLib.SourceFunc (GSourceFunc) for param func

// IOStreamAdapter is a wrapper around the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native *C.GIOStreamAdapter
}

func IOStreamAdapterNewFromC(u unsafe.Pointer) *IOStreamAdapter {
	c := (*C.GIOStreamAdapter)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamAdapter{native: c}

	return g
}

func (recv *IOStreamAdapter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOStreamAdapter with another IOStreamAdapter, and returns true if they represent the same GObject.
func (recv *IOStreamAdapter) Equals(other *IOStreamAdapter) bool {
	return other.ToC() == recv.ToC()
}

// IOStreamClass is a wrapper around the C record GIOStreamClass.
type IOStreamClass struct {
	native *C.GIOStreamClass
	// parent_class : record
	// no type for get_input_stream
	// no type for get_output_stream
	// no type for close_fn
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func IOStreamClassNewFromC(u unsafe.Pointer) *IOStreamClass {
	c := (*C.GIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamClass{native: c}

	return g
}

func (recv *IOStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOStreamClass with another IOStreamClass, and returns true if they represent the same GObject.
func (recv *IOStreamClass) Equals(other *IOStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// IOStreamPrivate is a wrapper around the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native *C.GIOStreamPrivate
}

func IOStreamPrivateNewFromC(u unsafe.Pointer) *IOStreamPrivate {
	c := (*C.GIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamPrivate{native: c}

	return g
}

func (recv *IOStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IOStreamPrivate with another IOStreamPrivate, and returns true if they represent the same GObject.
func (recv *IOStreamPrivate) Equals(other *IOStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// IconIface is a wrapper around the C record GIconIface.
type IconIface struct {
	native *C.GIconIface
	// g_iface : record
	// no type for hash
	// no type for equal
	// no type for to_tokens
	// no type for from_tokens
	// no type for serialize
}

func IconIfaceNewFromC(u unsafe.Pointer) *IconIface {
	c := (*C.GIconIface)(u)
	if c == nil {
		return nil
	}

	g := &IconIface{native: c}

	return g
}

func (recv *IconIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this IconIface with another IconIface, and returns true if they represent the same GObject.
func (recv *IconIface) Equals(other *IconIface) bool {
	return other.ToC() == recv.ToC()
}

// InetAddressClass is a wrapper around the C record GInetAddressClass.
type InetAddressClass struct {
	native *C.GInetAddressClass
	// parent_class : record
	// no type for to_string
	// no type for to_bytes
}

func InetAddressClassNewFromC(u unsafe.Pointer) *InetAddressClass {
	c := (*C.GInetAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressClass{native: c}

	return g
}

func (recv *InetAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddressClass with another InetAddressClass, and returns true if they represent the same GObject.
func (recv *InetAddressClass) Equals(other *InetAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// InetAddressMaskClass is a wrapper around the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native *C.GInetAddressMaskClass
	// parent_class : record
}

func InetAddressMaskClassNewFromC(u unsafe.Pointer) *InetAddressMaskClass {
	c := (*C.GInetAddressMaskClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskClass{native: c}

	return g
}

func (recv *InetAddressMaskClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddressMaskClass with another InetAddressMaskClass, and returns true if they represent the same GObject.
func (recv *InetAddressMaskClass) Equals(other *InetAddressMaskClass) bool {
	return other.ToC() == recv.ToC()
}

// InetAddressMaskPrivate is a wrapper around the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native *C.GInetAddressMaskPrivate
}

func InetAddressMaskPrivateNewFromC(u unsafe.Pointer) *InetAddressMaskPrivate {
	c := (*C.GInetAddressMaskPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskPrivate{native: c}

	return g
}

func (recv *InetAddressMaskPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddressMaskPrivate with another InetAddressMaskPrivate, and returns true if they represent the same GObject.
func (recv *InetAddressMaskPrivate) Equals(other *InetAddressMaskPrivate) bool {
	return other.ToC() == recv.ToC()
}

// InetAddressPrivate is a wrapper around the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native *C.GInetAddressPrivate
}

func InetAddressPrivateNewFromC(u unsafe.Pointer) *InetAddressPrivate {
	c := (*C.GInetAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressPrivate{native: c}

	return g
}

func (recv *InetAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetAddressPrivate with another InetAddressPrivate, and returns true if they represent the same GObject.
func (recv *InetAddressPrivate) Equals(other *InetAddressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// InetSocketAddressClass is a wrapper around the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native *C.GInetSocketAddressClass
	// parent_class : record
}

func InetSocketAddressClassNewFromC(u unsafe.Pointer) *InetSocketAddressClass {
	c := (*C.GInetSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressClass{native: c}

	return g
}

func (recv *InetSocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetSocketAddressClass with another InetSocketAddressClass, and returns true if they represent the same GObject.
func (recv *InetSocketAddressClass) Equals(other *InetSocketAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// InetSocketAddressPrivate is a wrapper around the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native *C.GInetSocketAddressPrivate
}

func InetSocketAddressPrivateNewFromC(u unsafe.Pointer) *InetSocketAddressPrivate {
	c := (*C.GInetSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressPrivate{native: c}

	return g
}

func (recv *InetSocketAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InetSocketAddressPrivate with another InetSocketAddressPrivate, and returns true if they represent the same GObject.
func (recv *InetSocketAddressPrivate) Equals(other *InetSocketAddressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// InputStreamClass is a wrapper around the C record GInputStreamClass.
type InputStreamClass struct {
	native *C.GInputStreamClass
	// parent_class : record
	// no type for read_fn
	// no type for skip
	// no type for close_fn
	// no type for read_async
	// no type for read_finish
	// no type for skip_async
	// no type for skip_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func InputStreamClassNewFromC(u unsafe.Pointer) *InputStreamClass {
	c := (*C.GInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamClass{native: c}

	return g
}

func (recv *InputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InputStreamClass with another InputStreamClass, and returns true if they represent the same GObject.
func (recv *InputStreamClass) Equals(other *InputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// InputStreamPrivate is a wrapper around the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native *C.GInputStreamPrivate
}

func InputStreamPrivateNewFromC(u unsafe.Pointer) *InputStreamPrivate {
	c := (*C.GInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamPrivate{native: c}

	return g
}

func (recv *InputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InputStreamPrivate with another InputStreamPrivate, and returns true if they represent the same GObject.
func (recv *InputStreamPrivate) Equals(other *InputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ListStoreClass is a wrapper around the C record GListStoreClass.
type ListStoreClass struct {
	native *C.GListStoreClass
	// parent_class : record
}

func ListStoreClassNewFromC(u unsafe.Pointer) *ListStoreClass {
	c := (*C.GListStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &ListStoreClass{native: c}

	return g
}

func (recv *ListStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ListStoreClass with another ListStoreClass, and returns true if they represent the same GObject.
func (recv *ListStoreClass) Equals(other *ListStoreClass) bool {
	return other.ToC() == recv.ToC()
}

// LoadableIconIface is a wrapper around the C record GLoadableIconIface.
type LoadableIconIface struct {
	native *C.GLoadableIconIface
	// g_iface : record
	// no type for load
	// no type for load_async
	// no type for load_finish
}

func LoadableIconIfaceNewFromC(u unsafe.Pointer) *LoadableIconIface {
	c := (*C.GLoadableIconIface)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIconIface{native: c}

	return g
}

func (recv *LoadableIconIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LoadableIconIface with another LoadableIconIface, and returns true if they represent the same GObject.
func (recv *LoadableIconIface) Equals(other *LoadableIconIface) bool {
	return other.ToC() == recv.ToC()
}

// MemoryInputStreamClass is a wrapper around the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native *C.GMemoryInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryInputStreamClassNewFromC(u unsafe.Pointer) *MemoryInputStreamClass {
	c := (*C.GMemoryInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamClass{native: c}

	return g
}

func (recv *MemoryInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryInputStreamClass with another MemoryInputStreamClass, and returns true if they represent the same GObject.
func (recv *MemoryInputStreamClass) Equals(other *MemoryInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// MemoryInputStreamPrivate is a wrapper around the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native *C.GMemoryInputStreamPrivate
}

func MemoryInputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryInputStreamPrivate {
	c := (*C.GMemoryInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamPrivate{native: c}

	return g
}

func (recv *MemoryInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryInputStreamPrivate with another MemoryInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *MemoryInputStreamPrivate) Equals(other *MemoryInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MemoryOutputStreamClass is a wrapper around the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native *C.GMemoryOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryOutputStreamClassNewFromC(u unsafe.Pointer) *MemoryOutputStreamClass {
	c := (*C.GMemoryOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamClass{native: c}

	return g
}

func (recv *MemoryOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryOutputStreamClass with another MemoryOutputStreamClass, and returns true if they represent the same GObject.
func (recv *MemoryOutputStreamClass) Equals(other *MemoryOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// MemoryOutputStreamPrivate is a wrapper around the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native *C.GMemoryOutputStreamPrivate
}

func MemoryOutputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryOutputStreamPrivate {
	c := (*C.GMemoryOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamPrivate{native: c}

	return g
}

func (recv *MemoryOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MemoryOutputStreamPrivate with another MemoryOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *MemoryOutputStreamPrivate) Equals(other *MemoryOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuAttributeIterClass is a wrapper around the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native *C.GMenuAttributeIterClass
	// parent_class : record
	// no type for get_next
}

func MenuAttributeIterClassNewFromC(u unsafe.Pointer) *MenuAttributeIterClass {
	c := (*C.GMenuAttributeIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterClass{native: c}

	return g
}

func (recv *MenuAttributeIterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuAttributeIterClass with another MenuAttributeIterClass, and returns true if they represent the same GObject.
func (recv *MenuAttributeIterClass) Equals(other *MenuAttributeIterClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuAttributeIterPrivate is a wrapper around the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native *C.GMenuAttributeIterPrivate
}

func MenuAttributeIterPrivateNewFromC(u unsafe.Pointer) *MenuAttributeIterPrivate {
	c := (*C.GMenuAttributeIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterPrivate{native: c}

	return g
}

func (recv *MenuAttributeIterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuAttributeIterPrivate with another MenuAttributeIterPrivate, and returns true if they represent the same GObject.
func (recv *MenuAttributeIterPrivate) Equals(other *MenuAttributeIterPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuLinkIterClass is a wrapper around the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native *C.GMenuLinkIterClass
	// parent_class : record
	// no type for get_next
}

func MenuLinkIterClassNewFromC(u unsafe.Pointer) *MenuLinkIterClass {
	c := (*C.GMenuLinkIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterClass{native: c}

	return g
}

func (recv *MenuLinkIterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuLinkIterClass with another MenuLinkIterClass, and returns true if they represent the same GObject.
func (recv *MenuLinkIterClass) Equals(other *MenuLinkIterClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuLinkIterPrivate is a wrapper around the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native *C.GMenuLinkIterPrivate
}

func MenuLinkIterPrivateNewFromC(u unsafe.Pointer) *MenuLinkIterPrivate {
	c := (*C.GMenuLinkIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterPrivate{native: c}

	return g
}

func (recv *MenuLinkIterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuLinkIterPrivate with another MenuLinkIterPrivate, and returns true if they represent the same GObject.
func (recv *MenuLinkIterPrivate) Equals(other *MenuLinkIterPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MenuModelClass is a wrapper around the C record GMenuModelClass.
type MenuModelClass struct {
	native *C.GMenuModelClass
	// parent_class : record
	// no type for is_mutable
	// no type for get_n_items
	// no type for get_item_attributes
	// no type for iterate_item_attributes
	// no type for get_item_attribute_value
	// no type for get_item_links
	// no type for iterate_item_links
	// no type for get_item_link
}

func MenuModelClassNewFromC(u unsafe.Pointer) *MenuModelClass {
	c := (*C.GMenuModelClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelClass{native: c}

	return g
}

func (recv *MenuModelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuModelClass with another MenuModelClass, and returns true if they represent the same GObject.
func (recv *MenuModelClass) Equals(other *MenuModelClass) bool {
	return other.ToC() == recv.ToC()
}

// MenuModelPrivate is a wrapper around the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native *C.GMenuModelPrivate
}

func MenuModelPrivateNewFromC(u unsafe.Pointer) *MenuModelPrivate {
	c := (*C.GMenuModelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelPrivate{native: c}

	return g
}

func (recv *MenuModelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MenuModelPrivate with another MenuModelPrivate, and returns true if they represent the same GObject.
func (recv *MenuModelPrivate) Equals(other *MenuModelPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MountIface is a wrapper around the C record GMountIface.
type MountIface struct {
	native *C.GMountIface
	// g_iface : record
	// no type for changed
	// no type for unmounted
	// no type for get_root
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_volume
	// no type for get_drive
	// no type for can_unmount
	// no type for can_eject
	// no type for unmount
	// no type for unmount_finish
	// no type for eject
	// no type for eject_finish
	// no type for remount
	// no type for remount_finish
	// no type for guess_content_type
	// no type for guess_content_type_finish
	// no type for guess_content_type_sync
	// no type for pre_unmount
	// no type for unmount_with_operation
	// no type for unmount_with_operation_finish
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_default_location
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func MountIfaceNewFromC(u unsafe.Pointer) *MountIface {
	c := (*C.GMountIface)(u)
	if c == nil {
		return nil
	}

	g := &MountIface{native: c}

	return g
}

func (recv *MountIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MountIface with another MountIface, and returns true if they represent the same GObject.
func (recv *MountIface) Equals(other *MountIface) bool {
	return other.ToC() == recv.ToC()
}

// MountOperationClass is a wrapper around the C record GMountOperationClass.
type MountOperationClass struct {
	native *C.GMountOperationClass
	// parent_class : record
	// no type for ask_password
	// no type for ask_question
	// no type for reply
	// no type for aborted
	// no type for show_processes
	// no type for show_unmount_progress
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
}

func MountOperationClassNewFromC(u unsafe.Pointer) *MountOperationClass {
	c := (*C.GMountOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationClass{native: c}

	return g
}

func (recv *MountOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MountOperationClass with another MountOperationClass, and returns true if they represent the same GObject.
func (recv *MountOperationClass) Equals(other *MountOperationClass) bool {
	return other.ToC() == recv.ToC()
}

// MountOperationPrivate is a wrapper around the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native *C.GMountOperationPrivate
}

func MountOperationPrivateNewFromC(u unsafe.Pointer) *MountOperationPrivate {
	c := (*C.GMountOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationPrivate{native: c}

	return g
}

func (recv *MountOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MountOperationPrivate with another MountOperationPrivate, and returns true if they represent the same GObject.
func (recv *MountOperationPrivate) Equals(other *MountOperationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NativeSocketAddress is a wrapper around the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native *C.GNativeSocketAddress
}

func NativeSocketAddressNewFromC(u unsafe.Pointer) *NativeSocketAddress {
	c := (*C.GNativeSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &NativeSocketAddress{native: c}

	return g
}

func (recv *NativeSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeSocketAddress with another NativeSocketAddress, and returns true if they represent the same GObject.
func (recv *NativeSocketAddress) Equals(other *NativeSocketAddress) bool {
	return other.ToC() == recv.ToC()
}

// NativeVolumeMonitorClass is a wrapper around the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native *C.GNativeVolumeMonitorClass
	// parent_class : record
	// no type for get_mount_for_mount_path
}

func NativeVolumeMonitorClassNewFromC(u unsafe.Pointer) *NativeVolumeMonitorClass {
	c := (*C.GNativeVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitorClass{native: c}

	return g
}

func (recv *NativeVolumeMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeVolumeMonitorClass with another NativeVolumeMonitorClass, and returns true if they represent the same GObject.
func (recv *NativeVolumeMonitorClass) Equals(other *NativeVolumeMonitorClass) bool {
	return other.ToC() == recv.ToC()
}

// NetworkAddressClass is a wrapper around the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native *C.GNetworkAddressClass
	// parent_class : record
}

func NetworkAddressClassNewFromC(u unsafe.Pointer) *NetworkAddressClass {
	c := (*C.GNetworkAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressClass{native: c}

	return g
}

func (recv *NetworkAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkAddressClass with another NetworkAddressClass, and returns true if they represent the same GObject.
func (recv *NetworkAddressClass) Equals(other *NetworkAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// NetworkAddressPrivate is a wrapper around the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native *C.GNetworkAddressPrivate
}

func NetworkAddressPrivateNewFromC(u unsafe.Pointer) *NetworkAddressPrivate {
	c := (*C.GNetworkAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressPrivate{native: c}

	return g
}

func (recv *NetworkAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkAddressPrivate with another NetworkAddressPrivate, and returns true if they represent the same GObject.
func (recv *NetworkAddressPrivate) Equals(other *NetworkAddressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NetworkServiceClass is a wrapper around the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native *C.GNetworkServiceClass
	// parent_class : record
}

func NetworkServiceClassNewFromC(u unsafe.Pointer) *NetworkServiceClass {
	c := (*C.GNetworkServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServiceClass{native: c}

	return g
}

func (recv *NetworkServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkServiceClass with another NetworkServiceClass, and returns true if they represent the same GObject.
func (recv *NetworkServiceClass) Equals(other *NetworkServiceClass) bool {
	return other.ToC() == recv.ToC()
}

// NetworkServicePrivate is a wrapper around the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native *C.GNetworkServicePrivate
}

func NetworkServicePrivateNewFromC(u unsafe.Pointer) *NetworkServicePrivate {
	c := (*C.GNetworkServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServicePrivate{native: c}

	return g
}

func (recv *NetworkServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkServicePrivate with another NetworkServicePrivate, and returns true if they represent the same GObject.
func (recv *NetworkServicePrivate) Equals(other *NetworkServicePrivate) bool {
	return other.ToC() == recv.ToC()
}

// OutputStreamClass is a wrapper around the C record GOutputStreamClass.
type OutputStreamClass struct {
	native *C.GOutputStreamClass
	// parent_class : record
	// no type for write_fn
	// no type for splice
	// no type for flush
	// no type for close_fn
	// no type for write_async
	// no type for write_finish
	// no type for splice_async
	// no type for splice_finish
	// no type for flush_async
	// no type for flush_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
}

func OutputStreamClassNewFromC(u unsafe.Pointer) *OutputStreamClass {
	c := (*C.GOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamClass{native: c}

	return g
}

func (recv *OutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OutputStreamClass with another OutputStreamClass, and returns true if they represent the same GObject.
func (recv *OutputStreamClass) Equals(other *OutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// OutputStreamPrivate is a wrapper around the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native *C.GOutputStreamPrivate
}

func OutputStreamPrivateNewFromC(u unsafe.Pointer) *OutputStreamPrivate {
	c := (*C.GOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamPrivate{native: c}

	return g
}

func (recv *OutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OutputStreamPrivate with another OutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *OutputStreamPrivate) Equals(other *OutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PermissionClass is a wrapper around the C record GPermissionClass.
type PermissionClass struct {
	native *C.GPermissionClass
	// parent_class : record
	// no type for acquire
	// no type for acquire_async
	// no type for acquire_finish
	// no type for release
	// no type for release_async
	// no type for release_finish
	// no type for reserved
}

func PermissionClassNewFromC(u unsafe.Pointer) *PermissionClass {
	c := (*C.GPermissionClass)(u)
	if c == nil {
		return nil
	}

	g := &PermissionClass{native: c}

	return g
}

func (recv *PermissionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PermissionClass with another PermissionClass, and returns true if they represent the same GObject.
func (recv *PermissionClass) Equals(other *PermissionClass) bool {
	return other.ToC() == recv.ToC()
}

// PermissionPrivate is a wrapper around the C record GPermissionPrivate.
type PermissionPrivate struct {
	native *C.GPermissionPrivate
}

func PermissionPrivateNewFromC(u unsafe.Pointer) *PermissionPrivate {
	c := (*C.GPermissionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PermissionPrivate{native: c}

	return g
}

func (recv *PermissionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PermissionPrivate with another PermissionPrivate, and returns true if they represent the same GObject.
func (recv *PermissionPrivate) Equals(other *PermissionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ProxyAddressEnumeratorClass is a wrapper around the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native *C.GProxyAddressEnumeratorClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func ProxyAddressEnumeratorClassNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorClass {
	c := (*C.GProxyAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorClass{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyAddressEnumeratorClass with another ProxyAddressEnumeratorClass, and returns true if they represent the same GObject.
func (recv *ProxyAddressEnumeratorClass) Equals(other *ProxyAddressEnumeratorClass) bool {
	return other.ToC() == recv.ToC()
}

// ProxyAddressEnumeratorPrivate is a wrapper around the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native *C.GProxyAddressEnumeratorPrivate
}

func ProxyAddressEnumeratorPrivateNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	c := (*C.GProxyAddressEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorPrivate{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyAddressEnumeratorPrivate with another ProxyAddressEnumeratorPrivate, and returns true if they represent the same GObject.
func (recv *ProxyAddressEnumeratorPrivate) Equals(other *ProxyAddressEnumeratorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ProxyAddressPrivate is a wrapper around the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native *C.GProxyAddressPrivate
}

func ProxyAddressPrivateNewFromC(u unsafe.Pointer) *ProxyAddressPrivate {
	c := (*C.GProxyAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressPrivate{native: c}

	return g
}

func (recv *ProxyAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyAddressPrivate with another ProxyAddressPrivate, and returns true if they represent the same GObject.
func (recv *ProxyAddressPrivate) Equals(other *ProxyAddressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ProxyResolverInterface is a wrapper around the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native *C.GProxyResolverInterface
	// g_iface : record
	// no type for is_supported
	// no type for lookup
	// no type for lookup_async
	// no type for lookup_finish
}

func ProxyResolverInterfaceNewFromC(u unsafe.Pointer) *ProxyResolverInterface {
	c := (*C.GProxyResolverInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolverInterface{native: c}

	return g
}

func (recv *ProxyResolverInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ProxyResolverInterface with another ProxyResolverInterface, and returns true if they represent the same GObject.
func (recv *ProxyResolverInterface) Equals(other *ProxyResolverInterface) bool {
	return other.ToC() == recv.ToC()
}

// ResolverClass is a wrapper around the C record GResolverClass.
type ResolverClass struct {
	native *C.GResolverClass
	// parent_class : record
	// no type for reload
	// no type for lookup_by_name
	// no type for lookup_by_name_async
	// no type for lookup_by_name_finish
	// no type for lookup_by_address
	// no type for lookup_by_address_async
	// no type for lookup_by_address_finish
	// no type for lookup_service
	// no type for lookup_service_async
	// no type for lookup_service_finish
	// no type for lookup_records
	// no type for lookup_records_async
	// no type for lookup_records_finish
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func ResolverClassNewFromC(u unsafe.Pointer) *ResolverClass {
	c := (*C.GResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &ResolverClass{native: c}

	return g
}

func (recv *ResolverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ResolverClass with another ResolverClass, and returns true if they represent the same GObject.
func (recv *ResolverClass) Equals(other *ResolverClass) bool {
	return other.ToC() == recv.ToC()
}

// ResolverPrivate is a wrapper around the C record GResolverPrivate.
type ResolverPrivate struct {
	native *C.GResolverPrivate
}

func ResolverPrivateNewFromC(u unsafe.Pointer) *ResolverPrivate {
	c := (*C.GResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ResolverPrivate{native: c}

	return g
}

func (recv *ResolverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ResolverPrivate with another ResolverPrivate, and returns true if they represent the same GObject.
func (recv *ResolverPrivate) Equals(other *ResolverPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SeekableIface is a wrapper around the C record GSeekableIface.
type SeekableIface struct {
	native *C.GSeekableIface
	// g_iface : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
}

func SeekableIfaceNewFromC(u unsafe.Pointer) *SeekableIface {
	c := (*C.GSeekableIface)(u)
	if c == nil {
		return nil
	}

	g := &SeekableIface{native: c}

	return g
}

func (recv *SeekableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SeekableIface with another SeekableIface, and returns true if they represent the same GObject.
func (recv *SeekableIface) Equals(other *SeekableIface) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// SettingsClass is a wrapper around the C record GSettingsClass.
type SettingsClass struct {
	native *C.GSettingsClass
	// parent_class : record
	// no type for writable_changed
	// no type for changed
	// no type for writable_change_event
	// no type for change_event
	// no type for padding
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	c := (*C.GSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SettingsClass{native: c}

	return g
}

func (recv *SettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsClass with another SettingsClass, and returns true if they represent the same GObject.
func (recv *SettingsClass) Equals(other *SettingsClass) bool {
	return other.ToC() == recv.ToC()
}

// SettingsPrivate is a wrapper around the C record GSettingsPrivate.
type SettingsPrivate struct {
	native *C.GSettingsPrivate
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	c := (*C.GSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SettingsPrivate{native: c}

	return g
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsPrivate with another SettingsPrivate, and returns true if they represent the same GObject.
func (recv *SettingsPrivate) Equals(other *SettingsPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SettingsSchemaKey is a wrapper around the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native *C.GSettingsSchemaKey
}

func SettingsSchemaKeyNewFromC(u unsafe.Pointer) *SettingsSchemaKey {
	c := (*C.GSettingsSchemaKey)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchemaKey{native: c}

	return g
}

func (recv *SettingsSchemaKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsSchemaKey with another SettingsSchemaKey, and returns true if they represent the same GObject.
func (recv *SettingsSchemaKey) Equals(other *SettingsSchemaKey) bool {
	return other.ToC() == recv.ToC()
}

// SimpleActionGroupClass is a wrapper around the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native *C.GSimpleActionGroupClass
	// Private : parent_class
	// Private : padding
}

func SimpleActionGroupClassNewFromC(u unsafe.Pointer) *SimpleActionGroupClass {
	c := (*C.GSimpleActionGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupClass{native: c}

	return g
}

func (recv *SimpleActionGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleActionGroupClass with another SimpleActionGroupClass, and returns true if they represent the same GObject.
func (recv *SimpleActionGroupClass) Equals(other *SimpleActionGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// SimpleActionGroupPrivate is a wrapper around the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native *C.GSimpleActionGroupPrivate
}

func SimpleActionGroupPrivateNewFromC(u unsafe.Pointer) *SimpleActionGroupPrivate {
	c := (*C.GSimpleActionGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupPrivate{native: c}

	return g
}

func (recv *SimpleActionGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleActionGroupPrivate with another SimpleActionGroupPrivate, and returns true if they represent the same GObject.
func (recv *SimpleActionGroupPrivate) Equals(other *SimpleActionGroupPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SimpleAsyncResultClass is a wrapper around the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native *C.GSimpleAsyncResultClass
}

func SimpleAsyncResultClassNewFromC(u unsafe.Pointer) *SimpleAsyncResultClass {
	c := (*C.GSimpleAsyncResultClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResultClass{native: c}

	return g
}

func (recv *SimpleAsyncResultClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleAsyncResultClass with another SimpleAsyncResultClass, and returns true if they represent the same GObject.
func (recv *SimpleAsyncResultClass) Equals(other *SimpleAsyncResultClass) bool {
	return other.ToC() == recv.ToC()
}

// SimpleProxyResolverClass is a wrapper around the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native *C.GSimpleProxyResolverClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SimpleProxyResolverClassNewFromC(u unsafe.Pointer) *SimpleProxyResolverClass {
	c := (*C.GSimpleProxyResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverClass{native: c}

	return g
}

func (recv *SimpleProxyResolverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleProxyResolverClass with another SimpleProxyResolverClass, and returns true if they represent the same GObject.
func (recv *SimpleProxyResolverClass) Equals(other *SimpleProxyResolverClass) bool {
	return other.ToC() == recv.ToC()
}

// SimpleProxyResolverPrivate is a wrapper around the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native *C.GSimpleProxyResolverPrivate
}

func SimpleProxyResolverPrivateNewFromC(u unsafe.Pointer) *SimpleProxyResolverPrivate {
	c := (*C.GSimpleProxyResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverPrivate{native: c}

	return g
}

func (recv *SimpleProxyResolverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleProxyResolverPrivate with another SimpleProxyResolverPrivate, and returns true if they represent the same GObject.
func (recv *SimpleProxyResolverPrivate) Equals(other *SimpleProxyResolverPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketAddressClass is a wrapper around the C record GSocketAddressClass.
type SocketAddressClass struct {
	native *C.GSocketAddressClass
	// parent_class : record
	// no type for get_family
	// no type for get_native_size
	// no type for to_native
}

func SocketAddressClassNewFromC(u unsafe.Pointer) *SocketAddressClass {
	c := (*C.GSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressClass{native: c}

	return g
}

func (recv *SocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketAddressClass with another SocketAddressClass, and returns true if they represent the same GObject.
func (recv *SocketAddressClass) Equals(other *SocketAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketAddressEnumeratorClass is a wrapper around the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native *C.GSocketAddressEnumeratorClass
	// parent_class : record
	// no type for next
	// no type for next_async
	// no type for next_finish
}

func SocketAddressEnumeratorClassNewFromC(u unsafe.Pointer) *SocketAddressEnumeratorClass {
	c := (*C.GSocketAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumeratorClass{native: c}

	return g
}

func (recv *SocketAddressEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketAddressEnumeratorClass with another SocketAddressEnumeratorClass, and returns true if they represent the same GObject.
func (recv *SocketAddressEnumeratorClass) Equals(other *SocketAddressEnumeratorClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketClass is a wrapper around the C record GSocketClass.
type SocketClass struct {
	native *C.GSocketClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.GSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same GObject.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketClientClass is a wrapper around the C record GSocketClientClass.
type SocketClientClass struct {
	native *C.GSocketClientClass
	// parent_class : record
	// no type for event
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func SocketClientClassNewFromC(u unsafe.Pointer) *SocketClientClass {
	c := (*C.GSocketClientClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientClass{native: c}

	return g
}

func (recv *SocketClientClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClientClass with another SocketClientClass, and returns true if they represent the same GObject.
func (recv *SocketClientClass) Equals(other *SocketClientClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketClientPrivate is a wrapper around the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native *C.GSocketClientPrivate
}

func SocketClientPrivateNewFromC(u unsafe.Pointer) *SocketClientPrivate {
	c := (*C.GSocketClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientPrivate{native: c}

	return g
}

func (recv *SocketClientPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClientPrivate with another SocketClientPrivate, and returns true if they represent the same GObject.
func (recv *SocketClientPrivate) Equals(other *SocketClientPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnectableIface is a wrapper around the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native *C.GSocketConnectableIface
	// g_iface : record
	// no type for enumerate
	// no type for proxy_enumerate
	// no type for to_string
}

func SocketConnectableIfaceNewFromC(u unsafe.Pointer) *SocketConnectableIface {
	c := (*C.GSocketConnectableIface)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectableIface{native: c}

	return g
}

func (recv *SocketConnectableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnectableIface with another SocketConnectableIface, and returns true if they represent the same GObject.
func (recv *SocketConnectableIface) Equals(other *SocketConnectableIface) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnectionClass is a wrapper around the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native *C.GSocketConnectionClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketConnectionClassNewFromC(u unsafe.Pointer) *SocketConnectionClass {
	c := (*C.GSocketConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionClass{native: c}

	return g
}

func (recv *SocketConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnectionClass with another SocketConnectionClass, and returns true if they represent the same GObject.
func (recv *SocketConnectionClass) Equals(other *SocketConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnectionPrivate is a wrapper around the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native *C.GSocketConnectionPrivate
}

func SocketConnectionPrivateNewFromC(u unsafe.Pointer) *SocketConnectionPrivate {
	c := (*C.GSocketConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionPrivate{native: c}

	return g
}

func (recv *SocketConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnectionPrivate with another SocketConnectionPrivate, and returns true if they represent the same GObject.
func (recv *SocketConnectionPrivate) Equals(other *SocketConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketControlMessageClass is a wrapper around the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native *C.GSocketControlMessageClass
	// parent_class : record
	// no type for get_size
	// no type for get_level
	// no type for get_type
	// no type for serialize
	// no type for deserialize
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SocketControlMessageClassNewFromC(u unsafe.Pointer) *SocketControlMessageClass {
	c := (*C.GSocketControlMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessageClass{native: c}

	return g
}

func (recv *SocketControlMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketControlMessageClass with another SocketControlMessageClass, and returns true if they represent the same GObject.
func (recv *SocketControlMessageClass) Equals(other *SocketControlMessageClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketControlMessagePrivate is a wrapper around the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native *C.GSocketControlMessagePrivate
}

func SocketControlMessagePrivateNewFromC(u unsafe.Pointer) *SocketControlMessagePrivate {
	c := (*C.GSocketControlMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessagePrivate{native: c}

	return g
}

func (recv *SocketControlMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketControlMessagePrivate with another SocketControlMessagePrivate, and returns true if they represent the same GObject.
func (recv *SocketControlMessagePrivate) Equals(other *SocketControlMessagePrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketListenerClass is a wrapper around the C record GSocketListenerClass.
type SocketListenerClass struct {
	native *C.GSocketListenerClass
	// parent_class : record
	// no type for changed
	// no type for event
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketListenerClassNewFromC(u unsafe.Pointer) *SocketListenerClass {
	c := (*C.GSocketListenerClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerClass{native: c}

	return g
}

func (recv *SocketListenerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketListenerClass with another SocketListenerClass, and returns true if they represent the same GObject.
func (recv *SocketListenerClass) Equals(other *SocketListenerClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketListenerPrivate is a wrapper around the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native *C.GSocketListenerPrivate
}

func SocketListenerPrivateNewFromC(u unsafe.Pointer) *SocketListenerPrivate {
	c := (*C.GSocketListenerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerPrivate{native: c}

	return g
}

func (recv *SocketListenerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketListenerPrivate with another SocketListenerPrivate, and returns true if they represent the same GObject.
func (recv *SocketListenerPrivate) Equals(other *SocketListenerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketPrivate is a wrapper around the C record GSocketPrivate.
type SocketPrivate struct {
	native *C.GSocketPrivate
}

func SocketPrivateNewFromC(u unsafe.Pointer) *SocketPrivate {
	c := (*C.GSocketPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketPrivate{native: c}

	return g
}

func (recv *SocketPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketPrivate with another SocketPrivate, and returns true if they represent the same GObject.
func (recv *SocketPrivate) Equals(other *SocketPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SocketServiceClass is a wrapper around the C record GSocketServiceClass.
type SocketServiceClass struct {
	native *C.GSocketServiceClass
	// parent_class : record
	// no type for incoming
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketServiceClassNewFromC(u unsafe.Pointer) *SocketServiceClass {
	c := (*C.GSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketServiceClass{native: c}

	return g
}

func (recv *SocketServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketServiceClass with another SocketServiceClass, and returns true if they represent the same GObject.
func (recv *SocketServiceClass) Equals(other *SocketServiceClass) bool {
	return other.ToC() == recv.ToC()
}

// SocketServicePrivate is a wrapper around the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native *C.GSocketServicePrivate
}

func SocketServicePrivateNewFromC(u unsafe.Pointer) *SocketServicePrivate {
	c := (*C.GSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketServicePrivate{native: c}

	return g
}

func (recv *SocketServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketServicePrivate with another SocketServicePrivate, and returns true if they represent the same GObject.
func (recv *SocketServicePrivate) Equals(other *SocketServicePrivate) bool {
	return other.ToC() == recv.ToC()
}

// SrvTarget is a wrapper around the C record GSrvTarget.
type SrvTarget struct {
	native *C.GSrvTarget
}

func SrvTargetNewFromC(u unsafe.Pointer) *SrvTarget {
	c := (*C.GSrvTarget)(u)
	if c == nil {
		return nil
	}

	g := &SrvTarget{native: c}

	return g
}

func (recv *SrvTarget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SrvTarget with another SrvTarget, and returns true if they represent the same GObject.
func (recv *SrvTarget) Equals(other *SrvTarget) bool {
	return other.ToC() == recv.ToC()
}

// StaticResource is a wrapper around the C record GStaticResource.
type StaticResource struct {
	native *C.GStaticResource
	// Private : data
	// Private : data_len
	// Private : resource
	// Private : next
	// Private : padding
}

func StaticResourceNewFromC(u unsafe.Pointer) *StaticResource {
	c := (*C.GStaticResource)(u)
	if c == nil {
		return nil
	}

	g := &StaticResource{native: c}

	return g
}

func (recv *StaticResource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StaticResource with another StaticResource, and returns true if they represent the same GObject.
func (recv *StaticResource) Equals(other *StaticResource) bool {
	return other.ToC() == recv.ToC()
}

// TaskClass is a wrapper around the C record GTaskClass.
type TaskClass struct {
	native *C.GTaskClass
}

func TaskClassNewFromC(u unsafe.Pointer) *TaskClass {
	c := (*C.GTaskClass)(u)
	if c == nil {
		return nil
	}

	g := &TaskClass{native: c}

	return g
}

func (recv *TaskClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TaskClass with another TaskClass, and returns true if they represent the same GObject.
func (recv *TaskClass) Equals(other *TaskClass) bool {
	return other.ToC() == recv.ToC()
}

// TcpConnectionClass is a wrapper around the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native *C.GTcpConnectionClass
	// parent_class : record
}

func TcpConnectionClassNewFromC(u unsafe.Pointer) *TcpConnectionClass {
	c := (*C.GTcpConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionClass{native: c}

	return g
}

func (recv *TcpConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpConnectionClass with another TcpConnectionClass, and returns true if they represent the same GObject.
func (recv *TcpConnectionClass) Equals(other *TcpConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// TcpConnectionPrivate is a wrapper around the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native *C.GTcpConnectionPrivate
}

func TcpConnectionPrivateNewFromC(u unsafe.Pointer) *TcpConnectionPrivate {
	c := (*C.GTcpConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionPrivate{native: c}

	return g
}

func (recv *TcpConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpConnectionPrivate with another TcpConnectionPrivate, and returns true if they represent the same GObject.
func (recv *TcpConnectionPrivate) Equals(other *TcpConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TcpWrapperConnectionClass is a wrapper around the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native *C.GTcpWrapperConnectionClass
	// parent_class : record
}

func TcpWrapperConnectionClassNewFromC(u unsafe.Pointer) *TcpWrapperConnectionClass {
	c := (*C.GTcpWrapperConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionClass{native: c}

	return g
}

func (recv *TcpWrapperConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpWrapperConnectionClass with another TcpWrapperConnectionClass, and returns true if they represent the same GObject.
func (recv *TcpWrapperConnectionClass) Equals(other *TcpWrapperConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// TcpWrapperConnectionPrivate is a wrapper around the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native *C.GTcpWrapperConnectionPrivate
}

func TcpWrapperConnectionPrivateNewFromC(u unsafe.Pointer) *TcpWrapperConnectionPrivate {
	c := (*C.GTcpWrapperConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionPrivate{native: c}

	return g
}

func (recv *TcpWrapperConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpWrapperConnectionPrivate with another TcpWrapperConnectionPrivate, and returns true if they represent the same GObject.
func (recv *TcpWrapperConnectionPrivate) Equals(other *TcpWrapperConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ThemedIconClass is a wrapper around the C record GThemedIconClass.
type ThemedIconClass struct {
	native *C.GThemedIconClass
}

func ThemedIconClassNewFromC(u unsafe.Pointer) *ThemedIconClass {
	c := (*C.GThemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIconClass{native: c}

	return g
}

func (recv *ThemedIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThemedIconClass with another ThemedIconClass, and returns true if they represent the same GObject.
func (recv *ThemedIconClass) Equals(other *ThemedIconClass) bool {
	return other.ToC() == recv.ToC()
}

// ThreadedSocketServiceClass is a wrapper around the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native *C.GThreadedSocketServiceClass
	// parent_class : record
	// no type for run
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ThreadedSocketServiceClassNewFromC(u unsafe.Pointer) *ThreadedSocketServiceClass {
	c := (*C.GThreadedSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServiceClass{native: c}

	return g
}

func (recv *ThreadedSocketServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadedSocketServiceClass with another ThreadedSocketServiceClass, and returns true if they represent the same GObject.
func (recv *ThreadedSocketServiceClass) Equals(other *ThreadedSocketServiceClass) bool {
	return other.ToC() == recv.ToC()
}

// ThreadedSocketServicePrivate is a wrapper around the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native *C.GThreadedSocketServicePrivate
}

func ThreadedSocketServicePrivateNewFromC(u unsafe.Pointer) *ThreadedSocketServicePrivate {
	c := (*C.GThreadedSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServicePrivate{native: c}

	return g
}

func (recv *ThreadedSocketServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadedSocketServicePrivate with another ThreadedSocketServicePrivate, and returns true if they represent the same GObject.
func (recv *ThreadedSocketServicePrivate) Equals(other *ThreadedSocketServicePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TlsCertificateClass is a wrapper around the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native *C.GTlsCertificateClass
	// parent_class : record
	// no type for verify
	// Private : padding
}

func TlsCertificateClassNewFromC(u unsafe.Pointer) *TlsCertificateClass {
	c := (*C.GTlsCertificateClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificateClass{native: c}

	return g
}

func (recv *TlsCertificateClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsCertificateClass with another TlsCertificateClass, and returns true if they represent the same GObject.
func (recv *TlsCertificateClass) Equals(other *TlsCertificateClass) bool {
	return other.ToC() == recv.ToC()
}

// TlsCertificatePrivate is a wrapper around the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native *C.GTlsCertificatePrivate
}

func TlsCertificatePrivateNewFromC(u unsafe.Pointer) *TlsCertificatePrivate {
	c := (*C.GTlsCertificatePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificatePrivate{native: c}

	return g
}

func (recv *TlsCertificatePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsCertificatePrivate with another TlsCertificatePrivate, and returns true if they represent the same GObject.
func (recv *TlsCertificatePrivate) Equals(other *TlsCertificatePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TlsConnectionClass is a wrapper around the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native *C.GTlsConnectionClass
	// parent_class : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// Private : padding
}

func TlsConnectionClassNewFromC(u unsafe.Pointer) *TlsConnectionClass {
	c := (*C.GTlsConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionClass{native: c}

	return g
}

func (recv *TlsConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsConnectionClass with another TlsConnectionClass, and returns true if they represent the same GObject.
func (recv *TlsConnectionClass) Equals(other *TlsConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// TlsConnectionPrivate is a wrapper around the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native *C.GTlsConnectionPrivate
}

func TlsConnectionPrivateNewFromC(u unsafe.Pointer) *TlsConnectionPrivate {
	c := (*C.GTlsConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionPrivate{native: c}

	return g
}

func (recv *TlsConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsConnectionPrivate with another TlsConnectionPrivate, and returns true if they represent the same GObject.
func (recv *TlsConnectionPrivate) Equals(other *TlsConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TlsDatabasePrivate is a wrapper around the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native *C.GTlsDatabasePrivate
}

func TlsDatabasePrivateNewFromC(u unsafe.Pointer) *TlsDatabasePrivate {
	c := (*C.GTlsDatabasePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabasePrivate{native: c}

	return g
}

func (recv *TlsDatabasePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsDatabasePrivate with another TlsDatabasePrivate, and returns true if they represent the same GObject.
func (recv *TlsDatabasePrivate) Equals(other *TlsDatabasePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TlsFileDatabaseInterface is a wrapper around the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native *C.GTlsFileDatabaseInterface
	// g_iface : record
	// Private : padding
}

func TlsFileDatabaseInterfaceNewFromC(u unsafe.Pointer) *TlsFileDatabaseInterface {
	c := (*C.GTlsFileDatabaseInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsFileDatabaseInterface{native: c}

	return g
}

func (recv *TlsFileDatabaseInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsFileDatabaseInterface with another TlsFileDatabaseInterface, and returns true if they represent the same GObject.
func (recv *TlsFileDatabaseInterface) Equals(other *TlsFileDatabaseInterface) bool {
	return other.ToC() == recv.ToC()
}

// TlsInteractionPrivate is a wrapper around the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native *C.GTlsInteractionPrivate
}

func TlsInteractionPrivateNewFromC(u unsafe.Pointer) *TlsInteractionPrivate {
	c := (*C.GTlsInteractionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteractionPrivate{native: c}

	return g
}

func (recv *TlsInteractionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsInteractionPrivate with another TlsInteractionPrivate, and returns true if they represent the same GObject.
func (recv *TlsInteractionPrivate) Equals(other *TlsInteractionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// TlsPasswordClass is a wrapper around the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native *C.GTlsPasswordClass
	// parent_class : record
	// no type for get_value
	// no type for set_value
	// no type for get_default_warning
	// Private : padding
}

func TlsPasswordClassNewFromC(u unsafe.Pointer) *TlsPasswordClass {
	c := (*C.GTlsPasswordClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordClass{native: c}

	return g
}

func (recv *TlsPasswordClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsPasswordClass with another TlsPasswordClass, and returns true if they represent the same GObject.
func (recv *TlsPasswordClass) Equals(other *TlsPasswordClass) bool {
	return other.ToC() == recv.ToC()
}

// TlsPasswordPrivate is a wrapper around the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native *C.GTlsPasswordPrivate
}

func TlsPasswordPrivateNewFromC(u unsafe.Pointer) *TlsPasswordPrivate {
	c := (*C.GTlsPasswordPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordPrivate{native: c}

	return g
}

func (recv *TlsPasswordPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsPasswordPrivate with another TlsPasswordPrivate, and returns true if they represent the same GObject.
func (recv *TlsPasswordPrivate) Equals(other *TlsPasswordPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixConnectionClass is a wrapper around the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native *C.GUnixConnectionClass
	// parent_class : record
}

func UnixConnectionClassNewFromC(u unsafe.Pointer) *UnixConnectionClass {
	c := (*C.GUnixConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionClass{native: c}

	return g
}

func (recv *UnixConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixConnectionClass with another UnixConnectionClass, and returns true if they represent the same GObject.
func (recv *UnixConnectionClass) Equals(other *UnixConnectionClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixConnectionPrivate is a wrapper around the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native *C.GUnixConnectionPrivate
}

func UnixConnectionPrivateNewFromC(u unsafe.Pointer) *UnixConnectionPrivate {
	c := (*C.GUnixConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionPrivate{native: c}

	return g
}

func (recv *UnixConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixConnectionPrivate with another UnixConnectionPrivate, and returns true if they represent the same GObject.
func (recv *UnixConnectionPrivate) Equals(other *UnixConnectionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixCredentialsMessagePrivate is a wrapper around the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native *C.GUnixCredentialsMessagePrivate
}

func UnixCredentialsMessagePrivateNewFromC(u unsafe.Pointer) *UnixCredentialsMessagePrivate {
	c := (*C.GUnixCredentialsMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessagePrivate{native: c}

	return g
}

func (recv *UnixCredentialsMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixCredentialsMessagePrivate with another UnixCredentialsMessagePrivate, and returns true if they represent the same GObject.
func (recv *UnixCredentialsMessagePrivate) Equals(other *UnixCredentialsMessagePrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixFDListClass is a wrapper around the C record GUnixFDListClass.
type UnixFDListClass struct {
	native *C.GUnixFDListClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixFDListClassNewFromC(u unsafe.Pointer) *UnixFDListClass {
	c := (*C.GUnixFDListClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListClass{native: c}

	return g
}

func (recv *UnixFDListClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDListClass with another UnixFDListClass, and returns true if they represent the same GObject.
func (recv *UnixFDListClass) Equals(other *UnixFDListClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixFDListPrivate is a wrapper around the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native *C.GUnixFDListPrivate
}

func UnixFDListPrivateNewFromC(u unsafe.Pointer) *UnixFDListPrivate {
	c := (*C.GUnixFDListPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListPrivate{native: c}

	return g
}

func (recv *UnixFDListPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDListPrivate with another UnixFDListPrivate, and returns true if they represent the same GObject.
func (recv *UnixFDListPrivate) Equals(other *UnixFDListPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixFDMessageClass is a wrapper around the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native *C.GUnixFDMessageClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func UnixFDMessageClassNewFromC(u unsafe.Pointer) *UnixFDMessageClass {
	c := (*C.GUnixFDMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessageClass{native: c}

	return g
}

func (recv *UnixFDMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDMessageClass with another UnixFDMessageClass, and returns true if they represent the same GObject.
func (recv *UnixFDMessageClass) Equals(other *UnixFDMessageClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixFDMessagePrivate is a wrapper around the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native *C.GUnixFDMessagePrivate
}

func UnixFDMessagePrivateNewFromC(u unsafe.Pointer) *UnixFDMessagePrivate {
	c := (*C.GUnixFDMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessagePrivate{native: c}

	return g
}

func (recv *UnixFDMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixFDMessagePrivate with another UnixFDMessagePrivate, and returns true if they represent the same GObject.
func (recv *UnixFDMessagePrivate) Equals(other *UnixFDMessagePrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixInputStreamClass is a wrapper around the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native *C.GUnixInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixInputStreamClassNewFromC(u unsafe.Pointer) *UnixInputStreamClass {
	c := (*C.GUnixInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamClass{native: c}

	return g
}

func (recv *UnixInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixInputStreamClass with another UnixInputStreamClass, and returns true if they represent the same GObject.
func (recv *UnixInputStreamClass) Equals(other *UnixInputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixInputStreamPrivate is a wrapper around the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native *C.GUnixInputStreamPrivate
}

func UnixInputStreamPrivateNewFromC(u unsafe.Pointer) *UnixInputStreamPrivate {
	c := (*C.GUnixInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamPrivate{native: c}

	return g
}

func (recv *UnixInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixInputStreamPrivate with another UnixInputStreamPrivate, and returns true if they represent the same GObject.
func (recv *UnixInputStreamPrivate) Equals(other *UnixInputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixMountEntry is a wrapper around the C record GUnixMountEntry.
type UnixMountEntry struct {
	native *C.GUnixMountEntry
}

func UnixMountEntryNewFromC(u unsafe.Pointer) *UnixMountEntry {
	c := (*C.GUnixMountEntry)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountEntry{native: c}

	return g
}

func (recv *UnixMountEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixMountEntry with another UnixMountEntry, and returns true if they represent the same GObject.
func (recv *UnixMountEntry) Equals(other *UnixMountEntry) bool {
	return other.ToC() == recv.ToC()
}

// UnixMountMonitorClass is a wrapper around the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native *C.GUnixMountMonitorClass
}

func UnixMountMonitorClassNewFromC(u unsafe.Pointer) *UnixMountMonitorClass {
	c := (*C.GUnixMountMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitorClass{native: c}

	return g
}

func (recv *UnixMountMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixMountMonitorClass with another UnixMountMonitorClass, and returns true if they represent the same GObject.
func (recv *UnixMountMonitorClass) Equals(other *UnixMountMonitorClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixMountPoint is a wrapper around the C record GUnixMountPoint.
type UnixMountPoint struct {
	native *C.GUnixMountPoint
}

func UnixMountPointNewFromC(u unsafe.Pointer) *UnixMountPoint {
	c := (*C.GUnixMountPoint)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountPoint{native: c}

	return g
}

func (recv *UnixMountPoint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixMountPoint with another UnixMountPoint, and returns true if they represent the same GObject.
func (recv *UnixMountPoint) Equals(other *UnixMountPoint) bool {
	return other.ToC() == recv.ToC()
}

// Compare is a wrapper around the C function g_unix_mount_point_compare.
func (recv *UnixMountPoint) Compare(mount2 *UnixMountPoint) int32 {
	c_mount2 := (*C.GUnixMountPoint)(C.NULL)
	if mount2 != nil {
		c_mount2 = (*C.GUnixMountPoint)(mount2.ToC())
	}

	retC := C.g_unix_mount_point_compare((*C.GUnixMountPoint)(recv.native), c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// Free is a wrapper around the C function g_unix_mount_point_free.
func (recv *UnixMountPoint) Free() {
	C.g_unix_mount_point_free((*C.GUnixMountPoint)(recv.native))

	return
}

// GetDevicePath is a wrapper around the C function g_unix_mount_point_get_device_path.
func (recv *UnixMountPoint) GetDevicePath() string {
	retC := C.g_unix_mount_point_get_device_path((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFsType is a wrapper around the C function g_unix_mount_point_get_fs_type.
func (recv *UnixMountPoint) GetFsType() string {
	retC := C.g_unix_mount_point_get_fs_type((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMountPath is a wrapper around the C function g_unix_mount_point_get_mount_path.
func (recv *UnixMountPoint) GetMountPath() string {
	retC := C.g_unix_mount_point_get_mount_path((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GuessCanEject is a wrapper around the C function g_unix_mount_point_guess_can_eject.
func (recv *UnixMountPoint) GuessCanEject() bool {
	retC := C.g_unix_mount_point_guess_can_eject((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GuessIcon is a wrapper around the C function g_unix_mount_point_guess_icon.
func (recv *UnixMountPoint) GuessIcon() *Icon {
	retC := C.g_unix_mount_point_guess_icon((*C.GUnixMountPoint)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GuessName is a wrapper around the C function g_unix_mount_point_guess_name.
func (recv *UnixMountPoint) GuessName() string {
	retC := C.g_unix_mount_point_guess_name((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsLoopback is a wrapper around the C function g_unix_mount_point_is_loopback.
func (recv *UnixMountPoint) IsLoopback() bool {
	retC := C.g_unix_mount_point_is_loopback((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsReadonly is a wrapper around the C function g_unix_mount_point_is_readonly.
func (recv *UnixMountPoint) IsReadonly() bool {
	retC := C.g_unix_mount_point_is_readonly((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsUserMountable is a wrapper around the C function g_unix_mount_point_is_user_mountable.
func (recv *UnixMountPoint) IsUserMountable() bool {
	retC := C.g_unix_mount_point_is_user_mountable((*C.GUnixMountPoint)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// UnixOutputStreamClass is a wrapper around the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native *C.GUnixOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixOutputStreamClassNewFromC(u unsafe.Pointer) *UnixOutputStreamClass {
	c := (*C.GUnixOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamClass{native: c}

	return g
}

func (recv *UnixOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixOutputStreamClass with another UnixOutputStreamClass, and returns true if they represent the same GObject.
func (recv *UnixOutputStreamClass) Equals(other *UnixOutputStreamClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixOutputStreamPrivate is a wrapper around the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native *C.GUnixOutputStreamPrivate
}

func UnixOutputStreamPrivateNewFromC(u unsafe.Pointer) *UnixOutputStreamPrivate {
	c := (*C.GUnixOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamPrivate{native: c}

	return g
}

func (recv *UnixOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixOutputStreamPrivate with another UnixOutputStreamPrivate, and returns true if they represent the same GObject.
func (recv *UnixOutputStreamPrivate) Equals(other *UnixOutputStreamPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UnixSocketAddressClass is a wrapper around the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native *C.GUnixSocketAddressClass
	// parent_class : record
}

func UnixSocketAddressClassNewFromC(u unsafe.Pointer) *UnixSocketAddressClass {
	c := (*C.GUnixSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressClass{native: c}

	return g
}

func (recv *UnixSocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixSocketAddressClass with another UnixSocketAddressClass, and returns true if they represent the same GObject.
func (recv *UnixSocketAddressClass) Equals(other *UnixSocketAddressClass) bool {
	return other.ToC() == recv.ToC()
}

// UnixSocketAddressPrivate is a wrapper around the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native *C.GUnixSocketAddressPrivate
}

func UnixSocketAddressPrivateNewFromC(u unsafe.Pointer) *UnixSocketAddressPrivate {
	c := (*C.GUnixSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressPrivate{native: c}

	return g
}

func (recv *UnixSocketAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UnixSocketAddressPrivate with another UnixSocketAddressPrivate, and returns true if they represent the same GObject.
func (recv *UnixSocketAddressPrivate) Equals(other *UnixSocketAddressPrivate) bool {
	return other.ToC() == recv.ToC()
}

// VfsClass is a wrapper around the C record GVfsClass.
type VfsClass struct {
	native *C.GVfsClass
	// parent_class : record
	// no type for is_active
	// no type for get_file_for_path
	// no type for get_file_for_uri
	// no type for get_supported_uri_schemes
	// no type for parse_name
	// no type for local_file_add_info
	// no type for add_writable_namespaces
	// no type for local_file_set_attributes
	// no type for local_file_removed
	// no type for local_file_moved
	// no type for deserialize_icon
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VfsClassNewFromC(u unsafe.Pointer) *VfsClass {
	c := (*C.GVfsClass)(u)
	if c == nil {
		return nil
	}

	g := &VfsClass{native: c}

	return g
}

func (recv *VfsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VfsClass with another VfsClass, and returns true if they represent the same GObject.
func (recv *VfsClass) Equals(other *VfsClass) bool {
	return other.ToC() == recv.ToC()
}

// VolumeIface is a wrapper around the C record GVolumeIface.
type VolumeIface struct {
	native *C.GVolumeIface
	// g_iface : record
	// no type for changed
	// no type for removed
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_drive
	// no type for get_mount
	// no type for can_mount
	// no type for can_eject
	// no type for mount_fn
	// no type for mount_finish
	// no type for eject
	// no type for eject_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for should_automount
	// no type for get_activation_root
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func VolumeIfaceNewFromC(u unsafe.Pointer) *VolumeIface {
	c := (*C.GVolumeIface)(u)
	if c == nil {
		return nil
	}

	g := &VolumeIface{native: c}

	return g
}

func (recv *VolumeIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VolumeIface with another VolumeIface, and returns true if they represent the same GObject.
func (recv *VolumeIface) Equals(other *VolumeIface) bool {
	return other.ToC() == recv.ToC()
}

// VolumeMonitorClass is a wrapper around the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native *C.GVolumeMonitorClass
	// parent_class : record
	// no type for volume_added
	// no type for volume_removed
	// no type for volume_changed
	// no type for mount_added
	// no type for mount_removed
	// no type for mount_pre_unmount
	// no type for mount_changed
	// no type for drive_connected
	// no type for drive_disconnected
	// no type for drive_changed
	// no type for is_supported
	// no type for get_connected_drives
	// no type for get_volumes
	// no type for get_mounts
	// no type for get_volume_for_uuid
	// no type for get_mount_for_uuid
	// no type for adopt_orphan_mount
	// no type for drive_eject_button
	// no type for drive_stop_button
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VolumeMonitorClassNewFromC(u unsafe.Pointer) *VolumeMonitorClass {
	c := (*C.GVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitorClass{native: c}

	return g
}

func (recv *VolumeMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this VolumeMonitorClass with another VolumeMonitorClass, and returns true if they represent the same GObject.
func (recv *VolumeMonitorClass) Equals(other *VolumeMonitorClass) bool {
	return other.ToC() == recv.ToC()
}

// ZlibCompressorClass is a wrapper around the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native *C.GZlibCompressorClass
	// parent_class : record
}

func ZlibCompressorClassNewFromC(u unsafe.Pointer) *ZlibCompressorClass {
	c := (*C.GZlibCompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressorClass{native: c}

	return g
}

func (recv *ZlibCompressorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ZlibCompressorClass with another ZlibCompressorClass, and returns true if they represent the same GObject.
func (recv *ZlibCompressorClass) Equals(other *ZlibCompressorClass) bool {
	return other.ToC() == recv.ToC()
}

// ZlibDecompressorClass is a wrapper around the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native *C.GZlibDecompressorClass
	// parent_class : record
}

func ZlibDecompressorClassNewFromC(u unsafe.Pointer) *ZlibDecompressorClass {
	c := (*C.GZlibDecompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressorClass{native: c}

	return g
}

func (recv *ZlibDecompressorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ZlibDecompressorClass with another ZlibDecompressorClass, and returns true if they represent the same GObject.
func (recv *ZlibDecompressorClass) Equals(other *ZlibDecompressorClass) bool {
	return other.ToC() == recv.ToC()
}
