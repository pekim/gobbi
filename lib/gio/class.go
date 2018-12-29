// This is a generated file - DO NOT EDIT

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

	void filenamecompleter_gotCompletionDataHandler(GObject *, gpointer);

	static gulong FilenameCompleter_signal_connect_got_completion_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "got-completion-data", G_CALLBACK(filenamecompleter_gotCompletionDataHandler), data);
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
import "C"

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
	c_buffer := &buffer[0]

	c_offset := (C.gsize)(offset)

	c_count := (C.gsize)(len(buffer))

	retC := C.g_buffered_input_stream_peek((*C.GBufferedInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_offset, c_count)
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
func (recv *FileInfo) GetSize() uint64 {
	retC := C.g_file_info_get_size((*C.GFileInfo)(recv.native))
	retGo := (uint64)(retC)

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

// SetAttribute is a wrapper around the C function g_file_info_set_attribute.
func (recv *FileInfo) SetAttribute(attribute string, type_ FileAttributeType, valueP uintptr) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_type := (C.GFileAttributeType)(type_)

	c_value_p := (C.gpointer)(valueP)

	C.g_file_info_set_attribute((*C.GFileInfo)(recv.native), c_attribute, c_type, c_value_p)

	return
}

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

// Unsupported : g_file_info_set_attribute_stringv : unsupported parameter attr_value :

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
func (recv *FileInfo) SetSize(size uint64) {
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

// Unsupported signal 'changed' for FileMonitor : unsupported parameter event_type : type FileMonitorEvent :

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
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read((*C.GInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
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
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	var c_bytes_read C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_all((*C.GInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, &c_bytes_read, c_cancellable, &cThrowableError)
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

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc (GReallocFunc) for param realloc_function

// GetData is a wrapper around the C function g_memory_output_stream_get_data.
func (recv *MemoryOutputStream) GetData() uintptr {
	retC := C.g_memory_output_stream_get_data((*C.GMemoryOutputStream)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported signal 'ask-password' for MountOperation : unsupported parameter flags : type AskPasswordFlags :

// Unsupported signal 'ask-question' for MountOperation : unsupported parameter choices :

// Unsupported signal 'reply' for MountOperation : unsupported parameter result : type MountOperationResult :

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
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_write((*C.GOutputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, c_cancellable, &cThrowableError)
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
	c_buffer := &buffer[0]

	c_count := (C.gsize)(len(buffer))

	var c_bytes_written C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_all((*C.GOutputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_count, &c_bytes_written, c_cancellable, &cThrowableError)
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

// GetOpResGpointer is a wrapper around the C function g_simple_async_result_get_op_res_gpointer.
func (recv *SimpleAsyncResult) GetOpResGpointer() uintptr {
	retC := C.g_simple_async_result_get_op_res_gpointer((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// GetOpResGssize is a wrapper around the C function g_simple_async_result_get_op_res_gssize.
func (recv *SimpleAsyncResult) GetOpResGssize() int64 {
	retC := C.g_simple_async_result_get_op_res_gssize((*C.GSimpleAsyncResult)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// GetSourceTag is a wrapper around the C function g_simple_async_result_get_source_tag.
func (recv *SimpleAsyncResult) GetSourceTag() uintptr {
	retC := C.g_simple_async_result_get_source_tag((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

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

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter destroy_op_res : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy_op_res

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

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames :

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
	c_path := &path[0]

	c_path_len := (C.gint)(len(path))

	retC := C.g_unix_socket_address_new_abstract((*C.gchar)(unsafe.Pointer(c_path)), c_path_len)
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
