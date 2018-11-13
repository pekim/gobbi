// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void appinfomonitor_changedHandler(GObject *, gpointer);

	static gulong AppInfoMonitor_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(appinfomonitor_changedHandler), data);
	}

*/
import "C"

// #GAppInfoMonitor is a very simple object used for monitoring the app
// info database for changes (ie: newly installed or removed
// applications).
//
// Call g_app_info_monitor_get() to get a #GAppInfoMonitor and connect
// to the "changed" signal.
//
// In the usual case, applications should try to make note of the change
// (doing things like invalidating caches) but not act on it.  In
// particular, applications should avoid making calls to #GAppInfo APIs
// in response to the change signal, deferring these until the time that
// the data is actually required.  The exception to this case is when
// application information is actually being displayed on the screen
// (eg: during a search or when the list of all applications is shown).
// The reason for this is that changes to the list of installed
// applications often come in groups (like during system updates) and
// rescanning the list on every change is pointless and expensive.
/*

C type

GAppInfoMonitor
*/
type AppInfoMonitor struct {
	native *C.GAppInfoMonitor
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	c := (*C.GAppInfoMonitor)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoMonitor{native: c}

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *AppInfoMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to AppInfoMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppInfoMonitor.
func CastToAppInfoMonitor(object *gobject.Object) *AppInfoMonitor {
	return AppInfoMonitorNewFromC(object.ToC())
}

type signalAppInfoMonitorChangedDetail struct {
	callback  AppInfoMonitorSignalChangedCallback
	handlerID C.gulong
}

var signalAppInfoMonitorChangedId int
var signalAppInfoMonitorChangedMap = make(map[int]signalAppInfoMonitorChangedDetail)
var signalAppInfoMonitorChangedLock sync.Mutex

// AppInfoMonitorSignalChangedCallback is a callback function for a 'changed' signal emitted from a AppInfoMonitor.
type AppInfoMonitorSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the AppInfoMonitor.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *AppInfoMonitor) ConnectChanged(callback AppInfoMonitorSignalChangedCallback) int {
	signalAppInfoMonitorChangedLock.Lock()
	defer signalAppInfoMonitorChangedLock.Unlock()

	signalAppInfoMonitorChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AppInfoMonitor_signal_connect_changed(instance, C.gpointer(uintptr(signalAppInfoMonitorChangedId)))

	detail := signalAppInfoMonitorChangedDetail{callback, handlerID}
	signalAppInfoMonitorChangedMap[signalAppInfoMonitorChangedId] = detail

	return signalAppInfoMonitorChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the AppInfoMonitor.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *AppInfoMonitor) DisconnectChanged(connectionID int) {
	signalAppInfoMonitorChangedLock.Lock()
	defer signalAppInfoMonitorChangedLock.Unlock()

	detail, exists := signalAppInfoMonitorChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAppInfoMonitorChangedMap, connectionID)
}

//export appinfomonitor_changedHandler
func appinfomonitor_changedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalAppInfoMonitorChangedMap[index].callback
	callback()
}

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries :

// Adds a #GOptionGroup to the commandline handling of @application.
//
// This function is comparable to g_option_context_add_group().
//
// Unlike g_application_add_main_option_entries(), this function does
// not deal with %NULL @arg_data and never transmits options to the
// primary instance.
//
// The reason for that is because, by the time the options arrive at the
// primary instance, it is typically too late to do anything with them.
// Taking the GTK option group as an example: GTK will already have been
// initialised by the time the #GApplication::command-line handler runs.
// In the case that this is not the first-running instance of the
// application, the existing instance may already have been running for
// a very long time.
//
// This means that the options from #GOptionGroup are only really usable
// in the case that the instance of the application being run is the
// first instance.  Passing options like `--display=` or `--gdk-debug=`
// on future runs will have no effect on the existing primary instance.
//
// Calling this function will cause the options in the supplied option
// group to be parsed, but it does not cause you to be "opted in" to the
// new functionality whereby unrecognised options are rejected even if
// %G_APPLICATION_HANDLES_COMMAND_LINE was given.
/*

C function

g_application_add_option_group
*/
func (recv *Application) AddOptionGroup(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_application_add_option_group((*C.GApplication)(recv.native), c_group)

	return
}

// Sends a notification on behalf of @application to the desktop shell.
// There is no guarantee that the notification is displayed immediately,
// or even at all.
//
// Notifications may persist after the application exits. It will be
// D-Bus-activated when the notification or one of its actions is
// activated.
//
// Modifying @notification after this call has no effect. However, the
// object can be reused for a later call to this function.
//
// @id may be any string that uniquely identifies the event for the
// application. It does not need to be in any special format. For
// example, "new-message" might be appropriate for a notification about
// new messages.
//
// If a previous notification was sent with the same @id, it will be
// replaced with @notification and shown again as if it was a new
// notification. This works even for notifications sent from a previous
// execution of the application, as long as @id is the same string.
//
// @id may be %NULL, but it is impossible to replace or withdraw
// notifications without an id.
//
// If @notification is no longer relevant, it can be withdrawn with
// g_application_withdraw_notification().
/*

C function

g_application_send_notification
*/
func (recv *Application) SendNotification(id string, notification *Notification) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_notification := (*C.GNotification)(C.NULL)
	if notification != nil {
		c_notification = (*C.GNotification)(notification.ToC())
	}

	C.g_application_send_notification((*C.GApplication)(recv.native), c_id, c_notification)

	return
}

// Withdraws a notification that was sent with
// g_application_send_notification().
//
// This call does nothing if a notification with @id doesn't exist or
// the notification was never sent.
//
// This function works even for notifications sent in previous
// executions of this application, as long @id is the same as it was for
// the sent notification.
//
// Note that notifications are dismissed when the user clicks on one
// of the buttons in a notification or triggers its default action, so
// there is no need to explicitly withdraw the notification in that case.
/*

C function

g_application_withdraw_notification
*/
func (recv *Application) WithdrawNotification(id string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	C.g_application_withdraw_notification((*C.GApplication)(recv.native), c_id)

	return
}

// Gets the options there were passed to g_application_command_line().
//
// If you did not override local_command_line() then these are the same
// options that were parsed according to the #GOptionEntrys added to the
// application with g_application_add_main_option_entries() and possibly
// modified from your GApplication::handle-local-options handler.
//
// If no options were sent then an empty dictionary is returned so that
// you don't need to check for %NULL.
/*

C function

g_application_command_line_get_options_dict
*/
func (recv *ApplicationCommandLine) GetOptionsDict() *glib.VariantDict {
	retC := C.g_application_command_line_get_options_dict((*C.GApplicationCommandLine)(recv.native))
	retGo := glib.VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GInetSocketAddress for @address and @port.
//
// If @address is an IPv6 address, it can also contain a scope ID
// (separated from the address by a `%`).
/*

C function

g_inet_socket_address_new_from_string
*/
func InetSocketAddressNewFromString(address string, port uint32) *InetSocketAddress {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_port := (C.guint)(port)

	retC := C.g_inet_socket_address_new_from_string(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// #GNotification is a mechanism for creating a notification to be shown
// to the user -- typically as a pop-up notification presented by the
// desktop environment shell.
//
// The key difference between #GNotification and other similar APIs is
// that, if supported by the desktop environment, notifications sent
// with #GNotification will persist after the application has exited,
// and even across system reboots.
//
// Since the user may click on a notification while the application is
// not running, applications using #GNotification should be able to be
// started as a D-Bus service, using #GApplication.
//
// User interaction with a notification (either the default action, or
// buttons) must be associated with actions on the application (ie:
// "app." actions).  It is not possible to route user interaction
// through the notification itself, because the object will not exist if
// the application is autostarted as a result of a notification being
// clicked.
//
// A notification can be sent with g_application_send_notification().
/*

C type

GNotification
*/
type Notification struct {
	native *C.GNotification
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.GNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Notification) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Notification.
// Exercise care, as this is a potentially dangerous function if the Object is not a Notification.
func CastToNotification(object *gobject.Object) *Notification {
	return NotificationNewFromC(object.ToC())
}

// Creates a new #GNotification with @title as its title.
//
// After populating @notification with more details, it can be sent to
// the desktop shell with g_application_send_notification(). Changing
// any properties after this call will not have any effect until
// resending @notification.
/*

C function

g_notification_new
*/
func NotificationNew(title string) *Notification {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.g_notification_new(c_title)
	retGo := NotificationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a button to @notification that activates the action in
// @detailed_action when clicked. That action must be an
// application-wide action (starting with "app."). If @detailed_action
// contains a target, the action will be activated with that target as
// its parameter.
//
// See g_action_parse_detailed_name() for a description of the format
// for @detailed_action.
/*

C function

g_notification_add_button
*/
func (recv *Notification) AddButton(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_add_button((*C.GNotification)(recv.native), c_label, c_detailed_action)

	return
}

// Unsupported : g_notification_add_button_with_target : unsupported parameter ... : varargs

// Unsupported : g_notification_add_button_with_target_value : unsupported parameter target : Blacklisted record : GVariant

// Sets the body of @notification to @body.
/*

C function

g_notification_set_body
*/
func (recv *Notification) SetBody(body string) {
	c_body := C.CString(body)
	defer C.free(unsafe.Pointer(c_body))

	C.g_notification_set_body((*C.GNotification)(recv.native), c_body)

	return
}

// Sets the default action of @notification to @detailed_action. This
// action is activated when the notification is clicked on.
//
// The action in @detailed_action must be an application-wide action (it
// must start with "app."). If @detailed_action contains a target, the
// given action will be activated with that target as its parameter.
// See g_action_parse_detailed_name() for a description of the format
// for @detailed_action.
//
// When no default action is set, the application that the notification
// was sent on is activated.
/*

C function

g_notification_set_default_action
*/
func (recv *Notification) SetDefaultAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_set_default_action((*C.GNotification)(recv.native), c_detailed_action)

	return
}

// Unsupported : g_notification_set_default_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_notification_set_default_action_and_target_value : unsupported parameter target : Blacklisted record : GVariant

// Sets the icon of @notification to @icon.
/*

C function

g_notification_set_icon
*/
func (recv *Notification) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_notification_set_icon((*C.GNotification)(recv.native), c_icon)

	return
}

// Sets the priority of @notification to @priority. See
// #GNotificationPriority for possible values.
/*

C function

g_notification_set_priority
*/
func (recv *Notification) SetPriority(priority NotificationPriority) {
	c_priority := (C.GNotificationPriority)(priority)

	C.g_notification_set_priority((*C.GNotification)(recv.native), c_priority)

	return
}

// Sets the title of @notification to @title.
/*

C function

g_notification_set_title
*/
func (recv *Notification) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_notification_set_title((*C.GNotification)(recv.native), c_title)

	return
}

// Deprecated in favor of g_notification_set_priority().
/*

C function

g_notification_set_urgent
*/
func (recv *Notification) SetUrgent(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.g_notification_set_urgent((*C.GNotification)(recv.native), c_urgent)

	return
}

// Unsupported : g_output_stream_printf : unsupported parameter ... : varargs

// Unsupported : g_output_stream_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_settings_get_default_value : return type : Blacklisted record : GVariant

// Unsupported : g_settings_get_user_value : return type : Blacklisted record : GVariant

// #GSubprocess allows the creation of and interaction with child
// processes.
//
// Processes can be communicated with using standard GIO-style APIs (ie:
// #GInputStream, #GOutputStream).  There are GIO-style APIs to wait for
// process termination (ie: cancellable and with an asynchronous
// variant).
//
// There is an API to force a process to terminate, as well as a
// race-free API for sending UNIX signals to a subprocess.
//
// One major advantage that GIO brings over the core GLib library is
// comprehensive API for asynchronous I/O, such
// g_output_stream_splice_async().  This makes GSubprocess
// significantly more powerful and flexible than equivalent APIs in
// some other languages such as the `subprocess.py`
// included with Python.  For example, using #GSubprocess one could
// create two child processes, reading standard output from the first,
// processing it, and writing to the input stream of the second, all
// without blocking the main loop.
//
// A powerful g_subprocess_communicate() API is provided similar to the
// `communicate()` method of `subprocess.py`. This enables very easy
// interaction with a subprocess that has been opened with pipes.
//
// #GSubprocess defaults to tight control over the file descriptors open
// in the child process, avoiding dangling-fd issues that are caused by
// a simple fork()/exec().  The only open file descriptors in the
// spawned process are ones that were explicitly specified by the
// #GSubprocess API (unless %G_SUBPROCESS_FLAGS_INHERIT_FDS was
// specified).
//
// #GSubprocess will quickly reap all child processes as they exit,
// avoiding "zombie processes" remaining around for long periods of
// time.  g_subprocess_wait() can be used to wait for this to happen,
// but it will happen even without the call being explicitly made.
//
// As a matter of principle, #GSubprocess has no API that accepts
// shell-style space-separated strings.  It will, however, match the
// typical shell behaviour of searching the PATH for executables that do
// not contain a directory separator in their name.
//
// #GSubprocess attempts to have a very simple API for most uses (ie:
// spawning a subprocess with arguments and support for most typical
// kinds of input and output redirection).  See g_subprocess_new(). The
// #GSubprocessLauncher API is provided for more complicated cases
// (advanced types of redirection, environment variable manipulation,
// change of working directory, child setup functions, etc).
//
// A typical use of #GSubprocess will involve calling
// g_subprocess_new(), followed by g_subprocess_wait_async() or
// g_subprocess_wait().  After the process exits, the status can be
// checked using functions such as g_subprocess_get_if_exited() (which
// are similar to the familiar WIFEXITED-style POSIX macros).
/*

C type

GSubprocess
*/
type Subprocess struct {
	native *C.GSubprocess
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	c := (*C.GSubprocess)(u)
	if c == nil {
		return nil
	}

	g := &Subprocess{native: c}

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Subprocess) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Subprocess.
// Exercise care, as this is a potentially dangerous function if the Object is not a Subprocess.
func CastToSubprocess(object *gobject.Object) *Subprocess {
	return SubprocessNewFromC(object.ToC())
}

// Unsupported : g_subprocess_new : unsupported parameter ... : varargs

// Unsupported : g_subprocess_newv : unsupported parameter argv :

// Communicate with the subprocess until it terminates, and all input
// and output has been completed.
//
// If @stdin_buf is given, the subprocess must have been created with
// %G_SUBPROCESS_FLAGS_STDIN_PIPE.  The given data is fed to the
// stdin of the subprocess and the pipe is closed (ie: EOF).
//
// At the same time (as not to cause blocking when dealing with large
// amounts of data), if %G_SUBPROCESS_FLAGS_STDOUT_PIPE or
// %G_SUBPROCESS_FLAGS_STDERR_PIPE were used, reads from those
// streams.  The data that was read is returned in @stdout and/or
// the @stderr.
//
// If the subprocess was created with %G_SUBPROCESS_FLAGS_STDOUT_PIPE,
// @stdout_buf will contain the data read from stdout.  Otherwise, for
// subprocesses not created with %G_SUBPROCESS_FLAGS_STDOUT_PIPE,
// @stdout_buf will be set to %NULL.  Similar provisions apply to
// @stderr_buf and %G_SUBPROCESS_FLAGS_STDERR_PIPE.
//
// As usual, any output variable may be given as %NULL to ignore it.
//
// If you desire the stdout and stderr data to be interleaved, create
// the subprocess with %G_SUBPROCESS_FLAGS_STDOUT_PIPE and
// %G_SUBPROCESS_FLAGS_STDERR_MERGE.  The merged result will be returned
// in @stdout_buf and @stderr_buf will be set to %NULL.
//
// In case of any error (including cancellation), %FALSE will be
// returned with @error set.  Some or all of the stdin data may have
// been written.  Any stdout or stderr data that has been read will be
// discarded. None of the out variables (aside from @error) will have
// been set to anything in particular and should not be inspected.
//
// In the case that %TRUE is returned, the subprocess has exited and the
// exit status inspection APIs (eg: g_subprocess_get_if_exited(),
// g_subprocess_get_exit_status()) may be used.
//
// You should not attempt to use any of the subprocess pipes after
// starting this function, since they may be left in strange states,
// even if the operation was cancelled.  You should especially not
// attempt to interact with the pipes while the operation is in progress
// (either from another thread or if using the asynchronous version).
/*

C function

g_subprocess_communicate
*/
func (recv *Subprocess) Communicate(stdinBuf *glib.Bytes, cancellable *Cancellable) (bool, *glib.Bytes, *glib.Bytes, error) {
	c_stdin_buf := (*C.GBytes)(C.NULL)
	if stdinBuf != nil {
		c_stdin_buf = (*C.GBytes)(stdinBuf.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_stdout_buf *C.GBytes

	var c_stderr_buf *C.GBytes

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := glib.BytesNewFromC(unsafe.Pointer(c_stdout_buf))

	stderrBuf := glib.BytesNewFromC(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Unsupported : g_subprocess_communicate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Complete an invocation of g_subprocess_communicate_async().
/*

C function

g_subprocess_communicate_finish
*/
func (recv *Subprocess) CommunicateFinish(result *AsyncResult) (bool, *glib.Bytes, *glib.Bytes, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_stdout_buf *C.GBytes

	var c_stderr_buf *C.GBytes

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_finish((*C.GSubprocess)(recv.native), c_result, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := glib.BytesNewFromC(unsafe.Pointer(c_stdout_buf))

	stderrBuf := glib.BytesNewFromC(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Like g_subprocess_communicate(), but validates the output of the
// process as UTF-8, and returns it as a regular NUL terminated string.
/*

C function

g_subprocess_communicate_utf8
*/
func (recv *Subprocess) CommunicateUtf8(stdinBuf string, cancellable *Cancellable) (bool, string, string, error) {
	c_stdin_buf := C.CString(stdinBuf)
	defer C.free(unsafe.Pointer(c_stdin_buf))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_stdout_buf *C.char

	var c_stderr_buf *C.char

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_utf8((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := C.GoString(c_stdout_buf)
	defer C.free(unsafe.Pointer(c_stdout_buf))

	stderrBuf := C.GoString(c_stderr_buf)
	defer C.free(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Unsupported : g_subprocess_communicate_utf8_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Complete an invocation of g_subprocess_communicate_utf8_async().
/*

C function

g_subprocess_communicate_utf8_finish
*/
func (recv *Subprocess) CommunicateUtf8Finish(result *AsyncResult) (bool, string, string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_stdout_buf *C.char

	var c_stderr_buf *C.char

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_utf8_finish((*C.GSubprocess)(recv.native), c_result, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	stdoutBuf := C.GoString(c_stdout_buf)
	defer C.free(unsafe.Pointer(c_stdout_buf))

	stderrBuf := C.GoString(c_stderr_buf)
	defer C.free(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goThrowableError
}

// Use an operating-system specific method to attempt an immediate,
// forceful termination of the process.  There is no mechanism to
// determine whether or not the request itself was successful;
// however, you can use g_subprocess_wait() to monitor the status of
// the process after calling this function.
//
// On Unix, this function sends %SIGKILL.
/*

C function

g_subprocess_force_exit
*/
func (recv *Subprocess) ForceExit() {
	C.g_subprocess_force_exit((*C.GSubprocess)(recv.native))

	return
}

// Check the exit status of the subprocess, given that it exited
// normally.  This is the value passed to the exit() system call or the
// return value from main.
//
// This is equivalent to the system WEXITSTATUS macro.
//
// It is an error to call this function before g_subprocess_wait() and
// unless g_subprocess_get_if_exited() returned %TRUE.
/*

C function

g_subprocess_get_exit_status
*/
func (recv *Subprocess) GetExitStatus() int32 {
	retC := C.g_subprocess_get_exit_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// On UNIX, returns the process ID as a decimal string.
// On Windows, returns the result of GetProcessId() also as a string.
/*

C function

g_subprocess_get_identifier
*/
func (recv *Subprocess) GetIdentifier() string {
	retC := C.g_subprocess_get_identifier((*C.GSubprocess)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Check if the given subprocess exited normally (ie: by way of exit()
// or return from main()).
//
// This is equivalent to the system WIFEXITED macro.
//
// It is an error to call this function before g_subprocess_wait() has
// returned.
/*

C function

g_subprocess_get_if_exited
*/
func (recv *Subprocess) GetIfExited() bool {
	retC := C.g_subprocess_get_if_exited((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Check if the given subprocess terminated in response to a signal.
//
// This is equivalent to the system WIFSIGNALED macro.
//
// It is an error to call this function before g_subprocess_wait() has
// returned.
/*

C function

g_subprocess_get_if_signaled
*/
func (recv *Subprocess) GetIfSignaled() bool {
	retC := C.g_subprocess_get_if_signaled((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the raw status code of the process, as from waitpid().
//
// This value has no particular meaning, but it can be used with the
// macros defined by the system headers such as WIFEXITED.  It can also
// be used with g_spawn_check_exit_status().
//
// It is more likely that you want to use g_subprocess_get_if_exited()
// followed by g_subprocess_get_exit_status().
//
// It is an error to call this function before g_subprocess_wait() has
// returned.
/*

C function

g_subprocess_get_status
*/
func (recv *Subprocess) GetStatus() int32 {
	retC := C.g_subprocess_get_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the #GInputStream from which to read the stderr output of
// @subprocess.
//
// The process must have been created with
// %G_SUBPROCESS_FLAGS_STDERR_PIPE.
/*

C function

g_subprocess_get_stderr_pipe
*/
func (recv *Subprocess) GetStderrPipe() *InputStream {
	retC := C.g_subprocess_get_stderr_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GOutputStream that you can write to in order to give data
// to the stdin of @subprocess.
//
// The process must have been created with
// %G_SUBPROCESS_FLAGS_STDIN_PIPE.
/*

C function

g_subprocess_get_stdin_pipe
*/
func (recv *Subprocess) GetStdinPipe() *OutputStream {
	retC := C.g_subprocess_get_stdin_pipe((*C.GSubprocess)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GInputStream from which to read the stdout output of
// @subprocess.
//
// The process must have been created with
// %G_SUBPROCESS_FLAGS_STDOUT_PIPE.
/*

C function

g_subprocess_get_stdout_pipe
*/
func (recv *Subprocess) GetStdoutPipe() *InputStream {
	retC := C.g_subprocess_get_stdout_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if the process was "successful".  A process is considered
// successful if it exited cleanly with an exit status of 0, either by
// way of the exit() system call or return from main().
//
// It is an error to call this function before g_subprocess_wait() has
// returned.
/*

C function

g_subprocess_get_successful
*/
func (recv *Subprocess) GetSuccessful() bool {
	retC := C.g_subprocess_get_successful((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Get the signal number that caused the subprocess to terminate, given
// that it terminated due to a signal.
//
// This is equivalent to the system WTERMSIG macro.
//
// It is an error to call this function before g_subprocess_wait() and
// unless g_subprocess_get_if_signaled() returned %TRUE.
/*

C function

g_subprocess_get_term_sig
*/
func (recv *Subprocess) GetTermSig() int32 {
	retC := C.g_subprocess_get_term_sig((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sends the UNIX signal @signal_num to the subprocess, if it is still
// running.
//
// This API is race-free.  If the subprocess has terminated, it will not
// be signalled.
//
// This API is not available on Windows.
/*

C function

g_subprocess_send_signal
*/
func (recv *Subprocess) SendSignal(signalNum int32) {
	c_signal_num := (C.gint)(signalNum)

	C.g_subprocess_send_signal((*C.GSubprocess)(recv.native), c_signal_num)

	return
}

// Synchronously wait for the subprocess to terminate.
//
// After the process terminates you can query its exit status with
// functions such as g_subprocess_get_if_exited() and
// g_subprocess_get_exit_status().
//
// This function does not fail in the case of the subprocess having
// abnormal termination.  See g_subprocess_wait_check() for that.
//
// Cancelling @cancellable doesn't kill the subprocess.  Call
// g_subprocess_force_exit() if it is desirable.
/*

C function

g_subprocess_wait
*/
func (recv *Subprocess) Wait(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_subprocess_wait_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Combines g_subprocess_wait() with g_spawn_check_exit_status().
/*

C function

g_subprocess_wait_check
*/
func (recv *Subprocess) WaitCheck(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_check((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_subprocess_wait_check_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Collects the result of a previous call to
// g_subprocess_wait_check_async().
/*

C function

g_subprocess_wait_check_finish
*/
func (recv *Subprocess) WaitCheckFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_check_finish((*C.GSubprocess)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Collects the result of a previous call to
// g_subprocess_wait_async().
/*

C function

g_subprocess_wait_finish
*/
func (recv *Subprocess) WaitFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_finish((*C.GSubprocess)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This class contains a set of options for launching child processes,
// such as where its standard input and output will be directed, the
// argument list, the environment, and more.
//
// While the #GSubprocess class has high level functions covering
// popular cases, use of this class allows access to more advanced
// options.  It can also be used to launch multiple subprocesses with
// a similar configuration.
/*

C type

GSubprocessLauncher
*/
type SubprocessLauncher struct {
	native *C.GSubprocessLauncher
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	c := (*C.GSubprocessLauncher)(u)
	if c == nil {
		return nil
	}

	g := &SubprocessLauncher{native: c}

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SubprocessLauncher) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SubprocessLauncher.
// Exercise care, as this is a potentially dangerous function if the Object is not a SubprocessLauncher.
func CastToSubprocessLauncher(object *gobject.Object) *SubprocessLauncher {
	return SubprocessLauncherNewFromC(object.ToC())
}

// Creates a new #GSubprocessLauncher.
//
// The launcher is created with the default options.  A copy of the
// environment of the calling process is made at the time of this call
// and will be used as the environment that the process is launched in.
/*

C function

g_subprocess_launcher_new
*/
func SubprocessLauncherNew(flags SubprocessFlags) *SubprocessLauncher {
	c_flags := (C.GSubprocessFlags)(flags)

	retC := C.g_subprocess_launcher_new(c_flags)
	retGo := SubprocessLauncherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the value of the environment variable @variable in the
// environment of processes launched from this launcher.
//
// On UNIX, the returned string can be an arbitrary byte string.
// On Windows, it will be UTF-8.
/*

C function

g_subprocess_launcher_getenv
*/
func (recv *SubprocessLauncher) Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_subprocess_launcher_getenv((*C.GSubprocessLauncher)(recv.native), c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_subprocess_launcher_set_child_setup : unsupported parameter child_setup : no type generator for GLib.SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Sets the current working directory that processes will be launched
// with.
//
// By default processes are launched with the current working directory
// of the launching process at the time of launch.
/*

C function

g_subprocess_launcher_set_cwd
*/
func (recv *SubprocessLauncher) SetCwd(cwd string) {
	c_cwd := C.CString(cwd)
	defer C.free(unsafe.Pointer(c_cwd))

	C.g_subprocess_launcher_set_cwd((*C.GSubprocessLauncher)(recv.native), c_cwd)

	return
}

// Unsupported : g_subprocess_launcher_set_environ : unsupported parameter env :

// Sets the flags on the launcher.
//
// The default flags are %G_SUBPROCESS_FLAGS_NONE.
//
// You may not set flags that specify conflicting options for how to
// handle a particular stdio stream (eg: specifying both
// %G_SUBPROCESS_FLAGS_STDIN_PIPE and
// %G_SUBPROCESS_FLAGS_STDIN_INHERIT).
//
// You may also not set a flag that conflicts with a previous call to a
// function like g_subprocess_launcher_set_stdin_file_path() or
// g_subprocess_launcher_take_stdout_fd().
/*

C function

g_subprocess_launcher_set_flags
*/
func (recv *SubprocessLauncher) SetFlags(flags SubprocessFlags) {
	c_flags := (C.GSubprocessFlags)(flags)

	C.g_subprocess_launcher_set_flags((*C.GSubprocessLauncher)(recv.native), c_flags)

	return
}

// Sets the file path to use as the stderr for spawned processes.
//
// If @path is %NULL then any previously given path is unset.
//
// The file will be created or truncated when the process is spawned, as
// would be the case if using '2>' at the shell.
//
// If you want to send both stdout and stderr to the same file then use
// %G_SUBPROCESS_FLAGS_STDERR_MERGE.
//
// You may not set a stderr file path if a stderr fd is already set or
// if the launcher flags contain any flags directing stderr elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_set_stderr_file_path
*/
func (recv *SubprocessLauncher) SetStderrFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stderr_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// Sets the file path to use as the stdin for spawned processes.
//
// If @path is %NULL then any previously given path is unset.
//
// The file must exist or spawning the process will fail.
//
// You may not set a stdin file path if a stdin fd is already set or if
// the launcher flags contain any flags directing stdin elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_set_stdin_file_path
*/
func (recv *SubprocessLauncher) SetStdinFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdin_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// Sets the file path to use as the stdout for spawned processes.
//
// If @path is %NULL then any previously given path is unset.
//
// The file will be created or truncated when the process is spawned, as
// would be the case if using '>' at the shell.
//
// You may not set a stdout file path if a stdout fd is already set or
// if the launcher flags contain any flags directing stdout elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_set_stdout_file_path
*/
func (recv *SubprocessLauncher) SetStdoutFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdout_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// Sets the environment variable @variable in the environment of
// processes launched from this launcher.
//
// On UNIX, both the variable's name and value can be arbitrary byte
// strings, except that the variable's name cannot contain '='.
// On Windows, they should be in UTF-8.
/*

C function

g_subprocess_launcher_setenv
*/
func (recv *SubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	C.g_subprocess_launcher_setenv((*C.GSubprocessLauncher)(recv.native), c_variable, c_value, c_overwrite)

	return
}

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter ... : varargs

// Unsupported : g_subprocess_launcher_spawnv : unsupported parameter argv :

// Transfer an arbitrary file descriptor from parent process to the
// child.  This function takes "ownership" of the fd; it will be closed
// in the parent when @self is freed.
//
// By default, all file descriptors from the parent will be closed.
// This function allows you to create (for example) a custom pipe() or
// socketpair() before launching the process, and choose the target
// descriptor in the child.
//
// An example use case is GNUPG, which has a command line argument
// --passphrase-fd providing a file descriptor number where it expects
// the passphrase to be written.
/*

C function

g_subprocess_launcher_take_fd
*/
func (recv *SubprocessLauncher) TakeFd(sourceFd int32, targetFd int32) {
	c_source_fd := (C.gint)(sourceFd)

	c_target_fd := (C.gint)(targetFd)

	C.g_subprocess_launcher_take_fd((*C.GSubprocessLauncher)(recv.native), c_source_fd, c_target_fd)

	return
}

// Sets the file descriptor to use as the stderr for spawned processes.
//
// If @fd is -1 then any previously given fd is unset.
//
// Note that the default behaviour is to pass stderr through to the
// stderr of the parent process.
//
// The passed @fd belongs to the #GSubprocessLauncher.  It will be
// automatically closed when the launcher is finalized.  The file
// descriptor will also be closed on the child side when executing the
// spawned process.
//
// You may not set a stderr fd if a stderr file path is already set or
// if the launcher flags contain any flags directing stderr elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_take_stderr_fd
*/
func (recv *SubprocessLauncher) TakeStderrFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stderr_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// Sets the file descriptor to use as the stdin for spawned processes.
//
// If @fd is -1 then any previously given fd is unset.
//
// Note that if your intention is to have the stdin of the calling
// process inherited by the child then %G_SUBPROCESS_FLAGS_STDIN_INHERIT
// is a better way to go about doing that.
//
// The passed @fd is noted but will not be touched in the current
// process.  It is therefore necessary that it be kept open by the
// caller until the subprocess is spawned.  The file descriptor will
// also not be explicitly closed on the child side, so it must be marked
// O_CLOEXEC if that's what you want.
//
// You may not set a stdin fd if a stdin file path is already set or if
// the launcher flags contain any flags directing stdin elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_take_stdin_fd
*/
func (recv *SubprocessLauncher) TakeStdinFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdin_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// Sets the file descriptor to use as the stdout for spawned processes.
//
// If @fd is -1 then any previously given fd is unset.
//
// Note that the default behaviour is to pass stdout through to the
// stdout of the parent process.
//
// The passed @fd is noted but will not be touched in the current
// process.  It is therefore necessary that it be kept open by the
// caller until the subprocess is spawned.  The file descriptor will
// also not be explicitly closed on the child side, so it must be marked
// O_CLOEXEC if that's what you want.
//
// You may not set a stdout fd if a stdout file path is already set or
// if the launcher flags contain any flags directing stdout elsewhere.
//
// This feature is only available on UNIX.
/*

C function

g_subprocess_launcher_take_stdout_fd
*/
func (recv *SubprocessLauncher) TakeStdoutFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdout_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// Removes the environment variable @variable from the environment of
// processes launched from this launcher.
//
// On UNIX, the variable's name can be an arbitrary byte string not
// containing '='. On Windows, it should be in UTF-8.
/*

C function

g_subprocess_launcher_unsetenv
*/
func (recv *SubprocessLauncher) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_subprocess_launcher_unsetenv((*C.GSubprocessLauncher)(recv.native), c_variable)

	return
}

// Invoke the interaction to ask the user to choose a certificate to
// use with the connection. It invokes this interaction in the main
// loop, specifically the #GMainContext returned by
// g_main_context_get_thread_default() when the interaction is
// created. This is called by called by #GTlsConnection when the peer
// requests a certificate during the handshake.
//
// Derived subclasses usually implement a certificate selector,
// although they may also choose to provide a certificate from
// elsewhere. Alternatively the user may abort this certificate
// request, which may or may not abort the TLS connection.
//
// The implementation can either be a synchronous (eg: modal dialog) or an
// asynchronous one (eg: modeless dialog). This function will take care of
// calling which ever one correctly.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
// not support immediate cancellation.
/*

C function

g_tls_interaction_invoke_request_certificate
*/
func (recv *TlsInteraction) InvokeRequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GTlsConnection)(connection.ToC())
	}

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_invoke_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Run synchronous interaction to ask the user to choose a certificate to use
// with the connection. In general, g_tls_interaction_invoke_request_certificate()
// should be used instead of this function.
//
// Derived subclasses usually implement a certificate selector, although they may
// also choose to provide a certificate from elsewhere. Alternatively the user may
// abort this certificate request, which will usually abort the TLS connection.
//
// If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsConnection
// passed to g_tls_interaction_request_certificate() will have had its
// #GTlsConnection:certificate filled in.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code. Certain implementations may
// not support immediate cancellation.
/*

C function

g_tls_interaction_request_certificate
*/
func (recv *TlsInteraction) RequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GTlsConnection)(connection.ToC())
	}

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Complete an request certificate user interaction request. This should be once
// the g_tls_interaction_request_certificate_async() completion callback is called.
//
// If %G_TLS_INTERACTION_HANDLED is returned, then the #GTlsConnection
// passed to g_tls_interaction_request_certificate_async() will have had its
// #GTlsConnection:certificate filled in.
//
// If the interaction is cancelled by the cancellation object, or by the
// user then %G_TLS_INTERACTION_FAILED will be returned with an error that
// contains a %G_IO_ERROR_CANCELLED error code.
/*

C function

g_tls_interaction_request_certificate_finish
*/
func (recv *TlsInteraction) RequestCertificateFinish(result *AsyncResult) (TlsInteractionResult, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_request_certificate_finish((*C.GTlsInteraction)(recv.native), c_result, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
