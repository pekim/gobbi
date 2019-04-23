// This is a generated file - DO NOT EDIT
// +build gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

	void applaunchcontext_launchFailedHandler(GObject *, gchar*, gpointer);

	static gulong AppLaunchContext_signal_connect_launch_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "launch-failed", G_CALLBACK(applaunchcontext_launchFailedHandler), data);
	}

*/
/*

	void applaunchcontext_launchedHandler(GObject *, GAppInfo *, GVariant *, gpointer);

	static gulong AppLaunchContext_signal_connect_launched(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "launched", G_CALLBACK(applaunchcontext_launchedHandler), data);
	}

*/
import "C"

type signalAppLaunchContextLaunchFailedDetail struct {
	callback  AppLaunchContextSignalLaunchFailedCallback
	handlerID C.gulong
}

var signalAppLaunchContextLaunchFailedId int
var signalAppLaunchContextLaunchFailedMap = make(map[int]signalAppLaunchContextLaunchFailedDetail)
var signalAppLaunchContextLaunchFailedLock sync.RWMutex

// AppLaunchContextSignalLaunchFailedCallback is a callback function for a 'launch-failed' signal emitted from a AppLaunchContext.
type AppLaunchContextSignalLaunchFailedCallback func(startupNotifyId string)

/*
ConnectLaunchFailed connects the callback to the 'launch-failed' signal for the AppLaunchContext.

The returned value represents the connection, and may be passed to DisconnectLaunchFailed to remove it.
*/
func (recv *AppLaunchContext) ConnectLaunchFailed(callback AppLaunchContextSignalLaunchFailedCallback) int {
	signalAppLaunchContextLaunchFailedLock.Lock()
	defer signalAppLaunchContextLaunchFailedLock.Unlock()

	signalAppLaunchContextLaunchFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AppLaunchContext_signal_connect_launch_failed(instance, C.gpointer(uintptr(signalAppLaunchContextLaunchFailedId)))

	detail := signalAppLaunchContextLaunchFailedDetail{callback, handlerID}
	signalAppLaunchContextLaunchFailedMap[signalAppLaunchContextLaunchFailedId] = detail

	return signalAppLaunchContextLaunchFailedId
}

/*
DisconnectLaunchFailed disconnects a callback from the 'launch-failed' signal for the AppLaunchContext.

The connectionID should be a value returned from a call to ConnectLaunchFailed.
*/
func (recv *AppLaunchContext) DisconnectLaunchFailed(connectionID int) {
	signalAppLaunchContextLaunchFailedLock.Lock()
	defer signalAppLaunchContextLaunchFailedLock.Unlock()

	detail, exists := signalAppLaunchContextLaunchFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAppLaunchContextLaunchFailedMap, connectionID)
}

//export applaunchcontext_launchFailedHandler
func applaunchcontext_launchFailedHandler(_ *C.GObject, c_startup_notify_id *C.gchar, data C.gpointer) {
	signalAppLaunchContextLaunchFailedLock.RLock()
	defer signalAppLaunchContextLaunchFailedLock.RUnlock()

	startupNotifyId := C.GoString(c_startup_notify_id)

	index := int(uintptr(data))
	callback := signalAppLaunchContextLaunchFailedMap[index].callback
	callback(startupNotifyId)
}

type signalAppLaunchContextLaunchedDetail struct {
	callback  AppLaunchContextSignalLaunchedCallback
	handlerID C.gulong
}

var signalAppLaunchContextLaunchedId int
var signalAppLaunchContextLaunchedMap = make(map[int]signalAppLaunchContextLaunchedDetail)
var signalAppLaunchContextLaunchedLock sync.RWMutex

// AppLaunchContextSignalLaunchedCallback is a callback function for a 'launched' signal emitted from a AppLaunchContext.
type AppLaunchContextSignalLaunchedCallback func(info *AppInfo, platformData *glib.Variant)

/*
ConnectLaunched connects the callback to the 'launched' signal for the AppLaunchContext.

The returned value represents the connection, and may be passed to DisconnectLaunched to remove it.
*/
func (recv *AppLaunchContext) ConnectLaunched(callback AppLaunchContextSignalLaunchedCallback) int {
	signalAppLaunchContextLaunchedLock.Lock()
	defer signalAppLaunchContextLaunchedLock.Unlock()

	signalAppLaunchContextLaunchedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AppLaunchContext_signal_connect_launched(instance, C.gpointer(uintptr(signalAppLaunchContextLaunchedId)))

	detail := signalAppLaunchContextLaunchedDetail{callback, handlerID}
	signalAppLaunchContextLaunchedMap[signalAppLaunchContextLaunchedId] = detail

	return signalAppLaunchContextLaunchedId
}

/*
DisconnectLaunched disconnects a callback from the 'launched' signal for the AppLaunchContext.

The connectionID should be a value returned from a call to ConnectLaunched.
*/
func (recv *AppLaunchContext) DisconnectLaunched(connectionID int) {
	signalAppLaunchContextLaunchedLock.Lock()
	defer signalAppLaunchContextLaunchedLock.Unlock()

	detail, exists := signalAppLaunchContextLaunchedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAppLaunchContextLaunchedMap, connectionID)
}

//export applaunchcontext_launchedHandler
func applaunchcontext_launchedHandler(_ *C.GObject, c_info *C.GAppInfo, c_platform_data *C.GVariant, data C.gpointer) {
	signalAppLaunchContextLaunchedLock.RLock()
	defer signalAppLaunchContextLaunchedLock.RUnlock()

	info := AppInfoNewFromC(unsafe.Pointer(c_info))

	platformData := glib.VariantNewFromC(unsafe.Pointer(c_platform_data))

	index := int(uintptr(data))
	callback := signalAppLaunchContextLaunchedMap[index].callback
	callback(info, platformData)
}

// Blacklisted : g_application_command_line_create_file_for_arg

// Blacklisted : g_credentials_get_unix_pid

// Blacklisted : g_desktop_app_info_get_boolean

// Blacklisted : g_desktop_app_info_get_string

// Blacklisted : g_desktop_app_info_has_key

// Blacklisted : g_file_enumerator_get_child

// Blacklisted : g_file_info_get_deletion_date

// Blacklisted : g_memory_output_stream_new_resizable

// g_simple_proxy_resolver_new : unsupported parameter ignore_hosts : in string with indirection level of 2
// Blacklisted : g_simple_proxy_resolver_set_default_proxy

// Unsupported : g_simple_proxy_resolver_set_ignore_hosts : unsupported parameter ignore_hosts : in string with indirection level of 2

// Blacklisted : g_simple_proxy_resolver_set_uri_proxy

// Blacklisted : g_socket_get_option

// Blacklisted : g_socket_set_option

// Blacklisted : g_socket_client_get_proxy_resolver

// Blacklisted : g_socket_client_set_proxy_resolver

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_task_is_valid

// g_task_report_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_task_report_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_task_attach_source : unsupported parameter callback : no type generator for GLib.SourceFunc (GSourceFunc) for param callback

// Blacklisted : g_task_get_cancellable

// Blacklisted : g_task_get_check_cancellable

// Blacklisted : g_task_get_context

// Blacklisted : g_task_get_priority

// Blacklisted : g_task_get_return_on_cancel

// Blacklisted : g_task_get_source_object

// Unsupported : g_task_get_source_tag : no return generator

// Unsupported : g_task_get_task_data : no return generator

// Blacklisted : g_task_had_error

// Blacklisted : g_task_propagate_boolean

// Blacklisted : g_task_propagate_int

// Unsupported : g_task_propagate_pointer : no return generator

// Blacklisted : g_task_return_boolean

// Blacklisted : g_task_return_error

// Blacklisted : g_task_return_error_if_cancelled

// Blacklisted : g_task_return_int

// Blacklisted : g_task_return_new_error

// Unsupported : g_task_return_pointer : unsupported parameter result : no type generator for gpointer (gpointer) for param result

// Unsupported : g_task_run_in_thread : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// Unsupported : g_task_run_in_thread_sync : unsupported parameter task_func : no type generator for TaskThreadFunc (GTaskThreadFunc) for param task_func

// Blacklisted : g_task_set_check_cancellable

// Blacklisted : g_task_set_priority

// Blacklisted : g_task_set_return_on_cancel

// Unsupported : g_task_set_source_tag : unsupported parameter source_tag : no type generator for gpointer (gpointer) for param source_tag

// Unsupported : g_task_set_task_data : unsupported parameter task_data : no type generator for gpointer (gpointer) for param task_data
