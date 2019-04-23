// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
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
import "C"

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

// Blacklisted : g_application_activate

// Blacklisted : g_application_get_application_id

// Blacklisted : g_application_get_flags

// Blacklisted : g_application_get_inactivity_timeout

// Blacklisted : g_application_get_is_registered

// Blacklisted : g_application_get_is_remote

// Blacklisted : g_application_hold

// Unsupported : g_application_open : unsupported parameter files :

// Blacklisted : g_application_register

// Blacklisted : g_application_release

// Blacklisted : g_application_run

// Blacklisted : g_application_set_action_group

// Blacklisted : g_application_set_application_id

// Blacklisted : g_application_set_flags

// Blacklisted : g_application_set_inactivity_timeout

// Blacklisted : g_application_command_line_get_arguments

// Blacklisted : g_application_command_line_get_cwd

// Blacklisted : g_application_command_line_get_environ

// Blacklisted : g_application_command_line_get_exit_status

// Blacklisted : g_application_command_line_get_is_remote

// Blacklisted : g_application_command_line_get_platform_data

// Blacklisted : g_application_command_line_getenv

// Blacklisted : g_application_command_line_print

// Blacklisted : g_application_command_line_printerr

// Blacklisted : g_application_command_line_set_exit_status

// Blacklisted : GSimpleActionGroup

// Blacklisted : GTlsCertificate

// Blacklisted : GTlsConnection
