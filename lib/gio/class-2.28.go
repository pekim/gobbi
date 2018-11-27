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

type signalApplicationCommandLineDetail struct {
	callback  ApplicationSignalCommandLineCallback
	handlerID C.gulong
}

var signalApplicationCommandLineId int
var signalApplicationCommandLineMap = make(map[int]signalApplicationCommandLineDetail)
var signalApplicationCommandLineLock sync.Mutex

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
var signalApplicationOpenLock sync.Mutex

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

// ApplicationNew is a wrapper around the C function g_application_new.
func ApplicationNew(applicationId string, flags ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.g_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
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

// Unsupported : g_application_command_line_get_arguments : no return type

// GetCwd is a wrapper around the C function g_application_command_line_get_cwd.
func (recv *ApplicationCommandLine) GetCwd() string {
	retC := C.g_application_command_line_get_cwd((*C.GApplicationCommandLine)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_get_environ : no return type

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

// Unsupported : g_application_command_line_get_platform_data : return type : Blacklisted record : GVariant

// Getenv is a wrapper around the C function g_application_command_line_getenv.
func (recv *ApplicationCommandLine) Getenv(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_application_command_line_getenv((*C.GApplicationCommandLine)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_application_command_line_print : unsupported parameter ... : varargs

// Unsupported : g_application_command_line_printerr : unsupported parameter ... : varargs

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

// Unsupported : g_io_stream_splice_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_settings_get_range : return type : Blacklisted record : GVariant

// Unsupported : g_settings_range_check : unsupported parameter value : Blacklisted record : GVariant

// Unsupported signal 'activate' for SimpleAction : unsupported parameter parameter : type GLib.Variant : Blacklisted record : GVariant

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

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

// SimpleActionGroupNew is a wrapper around the C function g_simple_action_group_new.
func SimpleActionGroupNew() *SimpleActionGroup {
	retC := C.g_simple_action_group_new()
	retGo := SimpleActionGroupNewFromC(unsafe.Pointer(retC))

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

// TlsCertificateNewFromFile is a wrapper around the C function g_tls_certificate_new_from_file.
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

// TlsCertificateNewFromFiles is a wrapper around the C function g_tls_certificate_new_from_files.
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

// TlsCertificateNewFromPem is a wrapper around the C function g_tls_certificate_new_from_pem.
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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_tls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// HandshakeFinish is a wrapper around the C function g_tls_connection_handshake_finish.
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
