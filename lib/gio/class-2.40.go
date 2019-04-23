// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	void appinfomonitor_changedHandler(GObject *, gpointer);

	static gulong AppInfoMonitor_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(appinfomonitor_changedHandler), data);
	}

*/
/*

	gint application_handleLocalOptionsHandler(GObject *, GVariantDict *, gpointer);

	static gulong Application_signal_connect_handle_local_options(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "handle-local-options", G_CALLBACK(application_handleLocalOptionsHandler), data);
	}

*/
/*

	static gboolean _g_output_stream_printf(GOutputStream* stream, gsize* bytes_written, GCancellable* cancellable, GError** error, const gchar* format) {
		return g_output_stream_printf(stream, bytes_written, cancellable, error, format);
    }
*/
import "C"

// AppInfoMonitor is a wrapper around the C record GAppInfoMonitor.
type AppInfoMonitor struct {
	native *C.GAppInfoMonitor
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	c := (*C.GAppInfoMonitor)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppInfoMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppInfoMonitor with another AppInfoMonitor, and returns true if they represent the same GObject.
func (recv *AppInfoMonitor) Equals(other *AppInfoMonitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AppInfoMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AppInfoMonitor.
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
var signalAppInfoMonitorChangedLock sync.RWMutex

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
	signalAppInfoMonitorChangedLock.RLock()
	defer signalAppInfoMonitorChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalAppInfoMonitorChangedMap[index].callback
	callback()
}

// Blacklisted : g_app_info_monitor_get

type signalApplicationHandleLocalOptionsDetail struct {
	callback  ApplicationSignalHandleLocalOptionsCallback
	handlerID C.gulong
}

var signalApplicationHandleLocalOptionsId int
var signalApplicationHandleLocalOptionsMap = make(map[int]signalApplicationHandleLocalOptionsDetail)
var signalApplicationHandleLocalOptionsLock sync.RWMutex

// ApplicationSignalHandleLocalOptionsCallback is a callback function for a 'handle-local-options' signal emitted from a Application.
type ApplicationSignalHandleLocalOptionsCallback func(options *glib.VariantDict) int32

/*
ConnectHandleLocalOptions connects the callback to the 'handle-local-options' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectHandleLocalOptions to remove it.
*/
func (recv *Application) ConnectHandleLocalOptions(callback ApplicationSignalHandleLocalOptionsCallback) int {
	signalApplicationHandleLocalOptionsLock.Lock()
	defer signalApplicationHandleLocalOptionsLock.Unlock()

	signalApplicationHandleLocalOptionsId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_handle_local_options(instance, C.gpointer(uintptr(signalApplicationHandleLocalOptionsId)))

	detail := signalApplicationHandleLocalOptionsDetail{callback, handlerID}
	signalApplicationHandleLocalOptionsMap[signalApplicationHandleLocalOptionsId] = detail

	return signalApplicationHandleLocalOptionsId
}

/*
DisconnectHandleLocalOptions disconnects a callback from the 'handle-local-options' signal for the Application.

The connectionID should be a value returned from a call to ConnectHandleLocalOptions.
*/
func (recv *Application) DisconnectHandleLocalOptions(connectionID int) {
	signalApplicationHandleLocalOptionsLock.Lock()
	defer signalApplicationHandleLocalOptionsLock.Unlock()

	detail, exists := signalApplicationHandleLocalOptionsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationHandleLocalOptionsMap, connectionID)
}

//export application_handleLocalOptionsHandler
func application_handleLocalOptionsHandler(_ *C.GObject, c_options *C.GVariantDict, data C.gpointer) C.gint {
	signalApplicationHandleLocalOptionsLock.RLock()
	defer signalApplicationHandleLocalOptionsLock.RUnlock()

	options := glib.VariantDictNewFromC(unsafe.Pointer(c_options))

	index := int(uintptr(data))
	callback := signalApplicationHandleLocalOptionsMap[index].callback
	retGo := callback(options)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries :

// Blacklisted : g_application_add_option_group

// Blacklisted : g_application_send_notification

// Blacklisted : g_application_withdraw_notification

// Blacklisted : g_application_command_line_get_options_dict

// Blacklisted : g_inet_socket_address_new_from_string

// Notification is a wrapper around the C record GNotification.
type Notification struct {
	native *C.GNotification
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.GNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Notification) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Notification with another Notification, and returns true if they represent the same GObject.
func (recv *Notification) Equals(other *Notification) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Notification) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Notification.
// Exercise care, as this is a potentially dangerous function if the Object is not a Notification.
func CastToNotification(object *gobject.Object) *Notification {
	return NotificationNewFromC(object.ToC())
}

// Blacklisted : g_notification_new

// Blacklisted : g_notification_add_button

// Blacklisted : g_notification_add_button_with_target

// Blacklisted : g_notification_add_button_with_target_value

// Blacklisted : g_notification_set_body

// Blacklisted : g_notification_set_default_action

// Blacklisted : g_notification_set_default_action_and_target

// Blacklisted : g_notification_set_default_action_and_target_value

// Blacklisted : g_notification_set_icon

// Blacklisted : g_notification_set_priority

// SetTitle is a wrapper around the C function g_notification_set_title.
func (recv *Notification) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_notification_set_title((*C.GNotification)(recv.native), c_title)

	return
}

// Blacklisted : g_notification_set_urgent

// Printf is a wrapper around the C function g_output_stream_printf.
func (recv *OutputStream) Printf(cancellable *Cancellable, error *glib.Error, format string, args ...interface{}) (bool, uint64) {
	var c_bytes_written C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	c_error := (**C.GError)(C.NULL)
	if error != nil {
		c_error = (**C.GError)(error.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_output_stream_printf((*C.GOutputStream)(recv.native), &c_bytes_written, c_cancellable, c_error, c_format)
	retGo := retC == C.TRUE

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten
}

// Unsupported : g_output_stream_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// Blacklisted : g_settings_get_default_value

// Blacklisted : g_settings_get_user_value

// Subprocess is a wrapper around the C record GSubprocess.
type Subprocess struct {
	native *C.GSubprocess
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	c := (*C.GSubprocess)(u)
	if c == nil {
		return nil
	}

	g := &Subprocess{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Subprocess) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Subprocess with another Subprocess, and returns true if they represent the same GObject.
func (recv *Subprocess) Equals(other *Subprocess) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Subprocess) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Subprocess.
// Exercise care, as this is a potentially dangerous function if the Object is not a Subprocess.
func CastToSubprocess(object *gobject.Object) *Subprocess {
	return SubprocessNewFromC(object.ToC())
}

// Unsupported : g_subprocess_new : unsupported parameter ... : varargs

// Blacklisted : g_subprocess_newv

// Blacklisted : g_subprocess_communicate

// Unsupported : g_subprocess_communicate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_subprocess_communicate_finish

// Blacklisted : g_subprocess_communicate_utf8

// Unsupported : g_subprocess_communicate_utf8_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_subprocess_communicate_utf8_finish

// Blacklisted : g_subprocess_force_exit

// Blacklisted : g_subprocess_get_exit_status

// Blacklisted : g_subprocess_get_identifier

// Blacklisted : g_subprocess_get_if_exited

// Blacklisted : g_subprocess_get_if_signaled

// Blacklisted : g_subprocess_get_status

// Blacklisted : g_subprocess_get_stderr_pipe

// Blacklisted : g_subprocess_get_stdin_pipe

// Blacklisted : g_subprocess_get_stdout_pipe

// Blacklisted : g_subprocess_get_successful

// Blacklisted : g_subprocess_get_term_sig

// Blacklisted : g_subprocess_send_signal

// Blacklisted : g_subprocess_wait

// Unsupported : g_subprocess_wait_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_subprocess_wait_check

// Unsupported : g_subprocess_wait_check_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_subprocess_wait_check_finish

// Blacklisted : g_subprocess_wait_finish

// SubprocessLauncher is a wrapper around the C record GSubprocessLauncher.
type SubprocessLauncher struct {
	native *C.GSubprocessLauncher
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	c := (*C.GSubprocessLauncher)(u)
	if c == nil {
		return nil
	}

	g := &SubprocessLauncher{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SubprocessLauncher) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SubprocessLauncher with another SubprocessLauncher, and returns true if they represent the same GObject.
func (recv *SubprocessLauncher) Equals(other *SubprocessLauncher) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SubprocessLauncher) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SubprocessLauncher.
// Exercise care, as this is a potentially dangerous function if the Object is not a SubprocessLauncher.
func CastToSubprocessLauncher(object *gobject.Object) *SubprocessLauncher {
	return SubprocessLauncherNewFromC(object.ToC())
}

// Blacklisted : g_subprocess_launcher_new

// Blacklisted : g_subprocess_launcher_getenv

// Unsupported : g_subprocess_launcher_set_child_setup : unsupported parameter child_setup : no type generator for GLib.SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// Blacklisted : g_subprocess_launcher_set_cwd

// Blacklisted : g_subprocess_launcher_set_environ

// Blacklisted : g_subprocess_launcher_set_flags

// Blacklisted : g_subprocess_launcher_set_stderr_file_path

// Blacklisted : g_subprocess_launcher_set_stdin_file_path

// Blacklisted : g_subprocess_launcher_set_stdout_file_path

// Blacklisted : g_subprocess_launcher_setenv

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter ... : varargs

// Blacklisted : g_subprocess_launcher_spawnv

// Blacklisted : g_subprocess_launcher_take_fd

// Blacklisted : g_subprocess_launcher_take_stderr_fd

// Blacklisted : g_subprocess_launcher_take_stdin_fd

// Blacklisted : g_subprocess_launcher_take_stdout_fd

// Blacklisted : g_subprocess_launcher_unsetenv

// Blacklisted : g_tls_interaction_invoke_request_certificate

// Blacklisted : g_tls_interaction_request_certificate

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_tls_interaction_request_certificate_finish
