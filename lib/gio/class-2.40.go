// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	gint application_handleLocalOptionsHandler(GObject *, GVariantDict *, gpointer);

	static gulong Application_signal_connect_handle_local_options(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "handle-local-options", G_CALLBACK(application_handleLocalOptionsHandler), data);
	}

*/
import "C"

// Blacklisted : GAppInfoMonitor

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

// Blacklisted : GNotification

// Blacklisted : GSubprocess

// Blacklisted : GSubprocessLauncher
