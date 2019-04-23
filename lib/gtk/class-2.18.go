// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void label_activateCurrentLinkHandler(GObject *, gpointer);

	static gulong Label_signal_connect_activate_current_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-current-link", G_CALLBACK(label_activateCurrentLinkHandler), data);
	}

*/
/*

	gboolean label_activateLinkHandler(GObject *, gchar*, gpointer);

	static gulong Label_signal_connect_activate_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-link", G_CALLBACK(label_activateLinkHandler), data);
	}

*/
import "C"

type signalLabelActivateCurrentLinkDetail struct {
	callback  LabelSignalActivateCurrentLinkCallback
	handlerID C.gulong
}

var signalLabelActivateCurrentLinkId int
var signalLabelActivateCurrentLinkMap = make(map[int]signalLabelActivateCurrentLinkDetail)
var signalLabelActivateCurrentLinkLock sync.RWMutex

// LabelSignalActivateCurrentLinkCallback is a callback function for a 'activate-current-link' signal emitted from a Label.
type LabelSignalActivateCurrentLinkCallback func()

/*
ConnectActivateCurrentLink connects the callback to the 'activate-current-link' signal for the Label.

The returned value represents the connection, and may be passed to DisconnectActivateCurrentLink to remove it.
*/
func (recv *Label) ConnectActivateCurrentLink(callback LabelSignalActivateCurrentLinkCallback) int {
	signalLabelActivateCurrentLinkLock.Lock()
	defer signalLabelActivateCurrentLinkLock.Unlock()

	signalLabelActivateCurrentLinkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Label_signal_connect_activate_current_link(instance, C.gpointer(uintptr(signalLabelActivateCurrentLinkId)))

	detail := signalLabelActivateCurrentLinkDetail{callback, handlerID}
	signalLabelActivateCurrentLinkMap[signalLabelActivateCurrentLinkId] = detail

	return signalLabelActivateCurrentLinkId
}

/*
DisconnectActivateCurrentLink disconnects a callback from the 'activate-current-link' signal for the Label.

The connectionID should be a value returned from a call to ConnectActivateCurrentLink.
*/
func (recv *Label) DisconnectActivateCurrentLink(connectionID int) {
	signalLabelActivateCurrentLinkLock.Lock()
	defer signalLabelActivateCurrentLinkLock.Unlock()

	detail, exists := signalLabelActivateCurrentLinkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLabelActivateCurrentLinkMap, connectionID)
}

//export label_activateCurrentLinkHandler
func label_activateCurrentLinkHandler(_ *C.GObject, data C.gpointer) {
	signalLabelActivateCurrentLinkLock.RLock()
	defer signalLabelActivateCurrentLinkLock.RUnlock()

	index := int(uintptr(data))
	callback := signalLabelActivateCurrentLinkMap[index].callback
	callback()
}

type signalLabelActivateLinkDetail struct {
	callback  LabelSignalActivateLinkCallback
	handlerID C.gulong
}

var signalLabelActivateLinkId int
var signalLabelActivateLinkMap = make(map[int]signalLabelActivateLinkDetail)
var signalLabelActivateLinkLock sync.RWMutex

// LabelSignalActivateLinkCallback is a callback function for a 'activate-link' signal emitted from a Label.
type LabelSignalActivateLinkCallback func(uri string) bool

/*
ConnectActivateLink connects the callback to the 'activate-link' signal for the Label.

The returned value represents the connection, and may be passed to DisconnectActivateLink to remove it.
*/
func (recv *Label) ConnectActivateLink(callback LabelSignalActivateLinkCallback) int {
	signalLabelActivateLinkLock.Lock()
	defer signalLabelActivateLinkLock.Unlock()

	signalLabelActivateLinkId++
	instance := C.gpointer(recv.native)
	handlerID := C.Label_signal_connect_activate_link(instance, C.gpointer(uintptr(signalLabelActivateLinkId)))

	detail := signalLabelActivateLinkDetail{callback, handlerID}
	signalLabelActivateLinkMap[signalLabelActivateLinkId] = detail

	return signalLabelActivateLinkId
}

/*
DisconnectActivateLink disconnects a callback from the 'activate-link' signal for the Label.

The connectionID should be a value returned from a call to ConnectActivateLink.
*/
func (recv *Label) DisconnectActivateLink(connectionID int) {
	signalLabelActivateLinkLock.Lock()
	defer signalLabelActivateLinkLock.Unlock()

	detail, exists := signalLabelActivateLinkMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalLabelActivateLinkMap, connectionID)
}

//export label_activateLinkHandler
func label_activateLinkHandler(_ *C.GObject, c_uri *C.gchar, data C.gpointer) C.gboolean {
	signalLabelActivateLinkLock.RLock()
	defer signalLabelActivateLinkLock.RUnlock()

	uri := C.GoString(c_uri)

	index := int(uintptr(data))
	callback := signalLabelActivateLinkMap[index].callback
	retGo := callback(uri)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted : gtk_label_get_current_uri

// Blacklisted : gtk_label_get_track_visited_links

// Blacklisted : gtk_label_set_track_visited_links

// Blacklisted : gtk_widget_get_allocation

// Blacklisted : gtk_widget_get_app_paintable

// Blacklisted : gtk_widget_get_can_default

// Blacklisted : gtk_widget_get_can_focus

// Blacklisted : gtk_widget_get_double_buffered

// Blacklisted : gtk_widget_get_has_window

// Blacklisted : gtk_widget_get_receives_default

// Blacklisted : gtk_widget_get_sensitive

// Blacklisted : gtk_widget_get_state

// Blacklisted : gtk_widget_get_visible

// Blacklisted : gtk_widget_has_default

// Blacklisted : gtk_widget_has_focus

// Blacklisted : gtk_widget_has_grab

// Blacklisted : gtk_widget_is_drawable

// Blacklisted : gtk_widget_is_sensitive

// Blacklisted : gtk_widget_is_toplevel

// Blacklisted : gtk_widget_set_allocation

// Blacklisted : gtk_widget_set_can_default

// Blacklisted : gtk_widget_set_can_focus

// Blacklisted : gtk_widget_set_has_window

// Blacklisted : gtk_widget_set_receives_default

// Blacklisted : gtk_widget_set_visible

// Blacklisted : gtk_widget_set_window
