// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void application_windowAddedHandler(GObject *, GtkWindow *, gpointer);

	static gulong Application_signal_connect_window_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-added", G_CALLBACK(application_windowAddedHandler), data);
	}

*/
/*

	void application_windowRemovedHandler(GObject *, GtkWindow *, gpointer);

	static gulong Application_signal_connect_window_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-removed", G_CALLBACK(application_windowRemovedHandler), data);
	}

*/
import "C"

type signalApplicationWindowAddedDetail struct {
	callback  ApplicationSignalWindowAddedCallback
	handlerID C.gulong
}

var signalApplicationWindowAddedId int
var signalApplicationWindowAddedMap = make(map[int]signalApplicationWindowAddedDetail)
var signalApplicationWindowAddedLock sync.RWMutex

// ApplicationSignalWindowAddedCallback is a callback function for a 'window-added' signal emitted from a Application.
type ApplicationSignalWindowAddedCallback func(window *Window)

/*
ConnectWindowAdded connects the callback to the 'window-added' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectWindowAdded to remove it.
*/
func (recv *Application) ConnectWindowAdded(callback ApplicationSignalWindowAddedCallback) int {
	signalApplicationWindowAddedLock.Lock()
	defer signalApplicationWindowAddedLock.Unlock()

	signalApplicationWindowAddedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_window_added(instance, C.gpointer(uintptr(signalApplicationWindowAddedId)))

	detail := signalApplicationWindowAddedDetail{callback, handlerID}
	signalApplicationWindowAddedMap[signalApplicationWindowAddedId] = detail

	return signalApplicationWindowAddedId
}

/*
DisconnectWindowAdded disconnects a callback from the 'window-added' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowAdded.
*/
func (recv *Application) DisconnectWindowAdded(connectionID int) {
	signalApplicationWindowAddedLock.Lock()
	defer signalApplicationWindowAddedLock.Unlock()

	detail, exists := signalApplicationWindowAddedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationWindowAddedMap, connectionID)
}

//export application_windowAddedHandler
func application_windowAddedHandler(_ *C.GObject, c_window *C.GtkWindow, data C.gpointer) {
	signalApplicationWindowAddedLock.RLock()
	defer signalApplicationWindowAddedLock.RUnlock()

	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := signalApplicationWindowAddedMap[index].callback
	callback(window)
}

type signalApplicationWindowRemovedDetail struct {
	callback  ApplicationSignalWindowRemovedCallback
	handlerID C.gulong
}

var signalApplicationWindowRemovedId int
var signalApplicationWindowRemovedMap = make(map[int]signalApplicationWindowRemovedDetail)
var signalApplicationWindowRemovedLock sync.RWMutex

// ApplicationSignalWindowRemovedCallback is a callback function for a 'window-removed' signal emitted from a Application.
type ApplicationSignalWindowRemovedCallback func(window *Window)

/*
ConnectWindowRemoved connects the callback to the 'window-removed' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectWindowRemoved to remove it.
*/
func (recv *Application) ConnectWindowRemoved(callback ApplicationSignalWindowRemovedCallback) int {
	signalApplicationWindowRemovedLock.Lock()
	defer signalApplicationWindowRemovedLock.Unlock()

	signalApplicationWindowRemovedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_window_removed(instance, C.gpointer(uintptr(signalApplicationWindowRemovedId)))

	detail := signalApplicationWindowRemovedDetail{callback, handlerID}
	signalApplicationWindowRemovedMap[signalApplicationWindowRemovedId] = detail

	return signalApplicationWindowRemovedId
}

/*
DisconnectWindowRemoved disconnects a callback from the 'window-removed' signal for the Application.

The connectionID should be a value returned from a call to ConnectWindowRemoved.
*/
func (recv *Application) DisconnectWindowRemoved(connectionID int) {
	signalApplicationWindowRemovedLock.Lock()
	defer signalApplicationWindowRemovedLock.Unlock()

	detail, exists := signalApplicationWindowRemovedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationWindowRemovedMap, connectionID)
}

//export application_windowRemovedHandler
func application_windowRemovedHandler(_ *C.GObject, c_window *C.GtkWindow, data C.gpointer) {
	signalApplicationWindowRemovedLock.RLock()
	defer signalApplicationWindowRemovedLock.RUnlock()

	window := WindowNewFromC(unsafe.Pointer(c_window))

	index := int(uintptr(data))
	callback := signalApplicationWindowRemovedMap[index].callback
	callback(window)
}

// Blacklisted : gtk_container_child_notify

// Blacklisted : gtk_drag_source_set_icon_gicon

// Blacklisted : gtk_widget_has_visible_focus

// Blacklisted : gtk_window_get_focus_visible

// Blacklisted : gtk_window_set_focus_visible
