// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void application_activateHandler(GObject *, gpointer);

	static gulong Application_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(application_activateHandler), data);
	}

*/
/*

	void application_shutdownHandler(GObject *, gpointer);

	static gulong Application_signal_connect_shutdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "shutdown", G_CALLBACK(application_shutdownHandler), data);
	}

*/
/*

	void application_startupHandler(GObject *, gpointer);

	static gulong Application_signal_connect_startup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "startup", G_CALLBACK(application_startupHandler), data);
	}

*/
import "C"

// A #GApplication is the foundation of an application.  It wraps some
// low-level platform-specific services and is intended to act as the
// foundation for higher-level application classes such as
// #GtkApplication or #MxApplication.  In general, you should not use
// this class outside of a higher level framework.
//
// GApplication provides convenient life cycle management by maintaining
// a "use count" for the primary application instance. The use count can
// be changed using g_application_hold() and g_application_release(). If
// it drops to zero, the application exits. Higher-level classes such as
// #GtkApplication employ the use count to ensure that the application
// stays alive as long as it has any opened windows.
//
// Another feature that GApplication (optionally) provides is process
// uniqueness. Applications can make use of this functionality by
// providing a unique application ID. If given, only one application
// with this ID can be running at a time per session. The session
// concept is platform-dependent, but corresponds roughly to a graphical
// desktop login. When your application is launched again, its
// arguments are passed through platform communication to the already
// running program. The already running instance of the program is
// called the "primary instance"; for non-unique applications this is
// the always the current instance. On Linux, the D-Bus session bus
// is used for communication.
//
// The use of #GApplication differs from some other commonly-used
// uniqueness libraries (such as libunique) in important ways. The
// application is not expected to manually register itself and check
// if it is the primary instance. Instead, the main() function of a
// #GApplication should do very little more than instantiating the
// application instance, possibly connecting signal handlers, then
// calling g_application_run(). All checks for uniqueness are done
// internally. If the application is the primary instance then the
// startup signal is emitted and the mainloop runs. If the application
// is not the primary instance then a signal is sent to the primary
// instance and g_application_run() promptly returns. See the code
// examples below.
//
// If used, the expected form of an application identifier is the same as
// that of of a
// [D-Bus well-known bus name](https://dbus.freedesktop.org/doc/dbus-specification.html#message-protocol-names-bus).
// Examples include: `com.example.MyApp`, `org.example.internal_apps.Calculator`,
// `org._7_zip.Archiver`.
// For details on valid application identifiers, see g_application_id_is_valid().
//
// On Linux, the application identifier is claimed as a well-known bus name
// on the user's session bus.  This means that the uniqueness of your
// application is scoped to the current session.  It also means that your
// application may provide additional services (through registration of other
// object paths) at that bus name.  The registration of these object paths
// should be done with the shared GDBus session bus.  Note that due to the
// internal architecture of GDBus, method calls can be dispatched at any time
// (even if a main loop is not running).  For this reason, you must ensure that
// any object paths that you wish to register are registered before #GApplication
// attempts to acquire the bus name of your application (which happens in
// g_application_register()).  Unfortunately, this means that you cannot use
// g_application_get_is_remote() to decide if you want to register object paths.
//
// GApplication also implements the #GActionGroup and #GActionMap
// interfaces and lets you easily export actions by adding them with
// g_action_map_add_action(). When invoking an action by calling
// g_action_group_activate_action() on the application, it is always
// invoked in the primary instance. The actions are also exported on
// the session bus, and GIO provides the #GDBusActionGroup wrapper to
// conveniently access them remotely. GIO provides a #GDBusMenuModel wrapper
// for remote access to exported #GMenuModels.
//
// There is a number of different entry points into a GApplication:
//
// - via 'Activate' (i.e. just starting the application)
//
// - via 'Open' (i.e. opening some files)
//
// - by handling a command-line
//
// - via activating an action
//
// The #GApplication::startup signal lets you handle the application
// initialization for all of these in a single place.
//
// Regardless of which of these entry points is used to start the
// application, GApplication passes some "platform data from the
// launching instance to the primary instance, in the form of a
// #GVariant dictionary mapping strings to variants. To use platform
// data, override the @before_emit or @after_emit virtual functions
// in your #GApplication subclass. When dealing with
// #GApplicationCommandLine objects, the platform data is
// directly available via g_application_command_line_get_cwd(),
// g_application_command_line_get_environ() and
// g_application_command_line_get_platform_data().
//
// As the name indicates, the platform data may vary depending on the
// operating system, but it always includes the current directory (key
// "cwd"), and optionally the environment (ie the set of environment
// variables and their values) of the calling process (key "environ").
// The environment is only added to the platform data if the
// %G_APPLICATION_SEND_ENVIRONMENT flag is set. #GApplication subclasses
// can add their own platform data by overriding the @add_platform_data
// virtual function. For instance, #GtkApplication adds startup notification
// data in this way.
//
// To parse commandline arguments you may handle the
// #GApplication::command-line signal or override the local_command_line()
// vfunc, to parse them in either the primary instance or the local instance,
// respectively.
//
// For an example of opening files with a GApplication, see
// [gapplication-example-open.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-open.c).
//
// For an example of using actions with GApplication, see
// [gapplication-example-actions.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-actions.c).
//
// For an example of using extra D-Bus hooks with GApplication, see
// [gapplication-example-dbushooks.c](https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-dbushooks.c).
/*

C type

GApplication
*/
type Application struct {
	native *C.GApplication
	// Private : parent_instance
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Application) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Application.
// Exercise care, as this is a potentially dangerous function if the Object is not a Application.
func CastToApplication(object *gobject.Object) *Application {
	return ApplicationNewFromC(object.ToC())
}

type signalApplicationActivateDetail struct {
	callback  ApplicationSignalActivateCallback
	handlerID C.gulong
}

var signalApplicationActivateId int
var signalApplicationActivateMap = make(map[int]signalApplicationActivateDetail)
var signalApplicationActivateLock sync.Mutex

// ApplicationSignalActivateCallback is a callback function for a 'activate' signal emitted from a Application.
type ApplicationSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *Application) ConnectActivate(callback ApplicationSignalActivateCallback) int {
	signalApplicationActivateLock.Lock()
	defer signalApplicationActivateLock.Unlock()

	signalApplicationActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_activate(instance, C.gpointer(uintptr(signalApplicationActivateId)))

	detail := signalApplicationActivateDetail{callback, handlerID}
	signalApplicationActivateMap[signalApplicationActivateId] = detail

	return signalApplicationActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the Application.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *Application) DisconnectActivate(connectionID int) {
	signalApplicationActivateLock.Lock()
	defer signalApplicationActivateLock.Unlock()

	detail, exists := signalApplicationActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationActivateMap, connectionID)
}

//export application_activateHandler
func application_activateHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalApplicationActivateMap[index].callback
	callback()
}

// Unsupported signal 'command-line' for Application : return value gint :

// Unsupported signal 'open' for Application : unsupported parameter files : no param type

type signalApplicationShutdownDetail struct {
	callback  ApplicationSignalShutdownCallback
	handlerID C.gulong
}

var signalApplicationShutdownId int
var signalApplicationShutdownMap = make(map[int]signalApplicationShutdownDetail)
var signalApplicationShutdownLock sync.Mutex

// ApplicationSignalShutdownCallback is a callback function for a 'shutdown' signal emitted from a Application.
type ApplicationSignalShutdownCallback func()

/*
ConnectShutdown connects the callback to the 'shutdown' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectShutdown to remove it.
*/
func (recv *Application) ConnectShutdown(callback ApplicationSignalShutdownCallback) int {
	signalApplicationShutdownLock.Lock()
	defer signalApplicationShutdownLock.Unlock()

	signalApplicationShutdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_shutdown(instance, C.gpointer(uintptr(signalApplicationShutdownId)))

	detail := signalApplicationShutdownDetail{callback, handlerID}
	signalApplicationShutdownMap[signalApplicationShutdownId] = detail

	return signalApplicationShutdownId
}

/*
DisconnectShutdown disconnects a callback from the 'shutdown' signal for the Application.

The connectionID should be a value returned from a call to ConnectShutdown.
*/
func (recv *Application) DisconnectShutdown(connectionID int) {
	signalApplicationShutdownLock.Lock()
	defer signalApplicationShutdownLock.Unlock()

	detail, exists := signalApplicationShutdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationShutdownMap, connectionID)
}

//export application_shutdownHandler
func application_shutdownHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalApplicationShutdownMap[index].callback
	callback()
}

type signalApplicationStartupDetail struct {
	callback  ApplicationSignalStartupCallback
	handlerID C.gulong
}

var signalApplicationStartupId int
var signalApplicationStartupMap = make(map[int]signalApplicationStartupDetail)
var signalApplicationStartupLock sync.Mutex

// ApplicationSignalStartupCallback is a callback function for a 'startup' signal emitted from a Application.
type ApplicationSignalStartupCallback func()

/*
ConnectStartup connects the callback to the 'startup' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectStartup to remove it.
*/
func (recv *Application) ConnectStartup(callback ApplicationSignalStartupCallback) int {
	signalApplicationStartupLock.Lock()
	defer signalApplicationStartupLock.Unlock()

	signalApplicationStartupId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_startup(instance, C.gpointer(uintptr(signalApplicationStartupId)))

	detail := signalApplicationStartupDetail{callback, handlerID}
	signalApplicationStartupMap[signalApplicationStartupId] = detail

	return signalApplicationStartupId
}

/*
DisconnectStartup disconnects a callback from the 'startup' signal for the Application.

The connectionID should be a value returned from a call to ConnectStartup.
*/
func (recv *Application) DisconnectStartup(connectionID int) {
	signalApplicationStartupLock.Lock()
	defer signalApplicationStartupLock.Unlock()

	detail, exists := signalApplicationStartupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationStartupMap, connectionID)
}

//export application_startupHandler
func application_startupHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalApplicationStartupMap[index].callback
	callback()
}

// Creates a new #GApplication instance.
//
// If non-%NULL, the application id must be valid.  See
// g_application_id_is_valid().
//
// If no application ID is given then some features of #GApplication
// (most notably application uniqueness) will be disabled.
/*

C function

g_application_new
*/
func ApplicationNew(applicationId string, flags ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.g_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activates the application.
//
// In essence, this results in the #GApplication::activate signal being
// emitted in the primary instance.
//
// The application must be registered before calling this function.
/*

C function

g_application_activate
*/
func (recv *Application) Activate() {
	C.g_application_activate((*C.GApplication)(recv.native))

	return
}

// Gets the unique identifier for @application.
/*

C function

g_application_get_application_id
*/
func (recv *Application) GetApplicationId() string {
	retC := C.g_application_get_application_id((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the flags for @application.
//
// See #GApplicationFlags.
/*

C function

g_application_get_flags
*/
func (recv *Application) GetFlags() ApplicationFlags {
	retC := C.g_application_get_flags((*C.GApplication)(recv.native))
	retGo := (ApplicationFlags)(retC)

	return retGo
}

// Gets the current inactivity timeout for the application.
//
// This is the amount of time (in milliseconds) after the last call to
// g_application_release() before the application stops running.
/*

C function

g_application_get_inactivity_timeout
*/
func (recv *Application) GetInactivityTimeout() uint32 {
	retC := C.g_application_get_inactivity_timeout((*C.GApplication)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Checks if @application is registered.
//
// An application is registered if g_application_register() has been
// successfully called.
/*

C function

g_application_get_is_registered
*/
func (recv *Application) GetIsRegistered() bool {
	retC := C.g_application_get_is_registered((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @application is remote.
//
// If @application is remote then it means that another instance of
// application already exists (the 'primary' instance).  Calls to
// perform actions on @application will result in the actions being
// performed by the primary instance.
//
// The value of this property cannot be accessed before
// g_application_register() has been called.  See
// g_application_get_is_registered().
/*

C function

g_application_get_is_remote
*/
func (recv *Application) GetIsRemote() bool {
	retC := C.g_application_get_is_remote((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Increases the use count of @application.
//
// Use this function to indicate that the application has a reason to
// continue to run.  For example, g_application_hold() is called by GTK+
// when a toplevel window is on the screen.
//
// To cancel the hold, call g_application_release().
/*

C function

g_application_hold
*/
func (recv *Application) Hold() {
	C.g_application_hold((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_open : unsupported parameter files :

// Attempts registration of the application.
//
// This is the point at which the application discovers if it is the
// primary instance or merely acting as a remote for an already-existing
// primary instance.  This is implemented by attempting to acquire the
// application identifier as a unique bus name on the session bus using
// GDBus.
//
// If there is no application ID or if %G_APPLICATION_NON_UNIQUE was
// given, then this process will always become the primary instance.
//
// Due to the internal architecture of GDBus, method calls can be
// dispatched at any time (even if a main loop is not running).  For
// this reason, you must ensure that any object paths that you wish to
// register are registered before calling this function.
//
// If the application has already been registered then %TRUE is
// returned with no work performed.
//
// The #GApplication::startup signal is emitted if registration succeeds
// and @application is the primary instance (including the non-unique
// case).
//
// In the event of an error (such as @cancellable being cancelled, or a
// failure to connect to the session bus), %FALSE is returned and @error
// is set appropriately.
//
// Note: the return value of this function is not an indicator that this
// instance is or is not the primary instance of the application.  See
// g_application_get_is_remote() for that.
/*

C function

g_application_register
*/
func (recv *Application) Register(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_application_register((*C.GApplication)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Decrease the use count of @application.
//
// When the use count reaches zero, the application will stop running.
//
// Never call this function except to cancel the effect of a previous
// call to g_application_hold().
/*

C function

g_application_release
*/
func (recv *Application) Release() {
	C.g_application_release((*C.GApplication)(recv.native))

	return
}

// Runs the application.
//
// This function is intended to be run from main() and its return value
// is intended to be returned by main(). Although you are expected to pass
// the @argc, @argv parameters from main() to this function, it is possible
// to pass %NULL if @argv is not available or commandline handling is not
// required.  Note that on Windows, @argc and @argv are ignored, and
// g_win32_get_command_line() is called internally (for proper support
// of Unicode commandline arguments).
//
// #GApplication will attempt to parse the commandline arguments.  You
// can add commandline flags to the list of recognised options by way of
// g_application_add_main_option_entries().  After this, the
// #GApplication::handle-local-options signal is emitted, from which the
// application can inspect the values of its #GOptionEntrys.
//
// #GApplication::handle-local-options is a good place to handle options
// such as `--version`, where an immediate reply from the local process is
// desired (instead of communicating with an already-running instance).
// A #GApplication::handle-local-options handler can stop further processing
// by returning a non-negative value, which then becomes the exit status of
// the process.
//
// What happens next depends on the flags: if
// %G_APPLICATION_HANDLES_COMMAND_LINE was specified then the remaining
// commandline arguments are sent to the primary instance, where a
// #GApplication::command-line signal is emitted.  Otherwise, the
// remaining commandline arguments are assumed to be a list of files.
// If there are no files listed, the application is activated via the
// #GApplication::activate signal.  If there are one or more files, and
// %G_APPLICATION_HANDLES_OPEN was specified then the files are opened
// via the #GApplication::open signal.
//
// If you are interested in doing more complicated local handling of the
// commandline then you should implement your own #GApplication subclass
// and override local_command_line(). In this case, you most likely want
// to return %TRUE from your local_command_line() implementation to
// suppress the default handling. See
// [gapplication-example-cmdline2.c][gapplication-example-cmdline2]
// for an example.
//
// If, after the above is done, the use count of the application is zero
// then the exit status is returned immediately.  If the use count is
// non-zero then the default main context is iterated until the use count
// falls to zero, at which point 0 is returned.
//
// If the %G_APPLICATION_IS_SERVICE flag is set, then the service will
// run for as much as 10 seconds with a use count of zero while waiting
// for the message that caused the activation to arrive.  After that,
// if the use count falls to zero the application will exit immediately,
// except in the case that g_application_set_inactivity_timeout() is in
// use.
//
// This function sets the prgname (g_set_prgname()), if not already set,
// to the basename of argv[0].
//
// Much like g_main_loop_run(), this function will acquire the main context
// for the duration that the application is running.
//
// Since 2.40, applications that are not explicitly flagged as services
// or launchers (ie: neither %G_APPLICATION_IS_SERVICE or
// %G_APPLICATION_IS_LAUNCHER are given as flags) will check (from the
// default handler for local_command_line) if "--gapplication-service"
// was given in the command line.  If this flag is present then normal
// commandline processing is interrupted and the
// %G_APPLICATION_IS_SERVICE flag is set.  This provides a "compromise"
// solution whereby running an application directly from the commandline
// will invoke it in the normal way (which can be useful for debugging)
// while still allowing applications to be D-Bus activated in service
// mode.  The D-Bus service file should invoke the executable with
// "--gapplication-service" as the sole commandline argument.  This
// approach is suitable for use by most graphical applications but
// should not be used from applications like editors that need precise
// control over when processes invoked via the commandline will exit and
// what their exit status will be.
/*

C function

g_application_run
*/
func (recv *Application) Run(args []string) int32 {
	cArgc, cArgv := argsIn(args)

	retC := C.g_application_run((*C.GApplication)(recv.native), cArgc, cArgv)
	retGo := (int32)(retC)

	return retGo
}

// This used to be how actions were associated with a #GApplication.
// Now there is #GActionMap for that.
/*

C function

g_application_set_action_group
*/
func (recv *Application) SetActionGroup(actionGroup *ActionGroup) {
	c_action_group := (*C.GActionGroup)(actionGroup.ToC())

	C.g_application_set_action_group((*C.GApplication)(recv.native), c_action_group)

	return
}

// Sets the unique identifier for @application.
//
// The application id can only be modified if @application has not yet
// been registered.
//
// If non-%NULL, the application id must be valid.  See
// g_application_id_is_valid().
/*

C function

g_application_set_application_id
*/
func (recv *Application) SetApplicationId(applicationId string) {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	C.g_application_set_application_id((*C.GApplication)(recv.native), c_application_id)

	return
}

// Sets the flags for @application.
//
// The flags can only be modified if @application has not yet been
// registered.
//
// See #GApplicationFlags.
/*

C function

g_application_set_flags
*/
func (recv *Application) SetFlags(flags ApplicationFlags) {
	c_flags := (C.GApplicationFlags)(flags)

	C.g_application_set_flags((*C.GApplication)(recv.native), c_flags)

	return
}

// Sets the current inactivity timeout for the application.
//
// This is the amount of time (in milliseconds) after the last call to
// g_application_release() before the application stops running.
//
// This call has no side effects of its own.  The value set here is only
// used for next time g_application_release() drops the use count to
// zero.  Any timeouts currently in progress are not impacted.
/*

C function

g_application_set_inactivity_timeout
*/
func (recv *Application) SetInactivityTimeout(inactivityTimeout uint32) {
	c_inactivity_timeout := (C.guint)(inactivityTimeout)

	C.g_application_set_inactivity_timeout((*C.GApplication)(recv.native), c_inactivity_timeout)

	return
}

// Unsupported : g_application_command_line_get_arguments : no return type

// Gets the working directory of the command line invocation.
// The string may contain non-utf8 data.
//
// It is possible that the remote application did not send a working
// directory, so this may be %NULL.
//
// The return value should not be modified or freed and is valid for as
// long as @cmdline exists.
/*

C function

g_application_command_line_get_cwd
*/
func (recv *ApplicationCommandLine) GetCwd() string {
	retC := C.g_application_command_line_get_cwd((*C.GApplicationCommandLine)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_get_environ : no return type

// Gets the exit status of @cmdline.  See
// g_application_command_line_set_exit_status() for more information.
/*

C function

g_application_command_line_get_exit_status
*/
func (recv *ApplicationCommandLine) GetExitStatus() int32 {
	retC := C.g_application_command_line_get_exit_status((*C.GApplicationCommandLine)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Determines if @cmdline represents a remote invocation.
/*

C function

g_application_command_line_get_is_remote
*/
func (recv *ApplicationCommandLine) GetIsRemote() bool {
	retC := C.g_application_command_line_get_is_remote((*C.GApplicationCommandLine)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// Gets the value of a particular environment variable of the command
// line invocation, as would be returned by g_getenv().  The strings may
// contain non-utf8 data.
//
// The remote application usually does not send an environment.  Use
// %G_APPLICATION_SEND_ENVIRONMENT to affect that.  Even with this flag
// set it is possible that the environment is still not available (due
// to invocation messages from other applications).
//
// The return value should not be modified or freed and is valid for as
// long as @cmdline exists.
/*

C function

g_application_command_line_getenv
*/
func (recv *ApplicationCommandLine) Getenv(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_application_command_line_getenv((*C.GApplicationCommandLine)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_print : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_printerr : unsupported parameter ... : varargs

// Sets the exit status that will be used when the invoking process
// exits.
//
// The return value of the #GApplication::command-line signal is
// passed to this function when the handler returns.  This is the usual
// way of setting the exit status.
//
// In the event that you want the remote invocation to continue running
// and want to decide on the exit status in the future, you can use this
// call.  For the case of a remote invocation, the remote process will
// typically exit when the last reference is dropped on @cmdline.  The
// exit status of the remote process will be equal to the last value
// that was set with this function.
//
// In the case that the commandline invocation is local, the situation
// is slightly more complicated.  If the commandline invocation results
// in the mainloop running (ie: because the use-count of the application
// increased to a non-zero value) then the application is considered to
// have been 'successful' in a certain sense, and the exit status is
// always zero.  If the application use count is zero, though, the exit
// status of the local #GApplicationCommandLine is used.
/*

C function

g_application_command_line_set_exit_status
*/
func (recv *ApplicationCommandLine) SetExitStatus(exitStatus int32) {
	c_exit_status := (C.int)(exitStatus)

	C.g_application_command_line_set_exit_status((*C.GApplicationCommandLine)(recv.native), c_exit_status)

	return
}

// Creates a source that triggers if @cancellable is cancelled and
// calls its callback of type #GCancellableSourceFunc. This is
// primarily useful for attaching to another (non-cancellable) source
// with g_source_add_child_source() to add cancellability to it.
//
// For convenience, you can call this with a %NULL #GCancellable,
// in which case the source will never trigger.
//
// The new #GSource will hold a reference to the #GCancellable.
/*

C function

g_cancellable_source_new
*/
func (recv *Cancellable) SourceNew() *glib.Source {
	retC := C.g_cancellable_source_new((*C.GCancellable)(recv.native))
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes all the emblems from @icon.
/*

C function

g_emblemed_icon_clear_emblems
*/
func (recv *EmblemedIcon) ClearEmblems() {
	C.g_emblemed_icon_clear_emblems((*C.GEmblemedIcon)(recv.native))

	return
}

// Unsupported : g_io_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_settings_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported signal 'activate' for SimpleAction : unsupported parameter parameter : type GLib.Variant : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Sets the action as enabled or not.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
//
// This should only be called by the implementor of the action.  Users
// of the action should not attempt to modify its enabled flag.
/*

C function

g_simple_action_set_enabled
*/
func (recv *SimpleAction) SetEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.g_simple_action_set_enabled((*C.GSimpleAction)(recv.native), c_enabled)

	return
}

// #GSimpleActionGroup is a hash table filled with #GAction objects,
// implementing the #GActionGroup and #GActionMap interfaces.
/*

C type

GSimpleActionGroup
*/
type SimpleActionGroup struct {
	native *C.GSimpleActionGroup
	// Private : parent_instance
	// Private : priv
}

func SimpleActionGroupNewFromC(u unsafe.Pointer) *SimpleActionGroup {
	c := (*C.GSimpleActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroup{native: c}

	return g
}

func (recv *SimpleActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SimpleActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SimpleActionGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleActionGroup.
func CastToSimpleActionGroup(object *gobject.Object) *SimpleActionGroup {
	return SimpleActionGroupNewFromC(object.ToC())
}

// Creates a new, empty, #GSimpleActionGroup.
/*

C function

g_simple_action_group_new
*/
func SimpleActionGroupNew() *SimpleActionGroup {
	retC := C.g_simple_action_group_new()
	retGo := SimpleActionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds an action to the action group.
//
// If the action group already contains an action with the same name as
// @action then the old action is dropped from the group.
//
// The action group takes its own reference on @action.
/*

C function

g_simple_action_group_insert
*/
func (recv *SimpleActionGroup) Insert(action *Action) {
	c_action := (*C.GAction)(action.ToC())

	C.g_simple_action_group_insert((*C.GSimpleActionGroup)(recv.native), c_action)

	return
}

// Looks up the action with the name @action_name in the group.
//
// If no such action exists, returns %NULL.
/*

C function

g_simple_action_group_lookup
*/
func (recv *SimpleActionGroup) Lookup(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_simple_action_group_lookup((*C.GSimpleActionGroup)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
/*

C function

g_simple_action_group_remove
*/
func (recv *SimpleActionGroup) Remove(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_simple_action_group_remove((*C.GSimpleActionGroup)(recv.native), c_action_name)

	return
}

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Sets the result from @error, and takes over the caller's ownership
// of @error, so the caller does not need to free it any more.
/*

C function

g_simple_async_result_take_error
*/
func (recv *SimpleAsyncResult) TakeError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_simple_async_result_take_error((*C.GSimpleAsyncResult)(recv.native), c_error)

	return
}

// Gets whether @client creates TLS connections. See
// g_socket_client_set_tls() for details.
/*

C function

g_socket_client_get_tls
*/
func (recv *SocketClient) GetTls() bool {
	retC := C.g_socket_client_get_tls((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the TLS validation flags used creating TLS connections via
// @client.
/*

C function

g_socket_client_get_tls_validation_flags
*/
func (recv *SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	retC := C.g_socket_client_get_tls_validation_flags((*C.GSocketClient)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Sets whether @client creates TLS (aka SSL) connections. If @tls is
// %TRUE, @client will wrap its connections in a #GTlsClientConnection
// and perform a TLS handshake when connecting.
//
// Note that since #GSocketClient must return a #GSocketConnection,
// but #GTlsClientConnection is not a #GSocketConnection, this
// actually wraps the resulting #GTlsClientConnection in a
// #GTcpWrapperConnection when returning it. You can use
// g_tcp_wrapper_connection_get_base_io_stream() on the return value
// to extract the #GTlsClientConnection.
//
// If you need to modify the behavior of the TLS handshake (eg, by
// setting a client-side certificate to use, or connecting to the
// #GTlsConnection::accept-certificate signal), you can connect to
// @client's #GSocketClient::event signal and wait for it to be
// emitted with %G_SOCKET_CLIENT_TLS_HANDSHAKING, which will give you
// a chance to see the #GTlsClientConnection before the handshake
// starts.
/*

C function

g_socket_client_set_tls
*/
func (recv *SocketClient) SetTls(tls bool) {
	c_tls :=
		boolToGboolean(tls)

	C.g_socket_client_set_tls((*C.GSocketClient)(recv.native), c_tls)

	return
}

// Sets the TLS validation flags used when creating TLS connections
// via @client. The default value is %G_TLS_CERTIFICATE_VALIDATE_ALL.
/*

C function

g_socket_client_set_tls_validation_flags
*/
func (recv *SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_socket_client_set_tls_validation_flags((*C.GSocketClient)(recv.native), c_flags)

	return
}

// Wraps @base_io_stream and @socket together as a #GSocketConnection.
/*

C function

g_tcp_wrapper_connection_new
*/
func TcpWrapperConnectionNew(baseIoStream *IOStream, socket *Socket) *TcpWrapperConnection {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_socket := (*C.GSocket)(C.NULL)
	if socket != nil {
		c_socket = (*C.GSocket)(socket.ToC())
	}

	retC := C.g_tcp_wrapper_connection_new(c_base_io_stream, c_socket)
	retGo := TcpWrapperConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// A certificate used for TLS authentication and encryption.
// This can represent either a certificate only (eg, the certificate
// received by a client from a server), or the combination of
// a certificate and a private key (which is needed when acting as a
// #GTlsServerConnection).
/*

C type

GTlsCertificate
*/
type TlsCertificate struct {
	native *C.GTlsCertificate
	// parent_instance : record
	// priv : record
}

func TlsCertificateNewFromC(u unsafe.Pointer) *TlsCertificate {
	c := (*C.GTlsCertificate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificate{native: c}

	return g
}

func (recv *TlsCertificate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TlsCertificate) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TlsCertificate.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsCertificate.
func CastToTlsCertificate(object *gobject.Object) *TlsCertificate {
	return TlsCertificateNewFromC(object.ToC())
}

// Creates a #GTlsCertificate from the PEM-encoded data in @file. The
// returned certificate will be the first certificate found in @file. As
// of GLib 2.44, if @file contains more certificates it will try to load
// a certificate chain. All certificates will be verified in the order
// found (top-level certificate should be the last one in the file) and
// the #GTlsCertificate:issuer property of each certificate will be set
// accordingly if the verification succeeds. If any certificate in the
// chain cannot be verified, the first certificate in the file will
// still be returned.
//
// If @file cannot be read or parsed, the function will return %NULL and
// set @error. Otherwise, this behaves like
// g_tls_certificate_new_from_pem().
/*

C function

g_tls_certificate_new_from_file
*/
func TlsCertificateNewFromFile(file string) (*TlsCertificate, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_file(c_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GTlsCertificate from the PEM-encoded data in @cert_file
// and @key_file. The returned certificate will be the first certificate
// found in @cert_file. As of GLib 2.44, if @cert_file contains more
// certificates it will try to load a certificate chain. All
// certificates will be verified in the order found (top-level
// certificate should be the last one in the file) and the
// #GTlsCertificate:issuer property of each certificate will be set
// accordingly if the verification succeeds. If any certificate in the
// chain cannot be verified, the first certificate in the file will
// still be returned.
//
// If either file cannot be read or parsed, the function will return
// %NULL and set @error. Otherwise, this behaves like
// g_tls_certificate_new_from_pem().
/*

C function

g_tls_certificate_new_from_files
*/
func TlsCertificateNewFromFiles(certFile string, keyFile string) (*TlsCertificate, error) {
	c_cert_file := C.CString(certFile)
	defer C.free(unsafe.Pointer(c_cert_file))

	c_key_file := C.CString(keyFile)
	defer C.free(unsafe.Pointer(c_key_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_files(c_cert_file, c_key_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GTlsCertificate from the PEM-encoded data in @data. If
// @data includes both a certificate and a private key, then the
// returned certificate will include the private key data as well. (See
// the #GTlsCertificate:private-key-pem property for information about
// supported formats.)
//
// The returned certificate will be the first certificate found in
// @data. As of GLib 2.44, if @data contains more certificates it will
// try to load a certificate chain. All certificates will be verified in
// the order found (top-level certificate should be the last one in the
// file) and the #GTlsCertificate:issuer property of each certificate
// will be set accordingly if the verification succeeds. If any
// certificate in the chain cannot be verified, the first certificate in
// the file will still be returned.
/*

C function

g_tls_certificate_new_from_pem
*/
func TlsCertificateNewFromPem(data string) (*TlsCertificate, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gssize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_pem(c_data, c_length, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the #GTlsCertificate representing @cert's issuer, if known
/*

C function

g_tls_certificate_get_issuer
*/
func (recv *TlsCertificate) GetIssuer() *TlsCertificate {
	retC := C.g_tls_certificate_get_issuer((*C.GTlsCertificate)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This verifies @cert and returns a set of #GTlsCertificateFlags
// indicating any problems found with it. This can be used to verify a
// certificate outside the context of making a connection, or to
// check a certificate against a CA that is not part of the system
// CA database.
//
// If @identity is not %NULL, @cert's name(s) will be compared against
// it, and %G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return
// value if it does not match. If @identity is %NULL, that bit will
// never be set in the return value.
//
// If @trusted_ca is not %NULL, then @cert (or one of the certificates
// in its chain) must be signed by it, or else
// %G_TLS_CERTIFICATE_UNKNOWN_CA will be set in the return value. If
// @trusted_ca is %NULL, that bit will never be set in the return
// value.
//
// (All other #GTlsCertificateFlags values will always be set or unset
// as appropriate.)
/*

C function

g_tls_certificate_verify
*/
func (recv *TlsCertificate) Verify(identity *SocketConnectable, trustedCa *TlsCertificate) TlsCertificateFlags {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	c_trusted_ca := (*C.GTlsCertificate)(C.NULL)
	if trustedCa != nil {
		c_trusted_ca = (*C.GTlsCertificate)(trustedCa.ToC())
	}

	retC := C.g_tls_certificate_verify((*C.GTlsCertificate)(recv.native), c_identity, c_trusted_ca)
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// #GTlsConnection is the base TLS connection class type, which wraps
// a #GIOStream and provides TLS encryption on top of it. Its
// subclasses, #GTlsClientConnection and #GTlsServerConnection,
// implement client-side and server-side TLS, respectively.
//
// For DTLS (Datagram TLS) support, see #GDtlsConnection.
/*

C type

GTlsConnection
*/
type TlsConnection struct {
	native *C.GTlsConnection
	// parent_instance : record
	// priv : record
}

func TlsConnectionNewFromC(u unsafe.Pointer) *TlsConnection {
	c := (*C.GTlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnection{native: c}

	return g
}

func (recv *TlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStream upcasts to *IOStream
func (recv *TlsConnection) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *TlsConnection) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitary Object to TlsConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsConnection.
func CastToTlsConnection(object *gobject.Object) *TlsConnection {
	return TlsConnectionNewFromC(object.ToC())
}

// Used by #GTlsConnection implementations to emit the
// #GTlsConnection::accept-certificate signal.
/*

C function

g_tls_connection_emit_accept_certificate
*/
func (recv *TlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	c_peer_cert := (*C.GTlsCertificate)(C.NULL)
	if peerCert != nil {
		c_peer_cert = (*C.GTlsCertificate)(peerCert.ToC())
	}

	c_errors := (C.GTlsCertificateFlags)(errors)

	retC := C.g_tls_connection_emit_accept_certificate((*C.GTlsConnection)(recv.native), c_peer_cert, c_errors)
	retGo := retC == C.TRUE

	return retGo
}

// Gets @conn's certificate, as set by
// g_tls_connection_set_certificate().
/*

C function

g_tls_connection_get_certificate
*/
func (recv *TlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @conn's peer's certificate after the handshake has completed.
// (It is not set during the emission of
// #GTlsConnection::accept-certificate.)
/*

C function

g_tls_connection_get_peer_certificate
*/
func (recv *TlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_peer_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the errors associated with validating @conn's peer's
// certificate, after the handshake has completed. (It is not set
// during the emission of #GTlsConnection::accept-certificate.)
/*

C function

g_tls_connection_get_peer_certificate_errors
*/
func (recv *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_tls_connection_get_peer_certificate_errors((*C.GTlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// Gets @conn rehandshaking mode. See
// g_tls_connection_set_rehandshake_mode() for details.
/*

C function

g_tls_connection_get_rehandshake_mode
*/
func (recv *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_tls_connection_get_rehandshake_mode((*C.GTlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// Tests whether or not @conn expects a proper TLS close notification
// when the connection is closed. See
// g_tls_connection_set_require_close_notify() for details.
/*

C function

g_tls_connection_get_require_close_notify
*/
func (recv *TlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_tls_connection_get_require_close_notify((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether @conn uses the system certificate database to verify
// peer certificates. See g_tls_connection_set_use_system_certdb().
/*

C function

g_tls_connection_get_use_system_certdb
*/
func (recv *TlsConnection) GetUseSystemCertdb() bool {
	retC := C.g_tls_connection_get_use_system_certdb((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method;
// although the connection needs to perform a handshake after
// connecting (or after sending a "STARTTLS"-type command) and may
// need to rehandshake later if the server requests it,
// #GTlsConnection will handle this for you automatically when you try
// to send or receive data on the connection. However, you can call
// g_tls_connection_handshake() manually if you want to know for sure
// whether the initial handshake succeeded or failed (as opposed to
// just immediately trying to write to @conn's output stream, in which
// case if it fails, it may not be possible to tell if it failed
// before or after completing the handshake).
//
// Likewise, on the server side, although a handshake is necessary at
// the beginning of the communication, you do not need to call this
// function explicitly unless you want clearer error reporting.
// However, you may call g_tls_connection_handshake() later on to
// renegotiate parameters (encryption methods, etc) with the client.
//
// #GTlsConnection::accept_certificate may be emitted during the
// handshake.
/*

C function

g_tls_connection_handshake
*/
func (recv *TlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_connection_handshake((*C.GTlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finish an asynchronous TLS handshake operation. See
// g_tls_connection_handshake() for more information.
/*

C function

g_tls_connection_handshake_finish
*/
func (recv *TlsConnection) HandshakeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_connection_handshake_finish((*C.GTlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This sets the certificate that @conn will present to its peer
// during the TLS handshake. For a #GTlsServerConnection, it is
// mandatory to set this, and that will normally be done at construct
// time.
//
// For a #GTlsClientConnection, this is optional. If a handshake fails
// with %G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server
// requires a certificate, and if you try connecting again, you should
// call this method first. You can call
// g_tls_client_connection_get_accepted_cas() on the failed connection
// to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with
// or without a certificate; in that case, if you don't provide a
// certificate, you can tell that the server requested one by the fact
// that g_tls_client_connection_get_accepted_cas() will return
// non-%NULL.)
/*

C function

g_tls_connection_set_certificate
*/
func (recv *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	C.g_tls_connection_set_certificate((*C.GTlsConnection)(recv.native), c_certificate)

	return
}

// Sets how @conn behaves with respect to rehandshaking requests.
//
// %G_TLS_REHANDSHAKE_NEVER means that it will never agree to
// rehandshake after the initial handshake is complete. (For a client,
// this means it will refuse rehandshake requests from the server, and
// for a server, this means it will close the connection with an error
// if the client attempts to rehandshake.)
//
// %G_TLS_REHANDSHAKE_SAFELY means that the connection will allow a
// rehandshake only if the other end of the connection supports the
// TLS `renegotiation_info` extension. This is the default behavior,
// but means that rehandshaking will not work against older
// implementations that do not support that extension.
//
// %G_TLS_REHANDSHAKE_UNSAFELY means that the connection will allow
// rehandshaking even without the `renegotiation_info` extension. On
// the server side in particular, this is not recommended, since it
// leaves the server open to certain attacks. However, this mode is
// necessary if you need to allow renegotiation with older client
// software.
/*

C function

g_tls_connection_set_rehandshake_mode
*/
func (recv *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_tls_connection_set_rehandshake_mode((*C.GTlsConnection)(recv.native), c_mode)

	return
}

// Sets whether or not @conn expects a proper TLS close notification
// before the connection is closed. If this is %TRUE (the default),
// then @conn will expect to receive a TLS close notification from its
// peer before the connection is closed, and will return a
// %G_TLS_ERROR_EOF error if the connection is closed without proper
// notification (since this may indicate a network error, or
// man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the
// connection was closed cleanly based on application-level data
// (because the application-level data includes a length field, or is
// somehow self-delimiting); in this case, the close notify is
// redundant and sometimes omitted. (TLS 1.1 explicitly allows this;
// in TLS 1.0 it is technically an error, but often done anyway.) You
// can use g_tls_connection_set_require_close_notify() to tell @conn
// to allow an "unannounced" connection close, in which case the close
// will show up as a 0-length read, as in a non-TLS
// #GSocketConnection, and it is up to the application to check that
// the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the
// connection; when the application calls g_io_stream_close() itself
// on @conn, this will send a close notification regardless of the
// setting of this property. If you explicitly want to do an unclean
// close, you can close @conn's #GTlsConnection:base-io-stream rather
// than closing @conn itself, but note that this may only be done when no other
// operations are pending on @conn or the base I/O stream.
/*

C function

g_tls_connection_set_require_close_notify
*/
func (recv *TlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_tls_connection_set_require_close_notify((*C.GTlsConnection)(recv.native), c_require_close_notify)

	return
}

// Sets whether @conn uses the system certificate database to verify
// peer certificates. This is %TRUE by default. If set to %FALSE, then
// peer certificate validation will always set the
// %G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// #GTlsConnection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// #GTlsClientConnection:validation-flags).
/*

C function

g_tls_connection_set_use_system_certdb
*/
func (recv *TlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	c_use_system_certdb :=
		boolToGboolean(useSystemCertdb)

	C.g_tls_connection_set_use_system_certdb((*C.GTlsConnection)(recv.native), c_use_system_certdb)

	return
}
