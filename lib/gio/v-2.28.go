// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	void application_activateHandler(GObject *, gpointer);

	static gulong Application_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(application_activateHandler), data);
	}

*/
/*

	gint application_commandLineHandler(GObject *, GApplicationCommandLine *, gpointer);

	static gulong Application_signal_connect_command_line(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "command-line", G_CALLBACK(application_commandLineHandler), data);
	}

*/
/*

	void application_openHandler(GObject *, gpointer, gint, gchar*, gpointer);

	static gulong Application_signal_connect_open(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "open", G_CALLBACK(application_openHandler), data);
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
/*

	static void _g_application_command_line_print(GApplicationCommandLine* cmdline, const gchar* format) {
		return g_application_command_line_print(cmdline, format);
    }
*/
/*

	static void _g_application_command_line_printerr(GApplicationCommandLine* cmdline, const gchar* format) {
		return g_application_command_line_printerr(cmdline, format);
    }
*/
/*

	void simpleaction_activateHandler(GObject *, GVariant *, gpointer);

	static gulong SimpleAction_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(simpleaction_activateHandler), data);
	}

*/
/*

	gboolean tlsconnection_acceptCertificateHandler(GObject *, GTlsCertificate *, GTlsCertificateFlags, gpointer);

	static gulong TlsConnection_signal_connect_accept_certificate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accept-certificate", G_CALLBACK(tlsconnection_acceptCertificateHandler), data);
	}

*/
/*

	void actiongroup_actionAddedHandler(GObject *, gchar*, gpointer);

	static gulong ActionGroup_signal_connect_action_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-added", G_CALLBACK(actiongroup_actionAddedHandler), data);
	}

*/
/*

	void actiongroup_actionEnabledChangedHandler(GObject *, gchar*, gboolean, gpointer);

	static gulong ActionGroup_signal_connect_action_enabled_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-enabled-changed", G_CALLBACK(actiongroup_actionEnabledChangedHandler), data);
	}

*/
/*

	void actiongroup_actionRemovedHandler(GObject *, gchar*, gpointer);

	static gulong ActionGroup_signal_connect_action_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-removed", G_CALLBACK(actiongroup_actionRemovedHandler), data);
	}

*/
/*

	void actiongroup_actionStateChangedHandler(GObject *, gchar*, GVariant *, gpointer);

	static gulong ActionGroup_signal_connect_action_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-state-changed", G_CALLBACK(actiongroup_actionStateChangedHandler), data);
	}

*/
import "C"

type ApplicationFlags C.GApplicationFlags

const (
	APPLICATION_FLAGS_NONE           ApplicationFlags = 0
	APPLICATION_IS_SERVICE           ApplicationFlags = 1
	APPLICATION_IS_LAUNCHER          ApplicationFlags = 2
	APPLICATION_HANDLES_OPEN         ApplicationFlags = 4
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = 8
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags = 16
	APPLICATION_NON_UNIQUE           ApplicationFlags = 32
	APPLICATION_CAN_OVERRIDE_APP_ID  ApplicationFlags = 64
)

type IOStreamSpliceFlags C.GIOStreamSpliceFlags

const (
	IO_STREAM_SPLICE_NONE          IOStreamSpliceFlags = 0
	IO_STREAM_SPLICE_CLOSE_STREAM1 IOStreamSpliceFlags = 1
	IO_STREAM_SPLICE_CLOSE_STREAM2 IOStreamSpliceFlags = 2
	IO_STREAM_SPLICE_WAIT_FOR_BOTH IOStreamSpliceFlags = 4
)

type TlsCertificateFlags C.GTlsCertificateFlags

const (
	TLS_CERTIFICATE_UNKNOWN_CA    TlsCertificateFlags = 1
	TLS_CERTIFICATE_BAD_IDENTITY  TlsCertificateFlags = 2
	TLS_CERTIFICATE_NOT_ACTIVATED TlsCertificateFlags = 4
	TLS_CERTIFICATE_EXPIRED       TlsCertificateFlags = 8
	TLS_CERTIFICATE_REVOKED       TlsCertificateFlags = 16
	TLS_CERTIFICATE_INSECURE      TlsCertificateFlags = 32
	TLS_CERTIFICATE_GENERIC_ERROR TlsCertificateFlags = 64
	TLS_CERTIFICATE_VALIDATE_ALL  TlsCertificateFlags = 127
)

// Application is a wrapper around the C record GApplication.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Application) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Application with another Application, and returns true if they represent the same GObject.
func (recv *Application) Equals(other *Application) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Application) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Application.
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
var signalApplicationActivateLock sync.RWMutex

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
	signalApplicationActivateLock.RLock()
	defer signalApplicationActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalApplicationActivateMap[index].callback
	callback()
}

type signalApplicationCommandLineDetail struct {
	callback  ApplicationSignalCommandLineCallback
	handlerID C.gulong
}

var signalApplicationCommandLineId int
var signalApplicationCommandLineMap = make(map[int]signalApplicationCommandLineDetail)
var signalApplicationCommandLineLock sync.RWMutex

// ApplicationSignalCommandLineCallback is a callback function for a 'command-line' signal emitted from a Application.
type ApplicationSignalCommandLineCallback func(commandLine *ApplicationCommandLine) int32

/*
ConnectCommandLine connects the callback to the 'command-line' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectCommandLine to remove it.
*/
func (recv *Application) ConnectCommandLine(callback ApplicationSignalCommandLineCallback) int {
	signalApplicationCommandLineLock.Lock()
	defer signalApplicationCommandLineLock.Unlock()

	signalApplicationCommandLineId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_command_line(instance, C.gpointer(uintptr(signalApplicationCommandLineId)))

	detail := signalApplicationCommandLineDetail{callback, handlerID}
	signalApplicationCommandLineMap[signalApplicationCommandLineId] = detail

	return signalApplicationCommandLineId
}

/*
DisconnectCommandLine disconnects a callback from the 'command-line' signal for the Application.

The connectionID should be a value returned from a call to ConnectCommandLine.
*/
func (recv *Application) DisconnectCommandLine(connectionID int) {
	signalApplicationCommandLineLock.Lock()
	defer signalApplicationCommandLineLock.Unlock()

	detail, exists := signalApplicationCommandLineMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationCommandLineMap, connectionID)
}

//export application_commandLineHandler
func application_commandLineHandler(_ *C.GObject, c_command_line *C.GApplicationCommandLine, data C.gpointer) C.gint {
	signalApplicationCommandLineLock.RLock()
	defer signalApplicationCommandLineLock.RUnlock()

	commandLine := ApplicationCommandLineNewFromC(unsafe.Pointer(c_command_line))

	index := int(uintptr(data))
	callback := signalApplicationCommandLineMap[index].callback
	retGo := callback(commandLine)
	retC :=
		(C.gint)(retGo)
	return retC
}

type signalApplicationOpenDetail struct {
	callback  ApplicationSignalOpenCallback
	handlerID C.gulong
}

var signalApplicationOpenId int
var signalApplicationOpenMap = make(map[int]signalApplicationOpenDetail)
var signalApplicationOpenLock sync.RWMutex

// ApplicationSignalOpenCallback is a callback function for a 'open' signal emitted from a Application.
type ApplicationSignalOpenCallback func(files []*File, hint string)

/*
ConnectOpen connects the callback to the 'open' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectOpen to remove it.
*/
func (recv *Application) ConnectOpen(callback ApplicationSignalOpenCallback) int {
	signalApplicationOpenLock.Lock()
	defer signalApplicationOpenLock.Unlock()

	signalApplicationOpenId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_open(instance, C.gpointer(uintptr(signalApplicationOpenId)))

	detail := signalApplicationOpenDetail{callback, handlerID}
	signalApplicationOpenMap[signalApplicationOpenId] = detail

	return signalApplicationOpenId
}

/*
DisconnectOpen disconnects a callback from the 'open' signal for the Application.

The connectionID should be a value returned from a call to ConnectOpen.
*/
func (recv *Application) DisconnectOpen(connectionID int) {
	signalApplicationOpenLock.Lock()
	defer signalApplicationOpenLock.Unlock()

	detail, exists := signalApplicationOpenMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationOpenMap, connectionID)
}

//export application_openHandler
func application_openHandler(_ *C.GObject, c_files C.gpointer, c_n_files C.gint, c_hint *C.gchar, data C.gpointer) {
	signalApplicationOpenLock.RLock()
	defer signalApplicationOpenLock.RUnlock()

	files := make([]*File, int(c_n_files), int(c_n_files))
	for i := 0; i < int(c_n_files); i++ {
		_item := FileNewFromC(unsafe.Pointer(*(*C.gpointer)(c_files)))
		files[i] = _item
		c_files = C.gpointer(uintptr(c_files) + uintptr(C.sizeof_gpointer))
	}

	hint := C.GoString(c_hint)

	index := int(uintptr(data))
	callback := signalApplicationOpenMap[index].callback
	callback(files, hint)
}

type signalApplicationShutdownDetail struct {
	callback  ApplicationSignalShutdownCallback
	handlerID C.gulong
}

var signalApplicationShutdownId int
var signalApplicationShutdownMap = make(map[int]signalApplicationShutdownDetail)
var signalApplicationShutdownLock sync.RWMutex

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
	signalApplicationShutdownLock.RLock()
	defer signalApplicationShutdownLock.RUnlock()

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
var signalApplicationStartupLock sync.RWMutex

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
	signalApplicationStartupLock.RLock()
	defer signalApplicationStartupLock.RUnlock()

	index := int(uintptr(data))
	callback := signalApplicationStartupMap[index].callback
	callback()
}

// ApplicationNew is a wrapper around the C function g_application_new.
func ApplicationNew(applicationId string, flags ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.g_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Activate is a wrapper around the C function g_application_activate.
func (recv *Application) Activate() {
	C.g_application_activate((*C.GApplication)(recv.native))

	return
}

// GetApplicationId is a wrapper around the C function g_application_get_application_id.
func (recv *Application) GetApplicationId() string {
	retC := C.g_application_get_application_id((*C.GApplication)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFlags is a wrapper around the C function g_application_get_flags.
func (recv *Application) GetFlags() ApplicationFlags {
	retC := C.g_application_get_flags((*C.GApplication)(recv.native))
	retGo := (ApplicationFlags)(retC)

	return retGo
}

// GetInactivityTimeout is a wrapper around the C function g_application_get_inactivity_timeout.
func (recv *Application) GetInactivityTimeout() uint32 {
	retC := C.g_application_get_inactivity_timeout((*C.GApplication)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetIsRegistered is a wrapper around the C function g_application_get_is_registered.
func (recv *Application) GetIsRegistered() bool {
	retC := C.g_application_get_is_registered((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsRemote is a wrapper around the C function g_application_get_is_remote.
func (recv *Application) GetIsRemote() bool {
	retC := C.g_application_get_is_remote((*C.GApplication)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hold is a wrapper around the C function g_application_hold.
func (recv *Application) Hold() {
	C.g_application_hold((*C.GApplication)(recv.native))

	return
}

// Unsupported : g_application_open : unsupported parameter files :

// Register is a wrapper around the C function g_application_register.
func (recv *Application) Register(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_application_register((*C.GApplication)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Release is a wrapper around the C function g_application_release.
func (recv *Application) Release() {
	C.g_application_release((*C.GApplication)(recv.native))

	return
}

// Run is a wrapper around the C function g_application_run.
func (recv *Application) Run(args []string) int32 {
	cArgc, cArgv := argsIn(args)

	retC := C.g_application_run((*C.GApplication)(recv.native), cArgc, cArgv)
	retGo := (int32)(retC)

	return retGo
}

// SetActionGroup is a wrapper around the C function g_application_set_action_group.
func (recv *Application) SetActionGroup(actionGroup *ActionGroup) {
	c_action_group := (*C.GActionGroup)(actionGroup.ToC())

	C.g_application_set_action_group((*C.GApplication)(recv.native), c_action_group)

	return
}

// SetApplicationId is a wrapper around the C function g_application_set_application_id.
func (recv *Application) SetApplicationId(applicationId string) {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	C.g_application_set_application_id((*C.GApplication)(recv.native), c_application_id)

	return
}

// SetFlags is a wrapper around the C function g_application_set_flags.
func (recv *Application) SetFlags(flags ApplicationFlags) {
	c_flags := (C.GApplicationFlags)(flags)

	C.g_application_set_flags((*C.GApplication)(recv.native), c_flags)

	return
}

// SetInactivityTimeout is a wrapper around the C function g_application_set_inactivity_timeout.
func (recv *Application) SetInactivityTimeout(inactivityTimeout uint32) {
	c_inactivity_timeout := (C.guint)(inactivityTimeout)

	C.g_application_set_inactivity_timeout((*C.GApplication)(recv.native), c_inactivity_timeout)

	return
}

// GetArguments is a wrapper around the C function g_application_command_line_get_arguments.
func (recv *ApplicationCommandLine) GetArguments() ([]string, int32) {
	var c_argc C.int

	retC := C.g_application_command_line_get_arguments((*C.GApplicationCommandLine)(recv.native), &c_argc)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	argc := (int32)(c_argc)

	return retGo, argc
}

// GetCwd is a wrapper around the C function g_application_command_line_get_cwd.
func (recv *ApplicationCommandLine) GetCwd() string {
	retC := C.g_application_command_line_get_cwd((*C.GApplicationCommandLine)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEnviron is a wrapper around the C function g_application_command_line_get_environ.
func (recv *ApplicationCommandLine) GetEnviron() []string {
	retC := C.g_application_command_line_get_environ((*C.GApplicationCommandLine)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetExitStatus is a wrapper around the C function g_application_command_line_get_exit_status.
func (recv *ApplicationCommandLine) GetExitStatus() int32 {
	retC := C.g_application_command_line_get_exit_status((*C.GApplicationCommandLine)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIsRemote is a wrapper around the C function g_application_command_line_get_is_remote.
func (recv *ApplicationCommandLine) GetIsRemote() bool {
	retC := C.g_application_command_line_get_is_remote((*C.GApplicationCommandLine)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPlatformData is a wrapper around the C function g_application_command_line_get_platform_data.
func (recv *ApplicationCommandLine) GetPlatformData() *glib.Variant {
	retC := C.g_application_command_line_get_platform_data((*C.GApplicationCommandLine)(recv.native))
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Getenv is a wrapper around the C function g_application_command_line_getenv.
func (recv *ApplicationCommandLine) Getenv(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_application_command_line_getenv((*C.GApplicationCommandLine)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// Print is a wrapper around the C function g_application_command_line_print.
func (recv *ApplicationCommandLine) Print(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_application_command_line_print((*C.GApplicationCommandLine)(recv.native), c_format)

	return
}

// Printerr is a wrapper around the C function g_application_command_line_printerr.
func (recv *ApplicationCommandLine) Printerr(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_application_command_line_printerr((*C.GApplicationCommandLine)(recv.native), c_format)

	return
}

// SetExitStatus is a wrapper around the C function g_application_command_line_set_exit_status.
func (recv *ApplicationCommandLine) SetExitStatus(exitStatus int32) {
	c_exit_status := (C.int)(exitStatus)

	C.g_application_command_line_set_exit_status((*C.GApplicationCommandLine)(recv.native), c_exit_status)

	return
}

// SourceNew is a wrapper around the C function g_cancellable_source_new.
func (recv *Cancellable) SourceNew() *glib.Source {
	retC := C.g_cancellable_source_new((*C.GCancellable)(recv.native))
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClearEmblems is a wrapper around the C function g_emblemed_icon_clear_emblems.
func (recv *EmblemedIcon) ClearEmblems() {
	C.g_emblemed_icon_clear_emblems((*C.GEmblemedIcon)(recv.native))

	return
}

// IOStreamSpliceFinish is a wrapper around the C function g_io_stream_splice_finish.
func IOStreamSpliceFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_io_stream_splice_finish(c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_io_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SettingsListRelocatableSchemas is a wrapper around the C function g_settings_list_relocatable_schemas.
func SettingsListRelocatableSchemas() []string {
	retC := C.g_settings_list_relocatable_schemas()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetRange is a wrapper around the C function g_settings_get_range.
func (recv *Settings) GetRange(key string) *glib.Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_range((*C.GSettings)(recv.native), c_key)
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RangeCheck is a wrapper around the C function g_settings_range_check.
func (recv *Settings) RangeCheck(key string, value *glib.Variant) bool {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_settings_range_check((*C.GSettings)(recv.native), c_key, c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : g_settings_backend_get_default

type signalSimpleActionActivateDetail struct {
	callback  SimpleActionSignalActivateCallback
	handlerID C.gulong
}

var signalSimpleActionActivateId int
var signalSimpleActionActivateMap = make(map[int]signalSimpleActionActivateDetail)
var signalSimpleActionActivateLock sync.RWMutex

// SimpleActionSignalActivateCallback is a callback function for a 'activate' signal emitted from a SimpleAction.
type SimpleActionSignalActivateCallback func(parameter *glib.Variant)

/*
ConnectActivate connects the callback to the 'activate' signal for the SimpleAction.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *SimpleAction) ConnectActivate(callback SimpleActionSignalActivateCallback) int {
	signalSimpleActionActivateLock.Lock()
	defer signalSimpleActionActivateLock.Unlock()

	signalSimpleActionActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.SimpleAction_signal_connect_activate(instance, C.gpointer(uintptr(signalSimpleActionActivateId)))

	detail := signalSimpleActionActivateDetail{callback, handlerID}
	signalSimpleActionActivateMap[signalSimpleActionActivateId] = detail

	return signalSimpleActionActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the SimpleAction.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *SimpleAction) DisconnectActivate(connectionID int) {
	signalSimpleActionActivateLock.Lock()
	defer signalSimpleActionActivateLock.Unlock()

	detail, exists := signalSimpleActionActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSimpleActionActivateMap, connectionID)
}

//export simpleaction_activateHandler
func simpleaction_activateHandler(_ *C.GObject, c_parameter *C.GVariant, data C.gpointer) {
	signalSimpleActionActivateLock.RLock()
	defer signalSimpleActionActivateLock.RUnlock()

	parameter := glib.VariantNewFromC(unsafe.Pointer(c_parameter))

	index := int(uintptr(data))
	callback := signalSimpleActionActivateMap[index].callback
	callback(parameter)
}

// SimpleActionNew is a wrapper around the C function g_simple_action_new.
func SimpleActionNew(name string, parameterType *glib.VariantType) *SimpleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_parameter_type := (*C.GVariantType)(C.NULL)
	if parameterType != nil {
		c_parameter_type = (*C.GVariantType)(parameterType.ToC())
	}

	retC := C.g_simple_action_new(c_name, c_parameter_type)
	retGo := SimpleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SimpleActionNewStateful is a wrapper around the C function g_simple_action_new_stateful.
func SimpleActionNewStateful(name string, parameterType *glib.VariantType, state *glib.Variant) *SimpleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_parameter_type := (*C.GVariantType)(C.NULL)
	if parameterType != nil {
		c_parameter_type = (*C.GVariantType)(parameterType.ToC())
	}

	c_state := (*C.GVariant)(C.NULL)
	if state != nil {
		c_state = (*C.GVariant)(state.ToC())
	}

	retC := C.g_simple_action_new_stateful(c_name, c_parameter_type, c_state)
	retGo := SimpleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SetEnabled is a wrapper around the C function g_simple_action_set_enabled.
func (recv *SimpleAction) SetEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.g_simple_action_set_enabled((*C.GSimpleAction)(recv.native), c_enabled)

	return
}

// SimpleActionGroup is a wrapper around the C record GSimpleActionGroup.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SimpleActionGroup with another SimpleActionGroup, and returns true if they represent the same GObject.
func (recv *SimpleActionGroup) Equals(other *SimpleActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SimpleActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SimpleActionGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a SimpleActionGroup.
func CastToSimpleActionGroup(object *gobject.Object) *SimpleActionGroup {
	return SimpleActionGroupNewFromC(object.ToC())
}

// SimpleActionGroupNew is a wrapper around the C function g_simple_action_group_new.
func SimpleActionGroupNew() *SimpleActionGroup {
	retC := C.g_simple_action_group_new()
	retGo := SimpleActionGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Insert is a wrapper around the C function g_simple_action_group_insert.
func (recv *SimpleActionGroup) Insert(action *Action) {
	c_action := (*C.GAction)(action.ToC())

	C.g_simple_action_group_insert((*C.GSimpleActionGroup)(recv.native), c_action)

	return
}

// Lookup is a wrapper around the C function g_simple_action_group_lookup.
func (recv *SimpleActionGroup) Lookup(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_simple_action_group_lookup((*C.GSimpleActionGroup)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function g_simple_action_group_remove.
func (recv *SimpleActionGroup) Remove(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_simple_action_group_remove((*C.GSimpleActionGroup)(recv.native), c_action_name)

	return
}

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// TakeError is a wrapper around the C function g_simple_async_result_take_error.
func (recv *SimpleAsyncResult) TakeError(error *glib.Error) {
	c_error := (*C.GError)(C.NULL)
	if error != nil {
		c_error = (*C.GError)(error.ToC())
	}

	C.g_simple_async_result_take_error((*C.GSimpleAsyncResult)(recv.native), c_error)

	return
}

// GetTls is a wrapper around the C function g_socket_client_get_tls.
func (recv *SocketClient) GetTls() bool {
	retC := C.g_socket_client_get_tls((*C.GSocketClient)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTlsValidationFlags is a wrapper around the C function g_socket_client_get_tls_validation_flags.
func (recv *SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	retC := C.g_socket_client_get_tls_validation_flags((*C.GSocketClient)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// SetTls is a wrapper around the C function g_socket_client_set_tls.
func (recv *SocketClient) SetTls(tls bool) {
	c_tls :=
		boolToGboolean(tls)

	C.g_socket_client_set_tls((*C.GSocketClient)(recv.native), c_tls)

	return
}

// SetTlsValidationFlags is a wrapper around the C function g_socket_client_set_tls_validation_flags.
func (recv *SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_socket_client_set_tls_validation_flags((*C.GSocketClient)(recv.native), c_flags)

	return
}

// TcpWrapperConnectionNew is a wrapper around the C function g_tcp_wrapper_connection_new.
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

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// TlsCertificate is a wrapper around the C record GTlsCertificate.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsCertificate) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsCertificate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsCertificate with another TlsCertificate, and returns true if they represent the same GObject.
func (recv *TlsCertificate) Equals(other *TlsCertificate) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *TlsCertificate) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to TlsCertificate.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsCertificate.
func CastToTlsCertificate(object *gobject.Object) *TlsCertificate {
	return TlsCertificateNewFromC(object.ToC())
}

// TlsCertificateNewFromFile is a wrapper around the C function g_tls_certificate_new_from_file.
func TlsCertificateNewFromFile(file string) (*TlsCertificate, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_file(c_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsCertificateNewFromFiles is a wrapper around the C function g_tls_certificate_new_from_files.
func TlsCertificateNewFromFiles(certFile string, keyFile string) (*TlsCertificate, error) {
	c_cert_file := C.CString(certFile)
	defer C.free(unsafe.Pointer(c_cert_file))

	c_key_file := C.CString(keyFile)
	defer C.free(unsafe.Pointer(c_key_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_files(c_cert_file, c_key_file, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsCertificateNewFromPem is a wrapper around the C function g_tls_certificate_new_from_pem.
func TlsCertificateNewFromPem(data string) (*TlsCertificate, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gssize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_new_from_pem(c_data, c_length, &cThrowableError)
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsCertificateListNewFromFile is a wrapper around the C function g_tls_certificate_list_new_from_file.
func TlsCertificateListNewFromFile(file string) (*glib.List, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var cThrowableError *C.GError

	retC := C.g_tls_certificate_list_new_from_file(c_file, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetIssuer is a wrapper around the C function g_tls_certificate_get_issuer.
func (recv *TlsCertificate) GetIssuer() *TlsCertificate {
	retC := C.g_tls_certificate_get_issuer((*C.GTlsCertificate)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Verify is a wrapper around the C function g_tls_certificate_verify.
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

// TlsConnection is a wrapper around the C record GTlsConnection.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TlsConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsConnection with another TlsConnection, and returns true if they represent the same GObject.
func (recv *TlsConnection) Equals(other *TlsConnection) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *TlsConnection) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *TlsConnection) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to TlsConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TlsConnection.
func CastToTlsConnection(object *gobject.Object) *TlsConnection {
	return TlsConnectionNewFromC(object.ToC())
}

type signalTlsConnectionAcceptCertificateDetail struct {
	callback  TlsConnectionSignalAcceptCertificateCallback
	handlerID C.gulong
}

var signalTlsConnectionAcceptCertificateId int
var signalTlsConnectionAcceptCertificateMap = make(map[int]signalTlsConnectionAcceptCertificateDetail)
var signalTlsConnectionAcceptCertificateLock sync.RWMutex

// TlsConnectionSignalAcceptCertificateCallback is a callback function for a 'accept-certificate' signal emitted from a TlsConnection.
type TlsConnectionSignalAcceptCertificateCallback func(peerCert *TlsCertificate, errors TlsCertificateFlags) bool

/*
ConnectAcceptCertificate connects the callback to the 'accept-certificate' signal for the TlsConnection.

The returned value represents the connection, and may be passed to DisconnectAcceptCertificate to remove it.
*/
func (recv *TlsConnection) ConnectAcceptCertificate(callback TlsConnectionSignalAcceptCertificateCallback) int {
	signalTlsConnectionAcceptCertificateLock.Lock()
	defer signalTlsConnectionAcceptCertificateLock.Unlock()

	signalTlsConnectionAcceptCertificateId++
	instance := C.gpointer(recv.native)
	handlerID := C.TlsConnection_signal_connect_accept_certificate(instance, C.gpointer(uintptr(signalTlsConnectionAcceptCertificateId)))

	detail := signalTlsConnectionAcceptCertificateDetail{callback, handlerID}
	signalTlsConnectionAcceptCertificateMap[signalTlsConnectionAcceptCertificateId] = detail

	return signalTlsConnectionAcceptCertificateId
}

/*
DisconnectAcceptCertificate disconnects a callback from the 'accept-certificate' signal for the TlsConnection.

The connectionID should be a value returned from a call to ConnectAcceptCertificate.
*/
func (recv *TlsConnection) DisconnectAcceptCertificate(connectionID int) {
	signalTlsConnectionAcceptCertificateLock.Lock()
	defer signalTlsConnectionAcceptCertificateLock.Unlock()

	detail, exists := signalTlsConnectionAcceptCertificateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalTlsConnectionAcceptCertificateMap, connectionID)
}

//export tlsconnection_acceptCertificateHandler
func tlsconnection_acceptCertificateHandler(_ *C.GObject, c_peer_cert *C.GTlsCertificate, c_errors C.GTlsCertificateFlags, data C.gpointer) C.gboolean {
	signalTlsConnectionAcceptCertificateLock.RLock()
	defer signalTlsConnectionAcceptCertificateLock.RUnlock()

	peerCert := TlsCertificateNewFromC(unsafe.Pointer(c_peer_cert))

	errors := TlsCertificateFlags(c_errors)

	index := int(uintptr(data))
	callback := signalTlsConnectionAcceptCertificateMap[index].callback
	retGo := callback(peerCert, errors)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// EmitAcceptCertificate is a wrapper around the C function g_tls_connection_emit_accept_certificate.
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

// GetCertificate is a wrapper around the C function g_tls_connection_get_certificate.
func (recv *TlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificate is a wrapper around the C function g_tls_connection_get_peer_certificate.
func (recv *TlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_tls_connection_get_peer_certificate((*C.GTlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificateErrors is a wrapper around the C function g_tls_connection_get_peer_certificate_errors.
func (recv *TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_tls_connection_get_peer_certificate_errors((*C.GTlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// GetRehandshakeMode is a wrapper around the C function g_tls_connection_get_rehandshake_mode.
func (recv *TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_tls_connection_get_rehandshake_mode((*C.GTlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// GetRequireCloseNotify is a wrapper around the C function g_tls_connection_get_require_close_notify.
func (recv *TlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_tls_connection_get_require_close_notify((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseSystemCertdb is a wrapper around the C function g_tls_connection_get_use_system_certdb.
func (recv *TlsConnection) GetUseSystemCertdb() bool {
	retC := C.g_tls_connection_get_use_system_certdb((*C.GTlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Handshake is a wrapper around the C function g_tls_connection_handshake.
func (recv *TlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_connection_handshake((*C.GTlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_tls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// HandshakeFinish is a wrapper around the C function g_tls_connection_handshake_finish.
func (recv *TlsConnection) HandshakeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_connection_handshake_finish((*C.GTlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCertificate is a wrapper around the C function g_tls_connection_set_certificate.
func (recv *TlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	C.g_tls_connection_set_certificate((*C.GTlsConnection)(recv.native), c_certificate)

	return
}

// SetRehandshakeMode is a wrapper around the C function g_tls_connection_set_rehandshake_mode.
func (recv *TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_tls_connection_set_rehandshake_mode((*C.GTlsConnection)(recv.native), c_mode)

	return
}

// SetRequireCloseNotify is a wrapper around the C function g_tls_connection_set_require_close_notify.
func (recv *TlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_tls_connection_set_require_close_notify((*C.GTlsConnection)(recv.native), c_require_close_notify)

	return
}

// SetUseSystemCertdb is a wrapper around the C function g_tls_connection_set_use_system_certdb.
func (recv *TlsConnection) SetUseSystemCertdb(useSystemCertdb bool) {
	c_use_system_certdb :=
		boolToGboolean(useSystemCertdb)

	C.g_tls_connection_set_use_system_certdb((*C.GTlsConnection)(recv.native), c_use_system_certdb)

	return
}

type TlsAuthenticationMode C.GTlsAuthenticationMode

const (
	TLS_AUTHENTICATION_NONE      TlsAuthenticationMode = 0
	TLS_AUTHENTICATION_REQUESTED TlsAuthenticationMode = 1
	TLS_AUTHENTICATION_REQUIRED  TlsAuthenticationMode = 2
)

type TlsError C.GTlsError

const (
	TLS_ERROR_UNAVAILABLE          TlsError = 0
	TLS_ERROR_MISC                 TlsError = 1
	TLS_ERROR_BAD_CERTIFICATE      TlsError = 2
	TLS_ERROR_NOT_TLS              TlsError = 3
	TLS_ERROR_HANDSHAKE            TlsError = 4
	TLS_ERROR_CERTIFICATE_REQUIRED TlsError = 5
	TLS_ERROR_EOF                  TlsError = 6
)

// TlsErrorQuark is a wrapper around the C function g_tls_error_quark.
func TlsErrorQuark() glib.Quark {
	retC := C.g_tls_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type TlsRehandshakeMode C.GTlsRehandshakeMode

const (
	TLS_REHANDSHAKE_NEVER    TlsRehandshakeMode = 0
	TLS_REHANDSHAKE_SAFELY   TlsRehandshakeMode = 1
	TLS_REHANDSHAKE_UNSAFELY TlsRehandshakeMode = 2
)

// Blacklisted : g_memory_settings_backend_new

// Blacklisted : g_null_settings_backend_new

// PollableSourceNew is a wrapper around the C function g_pollable_source_new.
func PollableSourceNew(pollableStream *gobject.Object) *glib.Source {
	c_pollable_stream := (*C.GObject)(C.NULL)
	if pollableStream != nil {
		c_pollable_stream = (*C.GObject)(pollableStream.ToC())
	}

	retC := C.g_pollable_source_new(c_pollable_stream)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_simple_async_report_take_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Activate is a wrapper around the C function g_action_activate.
func (recv *Action) Activate(parameter *glib.Variant) {
	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.g_action_activate((*C.GAction)(recv.native), c_parameter)

	return
}

// GetEnabled is a wrapper around the C function g_action_get_enabled.
func (recv *Action) GetEnabled() bool {
	retC := C.g_action_get_enabled((*C.GAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function g_action_get_name.
func (recv *Action) GetName() string {
	retC := C.g_action_get_name((*C.GAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetParameterType is a wrapper around the C function g_action_get_parameter_type.
func (recv *Action) GetParameterType() *glib.VariantType {
	retC := C.g_action_get_parameter_type((*C.GAction)(recv.native))
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetState is a wrapper around the C function g_action_get_state.
func (recv *Action) GetState() *glib.Variant {
	retC := C.g_action_get_state((*C.GAction)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStateHint is a wrapper around the C function g_action_get_state_hint.
func (recv *Action) GetStateHint() *glib.Variant {
	retC := C.g_action_get_state_hint((*C.GAction)(recv.native))
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStateType is a wrapper around the C function g_action_get_state_type.
func (recv *Action) GetStateType() *glib.VariantType {
	retC := C.g_action_get_state_type((*C.GAction)(recv.native))
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type signalActionGroupActionAddedDetail struct {
	callback  ActionGroupSignalActionAddedCallback
	handlerID C.gulong
}

var signalActionGroupActionAddedId int
var signalActionGroupActionAddedMap = make(map[int]signalActionGroupActionAddedDetail)
var signalActionGroupActionAddedLock sync.RWMutex

// ActionGroupSignalActionAddedCallback is a callback function for a 'action-added' signal emitted from a ActionGroup.
type ActionGroupSignalActionAddedCallback func(actionName string)

/*
ConnectActionAdded connects the callback to the 'action-added' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionAdded to remove it.
*/
func (recv *ActionGroup) ConnectActionAdded(callback ActionGroupSignalActionAddedCallback) int {
	signalActionGroupActionAddedLock.Lock()
	defer signalActionGroupActionAddedLock.Unlock()

	signalActionGroupActionAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_added(instance, C.gpointer(uintptr(signalActionGroupActionAddedId)))

	detail := signalActionGroupActionAddedDetail{callback, handlerID}
	signalActionGroupActionAddedMap[signalActionGroupActionAddedId] = detail

	return signalActionGroupActionAddedId
}

/*
DisconnectActionAdded disconnects a callback from the 'action-added' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionAdded.
*/
func (recv *ActionGroup) DisconnectActionAdded(connectionID int) {
	signalActionGroupActionAddedLock.Lock()
	defer signalActionGroupActionAddedLock.Unlock()

	detail, exists := signalActionGroupActionAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionAddedMap, connectionID)
}

//export actiongroup_actionAddedHandler
func actiongroup_actionAddedHandler(_ *C.GObject, c_action_name *C.gchar, data C.gpointer) {
	signalActionGroupActionAddedLock.RLock()
	defer signalActionGroupActionAddedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	index := int(uintptr(data))
	callback := signalActionGroupActionAddedMap[index].callback
	callback(actionName)
}

type signalActionGroupActionEnabledChangedDetail struct {
	callback  ActionGroupSignalActionEnabledChangedCallback
	handlerID C.gulong
}

var signalActionGroupActionEnabledChangedId int
var signalActionGroupActionEnabledChangedMap = make(map[int]signalActionGroupActionEnabledChangedDetail)
var signalActionGroupActionEnabledChangedLock sync.RWMutex

// ActionGroupSignalActionEnabledChangedCallback is a callback function for a 'action-enabled-changed' signal emitted from a ActionGroup.
type ActionGroupSignalActionEnabledChangedCallback func(actionName string, enabled bool)

/*
ConnectActionEnabledChanged connects the callback to the 'action-enabled-changed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionEnabledChanged to remove it.
*/
func (recv *ActionGroup) ConnectActionEnabledChanged(callback ActionGroupSignalActionEnabledChangedCallback) int {
	signalActionGroupActionEnabledChangedLock.Lock()
	defer signalActionGroupActionEnabledChangedLock.Unlock()

	signalActionGroupActionEnabledChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_enabled_changed(instance, C.gpointer(uintptr(signalActionGroupActionEnabledChangedId)))

	detail := signalActionGroupActionEnabledChangedDetail{callback, handlerID}
	signalActionGroupActionEnabledChangedMap[signalActionGroupActionEnabledChangedId] = detail

	return signalActionGroupActionEnabledChangedId
}

/*
DisconnectActionEnabledChanged disconnects a callback from the 'action-enabled-changed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionEnabledChanged.
*/
func (recv *ActionGroup) DisconnectActionEnabledChanged(connectionID int) {
	signalActionGroupActionEnabledChangedLock.Lock()
	defer signalActionGroupActionEnabledChangedLock.Unlock()

	detail, exists := signalActionGroupActionEnabledChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionEnabledChangedMap, connectionID)
}

//export actiongroup_actionEnabledChangedHandler
func actiongroup_actionEnabledChangedHandler(_ *C.GObject, c_action_name *C.gchar, c_enabled C.gboolean, data C.gpointer) {
	signalActionGroupActionEnabledChangedLock.RLock()
	defer signalActionGroupActionEnabledChangedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	enabled := c_enabled == C.TRUE

	index := int(uintptr(data))
	callback := signalActionGroupActionEnabledChangedMap[index].callback
	callback(actionName, enabled)
}

type signalActionGroupActionRemovedDetail struct {
	callback  ActionGroupSignalActionRemovedCallback
	handlerID C.gulong
}

var signalActionGroupActionRemovedId int
var signalActionGroupActionRemovedMap = make(map[int]signalActionGroupActionRemovedDetail)
var signalActionGroupActionRemovedLock sync.RWMutex

// ActionGroupSignalActionRemovedCallback is a callback function for a 'action-removed' signal emitted from a ActionGroup.
type ActionGroupSignalActionRemovedCallback func(actionName string)

/*
ConnectActionRemoved connects the callback to the 'action-removed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionRemoved to remove it.
*/
func (recv *ActionGroup) ConnectActionRemoved(callback ActionGroupSignalActionRemovedCallback) int {
	signalActionGroupActionRemovedLock.Lock()
	defer signalActionGroupActionRemovedLock.Unlock()

	signalActionGroupActionRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_removed(instance, C.gpointer(uintptr(signalActionGroupActionRemovedId)))

	detail := signalActionGroupActionRemovedDetail{callback, handlerID}
	signalActionGroupActionRemovedMap[signalActionGroupActionRemovedId] = detail

	return signalActionGroupActionRemovedId
}

/*
DisconnectActionRemoved disconnects a callback from the 'action-removed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionRemoved.
*/
func (recv *ActionGroup) DisconnectActionRemoved(connectionID int) {
	signalActionGroupActionRemovedLock.Lock()
	defer signalActionGroupActionRemovedLock.Unlock()

	detail, exists := signalActionGroupActionRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionRemovedMap, connectionID)
}

//export actiongroup_actionRemovedHandler
func actiongroup_actionRemovedHandler(_ *C.GObject, c_action_name *C.gchar, data C.gpointer) {
	signalActionGroupActionRemovedLock.RLock()
	defer signalActionGroupActionRemovedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	index := int(uintptr(data))
	callback := signalActionGroupActionRemovedMap[index].callback
	callback(actionName)
}

type signalActionGroupActionStateChangedDetail struct {
	callback  ActionGroupSignalActionStateChangedCallback
	handlerID C.gulong
}

var signalActionGroupActionStateChangedId int
var signalActionGroupActionStateChangedMap = make(map[int]signalActionGroupActionStateChangedDetail)
var signalActionGroupActionStateChangedLock sync.RWMutex

// ActionGroupSignalActionStateChangedCallback is a callback function for a 'action-state-changed' signal emitted from a ActionGroup.
type ActionGroupSignalActionStateChangedCallback func(actionName string, value *glib.Variant)

/*
ConnectActionStateChanged connects the callback to the 'action-state-changed' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectActionStateChanged to remove it.
*/
func (recv *ActionGroup) ConnectActionStateChanged(callback ActionGroupSignalActionStateChangedCallback) int {
	signalActionGroupActionStateChangedLock.Lock()
	defer signalActionGroupActionStateChangedLock.Unlock()

	signalActionGroupActionStateChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_action_state_changed(instance, C.gpointer(uintptr(signalActionGroupActionStateChangedId)))

	detail := signalActionGroupActionStateChangedDetail{callback, handlerID}
	signalActionGroupActionStateChangedMap[signalActionGroupActionStateChangedId] = detail

	return signalActionGroupActionStateChangedId
}

/*
DisconnectActionStateChanged disconnects a callback from the 'action-state-changed' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectActionStateChanged.
*/
func (recv *ActionGroup) DisconnectActionStateChanged(connectionID int) {
	signalActionGroupActionStateChangedLock.Lock()
	defer signalActionGroupActionStateChangedLock.Unlock()

	detail, exists := signalActionGroupActionStateChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupActionStateChangedMap, connectionID)
}

//export actiongroup_actionStateChangedHandler
func actiongroup_actionStateChangedHandler(_ *C.GObject, c_action_name *C.gchar, c_value *C.GVariant, data C.gpointer) {
	signalActionGroupActionStateChangedLock.RLock()
	defer signalActionGroupActionStateChangedLock.RUnlock()

	actionName := C.GoString(c_action_name)

	value := glib.VariantNewFromC(unsafe.Pointer(c_value))

	index := int(uintptr(data))
	callback := signalActionGroupActionStateChangedMap[index].callback
	callback(actionName, value)
}

// ActionAdded is a wrapper around the C function g_action_group_action_added.
func (recv *ActionGroup) ActionAdded(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_added((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// ActionEnabledChanged is a wrapper around the C function g_action_group_action_enabled_changed.
func (recv *ActionGroup) ActionEnabledChanged(actionName string, enabled bool) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_enabled :=
		boolToGboolean(enabled)

	C.g_action_group_action_enabled_changed((*C.GActionGroup)(recv.native), c_action_name, c_enabled)

	return
}

// ActionRemoved is a wrapper around the C function g_action_group_action_removed.
func (recv *ActionGroup) ActionRemoved(actionName string) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	C.g_action_group_action_removed((*C.GActionGroup)(recv.native), c_action_name)

	return
}

// ActionStateChanged is a wrapper around the C function g_action_group_action_state_changed.
func (recv *ActionGroup) ActionStateChanged(actionName string, state *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_state := (*C.GVariant)(C.NULL)
	if state != nil {
		c_state = (*C.GVariant)(state.ToC())
	}

	C.g_action_group_action_state_changed((*C.GActionGroup)(recv.native), c_action_name, c_state)

	return
}

// ActivateAction is a wrapper around the C function g_action_group_activate_action.
func (recv *ActionGroup) ActivateAction(actionName string, parameter *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_parameter := (*C.GVariant)(C.NULL)
	if parameter != nil {
		c_parameter = (*C.GVariant)(parameter.ToC())
	}

	C.g_action_group_activate_action((*C.GActionGroup)(recv.native), c_action_name, c_parameter)

	return
}

// ChangeActionState is a wrapper around the C function g_action_group_change_action_state.
func (recv *ActionGroup) ChangeActionState(actionName string, value *glib.Variant) {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	C.g_action_group_change_action_state((*C.GActionGroup)(recv.native), c_action_name, c_value)

	return
}

// GetActionEnabled is a wrapper around the C function g_action_group_get_action_enabled.
func (recv *ActionGroup) GetActionEnabled(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_enabled((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// GetActionParameterType is a wrapper around the C function g_action_group_get_action_parameter_type.
func (recv *ActionGroup) GetActionParameterType(actionName string) *glib.VariantType {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_parameter_type((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionState is a wrapper around the C function g_action_group_get_action_state.
func (recv *ActionGroup) GetActionState(actionName string) *glib.Variant {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionStateHint is a wrapper around the C function g_action_group_get_action_state_hint.
func (recv *ActionGroup) GetActionStateHint(actionName string) *glib.Variant {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state_hint((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionStateType is a wrapper around the C function g_action_group_get_action_state_type.
func (recv *ActionGroup) GetActionStateType(actionName string) *glib.VariantType {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_get_action_state_type((*C.GActionGroup)(recv.native), c_action_name)
	var retGo (*glib.VariantType)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantTypeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// HasAction is a wrapper around the C function g_action_group_has_action.
func (recv *ActionGroup) HasAction(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_group_has_action((*C.GActionGroup)(recv.native), c_action_name)
	retGo := retC == C.TRUE

	return retGo
}

// ListActions is a wrapper around the C function g_action_group_list_actions.
func (recv *ActionGroup) ListActions() []string {
	retC := C.g_action_group_list_actions((*C.GActionGroup)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// AppInfoGetFallbackForType is a wrapper around the C function g_app_info_get_fallback_for_type.
func AppInfoGetFallbackForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_fallback_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppInfoGetRecommendedForType is a wrapper around the C function g_app_info_get_recommended_for_type.
func AppInfoGetRecommendedForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_recommended_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PollableInputStream is a wrapper around the C record GPollableInputStream.
type PollableInputStream struct {
	native *C.GPollableInputStream
}

func PollableInputStreamNewFromC(u unsafe.Pointer) *PollableInputStream {
	c := (*C.GPollableInputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStream{native: c}

	return g
}

func (recv *PollableInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableInputStream with another PollableInputStream, and returns true if they represent the same GObject.
func (recv *PollableInputStream) Equals(other *PollableInputStream) bool {
	return other.ToC() == recv.ToC()
}

// CanPoll is a wrapper around the C function g_pollable_input_stream_can_poll.
func (recv *PollableInputStream) CanPoll() bool {
	retC := C.g_pollable_input_stream_can_poll((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_input_stream_create_source.
func (recv *PollableInputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_input_stream_create_source((*C.GPollableInputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsReadable is a wrapper around the C function g_pollable_input_stream_is_readable.
func (recv *PollableInputStream) IsReadable() bool {
	retC := C.g_pollable_input_stream_is_readable((*C.GPollableInputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ReadNonblocking is a wrapper around the C function g_pollable_input_stream_read_nonblocking.
func (recv *PollableInputStream) ReadNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
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

	retC := C.g_pollable_input_stream_read_nonblocking((*C.GPollableInputStream)(recv.native), c_buffer, c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PollableOutputStream is a wrapper around the C record GPollableOutputStream.
type PollableOutputStream struct {
	native *C.GPollableOutputStream
}

func PollableOutputStreamNewFromC(u unsafe.Pointer) *PollableOutputStream {
	c := (*C.GPollableOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStream{native: c}

	return g
}

func (recv *PollableOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableOutputStream with another PollableOutputStream, and returns true if they represent the same GObject.
func (recv *PollableOutputStream) Equals(other *PollableOutputStream) bool {
	return other.ToC() == recv.ToC()
}

// CanPoll is a wrapper around the C function g_pollable_output_stream_can_poll.
func (recv *PollableOutputStream) CanPoll() bool {
	retC := C.g_pollable_output_stream_can_poll((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CreateSource is a wrapper around the C function g_pollable_output_stream_create_source.
func (recv *PollableOutputStream) CreateSource(cancellable *Cancellable) *glib.Source {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_pollable_output_stream_create_source((*C.GPollableOutputStream)(recv.native), c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsWritable is a wrapper around the C function g_pollable_output_stream_is_writable.
func (recv *PollableOutputStream) IsWritable() bool {
	retC := C.g_pollable_output_stream_is_writable((*C.GPollableOutputStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// WriteNonblocking is a wrapper around the C function g_pollable_output_stream_write_nonblocking.
func (recv *PollableOutputStream) WriteNonblocking(buffer []uint8, cancellable *Cancellable) (int64, error) {
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

	retC := C.g_pollable_output_stream_write_nonblocking((*C.GPollableOutputStream)(recv.native), c_buffer, c_count, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TlsBackend is a wrapper around the C record GTlsBackend.
type TlsBackend struct {
	native *C.GTlsBackend
}

func TlsBackendNewFromC(u unsafe.Pointer) *TlsBackend {
	c := (*C.GTlsBackend)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackend{native: c}

	return g
}

func (recv *TlsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsBackend with another TlsBackend, and returns true if they represent the same GObject.
func (recv *TlsBackend) Equals(other *TlsBackend) bool {
	return other.ToC() == recv.ToC()
}

// TlsBackendGetDefault is a wrapper around the C function g_tls_backend_get_default.
func TlsBackendGetDefault() *TlsBackend {
	retC := C.g_tls_backend_get_default()
	retGo := TlsBackendNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCertificateType is a wrapper around the C function g_tls_backend_get_certificate_type.
func (recv *TlsBackend) GetCertificateType() gobject.Type {
	retC := C.g_tls_backend_get_certificate_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetClientConnectionType is a wrapper around the C function g_tls_backend_get_client_connection_type.
func (recv *TlsBackend) GetClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetServerConnectionType is a wrapper around the C function g_tls_backend_get_server_connection_type.
func (recv *TlsBackend) GetServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// SupportsTls is a wrapper around the C function g_tls_backend_supports_tls.
func (recv *TlsBackend) SupportsTls() bool {
	retC := C.g_tls_backend_supports_tls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// TlsClientConnection is a wrapper around the C record GTlsClientConnection.
type TlsClientConnection struct {
	native *C.GTlsClientConnection
}

func TlsClientConnectionNewFromC(u unsafe.Pointer) *TlsClientConnection {
	c := (*C.GTlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsClientConnection{native: c}

	return g
}

func (recv *TlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsClientConnection with another TlsClientConnection, and returns true if they represent the same GObject.
func (recv *TlsClientConnection) Equals(other *TlsClientConnection) bool {
	return other.ToC() == recv.ToC()
}

// TlsClientConnectionNew is a wrapper around the C function g_tls_client_connection_new.
func TlsClientConnectionNew(baseIoStream *IOStream, serverIdentity *SocketConnectable) (*TlsClientConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_server_identity := (*C.GSocketConnectable)(serverIdentity.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_client_connection_new(c_base_io_stream, c_server_identity, &cThrowableError)
	retGo := TlsClientConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAcceptedCas is a wrapper around the C function g_tls_client_connection_get_accepted_cas.
func (recv *TlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_tls_client_connection_get_accepted_cas((*C.GTlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetServerIdentity is a wrapper around the C function g_tls_client_connection_get_server_identity.
func (recv *TlsClientConnection) GetServerIdentity() *SocketConnectable {
	retC := C.g_tls_client_connection_get_server_identity((*C.GTlsClientConnection)(recv.native))
	retGo := SocketConnectableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUseSsl3 is a wrapper around the C function g_tls_client_connection_get_use_ssl3.
func (recv *TlsClientConnection) GetUseSsl3() bool {
	retC := C.g_tls_client_connection_get_use_ssl3((*C.GTlsClientConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValidationFlags is a wrapper around the C function g_tls_client_connection_get_validation_flags.
func (recv *TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_tls_client_connection_get_validation_flags((*C.GTlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// SetServerIdentity is a wrapper around the C function g_tls_client_connection_set_server_identity.
func (recv *TlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	C.g_tls_client_connection_set_server_identity((*C.GTlsClientConnection)(recv.native), c_identity)

	return
}

// SetUseSsl3 is a wrapper around the C function g_tls_client_connection_set_use_ssl3.
func (recv *TlsClientConnection) SetUseSsl3(useSsl3 bool) {
	c_use_ssl3 :=
		boolToGboolean(useSsl3)

	C.g_tls_client_connection_set_use_ssl3((*C.GTlsClientConnection)(recv.native), c_use_ssl3)

	return
}

// SetValidationFlags is a wrapper around the C function g_tls_client_connection_set_validation_flags.
func (recv *TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_tls_client_connection_set_validation_flags((*C.GTlsClientConnection)(recv.native), c_flags)

	return
}

// TlsServerConnection is a wrapper around the C record GTlsServerConnection.
type TlsServerConnection struct {
	native *C.GTlsServerConnection
}

func TlsServerConnectionNewFromC(u unsafe.Pointer) *TlsServerConnection {
	c := (*C.GTlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &TlsServerConnection{native: c}

	return g
}

func (recv *TlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsServerConnection with another TlsServerConnection, and returns true if they represent the same GObject.
func (recv *TlsServerConnection) Equals(other *TlsServerConnection) bool {
	return other.ToC() == recv.ToC()
}

// TlsServerConnectionNew is a wrapper around the C function g_tls_server_connection_new.
func TlsServerConnectionNew(baseIoStream *IOStream, certificate *TlsCertificate) (*TlsServerConnection, error) {
	c_base_io_stream := (*C.GIOStream)(C.NULL)
	if baseIoStream != nil {
		c_base_io_stream = (*C.GIOStream)(baseIoStream.ToC())
	}

	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_server_connection_new(c_base_io_stream, c_certificate, &cThrowableError)
	retGo := TlsServerConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ActionGroupInterface is a wrapper around the C record GActionGroupInterface.
type ActionGroupInterface struct {
	native *C.GActionGroupInterface
	// g_iface : record
	// no type for has_action
	// no type for list_actions
	// no type for get_action_enabled
	// no type for get_action_parameter_type
	// no type for get_action_state_type
	// no type for get_action_state_hint
	// no type for get_action_state
	// no type for change_action_state
	// no type for activate_action
	// no type for action_added
	// no type for action_removed
	// no type for action_enabled_changed
	// no type for action_state_changed
	// no type for query_action
}

func ActionGroupInterfaceNewFromC(u unsafe.Pointer) *ActionGroupInterface {
	c := (*C.GActionGroupInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroupInterface{native: c}

	return g
}

func (recv *ActionGroupInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionGroupInterface with another ActionGroupInterface, and returns true if they represent the same GObject.
func (recv *ActionGroupInterface) Equals(other *ActionGroupInterface) bool {
	return other.ToC() == recv.ToC()
}

// ActionInterface is a wrapper around the C record GActionInterface.
type ActionInterface struct {
	native *C.GActionInterface
	// g_iface : record
	// no type for get_name
	// no type for get_parameter_type
	// no type for get_state_type
	// no type for get_state_hint
	// no type for get_enabled
	// no type for get_state
	// no type for change_state
	// no type for activate
}

func ActionInterfaceNewFromC(u unsafe.Pointer) *ActionInterface {
	c := (*C.GActionInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionInterface{native: c}

	return g
}

func (recv *ActionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionInterface with another ActionInterface, and returns true if they represent the same GObject.
func (recv *ActionInterface) Equals(other *ActionInterface) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationClass is a wrapper around the C record GApplicationClass.
type ApplicationClass struct {
	native *C.GApplicationClass
	// Private : parent_class
	// no type for startup
	// no type for activate
	// no type for open
	// no type for command_line
	// no type for local_command_line
	// no type for before_emit
	// no type for after_emit
	// no type for add_platform_data
	// no type for quit_mainloop
	// no type for run_mainloop
	// no type for shutdown
	// no type for dbus_register
	// no type for dbus_unregister
	// no type for handle_local_options
	// Private : padding
}

func ApplicationClassNewFromC(u unsafe.Pointer) *ApplicationClass {
	c := (*C.GApplicationClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationClass{native: c}

	return g
}

func (recv *ApplicationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationClass with another ApplicationClass, and returns true if they represent the same GObject.
func (recv *ApplicationClass) Equals(other *ApplicationClass) bool {
	return other.ToC() == recv.ToC()
}

// ApplicationCommandLineClass is a wrapper around the C record GApplicationCommandLineClass.
type ApplicationCommandLineClass struct {
	native *C.GApplicationCommandLineClass
	// Private : parent_class
	// no type for print_literal
	// no type for printerr_literal
	// no type for get_stdin
	// Private : padding
}

func ApplicationCommandLineClassNewFromC(u unsafe.Pointer) *ApplicationCommandLineClass {
	c := (*C.GApplicationCommandLineClass)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLineClass{native: c}

	return g
}

func (recv *ApplicationCommandLineClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationCommandLineClass with another ApplicationCommandLineClass, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLineClass) Equals(other *ApplicationCommandLineClass) bool {
	return other.ToC() == recv.ToC()
}

// PollableInputStreamInterface is a wrapper around the C record GPollableInputStreamInterface.
type PollableInputStreamInterface struct {
	native *C.GPollableInputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_readable
	// no type for create_source
	// no type for read_nonblocking
}

func PollableInputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableInputStreamInterface {
	c := (*C.GPollableInputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableInputStreamInterface{native: c}

	return g
}

func (recv *PollableInputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableInputStreamInterface with another PollableInputStreamInterface, and returns true if they represent the same GObject.
func (recv *PollableInputStreamInterface) Equals(other *PollableInputStreamInterface) bool {
	return other.ToC() == recv.ToC()
}

// PollableOutputStreamInterface is a wrapper around the C record GPollableOutputStreamInterface.
type PollableOutputStreamInterface struct {
	native *C.GPollableOutputStreamInterface
	// g_iface : record
	// no type for can_poll
	// no type for is_writable
	// no type for create_source
	// no type for write_nonblocking
}

func PollableOutputStreamInterfaceNewFromC(u unsafe.Pointer) *PollableOutputStreamInterface {
	c := (*C.GPollableOutputStreamInterface)(u)
	if c == nil {
		return nil
	}

	g := &PollableOutputStreamInterface{native: c}

	return g
}

func (recv *PollableOutputStreamInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PollableOutputStreamInterface with another PollableOutputStreamInterface, and returns true if they represent the same GObject.
func (recv *PollableOutputStreamInterface) Equals(other *PollableOutputStreamInterface) bool {
	return other.ToC() == recv.ToC()
}

// TlsBackendInterface is a wrapper around the C record GTlsBackendInterface.
type TlsBackendInterface struct {
	native *C.GTlsBackendInterface
	// g_iface : record
	// no type for supports_tls
	// no type for get_certificate_type
	// no type for get_client_connection_type
	// no type for get_server_connection_type
	// no type for get_file_database_type
	// no type for get_default_database
	// no type for supports_dtls
	// no type for get_dtls_client_connection_type
	// no type for get_dtls_server_connection_type
}

func TlsBackendInterfaceNewFromC(u unsafe.Pointer) *TlsBackendInterface {
	c := (*C.GTlsBackendInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsBackendInterface{native: c}

	return g
}

func (recv *TlsBackendInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TlsBackendInterface with another TlsBackendInterface, and returns true if they represent the same GObject.
func (recv *TlsBackendInterface) Equals(other *TlsBackendInterface) bool {
	return other.ToC() == recv.ToC()
}
