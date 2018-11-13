// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

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

// Integrating the launch with the launching application. This is used to
// handle for instance startup notification and launching the new application
// on the same screen as the launching window.
/*

C record/class : GAppLaunchContext
*/
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

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *AppLaunchContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to AppLaunchContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppLaunchContext.
func CastToAppLaunchContext(object *gobject.Object) *AppLaunchContext {
	return AppLaunchContextNewFromC(object.ToC())
}

// Creates a new application launch context. This is not normally used,
// instead you instantiate a subclass of this, such as #GdkAppLaunchContext.
/*

C function : g_app_launch_context_new
*/
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.g_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the display string for the @context. This is used to ensure new
// applications are started on the same display as the launching
// application, by setting the `DISPLAY` environment variable.
/*

C function : g_app_launch_context_get_display
*/
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

// Initiates startup notification for the application and returns the
// `DESKTOP_STARTUP_ID` for the launched operation, if supported.
//
// Startup notification IDs are defined in the
// [FreeDesktop.Org Startup Notifications standard](http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt").
/*

C function : g_app_launch_context_get_startup_notify_id
*/
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

// Called when an application has failed to launch, so that it can cancel
// the application startup notification started in g_app_launch_context_get_startup_notify_id().
/*

C function : g_app_launch_context_launch_failed
*/
func (recv *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	c_startup_notify_id := C.CString(startupNotifyId)
	defer C.free(unsafe.Pointer(c_startup_notify_id))

	C.g_app_launch_context_launch_failed((*C.GAppLaunchContext)(recv.native), c_startup_notify_id)

	return
}

// #GApplicationCommandLine represents a command-line invocation of
// an application.  It is created by #GApplication and emitted
// in the #GApplication::command-line signal and virtual function.
//
// The class contains the list of arguments that the program was invoked
// with.  It is also possible to query if the commandline invocation was
// local (ie: the current process is running in direct response to the
// invocation) or remote (ie: some other process forwarded the
// commandline to this process).
//
// The GApplicationCommandLine object can provide the @argc and @argv
// parameters for use with the #GOptionContext command-line parsing API,
// with the g_application_command_line_get_arguments() function. See
// [gapplication-example-cmdline3.c][gapplication-example-cmdline3]
// for an example.
//
// The exit status of the originally-invoked process may be set and
// messages can be printed to stdout or stderr of that process.  The
// lifecycle of the originally-invoked process is tied to the lifecycle
// of this object (ie: the process exits when the last reference is
// dropped).
//
// The main use for #GApplicationCommandLine (and the
// #GApplication::command-line signal) is 'Emacs server' like use cases:
// You can set the `EDITOR` environment variable to have e.g. git use
// your favourite editor to edit commit messages, and if you already
// have an instance of the editor running, the editing will happen
// in the running instance, instead of opening a new one. An important
// aspect of this use case is that the process that gets started by git
// does not return until the editing is done.
//
// Normally, the commandline is completely handled in the
// #GApplication::command-line handler. The launching instance exits
// once the signal handler in the primary instance has returned, and
// the return value of the signal handler becomes the exit status
// of the launching instance.
// |[<!-- language="C" -->
// static int
// command_line (GApplication            *application,
// GApplicationCommandLine *cmdline)
// {
// gchar **argv;
// gint argc;
// gint i;
//
// argv = g_application_command_line_get_arguments (cmdline, &argc);
//
// g_application_command_line_print (cmdline,
// "This text is written back\n"
// "to stdout of the caller\n");
//
// for (i = 0; i < argc; i++)
// g_print ("argument %d: %s\n", i, argv[i]);
//
// g_strfreev (argv);
//
// return 0;
// }
// ]|
// The complete example can be found here:
// [gapplication-example-cmdline.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline.c)
//
// In more complicated cases, the handling of the comandline can be
// split between the launcher and the primary instance.
// |[<!-- language="C" -->
// static gboolean
// test_local_cmdline (GApplication   *application,
// gchar        ***arguments,
// gint           *exit_status)
// {
// gint i, j;
// gchar **argv;
//
// argv = *arguments;
//
// i = 1;
// while (argv[i])
// {
// if (g_str_has_prefix (argv[i], "--local-"))
// {
// g_print ("handling argument %s locally\n", argv[i]);
// g_free (argv[i]);
// for (j = i; argv[j]; j++)
// argv[j] = argv[j + 1];
// }
// else
// {
// g_print ("not handling argument %s locally\n", argv[i]);
// i++;
// }
// }
//
// *exit_status = 0;
//
// return FALSE;
// }
//
// static void
// test_application_class_init (TestApplicationClass *class)
// {
// G_APPLICATION_CLASS (class)->local_command_line = test_local_cmdline;
//
// ...
// }
// ]|
// In this example of split commandline handling, options that start
// with `--local-` are handled locally, all other options are passed
// to the #GApplication::command-line handler which runs in the primary
// instance.
//
// The complete example can be found here:
// [gapplication-example-cmdline2.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline2.c)
//
// If handling the commandline requires a lot of work, it may
// be better to defer it.
// |[<!-- language="C" -->
// static gboolean
// my_cmdline_handler (gpointer data)
// {
// GApplicationCommandLine *cmdline = data;
//
// do the heavy lifting in an idle
//
// g_application_command_line_set_exit_status (cmdline, 0);
// g_object_unref (cmdline); // this releases the application
//
// return G_SOURCE_REMOVE;
// }
//
// static int
// command_line (GApplication            *application,
// GApplicationCommandLine *cmdline)
// {
// keep the application running until we are done with this commandline
// g_application_hold (application);
//
// g_object_set_data_full (G_OBJECT (cmdline),
// "application", application,
// (GDestroyNotify)g_application_release);
//
// g_object_ref (cmdline);
// g_idle_add (my_cmdline_handler, cmdline);
//
// return 0;
// }
// ]|
// In this example the commandline is not completely handled before
// the #GApplication::command-line handler returns. Instead, we keep
// a reference to the #GApplicationCommandLine object and handle it
// later (in this example, in an idle). Note that it is necessary to
// hold the application until you are done with the commandline.
//
// The complete example can be found here:
// [gapplication-example-cmdline3.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline3.c)
/*

C record/class : GApplicationCommandLine
*/
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

	return g
}

func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ApplicationCommandLine) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ApplicationCommandLine.
// Exercise care, as this is a potentially dangerous function if the Object is not a ApplicationCommandLine.
func CastToApplicationCommandLine(object *gobject.Object) *ApplicationCommandLine {
	return ApplicationCommandLineNewFromC(object.ToC())
}

// Buffered input stream implements #GFilterInputStream and provides
// for buffered reads.
//
// By default, #GBufferedInputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered input stream, use g_buffered_input_stream_new(),
// or g_buffered_input_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_input_stream_get_buffer_size(). To change the size of a
// buffered input stream's buffer, use
// g_buffered_input_stream_set_buffer_size(). Note that the buffer's size
// cannot be reduced below the size of the data within the buffer.
/*

C record/class : GBufferedInputStream
*/
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

	return g
}

func (recv *BufferedInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to BufferedInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a BufferedInputStream.
func CastToBufferedInputStream(object *gobject.Object) *BufferedInputStream {
	return BufferedInputStreamNewFromC(object.ToC())
}

// Creates a new #GInputStream from the given @base_stream, with
// a buffer set to the default size (4 kilobytes).
/*

C function : g_buffered_input_stream_new
*/
func BufferedInputStreamNew(baseStream *InputStream) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_input_stream_new(c_base_stream)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GBufferedInputStream from the given @base_stream,
// with a buffer set to @size.
/*

C function : g_buffered_input_stream_new_sized
*/
func BufferedInputStreamNewSized(baseStream *InputStream, size uint64) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_input_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Tries to read @count bytes from the stream into the buffer.
// Will block during this read.
//
// If @count is zero, returns zero and does nothing. A value of @count
// larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. near the end of a file. Zero is returned on end of file
// (or if @count is zero),  but never otherwise.
//
// If @count is -1 then the attempted read size is equal to the number of
// bytes that are required to fill the buffer.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
/*

C function : g_buffered_input_stream_fill
*/
func (recv *BufferedInputStream) Fill(count int64, cancellable *Cancellable) (int64, error) {
	c_count := (C.gssize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_fill((*C.GBufferedInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_buffered_input_stream_fill_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous read.
/*

C function : g_buffered_input_stream_fill_finish
*/
func (recv *BufferedInputStream) FillFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_fill_finish((*C.GBufferedInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the size of the available data within the stream.
/*

C function : g_buffered_input_stream_get_available
*/
func (recv *BufferedInputStream) GetAvailable() uint64 {
	retC := C.g_buffered_input_stream_get_available((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Gets the size of the input buffer.
/*

C function : g_buffered_input_stream_get_buffer_size
*/
func (recv *BufferedInputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_input_stream_get_buffer_size((*C.GBufferedInputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Peeks in the buffer, copying data of size @count into @buffer,
// offset @offset bytes.
/*

C function : g_buffered_input_stream_peek
*/
func (recv *BufferedInputStream) Peek(buffer []uint8, offset uint64) uint64 {
	c_buffer := &buffer[0]

	c_offset := (C.gsize)(offset)

	c_count := (C.gsize)(len(buffer))

	retC := C.g_buffered_input_stream_peek((*C.GBufferedInputStream)(recv.native), (unsafe.Pointer(c_buffer)), c_offset, c_count)
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_buffered_input_stream_peek_buffer : no return type

// Tries to read a single byte from the stream or the buffer. Will block
// during this read.
//
// On success, the byte read from the stream is returned. On end of stream
// -1 is returned but it's not an exceptional error and @error is not set.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
/*

C function : g_buffered_input_stream_read_byte
*/
func (recv *BufferedInputStream) ReadByte(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_buffered_input_stream_read_byte((*C.GBufferedInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the size of the internal buffer of @stream to @size, or to the
// size of the contents of the buffer. The buffer can never be resized
// smaller than its current contents.
/*

C function : g_buffered_input_stream_set_buffer_size
*/
func (recv *BufferedInputStream) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_buffered_input_stream_set_buffer_size((*C.GBufferedInputStream)(recv.native), c_size)

	return
}

// Seekable returns the Seekable interface implemented by BufferedInputStream
func (recv *BufferedInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Buffered output stream implements #GFilterOutputStream and provides
// for buffered writes.
//
// By default, #GBufferedOutputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered output stream, use g_buffered_output_stream_new(),
// or g_buffered_output_stream_new_sized() to specify the buffer's size
// at construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_output_stream_get_buffer_size(). To change the size of a
// buffered output stream's buffer, use
// g_buffered_output_stream_set_buffer_size(). Note that the buffer's
// size cannot be reduced below the size of the data within the buffer.
/*

C record/class : GBufferedOutputStream
*/
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

	return g
}

func (recv *BufferedOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to BufferedOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a BufferedOutputStream.
func CastToBufferedOutputStream(object *gobject.Object) *BufferedOutputStream {
	return BufferedOutputStreamNewFromC(object.ToC())
}

// Creates a new buffered output stream for a base stream.
/*

C function : g_buffered_output_stream_new
*/
func BufferedOutputStreamNew(baseStream *OutputStream) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_output_stream_new(c_base_stream)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new buffered output stream with a given buffer size.
/*

C function : g_buffered_output_stream_new_sized
*/
func BufferedOutputStreamNewSized(baseStream *OutputStream, size uint64) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_output_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if the buffer automatically grows as data is added.
/*

C function : g_buffered_output_stream_get_auto_grow
*/
func (recv *BufferedOutputStream) GetAutoGrow() bool {
	retC := C.g_buffered_output_stream_get_auto_grow((*C.GBufferedOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the size of the buffer in the @stream.
/*

C function : g_buffered_output_stream_get_buffer_size
*/
func (recv *BufferedOutputStream) GetBufferSize() uint64 {
	retC := C.g_buffered_output_stream_get_buffer_size((*C.GBufferedOutputStream)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Sets whether or not the @stream's buffer should automatically grow.
// If @auto_grow is true, then each write will just make the buffer
// larger, and you must manually flush the buffer to actually write out
// the data to the underlying stream.
/*

C function : g_buffered_output_stream_set_auto_grow
*/
func (recv *BufferedOutputStream) SetAutoGrow(autoGrow bool) {
	c_auto_grow :=
		boolToGboolean(autoGrow)

	C.g_buffered_output_stream_set_auto_grow((*C.GBufferedOutputStream)(recv.native), c_auto_grow)

	return
}

// Sets the size of the internal buffer to @size.
/*

C function : g_buffered_output_stream_set_buffer_size
*/
func (recv *BufferedOutputStream) SetBufferSize(size uint64) {
	c_size := (C.gsize)(size)

	C.g_buffered_output_stream_set_buffer_size((*C.GBufferedOutputStream)(recv.native), c_size)

	return
}

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// #GBytesIcon specifies an image held in memory in a common format (usually
// png) to be used as icon.
/*

C record/class : GBytesIcon
*/
type BytesIcon struct {
	native *C.GBytesIcon
}

func BytesIconNewFromC(u unsafe.Pointer) *BytesIcon {
	c := (*C.GBytesIcon)(u)
	if c == nil {
		return nil
	}

	g := &BytesIcon{native: c}

	return g
}

func (recv *BytesIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *BytesIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to BytesIcon.
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

// GCancellable is a thread-safe operation cancellation stack used
// throughout GIO to allow for cancellation of synchronous and
// asynchronous operations.
/*

C record/class : GCancellable
*/
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

	return g
}

func (recv *Cancellable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Cancellable) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Cancellable.
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
var signalCancellableCancelledLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalCancellableCancelledMap[index].callback
	callback()
}

// Creates a new #GCancellable object.
//
// Applications that want to start one or more operations
// that should be cancellable should create a #GCancellable
// and pass it to the operations.
//
// One #GCancellable can be used in multiple consecutive
// operations or in multiple concurrent operations.
/*

C function : g_cancellable_new
*/
func CancellableNew() *Cancellable {
	retC := C.g_cancellable_new()
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Will set @cancellable to cancelled, and will emit the
// #GCancellable::cancelled signal. (However, see the warning about
// race conditions in the documentation for that signal if you are
// planning to connect to it.)
//
// This function is thread-safe. In other words, you can safely call
// it from a thread other than the one running the operation that was
// passed the @cancellable.
//
// If @cancellable is %NULL, this function returns immediately for convenience.
//
// The convention within GIO is that cancelling an asynchronous
// operation causes it to complete asynchronously. That is, if you
// cancel the operation from the same thread in which it is running,
// then the operation's #GAsyncReadyCallback will not be invoked until
// the application returns to the main loop.
/*

C function : g_cancellable_cancel
*/
func (recv *Cancellable) Cancel() {
	C.g_cancellable_cancel((*C.GCancellable)(recv.native))

	return
}

// Gets the file descriptor for a cancellable job. This can be used to
// implement cancellable operations on Unix systems. The returned fd will
// turn readable when @cancellable is cancelled.
//
// You are not supposed to read from the fd yourself, just check for
// readable status. Reading to unset the readable status is done
// with g_cancellable_reset().
//
// After a successful return from this function, you should use
// g_cancellable_release_fd() to free up resources allocated for
// the returned file descriptor.
//
// See also g_cancellable_make_pollfd().
/*

C function : g_cancellable_get_fd
*/
func (recv *Cancellable) GetFd() int32 {
	retC := C.g_cancellable_get_fd((*C.GCancellable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Checks if a cancellable job has been cancelled.
/*

C function : g_cancellable_is_cancelled
*/
func (recv *Cancellable) IsCancelled() bool {
	retC := C.g_cancellable_is_cancelled((*C.GCancellable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Pops @cancellable off the cancellable stack (verifying that @cancellable
// is on the top of the stack).
/*

C function : g_cancellable_pop_current
*/
func (recv *Cancellable) PopCurrent() {
	C.g_cancellable_pop_current((*C.GCancellable)(recv.native))

	return
}

// Pushes @cancellable onto the cancellable stack. The current
// cancellable can then be received using g_cancellable_get_current().
//
// This is useful when implementing cancellable operations in
// code that does not allow you to pass down the cancellable object.
//
// This is typically called automatically by e.g. #GFile operations,
// so you rarely have to call this yourself.
/*

C function : g_cancellable_push_current
*/
func (recv *Cancellable) PushCurrent() {
	C.g_cancellable_push_current((*C.GCancellable)(recv.native))

	return
}

// Resets @cancellable to its uncancelled state.
//
// If cancellable is currently in use by any cancellable operation
// then the behavior of this function is undefined.
//
// Note that it is generally not a good idea to reuse an existing
// cancellable for more operations after it has been cancelled once,
// as this function might tempt you to do. The recommended practice
// is to drop the reference to a cancellable after cancelling it,
// and let it die with the outstanding async operations. You should
// create a fresh cancellable for further async operations.
/*

C function : g_cancellable_reset
*/
func (recv *Cancellable) Reset() {
	C.g_cancellable_reset((*C.GCancellable)(recv.native))

	return
}

// If the @cancellable is cancelled, sets the error to notify
// that the operation was cancelled.
/*

C function : g_cancellable_set_error_if_cancelled
*/
func (recv *Cancellable) SetErrorIfCancelled() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_cancellable_set_error_if_cancelled((*C.GCancellable)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GCharsetConverter is an implementation of #GConverter based on
// GIConv.
/*

C record/class : GCharsetConverter
*/
type CharsetConverter struct {
	native *C.GCharsetConverter
}

func CharsetConverterNewFromC(u unsafe.Pointer) *CharsetConverter {
	c := (*C.GCharsetConverter)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverter{native: c}

	return g
}

func (recv *CharsetConverter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *CharsetConverter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to CharsetConverter.
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

// Converter input stream implements #GInputStream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, #GConverterInputStream implements
// #GPollableInputStream.
/*

C record/class : GConverterInputStream
*/
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

	return g
}

func (recv *ConverterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to ConverterInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a ConverterInputStream.
func CastToConverterInputStream(object *gobject.Object) *ConverterInputStream {
	return ConverterInputStreamNewFromC(object.ToC())
}

// Creates a new converter input stream for the @base_stream.
/*

C function : g_converter_input_stream_new
*/
func ConverterInputStreamNew(baseStream *InputStream, converter *Converter) *ConverterInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_input_stream_new(c_base_stream, c_converter)
	retGo := ConverterInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PollableInputStream returns the PollableInputStream interface implemented by ConverterInputStream
func (recv *ConverterInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Converter output stream implements #GOutputStream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, #GConverterOutputStream implements
// #GPollableOutputStream.
/*

C record/class : GConverterOutputStream
*/
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

	return g
}

func (recv *ConverterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to ConverterOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a ConverterOutputStream.
func CastToConverterOutputStream(object *gobject.Object) *ConverterOutputStream {
	return ConverterOutputStreamNewFromC(object.ToC())
}

// Creates a new converter output stream for the @base_stream.
/*

C function : g_converter_output_stream_new
*/
func ConverterOutputStreamNew(baseStream *OutputStream, converter *Converter) *ConverterOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_output_stream_new(c_base_stream, c_converter)
	retGo := ConverterOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// #GDBusActionGroup is an implementation of the #GActionGroup
// interface that can be used as a proxy for an action group
// that is exported over D-Bus with g_dbus_connection_export_action_group().
/*

C record/class : GDBusActionGroup
*/
type DBusActionGroup struct {
	native *C.GDBusActionGroup
}

func DBusActionGroupNewFromC(u unsafe.Pointer) *DBusActionGroup {
	c := (*C.GDBusActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &DBusActionGroup{native: c}

	return g
}

func (recv *DBusActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DBusActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DBusActionGroup.
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

// #GDBusMenuModel is an implementation of #GMenuModel that can be used
// as a proxy for a menu model that is exported over D-Bus with
// g_dbus_connection_export_menu_model().
/*

C record/class : GDBusMenuModel
*/
type DBusMenuModel struct {
	native *C.GDBusMenuModel
}

func DBusMenuModelNewFromC(u unsafe.Pointer) *DBusMenuModel {
	c := (*C.GDBusMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &DBusMenuModel{native: c}

	return g
}

func (recv *DBusMenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModel upcasts to *MenuModel
func (recv *DBusMenuModel) MenuModel() *MenuModel {
	return MenuModelNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *DBusMenuModel) Object() *gobject.Object {
	return recv.MenuModel().Object()
}

// CastToWidget down casts any arbitary Object to DBusMenuModel.
// Exercise care, as this is a potentially dangerous function if the Object is not a DBusMenuModel.
func CastToDBusMenuModel(object *gobject.Object) *DBusMenuModel {
	return DBusMenuModelNewFromC(object.ToC())
}

// Data input stream implements #GInputStream and includes functions for
// reading structured data directly from a binary input stream.
/*

C record/class : GDataInputStream
*/
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

	return g
}

func (recv *DataInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to DataInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a DataInputStream.
func CastToDataInputStream(object *gobject.Object) *DataInputStream {
	return DataInputStreamNewFromC(object.ToC())
}

// Creates a new data input stream for the @base_stream.
/*

C function : g_data_input_stream_new
*/
func DataInputStreamNew(baseStream *InputStream) *DataInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_data_input_stream_new(c_base_stream)
	retGo := DataInputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the byte order for the data input stream.
/*

C function : g_data_input_stream_get_byte_order
*/
func (recv *DataInputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_input_stream_get_byte_order((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// Gets the current newline type for the @stream.
/*

C function : g_data_input_stream_get_newline_type
*/
func (recv *DataInputStream) GetNewlineType() DataStreamNewlineType {
	retC := C.g_data_input_stream_get_newline_type((*C.GDataInputStream)(recv.native))
	retGo := (DataStreamNewlineType)(retC)

	return retGo
}

// Reads an unsigned 8-bit/1-byte value from @stream.
/*

C function : g_data_input_stream_read_byte
*/
func (recv *DataInputStream) ReadByte(cancellable *Cancellable) (uint8, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_byte((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint8)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads a 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
/*

C function : g_data_input_stream_read_int16
*/
func (recv *DataInputStream) ReadInt16(cancellable *Cancellable) (int16, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int16((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int16)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads a signed 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_data_input_stream_read_int32
*/
func (recv *DataInputStream) ReadInt32(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int32((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads a 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_data_input_stream_read_int64
*/
func (recv *DataInputStream) ReadInt64(cancellable *Cancellable) (int64, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_int64((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_data_input_stream_read_line : no return type

// Reads an unsigned 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
/*

C function : g_data_input_stream_read_uint16
*/
func (recv *DataInputStream) ReadUint16(cancellable *Cancellable) (uint16, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint16((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint16)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads an unsigned 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_data_input_stream_read_uint32
*/
func (recv *DataInputStream) ReadUint32(cancellable *Cancellable) (uint32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint32((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads an unsigned 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_data_input_stream_read_uint64
*/
func (recv *DataInputStream) ReadUint64(cancellable *Cancellable) (uint64, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_input_stream_read_uint64((*C.GDataInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := (uint64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// Note that, in contrast to g_data_input_stream_read_until_async(),
// this function consumes the stop character that it finds.
//
// Don't use this function in new code.  Its functionality is
// inconsistent with g_data_input_stream_read_until_async().  Both
// functions will be marked as deprecated in a future release.  Use
// g_data_input_stream_read_upto() instead, but note that that function
// does not consume the stop character.
/*

C function : g_data_input_stream_read_until
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// This function sets the byte order for the given @stream. All subsequent
// reads from the @stream will be read in the given @order.
/*

C function : g_data_input_stream_set_byte_order
*/
func (recv *DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_input_stream_set_byte_order((*C.GDataInputStream)(recv.native), c_order)

	return
}

// Sets the newline type for the @stream.
//
// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a read
// chunk ends in "CR" we must read an additional byte to know if this is "CR" or
// "CR LF", and this might block if there is no more data available.
/*

C function : g_data_input_stream_set_newline_type
*/
func (recv *DataInputStream) SetNewlineType(type_ DataStreamNewlineType) {
	c_type := (C.GDataStreamNewlineType)(type_)

	C.g_data_input_stream_set_newline_type((*C.GDataInputStream)(recv.native), c_type)

	return
}

// Seekable returns the Seekable interface implemented by DataInputStream
func (recv *DataInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Data output stream implements #GOutputStream and includes functions for
// writing data directly to an output stream.
/*

C record/class : GDataOutputStream
*/
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

	return g
}

func (recv *DataOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to DataOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a DataOutputStream.
func CastToDataOutputStream(object *gobject.Object) *DataOutputStream {
	return DataOutputStreamNewFromC(object.ToC())
}

// Creates a new data output stream for @base_stream.
/*

C function : g_data_output_stream_new
*/
func DataOutputStreamNew(baseStream *OutputStream) *DataOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_data_output_stream_new(c_base_stream)
	retGo := DataOutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the byte order for the stream.
/*

C function : g_data_output_stream_get_byte_order
*/
func (recv *DataOutputStream) GetByteOrder() DataStreamByteOrder {
	retC := C.g_data_output_stream_get_byte_order((*C.GDataOutputStream)(recv.native))
	retGo := (DataStreamByteOrder)(retC)

	return retGo
}

// Puts a byte into the output stream.
/*

C function : g_data_output_stream_put_byte
*/
func (recv *DataOutputStream) PutByte(data uint8, cancellable *Cancellable) (bool, error) {
	c_data := (C.guchar)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_byte((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts a signed 16-bit integer into the output stream.
/*

C function : g_data_output_stream_put_int16
*/
func (recv *DataOutputStream) PutInt16(data int16, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint16)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int16((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts a signed 32-bit integer into the output stream.
/*

C function : g_data_output_stream_put_int32
*/
func (recv *DataOutputStream) PutInt32(data int32, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint32)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int32((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts a signed 64-bit integer into the stream.
/*

C function : g_data_output_stream_put_int64
*/
func (recv *DataOutputStream) PutInt64(data int64, cancellable *Cancellable) (bool, error) {
	c_data := (C.gint64)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_int64((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts a string into the output stream.
/*

C function : g_data_output_stream_put_string
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts an unsigned 16-bit integer into the output stream.
/*

C function : g_data_output_stream_put_uint16
*/
func (recv *DataOutputStream) PutUint16(data uint16, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint16)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint16((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts an unsigned 32-bit integer into the stream.
/*

C function : g_data_output_stream_put_uint32
*/
func (recv *DataOutputStream) PutUint32(data uint32, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint32)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint32((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Puts an unsigned 64-bit integer into the stream.
/*

C function : g_data_output_stream_put_uint64
*/
func (recv *DataOutputStream) PutUint64(data uint64, cancellable *Cancellable) (bool, error) {
	c_data := (C.guint64)(data)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_data_output_stream_put_uint64((*C.GDataOutputStream)(recv.native), c_data, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the byte order of the data output stream to @order.
/*

C function : g_data_output_stream_set_byte_order
*/
func (recv *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	c_order := (C.GDataStreamByteOrder)(order)

	C.g_data_output_stream_set_byte_order((*C.GDataOutputStream)(recv.native), c_order)

	return
}

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// #GDesktopAppInfo is an implementation of #GAppInfo based on
// desktop files.
//
// Note that `<gio/gdesktopappinfo.h>` belongs to the UNIX-specific
// GIO interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
/*

C record/class : GDesktopAppInfo
*/
type DesktopAppInfo struct {
	native *C.GDesktopAppInfo
}

func DesktopAppInfoNewFromC(u unsafe.Pointer) *DesktopAppInfo {
	c := (*C.GDesktopAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfo{native: c}

	return g
}

func (recv *DesktopAppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *DesktopAppInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to DesktopAppInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a DesktopAppInfo.
func CastToDesktopAppInfo(object *gobject.Object) *DesktopAppInfo {
	return DesktopAppInfoNewFromC(object.ToC())
}

// Creates a new #GDesktopAppInfo based on a desktop file id.
//
// A desktop file id is the basename of the desktop file, including the
// .desktop extension. GIO is looking for a desktop file with this name
// in the `applications` subdirectories of the XDG
// data directories (i.e. the directories specified in the `XDG_DATA_HOME`
// and `XDG_DATA_DIRS` environment variables). GIO also supports the
// prefix-to-subdirectory mapping that is described in the
// [Menu Spec](http://standards.freedesktop.org/menu-spec/latest/)
// (i.e. a desktop id of kde-foo.desktop will match
// `/usr/share/applications/kde/foo.desktop`).
/*

C function : g_desktop_app_info_new
*/
func DesktopAppInfoNew(desktopId string) *DesktopAppInfo {
	c_desktop_id := C.CString(desktopId)
	defer C.free(unsafe.Pointer(c_desktop_id))

	retC := C.g_desktop_app_info_new(c_desktop_id)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GDesktopAppInfo.
/*

C function : g_desktop_app_info_new_from_filename
*/
func DesktopAppInfoNewFromFilename(filename string) *DesktopAppInfo {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_desktop_app_info_new_from_filename(c_filename)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the categories from the desktop file.
/*

C function : g_desktop_app_info_get_categories
*/
func (recv *DesktopAppInfo) GetCategories() string {
	retC := C.g_desktop_app_info_get_categories((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the generic name from the destkop file.
/*

C function : g_desktop_app_info_get_generic_name
*/
func (recv *DesktopAppInfo) GetGenericName() string {
	retC := C.g_desktop_app_info_get_generic_name((*C.GDesktopAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// A desktop file is hidden if the Hidden key in it is
// set to True.
/*

C function : g_desktop_app_info_get_is_hidden
*/
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

// #GEmblem is an implementation of #GIcon that supports
// having an emblem, which is an icon with additional properties.
// It can than be added to a #GEmblemedIcon.
//
// Currently, only metainformation about the emblem's origin is
// supported. More may be added in the future.
/*

C record/class : GEmblem
*/
type Emblem struct {
	native *C.GEmblem
}

func EmblemNewFromC(u unsafe.Pointer) *Emblem {
	c := (*C.GEmblem)(u)
	if c == nil {
		return nil
	}

	g := &Emblem{native: c}

	return g
}

func (recv *Emblem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Emblem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Emblem.
// Exercise care, as this is a potentially dangerous function if the Object is not a Emblem.
func CastToEmblem(object *gobject.Object) *Emblem {
	return EmblemNewFromC(object.ToC())
}

// Icon returns the Icon interface implemented by Emblem
func (recv *Emblem) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// #GEmblemedIcon is an implementation of #GIcon that supports
// adding an emblem to an icon. Adding multiple emblems to an
// icon is ensured via g_emblemed_icon_add_emblem().
//
// Note that #GEmblemedIcon allows no control over the position
// of the emblems. See also #GEmblem for more information.
/*

C record/class : GEmblemedIcon
*/
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

	return g
}

func (recv *EmblemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *EmblemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to EmblemedIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a EmblemedIcon.
func CastToEmblemedIcon(object *gobject.Object) *EmblemedIcon {
	return EmblemedIconNewFromC(object.ToC())
}

// Icon returns the Icon interface implemented by EmblemedIcon
func (recv *EmblemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// #GFileEnumerator allows you to operate on a set of #GFiles,
// returning a #GFileInfo structure for each file enumerated (e.g.
// g_file_enumerate_children() will return a #GFileEnumerator for each
// of the children within a directory).
//
// To get the next file's information from a #GFileEnumerator, use
// g_file_enumerator_next_file() or its asynchronous version,
// g_file_enumerator_next_files_async(). Note that the asynchronous
// version will return a list of #GFileInfos, whereas the
// synchronous will only return the next file in the enumerator.
//
// The ordering of returned files is unspecified for non-Unix
// platforms; for more information, see g_dir_read_name().  On Unix,
// when operating on local files, returned files will be sorted by
// inode number.  Effectively you can assume that the ordering of
// returned files will be stable between successive calls (and
// applications) assuming the directory is unchanged.
//
// If your application needs a specific ordering, such as by name or
// modification time, you will have to implement that in your
// application code.
//
// To close a #GFileEnumerator, use g_file_enumerator_close(), or
// its asynchronous version, g_file_enumerator_close_async(). Once
// a #GFileEnumerator is closed, no further actions may be performed
// on it, and it should be freed with g_object_unref().
/*

C record/class : GFileEnumerator
*/
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

	return g
}

func (recv *FileEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FileEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FileEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileEnumerator.
func CastToFileEnumerator(object *gobject.Object) *FileEnumerator {
	return FileEnumeratorNewFromC(object.ToC())
}

// Releases all resources used by this enumerator, making the
// enumerator return %G_IO_ERROR_CLOSED on all calls.
//
// This will be automatically called when the last reference
// is dropped, but you might want to call this function to make
// sure resources are released as early as possible.
/*

C function : g_file_enumerator_close
*/
func (recv *FileEnumerator) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_close((*C.GFileEnumerator)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_enumerator_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes closing a file enumerator, started from g_file_enumerator_close_async().
//
// If the file enumerator was already closed when g_file_enumerator_close_async()
// was called, then this function will report %G_IO_ERROR_CLOSED in @error, and
// return %FALSE. If the file enumerator had pending operation when the close
// operation was started, then this function will report %G_IO_ERROR_PENDING, and
// return %FALSE.  If @cancellable was not %NULL, then the operation may have been
// cancelled by triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %FALSE will be
// returned.
/*

C function : g_file_enumerator_close_finish
*/
func (recv *FileEnumerator) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_close_finish((*C.GFileEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if the file enumerator has pending operations.
/*

C function : g_file_enumerator_has_pending
*/
func (recv *FileEnumerator) HasPending() bool {
	retC := C.g_file_enumerator_has_pending((*C.GFileEnumerator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the file enumerator has been closed.
/*

C function : g_file_enumerator_is_closed
*/
func (recv *FileEnumerator) IsClosed() bool {
	retC := C.g_file_enumerator_is_closed((*C.GFileEnumerator)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns information for the next file in the enumerated object.
// Will block until the information is available. The #GFileInfo
// returned from this function will contain attributes that match the
// attribute string that was passed when the #GFileEnumerator was created.
//
// See the documentation of #GFileEnumerator for information about the
// order of returned files.
//
// On error, returns %NULL and sets @error to the error. If the
// enumerator is at the end, %NULL will be returned and @error will
// be unset.
/*

C function : g_file_enumerator_next_file
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_enumerator_next_files_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes the asynchronous operation started with g_file_enumerator_next_files_async().
/*

C function : g_file_enumerator_next_files_finish
*/
func (recv *FileEnumerator) NextFilesFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_enumerator_next_files_finish((*C.GFileEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the file enumerator as having pending operations.
/*

C function : g_file_enumerator_set_pending
*/
func (recv *FileEnumerator) SetPending(pending bool) {
	c_pending :=
		boolToGboolean(pending)

	C.g_file_enumerator_set_pending((*C.GFileEnumerator)(recv.native), c_pending)

	return
}

// GFileIOStream provides io streams that both read and write to the same
// file handle.
//
// GFileIOStream implements #GSeekable, which allows the io
// stream to jump to arbitrary positions in the file and to truncate
// the file, provided the filesystem of the file supports these
// operations.
//
// To find the position of a file io stream, use
// g_seekable_tell().
//
// To find out if a file io stream supports seeking, use g_seekable_can_seek().
// To position a file io stream, use g_seekable_seek().
// To find out if a file io stream supports truncating, use
// g_seekable_can_truncate(). To truncate a file io
// stream, use g_seekable_truncate().
//
// The default implementation of all the #GFileIOStream operations
// and the implementation of #GSeekable just call into the same operations
// on the output stream.
/*

C record/class : GFileIOStream
*/
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

	return g
}

func (recv *FileIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStream upcasts to *IOStream
func (recv *FileIOStream) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileIOStream) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitary Object to FileIOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileIOStream.
func CastToFileIOStream(object *gobject.Object) *FileIOStream {
	return FileIOStreamNewFromC(object.ToC())
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// #GFileIcon specifies an icon by pointing to an image file
// to be used as icon.
/*

C record/class : GFileIcon
*/
type FileIcon struct {
	native *C.GFileIcon
}

func FileIconNewFromC(u unsafe.Pointer) *FileIcon {
	c := (*C.GFileIcon)(u)
	if c == nil {
		return nil
	}

	g := &FileIcon{native: c}

	return g
}

func (recv *FileIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FileIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FileIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileIcon.
func CastToFileIcon(object *gobject.Object) *FileIcon {
	return FileIconNewFromC(object.ToC())
}

// Creates a new icon for a file.
/*

C function : g_file_icon_new
*/
func FileIconNew(file *File) *FileIcon {
	c_file := (*C.GFile)(file.ToC())

	retC := C.g_file_icon_new(c_file)
	retGo := FileIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GFile associated with the given @icon.
/*

C function : g_file_icon_get_file
*/
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

// Functionality for manipulating basic metadata for files. #GFileInfo
// implements methods for getting information that all files should
// contain, and allows for manipulation of extended attributes.
//
// See [GFileAttribute][gio-GFileAttribute] for more information on how
// GIO handles file attributes.
//
// To obtain a #GFileInfo for a #GFile, use g_file_query_info() (or its
// async variant). To obtain a #GFileInfo for a file input or output
// stream, use g_file_input_stream_query_info() or
// g_file_output_stream_query_info() (or their async variants).
//
// To change the actual attributes of a file, you should then set the
// attribute in the #GFileInfo and call g_file_set_attributes_from_info()
// or g_file_set_attributes_async() on a GFile.
//
// However, not all attributes can be changed in the file. For instance,
// the actual size of a file cannot be changed via g_file_info_set_size().
// You may call g_file_query_settable_attributes() and
// g_file_query_writable_namespaces() to discover the settable attributes
// of a particular file at runtime.
//
// #GFileAttributeMatcher allows for searching through a #GFileInfo for
// attributes.
/*

C record/class : GFileInfo
*/
type FileInfo struct {
	native *C.GFileInfo
}

func FileInfoNewFromC(u unsafe.Pointer) *FileInfo {
	c := (*C.GFileInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileInfo{native: c}

	return g
}

func (recv *FileInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FileInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FileInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileInfo.
func CastToFileInfo(object *gobject.Object) *FileInfo {
	return FileInfoNewFromC(object.ToC())
}

// Creates a new file info structure.
/*

C function : g_file_info_new
*/
func FileInfoNew() *FileInfo {
	retC := C.g_file_info_new()
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clears the status information from @info.
/*

C function : g_file_info_clear_status
*/
func (recv *FileInfo) ClearStatus() {
	C.g_file_info_clear_status((*C.GFileInfo)(recv.native))

	return
}

// First clears all of the [GFileAttribute][gio-GFileAttribute] of @dest_info,
// and then copies all of the file attributes from @src_info to @dest_info.
/*

C function : g_file_info_copy_into
*/
func (recv *FileInfo) CopyInto(destInfo *FileInfo) {
	c_dest_info := (*C.GFileInfo)(C.NULL)
	if destInfo != nil {
		c_dest_info = (*C.GFileInfo)(destInfo.ToC())
	}

	C.g_file_info_copy_into((*C.GFileInfo)(recv.native), c_dest_info)

	return
}

// Duplicates a file info structure.
/*

C function : g_file_info_dup
*/
func (recv *FileInfo) Dup() *FileInfo {
	retC := C.g_file_info_dup((*C.GFileInfo)(recv.native))
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the value of a attribute, formated as a string.
// This escapes things as needed to make the string valid
// utf8.
/*

C function : g_file_info_get_attribute_as_string
*/
func (recv *FileInfo) GetAttributeAsString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_as_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the value of a boolean attribute. If the attribute does not
// contain a boolean value, %FALSE will be returned.
/*

C function : g_file_info_get_attribute_boolean
*/
func (recv *FileInfo) GetAttributeBoolean(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_boolean((*C.GFileInfo)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the value of a byte string attribute. If the attribute does
// not contain a byte string, %NULL will be returned.
/*

C function : g_file_info_get_attribute_byte_string
*/
func (recv *FileInfo) GetAttributeByteString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_byte_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_file_info_get_attribute_data : unsupported parameter type : GFileAttributeType* with indirection level of 1

// Gets a signed 32-bit integer contained within the attribute. If the
// attribute does not contain a signed 32-bit integer, or is invalid,
// 0 will be returned.
/*

C function : g_file_info_get_attribute_int32
*/
func (recv *FileInfo) GetAttributeInt32(attribute string) int32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int32)(retC)

	return retGo
}

// Gets a signed 64-bit integer contained within the attribute. If the
// attribute does not contain an signed 64-bit integer, or is invalid,
// 0 will be returned.
/*

C function : g_file_info_get_attribute_int64
*/
func (recv *FileInfo) GetAttributeInt64(attribute string) int64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_int64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (int64)(retC)

	return retGo
}

// Gets the value of a #GObject attribute. If the attribute does
// not contain a #GObject, %NULL will be returned.
/*

C function : g_file_info_get_attribute_object
*/
func (recv *FileInfo) GetAttributeObject(attribute string) *gobject.Object {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_object((*C.GFileInfo)(recv.native), c_attribute)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the attribute status for an attribute key.
/*

C function : g_file_info_get_attribute_status
*/
func (recv *FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_status((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeStatus)(retC)

	return retGo
}

// Gets the value of a string attribute. If the attribute does
// not contain a string, %NULL will be returned.
/*

C function : g_file_info_get_attribute_string
*/
func (recv *FileInfo) GetAttributeString(attribute string) string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_string((*C.GFileInfo)(recv.native), c_attribute)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the attribute type for an attribute key.
/*

C function : g_file_info_get_attribute_type
*/
func (recv *FileInfo) GetAttributeType(attribute string) FileAttributeType {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_type((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (FileAttributeType)(retC)

	return retGo
}

// Gets an unsigned 32-bit integer contained within the attribute. If the
// attribute does not contain an unsigned 32-bit integer, or is invalid,
// 0 will be returned.
/*

C function : g_file_info_get_attribute_uint32
*/
func (recv *FileInfo) GetAttributeUint32(attribute string) uint32 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint32((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint32)(retC)

	return retGo
}

// Gets a unsigned 64-bit integer contained within the attribute. If the
// attribute does not contain an unsigned 64-bit integer, or is invalid,
// 0 will be returned.
/*

C function : g_file_info_get_attribute_uint64
*/
func (recv *FileInfo) GetAttributeUint64(attribute string) uint64 {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_uint64((*C.GFileInfo)(recv.native), c_attribute)
	retGo := (uint64)(retC)

	return retGo
}

// Gets the file's content type.
/*

C function : g_file_info_get_content_type
*/
func (recv *FileInfo) GetContentType() string {
	retC := C.g_file_info_get_content_type((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets a display name for a file.
/*

C function : g_file_info_get_display_name
*/
func (recv *FileInfo) GetDisplayName() string {
	retC := C.g_file_info_get_display_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the edit name for a file.
/*

C function : g_file_info_get_edit_name
*/
func (recv *FileInfo) GetEditName() string {
	retC := C.g_file_info_get_edit_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the [entity tag][gfile-etag] for a given
// #GFileInfo. See %G_FILE_ATTRIBUTE_ETAG_VALUE.
/*

C function : g_file_info_get_etag
*/
func (recv *FileInfo) GetEtag() string {
	retC := C.g_file_info_get_etag((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets a file's type (whether it is a regular file, symlink, etc).
// This is different from the file's content type, see g_file_info_get_content_type().
/*

C function : g_file_info_get_file_type
*/
func (recv *FileInfo) GetFileType() FileType {
	retC := C.g_file_info_get_file_type((*C.GFileInfo)(recv.native))
	retGo := (FileType)(retC)

	return retGo
}

// Gets the icon for a file.
/*

C function : g_file_info_get_icon
*/
func (recv *FileInfo) GetIcon() *Icon {
	retC := C.g_file_info_get_icon((*C.GFileInfo)(recv.native))
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if a file is a backup file.
/*

C function : g_file_info_get_is_backup
*/
func (recv *FileInfo) GetIsBackup() bool {
	retC := C.g_file_info_get_is_backup((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a file is hidden.
/*

C function : g_file_info_get_is_hidden
*/
func (recv *FileInfo) GetIsHidden() bool {
	retC := C.g_file_info_get_is_hidden((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a file is a symlink.
/*

C function : g_file_info_get_is_symlink
*/
func (recv *FileInfo) GetIsSymlink() bool {
	retC := C.g_file_info_get_is_symlink((*C.GFileInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the modification time of the current @info and sets it
// in @result.
/*

C function : g_file_info_get_modification_time
*/
func (recv *FileInfo) GetModificationTime() *glib.TimeVal {
	var c_result C.GTimeVal

	C.g_file_info_get_modification_time((*C.GFileInfo)(recv.native), &c_result)

	result := glib.TimeValNewFromC(unsafe.Pointer(&c_result))

	return result
}

// Gets the name for a file.
/*

C function : g_file_info_get_name
*/
func (recv *FileInfo) GetName() string {
	retC := C.g_file_info_get_name((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the file's size.
/*

C function : g_file_info_get_size
*/
func (recv *FileInfo) GetSize() uint64 {
	retC := C.g_file_info_get_size((*C.GFileInfo)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Gets the value of the sort_order attribute from the #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
/*

C function : g_file_info_get_sort_order
*/
func (recv *FileInfo) GetSortOrder() int32 {
	retC := C.g_file_info_get_sort_order((*C.GFileInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the symlink target for a given #GFileInfo.
/*

C function : g_file_info_get_symlink_target
*/
func (recv *FileInfo) GetSymlinkTarget() string {
	retC := C.g_file_info_get_symlink_target((*C.GFileInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Checks if a file info structure has an attribute named @attribute.
/*

C function : g_file_info_has_attribute
*/
func (recv *FileInfo) HasAttribute(attribute string) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_has_attribute((*C.GFileInfo)(recv.native), c_attribute)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_info_list_attributes : no return type

// Removes all cases of @attribute from @info if it exists.
/*

C function : g_file_info_remove_attribute
*/
func (recv *FileInfo) RemoveAttribute(attribute string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	C.g_file_info_remove_attribute((*C.GFileInfo)(recv.native), c_attribute)

	return
}

// Sets the @attribute to contain the given value, if possible. To unset the
// attribute, use %G_ATTRIBUTE_TYPE_INVALID for @type.
/*

C function : g_file_info_set_attribute
*/
func (recv *FileInfo) SetAttribute(attribute string, type_ FileAttributeType, valueP uintptr) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_type := (C.GFileAttributeType)(type_)

	c_value_p := (C.gpointer)(valueP)

	C.g_file_info_set_attribute((*C.GFileInfo)(recv.native), c_attribute, c_type, c_value_p)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_boolean
*/
func (recv *FileInfo) SetAttributeBoolean(attribute string, attrValue bool) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value :=
		boolToGboolean(attrValue)

	C.g_file_info_set_attribute_boolean((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_byte_string
*/
func (recv *FileInfo) SetAttributeByteString(attribute string, attrValue string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := C.CString(attrValue)
	defer C.free(unsafe.Pointer(c_attr_value))

	C.g_file_info_set_attribute_byte_string((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_int32
*/
func (recv *FileInfo) SetAttributeInt32(attribute string, attrValue int32) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.gint32)(attrValue)

	C.g_file_info_set_attribute_int32((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_int64
*/
func (recv *FileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.gint64)(attrValue)

	C.g_file_info_set_attribute_int64((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets @mask on @info to match specific attribute types.
/*

C function : g_file_info_set_attribute_mask
*/
func (recv *FileInfo) SetAttributeMask(mask *FileAttributeMatcher) {
	c_mask := (*C.GFileAttributeMatcher)(C.NULL)
	if mask != nil {
		c_mask = (*C.GFileAttributeMatcher)(mask.ToC())
	}

	C.g_file_info_set_attribute_mask((*C.GFileInfo)(recv.native), c_mask)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_object
*/
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

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_string
*/
func (recv *FileInfo) SetAttributeString(attribute string, attrValue string) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := C.CString(attrValue)
	defer C.free(unsafe.Pointer(c_attr_value))

	C.g_file_info_set_attribute_string((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Unsupported : g_file_info_set_attribute_stringv : unsupported parameter attr_value :

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_uint32
*/
func (recv *FileInfo) SetAttributeUint32(attribute string, attrValue uint32) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.guint32)(attrValue)

	C.g_file_info_set_attribute_uint32((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets the @attribute to contain the given @attr_value,
// if possible.
/*

C function : g_file_info_set_attribute_uint64
*/
func (recv *FileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_attr_value := (C.guint64)(attrValue)

	C.g_file_info_set_attribute_uint64((*C.GFileInfo)(recv.native), c_attribute, c_attr_value)

	return
}

// Sets the content type attribute for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
/*

C function : g_file_info_set_content_type
*/
func (recv *FileInfo) SetContentType(contentType string) {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	C.g_file_info_set_content_type((*C.GFileInfo)(recv.native), c_content_type)

	return
}

// Sets the display name for the current #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
/*

C function : g_file_info_set_display_name
*/
func (recv *FileInfo) SetDisplayName(displayName string) {
	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	C.g_file_info_set_display_name((*C.GFileInfo)(recv.native), c_display_name)

	return
}

// Sets the edit name for the current file.
// See %G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
/*

C function : g_file_info_set_edit_name
*/
func (recv *FileInfo) SetEditName(editName string) {
	c_edit_name := C.CString(editName)
	defer C.free(unsafe.Pointer(c_edit_name))

	C.g_file_info_set_edit_name((*C.GFileInfo)(recv.native), c_edit_name)

	return
}

// Sets the file type in a #GFileInfo to @type.
// See %G_FILE_ATTRIBUTE_STANDARD_TYPE.
/*

C function : g_file_info_set_file_type
*/
func (recv *FileInfo) SetFileType(type_ FileType) {
	c_type := (C.GFileType)(type_)

	C.g_file_info_set_file_type((*C.GFileInfo)(recv.native), c_type)

	return
}

// Sets the icon for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_ICON.
/*

C function : g_file_info_set_icon
*/
func (recv *FileInfo) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_file_info_set_icon((*C.GFileInfo)(recv.native), c_icon)

	return
}

// Sets the "is_hidden" attribute in a #GFileInfo according to @is_hidden.
// See %G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
/*

C function : g_file_info_set_is_hidden
*/
func (recv *FileInfo) SetIsHidden(isHidden bool) {
	c_is_hidden :=
		boolToGboolean(isHidden)

	C.g_file_info_set_is_hidden((*C.GFileInfo)(recv.native), c_is_hidden)

	return
}

// Sets the "is_symlink" attribute in a #GFileInfo according to @is_symlink.
// See %G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
/*

C function : g_file_info_set_is_symlink
*/
func (recv *FileInfo) SetIsSymlink(isSymlink bool) {
	c_is_symlink :=
		boolToGboolean(isSymlink)

	C.g_file_info_set_is_symlink((*C.GFileInfo)(recv.native), c_is_symlink)

	return
}

// Sets the %G_FILE_ATTRIBUTE_TIME_MODIFIED attribute in the file
// info to the given time value.
/*

C function : g_file_info_set_modification_time
*/
func (recv *FileInfo) SetModificationTime(mtime *glib.TimeVal) {
	c_mtime := (*C.GTimeVal)(C.NULL)
	if mtime != nil {
		c_mtime = (*C.GTimeVal)(mtime.ToC())
	}

	C.g_file_info_set_modification_time((*C.GFileInfo)(recv.native), c_mtime)

	return
}

// Sets the name attribute for the current #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_NAME.
/*

C function : g_file_info_set_name
*/
func (recv *FileInfo) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.g_file_info_set_name((*C.GFileInfo)(recv.native), c_name)

	return
}

// Sets the %G_FILE_ATTRIBUTE_STANDARD_SIZE attribute in the file info
// to the given size.
/*

C function : g_file_info_set_size
*/
func (recv *FileInfo) SetSize(size uint64) {
	c_size := (C.goffset)(size)

	C.g_file_info_set_size((*C.GFileInfo)(recv.native), c_size)

	return
}

// Sets the sort order attribute in the file info structure. See
// %G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
/*

C function : g_file_info_set_sort_order
*/
func (recv *FileInfo) SetSortOrder(sortOrder int32) {
	c_sort_order := (C.gint32)(sortOrder)

	C.g_file_info_set_sort_order((*C.GFileInfo)(recv.native), c_sort_order)

	return
}

// Sets the %G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET attribute in the file info
// to the given symlink target.
/*

C function : g_file_info_set_symlink_target
*/
func (recv *FileInfo) SetSymlinkTarget(symlinkTarget string) {
	c_symlink_target := C.CString(symlinkTarget)
	defer C.free(unsafe.Pointer(c_symlink_target))

	C.g_file_info_set_symlink_target((*C.GFileInfo)(recv.native), c_symlink_target)

	return
}

// Unsets a mask set by g_file_info_set_attribute_mask(), if one
// is set.
/*

C function : g_file_info_unset_attribute_mask
*/
func (recv *FileInfo) UnsetAttributeMask() {
	C.g_file_info_unset_attribute_mask((*C.GFileInfo)(recv.native))

	return
}

// GFileInputStream provides input streams that take their
// content from a file.
//
// GFileInputStream implements #GSeekable, which allows the input
// stream to jump to arbitrary positions in the file, provided the
// filesystem of the file allows it. To find the position of a file
// input stream, use g_seekable_tell(). To find out if a file input
// stream supports seeking, use g_seekable_can_seek().
// To position a file input stream, use g_seekable_seek().
/*

C record/class : GFileInputStream
*/
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

	return g
}

func (recv *FileInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStream upcasts to *InputStream
func (recv *FileInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitary Object to FileInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileInputStream.
func CastToFileInputStream(object *gobject.Object) *FileInputStream {
	return FileInputStreamNewFromC(object.ToC())
}

// Queries a file input stream the given @attributes. This function blocks
// while querying the stream. For the asynchronous (non-blocking) version
// of this function, see g_file_input_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and
// any other operations on the stream will fail with %G_IO_ERROR_PENDING.
/*

C function : g_file_input_stream_query_info
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_input_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous info query operation.
/*

C function : g_file_input_stream_query_info_finish
*/
func (recv *FileInputStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_input_stream_query_info_finish((*C.GFileInputStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Seekable returns the Seekable interface implemented by FileInputStream
func (recv *FileInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Monitors a file or directory for changes.
//
// To obtain a #GFileMonitor for a file or directory, use
// g_file_monitor(), g_file_monitor_file(), or
// g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are
// monitoring, connect to the #GFileMonitor::changed signal. The
// signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread that the monitor was created in
// (though if the global default main context is blocked, this may
// cause notifications to be blocked even if the thread-default
// context is still running).
/*

C record/class : GFileMonitor
*/
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

	return g
}

func (recv *FileMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FileMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FileMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileMonitor.
func CastToFileMonitor(object *gobject.Object) *FileMonitor {
	return FileMonitorNewFromC(object.ToC())
}

// Unsupported signal 'changed' for FileMonitor : unsupported parameter event_type : type FileMonitorEvent :

// Cancels a file monitor.
/*

C function : g_file_monitor_cancel
*/
func (recv *FileMonitor) Cancel() bool {
	retC := C.g_file_monitor_cancel((*C.GFileMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Emits the #GFileMonitor::changed signal if a change
// has taken place. Should be called from file monitor
// implementations only.
//
// Implementations are responsible to call this method from the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in.
/*

C function : g_file_monitor_emit_event
*/
func (recv *FileMonitor) EmitEvent(child *File, otherFile *File, eventType FileMonitorEvent) {
	c_child := (*C.GFile)(child.ToC())

	c_other_file := (*C.GFile)(otherFile.ToC())

	c_event_type := (C.GFileMonitorEvent)(eventType)

	C.g_file_monitor_emit_event((*C.GFileMonitor)(recv.native), c_child, c_other_file, c_event_type)

	return
}

// Returns whether the monitor is canceled.
/*

C function : g_file_monitor_is_cancelled
*/
func (recv *FileMonitor) IsCancelled() bool {
	retC := C.g_file_monitor_is_cancelled((*C.GFileMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the rate limit to which the @monitor will report
// consecutive change events to the same file.
/*

C function : g_file_monitor_set_rate_limit
*/
func (recv *FileMonitor) SetRateLimit(limitMsecs int32) {
	c_limit_msecs := (C.gint)(limitMsecs)

	C.g_file_monitor_set_rate_limit((*C.GFileMonitor)(recv.native), c_limit_msecs)

	return
}

// GFileOutputStream provides output streams that write their
// content to a file.
//
// GFileOutputStream implements #GSeekable, which allows the output
// stream to jump to arbitrary positions in the file and to truncate
// the file, provided the filesystem of the file supports these
// operations.
//
// To find the position of a file output stream, use g_seekable_tell().
// To find out if a file output stream supports seeking, use
// g_seekable_can_seek().To position a file output stream, use
// g_seekable_seek(). To find out if a file output stream supports
// truncating, use g_seekable_can_truncate(). To truncate a file output
// stream, use g_seekable_truncate().
/*

C record/class : GFileOutputStream
*/
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

	return g
}

func (recv *FileOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStream upcasts to *OutputStream
func (recv *FileOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitary Object to FileOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileOutputStream.
func CastToFileOutputStream(object *gobject.Object) *FileOutputStream {
	return FileOutputStreamNewFromC(object.ToC())
}

// Gets the entity tag for the file when it has been written.
// This must be called after the stream has been written
// and closed, as the etag can change while writing.
/*

C function : g_file_output_stream_get_etag
*/
func (recv *FileOutputStream) GetEtag() string {
	retC := C.g_file_output_stream_get_etag((*C.GFileOutputStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Queries a file output stream for the given @attributes.
// This function blocks while querying the stream. For the asynchronous
// version of this function, see g_file_output_stream_query_info_async().
// While the stream is blocked, the stream will set the pending flag
// internally, and any other operations on the stream will fail with
// %G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with @error being set to
// %G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
// set to %G_IO_ERROR_PENDING), or if querying info is not supported for
// the stream's interface (with @error being set to %G_IO_ERROR_NOT_SUPPORTED). In
// all cases of failure, %NULL will be returned.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %NULL will
// be returned.
/*

C function : g_file_output_stream_query_info
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_output_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finalizes the asynchronous query started
// by g_file_output_stream_query_info_async().
/*

C function : g_file_output_stream_query_info_finish
*/
func (recv *FileOutputStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_output_stream_query_info_finish((*C.GFileOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Completes partial file and directory names given a partial string by
// looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
/*

C record/class : GFilenameCompleter
*/
type FilenameCompleter struct {
	native *C.GFilenameCompleter
}

func FilenameCompleterNewFromC(u unsafe.Pointer) *FilenameCompleter {
	c := (*C.GFilenameCompleter)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleter{native: c}

	return g
}

func (recv *FilenameCompleter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *FilenameCompleter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to FilenameCompleter.
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
var signalFilenameCompleterGotCompletionDataLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalFilenameCompleterGotCompletionDataMap[index].callback
	callback()
}

// Creates a new filename completer.
/*

C function : g_filename_completer_new
*/
func FilenameCompleterNew() *FilenameCompleter {
	retC := C.g_filename_completer_new()
	retGo := FilenameCompleterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Obtains a completion for @initial_text from @completer.
/*

C function : g_filename_completer_get_completion_suffix
*/
func (recv *FilenameCompleter) GetCompletionSuffix(initialText string) string {
	c_initial_text := C.CString(initialText)
	defer C.free(unsafe.Pointer(c_initial_text))

	retC := C.g_filename_completer_get_completion_suffix((*C.GFilenameCompleter)(recv.native), c_initial_text)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_filename_completer_get_completions : no return type

// If @dirs_only is %TRUE, @completer will only
// complete directory names, and not file names.
/*

C function : g_filename_completer_set_dirs_only
*/
func (recv *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	c_dirs_only :=
		boolToGboolean(dirsOnly)

	C.g_filename_completer_set_dirs_only((*C.GFilenameCompleter)(recv.native), c_dirs_only)

	return
}

// Base class for input stream implementations that perform some
// kind of filtering operation on a base stream. Typical examples
// of filtering operations are character set conversion, compression
// and byte order flipping.
/*

C record/class : GFilterInputStream
*/
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

	return g
}

func (recv *FilterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStream upcasts to *InputStream
func (recv *FilterInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FilterInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitary Object to FilterInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FilterInputStream.
func CastToFilterInputStream(object *gobject.Object) *FilterInputStream {
	return FilterInputStreamNewFromC(object.ToC())
}

// Gets the base stream for the filter stream.
/*

C function : g_filter_input_stream_get_base_stream
*/
func (recv *FilterInputStream) GetBaseStream() *InputStream {
	retC := C.g_filter_input_stream_get_base_stream((*C.GFilterInputStream)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the base stream will be closed when @stream is
// closed.
/*

C function : g_filter_input_stream_get_close_base_stream
*/
func (recv *FilterInputStream) GetCloseBaseStream() bool {
	retC := C.g_filter_input_stream_get_close_base_stream((*C.GFilterInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the base stream will be closed when @stream is closed.
/*

C function : g_filter_input_stream_set_close_base_stream
*/
func (recv *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	c_close_base :=
		boolToGboolean(closeBase)

	C.g_filter_input_stream_set_close_base_stream((*C.GFilterInputStream)(recv.native), c_close_base)

	return
}

// Base class for output stream implementations that perform some
// kind of filtering operation on a base stream. Typical examples
// of filtering operations are character set conversion, compression
// and byte order flipping.
/*

C record/class : GFilterOutputStream
*/
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

	return g
}

func (recv *FilterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStream upcasts to *OutputStream
func (recv *FilterOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FilterOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitary Object to FilterOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a FilterOutputStream.
func CastToFilterOutputStream(object *gobject.Object) *FilterOutputStream {
	return FilterOutputStreamNewFromC(object.ToC())
}

// Gets the base stream for the filter stream.
/*

C function : g_filter_output_stream_get_base_stream
*/
func (recv *FilterOutputStream) GetBaseStream() *OutputStream {
	retC := C.g_filter_output_stream_get_base_stream((*C.GFilterOutputStream)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns whether the base stream will be closed when @stream is
// closed.
/*

C function : g_filter_output_stream_get_close_base_stream
*/
func (recv *FilterOutputStream) GetCloseBaseStream() bool {
	retC := C.g_filter_output_stream_get_close_base_stream((*C.GFilterOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether the base stream will be closed when @stream is closed.
/*

C function : g_filter_output_stream_set_close_base_stream
*/
func (recv *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	c_close_base :=
		boolToGboolean(closeBase)

	C.g_filter_output_stream_set_close_base_stream((*C.GFilterOutputStream)(recv.native), c_close_base)

	return
}

// Provides an interface and default functions for loading and unloading
// modules. This is used internally to make GIO extensible, but can also
// be used by others to implement module loading.
/*

C record/class : GIOModule
*/
type IOModule struct {
	native *C.GIOModule
}

func IOModuleNewFromC(u unsafe.Pointer) *IOModule {
	c := (*C.GIOModule)(u)
	if c == nil {
		return nil
	}

	g := &IOModule{native: c}

	return g
}

func (recv *IOModule) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CastToWidget down casts any arbitary Object to IOModule.
// Exercise care, as this is a potentially dangerous function if the Object is not a IOModule.
func CastToIOModule(object *gobject.Object) *IOModule {
	return IOModuleNewFromC(object.ToC())
}

// Creates a new GIOModule that will load the specific
// shared library when in use.
/*

C function : g_io_module_new
*/
func IOModuleNew(filename string) *IOModule {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_io_module_new(c_filename)
	retGo := IOModuleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : g_io_module_load

// Blacklisted : g_io_module_unload

// TypePlugin returns the TypePlugin interface implemented by IOModule
func (recv *IOModule) TypePlugin() *gobject.TypePlugin {
	return gobject.TypePluginNewFromC(recv.ToC())
}

// GIOStream represents an object that has both read and write streams.
// Generally the two streams act as separate input and output streams,
// but they share some common resources and state. For instance, for
// seekable streams, both streams may use the same position.
//
// Examples of #GIOStream objects are #GSocketConnection, which represents
// a two-way network connection; and #GFileIOStream, which represents a
// file handle opened in read-write mode.
//
// To do the actual reading and writing you need to get the substreams
// with g_io_stream_get_input_stream() and g_io_stream_get_output_stream().
//
// The #GIOStream object owns the input and the output streams, not the other
// way around, so keeping the substreams alive will not keep the #GIOStream
// object alive. If the #GIOStream object is freed it will be closed, thus
// closing the substreams, so even if the substreams stay alive they will
// always return %G_IO_ERROR_CLOSED for all operations.
//
// To close a stream use g_io_stream_close() which will close the common
// stream object and also the individual substreams. You can also close
// the substreams themselves. In most cases this only marks the
// substream as closed, so further I/O on it fails but common state in the
// #GIOStream may still be open. However, some streams may support
// "half-closed" states where one direction of the stream is actually shut down.
//
// Operations on #GIOStreams cannot be started while another operation on the
// #GIOStream or its substreams is in progress. Specifically, an application can
// read from the #GInputStream and write to the #GOutputStream simultaneously
// (either in separate threads, or as asynchronous operations in the same
// thread), but an application cannot start any #GIOStream operation while there
// is a #GIOStream, #GInputStream or #GOutputStream operation in progress, and
// an application can’t start any #GInputStream or #GOutputStream operation
// while there is a #GIOStream operation in progress.
//
// This is a product of individual stream operations being associated with a
// given #GMainContext (the thread-default context at the time the operation was
// started), rather than entire streams being associated with a single
// #GMainContext.
//
// GIO may run operations on #GIOStreams from other (worker) threads, and this
// may be exposed to application code in the behaviour of wrapper streams, such
// as #GBufferedInputStream or #GTlsConnection. With such wrapper APIs,
// application code may only run operations on the base (wrapped) stream when
// the wrapper stream is idle. Note that the semantics of such operations may
// not be well-defined due to the state the wrapper stream leaves the base
// stream in (though they are guaranteed not to crash).
/*

C record/class : GIOStream
*/
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

	return g
}

func (recv *IOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *IOStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to IOStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a IOStream.
func CastToIOStream(object *gobject.Object) *IOStream {
	return IOStreamNewFromC(object.ToC())
}

// #GInetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to
// look up the #GInetAddress for a hostname. Use
// g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a
// #GInetAddress.
//
// To actually connect to a remote host, you will need a
// #GInetSocketAddress (which includes a #GInetAddress as well as a
// port number).
/*

C record/class : GInetAddress
*/
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

	return g
}

func (recv *InetAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *InetAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to InetAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetAddress.
func CastToInetAddress(object *gobject.Object) *InetAddress {
	return InetAddressNewFromC(object.ToC())
}

// An IPv4 or IPv6 socket address; that is, the combination of a
// #GInetAddress and a port number.
/*

C record/class : GInetSocketAddress
*/
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

	return g
}

func (recv *InetSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddress upcasts to *SocketAddress
func (recv *InetSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *InetSocketAddress) Object() *gobject.Object {
	return recv.SocketAddress().Object()
}

// CastToWidget down casts any arbitary Object to InetSocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetSocketAddress.
func CastToInetSocketAddress(object *gobject.Object) *InetSocketAddress {
	return InetSocketAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by InetSocketAddress
func (recv *InetSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// #GInputStream has functions to read from a stream (g_input_stream_read()),
// to close a stream (g_input_stream_close()) and to skip some content
// (g_input_stream_skip()).
//
// To copy the content of an input stream to an output stream without
// manually handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for #GIOStream for details of thread safety of
// streaming APIs.
//
// All of these functions have async variants too.
/*

C record/class : GInputStream
*/
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

	return g
}

func (recv *InputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *InputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to InputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a InputStream.
func CastToInputStream(object *gobject.Object) *InputStream {
	return InputStreamNewFromC(object.ToC())
}

// Clears the pending flag on @stream.
/*

C function : g_input_stream_clear_pending
*/
func (recv *InputStream) ClearPending() {
	C.g_input_stream_clear_pending((*C.GInputStream)(recv.native))

	return
}

// Closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return %G_IO_ERROR_CLOSED.
// Closing a stream multiple times will not return an error.
//
// Streams will be automatically closed when the last reference
// is dropped, but you might want to call this function to make sure
// resources are released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file descriptor)
// open after the stream is closed. See the documentation for the individual
// stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to
// close will still return %G_IO_ERROR_CLOSED for all operations. Still, it
// is important to check and report the error to the user.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but some streams
// can use a faster close that doesn't block to e.g. check errors.
/*

C function : g_input_stream_close
*/
func (recv *InputStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_close((*C.GInputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes closing a stream asynchronously, started from g_input_stream_close_async().
/*

C function : g_input_stream_close_finish
*/
func (recv *InputStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_close_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if an input stream has pending actions.
/*

C function : g_input_stream_has_pending
*/
func (recv *InputStream) HasPending() bool {
	retC := C.g_input_stream_has_pending((*C.GInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if an input stream is closed.
/*

C function : g_input_stream_is_closed
*/
func (recv *InputStream) IsClosed() bool {
	retC := C.g_input_stream_is_closed((*C.GInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tries to read @count bytes from the stream into the buffer starting at
// @buffer. Will block during this read.
//
// If count is zero returns zero and does nothing. A value of @count
// larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. near the end of a file. Zero is returned on end of file
// (or if @count is zero),  but never otherwise.
//
// The returned @buffer is not a nul-terminated string, it can contain nul bytes
// at any position, and this function doesn't nul-terminate the @buffer.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
/*

C function : g_input_stream_read
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to read @count bytes from the stream into the buffer starting at
// @buffer. Will block during this read.
//
// This function is similar to g_input_stream_read(), except it tries to
// read as many bytes as requested, only stopping on an error or end of stream.
//
// On a successful read of @count bytes, or if we reached the end of the
// stream,  %TRUE is returned, and @bytes_read is set to the number of bytes
// read into @buffer.
//
// If there is an error during the operation %FALSE is returned and @error
// is set to indicate the error status.
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_read will be set to the number of bytes that were successfully
// read before the error was encountered.  This functionality is only
// available from C.  If you need it from another language then you must
// write your own loop around g_input_stream_read().
/*

C function : g_input_stream_read_all
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	return retGo, bytesRead, goThrowableError
}

// Unsupported : g_input_stream_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous stream read operation.
/*

C function : g_input_stream_read_finish
*/
func (recv *InputStream) ReadFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_read_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return %FALSE and set
// @error.
/*

C function : g_input_stream_set_pending
*/
func (recv *InputStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_input_stream_set_pending((*C.GInputStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to skip @count bytes from the stream. Will block during the operation.
//
// This is identical to g_input_stream_read(), from a behaviour standpoint,
// but the bytes that are skipped are not returned to the user. Some
// streams have an implementation that is more efficient than reading the data.
//
// This function is optional for inherited classes, as the default implementation
// emulates it using read.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
/*

C function : g_input_stream_skip
*/
func (recv *InputStream) Skip(count uint64, cancellable *Cancellable) (int64, error) {
	c_count := (C.gsize)(count)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_input_stream_skip((*C.GInputStream)(recv.native), c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_input_stream_skip_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a stream skip operation.
/*

C function : g_input_stream_skip_finish
*/
func (recv *InputStream) SkipFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_input_stream_skip_finish((*C.GInputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GListStore is a simple implementation of #GListModel that stores all
// items in memory.
//
// It provides insertions, deletions, and lookups in logarithmic time
// with a fast path for the common case of iterating the list linearly.
/*

C record/class : GListStore
*/
type ListStore struct {
	native *C.GListStore
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ListStore) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ListStore.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListStore.
func CastToListStore(object *gobject.Object) *ListStore {
	return ListStoreNewFromC(object.ToC())
}

// ListModel returns the ListModel interface implemented by ListStore
func (recv *ListStore) ListModel() *ListModel {
	return ListModelNewFromC(recv.ToC())
}

// #GMemoryInputStream is a class for using arbitrary
// memory chunks as input for GIO streaming input operations.
//
// As of GLib 2.34, #GMemoryInputStream implements
// #GPollableInputStream.
/*

C record/class : GMemoryInputStream
*/
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

	return g
}

func (recv *MemoryInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStream upcasts to *InputStream
func (recv *MemoryInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *MemoryInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitary Object to MemoryInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a MemoryInputStream.
func CastToMemoryInputStream(object *gobject.Object) *MemoryInputStream {
	return MemoryInputStreamNewFromC(object.ToC())
}

// Creates a new empty #GMemoryInputStream.
/*

C function : g_memory_input_stream_new
*/
func MemoryInputStreamNew() *MemoryInputStream {
	retC := C.g_memory_input_stream_new()
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

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

// #GMemoryOutputStream is a class for using arbitrary
// memory chunks as output for GIO streaming output operations.
//
// As of GLib 2.34, #GMemoryOutputStream trivially implements
// #GPollableOutputStream: it always polls as ready.
/*

C record/class : GMemoryOutputStream
*/
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

	return g
}

func (recv *MemoryOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStream upcasts to *OutputStream
func (recv *MemoryOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *MemoryOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitary Object to MemoryOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a MemoryOutputStream.
func CastToMemoryOutputStream(object *gobject.Object) *MemoryOutputStream {
	return MemoryOutputStreamNewFromC(object.ToC())
}

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc (GReallocFunc) for param realloc_function

// Gets any loaded data from the @ostream.
//
// Note that the returned pointer may become invalid on the next
// write or truncate operation on the stream.
/*

C function : g_memory_output_stream_get_data
*/
func (recv *MemoryOutputStream) GetData() uintptr {
	retC := C.g_memory_output_stream_get_data((*C.GMemoryOutputStream)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets the size of the currently allocated data area (available from
// g_memory_output_stream_get_data()).
//
// You probably don't want to use this function on resizable streams.
// See g_memory_output_stream_get_data_size() instead.  For resizable
// streams the size returned by this function is an implementation
// detail and may be change at any time in response to operations on the
// stream.
//
// If the stream is fixed-sized (ie: no realloc was passed to
// g_memory_output_stream_new()) then this is the maximum size of the
// stream and further writes will return %G_IO_ERROR_NO_SPACE.
//
// In any case, if you want the number of bytes currently written to the
// stream, use g_memory_output_stream_get_data_size().
/*

C function : g_memory_output_stream_get_size
*/
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

// #GMountOperation provides a mechanism for interacting with the user.
// It can be used for authenticating mountable operations, such as loop
// mounting files, hard drive partitions or server locations. It can
// also be used to ask the user questions or show a list of applications
// preventing unmount or eject operations from completing.
//
// Note that #GMountOperation is used for more than just #GMount
// objects – for example it is also used in g_drive_start() and
// g_drive_stop().
//
// Users should instantiate a subclass of this that implements all the
// various callbacks to show the required dialogs, such as
// #GtkMountOperation. If no user interaction is desired (for example
// when automounting filesystems at login time), usually %NULL can be
// passed, see each method taking a #GMountOperation for details.
/*

C record/class : GMountOperation
*/
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

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *MountOperation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to MountOperation.
// Exercise care, as this is a potentially dangerous function if the Object is not a MountOperation.
func CastToMountOperation(object *gobject.Object) *MountOperation {
	return MountOperationNewFromC(object.ToC())
}

// Unsupported signal 'ask-password' for MountOperation : unsupported parameter message : type utf8 :

// Unsupported signal 'ask-question' for MountOperation : unsupported parameter message : type utf8 :

// Unsupported signal 'reply' for MountOperation : unsupported parameter result : type MountOperationResult :

// Creates a new mount operation.
/*

C function : g_mount_operation_new
*/
func MountOperationNew() *MountOperation {
	retC := C.g_mount_operation_new()
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Check to see whether the mount operation is being used
// for an anonymous user.
/*

C function : g_mount_operation_get_anonymous
*/
func (recv *MountOperation) GetAnonymous() bool {
	retC := C.g_mount_operation_get_anonymous((*C.GMountOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets a choice from the mount operation.
/*

C function : g_mount_operation_get_choice
*/
func (recv *MountOperation) GetChoice() int32 {
	retC := C.g_mount_operation_get_choice((*C.GMountOperation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the domain of the mount operation.
/*

C function : g_mount_operation_get_domain
*/
func (recv *MountOperation) GetDomain() string {
	retC := C.g_mount_operation_get_domain((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets a password from the mount operation.
/*

C function : g_mount_operation_get_password
*/
func (recv *MountOperation) GetPassword() string {
	retC := C.g_mount_operation_get_password((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the state of saving passwords for the mount operation.
/*

C function : g_mount_operation_get_password_save
*/
func (recv *MountOperation) GetPasswordSave() PasswordSave {
	retC := C.g_mount_operation_get_password_save((*C.GMountOperation)(recv.native))
	retGo := (PasswordSave)(retC)

	return retGo
}

// Get the user name from the mount operation.
/*

C function : g_mount_operation_get_username
*/
func (recv *MountOperation) GetUsername() string {
	retC := C.g_mount_operation_get_username((*C.GMountOperation)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Emits the #GMountOperation::reply signal.
/*

C function : g_mount_operation_reply
*/
func (recv *MountOperation) Reply(result MountOperationResult) {
	c_result := (C.GMountOperationResult)(result)

	C.g_mount_operation_reply((*C.GMountOperation)(recv.native), c_result)

	return
}

// Sets the mount operation to use an anonymous user if @anonymous is %TRUE.
/*

C function : g_mount_operation_set_anonymous
*/
func (recv *MountOperation) SetAnonymous(anonymous bool) {
	c_anonymous :=
		boolToGboolean(anonymous)

	C.g_mount_operation_set_anonymous((*C.GMountOperation)(recv.native), c_anonymous)

	return
}

// Sets a default choice for the mount operation.
/*

C function : g_mount_operation_set_choice
*/
func (recv *MountOperation) SetChoice(choice int32) {
	c_choice := (C.int)(choice)

	C.g_mount_operation_set_choice((*C.GMountOperation)(recv.native), c_choice)

	return
}

// Sets the mount operation's domain.
/*

C function : g_mount_operation_set_domain
*/
func (recv *MountOperation) SetDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_mount_operation_set_domain((*C.GMountOperation)(recv.native), c_domain)

	return
}

// Sets the mount operation's password to @password.
/*

C function : g_mount_operation_set_password
*/
func (recv *MountOperation) SetPassword(password string) {
	c_password := C.CString(password)
	defer C.free(unsafe.Pointer(c_password))

	C.g_mount_operation_set_password((*C.GMountOperation)(recv.native), c_password)

	return
}

// Sets the state of saving passwords for the mount operation.
/*

C function : g_mount_operation_set_password_save
*/
func (recv *MountOperation) SetPasswordSave(save PasswordSave) {
	c_save := (C.GPasswordSave)(save)

	C.g_mount_operation_set_password_save((*C.GMountOperation)(recv.native), c_save)

	return
}

// Sets the user name within @op to @username.
/*

C function : g_mount_operation_set_username
*/
func (recv *MountOperation) SetUsername(username string) {
	c_username := C.CString(username)
	defer C.free(unsafe.Pointer(c_username))

	C.g_mount_operation_set_username((*C.GMountOperation)(recv.native), c_username)

	return
}

/*

C record/class : GNativeVolumeMonitor
*/
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

	return g
}

func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeMonitor upcasts to *VolumeMonitor
func (recv *NativeVolumeMonitor) VolumeMonitor() *VolumeMonitor {
	return VolumeMonitorNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *NativeVolumeMonitor) Object() *gobject.Object {
	return recv.VolumeMonitor().Object()
}

// CastToWidget down casts any arbitary Object to NativeVolumeMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a NativeVolumeMonitor.
func CastToNativeVolumeMonitor(object *gobject.Object) *NativeVolumeMonitor {
	return NativeVolumeMonitorNewFromC(object.ToC())
}

// #GNetworkAddress provides an easy way to resolve a hostname and
// then attempt to connect to that host, handling the possibility of
// multiple IP addresses and multiple address families.
//
// See #GSocketConnectable for and example of using the connectable
// interface.
/*

C record/class : GNetworkAddress
*/
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

	return g
}

func (recv *NetworkAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *NetworkAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to NetworkAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a NetworkAddress.
func CastToNetworkAddress(object *gobject.Object) *NetworkAddress {
	return NetworkAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkAddress
func (recv *NetworkAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Like #GNetworkAddress does with hostnames, #GNetworkService
// provides an easy way to resolve a SRV record, and then attempt to
// connect to one of the hosts that implements that service, handling
// service priority/weighting, multiple IP addresses, and multiple
// address families.
//
// See #GSrvTarget for more information about SRV records, and see
// #GSocketConnectable for and example of using the connectable
// interface.
/*

C record/class : GNetworkService
*/
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

	return g
}

func (recv *NetworkService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *NetworkService) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to NetworkService.
// Exercise care, as this is a potentially dangerous function if the Object is not a NetworkService.
func CastToNetworkService(object *gobject.Object) *NetworkService {
	return NetworkServiceNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkService
func (recv *NetworkService) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// #GOutputStream has functions to write to a stream (g_output_stream_write()),
// to close a stream (g_output_stream_close()) and to flush pending writes
// (g_output_stream_flush()).
//
// To copy the content of an input stream to an output stream without
// manually handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for #GIOStream for details of thread safety of
// streaming APIs.
//
// All of these functions have async variants too.
/*

C record/class : GOutputStream
*/
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

	return g
}

func (recv *OutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *OutputStream) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to OutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a OutputStream.
func CastToOutputStream(object *gobject.Object) *OutputStream {
	return OutputStreamNewFromC(object.ToC())
}

// Clears the pending flag on @stream.
/*

C function : g_output_stream_clear_pending
*/
func (recv *OutputStream) ClearPending() {
	C.g_output_stream_clear_pending((*C.GOutputStream)(recv.native))

	return
}

// Closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return %G_IO_ERROR_CLOSED.
// Closing a stream multiple times will not return an error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
//
// Streams will be automatically closed when the last reference
// is dropped, but you might want to call this function to make sure
// resources are released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file descriptor)
// open after the stream is closed. See the documentation for the individual
// stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to
// close will still return %G_IO_ERROR_CLOSED for all operations. Still, it
// is important to check and report the error to the user, otherwise
// there might be a loss of data as all data might not be written.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but there some streams
// can use a faster close that doesn't block to e.g. check errors. On
// cancellation (as with any error) there is no guarantee that all written
// data will reach the target.
/*

C function : g_output_stream_close
*/
func (recv *OutputStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_close((*C.GOutputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_output_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Closes an output stream.
/*

C function : g_output_stream_close_finish
*/
func (recv *OutputStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_close_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Forces a write of all user-space buffered data for the given
// @stream. Will block during the operation. Closing the stream will
// implicitly cause a flush.
//
// This function is optional for inherited classes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_output_stream_flush
*/
func (recv *OutputStream) Flush(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_output_stream_flush((*C.GOutputStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_output_stream_flush_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes flushing an output stream.
/*

C function : g_output_stream_flush_finish
*/
func (recv *OutputStream) FlushFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_flush_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if an output stream has pending actions.
/*

C function : g_output_stream_has_pending
*/
func (recv *OutputStream) HasPending() bool {
	retC := C.g_output_stream_has_pending((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if an output stream has already been closed.
/*

C function : g_output_stream_is_closed
*/
func (recv *OutputStream) IsClosed() bool {
	retC := C.g_output_stream_is_closed((*C.GOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return %FALSE and set
// @error.
/*

C function : g_output_stream_set_pending
*/
func (recv *OutputStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_output_stream_set_pending((*C.GOutputStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Splices an input stream into an output stream.
/*

C function : g_output_stream_splice
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_output_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous stream splice operation.
/*

C function : g_output_stream_splice_finish
*/
func (recv *OutputStream) SpliceFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_splice_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to write @count bytes from @buffer into the stream. Will block
// during the operation.
//
// If count is 0, returns 0 and does nothing. A value of @count
// larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes written to the stream is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. on a partial I/O error, or if there is not enough
// storage in the stream. All writes block until at least one byte
// is written or an error occurs; 0 is never returned (unless
// @count is 0).
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
/*

C function : g_output_stream_write
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tries to write @count bytes from @buffer into the stream. Will block
// during the operation.
//
// This function is similar to g_output_stream_write(), except it tries to
// write as many bytes as requested, only stopping on an error.
//
// On a successful write of @count bytes, %TRUE is returned, and @bytes_written
// is set to @count.
//
// If there is an error during the operation %FALSE is returned and @error
// is set to indicate the error status.
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C.  If you need it from another
// language then you must write your own loop around
// g_output_stream_write().
/*

C function : g_output_stream_write_all
*/
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten, goThrowableError
}

// Unsupported : g_output_stream_write_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_output_stream_write_bytes

// Unsupported : g_output_stream_write_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes a stream write-from-#GBytes operation.
/*

C function : g_output_stream_write_bytes_finish
*/
func (recv *OutputStream) WriteBytesFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_bytes_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Finishes a stream write operation.
/*

C function : g_output_stream_write_finish
*/
func (recv *OutputStream) WriteFinish(result *AsyncResult) (int64, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_output_stream_write_finish((*C.GOutputStream)(recv.native), c_result, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// A #GPermission represents the status of the caller's permission to
// perform a certain action.
//
// You can query if the action is currently allowed and if it is
// possible to acquire the permission so that the action will be allowed
// in the future.
//
// There is also an API to actually acquire the permission and one to
// release it.
//
// As an example, a #GPermission might represent the ability for the
// user to write to a #GSettings object.  This #GPermission object could
// then be used to decide if it is appropriate to show a "Click here to
// unlock" button in a dialog and to provide the mechanism to invoke
// when that button is clicked.
/*

C record/class : GPermission
*/
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

	return g
}

func (recv *Permission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Permission) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Permission.
// Exercise care, as this is a potentially dangerous function if the Object is not a Permission.
func CastToPermission(object *gobject.Object) *Permission {
	return PermissionNewFromC(object.ToC())
}

// A subclass of #GSocketAddressEnumerator that takes another address
// enumerator and wraps its results in #GProxyAddresses as
// directed by the default #GProxyResolver.
/*

C record/class : GProxyAddressEnumerator
*/
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

	return g
}

func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddressEnumerator upcasts to *SocketAddressEnumerator
func (recv *ProxyAddressEnumerator) SocketAddressEnumerator() *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *ProxyAddressEnumerator) Object() *gobject.Object {
	return recv.SocketAddressEnumerator().Object()
}

// CastToWidget down casts any arbitary Object to ProxyAddressEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProxyAddressEnumerator.
func CastToProxyAddressEnumerator(object *gobject.Object) *ProxyAddressEnumerator {
	return ProxyAddressEnumeratorNewFromC(object.ToC())
}

// #GResolver provides cancellable synchronous and asynchronous DNS
// resolution, for hostnames (g_resolver_lookup_by_address(),
// g_resolver_lookup_by_name() and their async variants) and SRV
// (service) records (g_resolver_lookup_service()).
//
// #GNetworkAddress and #GNetworkService provide wrappers around
// #GResolver functionality that also implement #GSocketConnectable,
// making it easy to connect to a remote host/service.
/*

C record/class : GResolver
*/
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

	return g
}

func (recv *Resolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Resolver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Resolver.
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
var signalResolverReloadLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalResolverReloadMap[index].callback
	callback()
}

// The #GSettings class provides a convenient API for storing and retrieving
// application settings.
//
// Reads and writes can be considered to be non-blocking.  Reading
// settings with #GSettings is typically extremely fast: on
// approximately the same order of magnitude (but slower than) a
// #GHashTable lookup.  Writing settings is also extremely fast in terms
// of time to return to your application, but can be extremely expensive
// for other threads and other processes.  Many settings backends
// (including dconf) have lazy initialisation which means in the common
// case of the user using their computer without modifying any settings
// a lot of work can be avoided.  For dconf, the D-Bus service doesn't
// even need to be started in this case.  For this reason, you should
// only ever modify #GSettings keys in response to explicit user action.
// Particular care should be paid to ensure that modifications are not
// made during startup -- for example, when setting the initial value
// of preferences widgets.  The built-in g_settings_bind() functionality
// is careful not to write settings in response to notify signals as a
// result of modifications that it makes to widgets.
//
// When creating a GSettings instance, you have to specify a schema
// that describes the keys in your settings and their types and default
// values, as well as some other information.
//
// Normally, a schema has a fixed path that determines where the settings
// are stored in the conceptual global tree of settings. However, schemas
// can also be '[relocatable][gsettings-relocatable]', i.e. not equipped with
// a fixed path. This is
// useful e.g. when the schema describes an 'account', and you want to be
// able to store a arbitrary number of accounts.
//
// Paths must start with and end with a forward slash character ('/')
// and must not contain two sequential slash characters.  Paths should
// be chosen based on a domain name associated with the program or
// library to which the settings belong.  Examples of paths are
// "/org/gtk/settings/file-chooser/" and "/ca/desrt/dconf-editor/".
// Paths should not start with "/apps/", "/desktop/" or "/system/" as
// they often did in GConf.
//
// Unlike other configuration systems (like GConf), GSettings does not
// restrict keys to basic types like strings and numbers. GSettings stores
// values as #GVariant, and allows any #GVariantType for keys. Key names
// are restricted to lowercase characters, numbers and '-'. Furthermore,
// the names must begin with a lowercase character, must not end
// with a '-', and must not contain consecutive dashes.
//
// Similar to GConf, the default values in GSettings schemas can be
// localized, but the localized values are stored in gettext catalogs
// and looked up with the domain that is specified in the
// `gettext-domain` attribute of the <schemalist> or <schema>
// elements and the category that is specified in the `l10n` attribute of
// the <default> element. The string which is translated includes all text in
// the <default> element, including any surrounding quotation marks.
//
// The `l10n` attribute must be set to `messages` or `time`, and sets the
// [locale category for
// translation](https://www.gnu.org/software/gettext/manual/html_node/Aspects.html#index-locale-categories-1).
// The `messages` category should be used by default; use `time` for
// translatable date or time formats. A translation comment can be added as an
// XML comment immediately above the <default> element — it is recommended to
// add these comments to aid translators understand the meaning and
// implications of the default value. An optional translation `context`
// attribute can be set on the <default> element to disambiguate multiple
// defaults which use the same string.
//
// For example:
// |[
// <!-- Translators: A list of words which are not allowed to be typed, in
// GVariant serialization syntax.
// See: https://developer.gnome.org/glib/stable/gvariant-text.html -->
// <default l10n='messages' context='Banned words'>['bad', 'words']</default>
// ]|
//
// Translations of default values must remain syntactically valid serialized
// #GVariants (e.g. retaining any surrounding quotation marks) or runtime
// errors will occur.
//
// GSettings uses schemas in a compact binary form that is created
// by the [glib-compile-schemas][glib-compile-schemas]
// utility. The input is a schema description in an XML format.
//
// A DTD for the gschema XML format can be found here:
// [gschema.dtd](https://git.gnome.org/browse/glib/tree/gio/gschema.dtd)
//
// The [glib-compile-schemas][glib-compile-schemas] tool expects schema
// files to have the extension `.gschema.xml`.
//
// At runtime, schemas are identified by their id (as specified in the
// id attribute of the <schema> element). The convention for schema
// ids is to use a dotted name, similar in style to a D-Bus bus name,
// e.g. "org.gnome.SessionManager". In particular, if the settings are
// for a specific service that owns a D-Bus bus name, the D-Bus bus name
// and schema id should match. For schemas which deal with settings not
// associated with one named application, the id should not use
// StudlyCaps, e.g. "org.gnome.font-rendering".
//
// In addition to #GVariant types, keys can have types that have
// enumerated types. These can be described by a <choice>,
// <enum> or <flags> element, as seen in the
// [example][schema-enumerated]. The underlying type of such a key
// is string, but you can use g_settings_get_enum(), g_settings_set_enum(),
// g_settings_get_flags(), g_settings_set_flags() access the numeric values
// corresponding to the string value of enum and flags keys.
//
// An example for default value:
// |[
// <schemalist>
// <schema id="org.gtk.Test" path="/org/gtk/Test/" gettext-domain="test">
//
// <key name="greeting" type="s">
// <default l10n="messages">"Hello, earthlings"</default>
// <summary>A greeting</summary>
// <description>
// Greeting of the invading martians
// </description>
// </key>
//
// <key name="box" type="(ii)">
// <default>(20,30)</default>
// </key>
//
// </schema>
// </schemalist>
// ]|
//
// An example for ranges, choices and enumerated types:
// |[
// <schemalist>
//
// <enum id="org.gtk.Test.myenum">
// <value nick="first" value="1"/>
// <value nick="second" value="2"/>
// </enum>
//
// <flags id="org.gtk.Test.myflags">
// <value nick="flag1" value="1"/>
// <value nick="flag2" value="2"/>
// <value nick="flag3" value="4"/>
// </flags>
//
// <schema id="org.gtk.Test">
//
// <key name="key-with-range" type="i">
// <range min="1" max="100"/>
// <default>10</default>
// </key>
//
// <key name="key-with-choices" type="s">
// <choices>
// <choice value='Elisabeth'/>
// <choice value='Annabeth'/>
// <choice value='Joe'/>
// </choices>
// <aliases>
// <alias value='Anna' target='Annabeth'/>
// <alias value='Beth' target='Elisabeth'/>
// </aliases>
// <default>'Joe'</default>
// </key>
//
// <key name='enumerated-key' enum='org.gtk.Test.myenum'>
// <default>'first'</default>
// </key>
//
// <key name='flags-key' flags='org.gtk.Test.myflags'>
// <default>["flag1","flag2"]</default>
// </key>
// </schema>
// </schemalist>
// ]|
//
// ## Vendor overrides
//
// Default values are defined in the schemas that get installed by
// an application. Sometimes, it is necessary for a vendor or distributor
// to adjust these defaults. Since patching the XML source for the schema
// is inconvenient and error-prone,
// [glib-compile-schemas][glib-compile-schemas] reads so-called vendor
// override' files. These are keyfiles in the same directory as the XML
// schema sources which can override default values. The schema id serves
// as the group name in the key file, and the values are expected in
// serialized GVariant form, as in the following example:
// |[
// [org.gtk.Example]
// key1='string'
// key2=1.5
// ]|
//
// glib-compile-schemas expects schema files to have the extension
// `.gschema.override`.
//
// ## Binding
//
// A very convenient feature of GSettings lets you bind #GObject properties
// directly to settings, using g_settings_bind(). Once a GObject property
// has been bound to a setting, changes on either side are automatically
// propagated to the other side. GSettings handles details like mapping
// between GObject and GVariant types, and preventing infinite cycles.
//
// This makes it very easy to hook up a preferences dialog to the
// underlying settings. To make this even more convenient, GSettings
// looks for a boolean property with the name "sensitivity" and
// automatically binds it to the writability of the bound setting.
// If this 'magic' gets in the way, it can be suppressed with the
// #G_SETTINGS_BIND_NO_SENSITIVITY flag.
//
// ## Relocatable schemas # {#gsettings-relocatable}
//
// A relocatable schema is one with no `path` attribute specified on its
// <schema> element. By using g_settings_new_with_path(), a #GSettings object
// can be instantiated for a relocatable schema, assigning a path to the
// instance. Paths passed to g_settings_new_with_path() will typically be
// constructed dynamically from a constant prefix plus some form of instance
// identifier; but they must still be valid GSettings paths. Paths could also
// be constant and used with a globally installed schema originating from a
// dependency library.
//
// For example, a relocatable schema could be used to store geometry information
// for different windows in an application. If the schema ID was
// `org.foo.MyApp.Window`, it could be instantiated for paths
// `/org/foo/MyApp/main/`, `/org/foo/MyApp/document-1/`,
// `/org/foo/MyApp/document-2/`, etc. If any of the paths are well-known
// they can be specified as <child> elements in the parent schema, e.g.:
// |[
// <schema id="org.foo.MyApp" path="/org/foo/MyApp/">
// <child name="main" schema="org.foo.MyApp.Window"/>
// </schema>
// ]|
//
// ## Build system integration # {#gsettings-build-system}
//
// GSettings comes with autotools integration to simplify compiling and
// installing schemas. To add GSettings support to an application, add the
// following to your `configure.ac`:
// |[
// GLIB_GSETTINGS
// ]|
//
// In the appropriate `Makefile.am`, use the following snippet to compile and
// install the named schema:
// |[
// gsettings_SCHEMAS = org.foo.MyApp.gschema.xml
// EXTRA_DIST = $(gsettings_SCHEMAS)
//
// @GSETTINGS_RULES@
// ]|
//
// No changes are needed to the build system to mark a schema XML file for
// translation. Assuming it sets the `gettext-domain` attribute, a schema may
// be marked for translation by adding it to `POTFILES.in`, assuming gettext
// 0.19 is in use (the preferred method for translation):
// |[
// data/org.foo.MyApp.gschema.xml
// ]|
//
// Alternatively, if intltool 0.50.1 is in use:
// |[
// [type: gettext/gsettings]data/org.foo.MyApp.gschema.xml
// ]|
//
// GSettings will use gettext to look up translations for the <summary> and
// <description> elements, and also any <default> elements which have a `l10n`
// attribute set. Translations must not be included in the `.gschema.xml` file
// by the build system, for example by using intltool XML rules with a
// `.gschema.xml.in` template.
//
// If an enumerated type defined in a C header file is to be used in a GSettings
// schema, it can either be defined manually using an <enum> element in the
// schema XML, or it can be extracted automatically from the C header. This
// approach is preferred, as it ensures the two representations are always
// synchronised. To do so, add the following to the relevant `Makefile.am`:
// |[
// gsettings_ENUM_NAMESPACE = org.foo.MyApp
// gsettings_ENUM_FILES = my-app-enums.h my-app-misc.h
// ]|
//
// `gsettings_ENUM_NAMESPACE` specifies the schema namespace for the enum files,
// which are specified in `gsettings_ENUM_FILES`. This will generate a
// `org.foo.MyApp.enums.xml` file containing the extracted enums, which will be
// automatically included in the schema compilation, install and uninstall
// rules. It should not be committed to version control or included in
// `EXTRA_DIST`.
/*

C record/class : GSettings
*/
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

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Settings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Settings.
// Exercise care, as this is a potentially dangerous function if the Object is not a Settings.
func CastToSettings(object *gobject.Object) *Settings {
	return SettingsNewFromC(object.ToC())
}

// Unsupported signal 'change-event' for Settings : unsupported parameter keys : no param type

// Unsupported signal 'changed' for Settings : unsupported parameter key : type utf8 :

// Unsupported signal 'writable-change-event' for Settings : unsupported parameter key : type guint :

// Unsupported signal 'writable-changed' for Settings : unsupported parameter key : type utf8 :

// Applies any changes that have been made to the settings.  This
// function does nothing unless @settings is in 'delay-apply' mode;
// see g_settings_delay().  In the normal case settings are always
// applied immediately.
/*

C function : g_settings_apply
*/
func (recv *Settings) Apply() {
	C.g_settings_apply((*C.GSettings)(recv.native))

	return
}

// Unsupported : g_settings_get_mapped : unsupported parameter mapping : no type generator for SettingsGetMapping (GSettingsGetMapping) for param mapping

// Unsupported : g_settings_list_children : no return type

// Unsupported : g_settings_list_keys : no return type

// Resets @key to its default value.
//
// This call resets the key, as much as possible, to its default value.
// That might the value specified in the schema or the one set by the
// administrator.
/*

C function : g_settings_reset
*/
func (recv *Settings) Reset(key string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	C.g_settings_reset((*C.GSettings)(recv.native), c_key)

	return
}

// Reverts all non-applied changes to the settings.  This function
// does nothing unless @settings is in 'delay-apply' mode; see
// g_settings_delay().  In the normal case settings are always applied
// immediately.
//
// Change notifications will be emitted for affected keys.
/*

C function : g_settings_revert
*/
func (recv *Settings) Revert() {
	C.g_settings_revert((*C.GSettings)(recv.native))

	return
}

// Looks up the enumerated type nick for @value and writes it to @key,
// within @settings.
//
// It is a programmer error to give a @key that isn't contained in the
// schema for @settings or is not marked as an enumerated type, or for
// @value not to be a valid value for the named type.
//
// After performing the write, accessing @key directly with
// g_settings_get_string() will return the 'nick' associated with
// @value.
/*

C function : g_settings_set_enum
*/
func (recv *Settings) SetEnum(key string, value int32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	retC := C.g_settings_set_enum((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Looks up the flags type nicks for the bits specified by @value, puts
// them in an array of strings and writes the array to @key, within
// @settings.
//
// It is a programmer error to give a @key that isn't contained in the
// schema for @settings or is not marked as a flags type, or for @value
// to contain any bits that are not value for the named type.
//
// After performing the write, accessing @key directly with
// g_settings_get_strv() will return an array of 'nicks'; one for each
// bit in @value.
/*

C function : g_settings_set_flags
*/
func (recv *Settings) SetFlags(key string, value uint32) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.guint)(value)

	retC := C.g_settings_set_flags((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// The #GSettingsBackend interface defines a generic interface for
// non-strictly-typed data that is stored in a hierarchy. To implement
// an alternative storage backend for #GSettings, you need to implement
// the #GSettingsBackend interface and then make it implement the
// extension point #G_SETTINGS_BACKEND_EXTENSION_POINT_NAME.
//
// The interface defines methods for reading and writing values, a
// method for determining if writing of certain values will fail
// (lockdown) and a change notification mechanism.
//
// The semantics of the interface are very precisely defined and
// implementations must carefully adhere to the expectations of
// callers that are documented on each of the interface methods.
//
// Some of the #GSettingsBackend functions accept or return a #GTree.
// These trees always have strings as keys and #GVariant as values.
// g_settings_backend_create_tree() is a convenience function to create
// suitable trees.
//
// The #GSettingsBackend API is exported to allow third-party
// implementations, but does not carry the same stability guarantees
// as the public GIO API. For this reason, you have to define the
// C preprocessor symbol %G_SETTINGS_ENABLE_BACKEND before including
// `gio/gsettingsbackend.h`.
/*

C record/class : GSettingsBackend
*/
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

	return g
}

func (recv *SettingsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SettingsBackend) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SettingsBackend.
// Exercise care, as this is a potentially dangerous function if the Object is not a SettingsBackend.
func CastToSettingsBackend(object *gobject.Object) *SettingsBackend {
	return SettingsBackendNewFromC(object.ToC())
}

// A #GSimpleAction is the obvious simple implementation of the #GAction
// interface. This is the easiest way to create an action for purposes of
// adding it to a #GSimpleActionGroup.
//
// See also #GtkAction.
/*

C record/class : GSimpleAction
*/
type SimpleAction struct {
	native *C.GSimpleAction
}

func SimpleActionNewFromC(u unsafe.Pointer) *SimpleAction {
	c := (*C.GSimpleAction)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAction{native: c}

	return g
}

func (recv *SimpleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SimpleAction) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SimpleAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleAction.
func CastToSimpleAction(object *gobject.Object) *SimpleAction {
	return SimpleActionNewFromC(object.ToC())
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// As of GLib 2.46, #GSimpleAsyncResult is deprecated in favor of
// #GTask, which provides a simpler API.
//
// #GSimpleAsyncResult implements #GAsyncResult.
//
// GSimpleAsyncResult handles #GAsyncReadyCallbacks, error
// reporting, operation cancellation and the final state of an operation,
// completely transparent to the application. Results can be returned
// as a pointer e.g. for functions that return data that is collected
// asynchronously, a boolean value for checking the success or failure
// of an operation, or a #gssize for operations which return the number
// of bytes modified by the operation; all of the simple return cases
// are covered.
//
// Most of the time, an application will not need to know of the details
// of this API; it is handled transparently, and any necessary operations
// are handled by #GAsyncResult's interface. However, if implementing a
// new GIO module, for writing language bindings, or for complex
// applications that need better control of how asynchronous operations
// are completed, it is important to understand this functionality.
//
// GSimpleAsyncResults are tagged with the calling function to ensure
// that asynchronous functions and their finishing functions are used
// together correctly.
//
// To create a new #GSimpleAsyncResult, call g_simple_async_result_new().
// If the result needs to be created for a #GError, use
// g_simple_async_result_new_from_error() or
// g_simple_async_result_new_take_error(). If a #GError is not available
// (e.g. the asynchronous operation's doesn't take a #GError argument),
// but the result still needs to be created for an error condition, use
// g_simple_async_result_new_error() (or g_simple_async_result_set_error_va()
// if your application or binding requires passing a variable argument list
// directly), and the error can then be propagated through the use of
// g_simple_async_result_propagate_error().
//
// An asynchronous operation can be made to ignore a cancellation event by
// calling g_simple_async_result_set_handle_cancellation() with a
// #GSimpleAsyncResult for the operation and %FALSE. This is useful for
// operations that are dangerous to cancel, such as close (which would
// cause a leak if cancelled before being run).
//
// GSimpleAsyncResult can integrate into GLib's event loop, #GMainLoop,
// or it can use #GThreads.
// g_simple_async_result_complete() will finish an I/O task directly
// from the point where it is called. g_simple_async_result_complete_in_idle()
// will finish it from an idle handler in the
// [thread-default main context][g-main-context-push-thread-default]
// where the #GSimpleAsyncResult was created.
// g_simple_async_result_run_in_thread() will run the job in a
// separate thread and then use
// g_simple_async_result_complete_in_idle() to deliver the result.
//
// To set the results of an asynchronous function,
// g_simple_async_result_set_op_res_gpointer(),
// g_simple_async_result_set_op_res_gboolean(), and
// g_simple_async_result_set_op_res_gssize()
// are provided, setting the operation's result to a gpointer, gboolean, or
// gssize, respectively.
//
// Likewise, to get the result of an asynchronous function,
// g_simple_async_result_get_op_res_gpointer(),
// g_simple_async_result_get_op_res_gboolean(), and
// g_simple_async_result_get_op_res_gssize() are
// provided, getting the operation's result as a gpointer, gboolean, and
// gssize, respectively.
//
// For the details of the requirements implementations must respect, see
// #GAsyncResult.  A typical implementation of an asynchronous operation
// using GSimpleAsyncResult looks something like this:
//
// |[<!-- language="C" -->
// static void
// baked_cb (Cake    *cake,
// gpointer user_data)
// {
// In this example, this callback is not given a reference to the cake,
// so the GSimpleAsyncResult has to take a reference to it.
// GSimpleAsyncResult *result = user_data;
//
// if (cake == NULL)
// g_simple_async_result_set_error (result,
// BAKER_ERRORS,
// BAKER_ERROR_NO_FLOUR,
// "Go to the supermarket");
// else
// g_simple_async_result_set_op_res_gpointer (result,
// g_object_ref (cake),
// g_object_unref);
//
//
// In this example, we assume that baked_cb is called as a callback from
// the mainloop, so it's safe to complete the operation synchronously here.
// If, however, _baker_prepare_cake () might call its callback without
// first returning to the mainloop — inadvisable, but some APIs do so —
// we would need to use g_simple_async_result_complete_in_idle().
// g_simple_async_result_complete (result);
// g_object_unref (result);
// }
//
// void
// baker_bake_cake_async (Baker              *self,
// guint               radius,
// GAsyncReadyCallback callback,
// gpointer            user_data)
// {
// GSimpleAsyncResult *simple;
// Cake               *cake;
//
// if (radius < 3)
// {
// g_simple_async_report_error_in_idle (G_OBJECT (self),
// callback,
// user_data,
// BAKER_ERRORS,
// BAKER_ERROR_TOO_SMALL,
// "%ucm radius cakes are silly",
// radius);
// return;
// }
//
// simple = g_simple_async_result_new (G_OBJECT (self),
// callback,
// user_data,
// baker_bake_cake_async);
// cake = _baker_get_cached_cake (self, radius);
//
// if (cake != NULL)
// {
// g_simple_async_result_set_op_res_gpointer (simple,
// g_object_ref (cake),
// g_object_unref);
// g_simple_async_result_complete_in_idle (simple);
// g_object_unref (simple);
// Drop the reference returned by _baker_get_cached_cake();
// the GSimpleAsyncResult has taken its own reference.
// g_object_unref (cake);
// return;
// }
//
// _baker_prepare_cake (self, radius, baked_cb, simple);
// }
//
// Cake *
// baker_bake_cake_finish (Baker        *self,
// GAsyncResult *result,
// GError      **error)
// {
// GSimpleAsyncResult *simple;
// Cake               *cake;
//
// g_return_val_if_fail (g_simple_async_result_is_valid (result,
// G_OBJECT (self),
// baker_bake_cake_async),
// NULL);
//
// simple = (GSimpleAsyncResult *) result;
//
// if (g_simple_async_result_propagate_error (simple, error))
// return NULL;
//
// cake = CAKE (g_simple_async_result_get_op_res_gpointer (simple));
// return g_object_ref (cake);
// }
// ]|
/*

C record/class : GSimpleAsyncResult
*/
type SimpleAsyncResult struct {
	native *C.GSimpleAsyncResult
}

func SimpleAsyncResultNewFromC(u unsafe.Pointer) *SimpleAsyncResult {
	c := (*C.GSimpleAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResult{native: c}

	return g
}

func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SimpleAsyncResult) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SimpleAsyncResult.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleAsyncResult.
func CastToSimpleAsyncResult(object *gobject.Object) *SimpleAsyncResult {
	return SimpleAsyncResultNewFromC(object.ToC())
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Completes an asynchronous I/O job immediately. Must be called in
// the thread where the asynchronous result was to be delivered, as it
// invokes the callback directly. If you are in a different thread use
// g_simple_async_result_complete_in_idle().
//
// Calling this function takes a reference to @simple for as long as
// is needed to complete the call.
/*

C function : g_simple_async_result_complete
*/
func (recv *SimpleAsyncResult) Complete() {
	C.g_simple_async_result_complete((*C.GSimpleAsyncResult)(recv.native))

	return
}

// Completes an asynchronous function in an idle handler in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread that @simple was initially created in
// (and re-pushes that context around the invocation of the callback).
//
// Calling this function takes a reference to @simple for as long as
// is needed to complete the call.
/*

C function : g_simple_async_result_complete_in_idle
*/
func (recv *SimpleAsyncResult) CompleteInIdle() {
	C.g_simple_async_result_complete_in_idle((*C.GSimpleAsyncResult)(recv.native))

	return
}

// Gets the operation result boolean from within the asynchronous result.
/*

C function : g_simple_async_result_get_op_res_gboolean
*/
func (recv *SimpleAsyncResult) GetOpResGboolean() bool {
	retC := C.g_simple_async_result_get_op_res_gboolean((*C.GSimpleAsyncResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets a pointer result as returned by the asynchronous function.
/*

C function : g_simple_async_result_get_op_res_gpointer
*/
func (recv *SimpleAsyncResult) GetOpResGpointer() uintptr {
	retC := C.g_simple_async_result_get_op_res_gpointer((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Gets a gssize from the asynchronous result.
/*

C function : g_simple_async_result_get_op_res_gssize
*/
func (recv *SimpleAsyncResult) GetOpResGssize() int64 {
	retC := C.g_simple_async_result_get_op_res_gssize((*C.GSimpleAsyncResult)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Gets the source tag for the #GSimpleAsyncResult.
/*

C function : g_simple_async_result_get_source_tag
*/
func (recv *SimpleAsyncResult) GetSourceTag() uintptr {
	retC := C.g_simple_async_result_get_source_tag((*C.GSimpleAsyncResult)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Propagates an error from within the simple asynchronous result to
// a given destination.
//
// If the #GCancellable given to a prior call to
// g_simple_async_result_set_check_cancellable() is cancelled then this
// function will return %TRUE with @dest set appropriately.
/*

C function : g_simple_async_result_propagate_error
*/
func (recv *SimpleAsyncResult) PropagateError() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_simple_async_result_propagate_error((*C.GSimpleAsyncResult)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_simple_async_result_run_in_thread : unsupported parameter func : no type generator for SimpleAsyncThreadFunc (GSimpleAsyncThreadFunc) for param func

// Unsupported : g_simple_async_result_set_error : unsupported parameter ... : varargs

// Unsupported : g_simple_async_result_set_error_va : unsupported parameter args : no type generator for va_list (va_list) for param args

// Sets the result from a #GError.
/*

C function : g_simple_async_result_set_from_error
*/
func (recv *SimpleAsyncResult) SetFromError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_simple_async_result_set_from_error((*C.GSimpleAsyncResult)(recv.native), c_error)

	return
}

// Sets whether to handle cancellation within the asynchronous operation.
//
// This function has nothing to do with
// g_simple_async_result_set_check_cancellable().  It only refers to the
// #GCancellable passed to g_simple_async_result_run_in_thread().
/*

C function : g_simple_async_result_set_handle_cancellation
*/
func (recv *SimpleAsyncResult) SetHandleCancellation(handleCancellation bool) {
	c_handle_cancellation :=
		boolToGboolean(handleCancellation)

	C.g_simple_async_result_set_handle_cancellation((*C.GSimpleAsyncResult)(recv.native), c_handle_cancellation)

	return
}

// Sets the operation result to a boolean within the asynchronous result.
/*

C function : g_simple_async_result_set_op_res_gboolean
*/
func (recv *SimpleAsyncResult) SetOpResGboolean(opRes bool) {
	c_op_res :=
		boolToGboolean(opRes)

	C.g_simple_async_result_set_op_res_gboolean((*C.GSimpleAsyncResult)(recv.native), c_op_res)

	return
}

// Unsupported : g_simple_async_result_set_op_res_gpointer : unsupported parameter destroy_op_res : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy_op_res

// Sets the operation result within the asynchronous result to
// the given @op_res.
/*

C function : g_simple_async_result_set_op_res_gssize
*/
func (recv *SimpleAsyncResult) SetOpResGssize(opRes int64) {
	c_op_res := (C.gssize)(opRes)

	C.g_simple_async_result_set_op_res_gssize((*C.GSimpleAsyncResult)(recv.native), c_op_res)

	return
}

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// #GSimplePermission is a trivial implementation of #GPermission that
// represents a permission that is either always or never allowed.  The
// value is given at construction and doesn't change.
//
// Calling request or release will result in errors.
/*

C record/class : GSimplePermission
*/
type SimplePermission struct {
	native *C.GSimplePermission
}

func SimplePermissionNewFromC(u unsafe.Pointer) *SimplePermission {
	c := (*C.GSimplePermission)(u)
	if c == nil {
		return nil
	}

	g := &SimplePermission{native: c}

	return g
}

func (recv *SimplePermission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Permission upcasts to *Permission
func (recv *SimplePermission) Permission() *Permission {
	return PermissionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SimplePermission) Object() *gobject.Object {
	return recv.Permission().Object()
}

// CastToWidget down casts any arbitary Object to SimplePermission.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimplePermission.
func CastToSimplePermission(object *gobject.Object) *SimplePermission {
	return SimplePermissionNewFromC(object.ToC())
}

// #GSimpleProxyResolver is a simple #GProxyResolver implementation
// that handles a single default proxy, multiple URI-scheme-specific
// proxies, and a list of hosts that proxies should not be used for.
//
// #GSimpleProxyResolver is never the default proxy resolver, but it
// can be used as the base class for another proxy resolver
// implementation, or it can be created and used manually, such as
// with g_socket_client_set_proxy_resolver().
/*

C record/class : GSimpleProxyResolver
*/
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

	return g
}

func (recv *SimpleProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SimpleProxyResolver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SimpleProxyResolver.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleProxyResolver.
func CastToSimpleProxyResolver(object *gobject.Object) *SimpleProxyResolver {
	return SimpleProxyResolverNewFromC(object.ToC())
}

// ProxyResolver returns the ProxyResolver interface implemented by SimpleProxyResolver
func (recv *SimpleProxyResolver) ProxyResolver() *ProxyResolver {
	return ProxyResolverNewFromC(recv.ToC())
}

// #GSocketAddress is the equivalent of struct sockaddr in the BSD
// sockets API. This is an abstract class; use #GInetSocketAddress
// for internet sockets, or #GUnixSocketAddress for UNIX domain sockets.
/*

C record/class : GSocketAddress
*/
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

	return g
}

func (recv *SocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SocketAddress) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketAddress.
func CastToSocketAddress(object *gobject.Object) *SocketAddress {
	return SocketAddressNewFromC(object.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by SocketAddress
func (recv *SocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Enumerator type for objects that contain or generate
// #GSocketAddress instances.
/*

C record/class : GSocketAddressEnumerator
*/
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

	return g
}

func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SocketAddressEnumerator) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SocketAddressEnumerator.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketAddressEnumerator.
func CastToSocketAddressEnumerator(object *gobject.Object) *SocketAddressEnumerator {
	return SocketAddressEnumeratorNewFromC(object.ToC())
}

// Retrieves the next #GSocketAddress from @enumerator. Note that this
// may block for some amount of time. (Eg, a #GNetworkAddress may need
// to do a DNS lookup before it can return an address.) Use
// g_socket_address_enumerator_next_async() if you need to avoid
// blocking.
//
// If @enumerator is expected to yield addresses, but for some reason
// is unable to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error
// in *@error. However, if the first call to
// g_socket_address_enumerator_next() succeeds, then any further
// internal errors (other than @cancellable being triggered) will be
// ignored.
/*

C function : g_socket_address_enumerator_next
*/
func (recv *SocketAddressEnumerator) Next(cancellable *Cancellable) (*SocketAddress, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_address_enumerator_next((*C.GSocketAddressEnumerator)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_address_enumerator_next_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about
// error handling.
/*

C function : g_socket_address_enumerator_next_finish
*/
func (recv *SocketAddressEnumerator) NextFinish(result *AsyncResult) (*SocketAddress, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_address_enumerator_next_finish((*C.GSocketAddressEnumerator)(recv.native), c_result, &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// A #GSocketControlMessage is a special-purpose utility message that
// can be sent to or received from a #GSocket. These types of
// messages are often called "ancillary data".
//
// The message can represent some sort of special instruction to or
// information from the socket or can represent a special kind of
// transfer to the peer (for example, sending a file descriptor over
// a UNIX socket).
//
// These messages are sent with g_socket_send_message() and received
// with g_socket_receive_message().
//
// To extend the set of control message that can be sent, subclass this
// class and override the get_size, get_level, get_type and serialize
// methods.
//
// To extend the set of control messages that can be received, subclass
// this class and implement the deserialize method. Also, make sure your
// class is registered with the GType typesystem before calling
// g_socket_receive_message() to read such a message.
/*

C record/class : GSocketControlMessage
*/
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

	return g
}

func (recv *SocketControlMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SocketControlMessage) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SocketControlMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketControlMessage.
func CastToSocketControlMessage(object *gobject.Object) *SocketControlMessage {
	return SocketControlMessageNewFromC(object.ToC())
}

// A #GTask represents and manages a cancellable "task".
//
// ## Asynchronous operations
//
// The most common usage of #GTask is as a #GAsyncResult, to
// manage data during an asynchronous operation. You call
// g_task_new() in the "start" method, followed by
// g_task_set_task_data() and the like if you need to keep some
// additional data associated with the task, and then pass the
// task object around through your asynchronous operation.
// Eventually, you will call a method such as
// g_task_return_pointer() or g_task_return_error(), which will
// save the value you give it and then invoke the task's callback
// function in the
// [thread-default main context][g-main-context-push-thread-default]
// where it was created (waiting until the next iteration of the main
// loop first, if necessary). The caller will pass the #GTask back to
// the operation's finish function (as a #GAsyncResult), and you can
// can use g_task_propagate_pointer() or the like to extract the
// return value.
//
// Here is an example for using GTask as a GAsyncResult:
// |[<!-- language="C" -->
// typedef struct {
// CakeFrostingType frosting;
// char *message;
// } DecorationData;
//
// static void
// decoration_data_free (DecorationData *decoration)
// {
// g_free (decoration->message);
// g_slice_free (DecorationData, decoration);
// }
//
// static void
// baked_cb (Cake     *cake,
// gpointer  user_data)
// {
// GTask *task = user_data;
// DecorationData *decoration = g_task_get_task_data (task);
// GError *error = NULL;
//
// if (cake == NULL)
// {
// g_task_return_new_error (task, BAKER_ERROR, BAKER_ERROR_NO_FLOUR,
// "Go to the supermarket");
// g_object_unref (task);
// return;
// }
//
// if (!cake_decorate (cake, decoration->frosting, decoration->message, &error))
// {
// g_object_unref (cake);
// g_task_return_error() takes ownership of error
// g_task_return_error (task, error);
// g_object_unref (task);
// return;
// }
//
// g_task_return_pointer (task, cake, g_object_unref);
// g_object_unref (task);
// }
//
// void
// baker_bake_cake_async (Baker               *self,
// guint                radius,
// CakeFlavor           flavor,
// CakeFrostingType     frosting,
// const char          *message,
// GCancellable        *cancellable,
// GAsyncReadyCallback  callback,
// gpointer             user_data)
// {
// GTask *task;
// DecorationData *decoration;
// Cake  *cake;
//
// task = g_task_new (self, cancellable, callback, user_data);
// if (radius < 3)
// {
// g_task_return_new_error (task, BAKER_ERROR, BAKER_ERROR_TOO_SMALL,
// "%ucm radius cakes are silly",
// radius);
// g_object_unref (task);
// return;
// }
//
// cake = _baker_get_cached_cake (self, radius, flavor, frosting, message);
// if (cake != NULL)
// {
// _baker_get_cached_cake() returns a reffed cake
// g_task_return_pointer (task, cake, g_object_unref);
// g_object_unref (task);
// return;
// }
//
// decoration = g_slice_new (DecorationData);
// decoration->frosting = frosting;
// decoration->message = g_strdup (message);
// g_task_set_task_data (task, decoration, (GDestroyNotify) decoration_data_free);
//
// _baker_begin_cake (self, radius, flavor, cancellable, baked_cb, task);
// }
//
// Cake *
// baker_bake_cake_finish (Baker         *self,
// GAsyncResult  *result,
// GError       **error)
// {
// g_return_val_if_fail (g_task_is_valid (result, self), NULL);
//
// return g_task_propagate_pointer (G_TASK (result), error);
// }
// ]|
//
// ## Chained asynchronous operations
//
// #GTask also tries to simplify asynchronous operations that
// internally chain together several smaller asynchronous
// operations. g_task_get_cancellable(), g_task_get_context(),
// and g_task_get_priority() allow you to get back the task's
// #GCancellable, #GMainContext, and [I/O priority][io-priority]
// when starting a new subtask, so you don't have to keep track
// of them yourself. g_task_attach_source() simplifies the case
// of waiting for a source to fire (automatically using the correct
// #GMainContext and priority).
//
// Here is an example for chained asynchronous operations:
// |[<!-- language="C" -->
// typedef struct {
// Cake *cake;
// CakeFrostingType frosting;
// char *message;
// } BakingData;
//
// static void
// decoration_data_free (BakingData *bd)
// {
// if (bd->cake)
// g_object_unref (bd->cake);
// g_free (bd->message);
// g_slice_free (BakingData, bd);
// }
//
// static void
// decorated_cb (Cake         *cake,
// GAsyncResult *result,
// gpointer      user_data)
// {
// GTask *task = user_data;
// GError *error = NULL;
//
// if (!cake_decorate_finish (cake, result, &error))
// {
// g_object_unref (cake);
// g_task_return_error (task, error);
// g_object_unref (task);
// return;
// }
//
// baking_data_free() will drop its ref on the cake, so we have to
// take another here to give to the caller.
// g_task_return_pointer (task, g_object_ref (cake), g_object_unref);
// g_object_unref (task);
// }
//
// static gboolean
// decorator_ready (gpointer user_data)
// {
// GTask *task = user_data;
// BakingData *bd = g_task_get_task_data (task);
//
// cake_decorate_async (bd->cake, bd->frosting, bd->message,
// g_task_get_cancellable (task),
// decorated_cb, task);
//
// return G_SOURCE_REMOVE;
// }
//
// static void
// baked_cb (Cake     *cake,
// gpointer  user_data)
// {
// GTask *task = user_data;
// BakingData *bd = g_task_get_task_data (task);
// GError *error = NULL;
//
// if (cake == NULL)
// {
// g_task_return_new_error (task, BAKER_ERROR, BAKER_ERROR_NO_FLOUR,
// "Go to the supermarket");
// g_object_unref (task);
// return;
// }
//
// bd->cake = cake;
//
// Bail out now if the user has already cancelled
// if (g_task_return_error_if_cancelled (task))
// {
// g_object_unref (task);
// return;
// }
//
// if (cake_decorator_available (cake))
// decorator_ready (task);
// else
// {
// GSource *source;
//
// source = cake_decorator_wait_source_new (cake);
// Attach @source to @task's GMainContext and have it call
// decorator_ready() when it is ready.
// g_task_attach_source (task, source, decorator_ready);
// g_source_unref (source);
// }
// }
//
// void
// baker_bake_cake_async (Baker               *self,
// guint                radius,
// CakeFlavor           flavor,
// CakeFrostingType     frosting,
// const char          *message,
// gint                 priority,
// GCancellable        *cancellable,
// GAsyncReadyCallback  callback,
// gpointer             user_data)
// {
// GTask *task;
// BakingData *bd;
//
// task = g_task_new (self, cancellable, callback, user_data);
// g_task_set_priority (task, priority);
//
// bd = g_slice_new0 (BakingData);
// bd->frosting = frosting;
// bd->message = g_strdup (message);
// g_task_set_task_data (task, bd, (GDestroyNotify) baking_data_free);
//
// _baker_begin_cake (self, radius, flavor, cancellable, baked_cb, task);
// }
//
// Cake *
// baker_bake_cake_finish (Baker         *self,
// GAsyncResult  *result,
// GError       **error)
// {
// g_return_val_if_fail (g_task_is_valid (result, self), NULL);
//
// return g_task_propagate_pointer (G_TASK (result), error);
// }
// ]|
//
// ## Asynchronous operations from synchronous ones
//
// You can use g_task_run_in_thread() to turn a synchronous
// operation into an asynchronous one, by running it in a thread.
// When it completes, the result will be dispatched to the
// [thread-default main context][g-main-context-push-thread-default]
// where the #GTask was created.
//
// Running a task in a thread:
// |[<!-- language="C" -->
// typedef struct {
// guint radius;
// CakeFlavor flavor;
// CakeFrostingType frosting;
// char *message;
// } CakeData;
//
// static void
// cake_data_free (CakeData *cake_data)
// {
// g_free (cake_data->message);
// g_slice_free (CakeData, cake_data);
// }
//
// static void
// bake_cake_thread (GTask         *task,
// gpointer       source_object,
// gpointer       task_data,
// GCancellable  *cancellable)
// {
// Baker *self = source_object;
// CakeData *cake_data = task_data;
// Cake *cake;
// GError *error = NULL;
//
// cake = bake_cake (baker, cake_data->radius, cake_data->flavor,
// cake_data->frosting, cake_data->message,
// cancellable, &error);
// if (cake)
// g_task_return_pointer (task, cake, g_object_unref);
// else
// g_task_return_error (task, error);
// }
//
// void
// baker_bake_cake_async (Baker               *self,
// guint                radius,
// CakeFlavor           flavor,
// CakeFrostingType     frosting,
// const char          *message,
// GCancellable        *cancellable,
// GAsyncReadyCallback  callback,
// gpointer             user_data)
// {
// CakeData *cake_data;
// GTask *task;
//
// cake_data = g_slice_new (CakeData);
// cake_data->radius = radius;
// cake_data->flavor = flavor;
// cake_data->frosting = frosting;
// cake_data->message = g_strdup (message);
// task = g_task_new (self, cancellable, callback, user_data);
// g_task_set_task_data (task, cake_data, (GDestroyNotify) cake_data_free);
// g_task_run_in_thread (task, bake_cake_thread);
// g_object_unref (task);
// }
//
// Cake *
// baker_bake_cake_finish (Baker         *self,
// GAsyncResult  *result,
// GError       **error)
// {
// g_return_val_if_fail (g_task_is_valid (result, self), NULL);
//
// return g_task_propagate_pointer (G_TASK (result), error);
// }
// ]|
//
// ## Adding cancellability to uncancellable tasks
//
// Finally, g_task_run_in_thread() and g_task_run_in_thread_sync()
// can be used to turn an uncancellable operation into a
// cancellable one. If you call g_task_set_return_on_cancel(),
// passing %TRUE, then if the task's #GCancellable is cancelled,
// it will return control back to the caller immediately, while
// allowing the task thread to continue running in the background
// (and simply discarding its result when it finally does finish).
// Provided that the task thread is careful about how it uses
// locks and other externally-visible resources, this allows you
// to make "GLib-friendly" asynchronous and cancellable
// synchronous variants of blocking APIs.
//
// Cancelling a task:
// |[<!-- language="C" -->
// static void
// bake_cake_thread (GTask         *task,
// gpointer       source_object,
// gpointer       task_data,
// GCancellable  *cancellable)
// {
// Baker *self = source_object;
// CakeData *cake_data = task_data;
// Cake *cake;
// GError *error = NULL;
//
// cake = bake_cake (baker, cake_data->radius, cake_data->flavor,
// cake_data->frosting, cake_data->message,
// &error);
// if (error)
// {
// g_task_return_error (task, error);
// return;
// }
//
// If the task has already been cancelled, then we don't want to add
// the cake to the cake cache. Likewise, we don't  want to have the
// task get cancelled in the middle of updating the cache.
// g_task_set_return_on_cancel() will return %TRUE here if it managed
// to disable return-on-cancel, or %FALSE if the task was cancelled
// before it could.
// if (g_task_set_return_on_cancel (task, FALSE))
// {
// If the caller cancels at this point, their
// GAsyncReadyCallback won't be invoked until we return,
// so we don't have to worry that this code will run at
// the same time as that code does. But if there were
// other functions that might look at the cake cache,
// then we'd probably need a GMutex here as well.
// baker_add_cake_to_cache (baker, cake);
// g_task_return_pointer (task, cake, g_object_unref);
// }
// }
//
// void
// baker_bake_cake_async (Baker               *self,
// guint                radius,
// CakeFlavor           flavor,
// CakeFrostingType     frosting,
// const char          *message,
// GCancellable        *cancellable,
// GAsyncReadyCallback  callback,
// gpointer             user_data)
// {
// CakeData *cake_data;
// GTask *task;
//
// cake_data = g_slice_new (CakeData);
//
// ...
//
// task = g_task_new (self, cancellable, callback, user_data);
// g_task_set_task_data (task, cake_data, (GDestroyNotify) cake_data_free);
// g_task_set_return_on_cancel (task, TRUE);
// g_task_run_in_thread (task, bake_cake_thread);
// }
//
// Cake *
// baker_bake_cake_sync (Baker               *self,
// guint                radius,
// CakeFlavor           flavor,
// CakeFrostingType     frosting,
// const char          *message,
// GCancellable        *cancellable,
// GError             **error)
// {
// CakeData *cake_data;
// GTask *task;
// Cake *cake;
//
// cake_data = g_slice_new (CakeData);
//
// ...
//
// task = g_task_new (self, cancellable, NULL, NULL);
// g_task_set_task_data (task, cake_data, (GDestroyNotify) cake_data_free);
// g_task_set_return_on_cancel (task, TRUE);
// g_task_run_in_thread_sync (task, bake_cake_thread);
//
// cake = g_task_propagate_pointer (task, error);
// g_object_unref (task);
// return cake;
// }
// ]|
//
// ## Porting from GSimpleAsyncResult
//
// #GTask's API attempts to be simpler than #GSimpleAsyncResult's
// in several ways:
// - You can save task-specific data with g_task_set_task_data(), and
// retrieve it later with g_task_get_task_data(). This replaces the
// abuse of g_simple_async_result_set_op_res_gpointer() for the same
// purpose with #GSimpleAsyncResult.
// - In addition to the task data, #GTask also keeps track of the
// [priority][io-priority], #GCancellable, and
// #GMainContext associated with the task, so tasks that consist of
// a chain of simpler asynchronous operations will have easy access
// to those values when starting each sub-task.
// - g_task_return_error_if_cancelled() provides simplified
// handling for cancellation. In addition, cancellation
// overrides any other #GTask return value by default, like
// #GSimpleAsyncResult does when
// g_simple_async_result_set_check_cancellable() is called.
// (You can use g_task_set_check_cancellable() to turn off that
// behavior.) On the other hand, g_task_run_in_thread()
// guarantees that it will always run your
// `task_func`, even if the task's #GCancellable
// is already cancelled before the task gets a chance to run;
// you can start your `task_func` with a
// g_task_return_error_if_cancelled() check if you need the
// old behavior.
// - The "return" methods (eg, g_task_return_pointer())
// automatically cause the task to be "completed" as well, and
// there is no need to worry about the "complete" vs "complete
// in idle" distinction. (#GTask automatically figures out
// whether the task's callback can be invoked directly, or
// if it needs to be sent to another #GMainContext, or delayed
// until the next iteration of the current #GMainContext.)
// - The "finish" functions for #GTask based operations are generally
// much simpler than #GSimpleAsyncResult ones, normally consisting
// of only a single call to g_task_propagate_pointer() or the like.
// Since g_task_propagate_pointer() "steals" the return value from
// the #GTask, it is not necessary to juggle pointers around to
// prevent it from being freed twice.
// - With #GSimpleAsyncResult, it was common to call
// g_simple_async_result_propagate_error() from the
// `_finish()` wrapper function, and have
// virtual method implementations only deal with successful
// returns. This behavior is deprecated, because it makes it
// difficult for a subclass to chain to a parent class's async
// methods. Instead, the wrapper function should just be a
// simple wrapper, and the virtual method should call an
// appropriate `g_task_propagate_` function.
// Note that wrapper methods can now use
// g_async_result_legacy_propagate_error() to do old-style
// #GSimpleAsyncResult error-returning behavior, and
// g_async_result_is_tagged() to check if a result is tagged as
// having come from the `_async()` wrapper
// function (for "short-circuit" results, such as when passing
// 0 to g_input_stream_read_async()).
/*

C record/class : GTask
*/
type Task struct {
	native *C.GTask
}

func TaskNewFromC(u unsafe.Pointer) *Task {
	c := (*C.GTask)(u)
	if c == nil {
		return nil
	}

	g := &Task{native: c}

	return g
}

func (recv *Task) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Task) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Task.
// Exercise care, as this is a potentially dangerous function if the Object is not a Task.
func CastToTask(object *gobject.Object) *Task {
	return TaskNewFromC(object.ToC())
}

// AsyncResult returns the AsyncResult interface implemented by Task
func (recv *Task) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// A #GTcpWrapperConnection can be used to wrap a #GIOStream that is
// based on a #GSocket, but which is not actually a
// #GSocketConnection. This is used by #GSocketClient so that it can
// always return a #GSocketConnection, even when the connection it has
// actually created is not directly a #GSocketConnection.
/*

C record/class : GTcpWrapperConnection
*/
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

	return g
}

func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to TcpWrapperConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TcpWrapperConnection.
func CastToTcpWrapperConnection(object *gobject.Object) *TcpWrapperConnection {
	return TcpWrapperConnectionNewFromC(object.ToC())
}

// Get's @conn's base #GIOStream
/*

C function : g_tcp_wrapper_connection_get_base_io_stream
*/
func (recv *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	retC := C.g_tcp_wrapper_connection_get_base_io_stream((*C.GTcpWrapperConnection)(recv.native))
	retGo := IOStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// #GThemedIcon is an implementation of #GIcon that supports icon themes.
// #GThemedIcon contains a list of all of the icons present in an icon
// theme, so that icons can be looked up quickly. #GThemedIcon does
// not provide actual pixmaps for icons, just the icon names.
// Ideally something like gtk_icon_theme_choose_icon() should be used to
// resolve the list of names so that fallback icons work nicely with
// themes that inherit other themes.
/*

C record/class : GThemedIcon
*/
type ThemedIcon struct {
	native *C.GThemedIcon
}

func ThemedIconNewFromC(u unsafe.Pointer) *ThemedIcon {
	c := (*C.GThemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIcon{native: c}

	return g
}

func (recv *ThemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ThemedIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ThemedIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThemedIcon.
func CastToThemedIcon(object *gobject.Object) *ThemedIcon {
	return ThemedIconNewFromC(object.ToC())
}

// Creates a new themed icon for @iconname.
/*

C function : g_themed_icon_new
*/
func ThemedIconNew(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames :

// Creates a new themed icon for @iconname, and all the names
// that can be created by shortening @iconname at '-' characters.
//
// In the following example, @icon1 and @icon2 are equivalent:
// |[<!-- language="C" -->
// const char *names[] = {
// "gnome-dev-cdrom-audio",
// "gnome-dev-cdrom",
// "gnome-dev",
// "gnome"
// };
//
// icon1 = g_themed_icon_new_from_names (names, 4);
// icon2 = g_themed_icon_new_with_default_fallbacks ("gnome-dev-cdrom-audio");
// ]|
/*

C function : g_themed_icon_new_with_default_fallbacks
*/
func ThemedIconNewWithDefaultFallbacks(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new_with_default_fallbacks(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Append a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls
// to g_icon_hash().
/*

C function : g_themed_icon_append_name
*/
func (recv *ThemedIcon) AppendName(iconname string) {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	C.g_themed_icon_append_name((*C.GThemedIcon)(recv.native), c_iconname)

	return
}

// Unsupported : g_themed_icon_get_names : no return type

// Icon returns the Icon interface implemented by ThemedIcon
func (recv *ThemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// This is the subclass of #GSocketConnection that is created
// for UNIX domain sockets.
//
// It contains functions to do some of the UNIX socket specific
// functionality like passing file descriptors.
//
// Note that `<gio/gunixconnection.h>` belongs to the UNIX-specific
// GIO interfaces, thus you have to use the `gio-unix-2.0.pc`
// pkg-config file when using it.
/*

C record/class : GUnixConnection
*/
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

	return g
}

func (recv *UnixConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// CastToWidget down casts any arbitary Object to UnixConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixConnection.
func CastToUnixConnection(object *gobject.Object) *UnixConnection {
	return UnixConnectionNewFromC(object.ToC())
}

// A #GUnixFDList contains a list of file descriptors.  It owns the file
// descriptors that it contains, closing them when finalized.
//
// It may be wrapped in a #GUnixFDMessage and sent over a #GSocket in
// the %G_SOCKET_ADDRESS_UNIX family by using g_socket_send_message()
// and received using g_socket_receive_message().
//
// Note that `<gio/gunixfdlist.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
/*

C record/class : GUnixFDList
*/
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

	return g
}

func (recv *UnixFDList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *UnixFDList) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to UnixFDList.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixFDList.
func CastToUnixFDList(object *gobject.Object) *UnixFDList {
	return UnixFDListNewFromC(object.ToC())
}

// This #GSocketControlMessage contains a #GUnixFDList.
// It may be sent using g_socket_send_message() and received using
// g_socket_receive_message() over UNIX sockets (ie: sockets in the
// %G_SOCKET_ADDRESS_UNIX family). The file descriptors are copied
// between processes by the kernel.
//
// For an easier way to send and receive file descriptors over
// stream-oriented UNIX sockets, see g_unix_connection_send_fd() and
// g_unix_connection_receive_fd().
//
// Note that `<gio/gunixfdmessage.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
/*

C record/class : GUnixFDMessage
*/
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

	return g
}

func (recv *UnixFDMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessage upcasts to *SocketControlMessage
func (recv *UnixFDMessage) SocketControlMessage() *SocketControlMessage {
	return SocketControlMessageNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixFDMessage) Object() *gobject.Object {
	return recv.SocketControlMessage().Object()
}

// CastToWidget down casts any arbitary Object to UnixFDMessage.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixFDMessage.
func CastToUnixFDMessage(object *gobject.Object) *UnixFDMessage {
	return UnixFDMessageNewFromC(object.ToC())
}

// #GUnixInputStream implements #GInputStream for reading from a UNIX
// file descriptor, including asynchronous operations. (If the file
// descriptor refers to a socket or pipe, this will use poll() to do
// asynchronous I/O. If it refers to a regular file, it will fall back
// to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixinputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
/*

C record/class : GUnixInputStream
*/
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

	return g
}

func (recv *UnixInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStream upcasts to *InputStream
func (recv *UnixInputStream) InputStream() *InputStream {
	return InputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixInputStream) Object() *gobject.Object {
	return recv.InputStream().Object()
}

// CastToWidget down casts any arbitary Object to UnixInputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixInputStream.
func CastToUnixInputStream(object *gobject.Object) *UnixInputStream {
	return UnixInputStreamNewFromC(object.ToC())
}

// Creates a new #GUnixInputStream for the given @fd.
//
// If @close_fd is %TRUE, the file descriptor will be closed
// when the stream is closed.
/*

C function : g_unix_input_stream_new
*/
func UnixInputStreamNew(fd int32, closeFd bool) *UnixInputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_input_stream_new(c_fd, c_close_fd)
	retGo := UnixInputStreamNewFromC(unsafe.Pointer(retC))

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

// Watches #GUnixMounts for changes.
/*

C record/class : GUnixMountMonitor
*/
type UnixMountMonitor struct {
	native *C.GUnixMountMonitor
}

func UnixMountMonitorNewFromC(u unsafe.Pointer) *UnixMountMonitor {
	c := (*C.GUnixMountMonitor)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitor{native: c}

	return g
}

func (recv *UnixMountMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *UnixMountMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to UnixMountMonitor.
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
var signalUnixMountMonitorMountpointsChangedLock sync.Mutex

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
var signalUnixMountMonitorMountsChangedLock sync.Mutex

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
	index := int(uintptr(data))
	callback := signalUnixMountMonitorMountsChangedMap[index].callback
	callback()
}

// Deprecated alias for g_unix_mount_monitor_get().
//
// This function was never a true constructor, which is why it was
// renamed.
/*

C function : g_unix_mount_monitor_new
*/
func UnixMountMonitorNew() *UnixMountMonitor {
	retC := C.g_unix_mount_monitor_new()
	retGo := UnixMountMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// #GUnixOutputStream implements #GOutputStream for writing to a UNIX
// file descriptor, including asynchronous operations. (If the file
// descriptor refers to a socket or pipe, this will use poll() to do
// asynchronous I/O. If it refers to a regular file, it will fall back
// to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixoutputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file
// when using it.
/*

C record/class : GUnixOutputStream
*/
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

	return g
}

func (recv *UnixOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStream upcasts to *OutputStream
func (recv *UnixOutputStream) OutputStream() *OutputStream {
	return OutputStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixOutputStream) Object() *gobject.Object {
	return recv.OutputStream().Object()
}

// CastToWidget down casts any arbitary Object to UnixOutputStream.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixOutputStream.
func CastToUnixOutputStream(object *gobject.Object) *UnixOutputStream {
	return UnixOutputStreamNewFromC(object.ToC())
}

// Creates a new #GUnixOutputStream for the given @fd.
//
// If @close_fd, is %TRUE, the file descriptor will be closed when
// the output stream is destroyed.
/*

C function : g_unix_output_stream_new
*/
func UnixOutputStreamNew(fd int32, closeFd bool) *UnixOutputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_output_stream_new(c_fd, c_close_fd)
	retGo := UnixOutputStreamNewFromC(unsafe.Pointer(retC))

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

// Support for UNIX-domain (also known as local) sockets.
//
// UNIX domain sockets are generally visible in the filesystem.
// However, some systems support abstract socket names which are not
// visible in the filesystem and not affected by the filesystem
// permissions, visibility, etc. Currently this is only supported
// under Linux. If you attempt to use abstract sockets on other
// systems, function calls may return %G_IO_ERROR_NOT_SUPPORTED
// errors. You can use g_unix_socket_address_abstract_names_supported()
// to see if abstract names are supported.
//
// Note that `<gio/gunixsocketaddress.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file
// when using it.
/*

C record/class : GUnixSocketAddress
*/
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

	return g
}

func (recv *UnixSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddress upcasts to *SocketAddress
func (recv *UnixSocketAddress) SocketAddress() *SocketAddress {
	return SocketAddressNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *UnixSocketAddress) Object() *gobject.Object {
	return recv.SocketAddress().Object()
}

// CastToWidget down casts any arbitary Object to UnixSocketAddress.
// Exercise care, as this is a potentially dangerous function if the Object is not a UnixSocketAddress.
func CastToUnixSocketAddress(object *gobject.Object) *UnixSocketAddress {
	return UnixSocketAddressNewFromC(object.ToC())
}

// Creates a new %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED
// #GUnixSocketAddress for @path.
/*

C function : g_unix_socket_address_new_abstract
*/
func UnixSocketAddressNewAbstract(path []rune) *UnixSocketAddress {
	c_path := &path[0]

	c_path_len := (C.gint)(len(path))

	retC := C.g_unix_socket_address_new_abstract((*C.gchar)(unsafe.Pointer(c_path)), c_path_len)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Entry point for using GIO functionality.
/*

C record/class : GVfs
*/
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

	return g
}

func (recv *Vfs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Vfs) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Vfs.
// Exercise care, as this is a potentially dangerous function if the Object is not a Vfs.
func CastToVfs(object *gobject.Object) *Vfs {
	return VfsNewFromC(object.ToC())
}

// Gets a #GFile for @path.
/*

C function : g_vfs_get_file_for_path
*/
func (recv *Vfs) GetFileForPath(path string) *File {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_vfs_get_file_for_path((*C.GVfs)(recv.native), c_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a #GFile for @uri.
//
// This operation never fails, but the returned object
// might not support any I/O operation if the URI
// is malformed or if the URI scheme is not supported.
/*

C function : g_vfs_get_file_for_uri
*/
func (recv *Vfs) GetFileForUri(uri string) *File {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_vfs_get_file_for_uri((*C.GVfs)(recv.native), c_uri)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_vfs_get_supported_uri_schemes : no return type

// Checks if the VFS is active.
/*

C function : g_vfs_is_active
*/
func (recv *Vfs) IsActive() bool {
	retC := C.g_vfs_is_active((*C.GVfs)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This operation never fails, but the returned object might
// not support any I/O operations if the @parse_name cannot
// be parsed by the #GVfs module.
/*

C function : g_vfs_parse_name
*/
func (recv *Vfs) ParseName(parseName string) *File {
	c_parse_name := C.CString(parseName)
	defer C.free(unsafe.Pointer(c_parse_name))

	retC := C.g_vfs_parse_name((*C.GVfs)(recv.native), c_parse_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// #GVolumeMonitor is for listing the user interesting devices and volumes
// on the computer. In other words, what a file selector or file manager
// would show in a sidebar.
//
// #GVolumeMonitor is not
// [thread-default-context aware][g-main-context-push-thread-default],
// and so should not be used other than from the main thread, with no
// thread-default-context active.
/*

C record/class : GVolumeMonitor
*/
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

	return g
}

func (recv *VolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *VolumeMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to VolumeMonitor.
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
var signalVolumeMonitorDriveChangedLock sync.Mutex

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
var signalVolumeMonitorDriveConnectedLock sync.Mutex

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
var signalVolumeMonitorDriveDisconnectedLock sync.Mutex

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
var signalVolumeMonitorMountAddedLock sync.Mutex

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
var signalVolumeMonitorMountChangedLock sync.Mutex

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
var signalVolumeMonitorMountPreUnmountLock sync.Mutex

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
var signalVolumeMonitorMountRemovedLock sync.Mutex

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
var signalVolumeMonitorVolumeAddedLock sync.Mutex

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
var signalVolumeMonitorVolumeChangedLock sync.Mutex

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
var signalVolumeMonitorVolumeRemovedLock sync.Mutex

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
	volume := VolumeNewFromC(unsafe.Pointer(c_volume))

	index := int(uintptr(data))
	callback := signalVolumeMonitorVolumeRemovedMap[index].callback
	callback(volume)
}

// Gets a list of drives connected to the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
/*

C function : g_volume_monitor_get_connected_drives
*/
func (recv *VolumeMonitor) GetConnectedDrives() *glib.List {
	retC := C.g_volume_monitor_get_connected_drives((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds a #GMount object by its UUID (see g_mount_get_uuid())
/*

C function : g_volume_monitor_get_mount_for_uuid
*/
func (recv *VolumeMonitor) GetMountForUuid(uuid string) *Mount {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	retC := C.g_volume_monitor_get_mount_for_uuid((*C.GVolumeMonitor)(recv.native), c_uuid)
	retGo := MountNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a list of the mounts on the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
/*

C function : g_volume_monitor_get_mounts
*/
func (recv *VolumeMonitor) GetMounts() *glib.List {
	retC := C.g_volume_monitor_get_mounts((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds a #GVolume object by its UUID (see g_volume_get_uuid())
/*

C function : g_volume_monitor_get_volume_for_uuid
*/
func (recv *VolumeMonitor) GetVolumeForUuid(uuid string) *Volume {
	c_uuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(c_uuid))

	retC := C.g_volume_monitor_get_volume_for_uuid((*C.GVolumeMonitor)(recv.native), c_uuid)
	retGo := VolumeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a list of the volumes on the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
/*

C function : g_volume_monitor_get_volumes
*/
func (recv *VolumeMonitor) GetVolumes() *glib.List {
	retC := C.g_volume_monitor_get_volumes((*C.GVolumeMonitor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Zlib decompression
/*

C record/class : GZlibCompressor
*/
type ZlibCompressor struct {
	native *C.GZlibCompressor
}

func ZlibCompressorNewFromC(u unsafe.Pointer) *ZlibCompressor {
	c := (*C.GZlibCompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressor{native: c}

	return g
}

func (recv *ZlibCompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ZlibCompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ZlibCompressor.
// Exercise care, as this is a potentially dangerous function if the Object is not a ZlibCompressor.
func CastToZlibCompressor(object *gobject.Object) *ZlibCompressor {
	return ZlibCompressorNewFromC(object.ToC())
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Zlib decompression
/*

C record/class : GZlibDecompressor
*/
type ZlibDecompressor struct {
	native *C.GZlibDecompressor
}

func ZlibDecompressorNewFromC(u unsafe.Pointer) *ZlibDecompressor {
	c := (*C.GZlibDecompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressor{native: c}

	return g
}

func (recv *ZlibDecompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ZlibDecompressor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ZlibDecompressor.
// Exercise care, as this is a potentially dangerous function if the Object is not a ZlibDecompressor.
func CastToZlibDecompressor(object *gobject.Object) *ZlibDecompressor {
	return ZlibDecompressorNewFromC(object.ToC())
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}
